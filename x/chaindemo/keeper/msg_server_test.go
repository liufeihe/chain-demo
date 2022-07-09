package keeper_test

import (
	"context"
	"testing"

	keepertest "chain-demo/testutil/keeper"
	"chain-demo/x/chaindemo/keeper"
	"chain-demo/x/chaindemo/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.ChaindemoKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
