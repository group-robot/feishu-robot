package feishu

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/hb0730/feishu-robot/internal/security"
	"strconv"
	"time"
)

// Client feishu robot client
type Client struct {
	WebHok string
	Secret string
}

// NewClient new feishu robot client
func NewClient(webhok, secret string) *Client {
	return &Client{webhok, secret}
}

// Response response struct
type Response struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}

// Send send message
func (c *Client) Send(message Message) (*Response, error) {
	res := &Response{}
	if len(c.WebHok) < 1 {
		return res, fmt.Errorf("webhok is blank")
	}
	timestamp := time.Now().Unix()
	sign, err := security.GenSign(c.Secret, timestamp)
	if err != nil {
		return res, err
	}

	body := message.ToMessageMap()
	body["timestamp"] = strconv.FormatInt(timestamp, 10)
	body["sign"] = sign
	client := resty.New()
	resp, err := client.SetRetryCount(3).R().
		SetBody(body).
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/json").
		SetResult(&Response{}).
		ForceContentType("application/json").
		Post(c.WebHok)
	if err != nil {
		return nil, err
	}
	result := resp.Result().(*Response)
	if result.Code != 0 {
		return res, fmt.Errorf("send message to feishu error = %s", result.Msg)
	}
	return result, nil

}
