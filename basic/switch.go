package basic

func typeAssertion(t interface{}) string {
	switch t.(type) {
	case string:
		return "string"
	case int:
		return "int"
	default:
		return "unknown"
	}
}
