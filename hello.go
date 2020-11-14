package main

import (
	"fmt"

	"github.com/binaryphile/hello-service/services/msgsvc"
)

func main() {
	fmt.Println(msgsvc.GetMessage())
}
