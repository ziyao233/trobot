/*
 *	trobot
 *	/methods/sendmessage.go
 *	By Mozilla Public License Version 2.0
 *	Copyright (c) 2023 Yao Zi. All rights reserved.
 */

package methods

import (
	"github.com/ziyao233/trobot/logger"
       )

type SendMessageParam struct {
	ChatID			int64	`json:"chat_id"`
	Text			string	`json:"text"`
	ReplyTo			int64	`json:"reply_to_message_id"`
}

func SendMessage(p SendMessageParam, force bool) error {
	_, err := call("sendMessage", p)
	if err != nil {
		logger.Error(err)
	}
	return err
}
