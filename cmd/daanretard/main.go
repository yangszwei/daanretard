package main

import "daanretard/internal/registry"

func main() {
	app := registry.Prepare()
	err := app.Run()
	if err != nil {
		panic(err)
	}
}
