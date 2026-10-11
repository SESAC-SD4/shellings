package exercises

import "shellings/internal/exercise"

func init() { register(3, "03-first-entry", "03 첫 기록", false, setupFirstEntry, checkFirstEntry) }

func setupFirstEntry(dir string) error {
	return exercise.WriteFiles(dir, map[string]string{"TIL/": ""})
}

func checkFirstEntry(dir string) exercise.Result {
	body, ok := exercise.Read(dir, "TIL/10-12.md")
	if !ok {
		if exercise.IsFile(dir, "10-12.md") {
			return exercise.Fail("`10-12.md`가 `TIL` 폴더 밖(연습 폴더 맨 위)에 생겼습니다. `TIL` 안에 만들어야 합니다")
		}
		return exercise.Missing(dir, "TIL/10-12.md")
	}
	if len(exercise.NonEmptyLines(body)) == 0 {
		return exercise.Fail("`TIL/10-12.md`는 생겼지만 비어 있습니다. 오늘 배운 것을 한 줄 이상 써 넣으세요")
	}
	return exercise.Pass("첫 기록을 남겼습니다")
}
