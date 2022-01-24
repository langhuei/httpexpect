package httpexpect

type Integer struct {
	chain chain
	value int64
}

func NewInteger(reporter Reporter, value int64) *Integer {
	return &Integer{makeChain(reporter), value}
}

func (n *Integer) Raw() int64 {
	return n.value
}

func (n *Integer) Schema(schema interface{}) *Integer {
	checkSchema(&n.chain, n.value, schema)
	return n
}

func (n *Integer) Equal(value interface{}) *Integer {
	v, ok := canonInteger(&n.chain, value)
	if !ok {
		return n
	}
	if !(n.value == v) {
		n.chain.fail("\nexpected number equal to:\n %v\n\nbut got:\n %v",
			v, n.value)
	}
	return n
}

func (n *Integer) NotEqual(value interface{}) *Integer {
	v, ok := canonInteger(&n.chain, value)
	if !ok {
		return n
	}
	if !(n.value != v) {
		n.chain.fail("\nexpected number not equal to:\n %v\n\nbut got:\n %v",
			v, n.value)
	}
	return n
}
