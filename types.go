package wlanapi

import "syscall"

const (
	MAX_INDEX = 1000
	S_OK      = 0
)

//A DOT11_SSID structure contains the SSID of an interface.
//https://docs.microsoft.com/en-us/windows/win32/nativewifi/dot11-ssid
type DOT11_SSID struct {
	uSSIDLength uint32
	ucSSID      [32]byte
}

//The WLAN_AVAILABLE_NETWORK_LIST structure contains an array of information about available networks.
//https://docs.microsoft.com/en-us/windows/win32/api/wlanapi/ns-wlanapi-wlan_available_network_list
//typedef struct _WLAN_AVAILABLE_NETWORK_LIST {
//  DWORD                  dwNumberOfItems;
//  DWORD                  dwIndex;
//#if ...
//  WLAN_AVAILABLE_NETWORK *Network[];
//#else
//  WLAN_AVAILABLE_NETWORK Network[1];
//#endif
//} WLAN_AVAILABLE_NETWORK_LIST, *PWLAN_AVAILABLE_NETWORK_LIST;
type WLAN_AVAILABLE_NETWORK_LIST struct {
	dwNumberOfItems uint32
	dwIndex         uint32
	Network         [MAX_INDEX + 1]WLAN_AVAILABLE_NETWORK
}

//The WLAN_AVAILABLE_NETWORK structure contains information about an available wireless network.
//https://docs.microsoft.com/en-us/windows/win32/api/wlanapi/ns-wlanapi-wlan_available_network
type WLAN_AVAILABLE_NETWORK struct {
	strProfileName              [256]uint16
	dot11Ssid                   DOT11_SSID
	dot11BssType                uint32
	uNumberOfBssids             uint32
	bNetworkConnectable         int32
	wlanNotConnectableReason    uint32
	uNumberOfPhyTypes           uint32
	dot11PhyTypes               [8]uint32
	bMorePhyTypes               int32
	wlanSignalQuality           uint32
	bSecurityEnabled            int32
	dot11DefaultAuthAlgorithm   uint32
	dot11DefaultCipherAlgorithm uint32
	dwFlags                     uint32
	dwReserved                  uint32
}

type WLAN_BSS_LIST struct {
	dwTotalSize     uint32
	dwNumberOfItems uint32
	wlanBssEntries  [MAX_INDEX + 1]WLAN_BSS_ENTRY
}

type WLAN_BSS_ENTRY struct {
	dot11Ssid               DOT11_SSID
	uPhyId                  uint32
	dot11Bssid              [6]byte
	dot11BssType            uint32
	dot11BssPhyType         uint32
	lRssi                   int32
	uLinkQuality            uint32
	bInRegDomain            int32
	usBeaconPeriod          uint16
	ullTimestamp            uint64
	ullHostTimestamp        uint64
	usCapabilityInformation uint16
	ulChCenterFrequency     uint32
	wlanRateSet             WLAN_RATE_SET
	ulIeOffset              uint32
	ulIeSize                uint32
}

type WLAN_RATE_SET struct {
	uRateSetLength uint32
	usRateSet      [126]uint16
}

type WLAN_INTERFACE_INFO struct {
	InterfaceGuid           syscall.GUID
	strInterfaceDescription [256]uint16
	isState                 uint32
}

type WLAN_INTERFACE_INFO_LIST struct {
	dwNumberOfItems uint32
	dwIndex         uint32
	InterfaceInfo   [MAX_INDEX + 1]WLAN_INTERFACE_INFO
}

type WLAN_PROFILE_INFO struct {
	ProfileName uint8
	Flags       uint32
}

type WLAN_PROFILE_INFO_LIST struct {
	NumberOfItems uint32
	Index         uint32
	ProfileInfo   WLAN_PROFILE_INFO
}

type WLAN_RAW_DATA struct {
	dwDataSize uint32
	DataBlob   [257]byte
}

type WLAN_CONNECTION_PARAMETERS struct {
	wlanConnectionMode uint32
	strProfile
}
