package brackets

var b = map[string]string{"{": "}", "[": "]", "(": ")"}

func Bracket(input string) (bool, error) {
	tmp := " "
	for _, tok := range input {
		char := b[tmp[len(tmp)-1:]]
		tmp += string(tok)
		if string(tok) == char {
			tmp = tmp[:len(tmp)-2]
		}
	}
	if len(tmp) == 1 {
		return true, nil
	} else {
		return false, nil
	}
}
