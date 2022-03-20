package schema

import "fmt"

type Mount struct {
	Common
	PidTree  string `json:"pid_tree"`
	Dev      string `json:"dev"`
	FilePath string `json:"file_path"`
	FsType   string `json:"fstype"`
	Flags    string `json:"flags"`
}

func (self *Mount) Name() string {
	return self.dataType()
}

func (self *Mount) dataType() string {
	return "mount"
}

func (self *Mount) dataId() string {
	return "165"
}

func (self *Mount) dataLength() int {
	return 18
}

func (self *Mount) Parse(source []string) error {
	var err error
	self.Common = Common{DataType: self.dataType()}

	if source[0] != self.dataId() || len(source) != self.dataLength() {
		return fmt.Errorf("%s data format error", self.dataType())
	}

	if err = self.Common.Parse(source[0:13]); err != nil {
		return fmt.Errorf("%s common: %v", self.dataType(), err)
	}

	self.PidTree = source[13]
	self.Dev = source[14]
	self.FilePath = source[15]
	self.FsType = source[16]
	self.Flags = source[17]

	return nil
}
