package derivativeinstrumentattribute

//New returns an initialized DerivativeInstrumentAttribute instance
func New() *DerivativeInstrumentAttribute {
	var m DerivativeInstrumentAttribute
	return &m
}

//NoDerivativeInstrAttrib is a repeating group in DerivativeInstrumentAttribute
type NoDerivativeInstrAttrib struct {
	//DerivativeInstrAttribType is a non-required field for NoDerivativeInstrAttrib.
	DerivativeInstrAttribType *int `fix:"1313"`
	//DerivativeInstrAttribValue is a non-required field for NoDerivativeInstrAttrib.
	DerivativeInstrAttribValue *string `fix:"1314"`
}

//NewNoDerivativeInstrAttrib returns an initialized NoDerivativeInstrAttrib instance
func NewNoDerivativeInstrAttrib() *NoDerivativeInstrAttrib {
	var m NoDerivativeInstrAttrib
	return &m
}

func (m *NoDerivativeInstrAttrib) SetDerivativeInstrAttribType(v int) {
	m.DerivativeInstrAttribType = &v
}
func (m *NoDerivativeInstrAttrib) SetDerivativeInstrAttribValue(v string) {
	m.DerivativeInstrAttribValue = &v
}

//DerivativeInstrumentAttribute is a fix50sp2 Component
type DerivativeInstrumentAttribute struct {
	//NoDerivativeInstrAttrib is a non-required field for DerivativeInstrumentAttribute.
	NoDerivativeInstrAttrib []NoDerivativeInstrAttrib `fix:"1311,omitempty"`
}

func (m *DerivativeInstrumentAttribute) SetNoDerivativeInstrAttrib(v []NoDerivativeInstrAttrib) {
	m.NoDerivativeInstrAttrib = v
}
