package api

import (
	"fmt"

	"github.com/go-vgo/robotgo"
	_ "github.com/go-vgo/robotgo/base"
	_ "github.com/go-vgo/robotgo/key"
	_ "github.com/go-vgo/robotgo/mouse"
	_ "github.com/go-vgo/robotgo/screen"
	_ "github.com/go-vgo/robotgo/window"
	"github.com/syzhang42/oxhorse/interception"
)

var (
	width  int
	height int
)

func init() {
	width, height = robotgo.GetScreenSize()
}

type MOUSE_BUTTON int

const (
	LEFT   MOUSE_BUTTON = 0
	RIGHT  MOUSE_BUTTON = 1
	MIDDLE MOUSE_BUTTON = 2
)

// 移动到目标位置（分辨率）并点击  坐标  左键还是右键   几下
func MouseMoveAndClick1(x, y int, t MOUSE_BUTTON, nums int) error {
	ct := interception.InterceptionCreateContext()
	defer interception.InterceptionDestroyContext(ct)
	mouseStroke := make([]interception.InterceptionMouseStroke, nums*2+1)

	mouseStroke[1].State = interception.INTERCEPTION_MOUSE_RIGHT_BUTTON_DOWN

	mouseStroke[2].State = interception.INTERCEPTION_MOUSE_RIGHT_BUTTON_UP

	_x, _y := 65535*x/width, 65535*y/height

	mouseStroke[0].Flags = interception.INTERCEPTION_MOUSE_MOVE_ABSOLUTE

	mouseStroke[0].X = int32(_x)

	mouseStroke[0].Y = int32(_y)

	switch t {
	case RIGHT:
		for i := 0; i < nums; i++ {
			mouseStroke[2*i+1].State = interception.INTERCEPTION_MOUSE_RIGHT_BUTTON_DOWN
			mouseStroke[2*i+2].State = interception.INTERCEPTION_MOUSE_RIGHT_BUTTON_UP
		}
	case LEFT:
		for i := 0; i < nums; i++ {
			mouseStroke[2*i+1].State = interception.INTERCEPTION_MOUSE_LEFT_BUTTON_DOWN
			mouseStroke[2*i+2].State = interception.INTERCEPTION_MOUSE_LEFT_BUTTON_UP
		}
	case MIDDLE:
		for i := 0; i < nums; i++ {
			mouseStroke[2*i+1].State = interception.INTERCEPTION_MOUSE_MIDDLE_BUTTON_DOWN
			mouseStroke[2*i+2].State = interception.INTERCEPTION_MOUSE_MIDDLE_BUTTON_UP
		}
	}
	interception.InterceptionSend(ct, interception.InterceptionDevice(interception.INTERCEPTION_MOUSE(0)), mouseStroke, uint32(len(mouseStroke)))
	return nil
}

// 移动到目标位置（屏幕比例）并点击  坐标（0-65535）  左键还是右键   几下
func MouseMoveAndClick2(x, y int, t MOUSE_BUTTON, nums int) error {
	if x > 65535 || x < 0 || y > 65535 || y < 0 {
		return fmt.Errorf("x,y must (0-65535)")
	}
	ct := interception.InterceptionCreateContext()
	defer interception.InterceptionDestroyContext(ct)
	mouseStroke := make([]interception.InterceptionMouseStroke, nums*2+1)

	mouseStroke[1].State = interception.INTERCEPTION_MOUSE_RIGHT_BUTTON_DOWN

	mouseStroke[2].State = interception.INTERCEPTION_MOUSE_RIGHT_BUTTON_UP

	mouseStroke[0].Flags = interception.INTERCEPTION_MOUSE_MOVE_ABSOLUTE

	mouseStroke[0].X = int32(x)

	mouseStroke[0].Y = int32(y)

	switch t {
	case RIGHT:
		for i := 0; i < nums; i++ {
			mouseStroke[2*i+1].State = interception.INTERCEPTION_MOUSE_RIGHT_BUTTON_DOWN
			mouseStroke[2*i+2].State = interception.INTERCEPTION_MOUSE_RIGHT_BUTTON_UP
		}
	case LEFT:
		for i := 0; i < nums; i++ {
			mouseStroke[2*i+1].State = interception.INTERCEPTION_MOUSE_LEFT_BUTTON_DOWN
			mouseStroke[2*i+2].State = interception.INTERCEPTION_MOUSE_LEFT_BUTTON_UP
		}
	case MIDDLE:
		for i := 0; i < nums; i++ {
			mouseStroke[2*i+1].State = interception.INTERCEPTION_MOUSE_MIDDLE_BUTTON_DOWN
			mouseStroke[2*i+2].State = interception.INTERCEPTION_MOUSE_MIDDLE_BUTTON_UP
		}
	}
	interception.InterceptionSend(ct, interception.InterceptionDevice(interception.INTERCEPTION_MOUSE(0)), mouseStroke, uint32(len(mouseStroke)))
	return nil
}
