package keeper

import (
	"context"
	"fmt"

	"votesystem/x/identity/types" 
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) RegisterVoter(goCtx context.Context, msg *types.MsgRegisterVoter) (*types.MsgRegisterVoterResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// 1. 驗證 identityProofHash 是否為空
	if msg.IdentityProofHash == "" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "identity proof hash cannot be empty")
	}

	// 2. 獲取簽名者地址 (即選民地址)
	voterAddress, err := sdk.AccAddressFromBech32(msg.Voter)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid voter address")
	}
	// 添加日誌：記錄嘗試註冊的選民地址和雜湊
	k.Logger(ctx).Info("Attempting to register voter", "address", voterAddress.String(), "hash", msg.IdentityProofHash)

	// 3. 檢查選民是否已註冊
	allRegisteredVoters := k.GetAllRegisteredVoter(ctx)
	// 添加日誌：記錄 GetAllRegisteredVoter 返回的結果
	k.Logger(ctx).Info("Retrieved voters for check", "count", len(allRegisteredVoters))
	if len(allRegisteredVoters) > 0 {
		// 如果列表不為空，記錄列表中第一個選民的詳細資訊以供檢查
		k.Logger(ctx).Info("First voter in list", "address", allRegisteredVoters[0].VoterAddress, "hash", allRegisteredVoters[0].IdentityProofHash)
	}


	for i, registeredVoter := range allRegisteredVoters {
		// 添加日誌：記錄正在檢查的每個選民以及比較結果
		k.Logger(ctx).Info("Checking voter from store",
			"index", i,
			"storedVoterAddress", registeredVoter.VoterAddress,
			"msgVoterAddress", voterAddress.String(),
			"isMatch", registeredVoter.VoterAddress == voterAddress.String())

		if registeredVoter.VoterAddress == voterAddress.String() {
			// 添加日誌：記錄發現選民已註冊
			k.Logger(ctx).Info("Voter already registered", "address", voterAddress.String())
			return nil, sdkerrors.Wrap(types.ErrVoterAlreadyRegistered, fmt.Sprintf("voter %s already registered", voterAddress.String()))
		}
	}

	// 4. 創建新的 RegisteredVoter 對象
	newVoter := types.RegisteredVoter{
		VoterAddress:      voterAddress.String(),
		IdentityProofHash: msg.IdentityProofHash,
	}

	// 5. 將新選民添加到儲存中
	k.AppendRegisteredVoter(ctx, newVoter)
	// 添加日誌：記錄已添加的新選民
	k.Logger(ctx).Info("Appended new voter", "address", newVoter.VoterAddress, "hash", newVoter.IdentityProofHash)

	// ... 函數的其餘部分
	err = ctx.EventManager().EmitTypedEvent(&types.EventVoterRegistered{
		VoterAddress:      voterAddress.String(),
		IdentityProofHash: msg.IdentityProofHash,
	})
	if err != nil {
		k.Logger(ctx).Error("failed to emit EventVoterRegistered", "error", err)
	}

	return &types.MsgRegisterVoterResponse{}, nil
}