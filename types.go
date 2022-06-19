package wlanapi

import (
	"golang.org/x/sys/windows"
	"syscall"
)

const (
	MAX_INDEX = 1000
	S_OK      = 0
)

//A DOT11_SSID structure contains the SSID of an interface.
//https://docs.microsoft.com/en-us/windows/win32/nativewifi/dot11-ssid
type DOT11_SSID struct {
	uSSIDLength ULONG
	ucSSID      [32]byte
}

type DOT11_MAC_ADDRESS [6]UCHAR

//The DOT11_BSSID_LIST structure contains a list of basic service set (BSS) identifiers.
//https://docs.microsoft.com/en-us/windows/win32/nativewifi/dot11-bssid-list
type DOT11_BSSID_LIST struct {
	Header             NDIS_OBJECT_HEADER
	uNumOfEntries      ULONG
	uTotalNumOfEntries ULONG
	BSSIDs             [1]DOT11_MAC_ADDRESS
}

//The DOT11_NETWORK structure contains information about an available wireless network.
//https://docs.microsoft.com/en-us/windows/win32/api/wlanapi/ns-wlanapi-dot11_network
type DOT11_NETWORK struct {
	dot11Ssid    DOT11_SSID
	dot11BssType DOT11_BSS_TYPE
}

//The DOT11_NETWORK_LIST structure contains a list of 802.11 wireless networks.
//https://docs.microsoft.com/en-us/windows/win32/api/wlanapi/ns-wlanapi-dot11_network_list
type DOT11_NETWORK_LIST struct {
	dwNumberOfItems DWORD
	dwIndex         DWORD
	Network         [MAX_INDEX + 1]DOT11_NETWORK
}

type NDIS_OBJECT_HEADER struct {
	Type     UCHAR
	Revision UCHAR
	Size     USHORT
}

//The WLAN_AVAILABLE_NETWORK_LIST structure contains an array of information about available networks.
//https://docs.microsoft.com/en-us/windows/win32/api/wlanapi/ns-wlanapi-wlan_available_network_list
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

func (avn WLAN_AVAILABLE_NETWORK) GetStrProfileName() (profileName string) {
	return windows.UTF16ToString(avn.strProfileName[:])
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
	InterfaceGuid           windows.GUID
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

//The WLAN_CONNECTION_PARAMETERS structure specifies the parameters used when using the WlanConnect function.
//https://docs.microsoft.com/en-us/windows/win32/api/wlanapi/ns-wlanapi-wlan_connection_parameters
type WLAN_CONNECTION_PARAMETERS struct {
	wlanConnectionMode WLAN_CONNECTION_MODE
	strProfile         string
	pDot11Ssid         DOT11_SSID
	pDesiredBssidList  DOT11_BSSID_LIST
	dot11BssType       DOT11_BSS_TYPE
	dwFlags            DWORD
}

//The WLAN_INTERFACE_CAPABILITY structure contains information about the capabilities of an interface.
//https://docs.microsoft.com/en-us/windows/win32/api/wlanapi/ns-wlanapi-wlan_interface_capability
type WLAN_INTERFACE_CAPABILITY struct {
	interfaceType             WLAN_INTERFACE_TYPE
	bDot11DSupported          BOOL
	dwMaxDesiredSsidListSize  DWORD
	dwMaxDesiredBssidListSize DWORD
	dwNumberOfSupportedPhys   DWORD
	dot11PhyTypes             [MAX_INDEX + 1]DOT11_PHY_TYPE
}

//WLAN_DEVICE_SERVICE_GUID_LIST Contains an array of device service GUIDs.
//https://docs.microsoft.com/zh-cn/windows/win32/api/wlanapi/ns-wlanapi-wlan_device_service_guid_list
type WLAN_DEVICE_SERVICE_GUID_LIST struct {
	dwNumberOfItems DWORD
	dwIndex         DWORD
	DeviceService   [MAX_INDEX + 1]syscall.GUID
}

//The WLAN_HOSTED_NETWORK_PEER_STATE structure contains information about the peer state for a peer on the wireless Hosted Network.
//https://docs.microsoft.com/en-us/windows/win32/api/wlanapi/ns-wlanapi-wlan_hosted_network_peer_state
type WLAN_HOSTED_NETWORK_PEER_STATE struct {
	PeerMacAddress DOT11_MAC_ADDRESS
	PeerAuthState  WLAN_HOSTED_NETWORK_PEER_AUTH_STATE
}

//The WLAN_HOSTED_NETWORK_STATUS structure contains information about the status of the wireless Hosted Network.
//https://docs.microsoft.com/en-us/windows/win32/api/wlanapi/ns-wlanapi-wlan_hosted_network_status
type WLAN_HOSTED_NETWORK_STATUS struct {
	HostedNetworkState     WLAN_HOSTED_NETWORK_STATE
	IPDeviceID             syscall.GUID
	wlanHostedNetworkBSSID DOT11_MAC_ADDRESS
	dot11PhyType           DOT11_PHY_TYPE
	ulChannelFrequency     ULONG
	dwNumberOfPeers        DWORD
	PeerList               [MAX_INDEX + 1]WLAN_HOSTED_NETWORK_PEER_STATE
}

//The EAP_TYPE structure contains type and vendor identification information for an EAP method.
//https://docs.microsoft.com/en-us/windows/win32/api/eaptypes/ns-eaptypes-eap_type
type EAP_TYPE struct {
	Type         BYTE
	dwVendorId   DWORD
	dwVendorType DWORD
}

//The EAP_METHOD_TYPE structure contains type, identification, and author information about an EAP method.
//https://docs.microsoft.com/en-us/windows/win32/api/eaptypes/ns-eaptypes-eap_method_type
type EAP_METHOD_TYPE struct {
	eapType    EAP_TYPE
	dwAuthorId DWORD
}

//The WLAN_RAW_DATA_LIST structure contains raw data in the form of an array of data blobs that are used by some Native Wifi functions.
//https://docs.microsoft.com/en-us/windows/win32/api/wlanapi/ns-wlanapi-wlan_raw_data_list
//TODO:
type WLAN_RAW_DATA_LIST struct {
	dwTotalSize     DWORD
	dwNumberOfItems DWORD
}

type WLAN_REASON_CODE uint32
