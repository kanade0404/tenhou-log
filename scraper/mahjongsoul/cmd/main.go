package main

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/runtime"
	"github.com/chromedp/chromedp"
	"github.com/sirupsen/logrus"
	"log"
	"strings"
)

type User struct {
	Name    string
	ID      string
	RoomIDs []string
}

const (
	baseURL = "https://amae-koromo.sapk.ch/player"
)

func main() {
	var (
		users = []User{
			{
				Name: "kanade0404",
				ID:   "73281085",
				RoomIDs: []string{
					"12",
					"11",
				},
			},
		}
	)
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false), // headless=false に変更
		chromedp.Flag("disable-gpu", false),
		chromedp.Flag("enable-automation", false),
		chromedp.Flag("disable-extensions", false),
		chromedp.Flag("hide-scrollbars", false),
		chromedp.Flag("mute-audio", false))
	alc, _ := chromedp.NewExecAllocator(context.Background(), opts...)
	ctx, cancel := chromedp.NewContext(alc, chromedp.WithLogf(log.Printf))
	defer cancel()
	for _, user := range users {
		url := CreateScrapingURL(user.ID, user.RoomIDs...)
		fmt.Println(url)
		if err := chromedp.Run(ctx, chromedp.Navigate(url)); err != nil {
			log.Fatal("error navigate page", err)
		}
		sel := ".ReactVirtualized__Grid.ReactVirtualized__Table__Grid"
		if err := chromedp.Run(ctx, chromedp.WaitVisible(sel)); err != nil {
			log.Fatal(fmt.Sprintf("error wait visible(sel:%s)", sel), err)
		}
		rowSel := ".ReactVirtualized__Table__row"
		var (
			nodeMap = make(map[string]*cdp.Node)
			count   = 1
		)
		isComplete := false
		for !isComplete {
			fmt.Println("start loop")
			fmt.Println("isComplete: ", isComplete)
			var nodes []*cdp.Node
			if err := chromedp.Run(ctx, chromedp.Nodes(rowSel, &nodes)); err != nil {
				log.Fatal(fmt.Sprintf("error query selector(%s)", sel), err)
			}
			fmt.Println(fmt.Sprintf("node len: %d", len(nodes)))
			if len(nodes) == 0 {
				isComplete = true
			}

			for _, node := range nodes {
				ariaRowIndex, exist := node.Attribute("aria-rowindex")
				if !exist {
					continue
				}
				fmt.Println(fmt.Sprintf("ariaRowIndex: %s", ariaRowIndex))
				if _, ok := nodeMap[ariaRowIndex]; ok {
					continue
				}
				fmt.Println(fmt.Sprintf("%+v", nodeMap))
				nodeMap[ariaRowIndex] = node
				var children []*cdp.Node
				if err := chromedp.Run(ctx, chromedp.Nodes(".ReactVirtualized__Table__rowColumn > div > span > a[href^=\"https://\"]", &children, chromedp.ByQueryAll, chromedp.FromNode(node))); err != nil {
					log.Fatal("error children query", err)
				}
				fmt.Println(fmt.Sprintf("children len(ariaRowIndex: %s): %d", ariaRowIndex, len(children)))
				for _, child := range children {
					href, exist := child.Attribute("href")
					logrus.Info(fmt.Sprintf("href: %s, exist: %t", href, exist))
					if !exist {
						continue
					}
					textContent, exist := child.Attribute("textContent")
					logrus.Info(fmt.Sprintf("textContent: %s, %t", textContent, exist))
					if !exist || strings.Contains(textContent, user.Name) {
						continue
					}
					fmt.Println(href, textContent)
				}

			}
			// 下にスクロールして表示を待つ
			if err := chromedp.Run(ctx, chromedp.ActionFunc(func(ctx context.Context) error {
				_, exp, err := runtime.Evaluate(fmt.Sprintf("window.scrollTo(0, window.innerHeight * %d);", count)).Do(ctx)
				if err != nil {
					return err
				}
				if exp != nil {
					return exp
				}
				return nil
			})); err != nil {
				log.Fatal("error scroll to bottom", err)
			}
			count += 1
			fmt.Println("count: ", count)
			fmt.Println("end loop")
		}
	}
}

func CreateScrapingURL(userID string, roomIDs ...string) string {
	fmt.Println(userID, roomIDs)
	return fmt.Sprintf("%s/%s/%s", baseURL, userID, strings.Join(roomIDs, "."))
}
