/*


Copyright 2020 Red Hat, Inc.

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as
published by the Free Software Foundation, either version 3 of the
License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package sources

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"testing"

	logr "github.com/go-logr/logr/testing"
	"github.com/project-koku/korekuta-operator-go/crhchttp"
)

var (
	// cost       = &crhchttp.CostManagementConfig{Log: zap.New()}
	errSources = errors.New("test error")
	// log        logr.TestLogger
)

// https://www.thegreatcodeadventure.com/mocking-http-requests-in-golang/
type MockClient struct {
	res *http.Response
	err error
}

// Do is the mock client's `Do` func
func (m MockClient) Do(req *http.Request) (*http.Response, error) {
	return m.res, m.err
}

func TestGetSourceTypeID(t *testing.T) {
	log := logr.TestLogger{T: t}
	cost := &crhchttp.CostManagementConfig{Log: log}

	getSourceTypeIDTests := []struct {
		name        string
		response    *http.Response
		responseErr error
		expected    string
		expectedErr error
	}{
		{
			name: "successful response with data",
			response: &http.Response{
				StatusCode: 200,
				Body:       ioutil.NopCloser(strings.NewReader("{\"meta\":{\"count\":1},\"data\":[{\"id\":\"1\",\"name\":\"openshift\"}]}")), // type is io.ReadCloser,
				Request:    &http.Request{Method: "GET", URL: &url.URL{}},
			},
			responseErr: nil,
			expected:    "1",
			expectedErr: nil,
		},
		{
			name: "400 bad response - malformed url",
			response: &http.Response{
				StatusCode: 400,
				Body:       ioutil.NopCloser(strings.NewReader("{\"errors\":[{\"status\":\"400\",\"detail\":\"ArgumentError: Failed to find definition for Name\"}]}")), // type is io.ReadCloser,
				Request:    &http.Request{Method: "GET", URL: &url.URL{}},
			},
			responseErr: nil,
			expected:    "",
			expectedErr: errSources,
		},
	}
	for _, tt := range getSourceTypeIDTests {
		t.Run(tt.name, func(t *testing.T) {
			clt := MockClient{res: tt.response, err: tt.responseErr}
			got, err := GetSourceTypeID(cost, clt)
			if tt.expectedErr != nil && err == nil {
				t.Errorf("%s expected error, got: %v", tt.name, err)
			}
			if tt.expectedErr == nil && err != nil {
				t.Errorf("%s got unexpected error: %v", tt.name, err)
			}
			if got != tt.expected {
				t.Errorf("%s got %s want %s", tt.name, got, tt.expected)
			}
		})
	}
}
