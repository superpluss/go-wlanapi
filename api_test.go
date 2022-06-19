package wlanapi

import (
	"golang.org/x/sys/windows"
	"log"
	"testing"
)

func handleSession() (handle windows.Handle) {
	handle, err := WlanOpenHandle()
	if err != nil {
		panic(err)
	}
	return handle
}

func Test_defaultInterface(t *testing.T) {
	session := handleSession()
	wii, err := defaultInterface(session)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(windows.UTF16ToString(wii.strInterfaceDescription[:]), wii.InterfaceGuid.String(), wii.isState)
}

func TestWlanScan(t *testing.T) {
	session := handleSession()
	defer WlanCloseHandle(session)
	wii, err := defaultInterface(session)
	if err != nil {
		log.Println(err)
		return
	}
	err = WlanScan(session, &wii.InterfaceGuid, nil, nil)
	if err != nil {
		t.Error(err)
	}
}

// 获取可用的WiFi列表
func TestWlanGetAvailableNetworkList(t *testing.T) {
	session := handleSession()
	defer WlanCloseHandle(session)
	wii, err := defaultInterface(session)
	if err != nil {
		log.Println(err)
		return
	}
	ppAvailableNetworkList, err := WlanGetAvailableNetworkList(session, &wii.InterfaceGuid, 0x00000002)
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("dwIndex:%d dwNumberOfItems:%d", ppAvailableNetworkList.dwIndex, ppAvailableNetworkList.dwNumberOfItems)
	for i := ppAvailableNetworkList.dwIndex; i < ppAvailableNetworkList.dwNumberOfItems; i++ {
		availableNetwork := ppAvailableNetworkList.Network[i]
		t.Log(availableNetwork.GetStrProfileName())
	}
}

//TODO
func TestWlanQueryInterface(t *testing.T) {
	session := handleSession()
	defer WlanCloseHandle(session)
	wii, err := defaultInterface(session)
	if err != nil {
		log.Println(err)
		return
	}
	queryInterface, data, valueType, err := WlanQueryInterface(session, &wii.InterfaceGuid, WlanIntfOpcodeCurrentConnection)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(*queryInterface, *data, *valueType)
}

func TestWlanGetProfile(t *testing.T) {

}
