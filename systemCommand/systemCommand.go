package systemCommand

import (
	"fmt"
	"os/exec"
)

func main() {
	//cmd := exec.Command("calc.exe")
	cmd := exec.Command("ls", "-la")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(output))
}
