package ops

import (
	"github.com/dfuse-io/eosio-boot/config"
	"github.com/FSharesSaaS/fshares.fsgo"
	"github.com/FSharesSaaS/fshares.fsgo/ecc"
	"github.com/FSharesSaaS/fshares.fsgo/system"
)

func init() {
	Register("system.delegate_bw", &OpDelegateBW{})
}

type OpDelegateBW struct {
	From     eos.AccountName
	To       eos.AccountName
	StakeCPU int64 `json:"stake_cpu"`
	StakeNet int64 `json:"stake_net"`
	Transfer bool
}

func (op *OpDelegateBW) Actions(opPubkey ecc.PublicKey, c *config.OpConfig, in chan interface{}) error {
	in <- (*TransactionAction)(system.NewDelegateBW(op.From, op.To, eos.NewEOSAsset(op.StakeCPU), eos.NewEOSAsset(op.StakeNet), op.Transfer))
	in <- EndTransaction(opPubkey) // end transaction
	return nil
}

func (op *OpDelegateBW) RequireValidation() bool {
	return true
}
