package main

import (
	"fmt"
	"github.com/tebeka/selenium"
	"os"
	"time"
)

func main(){
	Example()
}
// This example shows how to navigate to a http://play.golang.org page, input a
// short program, run it, and inspect its output.
//
// If you want to actually run this example:
//
//   1. Ensure the file paths at the top of the function are correct.
//   2. Remove the word "Example" from the comment at the bottom of the
//      function.
//   3. Run:
//      go test -test.run=Example$ github.com/tebeka/selenium
func Example() {
	// Start a Selenium WebDriver server instance (if one is not already
	// running).
	const (
		// These paths will be different on your system.
		seleniumPath    = "selenium-server-standalone-3.9.0.jar"
		geckoDriverPath = "geckodriver-v24.0.exe"
		port            = 8080
	)
	opts := []selenium.ServiceOption{
		//selenium.StartFrameBuffer(),         // Start an X frame buffer for the browser to run in.
		selenium.GeckoDriver(geckoDriverPath), // Specify the path to GeckoDriver in order to use Firefox.
		selenium.Output(os.Stderr) ,          // Output debug information to STDERR.
	}
	selenium.SetDebug(true)
	service, err := selenium.NewSeleniumService(seleniumPath, port, opts...)
	if err != nil {
		panic(err) // panic is used only as an example and is not otherwise recommended.
	}
	defer service.Stop()

	// Connect to the WebDriver instance running locally.
	caps := selenium.Capabilities{"browserName": "firefox"}
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		panic(err)
	}


	// Navigate to the simple playground interface.
	if err := wd.Get("https://mail.qq.com"); err != nil {
		panic(err)
	}

	// Get a reference to the text box containing code.
	time.Sleep(7*time.Second)


	//result2,err:=wd.FindElement(selenium.ByXPATH,"/html/body/div[1]/div[1]/div/div[3]/a[2]")
	//println(result2,"33344444")
	//time.Sleep(5*time.Second)
	//result2.Click()
	//searchbench,err:=wd.FindElement(selenium.ByXPATH,"//*[@id='uinArea']")
	//searchbench2,err:=wd.FindElement(selenium.ByXPATH,"//*[@id='u']")
	a,err:=wd.FindElement(selenium.ByXPATH,"//*[@id='qqLoginTab']")
	if err !=nil{
		println("select error")
	}
	a.Click()
	time.Sleep(2*time.Second)
	s,err:=wd.FindElement(selenium.ByXPATH,"//*[@id='uinArea']")
	println(s,"3333333333333")
	//searchbench.SendKeys("1030637731")
	//pwd,err:=wd.FindElement(selenium.ByXPATH,"//*[@id='pwdArea']")
	//pwd.SendKeys("aaa")
	//searchasure,err:=wd.FindElement(selenium.ByXPATH,"//*[@id='login_button']")
	//time.Sleep(1*time.Second)
	//searchasure.Click()
	time.Sleep(100*time.Second)
	defer wd.Quit()
}