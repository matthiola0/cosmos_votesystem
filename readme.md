# Cosmos VoteSystem

Cosmos VoteSystem 是一個基於 Cosmos SDK 和 Ignite CLI 開發的區塊鏈投票系統概念驗證 (PoC) 專案。它旨在模擬一個去中心化、安全且可驗證的電子投票流程，涵蓋了從選民身份註冊、投票憑證發放到最終計票的完整生命週期。

此專案是為區塊鏈課程期末專題而設計，完整記錄了從環境建置、模組開發、多節點網路設定到效能測試的整個過程。

## 核心功能

### 身份驗證模組 (identity)

* **選民註冊**：將使用者地址與身份證明雜湊綁定。
* **狀態查詢**：查詢特定地址是否為已註冊的選民。

### 投票代幣模組 (votecoin)

* **代幣發行**：由管理員向已驗證的選民發放 `votecredit`。
* **餘額管理**：整合 Cosmos SDK 的 bank 模組進行代幣管理。

### 投票核心模組 (voting)

* **投票執行**：選民消耗 `votecredit` 來投票給特定候選人。
* **防止重複投票**：確保每位選民在單一選舉中只能投票一次。
* **結果查詢**：提供即時、透明的票數統計與查詢功能。

### 權威證明 (Proof-of-Authority, POA)

可配置為 POA 共識機制，適用於需要許可控制的聯盟鏈或私有鏈場景，限制只有授權節點才能生成區塊。

### 多節點與分區部署

支援標準的多節點網路設定，並可部署為獨立運作的邏輯分區 (Zone)。

## 技術棧

* 區塊鏈框架：Cosmos SDK
* 開發工具：Ignite CLI
* 共識引擎：CometBFT（前身為 Tendermint Core）
* 程式語言：Go

## 快速開始

### 1. 環境準備

請確保您已安裝以下工具：

* Go 1.22 或更高版本
* Git
* Ignite CLI v0.27.2

### 2. 安裝與啟動

```bash
# 1. 複製專案原始碼
git clone https://github.com/matthiola0/cosmos_votesystem.git
cd cosmos_votesystem

# 2. 安裝 Go 依賴項
go mod tidy

# 3. 編譯並安裝鏈的二進制檔案 (votesystemd)
ignite chain build

# 4. (可選) 清理舊的測試數據
rm -rf ~/.votesystem

# 5. 初始化單節點測試網路
ignite chain serve -c config.yml --reset-once
```

上述 `ignite chain serve` 命令將會自動初始化鏈、建立測試帳戶（alice, bob）並啟動一個單節點網路。

## CLI 使用範例

以下指令建議在另一個終端機視窗中執行。

### 1. 準備工作 - 取得帳戶地址

```bash
# 取得 alice 和 bob 的地址
ALICE_ADDR=$(votesystemd keys show alice -a --keyring-backend test)
BOB_ADDR=$(votesystemd keys show bob -a --keyring-backend test)
```

### 2. 身份註冊 (Identity Module)

alice 作為管理員，將 bob 註冊為合格選民。

```bash
votesystemd tx identity register-voter "bob_identity_hash_123" \
  --from $BOB_ADDR \
  --chain-id votesystem \
  --keyring-backend test \
  --fees 5stake -y
```

### 3. 發放投票代幣 (Votecoin Module)

alice（作為管理員）向 bob 發放 100 `votecredit`。

```bash
votesystemd tx votecoin issue-vote-coin $BOB_ADDR 100 \
  --from $ALICE_ADDR \
  --chain-id votesystem \
  --keyring-backend test \
  --fees 5stake -y
```

### 4. 進行投票 (Voting Module)

bob 參與 `election_2024` 選舉，投票給 `candidate_A`，此操作將消耗 1 `votecredit`。

```bash
votesystemd tx voting cast-vote "candidate_A" "election_2024" \
  --from $BOB_ADDR \
  --chain-id votesystem \
  --keyring-backend test \
  --fees 5stake -y
```

### 5. 查詢投票結果

任何人都可以查詢 `election_2024` 選舉的當前結果：

```bash
votesystemd query voting get-election-result "election_2024"
```

預期輸出：

```json
{
  "election_id": "election_2024",
  "results": {
    "candidate_A": "1"
  }
}
```

## 授權

本專案採用 Apache 2.0 License。
