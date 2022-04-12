/*
@Time : 2022/4/12 3:21 下午
@Author : fushisanlang
@File : clean
@Software: GoLand
*/
package tools

import (
	"os"
	"os/exec"
	"runtime"
)

var clear map[string]func() //create a map for storing clear funcs

func init() {
	clear = make(map[string]func()) //Initialize it

	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested

		cmd.Stdout = os.Stdout

		cmd.Run()

	}
	clear["darwin"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested

		cmd.Stdout = os.Stdout

		cmd.Run()

	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested

		cmd.Stdout = os.Stdout

		cmd.Run()

	}

}

func CallClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.

	if ok { //if we defined a clear func for that platform:

		value() //we execute it

	} else { //unsupported platform

		panic("Your platform is unsupported! I can't clear terminal screen :(")

	}

}

func Clean() {
	//time.Sleep(2 * time.Second)
	//CallClear()
	//fmt.Println("=====clean========")
}