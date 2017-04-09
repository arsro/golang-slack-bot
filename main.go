package main

import (
	"encoding/json"
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

func createSlackClient() *Slack {
	sp := &Slack{}
	return sp
}

func (sp *Slack) setPostParamter(bot_name string, icon_url string, channel string) {
	sp.Username = bot_name
	if icon_url == "" {
		sp.IconEmoji = ":japanese_goblin:"
	} else {
		sp.IconUrl = icon_url
	}
	sp.Channel = channel
}

func (sp *Slack) sendMessage(msg string, incomming_webhook_url string) (string, error) {
	sp.Text = msg

	params, err := json.Marshal(*sp)
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
