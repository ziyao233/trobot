/*
 *	trobot
 *	/methods/methods.go
 *	By Mozilla Public License Version 2.0
 *	Copyright (c) 2023 Yao Zi. All rights reserved.
 */

package methods

import (
	"fmt"
	"encoding/json"
	"net/http"
	"io"
	"bytes"
	"errors"

	"github.com/ziyao233/trobot/types"
	"github.com/ziyao233/trobot/logger"
       )

var apiURL		string		= "https://api.telegram.org/bot"
var apiToken		string		= ""

type Response struct {
	Okay		bool		`json:"ok"`
	Result		interface{}	`json:"result"`
	Description	string		`json:"description"`
}

func getMethodURL(method string) string {
	return fmt.Sprintf("%s%s/%s", apiURL, apiToken, method);
}

func SetAPIURL(api string) {
	apiURL = api
}

func SetAPIToken(token string) {
	apiToken = token
}

func call(method string, param interface{}) (result interface{}, err error) {
	rawParam, err := json.Marshal(param)
	res, err := http.Post(getMethodURL(method), "application/json",
			      bytes.NewReader(rawParam))
	if err != nil {
		return
	}

	rawRes, err := io.ReadAll(res.Body)
	if err != nil {
		return
	}

	jsonRes := &Response{}
	if err = json.Unmarshal(rawRes, jsonRes); err != nil {
		return
	}

	if !jsonRes.Okay {
		if jsonRes.Description != "" {
			err = errors.New(jsonRes.Description)
		} else {
			err = errors.New("Error when sending request to the server")
		}
		return
	}

	result = jsonRes.Result
	return
}

type GetUpdatesParam struct {
	Offset		int64		`json:"offset"`
	Timeout		int		`json:"timeout"`
	Allowed		[]string	`json:"allowed_updates,omitempty"`
}

type Update struct {
	ID		int64
	Message		types.Message
}

func GetUpdates(p GetUpdatesParam) (update []Update, err error) {
	gUpdates, err := call("getUpdates", p)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	rawUpdates := gUpdates.([]interface{})

	update = make([]Update, len(rawUpdates))
	for i, v := range(rawUpdates) {
		gUpdate := v.(map[string]interface{})
		update[i] = parseUpdate(gUpdate)
		i++
	}

	return
}

func parseUpdate(i interface{}) Update {
	return Update{
			ID:		int64(types.FFloat64(i, "update_id")),
			Message:
				types.ToMessage(types.FGeneric(i, "message")),
		     }
}
