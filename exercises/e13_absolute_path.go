package exercises

import (
	"strings"

	"shellings-ko/internal/exercise"
)

func init() {
	register(13, "13-absolute-path", "13 어디서든 같은 곳", false, setupAbsolutePath, checkAbsolutePath)
}

func setupAbsolutePath(dir string) error {
	return exercise.WriteFiles(dir, map[string]string{
		"TIL/2026-10/week1/day1.md": til1012,
		"portfolio/":                "",
	})
}

// where.txt에는 학생이 `pwd`를 친 위치가 남으므로, "그 자리에서" 했는지 확인할 수 있다.
func checkAbsolutePath(dir string) exercise.Result {
	body, ok := exercise.Read(dir, "portfolio/where.txt")
	if !ok {
		if exercise.IsFile(dir, "TIL/2026-10/week1/where.txt") {
			return exercise.Fail("`where.txt`가 `TIL/2026-10/week1` 안에 생겼습니다. 지금 위치 기준이 아니라 `~`로 시작하는 경로로 `portfolio` 안에 저장해야 합니다")
		}
		return exercise.Missing(dir, "portfolio/where.txt")
	}
	where := strings.TrimRight(strings.ReplaceAll(strings.TrimSpace(body), `\`, "/"), "/")
	if where == "" {
		return exercise.Fail("`where.txt`는 생겼지만 비어 있습니다")
	}
	if !strings.HasSuffix(where, "/TIL/2026-10/week1") {
		return exercise.Fail("`where.txt`에 적힌 위치가 `week1`이 아닙니다: `" + where + "`. `week1` 안으로 들어간 뒤 그 자리에서 저장하세요")
	}
	return exercise.Pass("깊은 폴더 안에서도 `~`로 시작하는 경로로 원하는 곳에 저장했습니다")
}
