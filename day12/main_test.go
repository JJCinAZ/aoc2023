package main

import "testing"

func Test_matcher(t *testing.T) {
	type args struct {
		pattern string
		input   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{"?###????????", "###...##...#"}, false},
		{"2", args{"?###????????", ".###..##...#"}, true},
		{"3", args{"?###????????", ".###..##..#"}, false},
		{"4", args{"?###.?#??.???#.?#?", "####.##.#.#.#.###"}, false},
		{"6", args{"?###.?#??.???#.?#?", "####.##.#.#..#.###"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := matcher(tt.args.pattern, tt.args.input); got != tt.want {
				t.Errorf("matcher() = %v, want %v", got, tt.want)
			}
		})
	}
}
