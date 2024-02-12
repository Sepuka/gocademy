package arr

import (
	"errors"
	"fmt"
	"testing"
)

type simpleItem struct {
	date   []string
	values []uint64
}

func (s *simpleItem) toList(isShows bool) *StatList {
	n := len(s.values)
	if n == 0 {
		return nil
	}

	result := &StatList{}
	node := result

	for i := 0; i < n-1; i++ {
		node.Date = s.date[i]
		if isShows {
			node.Shows = s.values[i]
		} else {
			node.Clicks = s.values[i]
		}

		node.Next = &StatList{}
		node = node.Next
	}

	node.Date = s.date[n-1]
	if isShows {
		node.Shows = s.values[n-1]
	} else {
		node.Clicks = s.values[n-1]
	}

	return result
}

func compare(s *StatList, values []StatList) error {
	if s == nil {
		if len(values) == 0 {
			return nil
		} else {
			return errors.New("StatList is nil, but values is not empty")
		}
	}

	for i, v := range values {
		switch {
		case s == nil:
			return fmt.Errorf("StatList is nil, but values[%d] is %v", i, v)
		case s.Date != v.Date:
			return fmt.Errorf("StatList date is %s, but values[%d] is %s", s.Date, i, v.Date)
		case s.Shows != v.Shows:
			return fmt.Errorf("StatList shows is %d, but values[%d] is %d", i, s.Shows, v.Shows)
		case s.Clicks != v.Clicks:
			return fmt.Errorf("StatList clicks is %d, but values[%d] is %d", i, s.Clicks, v.Clicks)
		default:
			s = s.Next
		}
	}

	if s != nil {
		return errors.New("next StatList is not nil, but values is empty")
	}

	return nil
}

func TestCommonStatistics(t *testing.T) {
	testCases := []struct {
		name     string
		shows    simpleItem
		clicks   simpleItem
		expected []StatList
	}{
		{
			name:   "empty",
			shows:  simpleItem{},
			clicks: simpleItem{},
		},
		{
			name: "only_shows",
			shows: simpleItem{
				date:   []string{"2023-04-05", "2023-04-06", "2023-04-07"},
				values: []uint64{1000, 2000, 3000},
			},
			expected: []StatList{
				{Date: "2023-04-05", Shows: 1000},
				{Date: "2023-04-06", Shows: 2000},
				{Date: "2023-04-07", Shows: 3000},
			},
		},
		{
			name: "only_clicks",
			clicks: simpleItem{
				date:   []string{"2023-04-05", "2023-04-06", "2023-04-07"},
				values: []uint64{100, 200, 300},
			},
			expected: []StatList{
				{Date: "2023-04-05", Clicks: 100},
				{Date: "2023-04-06", Clicks: 200},
				{Date: "2023-04-07", Clicks: 300},
			},
		},
		{
			name: "shows_and_clicks",
			shows: simpleItem{
				date:   []string{"2023-04-05", "2023-04-06", "2023-04-07"},
				values: []uint64{1000, 2000, 3000},
			},
			clicks: simpleItem{
				date:   []string{"2023-04-05", "2023-04-07", "2023-04-08"},
				values: []uint64{100, 300, 400},
			},
			expected: []StatList{
				{Date: "2023-04-05", Shows: 1000, Clicks: 100},
				{Date: "2023-04-06", Shows: 2000},
				{Date: "2023-04-07", Shows: 3000, Clicks: 300},
				{Date: "2023-04-08", Clicks: 400},
			},
		},
		{
			name: "different_dates",
			shows: simpleItem{
				date:   []string{"2023-04-05", "2023-04-06", "2023-04-07"},
				values: []uint64{1000, 2000, 3000},
			},
			clicks: simpleItem{
				date:   []string{"2023-04-08", "2023-04-09", "2023-04-10", "2023-04-11", "2023-04-12"},
				values: []uint64{100, 200, 300, 400, 500},
			},
			expected: []StatList{
				{Date: "2023-04-05", Shows: 1000},
				{Date: "2023-04-06", Shows: 2000},
				{Date: "2023-04-07", Shows: 3000},
				{Date: "2023-04-08", Clicks: 100},
				{Date: "2023-04-09", Clicks: 200},
				{Date: "2023-04-10", Clicks: 300},
				{Date: "2023-04-11", Clicks: 400},
				{Date: "2023-04-12", Clicks: 500},
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			shows := tc.shows.toList(true)
			clicks := tc.clicks.toList(false)

			result := CommonStatistics(shows, clicks)

			if err := compare(result, tc.expected); err != nil {
				t.Errorf("unexpected result: %v", err)
			}
		})
	}
}
