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
				"B":   []string{"COM", "C", "G"},
				"C":   []string{"B", "D"},
				"D":   []string{"C", "E", "I"},
				"E":   []string{"D", "F", "J"},
				"F":   []string{"E"},
				"G":   []string{"B", "H"},
				"H":   []string{"G"},
				"I":   []string{"D"},
				"J":   []string{"E", "K"},
				"K":   []string{"J", "L"},
				"L":   []string{"K"},
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

func TestFindMinDistance(t *testing.T) {

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
				"I)SAN",
				"K)YOU",
			},
			want: 4,
		},
	}

	for _, tc := range testCases {
		if got := findMinDistance(parsemap(tc.m), "YOU", "SAN"); !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("got %v, want %v", got, tc.want)
		}
	}

}
