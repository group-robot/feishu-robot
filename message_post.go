package feishu

// PostMessage 富文本消息
type PostMessage struct {
	BaseMessage
	// Content 内容
	Content map[string]interface{}
}

// AddContent add PostMessage.Content
func (message *PostMessage) AddContent(content *PostContent) *PostMessage {
	for k, v := range content.ToMessageMap() {
		message.Content[k] = v
	}
	return message
}

// AddContents add PostMessage.Content
func (message *PostMessage) AddContents(contents ...*PostContent) *PostMessage {
	for _, content := range contents {
		message.AddContent(content)
	}
	return message
}
func (message *PostMessage) ToMessageMap() map[string]interface{} {
	contentMessage := map[string]interface{}{}
	contentMessage["post"] = message.Content
	postMessage := map[string]interface{}{}
	postMessage["msg_type"] = message.GetMsgType()
	postMessage["content"] = contentMessage
	return postMessage
}
func (message *PostMessage) GetMsgType() MsgType {
	return MsgTypePost
}
func NewPostMessage() *PostMessage {
	return &PostMessage{
		Content: map[string]interface{}{},
	}
}

// PostContent 富文本内容
type PostContent struct {
	// Lang 语言 默认:zh-cn
	Lang string
	// PostUnit 富文本具体内容
	PostUnit *PostUnit
}

func (postContent *PostContent) ToMessageMap() map[string]interface{} {
	langMap := map[string]interface{}{}
	lang := "zh-cn"
	if postContent.Lang != "" {
		lang = postContent.Lang
	}
	langMap[lang] = postContent.PostUnit.ToMessageMap()
	return langMap
}

// NewPostContent create PostContent
func NewPostContent() *PostContent {
	return new(PostContent)
}

// SetLang set PostContent.Lang
func (postContent *PostContent) SetLang(lang string) *PostContent {
	postContent.Lang = lang
	return postContent
}

// SetPostUnit set 富文本每组段落
func (postContent *PostContent) SetPostUnit(unit *PostUnit) *PostContent {
	postContent.PostUnit = unit
	return postContent
}

// PostUnit 富文本每组内容
type PostUnit struct {
	// Title 标题
	Title string
	// Content 段落
	Content [][]map[string]interface{}
}

func (p *PostUnit) ToMessageMap() map[string]interface{} {
	postUnitMap := map[string]interface{}{}
	postUnitMap["title"] = p.Title
	postUnitMap["content"] = p.Content
	return postUnitMap
}

// NewPostUnit create PostUnit
func NewPostUnit() *PostUnit {
	return &PostUnit{
		Content: [][]map[string]interface{}{},
	}
}

// SetTitle set 标题
func (p *PostUnit) SetTitle(title string) *PostUnit {
	p.Title = title
	return p
}

// SetContent set 富文本段落
func (p *PostUnit) SetContent(content [][]map[string]interface{}) *PostUnit {
	p.Content = content
	return p
}

// AddTags add PostTags
func (p *PostUnit) AddTags(tags []PostTag) *PostUnit {
	if p.Content == nil {
		p.Content = make([][]map[string]interface{}, 0)
	}
	var tag []map[string]interface{}
	for _, v := range tags {
		tag = append([]map[string]interface{}{}, v.ToMessageMap())
	}
	p.Content = append(p.Content, tag)
	return p
}

// AddPostTag add PostTags
func (p *PostUnit) AddPostTag(postTags *PostTags) *PostUnit {
	var tag []map[string]interface{}
	for _, v := range postTags.PostTags {
		tag = append([]map[string]interface{}{}, v.ToMessageMap())
	}
	p.Content = append(p.Content, tag)
	return p
}

// AddPostTags add PostTags
func (p *PostUnit) AddPostTags(tags ...*PostTags) *PostUnit {
	var tag []map[string]interface{}
	for _, postTags := range tags {
		for _, v := range postTags.PostTags {
			tag = append(tag, v.ToMessageMap())
		}
	}
	p.Content = append(p.Content, tag)
	return p
}

// PostTag 富文本标签
type PostTag interface {
	Message
	// getTag 标签名
	getTag() string
}

// TextTag 文本标签
type TextTag struct {
	// Content 	文本内容
	Content string
	// UnEsCape 表示是不是,默认false
	UnEsCape bool
}

func (tag *TextTag) getTag() string {
	return "text"
}

func (tag *TextTag) ToMessageMap() map[string]interface{} {
	contentMap := map[string]interface{}{}
	contentMap["tag"] = tag.getTag()
	contentMap["text"] = tag.Content
	if tag.UnEsCape {
		contentMap["un_escape"] = tag.UnEsCape
	}
	return contentMap
}

// NewTextTag create TextTag
func NewTextTag() *TextTag {
	return new(TextTag)
}

//SetContent set TextTag.Content
func (tag *TextTag) SetContent(content string) *TextTag {
	tag.Content = content
	return tag
}

//SetUnEsCape set TextTag.UnEsCape is true
func (tag *TextTag) SetUnEsCape() *TextTag {
	tag.UnEsCape = true
	return tag
}

// ATag a标签
type ATag struct {
	// Href 默认的链接地址
	Href string
	// Content 文本内容
	Content string
}

func (tag *ATag) getTag() string {
	return "a"
}
func (tag *ATag) ToMessageMap() map[string]interface{} {
	tagMessage := map[string]interface{}{}
	tagMessage["tag"] = tag.getTag()
	tagMessage["href"] = tag.Href
	tagMessage["text"] = tag.Content
	return tagMessage
}

// NewATag create ATag
func NewATag() *ATag {
	return new(ATag)
}

// SetHref set ATag.Href
func (tag *ATag) SetHref(href string) *ATag {
	tag.Href = href
	return tag
}

// SetContent set ATag.Content
func (tag *ATag) SetContent(content string) *ATag {
	tag.Content = content
	return tag
}

// AtTag At标签
type AtTag struct {
	// UserId open_id
	UserId string
	// AtAll 是否at全部，与UserId互斥
	AtAll bool
	// Username 	用户姓名
	Username string
}

func (tag *AtTag) getTag() string {
	return "at"
}

func (tag *AtTag) ToMessageMap() map[string]interface{} {
	tagMap := map[string]interface{}{}
	tagMap["tag"] = tag.getTag()
	if tag.AtAll {
		tagMap["user_id"] = "all"
		tagMap["user_name"] = "所有人"
	} else {
		tagMap["user_id"] = tag.UserId
		if tag.Username != "" {
			tagMap["user_name"] = tag.Username
		}
	}
	return tagMap
}

// NewAtTag create AtTag
func NewAtTag() *AtTag {
	return new(AtTag)
}

// SetUserId set AtTag.UserId
func (tag *AtTag) SetUserId(userId string) *AtTag {
	tag.UserId = userId
	return tag
}

// SetUserName set AtTag.Username
func (tag *AtTag) SetUserName(username string) *AtTag {
	tag.Username = username
	return tag
}

// SetAtAll set AtTag.AtAll
func (tag *AtTag) SetAtAll(atAll bool) *AtTag {
	tag.AtAll = atAll
	return tag
}

// IsAtAll set AtTag.AtAll is true
func (tag *AtTag) IsAtAll() *AtTag {
	tag.AtAll = true
	return tag
}

// ImgTag img标签
type ImgTag struct {
	// ImageKey 图片的唯一标识
	ImageKey string
	// Height 图片的高
	Height int
	// Width 图片的宽
	Width int
}

func (tag *ImgTag) getTag() string {
	return "img"
}
func (tag *ImgTag) ToMessageMap() map[string]interface{} {
	tagMap := map[string]interface{}{}
	tagMap["tag"] = tag.getTag()
	tagMap["image_key"] = tag.ImageKey
	if tag.Width != 0 {
		tagMap["width"] = tag.Width
	}
	if tag.Height != 0 {
		tagMap["height"] = tag.Height
	}
	return tagMap
}

// NewImgTag create ImgTag
func NewImgTag() *ImgTag {
	return new(ImgTag)
}

// SetImageKey set ImgTag.ImageKey
func (tag *ImgTag) SetImageKey(imageKey string) *ImgTag {
	tag.ImageKey = imageKey
	return tag
}

// SetHeight set ImgTag.Height
func (tag *ImgTag) SetHeight(height int) *ImgTag {
	tag.Height = height
	return tag
}

// SetWidth set ImgTag.Width
func (tag *ImgTag) SetWidth(width int) *ImgTag {
	tag.Width = width
	return tag
}

// PostTags 富文本标签
type PostTags struct {
	// PostTags 富文本标签集
	PostTags []PostTag
}

// NewPostTags create PostTags
func NewPostTags() *PostTags {
	tags := &PostTags{
		PostTags: []PostTag{},
	}
	return tags
}

// AddTags add post tags
func (tag *PostTags) AddTags(tags ...PostTag) *PostTags {
	tag.PostTags = append(tag.PostTags, tags...)
	return tag
}

// AddTag add post tag
func (tag *PostTags) AddTag(postTag PostTag) *PostTags {
	tag.PostTags = append(tag.PostTags, postTag)
	return tag
}
