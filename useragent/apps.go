package useragent

import (
	"regexp"
	"strings"
)

var (
	chromeRegex        = regexp.MustCompile("(Chrome|CriOS)/([\\d\\.]+)")
	edgeRegex          = regexp.MustCompile("Edge/([\\d\\.]+)")
	firefoxRegex       = regexp.MustCompile("Firefox/([\\d\\.]+)")
	ieRegex            = regexp.MustCompile("; MSIE( )?([\\d\\.]+)")
	appPrefixRegex     = regexp.MustCompile("^(.{1,100})/([\\d\\.]+)")
	webViewSafariRegex = regexp.MustCompile("Mobile/[A-Z0-9]+$")
)

// getAppName gets the app name from a user agent
func (device *Device) getAppDetails() {
	if strings.Contains(device.UserAgent, " Edge/") {
		device.AppName = "Edge"
		versionMatch := edgeRegex.FindStringSubmatch(device.UserAgent)
		if len(versionMatch) >= 1 {
			device.AppVersion = versionMatch[1]
		}
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
		versionMatch := firefoxRegex.FindStringSubmatch(device.UserAgent)
		if len(versionMatch) >= 1 {
			device.AppVersion = versionMatch[1]
		}
	} else if strings.Contains(device.UserAgent, "Trident/7.0; rv:11.0") {
		device.AppName = "Internet Explorer"
		device.AppVersion = "11.0"
	} else if strings.Contains(device.UserAgent, "; MSIE") {
		device.AppName = "Internet Explorer"
		versionMatch := ieRegex.FindStringSubmatch(device.UserAgent)
		if len(versionMatch) >= 1 {
			device.AppVersion = versionMatch[2]
		}
	} else if !strings.HasPrefix(device.UserAgent, "Mozilla/") && appPrefixRegex.MatchString(device.UserAgent) {
		match := appPrefixRegex.FindStringSubmatch(device.UserAgent)
		device.AppName = match[1]
		device.AppVersion = match[2]
	}
}

func (device *Device) getChromeDetails() {
	version := chromeRegex.FindStringSubmatch(device.UserAgent)
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
	} else if webViewSafariRegex.MatchString(device.UserAgent) {
		device.IsWebView = true
	}
}
