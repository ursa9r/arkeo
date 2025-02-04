package keeper

import (
	"context"

	"github.com/arkeonetwork/arkeo/common"
	"github.com/arkeonetwork/arkeo/common/cosmos"
	"github.com/arkeonetwork/arkeo/x/arkeo/configs"
	"github.com/arkeonetwork/arkeo/x/arkeo/types"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CloseContract(goCtx context.Context, msg *types.MsgCloseContract) (*types.MsgCloseContractResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	ctx.Logger().Info(
		"receive MsgCloseContract",
		"pubkey", msg.PubKey,
		"chain", msg.Chain,
		"client", msg.Client,
		"delegate", msg.Delegate,
	)

	cacheCtx, commit := ctx.CacheContext()
	if err := k.CloseContractValidate(cacheCtx, msg); err != nil {
		ctx.Logger().Error("failed close contract validation", "err", err)
		return nil, err
	}

	if err := k.CloseContractHandle(cacheCtx, msg); err != nil {
		ctx.Logger().Error("failed close contract handler", "err", err)
		return nil, err
	}

	commit()
	return &types.MsgCloseContractResponse{}, nil
}

func (k msgServer) CloseContractValidate(ctx cosmos.Context, msg *types.MsgCloseContract) error {
	if k.FetchConfig(ctx, configs.HandlerCloseContract) > 0 {
		return errors.Wrapf(types.ErrDisabledHandler, "close contract")
	}

	chain, err := common.NewChain(msg.Chain)
	if err != nil {
		return err
	}
	contract, err := k.GetContract(ctx, msg.PubKey, chain, msg.FetchSpender())
	if err != nil {
		return err
	}

	// if client is provided, ensure contract client matches msg client
	if len(msg.Client) > 0 {
		if !contract.Client.Equals(msg.Client) {
			return errors.Wrapf(sdkerrors.ErrUnauthorized, "unauthorized contract client")
		}

		if contract.Type == types.ContractType_PAY_AS_YOU_GO {
			// clients are not allowed to cancel a pay-as-you-go contract as it
			// could be a way to game providers. IE, the client make 1,000 requests
			// and before the provider can claim the rewards, the client cancels
			// the contract. We do not want providers to feel "rushed" to claim
			// their rewards or the income is gone.
			return errors.Wrapf(types.ErrCloseContractUnauthorized, "client cannot cancel a pay-as-you-go contract")
		}
	}

	if contract.IsClose(ctx.BlockHeight()) {
		return errors.Wrapf(types.ErrCloseContractAlreadyClosed, "closed %d", contract.Expiration())
	}

	return nil
}

func (k msgServer) CloseContractHandle(ctx cosmos.Context, msg *types.MsgCloseContract) error {
	chain, err := common.NewChain(msg.Chain)
	if err != nil {
		return err
	}
	contract, err := k.GetContract(ctx, msg.PubKey, chain, msg.FetchSpender())
	if err != nil {
		return err
	}

	_, err = k.mgr.SettleContract(ctx, contract, 0, true)
	if err != nil {
		return err
	}

	k.CloseContractEvent(ctx, msg)
	return nil
}
