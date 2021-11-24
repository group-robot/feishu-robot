package feishu

type TextMessage struct {
	BaseMessage
	Content string
	isAtAll bool
}

func (message *TextMessage) GetMsgType() MsgType {
	return MsgtypeText
}
func (message TextMessage) ToMessageMap() map[string]interface{} {
	contentMap := map[string]string{}
	content := ""
	if message.isAtAll {
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

func NewTextMessage() *TextMessage {
	return new(TextMessage)
}

func (message *TextMessage) SetContent(content string) *TextMessage {
	message.Content = content
	return message
}
func (message *TextMessage) SetAtAll(atAll bool) *TextMessage {
	message.isAtAll = atAll
	return message
}
func (message *TextMessage) IsAtAll() *TextMessage {
	message.isAtAll = true
	return message
}
