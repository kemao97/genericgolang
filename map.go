package generics

func Map[A any, B any](fx func(A) B, slice []A) (r []B) {
	for _, item := range slice {
		r = append(r, fx(item))
	}
	return
}
