package rest

import "testing"

func TestGuildAuditLog(t *testing.T) {
	_, err := testingRestClient.GetGuildAuditLog(testingGuildID)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log("Successfully got the guild audit log")
}
