package main

import . "github.com/pyrex41/shen-go/kl"

var SequentMain = MakeNative(func(__e *ControlFlow) {
tmp12236 := MakeNative(func(__e *ControlFlow) {
V3019 := __e.Get(1)
_ = V3019
tmp12237 := MakeNative(func(__e *ControlFlow) {
W3020 := __e.Get(1)
_ = W3020
tmp12239 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3020)


if True == tmp12239 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3020)
return
}


}, 1)

tmp12259 := PrimIsPair(V3019)

var ifres12240 Obj

if True == tmp12259 {
tmp12241 := MakeNative(func(__e *ControlFlow) {
W3021 := __e.Get(1)
_ = W3021
tmp12242 := MakeNative(func(__e *ControlFlow) {
W3022 := __e.Get(1)
_ = W3022
tmp12243 := MakeNative(func(__e *ControlFlow) {
W3023 := __e.Get(1)
_ = W3023
tmp12253 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3023)


if True == tmp12253 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12244 := MakeNative(func(__e *ControlFlow) {
W3024 := __e.Get(1)
_ = W3024
tmp12245 := MakeNative(func(__e *ControlFlow) {
W3025 := __e.Get(1)
_ = W3025
tmp12246 := MakeNative(func(__e *ControlFlow) {
W3026 := __e.Get(1)
_ = W3026
tmp12247 := Call(__e, PrimFunc(symfn), W3021)


__e.TailApply(PrimFunc(symshen_4remember_1datatype), W3021, tmp12247)
return


}, 1)

tmp12248 := Call(__e, PrimFunc(symshen_4rules_1_6prolog), W3021, W3024)


tmp12249 := Call(__e, tmp12246, tmp12248)


__e.TailApply(PrimFunc(symshen_4comb), W3025, tmp12249)
return


}, 1)

tmp12250 := Call(__e, PrimFunc(symshen_4in_1_6), W3023)


__e.TailApply(tmp12245, tmp12250)
return


}, 1)

tmp12251 := Call(__e, PrimFunc(symshen_4_5_1out), W3023)


__e.TailApply(tmp12244, tmp12251)
return


}


}, 1)

tmp12254 := Call(__e, PrimFunc(symshen_4_5datatype_1rules_6), W3022)


__e.TailApply(tmp12243, tmp12254)
return


}, 1)

tmp12255 := Call(__e, PrimFunc(symtail), V3019)


__e.TailApply(tmp12242, tmp12255)
return


}, 1)

tmp12256 := Call(__e, PrimFunc(symhead), V3019)


tmp12257 := Call(__e, tmp12241, tmp12256)


ifres12240 = tmp12257


} else {
tmp12258 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12240 = tmp12258


}

__e.TailApply(tmp12237, ifres12240)
return


}, 1)

tmp12260 := Call(__e, ns2_1set, symshen_4_5datatype_6, tmp12236)


_ = tmp12260

tmp12261 := MakeNative(func(__e *ControlFlow) {
V3027 := __e.Get(1)
_ = V3027
V3028 := __e.Get(2)
_ = V3028
tmp12262 := PrimValue(symshen_4_ddatatypes_d)

tmp12263 := Call(__e, PrimFunc(symshen_4assoc_1_6), V3027, V3028, tmp12262)


tmp12264 := PrimSet(symshen_4_ddatatypes_d, tmp12263)

_ = tmp12264

tmp12265 := PrimValue(symshen_4_dalldatatypes_d)

tmp12266 := Call(__e, PrimFunc(symshen_4assoc_1_6), V3027, V3028, tmp12265)


tmp12267 := PrimSet(symshen_4_dalldatatypes_d, tmp12266)

_ = tmp12267

__e.Return(V3027)
return


}, 2)

tmp12268 := Call(__e, ns2_1set, symshen_4remember_1datatype, tmp12261)


_ = tmp12268

tmp12269 := MakeNative(func(__e *ControlFlow) {
V3029 := __e.Get(1)
_ = V3029
tmp12270 := MakeNative(func(__e *ControlFlow) {
W3030 := __e.Get(1)
_ = W3030
tmp12289 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3030)


if True == tmp12289 {
tmp12271 := MakeNative(func(__e *ControlFlow) {
W3037 := __e.Get(1)
_ = W3037
tmp12273 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3037)


if True == tmp12273 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3037)
return
}


}, 1)

tmp12274 := MakeNative(func(__e *ControlFlow) {
W3038 := __e.Get(1)
_ = W3038
tmp12285 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3038)


if True == tmp12285 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12275 := MakeNative(func(__e *ControlFlow) {
W3039 := __e.Get(1)
_ = W3039
tmp12276 := MakeNative(func(__e *ControlFlow) {
W3040 := __e.Get(1)
_ = W3040
tmp12281 := Call(__e, PrimFunc(symempty_2), W3039)


var ifres12277 Obj

if True == tmp12281 {
ifres12277 = Nil


} else {
tmp12278 := Call(__e, PrimFunc(symshen_4app), W3039, MakeString("\n ..."), symshen_4r)


tmp12279 := PrimStringConcat(MakeString("datatype syntax error here:\n "), tmp12278)

tmp12280 := PrimSimpleError(tmp12279)

ifres12277 = tmp12280


}

__e.TailApply(PrimFunc(symshen_4comb), W3040, ifres12277)
return


}, 1)

tmp12282 := Call(__e, PrimFunc(symshen_4in_1_6), W3038)


__e.TailApply(tmp12276, tmp12282)
return


}, 1)

tmp12283 := Call(__e, PrimFunc(symshen_4_5_1out), W3038)


__e.TailApply(tmp12275, tmp12283)
return


}


}, 1)

tmp12286 := Call(__e, PrimFunc(sym_5_b_6), V3029)


tmp12287 := Call(__e, tmp12274, tmp12286)


__e.TailApply(tmp12271, tmp12287)
return


} else {
__e.Return(W3030)
return
}


}, 1)

tmp12290 := MakeNative(func(__e *ControlFlow) {
W3031 := __e.Get(1)
_ = W3031
tmp12305 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3031)


if True == tmp12305 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12291 := MakeNative(func(__e *ControlFlow) {
W3032 := __e.Get(1)
_ = W3032
tmp12292 := MakeNative(func(__e *ControlFlow) {
W3033 := __e.Get(1)
_ = W3033
tmp12293 := MakeNative(func(__e *ControlFlow) {
W3034 := __e.Get(1)
_ = W3034
tmp12300 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3034)


if True == tmp12300 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12294 := MakeNative(func(__e *ControlFlow) {
W3035 := __e.Get(1)
_ = W3035
tmp12295 := MakeNative(func(__e *ControlFlow) {
W3036 := __e.Get(1)
_ = W3036
tmp12296 := Call(__e, PrimFunc(symappend), W3032, W3035)


__e.TailApply(PrimFunc(symshen_4comb), W3036, tmp12296)
return


}, 1)

tmp12297 := Call(__e, PrimFunc(symshen_4in_1_6), W3034)


__e.TailApply(tmp12295, tmp12297)
return


}, 1)

tmp12298 := Call(__e, PrimFunc(symshen_4_5_1out), W3034)


__e.TailApply(tmp12294, tmp12298)
return


}


}, 1)

tmp12301 := Call(__e, PrimFunc(symshen_4_5datatype_1rules_6), W3033)


__e.TailApply(tmp12293, tmp12301)
return


}, 1)

tmp12302 := Call(__e, PrimFunc(symshen_4in_1_6), W3031)


__e.TailApply(tmp12292, tmp12302)
return


}, 1)

tmp12303 := Call(__e, PrimFunc(symshen_4_5_1out), W3031)


__e.TailApply(tmp12291, tmp12303)
return


}


}, 1)

tmp12306 := Call(__e, PrimFunc(symshen_4_5datatype_1rule_6), V3029)


tmp12307 := Call(__e, tmp12290, tmp12306)


__e.TailApply(tmp12270, tmp12307)
return


}, 1)

tmp12308 := Call(__e, ns2_1set, symshen_4_5datatype_1rules_6, tmp12269)


_ = tmp12308

tmp12309 := MakeNative(func(__e *ControlFlow) {
V3041 := __e.Get(1)
_ = V3041
tmp12310 := MakeNative(func(__e *ControlFlow) {
W3042 := __e.Get(1)
_ = W3042
tmp12324 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3042)


if True == tmp12324 {
tmp12311 := MakeNative(func(__e *ControlFlow) {
W3046 := __e.Get(1)
_ = W3046
tmp12313 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3046)


if True == tmp12313 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3046)
return
}


}, 1)

tmp12314 := MakeNative(func(__e *ControlFlow) {
W3047 := __e.Get(1)
_ = W3047
tmp12320 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3047)


if True == tmp12320 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12315 := MakeNative(func(__e *ControlFlow) {
W3048 := __e.Get(1)
_ = W3048
tmp12316 := MakeNative(func(__e *ControlFlow) {
W3049 := __e.Get(1)
_ = W3049
__e.TailApply(PrimFunc(symshen_4comb), W3049, W3048)
return
}, 1)

tmp12317 := Call(__e, PrimFunc(symshen_4in_1_6), W3047)


__e.TailApply(tmp12316, tmp12317)
return


}, 1)

tmp12318 := Call(__e, PrimFunc(symshen_4_5_1out), W3047)


__e.TailApply(tmp12315, tmp12318)
return


}


}, 1)

tmp12321 := Call(__e, PrimFunc(symshen_4_5double_6), V3041)


tmp12322 := Call(__e, tmp12314, tmp12321)


__e.TailApply(tmp12311, tmp12322)
return


} else {
__e.Return(W3042)
return
}


}, 1)

tmp12325 := MakeNative(func(__e *ControlFlow) {
W3043 := __e.Get(1)
_ = W3043
tmp12331 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3043)


if True == tmp12331 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12326 := MakeNative(func(__e *ControlFlow) {
W3044 := __e.Get(1)
_ = W3044
tmp12327 := MakeNative(func(__e *ControlFlow) {
W3045 := __e.Get(1)
_ = W3045
__e.TailApply(PrimFunc(symshen_4comb), W3045, W3044)
return
}, 1)

tmp12328 := Call(__e, PrimFunc(symshen_4in_1_6), W3043)


__e.TailApply(tmp12327, tmp12328)
return


}, 1)

tmp12329 := Call(__e, PrimFunc(symshen_4_5_1out), W3043)


__e.TailApply(tmp12326, tmp12329)
return


}


}, 1)

tmp12332 := Call(__e, PrimFunc(symshen_4_5single_6), V3041)


tmp12333 := Call(__e, tmp12325, tmp12332)


__e.TailApply(tmp12310, tmp12333)
return


}, 1)

tmp12334 := Call(__e, ns2_1set, symshen_4_5datatype_1rule_6, tmp12309)


_ = tmp12334

tmp12335 := MakeNative(func(__e *ControlFlow) {
V3050 := __e.Get(1)
_ = V3050
tmp12336 := MakeNative(func(__e *ControlFlow) {
W3051 := __e.Get(1)
_ = W3051
tmp12338 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3051)


if True == tmp12338 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3051)
return
}


}, 1)

tmp12339 := MakeNative(func(__e *ControlFlow) {
W3052 := __e.Get(1)
_ = W3052
tmp12377 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3052)


if True == tmp12377 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12340 := MakeNative(func(__e *ControlFlow) {
W3053 := __e.Get(1)
_ = W3053
tmp12341 := MakeNative(func(__e *ControlFlow) {
W3054 := __e.Get(1)
_ = W3054
tmp12342 := MakeNative(func(__e *ControlFlow) {
W3055 := __e.Get(1)
_ = W3055
tmp12372 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3055)


if True == tmp12372 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12343 := MakeNative(func(__e *ControlFlow) {
W3056 := __e.Get(1)
_ = W3056
tmp12344 := MakeNative(func(__e *ControlFlow) {
W3057 := __e.Get(1)
_ = W3057
tmp12345 := MakeNative(func(__e *ControlFlow) {
W3058 := __e.Get(1)
_ = W3058
tmp12367 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3058)


if True == tmp12367 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12346 := MakeNative(func(__e *ControlFlow) {
W3059 := __e.Get(1)
_ = W3059
tmp12347 := MakeNative(func(__e *ControlFlow) {
W3060 := __e.Get(1)
_ = W3060
tmp12363 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3060)


if True == tmp12363 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12348 := MakeNative(func(__e *ControlFlow) {
W3061 := __e.Get(1)
_ = W3061
tmp12349 := MakeNative(func(__e *ControlFlow) {
W3062 := __e.Get(1)
_ = W3062
tmp12350 := MakeNative(func(__e *ControlFlow) {
W3063 := __e.Get(1)
_ = W3063
tmp12358 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3063)


if True == tmp12358 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12351 := MakeNative(func(__e *ControlFlow) {
W3064 := __e.Get(1)
_ = W3064
tmp12352 := PrimCons(W3061, Nil)

tmp12353 := PrimCons(W3056, tmp12352)

tmp12354 := PrimCons(W3053, tmp12353)

tmp12355 := PrimCons(tmp12354, Nil)

__e.TailApply(PrimFunc(symshen_4comb), W3064, tmp12355)
return


}, 1)

tmp12356 := Call(__e, PrimFunc(symshen_4in_1_6), W3063)


__e.TailApply(tmp12351, tmp12356)
return


}


}, 1)

tmp12359 := Call(__e, PrimFunc(symshen_4_5sc_6), W3062)


__e.TailApply(tmp12350, tmp12359)
return


}, 1)

tmp12360 := Call(__e, PrimFunc(symshen_4in_1_6), W3060)


__e.TailApply(tmp12349, tmp12360)
return


}, 1)

tmp12361 := Call(__e, PrimFunc(symshen_4_5_1out), W3060)


__e.TailApply(tmp12348, tmp12361)
return


}


}, 1)

tmp12364 := Call(__e, PrimFunc(symshen_4_5conc_6), W3059)


__e.TailApply(tmp12347, tmp12364)
return


}, 1)

tmp12365 := Call(__e, PrimFunc(symshen_4in_1_6), W3058)


__e.TailApply(tmp12346, tmp12365)
return


}


}, 1)

tmp12368 := Call(__e, PrimFunc(symshen_4_5sng_6), W3057)


__e.TailApply(tmp12345, tmp12368)
return


}, 1)

tmp12369 := Call(__e, PrimFunc(symshen_4in_1_6), W3055)


__e.TailApply(tmp12344, tmp12369)
return


}, 1)

tmp12370 := Call(__e, PrimFunc(symshen_4_5_1out), W3055)


__e.TailApply(tmp12343, tmp12370)
return


}


}, 1)

tmp12373 := Call(__e, PrimFunc(symshen_4_5prems_6), W3054)


__e.TailApply(tmp12342, tmp12373)
return


}, 1)

tmp12374 := Call(__e, PrimFunc(symshen_4in_1_6), W3052)


__e.TailApply(tmp12341, tmp12374)
return


}, 1)

tmp12375 := Call(__e, PrimFunc(symshen_4_5_1out), W3052)


__e.TailApply(tmp12340, tmp12375)
return


}


}, 1)

tmp12378 := Call(__e, PrimFunc(symshen_4_5sides_6), V3050)


tmp12379 := Call(__e, tmp12339, tmp12378)


__e.TailApply(tmp12336, tmp12379)
return


}, 1)

tmp12380 := Call(__e, ns2_1set, symshen_4_5single_6, tmp12335)


_ = tmp12380

tmp12381 := MakeNative(func(__e *ControlFlow) {
V3065 := __e.Get(1)
_ = V3065
tmp12382 := MakeNative(func(__e *ControlFlow) {
W3066 := __e.Get(1)
_ = W3066
tmp12384 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3066)


if True == tmp12384 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3066)
return
}


}, 1)

tmp12385 := MakeNative(func(__e *ControlFlow) {
W3067 := __e.Get(1)
_ = W3067
tmp12422 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3067)


if True == tmp12422 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12386 := MakeNative(func(__e *ControlFlow) {
W3068 := __e.Get(1)
_ = W3068
tmp12387 := MakeNative(func(__e *ControlFlow) {
W3069 := __e.Get(1)
_ = W3069
tmp12388 := MakeNative(func(__e *ControlFlow) {
W3070 := __e.Get(1)
_ = W3070
tmp12417 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3070)


if True == tmp12417 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12389 := MakeNative(func(__e *ControlFlow) {
W3071 := __e.Get(1)
_ = W3071
tmp12390 := MakeNative(func(__e *ControlFlow) {
W3072 := __e.Get(1)
_ = W3072
tmp12391 := MakeNative(func(__e *ControlFlow) {
W3073 := __e.Get(1)
_ = W3073
tmp12412 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3073)


if True == tmp12412 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12392 := MakeNative(func(__e *ControlFlow) {
W3074 := __e.Get(1)
_ = W3074
tmp12393 := MakeNative(func(__e *ControlFlow) {
W3075 := __e.Get(1)
_ = W3075
tmp12408 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3075)


if True == tmp12408 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12394 := MakeNative(func(__e *ControlFlow) {
W3076 := __e.Get(1)
_ = W3076
tmp12395 := MakeNative(func(__e *ControlFlow) {
W3077 := __e.Get(1)
_ = W3077
tmp12396 := MakeNative(func(__e *ControlFlow) {
W3078 := __e.Get(1)
_ = W3078
tmp12403 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3078)


if True == tmp12403 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12397 := MakeNative(func(__e *ControlFlow) {
W3079 := __e.Get(1)
_ = W3079
tmp12398 := PrimCons(W3076, Nil)

tmp12399 := PrimCons(Nil, tmp12398)

tmp12400 := Call(__e, PrimFunc(symshen_4lr_1rule), W3068, W3071, tmp12399)


__e.TailApply(PrimFunc(symshen_4comb), W3079, tmp12400)
return


}, 1)

tmp12401 := Call(__e, PrimFunc(symshen_4in_1_6), W3078)


__e.TailApply(tmp12397, tmp12401)
return


}


}, 1)

tmp12404 := Call(__e, PrimFunc(symshen_4_5sc_6), W3077)


__e.TailApply(tmp12396, tmp12404)
return


}, 1)

tmp12405 := Call(__e, PrimFunc(symshen_4in_1_6), W3075)


__e.TailApply(tmp12395, tmp12405)
return


}, 1)

tmp12406 := Call(__e, PrimFunc(symshen_4_5_1out), W3075)


__e.TailApply(tmp12394, tmp12406)
return


}


}, 1)

tmp12409 := Call(__e, PrimFunc(symshen_4_5formula_6), W3074)


__e.TailApply(tmp12393, tmp12409)
return


}, 1)

tmp12410 := Call(__e, PrimFunc(symshen_4in_1_6), W3073)


__e.TailApply(tmp12392, tmp12410)
return


}


}, 1)

tmp12413 := Call(__e, PrimFunc(symshen_4_5dbl_6), W3072)


__e.TailApply(tmp12391, tmp12413)
return


}, 1)

tmp12414 := Call(__e, PrimFunc(symshen_4in_1_6), W3070)


__e.TailApply(tmp12390, tmp12414)
return


}, 1)

tmp12415 := Call(__e, PrimFunc(symshen_4_5_1out), W3070)


__e.TailApply(tmp12389, tmp12415)
return


}


}, 1)

tmp12418 := Call(__e, PrimFunc(symshen_4_5formulae_6), W3069)


__e.TailApply(tmp12388, tmp12418)
return


}, 1)

tmp12419 := Call(__e, PrimFunc(symshen_4in_1_6), W3067)


__e.TailApply(tmp12387, tmp12419)
return


}, 1)

tmp12420 := Call(__e, PrimFunc(symshen_4_5_1out), W3067)


__e.TailApply(tmp12386, tmp12420)
return


}


}, 1)

tmp12423 := Call(__e, PrimFunc(symshen_4_5sides_6), V3065)


tmp12424 := Call(__e, tmp12385, tmp12423)


__e.TailApply(tmp12382, tmp12424)
return


}, 1)

tmp12425 := Call(__e, ns2_1set, symshen_4_5double_6, tmp12381)


_ = tmp12425

tmp12426 := MakeNative(func(__e *ControlFlow) {
V3080 := __e.Get(1)
_ = V3080
tmp12427 := MakeNative(func(__e *ControlFlow) {
W3081 := __e.Get(1)
_ = W3081
tmp12450 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3081)


if True == tmp12450 {
tmp12428 := MakeNative(func(__e *ControlFlow) {
W3090 := __e.Get(1)
_ = W3090
tmp12430 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3090)


if True == tmp12430 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3090)
return
}


}, 1)

tmp12431 := MakeNative(func(__e *ControlFlow) {
W3091 := __e.Get(1)
_ = W3091
tmp12446 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3091)


if True == tmp12446 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12432 := MakeNative(func(__e *ControlFlow) {
W3092 := __e.Get(1)
_ = W3092
tmp12433 := MakeNative(func(__e *ControlFlow) {
W3093 := __e.Get(1)
_ = W3093
tmp12434 := MakeNative(func(__e *ControlFlow) {
W3094 := __e.Get(1)
_ = W3094
tmp12441 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3094)


if True == tmp12441 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12435 := MakeNative(func(__e *ControlFlow) {
W3095 := __e.Get(1)
_ = W3095
tmp12436 := PrimCons(W3092, Nil)

tmp12437 := PrimCons(Nil, tmp12436)

tmp12438 := PrimCons(tmp12437, Nil)

__e.TailApply(PrimFunc(symshen_4comb), W3095, tmp12438)
return


}, 1)

tmp12439 := Call(__e, PrimFunc(symshen_4in_1_6), W3094)


__e.TailApply(tmp12435, tmp12439)
return


}


}, 1)

tmp12442 := Call(__e, PrimFunc(symshen_4_5sc_6), W3093)


__e.TailApply(tmp12434, tmp12442)
return


}, 1)

tmp12443 := Call(__e, PrimFunc(symshen_4in_1_6), W3091)


__e.TailApply(tmp12433, tmp12443)
return


}, 1)

tmp12444 := Call(__e, PrimFunc(symshen_4_5_1out), W3091)


__e.TailApply(tmp12432, tmp12444)
return


}


}, 1)

tmp12447 := Call(__e, PrimFunc(symshen_4_5formula_6), V3080)


tmp12448 := Call(__e, tmp12431, tmp12447)


__e.TailApply(tmp12428, tmp12448)
return


} else {
__e.Return(W3081)
return
}


}, 1)

tmp12451 := MakeNative(func(__e *ControlFlow) {
W3082 := __e.Get(1)
_ = W3082
tmp12474 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3082)


if True == tmp12474 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12452 := MakeNative(func(__e *ControlFlow) {
W3083 := __e.Get(1)
_ = W3083
tmp12453 := MakeNative(func(__e *ControlFlow) {
W3084 := __e.Get(1)
_ = W3084
tmp12454 := MakeNative(func(__e *ControlFlow) {
W3085 := __e.Get(1)
_ = W3085
tmp12469 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3085)


if True == tmp12469 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12455 := MakeNative(func(__e *ControlFlow) {
W3086 := __e.Get(1)
_ = W3086
tmp12456 := MakeNative(func(__e *ControlFlow) {
W3087 := __e.Get(1)
_ = W3087
tmp12465 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3087)


if True == tmp12465 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12457 := MakeNative(func(__e *ControlFlow) {
W3088 := __e.Get(1)
_ = W3088
tmp12458 := MakeNative(func(__e *ControlFlow) {
W3089 := __e.Get(1)
_ = W3089
tmp12459 := PrimCons(W3083, Nil)

tmp12460 := PrimCons(Nil, tmp12459)

tmp12461 := PrimCons(tmp12460, W3088)

__e.TailApply(PrimFunc(symshen_4comb), W3089, tmp12461)
return


}, 1)

tmp12462 := Call(__e, PrimFunc(symshen_4in_1_6), W3087)


__e.TailApply(tmp12458, tmp12462)
return


}, 1)

tmp12463 := Call(__e, PrimFunc(symshen_4_5_1out), W3087)


__e.TailApply(tmp12457, tmp12463)
return


}


}, 1)

tmp12466 := Call(__e, PrimFunc(symshen_4_5formulae_6), W3086)


__e.TailApply(tmp12456, tmp12466)
return


}, 1)

tmp12467 := Call(__e, PrimFunc(symshen_4in_1_6), W3085)


__e.TailApply(tmp12455, tmp12467)
return


}


}, 1)

tmp12470 := Call(__e, PrimFunc(symshen_4_5sc_6), W3084)


__e.TailApply(tmp12454, tmp12470)
return


}, 1)

tmp12471 := Call(__e, PrimFunc(symshen_4in_1_6), W3082)


__e.TailApply(tmp12453, tmp12471)
return


}, 1)

tmp12472 := Call(__e, PrimFunc(symshen_4_5_1out), W3082)


__e.TailApply(tmp12452, tmp12472)
return


}


}, 1)

tmp12475 := Call(__e, PrimFunc(symshen_4_5formula_6), V3080)


tmp12476 := Call(__e, tmp12451, tmp12475)


__e.TailApply(tmp12427, tmp12476)
return


}, 1)

tmp12477 := Call(__e, ns2_1set, symshen_4_5formulae_6, tmp12426)


_ = tmp12477

tmp12478 := MakeNative(func(__e *ControlFlow) {
V3096 := __e.Get(1)
_ = V3096
tmp12479 := MakeNative(func(__e *ControlFlow) {
W3097 := __e.Get(1)
_ = W3097
tmp12495 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3097)


if True == tmp12495 {
tmp12480 := MakeNative(func(__e *ControlFlow) {
W3105 := __e.Get(1)
_ = W3105
tmp12482 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3105)


if True == tmp12482 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3105)
return
}


}, 1)

tmp12483 := MakeNative(func(__e *ControlFlow) {
W3106 := __e.Get(1)
_ = W3106
tmp12491 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3106)


if True == tmp12491 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12484 := MakeNative(func(__e *ControlFlow) {
W3107 := __e.Get(1)
_ = W3107
tmp12485 := MakeNative(func(__e *ControlFlow) {
W3108 := __e.Get(1)
_ = W3108
tmp12486 := PrimCons(W3107, Nil)

tmp12487 := PrimCons(Nil, tmp12486)

__e.TailApply(PrimFunc(symshen_4comb), W3108, tmp12487)
return


}, 1)

tmp12488 := Call(__e, PrimFunc(symshen_4in_1_6), W3106)


__e.TailApply(tmp12485, tmp12488)
return


}, 1)

tmp12489 := Call(__e, PrimFunc(symshen_4_5_1out), W3106)


__e.TailApply(tmp12484, tmp12489)
return


}


}, 1)

tmp12492 := Call(__e, PrimFunc(symshen_4_5formula_6), V3096)


tmp12493 := Call(__e, tmp12483, tmp12492)


__e.TailApply(tmp12480, tmp12493)
return


} else {
__e.Return(W3097)
return
}


}, 1)

tmp12496 := MakeNative(func(__e *ControlFlow) {
W3098 := __e.Get(1)
_ = W3098
tmp12516 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3098)


if True == tmp12516 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12497 := MakeNative(func(__e *ControlFlow) {
W3099 := __e.Get(1)
_ = W3099
tmp12498 := MakeNative(func(__e *ControlFlow) {
W3100 := __e.Get(1)
_ = W3100
tmp12512 := Call(__e, PrimFunc(symshen_4hds_a_2), W3100, sym_6_6)


if True == tmp12512 {
tmp12499 := MakeNative(func(__e *ControlFlow) {
W3101 := __e.Get(1)
_ = W3101
tmp12500 := MakeNative(func(__e *ControlFlow) {
W3102 := __e.Get(1)
_ = W3102
tmp12508 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3102)


if True == tmp12508 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12501 := MakeNative(func(__e *ControlFlow) {
W3103 := __e.Get(1)
_ = W3103
tmp12502 := MakeNative(func(__e *ControlFlow) {
W3104 := __e.Get(1)
_ = W3104
tmp12503 := PrimCons(W3103, Nil)

tmp12504 := PrimCons(W3099, tmp12503)

__e.TailApply(PrimFunc(symshen_4comb), W3104, tmp12504)
return


}, 1)

tmp12505 := Call(__e, PrimFunc(symshen_4in_1_6), W3102)


__e.TailApply(tmp12502, tmp12505)
return


}, 1)

tmp12506 := Call(__e, PrimFunc(symshen_4_5_1out), W3102)


__e.TailApply(tmp12501, tmp12506)
return


}


}, 1)

tmp12509 := Call(__e, PrimFunc(symshen_4_5formula_6), W3101)


__e.TailApply(tmp12500, tmp12509)
return


}, 1)

tmp12510 := Call(__e, PrimFunc(symtail), W3100)


__e.TailApply(tmp12499, tmp12510)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp12513 := Call(__e, PrimFunc(symshen_4in_1_6), W3098)


__e.TailApply(tmp12498, tmp12513)
return


}, 1)

tmp12514 := Call(__e, PrimFunc(symshen_4_5_1out), W3098)


__e.TailApply(tmp12497, tmp12514)
return


}


}, 1)

tmp12517 := Call(__e, PrimFunc(symshen_4_5ass_6), V3096)


tmp12518 := Call(__e, tmp12496, tmp12517)


__e.TailApply(tmp12479, tmp12518)
return


}, 1)

tmp12519 := Call(__e, ns2_1set, symshen_4_5conc_6, tmp12478)


_ = tmp12519

tmp12520 := MakeNative(func(__e *ControlFlow) {
V3109 := __e.Get(1)
_ = V3109
tmp12521 := MakeNative(func(__e *ControlFlow) {
W3110 := __e.Get(1)
_ = W3110
tmp12533 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3110)


if True == tmp12533 {
tmp12522 := MakeNative(func(__e *ControlFlow) {
W3119 := __e.Get(1)
_ = W3119
tmp12524 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3119)


if True == tmp12524 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3119)
return
}


}, 1)

tmp12525 := MakeNative(func(__e *ControlFlow) {
W3120 := __e.Get(1)
_ = W3120
tmp12529 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3120)


if True == tmp12529 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12526 := MakeNative(func(__e *ControlFlow) {
W3121 := __e.Get(1)
_ = W3121
__e.TailApply(PrimFunc(symshen_4comb), W3121, Nil)
return
}, 1)

tmp12527 := Call(__e, PrimFunc(symshen_4in_1_6), W3120)


__e.TailApply(tmp12526, tmp12527)
return


}


}, 1)

tmp12530 := Call(__e, PrimFunc(sym_5e_6), V3109)


tmp12531 := Call(__e, tmp12525, tmp12530)


__e.TailApply(tmp12522, tmp12531)
return


} else {
__e.Return(W3110)
return
}


}, 1)

tmp12534 := MakeNative(func(__e *ControlFlow) {
W3111 := __e.Get(1)
_ = W3111
tmp12555 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3111)


if True == tmp12555 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12535 := MakeNative(func(__e *ControlFlow) {
W3112 := __e.Get(1)
_ = W3112
tmp12536 := MakeNative(func(__e *ControlFlow) {
W3113 := __e.Get(1)
_ = W3113
tmp12537 := MakeNative(func(__e *ControlFlow) {
W3114 := __e.Get(1)
_ = W3114
tmp12550 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3114)


if True == tmp12550 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12538 := MakeNative(func(__e *ControlFlow) {
W3115 := __e.Get(1)
_ = W3115
tmp12539 := MakeNative(func(__e *ControlFlow) {
W3116 := __e.Get(1)
_ = W3116
tmp12546 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3116)


if True == tmp12546 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12540 := MakeNative(func(__e *ControlFlow) {
W3117 := __e.Get(1)
_ = W3117
tmp12541 := MakeNative(func(__e *ControlFlow) {
W3118 := __e.Get(1)
_ = W3118
tmp12542 := PrimCons(W3112, W3117)

__e.TailApply(PrimFunc(symshen_4comb), W3118, tmp12542)
return


}, 1)

tmp12543 := Call(__e, PrimFunc(symshen_4in_1_6), W3116)


__e.TailApply(tmp12541, tmp12543)
return


}, 1)

tmp12544 := Call(__e, PrimFunc(symshen_4_5_1out), W3116)


__e.TailApply(tmp12540, tmp12544)
return


}


}, 1)

tmp12547 := Call(__e, PrimFunc(symshen_4_5prems_6), W3115)


__e.TailApply(tmp12539, tmp12547)
return


}, 1)

tmp12548 := Call(__e, PrimFunc(symshen_4in_1_6), W3114)


__e.TailApply(tmp12538, tmp12548)
return


}


}, 1)

tmp12551 := Call(__e, PrimFunc(symshen_4_5sc_6), W3113)


__e.TailApply(tmp12537, tmp12551)
return


}, 1)

tmp12552 := Call(__e, PrimFunc(symshen_4in_1_6), W3111)


__e.TailApply(tmp12536, tmp12552)
return


}, 1)

tmp12553 := Call(__e, PrimFunc(symshen_4_5_1out), W3111)


__e.TailApply(tmp12535, tmp12553)
return


}


}, 1)

tmp12556 := Call(__e, PrimFunc(symshen_4_5prem_6), V3109)


tmp12557 := Call(__e, tmp12534, tmp12556)


__e.TailApply(tmp12521, tmp12557)
return


}, 1)

tmp12558 := Call(__e, ns2_1set, symshen_4_5prems_6, tmp12520)


_ = tmp12558

tmp12559 := MakeNative(func(__e *ControlFlow) {
V3122 := __e.Get(1)
_ = V3122
tmp12560 := MakeNative(func(__e *ControlFlow) {
W3123 := __e.Get(1)
_ = W3123
tmp12602 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3123)


if True == tmp12602 {
tmp12561 := MakeNative(func(__e *ControlFlow) {
W3125 := __e.Get(1)
_ = W3125
tmp12577 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3125)


if True == tmp12577 {
tmp12562 := MakeNative(func(__e *ControlFlow) {
W3133 := __e.Get(1)
_ = W3133
tmp12564 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3133)


if True == tmp12564 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3133)
return
}


}, 1)

tmp12565 := MakeNative(func(__e *ControlFlow) {
W3134 := __e.Get(1)
_ = W3134
tmp12573 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3134)


if True == tmp12573 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12566 := MakeNative(func(__e *ControlFlow) {
W3135 := __e.Get(1)
_ = W3135
tmp12567 := MakeNative(func(__e *ControlFlow) {
W3136 := __e.Get(1)
_ = W3136
tmp12568 := PrimCons(W3135, Nil)

tmp12569 := PrimCons(Nil, tmp12568)

__e.TailApply(PrimFunc(symshen_4comb), W3136, tmp12569)
return


}, 1)

tmp12570 := Call(__e, PrimFunc(symshen_4in_1_6), W3134)


__e.TailApply(tmp12567, tmp12570)
return


}, 1)

tmp12571 := Call(__e, PrimFunc(symshen_4_5_1out), W3134)


__e.TailApply(tmp12566, tmp12571)
return


}


}, 1)

tmp12574 := Call(__e, PrimFunc(symshen_4_5formula_6), V3122)


tmp12575 := Call(__e, tmp12565, tmp12574)


__e.TailApply(tmp12562, tmp12575)
return


} else {
__e.Return(W3125)
return
}


}, 1)

tmp12578 := MakeNative(func(__e *ControlFlow) {
W3126 := __e.Get(1)
_ = W3126
tmp12598 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3126)


if True == tmp12598 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12579 := MakeNative(func(__e *ControlFlow) {
W3127 := __e.Get(1)
_ = W3127
tmp12580 := MakeNative(func(__e *ControlFlow) {
W3128 := __e.Get(1)
_ = W3128
tmp12594 := Call(__e, PrimFunc(symshen_4hds_a_2), W3128, sym_6_6)


if True == tmp12594 {
tmp12581 := MakeNative(func(__e *ControlFlow) {
W3129 := __e.Get(1)
_ = W3129
tmp12582 := MakeNative(func(__e *ControlFlow) {
W3130 := __e.Get(1)
_ = W3130
tmp12590 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3130)


if True == tmp12590 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12583 := MakeNative(func(__e *ControlFlow) {
W3131 := __e.Get(1)
_ = W3131
tmp12584 := MakeNative(func(__e *ControlFlow) {
W3132 := __e.Get(1)
_ = W3132
tmp12585 := PrimCons(W3131, Nil)

tmp12586 := PrimCons(W3127, tmp12585)

__e.TailApply(PrimFunc(symshen_4comb), W3132, tmp12586)
return


}, 1)

tmp12587 := Call(__e, PrimFunc(symshen_4in_1_6), W3130)


__e.TailApply(tmp12584, tmp12587)
return


}, 1)

tmp12588 := Call(__e, PrimFunc(symshen_4_5_1out), W3130)


__e.TailApply(tmp12583, tmp12588)
return


}


}, 1)

tmp12591 := Call(__e, PrimFunc(symshen_4_5formula_6), W3129)


__e.TailApply(tmp12582, tmp12591)
return


}, 1)

tmp12592 := Call(__e, PrimFunc(symtail), W3128)


__e.TailApply(tmp12581, tmp12592)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp12595 := Call(__e, PrimFunc(symshen_4in_1_6), W3126)


__e.TailApply(tmp12580, tmp12595)
return


}, 1)

tmp12596 := Call(__e, PrimFunc(symshen_4_5_1out), W3126)


__e.TailApply(tmp12579, tmp12596)
return


}


}, 1)

tmp12599 := Call(__e, PrimFunc(symshen_4_5ass_6), V3122)


tmp12600 := Call(__e, tmp12578, tmp12599)


__e.TailApply(tmp12561, tmp12600)
return


} else {
__e.Return(W3123)
return
}


}, 1)

tmp12608 := Call(__e, PrimFunc(symshen_4hds_a_2), V3122, sym_b)


var ifres12603 Obj

if True == tmp12608 {
tmp12604 := MakeNative(func(__e *ControlFlow) {
W3124 := __e.Get(1)
_ = W3124
__e.TailApply(PrimFunc(symshen_4comb), W3124, sym_b)
return
}, 1)

tmp12605 := Call(__e, PrimFunc(symtail), V3122)


tmp12606 := Call(__e, tmp12604, tmp12605)


ifres12603 = tmp12606


} else {
tmp12607 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12603 = tmp12607


}

__e.TailApply(tmp12560, ifres12603)
return


}, 1)

tmp12609 := Call(__e, ns2_1set, symshen_4_5prem_6, tmp12559)


_ = tmp12609

tmp12610 := MakeNative(func(__e *ControlFlow) {
V3137 := __e.Get(1)
_ = V3137
tmp12611 := MakeNative(func(__e *ControlFlow) {
W3138 := __e.Get(1)
_ = W3138
tmp12636 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3138)


if True == tmp12636 {
tmp12612 := MakeNative(func(__e *ControlFlow) {
W3147 := __e.Get(1)
_ = W3147
tmp12624 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3147)


if True == tmp12624 {
tmp12613 := MakeNative(func(__e *ControlFlow) {
W3151 := __e.Get(1)
_ = W3151
tmp12615 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3151)


if True == tmp12615 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3151)
return
}


}, 1)

tmp12616 := MakeNative(func(__e *ControlFlow) {
W3152 := __e.Get(1)
_ = W3152
tmp12620 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3152)


if True == tmp12620 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12617 := MakeNative(func(__e *ControlFlow) {
W3153 := __e.Get(1)
_ = W3153
__e.TailApply(PrimFunc(symshen_4comb), W3153, Nil)
return
}, 1)

tmp12618 := Call(__e, PrimFunc(symshen_4in_1_6), W3152)


__e.TailApply(tmp12617, tmp12618)
return


}


}, 1)

tmp12621 := Call(__e, PrimFunc(sym_5e_6), V3137)


tmp12622 := Call(__e, tmp12616, tmp12621)


__e.TailApply(tmp12613, tmp12622)
return


} else {
__e.Return(W3147)
return
}


}, 1)

tmp12625 := MakeNative(func(__e *ControlFlow) {
W3148 := __e.Get(1)
_ = W3148
tmp12632 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3148)


if True == tmp12632 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12626 := MakeNative(func(__e *ControlFlow) {
W3149 := __e.Get(1)
_ = W3149
tmp12627 := MakeNative(func(__e *ControlFlow) {
W3150 := __e.Get(1)
_ = W3150
tmp12628 := PrimCons(W3149, Nil)

__e.TailApply(PrimFunc(symshen_4comb), W3150, tmp12628)
return


}, 1)

tmp12629 := Call(__e, PrimFunc(symshen_4in_1_6), W3148)


__e.TailApply(tmp12627, tmp12629)
return


}, 1)

tmp12630 := Call(__e, PrimFunc(symshen_4_5_1out), W3148)


__e.TailApply(tmp12626, tmp12630)
return


}


}, 1)

tmp12633 := Call(__e, PrimFunc(symshen_4_5formula_6), V3137)


tmp12634 := Call(__e, tmp12625, tmp12633)


__e.TailApply(tmp12612, tmp12634)
return


} else {
__e.Return(W3138)
return
}


}, 1)

tmp12637 := MakeNative(func(__e *ControlFlow) {
W3139 := __e.Get(1)
_ = W3139
tmp12658 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3139)


if True == tmp12658 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12638 := MakeNative(func(__e *ControlFlow) {
W3140 := __e.Get(1)
_ = W3140
tmp12639 := MakeNative(func(__e *ControlFlow) {
W3141 := __e.Get(1)
_ = W3141
tmp12640 := MakeNative(func(__e *ControlFlow) {
W3142 := __e.Get(1)
_ = W3142
tmp12653 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3142)


if True == tmp12653 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12641 := MakeNative(func(__e *ControlFlow) {
W3143 := __e.Get(1)
_ = W3143
tmp12642 := MakeNative(func(__e *ControlFlow) {
W3144 := __e.Get(1)
_ = W3144
tmp12649 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3144)


if True == tmp12649 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12643 := MakeNative(func(__e *ControlFlow) {
W3145 := __e.Get(1)
_ = W3145
tmp12644 := MakeNative(func(__e *ControlFlow) {
W3146 := __e.Get(1)
_ = W3146
tmp12645 := PrimCons(W3140, W3145)

__e.TailApply(PrimFunc(symshen_4comb), W3146, tmp12645)
return


}, 1)

tmp12646 := Call(__e, PrimFunc(symshen_4in_1_6), W3144)


__e.TailApply(tmp12644, tmp12646)
return


}, 1)

tmp12647 := Call(__e, PrimFunc(symshen_4_5_1out), W3144)


__e.TailApply(tmp12643, tmp12647)
return


}


}, 1)

tmp12650 := Call(__e, PrimFunc(symshen_4_5ass_6), W3143)


__e.TailApply(tmp12642, tmp12650)
return


}, 1)

tmp12651 := Call(__e, PrimFunc(symshen_4in_1_6), W3142)


__e.TailApply(tmp12641, tmp12651)
return


}


}, 1)

tmp12654 := Call(__e, PrimFunc(symshen_4_5iscomma_6), W3141)


__e.TailApply(tmp12640, tmp12654)
return


}, 1)

tmp12655 := Call(__e, PrimFunc(symshen_4in_1_6), W3139)


__e.TailApply(tmp12639, tmp12655)
return


}, 1)

tmp12656 := Call(__e, PrimFunc(symshen_4_5_1out), W3139)


__e.TailApply(tmp12638, tmp12656)
return


}


}, 1)

tmp12659 := Call(__e, PrimFunc(symshen_4_5formula_6), V3137)


tmp12660 := Call(__e, tmp12637, tmp12659)


__e.TailApply(tmp12611, tmp12660)
return


}, 1)

tmp12661 := Call(__e, ns2_1set, symshen_4_5ass_6, tmp12610)


_ = tmp12661

tmp12662 := MakeNative(func(__e *ControlFlow) {
V3154 := __e.Get(1)
_ = V3154
tmp12663 := MakeNative(func(__e *ControlFlow) {
W3155 := __e.Get(1)
_ = W3155
tmp12665 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3155)


if True == tmp12665 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3155)
return
}


}, 1)

tmp12676 := PrimIsPair(V3154)

var ifres12666 Obj

if True == tmp12676 {
tmp12667 := MakeNative(func(__e *ControlFlow) {
W3156 := __e.Get(1)
_ = W3156
tmp12668 := MakeNative(func(__e *ControlFlow) {
W3157 := __e.Get(1)
_ = W3157
tmp12670 := PrimIntern(MakeString(","))

tmp12671 := PrimEqual(W3156, tmp12670)

if True == tmp12671 {
__e.TailApply(PrimFunc(symshen_4comb), W3157, symshen_4skip)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp12672 := Call(__e, PrimFunc(symtail), V3154)


__e.TailApply(tmp12668, tmp12672)
return


}, 1)

tmp12673 := Call(__e, PrimFunc(symhead), V3154)


tmp12674 := Call(__e, tmp12667, tmp12673)


ifres12666 = tmp12674


} else {
tmp12675 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12666 = tmp12675


}

__e.TailApply(tmp12663, ifres12666)
return


}, 1)

tmp12677 := Call(__e, ns2_1set, symshen_4_5iscomma_6, tmp12662)


_ = tmp12677

tmp12678 := MakeNative(func(__e *ControlFlow) {
V3158 := __e.Get(1)
_ = V3158
tmp12679 := MakeNative(func(__e *ControlFlow) {
W3159 := __e.Get(1)
_ = W3159
tmp12693 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3159)


if True == tmp12693 {
tmp12680 := MakeNative(func(__e *ControlFlow) {
W3168 := __e.Get(1)
_ = W3168
tmp12682 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3168)


if True == tmp12682 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3168)
return
}


}, 1)

tmp12683 := MakeNative(func(__e *ControlFlow) {
W3169 := __e.Get(1)
_ = W3169
tmp12689 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3169)


if True == tmp12689 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12684 := MakeNative(func(__e *ControlFlow) {
W3170 := __e.Get(1)
_ = W3170
tmp12685 := MakeNative(func(__e *ControlFlow) {
W3171 := __e.Get(1)
_ = W3171
__e.TailApply(PrimFunc(symshen_4comb), W3171, W3170)
return
}, 1)

tmp12686 := Call(__e, PrimFunc(symshen_4in_1_6), W3169)


__e.TailApply(tmp12685, tmp12686)
return


}, 1)

tmp12687 := Call(__e, PrimFunc(symshen_4_5_1out), W3169)


__e.TailApply(tmp12684, tmp12687)
return


}


}, 1)

tmp12690 := Call(__e, PrimFunc(symshen_4_5expr_6), V3158)


tmp12691 := Call(__e, tmp12683, tmp12690)


__e.TailApply(tmp12680, tmp12691)
return


} else {
__e.Return(W3159)
return
}


}, 1)

tmp12694 := MakeNative(func(__e *ControlFlow) {
W3160 := __e.Get(1)
_ = W3160
tmp12720 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3160)


if True == tmp12720 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12695 := MakeNative(func(__e *ControlFlow) {
W3161 := __e.Get(1)
_ = W3161
tmp12696 := MakeNative(func(__e *ControlFlow) {
W3162 := __e.Get(1)
_ = W3162
tmp12697 := MakeNative(func(__e *ControlFlow) {
W3163 := __e.Get(1)
_ = W3163
tmp12715 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3163)


if True == tmp12715 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12698 := MakeNative(func(__e *ControlFlow) {
W3164 := __e.Get(1)
_ = W3164
tmp12699 := MakeNative(func(__e *ControlFlow) {
W3165 := __e.Get(1)
_ = W3165
tmp12711 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3165)


if True == tmp12711 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12700 := MakeNative(func(__e *ControlFlow) {
W3166 := __e.Get(1)
_ = W3166
tmp12701 := MakeNative(func(__e *ControlFlow) {
W3167 := __e.Get(1)
_ = W3167
tmp12702 := Call(__e, PrimFunc(symshen_4curry), W3161)


tmp12703 := PrimIntern(MakeString(":"))

tmp12704 := Call(__e, PrimFunc(symshen_4rectify_1type), W3166)


tmp12705 := PrimCons(tmp12704, Nil)

tmp12706 := PrimCons(tmp12703, tmp12705)

tmp12707 := PrimCons(tmp12702, tmp12706)

__e.TailApply(PrimFunc(symshen_4comb), W3167, tmp12707)
return


}, 1)

tmp12708 := Call(__e, PrimFunc(symshen_4in_1_6), W3165)


__e.TailApply(tmp12701, tmp12708)
return


}, 1)

tmp12709 := Call(__e, PrimFunc(symshen_4_5_1out), W3165)


__e.TailApply(tmp12700, tmp12709)
return


}


}, 1)

tmp12712 := Call(__e, PrimFunc(symshen_4_5type_6), W3164)


__e.TailApply(tmp12699, tmp12712)
return


}, 1)

tmp12713 := Call(__e, PrimFunc(symshen_4in_1_6), W3163)


__e.TailApply(tmp12698, tmp12713)
return


}


}, 1)

tmp12716 := Call(__e, PrimFunc(symshen_4_5iscolon_6), W3162)


__e.TailApply(tmp12697, tmp12716)
return


}, 1)

tmp12717 := Call(__e, PrimFunc(symshen_4in_1_6), W3160)


__e.TailApply(tmp12696, tmp12717)
return


}, 1)

tmp12718 := Call(__e, PrimFunc(symshen_4_5_1out), W3160)


__e.TailApply(tmp12695, tmp12718)
return


}


}, 1)

tmp12721 := Call(__e, PrimFunc(symshen_4_5expr_6), V3158)


tmp12722 := Call(__e, tmp12694, tmp12721)


__e.TailApply(tmp12679, tmp12722)
return


}, 1)

tmp12723 := Call(__e, ns2_1set, symshen_4_5formula_6, tmp12678)


_ = tmp12723

tmp12724 := MakeNative(func(__e *ControlFlow) {
V3172 := __e.Get(1)
_ = V3172
tmp12725 := MakeNative(func(__e *ControlFlow) {
W3173 := __e.Get(1)
_ = W3173
tmp12727 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3173)


if True == tmp12727 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3173)
return
}


}, 1)

tmp12738 := PrimIsPair(V3172)

var ifres12728 Obj

if True == tmp12738 {
tmp12729 := MakeNative(func(__e *ControlFlow) {
W3174 := __e.Get(1)
_ = W3174
tmp12730 := MakeNative(func(__e *ControlFlow) {
W3175 := __e.Get(1)
_ = W3175
tmp12732 := PrimIntern(MakeString(":"))

tmp12733 := PrimEqual(W3174, tmp12732)

if True == tmp12733 {
__e.TailApply(PrimFunc(symshen_4comb), W3175, symshen_4skip)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp12734 := Call(__e, PrimFunc(symtail), V3172)


__e.TailApply(tmp12730, tmp12734)
return


}, 1)

tmp12735 := Call(__e, PrimFunc(symhead), V3172)


tmp12736 := Call(__e, tmp12729, tmp12735)


ifres12728 = tmp12736


} else {
tmp12737 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12728 = tmp12737


}

__e.TailApply(tmp12725, ifres12728)
return


}, 1)

tmp12739 := Call(__e, ns2_1set, symshen_4_5iscolon_6, tmp12724)


_ = tmp12739

tmp12740 := MakeNative(func(__e *ControlFlow) {
V3176 := __e.Get(1)
_ = V3176
tmp12741 := MakeNative(func(__e *ControlFlow) {
W3177 := __e.Get(1)
_ = W3177
tmp12753 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3177)


if True == tmp12753 {
tmp12742 := MakeNative(func(__e *ControlFlow) {
W3184 := __e.Get(1)
_ = W3184
tmp12744 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3184)


if True == tmp12744 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3184)
return
}


}, 1)

tmp12745 := MakeNative(func(__e *ControlFlow) {
W3185 := __e.Get(1)
_ = W3185
tmp12749 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3185)


if True == tmp12749 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12746 := MakeNative(func(__e *ControlFlow) {
W3186 := __e.Get(1)
_ = W3186
__e.TailApply(PrimFunc(symshen_4comb), W3186, Nil)
return
}, 1)

tmp12747 := Call(__e, PrimFunc(symshen_4in_1_6), W3185)


__e.TailApply(tmp12746, tmp12747)
return


}


}, 1)

tmp12750 := Call(__e, PrimFunc(sym_5e_6), V3176)


tmp12751 := Call(__e, tmp12745, tmp12750)


__e.TailApply(tmp12742, tmp12751)
return


} else {
__e.Return(W3177)
return
}


}, 1)

tmp12754 := MakeNative(func(__e *ControlFlow) {
W3178 := __e.Get(1)
_ = W3178
tmp12769 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3178)


if True == tmp12769 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12755 := MakeNative(func(__e *ControlFlow) {
W3179 := __e.Get(1)
_ = W3179
tmp12756 := MakeNative(func(__e *ControlFlow) {
W3180 := __e.Get(1)
_ = W3180
tmp12757 := MakeNative(func(__e *ControlFlow) {
W3181 := __e.Get(1)
_ = W3181
tmp12764 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3181)


if True == tmp12764 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12758 := MakeNative(func(__e *ControlFlow) {
W3182 := __e.Get(1)
_ = W3182
tmp12759 := MakeNative(func(__e *ControlFlow) {
W3183 := __e.Get(1)
_ = W3183
tmp12760 := PrimCons(W3179, W3182)

__e.TailApply(PrimFunc(symshen_4comb), W3183, tmp12760)
return


}, 1)

tmp12761 := Call(__e, PrimFunc(symshen_4in_1_6), W3181)


__e.TailApply(tmp12759, tmp12761)
return


}, 1)

tmp12762 := Call(__e, PrimFunc(symshen_4_5_1out), W3181)


__e.TailApply(tmp12758, tmp12762)
return


}


}, 1)

tmp12765 := Call(__e, PrimFunc(symshen_4_5sides_6), W3180)


__e.TailApply(tmp12757, tmp12765)
return


}, 1)

tmp12766 := Call(__e, PrimFunc(symshen_4in_1_6), W3178)


__e.TailApply(tmp12756, tmp12766)
return


}, 1)

tmp12767 := Call(__e, PrimFunc(symshen_4_5_1out), W3178)


__e.TailApply(tmp12755, tmp12767)
return


}


}, 1)

tmp12770 := Call(__e, PrimFunc(symshen_4_5side_6), V3176)


tmp12771 := Call(__e, tmp12754, tmp12770)


__e.TailApply(tmp12741, tmp12771)
return


}, 1)

tmp12772 := Call(__e, ns2_1set, symshen_4_5sides_6, tmp12740)


_ = tmp12772

tmp12773 := MakeNative(func(__e *ControlFlow) {
V3187 := __e.Get(1)
_ = V3187
tmp12774 := MakeNative(func(__e *ControlFlow) {
W3188 := __e.Get(1)
_ = W3188
tmp12838 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3188)


if True == tmp12838 {
tmp12775 := MakeNative(func(__e *ControlFlow) {
W3192 := __e.Get(1)
_ = W3192
tmp12815 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3192)


if True == tmp12815 {
tmp12776 := MakeNative(func(__e *ControlFlow) {
W3198 := __e.Get(1)
_ = W3198
tmp12797 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3198)


if True == tmp12797 {
tmp12777 := MakeNative(func(__e *ControlFlow) {
W3202 := __e.Get(1)
_ = W3202
tmp12779 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3202)


if True == tmp12779 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3202)
return
}


}, 1)

tmp12795 := Call(__e, PrimFunc(symshen_4hds_a_2), V3187, symsqts)


var ifres12780 Obj

if True == tmp12795 {
tmp12781 := MakeNative(func(__e *ControlFlow) {
W3203 := __e.Get(1)
_ = W3203
tmp12791 := PrimIsPair(W3203)

if True == tmp12791 {
tmp12782 := MakeNative(func(__e *ControlFlow) {
W3204 := __e.Get(1)
_ = W3204
tmp12783 := MakeNative(func(__e *ControlFlow) {
W3205 := __e.Get(1)
_ = W3205
tmp12787 := PrimIsVariable(W3204)

if True == tmp12787 {
tmp12784 := PrimCons(W3204, Nil)

tmp12785 := PrimCons(symsqts, tmp12784)

__e.TailApply(PrimFunc(symshen_4comb), W3205, tmp12785)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp12788 := Call(__e, PrimFunc(symtail), W3203)


__e.TailApply(tmp12783, tmp12788)
return


}, 1)

tmp12789 := Call(__e, PrimFunc(symhead), W3203)


__e.TailApply(tmp12782, tmp12789)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp12792 := Call(__e, PrimFunc(symtail), V3187)


tmp12793 := Call(__e, tmp12781, tmp12792)


ifres12780 = tmp12793


} else {
tmp12794 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12780 = tmp12794


}

__e.TailApply(tmp12777, ifres12780)
return


} else {
__e.Return(W3198)
return
}


}, 1)

tmp12813 := Call(__e, PrimFunc(symshen_4hds_a_2), V3187, symctxt)


var ifres12798 Obj

if True == tmp12813 {
tmp12799 := MakeNative(func(__e *ControlFlow) {
W3199 := __e.Get(1)
_ = W3199
tmp12809 := PrimIsPair(W3199)

if True == tmp12809 {
tmp12800 := MakeNative(func(__e *ControlFlow) {
W3200 := __e.Get(1)
_ = W3200
tmp12801 := MakeNative(func(__e *ControlFlow) {
W3201 := __e.Get(1)
_ = W3201
tmp12805 := PrimIsVariable(W3200)

if True == tmp12805 {
tmp12802 := PrimCons(W3200, Nil)

tmp12803 := PrimCons(symctxt, tmp12802)

__e.TailApply(PrimFunc(symshen_4comb), W3201, tmp12803)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp12806 := Call(__e, PrimFunc(symtail), W3199)


__e.TailApply(tmp12801, tmp12806)
return


}, 1)

tmp12807 := Call(__e, PrimFunc(symhead), W3199)


__e.TailApply(tmp12800, tmp12807)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp12810 := Call(__e, PrimFunc(symtail), V3187)


tmp12811 := Call(__e, tmp12799, tmp12810)


ifres12798 = tmp12811


} else {
tmp12812 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12798 = tmp12812


}

__e.TailApply(tmp12776, ifres12798)
return


} else {
__e.Return(W3192)
return
}


}, 1)

tmp12836 := Call(__e, PrimFunc(symshen_4hds_a_2), V3187, symlet)


var ifres12816 Obj

if True == tmp12836 {
tmp12817 := MakeNative(func(__e *ControlFlow) {
W3193 := __e.Get(1)
_ = W3193
tmp12832 := PrimIsPair(W3193)

if True == tmp12832 {
tmp12818 := MakeNative(func(__e *ControlFlow) {
W3194 := __e.Get(1)
_ = W3194
tmp12819 := MakeNative(func(__e *ControlFlow) {
W3195 := __e.Get(1)
_ = W3195
tmp12828 := PrimIsPair(W3195)

if True == tmp12828 {
tmp12820 := MakeNative(func(__e *ControlFlow) {
W3196 := __e.Get(1)
_ = W3196
tmp12821 := MakeNative(func(__e *ControlFlow) {
W3197 := __e.Get(1)
_ = W3197
tmp12822 := PrimCons(W3196, Nil)

tmp12823 := PrimCons(W3194, tmp12822)

tmp12824 := PrimCons(symlet, tmp12823)

__e.TailApply(PrimFunc(symshen_4comb), W3197, tmp12824)
return


}, 1)

tmp12825 := Call(__e, PrimFunc(symtail), W3195)


__e.TailApply(tmp12821, tmp12825)
return


}, 1)

tmp12826 := Call(__e, PrimFunc(symhead), W3195)


__e.TailApply(tmp12820, tmp12826)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp12829 := Call(__e, PrimFunc(symtail), W3193)


__e.TailApply(tmp12819, tmp12829)
return


}, 1)

tmp12830 := Call(__e, PrimFunc(symhead), W3193)


__e.TailApply(tmp12818, tmp12830)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp12833 := Call(__e, PrimFunc(symtail), V3187)


tmp12834 := Call(__e, tmp12817, tmp12833)


ifres12816 = tmp12834


} else {
tmp12835 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12816 = tmp12835


}

__e.TailApply(tmp12775, ifres12816)
return


} else {
__e.Return(W3188)
return
}


}, 1)

tmp12852 := Call(__e, PrimFunc(symshen_4hds_a_2), V3187, symif)


var ifres12839 Obj

if True == tmp12852 {
tmp12840 := MakeNative(func(__e *ControlFlow) {
W3189 := __e.Get(1)
_ = W3189
tmp12848 := PrimIsPair(W3189)

if True == tmp12848 {
tmp12841 := MakeNative(func(__e *ControlFlow) {
W3190 := __e.Get(1)
_ = W3190
tmp12842 := MakeNative(func(__e *ControlFlow) {
W3191 := __e.Get(1)
_ = W3191
tmp12843 := PrimCons(W3190, Nil)

tmp12844 := PrimCons(symif, tmp12843)

__e.TailApply(PrimFunc(symshen_4comb), W3191, tmp12844)
return


}, 1)

tmp12845 := Call(__e, PrimFunc(symtail), W3189)


__e.TailApply(tmp12842, tmp12845)
return


}, 1)

tmp12846 := Call(__e, PrimFunc(symhead), W3189)


__e.TailApply(tmp12841, tmp12846)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp12849 := Call(__e, PrimFunc(symtail), V3187)


tmp12850 := Call(__e, tmp12840, tmp12849)


ifres12839 = tmp12850


} else {
tmp12851 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12839 = tmp12851


}

__e.TailApply(tmp12774, ifres12839)
return


}, 1)

tmp12853 := Call(__e, ns2_1set, symshen_4_5side_6, tmp12773)


_ = tmp12853

tmp12854 := MakeNative(func(__e *ControlFlow) {
V3212 := __e.Get(1)
_ = V3212
V3213 := __e.Get(2)
_ = V3213
V3214 := __e.Get(3)
_ = V3214
tmp12889 := PrimIsPair(V3214)

var ifres12876 Obj

if True == tmp12889 {
tmp12887 := PrimHead(V3214)

tmp12888 := PrimEqual(Nil, tmp12887)

var ifres12878 Obj

if True == tmp12888 {
tmp12885 := PrimTail(V3214)

tmp12886 := PrimIsPair(tmp12885)

var ifres12880 Obj

if True == tmp12886 {
tmp12882 := PrimTail(V3214)

tmp12883 := PrimTail(tmp12882)

tmp12884 := PrimEqual(Nil, tmp12883)

var ifres12881 Obj

if True == tmp12884 {
ifres12881 = True


} else {
ifres12881 = False


}

ifres12880 = ifres12881


} else {
ifres12880 = False


}

var ifres12879 Obj

if True == ifres12880 {
ifres12879 = True


} else {
ifres12879 = False


}

ifres12878 = ifres12879


} else {
ifres12878 = False


}

var ifres12877 Obj

if True == ifres12878 {
ifres12877 = True


} else {
ifres12877 = False


}

ifres12876 = ifres12877


} else {
ifres12876 = False


}

if True == ifres12876 {
tmp12855 := MakeNative(func(__e *ControlFlow) {
W3215 := __e.Get(1)
_ = W3215
tmp12856 := MakeNative(func(__e *ControlFlow) {
W3216 := __e.Get(1)
_ = W3216
tmp12857 := MakeNative(func(__e *ControlFlow) {
W3217 := __e.Get(1)
_ = W3217
tmp12858 := MakeNative(func(__e *ControlFlow) {
W3218 := __e.Get(1)
_ = W3218
tmp12859 := MakeNative(func(__e *ControlFlow) {
W3219 := __e.Get(1)
_ = W3219
tmp12860 := PrimCons(W3218, Nil)

__e.Return(PrimCons(W3219, tmp12860))
return


}, 1)

tmp12861 := PrimCons(V3214, Nil)

tmp12862 := PrimCons(V3213, tmp12861)

tmp12863 := PrimCons(V3212, tmp12862)

__e.TailApply(tmp12859, tmp12863)
return


}, 1)

tmp12864 := PrimCons(W3217, Nil)

tmp12865 := PrimCons(W3216, Nil)

tmp12866 := PrimCons(tmp12864, tmp12865)

tmp12867 := PrimCons(V3212, tmp12866)

__e.TailApply(tmp12858, tmp12867)
return


}, 1)

tmp12868 := Call(__e, PrimFunc(symshen_4coll_1formulae), V3213)


tmp12869 := PrimCons(W3215, Nil)

tmp12870 := PrimCons(tmp12868, tmp12869)

__e.TailApply(tmp12857, tmp12870)
return


}, 1)

tmp12871 := PrimTail(V3214)

tmp12872 := PrimCons(W3215, Nil)

tmp12873 := PrimCons(tmp12871, tmp12872)

__e.TailApply(tmp12856, tmp12873)
return


}, 1)

tmp12874 := Call(__e, PrimFunc(symgensym), symP)


__e.TailApply(tmp12855, tmp12874)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.lr-rule")))
return
}


}, 3)

tmp12890 := Call(__e, ns2_1set, symshen_4lr_1rule, tmp12854)


_ = tmp12890

tmp12891 := MakeNative(func(__e *ControlFlow) {
V3222 := __e.Get(1)
_ = V3222
tmp12920 := PrimEqual(Nil, V3222)

if True == tmp12920 {
__e.Return(Nil)
return
} else {
tmp12918 := PrimIsPair(V3222)

var ifres12898 Obj

if True == tmp12918 {
tmp12916 := PrimHead(V3222)

tmp12917 := PrimIsPair(tmp12916)

var ifres12900 Obj

if True == tmp12917 {
tmp12913 := PrimHead(V3222)

tmp12914 := PrimHead(tmp12913)

tmp12915 := PrimEqual(Nil, tmp12914)

var ifres12902 Obj

if True == tmp12915 {
tmp12910 := PrimHead(V3222)

tmp12911 := PrimTail(tmp12910)

tmp12912 := PrimIsPair(tmp12911)

var ifres12904 Obj

if True == tmp12912 {
tmp12906 := PrimHead(V3222)

tmp12907 := PrimTail(tmp12906)

tmp12908 := PrimTail(tmp12907)

tmp12909 := PrimEqual(Nil, tmp12908)

var ifres12905 Obj

if True == tmp12909 {
ifres12905 = True


} else {
ifres12905 = False


}

ifres12904 = ifres12905


} else {
ifres12904 = False


}

var ifres12903 Obj

if True == ifres12904 {
ifres12903 = True


} else {
ifres12903 = False


}

ifres12902 = ifres12903


} else {
ifres12902 = False


}

var ifres12901 Obj

if True == ifres12902 {
ifres12901 = True


} else {
ifres12901 = False


}

ifres12900 = ifres12901


} else {
ifres12900 = False


}

var ifres12899 Obj

if True == ifres12900 {
ifres12899 = True


} else {
ifres12899 = False


}

ifres12898 = ifres12899


} else {
ifres12898 = False


}

if True == ifres12898 {
tmp12892 := PrimHead(V3222)

tmp12893 := PrimTail(tmp12892)

tmp12894 := PrimHead(tmp12893)

tmp12895 := PrimTail(V3222)

tmp12896 := Call(__e, PrimFunc(symshen_4coll_1formulae), tmp12895)


__e.Return(PrimCons(tmp12894, tmp12896))
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.coll-formulae")))
return
}


}


}, 1)

tmp12921 := Call(__e, ns2_1set, symshen_4coll_1formulae, tmp12891)


_ = tmp12921

tmp12922 := MakeNative(func(__e *ControlFlow) {
V3223 := __e.Get(1)
_ = V3223
tmp12923 := MakeNative(func(__e *ControlFlow) {
W3224 := __e.Get(1)
_ = W3224
tmp12925 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3224)


if True == tmp12925 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3224)
return
}


}, 1)

tmp12937 := PrimIsPair(V3223)

var ifres12926 Obj

if True == tmp12937 {
tmp12927 := MakeNative(func(__e *ControlFlow) {
W3225 := __e.Get(1)
_ = W3225
tmp12928 := MakeNative(func(__e *ControlFlow) {
W3226 := __e.Get(1)
_ = W3226
tmp12931 := Call(__e, PrimFunc(symshen_4key_1in_1sequent_1calculus_2), W3225)


tmp12932 := PrimNot(tmp12931)

if True == tmp12932 {
tmp12929 := Call(__e, PrimFunc(symmacroexpand), W3225)


__e.TailApply(PrimFunc(symshen_4comb), W3226, tmp12929)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp12933 := Call(__e, PrimFunc(symtail), V3223)


__e.TailApply(tmp12928, tmp12933)
return


}, 1)

tmp12934 := Call(__e, PrimFunc(symhead), V3223)


tmp12935 := Call(__e, tmp12927, tmp12934)


ifres12926 = tmp12935


} else {
tmp12936 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12926 = tmp12936


}

__e.TailApply(tmp12923, ifres12926)
return


}, 1)

tmp12938 := Call(__e, ns2_1set, symshen_4_5expr_6, tmp12922)


_ = tmp12938

tmp12939 := MakeNative(func(__e *ControlFlow) {
V3227 := __e.Get(1)
_ = V3227
tmp12946 := PrimIntern(MakeString(";"))

tmp12947 := PrimIntern(MakeString(","))

tmp12948 := PrimIntern(MakeString(":"))

tmp12949 := PrimCons(sym_5_1_1, Nil)

tmp12950 := PrimCons(tmp12948, tmp12949)

tmp12951 := PrimCons(tmp12947, tmp12950)

tmp12952 := PrimCons(tmp12946, tmp12951)

tmp12953 := PrimCons(sym_6_6, tmp12952)

tmp12954 := Call(__e, PrimFunc(symelement_2), V3227, tmp12953)


if True == tmp12954 {
__e.Return(True)
return
} else {
tmp12944 := Call(__e, PrimFunc(symshen_4sng_2), V3227)


var ifres12941 Obj

if True == tmp12944 {
ifres12941 = True


} else {
tmp12943 := Call(__e, PrimFunc(symshen_4dbl_2), V3227)


var ifres12942 Obj

if True == tmp12943 {
ifres12942 = True


} else {
ifres12942 = False


}

ifres12941 = ifres12942


}

if True == ifres12941 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp12955 := Call(__e, ns2_1set, symshen_4key_1in_1sequent_1calculus_2, tmp12939)


_ = tmp12955

tmp12956 := MakeNative(func(__e *ControlFlow) {
V3228 := __e.Get(1)
_ = V3228
tmp12957 := MakeNative(func(__e *ControlFlow) {
W3229 := __e.Get(1)
_ = W3229
tmp12959 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3229)


if True == tmp12959 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3229)
return
}


}, 1)

tmp12960 := MakeNative(func(__e *ControlFlow) {
W3230 := __e.Get(1)
_ = W3230
tmp12966 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3230)


if True == tmp12966 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12961 := MakeNative(func(__e *ControlFlow) {
W3231 := __e.Get(1)
_ = W3231
tmp12962 := MakeNative(func(__e *ControlFlow) {
W3232 := __e.Get(1)
_ = W3232
__e.TailApply(PrimFunc(symshen_4comb), W3232, W3231)
return
}, 1)

tmp12963 := Call(__e, PrimFunc(symshen_4in_1_6), W3230)


__e.TailApply(tmp12962, tmp12963)
return


}, 1)

tmp12964 := Call(__e, PrimFunc(symshen_4_5_1out), W3230)


__e.TailApply(tmp12961, tmp12964)
return


}


}, 1)

tmp12967 := Call(__e, PrimFunc(symshen_4_5expr_6), V3228)


tmp12968 := Call(__e, tmp12960, tmp12967)


__e.TailApply(tmp12957, tmp12968)
return


}, 1)

tmp12969 := Call(__e, ns2_1set, symshen_4_5type_6, tmp12956)


_ = tmp12969

tmp12970 := MakeNative(func(__e *ControlFlow) {
V3233 := __e.Get(1)
_ = V3233
tmp12971 := MakeNative(func(__e *ControlFlow) {
W3234 := __e.Get(1)
_ = W3234
tmp12973 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3234)


if True == tmp12973 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3234)
return
}


}, 1)

tmp12983 := PrimIsPair(V3233)

var ifres12974 Obj

if True == tmp12983 {
tmp12975 := MakeNative(func(__e *ControlFlow) {
W3235 := __e.Get(1)
_ = W3235
tmp12976 := MakeNative(func(__e *ControlFlow) {
W3236 := __e.Get(1)
_ = W3236
tmp12978 := Call(__e, PrimFunc(symshen_4dbl_2), W3235)


if True == tmp12978 {
__e.TailApply(PrimFunc(symshen_4comb), W3236, W3235)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp12979 := Call(__e, PrimFunc(symtail), V3233)


__e.TailApply(tmp12976, tmp12979)
return


}, 1)

tmp12980 := Call(__e, PrimFunc(symhead), V3233)


tmp12981 := Call(__e, tmp12975, tmp12980)


ifres12974 = tmp12981


} else {
tmp12982 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12974 = tmp12982


}

__e.TailApply(tmp12971, ifres12974)
return


}, 1)

tmp12984 := Call(__e, ns2_1set, symshen_4_5dbl_6, tmp12970)


_ = tmp12984

tmp12985 := MakeNative(func(__e *ControlFlow) {
V3237 := __e.Get(1)
_ = V3237
tmp12986 := MakeNative(func(__e *ControlFlow) {
W3238 := __e.Get(1)
_ = W3238
tmp12988 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3238)


if True == tmp12988 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3238)
return
}


}, 1)

tmp12998 := PrimIsPair(V3237)

var ifres12989 Obj

if True == tmp12998 {
tmp12990 := MakeNative(func(__e *ControlFlow) {
W3239 := __e.Get(1)
_ = W3239
tmp12991 := MakeNative(func(__e *ControlFlow) {
W3240 := __e.Get(1)
_ = W3240
tmp12993 := Call(__e, PrimFunc(symshen_4sng_2), W3239)


if True == tmp12993 {
__e.TailApply(PrimFunc(symshen_4comb), W3240, W3239)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp12994 := Call(__e, PrimFunc(symtail), V3237)


__e.TailApply(tmp12991, tmp12994)
return


}, 1)

tmp12995 := Call(__e, PrimFunc(symhead), V3237)


tmp12996 := Call(__e, tmp12990, tmp12995)


ifres12989 = tmp12996


} else {
tmp12997 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12989 = tmp12997


}

__e.TailApply(tmp12986, ifres12989)
return


}, 1)

tmp12999 := Call(__e, ns2_1set, symshen_4_5sng_6, tmp12985)


_ = tmp12999

tmp13000 := MakeNative(func(__e *ControlFlow) {
V3241 := __e.Get(1)
_ = V3241
tmp13005 := PrimIsSymbol(V3241)

if True == tmp13005 {
tmp13002 := PrimStr(V3241)

tmp13003 := Call(__e, PrimFunc(symshen_4sng_1h_2), tmp13002)


if True == tmp13003 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


} else {
__e.Return(False)
return
}


}, 1)

tmp13006 := Call(__e, ns2_1set, symshen_4sng_2, tmp13000)


_ = tmp13006

tmp13007 := MakeNative(func(__e *ControlFlow) {
V3244 := __e.Get(1)
_ = V3244
tmp13016 := PrimEqual(MakeString("___"), V3244)

if True == tmp13016 {
__e.Return(True)
return
} else {
tmp13014 := Call(__e, PrimFunc(symshen_4_7string_2), V3244)


var ifres13010 Obj

if True == tmp13014 {
tmp13012 := Call(__e, PrimFunc(symhdstr), V3244)


tmp13013 := PrimEqual(MakeString("_"), tmp13012)

var ifres13011 Obj

if True == tmp13013 {
ifres13011 = True


} else {
ifres13011 = False


}

ifres13010 = ifres13011


} else {
ifres13010 = False


}

if True == ifres13010 {
tmp13008 := PrimTailString(V3244)

__e.TailApply(PrimFunc(symshen_4sng_1h_2), tmp13008)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp13017 := Call(__e, ns2_1set, symshen_4sng_1h_2, tmp13007)


_ = tmp13017

tmp13018 := MakeNative(func(__e *ControlFlow) {
V3245 := __e.Get(1)
_ = V3245
tmp13023 := PrimIsSymbol(V3245)

if True == tmp13023 {
tmp13020 := PrimStr(V3245)

tmp13021 := Call(__e, PrimFunc(symshen_4dbl_1h_2), tmp13020)


if True == tmp13021 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


} else {
__e.Return(False)
return
}


}, 1)

tmp13024 := Call(__e, ns2_1set, symshen_4dbl_2, tmp13018)


_ = tmp13024

tmp13025 := MakeNative(func(__e *ControlFlow) {
V3248 := __e.Get(1)
_ = V3248
tmp13034 := PrimEqual(MakeString("==="), V3248)

if True == tmp13034 {
__e.Return(True)
return
} else {
tmp13032 := Call(__e, PrimFunc(symshen_4_7string_2), V3248)


var ifres13028 Obj

if True == tmp13032 {
tmp13030 := Call(__e, PrimFunc(symhdstr), V3248)


tmp13031 := PrimEqual(MakeString("="), tmp13030)

var ifres13029 Obj

if True == tmp13031 {
ifres13029 = True


} else {
ifres13029 = False


}

ifres13028 = ifres13029


} else {
ifres13028 = False


}

if True == ifres13028 {
tmp13026 := PrimTailString(V3248)

__e.TailApply(PrimFunc(symshen_4dbl_1h_2), tmp13026)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp13035 := Call(__e, ns2_1set, symshen_4dbl_1h_2, tmp13025)


_ = tmp13035

tmp13036 := MakeNative(func(__e *ControlFlow) {
V3249 := __e.Get(1)
_ = V3249
V3250 := __e.Get(2)
_ = V3250
tmp13037 := MakeNative(func(__e *ControlFlow) {
W3251 := __e.Get(1)
_ = W3251
tmp13038 := MakeNative(func(__e *ControlFlow) {
W3253 := __e.Get(1)
_ = W3253
__e.TailApply(PrimFunc(symeval), W3253)
return
}, 1)

tmp13039 := PrimCons(V3249, W3251)

tmp13040 := PrimCons(symdefprolog, tmp13039)

__e.TailApply(tmp13038, tmp13040)
return


}, 1)

tmp13041 := MakeNative(func(__e *ControlFlow) {
Z3252 := __e.Get(1)
_ = Z3252
__e.TailApply(PrimFunc(symshen_4rule_1_6clause), Z3252)
return
}, 1)

tmp13042 := Call(__e, PrimFunc(symmapcan), tmp13041, V3250)


__e.TailApply(tmp13037, tmp13042)
return


}, 2)

tmp13043 := Call(__e, ns2_1set, symshen_4rules_1_6prolog, tmp13036)


_ = tmp13043

tmp13044 := MakeNative(func(__e *ControlFlow) {
V3254 := __e.Get(1)
_ = V3254
tmp13105 := PrimIsPair(V3254)

var ifres13069 Obj

if True == tmp13105 {
tmp13103 := PrimTail(V3254)

tmp13104 := PrimIsPair(tmp13103)

var ifres13071 Obj

if True == tmp13104 {
tmp13100 := PrimTail(V3254)

tmp13101 := PrimTail(tmp13100)

tmp13102 := PrimIsPair(tmp13101)

var ifres13073 Obj

if True == tmp13102 {
tmp13096 := PrimTail(V3254)

tmp13097 := PrimTail(tmp13096)

tmp13098 := PrimHead(tmp13097)

tmp13099 := PrimIsPair(tmp13098)

var ifres13075 Obj

if True == tmp13099 {
tmp13091 := PrimTail(V3254)

tmp13092 := PrimTail(tmp13091)

tmp13093 := PrimHead(tmp13092)

tmp13094 := PrimTail(tmp13093)

tmp13095 := PrimIsPair(tmp13094)

var ifres13077 Obj

if True == tmp13095 {
tmp13085 := PrimTail(V3254)

tmp13086 := PrimTail(tmp13085)

tmp13087 := PrimHead(tmp13086)

tmp13088 := PrimTail(tmp13087)

tmp13089 := PrimTail(tmp13088)

tmp13090 := PrimEqual(Nil, tmp13089)

var ifres13079 Obj

if True == tmp13090 {
tmp13081 := PrimTail(V3254)

tmp13082 := PrimTail(tmp13081)

tmp13083 := PrimTail(tmp13082)

tmp13084 := PrimEqual(Nil, tmp13083)

var ifres13080 Obj

if True == tmp13084 {
ifres13080 = True


} else {
ifres13080 = False


}

ifres13079 = ifres13080


} else {
ifres13079 = False


}

var ifres13078 Obj

if True == ifres13079 {
ifres13078 = True


} else {
ifres13078 = False


}

ifres13077 = ifres13078


} else {
ifres13077 = False


}

var ifres13076 Obj

if True == ifres13077 {
ifres13076 = True


} else {
ifres13076 = False


}

ifres13075 = ifres13076


} else {
ifres13075 = False


}

var ifres13074 Obj

if True == ifres13075 {
ifres13074 = True


} else {
ifres13074 = False


}

ifres13073 = ifres13074


} else {
ifres13073 = False


}

var ifres13072 Obj

if True == ifres13073 {
ifres13072 = True


} else {
ifres13072 = False


}

ifres13071 = ifres13072


} else {
ifres13071 = False


}

var ifres13070 Obj

if True == ifres13071 {
ifres13070 = True


} else {
ifres13070 = False


}

ifres13069 = ifres13070


} else {
ifres13069 = False


}

if True == ifres13069 {
tmp13045 := MakeNative(func(__e *ControlFlow) {
W3255 := __e.Get(1)
_ = W3255
tmp13046 := PrimTail(V3254)

tmp13047 := PrimTail(tmp13046)

tmp13048 := PrimHead(tmp13047)

tmp13049 := PrimTail(tmp13048)

tmp13050 := PrimHead(tmp13049)

tmp13051 := Call(__e, PrimFunc(symshen_4rule_1_6head), tmp13050)


tmp13052 := PrimCons(sym_5_1_1, Nil)

tmp13053 := PrimHead(V3254)

tmp13054 := PrimTail(V3254)

tmp13055 := PrimHead(tmp13054)

tmp13056 := PrimTail(V3254)

tmp13057 := PrimTail(tmp13056)

tmp13058 := PrimHead(tmp13057)

tmp13059 := PrimHead(tmp13058)

tmp13060 := Call(__e, PrimFunc(symshen_4rule_1_6body), W3255, symAssumptions, tmp13053, tmp13055, tmp13059)


tmp13061 := Call(__e, PrimFunc(symappend), tmp13052, tmp13060)


__e.TailApply(PrimFunc(symappend), tmp13051, tmp13061)
return


}, 1)

tmp13062 := PrimTail(V3254)

tmp13063 := PrimTail(tmp13062)

tmp13064 := PrimHead(tmp13063)

tmp13065 := PrimTail(tmp13064)

tmp13066 := PrimHead(tmp13065)

tmp13067 := Call(__e, PrimFunc(symshen_4extract_1vars), tmp13066)


__e.TailApply(tmp13045, tmp13067)
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.rule->clause")))
return
}


}, 1)

tmp13106 := Call(__e, ns2_1set, symshen_4rule_1_6clause, tmp13044)


_ = tmp13106

tmp13107 := MakeNative(func(__e *ControlFlow) {
V3256 := __e.Get(1)
_ = V3256
tmp13108 := Call(__e, PrimFunc(symshen_4macro_1_8ch), V3256)


tmp13109 := PrimCons(symAssumptions, Nil)

__e.Return(PrimCons(tmp13108, tmp13109))
return


}, 1)

tmp13110 := Call(__e, ns2_1set, symshen_4rule_1_6head, tmp13107)


_ = tmp13110

tmp13111 := MakeNative(func(__e *ControlFlow) {
V3257 := __e.Get(1)
_ = V3257
tmp13112 := PrimCons(V3257, Nil)

__e.Return(PrimCons(symshen_4_8ch, tmp13112))
return


}, 1)

tmp13113 := Call(__e, ns2_1set, symshen_4macro_1_8ch, tmp13111)


_ = tmp13113

tmp13114 := MakeNative(func(__e *ControlFlow) {
V3258 := __e.Get(1)
_ = V3258
tmp13115 := PrimCons(V3258, Nil)

__e.Return(PrimCons(symshen_4_8c, tmp13115))
return


}, 1)

tmp13116 := Call(__e, ns2_1set, symshen_4macro_1_8c, tmp13114)


_ = tmp13116

tmp13117 := MakeNative(func(__e *ControlFlow) {
V3259 := __e.Get(1)
_ = V3259
V3260 := __e.Get(2)
_ = V3260
V3261 := __e.Get(3)
_ = V3261
V3262 := __e.Get(4)
_ = V3262
V3263 := __e.Get(5)
_ = V3263
tmp13152 := PrimEqual(Nil, V3263)

if True == tmp13152 {
__e.TailApply(PrimFunc(symshen_4side_1conditions_1_6goals), Nil, V3259, V3260, V3261, V3262)
return
} else {
tmp13150 := PrimEqual(Nil, V3262)

var ifres13143 Obj

if True == tmp13150 {
tmp13149 := PrimIsPair(V3263)

var ifres13145 Obj

if True == tmp13149 {
tmp13147 := PrimTail(V3263)

tmp13148 := PrimEqual(Nil, tmp13147)

var ifres13146 Obj

if True == tmp13148 {
ifres13146 = True


} else {
ifres13146 = False


}

ifres13145 = ifres13146


} else {
ifres13145 = False


}

var ifres13144 Obj

if True == ifres13145 {
ifres13144 = True


} else {
ifres13144 = False


}

ifres13143 = ifres13144


} else {
ifres13143 = False


}

if True == ifres13143 {
tmp13118 := MakeNative(func(__e *ControlFlow) {
W3264 := __e.Get(1)
_ = W3264
tmp13119 := MakeNative(func(__e *ControlFlow) {
W3265 := __e.Get(1)
_ = W3265
tmp13120 := PrimHead(V3263)

tmp13121 := Call(__e, PrimFunc(symshen_4specialise_1member), tmp13120, V3260, W3265, W3264)


tmp13122 := Call(__e, PrimFunc(symshen_4side_1conditions_1_6goals), Nil, V3259, V3260, V3261, Nil)


__e.Return(PrimCons(tmp13121, tmp13122))
return


}, 1)

tmp13123 := PrimHead(V3263)

tmp13124 := Call(__e, PrimFunc(symshen_4remove_1bystanders), V3259, tmp13123)


__e.TailApply(tmp13119, tmp13124)
return


}, 1)

tmp13125 := PrimHead(V3263)

tmp13126 := Call(__e, PrimFunc(symshen_4passive_1variables), tmp13125, V3259)


__e.TailApply(tmp13118, tmp13126)
return


} else {
tmp13141 := PrimIsPair(V3263)

if True == tmp13141 {
tmp13127 := MakeNative(func(__e *ControlFlow) {
W3266 := __e.Get(1)
_ = W3266
tmp13128 := MakeNative(func(__e *ControlFlow) {
W3267 := __e.Get(1)
_ = W3267
tmp13129 := MakeNative(func(__e *ControlFlow) {
W3268 := __e.Get(1)
_ = W3268
tmp13130 := PrimHead(V3263)

tmp13131 := Call(__e, PrimFunc(symshen_4specialise_1consume), tmp13130, V3260, W3268, W3267, W3266)


tmp13132 := Call(__e, PrimFunc(symappend), V3259, W3267)


tmp13133 := PrimTail(V3263)

tmp13134 := Call(__e, PrimFunc(symshen_4rule_1_6body), tmp13132, W3266, V3261, V3262, tmp13133)


__e.Return(PrimCons(tmp13131, tmp13134))
return


}, 1)

tmp13135 := PrimHead(V3263)

tmp13136 := Call(__e, PrimFunc(symshen_4remove_1bystanders), V3259, tmp13135)


__e.TailApply(tmp13129, tmp13136)
return


}, 1)

tmp13137 := PrimHead(V3263)

tmp13138 := Call(__e, PrimFunc(symshen_4passive_1variables), tmp13137, V3259)


__e.TailApply(tmp13128, tmp13138)
return


}, 1)

tmp13139 := Call(__e, PrimFunc(symgensym), symNewAssumptions)


__e.TailApply(tmp13127, tmp13139)
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.rule->body")))
return
}


}


}


}, 5)

tmp13153 := Call(__e, ns2_1set, symshen_4rule_1_6body, tmp13117)


_ = tmp13153

tmp13154 := MakeNative(func(__e *ControlFlow) {
V3269 := __e.Get(1)
_ = V3269
V3270 := __e.Get(2)
_ = V3270
V3271 := __e.Get(3)
_ = V3271
V3272 := __e.Get(4)
_ = V3272
tmp13155 := MakeNative(func(__e *ControlFlow) {
W3273 := __e.Get(1)
_ = W3273
tmp13156 := MakeNative(func(__e *ControlFlow) {
W3274 := __e.Get(1)
_ = W3274
tmp13157 := Call(__e, PrimFunc(symappend), V3271, V3272)


tmp13158 := PrimCons(V3270, tmp13157)

__e.Return(PrimCons(W3273, tmp13158))
return


}, 1)

tmp13159 := Call(__e, PrimFunc(symshen_4member_1clause), W3273, V3269, V3271, V3272)


__e.TailApply(tmp13156, tmp13159)
return


}, 1)

tmp13160 := Call(__e, PrimFunc(symgensym), symshen_4member)


__e.TailApply(tmp13155, tmp13160)
return


}, 4)

tmp13161 := Call(__e, ns2_1set, symshen_4specialise_1member, tmp13154)


_ = tmp13161

tmp13162 := MakeNative(func(__e *ControlFlow) {
V3277 := __e.Get(1)
_ = V3277
V3278 := __e.Get(2)
_ = V3278
tmp13176 := PrimEqual(Nil, V3277)

if True == tmp13176 {
__e.Return(Nil)
return
} else {
tmp13174 := PrimIsPair(V3277)

var ifres13170 Obj

if True == tmp13174 {
tmp13172 := PrimHead(V3277)

tmp13173 := Call(__e, PrimFunc(symshen_4occurs_1check_2), tmp13172, V3278)


var ifres13171 Obj

if True == tmp13173 {
ifres13171 = True


} else {
ifres13171 = False


}

ifres13170 = ifres13171


} else {
ifres13170 = False


}

if True == ifres13170 {
tmp13163 := PrimHead(V3277)

tmp13164 := PrimTail(V3277)

tmp13165 := Call(__e, PrimFunc(symshen_4remove_1bystanders), tmp13164, V3278)


__e.Return(PrimCons(tmp13163, tmp13165))
return


} else {
tmp13168 := PrimIsPair(V3277)

if True == tmp13168 {
tmp13166 := PrimTail(V3277)

__e.TailApply(PrimFunc(symshen_4remove_1bystanders), tmp13166, V3278)
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.remove-bystanders")))
return
}


}


}


}, 2)

tmp13177 := Call(__e, ns2_1set, symshen_4remove_1bystanders, tmp13162)


_ = tmp13177

tmp13178 := MakeNative(func(__e *ControlFlow) {
V3279 := __e.Get(1)
_ = V3279
V3280 := __e.Get(2)
_ = V3280
V3281 := __e.Get(3)
_ = V3281
V3282 := __e.Get(4)
_ = V3282
tmp13179 := MakeNative(func(__e *ControlFlow) {
W3283 := __e.Get(1)
_ = W3283
tmp13180 := MakeNative(func(__e *ControlFlow) {
W3284 := __e.Get(1)
_ = W3284
tmp13181 := MakeNative(func(__e *ControlFlow) {
W3285 := __e.Get(1)
_ = W3285
tmp13182 := MakeNative(func(__e *ControlFlow) {
W3290 := __e.Get(1)
_ = W3290
__e.TailApply(PrimFunc(symeval), W3290)
return
}, 1)

tmp13183 := Call(__e, PrimFunc(symappend), W3284, W3285)


tmp13184 := PrimCons(V3279, tmp13183)

tmp13185 := PrimCons(symdefprolog, tmp13184)

__e.TailApply(tmp13182, tmp13185)
return


}, 1)

tmp13186 := MakeNative(func(__e *ControlFlow) {
W3286 := __e.Get(1)
_ = W3286
tmp13187 := MakeNative(func(__e *ControlFlow) {
W3287 := __e.Get(1)
_ = W3287
tmp13188 := MakeNative(func(__e *ControlFlow) {
W3288 := __e.Get(1)
_ = W3288
tmp13189 := MakeNative(func(__e *ControlFlow) {
W3289 := __e.Get(1)
_ = W3289
tmp13190 := PrimCons(sym_5_1_1, Nil)

tmp13191 := PrimIntern(MakeString(";"))

tmp13192 := PrimCons(tmp13191, Nil)

tmp13193 := Call(__e, PrimFunc(symappend), W3289, tmp13192)


tmp13194 := Call(__e, PrimFunc(symappend), tmp13190, tmp13193)


__e.TailApply(PrimFunc(symappend), W3288, tmp13194)
return


}, 1)

tmp13195 := PrimCons(W3286, W3287)

tmp13196 := PrimCons(V3279, tmp13195)

tmp13197 := PrimCons(tmp13196, Nil)

__e.TailApply(tmp13189, tmp13197)
return


}, 1)

tmp13198 := PrimCons(W3286, Nil)

tmp13199 := PrimCons(sym__, tmp13198)

tmp13200 := PrimCons(symcons, tmp13199)

tmp13201 := PrimCons(tmp13200, Nil)

tmp13202 := PrimCons(sym_1, tmp13201)

tmp13203 := PrimCons(tmp13202, Nil)

tmp13204 := Call(__e, PrimFunc(symappend), tmp13203, W3287)


__e.TailApply(tmp13188, tmp13204)
return


}, 1)

tmp13205 := Call(__e, PrimFunc(symappend), V3281, V3282)


__e.TailApply(tmp13187, tmp13205)
return


}, 1)

tmp13206 := Call(__e, PrimFunc(symgensym), symHypotheses)


tmp13207 := Call(__e, tmp13186, tmp13206)


__e.TailApply(tmp13181, tmp13207)
return


}, 1)

tmp13208 := Call(__e, PrimFunc(symshen_4macro_1_8ch), V3280)


tmp13209 := PrimCons(sym__, Nil)

tmp13210 := PrimCons(tmp13208, tmp13209)

tmp13211 := PrimCons(symcons, tmp13210)

tmp13212 := PrimCons(tmp13211, Nil)

tmp13213 := PrimCons(sym_1, tmp13212)

tmp13214 := PrimCons(tmp13213, Nil)

tmp13215 := PrimCons(sym_5_1_1, Nil)

tmp13216 := Call(__e, PrimFunc(symshen_4passive_1bind), V3282, W3283)


tmp13217 := PrimIntern(MakeString(";"))

tmp13218 := PrimCons(tmp13217, Nil)

tmp13219 := Call(__e, PrimFunc(symappend), tmp13216, tmp13218)


tmp13220 := Call(__e, PrimFunc(symappend), tmp13215, tmp13219)


tmp13221 := Call(__e, PrimFunc(symappend), W3283, tmp13220)


tmp13222 := Call(__e, PrimFunc(symappend), V3281, tmp13221)


tmp13223 := Call(__e, PrimFunc(symappend), tmp13214, tmp13222)


__e.TailApply(tmp13180, tmp13223)
return


}, 1)

tmp13224 := Call(__e, PrimFunc(symlength), V3282)


tmp13225 := Call(__e, PrimFunc(symshen_4nvars), tmp13224)


__e.TailApply(tmp13179, tmp13225)
return


}, 4)

tmp13226 := Call(__e, ns2_1set, symshen_4member_1clause, tmp13178)


_ = tmp13226

tmp13227 := MakeNative(func(__e *ControlFlow) {
V3291 := __e.Get(1)
_ = V3291
tmp13232 := PrimEqual(MakeNumber(0), V3291)

if True == tmp13232 {
__e.Return(Nil)
return
} else {
tmp13228 := Call(__e, PrimFunc(symgensym), symNewV)


tmp13229 := PrimNumberSubtract(V3291, MakeNumber(1))

tmp13230 := Call(__e, PrimFunc(symshen_4nvars), tmp13229)


__e.Return(PrimCons(tmp13228, tmp13230))
return


}


}, 1)

tmp13233 := Call(__e, ns2_1set, symshen_4nvars, tmp13227)


_ = tmp13233

tmp13234 := MakeNative(func(__e *ControlFlow) {
V3292 := __e.Get(1)
_ = V3292
V3293 := __e.Get(2)
_ = V3293
tmp13252 := PrimEqual(Nil, V3292)

var ifres13249 Obj

if True == tmp13252 {
tmp13251 := PrimEqual(Nil, V3293)

var ifres13250 Obj

if True == tmp13251 {
ifres13250 = True


} else {
ifres13250 = False


}

ifres13249 = ifres13250


} else {
ifres13249 = False


}

if True == ifres13249 {
__e.Return(Nil)
return
} else {
tmp13247 := PrimIsPair(V3292)

var ifres13244 Obj

if True == tmp13247 {
tmp13246 := PrimIsPair(V3293)

var ifres13245 Obj

if True == tmp13246 {
ifres13245 = True


} else {
ifres13245 = False


}

ifres13244 = ifres13245


} else {
ifres13244 = False


}

if True == ifres13244 {
tmp13235 := PrimHead(V3293)

tmp13236 := PrimHead(V3292)

tmp13237 := PrimCons(tmp13236, Nil)

tmp13238 := PrimCons(tmp13235, tmp13237)

tmp13239 := PrimCons(symbind, tmp13238)

tmp13240 := PrimTail(V3292)

tmp13241 := PrimTail(V3293)

tmp13242 := Call(__e, PrimFunc(symshen_4passive_1bind), tmp13240, tmp13241)


__e.Return(PrimCons(tmp13239, tmp13242))
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.passive-bind")))
return
}


}


}, 2)

tmp13253 := Call(__e, ns2_1set, symshen_4passive_1bind, tmp13234)


_ = tmp13253

tmp13254 := MakeNative(func(__e *ControlFlow) {
V3294 := __e.Get(1)
_ = V3294
V3295 := __e.Get(2)
_ = V3295
V3296 := __e.Get(3)
_ = V3296
V3297 := __e.Get(4)
_ = V3297
V3298 := __e.Get(5)
_ = V3298
tmp13255 := MakeNative(func(__e *ControlFlow) {
W3299 := __e.Get(1)
_ = W3299
tmp13256 := MakeNative(func(__e *ControlFlow) {
W3300 := __e.Get(1)
_ = W3300
tmp13257 := Call(__e, PrimFunc(symappend), V3296, V3297)


tmp13258 := PrimCons(V3298, tmp13257)

tmp13259 := PrimCons(V3295, tmp13258)

__e.Return(PrimCons(W3299, tmp13259))
return


}, 1)

tmp13260 := Call(__e, PrimFunc(symshen_4consume_1clause), W3299, V3294, V3296, V3297, V3298)


__e.TailApply(tmp13256, tmp13260)
return


}, 1)

tmp13261 := Call(__e, PrimFunc(symgensym), symshen_4consume)


__e.TailApply(tmp13255, tmp13261)
return


}, 5)

tmp13262 := Call(__e, ns2_1set, symshen_4specialise_1consume, tmp13254)


_ = tmp13262

tmp13263 := MakeNative(func(__e *ControlFlow) {
V3301 := __e.Get(1)
_ = V3301
V3302 := __e.Get(2)
_ = V3302
V3303 := __e.Get(3)
_ = V3303
V3304 := __e.Get(4)
_ = V3304
V3305 := __e.Get(5)
_ = V3305
tmp13264 := MakeNative(func(__e *ControlFlow) {
W3306 := __e.Get(1)
_ = W3306
tmp13265 := MakeNative(func(__e *ControlFlow) {
W3307 := __e.Get(1)
_ = W3307
tmp13266 := MakeNative(func(__e *ControlFlow) {
W3308 := __e.Get(1)
_ = W3308
tmp13267 := MakeNative(func(__e *ControlFlow) {
W3309 := __e.Get(1)
_ = W3309
tmp13268 := MakeNative(func(__e *ControlFlow) {
W3315 := __e.Get(1)
_ = W3315
__e.TailApply(PrimFunc(symeval), W3315)
return
}, 1)

tmp13269 := Call(__e, PrimFunc(symappend), W3308, W3309)


tmp13270 := PrimCons(V3301, tmp13269)

tmp13271 := PrimCons(symdefprolog, tmp13270)

__e.TailApply(tmp13268, tmp13271)
return


}, 1)

tmp13272 := MakeNative(func(__e *ControlFlow) {
W3310 := __e.Get(1)
_ = W3310
tmp13273 := MakeNative(func(__e *ControlFlow) {
W3311 := __e.Get(1)
_ = W3311
tmp13274 := MakeNative(func(__e *ControlFlow) {
W3312 := __e.Get(1)
_ = W3312
tmp13275 := MakeNative(func(__e *ControlFlow) {
W3313 := __e.Get(1)
_ = W3313
tmp13276 := MakeNative(func(__e *ControlFlow) {
W3314 := __e.Get(1)
_ = W3314
tmp13277 := PrimCons(sym_5_1_1, Nil)

tmp13278 := PrimIntern(MakeString(";"))

tmp13279 := PrimCons(tmp13278, Nil)

tmp13280 := Call(__e, PrimFunc(symappend), W3314, tmp13279)


tmp13281 := Call(__e, PrimFunc(symappend), tmp13277, tmp13280)


__e.TailApply(PrimFunc(symappend), W3313, tmp13281)
return


}, 1)

tmp13282 := PrimCons(W3307, Nil)

tmp13283 := PrimCons(W3312, tmp13282)

tmp13284 := PrimCons(symbind, tmp13283)

tmp13285 := PrimCons(V3305, W3311)

tmp13286 := PrimCons(W3310, tmp13285)

tmp13287 := PrimCons(V3301, tmp13286)

tmp13288 := PrimCons(tmp13287, Nil)

tmp13289 := PrimCons(tmp13284, tmp13288)

__e.TailApply(tmp13276, tmp13289)
return


}, 1)

tmp13290 := PrimCons(W3310, Nil)

tmp13291 := PrimCons(W3307, tmp13290)

tmp13292 := PrimCons(symcons, tmp13291)

tmp13293 := PrimCons(tmp13292, Nil)

tmp13294 := PrimCons(sym_1, tmp13293)

tmp13295 := PrimCons(V3305, Nil)

tmp13296 := PrimCons(W3312, tmp13295)

tmp13297 := PrimCons(symcons, tmp13296)

tmp13298 := PrimCons(tmp13297, W3311)

tmp13299 := PrimCons(tmp13294, tmp13298)

__e.TailApply(tmp13275, tmp13299)
return


}, 1)

tmp13300 := Call(__e, PrimFunc(symgensym), symAssumptions)


__e.TailApply(tmp13274, tmp13300)
return


}, 1)

tmp13301 := Call(__e, PrimFunc(symappend), V3303, V3304)


__e.TailApply(tmp13273, tmp13301)
return


}, 1)

tmp13302 := Call(__e, PrimFunc(symgensym), symHypotheses)


tmp13303 := Call(__e, tmp13272, tmp13302)


__e.TailApply(tmp13267, tmp13303)
return


}, 1)

tmp13304 := Call(__e, PrimFunc(symshen_4macro_1_8ch), V3302)


tmp13305 := PrimCons(W3307, Nil)

tmp13306 := PrimCons(tmp13304, tmp13305)

tmp13307 := PrimCons(symcons, tmp13306)

tmp13308 := PrimCons(tmp13307, Nil)

tmp13309 := PrimCons(sym_1, tmp13308)

tmp13310 := PrimCons(sym_5_1_1, Nil)

tmp13311 := Call(__e, PrimFunc(symshen_4passive_1bind), V3304, W3306)


tmp13312 := PrimCons(W3307, Nil)

tmp13313 := PrimCons(V3305, tmp13312)

tmp13314 := PrimCons(symbind, tmp13313)

tmp13315 := PrimIntern(MakeString(";"))

tmp13316 := PrimCons(tmp13315, Nil)

tmp13317 := PrimCons(tmp13314, tmp13316)

tmp13318 := Call(__e, PrimFunc(symappend), tmp13311, tmp13317)


tmp13319 := Call(__e, PrimFunc(symappend), tmp13310, tmp13318)


tmp13320 := Call(__e, PrimFunc(symappend), W3306, tmp13319)


tmp13321 := Call(__e, PrimFunc(symappend), V3303, tmp13320)


tmp13322 := PrimCons(V3305, tmp13321)

tmp13323 := PrimCons(tmp13309, tmp13322)

__e.TailApply(tmp13266, tmp13323)
return


}, 1)

tmp13324 := Call(__e, PrimFunc(symgensym), symAssumption)


__e.TailApply(tmp13265, tmp13324)
return


}, 1)

tmp13325 := Call(__e, PrimFunc(symlength), V3304)


tmp13326 := Call(__e, PrimFunc(symshen_4nvars), tmp13325)


__e.TailApply(tmp13264, tmp13326)
return


}, 5)

tmp13327 := Call(__e, ns2_1set, symshen_4consume_1clause, tmp13263)


_ = tmp13327

tmp13328 := MakeNative(func(__e *ControlFlow) {
V3316 := __e.Get(1)
_ = V3316
V3317 := __e.Get(2)
_ = V3317
tmp13329 := Call(__e, PrimFunc(symshen_4extract_1vars), V3316)


__e.TailApply(PrimFunc(symdifference), tmp13329, V3317)
return


}, 2)

tmp13330 := Call(__e, ns2_1set, symshen_4passive_1variables, tmp13328)


_ = tmp13330

tmp13331 := MakeNative(func(__e *ControlFlow) {
V3322 := __e.Get(1)
_ = V3322
V3323 := __e.Get(2)
_ = V3323
V3324 := __e.Get(3)
_ = V3324
V3325 := __e.Get(4)
_ = V3325
V3326 := __e.Get(5)
_ = V3326
tmp13482 := PrimEqual(Nil, V3325)

if True == tmp13482 {
__e.TailApply(PrimFunc(symshen_4premises_1_6goals), V3322, V3324, V3326)
return
} else {
tmp13480 := PrimIsPair(V3325)

var ifres13460 Obj

if True == tmp13480 {
tmp13478 := PrimHead(V3325)

tmp13479 := PrimIsPair(tmp13478)

var ifres13462 Obj

if True == tmp13479 {
tmp13475 := PrimHead(V3325)

tmp13476 := PrimHead(tmp13475)

tmp13477 := PrimEqual(symif, tmp13476)

var ifres13464 Obj

if True == tmp13477 {
tmp13472 := PrimHead(V3325)

tmp13473 := PrimTail(tmp13472)

tmp13474 := PrimIsPair(tmp13473)

var ifres13466 Obj

if True == tmp13474 {
tmp13468 := PrimHead(V3325)

tmp13469 := PrimTail(tmp13468)

tmp13470 := PrimTail(tmp13469)

tmp13471 := PrimEqual(Nil, tmp13470)

var ifres13467 Obj

if True == tmp13471 {
ifres13467 = True


} else {
ifres13467 = False


}

ifres13466 = ifres13467


} else {
ifres13466 = False


}

var ifres13465 Obj

if True == ifres13466 {
ifres13465 = True


} else {
ifres13465 = False


}

ifres13464 = ifres13465


} else {
ifres13464 = False


}

var ifres13463 Obj

if True == ifres13464 {
ifres13463 = True


} else {
ifres13463 = False


}

ifres13462 = ifres13463


} else {
ifres13462 = False


}

var ifres13461 Obj

if True == ifres13462 {
ifres13461 = True


} else {
ifres13461 = False


}

ifres13460 = ifres13461


} else {
ifres13460 = False


}

if True == ifres13460 {
tmp13332 := PrimHead(V3325)

tmp13333 := PrimTail(tmp13332)

tmp13334 := PrimCons(symwhen, tmp13333)

tmp13335 := PrimTail(V3325)

tmp13336 := Call(__e, PrimFunc(symshen_4side_1conditions_1_6goals), V3322, V3323, V3324, tmp13335, V3326)


__e.Return(PrimCons(tmp13334, tmp13336))
return


} else {
tmp13458 := PrimIsPair(V3325)

var ifres13431 Obj

if True == tmp13458 {
tmp13456 := PrimHead(V3325)

tmp13457 := PrimIsPair(tmp13456)

var ifres13433 Obj

if True == tmp13457 {
tmp13453 := PrimHead(V3325)

tmp13454 := PrimHead(tmp13453)

tmp13455 := PrimEqual(symlet, tmp13454)

var ifres13435 Obj

if True == tmp13455 {
tmp13450 := PrimHead(V3325)

tmp13451 := PrimTail(tmp13450)

tmp13452 := PrimIsPair(tmp13451)

var ifres13437 Obj

if True == tmp13452 {
tmp13446 := PrimHead(V3325)

tmp13447 := PrimTail(tmp13446)

tmp13448 := PrimTail(tmp13447)

tmp13449 := PrimIsPair(tmp13448)

var ifres13439 Obj

if True == tmp13449 {
tmp13441 := PrimHead(V3325)

tmp13442 := PrimTail(tmp13441)

tmp13443 := PrimTail(tmp13442)

tmp13444 := PrimTail(tmp13443)

tmp13445 := PrimEqual(Nil, tmp13444)

var ifres13440 Obj

if True == tmp13445 {
ifres13440 = True


} else {
ifres13440 = False


}

ifres13439 = ifres13440


} else {
ifres13439 = False


}

var ifres13438 Obj

if True == ifres13439 {
ifres13438 = True


} else {
ifres13438 = False


}

ifres13437 = ifres13438


} else {
ifres13437 = False


}

var ifres13436 Obj

if True == ifres13437 {
ifres13436 = True


} else {
ifres13436 = False


}

ifres13435 = ifres13436


} else {
ifres13435 = False


}

var ifres13434 Obj

if True == ifres13435 {
ifres13434 = True


} else {
ifres13434 = False


}

ifres13433 = ifres13434


} else {
ifres13433 = False


}

var ifres13432 Obj

if True == ifres13433 {
ifres13432 = True


} else {
ifres13432 = False


}

ifres13431 = ifres13432


} else {
ifres13431 = False


}

if True == ifres13431 {
tmp13352 := PrimHead(V3325)

tmp13353 := PrimTail(tmp13352)

tmp13354 := PrimHead(tmp13353)

tmp13355 := Call(__e, PrimFunc(symelement_2), tmp13354, V3323)


if True == tmp13355 {
tmp13337 := PrimHead(V3325)

tmp13338 := PrimTail(tmp13337)

tmp13339 := PrimCons(symis_b, tmp13338)

tmp13340 := PrimTail(V3325)

tmp13341 := Call(__e, PrimFunc(symshen_4side_1conditions_1_6goals), V3322, V3323, V3324, tmp13340, V3326)


__e.Return(PrimCons(tmp13339, tmp13341))
return


} else {
tmp13342 := PrimHead(V3325)

tmp13343 := PrimTail(tmp13342)

tmp13344 := PrimCons(symbind, tmp13343)

tmp13345 := PrimHead(V3325)

tmp13346 := PrimTail(tmp13345)

tmp13347 := PrimHead(tmp13346)

tmp13348 := PrimCons(tmp13347, V3323)

tmp13349 := PrimTail(V3325)

tmp13350 := Call(__e, PrimFunc(symshen_4side_1conditions_1_6goals), V3322, tmp13348, V3324, tmp13349, V3326)


__e.Return(PrimCons(tmp13344, tmp13350))
return


}


} else {
tmp13429 := PrimIsPair(V3325)

var ifres13409 Obj

if True == tmp13429 {
tmp13427 := PrimHead(V3325)

tmp13428 := PrimIsPair(tmp13427)

var ifres13411 Obj

if True == tmp13428 {
tmp13424 := PrimHead(V3325)

tmp13425 := PrimHead(tmp13424)

tmp13426 := PrimEqual(symctxt, tmp13425)

var ifres13413 Obj

if True == tmp13426 {
tmp13421 := PrimHead(V3325)

tmp13422 := PrimTail(tmp13421)

tmp13423 := PrimIsPair(tmp13422)

var ifres13415 Obj

if True == tmp13423 {
tmp13417 := PrimHead(V3325)

tmp13418 := PrimTail(tmp13417)

tmp13419 := PrimTail(tmp13418)

tmp13420 := PrimEqual(Nil, tmp13419)

var ifres13416 Obj

if True == tmp13420 {
ifres13416 = True


} else {
ifres13416 = False


}

ifres13415 = ifres13416


} else {
ifres13415 = False


}

var ifres13414 Obj

if True == ifres13415 {
ifres13414 = True


} else {
ifres13414 = False


}

ifres13413 = ifres13414


} else {
ifres13413 = False


}

var ifres13412 Obj

if True == ifres13413 {
ifres13412 = True


} else {
ifres13412 = False


}

ifres13411 = ifres13412


} else {
ifres13411 = False


}

var ifres13410 Obj

if True == ifres13411 {
ifres13410 = True


} else {
ifres13410 = False


}

ifres13409 = ifres13410


} else {
ifres13409 = False


}

if True == ifres13409 {
tmp13381 := PrimHead(V3325)

tmp13382 := PrimTail(tmp13381)

tmp13383 := PrimHead(tmp13382)

tmp13384 := Call(__e, PrimFunc(symelement_2), tmp13383, V3323)


if True == tmp13384 {
tmp13356 := PrimHead(V3325)

tmp13357 := PrimTail(tmp13356)

tmp13358 := PrimHead(tmp13357)

tmp13359 := PrimCons(tmp13358, V3322)

tmp13360 := PrimTail(V3325)

__e.TailApply(PrimFunc(symshen_4side_1conditions_1_6goals), tmp13359, V3323, V3324, tmp13360, V3326)
return


} else {
tmp13361 := PrimHead(V3325)

tmp13362 := PrimTail(tmp13361)

tmp13363 := PrimHead(tmp13362)

tmp13364 := PrimCons(V3324, Nil)

tmp13365 := PrimCons(tmp13363, tmp13364)

tmp13366 := PrimCons(symbind, tmp13365)

tmp13367 := PrimHead(V3325)

tmp13368 := PrimTail(tmp13367)

tmp13369 := PrimHead(tmp13368)

tmp13370 := PrimCons(tmp13369, V3322)

tmp13371 := PrimHead(V3325)

tmp13372 := PrimTail(tmp13371)

tmp13373 := PrimHead(tmp13372)

tmp13374 := PrimCons(tmp13373, V3323)

tmp13375 := PrimHead(V3325)

tmp13376 := PrimTail(tmp13375)

tmp13377 := PrimHead(tmp13376)

tmp13378 := PrimTail(V3325)

tmp13379 := Call(__e, PrimFunc(symshen_4side_1conditions_1_6goals), tmp13370, tmp13374, tmp13377, tmp13378, V3326)


__e.Return(PrimCons(tmp13366, tmp13379))
return


}


} else {
tmp13407 := PrimIsPair(V3325)

var ifres13387 Obj

if True == tmp13407 {
tmp13405 := PrimHead(V3325)

tmp13406 := PrimIsPair(tmp13405)

var ifres13389 Obj

if True == tmp13406 {
tmp13402 := PrimHead(V3325)

tmp13403 := PrimHead(tmp13402)

tmp13404 := PrimEqual(symsqts, tmp13403)

var ifres13391 Obj

if True == tmp13404 {
tmp13399 := PrimHead(V3325)

tmp13400 := PrimTail(tmp13399)

tmp13401 := PrimIsPair(tmp13400)

var ifres13393 Obj

if True == tmp13401 {
tmp13395 := PrimHead(V3325)

tmp13396 := PrimTail(tmp13395)

tmp13397 := PrimTail(tmp13396)

tmp13398 := PrimEqual(Nil, tmp13397)

var ifres13394 Obj

if True == tmp13398 {
ifres13394 = True


} else {
ifres13394 = False


}

ifres13393 = ifres13394


} else {
ifres13393 = False


}

var ifres13392 Obj

if True == ifres13393 {
ifres13392 = True


} else {
ifres13392 = False


}

ifres13391 = ifres13392


} else {
ifres13391 = False


}

var ifres13390 Obj

if True == ifres13391 {
ifres13390 = True


} else {
ifres13390 = False


}

ifres13389 = ifres13390


} else {
ifres13389 = False


}

var ifres13388 Obj

if True == ifres13389 {
ifres13388 = True


} else {
ifres13388 = False


}

ifres13387 = ifres13388


} else {
ifres13387 = False


}

if True == ifres13387 {
tmp13385 := PrimTail(V3325)

__e.TailApply(PrimFunc(symshen_4side_1conditions_1_6goals), V3322, V3323, V3324, tmp13385, V3326)
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.side-conditions->goals")))
return
}


}


}


}


}


}, 5)

tmp13483 := Call(__e, ns2_1set, symshen_4side_1conditions_1_6goals, tmp13331)


_ = tmp13483

tmp13484 := MakeNative(func(__e *ControlFlow) {
V3331 := __e.Get(1)
_ = V3331
V3332 := __e.Get(2)
_ = V3332
V3333 := __e.Get(3)
_ = V3333
tmp13534 := PrimEqual(Nil, V3333)

if True == tmp13534 {
tmp13485 := PrimIntern(MakeString(";"))

__e.Return(PrimCons(tmp13485, Nil))
return


} else {
tmp13532 := PrimIsPair(V3333)

var ifres13528 Obj

if True == tmp13532 {
tmp13530 := PrimHead(V3333)

tmp13531 := PrimEqual(sym_b, tmp13530)

var ifres13529 Obj

if True == tmp13531 {
ifres13529 = True


} else {
ifres13529 = False


}

ifres13528 = ifres13529


} else {
ifres13528 = False


}

if True == ifres13528 {
tmp13486 := PrimTail(V3333)

tmp13487 := Call(__e, PrimFunc(symshen_4premises_1_6goals), V3331, V3332, tmp13486)


__e.Return(PrimCons(sym_b, tmp13487))
return


} else {
tmp13526 := PrimIsPair(V3333)

var ifres13522 Obj

if True == tmp13526 {
tmp13524 := PrimHead(V3333)

tmp13525 := PrimEqual(symfail, tmp13524)

var ifres13523 Obj

if True == tmp13525 {
ifres13523 = True


} else {
ifres13523 = False


}

ifres13522 = ifres13523


} else {
ifres13522 = False


}

if True == ifres13522 {
tmp13488 := PrimCons(False, Nil)

tmp13489 := PrimCons(symwhen, tmp13488)

tmp13490 := PrimTail(V3333)

tmp13491 := Call(__e, PrimFunc(symshen_4premises_1_6goals), V3331, V3332, tmp13490)


__e.Return(PrimCons(tmp13489, tmp13491))
return


} else {
tmp13520 := PrimIsPair(V3333)

var ifres13505 Obj

if True == tmp13520 {
tmp13518 := PrimHead(V3333)

tmp13519 := PrimIsPair(tmp13518)

var ifres13507 Obj

if True == tmp13519 {
tmp13515 := PrimHead(V3333)

tmp13516 := PrimTail(tmp13515)

tmp13517 := PrimIsPair(tmp13516)

var ifres13509 Obj

if True == tmp13517 {
tmp13511 := PrimHead(V3333)

tmp13512 := PrimTail(tmp13511)

tmp13513 := PrimTail(tmp13512)

tmp13514 := PrimEqual(Nil, tmp13513)

var ifres13510 Obj

if True == tmp13514 {
ifres13510 = True


} else {
ifres13510 = False


}

ifres13509 = ifres13510


} else {
ifres13509 = False


}

var ifres13508 Obj

if True == ifres13509 {
ifres13508 = True


} else {
ifres13508 = False


}

ifres13507 = ifres13508


} else {
ifres13507 = False


}

var ifres13506 Obj

if True == ifres13507 {
ifres13506 = True


} else {
ifres13506 = False


}

ifres13505 = ifres13506


} else {
ifres13505 = False


}

if True == ifres13505 {
tmp13492 := PrimHead(V3333)

tmp13493 := PrimTail(tmp13492)

tmp13494 := PrimHead(tmp13493)

tmp13495 := Call(__e, PrimFunc(symshen_4macro_1_8c), tmp13494)


tmp13496 := PrimHead(V3333)

tmp13497 := PrimHead(tmp13496)

tmp13498 := Call(__e, PrimFunc(symshen_4construct_1context), V3331, tmp13497, V3332)


tmp13499 := PrimCons(tmp13498, Nil)

tmp13500 := PrimCons(tmp13495, tmp13499)

tmp13501 := PrimCons(symshen_4system_1S, tmp13500)

tmp13502 := PrimTail(V3333)

tmp13503 := Call(__e, PrimFunc(symshen_4premises_1_6goals), V3331, V3332, tmp13502)


__e.Return(PrimCons(tmp13501, tmp13503))
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.premises->goals")))
return
}


}


}


}


}, 3)

tmp13535 := Call(__e, ns2_1set, symshen_4premises_1_6goals, tmp13484)


_ = tmp13535

tmp13536 := MakeNative(func(__e *ControlFlow) {
V3337 := __e.Get(1)
_ = V3337
V3338 := __e.Get(2)
_ = V3338
V3339 := __e.Get(3)
_ = V3339
tmp13556 := PrimEqual(Nil, V3338)

if True == tmp13556 {
__e.Return(V3339)
return
} else {
tmp13554 := PrimIsPair(V3338)

var ifres13546 Obj

if True == tmp13554 {
tmp13552 := PrimTail(V3338)

tmp13553 := PrimEqual(Nil, tmp13552)

var ifres13548 Obj

if True == tmp13553 {
tmp13550 := PrimHead(V3338)

tmp13551 := Call(__e, PrimFunc(symelement_2), tmp13550, V3337)


var ifres13549 Obj

if True == tmp13551 {
ifres13549 = True


} else {
ifres13549 = False


}

ifres13548 = ifres13549


} else {
ifres13548 = False


}

var ifres13547 Obj

if True == ifres13548 {
ifres13547 = True


} else {
ifres13547 = False


}

ifres13546 = ifres13547


} else {
ifres13546 = False


}

if True == ifres13546 {
__e.Return(PrimHead(V3338))
return
} else {
tmp13544 := PrimIsPair(V3338)

if True == tmp13544 {
tmp13537 := PrimHead(V3338)

tmp13538 := Call(__e, PrimFunc(symshen_4macro_1_8c), tmp13537)


tmp13539 := PrimTail(V3338)

tmp13540 := Call(__e, PrimFunc(symshen_4construct_1context), V3337, tmp13539, V3339)


tmp13541 := PrimCons(tmp13540, Nil)

tmp13542 := PrimCons(tmp13538, tmp13541)

__e.Return(PrimCons(symcons, tmp13542))
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.construct-context")))
return
}


}


}


}, 3)

tmp13557 := Call(__e, ns2_1set, symshen_4construct_1context, tmp13536)


_ = tmp13557

tmp13558 := MakeNative(func(__e *ControlFlow) {
V3340 := __e.Get(1)
_ = V3340
tmp13559 := MakeNative(func(__e *ControlFlow) {
W3341 := __e.Get(1)
_ = W3341
tmp13560 := MakeNative(func(__e *ControlFlow) {
W3343 := __e.Get(1)
_ = W3343
tmp13561 := MakeNative(func(__e *ControlFlow) {
W3344 := __e.Get(1)
_ = W3344
tmp13562 := MakeNative(func(__e *ControlFlow) {
W3345 := __e.Get(1)
_ = W3345
__e.TailApply(PrimFunc(symshen_4show_1datatypes), W3345)
return
}, 1)

tmp13563 := PrimSet(symshen_4_ddatatypes_d, W3344)

__e.TailApply(tmp13562, tmp13563)
return


}, 1)

tmp13564 := Call(__e, PrimFunc(symshen_4remove_1datatypes), W3341, W3343)


__e.TailApply(tmp13561, tmp13564)
return


}, 1)

tmp13565 := PrimValue(symshen_4_ddatatypes_d)

__e.TailApply(tmp13560, tmp13565)
return


}, 1)

tmp13566 := MakeNative(func(__e *ControlFlow) {
Z3342 := __e.Get(1)
_ = Z3342
__e.TailApply(PrimFunc(symshen_4intern_1type), Z3342)
return
}, 1)

tmp13567 := Call(__e, PrimFunc(symmap), tmp13566, V3340)


__e.TailApply(tmp13559, tmp13567)
return


}, 1)

tmp13568 := Call(__e, ns2_1set, sympreclude, tmp13558)


_ = tmp13568

tmp13569 := MakeNative(func(__e *ControlFlow) {
V3350 := __e.Get(1)
_ = V3350
V3351 := __e.Get(2)
_ = V3351
tmp13576 := PrimEqual(Nil, V3350)

if True == tmp13576 {
__e.Return(V3351)
return
} else {
tmp13574 := PrimIsPair(V3350)

if True == tmp13574 {
tmp13570 := PrimTail(V3350)

tmp13571 := PrimHead(V3350)

tmp13572 := Call(__e, PrimFunc(symshen_4unassoc), tmp13571, V3351)


__e.TailApply(PrimFunc(symshen_4remove_1datatypes), tmp13570, tmp13572)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.remove-datatypes")))
return
}


}


}, 2)

tmp13577 := Call(__e, ns2_1set, symshen_4remove_1datatypes, tmp13569)


_ = tmp13577

tmp13578 := MakeNative(func(__e *ControlFlow) {
V3352 := __e.Get(1)
_ = V3352
tmp13579 := MakeNative(func(__e *ControlFlow) {
Z3353 := __e.Get(1)
_ = Z3353
__e.Return(PrimHead(Z3353))
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp13579, V3352)
return


}, 1)

tmp13580 := Call(__e, ns2_1set, symshen_4show_1datatypes, tmp13578)


_ = tmp13580

tmp13581 := MakeNative(func(__e *ControlFlow) {
V3354 := __e.Get(1)
_ = V3354
tmp13582 := MakeNative(func(__e *ControlFlow) {
W3355 := __e.Get(1)
_ = W3355
tmp13583 := MakeNative(func(__e *ControlFlow) {
W3357 := __e.Get(1)
_ = W3357
tmp13584 := MakeNative(func(__e *ControlFlow) {
W3359 := __e.Get(1)
_ = W3359
__e.TailApply(PrimFunc(symshen_4show_1datatypes), W3359)
return
}, 1)

tmp13585 := PrimValue(symshen_4_ddatatypes_d)

__e.TailApply(tmp13584, tmp13585)
return


}, 1)

tmp13586 := MakeNative(func(__e *ControlFlow) {
Z3358 := __e.Get(1)
_ = Z3358
tmp13587 := Call(__e, PrimFunc(symfn), Z3358)


__e.TailApply(PrimFunc(symshen_4remember_1datatype), Z3358, tmp13587)
return


}, 1)

tmp13588 := Call(__e, PrimFunc(symmap), tmp13586, W3355)


__e.TailApply(tmp13583, tmp13588)
return


}, 1)

tmp13589 := MakeNative(func(__e *ControlFlow) {
Z3356 := __e.Get(1)
_ = Z3356
__e.TailApply(PrimFunc(symshen_4intern_1type), Z3356)
return
}, 1)

tmp13590 := Call(__e, PrimFunc(symmap), tmp13589, V3354)


__e.TailApply(tmp13582, tmp13590)
return


}, 1)

tmp13591 := Call(__e, ns2_1set, syminclude, tmp13581)


_ = tmp13591

tmp13592 := MakeNative(func(__e *ControlFlow) {
V3360 := __e.Get(1)
_ = V3360
tmp13593 := MakeNative(func(__e *ControlFlow) {
W3361 := __e.Get(1)
_ = W3361
tmp13594 := MakeNative(func(__e *ControlFlow) {
W3362 := __e.Get(1)
_ = W3362
tmp13595 := MakeNative(func(__e *ControlFlow) {
W3364 := __e.Get(1)
_ = W3364
tmp13596 := PrimValue(symshen_4_ddatatypes_d)

__e.TailApply(PrimFunc(symshen_4show_1datatypes), tmp13596)
return


}, 1)

tmp13597 := MakeNative(func(__e *ControlFlow) {
Z3365 := __e.Get(1)
_ = Z3365
tmp13598 := Call(__e, PrimFunc(symfn), Z3365)


__e.TailApply(PrimFunc(symshen_4remember_1datatype), Z3365, tmp13598)
return


}, 1)

tmp13599 := Call(__e, PrimFunc(symmap), tmp13597, W3362)


__e.TailApply(tmp13595, tmp13599)
return


}, 1)

tmp13600 := MakeNative(func(__e *ControlFlow) {
Z3363 := __e.Get(1)
_ = Z3363
__e.TailApply(PrimFunc(symshen_4intern_1type), Z3363)
return
}, 1)

tmp13601 := Call(__e, PrimFunc(symmap), tmp13600, V3360)


__e.TailApply(tmp13594, tmp13601)
return


}, 1)

tmp13602 := PrimSet(symshen_4_ddatatypes_d, Nil)

__e.TailApply(tmp13593, tmp13602)
return


}, 1)

tmp13603 := Call(__e, ns2_1set, sympreclude_1all_1but, tmp13592)


_ = tmp13603

tmp13604 := MakeNative(func(__e *ControlFlow) {
V3366 := __e.Get(1)
_ = V3366
tmp13605 := MakeNative(func(__e *ControlFlow) {
W3367 := __e.Get(1)
_ = W3367
tmp13606 := MakeNative(func(__e *ControlFlow) {
W3369 := __e.Get(1)
_ = W3369
tmp13607 := MakeNative(func(__e *ControlFlow) {
W3370 := __e.Get(1)
_ = W3370
__e.TailApply(PrimFunc(symshen_4show_1datatypes), W3370)
return
}, 1)

tmp13608 := Call(__e, PrimFunc(symshen_4remove_1datatypes), W3367, W3369)


tmp13609 := PrimSet(symshen_4_ddatatypes_d, tmp13608)

__e.TailApply(tmp13607, tmp13609)
return


}, 1)

tmp13610 := PrimValue(symshen_4_dalldatatypes_d)

__e.TailApply(tmp13606, tmp13610)
return


}, 1)

tmp13611 := MakeNative(func(__e *ControlFlow) {
Z3368 := __e.Get(1)
_ = Z3368
__e.TailApply(PrimFunc(symshen_4intern_1type), Z3368)
return
}, 1)

tmp13612 := Call(__e, PrimFunc(symmap), tmp13611, V3366)


__e.TailApply(tmp13605, tmp13612)
return


}, 1)

__e.TailApply(ns2_1set, syminclude_1all_1but, tmp13604)
return




}, 0)

