package schema

import (
	"fmt"
	"strconv"
)

type UsbDeviceEvent struct {
	Common
	ProductInfo  string `json:"product_info"`
	Manufacturer string `json:"manufacturer"`
	Serial       string `json:"serial"`
	Action       int    `json:"action"`
}

func (self *UsbDeviceEvent) Name() string {
	return self.dataType()
}

func (self *UsbDeviceEvent) dataType() string {
	return "usb_device_event"
}

func (self *UsbDeviceEvent) dataId() string {
	return "610"
}

func (self *UsbDeviceEvent) dataLength() int {
	return 17
}

func (self *UsbDeviceEvent) Parse(source []string) error {
	var err error
	self.Common = Common{DataType: self.dataType()}

	if source[0] != self.dataId() || len(source) != self.dataLength() {
		return fmt.Errorf("%s data format error", self.dataType())
	}

	if err = self.Common.Parse(source[0:13]); err != nil {
		return fmt.Errorf("%s common: %v", self.dataType(), err)
	}

	self.ProductInfo = source[13]
	self.Manufacturer = source[14]
	self.Serial = source[15]
	if self.Action, err = strconv.Atoi(source[16]); err != nil {
		return fmt.Errorf("%s action: %v", self.dataType(), err)
	}

	return nil
}
