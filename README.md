# Ontology Test Framework
Ontology Test Framework is a light-weight test framework for ontology. Integration ontology-sdk to run Ontology test case.

## How to use?

1. Copy wallet file from ontology-test to your test ontology node, the password is 123123.
2. Set rpc server address of ontology, wallet file and password in config_test.json config file.

```
{
  "JsonRpcAddress":"http://localhost:20336",
  "RestfulAddress":"http://localhost:20334",
  "WebSocketAddress":"http://localhost:20335",
  "WalletFile":"./wallet.dat",
  "Password":"your wallet password"
}
```

Then start to run.

Notice:

​	1.If you want to use your own wallet, please change the fromAccount in contract (In testcase\Ontology\Native\native.go) then compiler and copy the new code to function.

​	2.Only test python-compiler.