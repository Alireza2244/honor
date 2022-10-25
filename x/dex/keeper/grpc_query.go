package keeper

import (
	"honor/x/dex/types"
)

var _ types.QueryServer = Keeper{}
