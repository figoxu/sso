package main

func convertIntArray2Int32Array(vs []int) []int32 {
	v32s := make([]int32, 0)
	for v, _ := range vs {
		v32s = append(v32s, int32(v))
	}
	return v32s
}
