package interception

import (
	"fmt"
	"testing"
	"time"
)

func TestMoveAndClick(t *testing.T) {
	ct := InterceptionCreateContext()
	defer InterceptionDestroyContext(ct)
	mouseStroke := make([]InterceptionMouseStroke, 3)
	mouseStroke[0].Flags = INTERCEPTION_MOUSE_MOVE_ABSOLUTE

	mouseStroke[0].X = 65535 / 2

	mouseStroke[0].Y = 65535 / 2

	mouseStroke[1].State = INTERCEPTION_MOUSE_RIGHT_BUTTON_DOWN

	mouseStroke[2].State = INTERCEPTION_MOUSE_RIGHT_BUTTON_UP

	InterceptionSend(ct, InterceptionDevice(INTERCEPTION_MOUSE(0)), mouseStroke, uint32(len(mouseStroke)))
}

func TestKeyClick(t *testing.T) {
	start := time.Now()
	ct := InterceptionCreateContext()
	defer func() {
		InterceptionDestroyContext(ct)
		elapsed := time.Since(start)
		fmt.Println(elapsed)
	}()

	mouseStroke := make([]InterceptionKeyStroke, 2)
	mouseStroke[0].Code = uint16(MapVirtualKey(uint32('F'), MAPVK_VK_TO_VSC))
	mouseStroke[0].State = INTERCEPTION_KEY_DOWN
	mouseStroke[1].Code = mouseStroke[0].Code
	mouseStroke[1].State = INTERCEPTION_KEY_UP

	InterceptionSend(ct, InterceptionDevice(INTERCEPTION_KEYBOARD(0)), mouseStroke, uint32(len(mouseStroke)))
}
