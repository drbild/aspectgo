package validate

type ValidateFunc func(interface{}) bool

func StringInSlice(valid []string, ignoreCase bool) ValidateFunc {
	return func(i interface{}) bool {
		v, ok := i.(string)
		if !ok {
			return false
		}

		for _, str := range valid {
			if v == str {
				return true
			}
		}

		return false
	}
}

func ValidValues() []string {
	return []string{"hello", "hola", "bonjour", "nĭ hăo"}
}
