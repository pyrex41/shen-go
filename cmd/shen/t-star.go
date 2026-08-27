package main

import . "github.com/pyrex41/shen-go/kl"

var TStarMain = MakeNative(func(__e *ControlFlow) {
tmp14176 := MakeNative(func(__e *ControlFlow) {
V4448 := __e.Get(1)
_ = V4448
V4449 := __e.Get(2)
_ = V4449
tmp14177 := MakeNative(func(__e *ControlFlow) {
W4450 := __e.Get(1)
_ = W4450
tmp14178 := MakeNative(func(__e *ControlFlow) {
W4451 := __e.Get(1)
_ = W4451
tmp14179 := MakeNative(func(__e *ControlFlow) {
W4452 := __e.Get(1)
_ = W4452
tmp14180 := MakeNative(func(__e *ControlFlow) {
Z4453 := __e.Get(1)
_ = Z4453
__e.Return(MakeNative(func(__e *ControlFlow) {
Z4454 := __e.Get(1)
_ = Z4454
__e.Return(MakeNative(func(__e *ControlFlow) {
Z4455 := __e.Get(1)
_ = Z4455
__e.Return(MakeNative(func(__e *ControlFlow) {
Z4456 := __e.Get(1)
_ = Z4456
tmp14181 := MakeNative(func(__e *ControlFlow) {
W4457 := __e.Get(1)
_ = W4457
tmp14182 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14182

tmp14183 := Call(__e, PrimFunc(symshen_4deref), W4450, Z4453)


tmp14184 := Call(__e, PrimFunc(symreceive), tmp14183)


tmp14185 := Call(__e, PrimFunc(symshen_4deref), W4451, Z4453)


tmp14186 := Call(__e, PrimFunc(symreceive), tmp14185)


tmp14187 := MakeNative(func(__e *ControlFlow) {
tmp14188 := Call(__e, PrimFunc(symshen_4deref), W4452, Z4453)


tmp14189 := Call(__e, PrimFunc(symreceive), tmp14188)


tmp14190 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symreturn), W4457, Z4453, Z4454, Z4455, Z4456)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4toplevel_1forms), tmp14189, W4457, Z4453, Z4454, Z4455, tmp14190)
return


}, 0)

tmp14191 := Call(__e, PrimFunc(symshen_4insert_1prolog_1variables), tmp14184, tmp14186, W4457, Z4453, Z4454, Z4455, tmp14187)


__e.TailApply(PrimFunc(symshen_4gc), Z4453, tmp14191)
return


}, 1)

tmp14192 := Call(__e, PrimFunc(symshen_4newpv), Z4453)


__e.TailApply(tmp14181, tmp14192)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp14193 := Call(__e, PrimFunc(symshen_4prolog_1vector))


tmp14194 := Call(__e, tmp14180, tmp14193)


tmp14195 := Call(__e, PrimFunc(symvector), MakeNumber(0))


tmp14196 := Call(__e, PrimFunc(sym_8v), MakeNumber(0), tmp14195)


tmp14197 := Call(__e, PrimFunc(sym_8v), True, tmp14196)


tmp14198 := Call(__e, tmp14194, tmp14197)


tmp14199 := Call(__e, tmp14198, MakeNumber(0))


tmp14200 := MakeNative(func(__e *ControlFlow) {
__e.Return(True)
return
}, 0)

__e.TailApply(tmp14199, tmp14200)
return


}, 1)

tmp14201 := Call(__e, PrimFunc(symshen_4curry), V4448)


__e.TailApply(tmp14179, tmp14201)
return


}, 1)

tmp14202 := Call(__e, PrimFunc(symshen_4rectify_1type), V4449)


__e.TailApply(tmp14178, tmp14202)
return


}, 1)

tmp14203 := Call(__e, PrimFunc(symshen_4extract_1vars), V4449)


__e.TailApply(tmp14177, tmp14203)
return


}, 2)

tmp14204 := Call(__e, ns2_1set, symshen_4typecheck, tmp14176)


_ = tmp14204

tmp14205 := MakeNative(func(__e *ControlFlow) {
V4458 := __e.Get(1)
_ = V4458
V4459 := __e.Get(2)
_ = V4459
V4460 := __e.Get(3)
_ = V4460
V4461 := __e.Get(4)
_ = V4461
V4462 := __e.Get(5)
_ = V4462
V4463 := __e.Get(6)
_ = V4463
V4464 := __e.Get(7)
_ = V4464
tmp14206 := MakeNative(func(__e *ControlFlow) {
W4465 := __e.Get(1)
_ = W4465
tmp14224 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4465, False)
}
__typedArg0 := W4465
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp14224 {
tmp14222 := Call(__e, PrimFunc(symshen_4unlocked_2), V4462)


if True == tmp14222 {
tmp14207 := MakeNative(func(__e *ControlFlow) {
W4467 := __e.Get(1)
_ = W4467
tmp14219 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4467)
}
__typedArg0 := W4467
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp14219 {
tmp14208 := MakeNative(func(__e *ControlFlow) {
W4468 := __e.Get(1)
_ = W4468
tmp14209 := MakeNative(func(__e *ControlFlow) {
W4469 := __e.Get(1)
_ = W4469
tmp14210 := MakeNative(func(__e *ControlFlow) {
W4470 := __e.Get(1)
_ = W4470
tmp14211 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14211

tmp14212 := Call(__e, PrimFunc(symshen_4deref), W4470, V4461)


tmp14213 := Call(__e, PrimFunc(symsubst), tmp14212, W4468, V4459)


tmp14214 := Call(__e, PrimFunc(symshen_4insert_1prolog_1variables), W4469, tmp14213, V4460, V4461, V4462, V4463, V4464)


__e.TailApply(PrimFunc(symshen_4gc), V4461, tmp14214)
return


}, 1)

tmp14215 := Call(__e, PrimFunc(symshen_4newpv), V4461)


__e.TailApply(tmp14210, tmp14215)
return


}, 1)

tmp14216 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4467)
}
__typedArg0 := W4467
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(tmp14209, tmp14216)
return


}, 1)

tmp14217 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4467)
}
__typedArg0 := W4467
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp14208, tmp14217)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14220 := Call(__e, PrimFunc(symshen_4lazyderef), V4458, V4461)


__e.TailApply(tmp14207, tmp14220)
return


} else {
__e.Return(False)
return
}


} else {
__e.Return(W4465)
return
}


}, 1)

tmp14232 := Call(__e, PrimFunc(symshen_4unlocked_2), V4462)


var ifres14225 Obj

if True == tmp14232 {
tmp14226 := MakeNative(func(__e *ControlFlow) {
W4466 := __e.Get(1)
_ = W4466
tmp14229 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4466, Nil)
}
__typedArg0 := W4466
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp14229 {
tmp14227 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14227

__e.TailApply(PrimFunc(symis_b), V4459, V4460, V4461, V4462, V4463, V4464)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14230 := Call(__e, PrimFunc(symshen_4lazyderef), V4458, V4461)


tmp14231 := Call(__e, tmp14226, tmp14230)


ifres14225 = tmp14231


} else {
ifres14225 = False


}

__e.TailApply(tmp14206, ifres14225)
return


}, 7)

tmp14233 := Call(__e, ns2_1set, symshen_4insert_1prolog_1variables, tmp14205)


_ = tmp14233

tmp14234 := MakeNative(func(__e *ControlFlow) {
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
tmp14235 := MakeNative(func(__e *ControlFlow) {
W4477 := __e.Get(1)
_ = W4477
tmp14236 := MakeNative(func(__e *ControlFlow) {
W4478 := __e.Get(1)
_ = W4478
tmp14249 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4478, False)
}
__typedArg0 := W4478
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp14249 {
tmp14237 := MakeNative(func(__e *ControlFlow) {
W4484 := __e.Get(1)
_ = W4484
tmp14239 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4484, False)
}
__typedArg0 := W4484
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp14239 {
__e.TailApply(PrimFunc(symshen_4unlock), V4474, W4477)
return
} else {
__e.Return(W4484)
return
}


}, 1)

tmp14247 := Call(__e, PrimFunc(symshen_4unlocked_2), V4474)


var ifres14240 Obj

if True == tmp14247 {
tmp14241 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14241

tmp14242 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp14243 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V4472, Nil)
}
__typedArg0 := V4472
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14244 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14242, tmp14243)
}
__typedArg0 := tmp14242
__typedArg1 := tmp14243
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14245 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V4471, tmp14244)
}
__typedArg0 := V4471
__typedArg1 := tmp14244
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14246 := Call(__e, PrimFunc(symshen_4system_1S), tmp14245, Nil, V4473, V4474, W4477, V4476)


ifres14240 = tmp14246


} else {
ifres14240 = False


}

__e.TailApply(tmp14237, ifres14240)
return


} else {
__e.Return(W4478)
return
}


}, 1)

tmp14278 := Call(__e, PrimFunc(symshen_4unlocked_2), V4474)


var ifres14250 Obj

if True == tmp14278 {
tmp14251 := MakeNative(func(__e *ControlFlow) {
W4479 := __e.Get(1)
_ = W4479
tmp14275 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4479)
}
__typedArg0 := W4479
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp14275 {
tmp14252 := MakeNative(func(__e *ControlFlow) {
W4480 := __e.Get(1)
_ = W4480
tmp14271 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4480, symdefine)
}
__typedArg0 := W4480
__typedArg1 := symdefine
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp14271 {
tmp14253 := MakeNative(func(__e *ControlFlow) {
W4481 := __e.Get(1)
_ = W4481
tmp14267 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4481)
}
__typedArg0 := W4481
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp14267 {
tmp14254 := MakeNative(func(__e *ControlFlow) {
W4482 := __e.Get(1)
_ = W4482
tmp14255 := MakeNative(func(__e *ControlFlow) {
W4483 := __e.Get(1)
_ = W4483
tmp14256 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14256

tmp14257 := Call(__e, PrimFunc(symshen_4type_1theory_1enabled_2))


tmp14258 := MakeNative(func(__e *ControlFlow) {
tmp14259 := MakeNative(func(__e *ControlFlow) {
tmp14260 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dspy_d)
}
__typedArg0 := symshen_4_dspy_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp14261 := MakeNative(func(__e *ControlFlow) {
tmp14262 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W4482, W4483)
}
__typedArg0 := W4482
__typedArg1 := W4483
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14263 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefine, tmp14262)
}
__typedArg0 := symdefine
__typedArg1 := tmp14262
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4t_d), tmp14263, V4472, V4473, V4474, W4477, V4476)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4signal_1def), tmp14260, W4482, V4473, V4474, W4477, tmp14261)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V4473, V4474, W4477, tmp14259)
return


}, 0)

__e.TailApply(PrimFunc(symwhen), tmp14257, V4473, V4474, W4477, tmp14258)
return


}, 1)

tmp14264 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4481)
}
__typedArg0 := W4481
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(tmp14255, tmp14264)
return


}, 1)

tmp14265 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4481)
}
__typedArg0 := W4481
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp14254, tmp14265)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14268 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4479)
}
__typedArg0 := W4479
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14269 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14268, V4473)


__e.TailApply(tmp14253, tmp14269)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14272 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4479)
}
__typedArg0 := W4479
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14273 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14272, V4473)


__e.TailApply(tmp14252, tmp14273)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14276 := Call(__e, PrimFunc(symshen_4lazyderef), V4471, V4473)


tmp14277 := Call(__e, tmp14251, tmp14276)


ifres14250 = tmp14277


} else {
ifres14250 = False


}

__e.TailApply(tmp14236, ifres14250)
return


}, 1)

__e.TailApply(tmp14235, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V4475)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V4475
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())
return


}, 6)

tmp14280 := Call(__e, ns2_1set, symshen_4toplevel_1forms, tmp14234)


_ = tmp14280

tmp14281 := MakeNative(func(__e *ControlFlow) {
V4485 := __e.Get(1)
_ = V4485
V4486 := __e.Get(2)
_ = V4486
V4487 := __e.Get(3)
_ = V4487
V4488 := __e.Get(4)
_ = V4488
V4489 := __e.Get(5)
_ = V4489
V4490 := __e.Get(6)
_ = V4490
tmp14282 := MakeNative(func(__e *ControlFlow) {
W4491 := __e.Get(1)
_ = W4491
tmp14299 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4491, False)
}
__typedArg0 := W4491
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp14299 {
tmp14297 := Call(__e, PrimFunc(symshen_4unlocked_2), V4488)


if True == tmp14297 {
tmp14283 := MakeNative(func(__e *ControlFlow) {
W4493 := __e.Get(1)
_ = W4493
tmp14294 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4493, True)
}
__typedArg0 := W4493
__typedArg1 := True
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp14294 {
tmp14284 := MakeNative(func(__e *ControlFlow) {
W4494 := __e.Get(1)
_ = W4494
tmp14285 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14285

tmp14286 := Call(__e, PrimFunc(symshen_4deref), V4486, V4487)


tmp14287 := Call(__e, PrimFunc(symshen_4app), tmp14286, MakeString(")\n"), symshen_4a)


tmp14288 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("\ntypechecking (fn "))
__typedS1, __typedOK1 := TypedString(tmp14287)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("\ntypechecking (fn ")
__typedArg1 := tmp14287
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp14289 := Call(__e, PrimFunc(symstoutput))


tmp14290 := Call(__e, PrimFunc(sympr), tmp14288, tmp14289)


tmp14291 := Call(__e, PrimFunc(symis), W4494, tmp14290, V4487, V4488, V4489, V4490)


__e.TailApply(PrimFunc(symshen_4gc), V4487, tmp14291)
return


}, 1)

tmp14292 := Call(__e, PrimFunc(symshen_4newpv), V4487)


__e.TailApply(tmp14284, tmp14292)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14295 := Call(__e, PrimFunc(symshen_4lazyderef), V4485, V4487)


__e.TailApply(tmp14283, tmp14295)
return


} else {
__e.Return(False)
return
}


} else {
__e.Return(W4491)
return
}


}, 1)

tmp14307 := Call(__e, PrimFunc(symshen_4unlocked_2), V4488)


var ifres14300 Obj

if True == tmp14307 {
tmp14301 := MakeNative(func(__e *ControlFlow) {
W4492 := __e.Get(1)
_ = W4492
tmp14304 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4492, False)
}
__typedArg0 := W4492
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp14304 {
tmp14302 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14302

__e.TailApply(PrimFunc(symthaw), V4490)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14305 := Call(__e, PrimFunc(symshen_4lazyderef), V4485, V4487)


tmp14306 := Call(__e, tmp14301, tmp14305)


ifres14300 = tmp14306


} else {
ifres14300 = False


}

__e.TailApply(tmp14282, ifres14300)
return


}, 6)

tmp14308 := Call(__e, ns2_1set, symshen_4signal_1def, tmp14281)


_ = tmp14308

tmp14309 := MakeNative(func(__e *ControlFlow) {
V4495 := __e.Get(1)
_ = V4495
tmp14310 := Call(__e, PrimFunc(symshen_4curry_1type), V4495)


__e.TailApply(PrimFunc(symshen_4demodulate), tmp14310)
return


}, 1)

tmp14311 := Call(__e, ns2_1set, symshen_4rectify_1type, tmp14309)


_ = tmp14311

tmp14312 := MakeNative(func(__e *ControlFlow) {
V4496 := __e.Get(1)
_ = V4496
tmp14313 := MakeNative(func(__e *ControlFlow) {
tmp14314 := MakeNative(func(__e *ControlFlow) {
W4497 := __e.Get(1)
_ = W4497
tmp14316 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4497, V4496)
}
__typedArg0 := W4497
__typedArg1 := V4496
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp14316 {
__e.Return(V4496)
return
} else {
__e.TailApply(PrimFunc(symshen_4demodulate), W4497)
return
}


}, 1)

tmp14317 := MakeNative(func(__e *ControlFlow) {
Z4498 := __e.Get(1)
_ = Z4498
__e.TailApply(PrimFunc(symshen_4demod), Z4498)
return
}, 1)

tmp14318 := Call(__e, PrimFunc(symshen_4walk), tmp14317, V4496)


__e.TailApply(tmp14314, tmp14318)
return


}, 0)

tmp14319 := MakeNative(func(__e *ControlFlow) {
Z4499 := __e.Get(1)
_ = Z4499
__e.Return(V4496)
return
}, 1)

__e.TailApply(try_1catch, tmp14313, tmp14319)
return


}, 1)

tmp14320 := Call(__e, ns2_1set, symshen_4demodulate, tmp14312)


_ = tmp14320

tmp14321 := MakeNative(func(__e *ControlFlow) {
V4500 := __e.Get(1)
_ = V4500
tmp14445 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4500)
}
__typedArg0 := V4500
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14418 Obj

if True == tmp14445 {
tmp14443 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4500)
}
__typedArg0 := V4500
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14444 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14443)
}
__typedArg0 := tmp14443
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14420 Obj

if True == tmp14444 {
tmp14440 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4500)
}
__typedArg0 := V4500
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14441 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp14440)
}
__typedArg0 := tmp14440
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14442 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_1_1_6, tmp14441)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp14441
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14422 Obj

if True == tmp14442 {
tmp14437 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4500)
}
__typedArg0 := V4500
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14438 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14437)
}
__typedArg0 := tmp14437
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14439 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14438)
}
__typedArg0 := tmp14438
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14424 Obj

if True == tmp14439 {
tmp14433 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4500)
}
__typedArg0 := V4500
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14434 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14433)
}
__typedArg0 := tmp14433
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14435 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14434)
}
__typedArg0 := tmp14434
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14436 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14435)
}
__typedArg0 := tmp14435
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14426 Obj

if True == tmp14436 {
tmp14428 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4500)
}
__typedArg0 := V4500
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14429 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14428)
}
__typedArg0 := tmp14428
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14430 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14429)
}
__typedArg0 := tmp14429
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14431 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp14430)
}
__typedArg0 := tmp14430
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14432 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_1_1_6, tmp14431)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp14431
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14427 Obj

if True == tmp14432 {
ifres14427 = True


} else {
ifres14427 = False


}

ifres14426 = ifres14427


} else {
ifres14426 = False


}

var ifres14425 Obj

if True == ifres14426 {
ifres14425 = True


} else {
ifres14425 = False


}

ifres14424 = ifres14425


} else {
ifres14424 = False


}

var ifres14423 Obj

if True == ifres14424 {
ifres14423 = True


} else {
ifres14423 = False


}

ifres14422 = ifres14423


} else {
ifres14422 = False


}

var ifres14421 Obj

if True == ifres14422 {
ifres14421 = True


} else {
ifres14421 = False


}

ifres14420 = ifres14421


} else {
ifres14420 = False


}

var ifres14419 Obj

if True == ifres14420 {
ifres14419 = True


} else {
ifres14419 = False


}

ifres14418 = ifres14419


} else {
ifres14418 = False


}

if True == ifres14418 {
tmp14322 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4500)
}
__typedArg0 := V4500
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14323 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4500)
}
__typedArg0 := V4500
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14324 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14323)
}
__typedArg0 := tmp14323
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14325 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14324, Nil)
}
__typedArg0 := tmp14324
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14326 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp14325)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp14325
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14327 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14322, tmp14326)
}
__typedArg0 := tmp14322
__typedArg1 := tmp14326
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4curry_1type), tmp14327)
return


} else {
tmp14416 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4500)
}
__typedArg0 := V4500
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14376 Obj

if True == tmp14416 {
tmp14414 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4500)
}
__typedArg0 := V4500
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14415 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14414)
}
__typedArg0 := tmp14414
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14378 Obj

if True == tmp14415 {
tmp14411 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4500)
}
__typedArg0 := V4500
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14412 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp14411)
}
__typedArg0 := tmp14411
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14413 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlist, tmp14412)
}
__typedArg0 := symlist
__typedArg1 := tmp14412
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14380 Obj

if True == tmp14413 {
tmp14408 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4500)
}
__typedArg0 := V4500
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14409 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14408)
}
__typedArg0 := tmp14408
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14410 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14409)
}
__typedArg0 := tmp14409
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14382 Obj

if True == tmp14410 {
tmp14404 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4500)
}
__typedArg0 := V4500
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14405 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14404)
}
__typedArg0 := tmp14404
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14406 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14405)
}
__typedArg0 := tmp14405
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14407 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp14406)
}
__typedArg0 := Nil
__typedArg1 := tmp14406
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14384 Obj

if True == tmp14407 {
tmp14402 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4500)
}
__typedArg0 := V4500
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14403 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14402)
}
__typedArg0 := tmp14402
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14386 Obj

if True == tmp14403 {
tmp14399 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4500)
}
__typedArg0 := V4500
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14400 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp14399)
}
__typedArg0 := tmp14399
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14401 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_a_a_6, tmp14400)
}
__typedArg0 := sym_a_a_6
__typedArg1 := tmp14400
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14388 Obj

if True == tmp14401 {
tmp14396 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4500)
}
__typedArg0 := V4500
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14397 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14396)
}
__typedArg0 := tmp14396
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14398 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14397)
}
__typedArg0 := tmp14397
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14390 Obj

if True == tmp14398 {
tmp14392 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4500)
}
__typedArg0 := V4500
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14393 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14392)
}
__typedArg0 := tmp14392
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14394 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14393)
}
__typedArg0 := tmp14393
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14395 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp14394)
}
__typedArg0 := Nil
__typedArg1 := tmp14394
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14391 Obj

if True == tmp14395 {
ifres14391 = True


} else {
ifres14391 = False


}

ifres14390 = ifres14391


} else {
ifres14390 = False


}

var ifres14389 Obj

if True == ifres14390 {
ifres14389 = True


} else {
ifres14389 = False


}

ifres14388 = ifres14389


} else {
ifres14388 = False


}

var ifres14387 Obj

if True == ifres14388 {
ifres14387 = True


} else {
ifres14387 = False


}

ifres14386 = ifres14387


} else {
ifres14386 = False


}

var ifres14385 Obj

if True == ifres14386 {
ifres14385 = True


} else {
ifres14385 = False


}

ifres14384 = ifres14385


} else {
ifres14384 = False


}

var ifres14383 Obj

if True == ifres14384 {
ifres14383 = True


} else {
ifres14383 = False


}

ifres14382 = ifres14383


} else {
ifres14382 = False


}

var ifres14381 Obj

if True == ifres14382 {
ifres14381 = True


} else {
ifres14381 = False


}

ifres14380 = ifres14381


} else {
ifres14380 = False


}

var ifres14379 Obj

if True == ifres14380 {
ifres14379 = True


} else {
ifres14379 = False


}

ifres14378 = ifres14379


} else {
ifres14378 = False


}

var ifres14377 Obj

if True == ifres14378 {
ifres14377 = True


} else {
ifres14377 = False


}

ifres14376 = ifres14377


} else {
ifres14376 = False


}

if True == ifres14376 {
tmp14328 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4500)
}
__typedArg0 := V4500
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14329 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4500)
}
__typedArg0 := V4500
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14330 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4500)
}
__typedArg0 := V4500
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14331 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14330)
}
__typedArg0 := tmp14330
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14332 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14329, tmp14331)
}
__typedArg0 := tmp14329
__typedArg1 := tmp14331
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14333 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstr, tmp14332)
}
__typedArg0 := symstr
__typedArg1 := tmp14332
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14334 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14333, Nil)
}
__typedArg0 := tmp14333
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14335 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp14334)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp14334
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14336 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14328, tmp14335)
}
__typedArg0 := tmp14328
__typedArg1 := tmp14335
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4curry_1type), tmp14336)
return


} else {
tmp14374 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4500)
}
__typedArg0 := V4500
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14347 Obj

if True == tmp14374 {
tmp14372 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4500)
}
__typedArg0 := V4500
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14373 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14372)
}
__typedArg0 := tmp14372
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14349 Obj

if True == tmp14373 {
tmp14369 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4500)
}
__typedArg0 := V4500
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14370 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp14369)
}
__typedArg0 := tmp14369
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14371 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_d, tmp14370)
}
__typedArg0 := sym_d
__typedArg1 := tmp14370
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14351 Obj

if True == tmp14371 {
tmp14366 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4500)
}
__typedArg0 := V4500
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14367 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14366)
}
__typedArg0 := tmp14366
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14368 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14367)
}
__typedArg0 := tmp14367
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14353 Obj

if True == tmp14368 {
tmp14362 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4500)
}
__typedArg0 := V4500
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14363 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14362)
}
__typedArg0 := tmp14362
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14364 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14363)
}
__typedArg0 := tmp14363
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14365 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14364)
}
__typedArg0 := tmp14364
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14355 Obj

if True == tmp14365 {
tmp14357 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4500)
}
__typedArg0 := V4500
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14358 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14357)
}
__typedArg0 := tmp14357
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14359 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14358)
}
__typedArg0 := tmp14358
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14360 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp14359)
}
__typedArg0 := tmp14359
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14361 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_d, tmp14360)
}
__typedArg0 := sym_d
__typedArg1 := tmp14360
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14356 Obj

if True == tmp14361 {
ifres14356 = True


} else {
ifres14356 = False


}

ifres14355 = ifres14356


} else {
ifres14355 = False


}

var ifres14354 Obj

if True == ifres14355 {
ifres14354 = True


} else {
ifres14354 = False


}

ifres14353 = ifres14354


} else {
ifres14353 = False


}

var ifres14352 Obj

if True == ifres14353 {
ifres14352 = True


} else {
ifres14352 = False


}

ifres14351 = ifres14352


} else {
ifres14351 = False


}

var ifres14350 Obj

if True == ifres14351 {
ifres14350 = True


} else {
ifres14350 = False


}

ifres14349 = ifres14350


} else {
ifres14349 = False


}

var ifres14348 Obj

if True == ifres14349 {
ifres14348 = True


} else {
ifres14348 = False


}

ifres14347 = ifres14348


} else {
ifres14347 = False


}

if True == ifres14347 {
tmp14337 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4500)
}
__typedArg0 := V4500
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14338 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4500)
}
__typedArg0 := V4500
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14339 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14338)
}
__typedArg0 := tmp14338
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14340 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14339, Nil)
}
__typedArg0 := tmp14339
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14341 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_d, tmp14340)
}
__typedArg0 := sym_d
__typedArg1 := tmp14340
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14342 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14337, tmp14341)
}
__typedArg0 := tmp14337
__typedArg1 := tmp14341
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4curry_1type), tmp14342)
return


} else {
tmp14345 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4500)
}
__typedArg0 := V4500
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp14345 {
tmp14343 := MakeNative(func(__e *ControlFlow) {
Z4501 := __e.Get(1)
_ = Z4501
__e.TailApply(PrimFunc(symshen_4curry_1type), Z4501)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp14343, V4500)
return


} else {
__e.Return(V4500)
return
}


}


}


}


}, 1)

tmp14446 := Call(__e, ns2_1set, symshen_4curry_1type, tmp14321)


_ = tmp14446

tmp14447 := MakeNative(func(__e *ControlFlow) {
V4502 := __e.Get(1)
_ = V4502
tmp14536 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4502)
}
__typedArg0 := V4502
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14528 Obj

if True == tmp14536 {
tmp14534 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4502)
}
__typedArg0 := V4502
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14535 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefine, tmp14534)
}
__typedArg0 := symdefine
__typedArg1 := tmp14534
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14530 Obj

if True == tmp14535 {
tmp14532 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4502)
}
__typedArg0 := V4502
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14533 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14532)
}
__typedArg0 := tmp14532
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14531 Obj

if True == tmp14533 {
ifres14531 = True


} else {
ifres14531 = False


}

ifres14530 = ifres14531


} else {
ifres14530 = False


}

var ifres14529 Obj

if True == ifres14530 {
ifres14529 = True


} else {
ifres14529 = False


}

ifres14528 = ifres14529


} else {
ifres14528 = False


}

if True == ifres14528 {
__e.Return(V4502)
return
} else {
tmp14526 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4502)
}
__typedArg0 := V4502
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14507 Obj

if True == tmp14526 {
tmp14524 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4502)
}
__typedArg0 := V4502
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14525 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symtype, tmp14524)
}
__typedArg0 := symtype
__typedArg1 := tmp14524
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14509 Obj

if True == tmp14525 {
tmp14522 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4502)
}
__typedArg0 := V4502
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14523 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14522)
}
__typedArg0 := tmp14522
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14511 Obj

if True == tmp14523 {
tmp14519 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4502)
}
__typedArg0 := V4502
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14520 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14519)
}
__typedArg0 := tmp14519
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14521 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14520)
}
__typedArg0 := tmp14520
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14513 Obj

if True == tmp14521 {
tmp14515 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4502)
}
__typedArg0 := V4502
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14516 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14515)
}
__typedArg0 := tmp14515
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14517 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14516)
}
__typedArg0 := tmp14516
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14518 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp14517)
}
__typedArg0 := Nil
__typedArg1 := tmp14517
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14514 Obj

if True == tmp14518 {
ifres14514 = True


} else {
ifres14514 = False


}

ifres14513 = ifres14514


} else {
ifres14513 = False


}

var ifres14512 Obj

if True == ifres14513 {
ifres14512 = True


} else {
ifres14512 = False


}

ifres14511 = ifres14512


} else {
ifres14511 = False


}

var ifres14510 Obj

if True == ifres14511 {
ifres14510 = True


} else {
ifres14510 = False


}

ifres14509 = ifres14510


} else {
ifres14509 = False


}

var ifres14508 Obj

if True == ifres14509 {
ifres14508 = True


} else {
ifres14508 = False


}

ifres14507 = ifres14508


} else {
ifres14507 = False


}

if True == ifres14507 {
tmp14448 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4502)
}
__typedArg0 := V4502
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14449 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp14448)
}
__typedArg0 := tmp14448
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14450 := Call(__e, PrimFunc(symshen_4curry), tmp14449)


tmp14451 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4502)
}
__typedArg0 := V4502
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14452 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14451)
}
__typedArg0 := tmp14451
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14453 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14450, tmp14452)
}
__typedArg0 := tmp14450
__typedArg1 := tmp14452
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtype, tmp14453)
}
__typedArg0 := symtype
__typedArg1 := tmp14453
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp14505 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4502)
}
__typedArg0 := V4502
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14501 Obj

if True == tmp14505 {
tmp14503 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4502)
}
__typedArg0 := V4502
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14504 := Call(__e, PrimFunc(symshen_4special_2), tmp14503)


var ifres14502 Obj

if True == tmp14504 {
ifres14502 = True


} else {
ifres14502 = False


}

ifres14501 = ifres14502


} else {
ifres14501 = False


}

if True == ifres14501 {
tmp14454 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4502)
}
__typedArg0 := V4502
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14455 := MakeNative(func(__e *ControlFlow) {
Z4503 := __e.Get(1)
_ = Z4503
__e.TailApply(PrimFunc(symshen_4curry), Z4503)
return
}, 1)

tmp14456 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4502)
}
__typedArg0 := V4502
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14457 := Call(__e, PrimFunc(symmap), tmp14455, tmp14456)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14454, tmp14457)
}
__typedArg0 := tmp14454
__typedArg1 := tmp14457
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp14499 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4502)
}
__typedArg0 := V4502
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14495 Obj

if True == tmp14499 {
tmp14497 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4502)
}
__typedArg0 := V4502
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14498 := Call(__e, PrimFunc(symshen_4extraspecial_2), tmp14497)


var ifres14496 Obj

if True == tmp14498 {
ifres14496 = True


} else {
ifres14496 = False


}

ifres14495 = ifres14496


} else {
ifres14495 = False


}

if True == ifres14495 {
__e.Return(V4502)
return
} else {
tmp14493 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4502)
}
__typedArg0 := V4502
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14484 Obj

if True == tmp14493 {
tmp14491 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4502)
}
__typedArg0 := V4502
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14492 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14491)
}
__typedArg0 := tmp14491
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14486 Obj

if True == tmp14492 {
tmp14488 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4502)
}
__typedArg0 := V4502
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14489 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14488)
}
__typedArg0 := tmp14488
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14490 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14489)
}
__typedArg0 := tmp14489
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14487 Obj

if True == tmp14490 {
ifres14487 = True


} else {
ifres14487 = False


}

ifres14486 = ifres14487


} else {
ifres14486 = False


}

var ifres14485 Obj

if True == ifres14486 {
ifres14485 = True


} else {
ifres14485 = False


}

ifres14484 = ifres14485


} else {
ifres14484 = False


}

if True == ifres14484 {
tmp14458 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4502)
}
__typedArg0 := V4502
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14459 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4502)
}
__typedArg0 := V4502
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14460 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp14459)
}
__typedArg0 := tmp14459
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14461 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14460, Nil)
}
__typedArg0 := tmp14460
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14462 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14458, tmp14461)
}
__typedArg0 := tmp14458
__typedArg1 := tmp14461
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14463 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4502)
}
__typedArg0 := V4502
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14464 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14463)
}
__typedArg0 := tmp14463
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14465 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14462, tmp14464)
}
__typedArg0 := tmp14462
__typedArg1 := tmp14464
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4curry), tmp14465)
return


} else {
tmp14482 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4502)
}
__typedArg0 := V4502
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14473 Obj

if True == tmp14482 {
tmp14480 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4502)
}
__typedArg0 := V4502
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14481 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14480)
}
__typedArg0 := tmp14480
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14475 Obj

if True == tmp14481 {
tmp14477 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4502)
}
__typedArg0 := V4502
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14478 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14477)
}
__typedArg0 := tmp14477
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14479 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp14478)
}
__typedArg0 := Nil
__typedArg1 := tmp14478
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14476 Obj

if True == tmp14479 {
ifres14476 = True


} else {
ifres14476 = False


}

ifres14475 = ifres14476


} else {
ifres14475 = False


}

var ifres14474 Obj

if True == ifres14475 {
ifres14474 = True


} else {
ifres14474 = False


}

ifres14473 = ifres14474


} else {
ifres14473 = False


}

if True == ifres14473 {
tmp14466 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4502)
}
__typedArg0 := V4502
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14467 := Call(__e, PrimFunc(symshen_4curry), tmp14466)


tmp14468 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4502)
}
__typedArg0 := V4502
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14469 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp14468)
}
__typedArg0 := tmp14468
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14470 := Call(__e, PrimFunc(symshen_4curry), tmp14469)


tmp14471 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14470, Nil)
}
__typedArg0 := tmp14470
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14467, tmp14471)
}
__typedArg0 := tmp14467
__typedArg1 := tmp14471
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return(V4502)
return
}


}


}


}


}


}


}, 1)

tmp14537 := Call(__e, ns2_1set, symshen_4curry, tmp14447)


_ = tmp14537

tmp14538 := MakeNative(func(__e *ControlFlow) {
V4504 := __e.Get(1)
_ = V4504
tmp14539 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dspecial_d)
}
__typedArg0 := symshen_4_dspecial_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symelement_2), V4504, tmp14539)
return


}, 1)

tmp14540 := Call(__e, ns2_1set, symshen_4special_2, tmp14538)


_ = tmp14540

tmp14541 := MakeNative(func(__e *ControlFlow) {
V4505 := __e.Get(1)
_ = V4505
tmp14542 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dextraspecial_d)
}
__typedArg0 := symshen_4_dextraspecial_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symelement_2), V4505, tmp14542)
return


}, 1)

tmp14543 := Call(__e, ns2_1set, symshen_4extraspecial_2, tmp14541)


_ = tmp14543

tmp14544 := MakeNative(func(__e *ControlFlow) {
V4506 := __e.Get(1)
_ = V4506
V4507 := __e.Get(2)
_ = V4507
V4508 := __e.Get(3)
_ = V4508
V4509 := __e.Get(4)
_ = V4509
V4510 := __e.Get(5)
_ = V4510
V4511 := __e.Get(6)
_ = V4511
tmp14545 := MakeNative(func(__e *ControlFlow) {
W4512 := __e.Get(1)
_ = W4512
tmp14546 := MakeNative(func(__e *ControlFlow) {
W4513 := __e.Get(1)
_ = W4513
tmp14604 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4513, False)
}
__typedArg0 := W4513
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp14604 {
tmp14547 := MakeNative(func(__e *ControlFlow) {
W4514 := __e.Get(1)
_ = W4514
tmp14566 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4514, False)
}
__typedArg0 := W4514
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp14566 {
tmp14548 := MakeNative(func(__e *ControlFlow) {
W4522 := __e.Get(1)
_ = W4522
tmp14558 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4522, False)
}
__typedArg0 := W4522
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp14558 {
tmp14549 := MakeNative(func(__e *ControlFlow) {
W4523 := __e.Get(1)
_ = W4523
tmp14551 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4523, False)
}
__typedArg0 := W4523
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp14551 {
__e.TailApply(PrimFunc(symshen_4unlock), V4509, W4512)
return
} else {
__e.Return(W4523)
return
}


}, 1)

tmp14556 := Call(__e, PrimFunc(symshen_4unlocked_2), V4509)


var ifres14552 Obj

if True == tmp14556 {
tmp14553 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14553

tmp14554 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_ddatatypes_d)
}
__typedArg0 := symshen_4_ddatatypes_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp14555 := Call(__e, PrimFunc(symshen_4search_1user_1datatypes), V4506, V4507, tmp14554, V4508, V4509, W4512, V4511)


ifres14552 = tmp14555


} else {
ifres14552 = False


}

__e.TailApply(tmp14549, ifres14552)
return


} else {
__e.Return(W4522)
return
}


}, 1)

tmp14564 := Call(__e, PrimFunc(symshen_4unlocked_2), V4509)


var ifres14559 Obj

if True == tmp14564 {
tmp14560 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14560

tmp14561 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dspy_d)
}
__typedArg0 := symshen_4_dspy_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp14562 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4show), V4506, V4507, V4508, V4509, W4512, V4511)
return
}, 0)

tmp14563 := Call(__e, PrimFunc(symwhen), tmp14561, V4508, V4509, W4512, tmp14562)


ifres14559 = tmp14563


} else {
ifres14559 = False


}

__e.TailApply(tmp14548, ifres14559)
return


} else {
__e.Return(W4514)
return
}


}, 1)

tmp14602 := Call(__e, PrimFunc(symshen_4unlocked_2), V4509)


var ifres14567 Obj

if True == tmp14602 {
tmp14568 := MakeNative(func(__e *ControlFlow) {
W4515 := __e.Get(1)
_ = W4515
tmp14599 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4515)
}
__typedArg0 := W4515
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp14599 {
tmp14569 := MakeNative(func(__e *ControlFlow) {
W4516 := __e.Get(1)
_ = W4516
tmp14570 := MakeNative(func(__e *ControlFlow) {
W4517 := __e.Get(1)
_ = W4517
tmp14594 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4517)
}
__typedArg0 := W4517
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp14594 {
tmp14571 := MakeNative(func(__e *ControlFlow) {
W4518 := __e.Get(1)
_ = W4518
tmp14572 := MakeNative(func(__e *ControlFlow) {
W4519 := __e.Get(1)
_ = W4519
tmp14589 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4519)
}
__typedArg0 := W4519
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp14589 {
tmp14573 := MakeNative(func(__e *ControlFlow) {
W4520 := __e.Get(1)
_ = W4520
tmp14574 := MakeNative(func(__e *ControlFlow) {
W4521 := __e.Get(1)
_ = W4521
tmp14584 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4521, Nil)
}
__typedArg0 := W4521
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp14584 {
tmp14575 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14575

tmp14576 := Call(__e, PrimFunc(symshen_4deref), W4518, V4508)


tmp14577 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp14578 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp14576, tmp14577)
}
__typedArg0 := tmp14576
__typedArg1 := tmp14577
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

tmp14579 := MakeNative(func(__e *ControlFlow) {
tmp14580 := Call(__e, PrimFunc(symshen_4type_1theory_1enabled_2))


tmp14581 := MakeNative(func(__e *ControlFlow) {
tmp14582 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4516, W4520, V4507, V4508, V4509, W4512, V4511)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V4508, V4509, W4512, tmp14582)
return


}, 0)

__e.TailApply(PrimFunc(symwhen), tmp14580, V4508, V4509, W4512, tmp14581)
return


}, 0)

__e.TailApply(PrimFunc(symwhen), tmp14578, V4508, V4509, W4512, tmp14579)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14585 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4519)
}
__typedArg0 := W4519
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14586 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14585, V4508)


__e.TailApply(tmp14574, tmp14586)
return


}, 1)

tmp14587 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4519)
}
__typedArg0 := W4519
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp14573, tmp14587)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14590 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4517)
}
__typedArg0 := W4517
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14591 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14590, V4508)


__e.TailApply(tmp14572, tmp14591)
return


}, 1)

tmp14592 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4517)
}
__typedArg0 := W4517
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp14571, tmp14592)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14595 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4515)
}
__typedArg0 := W4515
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14596 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14595, V4508)


__e.TailApply(tmp14570, tmp14596)
return


}, 1)

tmp14597 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4515)
}
__typedArg0 := W4515
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp14569, tmp14597)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14600 := Call(__e, PrimFunc(symshen_4lazyderef), V4506, V4508)


tmp14601 := Call(__e, tmp14568, tmp14600)


ifres14567 = tmp14601


} else {
ifres14567 = False


}

__e.TailApply(tmp14547, ifres14567)
return


} else {
__e.Return(W4513)
return
}


}, 1)

tmp14609 := Call(__e, PrimFunc(symshen_4unlocked_2), V4509)


var ifres14605 Obj

if True == tmp14609 {
tmp14606 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14606

tmp14607 := Call(__e, PrimFunc(symshen_4maxinfexceeded_2))


tmp14608 := Call(__e, PrimFunc(symwhen), tmp14607, V4508, V4509, W4512, V4511)


ifres14605 = tmp14608


} else {
ifres14605 = False


}

__e.TailApply(tmp14546, ifres14605)
return


}, 1)

__e.TailApply(tmp14545, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V4510)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V4510
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())
return


}, 6)

tmp14611 := Call(__e, ns2_1set, symshen_4system_1S, tmp14544)


_ = tmp14611

tmp14612 := MakeNative(func(__e *ControlFlow) {
V4530 := __e.Get(1)
_ = V4530
V4531 := __e.Get(2)
_ = V4531
V4532 := __e.Get(3)
_ = V4532
V4533 := __e.Get(4)
_ = V4533
V4534 := __e.Get(5)
_ = V4534
V4535 := __e.Get(6)
_ = V4535
tmp14613 := Call(__e, PrimFunc(symshen_4line))


_ = tmp14613

tmp14614 := Call(__e, PrimFunc(symshen_4deref), V4530, V4532)


tmp14615 := Call(__e, PrimFunc(symshen_4show_1p), tmp14614)


_ = tmp14615

tmp14616 := Call(__e, PrimFunc(symnl), MakeNumber(2))


_ = tmp14616

tmp14617 := Call(__e, PrimFunc(symshen_4deref), V4531, V4532)


tmp14618 := Call(__e, PrimFunc(symshen_4show_1assumptions), tmp14617, MakeNumber(1))


_ = tmp14618

tmp14619 := Call(__e, PrimFunc(symshen_4pause_1for_1user))


_ = tmp14619

__e.Return(False)
return


}, 6)

tmp14620 := Call(__e, ns2_1set, symshen_4show, tmp14612)


_ = tmp14620

tmp14621 := MakeNative(func(__e *ControlFlow) {
tmp14622 := MakeNative(func(__e *ControlFlow) {
W4536 := __e.Get(1)
_ = W4536
tmp14624 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeNumber(1), W4536)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := W4536
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14623 Obj

if True == tmp14624 {
ifres14623 = MakeString("")


} else {
ifres14623 = MakeString("s")


}

tmp14625 := Call(__e, PrimFunc(symshen_4app), ifres14623, MakeString(" \n?- "), symshen_4a)


tmp14627 := Call(__e, PrimFunc(symshen_4app), W4536, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(" inference"))
__typedS1, __typedOK1 := TypedString(tmp14625)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(" inference")
__typedArg1 := tmp14625
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


tmp14628 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("____________________________________________________________ "))
__typedS1, __typedOK1 := TypedString(tmp14627)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("____________________________________________________________ ")
__typedArg1 := tmp14627
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp14629 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp14628, tmp14629)
return


}, 1)

tmp14630 := Call(__e, PrimFunc(syminferences))


__e.TailApply(tmp14622, tmp14630)
return


}, 0)

tmp14631 := Call(__e, ns2_1set, symshen_4line, tmp14621)


_ = tmp14631

tmp14632 := MakeNative(func(__e *ControlFlow) {
V4537 := __e.Get(1)
_ = V4537
tmp14664 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4537)
}
__typedArg0 := V4537
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14643 Obj

if True == tmp14664 {
tmp14662 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4537)
}
__typedArg0 := V4537
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14663 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14662)
}
__typedArg0 := tmp14662
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14645 Obj

if True == tmp14663 {
tmp14659 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4537)
}
__typedArg0 := V4537
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14660 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14659)
}
__typedArg0 := tmp14659
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14661 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14660)
}
__typedArg0 := tmp14660
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14647 Obj

if True == tmp14661 {
tmp14655 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4537)
}
__typedArg0 := V4537
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14656 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14655)
}
__typedArg0 := tmp14655
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14657 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14656)
}
__typedArg0 := tmp14656
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14658 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp14657)
}
__typedArg0 := Nil
__typedArg1 := tmp14657
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14649 Obj

if True == tmp14658 {
tmp14651 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4537)
}
__typedArg0 := V4537
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14652 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp14651)
}
__typedArg0 := tmp14651
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14653 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp14654 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp14652, tmp14653)
}
__typedArg0 := tmp14652
__typedArg1 := tmp14653
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14650 Obj

if True == tmp14654 {
ifres14650 = True


} else {
ifres14650 = False


}

ifres14649 = ifres14650


} else {
ifres14649 = False


}

var ifres14648 Obj

if True == ifres14649 {
ifres14648 = True


} else {
ifres14648 = False


}

ifres14647 = ifres14648


} else {
ifres14647 = False


}

var ifres14646 Obj

if True == ifres14647 {
ifres14646 = True


} else {
ifres14646 = False


}

ifres14645 = ifres14646


} else {
ifres14645 = False


}

var ifres14644 Obj

if True == ifres14645 {
ifres14644 = True


} else {
ifres14644 = False


}

ifres14643 = ifres14644


} else {
ifres14643 = False


}

if True == ifres14643 {
tmp14633 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4537)
}
__typedArg0 := V4537
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14634 := Call(__e, PrimFunc(symshen_4prterm), tmp14633)


_ = tmp14634

tmp14635 := Call(__e, PrimFunc(symstoutput))


tmp14636 := Call(__e, PrimFunc(sympr), MakeString(" : "), tmp14635)


_ = tmp14636

tmp14637 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4537)
}
__typedArg0 := V4537
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14638 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14637)
}
__typedArg0 := tmp14637
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14639 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp14638)
}
__typedArg0 := tmp14638
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14640 := Call(__e, PrimFunc(symshen_4app), tmp14639, MakeString(""), symshen_4r)


tmp14641 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp14640, tmp14641)
return


} else {
__e.TailApply(PrimFunc(symshen_4prterm), V4537)
return
}


}, 1)

tmp14665 := Call(__e, ns2_1set, symshen_4show_1p, tmp14632)


_ = tmp14665

tmp14666 := MakeNative(func(__e *ControlFlow) {
V4538 := __e.Get(1)
_ = V4538
tmp14709 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4538)
}
__typedArg0 := V4538
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14690 Obj

if True == tmp14709 {
tmp14707 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4538)
}
__typedArg0 := V4538
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14708 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcons, tmp14707)
}
__typedArg0 := symcons
__typedArg1 := tmp14707
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14692 Obj

if True == tmp14708 {
tmp14705 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4538)
}
__typedArg0 := V4538
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14706 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14705)
}
__typedArg0 := tmp14705
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14694 Obj

if True == tmp14706 {
tmp14702 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4538)
}
__typedArg0 := V4538
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14703 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14702)
}
__typedArg0 := tmp14702
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14704 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14703)
}
__typedArg0 := tmp14703
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14696 Obj

if True == tmp14704 {
tmp14698 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4538)
}
__typedArg0 := V4538
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14699 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14698)
}
__typedArg0 := tmp14698
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14700 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14699)
}
__typedArg0 := tmp14699
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14701 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp14700)
}
__typedArg0 := Nil
__typedArg1 := tmp14700
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14697 Obj

if True == tmp14701 {
ifres14697 = True


} else {
ifres14697 = False


}

ifres14696 = ifres14697


} else {
ifres14696 = False


}

var ifres14695 Obj

if True == ifres14696 {
ifres14695 = True


} else {
ifres14695 = False


}

ifres14694 = ifres14695


} else {
ifres14694 = False


}

var ifres14693 Obj

if True == ifres14694 {
ifres14693 = True


} else {
ifres14693 = False


}

ifres14692 = ifres14693


} else {
ifres14692 = False


}

var ifres14691 Obj

if True == ifres14692 {
ifres14691 = True


} else {
ifres14691 = False


}

ifres14690 = ifres14691


} else {
ifres14690 = False


}

if True == ifres14690 {
tmp14667 := Call(__e, PrimFunc(symstoutput))


tmp14668 := Call(__e, PrimFunc(sympr), MakeString("["), tmp14667)


_ = tmp14668

tmp14669 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4538)
}
__typedArg0 := V4538
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14670 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp14669)
}
__typedArg0 := tmp14669
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14671 := Call(__e, PrimFunc(symshen_4prterm), tmp14670)


_ = tmp14671

tmp14672 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4538)
}
__typedArg0 := V4538
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14673 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14672)
}
__typedArg0 := tmp14672
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14674 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp14673)
}
__typedArg0 := tmp14673
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14675 := Call(__e, PrimFunc(symshen_4prtl), tmp14674)


_ = tmp14675

tmp14676 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), MakeString("]"), tmp14676)
return


} else {
tmp14688 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4538)
}
__typedArg0 := V4538
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp14688 {
tmp14677 := Call(__e, PrimFunc(symstoutput))


tmp14678 := Call(__e, PrimFunc(sympr), MakeString("("), tmp14677)


_ = tmp14678

tmp14679 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4538)
}
__typedArg0 := V4538
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14680 := Call(__e, PrimFunc(symshen_4prterm), tmp14679)


_ = tmp14680

tmp14681 := MakeNative(func(__e *ControlFlow) {
Z4539 := __e.Get(1)
_ = Z4539
tmp14682 := Call(__e, PrimFunc(symstoutput))


tmp14683 := Call(__e, PrimFunc(sympr), MakeString(" "), tmp14682)


_ = tmp14683

__e.TailApply(PrimFunc(symshen_4prterm), Z4539)
return


}, 1)

tmp14684 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4538)
}
__typedArg0 := V4538
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14685 := Call(__e, PrimFunc(symmap), tmp14681, tmp14684)


_ = tmp14685

tmp14686 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), MakeString(")"), tmp14686)
return


} else {
__e.TailApply(PrimFunc(symprint), V4538)
return
}


}


}, 1)

tmp14710 := Call(__e, ns2_1set, symshen_4prterm, tmp14666)


_ = tmp14710

tmp14711 := MakeNative(func(__e *ControlFlow) {
V4540 := __e.Get(1)
_ = V4540
tmp14744 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V4540)
}
__typedArg0 := Nil
__typedArg1 := V4540
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp14744 {
__e.Return(MakeString(""))
return
} else {
tmp14742 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4540)
}
__typedArg0 := V4540
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14723 Obj

if True == tmp14742 {
tmp14740 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4540)
}
__typedArg0 := V4540
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14741 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcons, tmp14740)
}
__typedArg0 := symcons
__typedArg1 := tmp14740
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14725 Obj

if True == tmp14741 {
tmp14738 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4540)
}
__typedArg0 := V4540
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14739 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14738)
}
__typedArg0 := tmp14738
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14727 Obj

if True == tmp14739 {
tmp14735 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4540)
}
__typedArg0 := V4540
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14736 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14735)
}
__typedArg0 := tmp14735
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14737 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14736)
}
__typedArg0 := tmp14736
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14729 Obj

if True == tmp14737 {
tmp14731 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4540)
}
__typedArg0 := V4540
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14732 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14731)
}
__typedArg0 := tmp14731
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14733 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14732)
}
__typedArg0 := tmp14732
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14734 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp14733)
}
__typedArg0 := Nil
__typedArg1 := tmp14733
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14730 Obj

if True == tmp14734 {
ifres14730 = True


} else {
ifres14730 = False


}

ifres14729 = ifres14730


} else {
ifres14729 = False


}

var ifres14728 Obj

if True == ifres14729 {
ifres14728 = True


} else {
ifres14728 = False


}

ifres14727 = ifres14728


} else {
ifres14727 = False


}

var ifres14726 Obj

if True == ifres14727 {
ifres14726 = True


} else {
ifres14726 = False


}

ifres14725 = ifres14726


} else {
ifres14725 = False


}

var ifres14724 Obj

if True == ifres14725 {
ifres14724 = True


} else {
ifres14724 = False


}

ifres14723 = ifres14724


} else {
ifres14723 = False


}

if True == ifres14723 {
tmp14712 := Call(__e, PrimFunc(symstoutput))


tmp14713 := Call(__e, PrimFunc(sympr), MakeString(" "), tmp14712)


_ = tmp14713

tmp14714 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4540)
}
__typedArg0 := V4540
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14715 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp14714)
}
__typedArg0 := tmp14714
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14716 := Call(__e, PrimFunc(symshen_4prterm), tmp14715)


_ = tmp14716

tmp14717 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4540)
}
__typedArg0 := V4540
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14718 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14717)
}
__typedArg0 := tmp14717
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14719 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp14718)
}
__typedArg0 := tmp14718
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4prtl), tmp14719)
return


} else {
tmp14720 := Call(__e, PrimFunc(symstoutput))


tmp14721 := Call(__e, PrimFunc(sympr), MakeString(" | "), tmp14720)


_ = tmp14721

__e.TailApply(PrimFunc(symshen_4prterm), V4540)
return


}


}


}, 1)

tmp14745 := Call(__e, ns2_1set, symshen_4prtl, tmp14711)


_ = tmp14745

tmp14746 := MakeNative(func(__e *ControlFlow) {
V4547 := __e.Get(1)
_ = V4547
V4548 := __e.Get(2)
_ = V4548
tmp14759 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V4547)
}
__typedArg0 := Nil
__typedArg1 := V4547
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp14759 {
tmp14747 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), MakeString("\n> "), tmp14747)
return


} else {
tmp14757 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4547)
}
__typedArg0 := V4547
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp14757 {
tmp14748 := Call(__e, PrimFunc(symshen_4app), V4548, MakeString(". "), symshen_4a)


tmp14749 := Call(__e, PrimFunc(symstoutput))


tmp14750 := Call(__e, PrimFunc(sympr), tmp14748, tmp14749)


_ = tmp14750

tmp14751 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4547)
}
__typedArg0 := V4547
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14752 := Call(__e, PrimFunc(symshen_4show_1p), tmp14751)


_ = tmp14752

tmp14753 := Call(__e, PrimFunc(symnl), MakeNumber(1))


_ = tmp14753

tmp14754 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4547)
}
__typedArg0 := V4547
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4show_1assumptions), tmp14754, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V4548)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V4548
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.show-assumptions"))
}
__typedArg0 := MakeString("implementation error in shen.show-assumptions")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 2)

tmp14760 := Call(__e, ns2_1set, symshen_4show_1assumptions, tmp14746)


_ = tmp14760

tmp14761 := MakeNative(func(__e *ControlFlow) {
tmp14762 := MakeNative(func(__e *ControlFlow) {
W4549 := __e.Get(1)
_ = W4549
tmp14764 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4549, MakeNumber(94))
}
__typedArg0 := W4549
__typedArg1 := MakeNumber(94)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp14764 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("input aborted\n"))
}
__typedArg0 := MakeString("input aborted\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
} else {
__e.TailApply(PrimFunc(symnl), MakeNumber(1))
return
}


}, 1)

tmp14765 := Call(__e, PrimFunc(symstinput))


tmp14766 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symread_1byte) {
return PrimReadByte(tmp14765)
}
__typedArg0 := tmp14765
return Call(__e, PrimFunc(symread_1byte), __typedArg0)
})()

__e.TailApply(tmp14762, tmp14766)
return


}, 0)

tmp14767 := Call(__e, ns2_1set, symshen_4pause_1for_1user, tmp14761)


_ = tmp14767

tmp14768 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dshen_1type_1theory_1enabled_2_d)
}
__typedArg0 := symshen_4_dshen_1type_1theory_1enabled_2_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp14769 := Call(__e, ns2_1set, symshen_4type_1theory_1enabled_2, tmp14768)


_ = tmp14769

tmp14770 := MakeNative(func(__e *ControlFlow) {
tmp14772 := Call(__e, PrimFunc(syminferences))


tmp14773 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dmaxinferences_d)
}
__typedArg0 := symshen_4_dmaxinferences_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_6) {
__typedN0, __typedOK0 := TypedFloat64(tmp14772)
__typedN1, __typedOK1 := TypedFloat64(tmp14773)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_6) {
return TypedMaterializeBoolean((__typedN0 > __typedN1))
}}
__typedArg0 := tmp14772
__typedArg1 := tmp14773
return Call(__e, PrimFunc(sym_6), __typedArg0, __typedArg1)
})() {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("maximum inferences exceeded"))
}
__typedArg0 := MakeString("maximum inferences exceeded")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
} else {
__e.Return(False)
return
}


}, 0)

tmp14775 := Call(__e, ns2_1set, symshen_4maxinfexceeded_2, tmp14770)


_ = tmp14775

tmp14776 := MakeNative(func(__e *ControlFlow) {
V4550 := __e.Get(1)
_ = V4550
V4551 := __e.Get(2)
_ = V4551
V4552 := __e.Get(3)
_ = V4552
V4553 := __e.Get(4)
_ = V4553
V4554 := __e.Get(5)
_ = V4554
V4555 := __e.Get(6)
_ = V4555
V4556 := __e.Get(7)
_ = V4556
tmp14777 := MakeNative(func(__e *ControlFlow) {
W4557 := __e.Get(1)
_ = W4557
tmp14778 := MakeNative(func(__e *ControlFlow) {
W4558 := __e.Get(1)
_ = W4558
tmp15693 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4558, False)
}
__typedArg0 := W4558
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15693 {
tmp14779 := MakeNative(func(__e *ControlFlow) {
W4559 := __e.Get(1)
_ = W4559
tmp15683 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4559, False)
}
__typedArg0 := W4559
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15683 {
tmp14780 := MakeNative(func(__e *ControlFlow) {
W4560 := __e.Get(1)
_ = W4560
tmp15677 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4560, False)
}
__typedArg0 := W4560
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15677 {
tmp14781 := MakeNative(func(__e *ControlFlow) {
W4561 := __e.Get(1)
_ = W4561
tmp15658 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4561, False)
}
__typedArg0 := W4561
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15658 {
tmp14782 := MakeNative(func(__e *ControlFlow) {
W4565 := __e.Get(1)
_ = W4565
tmp15625 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4565, False)
}
__typedArg0 := W4565
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15625 {
tmp14783 := MakeNative(func(__e *ControlFlow) {
W4571 := __e.Get(1)
_ = W4571
tmp15598 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4571, False)
}
__typedArg0 := W4571
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15598 {
tmp14784 := MakeNative(func(__e *ControlFlow) {
W4577 := __e.Get(1)
_ = W4577
tmp15563 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4577, False)
}
__typedArg0 := W4577
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15563 {
tmp14785 := MakeNative(func(__e *ControlFlow) {
W4584 := __e.Get(1)
_ = W4584
tmp15532 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4584, False)
}
__typedArg0 := W4584
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15532 {
tmp14786 := MakeNative(func(__e *ControlFlow) {
W4591 := __e.Get(1)
_ = W4591
tmp15447 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4591, False)
}
__typedArg0 := W4591
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15447 {
tmp14787 := MakeNative(func(__e *ControlFlow) {
W4612 := __e.Get(1)
_ = W4612
tmp15341 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4612, False)
}
__typedArg0 := W4612
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15341 {
tmp14788 := MakeNative(func(__e *ControlFlow) {
W4640 := __e.Get(1)
_ = W4640
tmp15256 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4640, False)
}
__typedArg0 := W4640
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15256 {
tmp14789 := MakeNative(func(__e *ControlFlow) {
W4661 := __e.Get(1)
_ = W4661
tmp15213 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4661, False)
}
__typedArg0 := W4661
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15213 {
tmp14790 := MakeNative(func(__e *ControlFlow) {
W4671 := __e.Get(1)
_ = W4671
tmp15089 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4671, False)
}
__typedArg0 := W4671
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15089 {
tmp14791 := MakeNative(func(__e *ControlFlow) {
W4701 := __e.Get(1)
_ = W4701
tmp15025 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4701, False)
}
__typedArg0 := W4701
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15025 {
tmp14792 := MakeNative(func(__e *ControlFlow) {
W4714 := __e.Get(1)
_ = W4714
tmp14937 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4714, False)
}
__typedArg0 := W4714
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp14937 {
tmp14793 := MakeNative(func(__e *ControlFlow) {
W4735 := __e.Get(1)
_ = W4735
tmp14899 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4735, False)
}
__typedArg0 := W4735
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp14899 {
tmp14794 := MakeNative(func(__e *ControlFlow) {
W4743 := __e.Get(1)
_ = W4743
tmp14859 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4743, False)
}
__typedArg0 := W4743
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp14859 {
tmp14795 := MakeNative(func(__e *ControlFlow) {
W4751 := __e.Get(1)
_ = W4751
tmp14821 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4751, False)
}
__typedArg0 := W4751
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp14821 {
tmp14796 := MakeNative(func(__e *ControlFlow) {
W4759 := __e.Get(1)
_ = W4759
tmp14810 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4759, False)
}
__typedArg0 := W4759
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp14810 {
tmp14797 := MakeNative(func(__e *ControlFlow) {
W4761 := __e.Get(1)
_ = W4761
tmp14799 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4761, False)
}
__typedArg0 := W4761
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp14799 {
__e.TailApply(PrimFunc(symshen_4unlock), V4554, W4557)
return
} else {
__e.Return(W4761)
return
}


}, 1)

tmp14808 := Call(__e, PrimFunc(symshen_4unlocked_2), V4554)


var ifres14800 Obj

if True == tmp14808 {
tmp14801 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14801

tmp14802 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp14803 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V4551, Nil)
}
__typedArg0 := V4551
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14804 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14802, tmp14803)
}
__typedArg0 := tmp14802
__typedArg1 := tmp14803
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14805 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V4550, tmp14804)
}
__typedArg0 := V4550
__typedArg1 := tmp14804
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14806 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_ddatatypes_d)
}
__typedArg0 := symshen_4_ddatatypes_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp14807 := Call(__e, PrimFunc(symshen_4search_1user_1datatypes), tmp14805, V4552, tmp14806, V4553, V4554, W4557, V4556)


ifres14800 = tmp14807


} else {
ifres14800 = False


}

__e.TailApply(tmp14797, ifres14800)
return


} else {
__e.Return(W4759)
return
}


}, 1)

tmp14819 := Call(__e, PrimFunc(symshen_4unlocked_2), V4554)


var ifres14811 Obj

if True == tmp14819 {
tmp14812 := MakeNative(func(__e *ControlFlow) {
W4760 := __e.Get(1)
_ = W4760
tmp14813 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14813

tmp14814 := MakeNative(func(__e *ControlFlow) {
tmp14815 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), V4550, V4551, W4760, V4553, V4554, W4557, V4556)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V4553, V4554, W4557, tmp14815)
return


}, 0)

tmp14816 := Call(__e, PrimFunc(symshen_4l_1rules), V4552, W4760, False, V4553, V4554, W4557, tmp14814)


__e.TailApply(PrimFunc(symshen_4gc), V4553, tmp14816)
return


}, 1)

tmp14817 := Call(__e, PrimFunc(symshen_4newpv), V4553)


tmp14818 := Call(__e, tmp14812, tmp14817)


ifres14811 = tmp14818


} else {
ifres14811 = False


}

__e.TailApply(tmp14796, ifres14811)
return


} else {
__e.Return(W4751)
return
}


}, 1)

tmp14857 := Call(__e, PrimFunc(symshen_4unlocked_2), V4554)


var ifres14822 Obj

if True == tmp14857 {
tmp14823 := MakeNative(func(__e *ControlFlow) {
W4752 := __e.Get(1)
_ = W4752
tmp14854 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4752)
}
__typedArg0 := W4752
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp14854 {
tmp14824 := MakeNative(func(__e *ControlFlow) {
W4753 := __e.Get(1)
_ = W4753
tmp14850 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4753, symset)
}
__typedArg0 := W4753
__typedArg1 := symset
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp14850 {
tmp14825 := MakeNative(func(__e *ControlFlow) {
W4754 := __e.Get(1)
_ = W4754
tmp14846 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4754)
}
__typedArg0 := W4754
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp14846 {
tmp14826 := MakeNative(func(__e *ControlFlow) {
W4755 := __e.Get(1)
_ = W4755
tmp14827 := MakeNative(func(__e *ControlFlow) {
W4756 := __e.Get(1)
_ = W4756
tmp14841 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4756)
}
__typedArg0 := W4756
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp14841 {
tmp14828 := MakeNative(func(__e *ControlFlow) {
W4757 := __e.Get(1)
_ = W4757
tmp14829 := MakeNative(func(__e *ControlFlow) {
W4758 := __e.Get(1)
_ = W4758
tmp14836 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4758, Nil)
}
__typedArg0 := W4758
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp14836 {
tmp14830 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14830

tmp14831 := MakeNative(func(__e *ControlFlow) {
tmp14832 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W4755, Nil)
}
__typedArg0 := W4755
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14833 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvalue, tmp14832)
}
__typedArg0 := symvalue
__typedArg1 := tmp14832
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14834 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4757, V4551, V4552, V4553, V4554, W4557, V4556)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), tmp14833, V4551, V4552, V4553, V4554, W4557, tmp14834)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4755, symsymbol, V4552, V4553, V4554, W4557, tmp14831)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14837 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4756)
}
__typedArg0 := W4756
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14838 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14837, V4553)


__e.TailApply(tmp14829, tmp14838)
return


}, 1)

tmp14839 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4756)
}
__typedArg0 := W4756
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp14828, tmp14839)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14842 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4754)
}
__typedArg0 := W4754
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14843 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14842, V4553)


__e.TailApply(tmp14827, tmp14843)
return


}, 1)

tmp14844 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4754)
}
__typedArg0 := W4754
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp14826, tmp14844)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14847 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4752)
}
__typedArg0 := W4752
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14848 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14847, V4553)


__e.TailApply(tmp14825, tmp14848)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14851 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4752)
}
__typedArg0 := W4752
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14852 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14851, V4553)


__e.TailApply(tmp14824, tmp14852)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14855 := Call(__e, PrimFunc(symshen_4lazyderef), V4550, V4553)


tmp14856 := Call(__e, tmp14823, tmp14855)


ifres14822 = tmp14856


} else {
ifres14822 = False


}

__e.TailApply(tmp14795, ifres14822)
return


} else {
__e.Return(W4743)
return
}


}, 1)

tmp14897 := Call(__e, PrimFunc(symshen_4unlocked_2), V4554)


var ifres14860 Obj

if True == tmp14897 {
tmp14861 := MakeNative(func(__e *ControlFlow) {
W4744 := __e.Get(1)
_ = W4744
tmp14894 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4744)
}
__typedArg0 := W4744
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp14894 {
tmp14862 := MakeNative(func(__e *ControlFlow) {
W4745 := __e.Get(1)
_ = W4745
tmp14890 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4745, symshen_4input_1h_7)
}
__typedArg0 := W4745
__typedArg1 := symshen_4input_1h_7
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp14890 {
tmp14863 := MakeNative(func(__e *ControlFlow) {
W4746 := __e.Get(1)
_ = W4746
tmp14886 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4746)
}
__typedArg0 := W4746
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp14886 {
tmp14864 := MakeNative(func(__e *ControlFlow) {
W4747 := __e.Get(1)
_ = W4747
tmp14865 := MakeNative(func(__e *ControlFlow) {
W4748 := __e.Get(1)
_ = W4748
tmp14881 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4748)
}
__typedArg0 := W4748
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp14881 {
tmp14866 := MakeNative(func(__e *ControlFlow) {
W4749 := __e.Get(1)
_ = W4749
tmp14867 := MakeNative(func(__e *ControlFlow) {
W4750 := __e.Get(1)
_ = W4750
tmp14876 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4750, Nil)
}
__typedArg0 := W4750
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp14876 {
tmp14868 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14868

tmp14869 := Call(__e, PrimFunc(symshen_4deref), W4747, V4553)


tmp14870 := Call(__e, PrimFunc(symshen_4rdecons), tmp14869)


tmp14871 := Call(__e, PrimFunc(symshen_4rectify_1type), tmp14870)


tmp14872 := MakeNative(func(__e *ControlFlow) {
tmp14873 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symin, Nil)
}
__typedArg0 := symin
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14874 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp14873)
}
__typedArg0 := symstream
__typedArg1 := tmp14873
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4749, tmp14874, V4552, V4553, V4554, W4557, V4556)
return


}, 0)

__e.TailApply(PrimFunc(symis), V4551, tmp14871, V4553, V4554, W4557, tmp14872)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14877 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4748)
}
__typedArg0 := W4748
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14878 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14877, V4553)


__e.TailApply(tmp14867, tmp14878)
return


}, 1)

tmp14879 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4748)
}
__typedArg0 := W4748
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp14866, tmp14879)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14882 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4746)
}
__typedArg0 := W4746
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14883 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14882, V4553)


__e.TailApply(tmp14865, tmp14883)
return


}, 1)

tmp14884 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4746)
}
__typedArg0 := W4746
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp14864, tmp14884)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14887 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4744)
}
__typedArg0 := W4744
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14888 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14887, V4553)


__e.TailApply(tmp14863, tmp14888)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14891 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4744)
}
__typedArg0 := W4744
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14892 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14891, V4553)


__e.TailApply(tmp14862, tmp14892)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14895 := Call(__e, PrimFunc(symshen_4lazyderef), V4550, V4553)


tmp14896 := Call(__e, tmp14861, tmp14895)


ifres14860 = tmp14896


} else {
ifres14860 = False


}

__e.TailApply(tmp14794, ifres14860)
return


} else {
__e.Return(W4735)
return
}


}, 1)

tmp14935 := Call(__e, PrimFunc(symshen_4unlocked_2), V4554)


var ifres14900 Obj

if True == tmp14935 {
tmp14901 := MakeNative(func(__e *ControlFlow) {
W4736 := __e.Get(1)
_ = W4736
tmp14932 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4736)
}
__typedArg0 := W4736
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp14932 {
tmp14902 := MakeNative(func(__e *ControlFlow) {
W4737 := __e.Get(1)
_ = W4737
tmp14928 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4737, symtype)
}
__typedArg0 := W4737
__typedArg1 := symtype
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp14928 {
tmp14903 := MakeNative(func(__e *ControlFlow) {
W4738 := __e.Get(1)
_ = W4738
tmp14924 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4738)
}
__typedArg0 := W4738
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp14924 {
tmp14904 := MakeNative(func(__e *ControlFlow) {
W4739 := __e.Get(1)
_ = W4739
tmp14905 := MakeNative(func(__e *ControlFlow) {
W4740 := __e.Get(1)
_ = W4740
tmp14919 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4740)
}
__typedArg0 := W4740
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp14919 {
tmp14906 := MakeNative(func(__e *ControlFlow) {
W4741 := __e.Get(1)
_ = W4741
tmp14907 := MakeNative(func(__e *ControlFlow) {
W4742 := __e.Get(1)
_ = W4742
tmp14914 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4742, Nil)
}
__typedArg0 := W4742
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp14914 {
tmp14908 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14908

tmp14909 := MakeNative(func(__e *ControlFlow) {
tmp14910 := Call(__e, PrimFunc(symshen_4deref), W4741, V4553)


tmp14911 := Call(__e, PrimFunc(symshen_4rectify_1type), tmp14910)


tmp14912 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4739, V4551, V4552, V4553, V4554, W4557, V4556)
return
}, 0)

__e.TailApply(PrimFunc(symis_b), tmp14911, V4551, V4553, V4554, W4557, tmp14912)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V4553, V4554, W4557, tmp14909)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14915 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4740)
}
__typedArg0 := W4740
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14916 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14915, V4553)


__e.TailApply(tmp14907, tmp14916)
return


}, 1)

tmp14917 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4740)
}
__typedArg0 := W4740
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp14906, tmp14917)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14920 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4738)
}
__typedArg0 := W4738
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14921 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14920, V4553)


__e.TailApply(tmp14905, tmp14921)
return


}, 1)

tmp14922 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4738)
}
__typedArg0 := W4738
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp14904, tmp14922)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14925 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4736)
}
__typedArg0 := W4736
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14926 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14925, V4553)


__e.TailApply(tmp14903, tmp14926)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14929 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4736)
}
__typedArg0 := W4736
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14930 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14929, V4553)


__e.TailApply(tmp14902, tmp14930)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14933 := Call(__e, PrimFunc(symshen_4lazyderef), V4550, V4553)


tmp14934 := Call(__e, tmp14901, tmp14933)


ifres14900 = tmp14934


} else {
ifres14900 = False


}

__e.TailApply(tmp14793, ifres14900)
return


} else {
__e.Return(W4714)
return
}


}, 1)

tmp15023 := Call(__e, PrimFunc(symshen_4unlocked_2), V4554)


var ifres14938 Obj

if True == tmp15023 {
tmp14939 := MakeNative(func(__e *ControlFlow) {
W4715 := __e.Get(1)
_ = W4715
tmp15020 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4715)
}
__typedArg0 := W4715
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15020 {
tmp14940 := MakeNative(func(__e *ControlFlow) {
W4716 := __e.Get(1)
_ = W4716
tmp15016 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4716, symopen)
}
__typedArg0 := W4716
__typedArg1 := symopen
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15016 {
tmp14941 := MakeNative(func(__e *ControlFlow) {
W4717 := __e.Get(1)
_ = W4717
tmp15012 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4717)
}
__typedArg0 := W4717
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15012 {
tmp14942 := MakeNative(func(__e *ControlFlow) {
W4718 := __e.Get(1)
_ = W4718
tmp14943 := MakeNative(func(__e *ControlFlow) {
W4719 := __e.Get(1)
_ = W4719
tmp15007 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4719)
}
__typedArg0 := W4719
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15007 {
tmp14944 := MakeNative(func(__e *ControlFlow) {
W4720 := __e.Get(1)
_ = W4720
tmp14945 := MakeNative(func(__e *ControlFlow) {
W4721 := __e.Get(1)
_ = W4721
tmp15002 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4721, Nil)
}
__typedArg0 := W4721
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15002 {
tmp14946 := MakeNative(func(__e *ControlFlow) {
W4722 := __e.Get(1)
_ = W4722
tmp14947 := MakeNative(func(__e *ControlFlow) {
W4723 := __e.Get(1)
_ = W4723
tmp14991 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4722)
}
__typedArg0 := W4722
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp14991 {
tmp14948 := MakeNative(func(__e *ControlFlow) {
W4725 := __e.Get(1)
_ = W4725
tmp14949 := MakeNative(func(__e *ControlFlow) {
W4726 := __e.Get(1)
_ = W4726
tmp14953 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4725, symstream)
}
__typedArg0 := W4725
__typedArg1 := symstream
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp14953 {
__e.TailApply(PrimFunc(symthaw), W4726)
return
} else {
tmp14951 := Call(__e, PrimFunc(symshen_4pvar_2), W4725)


if True == tmp14951 {
__e.TailApply(PrimFunc(symshen_4bind_b), W4725, symstream, V4553, W4726)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp14954 := MakeNative(func(__e *ControlFlow) {
tmp14955 := MakeNative(func(__e *ControlFlow) {
W4727 := __e.Get(1)
_ = W4727
tmp14956 := MakeNative(func(__e *ControlFlow) {
W4728 := __e.Get(1)
_ = W4728
tmp14976 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4727)
}
__typedArg0 := W4727
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp14976 {
tmp14957 := MakeNative(func(__e *ControlFlow) {
W4730 := __e.Get(1)
_ = W4730
tmp14958 := MakeNative(func(__e *ControlFlow) {
W4731 := __e.Get(1)
_ = W4731
tmp14959 := MakeNative(func(__e *ControlFlow) {
W4732 := __e.Get(1)
_ = W4732
tmp14963 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4731, Nil)
}
__typedArg0 := W4731
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp14963 {
__e.TailApply(PrimFunc(symthaw), W4732)
return
} else {
tmp14961 := Call(__e, PrimFunc(symshen_4pvar_2), W4731)


if True == tmp14961 {
__e.TailApply(PrimFunc(symshen_4bind_b), W4731, Nil, V4553, W4732)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp14964 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W4728, W4730)
return
}, 0)

__e.TailApply(tmp14959, tmp14964)
return


}, 1)

tmp14965 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4727)
}
__typedArg0 := W4727
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14966 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14965, V4553)


__e.TailApply(tmp14958, tmp14966)
return


}, 1)

tmp14967 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4727)
}
__typedArg0 := W4727
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp14957, tmp14967)
return


} else {
tmp14974 := Call(__e, PrimFunc(symshen_4pvar_2), W4727)


if True == tmp14974 {
tmp14968 := MakeNative(func(__e *ControlFlow) {
W4733 := __e.Get(1)
_ = W4733
tmp14969 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W4733, Nil)
}
__typedArg0 := W4733
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14970 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W4728, W4733)
return
}, 0)

tmp14971 := Call(__e, PrimFunc(symshen_4bind_b), W4727, tmp14969, V4553, tmp14970)


__e.TailApply(PrimFunc(symshen_4gc), V4553, tmp14971)
return


}, 1)

tmp14972 := Call(__e, PrimFunc(symshen_4newpv), V4553)


__e.TailApply(tmp14968, tmp14972)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp14977 := MakeNative(func(__e *ControlFlow) {
Z4729 := __e.Get(1)
_ = Z4729
__e.TailApply(W4723, Z4729)
return
}, 1)

__e.TailApply(tmp14956, tmp14977)
return


}, 1)

tmp14978 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4722)
}
__typedArg0 := W4722
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14979 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14978, V4553)


__e.TailApply(tmp14955, tmp14979)
return


}, 0)

__e.TailApply(tmp14949, tmp14954)
return


}, 1)

tmp14980 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4722)
}
__typedArg0 := W4722
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14981 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14980, V4553)


__e.TailApply(tmp14948, tmp14981)
return


} else {
tmp14989 := Call(__e, PrimFunc(symshen_4pvar_2), W4722)


if True == tmp14989 {
tmp14982 := MakeNative(func(__e *ControlFlow) {
W4734 := __e.Get(1)
_ = W4734
tmp14983 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W4734, Nil)
}
__typedArg0 := W4734
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14984 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp14983)
}
__typedArg0 := symstream
__typedArg1 := tmp14983
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14985 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W4723, W4734)
return
}, 0)

tmp14986 := Call(__e, PrimFunc(symshen_4bind_b), W4722, tmp14984, V4553, tmp14985)


__e.TailApply(PrimFunc(symshen_4gc), V4553, tmp14986)
return


}, 1)

tmp14987 := Call(__e, PrimFunc(symshen_4newpv), V4553)


__e.TailApply(tmp14982, tmp14987)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp14992 := MakeNative(func(__e *ControlFlow) {
Z4724 := __e.Get(1)
_ = Z4724
tmp14993 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14993

tmp14994 := MakeNative(func(__e *ControlFlow) {
tmp14995 := Call(__e, PrimFunc(symshen_4lazyderef), Z4724, V4553)


tmp14996 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symout, Nil)
}
__typedArg0 := symout
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14997 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symin, tmp14996)
}
__typedArg0 := symin
__typedArg1 := tmp14996
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14998 := Call(__e, PrimFunc(symelement_2), tmp14995, tmp14997)


tmp14999 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4718, symstring, V4552, V4553, V4554, W4557, V4556)
return
}, 0)

__e.TailApply(PrimFunc(symwhen), tmp14998, V4553, V4554, W4557, tmp14999)
return


}, 0)

__e.TailApply(PrimFunc(symis_b), W4720, Z4724, V4553, V4554, W4557, tmp14994)
return


}, 1)

__e.TailApply(tmp14947, tmp14992)
return


}, 1)

tmp15000 := Call(__e, PrimFunc(symshen_4lazyderef), V4551, V4553)


__e.TailApply(tmp14946, tmp15000)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15003 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4719)
}
__typedArg0 := W4719
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15004 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15003, V4553)


__e.TailApply(tmp14945, tmp15004)
return


}, 1)

tmp15005 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4719)
}
__typedArg0 := W4719
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp14944, tmp15005)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15008 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4717)
}
__typedArg0 := W4717
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15009 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15008, V4553)


__e.TailApply(tmp14943, tmp15009)
return


}, 1)

tmp15010 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4717)
}
__typedArg0 := W4717
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp14942, tmp15010)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15013 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4715)
}
__typedArg0 := W4715
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15014 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15013, V4553)


__e.TailApply(tmp14941, tmp15014)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15017 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4715)
}
__typedArg0 := W4715
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15018 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15017, V4553)


__e.TailApply(tmp14940, tmp15018)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15021 := Call(__e, PrimFunc(symshen_4lazyderef), V4550, V4553)


tmp15022 := Call(__e, tmp14939, tmp15021)


ifres14938 = tmp15022


} else {
ifres14938 = False


}

__e.TailApply(tmp14792, ifres14938)
return


} else {
__e.Return(W4701)
return
}


}, 1)

tmp15087 := Call(__e, PrimFunc(symshen_4unlocked_2), V4554)


var ifres15026 Obj

if True == tmp15087 {
tmp15027 := MakeNative(func(__e *ControlFlow) {
W4702 := __e.Get(1)
_ = W4702
tmp15084 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4702)
}
__typedArg0 := W4702
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15084 {
tmp15028 := MakeNative(func(__e *ControlFlow) {
W4703 := __e.Get(1)
_ = W4703
tmp15080 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4703, symlet)
}
__typedArg0 := W4703
__typedArg1 := symlet
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15080 {
tmp15029 := MakeNative(func(__e *ControlFlow) {
W4704 := __e.Get(1)
_ = W4704
tmp15076 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4704)
}
__typedArg0 := W4704
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15076 {
tmp15030 := MakeNative(func(__e *ControlFlow) {
W4705 := __e.Get(1)
_ = W4705
tmp15031 := MakeNative(func(__e *ControlFlow) {
W4706 := __e.Get(1)
_ = W4706
tmp15071 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4706)
}
__typedArg0 := W4706
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15071 {
tmp15032 := MakeNative(func(__e *ControlFlow) {
W4707 := __e.Get(1)
_ = W4707
tmp15033 := MakeNative(func(__e *ControlFlow) {
W4708 := __e.Get(1)
_ = W4708
tmp15066 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4708)
}
__typedArg0 := W4708
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15066 {
tmp15034 := MakeNative(func(__e *ControlFlow) {
W4709 := __e.Get(1)
_ = W4709
tmp15035 := MakeNative(func(__e *ControlFlow) {
W4710 := __e.Get(1)
_ = W4710
tmp15061 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4710, Nil)
}
__typedArg0 := W4710
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15061 {
tmp15036 := MakeNative(func(__e *ControlFlow) {
W4711 := __e.Get(1)
_ = W4711
tmp15037 := MakeNative(func(__e *ControlFlow) {
W4712 := __e.Get(1)
_ = W4712
tmp15038 := MakeNative(func(__e *ControlFlow) {
W4713 := __e.Get(1)
_ = W4713
tmp15039 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15039

tmp15040 := MakeNative(func(__e *ControlFlow) {
tmp15041 := Call(__e, PrimFunc(symshen_4lazyderef), W4705, V4553)


tmp15042 := Call(__e, PrimFunc(symshen_4freshterm), tmp15041)


tmp15043 := MakeNative(func(__e *ControlFlow) {
tmp15044 := Call(__e, PrimFunc(symshen_4lazyderef), W4705, V4553)


tmp15045 := Call(__e, PrimFunc(symshen_4lazyderef), W4712, V4553)


tmp15046 := Call(__e, PrimFunc(symshen_4lazyderef), W4709, V4553)


tmp15047 := Call(__e, PrimFunc(symshen_4beta), tmp15044, tmp15045, tmp15046)


tmp15048 := MakeNative(func(__e *ControlFlow) {
tmp15049 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp15050 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W4713, Nil)
}
__typedArg0 := W4713
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15051 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp15049, tmp15050)
}
__typedArg0 := tmp15049
__typedArg1 := tmp15050
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15052 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W4712, tmp15051)
}
__typedArg0 := W4712
__typedArg1 := tmp15051
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15053 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp15052, V4552)
}
__typedArg0 := tmp15052
__typedArg1 := V4552
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4711, V4551, tmp15053, V4553, V4554, W4557, V4556)
return


}, 0)

__e.TailApply(PrimFunc(symbind), W4711, tmp15047, V4553, V4554, W4557, tmp15048)
return


}, 0)

__e.TailApply(PrimFunc(symbind), W4712, tmp15042, V4553, V4554, W4557, tmp15043)
return


}, 0)

tmp15054 := Call(__e, PrimFunc(symshen_4system_1S_1h), W4707, W4713, V4552, V4553, V4554, W4557, tmp15040)


__e.TailApply(PrimFunc(symshen_4gc), V4553, tmp15054)
return


}, 1)

tmp15055 := Call(__e, PrimFunc(symshen_4newpv), V4553)


tmp15056 := Call(__e, tmp15038, tmp15055)


__e.TailApply(PrimFunc(symshen_4gc), V4553, tmp15056)
return


}, 1)

tmp15057 := Call(__e, PrimFunc(symshen_4newpv), V4553)


tmp15058 := Call(__e, tmp15037, tmp15057)


__e.TailApply(PrimFunc(symshen_4gc), V4553, tmp15058)
return


}, 1)

tmp15059 := Call(__e, PrimFunc(symshen_4newpv), V4553)


__e.TailApply(tmp15036, tmp15059)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15062 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4708)
}
__typedArg0 := W4708
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15063 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15062, V4553)


__e.TailApply(tmp15035, tmp15063)
return


}, 1)

tmp15064 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4708)
}
__typedArg0 := W4708
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp15034, tmp15064)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15067 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4706)
}
__typedArg0 := W4706
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15068 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15067, V4553)


__e.TailApply(tmp15033, tmp15068)
return


}, 1)

tmp15069 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4706)
}
__typedArg0 := W4706
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp15032, tmp15069)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15072 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4704)
}
__typedArg0 := W4704
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15073 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15072, V4553)


__e.TailApply(tmp15031, tmp15073)
return


}, 1)

tmp15074 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4704)
}
__typedArg0 := W4704
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp15030, tmp15074)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15077 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4702)
}
__typedArg0 := W4702
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15078 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15077, V4553)


__e.TailApply(tmp15029, tmp15078)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15081 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4702)
}
__typedArg0 := W4702
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15082 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15081, V4553)


__e.TailApply(tmp15028, tmp15082)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15085 := Call(__e, PrimFunc(symshen_4lazyderef), V4550, V4553)


tmp15086 := Call(__e, tmp15027, tmp15085)


ifres15026 = tmp15086


} else {
ifres15026 = False


}

__e.TailApply(tmp14791, ifres15026)
return


} else {
__e.Return(W4671)
return
}


}, 1)

tmp15211 := Call(__e, PrimFunc(symshen_4unlocked_2), V4554)


var ifres15090 Obj

if True == tmp15211 {
tmp15091 := MakeNative(func(__e *ControlFlow) {
W4672 := __e.Get(1)
_ = W4672
tmp15208 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4672)
}
__typedArg0 := W4672
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15208 {
tmp15092 := MakeNative(func(__e *ControlFlow) {
W4673 := __e.Get(1)
_ = W4673
tmp15204 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4673, symlambda)
}
__typedArg0 := W4673
__typedArg1 := symlambda
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15204 {
tmp15093 := MakeNative(func(__e *ControlFlow) {
W4674 := __e.Get(1)
_ = W4674
tmp15200 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4674)
}
__typedArg0 := W4674
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15200 {
tmp15094 := MakeNative(func(__e *ControlFlow) {
W4675 := __e.Get(1)
_ = W4675
tmp15095 := MakeNative(func(__e *ControlFlow) {
W4676 := __e.Get(1)
_ = W4676
tmp15195 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4676)
}
__typedArg0 := W4676
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15195 {
tmp15096 := MakeNative(func(__e *ControlFlow) {
W4677 := __e.Get(1)
_ = W4677
tmp15097 := MakeNative(func(__e *ControlFlow) {
W4678 := __e.Get(1)
_ = W4678
tmp15190 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4678, Nil)
}
__typedArg0 := W4678
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15190 {
tmp15098 := MakeNative(func(__e *ControlFlow) {
W4679 := __e.Get(1)
_ = W4679
tmp15099 := MakeNative(func(__e *ControlFlow) {
W4680 := __e.Get(1)
_ = W4680
tmp15166 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4679)
}
__typedArg0 := W4679
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15166 {
tmp15100 := MakeNative(func(__e *ControlFlow) {
W4685 := __e.Get(1)
_ = W4685
tmp15101 := MakeNative(func(__e *ControlFlow) {
W4686 := __e.Get(1)
_ = W4686
tmp15102 := MakeNative(func(__e *ControlFlow) {
W4687 := __e.Get(1)
_ = W4687
tmp15146 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4686)
}
__typedArg0 := W4686
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15146 {
tmp15103 := MakeNative(func(__e *ControlFlow) {
W4689 := __e.Get(1)
_ = W4689
tmp15104 := MakeNative(func(__e *ControlFlow) {
W4690 := __e.Get(1)
_ = W4690
tmp15108 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4689, sym_1_1_6)
}
__typedArg0 := W4689
__typedArg1 := sym_1_1_6
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15108 {
__e.TailApply(PrimFunc(symthaw), W4690)
return
} else {
tmp15106 := Call(__e, PrimFunc(symshen_4pvar_2), W4689)


if True == tmp15106 {
__e.TailApply(PrimFunc(symshen_4bind_b), W4689, sym_1_1_6, V4553, W4690)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp15109 := MakeNative(func(__e *ControlFlow) {
tmp15110 := MakeNative(func(__e *ControlFlow) {
W4691 := __e.Get(1)
_ = W4691
tmp15111 := MakeNative(func(__e *ControlFlow) {
W4692 := __e.Get(1)
_ = W4692
tmp15131 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4691)
}
__typedArg0 := W4691
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15131 {
tmp15112 := MakeNative(func(__e *ControlFlow) {
W4694 := __e.Get(1)
_ = W4694
tmp15113 := MakeNative(func(__e *ControlFlow) {
W4695 := __e.Get(1)
_ = W4695
tmp15114 := MakeNative(func(__e *ControlFlow) {
W4696 := __e.Get(1)
_ = W4696
tmp15118 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4695, Nil)
}
__typedArg0 := W4695
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15118 {
__e.TailApply(PrimFunc(symthaw), W4696)
return
} else {
tmp15116 := Call(__e, PrimFunc(symshen_4pvar_2), W4695)


if True == tmp15116 {
__e.TailApply(PrimFunc(symshen_4bind_b), W4695, Nil, V4553, W4696)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp15119 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W4692, W4694)
return
}, 0)

__e.TailApply(tmp15114, tmp15119)
return


}, 1)

tmp15120 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4691)
}
__typedArg0 := W4691
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15121 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15120, V4553)


__e.TailApply(tmp15113, tmp15121)
return


}, 1)

tmp15122 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4691)
}
__typedArg0 := W4691
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp15112, tmp15122)
return


} else {
tmp15129 := Call(__e, PrimFunc(symshen_4pvar_2), W4691)


if True == tmp15129 {
tmp15123 := MakeNative(func(__e *ControlFlow) {
W4697 := __e.Get(1)
_ = W4697
tmp15124 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W4697, Nil)
}
__typedArg0 := W4697
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15125 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W4692, W4697)
return
}, 0)

tmp15126 := Call(__e, PrimFunc(symshen_4bind_b), W4691, tmp15124, V4553, tmp15125)


__e.TailApply(PrimFunc(symshen_4gc), V4553, tmp15126)
return


}, 1)

tmp15127 := Call(__e, PrimFunc(symshen_4newpv), V4553)


__e.TailApply(tmp15123, tmp15127)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp15132 := MakeNative(func(__e *ControlFlow) {
Z4693 := __e.Get(1)
_ = Z4693
__e.TailApply(W4687, Z4693)
return
}, 1)

__e.TailApply(tmp15111, tmp15132)
return


}, 1)

tmp15133 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4686)
}
__typedArg0 := W4686
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15134 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15133, V4553)


__e.TailApply(tmp15110, tmp15134)
return


}, 0)

__e.TailApply(tmp15104, tmp15109)
return


}, 1)

tmp15135 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4686)
}
__typedArg0 := W4686
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15136 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15135, V4553)


__e.TailApply(tmp15103, tmp15136)
return


} else {
tmp15144 := Call(__e, PrimFunc(symshen_4pvar_2), W4686)


if True == tmp15144 {
tmp15137 := MakeNative(func(__e *ControlFlow) {
W4698 := __e.Get(1)
_ = W4698
tmp15138 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W4698, Nil)
}
__typedArg0 := W4698
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15139 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp15138)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp15138
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15140 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W4687, W4698)
return
}, 0)

tmp15141 := Call(__e, PrimFunc(symshen_4bind_b), W4686, tmp15139, V4553, tmp15140)


__e.TailApply(PrimFunc(symshen_4gc), V4553, tmp15141)
return


}, 1)

tmp15142 := Call(__e, PrimFunc(symshen_4newpv), V4553)


__e.TailApply(tmp15137, tmp15142)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp15147 := MakeNative(func(__e *ControlFlow) {
Z4688 := __e.Get(1)
_ = Z4688
tmp15148 := Call(__e, W4680, W4685)


__e.TailApply(tmp15148, Z4688)
return


}, 1)

__e.TailApply(tmp15102, tmp15147)
return


}, 1)

tmp15149 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4679)
}
__typedArg0 := W4679
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15150 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15149, V4553)


__e.TailApply(tmp15101, tmp15150)
return


}, 1)

tmp15151 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4679)
}
__typedArg0 := W4679
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp15100, tmp15151)
return


} else {
tmp15164 := Call(__e, PrimFunc(symshen_4pvar_2), W4679)


if True == tmp15164 {
tmp15152 := MakeNative(func(__e *ControlFlow) {
W4699 := __e.Get(1)
_ = W4699
tmp15153 := MakeNative(func(__e *ControlFlow) {
W4700 := __e.Get(1)
_ = W4700
tmp15154 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W4700, Nil)
}
__typedArg0 := W4700
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15155 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp15154)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp15154
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15156 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W4699, tmp15155)
}
__typedArg0 := W4699
__typedArg1 := tmp15155
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15157 := MakeNative(func(__e *ControlFlow) {
tmp15158 := Call(__e, W4680, W4699)


__e.TailApply(tmp15158, W4700)
return


}, 0)

tmp15159 := Call(__e, PrimFunc(symshen_4bind_b), W4679, tmp15156, V4553, tmp15157)


__e.TailApply(PrimFunc(symshen_4gc), V4553, tmp15159)
return


}, 1)

tmp15160 := Call(__e, PrimFunc(symshen_4newpv), V4553)


tmp15161 := Call(__e, tmp15153, tmp15160)


__e.TailApply(PrimFunc(symshen_4gc), V4553, tmp15161)
return


}, 1)

tmp15162 := Call(__e, PrimFunc(symshen_4newpv), V4553)


__e.TailApply(tmp15152, tmp15162)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp15167 := MakeNative(func(__e *ControlFlow) {
Z4681 := __e.Get(1)
_ = Z4681
__e.Return(MakeNative(func(__e *ControlFlow) {
Z4682 := __e.Get(1)
_ = Z4682
tmp15168 := MakeNative(func(__e *ControlFlow) {
W4683 := __e.Get(1)
_ = W4683
tmp15169 := MakeNative(func(__e *ControlFlow) {
W4684 := __e.Get(1)
_ = W4684
tmp15170 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15170

tmp15171 := Call(__e, PrimFunc(symshen_4lazyderef), W4675, V4553)


tmp15172 := Call(__e, PrimFunc(symshen_4freshterm), tmp15171)


tmp15173 := MakeNative(func(__e *ControlFlow) {
tmp15174 := Call(__e, PrimFunc(symshen_4lazyderef), W4675, V4553)


tmp15175 := Call(__e, PrimFunc(symshen_4deref), W4684, V4553)


tmp15176 := Call(__e, PrimFunc(symshen_4deref), W4677, V4553)


tmp15177 := Call(__e, PrimFunc(symshen_4beta), tmp15174, tmp15175, tmp15176)


tmp15178 := MakeNative(func(__e *ControlFlow) {
tmp15179 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp15180 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Z4681, Nil)
}
__typedArg0 := Z4681
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15181 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp15179, tmp15180)
}
__typedArg0 := tmp15179
__typedArg1 := tmp15180
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15182 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W4684, tmp15181)
}
__typedArg0 := W4684
__typedArg1 := tmp15181
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15183 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp15182, V4552)
}
__typedArg0 := tmp15182
__typedArg1 := V4552
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4683, Z4682, tmp15183, V4553, V4554, W4557, V4556)
return


}, 0)

__e.TailApply(PrimFunc(symbind), W4683, tmp15177, V4553, V4554, W4557, tmp15178)
return


}, 0)

tmp15184 := Call(__e, PrimFunc(symbind), W4684, tmp15172, V4553, V4554, W4557, tmp15173)


__e.TailApply(PrimFunc(symshen_4gc), V4553, tmp15184)
return


}, 1)

tmp15185 := Call(__e, PrimFunc(symshen_4newpv), V4553)


tmp15186 := Call(__e, tmp15169, tmp15185)


__e.TailApply(PrimFunc(symshen_4gc), V4553, tmp15186)
return


}, 1)

tmp15187 := Call(__e, PrimFunc(symshen_4newpv), V4553)


__e.TailApply(tmp15168, tmp15187)
return


}, 1))
return
}, 1)

__e.TailApply(tmp15099, tmp15167)
return


}, 1)

tmp15188 := Call(__e, PrimFunc(symshen_4lazyderef), V4551, V4553)


__e.TailApply(tmp15098, tmp15188)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15191 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4676)
}
__typedArg0 := W4676
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15192 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15191, V4553)


__e.TailApply(tmp15097, tmp15192)
return


}, 1)

tmp15193 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4676)
}
__typedArg0 := W4676
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp15096, tmp15193)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15196 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4674)
}
__typedArg0 := W4674
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15197 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15196, V4553)


__e.TailApply(tmp15095, tmp15197)
return


}, 1)

tmp15198 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4674)
}
__typedArg0 := W4674
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp15094, tmp15198)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15201 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4672)
}
__typedArg0 := W4672
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15202 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15201, V4553)


__e.TailApply(tmp15093, tmp15202)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15205 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4672)
}
__typedArg0 := W4672
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15206 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15205, V4553)


__e.TailApply(tmp15092, tmp15206)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15209 := Call(__e, PrimFunc(symshen_4lazyderef), V4550, V4553)


tmp15210 := Call(__e, tmp15091, tmp15209)


ifres15090 = tmp15210


} else {
ifres15090 = False


}

__e.TailApply(tmp14790, ifres15090)
return


} else {
__e.Return(W4661)
return
}


}, 1)

tmp15254 := Call(__e, PrimFunc(symshen_4unlocked_2), V4554)


var ifres15214 Obj

if True == tmp15254 {
tmp15215 := MakeNative(func(__e *ControlFlow) {
W4662 := __e.Get(1)
_ = W4662
tmp15251 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4662)
}
__typedArg0 := W4662
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15251 {
tmp15216 := MakeNative(func(__e *ControlFlow) {
W4663 := __e.Get(1)
_ = W4663
tmp15247 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4663, sym_8s)
}
__typedArg0 := W4663
__typedArg1 := sym_8s
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15247 {
tmp15217 := MakeNative(func(__e *ControlFlow) {
W4664 := __e.Get(1)
_ = W4664
tmp15243 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4664)
}
__typedArg0 := W4664
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15243 {
tmp15218 := MakeNative(func(__e *ControlFlow) {
W4665 := __e.Get(1)
_ = W4665
tmp15219 := MakeNative(func(__e *ControlFlow) {
W4666 := __e.Get(1)
_ = W4666
tmp15238 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4666)
}
__typedArg0 := W4666
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15238 {
tmp15220 := MakeNative(func(__e *ControlFlow) {
W4667 := __e.Get(1)
_ = W4667
tmp15221 := MakeNative(func(__e *ControlFlow) {
W4668 := __e.Get(1)
_ = W4668
tmp15233 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4668, Nil)
}
__typedArg0 := W4668
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15233 {
tmp15222 := MakeNative(func(__e *ControlFlow) {
W4669 := __e.Get(1)
_ = W4669
tmp15223 := MakeNative(func(__e *ControlFlow) {
W4670 := __e.Get(1)
_ = W4670
tmp15227 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4669, symstring)
}
__typedArg0 := W4669
__typedArg1 := symstring
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15227 {
__e.TailApply(PrimFunc(symthaw), W4670)
return
} else {
tmp15225 := Call(__e, PrimFunc(symshen_4pvar_2), W4669)


if True == tmp15225 {
__e.TailApply(PrimFunc(symshen_4bind_b), W4669, symstring, V4553, W4670)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp15228 := MakeNative(func(__e *ControlFlow) {
tmp15229 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15229

tmp15230 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4667, symstring, V4552, V4553, V4554, W4557, V4556)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4665, symstring, V4552, V4553, V4554, W4557, tmp15230)
return


}, 0)

__e.TailApply(tmp15223, tmp15228)
return


}, 1)

tmp15231 := Call(__e, PrimFunc(symshen_4lazyderef), V4551, V4553)


__e.TailApply(tmp15222, tmp15231)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15234 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4666)
}
__typedArg0 := W4666
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15235 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15234, V4553)


__e.TailApply(tmp15221, tmp15235)
return


}, 1)

tmp15236 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4666)
}
__typedArg0 := W4666
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp15220, tmp15236)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15239 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4664)
}
__typedArg0 := W4664
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15240 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15239, V4553)


__e.TailApply(tmp15219, tmp15240)
return


}, 1)

tmp15241 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4664)
}
__typedArg0 := W4664
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp15218, tmp15241)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15244 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4662)
}
__typedArg0 := W4662
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15245 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15244, V4553)


__e.TailApply(tmp15217, tmp15245)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15248 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4662)
}
__typedArg0 := W4662
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15249 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15248, V4553)


__e.TailApply(tmp15216, tmp15249)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15252 := Call(__e, PrimFunc(symshen_4lazyderef), V4550, V4553)


tmp15253 := Call(__e, tmp15215, tmp15252)


ifres15214 = tmp15253


} else {
ifres15214 = False


}

__e.TailApply(tmp14789, ifres15214)
return


} else {
__e.Return(W4640)
return
}


}, 1)

tmp15339 := Call(__e, PrimFunc(symshen_4unlocked_2), V4554)


var ifres15257 Obj

if True == tmp15339 {
tmp15258 := MakeNative(func(__e *ControlFlow) {
W4641 := __e.Get(1)
_ = W4641
tmp15336 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4641)
}
__typedArg0 := W4641
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15336 {
tmp15259 := MakeNative(func(__e *ControlFlow) {
W4642 := __e.Get(1)
_ = W4642
tmp15332 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4642, sym_8v)
}
__typedArg0 := W4642
__typedArg1 := sym_8v
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15332 {
tmp15260 := MakeNative(func(__e *ControlFlow) {
W4643 := __e.Get(1)
_ = W4643
tmp15328 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4643)
}
__typedArg0 := W4643
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15328 {
tmp15261 := MakeNative(func(__e *ControlFlow) {
W4644 := __e.Get(1)
_ = W4644
tmp15262 := MakeNative(func(__e *ControlFlow) {
W4645 := __e.Get(1)
_ = W4645
tmp15323 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4645)
}
__typedArg0 := W4645
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15323 {
tmp15263 := MakeNative(func(__e *ControlFlow) {
W4646 := __e.Get(1)
_ = W4646
tmp15264 := MakeNative(func(__e *ControlFlow) {
W4647 := __e.Get(1)
_ = W4647
tmp15318 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4647, Nil)
}
__typedArg0 := W4647
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15318 {
tmp15265 := MakeNative(func(__e *ControlFlow) {
W4648 := __e.Get(1)
_ = W4648
tmp15266 := MakeNative(func(__e *ControlFlow) {
W4649 := __e.Get(1)
_ = W4649
tmp15310 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4648)
}
__typedArg0 := W4648
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15310 {
tmp15267 := MakeNative(func(__e *ControlFlow) {
W4651 := __e.Get(1)
_ = W4651
tmp15268 := MakeNative(func(__e *ControlFlow) {
W4652 := __e.Get(1)
_ = W4652
tmp15272 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4651, symvector)
}
__typedArg0 := W4651
__typedArg1 := symvector
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15272 {
__e.TailApply(PrimFunc(symthaw), W4652)
return
} else {
tmp15270 := Call(__e, PrimFunc(symshen_4pvar_2), W4651)


if True == tmp15270 {
__e.TailApply(PrimFunc(symshen_4bind_b), W4651, symvector, V4553, W4652)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp15273 := MakeNative(func(__e *ControlFlow) {
tmp15274 := MakeNative(func(__e *ControlFlow) {
W4653 := __e.Get(1)
_ = W4653
tmp15275 := MakeNative(func(__e *ControlFlow) {
W4654 := __e.Get(1)
_ = W4654
tmp15295 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4653)
}
__typedArg0 := W4653
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15295 {
tmp15276 := MakeNative(func(__e *ControlFlow) {
W4656 := __e.Get(1)
_ = W4656
tmp15277 := MakeNative(func(__e *ControlFlow) {
W4657 := __e.Get(1)
_ = W4657
tmp15278 := MakeNative(func(__e *ControlFlow) {
W4658 := __e.Get(1)
_ = W4658
tmp15282 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4657, Nil)
}
__typedArg0 := W4657
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15282 {
__e.TailApply(PrimFunc(symthaw), W4658)
return
} else {
tmp15280 := Call(__e, PrimFunc(symshen_4pvar_2), W4657)


if True == tmp15280 {
__e.TailApply(PrimFunc(symshen_4bind_b), W4657, Nil, V4553, W4658)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp15283 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W4654, W4656)
return
}, 0)

__e.TailApply(tmp15278, tmp15283)
return


}, 1)

tmp15284 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4653)
}
__typedArg0 := W4653
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15285 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15284, V4553)


__e.TailApply(tmp15277, tmp15285)
return


}, 1)

tmp15286 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4653)
}
__typedArg0 := W4653
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp15276, tmp15286)
return


} else {
tmp15293 := Call(__e, PrimFunc(symshen_4pvar_2), W4653)


if True == tmp15293 {
tmp15287 := MakeNative(func(__e *ControlFlow) {
W4659 := __e.Get(1)
_ = W4659
tmp15288 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W4659, Nil)
}
__typedArg0 := W4659
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15289 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W4654, W4659)
return
}, 0)

tmp15290 := Call(__e, PrimFunc(symshen_4bind_b), W4653, tmp15288, V4553, tmp15289)


__e.TailApply(PrimFunc(symshen_4gc), V4553, tmp15290)
return


}, 1)

tmp15291 := Call(__e, PrimFunc(symshen_4newpv), V4553)


__e.TailApply(tmp15287, tmp15291)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp15296 := MakeNative(func(__e *ControlFlow) {
Z4655 := __e.Get(1)
_ = Z4655
__e.TailApply(W4649, Z4655)
return
}, 1)

__e.TailApply(tmp15275, tmp15296)
return


}, 1)

tmp15297 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4648)
}
__typedArg0 := W4648
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15298 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15297, V4553)


__e.TailApply(tmp15274, tmp15298)
return


}, 0)

__e.TailApply(tmp15268, tmp15273)
return


}, 1)

tmp15299 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4648)
}
__typedArg0 := W4648
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15300 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15299, V4553)


__e.TailApply(tmp15267, tmp15300)
return


} else {
tmp15308 := Call(__e, PrimFunc(symshen_4pvar_2), W4648)


if True == tmp15308 {
tmp15301 := MakeNative(func(__e *ControlFlow) {
W4660 := __e.Get(1)
_ = W4660
tmp15302 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W4660, Nil)
}
__typedArg0 := W4660
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15303 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp15302)
}
__typedArg0 := symvector
__typedArg1 := tmp15302
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15304 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W4649, W4660)
return
}, 0)

tmp15305 := Call(__e, PrimFunc(symshen_4bind_b), W4648, tmp15303, V4553, tmp15304)


__e.TailApply(PrimFunc(symshen_4gc), V4553, tmp15305)
return


}, 1)

tmp15306 := Call(__e, PrimFunc(symshen_4newpv), V4553)


__e.TailApply(tmp15301, tmp15306)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp15311 := MakeNative(func(__e *ControlFlow) {
Z4650 := __e.Get(1)
_ = Z4650
tmp15312 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15312

tmp15313 := MakeNative(func(__e *ControlFlow) {
tmp15314 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Z4650, Nil)
}
__typedArg0 := Z4650
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15315 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp15314)
}
__typedArg0 := symvector
__typedArg1 := tmp15314
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4646, tmp15315, V4552, V4553, V4554, W4557, V4556)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4644, Z4650, V4552, V4553, V4554, W4557, tmp15313)
return


}, 1)

__e.TailApply(tmp15266, tmp15311)
return


}, 1)

tmp15316 := Call(__e, PrimFunc(symshen_4lazyderef), V4551, V4553)


__e.TailApply(tmp15265, tmp15316)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15319 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4645)
}
__typedArg0 := W4645
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15320 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15319, V4553)


__e.TailApply(tmp15264, tmp15320)
return


}, 1)

tmp15321 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4645)
}
__typedArg0 := W4645
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp15263, tmp15321)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15324 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4643)
}
__typedArg0 := W4643
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15325 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15324, V4553)


__e.TailApply(tmp15262, tmp15325)
return


}, 1)

tmp15326 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4643)
}
__typedArg0 := W4643
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp15261, tmp15326)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15329 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4641)
}
__typedArg0 := W4641
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15330 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15329, V4553)


__e.TailApply(tmp15260, tmp15330)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15333 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4641)
}
__typedArg0 := W4641
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15334 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15333, V4553)


__e.TailApply(tmp15259, tmp15334)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15337 := Call(__e, PrimFunc(symshen_4lazyderef), V4550, V4553)


tmp15338 := Call(__e, tmp15258, tmp15337)


ifres15257 = tmp15338


} else {
ifres15257 = False


}

__e.TailApply(tmp14788, ifres15257)
return


} else {
__e.Return(W4612)
return
}


}, 1)

tmp15445 := Call(__e, PrimFunc(symshen_4unlocked_2), V4554)


var ifres15342 Obj

if True == tmp15445 {
tmp15343 := MakeNative(func(__e *ControlFlow) {
W4613 := __e.Get(1)
_ = W4613
tmp15442 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4613)
}
__typedArg0 := W4613
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15442 {
tmp15344 := MakeNative(func(__e *ControlFlow) {
W4614 := __e.Get(1)
_ = W4614
tmp15438 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4614, sym_8p)
}
__typedArg0 := W4614
__typedArg1 := sym_8p
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15438 {
tmp15345 := MakeNative(func(__e *ControlFlow) {
W4615 := __e.Get(1)
_ = W4615
tmp15434 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4615)
}
__typedArg0 := W4615
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15434 {
tmp15346 := MakeNative(func(__e *ControlFlow) {
W4616 := __e.Get(1)
_ = W4616
tmp15347 := MakeNative(func(__e *ControlFlow) {
W4617 := __e.Get(1)
_ = W4617
tmp15429 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4617)
}
__typedArg0 := W4617
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15429 {
tmp15348 := MakeNative(func(__e *ControlFlow) {
W4618 := __e.Get(1)
_ = W4618
tmp15349 := MakeNative(func(__e *ControlFlow) {
W4619 := __e.Get(1)
_ = W4619
tmp15424 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4619, Nil)
}
__typedArg0 := W4619
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15424 {
tmp15350 := MakeNative(func(__e *ControlFlow) {
W4620 := __e.Get(1)
_ = W4620
tmp15351 := MakeNative(func(__e *ControlFlow) {
W4621 := __e.Get(1)
_ = W4621
tmp15418 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4620)
}
__typedArg0 := W4620
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15418 {
tmp15352 := MakeNative(func(__e *ControlFlow) {
W4624 := __e.Get(1)
_ = W4624
tmp15353 := MakeNative(func(__e *ControlFlow) {
W4625 := __e.Get(1)
_ = W4625
tmp15354 := MakeNative(func(__e *ControlFlow) {
W4626 := __e.Get(1)
_ = W4626
tmp15398 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4625)
}
__typedArg0 := W4625
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15398 {
tmp15355 := MakeNative(func(__e *ControlFlow) {
W4628 := __e.Get(1)
_ = W4628
tmp15356 := MakeNative(func(__e *ControlFlow) {
W4629 := __e.Get(1)
_ = W4629
tmp15360 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4628, sym_d)
}
__typedArg0 := W4628
__typedArg1 := sym_d
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15360 {
__e.TailApply(PrimFunc(symthaw), W4629)
return
} else {
tmp15358 := Call(__e, PrimFunc(symshen_4pvar_2), W4628)


if True == tmp15358 {
__e.TailApply(PrimFunc(symshen_4bind_b), W4628, sym_d, V4553, W4629)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp15361 := MakeNative(func(__e *ControlFlow) {
tmp15362 := MakeNative(func(__e *ControlFlow) {
W4630 := __e.Get(1)
_ = W4630
tmp15363 := MakeNative(func(__e *ControlFlow) {
W4631 := __e.Get(1)
_ = W4631
tmp15383 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4630)
}
__typedArg0 := W4630
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15383 {
tmp15364 := MakeNative(func(__e *ControlFlow) {
W4633 := __e.Get(1)
_ = W4633
tmp15365 := MakeNative(func(__e *ControlFlow) {
W4634 := __e.Get(1)
_ = W4634
tmp15366 := MakeNative(func(__e *ControlFlow) {
W4635 := __e.Get(1)
_ = W4635
tmp15370 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4634, Nil)
}
__typedArg0 := W4634
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15370 {
__e.TailApply(PrimFunc(symthaw), W4635)
return
} else {
tmp15368 := Call(__e, PrimFunc(symshen_4pvar_2), W4634)


if True == tmp15368 {
__e.TailApply(PrimFunc(symshen_4bind_b), W4634, Nil, V4553, W4635)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp15371 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W4631, W4633)
return
}, 0)

__e.TailApply(tmp15366, tmp15371)
return


}, 1)

tmp15372 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4630)
}
__typedArg0 := W4630
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15373 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15372, V4553)


__e.TailApply(tmp15365, tmp15373)
return


}, 1)

tmp15374 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4630)
}
__typedArg0 := W4630
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp15364, tmp15374)
return


} else {
tmp15381 := Call(__e, PrimFunc(symshen_4pvar_2), W4630)


if True == tmp15381 {
tmp15375 := MakeNative(func(__e *ControlFlow) {
W4636 := __e.Get(1)
_ = W4636
tmp15376 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W4636, Nil)
}
__typedArg0 := W4636
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15377 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W4631, W4636)
return
}, 0)

tmp15378 := Call(__e, PrimFunc(symshen_4bind_b), W4630, tmp15376, V4553, tmp15377)


__e.TailApply(PrimFunc(symshen_4gc), V4553, tmp15378)
return


}, 1)

tmp15379 := Call(__e, PrimFunc(symshen_4newpv), V4553)


__e.TailApply(tmp15375, tmp15379)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp15384 := MakeNative(func(__e *ControlFlow) {
Z4632 := __e.Get(1)
_ = Z4632
__e.TailApply(W4626, Z4632)
return
}, 1)

__e.TailApply(tmp15363, tmp15384)
return


}, 1)

tmp15385 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4625)
}
__typedArg0 := W4625
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15386 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15385, V4553)


__e.TailApply(tmp15362, tmp15386)
return


}, 0)

__e.TailApply(tmp15356, tmp15361)
return


}, 1)

tmp15387 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4625)
}
__typedArg0 := W4625
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15388 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15387, V4553)


__e.TailApply(tmp15355, tmp15388)
return


} else {
tmp15396 := Call(__e, PrimFunc(symshen_4pvar_2), W4625)


if True == tmp15396 {
tmp15389 := MakeNative(func(__e *ControlFlow) {
W4637 := __e.Get(1)
_ = W4637
tmp15390 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W4637, Nil)
}
__typedArg0 := W4637
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15391 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_d, tmp15390)
}
__typedArg0 := sym_d
__typedArg1 := tmp15390
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15392 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W4626, W4637)
return
}, 0)

tmp15393 := Call(__e, PrimFunc(symshen_4bind_b), W4625, tmp15391, V4553, tmp15392)


__e.TailApply(PrimFunc(symshen_4gc), V4553, tmp15393)
return


}, 1)

tmp15394 := Call(__e, PrimFunc(symshen_4newpv), V4553)


__e.TailApply(tmp15389, tmp15394)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp15399 := MakeNative(func(__e *ControlFlow) {
Z4627 := __e.Get(1)
_ = Z4627
tmp15400 := Call(__e, W4621, W4624)


__e.TailApply(tmp15400, Z4627)
return


}, 1)

__e.TailApply(tmp15354, tmp15399)
return


}, 1)

tmp15401 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4620)
}
__typedArg0 := W4620
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15402 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15401, V4553)


__e.TailApply(tmp15353, tmp15402)
return


}, 1)

tmp15403 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4620)
}
__typedArg0 := W4620
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp15352, tmp15403)
return


} else {
tmp15416 := Call(__e, PrimFunc(symshen_4pvar_2), W4620)


if True == tmp15416 {
tmp15404 := MakeNative(func(__e *ControlFlow) {
W4638 := __e.Get(1)
_ = W4638
tmp15405 := MakeNative(func(__e *ControlFlow) {
W4639 := __e.Get(1)
_ = W4639
tmp15406 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W4639, Nil)
}
__typedArg0 := W4639
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15407 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_d, tmp15406)
}
__typedArg0 := sym_d
__typedArg1 := tmp15406
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15408 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W4638, tmp15407)
}
__typedArg0 := W4638
__typedArg1 := tmp15407
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15409 := MakeNative(func(__e *ControlFlow) {
tmp15410 := Call(__e, W4621, W4638)


__e.TailApply(tmp15410, W4639)
return


}, 0)

tmp15411 := Call(__e, PrimFunc(symshen_4bind_b), W4620, tmp15408, V4553, tmp15409)


__e.TailApply(PrimFunc(symshen_4gc), V4553, tmp15411)
return


}, 1)

tmp15412 := Call(__e, PrimFunc(symshen_4newpv), V4553)


tmp15413 := Call(__e, tmp15405, tmp15412)


__e.TailApply(PrimFunc(symshen_4gc), V4553, tmp15413)
return


}, 1)

tmp15414 := Call(__e, PrimFunc(symshen_4newpv), V4553)


__e.TailApply(tmp15404, tmp15414)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp15419 := MakeNative(func(__e *ControlFlow) {
Z4622 := __e.Get(1)
_ = Z4622
__e.Return(MakeNative(func(__e *ControlFlow) {
Z4623 := __e.Get(1)
_ = Z4623
tmp15420 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15420

tmp15421 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4618, Z4623, V4552, V4553, V4554, W4557, V4556)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4616, Z4622, V4552, V4553, V4554, W4557, tmp15421)
return


}, 1))
return
}, 1)

__e.TailApply(tmp15351, tmp15419)
return


}, 1)

tmp15422 := Call(__e, PrimFunc(symshen_4lazyderef), V4551, V4553)


__e.TailApply(tmp15350, tmp15422)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15425 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4617)
}
__typedArg0 := W4617
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15426 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15425, V4553)


__e.TailApply(tmp15349, tmp15426)
return


}, 1)

tmp15427 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4617)
}
__typedArg0 := W4617
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp15348, tmp15427)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15430 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4615)
}
__typedArg0 := W4615
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15431 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15430, V4553)


__e.TailApply(tmp15347, tmp15431)
return


}, 1)

tmp15432 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4615)
}
__typedArg0 := W4615
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp15346, tmp15432)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15435 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4613)
}
__typedArg0 := W4613
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15436 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15435, V4553)


__e.TailApply(tmp15345, tmp15436)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15439 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4613)
}
__typedArg0 := W4613
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15440 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15439, V4553)


__e.TailApply(tmp15344, tmp15440)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15443 := Call(__e, PrimFunc(symshen_4lazyderef), V4550, V4553)


tmp15444 := Call(__e, tmp15343, tmp15443)


ifres15342 = tmp15444


} else {
ifres15342 = False


}

__e.TailApply(tmp14787, ifres15342)
return


} else {
__e.Return(W4591)
return
}


}, 1)

tmp15530 := Call(__e, PrimFunc(symshen_4unlocked_2), V4554)


var ifres15448 Obj

if True == tmp15530 {
tmp15449 := MakeNative(func(__e *ControlFlow) {
W4592 := __e.Get(1)
_ = W4592
tmp15527 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4592)
}
__typedArg0 := W4592
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15527 {
tmp15450 := MakeNative(func(__e *ControlFlow) {
W4593 := __e.Get(1)
_ = W4593
tmp15523 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4593, symcons)
}
__typedArg0 := W4593
__typedArg1 := symcons
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15523 {
tmp15451 := MakeNative(func(__e *ControlFlow) {
W4594 := __e.Get(1)
_ = W4594
tmp15519 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4594)
}
__typedArg0 := W4594
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15519 {
tmp15452 := MakeNative(func(__e *ControlFlow) {
W4595 := __e.Get(1)
_ = W4595
tmp15453 := MakeNative(func(__e *ControlFlow) {
W4596 := __e.Get(1)
_ = W4596
tmp15514 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4596)
}
__typedArg0 := W4596
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15514 {
tmp15454 := MakeNative(func(__e *ControlFlow) {
W4597 := __e.Get(1)
_ = W4597
tmp15455 := MakeNative(func(__e *ControlFlow) {
W4598 := __e.Get(1)
_ = W4598
tmp15509 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4598, Nil)
}
__typedArg0 := W4598
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15509 {
tmp15456 := MakeNative(func(__e *ControlFlow) {
W4599 := __e.Get(1)
_ = W4599
tmp15457 := MakeNative(func(__e *ControlFlow) {
W4600 := __e.Get(1)
_ = W4600
tmp15501 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4599)
}
__typedArg0 := W4599
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15501 {
tmp15458 := MakeNative(func(__e *ControlFlow) {
W4602 := __e.Get(1)
_ = W4602
tmp15459 := MakeNative(func(__e *ControlFlow) {
W4603 := __e.Get(1)
_ = W4603
tmp15463 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4602, symlist)
}
__typedArg0 := W4602
__typedArg1 := symlist
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15463 {
__e.TailApply(PrimFunc(symthaw), W4603)
return
} else {
tmp15461 := Call(__e, PrimFunc(symshen_4pvar_2), W4602)


if True == tmp15461 {
__e.TailApply(PrimFunc(symshen_4bind_b), W4602, symlist, V4553, W4603)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp15464 := MakeNative(func(__e *ControlFlow) {
tmp15465 := MakeNative(func(__e *ControlFlow) {
W4604 := __e.Get(1)
_ = W4604
tmp15466 := MakeNative(func(__e *ControlFlow) {
W4605 := __e.Get(1)
_ = W4605
tmp15486 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4604)
}
__typedArg0 := W4604
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15486 {
tmp15467 := MakeNative(func(__e *ControlFlow) {
W4607 := __e.Get(1)
_ = W4607
tmp15468 := MakeNative(func(__e *ControlFlow) {
W4608 := __e.Get(1)
_ = W4608
tmp15469 := MakeNative(func(__e *ControlFlow) {
W4609 := __e.Get(1)
_ = W4609
tmp15473 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4608, Nil)
}
__typedArg0 := W4608
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15473 {
__e.TailApply(PrimFunc(symthaw), W4609)
return
} else {
tmp15471 := Call(__e, PrimFunc(symshen_4pvar_2), W4608)


if True == tmp15471 {
__e.TailApply(PrimFunc(symshen_4bind_b), W4608, Nil, V4553, W4609)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp15474 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W4605, W4607)
return
}, 0)

__e.TailApply(tmp15469, tmp15474)
return


}, 1)

tmp15475 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4604)
}
__typedArg0 := W4604
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15476 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15475, V4553)


__e.TailApply(tmp15468, tmp15476)
return


}, 1)

tmp15477 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4604)
}
__typedArg0 := W4604
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp15467, tmp15477)
return


} else {
tmp15484 := Call(__e, PrimFunc(symshen_4pvar_2), W4604)


if True == tmp15484 {
tmp15478 := MakeNative(func(__e *ControlFlow) {
W4610 := __e.Get(1)
_ = W4610
tmp15479 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W4610, Nil)
}
__typedArg0 := W4610
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15480 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W4605, W4610)
return
}, 0)

tmp15481 := Call(__e, PrimFunc(symshen_4bind_b), W4604, tmp15479, V4553, tmp15480)


__e.TailApply(PrimFunc(symshen_4gc), V4553, tmp15481)
return


}, 1)

tmp15482 := Call(__e, PrimFunc(symshen_4newpv), V4553)


__e.TailApply(tmp15478, tmp15482)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp15487 := MakeNative(func(__e *ControlFlow) {
Z4606 := __e.Get(1)
_ = Z4606
__e.TailApply(W4600, Z4606)
return
}, 1)

__e.TailApply(tmp15466, tmp15487)
return


}, 1)

tmp15488 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4599)
}
__typedArg0 := W4599
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15489 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15488, V4553)


__e.TailApply(tmp15465, tmp15489)
return


}, 0)

__e.TailApply(tmp15459, tmp15464)
return


}, 1)

tmp15490 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4599)
}
__typedArg0 := W4599
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15491 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15490, V4553)


__e.TailApply(tmp15458, tmp15491)
return


} else {
tmp15499 := Call(__e, PrimFunc(symshen_4pvar_2), W4599)


if True == tmp15499 {
tmp15492 := MakeNative(func(__e *ControlFlow) {
W4611 := __e.Get(1)
_ = W4611
tmp15493 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W4611, Nil)
}
__typedArg0 := W4611
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15494 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp15493)
}
__typedArg0 := symlist
__typedArg1 := tmp15493
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15495 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W4600, W4611)
return
}, 0)

tmp15496 := Call(__e, PrimFunc(symshen_4bind_b), W4599, tmp15494, V4553, tmp15495)


__e.TailApply(PrimFunc(symshen_4gc), V4553, tmp15496)
return


}, 1)

tmp15497 := Call(__e, PrimFunc(symshen_4newpv), V4553)


__e.TailApply(tmp15492, tmp15497)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp15502 := MakeNative(func(__e *ControlFlow) {
Z4601 := __e.Get(1)
_ = Z4601
tmp15503 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15503

tmp15504 := MakeNative(func(__e *ControlFlow) {
tmp15505 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Z4601, Nil)
}
__typedArg0 := Z4601
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15506 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp15505)
}
__typedArg0 := symlist
__typedArg1 := tmp15505
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4597, tmp15506, V4552, V4553, V4554, W4557, V4556)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4595, Z4601, V4552, V4553, V4554, W4557, tmp15504)
return


}, 1)

__e.TailApply(tmp15457, tmp15502)
return


}, 1)

tmp15507 := Call(__e, PrimFunc(symshen_4lazyderef), V4551, V4553)


__e.TailApply(tmp15456, tmp15507)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15510 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4596)
}
__typedArg0 := W4596
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15511 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15510, V4553)


__e.TailApply(tmp15455, tmp15511)
return


}, 1)

tmp15512 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4596)
}
__typedArg0 := W4596
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp15454, tmp15512)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15515 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4594)
}
__typedArg0 := W4594
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15516 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15515, V4553)


__e.TailApply(tmp15453, tmp15516)
return


}, 1)

tmp15517 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4594)
}
__typedArg0 := W4594
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp15452, tmp15517)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15520 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4592)
}
__typedArg0 := W4592
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15521 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15520, V4553)


__e.TailApply(tmp15451, tmp15521)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15524 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4592)
}
__typedArg0 := W4592
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15525 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15524, V4553)


__e.TailApply(tmp15450, tmp15525)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15528 := Call(__e, PrimFunc(symshen_4lazyderef), V4550, V4553)


tmp15529 := Call(__e, tmp15449, tmp15528)


ifres15448 = tmp15529


} else {
ifres15448 = False


}

__e.TailApply(tmp14786, ifres15448)
return


} else {
__e.Return(W4584)
return
}


}, 1)

tmp15561 := Call(__e, PrimFunc(symshen_4unlocked_2), V4554)


var ifres15533 Obj

if True == tmp15561 {
tmp15534 := MakeNative(func(__e *ControlFlow) {
W4585 := __e.Get(1)
_ = W4585
tmp15558 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4585)
}
__typedArg0 := W4585
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15558 {
tmp15535 := MakeNative(func(__e *ControlFlow) {
W4586 := __e.Get(1)
_ = W4586
tmp15536 := MakeNative(func(__e *ControlFlow) {
W4587 := __e.Get(1)
_ = W4587
tmp15553 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4587)
}
__typedArg0 := W4587
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15553 {
tmp15537 := MakeNative(func(__e *ControlFlow) {
W4588 := __e.Get(1)
_ = W4588
tmp15538 := MakeNative(func(__e *ControlFlow) {
W4589 := __e.Get(1)
_ = W4589
tmp15548 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4589, Nil)
}
__typedArg0 := W4589
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15548 {
tmp15539 := MakeNative(func(__e *ControlFlow) {
W4590 := __e.Get(1)
_ = W4590
tmp15540 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15540

tmp15541 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V4551, Nil)
}
__typedArg0 := V4551
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15542 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp15541)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp15541
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15543 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W4590, tmp15542)
}
__typedArg0 := W4590
__typedArg1 := tmp15542
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15544 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4588, W4590, V4552, V4553, V4554, W4557, V4556)
return
}, 0)

tmp15545 := Call(__e, PrimFunc(symshen_4system_1S_1h), W4586, tmp15543, V4552, V4553, V4554, W4557, tmp15544)


__e.TailApply(PrimFunc(symshen_4gc), V4553, tmp15545)
return


}, 1)

tmp15546 := Call(__e, PrimFunc(symshen_4newpv), V4553)


__e.TailApply(tmp15539, tmp15546)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15549 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4587)
}
__typedArg0 := W4587
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15550 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15549, V4553)


__e.TailApply(tmp15538, tmp15550)
return


}, 1)

tmp15551 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4587)
}
__typedArg0 := W4587
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp15537, tmp15551)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15554 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4585)
}
__typedArg0 := W4585
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15555 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15554, V4553)


__e.TailApply(tmp15536, tmp15555)
return


}, 1)

tmp15556 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4585)
}
__typedArg0 := W4585
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp15535, tmp15556)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15559 := Call(__e, PrimFunc(symshen_4lazyderef), V4550, V4553)


tmp15560 := Call(__e, tmp15534, tmp15559)


ifres15533 = tmp15560


} else {
ifres15533 = False


}

__e.TailApply(tmp14785, ifres15533)
return


} else {
__e.Return(W4577)
return
}


}, 1)

tmp15596 := Call(__e, PrimFunc(symshen_4unlocked_2), V4554)


var ifres15564 Obj

if True == tmp15596 {
tmp15565 := MakeNative(func(__e *ControlFlow) {
W4578 := __e.Get(1)
_ = W4578
tmp15593 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4578)
}
__typedArg0 := W4578
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15593 {
tmp15566 := MakeNative(func(__e *ControlFlow) {
W4579 := __e.Get(1)
_ = W4579
tmp15567 := MakeNative(func(__e *ControlFlow) {
W4580 := __e.Get(1)
_ = W4580
tmp15588 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4580)
}
__typedArg0 := W4580
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15588 {
tmp15568 := MakeNative(func(__e *ControlFlow) {
W4581 := __e.Get(1)
_ = W4581
tmp15569 := MakeNative(func(__e *ControlFlow) {
W4582 := __e.Get(1)
_ = W4582
tmp15583 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4582, Nil)
}
__typedArg0 := W4582
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15583 {
tmp15570 := MakeNative(func(__e *ControlFlow) {
W4583 := __e.Get(1)
_ = W4583
tmp15571 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15571

tmp15572 := Call(__e, PrimFunc(symshen_4lazyderef), W4579, V4553)


tmp15573 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp15572)
}
__typedArg0 := tmp15572
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

tmp15574 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp15573)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp15573
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

tmp15575 := MakeNative(func(__e *ControlFlow) {
tmp15576 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V4551, Nil)
}
__typedArg0 := V4551
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15577 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp15576)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp15576
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15578 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W4583, tmp15577)
}
__typedArg0 := W4583
__typedArg1 := tmp15577
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15579 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4581, W4583, V4552, V4553, V4554, W4557, V4556)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4lookupsig), W4579, tmp15578, V4553, V4554, W4557, tmp15579)
return


}, 0)

tmp15580 := Call(__e, PrimFunc(symwhen), tmp15574, V4553, V4554, W4557, tmp15575)


__e.TailApply(PrimFunc(symshen_4gc), V4553, tmp15580)
return


}, 1)

tmp15581 := Call(__e, PrimFunc(symshen_4newpv), V4553)


__e.TailApply(tmp15570, tmp15581)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15584 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4580)
}
__typedArg0 := W4580
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15585 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15584, V4553)


__e.TailApply(tmp15569, tmp15585)
return


}, 1)

tmp15586 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4580)
}
__typedArg0 := W4580
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp15568, tmp15586)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15589 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4578)
}
__typedArg0 := W4578
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15590 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15589, V4553)


__e.TailApply(tmp15567, tmp15590)
return


}, 1)

tmp15591 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4578)
}
__typedArg0 := W4578
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp15566, tmp15591)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15594 := Call(__e, PrimFunc(symshen_4lazyderef), V4550, V4553)


tmp15595 := Call(__e, tmp15565, tmp15594)


ifres15564 = tmp15595


} else {
ifres15564 = False


}

__e.TailApply(tmp14784, ifres15564)
return


} else {
__e.Return(W4571)
return
}


}, 1)

tmp15623 := Call(__e, PrimFunc(symshen_4unlocked_2), V4554)


var ifres15599 Obj

if True == tmp15623 {
tmp15600 := MakeNative(func(__e *ControlFlow) {
W4572 := __e.Get(1)
_ = W4572
tmp15620 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4572)
}
__typedArg0 := W4572
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15620 {
tmp15601 := MakeNative(func(__e *ControlFlow) {
W4573 := __e.Get(1)
_ = W4573
tmp15616 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4573, symfn)
}
__typedArg0 := W4573
__typedArg1 := symfn
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15616 {
tmp15602 := MakeNative(func(__e *ControlFlow) {
W4574 := __e.Get(1)
_ = W4574
tmp15612 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4574)
}
__typedArg0 := W4574
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15612 {
tmp15603 := MakeNative(func(__e *ControlFlow) {
W4575 := __e.Get(1)
_ = W4575
tmp15604 := MakeNative(func(__e *ControlFlow) {
W4576 := __e.Get(1)
_ = W4576
tmp15607 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4576, Nil)
}
__typedArg0 := W4576
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15607 {
tmp15605 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15605

__e.TailApply(PrimFunc(symshen_4lookupsig), W4575, V4551, V4553, V4554, W4557, V4556)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15608 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4574)
}
__typedArg0 := W4574
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15609 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15608, V4553)


__e.TailApply(tmp15604, tmp15609)
return


}, 1)

tmp15610 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4574)
}
__typedArg0 := W4574
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp15603, tmp15610)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15613 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4572)
}
__typedArg0 := W4572
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15614 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15613, V4553)


__e.TailApply(tmp15602, tmp15614)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15617 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4572)
}
__typedArg0 := W4572
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15618 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15617, V4553)


__e.TailApply(tmp15601, tmp15618)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15621 := Call(__e, PrimFunc(symshen_4lazyderef), V4550, V4553)


tmp15622 := Call(__e, tmp15600, tmp15621)


ifres15599 = tmp15622


} else {
ifres15599 = False


}

__e.TailApply(tmp14783, ifres15599)
return


} else {
__e.Return(W4565)
return
}


}, 1)

tmp15656 := Call(__e, PrimFunc(symshen_4unlocked_2), V4554)


var ifres15626 Obj

if True == tmp15656 {
tmp15627 := MakeNative(func(__e *ControlFlow) {
W4566 := __e.Get(1)
_ = W4566
tmp15653 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4566)
}
__typedArg0 := W4566
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15653 {
tmp15628 := MakeNative(func(__e *ControlFlow) {
W4567 := __e.Get(1)
_ = W4567
tmp15649 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4567, symfn)
}
__typedArg0 := W4567
__typedArg1 := symfn
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15649 {
tmp15629 := MakeNative(func(__e *ControlFlow) {
W4568 := __e.Get(1)
_ = W4568
tmp15645 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4568)
}
__typedArg0 := W4568
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15645 {
tmp15630 := MakeNative(func(__e *ControlFlow) {
W4569 := __e.Get(1)
_ = W4569
tmp15631 := MakeNative(func(__e *ControlFlow) {
W4570 := __e.Get(1)
_ = W4570
tmp15640 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4570, Nil)
}
__typedArg0 := W4570
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15640 {
tmp15632 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15632

tmp15633 := Call(__e, PrimFunc(symshen_4deref), W4569, V4553)


tmp15634 := Call(__e, PrimFunc(symarity), tmp15633)


tmp15635 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp15634, MakeNumber(0))
}
__typedArg0 := tmp15634
__typedArg1 := MakeNumber(0)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

tmp15636 := MakeNative(func(__e *ControlFlow) {
tmp15637 := MakeNative(func(__e *ControlFlow) {
tmp15638 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W4569, Nil)
}
__typedArg0 := W4569
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4system_1S_1h), tmp15638, V4551, V4552, V4553, V4554, W4557, V4556)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V4553, V4554, W4557, tmp15637)
return


}, 0)

__e.TailApply(PrimFunc(symwhen), tmp15635, V4553, V4554, W4557, tmp15636)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15641 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4568)
}
__typedArg0 := W4568
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15642 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15641, V4553)


__e.TailApply(tmp15631, tmp15642)
return


}, 1)

tmp15643 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4568)
}
__typedArg0 := W4568
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp15630, tmp15643)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15646 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4566)
}
__typedArg0 := W4566
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15647 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15646, V4553)


__e.TailApply(tmp15629, tmp15647)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15650 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4566)
}
__typedArg0 := W4566
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15651 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15650, V4553)


__e.TailApply(tmp15628, tmp15651)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15654 := Call(__e, PrimFunc(symshen_4lazyderef), V4550, V4553)


tmp15655 := Call(__e, tmp15627, tmp15654)


ifres15626 = tmp15655


} else {
ifres15626 = False


}

__e.TailApply(tmp14782, ifres15626)
return


} else {
__e.Return(W4561)
return
}


}, 1)

tmp15675 := Call(__e, PrimFunc(symshen_4unlocked_2), V4554)


var ifres15659 Obj

if True == tmp15675 {
tmp15660 := MakeNative(func(__e *ControlFlow) {
W4562 := __e.Get(1)
_ = W4562
tmp15672 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4562)
}
__typedArg0 := W4562
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15672 {
tmp15661 := MakeNative(func(__e *ControlFlow) {
W4563 := __e.Get(1)
_ = W4563
tmp15662 := MakeNative(func(__e *ControlFlow) {
W4564 := __e.Get(1)
_ = W4564
tmp15667 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4564, Nil)
}
__typedArg0 := W4564
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15667 {
tmp15663 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15663

tmp15664 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V4551, Nil)
}
__typedArg0 := V4551
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15665 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp15664)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp15664
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4lookupsig), W4563, tmp15665, V4553, V4554, W4557, V4556)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15668 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4562)
}
__typedArg0 := W4562
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15669 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15668, V4553)


__e.TailApply(tmp15662, tmp15669)
return


}, 1)

tmp15670 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4562)
}
__typedArg0 := W4562
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp15661, tmp15670)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15673 := Call(__e, PrimFunc(symshen_4lazyderef), V4550, V4553)


tmp15674 := Call(__e, tmp15660, tmp15673)


ifres15659 = tmp15674


} else {
ifres15659 = False


}

__e.TailApply(tmp14781, ifres15659)
return


} else {
__e.Return(W4560)
return
}


}, 1)

tmp15681 := Call(__e, PrimFunc(symshen_4unlocked_2), V4554)


var ifres15678 Obj

if True == tmp15681 {
tmp15679 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15679

tmp15680 := Call(__e, PrimFunc(symshen_4by_1hypothesis), V4550, V4551, V4552, V4553, V4554, W4557, V4556)


ifres15678 = tmp15680


} else {
ifres15678 = False


}

__e.TailApply(tmp14780, ifres15678)
return


} else {
__e.Return(W4559)
return
}


}, 1)

tmp15691 := Call(__e, PrimFunc(symshen_4unlocked_2), V4554)


var ifres15684 Obj

if True == tmp15691 {
tmp15685 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15685

tmp15686 := Call(__e, PrimFunc(symshen_4lazyderef), V4550, V4553)


tmp15687 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp15686)
}
__typedArg0 := tmp15686
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

tmp15688 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp15687)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp15687
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

tmp15689 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4primitive), V4550, V4551, V4553, V4554, W4557, V4556)
return
}, 0)

tmp15690 := Call(__e, PrimFunc(symwhen), tmp15688, V4553, V4554, W4557, tmp15689)


ifres15684 = tmp15690


} else {
ifres15684 = False


}

__e.TailApply(tmp14779, ifres15684)
return


} else {
__e.Return(W4558)
return
}


}, 1)

tmp15703 := Call(__e, PrimFunc(symshen_4unlocked_2), V4554)


var ifres15694 Obj

if True == tmp15703 {
tmp15695 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15695

tmp15696 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dspy_d)
}
__typedArg0 := symshen_4_dspy_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp15697 := MakeNative(func(__e *ControlFlow) {
tmp15698 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp15699 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V4551, Nil)
}
__typedArg0 := V4551
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15700 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp15698, tmp15699)
}
__typedArg0 := tmp15698
__typedArg1 := tmp15699
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15701 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V4550, tmp15700)
}
__typedArg0 := V4550
__typedArg1 := tmp15700
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4show), tmp15701, V4552, V4553, V4554, W4557, V4556)
return


}, 0)

tmp15702 := Call(__e, PrimFunc(symwhen), tmp15696, V4553, V4554, W4557, tmp15697)


ifres15694 = tmp15702


} else {
ifres15694 = False


}

__e.TailApply(tmp14778, ifres15694)
return


}, 1)

__e.TailApply(tmp14777, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V4555)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V4555
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())
return


}, 7)

tmp15705 := Call(__e, ns2_1set, symshen_4system_1S_1h, tmp14776)


_ = tmp15705

tmp15706 := MakeNative(func(__e *ControlFlow) {
V4762 := __e.Get(1)
_ = V4762
tmp15740 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4762)
}
__typedArg0 := V4762
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15721 Obj

if True == tmp15740 {
tmp15738 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4762)
}
__typedArg0 := V4762
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15739 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcons, tmp15738)
}
__typedArg0 := symcons
__typedArg1 := tmp15738
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15723 Obj

if True == tmp15739 {
tmp15736 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4762)
}
__typedArg0 := V4762
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15737 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp15736)
}
__typedArg0 := tmp15736
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15725 Obj

if True == tmp15737 {
tmp15733 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4762)
}
__typedArg0 := V4762
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15734 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp15733)
}
__typedArg0 := tmp15733
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15735 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp15734)
}
__typedArg0 := tmp15734
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15727 Obj

if True == tmp15735 {
tmp15729 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4762)
}
__typedArg0 := V4762
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15730 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp15729)
}
__typedArg0 := tmp15729
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15731 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp15730)
}
__typedArg0 := tmp15730
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15732 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp15731)
}
__typedArg0 := Nil
__typedArg1 := tmp15731
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15728 Obj

if True == tmp15732 {
ifres15728 = True


} else {
ifres15728 = False


}

ifres15727 = ifres15728


} else {
ifres15727 = False


}

var ifres15726 Obj

if True == ifres15727 {
ifres15726 = True


} else {
ifres15726 = False


}

ifres15725 = ifres15726


} else {
ifres15725 = False


}

var ifres15724 Obj

if True == ifres15725 {
ifres15724 = True


} else {
ifres15724 = False


}

ifres15723 = ifres15724


} else {
ifres15723 = False


}

var ifres15722 Obj

if True == ifres15723 {
ifres15722 = True


} else {
ifres15722 = False


}

ifres15721 = ifres15722


} else {
ifres15721 = False


}

if True == ifres15721 {
tmp15707 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4762)
}
__typedArg0 := V4762
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15708 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp15707)
}
__typedArg0 := tmp15707
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15709 := Call(__e, PrimFunc(symshen_4rdecons), tmp15708)


tmp15710 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4762)
}
__typedArg0 := V4762
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15711 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp15710)
}
__typedArg0 := tmp15710
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15712 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp15711)
}
__typedArg0 := tmp15711
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15713 := Call(__e, PrimFunc(symshen_4rdecons), tmp15712)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp15709, tmp15713)
}
__typedArg0 := tmp15709
__typedArg1 := tmp15713
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp15719 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4762)
}
__typedArg0 := V4762
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15719 {
tmp15714 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4762)
}
__typedArg0 := V4762
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15715 := Call(__e, PrimFunc(symshen_4rdecons), tmp15714)


tmp15716 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4762)
}
__typedArg0 := V4762
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15717 := Call(__e, PrimFunc(symshen_4rdecons), tmp15716)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp15715, tmp15717)
}
__typedArg0 := tmp15715
__typedArg1 := tmp15717
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return(V4762)
return
}


}


}, 1)

tmp15741 := Call(__e, ns2_1set, symshen_4rdecons, tmp15706)


_ = tmp15741

tmp15742 := MakeNative(func(__e *ControlFlow) {
V4763 := __e.Get(1)
_ = V4763
V4764 := __e.Get(2)
_ = V4764
V4765 := __e.Get(3)
_ = V4765
V4766 := __e.Get(4)
_ = V4766
V4767 := __e.Get(5)
_ = V4767
V4768 := __e.Get(6)
_ = V4768
tmp15743 := MakeNative(func(__e *ControlFlow) {
W4769 := __e.Get(1)
_ = W4769
tmp15851 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4769, False)
}
__typedArg0 := W4769
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15851 {
tmp15744 := MakeNative(func(__e *ControlFlow) {
W4772 := __e.Get(1)
_ = W4772
tmp15835 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4772, False)
}
__typedArg0 := W4772
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15835 {
tmp15745 := MakeNative(func(__e *ControlFlow) {
W4775 := __e.Get(1)
_ = W4775
tmp15819 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4775, False)
}
__typedArg0 := W4775
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15819 {
tmp15746 := MakeNative(func(__e *ControlFlow) {
W4778 := __e.Get(1)
_ = W4778
tmp15803 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4778, False)
}
__typedArg0 := W4778
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15803 {
tmp15801 := Call(__e, PrimFunc(symshen_4unlocked_2), V4766)


if True == tmp15801 {
tmp15747 := MakeNative(func(__e *ControlFlow) {
W4781 := __e.Get(1)
_ = W4781
tmp15798 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4781, Nil)
}
__typedArg0 := W4781
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15798 {
tmp15748 := MakeNative(func(__e *ControlFlow) {
W4782 := __e.Get(1)
_ = W4782
tmp15749 := MakeNative(func(__e *ControlFlow) {
W4783 := __e.Get(1)
_ = W4783
tmp15793 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4782)
}
__typedArg0 := W4782
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15793 {
tmp15750 := MakeNative(func(__e *ControlFlow) {
W4785 := __e.Get(1)
_ = W4785
tmp15751 := MakeNative(func(__e *ControlFlow) {
W4786 := __e.Get(1)
_ = W4786
tmp15755 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4785, symlist)
}
__typedArg0 := W4785
__typedArg1 := symlist
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15755 {
__e.TailApply(PrimFunc(symthaw), W4786)
return
} else {
tmp15753 := Call(__e, PrimFunc(symshen_4pvar_2), W4785)


if True == tmp15753 {
__e.TailApply(PrimFunc(symshen_4bind_b), W4785, symlist, V4765, W4786)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp15756 := MakeNative(func(__e *ControlFlow) {
tmp15757 := MakeNative(func(__e *ControlFlow) {
W4787 := __e.Get(1)
_ = W4787
tmp15758 := MakeNative(func(__e *ControlFlow) {
W4788 := __e.Get(1)
_ = W4788
tmp15778 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4787)
}
__typedArg0 := W4787
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15778 {
tmp15759 := MakeNative(func(__e *ControlFlow) {
W4790 := __e.Get(1)
_ = W4790
tmp15760 := MakeNative(func(__e *ControlFlow) {
W4791 := __e.Get(1)
_ = W4791
tmp15761 := MakeNative(func(__e *ControlFlow) {
W4792 := __e.Get(1)
_ = W4792
tmp15765 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4791, Nil)
}
__typedArg0 := W4791
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15765 {
__e.TailApply(PrimFunc(symthaw), W4792)
return
} else {
tmp15763 := Call(__e, PrimFunc(symshen_4pvar_2), W4791)


if True == tmp15763 {
__e.TailApply(PrimFunc(symshen_4bind_b), W4791, Nil, V4765, W4792)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp15766 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W4788, W4790)
return
}, 0)

__e.TailApply(tmp15761, tmp15766)
return


}, 1)

tmp15767 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4787)
}
__typedArg0 := W4787
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15768 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15767, V4765)


__e.TailApply(tmp15760, tmp15768)
return


}, 1)

tmp15769 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4787)
}
__typedArg0 := W4787
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp15759, tmp15769)
return


} else {
tmp15776 := Call(__e, PrimFunc(symshen_4pvar_2), W4787)


if True == tmp15776 {
tmp15770 := MakeNative(func(__e *ControlFlow) {
W4793 := __e.Get(1)
_ = W4793
tmp15771 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W4793, Nil)
}
__typedArg0 := W4793
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15772 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W4788, W4793)
return
}, 0)

tmp15773 := Call(__e, PrimFunc(symshen_4bind_b), W4787, tmp15771, V4765, tmp15772)


__e.TailApply(PrimFunc(symshen_4gc), V4765, tmp15773)
return


}, 1)

tmp15774 := Call(__e, PrimFunc(symshen_4newpv), V4765)


__e.TailApply(tmp15770, tmp15774)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp15779 := MakeNative(func(__e *ControlFlow) {
Z4789 := __e.Get(1)
_ = Z4789
__e.TailApply(W4783, Z4789)
return
}, 1)

__e.TailApply(tmp15758, tmp15779)
return


}, 1)

tmp15780 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4782)
}
__typedArg0 := W4782
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15781 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15780, V4765)


__e.TailApply(tmp15757, tmp15781)
return


}, 0)

__e.TailApply(tmp15751, tmp15756)
return


}, 1)

tmp15782 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4782)
}
__typedArg0 := W4782
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15783 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15782, V4765)


__e.TailApply(tmp15750, tmp15783)
return


} else {
tmp15791 := Call(__e, PrimFunc(symshen_4pvar_2), W4782)


if True == tmp15791 {
tmp15784 := MakeNative(func(__e *ControlFlow) {
W4794 := __e.Get(1)
_ = W4794
tmp15785 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W4794, Nil)
}
__typedArg0 := W4794
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15786 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp15785)
}
__typedArg0 := symlist
__typedArg1 := tmp15785
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15787 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W4783, W4794)
return
}, 0)

tmp15788 := Call(__e, PrimFunc(symshen_4bind_b), W4782, tmp15786, V4765, tmp15787)


__e.TailApply(PrimFunc(symshen_4gc), V4765, tmp15788)
return


}, 1)

tmp15789 := Call(__e, PrimFunc(symshen_4newpv), V4765)


__e.TailApply(tmp15784, tmp15789)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp15794 := MakeNative(func(__e *ControlFlow) {
Z4784 := __e.Get(1)
_ = Z4784
tmp15795 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15795

__e.TailApply(PrimFunc(symthaw), V4768)
return


}, 1)

__e.TailApply(tmp15749, tmp15794)
return


}, 1)

tmp15796 := Call(__e, PrimFunc(symshen_4lazyderef), V4764, V4765)


__e.TailApply(tmp15748, tmp15796)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15799 := Call(__e, PrimFunc(symshen_4lazyderef), V4763, V4765)


__e.TailApply(tmp15747, tmp15799)
return


} else {
__e.Return(False)
return
}


} else {
__e.Return(W4778)
return
}


}, 1)

tmp15817 := Call(__e, PrimFunc(symshen_4unlocked_2), V4766)


var ifres15804 Obj

if True == tmp15817 {
tmp15805 := MakeNative(func(__e *ControlFlow) {
W4779 := __e.Get(1)
_ = W4779
tmp15806 := MakeNative(func(__e *ControlFlow) {
W4780 := __e.Get(1)
_ = W4780
tmp15810 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4779, symsymbol)
}
__typedArg0 := W4779
__typedArg1 := symsymbol
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15810 {
__e.TailApply(PrimFunc(symthaw), W4780)
return
} else {
tmp15808 := Call(__e, PrimFunc(symshen_4pvar_2), W4779)


if True == tmp15808 {
__e.TailApply(PrimFunc(symshen_4bind_b), W4779, symsymbol, V4765, W4780)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp15811 := MakeNative(func(__e *ControlFlow) {
tmp15812 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15812

tmp15813 := Call(__e, PrimFunc(symshen_4lazyderef), V4763, V4765)


tmp15814 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsymbol_2) {
return PrimIsSymbol(tmp15813)
}
__typedArg0 := tmp15813
return Call(__e, PrimFunc(symsymbol_2), __typedArg0)
})()

__e.TailApply(PrimFunc(symwhen), tmp15814, V4765, V4766, V4767, V4768)
return


}, 0)

__e.TailApply(tmp15806, tmp15811)
return


}, 1)

tmp15815 := Call(__e, PrimFunc(symshen_4lazyderef), V4764, V4765)


tmp15816 := Call(__e, tmp15805, tmp15815)


ifres15804 = tmp15816


} else {
ifres15804 = False


}

__e.TailApply(tmp15746, ifres15804)
return


} else {
__e.Return(W4775)
return
}


}, 1)

tmp15833 := Call(__e, PrimFunc(symshen_4unlocked_2), V4766)


var ifres15820 Obj

if True == tmp15833 {
tmp15821 := MakeNative(func(__e *ControlFlow) {
W4776 := __e.Get(1)
_ = W4776
tmp15822 := MakeNative(func(__e *ControlFlow) {
W4777 := __e.Get(1)
_ = W4777
tmp15826 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4776, symstring)
}
__typedArg0 := W4776
__typedArg1 := symstring
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15826 {
__e.TailApply(PrimFunc(symthaw), W4777)
return
} else {
tmp15824 := Call(__e, PrimFunc(symshen_4pvar_2), W4776)


if True == tmp15824 {
__e.TailApply(PrimFunc(symshen_4bind_b), W4776, symstring, V4765, W4777)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp15827 := MakeNative(func(__e *ControlFlow) {
tmp15828 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15828

tmp15829 := Call(__e, PrimFunc(symshen_4lazyderef), V4763, V4765)


tmp15830 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_2) {
return PrimIsString(tmp15829)
}
__typedArg0 := tmp15829
return Call(__e, PrimFunc(symstring_2), __typedArg0)
})()

__e.TailApply(PrimFunc(symwhen), tmp15830, V4765, V4766, V4767, V4768)
return


}, 0)

__e.TailApply(tmp15822, tmp15827)
return


}, 1)

tmp15831 := Call(__e, PrimFunc(symshen_4lazyderef), V4764, V4765)


tmp15832 := Call(__e, tmp15821, tmp15831)


ifres15820 = tmp15832


} else {
ifres15820 = False


}

__e.TailApply(tmp15745, ifres15820)
return


} else {
__e.Return(W4772)
return
}


}, 1)

tmp15849 := Call(__e, PrimFunc(symshen_4unlocked_2), V4766)


var ifres15836 Obj

if True == tmp15849 {
tmp15837 := MakeNative(func(__e *ControlFlow) {
W4773 := __e.Get(1)
_ = W4773
tmp15838 := MakeNative(func(__e *ControlFlow) {
W4774 := __e.Get(1)
_ = W4774
tmp15842 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4773, symboolean)
}
__typedArg0 := W4773
__typedArg1 := symboolean
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15842 {
__e.TailApply(PrimFunc(symthaw), W4774)
return
} else {
tmp15840 := Call(__e, PrimFunc(symshen_4pvar_2), W4773)


if True == tmp15840 {
__e.TailApply(PrimFunc(symshen_4bind_b), W4773, symboolean, V4765, W4774)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp15843 := MakeNative(func(__e *ControlFlow) {
tmp15844 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15844

tmp15845 := Call(__e, PrimFunc(symshen_4lazyderef), V4763, V4765)


tmp15846 := Call(__e, PrimFunc(symboolean_2), tmp15845)


__e.TailApply(PrimFunc(symwhen), tmp15846, V4765, V4766, V4767, V4768)
return


}, 0)

__e.TailApply(tmp15838, tmp15843)
return


}, 1)

tmp15847 := Call(__e, PrimFunc(symshen_4lazyderef), V4764, V4765)


tmp15848 := Call(__e, tmp15837, tmp15847)


ifres15836 = tmp15848


} else {
ifres15836 = False


}

__e.TailApply(tmp15744, ifres15836)
return


} else {
__e.Return(W4769)
return
}


}, 1)

tmp15865 := Call(__e, PrimFunc(symshen_4unlocked_2), V4766)


var ifres15852 Obj

if True == tmp15865 {
tmp15853 := MakeNative(func(__e *ControlFlow) {
W4770 := __e.Get(1)
_ = W4770
tmp15854 := MakeNative(func(__e *ControlFlow) {
W4771 := __e.Get(1)
_ = W4771
tmp15858 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4770, symnumber)
}
__typedArg0 := W4770
__typedArg1 := symnumber
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15858 {
__e.TailApply(PrimFunc(symthaw), W4771)
return
} else {
tmp15856 := Call(__e, PrimFunc(symshen_4pvar_2), W4770)


if True == tmp15856 {
__e.TailApply(PrimFunc(symshen_4bind_b), W4770, symnumber, V4765, W4771)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp15859 := MakeNative(func(__e *ControlFlow) {
tmp15860 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15860

tmp15861 := Call(__e, PrimFunc(symshen_4lazyderef), V4763, V4765)


tmp15862 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnumber_2) {
return PrimIsNumber(tmp15861)
}
__typedArg0 := tmp15861
return Call(__e, PrimFunc(symnumber_2), __typedArg0)
})()

__e.TailApply(PrimFunc(symwhen), tmp15862, V4765, V4766, V4767, V4768)
return


}, 0)

__e.TailApply(tmp15854, tmp15859)
return


}, 1)

tmp15863 := Call(__e, PrimFunc(symshen_4lazyderef), V4764, V4765)


tmp15864 := Call(__e, tmp15853, tmp15863)


ifres15852 = tmp15864


} else {
ifres15852 = False


}

__e.TailApply(tmp15743, ifres15852)
return


}, 6)

tmp15866 := Call(__e, ns2_1set, symshen_4primitive, tmp15742)


_ = tmp15866

tmp15867 := MakeNative(func(__e *ControlFlow) {
V4795 := __e.Get(1)
_ = V4795
V4796 := __e.Get(2)
_ = V4796
V4797 := __e.Get(3)
_ = V4797
V4798 := __e.Get(4)
_ = V4798
V4799 := __e.Get(5)
_ = V4799
V4800 := __e.Get(6)
_ = V4800
V4801 := __e.Get(7)
_ = V4801
tmp15868 := MakeNative(func(__e *ControlFlow) {
W4802 := __e.Get(1)
_ = W4802
tmp15879 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4802, False)
}
__typedArg0 := W4802
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15879 {
tmp15877 := Call(__e, PrimFunc(symshen_4unlocked_2), V4799)


if True == tmp15877 {
tmp15869 := MakeNative(func(__e *ControlFlow) {
W4811 := __e.Get(1)
_ = W4811
tmp15874 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4811)
}
__typedArg0 := W4811
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15874 {
tmp15870 := MakeNative(func(__e *ControlFlow) {
W4812 := __e.Get(1)
_ = W4812
tmp15871 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15871

__e.TailApply(PrimFunc(symshen_4by_1hypothesis), V4795, V4796, W4812, V4798, V4799, V4800, V4801)
return


}, 1)

tmp15872 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4811)
}
__typedArg0 := W4811
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(tmp15870, tmp15872)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15875 := Call(__e, PrimFunc(symshen_4lazyderef), V4797, V4798)


__e.TailApply(tmp15869, tmp15875)
return


} else {
__e.Return(False)
return
}


} else {
__e.Return(W4802)
return
}


}, 1)

tmp15921 := Call(__e, PrimFunc(symshen_4unlocked_2), V4799)


var ifres15880 Obj

if True == tmp15921 {
tmp15881 := MakeNative(func(__e *ControlFlow) {
W4803 := __e.Get(1)
_ = W4803
tmp15918 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4803)
}
__typedArg0 := W4803
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15918 {
tmp15882 := MakeNative(func(__e *ControlFlow) {
W4804 := __e.Get(1)
_ = W4804
tmp15914 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4804)
}
__typedArg0 := W4804
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15914 {
tmp15883 := MakeNative(func(__e *ControlFlow) {
W4805 := __e.Get(1)
_ = W4805
tmp15884 := MakeNative(func(__e *ControlFlow) {
W4806 := __e.Get(1)
_ = W4806
tmp15909 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4806)
}
__typedArg0 := W4806
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15909 {
tmp15885 := MakeNative(func(__e *ControlFlow) {
W4807 := __e.Get(1)
_ = W4807
tmp15886 := MakeNative(func(__e *ControlFlow) {
W4808 := __e.Get(1)
_ = W4808
tmp15904 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4808)
}
__typedArg0 := W4808
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15904 {
tmp15887 := MakeNative(func(__e *ControlFlow) {
W4809 := __e.Get(1)
_ = W4809
tmp15888 := MakeNative(func(__e *ControlFlow) {
W4810 := __e.Get(1)
_ = W4810
tmp15899 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4810, Nil)
}
__typedArg0 := W4810
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15899 {
tmp15889 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15889

tmp15890 := Call(__e, PrimFunc(symshen_4deref), W4807, V4798)


tmp15891 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp15892 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp15890, tmp15891)
}
__typedArg0 := tmp15890
__typedArg1 := tmp15891
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

tmp15893 := MakeNative(func(__e *ControlFlow) {
tmp15894 := Call(__e, PrimFunc(symshen_4deref), V4795, V4798)


tmp15895 := Call(__e, PrimFunc(symshen_4deref), W4805, V4798)


tmp15896 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp15894, tmp15895)
}
__typedArg0 := tmp15894
__typedArg1 := tmp15895
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

tmp15897 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symis_b), V4796, W4809, V4798, V4799, V4800, V4801)
return
}, 0)

__e.TailApply(PrimFunc(symwhen), tmp15896, V4798, V4799, V4800, tmp15897)
return


}, 0)

__e.TailApply(PrimFunc(symwhen), tmp15892, V4798, V4799, V4800, tmp15893)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15900 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4808)
}
__typedArg0 := W4808
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15901 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15900, V4798)


__e.TailApply(tmp15888, tmp15901)
return


}, 1)

tmp15902 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4808)
}
__typedArg0 := W4808
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp15887, tmp15902)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15905 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4806)
}
__typedArg0 := W4806
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15906 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15905, V4798)


__e.TailApply(tmp15886, tmp15906)
return


}, 1)

tmp15907 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4806)
}
__typedArg0 := W4806
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp15885, tmp15907)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15910 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4804)
}
__typedArg0 := W4804
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15911 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15910, V4798)


__e.TailApply(tmp15884, tmp15911)
return


}, 1)

tmp15912 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4804)
}
__typedArg0 := W4804
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp15883, tmp15912)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15915 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4803)
}
__typedArg0 := W4803
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15916 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15915, V4798)


__e.TailApply(tmp15882, tmp15916)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15919 := Call(__e, PrimFunc(symshen_4lazyderef), V4797, V4798)


tmp15920 := Call(__e, tmp15881, tmp15919)


ifres15880 = tmp15920


} else {
ifres15880 = False


}

__e.TailApply(tmp15868, ifres15880)
return


}, 7)

tmp15922 := Call(__e, ns2_1set, symshen_4by_1hypothesis, tmp15867)


_ = tmp15922

tmp15923 := MakeNative(func(__e *ControlFlow) {
V4813 := __e.Get(1)
_ = V4813
V4814 := __e.Get(2)
_ = V4814
V4815 := __e.Get(3)
_ = V4815
V4816 := __e.Get(4)
_ = V4816
V4817 := __e.Get(5)
_ = V4817
V4818 := __e.Get(6)
_ = V4818
tmp15928 := Call(__e, PrimFunc(symshen_4unlocked_2), V4816)


if True == tmp15928 {
tmp15924 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15924

tmp15925 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp15926 := Call(__e, PrimFunc(symassoc), V4813, tmp15925)


__e.TailApply(PrimFunc(symshen_4sigf), tmp15926, V4814, V4815, V4816, V4817, V4818)
return


} else {
__e.Return(False)
return
}


}, 6)

tmp15929 := Call(__e, ns2_1set, symshen_4lookupsig, tmp15923)


_ = tmp15929

tmp15930 := MakeNative(func(__e *ControlFlow) {
V4833 := __e.Get(1)
_ = V4833
V4834 := __e.Get(2)
_ = V4834
V4835 := __e.Get(3)
_ = V4835
V4836 := __e.Get(4)
_ = V4836
V4837 := __e.Get(5)
_ = V4837
V4838 := __e.Get(6)
_ = V4838
tmp15937 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4833)
}
__typedArg0 := V4833
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15937 {
tmp15931 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4833)
}
__typedArg0 := V4833
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15932 := Call(__e, tmp15931, V4834)


tmp15933 := Call(__e, tmp15932, V4835)


tmp15934 := Call(__e, tmp15933, V4836)


tmp15935 := Call(__e, tmp15934, V4837)


__e.TailApply(tmp15935, V4838)
return


} else {
__e.Return(False)
return
}


}, 6)

tmp15938 := Call(__e, ns2_1set, symshen_4sigf, tmp15930)


_ = tmp15938

tmp15939 := MakeNative(func(__e *ControlFlow) {
V4839 := __e.Get(1)
_ = V4839
tmp15940 := MakeNative(func(__e *ControlFlow) {
W4840 := __e.Get(1)
_ = W4840
tmp15941 := MakeNative(func(__e *ControlFlow) {
W4841 := __e.Get(1)
_ = W4841
tmp15942 := MakeNative(func(__e *ControlFlow) {
W4842 := __e.Get(1)
_ = W4842
tmp15943 := MakeNative(func(__e *ControlFlow) {
W4843 := __e.Get(1)
_ = W4843
__e.Return(W4843)
return
}, 1)

tmp15944 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dgensym_d)
}
__typedArg0 := symshen_4_dgensym_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp15946 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dgensym_d, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(1))
__typedN1, __typedOK1 := TypedFloat64(tmp15944)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp15944
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())
}
__typedArg0 := symshen_4_dgensym_d
__typedArg1 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(1))
__typedN1, __typedOK1 := TypedFloat64(tmp15944)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp15944
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

tmp15947 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(W4842, MakeNumber(2), tmp15946)
}
__typedArg0 := W4842
__typedArg1 := MakeNumber(2)
__typedArg2 := tmp15946
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

__e.TailApply(tmp15943, tmp15947)
return


}, 1)

tmp15948 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(W4841, MakeNumber(1), V4839)
}
__typedArg0 := W4841
__typedArg1 := MakeNumber(1)
__typedArg2 := V4839
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

__e.TailApply(tmp15942, tmp15948)
return


}, 1)

tmp15949 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(W4840, MakeNumber(0), symshen_4print_1freshterm)
}
__typedArg0 := W4840
__typedArg1 := MakeNumber(0)
__typedArg2 := symshen_4print_1freshterm
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

__e.TailApply(tmp15941, tmp15949)
return


}, 1)

tmp15950 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symabsvector) {
return PrimAbsvector(MakeNumber(3))
}
__typedArg0 := MakeNumber(3)
return Call(__e, PrimFunc(symabsvector), __typedArg0)
})()

__e.TailApply(tmp15940, tmp15950)
return


}, 1)

tmp15951 := Call(__e, ns2_1set, symshen_4freshterm, tmp15939)


_ = tmp15951

tmp15952 := MakeNative(func(__e *ControlFlow) {
V4844 := __e.Get(1)
_ = V4844
tmp15953 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V4844, MakeNumber(1))
}
__typedArg0 := V4844
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})()

tmp15954 := Call(__e, PrimFunc(symshen_4app), tmp15953, MakeString(""), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("&&"))
__typedS1, __typedOK1 := TypedString(tmp15954)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("&&")
__typedArg1 := tmp15954
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp15955 := Call(__e, ns2_1set, symshen_4print_1freshterm, tmp15952)


_ = tmp15955

tmp15956 := MakeNative(func(__e *ControlFlow) {
V4845 := __e.Get(1)
_ = V4845
V4846 := __e.Get(2)
_ = V4846
V4847 := __e.Get(3)
_ = V4847
V4848 := __e.Get(4)
_ = V4848
V4849 := __e.Get(5)
_ = V4849
V4850 := __e.Get(6)
_ = V4850
V4851 := __e.Get(7)
_ = V4851
tmp15957 := MakeNative(func(__e *ControlFlow) {
W4852 := __e.Get(1)
_ = W4852
tmp15968 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4852, False)
}
__typedArg0 := W4852
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15968 {
tmp15966 := Call(__e, PrimFunc(symshen_4unlocked_2), V4849)


if True == tmp15966 {
tmp15958 := MakeNative(func(__e *ControlFlow) {
W4856 := __e.Get(1)
_ = W4856
tmp15963 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4856)
}
__typedArg0 := W4856
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15963 {
tmp15959 := MakeNative(func(__e *ControlFlow) {
W4857 := __e.Get(1)
_ = W4857
tmp15960 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15960

__e.TailApply(PrimFunc(symshen_4search_1user_1datatypes), V4845, V4846, W4857, V4848, V4849, V4850, V4851)
return


}, 1)

tmp15961 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4856)
}
__typedArg0 := W4856
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(tmp15959, tmp15961)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15964 := Call(__e, PrimFunc(symshen_4lazyderef), V4847, V4848)


__e.TailApply(tmp15958, tmp15964)
return


} else {
__e.Return(False)
return
}


} else {
__e.Return(W4852)
return
}


}, 1)

tmp15988 := Call(__e, PrimFunc(symshen_4unlocked_2), V4849)


var ifres15969 Obj

if True == tmp15988 {
tmp15970 := MakeNative(func(__e *ControlFlow) {
W4853 := __e.Get(1)
_ = W4853
tmp15985 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4853)
}
__typedArg0 := W4853
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15985 {
tmp15971 := MakeNative(func(__e *ControlFlow) {
W4854 := __e.Get(1)
_ = W4854
tmp15981 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4854)
}
__typedArg0 := W4854
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15981 {
tmp15972 := MakeNative(func(__e *ControlFlow) {
W4855 := __e.Get(1)
_ = W4855
tmp15973 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15973

tmp15974 := Call(__e, PrimFunc(symshen_4deref), W4855, V4848)


tmp15975 := Call(__e, PrimFunc(symshen_4deref), V4845, V4848)


tmp15976 := Call(__e, tmp15974, tmp15975)


tmp15977 := Call(__e, PrimFunc(symshen_4deref), V4846, V4848)


tmp15978 := Call(__e, tmp15976, tmp15977)


__e.TailApply(PrimFunc(symcall), tmp15978, V4848, V4849, V4850, V4851)
return


}, 1)

tmp15979 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4854)
}
__typedArg0 := W4854
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(tmp15972, tmp15979)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15982 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4853)
}
__typedArg0 := W4853
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15983 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15982, V4848)


__e.TailApply(tmp15971, tmp15983)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15986 := Call(__e, PrimFunc(symshen_4lazyderef), V4847, V4848)


tmp15987 := Call(__e, tmp15970, tmp15986)


ifres15969 = tmp15987


} else {
ifres15969 = False


}

__e.TailApply(tmp15957, ifres15969)
return


}, 7)

tmp15989 := Call(__e, ns2_1set, symshen_4search_1user_1datatypes, tmp15956)


_ = tmp15989

tmp15990 := MakeNative(func(__e *ControlFlow) {
V4858 := __e.Get(1)
_ = V4858
V4859 := __e.Get(2)
_ = V4859
V4860 := __e.Get(3)
_ = V4860
V4861 := __e.Get(4)
_ = V4861
V4862 := __e.Get(5)
_ = V4862
V4863 := __e.Get(6)
_ = V4863
V4864 := __e.Get(7)
_ = V4864
tmp15991 := MakeNative(func(__e *ControlFlow) {
W4865 := __e.Get(1)
_ = W4865
tmp15992 := MakeNative(func(__e *ControlFlow) {
W4866 := __e.Get(1)
_ = W4866
tmp16422 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4866, False)
}
__typedArg0 := W4866
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16422 {
tmp15993 := MakeNative(func(__e *ControlFlow) {
W4869 := __e.Get(1)
_ = W4869
tmp16322 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4869, False)
}
__typedArg0 := W4869
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16322 {
tmp15994 := MakeNative(func(__e *ControlFlow) {
W4889 := __e.Get(1)
_ = W4889
tmp16217 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4889, False)
}
__typedArg0 := W4889
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16217 {
tmp15995 := MakeNative(func(__e *ControlFlow) {
W4911 := __e.Get(1)
_ = W4911
tmp16136 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4911, False)
}
__typedArg0 := W4911
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16136 {
tmp15996 := MakeNative(func(__e *ControlFlow) {
W4927 := __e.Get(1)
_ = W4927
tmp16036 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4927, False)
}
__typedArg0 := W4927
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16036 {
tmp15997 := MakeNative(func(__e *ControlFlow) {
W4947 := __e.Get(1)
_ = W4947
tmp15999 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4947, False)
}
__typedArg0 := W4947
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15999 {
__e.TailApply(PrimFunc(symshen_4unlock), V4862, W4865)
return
} else {
__e.Return(W4947)
return
}


}, 1)

tmp16034 := Call(__e, PrimFunc(symshen_4unlocked_2), V4862)


var ifres16000 Obj

if True == tmp16034 {
tmp16001 := MakeNative(func(__e *ControlFlow) {
W4948 := __e.Get(1)
_ = W4948
tmp16031 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4948)
}
__typedArg0 := W4948
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp16031 {
tmp16002 := MakeNative(func(__e *ControlFlow) {
W4949 := __e.Get(1)
_ = W4949
tmp16003 := MakeNative(func(__e *ControlFlow) {
W4950 := __e.Get(1)
_ = W4950
tmp16004 := MakeNative(func(__e *ControlFlow) {
W4951 := __e.Get(1)
_ = W4951
tmp16005 := MakeNative(func(__e *ControlFlow) {
W4952 := __e.Get(1)
_ = W4952
tmp16023 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4951)
}
__typedArg0 := W4951
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp16023 {
tmp16006 := MakeNative(func(__e *ControlFlow) {
W4955 := __e.Get(1)
_ = W4955
tmp16007 := MakeNative(func(__e *ControlFlow) {
W4956 := __e.Get(1)
_ = W4956
tmp16008 := Call(__e, W4952, W4955)


__e.TailApply(tmp16008, W4956)
return


}, 1)

tmp16009 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4951)
}
__typedArg0 := W4951
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(tmp16007, tmp16009)
return


}, 1)

tmp16010 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4951)
}
__typedArg0 := W4951
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp16006, tmp16010)
return


} else {
tmp16021 := Call(__e, PrimFunc(symshen_4pvar_2), W4951)


if True == tmp16021 {
tmp16011 := MakeNative(func(__e *ControlFlow) {
W4957 := __e.Get(1)
_ = W4957
tmp16012 := MakeNative(func(__e *ControlFlow) {
W4958 := __e.Get(1)
_ = W4958
tmp16013 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W4957, W4958)
}
__typedArg0 := W4957
__typedArg1 := W4958
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16014 := MakeNative(func(__e *ControlFlow) {
tmp16015 := Call(__e, W4952, W4957)


__e.TailApply(tmp16015, W4958)
return


}, 0)

tmp16016 := Call(__e, PrimFunc(symshen_4bind_b), W4951, tmp16013, V4861, tmp16014)


__e.TailApply(PrimFunc(symshen_4gc), V4861, tmp16016)
return


}, 1)

tmp16017 := Call(__e, PrimFunc(symshen_4newpv), V4861)


tmp16018 := Call(__e, tmp16012, tmp16017)


__e.TailApply(PrimFunc(symshen_4gc), V4861, tmp16018)
return


}, 1)

tmp16019 := Call(__e, PrimFunc(symshen_4newpv), V4861)


__e.TailApply(tmp16011, tmp16019)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp16024 := MakeNative(func(__e *ControlFlow) {
Z4953 := __e.Get(1)
_ = Z4953
__e.Return(MakeNative(func(__e *ControlFlow) {
Z4954 := __e.Get(1)
_ = Z4954
tmp16025 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16025

tmp16026 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4l_1rules), W4950, Z4954, V4860, V4861, V4862, W4865, V4864)
return
}, 0)

__e.TailApply(PrimFunc(symbind), Z4953, W4949, V4861, V4862, W4865, tmp16026)
return


}, 1))
return
}, 1)

__e.TailApply(tmp16005, tmp16024)
return


}, 1)

tmp16027 := Call(__e, PrimFunc(symshen_4lazyderef), V4859, V4861)


__e.TailApply(tmp16004, tmp16027)
return


}, 1)

tmp16028 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4948)
}
__typedArg0 := W4948
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(tmp16003, tmp16028)
return


}, 1)

tmp16029 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4948)
}
__typedArg0 := W4948
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp16002, tmp16029)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16032 := Call(__e, PrimFunc(symshen_4lazyderef), V4858, V4861)


tmp16033 := Call(__e, tmp16001, tmp16032)


ifres16000 = tmp16033


} else {
ifres16000 = False


}

__e.TailApply(tmp15997, ifres16000)
return


} else {
__e.Return(W4927)
return
}


}, 1)

tmp16134 := Call(__e, PrimFunc(symshen_4unlocked_2), V4862)


var ifres16037 Obj

if True == tmp16134 {
tmp16038 := MakeNative(func(__e *ControlFlow) {
W4928 := __e.Get(1)
_ = W4928
tmp16131 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4928)
}
__typedArg0 := W4928
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp16131 {
tmp16039 := MakeNative(func(__e *ControlFlow) {
W4929 := __e.Get(1)
_ = W4929
tmp16127 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4929)
}
__typedArg0 := W4929
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp16127 {
tmp16040 := MakeNative(func(__e *ControlFlow) {
W4930 := __e.Get(1)
_ = W4930
tmp16123 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4930)
}
__typedArg0 := W4930
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp16123 {
tmp16041 := MakeNative(func(__e *ControlFlow) {
W4931 := __e.Get(1)
_ = W4931
tmp16119 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4931, sym_8v)
}
__typedArg0 := W4931
__typedArg1 := sym_8v
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16119 {
tmp16042 := MakeNative(func(__e *ControlFlow) {
W4932 := __e.Get(1)
_ = W4932
tmp16115 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4932)
}
__typedArg0 := W4932
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp16115 {
tmp16043 := MakeNative(func(__e *ControlFlow) {
W4933 := __e.Get(1)
_ = W4933
tmp16044 := MakeNative(func(__e *ControlFlow) {
W4934 := __e.Get(1)
_ = W4934
tmp16110 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4934)
}
__typedArg0 := W4934
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp16110 {
tmp16045 := MakeNative(func(__e *ControlFlow) {
W4935 := __e.Get(1)
_ = W4935
tmp16046 := MakeNative(func(__e *ControlFlow) {
W4936 := __e.Get(1)
_ = W4936
tmp16105 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4936, Nil)
}
__typedArg0 := W4936
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16105 {
tmp16047 := MakeNative(func(__e *ControlFlow) {
W4937 := __e.Get(1)
_ = W4937
tmp16101 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4937)
}
__typedArg0 := W4937
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp16101 {
tmp16048 := MakeNative(func(__e *ControlFlow) {
W4938 := __e.Get(1)
_ = W4938
tmp16049 := MakeNative(func(__e *ControlFlow) {
W4939 := __e.Get(1)
_ = W4939
tmp16096 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4939)
}
__typedArg0 := W4939
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp16096 {
tmp16050 := MakeNative(func(__e *ControlFlow) {
W4940 := __e.Get(1)
_ = W4940
tmp16092 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4940)
}
__typedArg0 := W4940
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp16092 {
tmp16051 := MakeNative(func(__e *ControlFlow) {
W4941 := __e.Get(1)
_ = W4941
tmp16088 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4941, symvector)
}
__typedArg0 := W4941
__typedArg1 := symvector
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16088 {
tmp16052 := MakeNative(func(__e *ControlFlow) {
W4942 := __e.Get(1)
_ = W4942
tmp16084 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4942)
}
__typedArg0 := W4942
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp16084 {
tmp16053 := MakeNative(func(__e *ControlFlow) {
W4943 := __e.Get(1)
_ = W4943
tmp16054 := MakeNative(func(__e *ControlFlow) {
W4944 := __e.Get(1)
_ = W4944
tmp16079 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4944, Nil)
}
__typedArg0 := W4944
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16079 {
tmp16055 := MakeNative(func(__e *ControlFlow) {
W4945 := __e.Get(1)
_ = W4945
tmp16075 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4945, Nil)
}
__typedArg0 := W4945
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16075 {
tmp16056 := MakeNative(func(__e *ControlFlow) {
W4946 := __e.Get(1)
_ = W4946
tmp16057 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16057

tmp16058 := Call(__e, PrimFunc(symshen_4deref), W4938, V4861)


tmp16059 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp16060 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp16058, tmp16059)
}
__typedArg0 := tmp16058
__typedArg1 := tmp16059
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

tmp16061 := MakeNative(func(__e *ControlFlow) {
tmp16062 := MakeNative(func(__e *ControlFlow) {
tmp16063 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W4943, Nil)
}
__typedArg0 := W4943
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16064 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W4938, tmp16063)
}
__typedArg0 := W4938
__typedArg1 := tmp16063
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16065 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W4933, tmp16064)
}
__typedArg0 := W4933
__typedArg1 := tmp16064
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16066 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W4943, Nil)
}
__typedArg0 := W4943
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16067 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp16066)
}
__typedArg0 := symvector
__typedArg1 := tmp16066
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16068 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16067, Nil)
}
__typedArg0 := tmp16067
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16069 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W4938, tmp16068)
}
__typedArg0 := W4938
__typedArg1 := tmp16068
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16070 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W4935, tmp16069)
}
__typedArg0 := W4935
__typedArg1 := tmp16069
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16071 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16070, W4946)
}
__typedArg0 := tmp16070
__typedArg1 := W4946
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16072 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16065, tmp16071)
}
__typedArg0 := tmp16065
__typedArg1 := tmp16071
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4l_1rules), tmp16072, V4859, True, V4861, V4862, W4865, V4864)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V4861, V4862, W4865, tmp16062)
return


}, 0)

__e.TailApply(PrimFunc(symwhen), tmp16060, V4861, V4862, W4865, tmp16061)
return


}, 1)

tmp16073 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4928)
}
__typedArg0 := W4928
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(tmp16056, tmp16073)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16076 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4939)
}
__typedArg0 := W4939
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16077 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16076, V4861)


__e.TailApply(tmp16055, tmp16077)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16080 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4942)
}
__typedArg0 := W4942
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16081 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16080, V4861)


__e.TailApply(tmp16054, tmp16081)
return


}, 1)

tmp16082 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4942)
}
__typedArg0 := W4942
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp16053, tmp16082)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16085 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4940)
}
__typedArg0 := W4940
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16086 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16085, V4861)


__e.TailApply(tmp16052, tmp16086)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16089 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4940)
}
__typedArg0 := W4940
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16090 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16089, V4861)


__e.TailApply(tmp16051, tmp16090)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16093 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4939)
}
__typedArg0 := W4939
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16094 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16093, V4861)


__e.TailApply(tmp16050, tmp16094)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16097 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4937)
}
__typedArg0 := W4937
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16098 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16097, V4861)


__e.TailApply(tmp16049, tmp16098)
return


}, 1)

tmp16099 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4937)
}
__typedArg0 := W4937
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp16048, tmp16099)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16102 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4929)
}
__typedArg0 := W4929
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16103 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16102, V4861)


__e.TailApply(tmp16047, tmp16103)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16106 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4934)
}
__typedArg0 := W4934
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16107 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16106, V4861)


__e.TailApply(tmp16046, tmp16107)
return


}, 1)

tmp16108 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4934)
}
__typedArg0 := W4934
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp16045, tmp16108)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16111 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4932)
}
__typedArg0 := W4932
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16112 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16111, V4861)


__e.TailApply(tmp16044, tmp16112)
return


}, 1)

tmp16113 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4932)
}
__typedArg0 := W4932
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp16043, tmp16113)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16116 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4930)
}
__typedArg0 := W4930
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16117 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16116, V4861)


__e.TailApply(tmp16042, tmp16117)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16120 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4930)
}
__typedArg0 := W4930
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16121 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16120, V4861)


__e.TailApply(tmp16041, tmp16121)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16124 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4929)
}
__typedArg0 := W4929
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16125 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16124, V4861)


__e.TailApply(tmp16040, tmp16125)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16128 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4928)
}
__typedArg0 := W4928
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16129 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16128, V4861)


__e.TailApply(tmp16039, tmp16129)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16132 := Call(__e, PrimFunc(symshen_4lazyderef), V4858, V4861)


tmp16133 := Call(__e, tmp16038, tmp16132)


ifres16037 = tmp16133


} else {
ifres16037 = False


}

__e.TailApply(tmp15996, ifres16037)
return


} else {
__e.Return(W4911)
return
}


}, 1)

tmp16215 := Call(__e, PrimFunc(symshen_4unlocked_2), V4862)


var ifres16137 Obj

if True == tmp16215 {
tmp16138 := MakeNative(func(__e *ControlFlow) {
W4912 := __e.Get(1)
_ = W4912
tmp16212 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4912)
}
__typedArg0 := W4912
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp16212 {
tmp16139 := MakeNative(func(__e *ControlFlow) {
W4913 := __e.Get(1)
_ = W4913
tmp16208 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4913)
}
__typedArg0 := W4913
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp16208 {
tmp16140 := MakeNative(func(__e *ControlFlow) {
W4914 := __e.Get(1)
_ = W4914
tmp16204 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4914)
}
__typedArg0 := W4914
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp16204 {
tmp16141 := MakeNative(func(__e *ControlFlow) {
W4915 := __e.Get(1)
_ = W4915
tmp16200 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4915, sym_8s)
}
__typedArg0 := W4915
__typedArg1 := sym_8s
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16200 {
tmp16142 := MakeNative(func(__e *ControlFlow) {
W4916 := __e.Get(1)
_ = W4916
tmp16196 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4916)
}
__typedArg0 := W4916
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp16196 {
tmp16143 := MakeNative(func(__e *ControlFlow) {
W4917 := __e.Get(1)
_ = W4917
tmp16144 := MakeNative(func(__e *ControlFlow) {
W4918 := __e.Get(1)
_ = W4918
tmp16191 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4918)
}
__typedArg0 := W4918
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp16191 {
tmp16145 := MakeNative(func(__e *ControlFlow) {
W4919 := __e.Get(1)
_ = W4919
tmp16146 := MakeNative(func(__e *ControlFlow) {
W4920 := __e.Get(1)
_ = W4920
tmp16186 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4920, Nil)
}
__typedArg0 := W4920
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16186 {
tmp16147 := MakeNative(func(__e *ControlFlow) {
W4921 := __e.Get(1)
_ = W4921
tmp16182 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4921)
}
__typedArg0 := W4921
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp16182 {
tmp16148 := MakeNative(func(__e *ControlFlow) {
W4922 := __e.Get(1)
_ = W4922
tmp16149 := MakeNative(func(__e *ControlFlow) {
W4923 := __e.Get(1)
_ = W4923
tmp16177 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4923)
}
__typedArg0 := W4923
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp16177 {
tmp16150 := MakeNative(func(__e *ControlFlow) {
W4924 := __e.Get(1)
_ = W4924
tmp16173 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4924, symstring)
}
__typedArg0 := W4924
__typedArg1 := symstring
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16173 {
tmp16151 := MakeNative(func(__e *ControlFlow) {
W4925 := __e.Get(1)
_ = W4925
tmp16169 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4925, Nil)
}
__typedArg0 := W4925
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16169 {
tmp16152 := MakeNative(func(__e *ControlFlow) {
W4926 := __e.Get(1)
_ = W4926
tmp16153 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16153

tmp16154 := Call(__e, PrimFunc(symshen_4deref), W4922, V4861)


tmp16155 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp16156 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp16154, tmp16155)
}
__typedArg0 := tmp16154
__typedArg1 := tmp16155
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

tmp16157 := MakeNative(func(__e *ControlFlow) {
tmp16158 := MakeNative(func(__e *ControlFlow) {
tmp16159 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16160 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W4922, tmp16159)
}
__typedArg0 := W4922
__typedArg1 := tmp16159
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16161 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W4917, tmp16160)
}
__typedArg0 := W4917
__typedArg1 := tmp16160
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16162 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16163 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W4922, tmp16162)
}
__typedArg0 := W4922
__typedArg1 := tmp16162
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16164 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W4919, tmp16163)
}
__typedArg0 := W4919
__typedArg1 := tmp16163
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16165 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16164, W4926)
}
__typedArg0 := tmp16164
__typedArg1 := W4926
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16166 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16161, tmp16165)
}
__typedArg0 := tmp16161
__typedArg1 := tmp16165
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4l_1rules), tmp16166, V4859, True, V4861, V4862, W4865, V4864)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V4861, V4862, W4865, tmp16158)
return


}, 0)

__e.TailApply(PrimFunc(symwhen), tmp16156, V4861, V4862, W4865, tmp16157)
return


}, 1)

tmp16167 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4912)
}
__typedArg0 := W4912
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(tmp16152, tmp16167)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16170 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4923)
}
__typedArg0 := W4923
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16171 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16170, V4861)


__e.TailApply(tmp16151, tmp16171)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16174 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4923)
}
__typedArg0 := W4923
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16175 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16174, V4861)


__e.TailApply(tmp16150, tmp16175)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16178 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4921)
}
__typedArg0 := W4921
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16179 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16178, V4861)


__e.TailApply(tmp16149, tmp16179)
return


}, 1)

tmp16180 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4921)
}
__typedArg0 := W4921
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp16148, tmp16180)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16183 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4913)
}
__typedArg0 := W4913
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16184 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16183, V4861)


__e.TailApply(tmp16147, tmp16184)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16187 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4918)
}
__typedArg0 := W4918
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16188 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16187, V4861)


__e.TailApply(tmp16146, tmp16188)
return


}, 1)

tmp16189 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4918)
}
__typedArg0 := W4918
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp16145, tmp16189)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16192 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4916)
}
__typedArg0 := W4916
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16193 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16192, V4861)


__e.TailApply(tmp16144, tmp16193)
return


}, 1)

tmp16194 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4916)
}
__typedArg0 := W4916
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp16143, tmp16194)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16197 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4914)
}
__typedArg0 := W4914
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16198 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16197, V4861)


__e.TailApply(tmp16142, tmp16198)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16201 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4914)
}
__typedArg0 := W4914
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16202 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16201, V4861)


__e.TailApply(tmp16141, tmp16202)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16205 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4913)
}
__typedArg0 := W4913
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16206 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16205, V4861)


__e.TailApply(tmp16140, tmp16206)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16209 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4912)
}
__typedArg0 := W4912
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16210 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16209, V4861)


__e.TailApply(tmp16139, tmp16210)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16213 := Call(__e, PrimFunc(symshen_4lazyderef), V4858, V4861)


tmp16214 := Call(__e, tmp16138, tmp16213)


ifres16137 = tmp16214


} else {
ifres16137 = False


}

__e.TailApply(tmp15995, ifres16137)
return


} else {
__e.Return(W4889)
return
}


}, 1)

tmp16320 := Call(__e, PrimFunc(symshen_4unlocked_2), V4862)


var ifres16218 Obj

if True == tmp16320 {
tmp16219 := MakeNative(func(__e *ControlFlow) {
W4890 := __e.Get(1)
_ = W4890
tmp16317 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4890)
}
__typedArg0 := W4890
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp16317 {
tmp16220 := MakeNative(func(__e *ControlFlow) {
W4891 := __e.Get(1)
_ = W4891
tmp16313 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4891)
}
__typedArg0 := W4891
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp16313 {
tmp16221 := MakeNative(func(__e *ControlFlow) {
W4892 := __e.Get(1)
_ = W4892
tmp16309 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4892)
}
__typedArg0 := W4892
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp16309 {
tmp16222 := MakeNative(func(__e *ControlFlow) {
W4893 := __e.Get(1)
_ = W4893
tmp16305 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4893, sym_8p)
}
__typedArg0 := W4893
__typedArg1 := sym_8p
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16305 {
tmp16223 := MakeNative(func(__e *ControlFlow) {
W4894 := __e.Get(1)
_ = W4894
tmp16301 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4894)
}
__typedArg0 := W4894
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp16301 {
tmp16224 := MakeNative(func(__e *ControlFlow) {
W4895 := __e.Get(1)
_ = W4895
tmp16225 := MakeNative(func(__e *ControlFlow) {
W4896 := __e.Get(1)
_ = W4896
tmp16296 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4896)
}
__typedArg0 := W4896
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp16296 {
tmp16226 := MakeNative(func(__e *ControlFlow) {
W4897 := __e.Get(1)
_ = W4897
tmp16227 := MakeNative(func(__e *ControlFlow) {
W4898 := __e.Get(1)
_ = W4898
tmp16291 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4898, Nil)
}
__typedArg0 := W4898
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16291 {
tmp16228 := MakeNative(func(__e *ControlFlow) {
W4899 := __e.Get(1)
_ = W4899
tmp16287 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4899)
}
__typedArg0 := W4899
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp16287 {
tmp16229 := MakeNative(func(__e *ControlFlow) {
W4900 := __e.Get(1)
_ = W4900
tmp16230 := MakeNative(func(__e *ControlFlow) {
W4901 := __e.Get(1)
_ = W4901
tmp16282 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4901)
}
__typedArg0 := W4901
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp16282 {
tmp16231 := MakeNative(func(__e *ControlFlow) {
W4902 := __e.Get(1)
_ = W4902
tmp16278 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4902)
}
__typedArg0 := W4902
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp16278 {
tmp16232 := MakeNative(func(__e *ControlFlow) {
W4903 := __e.Get(1)
_ = W4903
tmp16233 := MakeNative(func(__e *ControlFlow) {
W4904 := __e.Get(1)
_ = W4904
tmp16273 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4904)
}
__typedArg0 := W4904
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp16273 {
tmp16234 := MakeNative(func(__e *ControlFlow) {
W4905 := __e.Get(1)
_ = W4905
tmp16269 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4905, sym_d)
}
__typedArg0 := W4905
__typedArg1 := sym_d
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16269 {
tmp16235 := MakeNative(func(__e *ControlFlow) {
W4906 := __e.Get(1)
_ = W4906
tmp16265 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4906)
}
__typedArg0 := W4906
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp16265 {
tmp16236 := MakeNative(func(__e *ControlFlow) {
W4907 := __e.Get(1)
_ = W4907
tmp16237 := MakeNative(func(__e *ControlFlow) {
W4908 := __e.Get(1)
_ = W4908
tmp16260 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4908, Nil)
}
__typedArg0 := W4908
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16260 {
tmp16238 := MakeNative(func(__e *ControlFlow) {
W4909 := __e.Get(1)
_ = W4909
tmp16256 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4909, Nil)
}
__typedArg0 := W4909
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16256 {
tmp16239 := MakeNative(func(__e *ControlFlow) {
W4910 := __e.Get(1)
_ = W4910
tmp16240 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16240

tmp16241 := Call(__e, PrimFunc(symshen_4deref), W4900, V4861)


tmp16242 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp16243 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp16241, tmp16242)
}
__typedArg0 := tmp16241
__typedArg1 := tmp16242
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

tmp16244 := MakeNative(func(__e *ControlFlow) {
tmp16245 := MakeNative(func(__e *ControlFlow) {
tmp16246 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W4903, Nil)
}
__typedArg0 := W4903
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16247 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W4900, tmp16246)
}
__typedArg0 := W4900
__typedArg1 := tmp16246
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16248 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W4895, tmp16247)
}
__typedArg0 := W4895
__typedArg1 := tmp16247
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16249 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W4907, Nil)
}
__typedArg0 := W4907
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16250 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W4900, tmp16249)
}
__typedArg0 := W4900
__typedArg1 := tmp16249
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16251 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W4897, tmp16250)
}
__typedArg0 := W4897
__typedArg1 := tmp16250
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16252 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16251, W4910)
}
__typedArg0 := tmp16251
__typedArg1 := W4910
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16253 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16248, tmp16252)
}
__typedArg0 := tmp16248
__typedArg1 := tmp16252
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4l_1rules), tmp16253, V4859, True, V4861, V4862, W4865, V4864)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V4861, V4862, W4865, tmp16245)
return


}, 0)

__e.TailApply(PrimFunc(symwhen), tmp16243, V4861, V4862, W4865, tmp16244)
return


}, 1)

tmp16254 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4890)
}
__typedArg0 := W4890
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(tmp16239, tmp16254)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16257 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4901)
}
__typedArg0 := W4901
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16258 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16257, V4861)


__e.TailApply(tmp16238, tmp16258)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16261 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4906)
}
__typedArg0 := W4906
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16262 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16261, V4861)


__e.TailApply(tmp16237, tmp16262)
return


}, 1)

tmp16263 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4906)
}
__typedArg0 := W4906
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp16236, tmp16263)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16266 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4904)
}
__typedArg0 := W4904
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16267 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16266, V4861)


__e.TailApply(tmp16235, tmp16267)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16270 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4904)
}
__typedArg0 := W4904
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16271 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16270, V4861)


__e.TailApply(tmp16234, tmp16271)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16274 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4902)
}
__typedArg0 := W4902
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16275 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16274, V4861)


__e.TailApply(tmp16233, tmp16275)
return


}, 1)

tmp16276 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4902)
}
__typedArg0 := W4902
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp16232, tmp16276)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16279 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4901)
}
__typedArg0 := W4901
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16280 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16279, V4861)


__e.TailApply(tmp16231, tmp16280)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16283 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4899)
}
__typedArg0 := W4899
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16284 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16283, V4861)


__e.TailApply(tmp16230, tmp16284)
return


}, 1)

tmp16285 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4899)
}
__typedArg0 := W4899
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp16229, tmp16285)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16288 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4891)
}
__typedArg0 := W4891
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16289 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16288, V4861)


__e.TailApply(tmp16228, tmp16289)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16292 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4896)
}
__typedArg0 := W4896
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16293 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16292, V4861)


__e.TailApply(tmp16227, tmp16293)
return


}, 1)

tmp16294 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4896)
}
__typedArg0 := W4896
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp16226, tmp16294)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16297 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4894)
}
__typedArg0 := W4894
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16298 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16297, V4861)


__e.TailApply(tmp16225, tmp16298)
return


}, 1)

tmp16299 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4894)
}
__typedArg0 := W4894
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp16224, tmp16299)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16302 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4892)
}
__typedArg0 := W4892
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16303 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16302, V4861)


__e.TailApply(tmp16223, tmp16303)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16306 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4892)
}
__typedArg0 := W4892
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16307 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16306, V4861)


__e.TailApply(tmp16222, tmp16307)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16310 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4891)
}
__typedArg0 := W4891
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16311 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16310, V4861)


__e.TailApply(tmp16221, tmp16311)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16314 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4890)
}
__typedArg0 := W4890
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16315 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16314, V4861)


__e.TailApply(tmp16220, tmp16315)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16318 := Call(__e, PrimFunc(symshen_4lazyderef), V4858, V4861)


tmp16319 := Call(__e, tmp16219, tmp16318)


ifres16218 = tmp16319


} else {
ifres16218 = False


}

__e.TailApply(tmp15994, ifres16218)
return


} else {
__e.Return(W4869)
return
}


}, 1)

tmp16420 := Call(__e, PrimFunc(symshen_4unlocked_2), V4862)


var ifres16323 Obj

if True == tmp16420 {
tmp16324 := MakeNative(func(__e *ControlFlow) {
W4870 := __e.Get(1)
_ = W4870
tmp16417 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4870)
}
__typedArg0 := W4870
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp16417 {
tmp16325 := MakeNative(func(__e *ControlFlow) {
W4871 := __e.Get(1)
_ = W4871
tmp16413 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4871)
}
__typedArg0 := W4871
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp16413 {
tmp16326 := MakeNative(func(__e *ControlFlow) {
W4872 := __e.Get(1)
_ = W4872
tmp16409 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4872)
}
__typedArg0 := W4872
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp16409 {
tmp16327 := MakeNative(func(__e *ControlFlow) {
W4873 := __e.Get(1)
_ = W4873
tmp16405 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4873, symcons)
}
__typedArg0 := W4873
__typedArg1 := symcons
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16405 {
tmp16328 := MakeNative(func(__e *ControlFlow) {
W4874 := __e.Get(1)
_ = W4874
tmp16401 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4874)
}
__typedArg0 := W4874
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp16401 {
tmp16329 := MakeNative(func(__e *ControlFlow) {
W4875 := __e.Get(1)
_ = W4875
tmp16330 := MakeNative(func(__e *ControlFlow) {
W4876 := __e.Get(1)
_ = W4876
tmp16396 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4876)
}
__typedArg0 := W4876
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp16396 {
tmp16331 := MakeNative(func(__e *ControlFlow) {
W4877 := __e.Get(1)
_ = W4877
tmp16332 := MakeNative(func(__e *ControlFlow) {
W4878 := __e.Get(1)
_ = W4878
tmp16391 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4878, Nil)
}
__typedArg0 := W4878
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16391 {
tmp16333 := MakeNative(func(__e *ControlFlow) {
W4879 := __e.Get(1)
_ = W4879
tmp16387 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4879)
}
__typedArg0 := W4879
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp16387 {
tmp16334 := MakeNative(func(__e *ControlFlow) {
W4880 := __e.Get(1)
_ = W4880
tmp16335 := MakeNative(func(__e *ControlFlow) {
W4881 := __e.Get(1)
_ = W4881
tmp16382 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4881)
}
__typedArg0 := W4881
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp16382 {
tmp16336 := MakeNative(func(__e *ControlFlow) {
W4882 := __e.Get(1)
_ = W4882
tmp16378 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4882)
}
__typedArg0 := W4882
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp16378 {
tmp16337 := MakeNative(func(__e *ControlFlow) {
W4883 := __e.Get(1)
_ = W4883
tmp16374 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4883, symlist)
}
__typedArg0 := W4883
__typedArg1 := symlist
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16374 {
tmp16338 := MakeNative(func(__e *ControlFlow) {
W4884 := __e.Get(1)
_ = W4884
tmp16370 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4884)
}
__typedArg0 := W4884
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp16370 {
tmp16339 := MakeNative(func(__e *ControlFlow) {
W4885 := __e.Get(1)
_ = W4885
tmp16340 := MakeNative(func(__e *ControlFlow) {
W4886 := __e.Get(1)
_ = W4886
tmp16365 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4886, Nil)
}
__typedArg0 := W4886
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16365 {
tmp16341 := MakeNative(func(__e *ControlFlow) {
W4887 := __e.Get(1)
_ = W4887
tmp16361 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4887, Nil)
}
__typedArg0 := W4887
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16361 {
tmp16342 := MakeNative(func(__e *ControlFlow) {
W4888 := __e.Get(1)
_ = W4888
tmp16343 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16343

tmp16344 := Call(__e, PrimFunc(symshen_4deref), W4880, V4861)


tmp16345 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp16346 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp16344, tmp16345)
}
__typedArg0 := tmp16344
__typedArg1 := tmp16345
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

tmp16347 := MakeNative(func(__e *ControlFlow) {
tmp16348 := MakeNative(func(__e *ControlFlow) {
tmp16349 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W4885, Nil)
}
__typedArg0 := W4885
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16350 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W4880, tmp16349)
}
__typedArg0 := W4880
__typedArg1 := tmp16349
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16351 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W4875, tmp16350)
}
__typedArg0 := W4875
__typedArg1 := tmp16350
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16352 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W4885, Nil)
}
__typedArg0 := W4885
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16353 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp16352)
}
__typedArg0 := symlist
__typedArg1 := tmp16352
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16354 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16353, Nil)
}
__typedArg0 := tmp16353
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16355 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W4880, tmp16354)
}
__typedArg0 := W4880
__typedArg1 := tmp16354
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16356 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W4877, tmp16355)
}
__typedArg0 := W4877
__typedArg1 := tmp16355
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16357 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16356, W4888)
}
__typedArg0 := tmp16356
__typedArg1 := W4888
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16358 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16351, tmp16357)
}
__typedArg0 := tmp16351
__typedArg1 := tmp16357
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4l_1rules), tmp16358, V4859, True, V4861, V4862, W4865, V4864)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V4861, V4862, W4865, tmp16348)
return


}, 0)

__e.TailApply(PrimFunc(symwhen), tmp16346, V4861, V4862, W4865, tmp16347)
return


}, 1)

tmp16359 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4870)
}
__typedArg0 := W4870
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(tmp16342, tmp16359)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16362 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4881)
}
__typedArg0 := W4881
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16363 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16362, V4861)


__e.TailApply(tmp16341, tmp16363)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16366 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4884)
}
__typedArg0 := W4884
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16367 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16366, V4861)


__e.TailApply(tmp16340, tmp16367)
return


}, 1)

tmp16368 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4884)
}
__typedArg0 := W4884
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp16339, tmp16368)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16371 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4882)
}
__typedArg0 := W4882
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16372 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16371, V4861)


__e.TailApply(tmp16338, tmp16372)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16375 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4882)
}
__typedArg0 := W4882
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16376 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16375, V4861)


__e.TailApply(tmp16337, tmp16376)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16379 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4881)
}
__typedArg0 := W4881
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16380 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16379, V4861)


__e.TailApply(tmp16336, tmp16380)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16383 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4879)
}
__typedArg0 := W4879
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16384 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16383, V4861)


__e.TailApply(tmp16335, tmp16384)
return


}, 1)

tmp16385 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4879)
}
__typedArg0 := W4879
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp16334, tmp16385)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16388 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4871)
}
__typedArg0 := W4871
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16389 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16388, V4861)


__e.TailApply(tmp16333, tmp16389)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16392 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4876)
}
__typedArg0 := W4876
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16393 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16392, V4861)


__e.TailApply(tmp16332, tmp16393)
return


}, 1)

tmp16394 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4876)
}
__typedArg0 := W4876
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp16331, tmp16394)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16397 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4874)
}
__typedArg0 := W4874
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16398 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16397, V4861)


__e.TailApply(tmp16330, tmp16398)
return


}, 1)

tmp16399 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4874)
}
__typedArg0 := W4874
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp16329, tmp16399)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16402 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4872)
}
__typedArg0 := W4872
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16403 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16402, V4861)


__e.TailApply(tmp16328, tmp16403)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16406 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4872)
}
__typedArg0 := W4872
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16407 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16406, V4861)


__e.TailApply(tmp16327, tmp16407)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16410 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4871)
}
__typedArg0 := W4871
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16411 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16410, V4861)


__e.TailApply(tmp16326, tmp16411)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16414 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4870)
}
__typedArg0 := W4870
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16415 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16414, V4861)


__e.TailApply(tmp16325, tmp16415)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16418 := Call(__e, PrimFunc(symshen_4lazyderef), V4858, V4861)


tmp16419 := Call(__e, tmp16324, tmp16418)


ifres16323 = tmp16419


} else {
ifres16323 = False


}

__e.TailApply(tmp15993, ifres16323)
return


} else {
__e.Return(W4866)
return
}


}, 1)

tmp16435 := Call(__e, PrimFunc(symshen_4unlocked_2), V4862)


var ifres16423 Obj

if True == tmp16435 {
tmp16424 := MakeNative(func(__e *ControlFlow) {
W4867 := __e.Get(1)
_ = W4867
tmp16432 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4867, Nil)
}
__typedArg0 := W4867
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16432 {
tmp16425 := MakeNative(func(__e *ControlFlow) {
W4868 := __e.Get(1)
_ = W4868
tmp16429 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4868, True)
}
__typedArg0 := W4868
__typedArg1 := True
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16429 {
tmp16426 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16426

tmp16427 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symbind), V4859, Nil, V4861, V4862, W4865, V4864)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V4861, V4862, W4865, tmp16427)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16430 := Call(__e, PrimFunc(symshen_4lazyderef), V4860, V4861)


__e.TailApply(tmp16425, tmp16430)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16433 := Call(__e, PrimFunc(symshen_4lazyderef), V4858, V4861)


tmp16434 := Call(__e, tmp16424, tmp16433)


ifres16423 = tmp16434


} else {
ifres16423 = False


}

__e.TailApply(tmp15992, ifres16423)
return


}, 1)

__e.TailApply(tmp15991, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V4863)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V4863
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())
return


}, 7)

tmp16437 := Call(__e, ns2_1set, symshen_4l_1rules, tmp15990)


_ = tmp16437

tmp16438 := MakeNative(func(__e *ControlFlow) {
V4959 := __e.Get(1)
_ = V4959
V4960 := __e.Get(2)
_ = V4960
V4961 := __e.Get(3)
_ = V4961
V4962 := __e.Get(4)
_ = V4962
V4963 := __e.Get(5)
_ = V4963
V4964 := __e.Get(6)
_ = V4964
tmp16439 := MakeNative(func(__e *ControlFlow) {
W4965 := __e.Get(1)
_ = W4965
tmp16440 := MakeNative(func(__e *ControlFlow) {
W4966 := __e.Get(1)
_ = W4966
tmp16442 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4966, False)
}
__typedArg0 := W4966
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16442 {
__e.TailApply(PrimFunc(symshen_4unlock), V4962, W4965)
return
} else {
__e.Return(W4966)
return
}


}, 1)

tmp16490 := Call(__e, PrimFunc(symshen_4unlocked_2), V4962)


var ifres16443 Obj

if True == tmp16490 {
tmp16444 := MakeNative(func(__e *ControlFlow) {
W4967 := __e.Get(1)
_ = W4967
tmp16487 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4967)
}
__typedArg0 := W4967
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp16487 {
tmp16445 := MakeNative(func(__e *ControlFlow) {
W4968 := __e.Get(1)
_ = W4968
tmp16483 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W4968, symdefine)
}
__typedArg0 := W4968
__typedArg1 := symdefine
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16483 {
tmp16446 := MakeNative(func(__e *ControlFlow) {
W4969 := __e.Get(1)
_ = W4969
tmp16479 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W4969)
}
__typedArg0 := W4969
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp16479 {
tmp16447 := MakeNative(func(__e *ControlFlow) {
W4970 := __e.Get(1)
_ = W4970
tmp16448 := MakeNative(func(__e *ControlFlow) {
W4971 := __e.Get(1)
_ = W4971
tmp16449 := MakeNative(func(__e *ControlFlow) {
W4972 := __e.Get(1)
_ = W4972
tmp16450 := MakeNative(func(__e *ControlFlow) {
W4973 := __e.Get(1)
_ = W4973
tmp16451 := MakeNative(func(__e *ControlFlow) {
W4974 := __e.Get(1)
_ = W4974
tmp16452 := MakeNative(func(__e *ControlFlow) {
W4975 := __e.Get(1)
_ = W4975
tmp16453 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16453

tmp16454 := MakeNative(func(__e *ControlFlow) {
tmp16455 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W4970, W4971)
}
__typedArg0 := W4970
__typedArg1 := W4971
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16456 := Call(__e, PrimFunc(symshen_4sigxrules), tmp16455)


tmp16457 := MakeNative(func(__e *ControlFlow) {
tmp16458 := Call(__e, PrimFunc(symshen_4lazyderef), W4972, V4961)


tmp16459 := Call(__e, PrimFunc(symfst), tmp16458)


tmp16460 := MakeNative(func(__e *ControlFlow) {
tmp16461 := Call(__e, PrimFunc(symshen_4lazyderef), W4972, V4961)


tmp16462 := Call(__e, PrimFunc(symsnd), tmp16461)


tmp16463 := MakeNative(func(__e *ControlFlow) {
tmp16464 := Call(__e, PrimFunc(symshen_4deref), W4975, V4961)


tmp16465 := Call(__e, PrimFunc(symshen_4freshen_1sig), tmp16464)


tmp16466 := MakeNative(func(__e *ControlFlow) {
tmp16467 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symis), W4975, V4960, V4961, V4962, W4965, V4964)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4t_d_1rules), W4970, W4973, W4974, MakeNumber(1), V4961, V4962, W4965, tmp16467)
return


}, 0)

__e.TailApply(PrimFunc(symbind), W4974, tmp16465, V4961, V4962, W4965, tmp16466)
return


}, 0)

__e.TailApply(PrimFunc(symbind), W4973, tmp16462, V4961, V4962, W4965, tmp16463)
return


}, 0)

__e.TailApply(PrimFunc(symbind), W4975, tmp16459, V4961, V4962, W4965, tmp16460)
return


}, 0)

__e.TailApply(PrimFunc(symbind), W4972, tmp16456, V4961, V4962, W4965, tmp16457)
return


}, 0)

tmp16468 := Call(__e, PrimFunc(symshen_4cut), V4961, V4962, W4965, tmp16454)


__e.TailApply(PrimFunc(symshen_4gc), V4961, tmp16468)
return


}, 1)

tmp16469 := Call(__e, PrimFunc(symshen_4newpv), V4961)


tmp16470 := Call(__e, tmp16452, tmp16469)


__e.TailApply(PrimFunc(symshen_4gc), V4961, tmp16470)
return


}, 1)

tmp16471 := Call(__e, PrimFunc(symshen_4newpv), V4961)


tmp16472 := Call(__e, tmp16451, tmp16471)


__e.TailApply(PrimFunc(symshen_4gc), V4961, tmp16472)
return


}, 1)

tmp16473 := Call(__e, PrimFunc(symshen_4newpv), V4961)


tmp16474 := Call(__e, tmp16450, tmp16473)


__e.TailApply(PrimFunc(symshen_4gc), V4961, tmp16474)
return


}, 1)

tmp16475 := Call(__e, PrimFunc(symshen_4newpv), V4961)


__e.TailApply(tmp16449, tmp16475)
return


}, 1)

tmp16476 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4969)
}
__typedArg0 := W4969
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(tmp16448, tmp16476)
return


}, 1)

tmp16477 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4969)
}
__typedArg0 := W4969
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp16447, tmp16477)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16480 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W4967)
}
__typedArg0 := W4967
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16481 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16480, V4961)


__e.TailApply(tmp16446, tmp16481)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16484 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W4967)
}
__typedArg0 := W4967
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16485 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16484, V4961)


__e.TailApply(tmp16445, tmp16485)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16488 := Call(__e, PrimFunc(symshen_4lazyderef), V4959, V4961)


tmp16489 := Call(__e, tmp16444, tmp16488)


ifres16443 = tmp16489


} else {
ifres16443 = False


}

__e.TailApply(tmp16440, ifres16443)
return


}, 1)

__e.TailApply(tmp16439, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V4963)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V4963
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())
return


}, 6)

tmp16492 := Call(__e, ns2_1set, symshen_4t_d, tmp16438)


_ = tmp16492

tmp16493 := MakeNative(func(__e *ControlFlow) {
V4976 := __e.Get(1)
_ = V4976
tmp16494 := MakeNative(func(__e *ControlFlow) {
Z4977 := __e.Get(1)
_ = Z4977
__e.TailApply(PrimFunc(symshen_4_5sig_drules_6), Z4977)
return
}, 1)

__e.TailApply(PrimFunc(symcompile), tmp16494, V4976)
return


}, 1)

tmp16495 := Call(__e, ns2_1set, symshen_4sigxrules, tmp16493)


_ = tmp16495

tmp16496 := MakeNative(func(__e *ControlFlow) {
V4978 := __e.Get(1)
_ = V4978
tmp16497 := MakeNative(func(__e *ControlFlow) {
W4979 := __e.Get(1)
_ = W4979
tmp16499 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W4979)


if True == tmp16499 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W4979)
return
}


}, 1)

tmp16532 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4978)
}
__typedArg0 := V4978
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16500 Obj

if True == tmp16532 {
tmp16501 := MakeNative(func(__e *ControlFlow) {
W4980 := __e.Get(1)
_ = W4980
tmp16528 := Call(__e, PrimFunc(symshen_4hds_a_2), W4980, sym_i)


if True == tmp16528 {
tmp16502 := MakeNative(func(__e *ControlFlow) {
W4981 := __e.Get(1)
_ = W4981
tmp16503 := MakeNative(func(__e *ControlFlow) {
W4982 := __e.Get(1)
_ = W4982
tmp16524 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W4982)


if True == tmp16524 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp16504 := MakeNative(func(__e *ControlFlow) {
W4983 := __e.Get(1)
_ = W4983
tmp16505 := MakeNative(func(__e *ControlFlow) {
W4984 := __e.Get(1)
_ = W4984
tmp16520 := Call(__e, PrimFunc(symshen_4hds_a_2), W4984, sym_j)


if True == tmp16520 {
tmp16506 := MakeNative(func(__e *ControlFlow) {
W4985 := __e.Get(1)
_ = W4985
tmp16507 := MakeNative(func(__e *ControlFlow) {
W4986 := __e.Get(1)
_ = W4986
tmp16516 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W4986)


if True == tmp16516 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp16508 := MakeNative(func(__e *ControlFlow) {
W4987 := __e.Get(1)
_ = W4987
tmp16509 := MakeNative(func(__e *ControlFlow) {
W4988 := __e.Get(1)
_ = W4988
tmp16510 := MakeNative(func(__e *ControlFlow) {
W4989 := __e.Get(1)
_ = W4989
__e.TailApply(PrimFunc(sym_8p), W4989, W4987)
return
}, 1)

tmp16511 := Call(__e, PrimFunc(symshen_4rectify_1type), W4983)


tmp16512 := Call(__e, tmp16510, tmp16511)


__e.TailApply(PrimFunc(symshen_4comb), W4988, tmp16512)
return


}, 1)

tmp16513 := Call(__e, PrimFunc(symshen_4in_1_6), W4986)


__e.TailApply(tmp16509, tmp16513)
return


}, 1)

tmp16514 := Call(__e, PrimFunc(symshen_4_5_1out), W4986)


__e.TailApply(tmp16508, tmp16514)
return


}


}, 1)

tmp16517 := Call(__e, PrimFunc(symshen_4_5rules_d_6), W4985)


__e.TailApply(tmp16507, tmp16517)
return


}, 1)

tmp16518 := Call(__e, PrimFunc(symtail), W4984)


__e.TailApply(tmp16506, tmp16518)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp16521 := Call(__e, PrimFunc(symshen_4in_1_6), W4982)


__e.TailApply(tmp16505, tmp16521)
return


}, 1)

tmp16522 := Call(__e, PrimFunc(symshen_4_5_1out), W4982)


__e.TailApply(tmp16504, tmp16522)
return


}


}, 1)

tmp16525 := Call(__e, PrimFunc(symshen_4_5signature_6), W4981)


__e.TailApply(tmp16503, tmp16525)
return


}, 1)

tmp16526 := Call(__e, PrimFunc(symtail), W4980)


__e.TailApply(tmp16502, tmp16526)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp16529 := Call(__e, PrimFunc(symtail), V4978)


tmp16530 := Call(__e, tmp16501, tmp16529)


ifres16500 = tmp16530


} else {
tmp16531 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16500 = tmp16531


}

__e.TailApply(tmp16497, ifres16500)
return


}, 1)

tmp16533 := Call(__e, ns2_1set, symshen_4_5sig_drules_6, tmp16496)


_ = tmp16533

tmp16534 := MakeNative(func(__e *ControlFlow) {
V4990 := __e.Get(1)
_ = V4990
tmp16535 := MakeNative(func(__e *ControlFlow) {
W4991 := __e.Get(1)
_ = W4991
tmp16536 := MakeNative(func(__e *ControlFlow) {
W4992 := __e.Get(1)
_ = W4992
__e.TailApply(PrimFunc(symshen_4freshen_1type), W4992, V4990)
return
}, 1)

tmp16537 := MakeNative(func(__e *ControlFlow) {
Z4993 := __e.Get(1)
_ = Z4993
tmp16538 := Call(__e, PrimFunc(symconcat), sym_e, Z4993)


tmp16539 := Call(__e, PrimFunc(symshen_4freshterm), tmp16538)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Z4993, tmp16539)
}
__typedArg0 := Z4993
__typedArg1 := tmp16539
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp16540 := Call(__e, PrimFunc(symmap), tmp16537, W4991)


__e.TailApply(tmp16536, tmp16540)
return


}, 1)

tmp16541 := Call(__e, PrimFunc(symshen_4extract_1vars), V4990)


__e.TailApply(tmp16535, tmp16541)
return


}, 1)

tmp16542 := Call(__e, ns2_1set, symshen_4freshen_1sig, tmp16534)


_ = tmp16542

tmp16543 := MakeNative(func(__e *ControlFlow) {
V4994 := __e.Get(1)
_ = V4994
V4995 := __e.Get(2)
_ = V4995
tmp16557 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V4994)
}
__typedArg0 := Nil
__typedArg1 := V4994
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16557 {
__e.Return(V4995)
return
} else {
tmp16555 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4994)
}
__typedArg0 := V4994
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16551 Obj

if True == tmp16555 {
tmp16553 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4994)
}
__typedArg0 := V4994
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16554 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp16553)
}
__typedArg0 := tmp16553
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16552 Obj

if True == tmp16554 {
ifres16552 = True


} else {
ifres16552 = False


}

ifres16551 = ifres16552


} else {
ifres16551 = False


}

if True == ifres16551 {
tmp16544 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4994)
}
__typedArg0 := V4994
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16545 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4994)
}
__typedArg0 := V4994
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16546 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16545)
}
__typedArg0 := tmp16545
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16547 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4994)
}
__typedArg0 := V4994
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16548 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp16547)
}
__typedArg0 := tmp16547
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16549 := Call(__e, PrimFunc(symsubst), tmp16546, tmp16548, V4995)


__e.TailApply(PrimFunc(symshen_4freshen_1type), tmp16544, tmp16549)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("partial function shen.freshen-type"))
}
__typedArg0 := MakeString("partial function shen.freshen-type")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 2)

tmp16558 := Call(__e, ns2_1set, symshen_4freshen_1type, tmp16543)


_ = tmp16558

tmp16559 := MakeNative(func(__e *ControlFlow) {
V4996 := __e.Get(1)
_ = V4996
tmp16560 := MakeNative(func(__e *ControlFlow) {
W4997 := __e.Get(1)
_ = W4997
tmp16575 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W4997)


if True == tmp16575 {
tmp16561 := MakeNative(func(__e *ControlFlow) {
W5004 := __e.Get(1)
_ = W5004
tmp16563 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W5004)


if True == tmp16563 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W5004)
return
}


}, 1)

tmp16564 := MakeNative(func(__e *ControlFlow) {
W5005 := __e.Get(1)
_ = W5005
tmp16571 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W5005)


if True == tmp16571 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp16565 := MakeNative(func(__e *ControlFlow) {
W5006 := __e.Get(1)
_ = W5006
tmp16566 := MakeNative(func(__e *ControlFlow) {
W5007 := __e.Get(1)
_ = W5007
tmp16567 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W5006, Nil)
}
__typedArg0 := W5006
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4comb), W5007, tmp16567)
return


}, 1)

tmp16568 := Call(__e, PrimFunc(symshen_4in_1_6), W5005)


__e.TailApply(tmp16566, tmp16568)
return


}, 1)

tmp16569 := Call(__e, PrimFunc(symshen_4_5_1out), W5005)


__e.TailApply(tmp16565, tmp16569)
return


}


}, 1)

tmp16572 := Call(__e, PrimFunc(symshen_4_5rule_d_6), V4996)


tmp16573 := Call(__e, tmp16564, tmp16572)


__e.TailApply(tmp16561, tmp16573)
return


} else {
__e.Return(W4997)
return
}


}, 1)

tmp16576 := MakeNative(func(__e *ControlFlow) {
W4998 := __e.Get(1)
_ = W4998
tmp16591 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W4998)


if True == tmp16591 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp16577 := MakeNative(func(__e *ControlFlow) {
W4999 := __e.Get(1)
_ = W4999
tmp16578 := MakeNative(func(__e *ControlFlow) {
W5000 := __e.Get(1)
_ = W5000
tmp16579 := MakeNative(func(__e *ControlFlow) {
W5001 := __e.Get(1)
_ = W5001
tmp16586 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W5001)


if True == tmp16586 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp16580 := MakeNative(func(__e *ControlFlow) {
W5002 := __e.Get(1)
_ = W5002
tmp16581 := MakeNative(func(__e *ControlFlow) {
W5003 := __e.Get(1)
_ = W5003
tmp16582 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W4999, W5002)
}
__typedArg0 := W4999
__typedArg1 := W5002
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4comb), W5003, tmp16582)
return


}, 1)

tmp16583 := Call(__e, PrimFunc(symshen_4in_1_6), W5001)


__e.TailApply(tmp16581, tmp16583)
return


}, 1)

tmp16584 := Call(__e, PrimFunc(symshen_4_5_1out), W5001)


__e.TailApply(tmp16580, tmp16584)
return


}


}, 1)

tmp16587 := Call(__e, PrimFunc(symshen_4_5rules_d_6), W5000)


__e.TailApply(tmp16579, tmp16587)
return


}, 1)

tmp16588 := Call(__e, PrimFunc(symshen_4in_1_6), W4998)


__e.TailApply(tmp16578, tmp16588)
return


}, 1)

tmp16589 := Call(__e, PrimFunc(symshen_4_5_1out), W4998)


__e.TailApply(tmp16577, tmp16589)
return


}


}, 1)

tmp16592 := Call(__e, PrimFunc(symshen_4_5rule_d_6), V4996)


tmp16593 := Call(__e, tmp16576, tmp16592)


__e.TailApply(tmp16560, tmp16593)
return


}, 1)

tmp16594 := Call(__e, ns2_1set, symshen_4_5rules_d_6, tmp16559)


_ = tmp16594

tmp16595 := MakeNative(func(__e *ControlFlow) {
V5008 := __e.Get(1)
_ = V5008
tmp16596 := MakeNative(func(__e *ControlFlow) {
W5009 := __e.Get(1)
_ = W5009
tmp16682 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W5009)


if True == tmp16682 {
tmp16597 := MakeNative(func(__e *ControlFlow) {
W5019 := __e.Get(1)
_ = W5019
tmp16646 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W5019)


if True == tmp16646 {
tmp16598 := MakeNative(func(__e *ControlFlow) {
W5029 := __e.Get(1)
_ = W5029
tmp16623 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W5029)


if True == tmp16623 {
tmp16599 := MakeNative(func(__e *ControlFlow) {
W5036 := __e.Get(1)
_ = W5036
tmp16601 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W5036)


if True == tmp16601 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W5036)
return
}


}, 1)

tmp16602 := MakeNative(func(__e *ControlFlow) {
W5037 := __e.Get(1)
_ = W5037
tmp16619 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W5037)


if True == tmp16619 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp16603 := MakeNative(func(__e *ControlFlow) {
W5038 := __e.Get(1)
_ = W5038
tmp16604 := MakeNative(func(__e *ControlFlow) {
W5039 := __e.Get(1)
_ = W5039
tmp16615 := Call(__e, PrimFunc(symshen_4hds_a_2), W5039, sym_1_6)


if True == tmp16615 {
tmp16605 := MakeNative(func(__e *ControlFlow) {
W5040 := __e.Get(1)
_ = W5040
tmp16612 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W5040)
}
__typedArg0 := W5040
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp16612 {
tmp16606 := MakeNative(func(__e *ControlFlow) {
W5041 := __e.Get(1)
_ = W5041
tmp16607 := MakeNative(func(__e *ControlFlow) {
W5042 := __e.Get(1)
_ = W5042
tmp16608 := Call(__e, PrimFunc(sym_8p), W5038, W5041)


__e.TailApply(PrimFunc(symshen_4comb), W5042, tmp16608)
return


}, 1)

tmp16609 := Call(__e, PrimFunc(symtail), W5040)


__e.TailApply(tmp16607, tmp16609)
return


}, 1)

tmp16610 := Call(__e, PrimFunc(symhead), W5040)


__e.TailApply(tmp16606, tmp16610)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp16613 := Call(__e, PrimFunc(symtail), W5039)


__e.TailApply(tmp16605, tmp16613)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp16616 := Call(__e, PrimFunc(symshen_4in_1_6), W5037)


__e.TailApply(tmp16604, tmp16616)
return


}, 1)

tmp16617 := Call(__e, PrimFunc(symshen_4_5_1out), W5037)


__e.TailApply(tmp16603, tmp16617)
return


}


}, 1)

tmp16620 := Call(__e, PrimFunc(symshen_4_5patterns_6), V5008)


tmp16621 := Call(__e, tmp16602, tmp16620)


__e.TailApply(tmp16599, tmp16621)
return


} else {
__e.Return(W5029)
return
}


}, 1)

tmp16624 := MakeNative(func(__e *ControlFlow) {
W5030 := __e.Get(1)
_ = W5030
tmp16642 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W5030)


if True == tmp16642 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp16625 := MakeNative(func(__e *ControlFlow) {
W5031 := __e.Get(1)
_ = W5031
tmp16626 := MakeNative(func(__e *ControlFlow) {
W5032 := __e.Get(1)
_ = W5032
tmp16638 := Call(__e, PrimFunc(symshen_4hds_a_2), W5032, sym_5_1)


if True == tmp16638 {
tmp16627 := MakeNative(func(__e *ControlFlow) {
W5033 := __e.Get(1)
_ = W5033
tmp16635 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W5033)
}
__typedArg0 := W5033
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp16635 {
tmp16628 := MakeNative(func(__e *ControlFlow) {
W5034 := __e.Get(1)
_ = W5034
tmp16629 := MakeNative(func(__e *ControlFlow) {
W5035 := __e.Get(1)
_ = W5035
tmp16630 := Call(__e, PrimFunc(symshen_4correct), W5034)


tmp16631 := Call(__e, PrimFunc(sym_8p), W5031, tmp16630)


__e.TailApply(PrimFunc(symshen_4comb), W5035, tmp16631)
return


}, 1)

tmp16632 := Call(__e, PrimFunc(symtail), W5033)


__e.TailApply(tmp16629, tmp16632)
return


}, 1)

tmp16633 := Call(__e, PrimFunc(symhead), W5033)


__e.TailApply(tmp16628, tmp16633)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp16636 := Call(__e, PrimFunc(symtail), W5032)


__e.TailApply(tmp16627, tmp16636)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp16639 := Call(__e, PrimFunc(symshen_4in_1_6), W5030)


__e.TailApply(tmp16626, tmp16639)
return


}, 1)

tmp16640 := Call(__e, PrimFunc(symshen_4_5_1out), W5030)


__e.TailApply(tmp16625, tmp16640)
return


}


}, 1)

tmp16643 := Call(__e, PrimFunc(symshen_4_5patterns_6), V5008)


tmp16644 := Call(__e, tmp16624, tmp16643)


__e.TailApply(tmp16598, tmp16644)
return


} else {
__e.Return(W5019)
return
}


}, 1)

tmp16647 := MakeNative(func(__e *ControlFlow) {
W5020 := __e.Get(1)
_ = W5020
tmp16678 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W5020)


if True == tmp16678 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp16648 := MakeNative(func(__e *ControlFlow) {
W5021 := __e.Get(1)
_ = W5021
tmp16649 := MakeNative(func(__e *ControlFlow) {
W5022 := __e.Get(1)
_ = W5022
tmp16674 := Call(__e, PrimFunc(symshen_4hds_a_2), W5022, sym_5_1)


if True == tmp16674 {
tmp16650 := MakeNative(func(__e *ControlFlow) {
W5023 := __e.Get(1)
_ = W5023
tmp16671 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W5023)
}
__typedArg0 := W5023
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp16671 {
tmp16651 := MakeNative(func(__e *ControlFlow) {
W5024 := __e.Get(1)
_ = W5024
tmp16652 := MakeNative(func(__e *ControlFlow) {
W5025 := __e.Get(1)
_ = W5025
tmp16667 := Call(__e, PrimFunc(symshen_4hds_a_2), W5025, symwhere)


if True == tmp16667 {
tmp16653 := MakeNative(func(__e *ControlFlow) {
W5026 := __e.Get(1)
_ = W5026
tmp16664 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W5026)
}
__typedArg0 := W5026
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp16664 {
tmp16654 := MakeNative(func(__e *ControlFlow) {
W5027 := __e.Get(1)
_ = W5027
tmp16655 := MakeNative(func(__e *ControlFlow) {
W5028 := __e.Get(1)
_ = W5028
tmp16656 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W5024, Nil)
}
__typedArg0 := W5024
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16657 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W5027, tmp16656)
}
__typedArg0 := W5027
__typedArg1 := tmp16656
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16658 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhere, tmp16657)
}
__typedArg0 := symwhere
__typedArg1 := tmp16657
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16659 := Call(__e, PrimFunc(symshen_4correct), tmp16658)


tmp16660 := Call(__e, PrimFunc(sym_8p), W5021, tmp16659)


__e.TailApply(PrimFunc(symshen_4comb), W5028, tmp16660)
return


}, 1)

tmp16661 := Call(__e, PrimFunc(symtail), W5026)


__e.TailApply(tmp16655, tmp16661)
return


}, 1)

tmp16662 := Call(__e, PrimFunc(symhead), W5026)


__e.TailApply(tmp16654, tmp16662)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp16665 := Call(__e, PrimFunc(symtail), W5025)


__e.TailApply(tmp16653, tmp16665)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp16668 := Call(__e, PrimFunc(symtail), W5023)


__e.TailApply(tmp16652, tmp16668)
return


}, 1)

tmp16669 := Call(__e, PrimFunc(symhead), W5023)


__e.TailApply(tmp16651, tmp16669)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp16672 := Call(__e, PrimFunc(symtail), W5022)


__e.TailApply(tmp16650, tmp16672)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp16675 := Call(__e, PrimFunc(symshen_4in_1_6), W5020)


__e.TailApply(tmp16649, tmp16675)
return


}, 1)

tmp16676 := Call(__e, PrimFunc(symshen_4_5_1out), W5020)


__e.TailApply(tmp16648, tmp16676)
return


}


}, 1)

tmp16679 := Call(__e, PrimFunc(symshen_4_5patterns_6), V5008)


tmp16680 := Call(__e, tmp16647, tmp16679)


__e.TailApply(tmp16597, tmp16680)
return


} else {
__e.Return(W5009)
return
}


}, 1)

tmp16683 := MakeNative(func(__e *ControlFlow) {
W5010 := __e.Get(1)
_ = W5010
tmp16713 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W5010)


if True == tmp16713 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp16684 := MakeNative(func(__e *ControlFlow) {
W5011 := __e.Get(1)
_ = W5011
tmp16685 := MakeNative(func(__e *ControlFlow) {
W5012 := __e.Get(1)
_ = W5012
tmp16709 := Call(__e, PrimFunc(symshen_4hds_a_2), W5012, sym_1_6)


if True == tmp16709 {
tmp16686 := MakeNative(func(__e *ControlFlow) {
W5013 := __e.Get(1)
_ = W5013
tmp16706 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W5013)
}
__typedArg0 := W5013
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp16706 {
tmp16687 := MakeNative(func(__e *ControlFlow) {
W5014 := __e.Get(1)
_ = W5014
tmp16688 := MakeNative(func(__e *ControlFlow) {
W5015 := __e.Get(1)
_ = W5015
tmp16702 := Call(__e, PrimFunc(symshen_4hds_a_2), W5015, symwhere)


if True == tmp16702 {
tmp16689 := MakeNative(func(__e *ControlFlow) {
W5016 := __e.Get(1)
_ = W5016
tmp16699 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W5016)
}
__typedArg0 := W5016
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp16699 {
tmp16690 := MakeNative(func(__e *ControlFlow) {
W5017 := __e.Get(1)
_ = W5017
tmp16691 := MakeNative(func(__e *ControlFlow) {
W5018 := __e.Get(1)
_ = W5018
tmp16692 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W5014, Nil)
}
__typedArg0 := W5014
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16693 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W5017, tmp16692)
}
__typedArg0 := W5017
__typedArg1 := tmp16692
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16694 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhere, tmp16693)
}
__typedArg0 := symwhere
__typedArg1 := tmp16693
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16695 := Call(__e, PrimFunc(sym_8p), W5011, tmp16694)


__e.TailApply(PrimFunc(symshen_4comb), W5018, tmp16695)
return


}, 1)

tmp16696 := Call(__e, PrimFunc(symtail), W5016)


__e.TailApply(tmp16691, tmp16696)
return


}, 1)

tmp16697 := Call(__e, PrimFunc(symhead), W5016)


__e.TailApply(tmp16690, tmp16697)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp16700 := Call(__e, PrimFunc(symtail), W5015)


__e.TailApply(tmp16689, tmp16700)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp16703 := Call(__e, PrimFunc(symtail), W5013)


__e.TailApply(tmp16688, tmp16703)
return


}, 1)

tmp16704 := Call(__e, PrimFunc(symhead), W5013)


__e.TailApply(tmp16687, tmp16704)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp16707 := Call(__e, PrimFunc(symtail), W5012)


__e.TailApply(tmp16686, tmp16707)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp16710 := Call(__e, PrimFunc(symshen_4in_1_6), W5010)


__e.TailApply(tmp16685, tmp16710)
return


}, 1)

tmp16711 := Call(__e, PrimFunc(symshen_4_5_1out), W5010)


__e.TailApply(tmp16684, tmp16711)
return


}


}, 1)

tmp16714 := Call(__e, PrimFunc(symshen_4_5patterns_6), V5008)


tmp16715 := Call(__e, tmp16683, tmp16714)


__e.TailApply(tmp16596, tmp16715)
return


}, 1)

tmp16716 := Call(__e, ns2_1set, symshen_4_5rule_d_6, tmp16595)


_ = tmp16716

tmp16717 := MakeNative(func(__e *ControlFlow) {
V5043 := __e.Get(1)
_ = V5043
tmp16865 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5043)
}
__typedArg0 := V5043
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16809 Obj

if True == tmp16865 {
tmp16863 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5043)
}
__typedArg0 := V5043
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16864 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symwhere, tmp16863)
}
__typedArg0 := symwhere
__typedArg1 := tmp16863
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16811 Obj

if True == tmp16864 {
tmp16861 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5043)
}
__typedArg0 := V5043
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16862 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp16861)
}
__typedArg0 := tmp16861
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16813 Obj

if True == tmp16862 {
tmp16858 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5043)
}
__typedArg0 := V5043
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16859 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16858)
}
__typedArg0 := tmp16858
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16860 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp16859)
}
__typedArg0 := tmp16859
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16815 Obj

if True == tmp16860 {
tmp16854 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5043)
}
__typedArg0 := V5043
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16855 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16854)
}
__typedArg0 := tmp16854
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16856 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp16855)
}
__typedArg0 := tmp16855
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16857 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp16856)
}
__typedArg0 := tmp16856
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16817 Obj

if True == tmp16857 {
tmp16849 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5043)
}
__typedArg0 := V5043
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16850 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16849)
}
__typedArg0 := tmp16849
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16851 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp16850)
}
__typedArg0 := tmp16850
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16852 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp16851)
}
__typedArg0 := tmp16851
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16853 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symfail_1if, tmp16852)
}
__typedArg0 := symfail_1if
__typedArg1 := tmp16852
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16819 Obj

if True == tmp16853 {
tmp16844 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5043)
}
__typedArg0 := V5043
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16845 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16844)
}
__typedArg0 := tmp16844
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16846 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp16845)
}
__typedArg0 := tmp16845
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16847 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16846)
}
__typedArg0 := tmp16846
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16848 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp16847)
}
__typedArg0 := tmp16847
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16821 Obj

if True == tmp16848 {
tmp16838 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5043)
}
__typedArg0 := V5043
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16839 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16838)
}
__typedArg0 := tmp16838
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16840 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp16839)
}
__typedArg0 := tmp16839
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16841 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16840)
}
__typedArg0 := tmp16840
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16842 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16841)
}
__typedArg0 := tmp16841
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16843 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp16842)
}
__typedArg0 := tmp16842
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16823 Obj

if True == tmp16843 {
tmp16831 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5043)
}
__typedArg0 := V5043
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16832 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16831)
}
__typedArg0 := tmp16831
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16833 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp16832)
}
__typedArg0 := tmp16832
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16834 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16833)
}
__typedArg0 := tmp16833
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16835 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16834)
}
__typedArg0 := tmp16834
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16836 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16835)
}
__typedArg0 := tmp16835
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16837 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp16836)
}
__typedArg0 := Nil
__typedArg1 := tmp16836
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16825 Obj

if True == tmp16837 {
tmp16827 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5043)
}
__typedArg0 := V5043
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16828 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16827)
}
__typedArg0 := tmp16827
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16829 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16828)
}
__typedArg0 := tmp16828
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16830 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp16829)
}
__typedArg0 := Nil
__typedArg1 := tmp16829
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16826 Obj

if True == tmp16830 {
ifres16826 = True


} else {
ifres16826 = False


}

ifres16825 = ifres16826


} else {
ifres16825 = False


}

var ifres16824 Obj

if True == ifres16825 {
ifres16824 = True


} else {
ifres16824 = False


}

ifres16823 = ifres16824


} else {
ifres16823 = False


}

var ifres16822 Obj

if True == ifres16823 {
ifres16822 = True


} else {
ifres16822 = False


}

ifres16821 = ifres16822


} else {
ifres16821 = False


}

var ifres16820 Obj

if True == ifres16821 {
ifres16820 = True


} else {
ifres16820 = False


}

ifres16819 = ifres16820


} else {
ifres16819 = False


}

var ifres16818 Obj

if True == ifres16819 {
ifres16818 = True


} else {
ifres16818 = False


}

ifres16817 = ifres16818


} else {
ifres16817 = False


}

var ifres16816 Obj

if True == ifres16817 {
ifres16816 = True


} else {
ifres16816 = False


}

ifres16815 = ifres16816


} else {
ifres16815 = False


}

var ifres16814 Obj

if True == ifres16815 {
ifres16814 = True


} else {
ifres16814 = False


}

ifres16813 = ifres16814


} else {
ifres16813 = False


}

var ifres16812 Obj

if True == ifres16813 {
ifres16812 = True


} else {
ifres16812 = False


}

ifres16811 = ifres16812


} else {
ifres16811 = False


}

var ifres16810 Obj

if True == ifres16811 {
ifres16810 = True


} else {
ifres16810 = False


}

ifres16809 = ifres16810


} else {
ifres16809 = False


}

if True == ifres16809 {
tmp16718 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5043)
}
__typedArg0 := V5043
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16719 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp16718)
}
__typedArg0 := tmp16718
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16720 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5043)
}
__typedArg0 := V5043
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16721 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16720)
}
__typedArg0 := tmp16720
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16722 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp16721)
}
__typedArg0 := tmp16721
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16723 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16722)
}
__typedArg0 := tmp16722
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16724 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16723, Nil)
}
__typedArg0 := tmp16723
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16725 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnot, tmp16724)
}
__typedArg0 := symnot
__typedArg1 := tmp16724
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16726 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16725, Nil)
}
__typedArg0 := tmp16725
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16727 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16719, tmp16726)
}
__typedArg0 := tmp16719
__typedArg1 := tmp16726
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16728 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symand, tmp16727)
}
__typedArg0 := symand
__typedArg1 := tmp16727
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16729 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5043)
}
__typedArg0 := V5043
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16730 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16729)
}
__typedArg0 := tmp16729
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16731 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp16730)
}
__typedArg0 := tmp16730
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16732 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16731)
}
__typedArg0 := tmp16731
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16733 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16732)
}
__typedArg0 := tmp16732
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16734 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16728, tmp16733)
}
__typedArg0 := tmp16728
__typedArg1 := tmp16733
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhere, tmp16734)
}
__typedArg0 := symwhere
__typedArg1 := tmp16734
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp16807 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5043)
}
__typedArg0 := V5043
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16788 Obj

if True == tmp16807 {
tmp16805 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5043)
}
__typedArg0 := V5043
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16806 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symwhere, tmp16805)
}
__typedArg0 := symwhere
__typedArg1 := tmp16805
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16790 Obj

if True == tmp16806 {
tmp16803 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5043)
}
__typedArg0 := V5043
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16804 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp16803)
}
__typedArg0 := tmp16803
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16792 Obj

if True == tmp16804 {
tmp16800 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5043)
}
__typedArg0 := V5043
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16801 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16800)
}
__typedArg0 := tmp16800
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16802 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp16801)
}
__typedArg0 := tmp16801
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16794 Obj

if True == tmp16802 {
tmp16796 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5043)
}
__typedArg0 := V5043
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16797 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16796)
}
__typedArg0 := tmp16796
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16798 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16797)
}
__typedArg0 := tmp16797
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16799 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp16798)
}
__typedArg0 := Nil
__typedArg1 := tmp16798
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16795 Obj

if True == tmp16799 {
ifres16795 = True


} else {
ifres16795 = False


}

ifres16794 = ifres16795


} else {
ifres16794 = False


}

var ifres16793 Obj

if True == ifres16794 {
ifres16793 = True


} else {
ifres16793 = False


}

ifres16792 = ifres16793


} else {
ifres16792 = False


}

var ifres16791 Obj

if True == ifres16792 {
ifres16791 = True


} else {
ifres16791 = False


}

ifres16790 = ifres16791


} else {
ifres16790 = False


}

var ifres16789 Obj

if True == ifres16790 {
ifres16789 = True


} else {
ifres16789 = False


}

ifres16788 = ifres16789


} else {
ifres16788 = False


}

if True == ifres16788 {
tmp16735 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5043)
}
__typedArg0 := V5043
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16736 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp16735)
}
__typedArg0 := tmp16735
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16737 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5043)
}
__typedArg0 := V5043
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16738 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16737)
}
__typedArg0 := tmp16737
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16739 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp16738)
}
__typedArg0 := tmp16738
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16740 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfail, Nil)
}
__typedArg0 := symfail
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16741 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16740, Nil)
}
__typedArg0 := tmp16740
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16742 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16739, tmp16741)
}
__typedArg0 := tmp16739
__typedArg1 := tmp16741
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16743 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_a_a, tmp16742)
}
__typedArg0 := sym_a_a
__typedArg1 := tmp16742
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16744 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16743, Nil)
}
__typedArg0 := tmp16743
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16745 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnot, tmp16744)
}
__typedArg0 := symnot
__typedArg1 := tmp16744
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16746 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16745, Nil)
}
__typedArg0 := tmp16745
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16747 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16736, tmp16746)
}
__typedArg0 := tmp16736
__typedArg1 := tmp16746
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16748 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symand, tmp16747)
}
__typedArg0 := symand
__typedArg1 := tmp16747
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16749 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5043)
}
__typedArg0 := V5043
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16750 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16749)
}
__typedArg0 := tmp16749
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16751 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16748, tmp16750)
}
__typedArg0 := tmp16748
__typedArg1 := tmp16750
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhere, tmp16751)
}
__typedArg0 := symwhere
__typedArg1 := tmp16751
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp16786 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5043)
}
__typedArg0 := V5043
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16767 Obj

if True == tmp16786 {
tmp16784 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5043)
}
__typedArg0 := V5043
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16785 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symfail_1if, tmp16784)
}
__typedArg0 := symfail_1if
__typedArg1 := tmp16784
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16769 Obj

if True == tmp16785 {
tmp16782 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5043)
}
__typedArg0 := V5043
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16783 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp16782)
}
__typedArg0 := tmp16782
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16771 Obj

if True == tmp16783 {
tmp16779 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5043)
}
__typedArg0 := V5043
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16780 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16779)
}
__typedArg0 := tmp16779
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16781 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp16780)
}
__typedArg0 := tmp16780
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16773 Obj

if True == tmp16781 {
tmp16775 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5043)
}
__typedArg0 := V5043
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16776 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16775)
}
__typedArg0 := tmp16775
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16777 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16776)
}
__typedArg0 := tmp16776
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16778 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp16777)
}
__typedArg0 := Nil
__typedArg1 := tmp16777
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16774 Obj

if True == tmp16778 {
ifres16774 = True


} else {
ifres16774 = False


}

ifres16773 = ifres16774


} else {
ifres16773 = False


}

var ifres16772 Obj

if True == ifres16773 {
ifres16772 = True


} else {
ifres16772 = False


}

ifres16771 = ifres16772


} else {
ifres16771 = False


}

var ifres16770 Obj

if True == ifres16771 {
ifres16770 = True


} else {
ifres16770 = False


}

ifres16769 = ifres16770


} else {
ifres16769 = False


}

var ifres16768 Obj

if True == ifres16769 {
ifres16768 = True


} else {
ifres16768 = False


}

ifres16767 = ifres16768


} else {
ifres16767 = False


}

if True == ifres16767 {
tmp16752 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5043)
}
__typedArg0 := V5043
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16753 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16752, Nil)
}
__typedArg0 := tmp16752
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16754 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnot, tmp16753)
}
__typedArg0 := symnot
__typedArg1 := tmp16753
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16755 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5043)
}
__typedArg0 := V5043
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16756 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16755)
}
__typedArg0 := tmp16755
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16757 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16754, tmp16756)
}
__typedArg0 := tmp16754
__typedArg1 := tmp16756
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhere, tmp16757)
}
__typedArg0 := symwhere
__typedArg1 := tmp16757
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp16758 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfail, Nil)
}
__typedArg0 := symfail
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16759 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16758, Nil)
}
__typedArg0 := tmp16758
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16760 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5043, tmp16759)
}
__typedArg0 := V5043
__typedArg1 := tmp16759
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16761 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_a_a, tmp16760)
}
__typedArg0 := sym_a_a
__typedArg1 := tmp16760
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16762 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16761, Nil)
}
__typedArg0 := tmp16761
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16763 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnot, tmp16762)
}
__typedArg0 := symnot
__typedArg1 := tmp16762
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16764 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5043, Nil)
}
__typedArg0 := V5043
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16765 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16763, tmp16764)
}
__typedArg0 := tmp16763
__typedArg1 := tmp16764
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhere, tmp16765)
}
__typedArg0 := symwhere
__typedArg1 := tmp16765
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}


}


}


}, 1)

tmp16866 := Call(__e, ns2_1set, symshen_4correct, tmp16717)


_ = tmp16866

tmp16867 := MakeNative(func(__e *ControlFlow) {
V5044 := __e.Get(1)
_ = V5044
V5045 := __e.Get(2)
_ = V5045
V5046 := __e.Get(3)
_ = V5046
V5047 := __e.Get(4)
_ = V5047
V5048 := __e.Get(5)
_ = V5048
V5049 := __e.Get(6)
_ = V5049
V5050 := __e.Get(7)
_ = V5050
V5051 := __e.Get(8)
_ = V5051
tmp16868 := MakeNative(func(__e *ControlFlow) {
W5052 := __e.Get(1)
_ = W5052
tmp16869 := MakeNative(func(__e *ControlFlow) {
W5053 := __e.Get(1)
_ = W5053
tmp16899 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W5053, False)
}
__typedArg0 := W5053
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16899 {
tmp16870 := MakeNative(func(__e *ControlFlow) {
W5055 := __e.Get(1)
_ = W5055
tmp16872 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W5055, False)
}
__typedArg0 := W5055
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16872 {
__e.TailApply(PrimFunc(symshen_4unlock), V5049, W5052)
return
} else {
__e.Return(W5055)
return
}


}, 1)

tmp16897 := Call(__e, PrimFunc(symshen_4unlocked_2), V5049)


var ifres16873 Obj

if True == tmp16897 {
tmp16874 := MakeNative(func(__e *ControlFlow) {
W5056 := __e.Get(1)
_ = W5056
tmp16894 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W5056)
}
__typedArg0 := W5056
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp16894 {
tmp16875 := MakeNative(func(__e *ControlFlow) {
W5057 := __e.Get(1)
_ = W5057
tmp16876 := MakeNative(func(__e *ControlFlow) {
W5058 := __e.Get(1)
_ = W5058
tmp16877 := MakeNative(func(__e *ControlFlow) {
W5059 := __e.Get(1)
_ = W5059
tmp16878 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16878

tmp16879 := Call(__e, PrimFunc(symshen_4deref), W5057, V5048)


tmp16880 := Call(__e, PrimFunc(symshen_4freshen_1rule), tmp16879)


tmp16881 := MakeNative(func(__e *ControlFlow) {
tmp16882 := Call(__e, PrimFunc(symshen_4lazyderef), W5059, V5048)


tmp16883 := Call(__e, PrimFunc(symfst), tmp16882)


tmp16884 := Call(__e, PrimFunc(symshen_4lazyderef), W5059, V5048)


tmp16885 := Call(__e, PrimFunc(symsnd), tmp16884)


tmp16886 := MakeNative(func(__e *ControlFlow) {
tmp16887 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4t_d_1rules), V5044, W5058, V5046, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V5047)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V5047
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})(), V5048, V5049, W5052, V5051)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V5048, V5049, W5052, tmp16887)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4t_d_1rule), V5044, V5047, tmp16883, tmp16885, V5046, V5048, V5049, W5052, tmp16886)
return


}, 0)

tmp16889 := Call(__e, PrimFunc(symbind), W5059, tmp16880, V5048, V5049, W5052, tmp16881)


__e.TailApply(PrimFunc(symshen_4gc), V5048, tmp16889)
return


}, 1)

tmp16890 := Call(__e, PrimFunc(symshen_4newpv), V5048)


__e.TailApply(tmp16877, tmp16890)
return


}, 1)

tmp16891 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W5056)
}
__typedArg0 := W5056
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(tmp16876, tmp16891)
return


}, 1)

tmp16892 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W5056)
}
__typedArg0 := W5056
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp16875, tmp16892)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16895 := Call(__e, PrimFunc(symshen_4lazyderef), V5045, V5048)


tmp16896 := Call(__e, tmp16874, tmp16895)


ifres16873 = tmp16896


} else {
ifres16873 = False


}

__e.TailApply(tmp16870, ifres16873)
return


} else {
__e.Return(W5053)
return
}


}, 1)

tmp16907 := Call(__e, PrimFunc(symshen_4unlocked_2), V5049)


var ifres16900 Obj

if True == tmp16907 {
tmp16901 := MakeNative(func(__e *ControlFlow) {
W5054 := __e.Get(1)
_ = W5054
tmp16904 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W5054, Nil)
}
__typedArg0 := W5054
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16904 {
tmp16902 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16902

__e.TailApply(PrimFunc(symthaw), V5051)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16905 := Call(__e, PrimFunc(symshen_4lazyderef), V5045, V5048)


tmp16906 := Call(__e, tmp16901, tmp16905)


ifres16900 = tmp16906


} else {
ifres16900 = False


}

__e.TailApply(tmp16869, ifres16900)
return


}, 1)

__e.TailApply(tmp16868, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V5050)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V5050
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())
return


}, 8)

tmp16909 := Call(__e, ns2_1set, symshen_4t_d_1rules, tmp16867)


_ = tmp16909

tmp16910 := MakeNative(func(__e *ControlFlow) {
V5060 := __e.Get(1)
_ = V5060
tmp16923 := Call(__e, PrimFunc(symtuple_2), V5060)


if True == tmp16923 {
tmp16911 := MakeNative(func(__e *ControlFlow) {
W5061 := __e.Get(1)
_ = W5061
tmp16912 := MakeNative(func(__e *ControlFlow) {
W5062 := __e.Get(1)
_ = W5062
tmp16913 := Call(__e, PrimFunc(symfst), V5060)


tmp16914 := Call(__e, PrimFunc(symshen_4freshen), W5062, tmp16913)


tmp16915 := Call(__e, PrimFunc(symsnd), V5060)


tmp16916 := Call(__e, PrimFunc(symshen_4freshen), W5062, tmp16915)


__e.TailApply(PrimFunc(sym_8p), tmp16914, tmp16916)
return


}, 1)

tmp16917 := MakeNative(func(__e *ControlFlow) {
Z5063 := __e.Get(1)
_ = Z5063
tmp16918 := Call(__e, PrimFunc(symshen_4freshterm), Z5063)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Z5063, tmp16918)
}
__typedArg0 := Z5063
__typedArg1 := tmp16918
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp16919 := Call(__e, PrimFunc(symmap), tmp16917, W5061)


__e.TailApply(tmp16912, tmp16919)
return


}, 1)

tmp16920 := Call(__e, PrimFunc(symfst), V5060)


tmp16921 := Call(__e, PrimFunc(symshen_4extract_1vars), tmp16920)


__e.TailApply(tmp16911, tmp16921)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("partial function shen.freshen-rule"))
}
__typedArg0 := MakeString("partial function shen.freshen-rule")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}, 1)

tmp16924 := Call(__e, ns2_1set, symshen_4freshen_1rule, tmp16910)


_ = tmp16924

tmp16925 := MakeNative(func(__e *ControlFlow) {
V5064 := __e.Get(1)
_ = V5064
V5065 := __e.Get(2)
_ = V5065
tmp16939 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V5064)
}
__typedArg0 := Nil
__typedArg1 := V5064
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16939 {
__e.Return(V5065)
return
} else {
tmp16937 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5064)
}
__typedArg0 := V5064
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16933 Obj

if True == tmp16937 {
tmp16935 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5064)
}
__typedArg0 := V5064
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16936 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp16935)
}
__typedArg0 := tmp16935
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16934 Obj

if True == tmp16936 {
ifres16934 = True


} else {
ifres16934 = False


}

ifres16933 = ifres16934


} else {
ifres16933 = False


}

if True == ifres16933 {
tmp16926 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5064)
}
__typedArg0 := V5064
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16927 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5064)
}
__typedArg0 := V5064
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16928 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp16927)
}
__typedArg0 := tmp16927
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16929 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5064)
}
__typedArg0 := V5064
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16930 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16929)
}
__typedArg0 := tmp16929
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16931 := Call(__e, PrimFunc(symshen_4beta), tmp16928, tmp16930, V5065)


__e.TailApply(PrimFunc(symshen_4freshen), tmp16926, tmp16931)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("partial function shen.freshen"))
}
__typedArg0 := MakeString("partial function shen.freshen")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 2)

tmp16940 := Call(__e, ns2_1set, symshen_4freshen, tmp16925)


_ = tmp16940

tmp16941 := MakeNative(func(__e *ControlFlow) {
V5066 := __e.Get(1)
_ = V5066
V5067 := __e.Get(2)
_ = V5067
V5068 := __e.Get(3)
_ = V5068
V5069 := __e.Get(4)
_ = V5069
V5070 := __e.Get(5)
_ = V5070
V5071 := __e.Get(6)
_ = V5071
V5072 := __e.Get(7)
_ = V5072
V5073 := __e.Get(8)
_ = V5073
V5074 := __e.Get(9)
_ = V5074
tmp16942 := MakeNative(func(__e *ControlFlow) {
W5075 := __e.Get(1)
_ = W5075
tmp16955 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W5075, False)
}
__typedArg0 := W5075
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16955 {
tmp16953 := Call(__e, PrimFunc(symshen_4unlocked_2), V5072)


if True == tmp16953 {
tmp16943 := MakeNative(func(__e *ControlFlow) {
W5076 := __e.Get(1)
_ = W5076
tmp16944 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16944

tmp16945 := Call(__e, PrimFunc(symshen_4app), V5066, MakeString("\n"), symshen_4a)


tmp16947 := Call(__e, PrimFunc(symshen_4app), V5067, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(" of "))
__typedS1, __typedOK1 := TypedString(tmp16945)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(" of ")
__typedArg1 := tmp16945
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


tmp16949 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("type error in rule "))
__typedS1, __typedOK1 := TypedString(tmp16947)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("type error in rule ")
__typedArg1 := tmp16947
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("type error in rule "))
__typedS1, __typedOK1 := TypedString(tmp16947)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("type error in rule ")
__typedArg1 := tmp16947
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})()

tmp16950 := Call(__e, PrimFunc(symbind), W5076, tmp16949, V5071, V5072, V5073, V5074)


__e.TailApply(PrimFunc(symshen_4gc), V5071, tmp16950)
return


}, 1)

tmp16951 := Call(__e, PrimFunc(symshen_4newpv), V5071)


__e.TailApply(tmp16943, tmp16951)
return


} else {
__e.Return(False)
return
}


} else {
__e.Return(W5075)
return
}


}, 1)

tmp16959 := Call(__e, PrimFunc(symshen_4unlocked_2), V5072)


var ifres16956 Obj

if True == tmp16959 {
tmp16957 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16957

tmp16958 := Call(__e, PrimFunc(symshen_4t_d_1rule_1h), V5068, V5069, V5070, V5071, V5072, V5073, V5074)


ifres16956 = tmp16958


} else {
ifres16956 = False


}

__e.TailApply(tmp16942, ifres16956)
return


}, 9)

tmp16960 := Call(__e, ns2_1set, symshen_4t_d_1rule, tmp16941)


_ = tmp16960

tmp16961 := MakeNative(func(__e *ControlFlow) {
V5077 := __e.Get(1)
_ = V5077
V5078 := __e.Get(2)
_ = V5078
V5079 := __e.Get(3)
_ = V5079
V5080 := __e.Get(4)
_ = V5080
V5081 := __e.Get(5)
_ = V5081
V5082 := __e.Get(6)
_ = V5082
V5083 := __e.Get(7)
_ = V5083
tmp16962 := MakeNative(func(__e *ControlFlow) {
W5084 := __e.Get(1)
_ = W5084
tmp16963 := MakeNative(func(__e *ControlFlow) {
W5085 := __e.Get(1)
_ = W5085
tmp16986 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W5085, False)
}
__typedArg0 := W5085
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16986 {
tmp16964 := MakeNative(func(__e *ControlFlow) {
W5092 := __e.Get(1)
_ = W5092
tmp16966 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W5092, False)
}
__typedArg0 := W5092
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16966 {
__e.TailApply(PrimFunc(symshen_4unlock), V5081, W5084)
return
} else {
__e.Return(W5092)
return
}


}, 1)

tmp16984 := Call(__e, PrimFunc(symshen_4unlocked_2), V5081)


var ifres16967 Obj

if True == tmp16984 {
tmp16968 := MakeNative(func(__e *ControlFlow) {
W5093 := __e.Get(1)
_ = W5093
tmp16969 := MakeNative(func(__e *ControlFlow) {
W5094 := __e.Get(1)
_ = W5094
tmp16970 := MakeNative(func(__e *ControlFlow) {
W5095 := __e.Get(1)
_ = W5095
tmp16971 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16971

tmp16972 := Call(__e, PrimFunc(symshen_4freshterms), V5077)


tmp16973 := MakeNative(func(__e *ControlFlow) {
tmp16974 := MakeNative(func(__e *ControlFlow) {
tmp16975 := MakeNative(func(__e *ControlFlow) {
tmp16976 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4t_d_1correct), V5078, W5094, W5095, V5080, V5081, W5084, V5083)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4myassume), V5077, V5079, W5095, V5080, V5081, W5084, tmp16976)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V5080, V5081, W5084, tmp16975)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4t_d_1integrity), V5077, V5079, W5093, W5094, V5080, V5081, W5084, tmp16974)
return


}, 0)

tmp16977 := Call(__e, PrimFunc(symshen_4p_1hyps), tmp16972, W5093, V5080, V5081, W5084, tmp16973)


__e.TailApply(PrimFunc(symshen_4gc), V5080, tmp16977)
return


}, 1)

tmp16978 := Call(__e, PrimFunc(symshen_4newpv), V5080)


tmp16979 := Call(__e, tmp16970, tmp16978)


__e.TailApply(PrimFunc(symshen_4gc), V5080, tmp16979)
return


}, 1)

tmp16980 := Call(__e, PrimFunc(symshen_4newpv), V5080)


tmp16981 := Call(__e, tmp16969, tmp16980)


__e.TailApply(PrimFunc(symshen_4gc), V5080, tmp16981)
return


}, 1)

tmp16982 := Call(__e, PrimFunc(symshen_4newpv), V5080)


tmp16983 := Call(__e, tmp16968, tmp16982)


ifres16967 = tmp16983


} else {
ifres16967 = False


}

__e.TailApply(tmp16964, ifres16967)
return


} else {
__e.Return(W5085)
return
}


}, 1)

tmp17016 := Call(__e, PrimFunc(symshen_4unlocked_2), V5081)


var ifres16987 Obj

if True == tmp17016 {
tmp16988 := MakeNative(func(__e *ControlFlow) {
W5086 := __e.Get(1)
_ = W5086
tmp17013 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W5086, Nil)
}
__typedArg0 := W5086
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17013 {
tmp16989 := MakeNative(func(__e *ControlFlow) {
W5087 := __e.Get(1)
_ = W5087
tmp17010 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W5087)
}
__typedArg0 := W5087
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17010 {
tmp16990 := MakeNative(func(__e *ControlFlow) {
W5088 := __e.Get(1)
_ = W5088
tmp17006 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W5088, sym_1_1_6)
}
__typedArg0 := W5088
__typedArg1 := sym_1_1_6
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17006 {
tmp16991 := MakeNative(func(__e *ControlFlow) {
W5089 := __e.Get(1)
_ = W5089
tmp17002 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W5089)
}
__typedArg0 := W5089
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17002 {
tmp16992 := MakeNative(func(__e *ControlFlow) {
W5090 := __e.Get(1)
_ = W5090
tmp16993 := MakeNative(func(__e *ControlFlow) {
W5091 := __e.Get(1)
_ = W5091
tmp16997 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W5091, Nil)
}
__typedArg0 := W5091
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16997 {
tmp16994 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16994

tmp16995 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4t_d_1correct), V5078, W5090, Nil, V5080, V5081, W5084, V5083)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V5080, V5081, W5084, tmp16995)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp16998 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W5089)
}
__typedArg0 := W5089
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16999 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16998, V5080)


__e.TailApply(tmp16993, tmp16999)
return


}, 1)

tmp17000 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W5089)
}
__typedArg0 := W5089
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp16992, tmp17000)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp17003 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W5087)
}
__typedArg0 := W5087
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17004 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17003, V5080)


__e.TailApply(tmp16991, tmp17004)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp17007 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W5087)
}
__typedArg0 := W5087
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17008 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17007, V5080)


__e.TailApply(tmp16990, tmp17008)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp17011 := Call(__e, PrimFunc(symshen_4lazyderef), V5079, V5080)


__e.TailApply(tmp16989, tmp17011)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp17014 := Call(__e, PrimFunc(symshen_4lazyderef), V5077, V5080)


tmp17015 := Call(__e, tmp16988, tmp17014)


ifres16987 = tmp17015


} else {
ifres16987 = False


}

__e.TailApply(tmp16963, ifres16987)
return


}, 1)

__e.TailApply(tmp16962, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V5082)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V5082
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())
return


}, 7)

tmp17018 := Call(__e, ns2_1set, symshen_4t_d_1rule_1h, tmp16961)


_ = tmp17018

tmp17019 := MakeNative(func(__e *ControlFlow) {
V5096 := __e.Get(1)
_ = V5096
V5097 := __e.Get(2)
_ = V5097
V5098 := __e.Get(3)
_ = V5098
V5099 := __e.Get(4)
_ = V5099
V5100 := __e.Get(5)
_ = V5100
V5101 := __e.Get(6)
_ = V5101
V5102 := __e.Get(7)
_ = V5102
tmp17020 := MakeNative(func(__e *ControlFlow) {
W5103 := __e.Get(1)
_ = W5103
tmp17173 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W5103, False)
}
__typedArg0 := W5103
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17173 {
tmp17171 := Call(__e, PrimFunc(symshen_4unlocked_2), V5100)


if True == tmp17171 {
tmp17021 := MakeNative(func(__e *ControlFlow) {
W5107 := __e.Get(1)
_ = W5107
tmp17168 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W5107)
}
__typedArg0 := W5107
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17168 {
tmp17022 := MakeNative(func(__e *ControlFlow) {
W5108 := __e.Get(1)
_ = W5108
tmp17023 := MakeNative(func(__e *ControlFlow) {
W5109 := __e.Get(1)
_ = W5109
tmp17024 := MakeNative(func(__e *ControlFlow) {
W5110 := __e.Get(1)
_ = W5110
tmp17163 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W5110)
}
__typedArg0 := W5110
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17163 {
tmp17025 := MakeNative(func(__e *ControlFlow) {
W5111 := __e.Get(1)
_ = W5111
tmp17026 := MakeNative(func(__e *ControlFlow) {
W5112 := __e.Get(1)
_ = W5112
tmp17158 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W5112)
}
__typedArg0 := W5112
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17158 {
tmp17027 := MakeNative(func(__e *ControlFlow) {
W5113 := __e.Get(1)
_ = W5113
tmp17154 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W5113, sym_1_1_6)
}
__typedArg0 := W5113
__typedArg1 := sym_1_1_6
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17154 {
tmp17028 := MakeNative(func(__e *ControlFlow) {
W5114 := __e.Get(1)
_ = W5114
tmp17150 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W5114)
}
__typedArg0 := W5114
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17150 {
tmp17029 := MakeNative(func(__e *ControlFlow) {
W5115 := __e.Get(1)
_ = W5115
tmp17030 := MakeNative(func(__e *ControlFlow) {
W5116 := __e.Get(1)
_ = W5116
tmp17145 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W5116, Nil)
}
__typedArg0 := W5116
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17145 {
tmp17031 := MakeNative(func(__e *ControlFlow) {
W5117 := __e.Get(1)
_ = W5117
tmp17032 := MakeNative(func(__e *ControlFlow) {
W5118 := __e.Get(1)
_ = W5118
tmp17136 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W5117)
}
__typedArg0 := W5117
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17136 {
tmp17033 := MakeNative(func(__e *ControlFlow) {
W5123 := __e.Get(1)
_ = W5123
tmp17034 := MakeNative(func(__e *ControlFlow) {
W5124 := __e.Get(1)
_ = W5124
tmp17104 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W5123)
}
__typedArg0 := W5123
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17104 {
tmp17035 := MakeNative(func(__e *ControlFlow) {
W5129 := __e.Get(1)
_ = W5129
tmp17036 := MakeNative(func(__e *ControlFlow) {
W5130 := __e.Get(1)
_ = W5130
tmp17037 := MakeNative(func(__e *ControlFlow) {
W5131 := __e.Get(1)
_ = W5131
tmp17079 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W5130)
}
__typedArg0 := W5130
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17079 {
tmp17038 := MakeNative(func(__e *ControlFlow) {
W5134 := __e.Get(1)
_ = W5134
tmp17039 := MakeNative(func(__e *ControlFlow) {
W5135 := __e.Get(1)
_ = W5135
tmp17040 := MakeNative(func(__e *ControlFlow) {
W5136 := __e.Get(1)
_ = W5136
tmp17060 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W5135)
}
__typedArg0 := W5135
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17060 {
tmp17041 := MakeNative(func(__e *ControlFlow) {
W5138 := __e.Get(1)
_ = W5138
tmp17042 := MakeNative(func(__e *ControlFlow) {
W5139 := __e.Get(1)
_ = W5139
tmp17043 := MakeNative(func(__e *ControlFlow) {
W5140 := __e.Get(1)
_ = W5140
tmp17047 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W5139, Nil)
}
__typedArg0 := W5139
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17047 {
__e.TailApply(PrimFunc(symthaw), W5140)
return
} else {
tmp17045 := Call(__e, PrimFunc(symshen_4pvar_2), W5139)


if True == tmp17045 {
__e.TailApply(PrimFunc(symshen_4bind_b), W5139, Nil, V5099, W5140)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp17048 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W5136, W5138)
return
}, 0)

__e.TailApply(tmp17043, tmp17048)
return


}, 1)

tmp17049 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W5135)
}
__typedArg0 := W5135
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17050 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17049, V5099)


__e.TailApply(tmp17042, tmp17050)
return


}, 1)

tmp17051 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W5135)
}
__typedArg0 := W5135
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp17041, tmp17051)
return


} else {
tmp17058 := Call(__e, PrimFunc(symshen_4pvar_2), W5135)


if True == tmp17058 {
tmp17052 := MakeNative(func(__e *ControlFlow) {
W5141 := __e.Get(1)
_ = W5141
tmp17053 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W5141, Nil)
}
__typedArg0 := W5141
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17054 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W5136, W5141)
return
}, 0)

tmp17055 := Call(__e, PrimFunc(symshen_4bind_b), W5135, tmp17053, V5099, tmp17054)


__e.TailApply(PrimFunc(symshen_4gc), V5099, tmp17055)
return


}, 1)

tmp17056 := Call(__e, PrimFunc(symshen_4newpv), V5099)


__e.TailApply(tmp17052, tmp17056)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp17061 := MakeNative(func(__e *ControlFlow) {
Z5137 := __e.Get(1)
_ = Z5137
tmp17062 := Call(__e, W5131, W5134)


__e.TailApply(tmp17062, Z5137)
return


}, 1)

__e.TailApply(tmp17040, tmp17061)
return


}, 1)

tmp17063 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W5130)
}
__typedArg0 := W5130
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17064 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17063, V5099)


__e.TailApply(tmp17039, tmp17064)
return


}, 1)

tmp17065 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W5130)
}
__typedArg0 := W5130
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp17038, tmp17065)
return


} else {
tmp17077 := Call(__e, PrimFunc(symshen_4pvar_2), W5130)


if True == tmp17077 {
tmp17066 := MakeNative(func(__e *ControlFlow) {
W5142 := __e.Get(1)
_ = W5142
tmp17067 := MakeNative(func(__e *ControlFlow) {
W5143 := __e.Get(1)
_ = W5143
tmp17068 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W5143, Nil)
}
__typedArg0 := W5143
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17069 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W5142, tmp17068)
}
__typedArg0 := W5142
__typedArg1 := tmp17068
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17070 := MakeNative(func(__e *ControlFlow) {
tmp17071 := Call(__e, W5131, W5142)


__e.TailApply(tmp17071, W5143)
return


}, 0)

tmp17072 := Call(__e, PrimFunc(symshen_4bind_b), W5130, tmp17069, V5099, tmp17070)


__e.TailApply(PrimFunc(symshen_4gc), V5099, tmp17072)
return


}, 1)

tmp17073 := Call(__e, PrimFunc(symshen_4newpv), V5099)


tmp17074 := Call(__e, tmp17067, tmp17073)


__e.TailApply(PrimFunc(symshen_4gc), V5099, tmp17074)
return


}, 1)

tmp17075 := Call(__e, PrimFunc(symshen_4newpv), V5099)


__e.TailApply(tmp17066, tmp17075)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp17080 := MakeNative(func(__e *ControlFlow) {
Z5132 := __e.Get(1)
_ = Z5132
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5133 := __e.Get(1)
_ = Z5133
tmp17081 := Call(__e, W5124, W5129)


tmp17082 := Call(__e, tmp17081, Z5132)


__e.TailApply(tmp17082, Z5133)
return


}, 1))
return
}, 1)

__e.TailApply(tmp17037, tmp17080)
return


}, 1)

tmp17083 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W5123)
}
__typedArg0 := W5123
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17084 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17083, V5099)


__e.TailApply(tmp17036, tmp17084)
return


}, 1)

tmp17085 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W5123)
}
__typedArg0 := W5123
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp17035, tmp17085)
return


} else {
tmp17102 := Call(__e, PrimFunc(symshen_4pvar_2), W5123)


if True == tmp17102 {
tmp17086 := MakeNative(func(__e *ControlFlow) {
W5144 := __e.Get(1)
_ = W5144
tmp17087 := MakeNative(func(__e *ControlFlow) {
W5145 := __e.Get(1)
_ = W5145
tmp17088 := MakeNative(func(__e *ControlFlow) {
W5146 := __e.Get(1)
_ = W5146
tmp17089 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W5146, Nil)
}
__typedArg0 := W5146
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17090 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W5145, tmp17089)
}
__typedArg0 := W5145
__typedArg1 := tmp17089
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17091 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W5144, tmp17090)
}
__typedArg0 := W5144
__typedArg1 := tmp17090
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17092 := MakeNative(func(__e *ControlFlow) {
tmp17093 := Call(__e, W5124, W5144)


tmp17094 := Call(__e, tmp17093, W5145)


__e.TailApply(tmp17094, W5146)
return


}, 0)

tmp17095 := Call(__e, PrimFunc(symshen_4bind_b), W5123, tmp17091, V5099, tmp17092)


__e.TailApply(PrimFunc(symshen_4gc), V5099, tmp17095)
return


}, 1)

tmp17096 := Call(__e, PrimFunc(symshen_4newpv), V5099)


tmp17097 := Call(__e, tmp17088, tmp17096)


__e.TailApply(PrimFunc(symshen_4gc), V5099, tmp17097)
return


}, 1)

tmp17098 := Call(__e, PrimFunc(symshen_4newpv), V5099)


tmp17099 := Call(__e, tmp17087, tmp17098)


__e.TailApply(PrimFunc(symshen_4gc), V5099, tmp17099)
return


}, 1)

tmp17100 := Call(__e, PrimFunc(symshen_4newpv), V5099)


__e.TailApply(tmp17086, tmp17100)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp17105 := MakeNative(func(__e *ControlFlow) {
Z5125 := __e.Get(1)
_ = Z5125
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5126 := __e.Get(1)
_ = Z5126
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5127 := __e.Get(1)
_ = Z5127
tmp17106 := MakeNative(func(__e *ControlFlow) {
W5128 := __e.Get(1)
_ = W5128
tmp17107 := Call(__e, W5118, Z5125)


tmp17108 := Call(__e, tmp17107, Z5126)


tmp17109 := Call(__e, tmp17108, Z5127)


__e.TailApply(tmp17109, W5128)
return


}, 1)

tmp17110 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W5117)
}
__typedArg0 := W5117
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(tmp17106, tmp17110)
return


}, 1))
return
}, 1))
return
}, 1)

__e.TailApply(tmp17034, tmp17105)
return


}, 1)

tmp17111 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W5117)
}
__typedArg0 := W5117
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17112 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17111, V5099)


__e.TailApply(tmp17033, tmp17112)
return


} else {
tmp17134 := Call(__e, PrimFunc(symshen_4pvar_2), W5117)


if True == tmp17134 {
tmp17113 := MakeNative(func(__e *ControlFlow) {
W5147 := __e.Get(1)
_ = W5147
tmp17114 := MakeNative(func(__e *ControlFlow) {
W5148 := __e.Get(1)
_ = W5148
tmp17115 := MakeNative(func(__e *ControlFlow) {
W5149 := __e.Get(1)
_ = W5149
tmp17116 := MakeNative(func(__e *ControlFlow) {
W5150 := __e.Get(1)
_ = W5150
tmp17117 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W5149, Nil)
}
__typedArg0 := W5149
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17118 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W5148, tmp17117)
}
__typedArg0 := W5148
__typedArg1 := tmp17117
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17119 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W5147, tmp17118)
}
__typedArg0 := W5147
__typedArg1 := tmp17118
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17120 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17119, W5150)
}
__typedArg0 := tmp17119
__typedArg1 := W5150
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17121 := MakeNative(func(__e *ControlFlow) {
tmp17122 := Call(__e, W5118, W5147)


tmp17123 := Call(__e, tmp17122, W5148)


tmp17124 := Call(__e, tmp17123, W5149)


__e.TailApply(tmp17124, W5150)
return


}, 0)

tmp17125 := Call(__e, PrimFunc(symshen_4bind_b), W5117, tmp17120, V5099, tmp17121)


__e.TailApply(PrimFunc(symshen_4gc), V5099, tmp17125)
return


}, 1)

tmp17126 := Call(__e, PrimFunc(symshen_4newpv), V5099)


tmp17127 := Call(__e, tmp17116, tmp17126)


__e.TailApply(PrimFunc(symshen_4gc), V5099, tmp17127)
return


}, 1)

tmp17128 := Call(__e, PrimFunc(symshen_4newpv), V5099)


tmp17129 := Call(__e, tmp17115, tmp17128)


__e.TailApply(PrimFunc(symshen_4gc), V5099, tmp17129)
return


}, 1)

tmp17130 := Call(__e, PrimFunc(symshen_4newpv), V5099)


tmp17131 := Call(__e, tmp17114, tmp17130)


__e.TailApply(PrimFunc(symshen_4gc), V5099, tmp17131)
return


}, 1)

tmp17132 := Call(__e, PrimFunc(symshen_4newpv), V5099)


__e.TailApply(tmp17113, tmp17132)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp17137 := MakeNative(func(__e *ControlFlow) {
Z5119 := __e.Get(1)
_ = Z5119
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5120 := __e.Get(1)
_ = Z5120
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5121 := __e.Get(1)
_ = Z5121
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5122 := __e.Get(1)
_ = Z5122
tmp17138 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp17138

tmp17139 := MakeNative(func(__e *ControlFlow) {
tmp17140 := MakeNative(func(__e *ControlFlow) {
tmp17141 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp17142 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4myassume), W5109, W5115, Z5122, V5099, V5100, V5101, V5102)
return
}, 0)

__e.TailApply(PrimFunc(symbind), Z5120, tmp17141, V5099, V5100, V5101, tmp17142)
return


}, 0)

__e.TailApply(PrimFunc(symis_b), W5108, Z5119, V5099, V5100, V5101, tmp17140)
return


}, 0)

__e.TailApply(PrimFunc(symis_b), W5111, Z5121, V5099, V5100, V5101, tmp17139)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

__e.TailApply(tmp17032, tmp17137)
return


}, 1)

tmp17143 := Call(__e, PrimFunc(symshen_4lazyderef), V5098, V5099)


__e.TailApply(tmp17031, tmp17143)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp17146 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W5114)
}
__typedArg0 := W5114
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17147 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17146, V5099)


__e.TailApply(tmp17030, tmp17147)
return


}, 1)

tmp17148 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W5114)
}
__typedArg0 := W5114
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp17029, tmp17148)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp17151 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W5112)
}
__typedArg0 := W5112
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17152 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17151, V5099)


__e.TailApply(tmp17028, tmp17152)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp17155 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W5112)
}
__typedArg0 := W5112
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17156 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17155, V5099)


__e.TailApply(tmp17027, tmp17156)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp17159 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W5110)
}
__typedArg0 := W5110
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17160 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17159, V5099)


__e.TailApply(tmp17026, tmp17160)
return


}, 1)

tmp17161 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W5110)
}
__typedArg0 := W5110
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp17025, tmp17161)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp17164 := Call(__e, PrimFunc(symshen_4lazyderef), V5097, V5099)


__e.TailApply(tmp17024, tmp17164)
return


}, 1)

tmp17165 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W5107)
}
__typedArg0 := W5107
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(tmp17023, tmp17165)
return


}, 1)

tmp17166 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W5107)
}
__typedArg0 := W5107
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp17022, tmp17166)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp17169 := Call(__e, PrimFunc(symshen_4lazyderef), V5096, V5099)


__e.TailApply(tmp17021, tmp17169)
return


} else {
__e.Return(False)
return
}


} else {
__e.Return(W5103)
return
}


}, 1)

tmp17189 := Call(__e, PrimFunc(symshen_4unlocked_2), V5100)


var ifres17174 Obj

if True == tmp17189 {
tmp17175 := MakeNative(func(__e *ControlFlow) {
W5104 := __e.Get(1)
_ = W5104
tmp17186 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W5104, Nil)
}
__typedArg0 := W5104
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17186 {
tmp17176 := MakeNative(func(__e *ControlFlow) {
W5105 := __e.Get(1)
_ = W5105
tmp17177 := MakeNative(func(__e *ControlFlow) {
W5106 := __e.Get(1)
_ = W5106
tmp17181 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W5105, Nil)
}
__typedArg0 := W5105
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17181 {
__e.TailApply(PrimFunc(symthaw), W5106)
return
} else {
tmp17179 := Call(__e, PrimFunc(symshen_4pvar_2), W5105)


if True == tmp17179 {
__e.TailApply(PrimFunc(symshen_4bind_b), W5105, Nil, V5099, W5106)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp17182 := MakeNative(func(__e *ControlFlow) {
tmp17183 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp17183

__e.TailApply(PrimFunc(symthaw), V5102)
return


}, 0)

__e.TailApply(tmp17177, tmp17182)
return


}, 1)

tmp17184 := Call(__e, PrimFunc(symshen_4lazyderef), V5098, V5099)


__e.TailApply(tmp17176, tmp17184)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp17187 := Call(__e, PrimFunc(symshen_4lazyderef), V5096, V5099)


tmp17188 := Call(__e, tmp17175, tmp17187)


ifres17174 = tmp17188


} else {
ifres17174 = False


}

__e.TailApply(tmp17020, ifres17174)
return


}, 7)

tmp17190 := Call(__e, ns2_1set, symshen_4myassume, tmp17019)


_ = tmp17190

tmp17191 := MakeNative(func(__e *ControlFlow) {
V5153 := __e.Get(1)
_ = V5153
tmp17214 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V5153)
}
__typedArg0 := Nil
__typedArg1 := V5153
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17214 {
__e.Return(Nil)
return
} else {
tmp17212 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5153)
}
__typedArg0 := V5153
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17208 Obj

if True == tmp17212 {
tmp17210 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5153)
}
__typedArg0 := V5153
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17211 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp17210)
}
__typedArg0 := tmp17210
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17209 Obj

if True == tmp17211 {
ifres17209 = True


} else {
ifres17209 = False


}

ifres17208 = ifres17209


} else {
ifres17208 = False


}

if True == ifres17208 {
tmp17192 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5153)
}
__typedArg0 := V5153
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17193 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5153)
}
__typedArg0 := V5153
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17194 := Call(__e, PrimFunc(symappend), tmp17192, tmp17193)


__e.TailApply(PrimFunc(symshen_4freshterms), tmp17194)
return


} else {
tmp17206 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5153)
}
__typedArg0 := V5153
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17202 Obj

if True == tmp17206 {
tmp17204 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5153)
}
__typedArg0 := V5153
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17205 := Call(__e, PrimFunc(symshen_4freshterm_2), tmp17204)


var ifres17203 Obj

if True == tmp17205 {
ifres17203 = True


} else {
ifres17203 = False


}

ifres17202 = ifres17203


} else {
ifres17202 = False


}

if True == ifres17202 {
tmp17195 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5153)
}
__typedArg0 := V5153
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17196 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5153)
}
__typedArg0 := V5153
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17197 := Call(__e, PrimFunc(symshen_4freshterms), tmp17196)


__e.TailApply(PrimFunc(symadjoin), tmp17195, tmp17197)
return


} else {
tmp17200 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5153)
}
__typedArg0 := V5153
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17200 {
tmp17198 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5153)
}
__typedArg0 := V5153
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4freshterms), tmp17198)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("partial function shen.freshterms"))
}
__typedArg0 := MakeString("partial function shen.freshterms")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}


}


}, 1)

tmp17215 := Call(__e, ns2_1set, symshen_4freshterms, tmp17191)


_ = tmp17215

tmp17216 := MakeNative(func(__e *ControlFlow) {
V5154 := __e.Get(1)
_ = V5154
V5155 := __e.Get(2)
_ = V5155
V5156 := __e.Get(3)
_ = V5156
V5157 := __e.Get(4)
_ = V5157
V5158 := __e.Get(5)
_ = V5158
V5159 := __e.Get(6)
_ = V5159
tmp17217 := MakeNative(func(__e *ControlFlow) {
W5160 := __e.Get(1)
_ = W5160
tmp17341 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W5160, False)
}
__typedArg0 := W5160
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17341 {
tmp17339 := Call(__e, PrimFunc(symshen_4unlocked_2), V5157)


if True == tmp17339 {
tmp17218 := MakeNative(func(__e *ControlFlow) {
W5164 := __e.Get(1)
_ = W5164
tmp17336 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W5164)
}
__typedArg0 := W5164
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17336 {
tmp17219 := MakeNative(func(__e *ControlFlow) {
W5165 := __e.Get(1)
_ = W5165
tmp17220 := MakeNative(func(__e *ControlFlow) {
W5166 := __e.Get(1)
_ = W5166
tmp17221 := MakeNative(func(__e *ControlFlow) {
W5167 := __e.Get(1)
_ = W5167
tmp17222 := MakeNative(func(__e *ControlFlow) {
W5168 := __e.Get(1)
_ = W5168
tmp17326 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W5167)
}
__typedArg0 := W5167
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17326 {
tmp17223 := MakeNative(func(__e *ControlFlow) {
W5173 := __e.Get(1)
_ = W5173
tmp17224 := MakeNative(func(__e *ControlFlow) {
W5174 := __e.Get(1)
_ = W5174
tmp17294 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W5173)
}
__typedArg0 := W5173
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17294 {
tmp17225 := MakeNative(func(__e *ControlFlow) {
W5179 := __e.Get(1)
_ = W5179
tmp17226 := MakeNative(func(__e *ControlFlow) {
W5180 := __e.Get(1)
_ = W5180
tmp17227 := MakeNative(func(__e *ControlFlow) {
W5181 := __e.Get(1)
_ = W5181
tmp17269 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W5180)
}
__typedArg0 := W5180
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17269 {
tmp17228 := MakeNative(func(__e *ControlFlow) {
W5184 := __e.Get(1)
_ = W5184
tmp17229 := MakeNative(func(__e *ControlFlow) {
W5185 := __e.Get(1)
_ = W5185
tmp17230 := MakeNative(func(__e *ControlFlow) {
W5186 := __e.Get(1)
_ = W5186
tmp17250 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W5185)
}
__typedArg0 := W5185
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17250 {
tmp17231 := MakeNative(func(__e *ControlFlow) {
W5188 := __e.Get(1)
_ = W5188
tmp17232 := MakeNative(func(__e *ControlFlow) {
W5189 := __e.Get(1)
_ = W5189
tmp17233 := MakeNative(func(__e *ControlFlow) {
W5190 := __e.Get(1)
_ = W5190
tmp17237 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W5189, Nil)
}
__typedArg0 := W5189
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17237 {
__e.TailApply(PrimFunc(symthaw), W5190)
return
} else {
tmp17235 := Call(__e, PrimFunc(symshen_4pvar_2), W5189)


if True == tmp17235 {
__e.TailApply(PrimFunc(symshen_4bind_b), W5189, Nil, V5156, W5190)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp17238 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W5186, W5188)
return
}, 0)

__e.TailApply(tmp17233, tmp17238)
return


}, 1)

tmp17239 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W5185)
}
__typedArg0 := W5185
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17240 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17239, V5156)


__e.TailApply(tmp17232, tmp17240)
return


}, 1)

tmp17241 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W5185)
}
__typedArg0 := W5185
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp17231, tmp17241)
return


} else {
tmp17248 := Call(__e, PrimFunc(symshen_4pvar_2), W5185)


if True == tmp17248 {
tmp17242 := MakeNative(func(__e *ControlFlow) {
W5191 := __e.Get(1)
_ = W5191
tmp17243 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W5191, Nil)
}
__typedArg0 := W5191
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17244 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W5186, W5191)
return
}, 0)

tmp17245 := Call(__e, PrimFunc(symshen_4bind_b), W5185, tmp17243, V5156, tmp17244)


__e.TailApply(PrimFunc(symshen_4gc), V5156, tmp17245)
return


}, 1)

tmp17246 := Call(__e, PrimFunc(symshen_4newpv), V5156)


__e.TailApply(tmp17242, tmp17246)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp17251 := MakeNative(func(__e *ControlFlow) {
Z5187 := __e.Get(1)
_ = Z5187
tmp17252 := Call(__e, W5181, W5184)


__e.TailApply(tmp17252, Z5187)
return


}, 1)

__e.TailApply(tmp17230, tmp17251)
return


}, 1)

tmp17253 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W5180)
}
__typedArg0 := W5180
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17254 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17253, V5156)


__e.TailApply(tmp17229, tmp17254)
return


}, 1)

tmp17255 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W5180)
}
__typedArg0 := W5180
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp17228, tmp17255)
return


} else {
tmp17267 := Call(__e, PrimFunc(symshen_4pvar_2), W5180)


if True == tmp17267 {
tmp17256 := MakeNative(func(__e *ControlFlow) {
W5192 := __e.Get(1)
_ = W5192
tmp17257 := MakeNative(func(__e *ControlFlow) {
W5193 := __e.Get(1)
_ = W5193
tmp17258 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W5193, Nil)
}
__typedArg0 := W5193
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17259 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W5192, tmp17258)
}
__typedArg0 := W5192
__typedArg1 := tmp17258
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17260 := MakeNative(func(__e *ControlFlow) {
tmp17261 := Call(__e, W5181, W5192)


__e.TailApply(tmp17261, W5193)
return


}, 0)

tmp17262 := Call(__e, PrimFunc(symshen_4bind_b), W5180, tmp17259, V5156, tmp17260)


__e.TailApply(PrimFunc(symshen_4gc), V5156, tmp17262)
return


}, 1)

tmp17263 := Call(__e, PrimFunc(symshen_4newpv), V5156)


tmp17264 := Call(__e, tmp17257, tmp17263)


__e.TailApply(PrimFunc(symshen_4gc), V5156, tmp17264)
return


}, 1)

tmp17265 := Call(__e, PrimFunc(symshen_4newpv), V5156)


__e.TailApply(tmp17256, tmp17265)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp17270 := MakeNative(func(__e *ControlFlow) {
Z5182 := __e.Get(1)
_ = Z5182
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5183 := __e.Get(1)
_ = Z5183
tmp17271 := Call(__e, W5174, W5179)


tmp17272 := Call(__e, tmp17271, Z5182)


__e.TailApply(tmp17272, Z5183)
return


}, 1))
return
}, 1)

__e.TailApply(tmp17227, tmp17270)
return


}, 1)

tmp17273 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W5173)
}
__typedArg0 := W5173
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17274 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17273, V5156)


__e.TailApply(tmp17226, tmp17274)
return


}, 1)

tmp17275 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W5173)
}
__typedArg0 := W5173
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp17225, tmp17275)
return


} else {
tmp17292 := Call(__e, PrimFunc(symshen_4pvar_2), W5173)


if True == tmp17292 {
tmp17276 := MakeNative(func(__e *ControlFlow) {
W5194 := __e.Get(1)
_ = W5194
tmp17277 := MakeNative(func(__e *ControlFlow) {
W5195 := __e.Get(1)
_ = W5195
tmp17278 := MakeNative(func(__e *ControlFlow) {
W5196 := __e.Get(1)
_ = W5196
tmp17279 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W5196, Nil)
}
__typedArg0 := W5196
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17280 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W5195, tmp17279)
}
__typedArg0 := W5195
__typedArg1 := tmp17279
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17281 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W5194, tmp17280)
}
__typedArg0 := W5194
__typedArg1 := tmp17280
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17282 := MakeNative(func(__e *ControlFlow) {
tmp17283 := Call(__e, W5174, W5194)


tmp17284 := Call(__e, tmp17283, W5195)


__e.TailApply(tmp17284, W5196)
return


}, 0)

tmp17285 := Call(__e, PrimFunc(symshen_4bind_b), W5173, tmp17281, V5156, tmp17282)


__e.TailApply(PrimFunc(symshen_4gc), V5156, tmp17285)
return


}, 1)

tmp17286 := Call(__e, PrimFunc(symshen_4newpv), V5156)


tmp17287 := Call(__e, tmp17278, tmp17286)


__e.TailApply(PrimFunc(symshen_4gc), V5156, tmp17287)
return


}, 1)

tmp17288 := Call(__e, PrimFunc(symshen_4newpv), V5156)


tmp17289 := Call(__e, tmp17277, tmp17288)


__e.TailApply(PrimFunc(symshen_4gc), V5156, tmp17289)
return


}, 1)

tmp17290 := Call(__e, PrimFunc(symshen_4newpv), V5156)


__e.TailApply(tmp17276, tmp17290)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp17295 := MakeNative(func(__e *ControlFlow) {
Z5175 := __e.Get(1)
_ = Z5175
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5176 := __e.Get(1)
_ = Z5176
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5177 := __e.Get(1)
_ = Z5177
tmp17296 := MakeNative(func(__e *ControlFlow) {
W5178 := __e.Get(1)
_ = W5178
tmp17297 := Call(__e, W5168, Z5175)


tmp17298 := Call(__e, tmp17297, Z5176)


tmp17299 := Call(__e, tmp17298, Z5177)


__e.TailApply(tmp17299, W5178)
return


}, 1)

tmp17300 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W5167)
}
__typedArg0 := W5167
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(tmp17296, tmp17300)
return


}, 1))
return
}, 1))
return
}, 1)

__e.TailApply(tmp17224, tmp17295)
return


}, 1)

tmp17301 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W5167)
}
__typedArg0 := W5167
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17302 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17301, V5156)


__e.TailApply(tmp17223, tmp17302)
return


} else {
tmp17324 := Call(__e, PrimFunc(symshen_4pvar_2), W5167)


if True == tmp17324 {
tmp17303 := MakeNative(func(__e *ControlFlow) {
W5197 := __e.Get(1)
_ = W5197
tmp17304 := MakeNative(func(__e *ControlFlow) {
W5198 := __e.Get(1)
_ = W5198
tmp17305 := MakeNative(func(__e *ControlFlow) {
W5199 := __e.Get(1)
_ = W5199
tmp17306 := MakeNative(func(__e *ControlFlow) {
W5200 := __e.Get(1)
_ = W5200
tmp17307 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W5199, Nil)
}
__typedArg0 := W5199
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17308 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W5198, tmp17307)
}
__typedArg0 := W5198
__typedArg1 := tmp17307
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17309 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W5197, tmp17308)
}
__typedArg0 := W5197
__typedArg1 := tmp17308
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17310 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17309, W5200)
}
__typedArg0 := tmp17309
__typedArg1 := W5200
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17311 := MakeNative(func(__e *ControlFlow) {
tmp17312 := Call(__e, W5168, W5197)


tmp17313 := Call(__e, tmp17312, W5198)


tmp17314 := Call(__e, tmp17313, W5199)


__e.TailApply(tmp17314, W5200)
return


}, 0)

tmp17315 := Call(__e, PrimFunc(symshen_4bind_b), W5167, tmp17310, V5156, tmp17311)


__e.TailApply(PrimFunc(symshen_4gc), V5156, tmp17315)
return


}, 1)

tmp17316 := Call(__e, PrimFunc(symshen_4newpv), V5156)


tmp17317 := Call(__e, tmp17306, tmp17316)


__e.TailApply(PrimFunc(symshen_4gc), V5156, tmp17317)
return


}, 1)

tmp17318 := Call(__e, PrimFunc(symshen_4newpv), V5156)


tmp17319 := Call(__e, tmp17305, tmp17318)


__e.TailApply(PrimFunc(symshen_4gc), V5156, tmp17319)
return


}, 1)

tmp17320 := Call(__e, PrimFunc(symshen_4newpv), V5156)


tmp17321 := Call(__e, tmp17304, tmp17320)


__e.TailApply(PrimFunc(symshen_4gc), V5156, tmp17321)
return


}, 1)

tmp17322 := Call(__e, PrimFunc(symshen_4newpv), V5156)


__e.TailApply(tmp17303, tmp17322)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp17327 := MakeNative(func(__e *ControlFlow) {
Z5169 := __e.Get(1)
_ = Z5169
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5170 := __e.Get(1)
_ = Z5170
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5171 := __e.Get(1)
_ = Z5171
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5172 := __e.Get(1)
_ = Z5172
tmp17328 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp17328

tmp17329 := MakeNative(func(__e *ControlFlow) {
tmp17330 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp17331 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4p_1hyps), W5166, Z5172, V5156, V5157, V5158, V5159)
return
}, 0)

__e.TailApply(PrimFunc(symbind), Z5170, tmp17330, V5156, V5157, V5158, tmp17331)
return


}, 0)

__e.TailApply(PrimFunc(symbind), Z5169, W5165, V5156, V5157, V5158, tmp17329)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

__e.TailApply(tmp17222, tmp17327)
return


}, 1)

tmp17332 := Call(__e, PrimFunc(symshen_4lazyderef), V5155, V5156)


__e.TailApply(tmp17221, tmp17332)
return


}, 1)

tmp17333 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W5164)
}
__typedArg0 := W5164
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(tmp17220, tmp17333)
return


}, 1)

tmp17334 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W5164)
}
__typedArg0 := W5164
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp17219, tmp17334)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp17337 := Call(__e, PrimFunc(symshen_4lazyderef), V5154, V5156)


__e.TailApply(tmp17218, tmp17337)
return


} else {
__e.Return(False)
return
}


} else {
__e.Return(W5160)
return
}


}, 1)

tmp17357 := Call(__e, PrimFunc(symshen_4unlocked_2), V5157)


var ifres17342 Obj

if True == tmp17357 {
tmp17343 := MakeNative(func(__e *ControlFlow) {
W5161 := __e.Get(1)
_ = W5161
tmp17354 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W5161, Nil)
}
__typedArg0 := W5161
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17354 {
tmp17344 := MakeNative(func(__e *ControlFlow) {
W5162 := __e.Get(1)
_ = W5162
tmp17345 := MakeNative(func(__e *ControlFlow) {
W5163 := __e.Get(1)
_ = W5163
tmp17349 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W5162, Nil)
}
__typedArg0 := W5162
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17349 {
__e.TailApply(PrimFunc(symthaw), W5163)
return
} else {
tmp17347 := Call(__e, PrimFunc(symshen_4pvar_2), W5162)


if True == tmp17347 {
__e.TailApply(PrimFunc(symshen_4bind_b), W5162, Nil, V5156, W5163)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp17350 := MakeNative(func(__e *ControlFlow) {
tmp17351 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp17351

__e.TailApply(PrimFunc(symthaw), V5159)
return


}, 0)

__e.TailApply(tmp17345, tmp17350)
return


}, 1)

tmp17352 := Call(__e, PrimFunc(symshen_4lazyderef), V5155, V5156)


__e.TailApply(tmp17344, tmp17352)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp17355 := Call(__e, PrimFunc(symshen_4lazyderef), V5154, V5156)


tmp17356 := Call(__e, tmp17343, tmp17355)


ifres17342 = tmp17356


} else {
ifres17342 = False


}

__e.TailApply(tmp17217, ifres17342)
return


}, 6)

tmp17358 := Call(__e, ns2_1set, symshen_4p_1hyps, tmp17216)


_ = tmp17358

tmp17359 := MakeNative(func(__e *ControlFlow) {
V5201 := __e.Get(1)
_ = V5201
V5202 := __e.Get(2)
_ = V5202
V5203 := __e.Get(3)
_ = V5203
V5204 := __e.Get(4)
_ = V5204
V5205 := __e.Get(5)
_ = V5205
V5206 := __e.Get(6)
_ = V5206
V5207 := __e.Get(7)
_ = V5207
tmp17360 := MakeNative(func(__e *ControlFlow) {
W5208 := __e.Get(1)
_ = W5208
tmp17361 := MakeNative(func(__e *ControlFlow) {
W5209 := __e.Get(1)
_ = W5209
tmp17371 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W5209, False)
}
__typedArg0 := W5209
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17371 {
tmp17362 := MakeNative(func(__e *ControlFlow) {
W5218 := __e.Get(1)
_ = W5218
tmp17364 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W5218, False)
}
__typedArg0 := W5218
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17364 {
__e.TailApply(PrimFunc(symshen_4unlock), V5205, W5208)
return
} else {
__e.Return(W5218)
return
}


}, 1)

tmp17369 := Call(__e, PrimFunc(symshen_4unlocked_2), V5205)


var ifres17365 Obj

if True == tmp17369 {
tmp17366 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp17366

tmp17367 := Call(__e, PrimFunc(symshen_4curry), V5201)


tmp17368 := Call(__e, PrimFunc(symshen_4system_1S_1h), tmp17367, V5202, V5203, V5204, V5205, W5208, V5207)


ifres17365 = tmp17368


} else {
ifres17365 = False


}

__e.TailApply(tmp17362, ifres17365)
return


} else {
__e.Return(W5209)
return
}


}, 1)

tmp17416 := Call(__e, PrimFunc(symshen_4unlocked_2), V5205)


var ifres17372 Obj

if True == tmp17416 {
tmp17373 := MakeNative(func(__e *ControlFlow) {
W5210 := __e.Get(1)
_ = W5210
tmp17413 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W5210)
}
__typedArg0 := W5210
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17413 {
tmp17374 := MakeNative(func(__e *ControlFlow) {
W5211 := __e.Get(1)
_ = W5211
tmp17409 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W5211, symwhere)
}
__typedArg0 := W5211
__typedArg1 := symwhere
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17409 {
tmp17375 := MakeNative(func(__e *ControlFlow) {
W5212 := __e.Get(1)
_ = W5212
tmp17405 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W5212)
}
__typedArg0 := W5212
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17405 {
tmp17376 := MakeNative(func(__e *ControlFlow) {
W5213 := __e.Get(1)
_ = W5213
tmp17377 := MakeNative(func(__e *ControlFlow) {
W5214 := __e.Get(1)
_ = W5214
tmp17400 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W5214)
}
__typedArg0 := W5214
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17400 {
tmp17378 := MakeNative(func(__e *ControlFlow) {
W5215 := __e.Get(1)
_ = W5215
tmp17379 := MakeNative(func(__e *ControlFlow) {
W5216 := __e.Get(1)
_ = W5216
tmp17395 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W5216, Nil)
}
__typedArg0 := W5216
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17395 {
tmp17380 := MakeNative(func(__e *ControlFlow) {
W5217 := __e.Get(1)
_ = W5217
tmp17381 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp17381

tmp17382 := MakeNative(func(__e *ControlFlow) {
tmp17383 := Call(__e, PrimFunc(symshen_4curry), W5213)


tmp17384 := MakeNative(func(__e *ControlFlow) {
tmp17385 := MakeNative(func(__e *ControlFlow) {
tmp17386 := MakeNative(func(__e *ControlFlow) {
tmp17387 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp17388 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symverified, Nil)
}
__typedArg0 := symverified
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17389 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17387, tmp17388)
}
__typedArg0 := tmp17387
__typedArg1 := tmp17388
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17390 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W5217, tmp17389)
}
__typedArg0 := W5217
__typedArg1 := tmp17389
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17391 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17390, V5203)
}
__typedArg0 := tmp17390
__typedArg1 := V5203
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4t_d_1correct), W5215, V5202, tmp17391, V5204, V5205, W5208, V5207)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V5204, V5205, W5208, tmp17386)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), W5217, symboolean, V5203, V5204, V5205, W5208, tmp17385)
return


}, 0)

__e.TailApply(PrimFunc(symbind), W5217, tmp17383, V5204, V5205, W5208, tmp17384)
return


}, 0)

tmp17392 := Call(__e, PrimFunc(symshen_4cut), V5204, V5205, W5208, tmp17382)


__e.TailApply(PrimFunc(symshen_4gc), V5204, tmp17392)
return


}, 1)

tmp17393 := Call(__e, PrimFunc(symshen_4newpv), V5204)


__e.TailApply(tmp17380, tmp17393)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp17396 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W5214)
}
__typedArg0 := W5214
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17397 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17396, V5204)


__e.TailApply(tmp17379, tmp17397)
return


}, 1)

tmp17398 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W5214)
}
__typedArg0 := W5214
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp17378, tmp17398)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp17401 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W5212)
}
__typedArg0 := W5212
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17402 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17401, V5204)


__e.TailApply(tmp17377, tmp17402)
return


}, 1)

tmp17403 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W5212)
}
__typedArg0 := W5212
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp17376, tmp17403)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp17406 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W5210)
}
__typedArg0 := W5210
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17407 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17406, V5204)


__e.TailApply(tmp17375, tmp17407)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp17410 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W5210)
}
__typedArg0 := W5210
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17411 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17410, V5204)


__e.TailApply(tmp17374, tmp17411)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp17414 := Call(__e, PrimFunc(symshen_4lazyderef), V5201, V5204)


tmp17415 := Call(__e, tmp17373, tmp17414)


ifres17372 = tmp17415


} else {
ifres17372 = False


}

__e.TailApply(tmp17361, ifres17372)
return


}, 1)

__e.TailApply(tmp17360, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V5206)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V5206
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())
return


}, 7)

tmp17418 := Call(__e, ns2_1set, symshen_4t_d_1correct, tmp17359)


_ = tmp17418

tmp17419 := MakeNative(func(__e *ControlFlow) {
V5219 := __e.Get(1)
_ = V5219
V5220 := __e.Get(2)
_ = V5220
V5221 := __e.Get(3)
_ = V5221
V5222 := __e.Get(4)
_ = V5222
V5223 := __e.Get(5)
_ = V5223
V5224 := __e.Get(6)
_ = V5224
V5225 := __e.Get(7)
_ = V5225
V5226 := __e.Get(8)
_ = V5226
tmp17420 := MakeNative(func(__e *ControlFlow) {
W5227 := __e.Get(1)
_ = W5227
tmp17462 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W5227, False)
}
__typedArg0 := W5227
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17462 {
tmp17460 := Call(__e, PrimFunc(symshen_4unlocked_2), V5224)


if True == tmp17460 {
tmp17421 := MakeNative(func(__e *ControlFlow) {
W5229 := __e.Get(1)
_ = W5229
tmp17457 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W5229)
}
__typedArg0 := W5229
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17457 {
tmp17422 := MakeNative(func(__e *ControlFlow) {
W5230 := __e.Get(1)
_ = W5230
tmp17423 := MakeNative(func(__e *ControlFlow) {
W5231 := __e.Get(1)
_ = W5231
tmp17424 := MakeNative(func(__e *ControlFlow) {
W5232 := __e.Get(1)
_ = W5232
tmp17452 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W5232)
}
__typedArg0 := W5232
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17452 {
tmp17425 := MakeNative(func(__e *ControlFlow) {
W5233 := __e.Get(1)
_ = W5233
tmp17426 := MakeNative(func(__e *ControlFlow) {
W5234 := __e.Get(1)
_ = W5234
tmp17447 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W5234)
}
__typedArg0 := W5234
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17447 {
tmp17427 := MakeNative(func(__e *ControlFlow) {
W5235 := __e.Get(1)
_ = W5235
tmp17443 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W5235, sym_1_1_6)
}
__typedArg0 := W5235
__typedArg1 := sym_1_1_6
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17443 {
tmp17428 := MakeNative(func(__e *ControlFlow) {
W5236 := __e.Get(1)
_ = W5236
tmp17439 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W5236)
}
__typedArg0 := W5236
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17439 {
tmp17429 := MakeNative(func(__e *ControlFlow) {
W5237 := __e.Get(1)
_ = W5237
tmp17430 := MakeNative(func(__e *ControlFlow) {
W5238 := __e.Get(1)
_ = W5238
tmp17434 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W5238, Nil)
}
__typedArg0 := W5238
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17434 {
tmp17431 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp17431

tmp17432 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4t_d_1integrity), W5231, W5237, V5221, V5222, V5223, V5224, V5225, V5226)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), W5230, W5233, V5221, V5223, V5224, V5225, tmp17432)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp17435 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W5236)
}
__typedArg0 := W5236
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17436 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17435, V5223)


__e.TailApply(tmp17430, tmp17436)
return


}, 1)

tmp17437 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W5236)
}
__typedArg0 := W5236
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp17429, tmp17437)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp17440 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W5234)
}
__typedArg0 := W5234
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17441 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17440, V5223)


__e.TailApply(tmp17428, tmp17441)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp17444 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W5234)
}
__typedArg0 := W5234
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17445 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17444, V5223)


__e.TailApply(tmp17427, tmp17445)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp17448 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W5232)
}
__typedArg0 := W5232
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17449 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17448, V5223)


__e.TailApply(tmp17426, tmp17449)
return


}, 1)

tmp17450 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W5232)
}
__typedArg0 := W5232
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp17425, tmp17450)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp17453 := Call(__e, PrimFunc(symshen_4lazyderef), V5220, V5223)


__e.TailApply(tmp17424, tmp17453)
return


}, 1)

tmp17454 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W5229)
}
__typedArg0 := W5229
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(tmp17423, tmp17454)
return


}, 1)

tmp17455 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W5229)
}
__typedArg0 := W5229
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp17422, tmp17455)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp17458 := Call(__e, PrimFunc(symshen_4lazyderef), V5219, V5223)


__e.TailApply(tmp17421, tmp17458)
return


} else {
__e.Return(False)
return
}


} else {
__e.Return(W5227)
return
}


}, 1)

tmp17470 := Call(__e, PrimFunc(symshen_4unlocked_2), V5224)


var ifres17463 Obj

if True == tmp17470 {
tmp17464 := MakeNative(func(__e *ControlFlow) {
W5228 := __e.Get(1)
_ = W5228
tmp17467 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W5228, Nil)
}
__typedArg0 := W5228
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17467 {
tmp17465 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp17465

__e.TailApply(PrimFunc(symis_b), V5220, V5222, V5223, V5224, V5225, V5226)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp17468 := Call(__e, PrimFunc(symshen_4lazyderef), V5219, V5223)


tmp17469 := Call(__e, tmp17464, tmp17468)


ifres17463 = tmp17469


} else {
ifres17463 = False


}

__e.TailApply(tmp17420, ifres17463)
return


}, 8)

tmp17471 := Call(__e, ns2_1set, symshen_4t_d_1integrity, tmp17419)


_ = tmp17471

tmp17472 := MakeNative(func(__e *ControlFlow) {
V5239 := __e.Get(1)
_ = V5239
tmp17481 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symabsvector_2) {
return PrimIsVector(V5239)
}
__typedArg0 := V5239
return Call(__e, PrimFunc(symabsvector_2), __typedArg0)
})()

if True == tmp17481 {
tmp17478 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_2) {
return PrimIsString(V5239)
}
__typedArg0 := V5239
return Call(__e, PrimFunc(symstring_2), __typedArg0)
})()

tmp17479 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp17478)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp17478
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres17474 Obj

if True == tmp17479 {
tmp17476 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V5239, MakeNumber(0))
}
__typedArg0 := V5239
__typedArg1 := MakeNumber(0)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})()

tmp17477 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp17476, symshen_4print_1freshterm)
}
__typedArg0 := tmp17476
__typedArg1 := symshen_4print_1freshterm
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17475 Obj

if True == tmp17477 {
ifres17475 = True


} else {
ifres17475 = False


}

ifres17474 = ifres17475


} else {
ifres17474 = False


}

if True == ifres17474 {
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

__e.TailApply(ns2_1set, symshen_4freshterm_2, tmp17472)
return




}, 0)

