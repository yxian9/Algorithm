package leetcode

import (
	"strconv"
	"strings"
)

// @leet start
func groupStrings(strings []string) [][]string {
	dict := map[string][]string{}
	for _, s := range strings {
		curID := getID(s)
		dict[curID] = append(dict[curID], s)
	}
	ret := make([][]string, 0, len(dict))
	for _, v := range dict {
		ret = append(ret, v)
	}
	return ret
}

func getID(s string) string {
	var sb strings.Builder
	for i := 1; i < len(s); i++ {
		diff := int(s[i]) - int(s[i-1])
		if diff < 0 {
			diff = 26 + diff
		}
		if i > 1 {
			sb.WriteString(",")
		}
		sb.WriteString(strconv.Itoa(diff))
	}
	return sb.String()
}

func find1(strings []string) [][]string {
	mp := map[string][]string{}
	for _, v := range strings {
		mp[getID(v)] = append(mp[getID(v)], v)
	}
	ret := make([][]string, 0, len(mp))
	for _, v := range mp {
		ret = append(ret, v)
	}
	return ret
}

// @leet end

func groupStrings(strings []string) [][]string {
	var ret [][]string
	lenMap := map[int][]string{}
	for _, item := range strings {
		lenMap[len(item)] = append(lenMap[len(item)], item)
	}
	for _, strings := range lenMap {
		// if k == 1 { // empty string
		// 	ret = append(ret, strings)
		// 	continue
		// }
		groups := find1(strings)
		for _, strings := range groups {
			ret = append(ret, strings)
		}
	}
	return ret
}

func getID(s string) string {
	var sb strings.Builder
	for i := 1; i < len(s); i++ {
		diff := int(s[i]) - int(s[i-1])
		if diff < 0 {
			diff = 26 + diff
		}
		str := strconv.Itoa(diff)
		if i > 1 {
			sb.WriteString(",")
		}
		sb.WriteString(str)
	}
	return sb.String()
}

func find1(strings []string) [][]string {
	mp := map[string][]string{}
	for _, v := range strings {
		mp[getID(v)] = append(mp[getID(v)], v)
	}
	ret := make([][]string, 0, len(mp))
	for _, v := range mp {
		ret = append(ret, v)
	}
	return ret
}

