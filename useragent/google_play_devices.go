package useragent

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

// AndroidDevice is a struct for holding the details from
// the list of Google Play devices
type AndroidDevice struct {
	RetailBranding string
	MarketingName  string
	Device         string
	Model          string
	prettyName     string
}

var androidDeviceMapping = map[string]AndroidDevice{}

// PrettyName is a combination of the manufacturer, branding name
// and model
func (android *AndroidDevice) PrettyName() string {
	if android.prettyName != "" {
		return android.prettyName
	}
	if android.Model != android.MarketingName {
		android.prettyName = fmt.Sprintf("%s %s (%s)", android.RetailBranding, android.MarketingName, android.Model)
	} else {
		android.prettyName = fmt.Sprintf("%s %s", android.RetailBranding, android.Model)
	}
	return android.prettyName
}

// Print prints a JSON representation of the device
func (android *AndroidDevice) Print() {
	jsonOutput, err := json.MarshalIndent(android, "", "    ")
	if err == nil {
		fmt.Println(string(jsonOutput))
	} else {
		fmt.Println("ERROR Marshaling to Json:", err)
	}
}

// AndroidDeviceMappings is a map of model -> androidDevice
// taken from the list of Google Play approved devices
func AndroidDeviceMappings() map[string]AndroidDevice {
	if len(androidDeviceMapping) > 0 {
		return androidDeviceMapping
	}

	deviceCsv, _ := Asset("reference/supported_devices.csv")
	deviceCsvString, _ := DecodeUTF16(deviceCsv)

	reader := csv.NewReader(strings.NewReader(deviceCsvString))

	rowIndex := 0
	var headers []string
	for {
		record, err := reader.Read()
		// Stop at EOF.
		if err == io.EOF {
			break
		}

		if rowIndex == 0 {
			// This is the header row
			headers = record
			rowIndex++
			continue
		}

		if len(record) < 4 {
			continue
		}

		recordMap := csvRowToMap(record, headers)
		var android AndroidDevice
		android.RetailBranding = recordMap["Retail Branding"]
		android.MarketingName = recordMap["Marketing Name"]
		android.Device = recordMap["Device"]
		android.Model = recordMap["Model"]

		androidDeviceMapping[android.Model] = android
		rowIndex++
	}
	return androidDeviceMapping
}
