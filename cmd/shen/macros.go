package main

import . "github.com/tiancaiamao/shen-go/kl"

var MacrosMain = MakeNative(func(__e *ControlFlow) {
tmp7112 := MakeNative(func(__e *ControlFlow) {
V6939 := __e.Get(1)
_ = V6939
tmp7113 := MakeNative(func(__e *ControlFlow) {
W6940 := __e.Get(1)
_ = W6940
__e.TailApply(PrimFunc(symshen_4macroexpand_1h), V6939, W6940, W6940)
return
}, 1)

tmp7114 := MakeNative(func(__e *ControlFlow) {
Z6941 := __e.Get(1)
_ = Z6941
__e.Return(PrimTail(Z6941))
return
}, 1)

tmp7115 := PrimValue(sym_dmacros_d)

tmp7116 := Call(__e, PrimFunc(symmap), tmp7114, tmp7115)


__e.TailApply(tmp7113, tmp7116)
return


}, 1)

tmp7117 := Call(__e, ns2_1set, symmacroexpand, tmp7112)


_ = tmp7117

tmp7118 := MakeNative(func(__e *ControlFlow) {
V6950 := __e.Get(1)
_ = V6950
V6951 := __e.Get(2)
_ = V6951
V6952 := __e.Get(3)
_ = V6952
tmp7128 := PrimEqual(Nil, V6951)

if True == tmp7128 {
__e.Return(V6950)
return
} else {
tmp7126 := PrimIsPair(V6951)

if True == tmp7126 {
tmp7119 := MakeNative(func(__e *ControlFlow) {
W6953 := __e.Get(1)
_ = W6953
tmp7122 := PrimEqual(V6950, W6953)

if True == tmp7122 {
tmp7120 := PrimTail(V6951)

__e.TailApply(PrimFunc(symshen_4macroexpand_1h), V6950, tmp7120, V6952)
return


} else {
__e.TailApply(PrimFunc(symshen_4macroexpand_1h), W6953, V6952, V6952)
return
}


}, 1)

tmp7123 := PrimHead(V6951)

tmp7124 := Call(__e, PrimFunc(symshen_4walk), tmp7123, V6950)


__e.TailApply(tmp7119, tmp7124)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.macroexpand-h")))
return
}


}


}, 3)

tmp7129 := Call(__e, ns2_1set, symshen_4macroexpand_1h, tmp7118)


_ = tmp7129

tmp7130 := MakeNative(func(__e *ControlFlow) {
V6954 := __e.Get(1)
_ = V6954
V6955 := __e.Get(2)
_ = V6955
tmp7134 := PrimIsPair(V6955)

if True == tmp7134 {
tmp7131 := MakeNative(func(__e *ControlFlow) {
Z6956 := __e.Get(1)
_ = Z6956
__e.TailApply(PrimFunc(symshen_4walk), V6954, Z6956)
return
}, 1)

tmp7132 := Call(__e, PrimFunc(symmap), tmp7131, V6955)


__e.TailApply(V6954, tmp7132)
return


} else {
__e.TailApply(V6954, V6955)
return
}


}, 2)

tmp7135 := Call(__e, ns2_1set, symshen_4walk, tmp7130)


_ = tmp7135

tmp7136 := MakeNative(func(__e *ControlFlow) {
V6957 := __e.Get(1)
_ = V6957
tmp7137 := MakeNative(func(__e *ControlFlow) {
GoTo6958 := __e.Get(1)
_ = GoTo6958
tmp7448 := PrimIsPair(V6957)

if True == tmp7448 {
tmp7138 := MakeNative(func(__e *ControlFlow) {
Select6963 := __e.Get(1)
_ = Select6963
tmp7139 := MakeNative(func(__e *ControlFlow) {
Select6964 := __e.Get(1)
_ = Select6964
tmp7444 := PrimEqual(symdefmacro, Select6963)

var ifres7441 Obj

if True == tmp7444 {
tmp7443 := PrimIsPair(Select6964)

var ifres7442 Obj

if True == tmp7443 {
ifres7442 = True


} else {
ifres7442 = False


}

ifres7441 = ifres7442


} else {
ifres7441 = False


}

if True == ifres7441 {
tmp7140 := PrimHead(Select6964)

tmp7141 := PrimTail(Select6964)

__e.TailApply(PrimFunc(symshen_4process_1def), tmp7140, tmp7141)
return


} else {
tmp7439 := PrimEqual(symdefcc, Select6963)

if True == tmp7439 {
__e.TailApply(PrimFunc(symshen_4yacc_1_6shen), Select6964)
return
} else {
tmp7437 := PrimEqual(symu_b, Select6963)

var ifres7430 Obj

if True == tmp7437 {
tmp7436 := PrimIsPair(Select6964)

var ifres7432 Obj

if True == tmp7436 {
tmp7434 := PrimTail(Select6964)

tmp7435 := PrimEqual(Nil, tmp7434)

var ifres7433 Obj

if True == tmp7435 {
ifres7433 = True


} else {
ifres7433 = False


}

ifres7432 = ifres7433


} else {
ifres7432 = False


}

var ifres7431 Obj

if True == ifres7432 {
ifres7431 = True


} else {
ifres7431 = False


}

ifres7430 = ifres7431


} else {
ifres7430 = False


}

if True == ifres7430 {
tmp7142 := PrimHead(Select6964)

tmp7143 := Call(__e, PrimFunc(symshen_4make_1uppercase), tmp7142)


tmp7144 := PrimCons(tmp7143, Nil)

__e.Return(PrimCons(symprotect, tmp7144))
return


} else {
tmp7428 := PrimEqual(symerror, Select6963)

var ifres7425 Obj

if True == tmp7428 {
tmp7427 := PrimIsPair(Select6964)

var ifres7426 Obj

if True == tmp7427 {
ifres7426 = True


} else {
ifres7426 = False


}

ifres7425 = ifres7426


} else {
ifres7425 = False


}

if True == ifres7425 {
tmp7145 := PrimHead(Select6964)

tmp7146 := PrimTail(Select6964)

tmp7147 := Call(__e, PrimFunc(symshen_4mkstr), tmp7145, tmp7146)


tmp7148 := PrimCons(tmp7147, Nil)

__e.Return(PrimCons(symsimple_1error, tmp7148))
return


} else {
tmp7423 := PrimEqual(symoutput, Select6963)

var ifres7420 Obj

if True == tmp7423 {
tmp7422 := PrimIsPair(Select6964)

var ifres7421 Obj

if True == tmp7422 {
ifres7421 = True


} else {
ifres7421 = False


}

ifres7420 = ifres7421


} else {
ifres7420 = False


}

if True == ifres7420 {
tmp7149 := PrimHead(Select6964)

tmp7150 := PrimTail(Select6964)

tmp7151 := Call(__e, PrimFunc(symshen_4mkstr), tmp7149, tmp7150)


tmp7152 := PrimCons(symstoutput, Nil)

tmp7153 := PrimCons(tmp7152, Nil)

tmp7154 := PrimCons(tmp7151, tmp7153)

__e.Return(PrimCons(sympr, tmp7154))
return


} else {
tmp7418 := PrimEqual(sympr, Select6963)

var ifres7411 Obj

if True == tmp7418 {
tmp7417 := PrimIsPair(Select6964)

var ifres7413 Obj

if True == tmp7417 {
tmp7415 := PrimTail(Select6964)

tmp7416 := PrimEqual(Nil, tmp7415)

var ifres7414 Obj

if True == tmp7416 {
ifres7414 = True


} else {
ifres7414 = False


}

ifres7413 = ifres7414


} else {
ifres7413 = False


}

var ifres7412 Obj

if True == ifres7413 {
ifres7412 = True


} else {
ifres7412 = False


}

ifres7411 = ifres7412


} else {
ifres7411 = False


}

if True == ifres7411 {
tmp7155 := PrimHead(Select6964)

tmp7156 := PrimCons(symstoutput, Nil)

tmp7157 := PrimCons(tmp7156, Nil)

tmp7158 := PrimCons(tmp7155, tmp7157)

__e.Return(PrimCons(sympr, tmp7158))
return


} else {
tmp7409 := PrimEqual(symmake_1string, Select6963)

var ifres7406 Obj

if True == tmp7409 {
tmp7408 := PrimIsPair(Select6964)

var ifres7407 Obj

if True == tmp7408 {
ifres7407 = True


} else {
ifres7407 = False


}

ifres7406 = ifres7407


} else {
ifres7406 = False


}

if True == ifres7406 {
tmp7159 := PrimHead(Select6964)

tmp7160 := PrimTail(Select6964)

__e.TailApply(PrimFunc(symshen_4mkstr), tmp7159, tmp7160)
return


} else {
tmp7404 := PrimEqual(symlineread, Select6963)

var ifres7401 Obj

if True == tmp7404 {
tmp7403 := PrimEqual(Nil, Select6964)

var ifres7402 Obj

if True == tmp7403 {
ifres7402 = True


} else {
ifres7402 = False


}

ifres7401 = ifres7402


} else {
ifres7401 = False


}

if True == ifres7401 {
tmp7161 := PrimCons(symstinput, Nil)

tmp7162 := PrimCons(tmp7161, Nil)

__e.Return(PrimCons(symlineread, tmp7162))
return


} else {
tmp7399 := PrimEqual(syminput, Select6963)

var ifres7396 Obj

if True == tmp7399 {
tmp7398 := PrimEqual(Nil, Select6964)

var ifres7397 Obj

if True == tmp7398 {
ifres7397 = True


} else {
ifres7397 = False


}

ifres7396 = ifres7397


} else {
ifres7396 = False


}

if True == ifres7396 {
tmp7163 := PrimCons(symstinput, Nil)

tmp7164 := PrimCons(tmp7163, Nil)

__e.Return(PrimCons(syminput, tmp7164))
return


} else {
tmp7394 := PrimEqual(symread, Select6963)

var ifres7391 Obj

if True == tmp7394 {
tmp7393 := PrimEqual(Nil, Select6964)

var ifres7392 Obj

if True == tmp7393 {
ifres7392 = True


} else {
ifres7392 = False


}

ifres7391 = ifres7392


} else {
ifres7391 = False


}

if True == ifres7391 {
tmp7165 := PrimCons(symstinput, Nil)

tmp7166 := PrimCons(tmp7165, Nil)

__e.Return(PrimCons(symread, tmp7166))
return


} else {
tmp7389 := PrimEqual(syminput_7, Select6963)

var ifres7382 Obj

if True == tmp7389 {
tmp7388 := PrimIsPair(Select6964)

var ifres7384 Obj

if True == tmp7388 {
tmp7386 := PrimTail(Select6964)

tmp7387 := PrimEqual(Nil, tmp7386)

var ifres7385 Obj

if True == tmp7387 {
ifres7385 = True


} else {
ifres7385 = False


}

ifres7384 = ifres7385


} else {
ifres7384 = False


}

var ifres7383 Obj

if True == ifres7384 {
ifres7383 = True


} else {
ifres7383 = False


}

ifres7382 = ifres7383


} else {
ifres7382 = False


}

if True == ifres7382 {
tmp7167 := PrimHead(Select6964)

tmp7168 := PrimCons(symstinput, Nil)

tmp7169 := PrimCons(tmp7168, Nil)

tmp7170 := PrimCons(tmp7167, tmp7169)

__e.Return(PrimCons(syminput_7, tmp7170))
return


} else {
tmp7380 := PrimEqual(symread_1byte, Select6963)

var ifres7377 Obj

if True == tmp7380 {
tmp7379 := PrimEqual(Nil, Select6964)

var ifres7378 Obj

if True == tmp7379 {
ifres7378 = True


} else {
ifres7378 = False


}

ifres7377 = ifres7378


} else {
ifres7377 = False


}

if True == ifres7377 {
__e.TailApply(PrimFunc(symshen_4process_1read_1byte))
return
} else {
tmp7375 := PrimEqual(symprolog_2, Select6963)

if True == tmp7375 {
__e.TailApply(PrimFunc(symshen_4call_1prolog), Select6964)
return
} else {
tmp7373 := PrimEqual(symdefprolog, Select6963)

var ifres7370 Obj

if True == tmp7373 {
tmp7372 := PrimIsPair(Select6964)

var ifres7371 Obj

if True == tmp7372 {
ifres7371 = True


} else {
ifres7371 = False


}

ifres7370 = ifres7371


} else {
ifres7370 = False


}

if True == ifres7370 {
tmp7171 := PrimHead(Select6964)

tmp7172 := PrimTail(Select6964)

__e.TailApply(PrimFunc(symshen_4compile_1prolog), tmp7171, tmp7172)
return


} else {
tmp7368 := PrimEqual(symdatatype, Select6963)

var ifres7365 Obj

if True == tmp7368 {
tmp7367 := PrimIsPair(Select6964)

var ifres7366 Obj

if True == tmp7367 {
ifres7366 = True


} else {
ifres7366 = False


}

ifres7365 = ifres7366


} else {
ifres7365 = False


}

if True == ifres7365 {
tmp7173 := PrimHead(Select6964)

tmp7174 := PrimTail(Select6964)

__e.TailApply(PrimFunc(symshen_4process_1datatype), tmp7173, tmp7174)
return


} else {
tmp7363 := PrimEqual(sym_8s, Select6963)

if True == tmp7363 {
__e.TailApply(PrimFunc(symshen_4process_1_8s), V6957)
return
} else {
tmp7361 := PrimEqual(symsynonyms, Select6963)

if True == tmp7361 {
__e.TailApply(PrimFunc(symshen_4process_1synonyms), Select6964)
return
} else {
tmp7359 := PrimEqual(symnl, Select6963)

var ifres7356 Obj

if True == tmp7359 {
tmp7358 := PrimEqual(Nil, Select6964)

var ifres7357 Obj

if True == tmp7358 {
ifres7357 = True


} else {
ifres7357 = False


}

ifres7356 = ifres7357


} else {
ifres7356 = False


}

if True == ifres7356 {
tmp7175 := PrimCons(MakeNumber(1), Nil)

__e.Return(PrimCons(symnl, tmp7175))
return


} else {
tmp7354 := PrimEqual(symlet, Select6963)

if True == tmp7354 {
__e.TailApply(PrimFunc(symshen_4process_1let), V6957)
return
} else {
tmp7352 := PrimEqual(sym_c_4, Select6963)

if True == tmp7352 {
__e.TailApply(PrimFunc(symshen_4process_1lambda), V6957)
return
} else {
tmp7350 := PrimEqual(symcases, Select6963)

if True == tmp7350 {
__e.TailApply(PrimFunc(symshen_4process_1cases), V6957)
return
} else {
tmp7348 := PrimEqual(symtime, Select6963)

var ifres7341 Obj

if True == tmp7348 {
tmp7347 := PrimIsPair(Select6964)

var ifres7343 Obj

if True == tmp7347 {
tmp7345 := PrimTail(Select6964)

tmp7346 := PrimEqual(Nil, tmp7345)

var ifres7344 Obj

if True == tmp7346 {
ifres7344 = True


} else {
ifres7344 = False


}

ifres7343 = ifres7344


} else {
ifres7343 = False


}

var ifres7342 Obj

if True == ifres7343 {
ifres7342 = True


} else {
ifres7342 = False


}

ifres7341 = ifres7342


} else {
ifres7341 = False


}

if True == ifres7341 {
tmp7176 := PrimHead(Select6964)

__e.TailApply(PrimFunc(symshen_4process_1time), tmp7176)
return


} else {
tmp7339 := PrimEqual(symput, Select6963)

var ifres7321 Obj

if True == tmp7339 {
tmp7338 := PrimIsPair(Select6964)

var ifres7323 Obj

if True == tmp7338 {
tmp7336 := PrimTail(Select6964)

tmp7337 := PrimIsPair(tmp7336)

var ifres7325 Obj

if True == tmp7337 {
tmp7333 := PrimTail(Select6964)

tmp7334 := PrimTail(tmp7333)

tmp7335 := PrimIsPair(tmp7334)

var ifres7327 Obj

if True == tmp7335 {
tmp7329 := PrimTail(Select6964)

tmp7330 := PrimTail(tmp7329)

tmp7331 := PrimTail(tmp7330)

tmp7332 := PrimEqual(Nil, tmp7331)

var ifres7328 Obj

if True == tmp7332 {
ifres7328 = True


} else {
ifres7328 = False


}

ifres7327 = ifres7328


} else {
ifres7327 = False


}

var ifres7326 Obj

if True == ifres7327 {
ifres7326 = True


} else {
ifres7326 = False


}

ifres7325 = ifres7326


} else {
ifres7325 = False


}

var ifres7324 Obj

if True == ifres7325 {
ifres7324 = True


} else {
ifres7324 = False


}

ifres7323 = ifres7324


} else {
ifres7323 = False


}

var ifres7322 Obj

if True == ifres7323 {
ifres7322 = True


} else {
ifres7322 = False


}

ifres7321 = ifres7322


} else {
ifres7321 = False


}

if True == ifres7321 {
tmp7177 := PrimHead(Select6964)

tmp7178 := PrimTail(Select6964)

tmp7179 := PrimHead(tmp7178)

tmp7180 := PrimTail(Select6964)

tmp7181 := PrimTail(tmp7180)

tmp7182 := PrimHead(tmp7181)

tmp7183 := PrimCons(sym_dproperty_1vector_d, Nil)

tmp7184 := PrimCons(symvalue, tmp7183)

tmp7185 := PrimCons(tmp7184, Nil)

tmp7186 := PrimCons(tmp7182, tmp7185)

tmp7187 := PrimCons(tmp7179, tmp7186)

tmp7188 := PrimCons(tmp7177, tmp7187)

__e.Return(PrimCons(symput, tmp7188))
return


} else {
tmp7319 := PrimEqual(symget, Select6963)

var ifres7307 Obj

if True == tmp7319 {
tmp7318 := PrimIsPair(Select6964)

var ifres7309 Obj

if True == tmp7318 {
tmp7316 := PrimTail(Select6964)

tmp7317 := PrimIsPair(tmp7316)

var ifres7311 Obj

if True == tmp7317 {
tmp7313 := PrimTail(Select6964)

tmp7314 := PrimTail(tmp7313)

tmp7315 := PrimEqual(Nil, tmp7314)

var ifres7312 Obj

if True == tmp7315 {
ifres7312 = True


} else {
ifres7312 = False


}

ifres7311 = ifres7312


} else {
ifres7311 = False


}

var ifres7310 Obj

if True == ifres7311 {
ifres7310 = True


} else {
ifres7310 = False


}

ifres7309 = ifres7310


} else {
ifres7309 = False


}

var ifres7308 Obj

if True == ifres7309 {
ifres7308 = True


} else {
ifres7308 = False


}

ifres7307 = ifres7308


} else {
ifres7307 = False


}

if True == ifres7307 {
tmp7189 := PrimHead(Select6964)

tmp7190 := PrimTail(Select6964)

tmp7191 := PrimHead(tmp7190)

tmp7192 := PrimCons(sym_dproperty_1vector_d, Nil)

tmp7193 := PrimCons(symvalue, tmp7192)

tmp7194 := PrimCons(tmp7193, Nil)

tmp7195 := PrimCons(tmp7191, tmp7194)

tmp7196 := PrimCons(tmp7189, tmp7195)

__e.Return(PrimCons(symget, tmp7196))
return


} else {
tmp7305 := PrimEqual(symunput, Select6963)

var ifres7293 Obj

if True == tmp7305 {
tmp7304 := PrimIsPair(Select6964)

var ifres7295 Obj

if True == tmp7304 {
tmp7302 := PrimTail(Select6964)

tmp7303 := PrimIsPair(tmp7302)

var ifres7297 Obj

if True == tmp7303 {
tmp7299 := PrimTail(Select6964)

tmp7300 := PrimTail(tmp7299)

tmp7301 := PrimEqual(Nil, tmp7300)

var ifres7298 Obj

if True == tmp7301 {
ifres7298 = True


} else {
ifres7298 = False


}

ifres7297 = ifres7298


} else {
ifres7297 = False


}

var ifres7296 Obj

if True == ifres7297 {
ifres7296 = True


} else {
ifres7296 = False


}

ifres7295 = ifres7296


} else {
ifres7295 = False


}

var ifres7294 Obj

if True == ifres7295 {
ifres7294 = True


} else {
ifres7294 = False


}

ifres7293 = ifres7294


} else {
ifres7293 = False


}

if True == ifres7293 {
tmp7197 := PrimHead(Select6964)

tmp7198 := PrimTail(Select6964)

tmp7199 := PrimHead(tmp7198)

tmp7200 := PrimCons(sym_dproperty_1vector_d, Nil)

tmp7201 := PrimCons(symvalue, tmp7200)

tmp7202 := PrimCons(tmp7201, Nil)

tmp7203 := PrimCons(tmp7199, tmp7202)

tmp7204 := PrimCons(tmp7197, tmp7203)

__e.Return(PrimCons(symunput, tmp7204))
return


} else {
tmp7291 := PrimEqual(symshen_4_8c, Select6963)

var ifres7284 Obj

if True == tmp7291 {
tmp7290 := PrimIsPair(Select6964)

var ifres7286 Obj

if True == tmp7290 {
tmp7288 := PrimTail(Select6964)

tmp7289 := PrimEqual(Nil, tmp7288)

var ifres7287 Obj

if True == tmp7289 {
ifres7287 = True


} else {
ifres7287 = False


}

ifres7286 = ifres7287


} else {
ifres7286 = False


}

var ifres7285 Obj

if True == ifres7286 {
ifres7285 = True


} else {
ifres7285 = False


}

ifres7284 = ifres7285


} else {
ifres7284 = False


}

if True == ifres7284 {
tmp7205 := PrimHead(Select6964)

__e.TailApply(PrimFunc(symshen_4rcons__form), tmp7205)
return


} else {
tmp7206 := MakeNative(func(__e *ControlFlow) {
GoTo6959 := __e.Get(1)
_ = GoTo6959
tmp7253 := PrimEqual(symshen_4_8ch, Select6963)

if True == tmp7253 {
tmp7251 := PrimIsPair(Select6964)

if True == tmp7251 {
tmp7207 := MakeNative(func(__e *ControlFlow) {
Select6961 := __e.Get(1)
_ = Select6961
tmp7208 := MakeNative(func(__e *ControlFlow) {
Select6962 := __e.Get(1)
_ = Select6962
tmp7247 := PrimIsPair(Select6961)

var ifres7223 Obj

if True == tmp7247 {
tmp7245 := PrimTail(Select6961)

tmp7246 := PrimIsPair(tmp7245)

var ifres7225 Obj

if True == tmp7246 {
tmp7242 := PrimTail(Select6961)

tmp7243 := PrimTail(tmp7242)

tmp7244 := PrimIsPair(tmp7243)

var ifres7227 Obj

if True == tmp7244 {
tmp7238 := PrimTail(Select6961)

tmp7239 := PrimTail(tmp7238)

tmp7240 := PrimTail(tmp7239)

tmp7241 := PrimEqual(Nil, tmp7240)

var ifres7229 Obj

if True == tmp7241 {
tmp7237 := PrimEqual(Nil, Select6962)

var ifres7231 Obj

if True == tmp7237 {
tmp7233 := PrimTail(Select6961)

tmp7234 := PrimHead(tmp7233)

tmp7235 := PrimIntern(MakeString(":"))

tmp7236 := PrimEqual(tmp7234, tmp7235)

var ifres7232 Obj

if True == tmp7236 {
ifres7232 = True


} else {
ifres7232 = False


}

ifres7231 = ifres7232


} else {
ifres7231 = False


}

var ifres7230 Obj

if True == ifres7231 {
ifres7230 = True


} else {
ifres7230 = False


}

ifres7229 = ifres7230


} else {
ifres7229 = False


}

var ifres7228 Obj

if True == ifres7229 {
ifres7228 = True


} else {
ifres7228 = False


}

ifres7227 = ifres7228


} else {
ifres7227 = False


}

var ifres7226 Obj

if True == ifres7227 {
ifres7226 = True


} else {
ifres7226 = False


}

ifres7225 = ifres7226


} else {
ifres7225 = False


}

var ifres7224 Obj

if True == ifres7225 {
ifres7224 = True


} else {
ifres7224 = False


}

ifres7223 = ifres7224


} else {
ifres7223 = False


}

if True == ifres7223 {
tmp7209 := PrimHead(Select6961)

tmp7210 := PrimTail(Select6961)

tmp7211 := PrimHead(tmp7210)

tmp7212 := PrimTail(Select6961)

tmp7213 := PrimTail(tmp7212)

tmp7214 := PrimCons(sym_7, tmp7213)

tmp7215 := PrimCons(tmp7214, Nil)

tmp7216 := PrimCons(tmp7211, tmp7215)

tmp7217 := PrimCons(tmp7209, tmp7216)

tmp7218 := PrimCons(tmp7217, Nil)

tmp7219 := PrimCons(sym_1, tmp7218)

__e.TailApply(PrimFunc(symshen_4cons_1form_1respect_1modes), tmp7219)
return


} else {
tmp7221 := PrimEqual(Nil, Select6962)

if True == tmp7221 {
__e.TailApply(PrimFunc(symshen_4cons_1form_1respect_1modes), Select6961)
return
} else {
__e.TailApply(PrimFunc(symthaw), GoTo6959)
return
}


}


}, 1)

tmp7248 := PrimTail(Select6964)

__e.TailApply(tmp7208, tmp7248)
return


}, 1)

tmp7249 := PrimHead(Select6964)

__e.TailApply(tmp7207, tmp7249)
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo6959)
return
}


} else {
__e.TailApply(PrimFunc(symthaw), GoTo6959)
return
}


}, 1)

tmp7254 := MakeNative(func(__e *ControlFlow) {
tmp7282 := PrimIsPair(Select6964)

var ifres7262 Obj

if True == tmp7282 {
tmp7280 := PrimTail(Select6964)

tmp7281 := PrimIsPair(tmp7280)

var ifres7264 Obj

if True == tmp7281 {
tmp7277 := PrimTail(Select6964)

tmp7278 := PrimTail(tmp7277)

tmp7279 := PrimIsPair(tmp7278)

var ifres7266 Obj

if True == tmp7279 {
tmp7268 := PrimCons(symdo, Nil)

tmp7269 := PrimCons(sym_d, tmp7268)

tmp7270 := PrimCons(sym_7, tmp7269)

tmp7271 := PrimCons(symor, tmp7270)

tmp7272 := PrimCons(symand, tmp7271)

tmp7273 := PrimCons(symappend, tmp7272)

tmp7274 := PrimCons(sym_8v, tmp7273)

tmp7275 := PrimCons(sym_8p, tmp7274)

tmp7276 := Call(__e, PrimFunc(symelement_2), Select6963, tmp7275)


var ifres7267 Obj

if True == tmp7276 {
ifres7267 = True


} else {
ifres7267 = False


}

ifres7266 = ifres7267


} else {
ifres7266 = False


}

var ifres7265 Obj

if True == ifres7266 {
ifres7265 = True


} else {
ifres7265 = False


}

ifres7264 = ifres7265


} else {
ifres7264 = False


}

var ifres7263 Obj

if True == ifres7264 {
ifres7263 = True


} else {
ifres7263 = False


}

ifres7262 = ifres7263


} else {
ifres7262 = False


}

if True == ifres7262 {
tmp7255 := PrimHead(Select6964)

tmp7256 := PrimTail(Select6964)

tmp7257 := PrimCons(Select6963, tmp7256)

tmp7258 := Call(__e, PrimFunc(symshen_4process_1assoc), tmp7257)


tmp7259 := PrimCons(tmp7258, Nil)

tmp7260 := PrimCons(tmp7255, tmp7259)

__e.Return(PrimCons(Select6963, tmp7260))
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo6958)
return
}


}, 0)

__e.TailApply(tmp7206, tmp7254)
return


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


}, 1)

tmp7445 := PrimTail(V6957)

__e.TailApply(tmp7139, tmp7445)
return


}, 1)

tmp7446 := PrimHead(V6957)

__e.TailApply(tmp7138, tmp7446)
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo6958)
return
}


}, 1)

tmp7449 := MakeNative(func(__e *ControlFlow) {
__e.Return(V6957)
return
}, 0)

__e.TailApply(tmp7137, tmp7449)
return


}, 1)

tmp7450 := Call(__e, ns2_1set, symshen_4macros, tmp7136)


_ = tmp7450

tmp7451 := MakeNative(func(__e *ControlFlow) {
V6965 := __e.Get(1)
_ = V6965
tmp7452 := MakeNative(func(__e *ControlFlow) {
GoTo6966 := __e.Get(1)
_ = GoTo6966
tmp7486 := PrimIsPair(V6965)

if True == tmp7486 {
tmp7453 := MakeNative(func(__e *ControlFlow) {
Select6967 := __e.Get(1)
_ = Select6967
tmp7454 := MakeNative(func(__e *ControlFlow) {
Select6968 := __e.Get(1)
_ = Select6968
tmp7482 := PrimEqual(sym_7, Select6967)

var ifres7475 Obj

if True == tmp7482 {
tmp7481 := PrimIsPair(Select6968)

var ifres7477 Obj

if True == tmp7481 {
tmp7479 := PrimTail(Select6968)

tmp7480 := PrimEqual(Nil, tmp7479)

var ifres7478 Obj

if True == tmp7480 {
ifres7478 = True


} else {
ifres7478 = False


}

ifres7477 = ifres7478


} else {
ifres7477 = False


}

var ifres7476 Obj

if True == ifres7477 {
ifres7476 = True


} else {
ifres7476 = False


}

ifres7475 = ifres7476


} else {
ifres7475 = False


}

if True == ifres7475 {
tmp7455 := PrimHead(Select6968)

tmp7456 := Call(__e, PrimFunc(symshen_4cons_1form_1respect_1modes), tmp7455)


tmp7457 := PrimCons(tmp7456, Nil)

__e.Return(PrimCons(sym_7, tmp7457))
return


} else {
tmp7473 := PrimEqual(sym_1, Select6967)

var ifres7466 Obj

if True == tmp7473 {
tmp7472 := PrimIsPair(Select6968)

var ifres7468 Obj

if True == tmp7472 {
tmp7470 := PrimTail(Select6968)

tmp7471 := PrimEqual(Nil, tmp7470)

var ifres7469 Obj

if True == tmp7471 {
ifres7469 = True


} else {
ifres7469 = False


}

ifres7468 = ifres7469


} else {
ifres7468 = False


}

var ifres7467 Obj

if True == ifres7468 {
ifres7467 = True


} else {
ifres7467 = False


}

ifres7466 = ifres7467


} else {
ifres7466 = False


}

if True == ifres7466 {
tmp7458 := PrimHead(Select6968)

tmp7459 := Call(__e, PrimFunc(symshen_4cons_1form_1respect_1modes), tmp7458)


tmp7460 := PrimCons(tmp7459, Nil)

__e.Return(PrimCons(sym_1, tmp7460))
return


} else {
tmp7461 := Call(__e, PrimFunc(symshen_4cons_1form_1respect_1modes), Select6967)


tmp7462 := Call(__e, PrimFunc(symshen_4cons_1form_1respect_1modes), Select6968)


tmp7463 := PrimCons(tmp7462, Nil)

tmp7464 := PrimCons(tmp7461, tmp7463)

__e.Return(PrimCons(symcons, tmp7464))
return


}


}


}, 1)

tmp7483 := PrimTail(V6965)

__e.TailApply(tmp7454, tmp7483)
return


}, 1)

tmp7484 := PrimHead(V6965)

__e.TailApply(tmp7453, tmp7484)
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo6966)
return
}


}, 1)

tmp7487 := MakeNative(func(__e *ControlFlow) {
__e.Return(V6965)
return
}, 0)

__e.TailApply(tmp7452, tmp7487)
return


}, 1)

tmp7488 := Call(__e, ns2_1set, symshen_4cons_1form_1respect_1modes, tmp7451)


_ = tmp7488

tmp7489 := MakeNative(func(__e *ControlFlow) {
V6969 := __e.Get(1)
_ = V6969
V6970 := __e.Get(2)
_ = V6970
tmp7490 := MakeNative(func(__e *ControlFlow) {
W6971 := __e.Get(1)
_ = W6971
tmp7491 := MakeNative(func(__e *ControlFlow) {
W6972 := __e.Get(1)
_ = W6972
tmp7492 := MakeNative(func(__e *ControlFlow) {
W6973 := __e.Get(1)
_ = W6973
__e.Return(V6969)
return
}, 1)

tmp7493 := Call(__e, PrimFunc(symfn), V6969)


tmp7494 := Call(__e, PrimFunc(symshen_4record_1macro), V6969, tmp7493)


__e.TailApply(tmp7492, tmp7494)
return


}, 1)

tmp7495 := Call(__e, PrimFunc(symappend), V6970, W6971)


tmp7496 := PrimCons(V6969, tmp7495)

tmp7497 := PrimCons(symdefine, tmp7496)

tmp7498 := Call(__e, PrimFunc(symeval), tmp7497)


__e.TailApply(tmp7491, tmp7498)
return


}, 1)

tmp7499 := PrimCons(symX, Nil)

tmp7500 := PrimCons(sym_1_6, tmp7499)

tmp7501 := PrimCons(symX, tmp7500)

__e.TailApply(tmp7490, tmp7501)
return


}, 2)

tmp7502 := Call(__e, ns2_1set, symshen_4process_1def, tmp7489)


_ = tmp7502

tmp7503 := MakeNative(func(__e *ControlFlow) {
V6974 := __e.Get(1)
_ = V6974
tmp7543 := PrimIsPair(V6974)

var ifres7517 Obj

if True == tmp7543 {
tmp7541 := PrimHead(V6974)

tmp7542 := PrimEqual(symlet, tmp7541)

var ifres7519 Obj

if True == tmp7542 {
tmp7539 := PrimTail(V6974)

tmp7540 := PrimIsPair(tmp7539)

var ifres7521 Obj

if True == tmp7540 {
tmp7536 := PrimTail(V6974)

tmp7537 := PrimTail(tmp7536)

tmp7538 := PrimIsPair(tmp7537)

var ifres7523 Obj

if True == tmp7538 {
tmp7532 := PrimTail(V6974)

tmp7533 := PrimTail(tmp7532)

tmp7534 := PrimTail(tmp7533)

tmp7535 := PrimIsPair(tmp7534)

var ifres7525 Obj

if True == tmp7535 {
tmp7527 := PrimTail(V6974)

tmp7528 := PrimTail(tmp7527)

tmp7529 := PrimTail(tmp7528)

tmp7530 := PrimTail(tmp7529)

tmp7531 := PrimIsPair(tmp7530)

var ifres7526 Obj

if True == tmp7531 {
ifres7526 = True


} else {
ifres7526 = False


}

ifres7525 = ifres7526


} else {
ifres7525 = False


}

var ifres7524 Obj

if True == ifres7525 {
ifres7524 = True


} else {
ifres7524 = False


}

ifres7523 = ifres7524


} else {
ifres7523 = False


}

var ifres7522 Obj

if True == ifres7523 {
ifres7522 = True


} else {
ifres7522 = False


}

ifres7521 = ifres7522


} else {
ifres7521 = False


}

var ifres7520 Obj

if True == ifres7521 {
ifres7520 = True


} else {
ifres7520 = False


}

ifres7519 = ifres7520


} else {
ifres7519 = False


}

var ifres7518 Obj

if True == ifres7519 {
ifres7518 = True


} else {
ifres7518 = False


}

ifres7517 = ifres7518


} else {
ifres7517 = False


}

if True == ifres7517 {
tmp7504 := PrimTail(V6974)

tmp7505 := PrimHead(tmp7504)

tmp7506 := PrimTail(V6974)

tmp7507 := PrimTail(tmp7506)

tmp7508 := PrimHead(tmp7507)

tmp7509 := PrimTail(V6974)

tmp7510 := PrimTail(tmp7509)

tmp7511 := PrimTail(tmp7510)

tmp7512 := PrimCons(symlet, tmp7511)

tmp7513 := PrimCons(tmp7512, Nil)

tmp7514 := PrimCons(tmp7508, tmp7513)

tmp7515 := PrimCons(tmp7505, tmp7514)

__e.Return(PrimCons(symlet, tmp7515))
return


} else {
__e.Return(V6974)
return
}


}, 1)

tmp7544 := Call(__e, ns2_1set, symshen_4process_1let, tmp7503)


_ = tmp7544

tmp7545 := MakeNative(func(__e *ControlFlow) {
V6975 := __e.Get(1)
_ = V6975
tmp7546 := MakeNative(func(__e *ControlFlow) {
GoTo6977 := __e.Get(1)
_ = GoTo6977
tmp7581 := PrimIsPair(V6975)

if True == tmp7581 {
tmp7547 := MakeNative(func(__e *ControlFlow) {
Select6984 := __e.Get(1)
_ = Select6984
tmp7577 := PrimHead(V6975)

tmp7578 := PrimEqual(sym_8s, tmp7577)

if True == tmp7578 {
tmp7575 := PrimIsPair(Select6984)

if True == tmp7575 {
tmp7548 := MakeNative(func(__e *ControlFlow) {
Select6982 := __e.Get(1)
_ = Select6982
tmp7549 := MakeNative(func(__e *ControlFlow) {
Select6983 := __e.Get(1)
_ = Select6983
tmp7571 := PrimIsPair(Select6983)

if True == tmp7571 {
tmp7550 := MakeNative(func(__e *ControlFlow) {
Select6981 := __e.Get(1)
_ = Select6981
tmp7568 := PrimIsPair(Select6981)

if True == tmp7568 {
tmp7551 := PrimCons(sym_8s, Select6983)

tmp7552 := Call(__e, PrimFunc(symshen_4process_1_8s), tmp7551)


tmp7553 := PrimCons(tmp7552, Nil)

tmp7554 := PrimCons(Select6982, tmp7553)

__e.Return(PrimCons(sym_8s, tmp7554))
return


} else {
tmp7566 := PrimEqual(Nil, Select6981)

var ifres7563 Obj

if True == tmp7566 {
tmp7565 := PrimIsString(Select6982)

var ifres7564 Obj

if True == tmp7565 {
ifres7564 = True


} else {
ifres7564 = False


}

ifres7563 = ifres7564


} else {
ifres7563 = False


}

if True == ifres7563 {
tmp7555 := MakeNative(func(__e *ControlFlow) {
W6976 := __e.Get(1)
_ = W6976
tmp7559 := Call(__e, PrimFunc(symlength), W6976)


tmp7560 := PrimGreatThan(tmp7559, MakeNumber(1))

if True == tmp7560 {
tmp7556 := Call(__e, PrimFunc(symappend), W6976, Select6983)


tmp7557 := PrimCons(sym_8s, tmp7556)

__e.TailApply(PrimFunc(symshen_4process_1_8s), tmp7557)
return


} else {
__e.Return(V6975)
return
}


}, 1)

tmp7561 := Call(__e, PrimFunc(symexplode), Select6982)


__e.TailApply(tmp7555, tmp7561)
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo6977)
return
}


}


}, 1)

tmp7569 := PrimTail(Select6983)

__e.TailApply(tmp7550, tmp7569)
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo6977)
return
}


}, 1)

tmp7572 := PrimTail(Select6984)

__e.TailApply(tmp7549, tmp7572)
return


}, 1)

tmp7573 := PrimHead(Select6984)

__e.TailApply(tmp7548, tmp7573)
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo6977)
return
}


} else {
__e.TailApply(PrimFunc(symthaw), GoTo6977)
return
}


}, 1)

tmp7579 := PrimTail(V6975)

__e.TailApply(tmp7547, tmp7579)
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo6977)
return
}


}, 1)

tmp7582 := MakeNative(func(__e *ControlFlow) {
__e.Return(V6975)
return
}, 0)

__e.TailApply(tmp7546, tmp7582)
return


}, 1)

tmp7583 := Call(__e, ns2_1set, symshen_4process_1_8s, tmp7545)


_ = tmp7583

tmp7584 := MakeNative(func(__e *ControlFlow) {
V6985 := __e.Get(1)
_ = V6985
V6986 := __e.Get(2)
_ = V6986
tmp7585 := MakeNative(func(__e *ControlFlow) {
W6987 := __e.Get(1)
_ = W6987
tmp7586 := MakeNative(func(__e *ControlFlow) {
W6988 := __e.Get(1)
_ = W6988
__e.Return(W6987)
return
}, 1)

tmp7587 := MakeNative(func(__e *ControlFlow) {
Z6989 := __e.Get(1)
_ = Z6989
__e.TailApply(PrimFunc(symshen_4_5datatype_6), Z6989)
return
}, 1)

tmp7588 := PrimCons(W6987, V6986)

tmp7589 := Call(__e, PrimFunc(symcompile), tmp7587, tmp7588)


__e.TailApply(tmp7586, tmp7589)
return


}, 1)

tmp7590 := Call(__e, PrimFunc(symshen_4intern_1type), V6985)


__e.TailApply(tmp7585, tmp7590)
return


}, 2)

tmp7591 := Call(__e, ns2_1set, symshen_4process_1datatype, tmp7584)


_ = tmp7591

tmp7592 := MakeNative(func(__e *ControlFlow) {
V6990 := __e.Get(1)
_ = V6990
tmp7593 := PrimStr(V6990)

tmp7594 := PrimStringConcat(tmp7593, MakeString("#type"))

__e.Return(PrimIntern(tmp7594))
return


}, 1)

tmp7595 := Call(__e, ns2_1set, symshen_4intern_1type, tmp7592)


_ = tmp7595

tmp7596 := MakeNative(func(__e *ControlFlow) {
V6991 := __e.Get(1)
_ = V6991
tmp7597 := PrimValue(symshen_4_dsynonyms_d)

tmp7598 := Call(__e, PrimFunc(symappend), V6991, tmp7597)


tmp7599 := PrimSet(symshen_4_dsynonyms_d, tmp7598)

__e.TailApply(PrimFunc(symshen_4synonyms_1h), tmp7599)
return


}, 1)

tmp7600 := Call(__e, ns2_1set, symshen_4process_1synonyms, tmp7596)


_ = tmp7600

tmp7601 := MakeNative(func(__e *ControlFlow) {
V6994 := __e.Get(1)
_ = V6994
tmp7651 := PrimIsPair(V6994)

var ifres7612 Obj

if True == tmp7651 {
tmp7649 := PrimHead(V6994)

tmp7650 := PrimEqual(symdefun, tmp7649)

var ifres7614 Obj

if True == tmp7650 {
tmp7647 := PrimTail(V6994)

tmp7648 := PrimIsPair(tmp7647)

var ifres7616 Obj

if True == tmp7648 {
tmp7644 := PrimTail(V6994)

tmp7645 := PrimTail(tmp7644)

tmp7646 := PrimIsPair(tmp7645)

var ifres7618 Obj

if True == tmp7646 {
tmp7640 := PrimTail(V6994)

tmp7641 := PrimTail(tmp7640)

tmp7642 := PrimHead(tmp7641)

tmp7643 := PrimIsPair(tmp7642)

var ifres7620 Obj

if True == tmp7643 {
tmp7635 := PrimTail(V6994)

tmp7636 := PrimTail(tmp7635)

tmp7637 := PrimHead(tmp7636)

tmp7638 := PrimTail(tmp7637)

tmp7639 := PrimEqual(Nil, tmp7638)

var ifres7622 Obj

if True == tmp7639 {
tmp7631 := PrimTail(V6994)

tmp7632 := PrimTail(tmp7631)

tmp7633 := PrimTail(tmp7632)

tmp7634 := PrimIsPair(tmp7633)

var ifres7624 Obj

if True == tmp7634 {
tmp7626 := PrimTail(V6994)

tmp7627 := PrimTail(tmp7626)

tmp7628 := PrimTail(tmp7627)

tmp7629 := PrimTail(tmp7628)

tmp7630 := PrimEqual(Nil, tmp7629)

var ifres7625 Obj

if True == tmp7630 {
ifres7625 = True


} else {
ifres7625 = False


}

ifres7624 = ifres7625


} else {
ifres7624 = False


}

var ifres7623 Obj

if True == ifres7624 {
ifres7623 = True


} else {
ifres7623 = False


}

ifres7622 = ifres7623


} else {
ifres7622 = False


}

var ifres7621 Obj

if True == ifres7622 {
ifres7621 = True


} else {
ifres7621 = False


}

ifres7620 = ifres7621


} else {
ifres7620 = False


}

var ifres7619 Obj

if True == ifres7620 {
ifres7619 = True


} else {
ifres7619 = False


}

ifres7618 = ifres7619


} else {
ifres7618 = False


}

var ifres7617 Obj

if True == ifres7618 {
ifres7617 = True


} else {
ifres7617 = False


}

ifres7616 = ifres7617


} else {
ifres7616 = False


}

var ifres7615 Obj

if True == ifres7616 {
ifres7615 = True


} else {
ifres7615 = False


}

ifres7614 = ifres7615


} else {
ifres7614 = False


}

var ifres7613 Obj

if True == ifres7614 {
ifres7613 = True


} else {
ifres7613 = False


}

ifres7612 = ifres7613


} else {
ifres7612 = False


}

if True == ifres7612 {
tmp7602 := PrimTail(V6994)

tmp7603 := PrimTail(tmp7602)

tmp7604 := PrimHead(tmp7603)

tmp7605 := PrimHead(tmp7604)

tmp7606 := PrimTail(V6994)

tmp7607 := PrimTail(tmp7606)

tmp7608 := PrimTail(tmp7607)

tmp7609 := PrimCons(tmp7605, tmp7608)

tmp7610 := PrimCons(sym_c_4, tmp7609)

__e.TailApply(PrimFunc(symeval), tmp7610)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4lambda_1of_1defun)
return
}


}, 1)

tmp7652 := Call(__e, ns2_1set, symshen_4lambda_1of_1defun, tmp7601)


_ = tmp7652

tmp7653 := MakeNative(func(__e *ControlFlow) {
V6995 := __e.Get(1)
_ = V6995
tmp7654 := MakeNative(func(__e *ControlFlow) {
W6996 := __e.Get(1)
_ = W6996
tmp7655 := MakeNative(func(__e *ControlFlow) {
W6998 := __e.Get(1)
_ = W6998
tmp7656 := MakeNative(func(__e *ControlFlow) {
W6999 := __e.Get(1)
_ = W6999
__e.Return(symsynonyms)
return
}, 1)

tmp7657 := PrimSet(symshen_4_ddemodulation_1function_d, W6998)

__e.TailApply(tmp7656, tmp7657)
return


}, 1)

tmp7658 := Call(__e, PrimFunc(symshen_4compile_1synonyms), W6996)


tmp7659 := Call(__e, PrimFunc(symshen_4shendef_1_6kldef), symshen_4demod, tmp7658)


tmp7660 := Call(__e, PrimFunc(symshen_4lambda_1of_1defun), tmp7659)


__e.TailApply(tmp7655, tmp7660)
return


}, 1)

tmp7661 := MakeNative(func(__e *ControlFlow) {
Z6997 := __e.Get(1)
_ = Z6997
__e.TailApply(PrimFunc(symshen_4curry_1type), Z6997)
return
}, 1)

tmp7662 := Call(__e, PrimFunc(symmap), tmp7661, V6995)


__e.TailApply(tmp7654, tmp7662)
return


}, 1)

tmp7663 := Call(__e, ns2_1set, symshen_4synonyms_1h, tmp7653)


_ = tmp7663

tmp7664 := MakeNative(func(__e *ControlFlow) {
V7002 := __e.Get(1)
_ = V7002
tmp7686 := PrimEqual(Nil, V7002)

if True == tmp7686 {
tmp7665 := MakeNative(func(__e *ControlFlow) {
W7003 := __e.Get(1)
_ = W7003
tmp7666 := PrimCons(W7003, Nil)

tmp7667 := PrimCons(sym_1_6, tmp7666)

__e.Return(PrimCons(W7003, tmp7667))
return


}, 1)

tmp7668 := Call(__e, PrimFunc(symgensym), symX)


__e.TailApply(tmp7665, tmp7668)
return


} else {
tmp7684 := PrimIsPair(V7002)

var ifres7680 Obj

if True == tmp7684 {
tmp7682 := PrimTail(V7002)

tmp7683 := PrimIsPair(tmp7682)

var ifres7681 Obj

if True == tmp7683 {
ifres7681 = True


} else {
ifres7681 = False


}

ifres7680 = ifres7681


} else {
ifres7680 = False


}

if True == ifres7680 {
tmp7669 := PrimHead(V7002)

tmp7670 := Call(__e, PrimFunc(symshen_4rcons__form), tmp7669)


tmp7671 := PrimTail(V7002)

tmp7672 := PrimHead(tmp7671)

tmp7673 := Call(__e, PrimFunc(symshen_4rcons__form), tmp7672)


tmp7674 := PrimTail(V7002)

tmp7675 := PrimTail(tmp7674)

tmp7676 := Call(__e, PrimFunc(symshen_4compile_1synonyms), tmp7675)


tmp7677 := PrimCons(tmp7673, tmp7676)

tmp7678 := PrimCons(sym_1_6, tmp7677)

__e.Return(PrimCons(tmp7670, tmp7678))
return


} else {
__e.Return(PrimSimpleError(MakeString("synonyms requires an even number of arguments\n")))
return
}


}


}, 1)

tmp7687 := Call(__e, ns2_1set, symshen_4compile_1synonyms, tmp7664)


_ = tmp7687

tmp7688 := MakeNative(func(__e *ControlFlow) {
V7004 := __e.Get(1)
_ = V7004
tmp7689 := MakeNative(func(__e *ControlFlow) {
GoTo7005 := __e.Get(1)
_ = GoTo7005
tmp7717 := PrimIsPair(V7004)

if True == tmp7717 {
tmp7690 := MakeNative(func(__e *ControlFlow) {
Select7012 := __e.Get(1)
_ = Select7012
tmp7713 := PrimHead(V7004)

tmp7714 := PrimEqual(sym_c_4, tmp7713)

if True == tmp7714 {
tmp7711 := PrimIsPair(Select7012)

if True == tmp7711 {
tmp7691 := MakeNative(func(__e *ControlFlow) {
Select7010 := __e.Get(1)
_ = Select7010
tmp7692 := MakeNative(func(__e *ControlFlow) {
Select7011 := __e.Get(1)
_ = Select7011
tmp7707 := PrimIsPair(Select7011)

if True == tmp7707 {
tmp7693 := MakeNative(func(__e *ControlFlow) {
Select7009 := __e.Get(1)
_ = Select7009
tmp7704 := PrimIsPair(Select7009)

if True == tmp7704 {
tmp7694 := PrimCons(sym_c_4, Select7011)

tmp7695 := Call(__e, PrimFunc(symshen_4process_1lambda), tmp7694)


tmp7696 := PrimCons(tmp7695, Nil)

tmp7697 := PrimCons(Select7010, tmp7696)

__e.Return(PrimCons(symlambda, tmp7697))
return


} else {
tmp7702 := PrimEqual(Nil, Select7009)

if True == tmp7702 {
tmp7700 := PrimIsVariable(Select7010)

if True == tmp7700 {
__e.Return(PrimCons(symlambda, Select7012))
return
} else {
tmp7698 := Call(__e, PrimFunc(symshen_4app), Select7010, MakeString(" is not a variable\n"), symshen_4s)


__e.Return(PrimSimpleError(tmp7698))
return


}


} else {
__e.TailApply(PrimFunc(symthaw), GoTo7005)
return
}


}


}, 1)

tmp7705 := PrimTail(Select7011)

__e.TailApply(tmp7693, tmp7705)
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo7005)
return
}


}, 1)

tmp7708 := PrimTail(Select7012)

__e.TailApply(tmp7692, tmp7708)
return


}, 1)

tmp7709 := PrimHead(Select7012)

__e.TailApply(tmp7691, tmp7709)
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo7005)
return
}


} else {
__e.TailApply(PrimFunc(symthaw), GoTo7005)
return
}


}, 1)

tmp7715 := PrimTail(V7004)

__e.TailApply(tmp7690, tmp7715)
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo7005)
return
}


}, 1)

tmp7718 := MakeNative(func(__e *ControlFlow) {
__e.Return(V7004)
return
}, 0)

__e.TailApply(tmp7689, tmp7718)
return


}, 1)

tmp7719 := Call(__e, ns2_1set, symshen_4process_1lambda, tmp7688)


_ = tmp7719

tmp7720 := MakeNative(func(__e *ControlFlow) {
V7015 := __e.Get(1)
_ = V7015
tmp7721 := MakeNative(func(__e *ControlFlow) {
GoTo7016 := __e.Get(1)
_ = GoTo7016
tmp7761 := PrimIsPair(V7015)

if True == tmp7761 {
tmp7722 := MakeNative(func(__e *ControlFlow) {
Select7024 := __e.Get(1)
_ = Select7024
tmp7757 := PrimHead(V7015)

tmp7758 := PrimEqual(symcases, tmp7757)

if True == tmp7758 {
tmp7755 := PrimIsPair(Select7024)

if True == tmp7755 {
tmp7723 := MakeNative(func(__e *ControlFlow) {
Select7022 := __e.Get(1)
_ = Select7022
tmp7724 := MakeNative(func(__e *ControlFlow) {
Select7023 := __e.Get(1)
_ = Select7023
tmp7751 := PrimEqual(True, Select7022)

var ifres7748 Obj

if True == tmp7751 {
tmp7750 := PrimIsPair(Select7023)

var ifres7749 Obj

if True == tmp7750 {
ifres7749 = True


} else {
ifres7749 = False


}

ifres7748 = ifres7749


} else {
ifres7748 = False


}

if True == ifres7748 {
__e.Return(PrimHead(Select7023))
return
} else {
tmp7725 := MakeNative(func(__e *ControlFlow) {
GoTo7019 := __e.Get(1)
_ = GoTo7019
tmp7743 := PrimIsPair(Select7023)

if True == tmp7743 {
tmp7726 := MakeNative(func(__e *ControlFlow) {
Select7020 := __e.Get(1)
_ = Select7020
tmp7727 := MakeNative(func(__e *ControlFlow) {
Select7021 := __e.Get(1)
_ = Select7021
tmp7739 := PrimEqual(Nil, Select7021)

if True == tmp7739 {
tmp7728 := PrimCons(MakeString("error: cases exhausted"), Nil)

tmp7729 := PrimCons(symsimple_1error, tmp7728)

tmp7730 := PrimCons(tmp7729, Nil)

tmp7731 := PrimCons(Select7020, tmp7730)

tmp7732 := PrimCons(Select7022, tmp7731)

__e.Return(PrimCons(symif, tmp7732))
return


} else {
tmp7733 := PrimCons(symcases, Select7021)

tmp7734 := Call(__e, PrimFunc(symshen_4process_1cases), tmp7733)


tmp7735 := PrimCons(tmp7734, Nil)

tmp7736 := PrimCons(Select7020, tmp7735)

tmp7737 := PrimCons(Select7022, tmp7736)

__e.Return(PrimCons(symif, tmp7737))
return


}


}, 1)

tmp7740 := PrimTail(Select7023)

__e.TailApply(tmp7727, tmp7740)
return


}, 1)

tmp7741 := PrimHead(Select7023)

__e.TailApply(tmp7726, tmp7741)
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo7019)
return
}


}, 1)

tmp7744 := MakeNative(func(__e *ControlFlow) {
tmp7746 := PrimEqual(Nil, Select7023)

if True == tmp7746 {
__e.Return(PrimSimpleError(MakeString("error: odd number of case elements\n")))
return
} else {
__e.TailApply(PrimFunc(symthaw), GoTo7016)
return
}


}, 0)

__e.TailApply(tmp7725, tmp7744)
return


}


}, 1)

tmp7752 := PrimTail(Select7024)

__e.TailApply(tmp7724, tmp7752)
return


}, 1)

tmp7753 := PrimHead(Select7024)

__e.TailApply(tmp7723, tmp7753)
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo7016)
return
}


} else {
__e.TailApply(PrimFunc(symthaw), GoTo7016)
return
}


}, 1)

tmp7759 := PrimTail(V7015)

__e.TailApply(tmp7722, tmp7759)
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo7016)
return
}


}, 1)

tmp7762 := MakeNative(func(__e *ControlFlow) {
__e.Return(V7015)
return
}, 0)

__e.TailApply(tmp7721, tmp7762)
return


}, 1)

tmp7763 := Call(__e, ns2_1set, symshen_4process_1cases, tmp7720)


_ = tmp7763

tmp7764 := MakeNative(func(__e *ControlFlow) {
V7025 := __e.Get(1)
_ = V7025
tmp7765 := PrimCons(symrun, Nil)

tmp7766 := PrimCons(symget_1time, tmp7765)

tmp7767 := PrimCons(symrun, Nil)

tmp7768 := PrimCons(symget_1time, tmp7767)

tmp7769 := PrimCons(symStart, Nil)

tmp7770 := PrimCons(symFinish, tmp7769)

tmp7771 := PrimCons(sym_1, tmp7770)

tmp7772 := PrimCons(symTime, Nil)

tmp7773 := PrimCons(symstr, tmp7772)

tmp7774 := PrimCons(MakeString(" secs\n"), Nil)

tmp7775 := PrimCons(tmp7773, tmp7774)

tmp7776 := PrimCons(symcn, tmp7775)

tmp7777 := PrimCons(tmp7776, Nil)

tmp7778 := PrimCons(MakeString("\nrun time: "), tmp7777)

tmp7779 := PrimCons(symcn, tmp7778)

tmp7780 := PrimCons(symstoutput, Nil)

tmp7781 := PrimCons(tmp7780, Nil)

tmp7782 := PrimCons(tmp7779, tmp7781)

tmp7783 := PrimCons(sympr, tmp7782)

tmp7784 := PrimCons(symResult, Nil)

tmp7785 := PrimCons(tmp7783, tmp7784)

tmp7786 := PrimCons(symMessage, tmp7785)

tmp7787 := PrimCons(tmp7771, tmp7786)

tmp7788 := PrimCons(symTime, tmp7787)

tmp7789 := PrimCons(tmp7768, tmp7788)

tmp7790 := PrimCons(symFinish, tmp7789)

tmp7791 := PrimCons(V7025, tmp7790)

tmp7792 := PrimCons(symResult, tmp7791)

tmp7793 := PrimCons(tmp7766, tmp7792)

tmp7794 := PrimCons(symStart, tmp7793)

__e.Return(PrimCons(symlet, tmp7794))
return


}, 1)

tmp7795 := Call(__e, ns2_1set, symshen_4process_1time, tmp7764)


_ = tmp7795

tmp7796 := MakeNative(func(__e *ControlFlow) {
V7026 := __e.Get(1)
_ = V7026
tmp7822 := PrimIsPair(V7026)

var ifres7807 Obj

if True == tmp7822 {
tmp7820 := PrimTail(V7026)

tmp7821 := PrimIsPair(tmp7820)

var ifres7809 Obj

if True == tmp7821 {
tmp7817 := PrimTail(V7026)

tmp7818 := PrimTail(tmp7817)

tmp7819 := PrimIsPair(tmp7818)

var ifres7811 Obj

if True == tmp7819 {
tmp7813 := PrimTail(V7026)

tmp7814 := PrimTail(tmp7813)

tmp7815 := PrimTail(tmp7814)

tmp7816 := PrimIsPair(tmp7815)

var ifres7812 Obj

if True == tmp7816 {
ifres7812 = True


} else {
ifres7812 = False


}

ifres7811 = ifres7812


} else {
ifres7811 = False


}

var ifres7810 Obj

if True == ifres7811 {
ifres7810 = True


} else {
ifres7810 = False


}

ifres7809 = ifres7810


} else {
ifres7809 = False


}

var ifres7808 Obj

if True == ifres7809 {
ifres7808 = True


} else {
ifres7808 = False


}

ifres7807 = ifres7808


} else {
ifres7807 = False


}

if True == ifres7807 {
tmp7797 := PrimHead(V7026)

tmp7798 := PrimTail(V7026)

tmp7799 := PrimHead(tmp7798)

tmp7800 := PrimHead(V7026)

tmp7801 := PrimTail(V7026)

tmp7802 := PrimTail(tmp7801)

tmp7803 := PrimCons(tmp7800, tmp7802)

tmp7804 := PrimCons(tmp7803, Nil)

tmp7805 := PrimCons(tmp7799, tmp7804)

__e.Return(PrimCons(tmp7797, tmp7805))
return


} else {
__e.Return(V7026)
return
}


}, 1)

tmp7823 := Call(__e, ns2_1set, symshen_4process_1assoc, tmp7796)


_ = tmp7823

tmp7824 := MakeNative(func(__e *ControlFlow) {
V7027 := __e.Get(1)
_ = V7027
tmp7825 := PrimStr(V7027)

tmp7826 := Call(__e, PrimFunc(symshen_4mu_1h), tmp7825)


__e.Return(PrimIntern(tmp7826))
return


}, 1)

tmp7827 := Call(__e, ns2_1set, symshen_4make_1uppercase, tmp7824)


_ = tmp7827

tmp7828 := MakeNative(func(__e *ControlFlow) {
V7028 := __e.Get(1)
_ = V7028
tmp7847 := PrimEqual(MakeString(""), V7028)

if True == tmp7847 {
__e.Return(MakeString(""))
return
} else {
tmp7845 := Call(__e, PrimFunc(symshen_4_7string_2), V7028)


if True == tmp7845 {
tmp7829 := MakeNative(func(__e *ControlFlow) {
W7029 := __e.Get(1)
_ = W7029
tmp7830 := MakeNative(func(__e *ControlFlow) {
W7030 := __e.Get(1)
_ = W7030
tmp7831 := MakeNative(func(__e *ControlFlow) {
W7031 := __e.Get(1)
_ = W7031
tmp7832 := PrimTailString(V7028)

tmp7833 := Call(__e, PrimFunc(symshen_4mu_1h), tmp7832)


__e.TailApply(PrimFunc(sym_8s), W7031, tmp7833)
return


}, 1)

tmp7840 := PrimGreatEqual(W7029, MakeNumber(97))

var ifres7837 Obj

if True == tmp7840 {
tmp7839 := PrimLessEqual(W7029, MakeNumber(122))

var ifres7838 Obj

if True == tmp7839 {
ifres7838 = True


} else {
ifres7838 = False


}

ifres7837 = ifres7838


} else {
ifres7837 = False


}

var ifres7834 Obj

if True == ifres7837 {
tmp7835 := PrimNumberToString(W7030)

ifres7834 = tmp7835


} else {
tmp7836 := Call(__e, PrimFunc(symhdstr), V7028)


ifres7834 = tmp7836


}

__e.TailApply(tmp7831, ifres7834)
return


}, 1)

tmp7841 := PrimNumberSubtract(W7029, MakeNumber(32))

__e.TailApply(tmp7830, tmp7841)
return


}, 1)

tmp7842 := Call(__e, PrimFunc(symhdstr), V7028)


tmp7843 := PrimStringToNumber(tmp7842)

__e.TailApply(tmp7829, tmp7843)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4mu_1h)
return
}


}


}, 1)

tmp7848 := Call(__e, ns2_1set, symshen_4mu_1h, tmp7828)


_ = tmp7848

tmp7849 := MakeNative(func(__e *ControlFlow) {
V7032 := __e.Get(1)
_ = V7032
V7033 := __e.Get(2)
_ = V7033
tmp7850 := PrimValue(sym_dmacros_d)

tmp7851 := Call(__e, PrimFunc(symshen_4update_1assoc), V7032, V7033, tmp7850)


__e.Return(PrimSet(sym_dmacros_d, tmp7851))
return


}, 2)

tmp7852 := Call(__e, ns2_1set, symshen_4record_1macro, tmp7849)


_ = tmp7852

tmp7853 := MakeNative(func(__e *ControlFlow) {
V7043 := __e.Get(1)
_ = V7043
V7044 := __e.Get(2)
_ = V7044
V7045 := __e.Get(3)
_ = V7045
tmp7873 := PrimEqual(Nil, V7045)

if True == tmp7873 {
tmp7854 := PrimCons(V7043, V7044)

__e.Return(PrimCons(tmp7854, Nil))
return


} else {
tmp7855 := MakeNative(func(__e *ControlFlow) {
GoTo7046 := __e.Get(1)
_ = GoTo7046
tmp7870 := PrimIsPair(V7045)

if True == tmp7870 {
tmp7856 := MakeNative(func(__e *ControlFlow) {
Select7047 := __e.Get(1)
_ = Select7047
tmp7857 := MakeNative(func(__e *ControlFlow) {
Select7048 := __e.Get(1)
_ = Select7048
tmp7866 := PrimIsPair(Select7047)

var ifres7862 Obj

if True == tmp7866 {
tmp7864 := PrimHead(Select7047)

tmp7865 := PrimEqual(V7043, tmp7864)

var ifres7863 Obj

if True == tmp7865 {
ifres7863 = True


} else {
ifres7863 = False


}

ifres7862 = ifres7863


} else {
ifres7862 = False


}

if True == ifres7862 {
tmp7858 := PrimHead(Select7047)

tmp7859 := PrimCons(tmp7858, V7044)

__e.Return(PrimCons(tmp7859, Select7048))
return


} else {
tmp7860 := Call(__e, PrimFunc(symshen_4update_1assoc), V7043, V7044, Select7048)


__e.Return(PrimCons(Select7047, tmp7860))
return


}


}, 1)

tmp7867 := PrimTail(V7045)

__e.TailApply(tmp7857, tmp7867)
return


}, 1)

tmp7868 := PrimHead(V7045)

__e.TailApply(tmp7856, tmp7868)
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo7046)
return
}


}, 1)

tmp7871 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.update-assoc")))
return
}, 0)

__e.TailApply(tmp7855, tmp7871)
return


}


}, 3)

tmp7874 := Call(__e, ns2_1set, symshen_4update_1assoc, tmp7853)


_ = tmp7874

tmp7875 := MakeNative(func(__e *ControlFlow) {
tmp7883 := Call(__e, PrimFunc(symstinput))


tmp7884 := Call(__e, PrimFunc(symshen_4char_1stinput_2), tmp7883)


if True == tmp7884 {
tmp7876 := PrimCons(symstinput, Nil)

tmp7877 := PrimCons(tmp7876, Nil)

tmp7878 := PrimCons(symshen_4read_1unit_1string, tmp7877)

tmp7879 := PrimCons(tmp7878, Nil)

__e.Return(PrimCons(symstring_1_6n, tmp7879))
return


} else {
tmp7880 := PrimCons(symstinput, Nil)

tmp7881 := PrimCons(tmp7880, Nil)

__e.Return(PrimCons(symread_1byte, tmp7881))
return


}


}, 0)

tmp7885 := Call(__e, ns2_1set, symshen_4process_1read_1byte, tmp7875)


_ = tmp7885

tmp7886 := MakeNative(func(__e *ControlFlow) {
V7049 := __e.Get(1)
_ = V7049
tmp7887 := MakeNative(func(__e *ControlFlow) {
W7050 := __e.Get(1)
_ = W7050
tmp7888 := MakeNative(func(__e *ControlFlow) {
W7051 := __e.Get(1)
_ = W7051
tmp7889 := MakeNative(func(__e *ControlFlow) {
W7052 := __e.Get(1)
_ = W7052
tmp7890 := MakeNative(func(__e *ControlFlow) {
W7053 := __e.Get(1)
_ = W7053
tmp7891 := MakeNative(func(__e *ControlFlow) {
W7054 := __e.Get(1)
_ = W7054
tmp7892 := MakeNative(func(__e *ControlFlow) {
W7056 := __e.Get(1)
_ = W7056
tmp7893 := MakeNative(func(__e *ControlFlow) {
W7057 := __e.Get(1)
_ = W7057
tmp7894 := MakeNative(func(__e *ControlFlow) {
W7058 := __e.Get(1)
_ = W7058
tmp7895 := MakeNative(func(__e *ControlFlow) {
W7059 := __e.Get(1)
_ = W7059
tmp7896 := MakeNative(func(__e *ControlFlow) {
W7060 := __e.Get(1)
_ = W7060
tmp7897 := MakeNative(func(__e *ControlFlow) {
W7061 := __e.Get(1)
_ = W7061
tmp7898 := PrimCons(W7053, Nil)

tmp7899 := PrimCons(W7052, tmp7898)

tmp7900 := PrimCons(W7051, tmp7899)

tmp7901 := PrimCons(W7050, tmp7900)

__e.Return(PrimCons(W7061, tmp7901))
return


}, 1)

tmp7902 := Call(__e, PrimFunc(symshen_4continue), W7056, W7054, W7057, W7058, W7059, W7060)


tmp7903 := PrimCons(tmp7902, Nil)

tmp7904 := PrimCons(W7060, tmp7903)

tmp7905 := PrimCons(symlambda, tmp7904)

tmp7906 := PrimCons(tmp7905, Nil)

tmp7907 := PrimCons(W7059, tmp7906)

tmp7908 := PrimCons(symlambda, tmp7907)

tmp7909 := PrimCons(tmp7908, Nil)

tmp7910 := PrimCons(W7058, tmp7909)

tmp7911 := PrimCons(symlambda, tmp7910)

tmp7912 := PrimCons(tmp7911, Nil)

tmp7913 := PrimCons(W7057, tmp7912)

tmp7914 := PrimCons(symlambda, tmp7913)

__e.TailApply(tmp7897, tmp7914)
return


}, 1)

tmp7915 := Call(__e, PrimFunc(symgensym), symC)


__e.TailApply(tmp7896, tmp7915)
return


}, 1)

tmp7916 := Call(__e, PrimFunc(symgensym), symK)


__e.TailApply(tmp7895, tmp7916)
return


}, 1)

tmp7917 := Call(__e, PrimFunc(symgensym), symL)


__e.TailApply(tmp7894, tmp7917)
return


}, 1)

tmp7918 := Call(__e, PrimFunc(symgensym), symV)


__e.TailApply(tmp7893, tmp7918)
return


}, 1)

tmp7919 := Call(__e, PrimFunc(symshen_4received), V7049)


__e.TailApply(tmp7892, tmp7919)
return


}, 1)

tmp7920 := MakeNative(func(__e *ControlFlow) {
Z7055 := __e.Get(1)
_ = Z7055
__e.TailApply(PrimFunc(symshen_4_5body_6), Z7055)
return
}, 1)

tmp7921 := Call(__e, PrimFunc(symcompile), tmp7920, V7049)


__e.TailApply(tmp7891, tmp7921)
return


}, 1)

tmp7922 := PrimCons(True, Nil)

tmp7923 := PrimCons(symfreeze, tmp7922)

__e.TailApply(tmp7890, tmp7923)
return


}, 1)

__e.TailApply(tmp7889, MakeNumber(0))
return


}, 1)

tmp7924 := PrimCons(MakeNumber(0), Nil)

tmp7925 := PrimCons(symvector, tmp7924)

tmp7926 := PrimCons(tmp7925, Nil)

tmp7927 := PrimCons(MakeNumber(0), tmp7926)

tmp7928 := PrimCons(True, tmp7927)

tmp7929 := PrimCons(sym_8v, tmp7928)

__e.TailApply(tmp7888, tmp7929)
return


}, 1)

tmp7930 := PrimCons(symshen_4prolog_1vector, Nil)

__e.TailApply(tmp7887, tmp7930)
return


}, 1)

tmp7931 := Call(__e, ns2_1set, symshen_4call_1prolog, tmp7886)


_ = tmp7931

tmp7932 := MakeNative(func(__e *ControlFlow) {
V7064 := __e.Get(1)
_ = V7064
tmp7933 := MakeNative(func(__e *ControlFlow) {
GoTo7065 := __e.Get(1)
_ = GoTo7065
tmp7950 := PrimIsPair(V7064)

if True == tmp7950 {
tmp7934 := MakeNative(func(__e *ControlFlow) {
Select7066 := __e.Get(1)
_ = Select7066
tmp7935 := MakeNative(func(__e *ControlFlow) {
Select7067 := __e.Get(1)
_ = Select7067
tmp7946 := PrimEqual(symreceive, Select7066)

var ifres7939 Obj

if True == tmp7946 {
tmp7945 := PrimIsPair(Select7067)

var ifres7941 Obj

if True == tmp7945 {
tmp7943 := PrimTail(Select7067)

tmp7944 := PrimEqual(Nil, tmp7943)

var ifres7942 Obj

if True == tmp7944 {
ifres7942 = True


} else {
ifres7942 = False


}

ifres7941 = ifres7942


} else {
ifres7941 = False


}

var ifres7940 Obj

if True == ifres7941 {
ifres7940 = True


} else {
ifres7940 = False


}

ifres7939 = ifres7940


} else {
ifres7939 = False


}

if True == ifres7939 {
__e.Return(Select7067)
return
} else {
tmp7936 := Call(__e, PrimFunc(symshen_4received), Select7066)


tmp7937 := Call(__e, PrimFunc(symshen_4received), Select7067)


__e.TailApply(PrimFunc(symunion), tmp7936, tmp7937)
return


}


}, 1)

tmp7947 := PrimTail(V7064)

__e.TailApply(tmp7935, tmp7947)
return


}, 1)

tmp7948 := PrimHead(V7064)

__e.TailApply(tmp7934, tmp7948)
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo7065)
return
}


}, 1)

tmp7951 := MakeNative(func(__e *ControlFlow) {
__e.Return(Nil)
return
}, 0)

__e.TailApply(tmp7933, tmp7951)
return


}, 1)

tmp7952 := Call(__e, ns2_1set, symshen_4received, tmp7932)


_ = tmp7952

tmp7953 := MakeNative(func(__e *ControlFlow) {
tmp7954 := MakeNative(func(__e *ControlFlow) {
W7068 := __e.Get(1)
_ = W7068
tmp7955 := MakeNative(func(__e *ControlFlow) {
W7069 := __e.Get(1)
_ = W7069
tmp7956 := MakeNative(func(__e *ControlFlow) {
W7070 := __e.Get(1)
_ = W7070
__e.Return(W7070)
return
}, 1)

tmp7957 := PrimVectorSet(W7068, MakeNumber(1), MakeNumber(2))

__e.TailApply(tmp7956, tmp7957)
return


}, 1)

tmp7958 := PrimVectorSet(W7068, MakeNumber(0), symshen_4print_1prolog_1vector)

__e.TailApply(tmp7955, tmp7958)
return


}, 1)

tmp7959 := PrimValue(symshen_4_dprolog_1memory_d)

tmp7960 := PrimAbsvector(tmp7959)

__e.TailApply(tmp7954, tmp7960)
return


}, 0)

tmp7961 := Call(__e, ns2_1set, symshen_4prolog_1vector, tmp7953)


_ = tmp7961

tmp7962 := MakeNative(func(__e *ControlFlow) {
V7071 := __e.Get(1)
_ = V7071
__e.Return(V7071)
return
}, 1)

tmp7963 := Call(__e, ns2_1set, symreceive, tmp7962)


_ = tmp7963

tmp7964 := MakeNative(func(__e *ControlFlow) {
V7072 := __e.Get(1)
_ = V7072
tmp7972 := PrimIsPair(V7072)

if True == tmp7972 {
tmp7965 := PrimHead(V7072)

tmp7966 := Call(__e, PrimFunc(symshen_4rcons__form), tmp7965)


tmp7967 := PrimTail(V7072)

tmp7968 := Call(__e, PrimFunc(symshen_4rcons__form), tmp7967)


tmp7969 := PrimCons(tmp7968, Nil)

tmp7970 := PrimCons(tmp7966, tmp7969)

__e.Return(PrimCons(symcons, tmp7970))
return


} else {
__e.Return(V7072)
return
}


}, 1)

tmp7973 := Call(__e, ns2_1set, symshen_4rcons__form, tmp7964)


_ = tmp7973

tmp7974 := MakeNative(func(__e *ControlFlow) {
V7073 := __e.Get(1)
_ = V7073
tmp7981 := PrimIsPair(V7073)

if True == tmp7981 {
tmp7975 := PrimHead(V7073)

tmp7976 := PrimTail(V7073)

tmp7977 := Call(__e, PrimFunc(symshen_4tuple_1up), tmp7976)


tmp7978 := PrimCons(tmp7977, Nil)

tmp7979 := PrimCons(tmp7975, tmp7978)

__e.Return(PrimCons(sym_8p, tmp7979))
return


} else {
__e.Return(V7073)
return
}


}, 1)

tmp7982 := Call(__e, ns2_1set, symshen_4tuple_1up, tmp7974)


_ = tmp7982

tmp7983 := MakeNative(func(__e *ControlFlow) {
V7074 := __e.Get(1)
_ = V7074
tmp7984 := PrimValue(sym_dmacros_d)

tmp7985 := Call(__e, PrimFunc(symassoc), V7074, tmp7984)


tmp7986 := PrimValue(sym_dmacros_d)

tmp7987 := Call(__e, PrimFunc(symremove), tmp7985, tmp7986)


tmp7988 := PrimSet(sym_dmacros_d, tmp7987)

_ = tmp7988

__e.Return(V7074)
return


}, 1)

__e.TailApply(ns2_1set, symundefmacro, tmp7983)
return




}, 0)

