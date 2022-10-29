package filterbuilder

type IOperatorMap interface {
	Get(name string) IOperator
	Set(operator IOperator) IOperator
	SetSource(source map[string]IOperator)
}

type operatorMap struct {
	source map[string]IOperator
}

func (o operatorMap) Get(name string) IOperator {
	//TODO implement me
	panic("implement me")
}

func (o operatorMap) Set(operator IOperator) IOperator {
	//TODO implement me
	panic("implement me")
}

func (o operatorMap) SetSource(source map[string]IOperator) {
	//TODO implement me
	panic("implement me")
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
