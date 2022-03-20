package schema

import (
	"fmt"
	"strconv"
)

type CreateFile struct {
	Common
	FilePath  string `json:"file_path"`
	Dip       string `json:"dip"`
	Dport     int    `json:"dport"`
	Sip       string `json:"sip"`
	Sport     int    `json:"sport"`
	SaFamily  int    `json:"sa_family"`
	SocketPid int    `json:"socket_pid"`
	SbId      string `json:"sb_id"`
}

func (self *CreateFile) Name() string {
	return self.dataType()
}

func (self *CreateFile) dataType() string {
	return "create_file"
}

func (self *CreateFile) dataId() string {
	return "602"
}

func (self *CreateFile) dataLength() int {
	return 21
}

func (self *CreateFile) Parse(source []string) error {
	var err error
	self.Common = Common{DataType: self.dataType()}

	if source[0] != self.dataId() || len(source) != self.dataLength() {
		return fmt.Errorf("%s data format error", self.dataType())
	}

	if err = self.Common.Parse(source[0:13]); err != nil {
		return fmt.Errorf("%s common: %v", self.dataType(), err)
	}

	self.FilePath = source[13]

	self.Dip = source[14]

	if self.Dport, err = strconv.Atoi(source[15]); err != nil {
		return fmt.Errorf("%s dport: %v", self.dataType(), err)
	}

	self.Sip = source[16]

	if self.Sport, err = strconv.Atoi(source[17]); err != nil {
		return fmt.Errorf("%s sport: %v", self.dataType(), err)
	}

	if self.SaFamily, err = strconv.Atoi(source[18]); err != nil {
		return fmt.Errorf("%s sa_family: %v", self.dataType(), err)
	}

	if self.SocketPid, err = strconv.Atoi(source[19]); err != nil {
		return fmt.Errorf("%s socket_pid: %v", self.dataType(), err)
	}

	self.SbId = source[20]

	return nil
}
