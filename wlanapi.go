package wlanapi

import (
	"syscall"
)

var (
	wlanapi = syscall.NewLazyDLL("wlanapi.dll")

	procWlanFreeMemory = wlanapi.NewProc("WlanFreeMemory")

	procWlanOpenHandle                       = wlanapi.NewProc("WlanOpenHandle")
	procWlanCloseHandle                      = wlanapi.NewProc("WlanCloseHandle")
	procWlanEnumInterfaces                   = wlanapi.NewProc("WlanEnumInterfaces")
	procWlanScan                             = wlanapi.NewProc("WlanScan")
	procWlanGetAvailableNetworkList          = wlanapi.NewProc("WlanGetAvailableNetworkList")
	procWlanQueryInterface                   = wlanapi.NewProc("WlanQueryInterface")
	wlanConnect                              = wlanapi.NewProc("WlanConnect")
	wlanDisconnect                           = wlanapi.NewProc("WlanDisconnect")
	wlanDeleteProfile                        = wlanapi.NewProc("WlanDeleteProfile")
	wlanGetFilterList                        = wlanapi.NewProc("WlanGetFilterList")
	wlanGetInterfaceCapability               = wlanapi.NewProc("WlanGetInterfaceCapability")
	wlanGetProfile                           = wlanapi.NewProc("WlanGetProfile")
	wlanGetNetworkBssList                    = wlanapi.NewProc("WlanGetNetworkBssList")
	wlanGetProfileCustomUserData             = wlanapi.NewProc("WlanGetProfileCustomUserData")
	wlanGetProfileList                       = wlanapi.NewProc("WlanGetProfileList")
	wlanGetSupportedDeviceServices           = wlanapi.NewProc("WlanGetSupportedDeviceServices")
	wlanGetSecuritySettings                  = wlanapi.NewProc("WlanGetSecuritySettings")
	wlanHostedNetworkForceStart              = wlanapi.NewProc("WlanHostedNetworkForceStart")
	wlanHostedNetworkForceStop               = wlanapi.NewProc("WlanHostedNetworkForceStop")
	wlanHostedNetworkInitSettings            = wlanapi.NewProc("WlanHostedNetworkInitSettings")
	wlanHostedNetworkQueryProperty           = wlanapi.NewProc("WlanHostedNetworkQueryProperty")
	wlanHostedNetworkQuerySecondaryKey       = wlanapi.NewProc("WlanHostedNetworkQuerySecondaryKey")
	wlanHostedNetworkQueryStatus             = wlanapi.NewProc("WlanHostedNetworkQueryStatus")
	wlanHostedNetworkRefreshSecuritySettings = wlanapi.NewProc("WlanHostedNetworkRefreshSecuritySettings")
	wlanHostedNetworkSetProperty             = wlanapi.NewProc("WlanHostedNetworkSetProperty")
	wlanHostedNetworkSetSecondaryKey         = wlanapi.NewProc("WlanHostedNetworkSetSecondaryKey")
	wlanHostedNetworkStartUsing              = wlanapi.NewProc("WlanHostedNetworkStartUsing")
	wlanHostedNetworkStopUsing               = wlanapi.NewProc("WlanHostedNetworkStopUsing")
	wlanIhvControl                           = wlanapi.NewProc("WlanIhvControl")
)
