package keeper_test

import (
	"testing"

	testkeeper "chain-demo/testutil/keeper"
	"chain-demo/x/chaindemo/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.ChaindemoKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
