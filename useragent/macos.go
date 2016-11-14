package useragent

import (
	"github.com/ezzizzle/uaparser/regex"
	"regexp"
	"strings"
)

var (
	macRegex    = regexp.MustCompile("\\((?P<Model>[^;]+); Intel Mac OS X (?P<OsVersion>[\\d\\_\\.]+)")
	safariRegex = regexp.MustCompile("Safari/([\\d\\.]+)")
)

// iOS Web View and Safari User Agents are in the format
// Mozilla/5.0 (iPhone; CPU iPhone OS 5_0_1 like Mac OS X) AppleWebKit/534.46 (KHTML, like Gecko) Mobile/9A405

func (device *Device) parseMacOSUA() {
	match := macRegex.FindStringSubmatch(device.UserAgent)

	if len(match) > 0 {
		deviceMatch := regex.MapRegexNames(match, macRegex.SubexpNames())

		device.AppName = "Safari"
		if deviceMatch["Model"] != "" {
			device.Model = deviceMatch["Model"]
		}

		if deviceMatch["OsVersion"] != "" {
			device.OsVersion = strings.Replace(deviceMatch["OsVersion"], "_", ".", -1)
		}
	}

	if device.AppName == "Safari" && strings.Contains(device.UserAgent, "Safari/") {
		versionMatch := safariRegex.FindStringSubmatch(device.UserAgent)
		if len(versionMatch) >= 1 {
			device.AppVersion = versionMatch[1]
		}
	}

	device.getAppDetails()
	device.getWebView()
}
