package types

// DONTCOVER

import (
	"cosmossdk.io/errors"
)

// x/arkeo module sentinel errors
var (
	ErrProviderBadSigner                     = errors.Register(ModuleName, 2, "unauthorized: bad provider pubkey and signer")
	ErrProviderAlreadyExists                 = errors.Register(ModuleName, 3, "provider already exists")
	ErrInsufficientFunds                     = errors.Register(ModuleName, 4, "insufficient funds")
	ErrInvalidBond                           = errors.Register(ModuleName, 5, "invalid bond")
	ErrInvalidModProviderMetdataURI          = errors.Register(ModuleName, 6, "invalid mod provider metadata uri")
	ErrInvalidModProviderMaxContractDuration = errors.Register(ModuleName, 7, "invalid mod provider max contract duration")
	ErrInvalidModProviderMinContractDuration = errors.Register(ModuleName, 8, "invalid mod provider min contract duration")
	ErrInvalidModProviderStatus              = errors.Register(ModuleName, 9, "invalid mod provider bad provider status")
	ErrInvalidModProviderNoBond              = errors.Register(ModuleName, 10, "no bond")
	ErrDisabledHandler                       = errors.Register(ModuleName, 11, "disabled handler")
	ErrInvalidChain                          = errors.Register(ModuleName, 12, "invalid chain")
	ErrOpenContractBadProviderStatus         = errors.Register(ModuleName, 13, "provider must have status 'online'")
	ErrOpenContractDuration                  = errors.Register(ModuleName, 14, "invalid contract duration")
	ErrOpenContractMismatchRate              = errors.Register(ModuleName, 15, "mismatch contract rate")
	ErrOpenContractAlreadyOpen               = errors.Register(ModuleName, 16, "contract is already open")
	ErrOpenContractRate                      = errors.Register(ModuleName, 17, "invalid contract rate")
	ErrInvalidContractType                   = errors.Register(ModuleName, 18, "invalid contract type")
	ErrInvalidPubKey                         = errors.Register(ModuleName, 19, "invalid pubkey")
	ErrCloseContractAlreadyClosed            = errors.Register(ModuleName, 20, "contract is already closed")
	ErrCloseContractUnauthorized             = errors.Register(ModuleName, 21, "unauthorized to close contract")
	ErrClaimContractIncomeBadHeight          = errors.Register(ModuleName, 22, "height must be greater than zero")
	ErrClaimContractIncomeBadNonce           = errors.Register(ModuleName, 23, "nonce must be greater than zero")
	ErrClaimContractIncomeClosed             = errors.Register(ModuleName, 24, "contract is closed")
	ErrClaimContractIncomeInvalidSignature   = errors.Register(ModuleName, 25, "invalid signature")
)
