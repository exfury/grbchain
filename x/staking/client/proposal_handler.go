package client

import (
	"github.com/exfury/grbchain/x/staking/client/cli"
	"github.com/exfury/grbchain/x/staking/client/rest"
	govcli "github.com/exfury/grbchain/x/gov/client"
)

var (
	ProposeValidatorProposalHandler = govcli.NewProposalHandler(
		cli.GetCmdProposeValidatorProposal,
		rest.ProposeValidatorProposalRESTHandler,
	)
)

