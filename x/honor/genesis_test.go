package honor_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "honor/testutil/keeper"
	"honor/testutil/nullify"
	"honor/x/honor"
	"honor/x/honor/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.HonorKeeper(t)
	honor.InitGenesis(ctx, *k, genesisState)
	got := honor.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
