package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/NTSka/clicker/config"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"io/ioutil"
	"log"
	"os"
	"path"
	"time"
)

func main() {
	root, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	conf := new(config.Config)

	confPath := path.Join(root, "./config.json")
	file, err := ioutil.ReadFile(confPath)
	if os.IsNotExist(err) {
		log.Fatal("Create config file first")
	}

	if err := json.Unmarshal(file, conf); err != nil {
		log.Fatal("Can't parse config file")
	}

	ticker := time.NewTicker(time.Second)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		robotgo.EventHook(hook.KeyDown, []string{"ctrl", "shift", "q"}, func(e hook.Event) {
			robotgo.EventEnd()
		})

		s := robotgo.EventStart()
		<-robotgo.EventProcess(s)
		cancel()
	}()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Done")
			return
		case <-ticker.C:
			fmt.Println("Click")
			robotgo.MoveClick(conf.X, conf.Y, "left", true)
		}
	}
}
