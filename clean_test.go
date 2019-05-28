package clean_test

import (
	"io/ioutil"
	"log"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/recoilme/clean"
	"github.com/stretchr/testify/assert"
)

func url2file(t *testing.T, u string) {
	s, err := clean.URI(u, true)
	assert.NoError(t, err)
	assert.NotEmpty(t, s)
	ioutil.WriteFile("out.htm", []byte(s), 0666)
}

func Test01(t *testing.T) {
	url2file(t, "https://vc.ru/services/67441-baza-dannyh-turisticheskogo-servisa-sletat-ru-s-pasportnymi-dannymi-klientov-vremenno-okazalas-v-otkrytom-dostupe")

}
func Test02(t *testing.T) {
	url2file(t, "https://stackoverflow.com/questions/15081119/any-way-to-use-html-parse-without-it-adding-nodes-to-make-a-well-formed-tree")
}

func Test03(t *testing.T) {
	url2file(t, "adme.ru")
}

func Test04(t *testing.T) {
	url2file(t, "https://www.adme.ru/svoboda-narodnoe-tvorchestvo/25-muzhchin-kotorye-nashli-sposob-sdelat-otcovstvo-chutochku-legche-2084015/")
}

func Test05(t *testing.T) {
	url2file(t, "https://t.me/s/recoilmeblog/245")
}

func Test06(t *testing.T) {
	url2file(t, "https://vc.ru/contest/67338-kpi")
}

func Test07(t *testing.T) {
	url2file(t, "https://habr.com")

}
func Test08(t *testing.T) {
	url2file(t, "https://www.producthunt.com/posts/emtech-brew")
}

func Test09(t *testing.T) {
	url2file(t, "https://www.zakon.kz/top/")
}

func Test10(t *testing.T) {
	url2file(t, "https://bash.im/quote/455905")
}

func Test11(t *testing.T) {
	url2file(t, "https://roem.ru/19-05-2019/277784/huawei-tykva/")
}

func Test12(t *testing.T) {
	url2file(t, "https://sntch.com/6-glavnyh-tsitat-iz-rechi-vladimira-zelenskogo-na-inauguratsii/")
}

func Test13(t *testing.T) {
	url2file(t, "https://sntch.com/6-glavnyh-faktov-o-prikvele-igry-prestolov-v-kotorom-rasskazhut-o-proishozhdenii-belyh-hodokov/")
}

func Test14(t *testing.T) {
	url2file(t, "https://www.niemanlab.org/2019/04/what-kind-of-local-news-is-facebook-featuring-on-today-in-crime-car-crashes-and-not-too-much-community/")
}

func Test15(t *testing.T) {
	url2file(t, "https://habr.com/ru/post/220983/")
}

func Test16(t *testing.T) {
	url2file(t, "http://www.rbc.ru/technology_and_media/20/05/2019/5ce2ebb19a794748b13ed662")
}

func Test17(t *testing.T) {
	url2file(t, "http://feedproxy.google.com/~r/cosmo/fashion/~3/hsbD7Vq278c/")
}

func Test20(t *testing.T) {
	url2file(t, "https://www.bfm.ru/news/414680")
}

func Test21(t *testing.T) {
	url2file(t, "https://www.nur.kz/1795089-cinovnika-izbili-do-komy-v-stepnogorske-napadavsego-ne-nasli-spusta-poltora-goda.html")
}

func Test22(t *testing.T) {
	url2file(t, "https://www.nur.kz/1795094-narkomanov-i-prostitutok-net-almatincy-o-zizni-v-alatauskom-rajone-video.html")
}

func Test23(t *testing.T) {
	url2file(t, "https://tengrinews.kz/kazakhstan_news/stoit-udivlyatsya-poyavleniyu-elbasyi-mejdunarodnoy-arene-369526/")
}

func Test24(t *testing.T) {
	url2file(t, "https://sputniknews.kz/society/20190520/10197904/Sakhnany-syylauy-kerek-Belgili-galym-Qabatovtyn-soyleu-manerin-synga-aldy.html")
}

func Test25(t *testing.T) {
	url2file(t, "https://365info.kz/2019/05/za-schet-truboprovoda-v-knr-obemy-prodazh-tovarnogo-gaza-uvelichilis-pochti-na-tret")
}

func Test26(t *testing.T) {
	url2file(t, "https://kolesa.kz/content/news/bolee-30-tysyach-largusov-popali-pod-otzyv/")
}

func Test29(t *testing.T) {

	s := `<html><body> <strong><div>note:</div></strong> <p> a pargraph<a href='http://ya.ru'>with a link</a>in it. </p><ul><li>some <em>emphatic words</em> here.</li><li>more words.</li></ul></body></html>`
	s, _ = clean.Preprocess(s, true, nil)
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(s))
	maxsel := doc.Find("body")
	txt, link := clean.NodeDen(doc, maxsel, 0.3)
	pageText := txt
	log.Println(txt, link)

	max := -1.
	var mainNode *goquery.Selection
	var d func(*goquery.Selection)
	d = func(s *goquery.Selection) {
		nodes := s.Nodes
		if len(nodes) == 0 {
			return
		}

		for _, n := range nodes {
			nodesel := doc.FindNodes(n)
			txt, link := clean.NodeDen(doc, nodesel, 0.3)
			score := 0.
			if txt != 0 && pageText != 0 {
				score = 0.9*((txt-link)/txt) + 0.1*(txt/pageText)
			}

			if score > max {
				max = score
				mainNode = nodesel
			}
			log.Println(n.Data, txt, link, score)
		}
		s = s.Children()
		d(s)
	}
	d(maxsel.Children())
	log.Println(mainNode.Text())
}
