//定义了合约的某些方法能够被A用户调用，
package main

import (
	"fmt"
	"github.com/casbin/casbin"
	"github.com/casbin/casbin/model"
)

func QueryPermissionRule(contractAddr string) string {
	//GroupA能够访问ContractA的方法Fun1,Fun2
	//GroupB能够访问ContractA的方法Fun3，Fun4
	//ContractA创建者能够访问ContractA的所有方法
	ruleA := "m = (r.sub.Group==\"A\" && r.obj in ('Fun1','Fun2'))"
	ruleA += "|| (r.sub.Group==\"B\" && r.obj in (\"Fun3\",\"Fun4\"))"
	// ruleA+="|| r.sub.UserName==r.obj.Owner()"

	return ruleA
}
func CheckContractPermission(u *User, contractAddr string, function string) {
	m := model.Model{}
	rule := modelText + QueryPermissionRule(contractAddr)
	m.LoadModelFromText(rule)
	e := casbin.NewEnforcer(m)
	// contract := &Contract{Creator: "User1", Address: contractAddr, FuncName: function}
	//fmt.Println(rule)
	flag := e.Enforce(u, function)
	fmt.Println(contractAddr, u.UserName, function, flag)
}
func TestContractPermission(users []*User) {
	for _, u := range users {
		CheckContractPermission(u, "ContractA", "Fun1")
		CheckContractPermission(u, "ContractA", "Fun2")
		CheckContractPermission(u, "ContractA", "Fun3")
		CheckContractPermission(u, "ContractA", "Fun4")
	}
}
