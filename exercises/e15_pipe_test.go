package exercises

import "testing"

func TestPipe(t *testing.T) {
	ex := ByID(15)
	t.Run("setup 직후에는 실패", func(t *testing.T) {
		mustFail(t, ex, prepare(t, ex), "아직 없습니다: `linecount.txt`")
	})
	t.Run("cat | wc -l이면 통과(mac wc의 앞 공백 포함)", func(t *testing.T) {
		dir := prepare(t, ex)
		sh(t, dir, "cat TIL/*.md | wc -l > linecount.txt")
		mustPass(t, ex, dir)
	})
	t.Run("wc -l 파일들이면 total 안내", func(t *testing.T) {
		dir := prepare(t, ex)
		sh(t, dir, "wc -l TIL/*.md > linecount.txt")
		mustFail(t, ex, dir, "합계는 맞게")
	})
	t.Run("한 파일만 세면 숫자가 다르다고 안내", func(t *testing.T) {
		dir := prepare(t, ex)
		sh(t, dir, "cat TIL/10-12.md | wc -l > linecount.txt")
		mustFail(t, ex, dir, "숫자가 다릅니다")
	})
}
