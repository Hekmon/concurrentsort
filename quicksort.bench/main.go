package main

import "runtime"

func main() {
	QuickSortMinSizeCompute(0, 2, 20, 1<<24, 1<<7, runtime.NumCPU())
}
