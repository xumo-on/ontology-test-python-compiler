package Cond_loop

import (
	"github.com/xumo-on/ontology-test-python-compiler/testframework"
)

func TestCondLoop() {
	testframework.TFramework.RegTestCase("TestWhile", TestWhile)
	testframework.TFramework.RegTestCase("TestIfElse", TestIfElse)
}
