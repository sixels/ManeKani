package util

// Return the elements in `b` that are not in `a`
func DiffStrings(a, b []string) []string {
	set := make(map[string]struct{}, len(a))
	for _, s := range a {
		set[s] = struct{}{}
	}

	diff := make(map[string]struct{}, len(b))
	for _, s := range b {
		if _, exists := set[s]; !exists {
			diff[s] = struct{}{}
		}
	}

	if len(diff) > 0 {
		return MapKeys(diff)
	}

	return nil
}

func MapKeys[M ~map[K]V, K comparable, V any](m M) []K {
	slice := make([]K, 0, len(m))
	for key := range m {
		slice = append(slice, key)
	}
	return slice
}
