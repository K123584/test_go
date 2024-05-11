package main

import (
	"fmt"

	"github.com//K123584/test_go/namer"
)

func main() {
	user1 := namer.User{"田中", "裕二", 35}
	fmt.Println(user1.UserName("tana"))
}
