package keeper

import (
	"fmt"

	"cosmossdk.io/errors"
	"github.com/circlefin/noble-fiattokenfactory/x/fiattokenfactory/types"
	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   storetypes.StoreKey
		paramstore paramtypes.Subspace

		bankKeeper types.BankKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey storetypes.StoreKey,
	bankKeeper types.BankKeeper,
) *Keeper {
	return &Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		bankKeeper: bankKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// ValidatePrivileges checks if a specified address has already been assigned to a privileged role.
func (k Keeper) ValidatePrivileges(ctx sdk.Context, address string) error {
	acc, err := sdk.AccAddressFromBech32(address)
	if err != nil {
		return err
	}

	owner, found := k.GetOwner(ctx)
	if found && owner.Address == acc.String() {
		return errors.Wrapf(types.ErrAlreadyPrivileged, "cannot assign (%s) to owner role", acc.String())
	}

	blacklister, found := k.GetBlacklister(ctx)
	if found && blacklister.Address == acc.String() {
		return errors.Wrapf(types.ErrAlreadyPrivileged, "cannot assign (%s) to black lister role", acc.String())
	}

	masterminter, found := k.GetMasterMinter(ctx)
	if found && masterminter.Address == acc.String() {
		return errors.Wrapf(types.ErrAlreadyPrivileged, "cannot assign (%s) to master minter role", acc.String())
	}

	pauser, found := k.GetPauser(ctx)
	if found && pauser.Address == acc.String() {
		return errors.Wrapf(types.ErrAlreadyPrivileged, "cannot assign (%s) to pauser role", acc.String())
	}

	return nil
}
