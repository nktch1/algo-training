package main

func findTheDifference(s string, t string) byte {

	var r rune

	for _, rr := range s {
		r ^= rr
	}

	for _, rr := range t {
		r ^= rr
	}

	return byte(r)
}
