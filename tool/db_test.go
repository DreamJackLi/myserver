package tool

import (
	"fmt"
	"math/rand"
	"myserver/api/apicompany"
	"myserver/api/apiemployee"
	"testing"
	"time"

	"github.com/jinzhu/gorm"
)

var arr [3500]rune
var db *gorm.DB

func init() {
	str := "早天才长安元字太其生地今般认为是唐剑南道绵州巴西郡昌隆后避玄宗讳改为昌明青莲乡祖籍为甘肃天水其家世家族皆不详据新唐书记载为兴圣皇帝凉武昭王暠世孙按照这个说法与唐诸王同宗是唐太宗世民的同辈族弟亦有说其祖是成或元吉神龙元月武则天去世岁发蒙读书始于是上安州裴长史书云岁诵甲甲唐代的小学识字课本长史州之次官开元岁已有诗赋多首并得到些社会名流的推崇与奖掖开始从事社会干谒活动亦开始接受道家思想的影响好剑术喜任是岑参生开元岁隐居戴天大匡山在今川省江油县内读书往来于旁郡先后出游江油剑阁梓州州治在今川省境内等地增长不少阅历与见识辞亲远游开元岁离开故乡而踏上远游的征途再游成都峨眉山然后舟行东下至渝州今重庆市开元出蜀仗剑去国辞亲远游开元岁春往扬州今江苏省扬州市秋病卧扬州冬离扬州北游汝州今河南省临汝县至安陆今湖北省安陆县途经陈州时与邕相识结识孟浩然开元是诏令民间有文武之高才者可到朝廷自荐秋全国州水灾州霜旱岁居于安陆寿山与故宰相许圉师之孙女结婚逐家于安陆是王昌龄进士及第开元土蕃屡次入侵岁早春出游江夏今湖北省武汉市与孟浩然相会于斯开元月日唐玄宗为自己岁生日举行盛大的庆贺活动并以每月日为千秋节诏令天下诸州宴乐休假日以宇文融管理全国财赋强制税法广为聚敛供朝廷奢侈之用岁在安陆蹉跎岁月开元岁春在安陆前此曾多次谒见本州裴史因遭人谗谤于近日上书自终为所拒初夏往长安谒宰相张说并结识其子张垍寓居终南山玉真公主玄宗御妹别馆又曾谒见其它王公大臣均无结果暮秋游邠州在长安之西冬游坊州在长安之北是杜甫岁游于晋今山西省开元玄宗多任宦官尤宠高力士时方表奏举杯邀明月的举杯邀明月的皆先为高力士所决月玄宗驾幸洛阳岁穷愁潦倒于长安自暴自弃与长安市井之徒交往初夏离长安经开封今河南省开封市到宋城今河南省商丘县秋到嵩山岳之的中岳为河南省登封县的名山恋故友元丹丘的山居所在逐有隐居之意暮秋滞留洛阳开元月玄宗出巡诏令巡幸所至地方官员可将本地区贤才直接向朝廷推荐月归还洛阳是全国户数为万余人口万余是有史以来的最高记录自春历夏在洛阳与元演崔成甫结识秋自洛阳返安陆途经南阳今河南省南阳市结识崔宗之冬元演阳到安陆相访人同游随州今湖北省随县岁未归家安陆开元正月唐玄宗亲注老子道德经令天下士庶身份很低的役人与庶民家藏册每贡举时加试老子策岁构石室于安陆兆山桃花岩开山田日以耕种读书为生活献赋谋仕开元正月为宗献上著作明堂赋赋云穹崇明堂倚天开兮又云门启兮万国来考休征兮进贤才俨若皇居而作固穷千祀兮悠哉！按赋中有臣美颂等字样疑太曾以此赋在东都洛阳进献玄宗此赋盛赞明堂之宏大壮丽写尽开元盛世的雄伟气象以及作者的政治理想明堂赋的写作目的是为谋求官位其写作时间为开元拆毁明堂之前他赋明堂是为谋仕的需要是以大道匡君的需要由于家庭的缘故不能应常举和制举以入仕途只能走献赋之路这是真献赋谋仕的原因开元玄宗又次狩猎正好也在西游乘机献上大猎赋希望能博得玄宗的赏识他的大猎赋希图以大道匡君示物周博而圣朝园池遐荒殚穷合幅员辽阔境况与前代大不相同夸耀本朝远胜汉朝并在结尾处宣讲道教的玄埋以契合玄宗当时崇尚道教的心情是进长安后结识卫尉张卿并通过他向玉真公主献诗最后两句说几时入少室王母应相逢是祝她入道成仙由此他步步地接近统治阶级的上层这次在长安还结识贺知章去紫极宫在那里遇见贺知章立刻上前拜见并呈上袖中的诗本贺知章颇为欣赏蜀道难和乌栖曲瑰丽的诗歌和潇洒出尘的风采令贺知章惊异万分竟说公非人世之人可不是星精耶？贺知章称他为谪仙人后发出行路难归去来的感叹离开长安供奉翰林图片图片天宝元公元由于玉真公主和贺知章的交口称赞玄宗看的诗赋对其分钦慕便召进宫进宫朝见那天玄宗降辇步迎以宝床赐食于前亲手调羹玄宗问到些当世事务凭半生饱学及长期对社会的观察胸有成竹对答如流玄宗大为赞赏随即令供奉翰林职务是给皇上写诗文娱乐陪侍皇帝左右玄宗每有宴请或郊游必命侍从利用他敏捷的诗才赋诗纪实虽非记功也将其文字流传后世以盛况向后人夸示受到玄宗如此的信同僚不胜艳羡但也有人因此而产生嫉恨之心天宝岁诏翰林院初春玄宗于宫中行乐奉诏作宫中行乐词赐宫锦袍暮春兴庆池牡丹盛开玄宗与杨玉环同赏又奉诏作清平调对御用文人生活日渐厌倦始纵酒以自昏秽与贺知章等人结酒中人仙之游玄宗之不朝尝奉诏醉中起草诏书引足令高力士脱靴宫中人恨之谗谤于玄宗玄宗疏之杜相识天宝载注天宝至载至德号期间称载而不称夏天到东都洛阳在这里他遇到杜甫中国文学史上最伟大的两位诗人见面此时已名扬全国而杜甫风华正茂却困守城比杜甫长岁但他并没有以自己的才名在杜甫面前倨傲而性豪也嗜酒结交皆老苍的杜甫也没有在面前味低头称颂两人以平等的身份建立深厚的友情在洛阳时他们约好下次在梁宋今开封商丘带会面访道求仙同秋天两人如约到梁宋两人在此抒怀遣借古评今他们还在这里遇到诗人高适高适此时也还没有禄位然而人各有大志理想相同人畅游甚欢评文论诗纵谈天下大势都为国家的隐患而担忧这时的杜都值壮此次两人在创作上的切磋对他们今后产生积极影响这的秋冬之际杜又次分手到齐州今山东济南带紫极宫请道士高天师如贵授道箓-从此他算是正式履行道教仪式成为道士其后又赴德州安陵县遇见这带善写符箓的盖还为他造真箓此次的求仙访道得到完满的结果天宝载秋天与杜甫在东鲁第次会见短短多的时间他们两次相约会见知交之情不断加深他们道寻访隐士高人也偕同去济州拜访过当时驰名天下的文章家书法家邕就在这冬天杜两人分手安史入幕天宝载公元安史之乱爆发与妻子宗氏道南与安史之乱与安史之乱奔避难春在当涂旋闻洛阳失陷中原横溃乃自当返宣城避难剡中今浙江省嵊州至溧阳今江苏省溧阳市与张旭相遇夏至越中闻郭子仪光弼在河北大胜又返金陵秋闻玄宗奔蜀遂沿长江西上入庐山屏风叠隐居天宝载至德元载正月安禄山在洛阳自称大燕皇帝月郭子仪光弼大破史思明收复河北余郡月安禄山率攻破潼关生擒哥舒翰至德载岁正月在永王军营作组诗永王东巡歌抒发建功报国情怀永王擅自引兵东巡导致征剿兵败在浔阳入狱被宋若思崔涣营救成为宋若思的幕僚后为宋写过些文表并跟随他到武昌在宋若思幕下很受重视并以宋的名义再次朝廷推荐希望再度能得到朝廷的任用终以参加永王东巡而被判罪长流夜郎今贵州桐梓是杜甫岁月从贼营逃出谒肃宗于风翔授左拾遗乾元元月史思明反月肃宗罢张镐宰相出为荆州大都督长史月史思明陷魏州今河北省南部岁自浔阳出开始长流夜郎妻弟宗嫌相送春末夏初途经西塞驿今武昌县东至江夏访邕故居登黄鹤楼眺望鹦鹉洲秋至江陵冬入峡是杜甫岁为华州司功参军溘然病逝乾元朝廷因关中遭遇大旱宣布大赦规定死者静夜思静夜思从流流以下完全赦免经过长期的辗转流离终于获得自由他随即顺着长江疾驶而下而那首著名的早发帝城最能反映他当时的心情到江夏由于老友良宰正在当地做太守便逗留阵乾元应友人之邀再次与被谪贬的贾至泛舟赏月于洞庭之上发思古之幽情赋诗抒怀不久又回到宣城金陵旧游地差不多有两的时间他往来于两地之间仍然依人为生上元已出头的因病返回金陵在金陵他的生活相当窘迫不得已只好投奔在当涂做县令的族叔阳冰上元病重在病榻上把手稿交给阳冰赋临终歌而与世长辞关于之死历来众说纷纭莫衷是总体可以概括为种死法其是醉死其是病死其是溺死第种死法见诸旧唐书说以饮酒过度醉死于宣城第种死法亦见诸其他正史或专家学者的考证之说说当光弼东镇临淮时不顾岁的高龄闻讯前往请缨杀敌希望在垂暮之为挽救国家危亡尽力因病中途返回次病死于当涂县令唐代最有名的篆书家阳冰处而第种死法则多见诸民间传说极富浪漫色彩说在当涂的江上饮酒因醉跳入水中捉月而溺死与诗人性格非常吻合但是不管哪种死法都因参与永王璘谋反作乱有着直接的关系因为流放夜郎遇赦得还后不久就结束他传奇而坎坷的生这是个不争的事实"
	src := []rune(str)
	for i, v := range src {
		if i >= len(arr) {
			break
		}
		arr[i] = v
	}
	db = GetDB()
}

func TestAutoMigrate(t *testing.T) {
	db.AutoMigrate(apicompany.CompanyInfo{}, apicompany.PreventionRecord{}, apiemployee.EmployeeInfo{}, apiemployee.EmployeeHealthInfo{}, apiemployee.EmployeeHealthRecord{})
}

func TestData(t *testing.T) {
	for i := 0; i < 100; i++ {
		companyName := "四川%s%s科技有限公司"
		company := &apicompany.CompanyInfo{
			Name:            fmt.Sprintf(companyName, string(arr[random()]), string(arr[random()])),
			Phone:           phone(),
			Province:        "四川省",
			ProvinceId:      203,
			City:            "成都市",
			CityId:          423,
			District:        "温江区",
			DistrictId:      3517,
			Street:          "柳城街道",
			StreetId:        1,
			Community:       "光华社区",
			CommunityId:     16,
			DetailedAddress: "珠江国际写字楼14F 1001-1018",
			Adcode:          "510115",
			Identity:        "AHFHJ654DFG21DFG65SDG22SDG",
			Industry:        "互联网",
			IsDelete:        1,
			CreateTime:      time.Now().Unix(),
			UpdateTime:      time.Now().Unix(),
		}
		db.Model(apicompany.CompanyInfo{}).Create(company)
		for i := 0; i < 10; i++ {
			preventionRecord := &apicompany.PreventionRecord{
				CompanyId:            company.Id,
				ReturneeNum:          18,
				MaterialImage:        "Fo40PPj9EfYRid4TNydZh5aOiVXP,FtposE7I9Q2PIYhfnYOtv_fgyjH",
				DisinfectionImage:    "FvCuxcMePsabIK-UVFmqAlXugJVl,FqMhDEcJHkaRQ-5EEL_9R6l8LJJ-",
				DisinfectionTime:     time.Now().Unix(),
				DisinfectionRange:    "前台，走廊，研发大厅，活动室，会议室",
				DisinfectionOperator: "黄锦绣",
				ReportEmployeeNum:    15,
				CreateTime:           time.Now().Unix(),
				UpdateTime:           time.Now().Unix(),
			}
			db.Model(apicompany.PreventionRecord{}).Create(preventionRecord)
			employeeInfo := &apiemployee.EmployeeInfo{
				Name:                     fmt.Sprintf(companyName, string(arr[random()]), string(arr[random()])),
				Phone:                    phone(),
				CompanyProvince:          "四川省",
				CompanyProvinceId:        203,
				CompanyCity:              "成都市",
				CompanyCityId:            423,
				CompanyDistrict:          "温江区",
				CompanyDistrictId:        3517,
				CompanyStreet:            "柳城街道",
				CompanyStreetId:          1,
				CompanyCommunity:         "光华社区",
				CompanyCommunityId:       16,
				CompanyName:              "四川小咖科技有限公司",
				CompanyId:                company.Id,
				IdCardNumber:             "510922198909010694",
				Gender:                   2,
				Age:                      30,
				ResidenceProvince:        "四川省",
				ResidenceProvinceId:      203,
				ResidenceCity:            "成都市",
				ResidenceCityId:          423,
				ResidenceDistrict:        "温江区",
				ResidenceDistrictId:      3517,
				ResidenceStreet:          "天府街道",
				ResidenceStreetId:        3,
				ResidenceCommunity:       "天府家园社区",
				ResidenceCommunityId:     38,
				ResidenceDetailedAddress: "天府家园41栋3单元509",
				CreateTime:               time.Now().Unix(),
				UpdateTime:               time.Now().Unix(),
				IsDelete:                 1,
			}
			db.Model(apiemployee.EmployeeInfo{}).Create(employeeInfo)
			employeeHealthInfo := &apiemployee.EmployeeHealthInfo{
				SymptomWithin14Days:          "无上述症状",
				FormHuBei:                    2,
				Touch_NCP:                    2,
				BeenToHuBei:                  2,
				ReturnChengDuTime:            time.Date(2020, 2, 2, 21, 0, 0, 0, time.Local).Unix(),
				ReturnChengDuDetailedAddress: "成都市温江区天府家园41栋3单元509",
				ReturnChengDuTransport:       "自驾",
				TransportDetailInfo:          "川A41LW2",
				TouchHuBeiPeopleWithin14Days: 2,
				LastTouchTime:                0,
				FromChengDuOutside:           2,
				ThisPlaceTime:                time.Date(2020, 2, 2, 21, 0, 0, 0, time.Local).Unix(),
				ReturnThisPlaceTransport:     "自驾",
				Profession:                   "一般员工/工人/服务员",
				LifeTrace:                    "居家，自驾，家人",
				SignImage:                    "FsbPd_zJykUK1BGQflsiyUUoN44B",
				EmployeeId:                   employeeInfo.Id,
				CompanyId:                    company.Id,
				UpdateTime:                   time.Now().Unix(),
				ComeFromWenjiang:             "市外",
				ComeFromAddress:              "四川省南充市阆中市",
				ComeFromDetailedAddress:      "千佛镇峰占乡红瓦店村",
			}
			db.Model(apiemployee.EmployeeHealthInfo{}).Create(employeeHealthInfo)
			for i := 0; i < 10; i++ {
				record := &apiemployee.EmployeeHealthRecord{
					EmployeeId:        employeeInfo.Id,
					Department:        "研发部",
					Temperature:       36.5,
					EpidemicSymptom:   2,
					DetectionResult:   "健康",
					DetectionOperator: "黄锦绣",
					CreatedTime:       time.Now().Unix(),
					CompanyId:         time.Now().Unix(),
				}
				db.Model(apiemployee.EmployeeHealthRecord{}).Create(record)
			}
		}
	}
}

func random() int {
	//将时间戳设置成种子数
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(1000) + 1
}

func phone() string {
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(10)
	return fmt.Sprintf("%d%d", time.Now().Unix(), random)
}
