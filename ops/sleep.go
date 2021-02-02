package ops

import (
	"fmt"
	"time"

	"github.com/dfuse-io/eosio-boot/config"
	"github.com/FSharesSaaS/fshares.fsgo/ecc"
)

func init() {
	Register("sleep", &OpSleep{})
}

type OpSleep struct {
	Duration string `json:"duration"`
}

func (op *OpSleep) RequireValidation() bool {
	return true
}

func (op *OpSleep) Actions(opPubkey ecc.PublicKey, c *config.OpConfig, in chan interface{}) error {
	duration, err := time.ParseDuration(op.Duration)
	if err != nil {
		return fmt.Errorf("invalid format for sleep operation duration: %w", err)
	}

	time.Sleep(duration)
	return nil
}
