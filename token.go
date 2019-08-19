package main

import (
	"fmt"
	"github.com/casbin/casbin"
	"github.com/casbin/casbin/model"
)

type Token struct{
	TokenName string
	PermissionRule string
}
func TestTokenPermission(users []*User){
	tokenA:=&Token{TokenName:"TokenA",
		PermissionRule:"m = r.sub.Group == \"A\""}
	tokenB:=&Token{TokenName:"TokenB",
		PermissionRule:"m = r.sub.Group == \"B\""}
	tokenC:=&Token{TokenName:"TokenC",
		PermissionRule:"m = r.sub.Credit >= 60"}
	tokens:=[]*Token{tokenA,tokenB,tokenC}
	for _,token:=range tokens{
		m := model.Model{}
		m.LoadModelFromText(modelText+token.PermissionRule)
		e := casbin.NewEnforcer(m)
		for _, u := range users {
			//强制决定一个“subject”是否可以通过操作“action”访问一个“object”，输入参数通常是:(sub, obj, act)。
			flag := e.Enforce(u, token)
			if flag {
				fmt.Println(token.TokenName,u.UserName,"Pass")
			} else {
				fmt.Println(token.TokenName,u.UserName,"Deny")
			}
		}
	}
}