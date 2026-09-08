package gui

import (
	"syscall"
	"unsafe"
)

// ShowErrorPopup displays a native Windows error message box
func ShowErrorPopup(message string) {
	// Load user32.dll and get MessageBoxW procedure
	user32 := syscall.NewLazyDLL("user32.dll")
	messageBox := user32.NewProc("MessageBoxW")

	// Convert Go strings to UTF-16 pointers
	msgPtr, _ := syscall.UTF16PtrFromString(message)
	titlePtr, _ := syscall.UTF16PtrFromString("Error")

	// Flags: 0x00000010 = MB_ICONERROR, 0x00000000 = MB_OK
	const MB_OK = 0x00000000
	const MB_ICONERROR = 0x00000010

	// Call MessageBoxW
	messageBox.Call(
		0, // hWnd = 0 (no parent window)
		uintptr(unsafe.Pointer(msgPtr)),
		uintptr(unsafe.Pointer(titlePtr)),
		uintptr(MB_OK|MB_ICONERROR),
	)
}
