package Datatype

import (
	"time"
	"github.com/ontio/ontology-go-sdk/utils"
	"github.com/xumo-on/ontology-test-python-compiler/testframework"
	"github.com/ontio/ontology/common"
)

func TestArray(ctx *testframework.TestFrameworkContext) bool {
	code := "5ec56b6a00527ac46a51527ac46a00c30541727261799c76640300640f006a51c365dd006c75666203006a00c307426f6f6c65616e9c76640300640c0065d0006c75666203006a00c3094279746561727261799c766403006426006a51c300c36a52527ac46a51c351c36a53527ac46a53c36a52c3659f006c75666203006a00c306496e746765729c76640300640c00659f006c75666203006a00c30a52657475726e747970659c766403006433006a51c300c36a52527ac46a51c351c36a53527ac46a51c352c36a54527ac46a54c36a53c36a52c36560006c75666203006a00c306537472696e679c76640300640c00657b006c7566620300006c756652c56b6a00527ac46a00c3c06c756651c56b516c756653c56b6a00527ac46a51527ac46a00c36a51c39c766403006c756651c56b5a6c756655c56b6a00527ac46a51527ac46a52527ac400c176c96a53527ac46a53c36a00c3c86a53c36a51c3c86a53c36a52c3c86a53c36c756651c56b0b48656c6c6f20576f726c646c7566"
	codeAddress, _ := utils.GetContractAddress(code)
	signer, err := ctx.GetDefaultAccount()
	if err != nil {
		ctx.LogError("TestArray GetDefaultAccount error:%s", err)
		return false
	}
	_, err = ctx.Ont.NeoVM.DeployNeoVMSmartContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		false,
		code,
		"TestArray",
		"1.0",
		"",
		"",
		"",
	)
	if err != nil {
		ctx.LogError("TestArray DeploySmartContract error:%s", err)
		return false
	}
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestArray WaitForGenerateBlock error:%s", err)
		return false
	}
	params := []interface{}{[]byte("Hello"), []byte("world")}
	if !testArray(ctx, codeAddress, params) {
		return false
	}

	params = []interface{}{[]byte("Hello"), []byte("world"), "123456", 8}
	if !testArray(ctx, codeAddress, params) {
		return false
	}
	return true
}

func testArray(ctx *testframework.TestFrameworkContext, code common.Address, params []interface{}) bool {
	res, err := ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(
		code,
		[]interface{}{"Array", params},
	)
	if err != nil {
		ctx.LogError("TestArray InvokeSmartContract error:%s", err)
		return false
	}
	resValue, err := res.Result.ToInteger()
	if err != nil {
		ctx.LogError("TestArray Result.ToInteger error:%s", err)
		return false
	}
	err = ctx.AssertToInt(resValue, len(params))
	if err != nil {
		ctx.LogError("TestArray test failed %s", err)
		return false
	}
	return true
}
