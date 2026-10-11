package main

import . "github.com/pyrex41/shen-go/kl"

var YaccMain = MakeNative(func(__e *ControlFlow) {
tmp17501 := MakeNative(func(__e *ControlFlow) {
V112 := __e.Get(1)
_ = V112
V113 := __e.Get(2)
_ = V113
tmp17502 := Call(__e, V112, V113)


let__17363 := tmp17502
_ = let__17363

tmp17509 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17363)


if True == tmp17509 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("parse failure\n"))
}
__typedArg0 := MakeString("parse failure\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
} else {
tmp17507 := Call(__e, PrimFunc(symshen_4partial_1parse_1failure_2), let__17363)


if True == tmp17507 {
tmp17503 := Call(__e, PrimFunc(symshen_4in_1_6), let__17363)


tmp17504 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dresidue_d, tmp17503)
}
__typedArg0 := symshen_4_dresidue_d
__typedArg1 := tmp17503
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp17504

tmp17505 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dresidue_d)
}
__typedArg0 := symshen_4_dresidue_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4raise_1syntax_1error), tmp17505)
return


} else {
__e.TailApply(PrimFunc(symshen_4_5_1out), let__17363)
return
}


}


}, 2)

tmp17510 := Call(__e, ns2_1set, symcompile, tmp17501)


_ = tmp17510

tmp17511 := MakeNative(func(__e *ControlFlow) {
V115 := __e.Get(1)
_ = V115
tmp17512 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dmaximum_1print_1sequence_1size_d)
}
__typedArg0 := sym_dmaximum_1print_1sequence_1size_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp17513 := Call(__e, PrimFunc(symshen_4syntax_1error_1message), tmp17512, MakeInteger(0), V115)


tmp17515 := Call(__e, PrimFunc(symshen_4proc_1nl), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("syntax error here: "))
__typedS1, __typedOK1 := TypedString(tmp17513)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("syntax error here: ")
__typedArg1 := tmp17513
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(tmp17515)
}
__typedArg0 := tmp17515
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}, 1)

tmp17516 := Call(__e, ns2_1set, symshen_4raise_1syntax_1error, tmp17511)


_ = tmp17516

tmp17517 := MakeNative(func(__e *ControlFlow) {
V123 := __e.Get(1)
_ = V123
V124 := __e.Get(2)
_ = V124
V125 := __e.Get(3)
_ = V125
tmp17528 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V125)
}
__typedArg0 := Nil
__typedArg1 := V125
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17528 {
__e.Return(MakeString("\n"))
return
} else {
tmp17526 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V123, V124)
}
__typedArg0 := V123
__typedArg1 := V124
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17526 {
__e.Return(MakeString("...etc \n"))
return
} else {
tmp17524 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V125)
}
__typedArg0 := V125
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17524 {
tmp17518 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V125)
}
__typedArg0 := V125
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17519 := Call(__e, PrimFunc(symshen_4app), tmp17518, MakeString(" "), symshen_4s)


tmp17520 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V124)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V124
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()

tmp17521 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V125)
}
__typedArg0 := V125
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17522 := Call(__e, PrimFunc(symshen_4syntax_1error_1message), V123, tmp17520, tmp17521)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(tmp17519)
__typedS1, __typedOK1 := TypedString(tmp17522)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := tmp17519
__typedArg1 := tmp17522
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4syntax_1error_1message)
return
}


}


}


}, 3)

tmp17529 := Call(__e, ns2_1set, symshen_4syntax_1error_1message, tmp17517)


_ = tmp17529

tmp17530 := MakeNative(func(__e *ControlFlow) {
V126 := __e.Get(1)
_ = V126
tmp17531 := Call(__e, PrimFunc(symfail))


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V126, tmp17531)
}
__typedArg0 := V126
__typedArg1 := tmp17531
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp17532 := Call(__e, ns2_1set, symshen_4parse_1failure_2, tmp17530)


_ = tmp17532

tmp17533 := MakeNative(func(__e *ControlFlow) {
V127 := __e.Get(1)
_ = V127
tmp17534 := Call(__e, PrimFunc(symshen_4in_1_6), V127)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp17534)
}
__typedArg0 := tmp17534
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})())
return


}, 1)

tmp17535 := Call(__e, ns2_1set, symshen_4partial_1parse_1failure_2, tmp17533)


_ = tmp17535

tmp17536 := MakeNative(func(__e *ControlFlow) {
V130 := __e.Get(1)
_ = V130
tmp17549 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V130)
}
__typedArg0 := V130
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17540 Obj

if True == tmp17549 {
tmp17547 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V130)
}
__typedArg0 := V130
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17548 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp17547)
}
__typedArg0 := tmp17547
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17542 Obj

if True == tmp17548 {
tmp17544 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V130)
}
__typedArg0 := V130
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17545 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp17544)
}
__typedArg0 := tmp17544
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17546 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp17545)
}
__typedArg0 := Nil
__typedArg1 := tmp17545
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17543 Obj

if True == tmp17546 {
ifres17543 = True


} else {
ifres17543 = False


}

ifres17542 = ifres17543


} else {
ifres17542 = False


}

var ifres17541 Obj

if True == ifres17542 {
ifres17541 = True


} else {
ifres17541 = False


}

ifres17540 = ifres17541


} else {
ifres17540 = False


}

if True == ifres17540 {
tmp17537 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V130)
}
__typedArg0 := V130
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp17537)
}
__typedArg0 := tmp17537
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return


} else {
tmp17538 := Call(__e, PrimFunc(symshen_4app), V130, MakeString(" is not a YACC stream\n"), symshen_4s)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(tmp17538)
}
__typedArg0 := tmp17538
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}


}, 1)

tmp17550 := Call(__e, ns2_1set, symshen_4objectcode, tmp17536)


_ = tmp17550

tmp17551 := MakeNative(func(__e *ControlFlow) {
V131 := __e.Get(1)
_ = V131
tmp17552 := MakeNative(func(__e *ControlFlow) {
Z132 := __e.Get(1)
_ = Z132
__e.TailApply(PrimFunc(symshen_4_5yacc_6), Z132)
return
}, 1)

__e.TailApply(PrimFunc(symcompile), tmp17552, V131)
return


}, 1)

tmp17553 := Call(__e, ns2_1set, symshen_4yacc_1_6shen, tmp17551)


_ = tmp17553

tmp17554 := MakeNative(func(__e *ControlFlow) {
V133 := __e.Get(1)
_ = V133
tmp17581 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V133)
}
__typedArg0 := V133
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17555 Obj

if True == tmp17581 {
tmp17556 := Call(__e, PrimFunc(symhead), V133)


let__17365 := tmp17556
_ = let__17365

tmp17557 := Call(__e, PrimFunc(symtail), V133)


let__17366 := tmp17557
_ = let__17366

tmp17558 := Call(__e, PrimFunc(symshen_4_5yaccsig_6), let__17366)


let__17367 := tmp17558
_ = let__17367

tmp17579 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17367)


var ifres17559 Obj

if True == tmp17579 {
tmp17560 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17559 = tmp17560


} else {
tmp17561 := Call(__e, PrimFunc(symshen_4_5_1out), let__17367)


let__17368 := tmp17561
_ = let__17368

tmp17562 := Call(__e, PrimFunc(symshen_4in_1_6), let__17367)


let__17369 := tmp17562
_ = let__17369

tmp17563 := Call(__e, PrimFunc(symshen_4_5c_1rules_6), let__17369)


let__17370 := tmp17563
_ = let__17370

tmp17578 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17370)


var ifres17564 Obj

if True == tmp17578 {
tmp17565 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17564 = tmp17565


} else {
tmp17566 := Call(__e, PrimFunc(symshen_4_5_1out), let__17370)


let__17371 := tmp17566
_ = let__17371

tmp17567 := Call(__e, PrimFunc(symshen_4in_1_6), let__17370)


let__17372 := tmp17567
_ = let__17372

tmp17568 := Call(__e, PrimFunc(symgensym), symS)


let__17373 := tmp17568
_ = let__17373

tmp17569 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17365, Nil)
}
__typedArg0 := let__17365
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17570 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefine, tmp17569)
}
__typedArg0 := symdefine
__typedArg1 := tmp17569
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17571 := Call(__e, PrimFunc(symshen_4c_1rules_1_6shen), let__17368, let__17373, let__17371)


tmp17572 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17571, Nil)
}
__typedArg0 := tmp17571
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17573 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_6, tmp17572)
}
__typedArg0 := sym_1_6
__typedArg1 := tmp17572
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17574 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17373, tmp17573)
}
__typedArg0 := let__17373
__typedArg1 := tmp17573
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17575 := Call(__e, PrimFunc(symappend), let__17368, tmp17574)


tmp17576 := Call(__e, PrimFunc(symappend), tmp17570, tmp17575)


let__17374 := tmp17576
_ = let__17374

tmp17577 := Call(__e, PrimFunc(symshen_4comb), let__17372, let__17374)


ifres17564 = tmp17577


}

ifres17559 = ifres17564


}

ifres17555 = ifres17559


} else {
tmp17580 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17555 = tmp17580


}

let__17364 := ifres17555
_ = let__17364

tmp17583 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17364)


if True == tmp17583 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__17364)
return
}


}, 1)

tmp17584 := Call(__e, ns2_1set, symshen_4_5yacc_6, tmp17554)


_ = tmp17584

tmp17585 := MakeNative(func(__e *ControlFlow) {
V145 := __e.Get(1)
_ = V145
tmp17642 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V145)
}
__typedArg0 := V145
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17586 Obj

if True == tmp17642 {
tmp17587 := Call(__e, PrimFunc(symhead), V145)


let__17376 := tmp17587
_ = let__17376

tmp17588 := Call(__e, PrimFunc(symtail), V145)


let__17377 := tmp17588
_ = let__17377

tmp17640 := Call(__e, PrimFunc(symshen_4ccons_2), let__17377)


var ifres17589 Obj

if True == tmp17640 {
tmp17590 := Call(__e, PrimFunc(symhead), let__17377)


let__17378 := tmp17590
_ = let__17378

tmp17591 := Call(__e, PrimFunc(symtail), let__17377)


let__17379 := tmp17591
_ = let__17379

tmp17638 := Call(__e, PrimFunc(symshen_4hds_a_2), let__17378, symlist)


var ifres17592 Obj

if True == tmp17638 {
tmp17593 := Call(__e, PrimFunc(symtail), let__17378)


let__17380 := tmp17593
_ = let__17380

tmp17636 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__17380)
}
__typedArg0 := let__17380
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17594 Obj

if True == tmp17636 {
tmp17595 := Call(__e, PrimFunc(symhead), let__17380)


let__17381 := tmp17595
_ = let__17381

tmp17596 := Call(__e, PrimFunc(symtail), let__17380)


let__17382 := tmp17596
_ = let__17382

tmp17597 := Call(__e, PrimFunc(sym_5end_6), let__17382)


let__17383 := tmp17597
_ = let__17383

tmp17634 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17383)


var ifres17598 Obj

if True == tmp17634 {
tmp17599 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17598 = tmp17599


} else {
tmp17600 := Call(__e, PrimFunc(symshen_4in_1_6), let__17383)


let__17384 := tmp17600
_ = let__17384

tmp17633 := Call(__e, PrimFunc(symshen_4hds_a_2), let__17379, sym_a_a_6)


var ifres17601 Obj

if True == tmp17633 {
tmp17602 := Call(__e, PrimFunc(symtail), let__17379)


let__17385 := tmp17602
_ = let__17385

tmp17631 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__17385)
}
__typedArg0 := let__17385
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17603 Obj

if True == tmp17631 {
tmp17604 := Call(__e, PrimFunc(symhead), let__17385)


let__17386 := tmp17604
_ = let__17386

tmp17605 := Call(__e, PrimFunc(symtail), let__17385)


let__17387 := tmp17605
_ = let__17387

tmp17629 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__17387)
}
__typedArg0 := let__17387
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17606 Obj

if True == tmp17629 {
tmp17607 := Call(__e, PrimFunc(symhead), let__17387)


let__17388 := tmp17607
_ = let__17388

tmp17608 := Call(__e, PrimFunc(symtail), let__17387)


let__17389 := tmp17608
_ = let__17389

tmp17627 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_i, let__17376)
}
__typedArg0 := sym_i
__typedArg1 := let__17376
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17624 Obj

if True == tmp17627 {
tmp17626 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_j, let__17388)
}
__typedArg0 := sym_j
__typedArg1 := let__17388
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17625 Obj

if True == tmp17626 {
ifres17625 = True


} else {
ifres17625 = False


}

ifres17624 = ifres17625


} else {
ifres17624 = False


}

var ifres17609 Obj

if True == ifres17624 {
tmp17610 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17381, Nil)
}
__typedArg0 := let__17381
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17611 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp17610)
}
__typedArg0 := symlist
__typedArg1 := tmp17610
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17612 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17381, Nil)
}
__typedArg0 := let__17381
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17613 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp17612)
}
__typedArg0 := symlist
__typedArg1 := tmp17612
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17614 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17386, Nil)
}
__typedArg0 := let__17386
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17615 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17613, tmp17614)
}
__typedArg0 := tmp17613
__typedArg1 := tmp17614
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17616 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstr, tmp17615)
}
__typedArg0 := symstr
__typedArg1 := tmp17615
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17617 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_j, Nil)
}
__typedArg0 := sym_j
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17618 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17616, tmp17617)
}
__typedArg0 := tmp17616
__typedArg1 := tmp17617
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17619 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp17618)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp17618
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17620 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17611, tmp17619)
}
__typedArg0 := tmp17611
__typedArg1 := tmp17619
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17621 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_i, tmp17620)
}
__typedArg0 := sym_i
__typedArg1 := tmp17620
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17622 := Call(__e, PrimFunc(symshen_4comb), let__17389, tmp17621)


ifres17609 = tmp17622


} else {
tmp17623 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17609 = tmp17623


}

ifres17606 = ifres17609


} else {
tmp17628 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17606 = tmp17628


}

ifres17603 = ifres17606


} else {
tmp17630 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17603 = tmp17630


}

ifres17601 = ifres17603


} else {
tmp17632 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17601 = tmp17632


}

ifres17598 = ifres17601


}

ifres17594 = ifres17598


} else {
tmp17635 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17594 = tmp17635


}

ifres17592 = ifres17594


} else {
tmp17637 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17592 = tmp17637


}

ifres17589 = ifres17592


} else {
tmp17639 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17589 = tmp17639


}

ifres17586 = ifres17589


} else {
tmp17641 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17586 = tmp17641


}

let__17375 := ifres17586
_ = let__17375

tmp17652 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17375)


if True == tmp17652 {
tmp17643 := Call(__e, PrimFunc(sym_5e_6), V145)


let__17391 := tmp17643
_ = let__17391

tmp17648 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17391)


var ifres17644 Obj

if True == tmp17648 {
tmp17645 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17644 = tmp17645


} else {
tmp17646 := Call(__e, PrimFunc(symshen_4in_1_6), let__17391)


let__17392 := tmp17646
_ = let__17392

tmp17647 := Call(__e, PrimFunc(symshen_4comb), let__17392, Nil)


ifres17644 = tmp17647


}

let__17390 := ifres17644
_ = let__17390

tmp17650 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17390)


if True == tmp17650 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__17390)
return
}


} else {
__e.Return(let__17375)
return
}


}, 1)

tmp17653 := Call(__e, ns2_1set, symshen_4_5yaccsig_6, tmp17585)


_ = tmp17653

tmp17654 := MakeNative(func(__e *ControlFlow) {
V164 := __e.Get(1)
_ = V164
tmp17655 := Call(__e, PrimFunc(symshen_4_5c_1rule_6), V164)


let__17394 := tmp17655
_ = let__17394

tmp17668 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17394)


var ifres17656 Obj

if True == tmp17668 {
tmp17657 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17656 = tmp17657


} else {
tmp17658 := Call(__e, PrimFunc(symshen_4_5_1out), let__17394)


let__17395 := tmp17658
_ = let__17395

tmp17659 := Call(__e, PrimFunc(symshen_4in_1_6), let__17394)


let__17396 := tmp17659
_ = let__17396

tmp17660 := Call(__e, PrimFunc(symshen_4_5c_1rules_6), let__17396)


let__17397 := tmp17660
_ = let__17397

tmp17667 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17397)


var ifres17661 Obj

if True == tmp17667 {
tmp17662 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17661 = tmp17662


} else {
tmp17663 := Call(__e, PrimFunc(symshen_4_5_1out), let__17397)


let__17398 := tmp17663
_ = let__17398

tmp17664 := Call(__e, PrimFunc(symshen_4in_1_6), let__17397)


let__17399 := tmp17664
_ = let__17399

tmp17665 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17395, let__17398)
}
__typedArg0 := let__17395
__typedArg1 := let__17398
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17666 := Call(__e, PrimFunc(symshen_4comb), let__17399, tmp17665)


ifres17661 = tmp17666


}

ifres17656 = ifres17661


}

let__17393 := ifres17656
_ = let__17393

tmp17684 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17393)


if True == tmp17684 {
tmp17669 := Call(__e, PrimFunc(sym_5_b_6), V164)


let__17401 := tmp17669
_ = let__17401

tmp17680 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17401)


var ifres17670 Obj

if True == tmp17680 {
tmp17671 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17670 = tmp17671


} else {
tmp17672 := Call(__e, PrimFunc(symshen_4_5_1out), let__17401)


let__17402 := tmp17672
_ = let__17402

tmp17673 := Call(__e, PrimFunc(symshen_4in_1_6), let__17401)


let__17403 := tmp17673
_ = let__17403

tmp17678 := Call(__e, PrimFunc(symempty_2), let__17402)


var ifres17674 Obj

if True == tmp17678 {
ifres17674 = Nil


} else {
tmp17675 := Call(__e, PrimFunc(symshen_4app), let__17402, MakeString("\n ..."), symshen_4r)


tmp17677 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("YACC syntax error here:\n "))
__typedS1, __typedOK1 := TypedString(tmp17675)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("YACC syntax error here:\n ")
__typedArg1 := tmp17675
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("YACC syntax error here:\n "))
__typedS1, __typedOK1 := TypedString(tmp17675)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("YACC syntax error here:\n ")
__typedArg1 := tmp17675
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})()

ifres17674 = tmp17677


}

tmp17679 := Call(__e, PrimFunc(symshen_4comb), let__17403, ifres17674)


ifres17670 = tmp17679


}

let__17400 := ifres17670
_ = let__17400

tmp17682 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17400)


if True == tmp17682 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__17400)
return
}


} else {
__e.Return(let__17393)
return
}


}, 1)

tmp17685 := Call(__e, ns2_1set, symshen_4_5c_1rules_6, tmp17654)


_ = tmp17685

tmp17686 := MakeNative(func(__e *ControlFlow) {
V176 := __e.Get(1)
_ = V176
tmp17687 := Call(__e, PrimFunc(symshen_4_5syntax_6), V176)


let__17405 := tmp17687
_ = let__17405

tmp17706 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17405)


var ifres17688 Obj

if True == tmp17706 {
tmp17689 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17688 = tmp17689


} else {
tmp17690 := Call(__e, PrimFunc(symshen_4_5_1out), let__17405)


let__17406 := tmp17690
_ = let__17406

tmp17691 := Call(__e, PrimFunc(symshen_4in_1_6), let__17405)


let__17407 := tmp17691
_ = let__17407

tmp17692 := Call(__e, PrimFunc(symshen_4_5semantics_6), let__17407)


let__17408 := tmp17692
_ = let__17408

tmp17705 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17408)


var ifres17693 Obj

if True == tmp17705 {
tmp17694 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17693 = tmp17694


} else {
tmp17695 := Call(__e, PrimFunc(symshen_4_5_1out), let__17408)


let__17409 := tmp17695
_ = let__17409

tmp17696 := Call(__e, PrimFunc(symshen_4in_1_6), let__17408)


let__17410 := tmp17696
_ = let__17410

tmp17697 := Call(__e, PrimFunc(symshen_4_5sc_6), let__17410)


let__17411 := tmp17697
_ = let__17411

tmp17704 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17411)


var ifres17698 Obj

if True == tmp17704 {
tmp17699 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17698 = tmp17699


} else {
tmp17700 := Call(__e, PrimFunc(symshen_4in_1_6), let__17411)


let__17412 := tmp17700
_ = let__17412

tmp17701 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17409, Nil)
}
__typedArg0 := let__17409
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17702 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17406, tmp17701)
}
__typedArg0 := let__17406
__typedArg1 := tmp17701
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17703 := Call(__e, PrimFunc(symshen_4comb), let__17412, tmp17702)


ifres17698 = tmp17703


}

ifres17693 = ifres17698


}

ifres17688 = ifres17693


}

let__17404 := ifres17688
_ = let__17404

tmp17725 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17404)


if True == tmp17725 {
tmp17707 := Call(__e, PrimFunc(symshen_4_5syntax_6), V176)


let__17414 := tmp17707
_ = let__17414

tmp17721 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17414)


var ifres17708 Obj

if True == tmp17721 {
tmp17709 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17708 = tmp17709


} else {
tmp17710 := Call(__e, PrimFunc(symshen_4_5_1out), let__17414)


let__17415 := tmp17710
_ = let__17415

tmp17711 := Call(__e, PrimFunc(symshen_4in_1_6), let__17414)


let__17416 := tmp17711
_ = let__17416

tmp17712 := Call(__e, PrimFunc(symshen_4_5sc_6), let__17416)


let__17417 := tmp17712
_ = let__17417

tmp17720 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17417)


var ifres17713 Obj

if True == tmp17720 {
tmp17714 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17713 = tmp17714


} else {
tmp17715 := Call(__e, PrimFunc(symshen_4in_1_6), let__17417)


let__17418 := tmp17715
_ = let__17418

tmp17716 := Call(__e, PrimFunc(symshen_4autocomplete), let__17415)


tmp17717 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17716, Nil)
}
__typedArg0 := tmp17716
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17718 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17415, tmp17717)
}
__typedArg0 := let__17415
__typedArg1 := tmp17717
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17719 := Call(__e, PrimFunc(symshen_4comb), let__17418, tmp17718)


ifres17713 = tmp17719


}

ifres17708 = ifres17713


}

let__17413 := ifres17708
_ = let__17413

tmp17723 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17413)


if True == tmp17723 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__17413)
return
}


} else {
__e.Return(let__17404)
return
}


}, 1)

tmp17726 := Call(__e, ns2_1set, symshen_4_5c_1rule_6, tmp17686)


_ = tmp17726

tmp17727 := MakeNative(func(__e *ControlFlow) {
V192 := __e.Get(1)
_ = V192
tmp17756 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V192)
}
__typedArg0 := V192
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17748 Obj

if True == tmp17756 {
tmp17754 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V192)
}
__typedArg0 := V192
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17755 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp17754)
}
__typedArg0 := Nil
__typedArg1 := tmp17754
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17750 Obj

if True == tmp17755 {
tmp17752 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V192)
}
__typedArg0 := V192
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17753 := Call(__e, PrimFunc(symshen_4non_1terminal_2), tmp17752)


var ifres17751 Obj

if True == tmp17753 {
ifres17751 = True


} else {
ifres17751 = False


}

ifres17750 = ifres17751


} else {
ifres17750 = False


}

var ifres17749 Obj

if True == ifres17750 {
ifres17749 = True


} else {
ifres17749 = False


}

ifres17748 = ifres17749


} else {
ifres17748 = False


}

if True == ifres17748 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V192)
}
__typedArg0 := V192
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return
} else {
tmp17746 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V192)
}
__typedArg0 := V192
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17742 Obj

if True == tmp17746 {
tmp17744 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V192)
}
__typedArg0 := V192
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17745 := Call(__e, PrimFunc(symshen_4non_1terminal_2), tmp17744)


var ifres17743 Obj

if True == tmp17745 {
ifres17743 = True


} else {
ifres17743 = False


}

ifres17742 = ifres17743


} else {
ifres17742 = False


}

if True == ifres17742 {
tmp17728 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V192)
}
__typedArg0 := V192
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17729 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V192)
}
__typedArg0 := V192
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17730 := Call(__e, PrimFunc(symshen_4autocomplete), tmp17729)


tmp17731 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17730, Nil)
}
__typedArg0 := tmp17730
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17732 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17728, tmp17731)
}
__typedArg0 := tmp17728
__typedArg1 := tmp17731
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symappend, tmp17732)
}
__typedArg0 := symappend
__typedArg1 := tmp17732
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp17740 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V192)
}
__typedArg0 := V192
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17740 {
tmp17733 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V192)
}
__typedArg0 := V192
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17734 := Call(__e, PrimFunc(symshen_4autocomplete), tmp17733)


tmp17735 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V192)
}
__typedArg0 := V192
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17736 := Call(__e, PrimFunc(symshen_4autocomplete), tmp17735)


tmp17737 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17736, Nil)
}
__typedArg0 := tmp17736
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17738 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17734, tmp17737)
}
__typedArg0 := tmp17734
__typedArg1 := tmp17737
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp17738)
}
__typedArg0 := symcons
__typedArg1 := tmp17738
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return(V192)
return
}


}


}


}, 1)

tmp17757 := Call(__e, ns2_1set, symshen_4autocomplete, tmp17727)


_ = tmp17757

tmp17758 := MakeNative(func(__e *ControlFlow) {
V193 := __e.Get(1)
_ = V193
tmp17764 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsymbol_2) {
return PrimIsSymbol(V193)
}
__typedArg0 := V193
return Call(__e, PrimFunc(symsymbol_2), __typedArg0)
})()

if True == tmp17764 {
tmp17760 := Call(__e, PrimFunc(symexplode), V193)


let__17419 := tmp17760
_ = let__17419

tmp17761 := MakeNative(func(__e *ControlFlow) {
Z195 := __e.Get(1)
_ = Z195
__e.TailApply(PrimFunc(symshen_4_5non_1terminal_2_6), Z195)
return
}, 1)

tmp17762 := Call(__e, PrimFunc(symcompile), tmp17761, let__17419)


if True == tmp17762 {
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

tmp17765 := Call(__e, ns2_1set, symshen_4non_1terminal_2, tmp17758)


_ = tmp17765

tmp17766 := MakeNative(func(__e *ControlFlow) {
V196 := __e.Get(1)
_ = V196
tmp17767 := Call(__e, PrimFunc(symshen_4_5packagenames_6), V196)


let__17421 := tmp17767
_ = let__17421

tmp17777 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17421)


var ifres17768 Obj

if True == tmp17777 {
tmp17769 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17768 = tmp17769


} else {
tmp17770 := Call(__e, PrimFunc(symshen_4in_1_6), let__17421)


let__17422 := tmp17770
_ = let__17422

tmp17771 := Call(__e, PrimFunc(symshen_4_5non_1terminal_1name_6), let__17422)


let__17423 := tmp17771
_ = let__17423

tmp17776 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17423)


var ifres17772 Obj

if True == tmp17776 {
tmp17773 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17772 = tmp17773


} else {
tmp17774 := Call(__e, PrimFunc(symshen_4in_1_6), let__17423)


let__17424 := tmp17774
_ = let__17424

tmp17775 := Call(__e, PrimFunc(symshen_4comb), let__17424, True)


ifres17772 = tmp17775


}

ifres17768 = ifres17772


}

let__17420 := ifres17768
_ = let__17420

tmp17795 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17420)


if True == tmp17795 {
tmp17778 := Call(__e, PrimFunc(symshen_4_5non_1terminal_1name_6), V196)


let__17426 := tmp17778
_ = let__17426

tmp17783 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17426)


var ifres17779 Obj

if True == tmp17783 {
tmp17780 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17779 = tmp17780


} else {
tmp17781 := Call(__e, PrimFunc(symshen_4in_1_6), let__17426)


let__17427 := tmp17781
_ = let__17427

tmp17782 := Call(__e, PrimFunc(symshen_4comb), let__17427, True)


ifres17779 = tmp17782


}

let__17425 := ifres17779
_ = let__17425

tmp17793 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17425)


if True == tmp17793 {
tmp17784 := Call(__e, PrimFunc(sym_5_b_6), V196)


let__17429 := tmp17784
_ = let__17429

tmp17789 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17429)


var ifres17785 Obj

if True == tmp17789 {
tmp17786 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17785 = tmp17786


} else {
tmp17787 := Call(__e, PrimFunc(symshen_4in_1_6), let__17429)


let__17430 := tmp17787
_ = let__17430

tmp17788 := Call(__e, PrimFunc(symshen_4comb), let__17430, False)


ifres17785 = tmp17788


}

let__17428 := ifres17785
_ = let__17428

tmp17791 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17428)


if True == tmp17791 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__17428)
return
}


} else {
__e.Return(let__17425)
return
}


} else {
__e.Return(let__17420)
return
}


}, 1)

tmp17796 := Call(__e, ns2_1set, symshen_4_5non_1terminal_2_6, tmp17766)


_ = tmp17796

tmp17797 := MakeNative(func(__e *ControlFlow) {
V208 := __e.Get(1)
_ = V208
tmp17798 := Call(__e, PrimFunc(symshen_4_5packagename_6), V208)


let__17432 := tmp17798
_ = let__17432

tmp17812 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17432)


var ifres17799 Obj

if True == tmp17812 {
tmp17800 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17799 = tmp17800


} else {
tmp17801 := Call(__e, PrimFunc(symshen_4in_1_6), let__17432)


let__17433 := tmp17801
_ = let__17433

tmp17811 := Call(__e, PrimFunc(symshen_4hds_a_2), let__17433, MakeString("."))


var ifres17802 Obj

if True == tmp17811 {
tmp17803 := Call(__e, PrimFunc(symtail), let__17433)


let__17434 := tmp17803
_ = let__17434

tmp17804 := Call(__e, PrimFunc(symshen_4_5packagenames_6), let__17434)


let__17435 := tmp17804
_ = let__17435

tmp17809 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17435)


var ifres17805 Obj

if True == tmp17809 {
tmp17806 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17805 = tmp17806


} else {
tmp17807 := Call(__e, PrimFunc(symshen_4in_1_6), let__17435)


let__17436 := tmp17807
_ = let__17436

tmp17808 := Call(__e, PrimFunc(symshen_4comb), let__17436, symshen_4skip)


ifres17805 = tmp17808


}

ifres17802 = ifres17805


} else {
tmp17810 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17802 = tmp17810


}

ifres17799 = ifres17802


}

let__17431 := ifres17799
_ = let__17431

tmp17826 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17431)


if True == tmp17826 {
tmp17813 := Call(__e, PrimFunc(symshen_4_5packagename_6), V208)


let__17438 := tmp17813
_ = let__17438

tmp17822 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17438)


var ifres17814 Obj

if True == tmp17822 {
tmp17815 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17814 = tmp17815


} else {
tmp17816 := Call(__e, PrimFunc(symshen_4in_1_6), let__17438)


let__17439 := tmp17816
_ = let__17439

tmp17821 := Call(__e, PrimFunc(symshen_4hds_a_2), let__17439, MakeString("."))


var ifres17817 Obj

if True == tmp17821 {
tmp17818 := Call(__e, PrimFunc(symtail), let__17439)


let__17440 := tmp17818
_ = let__17440

tmp17819 := Call(__e, PrimFunc(symshen_4comb), let__17440, symshen_4skip)


ifres17817 = tmp17819


} else {
tmp17820 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17817 = tmp17820


}

ifres17814 = ifres17817


}

let__17437 := ifres17814
_ = let__17437

tmp17824 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17437)


if True == tmp17824 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__17437)
return
}


} else {
__e.Return(let__17431)
return
}


}, 1)

tmp17827 := Call(__e, ns2_1set, symshen_4_5packagenames_6, tmp17797)


_ = tmp17827

tmp17828 := MakeNative(func(__e *ControlFlow) {
V219 := __e.Get(1)
_ = V219
tmp17829 := Call(__e, PrimFunc(symshen_4_5packagechar_6), V219)


let__17442 := tmp17829
_ = let__17442

tmp17839 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17442)


var ifres17830 Obj

if True == tmp17839 {
tmp17831 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17830 = tmp17831


} else {
tmp17832 := Call(__e, PrimFunc(symshen_4in_1_6), let__17442)


let__17443 := tmp17832
_ = let__17443

tmp17833 := Call(__e, PrimFunc(symshen_4_5packagename_6), let__17443)


let__17444 := tmp17833
_ = let__17444

tmp17838 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17444)


var ifres17834 Obj

if True == tmp17838 {
tmp17835 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17834 = tmp17835


} else {
tmp17836 := Call(__e, PrimFunc(symshen_4in_1_6), let__17444)


let__17445 := tmp17836
_ = let__17445

tmp17837 := Call(__e, PrimFunc(symshen_4comb), let__17445, symshen_4skip)


ifres17834 = tmp17837


}

ifres17830 = ifres17834


}

let__17441 := ifres17830
_ = let__17441

tmp17849 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17441)


if True == tmp17849 {
tmp17840 := Call(__e, PrimFunc(sym_5e_6), V219)


let__17447 := tmp17840
_ = let__17447

tmp17845 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17447)


var ifres17841 Obj

if True == tmp17845 {
tmp17842 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17841 = tmp17842


} else {
tmp17843 := Call(__e, PrimFunc(symshen_4in_1_6), let__17447)


let__17448 := tmp17843
_ = let__17448

tmp17844 := Call(__e, PrimFunc(symshen_4comb), let__17448, symshen_4skip)


ifres17841 = tmp17844


}

let__17446 := ifres17841
_ = let__17446

tmp17847 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17446)


if True == tmp17847 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__17446)
return
}


} else {
__e.Return(let__17441)
return
}


}, 1)

tmp17850 := Call(__e, ns2_1set, symshen_4_5packagename_6, tmp17828)


_ = tmp17850

tmp17851 := MakeNative(func(__e *ControlFlow) {
V228 := __e.Get(1)
_ = V228
tmp17861 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V228)
}
__typedArg0 := V228
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17852 Obj

if True == tmp17861 {
tmp17853 := Call(__e, PrimFunc(symhead), V228)


let__17450 := tmp17853
_ = let__17450

tmp17854 := Call(__e, PrimFunc(symtail), V228)


let__17451 := tmp17854
_ = let__17451

tmp17858 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__17450, MakeString("."))
}
__typedArg0 := let__17450
__typedArg1 := MakeString(".")
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

tmp17859 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp17858)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp17858
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres17855 Obj

if True == tmp17859 {
tmp17856 := Call(__e, PrimFunc(symshen_4comb), let__17451, symshen_4skip)


ifres17855 = tmp17856


} else {
tmp17857 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17855 = tmp17857


}

ifres17852 = ifres17855


} else {
tmp17860 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17852 = tmp17860


}

let__17449 := ifres17852
_ = let__17449

tmp17863 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17449)


if True == tmp17863 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__17449)
return
}


}, 1)

tmp17864 := Call(__e, ns2_1set, symshen_4_5packagechar_6, tmp17851)


_ = tmp17864

tmp17865 := MakeNative(func(__e *ControlFlow) {
V232 := __e.Get(1)
_ = V232
tmp17884 := Call(__e, PrimFunc(symshen_4hds_a_2), V232, MakeString("<"))


var ifres17866 Obj

if True == tmp17884 {
tmp17867 := Call(__e, PrimFunc(symtail), V232)


let__17453 := tmp17867
_ = let__17453

tmp17868 := Call(__e, PrimFunc(sym_5_b_6), let__17453)


let__17454 := tmp17868
_ = let__17454

tmp17882 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17454)


var ifres17869 Obj

if True == tmp17882 {
tmp17870 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17869 = tmp17870


} else {
tmp17871 := Call(__e, PrimFunc(symshen_4_5_1out), let__17454)


let__17455 := tmp17871
_ = let__17455

tmp17872 := Call(__e, PrimFunc(symshen_4in_1_6), let__17454)


let__17456 := tmp17872
_ = let__17456

tmp17876 := Call(__e, PrimFunc(symreverse), let__17455)


let__17457 := tmp17876
_ = let__17457

tmp17881 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__17457)
}
__typedArg0 := let__17457
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17877 Obj

if True == tmp17881 {
tmp17879 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__17457)
}
__typedArg0 := let__17457
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17880 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp17879, MakeString(">"))
}
__typedArg0 := tmp17879
__typedArg1 := MakeString(">")
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17878 Obj

if True == tmp17880 {
ifres17878 = True


} else {
ifres17878 = False


}

ifres17877 = ifres17878


} else {
ifres17877 = False


}

var ifres17873 Obj

if True == ifres17877 {
tmp17874 := Call(__e, PrimFunc(symshen_4comb), let__17456, symshen_4skip)


ifres17873 = tmp17874


} else {
tmp17875 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17873 = tmp17875


}

ifres17869 = ifres17873


}

ifres17866 = ifres17869


} else {
tmp17883 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17866 = tmp17883


}

let__17452 := ifres17866
_ = let__17452

tmp17886 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17452)


if True == tmp17886 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__17452)
return
}


}, 1)

tmp17887 := Call(__e, ns2_1set, symshen_4_5non_1terminal_1name_6, tmp17865)


_ = tmp17887

tmp17888 := MakeNative(func(__e *ControlFlow) {
V239 := __e.Get(1)
_ = V239
tmp17889 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(";"))
}
__typedArg0 := MakeString(";")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V239, tmp17889)
}
__typedArg0 := V239
__typedArg1 := tmp17889
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp17890 := Call(__e, ns2_1set, symshen_4semicolon_2, tmp17888)


_ = tmp17890

tmp17891 := MakeNative(func(__e *ControlFlow) {
V240 := __e.Get(1)
_ = V240
tmp17900 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V240)
}
__typedArg0 := V240
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17892 Obj

if True == tmp17900 {
tmp17893 := Call(__e, PrimFunc(symhead), V240)


let__17459 := tmp17893
_ = let__17459

tmp17894 := Call(__e, PrimFunc(symtail), V240)


let__17460 := tmp17894
_ = let__17460

tmp17898 := Call(__e, PrimFunc(symshen_4colon_1equal_2), let__17459)


var ifres17895 Obj

if True == tmp17898 {
tmp17896 := Call(__e, PrimFunc(symshen_4comb), let__17460, symshen_4skip)


ifres17895 = tmp17896


} else {
tmp17897 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17895 = tmp17897


}

ifres17892 = ifres17895


} else {
tmp17899 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17892 = tmp17899


}

let__17458 := ifres17892
_ = let__17458

tmp17902 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17458)


if True == tmp17902 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__17458)
return
}


}, 1)

tmp17903 := Call(__e, ns2_1set, symshen_4_5colon_1equal_6, tmp17891)


_ = tmp17903

tmp17904 := MakeNative(func(__e *ControlFlow) {
V244 := __e.Get(1)
_ = V244
tmp17905 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":="))
}
__typedArg0 := MakeString(":=")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp17905, V244)
}
__typedArg0 := tmp17905
__typedArg1 := V244
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp17906 := Call(__e, ns2_1set, symshen_4colon_1equal_2, tmp17904)


_ = tmp17906

tmp17907 := MakeNative(func(__e *ControlFlow) {
V245 := __e.Get(1)
_ = V245
tmp17908 := Call(__e, PrimFunc(symshen_4_5syntax_1item_6), V245)


let__17462 := tmp17908
_ = let__17462

tmp17921 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17462)


var ifres17909 Obj

if True == tmp17921 {
tmp17910 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17909 = tmp17910


} else {
tmp17911 := Call(__e, PrimFunc(symshen_4_5_1out), let__17462)


let__17463 := tmp17911
_ = let__17463

tmp17912 := Call(__e, PrimFunc(symshen_4in_1_6), let__17462)


let__17464 := tmp17912
_ = let__17464

tmp17913 := Call(__e, PrimFunc(symshen_4_5syntax_6), let__17464)


let__17465 := tmp17913
_ = let__17465

tmp17920 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17465)


var ifres17914 Obj

if True == tmp17920 {
tmp17915 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17914 = tmp17915


} else {
tmp17916 := Call(__e, PrimFunc(symshen_4_5_1out), let__17465)


let__17466 := tmp17916
_ = let__17466

tmp17917 := Call(__e, PrimFunc(symshen_4in_1_6), let__17465)


let__17467 := tmp17917
_ = let__17467

tmp17918 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17463, let__17466)
}
__typedArg0 := let__17463
__typedArg1 := let__17466
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17919 := Call(__e, PrimFunc(symshen_4comb), let__17467, tmp17918)


ifres17914 = tmp17919


}

ifres17909 = ifres17914


}

let__17461 := ifres17909
_ = let__17461

tmp17933 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17461)


if True == tmp17933 {
tmp17922 := Call(__e, PrimFunc(symshen_4_5syntax_1item_6), V245)


let__17469 := tmp17922
_ = let__17469

tmp17929 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17469)


var ifres17923 Obj

if True == tmp17929 {
tmp17924 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17923 = tmp17924


} else {
tmp17925 := Call(__e, PrimFunc(symshen_4_5_1out), let__17469)


let__17470 := tmp17925
_ = let__17470

tmp17926 := Call(__e, PrimFunc(symshen_4in_1_6), let__17469)


let__17471 := tmp17926
_ = let__17471

tmp17927 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17470, Nil)
}
__typedArg0 := let__17470
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17928 := Call(__e, PrimFunc(symshen_4comb), let__17471, tmp17927)


ifres17923 = tmp17928


}

let__17468 := ifres17923
_ = let__17468

tmp17931 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17468)


if True == tmp17931 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__17468)
return
}


} else {
__e.Return(let__17461)
return
}


}, 1)

tmp17934 := Call(__e, ns2_1set, symshen_4_5syntax_6, tmp17907)


_ = tmp17934

tmp17935 := MakeNative(func(__e *ControlFlow) {
V257 := __e.Get(1)
_ = V257
tmp17944 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V257)
}
__typedArg0 := V257
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17936 Obj

if True == tmp17944 {
tmp17937 := Call(__e, PrimFunc(symhead), V257)


let__17473 := tmp17937
_ = let__17473

tmp17938 := Call(__e, PrimFunc(symtail), V257)


let__17474 := tmp17938
_ = let__17474

tmp17942 := Call(__e, PrimFunc(symshen_4syntax_1item_2), let__17473)


var ifres17939 Obj

if True == tmp17942 {
tmp17940 := Call(__e, PrimFunc(symshen_4comb), let__17474, let__17473)


ifres17939 = tmp17940


} else {
tmp17941 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17939 = tmp17941


}

ifres17936 = ifres17939


} else {
tmp17943 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17936 = tmp17943


}

let__17472 := ifres17936
_ = let__17472

tmp17946 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17472)


if True == tmp17946 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__17472)
return
}


}, 1)

tmp17947 := Call(__e, ns2_1set, symshen_4_5syntax_1item_6, tmp17935)


_ = tmp17947

tmp17948 := MakeNative(func(__e *ControlFlow) {
V263 := __e.Get(1)
_ = V263
tmp17984 := Call(__e, PrimFunc(symshen_4colon_1equal_2), V263)


if True == tmp17984 {
__e.Return(False)
return
} else {
tmp17982 := Call(__e, PrimFunc(symshen_4semicolon_2), V263)


if True == tmp17982 {
__e.Return(False)
return
} else {
tmp17980 := Call(__e, PrimFunc(symatom_2), V263)


if True == tmp17980 {
__e.Return(True)
return
} else {
tmp17978 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V263)
}
__typedArg0 := V263
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17959 Obj

if True == tmp17978 {
tmp17976 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V263)
}
__typedArg0 := V263
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17977 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcons, tmp17976)
}
__typedArg0 := symcons
__typedArg1 := tmp17976
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17961 Obj

if True == tmp17977 {
tmp17974 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V263)
}
__typedArg0 := V263
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17975 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp17974)
}
__typedArg0 := tmp17974
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17963 Obj

if True == tmp17975 {
tmp17971 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V263)
}
__typedArg0 := V263
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17972 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp17971)
}
__typedArg0 := tmp17971
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17973 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp17972)
}
__typedArg0 := tmp17972
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17965 Obj

if True == tmp17973 {
tmp17967 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V263)
}
__typedArg0 := V263
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17968 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp17967)
}
__typedArg0 := tmp17967
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17969 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp17968)
}
__typedArg0 := tmp17968
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17970 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp17969)
}
__typedArg0 := Nil
__typedArg1 := tmp17969
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17966 Obj

if True == tmp17970 {
ifres17966 = True


} else {
ifres17966 = False


}

ifres17965 = ifres17966


} else {
ifres17965 = False


}

var ifres17964 Obj

if True == ifres17965 {
ifres17964 = True


} else {
ifres17964 = False


}

ifres17963 = ifres17964


} else {
ifres17963 = False


}

var ifres17962 Obj

if True == ifres17963 {
ifres17962 = True


} else {
ifres17962 = False


}

ifres17961 = ifres17962


} else {
ifres17961 = False


}

var ifres17960 Obj

if True == ifres17961 {
ifres17960 = True


} else {
ifres17960 = False


}

ifres17959 = ifres17960


} else {
ifres17959 = False


}

if True == ifres17959 {
tmp17955 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V263)
}
__typedArg0 := V263
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17956 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp17955)
}
__typedArg0 := tmp17955
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17957 := Call(__e, PrimFunc(symshen_4syntax_1item_2), tmp17956)


if True == tmp17957 {
tmp17950 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V263)
}
__typedArg0 := V263
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17951 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp17950)
}
__typedArg0 := tmp17950
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17952 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp17951)
}
__typedArg0 := tmp17951
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17953 := Call(__e, PrimFunc(symshen_4syntax_1item_2), tmp17952)


if True == tmp17953 {
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


} else {
__e.Return(False)
return
}


}


}


}


}, 1)

tmp17985 := Call(__e, ns2_1set, symshen_4syntax_1item_2, tmp17948)


_ = tmp17985

tmp17986 := MakeNative(func(__e *ControlFlow) {
V264 := __e.Get(1)
_ = V264
tmp17987 := Call(__e, PrimFunc(symshen_4_5colon_1equal_6), V264)


let__17476 := tmp17987
_ = let__17476

tmp18013 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17476)


var ifres17988 Obj

if True == tmp18013 {
tmp17989 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17988 = tmp17989


} else {
tmp17990 := Call(__e, PrimFunc(symshen_4in_1_6), let__17476)


let__17477 := tmp17990
_ = let__17477

tmp18012 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__17477)
}
__typedArg0 := let__17477
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17991 Obj

if True == tmp18012 {
tmp17992 := Call(__e, PrimFunc(symhead), let__17477)


let__17478 := tmp17992
_ = let__17478

tmp17993 := Call(__e, PrimFunc(symtail), let__17477)


let__17479 := tmp17993
_ = let__17479

tmp18010 := Call(__e, PrimFunc(symshen_4hds_a_2), let__17479, symwhere)


var ifres17994 Obj

if True == tmp18010 {
tmp17995 := Call(__e, PrimFunc(symtail), let__17479)


let__17480 := tmp17995
_ = let__17480

tmp18008 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__17480)
}
__typedArg0 := let__17480
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17996 Obj

if True == tmp18008 {
tmp17997 := Call(__e, PrimFunc(symhead), let__17480)


let__17481 := tmp17997
_ = let__17481

tmp17998 := Call(__e, PrimFunc(symtail), let__17480)


let__17482 := tmp17998
_ = let__17482

tmp18005 := Call(__e, PrimFunc(symshen_4semicolon_2), let__17478)


tmp18006 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp18005)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp18005
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres17999 Obj

if True == tmp18006 {
tmp18000 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17478, Nil)
}
__typedArg0 := let__17478
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18001 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17481, tmp18000)
}
__typedArg0 := let__17481
__typedArg1 := tmp18000
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18002 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhere, tmp18001)
}
__typedArg0 := symwhere
__typedArg1 := tmp18001
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18003 := Call(__e, PrimFunc(symshen_4comb), let__17482, tmp18002)


ifres17999 = tmp18003


} else {
tmp18004 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17999 = tmp18004


}

ifres17996 = ifres17999


} else {
tmp18007 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17996 = tmp18007


}

ifres17994 = ifres17996


} else {
tmp18009 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17994 = tmp18009


}

ifres17991 = ifres17994


} else {
tmp18011 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17991 = tmp18011


}

ifres17988 = ifres17991


}

let__17475 := ifres17988
_ = let__17475

tmp18032 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17475)


if True == tmp18032 {
tmp18014 := Call(__e, PrimFunc(symshen_4_5colon_1equal_6), V264)


let__17484 := tmp18014
_ = let__17484

tmp18028 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17484)


var ifres18015 Obj

if True == tmp18028 {
tmp18016 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres18015 = tmp18016


} else {
tmp18017 := Call(__e, PrimFunc(symshen_4in_1_6), let__17484)


let__17485 := tmp18017
_ = let__17485

tmp18027 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__17485)
}
__typedArg0 := let__17485
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18018 Obj

if True == tmp18027 {
tmp18019 := Call(__e, PrimFunc(symhead), let__17485)


let__17486 := tmp18019
_ = let__17486

tmp18020 := Call(__e, PrimFunc(symtail), let__17485)


let__17487 := tmp18020
_ = let__17487

tmp18024 := Call(__e, PrimFunc(symshen_4semicolon_2), let__17486)


tmp18025 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp18024)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp18024
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres18021 Obj

if True == tmp18025 {
tmp18022 := Call(__e, PrimFunc(symshen_4comb), let__17487, let__17486)


ifres18021 = tmp18022


} else {
tmp18023 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres18021 = tmp18023


}

ifres18018 = ifres18021


} else {
tmp18026 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres18018 = tmp18026


}

ifres18015 = ifres18018


}

let__17483 := ifres18015
_ = let__17483

tmp18030 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17483)


if True == tmp18030 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__17483)
return
}


} else {
__e.Return(let__17475)
return
}


}, 1)

tmp18033 := Call(__e, ns2_1set, symshen_4_5semantics_6, tmp17986)


_ = tmp18033

tmp18034 := MakeNative(func(__e *ControlFlow) {
V286 := __e.Get(1)
_ = V286
V287 := __e.Get(2)
_ = V287
V288 := __e.Get(3)
_ = V288
tmp18042 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V288)
}
__typedArg0 := Nil
__typedArg1 := V288
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp18042 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4parse_1failure, Nil)
}
__typedArg0 := symshen_4parse_1failure
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return
} else {
tmp18040 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V288)
}
__typedArg0 := V288
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp18040 {
tmp18035 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V288)
}
__typedArg0 := V288
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18036 := Call(__e, PrimFunc(symshen_4c_1rule_1_6shen), V286, tmp18035, V287)


tmp18037 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V288)
}
__typedArg0 := V288
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18038 := Call(__e, PrimFunc(symshen_4c_1rules_1_6shen), V286, V287, tmp18037)


__e.TailApply(PrimFunc(symshen_4combine_1c_1code), tmp18036, tmp18038)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.c-rules->shen\n"))
}
__typedArg0 := MakeString("implementation error in shen.c-rules->shen\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 3)

tmp18043 := Call(__e, ns2_1set, symshen_4c_1rules_1_6shen, tmp18034)


_ = tmp18043

tmp18044 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symfail))
return
}, 0)

tmp18045 := Call(__e, ns2_1set, symshen_4parse_1failure, tmp18044)


_ = tmp18045

tmp18046 := MakeNative(func(__e *ControlFlow) {
V289 := __e.Get(1)
_ = V289
V290 := __e.Get(2)
_ = V290
tmp18047 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symResult, Nil)
}
__typedArg0 := symResult
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18048 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4parse_1failure_2, tmp18047)
}
__typedArg0 := symshen_4parse_1failure_2
__typedArg1 := tmp18047
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18049 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symResult, Nil)
}
__typedArg0 := symResult
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18050 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V290, tmp18049)
}
__typedArg0 := V290
__typedArg1 := tmp18049
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18051 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18048, tmp18050)
}
__typedArg0 := tmp18048
__typedArg1 := tmp18050
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18052 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp18051)
}
__typedArg0 := symif
__typedArg1 := tmp18051
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18053 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18052, Nil)
}
__typedArg0 := tmp18052
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18054 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V289, tmp18053)
}
__typedArg0 := V289
__typedArg1 := tmp18053
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18055 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symResult, tmp18054)
}
__typedArg0 := symResult
__typedArg1 := tmp18054
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp18055)
}
__typedArg0 := symlet
__typedArg1 := tmp18055
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 2)

tmp18056 := Call(__e, ns2_1set, symshen_4combine_1c_1code, tmp18046)


_ = tmp18056

tmp18057 := MakeNative(func(__e *ControlFlow) {
V297 := __e.Get(1)
_ = V297
V298 := __e.Get(2)
_ = V298
V299 := __e.Get(3)
_ = V299
tmp18071 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V298)
}
__typedArg0 := V298
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18062 Obj

if True == tmp18071 {
tmp18069 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V298)
}
__typedArg0 := V298
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18070 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18069)
}
__typedArg0 := tmp18069
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18064 Obj

if True == tmp18070 {
tmp18066 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V298)
}
__typedArg0 := V298
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18067 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18066)
}
__typedArg0 := tmp18066
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18068 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp18067)
}
__typedArg0 := Nil
__typedArg1 := tmp18067
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18065 Obj

if True == tmp18068 {
ifres18065 = True


} else {
ifres18065 = False


}

ifres18064 = ifres18065


} else {
ifres18064 = False


}

var ifres18063 Obj

if True == ifres18064 {
ifres18063 = True


} else {
ifres18063 = False


}

ifres18062 = ifres18063


} else {
ifres18062 = False


}

if True == ifres18062 {
tmp18058 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V298)
}
__typedArg0 := V298
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18059 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V298)
}
__typedArg0 := V298
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18060 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18059)
}
__typedArg0 := tmp18059
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4yacc_1syntax), V297, V299, tmp18058, tmp18060)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.c-rule->shen\n"))
}
__typedArg0 := MakeString("implementation error in shen.c-rule->shen\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}, 3)

tmp18072 := Call(__e, ns2_1set, symshen_4c_1rule_1_6shen, tmp18057)


_ = tmp18072

tmp18073 := MakeNative(func(__e *ControlFlow) {
V308 := __e.Get(1)
_ = V308
V309 := __e.Get(2)
_ = V309
V310 := __e.Get(3)
_ = V310
V311 := __e.Get(4)
_ = V311
tmp18137 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V310)
}
__typedArg0 := Nil
__typedArg1 := V310
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18115 Obj

if True == tmp18137 {
tmp18136 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V311)
}
__typedArg0 := V311
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18117 Obj

if True == tmp18136 {
tmp18134 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V311)
}
__typedArg0 := V311
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18135 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symwhere, tmp18134)
}
__typedArg0 := symwhere
__typedArg1 := tmp18134
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18119 Obj

if True == tmp18135 {
tmp18132 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V311)
}
__typedArg0 := V311
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18133 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18132)
}
__typedArg0 := tmp18132
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18121 Obj

if True == tmp18133 {
tmp18129 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V311)
}
__typedArg0 := V311
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18130 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18129)
}
__typedArg0 := tmp18129
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18131 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18130)
}
__typedArg0 := tmp18130
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18123 Obj

if True == tmp18131 {
tmp18125 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V311)
}
__typedArg0 := V311
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18126 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18125)
}
__typedArg0 := tmp18125
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18127 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18126)
}
__typedArg0 := tmp18126
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18128 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp18127)
}
__typedArg0 := Nil
__typedArg1 := tmp18127
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18124 Obj

if True == tmp18128 {
ifres18124 = True


} else {
ifres18124 = False


}

ifres18123 = ifres18124


} else {
ifres18123 = False


}

var ifres18122 Obj

if True == ifres18123 {
ifres18122 = True


} else {
ifres18122 = False


}

ifres18121 = ifres18122


} else {
ifres18121 = False


}

var ifres18120 Obj

if True == ifres18121 {
ifres18120 = True


} else {
ifres18120 = False


}

ifres18119 = ifres18120


} else {
ifres18119 = False


}

var ifres18118 Obj

if True == ifres18119 {
ifres18118 = True


} else {
ifres18118 = False


}

ifres18117 = ifres18118


} else {
ifres18117 = False


}

var ifres18116 Obj

if True == ifres18117 {
ifres18116 = True


} else {
ifres18116 = False


}

ifres18115 = ifres18116


} else {
ifres18115 = False


}

if True == ifres18115 {
tmp18074 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V311)
}
__typedArg0 := V311
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18075 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18074)
}
__typedArg0 := tmp18074
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18076 := Call(__e, PrimFunc(symshen_4process_1yacc_1semantics), tmp18075)


tmp18077 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V311)
}
__typedArg0 := V311
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18078 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18077)
}
__typedArg0 := tmp18077
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18079 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18078)
}
__typedArg0 := tmp18078
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18080 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V308, V309, Nil, tmp18079)


tmp18081 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4parse_1failure, Nil)
}
__typedArg0 := symshen_4parse_1failure
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18082 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18081, Nil)
}
__typedArg0 := tmp18081
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18083 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18080, tmp18082)
}
__typedArg0 := tmp18080
__typedArg1 := tmp18082
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18084 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18076, tmp18083)
}
__typedArg0 := tmp18076
__typedArg1 := tmp18083
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp18084)
}
__typedArg0 := symif
__typedArg1 := tmp18084
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp18113 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V310)
}
__typedArg0 := Nil
__typedArg1 := V310
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp18113 {
__e.TailApply(PrimFunc(symshen_4yacc_1semantics), V308, V309, V311)
return
} else {
tmp18111 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp18111 {
tmp18108 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18109 := Call(__e, PrimFunc(symshen_4non_1terminal_2), tmp18108)


if True == tmp18109 {
tmp18085 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18086 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4non_1terminalcode), V308, V309, tmp18085, tmp18086, V311)
return


} else {
tmp18105 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18106 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(tmp18105)
}
__typedArg0 := tmp18105
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

if True == tmp18106 {
tmp18087 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18088 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4variablecode), V308, V309, tmp18087, tmp18088, V311)
return


} else {
tmp18102 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18103 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym__, tmp18102)
}
__typedArg0 := sym__
__typedArg1 := tmp18102
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp18103 {
tmp18089 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18090 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4wildcardcode), V308, V309, tmp18089, tmp18090, V311)
return


} else {
tmp18099 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18100 := Call(__e, PrimFunc(symatom_2), tmp18099)


if True == tmp18100 {
tmp18091 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18092 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4terminalcode), V308, V309, tmp18091, tmp18092, V311)
return


} else {
tmp18096 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18097 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18096)
}
__typedArg0 := tmp18096
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp18097 {
tmp18093 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18094 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4conscode), V308, V309, tmp18093, tmp18094, V311)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.yacc-syntax\n"))
}
__typedArg0 := MakeString("implementation error in shen.yacc-syntax\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}


}


}


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.yacc-syntax\n"))
}
__typedArg0 := MakeString("implementation error in shen.yacc-syntax\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}


}, 4)

tmp18138 := Call(__e, ns2_1set, symshen_4yacc_1syntax, tmp18073)


_ = tmp18138

tmp18139 := MakeNative(func(__e *ControlFlow) {
V312 := __e.Get(1)
_ = V312
V313 := __e.Get(2)
_ = V313
V314 := __e.Get(3)
_ = V314
V315 := __e.Get(4)
_ = V315
V316 := __e.Get(5)
_ = V316
tmp18140 := Call(__e, PrimFunc(symconcat), symParse, V314)


let__17488 := tmp18140
_ = let__17488

tmp18141 := Call(__e, PrimFunc(symconcat), symAction, V314)


let__17489 := tmp18141
_ = let__17489

tmp18142 := Call(__e, PrimFunc(symconcat), symRemainder, V314)


let__17490 := tmp18142
_ = let__17490

tmp18143 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V313, Nil)
}
__typedArg0 := V313
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18144 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V314, tmp18143)
}
__typedArg0 := V314
__typedArg1 := tmp18143
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18145 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17488, Nil)
}
__typedArg0 := let__17488
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18146 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4parse_1failure_2, tmp18145)
}
__typedArg0 := symshen_4parse_1failure_2
__typedArg1 := tmp18145
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18147 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4parse_1failure, Nil)
}
__typedArg0 := symshen_4parse_1failure
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18148 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17488, Nil)
}
__typedArg0 := let__17488
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18149 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4in_1_6, tmp18148)
}
__typedArg0 := symshen_4in_1_6
__typedArg1 := tmp18148
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18150 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V312, let__17490, V315, V316)


tmp18151 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18150, Nil)
}
__typedArg0 := tmp18150
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18152 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18149, tmp18151)
}
__typedArg0 := tmp18149
__typedArg1 := tmp18151
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18153 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17490, tmp18152)
}
__typedArg0 := let__17490
__typedArg1 := tmp18152
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18154 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp18153)
}
__typedArg0 := symlet
__typedArg1 := tmp18153
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__17491 := tmp18154
_ = let__17491

tmp18165 := Call(__e, PrimFunc(symshen_4occurs_1check_2), V314, V316)


var ifres18162 Obj

if True == tmp18165 {
ifres18162 = True


} else {
tmp18164 := Call(__e, PrimFunc(symshen_4occurs_1check_2), let__17489, V316)


var ifres18163 Obj

if True == tmp18164 {
ifres18163 = True


} else {
ifres18163 = False


}

ifres18162 = ifres18163


}

var ifres18155 Obj

if True == ifres18162 {
tmp18156 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17488, Nil)
}
__typedArg0 := let__17488
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18157 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4_5_1out, tmp18156)
}
__typedArg0 := symshen_4_5_1out
__typedArg1 := tmp18156
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18158 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17491, Nil)
}
__typedArg0 := let__17491
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18159 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18157, tmp18158)
}
__typedArg0 := tmp18157
__typedArg1 := tmp18158
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18160 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17489, tmp18159)
}
__typedArg0 := let__17489
__typedArg1 := tmp18159
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18161 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp18160)
}
__typedArg0 := symlet
__typedArg1 := tmp18160
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

ifres18155 = tmp18161


} else {
ifres18155 = let__17491


}

tmp18166 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(ifres18155, Nil)
}
__typedArg0 := ifres18155
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18167 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18147, tmp18166)
}
__typedArg0 := tmp18147
__typedArg1 := tmp18166
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18168 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18146, tmp18167)
}
__typedArg0 := tmp18146
__typedArg1 := tmp18167
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18169 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp18168)
}
__typedArg0 := symif
__typedArg1 := tmp18168
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18170 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18169, Nil)
}
__typedArg0 := tmp18169
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18171 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18144, tmp18170)
}
__typedArg0 := tmp18144
__typedArg1 := tmp18170
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18172 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17488, tmp18171)
}
__typedArg0 := let__17488
__typedArg1 := tmp18171
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp18172)
}
__typedArg0 := symlet
__typedArg1 := tmp18172
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 5)

tmp18173 := Call(__e, ns2_1set, symshen_4non_1terminalcode, tmp18139)


_ = tmp18173

tmp18174 := MakeNative(func(__e *ControlFlow) {
V321 := __e.Get(1)
_ = V321
V322 := __e.Get(2)
_ = V322
V323 := __e.Get(3)
_ = V323
V324 := __e.Get(4)
_ = V324
V325 := __e.Get(5)
_ = V325
tmp18175 := Call(__e, PrimFunc(symgensym), symRemainder)


let__17492 := tmp18175
_ = let__17492

tmp18176 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V322, Nil)
}
__typedArg0 := V322
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18177 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons_2, tmp18176)
}
__typedArg0 := symcons_2
__typedArg1 := tmp18176
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18178 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V322, Nil)
}
__typedArg0 := V322
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18179 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtail, tmp18178)
}
__typedArg0 := symtail
__typedArg1 := tmp18178
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18180 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V321, let__17492, V324, V325)


tmp18181 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18180, Nil)
}
__typedArg0 := tmp18180
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18182 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18179, tmp18181)
}
__typedArg0 := tmp18179
__typedArg1 := tmp18181
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18183 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17492, tmp18182)
}
__typedArg0 := let__17492
__typedArg1 := tmp18182
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18184 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp18183)
}
__typedArg0 := symlet
__typedArg1 := tmp18183
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__17493 := tmp18184
_ = let__17493

tmp18192 := Call(__e, PrimFunc(symshen_4occurs_1check_2), V323, V325)


var ifres18185 Obj

if True == tmp18192 {
tmp18186 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V322, Nil)
}
__typedArg0 := V322
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18187 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhead, tmp18186)
}
__typedArg0 := symhead
__typedArg1 := tmp18186
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18188 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17493, Nil)
}
__typedArg0 := let__17493
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18189 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18187, tmp18188)
}
__typedArg0 := tmp18187
__typedArg1 := tmp18188
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18190 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V323, tmp18189)
}
__typedArg0 := V323
__typedArg1 := tmp18189
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18191 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp18190)
}
__typedArg0 := symlet
__typedArg1 := tmp18190
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

ifres18185 = tmp18191


} else {
ifres18185 = let__17493


}

tmp18193 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4parse_1failure, Nil)
}
__typedArg0 := symshen_4parse_1failure
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18194 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18193, Nil)
}
__typedArg0 := tmp18193
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18195 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(ifres18185, tmp18194)
}
__typedArg0 := ifres18185
__typedArg1 := tmp18194
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18196 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18177, tmp18195)
}
__typedArg0 := tmp18177
__typedArg1 := tmp18195
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp18196)
}
__typedArg0 := symif
__typedArg1 := tmp18196
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 5)

tmp18197 := Call(__e, ns2_1set, symshen_4variablecode, tmp18174)


_ = tmp18197

tmp18198 := MakeNative(func(__e *ControlFlow) {
V328 := __e.Get(1)
_ = V328
V329 := __e.Get(2)
_ = V329
V330 := __e.Get(3)
_ = V330
V331 := __e.Get(4)
_ = V331
V332 := __e.Get(5)
_ = V332
tmp18199 := Call(__e, PrimFunc(symgensym), symRemainder)


let__17494 := tmp18199
_ = let__17494

tmp18200 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V329, Nil)
}
__typedArg0 := V329
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18201 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons_2, tmp18200)
}
__typedArg0 := symcons_2
__typedArg1 := tmp18200
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18202 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V329, Nil)
}
__typedArg0 := V329
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18203 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtail, tmp18202)
}
__typedArg0 := symtail
__typedArg1 := tmp18202
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18204 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V328, let__17494, V331, V332)


tmp18205 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18204, Nil)
}
__typedArg0 := tmp18204
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18206 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18203, tmp18205)
}
__typedArg0 := tmp18203
__typedArg1 := tmp18205
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18207 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17494, tmp18206)
}
__typedArg0 := let__17494
__typedArg1 := tmp18206
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18208 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp18207)
}
__typedArg0 := symlet
__typedArg1 := tmp18207
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18209 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4parse_1failure, Nil)
}
__typedArg0 := symshen_4parse_1failure
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18210 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18209, Nil)
}
__typedArg0 := tmp18209
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18211 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18208, tmp18210)
}
__typedArg0 := tmp18208
__typedArg1 := tmp18210
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18212 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18201, tmp18211)
}
__typedArg0 := tmp18201
__typedArg1 := tmp18211
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp18212)
}
__typedArg0 := symif
__typedArg1 := tmp18212
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 5)

tmp18213 := Call(__e, ns2_1set, symshen_4wildcardcode, tmp18198)


_ = tmp18213

tmp18214 := MakeNative(func(__e *ControlFlow) {
V334 := __e.Get(1)
_ = V334
V335 := __e.Get(2)
_ = V335
V336 := __e.Get(3)
_ = V336
V337 := __e.Get(4)
_ = V337
V338 := __e.Get(5)
_ = V338
tmp18215 := Call(__e, PrimFunc(symgensym), symRemainder)


let__17495 := tmp18215
_ = let__17495

tmp18216 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V336, Nil)
}
__typedArg0 := V336
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18217 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V335, tmp18216)
}
__typedArg0 := V335
__typedArg1 := tmp18216
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18218 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4hds_a_2, tmp18217)
}
__typedArg0 := symshen_4hds_a_2
__typedArg1 := tmp18217
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18219 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V335, Nil)
}
__typedArg0 := V335
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18220 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtail, tmp18219)
}
__typedArg0 := symtail
__typedArg1 := tmp18219
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18221 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V334, let__17495, V337, V338)


tmp18222 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18221, Nil)
}
__typedArg0 := tmp18221
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18223 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18220, tmp18222)
}
__typedArg0 := tmp18220
__typedArg1 := tmp18222
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18224 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17495, tmp18223)
}
__typedArg0 := let__17495
__typedArg1 := tmp18223
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18225 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp18224)
}
__typedArg0 := symlet
__typedArg1 := tmp18224
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18226 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4parse_1failure, Nil)
}
__typedArg0 := symshen_4parse_1failure
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18227 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18226, Nil)
}
__typedArg0 := tmp18226
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18228 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18225, tmp18227)
}
__typedArg0 := tmp18225
__typedArg1 := tmp18227
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18229 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18218, tmp18228)
}
__typedArg0 := tmp18218
__typedArg1 := tmp18228
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp18229)
}
__typedArg0 := symif
__typedArg1 := tmp18229
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 5)

tmp18230 := Call(__e, ns2_1set, symshen_4terminalcode, tmp18214)


_ = tmp18230

tmp18231 := MakeNative(func(__e *ControlFlow) {
V347 := __e.Get(1)
_ = V347
V348 := __e.Get(2)
_ = V348
tmp18237 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V347)
}
__typedArg0 := V347
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18233 Obj

if True == tmp18237 {
tmp18235 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V347)
}
__typedArg0 := V347
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18236 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp18235, V348)
}
__typedArg0 := tmp18235
__typedArg1 := V348
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18234 Obj

if True == tmp18236 {
ifres18234 = True


} else {
ifres18234 = False


}

ifres18233 = ifres18234


} else {
ifres18233 = False


}

if True == ifres18233 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}, 2)

tmp18238 := Call(__e, ns2_1set, symshen_4hds_a_2, tmp18231)


_ = tmp18238

tmp18239 := MakeNative(func(__e *ControlFlow) {
V349 := __e.Get(1)
_ = V349
V350 := __e.Get(2)
_ = V350
V351 := __e.Get(3)
_ = V351
V352 := __e.Get(4)
_ = V352
V353 := __e.Get(5)
_ = V353
tmp18240 := Call(__e, PrimFunc(symgensym), symRemainder)


let__17496 := tmp18240
_ = let__17496

tmp18241 := Call(__e, PrimFunc(symgensym), symHd)


let__17497 := tmp18241
_ = let__17497

tmp18242 := Call(__e, PrimFunc(symgensym), symTl)


let__17498 := tmp18242
_ = let__17498

tmp18243 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V350, Nil)
}
__typedArg0 := V350
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18244 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4ccons_2, tmp18243)
}
__typedArg0 := symshen_4ccons_2
__typedArg1 := tmp18243
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18245 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V350, Nil)
}
__typedArg0 := V350
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18246 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhead, tmp18245)
}
__typedArg0 := symhead
__typedArg1 := tmp18245
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18247 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V350, Nil)
}
__typedArg0 := V350
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18248 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtail, tmp18247)
}
__typedArg0 := symtail
__typedArg1 := tmp18247
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18249 := Call(__e, PrimFunc(symshen_4decons), V351)


tmp18250 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5end_6, Nil)
}
__typedArg0 := sym_5end_6
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18251 := Call(__e, PrimFunc(symappend), tmp18249, tmp18250)


tmp18252 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V349, let__17498, V352, V353)


tmp18253 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18252, Nil)
}
__typedArg0 := tmp18252
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18254 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4processed, tmp18253)
}
__typedArg0 := symshen_4processed
__typedArg1 := tmp18253
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18255 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V349, let__17497, tmp18251, tmp18254)


tmp18256 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18255, Nil)
}
__typedArg0 := tmp18255
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18257 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18248, tmp18256)
}
__typedArg0 := tmp18248
__typedArg1 := tmp18256
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18258 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17498, tmp18257)
}
__typedArg0 := let__17498
__typedArg1 := tmp18257
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18259 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18246, tmp18258)
}
__typedArg0 := tmp18246
__typedArg1 := tmp18258
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18260 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17497, tmp18259)
}
__typedArg0 := let__17497
__typedArg1 := tmp18259
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18261 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp18260)
}
__typedArg0 := symlet
__typedArg1 := tmp18260
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18262 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4parse_1failure, Nil)
}
__typedArg0 := symshen_4parse_1failure
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18263 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18262, Nil)
}
__typedArg0 := tmp18262
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18264 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18261, tmp18263)
}
__typedArg0 := tmp18261
__typedArg1 := tmp18263
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18265 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18244, tmp18264)
}
__typedArg0 := tmp18244
__typedArg1 := tmp18264
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp18265)
}
__typedArg0 := symif
__typedArg1 := tmp18265
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 5)

tmp18266 := Call(__e, ns2_1set, symshen_4conscode, tmp18239)


_ = tmp18266

tmp18267 := MakeNative(func(__e *ControlFlow) {
V367 := __e.Get(1)
_ = V367
tmp18279 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V367)
}
__typedArg0 := V367
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18275 Obj

if True == tmp18279 {
tmp18277 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V367)
}
__typedArg0 := V367
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18278 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18277)
}
__typedArg0 := tmp18277
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18276 Obj

if True == tmp18278 {
ifres18276 = True


} else {
ifres18276 = False


}

ifres18275 = ifres18276


} else {
ifres18275 = False


}

if True == ifres18275 {
__e.Return(True)
return
} else {
tmp18273 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V367)
}
__typedArg0 := V367
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18269 Obj

if True == tmp18273 {
tmp18271 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V367)
}
__typedArg0 := V367
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18272 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp18271)
}
__typedArg0 := Nil
__typedArg1 := tmp18271
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18270 Obj

if True == tmp18272 {
ifres18270 = True


} else {
ifres18270 = False


}

ifres18269 = ifres18270


} else {
ifres18269 = False


}

if True == ifres18269 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp18280 := Call(__e, ns2_1set, symshen_4ccons_2, tmp18267)


_ = tmp18280

tmp18281 := MakeNative(func(__e *ControlFlow) {
V368 := __e.Get(1)
_ = V368
tmp18308 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V368)
}
__typedArg0 := V368
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18289 Obj

if True == tmp18308 {
tmp18306 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V368)
}
__typedArg0 := V368
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18307 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcons, tmp18306)
}
__typedArg0 := symcons
__typedArg1 := tmp18306
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18291 Obj

if True == tmp18307 {
tmp18304 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V368)
}
__typedArg0 := V368
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18305 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18304)
}
__typedArg0 := tmp18304
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18293 Obj

if True == tmp18305 {
tmp18301 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V368)
}
__typedArg0 := V368
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18302 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18301)
}
__typedArg0 := tmp18301
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18303 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18302)
}
__typedArg0 := tmp18302
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18295 Obj

if True == tmp18303 {
tmp18297 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V368)
}
__typedArg0 := V368
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18298 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18297)
}
__typedArg0 := tmp18297
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18299 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18298)
}
__typedArg0 := tmp18298
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18300 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp18299)
}
__typedArg0 := Nil
__typedArg1 := tmp18299
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18296 Obj

if True == tmp18300 {
ifres18296 = True


} else {
ifres18296 = False


}

ifres18295 = ifres18296


} else {
ifres18295 = False


}

var ifres18294 Obj

if True == ifres18295 {
ifres18294 = True


} else {
ifres18294 = False


}

ifres18293 = ifres18294


} else {
ifres18293 = False


}

var ifres18292 Obj

if True == ifres18293 {
ifres18292 = True


} else {
ifres18292 = False


}

ifres18291 = ifres18292


} else {
ifres18291 = False


}

var ifres18290 Obj

if True == ifres18291 {
ifres18290 = True


} else {
ifres18290 = False


}

ifres18289 = ifres18290


} else {
ifres18289 = False


}

if True == ifres18289 {
tmp18282 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V368)
}
__typedArg0 := V368
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18283 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18282)
}
__typedArg0 := tmp18282
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18284 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V368)
}
__typedArg0 := V368
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18285 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18284)
}
__typedArg0 := tmp18284
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18286 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18285)
}
__typedArg0 := tmp18285
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18287 := Call(__e, PrimFunc(symshen_4decons), tmp18286)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18283, tmp18287)
}
__typedArg0 := tmp18283
__typedArg1 := tmp18287
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return(V368)
return
}


}, 1)

tmp18309 := Call(__e, ns2_1set, symshen_4decons, tmp18281)


_ = tmp18309

tmp18310 := MakeNative(func(__e *ControlFlow) {
V369 := __e.Get(1)
_ = V369
V370 := __e.Get(2)
_ = V370
tmp18311 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V370, Nil)
}
__typedArg0 := V370
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V369, tmp18311)
}
__typedArg0 := V369
__typedArg1 := tmp18311
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 2)

tmp18312 := Call(__e, ns2_1set, symshen_4comb, tmp18310)


_ = tmp18312

tmp18313 := MakeNative(func(__e *ControlFlow) {
V375 := __e.Get(1)
_ = V375
V376 := __e.Get(2)
_ = V376
V377 := __e.Get(3)
_ = V377
tmp18333 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V377)
}
__typedArg0 := V377
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18320 Obj

if True == tmp18333 {
tmp18331 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V377)
}
__typedArg0 := V377
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18332 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4processed, tmp18331)
}
__typedArg0 := symshen_4processed
__typedArg1 := tmp18331
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18322 Obj

if True == tmp18332 {
tmp18329 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V377)
}
__typedArg0 := V377
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18330 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18329)
}
__typedArg0 := tmp18329
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18324 Obj

if True == tmp18330 {
tmp18326 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V377)
}
__typedArg0 := V377
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18327 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18326)
}
__typedArg0 := tmp18326
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18328 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp18327)
}
__typedArg0 := Nil
__typedArg1 := tmp18327
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18325 Obj

if True == tmp18328 {
ifres18325 = True


} else {
ifres18325 = False


}

ifres18324 = ifres18325


} else {
ifres18324 = False


}

var ifres18323 Obj

if True == ifres18324 {
ifres18323 = True


} else {
ifres18323 = False


}

ifres18322 = ifres18323


} else {
ifres18322 = False


}

var ifres18321 Obj

if True == ifres18322 {
ifres18321 = True


} else {
ifres18321 = False


}

ifres18320 = ifres18321


} else {
ifres18320 = False


}

if True == ifres18320 {
tmp18314 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V377)
}
__typedArg0 := V377
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18314)
}
__typedArg0 := tmp18314
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return


} else {
tmp18315 := Call(__e, PrimFunc(symshen_4process_1yacc_1semantics), V377)


let__17499 := tmp18315
_ = let__17499

tmp18316 := Call(__e, PrimFunc(symshen_4use_1type_1info), V375, let__17499)


let__17500 := tmp18316
_ = let__17500

tmp18317 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17500, Nil)
}
__typedArg0 := let__17500
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18318 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V376, tmp18317)
}
__typedArg0 := V376
__typedArg1 := tmp18317
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4comb, tmp18318)
}
__typedArg0 := symshen_4comb
__typedArg1 := tmp18318
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}


}, 3)

tmp18334 := Call(__e, ns2_1set, symshen_4yacc_1semantics, tmp18313)


_ = tmp18334

tmp18335 := MakeNative(func(__e *ControlFlow) {
V383 := __e.Get(1)
_ = V383
V384 := __e.Get(2)
_ = V384
tmp18523 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18344 Obj

if True == tmp18523 {
tmp18521 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18522 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_i, tmp18521)
}
__typedArg0 := sym_i
__typedArg1 := tmp18521
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18346 Obj

if True == tmp18522 {
tmp18519 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18520 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18519)
}
__typedArg0 := tmp18519
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18348 Obj

if True == tmp18520 {
tmp18516 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18517 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18516)
}
__typedArg0 := tmp18516
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18518 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18517)
}
__typedArg0 := tmp18517
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18350 Obj

if True == tmp18518 {
tmp18512 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18513 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18512)
}
__typedArg0 := tmp18512
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18514 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18513)
}
__typedArg0 := tmp18513
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18515 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlist, tmp18514)
}
__typedArg0 := symlist
__typedArg1 := tmp18514
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18352 Obj

if True == tmp18515 {
tmp18508 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18509 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18508)
}
__typedArg0 := tmp18508
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18510 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18509)
}
__typedArg0 := tmp18509
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18511 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18510)
}
__typedArg0 := tmp18510
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18354 Obj

if True == tmp18511 {
tmp18503 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18504 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18503)
}
__typedArg0 := tmp18503
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18505 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18504)
}
__typedArg0 := tmp18504
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18506 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18505)
}
__typedArg0 := tmp18505
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18507 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp18506)
}
__typedArg0 := Nil
__typedArg1 := tmp18506
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18356 Obj

if True == tmp18507 {
tmp18500 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18501 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18500)
}
__typedArg0 := tmp18500
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18502 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18501)
}
__typedArg0 := tmp18501
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18358 Obj

if True == tmp18502 {
tmp18496 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18497 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18496)
}
__typedArg0 := tmp18496
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18498 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18497)
}
__typedArg0 := tmp18497
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18499 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_1_1_6, tmp18498)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp18498
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18360 Obj

if True == tmp18499 {
tmp18492 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18493 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18492)
}
__typedArg0 := tmp18492
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18494 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18493)
}
__typedArg0 := tmp18493
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18495 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18494)
}
__typedArg0 := tmp18494
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18362 Obj

if True == tmp18495 {
tmp18487 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18488 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18487)
}
__typedArg0 := tmp18487
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18489 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18488)
}
__typedArg0 := tmp18488
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18490 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18489)
}
__typedArg0 := tmp18489
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18491 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18490)
}
__typedArg0 := tmp18490
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18364 Obj

if True == tmp18491 {
tmp18481 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18482 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18481)
}
__typedArg0 := tmp18481
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18483 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18482)
}
__typedArg0 := tmp18482
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18484 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18483)
}
__typedArg0 := tmp18483
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18485 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18484)
}
__typedArg0 := tmp18484
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18486 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symstr, tmp18485)
}
__typedArg0 := symstr
__typedArg1 := tmp18485
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18366 Obj

if True == tmp18486 {
tmp18475 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18476 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18475)
}
__typedArg0 := tmp18475
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18477 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18476)
}
__typedArg0 := tmp18476
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18478 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18477)
}
__typedArg0 := tmp18477
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18479 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18478)
}
__typedArg0 := tmp18478
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18480 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18479)
}
__typedArg0 := tmp18479
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18368 Obj

if True == tmp18480 {
tmp18468 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18469 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18468)
}
__typedArg0 := tmp18468
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18470 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18469)
}
__typedArg0 := tmp18469
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18471 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18470)
}
__typedArg0 := tmp18470
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18472 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18471)
}
__typedArg0 := tmp18471
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18473 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18472)
}
__typedArg0 := tmp18472
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18474 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18473)
}
__typedArg0 := tmp18473
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18370 Obj

if True == tmp18474 {
tmp18460 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18461 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18460)
}
__typedArg0 := tmp18460
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18462 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18461)
}
__typedArg0 := tmp18461
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18463 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18462)
}
__typedArg0 := tmp18462
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18464 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18463)
}
__typedArg0 := tmp18463
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18465 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18464)
}
__typedArg0 := tmp18464
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18466 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18465)
}
__typedArg0 := tmp18465
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18467 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlist, tmp18466)
}
__typedArg0 := symlist
__typedArg1 := tmp18466
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18372 Obj

if True == tmp18467 {
tmp18452 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18453 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18452)
}
__typedArg0 := tmp18452
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18454 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18453)
}
__typedArg0 := tmp18453
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18455 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18454)
}
__typedArg0 := tmp18454
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18456 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18455)
}
__typedArg0 := tmp18455
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18457 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18456)
}
__typedArg0 := tmp18456
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18458 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18457)
}
__typedArg0 := tmp18457
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18459 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18458)
}
__typedArg0 := tmp18458
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18374 Obj

if True == tmp18459 {
tmp18443 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18444 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18443)
}
__typedArg0 := tmp18443
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18445 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18444)
}
__typedArg0 := tmp18444
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18446 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18445)
}
__typedArg0 := tmp18445
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18447 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18446)
}
__typedArg0 := tmp18446
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18448 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18447)
}
__typedArg0 := tmp18447
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18449 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18448)
}
__typedArg0 := tmp18448
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18450 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18449)
}
__typedArg0 := tmp18449
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18451 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp18450)
}
__typedArg0 := Nil
__typedArg1 := tmp18450
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18376 Obj

if True == tmp18451 {
tmp18436 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18437 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18436)
}
__typedArg0 := tmp18436
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18438 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18437)
}
__typedArg0 := tmp18437
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18439 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18438)
}
__typedArg0 := tmp18438
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18440 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18439)
}
__typedArg0 := tmp18439
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18441 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18440)
}
__typedArg0 := tmp18440
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18442 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18441)
}
__typedArg0 := tmp18441
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18378 Obj

if True == tmp18442 {
tmp18428 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18429 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18428)
}
__typedArg0 := tmp18428
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18430 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18429)
}
__typedArg0 := tmp18429
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18431 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18430)
}
__typedArg0 := tmp18430
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18432 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18431)
}
__typedArg0 := tmp18431
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18433 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18432)
}
__typedArg0 := tmp18432
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18434 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18433)
}
__typedArg0 := tmp18433
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18435 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp18434)
}
__typedArg0 := Nil
__typedArg1 := tmp18434
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18380 Obj

if True == tmp18435 {
tmp18423 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18424 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18423)
}
__typedArg0 := tmp18423
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18425 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18424)
}
__typedArg0 := tmp18424
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18426 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18425)
}
__typedArg0 := tmp18425
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18427 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18426)
}
__typedArg0 := tmp18426
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18382 Obj

if True == tmp18427 {
tmp18417 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18418 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18417)
}
__typedArg0 := tmp18417
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18419 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18418)
}
__typedArg0 := tmp18418
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18420 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18419)
}
__typedArg0 := tmp18419
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18421 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18420)
}
__typedArg0 := tmp18420
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18422 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_j, tmp18421)
}
__typedArg0 := sym_j
__typedArg1 := tmp18421
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18384 Obj

if True == tmp18422 {
tmp18411 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18412 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18411)
}
__typedArg0 := tmp18411
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18413 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18412)
}
__typedArg0 := tmp18412
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18414 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18413)
}
__typedArg0 := tmp18413
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18415 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18414)
}
__typedArg0 := tmp18414
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18416 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp18415)
}
__typedArg0 := Nil
__typedArg1 := tmp18415
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18386 Obj

if True == tmp18416 {
tmp18398 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18399 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18398)
}
__typedArg0 := tmp18398
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18400 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18399)
}
__typedArg0 := tmp18399
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18401 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18400)
}
__typedArg0 := tmp18400
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18402 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18403 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18402)
}
__typedArg0 := tmp18402
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18404 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18403)
}
__typedArg0 := tmp18403
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18405 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18404)
}
__typedArg0 := tmp18404
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18406 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18405)
}
__typedArg0 := tmp18405
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18407 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18406)
}
__typedArg0 := tmp18406
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18408 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18407)
}
__typedArg0 := tmp18407
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18409 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18408)
}
__typedArg0 := tmp18408
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18410 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp18401, tmp18409)
}
__typedArg0 := tmp18401
__typedArg1 := tmp18409
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18388 Obj

if True == tmp18410 {
tmp18390 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18391 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18390)
}
__typedArg0 := tmp18390
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18392 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18391)
}
__typedArg0 := tmp18391
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18393 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18392)
}
__typedArg0 := tmp18392
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18394 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18393)
}
__typedArg0 := tmp18393
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18395 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18394)
}
__typedArg0 := tmp18394
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18396 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18395)
}
__typedArg0 := tmp18395
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18397 := Call(__e, PrimFunc(symshen_4monomorphic_2), tmp18396)


var ifres18389 Obj

if True == tmp18397 {
ifres18389 = True


} else {
ifres18389 = False


}

ifres18388 = ifres18389


} else {
ifres18388 = False


}

var ifres18387 Obj

if True == ifres18388 {
ifres18387 = True


} else {
ifres18387 = False


}

ifres18386 = ifres18387


} else {
ifres18386 = False


}

var ifres18385 Obj

if True == ifres18386 {
ifres18385 = True


} else {
ifres18385 = False


}

ifres18384 = ifres18385


} else {
ifres18384 = False


}

var ifres18383 Obj

if True == ifres18384 {
ifres18383 = True


} else {
ifres18383 = False


}

ifres18382 = ifres18383


} else {
ifres18382 = False


}

var ifres18381 Obj

if True == ifres18382 {
ifres18381 = True


} else {
ifres18381 = False


}

ifres18380 = ifres18381


} else {
ifres18380 = False


}

var ifres18379 Obj

if True == ifres18380 {
ifres18379 = True


} else {
ifres18379 = False


}

ifres18378 = ifres18379


} else {
ifres18378 = False


}

var ifres18377 Obj

if True == ifres18378 {
ifres18377 = True


} else {
ifres18377 = False


}

ifres18376 = ifres18377


} else {
ifres18376 = False


}

var ifres18375 Obj

if True == ifres18376 {
ifres18375 = True


} else {
ifres18375 = False


}

ifres18374 = ifres18375


} else {
ifres18374 = False


}

var ifres18373 Obj

if True == ifres18374 {
ifres18373 = True


} else {
ifres18373 = False


}

ifres18372 = ifres18373


} else {
ifres18372 = False


}

var ifres18371 Obj

if True == ifres18372 {
ifres18371 = True


} else {
ifres18371 = False


}

ifres18370 = ifres18371


} else {
ifres18370 = False


}

var ifres18369 Obj

if True == ifres18370 {
ifres18369 = True


} else {
ifres18369 = False


}

ifres18368 = ifres18369


} else {
ifres18368 = False


}

var ifres18367 Obj

if True == ifres18368 {
ifres18367 = True


} else {
ifres18367 = False


}

ifres18366 = ifres18367


} else {
ifres18366 = False


}

var ifres18365 Obj

if True == ifres18366 {
ifres18365 = True


} else {
ifres18365 = False


}

ifres18364 = ifres18365


} else {
ifres18364 = False


}

var ifres18363 Obj

if True == ifres18364 {
ifres18363 = True


} else {
ifres18363 = False


}

ifres18362 = ifres18363


} else {
ifres18362 = False


}

var ifres18361 Obj

if True == ifres18362 {
ifres18361 = True


} else {
ifres18361 = False


}

ifres18360 = ifres18361


} else {
ifres18360 = False


}

var ifres18359 Obj

if True == ifres18360 {
ifres18359 = True


} else {
ifres18359 = False


}

ifres18358 = ifres18359


} else {
ifres18358 = False


}

var ifres18357 Obj

if True == ifres18358 {
ifres18357 = True


} else {
ifres18357 = False


}

ifres18356 = ifres18357


} else {
ifres18356 = False


}

var ifres18355 Obj

if True == ifres18356 {
ifres18355 = True


} else {
ifres18355 = False


}

ifres18354 = ifres18355


} else {
ifres18354 = False


}

var ifres18353 Obj

if True == ifres18354 {
ifres18353 = True


} else {
ifres18353 = False


}

ifres18352 = ifres18353


} else {
ifres18352 = False


}

var ifres18351 Obj

if True == ifres18352 {
ifres18351 = True


} else {
ifres18351 = False


}

ifres18350 = ifres18351


} else {
ifres18350 = False


}

var ifres18349 Obj

if True == ifres18350 {
ifres18349 = True


} else {
ifres18349 = False


}

ifres18348 = ifres18349


} else {
ifres18348 = False


}

var ifres18347 Obj

if True == ifres18348 {
ifres18347 = True


} else {
ifres18347 = False


}

ifres18346 = ifres18347


} else {
ifres18346 = False


}

var ifres18345 Obj

if True == ifres18346 {
ifres18345 = True


} else {
ifres18345 = False


}

ifres18344 = ifres18345


} else {
ifres18344 = False


}

if True == ifres18344 {
tmp18336 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18337 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18336)
}
__typedArg0 := tmp18336
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18338 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18337)
}
__typedArg0 := tmp18337
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18339 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18338)
}
__typedArg0 := tmp18338
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18340 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18339)
}
__typedArg0 := tmp18339
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18341 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18340)
}
__typedArg0 := tmp18340
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18342 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V384, tmp18341)
}
__typedArg0 := V384
__typedArg1 := tmp18341
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtype, tmp18342)
}
__typedArg0 := symtype
__typedArg1 := tmp18342
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return(V384)
return
}


}, 2)

tmp18524 := Call(__e, ns2_1set, symshen_4use_1type_1info, tmp18335)


_ = tmp18524

tmp18525 := MakeNative(func(__e *ControlFlow) {
V387 := __e.Get(1)
_ = V387
tmp18535 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(V387)
}
__typedArg0 := V387
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

if True == tmp18535 {
__e.Return(False)
return
} else {
tmp18533 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V387)
}
__typedArg0 := V387
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp18533 {
tmp18530 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V387)
}
__typedArg0 := V387
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18531 := Call(__e, PrimFunc(symshen_4monomorphic_2), tmp18530)


if True == tmp18531 {
tmp18527 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V387)
}
__typedArg0 := V387
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18528 := Call(__e, PrimFunc(symshen_4monomorphic_2), tmp18527)


if True == tmp18528 {
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


} else {
__e.Return(True)
return
}


}


}, 1)

tmp18536 := Call(__e, ns2_1set, symshen_4monomorphic_2, tmp18525)


_ = tmp18536

tmp18537 := MakeNative(func(__e *ControlFlow) {
V388 := __e.Get(1)
_ = V388
tmp18563 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V388)
}
__typedArg0 := V388
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18545 Obj

if True == tmp18563 {
tmp18561 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V388)
}
__typedArg0 := V388
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18562 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symprotect, tmp18561)
}
__typedArg0 := symprotect
__typedArg1 := tmp18561
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18547 Obj

if True == tmp18562 {
tmp18559 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V388)
}
__typedArg0 := V388
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18560 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18559)
}
__typedArg0 := tmp18559
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18549 Obj

if True == tmp18560 {
tmp18556 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V388)
}
__typedArg0 := V388
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18557 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18556)
}
__typedArg0 := tmp18556
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18558 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp18557)
}
__typedArg0 := Nil
__typedArg1 := tmp18557
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18551 Obj

if True == tmp18558 {
tmp18553 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V388)
}
__typedArg0 := V388
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18554 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18553)
}
__typedArg0 := tmp18553
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18555 := Call(__e, PrimFunc(symshen_4non_1terminal_2), tmp18554)


var ifres18552 Obj

if True == tmp18555 {
ifres18552 = True


} else {
ifres18552 = False


}

ifres18551 = ifres18552


} else {
ifres18551 = False


}

var ifres18550 Obj

if True == ifres18551 {
ifres18550 = True


} else {
ifres18550 = False


}

ifres18549 = ifres18550


} else {
ifres18549 = False


}

var ifres18548 Obj

if True == ifres18549 {
ifres18548 = True


} else {
ifres18548 = False


}

ifres18547 = ifres18548


} else {
ifres18547 = False


}

var ifres18546 Obj

if True == ifres18547 {
ifres18546 = True


} else {
ifres18546 = False


}

ifres18545 = ifres18546


} else {
ifres18545 = False


}

if True == ifres18545 {
tmp18538 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V388)
}
__typedArg0 := V388
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18538)
}
__typedArg0 := tmp18538
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return


} else {
tmp18543 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V388)
}
__typedArg0 := V388
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp18543 {
tmp18539 := MakeNative(func(__e *ControlFlow) {
Z389 := __e.Get(1)
_ = Z389
__e.TailApply(PrimFunc(symshen_4process_1yacc_1semantics), Z389)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp18539, V388)
return


} else {
tmp18541 := Call(__e, PrimFunc(symshen_4non_1terminal_2), V388)


if True == tmp18541 {
__e.TailApply(PrimFunc(symconcat), symAction, V388)
return
} else {
__e.Return(V388)
return
}


}


}


}, 1)

tmp18564 := Call(__e, ns2_1set, symshen_4process_1yacc_1semantics, tmp18537)


_ = tmp18564

tmp18565 := MakeNative(func(__e *ControlFlow) {
V392 := __e.Get(1)
_ = V392
tmp18578 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V392)
}
__typedArg0 := V392
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18569 Obj

if True == tmp18578 {
tmp18576 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V392)
}
__typedArg0 := V392
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18577 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18576)
}
__typedArg0 := tmp18576
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18571 Obj

if True == tmp18577 {
tmp18573 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V392)
}
__typedArg0 := V392
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18574 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18573)
}
__typedArg0 := tmp18573
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18575 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp18574)
}
__typedArg0 := Nil
__typedArg1 := tmp18574
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18572 Obj

if True == tmp18575 {
ifres18572 = True


} else {
ifres18572 = False


}

ifres18571 = ifres18572


} else {
ifres18571 = False


}

var ifres18570 Obj

if True == ifres18571 {
ifres18570 = True


} else {
ifres18570 = False


}

ifres18569 = ifres18570


} else {
ifres18569 = False


}

if True == ifres18569 {
tmp18566 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V392)
}
__typedArg0 := V392
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18566)
}
__typedArg0 := tmp18566
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return


} else {
tmp18567 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V392)
}
__typedArg0 := V392
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18567)
}
__typedArg0 := tmp18567
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return


}


}, 1)

tmp18579 := Call(__e, ns2_1set, symshen_4_5_1out, tmp18565)


_ = tmp18579

tmp18580 := MakeNative(func(__e *ControlFlow) {
V393 := __e.Get(1)
_ = V393
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V393)
}
__typedArg0 := V393
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return
}, 1)

tmp18581 := Call(__e, ns2_1set, symshen_4in_1_6, tmp18580)


_ = tmp18581

tmp18582 := MakeNative(func(__e *ControlFlow) {
V394 := __e.Get(1)
_ = V394
tmp18583 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V394, Nil)
}
__typedArg0 := V394
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Nil, tmp18583)
}
__typedArg0 := Nil
__typedArg1 := tmp18583
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp18584 := Call(__e, ns2_1set, sym_5_b_6, tmp18582)


_ = tmp18584

tmp18585 := MakeNative(func(__e *ControlFlow) {
V395 := __e.Get(1)
_ = V395
tmp18586 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Nil, Nil)
}
__typedArg0 := Nil
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V395, tmp18586)
}
__typedArg0 := V395
__typedArg1 := tmp18586
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp18587 := Call(__e, ns2_1set, sym_5e_6, tmp18585)


_ = tmp18587

tmp18588 := MakeNative(func(__e *ControlFlow) {
V398 := __e.Get(1)
_ = V398
tmp18591 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V398)
}
__typedArg0 := Nil
__typedArg1 := V398
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp18591 {
tmp18589 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Nil, Nil)
}
__typedArg0 := Nil
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Nil, tmp18589)
}
__typedArg0 := Nil
__typedArg1 := tmp18589
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

__e.TailApply(ns2_1set, sym_5end_6, tmp18588)
return




}, 0)

