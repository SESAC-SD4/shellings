package exercises

import (
	"fmt"
	"strings"

	"shellings-ko/internal/exercise"
)

func init() {
	register(19, "19-wildcard", "19★ 한꺼번에 정리", true, setupWildcard, checkWildcard)
}

const (
	wildcardPNG = 8
	wildcardMD  = 6
)

func setupWildcard(dir string) error {
	files := map[string]string{"TIL/10-12.md": til1012, "images/": ""}
	for i := 1; i <= wildcardPNG; i++ {
		files[fmt.Sprintf("shot%d.png", i)] = "PNG\n"
	}
	for i := 0; i < wildcardMD; i++ {
		files[fmt.Sprintf("10-%d.md", 13+i)] = til1013
	}
	return exercise.WriteFiles(dir, files)
}

func countSuffix(names []string, suffix string) int {
	n := 0
	for _, name := range names {
		if strings.HasSuffix(name, suffix) {
			n++
		}
	}
	return n
}

func checkWildcard(dir string) exercise.Result {
	top := exercise.List(dir, "")
	if left := countSuffix(top, ".png") + countSuffix(top, ".md"); left > 0 {
		return exercise.Fail(fmt.Sprintf("아직 맨 위에 남아 있습니다: 파일 %d개", left))
	}
	images, til := exercise.List(dir, "images"), exercise.List(dir, "TIL")
	if countSuffix(til, ".png") > 0 {
		return exercise.Fail("`.png` 파일이 `TIL`에 들어갔습니다. `.png`는 `images`로 가야 합니다")
	}
	if countSuffix(images, ".md") > 0 {
		return exercise.Fail("`.md` 파일이 `images`에 들어갔습니다. `.md`는 `TIL`로 가야 합니다")
	}
	if countSuffix(images, ".png") != wildcardPNG {
		return exercise.Fail(fmt.Sprintf("`images`에 `.png`가 %d개 있어야 하는데 %d개입니다", wildcardPNG, countSuffix(images, ".png")))
	}
	if countSuffix(til, ".md") != wildcardMD+1 {
		return exercise.Fail(fmt.Sprintf("`TIL`에 `.md`가 %d개 있어야 하는데 %d개입니다", wildcardMD+1, countSuffix(til, ".md")))
	}
	return exercise.Pass("`*`로 파일 14개를 한꺼번에 정리했습니다")
}
