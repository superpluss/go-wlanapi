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
	r1, _, _ := wlanGetInterfaceCapability.Call(
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
func WlanGetProfile(handle *windows.Handle,
	pInterfaceGuid *syscall.GUID,
	strProfileName string) (pstrProfileXml string, pdwFlags *DWORD, pdwGrantedAccess *DWORD, err error) {
	pProfileName, err := syscall.UTF16PtrFromString(strProfileName)
	if err != nil {
		log.Println(err)
		return
	}
	var pstrProfile *uint16
	r1, _, _ := wlanGetInterfaceCapability.Call(
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
