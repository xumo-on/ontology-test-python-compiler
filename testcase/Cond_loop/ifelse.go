package Cond_loop

import (
	"github.com/ontio/ontology-go-sdk/utils"
	"github.com/ontio/ontology/common"
	"github.com/xumo-on/ontology-test-python-compiler/testframework"
)

func TestIfElse(ctx *testframework.TestFrameworkContext) bool {
	code := "58c56b6a00527ac46a51527ac46a00c3055768696c659c766403006419006a51c300c36a52527ac46a52c36542006c75666203006a00c3066966656c73659c766403006426006a51c300c36a52527ac46a51c351c36a53527ac46a53c36a52c3654e006c7566620300006c756656c56b6a00527ac4006a51527ac4006a52527ac46a52c36a00c39f76640300641c006a51c36a52c3936a51527ac46a52c351936a52527ac462dcff6a51c36c756655c56b6a00527ac46a51527ac46a00c36a51c3a076640300640a00516c7566621e006a00c36a51c39f76640300640c000051946c7566620700006c75666c7566"
	codeAddress, _ := utils.GetContractAddress(code)

	if !testIfElse(ctx, codeAddress, 23, 2) {
		return false
	}

	if !testIfElse(ctx, codeAddress, 2, 23) {
		return false
	}

	if !testIfElse(ctx, codeAddress, 0, 0) {
		return false
	}

	return true
}

func testIfElse(ctx *testframework.TestFrameworkContext, code common.Address, a, b int) bool {
	res, err := ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(
		code,
		[]interface{}{"ifelse",[]interface{}{a,b}},
	)
	if err != nil {
		ctx.LogError("TestIfElse InvokeSmartContract error:%s", err)
		return false
	}
	resValue, err := res.Result.ToInteger()
	if err != nil {
		ctx.LogError("TestIfElse Result.ToInteger error:%s", err)
		return false
	}
	err = ctx.AssertToInt(resValue, condIfElse(a, b))
	if err != nil {
		ctx.LogError("TestIfElse test %d ifelse %d failed %s", a, b, err)
		return false
	}
	return true
}

func condIfElse(a, b int) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	} else {
		return 0
	}
}