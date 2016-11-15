package useragent

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Device represents a device parsed from
// a user agent
type Device struct {
	UserAgent       string
	Model           string
	OperatingSystem string
	OsVersion       string
	AppName         string
	AppVersion      string
	IsAntiVirus     bool
	IsWebView       bool
}

// Print out a description of the device
func (device *Device) Print() {
	jsonOutput, err := json.MarshalIndent(device, "", "    ")
	if err == nil {
		fmt.Println(string(jsonOutput))
	} else {
		fmt.Println("ERROR Marshaling to Json:", err)
	}
}

// AsJSON returns a JSON representation of the device
func (device *Device) AsJSON() []byte {
	jsonOutput, err := json.Marshal(device)
	if err != nil {
		// Do something here when JSON can't be made

	}
	return jsonOutput
}

// DeviceFromUA creates a new device form a passed in
// user agent string
func DeviceFromUA(userAgent string) Device {
	device := Device{}
	device.UserAgent = userAgent
	device.profile()
	return device
}

// Profile a device from the user agent
func (device *Device) profile() {
	device.getOS()
}

// getOS Extracts the operating system from a user agent
func (device *Device) getOS() {
	// Android
	if strings.Contains(device.UserAgent, "Android") {
		device.OperatingSystem = "Android"
	} else if strings.Contains(device.UserAgent, "Windows") {
		device.OperatingSystem = "Windows"
	} else if strings.Contains(device.UserAgent, "iPhone") {
		device.OperatingSystem = "iOS"
	} else if strings.Contains(device.UserAgent, "like Mac OS X") {
		device.OperatingSystem = "iOS"
	} else if strings.Contains(device.UserAgent, "Macintosh") {
		device.OperatingSystem = "macOS"
	}

	switch device.OperatingSystem {
	case "Android":
		device.parseAndroidUA()
	case "Windows":
		device.parseWindowsUA()
	case "iOS":
		device.parseIOSUA()
	case "macOS":
		device.parseMacOSUA()
	}
}
