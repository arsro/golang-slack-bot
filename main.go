package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"golang_slack_bot/config"
)

type Slack struct {
	Text       string `json:"text"`       //投稿内容
	Username   string `json:"username"`   //投稿者名 or Bot名（存在しなくてOK）
	Icon_emoji string `json:"icon_emoji"` //アイコン絵文字
	Icon_url   string `json:"icon_url"`   //アイコンURL（icon_emojiが存在する場合は、適応されない）
	Channel    string `json:"channel"`    //#部屋名
}

func main() {
	slack := &Slack{
		"Hello World",
		"slack bot",
		"",
		"http://www.ensky.co.jp/item/images/save/07151842_53c4f7839cd4e.jpg",
		"#memo",
	}

	params, _ := json.Marshal(*slack)
	resp, _ := http.PostForm(
		config.IncomingUrl,
		url.Values{"payload": {string(params)}},
	)

	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	println(string(body))
}
