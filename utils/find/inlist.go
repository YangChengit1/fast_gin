package find

func InList[T comparable](list []T, key T) bool {
	for _, val := range list {
		if val == key {
			return true
		}
	}
	return false
}
