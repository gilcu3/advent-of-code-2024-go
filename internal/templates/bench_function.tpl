func Benchmark{{.Year}}{{.FormatDay}}(b *testing.B) {
        aoc.RegisterYears()
        input := aoc.TestInput({{.Year}}, {{.Day}})
        p := aoc.NewPuzzle({{.Year}}, {{.Day}})
        if p.PartA(input) != fmt.Sprint(nil) {
            b.Run("PartA", func(b *testing.B) {
                b.ResetTimer()
                for i := 0; i < b.N; i++ {
                    p.PartA(input)
                }
            })
        }
        if p.PartB(input) != fmt.Sprint(nil) {
            b.Run("PartB", func(b *testing.B) {
                b.ResetTimer()
                for i := 0; i < b.N; i++ {
                    p.PartB(input)
                }
            })
        }
}
