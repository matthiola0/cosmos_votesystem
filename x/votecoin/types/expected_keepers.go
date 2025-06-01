package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types" // Renamed import to avoid conflict if types from auth is directly used. Usually it's just 'types'
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"   // 新增 banktypes 導入
)

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) authtypes.AccountI // Changed 'types.AccountI' to 'authtypes.AccountI' for clarity
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	// --- 原有的方法 ---
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins

	// --- 為 genesis.go 添加的方法 ---
	GetDenomMetaData(ctx sdk.Context, denom string) (banktypes.Metadata, bool)
	SetDenomMetaData(ctx sdk.Context, denomMetaData banktypes.Metadata)

	// --- 為 msg_server_issue_vote_coin.go (第 6.2.B 步) 添加的方法 ---
	MintCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddress sdk.AccAddress, amt sdk.Coins) error
	BurnCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error // 用於 IssueVoteCoin 失敗時的回滾

	// --- 其他可能需要的方法 (根據實際 bank keeper 的使用情況添加) ---
	// GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin
	// SendCoins(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error
	// etc.
}