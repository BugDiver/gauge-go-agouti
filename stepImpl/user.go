package stepImpl

import (
	. "github.com/getgauge-contrib/gauge-go/gauge"
)

var _ = Step("On signup page", func() {
	signUpURL := "http://localhost:8080/signup"
	err := page.Navigate(signUpURL)
	assert(err, nil)
})

var _ = Step("Fill <name>, <email> and <password>", func(name, email, password string) {
	page.FindByID("user_username").SendKeys(name)
	page.FindByID("user_email").SendKeys(email)
	page.FindByID("user_password").SendKeys(password)
	page.FindByID("user_password_confirmation").SendKeys(password)

})

var _ = Step("Submit registration form", func() {
	page.FindByName("commit").Click()
})
