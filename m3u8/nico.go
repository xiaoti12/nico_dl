package m3u8

import (
	"bytes"
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"html"
	"log"
	"os"
	"time"
)

const pollFunction = `()=>{
  const e = document.getElementById("js-initial-watch-data");
  if (e) {
	return e.getAttribute("data-api-data");
  }
  // return a falsy value so that it will keep polling
  return "";
}`

func getAPIData(url string) []byte {
	resp, err := Client.R().
		SetHeader("Cookie", Cookies).
		Get(url)
	if err != nil {
		fmt.Printf("Error getting nico website: %s\n", err)
		return nil
	}
	time.Sleep(5 * time.Second)
	docReader := bytes.NewReader(resp.Body())
	doc, err := goquery.NewDocumentFromReader(docReader)
	if err != nil {
		fmt.Printf("Error parsing nico website: %s\n", err)
		return nil
	}
	// write doc content to local file
	err = os.WriteFile("nico.html", resp.Body(), 0644)
	if err != nil {
		fmt.Printf("Error writing file: %s\n", err)
		return nil
	}
	fmt.Println("save nico website to local file: nico.html")
	doc.Find("#js-initial-watch-data").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
	})
	return nil
}

func getActiveData(url string) string {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	var dataContent string
	domain := ".nicovideo.jp"
	err := chromedp.Run(ctx,
		SetCookie("nicosid", CookiesMap["nicosid"], domain, "/", false, false),
		SetCookie("_ss_pp_id", CookiesMap["_ss_pp_id"], domain, "/", false, false),
		SetCookie("_td", CookiesMap["_td"], domain, "/", false, false),
		//SetCookie("nico_gc", "", domain, "/", false, false),
		SetCookie("user_session", CookiesMap["user_session"], domain, "/", true, true),
		SetCookie("user_session_secure", CookiesMap["user_session_secure"], domain, "/", true, true),
		//SetCookie("domand_bid", "", domain, "/", true, true),
		chromedp.Navigate(url),
		chromedp.PollFunction(pollFunction, &dataContent, chromedp.WithPollingMutation()),
		chromedp.Stop(),
	)
	if err != nil {
		log.Fatal(err)
	}
	unescapedData := html.UnescapeString(dataContent)
	return unescapedData
}

func SetCookie(name, value, domain, path string, httpOnly, secure bool) chromedp.Action {
	return chromedp.ActionFunc(func(ctx context.Context) error {
		expr := cdp.TimeSinceEpoch(time.Now().Add(180 * 24 * time.Hour))
		err := network.SetCookie(name, value).
			WithExpires(&expr).
			WithDomain(domain).
			WithPath(path).
			WithHTTPOnly(httpOnly).
			WithSecure(secure).
			Do(ctx)
		if err != nil {
			return err
		}
		return nil
	})
}
