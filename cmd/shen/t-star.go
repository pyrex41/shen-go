package main

import . "github.com/pyrex41/shen-go/kl"

var TStarMain = MakeNative(func(__e *ControlFlow) {
tmp14058 := MakeNative(func(__e *ControlFlow) {
V4434 := __e.Get(1)
_ = V4434
V4435 := __e.Get(2)
_ = V4435
tmp14059 := MakeNative(func(__e *ControlFlow) {
W4436 := __e.Get(1)
_ = W4436
tmp14060 := MakeNative(func(__e *ControlFlow) {
W4437 := __e.Get(1)
_ = W4437
tmp14061 := MakeNative(func(__e *ControlFlow) {
W4438 := __e.Get(1)
_ = W4438
tmp14062 := MakeNative(func(__e *ControlFlow) {
Z4439 := __e.Get(1)
_ = Z4439
__e.Return(MakeNative(func(__e *ControlFlow) {
Z4440 := __e.Get(1)
_ = Z4440
__e.Return(MakeNative(func(__e *ControlFlow) {
Z4441 := __e.Get(1)
_ = Z4441
__e.Return(MakeNative(func(__e *ControlFlow) {
Z4442 := __e.Get(1)
_ = Z4442
tmp14063 := MakeNative(func(__e *ControlFlow) {
W4443 := __e.Get(1)
_ = W4443
tmp14064 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14064

tmp14065 := Call(__e, PrimFunc(symshen_4deref), W4436, Z4439)


tmp14066 := Call(__e, PrimFunc(symreceive), tmp14065)


tmp14067 := Call(__e, PrimFunc(symshen_4deref), W4437, Z4439)


tmp14068 := Call(__e, PrimFunc(symreceive), tmp14067)


tmp14069 := MakeNative(func(__e *ControlFlow) {
tmp14070 := Call(__e, PrimFunc(symshen_4deref), W4438, Z4439)


tmp14071 := Call(__e, PrimFunc(symreceive), tmp14070)


tmp14072 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symreturn), W4443, Z4439, Z4440, Z4441, Z4442)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4toplevel_1forms), tmp14071, W4443, Z4439, Z4440, Z4441, tmp14072)
return


}, 0)

tmp14073 := Call(__e, PrimFunc(symshen_4insert_1prolog_1variables), tmp14066, tmp14068, W4443, Z4439, Z4440, Z4441, tmp14069)


__e.TailApply(PrimFunc(symshen_4gc), Z4439, tmp14073)
return


}, 1)

tmp14074 := Call(__e, PrimFunc(symshen_4newpv), Z4439)


__e.TailApply(tmp14063, tmp14074)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp14075 := Call(__e, PrimFunc(symshen_4prolog_1vector))


tmp14076 := Call(__e, tmp14062, tmp14075)


tmp14077 := Call(__e, PrimFunc(symvector), MakeNumber(0))


tmp14078 := Call(__e, PrimFunc(sym_8v), MakeNumber(0), tmp14077)


tmp14079 := Call(__e, PrimFunc(sym_8v), True, tmp14078)


tmp14080 := Call(__e, tmp14076, tmp14079)


tmp14081 := Call(__e, tmp14080, MakeNumber(0))


tmp14082 := MakeNative(func(__e *ControlFlow) {
__e.Return(True)
return
}, 0)

__e.TailApply(tmp14081, tmp14082)
return


}, 1)

tmp14083 := Call(__e, PrimFunc(symshen_4curry), V4434)


__e.TailApply(tmp14061, tmp14083)
return


}, 1)

tmp14084 := Call(__e, PrimFunc(symshen_4rectify_1type), V4435)


__e.TailApply(tmp14060, tmp14084)
return


}, 1)

tmp14085 := Call(__e, PrimFunc(symshen_4extract_1vars), V4435)


__e.TailApply(tmp14059, tmp14085)
return


}, 2)

tmp14086 := Call(__e, ns2_1set, symshen_4typecheck, tmp14058)


_ = tmp14086

tmp14087 := MakeNative(func(__e *ControlFlow) {
V4444 := __e.Get(1)
_ = V4444
V4445 := __e.Get(2)
_ = V4445
V4446 := __e.Get(3)
_ = V4446
V4447 := __e.Get(4)
_ = V4447
V4448 := __e.Get(5)
_ = V4448
V4449 := __e.Get(6)
_ = V4449
V4450 := __e.Get(7)
_ = V4450
tmp14088 := MakeNative(func(__e *ControlFlow) {
W4451 := __e.Get(1)
_ = W4451
tmp14106 := PrimEqual(W4451, False)

if True == tmp14106 {
tmp14104 := Call(__e, PrimFunc(symshen_4unlocked_2), V4448)


if True == tmp14104 {
tmp14089 := MakeNative(func(__e *ControlFlow) {
W4453 := __e.Get(1)
_ = W4453
tmp14101 := PrimIsPair(W4453)

if True == tmp14101 {
tmp14090 := MakeNative(func(__e *ControlFlow) {
W4454 := __e.Get(1)
_ = W4454
tmp14091 := MakeNative(func(__e *ControlFlow) {
W4455 := __e.Get(1)
_ = W4455
tmp14092 := MakeNative(func(__e *ControlFlow) {
W4456 := __e.Get(1)
_ = W4456
tmp14093 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14093

tmp14094 := Call(__e, PrimFunc(symshen_4deref), W4456, V4447)


tmp14095 := Call(__e, PrimFunc(symsubst), tmp14094, W4454, V4445)


tmp14096 := Call(__e, PrimFunc(symshen_4insert_1prolog_1variables), W4455, tmp14095, V4446, V4447, V4448, V4449, V4450)


__e.TailApply(PrimFunc(symshen_4gc), V4447, tmp14096)
return


}, 1)

tmp14097 := Call(__e, PrimFunc(symshen_4newpv), V4447)


__e.TailApply(tmp14092, tmp14097)
return


}, 1)

tmp14098 := PrimTail(W4453)

__e.TailApply(tmp14091, tmp14098)
return


}, 1)

tmp14099 := PrimHead(W4453)

__e.TailApply(tmp14090, tmp14099)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14102 := Call(__e, PrimFunc(symshen_4lazyderef), V4444, V4447)


__e.TailApply(tmp14089, tmp14102)
return


} else {
__e.Return(False)
return
}


} else {
__e.Return(W4451)
return
}


}, 1)

tmp14114 := Call(__e, PrimFunc(symshen_4unlocked_2), V4448)


var ifres14107 Obj

if True == tmp14114 {
tmp14108 := MakeNative(func(__e *ControlFlow) {
W4452 := __e.Get(1)
_ = W4452
tmp14111 := PrimEqual(W4452, Nil)

if True == tmp14111 {
tmp14109 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14109

__e.TailApply(PrimFunc(symis_b), V4445, V4446, V4447, V4448, V4449, V4450)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14112 := Call(__e, PrimFunc(symshen_4lazyderef), V4444, V4447)


tmp14113 := Call(__e, tmp14108, tmp14112)


ifres14107 = tmp14113


} else {
ifres14107 = False


}

__e.TailApply(tmp14088, ifres14107)
return


}, 7)

tmp14115 := Call(__e, ns2_1set, symshen_4insert_1prolog_1variables, tmp14087)


_ = tmp14115

tmp14116 := MakeNative(func(__e *ControlFlow) {
V4457 := __e.Get(1)
_ = V4457
V4458 := __e.Get(2)
_ = V4458
V4459 := __e.Get(3)
_ = V4459
V4460 := __e.Get(4)
_ = V4460
V4461 := __e.Get(5)
_ = V4461
V4462 := __e.Get(6)
_ = V4462
tmp14117 := MakeNative(func(__e *ControlFlow) {
W4463 := __e.Get(1)
_ = W4463
tmp14118 := MakeNative(func(__e *ControlFlow) {
W4464 := __e.Get(1)
_ = W4464
tmp14131 := PrimEqual(W4464, False)

if True == tmp14131 {
tmp14119 := MakeNative(func(__e *ControlFlow) {
W4470 := __e.Get(1)
_ = W4470
tmp14121 := PrimEqual(W4470, False)

if True == tmp14121 {
__e.TailApply(PrimFunc(symshen_4unlock), V4460, W4463)
return
} else {
__e.Return(W4470)
return
}


}, 1)

tmp14129 := Call(__e, PrimFunc(symshen_4unlocked_2), V4460)


var ifres14122 Obj

if True == tmp14129 {
tmp14123 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14123

tmp14124 := PrimIntern(MakeString(":"))

tmp14125 := PrimCons(V4458, Nil)

tmp14126 := PrimCons(tmp14124, tmp14125)

tmp14127 := PrimCons(V4457, tmp14126)

tmp14128 := Call(__e, PrimFunc(symshen_4system_1S), tmp14127, Nil, V4459, V4460, W4463, V4462)


ifres14122 = tmp14128


} else {
ifres14122 = False


}

__e.TailApply(tmp14119, ifres14122)
return


} else {
__e.Return(W4464)
return
}


}, 1)

tmp14160 := Call(__e, PrimFunc(symshen_4unlocked_2), V4460)


var ifres14132 Obj

if True == tmp14160 {
tmp14133 := MakeNative(func(__e *ControlFlow) {
W4465 := __e.Get(1)
_ = W4465
tmp14157 := PrimIsPair(W4465)

if True == tmp14157 {
tmp14134 := MakeNative(func(__e *ControlFlow) {
W4466 := __e.Get(1)
_ = W4466
tmp14153 := PrimEqual(W4466, symdefine)

if True == tmp14153 {
tmp14135 := MakeNative(func(__e *ControlFlow) {
W4467 := __e.Get(1)
_ = W4467
tmp14149 := PrimIsPair(W4467)

if True == tmp14149 {
tmp14136 := MakeNative(func(__e *ControlFlow) {
W4468 := __e.Get(1)
_ = W4468
tmp14137 := MakeNative(func(__e *ControlFlow) {
W4469 := __e.Get(1)
_ = W4469
tmp14138 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14138

tmp14139 := Call(__e, PrimFunc(symshen_4type_1theory_1enabled_2))


tmp14140 := MakeNative(func(__e *ControlFlow) {
tmp14141 := MakeNative(func(__e *ControlFlow) {
tmp14142 := PrimValue(symshen_4_dspy_d)

tmp14143 := MakeNative(func(__e *ControlFlow) {
tmp14144 := PrimCons(W4468, W4469)

tmp14145 := PrimCons(symdefine, tmp14144)

__e.TailApply(PrimFunc(symshen_4t_d), tmp14145, V4458, V4459, V4460, W4463, V4462)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4signal_1def), tmp14142, W4468, V4459, V4460, W4463, tmp14143)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V4459, V4460, W4463, tmp14141)
return


}, 0)

__e.TailApply(PrimFunc(symwhen), tmp14139, V4459, V4460, W4463, tmp14140)
return


}, 1)

tmp14146 := PrimTail(W4467)

__e.TailApply(tmp14137, tmp14146)
return


}, 1)

tmp14147 := PrimHead(W4467)

__e.TailApply(tmp14136, tmp14147)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14150 := PrimTail(W4465)

tmp14151 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14150, V4459)


__e.TailApply(tmp14135, tmp14151)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14154 := PrimHead(W4465)

tmp14155 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14154, V4459)


__e.TailApply(tmp14134, tmp14155)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14158 := Call(__e, PrimFunc(symshen_4lazyderef), V4457, V4459)


tmp14159 := Call(__e, tmp14133, tmp14158)


ifres14132 = tmp14159


} else {
ifres14132 = False


}

__e.TailApply(tmp14118, ifres14132)
return


}, 1)

tmp14161 := PrimNumberAdd(V4461, MakeNumber(1))

__e.TailApply(tmp14117, tmp14161)
return


}, 6)

tmp14162 := Call(__e, ns2_1set, symshen_4toplevel_1forms, tmp14116)


_ = tmp14162

tmp14163 := MakeNative(func(__e *ControlFlow) {
V4471 := __e.Get(1)
_ = V4471
V4472 := __e.Get(2)
_ = V4472
V4473 := __e.Get(3)
_ = V4473
V4474 := __e.Get(4)
_ = V4474
V4475 := __e.Get(5)
_ = V4475
V4476 := __e.Get(6)
_ = V4476
tmp14164 := MakeNative(func(__e *ControlFlow) {
W4477 := __e.Get(1)
_ = W4477
tmp14181 := PrimEqual(W4477, False)

if True == tmp14181 {
tmp14179 := Call(__e, PrimFunc(symshen_4unlocked_2), V4474)


if True == tmp14179 {
tmp14165 := MakeNative(func(__e *ControlFlow) {
W4479 := __e.Get(1)
_ = W4479
tmp14176 := PrimEqual(W4479, True)

if True == tmp14176 {
tmp14166 := MakeNative(func(__e *ControlFlow) {
W4480 := __e.Get(1)
_ = W4480
tmp14167 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14167

tmp14168 := Call(__e, PrimFunc(symshen_4deref), V4472, V4473)


tmp14169 := Call(__e, PrimFunc(symshen_4app), tmp14168, MakeString(")\n"), symshen_4a)


tmp14170 := PrimStringConcat(MakeString("\ntypechecking (fn "), tmp14169)

tmp14171 := Call(__e, PrimFunc(symstoutput))


tmp14172 := Call(__e, PrimFunc(sympr), tmp14170, tmp14171)


tmp14173 := Call(__e, PrimFunc(symis), W4480, tmp14172, V4473, V4474, V4475, V4476)


__e.TailApply(PrimFunc(symshen_4gc), V4473, tmp14173)
return


}, 1)

tmp14174 := Call(__e, PrimFunc(symshen_4newpv), V4473)


__e.TailApply(tmp14166, tmp14174)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14177 := Call(__e, PrimFunc(symshen_4lazyderef), V4471, V4473)


__e.TailApply(tmp14165, tmp14177)
return


} else {
__e.Return(False)
return
}


} else {
__e.Return(W4477)
return
}


}, 1)

tmp14189 := Call(__e, PrimFunc(symshen_4unlocked_2), V4474)


var ifres14182 Obj

if True == tmp14189 {
tmp14183 := MakeNative(func(__e *ControlFlow) {
W4478 := __e.Get(1)
_ = W4478
tmp14186 := PrimEqual(W4478, False)

if True == tmp14186 {
tmp14184 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14184

__e.TailApply(PrimFunc(symthaw), V4476)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14187 := Call(__e, PrimFunc(symshen_4lazyderef), V4471, V4473)


tmp14188 := Call(__e, tmp14183, tmp14187)


ifres14182 = tmp14188


} else {
ifres14182 = False


}

__e.TailApply(tmp14164, ifres14182)
return


}, 6)

tmp14190 := Call(__e, ns2_1set, symshen_4signal_1def, tmp14163)


_ = tmp14190

tmp14191 := MakeNative(func(__e *ControlFlow) {
V4481 := __e.Get(1)
_ = V4481
tmp14192 := Call(__e, PrimFunc(symshen_4curry_1type), V4481)


__e.TailApply(PrimFunc(symshen_4demodulate), tmp14192)
return


}, 1)

tmp14193 := Call(__e, ns2_1set, symshen_4rectify_1type, tmp14191)


_ = tmp14193

tmp14194 := MakeNative(func(__e *ControlFlow) {
V4482 := __e.Get(1)
_ = V4482
tmp14195 := MakeNative(func(__e *ControlFlow) {
tmp14196 := MakeNative(func(__e *ControlFlow) {
W4483 := __e.Get(1)
_ = W4483
tmp14198 := PrimEqual(W4483, V4482)

if True == tmp14198 {
__e.Return(V4482)
return
} else {
__e.TailApply(PrimFunc(symshen_4demodulate), W4483)
return
}


}, 1)

tmp14199 := MakeNative(func(__e *ControlFlow) {
Z4484 := __e.Get(1)
_ = Z4484
__e.TailApply(PrimFunc(symshen_4demod), Z4484)
return
}, 1)

tmp14200 := Call(__e, PrimFunc(symshen_4walk), tmp14199, V4482)


__e.TailApply(tmp14196, tmp14200)
return


}, 0)

tmp14201 := MakeNative(func(__e *ControlFlow) {
Z4485 := __e.Get(1)
_ = Z4485
__e.Return(V4482)
return
}, 1)

__e.TailApply(try_1catch, tmp14195, tmp14201)
return


}, 1)

tmp14202 := Call(__e, ns2_1set, symshen_4demodulate, tmp14194)


_ = tmp14202

tmp14203 := MakeNative(func(__e *ControlFlow) {
V4486 := __e.Get(1)
_ = V4486
tmp14327 := PrimIsPair(V4486)

var ifres14300 Obj

if True == tmp14327 {
tmp14325 := PrimTail(V4486)

tmp14326 := PrimIsPair(tmp14325)

var ifres14302 Obj

if True == tmp14326 {
tmp14322 := PrimTail(V4486)

tmp14323 := PrimHead(tmp14322)

tmp14324 := PrimEqual(sym_1_1_6, tmp14323)

var ifres14304 Obj

if True == tmp14324 {
tmp14319 := PrimTail(V4486)

tmp14320 := PrimTail(tmp14319)

tmp14321 := PrimIsPair(tmp14320)

var ifres14306 Obj

if True == tmp14321 {
tmp14315 := PrimTail(V4486)

tmp14316 := PrimTail(tmp14315)

tmp14317 := PrimTail(tmp14316)

tmp14318 := PrimIsPair(tmp14317)

var ifres14308 Obj

if True == tmp14318 {
tmp14310 := PrimTail(V4486)

tmp14311 := PrimTail(tmp14310)

tmp14312 := PrimTail(tmp14311)

tmp14313 := PrimHead(tmp14312)

tmp14314 := PrimEqual(sym_1_1_6, tmp14313)

var ifres14309 Obj

if True == tmp14314 {
ifres14309 = True


} else {
ifres14309 = False


}

ifres14308 = ifres14309


} else {
ifres14308 = False


}

var ifres14307 Obj

if True == ifres14308 {
ifres14307 = True


} else {
ifres14307 = False


}

ifres14306 = ifres14307


} else {
ifres14306 = False


}

var ifres14305 Obj

if True == ifres14306 {
ifres14305 = True


} else {
ifres14305 = False


}

ifres14304 = ifres14305


} else {
ifres14304 = False


}

var ifres14303 Obj

if True == ifres14304 {
ifres14303 = True


} else {
ifres14303 = False


}

ifres14302 = ifres14303


} else {
ifres14302 = False


}

var ifres14301 Obj

if True == ifres14302 {
ifres14301 = True


} else {
ifres14301 = False


}

ifres14300 = ifres14301


} else {
ifres14300 = False


}

if True == ifres14300 {
tmp14204 := PrimHead(V4486)

tmp14205 := PrimTail(V4486)

tmp14206 := PrimTail(tmp14205)

tmp14207 := PrimCons(tmp14206, Nil)

tmp14208 := PrimCons(sym_1_1_6, tmp14207)

tmp14209 := PrimCons(tmp14204, tmp14208)

__e.TailApply(PrimFunc(symshen_4curry_1type), tmp14209)
return


} else {
tmp14298 := PrimIsPair(V4486)

var ifres14258 Obj

if True == tmp14298 {
tmp14296 := PrimHead(V4486)

tmp14297 := PrimIsPair(tmp14296)

var ifres14260 Obj

if True == tmp14297 {
tmp14293 := PrimHead(V4486)

tmp14294 := PrimHead(tmp14293)

tmp14295 := PrimEqual(symlist, tmp14294)

var ifres14262 Obj

if True == tmp14295 {
tmp14290 := PrimHead(V4486)

tmp14291 := PrimTail(tmp14290)

tmp14292 := PrimIsPair(tmp14291)

var ifres14264 Obj

if True == tmp14292 {
tmp14286 := PrimHead(V4486)

tmp14287 := PrimTail(tmp14286)

tmp14288 := PrimTail(tmp14287)

tmp14289 := PrimEqual(Nil, tmp14288)

var ifres14266 Obj

if True == tmp14289 {
tmp14284 := PrimTail(V4486)

tmp14285 := PrimIsPair(tmp14284)

var ifres14268 Obj

if True == tmp14285 {
tmp14281 := PrimTail(V4486)

tmp14282 := PrimHead(tmp14281)

tmp14283 := PrimEqual(sym_a_a_6, tmp14282)

var ifres14270 Obj

if True == tmp14283 {
tmp14278 := PrimTail(V4486)

tmp14279 := PrimTail(tmp14278)

tmp14280 := PrimIsPair(tmp14279)

var ifres14272 Obj

if True == tmp14280 {
tmp14274 := PrimTail(V4486)

tmp14275 := PrimTail(tmp14274)

tmp14276 := PrimTail(tmp14275)

tmp14277 := PrimEqual(Nil, tmp14276)

var ifres14273 Obj

if True == tmp14277 {
ifres14273 = True


} else {
ifres14273 = False


}

ifres14272 = ifres14273


} else {
ifres14272 = False


}

var ifres14271 Obj

if True == ifres14272 {
ifres14271 = True


} else {
ifres14271 = False


}

ifres14270 = ifres14271


} else {
ifres14270 = False


}

var ifres14269 Obj

if True == ifres14270 {
ifres14269 = True


} else {
ifres14269 = False


}

ifres14268 = ifres14269


} else {
ifres14268 = False


}

var ifres14267 Obj

if True == ifres14268 {
ifres14267 = True


} else {
ifres14267 = False


}

ifres14266 = ifres14267


} else {
ifres14266 = False


}

var ifres14265 Obj

if True == ifres14266 {
ifres14265 = True


} else {
ifres14265 = False


}

ifres14264 = ifres14265


} else {
ifres14264 = False


}

var ifres14263 Obj

if True == ifres14264 {
ifres14263 = True


} else {
ifres14263 = False


}

ifres14262 = ifres14263


} else {
ifres14262 = False


}

var ifres14261 Obj

if True == ifres14262 {
ifres14261 = True


} else {
ifres14261 = False


}

ifres14260 = ifres14261


} else {
ifres14260 = False


}

var ifres14259 Obj

if True == ifres14260 {
ifres14259 = True


} else {
ifres14259 = False


}

ifres14258 = ifres14259


} else {
ifres14258 = False


}

if True == ifres14258 {
tmp14210 := PrimHead(V4486)

tmp14211 := PrimHead(V4486)

tmp14212 := PrimTail(V4486)

tmp14213 := PrimTail(tmp14212)

tmp14214 := PrimCons(tmp14211, tmp14213)

tmp14215 := PrimCons(symstr, tmp14214)

tmp14216 := PrimCons(tmp14215, Nil)

tmp14217 := PrimCons(sym_1_1_6, tmp14216)

tmp14218 := PrimCons(tmp14210, tmp14217)

__e.TailApply(PrimFunc(symshen_4curry_1type), tmp14218)
return


} else {
tmp14256 := PrimIsPair(V4486)

var ifres14229 Obj

if True == tmp14256 {
tmp14254 := PrimTail(V4486)

tmp14255 := PrimIsPair(tmp14254)

var ifres14231 Obj

if True == tmp14255 {
tmp14251 := PrimTail(V4486)

tmp14252 := PrimHead(tmp14251)

tmp14253 := PrimEqual(sym_d, tmp14252)

var ifres14233 Obj

if True == tmp14253 {
tmp14248 := PrimTail(V4486)

tmp14249 := PrimTail(tmp14248)

tmp14250 := PrimIsPair(tmp14249)

var ifres14235 Obj

if True == tmp14250 {
tmp14244 := PrimTail(V4486)

tmp14245 := PrimTail(tmp14244)

tmp14246 := PrimTail(tmp14245)

tmp14247 := PrimIsPair(tmp14246)

var ifres14237 Obj

if True == tmp14247 {
tmp14239 := PrimTail(V4486)

tmp14240 := PrimTail(tmp14239)

tmp14241 := PrimTail(tmp14240)

tmp14242 := PrimHead(tmp14241)

tmp14243 := PrimEqual(sym_d, tmp14242)

var ifres14238 Obj

if True == tmp14243 {
ifres14238 = True


} else {
ifres14238 = False


}

ifres14237 = ifres14238


} else {
ifres14237 = False


}

var ifres14236 Obj

if True == ifres14237 {
ifres14236 = True


} else {
ifres14236 = False


}

ifres14235 = ifres14236


} else {
ifres14235 = False


}

var ifres14234 Obj

if True == ifres14235 {
ifres14234 = True


} else {
ifres14234 = False


}

ifres14233 = ifres14234


} else {
ifres14233 = False


}

var ifres14232 Obj

if True == ifres14233 {
ifres14232 = True


} else {
ifres14232 = False


}

ifres14231 = ifres14232


} else {
ifres14231 = False


}

var ifres14230 Obj

if True == ifres14231 {
ifres14230 = True


} else {
ifres14230 = False


}

ifres14229 = ifres14230


} else {
ifres14229 = False


}

if True == ifres14229 {
tmp14219 := PrimHead(V4486)

tmp14220 := PrimTail(V4486)

tmp14221 := PrimTail(tmp14220)

tmp14222 := PrimCons(tmp14221, Nil)

tmp14223 := PrimCons(sym_d, tmp14222)

tmp14224 := PrimCons(tmp14219, tmp14223)

__e.TailApply(PrimFunc(symshen_4curry_1type), tmp14224)
return


} else {
tmp14227 := PrimIsPair(V4486)

if True == tmp14227 {
tmp14225 := MakeNative(func(__e *ControlFlow) {
Z4487 := __e.Get(1)
_ = Z4487
__e.TailApply(PrimFunc(symshen_4curry_1type), Z4487)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp14225, V4486)
return


} else {
__e.Return(V4486)
return
}


}


}


}


}, 1)

tmp14328 := Call(__e, ns2_1set, symshen_4curry_1type, tmp14203)


_ = tmp14328

tmp14329 := MakeNative(func(__e *ControlFlow) {
V4488 := __e.Get(1)
_ = V4488
tmp14418 := PrimIsPair(V4488)

var ifres14410 Obj

if True == tmp14418 {
tmp14416 := PrimHead(V4488)

tmp14417 := PrimEqual(symdefine, tmp14416)

var ifres14412 Obj

if True == tmp14417 {
tmp14414 := PrimTail(V4488)

tmp14415 := PrimIsPair(tmp14414)

var ifres14413 Obj

if True == tmp14415 {
ifres14413 = True


} else {
ifres14413 = False


}

ifres14412 = ifres14413


} else {
ifres14412 = False


}

var ifres14411 Obj

if True == ifres14412 {
ifres14411 = True


} else {
ifres14411 = False


}

ifres14410 = ifres14411


} else {
ifres14410 = False


}

if True == ifres14410 {
__e.Return(V4488)
return
} else {
tmp14408 := PrimIsPair(V4488)

var ifres14389 Obj

if True == tmp14408 {
tmp14406 := PrimHead(V4488)

tmp14407 := PrimEqual(symtype, tmp14406)

var ifres14391 Obj

if True == tmp14407 {
tmp14404 := PrimTail(V4488)

tmp14405 := PrimIsPair(tmp14404)

var ifres14393 Obj

if True == tmp14405 {
tmp14401 := PrimTail(V4488)

tmp14402 := PrimTail(tmp14401)

tmp14403 := PrimIsPair(tmp14402)

var ifres14395 Obj

if True == tmp14403 {
tmp14397 := PrimTail(V4488)

tmp14398 := PrimTail(tmp14397)

tmp14399 := PrimTail(tmp14398)

tmp14400 := PrimEqual(Nil, tmp14399)

var ifres14396 Obj

if True == tmp14400 {
ifres14396 = True


} else {
ifres14396 = False


}

ifres14395 = ifres14396


} else {
ifres14395 = False


}

var ifres14394 Obj

if True == ifres14395 {
ifres14394 = True


} else {
ifres14394 = False


}

ifres14393 = ifres14394


} else {
ifres14393 = False


}

var ifres14392 Obj

if True == ifres14393 {
ifres14392 = True


} else {
ifres14392 = False


}

ifres14391 = ifres14392


} else {
ifres14391 = False


}

var ifres14390 Obj

if True == ifres14391 {
ifres14390 = True


} else {
ifres14390 = False


}

ifres14389 = ifres14390


} else {
ifres14389 = False


}

if True == ifres14389 {
tmp14330 := PrimTail(V4488)

tmp14331 := PrimHead(tmp14330)

tmp14332 := Call(__e, PrimFunc(symshen_4curry), tmp14331)


tmp14333 := PrimTail(V4488)

tmp14334 := PrimTail(tmp14333)

tmp14335 := PrimCons(tmp14332, tmp14334)

__e.Return(PrimCons(symtype, tmp14335))
return


} else {
tmp14387 := PrimIsPair(V4488)

var ifres14383 Obj

if True == tmp14387 {
tmp14385 := PrimHead(V4488)

tmp14386 := Call(__e, PrimFunc(symshen_4special_2), tmp14385)


var ifres14384 Obj

if True == tmp14386 {
ifres14384 = True


} else {
ifres14384 = False


}

ifres14383 = ifres14384


} else {
ifres14383 = False


}

if True == ifres14383 {
tmp14336 := PrimHead(V4488)

tmp14337 := MakeNative(func(__e *ControlFlow) {
Z4489 := __e.Get(1)
_ = Z4489
__e.TailApply(PrimFunc(symshen_4curry), Z4489)
return
}, 1)

tmp14338 := PrimTail(V4488)

tmp14339 := Call(__e, PrimFunc(symmap), tmp14337, tmp14338)


__e.Return(PrimCons(tmp14336, tmp14339))
return


} else {
tmp14381 := PrimIsPair(V4488)

var ifres14377 Obj

if True == tmp14381 {
tmp14379 := PrimHead(V4488)

tmp14380 := Call(__e, PrimFunc(symshen_4extraspecial_2), tmp14379)


var ifres14378 Obj

if True == tmp14380 {
ifres14378 = True


} else {
ifres14378 = False


}

ifres14377 = ifres14378


} else {
ifres14377 = False


}

if True == ifres14377 {
__e.Return(V4488)
return
} else {
tmp14375 := PrimIsPair(V4488)

var ifres14366 Obj

if True == tmp14375 {
tmp14373 := PrimTail(V4488)

tmp14374 := PrimIsPair(tmp14373)

var ifres14368 Obj

if True == tmp14374 {
tmp14370 := PrimTail(V4488)

tmp14371 := PrimTail(tmp14370)

tmp14372 := PrimIsPair(tmp14371)

var ifres14369 Obj

if True == tmp14372 {
ifres14369 = True


} else {
ifres14369 = False


}

ifres14368 = ifres14369


} else {
ifres14368 = False


}

var ifres14367 Obj

if True == ifres14368 {
ifres14367 = True


} else {
ifres14367 = False


}

ifres14366 = ifres14367


} else {
ifres14366 = False


}

if True == ifres14366 {
tmp14340 := PrimHead(V4488)

tmp14341 := PrimTail(V4488)

tmp14342 := PrimHead(tmp14341)

tmp14343 := PrimCons(tmp14342, Nil)

tmp14344 := PrimCons(tmp14340, tmp14343)

tmp14345 := PrimTail(V4488)

tmp14346 := PrimTail(tmp14345)

tmp14347 := PrimCons(tmp14344, tmp14346)

__e.TailApply(PrimFunc(symshen_4curry), tmp14347)
return


} else {
tmp14364 := PrimIsPair(V4488)

var ifres14355 Obj

if True == tmp14364 {
tmp14362 := PrimTail(V4488)

tmp14363 := PrimIsPair(tmp14362)

var ifres14357 Obj

if True == tmp14363 {
tmp14359 := PrimTail(V4488)

tmp14360 := PrimTail(tmp14359)

tmp14361 := PrimEqual(Nil, tmp14360)

var ifres14358 Obj

if True == tmp14361 {
ifres14358 = True


} else {
ifres14358 = False


}

ifres14357 = ifres14358


} else {
ifres14357 = False


}

var ifres14356 Obj

if True == ifres14357 {
ifres14356 = True


} else {
ifres14356 = False


}

ifres14355 = ifres14356


} else {
ifres14355 = False


}

if True == ifres14355 {
tmp14348 := PrimHead(V4488)

tmp14349 := Call(__e, PrimFunc(symshen_4curry), tmp14348)


tmp14350 := PrimTail(V4488)

tmp14351 := PrimHead(tmp14350)

tmp14352 := Call(__e, PrimFunc(symshen_4curry), tmp14351)


tmp14353 := PrimCons(tmp14352, Nil)

__e.Return(PrimCons(tmp14349, tmp14353))
return


} else {
__e.Return(V4488)
return
}


}


}


}


}


}


}, 1)

tmp14419 := Call(__e, ns2_1set, symshen_4curry, tmp14329)


_ = tmp14419

tmp14420 := MakeNative(func(__e *ControlFlow) {
V4490 := __e.Get(1)
_ = V4490
tmp14421 := PrimValue(symshen_4_dspecial_d)

__e.TailApply(PrimFunc(symelement_2), V4490, tmp14421)
return


}, 1)

tmp14422 := Call(__e, ns2_1set, symshen_4special_2, tmp14420)


_ = tmp14422

tmp14423 := MakeNative(func(__e *ControlFlow) {
V4491 := __e.Get(1)
_ = V4491
tmp14424 := PrimValue(symshen_4_dextraspecial_d)

__e.TailApply(PrimFunc(symelement_2), V4491, tmp14424)
return


}, 1)

tmp14425 := Call(__e, ns2_1set, symshen_4extraspecial_2, tmp14423)


_ = tmp14425

tmp14426 := MakeNative(func(__e *ControlFlow) {
V4492 := __e.Get(1)
_ = V4492
V4493 := __e.Get(2)
_ = V4493
V4494 := __e.Get(3)
_ = V4494
V4495 := __e.Get(4)
_ = V4495
V4496 := __e.Get(5)
_ = V4496
V4497 := __e.Get(6)
_ = V4497
tmp14427 := MakeNative(func(__e *ControlFlow) {
W4498 := __e.Get(1)
_ = W4498
tmp14428 := MakeNative(func(__e *ControlFlow) {
W4499 := __e.Get(1)
_ = W4499
tmp14486 := PrimEqual(W4499, False)

if True == tmp14486 {
tmp14429 := MakeNative(func(__e *ControlFlow) {
W4500 := __e.Get(1)
_ = W4500
tmp14448 := PrimEqual(W4500, False)

if True == tmp14448 {
tmp14430 := MakeNative(func(__e *ControlFlow) {
W4508 := __e.Get(1)
_ = W4508
tmp14440 := PrimEqual(W4508, False)

if True == tmp14440 {
tmp14431 := MakeNative(func(__e *ControlFlow) {
W4509 := __e.Get(1)
_ = W4509
tmp14433 := PrimEqual(W4509, False)

if True == tmp14433 {
__e.TailApply(PrimFunc(symshen_4unlock), V4495, W4498)
return
} else {
__e.Return(W4509)
return
}


}, 1)

tmp14438 := Call(__e, PrimFunc(symshen_4unlocked_2), V4495)


var ifres14434 Obj

if True == tmp14438 {
tmp14435 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14435

tmp14436 := PrimValue(symshen_4_ddatatypes_d)

tmp14437 := Call(__e, PrimFunc(symshen_4search_1user_1datatypes), V4492, V4493, tmp14436, V4494, V4495, W4498, V4497)


ifres14434 = tmp14437


} else {
ifres14434 = False


}

__e.TailApply(tmp14431, ifres14434)
return


} else {
__e.Return(W4508)
return
}


}, 1)

tmp14446 := Call(__e, PrimFunc(symshen_4unlocked_2), V4495)


var ifres14441 Obj

if True == tmp14446 {
tmp14442 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14442

tmp14443 := PrimValue(symshen_4_dspy_d)

tmp14444 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4show), V4492, V4493, V4494, V4495, W4498, V4497)
return
}, 0)

tmp14445 := Call(__e, PrimFunc(symwhen), tmp14443, V4494, V4495, W4498, tmp14444)


ifres14441 = tmp14445


} else {
ifres14441 = False


}

__e.TailApply(tmp14430, ifres14441)
return


} else {
__e.Return(W4500)
return
}


}, 1)

tmp14484 := Call(__e, PrimFunc(symshen_4unlocked_2), V4495)


var ifres14449 Obj

if True == tmp14484 {
tmp14450 := MakeNative(func(__e *ControlFlow) {
W4501 := __e.Get(1)
_ = W4501
tmp14481 := PrimIsPair(W4501)

if True == tmp14481 {
tmp14451 := MakeNative(func(__e *ControlFlow) {
W4502 := __e.Get(1)
_ = W4502
tmp14452 := MakeNative(func(__e *ControlFlow) {
W4503 := __e.Get(1)
_ = W4503
tmp14476 := PrimIsPair(W4503)

if True == tmp14476 {
tmp14453 := MakeNative(func(__e *ControlFlow) {
W4504 := __e.Get(1)
_ = W4504
tmp14454 := MakeNative(func(__e *ControlFlow) {
W4505 := __e.Get(1)
_ = W4505
tmp14471 := PrimIsPair(W4505)

if True == tmp14471 {
tmp14455 := MakeNative(func(__e *ControlFlow) {
W4506 := __e.Get(1)
_ = W4506
tmp14456 := MakeNative(func(__e *ControlFlow) {
W4507 := __e.Get(1)
_ = W4507
tmp14466 := PrimEqual(W4507, Nil)

if True == tmp14466 {
tmp14457 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14457

tmp14458 := Call(__e, PrimFunc(symshen_4deref), W4504, V4494)


tmp14459 := PrimIntern(MakeString(":"))

tmp14460 := PrimEqual(tmp14458, tmp14459)

tmp14461 := MakeNative(func(__e *ControlFlow) {
tmp14462 := Call(__e, PrimFunc(symshen_4type_1theory_1enabled_2))


tmp14463 := MakeNative(func(__e *ControlFlow) {
tmp14464 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4502, W4506, V4493, V4494, V4495, W4498, V4497)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V4494, V4495, W4498, tmp14464)
return


}, 0)

__e.TailApply(PrimFunc(symwhen), tmp14462, V4494, V4495, W4498, tmp14463)
return


}, 0)

__e.TailApply(PrimFunc(symwhen), tmp14460, V4494, V4495, W4498, tmp14461)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14467 := PrimTail(W4505)

tmp14468 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14467, V4494)


__e.TailApply(tmp14456, tmp14468)
return


}, 1)

tmp14469 := PrimHead(W4505)

__e.TailApply(tmp14455, tmp14469)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14472 := PrimTail(W4503)

tmp14473 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14472, V4494)


__e.TailApply(tmp14454, tmp14473)
return


}, 1)

tmp14474 := PrimHead(W4503)

__e.TailApply(tmp14453, tmp14474)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14477 := PrimTail(W4501)

tmp14478 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14477, V4494)


__e.TailApply(tmp14452, tmp14478)
return


}, 1)

tmp14479 := PrimHead(W4501)

__e.TailApply(tmp14451, tmp14479)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14482 := Call(__e, PrimFunc(symshen_4lazyderef), V4492, V4494)


tmp14483 := Call(__e, tmp14450, tmp14482)


ifres14449 = tmp14483


} else {
ifres14449 = False


}

__e.TailApply(tmp14429, ifres14449)
return


} else {
__e.Return(W4499)
return
}


}, 1)

tmp14491 := Call(__e, PrimFunc(symshen_4unlocked_2), V4495)


var ifres14487 Obj

if True == tmp14491 {
tmp14488 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14488

tmp14489 := Call(__e, PrimFunc(symshen_4maxinfexceeded_2))


tmp14490 := Call(__e, PrimFunc(symwhen), tmp14489, V4494, V4495, W4498, V4497)


ifres14487 = tmp14490


} else {
ifres14487 = False


}

__e.TailApply(tmp14428, ifres14487)
return


}, 1)

tmp14492 := PrimNumberAdd(V4496, MakeNumber(1))

__e.TailApply(tmp14427, tmp14492)
return


}, 6)

tmp14493 := Call(__e, ns2_1set, symshen_4system_1S, tmp14426)


_ = tmp14493

tmp14494 := MakeNative(func(__e *ControlFlow) {
V4516 := __e.Get(1)
_ = V4516
V4517 := __e.Get(2)
_ = V4517
V4518 := __e.Get(3)
_ = V4518
V4519 := __e.Get(4)
_ = V4519
V4520 := __e.Get(5)
_ = V4520
V4521 := __e.Get(6)
_ = V4521
tmp14495 := Call(__e, PrimFunc(symshen_4line))


_ = tmp14495

tmp14496 := Call(__e, PrimFunc(symshen_4deref), V4516, V4518)


tmp14497 := Call(__e, PrimFunc(symshen_4show_1p), tmp14496)


_ = tmp14497

tmp14498 := Call(__e, PrimFunc(symnl), MakeNumber(2))


_ = tmp14498

tmp14499 := Call(__e, PrimFunc(symshen_4deref), V4517, V4518)


tmp14500 := Call(__e, PrimFunc(symshen_4show_1assumptions), tmp14499, MakeNumber(1))


_ = tmp14500

tmp14501 := Call(__e, PrimFunc(symshen_4pause_1for_1user))


_ = tmp14501

__e.Return(False)
return


}, 6)

tmp14502 := Call(__e, ns2_1set, symshen_4show, tmp14494)


_ = tmp14502

tmp14503 := MakeNative(func(__e *ControlFlow) {
tmp14504 := MakeNative(func(__e *ControlFlow) {
W4522 := __e.Get(1)
_ = W4522
tmp14506 := PrimEqual(MakeNumber(1), W4522)

var ifres14505 Obj

if True == tmp14506 {
ifres14505 = MakeString("")


} else {
ifres14505 = MakeString("s")


}

tmp14507 := Call(__e, PrimFunc(symshen_4app), ifres14505, MakeString(" \n?- "), symshen_4a)


tmp14508 := PrimStringConcat(MakeString(" inference"), tmp14507)

tmp14509 := Call(__e, PrimFunc(symshen_4app), W4522, tmp14508, symshen_4a)


tmp14510 := PrimStringConcat(MakeString("____________________________________________________________ "), tmp14509)

tmp14511 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp14510, tmp14511)
return


}, 1)

tmp14512 := Call(__e, PrimFunc(syminferences))


__e.TailApply(tmp14504, tmp14512)
return


}, 0)

tmp14513 := Call(__e, ns2_1set, symshen_4line, tmp14503)


_ = tmp14513

tmp14514 := MakeNative(func(__e *ControlFlow) {
V4523 := __e.Get(1)
_ = V4523
tmp14546 := PrimIsPair(V4523)

var ifres14525 Obj

if True == tmp14546 {
tmp14544 := PrimTail(V4523)

tmp14545 := PrimIsPair(tmp14544)

var ifres14527 Obj

if True == tmp14545 {
tmp14541 := PrimTail(V4523)

tmp14542 := PrimTail(tmp14541)

tmp14543 := PrimIsPair(tmp14542)

var ifres14529 Obj

if True == tmp14543 {
tmp14537 := PrimTail(V4523)

tmp14538 := PrimTail(tmp14537)

tmp14539 := PrimTail(tmp14538)

tmp14540 := PrimEqual(Nil, tmp14539)

var ifres14531 Obj

if True == tmp14540 {
tmp14533 := PrimTail(V4523)

tmp14534 := PrimHead(tmp14533)

tmp14535 := PrimIntern(MakeString(":"))

tmp14536 := PrimEqual(tmp14534, tmp14535)

var ifres14532 Obj

if True == tmp14536 {
ifres14532 = True


} else {
ifres14532 = False


}

ifres14531 = ifres14532


} else {
ifres14531 = False


}

var ifres14530 Obj

if True == ifres14531 {
ifres14530 = True


} else {
ifres14530 = False


}

ifres14529 = ifres14530


} else {
ifres14529 = False


}

var ifres14528 Obj

if True == ifres14529 {
ifres14528 = True


} else {
ifres14528 = False


}

ifres14527 = ifres14528


} else {
ifres14527 = False


}

var ifres14526 Obj

if True == ifres14527 {
ifres14526 = True


} else {
ifres14526 = False


}

ifres14525 = ifres14526


} else {
ifres14525 = False


}

if True == ifres14525 {
tmp14515 := PrimHead(V4523)

tmp14516 := Call(__e, PrimFunc(symshen_4prterm), tmp14515)


_ = tmp14516

tmp14517 := Call(__e, PrimFunc(symstoutput))


tmp14518 := Call(__e, PrimFunc(sympr), MakeString(" : "), tmp14517)


_ = tmp14518

tmp14519 := PrimTail(V4523)

tmp14520 := PrimTail(tmp14519)

tmp14521 := PrimHead(tmp14520)

tmp14522 := Call(__e, PrimFunc(symshen_4app), tmp14521, MakeString(""), symshen_4r)


tmp14523 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp14522, tmp14523)
return


} else {
__e.TailApply(PrimFunc(symshen_4prterm), V4523)
return
}


}, 1)

tmp14547 := Call(__e, ns2_1set, symshen_4show_1p, tmp14514)


_ = tmp14547

tmp14548 := MakeNative(func(__e *ControlFlow) {
V4524 := __e.Get(1)
_ = V4524
tmp14591 := PrimIsPair(V4524)

var ifres14572 Obj

if True == tmp14591 {
tmp14589 := PrimHead(V4524)

tmp14590 := PrimEqual(symcons, tmp14589)

var ifres14574 Obj

if True == tmp14590 {
tmp14587 := PrimTail(V4524)

tmp14588 := PrimIsPair(tmp14587)

var ifres14576 Obj

if True == tmp14588 {
tmp14584 := PrimTail(V4524)

tmp14585 := PrimTail(tmp14584)

tmp14586 := PrimIsPair(tmp14585)

var ifres14578 Obj

if True == tmp14586 {
tmp14580 := PrimTail(V4524)

tmp14581 := PrimTail(tmp14580)

tmp14582 := PrimTail(tmp14581)

tmp14583 := PrimEqual(Nil, tmp14582)

var ifres14579 Obj

if True == tmp14583 {
ifres14579 = True


} else {
ifres14579 = False


}

ifres14578 = ifres14579


} else {
ifres14578 = False


}

var ifres14577 Obj

if True == ifres14578 {
ifres14577 = True


} else {
ifres14577 = False


}

ifres14576 = ifres14577


} else {
ifres14576 = False


}

var ifres14575 Obj

if True == ifres14576 {
ifres14575 = True


} else {
ifres14575 = False


}

ifres14574 = ifres14575


} else {
ifres14574 = False


}

var ifres14573 Obj

if True == ifres14574 {
ifres14573 = True


} else {
ifres14573 = False


}

ifres14572 = ifres14573


} else {
ifres14572 = False


}

if True == ifres14572 {
tmp14549 := Call(__e, PrimFunc(symstoutput))


tmp14550 := Call(__e, PrimFunc(sympr), MakeString("["), tmp14549)


_ = tmp14550

tmp14551 := PrimTail(V4524)

tmp14552 := PrimHead(tmp14551)

tmp14553 := Call(__e, PrimFunc(symshen_4prterm), tmp14552)


_ = tmp14553

tmp14554 := PrimTail(V4524)

tmp14555 := PrimTail(tmp14554)

tmp14556 := PrimHead(tmp14555)

tmp14557 := Call(__e, PrimFunc(symshen_4prtl), tmp14556)


_ = tmp14557

tmp14558 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), MakeString("]"), tmp14558)
return


} else {
tmp14570 := PrimIsPair(V4524)

if True == tmp14570 {
tmp14559 := Call(__e, PrimFunc(symstoutput))


tmp14560 := Call(__e, PrimFunc(sympr), MakeString("("), tmp14559)


_ = tmp14560

tmp14561 := PrimHead(V4524)

tmp14562 := Call(__e, PrimFunc(symshen_4prterm), tmp14561)


_ = tmp14562

tmp14563 := MakeNative(func(__e *ControlFlow) {
Z4525 := __e.Get(1)
_ = Z4525
tmp14564 := Call(__e, PrimFunc(symstoutput))


tmp14565 := Call(__e, PrimFunc(sympr), MakeString(" "), tmp14564)


_ = tmp14565

__e.TailApply(PrimFunc(symshen_4prterm), Z4525)
return


}, 1)

tmp14566 := PrimTail(V4524)

tmp14567 := Call(__e, PrimFunc(symmap), tmp14563, tmp14566)


_ = tmp14567

tmp14568 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), MakeString(")"), tmp14568)
return


} else {
__e.TailApply(PrimFunc(symprint), V4524)
return
}


}


}, 1)

tmp14592 := Call(__e, ns2_1set, symshen_4prterm, tmp14548)


_ = tmp14592

tmp14593 := MakeNative(func(__e *ControlFlow) {
V4526 := __e.Get(1)
_ = V4526
tmp14626 := PrimEqual(Nil, V4526)

if True == tmp14626 {
__e.Return(MakeString(""))
return
} else {
tmp14624 := PrimIsPair(V4526)

var ifres14605 Obj

if True == tmp14624 {
tmp14622 := PrimHead(V4526)

tmp14623 := PrimEqual(symcons, tmp14622)

var ifres14607 Obj

if True == tmp14623 {
tmp14620 := PrimTail(V4526)

tmp14621 := PrimIsPair(tmp14620)

var ifres14609 Obj

if True == tmp14621 {
tmp14617 := PrimTail(V4526)

tmp14618 := PrimTail(tmp14617)

tmp14619 := PrimIsPair(tmp14618)

var ifres14611 Obj

if True == tmp14619 {
tmp14613 := PrimTail(V4526)

tmp14614 := PrimTail(tmp14613)

tmp14615 := PrimTail(tmp14614)

tmp14616 := PrimEqual(Nil, tmp14615)

var ifres14612 Obj

if True == tmp14616 {
ifres14612 = True


} else {
ifres14612 = False


}

ifres14611 = ifres14612


} else {
ifres14611 = False


}

var ifres14610 Obj

if True == ifres14611 {
ifres14610 = True


} else {
ifres14610 = False


}

ifres14609 = ifres14610


} else {
ifres14609 = False


}

var ifres14608 Obj

if True == ifres14609 {
ifres14608 = True


} else {
ifres14608 = False


}

ifres14607 = ifres14608


} else {
ifres14607 = False


}

var ifres14606 Obj

if True == ifres14607 {
ifres14606 = True


} else {
ifres14606 = False


}

ifres14605 = ifres14606


} else {
ifres14605 = False


}

if True == ifres14605 {
tmp14594 := Call(__e, PrimFunc(symstoutput))


tmp14595 := Call(__e, PrimFunc(sympr), MakeString(" "), tmp14594)


_ = tmp14595

tmp14596 := PrimTail(V4526)

tmp14597 := PrimHead(tmp14596)

tmp14598 := Call(__e, PrimFunc(symshen_4prterm), tmp14597)


_ = tmp14598

tmp14599 := PrimTail(V4526)

tmp14600 := PrimTail(tmp14599)

tmp14601 := PrimHead(tmp14600)

__e.TailApply(PrimFunc(symshen_4prtl), tmp14601)
return


} else {
tmp14602 := Call(__e, PrimFunc(symstoutput))


tmp14603 := Call(__e, PrimFunc(sympr), MakeString(" | "), tmp14602)


_ = tmp14603

__e.TailApply(PrimFunc(symshen_4prterm), V4526)
return


}


}


}, 1)

tmp14627 := Call(__e, ns2_1set, symshen_4prtl, tmp14593)


_ = tmp14627

tmp14628 := MakeNative(func(__e *ControlFlow) {
V4533 := __e.Get(1)
_ = V4533
V4534 := __e.Get(2)
_ = V4534
tmp14641 := PrimEqual(Nil, V4533)

if True == tmp14641 {
tmp14629 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), MakeString("\n> "), tmp14629)
return


} else {
tmp14639 := PrimIsPair(V4533)

if True == tmp14639 {
tmp14630 := Call(__e, PrimFunc(symshen_4app), V4534, MakeString(". "), symshen_4a)


tmp14631 := Call(__e, PrimFunc(symstoutput))


tmp14632 := Call(__e, PrimFunc(sympr), tmp14630, tmp14631)


_ = tmp14632

tmp14633 := PrimHead(V4533)

tmp14634 := Call(__e, PrimFunc(symshen_4show_1p), tmp14633)


_ = tmp14634

tmp14635 := Call(__e, PrimFunc(symnl), MakeNumber(1))


_ = tmp14635

tmp14636 := PrimTail(V4533)

tmp14637 := PrimNumberAdd(V4534, MakeNumber(1))

__e.TailApply(PrimFunc(symshen_4show_1assumptions), tmp14636, tmp14637)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.show-assumptions")))
return
}


}


}, 2)

tmp14642 := Call(__e, ns2_1set, symshen_4show_1assumptions, tmp14628)


_ = tmp14642

tmp14643 := MakeNative(func(__e *ControlFlow) {
tmp14644 := MakeNative(func(__e *ControlFlow) {
W4535 := __e.Get(1)
_ = W4535
tmp14646 := PrimEqual(W4535, MakeNumber(94))

if True == tmp14646 {
__e.Return(PrimSimpleError(MakeString("input aborted\n")))
return
} else {
__e.TailApply(PrimFunc(symnl), MakeNumber(1))
return
}


}, 1)

tmp14647 := Call(__e, PrimFunc(symstinput))


tmp14648 := PrimReadByte(tmp14647)

__e.TailApply(tmp14644, tmp14648)
return


}, 0)

tmp14649 := Call(__e, ns2_1set, symshen_4pause_1for_1user, tmp14643)


_ = tmp14649

tmp14650 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(symshen_4_dshen_1type_1theory_1enabled_2_d))
return
}, 0)

tmp14651 := Call(__e, ns2_1set, symshen_4type_1theory_1enabled_2, tmp14650)


_ = tmp14651

tmp14652 := MakeNative(func(__e *ControlFlow) {
tmp14654 := Call(__e, PrimFunc(syminferences))


tmp14655 := PrimValue(symshen_4_dmaxinferences_d)

tmp14656 := PrimGreatThan(tmp14654, tmp14655)

if True == tmp14656 {
__e.Return(PrimSimpleError(MakeString("maximum inferences exceeded")))
return
} else {
__e.Return(False)
return
}


}, 0)

tmp14657 := Call(__e, ns2_1set, symshen_4maxinfexceeded_2, tmp14652)


_ = tmp14657

tmp14658 := MakeNative(func(__e *ControlFlow) {
V4536 := __e.Get(1)
_ = V4536
V4537 := __e.Get(2)
_ = V4537
V4538 := __e.Get(3)
_ = V4538
V4539 := __e.Get(4)
_ = V4539
V4540 := __e.Get(5)
_ = V4540
V4541 := __e.Get(6)
_ = V4541
V4542 := __e.Get(7)
_ = V4542
tmp14659 := MakeNative(func(__e *ControlFlow) {
W4543 := __e.Get(1)
_ = W4543
tmp14660 := MakeNative(func(__e *ControlFlow) {
W4544 := __e.Get(1)
_ = W4544
tmp15575 := PrimEqual(W4544, False)

if True == tmp15575 {
tmp14661 := MakeNative(func(__e *ControlFlow) {
W4545 := __e.Get(1)
_ = W4545
tmp15565 := PrimEqual(W4545, False)

if True == tmp15565 {
tmp14662 := MakeNative(func(__e *ControlFlow) {
W4546 := __e.Get(1)
_ = W4546
tmp15559 := PrimEqual(W4546, False)

if True == tmp15559 {
tmp14663 := MakeNative(func(__e *ControlFlow) {
W4547 := __e.Get(1)
_ = W4547
tmp15540 := PrimEqual(W4547, False)

if True == tmp15540 {
tmp14664 := MakeNative(func(__e *ControlFlow) {
W4551 := __e.Get(1)
_ = W4551
tmp15507 := PrimEqual(W4551, False)

if True == tmp15507 {
tmp14665 := MakeNative(func(__e *ControlFlow) {
W4557 := __e.Get(1)
_ = W4557
tmp15480 := PrimEqual(W4557, False)

if True == tmp15480 {
tmp14666 := MakeNative(func(__e *ControlFlow) {
W4563 := __e.Get(1)
_ = W4563
tmp15445 := PrimEqual(W4563, False)

if True == tmp15445 {
tmp14667 := MakeNative(func(__e *ControlFlow) {
W4570 := __e.Get(1)
_ = W4570
tmp15414 := PrimEqual(W4570, False)

if True == tmp15414 {
tmp14668 := MakeNative(func(__e *ControlFlow) {
W4577 := __e.Get(1)
_ = W4577
tmp15329 := PrimEqual(W4577, False)

if True == tmp15329 {
tmp14669 := MakeNative(func(__e *ControlFlow) {
W4598 := __e.Get(1)
_ = W4598
tmp15223 := PrimEqual(W4598, False)

if True == tmp15223 {
tmp14670 := MakeNative(func(__e *ControlFlow) {
W4626 := __e.Get(1)
_ = W4626
tmp15138 := PrimEqual(W4626, False)

if True == tmp15138 {
tmp14671 := MakeNative(func(__e *ControlFlow) {
W4647 := __e.Get(1)
_ = W4647
tmp15095 := PrimEqual(W4647, False)

if True == tmp15095 {
tmp14672 := MakeNative(func(__e *ControlFlow) {
W4657 := __e.Get(1)
_ = W4657
tmp14971 := PrimEqual(W4657, False)

if True == tmp14971 {
tmp14673 := MakeNative(func(__e *ControlFlow) {
W4687 := __e.Get(1)
_ = W4687
tmp14907 := PrimEqual(W4687, False)

if True == tmp14907 {
tmp14674 := MakeNative(func(__e *ControlFlow) {
W4700 := __e.Get(1)
_ = W4700
tmp14819 := PrimEqual(W4700, False)

if True == tmp14819 {
tmp14675 := MakeNative(func(__e *ControlFlow) {
W4721 := __e.Get(1)
_ = W4721
tmp14781 := PrimEqual(W4721, False)

if True == tmp14781 {
tmp14676 := MakeNative(func(__e *ControlFlow) {
W4729 := __e.Get(1)
_ = W4729
tmp14741 := PrimEqual(W4729, False)

if True == tmp14741 {
tmp14677 := MakeNative(func(__e *ControlFlow) {
W4737 := __e.Get(1)
_ = W4737
tmp14703 := PrimEqual(W4737, False)

if True == tmp14703 {
tmp14678 := MakeNative(func(__e *ControlFlow) {
W4745 := __e.Get(1)
_ = W4745
tmp14692 := PrimEqual(W4745, False)

if True == tmp14692 {
tmp14679 := MakeNative(func(__e *ControlFlow) {
W4747 := __e.Get(1)
_ = W4747
tmp14681 := PrimEqual(W4747, False)

if True == tmp14681 {
__e.TailApply(PrimFunc(symshen_4unlock), V4540, W4543)
return
} else {
__e.Return(W4747)
return
}


}, 1)

tmp14690 := Call(__e, PrimFunc(symshen_4unlocked_2), V4540)


var ifres14682 Obj

if True == tmp14690 {
tmp14683 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14683

tmp14684 := PrimIntern(MakeString(":"))

tmp14685 := PrimCons(V4537, Nil)

tmp14686 := PrimCons(tmp14684, tmp14685)

tmp14687 := PrimCons(V4536, tmp14686)

tmp14688 := PrimValue(symshen_4_ddatatypes_d)

tmp14689 := Call(__e, PrimFunc(symshen_4search_1user_1datatypes), tmp14687, V4538, tmp14688, V4539, V4540, W4543, V4542)


ifres14682 = tmp14689


} else {
ifres14682 = False


}

__e.TailApply(tmp14679, ifres14682)
return


} else {
__e.Return(W4745)
return
}


}, 1)

tmp14701 := Call(__e, PrimFunc(symshen_4unlocked_2), V4540)


var ifres14693 Obj

if True == tmp14701 {
tmp14694 := MakeNative(func(__e *ControlFlow) {
W4746 := __e.Get(1)
_ = W4746
tmp14695 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14695

tmp14696 := MakeNative(func(__e *ControlFlow) {
tmp14697 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), V4536, V4537, W4746, V4539, V4540, W4543, V4542)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V4539, V4540, W4543, tmp14697)
return


}, 0)

tmp14698 := Call(__e, PrimFunc(symshen_4l_1rules), V4538, W4746, False, V4539, V4540, W4543, tmp14696)


__e.TailApply(PrimFunc(symshen_4gc), V4539, tmp14698)
return


}, 1)

tmp14699 := Call(__e, PrimFunc(symshen_4newpv), V4539)


tmp14700 := Call(__e, tmp14694, tmp14699)


ifres14693 = tmp14700


} else {
ifres14693 = False


}

__e.TailApply(tmp14678, ifres14693)
return


} else {
__e.Return(W4737)
return
}


}, 1)

tmp14739 := Call(__e, PrimFunc(symshen_4unlocked_2), V4540)


var ifres14704 Obj

if True == tmp14739 {
tmp14705 := MakeNative(func(__e *ControlFlow) {
W4738 := __e.Get(1)
_ = W4738
tmp14736 := PrimIsPair(W4738)

if True == tmp14736 {
tmp14706 := MakeNative(func(__e *ControlFlow) {
W4739 := __e.Get(1)
_ = W4739
tmp14732 := PrimEqual(W4739, symset)

if True == tmp14732 {
tmp14707 := MakeNative(func(__e *ControlFlow) {
W4740 := __e.Get(1)
_ = W4740
tmp14728 := PrimIsPair(W4740)

if True == tmp14728 {
tmp14708 := MakeNative(func(__e *ControlFlow) {
W4741 := __e.Get(1)
_ = W4741
tmp14709 := MakeNative(func(__e *ControlFlow) {
W4742 := __e.Get(1)
_ = W4742
tmp14723 := PrimIsPair(W4742)

if True == tmp14723 {
tmp14710 := MakeNative(func(__e *ControlFlow) {
W4743 := __e.Get(1)
_ = W4743
tmp14711 := MakeNative(func(__e *ControlFlow) {
W4744 := __e.Get(1)
_ = W4744
tmp14718 := PrimEqual(W4744, Nil)

if True == tmp14718 {
tmp14712 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14712

tmp14713 := MakeNative(func(__e *ControlFlow) {
tmp14714 := PrimCons(W4741, Nil)

tmp14715 := PrimCons(symvalue, tmp14714)

tmp14716 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4743, V4537, V4538, V4539, V4540, W4543, V4542)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), tmp14715, V4537, V4538, V4539, V4540, W4543, tmp14716)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4741, symsymbol, V4538, V4539, V4540, W4543, tmp14713)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14719 := PrimTail(W4742)

tmp14720 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14719, V4539)


__e.TailApply(tmp14711, tmp14720)
return


}, 1)

tmp14721 := PrimHead(W4742)

__e.TailApply(tmp14710, tmp14721)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14724 := PrimTail(W4740)

tmp14725 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14724, V4539)


__e.TailApply(tmp14709, tmp14725)
return


}, 1)

tmp14726 := PrimHead(W4740)

__e.TailApply(tmp14708, tmp14726)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14729 := PrimTail(W4738)

tmp14730 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14729, V4539)


__e.TailApply(tmp14707, tmp14730)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14733 := PrimHead(W4738)

tmp14734 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14733, V4539)


__e.TailApply(tmp14706, tmp14734)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14737 := Call(__e, PrimFunc(symshen_4lazyderef), V4536, V4539)


tmp14738 := Call(__e, tmp14705, tmp14737)


ifres14704 = tmp14738


} else {
ifres14704 = False


}

__e.TailApply(tmp14677, ifres14704)
return


} else {
__e.Return(W4729)
return
}


}, 1)

tmp14779 := Call(__e, PrimFunc(symshen_4unlocked_2), V4540)


var ifres14742 Obj

if True == tmp14779 {
tmp14743 := MakeNative(func(__e *ControlFlow) {
W4730 := __e.Get(1)
_ = W4730
tmp14776 := PrimIsPair(W4730)

if True == tmp14776 {
tmp14744 := MakeNative(func(__e *ControlFlow) {
W4731 := __e.Get(1)
_ = W4731
tmp14772 := PrimEqual(W4731, symshen_4input_1h_7)

if True == tmp14772 {
tmp14745 := MakeNative(func(__e *ControlFlow) {
W4732 := __e.Get(1)
_ = W4732
tmp14768 := PrimIsPair(W4732)

if True == tmp14768 {
tmp14746 := MakeNative(func(__e *ControlFlow) {
W4733 := __e.Get(1)
_ = W4733
tmp14747 := MakeNative(func(__e *ControlFlow) {
W4734 := __e.Get(1)
_ = W4734
tmp14763 := PrimIsPair(W4734)

if True == tmp14763 {
tmp14748 := MakeNative(func(__e *ControlFlow) {
W4735 := __e.Get(1)
_ = W4735
tmp14749 := MakeNative(func(__e *ControlFlow) {
W4736 := __e.Get(1)
_ = W4736
tmp14758 := PrimEqual(W4736, Nil)

if True == tmp14758 {
tmp14750 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14750

tmp14751 := Call(__e, PrimFunc(symshen_4deref), W4733, V4539)


tmp14752 := Call(__e, PrimFunc(symshen_4rdecons), tmp14751)


tmp14753 := Call(__e, PrimFunc(symshen_4rectify_1type), tmp14752)


tmp14754 := MakeNative(func(__e *ControlFlow) {
tmp14755 := PrimCons(symin, Nil)

tmp14756 := PrimCons(symstream, tmp14755)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4735, tmp14756, V4538, V4539, V4540, W4543, V4542)
return


}, 0)

__e.TailApply(PrimFunc(symis), V4537, tmp14753, V4539, V4540, W4543, tmp14754)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14759 := PrimTail(W4734)

tmp14760 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14759, V4539)


__e.TailApply(tmp14749, tmp14760)
return


}, 1)

tmp14761 := PrimHead(W4734)

__e.TailApply(tmp14748, tmp14761)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14764 := PrimTail(W4732)

tmp14765 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14764, V4539)


__e.TailApply(tmp14747, tmp14765)
return


}, 1)

tmp14766 := PrimHead(W4732)

__e.TailApply(tmp14746, tmp14766)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14769 := PrimTail(W4730)

tmp14770 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14769, V4539)


__e.TailApply(tmp14745, tmp14770)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14773 := PrimHead(W4730)

tmp14774 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14773, V4539)


__e.TailApply(tmp14744, tmp14774)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14777 := Call(__e, PrimFunc(symshen_4lazyderef), V4536, V4539)


tmp14778 := Call(__e, tmp14743, tmp14777)


ifres14742 = tmp14778


} else {
ifres14742 = False


}

__e.TailApply(tmp14676, ifres14742)
return


} else {
__e.Return(W4721)
return
}


}, 1)

tmp14817 := Call(__e, PrimFunc(symshen_4unlocked_2), V4540)


var ifres14782 Obj

if True == tmp14817 {
tmp14783 := MakeNative(func(__e *ControlFlow) {
W4722 := __e.Get(1)
_ = W4722
tmp14814 := PrimIsPair(W4722)

if True == tmp14814 {
tmp14784 := MakeNative(func(__e *ControlFlow) {
W4723 := __e.Get(1)
_ = W4723
tmp14810 := PrimEqual(W4723, symtype)

if True == tmp14810 {
tmp14785 := MakeNative(func(__e *ControlFlow) {
W4724 := __e.Get(1)
_ = W4724
tmp14806 := PrimIsPair(W4724)

if True == tmp14806 {
tmp14786 := MakeNative(func(__e *ControlFlow) {
W4725 := __e.Get(1)
_ = W4725
tmp14787 := MakeNative(func(__e *ControlFlow) {
W4726 := __e.Get(1)
_ = W4726
tmp14801 := PrimIsPair(W4726)

if True == tmp14801 {
tmp14788 := MakeNative(func(__e *ControlFlow) {
W4727 := __e.Get(1)
_ = W4727
tmp14789 := MakeNative(func(__e *ControlFlow) {
W4728 := __e.Get(1)
_ = W4728
tmp14796 := PrimEqual(W4728, Nil)

if True == tmp14796 {
tmp14790 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14790

tmp14791 := MakeNative(func(__e *ControlFlow) {
tmp14792 := Call(__e, PrimFunc(symshen_4deref), W4727, V4539)


tmp14793 := Call(__e, PrimFunc(symshen_4rectify_1type), tmp14792)


tmp14794 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4725, V4537, V4538, V4539, V4540, W4543, V4542)
return
}, 0)

__e.TailApply(PrimFunc(symis_b), tmp14793, V4537, V4539, V4540, W4543, tmp14794)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V4539, V4540, W4543, tmp14791)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14797 := PrimTail(W4726)

tmp14798 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14797, V4539)


__e.TailApply(tmp14789, tmp14798)
return


}, 1)

tmp14799 := PrimHead(W4726)

__e.TailApply(tmp14788, tmp14799)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14802 := PrimTail(W4724)

tmp14803 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14802, V4539)


__e.TailApply(tmp14787, tmp14803)
return


}, 1)

tmp14804 := PrimHead(W4724)

__e.TailApply(tmp14786, tmp14804)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14807 := PrimTail(W4722)

tmp14808 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14807, V4539)


__e.TailApply(tmp14785, tmp14808)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14811 := PrimHead(W4722)

tmp14812 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14811, V4539)


__e.TailApply(tmp14784, tmp14812)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14815 := Call(__e, PrimFunc(symshen_4lazyderef), V4536, V4539)


tmp14816 := Call(__e, tmp14783, tmp14815)


ifres14782 = tmp14816


} else {
ifres14782 = False


}

__e.TailApply(tmp14675, ifres14782)
return


} else {
__e.Return(W4700)
return
}


}, 1)

tmp14905 := Call(__e, PrimFunc(symshen_4unlocked_2), V4540)


var ifres14820 Obj

if True == tmp14905 {
tmp14821 := MakeNative(func(__e *ControlFlow) {
W4701 := __e.Get(1)
_ = W4701
tmp14902 := PrimIsPair(W4701)

if True == tmp14902 {
tmp14822 := MakeNative(func(__e *ControlFlow) {
W4702 := __e.Get(1)
_ = W4702
tmp14898 := PrimEqual(W4702, symopen)

if True == tmp14898 {
tmp14823 := MakeNative(func(__e *ControlFlow) {
W4703 := __e.Get(1)
_ = W4703
tmp14894 := PrimIsPair(W4703)

if True == tmp14894 {
tmp14824 := MakeNative(func(__e *ControlFlow) {
W4704 := __e.Get(1)
_ = W4704
tmp14825 := MakeNative(func(__e *ControlFlow) {
W4705 := __e.Get(1)
_ = W4705
tmp14889 := PrimIsPair(W4705)

if True == tmp14889 {
tmp14826 := MakeNative(func(__e *ControlFlow) {
W4706 := __e.Get(1)
_ = W4706
tmp14827 := MakeNative(func(__e *ControlFlow) {
W4707 := __e.Get(1)
_ = W4707
tmp14884 := PrimEqual(W4707, Nil)

if True == tmp14884 {
tmp14828 := MakeNative(func(__e *ControlFlow) {
W4708 := __e.Get(1)
_ = W4708
tmp14829 := MakeNative(func(__e *ControlFlow) {
W4709 := __e.Get(1)
_ = W4709
tmp14873 := PrimIsPair(W4708)

if True == tmp14873 {
tmp14830 := MakeNative(func(__e *ControlFlow) {
W4711 := __e.Get(1)
_ = W4711
tmp14831 := MakeNative(func(__e *ControlFlow) {
W4712 := __e.Get(1)
_ = W4712
tmp14835 := PrimEqual(W4711, symstream)

if True == tmp14835 {
__e.TailApply(PrimFunc(symthaw), W4712)
return
} else {
tmp14833 := Call(__e, PrimFunc(symshen_4pvar_2), W4711)


if True == tmp14833 {
__e.TailApply(PrimFunc(symshen_4bind_b), W4711, symstream, V4539, W4712)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp14836 := MakeNative(func(__e *ControlFlow) {
tmp14837 := MakeNative(func(__e *ControlFlow) {
W4713 := __e.Get(1)
_ = W4713
tmp14838 := MakeNative(func(__e *ControlFlow) {
W4714 := __e.Get(1)
_ = W4714
tmp14858 := PrimIsPair(W4713)

if True == tmp14858 {
tmp14839 := MakeNative(func(__e *ControlFlow) {
W4716 := __e.Get(1)
_ = W4716
tmp14840 := MakeNative(func(__e *ControlFlow) {
W4717 := __e.Get(1)
_ = W4717
tmp14841 := MakeNative(func(__e *ControlFlow) {
W4718 := __e.Get(1)
_ = W4718
tmp14845 := PrimEqual(W4717, Nil)

if True == tmp14845 {
__e.TailApply(PrimFunc(symthaw), W4718)
return
} else {
tmp14843 := Call(__e, PrimFunc(symshen_4pvar_2), W4717)


if True == tmp14843 {
__e.TailApply(PrimFunc(symshen_4bind_b), W4717, Nil, V4539, W4718)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp14846 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W4714, W4716)
return
}, 0)

__e.TailApply(tmp14841, tmp14846)
return


}, 1)

tmp14847 := PrimTail(W4713)

tmp14848 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14847, V4539)


__e.TailApply(tmp14840, tmp14848)
return


}, 1)

tmp14849 := PrimHead(W4713)

__e.TailApply(tmp14839, tmp14849)
return


} else {
tmp14856 := Call(__e, PrimFunc(symshen_4pvar_2), W4713)


if True == tmp14856 {
tmp14850 := MakeNative(func(__e *ControlFlow) {
W4719 := __e.Get(1)
_ = W4719
tmp14851 := PrimCons(W4719, Nil)

tmp14852 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W4714, W4719)
return
}, 0)

tmp14853 := Call(__e, PrimFunc(symshen_4bind_b), W4713, tmp14851, V4539, tmp14852)


__e.TailApply(PrimFunc(symshen_4gc), V4539, tmp14853)
return


}, 1)

tmp14854 := Call(__e, PrimFunc(symshen_4newpv), V4539)


__e.TailApply(tmp14850, tmp14854)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp14859 := MakeNative(func(__e *ControlFlow) {
Z4715 := __e.Get(1)
_ = Z4715
__e.TailApply(W4709, Z4715)
return
}, 1)

__e.TailApply(tmp14838, tmp14859)
return


}, 1)

tmp14860 := PrimTail(W4708)

tmp14861 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14860, V4539)


__e.TailApply(tmp14837, tmp14861)
return


}, 0)

__e.TailApply(tmp14831, tmp14836)
return


}, 1)

tmp14862 := PrimHead(W4708)

tmp14863 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14862, V4539)


__e.TailApply(tmp14830, tmp14863)
return


} else {
tmp14871 := Call(__e, PrimFunc(symshen_4pvar_2), W4708)


if True == tmp14871 {
tmp14864 := MakeNative(func(__e *ControlFlow) {
W4720 := __e.Get(1)
_ = W4720
tmp14865 := PrimCons(W4720, Nil)

tmp14866 := PrimCons(symstream, tmp14865)

tmp14867 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W4709, W4720)
return
}, 0)

tmp14868 := Call(__e, PrimFunc(symshen_4bind_b), W4708, tmp14866, V4539, tmp14867)


__e.TailApply(PrimFunc(symshen_4gc), V4539, tmp14868)
return


}, 1)

tmp14869 := Call(__e, PrimFunc(symshen_4newpv), V4539)


__e.TailApply(tmp14864, tmp14869)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp14874 := MakeNative(func(__e *ControlFlow) {
Z4710 := __e.Get(1)
_ = Z4710
tmp14875 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14875

tmp14876 := MakeNative(func(__e *ControlFlow) {
tmp14877 := Call(__e, PrimFunc(symshen_4lazyderef), Z4710, V4539)


tmp14878 := PrimCons(symout, Nil)

tmp14879 := PrimCons(symin, tmp14878)

tmp14880 := Call(__e, PrimFunc(symelement_2), tmp14877, tmp14879)


tmp14881 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4704, symstring, V4538, V4539, V4540, W4543, V4542)
return
}, 0)

__e.TailApply(PrimFunc(symwhen), tmp14880, V4539, V4540, W4543, tmp14881)
return


}, 0)

__e.TailApply(PrimFunc(symis_b), W4706, Z4710, V4539, V4540, W4543, tmp14876)
return


}, 1)

__e.TailApply(tmp14829, tmp14874)
return


}, 1)

tmp14882 := Call(__e, PrimFunc(symshen_4lazyderef), V4537, V4539)


__e.TailApply(tmp14828, tmp14882)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14885 := PrimTail(W4705)

tmp14886 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14885, V4539)


__e.TailApply(tmp14827, tmp14886)
return


}, 1)

tmp14887 := PrimHead(W4705)

__e.TailApply(tmp14826, tmp14887)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14890 := PrimTail(W4703)

tmp14891 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14890, V4539)


__e.TailApply(tmp14825, tmp14891)
return


}, 1)

tmp14892 := PrimHead(W4703)

__e.TailApply(tmp14824, tmp14892)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14895 := PrimTail(W4701)

tmp14896 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14895, V4539)


__e.TailApply(tmp14823, tmp14896)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14899 := PrimHead(W4701)

tmp14900 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14899, V4539)


__e.TailApply(tmp14822, tmp14900)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14903 := Call(__e, PrimFunc(symshen_4lazyderef), V4536, V4539)


tmp14904 := Call(__e, tmp14821, tmp14903)


ifres14820 = tmp14904


} else {
ifres14820 = False


}

__e.TailApply(tmp14674, ifres14820)
return


} else {
__e.Return(W4687)
return
}


}, 1)

tmp14969 := Call(__e, PrimFunc(symshen_4unlocked_2), V4540)


var ifres14908 Obj

if True == tmp14969 {
tmp14909 := MakeNative(func(__e *ControlFlow) {
W4688 := __e.Get(1)
_ = W4688
tmp14966 := PrimIsPair(W4688)

if True == tmp14966 {
tmp14910 := MakeNative(func(__e *ControlFlow) {
W4689 := __e.Get(1)
_ = W4689
tmp14962 := PrimEqual(W4689, symlet)

if True == tmp14962 {
tmp14911 := MakeNative(func(__e *ControlFlow) {
W4690 := __e.Get(1)
_ = W4690
tmp14958 := PrimIsPair(W4690)

if True == tmp14958 {
tmp14912 := MakeNative(func(__e *ControlFlow) {
W4691 := __e.Get(1)
_ = W4691
tmp14913 := MakeNative(func(__e *ControlFlow) {
W4692 := __e.Get(1)
_ = W4692
tmp14953 := PrimIsPair(W4692)

if True == tmp14953 {
tmp14914 := MakeNative(func(__e *ControlFlow) {
W4693 := __e.Get(1)
_ = W4693
tmp14915 := MakeNative(func(__e *ControlFlow) {
W4694 := __e.Get(1)
_ = W4694
tmp14948 := PrimIsPair(W4694)

if True == tmp14948 {
tmp14916 := MakeNative(func(__e *ControlFlow) {
W4695 := __e.Get(1)
_ = W4695
tmp14917 := MakeNative(func(__e *ControlFlow) {
W4696 := __e.Get(1)
_ = W4696
tmp14943 := PrimEqual(W4696, Nil)

if True == tmp14943 {
tmp14918 := MakeNative(func(__e *ControlFlow) {
W4697 := __e.Get(1)
_ = W4697
tmp14919 := MakeNative(func(__e *ControlFlow) {
W4698 := __e.Get(1)
_ = W4698
tmp14920 := MakeNative(func(__e *ControlFlow) {
W4699 := __e.Get(1)
_ = W4699
tmp14921 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14921

tmp14922 := MakeNative(func(__e *ControlFlow) {
tmp14923 := Call(__e, PrimFunc(symshen_4lazyderef), W4691, V4539)


tmp14924 := Call(__e, PrimFunc(symshen_4freshterm), tmp14923)


tmp14925 := MakeNative(func(__e *ControlFlow) {
tmp14926 := Call(__e, PrimFunc(symshen_4lazyderef), W4691, V4539)


tmp14927 := Call(__e, PrimFunc(symshen_4lazyderef), W4698, V4539)


tmp14928 := Call(__e, PrimFunc(symshen_4lazyderef), W4695, V4539)


tmp14929 := Call(__e, PrimFunc(symshen_4beta), tmp14926, tmp14927, tmp14928)


tmp14930 := MakeNative(func(__e *ControlFlow) {
tmp14931 := PrimIntern(MakeString(":"))

tmp14932 := PrimCons(W4699, Nil)

tmp14933 := PrimCons(tmp14931, tmp14932)

tmp14934 := PrimCons(W4698, tmp14933)

tmp14935 := PrimCons(tmp14934, V4538)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4697, V4537, tmp14935, V4539, V4540, W4543, V4542)
return


}, 0)

__e.TailApply(PrimFunc(symbind), W4697, tmp14929, V4539, V4540, W4543, tmp14930)
return


}, 0)

__e.TailApply(PrimFunc(symbind), W4698, tmp14924, V4539, V4540, W4543, tmp14925)
return


}, 0)

tmp14936 := Call(__e, PrimFunc(symshen_4system_1S_1h), W4693, W4699, V4538, V4539, V4540, W4543, tmp14922)


__e.TailApply(PrimFunc(symshen_4gc), V4539, tmp14936)
return


}, 1)

tmp14937 := Call(__e, PrimFunc(symshen_4newpv), V4539)


tmp14938 := Call(__e, tmp14920, tmp14937)


__e.TailApply(PrimFunc(symshen_4gc), V4539, tmp14938)
return


}, 1)

tmp14939 := Call(__e, PrimFunc(symshen_4newpv), V4539)


tmp14940 := Call(__e, tmp14919, tmp14939)


__e.TailApply(PrimFunc(symshen_4gc), V4539, tmp14940)
return


}, 1)

tmp14941 := Call(__e, PrimFunc(symshen_4newpv), V4539)


__e.TailApply(tmp14918, tmp14941)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14944 := PrimTail(W4694)

tmp14945 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14944, V4539)


__e.TailApply(tmp14917, tmp14945)
return


}, 1)

tmp14946 := PrimHead(W4694)

__e.TailApply(tmp14916, tmp14946)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14949 := PrimTail(W4692)

tmp14950 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14949, V4539)


__e.TailApply(tmp14915, tmp14950)
return


}, 1)

tmp14951 := PrimHead(W4692)

__e.TailApply(tmp14914, tmp14951)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14954 := PrimTail(W4690)

tmp14955 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14954, V4539)


__e.TailApply(tmp14913, tmp14955)
return


}, 1)

tmp14956 := PrimHead(W4690)

__e.TailApply(tmp14912, tmp14956)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14959 := PrimTail(W4688)

tmp14960 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14959, V4539)


__e.TailApply(tmp14911, tmp14960)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14963 := PrimHead(W4688)

tmp14964 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14963, V4539)


__e.TailApply(tmp14910, tmp14964)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14967 := Call(__e, PrimFunc(symshen_4lazyderef), V4536, V4539)


tmp14968 := Call(__e, tmp14909, tmp14967)


ifres14908 = tmp14968


} else {
ifres14908 = False


}

__e.TailApply(tmp14673, ifres14908)
return


} else {
__e.Return(W4657)
return
}


}, 1)

tmp15093 := Call(__e, PrimFunc(symshen_4unlocked_2), V4540)


var ifres14972 Obj

if True == tmp15093 {
tmp14973 := MakeNative(func(__e *ControlFlow) {
W4658 := __e.Get(1)
_ = W4658
tmp15090 := PrimIsPair(W4658)

if True == tmp15090 {
tmp14974 := MakeNative(func(__e *ControlFlow) {
W4659 := __e.Get(1)
_ = W4659
tmp15086 := PrimEqual(W4659, symlambda)

if True == tmp15086 {
tmp14975 := MakeNative(func(__e *ControlFlow) {
W4660 := __e.Get(1)
_ = W4660
tmp15082 := PrimIsPair(W4660)

if True == tmp15082 {
tmp14976 := MakeNative(func(__e *ControlFlow) {
W4661 := __e.Get(1)
_ = W4661
tmp14977 := MakeNative(func(__e *ControlFlow) {
W4662 := __e.Get(1)
_ = W4662
tmp15077 := PrimIsPair(W4662)

if True == tmp15077 {
tmp14978 := MakeNative(func(__e *ControlFlow) {
W4663 := __e.Get(1)
_ = W4663
tmp14979 := MakeNative(func(__e *ControlFlow) {
W4664 := __e.Get(1)
_ = W4664
tmp15072 := PrimEqual(W4664, Nil)

if True == tmp15072 {
tmp14980 := MakeNative(func(__e *ControlFlow) {
W4665 := __e.Get(1)
_ = W4665
tmp14981 := MakeNative(func(__e *ControlFlow) {
W4666 := __e.Get(1)
_ = W4666
tmp15048 := PrimIsPair(W4665)

if True == tmp15048 {
tmp14982 := MakeNative(func(__e *ControlFlow) {
W4671 := __e.Get(1)
_ = W4671
tmp14983 := MakeNative(func(__e *ControlFlow) {
W4672 := __e.Get(1)
_ = W4672
tmp14984 := MakeNative(func(__e *ControlFlow) {
W4673 := __e.Get(1)
_ = W4673
tmp15028 := PrimIsPair(W4672)

if True == tmp15028 {
tmp14985 := MakeNative(func(__e *ControlFlow) {
W4675 := __e.Get(1)
_ = W4675
tmp14986 := MakeNative(func(__e *ControlFlow) {
W4676 := __e.Get(1)
_ = W4676
tmp14990 := PrimEqual(W4675, sym_1_1_6)

if True == tmp14990 {
__e.TailApply(PrimFunc(symthaw), W4676)
return
} else {
tmp14988 := Call(__e, PrimFunc(symshen_4pvar_2), W4675)


if True == tmp14988 {
__e.TailApply(PrimFunc(symshen_4bind_b), W4675, sym_1_1_6, V4539, W4676)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp14991 := MakeNative(func(__e *ControlFlow) {
tmp14992 := MakeNative(func(__e *ControlFlow) {
W4677 := __e.Get(1)
_ = W4677
tmp14993 := MakeNative(func(__e *ControlFlow) {
W4678 := __e.Get(1)
_ = W4678
tmp15013 := PrimIsPair(W4677)

if True == tmp15013 {
tmp14994 := MakeNative(func(__e *ControlFlow) {
W4680 := __e.Get(1)
_ = W4680
tmp14995 := MakeNative(func(__e *ControlFlow) {
W4681 := __e.Get(1)
_ = W4681
tmp14996 := MakeNative(func(__e *ControlFlow) {
W4682 := __e.Get(1)
_ = W4682
tmp15000 := PrimEqual(W4681, Nil)

if True == tmp15000 {
__e.TailApply(PrimFunc(symthaw), W4682)
return
} else {
tmp14998 := Call(__e, PrimFunc(symshen_4pvar_2), W4681)


if True == tmp14998 {
__e.TailApply(PrimFunc(symshen_4bind_b), W4681, Nil, V4539, W4682)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp15001 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W4678, W4680)
return
}, 0)

__e.TailApply(tmp14996, tmp15001)
return


}, 1)

tmp15002 := PrimTail(W4677)

tmp15003 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15002, V4539)


__e.TailApply(tmp14995, tmp15003)
return


}, 1)

tmp15004 := PrimHead(W4677)

__e.TailApply(tmp14994, tmp15004)
return


} else {
tmp15011 := Call(__e, PrimFunc(symshen_4pvar_2), W4677)


if True == tmp15011 {
tmp15005 := MakeNative(func(__e *ControlFlow) {
W4683 := __e.Get(1)
_ = W4683
tmp15006 := PrimCons(W4683, Nil)

tmp15007 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W4678, W4683)
return
}, 0)

tmp15008 := Call(__e, PrimFunc(symshen_4bind_b), W4677, tmp15006, V4539, tmp15007)


__e.TailApply(PrimFunc(symshen_4gc), V4539, tmp15008)
return


}, 1)

tmp15009 := Call(__e, PrimFunc(symshen_4newpv), V4539)


__e.TailApply(tmp15005, tmp15009)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp15014 := MakeNative(func(__e *ControlFlow) {
Z4679 := __e.Get(1)
_ = Z4679
__e.TailApply(W4673, Z4679)
return
}, 1)

__e.TailApply(tmp14993, tmp15014)
return


}, 1)

tmp15015 := PrimTail(W4672)

tmp15016 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15015, V4539)


__e.TailApply(tmp14992, tmp15016)
return


}, 0)

__e.TailApply(tmp14986, tmp14991)
return


}, 1)

tmp15017 := PrimHead(W4672)

tmp15018 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15017, V4539)


__e.TailApply(tmp14985, tmp15018)
return


} else {
tmp15026 := Call(__e, PrimFunc(symshen_4pvar_2), W4672)


if True == tmp15026 {
tmp15019 := MakeNative(func(__e *ControlFlow) {
W4684 := __e.Get(1)
_ = W4684
tmp15020 := PrimCons(W4684, Nil)

tmp15021 := PrimCons(sym_1_1_6, tmp15020)

tmp15022 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W4673, W4684)
return
}, 0)

tmp15023 := Call(__e, PrimFunc(symshen_4bind_b), W4672, tmp15021, V4539, tmp15022)


__e.TailApply(PrimFunc(symshen_4gc), V4539, tmp15023)
return


}, 1)

tmp15024 := Call(__e, PrimFunc(symshen_4newpv), V4539)


__e.TailApply(tmp15019, tmp15024)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp15029 := MakeNative(func(__e *ControlFlow) {
Z4674 := __e.Get(1)
_ = Z4674
tmp15030 := Call(__e, W4666, W4671)


__e.TailApply(tmp15030, Z4674)
return


}, 1)

__e.TailApply(tmp14984, tmp15029)
return


}, 1)

tmp15031 := PrimTail(W4665)

tmp15032 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15031, V4539)


__e.TailApply(tmp14983, tmp15032)
return


}, 1)

tmp15033 := PrimHead(W4665)

__e.TailApply(tmp14982, tmp15033)
return


} else {
tmp15046 := Call(__e, PrimFunc(symshen_4pvar_2), W4665)


if True == tmp15046 {
tmp15034 := MakeNative(func(__e *ControlFlow) {
W4685 := __e.Get(1)
_ = W4685
tmp15035 := MakeNative(func(__e *ControlFlow) {
W4686 := __e.Get(1)
_ = W4686
tmp15036 := PrimCons(W4686, Nil)

tmp15037 := PrimCons(sym_1_1_6, tmp15036)

tmp15038 := PrimCons(W4685, tmp15037)

tmp15039 := MakeNative(func(__e *ControlFlow) {
tmp15040 := Call(__e, W4666, W4685)


__e.TailApply(tmp15040, W4686)
return


}, 0)

tmp15041 := Call(__e, PrimFunc(symshen_4bind_b), W4665, tmp15038, V4539, tmp15039)


__e.TailApply(PrimFunc(symshen_4gc), V4539, tmp15041)
return


}, 1)

tmp15042 := Call(__e, PrimFunc(symshen_4newpv), V4539)


tmp15043 := Call(__e, tmp15035, tmp15042)


__e.TailApply(PrimFunc(symshen_4gc), V4539, tmp15043)
return


}, 1)

tmp15044 := Call(__e, PrimFunc(symshen_4newpv), V4539)


__e.TailApply(tmp15034, tmp15044)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp15049 := MakeNative(func(__e *ControlFlow) {
Z4667 := __e.Get(1)
_ = Z4667
__e.Return(MakeNative(func(__e *ControlFlow) {
Z4668 := __e.Get(1)
_ = Z4668
tmp15050 := MakeNative(func(__e *ControlFlow) {
W4669 := __e.Get(1)
_ = W4669
tmp15051 := MakeNative(func(__e *ControlFlow) {
W4670 := __e.Get(1)
_ = W4670
tmp15052 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15052

tmp15053 := Call(__e, PrimFunc(symshen_4lazyderef), W4661, V4539)


tmp15054 := Call(__e, PrimFunc(symshen_4freshterm), tmp15053)


tmp15055 := MakeNative(func(__e *ControlFlow) {
tmp15056 := Call(__e, PrimFunc(symshen_4lazyderef), W4661, V4539)


tmp15057 := Call(__e, PrimFunc(symshen_4deref), W4670, V4539)


tmp15058 := Call(__e, PrimFunc(symshen_4deref), W4663, V4539)


tmp15059 := Call(__e, PrimFunc(symshen_4beta), tmp15056, tmp15057, tmp15058)


tmp15060 := MakeNative(func(__e *ControlFlow) {
tmp15061 := PrimIntern(MakeString(":"))

tmp15062 := PrimCons(Z4667, Nil)

tmp15063 := PrimCons(tmp15061, tmp15062)

tmp15064 := PrimCons(W4670, tmp15063)

tmp15065 := PrimCons(tmp15064, V4538)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4669, Z4668, tmp15065, V4539, V4540, W4543, V4542)
return


}, 0)

__e.TailApply(PrimFunc(symbind), W4669, tmp15059, V4539, V4540, W4543, tmp15060)
return


}, 0)

tmp15066 := Call(__e, PrimFunc(symbind), W4670, tmp15054, V4539, V4540, W4543, tmp15055)


__e.TailApply(PrimFunc(symshen_4gc), V4539, tmp15066)
return


}, 1)

tmp15067 := Call(__e, PrimFunc(symshen_4newpv), V4539)


tmp15068 := Call(__e, tmp15051, tmp15067)


__e.TailApply(PrimFunc(symshen_4gc), V4539, tmp15068)
return


}, 1)

tmp15069 := Call(__e, PrimFunc(symshen_4newpv), V4539)


__e.TailApply(tmp15050, tmp15069)
return


}, 1))
return
}, 1)

__e.TailApply(tmp14981, tmp15049)
return


}, 1)

tmp15070 := Call(__e, PrimFunc(symshen_4lazyderef), V4537, V4539)


__e.TailApply(tmp14980, tmp15070)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15073 := PrimTail(W4662)

tmp15074 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15073, V4539)


__e.TailApply(tmp14979, tmp15074)
return


}, 1)

tmp15075 := PrimHead(W4662)

__e.TailApply(tmp14978, tmp15075)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15078 := PrimTail(W4660)

tmp15079 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15078, V4539)


__e.TailApply(tmp14977, tmp15079)
return


}, 1)

tmp15080 := PrimHead(W4660)

__e.TailApply(tmp14976, tmp15080)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15083 := PrimTail(W4658)

tmp15084 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15083, V4539)


__e.TailApply(tmp14975, tmp15084)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15087 := PrimHead(W4658)

tmp15088 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15087, V4539)


__e.TailApply(tmp14974, tmp15088)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15091 := Call(__e, PrimFunc(symshen_4lazyderef), V4536, V4539)


tmp15092 := Call(__e, tmp14973, tmp15091)


ifres14972 = tmp15092


} else {
ifres14972 = False


}

__e.TailApply(tmp14672, ifres14972)
return


} else {
__e.Return(W4647)
return
}


}, 1)

tmp15136 := Call(__e, PrimFunc(symshen_4unlocked_2), V4540)


var ifres15096 Obj

if True == tmp15136 {
tmp15097 := MakeNative(func(__e *ControlFlow) {
W4648 := __e.Get(1)
_ = W4648
tmp15133 := PrimIsPair(W4648)

if True == tmp15133 {
tmp15098 := MakeNative(func(__e *ControlFlow) {
W4649 := __e.Get(1)
_ = W4649
tmp15129 := PrimEqual(W4649, sym_8s)

if True == tmp15129 {
tmp15099 := MakeNative(func(__e *ControlFlow) {
W4650 := __e.Get(1)
_ = W4650
tmp15125 := PrimIsPair(W4650)

if True == tmp15125 {
tmp15100 := MakeNative(func(__e *ControlFlow) {
W4651 := __e.Get(1)
_ = W4651
tmp15101 := MakeNative(func(__e *ControlFlow) {
W4652 := __e.Get(1)
_ = W4652
tmp15120 := PrimIsPair(W4652)

if True == tmp15120 {
tmp15102 := MakeNative(func(__e *ControlFlow) {
W4653 := __e.Get(1)
_ = W4653
tmp15103 := MakeNative(func(__e *ControlFlow) {
W4654 := __e.Get(1)
_ = W4654
tmp15115 := PrimEqual(W4654, Nil)

if True == tmp15115 {
tmp15104 := MakeNative(func(__e *ControlFlow) {
W4655 := __e.Get(1)
_ = W4655
tmp15105 := MakeNative(func(__e *ControlFlow) {
W4656 := __e.Get(1)
_ = W4656
tmp15109 := PrimEqual(W4655, symstring)

if True == tmp15109 {
__e.TailApply(PrimFunc(symthaw), W4656)
return
} else {
tmp15107 := Call(__e, PrimFunc(symshen_4pvar_2), W4655)


if True == tmp15107 {
__e.TailApply(PrimFunc(symshen_4bind_b), W4655, symstring, V4539, W4656)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp15110 := MakeNative(func(__e *ControlFlow) {
tmp15111 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15111

tmp15112 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4653, symstring, V4538, V4539, V4540, W4543, V4542)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4651, symstring, V4538, V4539, V4540, W4543, tmp15112)
return


}, 0)

__e.TailApply(tmp15105, tmp15110)
return


}, 1)

tmp15113 := Call(__e, PrimFunc(symshen_4lazyderef), V4537, V4539)


__e.TailApply(tmp15104, tmp15113)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15116 := PrimTail(W4652)

tmp15117 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15116, V4539)


__e.TailApply(tmp15103, tmp15117)
return


}, 1)

tmp15118 := PrimHead(W4652)

__e.TailApply(tmp15102, tmp15118)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15121 := PrimTail(W4650)

tmp15122 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15121, V4539)


__e.TailApply(tmp15101, tmp15122)
return


}, 1)

tmp15123 := PrimHead(W4650)

__e.TailApply(tmp15100, tmp15123)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15126 := PrimTail(W4648)

tmp15127 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15126, V4539)


__e.TailApply(tmp15099, tmp15127)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15130 := PrimHead(W4648)

tmp15131 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15130, V4539)


__e.TailApply(tmp15098, tmp15131)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15134 := Call(__e, PrimFunc(symshen_4lazyderef), V4536, V4539)


tmp15135 := Call(__e, tmp15097, tmp15134)


ifres15096 = tmp15135


} else {
ifres15096 = False


}

__e.TailApply(tmp14671, ifres15096)
return


} else {
__e.Return(W4626)
return
}


}, 1)

tmp15221 := Call(__e, PrimFunc(symshen_4unlocked_2), V4540)


var ifres15139 Obj

if True == tmp15221 {
tmp15140 := MakeNative(func(__e *ControlFlow) {
W4627 := __e.Get(1)
_ = W4627
tmp15218 := PrimIsPair(W4627)

if True == tmp15218 {
tmp15141 := MakeNative(func(__e *ControlFlow) {
W4628 := __e.Get(1)
_ = W4628
tmp15214 := PrimEqual(W4628, sym_8v)

if True == tmp15214 {
tmp15142 := MakeNative(func(__e *ControlFlow) {
W4629 := __e.Get(1)
_ = W4629
tmp15210 := PrimIsPair(W4629)

if True == tmp15210 {
tmp15143 := MakeNative(func(__e *ControlFlow) {
W4630 := __e.Get(1)
_ = W4630
tmp15144 := MakeNative(func(__e *ControlFlow) {
W4631 := __e.Get(1)
_ = W4631
tmp15205 := PrimIsPair(W4631)

if True == tmp15205 {
tmp15145 := MakeNative(func(__e *ControlFlow) {
W4632 := __e.Get(1)
_ = W4632
tmp15146 := MakeNative(func(__e *ControlFlow) {
W4633 := __e.Get(1)
_ = W4633
tmp15200 := PrimEqual(W4633, Nil)

if True == tmp15200 {
tmp15147 := MakeNative(func(__e *ControlFlow) {
W4634 := __e.Get(1)
_ = W4634
tmp15148 := MakeNative(func(__e *ControlFlow) {
W4635 := __e.Get(1)
_ = W4635
tmp15192 := PrimIsPair(W4634)

if True == tmp15192 {
tmp15149 := MakeNative(func(__e *ControlFlow) {
W4637 := __e.Get(1)
_ = W4637
tmp15150 := MakeNative(func(__e *ControlFlow) {
W4638 := __e.Get(1)
_ = W4638
tmp15154 := PrimEqual(W4637, symvector)

if True == tmp15154 {
__e.TailApply(PrimFunc(symthaw), W4638)
return
} else {
tmp15152 := Call(__e, PrimFunc(symshen_4pvar_2), W4637)


if True == tmp15152 {
__e.TailApply(PrimFunc(symshen_4bind_b), W4637, symvector, V4539, W4638)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp15155 := MakeNative(func(__e *ControlFlow) {
tmp15156 := MakeNative(func(__e *ControlFlow) {
W4639 := __e.Get(1)
_ = W4639
tmp15157 := MakeNative(func(__e *ControlFlow) {
W4640 := __e.Get(1)
_ = W4640
tmp15177 := PrimIsPair(W4639)

if True == tmp15177 {
tmp15158 := MakeNative(func(__e *ControlFlow) {
W4642 := __e.Get(1)
_ = W4642
tmp15159 := MakeNative(func(__e *ControlFlow) {
W4643 := __e.Get(1)
_ = W4643
tmp15160 := MakeNative(func(__e *ControlFlow) {
W4644 := __e.Get(1)
_ = W4644
tmp15164 := PrimEqual(W4643, Nil)

if True == tmp15164 {
__e.TailApply(PrimFunc(symthaw), W4644)
return
} else {
tmp15162 := Call(__e, PrimFunc(symshen_4pvar_2), W4643)


if True == tmp15162 {
__e.TailApply(PrimFunc(symshen_4bind_b), W4643, Nil, V4539, W4644)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp15165 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W4640, W4642)
return
}, 0)

__e.TailApply(tmp15160, tmp15165)
return


}, 1)

tmp15166 := PrimTail(W4639)

tmp15167 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15166, V4539)


__e.TailApply(tmp15159, tmp15167)
return


}, 1)

tmp15168 := PrimHead(W4639)

__e.TailApply(tmp15158, tmp15168)
return


} else {
tmp15175 := Call(__e, PrimFunc(symshen_4pvar_2), W4639)


if True == tmp15175 {
tmp15169 := MakeNative(func(__e *ControlFlow) {
W4645 := __e.Get(1)
_ = W4645
tmp15170 := PrimCons(W4645, Nil)

tmp15171 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W4640, W4645)
return
}, 0)

tmp15172 := Call(__e, PrimFunc(symshen_4bind_b), W4639, tmp15170, V4539, tmp15171)


__e.TailApply(PrimFunc(symshen_4gc), V4539, tmp15172)
return


}, 1)

tmp15173 := Call(__e, PrimFunc(symshen_4newpv), V4539)


__e.TailApply(tmp15169, tmp15173)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp15178 := MakeNative(func(__e *ControlFlow) {
Z4641 := __e.Get(1)
_ = Z4641
__e.TailApply(W4635, Z4641)
return
}, 1)

__e.TailApply(tmp15157, tmp15178)
return


}, 1)

tmp15179 := PrimTail(W4634)

tmp15180 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15179, V4539)


__e.TailApply(tmp15156, tmp15180)
return


}, 0)

__e.TailApply(tmp15150, tmp15155)
return


}, 1)

tmp15181 := PrimHead(W4634)

tmp15182 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15181, V4539)


__e.TailApply(tmp15149, tmp15182)
return


} else {
tmp15190 := Call(__e, PrimFunc(symshen_4pvar_2), W4634)


if True == tmp15190 {
tmp15183 := MakeNative(func(__e *ControlFlow) {
W4646 := __e.Get(1)
_ = W4646
tmp15184 := PrimCons(W4646, Nil)

tmp15185 := PrimCons(symvector, tmp15184)

tmp15186 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W4635, W4646)
return
}, 0)

tmp15187 := Call(__e, PrimFunc(symshen_4bind_b), W4634, tmp15185, V4539, tmp15186)


__e.TailApply(PrimFunc(symshen_4gc), V4539, tmp15187)
return


}, 1)

tmp15188 := Call(__e, PrimFunc(symshen_4newpv), V4539)


__e.TailApply(tmp15183, tmp15188)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp15193 := MakeNative(func(__e *ControlFlow) {
Z4636 := __e.Get(1)
_ = Z4636
tmp15194 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15194

tmp15195 := MakeNative(func(__e *ControlFlow) {
tmp15196 := PrimCons(Z4636, Nil)

tmp15197 := PrimCons(symvector, tmp15196)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4632, tmp15197, V4538, V4539, V4540, W4543, V4542)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4630, Z4636, V4538, V4539, V4540, W4543, tmp15195)
return


}, 1)

__e.TailApply(tmp15148, tmp15193)
return


}, 1)

tmp15198 := Call(__e, PrimFunc(symshen_4lazyderef), V4537, V4539)


__e.TailApply(tmp15147, tmp15198)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15201 := PrimTail(W4631)

tmp15202 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15201, V4539)


__e.TailApply(tmp15146, tmp15202)
return


}, 1)

tmp15203 := PrimHead(W4631)

__e.TailApply(tmp15145, tmp15203)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15206 := PrimTail(W4629)

tmp15207 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15206, V4539)


__e.TailApply(tmp15144, tmp15207)
return


}, 1)

tmp15208 := PrimHead(W4629)

__e.TailApply(tmp15143, tmp15208)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15211 := PrimTail(W4627)

tmp15212 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15211, V4539)


__e.TailApply(tmp15142, tmp15212)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15215 := PrimHead(W4627)

tmp15216 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15215, V4539)


__e.TailApply(tmp15141, tmp15216)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15219 := Call(__e, PrimFunc(symshen_4lazyderef), V4536, V4539)


tmp15220 := Call(__e, tmp15140, tmp15219)


ifres15139 = tmp15220


} else {
ifres15139 = False


}

__e.TailApply(tmp14670, ifres15139)
return


} else {
__e.Return(W4598)
return
}


}, 1)

tmp15327 := Call(__e, PrimFunc(symshen_4unlocked_2), V4540)


var ifres15224 Obj

if True == tmp15327 {
tmp15225 := MakeNative(func(__e *ControlFlow) {
W4599 := __e.Get(1)
_ = W4599
tmp15324 := PrimIsPair(W4599)

if True == tmp15324 {
tmp15226 := MakeNative(func(__e *ControlFlow) {
W4600 := __e.Get(1)
_ = W4600
tmp15320 := PrimEqual(W4600, sym_8p)

if True == tmp15320 {
tmp15227 := MakeNative(func(__e *ControlFlow) {
W4601 := __e.Get(1)
_ = W4601
tmp15316 := PrimIsPair(W4601)

if True == tmp15316 {
tmp15228 := MakeNative(func(__e *ControlFlow) {
W4602 := __e.Get(1)
_ = W4602
tmp15229 := MakeNative(func(__e *ControlFlow) {
W4603 := __e.Get(1)
_ = W4603
tmp15311 := PrimIsPair(W4603)

if True == tmp15311 {
tmp15230 := MakeNative(func(__e *ControlFlow) {
W4604 := __e.Get(1)
_ = W4604
tmp15231 := MakeNative(func(__e *ControlFlow) {
W4605 := __e.Get(1)
_ = W4605
tmp15306 := PrimEqual(W4605, Nil)

if True == tmp15306 {
tmp15232 := MakeNative(func(__e *ControlFlow) {
W4606 := __e.Get(1)
_ = W4606
tmp15233 := MakeNative(func(__e *ControlFlow) {
W4607 := __e.Get(1)
_ = W4607
tmp15300 := PrimIsPair(W4606)

if True == tmp15300 {
tmp15234 := MakeNative(func(__e *ControlFlow) {
W4610 := __e.Get(1)
_ = W4610
tmp15235 := MakeNative(func(__e *ControlFlow) {
W4611 := __e.Get(1)
_ = W4611
tmp15236 := MakeNative(func(__e *ControlFlow) {
W4612 := __e.Get(1)
_ = W4612
tmp15280 := PrimIsPair(W4611)

if True == tmp15280 {
tmp15237 := MakeNative(func(__e *ControlFlow) {
W4614 := __e.Get(1)
_ = W4614
tmp15238 := MakeNative(func(__e *ControlFlow) {
W4615 := __e.Get(1)
_ = W4615
tmp15242 := PrimEqual(W4614, sym_d)

if True == tmp15242 {
__e.TailApply(PrimFunc(symthaw), W4615)
return
} else {
tmp15240 := Call(__e, PrimFunc(symshen_4pvar_2), W4614)


if True == tmp15240 {
__e.TailApply(PrimFunc(symshen_4bind_b), W4614, sym_d, V4539, W4615)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp15243 := MakeNative(func(__e *ControlFlow) {
tmp15244 := MakeNative(func(__e *ControlFlow) {
W4616 := __e.Get(1)
_ = W4616
tmp15245 := MakeNative(func(__e *ControlFlow) {
W4617 := __e.Get(1)
_ = W4617
tmp15265 := PrimIsPair(W4616)

if True == tmp15265 {
tmp15246 := MakeNative(func(__e *ControlFlow) {
W4619 := __e.Get(1)
_ = W4619
tmp15247 := MakeNative(func(__e *ControlFlow) {
W4620 := __e.Get(1)
_ = W4620
tmp15248 := MakeNative(func(__e *ControlFlow) {
W4621 := __e.Get(1)
_ = W4621
tmp15252 := PrimEqual(W4620, Nil)

if True == tmp15252 {
__e.TailApply(PrimFunc(symthaw), W4621)
return
} else {
tmp15250 := Call(__e, PrimFunc(symshen_4pvar_2), W4620)


if True == tmp15250 {
__e.TailApply(PrimFunc(symshen_4bind_b), W4620, Nil, V4539, W4621)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp15253 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W4617, W4619)
return
}, 0)

__e.TailApply(tmp15248, tmp15253)
return


}, 1)

tmp15254 := PrimTail(W4616)

tmp15255 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15254, V4539)


__e.TailApply(tmp15247, tmp15255)
return


}, 1)

tmp15256 := PrimHead(W4616)

__e.TailApply(tmp15246, tmp15256)
return


} else {
tmp15263 := Call(__e, PrimFunc(symshen_4pvar_2), W4616)


if True == tmp15263 {
tmp15257 := MakeNative(func(__e *ControlFlow) {
W4622 := __e.Get(1)
_ = W4622
tmp15258 := PrimCons(W4622, Nil)

tmp15259 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W4617, W4622)
return
}, 0)

tmp15260 := Call(__e, PrimFunc(symshen_4bind_b), W4616, tmp15258, V4539, tmp15259)


__e.TailApply(PrimFunc(symshen_4gc), V4539, tmp15260)
return


}, 1)

tmp15261 := Call(__e, PrimFunc(symshen_4newpv), V4539)


__e.TailApply(tmp15257, tmp15261)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp15266 := MakeNative(func(__e *ControlFlow) {
Z4618 := __e.Get(1)
_ = Z4618
__e.TailApply(W4612, Z4618)
return
}, 1)

__e.TailApply(tmp15245, tmp15266)
return


}, 1)

tmp15267 := PrimTail(W4611)

tmp15268 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15267, V4539)


__e.TailApply(tmp15244, tmp15268)
return


}, 0)

__e.TailApply(tmp15238, tmp15243)
return


}, 1)

tmp15269 := PrimHead(W4611)

tmp15270 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15269, V4539)


__e.TailApply(tmp15237, tmp15270)
return


} else {
tmp15278 := Call(__e, PrimFunc(symshen_4pvar_2), W4611)


if True == tmp15278 {
tmp15271 := MakeNative(func(__e *ControlFlow) {
W4623 := __e.Get(1)
_ = W4623
tmp15272 := PrimCons(W4623, Nil)

tmp15273 := PrimCons(sym_d, tmp15272)

tmp15274 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W4612, W4623)
return
}, 0)

tmp15275 := Call(__e, PrimFunc(symshen_4bind_b), W4611, tmp15273, V4539, tmp15274)


__e.TailApply(PrimFunc(symshen_4gc), V4539, tmp15275)
return


}, 1)

tmp15276 := Call(__e, PrimFunc(symshen_4newpv), V4539)


__e.TailApply(tmp15271, tmp15276)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp15281 := MakeNative(func(__e *ControlFlow) {
Z4613 := __e.Get(1)
_ = Z4613
tmp15282 := Call(__e, W4607, W4610)


__e.TailApply(tmp15282, Z4613)
return


}, 1)

__e.TailApply(tmp15236, tmp15281)
return


}, 1)

tmp15283 := PrimTail(W4606)

tmp15284 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15283, V4539)


__e.TailApply(tmp15235, tmp15284)
return


}, 1)

tmp15285 := PrimHead(W4606)

__e.TailApply(tmp15234, tmp15285)
return


} else {
tmp15298 := Call(__e, PrimFunc(symshen_4pvar_2), W4606)


if True == tmp15298 {
tmp15286 := MakeNative(func(__e *ControlFlow) {
W4624 := __e.Get(1)
_ = W4624
tmp15287 := MakeNative(func(__e *ControlFlow) {
W4625 := __e.Get(1)
_ = W4625
tmp15288 := PrimCons(W4625, Nil)

tmp15289 := PrimCons(sym_d, tmp15288)

tmp15290 := PrimCons(W4624, tmp15289)

tmp15291 := MakeNative(func(__e *ControlFlow) {
tmp15292 := Call(__e, W4607, W4624)


__e.TailApply(tmp15292, W4625)
return


}, 0)

tmp15293 := Call(__e, PrimFunc(symshen_4bind_b), W4606, tmp15290, V4539, tmp15291)


__e.TailApply(PrimFunc(symshen_4gc), V4539, tmp15293)
return


}, 1)

tmp15294 := Call(__e, PrimFunc(symshen_4newpv), V4539)


tmp15295 := Call(__e, tmp15287, tmp15294)


__e.TailApply(PrimFunc(symshen_4gc), V4539, tmp15295)
return


}, 1)

tmp15296 := Call(__e, PrimFunc(symshen_4newpv), V4539)


__e.TailApply(tmp15286, tmp15296)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp15301 := MakeNative(func(__e *ControlFlow) {
Z4608 := __e.Get(1)
_ = Z4608
__e.Return(MakeNative(func(__e *ControlFlow) {
Z4609 := __e.Get(1)
_ = Z4609
tmp15302 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15302

tmp15303 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4604, Z4609, V4538, V4539, V4540, W4543, V4542)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4602, Z4608, V4538, V4539, V4540, W4543, tmp15303)
return


}, 1))
return
}, 1)

__e.TailApply(tmp15233, tmp15301)
return


}, 1)

tmp15304 := Call(__e, PrimFunc(symshen_4lazyderef), V4537, V4539)


__e.TailApply(tmp15232, tmp15304)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15307 := PrimTail(W4603)

tmp15308 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15307, V4539)


__e.TailApply(tmp15231, tmp15308)
return


}, 1)

tmp15309 := PrimHead(W4603)

__e.TailApply(tmp15230, tmp15309)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15312 := PrimTail(W4601)

tmp15313 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15312, V4539)


__e.TailApply(tmp15229, tmp15313)
return


}, 1)

tmp15314 := PrimHead(W4601)

__e.TailApply(tmp15228, tmp15314)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15317 := PrimTail(W4599)

tmp15318 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15317, V4539)


__e.TailApply(tmp15227, tmp15318)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15321 := PrimHead(W4599)

tmp15322 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15321, V4539)


__e.TailApply(tmp15226, tmp15322)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15325 := Call(__e, PrimFunc(symshen_4lazyderef), V4536, V4539)


tmp15326 := Call(__e, tmp15225, tmp15325)


ifres15224 = tmp15326


} else {
ifres15224 = False


}

__e.TailApply(tmp14669, ifres15224)
return


} else {
__e.Return(W4577)
return
}


}, 1)

tmp15412 := Call(__e, PrimFunc(symshen_4unlocked_2), V4540)


var ifres15330 Obj

if True == tmp15412 {
tmp15331 := MakeNative(func(__e *ControlFlow) {
W4578 := __e.Get(1)
_ = W4578
tmp15409 := PrimIsPair(W4578)

if True == tmp15409 {
tmp15332 := MakeNative(func(__e *ControlFlow) {
W4579 := __e.Get(1)
_ = W4579
tmp15405 := PrimEqual(W4579, symcons)

if True == tmp15405 {
tmp15333 := MakeNative(func(__e *ControlFlow) {
W4580 := __e.Get(1)
_ = W4580
tmp15401 := PrimIsPair(W4580)

if True == tmp15401 {
tmp15334 := MakeNative(func(__e *ControlFlow) {
W4581 := __e.Get(1)
_ = W4581
tmp15335 := MakeNative(func(__e *ControlFlow) {
W4582 := __e.Get(1)
_ = W4582
tmp15396 := PrimIsPair(W4582)

if True == tmp15396 {
tmp15336 := MakeNative(func(__e *ControlFlow) {
W4583 := __e.Get(1)
_ = W4583
tmp15337 := MakeNative(func(__e *ControlFlow) {
W4584 := __e.Get(1)
_ = W4584
tmp15391 := PrimEqual(W4584, Nil)

if True == tmp15391 {
tmp15338 := MakeNative(func(__e *ControlFlow) {
W4585 := __e.Get(1)
_ = W4585
tmp15339 := MakeNative(func(__e *ControlFlow) {
W4586 := __e.Get(1)
_ = W4586
tmp15383 := PrimIsPair(W4585)

if True == tmp15383 {
tmp15340 := MakeNative(func(__e *ControlFlow) {
W4588 := __e.Get(1)
_ = W4588
tmp15341 := MakeNative(func(__e *ControlFlow) {
W4589 := __e.Get(1)
_ = W4589
tmp15345 := PrimEqual(W4588, symlist)

if True == tmp15345 {
__e.TailApply(PrimFunc(symthaw), W4589)
return
} else {
tmp15343 := Call(__e, PrimFunc(symshen_4pvar_2), W4588)


if True == tmp15343 {
__e.TailApply(PrimFunc(symshen_4bind_b), W4588, symlist, V4539, W4589)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp15346 := MakeNative(func(__e *ControlFlow) {
tmp15347 := MakeNative(func(__e *ControlFlow) {
W4590 := __e.Get(1)
_ = W4590
tmp15348 := MakeNative(func(__e *ControlFlow) {
W4591 := __e.Get(1)
_ = W4591
tmp15368 := PrimIsPair(W4590)

if True == tmp15368 {
tmp15349 := MakeNative(func(__e *ControlFlow) {
W4593 := __e.Get(1)
_ = W4593
tmp15350 := MakeNative(func(__e *ControlFlow) {
W4594 := __e.Get(1)
_ = W4594
tmp15351 := MakeNative(func(__e *ControlFlow) {
W4595 := __e.Get(1)
_ = W4595
tmp15355 := PrimEqual(W4594, Nil)

if True == tmp15355 {
__e.TailApply(PrimFunc(symthaw), W4595)
return
} else {
tmp15353 := Call(__e, PrimFunc(symshen_4pvar_2), W4594)


if True == tmp15353 {
__e.TailApply(PrimFunc(symshen_4bind_b), W4594, Nil, V4539, W4595)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp15356 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W4591, W4593)
return
}, 0)

__e.TailApply(tmp15351, tmp15356)
return


}, 1)

tmp15357 := PrimTail(W4590)

tmp15358 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15357, V4539)


__e.TailApply(tmp15350, tmp15358)
return


}, 1)

tmp15359 := PrimHead(W4590)

__e.TailApply(tmp15349, tmp15359)
return


} else {
tmp15366 := Call(__e, PrimFunc(symshen_4pvar_2), W4590)


if True == tmp15366 {
tmp15360 := MakeNative(func(__e *ControlFlow) {
W4596 := __e.Get(1)
_ = W4596
tmp15361 := PrimCons(W4596, Nil)

tmp15362 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W4591, W4596)
return
}, 0)

tmp15363 := Call(__e, PrimFunc(symshen_4bind_b), W4590, tmp15361, V4539, tmp15362)


__e.TailApply(PrimFunc(symshen_4gc), V4539, tmp15363)
return


}, 1)

tmp15364 := Call(__e, PrimFunc(symshen_4newpv), V4539)


__e.TailApply(tmp15360, tmp15364)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp15369 := MakeNative(func(__e *ControlFlow) {
Z4592 := __e.Get(1)
_ = Z4592
__e.TailApply(W4586, Z4592)
return
}, 1)

__e.TailApply(tmp15348, tmp15369)
return


}, 1)

tmp15370 := PrimTail(W4585)

tmp15371 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15370, V4539)


__e.TailApply(tmp15347, tmp15371)
return


}, 0)

__e.TailApply(tmp15341, tmp15346)
return


}, 1)

tmp15372 := PrimHead(W4585)

tmp15373 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15372, V4539)


__e.TailApply(tmp15340, tmp15373)
return


} else {
tmp15381 := Call(__e, PrimFunc(symshen_4pvar_2), W4585)


if True == tmp15381 {
tmp15374 := MakeNative(func(__e *ControlFlow) {
W4597 := __e.Get(1)
_ = W4597
tmp15375 := PrimCons(W4597, Nil)

tmp15376 := PrimCons(symlist, tmp15375)

tmp15377 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W4586, W4597)
return
}, 0)

tmp15378 := Call(__e, PrimFunc(symshen_4bind_b), W4585, tmp15376, V4539, tmp15377)


__e.TailApply(PrimFunc(symshen_4gc), V4539, tmp15378)
return


}, 1)

tmp15379 := Call(__e, PrimFunc(symshen_4newpv), V4539)


__e.TailApply(tmp15374, tmp15379)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp15384 := MakeNative(func(__e *ControlFlow) {
Z4587 := __e.Get(1)
_ = Z4587
tmp15385 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15385

tmp15386 := MakeNative(func(__e *ControlFlow) {
tmp15387 := PrimCons(Z4587, Nil)

tmp15388 := PrimCons(symlist, tmp15387)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4583, tmp15388, V4538, V4539, V4540, W4543, V4542)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4581, Z4587, V4538, V4539, V4540, W4543, tmp15386)
return


}, 1)

__e.TailApply(tmp15339, tmp15384)
return


}, 1)

tmp15389 := Call(__e, PrimFunc(symshen_4lazyderef), V4537, V4539)


__e.TailApply(tmp15338, tmp15389)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15392 := PrimTail(W4582)

tmp15393 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15392, V4539)


__e.TailApply(tmp15337, tmp15393)
return


}, 1)

tmp15394 := PrimHead(W4582)

__e.TailApply(tmp15336, tmp15394)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15397 := PrimTail(W4580)

tmp15398 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15397, V4539)


__e.TailApply(tmp15335, tmp15398)
return


}, 1)

tmp15399 := PrimHead(W4580)

__e.TailApply(tmp15334, tmp15399)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15402 := PrimTail(W4578)

tmp15403 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15402, V4539)


__e.TailApply(tmp15333, tmp15403)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15406 := PrimHead(W4578)

tmp15407 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15406, V4539)


__e.TailApply(tmp15332, tmp15407)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15410 := Call(__e, PrimFunc(symshen_4lazyderef), V4536, V4539)


tmp15411 := Call(__e, tmp15331, tmp15410)


ifres15330 = tmp15411


} else {
ifres15330 = False


}

__e.TailApply(tmp14668, ifres15330)
return


} else {
__e.Return(W4570)
return
}


}, 1)

tmp15443 := Call(__e, PrimFunc(symshen_4unlocked_2), V4540)


var ifres15415 Obj

if True == tmp15443 {
tmp15416 := MakeNative(func(__e *ControlFlow) {
W4571 := __e.Get(1)
_ = W4571
tmp15440 := PrimIsPair(W4571)

if True == tmp15440 {
tmp15417 := MakeNative(func(__e *ControlFlow) {
W4572 := __e.Get(1)
_ = W4572
tmp15418 := MakeNative(func(__e *ControlFlow) {
W4573 := __e.Get(1)
_ = W4573
tmp15435 := PrimIsPair(W4573)

if True == tmp15435 {
tmp15419 := MakeNative(func(__e *ControlFlow) {
W4574 := __e.Get(1)
_ = W4574
tmp15420 := MakeNative(func(__e *ControlFlow) {
W4575 := __e.Get(1)
_ = W4575
tmp15430 := PrimEqual(W4575, Nil)

if True == tmp15430 {
tmp15421 := MakeNative(func(__e *ControlFlow) {
W4576 := __e.Get(1)
_ = W4576
tmp15422 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15422

tmp15423 := PrimCons(V4537, Nil)

tmp15424 := PrimCons(sym_1_1_6, tmp15423)

tmp15425 := PrimCons(W4576, tmp15424)

tmp15426 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4574, W4576, V4538, V4539, V4540, W4543, V4542)
return
}, 0)

tmp15427 := Call(__e, PrimFunc(symshen_4system_1S_1h), W4572, tmp15425, V4538, V4539, V4540, W4543, tmp15426)


__e.TailApply(PrimFunc(symshen_4gc), V4539, tmp15427)
return


}, 1)

tmp15428 := Call(__e, PrimFunc(symshen_4newpv), V4539)


__e.TailApply(tmp15421, tmp15428)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15431 := PrimTail(W4573)

tmp15432 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15431, V4539)


__e.TailApply(tmp15420, tmp15432)
return


}, 1)

tmp15433 := PrimHead(W4573)

__e.TailApply(tmp15419, tmp15433)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15436 := PrimTail(W4571)

tmp15437 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15436, V4539)


__e.TailApply(tmp15418, tmp15437)
return


}, 1)

tmp15438 := PrimHead(W4571)

__e.TailApply(tmp15417, tmp15438)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15441 := Call(__e, PrimFunc(symshen_4lazyderef), V4536, V4539)


tmp15442 := Call(__e, tmp15416, tmp15441)


ifres15415 = tmp15442


} else {
ifres15415 = False


}

__e.TailApply(tmp14667, ifres15415)
return


} else {
__e.Return(W4563)
return
}


}, 1)

tmp15478 := Call(__e, PrimFunc(symshen_4unlocked_2), V4540)


var ifres15446 Obj

if True == tmp15478 {
tmp15447 := MakeNative(func(__e *ControlFlow) {
W4564 := __e.Get(1)
_ = W4564
tmp15475 := PrimIsPair(W4564)

if True == tmp15475 {
tmp15448 := MakeNative(func(__e *ControlFlow) {
W4565 := __e.Get(1)
_ = W4565
tmp15449 := MakeNative(func(__e *ControlFlow) {
W4566 := __e.Get(1)
_ = W4566
tmp15470 := PrimIsPair(W4566)

if True == tmp15470 {
tmp15450 := MakeNative(func(__e *ControlFlow) {
W4567 := __e.Get(1)
_ = W4567
tmp15451 := MakeNative(func(__e *ControlFlow) {
W4568 := __e.Get(1)
_ = W4568
tmp15465 := PrimEqual(W4568, Nil)

if True == tmp15465 {
tmp15452 := MakeNative(func(__e *ControlFlow) {
W4569 := __e.Get(1)
_ = W4569
tmp15453 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15453

tmp15454 := Call(__e, PrimFunc(symshen_4lazyderef), W4565, V4539)


tmp15455 := PrimIsPair(tmp15454)

tmp15456 := PrimNot(tmp15455)

tmp15457 := MakeNative(func(__e *ControlFlow) {
tmp15458 := PrimCons(V4537, Nil)

tmp15459 := PrimCons(sym_1_1_6, tmp15458)

tmp15460 := PrimCons(W4569, tmp15459)

tmp15461 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4567, W4569, V4538, V4539, V4540, W4543, V4542)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4lookupsig), W4565, tmp15460, V4539, V4540, W4543, tmp15461)
return


}, 0)

tmp15462 := Call(__e, PrimFunc(symwhen), tmp15456, V4539, V4540, W4543, tmp15457)


__e.TailApply(PrimFunc(symshen_4gc), V4539, tmp15462)
return


}, 1)

tmp15463 := Call(__e, PrimFunc(symshen_4newpv), V4539)


__e.TailApply(tmp15452, tmp15463)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15466 := PrimTail(W4566)

tmp15467 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15466, V4539)


__e.TailApply(tmp15451, tmp15467)
return


}, 1)

tmp15468 := PrimHead(W4566)

__e.TailApply(tmp15450, tmp15468)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15471 := PrimTail(W4564)

tmp15472 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15471, V4539)


__e.TailApply(tmp15449, tmp15472)
return


}, 1)

tmp15473 := PrimHead(W4564)

__e.TailApply(tmp15448, tmp15473)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15476 := Call(__e, PrimFunc(symshen_4lazyderef), V4536, V4539)


tmp15477 := Call(__e, tmp15447, tmp15476)


ifres15446 = tmp15477


} else {
ifres15446 = False


}

__e.TailApply(tmp14666, ifres15446)
return


} else {
__e.Return(W4557)
return
}


}, 1)

tmp15505 := Call(__e, PrimFunc(symshen_4unlocked_2), V4540)


var ifres15481 Obj

if True == tmp15505 {
tmp15482 := MakeNative(func(__e *ControlFlow) {
W4558 := __e.Get(1)
_ = W4558
tmp15502 := PrimIsPair(W4558)

if True == tmp15502 {
tmp15483 := MakeNative(func(__e *ControlFlow) {
W4559 := __e.Get(1)
_ = W4559
tmp15498 := PrimEqual(W4559, symfn)

if True == tmp15498 {
tmp15484 := MakeNative(func(__e *ControlFlow) {
W4560 := __e.Get(1)
_ = W4560
tmp15494 := PrimIsPair(W4560)

if True == tmp15494 {
tmp15485 := MakeNative(func(__e *ControlFlow) {
W4561 := __e.Get(1)
_ = W4561
tmp15486 := MakeNative(func(__e *ControlFlow) {
W4562 := __e.Get(1)
_ = W4562
tmp15489 := PrimEqual(W4562, Nil)

if True == tmp15489 {
tmp15487 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15487

__e.TailApply(PrimFunc(symshen_4lookupsig), W4561, V4537, V4539, V4540, W4543, V4542)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15490 := PrimTail(W4560)

tmp15491 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15490, V4539)


__e.TailApply(tmp15486, tmp15491)
return


}, 1)

tmp15492 := PrimHead(W4560)

__e.TailApply(tmp15485, tmp15492)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15495 := PrimTail(W4558)

tmp15496 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15495, V4539)


__e.TailApply(tmp15484, tmp15496)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15499 := PrimHead(W4558)

tmp15500 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15499, V4539)


__e.TailApply(tmp15483, tmp15500)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15503 := Call(__e, PrimFunc(symshen_4lazyderef), V4536, V4539)


tmp15504 := Call(__e, tmp15482, tmp15503)


ifres15481 = tmp15504


} else {
ifres15481 = False


}

__e.TailApply(tmp14665, ifres15481)
return


} else {
__e.Return(W4551)
return
}


}, 1)

tmp15538 := Call(__e, PrimFunc(symshen_4unlocked_2), V4540)


var ifres15508 Obj

if True == tmp15538 {
tmp15509 := MakeNative(func(__e *ControlFlow) {
W4552 := __e.Get(1)
_ = W4552
tmp15535 := PrimIsPair(W4552)

if True == tmp15535 {
tmp15510 := MakeNative(func(__e *ControlFlow) {
W4553 := __e.Get(1)
_ = W4553
tmp15531 := PrimEqual(W4553, symfn)

if True == tmp15531 {
tmp15511 := MakeNative(func(__e *ControlFlow) {
W4554 := __e.Get(1)
_ = W4554
tmp15527 := PrimIsPair(W4554)

if True == tmp15527 {
tmp15512 := MakeNative(func(__e *ControlFlow) {
W4555 := __e.Get(1)
_ = W4555
tmp15513 := MakeNative(func(__e *ControlFlow) {
W4556 := __e.Get(1)
_ = W4556
tmp15522 := PrimEqual(W4556, Nil)

if True == tmp15522 {
tmp15514 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15514

tmp15515 := Call(__e, PrimFunc(symshen_4deref), W4555, V4539)


tmp15516 := Call(__e, PrimFunc(symarity), tmp15515)


tmp15517 := PrimEqual(tmp15516, MakeNumber(0))

tmp15518 := MakeNative(func(__e *ControlFlow) {
tmp15519 := MakeNative(func(__e *ControlFlow) {
tmp15520 := PrimCons(W4555, Nil)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), tmp15520, V4537, V4538, V4539, V4540, W4543, V4542)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V4539, V4540, W4543, tmp15519)
return


}, 0)

__e.TailApply(PrimFunc(symwhen), tmp15517, V4539, V4540, W4543, tmp15518)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15523 := PrimTail(W4554)

tmp15524 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15523, V4539)


__e.TailApply(tmp15513, tmp15524)
return


}, 1)

tmp15525 := PrimHead(W4554)

__e.TailApply(tmp15512, tmp15525)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15528 := PrimTail(W4552)

tmp15529 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15528, V4539)


__e.TailApply(tmp15511, tmp15529)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15532 := PrimHead(W4552)

tmp15533 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15532, V4539)


__e.TailApply(tmp15510, tmp15533)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15536 := Call(__e, PrimFunc(symshen_4lazyderef), V4536, V4539)


tmp15537 := Call(__e, tmp15509, tmp15536)


ifres15508 = tmp15537


} else {
ifres15508 = False


}

__e.TailApply(tmp14664, ifres15508)
return


} else {
__e.Return(W4547)
return
}


}, 1)

tmp15557 := Call(__e, PrimFunc(symshen_4unlocked_2), V4540)


var ifres15541 Obj

if True == tmp15557 {
tmp15542 := MakeNative(func(__e *ControlFlow) {
W4548 := __e.Get(1)
_ = W4548
tmp15554 := PrimIsPair(W4548)

if True == tmp15554 {
tmp15543 := MakeNative(func(__e *ControlFlow) {
W4549 := __e.Get(1)
_ = W4549
tmp15544 := MakeNative(func(__e *ControlFlow) {
W4550 := __e.Get(1)
_ = W4550
tmp15549 := PrimEqual(W4550, Nil)

if True == tmp15549 {
tmp15545 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15545

tmp15546 := PrimCons(V4537, Nil)

tmp15547 := PrimCons(sym_1_1_6, tmp15546)

__e.TailApply(PrimFunc(symshen_4lookupsig), W4549, tmp15547, V4539, V4540, W4543, V4542)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15550 := PrimTail(W4548)

tmp15551 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15550, V4539)


__e.TailApply(tmp15544, tmp15551)
return


}, 1)

tmp15552 := PrimHead(W4548)

__e.TailApply(tmp15543, tmp15552)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15555 := Call(__e, PrimFunc(symshen_4lazyderef), V4536, V4539)


tmp15556 := Call(__e, tmp15542, tmp15555)


ifres15541 = tmp15556


} else {
ifres15541 = False


}

__e.TailApply(tmp14663, ifres15541)
return


} else {
__e.Return(W4546)
return
}


}, 1)

tmp15563 := Call(__e, PrimFunc(symshen_4unlocked_2), V4540)


var ifres15560 Obj

if True == tmp15563 {
tmp15561 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15561

tmp15562 := Call(__e, PrimFunc(symshen_4by_1hypothesis), V4536, V4537, V4538, V4539, V4540, W4543, V4542)


ifres15560 = tmp15562


} else {
ifres15560 = False


}

__e.TailApply(tmp14662, ifres15560)
return


} else {
__e.Return(W4545)
return
}


}, 1)

tmp15573 := Call(__e, PrimFunc(symshen_4unlocked_2), V4540)


var ifres15566 Obj

if True == tmp15573 {
tmp15567 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15567

tmp15568 := Call(__e, PrimFunc(symshen_4lazyderef), V4536, V4539)


tmp15569 := PrimIsPair(tmp15568)

tmp15570 := PrimNot(tmp15569)

tmp15571 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4primitive), V4536, V4537, V4539, V4540, W4543, V4542)
return
}, 0)

tmp15572 := Call(__e, PrimFunc(symwhen), tmp15570, V4539, V4540, W4543, tmp15571)


ifres15566 = tmp15572


} else {
ifres15566 = False


}

__e.TailApply(tmp14661, ifres15566)
return


} else {
__e.Return(W4544)
return
}


}, 1)

tmp15585 := Call(__e, PrimFunc(symshen_4unlocked_2), V4540)


var ifres15576 Obj

if True == tmp15585 {
tmp15577 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15577

tmp15578 := PrimValue(symshen_4_dspy_d)

tmp15579 := MakeNative(func(__e *ControlFlow) {
tmp15580 := PrimIntern(MakeString(":"))

tmp15581 := PrimCons(V4537, Nil)

tmp15582 := PrimCons(tmp15580, tmp15581)

tmp15583 := PrimCons(V4536, tmp15582)

__e.TailApply(PrimFunc(symshen_4show), tmp15583, V4538, V4539, V4540, W4543, V4542)
return


}, 0)

tmp15584 := Call(__e, PrimFunc(symwhen), tmp15578, V4539, V4540, W4543, tmp15579)


ifres15576 = tmp15584


} else {
ifres15576 = False


}

__e.TailApply(tmp14660, ifres15576)
return


}, 1)

tmp15586 := PrimNumberAdd(V4541, MakeNumber(1))

__e.TailApply(tmp14659, tmp15586)
return


}, 7)

tmp15587 := Call(__e, ns2_1set, symshen_4system_1S_1h, tmp14658)


_ = tmp15587

tmp15588 := MakeNative(func(__e *ControlFlow) {
V4748 := __e.Get(1)
_ = V4748
tmp15622 := PrimIsPair(V4748)

var ifres15603 Obj

if True == tmp15622 {
tmp15620 := PrimHead(V4748)

tmp15621 := PrimEqual(symcons, tmp15620)

var ifres15605 Obj

if True == tmp15621 {
tmp15618 := PrimTail(V4748)

tmp15619 := PrimIsPair(tmp15618)

var ifres15607 Obj

if True == tmp15619 {
tmp15615 := PrimTail(V4748)

tmp15616 := PrimTail(tmp15615)

tmp15617 := PrimIsPair(tmp15616)

var ifres15609 Obj

if True == tmp15617 {
tmp15611 := PrimTail(V4748)

tmp15612 := PrimTail(tmp15611)

tmp15613 := PrimTail(tmp15612)

tmp15614 := PrimEqual(Nil, tmp15613)

var ifres15610 Obj

if True == tmp15614 {
ifres15610 = True


} else {
ifres15610 = False


}

ifres15609 = ifres15610


} else {
ifres15609 = False


}

var ifres15608 Obj

if True == ifres15609 {
ifres15608 = True


} else {
ifres15608 = False


}

ifres15607 = ifres15608


} else {
ifres15607 = False


}

var ifres15606 Obj

if True == ifres15607 {
ifres15606 = True


} else {
ifres15606 = False


}

ifres15605 = ifres15606


} else {
ifres15605 = False


}

var ifres15604 Obj

if True == ifres15605 {
ifres15604 = True


} else {
ifres15604 = False


}

ifres15603 = ifres15604


} else {
ifres15603 = False


}

if True == ifres15603 {
tmp15589 := PrimTail(V4748)

tmp15590 := PrimHead(tmp15589)

tmp15591 := Call(__e, PrimFunc(symshen_4rdecons), tmp15590)


tmp15592 := PrimTail(V4748)

tmp15593 := PrimTail(tmp15592)

tmp15594 := PrimHead(tmp15593)

tmp15595 := Call(__e, PrimFunc(symshen_4rdecons), tmp15594)


__e.Return(PrimCons(tmp15591, tmp15595))
return


} else {
tmp15601 := PrimIsPair(V4748)

if True == tmp15601 {
tmp15596 := PrimHead(V4748)

tmp15597 := Call(__e, PrimFunc(symshen_4rdecons), tmp15596)


tmp15598 := PrimTail(V4748)

tmp15599 := Call(__e, PrimFunc(symshen_4rdecons), tmp15598)


__e.Return(PrimCons(tmp15597, tmp15599))
return


} else {
__e.Return(V4748)
return
}


}


}, 1)

tmp15623 := Call(__e, ns2_1set, symshen_4rdecons, tmp15588)


_ = tmp15623

tmp15624 := MakeNative(func(__e *ControlFlow) {
V4749 := __e.Get(1)
_ = V4749
V4750 := __e.Get(2)
_ = V4750
V4751 := __e.Get(3)
_ = V4751
V4752 := __e.Get(4)
_ = V4752
V4753 := __e.Get(5)
_ = V4753
V4754 := __e.Get(6)
_ = V4754
tmp15625 := MakeNative(func(__e *ControlFlow) {
W4755 := __e.Get(1)
_ = W4755
tmp15733 := PrimEqual(W4755, False)

if True == tmp15733 {
tmp15626 := MakeNative(func(__e *ControlFlow) {
W4758 := __e.Get(1)
_ = W4758
tmp15717 := PrimEqual(W4758, False)

if True == tmp15717 {
tmp15627 := MakeNative(func(__e *ControlFlow) {
W4761 := __e.Get(1)
_ = W4761
tmp15701 := PrimEqual(W4761, False)

if True == tmp15701 {
tmp15628 := MakeNative(func(__e *ControlFlow) {
W4764 := __e.Get(1)
_ = W4764
tmp15685 := PrimEqual(W4764, False)

if True == tmp15685 {
tmp15683 := Call(__e, PrimFunc(symshen_4unlocked_2), V4752)


if True == tmp15683 {
tmp15629 := MakeNative(func(__e *ControlFlow) {
W4767 := __e.Get(1)
_ = W4767
tmp15680 := PrimEqual(W4767, Nil)

if True == tmp15680 {
tmp15630 := MakeNative(func(__e *ControlFlow) {
W4768 := __e.Get(1)
_ = W4768
tmp15631 := MakeNative(func(__e *ControlFlow) {
W4769 := __e.Get(1)
_ = W4769
tmp15675 := PrimIsPair(W4768)

if True == tmp15675 {
tmp15632 := MakeNative(func(__e *ControlFlow) {
W4771 := __e.Get(1)
_ = W4771
tmp15633 := MakeNative(func(__e *ControlFlow) {
W4772 := __e.Get(1)
_ = W4772
tmp15637 := PrimEqual(W4771, symlist)

if True == tmp15637 {
__e.TailApply(PrimFunc(symthaw), W4772)
return
} else {
tmp15635 := Call(__e, PrimFunc(symshen_4pvar_2), W4771)


if True == tmp15635 {
__e.TailApply(PrimFunc(symshen_4bind_b), W4771, symlist, V4751, W4772)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp15638 := MakeNative(func(__e *ControlFlow) {
tmp15639 := MakeNative(func(__e *ControlFlow) {
W4773 := __e.Get(1)
_ = W4773
tmp15640 := MakeNative(func(__e *ControlFlow) {
W4774 := __e.Get(1)
_ = W4774
tmp15660 := PrimIsPair(W4773)

if True == tmp15660 {
tmp15641 := MakeNative(func(__e *ControlFlow) {
W4776 := __e.Get(1)
_ = W4776
tmp15642 := MakeNative(func(__e *ControlFlow) {
W4777 := __e.Get(1)
_ = W4777
tmp15643 := MakeNative(func(__e *ControlFlow) {
W4778 := __e.Get(1)
_ = W4778
tmp15647 := PrimEqual(W4777, Nil)

if True == tmp15647 {
__e.TailApply(PrimFunc(symthaw), W4778)
return
} else {
tmp15645 := Call(__e, PrimFunc(symshen_4pvar_2), W4777)


if True == tmp15645 {
__e.TailApply(PrimFunc(symshen_4bind_b), W4777, Nil, V4751, W4778)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp15648 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W4774, W4776)
return
}, 0)

__e.TailApply(tmp15643, tmp15648)
return


}, 1)

tmp15649 := PrimTail(W4773)

tmp15650 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15649, V4751)


__e.TailApply(tmp15642, tmp15650)
return


}, 1)

tmp15651 := PrimHead(W4773)

__e.TailApply(tmp15641, tmp15651)
return


} else {
tmp15658 := Call(__e, PrimFunc(symshen_4pvar_2), W4773)


if True == tmp15658 {
tmp15652 := MakeNative(func(__e *ControlFlow) {
W4779 := __e.Get(1)
_ = W4779
tmp15653 := PrimCons(W4779, Nil)

tmp15654 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W4774, W4779)
return
}, 0)

tmp15655 := Call(__e, PrimFunc(symshen_4bind_b), W4773, tmp15653, V4751, tmp15654)


__e.TailApply(PrimFunc(symshen_4gc), V4751, tmp15655)
return


}, 1)

tmp15656 := Call(__e, PrimFunc(symshen_4newpv), V4751)


__e.TailApply(tmp15652, tmp15656)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp15661 := MakeNative(func(__e *ControlFlow) {
Z4775 := __e.Get(1)
_ = Z4775
__e.TailApply(W4769, Z4775)
return
}, 1)

__e.TailApply(tmp15640, tmp15661)
return


}, 1)

tmp15662 := PrimTail(W4768)

tmp15663 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15662, V4751)


__e.TailApply(tmp15639, tmp15663)
return


}, 0)

__e.TailApply(tmp15633, tmp15638)
return


}, 1)

tmp15664 := PrimHead(W4768)

tmp15665 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15664, V4751)


__e.TailApply(tmp15632, tmp15665)
return


} else {
tmp15673 := Call(__e, PrimFunc(symshen_4pvar_2), W4768)


if True == tmp15673 {
tmp15666 := MakeNative(func(__e *ControlFlow) {
W4780 := __e.Get(1)
_ = W4780
tmp15667 := PrimCons(W4780, Nil)

tmp15668 := PrimCons(symlist, tmp15667)

tmp15669 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W4769, W4780)
return
}, 0)

tmp15670 := Call(__e, PrimFunc(symshen_4bind_b), W4768, tmp15668, V4751, tmp15669)


__e.TailApply(PrimFunc(symshen_4gc), V4751, tmp15670)
return


}, 1)

tmp15671 := Call(__e, PrimFunc(symshen_4newpv), V4751)


__e.TailApply(tmp15666, tmp15671)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp15676 := MakeNative(func(__e *ControlFlow) {
Z4770 := __e.Get(1)
_ = Z4770
tmp15677 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15677

__e.TailApply(PrimFunc(symthaw), V4754)
return


}, 1)

__e.TailApply(tmp15631, tmp15676)
return


}, 1)

tmp15678 := Call(__e, PrimFunc(symshen_4lazyderef), V4750, V4751)


__e.TailApply(tmp15630, tmp15678)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15681 := Call(__e, PrimFunc(symshen_4lazyderef), V4749, V4751)


__e.TailApply(tmp15629, tmp15681)
return


} else {
__e.Return(False)
return
}


} else {
__e.Return(W4764)
return
}


}, 1)

tmp15699 := Call(__e, PrimFunc(symshen_4unlocked_2), V4752)


var ifres15686 Obj

if True == tmp15699 {
tmp15687 := MakeNative(func(__e *ControlFlow) {
W4765 := __e.Get(1)
_ = W4765
tmp15688 := MakeNative(func(__e *ControlFlow) {
W4766 := __e.Get(1)
_ = W4766
tmp15692 := PrimEqual(W4765, symsymbol)

if True == tmp15692 {
__e.TailApply(PrimFunc(symthaw), W4766)
return
} else {
tmp15690 := Call(__e, PrimFunc(symshen_4pvar_2), W4765)


if True == tmp15690 {
__e.TailApply(PrimFunc(symshen_4bind_b), W4765, symsymbol, V4751, W4766)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp15693 := MakeNative(func(__e *ControlFlow) {
tmp15694 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15694

tmp15695 := Call(__e, PrimFunc(symshen_4lazyderef), V4749, V4751)


tmp15696 := PrimIsSymbol(tmp15695)

__e.TailApply(PrimFunc(symwhen), tmp15696, V4751, V4752, V4753, V4754)
return


}, 0)

__e.TailApply(tmp15688, tmp15693)
return


}, 1)

tmp15697 := Call(__e, PrimFunc(symshen_4lazyderef), V4750, V4751)


tmp15698 := Call(__e, tmp15687, tmp15697)


ifres15686 = tmp15698


} else {
ifres15686 = False


}

__e.TailApply(tmp15628, ifres15686)
return


} else {
__e.Return(W4761)
return
}


}, 1)

tmp15715 := Call(__e, PrimFunc(symshen_4unlocked_2), V4752)


var ifres15702 Obj

if True == tmp15715 {
tmp15703 := MakeNative(func(__e *ControlFlow) {
W4762 := __e.Get(1)
_ = W4762
tmp15704 := MakeNative(func(__e *ControlFlow) {
W4763 := __e.Get(1)
_ = W4763
tmp15708 := PrimEqual(W4762, symstring)

if True == tmp15708 {
__e.TailApply(PrimFunc(symthaw), W4763)
return
} else {
tmp15706 := Call(__e, PrimFunc(symshen_4pvar_2), W4762)


if True == tmp15706 {
__e.TailApply(PrimFunc(symshen_4bind_b), W4762, symstring, V4751, W4763)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp15709 := MakeNative(func(__e *ControlFlow) {
tmp15710 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15710

tmp15711 := Call(__e, PrimFunc(symshen_4lazyderef), V4749, V4751)


tmp15712 := PrimIsString(tmp15711)

__e.TailApply(PrimFunc(symwhen), tmp15712, V4751, V4752, V4753, V4754)
return


}, 0)

__e.TailApply(tmp15704, tmp15709)
return


}, 1)

tmp15713 := Call(__e, PrimFunc(symshen_4lazyderef), V4750, V4751)


tmp15714 := Call(__e, tmp15703, tmp15713)


ifres15702 = tmp15714


} else {
ifres15702 = False


}

__e.TailApply(tmp15627, ifres15702)
return


} else {
__e.Return(W4758)
return
}


}, 1)

tmp15731 := Call(__e, PrimFunc(symshen_4unlocked_2), V4752)


var ifres15718 Obj

if True == tmp15731 {
tmp15719 := MakeNative(func(__e *ControlFlow) {
W4759 := __e.Get(1)
_ = W4759
tmp15720 := MakeNative(func(__e *ControlFlow) {
W4760 := __e.Get(1)
_ = W4760
tmp15724 := PrimEqual(W4759, symboolean)

if True == tmp15724 {
__e.TailApply(PrimFunc(symthaw), W4760)
return
} else {
tmp15722 := Call(__e, PrimFunc(symshen_4pvar_2), W4759)


if True == tmp15722 {
__e.TailApply(PrimFunc(symshen_4bind_b), W4759, symboolean, V4751, W4760)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp15725 := MakeNative(func(__e *ControlFlow) {
tmp15726 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15726

tmp15727 := Call(__e, PrimFunc(symshen_4lazyderef), V4749, V4751)


tmp15728 := Call(__e, PrimFunc(symboolean_2), tmp15727)


__e.TailApply(PrimFunc(symwhen), tmp15728, V4751, V4752, V4753, V4754)
return


}, 0)

__e.TailApply(tmp15720, tmp15725)
return


}, 1)

tmp15729 := Call(__e, PrimFunc(symshen_4lazyderef), V4750, V4751)


tmp15730 := Call(__e, tmp15719, tmp15729)


ifres15718 = tmp15730


} else {
ifres15718 = False


}

__e.TailApply(tmp15626, ifres15718)
return


} else {
__e.Return(W4755)
return
}


}, 1)

tmp15747 := Call(__e, PrimFunc(symshen_4unlocked_2), V4752)


var ifres15734 Obj

if True == tmp15747 {
tmp15735 := MakeNative(func(__e *ControlFlow) {
W4756 := __e.Get(1)
_ = W4756
tmp15736 := MakeNative(func(__e *ControlFlow) {
W4757 := __e.Get(1)
_ = W4757
tmp15740 := PrimEqual(W4756, symnumber)

if True == tmp15740 {
__e.TailApply(PrimFunc(symthaw), W4757)
return
} else {
tmp15738 := Call(__e, PrimFunc(symshen_4pvar_2), W4756)


if True == tmp15738 {
__e.TailApply(PrimFunc(symshen_4bind_b), W4756, symnumber, V4751, W4757)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp15741 := MakeNative(func(__e *ControlFlow) {
tmp15742 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15742

tmp15743 := Call(__e, PrimFunc(symshen_4lazyderef), V4749, V4751)


tmp15744 := PrimIsNumber(tmp15743)

__e.TailApply(PrimFunc(symwhen), tmp15744, V4751, V4752, V4753, V4754)
return


}, 0)

__e.TailApply(tmp15736, tmp15741)
return


}, 1)

tmp15745 := Call(__e, PrimFunc(symshen_4lazyderef), V4750, V4751)


tmp15746 := Call(__e, tmp15735, tmp15745)


ifres15734 = tmp15746


} else {
ifres15734 = False


}

__e.TailApply(tmp15625, ifres15734)
return


}, 6)

tmp15748 := Call(__e, ns2_1set, symshen_4primitive, tmp15624)


_ = tmp15748

tmp15749 := MakeNative(func(__e *ControlFlow) {
V4781 := __e.Get(1)
_ = V4781
V4782 := __e.Get(2)
_ = V4782
V4783 := __e.Get(3)
_ = V4783
V4784 := __e.Get(4)
_ = V4784
V4785 := __e.Get(5)
_ = V4785
V4786 := __e.Get(6)
_ = V4786
V4787 := __e.Get(7)
_ = V4787
tmp15750 := MakeNative(func(__e *ControlFlow) {
W4788 := __e.Get(1)
_ = W4788
tmp15761 := PrimEqual(W4788, False)

if True == tmp15761 {
tmp15759 := Call(__e, PrimFunc(symshen_4unlocked_2), V4785)


if True == tmp15759 {
tmp15751 := MakeNative(func(__e *ControlFlow) {
W4797 := __e.Get(1)
_ = W4797
tmp15756 := PrimIsPair(W4797)

if True == tmp15756 {
tmp15752 := MakeNative(func(__e *ControlFlow) {
W4798 := __e.Get(1)
_ = W4798
tmp15753 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15753

__e.TailApply(PrimFunc(symshen_4by_1hypothesis), V4781, V4782, W4798, V4784, V4785, V4786, V4787)
return


}, 1)

tmp15754 := PrimTail(W4797)

__e.TailApply(tmp15752, tmp15754)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15757 := Call(__e, PrimFunc(symshen_4lazyderef), V4783, V4784)


__e.TailApply(tmp15751, tmp15757)
return


} else {
__e.Return(False)
return
}


} else {
__e.Return(W4788)
return
}


}, 1)

tmp15803 := Call(__e, PrimFunc(symshen_4unlocked_2), V4785)


var ifres15762 Obj

if True == tmp15803 {
tmp15763 := MakeNative(func(__e *ControlFlow) {
W4789 := __e.Get(1)
_ = W4789
tmp15800 := PrimIsPair(W4789)

if True == tmp15800 {
tmp15764 := MakeNative(func(__e *ControlFlow) {
W4790 := __e.Get(1)
_ = W4790
tmp15796 := PrimIsPair(W4790)

if True == tmp15796 {
tmp15765 := MakeNative(func(__e *ControlFlow) {
W4791 := __e.Get(1)
_ = W4791
tmp15766 := MakeNative(func(__e *ControlFlow) {
W4792 := __e.Get(1)
_ = W4792
tmp15791 := PrimIsPair(W4792)

if True == tmp15791 {
tmp15767 := MakeNative(func(__e *ControlFlow) {
W4793 := __e.Get(1)
_ = W4793
tmp15768 := MakeNative(func(__e *ControlFlow) {
W4794 := __e.Get(1)
_ = W4794
tmp15786 := PrimIsPair(W4794)

if True == tmp15786 {
tmp15769 := MakeNative(func(__e *ControlFlow) {
W4795 := __e.Get(1)
_ = W4795
tmp15770 := MakeNative(func(__e *ControlFlow) {
W4796 := __e.Get(1)
_ = W4796
tmp15781 := PrimEqual(W4796, Nil)

if True == tmp15781 {
tmp15771 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15771

tmp15772 := Call(__e, PrimFunc(symshen_4deref), W4793, V4784)


tmp15773 := PrimIntern(MakeString(":"))

tmp15774 := PrimEqual(tmp15772, tmp15773)

tmp15775 := MakeNative(func(__e *ControlFlow) {
tmp15776 := Call(__e, PrimFunc(symshen_4deref), V4781, V4784)


tmp15777 := Call(__e, PrimFunc(symshen_4deref), W4791, V4784)


tmp15778 := PrimEqual(tmp15776, tmp15777)

tmp15779 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symis_b), V4782, W4795, V4784, V4785, V4786, V4787)
return
}, 0)

__e.TailApply(PrimFunc(symwhen), tmp15778, V4784, V4785, V4786, tmp15779)
return


}, 0)

__e.TailApply(PrimFunc(symwhen), tmp15774, V4784, V4785, V4786, tmp15775)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15782 := PrimTail(W4794)

tmp15783 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15782, V4784)


__e.TailApply(tmp15770, tmp15783)
return


}, 1)

tmp15784 := PrimHead(W4794)

__e.TailApply(tmp15769, tmp15784)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15787 := PrimTail(W4792)

tmp15788 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15787, V4784)


__e.TailApply(tmp15768, tmp15788)
return


}, 1)

tmp15789 := PrimHead(W4792)

__e.TailApply(tmp15767, tmp15789)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15792 := PrimTail(W4790)

tmp15793 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15792, V4784)


__e.TailApply(tmp15766, tmp15793)
return


}, 1)

tmp15794 := PrimHead(W4790)

__e.TailApply(tmp15765, tmp15794)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15797 := PrimHead(W4789)

tmp15798 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15797, V4784)


__e.TailApply(tmp15764, tmp15798)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15801 := Call(__e, PrimFunc(symshen_4lazyderef), V4783, V4784)


tmp15802 := Call(__e, tmp15763, tmp15801)


ifres15762 = tmp15802


} else {
ifres15762 = False


}

__e.TailApply(tmp15750, ifres15762)
return


}, 7)

tmp15804 := Call(__e, ns2_1set, symshen_4by_1hypothesis, tmp15749)


_ = tmp15804

tmp15805 := MakeNative(func(__e *ControlFlow) {
V4799 := __e.Get(1)
_ = V4799
V4800 := __e.Get(2)
_ = V4800
V4801 := __e.Get(3)
_ = V4801
V4802 := __e.Get(4)
_ = V4802
V4803 := __e.Get(5)
_ = V4803
V4804 := __e.Get(6)
_ = V4804
tmp15810 := Call(__e, PrimFunc(symshen_4unlocked_2), V4802)


if True == tmp15810 {
tmp15806 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15806

tmp15807 := PrimValue(symshen_4_dsigf_d)

tmp15808 := Call(__e, PrimFunc(symassoc), V4799, tmp15807)


__e.TailApply(PrimFunc(symshen_4sigf), tmp15808, V4800, V4801, V4802, V4803, V4804)
return


} else {
__e.Return(False)
return
}


}, 6)

tmp15811 := Call(__e, ns2_1set, symshen_4lookupsig, tmp15805)


_ = tmp15811

tmp15812 := MakeNative(func(__e *ControlFlow) {
V4819 := __e.Get(1)
_ = V4819
V4820 := __e.Get(2)
_ = V4820
V4821 := __e.Get(3)
_ = V4821
V4822 := __e.Get(4)
_ = V4822
V4823 := __e.Get(5)
_ = V4823
V4824 := __e.Get(6)
_ = V4824
tmp15819 := PrimIsPair(V4819)

if True == tmp15819 {
tmp15813 := PrimTail(V4819)

tmp15814 := Call(__e, tmp15813, V4820)


tmp15815 := Call(__e, tmp15814, V4821)


tmp15816 := Call(__e, tmp15815, V4822)


tmp15817 := Call(__e, tmp15816, V4823)


__e.TailApply(tmp15817, V4824)
return


} else {
__e.Return(False)
return
}


}, 6)

tmp15820 := Call(__e, ns2_1set, symshen_4sigf, tmp15812)


_ = tmp15820

tmp15821 := MakeNative(func(__e *ControlFlow) {
V4825 := __e.Get(1)
_ = V4825
tmp15822 := MakeNative(func(__e *ControlFlow) {
W4826 := __e.Get(1)
_ = W4826
tmp15823 := MakeNative(func(__e *ControlFlow) {
W4827 := __e.Get(1)
_ = W4827
tmp15824 := MakeNative(func(__e *ControlFlow) {
W4828 := __e.Get(1)
_ = W4828
tmp15825 := MakeNative(func(__e *ControlFlow) {
W4829 := __e.Get(1)
_ = W4829
__e.Return(W4829)
return
}, 1)

tmp15826 := PrimValue(symshen_4_dgensym_d)

tmp15827 := PrimNumberAdd(MakeNumber(1), tmp15826)

tmp15828 := PrimSet(symshen_4_dgensym_d, tmp15827)

tmp15829 := PrimVectorSet(W4828, MakeNumber(2), tmp15828)

__e.TailApply(tmp15825, tmp15829)
return


}, 1)

tmp15830 := PrimVectorSet(W4827, MakeNumber(1), V4825)

__e.TailApply(tmp15824, tmp15830)
return


}, 1)

tmp15831 := PrimVectorSet(W4826, MakeNumber(0), symshen_4print_1freshterm)

__e.TailApply(tmp15823, tmp15831)
return


}, 1)

tmp15832 := PrimAbsvector(MakeNumber(3))

__e.TailApply(tmp15822, tmp15832)
return


}, 1)

tmp15833 := Call(__e, ns2_1set, symshen_4freshterm, tmp15821)


_ = tmp15833

tmp15834 := MakeNative(func(__e *ControlFlow) {
V4830 := __e.Get(1)
_ = V4830
tmp15835 := PrimVectorGet(V4830, MakeNumber(1))

tmp15836 := Call(__e, PrimFunc(symshen_4app), tmp15835, MakeString(""), symshen_4a)


__e.Return(PrimStringConcat(MakeString("&&"), tmp15836))
return


}, 1)

tmp15837 := Call(__e, ns2_1set, symshen_4print_1freshterm, tmp15834)


_ = tmp15837

tmp15838 := MakeNative(func(__e *ControlFlow) {
V4831 := __e.Get(1)
_ = V4831
V4832 := __e.Get(2)
_ = V4832
V4833 := __e.Get(3)
_ = V4833
V4834 := __e.Get(4)
_ = V4834
V4835 := __e.Get(5)
_ = V4835
V4836 := __e.Get(6)
_ = V4836
V4837 := __e.Get(7)
_ = V4837
tmp15839 := MakeNative(func(__e *ControlFlow) {
W4838 := __e.Get(1)
_ = W4838
tmp15850 := PrimEqual(W4838, False)

if True == tmp15850 {
tmp15848 := Call(__e, PrimFunc(symshen_4unlocked_2), V4835)


if True == tmp15848 {
tmp15840 := MakeNative(func(__e *ControlFlow) {
W4842 := __e.Get(1)
_ = W4842
tmp15845 := PrimIsPair(W4842)

if True == tmp15845 {
tmp15841 := MakeNative(func(__e *ControlFlow) {
W4843 := __e.Get(1)
_ = W4843
tmp15842 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15842

__e.TailApply(PrimFunc(symshen_4search_1user_1datatypes), V4831, V4832, W4843, V4834, V4835, V4836, V4837)
return


}, 1)

tmp15843 := PrimTail(W4842)

__e.TailApply(tmp15841, tmp15843)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15846 := Call(__e, PrimFunc(symshen_4lazyderef), V4833, V4834)


__e.TailApply(tmp15840, tmp15846)
return


} else {
__e.Return(False)
return
}


} else {
__e.Return(W4838)
return
}


}, 1)

tmp15870 := Call(__e, PrimFunc(symshen_4unlocked_2), V4835)


var ifres15851 Obj

if True == tmp15870 {
tmp15852 := MakeNative(func(__e *ControlFlow) {
W4839 := __e.Get(1)
_ = W4839
tmp15867 := PrimIsPair(W4839)

if True == tmp15867 {
tmp15853 := MakeNative(func(__e *ControlFlow) {
W4840 := __e.Get(1)
_ = W4840
tmp15863 := PrimIsPair(W4840)

if True == tmp15863 {
tmp15854 := MakeNative(func(__e *ControlFlow) {
W4841 := __e.Get(1)
_ = W4841
tmp15855 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15855

tmp15856 := Call(__e, PrimFunc(symshen_4deref), W4841, V4834)


tmp15857 := Call(__e, PrimFunc(symshen_4deref), V4831, V4834)


tmp15858 := Call(__e, tmp15856, tmp15857)


tmp15859 := Call(__e, PrimFunc(symshen_4deref), V4832, V4834)


tmp15860 := Call(__e, tmp15858, tmp15859)


__e.TailApply(PrimFunc(symcall), tmp15860, V4834, V4835, V4836, V4837)
return


}, 1)

tmp15861 := PrimTail(W4840)

__e.TailApply(tmp15854, tmp15861)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15864 := PrimHead(W4839)

tmp15865 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15864, V4834)


__e.TailApply(tmp15853, tmp15865)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15868 := Call(__e, PrimFunc(symshen_4lazyderef), V4833, V4834)


tmp15869 := Call(__e, tmp15852, tmp15868)


ifres15851 = tmp15869


} else {
ifres15851 = False


}

__e.TailApply(tmp15839, ifres15851)
return


}, 7)

tmp15871 := Call(__e, ns2_1set, symshen_4search_1user_1datatypes, tmp15838)


_ = tmp15871

tmp15872 := MakeNative(func(__e *ControlFlow) {
V4844 := __e.Get(1)
_ = V4844
V4845 := __e.Get(2)
_ = V4845
V4846 := __e.Get(3)
_ = V4846
V4847 := __e.Get(4)
_ = V4847
V4848 := __e.Get(5)
_ = V4848
V4849 := __e.Get(6)
_ = V4849
V4850 := __e.Get(7)
_ = V4850
tmp15873 := MakeNative(func(__e *ControlFlow) {
W4851 := __e.Get(1)
_ = W4851
tmp15874 := MakeNative(func(__e *ControlFlow) {
W4852 := __e.Get(1)
_ = W4852
tmp16304 := PrimEqual(W4852, False)

if True == tmp16304 {
tmp15875 := MakeNative(func(__e *ControlFlow) {
W4855 := __e.Get(1)
_ = W4855
tmp16204 := PrimEqual(W4855, False)

if True == tmp16204 {
tmp15876 := MakeNative(func(__e *ControlFlow) {
W4875 := __e.Get(1)
_ = W4875
tmp16099 := PrimEqual(W4875, False)

if True == tmp16099 {
tmp15877 := MakeNative(func(__e *ControlFlow) {
W4897 := __e.Get(1)
_ = W4897
tmp16018 := PrimEqual(W4897, False)

if True == tmp16018 {
tmp15878 := MakeNative(func(__e *ControlFlow) {
W4913 := __e.Get(1)
_ = W4913
tmp15918 := PrimEqual(W4913, False)

if True == tmp15918 {
tmp15879 := MakeNative(func(__e *ControlFlow) {
W4933 := __e.Get(1)
_ = W4933
tmp15881 := PrimEqual(W4933, False)

if True == tmp15881 {
__e.TailApply(PrimFunc(symshen_4unlock), V4848, W4851)
return
} else {
__e.Return(W4933)
return
}


}, 1)

tmp15916 := Call(__e, PrimFunc(symshen_4unlocked_2), V4848)


var ifres15882 Obj

if True == tmp15916 {
tmp15883 := MakeNative(func(__e *ControlFlow) {
W4934 := __e.Get(1)
_ = W4934
tmp15913 := PrimIsPair(W4934)

if True == tmp15913 {
tmp15884 := MakeNative(func(__e *ControlFlow) {
W4935 := __e.Get(1)
_ = W4935
tmp15885 := MakeNative(func(__e *ControlFlow) {
W4936 := __e.Get(1)
_ = W4936
tmp15886 := MakeNative(func(__e *ControlFlow) {
W4937 := __e.Get(1)
_ = W4937
tmp15887 := MakeNative(func(__e *ControlFlow) {
W4938 := __e.Get(1)
_ = W4938
tmp15905 := PrimIsPair(W4937)

if True == tmp15905 {
tmp15888 := MakeNative(func(__e *ControlFlow) {
W4941 := __e.Get(1)
_ = W4941
tmp15889 := MakeNative(func(__e *ControlFlow) {
W4942 := __e.Get(1)
_ = W4942
tmp15890 := Call(__e, W4938, W4941)


__e.TailApply(tmp15890, W4942)
return


}, 1)

tmp15891 := PrimTail(W4937)

__e.TailApply(tmp15889, tmp15891)
return


}, 1)

tmp15892 := PrimHead(W4937)

__e.TailApply(tmp15888, tmp15892)
return


} else {
tmp15903 := Call(__e, PrimFunc(symshen_4pvar_2), W4937)


if True == tmp15903 {
tmp15893 := MakeNative(func(__e *ControlFlow) {
W4943 := __e.Get(1)
_ = W4943
tmp15894 := MakeNative(func(__e *ControlFlow) {
W4944 := __e.Get(1)
_ = W4944
tmp15895 := PrimCons(W4943, W4944)

tmp15896 := MakeNative(func(__e *ControlFlow) {
tmp15897 := Call(__e, W4938, W4943)


__e.TailApply(tmp15897, W4944)
return


}, 0)

tmp15898 := Call(__e, PrimFunc(symshen_4bind_b), W4937, tmp15895, V4847, tmp15896)


__e.TailApply(PrimFunc(symshen_4gc), V4847, tmp15898)
return


}, 1)

tmp15899 := Call(__e, PrimFunc(symshen_4newpv), V4847)


tmp15900 := Call(__e, tmp15894, tmp15899)


__e.TailApply(PrimFunc(symshen_4gc), V4847, tmp15900)
return


}, 1)

tmp15901 := Call(__e, PrimFunc(symshen_4newpv), V4847)


__e.TailApply(tmp15893, tmp15901)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp15906 := MakeNative(func(__e *ControlFlow) {
Z4939 := __e.Get(1)
_ = Z4939
__e.Return(MakeNative(func(__e *ControlFlow) {
Z4940 := __e.Get(1)
_ = Z4940
tmp15907 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15907

tmp15908 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4l_1rules), W4936, Z4940, V4846, V4847, V4848, W4851, V4850)
return
}, 0)

__e.TailApply(PrimFunc(symbind), Z4939, W4935, V4847, V4848, W4851, tmp15908)
return


}, 1))
return
}, 1)

__e.TailApply(tmp15887, tmp15906)
return


}, 1)

tmp15909 := Call(__e, PrimFunc(symshen_4lazyderef), V4845, V4847)


__e.TailApply(tmp15886, tmp15909)
return


}, 1)

tmp15910 := PrimTail(W4934)

__e.TailApply(tmp15885, tmp15910)
return


}, 1)

tmp15911 := PrimHead(W4934)

__e.TailApply(tmp15884, tmp15911)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15914 := Call(__e, PrimFunc(symshen_4lazyderef), V4844, V4847)


tmp15915 := Call(__e, tmp15883, tmp15914)


ifres15882 = tmp15915


} else {
ifres15882 = False


}

__e.TailApply(tmp15879, ifres15882)
return


} else {
__e.Return(W4913)
return
}


}, 1)

tmp16016 := Call(__e, PrimFunc(symshen_4unlocked_2), V4848)


var ifres15919 Obj

if True == tmp16016 {
tmp15920 := MakeNative(func(__e *ControlFlow) {
W4914 := __e.Get(1)
_ = W4914
tmp16013 := PrimIsPair(W4914)

if True == tmp16013 {
tmp15921 := MakeNative(func(__e *ControlFlow) {
W4915 := __e.Get(1)
_ = W4915
tmp16009 := PrimIsPair(W4915)

if True == tmp16009 {
tmp15922 := MakeNative(func(__e *ControlFlow) {
W4916 := __e.Get(1)
_ = W4916
tmp16005 := PrimIsPair(W4916)

if True == tmp16005 {
tmp15923 := MakeNative(func(__e *ControlFlow) {
W4917 := __e.Get(1)
_ = W4917
tmp16001 := PrimEqual(W4917, sym_8v)

if True == tmp16001 {
tmp15924 := MakeNative(func(__e *ControlFlow) {
W4918 := __e.Get(1)
_ = W4918
tmp15997 := PrimIsPair(W4918)

if True == tmp15997 {
tmp15925 := MakeNative(func(__e *ControlFlow) {
W4919 := __e.Get(1)
_ = W4919
tmp15926 := MakeNative(func(__e *ControlFlow) {
W4920 := __e.Get(1)
_ = W4920
tmp15992 := PrimIsPair(W4920)

if True == tmp15992 {
tmp15927 := MakeNative(func(__e *ControlFlow) {
W4921 := __e.Get(1)
_ = W4921
tmp15928 := MakeNative(func(__e *ControlFlow) {
W4922 := __e.Get(1)
_ = W4922
tmp15987 := PrimEqual(W4922, Nil)

if True == tmp15987 {
tmp15929 := MakeNative(func(__e *ControlFlow) {
W4923 := __e.Get(1)
_ = W4923
tmp15983 := PrimIsPair(W4923)

if True == tmp15983 {
tmp15930 := MakeNative(func(__e *ControlFlow) {
W4924 := __e.Get(1)
_ = W4924
tmp15931 := MakeNative(func(__e *ControlFlow) {
W4925 := __e.Get(1)
_ = W4925
tmp15978 := PrimIsPair(W4925)

if True == tmp15978 {
tmp15932 := MakeNative(func(__e *ControlFlow) {
W4926 := __e.Get(1)
_ = W4926
tmp15974 := PrimIsPair(W4926)

if True == tmp15974 {
tmp15933 := MakeNative(func(__e *ControlFlow) {
W4927 := __e.Get(1)
_ = W4927
tmp15970 := PrimEqual(W4927, symvector)

if True == tmp15970 {
tmp15934 := MakeNative(func(__e *ControlFlow) {
W4928 := __e.Get(1)
_ = W4928
tmp15966 := PrimIsPair(W4928)

if True == tmp15966 {
tmp15935 := MakeNative(func(__e *ControlFlow) {
W4929 := __e.Get(1)
_ = W4929
tmp15936 := MakeNative(func(__e *ControlFlow) {
W4930 := __e.Get(1)
_ = W4930
tmp15961 := PrimEqual(W4930, Nil)

if True == tmp15961 {
tmp15937 := MakeNative(func(__e *ControlFlow) {
W4931 := __e.Get(1)
_ = W4931
tmp15957 := PrimEqual(W4931, Nil)

if True == tmp15957 {
tmp15938 := MakeNative(func(__e *ControlFlow) {
W4932 := __e.Get(1)
_ = W4932
tmp15939 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15939

tmp15940 := Call(__e, PrimFunc(symshen_4deref), W4924, V4847)


tmp15941 := PrimIntern(MakeString(":"))

tmp15942 := PrimEqual(tmp15940, tmp15941)

tmp15943 := MakeNative(func(__e *ControlFlow) {
tmp15944 := MakeNative(func(__e *ControlFlow) {
tmp15945 := PrimCons(W4929, Nil)

tmp15946 := PrimCons(W4924, tmp15945)

tmp15947 := PrimCons(W4919, tmp15946)

tmp15948 := PrimCons(W4929, Nil)

tmp15949 := PrimCons(symvector, tmp15948)

tmp15950 := PrimCons(tmp15949, Nil)

tmp15951 := PrimCons(W4924, tmp15950)

tmp15952 := PrimCons(W4921, tmp15951)

tmp15953 := PrimCons(tmp15952, W4932)

tmp15954 := PrimCons(tmp15947, tmp15953)

__e.TailApply(PrimFunc(symshen_4l_1rules), tmp15954, V4845, True, V4847, V4848, W4851, V4850)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V4847, V4848, W4851, tmp15944)
return


}, 0)

__e.TailApply(PrimFunc(symwhen), tmp15942, V4847, V4848, W4851, tmp15943)
return


}, 1)

tmp15955 := PrimTail(W4914)

__e.TailApply(tmp15938, tmp15955)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15958 := PrimTail(W4925)

tmp15959 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15958, V4847)


__e.TailApply(tmp15937, tmp15959)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15962 := PrimTail(W4928)

tmp15963 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15962, V4847)


__e.TailApply(tmp15936, tmp15963)
return


}, 1)

tmp15964 := PrimHead(W4928)

__e.TailApply(tmp15935, tmp15964)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15967 := PrimTail(W4926)

tmp15968 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15967, V4847)


__e.TailApply(tmp15934, tmp15968)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15971 := PrimHead(W4926)

tmp15972 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15971, V4847)


__e.TailApply(tmp15933, tmp15972)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15975 := PrimHead(W4925)

tmp15976 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15975, V4847)


__e.TailApply(tmp15932, tmp15976)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15979 := PrimTail(W4923)

tmp15980 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15979, V4847)


__e.TailApply(tmp15931, tmp15980)
return


}, 1)

tmp15981 := PrimHead(W4923)

__e.TailApply(tmp15930, tmp15981)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15984 := PrimTail(W4915)

tmp15985 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15984, V4847)


__e.TailApply(tmp15929, tmp15985)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15988 := PrimTail(W4920)

tmp15989 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15988, V4847)


__e.TailApply(tmp15928, tmp15989)
return


}, 1)

tmp15990 := PrimHead(W4920)

__e.TailApply(tmp15927, tmp15990)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15993 := PrimTail(W4918)

tmp15994 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15993, V4847)


__e.TailApply(tmp15926, tmp15994)
return


}, 1)

tmp15995 := PrimHead(W4918)

__e.TailApply(tmp15925, tmp15995)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15998 := PrimTail(W4916)

tmp15999 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15998, V4847)


__e.TailApply(tmp15924, tmp15999)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16002 := PrimHead(W4916)

tmp16003 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16002, V4847)


__e.TailApply(tmp15923, tmp16003)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16006 := PrimHead(W4915)

tmp16007 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16006, V4847)


__e.TailApply(tmp15922, tmp16007)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16010 := PrimHead(W4914)

tmp16011 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16010, V4847)


__e.TailApply(tmp15921, tmp16011)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16014 := Call(__e, PrimFunc(symshen_4lazyderef), V4844, V4847)


tmp16015 := Call(__e, tmp15920, tmp16014)


ifres15919 = tmp16015


} else {
ifres15919 = False


}

__e.TailApply(tmp15878, ifres15919)
return


} else {
__e.Return(W4897)
return
}


}, 1)

tmp16097 := Call(__e, PrimFunc(symshen_4unlocked_2), V4848)


var ifres16019 Obj

if True == tmp16097 {
tmp16020 := MakeNative(func(__e *ControlFlow) {
W4898 := __e.Get(1)
_ = W4898
tmp16094 := PrimIsPair(W4898)

if True == tmp16094 {
tmp16021 := MakeNative(func(__e *ControlFlow) {
W4899 := __e.Get(1)
_ = W4899
tmp16090 := PrimIsPair(W4899)

if True == tmp16090 {
tmp16022 := MakeNative(func(__e *ControlFlow) {
W4900 := __e.Get(1)
_ = W4900
tmp16086 := PrimIsPair(W4900)

if True == tmp16086 {
tmp16023 := MakeNative(func(__e *ControlFlow) {
W4901 := __e.Get(1)
_ = W4901
tmp16082 := PrimEqual(W4901, sym_8s)

if True == tmp16082 {
tmp16024 := MakeNative(func(__e *ControlFlow) {
W4902 := __e.Get(1)
_ = W4902
tmp16078 := PrimIsPair(W4902)

if True == tmp16078 {
tmp16025 := MakeNative(func(__e *ControlFlow) {
W4903 := __e.Get(1)
_ = W4903
tmp16026 := MakeNative(func(__e *ControlFlow) {
W4904 := __e.Get(1)
_ = W4904
tmp16073 := PrimIsPair(W4904)

if True == tmp16073 {
tmp16027 := MakeNative(func(__e *ControlFlow) {
W4905 := __e.Get(1)
_ = W4905
tmp16028 := MakeNative(func(__e *ControlFlow) {
W4906 := __e.Get(1)
_ = W4906
tmp16068 := PrimEqual(W4906, Nil)

if True == tmp16068 {
tmp16029 := MakeNative(func(__e *ControlFlow) {
W4907 := __e.Get(1)
_ = W4907
tmp16064 := PrimIsPair(W4907)

if True == tmp16064 {
tmp16030 := MakeNative(func(__e *ControlFlow) {
W4908 := __e.Get(1)
_ = W4908
tmp16031 := MakeNative(func(__e *ControlFlow) {
W4909 := __e.Get(1)
_ = W4909
tmp16059 := PrimIsPair(W4909)

if True == tmp16059 {
tmp16032 := MakeNative(func(__e *ControlFlow) {
W4910 := __e.Get(1)
_ = W4910
tmp16055 := PrimEqual(W4910, symstring)

if True == tmp16055 {
tmp16033 := MakeNative(func(__e *ControlFlow) {
W4911 := __e.Get(1)
_ = W4911
tmp16051 := PrimEqual(W4911, Nil)

if True == tmp16051 {
tmp16034 := MakeNative(func(__e *ControlFlow) {
W4912 := __e.Get(1)
_ = W4912
tmp16035 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16035

tmp16036 := Call(__e, PrimFunc(symshen_4deref), W4908, V4847)


tmp16037 := PrimIntern(MakeString(":"))

tmp16038 := PrimEqual(tmp16036, tmp16037)

tmp16039 := MakeNative(func(__e *ControlFlow) {
tmp16040 := MakeNative(func(__e *ControlFlow) {
tmp16041 := PrimCons(symstring, Nil)

tmp16042 := PrimCons(W4908, tmp16041)

tmp16043 := PrimCons(W4903, tmp16042)

tmp16044 := PrimCons(symstring, Nil)

tmp16045 := PrimCons(W4908, tmp16044)

tmp16046 := PrimCons(W4905, tmp16045)

tmp16047 := PrimCons(tmp16046, W4912)

tmp16048 := PrimCons(tmp16043, tmp16047)

__e.TailApply(PrimFunc(symshen_4l_1rules), tmp16048, V4845, True, V4847, V4848, W4851, V4850)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V4847, V4848, W4851, tmp16040)
return


}, 0)

__e.TailApply(PrimFunc(symwhen), tmp16038, V4847, V4848, W4851, tmp16039)
return


}, 1)

tmp16049 := PrimTail(W4898)

__e.TailApply(tmp16034, tmp16049)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16052 := PrimTail(W4909)

tmp16053 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16052, V4847)


__e.TailApply(tmp16033, tmp16053)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16056 := PrimHead(W4909)

tmp16057 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16056, V4847)


__e.TailApply(tmp16032, tmp16057)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16060 := PrimTail(W4907)

tmp16061 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16060, V4847)


__e.TailApply(tmp16031, tmp16061)
return


}, 1)

tmp16062 := PrimHead(W4907)

__e.TailApply(tmp16030, tmp16062)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16065 := PrimTail(W4899)

tmp16066 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16065, V4847)


__e.TailApply(tmp16029, tmp16066)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16069 := PrimTail(W4904)

tmp16070 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16069, V4847)


__e.TailApply(tmp16028, tmp16070)
return


}, 1)

tmp16071 := PrimHead(W4904)

__e.TailApply(tmp16027, tmp16071)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16074 := PrimTail(W4902)

tmp16075 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16074, V4847)


__e.TailApply(tmp16026, tmp16075)
return


}, 1)

tmp16076 := PrimHead(W4902)

__e.TailApply(tmp16025, tmp16076)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16079 := PrimTail(W4900)

tmp16080 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16079, V4847)


__e.TailApply(tmp16024, tmp16080)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16083 := PrimHead(W4900)

tmp16084 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16083, V4847)


__e.TailApply(tmp16023, tmp16084)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16087 := PrimHead(W4899)

tmp16088 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16087, V4847)


__e.TailApply(tmp16022, tmp16088)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16091 := PrimHead(W4898)

tmp16092 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16091, V4847)


__e.TailApply(tmp16021, tmp16092)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16095 := Call(__e, PrimFunc(symshen_4lazyderef), V4844, V4847)


tmp16096 := Call(__e, tmp16020, tmp16095)


ifres16019 = tmp16096


} else {
ifres16019 = False


}

__e.TailApply(tmp15877, ifres16019)
return


} else {
__e.Return(W4875)
return
}


}, 1)

tmp16202 := Call(__e, PrimFunc(symshen_4unlocked_2), V4848)


var ifres16100 Obj

if True == tmp16202 {
tmp16101 := MakeNative(func(__e *ControlFlow) {
W4876 := __e.Get(1)
_ = W4876
tmp16199 := PrimIsPair(W4876)

if True == tmp16199 {
tmp16102 := MakeNative(func(__e *ControlFlow) {
W4877 := __e.Get(1)
_ = W4877
tmp16195 := PrimIsPair(W4877)

if True == tmp16195 {
tmp16103 := MakeNative(func(__e *ControlFlow) {
W4878 := __e.Get(1)
_ = W4878
tmp16191 := PrimIsPair(W4878)

if True == tmp16191 {
tmp16104 := MakeNative(func(__e *ControlFlow) {
W4879 := __e.Get(1)
_ = W4879
tmp16187 := PrimEqual(W4879, sym_8p)

if True == tmp16187 {
tmp16105 := MakeNative(func(__e *ControlFlow) {
W4880 := __e.Get(1)
_ = W4880
tmp16183 := PrimIsPair(W4880)

if True == tmp16183 {
tmp16106 := MakeNative(func(__e *ControlFlow) {
W4881 := __e.Get(1)
_ = W4881
tmp16107 := MakeNative(func(__e *ControlFlow) {
W4882 := __e.Get(1)
_ = W4882
tmp16178 := PrimIsPair(W4882)

if True == tmp16178 {
tmp16108 := MakeNative(func(__e *ControlFlow) {
W4883 := __e.Get(1)
_ = W4883
tmp16109 := MakeNative(func(__e *ControlFlow) {
W4884 := __e.Get(1)
_ = W4884
tmp16173 := PrimEqual(W4884, Nil)

if True == tmp16173 {
tmp16110 := MakeNative(func(__e *ControlFlow) {
W4885 := __e.Get(1)
_ = W4885
tmp16169 := PrimIsPair(W4885)

if True == tmp16169 {
tmp16111 := MakeNative(func(__e *ControlFlow) {
W4886 := __e.Get(1)
_ = W4886
tmp16112 := MakeNative(func(__e *ControlFlow) {
W4887 := __e.Get(1)
_ = W4887
tmp16164 := PrimIsPair(W4887)

if True == tmp16164 {
tmp16113 := MakeNative(func(__e *ControlFlow) {
W4888 := __e.Get(1)
_ = W4888
tmp16160 := PrimIsPair(W4888)

if True == tmp16160 {
tmp16114 := MakeNative(func(__e *ControlFlow) {
W4889 := __e.Get(1)
_ = W4889
tmp16115 := MakeNative(func(__e *ControlFlow) {
W4890 := __e.Get(1)
_ = W4890
tmp16155 := PrimIsPair(W4890)

if True == tmp16155 {
tmp16116 := MakeNative(func(__e *ControlFlow) {
W4891 := __e.Get(1)
_ = W4891
tmp16151 := PrimEqual(W4891, sym_d)

if True == tmp16151 {
tmp16117 := MakeNative(func(__e *ControlFlow) {
W4892 := __e.Get(1)
_ = W4892
tmp16147 := PrimIsPair(W4892)

if True == tmp16147 {
tmp16118 := MakeNative(func(__e *ControlFlow) {
W4893 := __e.Get(1)
_ = W4893
tmp16119 := MakeNative(func(__e *ControlFlow) {
W4894 := __e.Get(1)
_ = W4894
tmp16142 := PrimEqual(W4894, Nil)

if True == tmp16142 {
tmp16120 := MakeNative(func(__e *ControlFlow) {
W4895 := __e.Get(1)
_ = W4895
tmp16138 := PrimEqual(W4895, Nil)

if True == tmp16138 {
tmp16121 := MakeNative(func(__e *ControlFlow) {
W4896 := __e.Get(1)
_ = W4896
tmp16122 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16122

tmp16123 := Call(__e, PrimFunc(symshen_4deref), W4886, V4847)


tmp16124 := PrimIntern(MakeString(":"))

tmp16125 := PrimEqual(tmp16123, tmp16124)

tmp16126 := MakeNative(func(__e *ControlFlow) {
tmp16127 := MakeNative(func(__e *ControlFlow) {
tmp16128 := PrimCons(W4889, Nil)

tmp16129 := PrimCons(W4886, tmp16128)

tmp16130 := PrimCons(W4881, tmp16129)

tmp16131 := PrimCons(W4893, Nil)

tmp16132 := PrimCons(W4886, tmp16131)

tmp16133 := PrimCons(W4883, tmp16132)

tmp16134 := PrimCons(tmp16133, W4896)

tmp16135 := PrimCons(tmp16130, tmp16134)

__e.TailApply(PrimFunc(symshen_4l_1rules), tmp16135, V4845, True, V4847, V4848, W4851, V4850)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V4847, V4848, W4851, tmp16127)
return


}, 0)

__e.TailApply(PrimFunc(symwhen), tmp16125, V4847, V4848, W4851, tmp16126)
return


}, 1)

tmp16136 := PrimTail(W4876)

__e.TailApply(tmp16121, tmp16136)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16139 := PrimTail(W4887)

tmp16140 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16139, V4847)


__e.TailApply(tmp16120, tmp16140)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16143 := PrimTail(W4892)

tmp16144 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16143, V4847)


__e.TailApply(tmp16119, tmp16144)
return


}, 1)

tmp16145 := PrimHead(W4892)

__e.TailApply(tmp16118, tmp16145)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16148 := PrimTail(W4890)

tmp16149 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16148, V4847)


__e.TailApply(tmp16117, tmp16149)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16152 := PrimHead(W4890)

tmp16153 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16152, V4847)


__e.TailApply(tmp16116, tmp16153)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16156 := PrimTail(W4888)

tmp16157 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16156, V4847)


__e.TailApply(tmp16115, tmp16157)
return


}, 1)

tmp16158 := PrimHead(W4888)

__e.TailApply(tmp16114, tmp16158)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16161 := PrimHead(W4887)

tmp16162 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16161, V4847)


__e.TailApply(tmp16113, tmp16162)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16165 := PrimTail(W4885)

tmp16166 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16165, V4847)


__e.TailApply(tmp16112, tmp16166)
return


}, 1)

tmp16167 := PrimHead(W4885)

__e.TailApply(tmp16111, tmp16167)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16170 := PrimTail(W4877)

tmp16171 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16170, V4847)


__e.TailApply(tmp16110, tmp16171)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16174 := PrimTail(W4882)

tmp16175 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16174, V4847)


__e.TailApply(tmp16109, tmp16175)
return


}, 1)

tmp16176 := PrimHead(W4882)

__e.TailApply(tmp16108, tmp16176)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16179 := PrimTail(W4880)

tmp16180 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16179, V4847)


__e.TailApply(tmp16107, tmp16180)
return


}, 1)

tmp16181 := PrimHead(W4880)

__e.TailApply(tmp16106, tmp16181)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16184 := PrimTail(W4878)

tmp16185 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16184, V4847)


__e.TailApply(tmp16105, tmp16185)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16188 := PrimHead(W4878)

tmp16189 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16188, V4847)


__e.TailApply(tmp16104, tmp16189)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16192 := PrimHead(W4877)

tmp16193 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16192, V4847)


__e.TailApply(tmp16103, tmp16193)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16196 := PrimHead(W4876)

tmp16197 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16196, V4847)


__e.TailApply(tmp16102, tmp16197)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16200 := Call(__e, PrimFunc(symshen_4lazyderef), V4844, V4847)


tmp16201 := Call(__e, tmp16101, tmp16200)


ifres16100 = tmp16201


} else {
ifres16100 = False


}

__e.TailApply(tmp15876, ifres16100)
return


} else {
__e.Return(W4855)
return
}


}, 1)

tmp16302 := Call(__e, PrimFunc(symshen_4unlocked_2), V4848)


var ifres16205 Obj

if True == tmp16302 {
tmp16206 := MakeNative(func(__e *ControlFlow) {
W4856 := __e.Get(1)
_ = W4856
tmp16299 := PrimIsPair(W4856)

if True == tmp16299 {
tmp16207 := MakeNative(func(__e *ControlFlow) {
W4857 := __e.Get(1)
_ = W4857
tmp16295 := PrimIsPair(W4857)

if True == tmp16295 {
tmp16208 := MakeNative(func(__e *ControlFlow) {
W4858 := __e.Get(1)
_ = W4858
tmp16291 := PrimIsPair(W4858)

if True == tmp16291 {
tmp16209 := MakeNative(func(__e *ControlFlow) {
W4859 := __e.Get(1)
_ = W4859
tmp16287 := PrimEqual(W4859, symcons)

if True == tmp16287 {
tmp16210 := MakeNative(func(__e *ControlFlow) {
W4860 := __e.Get(1)
_ = W4860
tmp16283 := PrimIsPair(W4860)

if True == tmp16283 {
tmp16211 := MakeNative(func(__e *ControlFlow) {
W4861 := __e.Get(1)
_ = W4861
tmp16212 := MakeNative(func(__e *ControlFlow) {
W4862 := __e.Get(1)
_ = W4862
tmp16278 := PrimIsPair(W4862)

if True == tmp16278 {
tmp16213 := MakeNative(func(__e *ControlFlow) {
W4863 := __e.Get(1)
_ = W4863
tmp16214 := MakeNative(func(__e *ControlFlow) {
W4864 := __e.Get(1)
_ = W4864
tmp16273 := PrimEqual(W4864, Nil)

if True == tmp16273 {
tmp16215 := MakeNative(func(__e *ControlFlow) {
W4865 := __e.Get(1)
_ = W4865
tmp16269 := PrimIsPair(W4865)

if True == tmp16269 {
tmp16216 := MakeNative(func(__e *ControlFlow) {
W4866 := __e.Get(1)
_ = W4866
tmp16217 := MakeNative(func(__e *ControlFlow) {
W4867 := __e.Get(1)
_ = W4867
tmp16264 := PrimIsPair(W4867)

if True == tmp16264 {
tmp16218 := MakeNative(func(__e *ControlFlow) {
W4868 := __e.Get(1)
_ = W4868
tmp16260 := PrimIsPair(W4868)

if True == tmp16260 {
tmp16219 := MakeNative(func(__e *ControlFlow) {
W4869 := __e.Get(1)
_ = W4869
tmp16256 := PrimEqual(W4869, symlist)

if True == tmp16256 {
tmp16220 := MakeNative(func(__e *ControlFlow) {
W4870 := __e.Get(1)
_ = W4870
tmp16252 := PrimIsPair(W4870)

if True == tmp16252 {
tmp16221 := MakeNative(func(__e *ControlFlow) {
W4871 := __e.Get(1)
_ = W4871
tmp16222 := MakeNative(func(__e *ControlFlow) {
W4872 := __e.Get(1)
_ = W4872
tmp16247 := PrimEqual(W4872, Nil)

if True == tmp16247 {
tmp16223 := MakeNative(func(__e *ControlFlow) {
W4873 := __e.Get(1)
_ = W4873
tmp16243 := PrimEqual(W4873, Nil)

if True == tmp16243 {
tmp16224 := MakeNative(func(__e *ControlFlow) {
W4874 := __e.Get(1)
_ = W4874
tmp16225 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16225

tmp16226 := Call(__e, PrimFunc(symshen_4deref), W4866, V4847)


tmp16227 := PrimIntern(MakeString(":"))

tmp16228 := PrimEqual(tmp16226, tmp16227)

tmp16229 := MakeNative(func(__e *ControlFlow) {
tmp16230 := MakeNative(func(__e *ControlFlow) {
tmp16231 := PrimCons(W4871, Nil)

tmp16232 := PrimCons(W4866, tmp16231)

tmp16233 := PrimCons(W4861, tmp16232)

tmp16234 := PrimCons(W4871, Nil)

tmp16235 := PrimCons(symlist, tmp16234)

tmp16236 := PrimCons(tmp16235, Nil)

tmp16237 := PrimCons(W4866, tmp16236)

tmp16238 := PrimCons(W4863, tmp16237)

tmp16239 := PrimCons(tmp16238, W4874)

tmp16240 := PrimCons(tmp16233, tmp16239)

__e.TailApply(PrimFunc(symshen_4l_1rules), tmp16240, V4845, True, V4847, V4848, W4851, V4850)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V4847, V4848, W4851, tmp16230)
return


}, 0)

__e.TailApply(PrimFunc(symwhen), tmp16228, V4847, V4848, W4851, tmp16229)
return


}, 1)

tmp16241 := PrimTail(W4856)

__e.TailApply(tmp16224, tmp16241)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16244 := PrimTail(W4867)

tmp16245 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16244, V4847)


__e.TailApply(tmp16223, tmp16245)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16248 := PrimTail(W4870)

tmp16249 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16248, V4847)


__e.TailApply(tmp16222, tmp16249)
return


}, 1)

tmp16250 := PrimHead(W4870)

__e.TailApply(tmp16221, tmp16250)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16253 := PrimTail(W4868)

tmp16254 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16253, V4847)


__e.TailApply(tmp16220, tmp16254)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16257 := PrimHead(W4868)

tmp16258 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16257, V4847)


__e.TailApply(tmp16219, tmp16258)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16261 := PrimHead(W4867)

tmp16262 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16261, V4847)


__e.TailApply(tmp16218, tmp16262)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16265 := PrimTail(W4865)

tmp16266 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16265, V4847)


__e.TailApply(tmp16217, tmp16266)
return


}, 1)

tmp16267 := PrimHead(W4865)

__e.TailApply(tmp16216, tmp16267)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16270 := PrimTail(W4857)

tmp16271 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16270, V4847)


__e.TailApply(tmp16215, tmp16271)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16274 := PrimTail(W4862)

tmp16275 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16274, V4847)


__e.TailApply(tmp16214, tmp16275)
return


}, 1)

tmp16276 := PrimHead(W4862)

__e.TailApply(tmp16213, tmp16276)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16279 := PrimTail(W4860)

tmp16280 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16279, V4847)


__e.TailApply(tmp16212, tmp16280)
return


}, 1)

tmp16281 := PrimHead(W4860)

__e.TailApply(tmp16211, tmp16281)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16284 := PrimTail(W4858)

tmp16285 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16284, V4847)


__e.TailApply(tmp16210, tmp16285)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16288 := PrimHead(W4858)

tmp16289 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16288, V4847)


__e.TailApply(tmp16209, tmp16289)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16292 := PrimHead(W4857)

tmp16293 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16292, V4847)


__e.TailApply(tmp16208, tmp16293)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16296 := PrimHead(W4856)

tmp16297 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16296, V4847)


__e.TailApply(tmp16207, tmp16297)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16300 := Call(__e, PrimFunc(symshen_4lazyderef), V4844, V4847)


tmp16301 := Call(__e, tmp16206, tmp16300)


ifres16205 = tmp16301


} else {
ifres16205 = False


}

__e.TailApply(tmp15875, ifres16205)
return


} else {
__e.Return(W4852)
return
}


}, 1)

tmp16317 := Call(__e, PrimFunc(symshen_4unlocked_2), V4848)


var ifres16305 Obj

if True == tmp16317 {
tmp16306 := MakeNative(func(__e *ControlFlow) {
W4853 := __e.Get(1)
_ = W4853
tmp16314 := PrimEqual(W4853, Nil)

if True == tmp16314 {
tmp16307 := MakeNative(func(__e *ControlFlow) {
W4854 := __e.Get(1)
_ = W4854
tmp16311 := PrimEqual(W4854, True)

if True == tmp16311 {
tmp16308 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16308

tmp16309 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symbind), V4845, Nil, V4847, V4848, W4851, V4850)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V4847, V4848, W4851, tmp16309)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16312 := Call(__e, PrimFunc(symshen_4lazyderef), V4846, V4847)


__e.TailApply(tmp16307, tmp16312)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16315 := Call(__e, PrimFunc(symshen_4lazyderef), V4844, V4847)


tmp16316 := Call(__e, tmp16306, tmp16315)


ifres16305 = tmp16316


} else {
ifres16305 = False


}

__e.TailApply(tmp15874, ifres16305)
return


}, 1)

tmp16318 := PrimNumberAdd(V4849, MakeNumber(1))

__e.TailApply(tmp15873, tmp16318)
return


}, 7)

tmp16319 := Call(__e, ns2_1set, symshen_4l_1rules, tmp15872)


_ = tmp16319

tmp16320 := MakeNative(func(__e *ControlFlow) {
V4945 := __e.Get(1)
_ = V4945
V4946 := __e.Get(2)
_ = V4946
V4947 := __e.Get(3)
_ = V4947
V4948 := __e.Get(4)
_ = V4948
V4949 := __e.Get(5)
_ = V4949
V4950 := __e.Get(6)
_ = V4950
tmp16321 := MakeNative(func(__e *ControlFlow) {
W4951 := __e.Get(1)
_ = W4951
tmp16322 := MakeNative(func(__e *ControlFlow) {
W4952 := __e.Get(1)
_ = W4952
tmp16324 := PrimEqual(W4952, False)

if True == tmp16324 {
__e.TailApply(PrimFunc(symshen_4unlock), V4948, W4951)
return
} else {
__e.Return(W4952)
return
}


}, 1)

tmp16372 := Call(__e, PrimFunc(symshen_4unlocked_2), V4948)


var ifres16325 Obj

if True == tmp16372 {
tmp16326 := MakeNative(func(__e *ControlFlow) {
W4953 := __e.Get(1)
_ = W4953
tmp16369 := PrimIsPair(W4953)

if True == tmp16369 {
tmp16327 := MakeNative(func(__e *ControlFlow) {
W4954 := __e.Get(1)
_ = W4954
tmp16365 := PrimEqual(W4954, symdefine)

if True == tmp16365 {
tmp16328 := MakeNative(func(__e *ControlFlow) {
W4955 := __e.Get(1)
_ = W4955
tmp16361 := PrimIsPair(W4955)

if True == tmp16361 {
tmp16329 := MakeNative(func(__e *ControlFlow) {
W4956 := __e.Get(1)
_ = W4956
tmp16330 := MakeNative(func(__e *ControlFlow) {
W4957 := __e.Get(1)
_ = W4957
tmp16331 := MakeNative(func(__e *ControlFlow) {
W4958 := __e.Get(1)
_ = W4958
tmp16332 := MakeNative(func(__e *ControlFlow) {
W4959 := __e.Get(1)
_ = W4959
tmp16333 := MakeNative(func(__e *ControlFlow) {
W4960 := __e.Get(1)
_ = W4960
tmp16334 := MakeNative(func(__e *ControlFlow) {
W4961 := __e.Get(1)
_ = W4961
tmp16335 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16335

tmp16336 := MakeNative(func(__e *ControlFlow) {
tmp16337 := PrimCons(W4956, W4957)

tmp16338 := Call(__e, PrimFunc(symshen_4sigxrules), tmp16337)


tmp16339 := MakeNative(func(__e *ControlFlow) {
tmp16340 := Call(__e, PrimFunc(symshen_4lazyderef), W4958, V4947)


tmp16341 := Call(__e, PrimFunc(symfst), tmp16340)


tmp16342 := MakeNative(func(__e *ControlFlow) {
tmp16343 := Call(__e, PrimFunc(symshen_4lazyderef), W4958, V4947)


tmp16344 := Call(__e, PrimFunc(symsnd), tmp16343)


tmp16345 := MakeNative(func(__e *ControlFlow) {
tmp16346 := Call(__e, PrimFunc(symshen_4deref), W4961, V4947)


tmp16347 := Call(__e, PrimFunc(symshen_4freshen_1sig), tmp16346)


tmp16348 := MakeNative(func(__e *ControlFlow) {
tmp16349 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symis), W4961, V4946, V4947, V4948, W4951, V4950)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4t_d_1rules), W4956, W4959, W4960, MakeNumber(1), V4947, V4948, W4951, tmp16349)
return


}, 0)

__e.TailApply(PrimFunc(symbind), W4960, tmp16347, V4947, V4948, W4951, tmp16348)
return


}, 0)

__e.TailApply(PrimFunc(symbind), W4959, tmp16344, V4947, V4948, W4951, tmp16345)
return


}, 0)

__e.TailApply(PrimFunc(symbind), W4961, tmp16341, V4947, V4948, W4951, tmp16342)
return


}, 0)

__e.TailApply(PrimFunc(symbind), W4958, tmp16338, V4947, V4948, W4951, tmp16339)
return


}, 0)

tmp16350 := Call(__e, PrimFunc(symshen_4cut), V4947, V4948, W4951, tmp16336)


__e.TailApply(PrimFunc(symshen_4gc), V4947, tmp16350)
return


}, 1)

tmp16351 := Call(__e, PrimFunc(symshen_4newpv), V4947)


tmp16352 := Call(__e, tmp16334, tmp16351)


__e.TailApply(PrimFunc(symshen_4gc), V4947, tmp16352)
return


}, 1)

tmp16353 := Call(__e, PrimFunc(symshen_4newpv), V4947)


tmp16354 := Call(__e, tmp16333, tmp16353)


__e.TailApply(PrimFunc(symshen_4gc), V4947, tmp16354)
return


}, 1)

tmp16355 := Call(__e, PrimFunc(symshen_4newpv), V4947)


tmp16356 := Call(__e, tmp16332, tmp16355)


__e.TailApply(PrimFunc(symshen_4gc), V4947, tmp16356)
return


}, 1)

tmp16357 := Call(__e, PrimFunc(symshen_4newpv), V4947)


__e.TailApply(tmp16331, tmp16357)
return


}, 1)

tmp16358 := PrimTail(W4955)

__e.TailApply(tmp16330, tmp16358)
return


}, 1)

tmp16359 := PrimHead(W4955)

__e.TailApply(tmp16329, tmp16359)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16362 := PrimTail(W4953)

tmp16363 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16362, V4947)


__e.TailApply(tmp16328, tmp16363)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16366 := PrimHead(W4953)

tmp16367 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16366, V4947)


__e.TailApply(tmp16327, tmp16367)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16370 := Call(__e, PrimFunc(symshen_4lazyderef), V4945, V4947)


tmp16371 := Call(__e, tmp16326, tmp16370)


ifres16325 = tmp16371


} else {
ifres16325 = False


}

__e.TailApply(tmp16322, ifres16325)
return


}, 1)

tmp16373 := PrimNumberAdd(V4949, MakeNumber(1))

__e.TailApply(tmp16321, tmp16373)
return


}, 6)

tmp16374 := Call(__e, ns2_1set, symshen_4t_d, tmp16320)


_ = tmp16374

tmp16375 := MakeNative(func(__e *ControlFlow) {
V4962 := __e.Get(1)
_ = V4962
tmp16376 := MakeNative(func(__e *ControlFlow) {
Z4963 := __e.Get(1)
_ = Z4963
__e.TailApply(PrimFunc(symshen_4_5sig_drules_6), Z4963)
return
}, 1)

__e.TailApply(PrimFunc(symcompile), tmp16376, V4962)
return


}, 1)

tmp16377 := Call(__e, ns2_1set, symshen_4sigxrules, tmp16375)


_ = tmp16377

tmp16378 := MakeNative(func(__e *ControlFlow) {
V4964 := __e.Get(1)
_ = V4964
tmp16379 := MakeNative(func(__e *ControlFlow) {
W4965 := __e.Get(1)
_ = W4965
tmp16381 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W4965)


if True == tmp16381 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W4965)
return
}


}, 1)

tmp16414 := PrimIsPair(V4964)

var ifres16382 Obj

if True == tmp16414 {
tmp16383 := MakeNative(func(__e *ControlFlow) {
W4966 := __e.Get(1)
_ = W4966
tmp16410 := Call(__e, PrimFunc(symshen_4hds_a_2), W4966, sym_i)


if True == tmp16410 {
tmp16384 := MakeNative(func(__e *ControlFlow) {
W4967 := __e.Get(1)
_ = W4967
tmp16385 := MakeNative(func(__e *ControlFlow) {
W4968 := __e.Get(1)
_ = W4968
tmp16406 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W4968)


if True == tmp16406 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp16386 := MakeNative(func(__e *ControlFlow) {
W4969 := __e.Get(1)
_ = W4969
tmp16387 := MakeNative(func(__e *ControlFlow) {
W4970 := __e.Get(1)
_ = W4970
tmp16402 := Call(__e, PrimFunc(symshen_4hds_a_2), W4970, sym_j)


if True == tmp16402 {
tmp16388 := MakeNative(func(__e *ControlFlow) {
W4971 := __e.Get(1)
_ = W4971
tmp16389 := MakeNative(func(__e *ControlFlow) {
W4972 := __e.Get(1)
_ = W4972
tmp16398 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W4972)


if True == tmp16398 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp16390 := MakeNative(func(__e *ControlFlow) {
W4973 := __e.Get(1)
_ = W4973
tmp16391 := MakeNative(func(__e *ControlFlow) {
W4974 := __e.Get(1)
_ = W4974
tmp16392 := MakeNative(func(__e *ControlFlow) {
W4975 := __e.Get(1)
_ = W4975
__e.TailApply(PrimFunc(sym_8p), W4975, W4973)
return
}, 1)

tmp16393 := Call(__e, PrimFunc(symshen_4rectify_1type), W4969)


tmp16394 := Call(__e, tmp16392, tmp16393)


__e.TailApply(PrimFunc(symshen_4comb), W4974, tmp16394)
return


}, 1)

tmp16395 := Call(__e, PrimFunc(symshen_4in_1_6), W4972)


__e.TailApply(tmp16391, tmp16395)
return


}, 1)

tmp16396 := Call(__e, PrimFunc(symshen_4_5_1out), W4972)


__e.TailApply(tmp16390, tmp16396)
return


}


}, 1)

tmp16399 := Call(__e, PrimFunc(symshen_4_5rules_d_6), W4971)


__e.TailApply(tmp16389, tmp16399)
return


}, 1)

tmp16400 := Call(__e, PrimFunc(symtail), W4970)


__e.TailApply(tmp16388, tmp16400)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp16403 := Call(__e, PrimFunc(symshen_4in_1_6), W4968)


__e.TailApply(tmp16387, tmp16403)
return


}, 1)

tmp16404 := Call(__e, PrimFunc(symshen_4_5_1out), W4968)


__e.TailApply(tmp16386, tmp16404)
return


}


}, 1)

tmp16407 := Call(__e, PrimFunc(symshen_4_5signature_6), W4967)


__e.TailApply(tmp16385, tmp16407)
return


}, 1)

tmp16408 := Call(__e, PrimFunc(symtail), W4966)


__e.TailApply(tmp16384, tmp16408)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp16411 := Call(__e, PrimFunc(symtail), V4964)


tmp16412 := Call(__e, tmp16383, tmp16411)


ifres16382 = tmp16412


} else {
tmp16413 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16382 = tmp16413


}

__e.TailApply(tmp16379, ifres16382)
return


}, 1)

tmp16415 := Call(__e, ns2_1set, symshen_4_5sig_drules_6, tmp16378)


_ = tmp16415

tmp16416 := MakeNative(func(__e *ControlFlow) {
V4976 := __e.Get(1)
_ = V4976
tmp16417 := MakeNative(func(__e *ControlFlow) {
W4977 := __e.Get(1)
_ = W4977
tmp16418 := MakeNative(func(__e *ControlFlow) {
W4978 := __e.Get(1)
_ = W4978
__e.TailApply(PrimFunc(symshen_4freshen_1type), W4978, V4976)
return
}, 1)

tmp16419 := MakeNative(func(__e *ControlFlow) {
Z4979 := __e.Get(1)
_ = Z4979
tmp16420 := Call(__e, PrimFunc(symconcat), sym_e, Z4979)


tmp16421 := Call(__e, PrimFunc(symshen_4freshterm), tmp16420)


__e.Return(PrimCons(Z4979, tmp16421))
return


}, 1)

tmp16422 := Call(__e, PrimFunc(symmap), tmp16419, W4977)


__e.TailApply(tmp16418, tmp16422)
return


}, 1)

tmp16423 := Call(__e, PrimFunc(symshen_4extract_1vars), V4976)


__e.TailApply(tmp16417, tmp16423)
return


}, 1)

tmp16424 := Call(__e, ns2_1set, symshen_4freshen_1sig, tmp16416)


_ = tmp16424

tmp16425 := MakeNative(func(__e *ControlFlow) {
V4980 := __e.Get(1)
_ = V4980
V4981 := __e.Get(2)
_ = V4981
tmp16439 := PrimEqual(Nil, V4980)

if True == tmp16439 {
__e.Return(V4981)
return
} else {
tmp16437 := PrimIsPair(V4980)

var ifres16433 Obj

if True == tmp16437 {
tmp16435 := PrimHead(V4980)

tmp16436 := PrimIsPair(tmp16435)

var ifres16434 Obj

if True == tmp16436 {
ifres16434 = True


} else {
ifres16434 = False


}

ifres16433 = ifres16434


} else {
ifres16433 = False


}

if True == ifres16433 {
tmp16426 := PrimTail(V4980)

tmp16427 := PrimHead(V4980)

tmp16428 := PrimTail(tmp16427)

tmp16429 := PrimHead(V4980)

tmp16430 := PrimHead(tmp16429)

tmp16431 := Call(__e, PrimFunc(symsubst), tmp16428, tmp16430, V4981)


__e.TailApply(PrimFunc(symshen_4freshen_1type), tmp16426, tmp16431)
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.freshen-type")))
return
}


}


}, 2)

tmp16440 := Call(__e, ns2_1set, symshen_4freshen_1type, tmp16425)


_ = tmp16440

tmp16441 := MakeNative(func(__e *ControlFlow) {
V4982 := __e.Get(1)
_ = V4982
tmp16442 := MakeNative(func(__e *ControlFlow) {
W4983 := __e.Get(1)
_ = W4983
tmp16457 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W4983)


if True == tmp16457 {
tmp16443 := MakeNative(func(__e *ControlFlow) {
W4990 := __e.Get(1)
_ = W4990
tmp16445 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W4990)


if True == tmp16445 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W4990)
return
}


}, 1)

tmp16446 := MakeNative(func(__e *ControlFlow) {
W4991 := __e.Get(1)
_ = W4991
tmp16453 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W4991)


if True == tmp16453 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp16447 := MakeNative(func(__e *ControlFlow) {
W4992 := __e.Get(1)
_ = W4992
tmp16448 := MakeNative(func(__e *ControlFlow) {
W4993 := __e.Get(1)
_ = W4993
tmp16449 := PrimCons(W4992, Nil)

__e.TailApply(PrimFunc(symshen_4comb), W4993, tmp16449)
return


}, 1)

tmp16450 := Call(__e, PrimFunc(symshen_4in_1_6), W4991)


__e.TailApply(tmp16448, tmp16450)
return


}, 1)

tmp16451 := Call(__e, PrimFunc(symshen_4_5_1out), W4991)


__e.TailApply(tmp16447, tmp16451)
return


}


}, 1)

tmp16454 := Call(__e, PrimFunc(symshen_4_5rule_d_6), V4982)


tmp16455 := Call(__e, tmp16446, tmp16454)


__e.TailApply(tmp16443, tmp16455)
return


} else {
__e.Return(W4983)
return
}


}, 1)

tmp16458 := MakeNative(func(__e *ControlFlow) {
W4984 := __e.Get(1)
_ = W4984
tmp16473 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W4984)


if True == tmp16473 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp16459 := MakeNative(func(__e *ControlFlow) {
W4985 := __e.Get(1)
_ = W4985
tmp16460 := MakeNative(func(__e *ControlFlow) {
W4986 := __e.Get(1)
_ = W4986
tmp16461 := MakeNative(func(__e *ControlFlow) {
W4987 := __e.Get(1)
_ = W4987
tmp16468 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W4987)


if True == tmp16468 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp16462 := MakeNative(func(__e *ControlFlow) {
W4988 := __e.Get(1)
_ = W4988
tmp16463 := MakeNative(func(__e *ControlFlow) {
W4989 := __e.Get(1)
_ = W4989
tmp16464 := PrimCons(W4985, W4988)

__e.TailApply(PrimFunc(symshen_4comb), W4989, tmp16464)
return


}, 1)

tmp16465 := Call(__e, PrimFunc(symshen_4in_1_6), W4987)


__e.TailApply(tmp16463, tmp16465)
return


}, 1)

tmp16466 := Call(__e, PrimFunc(symshen_4_5_1out), W4987)


__e.TailApply(tmp16462, tmp16466)
return


}


}, 1)

tmp16469 := Call(__e, PrimFunc(symshen_4_5rules_d_6), W4986)


__e.TailApply(tmp16461, tmp16469)
return


}, 1)

tmp16470 := Call(__e, PrimFunc(symshen_4in_1_6), W4984)


__e.TailApply(tmp16460, tmp16470)
return


}, 1)

tmp16471 := Call(__e, PrimFunc(symshen_4_5_1out), W4984)


__e.TailApply(tmp16459, tmp16471)
return


}


}, 1)

tmp16474 := Call(__e, PrimFunc(symshen_4_5rule_d_6), V4982)


tmp16475 := Call(__e, tmp16458, tmp16474)


__e.TailApply(tmp16442, tmp16475)
return


}, 1)

tmp16476 := Call(__e, ns2_1set, symshen_4_5rules_d_6, tmp16441)


_ = tmp16476

tmp16477 := MakeNative(func(__e *ControlFlow) {
V4994 := __e.Get(1)
_ = V4994
tmp16478 := MakeNative(func(__e *ControlFlow) {
W4995 := __e.Get(1)
_ = W4995
tmp16564 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W4995)


if True == tmp16564 {
tmp16479 := MakeNative(func(__e *ControlFlow) {
W5005 := __e.Get(1)
_ = W5005
tmp16528 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W5005)


if True == tmp16528 {
tmp16480 := MakeNative(func(__e *ControlFlow) {
W5015 := __e.Get(1)
_ = W5015
tmp16505 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W5015)


if True == tmp16505 {
tmp16481 := MakeNative(func(__e *ControlFlow) {
W5022 := __e.Get(1)
_ = W5022
tmp16483 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W5022)


if True == tmp16483 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W5022)
return
}


}, 1)

tmp16484 := MakeNative(func(__e *ControlFlow) {
W5023 := __e.Get(1)
_ = W5023
tmp16501 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W5023)


if True == tmp16501 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp16485 := MakeNative(func(__e *ControlFlow) {
W5024 := __e.Get(1)
_ = W5024
tmp16486 := MakeNative(func(__e *ControlFlow) {
W5025 := __e.Get(1)
_ = W5025
tmp16497 := Call(__e, PrimFunc(symshen_4hds_a_2), W5025, sym_1_6)


if True == tmp16497 {
tmp16487 := MakeNative(func(__e *ControlFlow) {
W5026 := __e.Get(1)
_ = W5026
tmp16494 := PrimIsPair(W5026)

if True == tmp16494 {
tmp16488 := MakeNative(func(__e *ControlFlow) {
W5027 := __e.Get(1)
_ = W5027
tmp16489 := MakeNative(func(__e *ControlFlow) {
W5028 := __e.Get(1)
_ = W5028
tmp16490 := Call(__e, PrimFunc(sym_8p), W5024, W5027)


__e.TailApply(PrimFunc(symshen_4comb), W5028, tmp16490)
return


}, 1)

tmp16491 := Call(__e, PrimFunc(symtail), W5026)


__e.TailApply(tmp16489, tmp16491)
return


}, 1)

tmp16492 := Call(__e, PrimFunc(symhead), W5026)


__e.TailApply(tmp16488, tmp16492)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp16495 := Call(__e, PrimFunc(symtail), W5025)


__e.TailApply(tmp16487, tmp16495)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp16498 := Call(__e, PrimFunc(symshen_4in_1_6), W5023)


__e.TailApply(tmp16486, tmp16498)
return


}, 1)

tmp16499 := Call(__e, PrimFunc(symshen_4_5_1out), W5023)


__e.TailApply(tmp16485, tmp16499)
return


}


}, 1)

tmp16502 := Call(__e, PrimFunc(symshen_4_5patterns_6), V4994)


tmp16503 := Call(__e, tmp16484, tmp16502)


__e.TailApply(tmp16481, tmp16503)
return


} else {
__e.Return(W5015)
return
}


}, 1)

tmp16506 := MakeNative(func(__e *ControlFlow) {
W5016 := __e.Get(1)
_ = W5016
tmp16524 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W5016)


if True == tmp16524 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp16507 := MakeNative(func(__e *ControlFlow) {
W5017 := __e.Get(1)
_ = W5017
tmp16508 := MakeNative(func(__e *ControlFlow) {
W5018 := __e.Get(1)
_ = W5018
tmp16520 := Call(__e, PrimFunc(symshen_4hds_a_2), W5018, sym_5_1)


if True == tmp16520 {
tmp16509 := MakeNative(func(__e *ControlFlow) {
W5019 := __e.Get(1)
_ = W5019
tmp16517 := PrimIsPair(W5019)

if True == tmp16517 {
tmp16510 := MakeNative(func(__e *ControlFlow) {
W5020 := __e.Get(1)
_ = W5020
tmp16511 := MakeNative(func(__e *ControlFlow) {
W5021 := __e.Get(1)
_ = W5021
tmp16512 := Call(__e, PrimFunc(symshen_4correct), W5020)


tmp16513 := Call(__e, PrimFunc(sym_8p), W5017, tmp16512)


__e.TailApply(PrimFunc(symshen_4comb), W5021, tmp16513)
return


}, 1)

tmp16514 := Call(__e, PrimFunc(symtail), W5019)


__e.TailApply(tmp16511, tmp16514)
return


}, 1)

tmp16515 := Call(__e, PrimFunc(symhead), W5019)


__e.TailApply(tmp16510, tmp16515)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp16518 := Call(__e, PrimFunc(symtail), W5018)


__e.TailApply(tmp16509, tmp16518)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp16521 := Call(__e, PrimFunc(symshen_4in_1_6), W5016)


__e.TailApply(tmp16508, tmp16521)
return


}, 1)

tmp16522 := Call(__e, PrimFunc(symshen_4_5_1out), W5016)


__e.TailApply(tmp16507, tmp16522)
return


}


}, 1)

tmp16525 := Call(__e, PrimFunc(symshen_4_5patterns_6), V4994)


tmp16526 := Call(__e, tmp16506, tmp16525)


__e.TailApply(tmp16480, tmp16526)
return


} else {
__e.Return(W5005)
return
}


}, 1)

tmp16529 := MakeNative(func(__e *ControlFlow) {
W5006 := __e.Get(1)
_ = W5006
tmp16560 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W5006)


if True == tmp16560 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp16530 := MakeNative(func(__e *ControlFlow) {
W5007 := __e.Get(1)
_ = W5007
tmp16531 := MakeNative(func(__e *ControlFlow) {
W5008 := __e.Get(1)
_ = W5008
tmp16556 := Call(__e, PrimFunc(symshen_4hds_a_2), W5008, sym_5_1)


if True == tmp16556 {
tmp16532 := MakeNative(func(__e *ControlFlow) {
W5009 := __e.Get(1)
_ = W5009
tmp16553 := PrimIsPair(W5009)

if True == tmp16553 {
tmp16533 := MakeNative(func(__e *ControlFlow) {
W5010 := __e.Get(1)
_ = W5010
tmp16534 := MakeNative(func(__e *ControlFlow) {
W5011 := __e.Get(1)
_ = W5011
tmp16549 := Call(__e, PrimFunc(symshen_4hds_a_2), W5011, symwhere)


if True == tmp16549 {
tmp16535 := MakeNative(func(__e *ControlFlow) {
W5012 := __e.Get(1)
_ = W5012
tmp16546 := PrimIsPair(W5012)

if True == tmp16546 {
tmp16536 := MakeNative(func(__e *ControlFlow) {
W5013 := __e.Get(1)
_ = W5013
tmp16537 := MakeNative(func(__e *ControlFlow) {
W5014 := __e.Get(1)
_ = W5014
tmp16538 := PrimCons(W5010, Nil)

tmp16539 := PrimCons(W5013, tmp16538)

tmp16540 := PrimCons(symwhere, tmp16539)

tmp16541 := Call(__e, PrimFunc(symshen_4correct), tmp16540)


tmp16542 := Call(__e, PrimFunc(sym_8p), W5007, tmp16541)


__e.TailApply(PrimFunc(symshen_4comb), W5014, tmp16542)
return


}, 1)

tmp16543 := Call(__e, PrimFunc(symtail), W5012)


__e.TailApply(tmp16537, tmp16543)
return


}, 1)

tmp16544 := Call(__e, PrimFunc(symhead), W5012)


__e.TailApply(tmp16536, tmp16544)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp16547 := Call(__e, PrimFunc(symtail), W5011)


__e.TailApply(tmp16535, tmp16547)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp16550 := Call(__e, PrimFunc(symtail), W5009)


__e.TailApply(tmp16534, tmp16550)
return


}, 1)

tmp16551 := Call(__e, PrimFunc(symhead), W5009)


__e.TailApply(tmp16533, tmp16551)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp16554 := Call(__e, PrimFunc(symtail), W5008)


__e.TailApply(tmp16532, tmp16554)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp16557 := Call(__e, PrimFunc(symshen_4in_1_6), W5006)


__e.TailApply(tmp16531, tmp16557)
return


}, 1)

tmp16558 := Call(__e, PrimFunc(symshen_4_5_1out), W5006)


__e.TailApply(tmp16530, tmp16558)
return


}


}, 1)

tmp16561 := Call(__e, PrimFunc(symshen_4_5patterns_6), V4994)


tmp16562 := Call(__e, tmp16529, tmp16561)


__e.TailApply(tmp16479, tmp16562)
return


} else {
__e.Return(W4995)
return
}


}, 1)

tmp16565 := MakeNative(func(__e *ControlFlow) {
W4996 := __e.Get(1)
_ = W4996
tmp16595 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W4996)


if True == tmp16595 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp16566 := MakeNative(func(__e *ControlFlow) {
W4997 := __e.Get(1)
_ = W4997
tmp16567 := MakeNative(func(__e *ControlFlow) {
W4998 := __e.Get(1)
_ = W4998
tmp16591 := Call(__e, PrimFunc(symshen_4hds_a_2), W4998, sym_1_6)


if True == tmp16591 {
tmp16568 := MakeNative(func(__e *ControlFlow) {
W4999 := __e.Get(1)
_ = W4999
tmp16588 := PrimIsPair(W4999)

if True == tmp16588 {
tmp16569 := MakeNative(func(__e *ControlFlow) {
W5000 := __e.Get(1)
_ = W5000
tmp16570 := MakeNative(func(__e *ControlFlow) {
W5001 := __e.Get(1)
_ = W5001
tmp16584 := Call(__e, PrimFunc(symshen_4hds_a_2), W5001, symwhere)


if True == tmp16584 {
tmp16571 := MakeNative(func(__e *ControlFlow) {
W5002 := __e.Get(1)
_ = W5002
tmp16581 := PrimIsPair(W5002)

if True == tmp16581 {
tmp16572 := MakeNative(func(__e *ControlFlow) {
W5003 := __e.Get(1)
_ = W5003
tmp16573 := MakeNative(func(__e *ControlFlow) {
W5004 := __e.Get(1)
_ = W5004
tmp16574 := PrimCons(W5000, Nil)

tmp16575 := PrimCons(W5003, tmp16574)

tmp16576 := PrimCons(symwhere, tmp16575)

tmp16577 := Call(__e, PrimFunc(sym_8p), W4997, tmp16576)


__e.TailApply(PrimFunc(symshen_4comb), W5004, tmp16577)
return


}, 1)

tmp16578 := Call(__e, PrimFunc(symtail), W5002)


__e.TailApply(tmp16573, tmp16578)
return


}, 1)

tmp16579 := Call(__e, PrimFunc(symhead), W5002)


__e.TailApply(tmp16572, tmp16579)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp16582 := Call(__e, PrimFunc(symtail), W5001)


__e.TailApply(tmp16571, tmp16582)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp16585 := Call(__e, PrimFunc(symtail), W4999)


__e.TailApply(tmp16570, tmp16585)
return


}, 1)

tmp16586 := Call(__e, PrimFunc(symhead), W4999)


__e.TailApply(tmp16569, tmp16586)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp16589 := Call(__e, PrimFunc(symtail), W4998)


__e.TailApply(tmp16568, tmp16589)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp16592 := Call(__e, PrimFunc(symshen_4in_1_6), W4996)


__e.TailApply(tmp16567, tmp16592)
return


}, 1)

tmp16593 := Call(__e, PrimFunc(symshen_4_5_1out), W4996)


__e.TailApply(tmp16566, tmp16593)
return


}


}, 1)

tmp16596 := Call(__e, PrimFunc(symshen_4_5patterns_6), V4994)


tmp16597 := Call(__e, tmp16565, tmp16596)


__e.TailApply(tmp16478, tmp16597)
return


}, 1)

tmp16598 := Call(__e, ns2_1set, symshen_4_5rule_d_6, tmp16477)


_ = tmp16598

tmp16599 := MakeNative(func(__e *ControlFlow) {
V5029 := __e.Get(1)
_ = V5029
tmp16747 := PrimIsPair(V5029)

var ifres16691 Obj

if True == tmp16747 {
tmp16745 := PrimHead(V5029)

tmp16746 := PrimEqual(symwhere, tmp16745)

var ifres16693 Obj

if True == tmp16746 {
tmp16743 := PrimTail(V5029)

tmp16744 := PrimIsPair(tmp16743)

var ifres16695 Obj

if True == tmp16744 {
tmp16740 := PrimTail(V5029)

tmp16741 := PrimTail(tmp16740)

tmp16742 := PrimIsPair(tmp16741)

var ifres16697 Obj

if True == tmp16742 {
tmp16736 := PrimTail(V5029)

tmp16737 := PrimTail(tmp16736)

tmp16738 := PrimHead(tmp16737)

tmp16739 := PrimIsPair(tmp16738)

var ifres16699 Obj

if True == tmp16739 {
tmp16731 := PrimTail(V5029)

tmp16732 := PrimTail(tmp16731)

tmp16733 := PrimHead(tmp16732)

tmp16734 := PrimHead(tmp16733)

tmp16735 := PrimEqual(symfail_1if, tmp16734)

var ifres16701 Obj

if True == tmp16735 {
tmp16726 := PrimTail(V5029)

tmp16727 := PrimTail(tmp16726)

tmp16728 := PrimHead(tmp16727)

tmp16729 := PrimTail(tmp16728)

tmp16730 := PrimIsPair(tmp16729)

var ifres16703 Obj

if True == tmp16730 {
tmp16720 := PrimTail(V5029)

tmp16721 := PrimTail(tmp16720)

tmp16722 := PrimHead(tmp16721)

tmp16723 := PrimTail(tmp16722)

tmp16724 := PrimTail(tmp16723)

tmp16725 := PrimIsPair(tmp16724)

var ifres16705 Obj

if True == tmp16725 {
tmp16713 := PrimTail(V5029)

tmp16714 := PrimTail(tmp16713)

tmp16715 := PrimHead(tmp16714)

tmp16716 := PrimTail(tmp16715)

tmp16717 := PrimTail(tmp16716)

tmp16718 := PrimTail(tmp16717)

tmp16719 := PrimEqual(Nil, tmp16718)

var ifres16707 Obj

if True == tmp16719 {
tmp16709 := PrimTail(V5029)

tmp16710 := PrimTail(tmp16709)

tmp16711 := PrimTail(tmp16710)

tmp16712 := PrimEqual(Nil, tmp16711)

var ifres16708 Obj

if True == tmp16712 {
ifres16708 = True


} else {
ifres16708 = False


}

ifres16707 = ifres16708


} else {
ifres16707 = False


}

var ifres16706 Obj

if True == ifres16707 {
ifres16706 = True


} else {
ifres16706 = False


}

ifres16705 = ifres16706


} else {
ifres16705 = False


}

var ifres16704 Obj

if True == ifres16705 {
ifres16704 = True


} else {
ifres16704 = False


}

ifres16703 = ifres16704


} else {
ifres16703 = False


}

var ifres16702 Obj

if True == ifres16703 {
ifres16702 = True


} else {
ifres16702 = False


}

ifres16701 = ifres16702


} else {
ifres16701 = False


}

var ifres16700 Obj

if True == ifres16701 {
ifres16700 = True


} else {
ifres16700 = False


}

ifres16699 = ifres16700


} else {
ifres16699 = False


}

var ifres16698 Obj

if True == ifres16699 {
ifres16698 = True


} else {
ifres16698 = False


}

ifres16697 = ifres16698


} else {
ifres16697 = False


}

var ifres16696 Obj

if True == ifres16697 {
ifres16696 = True


} else {
ifres16696 = False


}

ifres16695 = ifres16696


} else {
ifres16695 = False


}

var ifres16694 Obj

if True == ifres16695 {
ifres16694 = True


} else {
ifres16694 = False


}

ifres16693 = ifres16694


} else {
ifres16693 = False


}

var ifres16692 Obj

if True == ifres16693 {
ifres16692 = True


} else {
ifres16692 = False


}

ifres16691 = ifres16692


} else {
ifres16691 = False


}

if True == ifres16691 {
tmp16600 := PrimTail(V5029)

tmp16601 := PrimHead(tmp16600)

tmp16602 := PrimTail(V5029)

tmp16603 := PrimTail(tmp16602)

tmp16604 := PrimHead(tmp16603)

tmp16605 := PrimTail(tmp16604)

tmp16606 := PrimCons(tmp16605, Nil)

tmp16607 := PrimCons(symnot, tmp16606)

tmp16608 := PrimCons(tmp16607, Nil)

tmp16609 := PrimCons(tmp16601, tmp16608)

tmp16610 := PrimCons(symand, tmp16609)

tmp16611 := PrimTail(V5029)

tmp16612 := PrimTail(tmp16611)

tmp16613 := PrimHead(tmp16612)

tmp16614 := PrimTail(tmp16613)

tmp16615 := PrimTail(tmp16614)

tmp16616 := PrimCons(tmp16610, tmp16615)

__e.Return(PrimCons(symwhere, tmp16616))
return


} else {
tmp16689 := PrimIsPair(V5029)

var ifres16670 Obj

if True == tmp16689 {
tmp16687 := PrimHead(V5029)

tmp16688 := PrimEqual(symwhere, tmp16687)

var ifres16672 Obj

if True == tmp16688 {
tmp16685 := PrimTail(V5029)

tmp16686 := PrimIsPair(tmp16685)

var ifres16674 Obj

if True == tmp16686 {
tmp16682 := PrimTail(V5029)

tmp16683 := PrimTail(tmp16682)

tmp16684 := PrimIsPair(tmp16683)

var ifres16676 Obj

if True == tmp16684 {
tmp16678 := PrimTail(V5029)

tmp16679 := PrimTail(tmp16678)

tmp16680 := PrimTail(tmp16679)

tmp16681 := PrimEqual(Nil, tmp16680)

var ifres16677 Obj

if True == tmp16681 {
ifres16677 = True


} else {
ifres16677 = False


}

ifres16676 = ifres16677


} else {
ifres16676 = False


}

var ifres16675 Obj

if True == ifres16676 {
ifres16675 = True


} else {
ifres16675 = False


}

ifres16674 = ifres16675


} else {
ifres16674 = False


}

var ifres16673 Obj

if True == ifres16674 {
ifres16673 = True


} else {
ifres16673 = False


}

ifres16672 = ifres16673


} else {
ifres16672 = False


}

var ifres16671 Obj

if True == ifres16672 {
ifres16671 = True


} else {
ifres16671 = False


}

ifres16670 = ifres16671


} else {
ifres16670 = False


}

if True == ifres16670 {
tmp16617 := PrimTail(V5029)

tmp16618 := PrimHead(tmp16617)

tmp16619 := PrimTail(V5029)

tmp16620 := PrimTail(tmp16619)

tmp16621 := PrimHead(tmp16620)

tmp16622 := PrimCons(symfail, Nil)

tmp16623 := PrimCons(tmp16622, Nil)

tmp16624 := PrimCons(tmp16621, tmp16623)

tmp16625 := PrimCons(sym_a_a, tmp16624)

tmp16626 := PrimCons(tmp16625, Nil)

tmp16627 := PrimCons(symnot, tmp16626)

tmp16628 := PrimCons(tmp16627, Nil)

tmp16629 := PrimCons(tmp16618, tmp16628)

tmp16630 := PrimCons(symand, tmp16629)

tmp16631 := PrimTail(V5029)

tmp16632 := PrimTail(tmp16631)

tmp16633 := PrimCons(tmp16630, tmp16632)

__e.Return(PrimCons(symwhere, tmp16633))
return


} else {
tmp16668 := PrimIsPair(V5029)

var ifres16649 Obj

if True == tmp16668 {
tmp16666 := PrimHead(V5029)

tmp16667 := PrimEqual(symfail_1if, tmp16666)

var ifres16651 Obj

if True == tmp16667 {
tmp16664 := PrimTail(V5029)

tmp16665 := PrimIsPair(tmp16664)

var ifres16653 Obj

if True == tmp16665 {
tmp16661 := PrimTail(V5029)

tmp16662 := PrimTail(tmp16661)

tmp16663 := PrimIsPair(tmp16662)

var ifres16655 Obj

if True == tmp16663 {
tmp16657 := PrimTail(V5029)

tmp16658 := PrimTail(tmp16657)

tmp16659 := PrimTail(tmp16658)

tmp16660 := PrimEqual(Nil, tmp16659)

var ifres16656 Obj

if True == tmp16660 {
ifres16656 = True


} else {
ifres16656 = False


}

ifres16655 = ifres16656


} else {
ifres16655 = False


}

var ifres16654 Obj

if True == ifres16655 {
ifres16654 = True


} else {
ifres16654 = False


}

ifres16653 = ifres16654


} else {
ifres16653 = False


}

var ifres16652 Obj

if True == ifres16653 {
ifres16652 = True


} else {
ifres16652 = False


}

ifres16651 = ifres16652


} else {
ifres16651 = False


}

var ifres16650 Obj

if True == ifres16651 {
ifres16650 = True


} else {
ifres16650 = False


}

ifres16649 = ifres16650


} else {
ifres16649 = False


}

if True == ifres16649 {
tmp16634 := PrimTail(V5029)

tmp16635 := PrimCons(tmp16634, Nil)

tmp16636 := PrimCons(symnot, tmp16635)

tmp16637 := PrimTail(V5029)

tmp16638 := PrimTail(tmp16637)

tmp16639 := PrimCons(tmp16636, tmp16638)

__e.Return(PrimCons(symwhere, tmp16639))
return


} else {
tmp16640 := PrimCons(symfail, Nil)

tmp16641 := PrimCons(tmp16640, Nil)

tmp16642 := PrimCons(V5029, tmp16641)

tmp16643 := PrimCons(sym_a_a, tmp16642)

tmp16644 := PrimCons(tmp16643, Nil)

tmp16645 := PrimCons(symnot, tmp16644)

tmp16646 := PrimCons(V5029, Nil)

tmp16647 := PrimCons(tmp16645, tmp16646)

__e.Return(PrimCons(symwhere, tmp16647))
return


}


}


}


}, 1)

tmp16748 := Call(__e, ns2_1set, symshen_4correct, tmp16599)


_ = tmp16748

tmp16749 := MakeNative(func(__e *ControlFlow) {
V5030 := __e.Get(1)
_ = V5030
V5031 := __e.Get(2)
_ = V5031
V5032 := __e.Get(3)
_ = V5032
V5033 := __e.Get(4)
_ = V5033
V5034 := __e.Get(5)
_ = V5034
V5035 := __e.Get(6)
_ = V5035
V5036 := __e.Get(7)
_ = V5036
V5037 := __e.Get(8)
_ = V5037
tmp16750 := MakeNative(func(__e *ControlFlow) {
W5038 := __e.Get(1)
_ = W5038
tmp16751 := MakeNative(func(__e *ControlFlow) {
W5039 := __e.Get(1)
_ = W5039
tmp16781 := PrimEqual(W5039, False)

if True == tmp16781 {
tmp16752 := MakeNative(func(__e *ControlFlow) {
W5041 := __e.Get(1)
_ = W5041
tmp16754 := PrimEqual(W5041, False)

if True == tmp16754 {
__e.TailApply(PrimFunc(symshen_4unlock), V5035, W5038)
return
} else {
__e.Return(W5041)
return
}


}, 1)

tmp16779 := Call(__e, PrimFunc(symshen_4unlocked_2), V5035)


var ifres16755 Obj

if True == tmp16779 {
tmp16756 := MakeNative(func(__e *ControlFlow) {
W5042 := __e.Get(1)
_ = W5042
tmp16776 := PrimIsPair(W5042)

if True == tmp16776 {
tmp16757 := MakeNative(func(__e *ControlFlow) {
W5043 := __e.Get(1)
_ = W5043
tmp16758 := MakeNative(func(__e *ControlFlow) {
W5044 := __e.Get(1)
_ = W5044
tmp16759 := MakeNative(func(__e *ControlFlow) {
W5045 := __e.Get(1)
_ = W5045
tmp16760 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16760

tmp16761 := Call(__e, PrimFunc(symshen_4deref), W5043, V5034)


tmp16762 := Call(__e, PrimFunc(symshen_4freshen_1rule), tmp16761)


tmp16763 := MakeNative(func(__e *ControlFlow) {
tmp16764 := Call(__e, PrimFunc(symshen_4lazyderef), W5045, V5034)


tmp16765 := Call(__e, PrimFunc(symfst), tmp16764)


tmp16766 := Call(__e, PrimFunc(symshen_4lazyderef), W5045, V5034)


tmp16767 := Call(__e, PrimFunc(symsnd), tmp16766)


tmp16768 := MakeNative(func(__e *ControlFlow) {
tmp16769 := MakeNative(func(__e *ControlFlow) {
tmp16770 := PrimNumberAdd(V5033, MakeNumber(1))

__e.TailApply(PrimFunc(symshen_4t_d_1rules), V5030, W5044, V5032, tmp16770, V5034, V5035, W5038, V5037)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V5034, V5035, W5038, tmp16769)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4t_d_1rule), V5030, V5033, tmp16765, tmp16767, V5032, V5034, V5035, W5038, tmp16768)
return


}, 0)

tmp16771 := Call(__e, PrimFunc(symbind), W5045, tmp16762, V5034, V5035, W5038, tmp16763)


__e.TailApply(PrimFunc(symshen_4gc), V5034, tmp16771)
return


}, 1)

tmp16772 := Call(__e, PrimFunc(symshen_4newpv), V5034)


__e.TailApply(tmp16759, tmp16772)
return


}, 1)

tmp16773 := PrimTail(W5042)

__e.TailApply(tmp16758, tmp16773)
return


}, 1)

tmp16774 := PrimHead(W5042)

__e.TailApply(tmp16757, tmp16774)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16777 := Call(__e, PrimFunc(symshen_4lazyderef), V5031, V5034)


tmp16778 := Call(__e, tmp16756, tmp16777)


ifres16755 = tmp16778


} else {
ifres16755 = False


}

__e.TailApply(tmp16752, ifres16755)
return


} else {
__e.Return(W5039)
return
}


}, 1)

tmp16789 := Call(__e, PrimFunc(symshen_4unlocked_2), V5035)


var ifres16782 Obj

if True == tmp16789 {
tmp16783 := MakeNative(func(__e *ControlFlow) {
W5040 := __e.Get(1)
_ = W5040
tmp16786 := PrimEqual(W5040, Nil)

if True == tmp16786 {
tmp16784 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16784

__e.TailApply(PrimFunc(symthaw), V5037)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16787 := Call(__e, PrimFunc(symshen_4lazyderef), V5031, V5034)


tmp16788 := Call(__e, tmp16783, tmp16787)


ifres16782 = tmp16788


} else {
ifres16782 = False


}

__e.TailApply(tmp16751, ifres16782)
return


}, 1)

tmp16790 := PrimNumberAdd(V5036, MakeNumber(1))

__e.TailApply(tmp16750, tmp16790)
return


}, 8)

tmp16791 := Call(__e, ns2_1set, symshen_4t_d_1rules, tmp16749)


_ = tmp16791

tmp16792 := MakeNative(func(__e *ControlFlow) {
V5046 := __e.Get(1)
_ = V5046
tmp16805 := Call(__e, PrimFunc(symtuple_2), V5046)


if True == tmp16805 {
tmp16793 := MakeNative(func(__e *ControlFlow) {
W5047 := __e.Get(1)
_ = W5047
tmp16794 := MakeNative(func(__e *ControlFlow) {
W5048 := __e.Get(1)
_ = W5048
tmp16795 := Call(__e, PrimFunc(symfst), V5046)


tmp16796 := Call(__e, PrimFunc(symshen_4freshen), W5048, tmp16795)


tmp16797 := Call(__e, PrimFunc(symsnd), V5046)


tmp16798 := Call(__e, PrimFunc(symshen_4freshen), W5048, tmp16797)


__e.TailApply(PrimFunc(sym_8p), tmp16796, tmp16798)
return


}, 1)

tmp16799 := MakeNative(func(__e *ControlFlow) {
Z5049 := __e.Get(1)
_ = Z5049
tmp16800 := Call(__e, PrimFunc(symshen_4freshterm), Z5049)


__e.Return(PrimCons(Z5049, tmp16800))
return


}, 1)

tmp16801 := Call(__e, PrimFunc(symmap), tmp16799, W5047)


__e.TailApply(tmp16794, tmp16801)
return


}, 1)

tmp16802 := Call(__e, PrimFunc(symfst), V5046)


tmp16803 := Call(__e, PrimFunc(symshen_4extract_1vars), tmp16802)


__e.TailApply(tmp16793, tmp16803)
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.freshen-rule")))
return
}


}, 1)

tmp16806 := Call(__e, ns2_1set, symshen_4freshen_1rule, tmp16792)


_ = tmp16806

tmp16807 := MakeNative(func(__e *ControlFlow) {
V5050 := __e.Get(1)
_ = V5050
V5051 := __e.Get(2)
_ = V5051
tmp16821 := PrimEqual(Nil, V5050)

if True == tmp16821 {
__e.Return(V5051)
return
} else {
tmp16819 := PrimIsPair(V5050)

var ifres16815 Obj

if True == tmp16819 {
tmp16817 := PrimHead(V5050)

tmp16818 := PrimIsPair(tmp16817)

var ifres16816 Obj

if True == tmp16818 {
ifres16816 = True


} else {
ifres16816 = False


}

ifres16815 = ifres16816


} else {
ifres16815 = False


}

if True == ifres16815 {
tmp16808 := PrimTail(V5050)

tmp16809 := PrimHead(V5050)

tmp16810 := PrimHead(tmp16809)

tmp16811 := PrimHead(V5050)

tmp16812 := PrimTail(tmp16811)

tmp16813 := Call(__e, PrimFunc(symshen_4beta), tmp16810, tmp16812, V5051)


__e.TailApply(PrimFunc(symshen_4freshen), tmp16808, tmp16813)
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.freshen")))
return
}


}


}, 2)

tmp16822 := Call(__e, ns2_1set, symshen_4freshen, tmp16807)


_ = tmp16822

tmp16823 := MakeNative(func(__e *ControlFlow) {
V5052 := __e.Get(1)
_ = V5052
V5053 := __e.Get(2)
_ = V5053
V5054 := __e.Get(3)
_ = V5054
V5055 := __e.Get(4)
_ = V5055
V5056 := __e.Get(5)
_ = V5056
V5057 := __e.Get(6)
_ = V5057
V5058 := __e.Get(7)
_ = V5058
V5059 := __e.Get(8)
_ = V5059
V5060 := __e.Get(9)
_ = V5060
tmp16824 := MakeNative(func(__e *ControlFlow) {
W5061 := __e.Get(1)
_ = W5061
tmp16837 := PrimEqual(W5061, False)

if True == tmp16837 {
tmp16835 := Call(__e, PrimFunc(symshen_4unlocked_2), V5058)


if True == tmp16835 {
tmp16825 := MakeNative(func(__e *ControlFlow) {
W5062 := __e.Get(1)
_ = W5062
tmp16826 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16826

tmp16827 := Call(__e, PrimFunc(symshen_4app), V5052, MakeString("\n"), symshen_4a)


tmp16828 := PrimStringConcat(MakeString(" of "), tmp16827)

tmp16829 := Call(__e, PrimFunc(symshen_4app), V5053, tmp16828, symshen_4a)


tmp16830 := PrimStringConcat(MakeString("type error in rule "), tmp16829)

tmp16831 := PrimSimpleError(tmp16830)

tmp16832 := Call(__e, PrimFunc(symbind), W5062, tmp16831, V5057, V5058, V5059, V5060)


__e.TailApply(PrimFunc(symshen_4gc), V5057, tmp16832)
return


}, 1)

tmp16833 := Call(__e, PrimFunc(symshen_4newpv), V5057)


__e.TailApply(tmp16825, tmp16833)
return


} else {
__e.Return(False)
return
}


} else {
__e.Return(W5061)
return
}


}, 1)

tmp16841 := Call(__e, PrimFunc(symshen_4unlocked_2), V5058)


var ifres16838 Obj

if True == tmp16841 {
tmp16839 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16839

tmp16840 := Call(__e, PrimFunc(symshen_4t_d_1rule_1h), V5054, V5055, V5056, V5057, V5058, V5059, V5060)


ifres16838 = tmp16840


} else {
ifres16838 = False


}

__e.TailApply(tmp16824, ifres16838)
return


}, 9)

tmp16842 := Call(__e, ns2_1set, symshen_4t_d_1rule, tmp16823)


_ = tmp16842

tmp16843 := MakeNative(func(__e *ControlFlow) {
V5063 := __e.Get(1)
_ = V5063
V5064 := __e.Get(2)
_ = V5064
V5065 := __e.Get(3)
_ = V5065
V5066 := __e.Get(4)
_ = V5066
V5067 := __e.Get(5)
_ = V5067
V5068 := __e.Get(6)
_ = V5068
V5069 := __e.Get(7)
_ = V5069
tmp16844 := MakeNative(func(__e *ControlFlow) {
W5070 := __e.Get(1)
_ = W5070
tmp16845 := MakeNative(func(__e *ControlFlow) {
W5071 := __e.Get(1)
_ = W5071
tmp16868 := PrimEqual(W5071, False)

if True == tmp16868 {
tmp16846 := MakeNative(func(__e *ControlFlow) {
W5078 := __e.Get(1)
_ = W5078
tmp16848 := PrimEqual(W5078, False)

if True == tmp16848 {
__e.TailApply(PrimFunc(symshen_4unlock), V5067, W5070)
return
} else {
__e.Return(W5078)
return
}


}, 1)

tmp16866 := Call(__e, PrimFunc(symshen_4unlocked_2), V5067)


var ifres16849 Obj

if True == tmp16866 {
tmp16850 := MakeNative(func(__e *ControlFlow) {
W5079 := __e.Get(1)
_ = W5079
tmp16851 := MakeNative(func(__e *ControlFlow) {
W5080 := __e.Get(1)
_ = W5080
tmp16852 := MakeNative(func(__e *ControlFlow) {
W5081 := __e.Get(1)
_ = W5081
tmp16853 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16853

tmp16854 := Call(__e, PrimFunc(symshen_4freshterms), V5063)


tmp16855 := MakeNative(func(__e *ControlFlow) {
tmp16856 := MakeNative(func(__e *ControlFlow) {
tmp16857 := MakeNative(func(__e *ControlFlow) {
tmp16858 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4t_d_1correct), V5064, W5080, W5081, V5066, V5067, W5070, V5069)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4myassume), V5063, V5065, W5081, V5066, V5067, W5070, tmp16858)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V5066, V5067, W5070, tmp16857)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4t_d_1integrity), V5063, V5065, W5079, W5080, V5066, V5067, W5070, tmp16856)
return


}, 0)

tmp16859 := Call(__e, PrimFunc(symshen_4p_1hyps), tmp16854, W5079, V5066, V5067, W5070, tmp16855)


__e.TailApply(PrimFunc(symshen_4gc), V5066, tmp16859)
return


}, 1)

tmp16860 := Call(__e, PrimFunc(symshen_4newpv), V5066)


tmp16861 := Call(__e, tmp16852, tmp16860)


__e.TailApply(PrimFunc(symshen_4gc), V5066, tmp16861)
return


}, 1)

tmp16862 := Call(__e, PrimFunc(symshen_4newpv), V5066)


tmp16863 := Call(__e, tmp16851, tmp16862)


__e.TailApply(PrimFunc(symshen_4gc), V5066, tmp16863)
return


}, 1)

tmp16864 := Call(__e, PrimFunc(symshen_4newpv), V5066)


tmp16865 := Call(__e, tmp16850, tmp16864)


ifres16849 = tmp16865


} else {
ifres16849 = False


}

__e.TailApply(tmp16846, ifres16849)
return


} else {
__e.Return(W5071)
return
}


}, 1)

tmp16898 := Call(__e, PrimFunc(symshen_4unlocked_2), V5067)


var ifres16869 Obj

if True == tmp16898 {
tmp16870 := MakeNative(func(__e *ControlFlow) {
W5072 := __e.Get(1)
_ = W5072
tmp16895 := PrimEqual(W5072, Nil)

if True == tmp16895 {
tmp16871 := MakeNative(func(__e *ControlFlow) {
W5073 := __e.Get(1)
_ = W5073
tmp16892 := PrimIsPair(W5073)

if True == tmp16892 {
tmp16872 := MakeNative(func(__e *ControlFlow) {
W5074 := __e.Get(1)
_ = W5074
tmp16888 := PrimEqual(W5074, sym_1_1_6)

if True == tmp16888 {
tmp16873 := MakeNative(func(__e *ControlFlow) {
W5075 := __e.Get(1)
_ = W5075
tmp16884 := PrimIsPair(W5075)

if True == tmp16884 {
tmp16874 := MakeNative(func(__e *ControlFlow) {
W5076 := __e.Get(1)
_ = W5076
tmp16875 := MakeNative(func(__e *ControlFlow) {
W5077 := __e.Get(1)
_ = W5077
tmp16879 := PrimEqual(W5077, Nil)

if True == tmp16879 {
tmp16876 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16876

tmp16877 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4t_d_1correct), V5064, W5076, Nil, V5066, V5067, W5070, V5069)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V5066, V5067, W5070, tmp16877)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16880 := PrimTail(W5075)

tmp16881 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16880, V5066)


__e.TailApply(tmp16875, tmp16881)
return


}, 1)

tmp16882 := PrimHead(W5075)

__e.TailApply(tmp16874, tmp16882)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16885 := PrimTail(W5073)

tmp16886 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16885, V5066)


__e.TailApply(tmp16873, tmp16886)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16889 := PrimHead(W5073)

tmp16890 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16889, V5066)


__e.TailApply(tmp16872, tmp16890)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16893 := Call(__e, PrimFunc(symshen_4lazyderef), V5065, V5066)


__e.TailApply(tmp16871, tmp16893)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16896 := Call(__e, PrimFunc(symshen_4lazyderef), V5063, V5066)


tmp16897 := Call(__e, tmp16870, tmp16896)


ifres16869 = tmp16897


} else {
ifres16869 = False


}

__e.TailApply(tmp16845, ifres16869)
return


}, 1)

tmp16899 := PrimNumberAdd(V5068, MakeNumber(1))

__e.TailApply(tmp16844, tmp16899)
return


}, 7)

tmp16900 := Call(__e, ns2_1set, symshen_4t_d_1rule_1h, tmp16843)


_ = tmp16900

tmp16901 := MakeNative(func(__e *ControlFlow) {
V5082 := __e.Get(1)
_ = V5082
V5083 := __e.Get(2)
_ = V5083
V5084 := __e.Get(3)
_ = V5084
V5085 := __e.Get(4)
_ = V5085
V5086 := __e.Get(5)
_ = V5086
V5087 := __e.Get(6)
_ = V5087
V5088 := __e.Get(7)
_ = V5088
tmp16902 := MakeNative(func(__e *ControlFlow) {
W5089 := __e.Get(1)
_ = W5089
tmp17055 := PrimEqual(W5089, False)

if True == tmp17055 {
tmp17053 := Call(__e, PrimFunc(symshen_4unlocked_2), V5086)


if True == tmp17053 {
tmp16903 := MakeNative(func(__e *ControlFlow) {
W5093 := __e.Get(1)
_ = W5093
tmp17050 := PrimIsPair(W5093)

if True == tmp17050 {
tmp16904 := MakeNative(func(__e *ControlFlow) {
W5094 := __e.Get(1)
_ = W5094
tmp16905 := MakeNative(func(__e *ControlFlow) {
W5095 := __e.Get(1)
_ = W5095
tmp16906 := MakeNative(func(__e *ControlFlow) {
W5096 := __e.Get(1)
_ = W5096
tmp17045 := PrimIsPair(W5096)

if True == tmp17045 {
tmp16907 := MakeNative(func(__e *ControlFlow) {
W5097 := __e.Get(1)
_ = W5097
tmp16908 := MakeNative(func(__e *ControlFlow) {
W5098 := __e.Get(1)
_ = W5098
tmp17040 := PrimIsPair(W5098)

if True == tmp17040 {
tmp16909 := MakeNative(func(__e *ControlFlow) {
W5099 := __e.Get(1)
_ = W5099
tmp17036 := PrimEqual(W5099, sym_1_1_6)

if True == tmp17036 {
tmp16910 := MakeNative(func(__e *ControlFlow) {
W5100 := __e.Get(1)
_ = W5100
tmp17032 := PrimIsPair(W5100)

if True == tmp17032 {
tmp16911 := MakeNative(func(__e *ControlFlow) {
W5101 := __e.Get(1)
_ = W5101
tmp16912 := MakeNative(func(__e *ControlFlow) {
W5102 := __e.Get(1)
_ = W5102
tmp17027 := PrimEqual(W5102, Nil)

if True == tmp17027 {
tmp16913 := MakeNative(func(__e *ControlFlow) {
W5103 := __e.Get(1)
_ = W5103
tmp16914 := MakeNative(func(__e *ControlFlow) {
W5104 := __e.Get(1)
_ = W5104
tmp17018 := PrimIsPair(W5103)

if True == tmp17018 {
tmp16915 := MakeNative(func(__e *ControlFlow) {
W5109 := __e.Get(1)
_ = W5109
tmp16916 := MakeNative(func(__e *ControlFlow) {
W5110 := __e.Get(1)
_ = W5110
tmp16986 := PrimIsPair(W5109)

if True == tmp16986 {
tmp16917 := MakeNative(func(__e *ControlFlow) {
W5115 := __e.Get(1)
_ = W5115
tmp16918 := MakeNative(func(__e *ControlFlow) {
W5116 := __e.Get(1)
_ = W5116
tmp16919 := MakeNative(func(__e *ControlFlow) {
W5117 := __e.Get(1)
_ = W5117
tmp16961 := PrimIsPair(W5116)

if True == tmp16961 {
tmp16920 := MakeNative(func(__e *ControlFlow) {
W5120 := __e.Get(1)
_ = W5120
tmp16921 := MakeNative(func(__e *ControlFlow) {
W5121 := __e.Get(1)
_ = W5121
tmp16922 := MakeNative(func(__e *ControlFlow) {
W5122 := __e.Get(1)
_ = W5122
tmp16942 := PrimIsPair(W5121)

if True == tmp16942 {
tmp16923 := MakeNative(func(__e *ControlFlow) {
W5124 := __e.Get(1)
_ = W5124
tmp16924 := MakeNative(func(__e *ControlFlow) {
W5125 := __e.Get(1)
_ = W5125
tmp16925 := MakeNative(func(__e *ControlFlow) {
W5126 := __e.Get(1)
_ = W5126
tmp16929 := PrimEqual(W5125, Nil)

if True == tmp16929 {
__e.TailApply(PrimFunc(symthaw), W5126)
return
} else {
tmp16927 := Call(__e, PrimFunc(symshen_4pvar_2), W5125)


if True == tmp16927 {
__e.TailApply(PrimFunc(symshen_4bind_b), W5125, Nil, V5085, W5126)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp16930 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W5122, W5124)
return
}, 0)

__e.TailApply(tmp16925, tmp16930)
return


}, 1)

tmp16931 := PrimTail(W5121)

tmp16932 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16931, V5085)


__e.TailApply(tmp16924, tmp16932)
return


}, 1)

tmp16933 := PrimHead(W5121)

__e.TailApply(tmp16923, tmp16933)
return


} else {
tmp16940 := Call(__e, PrimFunc(symshen_4pvar_2), W5121)


if True == tmp16940 {
tmp16934 := MakeNative(func(__e *ControlFlow) {
W5127 := __e.Get(1)
_ = W5127
tmp16935 := PrimCons(W5127, Nil)

tmp16936 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W5122, W5127)
return
}, 0)

tmp16937 := Call(__e, PrimFunc(symshen_4bind_b), W5121, tmp16935, V5085, tmp16936)


__e.TailApply(PrimFunc(symshen_4gc), V5085, tmp16937)
return


}, 1)

tmp16938 := Call(__e, PrimFunc(symshen_4newpv), V5085)


__e.TailApply(tmp16934, tmp16938)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp16943 := MakeNative(func(__e *ControlFlow) {
Z5123 := __e.Get(1)
_ = Z5123
tmp16944 := Call(__e, W5117, W5120)


__e.TailApply(tmp16944, Z5123)
return


}, 1)

__e.TailApply(tmp16922, tmp16943)
return


}, 1)

tmp16945 := PrimTail(W5116)

tmp16946 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16945, V5085)


__e.TailApply(tmp16921, tmp16946)
return


}, 1)

tmp16947 := PrimHead(W5116)

__e.TailApply(tmp16920, tmp16947)
return


} else {
tmp16959 := Call(__e, PrimFunc(symshen_4pvar_2), W5116)


if True == tmp16959 {
tmp16948 := MakeNative(func(__e *ControlFlow) {
W5128 := __e.Get(1)
_ = W5128
tmp16949 := MakeNative(func(__e *ControlFlow) {
W5129 := __e.Get(1)
_ = W5129
tmp16950 := PrimCons(W5129, Nil)

tmp16951 := PrimCons(W5128, tmp16950)

tmp16952 := MakeNative(func(__e *ControlFlow) {
tmp16953 := Call(__e, W5117, W5128)


__e.TailApply(tmp16953, W5129)
return


}, 0)

tmp16954 := Call(__e, PrimFunc(symshen_4bind_b), W5116, tmp16951, V5085, tmp16952)


__e.TailApply(PrimFunc(symshen_4gc), V5085, tmp16954)
return


}, 1)

tmp16955 := Call(__e, PrimFunc(symshen_4newpv), V5085)


tmp16956 := Call(__e, tmp16949, tmp16955)


__e.TailApply(PrimFunc(symshen_4gc), V5085, tmp16956)
return


}, 1)

tmp16957 := Call(__e, PrimFunc(symshen_4newpv), V5085)


__e.TailApply(tmp16948, tmp16957)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp16962 := MakeNative(func(__e *ControlFlow) {
Z5118 := __e.Get(1)
_ = Z5118
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5119 := __e.Get(1)
_ = Z5119
tmp16963 := Call(__e, W5110, W5115)


tmp16964 := Call(__e, tmp16963, Z5118)


__e.TailApply(tmp16964, Z5119)
return


}, 1))
return
}, 1)

__e.TailApply(tmp16919, tmp16962)
return


}, 1)

tmp16965 := PrimTail(W5109)

tmp16966 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16965, V5085)


__e.TailApply(tmp16918, tmp16966)
return


}, 1)

tmp16967 := PrimHead(W5109)

__e.TailApply(tmp16917, tmp16967)
return


} else {
tmp16984 := Call(__e, PrimFunc(symshen_4pvar_2), W5109)


if True == tmp16984 {
tmp16968 := MakeNative(func(__e *ControlFlow) {
W5130 := __e.Get(1)
_ = W5130
tmp16969 := MakeNative(func(__e *ControlFlow) {
W5131 := __e.Get(1)
_ = W5131
tmp16970 := MakeNative(func(__e *ControlFlow) {
W5132 := __e.Get(1)
_ = W5132
tmp16971 := PrimCons(W5132, Nil)

tmp16972 := PrimCons(W5131, tmp16971)

tmp16973 := PrimCons(W5130, tmp16972)

tmp16974 := MakeNative(func(__e *ControlFlow) {
tmp16975 := Call(__e, W5110, W5130)


tmp16976 := Call(__e, tmp16975, W5131)


__e.TailApply(tmp16976, W5132)
return


}, 0)

tmp16977 := Call(__e, PrimFunc(symshen_4bind_b), W5109, tmp16973, V5085, tmp16974)


__e.TailApply(PrimFunc(symshen_4gc), V5085, tmp16977)
return


}, 1)

tmp16978 := Call(__e, PrimFunc(symshen_4newpv), V5085)


tmp16979 := Call(__e, tmp16970, tmp16978)


__e.TailApply(PrimFunc(symshen_4gc), V5085, tmp16979)
return


}, 1)

tmp16980 := Call(__e, PrimFunc(symshen_4newpv), V5085)


tmp16981 := Call(__e, tmp16969, tmp16980)


__e.TailApply(PrimFunc(symshen_4gc), V5085, tmp16981)
return


}, 1)

tmp16982 := Call(__e, PrimFunc(symshen_4newpv), V5085)


__e.TailApply(tmp16968, tmp16982)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp16987 := MakeNative(func(__e *ControlFlow) {
Z5111 := __e.Get(1)
_ = Z5111
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5112 := __e.Get(1)
_ = Z5112
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5113 := __e.Get(1)
_ = Z5113
tmp16988 := MakeNative(func(__e *ControlFlow) {
W5114 := __e.Get(1)
_ = W5114
tmp16989 := Call(__e, W5104, Z5111)


tmp16990 := Call(__e, tmp16989, Z5112)


tmp16991 := Call(__e, tmp16990, Z5113)


__e.TailApply(tmp16991, W5114)
return


}, 1)

tmp16992 := PrimTail(W5103)

__e.TailApply(tmp16988, tmp16992)
return


}, 1))
return
}, 1))
return
}, 1)

__e.TailApply(tmp16916, tmp16987)
return


}, 1)

tmp16993 := PrimHead(W5103)

tmp16994 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16993, V5085)


__e.TailApply(tmp16915, tmp16994)
return


} else {
tmp17016 := Call(__e, PrimFunc(symshen_4pvar_2), W5103)


if True == tmp17016 {
tmp16995 := MakeNative(func(__e *ControlFlow) {
W5133 := __e.Get(1)
_ = W5133
tmp16996 := MakeNative(func(__e *ControlFlow) {
W5134 := __e.Get(1)
_ = W5134
tmp16997 := MakeNative(func(__e *ControlFlow) {
W5135 := __e.Get(1)
_ = W5135
tmp16998 := MakeNative(func(__e *ControlFlow) {
W5136 := __e.Get(1)
_ = W5136
tmp16999 := PrimCons(W5135, Nil)

tmp17000 := PrimCons(W5134, tmp16999)

tmp17001 := PrimCons(W5133, tmp17000)

tmp17002 := PrimCons(tmp17001, W5136)

tmp17003 := MakeNative(func(__e *ControlFlow) {
tmp17004 := Call(__e, W5104, W5133)


tmp17005 := Call(__e, tmp17004, W5134)


tmp17006 := Call(__e, tmp17005, W5135)


__e.TailApply(tmp17006, W5136)
return


}, 0)

tmp17007 := Call(__e, PrimFunc(symshen_4bind_b), W5103, tmp17002, V5085, tmp17003)


__e.TailApply(PrimFunc(symshen_4gc), V5085, tmp17007)
return


}, 1)

tmp17008 := Call(__e, PrimFunc(symshen_4newpv), V5085)


tmp17009 := Call(__e, tmp16998, tmp17008)


__e.TailApply(PrimFunc(symshen_4gc), V5085, tmp17009)
return


}, 1)

tmp17010 := Call(__e, PrimFunc(symshen_4newpv), V5085)


tmp17011 := Call(__e, tmp16997, tmp17010)


__e.TailApply(PrimFunc(symshen_4gc), V5085, tmp17011)
return


}, 1)

tmp17012 := Call(__e, PrimFunc(symshen_4newpv), V5085)


tmp17013 := Call(__e, tmp16996, tmp17012)


__e.TailApply(PrimFunc(symshen_4gc), V5085, tmp17013)
return


}, 1)

tmp17014 := Call(__e, PrimFunc(symshen_4newpv), V5085)


__e.TailApply(tmp16995, tmp17014)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp17019 := MakeNative(func(__e *ControlFlow) {
Z5105 := __e.Get(1)
_ = Z5105
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5106 := __e.Get(1)
_ = Z5106
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5107 := __e.Get(1)
_ = Z5107
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5108 := __e.Get(1)
_ = Z5108
tmp17020 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp17020

tmp17021 := MakeNative(func(__e *ControlFlow) {
tmp17022 := MakeNative(func(__e *ControlFlow) {
tmp17023 := PrimIntern(MakeString(":"))

tmp17024 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4myassume), W5095, W5101, Z5108, V5085, V5086, V5087, V5088)
return
}, 0)

__e.TailApply(PrimFunc(symbind), Z5106, tmp17023, V5085, V5086, V5087, tmp17024)
return


}, 0)

__e.TailApply(PrimFunc(symis_b), W5094, Z5105, V5085, V5086, V5087, tmp17022)
return


}, 0)

__e.TailApply(PrimFunc(symis_b), W5097, Z5107, V5085, V5086, V5087, tmp17021)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

__e.TailApply(tmp16914, tmp17019)
return


}, 1)

tmp17025 := Call(__e, PrimFunc(symshen_4lazyderef), V5084, V5085)


__e.TailApply(tmp16913, tmp17025)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp17028 := PrimTail(W5100)

tmp17029 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17028, V5085)


__e.TailApply(tmp16912, tmp17029)
return


}, 1)

tmp17030 := PrimHead(W5100)

__e.TailApply(tmp16911, tmp17030)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp17033 := PrimTail(W5098)

tmp17034 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17033, V5085)


__e.TailApply(tmp16910, tmp17034)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp17037 := PrimHead(W5098)

tmp17038 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17037, V5085)


__e.TailApply(tmp16909, tmp17038)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp17041 := PrimTail(W5096)

tmp17042 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17041, V5085)


__e.TailApply(tmp16908, tmp17042)
return


}, 1)

tmp17043 := PrimHead(W5096)

__e.TailApply(tmp16907, tmp17043)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp17046 := Call(__e, PrimFunc(symshen_4lazyderef), V5083, V5085)


__e.TailApply(tmp16906, tmp17046)
return


}, 1)

tmp17047 := PrimTail(W5093)

__e.TailApply(tmp16905, tmp17047)
return


}, 1)

tmp17048 := PrimHead(W5093)

__e.TailApply(tmp16904, tmp17048)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp17051 := Call(__e, PrimFunc(symshen_4lazyderef), V5082, V5085)


__e.TailApply(tmp16903, tmp17051)
return


} else {
__e.Return(False)
return
}


} else {
__e.Return(W5089)
return
}


}, 1)

tmp17071 := Call(__e, PrimFunc(symshen_4unlocked_2), V5086)


var ifres17056 Obj

if True == tmp17071 {
tmp17057 := MakeNative(func(__e *ControlFlow) {
W5090 := __e.Get(1)
_ = W5090
tmp17068 := PrimEqual(W5090, Nil)

if True == tmp17068 {
tmp17058 := MakeNative(func(__e *ControlFlow) {
W5091 := __e.Get(1)
_ = W5091
tmp17059 := MakeNative(func(__e *ControlFlow) {
W5092 := __e.Get(1)
_ = W5092
tmp17063 := PrimEqual(W5091, Nil)

if True == tmp17063 {
__e.TailApply(PrimFunc(symthaw), W5092)
return
} else {
tmp17061 := Call(__e, PrimFunc(symshen_4pvar_2), W5091)


if True == tmp17061 {
__e.TailApply(PrimFunc(symshen_4bind_b), W5091, Nil, V5085, W5092)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp17064 := MakeNative(func(__e *ControlFlow) {
tmp17065 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp17065

__e.TailApply(PrimFunc(symthaw), V5088)
return


}, 0)

__e.TailApply(tmp17059, tmp17064)
return


}, 1)

tmp17066 := Call(__e, PrimFunc(symshen_4lazyderef), V5084, V5085)


__e.TailApply(tmp17058, tmp17066)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp17069 := Call(__e, PrimFunc(symshen_4lazyderef), V5082, V5085)


tmp17070 := Call(__e, tmp17057, tmp17069)


ifres17056 = tmp17070


} else {
ifres17056 = False


}

__e.TailApply(tmp16902, ifres17056)
return


}, 7)

tmp17072 := Call(__e, ns2_1set, symshen_4myassume, tmp16901)


_ = tmp17072

tmp17073 := MakeNative(func(__e *ControlFlow) {
V5139 := __e.Get(1)
_ = V5139
tmp17096 := PrimEqual(Nil, V5139)

if True == tmp17096 {
__e.Return(Nil)
return
} else {
tmp17094 := PrimIsPair(V5139)

var ifres17090 Obj

if True == tmp17094 {
tmp17092 := PrimHead(V5139)

tmp17093 := PrimIsPair(tmp17092)

var ifres17091 Obj

if True == tmp17093 {
ifres17091 = True


} else {
ifres17091 = False


}

ifres17090 = ifres17091


} else {
ifres17090 = False


}

if True == ifres17090 {
tmp17074 := PrimHead(V5139)

tmp17075 := PrimTail(V5139)

tmp17076 := Call(__e, PrimFunc(symappend), tmp17074, tmp17075)


__e.TailApply(PrimFunc(symshen_4freshterms), tmp17076)
return


} else {
tmp17088 := PrimIsPair(V5139)

var ifres17084 Obj

if True == tmp17088 {
tmp17086 := PrimHead(V5139)

tmp17087 := Call(__e, PrimFunc(symshen_4freshterm_2), tmp17086)


var ifres17085 Obj

if True == tmp17087 {
ifres17085 = True


} else {
ifres17085 = False


}

ifres17084 = ifres17085


} else {
ifres17084 = False


}

if True == ifres17084 {
tmp17077 := PrimHead(V5139)

tmp17078 := PrimTail(V5139)

tmp17079 := Call(__e, PrimFunc(symshen_4freshterms), tmp17078)


__e.TailApply(PrimFunc(symadjoin), tmp17077, tmp17079)
return


} else {
tmp17082 := PrimIsPair(V5139)

if True == tmp17082 {
tmp17080 := PrimTail(V5139)

__e.TailApply(PrimFunc(symshen_4freshterms), tmp17080)
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.freshterms")))
return
}


}


}


}


}, 1)

tmp17097 := Call(__e, ns2_1set, symshen_4freshterms, tmp17073)


_ = tmp17097

tmp17098 := MakeNative(func(__e *ControlFlow) {
V5140 := __e.Get(1)
_ = V5140
V5141 := __e.Get(2)
_ = V5141
V5142 := __e.Get(3)
_ = V5142
V5143 := __e.Get(4)
_ = V5143
V5144 := __e.Get(5)
_ = V5144
V5145 := __e.Get(6)
_ = V5145
tmp17099 := MakeNative(func(__e *ControlFlow) {
W5146 := __e.Get(1)
_ = W5146
tmp17223 := PrimEqual(W5146, False)

if True == tmp17223 {
tmp17221 := Call(__e, PrimFunc(symshen_4unlocked_2), V5143)


if True == tmp17221 {
tmp17100 := MakeNative(func(__e *ControlFlow) {
W5150 := __e.Get(1)
_ = W5150
tmp17218 := PrimIsPair(W5150)

if True == tmp17218 {
tmp17101 := MakeNative(func(__e *ControlFlow) {
W5151 := __e.Get(1)
_ = W5151
tmp17102 := MakeNative(func(__e *ControlFlow) {
W5152 := __e.Get(1)
_ = W5152
tmp17103 := MakeNative(func(__e *ControlFlow) {
W5153 := __e.Get(1)
_ = W5153
tmp17104 := MakeNative(func(__e *ControlFlow) {
W5154 := __e.Get(1)
_ = W5154
tmp17208 := PrimIsPair(W5153)

if True == tmp17208 {
tmp17105 := MakeNative(func(__e *ControlFlow) {
W5159 := __e.Get(1)
_ = W5159
tmp17106 := MakeNative(func(__e *ControlFlow) {
W5160 := __e.Get(1)
_ = W5160
tmp17176 := PrimIsPair(W5159)

if True == tmp17176 {
tmp17107 := MakeNative(func(__e *ControlFlow) {
W5165 := __e.Get(1)
_ = W5165
tmp17108 := MakeNative(func(__e *ControlFlow) {
W5166 := __e.Get(1)
_ = W5166
tmp17109 := MakeNative(func(__e *ControlFlow) {
W5167 := __e.Get(1)
_ = W5167
tmp17151 := PrimIsPair(W5166)

if True == tmp17151 {
tmp17110 := MakeNative(func(__e *ControlFlow) {
W5170 := __e.Get(1)
_ = W5170
tmp17111 := MakeNative(func(__e *ControlFlow) {
W5171 := __e.Get(1)
_ = W5171
tmp17112 := MakeNative(func(__e *ControlFlow) {
W5172 := __e.Get(1)
_ = W5172
tmp17132 := PrimIsPair(W5171)

if True == tmp17132 {
tmp17113 := MakeNative(func(__e *ControlFlow) {
W5174 := __e.Get(1)
_ = W5174
tmp17114 := MakeNative(func(__e *ControlFlow) {
W5175 := __e.Get(1)
_ = W5175
tmp17115 := MakeNative(func(__e *ControlFlow) {
W5176 := __e.Get(1)
_ = W5176
tmp17119 := PrimEqual(W5175, Nil)

if True == tmp17119 {
__e.TailApply(PrimFunc(symthaw), W5176)
return
} else {
tmp17117 := Call(__e, PrimFunc(symshen_4pvar_2), W5175)


if True == tmp17117 {
__e.TailApply(PrimFunc(symshen_4bind_b), W5175, Nil, V5142, W5176)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp17120 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W5172, W5174)
return
}, 0)

__e.TailApply(tmp17115, tmp17120)
return


}, 1)

tmp17121 := PrimTail(W5171)

tmp17122 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17121, V5142)


__e.TailApply(tmp17114, tmp17122)
return


}, 1)

tmp17123 := PrimHead(W5171)

__e.TailApply(tmp17113, tmp17123)
return


} else {
tmp17130 := Call(__e, PrimFunc(symshen_4pvar_2), W5171)


if True == tmp17130 {
tmp17124 := MakeNative(func(__e *ControlFlow) {
W5177 := __e.Get(1)
_ = W5177
tmp17125 := PrimCons(W5177, Nil)

tmp17126 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W5172, W5177)
return
}, 0)

tmp17127 := Call(__e, PrimFunc(symshen_4bind_b), W5171, tmp17125, V5142, tmp17126)


__e.TailApply(PrimFunc(symshen_4gc), V5142, tmp17127)
return


}, 1)

tmp17128 := Call(__e, PrimFunc(symshen_4newpv), V5142)


__e.TailApply(tmp17124, tmp17128)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp17133 := MakeNative(func(__e *ControlFlow) {
Z5173 := __e.Get(1)
_ = Z5173
tmp17134 := Call(__e, W5167, W5170)


__e.TailApply(tmp17134, Z5173)
return


}, 1)

__e.TailApply(tmp17112, tmp17133)
return


}, 1)

tmp17135 := PrimTail(W5166)

tmp17136 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17135, V5142)


__e.TailApply(tmp17111, tmp17136)
return


}, 1)

tmp17137 := PrimHead(W5166)

__e.TailApply(tmp17110, tmp17137)
return


} else {
tmp17149 := Call(__e, PrimFunc(symshen_4pvar_2), W5166)


if True == tmp17149 {
tmp17138 := MakeNative(func(__e *ControlFlow) {
W5178 := __e.Get(1)
_ = W5178
tmp17139 := MakeNative(func(__e *ControlFlow) {
W5179 := __e.Get(1)
_ = W5179
tmp17140 := PrimCons(W5179, Nil)

tmp17141 := PrimCons(W5178, tmp17140)

tmp17142 := MakeNative(func(__e *ControlFlow) {
tmp17143 := Call(__e, W5167, W5178)


__e.TailApply(tmp17143, W5179)
return


}, 0)

tmp17144 := Call(__e, PrimFunc(symshen_4bind_b), W5166, tmp17141, V5142, tmp17142)


__e.TailApply(PrimFunc(symshen_4gc), V5142, tmp17144)
return


}, 1)

tmp17145 := Call(__e, PrimFunc(symshen_4newpv), V5142)


tmp17146 := Call(__e, tmp17139, tmp17145)


__e.TailApply(PrimFunc(symshen_4gc), V5142, tmp17146)
return


}, 1)

tmp17147 := Call(__e, PrimFunc(symshen_4newpv), V5142)


__e.TailApply(tmp17138, tmp17147)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp17152 := MakeNative(func(__e *ControlFlow) {
Z5168 := __e.Get(1)
_ = Z5168
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5169 := __e.Get(1)
_ = Z5169
tmp17153 := Call(__e, W5160, W5165)


tmp17154 := Call(__e, tmp17153, Z5168)


__e.TailApply(tmp17154, Z5169)
return


}, 1))
return
}, 1)

__e.TailApply(tmp17109, tmp17152)
return


}, 1)

tmp17155 := PrimTail(W5159)

tmp17156 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17155, V5142)


__e.TailApply(tmp17108, tmp17156)
return


}, 1)

tmp17157 := PrimHead(W5159)

__e.TailApply(tmp17107, tmp17157)
return


} else {
tmp17174 := Call(__e, PrimFunc(symshen_4pvar_2), W5159)


if True == tmp17174 {
tmp17158 := MakeNative(func(__e *ControlFlow) {
W5180 := __e.Get(1)
_ = W5180
tmp17159 := MakeNative(func(__e *ControlFlow) {
W5181 := __e.Get(1)
_ = W5181
tmp17160 := MakeNative(func(__e *ControlFlow) {
W5182 := __e.Get(1)
_ = W5182
tmp17161 := PrimCons(W5182, Nil)

tmp17162 := PrimCons(W5181, tmp17161)

tmp17163 := PrimCons(W5180, tmp17162)

tmp17164 := MakeNative(func(__e *ControlFlow) {
tmp17165 := Call(__e, W5160, W5180)


tmp17166 := Call(__e, tmp17165, W5181)


__e.TailApply(tmp17166, W5182)
return


}, 0)

tmp17167 := Call(__e, PrimFunc(symshen_4bind_b), W5159, tmp17163, V5142, tmp17164)


__e.TailApply(PrimFunc(symshen_4gc), V5142, tmp17167)
return


}, 1)

tmp17168 := Call(__e, PrimFunc(symshen_4newpv), V5142)


tmp17169 := Call(__e, tmp17160, tmp17168)


__e.TailApply(PrimFunc(symshen_4gc), V5142, tmp17169)
return


}, 1)

tmp17170 := Call(__e, PrimFunc(symshen_4newpv), V5142)


tmp17171 := Call(__e, tmp17159, tmp17170)


__e.TailApply(PrimFunc(symshen_4gc), V5142, tmp17171)
return


}, 1)

tmp17172 := Call(__e, PrimFunc(symshen_4newpv), V5142)


__e.TailApply(tmp17158, tmp17172)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp17177 := MakeNative(func(__e *ControlFlow) {
Z5161 := __e.Get(1)
_ = Z5161
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5162 := __e.Get(1)
_ = Z5162
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5163 := __e.Get(1)
_ = Z5163
tmp17178 := MakeNative(func(__e *ControlFlow) {
W5164 := __e.Get(1)
_ = W5164
tmp17179 := Call(__e, W5154, Z5161)


tmp17180 := Call(__e, tmp17179, Z5162)


tmp17181 := Call(__e, tmp17180, Z5163)


__e.TailApply(tmp17181, W5164)
return


}, 1)

tmp17182 := PrimTail(W5153)

__e.TailApply(tmp17178, tmp17182)
return


}, 1))
return
}, 1))
return
}, 1)

__e.TailApply(tmp17106, tmp17177)
return


}, 1)

tmp17183 := PrimHead(W5153)

tmp17184 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17183, V5142)


__e.TailApply(tmp17105, tmp17184)
return


} else {
tmp17206 := Call(__e, PrimFunc(symshen_4pvar_2), W5153)


if True == tmp17206 {
tmp17185 := MakeNative(func(__e *ControlFlow) {
W5183 := __e.Get(1)
_ = W5183
tmp17186 := MakeNative(func(__e *ControlFlow) {
W5184 := __e.Get(1)
_ = W5184
tmp17187 := MakeNative(func(__e *ControlFlow) {
W5185 := __e.Get(1)
_ = W5185
tmp17188 := MakeNative(func(__e *ControlFlow) {
W5186 := __e.Get(1)
_ = W5186
tmp17189 := PrimCons(W5185, Nil)

tmp17190 := PrimCons(W5184, tmp17189)

tmp17191 := PrimCons(W5183, tmp17190)

tmp17192 := PrimCons(tmp17191, W5186)

tmp17193 := MakeNative(func(__e *ControlFlow) {
tmp17194 := Call(__e, W5154, W5183)


tmp17195 := Call(__e, tmp17194, W5184)


tmp17196 := Call(__e, tmp17195, W5185)


__e.TailApply(tmp17196, W5186)
return


}, 0)

tmp17197 := Call(__e, PrimFunc(symshen_4bind_b), W5153, tmp17192, V5142, tmp17193)


__e.TailApply(PrimFunc(symshen_4gc), V5142, tmp17197)
return


}, 1)

tmp17198 := Call(__e, PrimFunc(symshen_4newpv), V5142)


tmp17199 := Call(__e, tmp17188, tmp17198)


__e.TailApply(PrimFunc(symshen_4gc), V5142, tmp17199)
return


}, 1)

tmp17200 := Call(__e, PrimFunc(symshen_4newpv), V5142)


tmp17201 := Call(__e, tmp17187, tmp17200)


__e.TailApply(PrimFunc(symshen_4gc), V5142, tmp17201)
return


}, 1)

tmp17202 := Call(__e, PrimFunc(symshen_4newpv), V5142)


tmp17203 := Call(__e, tmp17186, tmp17202)


__e.TailApply(PrimFunc(symshen_4gc), V5142, tmp17203)
return


}, 1)

tmp17204 := Call(__e, PrimFunc(symshen_4newpv), V5142)


__e.TailApply(tmp17185, tmp17204)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp17209 := MakeNative(func(__e *ControlFlow) {
Z5155 := __e.Get(1)
_ = Z5155
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5156 := __e.Get(1)
_ = Z5156
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5157 := __e.Get(1)
_ = Z5157
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5158 := __e.Get(1)
_ = Z5158
tmp17210 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp17210

tmp17211 := MakeNative(func(__e *ControlFlow) {
tmp17212 := PrimIntern(MakeString(":"))

tmp17213 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4p_1hyps), W5152, Z5158, V5142, V5143, V5144, V5145)
return
}, 0)

__e.TailApply(PrimFunc(symbind), Z5156, tmp17212, V5142, V5143, V5144, tmp17213)
return


}, 0)

__e.TailApply(PrimFunc(symbind), Z5155, W5151, V5142, V5143, V5144, tmp17211)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

__e.TailApply(tmp17104, tmp17209)
return


}, 1)

tmp17214 := Call(__e, PrimFunc(symshen_4lazyderef), V5141, V5142)


__e.TailApply(tmp17103, tmp17214)
return


}, 1)

tmp17215 := PrimTail(W5150)

__e.TailApply(tmp17102, tmp17215)
return


}, 1)

tmp17216 := PrimHead(W5150)

__e.TailApply(tmp17101, tmp17216)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp17219 := Call(__e, PrimFunc(symshen_4lazyderef), V5140, V5142)


__e.TailApply(tmp17100, tmp17219)
return


} else {
__e.Return(False)
return
}


} else {
__e.Return(W5146)
return
}


}, 1)

tmp17239 := Call(__e, PrimFunc(symshen_4unlocked_2), V5143)


var ifres17224 Obj

if True == tmp17239 {
tmp17225 := MakeNative(func(__e *ControlFlow) {
W5147 := __e.Get(1)
_ = W5147
tmp17236 := PrimEqual(W5147, Nil)

if True == tmp17236 {
tmp17226 := MakeNative(func(__e *ControlFlow) {
W5148 := __e.Get(1)
_ = W5148
tmp17227 := MakeNative(func(__e *ControlFlow) {
W5149 := __e.Get(1)
_ = W5149
tmp17231 := PrimEqual(W5148, Nil)

if True == tmp17231 {
__e.TailApply(PrimFunc(symthaw), W5149)
return
} else {
tmp17229 := Call(__e, PrimFunc(symshen_4pvar_2), W5148)


if True == tmp17229 {
__e.TailApply(PrimFunc(symshen_4bind_b), W5148, Nil, V5142, W5149)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp17232 := MakeNative(func(__e *ControlFlow) {
tmp17233 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp17233

__e.TailApply(PrimFunc(symthaw), V5145)
return


}, 0)

__e.TailApply(tmp17227, tmp17232)
return


}, 1)

tmp17234 := Call(__e, PrimFunc(symshen_4lazyderef), V5141, V5142)


__e.TailApply(tmp17226, tmp17234)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp17237 := Call(__e, PrimFunc(symshen_4lazyderef), V5140, V5142)


tmp17238 := Call(__e, tmp17225, tmp17237)


ifres17224 = tmp17238


} else {
ifres17224 = False


}

__e.TailApply(tmp17099, ifres17224)
return


}, 6)

tmp17240 := Call(__e, ns2_1set, symshen_4p_1hyps, tmp17098)


_ = tmp17240

tmp17241 := MakeNative(func(__e *ControlFlow) {
V5187 := __e.Get(1)
_ = V5187
V5188 := __e.Get(2)
_ = V5188
V5189 := __e.Get(3)
_ = V5189
V5190 := __e.Get(4)
_ = V5190
V5191 := __e.Get(5)
_ = V5191
V5192 := __e.Get(6)
_ = V5192
V5193 := __e.Get(7)
_ = V5193
tmp17242 := MakeNative(func(__e *ControlFlow) {
W5194 := __e.Get(1)
_ = W5194
tmp17243 := MakeNative(func(__e *ControlFlow) {
W5195 := __e.Get(1)
_ = W5195
tmp17253 := PrimEqual(W5195, False)

if True == tmp17253 {
tmp17244 := MakeNative(func(__e *ControlFlow) {
W5204 := __e.Get(1)
_ = W5204
tmp17246 := PrimEqual(W5204, False)

if True == tmp17246 {
__e.TailApply(PrimFunc(symshen_4unlock), V5191, W5194)
return
} else {
__e.Return(W5204)
return
}


}, 1)

tmp17251 := Call(__e, PrimFunc(symshen_4unlocked_2), V5191)


var ifres17247 Obj

if True == tmp17251 {
tmp17248 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp17248

tmp17249 := Call(__e, PrimFunc(symshen_4curry), V5187)


tmp17250 := Call(__e, PrimFunc(symshen_4system_1S_1h), tmp17249, V5188, V5189, V5190, V5191, W5194, V5193)


ifres17247 = tmp17250


} else {
ifres17247 = False


}

__e.TailApply(tmp17244, ifres17247)
return


} else {
__e.Return(W5195)
return
}


}, 1)

tmp17298 := Call(__e, PrimFunc(symshen_4unlocked_2), V5191)


var ifres17254 Obj

if True == tmp17298 {
tmp17255 := MakeNative(func(__e *ControlFlow) {
W5196 := __e.Get(1)
_ = W5196
tmp17295 := PrimIsPair(W5196)

if True == tmp17295 {
tmp17256 := MakeNative(func(__e *ControlFlow) {
W5197 := __e.Get(1)
_ = W5197
tmp17291 := PrimEqual(W5197, symwhere)

if True == tmp17291 {
tmp17257 := MakeNative(func(__e *ControlFlow) {
W5198 := __e.Get(1)
_ = W5198
tmp17287 := PrimIsPair(W5198)

if True == tmp17287 {
tmp17258 := MakeNative(func(__e *ControlFlow) {
W5199 := __e.Get(1)
_ = W5199
tmp17259 := MakeNative(func(__e *ControlFlow) {
W5200 := __e.Get(1)
_ = W5200
tmp17282 := PrimIsPair(W5200)

if True == tmp17282 {
tmp17260 := MakeNative(func(__e *ControlFlow) {
W5201 := __e.Get(1)
_ = W5201
tmp17261 := MakeNative(func(__e *ControlFlow) {
W5202 := __e.Get(1)
_ = W5202
tmp17277 := PrimEqual(W5202, Nil)

if True == tmp17277 {
tmp17262 := MakeNative(func(__e *ControlFlow) {
W5203 := __e.Get(1)
_ = W5203
tmp17263 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp17263

tmp17264 := MakeNative(func(__e *ControlFlow) {
tmp17265 := Call(__e, PrimFunc(symshen_4curry), W5199)


tmp17266 := MakeNative(func(__e *ControlFlow) {
tmp17267 := MakeNative(func(__e *ControlFlow) {
tmp17268 := MakeNative(func(__e *ControlFlow) {
tmp17269 := PrimIntern(MakeString(":"))

tmp17270 := PrimCons(symverified, Nil)

tmp17271 := PrimCons(tmp17269, tmp17270)

tmp17272 := PrimCons(W5203, tmp17271)

tmp17273 := PrimCons(tmp17272, V5189)

__e.TailApply(PrimFunc(symshen_4t_d_1correct), W5201, V5188, tmp17273, V5190, V5191, W5194, V5193)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V5190, V5191, W5194, tmp17268)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), W5203, symboolean, V5189, V5190, V5191, W5194, tmp17267)
return


}, 0)

__e.TailApply(PrimFunc(symbind), W5203, tmp17265, V5190, V5191, W5194, tmp17266)
return


}, 0)

tmp17274 := Call(__e, PrimFunc(symshen_4cut), V5190, V5191, W5194, tmp17264)


__e.TailApply(PrimFunc(symshen_4gc), V5190, tmp17274)
return


}, 1)

tmp17275 := Call(__e, PrimFunc(symshen_4newpv), V5190)


__e.TailApply(tmp17262, tmp17275)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp17278 := PrimTail(W5200)

tmp17279 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17278, V5190)


__e.TailApply(tmp17261, tmp17279)
return


}, 1)

tmp17280 := PrimHead(W5200)

__e.TailApply(tmp17260, tmp17280)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp17283 := PrimTail(W5198)

tmp17284 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17283, V5190)


__e.TailApply(tmp17259, tmp17284)
return


}, 1)

tmp17285 := PrimHead(W5198)

__e.TailApply(tmp17258, tmp17285)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp17288 := PrimTail(W5196)

tmp17289 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17288, V5190)


__e.TailApply(tmp17257, tmp17289)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp17292 := PrimHead(W5196)

tmp17293 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17292, V5190)


__e.TailApply(tmp17256, tmp17293)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp17296 := Call(__e, PrimFunc(symshen_4lazyderef), V5187, V5190)


tmp17297 := Call(__e, tmp17255, tmp17296)


ifres17254 = tmp17297


} else {
ifres17254 = False


}

__e.TailApply(tmp17243, ifres17254)
return


}, 1)

tmp17299 := PrimNumberAdd(V5192, MakeNumber(1))

__e.TailApply(tmp17242, tmp17299)
return


}, 7)

tmp17300 := Call(__e, ns2_1set, symshen_4t_d_1correct, tmp17241)


_ = tmp17300

tmp17301 := MakeNative(func(__e *ControlFlow) {
V5205 := __e.Get(1)
_ = V5205
V5206 := __e.Get(2)
_ = V5206
V5207 := __e.Get(3)
_ = V5207
V5208 := __e.Get(4)
_ = V5208
V5209 := __e.Get(5)
_ = V5209
V5210 := __e.Get(6)
_ = V5210
V5211 := __e.Get(7)
_ = V5211
V5212 := __e.Get(8)
_ = V5212
tmp17302 := MakeNative(func(__e *ControlFlow) {
W5213 := __e.Get(1)
_ = W5213
tmp17344 := PrimEqual(W5213, False)

if True == tmp17344 {
tmp17342 := Call(__e, PrimFunc(symshen_4unlocked_2), V5210)


if True == tmp17342 {
tmp17303 := MakeNative(func(__e *ControlFlow) {
W5215 := __e.Get(1)
_ = W5215
tmp17339 := PrimIsPair(W5215)

if True == tmp17339 {
tmp17304 := MakeNative(func(__e *ControlFlow) {
W5216 := __e.Get(1)
_ = W5216
tmp17305 := MakeNative(func(__e *ControlFlow) {
W5217 := __e.Get(1)
_ = W5217
tmp17306 := MakeNative(func(__e *ControlFlow) {
W5218 := __e.Get(1)
_ = W5218
tmp17334 := PrimIsPair(W5218)

if True == tmp17334 {
tmp17307 := MakeNative(func(__e *ControlFlow) {
W5219 := __e.Get(1)
_ = W5219
tmp17308 := MakeNative(func(__e *ControlFlow) {
W5220 := __e.Get(1)
_ = W5220
tmp17329 := PrimIsPair(W5220)

if True == tmp17329 {
tmp17309 := MakeNative(func(__e *ControlFlow) {
W5221 := __e.Get(1)
_ = W5221
tmp17325 := PrimEqual(W5221, sym_1_1_6)

if True == tmp17325 {
tmp17310 := MakeNative(func(__e *ControlFlow) {
W5222 := __e.Get(1)
_ = W5222
tmp17321 := PrimIsPair(W5222)

if True == tmp17321 {
tmp17311 := MakeNative(func(__e *ControlFlow) {
W5223 := __e.Get(1)
_ = W5223
tmp17312 := MakeNative(func(__e *ControlFlow) {
W5224 := __e.Get(1)
_ = W5224
tmp17316 := PrimEqual(W5224, Nil)

if True == tmp17316 {
tmp17313 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp17313

tmp17314 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4t_d_1integrity), W5217, W5223, V5207, V5208, V5209, V5210, V5211, V5212)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), W5216, W5219, V5207, V5209, V5210, V5211, tmp17314)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp17317 := PrimTail(W5222)

tmp17318 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17317, V5209)


__e.TailApply(tmp17312, tmp17318)
return


}, 1)

tmp17319 := PrimHead(W5222)

__e.TailApply(tmp17311, tmp17319)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp17322 := PrimTail(W5220)

tmp17323 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17322, V5209)


__e.TailApply(tmp17310, tmp17323)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp17326 := PrimHead(W5220)

tmp17327 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17326, V5209)


__e.TailApply(tmp17309, tmp17327)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp17330 := PrimTail(W5218)

tmp17331 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17330, V5209)


__e.TailApply(tmp17308, tmp17331)
return


}, 1)

tmp17332 := PrimHead(W5218)

__e.TailApply(tmp17307, tmp17332)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp17335 := Call(__e, PrimFunc(symshen_4lazyderef), V5206, V5209)


__e.TailApply(tmp17306, tmp17335)
return


}, 1)

tmp17336 := PrimTail(W5215)

__e.TailApply(tmp17305, tmp17336)
return


}, 1)

tmp17337 := PrimHead(W5215)

__e.TailApply(tmp17304, tmp17337)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp17340 := Call(__e, PrimFunc(symshen_4lazyderef), V5205, V5209)


__e.TailApply(tmp17303, tmp17340)
return


} else {
__e.Return(False)
return
}


} else {
__e.Return(W5213)
return
}


}, 1)

tmp17352 := Call(__e, PrimFunc(symshen_4unlocked_2), V5210)


var ifres17345 Obj

if True == tmp17352 {
tmp17346 := MakeNative(func(__e *ControlFlow) {
W5214 := __e.Get(1)
_ = W5214
tmp17349 := PrimEqual(W5214, Nil)

if True == tmp17349 {
tmp17347 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp17347

__e.TailApply(PrimFunc(symis_b), V5206, V5208, V5209, V5210, V5211, V5212)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp17350 := Call(__e, PrimFunc(symshen_4lazyderef), V5205, V5209)


tmp17351 := Call(__e, tmp17346, tmp17350)


ifres17345 = tmp17351


} else {
ifres17345 = False


}

__e.TailApply(tmp17302, ifres17345)
return


}, 8)

tmp17353 := Call(__e, ns2_1set, symshen_4t_d_1integrity, tmp17301)


_ = tmp17353

tmp17354 := MakeNative(func(__e *ControlFlow) {
V5225 := __e.Get(1)
_ = V5225
tmp17363 := PrimIsVector(V5225)

if True == tmp17363 {
tmp17360 := PrimIsString(V5225)

tmp17361 := PrimNot(tmp17360)

var ifres17356 Obj

if True == tmp17361 {
tmp17358 := PrimVectorGet(V5225, MakeNumber(0))

tmp17359 := PrimEqual(tmp17358, symshen_4print_1freshterm)

var ifres17357 Obj

if True == tmp17359 {
ifres17357 = True


} else {
ifres17357 = False


}

ifres17356 = ifres17357


} else {
ifres17356 = False


}

if True == ifres17356 {
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

__e.TailApply(ns2_1set, symshen_4freshterm_2, tmp17354)
return




}, 0)

