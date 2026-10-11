package exercises

import "testing"

func TestLsToFile(t *testing.T) {
	ex := ByID(14)
	t.Run("setup 직후에는 실패", func(t *testing.T) {
		mustFail(t, ex, prepare(t, ex), "아직 없습니다: `filelist.txt`")
	})
	t.Run("ls TIL >면 통과", func(t *testing.T) {
		dir := prepare(t, ex)
		sh(t, dir, "ls TIL > filelist.txt")
		mustPass(t, ex, dir)
	})
	t.Run("맨 위 목록을 저장하면 안내", func(t *testing.T) {
		dir := prepare(t, ex)
		sh(t, dir, "ls > filelist.txt")
		mustFail(t, ex, dir, "`TIL` 폴더 안")
	})
	t.Run("TIL 안에 저장하면 위치 안내", func(t *testing.T) {
		dir := prepare(t, ex)
		sh(t, dir, "cd TIL && ls > filelist.txt")
		mustFail(t, ex, dir, "`TIL` 안에 생겼습니다")
	})
}
