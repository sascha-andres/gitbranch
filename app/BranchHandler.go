// Copyright © 2017 Sascha Andres <sascha.andres@outlook.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package app

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
)

type (
	// BranchResult represents a successful answer for a request
	BranchResult struct {
		Options []BranchInformation `json:"options"`
	}

	// BranchRequest represents the payload we ask for
	BranchRequest struct {
		Repository string `json:"repository"`
	}
)

var testRegex = regexp.MustCompile("[a-z@-_:/\\.]+git$")

// BranchHandler is the http.HandlerFunc for a branch list request
func BranchHandler(w http.ResponseWriter, r *http.Request) {
	var request BranchRequest
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &request); err != nil {
		handleError(w, err)
		return
	}

	if !testRegex.MatchString(request.Repository) {
		w.WriteHeader(500)
		return
	}

	branches, err := GetBranches(request.Repository)
	if err != nil {
		handleError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(BranchResult{Options: branches})
}

func handleError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(422) // unprocessable entity
	if err := json.NewEncoder(w).Encode(err); err != nil {
		panic(err)
	}
}
