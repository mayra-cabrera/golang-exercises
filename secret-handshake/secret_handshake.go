package secret

var steps = []string{"wink", "double blink", "close your eyes", "jump"}

func Handshake(number int) []string {
	var handshake = make([]string, 0)
	if number <= 0 {
		return handshake
	}
	for s, step := range steps {
		if (1<<uint(s))&number > 0 {
			handshake = append(handshake, step)
		}
	}
	if (1<<uint(len(steps)))&number > 0 {
		handshake = reverse(handshake)
	}
	return handshake
}

func reverse(signs []string) []string {
	for i := 0; i < len(signs)/2; i++ {
		j := len(signs) - i - 1
		signs[i], signs[j] = signs[j], signs[i]
	}
	return signs
}
