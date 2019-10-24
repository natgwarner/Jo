package main

import (
	"os/exec"
	"fmt"
)

func main() {
	exec.Command("javac", "Hello.java").Run()
	out, _ := exec.Command("java", "Hello").Output()
	fmt.Println(string(out))
}
