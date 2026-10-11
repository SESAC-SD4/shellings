package exercises

import "testing"

func TestOverwrite(t *testing.T) {
	ex := ByID(12)
	t.Run("setup 직후에는 실패", func(t *testing.T) {
		mustFail(t, ex, prepare(t, ex), "사고가 난 상태")
	})
	t.Run("cp로 되살리면 Git 안내와 함께 통과", func(t *testing.T) {
		dir := prepare(t, ex)
		sh(t, dir, "cp backup/10-12.md TIL/10-12.md")
		r := ex.Check(dir)
		if !r.Passed || r.Message != gitBridge {
			t.Fatalf("통과 메시지는 Git 안내여야 합니다: %+v", r)
		}
	})
	t.Run("mv로 되살리면 통과하되 백업이 사라졌다고 덧붙임", func(t *testing.T) {
		dir := prepare(t, ex)
		sh(t, dir, "mv backup/10-12.md TIL/10-12.md")
		r := ex.Check(dir)
		if !r.Passed || r.Message == gitBridge {
			t.Fatalf("통과하되 백업 안내가 붙어야 합니다: %+v", r)
		}
	})
	t.Run("다시 타이핑하면 안내", func(t *testing.T) {
		dir := prepare(t, ex)
		sh(t, dir, `echo "## 터미널" > TIL/10-12.md`)
		mustFail(t, ex, dir, "백업과 다릅니다")
	})
}
