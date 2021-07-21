package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrInvalidCollection = sdkerrors.Register(ModuleName, 24, "invalid nft collection")
	ErrUnknownCollection = sdkerrors.Register(ModuleName, 23, "unknown nft collection")
	ErrInvalidNFT        = sdkerrors.Register(ModuleName, 14, "invalid nft")
	ErrNFTAlreadyExists  = sdkerrors.Register(ModuleName, 15, "nft already exists")
	ErrUnknownNFT        = sdkerrors.Register(ModuleName, 16, "unknown nft")
	ErrEmptyTokenData    = sdkerrors.Register(ModuleName, 17, "nft data can't be empty")
	ErrUnauthorized      = sdkerrors.Register(ModuleName, 18, "unauthorized address")
	ErrInvalidDenom      = sdkerrors.Register(ModuleName, 19, "invalid denom")
	ErrInvalidTokenID    = sdkerrors.Register(ModuleName, 20, "invalid nft id")
	ErrInvalidTokenURI   = sdkerrors.Register(ModuleName, 21, "invalid nft uri")
	ErrInvalidDenomName  = sdkerrors.Register(ModuleName, 22, "invalid denom name")
)
