package sqlimporter

import "testing"

func TestGetFileList(t *testing.T) {
	t.Parallel()
	fileList := []string{"test1.sql", "test2.sql"}
	list, err := getFileList("files")
	if err != nil {
		t.Error(err)
	}
	if len(list) != len(fileList) {
		t.Error("List of files is different")
	}
}

func TestParseFiles(t *testing.T) {
	t.Parallel()
	queries, err := parseFiles("files/test1.sql")
	if err != nil {
		t.Error(err)
	}
	if len(queries) != 2 {
		t.Logf("%+v", queries)
		t.Error("mismatch queries count")
	}
}
