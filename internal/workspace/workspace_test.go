package workspace

import (
	"os"
	"path/filepath"
	"testing"

	"shellings-ko/internal/exercise"
)

func fake(id int, slug string) *exercise.Exercise {
	return &exercise.Exercise{ID: id, Slug: slug, Setup: func(dir string) error {
		return exercise.WriteFiles(dir, map[string]string{"seed.txt": "씨앗\n"})
	}}
}

func TestBaseFollowsHOME(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)
	if Base() != filepath.Join(home, "shellings-work") {
		t.Fatalf("HOME 아래여야 합니다: %s", Base())
	}
	if Display("01-a") != "~/shellings-work/01-a" {
		t.Fatalf("표시 경로가 다릅니다: %s", Display("01-a"))
	}
}

func TestPrepareRecreates(t *testing.T) {
	t.Setenv("HOME", t.TempDir())
	ex := fake(1, "01-a")
	if Exists("01-a") {
		t.Fatal("처음에는 없어야 합니다")
	}
	if err := Prepare(ex); err != nil {
		t.Fatal(err)
	}
	junk := filepath.Join(Dir("01-a"), "junk.txt")
	if err := os.WriteFile(junk, []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := Prepare(ex); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(junk); !os.IsNotExist(err) {
		t.Fatal("다시 만들면 학생이 만든 파일은 사라져야 합니다")
	}
	if !exercise.IsFile(Dir("01-a"), "seed.txt") {
		t.Fatal("setup이 실행되어야 합니다")
	}
}

func TestPrepareRefusesEmptySlug(t *testing.T) {
	t.Setenv("HOME", t.TempDir())
	if err := Prepare(fake(1, "")); err == nil {
		t.Fatal("slug가 비면 연습 폴더 전체를 지울 위험이 있으므로 거부해야 합니다")
	}
}

func TestDetect(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)
	exs := []*exercise.Exercise{fake(1, "01-a"), fake(2, "02-b")}
	inside := filepath.Join(Base(), "02-b", "TIL", "deep")
	if err := os.MkdirAll(inside, 0o755); err != nil {
		t.Fatal(err)
	}
	if got := Detect(inside, exs, 1); got == nil || got.ID != 2 {
		t.Fatalf("현재 경로의 문제를 골라야 합니다: %v", got)
	}
	if got := Detect(home, exs, 1); got == nil || got.ID != 1 {
		t.Fatalf("밖에서는 last_started를 써야 합니다: %v", got)
	}
	if got := Detect(home, exs, 0); got != nil {
		t.Fatal("아무것도 시작하지 않았으면 nil")
	}
	link := filepath.Join(t.TempDir(), "link")
	if err := os.Symlink(filepath.Join(Base(), "02-b"), link); err != nil {
		t.Fatal(err)
	}
	if got := Detect(link, exs, 1); got == nil || got.ID != 2 {
		t.Fatalf("심볼릭 링크로 들어와도 같은 문제여야 합니다: %v", got)
	}
}
