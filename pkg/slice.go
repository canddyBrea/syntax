package pkg

type Number interface {
	~int
}

// Delete delete slice
func Delete[T Number](stone []T, idx ...T) []T {
	var result []T
	for _, val := range idx {
		result = append(stone[:val-1], stone[val:]...)
	}
	return ShrinkSlice[T](result[:len(stone)-len(idx)])
}

// ShrinkSlice 缩容机制
func ShrinkSlice[T Number](slice []T) []T {
	length := len(slice)
	capacity := cap(slice)

	if length <= capacity/2 && capacity > 2 {
		newCapacity := capacity / 2
		if length < newCapacity {
			newCapacity = length
		} else if newCapacity > 2*length {
			newCapacity = 2 * length
		}
		newSlice := make([]T, length, newCapacity)
		copy(newSlice, slice)

		return newSlice
	}

	return slice
}
