# feishu-robot
[![Go Reference](https://pkg.go.dev/badge/github.com/group-robot/feishu-robot.svg)](https://pkg.go.dev/github.com/group-robot/feishu-robot/v2) 飞书机器人

# Example

## text
```go
text := NewTextMessage("新更新提醒", false)
client := NewClient()
client.Webhook = os.Getenv("webhook")
client.Secret = os.Getenv("secret")
rep, err := client.SendMessage(text)
```

## image
```go
message := NewImageMessage("img_ecffc3b9-8f14-400f-a014-05eca1a4310g")
client := NewClient()
client.Webhook = os.Getenv("webhook")
client.Secret = os.Getenv("secret")
rep, err := client.SendMessage(message)
```

### post
```go
	message := NewPostMessage(
		NewZhCnLangPostItem(
			NewPostItems(
				"项目更新通知",
				NewPostTags(
					NewTextTag("项目有更新: "),
				),
			).AddContent(
				NewPostTags(
					NewATag("请查看", "http://www.example.com/"),
					NewAtAllAtTag(),
				),
			),
		),
	)
	client := NewClient()
	client.Webhook = os.Getenv("webhook")
	client.Secret = os.Getenv("secret")
	rep, err := client.SendMessage(message)
```

## interactive
```go
	message := NewInteractiveMessage()
message.SetConfig(
NewCardConfig().
SetWideScreenMode(true).
SetEnableForward(true),
).AddElements(
NewCardElement(
NewCardText(Md, "**西湖**，位于浙江省杭州市西湖区龙井路1号，杭州市区西部，景区总面积49平方千米，汇水面积为21.22平方千米，湖面面积为6.38平方千米。"),
),
NewCardAction(
NewButtonActionElement(
NewCardText(Md, "更多景点介绍 :玫瑰:"),
).SetUrl("https://www.example.com").SetType(DefaultType),
),
).SetHeader(
NewCardHeader(NewCardTitle("今日旅游推荐", nil)),
)
client := NewClient()
client.Webhook = os.Getenv("webhook")
client.Secret = os.Getenv("secret")
rep, err := client.SendMessage(message)
```

```go
	json := `{
    "msg_type": "interactive",
    "card": {
        "config": {
                "wide_screen_mode": true,
                "enable_forward": true
        },
        "elements": [{
                "tag": "div",
                "text": {
                        "content": "**西湖**，位于浙江省杭州市西湖区龙井路1号，杭州市区西部，景区总面积49平方千米，汇水面积为21.22平方千米，湖面面积为6.38平方千米。",
                        "tag": "lark_md"
                }
        }, {
                "actions": [{
                        "tag": "button",
                        "text": {
                                "content": "更多景点介绍 :玫瑰:",
                                "tag": "lark_md"
                        },
                        "url": "https://www.example.com",
                        "type": "default",
                        "value": {}
                }],
                "tag": "action"
        }],
        "header": {
                "title": {
                        "content": "今日旅游推荐",
                        "tag": "plain_text"
                }
        }
    }
} `
client := NewClient()
client.Webhook = os.Getenv("webhook")
client.Secret = os.Getenv("secret")
rep, err := client.SendMessageStr(json)
```