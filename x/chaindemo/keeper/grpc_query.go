package keeper

import (
	"chain-demo/x/chaindemo/types"
)

var _ types.QueryServer = Keeper{}
