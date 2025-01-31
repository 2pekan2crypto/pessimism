package core

import (
	"math/big"
	"time"
)

// OracleConfig ... Configuration passed through to an oracle component constructor
type OracleConfig struct {
	RPCEndpoint  string
	PollInterval time.Duration
	NumOfRetries int
	StartHeight  *big.Int
	EndHeight    *big.Int
}

// PipelineConfig ... Configuration passed through to a pipeline constructor
type PipelineConfig struct {
	Network      Network
	DataType     RegisterType
	PipelineType PipelineType
	OracleCfg    *OracleConfig
}

// Backfill ... Returns true if the oracle is configured to backfill
func (oc *OracleConfig) Backfill() bool {
	return oc.StartHeight != nil
}

// Backtest ... Returns true if the oracle is configured to backtest
func (oc *OracleConfig) Backtest() bool {
	return oc.EndHeight != nil
}
