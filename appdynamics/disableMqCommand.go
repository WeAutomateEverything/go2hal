package appdynamics

import (
	"context"
	"github.com/weAutomateEverything/go2hal/alert"
	"github.com/weAutomateEverything/go2hal/telegram"
	"gopkg.in/telegram-bot-api.v4"
)

func NewDisableMqCommand(store Store, telegram telegram.Store, alert alert.Service) telegram.Command {
	return disableMq{
		store:    store,
		telegram: telegram,
		alert:    alert,
	}
}

type disableMq struct {
	store    Store
	telegram telegram.Store
	alert    alert.Service
}

func (disableMq) CommandIdentifier() string {
	return "disableMq"
}

func (disableMq) CommandDescription() string {
	return "Disabled MQ Monitoring and alerting"
}

func (disableMq) RestrictToAuthorised() bool {
	return true
}

func (s disableMq) Show(chat uint32) bool {
	i, err := s.store.GetAppDynamics(chat)
	if err != nil {
		return false

	}
	return len(i.MqEndpoints) > 0
}

func (s disableMq) Execute(ctx context.Context, update tgbotapi.Update) {
	id, err := s.telegram.GetUUID(update.Message.Chat.ID, update.Message.Chat.Title)
	if err != nil {
		s.alert.SendError(ctx, err)
		return
	}

	i, err := s.store.GetAppDynamics(id)
	if err != nil {
		s.alert.SendError(ctx, err)
		return
	}
	for x, _ := range i.MqEndpoints {
		i.MqEndpoints[x].Disabled = true
	}

	err = s.store.updateAppDynamics(*i)
	if err != nil {
		s.alert.SendError(ctx, err)
		return
	}
	s.alert.SendAlert(ctx, id, "MQ Monitoring has been successfully disabled.")
}
