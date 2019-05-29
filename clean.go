package clean

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"golang.org/x/net/html/charset"
)

var (
	whiteTags = " address, article, aside, a, abbr, b, bdi, bdo, big, blockquote, body, br, caption, cite, code, data, dd, del, dfn, div, dl, dt, em, figcaption, figure, footer, h1, h2, h3, h4, h5, h6, html, header, hr, i, img, ins, kbd, li, main, mark, nav, ol, p, pre, q, rp, rt, ruby, s, samp, section, small, span, strike, strong, sub, sup, table, tbody, td, template, tfoot, th, thead, time, tr, tt, u, ul, var,"
)

type density struct {
	chars    int
	tags     int
	linkchar int
	linktag  int
	txt      string
}

type stat struct {
	txt float64
	lin float64
}

// Preprocess some string
func Preprocess(fragment string, pretty bool, base *url.URL) (string, error) {

	var buf bytes.Buffer

	fixURL := func(u string) string {
		if base == nil {
			return u
		}
		uNew, err := url.Parse(u)
		if err == nil {
			uNew.Fragment = ""
			uNew.RawQuery = ""
			return base.ResolveReference(uNew).String()
		}
		return u
	}

	t := html.NewTokenizer(strings.NewReader(fragment))
	inTag := ""
	divCnt := 0
	for {
		switch tok := t.Next(); tok {
		case html.ErrorToken:
			if t.Err() == io.EOF {
				return buf.String(), nil
			}
			return buf.String(), t.Err()
		case html.SelfClosingTagToken:
			tagName, hasAttr := t.TagName()
			if string(tagName) == "br" {
				buf.WriteString("<br/>")
			}
			if string(tagName) == "img" {
				if hasAttr {
					for {
						k, v, m := t.TagAttr()
						if string(k) == "src" {
							if pretty {
								buf.WriteString(fmt.Sprintf("<img src=\"%s\"/><br/>", fixURL(string(v))))
							} else {
								buf.WriteString(fmt.Sprintf("<img src=\"%s\"/>", fixURL(string(v))))
							}
						}
						if !m {
							break
						}
					}
				}
			}
		case html.StartTagToken:
			tagName, hasAttr := t.TagName()
			if tag := atom.Lookup(tagName); tag != 0 {
				inTag = ""
				if strings.Contains(whiteTags, " "+tag.String()+",") { //_, ok := Cfg.whiteList[tag]; ok {

					switch tag.String() {
					case "html":
						head := "<!DOCTYPE html>\n<html>\n<head>\n <meta charset=\"utf-8\"><link rel=\"stylesheet\" href=\"https://cdn.jsdelivr.net/gh/kognise/water.css@latest/dist/dark.min.css\">\n</head>\n"
						if !pretty {
							head = "<html>"
						}
						buf.WriteString(head)
					case "a":
						if hasAttr {
							for {
								k, v, m := t.TagAttr()
								if string(k) == "href" {
									buf.WriteString(fmt.Sprintf(" <a href=\"%s\"> ", fixURL(string(v))))
								}
								if !m {
									break
								}
							}
						}
					case "img":
						if hasAttr {
							for {
								k, v, m := t.TagAttr()
								if string(k) == "src" {
									if pretty {
										buf.WriteString(fmt.Sprintf("<p><img src=\"%s\"></p>", fixURL(string(v))))
									} else {
										buf.WriteString(fmt.Sprintf("<img src=\"%s\">", fixURL(string(v))))
									}
								}
								if !m {
									break
								}
							}
						}
					case "span":
						buf.WriteByte(' ')
					case "p":
						buf.WriteString(fmt.Sprintf(" <%s>", tag.String()))
					case "em":
						buf.WriteString(fmt.Sprintf(" <%s>", tag.String()))
					case "div":
						if tag.String() == "div" && pretty {
							if divCnt < 0 {
								divCnt = 0
							}
							buf.WriteString(fmt.Sprintf("\n%s", strings.Repeat(" ", divCnt)))
							divCnt++
						}
						buf.WriteString(fmt.Sprintf("<%s>", tag.String()))
					default:

						if tag.String() == "nav" && pretty {
							buf.WriteString(fmt.Sprintf("<details><summary>Navigation</summary><%s>", tag.String()))
						}
						buf.WriteString(fmt.Sprintf("<%s>", tag.String()))
					}
					inTag = tag.String()

				} else {
					//log.Printf("not in tag %+v\n", tag)
				}
			}

		case html.TextToken:

			if inTag != "" {
				content := TrimBytes(t.Raw())
				if content != nil {
					buf.Write(content)
				}
			}
		case html.EndTagToken:
			//inTag = ""
			tagName, _ := t.TagName()
			if tag := atom.Lookup(tagName); tag != 0 {
				//if _, ok := Cfg.whiteList[tag]; ok {
				if strings.Contains(whiteTags, " "+tag.String()+",") {
					switch tag.String() {
					case "a":
						buf.WriteString(fmt.Sprintf("</%s> ", tag.String()))
					case "div":
						if pretty {
							divCnt--
							if divCnt < 0 {
								divCnt = 0
							}
							buf.WriteString(fmt.Sprintf("\n%s", strings.Repeat(" ", divCnt)))
						}
						buf.WriteString(fmt.Sprintf("</%s>", tag.String()))
					case "span":
						buf.WriteByte(' ')
					case "p":
						buf.WriteString(fmt.Sprintf(" </%s>", tag.String()))
					case "em":
						buf.WriteString(fmt.Sprintf(" </%s>", tag.String()))
					default:
						if tag.String() == "nav" && pretty {
							buf.WriteString(fmt.Sprintf("</%s></details>", tag.String()))
						}
						buf.WriteString(fmt.Sprintf("</%s>", tag.String()))
					}
				}
			}
		}
	}
}

//Clean return cleaned html
func Clean(s string, extract bool, baseURL *url.URL) (string, error) {
	s, err := Preprocess(s, false, baseURL)
	if err != nil {
		return s, err
	}
	ioutil.WriteFile("in.htm", []byte(s), 0666)

	n := MainNode(s, baseURL.Hostname())
	res := ""
	if n != nil {
		res = renderNode(n)
	}

	//ioutil.WriteFile("out.htm", []byte(res), 0666)

	return Preprocess(res, true, baseURL)
}

func MainNode(s, host string) *html.Node {
	var maxNode *html.Node

	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(s))

	maxsel := doc.Find("body")
	txt, link, linkint := NodeDen2(doc, maxsel, host)
	pageText := txt - link
	pageLink := link
	pageInternalLink := linkint

	log.Println(txt, link, pageInternalLink)

	max := -1.
	//var mainNode *goquery.Selection
	iter := 0
	var d func(*goquery.Selection)
	d = func(s *goquery.Selection) {
		nodes := s.Nodes
		if len(nodes) == 0 {
			return
		}

		for _, n := range nodes {
			nodesel := doc.FindNodes(n)
			txt, link, intern := NodeDen2(doc, nodesel, host)

			score1 := 0.
			if pageText > 100 {
				score1 = ((txt - link) / pageText)
			}
			score2 := 0.
			if pageLink > 100 {
				//	score2 = (1 - link/pageLink)
			}
			score3 := 0.
			if pageInternalLink > 10 {
				score3 = float64(1 - intern/pageInternalLink)
			}
			scoreElem := 0.
			if n.Data == "div" {
				scoreElem = .2
			}
			if n.Data == "main" {
				//log.Println(n.Data, max, score1, score2, score3, scoreElem, intern)
				scoreElem = .2
			}
			if n.Data == "article" {
				scoreElem = .2
			}
			if score1 > .05 && ((score1 + score2 + score3 + scoreElem) > max) {
				max = score1 + score2 + score3 + scoreElem
				//	mainNode = nodesel
				maxNode = n
				log.Println(n.Data, max, score1, score2, score3, scoreElem, intern)

				if iter == 1111 {
					return
				}
				iter++
			}
			_ = score3

		}
		//return

		s = s.Children()

		d(s)
	}
	d(maxsel.Children())

	/*
		doc, _ := goquery.NewDocumentFromReader(strings.NewReader(s))

		maxsel := doc.Find("body")
		txt, _ := NodeDen(doc, maxsel, 0.3)
		pageText := txt
		log.Println("pageText", pageText)

		max := -1.
		ld := 0.
		td := 0.
		var maxNode *html.Node
		var d func(*goquery.Selection)
		d = func(s *goquery.Selection) {
			nodes := s.Nodes
			if len(nodes) == 0 {
				return
			}

			for _, n := range nodes {
				nodesel := doc.FindNodes(n)
				txt, link := NodeDen(doc, nodesel, 0.3)
				score := 0.
				if txt != 0 && pageText != 0 {
					score = 0.8*(1-((txt-link)/txt)) + 0.2*(txt/pageText)
				}

				if score > max && (txt/pageText) > 0.2 {

					max = score
					maxNode = n
					ld = (txt - link) / txt
					td = txt / pageText
					log.Println(maxNode.Data, max, txt, link, (txt-link)/txt, (txt / pageText), link)
				}
				//log.Println(n.Data, txt, link, score)
			}
			s = s.Children()
			d(s)
		}
		d(maxsel.Children())

		log.Println("----")
		_ = ld
		_ = td
	*/
	/*
		if maxNode == nil {
			return maxNode
		}
		maxsel = doc.FindNodes(maxNode)
		var maxNode2 *html.Node
		ld2 := 0.
		td2 := 0.
		for {
			maxsel = maxsel.Children()
			max2 := -1.
			for i, n := range maxsel.Nodes {
				_ = i
				nodesel := doc.FindNodes(n)
				txt, link := NodeDen(doc, nodesel, 0.3)
				score := 0.
				if txt != 0 && pageText != 0 {
					score = 0.8*((txt-link)/txt) + 0.2*(txt/pageText)
				}
				if score > max2 && (txt/pageText) > 0.2 {
					max2 = score
					ld2 = (txt - link) / txt
					td2 = txt / pageText

					maxsel = nodesel
					maxNode2 = n
					log.Println(n.Data, score, txt, link, (txt-link)/txt, (txt / pageText), link)
				}

			}
			if max2 < max*1.0 || ld2 < ld*.8 || td2 < td*.64 {
				break
			} else {
				log.Println("newnode")
				ld = ld2
				td = td2
				max = max2
				maxNode = maxNode2
			}
		}
	*/
	return maxNode

}

// TrimBytes remove spaces and /r/n
func TrimBytes(input []byte) []byte {
	b := bytes.Replace(input, []byte("\r\n"), nil, -1)
	b = bytes.Replace(input, []byte("\n"), nil, -1)
	b = bytes.Replace(input, []byte("\t"), nil, -1)
	return bytes.TrimSpace(b)
}

// GetUtf8 return utf8 string from url
func GetUtf8(geturl string) (s string, err error) {

	defHeaders := make(map[string]string)
	defHeaders["User-Agent"] = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.14; rv:65.0) Gecko/20100101 Firefox/65.0"
	//"Mozilla/5.0 (iPhone; CPU iPhone OS 12_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/12.0 Mobile/15E148 Safari/604.1"
	defHeaders["Accept"] = "text/html,application/xhtml+xml,application/xml,application/rss+xml;q=0.9,image/webp,*/*;q=0.8"
	defHeaders["Accept-Language"] = "ru-RU,ru;q=0.8,en-US;q=0.5,en;q=0.3"
	t := time.Second * 10

	ctx, cncl := context.WithTimeout(context.Background(), t)
	defer cncl()
	q := geturl
	if !strings.HasPrefix(geturl, "http") {
		q = "http://" + geturl
	}
	req, err := http.NewRequest(http.MethodGet, q, nil)
	if err != nil {
		return s, err
	}
	//Host
	u, err := url.Parse(q)
	if err == nil && len(u.Host) > 2 {
		req.Header.Set("Host", u.Host)
	}
	for k, v := range defHeaders {
		req.Header.Set(k, v)
	}
	resp, err := http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		return s, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		utf8, err := charset.NewReader(resp.Body, resp.Header.Get("Content-Type"))
		if err != nil {
			return s, err
		}
		body, err := ioutil.ReadAll(utf8)
		if err != nil {
			return s, err
		}
		return string(body), err
	}
	return s, fmt.Errorf("Error, statusCode:%d", resp.StatusCode)
}

// URI get content from url and clean
func URI(u string, extract bool) (s string, err error) {
	s, err = GetUtf8(u)
	if err != nil {
		return
	}
	uNew, err := url.Parse(u)
	if err != nil {
		uNew = nil
	}
	s, err = Clean(s, extract, uNew)
	if err != nil {
		return
	}
	return
}

func renderNode(n *html.Node) string {
	var buf bytes.Buffer
	w := io.Writer(&buf)
	html.Render(w, n)
	return buf.String()
}

func NodeDen(doc *goquery.Document, s *goquery.Selection, threshold float64) (float64, float64) {
	tr := strings.TrimSpace(s.Text())
	txtCntN := float64(len([]rune(tr)))
	linkText := ""
	s.Find("a").Each(func(n int, s *goquery.Selection) {
		linkText += strings.TrimSpace(s.Text())
	})
	//log.Println("\n", linkText)
	linksCntN := float64(len([]rune(linkText)))
	return txtCntN, linksCntN

}

func NodeDen2(doc *goquery.Document, s *goquery.Selection, host string) (float64, float64, float64) {
	tr := strings.TrimSpace(s.Text())
	txtCntN := float64(len([]rune(tr)))
	linkText := ""
	s.Find("a").Each(func(n int, s *goquery.Selection) {
		linkText += strings.TrimSpace(s.Text())
	})

	//log.Println("\n", linkText)
	linksCntN := float64(len([]rune(linkText)))
	pageInternalLink := 0.
	s.Find("a").Each(func(n int, s *goquery.Selection) {
		if href, ok := s.Attr("href"); ok {
			if strings.Contains(href, host) {
				pageInternalLink++
			}
		}
	})
	return txtCntN, linksCntN, pageInternalLink

}
