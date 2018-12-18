/*
from boa.interop.System.Runtime import Notify
from boa.interop.System.StorageContext import AsReadOnly
from boa.interop.System.Storage	import GetReadOnlyContext,GetContext

def Main(operation, args):
    if operation == 'asReadOnly':
        return asReadOnly()
    return False

def asReadOnly():
    context = GetContext()
    AsReadOnly(context)
    context1 = GetReadOnlyContext()
    Notify(context)
    Notify(context1)
    return True
 */
package StorageContext

import (
	"github.com/ontio/ontology-go-sdk/utils"
	"github.com/xumo-on/ontology-test-python-compiler/testframework"
	"time"
)

func TestAsReadOnly(ctx *testframework.TestFrameworkContext) bool {
	//DeployContract
	code := "56c56b6a00527ac46a51527ac4681953797374656d2e53746f726167652e476574436f6e746578746a52527ac46a00c30a6173526561644f6e6c799c76640300640c006525006c75666203006a00c3037075749c76640300640c0065d4006c7566620300006c756654c56b681953797374656d2e53746f726167652e476574436f6e746578746a00527ac46a00c3682053797374656d2e53746f72616765436f6e746578742e4173526561644f6e6c796a51527ac4682153797374656d2e53746f726167652e476574526561644f6e6c79436f6e746578746a52527ac46a00c3681553797374656d2e52756e74696d652e4e6f746966796a52c3681553797374656d2e52756e74696d652e4e6f746966796a51c3681553797374656d2e52756e74696d652e4e6f74696679516c756653c56b681953797374656d2e53746f726167652e476574436f6e746578746a00527ac4682153797374656d2e53746f726167652e476574526561644f6e6c79436f6e746578746a51527ac4046f6e6c79036765746a51c3681253797374656d2e53746f726167652e507574516c7566"
	codeAddress, _ := utils.GetContractAddress(code)

	ctx.LogInfo("=====CodeAddress===%s", codeAddress.ToHexString())
	signer, err := ctx.GetDefaultAccount()
	if err != nil {
		ctx.LogError("TestDomainSmartContract GetDefaultAccount error:%s", err)
		return false
	}

	_, err = ctx.Ont.NeoVM.DeployNeoVMSmartContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		true,
		code,
		"TestDomainSmartContract",
		"1.0",
		"",
		"",
		"",
	)
	if err != nil {
		ctx.LogError("TestDomainSmartContract DeploySmartContract error: %s", err)
		return false
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestDomainSmartContract WaitForGenerateBlock error: %s", err)
		return false
	}

	//InvokeContract
	txHash, err := ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"asReadOnly", []interface{}{[]byte("asReadOnly")}})
	if err != nil {
		ctx.LogError("TestDomainSmartContract InvokeNeoVMSmartContract error: %s", err)
		return false
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestInvokeSmartContract WaitForGenerateBlock error:%s", err)
		return false
	}

	//GetEventOfContract
	events, err := ctx.Ont.GetSmartContractEvent(txHash.ToHexString())
	if err != nil {
		ctx.LogError("TestInvokeSmartContract GetSmartContractEvent error:%s", err)
		return false
	}
	if events.State == 0 {
		ctx.LogError("TestInvokeSmartContract failed invoked exec state return 0")
		return false
	}

	context := events.Notify[0].States.(string)
	context1 := events.Notify[1].States.(string)
	newContext := events.Notify[2].States.(string)


	if context != context1 || context != newContext {
		ctx.LogError("TestAsReadOnly error")
		return true
	}

	//TestGetKeyInReadOnly
	_, err = ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"put", []interface{}{}})
	if err == nil {
		ctx.LogError("TestInvokeSmartContract GetValue error:%s", err)
		return false
	}

	return true
}
