package config

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParseConfigFile(t *testing.T) {
	testCases := []struct {
		name     string
		filename string
		want     Config
		wantErr  bool
		errMsg   string
	}{
		{
			name:     "success",
			filename: "../../testdata/pkg/schema/valid.yml",
			want: Config{
				TargetServer: TargetServer{
					Host: "localhost",
					Port: 8080,
				},
				Services: []Service{
					{
						Name: "users",
						Methods: []Method{
							{
								Name: "createUser",
								Input: map[string]interface{}{
									"name":  "John Doe",
									"age":   30,
									"email": "XXXXXXXXXXXX",
									"phone": "XXXXXXXXXXXX",
									"address": map[string]interface{}{
										"street":  "123 Main St",
										"city":    "New York",
										"state":   "CA",
										"zip":     "94105",
										"country": "USA",
									},
									"password": "XXXXXXXXXXXX",
									"role":     "user",
									"status":   "active",
								},
							},
							{
								Name: "getUser",
								Input: map[string]interface{}{
									"id": "XXXXXXXXXXXX",
								},
							},
						},
					},
				},
				LoadPattern: LoadPattern{
					Type:            "ramp-up",
					ConcurrentUsers: 10,
					Duration:        10 * time.Second,
					RampUp: RampUp{
						Duration: time.Second * 10,
					},
					Cooldown: Cooldown{
						Duration: time.Second * 10,
					},
				},
				RateLimiting: RateLimiting{
					MaxRequestsPerSecond: 10,
				},
				Metadata: map[string]string{
					"name":        "test",
					"environment": "test",
					"version":     "1.0.0",
				},
				TLS: &TLS{
					Enabled:  false,
					CertFile: "testdata/cert.pem",
					KeyFile:  "testdata/key.pem",
				},
			},
			wantErr: false,
			errMsg:  "",
		},
		{
			name:     "invalid file",
			filename: "../../testdata/pkg/schema/invalid.yml",
			wantErr:  true,
			errMsg:   "error while parsing config file",
		},
		{
			name:     "absent file",
			filename: "absent.yml",
			wantErr:  true,
			errMsg:   "failed to read config file",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := ParseConfigFile(tc.filename)

			if tc.wantErr {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.errMsg)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func BenchmarkParseConfigFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		filename := "../../testdata/pkg/schema/valid.yml"
		ParseConfigFile(filename)
	}
}
