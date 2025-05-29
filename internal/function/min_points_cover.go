package function

import "sort"

type Section struct {
	Start float64
	End   float64
}

func MinPointsCover(sections []Section) int {
	if len(sections) == 0 {
		return 0
	}

	sort.Slice(sections, func(i, j int) bool {
		return sections[i].End < sections[j].End
	})

	count := 0
	currentPoint := -1e18

	for _, seg := range sections {
		if seg.Start > currentPoint {
			currentPoint = seg.End
			count++
		}
	}

	return count
}
