// Command pdf is a chromedp example demonstrating how to capture a pdf of a
// page.
package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

func main() {
	// create context

	opts := append(chromedp.DefaultExecAllocatorOptions[:], chromedp.Flag("headless", false))
	ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel = chromedp.NewContext(
		ctx,
		chromedp.WithLogf(log.Printf))
	defer cancel()

	// capture pdf
	var buf []byte
	if err := chromedp.Run(ctx, printToPDF(`https://mp.weixin.qq.com/s/Yf2pHXKwUjsLJ6J2wlnrOA`, &buf)); err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile("sample.pdf", buf, 0o644); err != nil {
		log.Fatal(err)
	}
	fmt.Println("wrote sample.pdf")
}

// var body = document.body,
//     html = document.documentElement;

// var height = Math.max( body.scrollHeight, body.offsetHeight,
//                        html.clientHeight, html.scrollHeight, html.offsetHeight );

// print a specific pdf page.
func printToPDF(urlstr string, res *[]byte) chromedp.Tasks {
	tasks := chromedp.Tasks{}

	tasks = append(tasks,
		chromedp.Navigate(urlstr),
		chromedp.EvaluateAsDevTools("document.documentElement.scrollTo(0, 0)", nil),
		chromedp.EvaluateAsDevTools("document.documentElement.scrollTo(Number.MAX_SAFE_INTEGER, Number.MAX_SAFE_INTEGER)", nil),
		// chromedp.WaitVisible(`document.querySelector("#js_content > section:nth-child(53) > section:nth-child(6) > img")`, chromedp.ByJSPath),
	)
	// var nodes []*cdp.Node
	tasks = append(tasks,
		// chromedp.Nodes(`document.querySelector("#js_content > section:nth-child(53) > section:nth-child(5) > p:nth-child(5)")`, &nodes, chromedp.ByJSPath),
		// chromedp.MouseClickNode(nodes[0], chromedp.ButtonType(input.Right)),
		// chromedp.EvaluateAsDevTools("document.documentElement.scrollTo(0, 0)", nil),
		// chromedp.EvaluateAsDevTools("document.documentElement.scrollTo(Number.MAX_SAFE_INTEGER, Number.MAX_SAFE_INTEGER)", nil),
		// chromedp.WaitVisible(`document.querySelector("#js_content > section:nth-child(53) > section:nth-child(5) > p:nth-child(5) > img")`, chromedp.ByJSPath),
		chromedp.Sleep(5*time.Second),
		chromedp.EvaluateAsDevTools("document.documentElement.scrollTo(0, 0)", nil),
		chromedp.EvaluateAsDevTools("document.documentElement.scrollTo(Number.MAX_SAFE_INTEGER, Number.MAX_SAFE_INTEGER)", nil),
		chromedp.ActionFunc(func(ctx context.Context) error {
			buf, _, err := page.PrintToPDF().WithPaperHeight(74.71875).WithPrintBackground(false).Do(ctx)
			if err != nil {
				return err
			}
			*res = buf
			return nil
		}),
	)

	return tasks
}
