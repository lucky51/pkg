package ext

// Clone clone slice
func Clone[T any](s []T) []T {
	news := make([]T, len(s))
	copy(news, s)
	return news
}
