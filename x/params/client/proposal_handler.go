package client

import (
	govclient "github.com/exfury/grbchain/x/gov/client"
	"github.com/exfury/grbchain/x/params/client/cli"
	"github.com/exfury/grbchain/x/params/client/rest"
)

// ProposalHandler is the param change proposal handler in cmsdk
var ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitProposal, rest.ProposalRESTHandler)

// UpgradeProposalHandler is the upgrade proposal handler
var UpgradeProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitUpgradeProposal, rest.ProposalUpgradeRESTHandler)
