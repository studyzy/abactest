package main

import (
	"fmt"
	"github.com/casbin/casbin"
	"github.com/casbin/casbin/model"
)

const modelText = `
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj,act

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m =  judgeObj(r.sub.Group,r.sub.Type)  &&  judgeAct(r.sub.Method)
`

//r.sub.Type == r.obj.GetType() &&  r.sub.Group == r.obj.Group  &&  r.sub.Method == r.act.Method

func DefineFunction(group, userType string) bool {
	if group == "gptn.mediator1" && userType == "user" {
		return true
	} else if group == "gptn.mediator1" && userType == "client" {
		return true
	} else {
		return false
	}
}

func DefineFunction2(method float64) bool {
	if method == 1 || method == 3 {
		return true
	} else {
		return false
	}
}

func DefineFunctionWrapper(args ...interface{}) (interface{}, error) {
	group := args[0].(string)
	userType := args[1].(string)
	return bool(DefineFunction(group,userType)),nil
}

func DefineFunctionWrapper2(args ...interface{}) (interface{}, error) {
	method := args[0].(float64)
	return bool(DefineFunction2(method)),nil
}

type User struct {
	Address  string
	UserName string
	CertId   int
	Type     string //user 、client
	Group    string //gptn.mediator1 or gptn.mediator2
	Method   int    //0 无权限  1 可以发起交易 2 可以执行合约 3  1&2
}

type Obj struct {
	Type  string //user 、client
	Group string //gptn.mediator1 or gptn.mediator2
}

type Act struct {
	Method int //0 无权限  1 可以发起交易 2 可以执行合约 3  1&2
}

func main() {
	m := model.Model{}
	m.LoadModelFromText(modelText)
	e := casbin.NewEnforcer(m)
	e.AddFunction("judgeObj",DefineFunctionWrapper)
	e.AddFunction("judgeAct",DefineFunctionWrapper2)
	user1 := User{
		Address:  "P13VBemDosoqQvQX6XaF84LsZMaRF7smxaF",
		UserName: "lk",
		CertId:   12345678,
		Type:     "user",
		Group:    "gptn.mediator1",
		Method:   1,
	}

	user2 := User{
		Address:  "P1R8oJsCypC2BgLRuxpnXW9S9gK9swYXQf",
		UserName: "lk",
		CertId:   87654321,
		Type:     "user",
		Group:    "gptn.mediator1",
		Method:   1,
	}

	//  用户 palletone  属于 group1 , group2
	users := []User{user1, user2}
	//资源属性
	obj := Obj{
		Type:  "user",
		Group: "gptn.mediator1",
	}
	//当执行操作权限为1 策略执行结果为allow
	//act := Act{
	//	Method: 1,
	//}
	// 检查 用户 palletone 所有的组  是否有权限
	for _, v := range users {
		//强制决定一个“subject”是否可以通过操作“action”访问一个“object”，输入参数通常是:(sub, obj, act)。
		flag := e.Enforce(v, obj)
		if flag {
			fmt.Println("权限正常")
		} else {
			fmt.Println("没有权限")
		}
	}
}
