package main

import (
	"fmt"

	"golang.org/x/sys/windows"
)

const enable_virtual_terminal_processing = 0x4 // https://docs.microsoft.com/en-us/windows/console/setconsolemode

func setVT(enable bool) error {
	var err error
	var screenBuf windows.Handle
	screenBuf, err = windows.Open("CONOUT$", windows.O_RDWR, 0)
	if err != nil {
		return err
	}

	var mode uint32
	err = windows.GetConsoleMode(screenBuf, &mode)
	if err != nil {
		return err
	}

	if enable {
		mode |= enable_virtual_terminal_processing
	} else {
		mode &^= enable_virtual_terminal_processing
	}
	err = windows.SetConsoleMode(screenBuf, mode)
	return err
}

func draw_256ramp() {
	var i int
	for i = -4; i < 256; i++ {
		if i < 0 {
			fmt.Print("  ")
		} else {
			fmt.Printf("\033[38;5;%dm%03d ", i, i)
		}
		if (i + 2) % 6 == 5 {
			fmt.Print("\n")
		}
	}
}

func draw_rgb_cube() {
	var r, g, b int
	b = 100
	for r = 0; r <= 255; r += 10 {
		fmt.Print("\033[0m\n")
		for g = 0; g <= 255; g += 5 {
			fmt.Printf("\033[48;2;%d;%d;%dm ", r, g, b)
		}
	}
}

func clearColor() {
	fmt.Print("\033[0m")
}

func main() {
	var err error = setVT(true)
	if err == nil {
		fmt.Println("\033[36mCyan")
		fmt.Println("\033[38;5;220mYellow")
		fmt.Println("\033[38;2;255;0;0mRed")
		draw_256ramp()
		draw_rgb_cube()
	} else {
		fmt.Println("error:")
		fmt.Println(err)
	}
	clearColor()
	err = setVT(false)
}