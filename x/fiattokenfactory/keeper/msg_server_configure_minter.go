package keeper

import (
	"context"

	"cosmossdk.io/errors"
	"github.com/circlefin/noble-fiattokenfactory/x/fiattokenfactory/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) ConfigureMinter(goCtx context.Context, msg *types.MsgConfigureMinter) (*types.MsgConfigureMinterResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	mintingDenom := k.GetMintingDenom(ctx)

	if msg.Allowance.Denom != mintingDenom.Denom {
		return nil, errors.Wrapf(types.ErrMint, "minting denom is incorrect")
	}

	minterController, found := k.GetMinterController(ctx, msg.From)
	if !found {
		return nil, errors.Wrapf(types.ErrUnauthorized, "minter controller not found")
	}

	if msg.From != minterController.Controller {
		return nil, errors.Wrapf(types.ErrUnauthorized, "you are not a controller of this minter")
	}

	if msg.Address != minterController.Minter {
		return nil, errors.Wrapf(
			types.ErrUnauthorized,
			"minter address ≠ minter controller's minter address, (%s≠%s)",
			msg.Address, minterController.Minter,
		)
	}

	k.SetMinters(ctx, types.Minters{
		Address:   msg.Address,
		Allowance: msg.Allowance,
	})

	err := ctx.EventManager().EmitTypedEvent(msg)

	return &types.MsgConfigureMinterResponse{}, err
}
