package interception

// #cgo CFLAGS: -I./include -Wno-attributes
// #cgo LDFLAGS: -L./lib/x64 -linterception
// #include "interception.h"
// #include <windows.h>
// int predicate_cgo(InterceptionDevice device);
import "C"

import (
	"unsafe"
)

const (
	MAPVK_VK_TO_VSC    = 0
	MAPVK_VSC_TO_VK    = 1
	MAPVK_VK_TO_CHAR   = 2
	MAPVK_VSC_TO_VK_EX = 3
)

func MapVirtualKey(code, mode uint32) uint32 {
	// 调用 Windows API 的 MapVirtualKey 函数
	return uint32(C.MapVirtualKey(C.UINT(code), C.UINT(mode)))
}

const (
	INTERCEPTION_MAX_KEYBOARD = 10
	INTERCEPTION_MAX_MOUSE    = 10
	INTERCEPTION_MAX_DEVICE   = INTERCEPTION_MAX_KEYBOARD + INTERCEPTION_MAX_MOUSE
)

func INTERCEPTION_KEYBOARD(index int) int {
	return index + 1
}
func INTERCEPTION_MOUSE(index int) int { return INTERCEPTION_MAX_KEYBOARD + index + 1 }

type (
	InterceptionContext    unsafe.Pointer                      //void *
	InterceptionDevice     int                                 // int
	InterceptionPrecedence int                                 //int
	InterceptionFilter     uint16                              // unsign short
	InterceptionPredicate  func(device InterceptionDevice) int //typedef int (*InterceptionPredicate)(InterceptionDevice device);
)

// InterceptionKeyState
type InterceptionKeyState int

const (
	INTERCEPTION_KEY_DOWN             = 0x00
	INTERCEPTION_KEY_UP               = 0x01
	INTERCEPTION_KEY_E0               = 0x02
	INTERCEPTION_KEY_E1               = 0x04
	INTERCEPTION_KEY_TERMSRV_SET_LED  = 0x08
	INTERCEPTION_KEY_TERMSRV_SHADOW   = 0x10
	INTERCEPTION_KEY_TERMSRV_VKPACKET = 0x20
)

type InterceptionFilterKeyState int

// InterceptionFilterKeyState
const (
	INTERCEPTION_FILTER_KEY_NONE             = 0x0000
	INTERCEPTION_FILTER_KEY_ALL              = 0xFFFF
	INTERCEPTION_FILTER_KEY_DOWN             = INTERCEPTION_KEY_UP
	INTERCEPTION_FILTER_KEY_UP               = INTERCEPTION_KEY_UP << 1
	INTERCEPTION_FILTER_KEY_E0               = INTERCEPTION_KEY_E0 << 1
	INTERCEPTION_FILTER_KEY_E1               = INTERCEPTION_KEY_E1 << 1
	INTERCEPTION_FILTER_KEY_TERMSRV_SET_LED  = INTERCEPTION_KEY_TERMSRV_SET_LED << 1
	INTERCEPTION_FILTER_KEY_TERMSRV_SHADOW   = INTERCEPTION_KEY_TERMSRV_SHADOW << 1
	INTERCEPTION_FILTER_KEY_TERMSRV_VKPACKET = INTERCEPTION_KEY_TERMSRV_VKPACKET << 1
)

type InterceptionMouseState int

// InterceptionMouseState
const (
	INTERCEPTION_MOUSE_LEFT_BUTTON_DOWN   = 0x001
	INTERCEPTION_MOUSE_LEFT_BUTTON_UP     = 0x002
	INTERCEPTION_MOUSE_RIGHT_BUTTON_DOWN  = 0x004
	INTERCEPTION_MOUSE_RIGHT_BUTTON_UP    = 0x008
	INTERCEPTION_MOUSE_MIDDLE_BUTTON_DOWN = 0x010
	INTERCEPTION_MOUSE_MIDDLE_BUTTON_UP   = 0x020

	INTERCEPTION_MOUSE_BUTTON_1_DOWN = INTERCEPTION_MOUSE_LEFT_BUTTON_DOWN
	INTERCEPTION_MOUSE_BUTTON_1_UP   = INTERCEPTION_MOUSE_LEFT_BUTTON_UP
	INTERCEPTION_MOUSE_BUTTON_2_DOWN = INTERCEPTION_MOUSE_RIGHT_BUTTON_DOWN
	INTERCEPTION_MOUSE_BUTTON_2_UP   = INTERCEPTION_MOUSE_RIGHT_BUTTON_UP
	INTERCEPTION_MOUSE_BUTTON_3_DOWN = INTERCEPTION_MOUSE_MIDDLE_BUTTON_DOWN
	INTERCEPTION_MOUSE_BUTTON_3_UP   = INTERCEPTION_MOUSE_MIDDLE_BUTTON_UP

	INTERCEPTION_MOUSE_BUTTON_4_DOWN = 0x040
	INTERCEPTION_MOUSE_BUTTON_4_UP   = 0x080
	INTERCEPTION_MOUSE_BUTTON_5_DOWN = 0x100
	INTERCEPTION_MOUSE_BUTTON_5_UP   = 0x200

	INTERCEPTION_MOUSE_WHEEL  = 0x400
	INTERCEPTION_MOUSE_HWHEEL = 0x800
)

type InterceptionFilterMouseState int

// InterceptionFilterMouseState
const (
	INTERCEPTION_FILTER_MOUSE_NONE = 0x0000
	INTERCEPTION_FILTER_MOUSE_ALL  = 0xFFFF

	INTERCEPTION_FILTER_MOUSE_LEFT_BUTTON_DOWN   = INTERCEPTION_MOUSE_LEFT_BUTTON_DOWN
	INTERCEPTION_FILTER_MOUSE_LEFT_BUTTON_UP     = INTERCEPTION_MOUSE_LEFT_BUTTON_UP
	INTERCEPTION_FILTER_MOUSE_RIGHT_BUTTON_DOWN  = INTERCEPTION_MOUSE_RIGHT_BUTTON_DOWN
	INTERCEPTION_FILTER_MOUSE_RIGHT_BUTTON_UP    = INTERCEPTION_MOUSE_RIGHT_BUTTON_UP
	INTERCEPTION_FILTER_MOUSE_MIDDLE_BUTTON_DOWN = INTERCEPTION_MOUSE_MIDDLE_BUTTON_DOWN
	INTERCEPTION_FILTER_MOUSE_MIDDLE_BUTTON_UP   = INTERCEPTION_MOUSE_MIDDLE_BUTTON_UP

	INTERCEPTION_FILTER_MOUSE_BUTTON_1_DOWN = INTERCEPTION_MOUSE_BUTTON_1_DOWN
	INTERCEPTION_FILTER_MOUSE_BUTTON_1_UP   = INTERCEPTION_MOUSE_BUTTON_1_UP
	INTERCEPTION_FILTER_MOUSE_BUTTON_2_DOWN = INTERCEPTION_MOUSE_BUTTON_2_DOWN
	INTERCEPTION_FILTER_MOUSE_BUTTON_2_UP   = INTERCEPTION_MOUSE_BUTTON_2_UP
	INTERCEPTION_FILTER_MOUSE_BUTTON_3_DOWN = INTERCEPTION_MOUSE_BUTTON_3_DOWN
	INTERCEPTION_FILTER_MOUSE_BUTTON_3_UP   = INTERCEPTION_MOUSE_BUTTON_3_UP

	INTERCEPTION_FILTER_MOUSE_BUTTON_4_DOWN = INTERCEPTION_MOUSE_BUTTON_4_DOWN
	INTERCEPTION_FILTER_MOUSE_BUTTON_4_UP   = INTERCEPTION_MOUSE_BUTTON_4_UP
	INTERCEPTION_FILTER_MOUSE_BUTTON_5_DOWN = INTERCEPTION_MOUSE_BUTTON_5_DOWN
	INTERCEPTION_FILTER_MOUSE_BUTTON_5_UP   = INTERCEPTION_MOUSE_BUTTON_5_UP

	INTERCEPTION_FILTER_MOUSE_WHEEL  = INTERCEPTION_MOUSE_WHEEL
	INTERCEPTION_FILTER_MOUSE_HWHEEL = INTERCEPTION_MOUSE_HWHEEL

	INTERCEPTION_FILTER_MOUSE_MOVE = 0x1000
)

type InterceptionMouseFlag int

// InterceptionMouseFlag
const (
	INTERCEPTION_MOUSE_MOVE_RELATIVE      = 0x000
	INTERCEPTION_MOUSE_MOVE_ABSOLUTE      = 0x001
	INTERCEPTION_MOUSE_VIRTUAL_DESKTOP    = 0x002
	INTERCEPTION_MOUSE_ATTRIBUTES_CHANGED = 0x004
	INTERCEPTION_MOUSE_MOVE_NOCOALESCE    = 0x008
	INTERCEPTION_MOUSE_TERMSRV_SRC_SHADOW = 0x100
)

var local InterceptionMouseStroke

//go:align 2
type InterceptionMouseStroke struct {
	State       uint16
	Flags       uint16
	Rolling     int16
	_           [2]byte // padding, to ensure proper alignment of x and y
	X           int32
	Y           int32
	Information uint32
}

//go:align 2
type InterceptionKeyStroke struct {
	Code        uint16
	State       uint16
	Information uint
}
type InterceptionStroke *InterceptionMouseStroke //数组地址

// Functions
func InterceptionCreateContext() InterceptionContext {
	return InterceptionContext(C.interception_create_context())
}

func InterceptionDestroyContext(context InterceptionContext) {
	C.interception_destroy_context(C.InterceptionContext(context))
}

func InterceptionGetPrecedence(context InterceptionContext, device InterceptionDevice) InterceptionPrecedence {
	return InterceptionPrecedence(C.interception_get_precedence(C.InterceptionContext(context), C.InterceptionDevice(device)))
}

func InterceptionSetPrecedence(context InterceptionContext, device InterceptionDevice, precedence InterceptionPrecedence) {
	C.interception_set_precedence(C.InterceptionContext(context), C.InterceptionDevice(device), C.InterceptionPrecedence(precedence))
}

func InterceptionGetFilter(context InterceptionContext, device InterceptionDevice) InterceptionFilter {
	return InterceptionFilter(C.interception_get_filter(C.InterceptionContext(context), C.InterceptionDevice(device)))
}

func InterceptionSetFilter(context InterceptionContext, predicate InterceptionPredicate, filter InterceptionFilter) {
	funcMap.Store("predicateFunc", predicate)
	C.interception_set_filter(C.InterceptionContext(context), C.InterceptionPredicate(unsafe.Pointer(C.predicate_cgo)), C.InterceptionFilter(filter))
}

func InterceptionWait(context InterceptionContext) InterceptionDevice {
	return InterceptionDevice(C.interception_wait(C.InterceptionContext(context)))
}

func InterceptionWaitWithTimeout(context InterceptionContext, milliseconds uint32) InterceptionDevice {
	return InterceptionDevice(C.interception_wait_with_timeout(C.InterceptionContext(context), C.ulong(milliseconds)))
}

func InterceptionSend[T any](context InterceptionContext, device InterceptionDevice, stroke []T, nstroke uint32) int {
	return int(C.interception_send(C.InterceptionContext(context), C.InterceptionDevice(device), (*[20]C.char)(unsafe.Pointer(&stroke[0])), C.uint(nstroke)))
}

func InterceptionReceive[T any](context InterceptionContext, device InterceptionDevice, stroke []T, nstroke uint32) int {
	return int(C.interception_receive(C.InterceptionContext(context), C.InterceptionDevice(device), (*C.InterceptionStroke)(unsafe.Pointer(&stroke[0])), C.uint(nstroke)))
}

func InterceptionGetHardwareID(context InterceptionContext, device InterceptionDevice, hardwareIDBuffer unsafe.Pointer, bufferSize uint32) uint32 {
	return uint32(C.interception_get_hardware_id(C.InterceptionContext(context), C.InterceptionDevice(device), hardwareIDBuffer, C.uint(bufferSize)))
}

func InterceptionIsInvalid(device InterceptionDevice) int {
	return int(C.interception_is_invalid(C.InterceptionDevice(device)))
}

func InterceptionIsKeyboard(device InterceptionDevice) int {
	return int(C.interception_is_keyboard(C.InterceptionDevice(device)))
}

func InterceptionIsMouse(device InterceptionDevice) int {
	return int(C.interception_is_mouse(C.InterceptionDevice(device)))
}
