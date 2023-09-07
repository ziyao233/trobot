package command

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/ziyao233/trobot/types"
	"github.com/ziyao233/trobot/methods"
       )

type Command struct {
	Message			*types.Message
	Args			[]string
	Markup			string
}

type handler struct {
	fn			func (Command) error
}

var commands map[string]handler = map[string]handler{}

func Register(cmd string, fn func (Command) error) {
	commands[cmd] = handler{ fn: fn }
}

func Handle(msg *types.Message) (bool, error) {
	text, found := strings.CutPrefix(msg.Text, "/")
	if !found {
		return false, nil
	}

	args := parseCommand(text)
	if len(args) < 1 {
		return false, nil
	}
	if handler, ok := commands[args[0]]; ok {
		return true, handler.fn(Command{ Message: msg, Args: args })
	} else {
		return false, nil
	}
}

var cmdlineRegex *regexp.Regexp = regexp.MustCompile(`\S+`)
func parseCommand(s string) []string {
	return cmdlineRegex.FindAllString(s, -1)
}

func (c Command) Println(args... any) {
	methods.SendMessage(methods.SendMessageParam{
						Text:	fmt.Sprintln(args...),
						ChatID:	c.Message.Chat.ID,
						ParseMode: c.Markup,
					}, true)
}

func (c Command) Printf(f string, args... any) {
	methods.SendMessage(methods.SendMessageParam{
						Text:	fmt.Sprintf(f, args...),
						ChatID:	c.Message.Chat.ID,
						ParseMode: c.Markup,
					}, true)
}

func (c Command) Reply(args... any) {
	methods.SendMessage(methods.SendMessageParam{
						Text:	fmt.Sprintln(args...),
						ChatID:	c.Message.Chat.ID,
						ReplyTo:c.Message.ID,
						ParseMode: c.Markup,
					}, true)
}

func (c Command) Replyf(f string, args... any) {
	methods.SendMessage(methods.SendMessageParam{
						Text:	fmt.Sprintf(f, args...),
						ChatID:	c.Message.Chat.ID,
						ReplyTo:c.Message.ID,
						ParseMode: c.Markup,
					}, true)
}

func (c Command) ReplyTo(ID int64, args... any) {
	methods.SendMessage(methods.SendMessageParam{
						Text:	fmt.Sprintln(args...),
						ChatID:	c.Message.Chat.ID,
						ReplyTo:ID,
						ParseMode: c.Markup,
					}, true)
}

func (c Command) ReplyTof(ID int64, f string, args... any) {
	methods.SendMessage(methods.SendMessageParam{
						Text:	fmt.Sprintf(f, args...),
						ChatID:	c.Message.Chat.ID,
						ReplyTo:ID,
						ParseMode: c.Markup,
					}, true)
}
