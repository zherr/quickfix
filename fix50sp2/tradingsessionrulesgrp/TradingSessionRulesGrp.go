package tradingsessionrulesgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp2/tradingsessionrules"
)

//New returns an initialized TradingSessionRulesGrp instance
func New() *TradingSessionRulesGrp {
	var m TradingSessionRulesGrp
	return &m
}

//NoTradingSessionRules is a repeating group in TradingSessionRulesGrp
type NoTradingSessionRules struct {
	//TradingSessionID is a non-required field for NoTradingSessionRules.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for NoTradingSessionRules.
	TradingSessionSubID *string `fix:"625"`
	//TradingSessionRules is a non-required component for NoTradingSessionRules.
	TradingSessionRules *tradingsessionrules.TradingSessionRules
}

//NewNoTradingSessionRules returns an initialized NoTradingSessionRules instance
func NewNoTradingSessionRules() *NoTradingSessionRules {
	var m NoTradingSessionRules
	return &m
}

func (m *NoTradingSessionRules) SetTradingSessionID(v string)    { m.TradingSessionID = &v }
func (m *NoTradingSessionRules) SetTradingSessionSubID(v string) { m.TradingSessionSubID = &v }
func (m *NoTradingSessionRules) SetTradingSessionRules(v tradingsessionrules.TradingSessionRules) {
	m.TradingSessionRules = &v
}

//TradingSessionRulesGrp is a fix50sp2 Component
type TradingSessionRulesGrp struct {
	//NoTradingSessionRules is a non-required field for TradingSessionRulesGrp.
	NoTradingSessionRules []NoTradingSessionRules `fix:"1309,omitempty"`
}

func (m *TradingSessionRulesGrp) SetNoTradingSessionRules(v []NoTradingSessionRules) {
	m.NoTradingSessionRules = v
}
