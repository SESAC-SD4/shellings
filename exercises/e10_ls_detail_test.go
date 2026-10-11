package exercises

import "testing"

func TestLsDetail(t *testing.T) {
	ex := ByID(10)
	t.Run("setup 직후에는 실패", func(t *testing.T) {
		mustFail(t, ex, prepare(t, ex), "`found.txt`가 없습니다")
	})
	t.Run("폴더 안에 만들면 통과", func(t *testing.T) {
		dir := prepare(t, ex)
		sh(t, dir, "ls -l mystery && touch mystery/notes2/found.txt")
		mustPass(t, ex, dir)
	})
	t.Run("폴더 밖에 만들면 위치 안내", func(t *testing.T) {
		dir := prepare(t, ex)
		sh(t, dir, "touch mystery/found.txt")
		mustFail(t, ex, dir, "`mystery/found.txt`")
	})
}
