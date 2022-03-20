package schema

import (
	"fmt"
	"strconv"
)

type DnsQuery struct {
	Common
	Query    string `json:"query"`
	SaFamily int    `json:"sa_family"`
	Dip      string `json:"dip"`
	Dport    int    `json:"dport"`
	Sip      string `json:"sip"`
	Sport    int    `json:"sport"`
	Opcode   int    `json:"opcode"`
	Rcode    int    `json:"rcode"`
}

func (self *DnsQuery) Name() string {
	return self.dataType()
}

func (self *DnsQuery) dataType() string {
	return "dns_query"
}

func (self *DnsQuery) dataId() string {
	return "601"
}

func (self *DnsQuery) dataLength() int {
	return 21
}

func (self *DnsQuery) Parse(source []string) error {
	var err error
	self.Common = Common{DataType: self.dataType()}

	if source[0] != self.dataId() || len(source) != self.dataLength() {
		return fmt.Errorf("%s data format error", self.dataType())
	}

	if err = self.Common.Parse(source[0:13]); err != nil {
		return fmt.Errorf("%s common: %v", self.dataType(), err)
	}

	self.Query = source[13]

	if self.SaFamily, err = strconv.Atoi(source[14]); err != nil {
		return fmt.Errorf("%s sa_family: %v", self.dataType(), err)
	}

	self.Dip = source[15]

	if self.Dport, err = strconv.Atoi(source[16]); err != nil {
		return fmt.Errorf("%s dport: %v", self.dataType(), err)
	}

	self.Sip = source[17]

	if self.Sport, err = strconv.Atoi(source[18]); err != nil {
		return fmt.Errorf("%s sport: %v", self.dataType(), err)
	}

	if self.Opcode, err = strconv.Atoi(source[19]); err != nil {
		return fmt.Errorf("%s opcode: %v", self.dataType(), err)
	}

	if self.Rcode, err = strconv.Atoi(source[20]); err != nil {
		return fmt.Errorf("%s rcode: %v", self.dataType(), err)
	}

	return nil
}
