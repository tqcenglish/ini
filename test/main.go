package main

import (
	"fmt"

	"github.com/tqcenglish/ini"
	"github.com/tqcenglish/ini/asterisk/models"
)

func main() {
	fmt.Print("hello world!")
	cfg, err := ini.Load("../testdata/users_callcenter.conf")
	if err != nil {
		panic(err)
	}
	p := new(models.Extension)
	err = cfg.Section("8702").MapTo(p)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(p)

	fmt.Printf("\nsection 8702 keys %v\n", cfg.Section("8702").KeyStrings())
	fmt.Printf("section sipCallCenter keys %v\n", cfg.Section("sipCallCenter").KeyStrings())
}
