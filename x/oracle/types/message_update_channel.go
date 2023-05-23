package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUpdateChannel = "update_channel"

var _ sdk.Msg = &MsgUpdateChannel{}

func NewMsgUpdateChannel(creator string, channel string) *MsgUpdateChannel {
	return &MsgUpdateChannel{
		Creator: creator,
		Channel: channel,
	}
}

func (msg *MsgUpdateChannel) Route() string {
	return RouterKey
}

func (msg *MsgUpdateChannel) Type() string {
	return TypeMsgUpdateChannel
}

func (msg *MsgUpdateChannel) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateChannel) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateChannel) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
