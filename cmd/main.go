package main

import (
	"fmt"
	"github.com/godepsresolve/plugin"
)

func main() {
	fmt.Println(plugin.MakeSometingPluggable("HelloWorld"))
}
