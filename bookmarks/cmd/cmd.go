package cmd

import (
	"fmt"
	"os"
	"os/exec"
)

// go exec command using pipes
func CMDPsWithPipes() {
	binName := os.Getenv("BIN_NAME")
	_ = binName

	grep := exec.Command("grep", binName+"3")
	ps := exec.Command("ps", "-ef")

	pipe, _ := ps.StdoutPipe()
	defer pipe.Close()

	grep.Stdin = pipe

	ps.Start()

	res, _ := grep.Output()

	fmt.Println(res)
	fmt.Println(string(res))
}
