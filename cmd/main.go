package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/sestack/elkeid-decoder/event"
	"github.com/sestack/elkeid-decoder/schema"
	ver "github.com/sestack/elkeid-decoder/version"
	"os"
	"strings"
)

var (
	filters = flag.String("filters", "", "event filtering ',' split multiple events")
	list = flag.Bool("list", false, "show event list ")
	version = flag.Bool("version", false, "show version")
)

func main() {
	flag.Parse()

	if *version {
		fmt.Printf("version: %s\n", ver.Version)
		os.Exit(0)
	}

	if *list {
		fmt.Printf("%s\t%s\n","ID","Event")
		for k,v := range schema.Handles{
			fmt.Printf("%s\t%s\n",k,v.Name())
		}
		os.Exit(0)
	}

	filterList := []string{}
	if *filters != "" {
	      filterList = strings.Split(*filters, ",")
	}

	receive := make(chan interface{})
	go func() {
		for obj := range receive {
			if data, err := json.Marshal(obj); err == nil {
				fmt.Println(string(data))
			}
		}
	}()

	eventManager := event.NewEventManager(filterList, receive)
	eventManager.Read()
}
