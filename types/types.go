/*
 *	trobot
 *	/types/types.go
 *	By Mozilla Public License Version 2.0
 *	Copyright (c) 2023 Yao Zi. All rights reserved.
 */

package types

type g interface{}

func FString(i g, f string) string {
	m := i.(map[string]interface{})
	if v, ok := m[f]; ok {
		return v.(string)
	} else {
		return ""
	}
}

func FInt64(i g, f string) int64 {
	m := i.(map[string]interface{})
	if v, ok := m[f]; ok {
		return v.(int64)
	} else {
		return 0
	}
}

func FBool(i g, f string) bool {
	m := i.(map[string]interface{})
	if v, ok := m[f]; ok {
		return v.(bool)
	} else {
		return false
	}
}

func FFloat64(i g, f string) float64 {
	m := i.(map[string]interface{})
	if v, ok := m[f]; ok {
		return v.(float64)
	} else {
		return 0
	}
}

func FGeneric(i g, f string) g {
	return i.(map[string]interface{})[f]
}
