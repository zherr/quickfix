package listordgrp

import (
	"github.com/quickfixgo/quickfix/fix50/commissiondata"
	"github.com/quickfixgo/quickfix/fix50/discretioninstructions"
	"github.com/quickfixgo/quickfix/fix50/displayinstruction"
	"github.com/quickfixgo/quickfix/fix50/instrument"
	"github.com/quickfixgo/quickfix/fix50/orderqtydata"
	"github.com/quickfixgo/quickfix/fix50/parties"
	"github.com/quickfixgo/quickfix/fix50/peginstructions"
	"github.com/quickfixgo/quickfix/fix50/preallocgrp"
	"github.com/quickfixgo/quickfix/fix50/spreadorbenchmarkcurvedata"
	"github.com/quickfixgo/quickfix/fix50/stipulations"
	"github.com/quickfixgo/quickfix/fix50/strategyparametersgrp"
	"github.com/quickfixgo/quickfix/fix50/trdgsesgrp"
	"github.com/quickfixgo/quickfix/fix50/triggeringinstruction"
	"github.com/quickfixgo/quickfix/fix50/undinstrmtgrp"
	"github.com/quickfixgo/quickfix/fix50/yielddata"
	"time"
)

//New returns an initialized ListOrdGrp instance
func New(noorders []NoOrders) *ListOrdGrp {
	var m ListOrdGrp
	m.SetNoOrders(noorders)
	return &m
}

//NoOrders is a repeating group in ListOrdGrp
type NoOrders struct {
	//ClOrdID is a required field for NoOrders.
	ClOrdID string `fix:"11"`
	//SecondaryClOrdID is a non-required field for NoOrders.
	SecondaryClOrdID *string `fix:"526"`
	//ListSeqNo is a required field for NoOrders.
	ListSeqNo int `fix:"67"`
	//ClOrdLinkID is a non-required field for NoOrders.
	ClOrdLinkID *string `fix:"583"`
	//SettlInstMode is a non-required field for NoOrders.
	SettlInstMode *string `fix:"160"`
	//Parties is a non-required component for NoOrders.
	Parties *parties.Parties
	//TradeOriginationDate is a non-required field for NoOrders.
	TradeOriginationDate *string `fix:"229"`
	//TradeDate is a non-required field for NoOrders.
	TradeDate *string `fix:"75"`
	//Account is a non-required field for NoOrders.
	Account *string `fix:"1"`
	//AcctIDSource is a non-required field for NoOrders.
	AcctIDSource *int `fix:"660"`
	//AccountType is a non-required field for NoOrders.
	AccountType *int `fix:"581"`
	//DayBookingInst is a non-required field for NoOrders.
	DayBookingInst *string `fix:"589"`
	//BookingUnit is a non-required field for NoOrders.
	BookingUnit *string `fix:"590"`
	//AllocID is a non-required field for NoOrders.
	AllocID *string `fix:"70"`
	//PreallocMethod is a non-required field for NoOrders.
	PreallocMethod *string `fix:"591"`
	//PreAllocGrp is a non-required component for NoOrders.
	PreAllocGrp *preallocgrp.PreAllocGrp
	//SettlType is a non-required field for NoOrders.
	SettlType *string `fix:"63"`
	//SettlDate is a non-required field for NoOrders.
	SettlDate *string `fix:"64"`
	//CashMargin is a non-required field for NoOrders.
	CashMargin *string `fix:"544"`
	//ClearingFeeIndicator is a non-required field for NoOrders.
	ClearingFeeIndicator *string `fix:"635"`
	//HandlInst is a non-required field for NoOrders.
	HandlInst *string `fix:"21"`
	//ExecInst is a non-required field for NoOrders.
	ExecInst *string `fix:"18"`
	//MinQty is a non-required field for NoOrders.
	MinQty *float64 `fix:"110"`
	//MaxFloor is a non-required field for NoOrders.
	MaxFloor *float64 `fix:"111"`
	//ExDestination is a non-required field for NoOrders.
	ExDestination *string `fix:"100"`
	//TrdgSesGrp is a non-required component for NoOrders.
	TrdgSesGrp *trdgsesgrp.TrdgSesGrp
	//ProcessCode is a non-required field for NoOrders.
	ProcessCode *string `fix:"81"`
	//Instrument is a required component for NoOrders.
	instrument.Instrument
	//UndInstrmtGrp is a non-required component for NoOrders.
	UndInstrmtGrp *undinstrmtgrp.UndInstrmtGrp
	//PrevClosePx is a non-required field for NoOrders.
	PrevClosePx *float64 `fix:"140"`
	//Side is a required field for NoOrders.
	Side string `fix:"54"`
	//SideValueInd is a non-required field for NoOrders.
	SideValueInd *int `fix:"401"`
	//LocateReqd is a non-required field for NoOrders.
	LocateReqd *bool `fix:"114"`
	//TransactTime is a non-required field for NoOrders.
	TransactTime *time.Time `fix:"60"`
	//Stipulations is a non-required component for NoOrders.
	Stipulations *stipulations.Stipulations
	//QtyType is a non-required field for NoOrders.
	QtyType *int `fix:"854"`
	//OrderQtyData is a required component for NoOrders.
	orderqtydata.OrderQtyData
	//OrdType is a non-required field for NoOrders.
	OrdType *string `fix:"40"`
	//PriceType is a non-required field for NoOrders.
	PriceType *int `fix:"423"`
	//Price is a non-required field for NoOrders.
	Price *float64 `fix:"44"`
	//StopPx is a non-required field for NoOrders.
	StopPx *float64 `fix:"99"`
	//SpreadOrBenchmarkCurveData is a non-required component for NoOrders.
	SpreadOrBenchmarkCurveData *spreadorbenchmarkcurvedata.SpreadOrBenchmarkCurveData
	//YieldData is a non-required component for NoOrders.
	YieldData *yielddata.YieldData
	//Currency is a non-required field for NoOrders.
	Currency *string `fix:"15"`
	//ComplianceID is a non-required field for NoOrders.
	ComplianceID *string `fix:"376"`
	//SolicitedFlag is a non-required field for NoOrders.
	SolicitedFlag *bool `fix:"377"`
	//IOIID is a non-required field for NoOrders.
	IOIID *string `fix:"23"`
	//QuoteID is a non-required field for NoOrders.
	QuoteID *string `fix:"117"`
	//TimeInForce is a non-required field for NoOrders.
	TimeInForce *string `fix:"59"`
	//EffectiveTime is a non-required field for NoOrders.
	EffectiveTime *time.Time `fix:"168"`
	//ExpireDate is a non-required field for NoOrders.
	ExpireDate *string `fix:"432"`
	//ExpireTime is a non-required field for NoOrders.
	ExpireTime *time.Time `fix:"126"`
	//GTBookingInst is a non-required field for NoOrders.
	GTBookingInst *int `fix:"427"`
	//CommissionData is a non-required component for NoOrders.
	CommissionData *commissiondata.CommissionData
	//OrderCapacity is a non-required field for NoOrders.
	OrderCapacity *string `fix:"528"`
	//OrderRestrictions is a non-required field for NoOrders.
	OrderRestrictions *string `fix:"529"`
	//CustOrderCapacity is a non-required field for NoOrders.
	CustOrderCapacity *int `fix:"582"`
	//ForexReq is a non-required field for NoOrders.
	ForexReq *bool `fix:"121"`
	//SettlCurrency is a non-required field for NoOrders.
	SettlCurrency *string `fix:"120"`
	//BookingType is a non-required field for NoOrders.
	BookingType *int `fix:"775"`
	//Text is a non-required field for NoOrders.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for NoOrders.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for NoOrders.
	EncodedText *string `fix:"355"`
	//SettlDate2 is a non-required field for NoOrders.
	SettlDate2 *string `fix:"193"`
	//OrderQty2 is a non-required field for NoOrders.
	OrderQty2 *float64 `fix:"192"`
	//Price2 is a non-required field for NoOrders.
	Price2 *float64 `fix:"640"`
	//PositionEffect is a non-required field for NoOrders.
	PositionEffect *string `fix:"77"`
	//CoveredOrUncovered is a non-required field for NoOrders.
	CoveredOrUncovered *int `fix:"203"`
	//MaxShow is a non-required field for NoOrders.
	MaxShow *float64 `fix:"210"`
	//PegInstructions is a non-required component for NoOrders.
	PegInstructions *peginstructions.PegInstructions
	//DiscretionInstructions is a non-required component for NoOrders.
	DiscretionInstructions *discretioninstructions.DiscretionInstructions
	//TargetStrategy is a non-required field for NoOrders.
	TargetStrategy *int `fix:"847"`
	//TargetStrategyParameters is a non-required field for NoOrders.
	TargetStrategyParameters *string `fix:"848"`
	//ParticipationRate is a non-required field for NoOrders.
	ParticipationRate *float64 `fix:"849"`
	//Designation is a non-required field for NoOrders.
	Designation *string `fix:"494"`
	//StrategyParametersGrp is a non-required component for NoOrders.
	StrategyParametersGrp *strategyparametersgrp.StrategyParametersGrp
	//MatchIncrement is a non-required field for NoOrders.
	MatchIncrement *float64 `fix:"1089"`
	//MaxPriceLevels is a non-required field for NoOrders.
	MaxPriceLevels *int `fix:"1090"`
	//DisplayInstruction is a non-required component for NoOrders.
	DisplayInstruction *displayinstruction.DisplayInstruction
	//PriceProtectionScope is a non-required field for NoOrders.
	PriceProtectionScope *string `fix:"1092"`
	//TriggeringInstruction is a non-required component for NoOrders.
	TriggeringInstruction *triggeringinstruction.TriggeringInstruction
	//RefOrderID is a non-required field for NoOrders.
	RefOrderID *string `fix:"1080"`
	//RefOrderIDSource is a non-required field for NoOrders.
	RefOrderIDSource *string `fix:"1081"`
	//PreTradeAnonymity is a non-required field for NoOrders.
	PreTradeAnonymity *bool `fix:"1091"`
	//ExDestinationIDSource is a non-required field for NoOrders.
	ExDestinationIDSource *string `fix:"1133"`
}

//NewNoOrders returns an initialized NoOrders instance
func NewNoOrders(clordid string, listseqno int, instrument instrument.Instrument, side string, orderqtydata orderqtydata.OrderQtyData) *NoOrders {
	var m NoOrders
	m.SetClOrdID(clordid)
	m.SetListSeqNo(listseqno)
	m.SetInstrument(instrument)
	m.SetSide(side)
	m.SetOrderQtyData(orderqtydata)
	return &m
}

func (m *NoOrders) SetClOrdID(v string)                            { m.ClOrdID = v }
func (m *NoOrders) SetSecondaryClOrdID(v string)                   { m.SecondaryClOrdID = &v }
func (m *NoOrders) SetListSeqNo(v int)                             { m.ListSeqNo = v }
func (m *NoOrders) SetClOrdLinkID(v string)                        { m.ClOrdLinkID = &v }
func (m *NoOrders) SetSettlInstMode(v string)                      { m.SettlInstMode = &v }
func (m *NoOrders) SetParties(v parties.Parties)                   { m.Parties = &v }
func (m *NoOrders) SetTradeOriginationDate(v string)               { m.TradeOriginationDate = &v }
func (m *NoOrders) SetTradeDate(v string)                          { m.TradeDate = &v }
func (m *NoOrders) SetAccount(v string)                            { m.Account = &v }
func (m *NoOrders) SetAcctIDSource(v int)                          { m.AcctIDSource = &v }
func (m *NoOrders) SetAccountType(v int)                           { m.AccountType = &v }
func (m *NoOrders) SetDayBookingInst(v string)                     { m.DayBookingInst = &v }
func (m *NoOrders) SetBookingUnit(v string)                        { m.BookingUnit = &v }
func (m *NoOrders) SetAllocID(v string)                            { m.AllocID = &v }
func (m *NoOrders) SetPreallocMethod(v string)                     { m.PreallocMethod = &v }
func (m *NoOrders) SetPreAllocGrp(v preallocgrp.PreAllocGrp)       { m.PreAllocGrp = &v }
func (m *NoOrders) SetSettlType(v string)                          { m.SettlType = &v }
func (m *NoOrders) SetSettlDate(v string)                          { m.SettlDate = &v }
func (m *NoOrders) SetCashMargin(v string)                         { m.CashMargin = &v }
func (m *NoOrders) SetClearingFeeIndicator(v string)               { m.ClearingFeeIndicator = &v }
func (m *NoOrders) SetHandlInst(v string)                          { m.HandlInst = &v }
func (m *NoOrders) SetExecInst(v string)                           { m.ExecInst = &v }
func (m *NoOrders) SetMinQty(v float64)                            { m.MinQty = &v }
func (m *NoOrders) SetMaxFloor(v float64)                          { m.MaxFloor = &v }
func (m *NoOrders) SetExDestination(v string)                      { m.ExDestination = &v }
func (m *NoOrders) SetTrdgSesGrp(v trdgsesgrp.TrdgSesGrp)          { m.TrdgSesGrp = &v }
func (m *NoOrders) SetProcessCode(v string)                        { m.ProcessCode = &v }
func (m *NoOrders) SetInstrument(v instrument.Instrument)          { m.Instrument = v }
func (m *NoOrders) SetUndInstrmtGrp(v undinstrmtgrp.UndInstrmtGrp) { m.UndInstrmtGrp = &v }
func (m *NoOrders) SetPrevClosePx(v float64)                       { m.PrevClosePx = &v }
func (m *NoOrders) SetSide(v string)                               { m.Side = v }
func (m *NoOrders) SetSideValueInd(v int)                          { m.SideValueInd = &v }
func (m *NoOrders) SetLocateReqd(v bool)                           { m.LocateReqd = &v }
func (m *NoOrders) SetTransactTime(v time.Time)                    { m.TransactTime = &v }
func (m *NoOrders) SetStipulations(v stipulations.Stipulations)    { m.Stipulations = &v }
func (m *NoOrders) SetQtyType(v int)                               { m.QtyType = &v }
func (m *NoOrders) SetOrderQtyData(v orderqtydata.OrderQtyData)    { m.OrderQtyData = v }
func (m *NoOrders) SetOrdType(v string)                            { m.OrdType = &v }
func (m *NoOrders) SetPriceType(v int)                             { m.PriceType = &v }
func (m *NoOrders) SetPrice(v float64)                             { m.Price = &v }
func (m *NoOrders) SetStopPx(v float64)                            { m.StopPx = &v }
func (m *NoOrders) SetSpreadOrBenchmarkCurveData(v spreadorbenchmarkcurvedata.SpreadOrBenchmarkCurveData) {
	m.SpreadOrBenchmarkCurveData = &v
}
func (m *NoOrders) SetYieldData(v yielddata.YieldData)                   { m.YieldData = &v }
func (m *NoOrders) SetCurrency(v string)                                 { m.Currency = &v }
func (m *NoOrders) SetComplianceID(v string)                             { m.ComplianceID = &v }
func (m *NoOrders) SetSolicitedFlag(v bool)                              { m.SolicitedFlag = &v }
func (m *NoOrders) SetIOIID(v string)                                    { m.IOIID = &v }
func (m *NoOrders) SetQuoteID(v string)                                  { m.QuoteID = &v }
func (m *NoOrders) SetTimeInForce(v string)                              { m.TimeInForce = &v }
func (m *NoOrders) SetEffectiveTime(v time.Time)                         { m.EffectiveTime = &v }
func (m *NoOrders) SetExpireDate(v string)                               { m.ExpireDate = &v }
func (m *NoOrders) SetExpireTime(v time.Time)                            { m.ExpireTime = &v }
func (m *NoOrders) SetGTBookingInst(v int)                               { m.GTBookingInst = &v }
func (m *NoOrders) SetCommissionData(v commissiondata.CommissionData)    { m.CommissionData = &v }
func (m *NoOrders) SetOrderCapacity(v string)                            { m.OrderCapacity = &v }
func (m *NoOrders) SetOrderRestrictions(v string)                        { m.OrderRestrictions = &v }
func (m *NoOrders) SetCustOrderCapacity(v int)                           { m.CustOrderCapacity = &v }
func (m *NoOrders) SetForexReq(v bool)                                   { m.ForexReq = &v }
func (m *NoOrders) SetSettlCurrency(v string)                            { m.SettlCurrency = &v }
func (m *NoOrders) SetBookingType(v int)                                 { m.BookingType = &v }
func (m *NoOrders) SetText(v string)                                     { m.Text = &v }
func (m *NoOrders) SetEncodedTextLen(v int)                              { m.EncodedTextLen = &v }
func (m *NoOrders) SetEncodedText(v string)                              { m.EncodedText = &v }
func (m *NoOrders) SetSettlDate2(v string)                               { m.SettlDate2 = &v }
func (m *NoOrders) SetOrderQty2(v float64)                               { m.OrderQty2 = &v }
func (m *NoOrders) SetPrice2(v float64)                                  { m.Price2 = &v }
func (m *NoOrders) SetPositionEffect(v string)                           { m.PositionEffect = &v }
func (m *NoOrders) SetCoveredOrUncovered(v int)                          { m.CoveredOrUncovered = &v }
func (m *NoOrders) SetMaxShow(v float64)                                 { m.MaxShow = &v }
func (m *NoOrders) SetPegInstructions(v peginstructions.PegInstructions) { m.PegInstructions = &v }
func (m *NoOrders) SetDiscretionInstructions(v discretioninstructions.DiscretionInstructions) {
	m.DiscretionInstructions = &v
}
func (m *NoOrders) SetTargetStrategy(v int)              { m.TargetStrategy = &v }
func (m *NoOrders) SetTargetStrategyParameters(v string) { m.TargetStrategyParameters = &v }
func (m *NoOrders) SetParticipationRate(v float64)       { m.ParticipationRate = &v }
func (m *NoOrders) SetDesignation(v string)              { m.Designation = &v }
func (m *NoOrders) SetStrategyParametersGrp(v strategyparametersgrp.StrategyParametersGrp) {
	m.StrategyParametersGrp = &v
}
func (m *NoOrders) SetMatchIncrement(v float64) { m.MatchIncrement = &v }
func (m *NoOrders) SetMaxPriceLevels(v int)     { m.MaxPriceLevels = &v }
func (m *NoOrders) SetDisplayInstruction(v displayinstruction.DisplayInstruction) {
	m.DisplayInstruction = &v
}
func (m *NoOrders) SetPriceProtectionScope(v string) { m.PriceProtectionScope = &v }
func (m *NoOrders) SetTriggeringInstruction(v triggeringinstruction.TriggeringInstruction) {
	m.TriggeringInstruction = &v
}
func (m *NoOrders) SetRefOrderID(v string)            { m.RefOrderID = &v }
func (m *NoOrders) SetRefOrderIDSource(v string)      { m.RefOrderIDSource = &v }
func (m *NoOrders) SetPreTradeAnonymity(v bool)       { m.PreTradeAnonymity = &v }
func (m *NoOrders) SetExDestinationIDSource(v string) { m.ExDestinationIDSource = &v }

//ListOrdGrp is a fix50 Component
type ListOrdGrp struct {
	//NoOrders is a required field for ListOrdGrp.
	NoOrders []NoOrders `fix:"73"`
}

func (m *ListOrdGrp) SetNoOrders(v []NoOrders) { m.NoOrders = v }
