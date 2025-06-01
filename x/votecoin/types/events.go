package types

// Event types for the votecoin module
const (
	EventTypeVoteCoinIssued = "votecoin_issued" // 事件類型字符串

	// Attribute keys
	AttributeKeyAdmin     = "admin"
	AttributeKeyRecipient = "recipient"
	AttributeKeyAmount    = "amount"
)

// EventVoteCoinIssued is an event triggered when new vote coins are issued.
// type EventVoteCoinIssued struct {
// 	Admin     string
// 	Recipient string
// 	Amount    string // 通常事件屬性值為字符串
// }