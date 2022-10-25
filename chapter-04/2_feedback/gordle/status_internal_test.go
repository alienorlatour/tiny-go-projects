package gordle

import "testing"

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
