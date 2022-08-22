package marketplace

import (
	"math/rand"

	"github.com/CudoVentures/cudos-node/testutil/sample"
	marketplacesimulation "github.com/CudoVentures/cudos-node/x/marketplace/simulation"
	"github.com/CudoVentures/cudos-node/x/marketplace/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = marketplacesimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgPublishCollection = "op_weight_msg_publish_collection"
	// TODO: Determine the simulation weight value
	defaultWeightMsgPublishCollection int = 100

	opWeightMsgPublishNft = "op_weight_msg_publish_nft"
	// TODO: Determine the simulation weight value
	defaultWeightMsgPublishNft int = 100

	opWeightMsgBuyNft = "op_weight_msg_buy_nft"
	// TODO: Determine the simulation weight value
	defaultWeightMsgBuyNft int = 100

	opWeightMsgMintNft = "op_weight_msg_mint_nft"
	// TODO: Determine the simulation weight value
	defaultWeightMsgMintNft int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	marketplaceGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&marketplaceGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgPublishCollection int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgPublishCollection, &weightMsgPublishCollection, nil,
		func(_ *rand.Rand) {
			weightMsgPublishCollection = defaultWeightMsgPublishCollection
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgPublishCollection,
		marketplacesimulation.SimulateMsgPublishCollection(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgPublishNft int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgPublishNft, &weightMsgPublishNft, nil,
		func(_ *rand.Rand) {
			weightMsgPublishNft = defaultWeightMsgPublishNft
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgPublishNft,
		marketplacesimulation.SimulateMsgPublishNft(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgBuyNft int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgBuyNft, &weightMsgBuyNft, nil,
		func(_ *rand.Rand) {
			weightMsgBuyNft = defaultWeightMsgBuyNft
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgBuyNft,
		marketplacesimulation.SimulateMsgBuyNft(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgMintNft int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgMintNft, &weightMsgMintNft, nil,
		func(_ *rand.Rand) {
			weightMsgMintNft = defaultWeightMsgMintNft
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgMintNft,
		marketplacesimulation.SimulateMsgMintNft(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
