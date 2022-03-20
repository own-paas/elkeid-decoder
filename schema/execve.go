package schema

import (
	"fmt"
	"strconv"
)

type Execve struct {
	Common
	Argv      string `json:"argv"`
	RunPath   string `json:"run_path"`
	Stdin     string `json:"stdin"`
	Stdout    string `json:"stdout"`
	Dip       string `json:"dip"`
	Dport     int    `json:"dport"`
	Sip       string `json:"sip"`
	Sport     int    `json:"sport"`
	SaFamily  string `json:"sa_family"`
	PidTree   string `json:"pid_tree"`
	TTY       string `json:"tty"`
	SocketPid int    `json:"socket_pid"`
	SSH       string `json:"ssh"`
	LdPreload int    `json:"ld_preload"`
	Res       int    `json:"res"`
}

func (self *Execve) Name() string {
	return self.dataType()
}

func (self *Execve) dataType() string {
	return "execve"
}

func (self *Execve) dataId() string {
	return "59"
}

func (self *Execve) dataLength() int {
	return 28
}

func (self *Execve) Parse(source []string) error {
	var err error
	self.Common = Common{DataType: self.dataType()}

	if source[0] != self.dataId() || len(source) != self.dataLength() {
		return fmt.Errorf("%s data format error", self.dataType())
	}

	if err = self.Common.Parse(source[0:13]); err != nil {
		return fmt.Errorf("%s common: %v", self.dataType(), err)
	}

	self.Argv = source[13]

	self.RunPath = source[14]

	self.Stdin = source[15]

	self.Stdout = source[16]

	self.Dip = source[17]

	if self.Dport, err = strconv.Atoi(source[18]); err != nil {
		return fmt.Errorf("%s dport: %v", self.dataType(), err)
	}

	self.Sip = source[19]

	if self.Sport, err = strconv.Atoi(source[20]); err != nil {
		return fmt.Errorf("%s sport: %v", self.dataType(), err)
	}

	self.SaFamily = source[21]

	self.PidTree = source[22]

	self.TTY = source[23]

	if self.SocketPid, err = strconv.Atoi(source[24]); err != nil {
		return fmt.Errorf("%s socket_pid: %v", self.dataType(), err)
	}

	self.SSH = source[25]

	if self.LdPreload, err = strconv.Atoi(source[26]); err != nil {
		return fmt.Errorf("%s ld_preload: %v", self.dataType(), err)
	}

	if self.Res, err = strconv.Atoi(source[27]); err != nil {
		return fmt.Errorf("%s res: %v", self.dataType(), err)
	}

	return nil
}
