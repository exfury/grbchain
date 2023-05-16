package types

import (
	sdk "github.com/exfury/grbchain/libs/cosmos-sdk/types"
	sdkerrors "github.com/exfury/grbchain/libs/cosmos-sdk/types/errors"
	"github.com/exfury/grbchain/libs/system"
)

// for denom convert wei to okb and reject okb direct
func (m *MsgTransfer) RulesFilter() (sdk.Msg, error) {
	if m.Token.Denom == sdk.DefaultBondDenom {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "ibc MsgTransfer not support "+system.Currency+" denom")
	}
	return m, nil
}
