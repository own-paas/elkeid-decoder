package schema

import (
	"fmt"
	"strconv"
)

type Common struct {
	DataType  string `json:"data_type"`
	UID       int    `json:"uid"`
	EXE       string `json:"exe"`
	Pid       int    `json:"pid"`
	PPid      int    `json:"ppid"`
	PGid      int    `json:"pgid"`
	TGid      int    `json:"tgid"`
	Sid       int    `json:"sid"`
	Comm      string `json:"comm"`
	NodeName  string `json:"node_name"`
	SessionID int    `json:"sessionid"`
	Pns       int    `json:"pns"`
	RootPns   int    `json:"root_pns"`
}

func (self *Common) Parse(source []string) error {
	var err error

	if len(source) != 13 {
		return fmt.Errorf("common data format error")
	}

	if self.UID, err = strconv.Atoi(source[1]); err != nil {
		return fmt.Errorf("common uid: %v", err)
	}

	self.EXE = source[2]

	if self.Pid, err = strconv.Atoi(source[3]); err != nil {
		return fmt.Errorf("common pid: %v", err)
	}

	if self.PPid, err = strconv.Atoi(source[4]); err != nil {
		return fmt.Errorf("common ppid: %v", err)
	}

	if self.PGid, err = strconv.Atoi(source[5]); err != nil {
		return fmt.Errorf("common pgid: %v", err)
	}

	if self.TGid, err = strconv.Atoi(source[6]); err != nil {
		return fmt.Errorf("common tgid: %v", err)
	}

	if self.Sid, err = strconv.Atoi(source[7]); err != nil {
		return fmt.Errorf("common sid: %v", err)
	}

	self.Comm = source[8]

	self.NodeName = source[9]

	if self.SessionID, err = strconv.Atoi(source[10]); err != nil {
		return fmt.Errorf("common session_id: %v", err)
	}

	if self.Pns, err = strconv.Atoi(source[11]); err != nil {
		return fmt.Errorf("common pns: %v", err)
	}

	if self.RootPns, err = strconv.Atoi(source[12]); err != nil {
		return fmt.Errorf("common root_pns: %v", err)
	}

	return nil
}
