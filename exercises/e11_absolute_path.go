package exercises

import "shellings-ko/internal/exercise"

func init() {
	register(11, "11-absolute-path", "11 어디서든 같은 곳", false, setupAbsolutePath, checkAbsolutePath)
}

func setupAbsolutePath(dir string) error {
	return exercise.WriteFiles(dir, map[string]string{
		"TIL/2026-10/week1/day1.md": til1012,
		"portfolio/":                "",
	})
}

func checkAbsolutePath(dir string) exercise.Result {
	body, ok := exercise.Read(dir, "portfolio/README.md")
	if !ok {
		if exercise.IsFile(dir, "TIL/2026-10/week1/README.md") {
			return exercise.Fail("`README.md`가 `TIL/2026-10/week1` 안에 생겼습니다. 지금 위치 기준이 아니라 `~`로 시작하는 경로를 써 보세요")
		}
		return exercise.Missing(dir, "portfolio/README.md")
	}
	if len(exercise.NonEmptyLines(body)) == 0 {
		return exercise.Fail("`portfolio/README.md`는 생겼지만 비어 있습니다. 한 줄 이상 써 넣으세요")
	}
	return exercise.Pass("깊은 폴더에서도 `~`로 시작하는 경로로 원하는 곳에 썼습니다")
}
