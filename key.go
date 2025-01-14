package woagh

import "github.com/syzhang42/oxhorse/interception"

func KeyClick(code rune, nums uint32) {
	ct := interception.InterceptionCreateContext()
	defer interception.InterceptionDestroyContext(ct)

	mouseStroke := make([]interception.InterceptionKeyStroke, 2*nums)
	for i := 0; i < int(nums); i++ {
		mouseStroke[2*i].Code = uint16(interception.MapVirtualKey(uint32(code), interception.MAPVK_VK_TO_VSC))
		mouseStroke[2*i].State = interception.INTERCEPTION_KEY_DOWN
		mouseStroke[2*i+1].Code = mouseStroke[0].Code
		mouseStroke[2*i+1].State = interception.INTERCEPTION_KEY_UP
	}
	interception.InterceptionSend(ct, interception.InterceptionDevice(interception.INTERCEPTION_KEYBOARD(0)), mouseStroke, uint32(len(mouseStroke)))
}
