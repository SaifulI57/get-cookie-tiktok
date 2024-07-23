package main

import (
	"fmt"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

// func main() {
// 	// Launch a new browser instance
// 	url := launcher.New().Headless(true).MustLaunch()
// 	browser := rod.New().ControlURL(url).MustConnect()
// 	defer browser.MustClose()

// 	// Create a new page
// 	page := browser.MustPage()
// 	var wg sync.WaitGroup
// 	// Enable request interception
// 	wg.Add(1)
// 	// Navigate to the target page
// 	go func() {
// 		defer wg.Done()
// 		page.MustNavigate("https://www.tiktok.com/explore")
// 		// Wait for navigation to complete
// 		page.MustWaitLoad()

// 	}()

// 	// Extract and print the cookies
// 	wg.Wait()

// 	cookies := page.MustCookies()
// 	fmt.Println("\nCookies:")
// 	for _, cookie := range cookies {
// 		fmt.Printf("%s: %s\n", cookie.Name, cookie.Value)
// 	}
// }

func main() {
	// Launch a new browser instance
	url := launcher.New().Headless(true).MustLaunch()
	browser := rod.New().ControlURL(url).MustConnect()
	defer browser.MustClose()

	// Create a new page
	page := browser.MustPage()

	// Enable request interception

	// Navigate to the target page
	page.MustNavigate("https://www.tiktok.com/explore")

	// Wait for navigation to complete
	page.MustWaitLoad()

	// Extract and print the cookies
	cookies := page.MustCookies()
	fmt.Println("\nCookies:")
	for _, cookie := range cookies {
		fmt.Printf("%s:%s;", cookie.Name, cookie.Value)
	}
}
