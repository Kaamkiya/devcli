package main

func includes(slice []string, query string) bool {
	for _, el := range slice {
		if el == query {
			return true
		}
	}
	return false
}
