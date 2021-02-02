package ops

import (
	"bytes"
	"fmt"

	"github.com/dfuse-io/eosio-boot/config"
	"github.com/FSharesSaaS/fshares.fsgo"
	"github.com/FSharesSaaS/fshares.fsgo/ecc"
	"github.com/FSharesSaaS/fshares.fsgo/system"
	"github.com/FSharesSaaS/fshares.fsgo/token"
)

func init() {
	Register("system.create_voters", &OpCreateVoters{})
}

type OpCreateVoters struct {
	Creator eos.AccountName
	Pubkey  string
	Count   int
}

func (op *OpCreateVoters) RequireValidation() bool {
	return true
}

func (op *OpCreateVoters) Actions(opPubkey ecc.PublicKey, c *config.OpConfig, in chan interface{}) error {
	pubKey, err := decodeOpPublicKey(c, op.Pubkey)
	if err != nil {
		return err
	}

	for i := 0; i < op.Count; i++ {
		voterName := eos.AccountName(voterName(i))
		fmt.Println("Creating voter: ", voterName)

		in <- (*TransactionAction)(system.NewNewAccount(op.Creator, voterName, pubKey))
		in <- (*TransactionAction)(token.NewTransfer(op.Creator, voterName, eos.NewEOSAsset(1000000000), ""))
		in <- (*TransactionAction)(system.NewBuyRAMBytes(AN("eosio"), voterName, 8192)) // 8kb gift !
		in <- (*TransactionAction)(system.NewDelegateBW(AN("eosio"), voterName, eos.NewEOSAsset(10000), eos.NewEOSAsset(10000), true))
	}
	in <- EndTransaction(opPubkey) // end transaction
	return nil
}

const charset = "abcdefghijklmnopqrstuvwxyz"

func voterName(index int) string {
	padding := string(bytes.Repeat([]byte{charset[index]}, 7))
	return "voter" + padding
}
