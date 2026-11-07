package main

import . "github.com/pyrex41/shen-go/kl"

var SequentMain = MakeNative(func(__e *ControlFlow) {
tmp12340 := MakeNative(func(__e *ControlFlow) {
V3418 := __e.Get(1)
_ = V3418
tmp12355 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3418)
}
__typedArg0 := V3418
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12341 Obj

if True == tmp12355 {
tmp12342 := Call(__e, PrimFunc(symhead), V3418)


let__12110 := tmp12342
_ = let__12110

tmp12343 := Call(__e, PrimFunc(symtail), V3418)


let__12111 := tmp12343
_ = let__12111

tmp12344 := Call(__e, PrimFunc(symshen_4_5datatype_1rules_6), let__12111)


let__12112 := tmp12344
_ = let__12112

tmp12353 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12112)


var ifres12345 Obj

if True == tmp12353 {
tmp12346 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12345 = tmp12346


} else {
tmp12347 := Call(__e, PrimFunc(symshen_4_5_1out), let__12112)


let__12113 := tmp12347
_ = let__12113

tmp12348 := Call(__e, PrimFunc(symshen_4in_1_6), let__12112)


let__12114 := tmp12348
_ = let__12114

tmp12349 := Call(__e, PrimFunc(symshen_4rules_1_6prolog), let__12110, let__12113)


let__12115 := tmp12349
_ = let__12115

tmp12350 := Call(__e, PrimFunc(symfn), let__12110)


tmp12351 := Call(__e, PrimFunc(symshen_4remember_1datatype), let__12110, tmp12350)


tmp12352 := Call(__e, PrimFunc(symshen_4comb), let__12114, tmp12351)


ifres12345 = tmp12352


}

ifres12341 = ifres12345


} else {
tmp12354 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12341 = tmp12354


}

let__12109 := ifres12341
_ = let__12109

tmp12357 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12109)


if True == tmp12357 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__12109)
return
}


}, 1)

tmp12358 := Call(__e, ns2_1set, symshen_4_5datatype_6, tmp12340)


_ = tmp12358

tmp12359 := MakeNative(func(__e *ControlFlow) {
V3426 := __e.Get(1)
_ = V3426
V3427 := __e.Get(2)
_ = V3427
tmp12360 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_ddatatypes_d)
}
__typedArg0 := symshen_4_ddatatypes_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp12361 := Call(__e, PrimFunc(symshen_4assoc_1_6), V3426, V3427, tmp12360)


tmp12362 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_ddatatypes_d, tmp12361)
}
__typedArg0 := symshen_4_ddatatypes_d
__typedArg1 := tmp12361
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp12362

tmp12363 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dalldatatypes_d)
}
__typedArg0 := symshen_4_dalldatatypes_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp12364 := Call(__e, PrimFunc(symshen_4assoc_1_6), V3426, V3427, tmp12363)


tmp12365 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dalldatatypes_d, tmp12364)
}
__typedArg0 := symshen_4_dalldatatypes_d
__typedArg1 := tmp12364
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp12365

__e.Return(V3426)
return


}, 2)

tmp12366 := Call(__e, ns2_1set, symshen_4remember_1datatype, tmp12359)


_ = tmp12366

tmp12367 := MakeNative(func(__e *ControlFlow) {
V3428 := __e.Get(1)
_ = V3428
tmp12368 := Call(__e, PrimFunc(symshen_4_5datatype_1rule_6), V3428)


let__12117 := tmp12368
_ = let__12117

tmp12381 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12117)


var ifres12369 Obj

if True == tmp12381 {
tmp12370 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12369 = tmp12370


} else {
tmp12371 := Call(__e, PrimFunc(symshen_4_5_1out), let__12117)


let__12118 := tmp12371
_ = let__12118

tmp12372 := Call(__e, PrimFunc(symshen_4in_1_6), let__12117)


let__12119 := tmp12372
_ = let__12119

tmp12373 := Call(__e, PrimFunc(symshen_4_5datatype_1rules_6), let__12119)


let__12120 := tmp12373
_ = let__12120

tmp12380 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12120)


var ifres12374 Obj

if True == tmp12380 {
tmp12375 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12374 = tmp12375


} else {
tmp12376 := Call(__e, PrimFunc(symshen_4_5_1out), let__12120)


let__12121 := tmp12376
_ = let__12121

tmp12377 := Call(__e, PrimFunc(symshen_4in_1_6), let__12120)


let__12122 := tmp12377
_ = let__12122

tmp12378 := Call(__e, PrimFunc(symappend), let__12118, let__12121)


tmp12379 := Call(__e, PrimFunc(symshen_4comb), let__12122, tmp12378)


ifres12374 = tmp12379


}

ifres12369 = ifres12374


}

let__12116 := ifres12369
_ = let__12116

tmp12397 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12116)


if True == tmp12397 {
tmp12382 := Call(__e, PrimFunc(sym_5_b_6), V3428)


let__12124 := tmp12382
_ = let__12124

tmp12393 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12124)


var ifres12383 Obj

if True == tmp12393 {
tmp12384 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12383 = tmp12384


} else {
tmp12385 := Call(__e, PrimFunc(symshen_4_5_1out), let__12124)


let__12125 := tmp12385
_ = let__12125

tmp12386 := Call(__e, PrimFunc(symshen_4in_1_6), let__12124)


let__12126 := tmp12386
_ = let__12126

tmp12391 := Call(__e, PrimFunc(symempty_2), let__12125)


var ifres12387 Obj

if True == tmp12391 {
ifres12387 = Nil


} else {
tmp12388 := Call(__e, PrimFunc(symshen_4app), let__12125, MakeString("\n ..."), symshen_4r)


tmp12390 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("datatype syntax error here:\n "))
__typedS1, __typedOK1 := TypedString(tmp12388)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("datatype syntax error here:\n ")
__typedArg1 := tmp12388
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("datatype syntax error here:\n "))
__typedS1, __typedOK1 := TypedString(tmp12388)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("datatype syntax error here:\n ")
__typedArg1 := tmp12388
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})()

ifres12387 = tmp12390


}

tmp12392 := Call(__e, PrimFunc(symshen_4comb), let__12126, ifres12387)


ifres12383 = tmp12392


}

let__12123 := ifres12383
_ = let__12123

tmp12395 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12123)


if True == tmp12395 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__12123)
return
}


} else {
__e.Return(let__12116)
return
}


}, 1)

tmp12398 := Call(__e, ns2_1set, symshen_4_5datatype_1rules_6, tmp12367)


_ = tmp12398

tmp12399 := MakeNative(func(__e *ControlFlow) {
V3440 := __e.Get(1)
_ = V3440
tmp12400 := Call(__e, PrimFunc(symshen_4_5single_6), V3440)


let__12128 := tmp12400
_ = let__12128

tmp12406 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12128)


var ifres12401 Obj

if True == tmp12406 {
tmp12402 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12401 = tmp12402


} else {
tmp12403 := Call(__e, PrimFunc(symshen_4_5_1out), let__12128)


let__12129 := tmp12403
_ = let__12129

tmp12404 := Call(__e, PrimFunc(symshen_4in_1_6), let__12128)


let__12130 := tmp12404
_ = let__12130

tmp12405 := Call(__e, PrimFunc(symshen_4comb), let__12130, let__12129)


ifres12401 = tmp12405


}

let__12127 := ifres12401
_ = let__12127

tmp12417 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12127)


if True == tmp12417 {
tmp12407 := Call(__e, PrimFunc(symshen_4_5double_6), V3440)


let__12132 := tmp12407
_ = let__12132

tmp12413 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12132)


var ifres12408 Obj

if True == tmp12413 {
tmp12409 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12408 = tmp12409


} else {
tmp12410 := Call(__e, PrimFunc(symshen_4_5_1out), let__12132)


let__12133 := tmp12410
_ = let__12133

tmp12411 := Call(__e, PrimFunc(symshen_4in_1_6), let__12132)


let__12134 := tmp12411
_ = let__12134

tmp12412 := Call(__e, PrimFunc(symshen_4comb), let__12134, let__12133)


ifres12408 = tmp12412


}

let__12131 := ifres12408
_ = let__12131

tmp12415 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12131)


if True == tmp12415 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__12131)
return
}


} else {
__e.Return(let__12127)
return
}


}, 1)

tmp12418 := Call(__e, ns2_1set, symshen_4_5datatype_1rule_6, tmp12399)


_ = tmp12418

tmp12419 := MakeNative(func(__e *ControlFlow) {
V3449 := __e.Get(1)
_ = V3449
tmp12420 := Call(__e, PrimFunc(symshen_4_5sides_6), V3449)


let__12136 := tmp12420
_ = let__12136

tmp12452 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12136)


var ifres12421 Obj

if True == tmp12452 {
tmp12422 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12421 = tmp12422


} else {
tmp12423 := Call(__e, PrimFunc(symshen_4_5_1out), let__12136)


let__12137 := tmp12423
_ = let__12137

tmp12424 := Call(__e, PrimFunc(symshen_4in_1_6), let__12136)


let__12138 := tmp12424
_ = let__12138

tmp12425 := Call(__e, PrimFunc(symshen_4_5prems_6), let__12138)


let__12139 := tmp12425
_ = let__12139

tmp12451 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12139)


var ifres12426 Obj

if True == tmp12451 {
tmp12427 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12426 = tmp12427


} else {
tmp12428 := Call(__e, PrimFunc(symshen_4_5_1out), let__12139)


let__12140 := tmp12428
_ = let__12140

tmp12429 := Call(__e, PrimFunc(symshen_4in_1_6), let__12139)


let__12141 := tmp12429
_ = let__12141

tmp12430 := Call(__e, PrimFunc(symshen_4_5sng_6), let__12141)


let__12142 := tmp12430
_ = let__12142

tmp12450 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12142)


var ifres12431 Obj

if True == tmp12450 {
tmp12432 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12431 = tmp12432


} else {
tmp12433 := Call(__e, PrimFunc(symshen_4in_1_6), let__12142)


let__12143 := tmp12433
_ = let__12143

tmp12434 := Call(__e, PrimFunc(symshen_4_5conc_6), let__12143)


let__12144 := tmp12434
_ = let__12144

tmp12449 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12144)


var ifres12435 Obj

if True == tmp12449 {
tmp12436 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12435 = tmp12436


} else {
tmp12437 := Call(__e, PrimFunc(symshen_4_5_1out), let__12144)


let__12145 := tmp12437
_ = let__12145

tmp12438 := Call(__e, PrimFunc(symshen_4in_1_6), let__12144)


let__12146 := tmp12438
_ = let__12146

tmp12439 := Call(__e, PrimFunc(symshen_4_5sc_6), let__12146)


let__12147 := tmp12439
_ = let__12147

tmp12448 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12147)


var ifres12440 Obj

if True == tmp12448 {
tmp12441 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12440 = tmp12441


} else {
tmp12442 := Call(__e, PrimFunc(symshen_4in_1_6), let__12147)


let__12148 := tmp12442
_ = let__12148

tmp12443 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12145, Nil)
}
__typedArg0 := let__12145
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12444 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12140, tmp12443)
}
__typedArg0 := let__12140
__typedArg1 := tmp12443
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12445 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12137, tmp12444)
}
__typedArg0 := let__12137
__typedArg1 := tmp12444
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12446 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12445, Nil)
}
__typedArg0 := tmp12445
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12447 := Call(__e, PrimFunc(symshen_4comb), let__12148, tmp12446)


ifres12440 = tmp12447


}

ifres12435 = ifres12440


}

ifres12431 = ifres12435


}

ifres12426 = ifres12431


}

ifres12421 = ifres12426


}

let__12135 := ifres12421
_ = let__12135

tmp12454 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12135)


if True == tmp12454 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__12135)
return
}


}, 1)

tmp12455 := Call(__e, ns2_1set, symshen_4_5single_6, tmp12419)


_ = tmp12455

tmp12456 := MakeNative(func(__e *ControlFlow) {
V3464 := __e.Get(1)
_ = V3464
tmp12457 := Call(__e, PrimFunc(symshen_4_5sides_6), V3464)


let__12150 := tmp12457
_ = let__12150

tmp12488 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12150)


var ifres12458 Obj

if True == tmp12488 {
tmp12459 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12458 = tmp12459


} else {
tmp12460 := Call(__e, PrimFunc(symshen_4_5_1out), let__12150)


let__12151 := tmp12460
_ = let__12151

tmp12461 := Call(__e, PrimFunc(symshen_4in_1_6), let__12150)


let__12152 := tmp12461
_ = let__12152

tmp12462 := Call(__e, PrimFunc(symshen_4_5formulae_6), let__12152)


let__12153 := tmp12462
_ = let__12153

tmp12487 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12153)


var ifres12463 Obj

if True == tmp12487 {
tmp12464 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12463 = tmp12464


} else {
tmp12465 := Call(__e, PrimFunc(symshen_4_5_1out), let__12153)


let__12154 := tmp12465
_ = let__12154

tmp12466 := Call(__e, PrimFunc(symshen_4in_1_6), let__12153)


let__12155 := tmp12466
_ = let__12155

tmp12467 := Call(__e, PrimFunc(symshen_4_5dbl_6), let__12155)


let__12156 := tmp12467
_ = let__12156

tmp12486 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12156)


var ifres12468 Obj

if True == tmp12486 {
tmp12469 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12468 = tmp12469


} else {
tmp12470 := Call(__e, PrimFunc(symshen_4in_1_6), let__12156)


let__12157 := tmp12470
_ = let__12157

tmp12471 := Call(__e, PrimFunc(symshen_4_5formula_6), let__12157)


let__12158 := tmp12471
_ = let__12158

tmp12485 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12158)


var ifres12472 Obj

if True == tmp12485 {
tmp12473 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12472 = tmp12473


} else {
tmp12474 := Call(__e, PrimFunc(symshen_4_5_1out), let__12158)


let__12159 := tmp12474
_ = let__12159

tmp12475 := Call(__e, PrimFunc(symshen_4in_1_6), let__12158)


let__12160 := tmp12475
_ = let__12160

tmp12476 := Call(__e, PrimFunc(symshen_4_5sc_6), let__12160)


let__12161 := tmp12476
_ = let__12161

tmp12484 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12161)


var ifres12477 Obj

if True == tmp12484 {
tmp12478 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12477 = tmp12478


} else {
tmp12479 := Call(__e, PrimFunc(symshen_4in_1_6), let__12161)


let__12162 := tmp12479
_ = let__12162

tmp12480 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12159, Nil)
}
__typedArg0 := let__12159
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

tmp12482 := Call(__e, PrimFunc(symshen_4lr_1rule), let__12151, let__12154, tmp12481)


tmp12483 := Call(__e, PrimFunc(symshen_4comb), let__12162, tmp12482)


ifres12477 = tmp12483


}

ifres12472 = ifres12477


}

ifres12468 = ifres12472


}

ifres12463 = ifres12468


}

ifres12458 = ifres12463


}

let__12149 := ifres12458
_ = let__12149

tmp12490 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12149)


if True == tmp12490 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__12149)
return
}


}, 1)

tmp12491 := Call(__e, ns2_1set, symshen_4_5double_6, tmp12456)


_ = tmp12491

tmp12492 := MakeNative(func(__e *ControlFlow) {
V3479 := __e.Get(1)
_ = V3479
tmp12493 := Call(__e, PrimFunc(symshen_4_5formula_6), V3479)


let__12164 := tmp12493
_ = let__12164

tmp12513 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12164)


var ifres12494 Obj

if True == tmp12513 {
tmp12495 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12494 = tmp12495


} else {
tmp12496 := Call(__e, PrimFunc(symshen_4_5_1out), let__12164)


let__12165 := tmp12496
_ = let__12165

tmp12497 := Call(__e, PrimFunc(symshen_4in_1_6), let__12164)


let__12166 := tmp12497
_ = let__12166

tmp12498 := Call(__e, PrimFunc(symshen_4_5sc_6), let__12166)


let__12167 := tmp12498
_ = let__12167

tmp12512 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12167)


var ifres12499 Obj

if True == tmp12512 {
tmp12500 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12499 = tmp12500


} else {
tmp12501 := Call(__e, PrimFunc(symshen_4in_1_6), let__12167)


let__12168 := tmp12501
_ = let__12168

tmp12502 := Call(__e, PrimFunc(symshen_4_5formulae_6), let__12168)


let__12169 := tmp12502
_ = let__12169

tmp12511 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12169)


var ifres12503 Obj

if True == tmp12511 {
tmp12504 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12503 = tmp12504


} else {
tmp12505 := Call(__e, PrimFunc(symshen_4_5_1out), let__12169)


let__12170 := tmp12505
_ = let__12170

tmp12506 := Call(__e, PrimFunc(symshen_4in_1_6), let__12169)


let__12171 := tmp12506
_ = let__12171

tmp12507 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12165, Nil)
}
__typedArg0 := let__12165
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12508 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Nil, tmp12507)
}
__typedArg0 := Nil
__typedArg1 := tmp12507
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12509 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12508, let__12170)
}
__typedArg0 := tmp12508
__typedArg1 := let__12170
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12510 := Call(__e, PrimFunc(symshen_4comb), let__12171, tmp12509)


ifres12503 = tmp12510


}

ifres12499 = ifres12503


}

ifres12494 = ifres12499


}

let__12163 := ifres12494
_ = let__12163

tmp12532 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12163)


if True == tmp12532 {
tmp12514 := Call(__e, PrimFunc(symshen_4_5formula_6), V3479)


let__12173 := tmp12514
_ = let__12173

tmp12528 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12173)


var ifres12515 Obj

if True == tmp12528 {
tmp12516 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12515 = tmp12516


} else {
tmp12517 := Call(__e, PrimFunc(symshen_4_5_1out), let__12173)


let__12174 := tmp12517
_ = let__12174

tmp12518 := Call(__e, PrimFunc(symshen_4in_1_6), let__12173)


let__12175 := tmp12518
_ = let__12175

tmp12519 := Call(__e, PrimFunc(symshen_4_5sc_6), let__12175)


let__12176 := tmp12519
_ = let__12176

tmp12527 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12176)


var ifres12520 Obj

if True == tmp12527 {
tmp12521 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12520 = tmp12521


} else {
tmp12522 := Call(__e, PrimFunc(symshen_4in_1_6), let__12176)


let__12177 := tmp12522
_ = let__12177

tmp12523 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12174, Nil)
}
__typedArg0 := let__12174
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12524 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Nil, tmp12523)
}
__typedArg0 := Nil
__typedArg1 := tmp12523
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12525 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12524, Nil)
}
__typedArg0 := tmp12524
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12526 := Call(__e, PrimFunc(symshen_4comb), let__12177, tmp12525)


ifres12520 = tmp12526


}

ifres12515 = ifres12520


}

let__12172 := ifres12515
_ = let__12172

tmp12530 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12172)


if True == tmp12530 {
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

tmp12533 := Call(__e, ns2_1set, symshen_4_5formulae_6, tmp12492)


_ = tmp12533

tmp12534 := MakeNative(func(__e *ControlFlow) {
V3495 := __e.Get(1)
_ = V3495
tmp12535 := Call(__e, PrimFunc(symshen_4_5ass_6), V3495)


let__12179 := tmp12535
_ = let__12179

tmp12553 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12179)


var ifres12536 Obj

if True == tmp12553 {
tmp12537 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12536 = tmp12537


} else {
tmp12538 := Call(__e, PrimFunc(symshen_4_5_1out), let__12179)


let__12180 := tmp12538
_ = let__12180

tmp12539 := Call(__e, PrimFunc(symshen_4in_1_6), let__12179)


let__12181 := tmp12539
_ = let__12181

tmp12552 := Call(__e, PrimFunc(symshen_4hds_a_2), let__12181, sym_6_6)


var ifres12540 Obj

if True == tmp12552 {
tmp12541 := Call(__e, PrimFunc(symtail), let__12181)


let__12182 := tmp12541
_ = let__12182

tmp12542 := Call(__e, PrimFunc(symshen_4_5formula_6), let__12182)


let__12183 := tmp12542
_ = let__12183

tmp12550 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12183)


var ifres12543 Obj

if True == tmp12550 {
tmp12544 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12543 = tmp12544


} else {
tmp12545 := Call(__e, PrimFunc(symshen_4_5_1out), let__12183)


let__12184 := tmp12545
_ = let__12184

tmp12546 := Call(__e, PrimFunc(symshen_4in_1_6), let__12183)


let__12185 := tmp12546
_ = let__12185

tmp12547 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12184, Nil)
}
__typedArg0 := let__12184
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12548 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12180, tmp12547)
}
__typedArg0 := let__12180
__typedArg1 := tmp12547
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12549 := Call(__e, PrimFunc(symshen_4comb), let__12185, tmp12548)


ifres12543 = tmp12549


}

ifres12540 = ifres12543


} else {
tmp12551 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12540 = tmp12551


}

ifres12536 = ifres12540


}

let__12178 := ifres12536
_ = let__12178

tmp12566 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12178)


if True == tmp12566 {
tmp12554 := Call(__e, PrimFunc(symshen_4_5formula_6), V3495)


let__12187 := tmp12554
_ = let__12187

tmp12562 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12187)


var ifres12555 Obj

if True == tmp12562 {
tmp12556 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12555 = tmp12556


} else {
tmp12557 := Call(__e, PrimFunc(symshen_4_5_1out), let__12187)


let__12188 := tmp12557
_ = let__12188

tmp12558 := Call(__e, PrimFunc(symshen_4in_1_6), let__12187)


let__12189 := tmp12558
_ = let__12189

tmp12559 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12188, Nil)
}
__typedArg0 := let__12188
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12560 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Nil, tmp12559)
}
__typedArg0 := Nil
__typedArg1 := tmp12559
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12561 := Call(__e, PrimFunc(symshen_4comb), let__12189, tmp12560)


ifres12555 = tmp12561


}

let__12186 := ifres12555
_ = let__12186

tmp12564 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12186)


if True == tmp12564 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__12186)
return
}


} else {
__e.Return(let__12178)
return
}


}, 1)

tmp12567 := Call(__e, ns2_1set, symshen_4_5conc_6, tmp12534)


_ = tmp12567

tmp12568 := MakeNative(func(__e *ControlFlow) {
V3508 := __e.Get(1)
_ = V3508
tmp12569 := Call(__e, PrimFunc(symshen_4_5prem_6), V3508)


let__12191 := tmp12569
_ = let__12191

tmp12587 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12191)


var ifres12570 Obj

if True == tmp12587 {
tmp12571 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12570 = tmp12571


} else {
tmp12572 := Call(__e, PrimFunc(symshen_4_5_1out), let__12191)


let__12192 := tmp12572
_ = let__12192

tmp12573 := Call(__e, PrimFunc(symshen_4in_1_6), let__12191)


let__12193 := tmp12573
_ = let__12193

tmp12574 := Call(__e, PrimFunc(symshen_4_5sc_6), let__12193)


let__12194 := tmp12574
_ = let__12194

tmp12586 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12194)


var ifres12575 Obj

if True == tmp12586 {
tmp12576 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12575 = tmp12576


} else {
tmp12577 := Call(__e, PrimFunc(symshen_4in_1_6), let__12194)


let__12195 := tmp12577
_ = let__12195

tmp12578 := Call(__e, PrimFunc(symshen_4_5prems_6), let__12195)


let__12196 := tmp12578
_ = let__12196

tmp12585 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12196)


var ifres12579 Obj

if True == tmp12585 {
tmp12580 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12579 = tmp12580


} else {
tmp12581 := Call(__e, PrimFunc(symshen_4_5_1out), let__12196)


let__12197 := tmp12581
_ = let__12197

tmp12582 := Call(__e, PrimFunc(symshen_4in_1_6), let__12196)


let__12198 := tmp12582
_ = let__12198

tmp12583 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12192, let__12197)
}
__typedArg0 := let__12192
__typedArg1 := let__12197
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12584 := Call(__e, PrimFunc(symshen_4comb), let__12198, tmp12583)


ifres12579 = tmp12584


}

ifres12575 = ifres12579


}

ifres12570 = ifres12575


}

let__12190 := ifres12570
_ = let__12190

tmp12597 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12190)


if True == tmp12597 {
tmp12588 := Call(__e, PrimFunc(sym_5e_6), V3508)


let__12200 := tmp12588
_ = let__12200

tmp12593 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12200)


var ifres12589 Obj

if True == tmp12593 {
tmp12590 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12589 = tmp12590


} else {
tmp12591 := Call(__e, PrimFunc(symshen_4in_1_6), let__12200)


let__12201 := tmp12591
_ = let__12201

tmp12592 := Call(__e, PrimFunc(symshen_4comb), let__12201, Nil)


ifres12589 = tmp12592


}

let__12199 := ifres12589
_ = let__12199

tmp12595 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12199)


if True == tmp12595 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__12199)
return
}


} else {
__e.Return(let__12190)
return
}


}, 1)

tmp12598 := Call(__e, ns2_1set, symshen_4_5prems_6, tmp12568)


_ = tmp12598

tmp12599 := MakeNative(func(__e *ControlFlow) {
V3521 := __e.Get(1)
_ = V3521
tmp12604 := Call(__e, PrimFunc(symshen_4hds_a_2), V3521, sym_b)


var ifres12600 Obj

if True == tmp12604 {
tmp12601 := Call(__e, PrimFunc(symtail), V3521)


let__12203 := tmp12601
_ = let__12203

tmp12602 := Call(__e, PrimFunc(symshen_4comb), let__12203, sym_b)


ifres12600 = tmp12602


} else {
tmp12603 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12600 = tmp12603


}

let__12202 := ifres12600
_ = let__12202

tmp12638 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12202)


if True == tmp12638 {
tmp12605 := Call(__e, PrimFunc(symshen_4_5ass_6), V3521)


let__12205 := tmp12605
_ = let__12205

tmp12623 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12205)


var ifres12606 Obj

if True == tmp12623 {
tmp12607 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12606 = tmp12607


} else {
tmp12608 := Call(__e, PrimFunc(symshen_4_5_1out), let__12205)


let__12206 := tmp12608
_ = let__12206

tmp12609 := Call(__e, PrimFunc(symshen_4in_1_6), let__12205)


let__12207 := tmp12609
_ = let__12207

tmp12622 := Call(__e, PrimFunc(symshen_4hds_a_2), let__12207, sym_6_6)


var ifres12610 Obj

if True == tmp12622 {
tmp12611 := Call(__e, PrimFunc(symtail), let__12207)


let__12208 := tmp12611
_ = let__12208

tmp12612 := Call(__e, PrimFunc(symshen_4_5formula_6), let__12208)


let__12209 := tmp12612
_ = let__12209

tmp12620 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12209)


var ifres12613 Obj

if True == tmp12620 {
tmp12614 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12613 = tmp12614


} else {
tmp12615 := Call(__e, PrimFunc(symshen_4_5_1out), let__12209)


let__12210 := tmp12615
_ = let__12210

tmp12616 := Call(__e, PrimFunc(symshen_4in_1_6), let__12209)


let__12211 := tmp12616
_ = let__12211

tmp12617 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12210, Nil)
}
__typedArg0 := let__12210
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12618 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12206, tmp12617)
}
__typedArg0 := let__12206
__typedArg1 := tmp12617
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12619 := Call(__e, PrimFunc(symshen_4comb), let__12211, tmp12618)


ifres12613 = tmp12619


}

ifres12610 = ifres12613


} else {
tmp12621 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12610 = tmp12621


}

ifres12606 = ifres12610


}

let__12204 := ifres12606
_ = let__12204

tmp12636 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12204)


if True == tmp12636 {
tmp12624 := Call(__e, PrimFunc(symshen_4_5formula_6), V3521)


let__12213 := tmp12624
_ = let__12213

tmp12632 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12213)


var ifres12625 Obj

if True == tmp12632 {
tmp12626 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12625 = tmp12626


} else {
tmp12627 := Call(__e, PrimFunc(symshen_4_5_1out), let__12213)


let__12214 := tmp12627
_ = let__12214

tmp12628 := Call(__e, PrimFunc(symshen_4in_1_6), let__12213)


let__12215 := tmp12628
_ = let__12215

tmp12629 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12214, Nil)
}
__typedArg0 := let__12214
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12630 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Nil, tmp12629)
}
__typedArg0 := Nil
__typedArg1 := tmp12629
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12631 := Call(__e, PrimFunc(symshen_4comb), let__12215, tmp12630)


ifres12625 = tmp12631


}

let__12212 := ifres12625
_ = let__12212

tmp12634 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12212)


if True == tmp12634 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__12212)
return
}


} else {
__e.Return(let__12204)
return
}


} else {
__e.Return(let__12202)
return
}


}, 1)

tmp12639 := Call(__e, ns2_1set, symshen_4_5prem_6, tmp12599)


_ = tmp12639

tmp12640 := MakeNative(func(__e *ControlFlow) {
V3536 := __e.Get(1)
_ = V3536
tmp12641 := Call(__e, PrimFunc(symshen_4_5formula_6), V3536)


let__12217 := tmp12641
_ = let__12217

tmp12659 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12217)


var ifres12642 Obj

if True == tmp12659 {
tmp12643 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12642 = tmp12643


} else {
tmp12644 := Call(__e, PrimFunc(symshen_4_5_1out), let__12217)


let__12218 := tmp12644
_ = let__12218

tmp12645 := Call(__e, PrimFunc(symshen_4in_1_6), let__12217)


let__12219 := tmp12645
_ = let__12219

tmp12646 := Call(__e, PrimFunc(symshen_4_5iscomma_6), let__12219)


let__12220 := tmp12646
_ = let__12220

tmp12658 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12220)


var ifres12647 Obj

if True == tmp12658 {
tmp12648 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12647 = tmp12648


} else {
tmp12649 := Call(__e, PrimFunc(symshen_4in_1_6), let__12220)


let__12221 := tmp12649
_ = let__12221

tmp12650 := Call(__e, PrimFunc(symshen_4_5ass_6), let__12221)


let__12222 := tmp12650
_ = let__12222

tmp12657 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12222)


var ifres12651 Obj

if True == tmp12657 {
tmp12652 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12651 = tmp12652


} else {
tmp12653 := Call(__e, PrimFunc(symshen_4_5_1out), let__12222)


let__12223 := tmp12653
_ = let__12223

tmp12654 := Call(__e, PrimFunc(symshen_4in_1_6), let__12222)


let__12224 := tmp12654
_ = let__12224

tmp12655 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12218, let__12223)
}
__typedArg0 := let__12218
__typedArg1 := let__12223
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12656 := Call(__e, PrimFunc(symshen_4comb), let__12224, tmp12655)


ifres12651 = tmp12656


}

ifres12647 = ifres12651


}

ifres12642 = ifres12647


}

let__12216 := ifres12642
_ = let__12216

tmp12679 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12216)


if True == tmp12679 {
tmp12660 := Call(__e, PrimFunc(symshen_4_5formula_6), V3536)


let__12226 := tmp12660
_ = let__12226

tmp12667 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12226)


var ifres12661 Obj

if True == tmp12667 {
tmp12662 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12661 = tmp12662


} else {
tmp12663 := Call(__e, PrimFunc(symshen_4_5_1out), let__12226)


let__12227 := tmp12663
_ = let__12227

tmp12664 := Call(__e, PrimFunc(symshen_4in_1_6), let__12226)


let__12228 := tmp12664
_ = let__12228

tmp12665 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12227, Nil)
}
__typedArg0 := let__12227
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12666 := Call(__e, PrimFunc(symshen_4comb), let__12228, tmp12665)


ifres12661 = tmp12666


}

let__12225 := ifres12661
_ = let__12225

tmp12677 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12225)


if True == tmp12677 {
tmp12668 := Call(__e, PrimFunc(sym_5e_6), V3536)


let__12230 := tmp12668
_ = let__12230

tmp12673 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12230)


var ifres12669 Obj

if True == tmp12673 {
tmp12670 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12669 = tmp12670


} else {
tmp12671 := Call(__e, PrimFunc(symshen_4in_1_6), let__12230)


let__12231 := tmp12671
_ = let__12231

tmp12672 := Call(__e, PrimFunc(symshen_4comb), let__12231, Nil)


ifres12669 = tmp12672


}

let__12229 := ifres12669
_ = let__12229

tmp12675 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12229)


if True == tmp12675 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__12229)
return
}


} else {
__e.Return(let__12225)
return
}


} else {
__e.Return(let__12216)
return
}


}, 1)

tmp12680 := Call(__e, ns2_1set, symshen_4_5ass_6, tmp12640)


_ = tmp12680

tmp12681 := MakeNative(func(__e *ControlFlow) {
V3553 := __e.Get(1)
_ = V3553
tmp12691 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3553)
}
__typedArg0 := V3553
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12682 Obj

if True == tmp12691 {
tmp12683 := Call(__e, PrimFunc(symhead), V3553)


let__12233 := tmp12683
_ = let__12233

tmp12684 := Call(__e, PrimFunc(symtail), V3553)


let__12234 := tmp12684
_ = let__12234

tmp12688 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(","))
}
__typedArg0 := MakeString(",")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp12689 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__12233, tmp12688)
}
__typedArg0 := let__12233
__typedArg1 := tmp12688
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres12685 Obj

if True == tmp12689 {
tmp12686 := Call(__e, PrimFunc(symshen_4comb), let__12234, symshen_4skip)


ifres12685 = tmp12686


} else {
tmp12687 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12685 = tmp12687


}

ifres12682 = ifres12685


} else {
tmp12690 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12682 = tmp12690


}

let__12232 := ifres12682
_ = let__12232

tmp12693 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12232)


if True == tmp12693 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__12232)
return
}


}, 1)

tmp12694 := Call(__e, ns2_1set, symshen_4_5iscomma_6, tmp12681)


_ = tmp12694

tmp12695 := MakeNative(func(__e *ControlFlow) {
V3557 := __e.Get(1)
_ = V3557
tmp12696 := Call(__e, PrimFunc(symshen_4_5expr_6), V3557)


let__12236 := tmp12696
_ = let__12236

tmp12719 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12236)


var ifres12697 Obj

if True == tmp12719 {
tmp12698 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12697 = tmp12698


} else {
tmp12699 := Call(__e, PrimFunc(symshen_4_5_1out), let__12236)


let__12237 := tmp12699
_ = let__12237

tmp12700 := Call(__e, PrimFunc(symshen_4in_1_6), let__12236)


let__12238 := tmp12700
_ = let__12238

tmp12701 := Call(__e, PrimFunc(symshen_4_5iscolon_6), let__12238)


let__12239 := tmp12701
_ = let__12239

tmp12718 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12239)


var ifres12702 Obj

if True == tmp12718 {
tmp12703 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12702 = tmp12703


} else {
tmp12704 := Call(__e, PrimFunc(symshen_4in_1_6), let__12239)


let__12240 := tmp12704
_ = let__12240

tmp12705 := Call(__e, PrimFunc(symshen_4_5type_6), let__12240)


let__12241 := tmp12705
_ = let__12241

tmp12717 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12241)


var ifres12706 Obj

if True == tmp12717 {
tmp12707 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12706 = tmp12707


} else {
tmp12708 := Call(__e, PrimFunc(symshen_4_5_1out), let__12241)


let__12242 := tmp12708
_ = let__12242

tmp12709 := Call(__e, PrimFunc(symshen_4in_1_6), let__12241)


let__12243 := tmp12709
_ = let__12243

tmp12710 := Call(__e, PrimFunc(symshen_4curry), let__12237)


tmp12711 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp12712 := Call(__e, PrimFunc(symshen_4rectify_1type), let__12242)


tmp12713 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12712, Nil)
}
__typedArg0 := tmp12712
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12714 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12711, tmp12713)
}
__typedArg0 := tmp12711
__typedArg1 := tmp12713
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12715 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12710, tmp12714)
}
__typedArg0 := tmp12710
__typedArg1 := tmp12714
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12716 := Call(__e, PrimFunc(symshen_4comb), let__12243, tmp12715)


ifres12706 = tmp12716


}

ifres12702 = ifres12706


}

ifres12697 = ifres12702


}

let__12235 := ifres12697
_ = let__12235

tmp12730 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12235)


if True == tmp12730 {
tmp12720 := Call(__e, PrimFunc(symshen_4_5expr_6), V3557)


let__12245 := tmp12720
_ = let__12245

tmp12726 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12245)


var ifres12721 Obj

if True == tmp12726 {
tmp12722 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12721 = tmp12722


} else {
tmp12723 := Call(__e, PrimFunc(symshen_4_5_1out), let__12245)


let__12246 := tmp12723
_ = let__12246

tmp12724 := Call(__e, PrimFunc(symshen_4in_1_6), let__12245)


let__12247 := tmp12724
_ = let__12247

tmp12725 := Call(__e, PrimFunc(symshen_4comb), let__12247, let__12246)


ifres12721 = tmp12725


}

let__12244 := ifres12721
_ = let__12244

tmp12728 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12244)


if True == tmp12728 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__12244)
return
}


} else {
__e.Return(let__12235)
return
}


}, 1)

tmp12731 := Call(__e, ns2_1set, symshen_4_5formula_6, tmp12695)


_ = tmp12731

tmp12732 := MakeNative(func(__e *ControlFlow) {
V3571 := __e.Get(1)
_ = V3571
tmp12742 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3571)
}
__typedArg0 := V3571
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12733 Obj

if True == tmp12742 {
tmp12734 := Call(__e, PrimFunc(symhead), V3571)


let__12249 := tmp12734
_ = let__12249

tmp12735 := Call(__e, PrimFunc(symtail), V3571)


let__12250 := tmp12735
_ = let__12250

tmp12739 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp12740 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__12249, tmp12739)
}
__typedArg0 := let__12249
__typedArg1 := tmp12739
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres12736 Obj

if True == tmp12740 {
tmp12737 := Call(__e, PrimFunc(symshen_4comb), let__12250, symshen_4skip)


ifres12736 = tmp12737


} else {
tmp12738 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12736 = tmp12738


}

ifres12733 = ifres12736


} else {
tmp12741 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12733 = tmp12741


}

let__12248 := ifres12733
_ = let__12248

tmp12744 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12248)


if True == tmp12744 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__12248)
return
}


}, 1)

tmp12745 := Call(__e, ns2_1set, symshen_4_5iscolon_6, tmp12732)


_ = tmp12745

tmp12746 := MakeNative(func(__e *ControlFlow) {
V3575 := __e.Get(1)
_ = V3575
tmp12747 := Call(__e, PrimFunc(symshen_4_5side_6), V3575)


let__12252 := tmp12747
_ = let__12252

tmp12760 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12252)


var ifres12748 Obj

if True == tmp12760 {
tmp12749 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12748 = tmp12749


} else {
tmp12750 := Call(__e, PrimFunc(symshen_4_5_1out), let__12252)


let__12253 := tmp12750
_ = let__12253

tmp12751 := Call(__e, PrimFunc(symshen_4in_1_6), let__12252)


let__12254 := tmp12751
_ = let__12254

tmp12752 := Call(__e, PrimFunc(symshen_4_5sides_6), let__12254)


let__12255 := tmp12752
_ = let__12255

tmp12759 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12255)


var ifres12753 Obj

if True == tmp12759 {
tmp12754 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12753 = tmp12754


} else {
tmp12755 := Call(__e, PrimFunc(symshen_4_5_1out), let__12255)


let__12256 := tmp12755
_ = let__12256

tmp12756 := Call(__e, PrimFunc(symshen_4in_1_6), let__12255)


let__12257 := tmp12756
_ = let__12257

tmp12757 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12253, let__12256)
}
__typedArg0 := let__12253
__typedArg1 := let__12256
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12758 := Call(__e, PrimFunc(symshen_4comb), let__12257, tmp12757)


ifres12753 = tmp12758


}

ifres12748 = ifres12753


}

let__12251 := ifres12748
_ = let__12251

tmp12770 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12251)


if True == tmp12770 {
tmp12761 := Call(__e, PrimFunc(sym_5e_6), V3575)


let__12259 := tmp12761
_ = let__12259

tmp12766 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12259)


var ifres12762 Obj

if True == tmp12766 {
tmp12763 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12762 = tmp12763


} else {
tmp12764 := Call(__e, PrimFunc(symshen_4in_1_6), let__12259)


let__12260 := tmp12764
_ = let__12260

tmp12765 := Call(__e, PrimFunc(symshen_4comb), let__12260, Nil)


ifres12762 = tmp12765


}

let__12258 := ifres12762
_ = let__12258

tmp12768 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12258)


if True == tmp12768 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__12258)
return
}


} else {
__e.Return(let__12251)
return
}


}, 1)

tmp12771 := Call(__e, ns2_1set, symshen_4_5sides_6, tmp12746)


_ = tmp12771

tmp12772 := MakeNative(func(__e *ControlFlow) {
V3586 := __e.Get(1)
_ = V3586
tmp12784 := Call(__e, PrimFunc(symshen_4hds_a_2), V3586, symif)


var ifres12773 Obj

if True == tmp12784 {
tmp12774 := Call(__e, PrimFunc(symtail), V3586)


let__12262 := tmp12774
_ = let__12262

tmp12782 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__12262)
}
__typedArg0 := let__12262
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12775 Obj

if True == tmp12782 {
tmp12776 := Call(__e, PrimFunc(symhead), let__12262)


let__12263 := tmp12776
_ = let__12263

tmp12777 := Call(__e, PrimFunc(symtail), let__12262)


let__12264 := tmp12777
_ = let__12264

tmp12778 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12263, Nil)
}
__typedArg0 := let__12263
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12779 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp12778)
}
__typedArg0 := symif
__typedArg1 := tmp12778
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12780 := Call(__e, PrimFunc(symshen_4comb), let__12264, tmp12779)


ifres12775 = tmp12780


} else {
tmp12781 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12775 = tmp12781


}

ifres12773 = ifres12775


} else {
tmp12783 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12773 = tmp12783


}

let__12261 := ifres12773
_ = let__12261

tmp12840 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12261)


if True == tmp12840 {
tmp12802 := Call(__e, PrimFunc(symshen_4hds_a_2), V3586, symlet)


var ifres12785 Obj

if True == tmp12802 {
tmp12786 := Call(__e, PrimFunc(symtail), V3586)


let__12266 := tmp12786
_ = let__12266

tmp12800 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__12266)
}
__typedArg0 := let__12266
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12787 Obj

if True == tmp12800 {
tmp12788 := Call(__e, PrimFunc(symhead), let__12266)


let__12267 := tmp12788
_ = let__12267

tmp12789 := Call(__e, PrimFunc(symtail), let__12266)


let__12268 := tmp12789
_ = let__12268

tmp12798 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__12268)
}
__typedArg0 := let__12268
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12790 Obj

if True == tmp12798 {
tmp12791 := Call(__e, PrimFunc(symhead), let__12268)


let__12269 := tmp12791
_ = let__12269

tmp12792 := Call(__e, PrimFunc(symtail), let__12268)


let__12270 := tmp12792
_ = let__12270

tmp12793 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12269, Nil)
}
__typedArg0 := let__12269
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12794 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12267, tmp12793)
}
__typedArg0 := let__12267
__typedArg1 := tmp12793
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12795 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp12794)
}
__typedArg0 := symlet
__typedArg1 := tmp12794
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12796 := Call(__e, PrimFunc(symshen_4comb), let__12270, tmp12795)


ifres12790 = tmp12796


} else {
tmp12797 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12790 = tmp12797


}

ifres12787 = ifres12790


} else {
tmp12799 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12787 = tmp12799


}

ifres12785 = ifres12787


} else {
tmp12801 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12785 = tmp12801


}

let__12265 := ifres12785
_ = let__12265

tmp12838 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12265)


if True == tmp12838 {
tmp12817 := Call(__e, PrimFunc(symshen_4hds_a_2), V3586, symctxt)


var ifres12803 Obj

if True == tmp12817 {
tmp12804 := Call(__e, PrimFunc(symtail), V3586)


let__12272 := tmp12804
_ = let__12272

tmp12815 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__12272)
}
__typedArg0 := let__12272
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12805 Obj

if True == tmp12815 {
tmp12806 := Call(__e, PrimFunc(symhead), let__12272)


let__12273 := tmp12806
_ = let__12273

tmp12807 := Call(__e, PrimFunc(symtail), let__12272)


let__12274 := tmp12807
_ = let__12274

tmp12813 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(let__12273)
}
__typedArg0 := let__12273
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

var ifres12808 Obj

if True == tmp12813 {
tmp12809 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12273, Nil)
}
__typedArg0 := let__12273
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12810 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symctxt, tmp12809)
}
__typedArg0 := symctxt
__typedArg1 := tmp12809
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12811 := Call(__e, PrimFunc(symshen_4comb), let__12274, tmp12810)


ifres12808 = tmp12811


} else {
tmp12812 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12808 = tmp12812


}

ifres12805 = ifres12808


} else {
tmp12814 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12805 = tmp12814


}

ifres12803 = ifres12805


} else {
tmp12816 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12803 = tmp12816


}

let__12271 := ifres12803
_ = let__12271

tmp12836 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12271)


if True == tmp12836 {
tmp12832 := Call(__e, PrimFunc(symshen_4hds_a_2), V3586, symsqts)


var ifres12818 Obj

if True == tmp12832 {
tmp12819 := Call(__e, PrimFunc(symtail), V3586)


let__12276 := tmp12819
_ = let__12276

tmp12830 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__12276)
}
__typedArg0 := let__12276
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12820 Obj

if True == tmp12830 {
tmp12821 := Call(__e, PrimFunc(symhead), let__12276)


let__12277 := tmp12821
_ = let__12277

tmp12822 := Call(__e, PrimFunc(symtail), let__12276)


let__12278 := tmp12822
_ = let__12278

tmp12828 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(let__12277)
}
__typedArg0 := let__12277
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

var ifres12823 Obj

if True == tmp12828 {
tmp12824 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12277, Nil)
}
__typedArg0 := let__12277
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12825 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsqts, tmp12824)
}
__typedArg0 := symsqts
__typedArg1 := tmp12824
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12826 := Call(__e, PrimFunc(symshen_4comb), let__12278, tmp12825)


ifres12823 = tmp12826


} else {
tmp12827 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12823 = tmp12827


}

ifres12820 = ifres12823


} else {
tmp12829 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12820 = tmp12829


}

ifres12818 = ifres12820


} else {
tmp12831 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12818 = tmp12831


}

let__12275 := ifres12818
_ = let__12275

tmp12834 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12275)


if True == tmp12834 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__12275)
return
}


} else {
__e.Return(let__12271)
return
}


} else {
__e.Return(let__12265)
return
}


} else {
__e.Return(let__12261)
return
}


}, 1)

tmp12841 := Call(__e, ns2_1set, symshen_4_5side_6, tmp12772)


_ = tmp12841

tmp12842 := MakeNative(func(__e *ControlFlow) {
V3611 := __e.Get(1)
_ = V3611
V3612 := __e.Get(2)
_ = V3612
V3613 := __e.Get(3)
_ = V3613
tmp12872 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3613)
}
__typedArg0 := V3613
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12859 Obj

if True == tmp12872 {
tmp12870 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3613)
}
__typedArg0 := V3613
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

var ifres12861 Obj

if True == tmp12871 {
tmp12868 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3613)
}
__typedArg0 := V3613
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12869 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp12868)
}
__typedArg0 := tmp12868
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12863 Obj

if True == tmp12869 {
tmp12865 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3613)
}
__typedArg0 := V3613
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12866 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp12865)
}
__typedArg0 := tmp12865
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12867 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp12866)
}
__typedArg0 := Nil
__typedArg1 := tmp12866
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres12864 Obj

if True == tmp12867 {
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

var ifres12860 Obj

if True == ifres12861 {
ifres12860 = True


} else {
ifres12860 = False


}

ifres12859 = ifres12860


} else {
ifres12859 = False


}

if True == ifres12859 {
tmp12843 := Call(__e, PrimFunc(symgensym), symP)


let__12279 := tmp12843
_ = let__12279

tmp12844 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3613)
}
__typedArg0 := V3613
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12845 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12279, Nil)
}
__typedArg0 := let__12279
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12846 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12844, tmp12845)
}
__typedArg0 := tmp12844
__typedArg1 := tmp12845
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__12280 := tmp12846
_ = let__12280

tmp12847 := Call(__e, PrimFunc(symshen_4coll_1formulae), V3612)


tmp12848 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12279, Nil)
}
__typedArg0 := let__12279
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12849 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12847, tmp12848)
}
__typedArg0 := tmp12847
__typedArg1 := tmp12848
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__12281 := tmp12849
_ = let__12281

tmp12850 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12281, Nil)
}
__typedArg0 := let__12281
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12851 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12280, Nil)
}
__typedArg0 := let__12280
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12852 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12850, tmp12851)
}
__typedArg0 := tmp12850
__typedArg1 := tmp12851
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12853 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3611, tmp12852)
}
__typedArg0 := V3611
__typedArg1 := tmp12852
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__12282 := tmp12853
_ = let__12282

tmp12854 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3613, Nil)
}
__typedArg0 := V3613
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12855 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3612, tmp12854)
}
__typedArg0 := V3612
__typedArg1 := tmp12854
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12856 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3611, tmp12855)
}
__typedArg0 := V3611
__typedArg1 := tmp12855
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__12283 := tmp12856
_ = let__12283

tmp12857 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12282, Nil)
}
__typedArg0 := let__12282
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12283, tmp12857)
}
__typedArg0 := let__12283
__typedArg1 := tmp12857
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

tmp12873 := Call(__e, ns2_1set, symshen_4lr_1rule, tmp12842)


_ = tmp12873

tmp12874 := MakeNative(func(__e *ControlFlow) {
V3621 := __e.Get(1)
_ = V3621
tmp12903 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3621)
}
__typedArg0 := Nil
__typedArg1 := V3621
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp12903 {
__e.Return(Nil)
return
} else {
tmp12901 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3621)
}
__typedArg0 := V3621
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12881 Obj

if True == tmp12901 {
tmp12899 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3621)
}
__typedArg0 := V3621
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12900 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp12899)
}
__typedArg0 := tmp12899
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12883 Obj

if True == tmp12900 {
tmp12896 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3621)
}
__typedArg0 := V3621
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12897 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp12896)
}
__typedArg0 := tmp12896
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12898 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp12897)
}
__typedArg0 := Nil
__typedArg1 := tmp12897
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres12885 Obj

if True == tmp12898 {
tmp12893 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3621)
}
__typedArg0 := V3621
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12894 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp12893)
}
__typedArg0 := tmp12893
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12895 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp12894)
}
__typedArg0 := tmp12894
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12887 Obj

if True == tmp12895 {
tmp12889 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3621)
}
__typedArg0 := V3621
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12890 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp12889)
}
__typedArg0 := tmp12889
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12891 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp12890)
}
__typedArg0 := tmp12890
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12892 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp12891)
}
__typedArg0 := Nil
__typedArg1 := tmp12891
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres12888 Obj

if True == tmp12892 {
ifres12888 = True


} else {
ifres12888 = False


}

ifres12887 = ifres12888


} else {
ifres12887 = False


}

var ifres12886 Obj

if True == ifres12887 {
ifres12886 = True


} else {
ifres12886 = False


}

ifres12885 = ifres12886


} else {
ifres12885 = False


}

var ifres12884 Obj

if True == ifres12885 {
ifres12884 = True


} else {
ifres12884 = False


}

ifres12883 = ifres12884


} else {
ifres12883 = False


}

var ifres12882 Obj

if True == ifres12883 {
ifres12882 = True


} else {
ifres12882 = False


}

ifres12881 = ifres12882


} else {
ifres12881 = False


}

if True == ifres12881 {
tmp12875 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3621)
}
__typedArg0 := V3621
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12876 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp12875)
}
__typedArg0 := tmp12875
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12877 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp12876)
}
__typedArg0 := tmp12876
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12878 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3621)
}
__typedArg0 := V3621
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12879 := Call(__e, PrimFunc(symshen_4coll_1formulae), tmp12878)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12877, tmp12879)
}
__typedArg0 := tmp12877
__typedArg1 := tmp12879
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

tmp12904 := Call(__e, ns2_1set, symshen_4coll_1formulae, tmp12874)


_ = tmp12904

tmp12905 := MakeNative(func(__e *ControlFlow) {
V3622 := __e.Get(1)
_ = V3622
tmp12916 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3622)
}
__typedArg0 := V3622
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12906 Obj

if True == tmp12916 {
tmp12907 := Call(__e, PrimFunc(symhead), V3622)


let__12285 := tmp12907
_ = let__12285

tmp12908 := Call(__e, PrimFunc(symtail), V3622)


let__12286 := tmp12908
_ = let__12286

tmp12913 := Call(__e, PrimFunc(symshen_4key_1in_1sequent_1calculus_2), let__12285)


tmp12914 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp12913)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp12913
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres12909 Obj

if True == tmp12914 {
tmp12910 := Call(__e, PrimFunc(symmacroexpand), let__12285)


tmp12911 := Call(__e, PrimFunc(symshen_4comb), let__12286, tmp12910)


ifres12909 = tmp12911


} else {
tmp12912 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12909 = tmp12912


}

ifres12906 = ifres12909


} else {
tmp12915 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12906 = tmp12915


}

let__12284 := ifres12906
_ = let__12284

tmp12918 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12284)


if True == tmp12918 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__12284)
return
}


}, 1)

tmp12919 := Call(__e, ns2_1set, symshen_4_5expr_6, tmp12905)


_ = tmp12919

tmp12920 := MakeNative(func(__e *ControlFlow) {
V3626 := __e.Get(1)
_ = V3626
tmp12927 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(";"))
}
__typedArg0 := MakeString(";")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp12928 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(","))
}
__typedArg0 := MakeString(",")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp12929 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp12930 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1_1, Nil)
}
__typedArg0 := sym_5_1_1
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12931 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12929, tmp12930)
}
__typedArg0 := tmp12929
__typedArg1 := tmp12930
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12932 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12928, tmp12931)
}
__typedArg0 := tmp12928
__typedArg1 := tmp12931
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12933 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12927, tmp12932)
}
__typedArg0 := tmp12927
__typedArg1 := tmp12932
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12934 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_6_6, tmp12933)
}
__typedArg0 := sym_6_6
__typedArg1 := tmp12933
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12935 := Call(__e, PrimFunc(symelement_2), V3626, tmp12934)


if True == tmp12935 {
__e.Return(True)
return
} else {
tmp12925 := Call(__e, PrimFunc(symshen_4sng_2), V3626)


var ifres12922 Obj

if True == tmp12925 {
ifres12922 = True


} else {
tmp12924 := Call(__e, PrimFunc(symshen_4dbl_2), V3626)


var ifres12923 Obj

if True == tmp12924 {
ifres12923 = True


} else {
ifres12923 = False


}

ifres12922 = ifres12923


}

if True == ifres12922 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp12936 := Call(__e, ns2_1set, symshen_4key_1in_1sequent_1calculus_2, tmp12920)


_ = tmp12936

tmp12937 := MakeNative(func(__e *ControlFlow) {
V3627 := __e.Get(1)
_ = V3627
tmp12938 := Call(__e, PrimFunc(symshen_4_5expr_6), V3627)


let__12288 := tmp12938
_ = let__12288

tmp12944 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12288)


var ifres12939 Obj

if True == tmp12944 {
tmp12940 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12939 = tmp12940


} else {
tmp12941 := Call(__e, PrimFunc(symshen_4_5_1out), let__12288)


let__12289 := tmp12941
_ = let__12289

tmp12942 := Call(__e, PrimFunc(symshen_4in_1_6), let__12288)


let__12290 := tmp12942
_ = let__12290

tmp12943 := Call(__e, PrimFunc(symshen_4comb), let__12290, let__12289)


ifres12939 = tmp12943


}

let__12287 := ifres12939
_ = let__12287

tmp12946 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12287)


if True == tmp12946 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__12287)
return
}


}, 1)

tmp12947 := Call(__e, ns2_1set, symshen_4_5type_6, tmp12937)


_ = tmp12947

tmp12948 := MakeNative(func(__e *ControlFlow) {
V3632 := __e.Get(1)
_ = V3632
tmp12957 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3632)
}
__typedArg0 := V3632
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12949 Obj

if True == tmp12957 {
tmp12950 := Call(__e, PrimFunc(symhead), V3632)


let__12292 := tmp12950
_ = let__12292

tmp12951 := Call(__e, PrimFunc(symtail), V3632)


let__12293 := tmp12951
_ = let__12293

tmp12955 := Call(__e, PrimFunc(symshen_4dbl_2), let__12292)


var ifres12952 Obj

if True == tmp12955 {
tmp12953 := Call(__e, PrimFunc(symshen_4comb), let__12293, let__12292)


ifres12952 = tmp12953


} else {
tmp12954 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12952 = tmp12954


}

ifres12949 = ifres12952


} else {
tmp12956 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12949 = tmp12956


}

let__12291 := ifres12949
_ = let__12291

tmp12959 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12291)


if True == tmp12959 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__12291)
return
}


}, 1)

tmp12960 := Call(__e, ns2_1set, symshen_4_5dbl_6, tmp12948)


_ = tmp12960

tmp12961 := MakeNative(func(__e *ControlFlow) {
V3636 := __e.Get(1)
_ = V3636
tmp12970 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3636)
}
__typedArg0 := V3636
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12962 Obj

if True == tmp12970 {
tmp12963 := Call(__e, PrimFunc(symhead), V3636)


let__12295 := tmp12963
_ = let__12295

tmp12964 := Call(__e, PrimFunc(symtail), V3636)


let__12296 := tmp12964
_ = let__12296

tmp12968 := Call(__e, PrimFunc(symshen_4sng_2), let__12295)


var ifres12965 Obj

if True == tmp12968 {
tmp12966 := Call(__e, PrimFunc(symshen_4comb), let__12296, let__12295)


ifres12965 = tmp12966


} else {
tmp12967 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12965 = tmp12967


}

ifres12962 = ifres12965


} else {
tmp12969 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12962 = tmp12969


}

let__12294 := ifres12962
_ = let__12294

tmp12972 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__12294)


if True == tmp12972 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__12294)
return
}


}, 1)

tmp12973 := Call(__e, ns2_1set, symshen_4_5sng_6, tmp12961)


_ = tmp12973

tmp12974 := MakeNative(func(__e *ControlFlow) {
V3640 := __e.Get(1)
_ = V3640
tmp12979 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsymbol_2) {
return PrimIsSymbol(V3640)
}
__typedArg0 := V3640
return Call(__e, PrimFunc(symsymbol_2), __typedArg0)
})()

if True == tmp12979 {
tmp12976 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(V3640)
}
__typedArg0 := V3640
return Call(__e, PrimFunc(symstr), __typedArg0)
})()

tmp12977 := Call(__e, PrimFunc(symshen_4sng_1h_2), tmp12976)


if True == tmp12977 {
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

tmp12980 := Call(__e, ns2_1set, symshen_4sng_2, tmp12974)


_ = tmp12980

tmp12981 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__selftop:
V3643 := __self1
_ = V3643
tmp12990 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("___"), V3643)
}
__typedArg0 := MakeString("___")
__typedArg1 := V3643
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp12990 {
__e.Return(True)
return
} else {
tmp12988 := Call(__e, PrimFunc(symshen_4_7string_2), V3643)


var ifres12984 Obj

if True == tmp12988 {
tmp12986 := Call(__e, PrimFunc(symhdstr), V3643)


tmp12987 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("_"), tmp12986)
}
__typedArg0 := MakeString("_")
__typedArg1 := tmp12986
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres12985 Obj

if True == tmp12987 {
ifres12985 = True


} else {
ifres12985 = False


}

ifres12984 = ifres12985


} else {
ifres12984 = False


}

if True == ifres12984 {
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

tmp12991 := Call(__e, ns2_1set, symshen_4sng_1h_2, tmp12981)


_ = tmp12991

tmp12992 := MakeNative(func(__e *ControlFlow) {
V3644 := __e.Get(1)
_ = V3644
tmp12997 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsymbol_2) {
return PrimIsSymbol(V3644)
}
__typedArg0 := V3644
return Call(__e, PrimFunc(symsymbol_2), __typedArg0)
})()

if True == tmp12997 {
tmp12994 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(V3644)
}
__typedArg0 := V3644
return Call(__e, PrimFunc(symstr), __typedArg0)
})()

tmp12995 := Call(__e, PrimFunc(symshen_4dbl_1h_2), tmp12994)


if True == tmp12995 {
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

tmp12998 := Call(__e, ns2_1set, symshen_4dbl_2, tmp12992)


_ = tmp12998

tmp12999 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__selftop:
V3647 := __self1
_ = V3647
tmp13008 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("==="), V3647)
}
__typedArg0 := MakeString("===")
__typedArg1 := V3647
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13008 {
__e.Return(True)
return
} else {
tmp13006 := Call(__e, PrimFunc(symshen_4_7string_2), V3647)


var ifres13002 Obj

if True == tmp13006 {
tmp13004 := Call(__e, PrimFunc(symhdstr), V3647)


tmp13005 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("="), tmp13004)
}
__typedArg0 := MakeString("=")
__typedArg1 := tmp13004
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13003 Obj

if True == tmp13005 {
ifres13003 = True


} else {
ifres13003 = False


}

ifres13002 = ifres13003


} else {
ifres13002 = False


}

if True == ifres13002 {
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

tmp13009 := Call(__e, ns2_1set, symshen_4dbl_1h_2, tmp12999)


_ = tmp13009

tmp13010 := MakeNative(func(__e *ControlFlow) {
V3648 := __e.Get(1)
_ = V3648
V3649 := __e.Get(2)
_ = V3649
tmp13011 := MakeNative(func(__e *ControlFlow) {
Z3651 := __e.Get(1)
_ = Z3651
__e.TailApply(PrimFunc(symshen_4rule_1_6clause), Z3651)
return
}, 1)

tmp13012 := Call(__e, PrimFunc(symmapcan), tmp13011, V3649)


let__12297 := tmp13012
_ = let__12297

tmp13013 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3648, let__12297)
}
__typedArg0 := V3648
__typedArg1 := let__12297
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13014 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefprolog, tmp13013)
}
__typedArg0 := symdefprolog
__typedArg1 := tmp13013
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__12298 := tmp13014
_ = let__12298

__e.TailApply(PrimFunc(symeval), let__12298)
return


}, 2)

tmp13015 := Call(__e, ns2_1set, symshen_4rules_1_6prolog, tmp13010)


_ = tmp13015

tmp13016 := MakeNative(func(__e *ControlFlow) {
V3653 := __e.Get(1)
_ = V3653
tmp13076 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3653)
}
__typedArg0 := V3653
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13040 Obj

if True == tmp13076 {
tmp13074 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3653)
}
__typedArg0 := V3653
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13075 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13074)
}
__typedArg0 := tmp13074
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13042 Obj

if True == tmp13075 {
tmp13071 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3653)
}
__typedArg0 := V3653
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13072 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13071)
}
__typedArg0 := tmp13071
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13073 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13072)
}
__typedArg0 := tmp13072
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13044 Obj

if True == tmp13073 {
tmp13067 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3653)
}
__typedArg0 := V3653
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13068 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13067)
}
__typedArg0 := tmp13067
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13069 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13068)
}
__typedArg0 := tmp13068
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13070 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13069)
}
__typedArg0 := tmp13069
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13046 Obj

if True == tmp13070 {
tmp13062 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3653)
}
__typedArg0 := V3653
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13063 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13062)
}
__typedArg0 := tmp13062
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13064 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13063)
}
__typedArg0 := tmp13063
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13065 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13064)
}
__typedArg0 := tmp13064
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13066 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13065)
}
__typedArg0 := tmp13065
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13048 Obj

if True == tmp13066 {
tmp13056 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3653)
}
__typedArg0 := V3653
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13057 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13056)
}
__typedArg0 := tmp13056
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13058 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13057)
}
__typedArg0 := tmp13057
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13059 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13058)
}
__typedArg0 := tmp13058
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13060 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13059)
}
__typedArg0 := tmp13059
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13061 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp13060)
}
__typedArg0 := Nil
__typedArg1 := tmp13060
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13050 Obj

if True == tmp13061 {
tmp13052 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3653)
}
__typedArg0 := V3653
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13053 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13052)
}
__typedArg0 := tmp13052
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13054 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13053)
}
__typedArg0 := tmp13053
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13055 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp13054)
}
__typedArg0 := Nil
__typedArg1 := tmp13054
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13051 Obj

if True == tmp13055 {
ifres13051 = True


} else {
ifres13051 = False


}

ifres13050 = ifres13051


} else {
ifres13050 = False


}

var ifres13049 Obj

if True == ifres13050 {
ifres13049 = True


} else {
ifres13049 = False


}

ifres13048 = ifres13049


} else {
ifres13048 = False


}

var ifres13047 Obj

if True == ifres13048 {
ifres13047 = True


} else {
ifres13047 = False


}

ifres13046 = ifres13047


} else {
ifres13046 = False


}

var ifres13045 Obj

if True == ifres13046 {
ifres13045 = True


} else {
ifres13045 = False


}

ifres13044 = ifres13045


} else {
ifres13044 = False


}

var ifres13043 Obj

if True == ifres13044 {
ifres13043 = True


} else {
ifres13043 = False


}

ifres13042 = ifres13043


} else {
ifres13042 = False


}

var ifres13041 Obj

if True == ifres13042 {
ifres13041 = True


} else {
ifres13041 = False


}

ifres13040 = ifres13041


} else {
ifres13040 = False


}

if True == ifres13040 {
tmp13017 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3653)
}
__typedArg0 := V3653
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13018 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13017)
}
__typedArg0 := tmp13017
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13019 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13018)
}
__typedArg0 := tmp13018
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13020 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13019)
}
__typedArg0 := tmp13019
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13021 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13020)
}
__typedArg0 := tmp13020
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13022 := Call(__e, PrimFunc(symshen_4extract_1vars), tmp13021)


let__12299 := tmp13022
_ = let__12299

tmp13023 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3653)
}
__typedArg0 := V3653
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13024 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13023)
}
__typedArg0 := tmp13023
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13025 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13024)
}
__typedArg0 := tmp13024
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13026 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13025)
}
__typedArg0 := tmp13025
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13027 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13026)
}
__typedArg0 := tmp13026
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13028 := Call(__e, PrimFunc(symshen_4rule_1_6head), tmp13027)


tmp13029 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1_1, Nil)
}
__typedArg0 := sym_5_1_1
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13030 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3653)
}
__typedArg0 := V3653
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13031 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3653)
}
__typedArg0 := V3653
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13032 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13031)
}
__typedArg0 := tmp13031
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13033 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3653)
}
__typedArg0 := V3653
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13034)
}
__typedArg0 := tmp13034
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13036 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13035)
}
__typedArg0 := tmp13035
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13037 := Call(__e, PrimFunc(symshen_4rule_1_6body), let__12299, symAssumptions, tmp13030, tmp13032, tmp13036)


tmp13038 := Call(__e, PrimFunc(symappend), tmp13029, tmp13037)


__e.TailApply(PrimFunc(symappend), tmp13028, tmp13038)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4rule_1_6clause)
return
}


}, 1)

tmp13077 := Call(__e, ns2_1set, symshen_4rule_1_6clause, tmp13016)


_ = tmp13077

tmp13078 := MakeNative(func(__e *ControlFlow) {
V3655 := __e.Get(1)
_ = V3655
tmp13079 := Call(__e, PrimFunc(symshen_4macro_1_8ch), V3655)


tmp13080 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symAssumptions, Nil)
}
__typedArg0 := symAssumptions
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13079, tmp13080)
}
__typedArg0 := tmp13079
__typedArg1 := tmp13080
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp13081 := Call(__e, ns2_1set, symshen_4rule_1_6head, tmp13078)


_ = tmp13081

tmp13082 := MakeNative(func(__e *ControlFlow) {
V3656 := __e.Get(1)
_ = V3656
tmp13083 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3656, Nil)
}
__typedArg0 := V3656
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4_8ch, tmp13083)
}
__typedArg0 := symshen_4_8ch
__typedArg1 := tmp13083
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp13084 := Call(__e, ns2_1set, symshen_4macro_1_8ch, tmp13082)


_ = tmp13084

tmp13085 := MakeNative(func(__e *ControlFlow) {
V3657 := __e.Get(1)
_ = V3657
tmp13086 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3657, Nil)
}
__typedArg0 := V3657
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4_8c, tmp13086)
}
__typedArg0 := symshen_4_8c
__typedArg1 := tmp13086
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp13087 := Call(__e, ns2_1set, symshen_4macro_1_8c, tmp13085)


_ = tmp13087

tmp13088 := MakeNative(func(__e *ControlFlow) {
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
tmp13118 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3662)
}
__typedArg0 := Nil
__typedArg1 := V3662
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13118 {
__e.TailApply(PrimFunc(symshen_4side_1conditions_1_6goals), Nil, V3658, V3659, V3660, V3661)
return
} else {
tmp13116 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3661)
}
__typedArg0 := Nil
__typedArg1 := V3661
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13109 Obj

if True == tmp13116 {
tmp13115 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3662)
}
__typedArg0 := V3662
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13111 Obj

if True == tmp13115 {
tmp13113 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3662)
}
__typedArg0 := V3662
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13114 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp13113)
}
__typedArg0 := Nil
__typedArg1 := tmp13113
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13112 Obj

if True == tmp13114 {
ifres13112 = True


} else {
ifres13112 = False


}

ifres13111 = ifres13112


} else {
ifres13111 = False


}

var ifres13110 Obj

if True == ifres13111 {
ifres13110 = True


} else {
ifres13110 = False


}

ifres13109 = ifres13110


} else {
ifres13109 = False


}

if True == ifres13109 {
tmp13089 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3662)
}
__typedArg0 := V3662
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13090 := Call(__e, PrimFunc(symshen_4passive_1variables), tmp13089, V3658)


let__12300 := tmp13090
_ = let__12300

tmp13091 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3662)
}
__typedArg0 := V3662
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13092 := Call(__e, PrimFunc(symshen_4remove_1bystanders), V3658, tmp13091)


let__12301 := tmp13092
_ = let__12301

tmp13093 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3662)
}
__typedArg0 := V3662
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13094 := Call(__e, PrimFunc(symshen_4specialise_1member), tmp13093, V3659, let__12301, let__12300)


tmp13095 := Call(__e, PrimFunc(symshen_4side_1conditions_1_6goals), Nil, V3658, V3659, V3660, Nil)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13094, tmp13095)
}
__typedArg0 := tmp13094
__typedArg1 := tmp13095
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp13107 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3662)
}
__typedArg0 := V3662
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp13107 {
tmp13096 := Call(__e, PrimFunc(symgensym), symNewAssumptions)


let__12302 := tmp13096
_ = let__12302

tmp13097 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3662)
}
__typedArg0 := V3662
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13098 := Call(__e, PrimFunc(symshen_4passive_1variables), tmp13097, V3658)


let__12303 := tmp13098
_ = let__12303

tmp13099 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3662)
}
__typedArg0 := V3662
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13100 := Call(__e, PrimFunc(symshen_4remove_1bystanders), V3658, tmp13099)


let__12304 := tmp13100
_ = let__12304

tmp13101 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3662)
}
__typedArg0 := V3662
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13102 := Call(__e, PrimFunc(symshen_4specialise_1consume), tmp13101, V3659, let__12304, let__12303, let__12302)


tmp13103 := Call(__e, PrimFunc(symappend), V3658, let__12303)


tmp13104 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3662)
}
__typedArg0 := V3662
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13105 := Call(__e, PrimFunc(symshen_4rule_1_6body), tmp13103, let__12302, V3660, V3661, tmp13104)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13102, tmp13105)
}
__typedArg0 := tmp13102
__typedArg1 := tmp13105
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

tmp13119 := Call(__e, ns2_1set, symshen_4rule_1_6body, tmp13088)


_ = tmp13119

tmp13120 := MakeNative(func(__e *ControlFlow) {
V3668 := __e.Get(1)
_ = V3668
V3669 := __e.Get(2)
_ = V3669
V3670 := __e.Get(3)
_ = V3670
V3671 := __e.Get(4)
_ = V3671
tmp13121 := Call(__e, PrimFunc(symgensym), symshen_4member)


let__12305 := tmp13121
_ = let__12305

tmp13122 := Call(__e, PrimFunc(symshen_4member_1clause), let__12305, V3668, V3670, V3671)


let__12306 := tmp13122
_ = let__12306

tmp13123 := Call(__e, PrimFunc(symappend), V3670, V3671)


tmp13124 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3669, tmp13123)
}
__typedArg0 := V3669
__typedArg1 := tmp13123
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12305, tmp13124)
}
__typedArg0 := let__12305
__typedArg1 := tmp13124
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 4)

tmp13125 := Call(__e, ns2_1set, symshen_4specialise_1member, tmp13120)


_ = tmp13125

tmp13126 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__selftop:
V3676 := __self1
_ = V3676
V3677 := __self2
_ = V3677
tmp13140 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3676)
}
__typedArg0 := Nil
__typedArg1 := V3676
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13140 {
__e.Return(Nil)
return
} else {
tmp13138 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3676)
}
__typedArg0 := V3676
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13134 Obj

if True == tmp13138 {
tmp13136 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3676)
}
__typedArg0 := V3676
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13137 := Call(__e, PrimFunc(symshen_4occurs_1check_2), tmp13136, V3677)


var ifres13135 Obj

if True == tmp13137 {
ifres13135 = True


} else {
ifres13135 = False


}

ifres13134 = ifres13135


} else {
ifres13134 = False


}

if True == ifres13134 {
tmp13127 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3676)
}
__typedArg0 := V3676
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13128 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3676)
}
__typedArg0 := V3676
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13129 := Call(__e, PrimFunc(symshen_4remove_1bystanders), tmp13128, V3677)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13127, tmp13129)
}
__typedArg0 := tmp13127
__typedArg1 := tmp13129
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp13132 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3676)
}
__typedArg0 := V3676
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp13132 {
tmp13130 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3676)
}
__typedArg0 := V3676
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4remove_1bystanders) == __self {
__self1, __self2 = tmp13130, V3677
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4remove_1bystanders), tmp13130, V3677)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4remove_1bystanders)
return
}


}


}


}, 2)

tmp13141 := Call(__e, ns2_1set, symshen_4remove_1bystanders, tmp13126)


_ = tmp13141

tmp13142 := MakeNative(func(__e *ControlFlow) {
V3678 := __e.Get(1)
_ = V3678
V3679 := __e.Get(2)
_ = V3679
V3680 := __e.Get(3)
_ = V3680
V3681 := __e.Get(4)
_ = V3681
tmp13143 := Call(__e, PrimFunc(symlength), V3681)


tmp13144 := Call(__e, PrimFunc(symshen_4nvars), tmp13143)


let__12307 := tmp13144
_ = let__12307

tmp13145 := Call(__e, PrimFunc(symshen_4macro_1_8ch), V3679)


tmp13146 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym__, Nil)
}
__typedArg0 := sym__
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13147 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13145, tmp13146)
}
__typedArg0 := tmp13145
__typedArg1 := tmp13146
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13148 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp13147)
}
__typedArg0 := symcons
__typedArg1 := tmp13147
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13149 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13148, Nil)
}
__typedArg0 := tmp13148
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13150 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1, tmp13149)
}
__typedArg0 := sym_1
__typedArg1 := tmp13149
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13151 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13150, Nil)
}
__typedArg0 := tmp13150
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13152 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1_1, Nil)
}
__typedArg0 := sym_5_1_1
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13153 := Call(__e, PrimFunc(symshen_4passive_1bind), V3681, let__12307)


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

tmp13156 := Call(__e, PrimFunc(symappend), tmp13153, tmp13155)


tmp13157 := Call(__e, PrimFunc(symappend), tmp13152, tmp13156)


tmp13158 := Call(__e, PrimFunc(symappend), let__12307, tmp13157)


tmp13159 := Call(__e, PrimFunc(symappend), V3680, tmp13158)


tmp13160 := Call(__e, PrimFunc(symappend), tmp13151, tmp13159)


let__12308 := tmp13160
_ = let__12308

tmp13161 := Call(__e, PrimFunc(symgensym), symHypotheses)


let__12310 := tmp13161
_ = let__12310

tmp13162 := Call(__e, PrimFunc(symappend), V3680, V3681)


let__12311 := tmp13162
_ = let__12311

tmp13163 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12310, Nil)
}
__typedArg0 := let__12310
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13164 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym__, tmp13163)
}
__typedArg0 := sym__
__typedArg1 := tmp13163
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13165 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp13164)
}
__typedArg0 := symcons
__typedArg1 := tmp13164
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
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
return PrimCons(sym_1, tmp13166)
}
__typedArg0 := sym_1
__typedArg1 := tmp13166
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13168 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13167, Nil)
}
__typedArg0 := tmp13167
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13169 := Call(__e, PrimFunc(symappend), tmp13168, let__12311)


let__12312 := tmp13169
_ = let__12312

tmp13170 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12310, let__12311)
}
__typedArg0 := let__12310
__typedArg1 := let__12311
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13171 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3678, tmp13170)
}
__typedArg0 := V3678
__typedArg1 := tmp13170
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13172 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13171, Nil)
}
__typedArg0 := tmp13171
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__12313 := tmp13172
_ = let__12313

tmp13173 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1_1, Nil)
}
__typedArg0 := sym_5_1_1
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13174 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(";"))
}
__typedArg0 := MakeString(";")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp13175 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13174, Nil)
}
__typedArg0 := tmp13174
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13176 := Call(__e, PrimFunc(symappend), let__12313, tmp13175)


tmp13177 := Call(__e, PrimFunc(symappend), tmp13173, tmp13176)


tmp13178 := Call(__e, PrimFunc(symappend), let__12312, tmp13177)


let__12309 := tmp13178
_ = let__12309

tmp13179 := Call(__e, PrimFunc(symappend), let__12308, let__12309)


tmp13180 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3678, tmp13179)
}
__typedArg0 := V3678
__typedArg1 := tmp13179
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13181 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefprolog, tmp13180)
}
__typedArg0 := symdefprolog
__typedArg1 := tmp13180
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__12314 := tmp13181
_ = let__12314

__e.TailApply(PrimFunc(symeval), let__12314)
return


}, 4)

tmp13182 := Call(__e, ns2_1set, symshen_4member_1clause, tmp13142)


_ = tmp13182

tmp13183 := MakeNative(func(__e *ControlFlow) {
V3690 := __e.Get(1)
_ = V3690
tmp13188 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(0), V3690)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := V3690
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13188 {
__e.Return(Nil)
return
} else {
tmp13184 := Call(__e, PrimFunc(symgensym), symNewV)


tmp13186 := Call(__e, PrimFunc(symshen_4nvars), (func() Obj {
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
return PrimCons(tmp13184, tmp13186)
}
__typedArg0 := tmp13184
__typedArg1 := tmp13186
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}


}, 1)

tmp13189 := Call(__e, ns2_1set, symshen_4nvars, tmp13183)


_ = tmp13189

tmp13190 := MakeNative(func(__e *ControlFlow) {
V3691 := __e.Get(1)
_ = V3691
V3692 := __e.Get(2)
_ = V3692
tmp13208 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3691)
}
__typedArg0 := Nil
__typedArg1 := V3691
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13205 Obj

if True == tmp13208 {
tmp13207 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3692)
}
__typedArg0 := Nil
__typedArg1 := V3692
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13206 Obj

if True == tmp13207 {
ifres13206 = True


} else {
ifres13206 = False


}

ifres13205 = ifres13206


} else {
ifres13205 = False


}

if True == ifres13205 {
__e.Return(Nil)
return
} else {
tmp13203 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3691)
}
__typedArg0 := V3691
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13200 Obj

if True == tmp13203 {
tmp13202 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3692)
}
__typedArg0 := V3692
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13201 Obj

if True == tmp13202 {
ifres13201 = True


} else {
ifres13201 = False


}

ifres13200 = ifres13201


} else {
ifres13200 = False


}

if True == ifres13200 {
tmp13191 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3692)
}
__typedArg0 := V3692
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13192 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3691)
}
__typedArg0 := V3691
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13193 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13192, Nil)
}
__typedArg0 := tmp13192
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13194 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13191, tmp13193)
}
__typedArg0 := tmp13191
__typedArg1 := tmp13193
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13195 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symbind, tmp13194)
}
__typedArg0 := symbind
__typedArg1 := tmp13194
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13196 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3691)
}
__typedArg0 := V3691
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13197 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3692)
}
__typedArg0 := V3692
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13198 := Call(__e, PrimFunc(symshen_4passive_1bind), tmp13196, tmp13197)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13195, tmp13198)
}
__typedArg0 := tmp13195
__typedArg1 := tmp13198
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4passive_1bind)
return
}


}


}, 2)

tmp13209 := Call(__e, ns2_1set, symshen_4passive_1bind, tmp13190)


_ = tmp13209

tmp13210 := MakeNative(func(__e *ControlFlow) {
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
tmp13211 := Call(__e, PrimFunc(symgensym), symshen_4consume)


let__12315 := tmp13211
_ = let__12315

tmp13212 := Call(__e, PrimFunc(symshen_4consume_1clause), let__12315, V3693, V3695, V3696, V3697)


let__12316 := tmp13212
_ = let__12316

tmp13213 := Call(__e, PrimFunc(symappend), V3695, V3696)


tmp13214 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3697, tmp13213)
}
__typedArg0 := V3697
__typedArg1 := tmp13213
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13215 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3694, tmp13214)
}
__typedArg0 := V3694
__typedArg1 := tmp13214
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12315, tmp13215)
}
__typedArg0 := let__12315
__typedArg1 := tmp13215
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 5)

tmp13216 := Call(__e, ns2_1set, symshen_4specialise_1consume, tmp13210)


_ = tmp13216

tmp13217 := MakeNative(func(__e *ControlFlow) {
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
tmp13218 := Call(__e, PrimFunc(symlength), V3703)


tmp13219 := Call(__e, PrimFunc(symshen_4nvars), tmp13218)


let__12317 := tmp13219
_ = let__12317

tmp13220 := Call(__e, PrimFunc(symgensym), symAssumption)


let__12318 := tmp13220
_ = let__12318

tmp13221 := Call(__e, PrimFunc(symshen_4macro_1_8ch), V3701)


tmp13222 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12318, Nil)
}
__typedArg0 := let__12318
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13223 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13221, tmp13222)
}
__typedArg0 := tmp13221
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
return PrimCons(tmp13224, Nil)
}
__typedArg0 := tmp13224
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13226 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1, tmp13225)
}
__typedArg0 := sym_1
__typedArg1 := tmp13225
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13227 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1_1, Nil)
}
__typedArg0 := sym_5_1_1
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13228 := Call(__e, PrimFunc(symshen_4passive_1bind), V3703, let__12317)


tmp13229 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12318, Nil)
}
__typedArg0 := let__12318
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13230 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3704, tmp13229)
}
__typedArg0 := V3704
__typedArg1 := tmp13229
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13231 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symbind, tmp13230)
}
__typedArg0 := symbind
__typedArg1 := tmp13230
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13232 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(";"))
}
__typedArg0 := MakeString(";")
return Call(__e, PrimFunc(symintern), __typedArg0)
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
return PrimCons(tmp13231, tmp13233)
}
__typedArg0 := tmp13231
__typedArg1 := tmp13233
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13235 := Call(__e, PrimFunc(symappend), tmp13228, tmp13234)


tmp13236 := Call(__e, PrimFunc(symappend), tmp13227, tmp13235)


tmp13237 := Call(__e, PrimFunc(symappend), let__12317, tmp13236)


tmp13238 := Call(__e, PrimFunc(symappend), V3702, tmp13237)


tmp13239 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3704, tmp13238)
}
__typedArg0 := V3704
__typedArg1 := tmp13238
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13240 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13226, tmp13239)
}
__typedArg0 := tmp13226
__typedArg1 := tmp13239
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__12319 := tmp13240
_ = let__12319

tmp13241 := Call(__e, PrimFunc(symgensym), symHypotheses)


let__12321 := tmp13241
_ = let__12321

tmp13242 := Call(__e, PrimFunc(symappend), V3702, V3703)


let__12322 := tmp13242
_ = let__12322

tmp13243 := Call(__e, PrimFunc(symgensym), symAssumptions)


let__12323 := tmp13243
_ = let__12323

tmp13244 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12321, Nil)
}
__typedArg0 := let__12321
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13245 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12318, tmp13244)
}
__typedArg0 := let__12318
__typedArg1 := tmp13244
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13246 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp13245)
}
__typedArg0 := symcons
__typedArg1 := tmp13245
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13247 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13246, Nil)
}
__typedArg0 := tmp13246
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13248 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1, tmp13247)
}
__typedArg0 := sym_1
__typedArg1 := tmp13247
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13249 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3704, Nil)
}
__typedArg0 := V3704
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13250 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12323, tmp13249)
}
__typedArg0 := let__12323
__typedArg1 := tmp13249
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13251 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp13250)
}
__typedArg0 := symcons
__typedArg1 := tmp13250
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13252 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13251, let__12322)
}
__typedArg0 := tmp13251
__typedArg1 := let__12322
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13253 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13248, tmp13252)
}
__typedArg0 := tmp13248
__typedArg1 := tmp13252
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__12324 := tmp13253
_ = let__12324

tmp13254 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12318, Nil)
}
__typedArg0 := let__12318
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13255 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12323, tmp13254)
}
__typedArg0 := let__12323
__typedArg1 := tmp13254
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13256 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symbind, tmp13255)
}
__typedArg0 := symbind
__typedArg1 := tmp13255
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13257 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3704, let__12322)
}
__typedArg0 := V3704
__typedArg1 := let__12322
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13258 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__12321, tmp13257)
}
__typedArg0 := let__12321
__typedArg1 := tmp13257
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13259 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3700, tmp13258)
}
__typedArg0 := V3700
__typedArg1 := tmp13258
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13260 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13259, Nil)
}
__typedArg0 := tmp13259
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13261 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13256, tmp13260)
}
__typedArg0 := tmp13256
__typedArg1 := tmp13260
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__12325 := tmp13261
_ = let__12325

tmp13262 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1_1, Nil)
}
__typedArg0 := sym_5_1_1
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13263 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(";"))
}
__typedArg0 := MakeString(";")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp13264 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13263, Nil)
}
__typedArg0 := tmp13263
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13265 := Call(__e, PrimFunc(symappend), let__12325, tmp13264)


tmp13266 := Call(__e, PrimFunc(symappend), tmp13262, tmp13265)


tmp13267 := Call(__e, PrimFunc(symappend), let__12324, tmp13266)


let__12320 := tmp13267
_ = let__12320

tmp13268 := Call(__e, PrimFunc(symappend), let__12319, let__12320)


tmp13269 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3700, tmp13268)
}
__typedArg0 := V3700
__typedArg1 := tmp13268
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13270 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefprolog, tmp13269)
}
__typedArg0 := symdefprolog
__typedArg1 := tmp13269
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__12326 := tmp13270
_ = let__12326

__e.TailApply(PrimFunc(symeval), let__12326)
return


}, 5)

tmp13271 := Call(__e, ns2_1set, symshen_4consume_1clause, tmp13217)


_ = tmp13271

tmp13272 := MakeNative(func(__e *ControlFlow) {
V3715 := __e.Get(1)
_ = V3715
V3716 := __e.Get(2)
_ = V3716
tmp13273 := Call(__e, PrimFunc(symshen_4extract_1vars), V3715)


__e.TailApply(PrimFunc(symdifference), tmp13273, V3716)
return


}, 2)

tmp13274 := Call(__e, ns2_1set, symshen_4passive_1variables, tmp13272)


_ = tmp13274

tmp13275 := MakeNative(func(__e *ControlFlow) {
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
tmp13426 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3724)
}
__typedArg0 := Nil
__typedArg1 := V3724
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13426 {
__e.TailApply(PrimFunc(symshen_4premises_1_6goals), V3721, V3723, V3725)
return
} else {
tmp13424 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3724)
}
__typedArg0 := V3724
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13404 Obj

if True == tmp13424 {
tmp13422 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3724)
}
__typedArg0 := V3724
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13423 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13422)
}
__typedArg0 := tmp13422
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13406 Obj

if True == tmp13423 {
tmp13419 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3724)
}
__typedArg0 := V3724
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13420 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13419)
}
__typedArg0 := tmp13419
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13421 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symif, tmp13420)
}
__typedArg0 := symif
__typedArg1 := tmp13420
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13408 Obj

if True == tmp13421 {
tmp13416 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3724)
}
__typedArg0 := V3724
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13417)
}
__typedArg0 := tmp13417
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13410 Obj

if True == tmp13418 {
tmp13412 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3724)
}
__typedArg0 := V3724
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13413 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13412)
}
__typedArg0 := tmp13412
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13414 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13413)
}
__typedArg0 := tmp13413
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13415 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp13414)
}
__typedArg0 := Nil
__typedArg1 := tmp13414
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13411 Obj

if True == tmp13415 {
ifres13411 = True


} else {
ifres13411 = False


}

ifres13410 = ifres13411


} else {
ifres13410 = False


}

var ifres13409 Obj

if True == ifres13410 {
ifres13409 = True


} else {
ifres13409 = False


}

ifres13408 = ifres13409


} else {
ifres13408 = False


}

var ifres13407 Obj

if True == ifres13408 {
ifres13407 = True


} else {
ifres13407 = False


}

ifres13406 = ifres13407


} else {
ifres13406 = False


}

var ifres13405 Obj

if True == ifres13406 {
ifres13405 = True


} else {
ifres13405 = False


}

ifres13404 = ifres13405


} else {
ifres13404 = False


}

if True == ifres13404 {
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhen, tmp13277)
}
__typedArg0 := symwhen
__typedArg1 := tmp13277
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13279 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3724)
}
__typedArg0 := V3724
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13280 := Call(__e, PrimFunc(symshen_4side_1conditions_1_6goals), V3721, V3722, V3723, tmp13279, V3725)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13278, tmp13280)
}
__typedArg0 := tmp13278
__typedArg1 := tmp13280
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp13402 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3724)
}
__typedArg0 := V3724
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13375 Obj

if True == tmp13402 {
tmp13400 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3724)
}
__typedArg0 := V3724
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13401 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13400)
}
__typedArg0 := tmp13400
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13377 Obj

if True == tmp13401 {
tmp13397 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3724)
}
__typedArg0 := V3724
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13398 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13397)
}
__typedArg0 := tmp13397
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13399 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlet, tmp13398)
}
__typedArg0 := symlet
__typedArg1 := tmp13398
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13379 Obj

if True == tmp13399 {
tmp13394 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3724)
}
__typedArg0 := V3724
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13395 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13394)
}
__typedArg0 := tmp13394
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13396 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13395)
}
__typedArg0 := tmp13395
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13381 Obj

if True == tmp13396 {
tmp13390 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3724)
}
__typedArg0 := V3724
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13391 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13390)
}
__typedArg0 := tmp13390
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13392 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13391)
}
__typedArg0 := tmp13391
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13393 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13392)
}
__typedArg0 := tmp13392
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13383 Obj

if True == tmp13393 {
tmp13385 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3724)
}
__typedArg0 := V3724
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13387)
}
__typedArg0 := tmp13387
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13389 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp13388)
}
__typedArg0 := Nil
__typedArg1 := tmp13388
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13384 Obj

if True == tmp13389 {
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

var ifres13376 Obj

if True == ifres13377 {
ifres13376 = True


} else {
ifres13376 = False


}

ifres13375 = ifres13376


} else {
ifres13375 = False


}

if True == ifres13375 {
tmp13296 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3724)
}
__typedArg0 := V3724
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13297 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13296)
}
__typedArg0 := tmp13296
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13298 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13297)
}
__typedArg0 := tmp13297
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13299 := Call(__e, PrimFunc(symelement_2), tmp13298, V3722)


if True == tmp13299 {
tmp13281 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3724)
}
__typedArg0 := V3724
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13282 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13281)
}
__typedArg0 := tmp13281
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13283 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symis_b, tmp13282)
}
__typedArg0 := symis_b
__typedArg1 := tmp13282
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13284 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3724)
}
__typedArg0 := V3724
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13285 := Call(__e, PrimFunc(symshen_4side_1conditions_1_6goals), V3721, V3722, V3723, tmp13284, V3725)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13283, tmp13285)
}
__typedArg0 := tmp13283
__typedArg1 := tmp13285
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp13286 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3724)
}
__typedArg0 := V3724
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13287 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13286)
}
__typedArg0 := tmp13286
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13288 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symbind, tmp13287)
}
__typedArg0 := symbind
__typedArg1 := tmp13287
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13289 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3724)
}
__typedArg0 := V3724
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13290 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13289)
}
__typedArg0 := tmp13289
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13291 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13290)
}
__typedArg0 := tmp13290
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13292 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13291, V3722)
}
__typedArg0 := tmp13291
__typedArg1 := V3722
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13293 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3724)
}
__typedArg0 := V3724
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13294 := Call(__e, PrimFunc(symshen_4side_1conditions_1_6goals), V3721, tmp13292, V3723, tmp13293, V3725)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13288, tmp13294)
}
__typedArg0 := tmp13288
__typedArg1 := tmp13294
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}


} else {
tmp13373 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3724)
}
__typedArg0 := V3724
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13353 Obj

if True == tmp13373 {
tmp13371 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3724)
}
__typedArg0 := V3724
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13372 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13371)
}
__typedArg0 := tmp13371
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13355 Obj

if True == tmp13372 {
tmp13368 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3724)
}
__typedArg0 := V3724
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13369 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13368)
}
__typedArg0 := tmp13368
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13370 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symctxt, tmp13369)
}
__typedArg0 := symctxt
__typedArg1 := tmp13369
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13357 Obj

if True == tmp13370 {
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13366)
}
__typedArg0 := tmp13366
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13359 Obj

if True == tmp13367 {
tmp13361 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3724)
}
__typedArg0 := V3724
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13362 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13361)
}
__typedArg0 := tmp13361
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13363 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13362)
}
__typedArg0 := tmp13362
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13364 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp13363)
}
__typedArg0 := Nil
__typedArg1 := tmp13363
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13360 Obj

if True == tmp13364 {
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

var ifres13354 Obj

if True == ifres13355 {
ifres13354 = True


} else {
ifres13354 = False


}

ifres13353 = ifres13354


} else {
ifres13353 = False


}

if True == ifres13353 {
tmp13325 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3724)
}
__typedArg0 := V3724
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13326 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13325)
}
__typedArg0 := tmp13325
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13327 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13326)
}
__typedArg0 := tmp13326
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13328 := Call(__e, PrimFunc(symelement_2), tmp13327, V3722)


if True == tmp13328 {
tmp13300 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3724)
}
__typedArg0 := V3724
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13301 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13300)
}
__typedArg0 := tmp13300
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13302 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13301)
}
__typedArg0 := tmp13301
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13303 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13302, V3721)
}
__typedArg0 := tmp13302
__typedArg1 := V3721
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13304 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3724)
}
__typedArg0 := V3724
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4side_1conditions_1_6goals) == __self {
__self1, __self2, __self3, __self4, __self5 = tmp13303, V3722, V3723, tmp13304, V3725
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4side_1conditions_1_6goals), tmp13303, V3722, V3723, tmp13304, V3725)
return


} else {
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

tmp13308 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3723, Nil)
}
__typedArg0 := V3723
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13309 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13307, tmp13308)
}
__typedArg0 := tmp13307
__typedArg1 := tmp13308
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13310 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symbind, tmp13309)
}
__typedArg0 := symbind
__typedArg1 := tmp13309
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13311 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3724)
}
__typedArg0 := V3724
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13312 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13311)
}
__typedArg0 := tmp13311
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13313 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13312)
}
__typedArg0 := tmp13312
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13314 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13313, V3721)
}
__typedArg0 := tmp13313
__typedArg1 := V3721
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13315 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3724)
}
__typedArg0 := V3724
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13316 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13315)
}
__typedArg0 := tmp13315
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13317 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13316)
}
__typedArg0 := tmp13316
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13318 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13317, V3722)
}
__typedArg0 := tmp13317
__typedArg1 := V3722
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13320)
}
__typedArg0 := tmp13320
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13322 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3724)
}
__typedArg0 := V3724
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13323 := Call(__e, PrimFunc(symshen_4side_1conditions_1_6goals), tmp13314, tmp13318, tmp13321, tmp13322, V3725)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13310, tmp13323)
}
__typedArg0 := tmp13310
__typedArg1 := tmp13323
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}


} else {
tmp13351 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3724)
}
__typedArg0 := V3724
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13331 Obj

if True == tmp13351 {
tmp13349 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3724)
}
__typedArg0 := V3724
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13350 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13349)
}
__typedArg0 := tmp13349
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13333 Obj

if True == tmp13350 {
tmp13346 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3724)
}
__typedArg0 := V3724
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13347 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13346)
}
__typedArg0 := tmp13346
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13348 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symsqts, tmp13347)
}
__typedArg0 := symsqts
__typedArg1 := tmp13347
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13335 Obj

if True == tmp13348 {
tmp13343 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3724)
}
__typedArg0 := V3724
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13344 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13343)
}
__typedArg0 := tmp13343
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13345 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13344)
}
__typedArg0 := tmp13344
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13337 Obj

if True == tmp13345 {
tmp13339 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3724)
}
__typedArg0 := V3724
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13340 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13339)
}
__typedArg0 := tmp13339
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13341 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13340)
}
__typedArg0 := tmp13340
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13342 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp13341)
}
__typedArg0 := Nil
__typedArg1 := tmp13341
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13338 Obj

if True == tmp13342 {
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

var ifres13332 Obj

if True == ifres13333 {
ifres13332 = True


} else {
ifres13332 = False


}

ifres13331 = ifres13332


} else {
ifres13331 = False


}

if True == ifres13331 {
tmp13329 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3724)
}
__typedArg0 := V3724
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4side_1conditions_1_6goals) == __self {
__self1, __self2, __self3, __self4, __self5 = V3721, V3722, V3723, tmp13329, V3725
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4side_1conditions_1_6goals), V3721, V3722, V3723, tmp13329, V3725)
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

tmp13427 := Call(__e, ns2_1set, symshen_4side_1conditions_1_6goals, tmp13275)


_ = tmp13427

tmp13428 := MakeNative(func(__e *ControlFlow) {
V3730 := __e.Get(1)
_ = V3730
V3731 := __e.Get(2)
_ = V3731
V3732 := __e.Get(3)
_ = V3732
tmp13478 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3732)
}
__typedArg0 := Nil
__typedArg1 := V3732
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13478 {
tmp13429 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(";"))
}
__typedArg0 := MakeString(";")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13429, Nil)
}
__typedArg0 := tmp13429
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp13476 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3732)
}
__typedArg0 := V3732
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13472 Obj

if True == tmp13476 {
tmp13474 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3732)
}
__typedArg0 := V3732
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13475 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_b, tmp13474)
}
__typedArg0 := sym_b
__typedArg1 := tmp13474
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

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

if True == ifres13472 {
tmp13430 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3732)
}
__typedArg0 := V3732
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13431 := Call(__e, PrimFunc(symshen_4premises_1_6goals), V3730, V3731, tmp13430)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_b, tmp13431)
}
__typedArg0 := sym_b
__typedArg1 := tmp13431
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp13470 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3732)
}
__typedArg0 := V3732
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13466 Obj

if True == tmp13470 {
tmp13468 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3732)
}
__typedArg0 := V3732
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13469 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symfail, tmp13468)
}
__typedArg0 := symfail
__typedArg1 := tmp13468
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13467 Obj

if True == tmp13469 {
ifres13467 = True


} else {
ifres13467 = False


}

ifres13466 = ifres13467


} else {
ifres13466 = False


}

if True == ifres13466 {
tmp13432 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(False, Nil)
}
__typedArg0 := False
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13433 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhen, tmp13432)
}
__typedArg0 := symwhen
__typedArg1 := tmp13432
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13434 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3732)
}
__typedArg0 := V3732
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13435 := Call(__e, PrimFunc(symshen_4premises_1_6goals), V3730, V3731, tmp13434)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13433, tmp13435)
}
__typedArg0 := tmp13433
__typedArg1 := tmp13435
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp13464 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3732)
}
__typedArg0 := V3732
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13449 Obj

if True == tmp13464 {
tmp13462 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3732)
}
__typedArg0 := V3732
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13463 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13462)
}
__typedArg0 := tmp13462
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13451 Obj

if True == tmp13463 {
tmp13459 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3732)
}
__typedArg0 := V3732
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13460 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13459)
}
__typedArg0 := tmp13459
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13461 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13460)
}
__typedArg0 := tmp13460
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13453 Obj

if True == tmp13461 {
tmp13455 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3732)
}
__typedArg0 := V3732
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13456 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13455)
}
__typedArg0 := tmp13455
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13457 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13456)
}
__typedArg0 := tmp13456
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13458 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp13457)
}
__typedArg0 := Nil
__typedArg1 := tmp13457
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13454 Obj

if True == tmp13458 {
ifres13454 = True


} else {
ifres13454 = False


}

ifres13453 = ifres13454


} else {
ifres13453 = False


}

var ifres13452 Obj

if True == ifres13453 {
ifres13452 = True


} else {
ifres13452 = False


}

ifres13451 = ifres13452


} else {
ifres13451 = False


}

var ifres13450 Obj

if True == ifres13451 {
ifres13450 = True


} else {
ifres13450 = False


}

ifres13449 = ifres13450


} else {
ifres13449 = False


}

if True == ifres13449 {
tmp13436 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3732)
}
__typedArg0 := V3732
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13437 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13436)
}
__typedArg0 := tmp13436
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13438 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13437)
}
__typedArg0 := tmp13437
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13439 := Call(__e, PrimFunc(symshen_4macro_1_8c), tmp13438)


tmp13440 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3732)
}
__typedArg0 := V3732
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13441 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13440)
}
__typedArg0 := tmp13440
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13442 := Call(__e, PrimFunc(symshen_4construct_1context), V3730, tmp13441, V3731)


tmp13443 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13442, Nil)
}
__typedArg0 := tmp13442
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13444 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13439, tmp13443)
}
__typedArg0 := tmp13439
__typedArg1 := tmp13443
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13445 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4system_1S, tmp13444)
}
__typedArg0 := symshen_4system_1S
__typedArg1 := tmp13444
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13446 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3732)
}
__typedArg0 := V3732
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13447 := Call(__e, PrimFunc(symshen_4premises_1_6goals), V3730, V3731, tmp13446)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13445, tmp13447)
}
__typedArg0 := tmp13445
__typedArg1 := tmp13447
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

tmp13479 := Call(__e, ns2_1set, symshen_4premises_1_6goals, tmp13428)


_ = tmp13479

tmp13480 := MakeNative(func(__e *ControlFlow) {
V3736 := __e.Get(1)
_ = V3736
V3737 := __e.Get(2)
_ = V3737
V3738 := __e.Get(3)
_ = V3738
tmp13500 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3737)
}
__typedArg0 := Nil
__typedArg1 := V3737
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13500 {
__e.Return(V3738)
return
} else {
tmp13498 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3737)
}
__typedArg0 := V3737
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13490 Obj

if True == tmp13498 {
tmp13496 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3737)
}
__typedArg0 := V3737
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13497 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp13496)
}
__typedArg0 := Nil
__typedArg1 := tmp13496
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13492 Obj

if True == tmp13497 {
tmp13494 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3737)
}
__typedArg0 := V3737
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13495 := Call(__e, PrimFunc(symelement_2), tmp13494, V3736)


var ifres13493 Obj

if True == tmp13495 {
ifres13493 = True


} else {
ifres13493 = False


}

ifres13492 = ifres13493


} else {
ifres13492 = False


}

var ifres13491 Obj

if True == ifres13492 {
ifres13491 = True


} else {
ifres13491 = False


}

ifres13490 = ifres13491


} else {
ifres13490 = False


}

if True == ifres13490 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3737)
}
__typedArg0 := V3737
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return
} else {
tmp13488 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3737)
}
__typedArg0 := V3737
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp13488 {
tmp13481 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3737)
}
__typedArg0 := V3737
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13482 := Call(__e, PrimFunc(symshen_4macro_1_8c), tmp13481)


tmp13483 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3737)
}
__typedArg0 := V3737
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13484 := Call(__e, PrimFunc(symshen_4construct_1context), V3736, tmp13483, V3738)


tmp13485 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13484, Nil)
}
__typedArg0 := tmp13484
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13486 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13482, tmp13485)
}
__typedArg0 := tmp13482
__typedArg1 := tmp13485
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp13486)
}
__typedArg0 := symcons
__typedArg1 := tmp13486
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

tmp13501 := Call(__e, ns2_1set, symshen_4construct_1context, tmp13480)


_ = tmp13501

tmp13502 := MakeNative(func(__e *ControlFlow) {
V3739 := __e.Get(1)
_ = V3739
tmp13503 := MakeNative(func(__e *ControlFlow) {
Z3741 := __e.Get(1)
_ = Z3741
__e.TailApply(PrimFunc(symshen_4intern_1type), Z3741)
return
}, 1)

tmp13504 := Call(__e, PrimFunc(symmap), tmp13503, V3739)


let__12327 := tmp13504
_ = let__12327

tmp13505 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_ddatatypes_d)
}
__typedArg0 := symshen_4_ddatatypes_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

let__12328 := tmp13505
_ = let__12328

tmp13506 := Call(__e, PrimFunc(symshen_4remove_1datatypes), let__12327, let__12328)


let__12329 := tmp13506
_ = let__12329

tmp13507 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_ddatatypes_d, let__12329)
}
__typedArg0 := symshen_4_ddatatypes_d
__typedArg1 := let__12329
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

let__12330 := tmp13507
_ = let__12330

__e.TailApply(PrimFunc(symshen_4show_1datatypes), let__12330)
return


}, 1)

tmp13508 := Call(__e, ns2_1set, sympreclude, tmp13502)


_ = tmp13508

tmp13509 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__selftop:
V3749 := __self1
_ = V3749
V3750 := __self2
_ = V3750
tmp13516 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3749)
}
__typedArg0 := Nil
__typedArg1 := V3749
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13516 {
__e.Return(V3750)
return
} else {
tmp13514 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3749)
}
__typedArg0 := V3749
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp13514 {
tmp13510 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3749)
}
__typedArg0 := V3749
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13511 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3749)
}
__typedArg0 := V3749
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13512 := Call(__e, PrimFunc(symshen_4unassoc), tmp13511, V3750)


if PrimFunc(symshen_4remove_1datatypes) == __self {
__self1, __self2 = tmp13510, tmp13512
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4remove_1datatypes), tmp13510, tmp13512)
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

tmp13517 := Call(__e, ns2_1set, symshen_4remove_1datatypes, tmp13509)


_ = tmp13517

tmp13518 := MakeNative(func(__e *ControlFlow) {
V3751 := __e.Get(1)
_ = V3751
tmp13519 := MakeNative(func(__e *ControlFlow) {
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

__e.TailApply(PrimFunc(symmap), tmp13519, V3751)
return


}, 1)

tmp13520 := Call(__e, ns2_1set, symshen_4show_1datatypes, tmp13518)


_ = tmp13520

tmp13521 := MakeNative(func(__e *ControlFlow) {
V3753 := __e.Get(1)
_ = V3753
tmp13522 := MakeNative(func(__e *ControlFlow) {
Z3755 := __e.Get(1)
_ = Z3755
__e.TailApply(PrimFunc(symshen_4intern_1type), Z3755)
return
}, 1)

tmp13523 := Call(__e, PrimFunc(symmap), tmp13522, V3753)


let__12331 := tmp13523
_ = let__12331

tmp13524 := MakeNative(func(__e *ControlFlow) {
Z3757 := __e.Get(1)
_ = Z3757
tmp13525 := Call(__e, PrimFunc(symfn), Z3757)


__e.TailApply(PrimFunc(symshen_4remember_1datatype), Z3757, tmp13525)
return


}, 1)

tmp13526 := Call(__e, PrimFunc(symmap), tmp13524, let__12331)


let__12332 := tmp13526
_ = let__12332

tmp13527 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_ddatatypes_d)
}
__typedArg0 := symshen_4_ddatatypes_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

let__12333 := tmp13527
_ = let__12333

__e.TailApply(PrimFunc(symshen_4show_1datatypes), let__12333)
return


}, 1)

tmp13528 := Call(__e, ns2_1set, syminclude, tmp13521)


_ = tmp13528

tmp13529 := MakeNative(func(__e *ControlFlow) {
V3759 := __e.Get(1)
_ = V3759
tmp13530 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_ddatatypes_d, Nil)
}
__typedArg0 := symshen_4_ddatatypes_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

let__12334 := tmp13530
_ = let__12334

tmp13531 := MakeNative(func(__e *ControlFlow) {
Z3762 := __e.Get(1)
_ = Z3762
__e.TailApply(PrimFunc(symshen_4intern_1type), Z3762)
return
}, 1)

tmp13532 := Call(__e, PrimFunc(symmap), tmp13531, V3759)


let__12335 := tmp13532
_ = let__12335

tmp13533 := MakeNative(func(__e *ControlFlow) {
Z3764 := __e.Get(1)
_ = Z3764
tmp13534 := Call(__e, PrimFunc(symfn), Z3764)


__e.TailApply(PrimFunc(symshen_4remember_1datatype), Z3764, tmp13534)
return


}, 1)

tmp13535 := Call(__e, PrimFunc(symmap), tmp13533, let__12335)


let__12336 := tmp13535
_ = let__12336

tmp13536 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_ddatatypes_d)
}
__typedArg0 := symshen_4_ddatatypes_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4show_1datatypes), tmp13536)
return


}, 1)

tmp13537 := Call(__e, ns2_1set, sympreclude_1all_1but, tmp13529)


_ = tmp13537

tmp13538 := MakeNative(func(__e *ControlFlow) {
V3765 := __e.Get(1)
_ = V3765
tmp13539 := MakeNative(func(__e *ControlFlow) {
Z3767 := __e.Get(1)
_ = Z3767
__e.TailApply(PrimFunc(symshen_4intern_1type), Z3767)
return
}, 1)

tmp13540 := Call(__e, PrimFunc(symmap), tmp13539, V3765)


let__12337 := tmp13540
_ = let__12337

tmp13541 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dalldatatypes_d)
}
__typedArg0 := symshen_4_dalldatatypes_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

let__12338 := tmp13541
_ = let__12338

tmp13542 := Call(__e, PrimFunc(symshen_4remove_1datatypes), let__12337, let__12338)


tmp13543 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_ddatatypes_d, tmp13542)
}
__typedArg0 := symshen_4_ddatatypes_d
__typedArg1 := tmp13542
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

let__12339 := tmp13543
_ = let__12339

__e.TailApply(PrimFunc(symshen_4show_1datatypes), let__12339)
return


}, 1)

__e.TailApply(ns2_1set, syminclude_1all_1but, tmp13538)
return




}, 0)

