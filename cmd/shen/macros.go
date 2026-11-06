package main

import . "github.com/pyrex41/shen-go/kl"

var MacrosMain = MakeNative0(func(__e *ControlFlow) Obj {
tmp8154 := MakeNative1(func(__e *ControlFlow, V7129 Obj) Obj {
tmp8155 := MakeNative1(func(__e *ControlFlow, Z7131 Obj) Obj {
return GTl(__e, Z7131)
panic("unreachable")
})

tmp8156 := GValue(__e, sym_dmacros_d)

tmp8157 := Call2(__e, SymFn(symmap), tmp8155, tmp8156)


let__8063 := tmp8157
_ = let__8063

tmp8158 := GValue(__e, symshen_4_dsource_1form_1handlers_d)

return TailCall4(__e, SymFn(symshen_4macroexpand_1h), V7129, let__8063, let__8063, tmp8158)


panic("unreachable")
})

tmp8159 := Call2(__e, ns2_1set, symmacroexpand, tmp8154)


_ = tmp8159

tmp8160 := (func() Obj {
var __self Obj
__self = MakeNative4(func(__e *ControlFlow, __self1, __self2, __self3, __self4 Obj) Obj {
__selftop:
V7144 := __self1
_ = V7144
V7145 := __self2
_ = V7145
V7146 := __self3
_ = V7146
V7147 := __self4
_ = V7147
tmp8169 := GEq(__e, Nil, V7145)

if True == tmp8169 {
return V7144
} else {
tmp8167 := GConsP(__e, V7145)

if True == tmp8167 {
tmp8161 := GHd(__e, V7145)

tmp8162 := Call3(__e, SymFn(symshen_4source_1macro_1walk), tmp8161, V7144, V7147)


let__8064 := tmp8162
_ = let__8064

tmp8165 := GEq(__e, V7144, let__8064)

if True == tmp8165 {
tmp8163 := GTl(__e, V7145)

if SymFn(symshen_4macroexpand_1h) == __self {
__self1, __self2, __self3, __self4 = V7144, tmp8163, V7146, V7147
__e.Tick()
goto __selftop
}
return TailCall4(__e, SymFn(symshen_4macroexpand_1h), V7144, tmp8163, V7146, V7147)


} else {
if SymFn(symshen_4macroexpand_1h) == __self {
__self1, __self2, __self3, __self4 = let__8064, V7146, V7146, V7147
__e.Tick()
goto __selftop
}
return TailCall4(__e, SymFn(symshen_4macroexpand_1h), let__8064, V7146, V7146, V7147)
}


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.macroexpand-h"))
}


}


panic("unreachable")
})
return __self
})()

tmp8170 := Call2(__e, ns2_1set, symshen_4macroexpand_1h, tmp8160)


_ = tmp8170

tmp8171 := MakeNative3(func(__e *ControlFlow, V7152, V7153, V7154 Obj) Obj {
tmp8180 := GEq(__e, Nil, V7154)

if True == tmp8180 {
return TailCall2(__e, SymFn(symshen_4walk), V7152, V7153)
} else {
tmp8177 := Call2(__e, SymFn(symshen_4source_1form_1handler), V7153, V7154)


tmp8178 := GConsP(__e, tmp8177)

if True == tmp8178 {
return V7153
} else {
tmp8175 := GConsP(__e, V7153)

if True == tmp8175 {
tmp8172 := MakeNative1(func(__e *ControlFlow, Z7155 Obj) Obj {
return TailCall3(__e, SymFn(symshen_4source_1macro_1walk), V7152, Z7155, V7154)
panic("unreachable")
})

tmp8173 := Call2(__e, SymFn(symmap), tmp8172, V7153)


return TailCall1(__e, V7152, tmp8173)


} else {
return TailCall1(__e, V7152, V7153)
}


}


}


panic("unreachable")
})

tmp8181 := Call2(__e, ns2_1set, symshen_4source_1macro_1walk, tmp8171)


_ = tmp8181

tmp8182 := MakeNative2(func(__e *ControlFlow, V7156, V7157 Obj) Obj {
tmp8186 := GConsP(__e, V7157)

if True == tmp8186 {
tmp8183 := MakeNative1(func(__e *ControlFlow, Z7158 Obj) Obj {
return TailCall2(__e, SymFn(symshen_4walk), V7156, Z7158)
panic("unreachable")
})

tmp8184 := Call2(__e, SymFn(symmap), tmp8183, V7157)


return TailCall1(__e, V7156, tmp8184)


} else {
return TailCall1(__e, V7156, V7157)
}


panic("unreachable")
})

tmp8187 := Call2(__e, ns2_1set, symshen_4walk, tmp8182)


_ = tmp8187

tmp8188 := MakeNative1(func(__e *ControlFlow, V7159 Obj) Obj {
tmp8189 := MakeNative0(func(__e *ControlFlow) Obj {
return V7159
panic("unreachable")
})

let__8065 := tmp8189
_ = let__8065

tmp8437 := GConsP(__e, V7159)

if True == tmp8437 {
tmp8190 := GHd(__e, V7159)

let__8066 := tmp8190
_ = let__8066

tmp8191 := GTl(__e, V7159)

let__8067 := tmp8191
_ = let__8067

tmp8435 := GEq(__e, symdefmacro, let__8066)

var ifres8433 Obj

if True == tmp8435 {
tmp8434 := GConsP(__e, let__8067)

ifres8433 = tmp8434


} else {
ifres8433 = False


}

if True == ifres8433 {
tmp8192 := GHd(__e, let__8067)

tmp8193 := GTl(__e, let__8067)

return TailCall2(__e, SymFn(symshen_4process_1def), tmp8192, tmp8193)


} else {
tmp8431 := GEq(__e, symdefcc, let__8066)

if True == tmp8431 {
return TailCall1(__e, SymFn(symshen_4yacc_1_6shen), let__8067)
} else {
tmp8429 := GEq(__e, symu_b, let__8066)

var ifres8424 Obj

if True == tmp8429 {
tmp8428 := GConsP(__e, let__8067)

var ifres8425 Obj

if True == tmp8428 {
tmp8426 := GTl(__e, let__8067)

tmp8427 := GEq(__e, Nil, tmp8426)

ifres8425 = tmp8427


} else {
ifres8425 = False


}

ifres8424 = ifres8425


} else {
ifres8424 = False


}

if True == ifres8424 {
tmp8194 := GHd(__e, let__8067)

tmp8195 := Call1(__e, SymFn(symshen_4make_1uppercase), tmp8194)


tmp8196 := GCons(__e, tmp8195, Nil)

return GCons(__e, symprotect, tmp8196)


} else {
tmp8422 := GEq(__e, symerror, let__8066)

var ifres8420 Obj

if True == tmp8422 {
tmp8421 := GConsP(__e, let__8067)

ifres8420 = tmp8421


} else {
ifres8420 = False


}

if True == ifres8420 {
tmp8197 := GHd(__e, let__8067)

tmp8198 := GTl(__e, let__8067)

tmp8199 := Call2(__e, SymFn(symshen_4mkstr), tmp8197, tmp8198)


tmp8200 := GCons(__e, tmp8199, Nil)

return GCons(__e, symsimple_1error, tmp8200)


} else {
tmp8418 := GEq(__e, symoutput, let__8066)

var ifres8416 Obj

if True == tmp8418 {
tmp8417 := GConsP(__e, let__8067)

ifres8416 = tmp8417


} else {
ifres8416 = False


}

if True == ifres8416 {
tmp8201 := GHd(__e, let__8067)

tmp8202 := GTl(__e, let__8067)

tmp8203 := Call2(__e, SymFn(symshen_4mkstr), tmp8201, tmp8202)


tmp8204 := GCons(__e, symstoutput, Nil)

tmp8205 := GCons(__e, tmp8204, Nil)

tmp8206 := GCons(__e, tmp8203, tmp8205)

return GCons(__e, sympr, tmp8206)


} else {
tmp8414 := GEq(__e, sympr, let__8066)

var ifres8409 Obj

if True == tmp8414 {
tmp8413 := GConsP(__e, let__8067)

var ifres8410 Obj

if True == tmp8413 {
tmp8411 := GTl(__e, let__8067)

tmp8412 := GEq(__e, Nil, tmp8411)

ifres8410 = tmp8412


} else {
ifres8410 = False


}

ifres8409 = ifres8410


} else {
ifres8409 = False


}

if True == ifres8409 {
tmp8207 := GHd(__e, let__8067)

tmp8208 := GCons(__e, symstoutput, Nil)

tmp8209 := GCons(__e, tmp8208, Nil)

tmp8210 := GCons(__e, tmp8207, tmp8209)

return GCons(__e, sympr, tmp8210)


} else {
tmp8407 := GEq(__e, symmake_1string, let__8066)

var ifres8405 Obj

if True == tmp8407 {
tmp8406 := GConsP(__e, let__8067)

ifres8405 = tmp8406


} else {
ifres8405 = False


}

if True == ifres8405 {
tmp8211 := GHd(__e, let__8067)

tmp8212 := GTl(__e, let__8067)

return TailCall2(__e, SymFn(symshen_4mkstr), tmp8211, tmp8212)


} else {
tmp8403 := GEq(__e, symlineread, let__8066)

var ifres8401 Obj

if True == tmp8403 {
tmp8402 := GEq(__e, Nil, let__8067)

ifres8401 = tmp8402


} else {
ifres8401 = False


}

if True == ifres8401 {
tmp8213 := GCons(__e, symstinput, Nil)

tmp8214 := GCons(__e, tmp8213, Nil)

return GCons(__e, symlineread, tmp8214)


} else {
tmp8399 := GEq(__e, syminput, let__8066)

var ifres8397 Obj

if True == tmp8399 {
tmp8398 := GEq(__e, Nil, let__8067)

ifres8397 = tmp8398


} else {
ifres8397 = False


}

if True == ifres8397 {
tmp8215 := GCons(__e, symstinput, Nil)

tmp8216 := GCons(__e, tmp8215, Nil)

return GCons(__e, syminput, tmp8216)


} else {
tmp8395 := GEq(__e, symread, let__8066)

var ifres8393 Obj

if True == tmp8395 {
tmp8394 := GEq(__e, Nil, let__8067)

ifres8393 = tmp8394


} else {
ifres8393 = False


}

if True == ifres8393 {
tmp8217 := GCons(__e, symstinput, Nil)

tmp8218 := GCons(__e, tmp8217, Nil)

return GCons(__e, symread, tmp8218)


} else {
tmp8391 := GEq(__e, syminput_7, let__8066)

var ifres8389 Obj

if True == tmp8391 {
tmp8390 := GConsP(__e, let__8067)

ifres8389 = tmp8390


} else {
ifres8389 = False


}

if True == ifres8389 {
return TailCall1(__e, SymFn(symshen_4process_1input_7), V7159)
} else {
tmp8387 := GEq(__e, symread_1byte, let__8066)

var ifres8385 Obj

if True == tmp8387 {
tmp8386 := GEq(__e, Nil, let__8067)

ifres8385 = tmp8386


} else {
ifres8385 = False


}

if True == ifres8385 {
return TailCall0(__e, SymFn(symshen_4process_1read_1byte))
} else {
tmp8383 := GEq(__e, symprolog_2, let__8066)

if True == tmp8383 {
return TailCall1(__e, SymFn(symshen_4call_1prolog), let__8067)
} else {
tmp8381 := GEq(__e, symdefprolog, let__8066)

var ifres8379 Obj

if True == tmp8381 {
tmp8380 := GConsP(__e, let__8067)

ifres8379 = tmp8380


} else {
ifres8379 = False


}

if True == ifres8379 {
tmp8219 := GHd(__e, let__8067)

tmp8220 := GTl(__e, let__8067)

return TailCall2(__e, SymFn(symshen_4compile_1prolog), tmp8219, tmp8220)


} else {
tmp8377 := GEq(__e, symdatatype, let__8066)

var ifres8375 Obj

if True == tmp8377 {
tmp8376 := GConsP(__e, let__8067)

ifres8375 = tmp8376


} else {
ifres8375 = False


}

if True == ifres8375 {
tmp8221 := GHd(__e, let__8067)

tmp8222 := GTl(__e, let__8067)

return TailCall2(__e, SymFn(symshen_4process_1datatype), tmp8221, tmp8222)


} else {
tmp8373 := GEq(__e, sym_8s, let__8066)

if True == tmp8373 {
return TailCall1(__e, SymFn(symshen_4process_1_8s), V7159)
} else {
tmp8371 := GEq(__e, symsynonyms, let__8066)

if True == tmp8371 {
return TailCall1(__e, SymFn(symshen_4process_1synonyms), let__8067)
} else {
tmp8369 := GEq(__e, symnl, let__8066)

var ifres8367 Obj

if True == tmp8369 {
tmp8368 := GEq(__e, Nil, let__8067)

ifres8367 = tmp8368


} else {
ifres8367 = False


}

if True == ifres8367 {
tmp8223 := GCons(__e, MakeInteger(1), Nil)

return GCons(__e, symnl, tmp8223)


} else {
tmp8365 := GEq(__e, symlet, let__8066)

if True == tmp8365 {
return TailCall1(__e, SymFn(symshen_4process_1let), V7159)
} else {
tmp8363 := GEq(__e, sym_c_4, let__8066)

if True == tmp8363 {
return TailCall1(__e, SymFn(symshen_4process_1lambda), V7159)
} else {
tmp8361 := GEq(__e, symcases, let__8066)

if True == tmp8361 {
return TailCall1(__e, SymFn(symshen_4process_1cases), V7159)
} else {
tmp8359 := GEq(__e, symtime, let__8066)

var ifres8354 Obj

if True == tmp8359 {
tmp8358 := GConsP(__e, let__8067)

var ifres8355 Obj

if True == tmp8358 {
tmp8356 := GTl(__e, let__8067)

tmp8357 := GEq(__e, Nil, tmp8356)

ifres8355 = tmp8357


} else {
ifres8355 = False


}

ifres8354 = ifres8355


} else {
ifres8354 = False


}

if True == ifres8354 {
tmp8224 := GHd(__e, let__8067)

return TailCall1(__e, SymFn(symshen_4process_1time), tmp8224)


} else {
tmp8352 := GEq(__e, symput, let__8066)

var ifres8341 Obj

if True == tmp8352 {
tmp8351 := GConsP(__e, let__8067)

var ifres8342 Obj

if True == tmp8351 {
tmp8343 := GTl(__e, let__8067)

let__8068 := tmp8343
_ = let__8068

tmp8350 := GConsP(__e, let__8068)

var ifres8344 Obj

if True == tmp8350 {
tmp8345 := GTl(__e, let__8068)

let__8069 := tmp8345
_ = let__8069

tmp8349 := GConsP(__e, let__8069)

var ifres8346 Obj

if True == tmp8349 {
tmp8347 := GTl(__e, let__8069)

tmp8348 := GEq(__e, Nil, tmp8347)

ifres8346 = tmp8348


} else {
ifres8346 = False


}

ifres8344 = ifres8346


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
tmp8225 := GTl(__e, let__8067)

let__8070 := tmp8225
_ = let__8070

tmp8226 := GHd(__e, let__8067)

tmp8227 := GHd(__e, let__8070)

tmp8228 := GTl(__e, let__8070)

tmp8229 := GHd(__e, tmp8228)

tmp8230 := GCons(__e, sym_dproperty_1vector_d, Nil)

tmp8231 := GCons(__e, symvalue, tmp8230)

tmp8232 := GCons(__e, tmp8231, Nil)

tmp8233 := GCons(__e, tmp8229, tmp8232)

tmp8234 := GCons(__e, tmp8227, tmp8233)

tmp8235 := GCons(__e, tmp8226, tmp8234)

return GCons(__e, symput, tmp8235)


} else {
tmp8339 := GEq(__e, symget, let__8066)

var ifres8331 Obj

if True == tmp8339 {
tmp8338 := GConsP(__e, let__8067)

var ifres8332 Obj

if True == tmp8338 {
tmp8333 := GTl(__e, let__8067)

let__8071 := tmp8333
_ = let__8071

tmp8337 := GConsP(__e, let__8071)

var ifres8334 Obj

if True == tmp8337 {
tmp8335 := GTl(__e, let__8071)

tmp8336 := GEq(__e, Nil, tmp8335)

ifres8334 = tmp8336


} else {
ifres8334 = False


}

ifres8332 = ifres8334


} else {
ifres8332 = False


}

ifres8331 = ifres8332


} else {
ifres8331 = False


}

if True == ifres8331 {
tmp8236 := GHd(__e, let__8067)

tmp8237 := GTl(__e, let__8067)

tmp8238 := GHd(__e, tmp8237)

tmp8239 := GCons(__e, sym_dproperty_1vector_d, Nil)

tmp8240 := GCons(__e, symvalue, tmp8239)

tmp8241 := GCons(__e, tmp8240, Nil)

tmp8242 := GCons(__e, tmp8238, tmp8241)

tmp8243 := GCons(__e, tmp8236, tmp8242)

return GCons(__e, symget, tmp8243)


} else {
tmp8329 := GEq(__e, symunput, let__8066)

var ifres8321 Obj

if True == tmp8329 {
tmp8328 := GConsP(__e, let__8067)

var ifres8322 Obj

if True == tmp8328 {
tmp8323 := GTl(__e, let__8067)

let__8072 := tmp8323
_ = let__8072

tmp8327 := GConsP(__e, let__8072)

var ifres8324 Obj

if True == tmp8327 {
tmp8325 := GTl(__e, let__8072)

tmp8326 := GEq(__e, Nil, tmp8325)

ifres8324 = tmp8326


} else {
ifres8324 = False


}

ifres8322 = ifres8324


} else {
ifres8322 = False


}

ifres8321 = ifres8322


} else {
ifres8321 = False


}

if True == ifres8321 {
tmp8244 := GHd(__e, let__8067)

tmp8245 := GTl(__e, let__8067)

tmp8246 := GHd(__e, tmp8245)

tmp8247 := GCons(__e, sym_dproperty_1vector_d, Nil)

tmp8248 := GCons(__e, symvalue, tmp8247)

tmp8249 := GCons(__e, tmp8248, Nil)

tmp8250 := GCons(__e, tmp8246, tmp8249)

tmp8251 := GCons(__e, tmp8244, tmp8250)

return GCons(__e, symunput, tmp8251)


} else {
tmp8319 := GEq(__e, symshen_4_8c, let__8066)

var ifres8314 Obj

if True == tmp8319 {
tmp8318 := GConsP(__e, let__8067)

var ifres8315 Obj

if True == tmp8318 {
tmp8316 := GTl(__e, let__8067)

tmp8317 := GEq(__e, Nil, tmp8316)

ifres8315 = tmp8317


} else {
ifres8315 = False


}

ifres8314 = ifres8315


} else {
ifres8314 = False


}

if True == ifres8314 {
tmp8252 := GHd(__e, let__8067)

return TailCall1(__e, SymFn(symshen_4rcons__form), tmp8252)


} else {
tmp8253 := MakeNative0(func(__e *ControlFlow) Obj {
tmp8277 := GConsP(__e, let__8067)

var ifres8261 Obj

if True == tmp8277 {
tmp8262 := GTl(__e, let__8067)

let__8074 := tmp8262
_ = let__8074

tmp8276 := GConsP(__e, let__8074)

var ifres8263 Obj

if True == tmp8276 {
tmp8274 := GTl(__e, let__8074)

tmp8275 := GConsP(__e, tmp8274)

var ifres8264 Obj

if True == tmp8275 {
tmp8265 := GCons(__e, symdo, Nil)

tmp8266 := GCons(__e, sym_d, tmp8265)

tmp8267 := GCons(__e, sym_7, tmp8266)

tmp8268 := GCons(__e, symor, tmp8267)

tmp8269 := GCons(__e, symand, tmp8268)

tmp8270 := GCons(__e, symappend, tmp8269)

tmp8271 := GCons(__e, sym_8v, tmp8270)

tmp8272 := GCons(__e, sym_8p, tmp8271)

tmp8273 := Call2(__e, SymFn(symelement_2), let__8066, tmp8272)


ifres8264 = tmp8273


} else {
ifres8264 = False


}

ifres8263 = ifres8264


} else {
ifres8263 = False


}

ifres8261 = ifres8263


} else {
ifres8261 = False


}

if True == ifres8261 {
tmp8254 := GHd(__e, let__8067)

tmp8255 := GTl(__e, let__8067)

tmp8256 := GCons(__e, let__8066, tmp8255)

tmp8257 := Call1(__e, SymFn(symshen_4process_1assoc), tmp8256)


tmp8258 := GCons(__e, tmp8257, Nil)

tmp8259 := GCons(__e, tmp8254, tmp8258)

return GCons(__e, let__8066, tmp8259)


} else {
return TailCall1(__e, SymFn(symthaw), let__8065)
}


panic("unreachable")
})

let__8073 := tmp8253
_ = let__8073

tmp8312 := GEq(__e, symshen_4_8ch, let__8066)

if True == tmp8312 {
tmp8310 := GConsP(__e, let__8067)

if True == tmp8310 {
tmp8278 := GHd(__e, let__8067)

let__8075 := tmp8278
_ = let__8075

tmp8279 := GTl(__e, let__8067)

let__8076 := tmp8279
_ = let__8076

tmp8308 := GConsP(__e, let__8075)

var ifres8293 Obj

if True == tmp8308 {
tmp8294 := GTl(__e, let__8075)

let__8077 := tmp8294
_ = let__8077

tmp8307 := GConsP(__e, let__8077)

var ifres8295 Obj

if True == tmp8307 {
tmp8296 := GTl(__e, let__8077)

let__8078 := tmp8296
_ = let__8078

tmp8306 := GConsP(__e, let__8078)

var ifres8297 Obj

if True == tmp8306 {
tmp8304 := GTl(__e, let__8078)

tmp8305 := GEq(__e, Nil, tmp8304)

var ifres8298 Obj

if True == tmp8305 {
tmp8303 := GEq(__e, Nil, let__8076)

var ifres8299 Obj

if True == tmp8303 {
tmp8300 := GHd(__e, let__8077)

tmp8301 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":"))

tmp8302 := GEq(__e, tmp8300, tmp8301)

ifres8299 = tmp8302


} else {
ifres8299 = False


}

ifres8298 = ifres8299


} else {
ifres8298 = False


}

ifres8297 = ifres8298


} else {
ifres8297 = False


}

ifres8295 = ifres8297


} else {
ifres8295 = False


}

ifres8293 = ifres8295


} else {
ifres8293 = False


}

if True == ifres8293 {
tmp8280 := GTl(__e, let__8075)

let__8079 := tmp8280
_ = let__8079

tmp8281 := GHd(__e, let__8075)

tmp8282 := GHd(__e, let__8079)

tmp8283 := GTl(__e, let__8079)

tmp8284 := GCons(__e, sym_7, tmp8283)

tmp8285 := GCons(__e, tmp8284, Nil)

tmp8286 := GCons(__e, tmp8282, tmp8285)

tmp8287 := GCons(__e, tmp8281, tmp8286)

tmp8288 := GCons(__e, tmp8287, Nil)

tmp8289 := GCons(__e, sym_1, tmp8288)

return TailCall1(__e, SymFn(symshen_4cons_1form_1respect_1modes), tmp8289)


} else {
tmp8291 := GEq(__e, Nil, let__8076)

if True == tmp8291 {
return TailCall1(__e, SymFn(symshen_4cons_1form_1respect_1modes), let__8075)
} else {
return TailCall1(__e, SymFn(symthaw), let__8073)
}


}


} else {
return TailCall1(__e, SymFn(symthaw), let__8073)
}


} else {
return TailCall1(__e, SymFn(symthaw), let__8073)
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
return TailCall1(__e, SymFn(symthaw), let__8065)
}


panic("unreachable")
})

tmp8438 := Call2(__e, ns2_1set, symshen_4macros, tmp8188)


_ = tmp8438

tmp8439 := MakeNative1(func(__e *ControlFlow, V7176 Obj) Obj {
tmp8440 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4process_1input_7)
panic("unreachable")
})

let__8080 := tmp8440
_ = let__8080

tmp8463 := GConsP(__e, V7176)

if True == tmp8463 {
tmp8441 := GTl(__e, V7176)

let__8081 := tmp8441
_ = let__8081

tmp8460 := GHd(__e, V7176)

tmp8461 := GEq(__e, syminput_7, tmp8460)

if True == tmp8461 {
tmp8458 := GConsP(__e, let__8081)

if True == tmp8458 {
tmp8442 := GHd(__e, let__8081)

let__8082 := tmp8442
_ = let__8082

tmp8443 := GTl(__e, let__8081)

let__8083 := tmp8443
_ = let__8083

tmp8456 := GEq(__e, Nil, let__8083)

if True == tmp8456 {
tmp8444 := Call1(__e, SymFn(symshen_4rcons__form), let__8082)


tmp8445 := GCons(__e, symstinput, Nil)

tmp8446 := GCons(__e, tmp8445, Nil)

tmp8447 := GCons(__e, tmp8444, tmp8446)

return GCons(__e, symshen_4input_1h_7, tmp8447)


} else {
tmp8454 := GConsP(__e, let__8083)

var ifres8451 Obj

if True == tmp8454 {
tmp8452 := GTl(__e, let__8083)

tmp8453 := GEq(__e, Nil, tmp8452)

ifres8451 = tmp8453


} else {
ifres8451 = False


}

if True == ifres8451 {
tmp8448 := Call1(__e, SymFn(symshen_4rcons__form), let__8082)


tmp8449 := GCons(__e, tmp8448, let__8083)

return GCons(__e, symshen_4input_1h_7, tmp8449)


} else {
return TailCall1(__e, SymFn(symthaw), let__8080)
}


}


} else {
return TailCall1(__e, SymFn(symthaw), let__8080)
}


} else {
return TailCall1(__e, SymFn(symthaw), let__8080)
}


} else {
return TailCall1(__e, SymFn(symthaw), let__8080)
}


panic("unreachable")
})

tmp8464 := Call2(__e, ns2_1set, symshen_4process_1input_7, tmp8439)


_ = tmp8464

tmp8465 := MakeNative1(func(__e *ControlFlow, V7183 Obj) Obj {
tmp8466 := MakeNative0(func(__e *ControlFlow) Obj {
return V7183
panic("unreachable")
})

let__8084 := tmp8466
_ = let__8084

tmp8501 := GConsP(__e, V7183)

if True == tmp8501 {
tmp8467 := GHd(__e, V7183)

let__8085 := tmp8467
_ = let__8085

tmp8468 := GTl(__e, V7183)

let__8086 := tmp8468
_ = let__8086

tmp8499 := GEq(__e, sym_7, let__8085)

var ifres8494 Obj

if True == tmp8499 {
tmp8498 := GConsP(__e, let__8086)

var ifres8495 Obj

if True == tmp8498 {
tmp8496 := GTl(__e, let__8086)

tmp8497 := GEq(__e, Nil, tmp8496)

ifres8495 = tmp8497


} else {
ifres8495 = False


}

ifres8494 = ifres8495


} else {
ifres8494 = False


}

if True == ifres8494 {
tmp8469 := GHd(__e, let__8086)

tmp8470 := Call1(__e, SymFn(symshen_4cons_1form_1respect_1modes), tmp8469)


tmp8471 := GCons(__e, tmp8470, Nil)

return GCons(__e, sym_7, tmp8471)


} else {
tmp8492 := GEq(__e, sym_1, let__8085)

var ifres8487 Obj

if True == tmp8492 {
tmp8491 := GConsP(__e, let__8086)

var ifres8488 Obj

if True == tmp8491 {
tmp8489 := GTl(__e, let__8086)

tmp8490 := GEq(__e, Nil, tmp8489)

ifres8488 = tmp8490


} else {
ifres8488 = False


}

ifres8487 = ifres8488


} else {
ifres8487 = False


}

if True == ifres8487 {
tmp8472 := GHd(__e, let__8086)

tmp8473 := Call1(__e, SymFn(symshen_4cons_1form_1respect_1modes), tmp8472)


tmp8474 := GCons(__e, tmp8473, Nil)

return GCons(__e, sym_1, tmp8474)


} else {
tmp8485 := GEq(__e, symbar_b, let__8085)

var ifres8480 Obj

if True == tmp8485 {
tmp8484 := GConsP(__e, let__8086)

var ifres8481 Obj

if True == tmp8484 {
tmp8482 := GTl(__e, let__8086)

tmp8483 := GEq(__e, Nil, tmp8482)

ifres8481 = tmp8483


} else {
ifres8481 = False


}

ifres8480 = ifres8481


} else {
ifres8480 = False


}

if True == ifres8480 {
return GHd(__e, let__8086)
} else {
tmp8475 := Call1(__e, SymFn(symshen_4cons_1form_1respect_1modes), let__8085)


tmp8476 := Call1(__e, SymFn(symshen_4cons_1form_1respect_1modes), let__8086)


tmp8477 := GCons(__e, tmp8476, Nil)

tmp8478 := GCons(__e, tmp8475, tmp8477)

return GCons(__e, symcons, tmp8478)


}


}


}


} else {
return TailCall1(__e, SymFn(symthaw), let__8084)
}


panic("unreachable")
})

tmp8502 := Call2(__e, ns2_1set, symshen_4cons_1form_1respect_1modes, tmp8465)


_ = tmp8502

tmp8503 := MakeNative2(func(__e *ControlFlow, V7187, V7188 Obj) Obj {
tmp8504 := GCons(__e, symX, Nil)

tmp8505 := GCons(__e, sym_1_6, tmp8504)

tmp8506 := GCons(__e, symX, tmp8505)

let__8087 := tmp8506
_ = let__8087

tmp8507 := Call2(__e, SymFn(symappend), V7188, let__8087)


tmp8508 := GCons(__e, V7187, tmp8507)

tmp8509 := GCons(__e, symdefine, tmp8508)

tmp8510 := Call1(__e, SymFn(symeval), tmp8509)


let__8088 := tmp8510
_ = let__8088

tmp8511 := Call1(__e, SymFn(symfn), V7187)


tmp8512 := Call2(__e, SymFn(symshen_4record_1macro), V7187, tmp8511)


let__8089 := tmp8512
_ = let__8089

return V7187


panic("unreachable")
})

tmp8513 := Call2(__e, ns2_1set, symshen_4process_1def, tmp8503)


_ = tmp8513

tmp8514 := MakeNative1(func(__e *ControlFlow, V7192 Obj) Obj {
tmp8541 := GConsP(__e, V7192)

var ifres8526 Obj

if True == tmp8541 {
tmp8527 := GTl(__e, V7192)

let__8090 := tmp8527
_ = let__8090

tmp8539 := GHd(__e, V7192)

tmp8540 := GEq(__e, symlet, tmp8539)

var ifres8528 Obj

if True == tmp8540 {
tmp8538 := GConsP(__e, let__8090)

var ifres8529 Obj

if True == tmp8538 {
tmp8530 := GTl(__e, let__8090)

let__8091 := tmp8530
_ = let__8091

tmp8537 := GConsP(__e, let__8091)

var ifres8531 Obj

if True == tmp8537 {
tmp8532 := GTl(__e, let__8091)

let__8092 := tmp8532
_ = let__8092

tmp8536 := GConsP(__e, let__8092)

var ifres8533 Obj

if True == tmp8536 {
tmp8534 := GTl(__e, let__8092)

tmp8535 := GConsP(__e, tmp8534)

ifres8533 = tmp8535


} else {
ifres8533 = False


}

ifres8531 = ifres8533


} else {
ifres8531 = False


}

ifres8529 = ifres8531


} else {
ifres8529 = False


}

ifres8528 = ifres8529


} else {
ifres8528 = False


}

ifres8526 = ifres8528


} else {
ifres8526 = False


}

if True == ifres8526 {
tmp8515 := GTl(__e, V7192)

let__8093 := tmp8515
_ = let__8093

tmp8516 := GHd(__e, let__8093)

tmp8517 := GTl(__e, let__8093)

tmp8518 := GHd(__e, tmp8517)

tmp8519 := GTl(__e, let__8093)

tmp8520 := GTl(__e, tmp8519)

tmp8521 := GCons(__e, symlet, tmp8520)

tmp8522 := GCons(__e, tmp8521, Nil)

tmp8523 := GCons(__e, tmp8518, tmp8522)

tmp8524 := GCons(__e, tmp8516, tmp8523)

return GCons(__e, symlet, tmp8524)


} else {
return V7192
}


panic("unreachable")
})

tmp8542 := Call2(__e, ns2_1set, symshen_4process_1let, tmp8514)


_ = tmp8542

tmp8543 := (func() Obj {
var __self Obj
__self = MakeNative1(func(__e *ControlFlow, __self1 Obj) Obj {
__selftop:
V7197 := __self1
_ = V7197
tmp8544 := MakeNative0(func(__e *ControlFlow) Obj {
return V7197
panic("unreachable")
})

let__8094 := tmp8544
_ = let__8094

tmp8573 := GConsP(__e, V7197)

if True == tmp8573 {
tmp8545 := GTl(__e, V7197)

let__8095 := tmp8545
_ = let__8095

tmp8570 := GHd(__e, V7197)

tmp8571 := GEq(__e, sym_8s, tmp8570)

if True == tmp8571 {
tmp8568 := GConsP(__e, let__8095)

if True == tmp8568 {
tmp8546 := GHd(__e, let__8095)

let__8096 := tmp8546
_ = let__8096

tmp8547 := GTl(__e, let__8095)

let__8097 := tmp8547
_ = let__8097

tmp8566 := GConsP(__e, let__8097)

if True == tmp8566 {
tmp8548 := GTl(__e, let__8097)

let__8098 := tmp8548
_ = let__8098

tmp8564 := GConsP(__e, let__8098)

if True == tmp8564 {
tmp8549 := GCons(__e, sym_8s, let__8097)

tmp8550 := Call1(__e, SymFn(symshen_4process_1_8s), tmp8549)


tmp8551 := GCons(__e, tmp8550, Nil)

tmp8552 := GCons(__e, let__8096, tmp8551)

return GCons(__e, sym_8s, tmp8552)


} else {
tmp8562 := GEq(__e, Nil, let__8098)

var ifres8560 Obj

if True == tmp8562 {
tmp8561 := PrimGuard1(__e, symstring_2, PrimIsString, let__8096)

ifres8560 = tmp8561


} else {
ifres8560 = False


}

if True == ifres8560 {
tmp8553 := Call1(__e, SymFn(symexplode), let__8096)


let__8099 := tmp8553
_ = let__8099

tmp8557 := Call1(__e, SymFn(symlength), let__8099)


if True == GGt(__e, tmp8557, MakeInteger(1)) {
tmp8554 := Call2(__e, SymFn(symappend), let__8099, let__8097)


tmp8555 := GCons(__e, sym_8s, tmp8554)

if SymFn(symshen_4process_1_8s) == __self {
__self1 = tmp8555
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4process_1_8s), tmp8555)


} else {
return V7197
}


} else {
return TailCall1(__e, SymFn(symthaw), let__8094)
}


}


} else {
return TailCall1(__e, SymFn(symthaw), let__8094)
}


} else {
return TailCall1(__e, SymFn(symthaw), let__8094)
}


} else {
return TailCall1(__e, SymFn(symthaw), let__8094)
}


} else {
return TailCall1(__e, SymFn(symthaw), let__8094)
}


panic("unreachable")
})
return __self
})()

tmp8574 := Call2(__e, ns2_1set, symshen_4process_1_8s, tmp8543)


_ = tmp8574

tmp8575 := MakeNative2(func(__e *ControlFlow, V7207, V7208 Obj) Obj {
tmp8576 := Call1(__e, SymFn(symshen_4intern_1type), V7207)


let__8100 := tmp8576
_ = let__8100

tmp8577 := MakeNative1(func(__e *ControlFlow, Z7211 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4_5datatype_6), Z7211)
panic("unreachable")
})

tmp8578 := GCons(__e, let__8100, V7208)

tmp8579 := Call2(__e, SymFn(symcompile), tmp8577, tmp8578)


let__8101 := tmp8579
_ = let__8101

return let__8100


panic("unreachable")
})

tmp8580 := Call2(__e, ns2_1set, symshen_4process_1datatype, tmp8575)


_ = tmp8580

tmp8581 := MakeNative1(func(__e *ControlFlow, V7212 Obj) Obj {
tmp8582 := PrimGuard1(__e, symstr, PrimStr, V7212)

return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(PrimGuard2(__e, symcn, PrimStringConcat, tmp8582, MakeString("#type")))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, tmp8582, MakeString("#type"))
return Call1(__e, SymFn(symintern), __typedArg0)
})()


panic("unreachable")
})

tmp8584 := Call2(__e, ns2_1set, symshen_4intern_1type, tmp8581)


_ = tmp8584

tmp8585 := MakeNative1(func(__e *ControlFlow, V7213 Obj) Obj {
tmp8586 := GValue(__e, symshen_4_dsynonyms_d)

tmp8587 := Call2(__e, SymFn(symappend), V7213, tmp8586)


tmp8588 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsynonyms_d, tmp8587)

return TailCall1(__e, SymFn(symshen_4synonyms_1h), tmp8588)


panic("unreachable")
})

tmp8589 := Call2(__e, ns2_1set, symshen_4process_1synonyms, tmp8585)


_ = tmp8589

tmp8590 := MakeNative1(func(__e *ControlFlow, V7216 Obj) Obj {
tmp8621 := GConsP(__e, V7216)

var ifres8600 Obj

if True == tmp8621 {
tmp8601 := GTl(__e, V7216)

let__8102 := tmp8601
_ = let__8102

tmp8619 := GHd(__e, V7216)

tmp8620 := GEq(__e, symdefun, tmp8619)

var ifres8602 Obj

if True == tmp8620 {
tmp8618 := GConsP(__e, let__8102)

var ifres8603 Obj

if True == tmp8618 {
tmp8604 := GTl(__e, let__8102)

let__8103 := tmp8604
_ = let__8103

tmp8617 := GConsP(__e, let__8103)

var ifres8605 Obj

if True == tmp8617 {
tmp8606 := GHd(__e, let__8103)

let__8104 := tmp8606
_ = let__8104

tmp8607 := GTl(__e, let__8103)

let__8105 := tmp8607
_ = let__8105

tmp8616 := GConsP(__e, let__8104)

var ifres8608 Obj

if True == tmp8616 {
tmp8614 := GTl(__e, let__8104)

tmp8615 := GEq(__e, Nil, tmp8614)

var ifres8609 Obj

if True == tmp8615 {
tmp8613 := GConsP(__e, let__8105)

var ifres8610 Obj

if True == tmp8613 {
tmp8611 := GTl(__e, let__8105)

tmp8612 := GEq(__e, Nil, tmp8611)

ifres8610 = tmp8612


} else {
ifres8610 = False


}

ifres8609 = ifres8610


} else {
ifres8609 = False


}

ifres8608 = ifres8609


} else {
ifres8608 = False


}

ifres8605 = ifres8608


} else {
ifres8605 = False


}

ifres8603 = ifres8605


} else {
ifres8603 = False


}

ifres8602 = ifres8603


} else {
ifres8602 = False


}

ifres8600 = ifres8602


} else {
ifres8600 = False


}

if True == ifres8600 {
tmp8591 := GTl(__e, V7216)

let__8106 := tmp8591
_ = let__8106

tmp8592 := GTl(__e, let__8106)

tmp8593 := GHd(__e, tmp8592)

tmp8594 := GHd(__e, tmp8593)

tmp8595 := GTl(__e, let__8106)

tmp8596 := GTl(__e, tmp8595)

tmp8597 := GCons(__e, tmp8594, tmp8596)

tmp8598 := GCons(__e, sym_c_4, tmp8597)

return TailCall1(__e, SymFn(symeval), tmp8598)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4lambda_1of_1defun)
}


panic("unreachable")
})

tmp8622 := Call2(__e, ns2_1set, symshen_4lambda_1of_1defun, tmp8590)


_ = tmp8622

tmp8623 := MakeNative1(func(__e *ControlFlow, V7222 Obj) Obj {
tmp8624 := MakeNative1(func(__e *ControlFlow, Z7224 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4curry_1type), Z7224)
panic("unreachable")
})

tmp8625 := Call2(__e, SymFn(symmap), tmp8624, V7222)


let__8107 := tmp8625
_ = let__8107

tmp8626 := Call1(__e, SymFn(symshen_4compile_1synonyms), let__8107)


tmp8627 := Call2(__e, SymFn(symshen_4shendef_1_6kldef), symshen_4demod, tmp8626)


tmp8628 := Call1(__e, SymFn(symshen_4lambda_1of_1defun), tmp8627)


let__8108 := tmp8628
_ = let__8108

tmp8629 := PrimGuard2(__e, symset, PrimSet, symshen_4_ddemodulation_1function_d, let__8108)

let__8109 := tmp8629
_ = let__8109

return symsynonyms


panic("unreachable")
})

tmp8630 := Call2(__e, ns2_1set, symshen_4synonyms_1h, tmp8623)


_ = tmp8630

tmp8631 := MakeNative1(func(__e *ControlFlow, V7229 Obj) Obj {
tmp8650 := GEq(__e, Nil, V7229)

if True == tmp8650 {
tmp8632 := Call1(__e, SymFn(symgensym), symX)


let__8110 := tmp8632
_ = let__8110

tmp8633 := GCons(__e, let__8110, Nil)

tmp8634 := GCons(__e, sym_1_6, tmp8633)

return GCons(__e, let__8110, tmp8634)


} else {
tmp8648 := GConsP(__e, V7229)

var ifres8645 Obj

if True == tmp8648 {
tmp8646 := GTl(__e, V7229)

tmp8647 := GConsP(__e, tmp8646)

ifres8645 = tmp8647


} else {
ifres8645 = False


}

if True == ifres8645 {
tmp8635 := GTl(__e, V7229)

let__8111 := tmp8635
_ = let__8111

tmp8636 := GHd(__e, V7229)

tmp8637 := Call1(__e, SymFn(symshen_4rcons__form), tmp8636)


tmp8638 := GHd(__e, let__8111)

tmp8639 := Call1(__e, SymFn(symshen_4rcons__form), tmp8638)


tmp8640 := GTl(__e, let__8111)

tmp8641 := Call1(__e, SymFn(symshen_4compile_1synonyms), tmp8640)


tmp8642 := GCons(__e, tmp8639, tmp8641)

tmp8643 := GCons(__e, sym_1_6, tmp8642)

return GCons(__e, tmp8637, tmp8643)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("synonyms requires an even number of arguments\n"))
}


}


panic("unreachable")
})

tmp8651 := Call2(__e, ns2_1set, symshen_4compile_1synonyms, tmp8631)


_ = tmp8651

tmp8652 := MakeNative1(func(__e *ControlFlow, V7232 Obj) Obj {
tmp8653 := MakeNative0(func(__e *ControlFlow) Obj {
return V7232
panic("unreachable")
})

let__8112 := tmp8653
_ = let__8112

tmp8677 := GConsP(__e, V7232)

if True == tmp8677 {
tmp8654 := GTl(__e, V7232)

let__8113 := tmp8654
_ = let__8113

tmp8674 := GHd(__e, V7232)

tmp8675 := GEq(__e, sym_c_4, tmp8674)

if True == tmp8675 {
tmp8672 := GConsP(__e, let__8113)

if True == tmp8672 {
tmp8655 := GHd(__e, let__8113)

let__8114 := tmp8655
_ = let__8114

tmp8656 := GTl(__e, let__8113)

let__8115 := tmp8656
_ = let__8115

tmp8670 := GConsP(__e, let__8115)

if True == tmp8670 {
tmp8657 := GTl(__e, let__8115)

let__8116 := tmp8657
_ = let__8116

tmp8668 := GConsP(__e, let__8116)

if True == tmp8668 {
tmp8658 := GCons(__e, sym_c_4, let__8115)

tmp8659 := Call1(__e, SymFn(symshen_4process_1lambda), tmp8658)


tmp8660 := GCons(__e, tmp8659, Nil)

tmp8661 := GCons(__e, let__8114, tmp8660)

return GCons(__e, symlambda, tmp8661)


} else {
tmp8666 := GEq(__e, Nil, let__8116)

if True == tmp8666 {
tmp8664 := PrimGuard1(__e, symvariable_2, PrimIsVariable, let__8114)

if True == tmp8664 {
return GCons(__e, symlambda, let__8113)
} else {
tmp8662 := Call3(__e, SymFn(symshen_4app), let__8114, MakeString(" is not a variable\n"), symshen_4s)


return PrimGuard1(__e, symsimple_1error, PrimSimpleError, tmp8662)


}


} else {
return TailCall1(__e, SymFn(symthaw), let__8112)
}


}


} else {
return TailCall1(__e, SymFn(symthaw), let__8112)
}


} else {
return TailCall1(__e, SymFn(symthaw), let__8112)
}


} else {
return TailCall1(__e, SymFn(symthaw), let__8112)
}


} else {
return TailCall1(__e, SymFn(symthaw), let__8112)
}


panic("unreachable")
})

tmp8678 := Call2(__e, ns2_1set, symshen_4process_1lambda, tmp8652)


_ = tmp8678

tmp8679 := MakeNative1(func(__e *ControlFlow, V7243 Obj) Obj {
tmp8680 := MakeNative0(func(__e *ControlFlow) Obj {
return V7243
panic("unreachable")
})

let__8117 := tmp8680
_ = let__8117

tmp8713 := GConsP(__e, V7243)

if True == tmp8713 {
tmp8681 := GTl(__e, V7243)

let__8118 := tmp8681
_ = let__8118

tmp8710 := GHd(__e, V7243)

tmp8711 := GEq(__e, symcases, tmp8710)

if True == tmp8711 {
tmp8708 := GConsP(__e, let__8118)

if True == tmp8708 {
tmp8682 := GHd(__e, let__8118)

let__8119 := tmp8682
_ = let__8119

tmp8683 := GTl(__e, let__8118)

let__8120 := tmp8683
_ = let__8120

tmp8706 := GEq(__e, True, let__8119)

var ifres8704 Obj

if True == tmp8706 {
tmp8705 := GConsP(__e, let__8120)

ifres8704 = tmp8705


} else {
ifres8704 = False


}

if True == ifres8704 {
return GHd(__e, let__8120)
} else {
tmp8684 := MakeNative0(func(__e *ControlFlow) Obj {
tmp8686 := GEq(__e, Nil, let__8120)

if True == tmp8686 {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("error: odd number of case elements\n"))
} else {
return TailCall1(__e, SymFn(symthaw), let__8117)
}


panic("unreachable")
})

let__8121 := tmp8684
_ = let__8121

tmp8702 := GConsP(__e, let__8120)

if True == tmp8702 {
tmp8687 := GHd(__e, let__8120)

let__8122 := tmp8687
_ = let__8122

tmp8688 := GTl(__e, let__8120)

let__8123 := tmp8688
_ = let__8123

tmp8700 := GEq(__e, Nil, let__8123)

if True == tmp8700 {
tmp8689 := GCons(__e, MakeString("error: cases exhausted"), Nil)

tmp8690 := GCons(__e, symsimple_1error, tmp8689)

tmp8691 := GCons(__e, tmp8690, Nil)

tmp8692 := GCons(__e, let__8122, tmp8691)

tmp8693 := GCons(__e, let__8119, tmp8692)

return GCons(__e, symif, tmp8693)


} else {
tmp8694 := GCons(__e, symcases, let__8123)

tmp8695 := Call1(__e, SymFn(symshen_4process_1cases), tmp8694)


tmp8696 := GCons(__e, tmp8695, Nil)

tmp8697 := GCons(__e, let__8122, tmp8696)

tmp8698 := GCons(__e, let__8119, tmp8697)

return GCons(__e, symif, tmp8698)


}


} else {
return TailCall1(__e, SymFn(symthaw), let__8121)
}


}


} else {
return TailCall1(__e, SymFn(symthaw), let__8117)
}


} else {
return TailCall1(__e, SymFn(symthaw), let__8117)
}


} else {
return TailCall1(__e, SymFn(symthaw), let__8117)
}


panic("unreachable")
})

tmp8714 := Call2(__e, ns2_1set, symshen_4process_1cases, tmp8679)


_ = tmp8714

tmp8715 := MakeNative1(func(__e *ControlFlow, V7253 Obj) Obj {
tmp8716 := GCons(__e, symrun, Nil)

tmp8717 := GCons(__e, symget_1time, tmp8716)

tmp8718 := GCons(__e, symrun, Nil)

tmp8719 := GCons(__e, symget_1time, tmp8718)

tmp8720 := GCons(__e, symStart, Nil)

tmp8721 := GCons(__e, symFinish, tmp8720)

tmp8722 := GCons(__e, sym_1, tmp8721)

tmp8723 := GCons(__e, symTime, Nil)

tmp8724 := GCons(__e, symstr, tmp8723)

tmp8725 := GCons(__e, MakeString(" secs\n"), Nil)

tmp8726 := GCons(__e, tmp8724, tmp8725)

tmp8727 := GCons(__e, symcn, tmp8726)

tmp8728 := GCons(__e, tmp8727, Nil)

tmp8729 := GCons(__e, MakeString("\nrun time: "), tmp8728)

tmp8730 := GCons(__e, symcn, tmp8729)

tmp8731 := GCons(__e, symstoutput, Nil)

tmp8732 := GCons(__e, tmp8731, Nil)

tmp8733 := GCons(__e, tmp8730, tmp8732)

tmp8734 := GCons(__e, sympr, tmp8733)

tmp8735 := GCons(__e, symResult, Nil)

tmp8736 := GCons(__e, tmp8734, tmp8735)

tmp8737 := GCons(__e, symMessage, tmp8736)

tmp8738 := GCons(__e, tmp8722, tmp8737)

tmp8739 := GCons(__e, symTime, tmp8738)

tmp8740 := GCons(__e, tmp8719, tmp8739)

tmp8741 := GCons(__e, symFinish, tmp8740)

tmp8742 := GCons(__e, V7253, tmp8741)

tmp8743 := GCons(__e, symResult, tmp8742)

tmp8744 := GCons(__e, tmp8717, tmp8743)

tmp8745 := GCons(__e, symStart, tmp8744)

return GCons(__e, symlet, tmp8745)


panic("unreachable")
})

tmp8746 := Call2(__e, ns2_1set, symshen_4process_1time, tmp8715)


_ = tmp8746

tmp8747 := MakeNative1(func(__e *ControlFlow, V7254 Obj) Obj {
tmp8765 := GConsP(__e, V7254)

var ifres8756 Obj

if True == tmp8765 {
tmp8757 := GTl(__e, V7254)

let__8124 := tmp8757
_ = let__8124

tmp8764 := GConsP(__e, let__8124)

var ifres8758 Obj

if True == tmp8764 {
tmp8759 := GTl(__e, let__8124)

let__8125 := tmp8759
_ = let__8125

tmp8763 := GConsP(__e, let__8125)

var ifres8760 Obj

if True == tmp8763 {
tmp8761 := GTl(__e, let__8125)

tmp8762 := GConsP(__e, tmp8761)

ifres8760 = tmp8762


} else {
ifres8760 = False


}

ifres8758 = ifres8760


} else {
ifres8758 = False


}

ifres8756 = ifres8758


} else {
ifres8756 = False


}

if True == ifres8756 {
tmp8748 := GHd(__e, V7254)

let__8126 := tmp8748
_ = let__8126

tmp8749 := GTl(__e, V7254)

let__8127 := tmp8749
_ = let__8127

tmp8750 := GHd(__e, let__8127)

tmp8751 := GTl(__e, let__8127)

tmp8752 := GCons(__e, let__8126, tmp8751)

tmp8753 := GCons(__e, tmp8752, Nil)

tmp8754 := GCons(__e, tmp8750, tmp8753)

return GCons(__e, let__8126, tmp8754)


} else {
return V7254
}


panic("unreachable")
})

tmp8766 := Call2(__e, ns2_1set, symshen_4process_1assoc, tmp8747)


_ = tmp8766

tmp8767 := MakeNative1(func(__e *ControlFlow, V7259 Obj) Obj {
tmp8768 := PrimGuard1(__e, symstr, PrimStr, V7259)

tmp8769 := Call1(__e, SymFn(symshen_4mu_1h), tmp8768)


return PrimGuard1(__e, symintern, PrimIntern, tmp8769)


panic("unreachable")
})

tmp8770 := Call2(__e, ns2_1set, symshen_4make_1uppercase, tmp8767)


_ = tmp8770

tmp8771 := MakeNative1(func(__e *ControlFlow, V7260 Obj) Obj {
tmp8786 := GEq(__e, MakeString(""), V7260)

if True == tmp8786 {
return MakeString("")
} else {
tmp8784 := Call1(__e, SymFn(symshen_4_7string_2), V7260)


if True == tmp8784 {
tmp8772 := Call1(__e, SymFn(symhdstr), V7260)


tmp8773 := PrimGuard1(__e, symstring_1_6n, PrimStringToNumber, tmp8772)

let__8128 := tmp8773
_ = let__8128

let__8129 := GSub(__e, let__8128, MakeInteger(32))
_ = let__8129

tmp8780 := GGe(__e, let__8128, MakeInteger(97))

var ifres8778 Obj

if True == tmp8780 {
ifres8778 = GLe(__e, let__8128, MakeInteger(122))


} else {
ifres8778 = False


}

var ifres8775 Obj

if True == ifres8778 {
tmp8776 := PrimGuard1(__e, symn_1_6string, PrimNumberToString, let__8129)

ifres8775 = tmp8776


} else {
tmp8777 := Call1(__e, SymFn(symhdstr), V7260)


ifres8775 = tmp8777


}

let__8130 := ifres8775
_ = let__8130

tmp8782 := Call1(__e, SymFn(symshen_4mu_1h), PrimGuard1(__e, symtlstr, PrimTailString, V7260))


return TailCall2(__e, SymFn(sym_8s), let__8130, tmp8782)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4mu_1h)
}


}


panic("unreachable")
})

tmp8787 := Call2(__e, ns2_1set, symshen_4mu_1h, tmp8771)


_ = tmp8787

tmp8788 := MakeNative2(func(__e *ControlFlow, V7264, V7265 Obj) Obj {
tmp8789 := GValue(__e, sym_dmacros_d)

tmp8790 := Call3(__e, SymFn(symshen_4update_1assoc), V7264, V7265, tmp8789)


return PrimGuard2(__e, symset, PrimSet, sym_dmacros_d, tmp8790)


panic("unreachable")
})

tmp8791 := Call2(__e, ns2_1set, symshen_4record_1macro, tmp8788)


_ = tmp8791

tmp8792 := MakeNative3(func(__e *ControlFlow, V7275, V7276, V7277 Obj) Obj {
tmp8808 := GEq(__e, Nil, V7277)

if True == tmp8808 {
tmp8793 := GCons(__e, V7275, V7276)

return GCons(__e, tmp8793, Nil)


} else {
tmp8794 := MakeNative0(func(__e *ControlFlow) Obj {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.update-assoc"))
panic("unreachable")
})

let__8131 := tmp8794
_ = let__8131

tmp8806 := GConsP(__e, V7277)

if True == tmp8806 {
tmp8795 := GHd(__e, V7277)

let__8132 := tmp8795
_ = let__8132

tmp8796 := GTl(__e, V7277)

let__8133 := tmp8796
_ = let__8133

tmp8804 := GConsP(__e, let__8132)

var ifres8801 Obj

if True == tmp8804 {
tmp8802 := GHd(__e, let__8132)

tmp8803 := GEq(__e, V7275, tmp8802)

ifres8801 = tmp8803


} else {
ifres8801 = False


}

if True == ifres8801 {
tmp8797 := GHd(__e, let__8132)

tmp8798 := GCons(__e, tmp8797, V7276)

return GCons(__e, tmp8798, let__8133)


} else {
tmp8799 := Call3(__e, SymFn(symshen_4update_1assoc), V7275, V7276, let__8133)


return GCons(__e, let__8132, tmp8799)


}


} else {
return TailCall1(__e, SymFn(symthaw), let__8131)
}


}


panic("unreachable")
})

tmp8809 := Call2(__e, ns2_1set, symshen_4update_1assoc, tmp8792)


_ = tmp8809

tmp8810 := MakeNative0(func(__e *ControlFlow) Obj {
tmp8818 := Call0(__e, SymFn(symstinput))


tmp8819 := Call1(__e, SymFn(symchar_1stream_2), tmp8818)


if True == tmp8819 {
tmp8811 := GCons(__e, symstinput, Nil)

tmp8812 := GCons(__e, tmp8811, Nil)

tmp8813 := GCons(__e, symread_1unit_1string, tmp8812)

tmp8814 := GCons(__e, tmp8813, Nil)

return GCons(__e, symstring_1_6n, tmp8814)


} else {
tmp8815 := GCons(__e, symstinput, Nil)

tmp8816 := GCons(__e, tmp8815, Nil)

return GCons(__e, symread_1byte, tmp8816)


}


panic("unreachable")
})

tmp8820 := Call2(__e, ns2_1set, symshen_4process_1read_1byte, tmp8810)


_ = tmp8820

tmp8821 := MakeNative1(func(__e *ControlFlow, V7281 Obj) Obj {
tmp8822 := GCons(__e, symshen_4prolog_1vector, Nil)

let__8134 := tmp8822
_ = let__8134

tmp8823 := GCons(__e, MakeInteger(0), Nil)

tmp8824 := GCons(__e, symvector, tmp8823)

tmp8825 := GCons(__e, tmp8824, Nil)

tmp8826 := GCons(__e, MakeInteger(0), tmp8825)

tmp8827 := GCons(__e, True, tmp8826)

tmp8828 := GCons(__e, sym_8v, tmp8827)

let__8135 := tmp8828
_ = let__8135

let__8136 := MakeInteger(0)
_ = let__8136

tmp8829 := GCons(__e, True, Nil)

tmp8830 := GCons(__e, symfreeze, tmp8829)

let__8137 := tmp8830
_ = let__8137

tmp8831 := MakeNative1(func(__e *ControlFlow, Z7287 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4_5body_6), Z7287)
panic("unreachable")
})

tmp8832 := Call2(__e, SymFn(symcompile), tmp8831, V7281)


let__8138 := tmp8832
_ = let__8138

tmp8833 := Call1(__e, SymFn(symshen_4received), V7281)


let__8139 := tmp8833
_ = let__8139

tmp8834 := Call1(__e, SymFn(symgensym), symV)


let__8140 := tmp8834
_ = let__8140

tmp8835 := Call1(__e, SymFn(symgensym), symL)


let__8141 := tmp8835
_ = let__8141

tmp8836 := Call1(__e, SymFn(symgensym), symK)


let__8142 := tmp8836
_ = let__8142

tmp8837 := Call1(__e, SymFn(symgensym), symC)


let__8143 := tmp8837
_ = let__8143

tmp8838 := Call6(__e, SymFn(symshen_4continue), let__8139, let__8138, let__8140, let__8141, let__8142, let__8143)


tmp8839 := GCons(__e, tmp8838, Nil)

tmp8840 := GCons(__e, let__8143, tmp8839)

tmp8841 := GCons(__e, symlambda, tmp8840)

tmp8842 := GCons(__e, tmp8841, Nil)

tmp8843 := GCons(__e, let__8142, tmp8842)

tmp8844 := GCons(__e, symlambda, tmp8843)

tmp8845 := GCons(__e, tmp8844, Nil)

tmp8846 := GCons(__e, let__8141, tmp8845)

tmp8847 := GCons(__e, symlambda, tmp8846)

tmp8848 := GCons(__e, tmp8847, Nil)

tmp8849 := GCons(__e, let__8140, tmp8848)

tmp8850 := GCons(__e, symlambda, tmp8849)

let__8144 := tmp8850
_ = let__8144

tmp8851 := GCons(__e, let__8137, Nil)

tmp8852 := GCons(__e, let__8136, tmp8851)

tmp8853 := GCons(__e, let__8135, tmp8852)

tmp8854 := GCons(__e, let__8134, tmp8853)

return GCons(__e, let__8144, tmp8854)


panic("unreachable")
})

tmp8855 := Call2(__e, ns2_1set, symshen_4call_1prolog, tmp8821)


_ = tmp8855

tmp8856 := MakeNative1(func(__e *ControlFlow, V7296 Obj) Obj {
tmp8857 := MakeNative0(func(__e *ControlFlow) Obj {
return Nil
panic("unreachable")
})

let__8145 := tmp8857
_ = let__8145

tmp8870 := GConsP(__e, V7296)

if True == tmp8870 {
tmp8858 := GHd(__e, V7296)

let__8146 := tmp8858
_ = let__8146

tmp8859 := GTl(__e, V7296)

let__8147 := tmp8859
_ = let__8147

tmp8868 := GEq(__e, symreceive, let__8146)

var ifres8863 Obj

if True == tmp8868 {
tmp8867 := GConsP(__e, let__8147)

var ifres8864 Obj

if True == tmp8867 {
tmp8865 := GTl(__e, let__8147)

tmp8866 := GEq(__e, Nil, tmp8865)

ifres8864 = tmp8866


} else {
ifres8864 = False


}

ifres8863 = ifres8864


} else {
ifres8863 = False


}

if True == ifres8863 {
return let__8147
} else {
tmp8860 := Call1(__e, SymFn(symshen_4received), let__8146)


tmp8861 := Call1(__e, SymFn(symshen_4received), let__8147)


return TailCall2(__e, SymFn(symunion), tmp8860, tmp8861)


}


} else {
return TailCall1(__e, SymFn(symthaw), let__8145)
}


panic("unreachable")
})

tmp8871 := Call2(__e, ns2_1set, symshen_4received, tmp8856)


_ = tmp8871

tmp8872 := MakeNative0(func(__e *ControlFlow) Obj {
tmp8873 := GValue(__e, symshen_4_dprolog_1memory_d)

tmp8874 := PrimGuard1(__e, symabsvector, PrimAbsvector, tmp8873)

let__8148 := tmp8874
_ = let__8148

tmp8875 := GVecSet(__e, let__8148, MakeInteger(0), symshen_4print_1prolog_1vector)

let__8149 := tmp8875
_ = let__8149

tmp8876 := GVecSet(__e, let__8148, MakeInteger(1), MakeInteger(2))

let__8150 := tmp8876
_ = let__8150

return let__8150


panic("unreachable")
})

tmp8877 := Call2(__e, ns2_1set, symshen_4prolog_1vector, tmp8872)


_ = tmp8877

tmp8878 := MakeNative1(func(__e *ControlFlow, V7303 Obj) Obj {
return V7303
panic("unreachable")
})

tmp8879 := Call2(__e, ns2_1set, symreceive, tmp8878)


_ = tmp8879

tmp8880 := MakeNative1(func(__e *ControlFlow, V7304 Obj) Obj {
tmp8881 := MakeNative0(func(__e *ControlFlow) Obj {
return V7304
panic("unreachable")
})

let__8151 := tmp8881
_ = let__8151

tmp8896 := GConsP(__e, V7304)

if True == tmp8896 {
tmp8882 := GHd(__e, V7304)

let__8152 := tmp8882
_ = let__8152

tmp8883 := GTl(__e, V7304)

let__8153 := tmp8883
_ = let__8153

tmp8894 := GEq(__e, symbar_b, let__8152)

var ifres8889 Obj

if True == tmp8894 {
tmp8893 := GConsP(__e, let__8153)

var ifres8890 Obj

if True == tmp8893 {
tmp8891 := GTl(__e, let__8153)

tmp8892 := GEq(__e, Nil, tmp8891)

ifres8890 = tmp8892


} else {
ifres8890 = False


}

ifres8889 = ifres8890


} else {
ifres8889 = False


}

if True == ifres8889 {
return GHd(__e, let__8153)
} else {
tmp8884 := Call1(__e, SymFn(symshen_4rcons__form), let__8152)


tmp8885 := Call1(__e, SymFn(symshen_4rcons__form), let__8153)


tmp8886 := GCons(__e, tmp8885, Nil)

tmp8887 := GCons(__e, tmp8884, tmp8886)

return GCons(__e, symcons, tmp8887)


}


} else {
return TailCall1(__e, SymFn(symthaw), let__8151)
}


panic("unreachable")
})

tmp8897 := Call2(__e, ns2_1set, symshen_4rcons__form, tmp8880)


_ = tmp8897

tmp8898 := MakeNative1(func(__e *ControlFlow, V7308 Obj) Obj {
tmp8905 := GConsP(__e, V7308)

if True == tmp8905 {
tmp8899 := GHd(__e, V7308)

tmp8900 := GTl(__e, V7308)

tmp8901 := Call1(__e, SymFn(symshen_4tuple_1up), tmp8900)


tmp8902 := GCons(__e, tmp8901, Nil)

tmp8903 := GCons(__e, tmp8899, tmp8902)

return GCons(__e, sym_8p, tmp8903)


} else {
return V7308
}


panic("unreachable")
})

tmp8906 := Call2(__e, ns2_1set, symshen_4tuple_1up, tmp8898)


_ = tmp8906

tmp8907 := MakeNative1(func(__e *ControlFlow, V7309 Obj) Obj {
tmp8908 := GValue(__e, sym_dmacros_d)

tmp8909 := Call2(__e, SymFn(symassoc), V7309, tmp8908)


tmp8910 := GValue(__e, sym_dmacros_d)

tmp8911 := Call2(__e, SymFn(symremove), tmp8909, tmp8910)


tmp8912 := PrimGuard2(__e, symset, PrimSet, sym_dmacros_d, tmp8911)

_ = tmp8912

return V7309


panic("unreachable")
})

return TailCall2(__e, ns2_1set, symundefmacro, tmp8907)




panic("unreachable")
})

