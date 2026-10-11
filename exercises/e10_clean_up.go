package exercises

import "shellings-ko/internal/exercise"

func init() {
	register(10, "10-clean-up", "10 임시 파일 지우기", false, setupCleanUp, checkCleanUp)
}

func setupCleanUp(dir string) error {
	return exercise.WriteFiles(dir, map[string]string{
		"TIL/10-12.md":     til1012,
		"TIL/temp-note.md": "임시 메모\n",
		"TIL/draft-old.md": "옛 초안\n",
	})
}

func checkCleanUp(dir string) exercise.Result {
	if !exercise.IsFile(dir, "TIL/10-12.md") {
		return exercise.Fail("지우면 안 되는 파일 `TIL/10-12.md`를 지웠습니다. 터미널에는 휴지통이 없습니다. 처음 상태로 되돌려 다시 해 보세요")
	}
	for _, name := range []string{"temp-note.md", "draft-old.md"} {
		if exercise.Exists(dir, "TIL/"+name) {
			return exercise.Fail("아직 남아 있습니다: `TIL/" + name + "`")
		}
	}
	return exercise.Pass("임시 파일만 골라서 지웠습니다")
}
