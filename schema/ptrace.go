package schema

import (
	"fmt"
	"strconv"
)

type Ptrace struct {
	Common
	PtraceRequest string `json:"ptrace_request"`
	TargetPid     int    `json:"target_pid"`
	Addr          string `json:"addr"`
	Data          string `json:"data"`
	PidTree       string `json:"pid_tree"`
}

func (self *Ptrace) Name() string {
	return self.dataType()
}

func (self *Ptrace) dataType() string {
	return "ptrace"
}

func (self *Ptrace) dataId() string {
	return "101"
}

func (self *Ptrace) dataLength() int {
	return 18
}

func (self *Ptrace) Parse(source []string) error {
	var err error
	self.Common = Common{DataType: self.dataType()}

	if source[0] != self.dataId() || len(source) != self.dataLength() {
		return fmt.Errorf("%s data format error", self.dataType())
	}

	if err = self.Common.Parse(source[0:13]); err != nil {
		return fmt.Errorf("%s common: %v", self.dataType(), err)
	}

	self.PtraceRequest = source[13]

	if self.TargetPid, err = strconv.Atoi(source[14]); err != nil {
		return fmt.Errorf("%s target_pid: %v", err, self.dataType())
	}

	self.Addr = source[15]

	self.Data = source[16]

	self.PidTree = source[17]

	return nil
}
