package main

import (
	"context"
	"fmt"
	"golang.design/x/clipboard"
	"os/exec"
)

func main() {
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}

	// ctx, cancel := context.WithTImeout(context.Background(), time.Second*2)
	// defer cancel()

	ch := clipboard.Watch(context.TODO(), clipboard.FmtText)
	for data := range ch {
		fmt.Println(string(data))
		_, err := exec.Command("notify-send", string(data)).Output()
		if err != nil {
			panic(err)
		}
	}
}
