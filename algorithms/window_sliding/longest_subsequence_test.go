package window_sliding

import (
	"testing"
)

/*
abcabcbb
*/
func lcs(s string) int {
	// <key:character, value:largest index of the character in the array>
	var start, end, max int
	m := map[int32]int{}
	runes := []rune(s)
	for end < len(runes) {
		// 如果在map中存在这个字符,并且它索引的位置处于滑窗中, 移动start到这个字符后面的位置
		if index, ok := m[runes[end]]; ok && index >= start {
			start = index + 1
		}
		// 更新字符的所在位置，使其保持最新的索引
		m[runes[end]] = end
		if windowSize := end - start + 1; windowSize > max {
			max = windowSize
		}
		end++
	}
	return max
}

func TestLcs(t *testing.T) {
	tests := []struct {
		input  string
		length int
	}{
		{input: "ABDEFGABEF", length: 6},
		{input: "BBBB", length: 1},
		{input: "GEEKSFORGEEKS", length: 7},
	}

	for _, test := range tests {
		output := lcs(test.input)
		if output != test.length {
			t.Errorf("got %d, want %d", output, test.length)
		}
	}
}
