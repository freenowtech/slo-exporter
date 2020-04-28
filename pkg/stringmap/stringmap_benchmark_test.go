package stringmap

import (
	"strconv"
	"strings"
	"testing"
)

type benchmarkCase struct {
	name string
	data StringMap
}

func generateStringMap(keysCount int, stringLen int) StringMap {
	newMap := make(StringMap, keysCount)
	for i := 0; i < keysCount; i++ {
		key := strings.Repeat(strconv.Itoa(i), stringLen)
		newMap[key] = key
	}
	return newMap
}

func BenchmarkStringMap_Matches(b *testing.B) {
	testCases := []benchmarkCase{
		{name: "small map/small keys", data: generateStringMap(3, 3)},
		{name: "small map/large keys", data: generateStringMap(3, 1000)},
		{name: "large map/small keys", data: generateStringMap(1000, 3)},
		{name: "large map/large keys", data: generateStringMap(1000, 1000)},
	}
	for _, tc := range testCases {
		b.Run("StringMap.Copy/"+tc.name, func(b *testing.B) {
			tc.data.Copy()
		})
		b.Run("StringMap.Merge on : "+tc.name, func(b *testing.B) {
			tc.data.Merge(tc.data)
		})
		b.Run("StringMap.Keys on : "+tc.name, func(b *testing.B) {
			tc.data.Keys()
		})
		b.Run("StringMap.NewWith on : "+tc.name, func(b *testing.B) {
			tc.data.NewWith("foo", "bar")
		})
		b.Run("StringMap.Select on : "+tc.name, func(b *testing.B) {
			tc.data.Select([]string{"a", "b", "c", "d", "e", "f", "g", "h"})
		})
		b.Run("StringMap.SortedKeys on : "+tc.name, func(b *testing.B) {
			tc.data.SortedKeys()
		})
		b.Run("StringMap.Without on : "+tc.name, func(b *testing.B) {
			tc.data.Without([]string{"a", "b", "c", "d", "e", "f", "g", "h"})
		})
		b.Run("StringMap.String on : "+tc.name, func(b *testing.B) {
			tc.data.String()
		})
	}
}