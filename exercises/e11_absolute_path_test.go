package exercises

import "testing"

func TestAbsolutePath(t *testing.T) {
	ex := ByID(11)
	t.Run("setup 직후에는 실패", func(t *testing.T) {
		mustFail(t, ex, prepare(t, ex), "아직 없습니다: `portfolio/README.md`")
	})
	t.Run("깊은 폴더에서 전체 경로로 쓰면 통과", func(t *testing.T) {
		dir := prepare(t, ex)
		sh(t, dir, `cd TIL/2026-10/week1 && echo "# 포트폴리오" > "$ROOT/portfolio/README.md"`)
		mustPass(t, ex, dir)
	})
	t.Run("week1 안에 만들면 위치 안내", func(t *testing.T) {
		dir := prepare(t, ex)
		sh(t, dir, `cd TIL/2026-10/week1 && echo "# 포트폴리오" > README.md`)
		mustFail(t, ex, dir, "`TIL/2026-10/week1` 안에")
	})
	t.Run("빈 파일이면 안내", func(t *testing.T) {
		dir := prepare(t, ex)
		sh(t, dir, "touch portfolio/README.md")
		mustFail(t, ex, dir, "비어 있습니다")
	})
}
