package main

import (
	"reflect"
	"testing"
)

func TestParsemap(t *testing.T) {

	testCases := []struct {
		ss   []string
		want map[string][]string
	}{
		{
			ss: []string{
				"COM)B",
				"B)C",
				"C)D",
				"D)E",
				"E)F",
				"B)G",
				"G)H",
				"D)I",
				"E)J",
				"J)K",
				"K)L",
			},
			want: map[string][]string{
				"COM": []string{"B"},
				"B":   []string{"C", "G"},
				"C":   []string{"D"},
				"D":   []string{"E", "I"},
				"E":   []string{"F", "J"},
				"F":   nil,
				"G":   []string{"H"},
				"H":   nil,
				"I":   nil,
				"J":   []string{"K"},
				"K":   []string{"L"},
				"L":   nil,
			},
		},
	}

	for _, tc := range testCases {
		if got := parsemap(tc.ss); !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("got %v, want %v", got, tc.want)
		}
	}

}

func TestCountOrbits(t *testing.T) {

	testCases := []struct {
		m    []string
		want int
	}{
		{
			m: []string{
				"COM)B",
				"B)C",
				"C)D",
				"D)E",
				"E)F",
				"B)G",
				"G)H",
				"D)I",
				"E)J",
				"J)K",
				"K)L",
			},
			want: 42,
		},
	}

	for _, tc := range testCases {
		if got := countOrbits(parsemap(tc.m)); !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("got %v, want %v", got, tc.want)
		}
	}

}
