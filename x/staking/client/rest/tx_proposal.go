package rest

import (
	"github.com/exfury/grbchain/libs/cosmos-sdk/client/context"
	govRest "github.com/exfury/grbchain/x/gov/client/rest"
)

// ProposeValidatorProposalRESTHandler defines propose validator proposal handler
func ProposeValidatorProposalRESTHandler(context.CLIContext) govRest.ProposalRESTHandler {
	return govRest.ProposalRESTHandler{}
}
