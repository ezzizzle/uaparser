package useragent

import (
    "regexp"
    "strings"
)

// getAppName gets the app name from a user agent
func (device *Device) getAppDetails() {
    appPrefixRegex := regexp.MustCompile("^(.{1,100})/([\\d\\.]+)")
    if strings.Contains(device.UserAgent, " Edge/") {
        device.AppName = "Edge"
        device.AppVersion = regexp.MustCompile("Edge/([\\d\\.]+)").FindStringSubmatch(device.UserAgent)[1]
    } else if strings.Contains(device.UserAgent, " Chrome/") {
        device.AppName = "Chrome"
        device.getChromeDetails()
        device.getWebView()
    } else if strings.Contains(device.UserAgent, "CriOS/") {
        // Chrome iOS
        device.AppName = "Chrome"
        device.getChromeDetails()
    } else if strings.Contains(device.UserAgent, "Firefox/") {
        device.AppName = "Firefox"
        device.AppVersion = regexp.MustCompile("Firefox/([\\d\\.]+)").FindStringSubmatch(device.UserAgent)[1]
    } else if strings.Contains(device.UserAgent, "Trident/7.0; rv:11.0") {
        device.AppName = "Internet Explorer"
        device.AppVersion = "11.0"
    } else if strings.Contains(device.UserAgent, "; MSIE") {
        device.AppName = "Internet Explorer"
        device.AppVersion = regexp.MustCompile("; MSIE ([\\d\\.]+)").FindStringSubmatch(device.UserAgent)[1]
    } else if !strings.HasPrefix(device.UserAgent, "Mozilla/") && appPrefixRegex.MatchString(device.UserAgent) {
        match := appPrefixRegex.FindStringSubmatch(device.UserAgent)
        device.AppName = match[1]
        device.AppVersion = match[2]
    }
}

func (device *Device) getChromeDetails() {
    versionRegex := regexp.MustCompile("(Chrome|CriOS)/([\\d\\.]+)")
    version := versionRegex.FindStringSubmatch(device.UserAgent)
    if len(version) > 0 {
        device.AppVersion = version[2]
    }
}

// getWebView determines if a browser is an embedded
// web view or not
func (device *Device) getWebView() {
    if strings.Contains(device.UserAgent, "; wv)") {
        device.IsWebView = true
    } else if strings.Contains(device.UserAgent, "Chrome/30.0.0.0") {
        device.IsWebView = true
    } else if device.OperatingSystem == "Android" && strings.HasSuffix(device.UserAgent, "AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30") {
        device.IsWebView = true
    } else if regexp.MustCompile("Mobile/[A-Z0-9]+$").MatchString(device.UserAgent) {
        device.IsWebView = true
    }
}
