package main

/*
	Array includes item.

Parameters
----------
slice : []string

	the array to search

query : string

	the value to search for

Returns
-------
bool: true if slice inclues query, false otherwise
*/
func includes(slice []string, query string) bool {
	for _, el := range slice {
		if el == query {
			return true
		}
	}
	return false
}
