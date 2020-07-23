package main

import (
	"context"
	"fmt"
	"github.com/jfrog-solutiontest/ci-basic-dep1"
)

func main () {
	fmt.Println ("Package Name: akita")
	ci_basic_dep1.PackageName()
	type Context = context.Context
	type CancelFunc = context.CancelFunc
}
