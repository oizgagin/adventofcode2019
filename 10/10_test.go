package main

import (
	"reflect"
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

func TestSortPoints(t *testing.T) {

	testCases := []struct {
		m    string
		from point
		want []point
	}{
		{
			m: `
				.#....#####...#..
				##...##.#####..##
				##...#...#.#####.
				..#.....#...###..
				..#.#.....#....##
			`,
			from: point{8, 3},
			want: []point{
				point{8, 3},
				point{8, 1},
				point{9, 0},
				point{9, 1},
				point{10, 0},
				point{9, 2},
				point{11, 1},
				point{12, 1},
				point{11, 2},
				point{15, 1},
				point{12, 2},
				point{13, 2},
				point{14, 2},
				point{15, 2},
				point{12, 3},
				point{16, 4},
				point{15, 4},
				point{10, 4},
				point{4, 4},
				point{2, 4},
				point{2, 3},
				point{0, 2},
				point{1, 2},
				point{0, 1},
				point{1, 1},
				point{5, 2},
				point{1, 0},
				point{5, 1},
				point{6, 1},
				point{6, 0},
				point{7, 0},
				point{8, 0},
				point{10, 1},
				point{14, 0},
				point{16, 1},
				point{13, 3},
				point{14, 3},
			},
		},
	}

	for i, tc := range testCases {
		v, err := parse(strings.Split(tc.m, "\n"))
		if err != nil {
			t.Fatalf("#%d: parsemap: %v", i, err)
		}
		m := v.(pointsmap)

		sortPoints(m.points, tc.from)
		if !reflect.DeepEqual(m.points, tc.want) {
			t.Fatalf("got\n\n%#v\nwant\n\n%#v", m.points, tc.want)
		}
	}

}
