package event

import (
	"bufio"
	"fmt"
	"github.com/sestack/elkeid-decoder/parse"
	"github.com/sestack/elkeid-decoder/schema"
	"io"
	"os"
	"time"
	"golang.org/x/sys/unix"
)

const PIPE_PATH = "/proc/elkeid-endpoint"

type EventManager struct {
	pipePath string
	handles  map[string]schema.Schema
	receive  chan interface{}
}

func NewEventManager(filters []string, receive chan interface{}) *EventManager {
	event := &EventManager{
		pipePath: PIPE_PATH,
		receive:  receive,
		handles:  map[string]schema.Schema{},
	}

	if len(filters) > 0 {
		for _, k := range filters {
			if v, ok := schema.Handles[k]; ok {
				event.handles[k] = v
			}
		}
	} else {
		event.handles = schema.Handles
	}

	return event
}

func (e *EventManager) Parse(source []byte) (data interface{}, err error) {
	sourceData := parse.Decoding(source)
	if len(sourceData) == 0 {
		return nil, fmt.Errorf("event data is empty")
	}

	obj, ok := e.handles[sourceData[0]]
	if !ok {
		return nil, fmt.Errorf("event handler does not exist")
	}

	err = obj.Parse(sourceData)
	return obj, err
}

func (e *EventManager) Read() {
	for {
		file, err := os.Open(e.pipePath)
		if err != nil {
			fmt.Println("open file %s failed ,err:%v", e.pipePath, err)
			time.Sleep(time.Second * 10)
			continue
		}

		_, err = file.Stat()
		if err != nil {
			fmt.Println("failed to get file %s status ,err:%v", e.pipePath, err)
			time.Sleep(time.Second * 10)
			continue
		}
		unix.Fadvise(int(file.Fd()), 0, 0, unix.FADV_SEQUENTIAL)

		reader := bufio.NewReader(file)
		for {
			obj, err := reader.ReadSlice('\x17')
			if len(obj) != 0 {
				if data, err := e.Parse(obj); err == nil {
					e.receive <- data
				}
			}
			if err == io.EOF {
				file.Close()
				time.Sleep(time.Second * 10)
				break
			} else if err != nil {
				fmt.Println(err)
				file.Close()
				time.Sleep(time.Second * 10)
				break
			}
		}
	}
}
