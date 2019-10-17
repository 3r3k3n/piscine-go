package piscine

func DivMod(a *int, b *int) {
	f := *a
	*a = *a / (*b)
	*b = f % (*b)
}
