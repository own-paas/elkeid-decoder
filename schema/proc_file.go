package schema

import "fmt"

type ProcFile struct {
	DataType string `json:"data_type"`
	ModuleName      string `json:"module_name"`
}

func (self *ProcFile) Name() string {
	return self.dataType()
}

func (self *ProcFile) dataType() string {
	return "proc_file"
}

func (self *ProcFile) dataId() string {
	return "700"
}

func (self *ProcFile) dataLength() int {
	return 2
}

func (self *ProcFile) Parse(source []string) error {
	if len(source) != 2 {
		return fmt.Errorf("%s data format error", self.dataType())
	}

	self.DataType = self.dataType()
	self.ModuleName = source[1]

	return nil
}
