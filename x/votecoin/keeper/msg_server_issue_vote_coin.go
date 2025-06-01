package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"votesystem/x/votecoin/types" // 確保這裡的路徑是正確的
)

// IssueVoteCoin 處理管理員向特定接收者發放 VoteCoin 的邏輯
func (k msgServer) IssueVoteCoin(goCtx context.Context, msg *types.MsgIssueVoteCoin) (*types.MsgIssueVoteCoinResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	k.Logger(ctx).Error("!!!!!!!!!! IssueVoteCoin function was called !!!!!!!!!!", "admin", msg.Admin, "recipient", msg.Recipient, "amount", msg.Amount) // 使用 Error 級別使其更顯眼
	
	// 1. 獲取簽名者 (管理員) 地址
	adminAddress, err := sdk.AccAddressFromBech32(msg.Admin)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid admin address")
	}

	// 2. 驗證簽名者是否是管理員 (簡化版，實際項目中需加強)
	// 在一個真實的應用中，您需要有一個明確的管理員列表或基於角色的訪問控制。
	// 這裡我們僅作日誌記錄，並假設交易發起者有權限。
	// 您可以從模組參數或一個固定的創世地址來驗證管理員。
	k.Logger(ctx).Info(fmt.Sprintf("IssueVoteCoin called by: %s. In a real app, verify this is an admin!", msg.Admin))
	// 例如:
	// expectedAdmin := sdk.MustAccAddressFromBech32("cosmos1...") // 從配置或參數中獲取
	// if !adminAddress.Equals(expectedAdmin) {
	// 	return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "signer is not the admin")
	// }

	// 3. 獲取接收者地址
	recipientAddress, err := sdk.AccAddressFromBech32(msg.Recipient)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid recipient address")
	}

	// 4. 驗證數量
	if msg.Amount <= 0 { // MsgIssueVoteCoin 中的 Amount 通常是 int64 (由 scaffold 生成)
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "amount must be positive")
	}
	amount := sdk.NewInt(int64(msg.Amount))
	// 使用在 types 包中定義的 VoteCoinDenom
	coinsToIssue := sdk.NewCoins(sdk.NewCoin(types.VoteCoinDenom, amount))

	// 5. 鑄造代幣到 votecoin 模組帳戶
	// types.ModuleName 應該在 x/votecoin/types/keys.go 中定義為 const ModuleName = "votecoin"
	err = k.BankKeeper.MintCoins(ctx, types.ModuleName, coinsToIssue)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "failed to mint coins")
	}

	// 6. 從模組帳戶發送代幣給接收者
	err = k.BankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, recipientAddress, coinsToIssue)
	if err != nil {
		// 如果發送失敗，嘗試銷毀之前鑄造的代幣以維持總量不變（關鍵操作）
		burnErr := k.BankKeeper.BurnCoins(ctx, types.ModuleName, coinsToIssue)
		if burnErr != nil {
			k.Logger(ctx).Error(fmt.Sprintf("CRITICAL: failed to burn coins after send failed: %v. Module account may have unallocated tokens.", burnErr))
			// 根據您的策略，這裡可能需要 panic 或返回一個更嚴重的錯誤
		}
		return nil, sdkerrors.Wrap(err, "failed to send coins to recipient")
	}

	// 7. (可選) 發出事件
	// 確保 EventVoteCoinIssued 在 types 包中定義
	err = ctx.EventManager().EmitTypedEvent(&types.EventVoteCoinIssued{
		Admin:     adminAddress.String(),
		Recipient: recipientAddress.String(),
		Amount:    amount.String(), // amount.String() 會得到數字的字符串形式
	})
	if err != nil {
		k.Logger(ctx).Error("failed to emit EventVoteCoinIssued", "error", err)
		// 即使事件發送失敗，核心邏輯也已完成，通常不應因此返回錯誤，但記錄是必要的。
	}

	return &types.MsgIssueVoteCoinResponse{}, nil
}