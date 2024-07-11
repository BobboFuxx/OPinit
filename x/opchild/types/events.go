package types

// Rollup module event types
const (
	EventTypeAddValidator            = "add_validator"
	EventTypeRemoveValidator         = "remove_validator"
	EventTypeFinalizeTokenDeposit    = "finalize_token_deposit"
	EventTypeInitiateTokenWithdrawal = "initiate_token_withdrawal"
	EventTypeExecuteMessages         = "execute_messages"
	EventTypeWhitelist               = "whitelist"
	EventTypeParams                  = "params"
	EventTypeUpdateOracle            = "update_oracle"
	EventTypeSetBridgeInfo           = "set_bridge_info"

	AttributeKeySender         = "sender"
	AttributeKeyBridgeId       = "bridge_id"
	AttributeKeyBridgeAddr     = "bridge_addr"
	AttributeKeyL1ChainId      = "l1_chain_id"
	AttributeKeyL1ClientId     = "l1_client_id"
	AttributeKeyRecipient      = "recipient"
	AttributeKeyAmount         = "amount"
	AttributeKeyDenom          = "denom"
	AttributeKeyBaseDenom      = "base_denom"
	AttributeKeyStructTag      = "struct_tag"
	AttributeKeyValidator      = "validator"
	AttributeKeyMetadata       = "metadata"
	AttributeKeyL1Sequence     = "l1_sequence"
	AttributeKeyFinalizeHeight = "finalize_height"
	AttributeKeyHookSuccess    = "hook_success"
	AttributeKeyFrom           = "from"
	AttributeKeyTo             = "to"
	AttributeKeyL2Sequence     = "l2_sequence"
	AttributeKeyHeight         = "height"
)
