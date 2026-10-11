package workspace

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	"shellings-ko/internal/exercise"
)

// Git Bash에서 HOME을 따로 정한 학생도 `cd ~/shellings-work`와 같은 곳을 보도록 HOME을 먼저 쓴다.
func Base() string {
	home := os.Getenv("HOME")
	if home == "" {
		home, _ = os.UserHomeDir()
	}
	return filepath.Join(home, "shellings-work")
}

func Dir(slug string) string     { return filepath.Join(Base(), slug) }
func Display(slug string) string { return "~/shellings-work/" + slug }

func Exists(slug string) bool {
	info, err := os.Stat(Dir(slug))
	return err == nil && info.IsDir()
}

func Prepare(ex *exercise.Exercise) error {
	if ex.Slug == "" {
		return errors.New("문제 slug가 비어 있습니다")
	}
	d := Dir(ex.Slug)
	if err := os.RemoveAll(d); err != nil {
		return err
	}
	if err := os.MkdirAll(d, 0o755); err != nil {
		return err
	}
	return ex.Setup(d)
}

func real(p string) string {
	if r, err := filepath.EvalSymlinks(p); err == nil {
		return r
	}
	return p
}

func Detect(cwd string, exs []*exercise.Exercise, lastStarted int) *exercise.Exercise {
	if rel, err := filepath.Rel(real(Base()), real(cwd)); err == nil && rel != "." && !strings.HasPrefix(rel, "..") {
		first := strings.Split(filepath.ToSlash(rel), "/")[0]
		for _, ex := range exs {
			if ex.Slug == first {
				return ex
			}
		}
	}
	for _, ex := range exs {
		if ex.ID == lastStarted {
			return ex
		}
	}
	return nil
}
