package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// test
type Sex struct {
	aname  string
	isMale bool
}

type Superman struct {
	values []int64
}

type CustomGroup struct {
	Name    string
	val     int
	Peolple []string
	mens    []Superman
	Sex
}

// main
func TestCompareAny(t *testing.T) {
	var k int64 = 3
	superman := Superman{
		values: []int64{1, 2, 3},
	}

	type test struct {
		slice         interface{}
		value         interface{}
		containsValue bool
	}

	aMetric := &CustomGroup{
		Name:    "Pino",
		val:     100,
		Peolple: []string{"tizio", "caio"},
		mens:    []Superman{superman},
		Sex: Sex{
			aname:  "alborg",
			isMale: true},
	}

	tests := []test{
		{[]int{3, 4, 5, 6}, 4, true},
		{[]int64{3, 4, 5, 6}, k, true},

		{[]*CustomGroup{aMetric,
			{
				Name:    "Balo",
				val:     111,
				Peolple: []string{"tizio", "caio"},
				Sex: Sex{
					aname:  "alborg",
					isMale: true},
			},
		}, aMetric, true},

		{[]CustomGroup{*aMetric,
			{
				Name:    "Balo",
				val:     111,
				Peolple: []string{"tizio", "caio"},
				Sex: Sex{
					aname:  "alborg",
					isMale: true},
			},
		}, CustomGroup{
			Name:    "Pino",
			val:     100,
			Peolple: []string{"tizio", "caio"},
			mens:    []Superman{superman},
			Sex: Sex{
				aname:  "alborg",
				isMale: true},
		}, true},

		{[]*CustomGroup{aMetric,
			{
				Name:    "Balo",
				val:     111,
				Peolple: []string{"tizio", "caio"},
				Sex: Sex{
					aname:  "alborg",
					isMale: true},
			},
		}, &CustomGroup{
			val: 22,
		}, false},
	}

	for _, tc := range tests {
		assert.Equal(t, containsAny(tc.slice, tc.value), tc.containsValue)
	}
}
