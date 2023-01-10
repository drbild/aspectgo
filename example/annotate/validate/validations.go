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

func IntBetween(min int, max int) ValidateFunc {
	return func(i interface{}) bool {
		v, ok := i.(int)
		if !ok {
			return false
		}

		return (min <= v) && (v <= max)
	}
}

func ValidHelloValues() []string {
	return []string{"hello", "hola", "bonjour", "nĭ hăo"}
}

func ValidGoodbyeValues() []string {
	return []string{"goodbye", "adios", "au revoir", "zia jian"}
}
