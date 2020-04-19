package libyiff

import "testing"

var testOptions Options = Options{"", ""}

func TestSum(t *testing.T) {
	var md5HashTest = md5Hash("/Users/puggo/Downloads/gitlab-icon-rgb.png")
	if md5HashTest != "d13bef02a7454ec7fc1a406ed303e8d7" {
		t.Error("Sum was incorrect, got:" + md5HashTest + ", wanted: " + "d13bef02a7454ec7fc1a406ed303e8d7")
	}
}
