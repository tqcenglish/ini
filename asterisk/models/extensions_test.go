package models

import (
	"fmt"
	"testing"

	"github.com/tqcenglish/ini"
	"github.com/tqcenglish/ini/asterisk/models"
)

func Test_Extension(t *testing.T) {
	cfg, err := ini.Load("../../testdata/users_callcenter.conf")
	if err != nil {
		panic(err)
	}
	p := new(models.Extension)
	err = cfg.Section("8702").MapTo(p)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(p)

	fmt.Printf("section 8702 keys %v", cfg.Section("8702").Keys())
}
