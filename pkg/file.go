package pkg

import (
	"os"
	"path/filepath"
)

func MoveDir(srcDir, destDir string) error {
	// 移动
	destFile := getDest(srcDir, destDir)
	_, err := os.Stat(destFile)
	if os.IsNotExist(err) {
		os.MkdirAll(destFile, os.ModePerm)
	} else if err != nil {
		return err
	}
	err = os.Rename(srcDir, filepath.Join(destFile, filepath.Base(srcDir)))
	if err != nil {
		return err
	}

	// 删除源目录
	err = os.RemoveAll(srcDir)
	if err != nil {
		return err
	}

	return nil
}

// 获取目标文件夹路径
func getDest(path, dest string) string {
	rel, _ := filepath.Rel(filepath.Dir(dest), filepath.Dir(path))
	return filepath.Join(dest, rel)
}
