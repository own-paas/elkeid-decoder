package schema

import "fmt"

type CallUserModeHelperExec struct {
	DataType string `json:"data_type"`
	Exe      string `json:"exe"`
	Argv     string `json:"argv"`
	Wait     string `json:"wait"`
}
func (self *CallUserModeHelperExec) Name() string {
	return self.dataType()
}

func (self *CallUserModeHelperExec) dataType() string {
	return "call_user_mode_helper_exec"
}

func (self *CallUserModeHelperExec) dataId() string {
	return "607"
}

func (self *CallUserModeHelperExec) dataLength() int {
	return 4
}

func (self *CallUserModeHelperExec) Parse(source []string) error {
	if len(source) != 4 {
		return fmt.Errorf("%s data format error", self.dataType())
	}

	self.DataType = self.dataType()
	self.Exe = source[1]
	self.Argv = source[2]
	self.Wait = source[3]

	return nil
}
