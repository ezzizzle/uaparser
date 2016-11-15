package main

import (
	"flag"
	"fmt"
	"github.com/ezzizzle/uaparser/useragent"
	"os"
)

type uaSlice []string

func (ua *uaSlice) String() string {
	return fmt.Sprintf("%s", *ua)
}

func (ua *uaSlice) Set(value string) error {
	*ua = append(*ua, value)
	return nil
}

func main() {
	var userAgents uaSlice

	flag.Var(&userAgents, "ua", "User Agents to Parse (Specify multiple with multiple -ua args)")
	flag.Parse()

	if len(userAgents) < 1 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	uaMap := map[string]useragent.Device{}

	for _, userAgent := range userAgents {
		var device useragent.Device
		if foundDevice, ok := uaMap[userAgent]; ok {
			// UA In Map, return this device
			device = foundDevice
		} else {
			device = useragent.DeviceFromUA(userAgent)
			uaMap[userAgent] = device
		}

		device.Print()
	}
}

// Sample run
// go run main.go -ua "Mozilla/5.0 (Linux; Android 4.0.4; Galaxy Nexus Build/IMM76B) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.133 Mobile Safari/535.19" -ua "Mozilla/5.0 (Linux; Android 5.1.1; Nexus 5 Build/LMY48B; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/43.0.2357.65 Mobile Safari/537.36" -ua "Mozilla/5.0 (Linux; Android 4.0.4; Galaxy Nexus Build/IMM76B) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.133 Mobile Safari/535.19" -ua "Mozilla/5.0 (iPhone; CPU iPhone OS 10_0 like Mac OS X) AppleWebKit/602.1.38 (KHTML, like Gecko) Version/10.0 Mobile/14A5297c Safari/602.1"
