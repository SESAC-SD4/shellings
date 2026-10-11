package exercises

import "testing"

func TestRename(t *testing.T) {
	ex := ByID(6)
	t.Run("setup 직후에는 실패", func(t *testing.T) {
		mustFail(t, ex, prepare(t, ex), "아직 없습니다: `TIL/10-14.md`")
	})
	t.Run("mv면 통과", func(t *testing.T) {
		dir := prepare(t, ex)
		sh(t, dir, "mv TIL/10_14.md TIL/10-14.md")
		mustPass(t, ex, dir)
	})
	t.Run("cp면 오타 파일이 남았다고 안내", func(t *testing.T) {
		dir := prepare(t, ex)
		sh(t, dir, "cp TIL/10_14.md TIL/10-14.md")
		mustFail(t, ex, dir, "오타 이름")
	})
}
