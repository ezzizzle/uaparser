package useragent

import (
	"github.com/ezzizzle/uaparser/regex"
	"regexp"
)

// Android User Agents are in the format
// Mozilla/5.0 (Linux; <Android Version>; <Build Tag etc.>) \
//      AppleWebKit/<WebKit Rev> (KHTML, like Gecko) Chrome/<Chrome Rev> \
//      Mobile Safari/<WebKit Rev>

var androidRegex = regexp.MustCompile("\\(Linux;( U;)? Android (?P<AVersion>[^;]+); ([a-z]{2}-[a-z]{2}; )?(?P<Model>.*?) Build/[^;\\)]+(?P<WebView>; wv)?\\)")

func (device *Device) parseAndroidUA() {

	match := androidRegex.FindStringSubmatch(device.UserAgent)

	if len(match) > 0 {
		deviceMatch := regex.MapRegexNames(match, androidRegex.SubexpNames())

		if deviceMatch["Model"] != "" {
			device.Model = deviceMatch["Model"]
		}

		if deviceMatch["AVersion"] != "" {
			device.OsVersion = deviceMatch["AVersion"]
		}

		if deviceMatch["WebView"] != "" {
			device.IsWebView = true
		}
	}

	if androidDevice, ok := AndroidDeviceMappings()[device.Model]; ok {
		// Use the pretty name
		device.Model = androidDevice.PrettyName()
	}

	device.getAppDetails()
}
