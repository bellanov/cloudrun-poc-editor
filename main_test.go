// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"os"
	"regexp"
	"testing"
)

func init() {
	os.Setenv("EDITOR_UPSTREAM_RENDER_URL", "http://testing.local")
}

func TestEditorHandler(t *testing.T) {

	data, _ := os.ReadFile("./templates/index.html")

	want := `<title>Markdown Editor</title>`
	re := regexp.MustCompile(`<title>.*</title>`)
	got := re.FindString(string(data))

	if got != want {
		t.Errorf("body: got %q, want %q", got, want)
	}

}
