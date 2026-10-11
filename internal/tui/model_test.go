package tui

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"

	"shellings-ko/exercises"
	"shellings-ko/internal/progress"
	"shellings-ko/internal/workspace"
)

func plain(md string, width int) string { return md }

func newModel(t *testing.T, prog *progress.Progress) Model {
	t.Helper()
	t.Setenv("HOME", t.TempDir())
	if prog == nil {
		prog = &progress.Progress{}
	}
	return New(exercises.All(), prog, plain)
}

func key(k string) tea.KeyMsg {
	switch k {
	case "up":
		return tea.KeyMsg{Type: tea.KeyUp}
	case "down":
		return tea.KeyMsg{Type: tea.KeyDown}
	case "ctrl+c":
		return tea.KeyMsg{Type: tea.KeyCtrlC}
	}
	return tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune(k)}
}

func press(m Model, keys ...string) Model {
	var tm tea.Model = m
	for _, k := range keys {
		tm, _ = tm.Update(key(k))
	}
	return tm.(Model)
}

func TestStartsAtFirstUnpassed(t *testing.T) {
	m := newModel(t, &progress.Progress{Passed: []int{1, 2}})
	if m.Cursor() != 2 {
		t.Fatalf("3번 문제(인덱스 2)에서 시작해야 합니다: %d", m.Cursor())
	}
}

func TestMoveStaysInBounds(t *testing.T) {
	m := newModel(t, nil)
	m = press(m, "k", "up")
	if m.Cursor() != 0 {
		t.Fatalf("맨 위에서 더 올라가면 안 됩니다: %d", m.Cursor())
	}
	for i := 0; i < 30; i++ {
		m = press(m, "j")
	}
	if m.Cursor() != len(exercises.All())-1 {
		t.Fatalf("맨 아래에서 멈춰야 합니다: %d", m.Cursor())
	}
	m = press(m, "down", "k")
	if m.Cursor() != len(exercises.All())-2 {
		t.Fatalf("down과 k가 동작해야 합니다: %d", m.Cursor())
	}
}

func TestCheckBeforeStart(t *testing.T) {
	m := press(newModel(t, nil), "c")
	if !strings.Contains(m.Status(), "먼저 s") {
		t.Fatalf("연습 폴더부터 만들라고 해야 합니다: %s", m.Status())
	}
}

func TestStartCreatesWorkspace(t *testing.T) {
	m := press(newModel(t, nil), "s")
	if !workspace.Exists("01-make-til") {
		t.Fatal("연습 폴더가 생겨야 합니다")
	}
	if !strings.Contains(m.Status(), "cd ~/shellings-work/01-make-til") {
		t.Fatalf("어디로 가야 하는지 보여 줘야 합니다: %s", m.Status())
	}
	if !strings.Contains(m.Content(), "cd ~/shellings-work/01-make-til") {
		t.Fatal("오른쪽 칸에도 연습 폴더 경로가 있어야 합니다")
	}
}

func TestStartTwiceAsksBeforeOverwriting(t *testing.T) {
	m := press(newModel(t, nil), "s")
	marker := filepath.Join(workspace.Dir("01-make-til"), "mine.txt")
	if err := os.WriteFile(marker, []byte("내 작업"), 0o644); err != nil {
		t.Fatal(err)
	}
	m = press(m, "s")
	if m.Mode() != ModeConfirm || !strings.Contains(m.Status(), "(y/n)") {
		t.Fatalf("덮어쓸지 물어야 합니다: %v %s", m.Mode(), m.Status())
	}
	m = press(m, "n")
	if _, err := os.Stat(marker); err != nil || !strings.Contains(m.Status(), "취소") {
		t.Fatal("n이면 학생 작업을 지키고 취소했다고 알려야 합니다")
	}
	m = press(m, "r", "y")
	if _, err := os.Stat(marker); !os.IsNotExist(err) {
		t.Fatal("r 다음 y면 처음 상태로 되돌려야 합니다")
	}
	if m.Mode() != ModePrompt {
		t.Fatal("확인이 끝나면 문제 화면으로 돌아가야 합니다")
	}
}

func TestCheckFailThenPass(t *testing.T) {
	prog := &progress.Progress{}
	m := press(newModel(t, prog), "s", "c")
	if !strings.Contains(m.Banner(), "아직입니다") || m.Cursor() != 0 {
		t.Fatalf("실패하면 그 자리에서 이유를 보여 줘야 합니다: %s", m.Banner())
	}
	if err := os.Mkdir(filepath.Join(workspace.Dir("01-make-til"), "TIL"), 0o755); err != nil {
		t.Fatal(err)
	}
	m = press(m, "c")
	if !strings.Contains(m.Banner(), "통과") || m.Cursor() != 1 {
		t.Fatalf("통과하면 배너를 남기고 다음 문제로 가야 합니다: %s %d", m.Banner(), m.Cursor())
	}
	if !strings.Contains(m.Content(), "통과") {
		t.Fatal("다음 문제 화면 위에 통과 배너가 보여야 합니다")
	}
	saved, err := progress.Load()
	if err != nil || !saved.IsPassed(1) {
		t.Fatal("통과 기록이 파일에 저장되어야 합니다")
	}
	m = press(m, "j")
	if m.Banner() != "" {
		t.Fatal("직접 이동하면 배너를 지워야 합니다")
	}
}

func TestHintLevels(t *testing.T) {
	m := press(newModel(t, nil), "h")
	if m.Mode() != ModeHint || m.HintLevel() != 1 || !strings.Contains(m.Content(), "힌트 1/3") {
		t.Fatalf("h 한 번이면 1단계: %v %d", m.Mode(), m.HintLevel())
	}
	m = press(m, "h", "h", "h")
	if m.HintLevel() != 3 || !strings.Contains(m.Content(), "힌트 3/3") || !strings.Contains(m.Content(), "힌트 1/3") {
		t.Fatalf("3단계에서 멈추고 앞 단계도 함께 보여야 합니다: %d", m.HintLevel())
	}
	m = press(m, "x")
	if m.Mode() != ModePrompt || m.HintLevel() != 3 {
		t.Fatal("다른 키면 문제 설명으로 돌아가되 단계는 기억해야 합니다")
	}
	m = press(m, "j")
	if m.HintLevel() != 0 {
		t.Fatal("문제를 바꾸면 힌트 단계는 처음으로")
	}
}

func TestQuit(t *testing.T) {
	m := newModel(t, nil)
	for _, k := range []string{"q", "ctrl+c"} {
		_, cmd := m.Update(key(k))
		if cmd == nil {
			t.Fatalf("%s는 종료 명령을 돌려줘야 합니다", k)
		}
		if _, ok := cmd().(tea.QuitMsg); !ok {
			t.Fatalf("%s는 tea.Quit이어야 합니다", k)
		}
	}
}
