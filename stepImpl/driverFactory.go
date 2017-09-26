package stepImpl

import (
	"fmt"

	"github.com/getgauge-contrib/gauge-go/gauge"
	. "github.com/getgauge-contrib/gauge-go/testsuit"
	"github.com/sclevine/agouti"
)

var driver *agouti.WebDriver
var page *agouti.Page

var setup = func() {
	driver = agouti.ChromeDriver(agouti.Desired(agouti.Capabilities{
		"chromeOptions": map[string][]string{
			"args": []string{"headless", "disable-gpu", "no-sandbox"},
		},
	}))
	err := driver.Start()
	if err != nil {
		T.Fail(fmt.Errorf("failed to open session: %s", err.Error()))
	}
	page, err = driver.NewPage()
	if err != nil {
		T.Fail(fmt.Errorf("failed to open page: %s", err.Error()))
	}

}

var quit = func() {
	if err := page.Destroy(); err != nil {
		T.Fail(fmt.Errorf("failed to the page: %s", err.Error()))
	}
	if err := driver.Stop(); err != nil {
		T.Fail(fmt.Errorf("failed to stop session: %s", err.Error()))
	}
}

var _ = gauge.BeforeSuite(setup, []string{}, AND)
var _ = gauge.AfterSuite(quit, []string{}, AND)
