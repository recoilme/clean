package clean_test

import (
	"io/ioutil"
	"testing"

	"github.com/recoilme/clean"
	"github.com/stretchr/testify/assert"
)

func url2file(t *testing.T, u string) {
	s, err := clean.URI(u, true)
	assert.NoError(t, err)
	assert.NotEmpty(t, s)
	ioutil.WriteFile("out.htm", []byte(s), 0666)
}

func Test0(t *testing.T) {
	html := `
	<!DOCTYPE html>
<html>
<head>
    <title>Title</title>
</head>
<!--comment-->
<body>
  <script>alert('1')</script>
	<div class="main">
	<img src="https://habrastorage.org/getpro/habr/post_images/d83/22c/e16/d8322ce16ac2929781cb2f4062960382.jpg" align="right">Поско
		<div class="article"><br>
		
			<b>in b</>

      <div id="someid" class="articleHeadline">South Korea to Hold Artillery Drills on Island

      </div>
    <div class="articleBody">....The announcement came <br/> as...
      <a href="http://ya.ru">Bill Richardson</a>...
		</div>
		<ul>
    <li>
        <a href="/profile/login/facebook/" title="Facebook" class="d-fb js-anchor-ext-auth" data-ga-category="User behavior" data-ga-action="Sign up" data-ga-label="Signup Fb start">
            <span></span>
        </a>
		</li>
		</ul>

		<h3 style="text-align: center;">1.</h3></div><div id="js-block-45877015" class="">
    <p>
    <a name="image4019465" href="#image4019465" style="display:block; clear:both; margin-left:auto; margin-right:auto; text-align:center;"
       class="article-integrated-image-anchor js-gif-play">
        <span class="article-pic js-article-image "
              data-id="4019465"><img src="h1556196624.jpg"/></span>
    </a>
</p>
  </div>
</div>
</body>
</html>`
	s, err := clean.Clean(html, true, nil)
	assert.NoError(t, err)
	assert.NotEmpty(t, s)
}

func Test01(t *testing.T) {
	url2file(t, "https://vc.ru/services/67441-baza-dannyh-turisticheskogo-servisa-sletat-ru-s-pasportnymi-dannymi-klientov-vremenno-okazalas-v-otkrytom-dostupe")

}
func Test02(t *testing.T) {
	url2file(t, "https://stackoverflow.com/questions/15081119/any-way-to-use-html-parse-without-it-adding-nodes-to-make-a-well-formed-tree")
}

func Test3(t *testing.T) {
	url2file(t, "adme.ru")
}

func Test04(t *testing.T) {
	url2file(t, "https://www.adme.ru/svoboda-narodnoe-tvorchestvo/25-muzhchin-kotorye-nashli-sposob-sdelat-otcovstvo-chutochku-legche-2084015/")
}

func Test5(t *testing.T) {
	url2file(t, "https://t.me/s/recoilmeblog/245")
}

func Test6(t *testing.T) {
	url2file(t, "https://vc.ru/contest/67338-kpi")
}

func Test7(t *testing.T) {
	url2file(t, "https://habr.com")

}
func Test8(t *testing.T) {
	url2file(t, "https://www.producthunt.com/posts/emtech-brew")
}

func Test9(t *testing.T) {
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
