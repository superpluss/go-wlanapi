package wlanapi

import (
	"syscall"
)

var (
	wlanapi             = syscall.NewLazyDLL("wlanapi.dll")
	fWlanOpenHandle     = wlanapi.NewProc("WlanOpenHandle")
	wlanConnect         = wlanapi.NewProc("WlanConnect")
	fWlanCloseHandle    = wlanapi.NewProc("WlanCloseHandle")
	fWlanEnumInterfaces = wlanapi.NewProc("WlanEnumInterfaces")
)
