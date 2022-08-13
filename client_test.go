package feishu

import (
	"os"
	"testing"
)

func TestClient_SendMessageStr(t *testing.T) {
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
	if err != nil {
		t.Error("send message error", err)
	}
	if rep.IsSuccess() {
		t.Log("send message success")
	} else {
		t.Errorf("send message failed: code %d, msg: %s", rep.Code, rep.Msg)
	}
}
