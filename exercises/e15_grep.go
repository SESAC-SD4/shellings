package exercises

import (
	"fmt"

	"shellings-ko/internal/exercise"
)

func init() { register(15, "15-grep", "15 단어로 찾기", false, setupGrep, checkGrep) }

var grepWant = []struct{ name, body string }{
	{"10-14.md", grep1014},
	{"10-16.md", grep1016},
}

func setupGrep(dir string) error {
	return exercise.WriteFiles(dir, map[string]string{
		"TIL/10-12.md": til1012,
		"TIL/10-13.md": til1013,
		"TIL/10-14.md": grep1014,
		"TIL/10-15.md": til1015,
		"TIL/10-16.md": grep1016,
		"TIL/10-17.md": til1017,
		"portfolio/":   "",
	})
}

func checkGrep(dir string) exercise.Result {
	wanted := map[string]bool{}
	for _, w := range grepWant {
		wanted[w.name] = true
	}
	for _, name := range exercise.List(dir, "portfolio") {
		if !wanted[name] {
			return exercise.Fail("`포트폴리오`라는 단어가 없는 글이 `portfolio`에 들어 있습니다: `" + name + "`")
		}
	}
	count := 0
	for _, w := range grepWant {
		if exercise.IsFile(dir, "portfolio/"+w.name) {
			count++
		}
	}
	if count < len(grepWant) {
		return exercise.Fail(fmt.Sprintf("`포트폴리오`가 들어 있는 글은 %d개입니다. 지금 `portfolio`에는 %d개뿐입니다", len(grepWant), count))
	}
	for _, w := range grepWant {
		if !exercise.IsFile(dir, "TIL/"+w.name) {
			return exercise.Fail("원본이 `TIL`에서 사라졌습니다: `" + w.name + "`. 옮기기(`mv`)가 아니라 복사(`cp`)하세요")
		}
		if body, _ := exercise.Read(dir, "portfolio/"+w.name); !exercise.Same(body, w.body) {
			return exercise.Fail("`portfolio/" + w.name + "`의 내용이 원본과 다릅니다")
		}
	}
	return exercise.Pass("`grep`으로 찾은 글을 포트폴리오에 복사했습니다")
}
