package feishu

// 可嵌入非交互元素

type CardTextTag string

const (
	PlainText CardTextTag = "plain_text"
	LarkMd    CardTextTag = "lark_md"
)

type Filed struct {
	IsShort bool
	Text    *Text
}

func NewFiled() *Filed {
	return new(Filed)
}

func (f *Filed) SetShort(short bool) *Filed {
	f.IsShort = short
	return f
}
func (f *Filed) Short() *Filed {
	f.IsShort = true
	return f
}
func (f *Filed) SetText(text *Text) *Filed {
	f.Text = text
	return f
}
func (f *Filed) ToMessageMap() map[string]interface{} {
	message := map[string]interface{}{}
	message["is_short"] = f.IsShort
	message["text"] = f.Text.ToMessageMap()
	return message
}

type Text struct {
	tag     CardTextTag
	Content string
	lines   int
}

func NewText() *Text {
	return new(Text)
}
func (t *Text) SetContent(content string) *Text {
	t.Content = content
	return t
}
func (t *Text) SetLines(lines int) *Text {
	t.lines = lines
	return t
}
func (t *Text) Tag(tag CardTextTag) *Text {
	t.tag = tag
	return t
}

func (t *Text) ToMessageMap() map[string]interface{} {
	message := map[string]interface{}{}
	message["tag"] = t.tag
	message["content"] = t.Content
	if t.lines > 0 {
		message["lines"] = t.lines
	}
	return message
}

type Image struct {
	ImgKey  string
	Alt     Text
	PreView bool
}

func NewImag() *Image {
	return new(Image)
}
func (image *Image) SetImgKey(imgKey string) *Image {
	image.ImgKey = imgKey
	return image
}
func (image *Image) SetAlt(alt Text) *Image {
	image.Alt = alt
	return image
}
func (image *Image) SetPreView(preview bool) *Image {
	image.PreView = preview
	return image
}

func (image *Image) ToMessageMap() map[string]interface{} {
	message := map[string]interface{}{}
	message["tag"] = "img"
	message["img_key"] = image.ImgKey
	message["alt"] = image.Alt.ToMessageMap()
	if image.PreView {
		message["preview"] = image.PreView
	}
	return message
}

// 交互

type LayoutType string

const (
	BisectedLayout   LayoutType = "bisected"
	TrisectionLayout LayoutType = "trisection"
	FlowLayout       LayoutType = "flow"
)

type Action interface {
	Message
	GetConfirm() *Confirm
}

type ActionModule struct {
	Actions []Action
	Layout  LayoutType
}

func NewActionModule() *ActionModule {
	return &ActionModule{
		Actions: []Action{},
		Layout:  FlowLayout,
	}
}
func (m *ActionModule) SetActions(actions []Action) *ActionModule {
	m.Actions = actions
	return m
}
func (m *ActionModule) AddAction(action Action) *ActionModule {
	m.Actions = append(m.Actions, action)
	return m
}
func (m *ActionModule) AddActions(actions ...Action) *ActionModule {
	m.Actions = append(m.Actions, actions...)
	return m
}

func (m *ActionModule) SetLayout(layout LayoutType) *ActionModule {
	m.Layout = layout
	return m
}
func (m *ActionModule) ToMessageMap() map[string]interface{} {
	message := map[string]interface{}{}
	message["tag"] = "action"
	var actionMessage []map[string]interface{}
	for _, action := range m.Actions {
		actionMessage = append(actionMessage, action.ToMessageMap())
	}
	message["actions"] = actionMessage
	message["layout"] = m.Layout
	return message

}

// 交互组件

type PickerTag string

const (
	datePicker     PickerTag = "date_picker"
	pickerTime     PickerTag = "picker_time"
	pickerDatetime PickerTag = "picker_datetime"
)

type DatePicker struct {
	Tag             *PickerTag
	InitialDate     string
	InitialTime     string
	InitialDatetime string
	Placeholder     string
	Value           string
	Confirm         *Confirm
}

func NewDatePicker() *DatePicker {
	return new(DatePicker)
}
func (picker *DatePicker) SetTag(tag *PickerTag) *DatePicker {
	picker.Tag = tag
	return picker
}
func (picker *DatePicker) SetInitialDate(initialDate string) *DatePicker {
	picker.InitialDate = initialDate
	return picker
}
func (picker *DatePicker) SetInitialTime(initialTime string) *DatePicker {
	picker.InitialTime = initialTime
	return picker
}
func (picker *DatePicker) SetInitialDatetime(initialDatetime string) *DatePicker {
	picker.InitialDatetime = initialDatetime
	return picker
}
func (picker *DatePicker) SetPlaceholder(placeholder string) *DatePicker {
	picker.Placeholder = placeholder
	return picker
}
func (picker *DatePicker) SetValue(value string) *DatePicker {
	picker.Value = value
	return picker
}
func (picker *DatePicker) SetConfirm(confirm *Confirm) *DatePicker {
	picker.Confirm = confirm
	return picker
}
func (picker *DatePicker) GetConfirm() *Confirm {
	return picker.Confirm
}
func (picker *DatePicker) ToMessageMap() map[string]interface{} {
	message := map[string]interface{}{}
	message["tag"] = picker.Tag
	if picker.InitialDate != "" {
		message["initial_date"] = picker.InitialDate
	}
	if picker.InitialTime != "" {
		message["initial_time"] = picker.InitialTime
	}
	if picker.InitialDatetime != "" {
		message["initial_datetime"] = picker.InitialDatetime
	}
	if picker.Placeholder != "" {
		placeholderMessage := map[string]interface{}{
			"tag":     "plain_text",
			"content": picker.Placeholder,
		}
		message["placeholder"] = placeholderMessage
	}
	if picker.Value != "" {
		message["value"] = picker.Value
	}
	if picker.Confirm != nil {
		message["confirm"] = picker.GetConfirm().ToMessageMap()
	}
	return message
}

type Overflow struct {
	Options []*Option
	Value   string
	Confirm *Confirm
}

func NewOverflow() *Overflow {
	return &Overflow{
		Options: []*Option{},
	}
}
func (o *Overflow) SetOptions(options []*Option) *Overflow {
	o.Options = options
	return o
}
func (o *Overflow) AddOption(option *Option) *Overflow {
	o.Options = append(o.Options, option)
	return o
}
func (o *Overflow) AddOptions(options ...*Option) *Overflow {
	o.Options = append(o.Options, options...)
	return o
}

func (o *Overflow) SetValue(value string) *Overflow {
	o.Value = value
	return o
}
func (o *Overflow) SetConfirm(confirm *Confirm) *Overflow {
	o.Confirm = confirm
	return o
}
func (o *Overflow) GetConfirm() *Confirm {
	return o.Confirm
}
func (o *Overflow) ToMessageMap() map[string]interface{} {
	var optionsMessage []map[string]interface{}
	for _, v := range o.Options {
		optionsMessage = append(optionsMessage, v.ToMessageMap())
	}

	message := map[string]interface{}{}
	message["tag"] = "overflow"
	message["options"] = optionsMessage
	if o.Value != "" {
		message["value"] = o.Value
	}
	if o.Confirm != nil {
		message["confirm"] = o.GetConfirm().ToMessageMap()
	}
	return message
}

type SelectTag string

const (
	SelectStatic SelectTag = "select_static"
	SelectPerson SelectTag = "select_person"
)

type SelectMenu struct {
	Tag           SelectTag
	Placeholder   *Text
	InitialOption string
	Options       []*Option
	Value         string
	Confirm       *Confirm
}

func NewSelectMenu() *SelectMenu {
	return &SelectMenu{
		Options: []*Option{},
	}
}
func (m *SelectMenu) SetTag(tag SelectTag) *SelectMenu {
	m.Tag = tag
	return m
}
func (m *SelectMenu) SetPlaceholder(placeholder *Text) *SelectMenu {
	m.Placeholder = placeholder
	return m
}
func (m *SelectMenu) SetInitialOption(initialOption string) *SelectMenu {
	m.InitialOption = initialOption
	return m
}
func (m *SelectMenu) SetOptions(options []*Option) *SelectMenu {
	m.Options = options
	return m
}
func (m *SelectMenu) AddOption(option *Option) *SelectMenu {
	m.Options = append(m.Options, option)
	return m
}
func (m *SelectMenu) AddOptions(options ...*Option) *SelectMenu {
	m.Options = append(m.Options, options...)
	return m
}

func (m *SelectMenu) SetValue(value string) *SelectMenu {
	m.Value = value
	return m
}
func (m *SelectMenu) SetConfirm(confirm *Confirm) *SelectMenu {
	m.Confirm = confirm
	return m
}
func (m *SelectMenu) GetConfirm() *Confirm {
	return m.Confirm
}
func (m *SelectMenu) ToMessageMap() map[string]interface{} {
	message := map[string]interface{}{}
	message["tag"] = m.Tag
	if m.Placeholder != nil {
		message["placeholder"] = m.Placeholder.ToMessageMap()
	}
	if m.InitialOption != "" {
		message["initial_option"] = m.InitialOption
	}
	if m.Options != nil {
		var optionsMessage []map[string]interface{}
		for _, v := range m.Options {
			optionsMessage = append(optionsMessage, v.ToMessageMap())
		}
		message["options"] = optionsMessage
	}
	if m.Value != "" {
		message["value"] = m.Value
	}
	if m.Confirm != nil {
		message["confirm"] = m.GetConfirm().ToMessageMap()
	}
	return message
}

type ButtonType string

const (
	DefaultButton ButtonType = "default"
	PrimaryButton ButtonType = "primary"
	DangerButton  ButtonType = "danger"
)

type Button struct {
	Text     *Text
	Url      string
	MultiUrl *Url
	Type     ButtonType
	Value    map[string]string
	Confirm  *Confirm
}

func NewButton() *Button {
	return &Button{
		Type:  DefaultButton,
		Value: map[string]string{},
	}
}
func (b *Button) SetText(text *Text) *Button {
	b.Text = text
	return b
}
func (b *Button) SetUrl(url string) *Button {
	b.Url = url
	return b
}
func (b *Button) SetMultiUrl(multiUrl *Url) *Button {
	b.MultiUrl = multiUrl
	return b
}
func (b *Button) SetType(buttonType ButtonType) *Button {
	b.Type = buttonType
	return b
}
func (b *Button) SetValue(key, value string) *Button {
	b.Value[key] = value
	return b
}
func (b *Button) SetConfirm(confirm *Confirm) *Button {
	b.Confirm = confirm
	return b
}
func (b *Button) GetConfirm() *Confirm {
	return b.Confirm
}
func (b *Button) ToMessageMap() map[string]interface{} {
	message := map[string]interface{}{}
	message["tag"] = "button"
	message["text"] = b.Text.ToMessageMap()
	if b.Url != "" {
		message["url"] = b.Url
	}
	if b.MultiUrl != nil {
		message["multi_url"] = b.MultiUrl.ToMessageMap()
	}
	message["type"] = b.Type
	if b.Value != nil {
		message["value"] = b.Value
	}
	if b.Confirm != nil {
		message["confirm"] = b.GetConfirm().ToMessageMap()
	}
	return message
}

// 可嵌入交互元素

type CardLink struct {
	Url
}

func NewCardLink() *CardLink {
	return new(CardLink)
}
func (link *CardLink) SetUrl(u string) *CardLink {
	link.Url.Url = u
	return link
}
func (link *CardLink) SetAndroidUrl(androidUrl string) *CardLink {
	link.AndroidUrl = androidUrl
	return link
}
func (link *CardLink) SetIosUrl(iosUrl string) *CardLink {
	link.IosUrl = iosUrl
	return link
}
func (link *CardLink) SetPcUrl(pcUrl string) *CardLink {
	link.PcUrl = pcUrl
	return link
}
func (link *CardLink) ToMessageMap() map[string]interface{} {
	urlMessage := map[string]interface{}{}
	urlMessage["url"] = link.Url
	urlMessage["android_url"] = link.AndroidUrl
	urlMessage["ios_url"] = link.IosUrl
	urlMessage["pc_url"] = link.PcUrl
	return urlMessage
}

type Url struct {
	Url        string
	AndroidUrl string
	IosUrl     string
	PcUrl      string
}

func NewUrl() *Url {
	return new(Url)
}

func (url *Url) SetUrl(u string) *Url {
	url.Url = u
	return url
}
func (url *Url) SetAndroidUrl(androidUrl string) *Url {
	url.AndroidUrl = androidUrl
	return url
}
func (url *Url) SetIosUrl(iosUrl string) *Url {
	url.IosUrl = iosUrl
	return url
}
func (url *Url) SetPcUrl(pcUrl string) *Url {
	url.PcUrl = pcUrl
	return url
}
func (url *Url) ToMessageMap() map[string]interface{} {
	message := map[string]interface{}{}
	message["url"] = url.Url
	message["android_url"] = url.AndroidUrl
	message["ios_url"] = url.IosUrl
	message["pc_url"] = url.PcUrl
	return message
}

type Confirm struct {
	Title *Text
	Text  *Text
}

func NewConfirm() *Confirm {
	return new(Confirm)
}
func (c *Confirm) SetTitle(title *Text) *Confirm {
	title.Tag(PlainText)
	c.Title = title
	return c
}
func (c *Confirm) SetText(text *Text) *Confirm {
	text.Tag(PlainText)
	c.Title = text
	return c
}
func (c *Confirm) ToMessageMap() map[string]interface{} {
	message := map[string]interface{}{}
	message["title"] = c.Title.ToMessageMap()
	message["text"] = c.Text.ToMessageMap()
	return message
}

type Option struct {
	Text     *Text
	value    string
	Url      string
	MultiUrl *Url
}

func NewOption() *Option {
	return new(Option)
}
func (o *Option) SetText(text *Text) *Option {
	o.Text = text
	return o
}
func (o *Option) SetValue(value string) *Option {
	o.value = value
	return o
}
func (o *Option) SetUrl(Url string) *Option {
	o.Url = Url
	return o
}
func (o *Option) SetMultiUrl(multiUrl *Url) *Option {
	o.MultiUrl = multiUrl
	return o
}
func (o *Option) ToMessageMap() map[string]interface{} {
	message := map[string]interface{}{}
	message["text"] = o.Text.ToMessageMap()
	message["value"] = o.value
	if o.value != "" {
		message["url"] = o.Url
	}
	if o.MultiUrl != nil {
		message["multi_url"] = o.MultiUrl.ToMessageMap()
	}
	return message
}
