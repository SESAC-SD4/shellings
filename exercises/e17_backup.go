package exercises

import "shellings/internal/exercise"

func init() {
	register(17, "17-backup", "17 폴더째 백업", false, setupBackup, checkBackup)
}

var backupFiles = []struct{ rel, body string }{
	{"10-12.md", til1012},
	{"10-13.md", til1013},
	{"10-14.md", til1014},
	{"2026-10/week1/day1.md", til1015},
}

func setupBackup(dir string) error {
	files := map[string]string{}
	for _, f := range backupFiles {
		files["TIL/"+f.rel] = f.body
	}
	return exercise.WriteFiles(dir, files)
}

func checkBackup(dir string) exercise.Result {
	for _, f := range backupFiles {
		if !exercise.IsFile(dir, "TIL/"+f.rel) {
			return exercise.Fail("원본 `TIL/" + f.rel + "`이 사라졌습니다. 옮기기가 아니라 복사여야 합니다. 처음 상태로 되돌려 다시 해 보세요")
		}
	}
	if !exercise.IsDir(dir, "backup") {
		return exercise.Missing(dir, "backup")
	}
	if exercise.Exists(dir, "backup/TIL") {
		return exercise.Fail("`backup` 안에 `TIL` 폴더가 한 겹 더 생겼습니다. `backup`을 지우고 다시 한 번만 복사하세요")
	}
	for _, f := range backupFiles {
		body, ok := exercise.Read(dir, "backup/"+f.rel)
		if !ok {
			return exercise.Missing(dir, "backup/"+f.rel)
		}
		if !exercise.Same(body, f.body) {
			return exercise.Fail("`backup/" + f.rel + "`의 내용이 원본과 다릅니다")
		}
	}
	return exercise.Pass("`TIL`을 안의 폴더까지 통째로 백업했습니다")
}
