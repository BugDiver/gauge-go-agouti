package stepImpl

import (
	. "github.com/getgauge-contrib/gauge-go/gauge"
	m "github.com/getgauge-contrib/gauge-go/models"
)

const (
	userNameID  = "q_username"
	bunttonName = "commit"
)

var _ = Step("On the customer page", func() {
	err := page.Navigate("http://localhost:8080/admin/customers")
	assert(err, nil)
})

var _ = Step("Search for customer <name>", searchUser)
var _ = Step("The customer <name> is listed", verifyUser)

var _ = Step("Search for customers <customersTable>", func(customersTable *m.Table) {
	for _, row := range customersTable.Rows {
		customerName := row.Cells[0]
		searchUser(customerName)
		verifyUser(customerName)
	}
})

func searchUser(name string) {
	box := page.FindByID(userNameID)
	assert(box.Clear(), nil)
	assert(box.SendKeys(name), nil)
	assert(page.FindByName(bunttonName).Click(), nil)
}

func verifyUser(name string) {
	actualUserName, err := page.Find("table#index_table_customers tbody tr:nth-child(1) td.col-username").Text()
	assert(err, nil)
	assert(actualUserName, name)
}
