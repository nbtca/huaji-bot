package utils

import (
	"net/http"
	"net/url"
)

func CallPushDeer(pushKey string, text string) error {
	pushAPI, _ := url.Parse("https://api2.pushdeer.com/message/push")
	params := url.Values{}
	params.Add("pushkey", pushKey)
	params.Add("text", text)
	pushAPI.RawQuery = params.Encode()
	_, err := http.Get(pushAPI.String())
	return err
}
