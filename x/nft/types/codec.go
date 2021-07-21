package types

import (
	// "github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	// // this line is used by starport scaffolding # 1
	// "github.com/cosmos/cosmos-sdk/types/msgservice"
	gogotypes "github.com/gogo/protobuf/types"

	"github.com/cosmos/cosmos-sdk/codec"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"

	"github.com/zebi/cric/x/nft/exported"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	// this line is used by starport scaffolding # 2
	cdc.RegisterConcrete(&MsgIssueDenom{}, "zebi/nft/MsgIssueDenom", nil)
	cdc.RegisterConcrete(&MsgTransferNFT{}, "zebi/nft/MsgTransferNFT", nil)
	cdc.RegisterConcrete(&MsgEditNFT{}, "zebi/nft/MsgEditNFT", nil)
	cdc.RegisterConcrete(&MsgMintNFT{}, "zebi/nft/MsgMintNFT", nil)
	cdc.RegisterConcrete(&MsgBurnNFT{}, "zebi/nft/MsgBurnNFT", nil)

	cdc.RegisterInterface((*exported.NFT)(nil), nil)
	cdc.RegisterConcrete(&BaseNFT{}, "zebi/nft/BaseNFT", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgIssueDenom{},
		&MsgTransferNFT{},
		&MsgEditNFT{},
		&MsgMintNFT{},
		&MsgBurnNFT{},
	)

	registry.RegisterImplementations((*exported.NFT)(nil),
		&BaseNFT{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}

// return supply protobuf code
func MustMarshalSupply(cdc codec.Marshaler, supply uint64) []byte {
	supplyWrap := gogotypes.UInt64Value{Value: supply}
	return cdc.MustMarshalBinaryBare(&supplyWrap)
}

// return th supply
func MustUnMarshalSupply(cdc codec.Marshaler, value []byte) uint64 {
	var supplyWrap gogotypes.UInt64Value
	cdc.MustUnmarshalBinaryBare(value, &supplyWrap)
	return supplyWrap.Value
}

// return the tokenID protobuf code
func MustMarshalTokenID(cdc codec.Marshaler, tokenID string) []byte {
	tokenIDWrap := gogotypes.StringValue{Value: tokenID}
	return cdc.MustMarshalBinaryBare(&tokenIDWrap)
}

// return th tokenID
func MustUnMarshalTokenID(cdc codec.Marshaler, value []byte) string {
	var tokenIDWrap gogotypes.StringValue
	cdc.MustUnmarshalBinaryBare(value, &tokenIDWrap)
	return tokenIDWrap.Value
}
