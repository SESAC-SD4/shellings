package exercises

import "testing"

func TestHiddenDraft(t *testing.T) {
	ex := ByID(4)
	t.Run("setup 직후에는 실패", func(t *testing.T) {
		mustFail(t, ex, prepare(t, ex), "아직 없습니다: `TIL/10-13.md`")
	})
	t.Run("mv로 이름을 바꾸면 통과", func(t *testing.T) {
		dir := prepare(t, ex)
		sh(t, dir, "ls -a TIL && mv TIL/.draft TIL/10-13.md")
		mustPass(t, ex, dir)
	})
	t.Run("rm으로 지우면 안내", func(t *testing.T) {
		dir := prepare(t, ex)
		sh(t, dir, "rm TIL/.draft")
		mustFail(t, ex, dir, "지우지 말고")
	})
	t.Run("새 파일을 만들면 숨김 파일이 남았다고 안내", func(t *testing.T) {
		dir := prepare(t, ex)
		sh(t, dir, "touch TIL/10-13.md")
		mustFail(t, ex, dir, "아직 남아 있습니다")
	})
}
