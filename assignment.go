package main

type Assignment [defaultSize]int

func (assignment Assignment) any(fn func(int) bool) bool {
	for _,v := range assignment {
		if fn(v) {
			return true
		}
	}
	return false
}
func (assignment Assignment) translateAssignment() (result Assignment) {
	for i := 0; i < defaultSize; i++ {
		result[assignment[i]] = i
	}
	return
}