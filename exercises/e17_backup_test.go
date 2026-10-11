package exercises

import "testing"

func TestBackup(t *testing.T) {
	ex := ByID(17)
	t.Run("setup 직후에는 실패", func(t *testing.T) {
		mustFail(t, ex, prepare(t, ex), "아직 없습니다: `backup`")
	})
	t.Run("cp -r이면 통과", func(t *testing.T) {
		dir := prepare(t, ex)
		sh(t, dir, "cp -r TIL backup")
		mustPass(t, ex, dir)
	})
	t.Run("파일만 복사하면 폴더가 빠졌다고 안내", func(t *testing.T) {
		dir := prepare(t, ex)
		sh(t, dir, "mkdir backup && cp TIL/10-12.md TIL/10-13.md TIL/10-14.md backup/")
		mustFail(t, ex, dir, "`backup/2026-10/week1/day1.md`")
	})
	t.Run("두 번 복사해 한 겹 더 생기면 안내", func(t *testing.T) {
		dir := prepare(t, ex)
		sh(t, dir, "cp -r TIL backup && cp -r TIL backup")
		mustFail(t, ex, dir, "한 겹 더")
	})
	t.Run("mv면 원본 안내", func(t *testing.T) {
		dir := prepare(t, ex)
		sh(t, dir, "mv TIL backup")
		mustFail(t, ex, dir, "원본")
	})
}
