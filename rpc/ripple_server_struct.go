package rpc

type ServerInfoResp struct {
	Result *ServerInfoResult
}

type ServerInfoResult struct {
	State  ServerInfoState
	Status string
}

type ServerInfoState struct {
	BuildVersion    string                    `json:"build_version"`
	CompleteLedgers string                    `json:"complete_ledgers"`
	ValidatedLedger ServerInfoValidatedLedger `json:"validated_ledger"`
}

type ServerInfoValidatedLedger struct {
	BaseFee     int64 `json:"base_fee"`
	CloseTime   int64 `json:"close_time"`
	Hash        string
	ReserveBase int64 `json:"reserve_base"`
	ReserveInc  int64 `json:"reserve_inc"`
	Seq         uint32
}

type ServerFee struct {
	Result *FeeResult
}
type Drops struct {
	BaseFee       string `json:"base_fee"`
	MedianFee     string `json:"median_fee"`
	MinimumFee    string `json:"minimum_fee"`
	OpenLedgerFee string `json:"open_ledger_fee"`
}
type FeeResult struct {
	CurrentLedgerSize  string `json:"current_ledger_size"`
	CurrentQueueSize   string `json:"current_queue_size"`
	Status             string
	Drops              *Drops
	ExpectedLedgerSize string `json:"expected_ledger_size"`
	LedgerCurrentIndex uint64 `json:"ledger_current_index"`
	MaxQueueSize       string `json:"max_queue_size"`
}
