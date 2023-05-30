package oracle_test

import (
	"testing"

	simapp "github.com/cascadiafoundation/cascadia/app"
	"github.com/cascadiafoundation/cascadia/testutil/nullify"
	feemarkettypes "github.com/cascadiafoundation/cascadia/x/feemarket/types"
	"github.com/cascadiafoundation/cascadia/x/oracle"
	"github.com/cascadiafoundation/cascadia/x/oracle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

const (
	initChain = true
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		AssetInfos: []types.AssetInfo{
			{
				Denom:   "satoshi",
				Display: "BTC",
			},
			{
				Denom:   "wei",
				Display: "ETH",
			},
		},
		Prices: []types.Price{
			{
				Asset: "BTC",
				Price: sdk.NewDec(30000),
			},
			{
				Asset: "ETH",
				Price: sdk.NewDec(2000),
			},
		},
		PriceFeeders: []types.PriceFeeder{
			{
				Feeder:   "elys10d07y265gmmuvt4z0w9aw880jnsr700j6z2zm3",
				IsActive: true,
			},
			{
				Feeder:   "elys12tzylat4udvjj56uuhu3vj2n4vgp7cf9fwna9w",
				IsActive: false,
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	isCheckTx := false
	app := simapp.Setup(isCheckTx, feemarkettypes.DefaultGenesisState())
	ctx := app.BaseApp.NewContext(initChain, tmproto.Header{})
	oracle.InitGenesis(ctx, app.OracleKeeper, genesisState)
	got := oracle.ExportGenesis(ctx, app.OracleKeeper)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	require.ElementsMatch(t, genesisState.AssetInfos, got.AssetInfos)
	require.ElementsMatch(t, genesisState.Prices, got.Prices)
	require.ElementsMatch(t, genesisState.PriceFeeders, got.PriceFeeders)
	// this line is used by starport scaffolding # genesis/test/assert
}
