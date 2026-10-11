package exercises

import "testing"

func TestGrepPipe(t *testing.T) {
	ex := ByID(16)
	t.Run("setup 직후에는 실패", func(t *testing.T) {
		mustFail(t, ex, prepare(t, ex), "아직 없습니다: `october.txt`")
	})
	t.Run("ls | grep이면 통과", func(t *testing.T) {
		dir := prepare(t, ex)
		sh(t, dir, "ls TIL | grep 10- > october.txt")
		mustPass(t, ex, dir)
	})
	t.Run("걸러내지 않으면 9월이 섞였다고 안내", func(t *testing.T) {
		dir := prepare(t, ex)
		sh(t, dir, "ls TIL > october.txt")
		mustFail(t, ex, dir, "9월")
	})
	t.Run("너무 좁게 거르면 빠진 파일 안내", func(t *testing.T) {
		dir := prepare(t, ex)
		sh(t, dir, "ls TIL | grep 10-12 > october.txt")
		mustFail(t, ex, dir, "빠진 10월 파일")
	})
}
