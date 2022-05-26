package porgrammers

func solution(lottos []int, win_nums []int) []int {
	grade := make(map[int]int)
	for i := 0; i <= len(win_nums); i++ {
		tmp := len(win_nums)
		if i < 2 {
			grade[i] = tmp
		} else {
			grade[i] = tmp - i + 1
		}
	}

	var zeroCnt, hitCnt int
	for i := 0; i < len(lottos); i++ {
		if lottos[i] == 0 {
			zeroCnt++
			continue
		}

		for j := 0; j < len(win_nums); j++ {
			if lottos[i] == win_nums[j] {
				hitCnt++
			}
		}
	}

	if zeroCnt > 0 {
		return []int{grade[hitCnt+zeroCnt], grade[hitCnt]}
	} else {
		return []int{grade[hitCnt], grade[hitCnt]}
	}

}
