package main

//fmt.Println(lengthOfLongestSubstring("abcabcbb"))
//fmt.Println(lengthOfLongestSubstring("bbbbb"))
//fmt.Println(lengthOfLongestSubstring("pwwkew"))
//fmt.Println(lengthOfLongestSubstring(""))
//fmt.Println(lengthOfLongestSubstring(" "))

func lengthOfLongestSubstring(s string) int {

	var (
		b, e, best int
		m          = map[rune]struct{}{}
		runes      = []rune(s)
	)

	for b >= 0 && b <= e && e <= len(runes) {

		var (
			keepExpanding = true
			nextPart      = runes[b:e]
		)

		for _, r := range nextPart {
			if _, isExists := m[r]; isExists {
				keepExpanding = false
				break
			}

			m[r] = struct{}{}
		}

		if keepExpanding {
			if best < len(nextPart) {
				best = len(nextPart)
			}

			e++
		} else {
			e++
			b++
		}

		m = map[rune]struct{}{}
	}

	return best
}
