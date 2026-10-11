package exercises

import "testing"

func TestGrep(t *testing.T) {
	ex := ByID(15)
	t.Run("setup 직후에는 실패", func(t *testing.T) {
		mustFail(t, ex, prepare(t, ex), "2개입니다")
	})
	t.Run("grep -l 결과를 복사하면 통과", func(t *testing.T) {
		dir := prepare(t, ex)
		sh(t, dir, `for f in $(grep -l 포트폴리오 TIL/*.md); do cp "$f" portfolio/; done`)
		mustPass(t, ex, dir)
	})
	t.Run("와일드카드 없이 grep -rl로 찾아도 통과", func(t *testing.T) {
		dir := prepare(t, ex)
		sh(t, dir, `for f in $(grep -rl 포트폴리오 TIL); do cp "$f" portfolio/; done`)
		mustPass(t, ex, dir)
	})
	t.Run("전부 복사하면 안내", func(t *testing.T) {
		dir := prepare(t, ex)
		sh(t, dir, "cp TIL/*.md portfolio/")
		mustFail(t, ex, dir, "단어가 없는 글")
	})
	t.Run("하나만 복사하면 개수 안내", func(t *testing.T) {
		dir := prepare(t, ex)
		sh(t, dir, "cp TIL/10-14.md portfolio/")
		mustFail(t, ex, dir, "1개뿐")
	})
	t.Run("mv면 원본 안내", func(t *testing.T) {
		dir := prepare(t, ex)
		sh(t, dir, "mv TIL/10-14.md TIL/10-16.md portfolio/")
		mustFail(t, ex, dir, "원본")
	})
}
