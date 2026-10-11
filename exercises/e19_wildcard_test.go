package exercises

import "testing"

func TestWildcard(t *testing.T) {
	ex := ByID(19)
	t.Run("setup 직후에는 실패", func(t *testing.T) {
		mustFail(t, ex, prepare(t, ex), "맨 위에 남아 있습니다")
	})
	t.Run("*로 한꺼번에 옮기면 통과", func(t *testing.T) {
		dir := prepare(t, ex)
		sh(t, dir, "mv *.png images/ && mv *.md TIL/")
		mustPass(t, ex, dir)
	})
	t.Run("한쪽만 옮기면 남았다고 안내", func(t *testing.T) {
		dir := prepare(t, ex)
		sh(t, dir, "mv *.png images/")
		mustFail(t, ex, dir, "맨 위에 남아 있습니다")
	})
	t.Run("섞어서 넣으면 안내", func(t *testing.T) {
		dir := prepare(t, ex)
		sh(t, dir, "mv *.png *.md TIL/")
		mustFail(t, ex, dir, "`.png` 파일이 `TIL`에")
	})
}
