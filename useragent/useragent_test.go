package useragent

import (
	"bufio"
	"encoding/csv"
	// "fmt"
	"io"
	"os"
	"reflect"
	"testing"
)

func compareDevices(t *testing.T, knownDevice Device, testDevice Device) {
	knownReflect := reflect.ValueOf(knownDevice)
	testReflect := reflect.ValueOf(testDevice)

	t.Errorf("%s", knownDevice.UserAgent)
	for i := 0; i < knownReflect.NumField(); i++ {
		knownValue := knownReflect.Field(i).Interface()
		testValue := testReflect.Field(i).Interface()
		if knownValue != testValue {
			val := reflect.Indirect(knownReflect)
			fieldName := val.Type().Field(i).Name
			t.Errorf("Device.%s Expected '%s' got '%s'", fieldName, knownValue, testValue)
		}
	}
}

// TestUAParsing reads in a csv file containing User Agents
// and their field mappings
// Test that the CSV gives the same results as the UAParser
func TestUAParsing(t *testing.T) {
	// Read in the list of UAs from CSV
	uafile, _ := os.Open("testdata/useragents.csv")
	reader := csv.NewReader(bufio.NewReader(uafile))

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

		recordMap := csvRowToMap(record, headers)
		var csvDevice Device
		csvDevice.UserAgent = recordMap["UserAgent"]
		csvDevice.Model = recordMap["Model"]
		csvDevice.OperatingSystem = recordMap["OperatingSystem"]
		csvDevice.OsVersion = recordMap["OsVersion"]
		csvDevice.AppName = recordMap["AppName"]
		csvDevice.AppVersion = recordMap["AppVersion"]
		if recordMap["IsAntiVirus"] == "TRUE" {
			csvDevice.IsAntiVirus = true
		} else {
			csvDevice.IsAntiVirus = false
		}
		if recordMap["IsWebView"] == "TRUE" {
			csvDevice.IsWebView = true
		} else {
			csvDevice.IsWebView = false
		}

		testDevice := DeviceFromUA(recordMap["UserAgent"])

		if csvDevice != testDevice {
			t.Error("UA failed to parse correctly")
			compareDevices(t, csvDevice, testDevice)
		}

		rowIndex++
	}
	t.Logf("Parsed %d User agents", rowIndex)
}

func BenchmarkWindowsParsing(b *testing.B) {
	//b.ResetTimer()
	for n := 0; n < b.N; n++ {
		userAgent := "Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"
		_ = DeviceFromUA(userAgent)
	}
}

func BenchmarkAndroidParsing(b *testing.B) {
	//b.ResetTimer()
	for n := 0; n < b.N; n++ {
		userAgent := "Mozilla/5.0 (Linux; Android 5.1.1; Nexus 5 Build/LMY48B; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/43.0.2357.65 Mobile Safari/537.36"
		_ = DeviceFromUA(userAgent)
	}
}

func BenchmarkIOSParsing(b *testing.B) {
	//b.ResetTimer()
	for n := 0; n < b.N; n++ {
		userAgent := "Mozilla/5.0 (iPhone; CPU iPhone OS 10_0 like Mac OS X) AppleWebKit/602.1.38 (KHTML, like Gecko) Version/10.0 Mobile/14A5297c Safari/602.1"
		_ = DeviceFromUA(userAgent)
	}
}

func BenchmarkMacParsing(b *testing.B) {
	//b.ResetTimer()
	for n := 0; n < b.N; n++ {
		userAgent := "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12) AppleWebKit/602.1.50 (KHTML, like Gecko) Version/10.0 Safari/602.1.50"
		_ = DeviceFromUA(userAgent)
	}
}
