package orka

import (
	"log"
	"os/exec"
	// "fmt"
)

func RunCommand(args ...string) (string, error) {

		// var stdout, stderr bytes.Buffer

		log.Printf("Executing args: %#v", args)
		cmd := exec.Command(args[0], args[1:]...)

    stdout, err := cmd.Output()
		return string(stdout), err

    if err != nil {
				log.Println("Error")
        log.Println(err.Error())
        return string(stdout), err
    }
		
    log.Println(string(stdout))
		return string(stdout), nil
}
