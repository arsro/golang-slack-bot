package main

import (
	"github.com/tomoryes/golang-slack-bot/command"
	"log"
)

func main() {

	s := command.CreateSlackClient()
	err := s.SetPostParamter(
		"hogehoge",
		":octocat:",
		"",
		"#memo",
	)
	if err != nil {
		log.Fatal(err)
	}

	response, err := s.SendMessage("Hello World!", "https://hooks.slack.com/services/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")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(response)
}
