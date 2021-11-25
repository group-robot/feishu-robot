package feishu

// DivCardContent 内容模块
type DivCardContent struct {
	// Text 单个文本展示，和fields至少要有一个
	Text *Text
	// Fields 多个文本展示，和text至少要有一个
	Fields []*Field
	// Extra 	附加的元素展示在文本内容右侧。可附加的元素包括 Image , Button , SelectMenu , Overflow , DatePicker
	Extra Message
}

// NewDivCardContent crete DivCardContent
func NewDivCardContent() *DivCardContent {
	return &DivCardContent{
		Fields: []*Field{},
	}
}

// SetText set DivCardContent.Text
func (c *DivCardContent) SetText(text *Text) *DivCardContent {
	c.Text = text
	return c
}

// SetFields set DivCardContent.Text
func (c *DivCardContent) SetFields(fields []*Field) *DivCardContent {
	c.Fields = fields
	return c
}

// AddField add DivCardContent.Fields
func (c *DivCardContent) AddField(filed *Field) *DivCardContent {
	c.Fields = append(c.Fields, filed)
	return c
}

// AddFields add DivCardContent.Fields
func (c *DivCardContent) AddFields(fields ...*Field) *DivCardContent {
	c.Fields = append(c.Fields, fields...)
	return c
}

// SetExtra set DivCardContent.Extra
func (c *DivCardContent) SetExtra(extra Message) *DivCardContent {
	c.Extra = extra
	return c
}
func (c *DivCardContent) ToMessageMap() map[string]interface{} {
	message := map[string]interface{}{}
	message["tag"] = "div"
	if c.Text != nil {
		message["text"] = c.Text.ToMessageMap()
	}
	var fieldsMessage []map[string]interface{}
	for _, field := range c.Fields {
		fieldsMessage = append(fieldsMessage, field.ToMessageMap())
	}
	message["fields"] = fieldsMessage
	if c.Extra != nil {
		message["extra"] = c.Extra.ToMessageMap()
	}
	return message
}

// MarkdownCardContent Markdown模块
type MarkdownCardContent struct {
	// Content 	markdown标签内容
	Content string
	// Href 差异化跳转
	Href *Url
}

// NewMarkdownCardContent create MarkdownCardContent
func NewMarkdownCardContent() *MarkdownCardContent {
	return new(MarkdownCardContent)
}

// SetContent set MarkdownCardContent.Content
func (c *MarkdownCardContent) SetContent(content string) *MarkdownCardContent {
	c.Content = content
	return c
}

// SetHref set MarkdownCardContent.Href
func (c *MarkdownCardContent) SetHref(href *Url) *MarkdownCardContent {
	c.Href = href
	return c
}
func (c *MarkdownCardContent) ToMessageMap() map[string]interface{} {
	message := map[string]interface{}{}
	message["tag"] = "markdown"
	message["content"] = c.Content
	if c.Href != nil {
		message["href"] = c.Href
	}
	return message
}

// HrCardContent 分割线模块
type HrCardContent struct {
}

// NewHrCardContent create HrCardContent
func NewHrCardContent() *HrCardContent {
	return new(HrCardContent)
}
func (c *HrCardContent) ToMessageMap() map[string]interface{} {
	message := map[string]interface{}{}
	message["tag"] = "hr"
	return message
}

// ImgModType 图片显示模式
type ImgModType string

const (
	// FitHorizontalMod 平铺模式
	FitHorizontalMod ImgModType = "fit_horizontal"
	// CropCenterMod 居中裁剪模式
	CropCenterMod ImgModType = "crop_center"
)

// ImgCardContent 图片模块
type ImgCardContent struct {
	// ImgKey 图片资源
	ImgKey string
	// Alt hover图片时弹出的Tips文案 ,content取值为空时则不展示
	Alt *Text
	// Title 图片标题
	Title *Text
	// CustomWidth	自定义图片的最大展示宽度。默认展示宽度撑满卡片的通栏图片
	CustomWidth int
	// CompactWidth	是否展示为紧凑型的图片。默认为false
	CompactWidth bool
	// Mode 图片显示模式。默认:居中裁剪模式
	Mode ImgModType
	// Preview	点击后是否放大图片，缺省为true
	Preview bool
}

// NewImgCardContent create ImgCardContent
func NewImgCardContent() *ImgCardContent {
	return &ImgCardContent{
		CompactWidth: false,
		Mode:         CropCenterMod,
		Preview:      true,
	}
}

// SetImgKey set ImgCardContent.ImgKey
func (i *ImgCardContent) SetImgKey(imgKey string) *ImgCardContent {
	i.ImgKey = imgKey
	return i
}

// SetAlt set ImgCardContent.Alt
func (i *ImgCardContent) SetAlt(alt *Text) *ImgCardContent {
	i.Alt = alt
	return i
}

// SetTitle set ImgCardContent.Title
func (i *ImgCardContent) SetTitle(title *Text) *ImgCardContent {
	i.Title = title
	return i
}

// SetCustomWidth set ImgCardContent.CustomWidth
func (i *ImgCardContent) SetCustomWidth(customWidth int) *ImgCardContent {
	i.CustomWidth = customWidth
	return i
}

// SetCompactWidth set ImgCardContent.CompactWidth
func (i *ImgCardContent) SetCompactWidth(compactWidth bool) *ImgCardContent {
	i.CompactWidth = compactWidth
	return i
}

// SetMode set ImgCardContent.Mode
func (i *ImgCardContent) SetMode(mode ImgModType) *ImgCardContent {
	i.Mode = mode
	return i
}

// SetPreview set ImgCardContent.Preview
func (i *ImgCardContent) SetPreview(preview bool) *ImgCardContent {
	i.Preview = preview
	return i
}
func (i *ImgCardContent) ToMessageMap() map[string]interface{} {
	message := map[string]interface{}{}
	message["tag"] = "img"
	message["img_key"] = i.ImgKey
	message["alt"] = i.Alt.ToMessageMap()
	if i.Title != nil {
		message["title"] = i.Title.ToMessageMap()
	}
	if i.CustomWidth != 0 {
		message["custom_width"] = i.CustomWidth
	}
	message["compact_width"] = i.CompactWidth
	message["mode"] = i.Mode
	message["preview"] = i.Preview
	return message
}

// NoteCardContent 备注模块
type NoteCardContent struct {
	// Elements 	备注信息
	Elements []Message
}

// NewNoteCardContent create NoteCardContent
func NewNoteCardContent() *NoteCardContent {
	return &NoteCardContent{
		Elements: []Message{},
	}
}

// SetElements  设置备注信息 ,只支持 Text 和 Image 元素
func (c *NoteCardContent) SetElements(elements []Message) *NoteCardContent {
	c.Elements = elements
	return c
}

// AddElement 添加备注信息 只支持 Text 和 Image 元素
func (c *NoteCardContent) AddElement(element Message) *NoteCardContent {
	c.Elements = append(c.Elements, element)
	return c
}

// AddElements 添加备注信息 只支持 Text 和 Image 元素
func (c *NoteCardContent) AddElements(elements ...Message) *NoteCardContent {
	c.Elements = append(c.Elements, elements...)
	return c
}

func (c *NoteCardContent) ToMessageMap() map[string]interface{} {
	message := map[string]interface{}{}
	message["tag"] = "note"
	var elementsMessage []map[string]interface{}
	for _, element := range c.Elements {
		elementsMessage = append(elementsMessage, element.ToMessageMap())
	}
	message["elements"] = elementsMessage
	return message
}
