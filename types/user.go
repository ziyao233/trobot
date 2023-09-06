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

	return User{
			ID:		int64(FFloat64(i, "id")),
			FirstName:	FString(i, "first_name"),
			LastName:	FString(i, "last_name"),
			Username:	FString(i, "username"),
			IsBot:		FBool(i, "is_bot"),
		   }
}
