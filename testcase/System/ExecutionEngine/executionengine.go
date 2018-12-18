package ExecutionEngine

import (
	"encoding/hex"
	"github.com/ontio/ontology-go-sdk/utils"
	"github.com/xumo-on/ontology-test-python-compiler/testframework"
	"github.com/ontio/ontology/common"
	"strings"
	"time"
)

var	codeAddress common.Address

func TestGetScriptContainer(ctx *testframework.TestFrameworkContext) bool {

	//DeployContract
	code := "57c56b6a00527ac46a51527ac4681953797374656d2e53746f726167652e476574436f6e746578746a52527ac46a00c312676574536372697074436f6e7461696e65729c76640300640c006561006c75666203006a00c316676574457865637574696e67536372697074486173689c76640300640c006587016c75666203006a00c31467657443616c6c696e67536372697074486173689c76640300640c0065f1016c7566620300006c756658c56b681953797374656d2e53746f726167652e476574436f6e746578746a00527ac4682953797374656d2e457865637574696f6e456e67696e652e476574536372697074436f6e7461696e65726a51527ac46a51c3681a53797374656d2e5472616e73616374696f6e2e476574486173686a52527ac468244f6e746f6c6f67792e52756e74696d652e47657443757272656e74426c6f636b486173686a53527ac46a53c3681a53797374656d2e426c6f636b636861696e2e476574426c6f636b6a54527ac4006a54c3681b53797374656d2e426c6f636b2e4765745472616e73616374696f6e6a55527ac46a55c3681a53797374656d2e5472616e73616374696f6e2e476574486173686a56527ac46a52c3036765746a00c3681253797374656d2e53746f726167652e5075746a56c304676574316a00c3681253797374656d2e53746f726167652e507574516c756653c56b681953797374656d2e53746f726167652e476574436f6e746578746a00527ac4682d53797374656d2e457865637574696f6e456e67696e652e476574457865637574696e67536372697074486173686a51527ac46a51c3036765746a00c3681253797374656d2e53746f726167652e5075746a51c3681553797374656d2e52756e74696d652e4e6f74696679516c756654c56b681953797374656d2e53746f726167652e476574436f6e746578746a00527ac4682b53797374656d2e457865637574696f6e456e67696e652e47657443616c6c696e67536372697074486173686a51527ac4682953797374656d2e457865637574696f6e456e67696e652e476574456e747279536372697074486173686a52527ac46a51c3681553797374656d2e52756e74696d652e4e6f746966796a52c3681553797374656d2e52756e74696d652e4e6f74696679516c7566"
	codeAddress, _ = utils.GetContractAddress(code)

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
	_, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"getScriptContainer", []interface{}{[]byte("getScriptContainer")}})
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

	//GetSvalueInStorage
	svalue, err := ctx.Ont.GetStorage(codeAddress.ToHexString(), []byte("get"))
	if err != nil {
		ctx.LogError("TestDomainSmartContract GetStorageItem key:hello error: %s", err)
		return false
	}
	value := hex.EncodeToString(svalue)

	svalue1, err := ctx.Ont.GetStorage(codeAddress.ToHexString(), []byte("get1"))
	if err != nil {
		ctx.LogError("TestDomainSmartContract GetStorageItem key:hello error: %s", err)
		return false
	}
	value1 := hex.EncodeToString(svalue1)

	ctx.LogInfo("hash:",value)
	ctx.LogInfo("hash1:",value1)

	if value != value1 {
		ctx.LogError("TestGetScriptContainer error")
		return false
	}
	return true
}

func TestGetExecutingScriptHash(ctx *testframework.TestFrameworkContext) bool {

	signer, err := ctx.GetDefaultAccount()
	if err != nil {
		ctx.LogError("TestDomainSmartContract GetDefaultAccount error:%s", err)
		return false
	}

	//InvokeContract
	_, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"getExecutingScriptHash", []interface{}{[]byte("getExecutingScriptHash")}})
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

	//GetSvalueInStorage
	svalue, err := ctx.Ont.GetStorage(codeAddress.ToHexString(), []byte("get"))
	if err != nil {
		ctx.LogError("TestDomainSmartContract GetStorageItem key:hello error: %s", err)
		return false
	}
	value := hex.EncodeToString(svalue)

	count := strings.Count(value, "") - 1
	s := []string{}
	for i := count; i > 0; i -= 2 {
		s = append(s, value[i-2:i])
	}
	s1 := strings.Join(s, "")
	ctx.LogInfo("hash:",s1)

	if s1 != codeAddress.ToHexString() {
		ctx.LogError("TestGetExecutingScriptHash error")
		return false
	}

	return true
}

func TestGetCallingScriptHash(ctx *testframework.TestFrameworkContext) bool {

	signer, err := ctx.GetDefaultAccount()
	if err != nil {
		ctx.LogError("TestDomainSmartContract GetDefaultAccount error:%s", err)
		return false
	}

	//InvokeContract
	txHash, err := ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"getCallingScriptHash", []interface{}{[]byte("getCallingScriptHash")}})
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
	notify := events.Notify[0]
	notify1 := events.Notify[1]
	ctx.LogInfo("notify: ", notify.States)
	ctx.LogInfo("notify1: ", notify1.States)

	if notify.States != notify1.States {
		ctx.LogError("getCallingScriptHash error")
		return false
	}

	return true
}
