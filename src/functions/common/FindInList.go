package common

func FindInList[T comparable](item T, list []T) int {
	for i, v := range list {
		if v == item {
			return i
		}
	}
	return -1
}
