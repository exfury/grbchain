package client

import (
	"github.com/exfury/grbchain/x/feesplit/client/cli"
	"github.com/exfury/grbchain/x/feesplit/client/rest"
	govcli "github.com/exfury/grbchain/x/gov/client"
)

var (
	// FeeSplitSharesProposalHandler alias gov NewProposalHandler
	FeeSplitSharesProposalHandler = govcli.NewProposalHandler(
		cli.GetCmdFeeSplitSharesProposal,
		rest.FeeSplitSharesProposalRESTHandler,
	)
)
