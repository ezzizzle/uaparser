package useragent

import (
    "github.com/ezzizzle/uaparser/regex"
    "regexp"
    "strings"
)

// iOS Web View and Safari User Agents are in the format
// Mozilla/5.0 (iPhone; CPU iPhone OS 5_0_1 like Mac OS X) AppleWebKit/534.46 (KHTML, like Gecko) Mobile/9A405

func (device *Device) parseIOSUA() {
    //iosRegex := regexp.MustCompile("\\((?P<Model>[^;]+); (CPU )?(iPhone )?OS (?P<OsVersion>[\\d\\_]+)")
    iosRegex := regexp.MustCompile("\\((?P<Model>[^;]+); (U; )?(CPU )?(iPhone )?OS (?P<OsVersion>[\\d\\_\\.]+)")

    match := iosRegex.FindStringSubmatch(device.UserAgent)

    if len(match) > 0 {
        deviceMatch := regex.MapRegexNames(match, iosRegex.SubexpNames())

        device.AppName = "Safari"
        if deviceMatch["Model"] != "" {
            device.Model = deviceMatch["Model"]
        }

        if deviceMatch["OsVersion"] != "" {
            device.OsVersion = strings.Replace(deviceMatch["OsVersion"], "_", ".", -1)
        }
    }

    if device.AppName == "Safari" && strings.Contains(device.UserAgent, "Safari/") {
        device.AppVersion = regexp.MustCompile("Safari/([\\d\\.]+)").FindStringSubmatch(device.UserAgent)[1]
    }

    // Chrome iOS
    if strings.Contains(device.UserAgent, "CriOS/") {
        device.AppName = "Google Chrome"
        device.AppVersion = regexp.MustCompile("CriOS/([\\d\\.]+)").FindStringSubmatch(device.UserAgent)[1]
    }

    device.getAppDetails()
    device.getWebView()
}
