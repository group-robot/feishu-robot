package feishu

type MsgType string

// MsgType
const (
	MsgtypeText        MsgType = "text"
	MsgTypePost        MsgType = "post"
	MsgTypeImage       MsgType = "image"
	MsgTypeShareChat   MsgType = "share_chat"
	MsgTypeInteractive MsgType = "interactive"
)

// Message interface
type Message interface {
	ToMessageMap() map[string]interface{}
	GetMsgType() MsgType
}

type BaseMessage struct {
	MsgType MsgType
}
