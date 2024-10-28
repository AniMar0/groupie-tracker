package TRC

func mm() {
	local := []string{}

	for v := range Location.Index {
		local = append(local, Location.Index[v].Locations...)
	}
	Alle.Local = UniqueLocation(local)
}

func UniqueLocation(local []string) []string {
	if len(local) < 2 {
		return local
	}

	med := len(local) / 2

	left := UniqueLocation(local[:med])
	right := UniqueLocation(local[med:])
	return merge(left, right)
}

func merge(left, right []string) []string {
	result := []string{}
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			if len(result) == 0 || result[len(result)-1] != left[i] {
				result = append(result, left[i])
			}
			i++
		} else if left[i] > right[j] {
			if len(result) == 0 || result[len(result)-1] != right[j] {
				result = append(result, right[j])
			}
			j++
		} else {
			if len(result) == 0 || result[len(result)-1] != left[i] {
				result = append(result, left[i])
			}
			i++
			j++
		}
	}

	for i < len(left) {
		if len(result) == 0 || result[len(result)-1] != left[i] {
			result = append(result, left[i])
		}
		i++
	}
	
	for j < len(right) {
		if len(result) == 0 || result[len(result)-1] != right[j] {
			result = append(result, right[j])
		}
		j++
	}

	return result
}
