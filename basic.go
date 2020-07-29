package main

import (
	"context"
	"fmt"
	"github.com/jfrog-solutiontest/test-basic-dep1"
)

func main () {
	fmt.Println ("Package Name: akita")
	test_basic_dep1.PackageName()
	type Context = context.Context
	type CancelFunc = context.CancelFunc
}
