package wlanapi

import (
	"fmt"
	"log"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	pReserved       = uintptr(0)
	dwClientVersion = uintptr(2)
)

func retValToError(errNo uintptr) error {
	return fmt.Errorf("return code: %x", errNo)
}

//The WlanOpenHandle function opens a connection to the server.
//https://docs.microsoft.com/en-us/windows/win32/api/wlanapi/nf-wlanapi-wlanopenhandle
func WlanOpenHandle() (handle *windows.Handle, err error) {
	var h windows.Handle
	r1, _, _ := wlanOpenHandle.Call(
		dwClientVersion,
		pReserved,
		uintptr(unsafe.Pointer(&dwClientVersion)),
		uintptr(unsafe.Pointer(&h)),
	)
	if r1 != S_OK {
		return nil, retValToError(r1)
	}
	return &h, nil
}

//The WlanCloseHandle function closes a connection to the server.
//https://docs.microsoft.com/en-us/windows/win32/api/wlanapi/nf-wlanapi-wlanclosehandle
func WlanCloseHandle(handle *windows.Handle) (err error) {
	r1, _, _ := fWlanCloseHandle.Call(
		uintptr(unsafe.Pointer(handle)),
		pReserved,
	)
	if r1 != S_OK {
		return retValToError(r1)
	}
	return nil
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
func WlanScan(hClientHandle uintptr, pInterfaceGuid *syscall.GUID, pDot11Ssid *DOT11_SSID, pIeData *WLAN_RAW_DATA) error {
	r1, _, _ := wlanScan.Call(hClientHandle,
		uintptr(unsafe.Pointer(pInterfaceGuid)),
		uintptr(unsafe.Pointer(pDot11Ssid)),
		uintptr(unsafe.Pointer(pIeData)),
		pReserved)
	return syscall.Errno(r1)
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
func WlanConnect(handle *windows.Handle, wlanConnectionParameters *WLAN_CONNECTION_PARAMETERS, guid *syscall.GUID) error {
	r1, _, _ := wlanConnect.Call(
		uintptr(unsafe.Pointer(handle)),
		uintptr(unsafe.Pointer(guid)),
		uintptr(unsafe.Pointer(wlanConnectionParameters)),
		pReserved,
	)
	return syscall.Errno(r1)
}

//The WlanDisconnect function disconnects an interface from its current network.
//https://docs.microsoft.com/en-us/windows/win32/api/wlanapi/nf-wlanapi-wlandisconnect
func WlanDisconnect(handle *windows.Handle, pInterfaceGuid *syscall.GUID) error {
	r1, _, _ := wlanDisconnect.Call(
		uintptr(unsafe.Pointer(handle)),
		uintptr(unsafe.Pointer(pInterfaceGuid)),
		pReserved,
	)
	return syscall.Errno(r1)
}

//The WlanDeleteProfile function deletes a wireless profile for a wireless interface on the local computer.
//https://docs.microsoft.com/en-us/windows/win32/api/wlanapi/nf-wlanapi-wlandeleteprofile
func WlanDeleteProfile(handle *windows.Handle, pInterfaceGuid *syscall.GUID, strProfileName string) error {
	pProfileName, err := syscall.UTF16PtrFromString(strProfileName)
	if err != nil {
		log.Println(err)
		return err
	}
	r1, _, _ := wlanDeleteProfile.Call(
		uintptr(unsafe.Pointer(handle)),
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
func WlanGetAvailableNetworkList(handle *windows.Handle, pInterfaceGuid *syscall.GUID, dwFlags DWORD) (
	ppAvailableNetworkList *WLAN_AVAILABLE_NETWORK_LIST, err error) {

	r1, _, _ := wlanGetAvailableNetworkList.Call(
		uintptr(unsafe.Pointer(handle)),
		uintptr(unsafe.Pointer(pInterfaceGuid)),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(ppAvailableNetworkList)),
		pReserved,
	)
	err = syscall.Errno(r1)
	return
}

//The WlanGetFilterList function retrieves a group policy or user permission list.
//https://docs.microsoft.com/en-us/windows/win32/api/wlanapi/nf-wlanapi-wlangetfilterlist
func WlanGetFilterList(handle *windows.Handle, wlanFilterListType WLAN_FILTER_LIST_TYPE) (ppNetworkList *DOT11_NETWORK_LIST, err error) {
	r1, _, _ := wlanGetFilterList.Call(
		uintptr(unsafe.Pointer(handle)),
		uintptr(wlanFilterListType),
		pReserved,
		uintptr(unsafe.Pointer(ppNetworkList)),
	)
	err = syscall.Errno(r1)
	return
}

//The WlanGetInterfaceCapability function retrieves the capabilities of an interface.
//https://docs.microsoft.com/en-us/windows/win32/api/wlanapi/nf-wlanapi-wlangetinterfacecapability
func WlanGetInterfaceCapability(handle *windows.Handle, pInterfaceGuid *syscall.GUID) (ppCapability *WLAN_INTERFACE_CAPABILITY, err error) {
	r1, _, _ := wlanGetInterfaceCapability.Call(
		uintptr(unsafe.Pointer(handle)),
		uintptr(unsafe.Pointer(pInterfaceGuid)),
		pReserved,
		uintptr(unsafe.Pointer(ppCapability)),
	)
	err = syscall.Errno(r1)
	return
}

//The WlanGetNetworkBssList function retrieves a list of the basic service set (BSS) entries of the wireless network or networks on a given wireless LAN interface.
//https://docs.microsoft.com/zh-cn/windows/win32/api/wlanapi/nf-wlanapi-wlangetnetworkbsslist
func WlanGetNetworkBssList(handle *windows.Handle,
	pInterfaceGuid *syscall.GUID,
	pDot11Ssid *DOT11_SSID,
	dot11BssType DOT11_BSS_TYPE,
	bSecurityEnabled BOOL) (ppWlanBssList *WLAN_BSS_LIST, err error) {
	r1, _, _ := wlanGetNetworkBssList.Call(
		uintptr(unsafe.Pointer(handle)),
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
	handle *windows.Handle,
	pInterfaceGuid *syscall.GUID,
	strProfileName string) (pstrProfileXml string, pdwFlags *DWORD, pdwGrantedAccess *DWORD, err error) {
	pProfileName, err := syscall.UTF16PtrFromString(strProfileName)
	if err != nil {
		log.Println(err)
		return
	}
	var pstrProfile *uint16
	r1, _, _ := wlanGetProfile.Call(
		uintptr(unsafe.Pointer(handle)),
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
	handle *windows.Handle,
	pInterfaceGuid *syscall.GUID,
	strProfileName string) (pdwDataSize *DWORD, ppData *BYTE, err error) {
	pProfileName, err := syscall.UTF16PtrFromString(strProfileName)
	if err != nil {
		log.Println(err)
		return
	}
	r1, _, _ := wlanGetProfileCustomUserData.Call(
		uintptr(unsafe.Pointer(handle)),
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
func WlanGetProfileList(handle *windows.Handle, pInterfaceGuid *syscall.GUID) (ppProfileList *WLAN_PROFILE_INFO_LIST, err error) {
	r1, _, _ := wlanGetProfileList.Call(
		uintptr(unsafe.Pointer(handle)),
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
func WlanGetSecuritySettings(handle *windows.Handle, SecurableObject WLAN_SECURABLE_OBJECT) (
	pValueType *WLAN_OPCODE_VALUE_TYPE, pdwGrantedAccess *DWORD, err error) {

	var pstrCurrentSDDL *uint16
	r1, _, _ := wlanGetSecuritySettings.Call(
		uintptr(unsafe.Pointer(handle)),
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
func WlanGetSupportedDeviceServices(handle *windows.Handle, pInterfaceGuid *syscall.GUID) (
	ppDevSvcGuidList *WLAN_DEVICE_SERVICE_GUID_LIST, err error) {

	r1, _, _ := wlanGetSupportedDeviceServices.Call(
		uintptr(unsafe.Pointer(handle)),
		uintptr(unsafe.Pointer(pInterfaceGuid)),
		// out
		uintptr(unsafe.Pointer(&ppDevSvcGuidList)),
	)
	err = syscall.Errno(r1)
	return
}

//The WlanHostedNetworkForceStart function transitions the wireless Hosted Network to the wlan_hosted_network_active state without associating the request with the application's calling handle.
//https://docs.microsoft.com/zh-cn/windows/win32/api/wlanapi/nf-wlanapi-wlanhostednetworkforcestart
func WlanHostedNetworkForceStart(handle *windows.Handle) (pFailReason *WLAN_HOSTED_NETWORK_REASON, err error) {
	r1, _, _ := wlanHostedNetworkForceStart.Call(
		uintptr(unsafe.Pointer(handle)),
		// out
		uintptr(unsafe.Pointer(&pFailReason)),
		pReserved,
	)
	err = syscall.Errno(r1)
	return
}

//The WlanHostedNetworkForceStop function transitions the wireless Hosted Network to the wlan_hosted_network_idle without associating the request with the application's calling handle.
//https://docs.microsoft.com/zh-cn/windows/win32/api/wlanapi/nf-wlanapi-wlanhostednetworkforcestop
func WlanHostedNetworkForceStop(handle *windows.Handle) (pFailReason *WLAN_HOSTED_NETWORK_REASON, err error) {
	r1, _, _ := wlanHostedNetworkForceStop.Call(
		uintptr(unsafe.Pointer(handle)),
		// out
		uintptr(unsafe.Pointer(&pFailReason)),
		pReserved,
	)
	err = syscall.Errno(r1)
	return
}

//The WlanHostedNetworkInitSettings function configures and persists to storage the network connection settings (SSID and maximum number of peers, for example) on the wireless Hosted Network if these settings are not already configured.
//https://docs.microsoft.com/zh-cn/windows/win32/api/wlanapi/nf-wlanapi-wlanhostednetworkinitsettings
func WlanHostedNetworkInitSettings(handle *windows.Handle) (pFailReason *WLAN_HOSTED_NETWORK_REASON, err error) {
	r1, _, _ := wlanHostedNetworkInitSettings.Call(
		uintptr(unsafe.Pointer(handle)),
		// out
		uintptr(unsafe.Pointer(&pFailReason)),
		pReserved,
	)
	err = syscall.Errno(r1)
	return
}

//The WlanHostedNetworkQueryProperty function queries the current static properties of the wireless Hosted Network.
//https://docs.microsoft.com/zh-cn/windows/win32/api/wlanapi/nf-wlanapi-wlanhostednetworkqueryproperty
func WlanHostedNetworkQueryProperty(handle *windows.Handle, OpCode WLAN_HOSTED_NETWORK_OPCODE) (
	pdwDataSize *DWORD, ppvData *PVOID, pWlanOpcodeValueType *WLAN_OPCODE_VALUE_TYPE, err error) {

	r1, _, _ := wlanHostedNetworkQueryProperty.Call(
		uintptr(unsafe.Pointer(handle)),
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
func WlanHostedNetworkQuerySecondaryKey(handle *windows.Handle) (
	pdwKeyLength *DWORD, ppucKeyData *UCHAR, pbIsPassPhrase *BOOL, pbPersistent *BOOL, pFailReason WLAN_HOSTED_NETWORK_REASON, err error) {

	r1, _, _ := wlanHostedNetworkQuerySecondaryKey.Call(
		uintptr(unsafe.Pointer(handle)),
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
func WlanHostedNetworkQueryStatus(handle *windows.Handle) (ppWlanHostedNetworkStatus *WLAN_HOSTED_NETWORK_STATUS, err error) {
	r1, _, _ := wlanHostedNetworkQueryStatus.Call(
		uintptr(unsafe.Pointer(handle)),
		// out
		uintptr(unsafe.Pointer(&ppWlanHostedNetworkStatus)),
		pReserved,
	)
	err = syscall.Errno(r1)
	return
}

//The WlanHostedNetworkRefreshSecuritySettings function refreshes the configurable and auto-generated parts of the wireless Hosted Network security settings.
//https://docs.microsoft.com/zh-cn/windows/win32/api/wlanapi/nf-wlanapi-wlanhostednetworkrefreshsecuritysettings
func WlanHostedNetworkRefreshSecuritySettings(handle *windows.Handle) (pFailReason *WLAN_HOSTED_NETWORK_REASON, err error) {
	r1, _, _ := wlanHostedNetworkRefreshSecuritySettings.Call(
		uintptr(unsafe.Pointer(handle)),
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
	handle *windows.Handle,
	OpCode WLAN_HOSTED_NETWORK_OPCODE,
	dwDataSize DWORD,
	pvData PVOID) (pFailReason *WLAN_HOSTED_NETWORK_REASON, err error) {

	r1, _, _ := wlanHostedNetworkSetProperty.Call(
		uintptr(unsafe.Pointer(handle)),
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
	handle *windows.Handle,
	dwKeyLength DWORD,
	pucKeyData *UCHAR,
	bIsPassPhrase, bPersistent BOOL) (pFailReason WLAN_HOSTED_NETWORK_REASON, err error) {

	r1, _, _ := wlanHostedNetworkSetSecondaryKey.Call(
		uintptr(unsafe.Pointer(handle)),
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
func WlanHostedNetworkStartUsing(handle *windows.Handle) (pFailReason WLAN_HOSTED_NETWORK_REASON, err error) {

	r1, _, _ := wlanHostedNetworkStartUsing.Call(
		uintptr(unsafe.Pointer(handle)),
		// out
		uintptr(unsafe.Pointer(&pFailReason)),
		pReserved,
	)
	err = syscall.Errno(r1)
	return
}

//The WlanHostedNetworkStopUsing function stops the wireless Hosted Network.
//https://docs.microsoft.com/zh-cn/windows/win32/api/wlanapi/nf-wlanapi-wlanhostednetworkstopusing
func WlanHostedNetworkStopUsing(handle *windows.Handle) (pFailReason WLAN_HOSTED_NETWORK_REASON, err error) {

	r1, _, _ := wlanHostedNetworkStopUsing.Call(
		uintptr(unsafe.Pointer(handle)),
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
	handle *windows.Handle,
	pInterfaceGuid *syscall.GUID,
	Type WLAN_IHV_CONTROL_TYPE,
	dwInBufferSize DWORD,
	pInBuffer PVOID,
	dwOutBufferSize DWORD) (pOutBuffer *PVOID, pdwBytesReturned *DWORD, err error) {

	r1, _, _ := wlanIhvControl.Call(
		uintptr(unsafe.Pointer(handle)),
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

func EnumInterfaces(handle *windows.Handle) (interfaceInfoList *WLAN_INTERFACE_INFO_LIST, err error) {

	var iil *WLAN_INTERFACE_INFO_LIST

	r1, _, _ := fWlanEnumInterfaces.Call(
		uintptr(unsafe.Pointer(handle)),
		uintptr(0),
		uintptr(unsafe.Pointer(&iil)),
	)

	if r1 != S_OK {
		return nil, retValToError(r1)
	}

	return iil, nil

	//TODO: Need to rework to cleanup unmanaged memory WlanFreeMemory(uintptr(unsafe.Pointer(iil)))
}

func defaultInterface(handle *windows.Handle) (guid *syscall.GUID, description string, err error) {

	iil, err := EnumInterfaces(handle)
	if err != nil {
		return guid, "", err
	}

	wli := iil.InterfaceInfo[iil.dwIndex]

	return &wli.InterfaceGuid, windows.UTF16ToString(wli.strInterfaceDescription[:]), nil
}
