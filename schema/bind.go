package schema

import (
	"fmt"
	"strconv"
)

type Bind struct {
	Common
	SaFamily int    `json:"sa_family"`
	Sip      string `json:"sip"`
	Sport    int    `json:"sport"`
	Res      int    `json:"res"`
}

func (self *Bind) Name() string {
	return self.dataType()
}

func (self *Bind) dataType() string {
	return "bind"
}

func (self *Bind) dataId() string {
	return "49"
}

func (self *Bind) dataLength() int {
	return 17
}

func (self *Bind) Parse(source []string) error {
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

	self.Sip = source[14]

	if self.Sport, err = strconv.Atoi(source[15]); err != nil {
		return fmt.Errorf("%s sport: %v", err, self.dataType())
	}

	if self.Res, err = strconv.Atoi(source[16]); err != nil {
		return fmt.Errorf("%s res: %v", self.dataType(), err)
	}

	return nil
}
