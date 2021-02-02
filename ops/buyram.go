package ops

import (
	"github.com/dfuse-io/eosio-boot/config"
	"github.com/FSharesSaaS/fshares.fsgo"
	"github.com/FSharesSaaS/fshares.fsgo/ecc"
	"github.com/FSharesSaaS/fshares.fsgo/system"
)

func init() {
	Register("system.buy_ram", &OpBuyRam{})
}

type OpBuyRam struct {
	Payer       eos.AccountName
	Receiver    eos.AccountName
	EOSQuantity uint64 `json:"eos_quantity"`
}

func (op *OpBuyRam) Actions(opPubkey ecc.PublicKey, c *config.OpConfig, in chan interface{}) error {
	in <- (*TransactionAction)(system.NewBuyRAM(op.Payer, op.Receiver, op.EOSQuantity))
	in <- EndTransaction(opPubkey) // end transaction
	return nil

}

func (op *OpBuyRam) RequireValidation() bool {
	return true
}
