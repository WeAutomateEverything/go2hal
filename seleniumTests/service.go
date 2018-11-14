package seleniumTests

import (
	"errors"
	"fmt"
	"github.com/tebeka/selenium"
	"github.com/weAutomateEverything/go2hal/alert"
	"github.com/weAutomateEverything/go2hal/callout"
	"golang.org/x/net/context"
	"gopkg.in/kyokomi/emoji.v1"
	"runtime/debug"
	"time"
)

type Service interface {
}

type service struct {
	store          Store
	alert          alert.Service
	calloutService callout.Service
}

func NewService(store Store, alertService alert.Service, calloutService callout.Service) Service {
	s := &service{store, alertService, calloutService}
	go func() {
		s.runTests()
	}()
	return s
}

func NewChromeClient(seleniumEndpoint string) (selenium.WebDriver, error) {
	caps := selenium.Capabilities(map[string]interface{}{"browserName": "chrome"})
	caps["chrome.switches"] = []string{"--ignore-certificate-errors"}
	return selenium.NewRemote(caps, seleniumEndpoint)

}

func (s *service) testSelenium(item Selenium) (err error) {
	ctx := context.Background()
	_, err = s.doSelenium(ctx, item)
	if err != nil {
		return err
	}

	err = s.store.AddSelenium(item)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) runTests() {
	var ctx context.Context
	defer func() {
		if err := recover(); err != nil {
			fmt.Print(err)
			s.alert.SendError(ctx, errors.New(fmt.Sprint(err)))
			s.alert.SendError(ctx, errors.New(string(debug.Stack())))

		}
	}()

	for true {
		tests, err := s.store.GetAllSeleniumTests()
		if err != nil {
			s.alert.SendError(ctx, err)
		} else {
			for _, test := range tests {
				image, err := s.doSelenium(ctx, test)
				if err != nil {
					s.handleError(ctx, test, image, err)
				} else {
					s.handleSuccess(ctx, test)
				}

			}
		}
		time.Sleep(5 * time.Minute)
	}
}

func (s service) handleError(ctx context.Context, test Selenium, image []byte, err error) {
	if error := s.store.SetSeleniumFailing(&test, err); error != nil {
		s.alert.SendError(ctx, fmt.Errorf("error setting seleniumTests test to failed. %s", error.Error()))
		return
	}
	if test.Threshold > 0 {
		if test.Threshold == test.ErrorCount {
			for _, chat := range test.Chats {
				s.calloutService.InvokeCallout(ctx, chat, fmt.Sprintf("halSelenium Error with  test %s", test.Name), err.Error(), false)
			}

		}

		if test.ErrorCount >= test.Threshold {
			for _, chat := range test.Chats {
				s.alert.SendAlert(ctx, chat, emoji.Sprintf(":computer: :x: Error executing seleniumTests test for %s. error: %s", test.Name, err.Error()))
				if image != nil {
					s.alert.SendImageToAlertGroup(ctx, chat, image)
				}
			}
		}
	}
}

func (s service) handleSuccess(ctx context.Context, test Selenium) {
	if err := s.store.SetSeleniumPassing(&test); err != nil {
		s.alert.SendError(ctx, fmt.Errorf("error setting seleniumTests test to passed. %s", err.Error()))
		return
	}
	if !test.Passing && test.ErrorCount >= test.Threshold {
		for _, chat := range test.Chats {
			s.alert.SendAlert(ctx, chat, emoji.Sprintf(":computer: :white_check_mark: halSelenium Test %s back to normal", test.Name))
		}
	}
}

func (s service) doSelenium(ctx context.Context, item Selenium) ([]byte, error) {
	if item.SeleniumServer == "" {
		return nil, errors.New("no halSelenium Server set")
	}
	if item.Name == "" {
		return nil, errors.New("no script name set")
	}
	if item.InitialURL == "" {
		return nil, errors.New("no initial url set")
	}
	if len(item.Pages) == 0 {
		return nil, errors.New("no pages detected in script")
	}

	var webDriver selenium.WebDriver
	var err error
	caps := selenium.Capabilities(map[string]interface{}{"browserName": "chrome"})
	caps["chrome.switches"] = []string{"--ignore-certificate-errors"}

	if webDriver, err = selenium.NewRemote(caps, item.SeleniumServer); err != nil {
		fmt.Printf("Failed to open session: %s\n", err)
		return nil, err
	}

	defer webDriver.Quit()

	err = webDriver.Get(item.InitialURL)
	if err != nil {
		return s.handleSeleniumError(ctx, item.Name, "Initial Page", "Load Page", err, webDriver)
	}

	for _, page := range item.Pages {
		if len(page.Actions) == 0 {
			return nil, fmt.Errorf("no pages found in test %s for page %s", item.Name, page.Name)
		}
		if page.PreCheck != nil {
			if page.PreCheck.SearchPattern == "" {
				return nil, fmt.Errorf("no search pattern found for precheck on test %s, page %s, check %s", item.Name, page.Name, page.PreCheck.Name)
			}
			err = doCheck(page.PreCheck, webDriver)
			if err != nil {
				return s.handleSeleniumError(ctx, item.Name, page.Name, page.PreCheck.Name, err, webDriver)
			}
		}
		for _, action := range page.Actions {
			if action.SearchPattern == "" {
				return nil, fmt.Errorf("no search pattern found for test %s, page %s, action %s", item.Name, page.Name, action.Name)
			}
			elems, err := findElement(action.SearchOption, webDriver)
			if err != nil {
				return s.handleSeleniumError(ctx, item.Name, page.Name, action.Name, err, webDriver)
			}
			elem := elems[0]
			executed := false
			if action.ClickLink != nil {
				elem.Click()
				executed = true
			}
			if action.ClickButton != nil && !executed {
				elem.Click()
				executed = true
			}
			if action.InputData != nil && !executed {
				elem.SendKeys(action.InputData.Value)
				executed = true
			}
			if !executed {
				return nil, fmt.Errorf("no action executed for test %s, page %s, action %s", item.Name, page.Name, action.Name)
			}

		}
		if page.PostCheck != nil {
			if page.PostCheck.SearchPattern == "" {
				return nil, fmt.Errorf("no search pattern found for post check on test %s, page %s, check %s", item.Name, page.Name, page.PostCheck.Name)
			}
			err := doCheck(page.PostCheck, webDriver)
			if err != nil {
				return s.handleSeleniumError(ctx, item.Name, page.Name, page.PostCheck.Name, err, webDriver)
			}
		}
	}
	return nil, nil
}

func doCheck(check *Check, driver selenium.WebDriver) error {
	waitfor := func(wb selenium.WebDriver) (bool, error) {

		elems, err := findElement(check.SearchOption, driver)
		if err != nil {
			return false, nil
		}

		for _, elem := range elems {
			dis, err := elem.IsDisplayed()
			if err != nil {
				return false, nil
			}
			if dis {
				if check.Value != nil {
					s, err := elem.Text()
					if err != nil {
						return false, nil
					}
					return *check.Value == s, nil
				}
				return true, nil
			}
		}
		return false, err
	}
	return driver.WaitWithTimeout(waitfor, 10*time.Second)

}

func (s service) handleSeleniumError(ctx context.Context, name, page, action string, err error, driver selenium.WebDriver) ([]byte, error) {
	bytes, error := driver.Screenshot()
	if error != nil {
		s.alert.SendError(ctx, error)
		return nil, err
	}
	return bytes, fmt.Errorf("application: %s,page: %s, action %s, Error: %s", name, page, action, err.Error())
}

func findElement(action SearchOption, driver selenium.WebDriver) ([]selenium.WebElement, error) {
	selector := selenium.ByCSSSelector
	if action.XPathSelector != nil {
		selector = selenium.ByXPATH
	}
	if action.PartialLinkTextSelect != nil {
		selector = selenium.ByPartialLinkText
	}
	if action.LinkTextSelector != nil {
		selector = selenium.ByLinkText
	}
	if action.IDSelector != nil {
		selector = selenium.ByID
	}
	if action.ClassNameSelector != nil {
		selector = selenium.ByCSSSelector
	}
	if action.NameSelector != nil {
		selector = selenium.ByName
	}
	if action.TagNameSelector != nil {
		selector = selenium.ByTagName
	}
	if action.Multiple {
		return driver.FindElements(selector, action.SearchPattern)
	}
	elem, err := driver.FindElement(selector, action.SearchPattern)
	return []selenium.WebElement{elem}, err

}
