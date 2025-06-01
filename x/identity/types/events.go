package types

// Event types for the identity module
const (
	EventTypeVoterRegistered = "voter_registered"
	// Attribute keys
	AttributeKeyVoterAddress      = "voter_address"
	AttributeKeyIdentityProofHash = "identity_proof_hash"
)

// EventVoterRegistered is an event triggered when a new voter is registered
// 由於 .proto 中沒有直接生成 Event 結構體，我們手動定義一個類似的結構體用於 EmitTypedEvent
// 或者您可以使用 sdk.NewEvent 和 sdk.NewAttribute 來構建事件
// type EventVoterRegistered struct {
// 	VoterAddress      string
// 	IdentityProofHash string
// }