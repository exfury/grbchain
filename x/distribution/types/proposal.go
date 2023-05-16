package types

import (
	"fmt"
	"github.com/exfury/grbchain/libs/system"
	"strings"

	sdk "github.com/exfury/grbchain/libs/cosmos-sdk/types"

	govtypes "github.com/exfury/grbchain/x/gov/types"
)

const (
	// ProposalTypeCommunityPoolSpend defines the type for a CommunityPoolSpendProposal
	ProposalTypeCommunityPoolSpend = "CommunityPoolSpend"
)

// Assert CommunityPoolSpendProposal implements govtypes.Content at compile-time
var _ govtypes.Content = CommunityPoolSpendProposal{}

func init() {
	govtypes.RegisterProposalType(ProposalTypeCommunityPoolSpend)
	govtypes.RegisterProposalType(ProposalTypeChangeDistributionType)
	govtypes.RegisterProposalType(ProposalTypeWithdrawRewardEnabled)
	govtypes.RegisterProposalType(ProposalTypeRewardTruncatePrecision)
	govtypes.RegisterProposalTypeCodec(CommunityPoolSpendProposal{}, system.Chain+"/distribution/CommunityPoolSpendProposal")
	govtypes.RegisterProposalTypeCodec(ChangeDistributionTypeProposal{}, system.Chain+"/distribution/ChangeDistributionTypeProposal")
	govtypes.RegisterProposalTypeCodec(WithdrawRewardEnabledProposal{}, system.Chain+"/distribution/WithdrawRewardEnabledProposal")
	govtypes.RegisterProposalTypeCodec(RewardTruncatePrecisionProposal{}, system.Chain+"/distribution/RewardTruncatePrecisionProposal")
}

// CommunityPoolSpendProposal spends from the community pool
type CommunityPoolSpendProposal struct {
	Title       string         `json:"title" yaml:"title"`
	Description string         `json:"description" yaml:"description"`
	Recipient   sdk.AccAddress `json:"recipient" yaml:"recipient"`
	Amount      sdk.SysCoins   `json:"amount" yaml:"amount"`
}

// NewCommunityPoolSpendProposal creates a new community pool spned proposal.
func NewCommunityPoolSpendProposal(title, description string, recipient sdk.AccAddress, amount sdk.SysCoins) CommunityPoolSpendProposal {
	return CommunityPoolSpendProposal{title, description, recipient, amount}
}

// GetTitle returns the title of a community pool spend proposal.
func (csp CommunityPoolSpendProposal) GetTitle() string { return csp.Title }

// GetDescription returns the description of a community pool spend proposal.
func (csp CommunityPoolSpendProposal) GetDescription() string { return csp.Description }

// GetDescription returns the routing key of a community pool spend proposal.
func (csp CommunityPoolSpendProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns the type of a community pool spend proposal.
func (csp CommunityPoolSpendProposal) ProposalType() string { return ProposalTypeCommunityPoolSpend }

// ValidateBasic runs basic stateless validity checks
func (csp CommunityPoolSpendProposal) ValidateBasic() error {
	err := govtypes.ValidateAbstract(ModuleName, csp)
	if err != nil {
		return err
	}
	if !csp.Amount.IsValid() {
		return ErrInvalidProposalAmount()
	}
	if csp.Recipient.Empty() {
		return ErrEmptyProposalRecipient()
	}
	return nil
}

// String implements the Stringer interface.
func (csp CommunityPoolSpendProposal) String() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf(`Community Pool Spend Proposal:
  Title:       %s
  Description: %s
  Recipient:   %s
  Amount:      %s
`, csp.Title, csp.Description, csp.Recipient, csp.Amount))
	return b.String()
}