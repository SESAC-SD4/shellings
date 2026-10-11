package exercises

import "testing"

func TestFind(t *testing.T) {
	ex := ByID(14)
	t.Run("setup 직후에는 실패", func(t *testing.T) {
		mustFail(t, ex, prepare(t, ex), "아직 없습니다: `filelist.txt`")
	})
	t.Run("find -name이면 통과", func(t *testing.T) {
		dir := prepare(t, ex)
		sh(t, dir, `find . -name "*.md" > filelist.txt`)
		mustPass(t, ex, dir)
	})
	t.Run("다른 파일이 섞이면 안내", func(t *testing.T) {
		dir := prepare(t, ex)
		sh(t, dir, `find . -type f > filelist.txt`)
		mustFail(t, ex, dir, "`.md`가 아닌")
	})
	t.Run("한 폴더만 보면 빠진 파일 안내", func(t *testing.T) {
		dir := prepare(t, ex)
		sh(t, dir, `find notes -name "*.md" > filelist.txt`)
		mustFail(t, ex, dir, "빠진 파일")
	})
}
