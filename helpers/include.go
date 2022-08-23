package helpers

import "path/filepath"

func Include(path string) ([]string, error) {
	files, err := filepath.Glob("site/views/templates/*.html")
	if err != nil {
		return nil, err
	}
	path_files, err := filepath.Glob("site/views/" + path + "/*.html")
	if err != nil {
		return nil, err
	}
	for _, file := range path_files {
		files = append(files, file)
	}
	return files, nil
}
