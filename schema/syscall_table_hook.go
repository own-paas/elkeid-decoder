package schema

import "fmt"

type SyscallTableHook struct {
	DataType      string `json:"data_type"`
	ModuleName    string `json:"module_name"`
	SyscallNumber string `json:"syscall_number"`
}

func (self *SyscallTableHook) Name() string {
	return self.dataType()
}

func (self *SyscallTableHook) dataType() string {
	return "syscall_table_hook"
}

func (self *SyscallTableHook) dataId() string {
	return "701"
}

func (self *SyscallTableHook) dataLength() int {
	return 3
}

func (self *SyscallTableHook) Parse(source []string) error {
	if len(source) != 3 {
		return fmt.Errorf("%s data format error", self.dataType())
	}

	self.DataType = self.dataType()
	self.ModuleName = source[1]
	self.SyscallNumber = source[2]

	return nil
}
