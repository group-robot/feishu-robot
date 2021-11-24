package feishu

type InteractiveMessage struct {
	BaseMessage
	CardJsonStr string
	Link        *CardLink
	Header      *CardHeader
	Config      *CardConfig
	Elements    *ElementsContent
	I18nElement []*I18nElementsContent
}

func NewInteractiveMessage() *InteractiveMessage {
	return &InteractiveMessage{
		I18nElement: []*I18nElementsContent{},
	}
}
func (m *InteractiveMessage) GetMsgType() MsgType {
	return MsgTypeInteractive
}
func (m *InteractiveMessage) SetHeader(header *CardHeader) *InteractiveMessage {
	m.Header = header
	return m
}
func (m *InteractiveMessage) SetConfig(config *CardConfig) *InteractiveMessage {
	m.Config = config
	return m
}
func (m *InteractiveMessage) SetElements(elements *ElementsContent) *InteractiveMessage {
	m.Elements = elements
	return m
}
func (m *InteractiveMessage) SetI18nElement(i18nElements []*I18nElementsContent) *InteractiveMessage {
	m.I18nElement = i18nElements
	return m
}
func (m *InteractiveMessage) SetCardJsonStr(jsonStr string) *InteractiveMessage {
	m.CardJsonStr = jsonStr
	return m
}
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

type CardConfig struct {
	EnableForward  bool
	WideScreenMode bool
}

func NewCardConfig() *CardConfig {
	return &CardConfig{
		EnableForward:  true,
		WideScreenMode: true,
	}
}
func (c *CardConfig) SetEnableForward(enableForward bool) *CardConfig {
	c.EnableForward = enableForward
	return c
}
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

type CardHeader struct {
	Title    *CardTitle
	Template string
}

func NewCardHeader() *CardHeader {
	return new(CardHeader)
}
func (h *CardHeader) SetTitle(title *CardTitle) *CardHeader {
	h.Title = title
	return h
}
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

type CardTitle struct {
	Content string
	I18n    map[string]string
}

func NewCardTitle() *CardTitle {
	return &CardTitle{
		I18n: map[string]string{},
	}
}
func (t *CardTitle) SetContent(content string) *CardTitle {
	t.Content = content
	return t
}
func (t *CardTitle) SetI18n(i18n map[string]string) *CardTitle {
	t.I18n = i18n
	return t
}
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

type ElementsContent struct {
	Elements []Message
}

func NewElementsContent() *ElementsContent {
	return &ElementsContent{
		Elements: []Message{},
	}
}
func (c *ElementsContent) SetElements(elements []Message) *ElementsContent {
	c.Elements = elements
	return c
}
func (c *ElementsContent) AddElement(element Message) *ElementsContent {
	c.Elements = append(c.Elements, element)
	return c
}
func (c *ElementsContent) ToMessageMap() []map[string]interface{} {
	var elementMessage []map[string]interface{}
	for _, element := range c.Elements {
		elementMessage = append(elementMessage, element.ToMessageMap())
	}
	return elementMessage
}

type I18nElementsContent struct {
	I18n     string
	Elements *ElementsContent
}

func NewI18nElementsContent() *I18nElementsContent {
	return &I18nElementsContent{}
}
func (c *I18nElementsContent) SetI18n(i18n string) *I18nElementsContent {
	c.I18n = i18n
	return c
}

func (c *I18nElementsContent) SetElementContent(element *ElementsContent) *I18nElementsContent {
	c.Elements = element
	return c
}
func (c *I18nElementsContent) ToMessageMap() map[string]interface{} {
	message := map[string]interface{}{}
	message[c.I18n] = c.Elements.ToMessageMap()
	return message
}
