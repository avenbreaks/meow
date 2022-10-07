package keeper

import (
	"meow/x/dex/types"
)

var _ types.QueryServer = Keeper{}
