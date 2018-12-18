package Ontology

import (
	"github.com/xumo-on/ontology-test-python-compiler/testcase/Ontology/Contract"
	"github.com/xumo-on/ontology-test-python-compiler/testcase/Ontology/Header"
	"github.com/xumo-on/ontology-test-python-compiler/testcase/Ontology/Native"
	"github.com/xumo-on/ontology-test-python-compiler/testcase/Ontology/Runtime"
	"github.com/xumo-on/ontology-test-python-compiler/testcase/Ontology/Transaction"
)

func TestOntology() {
	//Attribute.TestAttribute()
	Contract.TestContract()
	Header.TestHeader()
	Native.TestNative()
	Runtime.TestRuntime()
	Transaction.TestTransaction()
}
