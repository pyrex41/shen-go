package main

import . "github.com/pyrex41/shen-go/kl"

var SequentMain = MakeNative(func(__e *ControlFlow) {
tmp12320 := MakeNative(func(__e *ControlFlow) {
V3418 := __e.Get(1)
_ = V3418
tmp12335 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3418)
}
__typedArg0 := V3418
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12321 Obj

if True == tmp12335 {
tmp12322 := Call(__e, PrimFunc(symhead), V3418)


let__12090 := tmp12322
_ = let__12090

tmp12323 := Call(__e, PrimFunc(symtail), V3418)


let__12091 := tmp12323
_ = let__12091

tmp12324 := Call(__e, PrimFunc(symshen_4_5datatype_1rules_6), let__12091)


let__12092 := tmp12324
_ = let__12092

tmp12333 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12092)


var ifres12325 Obj

if True == tmp12333 {
tmp12326 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12325 = tmp12326


} else {
tmp12327 := Call(__e, PrimFunc(symshen_4_5_1out), let__12092)


let__12093 := tmp12327
_ = let__12093

tmp12328 := Call(__e, PrimFunc(symshen_4in_1_6), let__12092)


let__12094 := tmp12328
_ = let__12094

tmp12329 := Call(__e, PrimFunc(symshen_4rules_1_6prolog), let__12090, let__12093)


let__12095 := tmp12329
_ = let__12095

tmp12330 := Call(__e, PrimFunc(symfn), let__12090)


tmp12331 := Call(__e, PrimFunc(symshen_4remember_1datatype), let__12090, tmp12330)


tmp12332 := Call(__e, PrimFunc(symshen_4comb), let__12094, tmp12331)


ifres12325 = tmp12332


}

ifres12321 = ifres12325


} else {
tmp12334 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12321 = tmp12334


}

let__12089 := ifres12321
_ = let__12089

tmp12337 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12089)


if True == tmp12337 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__12089)
return
}


}, 1)

tmp12338 := Call(__e, ns2_1set, symshen_4_5datatype_6, tmp12320)


_ = tmp12338

tmp12339 := MakeNative(func(__e *ControlFlow) {
V3426 := __e.Get(1)
_ = V3426
V3427 := __e.Get(2)
_ = V3427
tmp12340 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_ddatatypes_d)
}
__typedArg0 := symshen_4_ddatatypes_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp12341 := Call(__e, PrimFunc(symshen_4assoc_1_6), V3426, V3427, tmp12340)


tmp12342 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_ddatatypes_d, tmp12341)
}
__typedArg0 := symshen_4_ddatatypes_d
__typedArg1 := tmp12341
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp12342

tmp12343 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dalldatatypes_d)
}
__typedArg0 := symshen_4_dalldatatypes_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp12344 := Call(__e, PrimFunc(symshen_4assoc_1_6), V3426, V3427, tmp12343)


tmp12345 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dalldatatypes_d, tmp12344)
}
__typedArg0 := symshen_4_dalldatatypes_d
__typedArg1 := tmp12344
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp12345

__e.Return(V3426)
return


}, 2)

tmp12346 := Call(__e, ns2_1set, symshen_4remember_1datatype, tmp12339)


_ = tmp12346

tmp12347 := MakeNative(func(__e *ControlFlow) {
V3428 := __e.Get(1)
_ = V3428
tmp12348 := Call(__e, PrimFunc(symshen_4_5datatype_1rule_6), V3428)


let__12097 := tmp12348
_ = let__12097

tmp12361 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12097)


var ifres12349 Obj

if True == tmp12361 {
tmp12350 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12349 = tmp12350


} else {
tmp12351 := Call(__e, PrimFunc(symshen_4_5_1out), let__12097)


let__12098 := tmp12351
_ = let__12098

tmp12352 := Call(__e, PrimFunc(symshen_4in_1_6), let__12097)


let__12099 := tmp12352
_ = let__12099

tmp12353 := Call(__e, PrimFunc(symshen_4_5datatype_1rules_6), let__12099)


let__12100 := tmp12353
_ = let__12100

tmp12360 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12100)


var ifres12354 Obj

if True == tmp12360 {
tmp12355 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12354 = tmp12355


} else {
tmp12356 := Call(__e, PrimFunc(symshen_4_5_1out), let__12100)


let__12101 := tmp12356
_ = let__12101

tmp12357 := Call(__e, PrimFunc(symshen_4in_1_6), let__12100)


let__12102 := tmp12357
_ = let__12102

tmp12358 := Call(__e, PrimFunc(symappend), let__12098, let__12101)


tmp12359 := Call(__e, PrimFunc(symshen_4comb), let__12102, tmp12358)


ifres12354 = tmp12359


}

ifres12349 = ifres12354


}

let__12096 := ifres12349
_ = let__12096

tmp12377 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12096)


if True == tmp12377 {
tmp12362 := Call(__e, PrimFunc(sym_5_b_6), V3428)


let__12104 := tmp12362
_ = let__12104

tmp12373 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12104)


var ifres12363 Obj

if True == tmp12373 {
tmp12364 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12363 = tmp12364


} else {
tmp12365 := Call(__e, PrimFunc(symshen_4_5_1out), let__12104)


let__12105 := tmp12365
_ = let__12105

tmp12366 := Call(__e, PrimFunc(symshen_4in_1_6), let__12104)


let__12106 := tmp12366
_ = let__12106

tmp12371 := Call(__e, PrimFunc(symempty_2), let__12105)


var ifres12367 Obj

if True == tmp12371 {
ifres12367 = Nil


} else {
tmp12368 := Call(__e, PrimFunc(symshen_4app), let__12105, MakeString("\n ..."), symshen_4r)


tmp12370 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("datatype syntax error here:\n "))
__typedS1, __typedOK1 := TypedString(tmp12368)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("datatype syntax error here:\n ")
__typedArg1 := tmp12368
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("datatype syntax error here:\n "))
__typedS1, __typedOK1 := TypedString(tmp12368)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("datatype syntax error here:\n ")
__typedArg1 := tmp12368
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})()

ifres12367 = tmp12370


}

tmp12372 := Call(__e, PrimFunc(symshen_4comb), let__12106, ifres12367)


ifres12363 = tmp12372


}

let__12103 := ifres12363
_ = let__12103

tmp12375 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12103)


if True == tmp12375 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__12103)
return
}


} else {
__e.Return(let__12096)
return
}


}, 1)

tmp12378 := Call(__e, ns2_1set, symshen_4_5datatype_1rules_6, tmp12347)


_ = tmp12378

tmp12379 := MakeNative(func(__e *ControlFlow) {
V3440 := __e.Get(1)
_ = V3440
tmp12380 := Call(__e, PrimFunc(symshen_4_5single_6), V3440)


let__12108 := tmp12380
_ = let__12108

tmp12386 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12108)


var ifres12381 Obj

if True == tmp12386 {
tmp12382 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12381 = tmp12382


} else {
tmp12383 := Call(__e, PrimFunc(symshen_4_5_1out), let__12108)


let__12109 := tmp12383
_ = let__12109

tmp12384 := Call(__e, PrimFunc(symshen_4in_1_6), let__12108)


let__12110 := tmp12384
_ = let__12110

tmp12385 := Call(__e, PrimFunc(symshen_4comb), let__12110, let__12109)


ifres12381 = tmp12385


}

let__12107 := ifres12381
_ = let__12107

tmp12397 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12107)


if True == tmp12397 {
tmp12387 := Call(__e, PrimFunc(symshen_4_5double_6), V3440)


let__12112 := tmp12387
_ = let__12112

tmp12393 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12112)


var ifres12388 Obj

if True == tmp12393 {
tmp12389 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12388 = tmp12389


} else {
tmp12390 := Call(__e, PrimFunc(symshen_4_5_1out), let__12112)


let__12113 := tmp12390
_ = let__12113

tmp12391 := Call(__e, PrimFunc(symshen_4in_1_6), let__12112)


let__12114 := tmp12391
_ = let__12114

tmp12392 := Call(__e, PrimFunc(symshen_4comb), let__12114, let__12113)


ifres12388 = tmp12392


}

let__12111 := ifres12388
_ = let__12111

tmp12395 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12111)


if True == tmp12395 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__12111)
return
}


} else {
__e.Return(let__12107)
return
}


}, 1)

tmp12398 := Call(__e, ns2_1set, symshen_4_5datatype_1rule_6, tmp12379)


_ = tmp12398

tmp12399 := MakeNative(func(__e *ControlFlow) {
V3449 := __e.Get(1)
_ = V3449
tmp12400 := Call(__e, PrimFunc(symshen_4_5sides_6), V3449)


let__12116 := tmp12400
_ = let__12116

tmp12432 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12116)


var ifres12401 Obj

if True == tmp12432 {
tmp12402 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12401 = tmp12402


} else {
tmp12403 := Call(__e, PrimFunc(symshen_4_5_1out), let__12116)


let__12117 := tmp12403
_ = let__12117

tmp12404 := Call(__e, PrimFunc(symshen_4in_1_6), let__12116)


let__12118 := tmp12404
_ = let__12118

tmp12405 := Call(__e, PrimFunc(symshen_4_5prems_6), let__12118)


let__12119 := tmp12405
_ = let__12119

tmp12431 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12119)


var ifres12406 Obj

if True == tmp12431 {
tmp12407 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12406 = tmp12407


} else {
tmp12408 := Call(__e, PrimFunc(symshen_4_5_1out), let__12119)


let__12120 := tmp12408
_ = let__12120

tmp12409 := Call(__e, PrimFunc(symshen_4in_1_6), let__12119)


let__12121 := tmp12409
_ = let__12121

tmp12410 := Call(__e, PrimFunc(symshen_4_5sng_6), let__12121)


let__12122 := tmp12410
_ = let__12122

tmp12430 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12122)


var ifres12411 Obj

if True == tmp12430 {
tmp12412 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12411 = tmp12412


} else {
tmp12413 := Call(__e, PrimFunc(symshen_4in_1_6), let__12122)


let__12123 := tmp12413
_ = let__12123

tmp12414 := Call(__e, PrimFunc(symshen_4_5conc_6), let__12123)


let__12124 := tmp12414
_ = let__12124

tmp12429 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12124)


var ifres12415 Obj

if True == tmp12429 {
tmp12416 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12415 = tmp12416


} else {
tmp12417 := Call(__e, PrimFunc(symshen_4_5_1out), let__12124)


let__12125 := tmp12417
_ = let__12125

tmp12418 := Call(__e, PrimFunc(symshen_4in_1_6), let__12124)


let__12126 := tmp12418
_ = let__12126

tmp12419 := Call(__e, PrimFunc(symshen_4_5sc_6), let__12126)


let__12127 := tmp12419
_ = let__12127

tmp12428 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12127)


var ifres12420 Obj

if True == tmp12428 {
tmp12421 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12420 = tmp12421


} else {
tmp12422 := Call(__e, PrimFunc(symshen_4in_1_6), let__12127)


let__12128 := tmp12422
_ = let__12128

tmp12423 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12125, Nil)
}
__typedArg0 := let__12125
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12424 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12120, tmp12423)
}
__typedArg0 := let__12120
__typedArg1 := tmp12423
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12425 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12117, tmp12424)
}
__typedArg0 := let__12117
__typedArg1 := tmp12424
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12426 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12425, Nil)
}
__typedArg0 := tmp12425
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12427 := Call(__e, PrimFunc(symshen_4comb), let__12128, tmp12426)


ifres12420 = tmp12427


}

ifres12415 = ifres12420


}

ifres12411 = ifres12415


}

ifres12406 = ifres12411


}

ifres12401 = ifres12406


}

let__12115 := ifres12401
_ = let__12115

tmp12434 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12115)


if True == tmp12434 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__12115)
return
}


}, 1)

tmp12435 := Call(__e, ns2_1set, symshen_4_5single_6, tmp12399)


_ = tmp12435

tmp12436 := MakeNative(func(__e *ControlFlow) {
V3464 := __e.Get(1)
_ = V3464
tmp12437 := Call(__e, PrimFunc(symshen_4_5sides_6), V3464)


let__12130 := tmp12437
_ = let__12130

tmp12468 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12130)


var ifres12438 Obj

if True == tmp12468 {
tmp12439 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12438 = tmp12439


} else {
tmp12440 := Call(__e, PrimFunc(symshen_4_5_1out), let__12130)


let__12131 := tmp12440
_ = let__12131

tmp12441 := Call(__e, PrimFunc(symshen_4in_1_6), let__12130)


let__12132 := tmp12441
_ = let__12132

tmp12442 := Call(__e, PrimFunc(symshen_4_5formulae_6), let__12132)


let__12133 := tmp12442
_ = let__12133

tmp12467 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12133)


var ifres12443 Obj

if True == tmp12467 {
tmp12444 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12443 = tmp12444


} else {
tmp12445 := Call(__e, PrimFunc(symshen_4_5_1out), let__12133)


let__12134 := tmp12445
_ = let__12134

tmp12446 := Call(__e, PrimFunc(symshen_4in_1_6), let__12133)


let__12135 := tmp12446
_ = let__12135

tmp12447 := Call(__e, PrimFunc(symshen_4_5dbl_6), let__12135)


let__12136 := tmp12447
_ = let__12136

tmp12466 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12136)


var ifres12448 Obj

if True == tmp12466 {
tmp12449 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12448 = tmp12449


} else {
tmp12450 := Call(__e, PrimFunc(symshen_4in_1_6), let__12136)


let__12137 := tmp12450
_ = let__12137

tmp12451 := Call(__e, PrimFunc(symshen_4_5formula_6), let__12137)


let__12138 := tmp12451
_ = let__12138

tmp12465 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12138)


var ifres12452 Obj

if True == tmp12465 {
tmp12453 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12452 = tmp12453


} else {
tmp12454 := Call(__e, PrimFunc(symshen_4_5_1out), let__12138)


let__12139 := tmp12454
_ = let__12139

tmp12455 := Call(__e, PrimFunc(symshen_4in_1_6), let__12138)


let__12140 := tmp12455
_ = let__12140

tmp12456 := Call(__e, PrimFunc(symshen_4_5sc_6), let__12140)


let__12141 := tmp12456
_ = let__12141

tmp12464 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12141)


var ifres12457 Obj

if True == tmp12464 {
tmp12458 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12457 = tmp12458


} else {
tmp12459 := Call(__e, PrimFunc(symshen_4in_1_6), let__12141)


let__12142 := tmp12459
_ = let__12142

tmp12460 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12139, Nil)
}
__typedArg0 := let__12139
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12461 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Nil, tmp12460)
}
__typedArg0 := Nil
__typedArg1 := tmp12460
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12462 := Call(__e, PrimFunc(symshen_4lr_1rule), let__12131, let__12134, tmp12461)


tmp12463 := Call(__e, PrimFunc(symshen_4comb), let__12142, tmp12462)


ifres12457 = tmp12463


}

ifres12452 = ifres12457


}

ifres12448 = ifres12452


}

ifres12443 = ifres12448


}

ifres12438 = ifres12443


}

let__12129 := ifres12438
_ = let__12129

tmp12470 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12129)


if True == tmp12470 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__12129)
return
}


}, 1)

tmp12471 := Call(__e, ns2_1set, symshen_4_5double_6, tmp12436)


_ = tmp12471

tmp12472 := MakeNative(func(__e *ControlFlow) {
V3479 := __e.Get(1)
_ = V3479
tmp12473 := Call(__e, PrimFunc(symshen_4_5formula_6), V3479)


let__12144 := tmp12473
_ = let__12144

tmp12493 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12144)


var ifres12474 Obj

if True == tmp12493 {
tmp12475 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12474 = tmp12475


} else {
tmp12476 := Call(__e, PrimFunc(symshen_4_5_1out), let__12144)


let__12145 := tmp12476
_ = let__12145

tmp12477 := Call(__e, PrimFunc(symshen_4in_1_6), let__12144)


let__12146 := tmp12477
_ = let__12146

tmp12478 := Call(__e, PrimFunc(symshen_4_5sc_6), let__12146)


let__12147 := tmp12478
_ = let__12147

tmp12492 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12147)


var ifres12479 Obj

if True == tmp12492 {
tmp12480 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12479 = tmp12480


} else {
tmp12481 := Call(__e, PrimFunc(symshen_4in_1_6), let__12147)


let__12148 := tmp12481
_ = let__12148

tmp12482 := Call(__e, PrimFunc(symshen_4_5formulae_6), let__12148)


let__12149 := tmp12482
_ = let__12149

tmp12491 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12149)


var ifres12483 Obj

if True == tmp12491 {
tmp12484 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12483 = tmp12484


} else {
tmp12485 := Call(__e, PrimFunc(symshen_4_5_1out), let__12149)


let__12150 := tmp12485
_ = let__12150

tmp12486 := Call(__e, PrimFunc(symshen_4in_1_6), let__12149)


let__12151 := tmp12486
_ = let__12151

tmp12487 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12145, Nil)
}
__typedArg0 := let__12145
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12488 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Nil, tmp12487)
}
__typedArg0 := Nil
__typedArg1 := tmp12487
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12489 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12488, let__12150)
}
__typedArg0 := tmp12488
__typedArg1 := let__12150
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12490 := Call(__e, PrimFunc(symshen_4comb), let__12151, tmp12489)


ifres12483 = tmp12490


}

ifres12479 = ifres12483


}

ifres12474 = ifres12479


}

let__12143 := ifres12474
_ = let__12143

tmp12512 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12143)


if True == tmp12512 {
tmp12494 := Call(__e, PrimFunc(symshen_4_5formula_6), V3479)


let__12153 := tmp12494
_ = let__12153

tmp12508 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12153)


var ifres12495 Obj

if True == tmp12508 {
tmp12496 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12495 = tmp12496


} else {
tmp12497 := Call(__e, PrimFunc(symshen_4_5_1out), let__12153)


let__12154 := tmp12497
_ = let__12154

tmp12498 := Call(__e, PrimFunc(symshen_4in_1_6), let__12153)


let__12155 := tmp12498
_ = let__12155

tmp12499 := Call(__e, PrimFunc(symshen_4_5sc_6), let__12155)


let__12156 := tmp12499
_ = let__12156

tmp12507 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12156)


var ifres12500 Obj

if True == tmp12507 {
tmp12501 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12500 = tmp12501


} else {
tmp12502 := Call(__e, PrimFunc(symshen_4in_1_6), let__12156)


let__12157 := tmp12502
_ = let__12157

tmp12503 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12154, Nil)
}
__typedArg0 := let__12154
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12504 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Nil, tmp12503)
}
__typedArg0 := Nil
__typedArg1 := tmp12503
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12505 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12504, Nil)
}
__typedArg0 := tmp12504
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12506 := Call(__e, PrimFunc(symshen_4comb), let__12157, tmp12505)


ifres12500 = tmp12506


}

ifres12495 = ifres12500


}

let__12152 := ifres12495
_ = let__12152

tmp12510 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12152)


if True == tmp12510 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__12152)
return
}


} else {
__e.Return(let__12143)
return
}


}, 1)

tmp12513 := Call(__e, ns2_1set, symshen_4_5formulae_6, tmp12472)


_ = tmp12513

tmp12514 := MakeNative(func(__e *ControlFlow) {
V3495 := __e.Get(1)
_ = V3495
tmp12515 := Call(__e, PrimFunc(symshen_4_5ass_6), V3495)


let__12159 := tmp12515
_ = let__12159

tmp12533 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12159)


var ifres12516 Obj

if True == tmp12533 {
tmp12517 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12516 = tmp12517


} else {
tmp12518 := Call(__e, PrimFunc(symshen_4_5_1out), let__12159)


let__12160 := tmp12518
_ = let__12160

tmp12519 := Call(__e, PrimFunc(symshen_4in_1_6), let__12159)


let__12161 := tmp12519
_ = let__12161

tmp12532 := Call(__e, PrimFunc(symshen_4hds_a_2), let__12161, sym_6_6)


var ifres12520 Obj

if True == tmp12532 {
tmp12521 := Call(__e, PrimFunc(symtail), let__12161)


let__12162 := tmp12521
_ = let__12162

tmp12522 := Call(__e, PrimFunc(symshen_4_5formula_6), let__12162)


let__12163 := tmp12522
_ = let__12163

tmp12530 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12163)


var ifres12523 Obj

if True == tmp12530 {
tmp12524 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12523 = tmp12524


} else {
tmp12525 := Call(__e, PrimFunc(symshen_4_5_1out), let__12163)


let__12164 := tmp12525
_ = let__12164

tmp12526 := Call(__e, PrimFunc(symshen_4in_1_6), let__12163)


let__12165 := tmp12526
_ = let__12165

tmp12527 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12164, Nil)
}
__typedArg0 := let__12164
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12528 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12160, tmp12527)
}
__typedArg0 := let__12160
__typedArg1 := tmp12527
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12529 := Call(__e, PrimFunc(symshen_4comb), let__12165, tmp12528)


ifres12523 = tmp12529


}

ifres12520 = ifres12523


} else {
tmp12531 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12520 = tmp12531


}

ifres12516 = ifres12520


}

let__12158 := ifres12516
_ = let__12158

tmp12546 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12158)


if True == tmp12546 {
tmp12534 := Call(__e, PrimFunc(symshen_4_5formula_6), V3495)


let__12167 := tmp12534
_ = let__12167

tmp12542 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12167)


var ifres12535 Obj

if True == tmp12542 {
tmp12536 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12535 = tmp12536


} else {
tmp12537 := Call(__e, PrimFunc(symshen_4_5_1out), let__12167)


let__12168 := tmp12537
_ = let__12168

tmp12538 := Call(__e, PrimFunc(symshen_4in_1_6), let__12167)


let__12169 := tmp12538
_ = let__12169

tmp12539 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12168, Nil)
}
__typedArg0 := let__12168
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12540 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Nil, tmp12539)
}
__typedArg0 := Nil
__typedArg1 := tmp12539
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12541 := Call(__e, PrimFunc(symshen_4comb), let__12169, tmp12540)


ifres12535 = tmp12541


}

let__12166 := ifres12535
_ = let__12166

tmp12544 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12166)


if True == tmp12544 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__12166)
return
}


} else {
__e.Return(let__12158)
return
}


}, 1)

tmp12547 := Call(__e, ns2_1set, symshen_4_5conc_6, tmp12514)


_ = tmp12547

tmp12548 := MakeNative(func(__e *ControlFlow) {
V3508 := __e.Get(1)
_ = V3508
tmp12549 := Call(__e, PrimFunc(symshen_4_5prem_6), V3508)


let__12171 := tmp12549
_ = let__12171

tmp12567 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12171)


var ifres12550 Obj

if True == tmp12567 {
tmp12551 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12550 = tmp12551


} else {
tmp12552 := Call(__e, PrimFunc(symshen_4_5_1out), let__12171)


let__12172 := tmp12552
_ = let__12172

tmp12553 := Call(__e, PrimFunc(symshen_4in_1_6), let__12171)


let__12173 := tmp12553
_ = let__12173

tmp12554 := Call(__e, PrimFunc(symshen_4_5sc_6), let__12173)


let__12174 := tmp12554
_ = let__12174

tmp12566 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12174)


var ifres12555 Obj

if True == tmp12566 {
tmp12556 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12555 = tmp12556


} else {
tmp12557 := Call(__e, PrimFunc(symshen_4in_1_6), let__12174)


let__12175 := tmp12557
_ = let__12175

tmp12558 := Call(__e, PrimFunc(symshen_4_5prems_6), let__12175)


let__12176 := tmp12558
_ = let__12176

tmp12565 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12176)


var ifres12559 Obj

if True == tmp12565 {
tmp12560 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12559 = tmp12560


} else {
tmp12561 := Call(__e, PrimFunc(symshen_4_5_1out), let__12176)


let__12177 := tmp12561
_ = let__12177

tmp12562 := Call(__e, PrimFunc(symshen_4in_1_6), let__12176)


let__12178 := tmp12562
_ = let__12178

tmp12563 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12172, let__12177)
}
__typedArg0 := let__12172
__typedArg1 := let__12177
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12564 := Call(__e, PrimFunc(symshen_4comb), let__12178, tmp12563)


ifres12559 = tmp12564


}

ifres12555 = ifres12559


}

ifres12550 = ifres12555


}

let__12170 := ifres12550
_ = let__12170

tmp12577 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12170)


if True == tmp12577 {
tmp12568 := Call(__e, PrimFunc(sym_5e_6), V3508)


let__12180 := tmp12568
_ = let__12180

tmp12573 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12180)


var ifres12569 Obj

if True == tmp12573 {
tmp12570 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12569 = tmp12570


} else {
tmp12571 := Call(__e, PrimFunc(symshen_4in_1_6), let__12180)


let__12181 := tmp12571
_ = let__12181

tmp12572 := Call(__e, PrimFunc(symshen_4comb), let__12181, Nil)


ifres12569 = tmp12572


}

let__12179 := ifres12569
_ = let__12179

tmp12575 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12179)


if True == tmp12575 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__12179)
return
}


} else {
__e.Return(let__12170)
return
}


}, 1)

tmp12578 := Call(__e, ns2_1set, symshen_4_5prems_6, tmp12548)


_ = tmp12578

tmp12579 := MakeNative(func(__e *ControlFlow) {
V3521 := __e.Get(1)
_ = V3521
tmp12584 := Call(__e, PrimFunc(symshen_4hds_a_2), V3521, sym_b)


var ifres12580 Obj

if True == tmp12584 {
tmp12581 := Call(__e, PrimFunc(symtail), V3521)


let__12183 := tmp12581
_ = let__12183

tmp12582 := Call(__e, PrimFunc(symshen_4comb), let__12183, sym_b)


ifres12580 = tmp12582


} else {
tmp12583 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12580 = tmp12583


}

let__12182 := ifres12580
_ = let__12182

tmp12618 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12182)


if True == tmp12618 {
tmp12585 := Call(__e, PrimFunc(symshen_4_5ass_6), V3521)


let__12185 := tmp12585
_ = let__12185

tmp12603 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12185)


var ifres12586 Obj

if True == tmp12603 {
tmp12587 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12586 = tmp12587


} else {
tmp12588 := Call(__e, PrimFunc(symshen_4_5_1out), let__12185)


let__12186 := tmp12588
_ = let__12186

tmp12589 := Call(__e, PrimFunc(symshen_4in_1_6), let__12185)


let__12187 := tmp12589
_ = let__12187

tmp12602 := Call(__e, PrimFunc(symshen_4hds_a_2), let__12187, sym_6_6)


var ifres12590 Obj

if True == tmp12602 {
tmp12591 := Call(__e, PrimFunc(symtail), let__12187)


let__12188 := tmp12591
_ = let__12188

tmp12592 := Call(__e, PrimFunc(symshen_4_5formula_6), let__12188)


let__12189 := tmp12592
_ = let__12189

tmp12600 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12189)


var ifres12593 Obj

if True == tmp12600 {
tmp12594 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12593 = tmp12594


} else {
tmp12595 := Call(__e, PrimFunc(symshen_4_5_1out), let__12189)


let__12190 := tmp12595
_ = let__12190

tmp12596 := Call(__e, PrimFunc(symshen_4in_1_6), let__12189)


let__12191 := tmp12596
_ = let__12191

tmp12597 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12190, Nil)
}
__typedArg0 := let__12190
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12598 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12186, tmp12597)
}
__typedArg0 := let__12186
__typedArg1 := tmp12597
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12599 := Call(__e, PrimFunc(symshen_4comb), let__12191, tmp12598)


ifres12593 = tmp12599


}

ifres12590 = ifres12593


} else {
tmp12601 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12590 = tmp12601


}

ifres12586 = ifres12590


}

let__12184 := ifres12586
_ = let__12184

tmp12616 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12184)


if True == tmp12616 {
tmp12604 := Call(__e, PrimFunc(symshen_4_5formula_6), V3521)


let__12193 := tmp12604
_ = let__12193

tmp12612 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12193)


var ifres12605 Obj

if True == tmp12612 {
tmp12606 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12605 = tmp12606


} else {
tmp12607 := Call(__e, PrimFunc(symshen_4_5_1out), let__12193)


let__12194 := tmp12607
_ = let__12194

tmp12608 := Call(__e, PrimFunc(symshen_4in_1_6), let__12193)


let__12195 := tmp12608
_ = let__12195

tmp12609 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12194, Nil)
}
__typedArg0 := let__12194
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12610 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Nil, tmp12609)
}
__typedArg0 := Nil
__typedArg1 := tmp12609
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12611 := Call(__e, PrimFunc(symshen_4comb), let__12195, tmp12610)


ifres12605 = tmp12611


}

let__12192 := ifres12605
_ = let__12192

tmp12614 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12192)


if True == tmp12614 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__12192)
return
}


} else {
__e.Return(let__12184)
return
}


} else {
__e.Return(let__12182)
return
}


}, 1)

tmp12619 := Call(__e, ns2_1set, symshen_4_5prem_6, tmp12579)


_ = tmp12619

tmp12620 := MakeNative(func(__e *ControlFlow) {
V3536 := __e.Get(1)
_ = V3536
tmp12621 := Call(__e, PrimFunc(symshen_4_5formula_6), V3536)


let__12197 := tmp12621
_ = let__12197

tmp12639 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12197)


var ifres12622 Obj

if True == tmp12639 {
tmp12623 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12622 = tmp12623


} else {
tmp12624 := Call(__e, PrimFunc(symshen_4_5_1out), let__12197)


let__12198 := tmp12624
_ = let__12198

tmp12625 := Call(__e, PrimFunc(symshen_4in_1_6), let__12197)


let__12199 := tmp12625
_ = let__12199

tmp12626 := Call(__e, PrimFunc(symshen_4_5iscomma_6), let__12199)


let__12200 := tmp12626
_ = let__12200

tmp12638 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12200)


var ifres12627 Obj

if True == tmp12638 {
tmp12628 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12627 = tmp12628


} else {
tmp12629 := Call(__e, PrimFunc(symshen_4in_1_6), let__12200)


let__12201 := tmp12629
_ = let__12201

tmp12630 := Call(__e, PrimFunc(symshen_4_5ass_6), let__12201)


let__12202 := tmp12630
_ = let__12202

tmp12637 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12202)


var ifres12631 Obj

if True == tmp12637 {
tmp12632 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12631 = tmp12632


} else {
tmp12633 := Call(__e, PrimFunc(symshen_4_5_1out), let__12202)


let__12203 := tmp12633
_ = let__12203

tmp12634 := Call(__e, PrimFunc(symshen_4in_1_6), let__12202)


let__12204 := tmp12634
_ = let__12204

tmp12635 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12198, let__12203)
}
__typedArg0 := let__12198
__typedArg1 := let__12203
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12636 := Call(__e, PrimFunc(symshen_4comb), let__12204, tmp12635)


ifres12631 = tmp12636


}

ifres12627 = ifres12631


}

ifres12622 = ifres12627


}

let__12196 := ifres12622
_ = let__12196

tmp12659 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12196)


if True == tmp12659 {
tmp12640 := Call(__e, PrimFunc(symshen_4_5formula_6), V3536)


let__12206 := tmp12640
_ = let__12206

tmp12647 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12206)


var ifres12641 Obj

if True == tmp12647 {
tmp12642 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12641 = tmp12642


} else {
tmp12643 := Call(__e, PrimFunc(symshen_4_5_1out), let__12206)


let__12207 := tmp12643
_ = let__12207

tmp12644 := Call(__e, PrimFunc(symshen_4in_1_6), let__12206)


let__12208 := tmp12644
_ = let__12208

tmp12645 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12207, Nil)
}
__typedArg0 := let__12207
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12646 := Call(__e, PrimFunc(symshen_4comb), let__12208, tmp12645)


ifres12641 = tmp12646


}

let__12205 := ifres12641
_ = let__12205

tmp12657 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12205)


if True == tmp12657 {
tmp12648 := Call(__e, PrimFunc(sym_5e_6), V3536)


let__12210 := tmp12648
_ = let__12210

tmp12653 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12210)


var ifres12649 Obj

if True == tmp12653 {
tmp12650 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12649 = tmp12650


} else {
tmp12651 := Call(__e, PrimFunc(symshen_4in_1_6), let__12210)


let__12211 := tmp12651
_ = let__12211

tmp12652 := Call(__e, PrimFunc(symshen_4comb), let__12211, Nil)


ifres12649 = tmp12652


}

let__12209 := ifres12649
_ = let__12209

tmp12655 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12209)


if True == tmp12655 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__12209)
return
}


} else {
__e.Return(let__12205)
return
}


} else {
__e.Return(let__12196)
return
}


}, 1)

tmp12660 := Call(__e, ns2_1set, symshen_4_5ass_6, tmp12620)


_ = tmp12660

tmp12661 := MakeNative(func(__e *ControlFlow) {
V3553 := __e.Get(1)
_ = V3553
tmp12671 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3553)
}
__typedArg0 := V3553
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12662 Obj

if True == tmp12671 {
tmp12663 := Call(__e, PrimFunc(symhead), V3553)


let__12213 := tmp12663
_ = let__12213

tmp12664 := Call(__e, PrimFunc(symtail), V3553)


let__12214 := tmp12664
_ = let__12214

tmp12668 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(","))
}
__typedArg0 := MakeString(",")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp12669 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__12213, tmp12668)
}
__typedArg0 := let__12213
__typedArg1 := tmp12668
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres12665 Obj

if True == tmp12669 {
tmp12666 := Call(__e, PrimFunc(symshen_4comb), let__12214, symshen_4skip)


ifres12665 = tmp12666


} else {
tmp12667 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12665 = tmp12667


}

ifres12662 = ifres12665


} else {
tmp12670 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12662 = tmp12670


}

let__12212 := ifres12662
_ = let__12212

tmp12673 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12212)


if True == tmp12673 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__12212)
return
}


}, 1)

tmp12674 := Call(__e, ns2_1set, symshen_4_5iscomma_6, tmp12661)


_ = tmp12674

tmp12675 := MakeNative(func(__e *ControlFlow) {
V3557 := __e.Get(1)
_ = V3557
tmp12676 := Call(__e, PrimFunc(symshen_4_5expr_6), V3557)


let__12216 := tmp12676
_ = let__12216

tmp12699 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12216)


var ifres12677 Obj

if True == tmp12699 {
tmp12678 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12677 = tmp12678


} else {
tmp12679 := Call(__e, PrimFunc(symshen_4_5_1out), let__12216)


let__12217 := tmp12679
_ = let__12217

tmp12680 := Call(__e, PrimFunc(symshen_4in_1_6), let__12216)


let__12218 := tmp12680
_ = let__12218

tmp12681 := Call(__e, PrimFunc(symshen_4_5iscolon_6), let__12218)


let__12219 := tmp12681
_ = let__12219

tmp12698 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12219)


var ifres12682 Obj

if True == tmp12698 {
tmp12683 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12682 = tmp12683


} else {
tmp12684 := Call(__e, PrimFunc(symshen_4in_1_6), let__12219)


let__12220 := tmp12684
_ = let__12220

tmp12685 := Call(__e, PrimFunc(symshen_4_5type_6), let__12220)


let__12221 := tmp12685
_ = let__12221

tmp12697 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12221)


var ifres12686 Obj

if True == tmp12697 {
tmp12687 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12686 = tmp12687


} else {
tmp12688 := Call(__e, PrimFunc(symshen_4_5_1out), let__12221)


let__12222 := tmp12688
_ = let__12222

tmp12689 := Call(__e, PrimFunc(symshen_4in_1_6), let__12221)


let__12223 := tmp12689
_ = let__12223

tmp12690 := Call(__e, PrimFunc(symshen_4curry), let__12217)


tmp12691 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp12692 := Call(__e, PrimFunc(symshen_4rectify_1type), let__12222)


tmp12693 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12692, Nil)
}
__typedArg0 := tmp12692
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12694 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12691, tmp12693)
}
__typedArg0 := tmp12691
__typedArg1 := tmp12693
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12695 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12690, tmp12694)
}
__typedArg0 := tmp12690
__typedArg1 := tmp12694
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12696 := Call(__e, PrimFunc(symshen_4comb), let__12223, tmp12695)


ifres12686 = tmp12696


}

ifres12682 = ifres12686


}

ifres12677 = ifres12682


}

let__12215 := ifres12677
_ = let__12215

tmp12710 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12215)


if True == tmp12710 {
tmp12700 := Call(__e, PrimFunc(symshen_4_5expr_6), V3557)


let__12225 := tmp12700
_ = let__12225

tmp12706 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12225)


var ifres12701 Obj

if True == tmp12706 {
tmp12702 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12701 = tmp12702


} else {
tmp12703 := Call(__e, PrimFunc(symshen_4_5_1out), let__12225)


let__12226 := tmp12703
_ = let__12226

tmp12704 := Call(__e, PrimFunc(symshen_4in_1_6), let__12225)


let__12227 := tmp12704
_ = let__12227

tmp12705 := Call(__e, PrimFunc(symshen_4comb), let__12227, let__12226)


ifres12701 = tmp12705


}

let__12224 := ifres12701
_ = let__12224

tmp12708 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12224)


if True == tmp12708 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__12224)
return
}


} else {
__e.Return(let__12215)
return
}


}, 1)

tmp12711 := Call(__e, ns2_1set, symshen_4_5formula_6, tmp12675)


_ = tmp12711

tmp12712 := MakeNative(func(__e *ControlFlow) {
V3571 := __e.Get(1)
_ = V3571
tmp12722 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3571)
}
__typedArg0 := V3571
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12713 Obj

if True == tmp12722 {
tmp12714 := Call(__e, PrimFunc(symhead), V3571)


let__12229 := tmp12714
_ = let__12229

tmp12715 := Call(__e, PrimFunc(symtail), V3571)


let__12230 := tmp12715
_ = let__12230

tmp12719 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp12720 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__12229, tmp12719)
}
__typedArg0 := let__12229
__typedArg1 := tmp12719
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres12716 Obj

if True == tmp12720 {
tmp12717 := Call(__e, PrimFunc(symshen_4comb), let__12230, symshen_4skip)


ifres12716 = tmp12717


} else {
tmp12718 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12716 = tmp12718


}

ifres12713 = ifres12716


} else {
tmp12721 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12713 = tmp12721


}

let__12228 := ifres12713
_ = let__12228

tmp12724 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12228)


if True == tmp12724 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__12228)
return
}


}, 1)

tmp12725 := Call(__e, ns2_1set, symshen_4_5iscolon_6, tmp12712)


_ = tmp12725

tmp12726 := MakeNative(func(__e *ControlFlow) {
V3575 := __e.Get(1)
_ = V3575
tmp12727 := Call(__e, PrimFunc(symshen_4_5side_6), V3575)


let__12232 := tmp12727
_ = let__12232

tmp12740 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12232)


var ifres12728 Obj

if True == tmp12740 {
tmp12729 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12728 = tmp12729


} else {
tmp12730 := Call(__e, PrimFunc(symshen_4_5_1out), let__12232)


let__12233 := tmp12730
_ = let__12233

tmp12731 := Call(__e, PrimFunc(symshen_4in_1_6), let__12232)


let__12234 := tmp12731
_ = let__12234

tmp12732 := Call(__e, PrimFunc(symshen_4_5sides_6), let__12234)


let__12235 := tmp12732
_ = let__12235

tmp12739 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12235)


var ifres12733 Obj

if True == tmp12739 {
tmp12734 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12733 = tmp12734


} else {
tmp12735 := Call(__e, PrimFunc(symshen_4_5_1out), let__12235)


let__12236 := tmp12735
_ = let__12236

tmp12736 := Call(__e, PrimFunc(symshen_4in_1_6), let__12235)


let__12237 := tmp12736
_ = let__12237

tmp12737 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12233, let__12236)
}
__typedArg0 := let__12233
__typedArg1 := let__12236
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12738 := Call(__e, PrimFunc(symshen_4comb), let__12237, tmp12737)


ifres12733 = tmp12738


}

ifres12728 = ifres12733


}

let__12231 := ifres12728
_ = let__12231

tmp12750 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12231)


if True == tmp12750 {
tmp12741 := Call(__e, PrimFunc(sym_5e_6), V3575)


let__12239 := tmp12741
_ = let__12239

tmp12746 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12239)


var ifres12742 Obj

if True == tmp12746 {
tmp12743 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12742 = tmp12743


} else {
tmp12744 := Call(__e, PrimFunc(symshen_4in_1_6), let__12239)


let__12240 := tmp12744
_ = let__12240

tmp12745 := Call(__e, PrimFunc(symshen_4comb), let__12240, Nil)


ifres12742 = tmp12745


}

let__12238 := ifres12742
_ = let__12238

tmp12748 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12238)


if True == tmp12748 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__12238)
return
}


} else {
__e.Return(let__12231)
return
}


}, 1)

tmp12751 := Call(__e, ns2_1set, symshen_4_5sides_6, tmp12726)


_ = tmp12751

tmp12752 := MakeNative(func(__e *ControlFlow) {
V3586 := __e.Get(1)
_ = V3586
tmp12764 := Call(__e, PrimFunc(symshen_4hds_a_2), V3586, symif)


var ifres12753 Obj

if True == tmp12764 {
tmp12754 := Call(__e, PrimFunc(symtail), V3586)


let__12242 := tmp12754
_ = let__12242

tmp12762 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__12242)
}
__typedArg0 := let__12242
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12755 Obj

if True == tmp12762 {
tmp12756 := Call(__e, PrimFunc(symhead), let__12242)


let__12243 := tmp12756
_ = let__12243

tmp12757 := Call(__e, PrimFunc(symtail), let__12242)


let__12244 := tmp12757
_ = let__12244

tmp12758 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12243, Nil)
}
__typedArg0 := let__12243
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12759 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp12758)
}
__typedArg0 := symif
__typedArg1 := tmp12758
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12760 := Call(__e, PrimFunc(symshen_4comb), let__12244, tmp12759)


ifres12755 = tmp12760


} else {
tmp12761 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12755 = tmp12761


}

ifres12753 = ifres12755


} else {
tmp12763 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12753 = tmp12763


}

let__12241 := ifres12753
_ = let__12241

tmp12820 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12241)


if True == tmp12820 {
tmp12782 := Call(__e, PrimFunc(symshen_4hds_a_2), V3586, symlet)


var ifres12765 Obj

if True == tmp12782 {
tmp12766 := Call(__e, PrimFunc(symtail), V3586)


let__12246 := tmp12766
_ = let__12246

tmp12780 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__12246)
}
__typedArg0 := let__12246
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12767 Obj

if True == tmp12780 {
tmp12768 := Call(__e, PrimFunc(symhead), let__12246)


let__12247 := tmp12768
_ = let__12247

tmp12769 := Call(__e, PrimFunc(symtail), let__12246)


let__12248 := tmp12769
_ = let__12248

tmp12778 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__12248)
}
__typedArg0 := let__12248
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12770 Obj

if True == tmp12778 {
tmp12771 := Call(__e, PrimFunc(symhead), let__12248)


let__12249 := tmp12771
_ = let__12249

tmp12772 := Call(__e, PrimFunc(symtail), let__12248)


let__12250 := tmp12772
_ = let__12250

tmp12773 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12249, Nil)
}
__typedArg0 := let__12249
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12774 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12247, tmp12773)
}
__typedArg0 := let__12247
__typedArg1 := tmp12773
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12775 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp12774)
}
__typedArg0 := symlet
__typedArg1 := tmp12774
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12776 := Call(__e, PrimFunc(symshen_4comb), let__12250, tmp12775)


ifres12770 = tmp12776


} else {
tmp12777 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12770 = tmp12777


}

ifres12767 = ifres12770


} else {
tmp12779 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12767 = tmp12779


}

ifres12765 = ifres12767


} else {
tmp12781 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12765 = tmp12781


}

let__12245 := ifres12765
_ = let__12245

tmp12818 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12245)


if True == tmp12818 {
tmp12797 := Call(__e, PrimFunc(symshen_4hds_a_2), V3586, symctxt)


var ifres12783 Obj

if True == tmp12797 {
tmp12784 := Call(__e, PrimFunc(symtail), V3586)


let__12252 := tmp12784
_ = let__12252

tmp12795 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__12252)
}
__typedArg0 := let__12252
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12785 Obj

if True == tmp12795 {
tmp12786 := Call(__e, PrimFunc(symhead), let__12252)


let__12253 := tmp12786
_ = let__12253

tmp12787 := Call(__e, PrimFunc(symtail), let__12252)


let__12254 := tmp12787
_ = let__12254

tmp12793 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(let__12253)
}
__typedArg0 := let__12253
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

var ifres12788 Obj

if True == tmp12793 {
tmp12789 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12253, Nil)
}
__typedArg0 := let__12253
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12790 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symctxt, tmp12789)
}
__typedArg0 := symctxt
__typedArg1 := tmp12789
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12791 := Call(__e, PrimFunc(symshen_4comb), let__12254, tmp12790)


ifres12788 = tmp12791


} else {
tmp12792 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12788 = tmp12792


}

ifres12785 = ifres12788


} else {
tmp12794 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12785 = tmp12794


}

ifres12783 = ifres12785


} else {
tmp12796 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12783 = tmp12796


}

let__12251 := ifres12783
_ = let__12251

tmp12816 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12251)


if True == tmp12816 {
tmp12812 := Call(__e, PrimFunc(symshen_4hds_a_2), V3586, symsqts)


var ifres12798 Obj

if True == tmp12812 {
tmp12799 := Call(__e, PrimFunc(symtail), V3586)


let__12256 := tmp12799
_ = let__12256

tmp12810 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__12256)
}
__typedArg0 := let__12256
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12800 Obj

if True == tmp12810 {
tmp12801 := Call(__e, PrimFunc(symhead), let__12256)


let__12257 := tmp12801
_ = let__12257

tmp12802 := Call(__e, PrimFunc(symtail), let__12256)


let__12258 := tmp12802
_ = let__12258

tmp12808 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(let__12257)
}
__typedArg0 := let__12257
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

var ifres12803 Obj

if True == tmp12808 {
tmp12804 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12257, Nil)
}
__typedArg0 := let__12257
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12805 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsqts, tmp12804)
}
__typedArg0 := symsqts
__typedArg1 := tmp12804
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12806 := Call(__e, PrimFunc(symshen_4comb), let__12258, tmp12805)


ifres12803 = tmp12806


} else {
tmp12807 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12803 = tmp12807


}

ifres12800 = ifres12803


} else {
tmp12809 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12800 = tmp12809


}

ifres12798 = ifres12800


} else {
tmp12811 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12798 = tmp12811


}

let__12255 := ifres12798
_ = let__12255

tmp12814 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12255)


if True == tmp12814 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__12255)
return
}


} else {
__e.Return(let__12251)
return
}


} else {
__e.Return(let__12245)
return
}


} else {
__e.Return(let__12241)
return
}


}, 1)

tmp12821 := Call(__e, ns2_1set, symshen_4_5side_6, tmp12752)


_ = tmp12821

tmp12822 := MakeNative(func(__e *ControlFlow) {
V3611 := __e.Get(1)
_ = V3611
V3612 := __e.Get(2)
_ = V3612
V3613 := __e.Get(3)
_ = V3613
tmp12852 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3613)
}
__typedArg0 := V3613
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12839 Obj

if True == tmp12852 {
tmp12850 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3613)
}
__typedArg0 := V3613
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12851 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp12850)
}
__typedArg0 := Nil
__typedArg1 := tmp12850
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres12841 Obj

if True == tmp12851 {
tmp12848 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3613)
}
__typedArg0 := V3613
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12849 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp12848)
}
__typedArg0 := tmp12848
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12843 Obj

if True == tmp12849 {
tmp12845 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3613)
}
__typedArg0 := V3613
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12846 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp12845)
}
__typedArg0 := tmp12845
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12847 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp12846)
}
__typedArg0 := Nil
__typedArg1 := tmp12846
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres12844 Obj

if True == tmp12847 {
ifres12844 = True


} else {
ifres12844 = False


}

ifres12843 = ifres12844


} else {
ifres12843 = False


}

var ifres12842 Obj

if True == ifres12843 {
ifres12842 = True


} else {
ifres12842 = False


}

ifres12841 = ifres12842


} else {
ifres12841 = False


}

var ifres12840 Obj

if True == ifres12841 {
ifres12840 = True


} else {
ifres12840 = False


}

ifres12839 = ifres12840


} else {
ifres12839 = False


}

if True == ifres12839 {
tmp12823 := Call(__e, PrimFunc(symgensym), symP)


let__12259 := tmp12823
_ = let__12259

tmp12824 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3613)
}
__typedArg0 := V3613
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12825 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12259, Nil)
}
__typedArg0 := let__12259
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12826 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12824, tmp12825)
}
__typedArg0 := tmp12824
__typedArg1 := tmp12825
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__12260 := tmp12826
_ = let__12260

tmp12827 := Call(__e, PrimFunc(symshen_4coll_1formulae), V3612)


tmp12828 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12259, Nil)
}
__typedArg0 := let__12259
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12829 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12827, tmp12828)
}
__typedArg0 := tmp12827
__typedArg1 := tmp12828
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__12261 := tmp12829
_ = let__12261

tmp12830 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12261, Nil)
}
__typedArg0 := let__12261
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12831 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12260, Nil)
}
__typedArg0 := let__12260
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12832 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12830, tmp12831)
}
__typedArg0 := tmp12830
__typedArg1 := tmp12831
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12833 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3611, tmp12832)
}
__typedArg0 := V3611
__typedArg1 := tmp12832
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__12262 := tmp12833
_ = let__12262

tmp12834 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3613, Nil)
}
__typedArg0 := V3613
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12835 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3612, tmp12834)
}
__typedArg0 := V3612
__typedArg1 := tmp12834
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12836 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3611, tmp12835)
}
__typedArg0 := V3611
__typedArg1 := tmp12835
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__12263 := tmp12836
_ = let__12263

tmp12837 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12262, Nil)
}
__typedArg0 := let__12262
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12263, tmp12837)
}
__typedArg0 := let__12263
__typedArg1 := tmp12837
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

tmp12853 := Call(__e, ns2_1set, symshen_4lr_1rule, tmp12822)


_ = tmp12853

tmp12854 := MakeNative(func(__e *ControlFlow) {
V3621 := __e.Get(1)
_ = V3621
tmp12883 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3621)
}
__typedArg0 := Nil
__typedArg1 := V3621
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp12883 {
__e.Return(Nil)
return
} else {
tmp12881 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3621)
}
__typedArg0 := V3621
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12861 Obj

if True == tmp12881 {
tmp12879 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3621)
}
__typedArg0 := V3621
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12880 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp12879)
}
__typedArg0 := tmp12879
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12863 Obj

if True == tmp12880 {
tmp12876 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3621)
}
__typedArg0 := V3621
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12877 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp12876)
}
__typedArg0 := tmp12876
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12878 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp12877)
}
__typedArg0 := Nil
__typedArg1 := tmp12877
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres12865 Obj

if True == tmp12878 {
tmp12873 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3621)
}
__typedArg0 := V3621
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12874 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp12873)
}
__typedArg0 := tmp12873
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12875 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp12874)
}
__typedArg0 := tmp12874
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12867 Obj

if True == tmp12875 {
tmp12869 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3621)
}
__typedArg0 := V3621
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12870 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp12869)
}
__typedArg0 := tmp12869
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12871 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp12870)
}
__typedArg0 := tmp12870
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12872 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp12871)
}
__typedArg0 := Nil
__typedArg1 := tmp12871
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres12868 Obj

if True == tmp12872 {
ifres12868 = True


} else {
ifres12868 = False


}

ifres12867 = ifres12868


} else {
ifres12867 = False


}

var ifres12866 Obj

if True == ifres12867 {
ifres12866 = True


} else {
ifres12866 = False


}

ifres12865 = ifres12866


} else {
ifres12865 = False


}

var ifres12864 Obj

if True == ifres12865 {
ifres12864 = True


} else {
ifres12864 = False


}

ifres12863 = ifres12864


} else {
ifres12863 = False


}

var ifres12862 Obj

if True == ifres12863 {
ifres12862 = True


} else {
ifres12862 = False


}

ifres12861 = ifres12862


} else {
ifres12861 = False


}

if True == ifres12861 {
tmp12855 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3621)
}
__typedArg0 := V3621
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12856 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp12855)
}
__typedArg0 := tmp12855
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12857 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp12856)
}
__typedArg0 := tmp12856
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12858 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3621)
}
__typedArg0 := V3621
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12859 := Call(__e, PrimFunc(symshen_4coll_1formulae), tmp12858)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12857, tmp12859)
}
__typedArg0 := tmp12857
__typedArg1 := tmp12859
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

tmp12884 := Call(__e, ns2_1set, symshen_4coll_1formulae, tmp12854)


_ = tmp12884

tmp12885 := MakeNative(func(__e *ControlFlow) {
V3622 := __e.Get(1)
_ = V3622
tmp12896 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3622)
}
__typedArg0 := V3622
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12886 Obj

if True == tmp12896 {
tmp12887 := Call(__e, PrimFunc(symhead), V3622)


let__12265 := tmp12887
_ = let__12265

tmp12888 := Call(__e, PrimFunc(symtail), V3622)


let__12266 := tmp12888
_ = let__12266

tmp12893 := Call(__e, PrimFunc(symshen_4key_1in_1sequent_1calculus_2), let__12265)


tmp12894 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp12893)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp12893
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres12889 Obj

if True == tmp12894 {
tmp12890 := Call(__e, PrimFunc(symmacroexpand), let__12265)


tmp12891 := Call(__e, PrimFunc(symshen_4comb), let__12266, tmp12890)


ifres12889 = tmp12891


} else {
tmp12892 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12889 = tmp12892


}

ifres12886 = ifres12889


} else {
tmp12895 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12886 = tmp12895


}

let__12264 := ifres12886
_ = let__12264

tmp12898 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12264)


if True == tmp12898 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__12264)
return
}


}, 1)

tmp12899 := Call(__e, ns2_1set, symshen_4_5expr_6, tmp12885)


_ = tmp12899

tmp12900 := MakeNative(func(__e *ControlFlow) {
V3626 := __e.Get(1)
_ = V3626
tmp12907 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(";"))
}
__typedArg0 := MakeString(";")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp12908 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(","))
}
__typedArg0 := MakeString(",")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp12909 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp12910 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1_1, Nil)
}
__typedArg0 := sym_5_1_1
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12911 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12909, tmp12910)
}
__typedArg0 := tmp12909
__typedArg1 := tmp12910
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12912 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12908, tmp12911)
}
__typedArg0 := tmp12908
__typedArg1 := tmp12911
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12913 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12907, tmp12912)
}
__typedArg0 := tmp12907
__typedArg1 := tmp12912
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12914 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_6_6, tmp12913)
}
__typedArg0 := sym_6_6
__typedArg1 := tmp12913
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12915 := Call(__e, PrimFunc(symelement_2), V3626, tmp12914)


if True == tmp12915 {
__e.Return(True)
return
} else {
tmp12905 := Call(__e, PrimFunc(symshen_4sng_2), V3626)


var ifres12902 Obj

if True == tmp12905 {
ifres12902 = True


} else {
tmp12904 := Call(__e, PrimFunc(symshen_4dbl_2), V3626)


var ifres12903 Obj

if True == tmp12904 {
ifres12903 = True


} else {
ifres12903 = False


}

ifres12902 = ifres12903


}

if True == ifres12902 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp12916 := Call(__e, ns2_1set, symshen_4key_1in_1sequent_1calculus_2, tmp12900)


_ = tmp12916

tmp12917 := MakeNative(func(__e *ControlFlow) {
V3627 := __e.Get(1)
_ = V3627
tmp12918 := Call(__e, PrimFunc(symshen_4_5expr_6), V3627)


let__12268 := tmp12918
_ = let__12268

tmp12924 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12268)


var ifres12919 Obj

if True == tmp12924 {
tmp12920 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12919 = tmp12920


} else {
tmp12921 := Call(__e, PrimFunc(symshen_4_5_1out), let__12268)


let__12269 := tmp12921
_ = let__12269

tmp12922 := Call(__e, PrimFunc(symshen_4in_1_6), let__12268)


let__12270 := tmp12922
_ = let__12270

tmp12923 := Call(__e, PrimFunc(symshen_4comb), let__12270, let__12269)


ifres12919 = tmp12923


}

let__12267 := ifres12919
_ = let__12267

tmp12926 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12267)


if True == tmp12926 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__12267)
return
}


}, 1)

tmp12927 := Call(__e, ns2_1set, symshen_4_5type_6, tmp12917)


_ = tmp12927

tmp12928 := MakeNative(func(__e *ControlFlow) {
V3632 := __e.Get(1)
_ = V3632
tmp12937 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3632)
}
__typedArg0 := V3632
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12929 Obj

if True == tmp12937 {
tmp12930 := Call(__e, PrimFunc(symhead), V3632)


let__12272 := tmp12930
_ = let__12272

tmp12931 := Call(__e, PrimFunc(symtail), V3632)


let__12273 := tmp12931
_ = let__12273

tmp12935 := Call(__e, PrimFunc(symshen_4dbl_2), let__12272)


var ifres12932 Obj

if True == tmp12935 {
tmp12933 := Call(__e, PrimFunc(symshen_4comb), let__12273, let__12272)


ifres12932 = tmp12933


} else {
tmp12934 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12932 = tmp12934


}

ifres12929 = ifres12932


} else {
tmp12936 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12929 = tmp12936


}

let__12271 := ifres12929
_ = let__12271

tmp12939 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12271)


if True == tmp12939 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__12271)
return
}


}, 1)

tmp12940 := Call(__e, ns2_1set, symshen_4_5dbl_6, tmp12928)


_ = tmp12940

tmp12941 := MakeNative(func(__e *ControlFlow) {
V3636 := __e.Get(1)
_ = V3636
tmp12950 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3636)
}
__typedArg0 := V3636
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12942 Obj

if True == tmp12950 {
tmp12943 := Call(__e, PrimFunc(symhead), V3636)


let__12275 := tmp12943
_ = let__12275

tmp12944 := Call(__e, PrimFunc(symtail), V3636)


let__12276 := tmp12944
_ = let__12276

tmp12948 := Call(__e, PrimFunc(symshen_4sng_2), let__12275)


var ifres12945 Obj

if True == tmp12948 {
tmp12946 := Call(__e, PrimFunc(symshen_4comb), let__12276, let__12275)


ifres12945 = tmp12946


} else {
tmp12947 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12945 = tmp12947


}

ifres12942 = ifres12945


} else {
tmp12949 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12942 = tmp12949


}

let__12274 := ifres12942
_ = let__12274

tmp12952 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12274)


if True == tmp12952 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__12274)
return
}


}, 1)

tmp12953 := Call(__e, ns2_1set, symshen_4_5sng_6, tmp12941)


_ = tmp12953

tmp12954 := MakeNative(func(__e *ControlFlow) {
V3640 := __e.Get(1)
_ = V3640
tmp12959 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsymbol_2) {
return PrimIsSymbol(V3640)
}
__typedArg0 := V3640
return Call(__e, PrimFunc(symsymbol_2), __typedArg0)
})()

if True == tmp12959 {
tmp12956 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(V3640)
}
__typedArg0 := V3640
return Call(__e, PrimFunc(symstr), __typedArg0)
})()

tmp12957 := Call(__e, PrimFunc(symshen_4sng_1h_2), tmp12956)


if True == tmp12957 {
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

tmp12960 := Call(__e, ns2_1set, symshen_4sng_2, tmp12954)


_ = tmp12960

tmp12961 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__selftop:
V3643 := __self1
_ = V3643
tmp12970 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("___"), V3643)
}
__typedArg0 := MakeString("___")
__typedArg1 := V3643
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp12970 {
__e.Return(True)
return
} else {
tmp12968 := Call(__e, PrimFunc(symshen_4_7string_2), V3643)


var ifres12964 Obj

if True == tmp12968 {
tmp12966 := Call(__e, PrimFunc(symhdstr), V3643)


tmp12967 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("_"), tmp12966)
}
__typedArg0 := MakeString("_")
__typedArg1 := tmp12966
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres12965 Obj

if True == tmp12967 {
ifres12965 = True


} else {
ifres12965 = False


}

ifres12964 = ifres12965


} else {
ifres12964 = False


}

if True == ifres12964 {
if PrimFunc(symshen_4sng_1h_2) == __self {
__self1 = (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V3643)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V3643
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4sng_1h_2), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V3643)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V3643
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp12971 := Call(__e, ns2_1set, symshen_4sng_1h_2, tmp12961)


_ = tmp12971

tmp12972 := MakeNative(func(__e *ControlFlow) {
V3644 := __e.Get(1)
_ = V3644
tmp12977 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsymbol_2) {
return PrimIsSymbol(V3644)
}
__typedArg0 := V3644
return Call(__e, PrimFunc(symsymbol_2), __typedArg0)
})()

if True == tmp12977 {
tmp12974 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(V3644)
}
__typedArg0 := V3644
return Call(__e, PrimFunc(symstr), __typedArg0)
})()

tmp12975 := Call(__e, PrimFunc(symshen_4dbl_1h_2), tmp12974)


if True == tmp12975 {
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

tmp12978 := Call(__e, ns2_1set, symshen_4dbl_2, tmp12972)


_ = tmp12978

tmp12979 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__selftop:
V3647 := __self1
_ = V3647
tmp12988 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("==="), V3647)
}
__typedArg0 := MakeString("===")
__typedArg1 := V3647
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp12988 {
__e.Return(True)
return
} else {
tmp12986 := Call(__e, PrimFunc(symshen_4_7string_2), V3647)


var ifres12982 Obj

if True == tmp12986 {
tmp12984 := Call(__e, PrimFunc(symhdstr), V3647)


tmp12985 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("="), tmp12984)
}
__typedArg0 := MakeString("=")
__typedArg1 := tmp12984
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres12983 Obj

if True == tmp12985 {
ifres12983 = True


} else {
ifres12983 = False


}

ifres12982 = ifres12983


} else {
ifres12982 = False


}

if True == ifres12982 {
if PrimFunc(symshen_4dbl_1h_2) == __self {
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
__e.TailApply(PrimFunc(symshen_4dbl_1h_2), (func() Obj {
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

tmp12989 := Call(__e, ns2_1set, symshen_4dbl_1h_2, tmp12979)


_ = tmp12989

tmp12990 := MakeNative(func(__e *ControlFlow) {
V3648 := __e.Get(1)
_ = V3648
V3649 := __e.Get(2)
_ = V3649
tmp12991 := MakeNative(func(__e *ControlFlow) {
Z3651 := __e.Get(1)
_ = Z3651
__e.TailApply(PrimFunc(symshen_4rule_1_6clause), Z3651)
return
}, 1)

tmp12992 := Call(__e, PrimFunc(symmapcan), tmp12991, V3649)


let__12277 := tmp12992
_ = let__12277

tmp12993 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3648, let__12277)
}
__typedArg0 := V3648
__typedArg1 := let__12277
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12994 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefprolog, tmp12993)
}
__typedArg0 := symdefprolog
__typedArg1 := tmp12993
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__12278 := tmp12994
_ = let__12278

__e.TailApply(PrimFunc(symeval), let__12278)
return


}, 2)

tmp12995 := Call(__e, ns2_1set, symshen_4rules_1_6prolog, tmp12990)


_ = tmp12995

tmp12996 := MakeNative(func(__e *ControlFlow) {
V3653 := __e.Get(1)
_ = V3653
tmp13056 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3653)
}
__typedArg0 := V3653
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13020 Obj

if True == tmp13056 {
tmp13054 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3653)
}
__typedArg0 := V3653
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13055 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13054)
}
__typedArg0 := tmp13054
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13022 Obj

if True == tmp13055 {
tmp13051 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3653)
}
__typedArg0 := V3653
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13052 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13051)
}
__typedArg0 := tmp13051
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13053 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13052)
}
__typedArg0 := tmp13052
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13024 Obj

if True == tmp13053 {
tmp13047 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3653)
}
__typedArg0 := V3653
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13048 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13047)
}
__typedArg0 := tmp13047
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13049 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13048)
}
__typedArg0 := tmp13048
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13050 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13049)
}
__typedArg0 := tmp13049
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13026 Obj

if True == tmp13050 {
tmp13042 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3653)
}
__typedArg0 := V3653
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13043 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13042)
}
__typedArg0 := tmp13042
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13044 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13043)
}
__typedArg0 := tmp13043
return Call(__e, PrimFunc(symhd), __typedArg0)
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

var ifres13028 Obj

if True == tmp13046 {
tmp13036 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3653)
}
__typedArg0 := V3653
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13037 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13036)
}
__typedArg0 := tmp13036
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13038 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13037)
}
__typedArg0 := tmp13037
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13039 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13038)
}
__typedArg0 := tmp13038
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13040 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13039)
}
__typedArg0 := tmp13039
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13041 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp13040)
}
__typedArg0 := Nil
__typedArg1 := tmp13040
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13030 Obj

if True == tmp13041 {
tmp13032 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3653)
}
__typedArg0 := V3653
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13033)
}
__typedArg0 := tmp13033
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13035 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp13034)
}
__typedArg0 := Nil
__typedArg1 := tmp13034
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13031 Obj

if True == tmp13035 {
ifres13031 = True


} else {
ifres13031 = False


}

ifres13030 = ifres13031


} else {
ifres13030 = False


}

var ifres13029 Obj

if True == ifres13030 {
ifres13029 = True


} else {
ifres13029 = False


}

ifres13028 = ifres13029


} else {
ifres13028 = False


}

var ifres13027 Obj

if True == ifres13028 {
ifres13027 = True


} else {
ifres13027 = False


}

ifres13026 = ifres13027


} else {
ifres13026 = False


}

var ifres13025 Obj

if True == ifres13026 {
ifres13025 = True


} else {
ifres13025 = False


}

ifres13024 = ifres13025


} else {
ifres13024 = False


}

var ifres13023 Obj

if True == ifres13024 {
ifres13023 = True


} else {
ifres13023 = False


}

ifres13022 = ifres13023


} else {
ifres13022 = False


}

var ifres13021 Obj

if True == ifres13022 {
ifres13021 = True


} else {
ifres13021 = False


}

ifres13020 = ifres13021


} else {
ifres13020 = False


}

if True == ifres13020 {
tmp12997 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3653)
}
__typedArg0 := V3653
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12998 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp12997)
}
__typedArg0 := tmp12997
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12999 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp12998)
}
__typedArg0 := tmp12998
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13000 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp12999)
}
__typedArg0 := tmp12999
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13001 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13000)
}
__typedArg0 := tmp13000
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13002 := Call(__e, PrimFunc(symshen_4extract_1vars), tmp13001)


let__12279 := tmp13002
_ = let__12279

tmp13003 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3653)
}
__typedArg0 := V3653
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13004 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13003)
}
__typedArg0 := tmp13003
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
return PrimTail(tmp13005)
}
__typedArg0 := tmp13005
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13007 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13006)
}
__typedArg0 := tmp13006
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13008 := Call(__e, PrimFunc(symshen_4rule_1_6head), tmp13007)


tmp13009 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1_1, Nil)
}
__typedArg0 := sym_5_1_1
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13010 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3653)
}
__typedArg0 := V3653
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13011 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3653)
}
__typedArg0 := V3653
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13012 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13011)
}
__typedArg0 := tmp13011
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13013 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3653)
}
__typedArg0 := V3653
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13014 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13013)
}
__typedArg0 := tmp13013
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13015 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13014)
}
__typedArg0 := tmp13014
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13016 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13015)
}
__typedArg0 := tmp13015
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13017 := Call(__e, PrimFunc(symshen_4rule_1_6body), let__12279, symAssumptions, tmp13010, tmp13012, tmp13016)


tmp13018 := Call(__e, PrimFunc(symappend), tmp13009, tmp13017)


__e.TailApply(PrimFunc(symappend), tmp13008, tmp13018)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4rule_1_6clause)
return
}


}, 1)

tmp13057 := Call(__e, ns2_1set, symshen_4rule_1_6clause, tmp12996)


_ = tmp13057

tmp13058 := MakeNative(func(__e *ControlFlow) {
V3655 := __e.Get(1)
_ = V3655
tmp13059 := Call(__e, PrimFunc(symshen_4macro_1_8ch), V3655)


tmp13060 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symAssumptions, Nil)
}
__typedArg0 := symAssumptions
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13059, tmp13060)
}
__typedArg0 := tmp13059
__typedArg1 := tmp13060
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp13061 := Call(__e, ns2_1set, symshen_4rule_1_6head, tmp13058)


_ = tmp13061

tmp13062 := MakeNative(func(__e *ControlFlow) {
V3656 := __e.Get(1)
_ = V3656
tmp13063 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3656, Nil)
}
__typedArg0 := V3656
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4_8ch, tmp13063)
}
__typedArg0 := symshen_4_8ch
__typedArg1 := tmp13063
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp13064 := Call(__e, ns2_1set, symshen_4macro_1_8ch, tmp13062)


_ = tmp13064

tmp13065 := MakeNative(func(__e *ControlFlow) {
V3657 := __e.Get(1)
_ = V3657
tmp13066 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3657, Nil)
}
__typedArg0 := V3657
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4_8c, tmp13066)
}
__typedArg0 := symshen_4_8c
__typedArg1 := tmp13066
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp13067 := Call(__e, ns2_1set, symshen_4macro_1_8c, tmp13065)


_ = tmp13067

tmp13068 := MakeNative(func(__e *ControlFlow) {
V3658 := __e.Get(1)
_ = V3658
V3659 := __e.Get(2)
_ = V3659
V3660 := __e.Get(3)
_ = V3660
V3661 := __e.Get(4)
_ = V3661
V3662 := __e.Get(5)
_ = V3662
tmp13098 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3662)
}
__typedArg0 := Nil
__typedArg1 := V3662
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13098 {
__e.TailApply(PrimFunc(symshen_4side_1conditions_1_6goals), Nil, V3658, V3659, V3660, V3661)
return
} else {
tmp13096 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3661)
}
__typedArg0 := Nil
__typedArg1 := V3661
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13089 Obj

if True == tmp13096 {
tmp13095 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3662)
}
__typedArg0 := V3662
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13091 Obj

if True == tmp13095 {
tmp13093 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3662)
}
__typedArg0 := V3662
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13094 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp13093)
}
__typedArg0 := Nil
__typedArg1 := tmp13093
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13092 Obj

if True == tmp13094 {
ifres13092 = True


} else {
ifres13092 = False


}

ifres13091 = ifres13092


} else {
ifres13091 = False


}

var ifres13090 Obj

if True == ifres13091 {
ifres13090 = True


} else {
ifres13090 = False


}

ifres13089 = ifres13090


} else {
ifres13089 = False


}

if True == ifres13089 {
tmp13069 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3662)
}
__typedArg0 := V3662
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13070 := Call(__e, PrimFunc(symshen_4passive_1variables), tmp13069, V3658)


let__12280 := tmp13070
_ = let__12280

tmp13071 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3662)
}
__typedArg0 := V3662
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13072 := Call(__e, PrimFunc(symshen_4remove_1bystanders), V3658, tmp13071)


let__12281 := tmp13072
_ = let__12281

tmp13073 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3662)
}
__typedArg0 := V3662
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13074 := Call(__e, PrimFunc(symshen_4specialise_1member), tmp13073, V3659, let__12281, let__12280)


tmp13075 := Call(__e, PrimFunc(symshen_4side_1conditions_1_6goals), Nil, V3658, V3659, V3660, Nil)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13074, tmp13075)
}
__typedArg0 := tmp13074
__typedArg1 := tmp13075
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp13087 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3662)
}
__typedArg0 := V3662
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp13087 {
tmp13076 := Call(__e, PrimFunc(symgensym), symNewAssumptions)


let__12282 := tmp13076
_ = let__12282

tmp13077 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3662)
}
__typedArg0 := V3662
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13078 := Call(__e, PrimFunc(symshen_4passive_1variables), tmp13077, V3658)


let__12283 := tmp13078
_ = let__12283

tmp13079 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3662)
}
__typedArg0 := V3662
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13080 := Call(__e, PrimFunc(symshen_4remove_1bystanders), V3658, tmp13079)


let__12284 := tmp13080
_ = let__12284

tmp13081 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3662)
}
__typedArg0 := V3662
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13082 := Call(__e, PrimFunc(symshen_4specialise_1consume), tmp13081, V3659, let__12284, let__12283, let__12282)


tmp13083 := Call(__e, PrimFunc(symappend), V3658, let__12283)


tmp13084 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3662)
}
__typedArg0 := V3662
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13085 := Call(__e, PrimFunc(symshen_4rule_1_6body), tmp13083, let__12282, V3660, V3661, tmp13084)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13082, tmp13085)
}
__typedArg0 := tmp13082
__typedArg1 := tmp13085
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

tmp13099 := Call(__e, ns2_1set, symshen_4rule_1_6body, tmp13068)


_ = tmp13099

tmp13100 := MakeNative(func(__e *ControlFlow) {
V3668 := __e.Get(1)
_ = V3668
V3669 := __e.Get(2)
_ = V3669
V3670 := __e.Get(3)
_ = V3670
V3671 := __e.Get(4)
_ = V3671
tmp13101 := Call(__e, PrimFunc(symgensym), symshen_4member)


let__12285 := tmp13101
_ = let__12285

tmp13102 := Call(__e, PrimFunc(symshen_4member_1clause), let__12285, V3668, V3670, V3671)


let__12286 := tmp13102
_ = let__12286

tmp13103 := Call(__e, PrimFunc(symappend), V3670, V3671)


tmp13104 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3669, tmp13103)
}
__typedArg0 := V3669
__typedArg1 := tmp13103
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12285, tmp13104)
}
__typedArg0 := let__12285
__typedArg1 := tmp13104
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 4)

tmp13105 := Call(__e, ns2_1set, symshen_4specialise_1member, tmp13100)


_ = tmp13105

tmp13106 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__selftop:
V3676 := __self1
_ = V3676
V3677 := __self2
_ = V3677
tmp13120 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3676)
}
__typedArg0 := Nil
__typedArg1 := V3676
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13120 {
__e.Return(Nil)
return
} else {
tmp13118 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3676)
}
__typedArg0 := V3676
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13114 Obj

if True == tmp13118 {
tmp13116 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3676)
}
__typedArg0 := V3676
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13117 := Call(__e, PrimFunc(symshen_4occurs_1check_2), tmp13116, V3677)


var ifres13115 Obj

if True == tmp13117 {
ifres13115 = True


} else {
ifres13115 = False


}

ifres13114 = ifres13115


} else {
ifres13114 = False


}

if True == ifres13114 {
tmp13107 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3676)
}
__typedArg0 := V3676
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13108 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3676)
}
__typedArg0 := V3676
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13109 := Call(__e, PrimFunc(symshen_4remove_1bystanders), tmp13108, V3677)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13107, tmp13109)
}
__typedArg0 := tmp13107
__typedArg1 := tmp13109
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp13112 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3676)
}
__typedArg0 := V3676
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp13112 {
tmp13110 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3676)
}
__typedArg0 := V3676
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4remove_1bystanders) == __self {
__self1, __self2 = tmp13110, V3677
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4remove_1bystanders), tmp13110, V3677)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4remove_1bystanders)
return
}


}


}


}, 2)

tmp13121 := Call(__e, ns2_1set, symshen_4remove_1bystanders, tmp13106)


_ = tmp13121

tmp13122 := MakeNative(func(__e *ControlFlow) {
V3678 := __e.Get(1)
_ = V3678
V3679 := __e.Get(2)
_ = V3679
V3680 := __e.Get(3)
_ = V3680
V3681 := __e.Get(4)
_ = V3681
tmp13123 := Call(__e, PrimFunc(symlength), V3681)


tmp13124 := Call(__e, PrimFunc(symshen_4nvars), tmp13123)


let__12287 := tmp13124
_ = let__12287

tmp13125 := Call(__e, PrimFunc(symshen_4macro_1_8ch), V3679)


tmp13126 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym__, Nil)
}
__typedArg0 := sym__
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13127 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13125, tmp13126)
}
__typedArg0 := tmp13125
__typedArg1 := tmp13126
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13128 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp13127)
}
__typedArg0 := symcons
__typedArg1 := tmp13127
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13129 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13128, Nil)
}
__typedArg0 := tmp13128
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13130 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1, tmp13129)
}
__typedArg0 := sym_1
__typedArg1 := tmp13129
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13131 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13130, Nil)
}
__typedArg0 := tmp13130
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13132 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1_1, Nil)
}
__typedArg0 := sym_5_1_1
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13133 := Call(__e, PrimFunc(symshen_4passive_1bind), V3681, let__12287)


tmp13134 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(";"))
}
__typedArg0 := MakeString(";")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp13135 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13134, Nil)
}
__typedArg0 := tmp13134
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13136 := Call(__e, PrimFunc(symappend), tmp13133, tmp13135)


tmp13137 := Call(__e, PrimFunc(symappend), tmp13132, tmp13136)


tmp13138 := Call(__e, PrimFunc(symappend), let__12287, tmp13137)


tmp13139 := Call(__e, PrimFunc(symappend), V3680, tmp13138)


tmp13140 := Call(__e, PrimFunc(symappend), tmp13131, tmp13139)


let__12288 := tmp13140
_ = let__12288

tmp13141 := Call(__e, PrimFunc(symgensym), symHypotheses)


let__12290 := tmp13141
_ = let__12290

tmp13142 := Call(__e, PrimFunc(symappend), V3680, V3681)


let__12291 := tmp13142
_ = let__12291

tmp13143 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12290, Nil)
}
__typedArg0 := let__12290
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13144 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym__, tmp13143)
}
__typedArg0 := sym__
__typedArg1 := tmp13143
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13145 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp13144)
}
__typedArg0 := symcons
__typedArg1 := tmp13144
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13146 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13145, Nil)
}
__typedArg0 := tmp13145
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13147 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1, tmp13146)
}
__typedArg0 := sym_1
__typedArg1 := tmp13146
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13148 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13147, Nil)
}
__typedArg0 := tmp13147
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13149 := Call(__e, PrimFunc(symappend), tmp13148, let__12291)


let__12292 := tmp13149
_ = let__12292

tmp13150 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12290, let__12291)
}
__typedArg0 := let__12290
__typedArg1 := let__12291
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13151 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3678, tmp13150)
}
__typedArg0 := V3678
__typedArg1 := tmp13150
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13152 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13151, Nil)
}
__typedArg0 := tmp13151
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__12293 := tmp13152
_ = let__12293

tmp13153 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1_1, Nil)
}
__typedArg0 := sym_5_1_1
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13154 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(";"))
}
__typedArg0 := MakeString(";")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp13155 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13154, Nil)
}
__typedArg0 := tmp13154
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13156 := Call(__e, PrimFunc(symappend), let__12293, tmp13155)


tmp13157 := Call(__e, PrimFunc(symappend), tmp13153, tmp13156)


tmp13158 := Call(__e, PrimFunc(symappend), let__12292, tmp13157)


let__12289 := tmp13158
_ = let__12289

tmp13159 := Call(__e, PrimFunc(symappend), let__12288, let__12289)


tmp13160 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3678, tmp13159)
}
__typedArg0 := V3678
__typedArg1 := tmp13159
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13161 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefprolog, tmp13160)
}
__typedArg0 := symdefprolog
__typedArg1 := tmp13160
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__12294 := tmp13161
_ = let__12294

__e.TailApply(PrimFunc(symeval), let__12294)
return


}, 4)

tmp13162 := Call(__e, ns2_1set, symshen_4member_1clause, tmp13122)


_ = tmp13162

tmp13163 := MakeNative(func(__e *ControlFlow) {
V3690 := __e.Get(1)
_ = V3690
tmp13168 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(0), V3690)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := V3690
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13168 {
__e.Return(Nil)
return
} else {
tmp13164 := Call(__e, PrimFunc(symgensym), symNewV)


tmp13166 := Call(__e, PrimFunc(symshen_4nvars), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(V3690)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := V3690
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})())


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13164, tmp13166)
}
__typedArg0 := tmp13164
__typedArg1 := tmp13166
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}


}, 1)

tmp13169 := Call(__e, ns2_1set, symshen_4nvars, tmp13163)


_ = tmp13169

tmp13170 := MakeNative(func(__e *ControlFlow) {
V3691 := __e.Get(1)
_ = V3691
V3692 := __e.Get(2)
_ = V3692
tmp13188 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3691)
}
__typedArg0 := Nil
__typedArg1 := V3691
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13185 Obj

if True == tmp13188 {
tmp13187 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3692)
}
__typedArg0 := Nil
__typedArg1 := V3692
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13186 Obj

if True == tmp13187 {
ifres13186 = True


} else {
ifres13186 = False


}

ifres13185 = ifres13186


} else {
ifres13185 = False


}

if True == ifres13185 {
__e.Return(Nil)
return
} else {
tmp13183 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3691)
}
__typedArg0 := V3691
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13180 Obj

if True == tmp13183 {
tmp13182 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3692)
}
__typedArg0 := V3692
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13181 Obj

if True == tmp13182 {
ifres13181 = True


} else {
ifres13181 = False


}

ifres13180 = ifres13181


} else {
ifres13180 = False


}

if True == ifres13180 {
tmp13171 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3692)
}
__typedArg0 := V3692
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13172 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3691)
}
__typedArg0 := V3691
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13173 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13172, Nil)
}
__typedArg0 := tmp13172
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13174 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13171, tmp13173)
}
__typedArg0 := tmp13171
__typedArg1 := tmp13173
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13175 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symbind, tmp13174)
}
__typedArg0 := symbind
__typedArg1 := tmp13174
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13176 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3691)
}
__typedArg0 := V3691
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13177 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3692)
}
__typedArg0 := V3692
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13178 := Call(__e, PrimFunc(symshen_4passive_1bind), tmp13176, tmp13177)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13175, tmp13178)
}
__typedArg0 := tmp13175
__typedArg1 := tmp13178
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4passive_1bind)
return
}


}


}, 2)

tmp13189 := Call(__e, ns2_1set, symshen_4passive_1bind, tmp13170)


_ = tmp13189

tmp13190 := MakeNative(func(__e *ControlFlow) {
V3693 := __e.Get(1)
_ = V3693
V3694 := __e.Get(2)
_ = V3694
V3695 := __e.Get(3)
_ = V3695
V3696 := __e.Get(4)
_ = V3696
V3697 := __e.Get(5)
_ = V3697
tmp13191 := Call(__e, PrimFunc(symgensym), symshen_4consume)


let__12295 := tmp13191
_ = let__12295

tmp13192 := Call(__e, PrimFunc(symshen_4consume_1clause), let__12295, V3693, V3695, V3696, V3697)


let__12296 := tmp13192
_ = let__12296

tmp13193 := Call(__e, PrimFunc(symappend), V3695, V3696)


tmp13194 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3697, tmp13193)
}
__typedArg0 := V3697
__typedArg1 := tmp13193
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13195 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3694, tmp13194)
}
__typedArg0 := V3694
__typedArg1 := tmp13194
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12295, tmp13195)
}
__typedArg0 := let__12295
__typedArg1 := tmp13195
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 5)

tmp13196 := Call(__e, ns2_1set, symshen_4specialise_1consume, tmp13190)


_ = tmp13196

tmp13197 := MakeNative(func(__e *ControlFlow) {
V3700 := __e.Get(1)
_ = V3700
V3701 := __e.Get(2)
_ = V3701
V3702 := __e.Get(3)
_ = V3702
V3703 := __e.Get(4)
_ = V3703
V3704 := __e.Get(5)
_ = V3704
tmp13198 := Call(__e, PrimFunc(symlength), V3703)


tmp13199 := Call(__e, PrimFunc(symshen_4nvars), tmp13198)


let__12297 := tmp13199
_ = let__12297

tmp13200 := Call(__e, PrimFunc(symgensym), symAssumption)


let__12298 := tmp13200
_ = let__12298

tmp13201 := Call(__e, PrimFunc(symshen_4macro_1_8ch), V3701)


tmp13202 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12298, Nil)
}
__typedArg0 := let__12298
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13203 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13201, tmp13202)
}
__typedArg0 := tmp13201
__typedArg1 := tmp13202
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13204 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp13203)
}
__typedArg0 := symcons
__typedArg1 := tmp13203
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13205 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13204, Nil)
}
__typedArg0 := tmp13204
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13206 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1, tmp13205)
}
__typedArg0 := sym_1
__typedArg1 := tmp13205
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13207 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1_1, Nil)
}
__typedArg0 := sym_5_1_1
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13208 := Call(__e, PrimFunc(symshen_4passive_1bind), V3703, let__12297)


tmp13209 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12298, Nil)
}
__typedArg0 := let__12298
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13210 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3704, tmp13209)
}
__typedArg0 := V3704
__typedArg1 := tmp13209
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13211 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symbind, tmp13210)
}
__typedArg0 := symbind
__typedArg1 := tmp13210
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13212 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(";"))
}
__typedArg0 := MakeString(";")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp13213 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13212, Nil)
}
__typedArg0 := tmp13212
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13214 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13211, tmp13213)
}
__typedArg0 := tmp13211
__typedArg1 := tmp13213
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13215 := Call(__e, PrimFunc(symappend), tmp13208, tmp13214)


tmp13216 := Call(__e, PrimFunc(symappend), tmp13207, tmp13215)


tmp13217 := Call(__e, PrimFunc(symappend), let__12297, tmp13216)


tmp13218 := Call(__e, PrimFunc(symappend), V3702, tmp13217)


tmp13219 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3704, tmp13218)
}
__typedArg0 := V3704
__typedArg1 := tmp13218
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13220 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13206, tmp13219)
}
__typedArg0 := tmp13206
__typedArg1 := tmp13219
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__12299 := tmp13220
_ = let__12299

tmp13221 := Call(__e, PrimFunc(symgensym), symHypotheses)


let__12301 := tmp13221
_ = let__12301

tmp13222 := Call(__e, PrimFunc(symappend), V3702, V3703)


let__12302 := tmp13222
_ = let__12302

tmp13223 := Call(__e, PrimFunc(symgensym), symAssumptions)


let__12303 := tmp13223
_ = let__12303

tmp13224 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12301, Nil)
}
__typedArg0 := let__12301
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13225 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12298, tmp13224)
}
__typedArg0 := let__12298
__typedArg1 := tmp13224
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13226 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp13225)
}
__typedArg0 := symcons
__typedArg1 := tmp13225
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13227 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13226, Nil)
}
__typedArg0 := tmp13226
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13228 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1, tmp13227)
}
__typedArg0 := sym_1
__typedArg1 := tmp13227
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13229 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3704, Nil)
}
__typedArg0 := V3704
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13230 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12303, tmp13229)
}
__typedArg0 := let__12303
__typedArg1 := tmp13229
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13231 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp13230)
}
__typedArg0 := symcons
__typedArg1 := tmp13230
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13232 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13231, let__12302)
}
__typedArg0 := tmp13231
__typedArg1 := let__12302
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13233 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13228, tmp13232)
}
__typedArg0 := tmp13228
__typedArg1 := tmp13232
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__12304 := tmp13233
_ = let__12304

tmp13234 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12298, Nil)
}
__typedArg0 := let__12298
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13235 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12303, tmp13234)
}
__typedArg0 := let__12303
__typedArg1 := tmp13234
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13236 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symbind, tmp13235)
}
__typedArg0 := symbind
__typedArg1 := tmp13235
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13237 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3704, let__12302)
}
__typedArg0 := V3704
__typedArg1 := let__12302
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13238 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12301, tmp13237)
}
__typedArg0 := let__12301
__typedArg1 := tmp13237
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13239 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3700, tmp13238)
}
__typedArg0 := V3700
__typedArg1 := tmp13238
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13240 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13239, Nil)
}
__typedArg0 := tmp13239
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13241 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13236, tmp13240)
}
__typedArg0 := tmp13236
__typedArg1 := tmp13240
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__12305 := tmp13241
_ = let__12305

tmp13242 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1_1, Nil)
}
__typedArg0 := sym_5_1_1
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13243 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(";"))
}
__typedArg0 := MakeString(";")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp13244 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13243, Nil)
}
__typedArg0 := tmp13243
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13245 := Call(__e, PrimFunc(symappend), let__12305, tmp13244)


tmp13246 := Call(__e, PrimFunc(symappend), tmp13242, tmp13245)


tmp13247 := Call(__e, PrimFunc(symappend), let__12304, tmp13246)


let__12300 := tmp13247
_ = let__12300

tmp13248 := Call(__e, PrimFunc(symappend), let__12299, let__12300)


tmp13249 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3700, tmp13248)
}
__typedArg0 := V3700
__typedArg1 := tmp13248
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13250 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefprolog, tmp13249)
}
__typedArg0 := symdefprolog
__typedArg1 := tmp13249
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__12306 := tmp13250
_ = let__12306

__e.TailApply(PrimFunc(symeval), let__12306)
return


}, 5)

tmp13251 := Call(__e, ns2_1set, symshen_4consume_1clause, tmp13197)


_ = tmp13251

tmp13252 := MakeNative(func(__e *ControlFlow) {
V3715 := __e.Get(1)
_ = V3715
V3716 := __e.Get(2)
_ = V3716
tmp13253 := Call(__e, PrimFunc(symshen_4extract_1vars), V3715)


__e.TailApply(PrimFunc(symdifference), tmp13253, V3716)
return


}, 2)

tmp13254 := Call(__e, ns2_1set, symshen_4passive_1variables, tmp13252)


_ = tmp13254

tmp13255 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__self3 := __e.Get(3)
__self4 := __e.Get(4)
__self5 := __e.Get(5)
__selftop:
V3721 := __self1
_ = V3721
V3722 := __self2
_ = V3722
V3723 := __self3
_ = V3723
V3724 := __self4
_ = V3724
V3725 := __self5
_ = V3725
tmp13406 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3724)
}
__typedArg0 := Nil
__typedArg1 := V3724
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13406 {
__e.TailApply(PrimFunc(symshen_4premises_1_6goals), V3721, V3723, V3725)
return
} else {
tmp13404 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3724)
}
__typedArg0 := V3724
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13384 Obj

if True == tmp13404 {
tmp13402 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3724)
}
__typedArg0 := V3724
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13403 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13402)
}
__typedArg0 := tmp13402
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13386 Obj

if True == tmp13403 {
tmp13399 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3724)
}
__typedArg0 := V3724
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13400 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13399)
}
__typedArg0 := tmp13399
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13401 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symif, tmp13400)
}
__typedArg0 := symif
__typedArg1 := tmp13400
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13388 Obj

if True == tmp13401 {
tmp13396 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3724)
}
__typedArg0 := V3724
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13397 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13396)
}
__typedArg0 := tmp13396
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13398 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13397)
}
__typedArg0 := tmp13397
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13390 Obj

if True == tmp13398 {
tmp13392 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3724)
}
__typedArg0 := V3724
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13393 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13392)
}
__typedArg0 := tmp13392
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13394 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13393)
}
__typedArg0 := tmp13393
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13395 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp13394)
}
__typedArg0 := Nil
__typedArg1 := tmp13394
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13391 Obj

if True == tmp13395 {
ifres13391 = True


} else {
ifres13391 = False


}

ifres13390 = ifres13391


} else {
ifres13390 = False


}

var ifres13389 Obj

if True == ifres13390 {
ifres13389 = True


} else {
ifres13389 = False


}

ifres13388 = ifres13389


} else {
ifres13388 = False


}

var ifres13387 Obj

if True == ifres13388 {
ifres13387 = True


} else {
ifres13387 = False


}

ifres13386 = ifres13387


} else {
ifres13386 = False


}

var ifres13385 Obj

if True == ifres13386 {
ifres13385 = True


} else {
ifres13385 = False


}

ifres13384 = ifres13385


} else {
ifres13384 = False


}

if True == ifres13384 {
tmp13256 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3724)
}
__typedArg0 := V3724
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13257 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13256)
}
__typedArg0 := tmp13256
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13258 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhen, tmp13257)
}
__typedArg0 := symwhen
__typedArg1 := tmp13257
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13259 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3724)
}
__typedArg0 := V3724
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13260 := Call(__e, PrimFunc(symshen_4side_1conditions_1_6goals), V3721, V3722, V3723, tmp13259, V3725)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13258, tmp13260)
}
__typedArg0 := tmp13258
__typedArg1 := tmp13260
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp13382 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3724)
}
__typedArg0 := V3724
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13355 Obj

if True == tmp13382 {
tmp13380 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3724)
}
__typedArg0 := V3724
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13381 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13380)
}
__typedArg0 := tmp13380
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13357 Obj

if True == tmp13381 {
tmp13377 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3724)
}
__typedArg0 := V3724
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13378 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13377)
}
__typedArg0 := tmp13377
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13379 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlet, tmp13378)
}
__typedArg0 := symlet
__typedArg1 := tmp13378
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13359 Obj

if True == tmp13379 {
tmp13374 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3724)
}
__typedArg0 := V3724
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13375 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13374)
}
__typedArg0 := tmp13374
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13376 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13375)
}
__typedArg0 := tmp13375
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13361 Obj

if True == tmp13376 {
tmp13370 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3724)
}
__typedArg0 := V3724
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13371 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13370)
}
__typedArg0 := tmp13370
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13372 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13371)
}
__typedArg0 := tmp13371
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13373 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13372)
}
__typedArg0 := tmp13372
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13363 Obj

if True == tmp13373 {
tmp13365 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3724)
}
__typedArg0 := V3724
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13366 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13365)
}
__typedArg0 := tmp13365
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13367 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13366)
}
__typedArg0 := tmp13366
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13368 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13367)
}
__typedArg0 := tmp13367
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13369 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp13368)
}
__typedArg0 := Nil
__typedArg1 := tmp13368
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13364 Obj

if True == tmp13369 {
ifres13364 = True


} else {
ifres13364 = False


}

ifres13363 = ifres13364


} else {
ifres13363 = False


}

var ifres13362 Obj

if True == ifres13363 {
ifres13362 = True


} else {
ifres13362 = False


}

ifres13361 = ifres13362


} else {
ifres13361 = False


}

var ifres13360 Obj

if True == ifres13361 {
ifres13360 = True


} else {
ifres13360 = False


}

ifres13359 = ifres13360


} else {
ifres13359 = False


}

var ifres13358 Obj

if True == ifres13359 {
ifres13358 = True


} else {
ifres13358 = False


}

ifres13357 = ifres13358


} else {
ifres13357 = False


}

var ifres13356 Obj

if True == ifres13357 {
ifres13356 = True


} else {
ifres13356 = False


}

ifres13355 = ifres13356


} else {
ifres13355 = False


}

if True == ifres13355 {
tmp13276 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3724)
}
__typedArg0 := V3724
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13277 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13276)
}
__typedArg0 := tmp13276
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13278 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13277)
}
__typedArg0 := tmp13277
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13279 := Call(__e, PrimFunc(symelement_2), tmp13278, V3722)


if True == tmp13279 {
tmp13261 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3724)
}
__typedArg0 := V3724
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13262 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13261)
}
__typedArg0 := tmp13261
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13263 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symis_b, tmp13262)
}
__typedArg0 := symis_b
__typedArg1 := tmp13262
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13264 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3724)
}
__typedArg0 := V3724
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13265 := Call(__e, PrimFunc(symshen_4side_1conditions_1_6goals), V3721, V3722, V3723, tmp13264, V3725)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13263, tmp13265)
}
__typedArg0 := tmp13263
__typedArg1 := tmp13265
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp13266 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3724)
}
__typedArg0 := V3724
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13267 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13266)
}
__typedArg0 := tmp13266
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13268 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symbind, tmp13267)
}
__typedArg0 := symbind
__typedArg1 := tmp13267
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13269 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3724)
}
__typedArg0 := V3724
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

tmp13272 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13271, V3722)
}
__typedArg0 := tmp13271
__typedArg1 := V3722
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13273 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3724)
}
__typedArg0 := V3724
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13274 := Call(__e, PrimFunc(symshen_4side_1conditions_1_6goals), V3721, tmp13272, V3723, tmp13273, V3725)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13268, tmp13274)
}
__typedArg0 := tmp13268
__typedArg1 := tmp13274
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}


} else {
tmp13353 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3724)
}
__typedArg0 := V3724
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13333 Obj

if True == tmp13353 {
tmp13351 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3724)
}
__typedArg0 := V3724
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13352 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13351)
}
__typedArg0 := tmp13351
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13335 Obj

if True == tmp13352 {
tmp13348 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3724)
}
__typedArg0 := V3724
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13349 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13348)
}
__typedArg0 := tmp13348
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13350 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symctxt, tmp13349)
}
__typedArg0 := symctxt
__typedArg1 := tmp13349
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13337 Obj

if True == tmp13350 {
tmp13345 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3724)
}
__typedArg0 := V3724
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13346 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13345)
}
__typedArg0 := tmp13345
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13347 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13346)
}
__typedArg0 := tmp13346
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13339 Obj

if True == tmp13347 {
tmp13341 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3724)
}
__typedArg0 := V3724
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13342 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13341)
}
__typedArg0 := tmp13341
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13343 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13342)
}
__typedArg0 := tmp13342
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13344 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp13343)
}
__typedArg0 := Nil
__typedArg1 := tmp13343
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13340 Obj

if True == tmp13344 {
ifres13340 = True


} else {
ifres13340 = False


}

ifres13339 = ifres13340


} else {
ifres13339 = False


}

var ifres13338 Obj

if True == ifres13339 {
ifres13338 = True


} else {
ifres13338 = False


}

ifres13337 = ifres13338


} else {
ifres13337 = False


}

var ifres13336 Obj

if True == ifres13337 {
ifres13336 = True


} else {
ifres13336 = False


}

ifres13335 = ifres13336


} else {
ifres13335 = False


}

var ifres13334 Obj

if True == ifres13335 {
ifres13334 = True


} else {
ifres13334 = False


}

ifres13333 = ifres13334


} else {
ifres13333 = False


}

if True == ifres13333 {
tmp13305 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3724)
}
__typedArg0 := V3724
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13306 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13305)
}
__typedArg0 := tmp13305
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13307 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13306)
}
__typedArg0 := tmp13306
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13308 := Call(__e, PrimFunc(symelement_2), tmp13307, V3722)


if True == tmp13308 {
tmp13280 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3724)
}
__typedArg0 := V3724
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13281 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13280)
}
__typedArg0 := tmp13280
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13282 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13281)
}
__typedArg0 := tmp13281
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13283 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13282, V3721)
}
__typedArg0 := tmp13282
__typedArg1 := V3721
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13284 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3724)
}
__typedArg0 := V3724
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4side_1conditions_1_6goals) == __self {
__self1, __self2, __self3, __self4, __self5 = tmp13283, V3722, V3723, tmp13284, V3725
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4side_1conditions_1_6goals), tmp13283, V3722, V3723, tmp13284, V3725)
return


} else {
tmp13285 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3724)
}
__typedArg0 := V3724
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13286 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13285)
}
__typedArg0 := tmp13285
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13287 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13286)
}
__typedArg0 := tmp13286
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13288 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3723, Nil)
}
__typedArg0 := V3723
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13289 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13287, tmp13288)
}
__typedArg0 := tmp13287
__typedArg1 := tmp13288
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13290 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symbind, tmp13289)
}
__typedArg0 := symbind
__typedArg1 := tmp13289
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13291 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3724)
}
__typedArg0 := V3724
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13292 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13291)
}
__typedArg0 := tmp13291
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13293 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13292)
}
__typedArg0 := tmp13292
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13294 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13293, V3721)
}
__typedArg0 := tmp13293
__typedArg1 := V3721
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13295 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3724)
}
__typedArg0 := V3724
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13296 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13295)
}
__typedArg0 := tmp13295
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13297 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13296)
}
__typedArg0 := tmp13296
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13298 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13297, V3722)
}
__typedArg0 := tmp13297
__typedArg1 := V3722
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13299 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3724)
}
__typedArg0 := V3724
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13300 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13299)
}
__typedArg0 := tmp13299
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13301 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13300)
}
__typedArg0 := tmp13300
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13302 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3724)
}
__typedArg0 := V3724
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13303 := Call(__e, PrimFunc(symshen_4side_1conditions_1_6goals), tmp13294, tmp13298, tmp13301, tmp13302, V3725)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13290, tmp13303)
}
__typedArg0 := tmp13290
__typedArg1 := tmp13303
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}


} else {
tmp13331 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3724)
}
__typedArg0 := V3724
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13311 Obj

if True == tmp13331 {
tmp13329 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3724)
}
__typedArg0 := V3724
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13330 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13329)
}
__typedArg0 := tmp13329
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13313 Obj

if True == tmp13330 {
tmp13326 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3724)
}
__typedArg0 := V3724
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13327 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13326)
}
__typedArg0 := tmp13326
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13328 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symsqts, tmp13327)
}
__typedArg0 := symsqts
__typedArg1 := tmp13327
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13315 Obj

if True == tmp13328 {
tmp13323 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3724)
}
__typedArg0 := V3724
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13324 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13323)
}
__typedArg0 := tmp13323
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13325 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13324)
}
__typedArg0 := tmp13324
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13317 Obj

if True == tmp13325 {
tmp13319 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3724)
}
__typedArg0 := V3724
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13320 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13319)
}
__typedArg0 := tmp13319
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13321 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13320)
}
__typedArg0 := tmp13320
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13322 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp13321)
}
__typedArg0 := Nil
__typedArg1 := tmp13321
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13318 Obj

if True == tmp13322 {
ifres13318 = True


} else {
ifres13318 = False


}

ifres13317 = ifres13318


} else {
ifres13317 = False


}

var ifres13316 Obj

if True == ifres13317 {
ifres13316 = True


} else {
ifres13316 = False


}

ifres13315 = ifres13316


} else {
ifres13315 = False


}

var ifres13314 Obj

if True == ifres13315 {
ifres13314 = True


} else {
ifres13314 = False


}

ifres13313 = ifres13314


} else {
ifres13313 = False


}

var ifres13312 Obj

if True == ifres13313 {
ifres13312 = True


} else {
ifres13312 = False


}

ifres13311 = ifres13312


} else {
ifres13311 = False


}

if True == ifres13311 {
tmp13309 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3724)
}
__typedArg0 := V3724
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4side_1conditions_1_6goals) == __self {
__self1, __self2, __self3, __self4, __self5 = V3721, V3722, V3723, tmp13309, V3725
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4side_1conditions_1_6goals), V3721, V3722, V3723, tmp13309, V3725)
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

tmp13407 := Call(__e, ns2_1set, symshen_4side_1conditions_1_6goals, tmp13255)


_ = tmp13407

tmp13408 := MakeNative(func(__e *ControlFlow) {
V3730 := __e.Get(1)
_ = V3730
V3731 := __e.Get(2)
_ = V3731
V3732 := __e.Get(3)
_ = V3732
tmp13458 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3732)
}
__typedArg0 := Nil
__typedArg1 := V3732
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13458 {
tmp13409 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(";"))
}
__typedArg0 := MakeString(";")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13409, Nil)
}
__typedArg0 := tmp13409
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp13456 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3732)
}
__typedArg0 := V3732
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13452 Obj

if True == tmp13456 {
tmp13454 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3732)
}
__typedArg0 := V3732
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13455 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_b, tmp13454)
}
__typedArg0 := sym_b
__typedArg1 := tmp13454
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13453 Obj

if True == tmp13455 {
ifres13453 = True


} else {
ifres13453 = False


}

ifres13452 = ifres13453


} else {
ifres13452 = False


}

if True == ifres13452 {
tmp13410 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3732)
}
__typedArg0 := V3732
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13411 := Call(__e, PrimFunc(symshen_4premises_1_6goals), V3730, V3731, tmp13410)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_b, tmp13411)
}
__typedArg0 := sym_b
__typedArg1 := tmp13411
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp13450 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3732)
}
__typedArg0 := V3732
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13446 Obj

if True == tmp13450 {
tmp13448 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3732)
}
__typedArg0 := V3732
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13449 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symfail, tmp13448)
}
__typedArg0 := symfail
__typedArg1 := tmp13448
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13447 Obj

if True == tmp13449 {
ifres13447 = True


} else {
ifres13447 = False


}

ifres13446 = ifres13447


} else {
ifres13446 = False


}

if True == ifres13446 {
tmp13412 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(False, Nil)
}
__typedArg0 := False
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13413 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhen, tmp13412)
}
__typedArg0 := symwhen
__typedArg1 := tmp13412
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13414 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3732)
}
__typedArg0 := V3732
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13415 := Call(__e, PrimFunc(symshen_4premises_1_6goals), V3730, V3731, tmp13414)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13413, tmp13415)
}
__typedArg0 := tmp13413
__typedArg1 := tmp13415
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp13444 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3732)
}
__typedArg0 := V3732
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13429 Obj

if True == tmp13444 {
tmp13442 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3732)
}
__typedArg0 := V3732
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13443 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13442)
}
__typedArg0 := tmp13442
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13431 Obj

if True == tmp13443 {
tmp13439 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3732)
}
__typedArg0 := V3732
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13440 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13439)
}
__typedArg0 := tmp13439
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13441 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13440)
}
__typedArg0 := tmp13440
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13433 Obj

if True == tmp13441 {
tmp13435 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3732)
}
__typedArg0 := V3732
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13436 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13435)
}
__typedArg0 := tmp13435
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13437 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13436)
}
__typedArg0 := tmp13436
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13438 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp13437)
}
__typedArg0 := Nil
__typedArg1 := tmp13437
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13434 Obj

if True == tmp13438 {
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

var ifres13430 Obj

if True == ifres13431 {
ifres13430 = True


} else {
ifres13430 = False


}

ifres13429 = ifres13430


} else {
ifres13429 = False


}

if True == ifres13429 {
tmp13416 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3732)
}
__typedArg0 := V3732
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13417 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13416)
}
__typedArg0 := tmp13416
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13418 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13417)
}
__typedArg0 := tmp13417
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13419 := Call(__e, PrimFunc(symshen_4macro_1_8c), tmp13418)


tmp13420 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3732)
}
__typedArg0 := V3732
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13421 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13420)
}
__typedArg0 := tmp13420
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13422 := Call(__e, PrimFunc(symshen_4construct_1context), V3730, tmp13421, V3731)


tmp13423 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13422, Nil)
}
__typedArg0 := tmp13422
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13424 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13419, tmp13423)
}
__typedArg0 := tmp13419
__typedArg1 := tmp13423
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13425 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4system_1S, tmp13424)
}
__typedArg0 := symshen_4system_1S
__typedArg1 := tmp13424
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13426 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3732)
}
__typedArg0 := V3732
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13427 := Call(__e, PrimFunc(symshen_4premises_1_6goals), V3730, V3731, tmp13426)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13425, tmp13427)
}
__typedArg0 := tmp13425
__typedArg1 := tmp13427
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

tmp13459 := Call(__e, ns2_1set, symshen_4premises_1_6goals, tmp13408)


_ = tmp13459

tmp13460 := MakeNative(func(__e *ControlFlow) {
V3736 := __e.Get(1)
_ = V3736
V3737 := __e.Get(2)
_ = V3737
V3738 := __e.Get(3)
_ = V3738
tmp13480 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3737)
}
__typedArg0 := Nil
__typedArg1 := V3737
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13480 {
__e.Return(V3738)
return
} else {
tmp13478 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3737)
}
__typedArg0 := V3737
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13470 Obj

if True == tmp13478 {
tmp13476 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3737)
}
__typedArg0 := V3737
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13477 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp13476)
}
__typedArg0 := Nil
__typedArg1 := tmp13476
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13472 Obj

if True == tmp13477 {
tmp13474 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3737)
}
__typedArg0 := V3737
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13475 := Call(__e, PrimFunc(symelement_2), tmp13474, V3736)


var ifres13473 Obj

if True == tmp13475 {
ifres13473 = True


} else {
ifres13473 = False


}

ifres13472 = ifres13473


} else {
ifres13472 = False


}

var ifres13471 Obj

if True == ifres13472 {
ifres13471 = True


} else {
ifres13471 = False


}

ifres13470 = ifres13471


} else {
ifres13470 = False


}

if True == ifres13470 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3737)
}
__typedArg0 := V3737
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return
} else {
tmp13468 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3737)
}
__typedArg0 := V3737
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp13468 {
tmp13461 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3737)
}
__typedArg0 := V3737
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13462 := Call(__e, PrimFunc(symshen_4macro_1_8c), tmp13461)


tmp13463 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3737)
}
__typedArg0 := V3737
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13464 := Call(__e, PrimFunc(symshen_4construct_1context), V3736, tmp13463, V3738)


tmp13465 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13464, Nil)
}
__typedArg0 := tmp13464
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13466 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13462, tmp13465)
}
__typedArg0 := tmp13462
__typedArg1 := tmp13465
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp13466)
}
__typedArg0 := symcons
__typedArg1 := tmp13466
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

tmp13481 := Call(__e, ns2_1set, symshen_4construct_1context, tmp13460)


_ = tmp13481

tmp13482 := MakeNative(func(__e *ControlFlow) {
V3739 := __e.Get(1)
_ = V3739
tmp13483 := MakeNative(func(__e *ControlFlow) {
Z3741 := __e.Get(1)
_ = Z3741
__e.TailApply(PrimFunc(symshen_4intern_1type), Z3741)
return
}, 1)

tmp13484 := Call(__e, PrimFunc(symmap), tmp13483, V3739)


let__12307 := tmp13484
_ = let__12307

tmp13485 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_ddatatypes_d)
}
__typedArg0 := symshen_4_ddatatypes_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

let__12308 := tmp13485
_ = let__12308

tmp13486 := Call(__e, PrimFunc(symshen_4remove_1datatypes), let__12307, let__12308)


let__12309 := tmp13486
_ = let__12309

tmp13487 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_ddatatypes_d, let__12309)
}
__typedArg0 := symshen_4_ddatatypes_d
__typedArg1 := let__12309
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

let__12310 := tmp13487
_ = let__12310

__e.TailApply(PrimFunc(symshen_4show_1datatypes), let__12310)
return


}, 1)

tmp13488 := Call(__e, ns2_1set, sympreclude, tmp13482)


_ = tmp13488

tmp13489 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__selftop:
V3749 := __self1
_ = V3749
V3750 := __self2
_ = V3750
tmp13496 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3749)
}
__typedArg0 := Nil
__typedArg1 := V3749
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13496 {
__e.Return(V3750)
return
} else {
tmp13494 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3749)
}
__typedArg0 := V3749
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp13494 {
tmp13490 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3749)
}
__typedArg0 := V3749
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13491 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3749)
}
__typedArg0 := V3749
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13492 := Call(__e, PrimFunc(symshen_4unassoc), tmp13491, V3750)


if PrimFunc(symshen_4remove_1datatypes) == __self {
__self1, __self2 = tmp13490, tmp13492
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4remove_1datatypes), tmp13490, tmp13492)
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

tmp13497 := Call(__e, ns2_1set, symshen_4remove_1datatypes, tmp13489)


_ = tmp13497

tmp13498 := MakeNative(func(__e *ControlFlow) {
V3751 := __e.Get(1)
_ = V3751
tmp13499 := MakeNative(func(__e *ControlFlow) {
Z3752 := __e.Get(1)
_ = Z3752
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(Z3752)
}
__typedArg0 := Z3752
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp13499, V3751)
return


}, 1)

tmp13500 := Call(__e, ns2_1set, symshen_4show_1datatypes, tmp13498)


_ = tmp13500

tmp13501 := MakeNative(func(__e *ControlFlow) {
V3753 := __e.Get(1)
_ = V3753
tmp13502 := MakeNative(func(__e *ControlFlow) {
Z3755 := __e.Get(1)
_ = Z3755
__e.TailApply(PrimFunc(symshen_4intern_1type), Z3755)
return
}, 1)

tmp13503 := Call(__e, PrimFunc(symmap), tmp13502, V3753)


let__12311 := tmp13503
_ = let__12311

tmp13504 := MakeNative(func(__e *ControlFlow) {
Z3757 := __e.Get(1)
_ = Z3757
tmp13505 := Call(__e, PrimFunc(symfn), Z3757)


__e.TailApply(PrimFunc(symshen_4remember_1datatype), Z3757, tmp13505)
return


}, 1)

tmp13506 := Call(__e, PrimFunc(symmap), tmp13504, let__12311)


let__12312 := tmp13506
_ = let__12312

tmp13507 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_ddatatypes_d)
}
__typedArg0 := symshen_4_ddatatypes_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

let__12313 := tmp13507
_ = let__12313

__e.TailApply(PrimFunc(symshen_4show_1datatypes), let__12313)
return


}, 1)

tmp13508 := Call(__e, ns2_1set, syminclude, tmp13501)


_ = tmp13508

tmp13509 := MakeNative(func(__e *ControlFlow) {
V3759 := __e.Get(1)
_ = V3759
tmp13510 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_ddatatypes_d, Nil)
}
__typedArg0 := symshen_4_ddatatypes_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

let__12314 := tmp13510
_ = let__12314

tmp13511 := MakeNative(func(__e *ControlFlow) {
Z3762 := __e.Get(1)
_ = Z3762
__e.TailApply(PrimFunc(symshen_4intern_1type), Z3762)
return
}, 1)

tmp13512 := Call(__e, PrimFunc(symmap), tmp13511, V3759)


let__12315 := tmp13512
_ = let__12315

tmp13513 := MakeNative(func(__e *ControlFlow) {
Z3764 := __e.Get(1)
_ = Z3764
tmp13514 := Call(__e, PrimFunc(symfn), Z3764)


__e.TailApply(PrimFunc(symshen_4remember_1datatype), Z3764, tmp13514)
return


}, 1)

tmp13515 := Call(__e, PrimFunc(symmap), tmp13513, let__12315)


let__12316 := tmp13515
_ = let__12316

tmp13516 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_ddatatypes_d)
}
__typedArg0 := symshen_4_ddatatypes_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4show_1datatypes), tmp13516)
return


}, 1)

tmp13517 := Call(__e, ns2_1set, sympreclude_1all_1but, tmp13509)


_ = tmp13517

tmp13518 := MakeNative(func(__e *ControlFlow) {
V3765 := __e.Get(1)
_ = V3765
tmp13519 := MakeNative(func(__e *ControlFlow) {
Z3767 := __e.Get(1)
_ = Z3767
__e.TailApply(PrimFunc(symshen_4intern_1type), Z3767)
return
}, 1)

tmp13520 := Call(__e, PrimFunc(symmap), tmp13519, V3765)


let__12317 := tmp13520
_ = let__12317

tmp13521 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dalldatatypes_d)
}
__typedArg0 := symshen_4_dalldatatypes_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

let__12318 := tmp13521
_ = let__12318

tmp13522 := Call(__e, PrimFunc(symshen_4remove_1datatypes), let__12317, let__12318)


tmp13523 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_ddatatypes_d, tmp13522)
}
__typedArg0 := symshen_4_ddatatypes_d
__typedArg1 := tmp13522
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

let__12319 := tmp13523
_ = let__12319

__e.TailApply(PrimFunc(symshen_4show_1datatypes), let__12319)
return


}, 1)

__e.TailApply(ns2_1set, syminclude_1all_1but, tmp13518)
return




}, 0)

