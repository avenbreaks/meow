package meow_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "meow/testutil/keeper"
	"meow/testutil/nullify"
	"meow/x/meow"
	"meow/x/meow/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.MeowKeeper(t)
	meow.InitGenesis(ctx, *k, genesisState)
	got := meow.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
