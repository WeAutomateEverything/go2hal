package httpSmoke

import (
	"fmt"
	"github.com/weAutomateEverything/go2hal/telegram"
	"golang.org/x/net/context"
	"gopkg.in/kyokomi/emoji.v1"
	"gopkg.in/telegram-bot-api.v4"
	"strconv"
)

type quietHttpAlertsCommand struct {
	telegramService telegram.Service
	service         Service
}

func NewQuietHttpAlertCommand(telegramService telegram.Service,
	service Service) telegram.Command {
	return &quietHttpAlertsCommand{telegramService: telegramService, service: service}
}

func (s quietHttpAlertsCommand) RestrictToAuthorised() bool {
	return true
}

func (quietHttpAlertsCommand) CommandIdentifier() string {
	return "SilenceSmoke"
}

func (quietHttpAlertsCommand) CommandDescription() string {
	return "Disables smoke alerts. The http checks will still run, and in the event it succeeds an alert will still be sent. Add an integer value to set the amount of time the alert will be quiet for"
}

func (s *quietHttpAlertsCommand)Show(chat uint32) bool {
	endpoints, err := s.service.getEndpoints(chat)
	if err != nil {
		return false
	}
	if len(endpoints) == 0 {
		return false
	}
	return true
}

func (s *quietHttpAlertsCommand) Execute(ctx context.Context, update tgbotapi.Update) {
	arg := update.Message.CommandArguments()
	if arg == "" {
		arg = "30"
	}

	interval, err := strconv.ParseInt(arg, 10, 16)
	if err != nil {
		s.telegramService.SendMessage(ctx, update.Message.Chat.ID, fmt.Sprintf("unable to use %v as an integer value", arg), update.Message.MessageID)
		return
	}
	s.service.setTimeOut(interval)
	s.telegramService.SendMessage(ctx, update.Message.Chat.ID, emoji.Sprintf(":zzz: smoke tests will now sleep for %v minutes", arg), update.Message.MessageID)
}
