package main

type User struct {
	Address  string
	UserName string
	CertId   int
	Type     string //
	Group    string //
	Credit   int    // 信用评分

}
func (u *User) IsDeveloper() bool{
	return u.UserName=="User2"
}

