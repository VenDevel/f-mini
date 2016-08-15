package main

import (
	"f-mini/utils"
	"fmt"
)

func init() {
	read_conf()
}

func read_conf() {
	bContext, err := utils.ReadFileAllByPath("./conf/f-mini.conf")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(bContext))
}
