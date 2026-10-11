package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"

	"shellings-ko/internal/exercise"
	"shellings-ko/internal/progress"
	"shellings-ko/internal/workspace"
)

type Renderer func(markdown string, width int) string

type Mode int

const (
	ModePrompt Mode = iota
	ModeHint
	ModeConfirm
)

const readyStatus = "s로 연습 폴더를 만들고, 다른 터미널에서 풀어 보세요"

type Model struct {
	exs       []*exercise.Exercise
	prog      *progress.Progress
	render    Renderer
	cursor    int
	mode      Mode
	hintLevel int
	status    string
	banner    string
	width     int
	height    int
	vp        viewport.Model
}

func New(exs []*exercise.Exercise, prog *progress.Progress, render Renderer) Model {
	m := Model{exs: exs, prog: prog, render: render, vp: viewport.New(0, 0), status: readyStatus}
	for i, ex := range exs {
		if !prog.IsPassed(ex.ID) {
			m.cursor = i
			break
		}
	}
	m.refresh()
	return m
}

func (m Model) Cursor() int             { return m.cursor }
func (m Model) Mode() Mode              { return m.mode }
func (m Model) HintLevel() int          { return m.hintLevel }
func (m Model) Status() string          { return m.status }
func (m Model) Banner() string          { return m.banner }
func (m Model) Init() tea.Cmd           { return nil }
func (m Model) View() string            { return m.Content() }
func (m Model) cur() *exercise.Exercise { return m.exs[m.cursor] }

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
		m.layout()
		m.refresh()
		return m, nil
	case tea.KeyMsg:
		return m.handleKey(msg)
	}
	var cmd tea.Cmd
	m.vp, cmd = m.vp.Update(msg)
	return m, cmd
}

func (m Model) handleKey(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	k := msg.String()
	if k == "ctrl+c" {
		return m, tea.Quit
	}
	if m.mode == ModeConfirm {
		m.mode = ModePrompt
		if k == "y" {
			m.prepare()
		} else {
			m.status = "취소했습니다. 연습 폴더는 그대로입니다"
		}
		m.refresh()
		return m, nil
	}
	if m.mode == ModeHint && k != "h" {
		m.mode = ModePrompt
	}
	switch k {
	case "q":
		return m, tea.Quit
	case "up", "k":
		m.move(-1)
	case "down", "j":
		m.move(1)
	case "s", "r":
		m.askOrPrepare()
	case "c":
		m.check()
	case "h":
		m.hint()
	default:
		var cmd tea.Cmd
		m.vp, cmd = m.vp.Update(msg)
		m.refresh()
		return m, cmd
	}
	m.refresh()
	return m, nil
}

func (m *Model) move(d int) {
	n := m.cursor + d
	if n < 0 || n >= len(m.exs) {
		return
	}
	m.cursor, m.hintLevel, m.banner, m.mode, m.status = n, 0, "", ModePrompt, readyStatus
}

func (m *Model) askOrPrepare() {
	if workspace.Exists(m.cur().Slug) {
		m.mode = ModeConfirm
		m.status = "이미 연습 폴더가 있습니다. 처음 상태로 덮어쓸까요? (y/n)"
		return
	}
	m.prepare()
}

func (m *Model) prepare() {
	ex := m.cur()
	if err := workspace.Prepare(ex); err != nil {
		m.status = "❌ 연습 폴더를 만들지 못했습니다: " + err.Error()
		return
	}
	m.prog.LastStarted = ex.ID
	m.prog.ResetHint(ex.ID)
	m.hintLevel, m.banner = 0, ""
	if err := m.prog.Save(); err != nil {
		m.status = "❌ 진행 상황을 저장하지 못했습니다: " + err.Error()
		return
	}
	m.status = "✅ 연습 폴더를 만들었습니다. 다른 터미널에서: cd " + workspace.Display(ex.Slug)
}

func (m *Model) check() {
	ex := m.cur()
	if !workspace.Exists(ex.Slug) {
		m.status = "먼저 s를 눌러 연습 폴더를 만드세요"
		return
	}
	r := ex.Check(workspace.Dir(ex.Slug))
	if !r.Passed {
		m.banner = "❌ 아직입니다 — " + r.Message
		m.status = "아직입니다. h로 힌트를 볼 수 있습니다"
		return
	}
	m.prog.MarkPassed(ex.ID)
	if err := m.prog.Save(); err != nil {
		m.status = "❌ 진행 상황을 저장하지 못했습니다: " + err.Error()
		return
	}
	m.banner = "✅ " + ex.Title + " 통과! " + r.Message
	m.status = "통과!"
	if m.cursor < len(m.exs)-1 {
		m.cursor++
		m.hintLevel = 0
	}
}

func (m *Model) hint() {
	ex := m.cur()
	if m.hintLevel < len(ex.Hints) {
		m.hintLevel++
	}
	m.mode = ModeHint
	m.status = fmt.Sprintf("힌트 %d/%d — h를 다시 누르면 다음 단계, 다른 키는 문제로 돌아갑니다", m.hintLevel, len(ex.Hints))
}

func (m Model) Content() string {
	ex := m.cur()
	var b strings.Builder
	if m.banner != "" {
		b.WriteString("> " + strings.ReplaceAll(m.banner, "\n", "\n> ") + "\n\n")
	}
	b.WriteString("# " + ex.Title + "\n\n")
	if m.mode == ModeHint {
		for i := 0; i < m.hintLevel; i++ {
			fmt.Fprintf(&b, "**힌트 %d/%d**\n\n%s\n\n", i+1, len(ex.Hints), ex.Hints[i])
		}
	} else {
		b.WriteString(strings.TrimSpace(ex.Prompt) + "\n\n")
	}
	b.WriteString("---\n\n연습 폴더: `cd " + workspace.Display(ex.Slug) + "`\n")
	if !workspace.Exists(ex.Slug) {
		b.WriteString("\n아직 만들지 않았습니다. `s`를 누르세요.\n")
	}
	return b.String()
}

// layout은 Task 9에서 창 크기에 맞춰 칸 너비를 정한다.
func (m *Model) layout() {}

func (m *Model) refresh() {
	m.vp.SetContent(m.render(m.Content(), m.vp.Width))
	m.vp.GotoTop()
}
