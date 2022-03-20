package schema

import (
	"fmt"
	"strconv"
)

type Connect struct {
	Common
	SaFamily int    `json:"sa_family"`
	Dip      string `json:"dip"`
	Dport    int    `json:"dport"`
	Sip      string `json:"sip"`
	Sport    int    `json:"sport"`
	Res      int    `json:"res"`
}

func (self *Connect) Name() string {
	return self.dataType()
}

func (self *Connect) dataType() string {
	return "connect"
}

func (self *Connect) dataId() string {
	return "48"
}

func (self *Connect) dataLength() int {
	return 17
}

func (self *Connect) Parse(source []string) error {
	var err error
	self.Common = Common{DataType: self.dataType()}

	if source[0] != self.dataId() || len(source) != self.dataLength() {
		return fmt.Errorf("%s data format error", self.dataType())
	}

	if err = self.Common.Parse(source[0:13]); err != nil {
		return fmt.Errorf("%s common: %v", self.dataType(), err)
	}

	if self.SaFamily, err = strconv.Atoi(source[13]); err != nil {
		return fmt.Errorf("%s sa_family: %v", self.dataType(), err)
	}

	if self.SaFamily, err = strconv.Atoi(source[13]); err != nil {
		return fmt.Errorf("%s sa_family: %v", self.dataType(), err)
	}

	self.Dip = source[14]

	if self.Dport, err = strconv.Atoi(source[15]); err != nil {
		return fmt.Errorf("%s dport: %v", self.dataType(), err)
	}

	self.Sip = source[16]

	if self.Sport, err = strconv.Atoi(source[17]); err != nil {
		return fmt.Errorf("%s sport: %v", self.dataType(), err)
	}

	if self.Res, err = strconv.Atoi(source[18]); err != nil {
		return fmt.Errorf("%s opcode: %v", self.dataType(), err)
	}

	return nil
}
