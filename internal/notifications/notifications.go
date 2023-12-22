package notifications

type Notifier interface {
	Notify(msg string) error
}
