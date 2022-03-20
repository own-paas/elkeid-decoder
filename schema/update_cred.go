package schema

import (
	"fmt"
	"strconv"
)

type UpdateCred struct {
	Common
	PidTree string `json:"pid_tree"`
	OldUid  int    `json:"old_uid"`
	Res     int    `json:"res"`
}

func (self *UpdateCred) Name() string {
	return self.dataType()
}

func (self *UpdateCred) dataType() string {
	return "update_cred"
}

func (self *UpdateCred) dataId() string {
	return "604"
}

func (self *UpdateCred) dataLength() int {
	return 16
}

func (self *UpdateCred) Parse(source []string) error {
	var err error
	self.Common = Common{DataType: self.dataType()}

	if source[0] != self.dataId() || len(source) != self.dataLength() {
		return fmt.Errorf("%s data format error", self.dataType())
	}

	if err = self.Common.Parse(source[0:13]); err != nil {
		return fmt.Errorf("%s common: %v", self.dataType(), err)
	}

	self.PidTree = source[13]

	if self.OldUid, err = strconv.Atoi(source[14]); err != nil {
		return fmt.Errorf("%s old_uid: %v", self.dataType(), err)
	}

	if self.Res, err = strconv.Atoi(source[15]); err != nil {
		return fmt.Errorf("%s res: %v", self.dataType(), err)
	}

	return nil
}
