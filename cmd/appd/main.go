package main

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func main() {
	setupConfig()
}

// setupConfig sets up config and seals it
func setupConfig() {
	config := sdk.GetConfig()

	// global coin type for HD wallets
	config.SetCoinType(999)

	// global prefix for bech32 account strings
	basePrefix := "testchain"
	config.SetBech32PrefixForAccount(basePrefix, basePrefix+sdk.PrefixPublic)

	// global prefix for bech32 validator strings
	baseValidatorPrefix := basePrefix + sdk.PrefixValidator + sdk.PrefixOperator
	config.SetBech32PrefixForValidator(baseValidatorPrefix, baseValidatorPrefix+sdk.PrefixPublic)

	// global prefix for bech32 consensus strings
	baseConsensusNodePrefix := basePrefix + sdk.PrefixValidator + sdk.PrefixConsensus
	config.SetBech32PrefixForConsensusNode(baseConsensusNodePrefix, baseConsensusNodePrefix+sdk.PrefixPublic)

	// seal config so it can not be modified
	config.Seal()
}
