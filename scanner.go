package filterbuilder

type IScanner interface {
	Scan(filterStruct interface{}) ([]IFilterField, error)
}

type scanner struct {
	OperatorMap     map[string]string
	LookupTagName   string
	OperatorTagName string
	RelationTagName string
}
