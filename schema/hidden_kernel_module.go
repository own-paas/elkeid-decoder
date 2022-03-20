package schema

import "fmt"

type HiddenKernelModule struct {
	DataType   string `json:"data_type"`
	ModuleName string `json:"module_name"`
}

func (self *HiddenKernelModule) Name() string {
	return self.dataType()
}

func (self *HiddenKernelModule) dataType() string {
	return "hidden_kernel_module"
}

func (self *HiddenKernelModule) dataId() string {
	return "702"
}

func (self *HiddenKernelModule) dataLength() int {
	return 2
}

func (self *HiddenKernelModule) Parse(source []string) error {
	if len(source) != 2 {
		return fmt.Errorf("%s data format error", self.dataType())
	}

	self.DataType = self.dataType()
	self.ModuleName = source[1]

	return nil
}
