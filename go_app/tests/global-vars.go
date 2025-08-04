package main

const (
	TypeInt = iota
	TypeFloat
	TypeString
	TypeBool
)

var ValueOperatorMap = map[int]map[string]bool{}

func PopulateValueOperatorMap() {
	ValueOperatorMap[TypeInt] = map[string]bool{
		"<":  true,
		"<=": true,
		"=":  true,
		"!=": true,
		">=": true,
		">":  true,
	}
	ValueOperatorMap[TypeFloat] = map[string]bool{
		"<":  true,
		"<=": true,
		"=":  true,
		"!=": true,
		">=": true,
		">":  true,
	}
	ValueOperatorMap[TypeString] = map[string]bool{
		"<":        true,
		"<=":       true,
		"=":        true,
		"<>":       true,
		">=":       true,
		">":        true,
		"LIKE":     true,
		"IN":       true,
		"NOT LIKE": true,
		"GLOB":     true,
	}
	ValueOperatorMap[TypeBool] = map[string]bool{
		"=":  true,
		"<>": true,
	}
}
