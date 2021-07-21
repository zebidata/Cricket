package types

import (
	// this line is used by starport scaffolding # genesis/types/import
	// this line is used by starport scaffolding # ibc/genesistype/import
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
// func DefaultGenesis() *GenesisState {
// 	return &GenesisState{
// 		// this line is used by starport scaffolding # ibc/genesistype/default
// 		// this line is used by starport scaffolding # genesis/types/default
// 	}
// }
func NewGenesisState(collections []Collection) *GenesisState {
	return &GenesisState{
		Collections: collections,
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
// func (gs GenesisState) Validate() error {
// 	// this line is used by starport scaffolding # ibc/genesistype/validate

// 	// this line is used by starport scaffolding # genesis/types/validate

// 	return nil
// }
func ValidateGenesis(data GenesisState) error {
	for _, c := range data.Collections {
		if err := ValidateDenomID(c.Denom.Name); err != nil {
			return err
		}

		for _, nft := range c.NFTs {
			if nft.GetOwner().Empty() {
				return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing owner")
			}

			if err := ValidateTokenID(nft.GetID()); err != nil {
				return err
			}

			if err := ValidateTokenURI(nft.GetURI()); err != nil {
				return err
			}
		}
	}
	return nil
}
