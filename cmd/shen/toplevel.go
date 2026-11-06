package main

import . "github.com/pyrex41/shen-go/kl"

var TopLevelMain = MakeNative0(func(__e *ControlFlow) Obj {
tmp8187 := MakeNative0(func(__e *ControlFlow) Obj {
tmp8188 := Call0(__e, SymFn(symshen_4credits))


_ = tmp8188

return TailCall0(__e, SymFn(symshen_4loop))


panic("unreachable")
})

tmp8189 := Call2(__e, ns2_1set, symshen_4repl, tmp8187)


_ = tmp8189

tmp8190 := (func() Obj {
var __self Obj
__self = MakeNative0(func(__e *ControlFlow) Obj {
__selftop:
tmp8191 := Call0(__e, SymFn(symshen_4initialise__environment))


_ = tmp8191

tmp8192 := Call0(__e, SymFn(symshen_4prompt))


_ = tmp8192

tmp8193 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall0(__e, SymFn(symshen_4read_1evaluate_1print))
panic("unreachable")
})

tmp8194 := MakeNative1(func(__e *ControlFlow, Z5748 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4toplevel_1display_1exception), Z5748)
panic("unreachable")
})

tmp8195 := Call2(__e, try_1catch, tmp8193, tmp8194)


_ = tmp8195

if SymFn(symshen_4loop) == __self {
__e.Tick()
goto __selftop
}
return TailCall0(__e, SymFn(symshen_4loop))


panic("unreachable")
})
return __self
})()

tmp8196 := Call2(__e, ns2_1set, symshen_4loop, tmp8190)


_ = tmp8196

tmp8197 := MakeNative1(func(__e *ControlFlow, V5749 Obj) Obj {
tmp8198 := PrimGuard1(__e, symerror_1to_1string, PrimErrorToString, V5749)

tmp8199 := Call0(__e, SymFn(symstoutput))


tmp8200 := Call2(__e, SymFn(sympr), tmp8198, tmp8199)


_ = tmp8200

return TailCall1(__e, SymFn(symnl), MakeInteger(0))


panic("unreachable")
})

tmp8201 := Call2(__e, ns2_1set, symshen_4toplevel_1display_1exception, tmp8197)


_ = tmp8201

tmp8202 := MakeNative0(func(__e *ControlFlow) Obj {
tmp8203 := Call0(__e, SymFn(symstoutput))


tmp8204 := Call2(__e, SymFn(sympr), MakeString("\nShen, www.shenlanguage.org, copyright (C) 2010-2026, Mark Tarver\n"), tmp8203)


_ = tmp8204

tmp8205 := GValue(__e, sym_dversion_d)

tmp8206 := GValue(__e, sym_dlanguage_d)

tmp8207 := GValue(__e, sym_dimplementation_d)

tmp8208 := GValue(__e, sym_drelease_d)

tmp8209 := Call3(__e, SymFn(symshen_4app), tmp8208, MakeString("\n"), symshen_4a)


tmp8211 := Call3(__e, SymFn(symshen_4app), tmp8207, PrimGuard2(__e, symcn, PrimStringConcat, MakeString(" "), tmp8209), symshen_4a)


tmp8213 := Call3(__e, SymFn(symshen_4app), tmp8206, PrimGuard2(__e, symcn, PrimStringConcat, MakeString(", platform: "), tmp8211), symshen_4a)


tmp8215 := Call3(__e, SymFn(symshen_4app), tmp8205, PrimGuard2(__e, symcn, PrimStringConcat, MakeString(", language: "), tmp8213), symshen_4a)


tmp8216 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("version: S"), tmp8215)

tmp8217 := Call0(__e, SymFn(symstoutput))


tmp8218 := Call2(__e, SymFn(sympr), tmp8216, tmp8217)


_ = tmp8218

tmp8219 := GValue(__e, sym_dport_d)

tmp8220 := GValue(__e, sym_dporters_d)

tmp8221 := Call3(__e, SymFn(symshen_4app), tmp8220, MakeString("\n\n"), symshen_4a)


tmp8223 := Call3(__e, SymFn(symshen_4app), tmp8219, PrimGuard2(__e, symcn, PrimStringConcat, MakeString(", ported by "), tmp8221), symshen_4a)


tmp8224 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("port "), tmp8223)

tmp8225 := Call0(__e, SymFn(symstoutput))


return TailCall2(__e, SymFn(sympr), tmp8224, tmp8225)


panic("unreachable")
})

tmp8226 := Call2(__e, ns2_1set, symshen_4credits, tmp8202)


_ = tmp8226

tmp8227 := MakeNative0(func(__e *ControlFlow) Obj {
tmp8228 := PrimGuard2(__e, symset, PrimSet, symshen_4_dcall_d, MakeInteger(0))

_ = tmp8228

return PrimGuard2(__e, symset, PrimSet, symshen_4_dinfs_d, MakeInteger(0))


panic("unreachable")
})

tmp8229 := Call2(__e, ns2_1set, symshen_4initialise__environment, tmp8227)


_ = tmp8229

tmp8230 := MakeNative0(func(__e *ControlFlow) Obj {
tmp8242 := GValue(__e, symshen_4_dtc_d)

if True == tmp8242 {
tmp8231 := GValue(__e, symshen_4_dhistory_d)

tmp8232 := Call1(__e, SymFn(symlength), tmp8231)


tmp8233 := Call3(__e, SymFn(symshen_4app), tmp8232, MakeString("+) "), symshen_4a)


tmp8234 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("\n("), tmp8233)

tmp8235 := Call0(__e, SymFn(symstoutput))


return TailCall2(__e, SymFn(sympr), tmp8234, tmp8235)


} else {
tmp8236 := GValue(__e, symshen_4_dhistory_d)

tmp8237 := Call1(__e, SymFn(symlength), tmp8236)


tmp8238 := Call3(__e, SymFn(symshen_4app), tmp8237, MakeString("-) "), symshen_4a)


tmp8239 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("\n("), tmp8238)

tmp8240 := Call0(__e, SymFn(symstoutput))


return TailCall2(__e, SymFn(sympr), tmp8239, tmp8240)


}


panic("unreachable")
})

tmp8243 := Call2(__e, ns2_1set, symshen_4prompt, tmp8230)


_ = tmp8243

tmp8244 := MakeNative0(func(__e *ControlFlow) Obj {
tmp8245 := GValue(__e, symshen_4_dpackage_d)

let__8171 := tmp8245
_ = let__8171

tmp8246 := Call0(__e, SymFn(symstinput))


tmp8247 := Call1(__e, SymFn(symlineread), tmp8246)


tmp8248 := Call2(__e, SymFn(symshen_4package_1user_1input), let__8171, tmp8247)


let__8172 := tmp8248
_ = let__8172

tmp8249 := Call0(__e, SymFn(symshen_4update_1history))


let__8173 := tmp8249
_ = let__8173

tmp8250 := GValue(__e, symshen_4_dtc_d)

return TailCall3(__e, SymFn(symshen_4evaluate_1lineread), let__8172, let__8173, tmp8250)


panic("unreachable")
})

tmp8251 := Call2(__e, ns2_1set, symshen_4read_1evaluate_1print, tmp8244)


_ = tmp8251

tmp8252 := MakeNative2(func(__e *ControlFlow, V5753, V5754 Obj) Obj {
tmp8257 := GEq(__e, symnull, V5753)

if True == tmp8257 {
return V5754
} else {
tmp8253 := PrimGuard1(__e, symstr, PrimStr, V5753)

let__8174 := tmp8253
_ = let__8174

tmp8254 := Call1(__e, SymFn(symexternal), V5753)


let__8175 := tmp8254
_ = let__8175

tmp8255 := MakeNative1(func(__e *ControlFlow, Z5757 Obj) Obj {
return TailCall3(__e, SymFn(symshen_4pui_1h), let__8174, let__8175, Z5757)
panic("unreachable")
})

return TailCall2(__e, SymFn(symmap), tmp8255, V5754)


}


panic("unreachable")
})

tmp8258 := Call2(__e, ns2_1set, symshen_4package_1user_1input, tmp8252)


_ = tmp8258

tmp8259 := MakeNative3(func(__e *ControlFlow, V5762, V5763, V5764 Obj) Obj {
tmp8300 := GConsP(__e, V5764)

var ifres8287 Obj

if True == tmp8300 {
tmp8298 := GHd(__e, V5764)

tmp8299 := GEq(__e, symfn, tmp8298)

var ifres8289 Obj

if True == tmp8299 {
tmp8296 := GTl(__e, V5764)

tmp8297 := GConsP(__e, tmp8296)

var ifres8291 Obj

if True == tmp8297 {
tmp8293 := GTl(__e, V5764)

tmp8294 := GTl(__e, tmp8293)

tmp8295 := GEq(__e, Nil, tmp8294)

var ifres8292 Obj

if True == tmp8295 {
ifres8292 = True


} else {
ifres8292 = False


}

ifres8291 = ifres8292


} else {
ifres8291 = False


}

var ifres8290 Obj

if True == ifres8291 {
ifres8290 = True


} else {
ifres8290 = False


}

ifres8289 = ifres8290


} else {
ifres8289 = False


}

var ifres8288 Obj

if True == ifres8289 {
ifres8288 = True


} else {
ifres8288 = False


}

ifres8287 = ifres8288


} else {
ifres8287 = False


}

if True == ifres8287 {
tmp8265 := GTl(__e, V5764)

tmp8266 := GHd(__e, tmp8265)

tmp8267 := Call3(__e, SymFn(symshen_4internal_2), tmp8266, V5762, V5763)


if True == tmp8267 {
tmp8260 := GTl(__e, V5764)

tmp8261 := GHd(__e, tmp8260)

tmp8262 := Call2(__e, SymFn(symshen_4intern_1in_1package), V5762, tmp8261)


tmp8263 := GCons(__e, tmp8262, Nil)

return GCons(__e, symfn, tmp8263)


} else {
return V5764
}


} else {
tmp8285 := GConsP(__e, V5764)

if True == tmp8285 {
tmp8282 := GHd(__e, V5764)

tmp8283 := Call3(__e, SymFn(symshen_4internal_2), tmp8282, V5762, V5763)


if True == tmp8283 {
tmp8268 := GHd(__e, V5764)

tmp8269 := Call2(__e, SymFn(symshen_4intern_1in_1package), V5762, tmp8268)


tmp8270 := MakeNative1(func(__e *ControlFlow, Z5765 Obj) Obj {
return TailCall3(__e, SymFn(symshen_4pui_1h), V5762, V5763, Z5765)
panic("unreachable")
})

tmp8271 := GTl(__e, V5764)

tmp8272 := Call2(__e, SymFn(symmap), tmp8270, tmp8271)


return GCons(__e, tmp8269, tmp8272)


} else {
tmp8279 := GHd(__e, V5764)

tmp8280 := GConsP(__e, tmp8279)

if True == tmp8280 {
tmp8273 := MakeNative1(func(__e *ControlFlow, Z5766 Obj) Obj {
return TailCall3(__e, SymFn(symshen_4pui_1h), V5762, V5763, Z5766)
panic("unreachable")
})

return TailCall2(__e, SymFn(symmap), tmp8273, V5764)


} else {
tmp8274 := GHd(__e, V5764)

tmp8275 := MakeNative1(func(__e *ControlFlow, Z5767 Obj) Obj {
return TailCall3(__e, SymFn(symshen_4pui_1h), V5762, V5763, Z5767)
panic("unreachable")
})

tmp8276 := GTl(__e, V5764)

tmp8277 := Call2(__e, SymFn(symmap), tmp8275, tmp8276)


return GCons(__e, tmp8274, tmp8277)


}


}


} else {
return V5764
}


}


panic("unreachable")
})

tmp8301 := Call2(__e, ns2_1set, symshen_4pui_1h, tmp8259)


_ = tmp8301

tmp8302 := MakeNative0(func(__e *ControlFlow) Obj {
tmp8303 := Call0(__e, SymFn(symit))


tmp8304 := Call1(__e, SymFn(symshen_4trim_1it), tmp8303)


tmp8305 := GValue(__e, symshen_4_dhistory_d)

tmp8306 := GCons(__e, tmp8304, tmp8305)

return PrimGuard2(__e, symset, PrimSet, symshen_4_dhistory_d, tmp8306)


panic("unreachable")
})

tmp8307 := Call2(__e, ns2_1set, symshen_4update_1history, tmp8302)


_ = tmp8307

tmp8308 := (func() Obj {
var __self Obj
__self = MakeNative1(func(__e *ControlFlow, __self1 Obj) Obj {
__selftop:
V5768 := __self1
_ = V5768
tmp8316 := Call1(__e, SymFn(symshen_4_7string_2), V5768)


var ifres8311 Obj

if True == tmp8316 {
tmp8313 := Call1(__e, SymFn(symhdstr), V5768)


tmp8314 := PrimGuard1(__e, symstring_1_6n, PrimStringToNumber, tmp8313)

tmp8315 := Call1(__e, SymFn(symshen_4whitespace_2), tmp8314)


var ifres8312 Obj

if True == tmp8315 {
ifres8312 = True


} else {
ifres8312 = False


}

ifres8311 = ifres8312


} else {
ifres8311 = False


}

if True == ifres8311 {
if SymFn(symshen_4trim_1it) == __self {
__self1 = PrimGuard1(__e, symtlstr, PrimTailString, V5768)
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4trim_1it), PrimGuard1(__e, symtlstr, PrimTailString, V5768))


} else {
return V5768
}


panic("unreachable")
})
return __self
})()

tmp8317 := Call2(__e, ns2_1set, symshen_4trim_1it, tmp8308)


_ = tmp8317

tmp8318 := (func() Obj {
var __self Obj
__self = MakeNative3(func(__e *ControlFlow, __self1, __self2, __self3 Obj) Obj {
__selftop:
V5787 := __self1
_ = V5787
V5788 := __self2
_ = V5788
V5789 := __self3
_ = V5789
tmp8437 := GConsP(__e, V5787)

var ifres8406 Obj

if True == tmp8437 {
tmp8435 := GTl(__e, V5787)

tmp8436 := GEq(__e, Nil, tmp8435)

var ifres8408 Obj

if True == tmp8436 {
tmp8434 := GConsP(__e, V5788)

var ifres8410 Obj

if True == tmp8434 {
tmp8432 := GHd(__e, V5788)

tmp8433 := Call1(__e, SymFn(symshen_4_7string_2), tmp8432)


var ifres8412 Obj

if True == tmp8433 {
tmp8429 := GHd(__e, V5788)

tmp8430 := Call1(__e, SymFn(symhdstr), tmp8429)


tmp8431 := GEq(__e, MakeString("!"), tmp8430)

var ifres8414 Obj

if True == tmp8431 {
tmp8426 := GHd(__e, V5788)

tmp8428 := Call1(__e, SymFn(symshen_4_7string_2), PrimGuard1(__e, symtlstr, PrimTailString, tmp8426))


var ifres8416 Obj

if True == tmp8428 {
tmp8422 := GHd(__e, V5788)

tmp8424 := Call1(__e, SymFn(symhdstr), PrimGuard1(__e, symtlstr, PrimTailString, tmp8422))


tmp8425 := GEq(__e, MakeString("!"), tmp8424)

var ifres8418 Obj

if True == tmp8425 {
tmp8420 := GTl(__e, V5788)

tmp8421 := GConsP(__e, tmp8420)

var ifres8419 Obj

if True == tmp8421 {
ifres8419 = True


} else {
ifres8419 = False


}

ifres8418 = ifres8419


} else {
ifres8418 = False


}

var ifres8417 Obj

if True == ifres8418 {
ifres8417 = True


} else {
ifres8417 = False


}

ifres8416 = ifres8417


} else {
ifres8416 = False


}

var ifres8415 Obj

if True == ifres8416 {
ifres8415 = True


} else {
ifres8415 = False


}

ifres8414 = ifres8415


} else {
ifres8414 = False


}

var ifres8413 Obj

if True == ifres8414 {
ifres8413 = True


} else {
ifres8413 = False


}

ifres8412 = ifres8413


} else {
ifres8412 = False


}

var ifres8411 Obj

if True == ifres8412 {
ifres8411 = True


} else {
ifres8411 = False


}

ifres8410 = ifres8411


} else {
ifres8410 = False


}

var ifres8409 Obj

if True == ifres8410 {
ifres8409 = True


} else {
ifres8409 = False


}

ifres8408 = ifres8409


} else {
ifres8408 = False


}

var ifres8407 Obj

if True == ifres8408 {
ifres8407 = True


} else {
ifres8407 = False


}

ifres8406 = ifres8407


} else {
ifres8406 = False


}

if True == ifres8406 {
tmp8319 := GTl(__e, V5788)

tmp8320 := GHd(__e, tmp8319)

tmp8321 := Call1(__e, SymFn(symread_1from_1string), tmp8320)


let__8176 := tmp8321
_ = let__8176

tmp8322 := GTl(__e, V5788)

tmp8323 := GHd(__e, tmp8322)

tmp8324 := GTl(__e, V5788)

tmp8325 := GCons(__e, tmp8323, tmp8324)

tmp8326 := PrimGuard2(__e, symset, PrimSet, symshen_4_dhistory_d, tmp8325)

let__8177 := tmp8326
_ = let__8177

tmp8327 := GTl(__e, V5788)

tmp8328 := GHd(__e, tmp8327)

tmp8329 := Call3(__e, SymFn(symshen_4app), tmp8328, MakeString("\n"), symshen_4a)


tmp8330 := Call0(__e, SymFn(symstoutput))


tmp8331 := Call2(__e, SymFn(sympr), tmp8329, tmp8330)


let__8178 := tmp8331
_ = let__8178

if SymFn(symshen_4evaluate_1lineread) == __self {
__self1, __self2, __self3 = let__8176, let__8177, V5789
__e.Tick()
goto __selftop
}
return TailCall3(__e, SymFn(symshen_4evaluate_1lineread), let__8176, let__8177, V5789)


} else {
tmp8404 := GConsP(__e, V5787)

var ifres8388 Obj

if True == tmp8404 {
tmp8402 := GTl(__e, V5787)

tmp8403 := GEq(__e, Nil, tmp8402)

var ifres8390 Obj

if True == tmp8403 {
tmp8401 := GConsP(__e, V5788)

var ifres8392 Obj

if True == tmp8401 {
tmp8399 := GHd(__e, V5788)

tmp8400 := Call1(__e, SymFn(symshen_4_7string_2), tmp8399)


var ifres8394 Obj

if True == tmp8400 {
tmp8396 := GHd(__e, V5788)

tmp8397 := Call1(__e, SymFn(symhdstr), tmp8396)


tmp8398 := GEq(__e, MakeString("!"), tmp8397)

var ifres8395 Obj

if True == tmp8398 {
ifres8395 = True


} else {
ifres8395 = False


}

ifres8394 = ifres8395


} else {
ifres8394 = False


}

var ifres8393 Obj

if True == ifres8394 {
ifres8393 = True


} else {
ifres8393 = False


}

ifres8392 = ifres8393


} else {
ifres8392 = False


}

var ifres8391 Obj

if True == ifres8392 {
ifres8391 = True


} else {
ifres8391 = False


}

ifres8390 = ifres8391


} else {
ifres8390 = False


}

var ifres8389 Obj

if True == ifres8390 {
ifres8389 = True


} else {
ifres8389 = False


}

ifres8388 = ifres8389


} else {
ifres8388 = False


}

if True == ifres8388 {
tmp8337 := GHd(__e, V5788)

tmp8339 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(PrimGuard1(__e, symtlstr, PrimTailString, tmp8337), MakeString(""))
}
__typedArg0 := PrimGuard1(__e, symtlstr, PrimTailString, tmp8337)
__typedArg1 := MakeString("")
return Call2(__e, SymFn(sym_a), __typedArg0, __typedArg1)
})()

var ifres8332 Obj

if True == tmp8339 {
ifres8332 = Nil


} else {
tmp8333 := GHd(__e, V5788)

tmp8335 := Call1(__e, SymFn(symread_1from_1string), PrimGuard1(__e, symtlstr, PrimTailString, tmp8333))


tmp8336 := GHd(__e, tmp8335)

ifres8332 = tmp8336


}

let__8179 := ifres8332
_ = let__8179

tmp8340 := GHd(__e, V5788)

tmp8341 := PrimGuard1(__e, symtlstr, PrimTailString, tmp8340)

tmp8342 := GTl(__e, V5788)

tmp8343 := Call3(__e, SymFn(symshen_4use_1history), let__8179, tmp8341, tmp8342)


let__8180 := tmp8343
_ = let__8180

tmp8344 := Call3(__e, SymFn(symshen_4app), let__8180, MakeString("\n"), symshen_4a)


tmp8345 := Call0(__e, SymFn(symstoutput))


tmp8346 := Call2(__e, SymFn(sympr), tmp8344, tmp8345)


let__8181 := tmp8346
_ = let__8181

tmp8347 := Call1(__e, SymFn(symread_1from_1string), let__8180)


let__8182 := tmp8347
_ = let__8182

tmp8348 := GTl(__e, V5788)

tmp8349 := GCons(__e, let__8180, tmp8348)

tmp8350 := PrimGuard2(__e, symset, PrimSet, symshen_4_dhistory_d, tmp8349)

let__8183 := tmp8350
_ = let__8183

if SymFn(symshen_4evaluate_1lineread) == __self {
__self1, __self2, __self3 = let__8182, let__8183, V5789
__e.Tick()
goto __selftop
}
return TailCall3(__e, SymFn(symshen_4evaluate_1lineread), let__8182, let__8183, V5789)


} else {
tmp8386 := GConsP(__e, V5787)

var ifres8370 Obj

if True == tmp8386 {
tmp8384 := GTl(__e, V5787)

tmp8385 := GEq(__e, Nil, tmp8384)

var ifres8372 Obj

if True == tmp8385 {
tmp8383 := GConsP(__e, V5788)

var ifres8374 Obj

if True == tmp8383 {
tmp8381 := GHd(__e, V5788)

tmp8382 := Call1(__e, SymFn(symshen_4_7string_2), tmp8381)


var ifres8376 Obj

if True == tmp8382 {
tmp8378 := GHd(__e, V5788)

tmp8379 := Call1(__e, SymFn(symhdstr), tmp8378)


tmp8380 := GEq(__e, MakeString("%"), tmp8379)

var ifres8377 Obj

if True == tmp8380 {
ifres8377 = True


} else {
ifres8377 = False


}

ifres8376 = ifres8377


} else {
ifres8376 = False


}

var ifres8375 Obj

if True == ifres8376 {
ifres8375 = True


} else {
ifres8375 = False


}

ifres8374 = ifres8375


} else {
ifres8374 = False


}

var ifres8373 Obj

if True == ifres8374 {
ifres8373 = True


} else {
ifres8373 = False


}

ifres8372 = ifres8373


} else {
ifres8372 = False


}

var ifres8371 Obj

if True == ifres8372 {
ifres8371 = True


} else {
ifres8371 = False


}

ifres8370 = ifres8371


} else {
ifres8370 = False


}

if True == ifres8370 {
tmp8356 := GHd(__e, V5788)

tmp8358 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(PrimGuard1(__e, symtlstr, PrimTailString, tmp8356), MakeString(""))
}
__typedArg0 := PrimGuard1(__e, symtlstr, PrimTailString, tmp8356)
__typedArg1 := MakeString("")
return Call2(__e, SymFn(sym_a), __typedArg0, __typedArg1)
})()

var ifres8351 Obj

if True == tmp8358 {
ifres8351 = Nil


} else {
tmp8352 := GHd(__e, V5788)

tmp8354 := Call1(__e, SymFn(symread_1from_1string), PrimGuard1(__e, symtlstr, PrimTailString, tmp8352))


tmp8355 := GHd(__e, tmp8354)

ifres8351 = tmp8355


}

let__8184 := ifres8351
_ = let__8184

tmp8359 := GHd(__e, V5788)

tmp8360 := PrimGuard1(__e, symtlstr, PrimTailString, tmp8359)

tmp8361 := GTl(__e, V5788)

tmp8362 := Call3(__e, SymFn(symshen_4peek_1history), let__8184, tmp8360, tmp8361)


let__8185 := tmp8362
_ = let__8185

tmp8363 := GTl(__e, V5788)

tmp8364 := PrimGuard2(__e, symset, PrimSet, symshen_4_dhistory_d, tmp8363)

let__8186 := tmp8364
_ = let__8186

return TailCall0(__e, SymFn(symabort))


} else {
tmp8368 := GEq(__e, True, V5789)

if True == tmp8368 {
return TailCall1(__e, SymFn(symshen_4check_1eval_1and_1print), V5787)
} else {
tmp8366 := GEq(__e, False, V5789)

if True == tmp8366 {
return TailCall1(__e, SymFn(symshen_4eval_1and_1print), V5787)
} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.evaluate-lineread"))
}


}


}


}


}


panic("unreachable")
})
return __self
})()

tmp8438 := Call2(__e, ns2_1set, symshen_4evaluate_1lineread, tmp8318)


_ = tmp8438

tmp8439 := MakeNative3(func(__e *ControlFlow, V5801, V5802, V5803 Obj) Obj {
tmp8445 := PrimGuard1(__e, syminteger_2, PrimIsInteger, V5801)

if True == tmp8445 {
tmp8440 := GAdd(__e, MakeInteger(1), V5801)

tmp8441 := Call1(__e, SymFn(symreverse), V5803)


return TailCall2(__e, SymFn(symnth), tmp8440, tmp8441)


} else {
tmp8443 := PrimGuard1(__e, symsymbol_2, PrimIsSymbol, V5801)

if True == tmp8443 {
return TailCall2(__e, SymFn(symshen_4string_1match), V5802, V5803)
} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("! expects a number or a symbol\n"))
}


}


panic("unreachable")
})

tmp8446 := Call2(__e, ns2_1set, symshen_4use_1history, tmp8439)


_ = tmp8446

tmp8447 := MakeNative3(func(__e *ControlFlow, V5804, V5805, V5806 Obj) Obj {
tmp8461 := PrimGuard1(__e, syminteger_2, PrimIsInteger, V5804)

if True == tmp8461 {
tmp8448 := GAdd(__e, MakeInteger(1), V5804)

tmp8449 := Call1(__e, SymFn(symreverse), V5806)


tmp8450 := Call2(__e, SymFn(symnth), tmp8448, tmp8449)


tmp8451 := Call3(__e, SymFn(symshen_4app), tmp8450, MakeString(""), symshen_4a)


tmp8452 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("\n"), tmp8451)

tmp8453 := Call0(__e, SymFn(symstoutput))


return TailCall2(__e, SymFn(sympr), tmp8452, tmp8453)


} else {
tmp8459 := GEq(__e, V5805, MakeString(""))

var ifres8456 Obj

if True == tmp8459 {
ifres8456 = True


} else {
tmp8458 := PrimGuard1(__e, symsymbol_2, PrimIsSymbol, V5804)

var ifres8457 Obj

if True == tmp8458 {
ifres8457 = True


} else {
ifres8457 = False


}

ifres8456 = ifres8457


}

if True == ifres8456 {
tmp8454 := Call1(__e, SymFn(symreverse), V5806)


return TailCall3(__e, SymFn(symshen_4recursive_1string_1match), MakeInteger(0), V5805, tmp8454)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("% expects a number or a symbol\n"))
}


}


panic("unreachable")
})

tmp8462 := Call2(__e, ns2_1set, symshen_4peek_1history, tmp8447)


_ = tmp8462

tmp8463 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V5816 := __self1
_ = V5816
V5817 := __self2
_ = V5817
tmp8474 := GEq(__e, Nil, V5817)

if True == tmp8474 {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("\ninput not found"))
} else {
tmp8472 := GConsP(__e, V5817)

var ifres8468 Obj

if True == tmp8472 {
tmp8470 := GHd(__e, V5817)

tmp8471 := Call2(__e, SymFn(symshen_4string_1prefix_2), V5816, tmp8470)


var ifres8469 Obj

if True == tmp8471 {
ifres8469 = True


} else {
ifres8469 = False


}

ifres8468 = ifres8469


} else {
ifres8468 = False


}

if True == ifres8468 {
return GHd(__e, V5817)
} else {
tmp8466 := GConsP(__e, V5817)

if True == tmp8466 {
tmp8464 := GTl(__e, V5817)

if SymFn(symshen_4string_1match) == __self {
__self1, __self2 = V5816, tmp8464
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4string_1match), V5816, tmp8464)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.string-match"))
}


}


}


panic("unreachable")
})
return __self
})()

tmp8475 := Call2(__e, ns2_1set, symshen_4string_1match, tmp8463)


_ = tmp8475

tmp8476 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V5825 := __self1
_ = V5825
V5826 := __self2
_ = V5826
tmp8513 := GEq(__e, MakeString(""), V5825)

if True == tmp8513 {
return True
} else {
tmp8511 := Call1(__e, SymFn(symshen_4_7string_2), V5825)


var ifres8506 Obj

if True == tmp8511 {
tmp8508 := Call1(__e, SymFn(symhdstr), V5825)


tmp8509 := PrimGuard1(__e, symstring_1_6n, PrimStringToNumber, tmp8508)

tmp8510 := Call1(__e, SymFn(symshen_4whitespace_2), tmp8509)


var ifres8507 Obj

if True == tmp8510 {
ifres8507 = True


} else {
ifres8507 = False


}

ifres8506 = ifres8507


} else {
ifres8506 = False


}

if True == ifres8506 {
if SymFn(symshen_4string_1prefix_2) == __self {
__self1, __self2 = PrimGuard1(__e, symtlstr, PrimTailString, V5825), V5826
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4string_1prefix_2), PrimGuard1(__e, symtlstr, PrimTailString, V5825), V5826)


} else {
tmp8504 := Call1(__e, SymFn(symshen_4_7string_2), V5826)


var ifres8499 Obj

if True == tmp8504 {
tmp8501 := Call1(__e, SymFn(symhdstr), V5826)


tmp8502 := PrimGuard1(__e, symstring_1_6n, PrimStringToNumber, tmp8501)

tmp8503 := Call1(__e, SymFn(symshen_4whitespace_2), tmp8502)


var ifres8500 Obj

if True == tmp8503 {
ifres8500 = True


} else {
ifres8500 = False


}

ifres8499 = ifres8500


} else {
ifres8499 = False


}

if True == ifres8499 {
if SymFn(symshen_4string_1prefix_2) == __self {
__self1, __self2 = V5825, PrimGuard1(__e, symtlstr, PrimTailString, V5826)
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4string_1prefix_2), V5825, PrimGuard1(__e, symtlstr, PrimTailString, V5826))


} else {
tmp8497 := Call1(__e, SymFn(symshen_4_7string_2), V5826)


var ifres8493 Obj

if True == tmp8497 {
tmp8495 := Call1(__e, SymFn(symhdstr), V5826)


tmp8496 := GEq(__e, MakeString("("), tmp8495)

var ifres8494 Obj

if True == tmp8496 {
ifres8494 = True


} else {
ifres8494 = False


}

ifres8493 = ifres8494


} else {
ifres8493 = False


}

if True == ifres8493 {
if SymFn(symshen_4string_1prefix_2) == __self {
__self1, __self2 = V5825, PrimGuard1(__e, symtlstr, PrimTailString, V5826)
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4string_1prefix_2), V5825, PrimGuard1(__e, symtlstr, PrimTailString, V5826))


} else {
tmp8491 := Call1(__e, SymFn(symshen_4_7string_2), V5825)


var ifres8483 Obj

if True == tmp8491 {
tmp8490 := Call1(__e, SymFn(symshen_4_7string_2), V5826)


var ifres8485 Obj

if True == tmp8490 {
tmp8487 := Call1(__e, SymFn(symhdstr), V5825)


tmp8488 := Call1(__e, SymFn(symhdstr), V5826)


tmp8489 := GEq(__e, tmp8487, tmp8488)

var ifres8486 Obj

if True == tmp8489 {
ifres8486 = True


} else {
ifres8486 = False


}

ifres8485 = ifres8486


} else {
ifres8485 = False


}

var ifres8484 Obj

if True == ifres8485 {
ifres8484 = True


} else {
ifres8484 = False


}

ifres8483 = ifres8484


} else {
ifres8483 = False


}

if True == ifres8483 {
tmp8480 := PrimGuard1(__e, symtlstr, PrimTailString, V5825)

if SymFn(symshen_4string_1prefix_2) == __self {
__self1, __self2 = tmp8480, PrimGuard1(__e, symtlstr, PrimTailString, V5826)
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4string_1prefix_2), tmp8480, PrimGuard1(__e, symtlstr, PrimTailString, V5826))


} else {
return False
}


}


}


}


}


panic("unreachable")
})
return __self
})()

tmp8514 := Call2(__e, ns2_1set, symshen_4string_1prefix_2, tmp8476)


_ = tmp8514

tmp8515 := (func() Obj {
var __self Obj
__self = MakeNative3(func(__e *ControlFlow, __self1, __self2, __self3 Obj) Obj {
__selftop:
V5837 := __self1
_ = V5837
V5838 := __self2
_ = V5838
V5839 := __self3
_ = V5839
tmp8530 := GEq(__e, Nil, V5839)

if True == tmp8530 {
return symshen_4skip
} else {
tmp8528 := GConsP(__e, V5839)

if True == tmp8528 {
tmp8523 := GHd(__e, V5839)

tmp8524 := Call2(__e, SymFn(symshen_4string_1prefix_2), V5838, tmp8523)


var ifres8516 Obj

if True == tmp8524 {
tmp8517 := GHd(__e, V5839)

tmp8518 := Call3(__e, SymFn(symshen_4app), tmp8517, MakeString("\n"), symshen_4a)


tmp8520 := Call3(__e, SymFn(symshen_4app), V5837, PrimGuard2(__e, symcn, PrimStringConcat, MakeString(". "), tmp8518), symshen_4a)


tmp8521 := Call0(__e, SymFn(symstoutput))


tmp8522 := Call2(__e, SymFn(sympr), tmp8520, tmp8521)


ifres8516 = tmp8522


} else {
ifres8516 = symshen_4skip


}

_ = ifres8516

tmp8525 := GAdd(__e, V5837, MakeInteger(1))

tmp8526 := GTl(__e, V5839)

if SymFn(symshen_4recursive_1string_1match) == __self {
__self1, __self2, __self3 = tmp8525, V5838, tmp8526
__e.Tick()
goto __selftop
}
return TailCall3(__e, SymFn(symshen_4recursive_1string_1match), tmp8525, V5838, tmp8526)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.recursive-string-match"))
}


}


panic("unreachable")
})
return __self
})()

return TailCall2(__e, ns2_1set, symshen_4recursive_1string_1match, tmp8515)




panic("unreachable")
})

