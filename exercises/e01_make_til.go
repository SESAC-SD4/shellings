package exercises

import "shellings-ko/internal/exercise"

func init() { register(1, "01-make-til", "01 TIL 폴더", false, setupMakeTil, checkMakeTil) }

func setupMakeTil(dir string) error { return nil }

func checkMakeTil(dir string) exercise.Result {
	if exercise.IsFile(dir, "TIL") {
		return exercise.Fail("`TIL`이 폴더가 아니라 파일입니다. 그 파일을 지우고 폴더로 다시 만드세요")
	}
	if !exercise.IsDir(dir, "TIL") {
		return exercise.Missing(dir, "TIL")
	}
	return exercise.Pass("`TIL` 폴더가 생겼습니다")
}
