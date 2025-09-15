package hamming

import "errors"

func Distance(a, b string) (int, error) {
	count := 0

	switch {
	case len(a) == 0 && len(b) == 0:
		return 0, nil
	case len(a) != len(b):
		return 0, errorDistance
	default:
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				count++
			}
		}

		return count, nil
	}
}

var errorDistance = errors.New("error")
