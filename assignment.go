package main

type Assignment [defaultSize]int

func NewAssignment(value int) (result Assignment) {
	return result.myMap(func(_ int) int { return value })
}

func (assignment Assignment) any(fn func(int) bool) bool {
	for _, v := range assignment {
		if fn(v) {
			return true
		}
	}
	return false
}

func (assignment Assignment) count(fn func(int) bool) (result int) {
	for _, v := range assignment {
		if fn(v) {
			result++
		}
	}
	return result
}

func (assignment Assignment) findIndex(fn func(int) bool) int {
	for i, v := range assignment {
		if fn(v) {
			return i
		}
	}
	return -1
}

func (assignment Assignment) myMap(fn func(int) int) (result Assignment) {
	for i, v := range assignment {
		result[i] = fn(v)
	}
	return
}

func (assignment Assignment) translateAssignment() (result Assignment) {
	for i := 0; i < defaultSize; i++ {
		result[assignment[i]] = i
	}
	return
}
