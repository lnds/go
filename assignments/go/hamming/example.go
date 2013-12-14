package hamming

func Distance(a, b string) (d int) {
	if len(b) < len(a) {
		a, b = b, a
	}

	for i := range a {
		if a[i] != b[i] {
			d++
		}
	}

	return
}
