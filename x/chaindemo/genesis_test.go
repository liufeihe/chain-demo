package chaindemo_test

import (
	"testing"

	keepertest "chain-demo/testutil/keeper"
	"chain-demo/testutil/nullify"
	"chain-demo/x/chaindemo"
	"chain-demo/x/chaindemo/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ChaindemoKeeper(t)
	chaindemo.InitGenesis(ctx, *k, genesisState)
	got := chaindemo.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
