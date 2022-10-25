package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "honor/testutil/keeper"
	"honor/x/honor/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.HonorKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
