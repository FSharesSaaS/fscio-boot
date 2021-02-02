package ops

import (
	"github.com/dfuse-io/eosio-boot/config"
	"github.com/FSharesSaaS/fshares.fsgo/ecc"
	"github.com/FSharesSaaS/fshares.fsgo/system"
)

func init() {
	Register("system.setram", &OpSetRAM{})
}

type OpSetRAM struct {
	MaxRAMSize uint64 `json:"max_ram_size"`
}

func (op *OpSetRAM) RequireValidation() bool {
	return true
}

func (op *OpSetRAM) Actions(opPubkey ecc.PublicKey, c *config.OpConfig, in chan interface{}) error {
	in <- (*TransactionAction)(system.NewSetRAM(op.MaxRAMSize))
	in <- EndTransaction(opPubkey) // end transaction
	return nil
}
