package exercises

import (
	"io/fs"
	"path/filepath"

	"shellings-ko/internal/exercise"
)

func init() { register(10, "10-ls-detail", "10 폴더 구별하기", false, setupLsDetail, checkLsDetail) }

const mysteryTarget = "mystery/notes2/found.txt"

func setupLsDetail(dir string) error {
	return exercise.WriteFiles(dir, map[string]string{
		"mystery/notes":     "메모\n",
		"mystery/notes.bak": "백업\n",
		"mystery/old":       "옛날 파일\n",
		"mystery/notes2/":   "",
	})
}

func checkLsDetail(dir string) exercise.Result {
	if exercise.IsDir(dir, "mystery/notes2") && exercise.Exists(dir, mysteryTarget) {
		return exercise.Pass("`ls -l`로 폴더를 가려내고 그 안에 파일을 만들었습니다")
	}
	var elsewhere string
	_ = filepath.WalkDir(dir, func(p string, d fs.DirEntry, err error) error {
		if err == nil && !d.IsDir() && d.Name() == "found.txt" && elsewhere == "" {
			rel, _ := filepath.Rel(dir, p)
			elsewhere = filepath.ToSlash(rel)
		}
		return nil
	})
	if elsewhere != "" {
		return exercise.Fail("`found.txt`가 다른 곳에 생겼습니다: `" + elsewhere + "`. 원래 있던 폴더는 하나뿐입니다. `ls -l`의 맨 앞 글자를 보세요")
	}
	return exercise.Fail("아직 `mystery` 안의 폴더 속에 `found.txt`가 없습니다")
}
