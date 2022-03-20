package schema

import "fmt"

type Link struct {
	Common
	OldName string `json:"old_name"`
	NewName string `json:"new_name"`
	SbId    string `json:"sb_id"`
}

func (self *Link) Name() string {
	return self.dataType()
}

func (self *Link) dataType() string {
	return "link"
}

func (self *Link) dataId() string {
	return "86"
}

func (self *Link) dataLength() int {
	return 16
}

func (self *Link) Parse(source []string) error {
	var err error
	self.Common = Common{DataType: self.dataType()}

	if source[0] != self.dataId() || len(source) != self.dataLength() {
		return fmt.Errorf("%s data format error", self.dataType())
	}

	if err = self.Common.Parse(source[0:13]); err != nil {
		return fmt.Errorf("%s common: %v", self.dataType(), err)
	}

	self.OldName = source[13]
	self.NewName = source[14]
	self.SbId = source[15]

	return nil
}
