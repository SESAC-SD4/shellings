package exercises

import "testing"

func TestAbsolutePath(t *testing.T) {
	ex := ByID(13)
	t.Run("setup 직후에는 실패", func(t *testing.T) {
		mustFail(t, ex, prepare(t, ex), "아직 없습니다: `portfolio/where.txt`")
	})
	t.Run("week1 안에서 pwd를 전체 경로 파일로 보내면 통과", func(t *testing.T) {
		dir := prepare(t, ex)
		sh(t, dir, `cd TIL/2026-10/week1 && pwd > "$ROOT/portfolio/where.txt"`)
		mustPass(t, ex, dir)
	})
	t.Run("다른 위치에서 pwd를 치면 안내", func(t *testing.T) {
		dir := prepare(t, ex)
		sh(t, dir, `pwd > portfolio/where.txt`)
		mustFail(t, ex, dir, "`week1`이 아닙니다")
	})
	t.Run("다른 문제 폴더의 week1에서 친 pwd면 안내", func(t *testing.T) {
		dir := prepare(t, ex)
		sh(t, dir, `echo "/c/Users/hong/shellings-work/17-backup/TIL/2026-10/week1" > portfolio/where.txt`)
		mustFail(t, ex, dir, "`week1`이 아닙니다")
	})
	t.Run("week1 안에 저장하면 위치 안내", func(t *testing.T) {
		dir := prepare(t, ex)
		sh(t, dir, `cd TIL/2026-10/week1 && pwd > where.txt`)
		mustFail(t, ex, dir, "`TIL/2026-10/week1` 안에")
	})
}
