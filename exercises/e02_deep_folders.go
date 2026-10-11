package exercises

import "shellings-ko/internal/exercise"

func init() {
	register(2, "02-deep-folders", "02 깊은 폴더", false, setupDeepFolders, checkDeepFolders)
}

const deepTarget = "TIL/2026-10/week2/day1"

func setupDeepFolders(dir string) error {
	return exercise.WriteFiles(dir, map[string]string{"TIL/10-12.md": til1012})
}

func checkDeepFolders(dir string) exercise.Result {
	if exercise.IsDir(dir, deepTarget) {
		return exercise.Pass("중간 폴더까지 한 번에 만들었습니다")
	}
	if exercise.IsFile(dir, deepTarget) {
		return exercise.Fail("`day1`이 폴더가 아니라 파일입니다. 지우고 폴더로 다시 만드세요")
	}
	if exercise.IsDir(dir, "2026-10/week2/day1") {
		return exercise.Fail("`2026-10/week2/day1`이 `TIL` 폴더 밖(연습 폴더 맨 위)에 생겼습니다. `TIL` 안에 만들어야 합니다")
	}
	return exercise.Missing(dir, deepTarget)
}
