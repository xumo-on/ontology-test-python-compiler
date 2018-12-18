package Header

import (
	"encoding/hex"
	"github.com/ontio/ontology-go-sdk/utils"
	"github.com/xumo-on/ontology-test-python-compiler/testframework"
	"github.com/ontio/ontology/common"
	"strconv"
	"strings"
	"time"
)

var	codeAddress common.Address

func TestGetVersion(ctx *testframework.TestFrameworkContext) bool {
	//DeployContract
	code := "57c56b6a00527ac46a51527ac46a00c30a67657456657273696f6e9c76640300640c006579006c75666203006a00c30d6765744d65726b6c65526f6f749c76640300640c0065d4006c75666203006a00c310676574436f6e73656e737573446174619c76640300640c006561016c75666203006a00c3106765744e657874436f6e73656e7375739c76640300640c0065f1016c7566620300006c756654c56b68244f6e746f6c6f67792e52756e74696d652e47657443757272656e74426c6f636b486173686a00527ac46a00c3681b53797374656d2e426c6f636b636861696e2e4765744865616465726a51527ac46a51c3681a4f6e746f6c6f67792e4865616465722e47657456657273696f6e6a52527ac46a52c36c756654c56b68244f6e746f6c6f67792e52756e74696d652e47657443757272656e74426c6f636b486173686a00527ac46a00c3681b53797374656d2e426c6f636b636861696e2e4765744865616465726a51527ac46a51c3681d4f6e746f6c6f67792e4865616465722e4765744d65726b6c65526f6f746a52527ac46a00c3681553797374656d2e52756e74696d652e4e6f746966796a52c3681553797374656d2e52756e74696d652e4e6f74696679516c756654c56b68244f6e746f6c6f67792e52756e74696d652e47657443757272656e74426c6f636b486173686a00527ac46a00c3681b53797374656d2e426c6f636b636861696e2e4765744865616465726a51527ac46a51c368204f6e746f6c6f67792e4865616465722e476574436f6e73656e737573446174616a52527ac46a00c3681553797374656d2e52756e74696d652e4e6f746966796a52c3681553797374656d2e52756e74696d652e4e6f74696679516c756654c56b68244f6e746f6c6f67792e52756e74696d652e47657443757272656e74426c6f636b486173686a00527ac46a00c3681b53797374656d2e426c6f636b636861696e2e4765744865616465726a51527ac46a51c368204f6e746f6c6f67792e4865616465722e4765744e657874436f6e73656e7375736a52527ac46a00c3681553797374656d2e52756e74696d652e4e6f746966796a52c3681553797374656d2e52756e74696d652e4e6f74696679516c7566"
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

	//PreExecInvokeContract
	txHash ,err := ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"getVersion", []interface{}{"getVersion"}})
	if err != nil {
		ctx.LogError("TestDomainSmartContract InvokeNeoVMSmartContract error: %s", err)
		return false
	}

	//GetResult
	result, err := txHash.Result.ToByteArray()
	if err != nil {
		ctx.LogError("TestDomainSmartContract GetResult error: %s", err)
		return false
	}

	if hex.EncodeToString(result) != "00" {
		ctx.LogError("TestGetVersion error")
		return false
	}

	return true
}

func TestGetMerkleRoot(ctx *testframework.TestFrameworkContext) bool {

	signer, err := ctx.GetDefaultAccount()
		if err != nil {
			ctx.LogError("TestDomainSmartContract GetDefaultAccount error:%s", err)
			return false
		}

		//InvokeContract
		txHash, err := ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
			signer,
			codeAddress,
			[]interface{}{"getMerkleRoot", []interface{}{[]byte("getMerkleRoot")}})
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

	Hash := events.Notify[0].States.(string)
	count := strings.Count(Hash, "") - 1
	s := []string{}
	for i := count; i > 0; i -= 2 {
		s = append(s, Hash[i-2:i])
	}
	Hash1 := strings.Join(s, "")
	block, err := ctx.Ont.GetBlockByHash(Hash1)
	if err != nil {
		ctx.LogError("GetBlockByHash error:%s", err)
		return false
	}
	txRoot := block.Header.TransactionsRoot.ToHexString()

	merkleRoot := events.Notify[1].States.(string)
	count = strings.Count(merkleRoot, "") - 1
	s = []string{}
	for i := count; i > 0; i -= 2 {
		s = append(s, merkleRoot[i-2:i])
	}
	merkleRoot1 := strings.Join(s, "")

	if txRoot != merkleRoot1 {
		ctx.LogError("TestGetMerkleRoot error")
		return false
	}

	return true
}

func TestGetConsensusData(ctx *testframework.TestFrameworkContext) bool {

	signer, err := ctx.GetDefaultAccount()
	if err != nil {
		ctx.LogError("TestDomainSmartContract GetDefaultAccount error:%s", err)
		return false
	}

	//InvokeContract
	txHash, err := ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"getConsensusData", []interface{}{[]byte("getConsensusData")}})
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

	Hash := events.Notify[0].States.(string)
	count := strings.Count(Hash, "") - 1
	s := []string{}
	for i := count; i > 0; i -= 2 {
		s = append(s, Hash[i-2:i])
	}
	Hash1 := strings.Join(s, "")
	block, err := ctx.Ont.GetBlockByHash(Hash1)
	if err != nil {
		ctx.LogError("GetBlockByHash error:%s", err)
		return false
	}
	cd := block.Header.ConsensusData

	ConsensusData := events.Notify[1].States.(string)
	count = strings.Count(ConsensusData, "") - 1
	s = []string{}
	for i := count; i > 0; i -= 2 {
		s = append(s, ConsensusData[i-2:i])
	}
	ConsensusData1, err := strconv.ParseUint(strings.Join(s, ""), 16, 64)
	if err != nil {
		ctx.LogError("ConsensusDataToUint error:%s", err)
		return false
	}

	if cd != ConsensusData1 {
		ctx.LogError("TestGetMerkleRoot error")
		return false
	}

	return true
}

func TestGetNextConsensus(ctx *testframework.TestFrameworkContext) bool {

	signer, err := ctx.GetDefaultAccount()
	if err != nil {
		ctx.LogError("TestDomainSmartContract GetDefaultAccount error:%s", err)
		return false
	}

	//InvokeContract
	txHash, err := ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"getNextConsensus", []interface{}{[]byte("getNextConsensus")}})
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

	Hash := events.Notify[0].States.(string)
	count := strings.Count(Hash, "") - 1
	s := []string{}
	for i := count; i > 0; i -= 2 {
		s = append(s, Hash[i-2:i])
	}
	Hash1 := strings.Join(s, "")
	block, err := ctx.Ont.GetBlockByHash(Hash1)
	if err != nil {
		ctx.LogError("GetBlockByHash error:%s", err)
		return false
	}
	nc := block.Header.NextBookkeeper.ToHexString()

	NextConsensus := events.Notify[1].States.(string)
	count = strings.Count(NextConsensus, "") - 1
	s = []string{}
	for i := count; i > 0; i -= 2 {
		s = append(s, NextConsensus[i-2:i])
	}
	NextConsensus1 := strings.Join(s, "")

	if nc != NextConsensus1 {
		ctx.LogError("TestGetMerkleRoot error")
		return false
	}

	return true
}
