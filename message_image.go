package feishu

// ImageMessage 图片
type ImageMessage struct {
	BaseMessage
	// ImageKey  图片资源
	ImageKey string
}

func (message *ImageMessage) ToMessageMap() map[string]interface{} {
	content := map[string]string{}
	content["image_key"] = message.ImageKey
	imageMessage := map[string]interface{}{}
	imageMessage["msg_type"] = message.GetMsgType()
	imageMessage["content"] = content
	return imageMessage
}

func (message *ImageMessage) GetMsgType() MsgType {
	return MsgTypeImage
}

// NewImageMessage create ImageMessage
func NewImageMessage() *ImageMessage {
	return new(ImageMessage)
}

// SetImageKey set ImageMessage.ImageKey
func (message *ImageMessage) SetImageKey(key string) *ImageMessage {
	message.ImageKey = key
	return message
}
