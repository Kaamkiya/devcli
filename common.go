package main

func includes(slice []string, query string) {
	for _, el := range slice {
		if el == query {
			return true
		}
	}
	return false
}
