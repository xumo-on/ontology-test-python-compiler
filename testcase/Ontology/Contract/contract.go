package Contract

import (
	"encoding/hex"
	"github.com/ontio/ontology-go-sdk/utils"
	"github.com/xumo-on/ontology-test-python-compiler/testframework"
	"github.com/ontio/ontology/common"
	"time"
)

var codeAddress common.Address
//Notice:Every time this program is executed, the migrateCode have to be replaced.Any contract undeployed is ok.
var migrateCode = "57c56b6a00527ac46a51527ac46a00c3036c6f679c640900650b006c756661006c756654c56b0e6161614141416161614161616161681253797374656d2e52756e74696d652e4c6f6761516c7566"

func TestGetScript(ctx *testframework.TestFrameworkContext) bool {

	//DeployContract
	code := "55c56b6a00527ac46a51527ac46a00c30f4d696772617465436f6e74726163749c766403006411006a51c300c3652b006c75666203006a00c3096765745363726970749c76640300640c006574006c7566620300006c756655c56b6a00527ac40131013101310131013101316a00c368194f6e746f6c6f67792e436f6e74726163742e4d6967726174656a51527ac46a51c36a52527ac46a51c3641e00144d696772617465207375636365737366756c6c796c7566620700006c75666c756654c56b682d53797374656d2e457865637574696f6e456e67696e652e476574457865637574696e67536372697074486173686a00527ac46a00c3681d53797374656d2e426c6f636b636861696e2e476574436f6e74726163746a51527ac46a51c3681b4f6e746f6c6f67792e436f6e74726163742e4765745363726970746a52527ac46a52c3681553797374656d2e52756e74696d652e4e6f746966796a52c36c7566"
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
	txHash ,err := ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"getScript", []interface{}{[]byte("getScript")}})
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

	if hex.EncodeToString(result) != "55c56b6a00527ac46a51527ac46a00c30f4d696772617465436f6e74726163749c766403006411006a51c300c3652b006c75666203006a00c3096765745363726970749c76640300640c006574006c7566620300006c756655c56b6a00527ac40131013101310131013101316a00c368194f6e746f6c6f67792e436f6e74726163742e4d6967726174656a51527ac46a51c36a52527ac46a51c3641e00144d696772617465207375636365737366756c6c796c7566620700006c75666c756654c56b682d53797374656d2e457865637574696f6e456e67696e652e476574457865637574696e67536372697074486173686a00527ac46a00c3681d53797374656d2e426c6f636b636861696e2e476574436f6e74726163746a51527ac46a51c3681b4f6e746f6c6f67792e436f6e74726163742e4765745363726970746a52527ac46a52c3681553797374656d2e52756e74696d652e4e6f746966796a52c36c7566" {
		ctx.LogError("TestGetScript error")
		return false
	}

	return true
}

func TestMigrate(ctx *testframework.TestFrameworkContext) bool {

	//PreExecInvokeContract
	txHash ,err := ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"MigrateContract", []interface{}{migrateCode}})
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

	if hex.EncodeToString(result) != "4d696772617465207375636365737366756c6c79" {
		ctx.LogError("TestGetContext error")
		return false
	}

	return true
}
