package main

import . "github.com/pyrex41/shen-go/kl"

var SequentMain = MakeNative0(func(__e *ControlFlow) Obj {
tmp12320 := MakeNative1(func(__e *ControlFlow, V3418 Obj) Obj {
tmp12335 := GConsP(__e, V3418)

var ifres12321 Obj

if True == tmp12335 {
tmp12322 := Call1(__e, SymFn(symhead), V3418)


let__12090 := tmp12322
_ = let__12090

tmp12323 := Call1(__e, SymFn(symtail), V3418)


let__12091 := tmp12323
_ = let__12091

tmp12324 := Call1(__e, SymFn(symshen_4_5datatype_1rules_6), let__12091)


let__12092 := tmp12324
_ = let__12092

tmp12333 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__12092)


var ifres12325 Obj

if True == tmp12333 {
tmp12326 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12325 = tmp12326


} else {
tmp12327 := Call1(__e, SymFn(symshen_4_5_1out), let__12092)


let__12093 := tmp12327
_ = let__12093

tmp12328 := Call1(__e, SymFn(symshen_4in_1_6), let__12092)


let__12094 := tmp12328
_ = let__12094

tmp12329 := Call2(__e, SymFn(symshen_4rules_1_6prolog), let__12090, let__12093)


let__12095 := tmp12329
_ = let__12095

tmp12330 := Call1(__e, SymFn(symfn), let__12090)


tmp12331 := Call2(__e, SymFn(symshen_4remember_1datatype), let__12090, tmp12330)


tmp12332 := Call2(__e, SymFn(symshen_4comb), let__12094, tmp12331)


ifres12325 = tmp12332


}

ifres12321 = ifres12325


} else {
tmp12334 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12321 = tmp12334


}

let__12089 := ifres12321
_ = let__12089

tmp12337 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__12089)


if True == tmp12337 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__12089
}


panic("unreachable")
})

tmp12338 := Call2(__e, ns2_1set, symshen_4_5datatype_6, tmp12320)


_ = tmp12338

tmp12339 := MakeNative2(func(__e *ControlFlow, V3426, V3427 Obj) Obj {
tmp12340 := GValue(__e, symshen_4_ddatatypes_d)

tmp12341 := Call3(__e, SymFn(symshen_4assoc_1_6), V3426, V3427, tmp12340)


tmp12342 := PrimGuard2(__e, symset, PrimSet, symshen_4_ddatatypes_d, tmp12341)

_ = tmp12342

tmp12343 := GValue(__e, symshen_4_dalldatatypes_d)

tmp12344 := Call3(__e, SymFn(symshen_4assoc_1_6), V3426, V3427, tmp12343)


tmp12345 := PrimGuard2(__e, symset, PrimSet, symshen_4_dalldatatypes_d, tmp12344)

_ = tmp12345

return V3426


panic("unreachable")
})

tmp12346 := Call2(__e, ns2_1set, symshen_4remember_1datatype, tmp12339)


_ = tmp12346

tmp12347 := MakeNative1(func(__e *ControlFlow, V3428 Obj) Obj {
tmp12348 := Call1(__e, SymFn(symshen_4_5datatype_1rule_6), V3428)


let__12097 := tmp12348
_ = let__12097

tmp12361 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__12097)


var ifres12349 Obj

if True == tmp12361 {
tmp12350 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12349 = tmp12350


} else {
tmp12351 := Call1(__e, SymFn(symshen_4_5_1out), let__12097)


let__12098 := tmp12351
_ = let__12098

tmp12352 := Call1(__e, SymFn(symshen_4in_1_6), let__12097)


let__12099 := tmp12352
_ = let__12099

tmp12353 := Call1(__e, SymFn(symshen_4_5datatype_1rules_6), let__12099)


let__12100 := tmp12353
_ = let__12100

tmp12360 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__12100)


var ifres12354 Obj

if True == tmp12360 {
tmp12355 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12354 = tmp12355


} else {
tmp12356 := Call1(__e, SymFn(symshen_4_5_1out), let__12100)


let__12101 := tmp12356
_ = let__12101

tmp12357 := Call1(__e, SymFn(symshen_4in_1_6), let__12100)


let__12102 := tmp12357
_ = let__12102

tmp12358 := Call2(__e, SymFn(symappend), let__12098, let__12101)


tmp12359 := Call2(__e, SymFn(symshen_4comb), let__12102, tmp12358)


ifres12354 = tmp12359


}

ifres12349 = ifres12354


}

let__12096 := ifres12349
_ = let__12096

tmp12377 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__12096)


if True == tmp12377 {
tmp12362 := Call1(__e, SymFn(sym_5_b_6), V3428)


let__12104 := tmp12362
_ = let__12104

tmp12373 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__12104)


var ifres12363 Obj

if True == tmp12373 {
tmp12364 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12363 = tmp12364


} else {
tmp12365 := Call1(__e, SymFn(symshen_4_5_1out), let__12104)


let__12105 := tmp12365
_ = let__12105

tmp12366 := Call1(__e, SymFn(symshen_4in_1_6), let__12104)


let__12106 := tmp12366
_ = let__12106

tmp12371 := Call1(__e, SymFn(symempty_2), let__12105)


var ifres12367 Obj

if True == tmp12371 {
ifres12367 = Nil


} else {
tmp12368 := Call3(__e, SymFn(symshen_4app), let__12105, MakeString("\n ..."), symshen_4r)


tmp12370 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("datatype syntax error here:\n "), tmp12368))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("datatype syntax error here:\n "), tmp12368)
return Call1(__e, SymFn(symsimple_1error), __typedArg0)
})()

ifres12367 = tmp12370


}

tmp12372 := Call2(__e, SymFn(symshen_4comb), let__12106, ifres12367)


ifres12363 = tmp12372


}

let__12103 := ifres12363
_ = let__12103

tmp12375 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__12103)


if True == tmp12375 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__12103
}


} else {
return let__12096
}


panic("unreachable")
})

tmp12378 := Call2(__e, ns2_1set, symshen_4_5datatype_1rules_6, tmp12347)


_ = tmp12378

tmp12379 := MakeNative1(func(__e *ControlFlow, V3440 Obj) Obj {
tmp12380 := Call1(__e, SymFn(symshen_4_5single_6), V3440)


let__12108 := tmp12380
_ = let__12108

tmp12386 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__12108)


var ifres12381 Obj

if True == tmp12386 {
tmp12382 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12381 = tmp12382


} else {
tmp12383 := Call1(__e, SymFn(symshen_4_5_1out), let__12108)


let__12109 := tmp12383
_ = let__12109

tmp12384 := Call1(__e, SymFn(symshen_4in_1_6), let__12108)


let__12110 := tmp12384
_ = let__12110

tmp12385 := Call2(__e, SymFn(symshen_4comb), let__12110, let__12109)


ifres12381 = tmp12385


}

let__12107 := ifres12381
_ = let__12107

tmp12397 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__12107)


if True == tmp12397 {
tmp12387 := Call1(__e, SymFn(symshen_4_5double_6), V3440)


let__12112 := tmp12387
_ = let__12112

tmp12393 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__12112)


var ifres12388 Obj

if True == tmp12393 {
tmp12389 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12388 = tmp12389


} else {
tmp12390 := Call1(__e, SymFn(symshen_4_5_1out), let__12112)


let__12113 := tmp12390
_ = let__12113

tmp12391 := Call1(__e, SymFn(symshen_4in_1_6), let__12112)


let__12114 := tmp12391
_ = let__12114

tmp12392 := Call2(__e, SymFn(symshen_4comb), let__12114, let__12113)


ifres12388 = tmp12392


}

let__12111 := ifres12388
_ = let__12111

tmp12395 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__12111)


if True == tmp12395 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__12111
}


} else {
return let__12107
}


panic("unreachable")
})

tmp12398 := Call2(__e, ns2_1set, symshen_4_5datatype_1rule_6, tmp12379)


_ = tmp12398

tmp12399 := MakeNative1(func(__e *ControlFlow, V3449 Obj) Obj {
tmp12400 := Call1(__e, SymFn(symshen_4_5sides_6), V3449)


let__12116 := tmp12400
_ = let__12116

tmp12432 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__12116)


var ifres12401 Obj

if True == tmp12432 {
tmp12402 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12401 = tmp12402


} else {
tmp12403 := Call1(__e, SymFn(symshen_4_5_1out), let__12116)


let__12117 := tmp12403
_ = let__12117

tmp12404 := Call1(__e, SymFn(symshen_4in_1_6), let__12116)


let__12118 := tmp12404
_ = let__12118

tmp12405 := Call1(__e, SymFn(symshen_4_5prems_6), let__12118)


let__12119 := tmp12405
_ = let__12119

tmp12431 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__12119)


var ifres12406 Obj

if True == tmp12431 {
tmp12407 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12406 = tmp12407


} else {
tmp12408 := Call1(__e, SymFn(symshen_4_5_1out), let__12119)


let__12120 := tmp12408
_ = let__12120

tmp12409 := Call1(__e, SymFn(symshen_4in_1_6), let__12119)


let__12121 := tmp12409
_ = let__12121

tmp12410 := Call1(__e, SymFn(symshen_4_5sng_6), let__12121)


let__12122 := tmp12410
_ = let__12122

tmp12430 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__12122)


var ifres12411 Obj

if True == tmp12430 {
tmp12412 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12411 = tmp12412


} else {
tmp12413 := Call1(__e, SymFn(symshen_4in_1_6), let__12122)


let__12123 := tmp12413
_ = let__12123

tmp12414 := Call1(__e, SymFn(symshen_4_5conc_6), let__12123)


let__12124 := tmp12414
_ = let__12124

tmp12429 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__12124)


var ifres12415 Obj

if True == tmp12429 {
tmp12416 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12415 = tmp12416


} else {
tmp12417 := Call1(__e, SymFn(symshen_4_5_1out), let__12124)


let__12125 := tmp12417
_ = let__12125

tmp12418 := Call1(__e, SymFn(symshen_4in_1_6), let__12124)


let__12126 := tmp12418
_ = let__12126

tmp12419 := Call1(__e, SymFn(symshen_4_5sc_6), let__12126)


let__12127 := tmp12419
_ = let__12127

tmp12428 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__12127)


var ifres12420 Obj

if True == tmp12428 {
tmp12421 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12420 = tmp12421


} else {
tmp12422 := Call1(__e, SymFn(symshen_4in_1_6), let__12127)


let__12128 := tmp12422
_ = let__12128

tmp12423 := GCons(__e, let__12125, Nil)

tmp12424 := GCons(__e, let__12120, tmp12423)

tmp12425 := GCons(__e, let__12117, tmp12424)

tmp12426 := GCons(__e, tmp12425, Nil)

tmp12427 := Call2(__e, SymFn(symshen_4comb), let__12128, tmp12426)


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

tmp12434 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__12115)


if True == tmp12434 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__12115
}


panic("unreachable")
})

tmp12435 := Call2(__e, ns2_1set, symshen_4_5single_6, tmp12399)


_ = tmp12435

tmp12436 := MakeNative1(func(__e *ControlFlow, V3464 Obj) Obj {
tmp12437 := Call1(__e, SymFn(symshen_4_5sides_6), V3464)


let__12130 := tmp12437
_ = let__12130

tmp12468 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__12130)


var ifres12438 Obj

if True == tmp12468 {
tmp12439 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12438 = tmp12439


} else {
tmp12440 := Call1(__e, SymFn(symshen_4_5_1out), let__12130)


let__12131 := tmp12440
_ = let__12131

tmp12441 := Call1(__e, SymFn(symshen_4in_1_6), let__12130)


let__12132 := tmp12441
_ = let__12132

tmp12442 := Call1(__e, SymFn(symshen_4_5formulae_6), let__12132)


let__12133 := tmp12442
_ = let__12133

tmp12467 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__12133)


var ifres12443 Obj

if True == tmp12467 {
tmp12444 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12443 = tmp12444


} else {
tmp12445 := Call1(__e, SymFn(symshen_4_5_1out), let__12133)


let__12134 := tmp12445
_ = let__12134

tmp12446 := Call1(__e, SymFn(symshen_4in_1_6), let__12133)


let__12135 := tmp12446
_ = let__12135

tmp12447 := Call1(__e, SymFn(symshen_4_5dbl_6), let__12135)


let__12136 := tmp12447
_ = let__12136

tmp12466 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__12136)


var ifres12448 Obj

if True == tmp12466 {
tmp12449 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12448 = tmp12449


} else {
tmp12450 := Call1(__e, SymFn(symshen_4in_1_6), let__12136)


let__12137 := tmp12450
_ = let__12137

tmp12451 := Call1(__e, SymFn(symshen_4_5formula_6), let__12137)


let__12138 := tmp12451
_ = let__12138

tmp12465 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__12138)


var ifres12452 Obj

if True == tmp12465 {
tmp12453 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12452 = tmp12453


} else {
tmp12454 := Call1(__e, SymFn(symshen_4_5_1out), let__12138)


let__12139 := tmp12454
_ = let__12139

tmp12455 := Call1(__e, SymFn(symshen_4in_1_6), let__12138)


let__12140 := tmp12455
_ = let__12140

tmp12456 := Call1(__e, SymFn(symshen_4_5sc_6), let__12140)


let__12141 := tmp12456
_ = let__12141

tmp12464 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__12141)


var ifres12457 Obj

if True == tmp12464 {
tmp12458 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12457 = tmp12458


} else {
tmp12459 := Call1(__e, SymFn(symshen_4in_1_6), let__12141)


let__12142 := tmp12459
_ = let__12142

tmp12460 := GCons(__e, let__12139, Nil)

tmp12461 := GCons(__e, Nil, tmp12460)

tmp12462 := Call3(__e, SymFn(symshen_4lr_1rule), let__12131, let__12134, tmp12461)


tmp12463 := Call2(__e, SymFn(symshen_4comb), let__12142, tmp12462)


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

tmp12470 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__12129)


if True == tmp12470 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__12129
}


panic("unreachable")
})

tmp12471 := Call2(__e, ns2_1set, symshen_4_5double_6, tmp12436)


_ = tmp12471

tmp12472 := MakeNative1(func(__e *ControlFlow, V3479 Obj) Obj {
tmp12473 := Call1(__e, SymFn(symshen_4_5formula_6), V3479)


let__12144 := tmp12473
_ = let__12144

tmp12493 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__12144)


var ifres12474 Obj

if True == tmp12493 {
tmp12475 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12474 = tmp12475


} else {
tmp12476 := Call1(__e, SymFn(symshen_4_5_1out), let__12144)


let__12145 := tmp12476
_ = let__12145

tmp12477 := Call1(__e, SymFn(symshen_4in_1_6), let__12144)


let__12146 := tmp12477
_ = let__12146

tmp12478 := Call1(__e, SymFn(symshen_4_5sc_6), let__12146)


let__12147 := tmp12478
_ = let__12147

tmp12492 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__12147)


var ifres12479 Obj

if True == tmp12492 {
tmp12480 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12479 = tmp12480


} else {
tmp12481 := Call1(__e, SymFn(symshen_4in_1_6), let__12147)


let__12148 := tmp12481
_ = let__12148

tmp12482 := Call1(__e, SymFn(symshen_4_5formulae_6), let__12148)


let__12149 := tmp12482
_ = let__12149

tmp12491 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__12149)


var ifres12483 Obj

if True == tmp12491 {
tmp12484 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12483 = tmp12484


} else {
tmp12485 := Call1(__e, SymFn(symshen_4_5_1out), let__12149)


let__12150 := tmp12485
_ = let__12150

tmp12486 := Call1(__e, SymFn(symshen_4in_1_6), let__12149)


let__12151 := tmp12486
_ = let__12151

tmp12487 := GCons(__e, let__12145, Nil)

tmp12488 := GCons(__e, Nil, tmp12487)

tmp12489 := GCons(__e, tmp12488, let__12150)

tmp12490 := Call2(__e, SymFn(symshen_4comb), let__12151, tmp12489)


ifres12483 = tmp12490


}

ifres12479 = ifres12483


}

ifres12474 = ifres12479


}

let__12143 := ifres12474
_ = let__12143

tmp12512 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__12143)


if True == tmp12512 {
tmp12494 := Call1(__e, SymFn(symshen_4_5formula_6), V3479)


let__12153 := tmp12494
_ = let__12153

tmp12508 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__12153)


var ifres12495 Obj

if True == tmp12508 {
tmp12496 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12495 = tmp12496


} else {
tmp12497 := Call1(__e, SymFn(symshen_4_5_1out), let__12153)


let__12154 := tmp12497
_ = let__12154

tmp12498 := Call1(__e, SymFn(symshen_4in_1_6), let__12153)


let__12155 := tmp12498
_ = let__12155

tmp12499 := Call1(__e, SymFn(symshen_4_5sc_6), let__12155)


let__12156 := tmp12499
_ = let__12156

tmp12507 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__12156)


var ifres12500 Obj

if True == tmp12507 {
tmp12501 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12500 = tmp12501


} else {
tmp12502 := Call1(__e, SymFn(symshen_4in_1_6), let__12156)


let__12157 := tmp12502
_ = let__12157

tmp12503 := GCons(__e, let__12154, Nil)

tmp12504 := GCons(__e, Nil, tmp12503)

tmp12505 := GCons(__e, tmp12504, Nil)

tmp12506 := Call2(__e, SymFn(symshen_4comb), let__12157, tmp12505)


ifres12500 = tmp12506


}

ifres12495 = ifres12500


}

let__12152 := ifres12495
_ = let__12152

tmp12510 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__12152)


if True == tmp12510 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__12152
}


} else {
return let__12143
}


panic("unreachable")
})

tmp12513 := Call2(__e, ns2_1set, symshen_4_5formulae_6, tmp12472)


_ = tmp12513

tmp12514 := MakeNative1(func(__e *ControlFlow, V3495 Obj) Obj {
tmp12515 := Call1(__e, SymFn(symshen_4_5ass_6), V3495)


let__12159 := tmp12515
_ = let__12159

tmp12533 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__12159)


var ifres12516 Obj

if True == tmp12533 {
tmp12517 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12516 = tmp12517


} else {
tmp12518 := Call1(__e, SymFn(symshen_4_5_1out), let__12159)


let__12160 := tmp12518
_ = let__12160

tmp12519 := Call1(__e, SymFn(symshen_4in_1_6), let__12159)


let__12161 := tmp12519
_ = let__12161

tmp12532 := Call2(__e, SymFn(symshen_4hds_a_2), let__12161, sym_6_6)


var ifres12520 Obj

if True == tmp12532 {
tmp12521 := Call1(__e, SymFn(symtail), let__12161)


let__12162 := tmp12521
_ = let__12162

tmp12522 := Call1(__e, SymFn(symshen_4_5formula_6), let__12162)


let__12163 := tmp12522
_ = let__12163

tmp12530 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__12163)


var ifres12523 Obj

if True == tmp12530 {
tmp12524 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12523 = tmp12524


} else {
tmp12525 := Call1(__e, SymFn(symshen_4_5_1out), let__12163)


let__12164 := tmp12525
_ = let__12164

tmp12526 := Call1(__e, SymFn(symshen_4in_1_6), let__12163)


let__12165 := tmp12526
_ = let__12165

tmp12527 := GCons(__e, let__12164, Nil)

tmp12528 := GCons(__e, let__12160, tmp12527)

tmp12529 := Call2(__e, SymFn(symshen_4comb), let__12165, tmp12528)


ifres12523 = tmp12529


}

ifres12520 = ifres12523


} else {
tmp12531 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12520 = tmp12531


}

ifres12516 = ifres12520


}

let__12158 := ifres12516
_ = let__12158

tmp12546 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__12158)


if True == tmp12546 {
tmp12534 := Call1(__e, SymFn(symshen_4_5formula_6), V3495)


let__12167 := tmp12534
_ = let__12167

tmp12542 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__12167)


var ifres12535 Obj

if True == tmp12542 {
tmp12536 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12535 = tmp12536


} else {
tmp12537 := Call1(__e, SymFn(symshen_4_5_1out), let__12167)


let__12168 := tmp12537
_ = let__12168

tmp12538 := Call1(__e, SymFn(symshen_4in_1_6), let__12167)


let__12169 := tmp12538
_ = let__12169

tmp12539 := GCons(__e, let__12168, Nil)

tmp12540 := GCons(__e, Nil, tmp12539)

tmp12541 := Call2(__e, SymFn(symshen_4comb), let__12169, tmp12540)


ifres12535 = tmp12541


}

let__12166 := ifres12535
_ = let__12166

tmp12544 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__12166)


if True == tmp12544 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__12166
}


} else {
return let__12158
}


panic("unreachable")
})

tmp12547 := Call2(__e, ns2_1set, symshen_4_5conc_6, tmp12514)


_ = tmp12547

tmp12548 := MakeNative1(func(__e *ControlFlow, V3508 Obj) Obj {
tmp12549 := Call1(__e, SymFn(symshen_4_5prem_6), V3508)


let__12171 := tmp12549
_ = let__12171

tmp12567 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__12171)


var ifres12550 Obj

if True == tmp12567 {
tmp12551 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12550 = tmp12551


} else {
tmp12552 := Call1(__e, SymFn(symshen_4_5_1out), let__12171)


let__12172 := tmp12552
_ = let__12172

tmp12553 := Call1(__e, SymFn(symshen_4in_1_6), let__12171)


let__12173 := tmp12553
_ = let__12173

tmp12554 := Call1(__e, SymFn(symshen_4_5sc_6), let__12173)


let__12174 := tmp12554
_ = let__12174

tmp12566 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__12174)


var ifres12555 Obj

if True == tmp12566 {
tmp12556 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12555 = tmp12556


} else {
tmp12557 := Call1(__e, SymFn(symshen_4in_1_6), let__12174)


let__12175 := tmp12557
_ = let__12175

tmp12558 := Call1(__e, SymFn(symshen_4_5prems_6), let__12175)


let__12176 := tmp12558
_ = let__12176

tmp12565 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__12176)


var ifres12559 Obj

if True == tmp12565 {
tmp12560 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12559 = tmp12560


} else {
tmp12561 := Call1(__e, SymFn(symshen_4_5_1out), let__12176)


let__12177 := tmp12561
_ = let__12177

tmp12562 := Call1(__e, SymFn(symshen_4in_1_6), let__12176)


let__12178 := tmp12562
_ = let__12178

tmp12563 := GCons(__e, let__12172, let__12177)

tmp12564 := Call2(__e, SymFn(symshen_4comb), let__12178, tmp12563)


ifres12559 = tmp12564


}

ifres12555 = ifres12559


}

ifres12550 = ifres12555


}

let__12170 := ifres12550
_ = let__12170

tmp12577 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__12170)


if True == tmp12577 {
tmp12568 := Call1(__e, SymFn(sym_5e_6), V3508)


let__12180 := tmp12568
_ = let__12180

tmp12573 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__12180)


var ifres12569 Obj

if True == tmp12573 {
tmp12570 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12569 = tmp12570


} else {
tmp12571 := Call1(__e, SymFn(symshen_4in_1_6), let__12180)


let__12181 := tmp12571
_ = let__12181

tmp12572 := Call2(__e, SymFn(symshen_4comb), let__12181, Nil)


ifres12569 = tmp12572


}

let__12179 := ifres12569
_ = let__12179

tmp12575 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__12179)


if True == tmp12575 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__12179
}


} else {
return let__12170
}


panic("unreachable")
})

tmp12578 := Call2(__e, ns2_1set, symshen_4_5prems_6, tmp12548)


_ = tmp12578

tmp12579 := MakeNative1(func(__e *ControlFlow, V3521 Obj) Obj {
tmp12584 := Call2(__e, SymFn(symshen_4hds_a_2), V3521, sym_b)


var ifres12580 Obj

if True == tmp12584 {
tmp12581 := Call1(__e, SymFn(symtail), V3521)


let__12183 := tmp12581
_ = let__12183

tmp12582 := Call2(__e, SymFn(symshen_4comb), let__12183, sym_b)


ifres12580 = tmp12582


} else {
tmp12583 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12580 = tmp12583


}

let__12182 := ifres12580
_ = let__12182

tmp12618 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__12182)


if True == tmp12618 {
tmp12585 := Call1(__e, SymFn(symshen_4_5ass_6), V3521)


let__12185 := tmp12585
_ = let__12185

tmp12603 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__12185)


var ifres12586 Obj

if True == tmp12603 {
tmp12587 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12586 = tmp12587


} else {
tmp12588 := Call1(__e, SymFn(symshen_4_5_1out), let__12185)


let__12186 := tmp12588
_ = let__12186

tmp12589 := Call1(__e, SymFn(symshen_4in_1_6), let__12185)


let__12187 := tmp12589
_ = let__12187

tmp12602 := Call2(__e, SymFn(symshen_4hds_a_2), let__12187, sym_6_6)


var ifres12590 Obj

if True == tmp12602 {
tmp12591 := Call1(__e, SymFn(symtail), let__12187)


let__12188 := tmp12591
_ = let__12188

tmp12592 := Call1(__e, SymFn(symshen_4_5formula_6), let__12188)


let__12189 := tmp12592
_ = let__12189

tmp12600 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__12189)


var ifres12593 Obj

if True == tmp12600 {
tmp12594 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12593 = tmp12594


} else {
tmp12595 := Call1(__e, SymFn(symshen_4_5_1out), let__12189)


let__12190 := tmp12595
_ = let__12190

tmp12596 := Call1(__e, SymFn(symshen_4in_1_6), let__12189)


let__12191 := tmp12596
_ = let__12191

tmp12597 := GCons(__e, let__12190, Nil)

tmp12598 := GCons(__e, let__12186, tmp12597)

tmp12599 := Call2(__e, SymFn(symshen_4comb), let__12191, tmp12598)


ifres12593 = tmp12599


}

ifres12590 = ifres12593


} else {
tmp12601 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12590 = tmp12601


}

ifres12586 = ifres12590


}

let__12184 := ifres12586
_ = let__12184

tmp12616 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__12184)


if True == tmp12616 {
tmp12604 := Call1(__e, SymFn(symshen_4_5formula_6), V3521)


let__12193 := tmp12604
_ = let__12193

tmp12612 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__12193)


var ifres12605 Obj

if True == tmp12612 {
tmp12606 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12605 = tmp12606


} else {
tmp12607 := Call1(__e, SymFn(symshen_4_5_1out), let__12193)


let__12194 := tmp12607
_ = let__12194

tmp12608 := Call1(__e, SymFn(symshen_4in_1_6), let__12193)


let__12195 := tmp12608
_ = let__12195

tmp12609 := GCons(__e, let__12194, Nil)

tmp12610 := GCons(__e, Nil, tmp12609)

tmp12611 := Call2(__e, SymFn(symshen_4comb), let__12195, tmp12610)


ifres12605 = tmp12611


}

let__12192 := ifres12605
_ = let__12192

tmp12614 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__12192)


if True == tmp12614 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__12192
}


} else {
return let__12184
}


} else {
return let__12182
}


panic("unreachable")
})

tmp12619 := Call2(__e, ns2_1set, symshen_4_5prem_6, tmp12579)


_ = tmp12619

tmp12620 := MakeNative1(func(__e *ControlFlow, V3536 Obj) Obj {
tmp12621 := Call1(__e, SymFn(symshen_4_5formula_6), V3536)


let__12197 := tmp12621
_ = let__12197

tmp12639 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__12197)


var ifres12622 Obj

if True == tmp12639 {
tmp12623 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12622 = tmp12623


} else {
tmp12624 := Call1(__e, SymFn(symshen_4_5_1out), let__12197)


let__12198 := tmp12624
_ = let__12198

tmp12625 := Call1(__e, SymFn(symshen_4in_1_6), let__12197)


let__12199 := tmp12625
_ = let__12199

tmp12626 := Call1(__e, SymFn(symshen_4_5iscomma_6), let__12199)


let__12200 := tmp12626
_ = let__12200

tmp12638 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__12200)


var ifres12627 Obj

if True == tmp12638 {
tmp12628 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12627 = tmp12628


} else {
tmp12629 := Call1(__e, SymFn(symshen_4in_1_6), let__12200)


let__12201 := tmp12629
_ = let__12201

tmp12630 := Call1(__e, SymFn(symshen_4_5ass_6), let__12201)


let__12202 := tmp12630
_ = let__12202

tmp12637 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__12202)


var ifres12631 Obj

if True == tmp12637 {
tmp12632 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12631 = tmp12632


} else {
tmp12633 := Call1(__e, SymFn(symshen_4_5_1out), let__12202)


let__12203 := tmp12633
_ = let__12203

tmp12634 := Call1(__e, SymFn(symshen_4in_1_6), let__12202)


let__12204 := tmp12634
_ = let__12204

tmp12635 := GCons(__e, let__12198, let__12203)

tmp12636 := Call2(__e, SymFn(symshen_4comb), let__12204, tmp12635)


ifres12631 = tmp12636


}

ifres12627 = ifres12631


}

ifres12622 = ifres12627


}

let__12196 := ifres12622
_ = let__12196

tmp12659 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__12196)


if True == tmp12659 {
tmp12640 := Call1(__e, SymFn(symshen_4_5formula_6), V3536)


let__12206 := tmp12640
_ = let__12206

tmp12647 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__12206)


var ifres12641 Obj

if True == tmp12647 {
tmp12642 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12641 = tmp12642


} else {
tmp12643 := Call1(__e, SymFn(symshen_4_5_1out), let__12206)


let__12207 := tmp12643
_ = let__12207

tmp12644 := Call1(__e, SymFn(symshen_4in_1_6), let__12206)


let__12208 := tmp12644
_ = let__12208

tmp12645 := GCons(__e, let__12207, Nil)

tmp12646 := Call2(__e, SymFn(symshen_4comb), let__12208, tmp12645)


ifres12641 = tmp12646


}

let__12205 := ifres12641
_ = let__12205

tmp12657 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__12205)


if True == tmp12657 {
tmp12648 := Call1(__e, SymFn(sym_5e_6), V3536)


let__12210 := tmp12648
_ = let__12210

tmp12653 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__12210)


var ifres12649 Obj

if True == tmp12653 {
tmp12650 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12649 = tmp12650


} else {
tmp12651 := Call1(__e, SymFn(symshen_4in_1_6), let__12210)


let__12211 := tmp12651
_ = let__12211

tmp12652 := Call2(__e, SymFn(symshen_4comb), let__12211, Nil)


ifres12649 = tmp12652


}

let__12209 := ifres12649
_ = let__12209

tmp12655 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__12209)


if True == tmp12655 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__12209
}


} else {
return let__12205
}


} else {
return let__12196
}


panic("unreachable")
})

tmp12660 := Call2(__e, ns2_1set, symshen_4_5ass_6, tmp12620)


_ = tmp12660

tmp12661 := MakeNative1(func(__e *ControlFlow, V3553 Obj) Obj {
tmp12671 := GConsP(__e, V3553)

var ifres12662 Obj

if True == tmp12671 {
tmp12663 := Call1(__e, SymFn(symhead), V3553)


let__12213 := tmp12663
_ = let__12213

tmp12664 := Call1(__e, SymFn(symtail), V3553)


let__12214 := tmp12664
_ = let__12214

tmp12668 := PrimGuard1(__e, symintern, PrimIntern, MakeString(","))

tmp12669 := GEq(__e, let__12213, tmp12668)

var ifres12665 Obj

if True == tmp12669 {
tmp12666 := Call2(__e, SymFn(symshen_4comb), let__12214, symshen_4skip)


ifres12665 = tmp12666


} else {
tmp12667 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12665 = tmp12667


}

ifres12662 = ifres12665


} else {
tmp12670 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12662 = tmp12670


}

let__12212 := ifres12662
_ = let__12212

tmp12673 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__12212)


if True == tmp12673 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__12212
}


panic("unreachable")
})

tmp12674 := Call2(__e, ns2_1set, symshen_4_5iscomma_6, tmp12661)


_ = tmp12674

tmp12675 := MakeNative1(func(__e *ControlFlow, V3557 Obj) Obj {
tmp12676 := Call1(__e, SymFn(symshen_4_5expr_6), V3557)


let__12216 := tmp12676
_ = let__12216

tmp12699 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__12216)


var ifres12677 Obj

if True == tmp12699 {
tmp12678 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12677 = tmp12678


} else {
tmp12679 := Call1(__e, SymFn(symshen_4_5_1out), let__12216)


let__12217 := tmp12679
_ = let__12217

tmp12680 := Call1(__e, SymFn(symshen_4in_1_6), let__12216)


let__12218 := tmp12680
_ = let__12218

tmp12681 := Call1(__e, SymFn(symshen_4_5iscolon_6), let__12218)


let__12219 := tmp12681
_ = let__12219

tmp12698 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__12219)


var ifres12682 Obj

if True == tmp12698 {
tmp12683 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12682 = tmp12683


} else {
tmp12684 := Call1(__e, SymFn(symshen_4in_1_6), let__12219)


let__12220 := tmp12684
_ = let__12220

tmp12685 := Call1(__e, SymFn(symshen_4_5type_6), let__12220)


let__12221 := tmp12685
_ = let__12221

tmp12697 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__12221)


var ifres12686 Obj

if True == tmp12697 {
tmp12687 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12686 = tmp12687


} else {
tmp12688 := Call1(__e, SymFn(symshen_4_5_1out), let__12221)


let__12222 := tmp12688
_ = let__12222

tmp12689 := Call1(__e, SymFn(symshen_4in_1_6), let__12221)


let__12223 := tmp12689
_ = let__12223

tmp12690 := Call1(__e, SymFn(symshen_4curry), let__12217)


tmp12691 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":"))

tmp12692 := Call1(__e, SymFn(symshen_4rectify_1type), let__12222)


tmp12693 := GCons(__e, tmp12692, Nil)

tmp12694 := GCons(__e, tmp12691, tmp12693)

tmp12695 := GCons(__e, tmp12690, tmp12694)

tmp12696 := Call2(__e, SymFn(symshen_4comb), let__12223, tmp12695)


ifres12686 = tmp12696


}

ifres12682 = ifres12686


}

ifres12677 = ifres12682


}

let__12215 := ifres12677
_ = let__12215

tmp12710 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__12215)


if True == tmp12710 {
tmp12700 := Call1(__e, SymFn(symshen_4_5expr_6), V3557)


let__12225 := tmp12700
_ = let__12225

tmp12706 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__12225)


var ifres12701 Obj

if True == tmp12706 {
tmp12702 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12701 = tmp12702


} else {
tmp12703 := Call1(__e, SymFn(symshen_4_5_1out), let__12225)


let__12226 := tmp12703
_ = let__12226

tmp12704 := Call1(__e, SymFn(symshen_4in_1_6), let__12225)


let__12227 := tmp12704
_ = let__12227

tmp12705 := Call2(__e, SymFn(symshen_4comb), let__12227, let__12226)


ifres12701 = tmp12705


}

let__12224 := ifres12701
_ = let__12224

tmp12708 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__12224)


if True == tmp12708 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__12224
}


} else {
return let__12215
}


panic("unreachable")
})

tmp12711 := Call2(__e, ns2_1set, symshen_4_5formula_6, tmp12675)


_ = tmp12711

tmp12712 := MakeNative1(func(__e *ControlFlow, V3571 Obj) Obj {
tmp12722 := GConsP(__e, V3571)

var ifres12713 Obj

if True == tmp12722 {
tmp12714 := Call1(__e, SymFn(symhead), V3571)


let__12229 := tmp12714
_ = let__12229

tmp12715 := Call1(__e, SymFn(symtail), V3571)


let__12230 := tmp12715
_ = let__12230

tmp12719 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":"))

tmp12720 := GEq(__e, let__12229, tmp12719)

var ifres12716 Obj

if True == tmp12720 {
tmp12717 := Call2(__e, SymFn(symshen_4comb), let__12230, symshen_4skip)


ifres12716 = tmp12717


} else {
tmp12718 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12716 = tmp12718


}

ifres12713 = ifres12716


} else {
tmp12721 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12713 = tmp12721


}

let__12228 := ifres12713
_ = let__12228

tmp12724 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__12228)


if True == tmp12724 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__12228
}


panic("unreachable")
})

tmp12725 := Call2(__e, ns2_1set, symshen_4_5iscolon_6, tmp12712)


_ = tmp12725

tmp12726 := MakeNative1(func(__e *ControlFlow, V3575 Obj) Obj {
tmp12727 := Call1(__e, SymFn(symshen_4_5side_6), V3575)


let__12232 := tmp12727
_ = let__12232

tmp12740 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__12232)


var ifres12728 Obj

if True == tmp12740 {
tmp12729 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12728 = tmp12729


} else {
tmp12730 := Call1(__e, SymFn(symshen_4_5_1out), let__12232)


let__12233 := tmp12730
_ = let__12233

tmp12731 := Call1(__e, SymFn(symshen_4in_1_6), let__12232)


let__12234 := tmp12731
_ = let__12234

tmp12732 := Call1(__e, SymFn(symshen_4_5sides_6), let__12234)


let__12235 := tmp12732
_ = let__12235

tmp12739 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__12235)


var ifres12733 Obj

if True == tmp12739 {
tmp12734 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12733 = tmp12734


} else {
tmp12735 := Call1(__e, SymFn(symshen_4_5_1out), let__12235)


let__12236 := tmp12735
_ = let__12236

tmp12736 := Call1(__e, SymFn(symshen_4in_1_6), let__12235)


let__12237 := tmp12736
_ = let__12237

tmp12737 := GCons(__e, let__12233, let__12236)

tmp12738 := Call2(__e, SymFn(symshen_4comb), let__12237, tmp12737)


ifres12733 = tmp12738


}

ifres12728 = ifres12733


}

let__12231 := ifres12728
_ = let__12231

tmp12750 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__12231)


if True == tmp12750 {
tmp12741 := Call1(__e, SymFn(sym_5e_6), V3575)


let__12239 := tmp12741
_ = let__12239

tmp12746 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__12239)


var ifres12742 Obj

if True == tmp12746 {
tmp12743 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12742 = tmp12743


} else {
tmp12744 := Call1(__e, SymFn(symshen_4in_1_6), let__12239)


let__12240 := tmp12744
_ = let__12240

tmp12745 := Call2(__e, SymFn(symshen_4comb), let__12240, Nil)


ifres12742 = tmp12745


}

let__12238 := ifres12742
_ = let__12238

tmp12748 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__12238)


if True == tmp12748 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__12238
}


} else {
return let__12231
}


panic("unreachable")
})

tmp12751 := Call2(__e, ns2_1set, symshen_4_5sides_6, tmp12726)


_ = tmp12751

tmp12752 := MakeNative1(func(__e *ControlFlow, V3586 Obj) Obj {
tmp12764 := Call2(__e, SymFn(symshen_4hds_a_2), V3586, symif)


var ifres12753 Obj

if True == tmp12764 {
tmp12754 := Call1(__e, SymFn(symtail), V3586)


let__12242 := tmp12754
_ = let__12242

tmp12762 := GConsP(__e, let__12242)

var ifres12755 Obj

if True == tmp12762 {
tmp12756 := Call1(__e, SymFn(symhead), let__12242)


let__12243 := tmp12756
_ = let__12243

tmp12757 := Call1(__e, SymFn(symtail), let__12242)


let__12244 := tmp12757
_ = let__12244

tmp12758 := GCons(__e, let__12243, Nil)

tmp12759 := GCons(__e, symif, tmp12758)

tmp12760 := Call2(__e, SymFn(symshen_4comb), let__12244, tmp12759)


ifres12755 = tmp12760


} else {
tmp12761 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12755 = tmp12761


}

ifres12753 = ifres12755


} else {
tmp12763 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12753 = tmp12763


}

let__12241 := ifres12753
_ = let__12241

tmp12820 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__12241)


if True == tmp12820 {
tmp12782 := Call2(__e, SymFn(symshen_4hds_a_2), V3586, symlet)


var ifres12765 Obj

if True == tmp12782 {
tmp12766 := Call1(__e, SymFn(symtail), V3586)


let__12246 := tmp12766
_ = let__12246

tmp12780 := GConsP(__e, let__12246)

var ifres12767 Obj

if True == tmp12780 {
tmp12768 := Call1(__e, SymFn(symhead), let__12246)


let__12247 := tmp12768
_ = let__12247

tmp12769 := Call1(__e, SymFn(symtail), let__12246)


let__12248 := tmp12769
_ = let__12248

tmp12778 := GConsP(__e, let__12248)

var ifres12770 Obj

if True == tmp12778 {
tmp12771 := Call1(__e, SymFn(symhead), let__12248)


let__12249 := tmp12771
_ = let__12249

tmp12772 := Call1(__e, SymFn(symtail), let__12248)


let__12250 := tmp12772
_ = let__12250

tmp12773 := GCons(__e, let__12249, Nil)

tmp12774 := GCons(__e, let__12247, tmp12773)

tmp12775 := GCons(__e, symlet, tmp12774)

tmp12776 := Call2(__e, SymFn(symshen_4comb), let__12250, tmp12775)


ifres12770 = tmp12776


} else {
tmp12777 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12770 = tmp12777


}

ifres12767 = ifres12770


} else {
tmp12779 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12767 = tmp12779


}

ifres12765 = ifres12767


} else {
tmp12781 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12765 = tmp12781


}

let__12245 := ifres12765
_ = let__12245

tmp12818 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__12245)


if True == tmp12818 {
tmp12797 := Call2(__e, SymFn(symshen_4hds_a_2), V3586, symctxt)


var ifres12783 Obj

if True == tmp12797 {
tmp12784 := Call1(__e, SymFn(symtail), V3586)


let__12252 := tmp12784
_ = let__12252

tmp12795 := GConsP(__e, let__12252)

var ifres12785 Obj

if True == tmp12795 {
tmp12786 := Call1(__e, SymFn(symhead), let__12252)


let__12253 := tmp12786
_ = let__12253

tmp12787 := Call1(__e, SymFn(symtail), let__12252)


let__12254 := tmp12787
_ = let__12254

tmp12793 := PrimGuard1(__e, symvariable_2, PrimIsVariable, let__12253)

var ifres12788 Obj

if True == tmp12793 {
tmp12789 := GCons(__e, let__12253, Nil)

tmp12790 := GCons(__e, symctxt, tmp12789)

tmp12791 := Call2(__e, SymFn(symshen_4comb), let__12254, tmp12790)


ifres12788 = tmp12791


} else {
tmp12792 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12788 = tmp12792


}

ifres12785 = ifres12788


} else {
tmp12794 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12785 = tmp12794


}

ifres12783 = ifres12785


} else {
tmp12796 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12783 = tmp12796


}

let__12251 := ifres12783
_ = let__12251

tmp12816 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__12251)


if True == tmp12816 {
tmp12812 := Call2(__e, SymFn(symshen_4hds_a_2), V3586, symsqts)


var ifres12798 Obj

if True == tmp12812 {
tmp12799 := Call1(__e, SymFn(symtail), V3586)


let__12256 := tmp12799
_ = let__12256

tmp12810 := GConsP(__e, let__12256)

var ifres12800 Obj

if True == tmp12810 {
tmp12801 := Call1(__e, SymFn(symhead), let__12256)


let__12257 := tmp12801
_ = let__12257

tmp12802 := Call1(__e, SymFn(symtail), let__12256)


let__12258 := tmp12802
_ = let__12258

tmp12808 := PrimGuard1(__e, symvariable_2, PrimIsVariable, let__12257)

var ifres12803 Obj

if True == tmp12808 {
tmp12804 := GCons(__e, let__12257, Nil)

tmp12805 := GCons(__e, symsqts, tmp12804)

tmp12806 := Call2(__e, SymFn(symshen_4comb), let__12258, tmp12805)


ifres12803 = tmp12806


} else {
tmp12807 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12803 = tmp12807


}

ifres12800 = ifres12803


} else {
tmp12809 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12800 = tmp12809


}

ifres12798 = ifres12800


} else {
tmp12811 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12798 = tmp12811


}

let__12255 := ifres12798
_ = let__12255

tmp12814 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__12255)


if True == tmp12814 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__12255
}


} else {
return let__12251
}


} else {
return let__12245
}


} else {
return let__12241
}


panic("unreachable")
})

tmp12821 := Call2(__e, ns2_1set, symshen_4_5side_6, tmp12752)


_ = tmp12821

tmp12822 := MakeNative3(func(__e *ControlFlow, V3611, V3612, V3613 Obj) Obj {
tmp12852 := GConsP(__e, V3613)

var ifres12839 Obj

if True == tmp12852 {
tmp12850 := GHd(__e, V3613)

tmp12851 := GEq(__e, Nil, tmp12850)

var ifres12841 Obj

if True == tmp12851 {
tmp12848 := GTl(__e, V3613)

tmp12849 := GConsP(__e, tmp12848)

var ifres12843 Obj

if True == tmp12849 {
tmp12845 := GTl(__e, V3613)

tmp12846 := GTl(__e, tmp12845)

tmp12847 := GEq(__e, Nil, tmp12846)

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
tmp12823 := Call1(__e, SymFn(symgensym), symP)


let__12259 := tmp12823
_ = let__12259

tmp12824 := GTl(__e, V3613)

tmp12825 := GCons(__e, let__12259, Nil)

tmp12826 := GCons(__e, tmp12824, tmp12825)

let__12260 := tmp12826
_ = let__12260

tmp12827 := Call1(__e, SymFn(symshen_4coll_1formulae), V3612)


tmp12828 := GCons(__e, let__12259, Nil)

tmp12829 := GCons(__e, tmp12827, tmp12828)

let__12261 := tmp12829
_ = let__12261

tmp12830 := GCons(__e, let__12261, Nil)

tmp12831 := GCons(__e, let__12260, Nil)

tmp12832 := GCons(__e, tmp12830, tmp12831)

tmp12833 := GCons(__e, V3611, tmp12832)

let__12262 := tmp12833
_ = let__12262

tmp12834 := GCons(__e, V3613, Nil)

tmp12835 := GCons(__e, V3612, tmp12834)

tmp12836 := GCons(__e, V3611, tmp12835)

let__12263 := tmp12836
_ = let__12263

tmp12837 := GCons(__e, let__12262, Nil)

return GCons(__e, let__12263, tmp12837)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.lr-rule"))
}


panic("unreachable")
})

tmp12853 := Call2(__e, ns2_1set, symshen_4lr_1rule, tmp12822)


_ = tmp12853

tmp12854 := MakeNative1(func(__e *ControlFlow, V3621 Obj) Obj {
tmp12883 := GEq(__e, Nil, V3621)

if True == tmp12883 {
return Nil
} else {
tmp12881 := GConsP(__e, V3621)

var ifres12861 Obj

if True == tmp12881 {
tmp12879 := GHd(__e, V3621)

tmp12880 := GConsP(__e, tmp12879)

var ifres12863 Obj

if True == tmp12880 {
tmp12876 := GHd(__e, V3621)

tmp12877 := GHd(__e, tmp12876)

tmp12878 := GEq(__e, Nil, tmp12877)

var ifres12865 Obj

if True == tmp12878 {
tmp12873 := GHd(__e, V3621)

tmp12874 := GTl(__e, tmp12873)

tmp12875 := GConsP(__e, tmp12874)

var ifres12867 Obj

if True == tmp12875 {
tmp12869 := GHd(__e, V3621)

tmp12870 := GTl(__e, tmp12869)

tmp12871 := GTl(__e, tmp12870)

tmp12872 := GEq(__e, Nil, tmp12871)

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
tmp12855 := GHd(__e, V3621)

tmp12856 := GTl(__e, tmp12855)

tmp12857 := GHd(__e, tmp12856)

tmp12858 := GTl(__e, V3621)

tmp12859 := Call1(__e, SymFn(symshen_4coll_1formulae), tmp12858)


return GCons(__e, tmp12857, tmp12859)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.coll-formulae"))
}


}


panic("unreachable")
})

tmp12884 := Call2(__e, ns2_1set, symshen_4coll_1formulae, tmp12854)


_ = tmp12884

tmp12885 := MakeNative1(func(__e *ControlFlow, V3622 Obj) Obj {
tmp12896 := GConsP(__e, V3622)

var ifres12886 Obj

if True == tmp12896 {
tmp12887 := Call1(__e, SymFn(symhead), V3622)


let__12265 := tmp12887
_ = let__12265

tmp12888 := Call1(__e, SymFn(symtail), V3622)


let__12266 := tmp12888
_ = let__12266

tmp12893 := Call1(__e, SymFn(symshen_4key_1in_1sequent_1calculus_2), let__12265)


tmp12894 := GNot(__e, tmp12893)

var ifres12889 Obj

if True == tmp12894 {
tmp12890 := Call1(__e, SymFn(symmacroexpand), let__12265)


tmp12891 := Call2(__e, SymFn(symshen_4comb), let__12266, tmp12890)


ifres12889 = tmp12891


} else {
tmp12892 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12889 = tmp12892


}

ifres12886 = ifres12889


} else {
tmp12895 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12886 = tmp12895


}

let__12264 := ifres12886
_ = let__12264

tmp12898 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__12264)


if True == tmp12898 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__12264
}


panic("unreachable")
})

tmp12899 := Call2(__e, ns2_1set, symshen_4_5expr_6, tmp12885)


_ = tmp12899

tmp12900 := MakeNative1(func(__e *ControlFlow, V3626 Obj) Obj {
tmp12907 := PrimGuard1(__e, symintern, PrimIntern, MakeString(";"))

tmp12908 := PrimGuard1(__e, symintern, PrimIntern, MakeString(","))

tmp12909 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":"))

tmp12910 := GCons(__e, sym_5_1_1, Nil)

tmp12911 := GCons(__e, tmp12909, tmp12910)

tmp12912 := GCons(__e, tmp12908, tmp12911)

tmp12913 := GCons(__e, tmp12907, tmp12912)

tmp12914 := GCons(__e, sym_6_6, tmp12913)

tmp12915 := Call2(__e, SymFn(symelement_2), V3626, tmp12914)


if True == tmp12915 {
return True
} else {
tmp12905 := Call1(__e, SymFn(symshen_4sng_2), V3626)


var ifres12902 Obj

if True == tmp12905 {
ifres12902 = True


} else {
tmp12904 := Call1(__e, SymFn(symshen_4dbl_2), V3626)


var ifres12903 Obj

if True == tmp12904 {
ifres12903 = True


} else {
ifres12903 = False


}

ifres12902 = ifres12903


}

if True == ifres12902 {
return True
} else {
return False
}


}


panic("unreachable")
})

tmp12916 := Call2(__e, ns2_1set, symshen_4key_1in_1sequent_1calculus_2, tmp12900)


_ = tmp12916

tmp12917 := MakeNative1(func(__e *ControlFlow, V3627 Obj) Obj {
tmp12918 := Call1(__e, SymFn(symshen_4_5expr_6), V3627)


let__12268 := tmp12918
_ = let__12268

tmp12924 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__12268)


var ifres12919 Obj

if True == tmp12924 {
tmp12920 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12919 = tmp12920


} else {
tmp12921 := Call1(__e, SymFn(symshen_4_5_1out), let__12268)


let__12269 := tmp12921
_ = let__12269

tmp12922 := Call1(__e, SymFn(symshen_4in_1_6), let__12268)


let__12270 := tmp12922
_ = let__12270

tmp12923 := Call2(__e, SymFn(symshen_4comb), let__12270, let__12269)


ifres12919 = tmp12923


}

let__12267 := ifres12919
_ = let__12267

tmp12926 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__12267)


if True == tmp12926 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__12267
}


panic("unreachable")
})

tmp12927 := Call2(__e, ns2_1set, symshen_4_5type_6, tmp12917)


_ = tmp12927

tmp12928 := MakeNative1(func(__e *ControlFlow, V3632 Obj) Obj {
tmp12937 := GConsP(__e, V3632)

var ifres12929 Obj

if True == tmp12937 {
tmp12930 := Call1(__e, SymFn(symhead), V3632)


let__12272 := tmp12930
_ = let__12272

tmp12931 := Call1(__e, SymFn(symtail), V3632)


let__12273 := tmp12931
_ = let__12273

tmp12935 := Call1(__e, SymFn(symshen_4dbl_2), let__12272)


var ifres12932 Obj

if True == tmp12935 {
tmp12933 := Call2(__e, SymFn(symshen_4comb), let__12273, let__12272)


ifres12932 = tmp12933


} else {
tmp12934 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12932 = tmp12934


}

ifres12929 = ifres12932


} else {
tmp12936 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12929 = tmp12936


}

let__12271 := ifres12929
_ = let__12271

tmp12939 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__12271)


if True == tmp12939 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__12271
}


panic("unreachable")
})

tmp12940 := Call2(__e, ns2_1set, symshen_4_5dbl_6, tmp12928)


_ = tmp12940

tmp12941 := MakeNative1(func(__e *ControlFlow, V3636 Obj) Obj {
tmp12950 := GConsP(__e, V3636)

var ifres12942 Obj

if True == tmp12950 {
tmp12943 := Call1(__e, SymFn(symhead), V3636)


let__12275 := tmp12943
_ = let__12275

tmp12944 := Call1(__e, SymFn(symtail), V3636)


let__12276 := tmp12944
_ = let__12276

tmp12948 := Call1(__e, SymFn(symshen_4sng_2), let__12275)


var ifres12945 Obj

if True == tmp12948 {
tmp12946 := Call2(__e, SymFn(symshen_4comb), let__12276, let__12275)


ifres12945 = tmp12946


} else {
tmp12947 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12945 = tmp12947


}

ifres12942 = ifres12945


} else {
tmp12949 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12942 = tmp12949


}

let__12274 := ifres12942
_ = let__12274

tmp12952 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__12274)


if True == tmp12952 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__12274
}


panic("unreachable")
})

tmp12953 := Call2(__e, ns2_1set, symshen_4_5sng_6, tmp12941)


_ = tmp12953

tmp12954 := MakeNative1(func(__e *ControlFlow, V3640 Obj) Obj {
tmp12959 := PrimGuard1(__e, symsymbol_2, PrimIsSymbol, V3640)

if True == tmp12959 {
tmp12956 := PrimGuard1(__e, symstr, PrimStr, V3640)

tmp12957 := Call1(__e, SymFn(symshen_4sng_1h_2), tmp12956)


if True == tmp12957 {
return True
} else {
return False
}


} else {
return False
}


panic("unreachable")
})

tmp12960 := Call2(__e, ns2_1set, symshen_4sng_2, tmp12954)


_ = tmp12960

tmp12961 := (func() Obj {
var __self Obj
__self = MakeNative1(func(__e *ControlFlow, __self1 Obj) Obj {
__selftop:
V3643 := __self1
_ = V3643
tmp12970 := GEq(__e, MakeString("___"), V3643)

if True == tmp12970 {
return True
} else {
tmp12968 := Call1(__e, SymFn(symshen_4_7string_2), V3643)


var ifres12964 Obj

if True == tmp12968 {
tmp12966 := Call1(__e, SymFn(symhdstr), V3643)


tmp12967 := GEq(__e, MakeString("_"), tmp12966)

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
if SymFn(symshen_4sng_1h_2) == __self {
__self1 = PrimGuard1(__e, symtlstr, PrimTailString, V3643)
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4sng_1h_2), PrimGuard1(__e, symtlstr, PrimTailString, V3643))


} else {
return False
}


}


panic("unreachable")
})
return __self
})()

tmp12971 := Call2(__e, ns2_1set, symshen_4sng_1h_2, tmp12961)


_ = tmp12971

tmp12972 := MakeNative1(func(__e *ControlFlow, V3644 Obj) Obj {
tmp12977 := PrimGuard1(__e, symsymbol_2, PrimIsSymbol, V3644)

if True == tmp12977 {
tmp12974 := PrimGuard1(__e, symstr, PrimStr, V3644)

tmp12975 := Call1(__e, SymFn(symshen_4dbl_1h_2), tmp12974)


if True == tmp12975 {
return True
} else {
return False
}


} else {
return False
}


panic("unreachable")
})

tmp12978 := Call2(__e, ns2_1set, symshen_4dbl_2, tmp12972)


_ = tmp12978

tmp12979 := (func() Obj {
var __self Obj
__self = MakeNative1(func(__e *ControlFlow, __self1 Obj) Obj {
__selftop:
V3647 := __self1
_ = V3647
tmp12988 := GEq(__e, MakeString("==="), V3647)

if True == tmp12988 {
return True
} else {
tmp12986 := Call1(__e, SymFn(symshen_4_7string_2), V3647)


var ifres12982 Obj

if True == tmp12986 {
tmp12984 := Call1(__e, SymFn(symhdstr), V3647)


tmp12985 := GEq(__e, MakeString("="), tmp12984)

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
if SymFn(symshen_4dbl_1h_2) == __self {
__self1 = PrimGuard1(__e, symtlstr, PrimTailString, V3647)
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4dbl_1h_2), PrimGuard1(__e, symtlstr, PrimTailString, V3647))


} else {
return False
}


}


panic("unreachable")
})
return __self
})()

tmp12989 := Call2(__e, ns2_1set, symshen_4dbl_1h_2, tmp12979)


_ = tmp12989

tmp12990 := MakeNative2(func(__e *ControlFlow, V3648, V3649 Obj) Obj {
tmp12991 := MakeNative1(func(__e *ControlFlow, Z3651 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4rule_1_6clause), Z3651)
panic("unreachable")
})

tmp12992 := Call2(__e, SymFn(symmapcan), tmp12991, V3649)


let__12277 := tmp12992
_ = let__12277

tmp12993 := GCons(__e, V3648, let__12277)

tmp12994 := GCons(__e, symdefprolog, tmp12993)

let__12278 := tmp12994
_ = let__12278

return TailCall1(__e, SymFn(symeval), let__12278)


panic("unreachable")
})

tmp12995 := Call2(__e, ns2_1set, symshen_4rules_1_6prolog, tmp12990)


_ = tmp12995

tmp12996 := MakeNative1(func(__e *ControlFlow, V3653 Obj) Obj {
tmp13056 := GConsP(__e, V3653)

var ifres13020 Obj

if True == tmp13056 {
tmp13054 := GTl(__e, V3653)

tmp13055 := GConsP(__e, tmp13054)

var ifres13022 Obj

if True == tmp13055 {
tmp13051 := GTl(__e, V3653)

tmp13052 := GTl(__e, tmp13051)

tmp13053 := GConsP(__e, tmp13052)

var ifres13024 Obj

if True == tmp13053 {
tmp13047 := GTl(__e, V3653)

tmp13048 := GTl(__e, tmp13047)

tmp13049 := GHd(__e, tmp13048)

tmp13050 := GConsP(__e, tmp13049)

var ifres13026 Obj

if True == tmp13050 {
tmp13042 := GTl(__e, V3653)

tmp13043 := GTl(__e, tmp13042)

tmp13044 := GHd(__e, tmp13043)

tmp13045 := GTl(__e, tmp13044)

tmp13046 := GConsP(__e, tmp13045)

var ifres13028 Obj

if True == tmp13046 {
tmp13036 := GTl(__e, V3653)

tmp13037 := GTl(__e, tmp13036)

tmp13038 := GHd(__e, tmp13037)

tmp13039 := GTl(__e, tmp13038)

tmp13040 := GTl(__e, tmp13039)

tmp13041 := GEq(__e, Nil, tmp13040)

var ifres13030 Obj

if True == tmp13041 {
tmp13032 := GTl(__e, V3653)

tmp13033 := GTl(__e, tmp13032)

tmp13034 := GTl(__e, tmp13033)

tmp13035 := GEq(__e, Nil, tmp13034)

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
tmp12997 := GTl(__e, V3653)

tmp12998 := GTl(__e, tmp12997)

tmp12999 := GHd(__e, tmp12998)

tmp13000 := GTl(__e, tmp12999)

tmp13001 := GHd(__e, tmp13000)

tmp13002 := Call1(__e, SymFn(symshen_4extract_1vars), tmp13001)


let__12279 := tmp13002
_ = let__12279

tmp13003 := GTl(__e, V3653)

tmp13004 := GTl(__e, tmp13003)

tmp13005 := GHd(__e, tmp13004)

tmp13006 := GTl(__e, tmp13005)

tmp13007 := GHd(__e, tmp13006)

tmp13008 := Call1(__e, SymFn(symshen_4rule_1_6head), tmp13007)


tmp13009 := GCons(__e, sym_5_1_1, Nil)

tmp13010 := GHd(__e, V3653)

tmp13011 := GTl(__e, V3653)

tmp13012 := GHd(__e, tmp13011)

tmp13013 := GTl(__e, V3653)

tmp13014 := GTl(__e, tmp13013)

tmp13015 := GHd(__e, tmp13014)

tmp13016 := GHd(__e, tmp13015)

tmp13017 := Call5(__e, SymFn(symshen_4rule_1_6body), let__12279, symAssumptions, tmp13010, tmp13012, tmp13016)


tmp13018 := Call2(__e, SymFn(symappend), tmp13009, tmp13017)


return TailCall2(__e, SymFn(symappend), tmp13008, tmp13018)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4rule_1_6clause)
}


panic("unreachable")
})

tmp13057 := Call2(__e, ns2_1set, symshen_4rule_1_6clause, tmp12996)


_ = tmp13057

tmp13058 := MakeNative1(func(__e *ControlFlow, V3655 Obj) Obj {
tmp13059 := Call1(__e, SymFn(symshen_4macro_1_8ch), V3655)


tmp13060 := GCons(__e, symAssumptions, Nil)

return GCons(__e, tmp13059, tmp13060)


panic("unreachable")
})

tmp13061 := Call2(__e, ns2_1set, symshen_4rule_1_6head, tmp13058)


_ = tmp13061

tmp13062 := MakeNative1(func(__e *ControlFlow, V3656 Obj) Obj {
tmp13063 := GCons(__e, V3656, Nil)

return GCons(__e, symshen_4_8ch, tmp13063)


panic("unreachable")
})

tmp13064 := Call2(__e, ns2_1set, symshen_4macro_1_8ch, tmp13062)


_ = tmp13064

tmp13065 := MakeNative1(func(__e *ControlFlow, V3657 Obj) Obj {
tmp13066 := GCons(__e, V3657, Nil)

return GCons(__e, symshen_4_8c, tmp13066)


panic("unreachable")
})

tmp13067 := Call2(__e, ns2_1set, symshen_4macro_1_8c, tmp13065)


_ = tmp13067

tmp13068 := MakeNative5(func(__e *ControlFlow, V3658, V3659, V3660, V3661, V3662 Obj) Obj {
tmp13098 := GEq(__e, Nil, V3662)

if True == tmp13098 {
return TailCall5(__e, SymFn(symshen_4side_1conditions_1_6goals), Nil, V3658, V3659, V3660, V3661)
} else {
tmp13096 := GEq(__e, Nil, V3661)

var ifres13089 Obj

if True == tmp13096 {
tmp13095 := GConsP(__e, V3662)

var ifres13091 Obj

if True == tmp13095 {
tmp13093 := GTl(__e, V3662)

tmp13094 := GEq(__e, Nil, tmp13093)

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
tmp13069 := GHd(__e, V3662)

tmp13070 := Call2(__e, SymFn(symshen_4passive_1variables), tmp13069, V3658)


let__12280 := tmp13070
_ = let__12280

tmp13071 := GHd(__e, V3662)

tmp13072 := Call2(__e, SymFn(symshen_4remove_1bystanders), V3658, tmp13071)


let__12281 := tmp13072
_ = let__12281

tmp13073 := GHd(__e, V3662)

tmp13074 := Call4(__e, SymFn(symshen_4specialise_1member), tmp13073, V3659, let__12281, let__12280)


tmp13075 := Call5(__e, SymFn(symshen_4side_1conditions_1_6goals), Nil, V3658, V3659, V3660, Nil)


return GCons(__e, tmp13074, tmp13075)


} else {
tmp13087 := GConsP(__e, V3662)

if True == tmp13087 {
tmp13076 := Call1(__e, SymFn(symgensym), symNewAssumptions)


let__12282 := tmp13076
_ = let__12282

tmp13077 := GHd(__e, V3662)

tmp13078 := Call2(__e, SymFn(symshen_4passive_1variables), tmp13077, V3658)


let__12283 := tmp13078
_ = let__12283

tmp13079 := GHd(__e, V3662)

tmp13080 := Call2(__e, SymFn(symshen_4remove_1bystanders), V3658, tmp13079)


let__12284 := tmp13080
_ = let__12284

tmp13081 := GHd(__e, V3662)

tmp13082 := Call5(__e, SymFn(symshen_4specialise_1consume), tmp13081, V3659, let__12284, let__12283, let__12282)


tmp13083 := Call2(__e, SymFn(symappend), V3658, let__12283)


tmp13084 := GTl(__e, V3662)

tmp13085 := Call5(__e, SymFn(symshen_4rule_1_6body), tmp13083, let__12282, V3660, V3661, tmp13084)


return GCons(__e, tmp13082, tmp13085)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4rule_1_6body)
}


}


}


panic("unreachable")
})

tmp13099 := Call2(__e, ns2_1set, symshen_4rule_1_6body, tmp13068)


_ = tmp13099

tmp13100 := MakeNative4(func(__e *ControlFlow, V3668, V3669, V3670, V3671 Obj) Obj {
tmp13101 := Call1(__e, SymFn(symgensym), symshen_4member)


let__12285 := tmp13101
_ = let__12285

tmp13102 := Call4(__e, SymFn(symshen_4member_1clause), let__12285, V3668, V3670, V3671)


let__12286 := tmp13102
_ = let__12286

tmp13103 := Call2(__e, SymFn(symappend), V3670, V3671)


tmp13104 := GCons(__e, V3669, tmp13103)

return GCons(__e, let__12285, tmp13104)


panic("unreachable")
})

tmp13105 := Call2(__e, ns2_1set, symshen_4specialise_1member, tmp13100)


_ = tmp13105

tmp13106 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V3676 := __self1
_ = V3676
V3677 := __self2
_ = V3677
tmp13120 := GEq(__e, Nil, V3676)

if True == tmp13120 {
return Nil
} else {
tmp13118 := GConsP(__e, V3676)

var ifres13114 Obj

if True == tmp13118 {
tmp13116 := GHd(__e, V3676)

tmp13117 := Call2(__e, SymFn(symshen_4occurs_1check_2), tmp13116, V3677)


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
tmp13107 := GHd(__e, V3676)

tmp13108 := GTl(__e, V3676)

tmp13109 := Call2(__e, SymFn(symshen_4remove_1bystanders), tmp13108, V3677)


return GCons(__e, tmp13107, tmp13109)


} else {
tmp13112 := GConsP(__e, V3676)

if True == tmp13112 {
tmp13110 := GTl(__e, V3676)

if SymFn(symshen_4remove_1bystanders) == __self {
__self1, __self2 = tmp13110, V3677
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4remove_1bystanders), tmp13110, V3677)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4remove_1bystanders)
}


}


}


panic("unreachable")
})
return __self
})()

tmp13121 := Call2(__e, ns2_1set, symshen_4remove_1bystanders, tmp13106)


_ = tmp13121

tmp13122 := MakeNative4(func(__e *ControlFlow, V3678, V3679, V3680, V3681 Obj) Obj {
tmp13123 := Call1(__e, SymFn(symlength), V3681)


tmp13124 := Call1(__e, SymFn(symshen_4nvars), tmp13123)


let__12287 := tmp13124
_ = let__12287

tmp13125 := Call1(__e, SymFn(symshen_4macro_1_8ch), V3679)


tmp13126 := GCons(__e, sym__, Nil)

tmp13127 := GCons(__e, tmp13125, tmp13126)

tmp13128 := GCons(__e, symcons, tmp13127)

tmp13129 := GCons(__e, tmp13128, Nil)

tmp13130 := GCons(__e, sym_1, tmp13129)

tmp13131 := GCons(__e, tmp13130, Nil)

tmp13132 := GCons(__e, sym_5_1_1, Nil)

tmp13133 := Call2(__e, SymFn(symshen_4passive_1bind), V3681, let__12287)


tmp13134 := PrimGuard1(__e, symintern, PrimIntern, MakeString(";"))

tmp13135 := GCons(__e, tmp13134, Nil)

tmp13136 := Call2(__e, SymFn(symappend), tmp13133, tmp13135)


tmp13137 := Call2(__e, SymFn(symappend), tmp13132, tmp13136)


tmp13138 := Call2(__e, SymFn(symappend), let__12287, tmp13137)


tmp13139 := Call2(__e, SymFn(symappend), V3680, tmp13138)


tmp13140 := Call2(__e, SymFn(symappend), tmp13131, tmp13139)


let__12288 := tmp13140
_ = let__12288

tmp13141 := Call1(__e, SymFn(symgensym), symHypotheses)


let__12290 := tmp13141
_ = let__12290

tmp13142 := Call2(__e, SymFn(symappend), V3680, V3681)


let__12291 := tmp13142
_ = let__12291

tmp13143 := GCons(__e, let__12290, Nil)

tmp13144 := GCons(__e, sym__, tmp13143)

tmp13145 := GCons(__e, symcons, tmp13144)

tmp13146 := GCons(__e, tmp13145, Nil)

tmp13147 := GCons(__e, sym_1, tmp13146)

tmp13148 := GCons(__e, tmp13147, Nil)

tmp13149 := Call2(__e, SymFn(symappend), tmp13148, let__12291)


let__12292 := tmp13149
_ = let__12292

tmp13150 := GCons(__e, let__12290, let__12291)

tmp13151 := GCons(__e, V3678, tmp13150)

tmp13152 := GCons(__e, tmp13151, Nil)

let__12293 := tmp13152
_ = let__12293

tmp13153 := GCons(__e, sym_5_1_1, Nil)

tmp13154 := PrimGuard1(__e, symintern, PrimIntern, MakeString(";"))

tmp13155 := GCons(__e, tmp13154, Nil)

tmp13156 := Call2(__e, SymFn(symappend), let__12293, tmp13155)


tmp13157 := Call2(__e, SymFn(symappend), tmp13153, tmp13156)


tmp13158 := Call2(__e, SymFn(symappend), let__12292, tmp13157)


let__12289 := tmp13158
_ = let__12289

tmp13159 := Call2(__e, SymFn(symappend), let__12288, let__12289)


tmp13160 := GCons(__e, V3678, tmp13159)

tmp13161 := GCons(__e, symdefprolog, tmp13160)

let__12294 := tmp13161
_ = let__12294

return TailCall1(__e, SymFn(symeval), let__12294)


panic("unreachable")
})

tmp13162 := Call2(__e, ns2_1set, symshen_4member_1clause, tmp13122)


_ = tmp13162

tmp13163 := MakeNative1(func(__e *ControlFlow, V3690 Obj) Obj {
tmp13168 := GEq(__e, MakeInteger(0), V3690)

if True == tmp13168 {
return Nil
} else {
tmp13164 := Call1(__e, SymFn(symgensym), symNewV)


tmp13166 := Call1(__e, SymFn(symshen_4nvars), GSub(__e, V3690, MakeInteger(1)))


return GCons(__e, tmp13164, tmp13166)


}


panic("unreachable")
})

tmp13169 := Call2(__e, ns2_1set, symshen_4nvars, tmp13163)


_ = tmp13169

tmp13170 := MakeNative2(func(__e *ControlFlow, V3691, V3692 Obj) Obj {
tmp13188 := GEq(__e, Nil, V3691)

var ifres13185 Obj

if True == tmp13188 {
tmp13187 := GEq(__e, Nil, V3692)

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
return Nil
} else {
tmp13183 := GConsP(__e, V3691)

var ifres13180 Obj

if True == tmp13183 {
tmp13182 := GConsP(__e, V3692)

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
tmp13171 := GHd(__e, V3692)

tmp13172 := GHd(__e, V3691)

tmp13173 := GCons(__e, tmp13172, Nil)

tmp13174 := GCons(__e, tmp13171, tmp13173)

tmp13175 := GCons(__e, symbind, tmp13174)

tmp13176 := GTl(__e, V3691)

tmp13177 := GTl(__e, V3692)

tmp13178 := Call2(__e, SymFn(symshen_4passive_1bind), tmp13176, tmp13177)


return GCons(__e, tmp13175, tmp13178)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4passive_1bind)
}


}


panic("unreachable")
})

tmp13189 := Call2(__e, ns2_1set, symshen_4passive_1bind, tmp13170)


_ = tmp13189

tmp13190 := MakeNative5(func(__e *ControlFlow, V3693, V3694, V3695, V3696, V3697 Obj) Obj {
tmp13191 := Call1(__e, SymFn(symgensym), symshen_4consume)


let__12295 := tmp13191
_ = let__12295

tmp13192 := Call5(__e, SymFn(symshen_4consume_1clause), let__12295, V3693, V3695, V3696, V3697)


let__12296 := tmp13192
_ = let__12296

tmp13193 := Call2(__e, SymFn(symappend), V3695, V3696)


tmp13194 := GCons(__e, V3697, tmp13193)

tmp13195 := GCons(__e, V3694, tmp13194)

return GCons(__e, let__12295, tmp13195)


panic("unreachable")
})

tmp13196 := Call2(__e, ns2_1set, symshen_4specialise_1consume, tmp13190)


_ = tmp13196

tmp13197 := MakeNative5(func(__e *ControlFlow, V3700, V3701, V3702, V3703, V3704 Obj) Obj {
tmp13198 := Call1(__e, SymFn(symlength), V3703)


tmp13199 := Call1(__e, SymFn(symshen_4nvars), tmp13198)


let__12297 := tmp13199
_ = let__12297

tmp13200 := Call1(__e, SymFn(symgensym), symAssumption)


let__12298 := tmp13200
_ = let__12298

tmp13201 := Call1(__e, SymFn(symshen_4macro_1_8ch), V3701)


tmp13202 := GCons(__e, let__12298, Nil)

tmp13203 := GCons(__e, tmp13201, tmp13202)

tmp13204 := GCons(__e, symcons, tmp13203)

tmp13205 := GCons(__e, tmp13204, Nil)

tmp13206 := GCons(__e, sym_1, tmp13205)

tmp13207 := GCons(__e, sym_5_1_1, Nil)

tmp13208 := Call2(__e, SymFn(symshen_4passive_1bind), V3703, let__12297)


tmp13209 := GCons(__e, let__12298, Nil)

tmp13210 := GCons(__e, V3704, tmp13209)

tmp13211 := GCons(__e, symbind, tmp13210)

tmp13212 := PrimGuard1(__e, symintern, PrimIntern, MakeString(";"))

tmp13213 := GCons(__e, tmp13212, Nil)

tmp13214 := GCons(__e, tmp13211, tmp13213)

tmp13215 := Call2(__e, SymFn(symappend), tmp13208, tmp13214)


tmp13216 := Call2(__e, SymFn(symappend), tmp13207, tmp13215)


tmp13217 := Call2(__e, SymFn(symappend), let__12297, tmp13216)


tmp13218 := Call2(__e, SymFn(symappend), V3702, tmp13217)


tmp13219 := GCons(__e, V3704, tmp13218)

tmp13220 := GCons(__e, tmp13206, tmp13219)

let__12299 := tmp13220
_ = let__12299

tmp13221 := Call1(__e, SymFn(symgensym), symHypotheses)


let__12301 := tmp13221
_ = let__12301

tmp13222 := Call2(__e, SymFn(symappend), V3702, V3703)


let__12302 := tmp13222
_ = let__12302

tmp13223 := Call1(__e, SymFn(symgensym), symAssumptions)


let__12303 := tmp13223
_ = let__12303

tmp13224 := GCons(__e, let__12301, Nil)

tmp13225 := GCons(__e, let__12298, tmp13224)

tmp13226 := GCons(__e, symcons, tmp13225)

tmp13227 := GCons(__e, tmp13226, Nil)

tmp13228 := GCons(__e, sym_1, tmp13227)

tmp13229 := GCons(__e, V3704, Nil)

tmp13230 := GCons(__e, let__12303, tmp13229)

tmp13231 := GCons(__e, symcons, tmp13230)

tmp13232 := GCons(__e, tmp13231, let__12302)

tmp13233 := GCons(__e, tmp13228, tmp13232)

let__12304 := tmp13233
_ = let__12304

tmp13234 := GCons(__e, let__12298, Nil)

tmp13235 := GCons(__e, let__12303, tmp13234)

tmp13236 := GCons(__e, symbind, tmp13235)

tmp13237 := GCons(__e, V3704, let__12302)

tmp13238 := GCons(__e, let__12301, tmp13237)

tmp13239 := GCons(__e, V3700, tmp13238)

tmp13240 := GCons(__e, tmp13239, Nil)

tmp13241 := GCons(__e, tmp13236, tmp13240)

let__12305 := tmp13241
_ = let__12305

tmp13242 := GCons(__e, sym_5_1_1, Nil)

tmp13243 := PrimGuard1(__e, symintern, PrimIntern, MakeString(";"))

tmp13244 := GCons(__e, tmp13243, Nil)

tmp13245 := Call2(__e, SymFn(symappend), let__12305, tmp13244)


tmp13246 := Call2(__e, SymFn(symappend), tmp13242, tmp13245)


tmp13247 := Call2(__e, SymFn(symappend), let__12304, tmp13246)


let__12300 := tmp13247
_ = let__12300

tmp13248 := Call2(__e, SymFn(symappend), let__12299, let__12300)


tmp13249 := GCons(__e, V3700, tmp13248)

tmp13250 := GCons(__e, symdefprolog, tmp13249)

let__12306 := tmp13250
_ = let__12306

return TailCall1(__e, SymFn(symeval), let__12306)


panic("unreachable")
})

tmp13251 := Call2(__e, ns2_1set, symshen_4consume_1clause, tmp13197)


_ = tmp13251

tmp13252 := MakeNative2(func(__e *ControlFlow, V3715, V3716 Obj) Obj {
tmp13253 := Call1(__e, SymFn(symshen_4extract_1vars), V3715)


return TailCall2(__e, SymFn(symdifference), tmp13253, V3716)


panic("unreachable")
})

tmp13254 := Call2(__e, ns2_1set, symshen_4passive_1variables, tmp13252)


_ = tmp13254

tmp13255 := (func() Obj {
var __self Obj
__self = MakeNative5(func(__e *ControlFlow, __self1, __self2, __self3, __self4, __self5 Obj) Obj {
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
tmp13406 := GEq(__e, Nil, V3724)

if True == tmp13406 {
return TailCall3(__e, SymFn(symshen_4premises_1_6goals), V3721, V3723, V3725)
} else {
tmp13404 := GConsP(__e, V3724)

var ifres13384 Obj

if True == tmp13404 {
tmp13402 := GHd(__e, V3724)

tmp13403 := GConsP(__e, tmp13402)

var ifres13386 Obj

if True == tmp13403 {
tmp13399 := GHd(__e, V3724)

tmp13400 := GHd(__e, tmp13399)

tmp13401 := GEq(__e, symif, tmp13400)

var ifres13388 Obj

if True == tmp13401 {
tmp13396 := GHd(__e, V3724)

tmp13397 := GTl(__e, tmp13396)

tmp13398 := GConsP(__e, tmp13397)

var ifres13390 Obj

if True == tmp13398 {
tmp13392 := GHd(__e, V3724)

tmp13393 := GTl(__e, tmp13392)

tmp13394 := GTl(__e, tmp13393)

tmp13395 := GEq(__e, Nil, tmp13394)

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
tmp13256 := GHd(__e, V3724)

tmp13257 := GTl(__e, tmp13256)

tmp13258 := GCons(__e, symwhen, tmp13257)

tmp13259 := GTl(__e, V3724)

tmp13260 := Call5(__e, SymFn(symshen_4side_1conditions_1_6goals), V3721, V3722, V3723, tmp13259, V3725)


return GCons(__e, tmp13258, tmp13260)


} else {
tmp13382 := GConsP(__e, V3724)

var ifres13355 Obj

if True == tmp13382 {
tmp13380 := GHd(__e, V3724)

tmp13381 := GConsP(__e, tmp13380)

var ifres13357 Obj

if True == tmp13381 {
tmp13377 := GHd(__e, V3724)

tmp13378 := GHd(__e, tmp13377)

tmp13379 := GEq(__e, symlet, tmp13378)

var ifres13359 Obj

if True == tmp13379 {
tmp13374 := GHd(__e, V3724)

tmp13375 := GTl(__e, tmp13374)

tmp13376 := GConsP(__e, tmp13375)

var ifres13361 Obj

if True == tmp13376 {
tmp13370 := GHd(__e, V3724)

tmp13371 := GTl(__e, tmp13370)

tmp13372 := GTl(__e, tmp13371)

tmp13373 := GConsP(__e, tmp13372)

var ifres13363 Obj

if True == tmp13373 {
tmp13365 := GHd(__e, V3724)

tmp13366 := GTl(__e, tmp13365)

tmp13367 := GTl(__e, tmp13366)

tmp13368 := GTl(__e, tmp13367)

tmp13369 := GEq(__e, Nil, tmp13368)

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
tmp13276 := GHd(__e, V3724)

tmp13277 := GTl(__e, tmp13276)

tmp13278 := GHd(__e, tmp13277)

tmp13279 := Call2(__e, SymFn(symelement_2), tmp13278, V3722)


if True == tmp13279 {
tmp13261 := GHd(__e, V3724)

tmp13262 := GTl(__e, tmp13261)

tmp13263 := GCons(__e, symis_b, tmp13262)

tmp13264 := GTl(__e, V3724)

tmp13265 := Call5(__e, SymFn(symshen_4side_1conditions_1_6goals), V3721, V3722, V3723, tmp13264, V3725)


return GCons(__e, tmp13263, tmp13265)


} else {
tmp13266 := GHd(__e, V3724)

tmp13267 := GTl(__e, tmp13266)

tmp13268 := GCons(__e, symbind, tmp13267)

tmp13269 := GHd(__e, V3724)

tmp13270 := GTl(__e, tmp13269)

tmp13271 := GHd(__e, tmp13270)

tmp13272 := GCons(__e, tmp13271, V3722)

tmp13273 := GTl(__e, V3724)

tmp13274 := Call5(__e, SymFn(symshen_4side_1conditions_1_6goals), V3721, tmp13272, V3723, tmp13273, V3725)


return GCons(__e, tmp13268, tmp13274)


}


} else {
tmp13353 := GConsP(__e, V3724)

var ifres13333 Obj

if True == tmp13353 {
tmp13351 := GHd(__e, V3724)

tmp13352 := GConsP(__e, tmp13351)

var ifres13335 Obj

if True == tmp13352 {
tmp13348 := GHd(__e, V3724)

tmp13349 := GHd(__e, tmp13348)

tmp13350 := GEq(__e, symctxt, tmp13349)

var ifres13337 Obj

if True == tmp13350 {
tmp13345 := GHd(__e, V3724)

tmp13346 := GTl(__e, tmp13345)

tmp13347 := GConsP(__e, tmp13346)

var ifres13339 Obj

if True == tmp13347 {
tmp13341 := GHd(__e, V3724)

tmp13342 := GTl(__e, tmp13341)

tmp13343 := GTl(__e, tmp13342)

tmp13344 := GEq(__e, Nil, tmp13343)

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
tmp13305 := GHd(__e, V3724)

tmp13306 := GTl(__e, tmp13305)

tmp13307 := GHd(__e, tmp13306)

tmp13308 := Call2(__e, SymFn(symelement_2), tmp13307, V3722)


if True == tmp13308 {
tmp13280 := GHd(__e, V3724)

tmp13281 := GTl(__e, tmp13280)

tmp13282 := GHd(__e, tmp13281)

tmp13283 := GCons(__e, tmp13282, V3721)

tmp13284 := GTl(__e, V3724)

if SymFn(symshen_4side_1conditions_1_6goals) == __self {
__self1, __self2, __self3, __self4, __self5 = tmp13283, V3722, V3723, tmp13284, V3725
__e.Tick()
goto __selftop
}
return TailCall5(__e, SymFn(symshen_4side_1conditions_1_6goals), tmp13283, V3722, V3723, tmp13284, V3725)


} else {
tmp13285 := GHd(__e, V3724)

tmp13286 := GTl(__e, tmp13285)

tmp13287 := GHd(__e, tmp13286)

tmp13288 := GCons(__e, V3723, Nil)

tmp13289 := GCons(__e, tmp13287, tmp13288)

tmp13290 := GCons(__e, symbind, tmp13289)

tmp13291 := GHd(__e, V3724)

tmp13292 := GTl(__e, tmp13291)

tmp13293 := GHd(__e, tmp13292)

tmp13294 := GCons(__e, tmp13293, V3721)

tmp13295 := GHd(__e, V3724)

tmp13296 := GTl(__e, tmp13295)

tmp13297 := GHd(__e, tmp13296)

tmp13298 := GCons(__e, tmp13297, V3722)

tmp13299 := GHd(__e, V3724)

tmp13300 := GTl(__e, tmp13299)

tmp13301 := GHd(__e, tmp13300)

tmp13302 := GTl(__e, V3724)

tmp13303 := Call5(__e, SymFn(symshen_4side_1conditions_1_6goals), tmp13294, tmp13298, tmp13301, tmp13302, V3725)


return GCons(__e, tmp13290, tmp13303)


}


} else {
tmp13331 := GConsP(__e, V3724)

var ifres13311 Obj

if True == tmp13331 {
tmp13329 := GHd(__e, V3724)

tmp13330 := GConsP(__e, tmp13329)

var ifres13313 Obj

if True == tmp13330 {
tmp13326 := GHd(__e, V3724)

tmp13327 := GHd(__e, tmp13326)

tmp13328 := GEq(__e, symsqts, tmp13327)

var ifres13315 Obj

if True == tmp13328 {
tmp13323 := GHd(__e, V3724)

tmp13324 := GTl(__e, tmp13323)

tmp13325 := GConsP(__e, tmp13324)

var ifres13317 Obj

if True == tmp13325 {
tmp13319 := GHd(__e, V3724)

tmp13320 := GTl(__e, tmp13319)

tmp13321 := GTl(__e, tmp13320)

tmp13322 := GEq(__e, Nil, tmp13321)

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
tmp13309 := GTl(__e, V3724)

if SymFn(symshen_4side_1conditions_1_6goals) == __self {
__self1, __self2, __self3, __self4, __self5 = V3721, V3722, V3723, tmp13309, V3725
__e.Tick()
goto __selftop
}
return TailCall5(__e, SymFn(symshen_4side_1conditions_1_6goals), V3721, V3722, V3723, tmp13309, V3725)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4side_1conditions_1_6goals)
}


}


}


}


}


panic("unreachable")
})
return __self
})()

tmp13407 := Call2(__e, ns2_1set, symshen_4side_1conditions_1_6goals, tmp13255)


_ = tmp13407

tmp13408 := MakeNative3(func(__e *ControlFlow, V3730, V3731, V3732 Obj) Obj {
tmp13458 := GEq(__e, Nil, V3732)

if True == tmp13458 {
tmp13409 := PrimGuard1(__e, symintern, PrimIntern, MakeString(";"))

return GCons(__e, tmp13409, Nil)


} else {
tmp13456 := GConsP(__e, V3732)

var ifres13452 Obj

if True == tmp13456 {
tmp13454 := GHd(__e, V3732)

tmp13455 := GEq(__e, sym_b, tmp13454)

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
tmp13410 := GTl(__e, V3732)

tmp13411 := Call3(__e, SymFn(symshen_4premises_1_6goals), V3730, V3731, tmp13410)


return GCons(__e, sym_b, tmp13411)


} else {
tmp13450 := GConsP(__e, V3732)

var ifres13446 Obj

if True == tmp13450 {
tmp13448 := GHd(__e, V3732)

tmp13449 := GEq(__e, symfail, tmp13448)

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
tmp13412 := GCons(__e, False, Nil)

tmp13413 := GCons(__e, symwhen, tmp13412)

tmp13414 := GTl(__e, V3732)

tmp13415 := Call3(__e, SymFn(symshen_4premises_1_6goals), V3730, V3731, tmp13414)


return GCons(__e, tmp13413, tmp13415)


} else {
tmp13444 := GConsP(__e, V3732)

var ifres13429 Obj

if True == tmp13444 {
tmp13442 := GHd(__e, V3732)

tmp13443 := GConsP(__e, tmp13442)

var ifres13431 Obj

if True == tmp13443 {
tmp13439 := GHd(__e, V3732)

tmp13440 := GTl(__e, tmp13439)

tmp13441 := GConsP(__e, tmp13440)

var ifres13433 Obj

if True == tmp13441 {
tmp13435 := GHd(__e, V3732)

tmp13436 := GTl(__e, tmp13435)

tmp13437 := GTl(__e, tmp13436)

tmp13438 := GEq(__e, Nil, tmp13437)

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
tmp13416 := GHd(__e, V3732)

tmp13417 := GTl(__e, tmp13416)

tmp13418 := GHd(__e, tmp13417)

tmp13419 := Call1(__e, SymFn(symshen_4macro_1_8c), tmp13418)


tmp13420 := GHd(__e, V3732)

tmp13421 := GHd(__e, tmp13420)

tmp13422 := Call3(__e, SymFn(symshen_4construct_1context), V3730, tmp13421, V3731)


tmp13423 := GCons(__e, tmp13422, Nil)

tmp13424 := GCons(__e, tmp13419, tmp13423)

tmp13425 := GCons(__e, symshen_4system_1S, tmp13424)

tmp13426 := GTl(__e, V3732)

tmp13427 := Call3(__e, SymFn(symshen_4premises_1_6goals), V3730, V3731, tmp13426)


return GCons(__e, tmp13425, tmp13427)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4premises_1_6goals)
}


}


}


}


panic("unreachable")
})

tmp13459 := Call2(__e, ns2_1set, symshen_4premises_1_6goals, tmp13408)


_ = tmp13459

tmp13460 := MakeNative3(func(__e *ControlFlow, V3736, V3737, V3738 Obj) Obj {
tmp13480 := GEq(__e, Nil, V3737)

if True == tmp13480 {
return V3738
} else {
tmp13478 := GConsP(__e, V3737)

var ifres13470 Obj

if True == tmp13478 {
tmp13476 := GTl(__e, V3737)

tmp13477 := GEq(__e, Nil, tmp13476)

var ifres13472 Obj

if True == tmp13477 {
tmp13474 := GHd(__e, V3737)

tmp13475 := Call2(__e, SymFn(symelement_2), tmp13474, V3736)


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
return GHd(__e, V3737)
} else {
tmp13468 := GConsP(__e, V3737)

if True == tmp13468 {
tmp13461 := GHd(__e, V3737)

tmp13462 := Call1(__e, SymFn(symshen_4macro_1_8c), tmp13461)


tmp13463 := GTl(__e, V3737)

tmp13464 := Call3(__e, SymFn(symshen_4construct_1context), V3736, tmp13463, V3738)


tmp13465 := GCons(__e, tmp13464, Nil)

tmp13466 := GCons(__e, tmp13462, tmp13465)

return GCons(__e, symcons, tmp13466)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4construct_1context)
}


}


}


panic("unreachable")
})

tmp13481 := Call2(__e, ns2_1set, symshen_4construct_1context, tmp13460)


_ = tmp13481

tmp13482 := MakeNative1(func(__e *ControlFlow, V3739 Obj) Obj {
tmp13483 := MakeNative1(func(__e *ControlFlow, Z3741 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4intern_1type), Z3741)
panic("unreachable")
})

tmp13484 := Call2(__e, SymFn(symmap), tmp13483, V3739)


let__12307 := tmp13484
_ = let__12307

tmp13485 := GValue(__e, symshen_4_ddatatypes_d)

let__12308 := tmp13485
_ = let__12308

tmp13486 := Call2(__e, SymFn(symshen_4remove_1datatypes), let__12307, let__12308)


let__12309 := tmp13486
_ = let__12309

tmp13487 := PrimGuard2(__e, symset, PrimSet, symshen_4_ddatatypes_d, let__12309)

let__12310 := tmp13487
_ = let__12310

return TailCall1(__e, SymFn(symshen_4show_1datatypes), let__12310)


panic("unreachable")
})

tmp13488 := Call2(__e, ns2_1set, sympreclude, tmp13482)


_ = tmp13488

tmp13489 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V3749 := __self1
_ = V3749
V3750 := __self2
_ = V3750
tmp13496 := GEq(__e, Nil, V3749)

if True == tmp13496 {
return V3750
} else {
tmp13494 := GConsP(__e, V3749)

if True == tmp13494 {
tmp13490 := GTl(__e, V3749)

tmp13491 := GHd(__e, V3749)

tmp13492 := Call2(__e, SymFn(symshen_4unassoc), tmp13491, V3750)


if SymFn(symshen_4remove_1datatypes) == __self {
__self1, __self2 = tmp13490, tmp13492
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4remove_1datatypes), tmp13490, tmp13492)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.remove-datatypes"))
}


}


panic("unreachable")
})
return __self
})()

tmp13497 := Call2(__e, ns2_1set, symshen_4remove_1datatypes, tmp13489)


_ = tmp13497

tmp13498 := MakeNative1(func(__e *ControlFlow, V3751 Obj) Obj {
tmp13499 := MakeNative1(func(__e *ControlFlow, Z3752 Obj) Obj {
return GHd(__e, Z3752)
panic("unreachable")
})

return TailCall2(__e, SymFn(symmap), tmp13499, V3751)


panic("unreachable")
})

tmp13500 := Call2(__e, ns2_1set, symshen_4show_1datatypes, tmp13498)


_ = tmp13500

tmp13501 := MakeNative1(func(__e *ControlFlow, V3753 Obj) Obj {
tmp13502 := MakeNative1(func(__e *ControlFlow, Z3755 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4intern_1type), Z3755)
panic("unreachable")
})

tmp13503 := Call2(__e, SymFn(symmap), tmp13502, V3753)


let__12311 := tmp13503
_ = let__12311

tmp13504 := MakeNative1(func(__e *ControlFlow, Z3757 Obj) Obj {
tmp13505 := Call1(__e, SymFn(symfn), Z3757)


return TailCall2(__e, SymFn(symshen_4remember_1datatype), Z3757, tmp13505)


panic("unreachable")
})

tmp13506 := Call2(__e, SymFn(symmap), tmp13504, let__12311)


let__12312 := tmp13506
_ = let__12312

tmp13507 := GValue(__e, symshen_4_ddatatypes_d)

let__12313 := tmp13507
_ = let__12313

return TailCall1(__e, SymFn(symshen_4show_1datatypes), let__12313)


panic("unreachable")
})

tmp13508 := Call2(__e, ns2_1set, syminclude, tmp13501)


_ = tmp13508

tmp13509 := MakeNative1(func(__e *ControlFlow, V3759 Obj) Obj {
tmp13510 := PrimGuard2(__e, symset, PrimSet, symshen_4_ddatatypes_d, Nil)

let__12314 := tmp13510
_ = let__12314

tmp13511 := MakeNative1(func(__e *ControlFlow, Z3762 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4intern_1type), Z3762)
panic("unreachable")
})

tmp13512 := Call2(__e, SymFn(symmap), tmp13511, V3759)


let__12315 := tmp13512
_ = let__12315

tmp13513 := MakeNative1(func(__e *ControlFlow, Z3764 Obj) Obj {
tmp13514 := Call1(__e, SymFn(symfn), Z3764)


return TailCall2(__e, SymFn(symshen_4remember_1datatype), Z3764, tmp13514)


panic("unreachable")
})

tmp13515 := Call2(__e, SymFn(symmap), tmp13513, let__12315)


let__12316 := tmp13515
_ = let__12316

tmp13516 := GValue(__e, symshen_4_ddatatypes_d)

return TailCall1(__e, SymFn(symshen_4show_1datatypes), tmp13516)


panic("unreachable")
})

tmp13517 := Call2(__e, ns2_1set, sympreclude_1all_1but, tmp13509)


_ = tmp13517

tmp13518 := MakeNative1(func(__e *ControlFlow, V3765 Obj) Obj {
tmp13519 := MakeNative1(func(__e *ControlFlow, Z3767 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4intern_1type), Z3767)
panic("unreachable")
})

tmp13520 := Call2(__e, SymFn(symmap), tmp13519, V3765)


let__12317 := tmp13520
_ = let__12317

tmp13521 := GValue(__e, symshen_4_dalldatatypes_d)

let__12318 := tmp13521
_ = let__12318

tmp13522 := Call2(__e, SymFn(symshen_4remove_1datatypes), let__12317, let__12318)


tmp13523 := PrimGuard2(__e, symset, PrimSet, symshen_4_ddatatypes_d, tmp13522)

let__12319 := tmp13523
_ = let__12319

return TailCall1(__e, SymFn(symshen_4show_1datatypes), let__12319)


panic("unreachable")
})

return TailCall2(__e, ns2_1set, syminclude_1all_1but, tmp13518)




panic("unreachable")
})

