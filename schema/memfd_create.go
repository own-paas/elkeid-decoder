package schema

import "fmt"

type MemfdCreate struct {
	Common
	FdName string `json:"fdname"`
	Flags  string `json:"flags"`
}

func (self *MemfdCreate) Name() string {
	return self.dataType()
}

func (self *MemfdCreate) dataType() string {
	return "memfd_create "
}

func (self *MemfdCreate) dataId() string {
	return "356"
}

func (self *MemfdCreate) dataLength() int {
	return 15
}

func (self *MemfdCreate) Parse(source []string) error {
	var err error
	self.Common = Common{DataType: self.dataType()}

	if source[0] != self.dataId() || len(source) != self.dataLength() {
		return fmt.Errorf("%s data format error", self.dataType())
	}

	if err = self.Common.Parse(source[0:13]); err != nil {
		return fmt.Errorf("%s common: %v", self.dataType(), err)
	}

	self.FdName = source[13]
	self.Flags = source[14]

	return nil
}
