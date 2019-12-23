package main

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

func main() {
	err := keyboard.Open()
	if err != nil {
		panic(err)
	}
	defer keyboard.Close()

	fmt.Println("Press ESC to quit")
	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		} else if key == keyboard.KeyEsc {
			break
		}
		switch key {
		case keyboard.KeyArrowUp:
			fmt.Println("UP")
		case keyboard.KeyArrowDown:
			fmt.Println("DOWN")
		case keyboard.KeyArrowLeft:
			fmt.Println("LEFT")
		case keyboard.KeyArrowRight:
			fmt.Println("RIGHT")
		default:
			fmt.Printf("You pressed: %q\r\n", char)
		}
	}
}
