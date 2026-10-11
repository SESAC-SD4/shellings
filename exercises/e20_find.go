package exercises

import (
	"path"
	"strings"

	"shellings-ko/internal/exercise"
)

func init() { register(20, "20-find", "20★ 파일 찾기", true, setupFind, checkFind) }

var findWant = []string{"day1.md", "day2.md", "day3.md", "day4.md", "day5.md"}

func setupFind(dir string) error {
	return exercise.WriteFiles(dir, map[string]string{
		"notes/day1.md":         til1012,
		"notes/2026/day2.md":    til1013,
		"notes/2026/10/day3.md": til1014,
		"archive/day4.md":       til1015,
		"archive/old/day5.md":   til1016,
		"notes/todo.txt":        "할 일\n",
		"archive/old/photo.txt": "사진 설명\n",
		"readme.txt":            "읽어 보세요\n",
	})
}

func checkFind(dir string) exercise.Result {
	body, ok := exercise.Read(dir, "filelist.txt")
	if !ok {
		return exercise.Missing(dir, "filelist.txt")
	}
	seen := map[string]bool{}
	for _, line := range exercise.NonEmptyLines(body) {
		name := path.Base(strings.ReplaceAll(line, `\`, "/"))
		if !strings.HasSuffix(name, ".md") {
			return exercise.Fail("`.md`가 아닌 파일이 섞여 있습니다: `" + name + "`")
		}
		seen[name] = true
	}
	for _, n := range findWant {
		if !seen[n] {
			return exercise.Fail("목록에 빠진 파일이 있습니다: `" + n + "`. 안쪽 폴더까지 모두 찾아야 합니다")
		}
	}
	return exercise.Pass("`find`로 숨어 있던 `.md` 파일을 모두 찾았습니다")
}
