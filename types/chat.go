/*
 *	trobot
 *	/types/chat.go
 *	By Mozilla Public License Version 2.0
 *	Copyright (c) 2023 Yao Zi. All rights reserved.
 */

package types

type Chat struct {
	ID		int64		`json:"id"`
	Type		string		`json:"type"`
}

func ToChat(i interface{}) Chat {
	if i == nil {
		return Chat{}
	}

	return Chat{
			ID:		int64(FFloat64(i, "id")),
			Type:		FString(i, "type"),
		   }
}
