package shellcommand

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

func ExecShell(s_command string) (string, string) {
	var stdoutBuf, stderrBuf bytes.Buffer
	cmd := exec.Command("/bin/bash", "-c", s_command)
	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()
	var errStdout, errStderr error
	err := cmd.Start()
	if err != nil {
		log.Fatalf("cmd.Start() failed with '%s'\n", err)
	}
	stdout := io.MultiWriter(os.Stdout, &stdoutBuf)
	stderr := io.MultiWriter(os.Stderr, &stderrBuf)
	_, errStdout = io.Copy(stdout, stdoutIn)
	_, errStderr = io.Copy(stderr, stderrIn)
	err = cmd.Wait()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	if errStdout != nil || errStderr != nil {
		fmt.Printf("stdout: %v, stderr: %v\n", errStdout, errStderr)
		log.Fatal("failed to capture stdout or stderr\n")
	}
	outStr, errStr := string(stdoutBuf.Bytes()), string(stderrBuf.Bytes())
	return outStr, errStr
}
func RunCommand(commandStr string) string {
	cmdstr := commandStr
	out, _ := exec.Command("sh", "-c", cmdstr).Output()
	strout := string(out)

	return strout
}
