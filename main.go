package main

import (
	"context"
	"golang.design/x/clipboard"
	"os/exec"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	err := clipboard.Init()
	if err != nil {
		panic(err)
	}

	go func() {
		ch := clipboard.Watch(context.TODO(), clipboard.FmtText)
		for data := range ch {
			_, err := exec.Command("notify-send", string(data)).Output()
			if err != nil {
				panic(err)
			}
		}
	}()

	go func() {
		image := clipboard.Watch(context.TODO(), clipboard.FmtImage)
		for _ = range image {
			_, err := exec.Command("notify-send", string("image was copied")).Output()
			if err != nil {
				panic(err)
			}
		}
	}()

	wg.Wait()
}
