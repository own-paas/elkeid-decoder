package schema

import "fmt"

type InterruptTableHook struct {
	DataType        string `json:"data_type"`
	ModuleName      string `json:"module_name"`
	InterruptNumber string `json:"interrupt_number"`
}

func (self *InterruptTableHook) Name() string {
	return self.dataType()
}

func (self *InterruptTableHook) dataType() string {
	return "interrupt_table_hook"
}

func (self *InterruptTableHook) dataId() string {
	return "703"
}

func (self *InterruptTableHook) dataLength() int {
	return 3
}

func (self *InterruptTableHook) Parse(source []string) error {
	if len(source) != 3 {
		return fmt.Errorf("%s data format error", self.dataType())
	}

	self.DataType = self.dataType()
	self.ModuleName = source[1]
	self.InterruptNumber = source[2]

	return nil
}
