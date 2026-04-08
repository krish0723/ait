package profile

import (
	"crypto/sha256"
	"fmt"
	"io/fs"
	"sort"
)

// BundleDigest returns a stable sha256 over all embedded profile/preset bytes (paths + contents, sorted).
func BundleDigest() string {
	var paths []string
	_ = fs.WalkDir(bundle, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		paths = append(paths, path)
		return nil
	})
	sort.Strings(paths)

	h := sha256.New()
	for _, p := range paths {
		b, err := bundle.ReadFile(p)
		if err != nil {
			continue
		}
		_, _ = h.Write([]byte(p))
		_, _ = h.Write([]byte{0})
		_, _ = h.Write(b)
		_, _ = h.Write([]byte{0})
	}
	return fmt.Sprintf("sha256:%x", h.Sum(nil))
}
