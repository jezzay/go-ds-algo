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

type MaxBenchmark struct {
	name string
	list []int
}

func TestMaxKadane(t *testing.T) {
	tests := createMaxTests()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := MaxKadane(test.list)
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

func TestMaxBruteForce(t *testing.T) {
	tests := createMaxTests()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := MaxBruteForce(test.list)
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

func BenchmarkMaxBruteForce(b *testing.B) {
	benchmarks := createMaxBenchmarks()
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				MaxBruteForce(bm.list)
			}
		})
	}
}

func BenchmarkMaxKadane(b *testing.B) {
	benchmarks := createMaxBenchmarks()
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				MaxKadane(bm.list)
			}
		})
	}
}

func createMaxTests() []*MaxTest {
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
	return tests
}

func createMaxBenchmarks() []MaxBenchmark {
	benchmarks := make([]MaxBenchmark, 0)

	benchmarks = append(benchmarks, MaxBenchmark{name: "single", list: makeList(5)})
	benchmarks = append(benchmarks, MaxBenchmark{name: "double", list: makeList(5, 2)})
	benchmarks = append(benchmarks, MaxBenchmark{name: "five", list: makeList(5, 2, -1, 3, 0)})
	benchmarks = append(benchmarks, MaxBenchmark{name: "eight", list: makeList(5, 2, -1, 3, 0, -1, 3, 0)})
	benchmarks = append(benchmarks, MaxBenchmark{name: "ten", list: makeList(5, 2, -15, 2, 2, -3, 0, 8, 6)})
	benchmarks = append(benchmarks, MaxBenchmark{name: "13", list: makeList(-3, 5, 2, -15, 65, 2, 2, -3, 0, 8, 0, 8, 6)})
	benchmarks = append(benchmarks, MaxBenchmark{name: "15", list: makeList(-3, 5, 2, -15, 65, 2, 2, -3, 0, 8, 0, 8, 6, 7, 5)})
	benchmarks = append(benchmarks, MaxBenchmark{name: "17", list: makeList(65, 2, -15, 2, 2, -3, 5, 2, -15, 2, 2, -3, 0, 8, 0, 8, 6)})
	benchmarks = append(benchmarks, MaxBenchmark{name: "20", list: makeList(65, 2, -15, 2, 2, -3, 5, 2, -15, 2, 2, -3, 0, 8, 0, 8, 6, 16, 73, 34)})
	benchmarks = append(benchmarks, MaxBenchmark{name: "25", list: makeList(65, 2, -15, 2, 2, -3, 5, 8, 6, 16, -31, 4, 2, -15, 2, 2, -3, 0, 8, 0, 8, 6, 16, -31, 4)})

	list := makeList(-2, 1, -3, 4, -1, 10, 1, 5, 34, 0, 2, 4, 9, 2, 1, 4, 15, 1, 7, 4, 8, 2, 13, -15, 6, 24, -6)
	benchmarks = append(benchmarks, MaxBenchmark{name: "27", list: list})

	return benchmarks
}

func makeList(numbers ...int) []int {
	list := make([]int, len(numbers))
	for i, num := range numbers {
		list[i] = num
	}
	return list
}
