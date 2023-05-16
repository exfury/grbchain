package client

import (
	"github.com/exfury/grbchain/libs/cosmos-sdk/x/mint/client/cli"
	"github.com/exfury/grbchain/libs/cosmos-sdk/x/mint/client/rest"
	govcli "github.com/exfury/grbchain/x/gov/client"
)

var (
	ManageTreasuresProposalHandler = govcli.NewProposalHandler(
		cli.GetCmdManageTreasuresProposal,
		rest.ManageTreasuresProposalRESTHandler,
	)

	ExtraProposalHandler = govcli.NewProposalHandler(
		cli.GetCmdExtraProposal,
		rest.ExtraProposalRESTHandler,
	)
)
