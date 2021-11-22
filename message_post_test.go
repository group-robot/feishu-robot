package feishu

import (
	"encoding/json"
	"os"
	"testing"
)

func TestNewPostUnit(t *testing.T) {
	message := NewPostUnit().
		SetTitle("项目更新通知").
		AddPostTags(
			NewPostTags().addTag(
				NewTextTag().SetContent("项目有更新:"),
				NewATag().SetContent("请查看").SetHref("http://www.example.com/"),
				NewAtTag().SetAtAll(),
			),
		).ToMessageMap()
	bt, err := json.Marshal(message)
	if err != nil {
		t.Fatalf("error:%s", err.Error())
	}
	t.Logf("json:%s", string(bt))
}

func TestNewPostContent(t *testing.T) {
	message := NewPostContent().SetLang("zh_cn").SetPostUnit(
		NewPostUnit().
			SetTitle("项目更新通知").
			AddPostTags(
				NewPostTags().addTag(
					NewTextTag().SetContent("项目有更新:"),
					NewATag().SetContent("请查看").SetHref("http://www.example.com/"),
					NewAtTag().SetAtAll(),
				),
			),
	).ToMessageMap()
	bt, err := json.Marshal(message)
	if err != nil {
		t.Fatalf("error:%s", err.Error())
	}
	t.Logf("json:%s", string(bt))
}

func TestPostMessage_ToMessageMap(t *testing.T) {
	webhok := os.Getenv("webhok")
	secret := os.Getenv("secret")
	message := NewPostMessage().addContent(
		NewPostContent().SetLang("zh_cn").SetPostUnit(
			NewPostUnit().
				SetTitle("项目更新通知").
				AddPostTags(
					NewPostTags().addTag(
						NewTextTag().SetContent("项目有更新:"),
						NewATag().SetContent("请查看").SetHref("http://www.example.com/"),
						NewAtTag().SetAtAll(),
					),
				),
		),
	)
	bt, err := json.Marshal(message.ToMessageMap())
	if err != nil {
		t.Fatalf("error:%s", err.Error())
	}
	t.Logf("json:%s", string(bt))

	client := newClient(webhok, secret)
	_, err = client.Send(message)
	if err != nil {
		t.Logf("%s", err)
	}
}
