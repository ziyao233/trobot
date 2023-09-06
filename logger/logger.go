/*
 *	trobot
 *	/logger/logger.go
 *	By Mozilla Public License Version 2.0
 *	Copyright (c) 2023 Yao Zi. All rights reserved.
 */

package logger

import (
	"log"
	"os"
       )

type Level int

const (
	LFatal		= 0
	LError		= 1
	LWarning	= 2
	LInfo		= 3
	LDebug		= 4
      )

var loglevel Level = LWarning

func SetLogLevel(l Level) {
	loglevel = l
}

func Init(path string) {
	if path == "" {
		log.SetOutput(os.Stderr)
		path = "<stderr>"
	} else {
		f, err := os.OpenFile(path, os.O_CREATE | os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		log.SetOutput(f)
	}
	Logf(LInfo, "Start logging on %s with loglevel %d\n", path, loglevel)
}

func Log(l Level, args ...any) {
	if l < loglevel {
		return
	}
	log.Println(args...)
}

func Logf(l Level, f string, args ...any) {
	if l < loglevel {
		return
	}
	log.Printf(f, args...)
}

func Fatal(args ...any) {
	Log(LFatal, args...)
	os.Exit(1)
}

func Fatalf(f string, args ...any) {
	Logf(LFatal, f, args...)
	os.Exit(1)
}

func Error(args ...any) {
	Log(LError, args...)
}

func Errorf(f string, args ...any) {
	Logf(LError, f, args...)
}

func Debug(args ...any) {
	Log(LDebug, args...)
}

func Debugf(f string, args ...any) {
	Logf(LDebug, f, args...)
}
