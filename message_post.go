package feishu

type PostMessage struct {
	BaseMessage
	Content map[string]interface{}
}

func (message *PostMessage) addContent(content *PostContent) *PostMessage {
	for k, v := range content.ToMessageMap() {
		message.Content[k] = v
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

type PostContent struct {
	lang     string
	PostUnit *PostUnit
}

func (postContent *PostContent) ToMessageMap() map[string]interface{} {
	langMap := map[string]interface{}{}
	lang := "zh-cn"
	if postContent.lang != "" {
		lang = postContent.lang
	}
	langMap[lang] = postContent.PostUnit.ToMessageMap()
	return langMap
}
func NewPostContent() *PostContent {
	return new(PostContent)
}
func (postContent *PostContent) SetLang(lang string) *PostContent {
	postContent.lang = lang
	return postContent
}
func (postContent *PostContent) SetPostUnit(unit *PostUnit) *PostContent {
	postContent.PostUnit = unit
	return postContent
}

type PostUnit struct {
	Title   string
	Content [][]map[string]interface{}
}

func (p *PostUnit) ToMessageMap() map[string]interface{} {
	postUnitMap := map[string]interface{}{}
	postUnitMap["title"] = p.Title
	postUnitMap["content"] = p.Content
	return postUnitMap
}

func NewPostUnit() *PostUnit {
	return &PostUnit{}
}

func (p *PostUnit) SetTitle(title string) *PostUnit {
	p.Title = title
	return p
}
func (p *PostUnit) SetContent(content [][]map[string]interface{}) *PostUnit {
	p.Content = content
	return p
}
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
func (p *PostUnit) AddPostTag(postTags *PostTags) *PostUnit {
	var tag []map[string]interface{}
	for _, v := range postTags.postTags {
		tag = append([]map[string]interface{}{}, v.ToMessageMap())
	}
	p.Content = append(p.Content, tag)
	return p
}
func (p *PostUnit) AddPostTags(tags ...*PostTags) *PostUnit {
	var tag []map[string]interface{}
	for _, postTags := range tags {
		for _, v := range postTags.postTags {
			tag = append(tag, v.ToMessageMap())
		}
	}
	p.Content = append(p.Content, tag)
	return p
}

type PostTag interface {
	Message
	getTag() string
}
type TextTag struct {
	Content  string
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

func NewTextTag() *TextTag {
	return new(TextTag)
}
func (tag *TextTag) SetContent(content string) *TextTag {
	tag.Content = content
	return tag
}
func (tag *TextTag) SetUnEsCape() *TextTag {
	tag.UnEsCape = true
	return tag
}

type ATag struct {
	Href    string
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
func NewATag() *ATag {
	return new(ATag)
}
func (tag *ATag) SetHref(href string) *ATag {
	tag.Href = href
	return tag
}
func (tag *ATag) SetContent(content string) *ATag {
	tag.Content = content
	return tag
}

type AtTag struct {
	UserId   string
	AtAll    bool
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
func NewAtTag() *AtTag {
	return new(AtTag)
}
func (tag *AtTag) SetUserId(userId string) *AtTag {
	tag.UserId = userId
	return tag
}
func (tag *AtTag) SetUserName(username string) *AtTag {
	tag.Username = username
	return tag
}
func (tag *AtTag) SetAtAll() *AtTag {
	tag.AtAll = true
	return tag
}

type ImgTag struct {
	ImageKey string
	Height   int
	Width    int
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

func NewImgTag() *ImgTag {
	return new(ImgTag)
}
func (tag *ImgTag) SetImageKey(imageKey string) *ImgTag {
	tag.ImageKey = imageKey
	return tag
}

func (tag *ImgTag) SetHeight(height int) *ImgTag {
	tag.Height = height
	return tag
}
func (tag *ImgTag) SetWidth(width int) *ImgTag {
	tag.Width = width
	return tag
}

type PostTags struct {
	postTags []PostTag
}

func NewPostTags() *PostTags {
	tags := &PostTags{
		postTags: []PostTag{},
	}
	return tags
}

func (tags *PostTags) addTag(tag ...PostTag) *PostTags {
	tags.postTags = append(tags.postTags, tag...)
	return tags
}
