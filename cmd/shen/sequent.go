package main

import . "github.com/pyrex41/shen-go/kl"

var SequentMain = MakeNative(func(__e *ControlFlow) {
tmp12313 := MakeNative(func(__e *ControlFlow) {
V3422 := __e.Get(1)
_ = V3422
tmp12328 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3422)
}
__typedArg0 := V3422
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12314 Obj

if True == tmp12328 {
tmp12315 := Call(__e, PrimFunc(symhead), V3422)


let__12083 := tmp12315
_ = let__12083

tmp12316 := Call(__e, PrimFunc(symtail), V3422)


let__12084 := tmp12316
_ = let__12084

tmp12317 := Call(__e, PrimFunc(symshen_4_5datatype_1rules_6), let__12084)


let__12085 := tmp12317
_ = let__12085

tmp12326 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12085)


var ifres12318 Obj

if True == tmp12326 {
tmp12319 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12318 = tmp12319


} else {
tmp12320 := Call(__e, PrimFunc(symshen_4_5_1out), let__12085)


let__12086 := tmp12320
_ = let__12086

tmp12321 := Call(__e, PrimFunc(symshen_4in_1_6), let__12085)


let__12087 := tmp12321
_ = let__12087

tmp12322 := Call(__e, PrimFunc(symshen_4rules_1_6prolog), let__12083, let__12086)


let__12088 := tmp12322
_ = let__12088

tmp12323 := Call(__e, PrimFunc(symfn), let__12083)


tmp12324 := Call(__e, PrimFunc(symshen_4remember_1datatype), let__12083, tmp12323)


tmp12325 := Call(__e, PrimFunc(symshen_4comb), let__12087, tmp12324)


ifres12318 = tmp12325


}

ifres12314 = ifres12318


} else {
tmp12327 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12314 = tmp12327


}

let__12082 := ifres12314
_ = let__12082

tmp12330 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12082)


if True == tmp12330 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__12082)
return
}


}, 1)

tmp12331 := Call(__e, ns2_1set, symshen_4_5datatype_6, tmp12313)


_ = tmp12331

tmp12332 := MakeNative(func(__e *ControlFlow) {
V3430 := __e.Get(1)
_ = V3430
V3431 := __e.Get(2)
_ = V3431
tmp12333 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_ddatatypes_d)
}
__typedArg0 := symshen_4_ddatatypes_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp12334 := Call(__e, PrimFunc(symshen_4assoc_1_6), V3430, V3431, tmp12333)


tmp12335 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_ddatatypes_d, tmp12334)
}
__typedArg0 := symshen_4_ddatatypes_d
__typedArg1 := tmp12334
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp12335

tmp12336 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dalldatatypes_d)
}
__typedArg0 := symshen_4_dalldatatypes_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp12337 := Call(__e, PrimFunc(symshen_4assoc_1_6), V3430, V3431, tmp12336)


tmp12338 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dalldatatypes_d, tmp12337)
}
__typedArg0 := symshen_4_dalldatatypes_d
__typedArg1 := tmp12337
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp12338

__e.Return(V3430)
return


}, 2)

tmp12339 := Call(__e, ns2_1set, symshen_4remember_1datatype, tmp12332)


_ = tmp12339

tmp12340 := MakeNative(func(__e *ControlFlow) {
V3432 := __e.Get(1)
_ = V3432
tmp12341 := Call(__e, PrimFunc(symshen_4_5datatype_1rule_6), V3432)


let__12090 := tmp12341
_ = let__12090

tmp12354 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12090)


var ifres12342 Obj

if True == tmp12354 {
tmp12343 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12342 = tmp12343


} else {
tmp12344 := Call(__e, PrimFunc(symshen_4_5_1out), let__12090)


let__12091 := tmp12344
_ = let__12091

tmp12345 := Call(__e, PrimFunc(symshen_4in_1_6), let__12090)


let__12092 := tmp12345
_ = let__12092

tmp12346 := Call(__e, PrimFunc(symshen_4_5datatype_1rules_6), let__12092)


let__12093 := tmp12346
_ = let__12093

tmp12353 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12093)


var ifres12347 Obj

if True == tmp12353 {
tmp12348 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12347 = tmp12348


} else {
tmp12349 := Call(__e, PrimFunc(symshen_4_5_1out), let__12093)


let__12094 := tmp12349
_ = let__12094

tmp12350 := Call(__e, PrimFunc(symshen_4in_1_6), let__12093)


let__12095 := tmp12350
_ = let__12095

tmp12351 := Call(__e, PrimFunc(symappend), let__12091, let__12094)


tmp12352 := Call(__e, PrimFunc(symshen_4comb), let__12095, tmp12351)


ifres12347 = tmp12352


}

ifres12342 = ifres12347


}

let__12089 := ifres12342
_ = let__12089

tmp12370 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12089)


if True == tmp12370 {
tmp12355 := Call(__e, PrimFunc(sym_5_b_6), V3432)


let__12097 := tmp12355
_ = let__12097

tmp12366 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12097)


var ifres12356 Obj

if True == tmp12366 {
tmp12357 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12356 = tmp12357


} else {
tmp12358 := Call(__e, PrimFunc(symshen_4_5_1out), let__12097)


let__12098 := tmp12358
_ = let__12098

tmp12359 := Call(__e, PrimFunc(symshen_4in_1_6), let__12097)


let__12099 := tmp12359
_ = let__12099

tmp12364 := Call(__e, PrimFunc(symempty_2), let__12098)


var ifres12360 Obj

if True == tmp12364 {
ifres12360 = Nil


} else {
tmp12361 := Call(__e, PrimFunc(symshen_4app), let__12098, MakeString("\n ..."), symshen_4r)


tmp12363 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("datatype syntax error here:\n "))
__typedS1, __typedOK1 := TypedString(tmp12361)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("datatype syntax error here:\n ")
__typedArg1 := tmp12361
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("datatype syntax error here:\n "))
__typedS1, __typedOK1 := TypedString(tmp12361)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("datatype syntax error here:\n ")
__typedArg1 := tmp12361
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})()

ifres12360 = tmp12363


}

tmp12365 := Call(__e, PrimFunc(symshen_4comb), let__12099, ifres12360)


ifres12356 = tmp12365


}

let__12096 := ifres12356
_ = let__12096

tmp12368 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12096)


if True == tmp12368 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__12096)
return
}


} else {
__e.Return(let__12089)
return
}


}, 1)

tmp12371 := Call(__e, ns2_1set, symshen_4_5datatype_1rules_6, tmp12340)


_ = tmp12371

tmp12372 := MakeNative(func(__e *ControlFlow) {
V3444 := __e.Get(1)
_ = V3444
tmp12373 := Call(__e, PrimFunc(symshen_4_5single_6), V3444)


let__12101 := tmp12373
_ = let__12101

tmp12379 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12101)


var ifres12374 Obj

if True == tmp12379 {
tmp12375 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12374 = tmp12375


} else {
tmp12376 := Call(__e, PrimFunc(symshen_4_5_1out), let__12101)


let__12102 := tmp12376
_ = let__12102

tmp12377 := Call(__e, PrimFunc(symshen_4in_1_6), let__12101)


let__12103 := tmp12377
_ = let__12103

tmp12378 := Call(__e, PrimFunc(symshen_4comb), let__12103, let__12102)


ifres12374 = tmp12378


}

let__12100 := ifres12374
_ = let__12100

tmp12390 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12100)


if True == tmp12390 {
tmp12380 := Call(__e, PrimFunc(symshen_4_5double_6), V3444)


let__12105 := tmp12380
_ = let__12105

tmp12386 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12105)


var ifres12381 Obj

if True == tmp12386 {
tmp12382 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12381 = tmp12382


} else {
tmp12383 := Call(__e, PrimFunc(symshen_4_5_1out), let__12105)


let__12106 := tmp12383
_ = let__12106

tmp12384 := Call(__e, PrimFunc(symshen_4in_1_6), let__12105)


let__12107 := tmp12384
_ = let__12107

tmp12385 := Call(__e, PrimFunc(symshen_4comb), let__12107, let__12106)


ifres12381 = tmp12385


}

let__12104 := ifres12381
_ = let__12104

tmp12388 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12104)


if True == tmp12388 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__12104)
return
}


} else {
__e.Return(let__12100)
return
}


}, 1)

tmp12391 := Call(__e, ns2_1set, symshen_4_5datatype_1rule_6, tmp12372)


_ = tmp12391

tmp12392 := MakeNative(func(__e *ControlFlow) {
V3453 := __e.Get(1)
_ = V3453
tmp12393 := Call(__e, PrimFunc(symshen_4_5sides_6), V3453)


let__12109 := tmp12393
_ = let__12109

tmp12425 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12109)


var ifres12394 Obj

if True == tmp12425 {
tmp12395 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12394 = tmp12395


} else {
tmp12396 := Call(__e, PrimFunc(symshen_4_5_1out), let__12109)


let__12110 := tmp12396
_ = let__12110

tmp12397 := Call(__e, PrimFunc(symshen_4in_1_6), let__12109)


let__12111 := tmp12397
_ = let__12111

tmp12398 := Call(__e, PrimFunc(symshen_4_5prems_6), let__12111)


let__12112 := tmp12398
_ = let__12112

tmp12424 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12112)


var ifres12399 Obj

if True == tmp12424 {
tmp12400 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12399 = tmp12400


} else {
tmp12401 := Call(__e, PrimFunc(symshen_4_5_1out), let__12112)


let__12113 := tmp12401
_ = let__12113

tmp12402 := Call(__e, PrimFunc(symshen_4in_1_6), let__12112)


let__12114 := tmp12402
_ = let__12114

tmp12403 := Call(__e, PrimFunc(symshen_4_5sng_6), let__12114)


let__12115 := tmp12403
_ = let__12115

tmp12423 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12115)


var ifres12404 Obj

if True == tmp12423 {
tmp12405 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12404 = tmp12405


} else {
tmp12406 := Call(__e, PrimFunc(symshen_4in_1_6), let__12115)


let__12116 := tmp12406
_ = let__12116

tmp12407 := Call(__e, PrimFunc(symshen_4_5conc_6), let__12116)


let__12117 := tmp12407
_ = let__12117

tmp12422 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12117)


var ifres12408 Obj

if True == tmp12422 {
tmp12409 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12408 = tmp12409


} else {
tmp12410 := Call(__e, PrimFunc(symshen_4_5_1out), let__12117)


let__12118 := tmp12410
_ = let__12118

tmp12411 := Call(__e, PrimFunc(symshen_4in_1_6), let__12117)


let__12119 := tmp12411
_ = let__12119

tmp12412 := Call(__e, PrimFunc(symshen_4_5sc_6), let__12119)


let__12120 := tmp12412
_ = let__12120

tmp12421 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12120)


var ifres12413 Obj

if True == tmp12421 {
tmp12414 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12413 = tmp12414


} else {
tmp12415 := Call(__e, PrimFunc(symshen_4in_1_6), let__12120)


let__12121 := tmp12415
_ = let__12121

tmp12416 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12118, Nil)
}
__typedArg0 := let__12118
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12417 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12113, tmp12416)
}
__typedArg0 := let__12113
__typedArg1 := tmp12416
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12418 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12110, tmp12417)
}
__typedArg0 := let__12110
__typedArg1 := tmp12417
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12419 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12418, Nil)
}
__typedArg0 := tmp12418
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12420 := Call(__e, PrimFunc(symshen_4comb), let__12121, tmp12419)


ifres12413 = tmp12420


}

ifres12408 = ifres12413


}

ifres12404 = ifres12408


}

ifres12399 = ifres12404


}

ifres12394 = ifres12399


}

let__12108 := ifres12394
_ = let__12108

tmp12427 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12108)


if True == tmp12427 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__12108)
return
}


}, 1)

tmp12428 := Call(__e, ns2_1set, symshen_4_5single_6, tmp12392)


_ = tmp12428

tmp12429 := MakeNative(func(__e *ControlFlow) {
V3468 := __e.Get(1)
_ = V3468
tmp12430 := Call(__e, PrimFunc(symshen_4_5sides_6), V3468)


let__12123 := tmp12430
_ = let__12123

tmp12461 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12123)


var ifres12431 Obj

if True == tmp12461 {
tmp12432 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12431 = tmp12432


} else {
tmp12433 := Call(__e, PrimFunc(symshen_4_5_1out), let__12123)


let__12124 := tmp12433
_ = let__12124

tmp12434 := Call(__e, PrimFunc(symshen_4in_1_6), let__12123)


let__12125 := tmp12434
_ = let__12125

tmp12435 := Call(__e, PrimFunc(symshen_4_5formulae_6), let__12125)


let__12126 := tmp12435
_ = let__12126

tmp12460 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12126)


var ifres12436 Obj

if True == tmp12460 {
tmp12437 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12436 = tmp12437


} else {
tmp12438 := Call(__e, PrimFunc(symshen_4_5_1out), let__12126)


let__12127 := tmp12438
_ = let__12127

tmp12439 := Call(__e, PrimFunc(symshen_4in_1_6), let__12126)


let__12128 := tmp12439
_ = let__12128

tmp12440 := Call(__e, PrimFunc(symshen_4_5dbl_6), let__12128)


let__12129 := tmp12440
_ = let__12129

tmp12459 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12129)


var ifres12441 Obj

if True == tmp12459 {
tmp12442 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12441 = tmp12442


} else {
tmp12443 := Call(__e, PrimFunc(symshen_4in_1_6), let__12129)


let__12130 := tmp12443
_ = let__12130

tmp12444 := Call(__e, PrimFunc(symshen_4_5formula_6), let__12130)


let__12131 := tmp12444
_ = let__12131

tmp12458 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12131)


var ifres12445 Obj

if True == tmp12458 {
tmp12446 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12445 = tmp12446


} else {
tmp12447 := Call(__e, PrimFunc(symshen_4_5_1out), let__12131)


let__12132 := tmp12447
_ = let__12132

tmp12448 := Call(__e, PrimFunc(symshen_4in_1_6), let__12131)


let__12133 := tmp12448
_ = let__12133

tmp12449 := Call(__e, PrimFunc(symshen_4_5sc_6), let__12133)


let__12134 := tmp12449
_ = let__12134

tmp12457 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12134)


var ifres12450 Obj

if True == tmp12457 {
tmp12451 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12450 = tmp12451


} else {
tmp12452 := Call(__e, PrimFunc(symshen_4in_1_6), let__12134)


let__12135 := tmp12452
_ = let__12135

tmp12453 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12132, Nil)
}
__typedArg0 := let__12132
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12454 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Nil, tmp12453)
}
__typedArg0 := Nil
__typedArg1 := tmp12453
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12455 := Call(__e, PrimFunc(symshen_4lr_1rule), let__12124, let__12127, tmp12454)


tmp12456 := Call(__e, PrimFunc(symshen_4comb), let__12135, tmp12455)


ifres12450 = tmp12456


}

ifres12445 = ifres12450


}

ifres12441 = ifres12445


}

ifres12436 = ifres12441


}

ifres12431 = ifres12436


}

let__12122 := ifres12431
_ = let__12122

tmp12463 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12122)


if True == tmp12463 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__12122)
return
}


}, 1)

tmp12464 := Call(__e, ns2_1set, symshen_4_5double_6, tmp12429)


_ = tmp12464

tmp12465 := MakeNative(func(__e *ControlFlow) {
V3483 := __e.Get(1)
_ = V3483
tmp12466 := Call(__e, PrimFunc(symshen_4_5formula_6), V3483)


let__12137 := tmp12466
_ = let__12137

tmp12486 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12137)


var ifres12467 Obj

if True == tmp12486 {
tmp12468 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12467 = tmp12468


} else {
tmp12469 := Call(__e, PrimFunc(symshen_4_5_1out), let__12137)


let__12138 := tmp12469
_ = let__12138

tmp12470 := Call(__e, PrimFunc(symshen_4in_1_6), let__12137)


let__12139 := tmp12470
_ = let__12139

tmp12471 := Call(__e, PrimFunc(symshen_4_5sc_6), let__12139)


let__12140 := tmp12471
_ = let__12140

tmp12485 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12140)


var ifres12472 Obj

if True == tmp12485 {
tmp12473 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12472 = tmp12473


} else {
tmp12474 := Call(__e, PrimFunc(symshen_4in_1_6), let__12140)


let__12141 := tmp12474
_ = let__12141

tmp12475 := Call(__e, PrimFunc(symshen_4_5formulae_6), let__12141)


let__12142 := tmp12475
_ = let__12142

tmp12484 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12142)


var ifres12476 Obj

if True == tmp12484 {
tmp12477 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12476 = tmp12477


} else {
tmp12478 := Call(__e, PrimFunc(symshen_4_5_1out), let__12142)


let__12143 := tmp12478
_ = let__12143

tmp12479 := Call(__e, PrimFunc(symshen_4in_1_6), let__12142)


let__12144 := tmp12479
_ = let__12144

tmp12480 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12138, Nil)
}
__typedArg0 := let__12138
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12481 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Nil, tmp12480)
}
__typedArg0 := Nil
__typedArg1 := tmp12480
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12482 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12481, let__12143)
}
__typedArg0 := tmp12481
__typedArg1 := let__12143
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12483 := Call(__e, PrimFunc(symshen_4comb), let__12144, tmp12482)


ifres12476 = tmp12483


}

ifres12472 = ifres12476


}

ifres12467 = ifres12472


}

let__12136 := ifres12467
_ = let__12136

tmp12505 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12136)


if True == tmp12505 {
tmp12487 := Call(__e, PrimFunc(symshen_4_5formula_6), V3483)


let__12146 := tmp12487
_ = let__12146

tmp12501 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12146)


var ifres12488 Obj

if True == tmp12501 {
tmp12489 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12488 = tmp12489


} else {
tmp12490 := Call(__e, PrimFunc(symshen_4_5_1out), let__12146)


let__12147 := tmp12490
_ = let__12147

tmp12491 := Call(__e, PrimFunc(symshen_4in_1_6), let__12146)


let__12148 := tmp12491
_ = let__12148

tmp12492 := Call(__e, PrimFunc(symshen_4_5sc_6), let__12148)


let__12149 := tmp12492
_ = let__12149

tmp12500 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12149)


var ifres12493 Obj

if True == tmp12500 {
tmp12494 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12493 = tmp12494


} else {
tmp12495 := Call(__e, PrimFunc(symshen_4in_1_6), let__12149)


let__12150 := tmp12495
_ = let__12150

tmp12496 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12147, Nil)
}
__typedArg0 := let__12147
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12497 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Nil, tmp12496)
}
__typedArg0 := Nil
__typedArg1 := tmp12496
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12498 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12497, Nil)
}
__typedArg0 := tmp12497
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12499 := Call(__e, PrimFunc(symshen_4comb), let__12150, tmp12498)


ifres12493 = tmp12499


}

ifres12488 = ifres12493


}

let__12145 := ifres12488
_ = let__12145

tmp12503 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12145)


if True == tmp12503 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__12145)
return
}


} else {
__e.Return(let__12136)
return
}


}, 1)

tmp12506 := Call(__e, ns2_1set, symshen_4_5formulae_6, tmp12465)


_ = tmp12506

tmp12507 := MakeNative(func(__e *ControlFlow) {
V3499 := __e.Get(1)
_ = V3499
tmp12508 := Call(__e, PrimFunc(symshen_4_5ass_6), V3499)


let__12152 := tmp12508
_ = let__12152

tmp12526 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12152)


var ifres12509 Obj

if True == tmp12526 {
tmp12510 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12509 = tmp12510


} else {
tmp12511 := Call(__e, PrimFunc(symshen_4_5_1out), let__12152)


let__12153 := tmp12511
_ = let__12153

tmp12512 := Call(__e, PrimFunc(symshen_4in_1_6), let__12152)


let__12154 := tmp12512
_ = let__12154

tmp12525 := Call(__e, PrimFunc(symshen_4hds_a_2), let__12154, sym_6_6)


var ifres12513 Obj

if True == tmp12525 {
tmp12514 := Call(__e, PrimFunc(symtail), let__12154)


let__12155 := tmp12514
_ = let__12155

tmp12515 := Call(__e, PrimFunc(symshen_4_5formula_6), let__12155)


let__12156 := tmp12515
_ = let__12156

tmp12523 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12156)


var ifres12516 Obj

if True == tmp12523 {
tmp12517 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12516 = tmp12517


} else {
tmp12518 := Call(__e, PrimFunc(symshen_4_5_1out), let__12156)


let__12157 := tmp12518
_ = let__12157

tmp12519 := Call(__e, PrimFunc(symshen_4in_1_6), let__12156)


let__12158 := tmp12519
_ = let__12158

tmp12520 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12157, Nil)
}
__typedArg0 := let__12157
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12521 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12153, tmp12520)
}
__typedArg0 := let__12153
__typedArg1 := tmp12520
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12522 := Call(__e, PrimFunc(symshen_4comb), let__12158, tmp12521)


ifres12516 = tmp12522


}

ifres12513 = ifres12516


} else {
tmp12524 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12513 = tmp12524


}

ifres12509 = ifres12513


}

let__12151 := ifres12509
_ = let__12151

tmp12539 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12151)


if True == tmp12539 {
tmp12527 := Call(__e, PrimFunc(symshen_4_5formula_6), V3499)


let__12160 := tmp12527
_ = let__12160

tmp12535 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12160)


var ifres12528 Obj

if True == tmp12535 {
tmp12529 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12528 = tmp12529


} else {
tmp12530 := Call(__e, PrimFunc(symshen_4_5_1out), let__12160)


let__12161 := tmp12530
_ = let__12161

tmp12531 := Call(__e, PrimFunc(symshen_4in_1_6), let__12160)


let__12162 := tmp12531
_ = let__12162

tmp12532 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12161, Nil)
}
__typedArg0 := let__12161
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12533 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Nil, tmp12532)
}
__typedArg0 := Nil
__typedArg1 := tmp12532
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12534 := Call(__e, PrimFunc(symshen_4comb), let__12162, tmp12533)


ifres12528 = tmp12534


}

let__12159 := ifres12528
_ = let__12159

tmp12537 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12159)


if True == tmp12537 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__12159)
return
}


} else {
__e.Return(let__12151)
return
}


}, 1)

tmp12540 := Call(__e, ns2_1set, symshen_4_5conc_6, tmp12507)


_ = tmp12540

tmp12541 := MakeNative(func(__e *ControlFlow) {
V3512 := __e.Get(1)
_ = V3512
tmp12542 := Call(__e, PrimFunc(symshen_4_5prem_6), V3512)


let__12164 := tmp12542
_ = let__12164

tmp12560 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12164)


var ifres12543 Obj

if True == tmp12560 {
tmp12544 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12543 = tmp12544


} else {
tmp12545 := Call(__e, PrimFunc(symshen_4_5_1out), let__12164)


let__12165 := tmp12545
_ = let__12165

tmp12546 := Call(__e, PrimFunc(symshen_4in_1_6), let__12164)


let__12166 := tmp12546
_ = let__12166

tmp12547 := Call(__e, PrimFunc(symshen_4_5sc_6), let__12166)


let__12167 := tmp12547
_ = let__12167

tmp12559 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12167)


var ifres12548 Obj

if True == tmp12559 {
tmp12549 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12548 = tmp12549


} else {
tmp12550 := Call(__e, PrimFunc(symshen_4in_1_6), let__12167)


let__12168 := tmp12550
_ = let__12168

tmp12551 := Call(__e, PrimFunc(symshen_4_5prems_6), let__12168)


let__12169 := tmp12551
_ = let__12169

tmp12558 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12169)


var ifres12552 Obj

if True == tmp12558 {
tmp12553 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12552 = tmp12553


} else {
tmp12554 := Call(__e, PrimFunc(symshen_4_5_1out), let__12169)


let__12170 := tmp12554
_ = let__12170

tmp12555 := Call(__e, PrimFunc(symshen_4in_1_6), let__12169)


let__12171 := tmp12555
_ = let__12171

tmp12556 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12165, let__12170)
}
__typedArg0 := let__12165
__typedArg1 := let__12170
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12557 := Call(__e, PrimFunc(symshen_4comb), let__12171, tmp12556)


ifres12552 = tmp12557


}

ifres12548 = ifres12552


}

ifres12543 = ifres12548


}

let__12163 := ifres12543
_ = let__12163

tmp12570 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12163)


if True == tmp12570 {
tmp12561 := Call(__e, PrimFunc(sym_5e_6), V3512)


let__12173 := tmp12561
_ = let__12173

tmp12566 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12173)


var ifres12562 Obj

if True == tmp12566 {
tmp12563 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12562 = tmp12563


} else {
tmp12564 := Call(__e, PrimFunc(symshen_4in_1_6), let__12173)


let__12174 := tmp12564
_ = let__12174

tmp12565 := Call(__e, PrimFunc(symshen_4comb), let__12174, Nil)


ifres12562 = tmp12565


}

let__12172 := ifres12562
_ = let__12172

tmp12568 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12172)


if True == tmp12568 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__12172)
return
}


} else {
__e.Return(let__12163)
return
}


}, 1)

tmp12571 := Call(__e, ns2_1set, symshen_4_5prems_6, tmp12541)


_ = tmp12571

tmp12572 := MakeNative(func(__e *ControlFlow) {
V3525 := __e.Get(1)
_ = V3525
tmp12577 := Call(__e, PrimFunc(symshen_4hds_a_2), V3525, sym_b)


var ifres12573 Obj

if True == tmp12577 {
tmp12574 := Call(__e, PrimFunc(symtail), V3525)


let__12176 := tmp12574
_ = let__12176

tmp12575 := Call(__e, PrimFunc(symshen_4comb), let__12176, sym_b)


ifres12573 = tmp12575


} else {
tmp12576 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12573 = tmp12576


}

let__12175 := ifres12573
_ = let__12175

tmp12611 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12175)


if True == tmp12611 {
tmp12578 := Call(__e, PrimFunc(symshen_4_5ass_6), V3525)


let__12178 := tmp12578
_ = let__12178

tmp12596 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12178)


var ifres12579 Obj

if True == tmp12596 {
tmp12580 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12579 = tmp12580


} else {
tmp12581 := Call(__e, PrimFunc(symshen_4_5_1out), let__12178)


let__12179 := tmp12581
_ = let__12179

tmp12582 := Call(__e, PrimFunc(symshen_4in_1_6), let__12178)


let__12180 := tmp12582
_ = let__12180

tmp12595 := Call(__e, PrimFunc(symshen_4hds_a_2), let__12180, sym_6_6)


var ifres12583 Obj

if True == tmp12595 {
tmp12584 := Call(__e, PrimFunc(symtail), let__12180)


let__12181 := tmp12584
_ = let__12181

tmp12585 := Call(__e, PrimFunc(symshen_4_5formula_6), let__12181)


let__12182 := tmp12585
_ = let__12182

tmp12593 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12182)


var ifres12586 Obj

if True == tmp12593 {
tmp12587 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12586 = tmp12587


} else {
tmp12588 := Call(__e, PrimFunc(symshen_4_5_1out), let__12182)


let__12183 := tmp12588
_ = let__12183

tmp12589 := Call(__e, PrimFunc(symshen_4in_1_6), let__12182)


let__12184 := tmp12589
_ = let__12184

tmp12590 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12183, Nil)
}
__typedArg0 := let__12183
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12591 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12179, tmp12590)
}
__typedArg0 := let__12179
__typedArg1 := tmp12590
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12592 := Call(__e, PrimFunc(symshen_4comb), let__12184, tmp12591)


ifres12586 = tmp12592


}

ifres12583 = ifres12586


} else {
tmp12594 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12583 = tmp12594


}

ifres12579 = ifres12583


}

let__12177 := ifres12579
_ = let__12177

tmp12609 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12177)


if True == tmp12609 {
tmp12597 := Call(__e, PrimFunc(symshen_4_5formula_6), V3525)


let__12186 := tmp12597
_ = let__12186

tmp12605 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12186)


var ifres12598 Obj

if True == tmp12605 {
tmp12599 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12598 = tmp12599


} else {
tmp12600 := Call(__e, PrimFunc(symshen_4_5_1out), let__12186)


let__12187 := tmp12600
_ = let__12187

tmp12601 := Call(__e, PrimFunc(symshen_4in_1_6), let__12186)


let__12188 := tmp12601
_ = let__12188

tmp12602 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12187, Nil)
}
__typedArg0 := let__12187
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12603 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Nil, tmp12602)
}
__typedArg0 := Nil
__typedArg1 := tmp12602
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12604 := Call(__e, PrimFunc(symshen_4comb), let__12188, tmp12603)


ifres12598 = tmp12604


}

let__12185 := ifres12598
_ = let__12185

tmp12607 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12185)


if True == tmp12607 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__12185)
return
}


} else {
__e.Return(let__12177)
return
}


} else {
__e.Return(let__12175)
return
}


}, 1)

tmp12612 := Call(__e, ns2_1set, symshen_4_5prem_6, tmp12572)


_ = tmp12612

tmp12613 := MakeNative(func(__e *ControlFlow) {
V3540 := __e.Get(1)
_ = V3540
tmp12614 := Call(__e, PrimFunc(symshen_4_5formula_6), V3540)


let__12190 := tmp12614
_ = let__12190

tmp12632 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12190)


var ifres12615 Obj

if True == tmp12632 {
tmp12616 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12615 = tmp12616


} else {
tmp12617 := Call(__e, PrimFunc(symshen_4_5_1out), let__12190)


let__12191 := tmp12617
_ = let__12191

tmp12618 := Call(__e, PrimFunc(symshen_4in_1_6), let__12190)


let__12192 := tmp12618
_ = let__12192

tmp12619 := Call(__e, PrimFunc(symshen_4_5iscomma_6), let__12192)


let__12193 := tmp12619
_ = let__12193

tmp12631 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12193)


var ifres12620 Obj

if True == tmp12631 {
tmp12621 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12620 = tmp12621


} else {
tmp12622 := Call(__e, PrimFunc(symshen_4in_1_6), let__12193)


let__12194 := tmp12622
_ = let__12194

tmp12623 := Call(__e, PrimFunc(symshen_4_5ass_6), let__12194)


let__12195 := tmp12623
_ = let__12195

tmp12630 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12195)


var ifres12624 Obj

if True == tmp12630 {
tmp12625 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12624 = tmp12625


} else {
tmp12626 := Call(__e, PrimFunc(symshen_4_5_1out), let__12195)


let__12196 := tmp12626
_ = let__12196

tmp12627 := Call(__e, PrimFunc(symshen_4in_1_6), let__12195)


let__12197 := tmp12627
_ = let__12197

tmp12628 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12191, let__12196)
}
__typedArg0 := let__12191
__typedArg1 := let__12196
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12629 := Call(__e, PrimFunc(symshen_4comb), let__12197, tmp12628)


ifres12624 = tmp12629


}

ifres12620 = ifres12624


}

ifres12615 = ifres12620


}

let__12189 := ifres12615
_ = let__12189

tmp12652 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12189)


if True == tmp12652 {
tmp12633 := Call(__e, PrimFunc(symshen_4_5formula_6), V3540)


let__12199 := tmp12633
_ = let__12199

tmp12640 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12199)


var ifres12634 Obj

if True == tmp12640 {
tmp12635 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12634 = tmp12635


} else {
tmp12636 := Call(__e, PrimFunc(symshen_4_5_1out), let__12199)


let__12200 := tmp12636
_ = let__12200

tmp12637 := Call(__e, PrimFunc(symshen_4in_1_6), let__12199)


let__12201 := tmp12637
_ = let__12201

tmp12638 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12200, Nil)
}
__typedArg0 := let__12200
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12639 := Call(__e, PrimFunc(symshen_4comb), let__12201, tmp12638)


ifres12634 = tmp12639


}

let__12198 := ifres12634
_ = let__12198

tmp12650 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12198)


if True == tmp12650 {
tmp12641 := Call(__e, PrimFunc(sym_5e_6), V3540)


let__12203 := tmp12641
_ = let__12203

tmp12646 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12203)


var ifres12642 Obj

if True == tmp12646 {
tmp12643 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12642 = tmp12643


} else {
tmp12644 := Call(__e, PrimFunc(symshen_4in_1_6), let__12203)


let__12204 := tmp12644
_ = let__12204

tmp12645 := Call(__e, PrimFunc(symshen_4comb), let__12204, Nil)


ifres12642 = tmp12645


}

let__12202 := ifres12642
_ = let__12202

tmp12648 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12202)


if True == tmp12648 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__12202)
return
}


} else {
__e.Return(let__12198)
return
}


} else {
__e.Return(let__12189)
return
}


}, 1)

tmp12653 := Call(__e, ns2_1set, symshen_4_5ass_6, tmp12613)


_ = tmp12653

tmp12654 := MakeNative(func(__e *ControlFlow) {
V3557 := __e.Get(1)
_ = V3557
tmp12664 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3557)
}
__typedArg0 := V3557
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12655 Obj

if True == tmp12664 {
tmp12656 := Call(__e, PrimFunc(symhead), V3557)


let__12206 := tmp12656
_ = let__12206

tmp12657 := Call(__e, PrimFunc(symtail), V3557)


let__12207 := tmp12657
_ = let__12207

tmp12661 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(","))
}
__typedArg0 := MakeString(",")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp12662 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__12206, tmp12661)
}
__typedArg0 := let__12206
__typedArg1 := tmp12661
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres12658 Obj

if True == tmp12662 {
tmp12659 := Call(__e, PrimFunc(symshen_4comb), let__12207, symshen_4skip)


ifres12658 = tmp12659


} else {
tmp12660 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12658 = tmp12660


}

ifres12655 = ifres12658


} else {
tmp12663 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12655 = tmp12663


}

let__12205 := ifres12655
_ = let__12205

tmp12666 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12205)


if True == tmp12666 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__12205)
return
}


}, 1)

tmp12667 := Call(__e, ns2_1set, symshen_4_5iscomma_6, tmp12654)


_ = tmp12667

tmp12668 := MakeNative(func(__e *ControlFlow) {
V3561 := __e.Get(1)
_ = V3561
tmp12669 := Call(__e, PrimFunc(symshen_4_5expr_6), V3561)


let__12209 := tmp12669
_ = let__12209

tmp12692 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12209)


var ifres12670 Obj

if True == tmp12692 {
tmp12671 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12670 = tmp12671


} else {
tmp12672 := Call(__e, PrimFunc(symshen_4_5_1out), let__12209)


let__12210 := tmp12672
_ = let__12210

tmp12673 := Call(__e, PrimFunc(symshen_4in_1_6), let__12209)


let__12211 := tmp12673
_ = let__12211

tmp12674 := Call(__e, PrimFunc(symshen_4_5iscolon_6), let__12211)


let__12212 := tmp12674
_ = let__12212

tmp12691 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12212)


var ifres12675 Obj

if True == tmp12691 {
tmp12676 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12675 = tmp12676


} else {
tmp12677 := Call(__e, PrimFunc(symshen_4in_1_6), let__12212)


let__12213 := tmp12677
_ = let__12213

tmp12678 := Call(__e, PrimFunc(symshen_4_5type_6), let__12213)


let__12214 := tmp12678
_ = let__12214

tmp12690 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12214)


var ifres12679 Obj

if True == tmp12690 {
tmp12680 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12679 = tmp12680


} else {
tmp12681 := Call(__e, PrimFunc(symshen_4_5_1out), let__12214)


let__12215 := tmp12681
_ = let__12215

tmp12682 := Call(__e, PrimFunc(symshen_4in_1_6), let__12214)


let__12216 := tmp12682
_ = let__12216

tmp12683 := Call(__e, PrimFunc(symshen_4curry), let__12210)


tmp12684 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp12685 := Call(__e, PrimFunc(symshen_4rectify_1type), let__12215)


tmp12686 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12685, Nil)
}
__typedArg0 := tmp12685
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12687 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12684, tmp12686)
}
__typedArg0 := tmp12684
__typedArg1 := tmp12686
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12688 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12683, tmp12687)
}
__typedArg0 := tmp12683
__typedArg1 := tmp12687
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12689 := Call(__e, PrimFunc(symshen_4comb), let__12216, tmp12688)


ifres12679 = tmp12689


}

ifres12675 = ifres12679


}

ifres12670 = ifres12675


}

let__12208 := ifres12670
_ = let__12208

tmp12703 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12208)


if True == tmp12703 {
tmp12693 := Call(__e, PrimFunc(symshen_4_5expr_6), V3561)


let__12218 := tmp12693
_ = let__12218

tmp12699 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12218)


var ifres12694 Obj

if True == tmp12699 {
tmp12695 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12694 = tmp12695


} else {
tmp12696 := Call(__e, PrimFunc(symshen_4_5_1out), let__12218)


let__12219 := tmp12696
_ = let__12219

tmp12697 := Call(__e, PrimFunc(symshen_4in_1_6), let__12218)


let__12220 := tmp12697
_ = let__12220

tmp12698 := Call(__e, PrimFunc(symshen_4comb), let__12220, let__12219)


ifres12694 = tmp12698


}

let__12217 := ifres12694
_ = let__12217

tmp12701 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12217)


if True == tmp12701 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__12217)
return
}


} else {
__e.Return(let__12208)
return
}


}, 1)

tmp12704 := Call(__e, ns2_1set, symshen_4_5formula_6, tmp12668)


_ = tmp12704

tmp12705 := MakeNative(func(__e *ControlFlow) {
V3575 := __e.Get(1)
_ = V3575
tmp12715 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3575)
}
__typedArg0 := V3575
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12706 Obj

if True == tmp12715 {
tmp12707 := Call(__e, PrimFunc(symhead), V3575)


let__12222 := tmp12707
_ = let__12222

tmp12708 := Call(__e, PrimFunc(symtail), V3575)


let__12223 := tmp12708
_ = let__12223

tmp12712 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp12713 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__12222, tmp12712)
}
__typedArg0 := let__12222
__typedArg1 := tmp12712
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres12709 Obj

if True == tmp12713 {
tmp12710 := Call(__e, PrimFunc(symshen_4comb), let__12223, symshen_4skip)


ifres12709 = tmp12710


} else {
tmp12711 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12709 = tmp12711


}

ifres12706 = ifres12709


} else {
tmp12714 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12706 = tmp12714


}

let__12221 := ifres12706
_ = let__12221

tmp12717 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12221)


if True == tmp12717 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__12221)
return
}


}, 1)

tmp12718 := Call(__e, ns2_1set, symshen_4_5iscolon_6, tmp12705)


_ = tmp12718

tmp12719 := MakeNative(func(__e *ControlFlow) {
V3579 := __e.Get(1)
_ = V3579
tmp12720 := Call(__e, PrimFunc(symshen_4_5side_6), V3579)


let__12225 := tmp12720
_ = let__12225

tmp12733 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12225)


var ifres12721 Obj

if True == tmp12733 {
tmp12722 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12721 = tmp12722


} else {
tmp12723 := Call(__e, PrimFunc(symshen_4_5_1out), let__12225)


let__12226 := tmp12723
_ = let__12226

tmp12724 := Call(__e, PrimFunc(symshen_4in_1_6), let__12225)


let__12227 := tmp12724
_ = let__12227

tmp12725 := Call(__e, PrimFunc(symshen_4_5sides_6), let__12227)


let__12228 := tmp12725
_ = let__12228

tmp12732 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12228)


var ifres12726 Obj

if True == tmp12732 {
tmp12727 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12726 = tmp12727


} else {
tmp12728 := Call(__e, PrimFunc(symshen_4_5_1out), let__12228)


let__12229 := tmp12728
_ = let__12229

tmp12729 := Call(__e, PrimFunc(symshen_4in_1_6), let__12228)


let__12230 := tmp12729
_ = let__12230

tmp12730 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12226, let__12229)
}
__typedArg0 := let__12226
__typedArg1 := let__12229
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12731 := Call(__e, PrimFunc(symshen_4comb), let__12230, tmp12730)


ifres12726 = tmp12731


}

ifres12721 = ifres12726


}

let__12224 := ifres12721
_ = let__12224

tmp12743 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12224)


if True == tmp12743 {
tmp12734 := Call(__e, PrimFunc(sym_5e_6), V3579)


let__12232 := tmp12734
_ = let__12232

tmp12739 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12232)


var ifres12735 Obj

if True == tmp12739 {
tmp12736 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12735 = tmp12736


} else {
tmp12737 := Call(__e, PrimFunc(symshen_4in_1_6), let__12232)


let__12233 := tmp12737
_ = let__12233

tmp12738 := Call(__e, PrimFunc(symshen_4comb), let__12233, Nil)


ifres12735 = tmp12738


}

let__12231 := ifres12735
_ = let__12231

tmp12741 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12231)


if True == tmp12741 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__12231)
return
}


} else {
__e.Return(let__12224)
return
}


}, 1)

tmp12744 := Call(__e, ns2_1set, symshen_4_5sides_6, tmp12719)


_ = tmp12744

tmp12745 := MakeNative(func(__e *ControlFlow) {
V3590 := __e.Get(1)
_ = V3590
tmp12757 := Call(__e, PrimFunc(symshen_4hds_a_2), V3590, symif)


var ifres12746 Obj

if True == tmp12757 {
tmp12747 := Call(__e, PrimFunc(symtail), V3590)


let__12235 := tmp12747
_ = let__12235

tmp12755 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__12235)
}
__typedArg0 := let__12235
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12748 Obj

if True == tmp12755 {
tmp12749 := Call(__e, PrimFunc(symhead), let__12235)


let__12236 := tmp12749
_ = let__12236

tmp12750 := Call(__e, PrimFunc(symtail), let__12235)


let__12237 := tmp12750
_ = let__12237

tmp12751 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12236, Nil)
}
__typedArg0 := let__12236
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12752 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp12751)
}
__typedArg0 := symif
__typedArg1 := tmp12751
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12753 := Call(__e, PrimFunc(symshen_4comb), let__12237, tmp12752)


ifres12748 = tmp12753


} else {
tmp12754 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12748 = tmp12754


}

ifres12746 = ifres12748


} else {
tmp12756 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12746 = tmp12756


}

let__12234 := ifres12746
_ = let__12234

tmp12813 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12234)


if True == tmp12813 {
tmp12775 := Call(__e, PrimFunc(symshen_4hds_a_2), V3590, symlet)


var ifres12758 Obj

if True == tmp12775 {
tmp12759 := Call(__e, PrimFunc(symtail), V3590)


let__12239 := tmp12759
_ = let__12239

tmp12773 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__12239)
}
__typedArg0 := let__12239
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12760 Obj

if True == tmp12773 {
tmp12761 := Call(__e, PrimFunc(symhead), let__12239)


let__12240 := tmp12761
_ = let__12240

tmp12762 := Call(__e, PrimFunc(symtail), let__12239)


let__12241 := tmp12762
_ = let__12241

tmp12771 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__12241)
}
__typedArg0 := let__12241
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12763 Obj

if True == tmp12771 {
tmp12764 := Call(__e, PrimFunc(symhead), let__12241)


let__12242 := tmp12764
_ = let__12242

tmp12765 := Call(__e, PrimFunc(symtail), let__12241)


let__12243 := tmp12765
_ = let__12243

tmp12766 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12242, Nil)
}
__typedArg0 := let__12242
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12767 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12240, tmp12766)
}
__typedArg0 := let__12240
__typedArg1 := tmp12766
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12768 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp12767)
}
__typedArg0 := symlet
__typedArg1 := tmp12767
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12769 := Call(__e, PrimFunc(symshen_4comb), let__12243, tmp12768)


ifres12763 = tmp12769


} else {
tmp12770 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12763 = tmp12770


}

ifres12760 = ifres12763


} else {
tmp12772 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12760 = tmp12772


}

ifres12758 = ifres12760


} else {
tmp12774 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12758 = tmp12774


}

let__12238 := ifres12758
_ = let__12238

tmp12811 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12238)


if True == tmp12811 {
tmp12790 := Call(__e, PrimFunc(symshen_4hds_a_2), V3590, symctxt)


var ifres12776 Obj

if True == tmp12790 {
tmp12777 := Call(__e, PrimFunc(symtail), V3590)


let__12245 := tmp12777
_ = let__12245

tmp12788 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__12245)
}
__typedArg0 := let__12245
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12778 Obj

if True == tmp12788 {
tmp12779 := Call(__e, PrimFunc(symhead), let__12245)


let__12246 := tmp12779
_ = let__12246

tmp12780 := Call(__e, PrimFunc(symtail), let__12245)


let__12247 := tmp12780
_ = let__12247

tmp12786 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(let__12246)
}
__typedArg0 := let__12246
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

var ifres12781 Obj

if True == tmp12786 {
tmp12782 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12246, Nil)
}
__typedArg0 := let__12246
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12783 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symctxt, tmp12782)
}
__typedArg0 := symctxt
__typedArg1 := tmp12782
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12784 := Call(__e, PrimFunc(symshen_4comb), let__12247, tmp12783)


ifres12781 = tmp12784


} else {
tmp12785 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12781 = tmp12785


}

ifres12778 = ifres12781


} else {
tmp12787 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12778 = tmp12787


}

ifres12776 = ifres12778


} else {
tmp12789 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12776 = tmp12789


}

let__12244 := ifres12776
_ = let__12244

tmp12809 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12244)


if True == tmp12809 {
tmp12805 := Call(__e, PrimFunc(symshen_4hds_a_2), V3590, symsqts)


var ifres12791 Obj

if True == tmp12805 {
tmp12792 := Call(__e, PrimFunc(symtail), V3590)


let__12249 := tmp12792
_ = let__12249

tmp12803 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__12249)
}
__typedArg0 := let__12249
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12793 Obj

if True == tmp12803 {
tmp12794 := Call(__e, PrimFunc(symhead), let__12249)


let__12250 := tmp12794
_ = let__12250

tmp12795 := Call(__e, PrimFunc(symtail), let__12249)


let__12251 := tmp12795
_ = let__12251

tmp12801 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(let__12250)
}
__typedArg0 := let__12250
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

var ifres12796 Obj

if True == tmp12801 {
tmp12797 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12250, Nil)
}
__typedArg0 := let__12250
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12798 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsqts, tmp12797)
}
__typedArg0 := symsqts
__typedArg1 := tmp12797
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12799 := Call(__e, PrimFunc(symshen_4comb), let__12251, tmp12798)


ifres12796 = tmp12799


} else {
tmp12800 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12796 = tmp12800


}

ifres12793 = ifres12796


} else {
tmp12802 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12793 = tmp12802


}

ifres12791 = ifres12793


} else {
tmp12804 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12791 = tmp12804


}

let__12248 := ifres12791
_ = let__12248

tmp12807 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12248)


if True == tmp12807 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__12248)
return
}


} else {
__e.Return(let__12244)
return
}


} else {
__e.Return(let__12238)
return
}


} else {
__e.Return(let__12234)
return
}


}, 1)

tmp12814 := Call(__e, ns2_1set, symshen_4_5side_6, tmp12745)


_ = tmp12814

tmp12815 := MakeNative(func(__e *ControlFlow) {
V3615 := __e.Get(1)
_ = V3615
V3616 := __e.Get(2)
_ = V3616
V3617 := __e.Get(3)
_ = V3617
tmp12845 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3617)
}
__typedArg0 := V3617
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12832 Obj

if True == tmp12845 {
tmp12843 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3617)
}
__typedArg0 := V3617
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12844 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp12843)
}
__typedArg0 := Nil
__typedArg1 := tmp12843
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres12834 Obj

if True == tmp12844 {
tmp12841 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3617)
}
__typedArg0 := V3617
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12842 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp12841)
}
__typedArg0 := tmp12841
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12836 Obj

if True == tmp12842 {
tmp12838 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3617)
}
__typedArg0 := V3617
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12839 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp12838)
}
__typedArg0 := tmp12838
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12840 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp12839)
}
__typedArg0 := Nil
__typedArg1 := tmp12839
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres12837 Obj

if True == tmp12840 {
ifres12837 = True


} else {
ifres12837 = False


}

ifres12836 = ifres12837


} else {
ifres12836 = False


}

var ifres12835 Obj

if True == ifres12836 {
ifres12835 = True


} else {
ifres12835 = False


}

ifres12834 = ifres12835


} else {
ifres12834 = False


}

var ifres12833 Obj

if True == ifres12834 {
ifres12833 = True


} else {
ifres12833 = False


}

ifres12832 = ifres12833


} else {
ifres12832 = False


}

if True == ifres12832 {
tmp12816 := Call(__e, PrimFunc(symgensym), symP)


let__12252 := tmp12816
_ = let__12252

tmp12817 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3617)
}
__typedArg0 := V3617
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12818 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12252, Nil)
}
__typedArg0 := let__12252
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12819 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12817, tmp12818)
}
__typedArg0 := tmp12817
__typedArg1 := tmp12818
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__12253 := tmp12819
_ = let__12253

tmp12820 := Call(__e, PrimFunc(symshen_4coll_1formulae), V3616)


tmp12821 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12252, Nil)
}
__typedArg0 := let__12252
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12822 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12820, tmp12821)
}
__typedArg0 := tmp12820
__typedArg1 := tmp12821
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__12254 := tmp12822
_ = let__12254

tmp12823 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12254, Nil)
}
__typedArg0 := let__12254
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12824 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12253, Nil)
}
__typedArg0 := let__12253
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12825 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12823, tmp12824)
}
__typedArg0 := tmp12823
__typedArg1 := tmp12824
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12826 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3615, tmp12825)
}
__typedArg0 := V3615
__typedArg1 := tmp12825
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__12255 := tmp12826
_ = let__12255

tmp12827 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3617, Nil)
}
__typedArg0 := V3617
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12828 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3616, tmp12827)
}
__typedArg0 := V3616
__typedArg1 := tmp12827
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12829 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3615, tmp12828)
}
__typedArg0 := V3615
__typedArg1 := tmp12828
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__12256 := tmp12829
_ = let__12256

tmp12830 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12255, Nil)
}
__typedArg0 := let__12255
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12256, tmp12830)
}
__typedArg0 := let__12256
__typedArg1 := tmp12830
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.lr-rule"))
}
__typedArg0 := MakeString("implementation error in shen.lr-rule")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}, 3)

tmp12846 := Call(__e, ns2_1set, symshen_4lr_1rule, tmp12815)


_ = tmp12846

tmp12847 := MakeNative(func(__e *ControlFlow) {
V3625 := __e.Get(1)
_ = V3625
tmp12876 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3625)
}
__typedArg0 := Nil
__typedArg1 := V3625
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp12876 {
__e.Return(Nil)
return
} else {
tmp12874 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3625)
}
__typedArg0 := V3625
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12854 Obj

if True == tmp12874 {
tmp12872 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3625)
}
__typedArg0 := V3625
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12873 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp12872)
}
__typedArg0 := tmp12872
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12856 Obj

if True == tmp12873 {
tmp12869 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3625)
}
__typedArg0 := V3625
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12870 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp12869)
}
__typedArg0 := tmp12869
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12871 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp12870)
}
__typedArg0 := Nil
__typedArg1 := tmp12870
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres12858 Obj

if True == tmp12871 {
tmp12866 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3625)
}
__typedArg0 := V3625
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12867 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp12866)
}
__typedArg0 := tmp12866
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12868 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp12867)
}
__typedArg0 := tmp12867
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12860 Obj

if True == tmp12868 {
tmp12862 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3625)
}
__typedArg0 := V3625
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12863 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp12862)
}
__typedArg0 := tmp12862
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12864 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp12863)
}
__typedArg0 := tmp12863
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12865 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp12864)
}
__typedArg0 := Nil
__typedArg1 := tmp12864
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres12861 Obj

if True == tmp12865 {
ifres12861 = True


} else {
ifres12861 = False


}

ifres12860 = ifres12861


} else {
ifres12860 = False


}

var ifres12859 Obj

if True == ifres12860 {
ifres12859 = True


} else {
ifres12859 = False


}

ifres12858 = ifres12859


} else {
ifres12858 = False


}

var ifres12857 Obj

if True == ifres12858 {
ifres12857 = True


} else {
ifres12857 = False


}

ifres12856 = ifres12857


} else {
ifres12856 = False


}

var ifres12855 Obj

if True == ifres12856 {
ifres12855 = True


} else {
ifres12855 = False


}

ifres12854 = ifres12855


} else {
ifres12854 = False


}

if True == ifres12854 {
tmp12848 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3625)
}
__typedArg0 := V3625
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12849 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp12848)
}
__typedArg0 := tmp12848
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12850 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp12849)
}
__typedArg0 := tmp12849
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12851 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3625)
}
__typedArg0 := V3625
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12852 := Call(__e, PrimFunc(symshen_4coll_1formulae), tmp12851)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12850, tmp12852)
}
__typedArg0 := tmp12850
__typedArg1 := tmp12852
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.coll-formulae"))
}
__typedArg0 := MakeString("implementation error in shen.coll-formulae")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 1)

tmp12877 := Call(__e, ns2_1set, symshen_4coll_1formulae, tmp12847)


_ = tmp12877

tmp12878 := MakeNative(func(__e *ControlFlow) {
V3626 := __e.Get(1)
_ = V3626
tmp12889 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3626)
}
__typedArg0 := V3626
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12879 Obj

if True == tmp12889 {
tmp12880 := Call(__e, PrimFunc(symhead), V3626)


let__12258 := tmp12880
_ = let__12258

tmp12881 := Call(__e, PrimFunc(symtail), V3626)


let__12259 := tmp12881
_ = let__12259

tmp12886 := Call(__e, PrimFunc(symshen_4key_1in_1sequent_1calculus_2), let__12258)


tmp12887 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp12886)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp12886
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres12882 Obj

if True == tmp12887 {
tmp12883 := Call(__e, PrimFunc(symmacroexpand), let__12258)


tmp12884 := Call(__e, PrimFunc(symshen_4comb), let__12259, tmp12883)


ifres12882 = tmp12884


} else {
tmp12885 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12882 = tmp12885


}

ifres12879 = ifres12882


} else {
tmp12888 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12879 = tmp12888


}

let__12257 := ifres12879
_ = let__12257

tmp12891 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12257)


if True == tmp12891 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__12257)
return
}


}, 1)

tmp12892 := Call(__e, ns2_1set, symshen_4_5expr_6, tmp12878)


_ = tmp12892

tmp12893 := MakeNative(func(__e *ControlFlow) {
V3630 := __e.Get(1)
_ = V3630
tmp12900 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(";"))
}
__typedArg0 := MakeString(";")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp12901 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(","))
}
__typedArg0 := MakeString(",")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp12902 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp12903 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1_1, Nil)
}
__typedArg0 := sym_5_1_1
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12904 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12902, tmp12903)
}
__typedArg0 := tmp12902
__typedArg1 := tmp12903
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12905 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12901, tmp12904)
}
__typedArg0 := tmp12901
__typedArg1 := tmp12904
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12906 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12900, tmp12905)
}
__typedArg0 := tmp12900
__typedArg1 := tmp12905
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12907 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_6_6, tmp12906)
}
__typedArg0 := sym_6_6
__typedArg1 := tmp12906
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12908 := Call(__e, PrimFunc(symelement_2), V3630, tmp12907)


if True == tmp12908 {
__e.Return(True)
return
} else {
tmp12898 := Call(__e, PrimFunc(symshen_4sng_2), V3630)


var ifres12895 Obj

if True == tmp12898 {
ifres12895 = True


} else {
tmp12897 := Call(__e, PrimFunc(symshen_4dbl_2), V3630)


var ifres12896 Obj

if True == tmp12897 {
ifres12896 = True


} else {
ifres12896 = False


}

ifres12895 = ifres12896


}

if True == ifres12895 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp12909 := Call(__e, ns2_1set, symshen_4key_1in_1sequent_1calculus_2, tmp12893)


_ = tmp12909

tmp12910 := MakeNative(func(__e *ControlFlow) {
V3631 := __e.Get(1)
_ = V3631
tmp12911 := Call(__e, PrimFunc(symshen_4_5expr_6), V3631)


let__12261 := tmp12911
_ = let__12261

tmp12917 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12261)


var ifres12912 Obj

if True == tmp12917 {
tmp12913 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12912 = tmp12913


} else {
tmp12914 := Call(__e, PrimFunc(symshen_4_5_1out), let__12261)


let__12262 := tmp12914
_ = let__12262

tmp12915 := Call(__e, PrimFunc(symshen_4in_1_6), let__12261)


let__12263 := tmp12915
_ = let__12263

tmp12916 := Call(__e, PrimFunc(symshen_4comb), let__12263, let__12262)


ifres12912 = tmp12916


}

let__12260 := ifres12912
_ = let__12260

tmp12919 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12260)


if True == tmp12919 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__12260)
return
}


}, 1)

tmp12920 := Call(__e, ns2_1set, symshen_4_5type_6, tmp12910)


_ = tmp12920

tmp12921 := MakeNative(func(__e *ControlFlow) {
V3636 := __e.Get(1)
_ = V3636
tmp12930 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3636)
}
__typedArg0 := V3636
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12922 Obj

if True == tmp12930 {
tmp12923 := Call(__e, PrimFunc(symhead), V3636)


let__12265 := tmp12923
_ = let__12265

tmp12924 := Call(__e, PrimFunc(symtail), V3636)


let__12266 := tmp12924
_ = let__12266

tmp12928 := Call(__e, PrimFunc(symshen_4dbl_2), let__12265)


var ifres12925 Obj

if True == tmp12928 {
tmp12926 := Call(__e, PrimFunc(symshen_4comb), let__12266, let__12265)


ifres12925 = tmp12926


} else {
tmp12927 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12925 = tmp12927


}

ifres12922 = ifres12925


} else {
tmp12929 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12922 = tmp12929


}

let__12264 := ifres12922
_ = let__12264

tmp12932 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12264)


if True == tmp12932 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__12264)
return
}


}, 1)

tmp12933 := Call(__e, ns2_1set, symshen_4_5dbl_6, tmp12921)


_ = tmp12933

tmp12934 := MakeNative(func(__e *ControlFlow) {
V3640 := __e.Get(1)
_ = V3640
tmp12943 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3640)
}
__typedArg0 := V3640
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12935 Obj

if True == tmp12943 {
tmp12936 := Call(__e, PrimFunc(symhead), V3640)


let__12268 := tmp12936
_ = let__12268

tmp12937 := Call(__e, PrimFunc(symtail), V3640)


let__12269 := tmp12937
_ = let__12269

tmp12941 := Call(__e, PrimFunc(symshen_4sng_2), let__12268)


var ifres12938 Obj

if True == tmp12941 {
tmp12939 := Call(__e, PrimFunc(symshen_4comb), let__12269, let__12268)


ifres12938 = tmp12939


} else {
tmp12940 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12938 = tmp12940


}

ifres12935 = ifres12938


} else {
tmp12942 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12935 = tmp12942


}

let__12267 := ifres12935
_ = let__12267

tmp12945 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12267)


if True == tmp12945 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__12267)
return
}


}, 1)

tmp12946 := Call(__e, ns2_1set, symshen_4_5sng_6, tmp12934)


_ = tmp12946

tmp12947 := MakeNative(func(__e *ControlFlow) {
V3644 := __e.Get(1)
_ = V3644
tmp12952 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsymbol_2) {
return PrimIsSymbol(V3644)
}
__typedArg0 := V3644
return Call(__e, PrimFunc(symsymbol_2), __typedArg0)
})()

if True == tmp12952 {
tmp12949 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(V3644)
}
__typedArg0 := V3644
return Call(__e, PrimFunc(symstr), __typedArg0)
})()

tmp12950 := Call(__e, PrimFunc(symshen_4sng_1h_2), tmp12949)


if True == tmp12950 {
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

tmp12953 := Call(__e, ns2_1set, symshen_4sng_2, tmp12947)


_ = tmp12953

tmp12954 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__selftop:
V3647 := __self1
_ = V3647
tmp12963 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("___"), V3647)
}
__typedArg0 := MakeString("___")
__typedArg1 := V3647
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp12963 {
__e.Return(True)
return
} else {
tmp12961 := Call(__e, PrimFunc(symshen_4_7string_2), V3647)


var ifres12957 Obj

if True == tmp12961 {
tmp12959 := Call(__e, PrimFunc(symhdstr), V3647)


tmp12960 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("_"), tmp12959)
}
__typedArg0 := MakeString("_")
__typedArg1 := tmp12959
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres12958 Obj

if True == tmp12960 {
ifres12958 = True


} else {
ifres12958 = False


}

ifres12957 = ifres12958


} else {
ifres12957 = False


}

if True == ifres12957 {
if PrimFunc(symshen_4sng_1h_2) == __self {
__self1 = (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V3647)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V3647
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4sng_1h_2), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V3647)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V3647
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp12964 := Call(__e, ns2_1set, symshen_4sng_1h_2, tmp12954)


_ = tmp12964

tmp12965 := MakeNative(func(__e *ControlFlow) {
V3648 := __e.Get(1)
_ = V3648
tmp12970 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsymbol_2) {
return PrimIsSymbol(V3648)
}
__typedArg0 := V3648
return Call(__e, PrimFunc(symsymbol_2), __typedArg0)
})()

if True == tmp12970 {
tmp12967 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(V3648)
}
__typedArg0 := V3648
return Call(__e, PrimFunc(symstr), __typedArg0)
})()

tmp12968 := Call(__e, PrimFunc(symshen_4dbl_1h_2), tmp12967)


if True == tmp12968 {
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

tmp12971 := Call(__e, ns2_1set, symshen_4dbl_2, tmp12965)


_ = tmp12971

tmp12972 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__selftop:
V3651 := __self1
_ = V3651
tmp12981 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("==="), V3651)
}
__typedArg0 := MakeString("===")
__typedArg1 := V3651
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp12981 {
__e.Return(True)
return
} else {
tmp12979 := Call(__e, PrimFunc(symshen_4_7string_2), V3651)


var ifres12975 Obj

if True == tmp12979 {
tmp12977 := Call(__e, PrimFunc(symhdstr), V3651)


tmp12978 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("="), tmp12977)
}
__typedArg0 := MakeString("=")
__typedArg1 := tmp12977
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres12976 Obj

if True == tmp12978 {
ifres12976 = True


} else {
ifres12976 = False


}

ifres12975 = ifres12976


} else {
ifres12975 = False


}

if True == ifres12975 {
if PrimFunc(symshen_4dbl_1h_2) == __self {
__self1 = (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V3651)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V3651
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4dbl_1h_2), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V3651)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V3651
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp12982 := Call(__e, ns2_1set, symshen_4dbl_1h_2, tmp12972)


_ = tmp12982

tmp12983 := MakeNative(func(__e *ControlFlow) {
V3652 := __e.Get(1)
_ = V3652
V3653 := __e.Get(2)
_ = V3653
tmp12984 := MakeNative(func(__e *ControlFlow) {
Z3655 := __e.Get(1)
_ = Z3655
__e.TailApply(PrimFunc(symshen_4rule_1_6clause), Z3655)
return
}, 1)

tmp12985 := Call(__e, PrimFunc(symmapcan), tmp12984, V3653)


let__12270 := tmp12985
_ = let__12270

tmp12986 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3652, let__12270)
}
__typedArg0 := V3652
__typedArg1 := let__12270
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12987 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefprolog, tmp12986)
}
__typedArg0 := symdefprolog
__typedArg1 := tmp12986
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__12271 := tmp12987
_ = let__12271

__e.TailApply(PrimFunc(symeval), let__12271)
return


}, 2)

tmp12988 := Call(__e, ns2_1set, symshen_4rules_1_6prolog, tmp12983)


_ = tmp12988

tmp12989 := MakeNative(func(__e *ControlFlow) {
V3657 := __e.Get(1)
_ = V3657
tmp13049 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3657)
}
__typedArg0 := V3657
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13013 Obj

if True == tmp13049 {
tmp13047 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3657)
}
__typedArg0 := V3657
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13048 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13047)
}
__typedArg0 := tmp13047
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13015 Obj

if True == tmp13048 {
tmp13044 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3657)
}
__typedArg0 := V3657
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13045 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13044)
}
__typedArg0 := tmp13044
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13046 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13045)
}
__typedArg0 := tmp13045
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13017 Obj

if True == tmp13046 {
tmp13040 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3657)
}
__typedArg0 := V3657
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13041 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13040)
}
__typedArg0 := tmp13040
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13042 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13041)
}
__typedArg0 := tmp13041
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13043 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13042)
}
__typedArg0 := tmp13042
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13019 Obj

if True == tmp13043 {
tmp13035 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3657)
}
__typedArg0 := V3657
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13036 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13035)
}
__typedArg0 := tmp13035
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13037 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13036)
}
__typedArg0 := tmp13036
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13038 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13037)
}
__typedArg0 := tmp13037
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13039 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13038)
}
__typedArg0 := tmp13038
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13021 Obj

if True == tmp13039 {
tmp13029 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3657)
}
__typedArg0 := V3657
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13030 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13029)
}
__typedArg0 := tmp13029
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13031 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13030)
}
__typedArg0 := tmp13030
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13032 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13031)
}
__typedArg0 := tmp13031
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13033 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13032)
}
__typedArg0 := tmp13032
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13034 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp13033)
}
__typedArg0 := Nil
__typedArg1 := tmp13033
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13023 Obj

if True == tmp13034 {
tmp13025 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3657)
}
__typedArg0 := V3657
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13026 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13025)
}
__typedArg0 := tmp13025
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13027 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13026)
}
__typedArg0 := tmp13026
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13028 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp13027)
}
__typedArg0 := Nil
__typedArg1 := tmp13027
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13024 Obj

if True == tmp13028 {
ifres13024 = True


} else {
ifres13024 = False


}

ifres13023 = ifres13024


} else {
ifres13023 = False


}

var ifres13022 Obj

if True == ifres13023 {
ifres13022 = True


} else {
ifres13022 = False


}

ifres13021 = ifres13022


} else {
ifres13021 = False


}

var ifres13020 Obj

if True == ifres13021 {
ifres13020 = True


} else {
ifres13020 = False


}

ifres13019 = ifres13020


} else {
ifres13019 = False


}

var ifres13018 Obj

if True == ifres13019 {
ifres13018 = True


} else {
ifres13018 = False


}

ifres13017 = ifres13018


} else {
ifres13017 = False


}

var ifres13016 Obj

if True == ifres13017 {
ifres13016 = True


} else {
ifres13016 = False


}

ifres13015 = ifres13016


} else {
ifres13015 = False


}

var ifres13014 Obj

if True == ifres13015 {
ifres13014 = True


} else {
ifres13014 = False


}

ifres13013 = ifres13014


} else {
ifres13013 = False


}

if True == ifres13013 {
tmp12990 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3657)
}
__typedArg0 := V3657
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12991 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp12990)
}
__typedArg0 := tmp12990
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12992 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp12991)
}
__typedArg0 := tmp12991
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12993 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp12992)
}
__typedArg0 := tmp12992
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12994 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp12993)
}
__typedArg0 := tmp12993
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12995 := Call(__e, PrimFunc(symshen_4extract_1vars), tmp12994)


let__12272 := tmp12995
_ = let__12272

tmp12996 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3657)
}
__typedArg0 := V3657
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12997 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp12996)
}
__typedArg0 := tmp12996
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12998 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp12997)
}
__typedArg0 := tmp12997
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12999 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp12998)
}
__typedArg0 := tmp12998
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13000 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp12999)
}
__typedArg0 := tmp12999
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13001 := Call(__e, PrimFunc(symshen_4rule_1_6head), tmp13000)


tmp13002 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1_1, Nil)
}
__typedArg0 := sym_5_1_1
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13003 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3657)
}
__typedArg0 := V3657
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13004 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3657)
}
__typedArg0 := V3657
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13005 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13004)
}
__typedArg0 := tmp13004
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13006 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3657)
}
__typedArg0 := V3657
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13007 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13006)
}
__typedArg0 := tmp13006
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13008 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13007)
}
__typedArg0 := tmp13007
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13009 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13008)
}
__typedArg0 := tmp13008
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13010 := Call(__e, PrimFunc(symshen_4rule_1_6body), let__12272, symAssumptions, tmp13003, tmp13005, tmp13009)


tmp13011 := Call(__e, PrimFunc(symappend), tmp13002, tmp13010)


__e.TailApply(PrimFunc(symappend), tmp13001, tmp13011)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4rule_1_6clause)
return
}


}, 1)

tmp13050 := Call(__e, ns2_1set, symshen_4rule_1_6clause, tmp12989)


_ = tmp13050

tmp13051 := MakeNative(func(__e *ControlFlow) {
V3659 := __e.Get(1)
_ = V3659
tmp13052 := Call(__e, PrimFunc(symshen_4macro_1_8ch), V3659)


tmp13053 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symAssumptions, Nil)
}
__typedArg0 := symAssumptions
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13052, tmp13053)
}
__typedArg0 := tmp13052
__typedArg1 := tmp13053
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp13054 := Call(__e, ns2_1set, symshen_4rule_1_6head, tmp13051)


_ = tmp13054

tmp13055 := MakeNative(func(__e *ControlFlow) {
V3660 := __e.Get(1)
_ = V3660
tmp13056 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3660, Nil)
}
__typedArg0 := V3660
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4_8ch, tmp13056)
}
__typedArg0 := symshen_4_8ch
__typedArg1 := tmp13056
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp13057 := Call(__e, ns2_1set, symshen_4macro_1_8ch, tmp13055)


_ = tmp13057

tmp13058 := MakeNative(func(__e *ControlFlow) {
V3661 := __e.Get(1)
_ = V3661
tmp13059 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3661, Nil)
}
__typedArg0 := V3661
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4_8c, tmp13059)
}
__typedArg0 := symshen_4_8c
__typedArg1 := tmp13059
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp13060 := Call(__e, ns2_1set, symshen_4macro_1_8c, tmp13058)


_ = tmp13060

tmp13061 := MakeNative(func(__e *ControlFlow) {
V3662 := __e.Get(1)
_ = V3662
V3663 := __e.Get(2)
_ = V3663
V3664 := __e.Get(3)
_ = V3664
V3665 := __e.Get(4)
_ = V3665
V3666 := __e.Get(5)
_ = V3666
tmp13091 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3666)
}
__typedArg0 := Nil
__typedArg1 := V3666
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13091 {
__e.TailApply(PrimFunc(symshen_4side_1conditions_1_6goals), Nil, V3662, V3663, V3664, V3665)
return
} else {
tmp13089 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3665)
}
__typedArg0 := Nil
__typedArg1 := V3665
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13082 Obj

if True == tmp13089 {
tmp13088 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3666)
}
__typedArg0 := V3666
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13084 Obj

if True == tmp13088 {
tmp13086 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3666)
}
__typedArg0 := V3666
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13087 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp13086)
}
__typedArg0 := Nil
__typedArg1 := tmp13086
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13085 Obj

if True == tmp13087 {
ifres13085 = True


} else {
ifres13085 = False


}

ifres13084 = ifres13085


} else {
ifres13084 = False


}

var ifres13083 Obj

if True == ifres13084 {
ifres13083 = True


} else {
ifres13083 = False


}

ifres13082 = ifres13083


} else {
ifres13082 = False


}

if True == ifres13082 {
tmp13062 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3666)
}
__typedArg0 := V3666
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13063 := Call(__e, PrimFunc(symshen_4passive_1variables), tmp13062, V3662)


let__12273 := tmp13063
_ = let__12273

tmp13064 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3666)
}
__typedArg0 := V3666
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13065 := Call(__e, PrimFunc(symshen_4remove_1bystanders), V3662, tmp13064)


let__12274 := tmp13065
_ = let__12274

tmp13066 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3666)
}
__typedArg0 := V3666
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13067 := Call(__e, PrimFunc(symshen_4specialise_1member), tmp13066, V3663, let__12274, let__12273)


tmp13068 := Call(__e, PrimFunc(symshen_4side_1conditions_1_6goals), Nil, V3662, V3663, V3664, Nil)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13067, tmp13068)
}
__typedArg0 := tmp13067
__typedArg1 := tmp13068
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp13080 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3666)
}
__typedArg0 := V3666
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp13080 {
tmp13069 := Call(__e, PrimFunc(symgensym), symNewAssumptions)


let__12275 := tmp13069
_ = let__12275

tmp13070 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3666)
}
__typedArg0 := V3666
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13071 := Call(__e, PrimFunc(symshen_4passive_1variables), tmp13070, V3662)


let__12276 := tmp13071
_ = let__12276

tmp13072 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3666)
}
__typedArg0 := V3666
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13073 := Call(__e, PrimFunc(symshen_4remove_1bystanders), V3662, tmp13072)


let__12277 := tmp13073
_ = let__12277

tmp13074 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3666)
}
__typedArg0 := V3666
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13075 := Call(__e, PrimFunc(symshen_4specialise_1consume), tmp13074, V3663, let__12277, let__12276, let__12275)


tmp13076 := Call(__e, PrimFunc(symappend), V3662, let__12276)


tmp13077 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3666)
}
__typedArg0 := V3666
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13078 := Call(__e, PrimFunc(symshen_4rule_1_6body), tmp13076, let__12275, V3664, V3665, tmp13077)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13075, tmp13078)
}
__typedArg0 := tmp13075
__typedArg1 := tmp13078
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4rule_1_6body)
return
}


}


}


}, 5)

tmp13092 := Call(__e, ns2_1set, symshen_4rule_1_6body, tmp13061)


_ = tmp13092

tmp13093 := MakeNative(func(__e *ControlFlow) {
V3672 := __e.Get(1)
_ = V3672
V3673 := __e.Get(2)
_ = V3673
V3674 := __e.Get(3)
_ = V3674
V3675 := __e.Get(4)
_ = V3675
tmp13094 := Call(__e, PrimFunc(symgensym), symshen_4member)


let__12278 := tmp13094
_ = let__12278

tmp13095 := Call(__e, PrimFunc(symshen_4member_1clause), let__12278, V3672, V3674, V3675)


let__12279 := tmp13095
_ = let__12279

tmp13096 := Call(__e, PrimFunc(symappend), V3674, V3675)


tmp13097 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3673, tmp13096)
}
__typedArg0 := V3673
__typedArg1 := tmp13096
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12278, tmp13097)
}
__typedArg0 := let__12278
__typedArg1 := tmp13097
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 4)

tmp13098 := Call(__e, ns2_1set, symshen_4specialise_1member, tmp13093)


_ = tmp13098

tmp13099 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__selftop:
V3680 := __self1
_ = V3680
V3681 := __self2
_ = V3681
tmp13113 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3680)
}
__typedArg0 := Nil
__typedArg1 := V3680
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13113 {
__e.Return(Nil)
return
} else {
tmp13111 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3680)
}
__typedArg0 := V3680
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13107 Obj

if True == tmp13111 {
tmp13109 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3680)
}
__typedArg0 := V3680
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13110 := Call(__e, PrimFunc(symshen_4occurs_1check_2), tmp13109, V3681)


var ifres13108 Obj

if True == tmp13110 {
ifres13108 = True


} else {
ifres13108 = False


}

ifres13107 = ifres13108


} else {
ifres13107 = False


}

if True == ifres13107 {
tmp13100 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3680)
}
__typedArg0 := V3680
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13101 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3680)
}
__typedArg0 := V3680
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13102 := Call(__e, PrimFunc(symshen_4remove_1bystanders), tmp13101, V3681)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13100, tmp13102)
}
__typedArg0 := tmp13100
__typedArg1 := tmp13102
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp13105 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3680)
}
__typedArg0 := V3680
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp13105 {
tmp13103 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3680)
}
__typedArg0 := V3680
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4remove_1bystanders) == __self {
__self1, __self2 = tmp13103, V3681
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4remove_1bystanders), tmp13103, V3681)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4remove_1bystanders)
return
}


}


}


}, 2)

tmp13114 := Call(__e, ns2_1set, symshen_4remove_1bystanders, tmp13099)


_ = tmp13114

tmp13115 := MakeNative(func(__e *ControlFlow) {
V3682 := __e.Get(1)
_ = V3682
V3683 := __e.Get(2)
_ = V3683
V3684 := __e.Get(3)
_ = V3684
V3685 := __e.Get(4)
_ = V3685
tmp13116 := Call(__e, PrimFunc(symlength), V3685)


tmp13117 := Call(__e, PrimFunc(symshen_4nvars), tmp13116)


let__12280 := tmp13117
_ = let__12280

tmp13118 := Call(__e, PrimFunc(symshen_4macro_1_8ch), V3683)


tmp13119 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym__, Nil)
}
__typedArg0 := sym__
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13120 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13118, tmp13119)
}
__typedArg0 := tmp13118
__typedArg1 := tmp13119
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13121 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp13120)
}
__typedArg0 := symcons
__typedArg1 := tmp13120
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13122 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13121, Nil)
}
__typedArg0 := tmp13121
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13123 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1, tmp13122)
}
__typedArg0 := sym_1
__typedArg1 := tmp13122
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13124 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13123, Nil)
}
__typedArg0 := tmp13123
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13125 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1_1, Nil)
}
__typedArg0 := sym_5_1_1
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13126 := Call(__e, PrimFunc(symshen_4passive_1bind), V3685, let__12280)


tmp13127 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(";"))
}
__typedArg0 := MakeString(";")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp13128 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13127, Nil)
}
__typedArg0 := tmp13127
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13129 := Call(__e, PrimFunc(symappend), tmp13126, tmp13128)


tmp13130 := Call(__e, PrimFunc(symappend), tmp13125, tmp13129)


tmp13131 := Call(__e, PrimFunc(symappend), let__12280, tmp13130)


tmp13132 := Call(__e, PrimFunc(symappend), V3684, tmp13131)


tmp13133 := Call(__e, PrimFunc(symappend), tmp13124, tmp13132)


let__12281 := tmp13133
_ = let__12281

tmp13134 := Call(__e, PrimFunc(symgensym), symHypotheses)


let__12283 := tmp13134
_ = let__12283

tmp13135 := Call(__e, PrimFunc(symappend), V3684, V3685)


let__12284 := tmp13135
_ = let__12284

tmp13136 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12283, Nil)
}
__typedArg0 := let__12283
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13137 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym__, tmp13136)
}
__typedArg0 := sym__
__typedArg1 := tmp13136
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13138 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp13137)
}
__typedArg0 := symcons
__typedArg1 := tmp13137
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13139 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13138, Nil)
}
__typedArg0 := tmp13138
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13140 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1, tmp13139)
}
__typedArg0 := sym_1
__typedArg1 := tmp13139
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13141 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13140, Nil)
}
__typedArg0 := tmp13140
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13142 := Call(__e, PrimFunc(symappend), tmp13141, let__12284)


let__12285 := tmp13142
_ = let__12285

tmp13143 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12283, let__12284)
}
__typedArg0 := let__12283
__typedArg1 := let__12284
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13144 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3682, tmp13143)
}
__typedArg0 := V3682
__typedArg1 := tmp13143
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13145 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13144, Nil)
}
__typedArg0 := tmp13144
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__12286 := tmp13145
_ = let__12286

tmp13146 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1_1, Nil)
}
__typedArg0 := sym_5_1_1
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13147 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(";"))
}
__typedArg0 := MakeString(";")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp13148 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13147, Nil)
}
__typedArg0 := tmp13147
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13149 := Call(__e, PrimFunc(symappend), let__12286, tmp13148)


tmp13150 := Call(__e, PrimFunc(symappend), tmp13146, tmp13149)


tmp13151 := Call(__e, PrimFunc(symappend), let__12285, tmp13150)


let__12282 := tmp13151
_ = let__12282

tmp13152 := Call(__e, PrimFunc(symappend), let__12281, let__12282)


tmp13153 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3682, tmp13152)
}
__typedArg0 := V3682
__typedArg1 := tmp13152
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13154 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefprolog, tmp13153)
}
__typedArg0 := symdefprolog
__typedArg1 := tmp13153
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__12287 := tmp13154
_ = let__12287

__e.TailApply(PrimFunc(symeval), let__12287)
return


}, 4)

tmp13155 := Call(__e, ns2_1set, symshen_4member_1clause, tmp13115)


_ = tmp13155

tmp13156 := MakeNative(func(__e *ControlFlow) {
V3694 := __e.Get(1)
_ = V3694
tmp13161 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(0), V3694)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := V3694
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13161 {
__e.Return(Nil)
return
} else {
tmp13157 := Call(__e, PrimFunc(symgensym), symNewV)


tmp13159 := Call(__e, PrimFunc(symshen_4nvars), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(V3694)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := V3694
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})())


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13157, tmp13159)
}
__typedArg0 := tmp13157
__typedArg1 := tmp13159
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}


}, 1)

tmp13162 := Call(__e, ns2_1set, symshen_4nvars, tmp13156)


_ = tmp13162

tmp13163 := MakeNative(func(__e *ControlFlow) {
V3695 := __e.Get(1)
_ = V3695
V3696 := __e.Get(2)
_ = V3696
tmp13181 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3695)
}
__typedArg0 := Nil
__typedArg1 := V3695
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13178 Obj

if True == tmp13181 {
tmp13180 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3696)
}
__typedArg0 := Nil
__typedArg1 := V3696
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13179 Obj

if True == tmp13180 {
ifres13179 = True


} else {
ifres13179 = False


}

ifres13178 = ifres13179


} else {
ifres13178 = False


}

if True == ifres13178 {
__e.Return(Nil)
return
} else {
tmp13176 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3695)
}
__typedArg0 := V3695
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13173 Obj

if True == tmp13176 {
tmp13175 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3696)
}
__typedArg0 := V3696
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13174 Obj

if True == tmp13175 {
ifres13174 = True


} else {
ifres13174 = False


}

ifres13173 = ifres13174


} else {
ifres13173 = False


}

if True == ifres13173 {
tmp13164 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3696)
}
__typedArg0 := V3696
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13165 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3695)
}
__typedArg0 := V3695
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13166 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13165, Nil)
}
__typedArg0 := tmp13165
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13167 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13164, tmp13166)
}
__typedArg0 := tmp13164
__typedArg1 := tmp13166
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13168 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symbind, tmp13167)
}
__typedArg0 := symbind
__typedArg1 := tmp13167
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13169 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3695)
}
__typedArg0 := V3695
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13170 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3696)
}
__typedArg0 := V3696
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13171 := Call(__e, PrimFunc(symshen_4passive_1bind), tmp13169, tmp13170)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13168, tmp13171)
}
__typedArg0 := tmp13168
__typedArg1 := tmp13171
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4passive_1bind)
return
}


}


}, 2)

tmp13182 := Call(__e, ns2_1set, symshen_4passive_1bind, tmp13163)


_ = tmp13182

tmp13183 := MakeNative(func(__e *ControlFlow) {
V3697 := __e.Get(1)
_ = V3697
V3698 := __e.Get(2)
_ = V3698
V3699 := __e.Get(3)
_ = V3699
V3700 := __e.Get(4)
_ = V3700
V3701 := __e.Get(5)
_ = V3701
tmp13184 := Call(__e, PrimFunc(symgensym), symshen_4consume)


let__12288 := tmp13184
_ = let__12288

tmp13185 := Call(__e, PrimFunc(symshen_4consume_1clause), let__12288, V3697, V3699, V3700, V3701)


let__12289 := tmp13185
_ = let__12289

tmp13186 := Call(__e, PrimFunc(symappend), V3699, V3700)


tmp13187 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3701, tmp13186)
}
__typedArg0 := V3701
__typedArg1 := tmp13186
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13188 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3698, tmp13187)
}
__typedArg0 := V3698
__typedArg1 := tmp13187
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12288, tmp13188)
}
__typedArg0 := let__12288
__typedArg1 := tmp13188
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 5)

tmp13189 := Call(__e, ns2_1set, symshen_4specialise_1consume, tmp13183)


_ = tmp13189

tmp13190 := MakeNative(func(__e *ControlFlow) {
V3704 := __e.Get(1)
_ = V3704
V3705 := __e.Get(2)
_ = V3705
V3706 := __e.Get(3)
_ = V3706
V3707 := __e.Get(4)
_ = V3707
V3708 := __e.Get(5)
_ = V3708
tmp13191 := Call(__e, PrimFunc(symlength), V3707)


tmp13192 := Call(__e, PrimFunc(symshen_4nvars), tmp13191)


let__12290 := tmp13192
_ = let__12290

tmp13193 := Call(__e, PrimFunc(symgensym), symAssumption)


let__12291 := tmp13193
_ = let__12291

tmp13194 := Call(__e, PrimFunc(symshen_4macro_1_8ch), V3705)


tmp13195 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12291, Nil)
}
__typedArg0 := let__12291
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13196 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13194, tmp13195)
}
__typedArg0 := tmp13194
__typedArg1 := tmp13195
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13197 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp13196)
}
__typedArg0 := symcons
__typedArg1 := tmp13196
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13198 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13197, Nil)
}
__typedArg0 := tmp13197
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13199 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1, tmp13198)
}
__typedArg0 := sym_1
__typedArg1 := tmp13198
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13200 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1_1, Nil)
}
__typedArg0 := sym_5_1_1
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13201 := Call(__e, PrimFunc(symshen_4passive_1bind), V3707, let__12290)


tmp13202 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12291, Nil)
}
__typedArg0 := let__12291
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13203 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3708, tmp13202)
}
__typedArg0 := V3708
__typedArg1 := tmp13202
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13204 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symbind, tmp13203)
}
__typedArg0 := symbind
__typedArg1 := tmp13203
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13205 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(";"))
}
__typedArg0 := MakeString(";")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp13206 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13205, Nil)
}
__typedArg0 := tmp13205
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13207 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13204, tmp13206)
}
__typedArg0 := tmp13204
__typedArg1 := tmp13206
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13208 := Call(__e, PrimFunc(symappend), tmp13201, tmp13207)


tmp13209 := Call(__e, PrimFunc(symappend), tmp13200, tmp13208)


tmp13210 := Call(__e, PrimFunc(symappend), let__12290, tmp13209)


tmp13211 := Call(__e, PrimFunc(symappend), V3706, tmp13210)


tmp13212 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3708, tmp13211)
}
__typedArg0 := V3708
__typedArg1 := tmp13211
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13213 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13199, tmp13212)
}
__typedArg0 := tmp13199
__typedArg1 := tmp13212
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__12292 := tmp13213
_ = let__12292

tmp13214 := Call(__e, PrimFunc(symgensym), symHypotheses)


let__12294 := tmp13214
_ = let__12294

tmp13215 := Call(__e, PrimFunc(symappend), V3706, V3707)


let__12295 := tmp13215
_ = let__12295

tmp13216 := Call(__e, PrimFunc(symgensym), symAssumptions)


let__12296 := tmp13216
_ = let__12296

tmp13217 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12294, Nil)
}
__typedArg0 := let__12294
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13218 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12291, tmp13217)
}
__typedArg0 := let__12291
__typedArg1 := tmp13217
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13219 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp13218)
}
__typedArg0 := symcons
__typedArg1 := tmp13218
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13220 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13219, Nil)
}
__typedArg0 := tmp13219
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13221 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1, tmp13220)
}
__typedArg0 := sym_1
__typedArg1 := tmp13220
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13222 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3708, Nil)
}
__typedArg0 := V3708
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13223 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12296, tmp13222)
}
__typedArg0 := let__12296
__typedArg1 := tmp13222
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13224 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp13223)
}
__typedArg0 := symcons
__typedArg1 := tmp13223
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13225 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13224, let__12295)
}
__typedArg0 := tmp13224
__typedArg1 := let__12295
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13226 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13221, tmp13225)
}
__typedArg0 := tmp13221
__typedArg1 := tmp13225
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__12297 := tmp13226
_ = let__12297

tmp13227 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12291, Nil)
}
__typedArg0 := let__12291
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13228 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12296, tmp13227)
}
__typedArg0 := let__12296
__typedArg1 := tmp13227
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13229 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symbind, tmp13228)
}
__typedArg0 := symbind
__typedArg1 := tmp13228
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13230 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3708, let__12295)
}
__typedArg0 := V3708
__typedArg1 := let__12295
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13231 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12294, tmp13230)
}
__typedArg0 := let__12294
__typedArg1 := tmp13230
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13232 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3704, tmp13231)
}
__typedArg0 := V3704
__typedArg1 := tmp13231
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13233 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13232, Nil)
}
__typedArg0 := tmp13232
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13234 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13229, tmp13233)
}
__typedArg0 := tmp13229
__typedArg1 := tmp13233
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__12298 := tmp13234
_ = let__12298

tmp13235 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1_1, Nil)
}
__typedArg0 := sym_5_1_1
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13236 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(";"))
}
__typedArg0 := MakeString(";")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp13237 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13236, Nil)
}
__typedArg0 := tmp13236
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13238 := Call(__e, PrimFunc(symappend), let__12298, tmp13237)


tmp13239 := Call(__e, PrimFunc(symappend), tmp13235, tmp13238)


tmp13240 := Call(__e, PrimFunc(symappend), let__12297, tmp13239)


let__12293 := tmp13240
_ = let__12293

tmp13241 := Call(__e, PrimFunc(symappend), let__12292, let__12293)


tmp13242 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3704, tmp13241)
}
__typedArg0 := V3704
__typedArg1 := tmp13241
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13243 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefprolog, tmp13242)
}
__typedArg0 := symdefprolog
__typedArg1 := tmp13242
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__12299 := tmp13243
_ = let__12299

__e.TailApply(PrimFunc(symeval), let__12299)
return


}, 5)

tmp13244 := Call(__e, ns2_1set, symshen_4consume_1clause, tmp13190)


_ = tmp13244

tmp13245 := MakeNative(func(__e *ControlFlow) {
V3719 := __e.Get(1)
_ = V3719
V3720 := __e.Get(2)
_ = V3720
tmp13246 := Call(__e, PrimFunc(symshen_4extract_1vars), V3719)


__e.TailApply(PrimFunc(symdifference), tmp13246, V3720)
return


}, 2)

tmp13247 := Call(__e, ns2_1set, symshen_4passive_1variables, tmp13245)


_ = tmp13247

tmp13248 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__self3 := __e.Get(3)
__self4 := __e.Get(4)
__self5 := __e.Get(5)
__selftop:
V3725 := __self1
_ = V3725
V3726 := __self2
_ = V3726
V3727 := __self3
_ = V3727
V3728 := __self4
_ = V3728
V3729 := __self5
_ = V3729
tmp13399 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3728)
}
__typedArg0 := Nil
__typedArg1 := V3728
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13399 {
__e.TailApply(PrimFunc(symshen_4premises_1_6goals), V3725, V3727, V3729)
return
} else {
tmp13397 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3728)
}
__typedArg0 := V3728
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13377 Obj

if True == tmp13397 {
tmp13395 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3728)
}
__typedArg0 := V3728
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13396 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13395)
}
__typedArg0 := tmp13395
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13379 Obj

if True == tmp13396 {
tmp13392 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3728)
}
__typedArg0 := V3728
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13393 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13392)
}
__typedArg0 := tmp13392
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13394 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symif, tmp13393)
}
__typedArg0 := symif
__typedArg1 := tmp13393
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13381 Obj

if True == tmp13394 {
tmp13389 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3728)
}
__typedArg0 := V3728
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13390 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13389)
}
__typedArg0 := tmp13389
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13391 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13390)
}
__typedArg0 := tmp13390
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13383 Obj

if True == tmp13391 {
tmp13385 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3728)
}
__typedArg0 := V3728
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13386 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13385)
}
__typedArg0 := tmp13385
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13387 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13386)
}
__typedArg0 := tmp13386
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13388 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp13387)
}
__typedArg0 := Nil
__typedArg1 := tmp13387
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13384 Obj

if True == tmp13388 {
ifres13384 = True


} else {
ifres13384 = False


}

ifres13383 = ifres13384


} else {
ifres13383 = False


}

var ifres13382 Obj

if True == ifres13383 {
ifres13382 = True


} else {
ifres13382 = False


}

ifres13381 = ifres13382


} else {
ifres13381 = False


}

var ifres13380 Obj

if True == ifres13381 {
ifres13380 = True


} else {
ifres13380 = False


}

ifres13379 = ifres13380


} else {
ifres13379 = False


}

var ifres13378 Obj

if True == ifres13379 {
ifres13378 = True


} else {
ifres13378 = False


}

ifres13377 = ifres13378


} else {
ifres13377 = False


}

if True == ifres13377 {
tmp13249 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3728)
}
__typedArg0 := V3728
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13250 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13249)
}
__typedArg0 := tmp13249
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13251 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhen, tmp13250)
}
__typedArg0 := symwhen
__typedArg1 := tmp13250
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13252 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3728)
}
__typedArg0 := V3728
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13253 := Call(__e, PrimFunc(symshen_4side_1conditions_1_6goals), V3725, V3726, V3727, tmp13252, V3729)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13251, tmp13253)
}
__typedArg0 := tmp13251
__typedArg1 := tmp13253
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp13375 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3728)
}
__typedArg0 := V3728
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13348 Obj

if True == tmp13375 {
tmp13373 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3728)
}
__typedArg0 := V3728
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13374 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13373)
}
__typedArg0 := tmp13373
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13350 Obj

if True == tmp13374 {
tmp13370 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3728)
}
__typedArg0 := V3728
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13371 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13370)
}
__typedArg0 := tmp13370
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13372 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlet, tmp13371)
}
__typedArg0 := symlet
__typedArg1 := tmp13371
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13352 Obj

if True == tmp13372 {
tmp13367 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3728)
}
__typedArg0 := V3728
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13368 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13367)
}
__typedArg0 := tmp13367
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13369 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13368)
}
__typedArg0 := tmp13368
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13354 Obj

if True == tmp13369 {
tmp13363 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3728)
}
__typedArg0 := V3728
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13364 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13363)
}
__typedArg0 := tmp13363
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13365 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13364)
}
__typedArg0 := tmp13364
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13366 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13365)
}
__typedArg0 := tmp13365
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13356 Obj

if True == tmp13366 {
tmp13358 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3728)
}
__typedArg0 := V3728
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13359 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13358)
}
__typedArg0 := tmp13358
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13360 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13359)
}
__typedArg0 := tmp13359
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13361 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13360)
}
__typedArg0 := tmp13360
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13362 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp13361)
}
__typedArg0 := Nil
__typedArg1 := tmp13361
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13357 Obj

if True == tmp13362 {
ifres13357 = True


} else {
ifres13357 = False


}

ifres13356 = ifres13357


} else {
ifres13356 = False


}

var ifres13355 Obj

if True == ifres13356 {
ifres13355 = True


} else {
ifres13355 = False


}

ifres13354 = ifres13355


} else {
ifres13354 = False


}

var ifres13353 Obj

if True == ifres13354 {
ifres13353 = True


} else {
ifres13353 = False


}

ifres13352 = ifres13353


} else {
ifres13352 = False


}

var ifres13351 Obj

if True == ifres13352 {
ifres13351 = True


} else {
ifres13351 = False


}

ifres13350 = ifres13351


} else {
ifres13350 = False


}

var ifres13349 Obj

if True == ifres13350 {
ifres13349 = True


} else {
ifres13349 = False


}

ifres13348 = ifres13349


} else {
ifres13348 = False


}

if True == ifres13348 {
tmp13269 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3728)
}
__typedArg0 := V3728
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13270 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13269)
}
__typedArg0 := tmp13269
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13271 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13270)
}
__typedArg0 := tmp13270
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13272 := Call(__e, PrimFunc(symelement_2), tmp13271, V3726)


if True == tmp13272 {
tmp13254 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3728)
}
__typedArg0 := V3728
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13255 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13254)
}
__typedArg0 := tmp13254
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13256 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symis_b, tmp13255)
}
__typedArg0 := symis_b
__typedArg1 := tmp13255
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13257 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3728)
}
__typedArg0 := V3728
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13258 := Call(__e, PrimFunc(symshen_4side_1conditions_1_6goals), V3725, V3726, V3727, tmp13257, V3729)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13256, tmp13258)
}
__typedArg0 := tmp13256
__typedArg1 := tmp13258
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp13259 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3728)
}
__typedArg0 := V3728
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13260 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13259)
}
__typedArg0 := tmp13259
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13261 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symbind, tmp13260)
}
__typedArg0 := symbind
__typedArg1 := tmp13260
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13262 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3728)
}
__typedArg0 := V3728
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13263 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13262)
}
__typedArg0 := tmp13262
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13264 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13263)
}
__typedArg0 := tmp13263
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13265 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13264, V3726)
}
__typedArg0 := tmp13264
__typedArg1 := V3726
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13266 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3728)
}
__typedArg0 := V3728
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13267 := Call(__e, PrimFunc(symshen_4side_1conditions_1_6goals), V3725, tmp13265, V3727, tmp13266, V3729)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13261, tmp13267)
}
__typedArg0 := tmp13261
__typedArg1 := tmp13267
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}


} else {
tmp13346 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3728)
}
__typedArg0 := V3728
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13326 Obj

if True == tmp13346 {
tmp13344 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3728)
}
__typedArg0 := V3728
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13345 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13344)
}
__typedArg0 := tmp13344
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13328 Obj

if True == tmp13345 {
tmp13341 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3728)
}
__typedArg0 := V3728
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13342 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13341)
}
__typedArg0 := tmp13341
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13343 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symctxt, tmp13342)
}
__typedArg0 := symctxt
__typedArg1 := tmp13342
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13330 Obj

if True == tmp13343 {
tmp13338 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3728)
}
__typedArg0 := V3728
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13339 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13338)
}
__typedArg0 := tmp13338
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13340 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13339)
}
__typedArg0 := tmp13339
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13332 Obj

if True == tmp13340 {
tmp13334 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3728)
}
__typedArg0 := V3728
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13335 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13334)
}
__typedArg0 := tmp13334
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13336 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13335)
}
__typedArg0 := tmp13335
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13337 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp13336)
}
__typedArg0 := Nil
__typedArg1 := tmp13336
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13333 Obj

if True == tmp13337 {
ifres13333 = True


} else {
ifres13333 = False


}

ifres13332 = ifres13333


} else {
ifres13332 = False


}

var ifres13331 Obj

if True == ifres13332 {
ifres13331 = True


} else {
ifres13331 = False


}

ifres13330 = ifres13331


} else {
ifres13330 = False


}

var ifres13329 Obj

if True == ifres13330 {
ifres13329 = True


} else {
ifres13329 = False


}

ifres13328 = ifres13329


} else {
ifres13328 = False


}

var ifres13327 Obj

if True == ifres13328 {
ifres13327 = True


} else {
ifres13327 = False


}

ifres13326 = ifres13327


} else {
ifres13326 = False


}

if True == ifres13326 {
tmp13298 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3728)
}
__typedArg0 := V3728
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13299 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13298)
}
__typedArg0 := tmp13298
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13300 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13299)
}
__typedArg0 := tmp13299
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13301 := Call(__e, PrimFunc(symelement_2), tmp13300, V3726)


if True == tmp13301 {
tmp13273 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3728)
}
__typedArg0 := V3728
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13274 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13273)
}
__typedArg0 := tmp13273
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13275 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13274)
}
__typedArg0 := tmp13274
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13276 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13275, V3725)
}
__typedArg0 := tmp13275
__typedArg1 := V3725
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13277 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3728)
}
__typedArg0 := V3728
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4side_1conditions_1_6goals) == __self {
__self1, __self2, __self3, __self4, __self5 = tmp13276, V3726, V3727, tmp13277, V3729
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4side_1conditions_1_6goals), tmp13276, V3726, V3727, tmp13277, V3729)
return


} else {
tmp13278 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3728)
}
__typedArg0 := V3728
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13279 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13278)
}
__typedArg0 := tmp13278
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13280 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13279)
}
__typedArg0 := tmp13279
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13281 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3727, Nil)
}
__typedArg0 := V3727
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13282 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13280, tmp13281)
}
__typedArg0 := tmp13280
__typedArg1 := tmp13281
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13283 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symbind, tmp13282)
}
__typedArg0 := symbind
__typedArg1 := tmp13282
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13284 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3728)
}
__typedArg0 := V3728
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13285 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13284)
}
__typedArg0 := tmp13284
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13286 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13285)
}
__typedArg0 := tmp13285
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13287 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13286, V3725)
}
__typedArg0 := tmp13286
__typedArg1 := V3725
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13288 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3728)
}
__typedArg0 := V3728
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13289 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13288)
}
__typedArg0 := tmp13288
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13290 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13289)
}
__typedArg0 := tmp13289
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13291 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13290, V3726)
}
__typedArg0 := tmp13290
__typedArg1 := V3726
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13292 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3728)
}
__typedArg0 := V3728
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13293 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13292)
}
__typedArg0 := tmp13292
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13294 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13293)
}
__typedArg0 := tmp13293
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13295 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3728)
}
__typedArg0 := V3728
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13296 := Call(__e, PrimFunc(symshen_4side_1conditions_1_6goals), tmp13287, tmp13291, tmp13294, tmp13295, V3729)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13283, tmp13296)
}
__typedArg0 := tmp13283
__typedArg1 := tmp13296
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}


} else {
tmp13324 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3728)
}
__typedArg0 := V3728
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13304 Obj

if True == tmp13324 {
tmp13322 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3728)
}
__typedArg0 := V3728
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13323 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13322)
}
__typedArg0 := tmp13322
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13306 Obj

if True == tmp13323 {
tmp13319 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3728)
}
__typedArg0 := V3728
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13320 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13319)
}
__typedArg0 := tmp13319
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13321 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symsqts, tmp13320)
}
__typedArg0 := symsqts
__typedArg1 := tmp13320
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13308 Obj

if True == tmp13321 {
tmp13316 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3728)
}
__typedArg0 := V3728
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13317 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13316)
}
__typedArg0 := tmp13316
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13318 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13317)
}
__typedArg0 := tmp13317
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13310 Obj

if True == tmp13318 {
tmp13312 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3728)
}
__typedArg0 := V3728
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13313 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13312)
}
__typedArg0 := tmp13312
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13314 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13313)
}
__typedArg0 := tmp13313
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13315 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp13314)
}
__typedArg0 := Nil
__typedArg1 := tmp13314
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13311 Obj

if True == tmp13315 {
ifres13311 = True


} else {
ifres13311 = False


}

ifres13310 = ifres13311


} else {
ifres13310 = False


}

var ifres13309 Obj

if True == ifres13310 {
ifres13309 = True


} else {
ifres13309 = False


}

ifres13308 = ifres13309


} else {
ifres13308 = False


}

var ifres13307 Obj

if True == ifres13308 {
ifres13307 = True


} else {
ifres13307 = False


}

ifres13306 = ifres13307


} else {
ifres13306 = False


}

var ifres13305 Obj

if True == ifres13306 {
ifres13305 = True


} else {
ifres13305 = False


}

ifres13304 = ifres13305


} else {
ifres13304 = False


}

if True == ifres13304 {
tmp13302 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3728)
}
__typedArg0 := V3728
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4side_1conditions_1_6goals) == __self {
__self1, __self2, __self3, __self4, __self5 = V3725, V3726, V3727, tmp13302, V3729
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4side_1conditions_1_6goals), V3725, V3726, V3727, tmp13302, V3729)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4side_1conditions_1_6goals)
return
}


}


}


}


}


}, 5)

tmp13400 := Call(__e, ns2_1set, symshen_4side_1conditions_1_6goals, tmp13248)


_ = tmp13400

tmp13401 := MakeNative(func(__e *ControlFlow) {
V3734 := __e.Get(1)
_ = V3734
V3735 := __e.Get(2)
_ = V3735
V3736 := __e.Get(3)
_ = V3736
tmp13451 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3736)
}
__typedArg0 := Nil
__typedArg1 := V3736
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13451 {
tmp13402 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(";"))
}
__typedArg0 := MakeString(";")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13402, Nil)
}
__typedArg0 := tmp13402
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp13449 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3736)
}
__typedArg0 := V3736
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13445 Obj

if True == tmp13449 {
tmp13447 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3736)
}
__typedArg0 := V3736
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13448 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_b, tmp13447)
}
__typedArg0 := sym_b
__typedArg1 := tmp13447
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13446 Obj

if True == tmp13448 {
ifres13446 = True


} else {
ifres13446 = False


}

ifres13445 = ifres13446


} else {
ifres13445 = False


}

if True == ifres13445 {
tmp13403 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3736)
}
__typedArg0 := V3736
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13404 := Call(__e, PrimFunc(symshen_4premises_1_6goals), V3734, V3735, tmp13403)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_b, tmp13404)
}
__typedArg0 := sym_b
__typedArg1 := tmp13404
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp13443 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3736)
}
__typedArg0 := V3736
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13439 Obj

if True == tmp13443 {
tmp13441 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3736)
}
__typedArg0 := V3736
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13442 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symfail, tmp13441)
}
__typedArg0 := symfail
__typedArg1 := tmp13441
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13440 Obj

if True == tmp13442 {
ifres13440 = True


} else {
ifres13440 = False


}

ifres13439 = ifres13440


} else {
ifres13439 = False


}

if True == ifres13439 {
tmp13405 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(False, Nil)
}
__typedArg0 := False
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13406 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhen, tmp13405)
}
__typedArg0 := symwhen
__typedArg1 := tmp13405
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13407 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3736)
}
__typedArg0 := V3736
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13408 := Call(__e, PrimFunc(symshen_4premises_1_6goals), V3734, V3735, tmp13407)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13406, tmp13408)
}
__typedArg0 := tmp13406
__typedArg1 := tmp13408
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp13437 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3736)
}
__typedArg0 := V3736
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13422 Obj

if True == tmp13437 {
tmp13435 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3736)
}
__typedArg0 := V3736
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13436 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13435)
}
__typedArg0 := tmp13435
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13424 Obj

if True == tmp13436 {
tmp13432 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3736)
}
__typedArg0 := V3736
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13433 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13432)
}
__typedArg0 := tmp13432
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13434 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13433)
}
__typedArg0 := tmp13433
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13426 Obj

if True == tmp13434 {
tmp13428 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3736)
}
__typedArg0 := V3736
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13429 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13428)
}
__typedArg0 := tmp13428
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13430 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13429)
}
__typedArg0 := tmp13429
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13431 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp13430)
}
__typedArg0 := Nil
__typedArg1 := tmp13430
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13427 Obj

if True == tmp13431 {
ifres13427 = True


} else {
ifres13427 = False


}

ifres13426 = ifres13427


} else {
ifres13426 = False


}

var ifres13425 Obj

if True == ifres13426 {
ifres13425 = True


} else {
ifres13425 = False


}

ifres13424 = ifres13425


} else {
ifres13424 = False


}

var ifres13423 Obj

if True == ifres13424 {
ifres13423 = True


} else {
ifres13423 = False


}

ifres13422 = ifres13423


} else {
ifres13422 = False


}

if True == ifres13422 {
tmp13409 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3736)
}
__typedArg0 := V3736
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13410 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13409)
}
__typedArg0 := tmp13409
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13411 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13410)
}
__typedArg0 := tmp13410
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13412 := Call(__e, PrimFunc(symshen_4macro_1_8c), tmp13411)


tmp13413 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3736)
}
__typedArg0 := V3736
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13414 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13413)
}
__typedArg0 := tmp13413
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13415 := Call(__e, PrimFunc(symshen_4construct_1context), V3734, tmp13414, V3735)


tmp13416 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13415, Nil)
}
__typedArg0 := tmp13415
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13417 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13412, tmp13416)
}
__typedArg0 := tmp13412
__typedArg1 := tmp13416
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13418 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4system_1S, tmp13417)
}
__typedArg0 := symshen_4system_1S
__typedArg1 := tmp13417
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13419 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3736)
}
__typedArg0 := V3736
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13420 := Call(__e, PrimFunc(symshen_4premises_1_6goals), V3734, V3735, tmp13419)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13418, tmp13420)
}
__typedArg0 := tmp13418
__typedArg1 := tmp13420
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4premises_1_6goals)
return
}


}


}


}


}, 3)

tmp13452 := Call(__e, ns2_1set, symshen_4premises_1_6goals, tmp13401)


_ = tmp13452

tmp13453 := MakeNative(func(__e *ControlFlow) {
V3740 := __e.Get(1)
_ = V3740
V3741 := __e.Get(2)
_ = V3741
V3742 := __e.Get(3)
_ = V3742
tmp13473 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3741)
}
__typedArg0 := Nil
__typedArg1 := V3741
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13473 {
__e.Return(V3742)
return
} else {
tmp13471 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3741)
}
__typedArg0 := V3741
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13463 Obj

if True == tmp13471 {
tmp13469 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3741)
}
__typedArg0 := V3741
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13470 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp13469)
}
__typedArg0 := Nil
__typedArg1 := tmp13469
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13465 Obj

if True == tmp13470 {
tmp13467 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3741)
}
__typedArg0 := V3741
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13468 := Call(__e, PrimFunc(symelement_2), tmp13467, V3740)


var ifres13466 Obj

if True == tmp13468 {
ifres13466 = True


} else {
ifres13466 = False


}

ifres13465 = ifres13466


} else {
ifres13465 = False


}

var ifres13464 Obj

if True == ifres13465 {
ifres13464 = True


} else {
ifres13464 = False


}

ifres13463 = ifres13464


} else {
ifres13463 = False


}

if True == ifres13463 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3741)
}
__typedArg0 := V3741
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return
} else {
tmp13461 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3741)
}
__typedArg0 := V3741
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp13461 {
tmp13454 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3741)
}
__typedArg0 := V3741
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13455 := Call(__e, PrimFunc(symshen_4macro_1_8c), tmp13454)


tmp13456 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3741)
}
__typedArg0 := V3741
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13457 := Call(__e, PrimFunc(symshen_4construct_1context), V3740, tmp13456, V3742)


tmp13458 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13457, Nil)
}
__typedArg0 := tmp13457
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13459 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13455, tmp13458)
}
__typedArg0 := tmp13455
__typedArg1 := tmp13458
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp13459)
}
__typedArg0 := symcons
__typedArg1 := tmp13459
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4construct_1context)
return
}


}


}


}, 3)

tmp13474 := Call(__e, ns2_1set, symshen_4construct_1context, tmp13453)


_ = tmp13474

tmp13475 := MakeNative(func(__e *ControlFlow) {
V3743 := __e.Get(1)
_ = V3743
tmp13476 := MakeNative(func(__e *ControlFlow) {
Z3745 := __e.Get(1)
_ = Z3745
__e.TailApply(PrimFunc(symshen_4intern_1type), Z3745)
return
}, 1)

tmp13477 := Call(__e, PrimFunc(symmap), tmp13476, V3743)


let__12300 := tmp13477
_ = let__12300

tmp13478 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_ddatatypes_d)
}
__typedArg0 := symshen_4_ddatatypes_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

let__12301 := tmp13478
_ = let__12301

tmp13479 := Call(__e, PrimFunc(symshen_4remove_1datatypes), let__12300, let__12301)


let__12302 := tmp13479
_ = let__12302

tmp13480 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_ddatatypes_d, let__12302)
}
__typedArg0 := symshen_4_ddatatypes_d
__typedArg1 := let__12302
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

let__12303 := tmp13480
_ = let__12303

__e.TailApply(PrimFunc(symshen_4show_1datatypes), let__12303)
return


}, 1)

tmp13481 := Call(__e, ns2_1set, sympreclude, tmp13475)


_ = tmp13481

tmp13482 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__selftop:
V3753 := __self1
_ = V3753
V3754 := __self2
_ = V3754
tmp13489 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3753)
}
__typedArg0 := Nil
__typedArg1 := V3753
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13489 {
__e.Return(V3754)
return
} else {
tmp13487 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3753)
}
__typedArg0 := V3753
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp13487 {
tmp13483 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3753)
}
__typedArg0 := V3753
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13484 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3753)
}
__typedArg0 := V3753
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13485 := Call(__e, PrimFunc(symshen_4unassoc), tmp13484, V3754)


if PrimFunc(symshen_4remove_1datatypes) == __self {
__self1, __self2 = tmp13483, tmp13485
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4remove_1datatypes), tmp13483, tmp13485)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.remove-datatypes"))
}
__typedArg0 := MakeString("implementation error in shen.remove-datatypes")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 2)

tmp13490 := Call(__e, ns2_1set, symshen_4remove_1datatypes, tmp13482)


_ = tmp13490

tmp13491 := MakeNative(func(__e *ControlFlow) {
V3755 := __e.Get(1)
_ = V3755
tmp13492 := MakeNative(func(__e *ControlFlow) {
Z3756 := __e.Get(1)
_ = Z3756
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(Z3756)
}
__typedArg0 := Z3756
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp13492, V3755)
return


}, 1)

tmp13493 := Call(__e, ns2_1set, symshen_4show_1datatypes, tmp13491)


_ = tmp13493

tmp13494 := MakeNative(func(__e *ControlFlow) {
V3757 := __e.Get(1)
_ = V3757
tmp13495 := MakeNative(func(__e *ControlFlow) {
Z3759 := __e.Get(1)
_ = Z3759
__e.TailApply(PrimFunc(symshen_4intern_1type), Z3759)
return
}, 1)

tmp13496 := Call(__e, PrimFunc(symmap), tmp13495, V3757)


let__12304 := tmp13496
_ = let__12304

tmp13497 := MakeNative(func(__e *ControlFlow) {
Z3761 := __e.Get(1)
_ = Z3761
tmp13498 := Call(__e, PrimFunc(symfn), Z3761)


__e.TailApply(PrimFunc(symshen_4remember_1datatype), Z3761, tmp13498)
return


}, 1)

tmp13499 := Call(__e, PrimFunc(symmap), tmp13497, let__12304)


let__12305 := tmp13499
_ = let__12305

tmp13500 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_ddatatypes_d)
}
__typedArg0 := symshen_4_ddatatypes_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

let__12306 := tmp13500
_ = let__12306

__e.TailApply(PrimFunc(symshen_4show_1datatypes), let__12306)
return


}, 1)

tmp13501 := Call(__e, ns2_1set, syminclude, tmp13494)


_ = tmp13501

tmp13502 := MakeNative(func(__e *ControlFlow) {
V3763 := __e.Get(1)
_ = V3763
tmp13503 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_ddatatypes_d, Nil)
}
__typedArg0 := symshen_4_ddatatypes_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

let__12307 := tmp13503
_ = let__12307

tmp13504 := MakeNative(func(__e *ControlFlow) {
Z3766 := __e.Get(1)
_ = Z3766
__e.TailApply(PrimFunc(symshen_4intern_1type), Z3766)
return
}, 1)

tmp13505 := Call(__e, PrimFunc(symmap), tmp13504, V3763)


let__12308 := tmp13505
_ = let__12308

tmp13506 := MakeNative(func(__e *ControlFlow) {
Z3768 := __e.Get(1)
_ = Z3768
tmp13507 := Call(__e, PrimFunc(symfn), Z3768)


__e.TailApply(PrimFunc(symshen_4remember_1datatype), Z3768, tmp13507)
return


}, 1)

tmp13508 := Call(__e, PrimFunc(symmap), tmp13506, let__12308)


let__12309 := tmp13508
_ = let__12309

tmp13509 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_ddatatypes_d)
}
__typedArg0 := symshen_4_ddatatypes_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4show_1datatypes), tmp13509)
return


}, 1)

tmp13510 := Call(__e, ns2_1set, sympreclude_1all_1but, tmp13502)


_ = tmp13510

tmp13511 := MakeNative(func(__e *ControlFlow) {
V3769 := __e.Get(1)
_ = V3769
tmp13512 := MakeNative(func(__e *ControlFlow) {
Z3771 := __e.Get(1)
_ = Z3771
__e.TailApply(PrimFunc(symshen_4intern_1type), Z3771)
return
}, 1)

tmp13513 := Call(__e, PrimFunc(symmap), tmp13512, V3769)


let__12310 := tmp13513
_ = let__12310

tmp13514 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dalldatatypes_d)
}
__typedArg0 := symshen_4_dalldatatypes_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

let__12311 := tmp13514
_ = let__12311

tmp13515 := Call(__e, PrimFunc(symshen_4remove_1datatypes), let__12310, let__12311)


tmp13516 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_ddatatypes_d, tmp13515)
}
__typedArg0 := symshen_4_ddatatypes_d
__typedArg1 := tmp13515
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

let__12312 := tmp13516
_ = let__12312

__e.TailApply(PrimFunc(symshen_4show_1datatypes), let__12312)
return


}, 1)

__e.TailApply(ns2_1set, syminclude_1all_1but, tmp13511)
return




}, 0)

