package main

import . "github.com/pyrex41/shen-go/kl"

var YaccMain = MakeNative(func(__e *ControlFlow) {
tmp17508 := MakeNative(func(__e *ControlFlow) {
V112 := __e.Get(1)
_ = V112
V113 := __e.Get(2)
_ = V113
tmp17509 := Call(__e, V112, V113)


let__17370 := tmp17509
_ = let__17370

tmp17516 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17370)


if True == tmp17516 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("parse failure\n"))
}
__typedArg0 := MakeString("parse failure\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
} else {
tmp17514 := Call(__e, PrimFunc(symshen_4partial_1parse_1failure_2), let__17370)


if True == tmp17514 {
tmp17510 := Call(__e, PrimFunc(symshen_4in_1_6), let__17370)


tmp17511 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dresidue_d, tmp17510)
}
__typedArg0 := symshen_4_dresidue_d
__typedArg1 := tmp17510
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp17511

tmp17512 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dresidue_d)
}
__typedArg0 := symshen_4_dresidue_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4raise_1syntax_1error), tmp17512)
return


} else {
__e.TailApply(PrimFunc(symshen_4_5_1out), let__17370)
return
}


}


}, 2)

tmp17517 := Call(__e, ns2_1set, symcompile, tmp17508)


_ = tmp17517

tmp17518 := MakeNative(func(__e *ControlFlow) {
V115 := __e.Get(1)
_ = V115
tmp17519 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dmaximum_1print_1sequence_1size_d)
}
__typedArg0 := sym_dmaximum_1print_1sequence_1size_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp17520 := Call(__e, PrimFunc(symshen_4syntax_1error_1message), tmp17519, MakeInteger(0), V115)


tmp17522 := Call(__e, PrimFunc(symshen_4proc_1nl), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("syntax error here: "))
__typedS1, __typedOK1 := TypedString(tmp17520)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("syntax error here: ")
__typedArg1 := tmp17520
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(tmp17522)
}
__typedArg0 := tmp17522
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}, 1)

tmp17523 := Call(__e, ns2_1set, symshen_4raise_1syntax_1error, tmp17518)


_ = tmp17523

tmp17524 := MakeNative(func(__e *ControlFlow) {
V123 := __e.Get(1)
_ = V123
V124 := __e.Get(2)
_ = V124
V125 := __e.Get(3)
_ = V125
tmp17535 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V125)
}
__typedArg0 := Nil
__typedArg1 := V125
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17535 {
__e.Return(MakeString("\n"))
return
} else {
tmp17533 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V123, V124)
}
__typedArg0 := V123
__typedArg1 := V124
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17533 {
__e.Return(MakeString("...etc \n"))
return
} else {
tmp17531 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V125)
}
__typedArg0 := V125
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17531 {
tmp17525 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V125)
}
__typedArg0 := V125
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17526 := Call(__e, PrimFunc(symshen_4app), tmp17525, MakeString(" "), symshen_4s)


tmp17527 := (func() Obj {
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

tmp17528 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V125)
}
__typedArg0 := V125
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17529 := Call(__e, PrimFunc(symshen_4syntax_1error_1message), V123, tmp17527, tmp17528)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(tmp17526)
__typedS1, __typedOK1 := TypedString(tmp17529)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := tmp17526
__typedArg1 := tmp17529
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

tmp17536 := Call(__e, ns2_1set, symshen_4syntax_1error_1message, tmp17524)


_ = tmp17536

tmp17537 := MakeNative(func(__e *ControlFlow) {
V126 := __e.Get(1)
_ = V126
tmp17538 := Call(__e, PrimFunc(symfail))


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V126, tmp17538)
}
__typedArg0 := V126
__typedArg1 := tmp17538
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp17539 := Call(__e, ns2_1set, symshen_4parse_1failure_2, tmp17537)


_ = tmp17539

tmp17540 := MakeNative(func(__e *ControlFlow) {
V127 := __e.Get(1)
_ = V127
tmp17541 := Call(__e, PrimFunc(symshen_4in_1_6), V127)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp17541)
}
__typedArg0 := tmp17541
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})())
return


}, 1)

tmp17542 := Call(__e, ns2_1set, symshen_4partial_1parse_1failure_2, tmp17540)


_ = tmp17542

tmp17543 := MakeNative(func(__e *ControlFlow) {
V130 := __e.Get(1)
_ = V130
tmp17556 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V130)
}
__typedArg0 := V130
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17547 Obj

if True == tmp17556 {
tmp17554 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V130)
}
__typedArg0 := V130
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17555 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp17554)
}
__typedArg0 := tmp17554
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17549 Obj

if True == tmp17555 {
tmp17551 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V130)
}
__typedArg0 := V130
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17552 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp17551)
}
__typedArg0 := tmp17551
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17553 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp17552)
}
__typedArg0 := Nil
__typedArg1 := tmp17552
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17550 Obj

if True == tmp17553 {
ifres17550 = True


} else {
ifres17550 = False


}

ifres17549 = ifres17550


} else {
ifres17549 = False


}

var ifres17548 Obj

if True == ifres17549 {
ifres17548 = True


} else {
ifres17548 = False


}

ifres17547 = ifres17548


} else {
ifres17547 = False


}

if True == ifres17547 {
tmp17544 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V130)
}
__typedArg0 := V130
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp17544)
}
__typedArg0 := tmp17544
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return


} else {
tmp17545 := Call(__e, PrimFunc(symshen_4app), V130, MakeString(" is not a YACC stream\n"), symshen_4s)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(tmp17545)
}
__typedArg0 := tmp17545
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}


}, 1)

tmp17557 := Call(__e, ns2_1set, symshen_4objectcode, tmp17543)


_ = tmp17557

tmp17558 := MakeNative(func(__e *ControlFlow) {
V131 := __e.Get(1)
_ = V131
tmp17559 := MakeNative(func(__e *ControlFlow) {
Z132 := __e.Get(1)
_ = Z132
__e.TailApply(PrimFunc(symshen_4_5yacc_6), Z132)
return
}, 1)

__e.TailApply(PrimFunc(symcompile), tmp17559, V131)
return


}, 1)

tmp17560 := Call(__e, ns2_1set, symshen_4yacc_1_6shen, tmp17558)


_ = tmp17560

tmp17561 := MakeNative(func(__e *ControlFlow) {
V133 := __e.Get(1)
_ = V133
tmp17588 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V133)
}
__typedArg0 := V133
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17562 Obj

if True == tmp17588 {
tmp17563 := Call(__e, PrimFunc(symhead), V133)


let__17372 := tmp17563
_ = let__17372

tmp17564 := Call(__e, PrimFunc(symtail), V133)


let__17373 := tmp17564
_ = let__17373

tmp17565 := Call(__e, PrimFunc(symshen_4_5yaccsig_6), let__17373)


let__17374 := tmp17565
_ = let__17374

tmp17586 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17374)


var ifres17566 Obj

if True == tmp17586 {
tmp17567 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17566 = tmp17567


} else {
tmp17568 := Call(__e, PrimFunc(symshen_4_5_1out), let__17374)


let__17375 := tmp17568
_ = let__17375

tmp17569 := Call(__e, PrimFunc(symshen_4in_1_6), let__17374)


let__17376 := tmp17569
_ = let__17376

tmp17570 := Call(__e, PrimFunc(symshen_4_5c_1rules_6), let__17376)


let__17377 := tmp17570
_ = let__17377

tmp17585 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17377)


var ifres17571 Obj

if True == tmp17585 {
tmp17572 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17571 = tmp17572


} else {
tmp17573 := Call(__e, PrimFunc(symshen_4_5_1out), let__17377)


let__17378 := tmp17573
_ = let__17378

tmp17574 := Call(__e, PrimFunc(symshen_4in_1_6), let__17377)


let__17379 := tmp17574
_ = let__17379

tmp17575 := Call(__e, PrimFunc(symgensym), symS)


let__17380 := tmp17575
_ = let__17380

tmp17576 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17372, Nil)
}
__typedArg0 := let__17372
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17577 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefine, tmp17576)
}
__typedArg0 := symdefine
__typedArg1 := tmp17576
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17578 := Call(__e, PrimFunc(symshen_4c_1rules_1_6shen), let__17375, let__17380, let__17378)


tmp17579 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17578, Nil)
}
__typedArg0 := tmp17578
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17580 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_6, tmp17579)
}
__typedArg0 := sym_1_6
__typedArg1 := tmp17579
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17581 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17380, tmp17580)
}
__typedArg0 := let__17380
__typedArg1 := tmp17580
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17582 := Call(__e, PrimFunc(symappend), let__17375, tmp17581)


tmp17583 := Call(__e, PrimFunc(symappend), tmp17577, tmp17582)


let__17381 := tmp17583
_ = let__17381

tmp17584 := Call(__e, PrimFunc(symshen_4comb), let__17379, let__17381)


ifres17571 = tmp17584


}

ifres17566 = ifres17571


}

ifres17562 = ifres17566


} else {
tmp17587 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17562 = tmp17587


}

let__17371 := ifres17562
_ = let__17371

tmp17590 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17371)


if True == tmp17590 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__17371)
return
}


}, 1)

tmp17591 := Call(__e, ns2_1set, symshen_4_5yacc_6, tmp17561)


_ = tmp17591

tmp17592 := MakeNative(func(__e *ControlFlow) {
V145 := __e.Get(1)
_ = V145
tmp17649 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V145)
}
__typedArg0 := V145
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17593 Obj

if True == tmp17649 {
tmp17594 := Call(__e, PrimFunc(symhead), V145)


let__17383 := tmp17594
_ = let__17383

tmp17595 := Call(__e, PrimFunc(symtail), V145)


let__17384 := tmp17595
_ = let__17384

tmp17647 := Call(__e, PrimFunc(symshen_4ccons_2), let__17384)


var ifres17596 Obj

if True == tmp17647 {
tmp17597 := Call(__e, PrimFunc(symhead), let__17384)


let__17385 := tmp17597
_ = let__17385

tmp17598 := Call(__e, PrimFunc(symtail), let__17384)


let__17386 := tmp17598
_ = let__17386

tmp17645 := Call(__e, PrimFunc(symshen_4hds_a_2), let__17385, symlist)


var ifres17599 Obj

if True == tmp17645 {
tmp17600 := Call(__e, PrimFunc(symtail), let__17385)


let__17387 := tmp17600
_ = let__17387

tmp17643 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__17387)
}
__typedArg0 := let__17387
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17601 Obj

if True == tmp17643 {
tmp17602 := Call(__e, PrimFunc(symhead), let__17387)


let__17388 := tmp17602
_ = let__17388

tmp17603 := Call(__e, PrimFunc(symtail), let__17387)


let__17389 := tmp17603
_ = let__17389

tmp17604 := Call(__e, PrimFunc(sym_5end_6), let__17389)


let__17390 := tmp17604
_ = let__17390

tmp17641 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17390)


var ifres17605 Obj

if True == tmp17641 {
tmp17606 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17605 = tmp17606


} else {
tmp17607 := Call(__e, PrimFunc(symshen_4in_1_6), let__17390)


let__17391 := tmp17607
_ = let__17391

tmp17640 := Call(__e, PrimFunc(symshen_4hds_a_2), let__17386, sym_a_a_6)


var ifres17608 Obj

if True == tmp17640 {
tmp17609 := Call(__e, PrimFunc(symtail), let__17386)


let__17392 := tmp17609
_ = let__17392

tmp17638 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__17392)
}
__typedArg0 := let__17392
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17610 Obj

if True == tmp17638 {
tmp17611 := Call(__e, PrimFunc(symhead), let__17392)


let__17393 := tmp17611
_ = let__17393

tmp17612 := Call(__e, PrimFunc(symtail), let__17392)


let__17394 := tmp17612
_ = let__17394

tmp17636 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__17394)
}
__typedArg0 := let__17394
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17613 Obj

if True == tmp17636 {
tmp17614 := Call(__e, PrimFunc(symhead), let__17394)


let__17395 := tmp17614
_ = let__17395

tmp17615 := Call(__e, PrimFunc(symtail), let__17394)


let__17396 := tmp17615
_ = let__17396

tmp17634 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_i, let__17383)
}
__typedArg0 := sym_i
__typedArg1 := let__17383
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17631 Obj

if True == tmp17634 {
tmp17633 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_j, let__17395)
}
__typedArg0 := sym_j
__typedArg1 := let__17395
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17632 Obj

if True == tmp17633 {
ifres17632 = True


} else {
ifres17632 = False


}

ifres17631 = ifres17632


} else {
ifres17631 = False


}

var ifres17616 Obj

if True == ifres17631 {
tmp17617 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17388, Nil)
}
__typedArg0 := let__17388
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17618 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp17617)
}
__typedArg0 := symlist
__typedArg1 := tmp17617
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17619 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17388, Nil)
}
__typedArg0 := let__17388
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17620 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp17619)
}
__typedArg0 := symlist
__typedArg1 := tmp17619
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17621 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17393, Nil)
}
__typedArg0 := let__17393
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17622 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17620, tmp17621)
}
__typedArg0 := tmp17620
__typedArg1 := tmp17621
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17623 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstr, tmp17622)
}
__typedArg0 := symstr
__typedArg1 := tmp17622
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17624 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_j, Nil)
}
__typedArg0 := sym_j
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17625 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17623, tmp17624)
}
__typedArg0 := tmp17623
__typedArg1 := tmp17624
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17626 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp17625)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp17625
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17627 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17618, tmp17626)
}
__typedArg0 := tmp17618
__typedArg1 := tmp17626
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17628 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_i, tmp17627)
}
__typedArg0 := sym_i
__typedArg1 := tmp17627
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17629 := Call(__e, PrimFunc(symshen_4comb), let__17396, tmp17628)


ifres17616 = tmp17629


} else {
tmp17630 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17616 = tmp17630


}

ifres17613 = ifres17616


} else {
tmp17635 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17613 = tmp17635


}

ifres17610 = ifres17613


} else {
tmp17637 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17610 = tmp17637


}

ifres17608 = ifres17610


} else {
tmp17639 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17608 = tmp17639


}

ifres17605 = ifres17608


}

ifres17601 = ifres17605


} else {
tmp17642 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17601 = tmp17642


}

ifres17599 = ifres17601


} else {
tmp17644 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17599 = tmp17644


}

ifres17596 = ifres17599


} else {
tmp17646 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17596 = tmp17646


}

ifres17593 = ifres17596


} else {
tmp17648 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17593 = tmp17648


}

let__17382 := ifres17593
_ = let__17382

tmp17659 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17382)


if True == tmp17659 {
tmp17650 := Call(__e, PrimFunc(sym_5e_6), V145)


let__17398 := tmp17650
_ = let__17398

tmp17655 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17398)


var ifres17651 Obj

if True == tmp17655 {
tmp17652 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17651 = tmp17652


} else {
tmp17653 := Call(__e, PrimFunc(symshen_4in_1_6), let__17398)


let__17399 := tmp17653
_ = let__17399

tmp17654 := Call(__e, PrimFunc(symshen_4comb), let__17399, Nil)


ifres17651 = tmp17654


}

let__17397 := ifres17651
_ = let__17397

tmp17657 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17397)


if True == tmp17657 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__17397)
return
}


} else {
__e.Return(let__17382)
return
}


}, 1)

tmp17660 := Call(__e, ns2_1set, symshen_4_5yaccsig_6, tmp17592)


_ = tmp17660

tmp17661 := MakeNative(func(__e *ControlFlow) {
V164 := __e.Get(1)
_ = V164
tmp17662 := Call(__e, PrimFunc(symshen_4_5c_1rule_6), V164)


let__17401 := tmp17662
_ = let__17401

tmp17675 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17401)


var ifres17663 Obj

if True == tmp17675 {
tmp17664 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17663 = tmp17664


} else {
tmp17665 := Call(__e, PrimFunc(symshen_4_5_1out), let__17401)


let__17402 := tmp17665
_ = let__17402

tmp17666 := Call(__e, PrimFunc(symshen_4in_1_6), let__17401)


let__17403 := tmp17666
_ = let__17403

tmp17667 := Call(__e, PrimFunc(symshen_4_5c_1rules_6), let__17403)


let__17404 := tmp17667
_ = let__17404

tmp17674 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17404)


var ifres17668 Obj

if True == tmp17674 {
tmp17669 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17668 = tmp17669


} else {
tmp17670 := Call(__e, PrimFunc(symshen_4_5_1out), let__17404)


let__17405 := tmp17670
_ = let__17405

tmp17671 := Call(__e, PrimFunc(symshen_4in_1_6), let__17404)


let__17406 := tmp17671
_ = let__17406

tmp17672 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17402, let__17405)
}
__typedArg0 := let__17402
__typedArg1 := let__17405
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17673 := Call(__e, PrimFunc(symshen_4comb), let__17406, tmp17672)


ifres17668 = tmp17673


}

ifres17663 = ifres17668


}

let__17400 := ifres17663
_ = let__17400

tmp17691 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17400)


if True == tmp17691 {
tmp17676 := Call(__e, PrimFunc(sym_5_b_6), V164)


let__17408 := tmp17676
_ = let__17408

tmp17687 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17408)


var ifres17677 Obj

if True == tmp17687 {
tmp17678 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17677 = tmp17678


} else {
tmp17679 := Call(__e, PrimFunc(symshen_4_5_1out), let__17408)


let__17409 := tmp17679
_ = let__17409

tmp17680 := Call(__e, PrimFunc(symshen_4in_1_6), let__17408)


let__17410 := tmp17680
_ = let__17410

tmp17685 := Call(__e, PrimFunc(symempty_2), let__17409)


var ifres17681 Obj

if True == tmp17685 {
ifres17681 = Nil


} else {
tmp17682 := Call(__e, PrimFunc(symshen_4app), let__17409, MakeString("\n ..."), symshen_4r)


tmp17684 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("YACC syntax error here:\n "))
__typedS1, __typedOK1 := TypedString(tmp17682)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("YACC syntax error here:\n ")
__typedArg1 := tmp17682
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("YACC syntax error here:\n "))
__typedS1, __typedOK1 := TypedString(tmp17682)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("YACC syntax error here:\n ")
__typedArg1 := tmp17682
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})()

ifres17681 = tmp17684


}

tmp17686 := Call(__e, PrimFunc(symshen_4comb), let__17410, ifres17681)


ifres17677 = tmp17686


}

let__17407 := ifres17677
_ = let__17407

tmp17689 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17407)


if True == tmp17689 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__17407)
return
}


} else {
__e.Return(let__17400)
return
}


}, 1)

tmp17692 := Call(__e, ns2_1set, symshen_4_5c_1rules_6, tmp17661)


_ = tmp17692

tmp17693 := MakeNative(func(__e *ControlFlow) {
V176 := __e.Get(1)
_ = V176
tmp17694 := Call(__e, PrimFunc(symshen_4_5syntax_6), V176)


let__17412 := tmp17694
_ = let__17412

tmp17713 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17412)


var ifres17695 Obj

if True == tmp17713 {
tmp17696 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17695 = tmp17696


} else {
tmp17697 := Call(__e, PrimFunc(symshen_4_5_1out), let__17412)


let__17413 := tmp17697
_ = let__17413

tmp17698 := Call(__e, PrimFunc(symshen_4in_1_6), let__17412)


let__17414 := tmp17698
_ = let__17414

tmp17699 := Call(__e, PrimFunc(symshen_4_5semantics_6), let__17414)


let__17415 := tmp17699
_ = let__17415

tmp17712 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17415)


var ifres17700 Obj

if True == tmp17712 {
tmp17701 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17700 = tmp17701


} else {
tmp17702 := Call(__e, PrimFunc(symshen_4_5_1out), let__17415)


let__17416 := tmp17702
_ = let__17416

tmp17703 := Call(__e, PrimFunc(symshen_4in_1_6), let__17415)


let__17417 := tmp17703
_ = let__17417

tmp17704 := Call(__e, PrimFunc(symshen_4_5sc_6), let__17417)


let__17418 := tmp17704
_ = let__17418

tmp17711 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17418)


var ifres17705 Obj

if True == tmp17711 {
tmp17706 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17705 = tmp17706


} else {
tmp17707 := Call(__e, PrimFunc(symshen_4in_1_6), let__17418)


let__17419 := tmp17707
_ = let__17419

tmp17708 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17416, Nil)
}
__typedArg0 := let__17416
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17709 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17413, tmp17708)
}
__typedArg0 := let__17413
__typedArg1 := tmp17708
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17710 := Call(__e, PrimFunc(symshen_4comb), let__17419, tmp17709)


ifres17705 = tmp17710


}

ifres17700 = ifres17705


}

ifres17695 = ifres17700


}

let__17411 := ifres17695
_ = let__17411

tmp17732 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17411)


if True == tmp17732 {
tmp17714 := Call(__e, PrimFunc(symshen_4_5syntax_6), V176)


let__17421 := tmp17714
_ = let__17421

tmp17728 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17421)


var ifres17715 Obj

if True == tmp17728 {
tmp17716 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17715 = tmp17716


} else {
tmp17717 := Call(__e, PrimFunc(symshen_4_5_1out), let__17421)


let__17422 := tmp17717
_ = let__17422

tmp17718 := Call(__e, PrimFunc(symshen_4in_1_6), let__17421)


let__17423 := tmp17718
_ = let__17423

tmp17719 := Call(__e, PrimFunc(symshen_4_5sc_6), let__17423)


let__17424 := tmp17719
_ = let__17424

tmp17727 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17424)


var ifres17720 Obj

if True == tmp17727 {
tmp17721 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17720 = tmp17721


} else {
tmp17722 := Call(__e, PrimFunc(symshen_4in_1_6), let__17424)


let__17425 := tmp17722
_ = let__17425

tmp17723 := Call(__e, PrimFunc(symshen_4autocomplete), let__17422)


tmp17724 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17723, Nil)
}
__typedArg0 := tmp17723
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17725 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17422, tmp17724)
}
__typedArg0 := let__17422
__typedArg1 := tmp17724
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17726 := Call(__e, PrimFunc(symshen_4comb), let__17425, tmp17725)


ifres17720 = tmp17726


}

ifres17715 = ifres17720


}

let__17420 := ifres17715
_ = let__17420

tmp17730 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17420)


if True == tmp17730 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__17420)
return
}


} else {
__e.Return(let__17411)
return
}


}, 1)

tmp17733 := Call(__e, ns2_1set, symshen_4_5c_1rule_6, tmp17693)


_ = tmp17733

tmp17734 := MakeNative(func(__e *ControlFlow) {
V192 := __e.Get(1)
_ = V192
tmp17763 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V192)
}
__typedArg0 := V192
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17755 Obj

if True == tmp17763 {
tmp17761 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V192)
}
__typedArg0 := V192
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17762 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp17761)
}
__typedArg0 := Nil
__typedArg1 := tmp17761
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17757 Obj

if True == tmp17762 {
tmp17759 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V192)
}
__typedArg0 := V192
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17760 := Call(__e, PrimFunc(symshen_4non_1terminal_2), tmp17759)


var ifres17758 Obj

if True == tmp17760 {
ifres17758 = True


} else {
ifres17758 = False


}

ifres17757 = ifres17758


} else {
ifres17757 = False


}

var ifres17756 Obj

if True == ifres17757 {
ifres17756 = True


} else {
ifres17756 = False


}

ifres17755 = ifres17756


} else {
ifres17755 = False


}

if True == ifres17755 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V192)
}
__typedArg0 := V192
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return
} else {
tmp17753 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V192)
}
__typedArg0 := V192
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17749 Obj

if True == tmp17753 {
tmp17751 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V192)
}
__typedArg0 := V192
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17752 := Call(__e, PrimFunc(symshen_4non_1terminal_2), tmp17751)


var ifres17750 Obj

if True == tmp17752 {
ifres17750 = True


} else {
ifres17750 = False


}

ifres17749 = ifres17750


} else {
ifres17749 = False


}

if True == ifres17749 {
tmp17735 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V192)
}
__typedArg0 := V192
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17736 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V192)
}
__typedArg0 := V192
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17737 := Call(__e, PrimFunc(symshen_4autocomplete), tmp17736)


tmp17738 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17737, Nil)
}
__typedArg0 := tmp17737
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17739 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17735, tmp17738)
}
__typedArg0 := tmp17735
__typedArg1 := tmp17738
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symappend, tmp17739)
}
__typedArg0 := symappend
__typedArg1 := tmp17739
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp17747 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V192)
}
__typedArg0 := V192
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17747 {
tmp17740 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V192)
}
__typedArg0 := V192
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17741 := Call(__e, PrimFunc(symshen_4autocomplete), tmp17740)


tmp17742 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V192)
}
__typedArg0 := V192
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17743 := Call(__e, PrimFunc(symshen_4autocomplete), tmp17742)


tmp17744 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17743, Nil)
}
__typedArg0 := tmp17743
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17745 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17741, tmp17744)
}
__typedArg0 := tmp17741
__typedArg1 := tmp17744
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp17745)
}
__typedArg0 := symcons
__typedArg1 := tmp17745
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

tmp17764 := Call(__e, ns2_1set, symshen_4autocomplete, tmp17734)


_ = tmp17764

tmp17765 := MakeNative(func(__e *ControlFlow) {
V193 := __e.Get(1)
_ = V193
tmp17771 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsymbol_2) {
return PrimIsSymbol(V193)
}
__typedArg0 := V193
return Call(__e, PrimFunc(symsymbol_2), __typedArg0)
})()

if True == tmp17771 {
tmp17767 := Call(__e, PrimFunc(symexplode), V193)


let__17426 := tmp17767
_ = let__17426

tmp17768 := MakeNative(func(__e *ControlFlow) {
Z195 := __e.Get(1)
_ = Z195
__e.TailApply(PrimFunc(symshen_4_5non_1terminal_2_6), Z195)
return
}, 1)

tmp17769 := Call(__e, PrimFunc(symcompile), tmp17768, let__17426)


if True == tmp17769 {
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

tmp17772 := Call(__e, ns2_1set, symshen_4non_1terminal_2, tmp17765)


_ = tmp17772

tmp17773 := MakeNative(func(__e *ControlFlow) {
V196 := __e.Get(1)
_ = V196
tmp17774 := Call(__e, PrimFunc(symshen_4_5packagenames_6), V196)


let__17428 := tmp17774
_ = let__17428

tmp17784 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17428)


var ifres17775 Obj

if True == tmp17784 {
tmp17776 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17775 = tmp17776


} else {
tmp17777 := Call(__e, PrimFunc(symshen_4in_1_6), let__17428)


let__17429 := tmp17777
_ = let__17429

tmp17778 := Call(__e, PrimFunc(symshen_4_5non_1terminal_1name_6), let__17429)


let__17430 := tmp17778
_ = let__17430

tmp17783 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17430)


var ifres17779 Obj

if True == tmp17783 {
tmp17780 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17779 = tmp17780


} else {
tmp17781 := Call(__e, PrimFunc(symshen_4in_1_6), let__17430)


let__17431 := tmp17781
_ = let__17431

tmp17782 := Call(__e, PrimFunc(symshen_4comb), let__17431, True)


ifres17779 = tmp17782


}

ifres17775 = ifres17779


}

let__17427 := ifres17775
_ = let__17427

tmp17802 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17427)


if True == tmp17802 {
tmp17785 := Call(__e, PrimFunc(symshen_4_5non_1terminal_1name_6), V196)


let__17433 := tmp17785
_ = let__17433

tmp17790 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17433)


var ifres17786 Obj

if True == tmp17790 {
tmp17787 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17786 = tmp17787


} else {
tmp17788 := Call(__e, PrimFunc(symshen_4in_1_6), let__17433)


let__17434 := tmp17788
_ = let__17434

tmp17789 := Call(__e, PrimFunc(symshen_4comb), let__17434, True)


ifres17786 = tmp17789


}

let__17432 := ifres17786
_ = let__17432

tmp17800 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17432)


if True == tmp17800 {
tmp17791 := Call(__e, PrimFunc(sym_5_b_6), V196)


let__17436 := tmp17791
_ = let__17436

tmp17796 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17436)


var ifres17792 Obj

if True == tmp17796 {
tmp17793 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17792 = tmp17793


} else {
tmp17794 := Call(__e, PrimFunc(symshen_4in_1_6), let__17436)


let__17437 := tmp17794
_ = let__17437

tmp17795 := Call(__e, PrimFunc(symshen_4comb), let__17437, False)


ifres17792 = tmp17795


}

let__17435 := ifres17792
_ = let__17435

tmp17798 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17435)


if True == tmp17798 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__17435)
return
}


} else {
__e.Return(let__17432)
return
}


} else {
__e.Return(let__17427)
return
}


}, 1)

tmp17803 := Call(__e, ns2_1set, symshen_4_5non_1terminal_2_6, tmp17773)


_ = tmp17803

tmp17804 := MakeNative(func(__e *ControlFlow) {
V208 := __e.Get(1)
_ = V208
tmp17805 := Call(__e, PrimFunc(symshen_4_5packagename_6), V208)


let__17439 := tmp17805
_ = let__17439

tmp17819 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17439)


var ifres17806 Obj

if True == tmp17819 {
tmp17807 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17806 = tmp17807


} else {
tmp17808 := Call(__e, PrimFunc(symshen_4in_1_6), let__17439)


let__17440 := tmp17808
_ = let__17440

tmp17818 := Call(__e, PrimFunc(symshen_4hds_a_2), let__17440, MakeString("."))


var ifres17809 Obj

if True == tmp17818 {
tmp17810 := Call(__e, PrimFunc(symtail), let__17440)


let__17441 := tmp17810
_ = let__17441

tmp17811 := Call(__e, PrimFunc(symshen_4_5packagenames_6), let__17441)


let__17442 := tmp17811
_ = let__17442

tmp17816 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17442)


var ifres17812 Obj

if True == tmp17816 {
tmp17813 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17812 = tmp17813


} else {
tmp17814 := Call(__e, PrimFunc(symshen_4in_1_6), let__17442)


let__17443 := tmp17814
_ = let__17443

tmp17815 := Call(__e, PrimFunc(symshen_4comb), let__17443, symshen_4skip)


ifres17812 = tmp17815


}

ifres17809 = ifres17812


} else {
tmp17817 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17809 = tmp17817


}

ifres17806 = ifres17809


}

let__17438 := ifres17806
_ = let__17438

tmp17833 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17438)


if True == tmp17833 {
tmp17820 := Call(__e, PrimFunc(symshen_4_5packagename_6), V208)


let__17445 := tmp17820
_ = let__17445

tmp17829 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17445)


var ifres17821 Obj

if True == tmp17829 {
tmp17822 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17821 = tmp17822


} else {
tmp17823 := Call(__e, PrimFunc(symshen_4in_1_6), let__17445)


let__17446 := tmp17823
_ = let__17446

tmp17828 := Call(__e, PrimFunc(symshen_4hds_a_2), let__17446, MakeString("."))


var ifres17824 Obj

if True == tmp17828 {
tmp17825 := Call(__e, PrimFunc(symtail), let__17446)


let__17447 := tmp17825
_ = let__17447

tmp17826 := Call(__e, PrimFunc(symshen_4comb), let__17447, symshen_4skip)


ifres17824 = tmp17826


} else {
tmp17827 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17824 = tmp17827


}

ifres17821 = ifres17824


}

let__17444 := ifres17821
_ = let__17444

tmp17831 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17444)


if True == tmp17831 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__17444)
return
}


} else {
__e.Return(let__17438)
return
}


}, 1)

tmp17834 := Call(__e, ns2_1set, symshen_4_5packagenames_6, tmp17804)


_ = tmp17834

tmp17835 := MakeNative(func(__e *ControlFlow) {
V219 := __e.Get(1)
_ = V219
tmp17836 := Call(__e, PrimFunc(symshen_4_5packagechar_6), V219)


let__17449 := tmp17836
_ = let__17449

tmp17846 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17449)


var ifres17837 Obj

if True == tmp17846 {
tmp17838 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17837 = tmp17838


} else {
tmp17839 := Call(__e, PrimFunc(symshen_4in_1_6), let__17449)


let__17450 := tmp17839
_ = let__17450

tmp17840 := Call(__e, PrimFunc(symshen_4_5packagename_6), let__17450)


let__17451 := tmp17840
_ = let__17451

tmp17845 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17451)


var ifres17841 Obj

if True == tmp17845 {
tmp17842 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17841 = tmp17842


} else {
tmp17843 := Call(__e, PrimFunc(symshen_4in_1_6), let__17451)


let__17452 := tmp17843
_ = let__17452

tmp17844 := Call(__e, PrimFunc(symshen_4comb), let__17452, symshen_4skip)


ifres17841 = tmp17844


}

ifres17837 = ifres17841


}

let__17448 := ifres17837
_ = let__17448

tmp17856 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17448)


if True == tmp17856 {
tmp17847 := Call(__e, PrimFunc(sym_5e_6), V219)


let__17454 := tmp17847
_ = let__17454

tmp17852 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17454)


var ifres17848 Obj

if True == tmp17852 {
tmp17849 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17848 = tmp17849


} else {
tmp17850 := Call(__e, PrimFunc(symshen_4in_1_6), let__17454)


let__17455 := tmp17850
_ = let__17455

tmp17851 := Call(__e, PrimFunc(symshen_4comb), let__17455, symshen_4skip)


ifres17848 = tmp17851


}

let__17453 := ifres17848
_ = let__17453

tmp17854 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17453)


if True == tmp17854 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__17453)
return
}


} else {
__e.Return(let__17448)
return
}


}, 1)

tmp17857 := Call(__e, ns2_1set, symshen_4_5packagename_6, tmp17835)


_ = tmp17857

tmp17858 := MakeNative(func(__e *ControlFlow) {
V228 := __e.Get(1)
_ = V228
tmp17868 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V228)
}
__typedArg0 := V228
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17859 Obj

if True == tmp17868 {
tmp17860 := Call(__e, PrimFunc(symhead), V228)


let__17457 := tmp17860
_ = let__17457

tmp17861 := Call(__e, PrimFunc(symtail), V228)


let__17458 := tmp17861
_ = let__17458

tmp17865 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__17457, MakeString("."))
}
__typedArg0 := let__17457
__typedArg1 := MakeString(".")
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

tmp17866 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp17865)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp17865
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres17862 Obj

if True == tmp17866 {
tmp17863 := Call(__e, PrimFunc(symshen_4comb), let__17458, symshen_4skip)


ifres17862 = tmp17863


} else {
tmp17864 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17862 = tmp17864


}

ifres17859 = ifres17862


} else {
tmp17867 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17859 = tmp17867


}

let__17456 := ifres17859
_ = let__17456

tmp17870 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17456)


if True == tmp17870 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__17456)
return
}


}, 1)

tmp17871 := Call(__e, ns2_1set, symshen_4_5packagechar_6, tmp17858)


_ = tmp17871

tmp17872 := MakeNative(func(__e *ControlFlow) {
V232 := __e.Get(1)
_ = V232
tmp17891 := Call(__e, PrimFunc(symshen_4hds_a_2), V232, MakeString("<"))


var ifres17873 Obj

if True == tmp17891 {
tmp17874 := Call(__e, PrimFunc(symtail), V232)


let__17460 := tmp17874
_ = let__17460

tmp17875 := Call(__e, PrimFunc(sym_5_b_6), let__17460)


let__17461 := tmp17875
_ = let__17461

tmp17889 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17461)


var ifres17876 Obj

if True == tmp17889 {
tmp17877 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17876 = tmp17877


} else {
tmp17878 := Call(__e, PrimFunc(symshen_4_5_1out), let__17461)


let__17462 := tmp17878
_ = let__17462

tmp17879 := Call(__e, PrimFunc(symshen_4in_1_6), let__17461)


let__17463 := tmp17879
_ = let__17463

tmp17883 := Call(__e, PrimFunc(symreverse), let__17462)


let__17464 := tmp17883
_ = let__17464

tmp17888 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__17464)
}
__typedArg0 := let__17464
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17884 Obj

if True == tmp17888 {
tmp17886 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__17464)
}
__typedArg0 := let__17464
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17887 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp17886, MakeString(">"))
}
__typedArg0 := tmp17886
__typedArg1 := MakeString(">")
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17885 Obj

if True == tmp17887 {
ifres17885 = True


} else {
ifres17885 = False


}

ifres17884 = ifres17885


} else {
ifres17884 = False


}

var ifres17880 Obj

if True == ifres17884 {
tmp17881 := Call(__e, PrimFunc(symshen_4comb), let__17463, symshen_4skip)


ifres17880 = tmp17881


} else {
tmp17882 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17880 = tmp17882


}

ifres17876 = ifres17880


}

ifres17873 = ifres17876


} else {
tmp17890 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17873 = tmp17890


}

let__17459 := ifres17873
_ = let__17459

tmp17893 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17459)


if True == tmp17893 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__17459)
return
}


}, 1)

tmp17894 := Call(__e, ns2_1set, symshen_4_5non_1terminal_1name_6, tmp17872)


_ = tmp17894

tmp17895 := MakeNative(func(__e *ControlFlow) {
V239 := __e.Get(1)
_ = V239
tmp17896 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(";"))
}
__typedArg0 := MakeString(";")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V239, tmp17896)
}
__typedArg0 := V239
__typedArg1 := tmp17896
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp17897 := Call(__e, ns2_1set, symshen_4semicolon_2, tmp17895)


_ = tmp17897

tmp17898 := MakeNative(func(__e *ControlFlow) {
V240 := __e.Get(1)
_ = V240
tmp17907 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V240)
}
__typedArg0 := V240
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17899 Obj

if True == tmp17907 {
tmp17900 := Call(__e, PrimFunc(symhead), V240)


let__17466 := tmp17900
_ = let__17466

tmp17901 := Call(__e, PrimFunc(symtail), V240)


let__17467 := tmp17901
_ = let__17467

tmp17905 := Call(__e, PrimFunc(symshen_4colon_1equal_2), let__17466)


var ifres17902 Obj

if True == tmp17905 {
tmp17903 := Call(__e, PrimFunc(symshen_4comb), let__17467, symshen_4skip)


ifres17902 = tmp17903


} else {
tmp17904 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17902 = tmp17904


}

ifres17899 = ifres17902


} else {
tmp17906 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17899 = tmp17906


}

let__17465 := ifres17899
_ = let__17465

tmp17909 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17465)


if True == tmp17909 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__17465)
return
}


}, 1)

tmp17910 := Call(__e, ns2_1set, symshen_4_5colon_1equal_6, tmp17898)


_ = tmp17910

tmp17911 := MakeNative(func(__e *ControlFlow) {
V244 := __e.Get(1)
_ = V244
tmp17912 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":="))
}
__typedArg0 := MakeString(":=")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp17912, V244)
}
__typedArg0 := tmp17912
__typedArg1 := V244
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp17913 := Call(__e, ns2_1set, symshen_4colon_1equal_2, tmp17911)


_ = tmp17913

tmp17914 := MakeNative(func(__e *ControlFlow) {
V245 := __e.Get(1)
_ = V245
tmp17915 := Call(__e, PrimFunc(symshen_4_5syntax_1item_6), V245)


let__17469 := tmp17915
_ = let__17469

tmp17928 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17469)


var ifres17916 Obj

if True == tmp17928 {
tmp17917 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17916 = tmp17917


} else {
tmp17918 := Call(__e, PrimFunc(symshen_4_5_1out), let__17469)


let__17470 := tmp17918
_ = let__17470

tmp17919 := Call(__e, PrimFunc(symshen_4in_1_6), let__17469)


let__17471 := tmp17919
_ = let__17471

tmp17920 := Call(__e, PrimFunc(symshen_4_5syntax_6), let__17471)


let__17472 := tmp17920
_ = let__17472

tmp17927 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17472)


var ifres17921 Obj

if True == tmp17927 {
tmp17922 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17921 = tmp17922


} else {
tmp17923 := Call(__e, PrimFunc(symshen_4_5_1out), let__17472)


let__17473 := tmp17923
_ = let__17473

tmp17924 := Call(__e, PrimFunc(symshen_4in_1_6), let__17472)


let__17474 := tmp17924
_ = let__17474

tmp17925 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17470, let__17473)
}
__typedArg0 := let__17470
__typedArg1 := let__17473
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17926 := Call(__e, PrimFunc(symshen_4comb), let__17474, tmp17925)


ifres17921 = tmp17926


}

ifres17916 = ifres17921


}

let__17468 := ifres17916
_ = let__17468

tmp17940 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17468)


if True == tmp17940 {
tmp17929 := Call(__e, PrimFunc(symshen_4_5syntax_1item_6), V245)


let__17476 := tmp17929
_ = let__17476

tmp17936 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17476)


var ifres17930 Obj

if True == tmp17936 {
tmp17931 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17930 = tmp17931


} else {
tmp17932 := Call(__e, PrimFunc(symshen_4_5_1out), let__17476)


let__17477 := tmp17932
_ = let__17477

tmp17933 := Call(__e, PrimFunc(symshen_4in_1_6), let__17476)


let__17478 := tmp17933
_ = let__17478

tmp17934 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17477, Nil)
}
__typedArg0 := let__17477
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17935 := Call(__e, PrimFunc(symshen_4comb), let__17478, tmp17934)


ifres17930 = tmp17935


}

let__17475 := ifres17930
_ = let__17475

tmp17938 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17475)


if True == tmp17938 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__17475)
return
}


} else {
__e.Return(let__17468)
return
}


}, 1)

tmp17941 := Call(__e, ns2_1set, symshen_4_5syntax_6, tmp17914)


_ = tmp17941

tmp17942 := MakeNative(func(__e *ControlFlow) {
V257 := __e.Get(1)
_ = V257
tmp17951 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V257)
}
__typedArg0 := V257
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17943 Obj

if True == tmp17951 {
tmp17944 := Call(__e, PrimFunc(symhead), V257)


let__17480 := tmp17944
_ = let__17480

tmp17945 := Call(__e, PrimFunc(symtail), V257)


let__17481 := tmp17945
_ = let__17481

tmp17949 := Call(__e, PrimFunc(symshen_4syntax_1item_2), let__17480)


var ifres17946 Obj

if True == tmp17949 {
tmp17947 := Call(__e, PrimFunc(symshen_4comb), let__17481, let__17480)


ifres17946 = tmp17947


} else {
tmp17948 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17946 = tmp17948


}

ifres17943 = ifres17946


} else {
tmp17950 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17943 = tmp17950


}

let__17479 := ifres17943
_ = let__17479

tmp17953 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17479)


if True == tmp17953 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__17479)
return
}


}, 1)

tmp17954 := Call(__e, ns2_1set, symshen_4_5syntax_1item_6, tmp17942)


_ = tmp17954

tmp17955 := MakeNative(func(__e *ControlFlow) {
V263 := __e.Get(1)
_ = V263
tmp17991 := Call(__e, PrimFunc(symshen_4colon_1equal_2), V263)


if True == tmp17991 {
__e.Return(False)
return
} else {
tmp17989 := Call(__e, PrimFunc(symshen_4semicolon_2), V263)


if True == tmp17989 {
__e.Return(False)
return
} else {
tmp17987 := Call(__e, PrimFunc(symatom_2), V263)


if True == tmp17987 {
__e.Return(True)
return
} else {
tmp17985 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V263)
}
__typedArg0 := V263
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17966 Obj

if True == tmp17985 {
tmp17983 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V263)
}
__typedArg0 := V263
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17984 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcons, tmp17983)
}
__typedArg0 := symcons
__typedArg1 := tmp17983
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17968 Obj

if True == tmp17984 {
tmp17981 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V263)
}
__typedArg0 := V263
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17982 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp17981)
}
__typedArg0 := tmp17981
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17970 Obj

if True == tmp17982 {
tmp17978 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V263)
}
__typedArg0 := V263
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17979 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp17978)
}
__typedArg0 := tmp17978
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17980 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp17979)
}
__typedArg0 := tmp17979
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17972 Obj

if True == tmp17980 {
tmp17974 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V263)
}
__typedArg0 := V263
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17975 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp17974)
}
__typedArg0 := tmp17974
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17976 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp17975)
}
__typedArg0 := tmp17975
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17977 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp17976)
}
__typedArg0 := Nil
__typedArg1 := tmp17976
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17973 Obj

if True == tmp17977 {
ifres17973 = True


} else {
ifres17973 = False


}

ifres17972 = ifres17973


} else {
ifres17972 = False


}

var ifres17971 Obj

if True == ifres17972 {
ifres17971 = True


} else {
ifres17971 = False


}

ifres17970 = ifres17971


} else {
ifres17970 = False


}

var ifres17969 Obj

if True == ifres17970 {
ifres17969 = True


} else {
ifres17969 = False


}

ifres17968 = ifres17969


} else {
ifres17968 = False


}

var ifres17967 Obj

if True == ifres17968 {
ifres17967 = True


} else {
ifres17967 = False


}

ifres17966 = ifres17967


} else {
ifres17966 = False


}

if True == ifres17966 {
tmp17962 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V263)
}
__typedArg0 := V263
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17963 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp17962)
}
__typedArg0 := tmp17962
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17964 := Call(__e, PrimFunc(symshen_4syntax_1item_2), tmp17963)


if True == tmp17964 {
tmp17957 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V263)
}
__typedArg0 := V263
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17958 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp17957)
}
__typedArg0 := tmp17957
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17959 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp17958)
}
__typedArg0 := tmp17958
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17960 := Call(__e, PrimFunc(symshen_4syntax_1item_2), tmp17959)


if True == tmp17960 {
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

tmp17992 := Call(__e, ns2_1set, symshen_4syntax_1item_2, tmp17955)


_ = tmp17992

tmp17993 := MakeNative(func(__e *ControlFlow) {
V264 := __e.Get(1)
_ = V264
tmp17994 := Call(__e, PrimFunc(symshen_4_5colon_1equal_6), V264)


let__17483 := tmp17994
_ = let__17483

tmp18020 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17483)


var ifres17995 Obj

if True == tmp18020 {
tmp17996 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17995 = tmp17996


} else {
tmp17997 := Call(__e, PrimFunc(symshen_4in_1_6), let__17483)


let__17484 := tmp17997
_ = let__17484

tmp18019 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__17484)
}
__typedArg0 := let__17484
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17998 Obj

if True == tmp18019 {
tmp17999 := Call(__e, PrimFunc(symhead), let__17484)


let__17485 := tmp17999
_ = let__17485

tmp18000 := Call(__e, PrimFunc(symtail), let__17484)


let__17486 := tmp18000
_ = let__17486

tmp18017 := Call(__e, PrimFunc(symshen_4hds_a_2), let__17486, symwhere)


var ifres18001 Obj

if True == tmp18017 {
tmp18002 := Call(__e, PrimFunc(symtail), let__17486)


let__17487 := tmp18002
_ = let__17487

tmp18015 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__17487)
}
__typedArg0 := let__17487
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18003 Obj

if True == tmp18015 {
tmp18004 := Call(__e, PrimFunc(symhead), let__17487)


let__17488 := tmp18004
_ = let__17488

tmp18005 := Call(__e, PrimFunc(symtail), let__17487)


let__17489 := tmp18005
_ = let__17489

tmp18012 := Call(__e, PrimFunc(symshen_4semicolon_2), let__17485)


tmp18013 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp18012)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp18012
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres18006 Obj

if True == tmp18013 {
tmp18007 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17485, Nil)
}
__typedArg0 := let__17485
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18008 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17488, tmp18007)
}
__typedArg0 := let__17488
__typedArg1 := tmp18007
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18009 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhere, tmp18008)
}
__typedArg0 := symwhere
__typedArg1 := tmp18008
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18010 := Call(__e, PrimFunc(symshen_4comb), let__17489, tmp18009)


ifres18006 = tmp18010


} else {
tmp18011 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres18006 = tmp18011


}

ifres18003 = ifres18006


} else {
tmp18014 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres18003 = tmp18014


}

ifres18001 = ifres18003


} else {
tmp18016 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres18001 = tmp18016


}

ifres17998 = ifres18001


} else {
tmp18018 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17998 = tmp18018


}

ifres17995 = ifres17998


}

let__17482 := ifres17995
_ = let__17482

tmp18039 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17482)


if True == tmp18039 {
tmp18021 := Call(__e, PrimFunc(symshen_4_5colon_1equal_6), V264)


let__17491 := tmp18021
_ = let__17491

tmp18035 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17491)


var ifres18022 Obj

if True == tmp18035 {
tmp18023 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres18022 = tmp18023


} else {
tmp18024 := Call(__e, PrimFunc(symshen_4in_1_6), let__17491)


let__17492 := tmp18024
_ = let__17492

tmp18034 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__17492)
}
__typedArg0 := let__17492
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18025 Obj

if True == tmp18034 {
tmp18026 := Call(__e, PrimFunc(symhead), let__17492)


let__17493 := tmp18026
_ = let__17493

tmp18027 := Call(__e, PrimFunc(symtail), let__17492)


let__17494 := tmp18027
_ = let__17494

tmp18031 := Call(__e, PrimFunc(symshen_4semicolon_2), let__17493)


tmp18032 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp18031)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp18031
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres18028 Obj

if True == tmp18032 {
tmp18029 := Call(__e, PrimFunc(symshen_4comb), let__17494, let__17493)


ifres18028 = tmp18029


} else {
tmp18030 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres18028 = tmp18030


}

ifres18025 = ifres18028


} else {
tmp18033 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres18025 = tmp18033


}

ifres18022 = ifres18025


}

let__17490 := ifres18022
_ = let__17490

tmp18037 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17490)


if True == tmp18037 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__17490)
return
}


} else {
__e.Return(let__17482)
return
}


}, 1)

tmp18040 := Call(__e, ns2_1set, symshen_4_5semantics_6, tmp17993)


_ = tmp18040

tmp18041 := MakeNative(func(__e *ControlFlow) {
V286 := __e.Get(1)
_ = V286
V287 := __e.Get(2)
_ = V287
V288 := __e.Get(3)
_ = V288
tmp18049 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V288)
}
__typedArg0 := Nil
__typedArg1 := V288
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp18049 {
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
tmp18047 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V288)
}
__typedArg0 := V288
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp18047 {
tmp18042 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V288)
}
__typedArg0 := V288
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18043 := Call(__e, PrimFunc(symshen_4c_1rule_1_6shen), V286, tmp18042, V287)


tmp18044 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V288)
}
__typedArg0 := V288
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18045 := Call(__e, PrimFunc(symshen_4c_1rules_1_6shen), V286, V287, tmp18044)


__e.TailApply(PrimFunc(symshen_4combine_1c_1code), tmp18043, tmp18045)
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

tmp18050 := Call(__e, ns2_1set, symshen_4c_1rules_1_6shen, tmp18041)


_ = tmp18050

tmp18051 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symfail))
return
}, 0)

tmp18052 := Call(__e, ns2_1set, symshen_4parse_1failure, tmp18051)


_ = tmp18052

tmp18053 := MakeNative(func(__e *ControlFlow) {
V289 := __e.Get(1)
_ = V289
V290 := __e.Get(2)
_ = V290
tmp18054 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symResult, Nil)
}
__typedArg0 := symResult
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18055 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4parse_1failure_2, tmp18054)
}
__typedArg0 := symshen_4parse_1failure_2
__typedArg1 := tmp18054
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18056 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symResult, Nil)
}
__typedArg0 := symResult
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18057 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V290, tmp18056)
}
__typedArg0 := V290
__typedArg1 := tmp18056
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18058 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18055, tmp18057)
}
__typedArg0 := tmp18055
__typedArg1 := tmp18057
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18059 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp18058)
}
__typedArg0 := symif
__typedArg1 := tmp18058
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18060 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18059, Nil)
}
__typedArg0 := tmp18059
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18061 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V289, tmp18060)
}
__typedArg0 := V289
__typedArg1 := tmp18060
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18062 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symResult, tmp18061)
}
__typedArg0 := symResult
__typedArg1 := tmp18061
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp18062)
}
__typedArg0 := symlet
__typedArg1 := tmp18062
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 2)

tmp18063 := Call(__e, ns2_1set, symshen_4combine_1c_1code, tmp18053)


_ = tmp18063

tmp18064 := MakeNative(func(__e *ControlFlow) {
V297 := __e.Get(1)
_ = V297
V298 := __e.Get(2)
_ = V298
V299 := __e.Get(3)
_ = V299
tmp18078 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V298)
}
__typedArg0 := V298
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18069 Obj

if True == tmp18078 {
tmp18076 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V298)
}
__typedArg0 := V298
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18077 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18076)
}
__typedArg0 := tmp18076
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18071 Obj

if True == tmp18077 {
tmp18073 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V298)
}
__typedArg0 := V298
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18074 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18073)
}
__typedArg0 := tmp18073
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18075 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp18074)
}
__typedArg0 := Nil
__typedArg1 := tmp18074
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18072 Obj

if True == tmp18075 {
ifres18072 = True


} else {
ifres18072 = False


}

ifres18071 = ifres18072


} else {
ifres18071 = False


}

var ifres18070 Obj

if True == ifres18071 {
ifres18070 = True


} else {
ifres18070 = False


}

ifres18069 = ifres18070


} else {
ifres18069 = False


}

if True == ifres18069 {
tmp18065 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V298)
}
__typedArg0 := V298
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18066 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V298)
}
__typedArg0 := V298
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18067 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18066)
}
__typedArg0 := tmp18066
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4yacc_1syntax), V297, V299, tmp18065, tmp18067)
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

tmp18079 := Call(__e, ns2_1set, symshen_4c_1rule_1_6shen, tmp18064)


_ = tmp18079

tmp18080 := MakeNative(func(__e *ControlFlow) {
V308 := __e.Get(1)
_ = V308
V309 := __e.Get(2)
_ = V309
V310 := __e.Get(3)
_ = V310
V311 := __e.Get(4)
_ = V311
tmp18144 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V310)
}
__typedArg0 := Nil
__typedArg1 := V310
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18122 Obj

if True == tmp18144 {
tmp18143 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V311)
}
__typedArg0 := V311
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18124 Obj

if True == tmp18143 {
tmp18141 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V311)
}
__typedArg0 := V311
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18142 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symwhere, tmp18141)
}
__typedArg0 := symwhere
__typedArg1 := tmp18141
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18126 Obj

if True == tmp18142 {
tmp18139 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V311)
}
__typedArg0 := V311
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18140 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18139)
}
__typedArg0 := tmp18139
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18128 Obj

if True == tmp18140 {
tmp18136 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V311)
}
__typedArg0 := V311
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18137 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18136)
}
__typedArg0 := tmp18136
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18138 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18137)
}
__typedArg0 := tmp18137
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18130 Obj

if True == tmp18138 {
tmp18132 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V311)
}
__typedArg0 := V311
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18133 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18132)
}
__typedArg0 := tmp18132
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18134 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18133)
}
__typedArg0 := tmp18133
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18135 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp18134)
}
__typedArg0 := Nil
__typedArg1 := tmp18134
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18131 Obj

if True == tmp18135 {
ifres18131 = True


} else {
ifres18131 = False


}

ifres18130 = ifres18131


} else {
ifres18130 = False


}

var ifres18129 Obj

if True == ifres18130 {
ifres18129 = True


} else {
ifres18129 = False


}

ifres18128 = ifres18129


} else {
ifres18128 = False


}

var ifres18127 Obj

if True == ifres18128 {
ifres18127 = True


} else {
ifres18127 = False


}

ifres18126 = ifres18127


} else {
ifres18126 = False


}

var ifres18125 Obj

if True == ifres18126 {
ifres18125 = True


} else {
ifres18125 = False


}

ifres18124 = ifres18125


} else {
ifres18124 = False


}

var ifres18123 Obj

if True == ifres18124 {
ifres18123 = True


} else {
ifres18123 = False


}

ifres18122 = ifres18123


} else {
ifres18122 = False


}

if True == ifres18122 {
tmp18081 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V311)
}
__typedArg0 := V311
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18082 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18081)
}
__typedArg0 := tmp18081
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18083 := Call(__e, PrimFunc(symshen_4process_1yacc_1semantics), tmp18082)


tmp18084 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V311)
}
__typedArg0 := V311
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18085 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18084)
}
__typedArg0 := tmp18084
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18086 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18085)
}
__typedArg0 := tmp18085
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18087 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V308, V309, Nil, tmp18086)


tmp18088 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4parse_1failure, Nil)
}
__typedArg0 := symshen_4parse_1failure
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18089 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18088, Nil)
}
__typedArg0 := tmp18088
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18090 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18087, tmp18089)
}
__typedArg0 := tmp18087
__typedArg1 := tmp18089
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18091 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18083, tmp18090)
}
__typedArg0 := tmp18083
__typedArg1 := tmp18090
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp18091)
}
__typedArg0 := symif
__typedArg1 := tmp18091
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp18120 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V310)
}
__typedArg0 := Nil
__typedArg1 := V310
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp18120 {
__e.TailApply(PrimFunc(symshen_4yacc_1semantics), V308, V309, V311)
return
} else {
tmp18118 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp18118 {
tmp18115 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18116 := Call(__e, PrimFunc(symshen_4non_1terminal_2), tmp18115)


if True == tmp18116 {
tmp18092 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18093 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4non_1terminalcode), V308, V309, tmp18092, tmp18093, V311)
return


} else {
tmp18112 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18113 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(tmp18112)
}
__typedArg0 := tmp18112
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

if True == tmp18113 {
tmp18094 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18095 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4variablecode), V308, V309, tmp18094, tmp18095, V311)
return


} else {
tmp18109 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18110 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym__, tmp18109)
}
__typedArg0 := sym__
__typedArg1 := tmp18109
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp18110 {
tmp18096 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18097 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4wildcardcode), V308, V309, tmp18096, tmp18097, V311)
return


} else {
tmp18106 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18107 := Call(__e, PrimFunc(symatom_2), tmp18106)


if True == tmp18107 {
tmp18098 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18099 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4terminalcode), V308, V309, tmp18098, tmp18099, V311)
return


} else {
tmp18103 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18104 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18103)
}
__typedArg0 := tmp18103
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp18104 {
tmp18100 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18101 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4conscode), V308, V309, tmp18100, tmp18101, V311)
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

tmp18145 := Call(__e, ns2_1set, symshen_4yacc_1syntax, tmp18080)


_ = tmp18145

tmp18146 := MakeNative(func(__e *ControlFlow) {
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
tmp18147 := Call(__e, PrimFunc(symconcat), symParse, V314)


let__17495 := tmp18147
_ = let__17495

tmp18148 := Call(__e, PrimFunc(symconcat), symAction, V314)


let__17496 := tmp18148
_ = let__17496

tmp18149 := Call(__e, PrimFunc(symconcat), symRemainder, V314)


let__17497 := tmp18149
_ = let__17497

tmp18150 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V313, Nil)
}
__typedArg0 := V313
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18151 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V314, tmp18150)
}
__typedArg0 := V314
__typedArg1 := tmp18150
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18152 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17495, Nil)
}
__typedArg0 := let__17495
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18153 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4parse_1failure_2, tmp18152)
}
__typedArg0 := symshen_4parse_1failure_2
__typedArg1 := tmp18152
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18154 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4parse_1failure, Nil)
}
__typedArg0 := symshen_4parse_1failure
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18155 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17495, Nil)
}
__typedArg0 := let__17495
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18156 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4in_1_6, tmp18155)
}
__typedArg0 := symshen_4in_1_6
__typedArg1 := tmp18155
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18157 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V312, let__17497, V315, V316)


tmp18158 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18157, Nil)
}
__typedArg0 := tmp18157
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18159 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18156, tmp18158)
}
__typedArg0 := tmp18156
__typedArg1 := tmp18158
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18160 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17497, tmp18159)
}
__typedArg0 := let__17497
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

let__17498 := tmp18161
_ = let__17498

tmp18172 := Call(__e, PrimFunc(symshen_4occurs_1check_2), V314, V316)


var ifres18169 Obj

if True == tmp18172 {
ifres18169 = True


} else {
tmp18171 := Call(__e, PrimFunc(symshen_4occurs_1check_2), let__17496, V316)


var ifres18170 Obj

if True == tmp18171 {
ifres18170 = True


} else {
ifres18170 = False


}

ifres18169 = ifres18170


}

var ifres18162 Obj

if True == ifres18169 {
tmp18163 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17495, Nil)
}
__typedArg0 := let__17495
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18164 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4_5_1out, tmp18163)
}
__typedArg0 := symshen_4_5_1out
__typedArg1 := tmp18163
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18165 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17498, Nil)
}
__typedArg0 := let__17498
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18166 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18164, tmp18165)
}
__typedArg0 := tmp18164
__typedArg1 := tmp18165
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18167 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17496, tmp18166)
}
__typedArg0 := let__17496
__typedArg1 := tmp18166
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18168 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp18167)
}
__typedArg0 := symlet
__typedArg1 := tmp18167
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

ifres18162 = tmp18168


} else {
ifres18162 = let__17498


}

tmp18173 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(ifres18162, Nil)
}
__typedArg0 := ifres18162
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18174 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18154, tmp18173)
}
__typedArg0 := tmp18154
__typedArg1 := tmp18173
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18175 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18153, tmp18174)
}
__typedArg0 := tmp18153
__typedArg1 := tmp18174
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18176 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp18175)
}
__typedArg0 := symif
__typedArg1 := tmp18175
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18177 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18176, Nil)
}
__typedArg0 := tmp18176
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18178 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18151, tmp18177)
}
__typedArg0 := tmp18151
__typedArg1 := tmp18177
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18179 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17495, tmp18178)
}
__typedArg0 := let__17495
__typedArg1 := tmp18178
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp18179)
}
__typedArg0 := symlet
__typedArg1 := tmp18179
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 5)

tmp18180 := Call(__e, ns2_1set, symshen_4non_1terminalcode, tmp18146)


_ = tmp18180

tmp18181 := MakeNative(func(__e *ControlFlow) {
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
tmp18182 := Call(__e, PrimFunc(symgensym), symRemainder)


let__17499 := tmp18182
_ = let__17499

tmp18183 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V322, Nil)
}
__typedArg0 := V322
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18184 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons_2, tmp18183)
}
__typedArg0 := symcons_2
__typedArg1 := tmp18183
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18185 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V322, Nil)
}
__typedArg0 := V322
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18186 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtail, tmp18185)
}
__typedArg0 := symtail
__typedArg1 := tmp18185
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18187 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V321, let__17499, V324, V325)


tmp18188 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18187, Nil)
}
__typedArg0 := tmp18187
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18189 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18186, tmp18188)
}
__typedArg0 := tmp18186
__typedArg1 := tmp18188
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18190 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17499, tmp18189)
}
__typedArg0 := let__17499
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

let__17500 := tmp18191
_ = let__17500

tmp18199 := Call(__e, PrimFunc(symshen_4occurs_1check_2), V323, V325)


var ifres18192 Obj

if True == tmp18199 {
tmp18193 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V322, Nil)
}
__typedArg0 := V322
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18194 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhead, tmp18193)
}
__typedArg0 := symhead
__typedArg1 := tmp18193
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18195 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17500, Nil)
}
__typedArg0 := let__17500
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18196 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18194, tmp18195)
}
__typedArg0 := tmp18194
__typedArg1 := tmp18195
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18197 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V323, tmp18196)
}
__typedArg0 := V323
__typedArg1 := tmp18196
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18198 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp18197)
}
__typedArg0 := symlet
__typedArg1 := tmp18197
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

ifres18192 = tmp18198


} else {
ifres18192 = let__17500


}

tmp18200 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4parse_1failure, Nil)
}
__typedArg0 := symshen_4parse_1failure
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18201 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18200, Nil)
}
__typedArg0 := tmp18200
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18202 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(ifres18192, tmp18201)
}
__typedArg0 := ifres18192
__typedArg1 := tmp18201
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18203 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18184, tmp18202)
}
__typedArg0 := tmp18184
__typedArg1 := tmp18202
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp18203)
}
__typedArg0 := symif
__typedArg1 := tmp18203
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 5)

tmp18204 := Call(__e, ns2_1set, symshen_4variablecode, tmp18181)


_ = tmp18204

tmp18205 := MakeNative(func(__e *ControlFlow) {
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
tmp18206 := Call(__e, PrimFunc(symgensym), symRemainder)


let__17501 := tmp18206
_ = let__17501

tmp18207 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V329, Nil)
}
__typedArg0 := V329
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18208 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons_2, tmp18207)
}
__typedArg0 := symcons_2
__typedArg1 := tmp18207
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18209 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V329, Nil)
}
__typedArg0 := V329
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18210 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtail, tmp18209)
}
__typedArg0 := symtail
__typedArg1 := tmp18209
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18211 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V328, let__17501, V331, V332)


tmp18212 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18211, Nil)
}
__typedArg0 := tmp18211
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18213 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18210, tmp18212)
}
__typedArg0 := tmp18210
__typedArg1 := tmp18212
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18214 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17501, tmp18213)
}
__typedArg0 := let__17501
__typedArg1 := tmp18213
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18215 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp18214)
}
__typedArg0 := symlet
__typedArg1 := tmp18214
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18216 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4parse_1failure, Nil)
}
__typedArg0 := symshen_4parse_1failure
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18217 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18216, Nil)
}
__typedArg0 := tmp18216
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18218 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18215, tmp18217)
}
__typedArg0 := tmp18215
__typedArg1 := tmp18217
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18219 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18208, tmp18218)
}
__typedArg0 := tmp18208
__typedArg1 := tmp18218
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp18219)
}
__typedArg0 := symif
__typedArg1 := tmp18219
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 5)

tmp18220 := Call(__e, ns2_1set, symshen_4wildcardcode, tmp18205)


_ = tmp18220

tmp18221 := MakeNative(func(__e *ControlFlow) {
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
tmp18222 := Call(__e, PrimFunc(symgensym), symRemainder)


let__17502 := tmp18222
_ = let__17502

tmp18223 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V336, Nil)
}
__typedArg0 := V336
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18224 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V335, tmp18223)
}
__typedArg0 := V335
__typedArg1 := tmp18223
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18225 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4hds_a_2, tmp18224)
}
__typedArg0 := symshen_4hds_a_2
__typedArg1 := tmp18224
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18226 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V335, Nil)
}
__typedArg0 := V335
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18227 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtail, tmp18226)
}
__typedArg0 := symtail
__typedArg1 := tmp18226
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18228 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V334, let__17502, V337, V338)


tmp18229 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18228, Nil)
}
__typedArg0 := tmp18228
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18230 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18227, tmp18229)
}
__typedArg0 := tmp18227
__typedArg1 := tmp18229
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18231 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17502, tmp18230)
}
__typedArg0 := let__17502
__typedArg1 := tmp18230
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18232 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp18231)
}
__typedArg0 := symlet
__typedArg1 := tmp18231
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18233 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4parse_1failure, Nil)
}
__typedArg0 := symshen_4parse_1failure
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18234 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18233, Nil)
}
__typedArg0 := tmp18233
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18235 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18232, tmp18234)
}
__typedArg0 := tmp18232
__typedArg1 := tmp18234
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18236 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18225, tmp18235)
}
__typedArg0 := tmp18225
__typedArg1 := tmp18235
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp18236)
}
__typedArg0 := symif
__typedArg1 := tmp18236
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 5)

tmp18237 := Call(__e, ns2_1set, symshen_4terminalcode, tmp18221)


_ = tmp18237

tmp18238 := MakeNative(func(__e *ControlFlow) {
V347 := __e.Get(1)
_ = V347
V348 := __e.Get(2)
_ = V348
tmp18244 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V347)
}
__typedArg0 := V347
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18240 Obj

if True == tmp18244 {
tmp18242 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V347)
}
__typedArg0 := V347
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18243 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp18242, V348)
}
__typedArg0 := tmp18242
__typedArg1 := V348
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18241 Obj

if True == tmp18243 {
ifres18241 = True


} else {
ifres18241 = False


}

ifres18240 = ifres18241


} else {
ifres18240 = False


}

if True == ifres18240 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}, 2)

tmp18245 := Call(__e, ns2_1set, symshen_4hds_a_2, tmp18238)


_ = tmp18245

tmp18246 := MakeNative(func(__e *ControlFlow) {
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
tmp18247 := Call(__e, PrimFunc(symgensym), symRemainder)


let__17503 := tmp18247
_ = let__17503

tmp18248 := Call(__e, PrimFunc(symgensym), symHd)


let__17504 := tmp18248
_ = let__17504

tmp18249 := Call(__e, PrimFunc(symgensym), symTl)


let__17505 := tmp18249
_ = let__17505

tmp18250 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V350, Nil)
}
__typedArg0 := V350
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18251 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4ccons_2, tmp18250)
}
__typedArg0 := symshen_4ccons_2
__typedArg1 := tmp18250
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18252 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V350, Nil)
}
__typedArg0 := V350
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18253 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhead, tmp18252)
}
__typedArg0 := symhead
__typedArg1 := tmp18252
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18254 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V350, Nil)
}
__typedArg0 := V350
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18255 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtail, tmp18254)
}
__typedArg0 := symtail
__typedArg1 := tmp18254
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18256 := Call(__e, PrimFunc(symshen_4decons), V351)


tmp18257 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5end_6, Nil)
}
__typedArg0 := sym_5end_6
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18258 := Call(__e, PrimFunc(symappend), tmp18256, tmp18257)


tmp18259 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V349, let__17505, V352, V353)


tmp18260 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18259, Nil)
}
__typedArg0 := tmp18259
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18261 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4processed, tmp18260)
}
__typedArg0 := symshen_4processed
__typedArg1 := tmp18260
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18262 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V349, let__17504, tmp18258, tmp18261)


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
return PrimCons(tmp18255, tmp18263)
}
__typedArg0 := tmp18255
__typedArg1 := tmp18263
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18265 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17505, tmp18264)
}
__typedArg0 := let__17505
__typedArg1 := tmp18264
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18266 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18253, tmp18265)
}
__typedArg0 := tmp18253
__typedArg1 := tmp18265
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18267 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17504, tmp18266)
}
__typedArg0 := let__17504
__typedArg1 := tmp18266
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18268 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp18267)
}
__typedArg0 := symlet
__typedArg1 := tmp18267
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18269 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4parse_1failure, Nil)
}
__typedArg0 := symshen_4parse_1failure
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18270 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18269, Nil)
}
__typedArg0 := tmp18269
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18271 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18268, tmp18270)
}
__typedArg0 := tmp18268
__typedArg1 := tmp18270
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18272 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18251, tmp18271)
}
__typedArg0 := tmp18251
__typedArg1 := tmp18271
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp18272)
}
__typedArg0 := symif
__typedArg1 := tmp18272
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 5)

tmp18273 := Call(__e, ns2_1set, symshen_4conscode, tmp18246)


_ = tmp18273

tmp18274 := MakeNative(func(__e *ControlFlow) {
V367 := __e.Get(1)
_ = V367
tmp18286 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V367)
}
__typedArg0 := V367
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18282 Obj

if True == tmp18286 {
tmp18284 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V367)
}
__typedArg0 := V367
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18285 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18284)
}
__typedArg0 := tmp18284
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18283 Obj

if True == tmp18285 {
ifres18283 = True


} else {
ifres18283 = False


}

ifres18282 = ifres18283


} else {
ifres18282 = False


}

if True == ifres18282 {
__e.Return(True)
return
} else {
tmp18280 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V367)
}
__typedArg0 := V367
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18276 Obj

if True == tmp18280 {
tmp18278 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V367)
}
__typedArg0 := V367
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18279 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp18278)
}
__typedArg0 := Nil
__typedArg1 := tmp18278
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18277 Obj

if True == tmp18279 {
ifres18277 = True


} else {
ifres18277 = False


}

ifres18276 = ifres18277


} else {
ifres18276 = False


}

if True == ifres18276 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp18287 := Call(__e, ns2_1set, symshen_4ccons_2, tmp18274)


_ = tmp18287

tmp18288 := MakeNative(func(__e *ControlFlow) {
V368 := __e.Get(1)
_ = V368
tmp18315 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V368)
}
__typedArg0 := V368
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18296 Obj

if True == tmp18315 {
tmp18313 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V368)
}
__typedArg0 := V368
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18314 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcons, tmp18313)
}
__typedArg0 := symcons
__typedArg1 := tmp18313
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18298 Obj

if True == tmp18314 {
tmp18311 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V368)
}
__typedArg0 := V368
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18312 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18311)
}
__typedArg0 := tmp18311
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18300 Obj

if True == tmp18312 {
tmp18308 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V368)
}
__typedArg0 := V368
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18309 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18308)
}
__typedArg0 := tmp18308
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18310 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18309)
}
__typedArg0 := tmp18309
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18302 Obj

if True == tmp18310 {
tmp18304 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V368)
}
__typedArg0 := V368
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18305 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18304)
}
__typedArg0 := tmp18304
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18306 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18305)
}
__typedArg0 := tmp18305
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18307 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp18306)
}
__typedArg0 := Nil
__typedArg1 := tmp18306
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18303 Obj

if True == tmp18307 {
ifres18303 = True


} else {
ifres18303 = False


}

ifres18302 = ifres18303


} else {
ifres18302 = False


}

var ifres18301 Obj

if True == ifres18302 {
ifres18301 = True


} else {
ifres18301 = False


}

ifres18300 = ifres18301


} else {
ifres18300 = False


}

var ifres18299 Obj

if True == ifres18300 {
ifres18299 = True


} else {
ifres18299 = False


}

ifres18298 = ifres18299


} else {
ifres18298 = False


}

var ifres18297 Obj

if True == ifres18298 {
ifres18297 = True


} else {
ifres18297 = False


}

ifres18296 = ifres18297


} else {
ifres18296 = False


}

if True == ifres18296 {
tmp18289 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V368)
}
__typedArg0 := V368
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18290 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18289)
}
__typedArg0 := tmp18289
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18291 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V368)
}
__typedArg0 := V368
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18292 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18291)
}
__typedArg0 := tmp18291
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18293 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18292)
}
__typedArg0 := tmp18292
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18294 := Call(__e, PrimFunc(symshen_4decons), tmp18293)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18290, tmp18294)
}
__typedArg0 := tmp18290
__typedArg1 := tmp18294
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return(V368)
return
}


}, 1)

tmp18316 := Call(__e, ns2_1set, symshen_4decons, tmp18288)


_ = tmp18316

tmp18317 := MakeNative(func(__e *ControlFlow) {
V369 := __e.Get(1)
_ = V369
V370 := __e.Get(2)
_ = V370
tmp18318 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V370, Nil)
}
__typedArg0 := V370
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V369, tmp18318)
}
__typedArg0 := V369
__typedArg1 := tmp18318
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 2)

tmp18319 := Call(__e, ns2_1set, symshen_4comb, tmp18317)


_ = tmp18319

tmp18320 := MakeNative(func(__e *ControlFlow) {
V375 := __e.Get(1)
_ = V375
V376 := __e.Get(2)
_ = V376
V377 := __e.Get(3)
_ = V377
tmp18340 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V377)
}
__typedArg0 := V377
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18327 Obj

if True == tmp18340 {
tmp18338 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V377)
}
__typedArg0 := V377
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18339 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4processed, tmp18338)
}
__typedArg0 := symshen_4processed
__typedArg1 := tmp18338
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18329 Obj

if True == tmp18339 {
tmp18336 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V377)
}
__typedArg0 := V377
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18337 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18336)
}
__typedArg0 := tmp18336
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18331 Obj

if True == tmp18337 {
tmp18333 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V377)
}
__typedArg0 := V377
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18334 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18333)
}
__typedArg0 := tmp18333
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18335 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp18334)
}
__typedArg0 := Nil
__typedArg1 := tmp18334
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18332 Obj

if True == tmp18335 {
ifres18332 = True


} else {
ifres18332 = False


}

ifres18331 = ifres18332


} else {
ifres18331 = False


}

var ifres18330 Obj

if True == ifres18331 {
ifres18330 = True


} else {
ifres18330 = False


}

ifres18329 = ifres18330


} else {
ifres18329 = False


}

var ifres18328 Obj

if True == ifres18329 {
ifres18328 = True


} else {
ifres18328 = False


}

ifres18327 = ifres18328


} else {
ifres18327 = False


}

if True == ifres18327 {
tmp18321 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V377)
}
__typedArg0 := V377
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18321)
}
__typedArg0 := tmp18321
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return


} else {
tmp18322 := Call(__e, PrimFunc(symshen_4process_1yacc_1semantics), V377)


let__17506 := tmp18322
_ = let__17506

tmp18323 := Call(__e, PrimFunc(symshen_4use_1type_1info), V375, let__17506)


let__17507 := tmp18323
_ = let__17507

tmp18324 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17507, Nil)
}
__typedArg0 := let__17507
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18325 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V376, tmp18324)
}
__typedArg0 := V376
__typedArg1 := tmp18324
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4comb, tmp18325)
}
__typedArg0 := symshen_4comb
__typedArg1 := tmp18325
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}


}, 3)

tmp18341 := Call(__e, ns2_1set, symshen_4yacc_1semantics, tmp18320)


_ = tmp18341

tmp18342 := MakeNative(func(__e *ControlFlow) {
V383 := __e.Get(1)
_ = V383
V384 := __e.Get(2)
_ = V384
tmp18530 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18351 Obj

if True == tmp18530 {
tmp18528 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18529 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_i, tmp18528)
}
__typedArg0 := sym_i
__typedArg1 := tmp18528
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18353 Obj

if True == tmp18529 {
tmp18526 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18527 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18526)
}
__typedArg0 := tmp18526
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18355 Obj

if True == tmp18527 {
tmp18523 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18524 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18523)
}
__typedArg0 := tmp18523
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18525 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18524)
}
__typedArg0 := tmp18524
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18357 Obj

if True == tmp18525 {
tmp18519 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18520 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18519)
}
__typedArg0 := tmp18519
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18521 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18520)
}
__typedArg0 := tmp18520
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18522 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlist, tmp18521)
}
__typedArg0 := symlist
__typedArg1 := tmp18521
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18359 Obj

if True == tmp18522 {
tmp18515 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18516 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18515)
}
__typedArg0 := tmp18515
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18517 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18516)
}
__typedArg0 := tmp18516
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18518 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18517)
}
__typedArg0 := tmp18517
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18361 Obj

if True == tmp18518 {
tmp18510 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18511 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18510)
}
__typedArg0 := tmp18510
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18512 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18511)
}
__typedArg0 := tmp18511
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18513 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18512)
}
__typedArg0 := tmp18512
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18514 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp18513)
}
__typedArg0 := Nil
__typedArg1 := tmp18513
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18363 Obj

if True == tmp18514 {
tmp18507 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18508 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18507)
}
__typedArg0 := tmp18507
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18509 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18508)
}
__typedArg0 := tmp18508
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18365 Obj

if True == tmp18509 {
tmp18503 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18504 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18503)
}
__typedArg0 := tmp18503
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18505 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18504)
}
__typedArg0 := tmp18504
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18506 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_1_1_6, tmp18505)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp18505
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18367 Obj

if True == tmp18506 {
tmp18499 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18500 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18499)
}
__typedArg0 := tmp18499
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

var ifres18369 Obj

if True == tmp18502 {
tmp18494 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18495 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18494)
}
__typedArg0 := tmp18494
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18496 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18495)
}
__typedArg0 := tmp18495
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18497 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18496)
}
__typedArg0 := tmp18496
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18498 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18497)
}
__typedArg0 := tmp18497
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18371 Obj

if True == tmp18498 {
tmp18488 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18489)
}
__typedArg0 := tmp18489
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18491 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18490)
}
__typedArg0 := tmp18490
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18492 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18491)
}
__typedArg0 := tmp18491
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18493 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symstr, tmp18492)
}
__typedArg0 := symstr
__typedArg1 := tmp18492
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18373 Obj

if True == tmp18493 {
tmp18482 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18483)
}
__typedArg0 := tmp18483
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18485 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18484)
}
__typedArg0 := tmp18484
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18486 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18485)
}
__typedArg0 := tmp18485
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18487 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18486)
}
__typedArg0 := tmp18486
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18375 Obj

if True == tmp18487 {
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18479)
}
__typedArg0 := tmp18479
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18481 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18480)
}
__typedArg0 := tmp18480
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18377 Obj

if True == tmp18481 {
tmp18467 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18468 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18467)
}
__typedArg0 := tmp18467
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18469)
}
__typedArg0 := tmp18469
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18471 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18470)
}
__typedArg0 := tmp18470
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18472 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18471)
}
__typedArg0 := tmp18471
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18473 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18472)
}
__typedArg0 := tmp18472
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18474 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlist, tmp18473)
}
__typedArg0 := symlist
__typedArg1 := tmp18473
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18379 Obj

if True == tmp18474 {
tmp18459 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18460 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18459)
}
__typedArg0 := tmp18459
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18461)
}
__typedArg0 := tmp18461
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18463 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18462)
}
__typedArg0 := tmp18462
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18464 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18463)
}
__typedArg0 := tmp18463
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18465 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18464)
}
__typedArg0 := tmp18464
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18466 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18465)
}
__typedArg0 := tmp18465
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18381 Obj

if True == tmp18466 {
tmp18450 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18451 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18450)
}
__typedArg0 := tmp18450
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18452 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18451)
}
__typedArg0 := tmp18451
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18453 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18452)
}
__typedArg0 := tmp18452
return Call(__e, PrimFunc(symhd), __typedArg0)
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18456)
}
__typedArg0 := tmp18456
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18458 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp18457)
}
__typedArg0 := Nil
__typedArg1 := tmp18457
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18383 Obj

if True == tmp18458 {
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18447)
}
__typedArg0 := tmp18447
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18449 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18448)
}
__typedArg0 := tmp18448
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18385 Obj

if True == tmp18449 {
tmp18435 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18436 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18435)
}
__typedArg0 := tmp18435
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18437)
}
__typedArg0 := tmp18437
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18439 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18438)
}
__typedArg0 := tmp18438
return Call(__e, PrimFunc(symtl), __typedArg0)
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp18441)
}
__typedArg0 := Nil
__typedArg1 := tmp18441
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18387 Obj

if True == tmp18442 {
tmp18430 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18431 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18430)
}
__typedArg0 := tmp18430
return Call(__e, PrimFunc(symtl), __typedArg0)
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18433)
}
__typedArg0 := tmp18433
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18389 Obj

if True == tmp18434 {
tmp18424 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18426)
}
__typedArg0 := tmp18426
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18428 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18427)
}
__typedArg0 := tmp18427
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18429 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_j, tmp18428)
}
__typedArg0 := sym_j
__typedArg1 := tmp18428
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18391 Obj

if True == tmp18429 {
tmp18418 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18420)
}
__typedArg0 := tmp18420
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18422 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18421)
}
__typedArg0 := tmp18421
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18423 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp18422)
}
__typedArg0 := Nil
__typedArg1 := tmp18422
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18393 Obj

if True == tmp18423 {
tmp18405 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18406 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18405)
}
__typedArg0 := tmp18405
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18407 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18406)
}
__typedArg0 := tmp18406
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18408 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18407)
}
__typedArg0 := tmp18407
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18409 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18410 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18409)
}
__typedArg0 := tmp18409
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18411 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18410)
}
__typedArg0 := tmp18410
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18412 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18411)
}
__typedArg0 := tmp18411
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18413 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18412)
}
__typedArg0 := tmp18412
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18414 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18413)
}
__typedArg0 := tmp18413
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18415 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18414)
}
__typedArg0 := tmp18414
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18416 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18415)
}
__typedArg0 := tmp18415
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18417 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp18408, tmp18416)
}
__typedArg0 := tmp18408
__typedArg1 := tmp18416
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18395 Obj

if True == tmp18417 {
tmp18397 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18398 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18397)
}
__typedArg0 := tmp18397
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18399 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18398)
}
__typedArg0 := tmp18398
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18400 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18399)
}
__typedArg0 := tmp18399
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18401 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18400)
}
__typedArg0 := tmp18400
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18402 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18401)
}
__typedArg0 := tmp18401
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18403 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18402)
}
__typedArg0 := tmp18402
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18404 := Call(__e, PrimFunc(symshen_4monomorphic_2), tmp18403)


var ifres18396 Obj

if True == tmp18404 {
ifres18396 = True


} else {
ifres18396 = False


}

ifres18395 = ifres18396


} else {
ifres18395 = False


}

var ifres18394 Obj

if True == ifres18395 {
ifres18394 = True


} else {
ifres18394 = False


}

ifres18393 = ifres18394


} else {
ifres18393 = False


}

var ifres18392 Obj

if True == ifres18393 {
ifres18392 = True


} else {
ifres18392 = False


}

ifres18391 = ifres18392


} else {
ifres18391 = False


}

var ifres18390 Obj

if True == ifres18391 {
ifres18390 = True


} else {
ifres18390 = False


}

ifres18389 = ifres18390


} else {
ifres18389 = False


}

var ifres18388 Obj

if True == ifres18389 {
ifres18388 = True


} else {
ifres18388 = False


}

ifres18387 = ifres18388


} else {
ifres18387 = False


}

var ifres18386 Obj

if True == ifres18387 {
ifres18386 = True


} else {
ifres18386 = False


}

ifres18385 = ifres18386


} else {
ifres18385 = False


}

var ifres18384 Obj

if True == ifres18385 {
ifres18384 = True


} else {
ifres18384 = False


}

ifres18383 = ifres18384


} else {
ifres18383 = False


}

var ifres18382 Obj

if True == ifres18383 {
ifres18382 = True


} else {
ifres18382 = False


}

ifres18381 = ifres18382


} else {
ifres18381 = False


}

var ifres18380 Obj

if True == ifres18381 {
ifres18380 = True


} else {
ifres18380 = False


}

ifres18379 = ifres18380


} else {
ifres18379 = False


}

var ifres18378 Obj

if True == ifres18379 {
ifres18378 = True


} else {
ifres18378 = False


}

ifres18377 = ifres18378


} else {
ifres18377 = False


}

var ifres18376 Obj

if True == ifres18377 {
ifres18376 = True


} else {
ifres18376 = False


}

ifres18375 = ifres18376


} else {
ifres18375 = False


}

var ifres18374 Obj

if True == ifres18375 {
ifres18374 = True


} else {
ifres18374 = False


}

ifres18373 = ifres18374


} else {
ifres18373 = False


}

var ifres18372 Obj

if True == ifres18373 {
ifres18372 = True


} else {
ifres18372 = False


}

ifres18371 = ifres18372


} else {
ifres18371 = False


}

var ifres18370 Obj

if True == ifres18371 {
ifres18370 = True


} else {
ifres18370 = False


}

ifres18369 = ifres18370


} else {
ifres18369 = False


}

var ifres18368 Obj

if True == ifres18369 {
ifres18368 = True


} else {
ifres18368 = False


}

ifres18367 = ifres18368


} else {
ifres18367 = False


}

var ifres18366 Obj

if True == ifres18367 {
ifres18366 = True


} else {
ifres18366 = False


}

ifres18365 = ifres18366


} else {
ifres18365 = False


}

var ifres18364 Obj

if True == ifres18365 {
ifres18364 = True


} else {
ifres18364 = False


}

ifres18363 = ifres18364


} else {
ifres18363 = False


}

var ifres18362 Obj

if True == ifres18363 {
ifres18362 = True


} else {
ifres18362 = False


}

ifres18361 = ifres18362


} else {
ifres18361 = False


}

var ifres18360 Obj

if True == ifres18361 {
ifres18360 = True


} else {
ifres18360 = False


}

ifres18359 = ifres18360


} else {
ifres18359 = False


}

var ifres18358 Obj

if True == ifres18359 {
ifres18358 = True


} else {
ifres18358 = False


}

ifres18357 = ifres18358


} else {
ifres18357 = False


}

var ifres18356 Obj

if True == ifres18357 {
ifres18356 = True


} else {
ifres18356 = False


}

ifres18355 = ifres18356


} else {
ifres18355 = False


}

var ifres18354 Obj

if True == ifres18355 {
ifres18354 = True


} else {
ifres18354 = False


}

ifres18353 = ifres18354


} else {
ifres18353 = False


}

var ifres18352 Obj

if True == ifres18353 {
ifres18352 = True


} else {
ifres18352 = False


}

ifres18351 = ifres18352


} else {
ifres18351 = False


}

if True == ifres18351 {
tmp18343 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18344 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18343)
}
__typedArg0 := tmp18343
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18345 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18344)
}
__typedArg0 := tmp18344
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18346 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18345)
}
__typedArg0 := tmp18345
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18347 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18346)
}
__typedArg0 := tmp18346
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18348 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18347)
}
__typedArg0 := tmp18347
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18349 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V384, tmp18348)
}
__typedArg0 := V384
__typedArg1 := tmp18348
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtype, tmp18349)
}
__typedArg0 := symtype
__typedArg1 := tmp18349
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return(V384)
return
}


}, 2)

tmp18531 := Call(__e, ns2_1set, symshen_4use_1type_1info, tmp18342)


_ = tmp18531

tmp18532 := MakeNative(func(__e *ControlFlow) {
V387 := __e.Get(1)
_ = V387
tmp18542 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(V387)
}
__typedArg0 := V387
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

if True == tmp18542 {
__e.Return(False)
return
} else {
tmp18540 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V387)
}
__typedArg0 := V387
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp18540 {
tmp18537 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V387)
}
__typedArg0 := V387
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18538 := Call(__e, PrimFunc(symshen_4monomorphic_2), tmp18537)


if True == tmp18538 {
tmp18534 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V387)
}
__typedArg0 := V387
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18535 := Call(__e, PrimFunc(symshen_4monomorphic_2), tmp18534)


if True == tmp18535 {
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

tmp18543 := Call(__e, ns2_1set, symshen_4monomorphic_2, tmp18532)


_ = tmp18543

tmp18544 := MakeNative(func(__e *ControlFlow) {
V388 := __e.Get(1)
_ = V388
tmp18570 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V388)
}
__typedArg0 := V388
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18552 Obj

if True == tmp18570 {
tmp18568 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V388)
}
__typedArg0 := V388
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18569 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symprotect, tmp18568)
}
__typedArg0 := symprotect
__typedArg1 := tmp18568
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18554 Obj

if True == tmp18569 {
tmp18566 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V388)
}
__typedArg0 := V388
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18567 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18566)
}
__typedArg0 := tmp18566
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18556 Obj

if True == tmp18567 {
tmp18563 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V388)
}
__typedArg0 := V388
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18564 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18563)
}
__typedArg0 := tmp18563
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18565 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp18564)
}
__typedArg0 := Nil
__typedArg1 := tmp18564
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18558 Obj

if True == tmp18565 {
tmp18560 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V388)
}
__typedArg0 := V388
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18561 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18560)
}
__typedArg0 := tmp18560
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18562 := Call(__e, PrimFunc(symshen_4non_1terminal_2), tmp18561)


var ifres18559 Obj

if True == tmp18562 {
ifres18559 = True


} else {
ifres18559 = False


}

ifres18558 = ifres18559


} else {
ifres18558 = False


}

var ifres18557 Obj

if True == ifres18558 {
ifres18557 = True


} else {
ifres18557 = False


}

ifres18556 = ifres18557


} else {
ifres18556 = False


}

var ifres18555 Obj

if True == ifres18556 {
ifres18555 = True


} else {
ifres18555 = False


}

ifres18554 = ifres18555


} else {
ifres18554 = False


}

var ifres18553 Obj

if True == ifres18554 {
ifres18553 = True


} else {
ifres18553 = False


}

ifres18552 = ifres18553


} else {
ifres18552 = False


}

if True == ifres18552 {
tmp18545 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V388)
}
__typedArg0 := V388
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18545)
}
__typedArg0 := tmp18545
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return


} else {
tmp18550 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V388)
}
__typedArg0 := V388
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp18550 {
tmp18546 := MakeNative(func(__e *ControlFlow) {
Z389 := __e.Get(1)
_ = Z389
__e.TailApply(PrimFunc(symshen_4process_1yacc_1semantics), Z389)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp18546, V388)
return


} else {
tmp18548 := Call(__e, PrimFunc(symshen_4non_1terminal_2), V388)


if True == tmp18548 {
__e.TailApply(PrimFunc(symconcat), symAction, V388)
return
} else {
__e.Return(V388)
return
}


}


}


}, 1)

tmp18571 := Call(__e, ns2_1set, symshen_4process_1yacc_1semantics, tmp18544)


_ = tmp18571

tmp18572 := MakeNative(func(__e *ControlFlow) {
V392 := __e.Get(1)
_ = V392
tmp18585 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V392)
}
__typedArg0 := V392
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18576 Obj

if True == tmp18585 {
tmp18583 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V392)
}
__typedArg0 := V392
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18584 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18583)
}
__typedArg0 := tmp18583
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18578 Obj

if True == tmp18584 {
tmp18580 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V392)
}
__typedArg0 := V392
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18581 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18580)
}
__typedArg0 := tmp18580
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18582 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp18581)
}
__typedArg0 := Nil
__typedArg1 := tmp18581
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18579 Obj

if True == tmp18582 {
ifres18579 = True


} else {
ifres18579 = False


}

ifres18578 = ifres18579


} else {
ifres18578 = False


}

var ifres18577 Obj

if True == ifres18578 {
ifres18577 = True


} else {
ifres18577 = False


}

ifres18576 = ifres18577


} else {
ifres18576 = False


}

if True == ifres18576 {
tmp18573 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V392)
}
__typedArg0 := V392
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18573)
}
__typedArg0 := tmp18573
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return


} else {
tmp18574 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V392)
}
__typedArg0 := V392
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18574)
}
__typedArg0 := tmp18574
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return


}


}, 1)

tmp18586 := Call(__e, ns2_1set, symshen_4_5_1out, tmp18572)


_ = tmp18586

tmp18587 := MakeNative(func(__e *ControlFlow) {
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

tmp18588 := Call(__e, ns2_1set, symshen_4in_1_6, tmp18587)


_ = tmp18588

tmp18589 := MakeNative(func(__e *ControlFlow) {
V394 := __e.Get(1)
_ = V394
tmp18590 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V394, Nil)
}
__typedArg0 := V394
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Nil, tmp18590)
}
__typedArg0 := Nil
__typedArg1 := tmp18590
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp18591 := Call(__e, ns2_1set, sym_5_b_6, tmp18589)


_ = tmp18591

tmp18592 := MakeNative(func(__e *ControlFlow) {
V395 := __e.Get(1)
_ = V395
tmp18593 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Nil, Nil)
}
__typedArg0 := Nil
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V395, tmp18593)
}
__typedArg0 := V395
__typedArg1 := tmp18593
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp18594 := Call(__e, ns2_1set, sym_5e_6, tmp18592)


_ = tmp18594

tmp18595 := MakeNative(func(__e *ControlFlow) {
V398 := __e.Get(1)
_ = V398
tmp18598 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V398)
}
__typedArg0 := Nil
__typedArg1 := V398
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp18598 {
tmp18596 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Nil, Nil)
}
__typedArg0 := Nil
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Nil, tmp18596)
}
__typedArg0 := Nil
__typedArg1 := tmp18596
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

__e.TailApply(ns2_1set, sym_5end_6, tmp18595)
return




}, 0)

