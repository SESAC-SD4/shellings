package exercises

import "shellings/internal/exercise"

func init() {
	register(14, "14-ls-to-file", "14 목록을 파일로", false, setupLsToFile, checkLsToFile)
}

var lsToFileWant = []string{"10-12.md", "10-13.md", "10-14.md", "10-15.md"}

func setupLsToFile(dir string) error {
	return exercise.WriteFiles(dir, map[string]string{
		"TIL/10-12.md": til1012,
		"TIL/10-13.md": til1013,
		"TIL/10-14.md": til1014,
		"TIL/10-15.md": til1015,
		"portfolio/":   "",
	})
}

func checkLsToFile(dir string) exercise.Result {
	body, ok := exercise.Read(dir, "filelist.txt")
	if !ok {
		if exercise.IsFile(dir, "TIL/filelist.txt") {
			return exercise.Fail("`filelist.txt`가 `TIL` 안에 생겼습니다. 연습 폴더 맨 위에 있어야 합니다")
		}
		return exercise.Missing(dir, "filelist.txt")
	}
	got := names(body)
	for _, n := range lsToFileWant {
		if !got[n] {
			if got["TIL"] {
				return exercise.Fail("`filelist.txt`에 맨 위 폴더의 목록이 들어 있습니다. `TIL` 폴더 안의 목록이어야 합니다")
			}
			return exercise.Fail("목록에 빠진 파일이 있습니다: `" + n + "`")
		}
	}
	return exercise.Pass("`TIL` 안의 파일 목록을 파일로 남겼습니다")
}
