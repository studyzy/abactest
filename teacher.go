package main

import (
	"fmt"
	"github.com/casbin/casbin"
	"github.com/casbin/casbin/model"
	"time"
)

type Person struct{
	Role string
	Name string
}
type Gate struct{
	Name string
}
type Env struct{
	Time time.Time
	Location string
}
func (env *Env) IsSchooltime() bool{
	return env.Time.Hour()>=7&&env.Time.Hour()<=18
}
func TestTeacherEnterSchoolGate() {
	p1 := Person{Role: "Student", Name: "Yun"}
	p2 := Person{Role: "Teacher", Name: "Devin"}
	persons := []Person{p1, p2}
	g1 := Gate{Name: "School Gate"}
	g2 := Gate{Name: "Factory Gate"}
	gates := []Gate{g1, g2}
	const modelText = `
[request_definition]
r = sub, obj, act, env

[policy_definition]
p = sub, obj,act

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = r.sub.Role=='Teacher' && r.obj.Name=='School Gate' && r.act in('In','Out') && r.env.Time.Hour >7 && r.env.Time.Hour <= 18
`
	//m = r.sub.Role=='Teacher' && r.obj.Name=='School Gate' && r.act in('In','Out') && r.env.IsSchooltime()
	m := model.Model{}

	m.LoadModelFromText(modelText)
	e := casbin.NewEnforcer(m)
	envs := []*Env{InitEnv(9), InitEnv(23)}
	for _, env := range envs {
		fmt.Println("\r\nTime:",env.Time.Local())
		for _, p := range persons {
			for _, g := range gates {
				pass := e.Enforce(p, g, "In", env)
				fmt.Println(p.Role, p.Name, "In", g.Name, pass)
				pass = e.Enforce(p, g, "Control", env)
				fmt.Println(p.Role,p.Name, "Control", g.Name, pass)
			}
		}
	}
}

func InitEnv(hour int) *Env{
	env:=&Env{}
	env.Time=time.Date(2019,8,20,hour,0,0,0,time.Local)
	return env
}
