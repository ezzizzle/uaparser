package useragent

import (
	"strings"
)

func (device *Device) parseWindowsUA() {
	if strings.Contains(device.UserAgent, "Windows NT 10.0") {
		device.OsVersion = "Windows 10"
	} else if strings.Contains(device.UserAgent, "Windows NT 6.3") {
		device.OsVersion = "Windows 8.1"
	} else if strings.Contains(device.UserAgent, "Windows NT 6.2") {
		device.OsVersion = "Windows 8"
	} else if strings.Contains(device.UserAgent, "Windows NT 6.1") {
		device.OsVersion = "Windows 7"
	} else if strings.Contains(device.UserAgent, "Windows NT 6.0") {
		device.OsVersion = "Windows Vista"
	} else if strings.Contains(device.UserAgent, "Windows NT 5.2") {
		device.OsVersion = "Windows XP x64"
	} else if strings.Contains(device.UserAgent, "Windows NT 5.1") {
		device.OsVersion = "Windows XP"
	} else if strings.Contains(device.UserAgent, "Windows NT 5.0") {
		device.OsVersion = "Windows 2000"
	} else if strings.Contains(device.UserAgent, "Windows NT 5.01") {
		device.OsVersion = "Windows 2000 SP1"
	}

	device.getAppDetails()
}
