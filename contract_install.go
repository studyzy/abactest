package main

import (
	"fmt"
	"github.com/casbin/casbin"
	"github.com/casbin/casbin/model"
)

func TestContractInstallPermission(users []*User){
	m := model.Model{}
	rule:=modelText + "m = r.sub.IsDeveloper"
	m.LoadModelFromText(rule)
	e := casbin.NewEnforcer(m)
	for _,u:=range users{
		pass:=e.Enforce(u)
		fmt.Println(u.UserName, "InstallContract",pass)
	}
}
