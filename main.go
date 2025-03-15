package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/micmonay/keybd_event"
	"github.com/peterhellberg/acr122u"
	"os"
)

func main() {
	fileName := "mappings.json"
	mappings := make(map[uint32]string)

	file, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	if file != nil {
		json.Unmarshal(file, &mappings)
	}

	ctx, err := acr122u.EstablishContext()
	if err != nil {
		panic(err)
	}

	cleanup := func() {
		json, err := json.Marshal(mappings)
		if err != nil {
			panic(err)
		}
		os.WriteFile(fileName, json, 0644)
	}

	ctx.ServeFunc(func(c acr122u.Card) {
		var key string

		uid := binary.LittleEndian.Uint32(c.UID())

		if _, ok := mappings[uid]; ok {
			kb, err := keybd_event.NewKeyBonding()
			if err != nil {
				panic(err)
			}

			switch mappings[uid] {
			case "SHIFT":
				kb.AddKey(keyboardMapping_SHIFT)
			case "SPACE":
				kb.AddKey(keyboardMapping_SPACE)
			default:
				kb.AddKey(keyboardMapping[mappings[uid][0]])
			}

			fmt.Println(keyboardMapping[mappings[uid][0]])
			err = kb.Launching()
			if err != nil {
				panic(err)
			}

			return
		}

		fmt.Println("enter key: ")
		_, err = fmt.Scanln(&key)
		if err != nil {
			panic(err)
		}

		mappings[uid] = key
		fmt.Printf("%s: %x\n", key, uid)
		cleanup()
		fmt.Println("saved!")
		return
	})
}
