# feishu-robot
[![Go Reference](https://pkg.go.dev/badge/github.com/group-robot/feishu-robot.svg)](https://pkg.go.dev/github.com/group-robot/feishu-robot) 飞书机器人

# Example

## text
```go
webhok := os.Getenv("webhok")
secret := os.Getenv("secret")
client := newClient(webhok, secret)
_, err := client.Send(NewTextMessage().SetContent("test").SetAtAll())
```

## image
```go
webhok := os.Getenv("webhok")
secret := os.Getenv("secret")
client := newClient(webhok, secret)
_, err := client.Send(NewImageMessage().SetImageKey("img_7ea74629-9191-4176-998c-2e603c9c5e8g"))
```

## interactive
```go
	message := NewInteractiveMessage()
	message.SetConfig(NewCardConfig().SetWideScreenMode(true)).
		SetHeader(
			NewCardHeader().SetTemplate("indigo").SetTitle(newCardTitle().SetContent("🐈 英国短毛猫")),
		).SetElements(
		NewElementsContent().AddElement(
			NewDivCardContent().
				SetText(NewText().
					Tag(LarkMd).
					SetContent("英国短毛猫，体形圆胖，四肢短粗发达，毛短而密，头大脸圆，对人友善。 \n其历史可追溯至古罗马时期的家猫，由于拥有悠久的育种历史，称得上是猫家族中的典范。"),
				).SetExtra(NewImgCardContent().SetImgKey("img_1cad0e51-26f6-492a-8280-a47057b09a0g").SetAlt(NewText().Tag(PlainText).SetContent("图片"))),
		).AddElement(
			NewDivCardContent().AddField(
				NewFiled().SetShort(true).SetText(NewText().Tag(LarkMd).SetContent("**中文学名：**\n英国短毛猫")),
			).AddField(
				NewFiled().SetShort(true).SetText(NewText().Tag(LarkMd).SetContent("**拉丁学名：**\nFelinae")),
			).AddField(
				NewFiled().SetShort(false).SetText(NewText().Tag(LarkMd).SetContent("")),
			).AddField(
				NewFiled().SetShort(true).SetText(NewText().Tag(LarkMd).SetContent("**体形：**\n圆胖")),
			).AddField(
				NewFiled().SetShort(true).SetText(NewText().Tag(LarkMd).SetContent("**被毛：**\n短而浓密、俗称地毯毛")),
			),
		).AddElement(NewHrCardContent()).AddElement(
			NewDivCardContent().
				SetText(NewText().
					Tag(LarkMd).
					SetContent("**1 形态特征**\n\n 🔵 外形：身体厚实，胸部饱满宽阔，腿部粗壮，爪子浑圆，尾巴的根部粗壮，尾尖钝圆。\n\n🔵 毛色：共有十五种品种被承认，其中最著名的是蓝色系的英国短毛猫。 ")).
				SetExtra(NewImgCardContent().SetImgKey("img_70558e3a-2eef-4e8f-9a07-a701c165431g").SetAlt(NewText().Tag(PlainText).SetContent("图片"))),
		).AddElement(
			NewNoteCardContent().AddElement(
				NewImgCardContent().SetImgKey("img_7ea74629-9191-4176-998c-2e603c9c5e8g").SetAlt(NewText().Tag(PlainText).SetContent("图片")),
			).AddElement(
				NewText().Tag(PlainText).SetContent("以上资料来自百度百科"),
			),
		),
	)
	webhok := os.Getenv("webhok")
	secret := os.Getenv("secret")
	bt, err := json.Marshal(message.ToMessageMap())
	if err != nil {
		t.Fatalf("error:%s", err.Error())
	}
	t.Logf("json:%s", string(bt))

	client := newClient(webhok, secret)
	_, err = client.Send(message)
```

```go
	jsonStr := `{
        "config": {
            "enable_forward": true,
            "wide_screen_mode": true
        },
        "elements": [
            {
                "extra": {
                    "alt": {
                        "content": "图片",
                        "tag": "plain_text"
                    },
                    "compact_width": false,
                    "img_key": "img_1cad0e51-26f6-492a-8280-a47057b09a0g",
                    "mode": "crop_center",
                    "preview": true,
                    "tag": "img"
                },
                "fields": null,
                "tag": "div",
                "text": {
                    "content": "英国短毛猫，体形圆胖，四肢短粗发达，毛短而密，头大脸圆，对人友善。 \n其历史可追溯至古罗马时期的家猫，由于拥有悠久的育种历史，称得上是猫家族中的典范。",
                    "tag": "lark_md"
                }
            },
            {
                "fields": [
                    {
                        "is_short": true,
                        "text": {
                            "content": "**中文学名：**\n英国短毛猫",
                            "tag": "lark_md"
                        }
                    },
                    {
                        "is_short": true,
                        "text": {
                            "content": "**拉丁学名：**\nFelinae",
                            "tag": "lark_md"
                        }
                    },
                    {
                        "is_short": false,
                        "text": {
                            "content": "",
                            "tag": "lark_md"
                        }
                    },
                    {
                        "is_short": true,
                        "text": {
                            "content": "**体形：**\n圆胖",
                            "tag": "lark_md"
                        }
                    },
                    {
                        "is_short": true,
                        "text": {
                            "content": "**被毛：**\n短而浓密、俗称地毯毛",
                            "tag": "lark_md"
                        }
                    }
                ],
                "tag": "div"
            },
            {
                "tag": "hr"
            },
            {
                "extra": {
                    "alt": {
                        "content": "图片",
                        "tag": "plain_text"
                    },
                    "compact_width": false,
                    "img_key": "img_70558e3a-2eef-4e8f-9a07-a701c165431g",
                    "mode": "crop_center",
                    "preview": true,
                    "tag": "img"
                },
                "fields": null,
                "tag": "div",
                "text": {
                    "content": "**1 形态特征**\n\n 🔵 外形：身体厚实，胸部饱满宽阔，腿部粗壮，爪子浑圆，尾巴的根部粗壮，尾尖钝圆。\n\n🔵 毛色：共有十五种品种被承认，其中最著名的是蓝色系的英国短毛猫。 ",
                    "tag": "lark_md"
                }
            },
            {
                "elements": [
                    {
                        "alt": {
                            "content": "图片",
                            "tag": "plain_text"
                        },
                        "compact_width": false,
                        "img_key": "img_7ea74629-9191-4176-998c-2e603c9c5e8g",
                        "mode": "crop_center",
                        "preview": true,
                        "tag": "img"
                    },
                    {
                        "content": "以上资料来自百度百科",
                        "tag": "plain_text"
                    }
                ],
                "tag": "note"
            }
        ],
        "header": {
            "template": "indigo",
            "title": {
                "content": "🐈 英国短毛猫",
                "tag": "plain_text"
            }
        }
    }
`
	message := NewInteractiveMessage()
	message.SetCardJsonStr(jsonStr)
	webhok := os.Getenv("webhok")
	secret := os.Getenv("secret")

	client := newClient(webhok, secret)
```