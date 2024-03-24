package arr

type meet = struct {
	start int
	end   int
}

func getIntersection(mettings [][]int) [][]int {
	var intersections []meet
	var result [][]int
	var start, end int
	var i int = 1
	var j int = 0
	var current meet
	var next meet
	var wasIntersected bool

	for j < len(mettings) {
		current = meet{
			start: mettings[j][0],
			end:   mettings[j][1],
		}

		for i < len(mettings) {
			next = meet{
				start: mettings[i][0],
				end:   mettings[i][1],
			}

			if hasIntersect(current, next) {
				if current.start < next.start {
					start = current.start
				} else {
					start = next.start
				}

				if current.end > next.end {
					end = current.end
				} else {
					end = next.end
				}

				intersections = append(intersections, meet{start: start, end: end})
				wasIntersected = true
			} else {
				intersections = append(intersections, next)
			}
			i++
		}
		if !wasIntersected {
			intersections = append(intersections, current)
		}
		j++
	}

	for _, v := range intersections {
		result = append(result, []int{v.start, v.end})
	}

	return result
}

func hasIntersect(a meet, b meet) bool {
	return a.start <= b.end && a.end >= b.start
}
