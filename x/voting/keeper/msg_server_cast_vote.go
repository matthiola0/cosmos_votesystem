package keeper

import (
	"context"
	"fmt"

	identitytypes "votesystem/x/identity/types"   // 導入 identity types 包
	votecointypes "votesystem/x/votecoin/types" // 為了 ModuleName 和 VoteCoinDenom
	"votesystem/x/voting/types" 
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const VoteCost = 1 // 每票消耗 1 votecredit

// Helper function to create the composite index
func createCompositeHasVotedIndex(electionID string, voterAddress string) string {
	return electionID + "/" + voterAddress // 您可以選擇不同的分隔符
}

func (k msgServer) CastVote(goCtx context.Context, msg *types.MsgCastVote) (*types.MsgCastVoteResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	voterAddress, err := sdk.AccAddressFromBech32(msg.Voter)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid voter address")
	}

	if msg.ElectionID == "" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "election ID cannot be empty")
	}
	if msg.CandidateID == "" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "candidate ID cannot be empty")
	}

	// 3. 使用 identityKeeper 檢查選民是否已註冊
	// k.identityKeeper 應在 x/voting/keeper/keeper.go 中的 Keeper 結構體中定義，並在 app.go 中注入
	if !k.identityKeeper.IsVoterRegistered(ctx, voterAddress) { 
		return nil, sdkerrors.Wrap(identitytypes.ErrVoterNotRegistered, fmt.Sprintf("voter %s is not registered", voterAddress.String()))
	}

	// 4. 使用 bankKeeper 檢查選民是否有足夠的 VoteCoin
	voteCoinDenom := votecointypes.VoteCoinDenom // "votecredit" from x/votecoin/genesis.go
	balance := k.bankKeeper.GetBalance(ctx, voterAddress, voteCoinDenom)
	cost := sdk.NewInt(VoteCost)
	if balance.Amount.LT(cost) {
		return nil, sdkerrors.Wrap(types.ErrInsufficientVoteCoin, fmt.Sprintf("insufficient %s to vote. Required: %s, Have: %s", voteCoinDenom, cost.String(), balance.Amount.String()))
	}

	// 5. 構造複合索引並檢查選民是否已在該選舉中投過票
	compositeIndex := createCompositeHasVotedIndex(msg.ElectionID, voterAddress.String())
	hasVotedEntry, found := k.GetHasVotedVoter(ctx, compositeIndex) // 使用單個 compositeIndex 調用
	if found && hasVotedEntry.VotedStatus { // 假設 hasVotedEntry 包含 VotedStatus 字段
		return nil, sdkerrors.Wrap(types.ErrAlreadyVoted, fmt.Sprintf("voter %s has already voted in election %s", voterAddress.String(), msg.ElectionID))
	}

	// 6. 銷毀或鎖定已使用的 VoteCoin
	coinsToBurn := sdk.NewCoins(sdk.NewCoin(voteCoinDenom, cost))
	// 從選民帳戶發送到 voting 模組帳戶
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, voterAddress, types.ModuleName, coinsToBurn) // types.ModuleName is "voting"
	if err != nil {
		return nil, sdkerrors.Wrap(err, "failed to send coins from voter to module for burning")
	}
	// 從 voting 模組帳戶銷毀
	err = k.bankKeeper.BurnCoins(ctx, types.ModuleName, coinsToBurn)
	if err != nil {
		k.Logger(ctx).Error(fmt.Sprintf("critical: failed to burn coins from module account: %v", err))
		// 根據您的錯誤處理策略，決定是否返回錯誤或繼續
		// return nil, sdkerrors.Wrap(err, "failed to burn coins from module") 
	}

	// 7. 記錄選票到 voteRecord 列表
	voteRecord := types.VoteRecord{
		// Index/ID 由 AppendVoteRecord 自動處理 (如果適用)
		VoterAddress: voterAddress.String(),
		CandidateID:  msg.CandidateID,
		ElectionID:   msg.ElectionID,
	}
	k.AppendVoteRecord(ctx, voteRecord) // 確保此函數存在且按預期工作

	// 8. 更新 HasVotedVoter map
	newHasVotedEntry := types.HasVotedVoter{
		Index:        compositeIndex,          // 將複合鍵存入 Index 字段
		ElectionID:   msg.ElectionID,          // 同時存儲原始 ElectionID
		VoterAddress: voterAddress.String(), // 同時存儲原始 VoterAddress
		VotedStatus:  true,
	}
	k.SetHasVotedVoter(ctx, newHasVotedEntry)

	// 9. (可選) Emit event
	err = ctx.EventManager().EmitTypedEvent(&types.EventVoteCast{ // 確保 EventVoteCast 已定義
		VoterAddress: voterAddress.String(),
		ElectionId:   msg.ElectionID,
		CandidateId:  msg.CandidateID,
	})
	if err != nil {
		k.Logger(ctx).Error("failed to emit EventVoteCast", "error", err)
	}

	return &types.MsgCastVoteResponse{}, nil
}
