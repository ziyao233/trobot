/*
 *	trobot
 *	/types/user.go
 *	By Mozilla Public License Version 2.0
 *	Copyright (c) 2023 Yao Zi. All right reserved.
 */

package types

type User struct {
	ID		int64		`json:"id"`
	FirstName	string		`json:"first_name"`
	LastName	string		`json:"last_name,omitempty"`
	Username	string		`json:"username,omitempty"`
	IsBot		bool		`json:"is_bot,omitempty"`
}

func ToUser(i interface{}) User {
	if i == nil {
		return User{}
	}

	s := i.(map[string]interface{})
	return User{
			ID:		int64(s["id"].(float64)),
			FirstName:	s["first_name"].(string),
			LastName:	s["last_name"].(string),
			Username:	s["username"].(string),
			IsBot:		s["is_bot"].(bool),
		   }
}
