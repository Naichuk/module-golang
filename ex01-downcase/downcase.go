package downcase

func Downcase(str string) (string, error) {
	a := len(str)
	tmp := ""
	for i := 0; i < a; i++ {
		switch {
		case str[i] >= 'A' && str[i] < 'a':
			tmp += string(str[i] + 32)
		default:
			tmp += string(str[i])
		}
	}

	return tmp, nil
}
