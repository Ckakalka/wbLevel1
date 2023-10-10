package main

import (
	"fmt"
	"time"
)

type Human struct {
	Name      string
	BirthDate time.Time
}

func (h Human) String() string {
	return fmt.Sprintf("Name: %s, BirthDate: %s", h.Name, h.BirthDate.Format("02-01-2006"))
}

type Action struct {
	Human
	Info string
}

func (a Action) String() string {
	return fmt.Sprintf("%s, info: %s", a.Human, a.Info)
}

func main() {
	human := Human{"Valeriy", time.Now()}
	fmt.Println(human)
	action := Action{Human: human, Info: "Sleep"}
	fmt.Println(action)
}
