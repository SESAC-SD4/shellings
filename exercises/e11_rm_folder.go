package exercises

import "shellings/internal/exercise"

func init() { register(11, "11-rm-folder", "11 폴더 지우기", false, setupRmFolder, checkRmFolder) }

func setupRmFolder(dir string) error {
	return exercise.WriteFiles(dir, map[string]string{
		"TIL/10-12.md": til1012,
		"trash/a.txt":  "a\n",
		"trash/b.txt":  "b\n",
	})
}

func checkRmFolder(dir string) exercise.Result {
	if !exercise.IsFile(dir, "TIL/10-12.md") {
		return exercise.Fail("`TIL`에 있던 `10-12.md`가 사라졌습니다. `trash`만 지워야 합니다. 처음 상태로 되돌려 다시 해 보세요")
	}
	if exercise.Exists(dir, "trash") {
		if exercise.IsDir(dir, "trash") && len(exercise.List(dir, "trash")) == 0 {
			return exercise.Fail("`trash` 안은 비웠지만 폴더가 남아 있습니다. 폴더째 지우는 옵션을 찾아보세요")
		}
		return exercise.Fail("`trash` 폴더가 아직 있습니다")
	}
	return exercise.Pass("폴더를 안의 파일까지 통째로 지웠습니다")
}
