package main

import (
	"context"
	"fmt"
	"os/exec"
	"runtime"
	"sync"

	"golang.design/x/clipboard"
)

func main() {
	runtime.GOMAXPROCS(2)
	var wg sync.WaitGroup
	wg.Add(2)

	err := clipboard.Init()
	if err != nil {
		panic(err)
	}

	go func() {
		defer wg.Done()
		ch := clipboard.Watch(context.TODO(), clipboard.FmtText)
		for data := range ch {
			_, err := exec.Command("notify-send", string(data), "--app-name", "go-notif").Output()
			if err != nil {
				panic(err)
			}
		}
	}()

	go func() {
		defer wg.Done()
		ch := clipboard.Watch(context.TODO(), clipboard.FmtImage)
		for _ = range ch {
			_, err := exec.Command("notify-send", string("image was copied")).Output()
			if err != nil {
				fmt.Println(fmt.Sprint(err))
			}
		}
	}()

	wg.Wait()
}
