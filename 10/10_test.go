package main

import (
	"strings"
	"testing"
)

func TestSolve1(t *testing.T) {

	testCases := []struct {
		m    string
		want int
	}{
		{
			m: `
					.#..#
					.....
					#####
					....#
					...##
				`,
			want: 8,
		},
		{
			m: `
							......#.#.
							#..#.#....
							..#######.
							.#.#.###..
							.#..#.....
							..#....#.#
							#..#....#.
							.##.#..###
							##...#..#.
							.#....####
						`,
			want: 33,
		},
		{
			m: `
						#.#...#.#.
						.###....#.
						.#....#...
						##.#.#.#.#
						....#.#.#.
						.##..###.#
						..#...##..
						..##....##
						......#...
						.####.###.
					`,
			want: 35,
		},
		{
			m: `
					.#..#..###
					####.###.#
					....###.#.
					..###.##.#
					##.##.#.#.
					....###..#
					..#.#..#.#
					#..#.#.###
					.##...##.#
					.....#.#..
				`,
			want: 41,
		},
		{
			m: `
					.#..##.###...#######
					##.############..##.
					.#.######.########.#
					.###.#######.####.#.
					#####.##.#.##.###.##
					..#####..#.#########
					####################
					#.####....###.#.#.##
					##.#################
					#####.##.###..####..
					..######..##.#######
					####.##.####...##..#
					.#####..#.######.###
					##...#.##########...
					#.##########.#######
					.####.#.###.###.#.##
					....##.##.###..#####
					.#.#.###########.###
					#.#.#.#####.####.###
					###.##.####.##.#..##
				`,
			want: 210,
		},
	}

	for i, tc := range testCases {
		m, err := parse(strings.Split(tc.m, "\n"))
		if err != nil {
			t.Fatalf("#%d: parsemap: %v", i, err)
		}
		if got := solve1(m); got != tc.want {
			t.Fatalf("#%d: got %d, want %d", i, got, tc.want)
		}
	}

}
