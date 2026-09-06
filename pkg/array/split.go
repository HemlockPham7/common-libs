package array

// SplitIntoBatches splits a slice into batches with minimal memory allocation.
//
// Parameters:
//   - objects: the slice to split into batches.
//   - size: the maximum number of elements in each batch.
//
// Returns:
//   - batches that share the underlying array with the input slice.
func SplitIntoBatches[T any](objects []T, size int) [][]T {
	batches := make([][]T, 0, (len(objects)+size-1)/size)
	for size < len(objects) {
		objects, batches = objects[size:], append(batches, objects[0:size:size])
	}
	batches = append(batches, objects)
	return batches
}
