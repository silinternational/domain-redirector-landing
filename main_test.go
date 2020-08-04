package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func clearEnv(t *testing.T) {
	vars := []string{"NEW_HOST", "PORT", "MORE_INFO_URL", "REDIRECT_END_DATE", "ADDITIONAL_MESSAGE"}
	for _, k := range vars {
		if err := os.Unsetenv(k); err != nil {
			t.Errorf("failed to unset env var %s: %s", k, err.Error())
		}
	}
}

func loadEnv(t *testing.T, vars map[string]string) {
	for k, v := range vars {
		if err := os.Setenv(k, v); err != nil {
			t.Errorf("failed to set env var %s: %s", k, err.Error())
		}
	}
}

func Test_serveTemplate(t *testing.T) {
	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}
	tests := []struct {
		name          string
		args          args
		envConfig     map[string]string
		wantStatus    int
		wantInBody    []string
		notWantInBody []string
	}{
		{
			name: "only set NEW_HOST",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodGet, "https://old.url.com/path/to/page", nil),
			},
			envConfig:  map[string]string{"NEW_HOST": "new.url.com"},
			wantStatus: 200,
			wantInBody: []string{
				"<em>old.url.com</em> has changed to <em>new.url.com</em>",
				"https://new.url.com/path/to/page",
			},
			notWantInBody: []string{
				"https://old.url.com/path/to/page",
				"This message will be available until",
				"Learn more about this change at",
			},
		},
		{
			name: "use additionalMessage",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodGet, "https://old.url.com/path/to/page", nil),
			},
			envConfig: map[string]string{
				"NEW_HOST":           "new.url.com",
				"ADDITIONAL_MESSAGE": "y0! we've moved!",
			},
			wantStatus: 200,
			wantInBody: []string{
				"y0! we've moved!",
				"https://new.url.com/path/to/page",
			},
			notWantInBody: []string{
				"<em>old.url.com</em> has changed to <em>new.url.com</em>",
				"https://old.url.com/path/to/page",
				"This message will be available until",
				"Learn more about this change at",
			},
		},
		{
			name: "use additionalMessage",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodGet, "https://old.url.com/path/to/page", nil),
			},
			envConfig: map[string]string{
				"NEW_HOST":           "new.url.com",
				"ADDITIONAL_MESSAGE": "y0! we've moved!",
			},
			wantStatus: 200,
			wantInBody: []string{
				"y0! we've moved!",
				"https://new.url.com/path/to/page",
			},
			notWantInBody: []string{
				"<em>old.url.com</em> has changed to <em>new.url.com</em>",
				"https://old.url.com/path/to/page",
				"This message will be available until",
				"Learn more about this change at",
			},
		},
		{
			name: "use all optional vars",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodGet, "https://old.url.com/path/to/page?param1=something&param2=something%20else", nil),
			},
			envConfig: map[string]string{
				"NEW_HOST":           "new.url.com",
				"PORT":               "8080",
				"MORE_INFO_URL":      "https://moreinforurl.com",
				"REDIRECT_END_DATE":  "22 Aug 2020",
				"ADDITIONAL_MESSAGE": "Hi there, we've changed old.url.com to new.url.com, get with the program!",
			},
			wantStatus: 200,
			wantInBody: []string{
				"https://moreinforurl.com",
				"https://new.url.com/path/to/page?param1=something&amp;param2=something&#43;else",
				"22 Aug 2020",
				"Hi there, we've changed old.url.com to new.url.com, get with the program!",
			},
			notWantInBody: []string{
				"<em>old.url.com</em> has changed to <em>new.url.com</em>",
				"https://old.url.com/path/to/page?param1=something&param2=something%20else",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clearEnv(t)
			loadEnv(t, tt.envConfig)
			loadConfig()

			serveTemplate(tt.args.w, tt.args.r)
			resp := tt.args.w.Result()
			if resp.StatusCode != tt.wantStatus {
				t.Errorf("status code not as expected. wanted %v, got %v", tt.wantStatus, resp.StatusCode)
			}

			if tt.wantStatus != 200 {
				return
			}

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Errorf("got error trying to ready body: %s", err.Error())
			}
			for _, s := range tt.wantInBody {
				if !strings.Contains(string(body), s) {
					t.Errorf("%s \ndoes not contain \n%s", body, s)
				}
			}
			for _, s := range tt.notWantInBody {
				if strings.Contains(string(body), s) {
					t.Errorf("%s \n contains \n%s", body, s)
				}
			}
		})
	}
}
