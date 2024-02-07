package structure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJoin(t *testing.T) {
	var testCases = []struct {
		usr      []User
		grp      []Group
		expected []UserGroup
	}{
		{
			usr:      []User{},
			grp:      []Group{},
			expected: []UserGroup{},
		},
		{
			usr: []User{
				{
					Id:      1,
					GroupId: 1,
				},
			},
			grp: []Group{
				{
					Id:   1,
					Name: `First`,
				},
			},
			expected: []UserGroup{
				{
					Id:      1,
					GroupId: 1,
					Name:    `First`,
				},
			},
		},
		{
			usr: []User{
				{
					Id:      1,
					GroupId: 1,
				},
				{
					Id:      2,
					GroupId: 2,
				},
				{
					Id:      3,
					GroupId: 2,
				},
			},
			grp: []Group{
				{
					Id:   1,
					Name: `First`,
				},
				{
					Id:   2,
					Name: `Second`,
				},
			},
			expected: []UserGroup{
				{
					Id:      1,
					GroupId: 1,
					Name:    `First`,
				},
				{
					Id:      2,
					GroupId: 2,
					Name:    `Second`,
				},
				{
					Id:      3,
					GroupId: 2,
					Name:    `Second`,
				},
			},
		},
	}

	for _, v := range testCases {
		assert.Equal(t, join(v.usr, v.grp), v.expected)
	}
}
