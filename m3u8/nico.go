package m3u8

import (
	"context"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"html"
	"log"
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

func getAPIData(url string) string {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	var dataContent string
	domain := ".nicovideo.jp"
	err := chromedp.Run(ctx,
		setCookie("nicosid", CookiesMap["nicosid"], domain, "/", false, false),
		setCookie("_ss_pp_id", CookiesMap["_ss_pp_id"], domain, "/", false, false),
		setCookie("_td", CookiesMap["_td"], domain, "/", false, false),
		//SetCookie("nico_gc", "", domain, "/", false, false),
		setCookie("user_session", CookiesMap["user_session"], domain, "/", true, true),
		setCookie("user_session_secure", CookiesMap["user_session_secure"], domain, "/", true, true),
		//SetCookie("domand_bid", "", domain, "/", true, true),
		chromedp.Navigate(url),
		chromedp.PollFunction(pollFunction, &dataContent, chromedp.WithPollingMutation()),
		//chromedp.Stop(),
	)
	if err != nil {
		log.Fatal(err)
	}
	unescapedData := html.UnescapeString(dataContent)
	return unescapedData
}

func setCookie(name, value, domain, path string, httpOnly, secure bool) chromedp.Action {
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

func setHeaders(headers map[string]interface{}) chromedp.Tasks {
	return chromedp.Tasks{
		network.Enable(),
		network.SetExtraHTTPHeaders(headers),
	}
} // dummy header
