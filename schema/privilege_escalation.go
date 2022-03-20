package schema

import (
	"fmt"
	"strconv"
)

type PrivilegeEscalation struct {
	Common
	PPid    int    `json:"p_pid"`
	PidTree string `json:"pid_tree"`
	PCred   string `json:"p_cred"`
	CCred   string `json:"c_cred"`
}

func (self *PrivilegeEscalation) Name() string {
	return self.dataType()
}

func (self *PrivilegeEscalation) dataType() string {
	return "privilege_escalation"
}

func (self *PrivilegeEscalation) dataId() string {
	return "611"
}

func (self *PrivilegeEscalation) dataLength() int {
	return 17
}

func (self *PrivilegeEscalation) Parse(source []string) error {
	var err error
	self.Common = Common{DataType: self.dataType()}

	if source[0] != self.dataId() || len(source) != self.dataLength() {
		return fmt.Errorf("%s data format error", self.dataType())
	}

	if err = self.Common.Parse(source[0:13]); err != nil {
		return fmt.Errorf("%s common: %v", self.dataType(), err)
	}

	if self.PPid, err = strconv.Atoi(source[13]); err != nil {
		return fmt.Errorf("%s p_pid: %v", self.dataType(), err)
	}

	self.PidTree = source[14]
	self.PCred = source[15]
	self.CCred = source[16]

	return nil
}
