package tui

import (
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func TestViewShowsListContentAndKeys(t *testing.T) {
	m := newModel(t, nil)
	tm, _ := m.Update(tea.WindowSizeMsg{Width: 120, Height: 40})
	v := tm.View()
	for _, want := range []string{"01 TIL 폴더", "15★ 줄 수 세기", "⬜", ">", "(c)채점", "(h)힌트", "TIL"} {
		if !strings.Contains(v, want) {
			t.Errorf("화면에 %q가 없습니다", want)
		}
	}
	for i, line := range strings.Split(v, "\n") {
		if w := lipgloss.Width(line); w > 120 {
			t.Errorf("%d번째 줄이 창보다 넓습니다: %d", i+1, w)
		}
	}
}

func TestViewBeforeSize(t *testing.T) {
	if v := newModel(t, nil).View(); !strings.Contains(v, "불러오는 중") {
		t.Fatalf("창 크기를 받기 전에는 불러오는 중 표시: %q", v)
	}
}

func TestGlamourRenderer(t *testing.T) {
	for _, dark := range []bool{true, false} {
		out := GlamourRenderer(dark)("# 제목\n\n본문 `cd`", 60)
		if !strings.Contains(out, "제목") || !strings.Contains(out, "본문") {
			t.Fatalf("렌더링 결과에 글자가 남아 있어야 합니다: %q", out)
		}
	}
}
