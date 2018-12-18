package Blockchain

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
var code string

func TestGetHeight(ctx *testframework.TestFrameworkContext) bool {
	//DeployContract
	code = "5ac56b6a00527ac46a51527ac4681953797374656d2e53746f726167652e476574436f6e746578746a52527ac46a00c3096765744865696768749c76640300640c0065b4006c75666203006a00c3096765744865616465729c76640300640c006517016c75666203006a00c308676574426c6f636b9c76640300640c0065de016c75666203006a00c30e6765745472616e73616374696f6e9c76640300640c006579026c75666203006a00c30b676574436f6e74726163749c76640300640c0065ac036c75666203006a00c3146765745472616e73616374696f6e4865696768749c76640300640c006548046c7566620300006c756653c56b681953797374656d2e53746f726167652e476574436f6e746578746a00527ac4681b53797374656d2e426c6f636b636861696e2e4765744865696768746a51527ac46a51c3036765746a00c3681253797374656d2e53746f726167652e5075746a51c3681553797374656d2e52756e74696d652e4e6f74696679516c756655c56b681953797374656d2e53746f726167652e476574436f6e746578746a00527ac468244f6e746f6c6f67792e52756e74696d652e47657443757272656e74426c6f636b486173686a51527ac46a51c3681b53797374656d2e426c6f636b636861696e2e4765744865616465726a52527ac46a52c3681553797374656d2e4865616465722e476574486173686a53527ac46a51c307676574686173686a00c3681253797374656d2e53746f726167652e5075746a53c3116765746861736866726f6d6865616465726a00c3681253797374656d2e53746f726167652e507574516c756655c56b681953797374656d2e53746f726167652e476574436f6e746578746a00527ac468244f6e746f6c6f67792e52756e74696d652e47657443757272656e74426c6f636b486173686a51527ac46a51c3681a53797374656d2e426c6f636b636861696e2e476574426c6f636b6a52527ac46a52c3682053797374656d2e426c6f636b2e4765745472616e73616374696f6e436f756e746a53527ac46a53c3036765746a00c3681253797374656d2e53746f726167652e507574516c756658c56b681953797374656d2e53746f726167652e476574436f6e746578746a00527ac468244f6e746f6c6f67792e52756e74696d652e47657443757272656e74426c6f636b486173686a51527ac46a51c3681a53797374656d2e426c6f636b636861696e2e476574426c6f636b6a52527ac4006a52c3681b53797374656d2e426c6f636b2e4765745472616e73616374696f6e6a53527ac46a53c3681a53797374656d2e5472616e73616374696f6e2e476574486173686a54527ac46a54c3682053797374656d2e426c6f636b636861696e2e4765745472616e73616374696f6e6a55527ac46a55c3681a53797374656d2e5472616e73616374696f6e2e476574486173686a56527ac46a54c307676574486173686a00c3681253797374656d2e53746f726167652e5075746a56c30867657448617368316a00c3681253797374656d2e53746f726167652e507574516c756655c56b681953797374656d2e53746f726167652e476574436f6e746578746a00527ac4682d53797374656d2e457865637574696f6e456e67696e652e476574457865637574696e67536372697074486173686a51527ac46a51c3681d53797374656d2e426c6f636b636861696e2e476574436f6e74726163746a52527ac46a52c3681b4f6e746f6c6f67792e436f6e74726163742e4765745363726970746a53527ac46a53c3036765746a00c3681253797374656d2e53746f726167652e507574516c756657c56b681953797374656d2e53746f726167652e476574436f6e746578746a00527ac468244f6e746f6c6f67792e52756e74696d652e47657443757272656e74426c6f636b486173686a51527ac46a51c3681a53797374656d2e426c6f636b636861696e2e476574426c6f636b6a52527ac4006a52c3681b53797374656d2e426c6f636b2e4765745472616e73616374696f6e6a53527ac46a53c3681a53797374656d2e5472616e73616374696f6e2e476574486173686a54527ac46a54c3682653797374656d2e426c6f636b636861696e2e4765745472616e73616374696f6e4865696768746a55527ac46a55c3036765746a00c3681253797374656d2e53746f726167652e507574516c7566"
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
		[]interface{}{"getHeight", []interface{}{[]byte("getHeight")}})
	if err != nil {
		ctx.LogError("TestDomainSmartContract InvokeNeoVMSmartContract error: %s", err)
		return false
	}

	//SdkGetBlockHeight
	SdkGetBlockHeight, err := ctx.Ont.GetCurrentBlockHeight()
	if err != nil {
		ctx.LogError("ctx.Ont.GetCurrentBlockHeight error:%s", err)
		return false
	}
	ctx.LogInfo("SdkGetBlockHeight:", SdkGetBlockHeight)

	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestDomainSmartContract WaitForGenerateBlock error: %s", err)
		return false
	}

	//GetSvalueInStorage
	svalue, err := ctx.Ont.GetStorage(codeAddress.ToHexString(), []byte("get"))
	if err != nil {
		ctx.LogError("TestDomainSmartContract GetStorageItem key:hello error: %s", err)
		return false
	}

	//TransferValueToUint
	value := hex.EncodeToString(svalue)

	count := strings.Count(value, "") - 1
	s := []string{}
	for i := count; i > 0; i -= 2 {
		s = append(s, value[i-2:i])
	}
	s1 := strings.Join(s, "")

	height, err := strconv.ParseUint(s1, 16, 32)
	ctx.LogInfo("ContractGetBlockHeight:", height)

	if  uint32(height) - SdkGetBlockHeight > 2 {
		ctx.LogError("TestGetHeight error")
		return false
	}

	return true
}

func TestGetHeader(ctx *testframework.TestFrameworkContext) bool {

	signer, err := ctx.GetDefaultAccount()
	if err != nil {
		ctx.LogError("TestDomainSmartContract GetDefaultAccount error:%s", err)
		return false
	}

	//InvokeContract
	_, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"getHeader", []interface{}{[]byte("getHeader")}})
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
	hash, err := ctx.Ont.GetStorage(codeAddress.ToHexString(), []byte("gethash"))
	if err != nil {
		ctx.LogError("TestDomainSmartContract GetStorageItem key:hello error: %s", err)
		return false
	}
	shash := hex.EncodeToString(hash)
	count := strings.Count(shash, "") - 1
	s := []string{}
	for i := count; i > 0; i -= 2 {
		s = append(s, shash[i-2:i])
	}
	s1 := strings.Join(s, "")
	uhash, err := strconv.ParseUint(s1, 16, 32)
	ctx.LogInfo("hash:", uhash)

	header, err := ctx.Ont.GetStorage(codeAddress.ToHexString(), []byte("gethashfromheader"))
	if err != nil {
		ctx.LogError("TestDomainSmartContract GetStorageItem key:hello error: %s", err)
		return false
	}
	sheader := hex.EncodeToString(header)
	count1 := strings.Count(sheader, "") - 1
	s2 := []string{}
	for i := count1; i > 0; i -= 2 {
		s2 = append(s2, sheader[i-2:i])
	}
	s3 := strings.Join(s2, "")
	uheader, err := strconv.ParseUint(s3, 16, 32)
	ctx.LogInfo("header:", uheader)

	if uheader != uhash {
		ctx.LogError("TestHeader error")
		return false
	}
	return true
}

func TestGetBlock(ctx *testframework.TestFrameworkContext) bool {

	signer, err := ctx.GetDefaultAccount()
	if err != nil {
		ctx.LogError("TestDomainSmartContract GetDefaultAccount error:%s", err)
		return false
	}

	//InvokeContract
	_, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"getBlock", []interface{}{[]byte("getBlock")}})
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
		ctx.LogError("TestGetBlock error", err)
		return false
	}
	return true
}

func TestGetTransaction(ctx *testframework.TestFrameworkContext) bool {

	signer, err := ctx.GetDefaultAccount()
	if err != nil {
		ctx.LogError("TestDomainSmartContract GetDefaultAccount error:%s", err)
		return false
	}

	//InvokeContract
	_, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"getTransaction", []interface{}{[]byte("getTransaction")}})
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

func TestGetContract(ctx *testframework.TestFrameworkContext) bool {

	signer, err := ctx.GetDefaultAccount()
	if err != nil {
		ctx.LogError("TestDomainSmartContract GetDefaultAccount error:%s", err)
		return false
	}

	//InvokeContract
	_, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"getContract", []interface{}{[]byte("getContract")}})
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
	value, err := ctx.Ont.GetStorage(codeAddress.ToHexString(), []byte("get"))
	if err != nil {
		ctx.LogError("TestDomainSmartContract GetStorageItem key:hello error: %s", err)
		return false
	}
	script := hex.EncodeToString(value)
	ctx.LogInfo("	script:", script)

	if script != code {
		ctx.LogError("TestGetTransaction error")
		return false
	}
	return true
}

func TestGetTransactionHeight(ctx *testframework.TestFrameworkContext) bool {

	signer, err := ctx.GetDefaultAccount()
	if err != nil {
		ctx.LogError("TestDomainSmartContract GetDefaultAccount error:%s", err)
		return false
	}

	//InvokeContract
	_, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"getTransactionHeight", []interface{}{[]byte("getTransactionHeight")}})
	if err != nil {
		ctx.LogError("TestDomainSmartContract InvokeNeoVMSmartContract error: %s", err)
		return false
	}

	//SdkGetTransactionHeight
	SdkGetTransactionHeight, err := ctx.Ont.GetCurrentBlockHeight()
	if err != nil {
		ctx.LogError("ctx.Ont.GetCurrentTransactionHeighHeight error:%s", err)
		return false
	}
	ctx.LogInfo("SdkGetTransactionHeight:", SdkGetTransactionHeight)

	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestDomainSmartContract WaitForGenerateBlock error: %s", err)
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
	height, err := strconv.ParseUint(s1, 16, 32)
	ctx.LogInfo("ContractGetTransactionHeight:", height)

	if uint32(height) - SdkGetTransactionHeight > 2 {
		ctx.LogError("TestGetTransactionHeight error: %s")
		return false
	}

	return true
}
