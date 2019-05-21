package clean

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"

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
func Clean(s string, extract bool, baseURL *url.URL) (string, error) {
	s, err := Preprocess(s, false, baseURL)
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
	if extract {
		//ss := MainContent(s)
		s = MainContent(s)
	}
	return Preprocess(s, true, nil)
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

func MainContent(s string) string {
	type noded struct {
		node       *html.Node
		density    int
		densitySum int
	}

	nodesmap := make(map[*html.Node]int)
	node, err := html.Parse(strings.NewReader(s))
	removeBad(node)
	if err != nil {
		log.Fatal(err)
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if c.Type == html.ElementNode {
				nodesmap[c] = calcChars(c).textDensity()
			}
			f(c)
		}
	}
	f(node)

	var fSum func(int, *html.Node) int
	fSum = func(s int, node *html.Node) int {
		var ff func(*html.Node)
		sum := s
		ff = func(n *html.Node) {
			if n.Type == html.ElementNode {
				sum += nodesmap[n]
			}

			for c := n.FirstChild; c != nil; c = c.NextSibling {
				ff(c)
			}
		}
		ff(node)
		return sum
	}
	nodes := make([]*noded, 0)
	for n, d := range nodesmap {
		node := &noded{}
		node.density = d
		node.node = n
		node.densitySum = fSum(d, n)
		nodes = append(nodes, node)

	}

	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].densitySum > nodes[j].densitySum
	})

	var selected *noded
	for i, n := range nodes {
		if i > 9 {
			break
		}
		//log.Println(n.node.Data, n.density, n.densitySum)
		if n.node.Data == "html" {
			continue
		}
		if n.node.Data == "body" {
			continue
		}
		if selected == nil {
			selected = n
		}
		//log.Println(selected.node.Data, n.density, n.densitySum, int(float32(selected.densitySum)*0.8))
		if selected != nil && n.density >= int(float32(selected.density)*0.7) && n.densitySum >= int(float32(selected.densitySum)*0.7) {
			selected = n
			//log.Println("in")
		}
	}
	if selected != nil {
		//log.Println("selected", selected.node.Data, selected.density)
		return renderNode(selected.node)
	}
	return ""
}

func renderNode(n *html.Node) string {
	var buf bytes.Buffer
	w := io.Writer(&buf)
	html.Render(w, n)
	return buf.String()
}

func calcChars(n *html.Node) *density {
	txt := ""
	tags := -1
	linkchar := ""
	linktag := 0
	var f func(*html.Node, bool, bool)
	f = func(n *html.Node, isText, isHyper bool) {

		if n.Type == html.ElementNode {
			if isHyper {
				linktag++
			} else {
				tags++
			}

		}
		if n.Type == html.TextNode {

			if isHyper {
				linkchar += strings.TrimSpace(n.Data)
				txt += linkchar
			} else {
				if isText {
					txt += strings.TrimSpace(n.Data)
				}
			}

		}
		isText = isText || (n.Type == html.ElementNode)
		isHyper = isHyper || (n.Type == html.ElementNode && n.Data == "a")
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c, isText, isHyper)
		}
	}

	f(n, false, false)
	txt = (txt)
	linkchar = (linkchar)
	d := &density{chars: len(txt), tags: tags, linkchar: len(linkchar), linktag: linktag, txt: txt}
	return d
}

func (d *density) textDensity() int {
	//text_density = (1.0 * char_num / tag_num) * qLn((1.0 * char_num * tag_num) / (1.0 * linkchar_num * linktag_num))
	// / qLn(qLn(1.0 * char_num * linkchar_num / un_linkchar_num + ratio * char_num + qExp(1.0)));
	//return 0

	if d.chars == 0 {
		return 0
	}
	unlinkcharnum := d.chars - d.linkchar
	if d.tags <= 0 {
		d.tags = 1
	}
	if d.linkchar <= 0 {
		d.linkchar = 1
	}
	if d.linktag <= 0 {
		d.linktag = 1
	}
	if unlinkcharnum <= 0 {
		unlinkcharnum = 1
	}
	chisl := (float64(d.chars) / float64(d.tags)) * math.Log((float64(d.chars)*float64(d.tags))/(float64(d.linkchar)*float64(d.linktag)))
	//qLn(qLn(1.0*char_num*linkchar_num/un_linkchar_num + ratio*char_num + qExp(1.0)))
	znamen := math.Log( /*math.Log*/ (float64(d.chars)*float64(d.linkchar)/float64(unlinkcharnum) + 1.0*float64(d.chars) + math.Exp(1.0)))
	return int(chisl / znamen)
}

func removeNode(n *html.Node, bad map[string]struct{}) bool {
	var r bool
	// if note is script tag
	if n.Type == html.ElementNode {
		atom := strings.ToLower(n.Data)
		if _, ok := bad[atom]; ok {
			n.Parent.RemoveChild(n)
			return true
		}
	}
	// traverse DOM
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		r = removeNode(c, bad)
		if r {
			break
		}
	}
	return r
}

func removeBad(n *html.Node) {
	bad := map[string]struct{}{
		"style":    struct{}{},
		"script":   struct{}{},
		"svg":      struct{}{},
		"nav":      struct{}{},
		"aside":    struct{}{},
		"form":     struct{}{},
		"noscript": struct{}{},
		"xmp":      struct{}{},
		"textarea": struct{}{},
		"air":      struct{}{},
	}
	for {
		if !removeNode(n, bad) {
			break
		}
	}
}
