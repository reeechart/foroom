package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/signal"

	foroomerrors "github.com/reeechart/foroom/errors"
)

var (
	scanner *bufio.Scanner
)

func main() {
	user, room, err := parseUserAndRoom()
	checkError(err)
	initiateForoomSession(user, room)

	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt)

	<-interruptChan

	closeForoomSession()
}

func parseUserAndRoom() (string, string, error) {
	userPtr := flag.String("user", "", "user name in the room")
	roomPtr := flag.String("room", "", "room name to enter")
	flag.Parse()

	if *userPtr == "" && *roomPtr == "" {
		return "", "", foroomerrors.ErrInvalidArgs
	}
	return *userPtr, *roomPtr, nil
}

func initiateForoomSession(user, room string) {
	fmt.Printf("Welcome, %s to room %s @ Foroom\n", user, room)
	fmt.Println("Start chatting!")
}

func receiveRoomName() string {
	scanner.Scan()
	return scanner.Text()
}

func closeForoomSession() {
	fmt.Println()
	fmt.Println("See you next time!")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
