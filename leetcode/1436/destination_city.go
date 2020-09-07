package main

func destCity(paths [][]string) string {
	m := map[string]string{}

	for _, path := range paths {
		m[path[0]] = path[1]
	}

	for _, to := range m {
		if _, ok := m[to]; !ok {
			return to
		}
	}

	return ""
}
