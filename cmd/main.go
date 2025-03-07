package main

import (
	"fmt"
	"golang_auto_shop/internal/app"
	"time"
)

func main() {
	time.Sleep(5 * time.Second)
	
	app := app.NewApp()
	err := app.StartApp()
	fmt.Println(err)
	app.Storage.Example()
}