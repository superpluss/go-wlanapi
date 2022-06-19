package wlanapi

import (
	"log"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	pReserved       = uintptr(0)
	dwClientVersion = uintptr(2)
)

func WlanFreeMemory(pMemory PVOID) (err error) {
	r1, _, _ := procWlanOpenHandle.Call(
		uintptr(pMemory),
	)
	if r1 != S_OK {
		err = syscall.Errno(r1)
	}
	return
}

//The WlanOpenHandle function opens a connection to the server.
//https://docs.microsoft.com/en-us/windows/win32/api/wlanapi/nf-wlanapi-wlanopenhandle
func WlanOpenHandle() (handle windows.Handle, err error) {
	r1, _, _ := procWlanOpenHandle.Call(
		dwClientVersion,
		pReserved,
		uintptr(unsafe.Pointer(&dwClientVersion)),
		uintptr(unsafe.Pointer(&handle)),
	)
	if r1 != S_OK {
		err = syscall.Errno(r1)
	}
	return
}

//The WlanCloseHandle function closes a connection to the server.
//https://docs.microsoft.com/en-us/windows/win32/api/wlanapi/nf-wlanapi-wlanclosehandle
func WlanCloseHandle(handle windows.Handle) (err error) {
	r1, _, _ := procWlanCloseHandle.Call(
		uintptr(handle),
		pReserved,
	)
	if r1 != S_OK {
		return syscall.Errno(r1)
	}
	return
}

func WlanEnumInterfaces(handle windows.Handle) (interfaceInfoList *WLAN_INTERFACE_INFO_LIST, err error) {
	var iil *WLAN_INTERFACE_INFO_LIST
	r1, _, err := procWlanEnumInterfaces.Call(
		uintptr(handle),
		pReserved,
		uintptr(unsafe.Pointer(&iil)),
	)
	log.Println(err)
	if r1 != S_OK {
		err = syscall.Errno(r1)
	}
	return iil, nil

	//TODO: Need to rework to cleanup unmanaged memory WlanFreeMemory(uintptr(unsafe.Pointer(iil)))
}

//The WlanScan function requests a scan for available networks on the indicated interface.
//https://docs.microsoft.com/en-us/windows/win32/api/wlanapi/nf-wlanapi-wlanscan
/**
Sample Code: http://www.pinvoke.net/default.aspx/wlanapi.WlanScan
for (int i = 0; i < infoList.dwNumberOfItems; i++)
{
g = infoList.InterfaceInfo[i].InterfaceGuid;
uint resultCode = WlanScan(wlanHndl, ref g, IntPtr.Zero, IntPtr.Zero, IntPtr.Zero);
if (resultCode != 0)
    return;
}
*/
func WlanScan(hClientHandle windows.Handle, pInterfaceGuid *windows.GUID, pDot11Ssid *DOT11_SSID, pIeData *WLAN_RAW_DATA) (err error) {
	r1, _, _ := procWlanScan.Call(
		uintptr(hClientHandle),
		uintptr(unsafe.Pointer(pInterfaceGuid)),
		uintptr(unsafe.Pointer(pDot11Ssid)),
		uintptr(unsafe.Pointer(pIeData)),
		pReserved)
	if r1 != S_OK {
		return syscall.Errno(r1)
	}
	return
}

//The WlanConnect function attempts to connect to a specific network.
//https://docs.microsoft.com/en-us/windows/win32/api/wlanapi/nf-wlanapi-wlanconnect
/**
Sample Code: http://www.pinvoke.net/default.aspx/wlanapi.WlanConnect
WLAN_CONNECTION_PARAMETERS wlanConnectionParameters = new WLAN_CONNECTION_PARAMETERS();
wlanConnectionParameters.dot11BssType = DOT11_BSS_TYPE.dot11_BSS_type_any;
wlanConnectionParameters.dwFlags = 0;
wlanConnectionParameters.strProfile = "dlink";
wlanConnectionParameters.wlanConnectionMode = WLAN_CONNECTION_MODE.wlan_connection_mode_profile;
WlanConnect(ClientHandle,ref pInterfaceGuid,ref wlanConnectionParameters ,new IntPtr());
*/
func WlanConnect(handle windows.Handle, wlanConnectionParameters *WLAN_CONNECTION_PARAMETERS, guid *windows.GUID) error {
	r1, _, _ := wlanConnect.Call(
		uintptr(handle),
		uintptr(unsafe.Pointer(guid)),
		uintptr(unsafe.Pointer(wlanConnectionParameters)),
		pReserved,
	)
	return syscall.Errno(r1)
}

//The WlanDisconnect function disconnects an interface from its current network.
//https://docs.microsoft.com/en-us/windows/win32/api/wlanapi/nf-wlanapi-wlandisconnect
func WlanDisconnect(handle windows.Handle, pInterfaceGuid *windows.GUID) error {
	r1, _, _ := wlanDisconnect.Call(
		uintptr(handle),
		uintptr(unsafe.Pointer(pInterfaceGuid)),
		pReserved,
	)
	return syscall.Errno(r1)
}

//The WlanDeleteProfile function deletes a wireless profile for a wireless interface on the local computer.
//https://docs.microsoft.com/en-us/windows/win32/api/wlanapi/nf-wlanapi-wlandeleteprofile
func WlanDeleteProfile(handle windows.Handle, pInterfaceGuid *windows.GUID, strProfileName string) error {
	pProfileName, err := syscall.UTF16PtrFromString(strProfileName)
	if err != nil {
		log.Println(err)
		return err
	}
	r1, _, _ := wlanDeleteProfile.Call(
		uintptr(handle),
		uintptr(unsafe.Pointer(pInterfaceGuid)),
		uintptr(unsafe.Pointer(pProfileName)),
		pReserved,
	)
	return syscall.Errno(r1)
}

//The WlanGetAvailableNetworkList function retrieves the list of available networks on a wireless LAN interface.
//https://docs.microsoft.com/en-us/windows/win32/api/wlanapi/nf-wlanapi-wlangetavailablenetworklist
/**
Sample Code: http://www.pinvoke.net/default.aspx/wlanapi.WlanGetAvailableNetworkList
IntPtr ppAvailableNetworkList = new IntPtr();
Guid pInterfaceGuid = ((WLAN_INTERFACE_INFO)wlanInterfaceInfoList.InterfaceInfo[0]).InterfaceGuid;
WlanGetAvailableNetworkList(ClientHandle, ref pInterfaceGuid, WLAN_AVAILABLE_NETWORK_INCLUDE_ALL_MANUAL_HIDDEN_PROFILES, new  IntPtr(), ref  ppAvailableNetworkList);
WLAN_AVAILABLE_NETWORK_LIST wlanAvailableNetworkList = new WLAN_AVAILABLE_NETWORK_LIST(ppAvailableNetworkList);
WlanFreeMemory(ppAvailableNetworkList);
    for (int j = 0; j < wlanAvailableNetworkList .dwNumberOfItems; j++)
    {
    Interop.WLAN_AVAILABLE_NETWORK network = wlanAvailableNetworkList.wlanAvailableNetwork[j];
    Console.WriteLine("Available Network: ");
    Console.WriteLine("SSID: " + network.dot11Ssid.ucSSID);
    Console.WriteLine("Encrypted: " + network.bSecurityEnabled);
    Console.WriteLine("Signal Strength: " + network.wlanSignalQuality);
    Console.WriteLine("Default Authentication: " +
        network.dot11DefaultAuthAlgorithm.ToString());
    Console.WriteLine("Default Cipher: " + network.dot11DefaultCipherAlgorithm.ToString());
    Console.WriteLine();
    }
*/
func WlanGetAvailableNetworkList(handle windows.Handle, pInterfaceGuid *windows.GUID, dwFlags DWORD) (
	ppAvailableNetworkList *WLAN_AVAILABLE_NETWORK_LIST, err error) {

	r1, _, _ := procWlanGetAvailableNetworkList.Call(
		uintptr(handle),
		uintptr(unsafe.Pointer(pInterfaceGuid)),
		uintptr(dwFlags),
		pReserved,
		uintptr(unsafe.Pointer(&ppAvailableNetworkList)),
	)
	if r1 != S_OK {
		err = syscall.Errno(r1)
	}
	return
}

//The WlanGetFilterList function retrieves a group policy or user permission list.
//https://docs.microsoft.com/en-us/windows/win32/api/wlanapi/nf-wlanapi-wlangetfilterlist
func WlanGetFilterList(handle windows.Handle, wlanFilterListType WLAN_FILTER_LIST_TYPE) (ppNetworkList *DOT11_NETWORK_LIST, err error) {
	r1, _, _ := wlanGetFilterList.Call(
		uintptr(handle),
		uintptr(wlanFilterListType),
		pReserved,
		uintptr(unsafe.Pointer(ppNetworkList)),
	)
	if r1 != S_OK {
		err = syscall.Errno(r1)
	}
	return
}

//The WlanGetInterfaceCapability function retrieves the capabilities of an interface.
//https://docs.microsoft.com/en-us/windows/win32/api/wlanapi/nf-wlanapi-wlangetinterfacecapability
func WlanGetInterfaceCapability(handle windows.Handle, pInterfaceGuid *windows.GUID) (ppCapability *WLAN_INTERFACE_CAPABILITY, err error) {
	r1, _, _ := wlanGetInterfaceCapability.Call(
		uintptr(handle),
		uintptr(unsafe.Pointer(pInterfaceGuid)),
		pReserved,
		uintptr(unsafe.Pointer(ppCapability)),
	)
	err = syscall.Errno(r1)
	return
}

//The WlanGetNetworkBssList function retrieves a list of the basic service set (BSS) entries of the wireless network or networks on a given wireless LAN interface.
//https://docs.microsoft.com/zh-cn/windows/win32/api/wlanapi/nf-wlanapi-wlangetnetworkbsslist
func WlanGetNetworkBssList(handle windows.Handle,
	pInterfaceGuid *windows.GUID,
	pDot11Ssid *DOT11_SSID,
	dot11BssType DOT11_BSS_TYPE,
	bSecurityEnabled BOOL) (ppWlanBssList *WLAN_BSS_LIST, err error) {
	r1, _, _ := wlanGetNetworkBssList.Call(
		uintptr(handle),
		uintptr(unsafe.Pointer(pInterfaceGuid)),
		uintptr(unsafe.Pointer(pDot11Ssid)),
		uintptr(dot11BssType),
		uintptr(bSecurityEnabled),
		pReserved,
		// out
		uintptr(unsafe.Pointer(ppWlanBssList)),
	)
	err = syscall.Errno(r1)
	return
}

//The WlanGetProfile function retrieves all information about a specified wireless profile.
//https://docs.microsoft.com/zh-cn/windows/win32/api/wlanapi/nf-wlanapi-wlangetprofile
func WlanGetProfile(
	handle windows.Handle,
	pInterfaceGuid *windows.GUID,
	strProfileName string) (pstrProfileXml string, pdwFlags *DWORD, pdwGrantedAccess *DWORD, err error) {
	pProfileName, err := syscall.UTF16PtrFromString(strProfileName)
	if err != nil {
		log.Println(err)
		return
	}
	var pstrProfile *uint16
	r1, _, _ := wlanGetProfile.Call(
		uintptr(handle),
		uintptr(unsafe.Pointer(pInterfaceGuid)),
		uintptr(unsafe.Pointer(pProfileName)),
		pReserved,
		// out
		uintptr(unsafe.Pointer(pstrProfile)),
		uintptr(unsafe.Pointer(pdwFlags)),
		uintptr(unsafe.Pointer(pdwGrantedAccess)),
	)
	err = syscall.Errno(r1)
	//syscall.UTF16ToString(pstrProfile)
	return
}

//The WlanGetProfileCustomUserData function gets the custom user data associated with a wireless profile.
//https://docs.microsoft.com/zh-cn/windows/win32/api/wlanapi/nf-wlanapi-wlangetprofilecustomuserdata
func WlanGetProfileCustomUserData(
	handle windows.Handle,
	pInterfaceGuid *windows.GUID,
	strProfileName string) (pdwDataSize *DWORD, ppData *BYTE, err error) {
	pProfileName, err := syscall.UTF16PtrFromString(strProfileName)
	if err != nil {
		log.Println(err)
		return
	}
	r1, _, _ := wlanGetProfileCustomUserData.Call(
		uintptr(handle),
		uintptr(unsafe.Pointer(pInterfaceGuid)),
		uintptr(unsafe.Pointer(pProfileName)),
		pReserved,
		// out
		uintptr(unsafe.Pointer(pdwDataSize)),
		uintptr(unsafe.Pointer(ppData)),
	)
	err = syscall.Errno(r1)
	return
}

//The WlanGetProfileList function retrieves the list of profiles in preference order.
//https://docs.microsoft.com/zh-cn/windows/win32/api/wlanapi/nf-wlanapi-wlangetprofilelist
func WlanGetProfileList(handle windows.Handle, pInterfaceGuid *windows.GUID) (ppProfileList *WLAN_PROFILE_INFO_LIST, err error) {
	r1, _, _ := wlanGetProfileList.Call(
		uintptr(handle),
		uintptr(unsafe.Pointer(pInterfaceGuid)),
		pReserved,
		// out
		uintptr(unsafe.Pointer(ppProfileList)),
	)
	err = syscall.Errno(r1)
	return
}

//The WlanGetSecuritySettings function gets the security settings associated with a configurable object.
//https://docs.microsoft.com/zh-cn/windows/win32/api/wlanapi/nf-wlanapi-wlangetsecuritysettings
func WlanGetSecuritySettings(handle windows.Handle, SecurableObject WLAN_SECURABLE_OBJECT) (
	pValueType *WLAN_OPCODE_VALUE_TYPE, pdwGrantedAccess *DWORD, err error) {

	var pstrCurrentSDDL *uint16
	r1, _, _ := wlanGetSecuritySettings.Call(
		uintptr(handle),
		uintptr(SecurableObject),
		// out
		uintptr(unsafe.Pointer(&pValueType)),
		uintptr(unsafe.Pointer(&pstrCurrentSDDL)),
		uintptr(unsafe.Pointer(&pdwGrantedAccess)),
	)
	err = syscall.Errno(r1)
	return
}

//WlanGetSupportedDeviceServices Retrieves a list of the supported device services on a given wireless LAN interface.
//https://docs.microsoft.com/zh-cn/windows/win32/api/wlanapi/nf-wlanapi-wlangetsupporteddeviceservices
func WlanGetSupportedDeviceServices(handle windows.Handle, pInterfaceGuid *windows.GUID) (
	ppDevSvcGuidList *WLAN_DEVICE_SERVICE_GUID_LIST, err error) {

	r1, _, _ := wlanGetSupportedDeviceServices.Call(
		uintptr(handle),
		uintptr(unsafe.Pointer(pInterfaceGuid)),
		// out
		uintptr(unsafe.Pointer(&ppDevSvcGuidList)),
	)
	err = syscall.Errno(r1)
	return
}

//The WlanHostedNetworkForceStart function transitions the wireless Hosted Network to the wlan_hosted_network_active state without associating the request with the application's calling handle.
//https://docs.microsoft.com/zh-cn/windows/win32/api/wlanapi/nf-wlanapi-wlanhostednetworkforcestart
func WlanHostedNetworkForceStart(handle windows.Handle) (pFailReason *WLAN_HOSTED_NETWORK_REASON, err error) {
	r1, _, _ := wlanHostedNetworkForceStart.Call(
		uintptr(handle),
		// out
		uintptr(unsafe.Pointer(&pFailReason)),
		pReserved,
	)
	err = syscall.Errno(r1)
	return
}

//The WlanHostedNetworkForceStop function transitions the wireless Hosted Network to the wlan_hosted_network_idle without associating the request with the application's calling handle.
//https://docs.microsoft.com/zh-cn/windows/win32/api/wlanapi/nf-wlanapi-wlanhostednetworkforcestop
func WlanHostedNetworkForceStop(handle windows.Handle) (pFailReason *WLAN_HOSTED_NETWORK_REASON, err error) {
	r1, _, _ := wlanHostedNetworkForceStop.Call(
		uintptr(handle),
		// out
		uintptr(unsafe.Pointer(&pFailReason)),
		pReserved,
	)
	err = syscall.Errno(r1)
	return
}

//The WlanHostedNetworkInitSettings function configures and persists to storage the network connection settings (SSID and maximum number of peers, for example) on the wireless Hosted Network if these settings are not already configured.
//https://docs.microsoft.com/zh-cn/windows/win32/api/wlanapi/nf-wlanapi-wlanhostednetworkinitsettings
func WlanHostedNetworkInitSettings(handle windows.Handle) (pFailReason *WLAN_HOSTED_NETWORK_REASON, err error) {
	r1, _, _ := wlanHostedNetworkInitSettings.Call(
		uintptr(handle),
		// out
		uintptr(unsafe.Pointer(&pFailReason)),
		pReserved,
	)
	err = syscall.Errno(r1)
	return
}

//The WlanHostedNetworkQueryProperty function queries the current static properties of the wireless Hosted Network.
//https://docs.microsoft.com/zh-cn/windows/win32/api/wlanapi/nf-wlanapi-wlanhostednetworkqueryproperty
func WlanHostedNetworkQueryProperty(handle windows.Handle, OpCode WLAN_HOSTED_NETWORK_OPCODE) (
	pdwDataSize *DWORD, ppvData *PVOID, pWlanOpcodeValueType *WLAN_OPCODE_VALUE_TYPE, err error) {

	r1, _, _ := wlanHostedNetworkQueryProperty.Call(
		uintptr(handle),
		uintptr(OpCode),
		// out
		uintptr(unsafe.Pointer(&pdwDataSize)),
		uintptr(unsafe.Pointer(&ppvData)),
		uintptr(unsafe.Pointer(&pWlanOpcodeValueType)),
		pReserved,
	)
	err = syscall.Errno(r1)
	return
}

//The WlanHostedNetworkQuerySecondaryKey function queries the secondary security key that is configured to be used by the wireless Hosted Network.
//https://docs.microsoft.com/zh-cn/windows/win32/api/wlanapi/nf-wlanapi-wlanhostednetworkquerysecondarykey
func WlanHostedNetworkQuerySecondaryKey(handle windows.Handle) (
	pdwKeyLength *DWORD, ppucKeyData *UCHAR, pbIsPassPhrase *BOOL, pbPersistent *BOOL, pFailReason WLAN_HOSTED_NETWORK_REASON, err error) {

	r1, _, _ := wlanHostedNetworkQuerySecondaryKey.Call(
		uintptr(handle),
		// out
		uintptr(unsafe.Pointer(pdwKeyLength)),
		uintptr(unsafe.Pointer(&ppucKeyData)),
		uintptr(unsafe.Pointer(&pbIsPassPhrase)),
		uintptr(unsafe.Pointer(&pbPersistent)),
		uintptr(unsafe.Pointer(&pFailReason)),
		pReserved,
	)
	err = syscall.Errno(r1)
	return
}

//The WlanHostedNetworkQueryStatus function queries the current status of the wireless Hosted Network.
//https://docs.microsoft.com/zh-cn/windows/win32/api/wlanapi/nf-wlanapi-wlanhostednetworkquerystatus
func WlanHostedNetworkQueryStatus(handle windows.Handle) (ppWlanHostedNetworkStatus *WLAN_HOSTED_NETWORK_STATUS, err error) {
	r1, _, _ := wlanHostedNetworkQueryStatus.Call(
		uintptr(handle),
		// out
		uintptr(unsafe.Pointer(&ppWlanHostedNetworkStatus)),
		pReserved,
	)
	err = syscall.Errno(r1)
	return
}

//The WlanHostedNetworkRefreshSecuritySettings function refreshes the configurable and auto-generated parts of the wireless Hosted Network security settings.
//https://docs.microsoft.com/zh-cn/windows/win32/api/wlanapi/nf-wlanapi-wlanhostednetworkrefreshsecuritysettings
func WlanHostedNetworkRefreshSecuritySettings(handle windows.Handle) (pFailReason *WLAN_HOSTED_NETWORK_REASON, err error) {
	r1, _, _ := wlanHostedNetworkRefreshSecuritySettings.Call(
		uintptr(handle),
		// out
		uintptr(unsafe.Pointer(&pFailReason)),
		pReserved,
	)
	err = syscall.Errno(r1)
	return
}

//The WlanHostedNetworkSetProperty function sets static properties of the wireless Hosted Network.
//https://docs.microsoft.com/zh-cn/windows/win32/api/wlanapi/nf-wlanapi-wlanhostednetworksetproperty
func WlanHostedNetworkSetProperty(
	handle windows.Handle,
	OpCode WLAN_HOSTED_NETWORK_OPCODE,
	dwDataSize DWORD,
	pvData PVOID) (pFailReason *WLAN_HOSTED_NETWORK_REASON, err error) {

	r1, _, _ := wlanHostedNetworkSetProperty.Call(
		uintptr(handle),
		uintptr(OpCode),
		uintptr(dwDataSize),
		uintptr(pvData),
		// out
		uintptr(unsafe.Pointer(&pFailReason)),
		pReserved,
	)
	err = syscall.Errno(r1)
	return
}

//The WlanHostedNetworkSetSecondaryKey function configures the secondary security key that will be used by the wireless Hosted Network.
//https://docs.microsoft.com/zh-cn/windows/win32/api/wlanapi/nf-wlanapi-wlanhostednetworksetsecondarykey
func WlanHostedNetworkSetSecondaryKey(
	handle windows.Handle,
	dwKeyLength DWORD,
	pucKeyData *UCHAR,
	bIsPassPhrase, bPersistent BOOL) (pFailReason WLAN_HOSTED_NETWORK_REASON, err error) {

	r1, _, _ := wlanHostedNetworkSetSecondaryKey.Call(
		uintptr(handle),
		uintptr(dwKeyLength),
		uintptr(unsafe.Pointer(pucKeyData)),
		uintptr(bIsPassPhrase),
		uintptr(bPersistent),
		// out
		uintptr(unsafe.Pointer(&pFailReason)),
		pReserved,
	)
	err = syscall.Errno(r1)
	return
}

//The WlanHostedNetworkStartUsing function starts the wireless Hosted Network.
//https://docs.microsoft.com/zh-cn/windows/win32/api/wlanapi/nf-wlanapi-wlanhostednetworkstartusing
func WlanHostedNetworkStartUsing(handle windows.Handle) (pFailReason WLAN_HOSTED_NETWORK_REASON, err error) {

	r1, _, _ := wlanHostedNetworkStartUsing.Call(
		uintptr(handle),
		// out
		uintptr(unsafe.Pointer(&pFailReason)),
		pReserved,
	)
	err = syscall.Errno(r1)
	return
}

//The WlanHostedNetworkStopUsing function stops the wireless Hosted Network.
//https://docs.microsoft.com/zh-cn/windows/win32/api/wlanapi/nf-wlanapi-wlanhostednetworkstopusing
func WlanHostedNetworkStopUsing(handle windows.Handle) (pFailReason WLAN_HOSTED_NETWORK_REASON, err error) {

	r1, _, _ := wlanHostedNetworkStopUsing.Call(
		uintptr(handle),
		// out
		uintptr(unsafe.Pointer(&pFailReason)),
		pReserved,
	)
	err = syscall.Errno(r1)
	return
}

//The WlanIhvControl function provides a mechanism for independent hardware vendor (IHV) control of WLAN drivers or services.
//https://docs.microsoft.com/zh-cn/windows/win32/api/wlanapi/nf-wlanapi-wlanihvcontrol
func WlanIhvControl(
	handle windows.Handle,
	pInterfaceGuid *windows.GUID,
	Type WLAN_IHV_CONTROL_TYPE,
	dwInBufferSize DWORD,
	pInBuffer PVOID,
	dwOutBufferSize DWORD) (pOutBuffer *PVOID, pdwBytesReturned *DWORD, err error) {

	r1, _, _ := wlanIhvControl.Call(
		uintptr(handle),
		uintptr(unsafe.Pointer(pInterfaceGuid)),
		uintptr(Type),
		uintptr(dwInBufferSize),
		uintptr(pInBuffer),
		uintptr(dwOutBufferSize),
		// out
		uintptr(unsafe.Pointer(&pOutBuffer)),
		uintptr(unsafe.Pointer(&pdwBytesReturned)),
	)
	err = syscall.Errno(r1)
	return
}

//The WlanQueryAutoConfigParameter function queries for the parameters of the auto configuration service.
//https://docs.microsoft.com/zh-cn/windows/win32/api/wlanapi/nf-wlanapi-wlanqueryautoconfigparameter
func WlanQueryAutoConfigParameter(handle windows.Handle, OpCode WLAN_AUTOCONF_OPCODE) (
	pdwDataSize *DWORD, ppData *PVOID, pWlanOpcodeValueType *WLAN_OPCODE_VALUE_TYPE, err error) {

	r1, _, _ := wlanIhvControl.Call(
		uintptr(handle),
		uintptr(OpCode),
		pReserved,
		// out
		uintptr(unsafe.Pointer(&pdwDataSize)),
		uintptr(unsafe.Pointer(&pWlanOpcodeValueType)),
		uintptr(unsafe.Pointer(&pWlanOpcodeValueType)),
	)
	err = syscall.Errno(r1)
	return
}

//The WlanQueryInterface function queries various parameters of a specified interface.
//https://docs.microsoft.com/zh-cn/windows/win32/api/wlanapi/nf-wlanapi-wlanqueryinterface
func WlanQueryInterface(handle windows.Handle, pInterfaceGuid *windows.GUID, OpCode WLAN_INTF_OPCODE) (
	pdwDataSize *DWORD, ppData *PVOID, pWlanOpcodeValueType *WLAN_OPCODE_VALUE_TYPE, err error) {

	r1, _, _ := procWlanQueryInterface.Call(
		uintptr(handle),
		uintptr(unsafe.Pointer(pInterfaceGuid)),
		uintptr(OpCode),
		pReserved,
		// out
		uintptr(unsafe.Pointer(pdwDataSize)),
		uintptr(unsafe.Pointer(&ppData)),
		uintptr(unsafe.Pointer(pWlanOpcodeValueType)),
	)
	if r1 != S_OK {
		err = syscall.Errno(r1)
	}
	return
}

//The WlanReasonCodeToString function retrieves a string that describes a specified reason code.
//https://docs.microsoft.com/zh-cn/windows/win32/api/wlanapi/nf-wlanapi-wlanreasoncodetostring
func WlanReasonCodeToString(dwReasonCode DWORD, dwBufferSize DWORD, pStringBuffer *WCHAR) (err error) {
	r1, _, _ := wlanIhvControl.Call(
		uintptr(dwReasonCode),
		uintptr(dwBufferSize),
		uintptr(unsafe.Pointer(&pStringBuffer)),
		pReserved,
	)
	err = syscall.Errno(r1)
	return
}

//WlanRegisterDeviceServiceNotification Allows user mode clients with admin privileges, or User-Mode Driver Framework (UMDF) drivers, to register for unsolicited notifications corresponding to device services that they're interested in.
//https://docs.microsoft.com/zh-cn/windows/win32/api/wlanapi/nf-wlanapi-wlanregisterdeviceservicenotification
func WlanRegisterDeviceServiceNotification(handle windows.Handle, pDevSvcGuidList WLAN_DEVICE_SERVICE_GUID_LIST) (err error) {
	r1, _, _ := wlanIhvControl.Call(
		uintptr(handle),
		uintptr(unsafe.Pointer(&pDevSvcGuidList)),
	)
	err = syscall.Errno(r1)
	return
}

//The WlanRegisterNotification function is used to register and unregister notifications on all wireless interfaces.
//https://docs.microsoft.com/zh-cn/windows/win32/api/wlanapi/nf-wlanapi-wlanregisternotification
// TODO:
func WlanRegisterNotification(handle windows.Handle, dwNotifSource DWORD, bIgnoreDuplicate BOOL) (err error) {

	return
}

//The WlanRegisterVirtualStationNotification function is used to register and unregister notifications on a virtual station.
//https://docs.microsoft.com/zh-cn/windows/win32/api/wlanapi/nf-wlanapi-wlanregistervirtualstationnotification
func WlanRegisterVirtualStationNotification(handle windows.Handle, bRegister BOOL) (err error) {
	r1, _, _ := wlanIhvControl.Call(
		uintptr(handle),
		uintptr(bRegister),
	)
	err = syscall.Errno(r1)
	return
}

//The WlanRenameProfile function renames the specified profile.
//https://docs.microsoft.com/zh-cn/windows/win32/api/wlanapi/nf-wlanapi-wlanrenameprofile
func WlanRenameProfile(handle windows.Handle, pInterfaceGuid *windows.GUID, strOldProfileName, strNewProfileName string) (err error) {
	oldProfileName, err := syscall.UTF16FromString(strOldProfileName)
	if err != nil {
		log.Println(err)
		return
	}
	newProfileName, err := syscall.UTF16FromString(strNewProfileName)
	if err != nil {
		log.Println(err)
		return
	}
	r1, _, _ := wlanIhvControl.Call(
		uintptr(handle),
		uintptr(unsafe.Pointer(pInterfaceGuid)),
		uintptr(unsafe.Pointer(&oldProfileName)),
		uintptr(unsafe.Pointer(&newProfileName)),
	)
	err = syscall.Errno(r1)
	return
}

//The WlanSaveTemporaryProfile function saves a temporary profile to the profile store.
//https://docs.microsoft.com/zh-cn/windows/win32/api/wlanapi/nf-wlanapi-wlansavetemporaryprofile
func WlanSaveTemporaryProfile(handle windows.Handle,
	pInterfaceGuid *windows.GUID,
	strProfileName, strAllUserProfileSecurity string,
	dwFlags DWORD,
	bOverWrite BOOL) (err error) {
	profileName, err := syscall.UTF16FromString(strProfileName)
	if err != nil {
		log.Println(err)
		return
	}
	allUserProfileSecurity, err := syscall.UTF16FromString(strAllUserProfileSecurity)
	if err != nil {
		log.Println(err)
		return
	}
	r1, _, _ := wlanIhvControl.Call(
		uintptr(handle),
		uintptr(unsafe.Pointer(pInterfaceGuid)),
		uintptr(unsafe.Pointer(&profileName)),
		uintptr(unsafe.Pointer(&allUserProfileSecurity)),
		uintptr(dwFlags),
		uintptr(bOverWrite),
	)
	err = syscall.Errno(r1)
	return
}

//The WlanSetAutoConfigParameter function sets parameters for the automatic configuration service.
//https://docs.microsoft.com/zh-cn/windows/win32/api/wlanapi/nf-wlanapi-wlansetautoconfigparameter
func WlanSetAutoConfigParameter(handle windows.Handle, OpCode WLAN_AUTOCONF_OPCODE, dwDataSize DWORD, pData PVOID) (err error) {
	r1, _, _ := wlanIhvControl.Call(
		uintptr(handle),
		uintptr(OpCode),
		uintptr(dwDataSize),
		uintptr(pData),
		pReserved,
	)
	err = syscall.Errno(r1)
	return
}

//The WlanSetFilterList function sets the permit/deny list.
//https://docs.microsoft.com/zh-cn/windows/win32/api/wlanapi/nf-wlanapi-wlansetfilterlist
func WlanSetFilterList(handle windows.Handle, wlanFilterListType WLAN_FILTER_LIST_TYPE, pNetworkList *DOT11_NETWORK_LIST) (err error) {
	r1, _, _ := wlanIhvControl.Call(
		uintptr(handle),
		uintptr(wlanFilterListType),
		uintptr(unsafe.Pointer(pNetworkList)),
		pReserved,
	)
	err = syscall.Errno(r1)
	return
}

//The WlanSetInterface function sets user-configurable parameters for a specified interface.
//https://docs.microsoft.com/zh-cn/windows/win32/api/wlanapi/nf-wlanapi-wlansetinterface
func WlanSetInterface(handle windows.Handle,
	pInterfaceGuid *windows.GUID, OpCode WLAN_INTF_OPCODE, dwDataSize DWORD, pData PVOID) (err error) {
	r1, _, _ := wlanIhvControl.Call(
		uintptr(handle),
		uintptr(unsafe.Pointer(pInterfaceGuid)),
		uintptr(OpCode),
		uintptr(dwDataSize),
		uintptr(pData),
		pReserved,
	)
	err = syscall.Errno(r1)
	return
}

//The WlanSetProfile function sets the content of a specific profile.
//https://docs.microsoft.com/zh-cn/windows/win32/api/wlanapi/nf-wlanapi-wlansetprofile
func WlanSetProfile(handle windows.Handle,
	pInterfaceGuid *windows.GUID,
	dwFlags DWORD,
	strProfileXml, strAllUserProfileSecurity string,
	bOverwrite BOOL) (pdwReasonCode *DWORD, err error) {
	profileXml, err := syscall.UTF16FromString(strProfileXml)
	if err != nil {
		log.Println(err)
		return
	}
	allUserProfileSecurity, err := syscall.UTF16FromString(strAllUserProfileSecurity)
	if err != nil {
		log.Println(err)
		return
	}
	r1, _, _ := wlanIhvControl.Call(
		uintptr(handle),
		uintptr(unsafe.Pointer(pInterfaceGuid)),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(&profileXml)),
		uintptr(unsafe.Pointer(&allUserProfileSecurity)),
		uintptr(bOverwrite),
		pReserved,
		uintptr(unsafe.Pointer(&pdwReasonCode)),
	)
	err = syscall.Errno(r1)
	return
}

//The WlanSetProfileCustomUserData function sets the custom user data associated with a profile.
//https://docs.microsoft.com/zh-cn/windows/win32/api/wlanapi/nf-wlanapi-wlansetprofilecustomuserdata
func WlanSetProfileCustomUserData(
	handle windows.Handle,
	pInterfaceGuid *windows.GUID,
	strProfileName string,
	dwDataSize DWORD,
	pData *BYTE) (err error) {

	profileName, err := syscall.UTF16FromString(strProfileName)
	if err != nil {
		log.Println(err)
		return
	}
	r1, _, _ := wlanIhvControl.Call(
		uintptr(handle),
		uintptr(unsafe.Pointer(pInterfaceGuid)),
		uintptr(unsafe.Pointer(&profileName)),
		uintptr(dwDataSize),
		uintptr(unsafe.Pointer(&pData)),
		pReserved,
	)
	err = syscall.Errno(r1)
	return
}

//The WlanSetProfileEapUserData function sets the Extensible Authentication Protocol (EAP) user credentials as specified by raw EAP data. The user credentials apply to a profile on an interface.
//https://docs.microsoft.com/zh-cn/windows/win32/api/wlanapi/nf-wlanapi-wlansetprofileeapuserdata
func WlanSetProfileEapUserData(
	handle windows.Handle,
	pInterfaceGuid *windows.GUID,
	strProfileName string,
	eapType EAP_METHOD_TYPE,
	dwFlags DWORD,
	dwEapUserDataSize DWORD,
	pbEapUserData BYTE) (err error) {
	profileName, err := syscall.UTF16FromString(strProfileName)
	if err != nil {
		log.Println(err)
		return
	}
	r1, _, _ := wlanIhvControl.Call(
		uintptr(handle),
		uintptr(unsafe.Pointer(pInterfaceGuid)),
		uintptr(unsafe.Pointer(&profileName)),
		uintptr(unsafe.Pointer(&eapType)),
		uintptr(dwFlags),
		uintptr(dwEapUserDataSize),
		uintptr(pbEapUserData),
		pReserved,
	)
	err = syscall.Errno(r1)
	return
}

//The WlanSetProfileEapXmlUserData function sets the Extensible Authentication Protocol (EAP) user credentials as specified by an XML string. The user credentials apply to a profile on an adapter. These credentials can be used only by the caller.
//https://docs.microsoft.com/zh-cn/windows/win32/api/wlanapi/nf-wlanapi-wlansetprofileeapxmluserdata
func WlanSetProfileEapXmlUserData(
	handle windows.Handle,
	pInterfaceGuid *windows.GUID,
	strProfileName string,
	dwFlags DWORD,
	strEapXmlUserData string) (err error) {

	profileName, err := syscall.UTF16FromString(strProfileName)
	if err != nil {
		log.Println(err)
		return
	}

	eapXmlUserData, err := syscall.UTF16FromString(strEapXmlUserData)
	if err != nil {
		log.Println(err)
		return
	}
	r1, _, _ := wlanIhvControl.Call(
		uintptr(handle),
		uintptr(unsafe.Pointer(pInterfaceGuid)),
		uintptr(unsafe.Pointer(&profileName)),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(&eapXmlUserData)),
		pReserved,
	)
	err = syscall.Errno(r1)
	return
}

//The WlanSetProfileList function sets the preference order of profiles for a given interface.
//https://docs.microsoft.com/zh-cn/windows/win32/api/wlanapi/nf-wlanapi-wlansetprofilelist
func WlanSetProfileList(
	handle windows.Handle,
	pInterfaceGuid *windows.GUID,
	dwItems DWORD,
	strProfileNames string) (err error) {
	profileNames, err := syscall.UTF16FromString(strProfileNames)
	if err != nil {
		log.Println(err)
		return
	}
	r1, _, _ := wlanIhvControl.Call(
		uintptr(handle),
		uintptr(unsafe.Pointer(pInterfaceGuid)),
		uintptr(dwItems),
		uintptr(unsafe.Pointer(&profileNames)),
		pReserved,
	)
	err = syscall.Errno(r1)
	return
}

//The WlanSetProfilePosition function sets the position of a single, specified profile in the preference list.
//https://docs.microsoft.com/zh-cn/windows/win32/api/wlanapi/nf-wlanapi-wlansetprofileposition
func WlanSetProfilePosition(
	handle windows.Handle,
	pInterfaceGuid *windows.GUID,
	strProfileNames string, dwPosition DWORD) (err error) {
	profileNames, err := syscall.UTF16FromString(strProfileNames)
	if err != nil {
		log.Println(err)
		return
	}
	r1, _, _ := wlanIhvControl.Call(
		uintptr(handle),
		uintptr(unsafe.Pointer(pInterfaceGuid)),
		uintptr(unsafe.Pointer(&profileNames)),
		uintptr(dwPosition),
		pReserved,
	)
	err = syscall.Errno(r1)
	return
}

//The WlanSetPsdIEDataList function sets the proximity service discovery (PSD) information element (IE) data list.
//https://docs.microsoft.com/en-us/windows/win32/api/wlanapi/nf-wlanapi-wlansetpsdiedatalist
func WlanSetPsdIEDataList(handle windows.Handle, strFormat string, pPsdIEDataList *WLAN_RAW_DATA_LIST) (err error) {
	format, err := syscall.UTF16FromString(strFormat)
	if err != nil {
		log.Println(err)
		return
	}
	r1, _, _ := wlanIhvControl.Call(
		uintptr(handle),
		uintptr(unsafe.Pointer(&format)),
		uintptr(unsafe.Pointer(&pPsdIEDataList)),
		pReserved,
	)
	err = syscall.Errno(r1)
	return
}

//WlanSetSecuritySettings The WlanGetProfileList function sets the security settings for a configurable object.
//https://docs.microsoft.com/en-us/windows/win32/api/wlanapi/nf-wlanapi-wlansetsecuritysettings
func WlanSetSecuritySettings(handle windows.Handle, SecurableObject WLAN_SECURABLE_OBJECT, strModifiedSDDL string) (err error) {
	modifiedSDDL, err := syscall.UTF16FromString(strModifiedSDDL)
	if err != nil {
		log.Println(err)
		return
	}
	r1, _, _ := wlanIhvControl.Call(
		uintptr(handle),
		uintptr(SecurableObject),
		uintptr(unsafe.Pointer(&modifiedSDDL)),
		pReserved,
	)
	err = syscall.Errno(r1)
	return
}

//WlanUIEditProfile Displays the wireless profile user interface (UI). This UI is used to view and edit advanced settings of a wireless network profile.
//https://docs.microsoft.com/en-us/windows/win32/api/wlanapi/nf-wlanapi-wlanuieditprofile
func WlanUIEditProfile(
	dwClientVersion DWORD,
	wstrProfileName string,
	pInterfaceGuid *windows.GUID,
	hWnd HWND,
	wlStartPage WL_DISPLAY_PAGES) (pWlanReasonCode *WLAN_REASON_CODE, err error) {

	profileName, err := syscall.UTF16FromString(wstrProfileName)
	if err != nil {
		log.Println(err)
		return
	}
	r1, _, _ := wlanIhvControl.Call(
		uintptr(dwClientVersion),
		uintptr(unsafe.Pointer(&profileName)),
		uintptr(unsafe.Pointer(pInterfaceGuid)),
		uintptr(unsafe.Pointer(&hWnd)),
		uintptr(unsafe.Pointer(&wlStartPage)),
		pReserved,
		uintptr(unsafe.Pointer(&pWlanReasonCode)),
	)
	err = syscall.Errno(r1)
	return
}

func defaultInterface(handle windows.Handle) (wii WLAN_INTERFACE_INFO, err error) {
	iil, err := WlanEnumInterfaces(handle)
	log.Printf("dwIndex:%d dwNumberOfItems:%d", iil.dwIndex, iil.dwNumberOfItems)
	if err != nil {
		return
	}
	return iil.InterfaceInfo[iil.dwIndex], nil
}
