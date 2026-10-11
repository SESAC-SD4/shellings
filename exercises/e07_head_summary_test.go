package exercises

import "testing"

func TestHeadSummary(t *testing.T) {
	ex := ByID(7)
	t.Run("setup 직후에는 실패", func(t *testing.T) {
		mustFail(t, ex, prepare(t, ex), "아직 없습니다: `portfolio/summary.md`")
	})
	t.Run("head -3이면 통과", func(t *testing.T) {
		dir := prepare(t, ex)
		sh(t, dir, "head -3 TIL/10-12.md > portfolio/summary.md")
		mustPass(t, ex, dir)
	})
	t.Run("파일 전체면 안내", func(t *testing.T) {
		dir := prepare(t, ex)
		sh(t, dir, "cat TIL/10-12.md > portfolio/summary.md")
		mustFail(t, ex, dir, "전체")
	})
	t.Run("뒤 3줄이면 안내", func(t *testing.T) {
		dir := prepare(t, ex)
		sh(t, dir, "tail -3 TIL/10-12.md > portfolio/summary.md")
		mustFail(t, ex, dir, "뒷부분")
	})
}
