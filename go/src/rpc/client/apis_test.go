package client

import (
	"strings"
	"testing"
)

var rpc *Rpc

func init() {
	cli := NewClient("http://127.0.0.1:8820", "/rpc")
	user := &User{Account: "impressiver", Auth: "lanjianhao"}
	group := map[string]*User{"impressiver": user}
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
	rpc.User = &User{"mutou111", "mutou111"}
	words := "apology network ostrich dentist denial vehicle forget buzz about record barely course check clutch area game pioneer sort casino letter isolate style maximum father"
	mnemonic := strings.Split(words, " ")
	info, ex := rpc.ImportAccount(mnemonic, 0)
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
	rpc.User = &User{"mutou1234567", "mutou1234567"}
	info, ex := rpc.GetAccount("father")
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
	options3 := map[string]interface{}{"address": "3Pf24QBBdL7cLk27BpQeAaFYZqmDUc7ppb"}

	txs, ex := rpc.ListTxs(options3)
	//txs, ex := rpc.ListTxs(nil)
	if ex != nil {
		t.Fatal(ex.Message)
	}
	t.Log(txs)
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
	t.Log(bs[1].Balance.Address)
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
	tx, ex := rpc.Send("MSx1uiBk9vRnW9mcbNgmRP8Fn1svcGkNfR", 1)
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
