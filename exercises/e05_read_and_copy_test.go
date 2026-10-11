package exercises

import "testing"

func TestReadAndCopy(t *testing.T) {
	ex := ByID(5)
	t.Run("setup 직후에는 실패", func(t *testing.T) {
		mustFail(t, ex, prepare(t, ex), "`index.txt`")
	})
	t.Run("cat 후 cp면 통과", func(t *testing.T) {
		dir := prepare(t, ex)
		sh(t, dir, "cat index.txt && cp TIL/10-13.md portfolio/")
		mustPass(t, ex, dir)
	})
	t.Run("다른 글을 복사하면 안내", func(t *testing.T) {
		dir := prepare(t, ex)
		sh(t, dir, "cp TIL/10-12.md portfolio/")
		mustFail(t, ex, dir, "다른 글")
	})
	t.Run("mv로 옮기면 원본 안내", func(t *testing.T) {
		dir := prepare(t, ex)
		sh(t, dir, "mv TIL/10-13.md portfolio/")
		mustFail(t, ex, dir, "원본")
	})
}
