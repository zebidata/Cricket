package nft

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zebi/cric/x/nft/keeper"
	"github.com/zebi/cric/x/nft/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init

	// this line is used by starport scaffolding # ibc/genesis/init
	if err := types.ValidateGenesis(genState); err != nil {
		panic(err.Error())
	}

	for _, c := range genState.Collections {
		if err := k.SetDenom(ctx, c.Denom); err != nil {
			panic(err)
		}
		if err := k.SetCollection(ctx, c); err != nil {
			panic(err)
		}
	}
}

// ExportGenesis returns the capability module's exported genesis.
// func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
// 	genesis := types.DefaultGenesis()

// 	// this line is used by starport scaffolding # genesis/module/export

// 	// this line is used by starport scaffolding # ibc/genesis/export

// 	return genesis
// }

// ExportGenesis returns a GenesisState for a given context and keeper.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	return types.NewGenesisState(k.GetCollections(ctx))
}

// DefaultGenesisState returns a default genesis state
func DefaultGenesisState() *types.GenesisState {
	return types.NewGenesisState([]types.Collection{})
}
