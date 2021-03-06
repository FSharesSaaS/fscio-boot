package ops

import (
	"fmt"

	"github.com/dfuse-io/eosio-boot/config"
	"github.com/FSharesSaaS/fshares.fsgo"
	"github.com/FSharesSaaS/fshares.fsgo/ecc"
	"github.com/FSharesSaaS/fshares.fsgo/system"
)

func init() {
	Register("system.setcode", &OpSetCode{})
}

type OpSetCode struct {
	Account         eos.AccountName
	ContractNameRef string `json:"contract_name_ref"`
}

func (op *OpSetCode) RequireValidation() bool {
	return true
}

func (op *OpSetCode) Actions(opPubkey ecc.PublicKey, c *config.OpConfig, in chan interface{}) error {
	wasmFileRef, err := c.GetContentsCacheRef(fmt.Sprintf("%s.wasm", op.ContractNameRef))
	if err != nil {
		return err
	}

	abiFileRef, err := c.GetContentsCacheRef(fmt.Sprintf("%s.abi", op.ContractNameRef))
	if err != nil {
		return err
	}

	codeAction, err := system.NewSetCode(op.Account, c.FileNameFromCache(wasmFileRef))
	if err != nil {
		return fmt.Errorf("NewSetCode %s: %s", op.ContractNameRef, err)
	}

	abi, err := retrieveABIfromRef(c.FileNameFromCache(abiFileRef))
	if err != nil {
		return fmt.Errorf("unable to read ABI %s: %s", abiFileRef, err)
	}

	abiAction, err := system.NewSetAbiFromAbi(op.Account, *abi)
	if err != nil {
		return fmt.Errorf("NewSetAbiFromAbi %s: %s", op.ContractNameRef, err)
	}

	c.AbiCache.SetABI(op.Account, abi)
	for _, act := range []*eos.Action{codeAction, abiAction} {
		in <- (*TransactionAction)(act)
	}

	in <- EndTransaction(opPubkey) // end transaction
	return nil
}
