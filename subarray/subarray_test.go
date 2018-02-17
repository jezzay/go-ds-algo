package subarray

import "testing"

func TestMax(t *testing.T) {
	list := make([]int, 9)
	list[0] = -2
	list[1] = 1
	list[2] = -3
	list[3] = 4
	list[4] = -1
	list[5] = 2
	list[6] = 1
	list[7] = -5
	list[8] = 4
	result := Max(list)
	if len(result) != 4 {
		t.Errorf("Expected size of maximum sub array to be 4, received %v", len(result))
	} else if result[0] != 4 || result[1] != -1 || result[2] != 2 || result[3] != 1 {
		t.Errorf("Expected contiguous array to contain [4 -1 2 1], received %v", result)
	}

}
