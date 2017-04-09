package command

import (
	"encoding/json"
	"errors"
	"io/ioutil"

	"net/http"
	"net/url"
)

type Slack struct {
	Text      string `json:"text"`
	Username  string `json:"username"`
	IconEmoji string `json:"icon_emoji"`
	IconUrl   string `json:"icon_url"`
	Channel   string `json:"channel"`
}

func CreateSlackClient() *Slack {
	return &Slack{}
}

func (s *Slack) SetPostParamter(bot_name string, icon_emoji string, icon_url string, channel string) error {
	s.Username = bot_name
	s.IconEmoji = icon_emoji
	s.IconUrl = icon_url
	s.Channel = channel

	if icon_url == "" && icon_emoji == "" {
		return errors.New("invalid params")
	}
	return nil
}

func (s *Slack) SendMessage(msg string, incomming_webhook_url string) (string, error) {
	s.Text = msg

	params, err := json.Marshal(*s)
	if err != nil {
		return "", err
	}

	payload := string(string(params))
	post_value := url.Values{"payload": {payload}}

	// post
	resp, _ := http.PostForm(
		incomming_webhook_url,
		post_value,
	)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	return string(body), nil
}
