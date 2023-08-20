/*
 *	trobot
 *	/trobot.go
 *	By Mozilla Public License Version 2.0.
 *	Copyright (c) 2023 Yao Zi. All rights reserved.
 */

package trobot

import (
	"io"
	"fmt"
	"net/http"
       )

var apiURL	string = "https://api.telegram.org/bot"
var apiToken	string = ""

func SetAPIURL(api string) {
	apiURL = api;
}

func SetAPIToken(token string) {
	apiToken = token;
}

func GetMe() string {
	resp, err := http.Get(fmt.Sprintf("%s%s/getMe", apiURL, apiToken))
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	info, err := io.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	return string(info)
}
