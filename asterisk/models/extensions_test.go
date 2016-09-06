package models

import (
	"fmt"
	"testing"

	"github.com/ini"
)

func Test_Section(t *testing.T) {
	cfg, err := ini.Load("../../testdata/users_callcenter.conf")
	if err != nil {
		panic(err)
	}
	p := new(Extension)
	err = cfg.Section("8702").MapTo(p)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(p)
}
