package notifications

type SlackNotifier struct {
	Url string
}

func (sn *SlackNotifier) Notify(msg string) error {
	return nil
}
