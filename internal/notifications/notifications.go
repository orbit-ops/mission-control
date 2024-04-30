package notifications

import (
	"context"
	"log"

	"github.com/go-pkgz/notify"
	"github.com/slack-go/slack"
)

type Notifier interface {
	Notify(msg string) error
}

func Notifications() {
	wh := notify.NewSlack(
		"token",
		slack.OptionDebug(true), // optional, you can pass any slack.Options
	)
	err := wh.Send(context.Background(), "slack:general", "Hello, World!")
	if err != nil {
		log.Fatalf("problem sending message using slack, %v", err)
	}
}
