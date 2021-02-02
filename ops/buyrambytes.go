package ops

import (
	"github.com/dfuse-io/eosio-boot/config"
	"github.com/FSharesSaaS/fshares.fsgo"
	"github.com/FSharesSaaS/fshares.fsgo/ecc"
	"github.com/FSharesSaaS/fshares.fsgo/system"
)

func init() {
	Register("system.buy_ram_bytes", &OpBuyRamBytes{})
}

type OpBuyRamBytes struct {
	Payer    eos.AccountName
	Receiver eos.AccountName
	Bytes    uint32
}

func (op *OpBuyRamBytes) Actions(opPubkey ecc.PublicKey, c *config.OpConfig, in chan interface{}) error {
	in <- (*TransactionAction)(system.NewBuyRAMBytes(op.Payer, op.Receiver, op.Bytes))
	in <- EndTransaction(opPubkey) // end transaction
	return nil

}

func (op *OpBuyRamBytes) RequireValidation() bool {
	return true
}
