package main

import (
	"flag"
	"fmt"

	"github.com/ComputePractice2018/photohosting/backend/utils"
)

func main() {
	var name := flag.String("name", "slimz", "имя для приветствия")
	flag.Parse()
	fmt.Println(utils.GetHelloWorldString(*name))
}