package cipher

type Cipher interface {
	Encode(string) string
	Decode(string) string
}

type Caesar struct{}

type Shift struct {
	key int
}

type Vigenere struct {
	key string
}

func Downcase(str string) string {
	a := len(str)
	tmp := ""
	for i := 0; i < a; i++ {
		switch {
		case str[i] >= 'A' && str[i] <= 'Z':
			tmp += string(str[i] + 32)
		case str[i] >= 'a' && str[i] <= 'z':
			tmp += string(str[i])
		}

	}

	return tmp
}

func NewCaesar() Cipher {
	return Caesar{}
}

func (c Caesar) Encode(s string) string {
	tmp := Downcase(s)
	var str string
	for i := 0; i < len(tmp); i++ {
		switch {
		case tmp[i] >= 'x':
			str += string(tmp[i] - 23)
		default:
			str += string(tmp[i] + 3)
		}
	}
	return str
}

func (c Caesar) Decode(s string) string {
	tmp := Downcase(s)
	var str string
	for i := 0; i < len(tmp); i++ {
		switch {
		case tmp[i] <= 'c':
			str += string(tmp[i] + 23)
		default:
			str += string(tmp[i] - 3)
		}
	}
	return str
}

func NewShift(key int) Cipher {
	switch {
	case key >= 1 && key <= 25:
		return Shift{key}
	case key >= -25 && key <= -1:
		return Shift{key}
	default:
		return nil
	}
}

func (c Shift) Encode(s string) string {
	tmp := Downcase(s)
	var char uint8
	var str string
	for i := 0; i < len(tmp); i++ {
		char = tmp[i] + byte(c.key)
		switch {
		case char > 122:
			char = 'a' + ((char - 1) % 122)
		case char < 97:
			char = 'z' - (97 % (char + 1))
		}
		str += string(char)
	}
	return str
}

func (c Shift) Decode(s string) string {
	tmp := Downcase(s)
	var char uint8
	var str string
	for i := 0; i < len(tmp); i++ {
		char = tmp[i] - byte(c.key)
		switch {
		case char > 122:
			char = 'a' + ((char - 1) % 122)
		case char < 97:
			char = 'z' - (97 % (char + 1))
		}
		str += string(char)
	}
	return str
}

func NewVigenere(key string) Cipher {
	flag := false
	for i := 0; i < len(key); i++ {
		if key[i] < 97 || key[i] > 122 {
			return nil
		}
		if key[i] != 97 {
			flag = true
		}
	}
	if flag == true {
		return Vigenere{key}
	} else {
		return nil
	}
}

func (c Vigenere) Encode(s string) string {
	tmp := Downcase(s)
	var char uint8
	var str string
	lok := len(c.key)
	for i := 0; i < len(tmp); i++ {
		char = tmp[i] + c.key[i%lok] - 97
		switch {
		case char > 122:
			char = 'a' + ((char - 1) % 122)
		case char < 97:
			char = 'z' - (97 % (char + 1))
		}
		str += string(char)
	}
	return str
}

func (c Vigenere) Decode(s string) string {
	tmp := Downcase(s)
	var char uint8
	var str string
	lok := len(c.key)
	for i := 0; i < len(tmp); i++ {
		char = tmp[i] - c.key[i%lok] + 97
		switch {
		case char > 122:
			char = 'a' + ((char - 1) % 122)
		case char < 97:
			char = 'z' - (97 % (char + 1))
		}
		str += string(char)
	}
	return str
}
