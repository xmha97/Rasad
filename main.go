package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

type Site struct {
	Name string
	URL  string
}

func DefaultSites() []Site {
	// {Name: "Aparat", URL: "https://www.aparat.com/"},
	// {Name: "Digikala", URL: "https://www.digikala.com/"},
	// {Name: "Divaar", URL: "https://divar.ir/"},
	// {Name: "Bale", URL: "https://ble.ir/"},
	// {Name: "Eitaa", URL: "https://eitaa.com/"},
	// {Name: "Eitaa (Web)", URL: "https://web.eitaa.com/"},
	// {Name: "Eitaa (Map)", URL: "https://map.eitaa.com/"},
	// {Name: "Eitaa Server (Ali-com)", URL: "https://ali.eitaa.com/eitaa/"},
	// {Name: "Eitaa Server (Alireza-com)", URL: "https://alireza.eitaa.com/eitaa/"},
	// {Name: "Eitaa Server (Alireza-ir)", URL: "https://alireza.eitaa.ir/eitaa/"},
	// {Name: "Eitaa Server (Alzheimer-com)", URL: "https://alzheimer.eitaa.com/eitaa/"},
	// {Name: "Eitaa Server (Armita-com)", URL: "https://armita.eitaa.com/eitaa/"},
	// {Name: "Eitaa Server (Armita-ir)", URL: "https://armita.eitaa.ir/eitaa/"},
	// {Name: "Eitaa Server (Bagher-ir)", URL: "https://bagher.eitaa.ir/eitaa/"},
	// {Name: "Eitaa Server (Fateme-com)", URL: "https://fateme.eitaa.com/eitaa/"},
	// {Name: "Eitaa Server (Hasan-ir)", URL: "https://hasan.eitaa.ir/eitaa/"},
	// {Name: "Eitaa Server (Hosna-com)", URL: "https://hosna.eitaa.com/eitaa/"},
	// {Name: "Eitaa Server (Hosna-ir)", URL: "https://hosna.eitaa.ir/eitaa/"},
	// {Name: "Eitaa Server (Kazem-ir)", URL: "https://kazem.eitaa.ir/eitaa/"},
	// {Name: "Eitaa Server (Mahdi-com)", URL: "https://mahdi.eitaa.com/eitaa/"},
	// {Name: "Eitaa Server (Majid-com)", URL: "https://majid.eitaa.com/eitaa/"},
	// {Name: "Eitaa Server (Majid-ir)", URL: "https://majid.eitaa.ir/eitaa/"},
	// {Name: "Eitaa Server (Meysam-com)", URL: "https://meysam.eitaa.com/eitaa/"},
	// {Name: "Eitaa Server (Mostafa-com)", URL: "https://mostafa.eitaa.com/eitaa/"},
	// {Name: "Eitaa Server (Mostafa-ir)", URL: "https://mostafa.eitaa.ir/eitaa/"},
	// {Name: "Eitaa Server (Sadegh-ir)", URL: "https://sadegh.eitaa.ir/eitaa/"},
	// {Name: "Eitaa Server (Sajad-ir)", URL: "https://sajad.eitaa.ir/eitaa/"},
	// {Name: "Varzesh3", URL: "https://www.varzesh3.com/"},
	// {Name: "Shaparak", URL: "https://shaparak.ir/"},
	// {Name: "Blogfa", URL: "https://www.blogfa.com/"},
	// {Name: "Telewebion", URL: "https://telewebion.com/"},
	// {Name: "Irancell", URL: "https://irancell.ir/"},
	// {Name: "My Irancell", URL: "https://my.irancell.ir/"},
	// {Name: "Google", URL: "https://www.google.com/"},
	// {Name: "WhatsApp", URL: "https://www.whatsapp.com/"},
	// {Name: "WhatsApp (API)", URL: "https://api.whatsapp.com/"},
	// {Name: "WhatsApp (Shorten)", URL: "https://wa.me/"},
	// {Name: "FlatHub", URL: "https://www.flathub.org/"},
	// {Name: "FlatHub (Download)", URL: "https://dl.flathub.org/"},
	// {Name: "Docker", URL: "https://www.docker.com/"},
	// {Name: "GitHub", URL: "https://github.com/"},
	// {Name: "GitHub Pages", URL: "https://pages.github.com/"},
	// {Name: "GitHub (RAW Download)", URL: "https://raw.githubusercontent.com/"},
	// {Name: "Google Play", URL: "https://play.google.com/"},
	// {Name: "App Store", URL: "https://apps.apple.com/"},
	// {Name: "Google Translate", URL: "https://translate.google.com/"},
	// {Name: "Mozilla Developer", URL: "https://developer.mozilla.org/"},
	// {Name: "Google Docs", URL: "https://docs.google.com/"},

	// {Name: "Facebook", URL: "https://www.fb.com/"},
	// {Name: "Facebook", URL: "https://www.fb.me/"},
	// {Name: "Messenger", URL: "https://m.me/"},
	// {Name: "Telegram", URL: "https://web.telegram.org/"},
	// {Name: "Telegram", URL: "https://t.me/"},
	// {Name: "PornHub", URL: "https://www.pornhub.com/"},

	return []Site{
		{Name: "Google for Developers", URL: "https://developers.google.com/"},
		{Name: "Dart Packages", URL: "https://pub.dev/"},
		{Name: "Messenger", URL: "https://www.messenger.com/"},
		{Name: "Instagram", URL: "https://www.instagram.com/"},
		{Name: "Facebook", URL: "https://www.facebook.com/"},
		{Name: "Discord", URL: "https://discord.com/"},
		{Name: "Reddit", URL: "https://www.reddit.com/"},
		{Name: "Telegram", URL: "https://telegram.org/"},
		{Name: "Microsoft Store", URL: "https://apps.microsoft.com/"},
		{Name: "YouTube", URL: "https://www.youtube.com/"},
	}
}

func Test(name string, url string, count int) string {
	fmt.Printf(name + "                        ")
	outText := ""
	count++
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		outText += "❌ " + name + " | " + "NewRequest | " + err.Error()
		return outText
	}
	req.Header.Set("If-Modified-Since", "Wed, 21 Oct 2020 07:28:00 GMT")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		outText += handleError(name, url, count, err)
		return outText
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusNotModified:
		outText += "✅ " + name + " | " + "OK | NotModified"
	case http.StatusOK:
		outText += "✅ " + name + " | " + "OK | 200"
	case http.StatusForbidden:
		outText += "❌ " + name + " | " + "Forbidden | 403"
	case http.StatusUnavailableForLegalReasons:
		outText += "❌ " + name + " | " + "Forbidden | Illegal"
	default:
		outText += "❌ " + name + " | " + "Error | " + strconv.Itoa(resp.StatusCode) + http.StatusText(resp.StatusCode)
	}
	return outText
}

func handleError(name string, url string, count int, err error) string {
	outText := ""
	msg := err.Error()
	switch {
	case strings.Contains(msg, "dial tcp"):
		outText += "❌ " + name + " | " + "Blocked | Dial"
	case strings.Contains(msg, "read tcp"):
		outText += "❌ " + name + " | " + "Blocked | Read"
	case strings.Contains(msg, "TLS handshake timeout"):
		outText += "❌ " + name + " | " + "Timeout"
		if count < 3 {
			return Test(name, url, count)
		}
	case strings.Contains(msg, "forcibly closed"):
		outText += "❌ " + name + " | " + "Blocked | Closed"
	case strings.Contains(msg, "no such host"):
		outText += "❌ " + name + " | " + "Blocked | NoHost"
	default:
		outText += "❌ " + name + " | " + "SendRequest | " + err.Error()
	}
	return outText
}

func main() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	done := make(chan bool)
	var printData string
	go func() {
		spinner := []rune("⠋⠙⠹⠸⠼⠴⠦⠧⠇⠏")
		for {
			for _, r := range spinner {
				select {
				case <-done:
					fmt.Printf("\r" + printData)
					return
				default:
					fmt.Printf("\r%s Testing ", string(r))
					time.Sleep(100 * time.Millisecond)
				}
			}
		}
	}()

	printData = doTesting()
	done <- true
	time.Sleep(100 * time.Millisecond)
	fmt.Scanln()
}

func doTesting() string {
	sites := DefaultSites()
	outText := "╭────┬────────────────────────────────────────────╮\n"
	outText += "│  # │                    name                    │\n"
	outText += "├────┼────────────────────────────────────────────┤\n"
	var index int = 0
	for _, s := range sites {
		index++
		result := Test(s.Name, s.URL, 0)
		outText += "│ " + fmt.Sprintf("%2d", index) + " │ " + fmt.Sprintf("%-41s", result) + " │\n"
	}
	outText += "╰────┴────────────────────────────────────────────╯\n"
	return outText
}
