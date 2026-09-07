package common

func IsInList[T comparable](item T, list []T) bool {
	return FindInList(item, list) != -1
}
