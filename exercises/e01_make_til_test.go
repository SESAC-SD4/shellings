package exercises

import "testing"

func TestMakeTil(t *testing.T) {
	ex := ByID(1)
	t.Run("setup 직후에는 실패", func(t *testing.T) {
		mustFail(t, ex, prepare(t, ex), "아직 없습니다")
	})
	t.Run("mkdir TIL이면 통과", func(t *testing.T) {
		dir := prepare(t, ex)
		sh(t, dir, "mkdir TIL")
		mustPass(t, ex, dir)
	})
	t.Run("파일로 만들면 안내", func(t *testing.T) {
		dir := prepare(t, ex)
		sh(t, dir, "touch TIL")
		mustFail(t, ex, dir, "폴더가 아니라 파일")
	})
	t.Run("소문자 til이면 대소문자 안내", func(t *testing.T) {
		dir := prepare(t, ex)
		sh(t, dir, "mkdir til")
		mustFail(t, ex, dir, "대소문자")
	})
}
