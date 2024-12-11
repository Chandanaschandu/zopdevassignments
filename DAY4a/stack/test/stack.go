package test

func Push(value int, arr []int) {

	arr = append(arr, value)

}
func Pop(arr []int) int {
	res := arr[len(arr)-1]
	return res
}

func Stack(value int, arr []int) (s []int, err error) {
	Push(value, arr)
	Pop(arr)
	return arr, nil

}
