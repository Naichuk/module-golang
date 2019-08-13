package downcase

func Downcase(s string) (string, error) {
	var tmp string
	for i := 0; i < len(s); i++ {
		c := s[i]
		if 'A' <= c && c <= 'Z' {
			c += 'a' - 'A'
		}
		tmp += string(c)
	}
	return tmp, nil
}
