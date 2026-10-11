package exercises

import "testing"

func TestCleanUp(t *testing.T) {
	ex := ByID(10)
	t.Run("setup 직후에는 실패", func(t *testing.T) {
		mustFail(t, ex, prepare(t, ex), "아직 남아 있습니다")
	})
	t.Run("두 파일만 지우면 통과", func(t *testing.T) {
		dir := prepare(t, ex)
		sh(t, dir, "rm TIL/temp-note.md TIL/draft-old.md")
		mustPass(t, ex, dir)
	})
	t.Run("기록까지 지우면 안내", func(t *testing.T) {
		dir := prepare(t, ex)
		sh(t, dir, "rm TIL/*.md")
		mustFail(t, ex, dir, "지우면 안 되는")
	})
}
