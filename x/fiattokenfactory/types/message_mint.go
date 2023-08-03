package types

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	errorsTypes "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgMint = "mint"

var _ sdk.Msg = &MsgMint{}

func NewMsgMint(from string, address string, amount sdk.Coin) *MsgMint {
	return &MsgMint{
		From:    from,
		Address: address,
		Amount:  amount,
	}
}

func (msg *MsgMint) Route() string {
	return RouterKey
}

func (msg *MsgMint) Type() string {
	return TypeMsgMint
}

func (msg *MsgMint) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

func (msg *MsgMint) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgMint) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return errors.Wrapf(errorsTypes.ErrInvalidAddress, "invalid from address (%s)", err)
	}

	_, err = sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return errors.Wrapf(errorsTypes.ErrInvalidAddress, "invalid address (%s)", err)
	}

	if msg.Amount.IsNil() {
		return errors.Wrap(errorsTypes.ErrInvalidCoins, "mint amount cannot be nil")
	}

	if msg.Amount.IsNegative() {
		return errors.Wrap(errorsTypes.ErrInvalidCoins, "mint amount cannot be negative")
	}

	if msg.Amount.IsZero() {
		return errors.Wrap(errorsTypes.ErrInvalidCoins, "mint amount cannot be zero")
	}

	return nil
}
