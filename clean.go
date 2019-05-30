package clean

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
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
							buf.WriteString(fmt.Sprintf(" <img src=\"%s\"/> ", fixURL(string(v))))
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
				if strings.Contains(whiteTags, " "+tag.String()+",") {
					switch tag.String() {
					case "html":
						buf.WriteString("<html>")
					case "a":
						if hasAttr {
							for {
								k, v, m := t.TagAttr()
								if string(k) == "href" {
									buf.WriteString(fmt.Sprintf(" <a href=\"%s\">", fixURL(string(v))))
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
									buf.WriteString(fmt.Sprintf(" <img src=\"%s\">", fixURL(string(v))))
								}
								if !m {
									break
								}
							}
						}
					case "span":
						buf.WriteByte(' ')
					case "div":
						if pretty {
							buf.WriteByte(' ')
						} else {
							buf.WriteString(fmt.Sprintf("<%s>", tag.String()))
						}
					default:
						if pretty {
							buf.WriteString(fmt.Sprintf(" <%s> ", tag.String()))
						} else {
							buf.WriteString(fmt.Sprintf("<%s>", tag.String()))
						}
					}
					inTag = tag.String()
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
			tagName, _ := t.TagName()
			if tag := atom.Lookup(tagName); tag != 0 {
				if strings.Contains(whiteTags, " "+tag.String()+",") {
					switch tag.String() {
					case "div":
						if pretty {
							buf.WriteByte(' ')
						} else {
							buf.WriteString(fmt.Sprintf("</%s>", tag.String()))
						}
					case "span":
						buf.WriteByte(' ')
					default:
						if pretty {
							buf.WriteString(fmt.Sprintf(" </%s> ", tag.String()))
						} else {
							buf.WriteString(fmt.Sprintf("</%s>", tag.String()))
						}
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
	//ioutil.WriteFile("in.htm", []byte(s), 0666)
	n := MainNode(s, baseURL.Hostname())
	res := ""
	if n != nil {
		res = renderNode(n)
	}
	s, err = Preprocess(res, true, baseURL)
	if err != nil {
		return s, err
	}
	space := regexp.MustCompile(`\s+`)
	s = space.ReplaceAllString(s, " ")
	//re := regexp.MustCompile("<\\w*>\\s*</\\w*>")
	//for {
	//before := s
	//s = re.ReplaceAllString(s, "")
	//if before == s {
	//	break
	//}
	//}
	s = "<!DOCTYPE html>\n<html>\n<head>\n <meta charset=\"utf-8\"><link rel=\"stylesheet\" href=\"https://cdn.jsdelivr.net/gh/kognise/water.css@latest/dist/dark.min.css\">\n</head>\n<body>" + s
	s = s + "</body></html>"
	return s, nil
}

// MainNode score all nodes by textDencity and internal link dencity
func MainNode(s, host string) *html.Node {
	var maxNode *html.Node

	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(s))

	maxsel := doc.Find("body")
	txt, link, linkint := NodeDen(doc, maxsel, host)
	pageText := txt - link
	pageInternalLink := linkint

	max := -1.
	var d func(*goquery.Selection)
	d = func(s *goquery.Selection) {
		nodes := s.Nodes
		if len(nodes) == 0 {
			return
		}
		for _, n := range nodes {
			nodesel := doc.FindNodes(n)
			txt, link, intern := NodeDen(doc, nodesel, host)

			scoreTxt := 0.
			if pageText > 100 {
				scoreTxt = ((txt - link) / pageText)
			}
			scoreLink := 0.
			if pageInternalLink > 10 {
				scoreLink = float64(1 - intern/pageInternalLink)
			}
			scoreElem := 0.
			if n.Data == "div" || n.Data == "main" || n.Data == "article" {
				scoreElem = .2
			}
			if scoreTxt > .05 && ((scoreTxt + +scoreLink + scoreElem) > max) {
				max = scoreTxt + scoreLink + scoreElem
				maxNode = n
			}
		}
		s = s.Children()
		d(s)
	}
	d(maxsel.Children())
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

// NodeDen calc textNode, textinlinks and count of internal link in node
func NodeDen(doc *goquery.Document, s *goquery.Selection, host string) (float64, float64, float64) {
	tr := strings.TrimSpace(s.Text())
	txtCntN := float64(len([]rune(tr)))
	linkText := ""
	s.Find("a").Each(func(n int, s *goquery.Selection) {
		linkText += strings.TrimSpace(s.Text())
	})

	linksCntN := float64(len([]rune(linkText)))
	internalLink := 0.
	s.Find("a").Each(func(n int, s *goquery.Selection) {
		if href, ok := s.Attr("href"); ok {
			if strings.Contains(href, host) {
				internalLink++
			}
		}
	})
	return txtCntN, linksCntN, internalLink
}
