package exercises

import (
	"fmt"

	"shellings-ko/internal/exercise"
)

func init() {
	register(9, "09-organize", "09 흩어진 파일 정리", false, setupOrganize, checkOrganize)
}

type move struct{ name, dest, other, body string }

var organizeTargets = []move{
	{"10-15.md", "TIL", "portfolio", til1015},
	{"10-16.md", "TIL", "portfolio", til1016},
	{"about.md", "portfolio", "TIL", aboutMD},
}

func setupOrganize(dir string) error {
	return exercise.WriteFiles(dir, map[string]string{
		"TIL/10-12.md": til1012,
		"TIL/10-13.md": til1013,
		"TIL/10-14.md": til1014,
		"portfolio/":   "",
		"10-15.md":     til1015,
		"10-16.md":     til1016,
		"about.md":     aboutMD,
	})
}

func checkOrganize(dir string) exercise.Result {
	for _, m := range organizeTargets {
		if !exercise.Exists(dir, m.name) {
			continue
		}
		if exercise.Exists(dir, m.dest+"/"+m.name) {
			return exercise.Fail(fmt.Sprintf("`%s`가 `%s`에도 있고 맨 위 폴더에도 남아 있습니다. 복사(`cp`)가 아니라 옮기기(`mv`)를 쓰세요", m.name, m.dest))
		}
		return exercise.Fail(fmt.Sprintf("아직 맨 위 폴더에 있습니다: `%s`. `%s` 폴더로 옮기세요", m.name, m.dest))
	}
	for _, m := range organizeTargets {
		want := m.dest + "/" + m.name
		body, ok := exercise.Read(dir, want)
		if !ok {
			if exercise.Exists(dir, m.other+"/"+m.name) {
				return exercise.Fail(fmt.Sprintf("`%s`가 `%s`에 들어갔습니다. `%s`로 옮겨야 합니다", m.name, m.other, m.dest))
			}
			return exercise.Missing(dir, want)
		}
		if !exercise.Same(body, m.body) {
			return exercise.Fail(fmt.Sprintf("`%s`의 내용이 원래와 다릅니다. 처음 상태로 되돌려 다시 해 보세요", want))
		}
	}
	return exercise.Pass("세 파일이 모두 제자리에 들어갔습니다")
}
