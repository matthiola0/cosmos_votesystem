package types

import (
	// banktypes "github.com/cosmos/cosmos-sdk/x/bank/types" // 新增：為 BankKeeper 的方法簽名導入 banktypes
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types" 
)

// IdentityKeeper defines the expected interface for the identity keeper.
type IdentityKeeper interface {
	// IsVoterRegistered checks if a voter is registered.
	IsVoterRegistered(ctx sdk.Context, voterAddr sdk.AccAddress) bool
	// 如果 voting 模組需要 identity 模組的其他功能，在此處添加
}

// VotecoinKeeper defines the expected interface for the votecoin keeper.
type VotecoinKeeper interface {
	// 目前 CastVote 邏輯直接使用 votecointypes.VoteCoinDenom。
	// 如果將來需要從 votecoin keeper 獲取 denom 或其他信息，可以在此添加方法。
	// 例如： GetVoteCoinDenom(ctx sdk.Context) string
}

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) authtypes.AccountI // 使用重命名的 authtypes
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface needed to retrieve account balances and send coins.
type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	
	GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddress sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	BurnCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddress sdk.AccAddress, amt sdk.Coins) error
	MintCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
}