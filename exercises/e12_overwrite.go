package exercises

import "shellings-ko/internal/exercise"

func init() { register(12, "12-overwrite", "12 되돌릴 수 없는 사고", false, setupOverwrite, checkOverwrite) }

func setupOverwrite(dir string) error {
	return exercise.WriteFiles(dir, map[string]string{
		"TIL/10-12.md":    accidentLine,
		"backup/10-12.md": til1012,
	})
}

func checkOverwrite(dir string) exercise.Result {
	body, ok := exercise.Read(dir, "TIL/10-12.md")
	if !ok {
		return exercise.Missing(dir, "TIL/10-12.md")
	}
	if exercise.Same(body, accidentLine) {
		return exercise.Fail("아직 사고가 난 상태 그대로입니다. `backup` 폴더를 열어 보세요")
	}
	if !exercise.Same(body, til1012) {
		return exercise.Fail("`TIL/10-12.md`의 내용이 백업과 다릅니다. 다시 쓰지 말고 백업을 복사해 오세요")
	}
	if !exercise.IsFile(dir, "backup/10-12.md") {
		return exercise.Pass("백업 파일을 옮겨 와서 `backup`에는 이제 없습니다. 복사(`cp`)했다면 백업도 남았을 겁니다.\n" + gitBridge)
	}
	return exercise.Pass(gitBridge)
}
