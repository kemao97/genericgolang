package generics

func Reduce[A any](fx func(A, A) A, slice []A) (r A) {
	for i, item := range slice {
		if i == 0 {
			r = item
			continue
		}

		r = fx(r, item)
	}
	return
}
