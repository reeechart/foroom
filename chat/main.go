package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"

	"github.com/joho/godotenv"
	"github.com/reeechart/foroom/chat/receiver"
	"github.com/reeechart/foroom/chat/sender"
	foroomerrors "github.com/reeechart/foroom/errors"
)

func main() {
	user, room, err := parseUserAndRoom()
	checkError(err)
	initiateForoomSession(user, room)

	err = godotenv.Load()
	checkError(err)

	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt)

	sender := sender.NewSender(interruptChan)
	receiver := receiver.NewReceiver(interruptChan)

	go sender.ListenAndSendUserInputs(room)
	go receiver.ConsumeMessages(room, 0)

	<-interruptChan

	closeForoomSession()
}

func parseUserAndRoom() (string, string, error) {
	userPtr := flag.String("user", "", "user name in the room")
	roomPtr := flag.String("room", "", "room name to enter")
	flag.Parse()

	if !argValid(*userPtr, *roomPtr) {
		return "", "", foroomerrors.ErrInvalidArgs
	}
	return *userPtr, *roomPtr, nil
}

func argValid(user, room string) bool {
	return (user != "" && room != "")
}

func initiateForoomSession(user, room string) {
	fmt.Printf("Welcome, %s to room %s @ Foroom\n", user, room)
	fmt.Println("Start chatting!")
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
