package main

import (
	"flag"
	"fmt"
	"gssh/config"
	"gssh/utils"
)

var (
	Version  = ""
	mVersion = flag.Bool("v", false, "version")
	mHost    = flag.String("s", "", "host")
	mPort    = flag.String("P", "22", "port")
	mUser    = flag.String("u", "root", "user")
	mPwd     = flag.String("p", "", "password")
	mDel     = flag.String("del", "", "del host")
	mList    = flag.Bool("l", false, "host list")
)

func main() {
	flag.Parse()

	if *mVersion {
		fmt.Println(Version)
		return
	}

	if *mList {
		utils.HostList()
		return
	}

	if *mDel != "" {
		config.DelHost(*mDel)
		return
	}

	if *mHost != "" {
		if ok := utils.CheckFile(*mHost); ok {
			c, err := config.ReadHost(*mHost)
			if err != nil {
				fmt.Println(err)
				return
			}
			utils.SshToServer(c.Host, c.Port, c.User, c.Pwd)
			return
		}
	}

	if *mHost != "" && *mPwd != "" {
		if err := config.WriteHost(*mHost, *mPort, *mUser, *mPwd); err != nil {
			fmt.Println(err)
		}
		utils.SshToServer(*mHost, *mPort, *mUser, *mPwd)
	} else {
		fmt.Println("Host or Pwd is null.")
	}

}
