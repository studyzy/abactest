package main

import (
	"github.com/casbin/casbin"
	"github.com/casbin/casbin/model"
	"testing"
)

func BenchmarkCheckContractPermission(b *testing.B) {
	rule:=modelText + Init()
	u := &User{
		Address:  "P13VBemDosoqQvQX6XaF84LsZMaRF7smxaF",
		UserName: "User1",
		CertId:   12345678,
		Type:     "user",
		Group:    "A",
		Credit:   50,
	}

	for i:=0;i<b.N;i++{
		m := model.Model{}
		m.LoadModelFromText(rule)
		e := casbin.NewEnforcer(m)
		contract := &Contract{Creator: "User1", Address: "contractAddr", FuncName: "Fun1"}
		//fmt.Println(rule)
		e.Enforce(u, contract)
	}
}
