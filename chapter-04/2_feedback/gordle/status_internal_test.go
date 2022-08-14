package gordle

import "testing"

/*
Result of the benchmarks
$ go test -run=^$ -bench=. -benchmem

goos: darwin
goarch: arm64
pkg: github.com/ablqk/tiny-go-projects/chapter-04/2_feedback/gordle
BenchmarkStringConcat1-10       174882942                6.850 ns/op           0 B/op          0 allocs/op
BenchmarkStringConcat2-10       15633693                74.28 ns/op           24 B/op          2 allocs/op
BenchmarkStringConcat3-10        8609542               137.1 ns/op            56 B/op          4 allocs/op
BenchmarkStringConcat4-10        5873654               201.1 ns/op           104 B/op          6 allocs/op
BenchmarkStringConcat5-10        4455464               275.2 ns/op           160 B/op          8 allocs/op
BenchmarkStringBuilder1-10      71407850                16.69 ns/op            8 B/op          1 allocs/op
BenchmarkStringBuilder2-10      30721999                38.28 ns/op           24 B/op          2 allocs/op
BenchmarkStringBuilder3-10      27036134                45.64 ns/op           24 B/op          2 allocs/op
BenchmarkStringBuilder4-10      17278803                70.44 ns/op           56 B/op          3 allocs/op
BenchmarkStringBuilder5-10      16189770                73.27 ns/op           56 B/op          3 allocs/op
PASS
ok      github.com/ablqk/tiny-go-projects/chapter-04/2_feedback/gordle  13.762s
*/

func BenchmarkStringConcat1(b *testing.B) {
	fb := feedback{absentCharacter}
	for n := 0; n < b.N; n++ {
		_ = fb.StringConcat()
	}
}

func BenchmarkStringConcat2(b *testing.B) {
	fb := feedback{absentCharacter, wrongPosition}
	for n := 0; n < b.N; n++ {
		_ = fb.StringConcat()
	}
}

func BenchmarkStringConcat3(b *testing.B) {
	fb := feedback{absentCharacter, wrongPosition, correctPosition}
	for n := 0; n < b.N; n++ {
		_ = fb.StringConcat()
	}
}

func BenchmarkStringConcat4(b *testing.B) {
	fb := feedback{absentCharacter, wrongPosition, correctPosition, absentCharacter}

	for n := 0; n < b.N; n++ {
		_ = fb.StringConcat()
	}
}

func BenchmarkStringConcat5(b *testing.B) {
	fb := feedback{absentCharacter, wrongPosition, correctPosition, absentCharacter, wrongPosition}
	for n := 0; n < b.N; n++ {
		_ = fb.StringConcat()
	}
}

func BenchmarkStringBuilder1(b *testing.B) {
	fb := feedback{absentCharacter}
	for n := 0; n < b.N; n++ {
		_ = fb.String()
	}
}

func BenchmarkStringBuilder2(b *testing.B) {
	fb := feedback{absentCharacter, wrongPosition}
	for n := 0; n < b.N; n++ {
		_ = fb.String()
	}
}

func BenchmarkStringBuilder3(b *testing.B) {
	fb := feedback{absentCharacter, wrongPosition, correctPosition}
	for n := 0; n < b.N; n++ {
		_ = fb.String()
	}
}

func BenchmarkStringBuilder4(b *testing.B) {
	fb := feedback{absentCharacter, wrongPosition, correctPosition, absentCharacter}
	for n := 0; n < b.N; n++ {
		_ = fb.String()
	}
}

func BenchmarkStringBuilder5(b *testing.B) {
	fb := feedback{absentCharacter, wrongPosition, correctPosition, absentCharacter, wrongPosition}
	for n := 0; n < b.N; n++ {
		_ = fb.String()

	}
}
