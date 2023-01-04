package util

// Return the elements in `b` that are not in `a`
func DiffStrings(a, b []string) []string {
	set := make(map[string]struct{}, len(a))
	for _, s := range a {
		set[s] = struct{}{}
	}

	var diff []string
	for _, s := range b {
		if _, exists := set[s]; !exists {
			diff = append(diff, s)
		}
	}

	if len(diff) > 0 {
		return diff
	}

	return nil
}
