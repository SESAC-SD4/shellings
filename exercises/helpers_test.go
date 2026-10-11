package exercises

import (
	"os"
	"os/exec"
	"strings"
	"testing"

	"shellings-ko/internal/exercise"
)

func prepare(t *testing.T, ex *exercise.Exercise) string {
	t.Helper()
	if ex == nil {
		t.Fatal("문제가 등록되지 않았습니다")
	}
	dir := t.TempDir()
	if err := ex.Setup(dir); err != nil {
		t.Fatalf("setup 실패: %v", err)
	}
	return dir
}

// 학생이 치는 것과 같은 진짜 명령으로 풀이를 재현한다.
func sh(t *testing.T, dir, script string) {
	t.Helper()
	cmd := exec.Command("bash", "-c", script)
	cmd.Dir = dir
	cmd.Env = append(os.Environ(), "ROOT="+dir)
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("%q 실패: %v\n%s", script, err, out)
	}
}

func mustPass(t *testing.T, ex *exercise.Exercise, dir string) {
	t.Helper()
	if r := ex.Check(dir); !r.Passed {
		t.Fatalf("통과해야 하는데 실패: %s", r.Message)
	}
}

func mustFail(t *testing.T, ex *exercise.Exercise, dir, want string) {
	t.Helper()
	r := ex.Check(dir)
	if r.Passed {
		t.Fatalf("실패해야 하는데 통과: %s", r.Message)
	}
	if !strings.Contains(r.Message, want) {
		t.Fatalf("메시지에 %q가 없습니다: %s", want, r.Message)
	}
}
