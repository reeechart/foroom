package sender

import (
	"bufio"
	"os"
)

type InputListener struct {
	MsgChannel chan string
}

func (listener InputListener) GetInput() {
	scanner := bufio.NewScanner(os.Stdin)
	var msg string
	if scanner.Scan() {
		msg = scanner.Text()
	}
	listener.MsgChannel <- msg
}
