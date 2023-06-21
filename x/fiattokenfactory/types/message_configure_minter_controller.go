package types

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	errorsTypes "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgConfigureMinterController = "configure_minter_controller"

var _ sdk.Msg = &MsgConfigureMinterController{}

func NewMsgConfigureMinterController(from string, controller string, minter string) *MsgConfigureMinterController {
	return &MsgConfigureMinterController{
		From:       from,
		Controller: controller,
		Minter:     minter,
	}
}

func (msg *MsgConfigureMinterController) Route() string {
	return RouterKey
}

func (msg *MsgConfigureMinterController) Type() string {
	return TypeMsgConfigureMinterController
}

func (msg *MsgConfigureMinterController) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

func (msg *MsgConfigureMinterController) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgConfigureMinterController) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return errors.Wrapf(errorsTypes.ErrInvalidAddress, "invalid from address (%s)", err)
	}
	_, err = sdk.AccAddressFromBech32(msg.Controller)
	if err != nil {
		return errors.Wrapf(errorsTypes.ErrInvalidAddress, "invalid controller address (%s)", err)
	}
	_, err = sdk.AccAddressFromBech32(msg.Minter)
	if err != nil {
		return errors.Wrapf(errorsTypes.ErrInvalidAddress, "invalid minter address (%s)", err)
	}
	return nil
}
