// Code generated by aocgen; DO NOT EDIT.
package tests

import (
	"fmt"
	"testing"

	"aocgen/internal/aoc"
	"aocgen/internal/util"
)

func Benchmark202401(b *testing.B) {
	aoc.RegisterYears()
	input := util.TestInput(2024, 1)
	p := aoc.NewPuzzle(2024, 1)
	if p.Part1(input) != fmt.Sprint(nil) {
		b.Run("Part1", func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				p.Part1(input)
			}
		})
	}
	if p.Part2(input) != fmt.Sprint(nil) {
		b.Run("Part2", func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				p.Part2(input)
			}
		})
	}
}
func Benchmark202402(b *testing.B) {
	aoc.RegisterYears()
	input := util.TestInput(2024, 2)
	p := aoc.NewPuzzle(2024, 2)
	if p.Part1(input) != fmt.Sprint(nil) {
		b.Run("Part1", func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				p.Part1(input)
			}
		})
	}
	if p.Part2(input) != fmt.Sprint(nil) {
		b.Run("Part2", func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				p.Part2(input)
			}
		})
	}
}
