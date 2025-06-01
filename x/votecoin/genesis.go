package votecoin

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types" // 新增 banktypes 導入
	"votesystem/x/votecoin/keeper"
	"votesystem/x/votecoin/types"
)

// 這些常量可以放在這裡，或者更規範地放在 x/votecoin/types/ 包下的某個文件中 (例如 keys.go 或 coin.go)
const (
	VoteCoinDenom    = "votecredit" // 與您在第 5 步 add-genesis-account 中使用的面額一致
	VoteCoinDisplay  = "VoteCredit"
	VoteCoinExponent = 0 // 如果 votecredit 是最小單位，則為 0
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// 保留 SetParams (如果您的模組有參數的話)
	// this line is used by starport scaffolding # genesis/module/init
    // 檢查您的 types.GenesisState 是否真的有 Params 欄位，以及您的 keeper 是否有 SetParams 方法。
    // 如果 votecoin 模組沒有自定義參數，可以移除或註釋掉下面這行。
	k.SetParams(ctx, genState.Params)


	// 設置 votecredit 代幣的元數據
	// 確保您的 keeper.Keeper 結構體中注入了 BankKeeper，並且該 BankKeeper 具有 GetDenomMetaData 和 SetDenomMetaData 方法。
	// 通常 BankKeeper 欄位名為 BankKeeper (首字母大寫以便導出)。
	// k.BankKeeper 應該是您在 x/votecoin/keeper/keeper.go 中定義的 Keeper 結構體的一個欄位。
	metadata, found := k.BankKeeper.GetDenomMetaData(ctx, VoteCoinDenom)
	if !found {
		k.BankKeeper.SetDenomMetaData(ctx, banktypes.Metadata{
			Description: "A non-transferable token for voting purposes.",
			DenomUnits: []*banktypes.DenomUnit{
				{Denom: VoteCoinDenom, Exponent: uint32(VoteCoinExponent), Aliases: []string{"voting credit"}},
			},
			Base:    VoteCoinDenom,
			Display: VoteCoinDisplay,
			Name:    VoteCoinDisplay,
			Symbol:  "VCRD", // 您可以自訂一個代幣符號
		})
	} else {
		ctx.Logger().Info("VoteCoin metadata already exists", "denom", VoteCoinDenom, "description", metadata.Description)
	}
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
    // 檢查您的 keeper 是否有 GetParams 方法，以及 types.DefaultGenesis() 返回的 GenesisState 是否有 Params 欄位。
    // 如果 votecoin 模組沒有自定義參數，可以移除或註釋掉下面這行。
	genesis.Params = k.GetParams(ctx)

	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}