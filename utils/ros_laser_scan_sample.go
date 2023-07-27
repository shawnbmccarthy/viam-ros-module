package utils

import "math"

const Secs = 1690477544
const Nsecs =  502371703
const AngleMin = -3.12413907051
const AngleMax = 3.14159274101
const AngleIncrement = 0.00871450919658
const TimeIncrement = 0.000203415751457
const ScanTime = 0.146255925298
const RangeMin = 0.15000000596
const RangeMax = 12.0
var Ranges = []float64{3.5199999809265137, 3.503999948501587, 3.25600004196167, 3.25600004196167, 3.444000005722046, 3.4719998836517334, 3.4639999866485596, 3.4519999027252197, 3.447999954223633, 3.436000108718872, 3.4240000247955322, 3.4040000438690186, 3.4040000438690186, math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), inf, math.Inf(1), 1.4859999418258667, 1.4859999418258667, 1.4859999418258667, 1.4800000190734863, 1.4539999961853027, 1.4539999961853027, 1.496000051498413, math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), 1.5080000162124634, 1.5080000162124634, 1.5160000324249268, 2.7639999389648438, 2.7639999389648438, 2.740000009536743, 2.7200000286102295, 2.7239999771118164, 2.7239999771118164, 3.0399999618530273, 3.0399999618530273, 3.2320001125335693, 3.2160000801086426, 3.2119998931884766, 3.2200000286102295, 3.2160000801086426, 3.2119998931884766, 3.2160000801086426, 3.2079999446868896, 3.2039999961853027, 3.2119998931884766, 3.2119998931884766, 3.2160000801086426, 3.2200000286102295, 3.2239999771118164, 3.2239999771118164, 3.2239999771118164, 3.2239999771118164, 3.2279999256134033, 3.2360000610351562, 3.2320001125335693, 3.2360000610351562, 3.240000009536743, 3.24399995803833, 3.252000093460083, 3.252000093460083, 3.259999990463257, 3.2679998874664307, 3.2720000743865967, 3.2760000228881836, 3.2799999713897705, 3.2880001068115234, 3.2920000553131104, 3.303999900817871, 3.303999900817871, 3.312000036239624, 3.315999984741211, 3.328000068664551, 3.3320000171661377, 3.3359999656677246, 3.3480000495910645, 3.364000082015991, 3.375999927520752, 3.384000062942505, 3.3919999599456787, 3.4079999923706055, 3.4200000762939453, 3.436000108718872, 3.444000005722046, 3.4519999027252197, 3.4560000896453857, 3.4719998836517334, 3.4839999675750732, 3.496000051498413, 3.5, 3.51200008392334, 3.5160000324249268, 3.5360000133514404, 3.5480000972747803, 3.555999994277954, 3.5920000076293945, 3.611999988555908, 3.635999917984009, 3.6440000534057617, 3.6600000858306885, 3.6679999828338623, 3.687999963760376, 3.696000099182129, 3.7119998931884766, 3.7200000286102295, 3.752000093460083, 3.7760000228881836, 3.812000036239624, 3.8239998817443848, math.Inf(1), 2.6600000858306885, 2.6600000858306885, 2.6440000534057617, 2.635999917984009, 2.6080000400543213, 2.5959999561309814, 2.568000078201294, 2.559999942779541, 2.5480000972747803, 2.5439999103546143, 2.5320000648498535, 2.5239999294281006, 2.51200008392334, 2.5, 2.48799991607666, 2.4839999675750732, 2.4719998836517334, 2.4679999351501465, 2.431999921798706, 4.303999900817871, 4.303999900817871, 4.288000106811523, 4.23199987411499, 4.199999809265137, 4.136000156402588, 4.104000091552734, 4.076000213623047, 3.947999954223633, 3.9519999027252197, 3.9200000762939453, 3.9079999923706055, 3.880000114440918, 3.859999895095825, 3.819999933242798, 3.7920000553131104, 3.74399995803833, 3.7239999771118164, 3.696000099182129, 3.687999963760376, 3.6640000343322754, 3.6480000019073486, 3.615999937057495, 3.6040000915527344, 3.568000078201294, 3.552000045776367, 3.5199999809265137, 3.51200008392334, 3.4800000190734863, 3.4639999866485596, 2.936000108718872, 2.936000108718872, 2.9159998893737793, 2.940000057220459, 3.380000114440918, 3.368000030517578, 3.3440001010894775, 3.328000068664551, 3.312000036239624, 3.299999952316284, 3.2760000228881836, 3.259999990463257, 3.25600004196167, 3.259999990463257, 5.28000020980835, 5.28000020980835, 5.25600004196167, 5.223999977111816, 5.216000080108643, 5.208000183105469, 5.208000183105469, 5.208000183105469, 5.208000183105469, 5.199999809265137, 5.191999912261963, 5.184000015258789, 5.176000118255615, 5.144000053405762, 5.144000053405762, 5.111999988555908, 5.0960001945495605, 5.079999923706055, 5.079999923706055, 5.047999858856201, 5.039999961853027, 5.007999897003174, 5.0, 4.984000205993652, 4.97599983215332, 4.960000038146973, 4.943999767303467, 4.927999973297119, 4.927999973297119, 4.927999973297119, 4.927999973297119, 4.9120001792907715, 4.9120001792907715, 4.9120001792907715, 3.0320000648498535, 3.007999897003174, 3.007999897003174, 2.7960000038146973, 2.7960000038146973, 2.7839999198913574, 2.9800000190734863, 2.9800000190734863, 2.9800000190734863, 2.9719998836517334, 2.9760000705718994, 2.9719998836517334, 2.936000108718872, 2.9760000705718994, 2.9679999351501465, 2.9679999351501465, 2.9760000705718994, 2.9760000705718994, 2.9719998836517334, 2.9760000705718994, 2.9719998836517334, 2.9760000705718994, 2.9800000190734863, 2.9800000190734863, 2.9800000190734863, 2.9839999675750732, 2.9519999027252197, 2.7720000743865967, 2.7760000228881836, 2.7799999713897705, 2.7920000553131104, 2.7960000038146973, 2.808000087738037, math.Inf(1), math.Inf(1), math.Inf(1), 0.4490000009536743, 0.4490000009536743, 0.4449999928474426, 0.4440000057220459, 0.4480000138282776, 0.44999998807907104, 0.45100000500679016, 0.45100000500679016, 0.45500001311302185, 0.4560000002384186, 0.45899999141693115, 0.46000000834465027, 0.4620000123977661, 0.46399998664855957, 0.4650000035762787, 0.4690000116825104, 0.47099998593330383, 0.4729999899864197, 0.47999998927116394, 0.5090000033378601, 0.5099999904632568, 0.5120000243186951, 0.5139999985694885, 0.5180000066757202, 0.5220000147819519, 0.5260000228881836, 0.5299999713897705, 0.5339999794960022, 0.5379999876022339, 0.5440000295639038, 0.5479999780654907, 0.550000011920929, 0.5580000281333923, 0.5619999766349792, 0.5699999928474426, 0.5740000009536743, 0.5799999833106995, 0.5860000252723694, 0.5899999737739563, 0.6000000238418579, 0.6060000061988831, 0.6159999966621399, 0.621999979019165, 0.628000020980835, 0.6399999856948853, 0.6480000019073486, 0.6620000004768372, 0.671999990940094, 0.6800000071525574, 0.699999988079071, 0.7099999785423279, 0.7319999933242798, 0.7459999918937683, 0.7599999904632568, 0.7900000214576721, 0.7900000214576721, math.Inf(1), math.Inf(1), math.Inf(1), 2.187999963760376, 2.187999963760376, 2.1679999828338623, 2.1679999828338623, 2.1559998989105225, 2.1480000019073486, 2.140000104904175, 2.135999917984009, 2.1480000019073486, math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), 1.5299999713897705, 1.5299999713897705, 1.534000039100647, 1.534000039100647, 1.5360000133514404, 1.5399999618530273, 1.5440000295639038, 1.5499999523162842, 1.5499999523162842, 1.5440000295639038, 1.5379999876022339, 1.527999997138977, 1.5260000228881836, 1.5219999551773071, 1.5160000324249268, 1.503999948501587, 1.49399995803833, 1.4620000123977661, math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), 1.4420000314712524, 1.4420000314712524, 1.437999963760376, 1.4359999895095825, 1.4320000410079956, 1.4299999475479126, 1.4320000410079956, 1.4320000410079956, 1.434000015258789, 1.4359999895095825, 1.4359999895095825, 1.440000057220459, 1.440000057220459, 1.4420000314712524, 1.444000005722046, 1.4459999799728394, 1.4479999542236328, 1.4500000476837158, 1.4520000219345093, 1.4539999961853027, 1.3660000562667847, 1.3339999914169312, 1.3020000457763672, 1.2419999837875366, 1.190000057220459, 1.1679999828338623, 1.121999979019165, 1.0779999494552612, 1.0579999685287476, 1.0240000486373901, 1.003999948501587, 0.9700000286102295, 0.9559999704360962, 0.9279999732971191, 0.8999999761581421, 0.8859999775886536, 0.8560000061988831, 0.8479999899864197, 0.8259999752044678, 0.8019999861717224, 0.7919999957084656, 0.7739999890327454, 0.7599999904632568, 0.7540000081062317, 0.7419999837875366, 0.734000027179718, math.Inf(1), math.Inf(1), 0.6340000033378601, 0.6340000033378601, 0.6299999952316284, 0.6299999952316284, 0.6259999871253967, 0.6240000128746033, 0.6119999885559082, 0.6060000061988831, 0.5960000157356262, 0.5860000252723694, 0.5799999833106995, 0.5699999928474426, 0.5640000104904175, 0.5559999942779541, 0.5460000038146973, 0.5419999957084656, 0.5339999794960022, 0.5299999713897705, 0.5199999809265137, 0.5130000114440918, 0.5099999904632568, 0.5019999742507935, 0.49900001287460327, 0.49000000953674316, 0.4790000021457672, 0.47099998593330383, math.Inf(1), math.Inf(1), 0.46299999952316284, 0.46299999952316284, 0.4569999873638153, 0.45100000500679016, 0.44699999690055847, 0.44200000166893005, 0.4359999895095825, 0.4339999854564667, 0.42800000309944153, 0.4259999990463257, 0.421999990940094, 0.4189999997615814, 0.41499999165534973, 0.41100001335144043, 0.4090000092983246, 0.40400001406669617, 0.4020000100135803, 0.39899998903274536, 0.3970000147819519, 0.3930000066757202, 0.3889999985694885, 0.3869999945163727, 0.3840000033378601, 0.38199999928474426, 0.3779999911785126, 0.375, 0.37400001287460327, 0.3700000047683716, 0.36800000071525574, 0.36500000953674316, 0.36399999260902405, 0.3619999885559082, 0.35899999737739563, 0.3580000102519989, 0.3610000014305115, 0.3619999885559082, 0.3409999907016754, 0.3409999907016754, 0.3440000116825104, 0.3440000116825104, 0.34599998593330383, 0.34700000286102295, 0.34299999475479126, 0.3400000035762787, 0.33899998664855957, 0.335999995470047, 0.33500000834465027, 0.11599999666213989, 0.3330000042915344, 0.33000001311302185, 0.33000001311302185, 0.32899999618530273, 0.32600000500679016, 0.32499998807907104, 0.3230000138282776, 0.32100000977516174, 0.32199999690055847, 0.3230000138282776, math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), 0.3160000145435333, 0.3160000145435333, 0.3149999976158142, 0.31299999356269836, 0.3109999895095825, 0.3100000023841858, 0.3100000023841858, 0.3089999854564667, 0.3089999854564667, 0.3089999854564667, 0.3070000112056732, 0.3059999942779541, 0.3059999942779541, 0.3059999942779541, 0.3050000071525574, 0.30300000309944153, 0.3019999861717224, 0.29600000381469727, math.Inf(1), math.Inf(1), 0.3050000071525574, 0.3050000071525574, 0.30300000309944153, 0.3009999990463257, 0.3009999990463257, 0.30000001192092896, 0.30000001192092896, 0.29899999499320984, 0.3009999990463257, 0.29899999499320984, 0.2980000078678131, 0.29899999499320984, 0.29899999499320984, 0.29899999499320984, 0.29899999499320984, 0.2980000078678131, 0.2980000078678131, 0.29899999499320984, 0.29899999499320984, 0.29899999499320984, 0.2980000078678131, 0.2980000078678131, 0.2980000078678131, 0.29899999499320984, 0.29899999499320984, 0.29899999499320984, 0.30000001192092896, 0.3009999990463257, 0.3019999861717224, 0.30000001192092896, 0.3009999990463257, 0.3009999990463257, 0.3009999990463257, 0.3019999861717224, 0.30300000309944153, 0.30300000309944153, 0.30300000309944153, 0.30399999022483826, 0.3050000071525574, 0.3050000071525574, 0.3050000071525574, 0.3059999942779541, 0.3070000112056732, 0.3070000112056732, 0.30799999833106995, 0.3089999854564667, 0.3100000023841858, 0.3100000023841858, 0.3109999895095825, 0.3109999895095825, 0.31200000643730164, 0.31299999356269836, 0.3140000104904175, 0.3149999976158142, 0.3160000145435333, 0.3179999887943268, 0.3179999887943268, 0.3190000057220459, 0.3190000057220459, 0.3199999928474426, 0.32199999690055847, 0.3230000138282776, 0.3240000009536743, 0.32499998807907104, 0.3269999921321869, 0.3269999921321869, 0.33000001311302185, 0.3310000002384186, 0.3319999873638153, 0.33399999141693115, 0.33500000834465027, 0.3370000123977661, 0.33799999952316284, 0.3400000035762787, 0.3409999907016754, 0.34299999475479126, 0.3449999988079071, 0.34599998593330383, 0.3490000069141388, 0.3499999940395355, 0.3529999852180481, 0.3540000021457672, 0.3569999933242798, 0.35899999737739563, 0.36000001430511475, 0.3630000054836273, 0.36399999260902405, 0.367000013589859, 0.36899998784065247, 0.3709999918937683, 0.37400001287460327, 0.37599998712539673, 0.38100001215934753, 0.38199999928474426, 0.38600000739097595, 0.3880000114440918, 0.3889999985694885, 0.39399999380111694, 0.3959999978542328, 0.4000000059604645, math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), 0.4480000138282776, 0.4480000138282776, 0.44999998807907104, math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), 2.436000108718872, 2.436000108718872, 2.371999979019165, 2.359999895095825, 2.3440001010894775, math.Inf(1), 2.3519999980926514, 2.3519999980926514, 2.368000030517578, math.Inf(1), 2.48799991607666, 2.48799991607666, 2.5280001163482666, 2.552000045776367, 2.5999999046325684, 2.6440000534057617, 2.6679999828338623, 2.691999912261963, 3.6480000019073486, 3.6480000019073486, 3.640000104904175, 3.624000072479248, 3.5959999561309814, 3.568000078201294, 3.555999994277954, 3.5320000648498535}
var Intensities = []float64{47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 0.0, 0.0, 0.0, 0.0, 0.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 0.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 0.0, 0.0, 0.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 0.0, 0.0, 0.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 0.0, 0.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 0.0, 0.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 0.0, 0.0, 0.0, 0.0, 0.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 0.0, 0.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 47.0, 47.0, 47.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 47.0, 47.0, 47.0, 47.0, 47.0, 0.0, 47.0, 47.0, 47.0, 0.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0, 47.0}