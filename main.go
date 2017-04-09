package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
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

func main() {

	sp := createSlackClient()
	err := sp.setPostParamter(
		"test",
		":octocat:",
		"",
		"#memo",
	)
	if err != nil {
		log.Fatal(err)
	}

	response, err := sp.sendMessage("Hello World!", "https://hooks.slack.com/services/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")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(response)
}

func createSlackClient() *Slack {
	sp := &Slack{}
	return sp
}

func (sp *Slack) setPostParamter(bot_name string, icon_emoji string, icon_url string, channel string) error {
	sp.Username = bot_name
	sp.IconEmoji = icon_emoji
	sp.IconUrl = icon_url
	sp.Channel = channel

	if icon_url == "" && icon_emoji == "" {
		return errors.New("invalid params")
	}
	return nil
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
