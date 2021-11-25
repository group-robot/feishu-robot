package feishu

// InteractiveMessage 消息卡片
type InteractiveMessage struct {
	BaseMessage
	// CardJsonStr json格式内容，如果不为空，则其他元素失效
	CardJsonStr string
	// Link 整个卡片的链接
	Link *CardLink
	// Header 用于配置卡片标题内容。
	Header *CardHeader
	// Config 用于描述卡片的功能属性。
	Config *CardConfig
	// Elements 用于定义卡片正文内容,和i18n_elements至少必填其中1个
	Elements *ElementsContent
	// I18nElement 为卡片的正文部分定义多语言内容，和elements至少必填其中1个
	I18nElement []*I18nElementsContent
}

// NewInteractiveMessage create InteractiveMessage
func NewInteractiveMessage() *InteractiveMessage {
	return &InteractiveMessage{
		I18nElement: []*I18nElementsContent{},
	}
}
func (m *InteractiveMessage) GetMsgType() MsgType {
	return MsgTypeInteractive
}

// SetHeader set InteractiveMessage.Header
func (m *InteractiveMessage) SetHeader(header *CardHeader) *InteractiveMessage {
	m.Header = header
	return m
}

// SetConfig set InteractiveMessage.Config
func (m *InteractiveMessage) SetConfig(config *CardConfig) *InteractiveMessage {
	m.Config = config
	return m
}

// SetElements set InteractiveMessage.Elements
func (m *InteractiveMessage) SetElements(elements *ElementsContent) *InteractiveMessage {
	m.Elements = elements
	return m
}

// SetI18nElement set InteractiveMessage.I18nElement
func (m *InteractiveMessage) SetI18nElement(i18nElements []*I18nElementsContent) *InteractiveMessage {
	m.I18nElement = i18nElements
	return m
}

// SetCardJsonStr set InteractiveMessage.CardJsonStr
func (m *InteractiveMessage) SetCardJsonStr(jsonStr string) *InteractiveMessage {
	m.CardJsonStr = jsonStr
	return m
}

// AddI18nElement add InteractiveMessage.I18nElement
func (m *InteractiveMessage) AddI18nElement(i18nElement *I18nElementsContent) *InteractiveMessage {
	m.I18nElement = append(m.I18nElement, i18nElement)
	return m
}
func (m *InteractiveMessage) ToMessageMap() map[string]interface{} {
	message := map[string]interface{}{}
	if m.Config != nil {
		message["config"] = m.Config.ToMessageMap()
	}
	if m.Header != nil {
		message["header"] = m.Header.ToMessageMap()
	}
	if m.Elements != nil {
		message["elements"] = m.Elements.ToMessageMap()
	}
	if len(m.I18nElement) > 0 {
		var i18nElementsMessage map[string]interface{}
		for _, i18nElement := range m.I18nElement {
			for k, v := range i18nElement.ToMessageMap() {
				i18nElementsMessage[k] = v
			}
		}
		message["i18n_elements"] = i18nElementsMessage
	}
	if m.Link != nil {
		linkMessage := m.Link.ToMessageMap()
		message["card_link"] = linkMessage
	}
	cardMessage := map[string]interface{}{}
	cardMessage["msg_type"] = m.GetMsgType()
	if m.CardJsonStr == "" {
		cardMessage["card"] = message
	} else {
		cardMessage["card"] = m.CardJsonStr
	}
	return cardMessage
}

// CardConfig 配置卡片属性
type CardConfig struct {
	// EnableForward 是否允许卡片被转发。默认 true
	EnableForward bool
	// WideScreenMode 是否根据屏幕宽度动态调整消息卡片宽度，默认值为true（已废弃）
	WideScreenMode bool
}

// NewCardConfig create CardConfig
func NewCardConfig() *CardConfig {
	return &CardConfig{
		EnableForward:  true,
		WideScreenMode: true,
	}
}

// SetEnableForward set CardConfig.SetEnableForward
func (c *CardConfig) SetEnableForward(enableForward bool) *CardConfig {
	c.EnableForward = enableForward
	return c
}

// SetWideScreenMode set CardConfig.WideScreenMode
func (c *CardConfig) SetWideScreenMode(wideScreenMode bool) *CardConfig {
	c.WideScreenMode = wideScreenMode
	return c
}

func (c *CardConfig) ToMessageMap() map[string]interface{} {
	message := map[string]interface{}{}
	message["enable_forward"] = c.EnableForward
	message["wide_screen_mode"] = c.WideScreenMode
	return message
}

// CardHeader 卡片标题
type CardHeader struct {
	// Title 配置卡片标题内容
	Title *CardTitle
	// Template	控制标题背景颜色
	Template string
}

// NewCardHeader create CardHeader
func NewCardHeader() *CardHeader {
	return new(CardHeader)
}

// SetTitle set CardHeader.Title
func (h *CardHeader) SetTitle(title *CardTitle) *CardHeader {
	h.Title = title
	return h
}

// SetTemplate  set CardHeader.Template
func (h *CardHeader) SetTemplate(template string) *CardHeader {
	h.Template = template
	return h
}
func (h *CardHeader) ToMessageMap() map[string]interface{} {
	message := map[string]interface{}{}
	message["title"] = h.Title.ToMessageMap()
	if h.Template != "" {
		message["template"] = h.Template
	}
	return message
}

// CardTitle 卡片标题内容
type CardTitle struct {
	// Content 内容
	Content string
	// i18n 国际化内容
	I18n map[string]string
}

// NewCardTitle create CardTitle
func NewCardTitle() *CardTitle {
	return &CardTitle{
		I18n: map[string]string{},
	}
}

// SetContent set CardTitle.Content
func (t *CardTitle) SetContent(content string) *CardTitle {
	t.Content = content
	return t
}

// SetI18n set CardTitle.I18n
func (t *CardTitle) SetI18n(i18n map[string]string) *CardTitle {
	t.I18n = i18n
	return t
}

// AddI18n add CardTitle.I18n
func (t *CardTitle) AddI18n(key, value string) *CardTitle {
	t.I18n[key] = value
	return t
}
func (t *CardTitle) ToMessageMap() map[string]interface{} {
	message := map[string]interface{}{}
	message["tag"] = "plain_text"
	if len(t.I18n) > 0 {
		message["i18n"] = t.I18n
	} else {
		message["content"] = t.Content
	}
	return message
}

// ElementsContent 用于定义卡片正文内容
type ElementsContent struct {
	// Elements 正文内容, DivCardContent , HrCardContent, MarkdownCardContent, ImgCardContent, NoteCardContent
	Elements []Message
}

// NewElementsContent create ElementsContent
func NewElementsContent() *ElementsContent {
	return &ElementsContent{
		Elements: []Message{},
	}
}

// SetElements set ElementsContent.Elements
func (c *ElementsContent) SetElements(elements []Message) *ElementsContent {
	c.Elements = elements
	return c
}

// AddElement set ElementsContent.Elements
func (c *ElementsContent) AddElement(element Message) *ElementsContent {
	c.Elements = append(c.Elements, element)
	return c
}

// AddElements set ElementsContent.Elements
func (c *ElementsContent) AddElements(elements ...Message) *ElementsContent {
	c.Elements = append(c.Elements, elements...)
	return c
}

func (c *ElementsContent) ToMessageMap() []map[string]interface{} {
	var elementMessage []map[string]interface{}
	for _, element := range c.Elements {
		elementMessage = append(elementMessage, element.ToMessageMap())
	}
	return elementMessage
}

// I18nElementsContent 国际化正文内容
type I18nElementsContent struct {
	// I18n 国际化 如:zh-cn
	I18n string
	// Elements 内容
	Elements *ElementsContent
}

// NewI18nElementsContent create I18nElementsContent
func NewI18nElementsContent() *I18nElementsContent {
	return &I18nElementsContent{}
}

// SetI18n set I18nElementsContent.i18n
func (c *I18nElementsContent) SetI18n(i18n string) *I18nElementsContent {
	c.I18n = i18n
	return c
}

// SetElementContent set I18nElementsContent.Elements
func (c *I18nElementsContent) SetElementContent(element *ElementsContent) *I18nElementsContent {
	c.Elements = element
	return c
}
func (c *I18nElementsContent) ToMessageMap() map[string]interface{} {
	message := map[string]interface{}{}
	message[c.I18n] = c.Elements.ToMessageMap()
	return message
}
