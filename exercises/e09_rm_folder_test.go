package exercises

import "testing"

func TestRmFolder(t *testing.T) {
	ex := ByID(9)
	t.Run("setup 직후에는 실패", func(t *testing.T) {
		mustFail(t, ex, prepare(t, ex), "`trash` 폴더가 아직 있습니다")
	})
	t.Run("rm -r이면 통과", func(t *testing.T) {
		dir := prepare(t, ex)
		sh(t, dir, "rm -r trash")
		mustPass(t, ex, dir)
	})
	t.Run("안만 비우면 폴더가 남았다고 안내", func(t *testing.T) {
		dir := prepare(t, ex)
		sh(t, dir, "rm trash/*")
		mustFail(t, ex, dir, "폴더가 남아")
	})
	t.Run("TIL까지 지우면 안내", func(t *testing.T) {
		dir := prepare(t, ex)
		sh(t, dir, "rm -r trash TIL")
		mustFail(t, ex, dir, "`trash`만 지워야")
	})
}
