package main

import "testing"

func TestUser_IsDeveloper(t *testing.T) {
	user1 := &User{
		Address:  "P13VBemDosoqQvQX6XaF84LsZMaRF7smxaF",
		UserName: "User1",
		CertId:   12345678,
		Type:     "user",
		Group:    "A",
		Credit:   50,
	}
	t.Log(user1.IsDeveloper())
	user2 := User{
		Address:  "P1R8oJsCypC2BgLRuxpnXW9S9gK9swYXQf",
		UserName: "User2",
		CertId:   87654321,
		Type:     "user",
		Group:    "B",
		Credit:   90,
	}
	t.Log(user2.IsDeveloper())
}
