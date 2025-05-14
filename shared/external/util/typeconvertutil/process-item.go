package typeconvertutil

func ProcessItem[E any, D any](items []E, handlerFunc func(E) D) []D {
	res := make([]D, len(items))
	for i, item := range items {
		res[i] = handlerFunc(item)
	}
	return res
}