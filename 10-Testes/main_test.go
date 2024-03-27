package main

import (
	"testing"
)

func Test_find_simple(t *testing.T) {
	scenarios := []struct {
		list   []string
		item   string
		result string
	}{
		{
			list: []string{"a", "b"},
			item: "a", result: "a",
		},
	}

	for _, v := range scenarios {
		if result := find(v.list, v.item); result != v.result {
			t.Errorf("want %s, received %s", v.result, result)
		}
	}
}

func Test_find_tableDrive(t *testing.T) {
	type args struct {
		list []string
		item string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "find",
			args: args{list: []string{"a", "b"}, item: "b"},
			want: "b",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := find(tt.args.list, tt.args.item); got != tt.want {
				t.Errorf("find() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_find_subTest(t *testing.T) {
	t.Run("find b", func(t *testing.T) {
		result := find([]string{"a", "b"}, "b")
		want := "b"

		if result != want {
			t.Errorf("want %s, received %s", want, result)
		}
	})
	t.Run("find a", func(t *testing.T) {
		result := find([]string{"a", "b"}, "a")
		want := "a"

		if result != want {
			t.Errorf("want %s, received %s", want, result)
		}
	})
}
