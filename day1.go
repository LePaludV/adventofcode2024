package main

import (
	"fmt"
	"math"
	"slices"
)

func main() {
	slice1 := []int{97924,50267,98415,64856,73972,85009,16084,63342,36722,71800,13657,31646,13033,98175,12224,26745,91733,63235,34149,55819,89877,63876,35017,71693,67230,32107,29455,10294,11503,45021,65889,22646,29246,94894,97487,11658,86638,77389,20834,43948,25169,47469,77677,67863,66135,22510,39330,96958,65629,40599,41131,41361,44833,32056,15025,32170,95391,55371,48305,65862,40636,97450,77105,34812,23999,25503,66670,30994,76175,43005,53741,67639,68473,16078,51006,74520,17232,25445,39404,23041,54140,46419,49484,19164,39289,59326,47693,89369,78926,46255,53879,27064,21756,46201,83358,67814,36582,68742,63598,46069,43061,82842,47757,11761,55059,73640,87885,76819,99933,96701,33926,53211,63392,86076,22252,82469,15384,17446,60995,67563,81796,13959,67714,98828,66905,12690,24626,14738,58555,23470,38356,21420,85109,88215,13822,78720,17622,27946,42737,39857,98335,51079,75618,55716,44439,74943,15914,94892,89851,58970,89048,79217,22327,15513,52584,17822,50437,26861,39402,44176,93367,85462,70936,64975,69603,27447,48913,55729,55903,58608,13030,79586,85885,47696,89036,51691,74369,83246,34233,87921,12660,15408,82361,78963,87395,73712,97972,77207,26598,51812,34508,26436,71011,50890,50493,16217,44890,44429,99607,38107,10208,73098,26515,53750,74904,89898,71445,77746,20131,67840,36620,18892,73474,83450,45728,66927,35585,93945,76199,26536,58275,51874,88833,75871,67764,69526,19475,80585,34189,96205,11371,38079,99465,21120,15731,82309,33121,84859,93071,69403,47701,94758,71944,59059,23522,69897,96432,40161,99974,28726,43411,52521,10024,84954,33296,77780,89571,13582,27081,29330,35469,96455,44473,52280,36181,83515,77203,31113,93446,94776,65132,76822,76108,88359,58602,71874,78276,39092,62278,40915,55896,89049,73736,41226,91298,48613,48763,72203,83036,40070,34155,66614,99172,10261,75196,60212,82136,68902,61584,18196,12112,28990,21082,45932,71891,93524,36886,84927,90640,12240,50754,12243,21896,27361,92599,50881,19258,48109,39617,33422,24453,70071,61099,32270,18554,57108,72972,44178,58969,68638,59752,91349,42541,11241,56014,18240,83562,99794,49642,26318,20645,95467,85129,63111,21776,17413,39834,54442,25618,67027,91899,36861,41069,39818,41305,50844,35314,13646,21610,84316,37390,33072,50742,95333,67978,54472,65124,62367,88774,71725,27200,72832,93194,84331,83654,98518,26270,17471,40445,24801,71775,63886,11574,36092,70761,61291,97303,44506,91586,58301,81698,72921,50502,16298,75558,22495,33661,85183,57125,32610,61179,59372,74037,48967,57442,26047,23188,23633,15300,11726,59493,66596,78092,48254,23338,36327,14142,26946,51040,69209,15882,31997,12496,54801,92254,54184,50128,46755,68205,16526,33415,95561,58231,12000,13762,20783,49870,31028,73097,36095,15523,41560,30530,34128,91541,97402,14133,95923,21065,54343,17923,39271,54329,29173,73679,25043,23367,47733,50178,61856,12334,24075,66813,68591,60548,47476,65067,27139,90127,71661,76564,14606,33073,28929,20290,55600,44895,69018,19996,79035,59205,38746,86509,99613,21577,97880,91425,84219,82752,24031,21843,67480,89728,28954,88314,18726,87526,85951,64143,42517,13848,43260,53772,54544,38528,75966,66179,58180,53138,43795,75091,82097,58590,50521,27323,33112,27922,65787,40517,20919,30416,88719,79993,27124,35850,47642,56934,99900,21591,97673,29342,75687,45898,22897,23889,63662,65299,44563,60281,49181,50909,21600,44988,58351,84132,33795,30993,23660,56833,84296,27172,41123,70231,65636,24047,11869,44666,81159,75418,18982,54865,91042,21791,86188,92550,63942,55260,97155,44593,94520,16803,99284,10940,51986,94727,79564,67524,96074,32231,82182,36414,16281,21085,98238,53299,19359,46340,16355,70234,50841,81149,58011,80355,62858,25160,29287,82307,45594,52167,66037,16339,74466,68243,70410,18833,52684,99904,64936,67435,25168,48685,87833,16974,82852,54986,56112,64314,66607,61959,20294,22416,15287,42081,69276,19958,14337,43966,56772,37195,71264,45446,27319,71714,98898,15964,35923,92462,70969,60163,26799,57327,19339,56504,53113,78876,96200,82490,44316,71051,89880,27972,58858,31090,99124,77060,67020,28781,58289,89174,38323,42607,12543,23349,70058,15437,59862,30674,14449,42471,41124,77330,91002,12075,51581,17412,57784,78558,68326,33313,18353,91690,31047,41819,97743,81549,72159,38296,39901,26924,77429,90666,25384,61440,42764,26668,44171,96575,77777,65805,66448,66204,67751,48577,14034,82234,94447,21669,76758,16998,10808,91668,42411,83439,57286,26423,47687,75978,79027,13122,43173,33696,26050,29903,55762,22725,92748,75292,84901,82424,49756,99049,21549,45030,42687,24675,64330,26258,11886,10674,80983,73290,91623,68510,96050,29559,87297,16161,69453,79979,28404,95222,61168,39809,24627,48823,22852,67111,22838,19454,84203,97780,53203,89425,16614,19219,20569,48161,17213,70935,86566,24611,54359,95758,20532,49070,66014,69011,54241,61068,64852,60260,26299,58452,26134,84291,94851,90147,88435,26639,98570,66700,57045,30983,77419,61144,79828,38801,71141,30543,33787,93255,27104,43923,70069,23376,82293,37692,91328,47801,90140,40772,77630,26106,82191,29674,32858,57709,70834,34807,40646,17334,79995,22031,59137,23515,69665,78363,74969,51172,11229,86385,71049,90321,15381,17377,35732,92220,75143,93217,12345,66240,63460,14334,52629,67549,32589,40973,48086,62405,69385,73015,47116,89461,89978,45093,95580,90761,61770,91506,37148,73119,83826,11602,79485,51373,71333,41091,60255,70778,76205,67609,65250,81623,20105,53573,67594,30437,74127,43012,82823,87461,35267,99386,47451,51886,45000,93502,26127,61492,16858,77633,29357,91974,87979,12707,38824,46482,69244,10765,30619,30590,99292,28244,18494,91105,73930,30813,67308,54758,43568,15869,82010,33131,57985,93101,15835,24286,56944,47755,82704,93787,51466,51856,30252,44052,20691,57033,13300,33712,97080,91808,90652,64155,65463,44798,43936,18253,73289,71164,34646,42097,15990,56628,68636,67043,25652,87109,83865,54773,12778,41699,59666,94699,77535,19069,34344,24906,91226,58992,21084,95352,14593,57258,41550,59908,63930,52103,98431,64232,10638,52217,11797,12501,39316,62625,54433,23801,15292,28822,90501,52285,42385,34500,39882,12013,42691,57217,31320,90532,65216,78393,48315,37332,40774,16062,73931,54344,17846,61807,93272}
	slice2 := []int{12015,32019,10716,63472,11396,60876,81584,45754,57910,33139,42097,60883,45362,50742,82309,36160,22906,23212,43474,92797,82309,68418,24286,56515,56504,70146,43966,80966,97276,44890,45042,50684,40041,64304,32858,25139,37390,69681,20641,97853,18624,14334,51466,10346,44506,90140,56504,57081,98292,32865,50742,91497,76741,53741,58921,52656,10777,78010,51466,89174,54184,99374,83865,41764,77192,93924,24286,12382,60974,50445,14287,60769,53741,76080,54184,49015,40707,89174,16084,85801,17286,20749,46855,86911,12150,22031,53741,86087,99284,78050,65533,21120,63733,58079,45402,39702,27906,32858,80368,15449,21120,43966,41841,26946,78870,28789,16294,42168,82831,43783,50742,50456,97067,89174,52322,52960,19277,62973,85599,99208,42097,65386,81831,55717,59486,99284,21468,49115,17335,27139,94975,34794,99284,16754,56504,97465,49756,39646,85951,51466,73635,27139,80096,22031,79910,85951,69859,51466,26946,40772,21927,43966,21545,56526,39743,37390,45539,99284,43966,98943,38719,35314,30406,53741,40772,43966,50742,17033,55600,73308,62519,53741,78724,26716,95563,35314,27139,50742,53741,48574,51466,54159,21120,99205,10043,24286,97487,34050,21574,47558,90761,40460,60574,51466,24854,47118,73772,26946,43966,99307,32858,45260,70155,55600,44506,43966,37467,82076,34528,12429,80900,95044,85951,24551,11046,40772,85951,83865,18892,16084,26910,42097,57054,85642,32556,26946,16084,90761,82309,55600,24399,51194,44890,99284,84749,28950,47050,31889,22719,40772,55127,12405,40772,27139,48009,20227,53857,56742,31605,40772,19960,51466,51312,35314,99284,49638,26843,87187,74290,26946,42097,60260,20645,90488,17193,89920,22203,70569,50742,14334,74481,22765,25212,81025,89174,21120,98125,22656,65305,21120,26933,11422,82035,93643,94509,68016,50742,25900,90823,22471,77007,75044,34719,48242,91994,32858,42097,40955,61094,97838,99284,51912,40991,51466,70388,71891,20645,11097,82280,22031,79877,56504,69022,50742,91087,17965,92324,51450,66508,59580,21120,76583,99284,81634,90140,40772,82180,35314,35314,32858,97487,85951,91799,85951,40772,82309,82349,20645,24286,35314,16436,32543,29179,28393,47366,13453,96167,65122,90140,50742,59045,84429,24405,26532,57846,54184,52839,66988,72792,67867,27139,38672,68302,44890,10314,77425,64037,57908,40958,85651,73263,16084,20645,87651,96024,44506,85951,54184,40716,98329,48960,62621,58455,85951,10112,26292,19625,86054,95307,40772,27139,26946,36406,86304,37706,99284,82309,85951,55600,99284,38278,93008,14334,43966,16159,21120,89174,16084,57527,50556,94320,32570,75310,42097,50742,78906,94284,68750,40772,40772,54414,14334,32775,60587,34864,56224,37390,37390,42097,31984,89174,99284,11999,27968,24130,29706,13617,85455,31238,22031,97714,34014,90052,97988,54184,68624,22098,88909,76908,27139,52386,97487,47300,12083,66145,79638,27139,73848,19841,75136,99284,99000,16084,37390,26946,51466,34238,43966,90140,53741,86323,74629,20645,99210,16026,31919,95367,51466,10788,39402,55228,94339,76500,75251,12189,60387,85951,85951,51466,14334,51466,16084,72818,99145,89174,32858,25968,85741,78312,49699,22031,35799,81476,99120,26946,82991,95509,45829,99284,67657,16278,47194,11318,83696,81208,99284,49371,90869,36120,48795,63304,50742,75970,51458,70479,30113,42097,51466,33043,89174,22031,90761,50821,63031,77846,48377,75957,99284,51466,35314,52749,11254,13205,33964,90140,94103,90836,42097,67799,30712,23612,64339,50742,35314,72155,60279,91711,79500,66073,51466,26946,14334,89073,22769,56504,48743,92181,69816,25507,37390,31162,94474,16434,32334,54184,22031,42754,54184,87988,11698,29276,28226,86754,43817,42097,93023,72842,51322,51466,13798,71891,53741,58080,56427,14155,26946,37276,10162,67510,18892,26946,59737,96568,82183,81661,20077,14334,32858,35314,51466,52200,43966,56000,91407,42097,49272,67342,54466,45397,14988,99280,22031,17788,85883,89174,24625,45494,64995,65787,60398,14947,54184,56504,54184,53181,34140,99284,21414,81686,45525,92053,51466,91566,44890,51902,89174,70725,36065,88682,22031,74592,41411,64119,40772,35314,25937,22031,99284,50742,63710,42383,89055,38990,59526,50742,98161,37596,18706,22031,32858,44890,14369,40772,75361,76932,38581,99284,18310,54184,52449,32831,33504,19604,65016,20645,33409,27139,97487,14334,52218,42785,60133,92089,30200,94108,93942,54184,25665,84579,48617,20645,33623,64139,62738,33385,92116,80984,43995,91461,14581,92007,65834,51090,56504,85951,90581,35314,96347,76475,64852,85951,53741,49202,67069,93128,37390,46134,83013,53741,10496,95224,42097,90831,43966,43966,78889,21120,18810,55709,36441,20645,50742,82209,27139,40772,79655,22031,73178,36506,82655,38807,81721,74305,16084,82344,44523,26946,97487,33107,54184,35714,27139,97487,18100,99284,82309,40982,31332,40772,78489,91868,97873,54184,73488,58755,54184,69128,99864,82574,26325,90900,10028,27139,22031,31012,23134,21120,46853,32875,64925,61354,31104,77622,62276,34934,81916,34632,73192,33571,93024,97304,93186,23461,73947,79180,51960,27139,69682,37390,50742,21893,62221,24286,79832,85951,35263,50742,74783,23868,93606,72970,39402,14334,40772,23407,59362,31953,85136,40772,98531,53822,73230,64852,24392,14334,26946,98405,22448,22031,80271,40772,31057,59725,82139,57710,21120,51466,64809,56431,65345,28247,35314,53741,37390,43966,20645,22983,37390,17021,58211,29135,16084,46435,43966,14334,50742,71891,11833,44890,26946,32858,83865,12242,17154,89174,55600,56504,98641,24286,90761,43966,22031,71177,23934,43966,98769,69179,42438,44506,20657,89949,86929,85951,14334,83865,83968,25275,35249,99284,53741,55600,85951,43966,85951,44506,26946,85951,94250,85951,33487,34016,27139,52450,12045,51270,64414,23491,11795,10809,75027,90140,20555,14334,44890,31936,68431,63455,43535,84239,51559,22031,22031,44890,40772,52819,44156,27360,89174,56212,20645,89582,55600,41785,89174,16186,34811,70679,43966,24922,27139,62346,17557,37390,27125,18324,51466,41267,76330,17879,37390,27424,14167,79480,98477,49327,56501,49712,19293,64852,14334,35314,27139,53741,73238,54184,27493,35877,50742,55302,18892,33785,21120,97487,93726,47966,18536,14334,58173,42331,54184,42097,29944,48362,98927,28783,39391,21435,56504,42097,97487}
	sum := 0.
	slices.Sort(slice2)
	slices.Sort(slice1)

	for i := 0;i < len(slice1); i++ {
		a := float64(slice1[i])
		b := float64(slice2[i])
		sum += math.Abs(a-b)
	}
	fmt.Printf("%.f\n", sum)
	similarityScore := 0
	for i := 0; i < len(slice1); i++ {
		a := slice1[i]
		b := slice2[0]
		k:= 0
		multiplier := 0
		for(!(a<b || k+1>=len(slice1))) {
			if(a == b) {
				multiplier++
			}

			k++
			b= slice2[k]
			
		}
		similarityScore += a * multiplier
	}
	fmt.Println("final",similarityScore)

}
