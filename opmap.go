package filterbuilder

type IOperatorMap interface {
	Get(name string) IOperator
	Set(name string, operator IOperator) IOperator
	SetSource(source map[string]IOperator)
}

type operatorMap struct {
	source map[string]IOperator
}

func (o operatorMap) Get(name string) IOperator {
	return o.source[name]
}

func (o operatorMap) Set(name string, operator IOperator) IOperator {
	o.source[name] = operator
	return operator
}

func (o operatorMap) SetSource(source map[string]IOperator) {
	o.source = source
}

func NewOperatorMap() IOperatorMap {
	return &operatorMap{
		source: map[string]IOperator{
			"eq":    EQOperator{},
			"regex": RegexOperator{},
			"lt":    LTOperator{},
			"lte":   LTEOperator{},
			"gt":    GTOperator{},
			"gte":   GTEOperator{},
			"ne":    NEOperator{},
			"in":    INOperator{},
			"nin":   NINOperator{},
		},
	}
}
