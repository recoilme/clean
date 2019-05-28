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

func Test00(t *testing.T) {
	html := `
	<!DOCTYPE html>
<html>
<head>
    <title>Title</title>
</head>
<!--comment-->
<body>
<!--comment bodu-->
  <script>alert('1')</script>
	<div class="main">
	<img src="https://habrastorage.org/getpro/habr/post_images/d83/22c/e16/d8322ce16ac2929781cb2f4062960382.jpg" align="right">Поско
		<div class="article"><br>
		
			

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
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatal(err)
	}

	var d func(*goquery.Selection)
	d = func(s *goquery.Selection) {
		for _, n := range s.Nodes {
			if n.Data == "script" || n.Data == "style" {
				s.Remove()
			}
			if n.Type != 1 && n.Type != 3 {
				s.Remove()
			}
			if clean.TrimBytes([]byte(s.Text())) == nil {
				//s.Remove()
			}
			ch := s.Children().Nodes
			if len(ch) == 1 && ch[0].Type == 3 {
				s.ReplaceWithNodes(ch[0])
			}

			log.Println(n.DataAtom)
		}
	}
	doc.Find("*").Children().Each(func(_ int, s *goquery.Selection) {
		d(s)
	})

	doc.Find("body").Each(func(_ int, s *goquery.Selection) {
		for _, n := range s.Nodes {
			ch := s.Children().Nodes
			if len(ch) == 1 && ch[0].Type == 3 {
				log.Println("alo")
				//	s.ReplaceWithNodes(ch[0])
			}
			log.Println("2.", n.DataAtom)
		}
	})
	s, _ = doc.Html()
	_ = s
	log.Println(s)
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

func Test18(t *testing.T) {
	html := `
	<div>
	 <div>
		<div>
		 <div>
			<div> <a href="http://feedproxy.google.com/auth/register"> <i>Впервые на cosmo.ru?</i><i>Зарегистрироваться</i></a>  <a href="http://feedproxy.google.com/auth/login"> <i>Вы уже зарегистрированы?</i><i>Войти</i></a> 
			</div>
			<div>
			 <div>Войти
			 </div>
			 <div><span>Быстрый вход через соц. сети</span>
			 </div><h4><strong>или</strong><span>войти, указав адрес электронной почты</span></h4><p><span>Напомнить пароль</span></p><p>Регистрируясь на сайте cosmo.ru, вы подтверждаете, что ознакомились и
							принимаете <a href="http://feedproxy.google.com/usage/%20"> правила пользовательского соглашения</a> , <a href="http://feedproxy.google.com/privacy_policy/"> Политику по обработке и защите персональных данных</a> , <a href="http://feedproxy.google.com/privacy_policy_bc/"> Политику кофиденциальности Бонусного клуба Cosmo</a> и <a href="http://feedproxy.google.com/cosmoshop/offer/"> соглашение об услугах Бонусного клуба Cosmo</a> </p><span>Поле заполнено не верно</span>
			</div>
			<div>
			 <div>Зарегистрироваться
			 </div>
			 <div><span>Быстрый вход через соц. сети</span>
			 </div><h4><strong>или</strong><span>зарегистрироваться, указав адрес электронной почты</span></h4><p>Я хочу получать новости от Cosmo.ru</p><p>Регистрируясь на сайте cosmo.ru, вы подтверждаете, что ознакомились и
							принимаете <a href="http://feedproxy.google.com/usage/%20"> правила пользовательского соглашения</a> , <a href="http://feedproxy.google.com/privacy_policy/"> Политику по обработке и защите персональных данных</a> , <a href="http://feedproxy.google.com/privacy_policy_bc/"> Политику кофиденциальности Бонусного клуба Cosmo</a> и <a href="http://feedproxy.google.com/cosmoshop/offer/"> соглашение об услугах Бонусного клуба Cosmo</a> </p><span>Поле заполнено не верно</span>
			</div>
			<div>
			 <div>
				<div>Восстановить пароль
				</div><span>Поле заполнено не верно</span>
			 </div>
			</div>
		 </div>
		</div>
	 </div>
	 <div>
		<div> <a href="http://feedproxy.google.com/~r/cosmo/fashion/~3/hsbD7Vq278c/"> </a>  <a href="http://feedproxy.google.com/~r/cosmo/fashion/~3/hsbD7Vq278c/"> </a> 
		</div>
	 </div>
	 <div>
		<div>
		 <div>
		 </div>
		 <div>
			<div><h4>Хочешь получать кэшбэк-рубли за все покупки в интернет-магазинах?</h4><p>Тогда скорее на Cosmoshop! Регистрируйся и приступай к покупкам</p><p>Чтобы покупать было ещё приятнее, мы дарим тебе</p><span>150</span><span>кэшбэк-рублей</span><span><span>+<br/>золотой статус</span>на 30 дней</span>
			 <div> <a href="http://feedproxy.google.com/~r/cosmo/fashion/~3/hsbD7Vq278c/"> Регистрация</a>  <a href="http://feedproxy.google.com/~r/cosmo/fashion/~3/hsbD7Vq278c/"> Без регистрации</a>  <a href="http://feedproxy.google.com/~r/cosmo/fashion/~3/hsbD7Vq278c/"> У меня уже есть аккаунт</a> 
			 </div>
			</div>
		 </div>
		</div>
		<div>
		 <div>
			<div><h3>Оферта по предоставлению услуг</h3>
			</div>
			<div><p>Добро пожаловать в Cosmoshop!</p><p>Перед использованием сервиса, ознакомьтесь с <a href="http://feedproxy.google.com/cosmoshop_oferta/"> пользовательской офертой.</a> </p><p>Нажимая продолжить, Вы соглашаетесь с условием оферты.</p>
			</div>
			<div><strong>Я прочитал и принимаю данное соглашение об услугах</strong>
			</div>
		 </div>
		</div>
		<div>
		 <div>Свернуть
		 </div>
		</div><header>
		<div>
		 <div>
			<div>
			 <div> <a href="http://feedproxy.google.com/"> COSMOPOLITAN</a> 
			 </div>
			</div>
			<div>
			 <div>
				<div>Поиск
				</div><span>Например:</span>Кристина Агилера
			 </div>
			</div>
		 </div>
		</div></header><main>
		<div><section>
		 <div>
			<div>
			 <div><span><time>19 мая 2019 18:00</time></span>
			 </div><section><h1>Спорная вещь: лаковые туфли. Носить или не носить?</h1><p>Инна Morata<span>,</span>автор блога La Micina<span>,</span>перечисляет 6 случаев<span>,</span>когда носить лаковые туфли точно не стоит<span>,</span>и еще 5 случаев<span>,</span>когда использовать их вполне себе разрешается.</p>
			 <div>
				<div><article>
				 <div>Cosmo Online
				 </div>
				 <div>Мода
				 </div>
				 <div>Не новости
				 </div>
				 <div><img src="https://images11.cosmopolitan.ru/upload/img_cache/c77/c77ad493c5327d93e5fee0f08d950d99_fitted_740x0.jpg"/><br/>
				 </div>
				 <div><p></p>
					<div><img src="https://images11.cosmopolitan.ru/upload/img_cache/eab/eab0f436651b6f3d925eb5842944a2a9_cropped_150x150.jpg"/><br/>
					 <div>
						<div>Inna Morata
						</div>
						<div>Блогер
						</div>
					 </div>
					</div><p></p><p>Затрагиваю сегодня злободневную тему. Заранее предвосхищаю споры и повышенный эмоциональный накал<span>,</span>но я же обещала… Итак<span>,</span>черная лакированная обувь — быть или не быть?</p><p>Не так давно я писала статью про особенности провинциального стиля и упомянула<span>,</span>что часто<span>,</span>путешествуя по стране<span>,</span>вижу на девушках черные лакированные туфли. Что же они мне такого сделали? На мой взгляд<span>,</span>это очень обязывающая обувь. Требует полной согласованности всего образа и совсем никак не справляется с ролью<span>«</span>той самой» универсальной пары<span>,</span>подходящей практически ко всему<span>(</span>это уж скорее будут белые лаконичные кроссовки<span>,</span>чем черные лаковые туфли).</p><p>Как точно не стоит комбинировать черные лакированные лодочки?</p>
					<div>
					 <div><span>Популярное</span>
					 </div><ul><li> <a href="http://feedproxy.google.com/fashion/trends/maksimum-komforta-samye-udobnye-lodochki-v-kotoryh-mozhno-dazhe-begat/"> <img src="https://images11.cosmopolitan.ru/upload/img_cache/45f/45ffff50646f58a06cf771c0290a3fad_cropped_200x133.jpg"/><br/><p>Максимум комфорта: самые удобные лодочки<span>,</span>в которых можно даже бегать</p></a> </li></ul>
					</div><ul><li>Абсолютно точно не стоит носить огромные блестящие<span>«</span>копытоподобные» экземплярчики с встроенной<span>(</span>или не встроенной) платформой. Нет<span>,</span>и все. Это — модная смерть. Сама подобные аж из Рима тащила лет 5 назад — правда<span>,</span>уже давно выбросила.</li></ul><p></p><ul><li><img src="https://images11.cosmopolitan.ru/upload/img_cache/d4e/d4eddef0bbd0c81435f7178ecd4714a0_fitted_358x700.jpg"/><br/></li><li><img src="https://images11.cosmopolitan.ru/upload/img_cache/490/4900311a3fa5e94d698d017a778de92d_fitted_358x700.jpg"/><br/></li></ul><p></p><ul><li>Совершенно не сочетаются черные закрытые туфли и летящие цветочные платьица: подобные образы выглядят странно.</li></ul><p></p>
					<div><img src="https://images11.cosmopolitan.ru/upload/img_cache/e62/e62a0dedd2e865068b95c550365f030b_fitted_740x700.jpg"/><br/>
					</div><p></p><ul><li>Суперфишка<span>«</span>Я добавила черный пояс<span>,</span>сумку или очки» тоже не работает — получается<span>,</span>как правило<span>,</span>очень банально и скучно.</li></ul><p></p>
					<div><img src="https://images11.cosmopolitan.ru/upload/img_cache/687/6878092a188bfae7882cae8e3b56b16f_fitted_740x700.jpg"/><br/>
					</div><p></p><ul><li>Кэжуал-образ с потертыми джинсами и твидовым жакетом нуждается скорее в нюдовой замшевой паре<span>,</span>чем в черных лакированных друзьях<span>(</span>кстати<span>,</span>это и более универсальный вариант<span>,</span>если что).</li></ul><p></p>
					<div><img src="https://images11.cosmopolitan.ru/upload/img_cache/14a/14a6997a232f41d01013d64a47440515_fitted_740x700.jpg"/><br/>
					</div><p></p><ul><li>«Я же в черной юбочке<span>,</span>почему нет??? — А сверху что? — Ну как что<span>,</span>мягкий трикотаж хорошего свежего цвета». Огорчу тебя<span>,</span>но и тут претенциозные черные лакированные туфли не будут кстати.</li></ul><p></p>
					<div><img src="https://images11.cosmopolitan.ru/upload/img_cache/afa/afafc14d1470000d6f01decff27621a0_fitted_740x700.jpg"/><br/>
					</div><p></p><p>Итак<span>,</span>когда же они в тему?</p><ul><li>У тебя готично-романтический образ — с большим количеством кружева<span>,</span>бархата и прочих аристократических нюансов.</li></ul><p></p>
					<div><img src="https://images11.cosmopolitan.ru/upload/img_cache/851/8519e22f6306fe42a13833e2a7635b6a_fitted_740x700.jpg"/><br/>
					</div><p></p><ul><li>У тебя есть дико стильные лакированные лоферы<span>,</span>дополняющие интересный образ с намеком на унисекс. Джинсам тут находиться тоже простительно.</li></ul><p></p>
					<div><img src="https://images11.cosmopolitan.ru/upload/img_cache/f2a/f2a18afa93d958fe285e4f76f0eb3d24_fitted_740x700.jpg"/><br/>
					</div><p></p><ul><li>На тебе драматическая вариация черно-белого сочетания<span>(</span>так и хочется добавить: и природная контрастность твоей внешности все это выдерживает).</li></ul><p></p>
					<div><img src="https://images11.cosmopolitan.ru/upload/img_cache/2b3/2b39ad64d824894acbe9e080f6b4cf7d_fitted_740x700.jpg"/><br/>
					</div><p></p><ul><li>Сегодня ты глэм-рок дива: на тебе как минимум пачка и косуха<span>,</span>дополненные основательным начесом<span>(</span>да простит мне публика мою нескромность).</li></ul><p></p>
					<div><img src="https://images11.cosmopolitan.ru/upload/img_cache/44c/44cb9177e5143e8293abb07a356c44ab_fitted_740x700.jpg"/><br/>
					</div><p></p><ul><li>Ты Моника Беллуччи — в этом случае в лакированных туфельках можно даже спать.</li></ul><p></p>
					<div><img src="https://images11.cosmopolitan.ru/upload/img_cache/1a4/1a41c2405f96238a0d03b1e2545c6564_fitted_740x700.jpg"/><br/>
					</div><p></p><p>Инна Morata<span>,</span>автор блога <a href="http://lamicina.com/"> La Micicna</a> </p>
					<div>
					 <div><span><strong>Интересуешься этой темой?</strong><br/>Получай последние тенденции моды на почту.</span>
					 </div>
					 <div><span>ОК</span>
					 </div><span>Я соглашаюсь с <a href="http://feedproxy.google.com/privacy_policy/"> правилами сайта</a> </span>
					 <div>
						<div>Спасибо.
						</div><span>Мы отправили на ваш email письмо с подтверждением.</span>
					 </div>
					</div>
				 </div></article>
				 <div>
					<div>
					 <div>Комментарии
					 </div>
					 <div>
						<div>
						 <div><p>Комментировать могут только авторизированные пользователи. Пожалуйста, <a href="http://feedproxy.google.com/auth/login"> войди</a> или <a href="http://feedproxy.google.com/auth/register"> зарегистрируйся</a> .</p>
							<div><span>Ещё быстро можно войти с помощью соц. сетей</span>
							</div>
						 </div><img src="http://feedproxy.google.com/~r/cosmo/fashion/~3/hsbD7Vq278c/"/><br/>
						 <div>Текст комментария
						 </div>Отменить
						</div>
					 </div>
					</div>
				 </div>
				</div>
			 </div></section>
			</div>
		 </div></section>
		 <div> <a href="http://feedproxy.google.com/fashion/how_to/tebe-ponravitsya-6-idey-cvetovyh-sochetaniy-na-primere-blogera-bler-edi/"> Показать больше</a> 
		 </div>
		</div></main><footer>
		<div>
		 <div>
			<div><ul><li> <a href="http://feedproxy.google.com/issue/cosmo/"> Cosmopolitan</a> </li><li> <a href="http://feedproxy.google.com/cosmopolitan-beauty/"> Cosmopolitan Beauty</a> </li><li> <a href="http://feedproxy.google.com/cosmopolitan-shopping/"> Cosmopolitan Shopping</a> </li></ul><ul><li> <a href="http://feedproxy.google.com/editors/blogs/"> Блог редакции</a> </li><li> <a href="http://feedproxy.google.com/editors/contests/"> Спецпроекты</a> </li><li> <a href="http://feedproxy.google.com/archive/"> Архив</a> </li><li> <a href="http://feedproxy.google.com/contacts/"> Контакты</a> </li></ul>
			</div>
			<div>
			 <div>Cosmo в соц сетях
			 </div>
			</div>
		 </div>
		 <div> <a href="http://feedproxy.google.com/issue/iyun-2019/"> <img src="https://images11.cosmopolitan.ru/upload/custom/68e/68eb2086fdc89d0ed645b854d926b91b.gif"/><br/></a> 
			<div><p>Новый номер<span>Июнь 2019</span></p>
			</div>
		 </div>
		 <div>
			<div>16+
			</div><p>© 2007 — 2019 ООО «Фэшн Пресс». При
									размещении материалов на Сайте
									Пользователь безвозмездно
									предоставляет ООО «Фэшн Пресс» неисключительные права на использование,
									воспроизведение,
									распространение, создание
									производных произведений, а также на демонстрацию материалов и доведение их до всеобщего
									сведения через сайт
									www.cosmo.ru и на официальных страницах этого сайта на www.facebook.com, vk.com,
									twitter.com, instagram.com. Партнер «Рамблера».</p>
		 </div>
		</div></footer>
		<div>
		 <div> <a href="http://www.liveinternet.ru/click"> <img src="http://counter.yadro.ru/logo"/><br/></a> 
		 </div>
		</div>
	 </div>
	</div>`
	s := clean.MainContent(html)
	_ = s
	ioutil.WriteFile("out.htm", []byte(s), 0666)
	//log.Println(s)
}

func Test19(t *testing.T) {
	html := `<div><div><a href="http://www.adv.rbc.ru"> 1234<span><span>Реклама на РБК<span>www.adv.rbc.ru</span></span></span></a> 
</div>42</div>`

	c, _ := clean.Preprocess(html, false, nil)
	//log.Println(c)
	clean.MainContent(c)
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

func Test27(t *testing.T) {
	b, err := ioutil.ReadFile("3.htm")
	assert.NoError(t, err)
	s, err := clean.Preprocess(string(b), false, nil)
	assert.NoError(t, err)

	/*
		s = `<div> <a href="https://habr.com/ru/users/bak/"> <img src="//habrastorage.org/r/w48/getpro/habr/avatars/f1a/dba/c3e/f1adbac3ea3be738241838375a0bdd70.png"/><br/>bak</a> <time>26 апреля 2014 в 23:43</time><ul><li> <a href="#comment_7536257"> </a> </li><li> <a href="#"> </a> </li><li></li><li> <a href="#comment_7536245"> </a> </li><li> <a href="#comment_7536257"> </a> </li></ul>
		  <div>+2
		  </div>
		</div>`
	*/
	var d func(*goquery.Selection)
	_ = d
	d = func(s *goquery.Selection) {
		if len(s.Children().Nodes) == 0 {
			if len(s.Nodes) > 0 {
				if s.Nodes[0].Data == "div" {
					txt := clean.TrimBytes([]byte(s.Text()))
					if txt == nil {
						s.Remove()
					} else {
						inner, _ := s.Html()
						s.ReplaceWithHtml(inner)
					}
				}
			}
			return
		}
		inner, _ := s.Html()
		s.ReplaceWithHtml(inner)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(s))
	/*
		for i := 0; i < 5; i++ {
			doc.Find("div").Children().Each(func(_ int, s *goquery.Selection) {
				d(s)
			})
		}
	*/

	s, _ = doc.Html()
	ioutil.WriteFile("1_1.htm", []byte(s), 0666)

	sel := doc.Find("body")

	iter := 0
	maxsel := sel
	for {
		maxsel = maxsel.Children()
		iter++
		if iter > 8 {
			break
		}
		max := 0
		sumd := 0
		for i, n := range maxsel.Nodes {
			_ = i

			d := clean.NodeDencity(n)
			sumd += d
			if iter > 0 {
				log.Println(i, n.Data, d)
			}
			if d >= max {
				max = d
				maxsel = doc.FindNodes(n)
			}
		}
		log.Println(max, maxsel.Nodes[0].Data)
		if iter >= 2 {
			//log.Println(maxsel.Text())
		}

	}
	res, err := maxsel.Html() //doc.Html()
	assert.NoError(t, err)
	//log.Println(res)
	//res = clean.MainContent(res)
	res, _ = clean.Preprocess(res, true, nil)
	ioutil.WriteFile("1_2.htm", []byte(res), 0666)
}

func Test28(t *testing.T) {

	s := `
	<div><div>
<div><div><p> 44&nbsp;комментария  44 комм. </p><div><div><p>Популярные</p><p>По порядку</p></div></div></div></div><div><div> Написать комментарий... <div></div><div></div></div></div><div><div><div><div><div><div> <a href="https://vc.ru/u/37037-aleksandr-kalin"> <img src="data:image/gif;base64,R0lGODlhAQABAJAAAP8AAAAAACH5BAUQAAAALAAAAAABAAEAAAICBAEAOw=="><p> Александр Калин </p></a> </div><div></div> <a href="https://vc.ru/services/67441-baza-dannyh-turisticheskogo-servisa-sletat-ru-s-pasportnymi-dannymi-klientov-vremenno-okazalas-v-otkrytom-dostupe"> <time>16&nbsp;мая&nbsp;в&nbsp;11:56</time></a> </div><div><div><div></div></div><div> 3 </div><div><div></div></div></div><div><div><p>дайте ссыль на гитхаб</p></div></div><div>Ответить</div><div><div></div></div><div></div></div></div><div><div><div><div><div><div><div> <a href="https://vc.ru/u/293-vyacheslav-anzhiganov"> <img src="data:image/gif;base64,R0lGODlhAQABAJAAAP8AAAAAACH5BAUQAAAALAAAAAABAAEAAAICBAEAOw=="><p> Vyacheslav Anzhiganov </p></a>  <a href="https://vc.ru/services/67441-baza-dannyh-turisticheskogo-servisa-sletat-ru-s-pasportnymi-dannymi-klientov-vremenno-okazalas-v-otkrytom-dostupe"> </a> </div><div></div> <a href="https://vc.ru/services/67441-baza-dannyh-turisticheskogo-servisa-sletat-ru-s-pasportnymi-dannymi-klientov-vremenno-okazalas-v-otkrytom-dostupe"> <time>16&nbsp;мая&nbsp;в&nbsp;15:50</time></a> </div><div><div><div></div></div><div> 10 </div><div><div></div></div></div><div><div><p>Держи, бро.</p></div><div> <a href="https://github.com/"> <img src="data:image/gif;base64,R0lGODlhAQABAJAAAP8AAAAAACH5BAUQAAAALAAAAAABAAEAAAICBAEAOw=="><div>Build software better, together</div><div>GitHub is where people build software. More than 36 million people use GitHub to discover, fork…</div><div></div></a> </div></div><div>Ответить</div><div><div></div></div><div></div></div></div><div><div><div><div><div><div><div> <a href="https://vc.ru/u/37037-aleksandr-kalin"> <img src="data:image/gif;base64,R0lGODlhAQABAJAAAP8AAAAAACH5BAUQAAAALAAAAAABAAEAAAICBAEAOw=="><p> Александр Калин </p></a>  <a href="https://vc.ru/services/67441-baza-dannyh-turisticheskogo-servisa-sletat-ru-s-pasportnymi-dannymi-klientov-vremenno-okazalas-v-otkrytom-dostupe"> </a> </div> <a href="https://vc.ru/services/67441-baza-dannyh-turisticheskogo-servisa-sletat-ru-s-pasportnymi-dannymi-klientov-vremenno-okazalas-v-otkrytom-dostupe"> <time>16&nbsp;мая&nbsp;в&nbsp;19:54</time></a> </div><div><div><div></div></div><div> 3 </div><div><div></div></div></div><div><div><p>ты чет долго, ждал такую ссыль в первые полчаса, начал, было, разочаровываться в вц</p></div></div><div>Ответить</div><div><div></div></div><div></div></div></div><div><div></div><div></div></div></div></div><div></div></div></div></div><div></div></div></div><div><div><div><div><div> <a href="https://vc.ru/u/224572-anton-lisin"> <img src="data:image/gif;base64,R0lGODlhAQABAJAAAP8AAAAAACH5BAUQAAAALAAAAAABAAEAAAICBAEAOw=="><p> Антон Лисин </p></a> </div><div></div> <a href="https://vc.ru/services/67441-baza-dannyh-turisticheskogo-servisa-sletat-ru-s-pasportnymi-dannymi-klientov-vremenno-okazalas-v-otkrytom-dostupe"> <time>16&nbsp;мая&nbsp;в&nbsp;11:53</time></a> </div><div><div><div></div></div><div> 4 </div><div><div></div></div></div><div><div><p>Ну и где уголовные дела?</p></div></div><div>Ответить</div><div><div></div></div><div></div></div></div><div><div><div><div><div><div><div> <a href="https://vc.ru/u/169061-ware-wow"> <img src="data:image/gif;base64,R0lGODlhAQABAJAAAP8AAAAAACH5BAUQAAAALAAAAAABAAEAAAICBAEAOw=="><p> Ware Wow </p></a>  <a href="https://vc.ru/services/67441-baza-dannyh-turisticheskogo-servisa-sletat-ru-s-pasportnymi-dannymi-klientov-vremenno-okazalas-v-otkrytom-dostupe"> </a> </div><div></div><div></div> <a href="https://vc.ru/services/67441-baza-dannyh-turisticheskogo-servisa-sletat-ru-s-pasportnymi-dannymi-klientov-vremenno-okazalas-v-otkrytom-dostupe"> <time>16&nbsp;мая&nbsp;в&nbsp;12:55</time></a> </div><div><div><div></div></div><div> 4 </div><div><div></div></div></div><div><div><p>+1 взломщиков нужно искать.<br>но им не до того...</p><p>  Сегодня главу отдела управления «К» ФСБ обвинили в получении взяток на $850 тыс.  <br></p></div></div><div>Ответить</div><div><div></div></div><div></div></div></div><div><div></div><div></div></div></div><div><div><div><div><div> <a href="https://vc.ru/u/60441-johnny-vorony"> <img src="data:image/gif;base64,R0lGODlhAQABAJAAAP8AAAAAACH5BAUQAAAALAAAAAABAAEAAAICBAEAOw=="><p> Johnny Vorony </p></a>  <a href="https://vc.ru/services/67441-baza-dannyh-turisticheskogo-servisa-sletat-ru-s-pasportnymi-dannymi-klientov-vremenno-okazalas-v-otkrytom-dostupe"> </a> </div> <a href="https://itunes.apple.com/ru/app/cukerberg-pozvonit-novosti/id920638420"> </a> <div></div> <a href="https://vc.ru/services/67441-baza-dannyh-turisticheskogo-servisa-sletat-ru-s-pasportnymi-dannymi-klientov-vremenno-okazalas-v-otkrytom-dostupe"> <time>16&nbsp;мая&nbsp;в&nbsp;12:33</time></a> </div><div><div><div></div></div><div> 0 </div><div><div></div></div></div><div><div><p>В следственных органах МВД России и Следственном комитете, зависит от состава преступления.</p></div></div><div>Ответить</div><div><div></div></div><div></div></div></div><div><div></div><div></div></div></div></div><div></div></div></div><div><div><div><div><div> <a href="https://vc.ru/u/140640-andrew-solovov"> <img src="data:image/gif;base64,R0lGODlhAQABAJAAAP8AAAAAACH5BAUQAAAALAAAAAABAAEAAAICBAEAOw=="><p> Andrew Solovov </p></a> </div> <a href="https://play.google.com/store/apps/details"> </a> <div></div> <a href="https://vc.ru/services/67441-baza-dannyh-turisticheskogo-servisa-sletat-ru-s-pasportnymi-dannymi-klientov-vremenno-okazalas-v-otkrytom-dostupe"> <time>16&nbsp;мая&nbsp;в&nbsp;11:48</time></a> </div><div><div><div></div></div><div> 0 </div><div><div></div></div></div><div><div><p>Вот случилось такое. Бывает.<br>Но какого хрена вы не шифруете такие данные в базе?!!</p></div></div><div>Ответить</div><div><div></div></div><div></div></div></div><div><div><div><div><div><div><div> <a href="https://vc.ru/u/66554-kamaz-uzbekov"> <img src="data:image/gif;base64,R0lGODlhAQABAJAAAP8AAAAAACH5BAUQAAAALAAAAAABAAEAAAICBAEAOw=="><p> Камаз Узбеков </p></a>  <a href="https://vc.ru/services/67441-baza-dannyh-turisticheskogo-servisa-sletat-ru-s-pasportnymi-dannymi-klientov-vremenno-okazalas-v-otkrytom-dostupe"> </a> </div><div></div> <a href="https://vc.ru/services/67441-baza-dannyh-turisticheskogo-servisa-sletat-ru-s-pasportnymi-dannymi-klientov-vremenno-okazalas-v-otkrytom-dostupe"> <time>16&nbsp;мая&nbsp;в&nbsp;11:49</time></a> </div><div><div><div></div></div><div> 0 </div><div><div></div></div></div><div><div><p>Как ты себе это представляешь?</p></div></div><div>Ответить</div><div><div></div></div><div></div></div></div><div><div><div><div><div><div><div> <a href="https://vc.ru/u/145413-kirill-gerasimenko"> <img src="data:image/gif;base64,R0lGODlhAQABAJAAAP8AAAAAACH5BAUQAAAALAAAAAABAAEAAAICBAEAOw=="><p> Кирилл Герасименко </p></a>  <a href="https://vc.ru/services/67441-baza-dannyh-turisticheskogo-servisa-sletat-ru-s-pasportnymi-dannymi-klientov-vremenno-okazalas-v-otkrytom-dostupe"> </a> </div> <a href="https://vc.ru/services/67441-baza-dannyh-turisticheskogo-servisa-sletat-ru-s-pasportnymi-dannymi-klientov-vremenno-okazalas-v-otkrytom-dostupe"> <time>16&nbsp;мая&nbsp;в&nbsp;11:51</time></a> </div><div><div><div></div></div><div> 16 </div><div><div></div></div></div><div><div><p>Берешь и шифруешь</p></div></div><div>Ответить</div><div><div></div></div><div></div></div></div><div><div><div><div><div><div><div> <a href="https://vc.ru/u/66554-kamaz-uzbekov"> <img src="data:image/gif;base64,R0lGODlhAQABAJAAAP8AAAAAACH5BAUQAAAALAAAAAABAAEAAAICBAEAOw=="><p> Камаз Узбеков </p></a>  <a href="https://vc.ru/services/67441-baza-dannyh-turisticheskogo-servisa-sletat-ru-s-pasportnymi-dannymi-klientov-vremenno-okazalas-v-otkrytom-dostupe"> </a> </div> <a href="https://vc.ru/services/67441-baza-dannyh-turisticheskogo-servisa-sletat-ru-s-pasportnymi-dannymi-klientov-vremenno-okazalas-v-otkrytom-dostupe"> <time>16&nbsp;мая&nbsp;в&nbsp;11:54</time></a> </div><div><div><div></div></div><div> 3 </div><div><div></div></div></div><div><div><p>Ключи будешь хранить в этой же базе? Или другую базу рядом положишь, тоже с открытым доступом?</p></div></div><div>Ответить</div><div><div></div></div><div></div></div></div><div><div><div><div><div><div><div> <a href="https://vc.ru/u/293-vyacheslav-anzhiganov"> <img src="data:image/gif;base64,R0lGODlhAQABAJAAAP8AAAAAACH5BAUQAAAALAAAAAABAAEAAAICBAEAOw=="><p> Vyacheslav Anzhiganov </p></a>  <a href="https://vc.ru/services/67441-baza-dannyh-turisticheskogo-servisa-sletat-ru-s-pasportnymi-dannymi-klientov-vremenno-okazalas-v-otkrytom-dostupe"> </a> </div> <a href="https://vc.ru/services/67441-baza-dannyh-turisticheskogo-servisa-sletat-ru-s-pasportnymi-dannymi-klientov-vremenno-okazalas-v-otkrytom-dostupe"> <time>16&nbsp;мая&nbsp;в&nbsp;12:16</time></a> </div><div><div><div></div></div><div> 1 </div><div><div></div></div></div><div><div><p>Зачем ключи в базе хранить?</p></div></div><div>Ответить</div><div><div></div></div><div></div></div></div><div><div><div><div><div><div><div> <a href="https://vc.ru/u/25730-ave-ego"> <img src="data:image/gif;base64,R0lGODlhAQABAJAAAP8AAAAAACH5BAUQAAAALAAAAAABAAEAAAICBAEAOw=="><p> ave ego </p></a>  <a href="https://vc.ru/services/67441-baza-dannyh-turisticheskogo-servisa-sletat-ru-s-pasportnymi-dannymi-klientov-vremenno-okazalas-v-otkrytom-dostupe"> </a> </div> <a href="https://vc.ru/services/67441-baza-dannyh-turisticheskogo-servisa-sletat-ru-s-pasportnymi-dannymi-klientov-vremenno-okazalas-v-otkrytom-dostupe"> <time>16&nbsp;мая&nbsp;в&nbsp;23:13</time></a> </div><div><div><div></div></div><div> 4 </div><div><div></div></div></div><div><div><p>да, выкладывайте их на гитхаб))</p></div></div><div>Ответить</div><div><div></div></div><div></div></div>
</div><div><div></div></div></div></div><div></div></div></div><div><div><div><div><div> <a href="https://vc.ru/u/33452-nik-lezhnevich"> <img src="data:image/gif;base64,R0lGODlhAQABAJAAAP8AAAAAACH5BAUQAAAALAAAAAABAAEAAAICBAEAOw=="><p> Ник Лежневич </p></a>  <a href="https://vc.ru/services/67441-baza-dannyh-turisticheskogo-servisa-sletat-ru-s-pasportnymi-dannymi-klientov-vremenno-okazalas-v-otkrytom-dostupe"> </a> </div> <a href="https://vc.ru/services/67441-baza-dannyh-turisticheskogo-servisa-sletat-ru-s-pasportnymi-dannymi-klientov-vremenno-okazalas-v-otkrytom-dostupe"> <time>16&nbsp;мая&nbsp;в&nbsp;12:12</time></a> </div><div><div><div></div></div><div> 1 </div><div><div></div></div></div><div><div><p>ключи хранятся там, где данные обрабатываются (то бишь какой-то программный код)</p></div></div><div>Ответить</div><div><div></div></div><div></div></div></div><div><div></div><div></div></div></div></div><div></div></div></div></div><div></div></div></div><div><div><div><div><div> <a href="https://vc.ru/u/140640-andrew-solovov"> <img src="data:image/gif;base64,R0lGODlhAQABAJAAAP8AAAAAACH5BAUQAAAALAAAAAABAAEAAAICBAEAOw=="><p> Andrew Solovov </p></a>  <a href="https://vc.ru/services/67441-baza-dannyh-turisticheskogo-servisa-sletat-ru-s-pasportnymi-dannymi-klientov-vremenno-okazalas-v-otkrytom-dostupe"> </a> </div> <a href="https://play.google.com/store/apps/details"> </a>  <a href="https://vc.ru/services/67441-baza-dannyh-turisticheskogo-servisa-sletat-ru-s-pasportnymi-dannymi-klientov-vremenno-okazalas-v-otkrytom-dostupe"> <time>16&nbsp;мая&nbsp;в&nbsp;11:55</time></a> </div><div><div><div></div></div><div> 7 </div><div><div></div></div></div><div><div><p>Я себе это не представляю, а делаю так - паспортные данные, еmail, телефон и другие, чувствительны данные хранятся в бд в зашифрованном виде.<br>При извлечение - дешифруются.<br>Согласен, что это не панацея. Но слив бд происходит в разы чаще, чем взлом всего кода и это хоть какая-то защита.</p></div></div><div>Ответить</div><div><div></div></div><div></div></div></div><div><div><div><div><div><div><div> <a href="https://vc.ru/u/185181-sergey-ikrin"> <img src="data:image/gif;base64,R0lGODlhAQABAJAAAP8AAAAAACH5BAUQAAAALAAAAAABAAEAAAICBAEAOw=="><p> Сергей Икрин </p></a>  <a href="https://vc.ru/services/67441-baza-dannyh-turisticheskogo-servisa-sletat-ru-s-pasportnymi-dannymi-klientov-vremenno-okazalas-v-otkrytom-dostupe"> </a> </div> <a href="https://vc.ru/services/67441-baza-dannyh-turisticheskogo-servisa-sletat-ru-s-pasportnymi-dannymi-klientov-vremenno-okazalas-v-otkrytom-dostupe"> <time>16&nbsp;мая&nbsp;в&nbsp;12:26</time></a> </div><div><div><div></div></div><div> –1 </div><div><div></div></div></div><div><div><p>Как вы потом будете делать поиск по шифрованным данным?</p></div></div><div>Ответить</div><div><div></div></div><div></div></div></div><div><div><div><div><div><div><div> <a href="https://vc.ru/u/140640-andrew-solovov"> <img src="data:image/gif;base64,R0lGODlhAQABAJAAAP8AAAAAACH5BAUQAAAALAAAAAABAAEAAAICBAEAOw=="><p> Andrew Solovov </p></a>  <a href="https://vc.ru/services/67441-baza-dannyh-turisticheskogo-servisa-sletat-ru-s-pasportnymi-dannymi-klientov-vremenno-okazalas-v-otkrytom-dostupe"> </a> </div> <a href="https://play.google.com/store/apps/details"> </a>  <a href="https://vc.ru/services/67441-baza-dannyh-turisticheskogo-servisa-sletat-ru-s-pasportnymi-dannymi-klientov-vremenno-okazalas-v-otkrytom-dostupe"> <time>16&nbsp;мая&nbsp;в&nbsp;12:52</time></a> </div><div><div><div></div></div><div> 3 </div><div><div></div></div></div><div><div><p>$needle = MyClass::Encrypt(паспорт);<br>select * from table where passport = $needle;</p><p>Магия какая то?</p></div></div><div>Ответить</div><div><div></div></div><div></div></div></div><div><div><div><div><div><div><div> <a href="https://vc.ru/u/185181-sergey-ikrin"> <img src="data:image/gif;base64,R0lGODlhAQABAJAAAP8AAAAAACH5BAUQAAAALAAAAAABAAEAAAICBAEAOw=="><p> Сергей Икрин </p></a>  <a href="https://vc.ru/services/67441-baza-dannyh-turisticheskogo-servisa-sletat-ru-s-pasportnymi-dannymi-klientov-vremenno-okazalas-v-otkrytom-dostupe"> </a> </div> <a href="https://vc.ru/services/67441-baza-dannyh-turisticheskogo-servisa-sletat-ru-s-pasportnymi-dannymi-klientov-vremenno-okazalas-v-otkrytom-dostupe"> <time>16&nbsp;мая&nbsp;в&nbsp;15:06</time></a> </div><div><div><div></div></div><div> 3 </div><div><div></div></div></div><div><div><p>Это по полному совпадению. Здесь вопросов нет.</p><p>Напишите магический вариант для поиска по подстроке. Например по названию населенного пункта из адреса прописки.</p></div></div><div>Ответить</div><div><div></div></div><div></div></div></div><div><div><div><div><div><div><div> <a href="https://vc.ru/u/140640-andrew-solovov"> <img src="data:image/gif;base64,R0lGODlhAQABAJAAAP8AAAAAACH5BAUQAAAALAAAAAABAAEAAAICBAEAOw=="><p> Andrew Solovov </p></a>  
<a href="https://vc.ru/services/67441-baza-dannyh-turisticheskogo-servisa-sletat-ru-s-pasportnymi-dannymi-klientov-vremenno-okazalas-v-otkrytom-dostupe"> </a> </div> <a href="https://vc.ru/services/67441-baza-dannyh-turisticheskogo-servisa-sletat-ru-s-pasportnymi-dannymi-klientov-vremenno-okazalas-v-otkrytom-dostupe"> <time>16&nbsp;мая&nbsp;в&nbsp;15:29</time></a> </div><div><div><div></div></div><div> 2 </div><div><div></div></div></div><div><div><p>  Напишите магический вариант для поиска по подстроке. Например по названию населенного пункта из адреса прописки<br>  </p><p>Пишу сайты, android приложения, микросервисы недорого — обращайтесь )</p></div></div><div>Ответить</div><div><div></div></div><div></div></div></div><div><div></div></div></div><div><div><div><div><div> <a href="https://vc.ru/u/205161-aleksis-vtoroy"> <img src="data:image/gif;base64,R0lGODlhAQABAJAAAP8AAAAAACH5BAUQAAAALAAAAAABAAEAAAICBAEAOw=="><p> Алексис Второй </p></a>  <a href="https://vc.ru/services/67441-baza-dannyh-turisticheskogo-servisa-sletat-ru-s-pasportnymi-dannymi-klientov-vremenno-okazalas-v-otkrytom-dostupe"> </a> </div><div></div> <a href="https://vc.ru/services/67441-baza-dannyh-turisticheskogo-servisa-sletat-ru-s-pasportnymi-dannymi-klientov-vremenno-okazalas-v-otkrytom-dostupe"> <time>16&nbsp;мая&nbsp;в&nbsp;15:11</time></a> </div><div><div><div></div></div><div> 0 </div><div><div></div></div></div><div><div><p>Нечуйствительные данные можна дублировать в открытом виде. В столбце Syti, нопример.</p></div></div><div>Ответить</div><div><div></div></div><div></div></div></div><div><div><div><div><div><div><div> <a href="https://vc.ru/u/185181-sergey-ikrin"> <img src="data:image/gif;base64,R0lGODlhAQABAJAAAP8AAAAAACH5BAUQAAAALAAAAAABAAEAAAICBAEAOw=="><p> Сергей Икрин </p></a>  <a href="https://vc.ru/services/67441-baza-dannyh-turisticheskogo-servisa-sletat-ru-s-pasportnymi-dannymi-klientov-vremenno-okazalas-v-otkrytom-dostupe"> </a> </div> <a href="https://vc.ru/services/67441-baza-dannyh-turisticheskogo-servisa-sletat-ru-s-pasportnymi-dannymi-klientov-vremenno-okazalas-v-otkrytom-dostupe"> <time>16&nbsp;мая&nbsp;в&nbsp;18:09</time></a> </div><div><div><div></div></div><div> 2 </div><div><div></div></div></div><div><div><p>Вы побили рекорд количества ошибок в одном коротком комментарии.</p></div></div><div>Ответить</div><div><div></div></div><div></div></div></div><div><div></div></div></div></div></div></div><div><div><div><div><div> <a href="https://vc.ru/u/308-mike-kosulin"> <img src="data:image/gif;base64,R0lGODlhAQABAJAAAP8AAAAAACH5BAUQAAAALAAAAAABAAEAAAICBAEAOw=="><p> Mike Kosulin </p></a>  <a href="https://vc.ru/services/67441-baza-dannyh-turisticheskogo-servisa-sletat-ru-s-pasportnymi-dannymi-klientov-vremenno-okazalas-v-otkrytom-dostupe"> </a> </div> <a href="https://vc.ru/services/67441-baza-dannyh-turisticheskogo-servisa-sletat-ru-s-pasportnymi-dannymi-klientov-vremenno-okazalas-v-otkrytom-dostupe"> <time>17&nbsp;мая&nbsp;в&nbsp;00:41</time></a> </div><div><div><div></div></div><div> 0 </div><div><div></div></div></div><div><div><p>Адрес — так себе пример.<br>Адреса в едином виде, отсекается ненужная информация и все.</p><p>Сам справочник можно вообще не в бд держать. Плюс не использовать стандартные коды. Да и вообще в микросервис выделить.</p></div></div><div>Ответить</div><div><div></div></div><div></div></div></div><div><div></div></div></div></div></div></div><div><div><div><div><div> <a href="https://vc.ru/u/174309-sergei-timofeyev"> <img src="data:image/gif;base64,R0lGODlhAQABAJAAAP8AAAAAACH5BAUQAAAALAAAAAABAAEAAAICBAEAOw=="><p> Sergei Timofeyev </p></a>  <a href="https://vc.ru/services/67441-baza-dannyh-turisticheskogo-servisa-sletat-ru-s-pasportnymi-dannymi-klientov-vremenno-okazalas-v-otkrytom-dostupe"> </a> </div> <a href="https://vc.ru/services/67441-baza-dannyh-turisticheskogo-servisa-sletat-ru-s-pasportnymi-dannymi-klientov-vremenno-okazalas-v-otkrytom-dostupe"> <time>16&nbsp;мая&nbsp;в&nbsp;16:03</time></a> </div><div><div><div></div></div><div> 0 </div><div><div></div></div></div><div><div><p>' drop database;</p></div></div><div>Ответить</div><div><div></div></div><div></div></div></div><div><div></div></div></div></div><div></div></div></div><div><div><div><div><div> <a href="https://vc.ru/u/28026-anton-vlasov"> <img src="data:image/gif;base64,R0lGODlhAQABAJAAAP8AAAAAACH5BAUQAAAALAAAAAABAAEAAAICBAEAOw=="><p> Anton Vlasov </p></a>  <a href="https://vc.ru/services/67441-baza-dannyh-turisticheskogo-servisa-sletat-ru-s-pasportnymi-dannymi-klientov-vremenno-okazalas-v-otkrytom-dostupe"> </a> </div> <a href="https://vc.ru/services/67441-baza-dannyh-turisticheskogo-servisa-sletat-ru-s-pasportnymi-dannymi-klientov-vremenno-okazalas-v-otkrytom-dostupe"> <time>16&nbsp;мая&nbsp;в&nbsp;18:28</time></a> </div><div><div><div></div></div><div> 0 </div><div><div></div></div></div><div><div><p>вам нужен поиск регекспами по паспортным данным клиентов? да ну нафиг</p></div></div><div>Ответить</div><div><div></div></div><div></div></div></div><div><div></div><div></div></div></div></div><div></div></div></div><div><div><div><div><div> <a href="https://vc.ru/u/142710-egor-homakov"> <img src="data:image/gif;base64,R0lGODlhAQABAJAAAP8AAAAAACH5BAUQAAAALAAAAAABAAEAAAICBAEAOw=="><p> Egor Homakov </p></a>  <a href="https://vc.ru/services/67441-baza-dannyh-turisticheskogo-servisa-sletat-ru-s-pasportnymi-dannymi-klientov-vremenno-okazalas-v-otkrytom-dostupe"> </a> </div> <a href="https://vc.ru/services/67441-baza-dannyh-turisticheskogo-servisa-sletat-ru-s-pasportnymi-dannymi-klientov-vremenno-okazalas-v-otkrytom-dostupe"> <time>16&nbsp;мая&nbsp;в&nbsp;12:44</time></a> </div><div><div><div></div></div><div> 2 </div><div><div></div></div></div><div><div><p>Поиск по полям, всякие выборки перестают работать. Повезет если пароли хешируют, а чтобы данные по которым воронки делают шифровали, это редкость.</p></div></div><div>Ответить</div><div><div></div></div><div></div></div></div><div><div></div><div></div></div></div></div><div></div></div></div><div><div><div><div><div> <a href="https://vc.ru/u/139412-danila-vereshchakov"> <img src="data:image/gif;base64,R0lGODlhAQABAJAAAP8AAAAAACH5BAUQAAAALAAAAAABAAEAAAICBAEAOw=="><p> Данила Верещаков </p></a>  <a href="https://vc.ru/services/67441-baza-dannyh-turisticheskogo-servisa-sletat-ru-s-pasportnymi-dannymi-klientov-vremenno-okazalas-v-otkrytom-dostupe"> </a> </div> <a href="https://vc.ru/services/67441-baza-dannyh-turisticheskogo-servisa-sletat-ru-s-pasportnymi-dannymi-klientov-vremenno-okazalas-v-otkrytom-dostupe"> <time>16&nbsp;мая&nbsp;в&nbsp;12:17</time></a> </div><div><div><div></div></div><div> 4 </div><div><div></div></div></div><div><div><p>"содержала логины и пароли нескольких сотен подключенных к системе турагентств"<br>Ну уж пароли хранить в открытом виде - это совсем зашквар!</p></div></div><div>Ответить</div><div><div></div></div><div></div></div></div><div><div><div><div><div><div><div> <a href="https://vc.ru/u/103360-ashot-oganesyan"> <img src="data:image/gif;base64,R0lGODlhAQABAJAAAP8AAAAAACH5BAUQAAAALAAAAAABAAEAAAICBAEAOw=="><p> Ashot Oganesyan </p></a>  <a href="https://vc.ru/services/67441-baza-dannyh-turisticheskogo-servisa-sletat-ru-s-pasportnymi-dannymi-klientov-vremenno-okazalas-v-otkrytom-dostupe"> </a> </div> <a href="https://vc.ru/services/67441-baza-dannyh-turisticheskogo-servisa-sletat-ru-s-pasportnymi-dannymi-klientov-vremenno-okazalas-v-otkrytom-dostupe"> <time>16&nbsp;мая&nbsp;в&nbsp;15:18</time></a> </div><div><div><div></div></div><div> 0 </div><div><div></div></div></div><div><div></div><div><div><div></div></div></div></div><div>Ответить</div><div><div></div></div><div></div></div></div><div><div><div><div><div><div><div> <a href="https://vc.ru/u/174309-sergei-timofeyev"> <img src="data:image/gif;base64,R0lGODlhAQABAJAAAP8AAAAAACH5BAUQAAAALAAAAAABAAEAAAICBAEAOw=="><p> Sergei Timofeyev </p></a>  <a href="https://vc.ru/services/67441-baza-dannyh-turisticheskogo-servisa-sletat-ru-s-pasportnymi-dannymi-klientov-vremenno-okazalas-v-otkrytom-dostupe"> </a> </div> <a href="https://vc.ru/services/67441-baza-dannyh-turisticheskogo-servisa-sletat-ru-s-pasportnymi-dannymi-klientov-vremenno-okazalas-v-otkrytom-dostupe"> <time>16&nbsp;мая&nbsp;в&nbsp;16:07</time></a> </div><div><div><div></div></div><div> 2 </div><div><div></div></div></div><div><div><p>и без ssl :)</p></div></div><div>Ответить</div><div><div></div></div><div></div></div></div><div><div><div><div><div><div><div> <a href="https://vc.ru/u/308-mike-kosulin"> <img src="data:image/gif;base64,R0lGODlhAQABAJAAAP8AAAAAACH5BAUQAAAALAAAAAABAAEAAAICBAEAOw=="><p> Mike Kosulin </p></a>  <a href="https://vc.ru/services/67441-baza-dannyh-turisticheskogo-servisa-sletat-ru-s-pasportnymi-dannymi-klientov-vremenno-okazalas-v-otkrytom-dostupe"> </a> </div> <a href="https://vc.ru/services/67441-baza-dannyh-turisticheskogo-servisa-sletat-ru-s-pasportnymi-dannymi-klientov-vremenno-okazalas-v-otkrytom-dostupe"> <time>17&nbsp;мая&nbsp;в&nbsp;00:42</time></a> </div><div><div><div></div></div><div> 0 </div><div><div></div></div></div><div><div><p>Да ну, серты ещё обновлять.</p></div></div><div>Ответить</div><div><div></div></div><div></div></div></div><div><div><div><div><div><div><div> <a href="https://vc.ru/u/174309-sergei-timofeyev"> <img src="data:image/gif;base64,R0lGODlhAQABAJAAAP8AAAAAACH5BAUQAAAALAAAAAABAAEAAAICBAEAOw=="><p> Sergei Timofeyev </p></a>  <a href="https://vc.ru/services/67441-baza-dannyh-turisticheskogo-servisa-sletat-ru-s-pasportnymi-dannymi-klientov-vremenno-okazalas-v-otkrytom-dostupe"> </a> </div> <a href="https://vc.ru/services/67441-baza-dannyh-turisticheskogo-servisa-sletat-ru-s-pasportnymi-dannymi-klientov-vremenno-okazalas-v-otkrytom-dostupe"> 
<time>17&nbsp;мая&nbsp;в&nbsp;00:48</time></a> </div><div><div><div></div></div><div> 0 </div><div><div></div></div></div><div><div><p>У меня само...</p></div></div><div>Ответить</div><div><div></div></div><div></div></div></div><div><div></div></div></div></div></div></div></div><div></div></div></div></div><div></div></div></div></div><div></div></div></div><div><div><div><div><div> <a href="https://vc.ru/u/293-vyacheslav-anzhiganov"> <img src="data:image/gif;base64,R0lGODlhAQABAJAAAP8AAAAAACH5BAUQAAAALAAAAAABAAEAAAICBAEAOw=="><p> Vyacheslav Anzhiganov </p></a>  <a href="https://vc.ru/services/67441-baza-dannyh-turisticheskogo-servisa-sletat-ru-s-pasportnymi-dannymi-klientov-vremenno-okazalas-v-otkrytom-dostupe"> </a> </div><div></div> <a href="https://vc.ru/services/67441-baza-dannyh-turisticheskogo-servisa-sletat-ru-s-pasportnymi-dannymi-klientov-vremenno-okazalas-v-otkrytom-dostupe"> <time>16&nbsp;мая&nbsp;в&nbsp;12:20</time></a> </div><div><div><div></div></div><div> 0 </div><div><div></div></div></div><div><div><p>Направления для изучения на примере Mysql:<br>AES_ENCRYPT() - шифрование AES<br>AES_DECRYPT() - расшифровка AES<br>COMPRESS() - возвращение результата в бинарном виде<br>DES_ENCRYPT() - шифрование DES<br>DES_DECRYPT() - дешифрование DES<br>ENCODE() - шифрование строки поверхностным паролем (на выходе<br>получается шифрованное слово первоначальной "plaintext" длины<br>DECODE() - расшифровка текста, обработанного функцией ENCODE()<br>ENCRYPT() - шифрование с помощью Unix’ового системного вызова crypt<br>MD5() - подсчет MD-5 суммы<br>SHA1(), SHA() - подсчет SHA-1 (160-бит)</p><p>Конечно, нагрузка выростет, но это лучше, чем голой жопой в интернет смотреть.</p></div></div><div>Ответить</div><div><div></div></div><div></div></div></div><div><div><div><div><div><div><div> <a href="https://vc.ru/u/174309-sergei-timofeyev"> <img src="data:image/gif;base64,R0lGODlhAQABAJAAAP8AAAAAACH5BAUQAAAALAAAAAABAAEAAAICBAEAOw=="><p> Sergei Timofeyev </p></a>  <a href="https://vc.ru/services/67441-baza-dannyh-turisticheskogo-servisa-sletat-ru-s-pasportnymi-dannymi-klientov-vremenno-okazalas-v-otkrytom-dostupe"> </a> </div> <a href="https://vc.ru/services/67441-baza-dannyh-turisticheskogo-servisa-sletat-ru-s-pasportnymi-dannymi-klientov-vremenno-okazalas-v-otkrytom-dostupe"> <time>16&nbsp;мая&nbsp;в&nbsp;16:08</time></a> </div><div><div><div></div></div><div> 3 </div><div><div></div></div></div><div><div><p>База в интернет не должна смотреть ни при каких обстоятельствах.</p></div></div><div>Ответить</div><div><div></div></div><div></div></div></div><div><div><div><div><div><div><div> <a href="https://vc.ru/u/28026-anton-vlasov"> <img src="data:image/gif;base64,R0lGODlhAQABAJAAAP8AAAAAACH5BAUQAAAALAAAAAABAAEAAAICBAEAOw=="><p> Anton Vlasov </p></a>  <a href="https://vc.ru/services/67441-baza-dannyh-turisticheskogo-servisa-sletat-ru-s-pasportnymi-dannymi-klientov-vremenno-okazalas-v-otkrytom-dostupe"> </a> </div> <a href="https://vc.ru/services/67441-baza-dannyh-turisticheskogo-servisa-sletat-ru-s-pasportnymi-dannymi-klientov-vremenno-okazalas-v-otkrytom-dostupe"> <time>16&nbsp;мая&nbsp;в&nbsp;18:29</time></a> </div><div><div><div></div></div><div> 0 </div><div><div></div></div></div><div><div><p>ага, такие умники обычно рядом с неторчащей базой кладут скриптец вроде phpmyadmin :)</p></div></div><div>Ответить</div><div><div></div></div><div></div></div></div><div><div><div><div><div><div><div> <a href="https://vc.ru/u/174309-sergei-timofeyev"> <img src="data:image/gif;base64,R0lGODlhAQABAJAAAP8AAAAAACH5BAUQAAAALAAAAAABAAEAAAICBAEAOw=="><p> Sergei Timofeyev </p></a>  <a href="https://vc.ru/services/67441-baza-dannyh-turisticheskogo-servisa-sletat-ru-s-pasportnymi-dannymi-klientov-vremenno-okazalas-v-otkrytom-dostupe"> </a> </div> <a href="https://vc.ru/services/67441-baza-dannyh-turisticheskogo-servisa-sletat-ru-s-pasportnymi-dannymi-klientov-vremenno-okazalas-v-otkrytom-dostupe"> <time>16&nbsp;мая&nbsp;в&nbsp;19:47</time></a> </div><div><div><div></div></div><div> 1 </div><div><div></div></div></div><div><div><p>c root без пароля. :) Вообще последние пакеты MySQL без пароля очень сильно удивили. Весь мир за безопасность, а тут они решили что-то изменить...</p></div></div><div>Ответить</div><div><div></div></div><div></div></div></div><div><div></div></div></div></div><div></div></div></div><div><div><div><div><div> <a href="https://vc.ru/u/293-vyacheslav-anzhiganov"> <img src="data:image/gif;base64,R0lGODlhAQABAJAAAP8AAAAAACH5BAUQAAAALAAAAAABAAEAAAICBAEAOw=="><p> Vyacheslav Anzhiganov </p></a>  <a href="https://vc.ru/services/67441-baza-dannyh-turisticheskogo-servisa-sletat-ru-s-pasportnymi-dannymi-klientov-vremenno-okazalas-v-otkrytom-dostupe"> </a> </div> <a href="https://vc.ru/services/67441-baza-dannyh-turisticheskogo-servisa-sletat-ru-s-pasportnymi-dannymi-klientov-vremenno-okazalas-v-otkrytom-dostupe"> <time>16&nbsp;мая&nbsp;в&nbsp;16:47</time></a> </div><div><div><div></div></div><div> 0 </div><div><div></div></div></div><div><div><p>Да, а чём статья которую мы обсуждаем? =)</p></div></div><div>Ответить</div><div><div></div></div><div></div></div></div><div><div><div><div><div><div><div> <a href="https://vc.ru/u/174309-sergei-timofeyev"> <img src="data:image/gif;base64,R0lGODlhAQABAJAAAP8AAAAAACH5BAUQAAAALAAAAAABAAEAAAICBAEAOw=="><p> Sergei Timofeyev </p></a>  <a href="https://vc.ru/services/67441-baza-dannyh-turisticheskogo-servisa-sletat-ru-s-pasportnymi-dannymi-klientov-vremenno-okazalas-v-otkrytom-dostupe"> </a> </div> <a href="https://vc.ru/services/67441-baza-dannyh-turisticheskogo-servisa-sletat-ru-s-pasportnymi-dannymi-klientov-vremenno-okazalas-v-otkrytom-dostupe"> <time>16&nbsp;мая&nbsp;в&nbsp;17:01</time></a> </div><div><div><div></div></div><div> 0 </div><div><div></div></div></div><div><div><p>Про дурость некоторых индивидов. :)</p></div></div><div>Ответить</div><div><div></div></div><div></div></div></div><div><div></div></div></div></div><div></div></div></div></div><div></div></div></div><div><div><div><div><div> <a href="https://vc.ru/u/2327-alexander-matveev"> <img src="data:image/gif;base64,R0lGODlhAQABAJAAAP8AAAAAACH5BAUQAAAALAAAAAABAAEAAAICBAEAOw=="><p> Alexander Matveev </p></a>  <a href="https://vc.ru/services/67441-baza-dannyh-turisticheskogo-servisa-sletat-ru-s-pasportnymi-dannymi-klientov-vremenno-okazalas-v-otkrytom-dostupe"> </a> </div> <a href="https://vc.ru/services/67441-baza-dannyh-turisticheskogo-servisa-sletat-ru-s-pasportnymi-dannymi-klientov-vremenno-okazalas-v-otkrytom-dostupe"> <time>16&nbsp;мая&nbsp;в&nbsp;12:47</time></a> </div><div><div><div></div></div><div> –1 </div><div><div></div></div></div><div><div><p>(нет)</p></div></div><div>Ответить</div><div><div></div></div><div></div></div></div><div><div></div><div></div></div></div></div><div></div></div></div></div><div></div></div></div></div><div></div></div></div><div><div><div><div><div> <a href="https://vc.ru/u/248622-dmitriy-fisher"> <img src="data:image/gif;base64,R0lGODlhAQABAJAAAP8AAAAAACH5BAUQAAAALAAAAAABAAEAAAICBAEAOw=="><p> Дмитрий Фишер </p></a> </div> <a href="https://itunes.apple.com/ru/app/cukerberg-pozvonit-novosti/id920638420"> </a> <div></div> <a href="https://vc.ru/services/67441-baza-dannyh-turisticheskogo-servisa-sletat-ru-s-pasportnymi-dannymi-klientov-vremenno-okazalas-v-otkrytom-dostupe"> <time>16&nbsp;мая&nbsp;в&nbsp;11:35</time></a> </div><div><div><div></div></div><div> 2 </div><div><div></div></div></div><div><div><p>Сейчас эти ошибки с базами данных всплывают практически везде, о какой конфиденциальности вообще может быть речь.</p></div></div><div>Ответить</div><div><div></div></div><div></div></div></div><div><div></div><div></div></div></div><div><div><div><div><div> <a href="https://vc.ru/u/25730-ave-ego"> <img src="data:image/gif;base64,R0lGODlhAQABAJAAAP8AAAAAACH5BAUQAAAALAAAAAABAAEAAAICBAEAOw=="><p> ave ego </p></a> </div><div></div> <a href="https://vc.ru/services/67441-baza-dannyh-turisticheskogo-servisa-sletat-ru-s-pasportnymi-dannymi-klientov-vremenno-okazalas-v-otkrytom-dostupe"> <time>16&nbsp;мая&nbsp;в&nbsp;23:14</time></a> </div><div><div><div></div></div><div> 1 </div><div><div></div></div></div><div><div><p>ркн не найдет нарушений, тк данные обрабатываются в рамках договоров и законодательных актов) <a href="https://rkn.gov.ru/news/rsoc/news67040.htm"> https://rkn.gov.ru/news/rsoc/news67040.htm</a> </p></div></div><div>Ответить</div><div><div></div></div><div></div></div></div><div><div></div><div></div></div></div><div><div><div><div><div> <a href="https://vc.ru/u/294386-elena-loginova"> <img src="data:image/gif;base64,R0lGODlhAQABAJAAAP8AAAAAACH5BAUQAAAALAAAAAABAAEAAAICBAEAOw=="><p> Елена Логинова </p></a> </div><div></div> <a href="https://vc.ru/services/67441-baza-dannyh-turisticheskogo-servisa-sletat-ru-s-pasportnymi-dannymi-klientov-vremenno-okazalas-v-otkrytom-dostupe"> <time>16&nbsp;мая&nbsp;в&nbsp;12:52</time></a> </div><div><div><div></div></div><div> –5 </div><div><div></div></div></div><div><div><p>&lt;a href=" <a href="https://homeandhoby.blogspot.com/%22"> https://homeandhoby.blogspot.com/"</a> ; target="_blank"&gt;Дом и хобби&lt;/a&gt;</p></div></div><div>Ответить</div><div><div></div></div><div></div></div></div><div><div><div><div><div><div><div> <a href="https://vc.ru/u/85254-pavel-natalia-selivanovi"> <img src="data:image/gif;base64,R0lGODlhAQABAJAAAP8AAAAAACH5BAUQAAAALAAAAAABAAEAAAICBAEAOw=="><p> Pavel Natalia Selivanovi </p></a>  <a href="https://vc.ru/services/67441-baza-dannyh-turisticheskogo-servisa-sletat-ru-s-pasportnymi-dannymi-klientov-vremenno-okazalas-v-otkrytom-dostupe"> </a> </div><div></div> <a href="https://vc.ru/services/67441-baza-dannyh-turisticheskogo-servisa-sletat-ru-s-pasportnymi-dannymi-klientov-vremenno-okazalas-v-otkrytom-dostupe"> <time>16&nbsp;мая&nbsp;в&nbsp;14:40</time></a> </div><div><div><div></div></div><div> 0 </div><div><div></div></div></div><div><div><p>&lt;font size="9"&gt;&lt;b&gt;ВСЕМ ПРИВЕТ!&lt;/b&gt;&lt;/font&gt;</p></div></div><div>Ответить</div><div><div></div></div><div></div></div></div><div><div><div><div><div><div><div> <a href="https://vc.ru/u/174309-sergei-timofeyev"> <img src="data:image/gif;base64,R0lGODlhAQABAJAAAP8AAAAAACH5BAUQAAAALAAAAAABAAEAAAICBAEAOw=="><p> Sergei Timofeyev </p></a>  <a href="https://vc.ru/services/67441-baza-dannyh-turisticheskogo-servisa-sletat-ru-s-pasportnymi-dannymi-klientov-vremenno-okazalas-v-otkrytom-dostupe"> </a> </div> <a href="https://vc.ru/services/67441-baza-dannyh-turisticheskogo-servisa-sletat-ru-s-pasportnymi-dannymi-klientov-vremenno-okazalas-v-otkrytom-dostupe"> <time>17&nbsp;мая&nbsp;в&nbsp;00:50</time></a> </div><div><div><div></div></div><div> 0 </div><div><div></div></div></div><div><div><p>strong же</p></div></div><div>Ответить</div><div><div></div></div><div></div></div></div><div><div></div><div></div></div></div></div><div></div></div></div></div><div></div></div></div><div><div><div><div><div> <a href="https://vc.ru/u/66991-alex-lutnev"> <img src="data:image/gif;base64,R0lGODlhAQABAJAAAP8AAAAAACH5BAUQAAAALAAAAAABAAEAAAICBAEAOw=="><p> Alex Lutnev </p></a> </div><div></div> <a href="https://vc.ru/services/67441-baza-dannyh-turisticheskogo-servisa-sletat-ru-s-pasportnymi-dannymi-klientov-vremenno-okazalas-v-otkrytom-dostupe"> <time>16&nbsp;мая&nbsp;в&nbsp;12:30</time></a> </div><div><div><div></div></div><div> 0 </div><div><div></div></div></div><div><div><p>Вообще эту новость Коммерсу стоило выпустить хотя бы ради заголовка "Хакеров пригласили в тур". Ору тут потихонечку, держу вас в курсе</p></div></div><div>Ответить</div><div><div></div></div><div></div></div></div><div><div></div><div></div></div></div><div><div><div><div><div> <a href="https://vc.ru/u/296065-nikolya"> <img src="data:image/gif;base64,R0lGODlhAQABAJAAAP8AAAAAACH5BAUQAAAALAAAAAABAAEAAAICBAEAOw=="><p> Николя </p></a> </div><div></div> <a href="https://vc.ru/services/67441-baza-dannyh-turisticheskogo-servisa-sletat-ru-s-pasportnymi-dannymi-klientov-vremenno-okazalas-v-otkrytom-dostupe"> <time>16&nbsp;мая&nbsp;в&nbsp;15:54</time></a> </div><div><div><div></div></div><div> 0 </div><div><div></div></div></div><div><div><p>Легкую юбочку Строгово закона о защите Персональных Данных Россиян сдуло бризом, и она слетела, временно обнажив розовенькую, но уже давно не девственную, потрепанную попку</p></div></div><div>Ответить</div><div><div></div></div><div></div></div></div><div><div></div><div></div></div></div><div><div><div><div><div> <a href="https://vc.ru/u/86341-frederic-rossif"> <img src="data:image/gif;base64,R0lGODlhAQABAJAAAP8AAAAAACH5BAUQAAAALAAAAAABAAEAAAICBAEAOw=="><p> Frédéric Rossif </p></a> </div><div></div> <a href="https://vc.ru/services/67441-baza-dannyh-turisticheskogo-servisa-sletat-ru-s-pasportnymi-dannymi-klientov-vremenno-okazalas-v-otkrytom-dostupe"> <time>вчера&nbsp;в&nbsp;15:04</time></a> </div><div><div><div></div></div><div> 0 </div><div><div></div></div></div><div><div><p>Виртуальные комнаты данных с шифрованием уже отменили ?</p></div></div><div>Ответить</div><div><div></div></div><div></div></div></div><div><div></div><div></div></div></div></div><div><div><div></div><div><div></div><div><div>Отправить</div><div></div><div></div>        <div>Отменить</div></div></div></div><div> Написать комментарий... <div></div><div></div></div></div>
	
  wweweqw</div> </div>
`

	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(s))
	for i, n := range doc.Nodes {
		_ = i

		d := clean.NodeDencity(n)
		log.Println(i, d)
	}

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

/*
-- th 0.3
div 0.8331918533684844 153043 28235 0.815509366648589 0.9039218002480657
2019/05/28 17:39:52 ----
2019/05/28 17:39:52 div 0.8237873729292887 141698 25434 0.8205055822947395 0.8369145354674857
2019/05/28 17:39:52 newnode
2019/05/28 17:39:52 div 0.7855557087675903 113140 20944 0.8148842142478345 0.6682416868466127
2019/05/28 17:39:52 newnode
2019/05/28 17:39:52 div 0.7779414856368929 105960 19500 0.8159682899207248 0.6258342685015652
2019/05/28 17:39:52 newnode
2019/05/28 17:39:52 div 0.7704506775947658 98779 18056 0.8172081110357464 0.5834209438308429
2019/05/28 17:39:52 newnode
2019/05/28 17:39:52 div 0.7635270195138256 91747 16612 0.8189368589708655 0.5418876616856654
2019/05/28 17:39:52 newnode
2019/05/28 17:39:52 div 0.7499965657766894 72883 12399 0.8298780236817913 0.4304707341562814
2019/05/28 17:39:52 newnode
2019/05/28 17:39:52 div 0.7445435922683248 67798 11487 0.8305702233104222 0.400437068099935
th 0
2019/05/28 17:42:42 div 0.8332634223676355 153043 28235 0.815509366648589 0.904279645243821
2019/05/28 17:42:42 ----
2019/05/28 17:42:42 div 0.7921679857762634 155294 37144 0.7608149703143714 0.9175800476238308
2019/05/28 17:42:42 newnode
2019/05/28 17:42:42 div 0.7696693339853163 117456 24832 0.7885846614902602 0.6940080239655406
2019/05/28 17:42:42 newnode
2019/05/28 17:42:42 div 0.7606479325369819 110276 23388 0.7879139613333817 0.6515838173513824
2019/05/28 17:42:42 newnode
2019/05/28 17:42:42 div 0.7515489614729173 103095 21944 0.7871477763228091 0.6091537020733502
2019/05/28 17:42:42 newnode
2019/05/28 17:42:42 div 0.7427994980606432 96063 20500 0.786598378147674 0.5676039777125199
2019/05/28 17:42:42 newnode
2019/05/28 17:42:42 div 0.7343602764195393 75277 14547 0.8067537229166943 0.44478649043091883
2019/05/28 17:42:42 newnode
2019/05/28 17:42:42 div 0.7280795049383273 70114 13565 0.806529366460336 0.4142800588502922
2019/05/28 17:42:42 newnode
2019/05/28 17:42:42 div 0.7225717917522931 65707 12737 0.8061545954007945 0.3882405771582872
2019/05/28 17:42:42 newnode
2019/05/28 17:42:42 div 0.717020969501087 61300 11909 0.805725938009788 0.3622010954662822
2019/05/28 17:42:42 newnode
2019/05/28 17:42:42 div 0.7117116448565635 56533 10960 0.8061309323757805 0.3340344947796955
2019/05/28 17:42:42 newnode
2019/05/28 17:42:42 div 0.7060304389249629 52209 10159 0.8054166906088989 0.30848543218921903
2019/05/28 17:42:42 newnode

*/
