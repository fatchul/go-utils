package mail

import (
	"testing"
)

func TestSendingEmail_withRecipient(t *testing.T) {
	to, subject, body := "fatchulamin3@gmail.com", "Test Mail", "./sample.html"
	templateData := make(map[string]interface{})
	templateData["title"] = "Hello"
	templateData["sender"] = "Fatchul Amin"

	t.Logf("Start send email to %s", to)

	err := SendingEmail(to, subject, body, templateData)
	if err != nil {
		t.Errorf("Failed send email got %v", err)
		t.FailNow()
	}

	t.Log("Email sent to: ", to)
}
