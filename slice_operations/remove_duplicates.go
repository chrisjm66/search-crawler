package slice_operations

import "slices"

func RemoveDuplicates[T comparable](slice []T) []T {
	nonDupeList := []T{}
	for _, v := range slice {
		if slices.Contains(nonDupeList, v) {
			nonDupeList = append(nonDupeList, v)
		}	
	}

	return nonDupeList
}
