package Block

import (
	"encoding/hex"
	"github.com/ontio/ontology-go-sdk/utils"
	"github.com/xumo-on/ontology-test-python-compiler/testframework"
	"github.com/ontio/ontology/common"
	"strconv"
	"strings"
	"time"
)

var codeAddress common.Address

func TestGetTransactionCount(ctx *testframework.TestFrameworkContext) bool {
	//DeployContract
	code := "56c56b6a00527ac46a51527ac4681953797374656d2e53746f726167652e476574436f6e746578746a52527ac46a00c3136765745472616e73616374696f6e436f756e749c76640300640c006531006c75666203006a00c30f6765745472616e73616374696f6e739c76640300640c0065cb006c7566620300006c756655c56b681953797374656d2e53746f726167652e476574436f6e746578746a00527ac468244f6e746f6c6f67792e52756e74696d652e47657443757272656e74426c6f636b486173686a51527ac46a51c3681a53797374656d2e426c6f636b636861696e2e476574426c6f636b6a52527ac46a52c3682053797374656d2e426c6f636b2e4765745472616e73616374696f6e436f756e746a53527ac46a53c3036765746a00c3681253797374656d2e53746f726167652e507574516c756658c56b681953797374656d2e53746f726167652e476574436f6e746578746a00527ac468244f6e746f6c6f67792e52756e74696d652e47657443757272656e74426c6f636b486173686a51527ac46a51c3681a53797374656d2e426c6f636b636861696e2e476574426c6f636b6a52527ac46a52c3681c53797374656d2e426c6f636b2e4765745472616e73616374696f6e736a53527ac46a53c300c3681a53797374656d2e5472616e73616374696f6e2e476574486173686a51527ac4006a52c3681b53797374656d2e426c6f636b2e4765745472616e73616374696f6e6a54527ac46a54c3681a53797374656d2e5472616e73616374696f6e2e476574486173686a55527ac46a51c307676574486173686a00c3681253797374656d2e53746f726167652e5075746a55c30867657448617368316a00c3681253797374656d2e53746f726167652e507574516c7566"
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
		[]interface{}{"getTransactionCount", []interface{}{[]byte("getTransactionCount")}})
	if err != nil {
		ctx.LogError("TestDomainSmartContract InvokeNeoVMSmartContract error: %s", err)
		return false
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestDomainSmartContract WaitForGenerateBlock error: %s", err)
		return false
	}

	//GetSvalueInStorage
	hash, err := ctx.Ont.GetStorage(codeAddress.ToHexString(), []byte("get"))
	if err != nil {
		ctx.LogError("TestDomainSmartContract GetStorageItem key:hello error: %s", err)
		return false
	}
	str := hex.EncodeToString(hash)
	count := strings.Count(str, "") - 1
	s := []string{}
	for i := count; i > 0; i -= 2 {
		s = append(s, str[i-2:i])
	}
	s1 := strings.Join(s, "")
	Count, err := strconv.ParseUint(s1, 16, 32)
	ctx.LogInfo("Count:", Count)

	if Count != 1 {
		ctx.LogError("TestGetTransactionCount error", err)
		return false
	}

	return true
}

func TestGetTransactions(ctx *testframework.TestFrameworkContext) bool {

	signer, err := ctx.GetDefaultAccount()
	if err != nil {
		ctx.LogError("TestDomainSmartContract GetDefaultAccount error:%s", err)
		return false
	}

	//InvokeContract
	_, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"getTransactions", []interface{}{[]byte("getTransactions")}})
	if err != nil {
		ctx.LogError("TestDomainSmartContract InvokeNeoVMSmartContract error: %s", err)
		return false
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestDomainSmartContract WaitForGenerateBlock error: %s", err)
		return false
	}

	//GetSvalueInStorage
	hash, err := ctx.Ont.GetStorage(codeAddress.ToHexString(), []byte("getHash"))
	if err != nil {
		ctx.LogError("TestDomainSmartContract GetStorageItem key:hello error: %s", err)
		return false
	}
	txHash := hex.EncodeToString(hash)
	ctx.LogInfo("	TxHash:", txHash)

	hash, err = ctx.Ont.GetStorage(codeAddress.ToHexString(), []byte("getHash1"))
	if err != nil {
		ctx.LogError("TestDomainSmartContract GetStorageItem key:hello error: %s", err)
		return false
	}
	txHash1 := hex.EncodeToString(hash)
	ctx.LogInfo("	TxHash1:", txHash1)

	if txHash != txHash1 {
		ctx.LogError("TestGetTransaction error", err)
		return false
	}
	return true
}
