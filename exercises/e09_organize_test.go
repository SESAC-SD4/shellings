package exercises

import "testing"

func TestOrganize(t *testing.T) {
	ex := ByID(9)
	t.Run("setup 직후에는 실패", func(t *testing.T) {
		mustFail(t, ex, prepare(t, ex), "아직 맨 위 폴더에 있습니다")
	})
	t.Run("mv면 통과", func(t *testing.T) {
		dir := prepare(t, ex)
		sh(t, dir, "mv 10-15.md 10-16.md TIL/ && mv about.md portfolio/")
		mustPass(t, ex, dir)
	})
	t.Run("cp면 옮기기 안내", func(t *testing.T) {
		dir := prepare(t, ex)
		sh(t, dir, "cp 10-15.md 10-16.md TIL/ && cp about.md portfolio/")
		mustFail(t, ex, dir, "복사(`cp`)가 아니라")
	})
	t.Run("엉뚱한 폴더로 옮기면 안내", func(t *testing.T) {
		dir := prepare(t, ex)
		sh(t, dir, "mv 10-15.md 10-16.md about.md TIL/")
		mustFail(t, ex, dir, "`about.md`가 `TIL`에 들어갔습니다")
	})
}
