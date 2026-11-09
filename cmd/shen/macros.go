package main

import . "github.com/pyrex41/shen-go/kl"

var MacrosMain = MakeNative0(func(__e *ControlFlow) Obj {
tmp8174 := MakeNative1(func(__e *ControlFlow, V7150 Obj) Obj {
tmp8175 := MakeNative1(func(__e *ControlFlow, Z7152 Obj) Obj {
return GTl(__e, Z7152)
panic("unreachable")
})

tmp8176 := GValue(__e, sym_dmacros_d)

tmp8177 := Call2(__e, SymFn(symmap), tmp8175, tmp8176)


let__8083 := tmp8177
_ = let__8083

tmp8178 := GValue(__e, symshen_4_dsource_1form_1handlers_d)

return TailCall4(__e, SymFn(symshen_4macroexpand_1h), V7150, let__8083, let__8083, tmp8178)


panic("unreachable")
})

tmp8179 := Call2(__e, ns2_1set, symmacroexpand, tmp8174)


_ = tmp8179

tmp8180 := (func() Obj {
var __self Obj
__self = MakeNative4(func(__e *ControlFlow, __self1, __self2, __self3, __self4 Obj) Obj {
__selftop:
V7165 := __self1
_ = V7165
V7166 := __self2
_ = V7166
V7167 := __self3
_ = V7167
V7168 := __self4
_ = V7168
tmp8189 := GEq(__e, Nil, V7166)

if True == tmp8189 {
return V7165
} else {
tmp8187 := GConsP(__e, V7166)

if True == tmp8187 {
tmp8181 := GHd(__e, V7166)

tmp8182 := Call3(__e, SymFn(symshen_4source_1macro_1walk), tmp8181, V7165, V7168)


let__8084 := tmp8182
_ = let__8084

tmp8185 := GEq(__e, V7165, let__8084)

if True == tmp8185 {
tmp8183 := GTl(__e, V7166)

if SymFn(symshen_4macroexpand_1h) == __self {
__self1, __self2, __self3, __self4 = V7165, tmp8183, V7167, V7168
__e.Tick()
goto __selftop
}
return TailCall4(__e, SymFn(symshen_4macroexpand_1h), V7165, tmp8183, V7167, V7168)


} else {
if SymFn(symshen_4macroexpand_1h) == __self {
__self1, __self2, __self3, __self4 = let__8084, V7167, V7167, V7168
__e.Tick()
goto __selftop
}
return TailCall4(__e, SymFn(symshen_4macroexpand_1h), let__8084, V7167, V7167, V7168)
}


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.macroexpand-h"))
}


}


panic("unreachable")
})
return __self
})()

tmp8190 := Call2(__e, ns2_1set, symshen_4macroexpand_1h, tmp8180)


_ = tmp8190

tmp8191 := MakeNative3(func(__e *ControlFlow, V7173, V7174, V7175 Obj) Obj {
tmp8200 := GEq(__e, Nil, V7175)

if True == tmp8200 {
return TailCall2(__e, SymFn(symshen_4walk), V7173, V7174)
} else {
tmp8197 := Call2(__e, SymFn(symshen_4source_1form_1handler), V7174, V7175)


tmp8198 := GConsP(__e, tmp8197)

if True == tmp8198 {
return V7174
} else {
tmp8195 := GConsP(__e, V7174)

if True == tmp8195 {
tmp8192 := MakeNative1(func(__e *ControlFlow, Z7176 Obj) Obj {
return TailCall3(__e, SymFn(symshen_4source_1macro_1walk), V7173, Z7176, V7175)
panic("unreachable")
})

tmp8193 := Call2(__e, SymFn(symmap), tmp8192, V7174)


return TailCall1(__e, V7173, tmp8193)


} else {
return TailCall1(__e, V7173, V7174)
}


}


}


panic("unreachable")
})

tmp8201 := Call2(__e, ns2_1set, symshen_4source_1macro_1walk, tmp8191)


_ = tmp8201

tmp8202 := MakeNative2(func(__e *ControlFlow, V7177, V7178 Obj) Obj {
tmp8206 := GConsP(__e, V7178)

if True == tmp8206 {
tmp8203 := MakeNative1(func(__e *ControlFlow, Z7179 Obj) Obj {
return TailCall2(__e, SymFn(symshen_4walk), V7177, Z7179)
panic("unreachable")
})

tmp8204 := Call2(__e, SymFn(symmap), tmp8203, V7178)


return TailCall1(__e, V7177, tmp8204)


} else {
return TailCall1(__e, V7177, V7178)
}


panic("unreachable")
})

tmp8207 := Call2(__e, ns2_1set, symshen_4walk, tmp8202)


_ = tmp8207

tmp8208 := MakeNative1(func(__e *ControlFlow, V7180 Obj) Obj {
tmp8209 := MakeNative0(func(__e *ControlFlow) Obj {
return V7180
panic("unreachable")
})

let__8085 := tmp8209
_ = let__8085

tmp8457 := GConsP(__e, V7180)

if True == tmp8457 {
tmp8210 := GHd(__e, V7180)

let__8086 := tmp8210
_ = let__8086

tmp8211 := GTl(__e, V7180)

let__8087 := tmp8211
_ = let__8087

tmp8455 := GEq(__e, symdefmacro, let__8086)

var ifres8453 Obj

if True == tmp8455 {
tmp8454 := GConsP(__e, let__8087)

ifres8453 = tmp8454


} else {
ifres8453 = False


}

if True == ifres8453 {
tmp8212 := GHd(__e, let__8087)

tmp8213 := GTl(__e, let__8087)

return TailCall2(__e, SymFn(symshen_4process_1def), tmp8212, tmp8213)


} else {
tmp8451 := GEq(__e, symdefcc, let__8086)

if True == tmp8451 {
return TailCall1(__e, SymFn(symshen_4yacc_1_6shen), let__8087)
} else {
tmp8449 := GEq(__e, symu_b, let__8086)

var ifres8444 Obj

if True == tmp8449 {
tmp8448 := GConsP(__e, let__8087)

var ifres8445 Obj

if True == tmp8448 {
tmp8446 := GTl(__e, let__8087)

tmp8447 := GEq(__e, Nil, tmp8446)

ifres8445 = tmp8447


} else {
ifres8445 = False


}

ifres8444 = ifres8445


} else {
ifres8444 = False


}

if True == ifres8444 {
tmp8214 := GHd(__e, let__8087)

tmp8215 := Call1(__e, SymFn(symshen_4make_1uppercase), tmp8214)


tmp8216 := GCons(__e, tmp8215, Nil)

return GCons(__e, symprotect, tmp8216)


} else {
tmp8442 := GEq(__e, symerror, let__8086)

var ifres8440 Obj

if True == tmp8442 {
tmp8441 := GConsP(__e, let__8087)

ifres8440 = tmp8441


} else {
ifres8440 = False


}

if True == ifres8440 {
tmp8217 := GHd(__e, let__8087)

tmp8218 := GTl(__e, let__8087)

tmp8219 := Call2(__e, SymFn(symshen_4mkstr), tmp8217, tmp8218)


tmp8220 := GCons(__e, tmp8219, Nil)

return GCons(__e, symsimple_1error, tmp8220)


} else {
tmp8438 := GEq(__e, symoutput, let__8086)

var ifres8436 Obj

if True == tmp8438 {
tmp8437 := GConsP(__e, let__8087)

ifres8436 = tmp8437


} else {
ifres8436 = False


}

if True == ifres8436 {
tmp8221 := GHd(__e, let__8087)

tmp8222 := GTl(__e, let__8087)

tmp8223 := Call2(__e, SymFn(symshen_4mkstr), tmp8221, tmp8222)


tmp8224 := GCons(__e, symstoutput, Nil)

tmp8225 := GCons(__e, tmp8224, Nil)

tmp8226 := GCons(__e, tmp8223, tmp8225)

return GCons(__e, sympr, tmp8226)


} else {
tmp8434 := GEq(__e, sympr, let__8086)

var ifres8429 Obj

if True == tmp8434 {
tmp8433 := GConsP(__e, let__8087)

var ifres8430 Obj

if True == tmp8433 {
tmp8431 := GTl(__e, let__8087)

tmp8432 := GEq(__e, Nil, tmp8431)

ifres8430 = tmp8432


} else {
ifres8430 = False


}

ifres8429 = ifres8430


} else {
ifres8429 = False


}

if True == ifres8429 {
tmp8227 := GHd(__e, let__8087)

tmp8228 := GCons(__e, symstoutput, Nil)

tmp8229 := GCons(__e, tmp8228, Nil)

tmp8230 := GCons(__e, tmp8227, tmp8229)

return GCons(__e, sympr, tmp8230)


} else {
tmp8427 := GEq(__e, symmake_1string, let__8086)

var ifres8425 Obj

if True == tmp8427 {
tmp8426 := GConsP(__e, let__8087)

ifres8425 = tmp8426


} else {
ifres8425 = False


}

if True == ifres8425 {
tmp8231 := GHd(__e, let__8087)

tmp8232 := GTl(__e, let__8087)

return TailCall2(__e, SymFn(symshen_4mkstr), tmp8231, tmp8232)


} else {
tmp8423 := GEq(__e, symlineread, let__8086)

var ifres8421 Obj

if True == tmp8423 {
tmp8422 := GEq(__e, Nil, let__8087)

ifres8421 = tmp8422


} else {
ifres8421 = False


}

if True == ifres8421 {
tmp8233 := GCons(__e, symstinput, Nil)

tmp8234 := GCons(__e, tmp8233, Nil)

return GCons(__e, symlineread, tmp8234)


} else {
tmp8419 := GEq(__e, syminput, let__8086)

var ifres8417 Obj

if True == tmp8419 {
tmp8418 := GEq(__e, Nil, let__8087)

ifres8417 = tmp8418


} else {
ifres8417 = False


}

if True == ifres8417 {
tmp8235 := GCons(__e, symstinput, Nil)

tmp8236 := GCons(__e, tmp8235, Nil)

return GCons(__e, syminput, tmp8236)


} else {
tmp8415 := GEq(__e, symread, let__8086)

var ifres8413 Obj

if True == tmp8415 {
tmp8414 := GEq(__e, Nil, let__8087)

ifres8413 = tmp8414


} else {
ifres8413 = False


}

if True == ifres8413 {
tmp8237 := GCons(__e, symstinput, Nil)

tmp8238 := GCons(__e, tmp8237, Nil)

return GCons(__e, symread, tmp8238)


} else {
tmp8411 := GEq(__e, syminput_7, let__8086)

var ifres8409 Obj

if True == tmp8411 {
tmp8410 := GConsP(__e, let__8087)

ifres8409 = tmp8410


} else {
ifres8409 = False


}

if True == ifres8409 {
return TailCall1(__e, SymFn(symshen_4process_1input_7), V7180)
} else {
tmp8407 := GEq(__e, symread_1byte, let__8086)

var ifres8405 Obj

if True == tmp8407 {
tmp8406 := GEq(__e, Nil, let__8087)

ifres8405 = tmp8406


} else {
ifres8405 = False


}

if True == ifres8405 {
return TailCall0(__e, SymFn(symshen_4process_1read_1byte))
} else {
tmp8403 := GEq(__e, symprolog_2, let__8086)

if True == tmp8403 {
return TailCall1(__e, SymFn(symshen_4call_1prolog), let__8087)
} else {
tmp8401 := GEq(__e, symdefprolog, let__8086)

var ifres8399 Obj

if True == tmp8401 {
tmp8400 := GConsP(__e, let__8087)

ifres8399 = tmp8400


} else {
ifres8399 = False


}

if True == ifres8399 {
tmp8239 := GHd(__e, let__8087)

tmp8240 := GTl(__e, let__8087)

return TailCall2(__e, SymFn(symshen_4compile_1prolog), tmp8239, tmp8240)


} else {
tmp8397 := GEq(__e, symdatatype, let__8086)

var ifres8395 Obj

if True == tmp8397 {
tmp8396 := GConsP(__e, let__8087)

ifres8395 = tmp8396


} else {
ifres8395 = False


}

if True == ifres8395 {
tmp8241 := GHd(__e, let__8087)

tmp8242 := GTl(__e, let__8087)

return TailCall2(__e, SymFn(symshen_4process_1datatype), tmp8241, tmp8242)


} else {
tmp8393 := GEq(__e, sym_8s, let__8086)

if True == tmp8393 {
return TailCall1(__e, SymFn(symshen_4process_1_8s), V7180)
} else {
tmp8391 := GEq(__e, symsynonyms, let__8086)

if True == tmp8391 {
return TailCall1(__e, SymFn(symshen_4process_1synonyms), let__8087)
} else {
tmp8389 := GEq(__e, symnl, let__8086)

var ifres8387 Obj

if True == tmp8389 {
tmp8388 := GEq(__e, Nil, let__8087)

ifres8387 = tmp8388


} else {
ifres8387 = False


}

if True == ifres8387 {
tmp8243 := GCons(__e, MakeInteger(1), Nil)

return GCons(__e, symnl, tmp8243)


} else {
tmp8385 := GEq(__e, symlet, let__8086)

if True == tmp8385 {
return TailCall1(__e, SymFn(symshen_4process_1let), V7180)
} else {
tmp8383 := GEq(__e, sym_c_4, let__8086)

if True == tmp8383 {
return TailCall1(__e, SymFn(symshen_4process_1lambda), V7180)
} else {
tmp8381 := GEq(__e, symcases, let__8086)

if True == tmp8381 {
return TailCall1(__e, SymFn(symshen_4process_1cases), V7180)
} else {
tmp8379 := GEq(__e, symtime, let__8086)

var ifres8374 Obj

if True == tmp8379 {
tmp8378 := GConsP(__e, let__8087)

var ifres8375 Obj

if True == tmp8378 {
tmp8376 := GTl(__e, let__8087)

tmp8377 := GEq(__e, Nil, tmp8376)

ifres8375 = tmp8377


} else {
ifres8375 = False


}

ifres8374 = ifres8375


} else {
ifres8374 = False


}

if True == ifres8374 {
tmp8244 := GHd(__e, let__8087)

return TailCall1(__e, SymFn(symshen_4process_1time), tmp8244)


} else {
tmp8372 := GEq(__e, symput, let__8086)

var ifres8361 Obj

if True == tmp8372 {
tmp8371 := GConsP(__e, let__8087)

var ifres8362 Obj

if True == tmp8371 {
tmp8363 := GTl(__e, let__8087)

let__8088 := tmp8363
_ = let__8088

tmp8370 := GConsP(__e, let__8088)

var ifres8364 Obj

if True == tmp8370 {
tmp8365 := GTl(__e, let__8088)

let__8089 := tmp8365
_ = let__8089

tmp8369 := GConsP(__e, let__8089)

var ifres8366 Obj

if True == tmp8369 {
tmp8367 := GTl(__e, let__8089)

tmp8368 := GEq(__e, Nil, tmp8367)

ifres8366 = tmp8368


} else {
ifres8366 = False


}

ifres8364 = ifres8366


} else {
ifres8364 = False


}

ifres8362 = ifres8364


} else {
ifres8362 = False


}

ifres8361 = ifres8362


} else {
ifres8361 = False


}

if True == ifres8361 {
tmp8245 := GTl(__e, let__8087)

let__8090 := tmp8245
_ = let__8090

tmp8246 := GHd(__e, let__8087)

tmp8247 := GHd(__e, let__8090)

tmp8248 := GTl(__e, let__8090)

tmp8249 := GHd(__e, tmp8248)

tmp8250 := GCons(__e, sym_dproperty_1vector_d, Nil)

tmp8251 := GCons(__e, symvalue, tmp8250)

tmp8252 := GCons(__e, tmp8251, Nil)

tmp8253 := GCons(__e, tmp8249, tmp8252)

tmp8254 := GCons(__e, tmp8247, tmp8253)

tmp8255 := GCons(__e, tmp8246, tmp8254)

return GCons(__e, symput, tmp8255)


} else {
tmp8359 := GEq(__e, symget, let__8086)

var ifres8351 Obj

if True == tmp8359 {
tmp8358 := GConsP(__e, let__8087)

var ifres8352 Obj

if True == tmp8358 {
tmp8353 := GTl(__e, let__8087)

let__8091 := tmp8353
_ = let__8091

tmp8357 := GConsP(__e, let__8091)

var ifres8354 Obj

if True == tmp8357 {
tmp8355 := GTl(__e, let__8091)

tmp8356 := GEq(__e, Nil, tmp8355)

ifres8354 = tmp8356


} else {
ifres8354 = False


}

ifres8352 = ifres8354


} else {
ifres8352 = False


}

ifres8351 = ifres8352


} else {
ifres8351 = False


}

if True == ifres8351 {
tmp8256 := GHd(__e, let__8087)

tmp8257 := GTl(__e, let__8087)

tmp8258 := GHd(__e, tmp8257)

tmp8259 := GCons(__e, sym_dproperty_1vector_d, Nil)

tmp8260 := GCons(__e, symvalue, tmp8259)

tmp8261 := GCons(__e, tmp8260, Nil)

tmp8262 := GCons(__e, tmp8258, tmp8261)

tmp8263 := GCons(__e, tmp8256, tmp8262)

return GCons(__e, symget, tmp8263)


} else {
tmp8349 := GEq(__e, symunput, let__8086)

var ifres8341 Obj

if True == tmp8349 {
tmp8348 := GConsP(__e, let__8087)

var ifres8342 Obj

if True == tmp8348 {
tmp8343 := GTl(__e, let__8087)

let__8092 := tmp8343
_ = let__8092

tmp8347 := GConsP(__e, let__8092)

var ifres8344 Obj

if True == tmp8347 {
tmp8345 := GTl(__e, let__8092)

tmp8346 := GEq(__e, Nil, tmp8345)

ifres8344 = tmp8346


} else {
ifres8344 = False


}

ifres8342 = ifres8344


} else {
ifres8342 = False


}

ifres8341 = ifres8342


} else {
ifres8341 = False


}

if True == ifres8341 {
tmp8264 := GHd(__e, let__8087)

tmp8265 := GTl(__e, let__8087)

tmp8266 := GHd(__e, tmp8265)

tmp8267 := GCons(__e, sym_dproperty_1vector_d, Nil)

tmp8268 := GCons(__e, symvalue, tmp8267)

tmp8269 := GCons(__e, tmp8268, Nil)

tmp8270 := GCons(__e, tmp8266, tmp8269)

tmp8271 := GCons(__e, tmp8264, tmp8270)

return GCons(__e, symunput, tmp8271)


} else {
tmp8339 := GEq(__e, symshen_4_8c, let__8086)

var ifres8334 Obj

if True == tmp8339 {
tmp8338 := GConsP(__e, let__8087)

var ifres8335 Obj

if True == tmp8338 {
tmp8336 := GTl(__e, let__8087)

tmp8337 := GEq(__e, Nil, tmp8336)

ifres8335 = tmp8337


} else {
ifres8335 = False


}

ifres8334 = ifres8335


} else {
ifres8334 = False


}

if True == ifres8334 {
tmp8272 := GHd(__e, let__8087)

return TailCall1(__e, SymFn(symshen_4rcons__form), tmp8272)


} else {
tmp8273 := MakeNative0(func(__e *ControlFlow) Obj {
tmp8297 := GConsP(__e, let__8087)

var ifres8281 Obj

if True == tmp8297 {
tmp8282 := GTl(__e, let__8087)

let__8094 := tmp8282
_ = let__8094

tmp8296 := GConsP(__e, let__8094)

var ifres8283 Obj

if True == tmp8296 {
tmp8294 := GTl(__e, let__8094)

tmp8295 := GConsP(__e, tmp8294)

var ifres8284 Obj

if True == tmp8295 {
tmp8285 := GCons(__e, symdo, Nil)

tmp8286 := GCons(__e, sym_d, tmp8285)

tmp8287 := GCons(__e, sym_7, tmp8286)

tmp8288 := GCons(__e, symor, tmp8287)

tmp8289 := GCons(__e, symand, tmp8288)

tmp8290 := GCons(__e, symappend, tmp8289)

tmp8291 := GCons(__e, sym_8v, tmp8290)

tmp8292 := GCons(__e, sym_8p, tmp8291)

tmp8293 := Call2(__e, SymFn(symelement_2), let__8086, tmp8292)


ifres8284 = tmp8293


} else {
ifres8284 = False


}

ifres8283 = ifres8284


} else {
ifres8283 = False


}

ifres8281 = ifres8283


} else {
ifres8281 = False


}

if True == ifres8281 {
tmp8274 := GHd(__e, let__8087)

tmp8275 := GTl(__e, let__8087)

tmp8276 := GCons(__e, let__8086, tmp8275)

tmp8277 := Call1(__e, SymFn(symshen_4process_1assoc), tmp8276)


tmp8278 := GCons(__e, tmp8277, Nil)

tmp8279 := GCons(__e, tmp8274, tmp8278)

return GCons(__e, let__8086, tmp8279)


} else {
return TailCall1(__e, SymFn(symthaw), let__8085)
}


panic("unreachable")
})

let__8093 := tmp8273
_ = let__8093

tmp8332 := GEq(__e, symshen_4_8ch, let__8086)

if True == tmp8332 {
tmp8330 := GConsP(__e, let__8087)

if True == tmp8330 {
tmp8298 := GHd(__e, let__8087)

let__8095 := tmp8298
_ = let__8095

tmp8299 := GTl(__e, let__8087)

let__8096 := tmp8299
_ = let__8096

tmp8328 := GConsP(__e, let__8095)

var ifres8313 Obj

if True == tmp8328 {
tmp8314 := GTl(__e, let__8095)

let__8097 := tmp8314
_ = let__8097

tmp8327 := GConsP(__e, let__8097)

var ifres8315 Obj

if True == tmp8327 {
tmp8316 := GTl(__e, let__8097)

let__8098 := tmp8316
_ = let__8098

tmp8326 := GConsP(__e, let__8098)

var ifres8317 Obj

if True == tmp8326 {
tmp8324 := GTl(__e, let__8098)

tmp8325 := GEq(__e, Nil, tmp8324)

var ifres8318 Obj

if True == tmp8325 {
tmp8323 := GEq(__e, Nil, let__8096)

var ifres8319 Obj

if True == tmp8323 {
tmp8320 := GHd(__e, let__8097)

tmp8321 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":"))

tmp8322 := GEq(__e, tmp8320, tmp8321)

ifres8319 = tmp8322


} else {
ifres8319 = False


}

ifres8318 = ifres8319


} else {
ifres8318 = False


}

ifres8317 = ifres8318


} else {
ifres8317 = False


}

ifres8315 = ifres8317


} else {
ifres8315 = False


}

ifres8313 = ifres8315


} else {
ifres8313 = False


}

if True == ifres8313 {
tmp8300 := GTl(__e, let__8095)

let__8099 := tmp8300
_ = let__8099

tmp8301 := GHd(__e, let__8095)

tmp8302 := GHd(__e, let__8099)

tmp8303 := GTl(__e, let__8099)

tmp8304 := GCons(__e, sym_7, tmp8303)

tmp8305 := GCons(__e, tmp8304, Nil)

tmp8306 := GCons(__e, tmp8302, tmp8305)

tmp8307 := GCons(__e, tmp8301, tmp8306)

tmp8308 := GCons(__e, tmp8307, Nil)

tmp8309 := GCons(__e, sym_1, tmp8308)

return TailCall1(__e, SymFn(symshen_4cons_1form_1respect_1modes), tmp8309)


} else {
tmp8311 := GEq(__e, Nil, let__8096)

if True == tmp8311 {
return TailCall1(__e, SymFn(symshen_4cons_1form_1respect_1modes), let__8095)
} else {
return TailCall1(__e, SymFn(symthaw), let__8093)
}


}


} else {
return TailCall1(__e, SymFn(symthaw), let__8093)
}


} else {
return TailCall1(__e, SymFn(symthaw), let__8093)
}


}


}


}


}


}


}


}


}


}


}


}


}


}


}


}


}


}


}


}


}


}


}


}


}


}


}


} else {
return TailCall1(__e, SymFn(symthaw), let__8085)
}


panic("unreachable")
})

tmp8458 := Call2(__e, ns2_1set, symshen_4macros, tmp8208)


_ = tmp8458

tmp8459 := MakeNative1(func(__e *ControlFlow, V7197 Obj) Obj {
tmp8460 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4process_1input_7)
panic("unreachable")
})

let__8100 := tmp8460
_ = let__8100

tmp8483 := GConsP(__e, V7197)

if True == tmp8483 {
tmp8461 := GTl(__e, V7197)

let__8101 := tmp8461
_ = let__8101

tmp8480 := GHd(__e, V7197)

tmp8481 := GEq(__e, syminput_7, tmp8480)

if True == tmp8481 {
tmp8478 := GConsP(__e, let__8101)

if True == tmp8478 {
tmp8462 := GHd(__e, let__8101)

let__8102 := tmp8462
_ = let__8102

tmp8463 := GTl(__e, let__8101)

let__8103 := tmp8463
_ = let__8103

tmp8476 := GEq(__e, Nil, let__8103)

if True == tmp8476 {
tmp8464 := Call1(__e, SymFn(symshen_4rcons__form), let__8102)


tmp8465 := GCons(__e, symstinput, Nil)

tmp8466 := GCons(__e, tmp8465, Nil)

tmp8467 := GCons(__e, tmp8464, tmp8466)

return GCons(__e, symshen_4input_1h_7, tmp8467)


} else {
tmp8474 := GConsP(__e, let__8103)

var ifres8471 Obj

if True == tmp8474 {
tmp8472 := GTl(__e, let__8103)

tmp8473 := GEq(__e, Nil, tmp8472)

ifres8471 = tmp8473


} else {
ifres8471 = False


}

if True == ifres8471 {
tmp8468 := Call1(__e, SymFn(symshen_4rcons__form), let__8102)


tmp8469 := GCons(__e, tmp8468, let__8103)

return GCons(__e, symshen_4input_1h_7, tmp8469)


} else {
return TailCall1(__e, SymFn(symthaw), let__8100)
}


}


} else {
return TailCall1(__e, SymFn(symthaw), let__8100)
}


} else {
return TailCall1(__e, SymFn(symthaw), let__8100)
}


} else {
return TailCall1(__e, SymFn(symthaw), let__8100)
}


panic("unreachable")
})

tmp8484 := Call2(__e, ns2_1set, symshen_4process_1input_7, tmp8459)


_ = tmp8484

tmp8485 := MakeNative1(func(__e *ControlFlow, V7204 Obj) Obj {
tmp8486 := MakeNative0(func(__e *ControlFlow) Obj {
return V7204
panic("unreachable")
})

let__8104 := tmp8486
_ = let__8104

tmp8521 := GConsP(__e, V7204)

if True == tmp8521 {
tmp8487 := GHd(__e, V7204)

let__8105 := tmp8487
_ = let__8105

tmp8488 := GTl(__e, V7204)

let__8106 := tmp8488
_ = let__8106

tmp8519 := GEq(__e, sym_7, let__8105)

var ifres8514 Obj

if True == tmp8519 {
tmp8518 := GConsP(__e, let__8106)

var ifres8515 Obj

if True == tmp8518 {
tmp8516 := GTl(__e, let__8106)

tmp8517 := GEq(__e, Nil, tmp8516)

ifres8515 = tmp8517


} else {
ifres8515 = False


}

ifres8514 = ifres8515


} else {
ifres8514 = False


}

if True == ifres8514 {
tmp8489 := GHd(__e, let__8106)

tmp8490 := Call1(__e, SymFn(symshen_4cons_1form_1respect_1modes), tmp8489)


tmp8491 := GCons(__e, tmp8490, Nil)

return GCons(__e, sym_7, tmp8491)


} else {
tmp8512 := GEq(__e, sym_1, let__8105)

var ifres8507 Obj

if True == tmp8512 {
tmp8511 := GConsP(__e, let__8106)

var ifres8508 Obj

if True == tmp8511 {
tmp8509 := GTl(__e, let__8106)

tmp8510 := GEq(__e, Nil, tmp8509)

ifres8508 = tmp8510


} else {
ifres8508 = False


}

ifres8507 = ifres8508


} else {
ifres8507 = False


}

if True == ifres8507 {
tmp8492 := GHd(__e, let__8106)

tmp8493 := Call1(__e, SymFn(symshen_4cons_1form_1respect_1modes), tmp8492)


tmp8494 := GCons(__e, tmp8493, Nil)

return GCons(__e, sym_1, tmp8494)


} else {
tmp8505 := GEq(__e, symbar_b, let__8105)

var ifres8500 Obj

if True == tmp8505 {
tmp8504 := GConsP(__e, let__8106)

var ifres8501 Obj

if True == tmp8504 {
tmp8502 := GTl(__e, let__8106)

tmp8503 := GEq(__e, Nil, tmp8502)

ifres8501 = tmp8503


} else {
ifres8501 = False


}

ifres8500 = ifres8501


} else {
ifres8500 = False


}

if True == ifres8500 {
return GHd(__e, let__8106)
} else {
tmp8495 := Call1(__e, SymFn(symshen_4cons_1form_1respect_1modes), let__8105)


tmp8496 := Call1(__e, SymFn(symshen_4cons_1form_1respect_1modes), let__8106)


tmp8497 := GCons(__e, tmp8496, Nil)

tmp8498 := GCons(__e, tmp8495, tmp8497)

return GCons(__e, symcons, tmp8498)


}


}


}


} else {
return TailCall1(__e, SymFn(symthaw), let__8104)
}


panic("unreachable")
})

tmp8522 := Call2(__e, ns2_1set, symshen_4cons_1form_1respect_1modes, tmp8485)


_ = tmp8522

tmp8523 := MakeNative2(func(__e *ControlFlow, V7208, V7209 Obj) Obj {
tmp8524 := GCons(__e, symX, Nil)

tmp8525 := GCons(__e, sym_1_6, tmp8524)

tmp8526 := GCons(__e, symX, tmp8525)

let__8107 := tmp8526
_ = let__8107

tmp8527 := Call2(__e, SymFn(symappend), V7209, let__8107)


tmp8528 := GCons(__e, V7208, tmp8527)

tmp8529 := GCons(__e, symdefine, tmp8528)

tmp8530 := Call1(__e, SymFn(symeval), tmp8529)


let__8108 := tmp8530
_ = let__8108

tmp8531 := Call1(__e, SymFn(symfn), V7208)


tmp8532 := Call2(__e, SymFn(symshen_4record_1macro), V7208, tmp8531)


let__8109 := tmp8532
_ = let__8109

return V7208


panic("unreachable")
})

tmp8533 := Call2(__e, ns2_1set, symshen_4process_1def, tmp8523)


_ = tmp8533

tmp8534 := MakeNative1(func(__e *ControlFlow, V7213 Obj) Obj {
tmp8561 := GConsP(__e, V7213)

var ifres8546 Obj

if True == tmp8561 {
tmp8547 := GTl(__e, V7213)

let__8110 := tmp8547
_ = let__8110

tmp8559 := GHd(__e, V7213)

tmp8560 := GEq(__e, symlet, tmp8559)

var ifres8548 Obj

if True == tmp8560 {
tmp8558 := GConsP(__e, let__8110)

var ifres8549 Obj

if True == tmp8558 {
tmp8550 := GTl(__e, let__8110)

let__8111 := tmp8550
_ = let__8111

tmp8557 := GConsP(__e, let__8111)

var ifres8551 Obj

if True == tmp8557 {
tmp8552 := GTl(__e, let__8111)

let__8112 := tmp8552
_ = let__8112

tmp8556 := GConsP(__e, let__8112)

var ifres8553 Obj

if True == tmp8556 {
tmp8554 := GTl(__e, let__8112)

tmp8555 := GConsP(__e, tmp8554)

ifres8553 = tmp8555


} else {
ifres8553 = False


}

ifres8551 = ifres8553


} else {
ifres8551 = False


}

ifres8549 = ifres8551


} else {
ifres8549 = False


}

ifres8548 = ifres8549


} else {
ifres8548 = False


}

ifres8546 = ifres8548


} else {
ifres8546 = False


}

if True == ifres8546 {
tmp8535 := GTl(__e, V7213)

let__8113 := tmp8535
_ = let__8113

tmp8536 := GHd(__e, let__8113)

tmp8537 := GTl(__e, let__8113)

tmp8538 := GHd(__e, tmp8537)

tmp8539 := GTl(__e, let__8113)

tmp8540 := GTl(__e, tmp8539)

tmp8541 := GCons(__e, symlet, tmp8540)

tmp8542 := GCons(__e, tmp8541, Nil)

tmp8543 := GCons(__e, tmp8538, tmp8542)

tmp8544 := GCons(__e, tmp8536, tmp8543)

return GCons(__e, symlet, tmp8544)


} else {
return V7213
}


panic("unreachable")
})

tmp8562 := Call2(__e, ns2_1set, symshen_4process_1let, tmp8534)


_ = tmp8562

tmp8563 := (func() Obj {
var __self Obj
__self = MakeNative1(func(__e *ControlFlow, __self1 Obj) Obj {
__selftop:
V7218 := __self1
_ = V7218
tmp8564 := MakeNative0(func(__e *ControlFlow) Obj {
return V7218
panic("unreachable")
})

let__8114 := tmp8564
_ = let__8114

tmp8593 := GConsP(__e, V7218)

if True == tmp8593 {
tmp8565 := GTl(__e, V7218)

let__8115 := tmp8565
_ = let__8115

tmp8590 := GHd(__e, V7218)

tmp8591 := GEq(__e, sym_8s, tmp8590)

if True == tmp8591 {
tmp8588 := GConsP(__e, let__8115)

if True == tmp8588 {
tmp8566 := GHd(__e, let__8115)

let__8116 := tmp8566
_ = let__8116

tmp8567 := GTl(__e, let__8115)

let__8117 := tmp8567
_ = let__8117

tmp8586 := GConsP(__e, let__8117)

if True == tmp8586 {
tmp8568 := GTl(__e, let__8117)

let__8118 := tmp8568
_ = let__8118

tmp8584 := GConsP(__e, let__8118)

if True == tmp8584 {
tmp8569 := GCons(__e, sym_8s, let__8117)

tmp8570 := Call1(__e, SymFn(symshen_4process_1_8s), tmp8569)


tmp8571 := GCons(__e, tmp8570, Nil)

tmp8572 := GCons(__e, let__8116, tmp8571)

return GCons(__e, sym_8s, tmp8572)


} else {
tmp8582 := GEq(__e, Nil, let__8118)

var ifres8580 Obj

if True == tmp8582 {
tmp8581 := PrimGuard1(__e, symstring_2, PrimIsString, let__8116)

ifres8580 = tmp8581


} else {
ifres8580 = False


}

if True == ifres8580 {
tmp8573 := Call1(__e, SymFn(symexplode), let__8116)


let__8119 := tmp8573
_ = let__8119

tmp8577 := Call1(__e, SymFn(symlength), let__8119)


if True == GGt(__e, tmp8577, MakeInteger(1)) {
tmp8574 := Call2(__e, SymFn(symappend), let__8119, let__8117)


tmp8575 := GCons(__e, sym_8s, tmp8574)

if SymFn(symshen_4process_1_8s) == __self {
__self1 = tmp8575
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4process_1_8s), tmp8575)


} else {
return V7218
}


} else {
return TailCall1(__e, SymFn(symthaw), let__8114)
}


}


} else {
return TailCall1(__e, SymFn(symthaw), let__8114)
}


} else {
return TailCall1(__e, SymFn(symthaw), let__8114)
}


} else {
return TailCall1(__e, SymFn(symthaw), let__8114)
}


} else {
return TailCall1(__e, SymFn(symthaw), let__8114)
}


panic("unreachable")
})
return __self
})()

tmp8594 := Call2(__e, ns2_1set, symshen_4process_1_8s, tmp8563)


_ = tmp8594

tmp8595 := MakeNative2(func(__e *ControlFlow, V7228, V7229 Obj) Obj {
tmp8596 := Call1(__e, SymFn(symshen_4intern_1type), V7228)


let__8120 := tmp8596
_ = let__8120

tmp8597 := MakeNative1(func(__e *ControlFlow, Z7232 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4_5datatype_6), Z7232)
panic("unreachable")
})

tmp8598 := GCons(__e, let__8120, V7229)

tmp8599 := Call2(__e, SymFn(symcompile), tmp8597, tmp8598)


let__8121 := tmp8599
_ = let__8121

return let__8120


panic("unreachable")
})

tmp8600 := Call2(__e, ns2_1set, symshen_4process_1datatype, tmp8595)


_ = tmp8600

tmp8601 := MakeNative1(func(__e *ControlFlow, V7233 Obj) Obj {
tmp8602 := PrimGuard1(__e, symstr, PrimStr, V7233)

return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(PrimGuard2(__e, symcn, PrimStringConcat, tmp8602, MakeString("#type")))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, tmp8602, MakeString("#type"))
return Call1(__e, SymFn(symintern), __typedArg0)
})()


panic("unreachable")
})

tmp8604 := Call2(__e, ns2_1set, symshen_4intern_1type, tmp8601)


_ = tmp8604

tmp8605 := MakeNative1(func(__e *ControlFlow, V7234 Obj) Obj {
tmp8606 := GValue(__e, symshen_4_dsynonyms_d)

tmp8607 := Call2(__e, SymFn(symappend), V7234, tmp8606)


tmp8608 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsynonyms_d, tmp8607)

return TailCall1(__e, SymFn(symshen_4synonyms_1h), tmp8608)


panic("unreachable")
})

tmp8609 := Call2(__e, ns2_1set, symshen_4process_1synonyms, tmp8605)


_ = tmp8609

tmp8610 := MakeNative1(func(__e *ControlFlow, V7237 Obj) Obj {
tmp8641 := GConsP(__e, V7237)

var ifres8620 Obj

if True == tmp8641 {
tmp8621 := GTl(__e, V7237)

let__8122 := tmp8621
_ = let__8122

tmp8639 := GHd(__e, V7237)

tmp8640 := GEq(__e, symdefun, tmp8639)

var ifres8622 Obj

if True == tmp8640 {
tmp8638 := GConsP(__e, let__8122)

var ifres8623 Obj

if True == tmp8638 {
tmp8624 := GTl(__e, let__8122)

let__8123 := tmp8624
_ = let__8123

tmp8637 := GConsP(__e, let__8123)

var ifres8625 Obj

if True == tmp8637 {
tmp8626 := GHd(__e, let__8123)

let__8124 := tmp8626
_ = let__8124

tmp8627 := GTl(__e, let__8123)

let__8125 := tmp8627
_ = let__8125

tmp8636 := GConsP(__e, let__8124)

var ifres8628 Obj

if True == tmp8636 {
tmp8634 := GTl(__e, let__8124)

tmp8635 := GEq(__e, Nil, tmp8634)

var ifres8629 Obj

if True == tmp8635 {
tmp8633 := GConsP(__e, let__8125)

var ifres8630 Obj

if True == tmp8633 {
tmp8631 := GTl(__e, let__8125)

tmp8632 := GEq(__e, Nil, tmp8631)

ifres8630 = tmp8632


} else {
ifres8630 = False


}

ifres8629 = ifres8630


} else {
ifres8629 = False


}

ifres8628 = ifres8629


} else {
ifres8628 = False


}

ifres8625 = ifres8628


} else {
ifres8625 = False


}

ifres8623 = ifres8625


} else {
ifres8623 = False


}

ifres8622 = ifres8623


} else {
ifres8622 = False


}

ifres8620 = ifres8622


} else {
ifres8620 = False


}

if True == ifres8620 {
tmp8611 := GTl(__e, V7237)

let__8126 := tmp8611
_ = let__8126

tmp8612 := GTl(__e, let__8126)

tmp8613 := GHd(__e, tmp8612)

tmp8614 := GHd(__e, tmp8613)

tmp8615 := GTl(__e, let__8126)

tmp8616 := GTl(__e, tmp8615)

tmp8617 := GCons(__e, tmp8614, tmp8616)

tmp8618 := GCons(__e, sym_c_4, tmp8617)

return TailCall1(__e, SymFn(symeval), tmp8618)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4lambda_1of_1defun)
}


panic("unreachable")
})

tmp8642 := Call2(__e, ns2_1set, symshen_4lambda_1of_1defun, tmp8610)


_ = tmp8642

tmp8643 := MakeNative1(func(__e *ControlFlow, V7243 Obj) Obj {
tmp8644 := MakeNative1(func(__e *ControlFlow, Z7245 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4curry_1type), Z7245)
panic("unreachable")
})

tmp8645 := Call2(__e, SymFn(symmap), tmp8644, V7243)


let__8127 := tmp8645
_ = let__8127

tmp8646 := Call1(__e, SymFn(symshen_4compile_1synonyms), let__8127)


tmp8647 := Call2(__e, SymFn(symshen_4shendef_1_6kldef), symshen_4demod, tmp8646)


tmp8648 := Call1(__e, SymFn(symshen_4lambda_1of_1defun), tmp8647)


let__8128 := tmp8648
_ = let__8128

tmp8649 := PrimGuard2(__e, symset, PrimSet, symshen_4_ddemodulation_1function_d, let__8128)

let__8129 := tmp8649
_ = let__8129

return symsynonyms


panic("unreachable")
})

tmp8650 := Call2(__e, ns2_1set, symshen_4synonyms_1h, tmp8643)


_ = tmp8650

tmp8651 := MakeNative1(func(__e *ControlFlow, V7250 Obj) Obj {
tmp8670 := GEq(__e, Nil, V7250)

if True == tmp8670 {
tmp8652 := Call1(__e, SymFn(symgensym), symX)


let__8130 := tmp8652
_ = let__8130

tmp8653 := GCons(__e, let__8130, Nil)

tmp8654 := GCons(__e, sym_1_6, tmp8653)

return GCons(__e, let__8130, tmp8654)


} else {
tmp8668 := GConsP(__e, V7250)

var ifres8665 Obj

if True == tmp8668 {
tmp8666 := GTl(__e, V7250)

tmp8667 := GConsP(__e, tmp8666)

ifres8665 = tmp8667


} else {
ifres8665 = False


}

if True == ifres8665 {
tmp8655 := GTl(__e, V7250)

let__8131 := tmp8655
_ = let__8131

tmp8656 := GHd(__e, V7250)

tmp8657 := Call1(__e, SymFn(symshen_4rcons__form), tmp8656)


tmp8658 := GHd(__e, let__8131)

tmp8659 := Call1(__e, SymFn(symshen_4rcons__form), tmp8658)


tmp8660 := GTl(__e, let__8131)

tmp8661 := Call1(__e, SymFn(symshen_4compile_1synonyms), tmp8660)


tmp8662 := GCons(__e, tmp8659, tmp8661)

tmp8663 := GCons(__e, sym_1_6, tmp8662)

return GCons(__e, tmp8657, tmp8663)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("synonyms requires an even number of arguments\n"))
}


}


panic("unreachable")
})

tmp8671 := Call2(__e, ns2_1set, symshen_4compile_1synonyms, tmp8651)


_ = tmp8671

tmp8672 := MakeNative1(func(__e *ControlFlow, V7253 Obj) Obj {
tmp8673 := MakeNative0(func(__e *ControlFlow) Obj {
return V7253
panic("unreachable")
})

let__8132 := tmp8673
_ = let__8132

tmp8697 := GConsP(__e, V7253)

if True == tmp8697 {
tmp8674 := GTl(__e, V7253)

let__8133 := tmp8674
_ = let__8133

tmp8694 := GHd(__e, V7253)

tmp8695 := GEq(__e, sym_c_4, tmp8694)

if True == tmp8695 {
tmp8692 := GConsP(__e, let__8133)

if True == tmp8692 {
tmp8675 := GHd(__e, let__8133)

let__8134 := tmp8675
_ = let__8134

tmp8676 := GTl(__e, let__8133)

let__8135 := tmp8676
_ = let__8135

tmp8690 := GConsP(__e, let__8135)

if True == tmp8690 {
tmp8677 := GTl(__e, let__8135)

let__8136 := tmp8677
_ = let__8136

tmp8688 := GConsP(__e, let__8136)

if True == tmp8688 {
tmp8678 := GCons(__e, sym_c_4, let__8135)

tmp8679 := Call1(__e, SymFn(symshen_4process_1lambda), tmp8678)


tmp8680 := GCons(__e, tmp8679, Nil)

tmp8681 := GCons(__e, let__8134, tmp8680)

return GCons(__e, symlambda, tmp8681)


} else {
tmp8686 := GEq(__e, Nil, let__8136)

if True == tmp8686 {
tmp8684 := PrimGuard1(__e, symvariable_2, PrimIsVariable, let__8134)

if True == tmp8684 {
return GCons(__e, symlambda, let__8133)
} else {
tmp8682 := Call3(__e, SymFn(symshen_4app), let__8134, MakeString(" is not a variable\n"), symshen_4s)


return PrimGuard1(__e, symsimple_1error, PrimSimpleError, tmp8682)


}


} else {
return TailCall1(__e, SymFn(symthaw), let__8132)
}


}


} else {
return TailCall1(__e, SymFn(symthaw), let__8132)
}


} else {
return TailCall1(__e, SymFn(symthaw), let__8132)
}


} else {
return TailCall1(__e, SymFn(symthaw), let__8132)
}


} else {
return TailCall1(__e, SymFn(symthaw), let__8132)
}


panic("unreachable")
})

tmp8698 := Call2(__e, ns2_1set, symshen_4process_1lambda, tmp8672)


_ = tmp8698

tmp8699 := MakeNative1(func(__e *ControlFlow, V7264 Obj) Obj {
tmp8700 := MakeNative0(func(__e *ControlFlow) Obj {
return V7264
panic("unreachable")
})

let__8137 := tmp8700
_ = let__8137

tmp8733 := GConsP(__e, V7264)

if True == tmp8733 {
tmp8701 := GTl(__e, V7264)

let__8138 := tmp8701
_ = let__8138

tmp8730 := GHd(__e, V7264)

tmp8731 := GEq(__e, symcases, tmp8730)

if True == tmp8731 {
tmp8728 := GConsP(__e, let__8138)

if True == tmp8728 {
tmp8702 := GHd(__e, let__8138)

let__8139 := tmp8702
_ = let__8139

tmp8703 := GTl(__e, let__8138)

let__8140 := tmp8703
_ = let__8140

tmp8726 := GEq(__e, True, let__8139)

var ifres8724 Obj

if True == tmp8726 {
tmp8725 := GConsP(__e, let__8140)

ifres8724 = tmp8725


} else {
ifres8724 = False


}

if True == ifres8724 {
return GHd(__e, let__8140)
} else {
tmp8704 := MakeNative0(func(__e *ControlFlow) Obj {
tmp8706 := GEq(__e, Nil, let__8140)

if True == tmp8706 {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("error: odd number of case elements\n"))
} else {
return TailCall1(__e, SymFn(symthaw), let__8137)
}


panic("unreachable")
})

let__8141 := tmp8704
_ = let__8141

tmp8722 := GConsP(__e, let__8140)

if True == tmp8722 {
tmp8707 := GHd(__e, let__8140)

let__8142 := tmp8707
_ = let__8142

tmp8708 := GTl(__e, let__8140)

let__8143 := tmp8708
_ = let__8143

tmp8720 := GEq(__e, Nil, let__8143)

if True == tmp8720 {
tmp8709 := GCons(__e, MakeString("error: cases exhausted"), Nil)

tmp8710 := GCons(__e, symsimple_1error, tmp8709)

tmp8711 := GCons(__e, tmp8710, Nil)

tmp8712 := GCons(__e, let__8142, tmp8711)

tmp8713 := GCons(__e, let__8139, tmp8712)

return GCons(__e, symif, tmp8713)


} else {
tmp8714 := GCons(__e, symcases, let__8143)

tmp8715 := Call1(__e, SymFn(symshen_4process_1cases), tmp8714)


tmp8716 := GCons(__e, tmp8715, Nil)

tmp8717 := GCons(__e, let__8142, tmp8716)

tmp8718 := GCons(__e, let__8139, tmp8717)

return GCons(__e, symif, tmp8718)


}


} else {
return TailCall1(__e, SymFn(symthaw), let__8141)
}


}


} else {
return TailCall1(__e, SymFn(symthaw), let__8137)
}


} else {
return TailCall1(__e, SymFn(symthaw), let__8137)
}


} else {
return TailCall1(__e, SymFn(symthaw), let__8137)
}


panic("unreachable")
})

tmp8734 := Call2(__e, ns2_1set, symshen_4process_1cases, tmp8699)


_ = tmp8734

tmp8735 := MakeNative1(func(__e *ControlFlow, V7274 Obj) Obj {
tmp8736 := GCons(__e, symrun, Nil)

tmp8737 := GCons(__e, symget_1time, tmp8736)

tmp8738 := GCons(__e, symrun, Nil)

tmp8739 := GCons(__e, symget_1time, tmp8738)

tmp8740 := GCons(__e, symStart, Nil)

tmp8741 := GCons(__e, symFinish, tmp8740)

tmp8742 := GCons(__e, sym_1, tmp8741)

tmp8743 := GCons(__e, symTime, Nil)

tmp8744 := GCons(__e, symstr, tmp8743)

tmp8745 := GCons(__e, MakeString(" secs\n"), Nil)

tmp8746 := GCons(__e, tmp8744, tmp8745)

tmp8747 := GCons(__e, symcn, tmp8746)

tmp8748 := GCons(__e, tmp8747, Nil)

tmp8749 := GCons(__e, MakeString("\nrun time: "), tmp8748)

tmp8750 := GCons(__e, symcn, tmp8749)

tmp8751 := GCons(__e, symstoutput, Nil)

tmp8752 := GCons(__e, tmp8751, Nil)

tmp8753 := GCons(__e, tmp8750, tmp8752)

tmp8754 := GCons(__e, sympr, tmp8753)

tmp8755 := GCons(__e, symResult, Nil)

tmp8756 := GCons(__e, tmp8754, tmp8755)

tmp8757 := GCons(__e, symMessage, tmp8756)

tmp8758 := GCons(__e, tmp8742, tmp8757)

tmp8759 := GCons(__e, symTime, tmp8758)

tmp8760 := GCons(__e, tmp8739, tmp8759)

tmp8761 := GCons(__e, symFinish, tmp8760)

tmp8762 := GCons(__e, V7274, tmp8761)

tmp8763 := GCons(__e, symResult, tmp8762)

tmp8764 := GCons(__e, tmp8737, tmp8763)

tmp8765 := GCons(__e, symStart, tmp8764)

return GCons(__e, symlet, tmp8765)


panic("unreachable")
})

tmp8766 := Call2(__e, ns2_1set, symshen_4process_1time, tmp8735)


_ = tmp8766

tmp8767 := MakeNative1(func(__e *ControlFlow, V7275 Obj) Obj {
tmp8785 := GConsP(__e, V7275)

var ifres8776 Obj

if True == tmp8785 {
tmp8777 := GTl(__e, V7275)

let__8144 := tmp8777
_ = let__8144

tmp8784 := GConsP(__e, let__8144)

var ifres8778 Obj

if True == tmp8784 {
tmp8779 := GTl(__e, let__8144)

let__8145 := tmp8779
_ = let__8145

tmp8783 := GConsP(__e, let__8145)

var ifres8780 Obj

if True == tmp8783 {
tmp8781 := GTl(__e, let__8145)

tmp8782 := GConsP(__e, tmp8781)

ifres8780 = tmp8782


} else {
ifres8780 = False


}

ifres8778 = ifres8780


} else {
ifres8778 = False


}

ifres8776 = ifres8778


} else {
ifres8776 = False


}

if True == ifres8776 {
tmp8768 := GHd(__e, V7275)

let__8146 := tmp8768
_ = let__8146

tmp8769 := GTl(__e, V7275)

let__8147 := tmp8769
_ = let__8147

tmp8770 := GHd(__e, let__8147)

tmp8771 := GTl(__e, let__8147)

tmp8772 := GCons(__e, let__8146, tmp8771)

tmp8773 := GCons(__e, tmp8772, Nil)

tmp8774 := GCons(__e, tmp8770, tmp8773)

return GCons(__e, let__8146, tmp8774)


} else {
return V7275
}


panic("unreachable")
})

tmp8786 := Call2(__e, ns2_1set, symshen_4process_1assoc, tmp8767)


_ = tmp8786

tmp8787 := MakeNative1(func(__e *ControlFlow, V7280 Obj) Obj {
tmp8788 := PrimGuard1(__e, symstr, PrimStr, V7280)

tmp8789 := Call1(__e, SymFn(symshen_4mu_1h), tmp8788)


return PrimGuard1(__e, symintern, PrimIntern, tmp8789)


panic("unreachable")
})

tmp8790 := Call2(__e, ns2_1set, symshen_4make_1uppercase, tmp8787)


_ = tmp8790

tmp8791 := MakeNative1(func(__e *ControlFlow, V7281 Obj) Obj {
tmp8806 := GEq(__e, MakeString(""), V7281)

if True == tmp8806 {
return MakeString("")
} else {
tmp8804 := Call1(__e, SymFn(symshen_4_7string_2), V7281)


if True == tmp8804 {
tmp8792 := Call1(__e, SymFn(symhdstr), V7281)


tmp8793 := PrimGuard1(__e, symstring_1_6n, PrimStringToNumber, tmp8792)

let__8148 := tmp8793
_ = let__8148

let__8149 := GSub(__e, let__8148, MakeInteger(32))
_ = let__8149

tmp8800 := GGe(__e, let__8148, MakeInteger(97))

var ifres8798 Obj

if True == tmp8800 {
ifres8798 = GLe(__e, let__8148, MakeInteger(122))


} else {
ifres8798 = False


}

var ifres8795 Obj

if True == ifres8798 {
tmp8796 := PrimGuard1(__e, symn_1_6string, PrimNumberToString, let__8149)

ifres8795 = tmp8796


} else {
tmp8797 := Call1(__e, SymFn(symhdstr), V7281)


ifres8795 = tmp8797


}

let__8150 := ifres8795
_ = let__8150

tmp8802 := Call1(__e, SymFn(symshen_4mu_1h), PrimGuard1(__e, symtlstr, PrimTailString, V7281))


return TailCall2(__e, SymFn(sym_8s), let__8150, tmp8802)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4mu_1h)
}


}


panic("unreachable")
})

tmp8807 := Call2(__e, ns2_1set, symshen_4mu_1h, tmp8791)


_ = tmp8807

tmp8808 := MakeNative2(func(__e *ControlFlow, V7285, V7286 Obj) Obj {
tmp8809 := GValue(__e, sym_dmacros_d)

tmp8810 := Call3(__e, SymFn(symshen_4update_1assoc), V7285, V7286, tmp8809)


return PrimGuard2(__e, symset, PrimSet, sym_dmacros_d, tmp8810)


panic("unreachable")
})

tmp8811 := Call2(__e, ns2_1set, symshen_4record_1macro, tmp8808)


_ = tmp8811

tmp8812 := MakeNative3(func(__e *ControlFlow, V7296, V7297, V7298 Obj) Obj {
tmp8828 := GEq(__e, Nil, V7298)

if True == tmp8828 {
tmp8813 := GCons(__e, V7296, V7297)

return GCons(__e, tmp8813, Nil)


} else {
tmp8814 := MakeNative0(func(__e *ControlFlow) Obj {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.update-assoc"))
panic("unreachable")
})

let__8151 := tmp8814
_ = let__8151

tmp8826 := GConsP(__e, V7298)

if True == tmp8826 {
tmp8815 := GHd(__e, V7298)

let__8152 := tmp8815
_ = let__8152

tmp8816 := GTl(__e, V7298)

let__8153 := tmp8816
_ = let__8153

tmp8824 := GConsP(__e, let__8152)

var ifres8821 Obj

if True == tmp8824 {
tmp8822 := GHd(__e, let__8152)

tmp8823 := GEq(__e, V7296, tmp8822)

ifres8821 = tmp8823


} else {
ifres8821 = False


}

if True == ifres8821 {
tmp8817 := GHd(__e, let__8152)

tmp8818 := GCons(__e, tmp8817, V7297)

return GCons(__e, tmp8818, let__8153)


} else {
tmp8819 := Call3(__e, SymFn(symshen_4update_1assoc), V7296, V7297, let__8153)


return GCons(__e, let__8152, tmp8819)


}


} else {
return TailCall1(__e, SymFn(symthaw), let__8151)
}


}


panic("unreachable")
})

tmp8829 := Call2(__e, ns2_1set, symshen_4update_1assoc, tmp8812)


_ = tmp8829

tmp8830 := MakeNative0(func(__e *ControlFlow) Obj {
tmp8838 := Call0(__e, SymFn(symstinput))


tmp8839 := Call1(__e, SymFn(symchar_1stream_2), tmp8838)


if True == tmp8839 {
tmp8831 := GCons(__e, symstinput, Nil)

tmp8832 := GCons(__e, tmp8831, Nil)

tmp8833 := GCons(__e, symread_1unit_1string, tmp8832)

tmp8834 := GCons(__e, tmp8833, Nil)

return GCons(__e, symstring_1_6n, tmp8834)


} else {
tmp8835 := GCons(__e, symstinput, Nil)

tmp8836 := GCons(__e, tmp8835, Nil)

return GCons(__e, symread_1byte, tmp8836)


}


panic("unreachable")
})

tmp8840 := Call2(__e, ns2_1set, symshen_4process_1read_1byte, tmp8830)


_ = tmp8840

tmp8841 := MakeNative1(func(__e *ControlFlow, V7302 Obj) Obj {
tmp8842 := GCons(__e, symshen_4prolog_1vector, Nil)

let__8154 := tmp8842
_ = let__8154

tmp8843 := GCons(__e, MakeInteger(0), Nil)

tmp8844 := GCons(__e, symvector, tmp8843)

tmp8845 := GCons(__e, tmp8844, Nil)

tmp8846 := GCons(__e, MakeInteger(0), tmp8845)

tmp8847 := GCons(__e, True, tmp8846)

tmp8848 := GCons(__e, sym_8v, tmp8847)

let__8155 := tmp8848
_ = let__8155

let__8156 := MakeInteger(0)
_ = let__8156

tmp8849 := GCons(__e, True, Nil)

tmp8850 := GCons(__e, symfreeze, tmp8849)

let__8157 := tmp8850
_ = let__8157

tmp8851 := MakeNative1(func(__e *ControlFlow, Z7308 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4_5body_6), Z7308)
panic("unreachable")
})

tmp8852 := Call2(__e, SymFn(symcompile), tmp8851, V7302)


let__8158 := tmp8852
_ = let__8158

tmp8853 := Call1(__e, SymFn(symshen_4received), V7302)


let__8159 := tmp8853
_ = let__8159

tmp8854 := Call1(__e, SymFn(symgensym), symV)


let__8160 := tmp8854
_ = let__8160

tmp8855 := Call1(__e, SymFn(symgensym), symL)


let__8161 := tmp8855
_ = let__8161

tmp8856 := Call1(__e, SymFn(symgensym), symK)


let__8162 := tmp8856
_ = let__8162

tmp8857 := Call1(__e, SymFn(symgensym), symC)


let__8163 := tmp8857
_ = let__8163

tmp8858 := Call6(__e, SymFn(symshen_4continue), let__8159, let__8158, let__8160, let__8161, let__8162, let__8163)


tmp8859 := GCons(__e, tmp8858, Nil)

tmp8860 := GCons(__e, let__8163, tmp8859)

tmp8861 := GCons(__e, symlambda, tmp8860)

tmp8862 := GCons(__e, tmp8861, Nil)

tmp8863 := GCons(__e, let__8162, tmp8862)

tmp8864 := GCons(__e, symlambda, tmp8863)

tmp8865 := GCons(__e, tmp8864, Nil)

tmp8866 := GCons(__e, let__8161, tmp8865)

tmp8867 := GCons(__e, symlambda, tmp8866)

tmp8868 := GCons(__e, tmp8867, Nil)

tmp8869 := GCons(__e, let__8160, tmp8868)

tmp8870 := GCons(__e, symlambda, tmp8869)

let__8164 := tmp8870
_ = let__8164

tmp8871 := GCons(__e, let__8157, Nil)

tmp8872 := GCons(__e, let__8156, tmp8871)

tmp8873 := GCons(__e, let__8155, tmp8872)

tmp8874 := GCons(__e, let__8154, tmp8873)

return GCons(__e, let__8164, tmp8874)


panic("unreachable")
})

tmp8875 := Call2(__e, ns2_1set, symshen_4call_1prolog, tmp8841)


_ = tmp8875

tmp8876 := MakeNative1(func(__e *ControlFlow, V7317 Obj) Obj {
tmp8877 := MakeNative0(func(__e *ControlFlow) Obj {
return Nil
panic("unreachable")
})

let__8165 := tmp8877
_ = let__8165

tmp8890 := GConsP(__e, V7317)

if True == tmp8890 {
tmp8878 := GHd(__e, V7317)

let__8166 := tmp8878
_ = let__8166

tmp8879 := GTl(__e, V7317)

let__8167 := tmp8879
_ = let__8167

tmp8888 := GEq(__e, symreceive, let__8166)

var ifres8883 Obj

if True == tmp8888 {
tmp8887 := GConsP(__e, let__8167)

var ifres8884 Obj

if True == tmp8887 {
tmp8885 := GTl(__e, let__8167)

tmp8886 := GEq(__e, Nil, tmp8885)

ifres8884 = tmp8886


} else {
ifres8884 = False


}

ifres8883 = ifres8884


} else {
ifres8883 = False


}

if True == ifres8883 {
return let__8167
} else {
tmp8880 := Call1(__e, SymFn(symshen_4received), let__8166)


tmp8881 := Call1(__e, SymFn(symshen_4received), let__8167)


return TailCall2(__e, SymFn(symunion), tmp8880, tmp8881)


}


} else {
return TailCall1(__e, SymFn(symthaw), let__8165)
}


panic("unreachable")
})

tmp8891 := Call2(__e, ns2_1set, symshen_4received, tmp8876)


_ = tmp8891

tmp8892 := MakeNative0(func(__e *ControlFlow) Obj {
tmp8893 := GValue(__e, symshen_4_dprolog_1memory_d)

tmp8894 := PrimGuard1(__e, symabsvector, PrimAbsvector, tmp8893)

let__8168 := tmp8894
_ = let__8168

tmp8895 := GVecSet(__e, let__8168, MakeInteger(0), symshen_4print_1prolog_1vector)

let__8169 := tmp8895
_ = let__8169

tmp8896 := GVecSet(__e, let__8168, MakeInteger(1), MakeInteger(2))

let__8170 := tmp8896
_ = let__8170

return let__8170


panic("unreachable")
})

tmp8897 := Call2(__e, ns2_1set, symshen_4prolog_1vector, tmp8892)


_ = tmp8897

tmp8898 := MakeNative1(func(__e *ControlFlow, V7324 Obj) Obj {
return V7324
panic("unreachable")
})

tmp8899 := Call2(__e, ns2_1set, symreceive, tmp8898)


_ = tmp8899

tmp8900 := MakeNative1(func(__e *ControlFlow, V7325 Obj) Obj {
tmp8901 := MakeNative0(func(__e *ControlFlow) Obj {
return V7325
panic("unreachable")
})

let__8171 := tmp8901
_ = let__8171

tmp8916 := GConsP(__e, V7325)

if True == tmp8916 {
tmp8902 := GHd(__e, V7325)

let__8172 := tmp8902
_ = let__8172

tmp8903 := GTl(__e, V7325)

let__8173 := tmp8903
_ = let__8173

tmp8914 := GEq(__e, symbar_b, let__8172)

var ifres8909 Obj

if True == tmp8914 {
tmp8913 := GConsP(__e, let__8173)

var ifres8910 Obj

if True == tmp8913 {
tmp8911 := GTl(__e, let__8173)

tmp8912 := GEq(__e, Nil, tmp8911)

ifres8910 = tmp8912


} else {
ifres8910 = False


}

ifres8909 = ifres8910


} else {
ifres8909 = False


}

if True == ifres8909 {
return GHd(__e, let__8173)
} else {
tmp8904 := Call1(__e, SymFn(symshen_4rcons__form), let__8172)


tmp8905 := Call1(__e, SymFn(symshen_4rcons__form), let__8173)


tmp8906 := GCons(__e, tmp8905, Nil)

tmp8907 := GCons(__e, tmp8904, tmp8906)

return GCons(__e, symcons, tmp8907)


}


} else {
return TailCall1(__e, SymFn(symthaw), let__8171)
}


panic("unreachable")
})

tmp8917 := Call2(__e, ns2_1set, symshen_4rcons__form, tmp8900)


_ = tmp8917

tmp8918 := MakeNative1(func(__e *ControlFlow, V7329 Obj) Obj {
tmp8925 := GConsP(__e, V7329)

if True == tmp8925 {
tmp8919 := GHd(__e, V7329)

tmp8920 := GTl(__e, V7329)

tmp8921 := Call1(__e, SymFn(symshen_4tuple_1up), tmp8920)


tmp8922 := GCons(__e, tmp8921, Nil)

tmp8923 := GCons(__e, tmp8919, tmp8922)

return GCons(__e, sym_8p, tmp8923)


} else {
return V7329
}


panic("unreachable")
})

tmp8926 := Call2(__e, ns2_1set, symshen_4tuple_1up, tmp8918)


_ = tmp8926

tmp8927 := MakeNative1(func(__e *ControlFlow, V7330 Obj) Obj {
tmp8928 := GValue(__e, sym_dmacros_d)

tmp8929 := Call2(__e, SymFn(symassoc), V7330, tmp8928)


tmp8930 := GValue(__e, sym_dmacros_d)

tmp8931 := Call2(__e, SymFn(symremove), tmp8929, tmp8930)


tmp8932 := PrimGuard2(__e, symset, PrimSet, sym_dmacros_d, tmp8931)

_ = tmp8932

return V7330


panic("unreachable")
})

return TailCall2(__e, ns2_1set, symundefmacro, tmp8927)




panic("unreachable")
})

