package keeper

import (
	"cosmossdk.io/errors"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"

	"github.com/arkeonetwork/arkeo/common"
	"github.com/arkeonetwork/arkeo/common/cosmos"
	"github.com/arkeonetwork/arkeo/x/arkeo/configs"
	"github.com/arkeonetwork/arkeo/x/arkeo/types"
)

type Manager struct {
	keeper  Keeper
	sk      stakingkeeper.Keeper
	configs configs.ConfigValues
}

func NewManager(k Keeper, sk stakingkeeper.Keeper) Manager {
	ver := k.GetVersion()
	return Manager{
		keeper:  k,
		sk:      sk,
		configs: configs.GetConfigValues(ver),
	}
}

func (mgr Manager) EndBlock(ctx cosmos.Context) error {
	if err := mgr.ContractEndBlock(ctx); err != nil {
		ctx.Logger().Error("unable to settle contracts", "error", err)
	}
	if err := mgr.ValidatorEndBlock(ctx); err != nil {
		ctx.Logger().Error("unable to settle contracts", "error", err)
	}
	return nil
}

func (mgr Manager) ContractEndBlock(ctx cosmos.Context) error {
	set, err := mgr.keeper.GetContractExpirationSet(ctx, ctx.BlockHeight())
	if err != nil {
		return err
	}

	for _, exp := range set.Contracts {
		contract, err := mgr.keeper.GetContract(ctx, exp.ProviderPubKey, exp.Chain, exp.Client)
		if err != nil {
			ctx.Logger().Error("unable to fetch contract", "pubkey", exp.ProviderPubKey, "chain", exp.Chain, "client", exp.Client, "error", err)
			continue
		}
		_, err = mgr.SettleContract(ctx, contract, 0, true)
		if err != nil {
			ctx.Logger().Error("unable settle contract", "pubkey", exp.ProviderPubKey, "chain", exp.Chain, "client", exp.Client, "error", err)
			continue
		}
	}

	return nil
}

// This function pays out rewards to validators.
// TODO: the method of accomplishing this is admittedly quite inefficient. The
// better approach would be to track live allocation via assigning "units" to
// validators when they bond and unbond. The math for this is as follows
// U = total bond units
// T = tokens bonded
// t = new tokens being bonded
// units = U / (T / t)
// Since the development goal at the moment is to get this chain up and
// running, we can save this optimization for another day.
func (mgr Manager) ValidatorEndBlock(ctx cosmos.Context) error {
	valCycle := mgr.FetchConfig(ctx, configs.ValidatorPayoutCycle)
	if valCycle == 0 || ctx.BlockHeight()%valCycle != 0 {
		return nil
	}
	validators := mgr.sk.GetBondedValidatorsByPower(ctx)

	reserve := mgr.keeper.GetBalanceOfModule(ctx, types.ReserveName, configs.Denom)
	emissionCurve := mgr.FetchConfig(ctx, configs.EmissionCurve)
	blocksPerYear := mgr.FetchConfig(ctx, configs.BlocksPerYear)
	blockReward := mgr.calcBlockReward(reserve.Int64(), emissionCurve, (blocksPerYear / valCycle))

	if blockReward.IsZero() {
		ctx.Logger().Info("no validator rewards this block")
		return nil
	}

	// sum tokens
	total := cosmos.ZeroInt()
	for _, val := range validators {
		if !val.IsBonded() || val.IsJailed() {
			continue
		}
		total = total.Add(val.DelegatorShares.RoundInt())
	}

	for _, val := range validators {
		if !val.IsBonded() || val.IsJailed() {
			ctx.Logger().Info("validator rewards skipped due to status or jailed", "validator", val.String())
			continue
		}

		acc := cosmos.AccAddress(val.GetOperator())

		totalReward := common.GetSafeShare(val.DelegatorShares.RoundInt(), total, blockReward)
		validatorReward := cosmos.ZeroInt()
		rateBasisPts := val.Commission.CommissionRates.Rate.MulInt64(100).RoundInt()

		delegates := mgr.sk.GetValidatorDelegations(ctx, val.GetOperator())
		for _, delegate := range delegates {
			delegateAcc, err := cosmos.AccAddressFromBech32(delegate.DelegatorAddress)
			if err != nil {
				ctx.Logger().Error("unable to fetch delegate address", "delegate", delegate.DelegatorAddress, "error", err)
				continue
			}
			delegateReward := common.GetSafeShare(delegate.Shares.RoundInt(), val.DelegatorShares.RoundInt(), totalReward)
			if acc.String() != delegate.DelegatorAddress {
				valFee := common.GetSafeShare(rateBasisPts, cosmos.NewInt(configs.MaxBasisPoints), delegateReward)
				delegateReward = delegateReward.Sub(valFee)
				validatorReward = validatorReward.Add(valFee)
			}
			if err := mgr.keeper.SendFromModuleToAccount(ctx, types.ReserveName, delegateAcc, getCoins(delegateReward.Int64())); err != nil {
				ctx.Logger().Error("unable to pay rewards to delegate", "delegate", delegate.DelegatorAddress, "error", err)
			}
			ctx.Logger().Info("delegate rewarded", "delegate", delegateAcc.String(), "amount", delegateReward)
		}

		if !validatorReward.IsZero() {
			if err := mgr.keeper.SendFromModuleToAccount(ctx, types.ReserveName, acc, getCoins(validatorReward.Int64())); err != nil {
				ctx.Logger().Error("unable to pay rewards to validator", "validator", val.OperatorAddress, "error", err)
				continue
			}
			ctx.Logger().Info("validator additional rewards", "validator", acc.String(), "amount", validatorReward)
		}

		mgr.ValidatorPayoutEvent(ctx, acc, validatorReward)
	}

	return nil
}

func (mgr Manager) calcBlockReward(totalReserve, emissionCurve, blocksPerYear int64) cosmos.Int {
	// Block Rewards will take the latest reserve, divide it by the emission
	// curve factor, then divide by blocks per year
	if emissionCurve == 0 || blocksPerYear == 0 {
		return cosmos.ZeroInt()
	}
	trD := cosmos.NewDec(totalReserve)
	ecD := cosmos.NewDec(emissionCurve)
	bpyD := cosmos.NewDec(blocksPerYear)
	return trD.Quo(ecD).Quo(bpyD).RoundInt()
}

func (mgr Manager) FetchConfig(ctx cosmos.Context, name configs.ConfigName) int64 {
	// TODO: create a handler for admins to be able to change configs on the
	// fly and check them here before returning
	return mgr.configs.GetInt64Value(name)
}

// any owed debt is paid to data provider
func (mgr Manager) SettleContract(ctx cosmos.Context, contract types.Contract, nonce int64, closed bool) (types.Contract, error) {
	if nonce > contract.Nonce {
		contract.Nonce = nonce
	}
	totalDebt, err := mgr.contractDebt(ctx, contract)
	valIncome := common.GetSafeShare(cosmos.NewInt(mgr.FetchConfig(ctx, configs.ReserveTax)), cosmos.NewInt(configs.MaxBasisPoints), totalDebt)
	debt := totalDebt.Sub(valIncome)
	if err != nil {
		return contract, err
	}
	if !debt.IsZero() {
		provider, err := contract.ProviderPubKey.GetMyAddress()
		if err != nil {
			return contract, err
		}
		if err := mgr.keeper.SendFromModuleToAccount(ctx, types.ContractName, provider, cosmos.NewCoins(cosmos.NewCoin(configs.Denom, debt))); err != nil {
			return contract, err
		}
		if err := mgr.keeper.SendFromModuleToModule(ctx, types.ContractName, types.ReserveName, cosmos.NewCoins(cosmos.NewCoin(configs.Denom, valIncome))); err != nil {
			return contract, err
		}
	}

	contract.Paid = contract.Paid.Add(totalDebt)
	if closed {
		remainder := contract.Deposit.Sub(contract.Paid)
		if !remainder.IsZero() {
			client, err := contract.Client.GetMyAddress()
			if err != nil {
				return contract, err
			}
			if err := mgr.keeper.SendFromModuleToAccount(ctx, types.ContractName, client, cosmos.NewCoins(cosmos.NewCoin(configs.Denom, remainder))); err != nil {
				return contract, err
			}
		}
		contract.ClosedHeight = ctx.BlockHeight()
	}

	err = mgr.keeper.SetContract(ctx, contract)
	if err != nil {
		return contract, err
	}

	mgr.ContractSettlementEvent(ctx, debt, valIncome, contract)
	return contract, nil
}

func (mgr Manager) contractDebt(ctx cosmos.Context, contract types.Contract) (cosmos.Int, error) {
	var debt cosmos.Int
	switch contract.Type {
	case types.ContractType_SUBSCRIPTION:
		debt = cosmos.NewInt(contract.Rate * (ctx.BlockHeight() - contract.Height)).Sub(contract.Paid)
	case types.ContractType_PAY_AS_YOU_GO:
		debt = cosmos.NewInt(contract.Rate * contract.Nonce).Sub(contract.Paid)
	default:
		return cosmos.ZeroInt(), errors.Wrapf(types.ErrInvalidContractType, "%s", contract.Type.String())
	}

	if debt.IsNegative() {
		return cosmos.ZeroInt(), nil
	}

	// sanity check, ensure provider cannot take more than deposited into the contract
	if contract.Paid.Add(debt).GT(contract.Deposit) {
		return contract.Deposit.Sub(contract.Paid), nil
	}

	return debt, nil
}
