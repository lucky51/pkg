package ext

// Clone clone slice
func Clone[T any](s []T) []T {
	if s == nil {
		return make([]T, 0)
	}
	news := make([]T, len(s))
	copy(news, s)
	return news
}
