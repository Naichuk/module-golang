package letter

type FreqMap map[rune]int

func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(list []string) FreqMap {
	resChannel := make(chan FreqMap, len(list))
	for _, s := range list {
		s := s
		go func() {
			resChannel <- Frequency(s)
		}()
	}

	res := make(FreqMap)
	for range list {
		for letter, freq := range <-resChannel {
			res[letter] += freq
		}
	}
	return res
}
