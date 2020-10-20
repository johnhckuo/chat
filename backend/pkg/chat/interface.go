package chat

type Message interface {
	Push(string, string) error
	Pop(string) (*string, error)
}
