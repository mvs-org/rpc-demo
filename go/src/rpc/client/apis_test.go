package client

import (
	"strings"
	"testing"
)

var rpc *Rpc

func init() {
	cli := NewClient("http://127.0.0.1:8820", "/rpc")
	user := &User{Account: "mutou", Auth: "mutou"}
	group := map[string]*User{"mutou": user}
	rpc = &Rpc{User: user, Client: cli, Group: group}
}

func TestRpc_GetBestblockhash(t *testing.T) {
	hash, ex := rpc.GetBestblockhash()
	if ex != nil {
		t.Fatal(ex.Message)
	}
	t.Log(hash)
}

func TestRpc_GetBestBlockHeader(t *testing.T) {
	blockHeader, ex := rpc.GetBestBlockHeader()
	if ex != nil {
		t.Fatal(ex.Message)
	}
	t.Log(blockHeader)
}

func TestRpc_FetchHeader(t *testing.T) {
	options := make(map[string]interface{}, 0)
	options["hash"] = "abc0b70b271be9249c57e724667c7ce1330e7454c3697336a6b38aa0cf1516dc"
	options["height"] = "30000"
	b, ex := rpc.FetchHeader(options)
	if ex != nil {
		t.Fatal(ex.Message)
	}
	t.Log(b)
}

func TestRpc_FetchHeaderExt(t *testing.T) {
	b, ex := rpc.FetchHeaderExt(30000)
	if ex != nil {
		t.Fatal(ex.Message)
	}
	t.Log(b)
}

func TestRpc_FetchHeight(t *testing.T) {
	height, ex := rpc.FetchHeight()
	if ex != nil {
		t.Fatal(ex.Message)
	}
	t.Log(height)
}

func TestRpc_GetMiningInfo(t *testing.T) {
	info, ex := rpc.GetMiningInfo()
	if ex != nil {
		t.Fatal(ex.Message)
	}
	t.Log(info)
}

func TestRpc_GetPeerInfo(t *testing.T) {
	info, ex := rpc.GetPeerInfo()
	if ex != nil {
		t.Fatal(ex.Message)
	}
	t.Log(info)
}

func TestRpc_Start(t *testing.T) {
	result, ex := rpc.Start()
	if ex != nil {
		t.Fatal(ex.Message)
	}
	t.Log(result)
}

func TestRpc_Stop(t *testing.T) {
	result, ex := rpc.Stop()
	if ex != nil {
		t.Fatal(ex.Message)
	}
	t.Log(result)
}

func TestRpc_StopAll(t *testing.T) {
	admin := &User{Account: "admin", Auth: "admin"}
	rpc.Admin = admin
	result, ex := rpc.StopAll()
	if ex != nil {
		t.Fatal(ex.Message)
	}
	t.Log(result)
}

func TestRpc_GetNewAccount(t *testing.T) {
	newUser := &User{Account: "demo0612", Auth: "demo0610"}
	rpc.Group["demo0612"] = newUser
	SwitchUser(rpc, "demo0612")
	words, addr, ex := rpc.GetNewAccount()
	if ex != nil {
		t.Fatal(ex)
	}
	t.Log(words, "\n", addr)
}

func TestRpc_DeleteAccount(t *testing.T) {
	SwitchUser(rpc, "demo0611")
	result, ex := rpc.DeleteAccount("parade")
	if ex != nil {
		t.Fatal(ex.Message)
	}
	t.Log(result)
}

func TestRpc_ImportAccount(t *testing.T) {
	rpc.User = &User{"mutou", "mutou"}
	words := "decrease tank copper object melt twelve axis amateur half tornado snake huge estate need since giggle genius oil tiny daughter alcohol lab funny audit"
	mnemonic := strings.Split(words, " ")
	info, ex := rpc.ImportAccount(mnemonic, 1)
	if ex != nil {
		t.Fatal(ex.Message)
	}
	t.Log(info)
}

func TestRpc_GetNewAddress(t *testing.T) {
	rpc.User = &User{"mutou1234567", "mutou1234567"}
	addr, ex := rpc.GetNewAddress()
	if ex != nil {
		t.Fatal(ex.Message)
	}
	t.Log(addr)
}

func TestRpc_GetAccount(t *testing.T) {
	rpc.User = &User{"mutou", "mutou"}
	info, ex := rpc.GetAccount("audit")
	if ex != nil {
		t.Fatal(ex.Message)
	}
	t.Log(info)
}

func TestRpc_GetBalance(t *testing.T) {
	result, ex := rpc.GetBalance()
	if ex != nil {
		t.Fatal(ex.Message)
	}
	t.Log(result)
	t.Log(result.TotalAvailable)
}

func TestRpc_GetTransaction(t *testing.T) {
	tx, ex := rpc.GetTransaction("efff1560ef1f60af97284cf45db3444c74746c80d94634a6f12a7111e681de31")
	if ex != nil {
		t.Fatal(ex.Message)
	}
	t.Log(tx)
}

func TestRpc_ListTxs(t *testing.T) {
	//options1 := map[string]interface{}{"address": "MSx1uiBk9vRnW9mcbNgmRP8Fn1svcGkNfR", "height": "231:43827428"}
	//options2 := map[string]interface{}{"height": "438:4393123"}
	options3 := map[string]interface{}{"address": "tMhZMdz6Gk7RARMW1WxFvLfoodKG82CEKT"}

	txs, ex := rpc.ListTxs(options3)
	//txs, ex := rpc.ListTxs(nil)
	if ex != nil {
		t.Fatal(ex.Message)
	}
	t.Log(txs)
	t.Log(txs[1].Hash, txs[0].Inputs[0].Address)
	t.Log(txs[1].Hash)
}

func TestRpc_ListAddresses(t *testing.T) {
	addrs, ex := rpc.ListAddresses()
	if ex != nil {
		t.Fatal(ex.Message)
	}
	t.Log(addrs)
}

func TestRpc_ListBalances(t *testing.T) {
	bs, ex := rpc.ListBalances()
	if ex != nil {
		t.Fatal(ex.Message)
	}
	t.Log(bs)
	t.Log(bs[0].Balance.Address)
}

func TestRpc_ValidateAddress(t *testing.T) {
	result, ex := rpc.ValidateAddress("MSx1uiBk9vRnW9mcbNgmRP8Fn1svcGkNfR")
	if ex != nil {
		t.Fatal(ex.Message)
	}
	t.Log(result)
}

func TestRpc_XFetchBalance(t *testing.T) {
	bs, ex := rpc.XFetchBalance("MSx1uiBk9vRnW9mcbNgmRP8Fn1svcGkNfR", "etp")
	if ex != nil {
		t.Fatal(ex.Message)
	}
	t.Log(bs)
}

func TestRpc_XFetchUTXO(t *testing.T) {
	utxos, ex := rpc.XFetchUTXO(100, "MSx1uiBk9vRnW9mcbNgmRP8Fn1svcGkNfR", "etp")
	if ex != nil {
		t.Fatal(ex.Message)
	}
	t.Log(utxos)
}

func TestRpc_GetPublicKey(t *testing.T) {
	pairs, ex := rpc.GetPublicKey("MSx1uiBk9vRnW9mcbNgmRP8Fn1svcGkNfR")
	if ex != nil {
		t.Fatal(ex)
	}
	for k, v := range pairs {
		t.Log("addr:pubkey | ", k, ":", v)
	}
}

func TestRpc_Send(t *testing.T) {
	tx, ex := rpc.Send("tEDsZ5ob2MPNLmZQifevijY1CMq4bw8pH2", 1)
	if ex != nil {
		t.Fatal(ex.Message)
	}
	t.Log(tx)
}

func TestRpc_SendFrom(t *testing.T) {
	tx, ex := rpc.SendFrom("MSx1uiBk9vRnW9mcbNgmRP8Fn1svcGkNfR", "3Pf24QBBdL7cLk27BpQeAaFYZqmDUc7ppb", 1)
	if ex != nil {
		t.Fatal(ex.Message)
	}
	t.Log(tx)
}

func TestRpc_SendMore(t *testing.T) {
	r := make(map[string]uint64, 0)
	r["MSx1uiBk9vRnW9mcbNgmRP8Fn1svcGkNfR"] = 1
	txs, ex := rpc.SendMore(r, 0.0001, "MAmHAQd9hUmHG7EbnUvQ1pg2s3Tyd2tdjf")
	if ex != nil {
		t.Fatal(ex.Message)
	}
	t.Log(txs)
}

func TestRpc_GetBlock(t *testing.T) {
	block, ex := rpc.GetBlock("abc0b70b271be9249c57e724667c7ce1330e7454c3697336a6b38aa0cf1516dc", true)
	if ex != nil {
		t.Fatal(ex.Message)
	}
	t.Log(block.Header.Hash)
	t.Log(block.Txs[0].Hash)
}

func TestRpc_CreateAsset(t *testing.T) {
	options := make(map[string]interface{}, 0)
	options["issuer"] = "tester"
	options["description"] = "we are testing"
	options["decimalnumber"] = 2
	asset, ex := rpc.CreateAsset("TEST.TPC", 2000, options)
	if ex != nil {
		t.Fatal(ex.Message)
	}
	t.Log(asset)
	t.Log(asset.Symbol)
}

func TestRpc_GetAsset(t *testing.T) {
	assets, ex := rpc.GetAsset("TEST.LANJIAN")
	if ex != nil {
		t.Fatal(ex.Message)
	}
	t.Log(assets)
	t.Log(assets[0])
}

func TestRpc_GetAddressAsset(t *testing.T) {
	assets, ex := rpc.GetAddressAsset("tTxNwtVgwyguysyNcWFFgGMKzfhwYqUy3m")
	if ex != nil {
		t.Fatal(ex.Message)
	}
	t.Log(assets)
	t.Log(assets[0])
}

func TestRpc_GetAccountAsset(t *testing.T) {
	assets, ex := rpc.GetAccountAsset()
	if ex != nil {
		t.Fatal(ex.Message)
	}
	t.Log(assets)
	t.Log(assets[1])
}

func TestRpc_ListAssets(t *testing.T) {
	assets, ex := rpc.ListAssets()
	if ex != nil {
		t.Fatal(ex.Message)
	}
	t.Log(assets)
	t.Log(assets[0])
}

func TestRpc_Issue(t *testing.T) {
	tx, ex := rpc.Issue("TEST.TPC")
	if ex != nil {
		t.Fatal(ex.Message)
	}
	t.Log(tx)
	t.Log(tx.Outputs[0].Attachment.Symbol)
}

func TestRpc_IssueFrom(t *testing.T) {
	tx, ex := rpc.IssueFrom("TEST.TPC", "tMhZMdz6Gk7RARMW1WxFvLfoodKG82CEKT")
	if ex != nil {
		t.Fatal(ex.Message)
	}
	t.Log(tx)
}

func TestRpc_SendAsset(t *testing.T) {
	tx, ex := rpc.SendAsset("TEST.TPC", "tEDsZ5ob2MPNLmZQifevijY1CMq4bw8pH2", 200)
	if ex != nil {
		t.Fatal(ex.Message)
	}
	t.Log(tx)
	for _, v := range tx.Outputs {
		t.Log(v)
	}

	for _, v := range tx.Inputs {
		t.Log(v)
	}
}

func TestRpc_SendAssetFrom(t *testing.T) {
	tx, ex := rpc.SendAssetFrom("TEST.LANJIAN", "tMhZMdz6Gk7RARMW1WxFvLfoodKG82CEKT", "tEDsZ5ob2MPNLmZQifevijY1CMq4bw8pH2", 300)
	if ex != nil {
		t.Fatal(ex.Message)
	}
	t.Log(tx)
	for _, v := range tx.Outputs {
		t.Log(v)
	}

	for _, v := range tx.Inputs {
		t.Log(v)
	}
}

func TestRpc_Deposit(t *testing.T) {
	tx, ex := rpc.Deposit(200)
	if ex != nil {
		t.Fatal(ex.Message)
	}
	t.Log(tx)
	for _, v := range tx.Outputs {
		t.Log(v)
	}

	for _, v := range tx.Inputs {
		t.Log(v)
	}
}

/*
 tx := `{
    "transaction": {
        "hash": "af9a1ddf4ad11826bcde7d7969c86c96b74d897ceef18c49c0d553595849df1d",
        "inputs": [
            {
                "address": "tMhZMdz6Gk7RARMW1WxFvLfoodKG82CEKT",
                "previous_output": {
                    "hash": "c920d4c2eca699fdfe6b78ac250372002ef97a08028813e57ccfd774e5d5f33f",
                    "index": "1"
                },
                "script": "[ 30440220654bfd0cffea9ac4b5dfec0e492685d41e28b66b6da91a72b4cf57dd30e27e73022039b914b4c2c5517d75fd36b21bf3302b068e7c1294f793ff7403aa40e25ea82301 ] [ 022c40429b64080245118286d5f951169915ddd66b6188135ec67455a90b548a0e ]",
                "sequence": "4294967295"
            }
        ],
        "lock_time": "0",
        "outputs": [
            {
                "index": "0",
                "address": "tMhZMdz6Gk7RARMW1WxFvLfoodKG82CEKT",
                "script": "dup hash160 [ a2123b59f0b22170f3cdf4ebc7452af7da44c16c ] equalverify checksig",
                "value": "0",
                "attachment": {
                    "type": "asset-issue",
                    "symbol": "TEST.TPC",
                    "quantity": "2000",
                    "decimal_number": "2",
                    "issuer": "mutou",
                    "address": "tMhZMdz6Gk7RARMW1WxFvLfoodKG82CEKT",
                    "description": "we are testing"
                }
            },
            {
                "index": "1",
                "address": "tMhZMdz6Gk7RARMW1WxFvLfoodKG82CEKT",
                "script": "dup hash160 [ a2123b59f0b22170f3cdf4ebc7452af7da44c16c ] equalverify checksig",
                "value": "28000000000",
                "attachment": {
                    "type": "etp"
                }
            }
        ],
        "version": "2"
    }
}`
*/
