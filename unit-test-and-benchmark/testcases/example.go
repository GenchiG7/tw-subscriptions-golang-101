package unittest

func showRange(n int) string {

	if n >= 80 {
		return "A"
	}

	if n >= 60 {
		return "B"
	}

	if n >= 40 {
		return "C"
	}

	if n >= 20 {
		return "D"
	}

	return "E"
}
