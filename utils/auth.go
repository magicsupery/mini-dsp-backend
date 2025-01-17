package utils

import (
	"crypto/sha256"
	"fmt"
)

// HashPassword 示例 - 使用 SHA256 （仅演示，生产环境应加入盐值并使用更安全的哈希方式）
func HashPassword(raw string) string {
	h := sha256.New()
	h.Write([]byte(raw))
	return fmt.Sprintf("%x", h.Sum(nil))
}
