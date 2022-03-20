package schema

import "fmt"

type LoadModule struct {
	Common
	KOFile  string `json:"ko_file"`
	PidTree string `json:"pid_tree"`
	RunPath string `json:"run_path"`
}

func (self *LoadModule) Name() string {
	return self.dataType()
}

func (self *LoadModule) dataType() string {
	return "load_module"
}

func (self *LoadModule) dataId() string {
	return "603"
}

func (self *LoadModule) dataLength() int {
	return 16
}

func (self *LoadModule) Parse(source []string) error {
	var err error
	self.Common = Common{DataType: self.dataType()}

	if source[0] != self.dataId() || len(source) != self.dataLength() {
		return fmt.Errorf("%s data format error", self.dataType())
	}

	if err = self.Common.Parse(source[0:13]); err != nil {
		return fmt.Errorf("%s common: %v", self.dataType(), err)
	}

	self.KOFile = source[13]
	self.PidTree = source[14]
	self.RunPath = source[15]

	return nil
}
