package hw09structvalidator

func contains(strings []string, search string) bool {
	for _, el := range strings {
		if el == search {
			return true
		}
	}

	return false
}
