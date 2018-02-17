package subarray

import (
	"testing"
	"reflect"
)

type MaxTest struct {
	name         string
	list         []int
	expectedLen  int
	expectedList []int
}

func TestMax(t *testing.T) {
	tests := make([]*MaxTest, 0)

	tests = append(tests, &MaxTest{
		name:         "positive-negative-mix",
		list:         makeList(-2, 1, -3, 4, -1, 2, 1, -5, 4),
		expectedLen:  4,
		expectedList: makeList(4, -1, 2, 1),
	})

	tests = append(tests, &MaxTest{
		name:         "all-positive",
		list:         makeList(8, 2, 5, 3, 4),
		expectedLen:  5,
		expectedList: makeList(8, 2, 5, 3, 4),
	})

	//tests = append(tests, &MaxTest{
	//	name:         "all-negative",
	//	list:         makeList(-1, -24, -5, -3, -42),
	//	expectedLen:  1,
	//	expectedList: makeList(-42),
	//})

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Max(test.list)
			if len(result) != test.expectedLen {
				t.Errorf("%v: Expected size of maximum sub array to be %v, received %v",
					test.name, test.expectedLen, len(result))
			}
			if !reflect.DeepEqual(result, test.expectedList) {
				t.Errorf("%v: Expected contiguous array to contain %v , received %v", test.name, test.expectedList, result)
			}

		})
	}

}

func BenchmarkMax1(b *testing.B)  { benchmarkMax(makeList(5), b) }
func BenchmarkMax2(b *testing.B)  { benchmarkMax(makeList(5, 2), b) }
func BenchmarkMax5(b *testing.B)  { benchmarkMax(makeList(5, 2, -1, 3, 0), b) }
func BenchmarkMax10(b *testing.B) { benchmarkMax(makeList(5, 2, -15, 2, 2, -3, 0, 8, 6), b) }

func BenchmarkMax27(b *testing.B) {
	list := makeList(-2, 1, -3, 4, -1, 10, 1, 5, 34, 0, 2, 4, 9, 2, 1, 4, 15, 1, 7, 4, 8, 2, 13, -15, 6, 24, -6)
	benchmarkMax(list, b)
}

var result []int

func benchmarkMax(list []int, b *testing.B) {

	var r []int
	for i := 0; i < b.N; i++ {
		r = Max(list)
	}
	result = r
}

func makeList(numbers ...int) []int {
	list := make([]int, len(numbers))
	for i, num := range numbers {
		list[i] = num
	}
	return list
}
