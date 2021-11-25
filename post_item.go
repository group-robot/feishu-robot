package feishu

// 可嵌入非交互元素

// CardTextTag Text 元素标签
type CardTextTag string

const (
	// PlainText 支持普通文本内容
	PlainText CardTextTag = "plain_text"
	// LarkMd 支持部分markDown语法展示文本内容
	LarkMd CardTextTag = "lark_md"
)

// Field field对象可用于内容模块的field字段，通过"is_short"字段控制是否并排布局
type Field struct {
	// 是否并排布局
	IsShort bool
	// Text 国际化文本内容
	Text *Text
}

// NewField create Field
func NewField() *Field {
	return new(Field)
}

// SetShort set Field.IsShort
func (f *Field) SetShort(short bool) *Field {
	f.IsShort = short
	return f
}

// Short set Field.IsShort is true
func (f *Field) Short() *Field {
	f.IsShort = true
	return f
}

// SetText set Field.Text
func (f *Field) SetText(text *Text) *Field {
	f.Text = text
	return f
}
func (f *Field) ToMessageMap() map[string]interface{} {
	message := map[string]interface{}{}
	message["is_short"] = f.IsShort
	message["text"] = f.Text.ToMessageMap()
	return message
}

// Text 作为text对象被使用
type Text struct {
	// Tag 	元素标签
	Tag CardTextTag
	// Content 	文本内容
	Content string
	// Lines 内容显示行数
	Lines int
}

// NewText create Text
func NewText() *Text {
	return new(Text)
}

// SetContent set Text.Content
func (t *Text) SetContent(content string) *Text {
	t.Content = content
	return t
}

// SetLines set Text.Lines
func (t *Text) SetLines(lines int) *Text {
	t.Lines = lines
	return t
}

// SetTag set Text.Tag
func (t *Text) SetTag(tag CardTextTag) *Text {
	t.Tag = tag
	return t
}

func (t *Text) ToMessageMap() map[string]interface{} {
	message := map[string]interface{}{}
	message["tag"] = t.Tag
	message["content"] = t.Content
	if t.Lines > 0 {
		message["lines"] = t.Lines
	}
	return message
}

// Image 属于内容元素的一种，可用于内容块的extra字段和备注块的elements字段。
type Image struct {
	// ImgKey 图片资源
	ImgKey string
	// Alt 图片hover说明
	Alt *Text
	// PreView 点击后是否放大图片，缺省为true
	PreView bool
}

// NewImag create Image
func NewImag() *Image {
	return &Image{
		PreView: true,
	}
}

// SetImgKey set Image.ImgKey
func (image *Image) SetImgKey(imgKey string) *Image {
	image.ImgKey = imgKey
	return image
}

// SetAlt set Image.Alt
func (image *Image) SetAlt(alt *Text) *Image {
	image.Alt = alt
	return image
}

// SetPreView set Image.PreView
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

// LayoutType 交互元素布局
type LayoutType string

const (
	// BisectedLayout 二等分布局
	BisectedLayout LayoutType = "bisected"
	// TrisectionLayout 三等分布局
	TrisectionLayout LayoutType = "trisection"
	// FlowLayout 流式布局
	FlowLayout LayoutType = "flow"
)

// Action 交互模块
type Action interface {
	Message
	// GetConfirm 获取二次确认的弹框
	GetConfirm() *Confirm
}

// ActionModule 交互模块
type ActionModule struct {
	// Actions 交互元素
	Actions []Action
	// Layout 交互元素布局
	Layout LayoutType
}

// NewActionModule create ActionModule
func NewActionModule() *ActionModule {
	return &ActionModule{
		Actions: []Action{},
		Layout:  FlowLayout,
	}
}

// SetActions set ActionModule.Actions
func (m *ActionModule) SetActions(actions []Action) *ActionModule {
	m.Actions = actions
	return m
}

// AddAction add ActionModule.Actions
func (m *ActionModule) AddAction(action Action) *ActionModule {
	m.Actions = append(m.Actions, action)
	return m
}

// AddActions add ActionModule.Actions
func (m *ActionModule) AddActions(actions ...Action) *ActionModule {
	m.Actions = append(m.Actions, actions...)
	return m
}

// SetLayout set ActionModule.Layout
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

// PickerTag DatePicker 元素标签
type PickerTag string

const (
	// DatePickerTag  日期
	DatePickerTag PickerTag = "date_picker"
	// PickerTimeTag 时间
	PickerTimeTag PickerTag = "picker_time"
	// PickerDatetimeTag 日期+时间
	PickerDatetimeTag PickerTag = "picker_datetime"
)

// DatePicker 提供时间选择的功能,
//datePicker属于交互元素的一种，可用于内容块的extra字段和交互块的actions字段。
type DatePicker struct {
	// Tag 元素标签
	Tag *PickerTag
	// InitialDate 日期模式的初始值,格式"yyyy-MM-dd"
	InitialDate string
	// InitialTime 	时间模式的初始值,格式"HH:mm"
	InitialTime string
	// InitialDatetime 日期时间模式的初始值,格式"yyyy-MM-dd HH:mm"
	InitialDatetime string
	// Placeholder 占位符，无初始值时必填
	Placeholder string
	// Value 用户选定后返回业务方的数据
	Value string
	// Confirm 	二次确认的弹框
	Confirm *Confirm
}

// NewDatePicker create DatePicker
func NewDatePicker() *DatePicker {
	return new(DatePicker)
}

// SetTag set DatePicker.Tag
func (picker *DatePicker) SetTag(tag *PickerTag) *DatePicker {
	picker.Tag = tag
	return picker
}

// SetInitialDate set DatePicker.InitialDate
func (picker *DatePicker) SetInitialDate(initialDate string) *DatePicker {
	picker.InitialDate = initialDate
	return picker
}

// SetInitialTime set DatePicker.InitialTime
func (picker *DatePicker) SetInitialTime(initialTime string) *DatePicker {
	picker.InitialTime = initialTime
	return picker
}

// SetInitialDatetime set DatePicker.InitialDatetime
func (picker *DatePicker) SetInitialDatetime(initialDatetime string) *DatePicker {
	picker.InitialDatetime = initialDatetime
	return picker
}

// SetPlaceholder set DatePicker.Placeholder
func (picker *DatePicker) SetPlaceholder(placeholder string) *DatePicker {
	picker.Placeholder = placeholder
	return picker
}

// SetValue set DatePicker.Value
func (picker *DatePicker) SetValue(value string) *DatePicker {
	picker.Value = value
	return picker
}

// SetConfirm set DatePicker.Confirm
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

// Overflow 提供折叠的按钮型菜单,overflow属于交互元素的一种，可用于内容块的extra字段和交互块的actions字段。
type Overflow struct {
	// Options 	待选选项
	Options []*Option
	// Value 用户选定后返回业务方的数据
	Value string
	// Confirm 二次确认的弹框
	Confirm *Confirm
}

// NewOverflow create Overflow
func NewOverflow() *Overflow {
	return &Overflow{
		Options: []*Option{},
	}
}

// SetOptions set Overflow.Options
func (o *Overflow) SetOptions(options []*Option) *Overflow {
	o.Options = options
	return o
}

// AddOption add Overflow.Options
func (o *Overflow) AddOption(option *Option) *Overflow {
	o.Options = append(o.Options, option)
	return o
}

// AddOptions add Overflow.Options
func (o *Overflow) AddOptions(options ...*Option) *Overflow {
	o.Options = append(o.Options, options...)
	return o
}

// SetValue set Overflow.Value
func (o *Overflow) SetValue(value string) *Overflow {
	o.Value = value
	return o
}

// SetConfirm set Overflow.Confirm
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

// SelectTag SelectMenu 元素标签
type SelectTag string

const (
	// SelectStatic 选项模式
	SelectStatic SelectTag = "select_static"
	// SelectPerson 选人模式
	SelectPerson SelectTag = "select_person"
)

// SelectMenu 提供选项菜单的功能,selectMenu属于交互元素的一种，可用于内容块的extra字段和交互块的actions字段。
type SelectMenu struct {
	// Tag 元素标签
	Tag SelectTag
	// Placeholder 占位符，无默认选项时必须有
	Placeholder *Text
	// InitialOption	默认选项的value字段值
	InitialOption string
	// Options 	待选选项
	Options []*Option
	// Value 用户选定后返回业务方的数据
	Value string
	// Confirm	二次确认的弹框
	Confirm *Confirm
}

// NewSelectMenu create SelectMenu
func NewSelectMenu() *SelectMenu {
	return &SelectMenu{
		Options: []*Option{},
	}
}

// SetTag set SelectMenu.Tag
func (m *SelectMenu) SetTag(tag SelectTag) *SelectMenu {
	m.Tag = tag
	return m
}

// SetPlaceholder set SelectMenu.Placeholder
func (m *SelectMenu) SetPlaceholder(placeholder *Text) *SelectMenu {
	m.Placeholder = placeholder
	return m
}

// SetInitialOption set SelectMenu.InitialOption
func (m *SelectMenu) SetInitialOption(initialOption string) *SelectMenu {
	m.InitialOption = initialOption
	return m
}

// SetOptions set SelectMenu.Options
func (m *SelectMenu) SetOptions(options []*Option) *SelectMenu {
	m.Options = options
	return m
}

// AddOption add SelectMenu.Options
func (m *SelectMenu) AddOption(option *Option) *SelectMenu {
	m.Options = append(m.Options, option)
	return m
}

// AddOptions add SelectMenu.Options
func (m *SelectMenu) AddOptions(options ...*Option) *SelectMenu {
	m.Options = append(m.Options, options...)
	return m
}

// SetValue set SelectMenu.Value
func (m *SelectMenu) SetValue(value string) *SelectMenu {
	m.Value = value
	return m
}

// SetConfirm set SelectMenu.Confirm
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

// ButtonType Button 配置按钮样式
type ButtonType string

const (
	// DefaultButton default
	DefaultButton ButtonType = "default"
	// PrimaryButton primary
	PrimaryButton ButtonType = "primary"
	// DangerButton danger
	DangerButton ButtonType = "danger"
)

// Button button 属于交互元素的一种，可用于内容块的extra字段和交互块的actions字段。
type Button struct {
	// Text	按钮中的文本
	Text *Text
	// Url  跳转链接，和 Button.MultiUrl 互斥
	Url string
	// MultiUrl 多端跳转链接
	MultiUrl *Url
	// Type 配置按钮样式，默认为"default"
	Type ButtonType
	// Value 点击后返回业务方
	Value map[string]string
	// Confirm	二次确认的弹框
	Confirm *Confirm
}

// NewButton create Button
func NewButton() *Button {
	return &Button{
		Type:  DefaultButton,
		Value: map[string]string{},
	}
}

// SetText set Button.Text
func (b *Button) SetText(text *Text) *Button {
	b.Text = text
	return b
}

// SetUrl set Button.Url
func (b *Button) SetUrl(url string) *Button {
	b.Url = url
	return b
}

// SetMultiUrl set Button.MultiUrl
func (b *Button) SetMultiUrl(multiUrl *Url) *Button {
	b.MultiUrl = multiUrl
	return b
}

// SetType set Button.Type
func (b *Button) SetType(buttonType ButtonType) *Button {
	b.Type = buttonType
	return b
}

// SetValue set Button.Value
func (b *Button) SetValue(key, value string) *Button {
	b.Value[key] = value
	return b
}

// SetConfirm set Button.Confirm
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

// CardLink card_link用于指定卡片整体的点击跳转链接，可以配置默认链接，也可以分别配置不同终端的链接
type CardLink struct {
	Url
}

// NewCardLink create CardLink
func NewCardLink() *CardLink {
	return new(CardLink)
}

// SetUrl 设置 默认的链接地址
func (link *CardLink) SetUrl(u string) *CardLink {
	link.Url.Url = u
	return link
}

// SetAndroidUrl 设置 	Android 端的链接地址
func (link *CardLink) SetAndroidUrl(androidUrl string) *CardLink {
	link.AndroidUrl = androidUrl
	return link
}

// SetIosUrl 设置 	iOS 端的链接地址
func (link *CardLink) SetIosUrl(iosUrl string) *CardLink {
	link.IosUrl = iosUrl
	return link
}

// SetPcUrl 设置 PC 端的链接地址
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

// Url url对象用作多端差异跳转链接
type Url struct {
	// Url 	默认跳转链接，
	Url string
	// AndroidUrl 安卓端跳转链接
	AndroidUrl string
	// IosUrl ios端跳转链接
	IosUrl string
	// PcUrl pc端跳转链接
	PcUrl string
}

// NewUrl create Url
func NewUrl() *Url {
	return new(Url)
}

// SetUrl set Url.Url
func (url *Url) SetUrl(u string) *Url {
	url.Url = u
	return url
}

// SetAndroidUrl set Url.AndroidUrl
func (url *Url) SetAndroidUrl(androidUrl string) *Url {
	url.AndroidUrl = androidUrl
	return url
}

// SetIosUrl set Url.IosUrl
func (url *Url) SetIosUrl(iosUrl string) *Url {
	url.IosUrl = iosUrl
	return url
}

// SetPcUrl set Url.PcUrl
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

// Confirm 用于交互元素的二次确认
type Confirm struct {
	// Title 弹框标题
	Title *Text
	// Text 弹框内容
	Text *Text
}

// NewConfirm create Confirm
func NewConfirm() *Confirm {
	return new(Confirm)
}

// SetTitle set Confirm.Title
func (c *Confirm) SetTitle(title *Text) *Confirm {
	title.SetTag(PlainText)
	c.Title = title
	return c
}

// SetText set Confirm.Title
func (c *Confirm) SetText(text *Text) *Confirm {
	text.SetTag(PlainText)
	c.Title = text
	return c
}
func (c *Confirm) ToMessageMap() map[string]interface{} {
	message := map[string]interface{}{}
	message["title"] = c.Title.ToMessageMap()
	message["text"] = c.Text.ToMessageMap()
	return message
}

// Option
//
//作为selectMenu的选项对象
//
//作为overflow的选项对象
type Option struct {
	// Text 	选项显示内容，非待选人员时必填
	Text *Text
	// Value 选项选中后返回业务方的数据
	Value string
	// Url *仅支持overflow，跳转指定链接，和multi_url字段互斥
	Url string
	// MultiUrl *仅支持overflow，跳转对应链接，和url字段互斥
	MultiUrl *Url
}

// NewOption create Option
func NewOption() *Option {
	return new(Option)
}

//SetText set Option.Text
func (o *Option) SetText(text *Text) *Option {
	o.Text = text
	return o
}

//SetValue set Option.Value
func (o *Option) SetValue(value string) *Option {
	o.Value = value
	return o
}

//SetUrl set Option.Url
func (o *Option) SetUrl(Url string) *Option {
	o.Url = Url
	return o
}

//SetMultiUrl set Option.MultiUrl
func (o *Option) SetMultiUrl(multiUrl *Url) *Option {
	o.MultiUrl = multiUrl
	return o
}
func (o *Option) ToMessageMap() map[string]interface{} {
	message := map[string]interface{}{}
	message["text"] = o.Text.ToMessageMap()
	message["value"] = o.Value
	if o.Value != "" {
		message["url"] = o.Url
	}
	if o.MultiUrl != nil {
		message["multi_url"] = o.MultiUrl.ToMessageMap()
	}
	return message
}
