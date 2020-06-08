package main

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestConfigSetup(t *testing.T) {
	// sets up and seals config
	setupConfig()

	// get config
	config := sdk.GetConfig()

	// coin type
	require.Equal(t, uint32(999), config.GetCoinType(), "coin type set")

	// account prefix
	require.Equal(t, "testchain", config.GetBech32AccountAddrPrefix(), "bech 32 account addr prefix")
	require.Equal(t, "testchainpub", config.GetBech32AccountPubPrefix(), "bech 32 account pub prefix")

	// validator prefix
	require.Equal(t, "testchainvaloper", config.GetBech32ValidatorAddrPrefix(), "bech 32 validator addr prefix")
	require.Equal(t, "testchainvaloperpub", config.GetBech32ValidatorPubPrefix(), "bech 32 validator pub prefix")

	// consensus prefix
	require.Equal(t, "testchainvalcons", config.GetBech32ConsensusAddrPrefix(), "bech 32 consensus addr prefix")
	require.Equal(t, "testchainvalconspub", config.GetBech32ConsensusPubPrefix(), "bech 32 consensus pub prefix")

	// config is sealed and can not be modified
	require.Panics(t,
		func() { config.SetBech32PrefixForAccount("someprefix", "someprefixpub") },
		"panics when config is modified (config is sealed)")
}
