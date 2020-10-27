package main

type ValueWithCoords [3]int
type ValueWithCoordsArray [defaultSize * defaultSize]ValueWithCoords

func (vs ValueWithCoordsArray) Len() int {
	return defaultSize * defaultSize
}
func (vs ValueWithCoordsArray) Less(i, j int) bool {
	return vs[i][0] < vs[j][0]
}
func (vs *ValueWithCoordsArray) Swap(i, j int) {
	vs[i], vs[j] = vs[j], vs[i]
}
