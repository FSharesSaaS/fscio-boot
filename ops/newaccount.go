package ops

import (
	"github.com/dfuse-io/eosio-boot/config"
	"github.com/FSharesSaaS/fshares.fsgo"
	"github.com/FSharesSaaS/fshares.fsgo/ecc"
	"github.com/FSharesSaaS/fshares.fsgo/system"
)

func init() {
	Register("system.newaccount", &OpNewAccount{})
}

type OpNewAccount struct {
	Creator    eos.AccountName
	NewAccount eos.AccountName `json:"new_account"`
	Pubkey     string
	RamBytes   uint32 `json:"ram_bytes"`
}

func (op *OpNewAccount) RequireValidation() bool {
	return true
}

func (op *OpNewAccount) Actions(opPubkey ecc.PublicKey, c *config.OpConfig, in chan interface{}) error {
	pubKey, err := decodeOpPublicKey(c, op.Pubkey)
	if err != nil {
		return err
	}

	in <- (*TransactionAction)(system.NewNewAccount(op.Creator, op.NewAccount, pubKey))

	if op.RamBytes > 0 {
		in <- (*TransactionAction)(system.NewBuyRAMBytes(op.Creator, op.NewAccount, op.RamBytes))
	}
	in <- EndTransaction(opPubkey) // end transaction
	return nil
}
