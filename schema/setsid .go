package schema

import "fmt"

type Setsid struct {
	Common
}

func (self *Setsid) Name() string {
	return self.dataType()
}

func (self *Setsid) dataType() string {
	return "setsid"
}

func (self *Setsid) dataId() string {
	return "112"
}

func (self *Setsid) dataLength() int {
	return 13
}

func (self *Setsid) Parse(source []string) error {
	var err error
	self.Common = Common{DataType: self.dataType()}

	if source[0] != self.dataId() || len(source) != self.dataLength() {
		return fmt.Errorf("%s data format error", self.dataType())
	}

	if err = self.Common.Parse(source); err != nil {
		return fmt.Errorf("%s common: %v", self.dataType(), err)
	}

	return nil
}
