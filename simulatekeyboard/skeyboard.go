package main

import (
	"math/rand"
	"time"

	"github.com/micmonay/keybd_event"
)

func main() {
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		panic(err)
	}

	for {
		// Select keys to be pressed
		kb.SetKeys(rand.Intn(50))
		// Or you can use Press and Release
		kb.Press()
		time.Sleep(10 * time.Second)
		kb.Release()
	}

}
