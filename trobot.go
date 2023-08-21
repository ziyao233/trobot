/*
 *	trobot
 *	/trobot.go
 *	By Mozilla Public License Version 2.0.
 *	Copyright (c) 2023 Yao Zi. All rights reserved.
 */

package trobot

import (
	"github.com/ziyao233/trobot/methods"
	"github.com/ziyao233/trobot/logger"
       )

var pollingInterval	int    = 60

var running		bool
var logPath		string = ""

func SetPollingInterval(t int) {
	pollingInterval = t
}

func SetAPIToken(token string) {
	methods.SetAPIToken(token)
}

func SetAPIURL(api string) {
	methods.SetAPIURL(api)
}

func SetLogPath(p string) {
	logPath = p
}

func SetLogLevel(l logger.Level) {
	logger.SetLogLevel(l)
}

func doPolling(start int64) []methods.Update {
	p := methods.GetUpdatesParam {
					Offset:		start,
					Timeout:	pollingInterval,
				 }
	updates, err := methods.GetUpdates(p)

	if err != nil {
		logger.Log(logger.LError, err)
		return nil
	}

	return updates
}

func processUpdates(updates []methods.Update) int64 {
	var nextOff int64 = -1
	for _, v := range(updates) {
		logger.Debugf("Update %d\n", v.ID)
		logger.Debug(v.Message)
		if v.ID > nextOff {
			nextOff = v.ID
		}
	}
	return nextOff + 1
}

func Run() {
	logger.Init(logPath)
	running = true

	var off int64 = 0

	for updates := doPolling(-1);
	    running;
	    updates  = doPolling(off) {
		off = processUpdates(updates)
	}
}
