package feishu

import (
	"os"
	"testing"
)

func TestNewImageMessage(t *testing.T) {
	webhok := os.Getenv("webhok")
	secret := os.Getenv("secret")
	client := newClient(webhok, secret)
	_, err := client.Send(NewImageMessage().SetImageKey("img_7ea74629-9191-4176-998c-2e603c9c5e8g"))
	if err != nil {
		t.Logf("%s", err)
	}
}
