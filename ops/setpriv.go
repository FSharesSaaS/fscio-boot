package ops

import (
	"github.com/dfuse-io/eosio-boot/config"
	"github.com/FSharesSaaS/fshares.fsgo"
	"github.com/FSharesSaaS/fshares.fsgo/ecc"
	"github.com/FSharesSaaS/fshares.fsgo/system"
)

func init() {
	Register("system.setpriv", &OpSetPriv{})
}

type OpSetPriv struct {
	Account eos.AccountName
}

func (op *OpSetPriv) RequireValidation() bool {
	return true
}

func (op *OpSetPriv) Actions(opPubkey ecc.PublicKey, c *config.OpConfig, in chan interface{}) error {
	in <- (*TransactionAction)(system.NewSetPriv(op.Account))
	in <- EndTransaction(opPubkey) // end transaction
	return nil

}
