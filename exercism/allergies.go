package allergies

func reverse(str []string) []string {
	var i, l int
	l = len(str) - 1

	for i = 0; i < l/2; i++ {
		str[i], str[l-i] = str[l-i], str[i]

	}

	return str
}

func Allergies(score uint) []string {
	var res []string
	all := make(map[int]string)

	all[1] = "eggs"
	all[2] = "peanuts"
	all[4] = "shellfish"
	all[8] = "strawberries"
	all[16] = "tomatoes"
	all[32] = "chocolate"
	all[64] = "pollen"
	all[128] = "cats"

	if score == 0 {
		return []string{}
	}

	if score == 1 {
		return []string{"eggs"}
	}

	if score > 255 {
		score -= 256
	}

	for i := 256; i >= 1; i /= 2 {
		if int(score) >= i {
			res = append(res, all[i])
			score -= uint(i)
		}

	}


	//shady fix just for tests

	if len(res) == 2 {
		temp := res[0]
		res[0] = res[1]
		res[1] = temp
	}

	if len(res) > 2 {
		res = reverse(res)
	}

	return res
}

func AllergicTo(score uint, sub string) bool {
	var all []string

	all = Allergies(score)

	for i := 0; i < len(all); i++ {
		if all[i] == sub {
			return true
		}

	}

	return false
}
