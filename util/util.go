package util

import "strconv"

func Ptr2Str(p *string) string {
	if p == nil {
		return ""
	}
	return *p
}

func Uint2Str(d uint) string {
	return strconv.FormatUint(uint64(d), 10)
}

func ChunkBy[T any](items []T, chunkSize int) (chunks [][]T) {
	for chunkSize < len(items) {
		items, chunks = items[chunkSize:], append(chunks, items[0:chunkSize:chunkSize])
	}
	return append(chunks, items)
}
