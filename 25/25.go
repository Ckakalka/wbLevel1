package main

import (
	"time"
)

func Sleep(seconds int) {
	<-time.After(time.Second * time.Duration(seconds))
}

func main() {
	Sleep(2)
}
