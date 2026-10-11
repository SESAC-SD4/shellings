package exercises

import "testing"

func TestDiff(t *testing.T) {
	ex := ByID(22)
	t.Run("setup 직후에는 실패", func(t *testing.T) {
		mustFail(t, ex, prepare(t, ex), "아직 없습니다: `changed.txt`")
	})
	t.Run("diff 출력을 그대로 저장하면 통과", func(t *testing.T) {
		dir := prepare(t, ex)
		// diff는 차이가 있으면 종료 코드 1을 돌려준다.
		sh(t, dir, "diff backup/10-13.md TIL/10-13.md > changed.txt; true")
		mustPass(t, ex, dir)
	})
	t.Run("손으로 두 줄을 옮겨 적어도 통과", func(t *testing.T) {
		dir := prepare(t, ex)
		sh(t, dir, `printf '%s\n%s\n' "`+diffNew1+`" "`+diffNew2+`" > changed.txt`)
		mustPass(t, ex, dir)
	})
	t.Run("파일 전체면 안내", func(t *testing.T) {
		dir := prepare(t, ex)
		sh(t, dir, "cp TIL/10-13.md changed.txt")
		mustFail(t, ex, dir, "바뀌지 않은 줄")
	})
	t.Run("옛 줄만 적으면 안내", func(t *testing.T) {
		dir := prepare(t, ex)
		sh(t, dir, `echo "`+diffOld+`" > changed.txt`)
		mustFail(t, ex, dir, "옛 파일")
	})
}
