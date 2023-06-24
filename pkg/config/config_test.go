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
			filename: "testdata/valid.yaml",
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
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := ParseConfigFile(tc.filename)

			if tc.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tc.errMsg, err.Error())
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}
