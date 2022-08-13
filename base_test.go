package feishu

import (
	"os"
	"testing"
)

func TestTextMessage_ToMessageMap(t *testing.T) {
	text := NewTextMessage("新更新提醒", false)
	client := NewClient()
	client.Webhook = os.Getenv("webhook")
	client.Secret = os.Getenv("secret")
	rep, err := client.SendMessage(text)
	if err != nil {
		t.Error("send message error", err)
	}
	if rep.IsSuccess() {
		t.Log("send message success")
	} else {
		t.Errorf("send message failed: code %d, msg: %s", rep.Code, rep.Msg)
	}
}

func TestTextMessage_ToMessageMap_AtAll(t *testing.T) {
	text := NewTextMessage("新更新提醒", true)
	client := NewClient()
	client.Webhook = os.Getenv("webhook")
	client.Secret = os.Getenv("secret")
	rep, err := client.SendMessage(text)
	if err != nil {
		t.Error("send message error", err)
	}
	if rep.IsSuccess() {
		t.Log("send message success")
	} else {
		t.Errorf("send message failed: code %d, msg: %s", rep.Code, rep.Msg)
	}
}

func TestPostMessage_ToMessageMap(t *testing.T) {
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
	if err != nil {
		t.Error("send message error", err)
	}
	if rep.IsSuccess() {
		t.Log("send message success")
	} else {
		t.Errorf("send message failed: code %d, msg: %s", rep.Code, rep.Msg)
	}
}

func TestImageMessage_ToMessageMap(t *testing.T) {
	message := NewImageMessage("img_ecffc3b9-8f14-400f-a014-05eca1a4310g")
	client := NewClient()
	client.Webhook = os.Getenv("webhook")
	client.Secret = os.Getenv("secret")
	rep, err := client.SendMessage(message)
	if err != nil {
		t.Error("send message error", err)
	}
	if rep.IsSuccess() {
		t.Log("send message success")
	} else {
		t.Errorf("send message failed: code %d, msg: %s", rep.Code, rep.Msg)
	}
}

func TestInteractiveMessage_ToMessageMap(t *testing.T) {
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
	if err != nil {
		t.Error("send message error", err)
	}
	if rep.IsSuccess() {
		t.Log("send message success")
	} else {
		t.Errorf("send message failed: code %d, msg: %s", rep.Code, rep.Msg)
	}
}
func TestInteractiveMessage_ToMessageMap1(t *testing.T) {
	// 醒目的通知
	message := NewInteractiveMessage()
	message.SetConfig(
		NewCardConfig().SetWideScreenMode(true),
	).SetHeader(
		NewCardHeader(NewCardTitle("你有一个休假申请待审批", nil)).SetTemplate(Indigo),
	).AddElements(
		NewCardElement(nil).
			AddFields(
				NewCardField(true, NewCardText(Md, "**申请人**\n王晓磊")),
				NewCardField(true, NewCardText(Md, "**休假类型：**\n年假")),
				NewCardField(false, NewCardText(Md, "")),
				NewCardField(false, NewCardText(Md, "**时间：**\n2020-4-8 至 2020-4-10（共3天）")),
				NewCardField(false, NewCardText(Md, "")),
				NewCardField(true, NewCardText(Md, "**备注**\n因家中有急事，需往返老家，故请假")),
			),
		NewCardHr(),
		NewCardAction(
			NewButtonActionElement(NewCardText(Text, "批准")).
				SetType(PrimaryType).SetValue(map[string]string{"chosen": "approve"}),
			NewButtonActionElement(NewCardText(Text, "拒绝")).
				SetType(PrimaryType).SetValue(map[string]string{"chosen": "decline"}),
		).SetLayout(
			Bisected,
		),
	)
	client := NewClient()
	client.Webhook = os.Getenv("webhook")
	client.Secret = os.Getenv("secret")
	rep, err := client.SendMessage(message)
	if err != nil {
		t.Error("send message error", err)
	}
	if rep.IsSuccess() {
		t.Log("send message success")
	} else {
		t.Errorf("send message failed: code %d, msg: %s", rep.Code, rep.Msg)
	}
}
func TestInteractiveMessage_ToMessageMap2(t *testing.T) {
	// 醒目的通知
	message := NewInteractiveMessage()
	message.SetConfig(
		NewCardConfig().SetWideScreenMode(true),
	).SetHeader(
		NewCardHeader(NewCardTitle("🥤 下午的奶茶发车了，你要上车么", nil)).SetTemplate(Indigo),
	).AddElements(
		NewCardHr(),
		NewCardAction(
			NewButtonActionElement(NewCardText(Text, "😍 带我！带我！！")).
				SetType(DefaultType).SetValue(map[string]string{"chosen": "option1"}),
			NewButtonActionElement(NewCardText(Text, "🤐 告辞…")).
				SetType(DefaultType).SetValue(map[string]string{"chosen": "option2"}),
		),
		NewCardNote(
			NewCardText(Text, "创建者：王大明 🔐本投票为匿名投票"),
		),
	)
	client := NewClient()
	client.Webhook = os.Getenv("webhook")
	client.Secret = os.Getenv("secret")
	rep, err := client.SendMessage(message)
	if err != nil {
		t.Error("send message error", err)
	}
	if rep.IsSuccess() {
		t.Log("send message success")
	} else {
		t.Errorf("send message failed: code %d, msg: %s", rep.Code, rep.Msg)
	}
}
func TestInteractiveMessage_ToMessageMap3(t *testing.T) {
	// 醒目的通知
	message := NewInteractiveMessage()
	message.SetConfig(
		NewCardConfig().SetWideScreenMode(true),
	).SetHeader(
		NewCardHeader(NewCardTitle("🐈 英国短毛猫", nil)).SetTemplate(Indigo),
	).AddElements(
		NewCardElement(
			NewCardText(Md, "英国短毛猫，体形圆胖，四肢短粗发达，毛短而密，头大脸圆，对人友善。 \n其历史可追溯至古罗马时期的家猫，由于拥有悠久的育种历史，称得上是猫家族中的典范。"),
		).SetExtra(
			NewCardImage("img_1cad0e51-26f6-492a-8280-a47057b09a0g", NewCardText(Text, "图片")),
		),
		NewCardElement(
			nil,
		).AddFields(
			NewCardField(true, NewCardText(Md, "**中文学名：**\n英国短毛猫")),
			NewCardField(true, NewCardText(Md, "**拉丁学名：**\nFelinae")),
			NewCardField(false, NewCardText(Text, "")),
			NewCardField(true, NewCardText(Md, "**体形：**\n圆胖")),
			NewCardField(true, NewCardText(Md, "**被毛：**\n短而浓密、俗称地毯毛")),
		),
		NewCardHr(),
		NewCardElement(
			NewCardText(Md, "**1 形态特征**\n\n 🔵 外形：身体厚实，胸部饱满宽阔，腿部粗壮，爪子浑圆，尾巴的根部粗壮，尾尖钝圆。\n\n🔵 毛色：共有十五种品种被承认，其中最著名的是蓝色系的英国短毛猫。 "),
		).SetExtra(
			NewCardImage("img_70558e3a-2eef-4e8f-9a07-a701c165431g", NewCardText(Text, "图片")),
		),
		NewCardNote(
			NewCardImage("img_e61db329-2469-4da7-8f13-2d2f284c3b1g", NewCardText(Text, "图片")),
			NewCardText(Text, "以上资料来自头条百科"),
		),
	)
	client := NewClient()
	client.Webhook = os.Getenv("webhook")
	client.Secret = os.Getenv("secret")
	rep, err := client.SendMessage(message)
	if err != nil {
		t.Error("send message error", err)
	}
	if rep.IsSuccess() {
		t.Log("send message success")
	} else {
		t.Errorf("send message failed: code %d, msg: %s", rep.Code, rep.Msg)
	}
}
func TestInteractiveMessage_ToMessageMap4(t *testing.T) {
	// 醒目的通知
	message := NewInteractiveMessage()
	message.SetConfig(
		NewCardConfig().SetWideScreenMode(true),
	).SetHeader(
		NewCardHeader(NewCardTitle("🎉第一天，做点什么好呢？", nil)).SetTemplate(Wathet),
	).AddElements(
		NewCardImage("img_770020cd-f92e-4a1f-ac5f-20047cf1731g", NewCardText(Text, "")),
		NewCardElement(NewCardText(Md, "**换个头像吧！这6件事助你开启Lark第一天**  [点击查看>>](https://www.larksuite.com/hc/en-US/articles/360042500034)")),
		NewCardHr(),
		NewCardElement(NewCardText(Md, "**【聊天】消息串、表情回复？高效沟通很轻松** \n[点击查看>>](https://www.larksuite.com/hc/en-US/articles/360023545914)")).SetExtra(
			NewCardImage("img_50b4fd3f-4077-4e4b-b92c-664039a6153g", NewCardText(Text, "")),
		),
		NewCardHr(),
		NewCardElement(NewCardText(Md, "**【会议】多人异地办公难？音视频会议来帮你** \n[点击查看>>](https://www.larksuite.com/hc/en-US/articles/360035593213)")).
			SetExtra(
				NewCardImage("img_54b18a36-81b0-4776-9f47-854f9da6cd3g", NewCardText(Text, "")),
			),
		NewCardHr(),
		NewCardNote(
			NewCardImage("img_5aea87b0-63bc-4a52-9ca4-ce2e58468e5g", NewCardText(Text, "")),
			NewCardText(Text, "回复“退订”不再接收此订阅信息"),
		),
	)
	client := NewClient()
	client.Webhook = os.Getenv("webhook")
	client.Secret = os.Getenv("secret")
	rep, err := client.SendMessage(message)
	if err != nil {
		t.Error("send message error", err)
	}
	if rep.IsSuccess() {
		t.Log("send message success")
	} else {
		t.Errorf("send message failed: code %d, msg: %s", rep.Code, rep.Msg)
	}
}

func TestInteractiveMessage_ToMessageMap5(t *testing.T) {
	// 醒目的通知
	message := NewInteractiveMessage()
	message.SetConfig(
		NewCardConfig().SetWideScreenMode(true),
	).SetHeader(
		NewCardHeader(NewCardTitle("📚晒挚爱好书，赢读书礼金", nil)).SetTemplate(Turquoise),
	).AddElements(
		NewCardImage("img_7ea74629-9191-4176-998c-2e603c9c5e8g", NewCardText(Text, "图片")),
		NewCardElement(NewCardText(Md, "你是否曾因为一本书而产生心灵共振，开始感悟人生？\n你有哪些想极力推荐给他人的珍藏书单？\n\n加入 **4·23 飞书读书节**，分享你的**挚爱书单**及**读书笔记**，**赢取千元读书礼**！\n\n📬 填写问卷，晒出你的珍藏好书\n😍 想知道其他人都推荐了哪些好书？马上[入群围观](https://feishu.cn)\n📝 用[读书笔记模板](https://feishu.cn)（桌面端打开），记录你的心得体会\n🙌 更有惊喜特邀嘉宾 4月12日起带你共读")),
		NewCardAction(
			NewButtonActionElement(NewCardText(Text, "立即推荐好书")).SetType(PrimaryType).SetUrl("https://feishu.cn"),
			NewButtonActionElement(NewCardText(Text, "查看活动指南")).SetType(DefaultType).SetUrl("https://feishu.cn"),
		),
	)
	client := NewClient()
	client.Webhook = os.Getenv("webhook")
	client.Secret = os.Getenv("secret")
	rep, err := client.SendMessage(message)
	if err != nil {
		t.Error("send message error", err)
	}
	if rep.IsSuccess() {
		t.Log("send message success")
	} else {
		t.Errorf("send message failed: code %d, msg: %s", rep.Code, rep.Msg)
	}
}
