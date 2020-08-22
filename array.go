package helper

func FindArrayItem(slice []uint, val uint) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func UniqueArrayItem(intSlice []uint) []uint {
	keys := make(map[uint]bool)
	list := []uint{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
