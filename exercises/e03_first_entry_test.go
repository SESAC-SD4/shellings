package exercises

import "testing"

func TestFirstEntry(t *testing.T) {
	ex := ByID(3)
	t.Run("setup 직후에는 실패", func(t *testing.T) {
		mustFail(t, ex, prepare(t, ex), "아직 없습니다: `TIL/10-12.md`")
	})
	t.Run("echo >이면 통과", func(t *testing.T) {
		dir := prepare(t, ex)
		sh(t, dir, `echo "## 터미널" > TIL/10-12.md`)
		mustPass(t, ex, dir)
	})
	t.Run("touch로 빈 파일이면 안내", func(t *testing.T) {
		dir := prepare(t, ex)
		sh(t, dir, "touch TIL/10-12.md")
		mustFail(t, ex, dir, "비어 있습니다")
	})
	t.Run("TIL 밖에 만들면 위치 안내", func(t *testing.T) {
		dir := prepare(t, ex)
		sh(t, dir, `echo "## 터미널" > 10-12.md`)
		mustFail(t, ex, dir, "`TIL` 폴더 밖")
	})
}
