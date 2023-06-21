package types

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	errorsTypes "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUpdatePauser = "update_pauser"

var _ sdk.Msg = &MsgUpdatePauser{}

func NewMsgUpdatePauser(from string, address string) *MsgUpdatePauser {
	return &MsgUpdatePauser{
		From:    from,
		Address: address,
	}
}

func (msg *MsgUpdatePauser) Route() string {
	return RouterKey
}

func (msg *MsgUpdatePauser) Type() string {
	return TypeMsgUpdatePauser
}

func (msg *MsgUpdatePauser) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

func (msg *MsgUpdatePauser) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdatePauser) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return errors.Wrapf(errorsTypes.ErrInvalidAddress, "invalid from address (%s)", err)
	}
	_, err = sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return errors.Wrapf(errorsTypes.ErrInvalidAddress, "invalid pauser address (%s)", err)
	}
	return nil
}
