package exercises

import (
	"strings"

	"shellings/internal/exercise"
)

func init() {
	register(22, "22-diff", "22★ 바뀐 줄 찾기", true, setupDiff, checkDiff)
}

const (
	diffSame1 = "## 마크다운"
	diffSame2 = "- 제목은 #으로 쓴다"
	diffSame3 = "- 굵게는 별표 두 개로 감싼다"
	diffOld   = "- 목록은 -로 쓴다"
	diffNew1  = "- 목록은 -나 *로 쓴다"
	diffNew2  = "- 링크는 [글자](주소)로 쓴다"
)

func setupDiff(dir string) error {
	return exercise.WriteFiles(dir, map[string]string{
		"backup/10-13.md": strings.Join([]string{diffSame1, diffSame2, diffOld, diffSame3}, "\n") + "\n",
		"TIL/10-13.md":    strings.Join([]string{diffSame1, diffSame2, diffNew1, diffSame3, diffNew2}, "\n") + "\n",
	})
}

func checkDiff(dir string) exercise.Result {
	body, ok := exercise.Read(dir, "changed.txt")
	if !ok {
		return exercise.Missing(dir, "changed.txt")
	}
	for _, same := range []string{diffSame1, diffSame2, diffSame3} {
		if strings.Contains(body, same) {
			return exercise.Fail("바뀌지 않은 줄까지 들어 있습니다: `" + same + "`. 두 파일에서 다른 줄만 남겨야 합니다")
		}
	}
	if !strings.Contains(body, diffNew1) || !strings.Contains(body, diffNew2) {
		if strings.Contains(body, diffOld) {
			return exercise.Fail("옛 파일(`backup`)에만 있는 줄이 들어 있습니다. 새 파일(`TIL`)에서 바뀌거나 새로 생긴 줄을 적어야 합니다")
		}
		return exercise.Fail("오늘 바뀌거나 새로 생긴 두 줄이 모두 들어 있지 않습니다")
	}
	return exercise.Pass("`diff`로 바뀐 줄을 찾았습니다. 내일 `git diff`도 이렇게 읽습니다")
}
