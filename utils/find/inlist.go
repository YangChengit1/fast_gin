package find

// InList 查找list中是否存在key
func InList[T comparable](list []T, key T) bool {
	for _, val := range list {
		if val == key {
			return true
		}
	}
	return false
}
