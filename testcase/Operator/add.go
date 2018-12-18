package Operator

import (
	"time"

	"github.com/ontio/ontology-go-sdk/utils"
	"github.com/xumo-on/ontology-test-python-compiler/testframework"
	"github.com/ontio/ontology/common"
)

func TestOperationAdd(ctx *testframework.TestFrameworkContext) bool {
	code := "0136c56b6a00527ac46a51527ac46a00c3034164649c766403006426006a51c300c36a52527ac46a51c351c36a53527ac46a53c36a52c3654b036c75666203006a00c303416e649c766403006426006a51c300c36a52527ac46a51c351c36a53527ac46a53c36a52c36530036c75666203006a00c3064469766964659c766403006426006a51c300c36a52527ac46a51c351c36a53527ac46a53c36a52c36516036c75666203006a00c30265719c766403006426006a51c300c36a52527ac46a51c351c36a53527ac46a53c36a52c365fc026c75666203006a00c3026c719c766403006426006a51c300c36a52527ac46a51c351c36a53527ac46a53c36a52c365e6026c75666203006a00c3096c65667473686966749c766403006426006a51c300c36a52527ac46a51c351c36a53527ac46a53c36a52c365c9026c75666203006a00c3026c729c766403006426006a51c300c36a52527ac46a51c351c36a53527ac46a53c36a52c365af026c75666203006a00c3046d6f64659c766403006426006a51c300c36a52527ac46a51c351c36a53527ac46a53c36a52c36597026c75666203006a00c3056d756c74699c766403006426006a51c300c36a52527ac46a51c351c36a53527ac46a53c36a52c3657a026c75666203006a00c3026e659c766403006426006a51c300c36a52527ac46a51c351c36a53527ac46a53c36a52c36560026c75666203006a00c3026e679c766403006419006a51c300c36a52527ac46a52c36557026c75666203006a00c3024f729c766403006426006a51c300c36a52527ac46a51c351c36a53527ac46a53c36a52c3653d026c75666203006a00c30a726967687473686966749c766403006426006a51c300c36a52527ac46a51c351c36a53527ac46a53c36a52c3651f026c75666203006a00c30273659c766403006426006a51c300c36a52527ac46a51c351c36a53527ac46a53c36a52c36505026c75666203006a00c30773656c666164649c766403006419006a51c300c36a52527ac46a52c365f7016c75666203006a00c30773656c667375629c766403006419006a51c300c36a52527ac46a52c365e6016c75666203006a00c302736c9c766403006426006a51c300c36a52527ac46a51c351c36a53527ac46a53c36a52c365cd016c75666203006a00c3037375629c766403006426006a51c300c36a52527ac46a51c351c36a53527ac46a53c36a52c365b6016c7566620300006c756653c56b6a00527ac46a51527ac46a00c36a51c3936c756653c56b6a00527ac46a51527ac46a00c36a51c39a766403006c756653c56b6a00527ac46a51527ac46a00c36a51c3966c756653c56b6a00527ac46a51527ac46a00c36a51c39c766403006c756653c56b6a00527ac46a51527ac46a00c36a51c3a2766403006c756653c56b6a00527ac46a51527ac46a00c36a51c3986c756653c56b6a00527ac46a51527ac46a00c36a51c3a0766403006c756653c56b6a00527ac46a51527ac46a00c36a51c3976c756653c56b6a00527ac46a51527ac46a00c36a51c3956c756653c56b6a00527ac46a51527ac46a00c36a51c39e766403006c756653c56b6a00527ac46a00c3916a51527ac46a51c36c756653c56b6a00527ac46a51527ac46a00c36a51c39b766303006c756653c56b6a00527ac46a51527ac46a00c36a51c3996c756653c56b6a00527ac46a51527ac46a00c36a51c3a1766403006c756653c56b6a00527ac46a00c351936a00527ac46a00c36c756653c56b6a00527ac46a00c351946a00527ac46a00c36c756653c56b6a00527ac46a51527ac46a00c36a51c39f766403006c756653c56b6a00527ac46a51527ac46a00c36a51c3946c7566"
	codeAddress, err := utils.GetContractAddress(code)
	if err != nil {
		ctx.LogError("TestOperationAdd GetContractAddress error:%s", err)
		return false
	}
	ctx.LogInfo("TestOperationAdd contact address:%s", codeAddress.ToHexString())
	signer, err := ctx.GetDefaultAccount()
	if err != nil {
		ctx.LogError("TestOperationAdd GetDefaultAccount error:%s", err)
		return false
	}
	tx, err := ctx.Ont.NeoVM.DeployNeoVMSmartContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		false,
		code,
		"TestOperationAdd",
		"1.0",
		"",
		"",
		"",
	)
	if err != nil {
		ctx.LogError("TestOperationAdd DeploySmartContract error:%s", err)
		return false
	}
	ctx.LogInfo("DeployContract TxHash:%s", tx.ToHexString())
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestOperationAdd WaitForGenerateBlock error:%s", err)
		return false
	}

	if !testOperationAdd(ctx, codeAddress, 1, 2) {
		return false
	}

	if !testOperationAdd(ctx, codeAddress, -1, 1) {
		return false
	}

	if !testOperationAdd(ctx, codeAddress, -1, -2) {
		return false
	}

	if !testOperationAdd(ctx, codeAddress, 0, 0) {
		return false
	}

	return true
}

func testOperationAdd(ctx *testframework.TestFrameworkContext, codeAddress common.Address, a, b int) bool {
	res, err := ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(
		codeAddress,
		[]interface{}{"Add", []interface{}{a, b}},
	)
	if err != nil {
		ctx.LogError("TestOperationAdd InvokeSmartContract error:%s", err)
		return false
	}
	resValue,err := res.Result.ToInteger()
	if err != nil {
		ctx.LogError("TestOperationAdd Result.ToInteger error:%s", err)
		return false
	}
	err = ctx.AssertToInt(resValue, a+b)
	if err != nil {
		ctx.LogError("TestOperationAdd test failed %s , %d, %d", err, a, b)
		return false
	}
	return true
}
