package types

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	errorsTypes "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAcceptOwner = "accept_owner"

var _ sdk.Msg = &MsgAcceptOwner{}

func NewMsgAcceptOwner(from string) *MsgAcceptOwner {
	return &MsgAcceptOwner{
		From: from,
	}
}

func (msg *MsgAcceptOwner) Route() string {
	return RouterKey
}

func (msg *MsgAcceptOwner) Type() string {
	return TypeMsgAcceptOwner
}

func (msg *MsgAcceptOwner) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

func (msg *MsgAcceptOwner) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAcceptOwner) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return errors.Wrapf(errorsTypes.ErrInvalidAddress, "invalid from address (%s)", err)
	}
	return nil
}
