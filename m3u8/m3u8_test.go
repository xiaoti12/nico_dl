package m3u8

import (
	"testing"
)

func TestGenerateM3U8URL(t *testing.T) {
	Client.SetProxy("http://localhost:7890")
	Cookies = "" // hide for commit
	suffix := "so43332851"
	key := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiJ9.eyJqdGkiOiI2NWJkMGFmY2FkY2U1IiwiZXhwIjoxNzA2ODg4NTMyLCJ0eXAiOiJBY2Nlc3MtUmlnaHQtS2V5IiwidmlkIjoic280MzMzMjg1MSIsInJpZCI6Im5pY292aWRlby1zbzQzMzMyODUxIiwiZmlkIjo2LCJ1aWQiOiIxMzAwNTI2NTciLCJkIjoyNzA0LCJ2IjpbInZpZGVvLWgyNjQtMTA4MHAiLCJ2aWRlby1oMjY0LTcyMHAiLCJ2aWRlby1oMjY0LTQ4MHAiLCJ2aWRlby1oMjY0LTM2MHAiLCJ2aWRlby1oMjY0LTE0NHAiXSwiYSI6WyJhdWRpby1hYWMtMTkya2JwcyIsImF1ZGlvLWFhYy02NGticHMiXSwicyI6ZmFsc2UsInNoIjpmYWxzZX0.PO_aUIC_KnaIIO_wctZQ4vMQcdObL-baTdUbYYfegzBHKplP40eaKGG1R3Kd91ckRpC1QxpSY2ZiQHFvRlAOMw"
	trackID := "d811ywn71i_1706887932519"
	url := generateM3U8URL(suffix, key, trackID, 128)
	if url == "" {
		t.Fatalf("Error generating m3u8 url")
	}
	t.Log(url)
}
