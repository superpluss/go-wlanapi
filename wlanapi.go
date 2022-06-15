package wlanapi

import (
	"syscall"
)

var (
	wlanapi = syscall.NewLazyDLL("wlanapi.dll")

	wlanOpenHandle                           = wlanapi.NewProc("WlanOpenHandle")
	wlanScan                                 = wlanapi.NewProc("WlanScan")
	fWlanCloseHandle                         = wlanapi.NewProc("WlanCloseHandle")
	fWlanEnumInterfaces                      = wlanapi.NewProc("WlanEnumInterfaces")
	wlanConnect                              = wlanapi.NewProc("WlanConnect")
	wlanDisconnect                           = wlanapi.NewProc("WlanDisconnect")
	wlanDeleteProfile                        = wlanapi.NewProc("WlanDeleteProfile")
	wlanGetAvailableNetworkList              = wlanapi.NewProc("WlanGetAvailableNetworkList")
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
