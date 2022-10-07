package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "meow/testutil/keeper"
	"meow/x/meow/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.MeowKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
