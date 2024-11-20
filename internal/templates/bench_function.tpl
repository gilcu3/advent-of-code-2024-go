func Benchmark{{.Year}}{{.FormatDay}}(b *testing.B) {
        aoc.RegisterYears()
        input := util.TestInput({{.Year}}, {{.Day}})
        p := aoc.NewPuzzle({{.Year}}, {{.Day}})
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
