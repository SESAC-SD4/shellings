package exercises

import "testing"

func TestAddMore(t *testing.T) {
	ex := ByID(3)
	t.Run("setup 직후에는 실패", func(t *testing.T) {
		mustFail(t, ex, prepare(t, ex), "한 줄뿐")
	})
	t.Run("echo >>이면 통과", func(t *testing.T) {
		dir := prepare(t, ex)
		sh(t, dir, `echo "- cd로 폴더를 옮긴다" >> TIL/10-12.md`)
		mustPass(t, ex, dir)
	})
	t.Run("> 로 덮어쓰면 안내", func(t *testing.T) {
		dir := prepare(t, ex)
		sh(t, dir, `echo "- cd로 폴더를 옮긴다" > TIL/10-12.md`)
		mustFail(t, ex, dir, "사라졌습니다")
	})
}
