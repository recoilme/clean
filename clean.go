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

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"golang.org/x/net/html/charset"
)

var (
	whiteTags = " address, article, aside, a, abbr, b, bdi, bdo, big, blockquote, body, br, caption, cite, code, data, dd, del, dfn, div, dl, dt, em, figcaption, figure, footer, h1, h2, h3, h4, h5, h6, html, header, hr, i, img, ins, kbd, li, main, mark, nav, ol, p, pre, q, rp, rt, ruby, s, samp, section, small, span, strike, strong, sub, sup, table, tbody, td, template, tfoot, th, thead, time, tr, tt, u, ul, var,"
)

// Preprocess some string
func Preprocess(fragment string, pretty bool) (string, error) {

	var buf bytes.Buffer

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
								buf.WriteString(fmt.Sprintf("<img src=\"%s\"/><br/>", string(v)))
							} else {
								buf.WriteString(fmt.Sprintf("<img src=\"%s\"/>", string(v)))
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
									buf.WriteString(fmt.Sprintf(" <a href=\"%s\"> ", string(v)))
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
										buf.WriteString(fmt.Sprintf("<p><img src=\"%s\"></p>", string(v)))
									} else {
										buf.WriteString(fmt.Sprintf("<img src=\"%s\">", string(v)))
									}
								}
								if !m {
									break
								}
							}
						}
					default:
						if tag.String() == "div" && pretty {
							if divCnt < 0 {
								divCnt = 0
							}
							buf.WriteString(fmt.Sprintf("\n%s", strings.Repeat(" ", divCnt)))
							divCnt++
						}
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
func Clean(s string) (string, error) {
	s, err := Preprocess(s, false)
	if err != nil {
		return s, err
	}
	re := regexp.MustCompile("<\\w*>\\s*\\^*\\$*</\\w*>")
	for {
		before := s
		s = re.ReplaceAllString(s, "")
		if before == s {
			break
		}
	}
	return Preprocess(s, true)
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
func URI(u string) (s string, err error) {
	s, err = GetUtf8(u)
	if err != nil {
		return
	}
	s, err = Clean(s)
	if err != nil {
		return
	}
	return
}
