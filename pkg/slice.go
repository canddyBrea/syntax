package pkg

type Number interface {
	~int
}

// Delete delete slice
func Delete[T Number](s []T, idx ...T) []T {
	result := make([]T, len(s))
	for _, val := range idx {
		result = append(s[:val-1], s[val:]...)
	}
	return ShrinkSlice[T](result[:len(s)-len(idx)])
}

// ShrinkSlice 缩容机制
func ShrinkSlice[T Number](s []T) []T {
	length := len(s)
	capacity := cap(s)

	if length <= capacity/2 && capacity > 2 {
		newCapacity := capacity / 2
		if length < newCapacity {
			newCapacity = length
		} else if newCapacity > 2*length {
			newCapacity = 2 * length
		}
		newSlice := make([]T, length, newCapacity)
		copy(newSlice, s)

		return newSlice
	}

	return s
}
