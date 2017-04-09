package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"golang_slack_bot/config"
)

const ()

type Slack struct {
	Text      string `json:"text"`
	Username  string `json:"username"`
	IconEmoji string `json:"icon_emoji"`
	IconUrl   string `json:"icon_url"`
	Channel   string `json:"channel"`
}

func main() {

	sp := createSlackClient()
	sp.setPostParamter(
		"test",
		"https://applets.imgix.net/https%3A%2F%2Fassets.ifttt.com%2Fimages%2Fchannels%2F2107379463%2Ficons%2Fon_color_large.png%3Fversion%3D0?ixlib=rails-2.1.3&w=240&h=240&auto=compress&s=07c1117d9e046c1a26150728429d62db",
		"#memo",
	)
	response, err := sp.sendMessage("Hello World", "")
	if err != nil {
		log.Println(err)
	}
	log.Println(response)
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

	// config.IncomingUrlをslackのincomming webhookのurlに設定
	if incomming_webhook_url == "" {
		incomming_webhook_url = config.IncomingUrl
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
