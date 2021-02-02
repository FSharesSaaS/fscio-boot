package ops

import (
	"github.com/dfuse-io/eosio-boot/config"
	"github.com/FSharesSaaS/fshares.fsgo"
	"github.com/FSharesSaaS/fshares.fsgo/ecc"
	"github.com/FSharesSaaS/fshares.fsgo/token"
)

func init() {
	Register("token.transfer", &OpTransferToken{})
}

type OpTransferToken struct {
	From     eos.AccountName
	To       eos.AccountName
	Quantity eos.Asset
	Memo     string
}

func (op *OpTransferToken) RequireValidation() bool {
	return true
}

func (op *OpTransferToken) Actions(opPubkey ecc.PublicKey, c *config.OpConfig, in chan interface{}) error {
	in <- (*TransactionAction)(token.NewTransfer(op.From, op.To, op.Quantity, op.Memo))
	in <- EndTransaction(opPubkey) // end transaction
	return nil
}
