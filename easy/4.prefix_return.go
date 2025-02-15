package main

import "fmt"

func longestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 0; i < len(strs); i++ {

		if len(prefix) > len(strs[i]) || strs[i][:len(prefix)] != prefix {
			prefix = prefix[:len(prefix)-1]

			if len(prefix) == 0 {
				return ""
			}
			i--
		}
	}
	return prefix

}

func main() {
	strs := []string{"flower", "flow", "flight"}
	test := longestCommonPrefix(strs)

	fmt.Println(test)
}
