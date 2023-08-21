/*
 *	trobot
 *	/types/message.go
 *	By Mozilla Public License Version 2.0
 *	Copyright (c) 2023 Yao Zi. All rights reserved.
 */

package types

type Message struct {
	ID		int64		`json:"message_id"`
	From		User		`json:"from"`
	Date		int64		`json:"date"`
	Chat		Chat		`json:"chat"`
	Text		string		`json:"text"`
}

func ToMessage(i interface {}) Message {
	if i == nil {
		return Message{}
	}

	return Message{
			ID:		int64(FFloat64(i, "message_id")),
			From:		ToUser(FGeneric(i, "from")),
			Date:		int64(FFloat64(i, "Date")),
			Chat:		ToChat(FGeneric(i, "chat")),
			Text:		FString(i, "text"),
		      }
}

func (self Message) Type() string {
	return "message"
}
