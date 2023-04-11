package main

import (
	"fmt"
	"os"

	"github.com/seven-io/go-client/sms77api"
)

func main() {
	clientOptions := sms77api.Options{
		ApiKey: os.Getenv("PLUGIN_API_KEY"),
	}
	seven := sms77api.New(clientOptions)
	from := os.Getenv("PLUGIN_FROM")
	text := os.Getenv("PLUGIN_TEXT")
	to := os.Getenv("PLUGIN_TO")
	typ := os.Getenv("PLUGIN_TYPE")
	debug := parseBool("DEBUG")

	var err error

	if typ == "voice" {
		params := sms77api.VoiceParams{
			Debug: debug,
			From:  from,
			Text:  text,
			To:    to,
		}
		resp, _err := seven.Voice.Json(params)
		err = _err
		fmt.Println("Voice sent:", resp)
	} else {
		params := sms77api.SmsBaseParams{
			Debug: debug,
			Flash: parseBool("FLASH"),
			From:  from,
			Text:  text,
			To:    to,
		}

		resp, _err := seven.Sms.Json(params)
		err = _err
		fmt.Println("SMS sent:", resp)
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func parseBool(name string) bool {
	key := fmt.Sprintf("PLUGIN_%s", name)
	value := os.Getenv(key)
	return "1" == value || "true" == value
}
