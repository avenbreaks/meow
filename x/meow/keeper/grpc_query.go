package keeper

import (
	"meow/x/meow/types"
)

var _ types.QueryServer = Keeper{}
