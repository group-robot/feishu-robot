package feishu

import (
	"os"
	"testing"
)

func TestNewTextMessage(t *testing.T) {
	webhok := os.Getenv("webhok")
	secret := os.Getenv("secret")
	client := NewClient(webhok, secret)
	_, err := client.Send(NewTextMessage().SetContent("test").SetAtAll())
	if err != nil {
		t.Logf("%s", err)
	}
}

func TestTextMessage_NonAtAll(t *testing.T) {
	webhok := os.Getenv("webhok")
	secret := os.Getenv("secret")
	client := NewClient(webhok, secret)
	_, err := client.Send(NewTextMessage().SetContent("test"))
	if err != nil {
		t.Logf("%s", err)
	}
}
