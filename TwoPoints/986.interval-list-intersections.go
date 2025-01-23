package leetcode

// @leet start
func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	if len(firstList) == 0 || len(secondList) == 0 {
		return nil
	}
	var ret [][]int
	var prev []int
	prev = pop(&firstList, &secondList)
	for {
		cur := pop(&firstList, &secondList)
		if cur == nil {
			break
		}
		overlap, ok := intersect(prev, cur)
		if ok {
			ret = append(ret, overlap)
		}
		if cur[1] > prev[1] {
			prev = cur
		}

		// if len(firstList) == 0 && len(secondList) == 0 {
		// 	break
		// }
		// if len(firstList) == 0 || len(secondList) == 0 {
		// 	cur := pop(&firstList, &secondList)
		// 	overlap, ok := intersect(prev, cur)
		// 	if ok {
		// 		ret = append(ret, overlap)
		// 	}
		// 	break
		// } // even one list is emtpy (by pop/ not initial emty), the other may still
		// contribute
	}
	return ret
}

func intersect(l1, l2 []int) (res []int, find bool) {
	if l1[0] > l2[1] || l2[0] > l1[1] {
		return nil, false
	}
	return []int{max(l1[0], l2[0]), min(l1[1], l2[1])}, true
}

func pop(l1, l2 *[][]int) (res []int) {
	if len(*l1) == 0 && len(*l2) == 0 {
		return nil
	}
	var chose *[][]int
	if len(*l1) == 0 {
		chose = l2
	} else if len(*l2) == 0 {
		chose = l1
	} else if (*l1)[0][0] < (*l2)[0][0] {
		chose = l1
	} else {
		chose = l2
	}
	res = (*chose)[0]
	(*chose) = (*chose)[1:]
	return res
}

// @leet end

