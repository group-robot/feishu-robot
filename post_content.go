package feishu

type DivCardContent struct {
	Text   *Text
	Fields []*Filed
	Extra  Message
}

func NewDivCardContent() *DivCardContent {
	return &DivCardContent{
		Fields: []*Filed{},
	}
}
func (c *DivCardContent) SetText(text *Text) *DivCardContent {
	c.Text = text
	return c
}
func (c *DivCardContent) SetFields(fields []*Filed) *DivCardContent {
	c.Fields = fields
	return c
}
func (c *DivCardContent) AddField(filed *Filed) *DivCardContent {
	c.Fields = append(c.Fields, filed)
	return c
}
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

type MarkdownCardContent struct {
	Content string
	Href    *Url
}

func NewMarkdownCardContent() *MarkdownCardContent {
	return new(MarkdownCardContent)
}

func (c *MarkdownCardContent) SetContent(content string) *MarkdownCardContent {
	c.Content = content
	return c
}

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

type HrCardContent struct {
}

func NewHrCardContent() *HrCardContent {
	return new(HrCardContent)
}
func (c *HrCardContent) ToMessageMap() map[string]interface{} {
	message := map[string]interface{}{}
	message["tag"] = "hr"
	return message
}

type ImgModType string

const (
	FitHorizontalMod ImgModType = "fit_horizontal"
	CropCenterMod    ImgModType = "crop_center"
)

type ImgCardContent struct {
	ImgKey       string
	Alt          *Text
	Title        *Text
	CustomWidth  int
	CompactWidth bool
	Mode         ImgModType
	Preview      bool
}

func NewImgCardContent() *ImgCardContent {
	return &ImgCardContent{
		CompactWidth: false,
		Mode:         CropCenterMod,
		Preview:      true,
	}
}

func (i *ImgCardContent) SetImgKey(imgKey string) *ImgCardContent {
	i.ImgKey = imgKey
	return i
}
func (i *ImgCardContent) SetAlt(alt *Text) *ImgCardContent {
	i.Alt = alt
	return i
}

func (i *ImgCardContent) SetTitle(title *Text) *ImgCardContent {
	i.Title = title
	return i
}

func (i *ImgCardContent) SetCustomWidth(customWidth int) *ImgCardContent {
	i.CustomWidth = customWidth
	return i
}
func (i *ImgCardContent) SetCompactWidth(compactWidth bool) *ImgCardContent {
	i.CompactWidth = compactWidth
	return i
}
func (i *ImgCardContent) SetMode(mode ImgModType) *ImgCardContent {
	i.Mode = mode
	return i
}
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

type NoteCardContent struct {
	Elements []Message
}

func NewNoteCardContent() *NoteCardContent {
	return &NoteCardContent{
		Elements: []Message{},
	}
}
func (c *NoteCardContent) SetElements(elements []Message) *NoteCardContent {
	c.Elements = elements
	return c
}
func (c *NoteCardContent) AddElement(element Message) *NoteCardContent {
	c.Elements = append(c.Elements, element)
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
