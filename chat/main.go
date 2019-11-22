package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
)

var (
	scanner *bufio.Scanner
)

func main() {
	initiateForoomSession()

	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt)

	<-interruptChan
	fmt.Println()
	fmt.Println("Exited foroom")
}

func initiateForoomSession() {
	fmt.Println("Welcome to Foroom")
	fmt.Print("Please enter room name: ")
	scanner = bufio.NewScanner(os.Stdin)
	roomName := receiveRoomName()
	fmt.Printf("Welcome to room %s\n", roomName)
	fmt.Printf("Start chatting in room %s!\n", roomName)
}

func receiveRoomName() string {
	scanner.Scan()
	return scanner.Text()
}
