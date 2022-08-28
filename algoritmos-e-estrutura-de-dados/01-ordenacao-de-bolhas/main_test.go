package main

import (
	"reflect"
	"testing"
)

func TestBubbleSort_Ascending_1(t *testing.T) {
	tests := []struct {
		scenario string
		numbers  []int
		want     []int
	}{
		{
			scenario: "dever ordenar dois numeros",
			numbers:  []int{3, 1},
			want:     []int{1, 3},
		},
		{
			scenario: "dever ordenar três numeros",
			numbers:  []int{3, 2, 1},
			want:     []int{1, 2, 3},
		},
		{
			scenario: "dever ordenar 10 numeros",
			numbers:  []int{3, 2, 1, 5, 7, 8, 10, 4, 6, 9},
			want:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			scenario: "dever ordenar 13 numeros com itens repetidos",
			numbers:  []int{11, 11, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			want:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 11, 11},
		},
		{
			scenario: "dever ordenar 15 numeros  com itens repetidos",
			numbers:  []int{11, 11, 11, 10, 9, 8, 7, 6, 5, 5, 4, 4, 3, 2, 1},
			want:     []int{1, 2, 3, 4, 4, 5, 5, 6, 7, 8, 9, 10, 11, 11, 11},
		},
	}
	for _, tt := range tests {
		t.Run(tt.scenario, func(t *testing.T) {
			b := BubbleSort{}
			if got := b.Ascending_1(tt.numbers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BubbleSort.Ascending_1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBubbleSort_Ascending_2(t *testing.T) {
	tests := []struct {
		scenario string
		numbers  []int
		want     []int
	}{
		{
			scenario: "dever ordenar dois numeros",
			numbers:  []int{3, 1},
			want:     []int{1, 3},
		},
		{
			scenario: "dever ordenar três numeros",
			numbers:  []int{3, 2, 1},
			want:     []int{1, 2, 3},
		},
		{
			scenario: "dever ordenar 10 numeros",
			numbers:  []int{3, 2, 1, 5, 7, 8, 10, 4, 6, 9},
			want:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			scenario: "dever ordenar 13 numeros com itens repetidos",
			numbers:  []int{11, 11, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			want:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 11, 11},
		},
		{
			scenario: "dever ordenar 15 numeros  com itens repetidos",
			numbers:  []int{11, 11, 11, 10, 9, 8, 7, 6, 5, 5, 4, 4, 3, 2, 1},
			want:     []int{1, 2, 3, 4, 4, 5, 5, 6, 7, 8, 9, 10, 11, 11, 11},
		},
	}
	for _, tt := range tests {
		t.Run(tt.scenario, func(t *testing.T) {
			b := BubbleSort{}
			if got := b.Ascending_2(tt.numbers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BubbleSort.Ascending_1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBubbleSort_Descending(t *testing.T) {
	tests := []struct {
		scenario string
		numbers  []int
		want     []int
	}{
		{
			scenario: "dever ordenar dois numeros",
			numbers:  []int{1, 3},
			want:     []int{3, 1},
		},
		{
			scenario: "dever ordenar três numeros",
			numbers:  []int{1, 2, 3},
			want:     []int{3, 2, 1},
		},
		{
			scenario: "dever ordenar 10 numeros",
			numbers:  []int{3, 2, 1, 5, 7, 8, 10, 4, 6, 9},
			want:     []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		},
		{
			scenario: "dever ordenar 13 numeros com itens repetidos",
			numbers:  []int{0, 11, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			want:     []int{11, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
		},
		{
			scenario: "dever ordenar 15 numeros  com itens repetidos",
			numbers:  []int{1, 2, 3, 0, 11, 4, 5, 6, 1, 8, 9, 7, 11, 11, 10},
			want:     []int{11, 11, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.scenario, func(t *testing.T) {
			b := BubbleSort{}
			if got := b.Descending(tt.numbers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BubbleSort.Ascending_1() = %v, want %v", got, tt.want)
			}
		})
	}
}
