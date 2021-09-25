package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// phone call example

type State int

const (
	OffHook State = iota
	Connecting
	Connected
	OnHold
	OnHook
)

func (s State) String() string {
	switch s {
	case OffHook:
		return "OffHook"
	case Connecting:
		return "Connecting"
	case Connected:
		return "Connected"
	case OnHold:
		return "OnHold"
	case OnHook:
		return "OnHook"
	}
	return "Unknown"
}

type Trigger int

const (
	CallDialed Trigger = iota
	HungUp
	CallConnected
	PlaceOnHold
	TakeOffHold
	LeftMessage
)

func (t Trigger) String() string {
	switch t {
	case CallDialed:
		return "CallDialed"
	case HungUp:
		return "HungUp"
	case CallConnected:
		return "CallConnected"
	case PlaceOnHold:
		return "PlaceOnHold"
	case TakeOffHold:
		return "TakeOffHold"
	case LeftMessage:
		return "LeftMessage"
	}
	return "Unknown"
}

type TriggerResult struct {
	Trigger Trigger
	State   State
}

var rules = map[State][]TriggerResult{
	OffHook: {
		{Trigger: CallDialed, State: Connecting},
	},
	Connecting: {
		{Trigger: HungUp, State: OnHold},
		{Trigger: CallConnected, State: Connected},
	},
	Connected: {
		{Trigger: LeftMessage, State: OnHold},
		{Trigger: HungUp, State: OnHold},
		{Trigger: PlaceOnHold, State: OnHold},
	},
	OnHold: {
		{Trigger: TakeOffHold, State: Connected},
		{Trigger: HungUp, State: OnHook},
	},
}

func main() {
	state, exitState := OffHook, OnHook

	for ok := true; ok; ok = state != exitState {
		fmt.Println("The phone is currently ", state)
		fmt.Println("Select a trigger:")

		for i := 0; i < len(rules[state]); i++ {
			transition := rules[state][i]
			fmt.Printf("%d. %s\n", i, transition.Trigger)
		}

		input, _, _ := bufio.NewReader(os.Stdin).ReadLine()

		choice, _ := strconv.Atoi(string(input))

		transition := rules[state][choice]
		state = transition.State
	}

	fmt.Println("We are done using the phone")
}
