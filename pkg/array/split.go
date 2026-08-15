package array

// https://github.com/golang/go/wiki/SliceTricks#batching-with-minimal-allocation
func SplitIntoBatches[T any](objects []T, size int) [][]T {
	batches := make([][]T, 0, (len(objects)+size-1)/size)
	for size < len(objects) {
		objects, batches = objects[size:], append(batches, objects[0:size:size])
	}
	batches = append(batches, objects)
	return batches
}
