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
	b, err := ioutil.ReadFile("1.htm")
	assert.NoError(t, err)
	s, err := clean.Preprocess(string(b), false, nil)
	assert.NoError(t, err)
	ioutil.WriteFile("1_1.htm", []byte(s), 0666)

	var d func(*goquery.Selection)
	d = func(s *goquery.Selection) {
		alldiv := true
		if len(s.Children().Nodes) == 0 {
			//txt := clean.TrimBytes([]byte(s.Text()))
			//if string(txt) != "" {
			alldiv = false
			return
			//s.Parent().Remove() //s.Remove()
			//}
		}
		for _, n := range s.Nodes {
			if n.Type != 3 {
				alldiv = false
				break
			}
			if n.Data != "div" {
				alldiv = false
				break
			}
		}
		if alldiv {
			s.ReplaceWithNodes(s.Children().Nodes...)
		}

	}
	/*
		doc, err := goquery.NewDocumentFromReader(strings.NewReader(s))
		doc.Find("div").Children().Each(func(_ int, s *goquery.Selection) {
			//log.Println(len(s.Nodes))
			d(s)
		})

		res, _ := doc.Html()*/
	//res = clean.MainContent(res)
	//res, _ = clean.Preprocess(res, true, nil)
	ioutil.WriteFile("1_2.htm", []byte(res), 0666)
}

/*
func Test22(t *testing.T) {
	url2file(t, "")
}
*/
