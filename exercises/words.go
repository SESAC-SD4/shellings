package exercises

import (
	"path"
	"strings"
)

// names는 `ls`·`find` 출력처럼 공백·줄로 나뉜 파일 경로에서 마지막 이름만 모은다.
func names(body string) map[string]bool {
	out := map[string]bool{}
	for _, f := range strings.Fields(body) {
		out[path.Base(strings.ReplaceAll(f, `\`, "/"))] = true
	}
	return out
}
