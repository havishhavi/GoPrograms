package handlers

func ElementSearch(str []rune, s rune) bool {
	var searchE bool
	for _, v := range str {

		if string(s) == string(v) {
			searchE = true
		} else {
			searchE = false
		}

	}
	return searchE
}
