package command

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/ziyao233/trobot/types"
	"github.com/ziyao233/trobot/methods"
	"github.com/ziyao233/trobot/logger"
       )

type Command struct {
	Message			*types.Message
	Args			[]string
}

type handler struct {
	fn			func (Command) error
}

var commands map[string]handler = map[string]handler{}

func Register(cmd string, fn func (Command) error) {
	commands[cmd] = handler{ fn: fn }
}

func Handle(msg *types.Message) (bool, error) {
	if !strings.HasPrefix(msg.Text, "/") {
		return false, nil
	}

	args := parseCommand(msg.Text)
	logger.Debug([]byte(args[0]))
	if handler, ok := commands[args[0]]; ok {
		return true, handler.fn(Command{ Message: msg, Args: args })
	} else {
		return false, nil
	}
}

var cmdlineRegex *regexp.Regexp = regexp.MustCompile(`\w+`)
func parseCommand(s string) []string {
	return cmdlineRegex.FindAllString(s, -1)
}

func (c Command) Println(args... any) {
	methods.SendMessage(methods.SendMessageParam{
						Text:	fmt.Sprintln(args...),
						ChatID:	c.Message.Chat.ID,
					}, true)
}
