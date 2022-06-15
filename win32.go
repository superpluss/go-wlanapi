//go:build windows
// +build windows

package wlanapi

var (
	NULL  = uintptr(0)
	FALSE = BOOL(0)
	TRUE  = BOOL(1)
)

type (
	//HANDLE ...
	HANDLE uintptr

	//HWND ...
	HWND      HANDLE
	HMODULE   HANDLE
	HHOOK     HANDLE
	HINSTANCE HANDLE
	HKL       HANDLE

	LRESULT uintptr
	LPARAM  uintptr
	WPARAM  uintptr

	//byte
	BYTE  byte
	CCHAR byte
	CHAR  rune

	//bool
	BOOL    int32
	BOOLEAN BYTE

	//void
	PVOID   uintptr
	LPVOID  uintptr
	LPCVOID uintptr
	LVOID   uintptr
)
type COLORREF DWORD

type (
	UCHAR uint8

	WORD   uint16
	USHORT uint16
	WCHAR  uint16
	SHORT  int16

	LONG    int32
	ULONG   uint32
	UINT    uint32
	DWORD   uint32
	DWORD32 uint32

	LONGLONG  int64
	ULONGLONG uint64
	DWORDLONG uint64
	DWORD64   uint64

	SIZE_T    ULONG_PTR
	KAFFINITY ULONG_PTR

	ULONG_PTR uintptr
	LONG_PTR  int

	NTSTATUS    LONG
	KPRIORITY   LONG
	PPEB        uintptr //not sure
	PWSTR       *WCHAR
	ACCESS_MASK ULONG
)

//
const (
	INVALID_HANDLE = LONG_PTR(-1)
)

const (
	MAX_PATH = 260
)
