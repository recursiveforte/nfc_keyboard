package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/peterhellberg/acr122u"
	"os"
)

func main() {
	fileName := "mappings.json"
	mappings := make(map[string]uint32)

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

		fmt.Println("enter key: ")
		_, err = fmt.Scanln(&key)
		if err != nil {
			panic(err)
		}

		uid := binary.LittleEndian.Uint32(c.UID())
		mappings[key] = uid
		fmt.Printf("%s: %x\n", key, uid)
		cleanup()
		fmt.Println("saved!")
		return
	})
}
