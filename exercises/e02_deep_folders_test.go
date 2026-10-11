package exercises

import "testing"

func TestDeepFolders(t *testing.T) {
	ex := ByID(2)
	t.Run("setup 직후에는 실패", func(t *testing.T) {
		mustFail(t, ex, prepare(t, ex), "아직 없습니다")
	})
	t.Run("mkdir -p면 통과", func(t *testing.T) {
		dir := prepare(t, ex)
		sh(t, dir, "mkdir -p TIL/2026-10/week2/day1")
		mustPass(t, ex, dir)
	})
	t.Run("TIL 밖에 만들면 위치 안내", func(t *testing.T) {
		dir := prepare(t, ex)
		sh(t, dir, "mkdir -p 2026-10/week2/day1")
		mustFail(t, ex, dir, "`TIL` 폴더 밖")
	})
}
