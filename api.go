package wlanapi

import (
	"fmt"
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

var (
	pReserved = uintptr(0)
)

func retValToError(errNo uintptr) error {
	return fmt.Errorf("return code: %x", errNo)
}

/*
	https://docs.microsoft.com/en-us/windows/win32/api/wlanapi/nf-wlanapi-wlanclosehandle
	DWORD WlanCloseHandle(
		[in] HANDLE hClientHandle,
		PVOID  pReserved
	);
*/
// WlanCloseHandle closes a connection to the server
func WlanCloseHandle(handle *windows.Handle) (err error) {
	r1, _, _ := fWlanCloseHandle.Call(
		uintptr(unsafe.Pointer(handle)),
		uintptr(0),
	)
	if r1 != S_OK {
		return retValToError(r1)
	}
	return nil
}

//WlanConnect attempts to connect to a specific network
//DWORD WlanConnect(
//  [in] HANDLE                            hClientHandle,
//  [in] const GUID                        *pInterfaceGuid,
//  [in] const PWLAN_CONNECTION_PARAMETERS pConnectionParameters,
//       PVOID                             pReserved
//);
func WlanConnect(handle *windows.Handle, guid *syscall.GUID) (err error) {
	wlanConnect.Call(
		uintptr(unsafe.Pointer(handle)),
		uintptr(unsafe.Pointer(guid)),

		pReserved,
	)
}

func OpenHandle() (handle *windows.Handle, err error) {

	clientVer := uintptr(2)
	var h windows.Handle

	r1, _, _ := fWlanOpenHandle.Call(
		clientVer,
		uintptr(0),
		uintptr(unsafe.Pointer(&clientVer)),
		uintptr(unsafe.Pointer(&h)),
	)

	if r1 != S_OK {
		return nil, retValToError(r1)
	}

	return &h, nil
}

func WlanScan(hClientHandle uintptr,
	pInterfaceGuid *syscall.GUID,
	pDot11Ssid *DOT11_SSID,
	pIeData *WLAN_RAW_DATA,
	pReserved uintptr) syscall.Errno {
	e, _, _ := hWlanScan.Call(hClientHandle,
		uintptr(unsafe.Pointer(pInterfaceGuid)),
		uintptr(unsafe.Pointer(pDot11Ssid)),
		uintptr(unsafe.Pointer(pIeData)),
		pReserved)
	return syscall.Errno(e)
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
