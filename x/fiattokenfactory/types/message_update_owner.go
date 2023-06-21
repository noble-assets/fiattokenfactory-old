package types

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	errorsTypes "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUpdateOwner = "update_owner"

var _ sdk.Msg = &MsgUpdateOwner{}

func NewMsgUpdateOwner(from string, address string) *MsgUpdateOwner {
	return &MsgUpdateOwner{
		From:    from,
		Address: address,
	}
}

func (msg *MsgUpdateOwner) Route() string {
	return RouterKey
}

func (msg *MsgUpdateOwner) Type() string {
	return TypeMsgUpdateOwner
}

func (msg *MsgUpdateOwner) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

func (msg *MsgUpdateOwner) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateOwner) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return errors.Wrapf(errorsTypes.ErrInvalidAddress, "invalid from address (%s)", err)
	}
	_, err = sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return errors.Wrapf(errorsTypes.ErrInvalidAddress, "invalid owner address (%s)", err)
	}
	return nil
}
