package main

import (
	"context"
	"golang.design/x/clipboard"
	"os/exec"
)

func main() {
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}

	image := clipboard.Watch(context.TODO(), clipboard.FmtImage)

	go func() {
		ch := clipboard.Watch(context.TODO(), clipboard.FmtText)
		for data := range ch {
			_, err := exec.Command("notify-send", string(data)).Output()
			if err != nil {
				panic(err)
			}
		}
	}()

	for _ = range image {
		_, err := exec.Command("notify-send", string("image was copied")).Output()
		if err != nil {
			panic(err)
		}
	}
}
