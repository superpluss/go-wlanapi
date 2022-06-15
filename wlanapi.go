package wlanapi

import (
	"syscall"
)

var (
	wlanapi                     = syscall.NewLazyDLL("wlanapi.dll")
	wlanOpenHandle              = wlanapi.NewProc("WlanOpenHandle")
	wlanScan                    = wlanapi.NewProc("WlanScan")
	fWlanCloseHandle            = wlanapi.NewProc("WlanCloseHandle")
	fWlanEnumInterfaces         = wlanapi.NewProc("WlanEnumInterfaces")
	wlanConnect                 = wlanapi.NewProc("WlanConnect")
	wlanDisconnect              = wlanapi.NewProc("WlanDisconnect")
	wlanDeleteProfile           = wlanapi.NewProc("WlanDeleteProfile")
	wlanGetAvailableNetworkList = wlanapi.NewProc("WlanGetAvailableNetworkList")
	wlanGetFilterList           = wlanapi.NewProc("WlanGetFilterList")
	wlanGetInterfaceCapability  = wlanapi.NewProc("WlanGetInterfaceCapability")
	wlanGetNetworkBssList       = wlanapi.NewProc("WlanGetNetworkBssList")
)
