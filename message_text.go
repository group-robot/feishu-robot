package feishu

// TextMessage 文本消息
type TextMessage struct {
	BaseMessage
	// Content 文本内容
	Content string
	// AtAll 是否at全部
	AtAll bool
}

func (message *TextMessage) GetMsgType() MsgType {
	return MsgTypeText
}
func (message TextMessage) ToMessageMap() map[string]interface{} {
	contentMap := map[string]string{}
	content := ""
	if message.AtAll {
		content = "<at user_id=\"all\">所有</at> " + message.Content
	} else {
		content = message.Content
	}
	contentMap["text"] = content
	//text
	textMap := map[string]interface{}{}
	textMap["msg_type"] = message.GetMsgType()
	textMap["content"] = contentMap
	return textMap
}

// NewTextMessage create TextMessage
func NewTextMessage() *TextMessage {
	return new(TextMessage)
}

// SetContent set TextMessage.Content
func (message *TextMessage) SetContent(content string) *TextMessage {
	message.Content = content
	return message
}

// SetAtAll set TextMessage.AtAll
func (message *TextMessage) SetAtAll(atAll bool) *TextMessage {
	message.AtAll = atAll
	return message
}

// IsAtAll set TextMessage.AtAll is true
func (message *TextMessage) IsAtAll() *TextMessage {
	message.AtAll = true
	return message
}
