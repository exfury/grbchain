package client

import (
	govclient "github.com/exfury/grbchain/libs/cosmos-sdk/x/gov/client"
	"github.com/exfury/grbchain/libs/cosmos-sdk/x/upgrade/client/cli"
	"github.com/exfury/grbchain/libs/cosmos-sdk/x/upgrade/client/rest"
)

var ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitUpgradeProposal, rest.ProposalRESTHandler)
