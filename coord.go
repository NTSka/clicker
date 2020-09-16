package main

import (
	"encoding/json"
	"fmt"
	"github.com/NTSka/clicker/config"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"io/ioutil"
	"log"
	"os"
	"path"
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
		conf = &config.Config{
			X:    0,
			Y:    0,
			Time: make([]string, 1),
		}
	} else {
		if err := json.Unmarshal(file, conf); err != nil {
			conf = &config.Config{
				X:    0,
				Y:    0,
				Time: make([]string, 0),
			}
		}
	}

	robotgo.EventHook(hook.KeyDown, []string{"ctrl", "w"}, func(e hook.Event) {

		x, y := robotgo.GetMousePos()

		conf.X = x
		conf.Y = y

		raw, err := json.Marshal(conf)
		if err != nil {
			fmt.Println(err)
		}

		if err := ioutil.WriteFile(confPath, raw, os.ModePerm); err != nil {
			fmt.Println(err)
		}

		robotgo.EventEnd()
	})

	s := robotgo.EventStart()
	<-robotgo.EventProcess(s)
}
