package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	ibctransferkeeper "github.com/cosmos/ibc-go/v3/modules/apps/transfer/keeper"
	transfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	porttypes "github.com/cosmos/ibc-go/v3/modules/core/05-port/types"
	host "github.com/cosmos/ibc-go/v3/modules/core/24-host"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/pundix/pundix/x/ibc/applications/transfer/types"
)

// Keeper defines the IBC fungible transfer keeper
type Keeper struct {
	ibctransferkeeper.Keeper
	storeKey   sdk.StoreKey
	cdc        codec.BinaryCodec
	paramSpace paramtypes.Subspace

	ics4Wrapper   porttypes.ICS4Wrapper
	channelKeeper transfertypes.ChannelKeeper
	portKeeper    transfertypes.PortKeeper
	authKeeper    transfertypes.AccountKeeper
	bankKeeper    transfertypes.BankKeeper
	scopedKeeper  capabilitykeeper.ScopedKeeper
}

// NewKeeper creates a new IBC transfer Keeper instance
func NewKeeper(keeper ibctransferkeeper.Keeper,
	cdc codec.BinaryCodec, key sdk.StoreKey, paramSpace paramtypes.Subspace,
	ics4Wrapper porttypes.ICS4Wrapper, channelKeeper transfertypes.ChannelKeeper, portKeeper transfertypes.PortKeeper,
	authKeeper transfertypes.AccountKeeper, bankKeeper transfertypes.BankKeeper, scopedKeeper capabilitykeeper.ScopedKeeper,
) Keeper {
	// ensure ibc transfer module account is set
	if addr := authKeeper.GetModuleAddress(types.ModuleName); addr == nil {
		panic("the FX IBC transfer module account has not been set")
	}

	return Keeper{
		Keeper:        keeper,
		cdc:           cdc,
		storeKey:      key,
		paramSpace:    paramSpace,
		ics4Wrapper:   ics4Wrapper,
		channelKeeper: channelKeeper,
		portKeeper:    portKeeper,
		authKeeper:    authKeeper,
		bankKeeper:    bankKeeper,
		scopedKeeper:  scopedKeeper,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+host.ModuleName+"-"+types.CompatibleModuleName)
}
