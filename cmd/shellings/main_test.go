package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"shellings-ko/internal/progress"
	"shellings-ko/internal/workspace"
)

func call(t *testing.T, input string, args ...string) (int, string) {
	t.Helper()
	var out bytes.Buffer
	code := run(args, strings.NewReader(input), &out)
	return code, out.String()
}

func setHome(t *testing.T) {
	t.Helper()
	t.Setenv("HOME", t.TempDir())
	t.Chdir(t.TempDir())
}

func TestList(t *testing.T) {
	setHome(t)
	code, out := call(t, "", "list")
	if code != 0 || !strings.Contains(out, "⬜ 01 TIL 폴더") || !strings.Contains(out, "0/22 통과") {
		t.Fatalf("목록이 나와야 합니다: %d\n%s", code, out)
	}
}

func TestStart(t *testing.T) {
	setHome(t)
	code, out := call(t, "", "start", "1")
	if code != 0 || !workspace.Exists("01-make-til") || !strings.Contains(out, "cd ~/shellings-work/01-make-til") || !strings.Contains(out, "## 할 일") {
		t.Fatalf("연습 폴더를 만들고 문제와 이동 경로를 보여 줘야 합니다: %d\n%s", code, out)
	}
	p, _ := progress.Load()
	if p.LastStarted != 1 {
		t.Fatal("last_started가 저장되어야 합니다")
	}
}

func TestStartBadNumber(t *testing.T) {
	setHome(t)
	for _, args := range [][]string{{"start", "99"}, {"start", "abc"}, {"start"}} {
		if code, _ := call(t, "", args...); code != 1 {
			t.Errorf("%v는 실패해야 합니다", args)
		}
	}
}

func TestStartAsksBeforeOverwrite(t *testing.T) {
	setHome(t)
	call(t, "", "start", "1")
	marker := filepath.Join(workspace.Dir("01-make-til"), "mine.txt")
	if err := os.WriteFile(marker, []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	_, out := call(t, "n\n", "start", "1")
	if _, err := os.Stat(marker); err != nil || !strings.Contains(out, "그대로 두었습니다") {
		t.Fatal("n이면 덮어쓰지 않아야 합니다")
	}
	call(t, "y\n", "start", "1")
	if _, err := os.Stat(marker); !os.IsNotExist(err) {
		t.Fatal("y면 처음 상태로 덮어써야 합니다")
	}
}

func TestCheckFromInsideAndOutside(t *testing.T) {
	setHome(t)
	if code, out := call(t, "", "check"); code != 1 || !strings.Contains(out, "start N") {
		t.Fatalf("시작 전에는 start를 안내해야 합니다: %s", out)
	}
	call(t, "", "start", "1")
	if code, out := call(t, "", "check"); code != 1 || !strings.Contains(out, "아직입니다") {
		t.Fatalf("풀기 전에는 실패: %s", out)
	}
	if err := os.Mkdir(filepath.Join(workspace.Dir("01-make-til"), "TIL"), 0o755); err != nil {
		t.Fatal(err)
	}
	t.Chdir(workspace.Dir("01-make-til"))
	code, out := call(t, "", "check")
	if code != 0 || !strings.Contains(out, "통과") || !strings.Contains(out, "start 2") {
		t.Fatalf("통과하고 다음 문제를 안내해야 합니다: %s", out)
	}
	p, _ := progress.Load()
	if !p.IsPassed(1) {
		t.Fatal("통과가 저장되어야 합니다")
	}
}

func TestHintSteps(t *testing.T) {
	setHome(t)
	call(t, "", "start", "1")
	_, out := call(t, "", "hint")
	if !strings.Contains(out, "힌트 1/3") || strings.Contains(out, "힌트 2/3") {
		t.Fatalf("처음에는 1단계만: %s", out)
	}
	call(t, "", "hint")
	call(t, "", "hint")
	_, out = call(t, "", "hint")
	if !strings.Contains(out, "힌트 3/3") || !strings.Contains(out, "마지막 힌트") {
		t.Fatalf("3단계에서 멈춰야 합니다: %s", out)
	}
}

func TestReset(t *testing.T) {
	setHome(t)
	call(t, "", "start", "1")
	marker := filepath.Join(workspace.Dir("01-make-til"), "mine.txt")
	if err := os.WriteFile(marker, []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	if code, out := call(t, "", "reset"); code != 0 || !strings.Contains(out, "처음 상태") {
		t.Fatalf("reset이 성공해야 합니다: %s", out)
	}
	if _, err := os.Stat(marker); !os.IsNotExist(err) {
		t.Fatal("reset은 처음 상태로 되돌려야 합니다")
	}
}

func TestUnknownAndHelp(t *testing.T) {
	setHome(t)
	if code, out := call(t, "", "fly"); code != 1 || !strings.Contains(out, "모르는 명령") {
		t.Fatalf("모르는 명령은 실패: %s", out)
	}
	if code, out := call(t, "", "help"); code != 0 || !strings.Contains(out, "watch") {
		t.Fatalf("help에 명령 목록이 있어야 합니다: %s", out)
	}
}
