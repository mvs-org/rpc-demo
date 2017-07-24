package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/antonholmquist/jason"
	"strconv"
	"strings"
)

type User struct {
	Account string
	Auth    string
}

type Rpc struct {
	Admin  *User
	User   *User
	Group  map[string]*User
	Client *Client
}

func SwitchUser(rpc *Rpc, account string) error {
	if user, ok := rpc.Group[account]; ok {
		rpc.User = user
		return nil
	} else {
		return errors.New("account not found")
	}
}

func (r *Rpc) Stop() (string, *ResponseError) {
	resp, ex := r.Client.Request("stop", Params{r.User.Account, r.User.Auth})
	if ex != nil {
		return "", ex
	}
	return string(resp), nil
}

func (r *Rpc) StopAll() (string, *ResponseError) {
	resp, ex := r.Client.Request("stopall", Params{r.Admin.Account, r.User.Auth})
	if ex != nil {
		return "", ex
	}
	return string(resp), nil
}

func (r *Rpc) Start() (string, *ResponseError) {
	resp, ex := r.Client.Request("start", Params{r.User.Account, r.User.Auth})
	if ex != nil {
		return "", ex
	}
	return string(resp), nil
}

func (r *Rpc) GetPeerInfo() ([]string, *ResponseError) {
	resp, ex := r.Client.Request("getpeerinfo", Params{r.User.Account, r.User.Auth})
	if ex != nil {
		return nil, ex
	}
	jsonResp, _ := jason.NewObjectFromBytes(resp)
	peers, _ := jsonResp.GetStringArray("peers")
	return peers, nil
}

func (r *Rpc) GetMiningInfo() (*MiningInfo, *ResponseError) {
	resp, ex := r.Client.Request("getmininginfo", Params{r.User.Account, r.User.Auth})
	if ex != nil {
		return nil, ex
	}
	var mininginfo struct {
		Info *MiningInfo `json:"mining-info"`
	}
	if err := json.Unmarshal(resp, &mininginfo); err != nil {
		ex.Message = err.Error()
		return nil, ex
	}
	return mininginfo.Info, nil
}

func (r *Rpc) GetBestBlockHeader() (*BlockHeader, *ResponseError) {
	resp, ex := r.Client.Request("getbestblockheader", Params{})
	if ex != nil {
		return nil, ex
	}
	var result struct {
		BlockHeader BlockHeader `json:"result"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		ex.Message = err.Error()
		return nil, ex
	}
	return &result.BlockHeader, nil
}

func (r *Rpc) FetchHeaderExt(blocknum uint64) (*BlockHeader, *ResponseError) {
	resp, ex := r.Client.Request("fetchheaderext", Params{r.User.Account, r.User.Auth, blocknum})
	if ex != nil {
		return nil, ex
	}
	var result struct {
		BlockHeader *BlockHeader `json:"result"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		ex.Message = err.Error()
		return nil, ex
	}
	return result.BlockHeader, nil
}

func (r *Rpc) GetNewAccount() (string, string, *ResponseError) {
	resp, ex := r.Client.Request("getnewaccount", Params{r.User.Account, r.User.Auth})
	if ex != nil {
		return "", "", ex
	}
	var result struct {
		Mnemonic       string `json:"mnemonic"`
		DefaultAddress string `json:"default-address"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		ex.Message = err.Error()
		return "", "", ex
	}
	return result.Mnemonic, result.DefaultAddress, nil
}

func (r *Rpc) ImportAccount(mnemonic []string, hdIdx uint64) (*AccountInfo, *ResponseError) {
	ex := new(ResponseError)
	if hdIdx == 0 {
		ex.Message = "hd index should > 0"
		return nil, ex
	}
	params := Params{
		strings.Join(mnemonic, " "),
		fmt.Sprintf("--accoutname=%v", r.User.Account),
		fmt.Sprintf("--password=%v", r.User.Auth),
		fmt.Sprintf("--hd_index=%v", hdIdx),
	}
	resp, ex := r.Client.Request("importaccount", params)
	if ex != nil {
		return nil, ex
	}
	accInfo := new(AccountInfo)
	if err := json.Unmarshal(resp, accInfo); err != nil {
		fmt.Println(err)
		ex.Message = err.Error()
		return nil, ex
	}
	return accInfo, nil
}

func (r *Rpc) GetAccount(lastword string) (*AccountInfo, *ResponseError) {
	resp, ex := r.Client.Request("getaccount", Params{r.User.Account, r.User.Auth, lastword})
	if ex != nil {
		return nil, ex
	}
	var result struct {
		Name       string `json:"name"`
		Mnemonic   string `json:"mnemonic-key"`
		AddressCnt int64  `json:"address-count,string,omitempty"`
		UserStatus int    `json:"user-status,string,omitempty"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		ex.Message = err.Error()
		return nil, ex
	}
	accInfo := &AccountInfo{
		Name:       result.Name,
		Mnemonic:   result.Mnemonic,
		AddressCnt: result.AddressCnt,
		UserStatus: result.UserStatus,
	}
	return accInfo, ex
}

func (r *Rpc) DeleteAccount(lastword string) (string, *ResponseError) {
	resp, ex := r.Client.Request("deleteaccount", Params{r.User.Account, r.User.Auth, lastword})
	if ex != nil {
		return "", ex
	}
	return string(resp), nil
}

func (r *Rpc) ListAddresses() ([]string, *ResponseError) {
	resp, ex := r.Client.Request("listaddresses", Params{r.User.Account, r.User.Auth})
	if ex != nil {
		return nil, ex
	}
	var result struct {
		Addresses []string `json:"addresses"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		ex.Message = err.Error()
		return nil, ex
	}
	return result.Addresses, nil
}

func (r *Rpc) GetNewAddress() (string, *ResponseError) {
	resp, ex := r.Client.Request("getnewaddress", Params{r.User.Account, r.User.Auth})
	if ex != nil {
		return "", ex
	}
	return string(resp), nil
}

func (r *Rpc) GetPublicKey(optaddr string) (map[string]string, *ResponseError) {
	params := []interface{}{r.User.Account, r.User.Auth}
	if optaddr != "" {
		params = append(params, optaddr)
	}
	resp, ex := r.Client.Request("getpublickey", Params(params))
	if ex != nil {
		return nil, ex
	}
	var result struct {
		PubKey  string `json:"public-key"`
		Address string `json:"address"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		ex.Message = err.Error()
		return nil, ex
	}
	return map[string]string{result.Address: result.PubKey}, nil
}

func (r *Rpc) ValidateAddress(address string) (string, *ResponseError) {
	resp, ex := r.Client.Request("validateaddress", Params{r.User.Account, r.User.Auth, address})
	if ex != nil {
		return "", ex
	}
	return string(resp), nil
}

func (r *Rpc) GetBalance() (*BalanceStatistic, *ResponseError) {
	resp, ex := r.Client.Request("getbalance", Params{r.User.Account, r.User.Auth})
	if ex != nil {
		return nil, ex
	}
	bs := new(BalanceStatistic)
	if err := json.Unmarshal(resp, bs); err != nil {
		ex.Message = err.Error()
		return nil, ex
	}
	return bs, nil
}

func (r *Rpc) XFetchBalance(address string, assetType ...string) (*Balance, *ResponseError) {
	params := []interface{}{address}
	if assetType != nil && len(assetType) > 0 && assetType[0] == "etp" {
		params = append(params, fmt.Sprintf("--type=%v", assetType[0]))
	}
	resp, ex := r.Client.Request("xfetchbalance", Params(params))
	if ex != nil {
		return nil, ex
	}
	var result struct {
		Balance *Balance `json:"balance"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		ex.Message = err.Error()
		return nil, ex
	}
	return result.Balance, nil
}

func (r *Rpc) XFetchUTXO(amount uint64, address string, assetType ...string) (*UTXO, *ResponseError) {
	params := []interface{}{amount, address}
	if assetType != nil && len(assetType) > 0 && assetType[0] == "etp" {
		params = append(params, fmt.Sprintf("--type=%v", assetType[0]))
	}
	resp, ex := r.Client.Request("xfetchutxo", Params(params))
	if ex != nil {
		return nil, ex
	}
	utxo := new(UTXO)
	if err := json.Unmarshal(resp, utxo); err != nil {
		ex.Message = err.Error()
		return nil, ex
	}
	return utxo, nil
}

func (r *Rpc) ListBalances() ([]*BalanceSet, *ResponseError) {
	resp, ex := r.Client.Request("listbalances", Params{r.User.Account, r.User.Auth})
	if ex != nil {
		return nil, ex
	}
	var result struct {
		Balances []*BalanceSet `json:"balances"`
	}

	if err := json.Unmarshal(resp, &result); err != nil {
		ex.Message = err.Error()
		return nil, ex
	}
	return result.Balances, nil
}

func (r *Rpc) GetTransaction(txhash string) (*Transaction, *ResponseError) {
	resp, ex := r.Client.Request("gettransaction", Params{txhash})
	if ex != nil {
		return nil, ex
	}
	tx := new(Transaction)
	if err := json.Unmarshal(resp, tx); err != nil {
		fmt.Println(err)
		ex.Message = err.Error()
		return nil, ex
	}
	return tx, nil
}

func (r *Rpc) ListTxs(options map[string]interface{}) ([]*Transaction, *ResponseError) {
	params := []interface{}{r.User.Account, r.User.Auth}
	for k, v := range options {
		switch k {
		case "address":
			params = append(params, fmt.Sprintf("--address=%v", v))
		case "a":
			params = append(params, fmt.Sprintf("-a %v", v))
		case "height":
			params = append(params, fmt.Sprintf("--height=%v", v))
		case "e":
			params = append(params, fmt.Sprintf("-e %v", v))
		}
	}
	resp, ex := r.Client.Request("listtxs", Params(params))
	if ex != nil {
		return nil, ex
	}
	var result struct {
		Transactions []*Transaction `json:"transactions"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		ex.Message = err.Error()
		return nil, ex
	}
	return result.Transactions, nil
}

func (r *Rpc) GetBestblockhash() (string, *ResponseError) {
	resp, ex := r.Client.Request("getbestblockhash", Params{})
	if ex != nil {
		return "", ex
	}
	return string(resp), nil
}

func (r *Rpc) Send(address string, amount uint64) (*Transaction, *ResponseError) {
	resp, ex := r.Client.Request("send", Params{r.User.Account, r.User.Auth, address, amount})
	if ex != nil {
		return nil, ex
	}
	var result struct {
		Tx *Transaction `json:"transaction"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		ex.Message = err.Error()
		return nil, ex
	}
	return result.Tx, nil
}

func (r *Rpc) SendFrom(fromaddr, toaddr string, amount int64) (*Transaction, *ResponseError) {
	resp, ex := r.Client.Request("sendfrom", Params{r.User.Account, r.User.Auth, fromaddr, toaddr, amount})
	if ex != nil {
		return nil, ex
	}
	var result struct {
		Tx *Transaction `json:"transaction"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		ex.Message = err.Error()
		return nil, ex
	}
	return result.Tx, nil
}

func (r *Rpc) SendMore(receivers map[string]uint64, fee float64, mychangeAddr string) ([]*Transaction, *ResponseError) {
	params := []interface{}{r.User.Account, r.User.Auth}
	for addr, amount := range receivers {
		params = append(params, fmt.Sprintf("-r %v:%v ", addr, amount))
	}
	if fee > 0 {
		params = append(params, fee)
	}
	if mychangeAddr != "" {
		params = append(params, mychangeAddr)
	}
	resp, ex := r.Client.Request("sendmore", Params(params))
	if ex != nil {
		return nil, ex
	}
	txs := make([]*Transaction, 0)
	if err := json.Unmarshal(resp, txs); err != nil {
		ex.Message = err.Error()
		return nil, ex
	}
	return txs, nil

}

func (r *Rpc) FetchHeight() (uint64, *ResponseError) {
	resp, ex := r.Client.Request("fetch-height", Params{})
	if ex != nil {
		return 0, ex
	}

	height, err := strconv.ParseUint(strings.TrimRight(string(resp), "\n"), 10, 64)
	if err != nil {
		fmt.Println(err)
		ex.Message = err.Error()
		return 0, ex
	}
	return height, nil
}

func (r *Rpc) FetchHeader(options map[string]interface{}) (*BlockHeader, *ResponseError) {
	params := make([]interface{}, 0)
	for k, v := range options {
		switch k {
		case "hash":
			params = append(params, fmt.Sprintf("--hash=%v", v))
		case "s":
			params = append(params, fmt.Sprintf("-s %v", v))
		case "height":
			params = append(params, fmt.Sprintf("--height=%v", v))
		case "t":
			params = append(params, fmt.Sprintf("-t %v", v))
		}
	}

	resp, ex := r.Client.Request("fetch-header", Params(params))
	if ex != nil {
		return nil, ex
	}
	var result struct {
		Block *BlockHeader `json:"result"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		ex.Message = err.Error()
		return nil, ex
	}
	return result.Block, nil
}


