package main

import (
	"github.com/sivchari/ctxcause"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(ctxcause.Analyzer) }
