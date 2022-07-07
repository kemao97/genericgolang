package generics

func Filter[A any](fx func(A) bool, slice []A) (r []A) {
	for _, item := range slice {
		if fx(item) {
			r = append(r, item)
		}
	}
	return
}
