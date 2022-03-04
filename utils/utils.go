package utils

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
)

func GetHome() string {
	u, err := user.Current()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return u.HomeDir + `/.gssh/`
}

func CheckFile(host string) bool {
	name := GetHome() + host
	_, err := os.Lstat(name)
	return !os.IsNotExist(err)
}

func HostList() {
	home := GetHome()
	fs, err := os.ReadDir(home)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, f := range fs {
		fmt.Println(f.Name())
	}
}

func SshToServer(host, port, user, pwd string) {
	cmd := exec.Command("expect", "-c", `spawn ssh -o StrictHostKeyChecking=no -p `+port+` `+user+`@`+host+`;expect "*assword:*";send "`+pwd+`\r";interact`)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Run()
}
