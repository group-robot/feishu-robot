package feishu

// MsgType 消息类型
type MsgType string

// MsgType
const (
	// MsgTypeText 文本
	MsgTypeText MsgType = "text"
	// MsgTypePost 富文本
	MsgTypePost MsgType = "post"
	// MsgTypeImage 图片
	MsgTypeImage     MsgType = "image"
	MsgTypeShareChat MsgType = "share_chat"
	// MsgTypeInteractive 消息卡片
	MsgTypeInteractive MsgType = "interactive"
)

// Message interface
type Message interface {
	// ToMessageMap to map
	ToMessageMap() map[string]interface{}
}

// ContentMessage  消息内容
type ContentMessage interface {
	Message
	// GetMsgType 消息类型
	GetMsgType() MsgType
}

// BaseMessage 消息基础
type BaseMessage struct {
	// MsgType 消息类型
	MsgType MsgType
}
