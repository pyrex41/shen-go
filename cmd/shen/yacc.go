package main

import . "github.com/pyrex41/shen-go/kl"

var YaccMain = MakeNative(func(__e *ControlFlow) {
tmp17518 := MakeNative(func(__e *ControlFlow) {
V112 := __e.Get(1)
_ = V112
V113 := __e.Get(2)
_ = V113
tmp17519 := Call(__e, V112, V113)


let__17380 := tmp17519
_ = let__17380

tmp17526 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17380)


if True == tmp17526 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("parse failure\n"))
}
__typedArg0 := MakeString("parse failure\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
} else {
tmp17524 := Call(__e, PrimFunc(symshen_4partial_1parse_1failure_2), let__17380)


if True == tmp17524 {
tmp17520 := Call(__e, PrimFunc(symshen_4in_1_6), let__17380)


tmp17521 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dresidue_d, tmp17520)
}
__typedArg0 := symshen_4_dresidue_d
__typedArg1 := tmp17520
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp17521

tmp17522 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dresidue_d)
}
__typedArg0 := symshen_4_dresidue_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4raise_1syntax_1error), tmp17522)
return


} else {
__e.TailApply(PrimFunc(symshen_4_5_1out), let__17380)
return
}


}


}, 2)

tmp17527 := Call(__e, ns2_1set, symcompile, tmp17518)


_ = tmp17527

tmp17528 := MakeNative(func(__e *ControlFlow) {
V115 := __e.Get(1)
_ = V115
tmp17529 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dmaximum_1print_1sequence_1size_d)
}
__typedArg0 := sym_dmaximum_1print_1sequence_1size_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp17530 := Call(__e, PrimFunc(symshen_4syntax_1error_1message), tmp17529, MakeInteger(0), V115)


tmp17532 := Call(__e, PrimFunc(symshen_4proc_1nl), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("syntax error here: "))
__typedS1, __typedOK1 := TypedString(tmp17530)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("syntax error here: ")
__typedArg1 := tmp17530
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(tmp17532)
}
__typedArg0 := tmp17532
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}, 1)

tmp17533 := Call(__e, ns2_1set, symshen_4raise_1syntax_1error, tmp17528)


_ = tmp17533

tmp17534 := MakeNative(func(__e *ControlFlow) {
V123 := __e.Get(1)
_ = V123
V124 := __e.Get(2)
_ = V124
V125 := __e.Get(3)
_ = V125
tmp17545 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V125)
}
__typedArg0 := Nil
__typedArg1 := V125
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17545 {
__e.Return(MakeString("\n"))
return
} else {
tmp17543 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V123, V124)
}
__typedArg0 := V123
__typedArg1 := V124
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17543 {
__e.Return(MakeString("...etc \n"))
return
} else {
tmp17541 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V125)
}
__typedArg0 := V125
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17541 {
tmp17535 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V125)
}
__typedArg0 := V125
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17536 := Call(__e, PrimFunc(symshen_4app), tmp17535, MakeString(" "), symshen_4s)


tmp17537 := (func() Obj {
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

tmp17538 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V125)
}
__typedArg0 := V125
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17539 := Call(__e, PrimFunc(symshen_4syntax_1error_1message), V123, tmp17537, tmp17538)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(tmp17536)
__typedS1, __typedOK1 := TypedString(tmp17539)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := tmp17536
__typedArg1 := tmp17539
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

tmp17546 := Call(__e, ns2_1set, symshen_4syntax_1error_1message, tmp17534)


_ = tmp17546

tmp17547 := MakeNative(func(__e *ControlFlow) {
V126 := __e.Get(1)
_ = V126
tmp17548 := Call(__e, PrimFunc(symfail))


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V126, tmp17548)
}
__typedArg0 := V126
__typedArg1 := tmp17548
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp17549 := Call(__e, ns2_1set, symshen_4parse_1failure_2, tmp17547)


_ = tmp17549

tmp17550 := MakeNative(func(__e *ControlFlow) {
V127 := __e.Get(1)
_ = V127
tmp17551 := Call(__e, PrimFunc(symshen_4in_1_6), V127)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp17551)
}
__typedArg0 := tmp17551
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})())
return


}, 1)

tmp17552 := Call(__e, ns2_1set, symshen_4partial_1parse_1failure_2, tmp17550)


_ = tmp17552

tmp17553 := MakeNative(func(__e *ControlFlow) {
V130 := __e.Get(1)
_ = V130
tmp17566 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V130)
}
__typedArg0 := V130
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17557 Obj

if True == tmp17566 {
tmp17564 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V130)
}
__typedArg0 := V130
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17565 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp17564)
}
__typedArg0 := tmp17564
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17559 Obj

if True == tmp17565 {
tmp17561 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V130)
}
__typedArg0 := V130
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17562 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp17561)
}
__typedArg0 := tmp17561
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17563 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp17562)
}
__typedArg0 := Nil
__typedArg1 := tmp17562
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17560 Obj

if True == tmp17563 {
ifres17560 = True


} else {
ifres17560 = False


}

ifres17559 = ifres17560


} else {
ifres17559 = False


}

var ifres17558 Obj

if True == ifres17559 {
ifres17558 = True


} else {
ifres17558 = False


}

ifres17557 = ifres17558


} else {
ifres17557 = False


}

if True == ifres17557 {
tmp17554 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V130)
}
__typedArg0 := V130
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp17554)
}
__typedArg0 := tmp17554
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return


} else {
tmp17555 := Call(__e, PrimFunc(symshen_4app), V130, MakeString(" is not a YACC stream\n"), symshen_4s)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(tmp17555)
}
__typedArg0 := tmp17555
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}


}, 1)

tmp17567 := Call(__e, ns2_1set, symshen_4objectcode, tmp17553)


_ = tmp17567

tmp17568 := MakeNative(func(__e *ControlFlow) {
V131 := __e.Get(1)
_ = V131
tmp17569 := MakeNative(func(__e *ControlFlow) {
Z132 := __e.Get(1)
_ = Z132
__e.TailApply(PrimFunc(symshen_4_5yacc_6), Z132)
return
}, 1)

__e.TailApply(PrimFunc(symcompile), tmp17569, V131)
return


}, 1)

tmp17570 := Call(__e, ns2_1set, symshen_4yacc_1_6shen, tmp17568)


_ = tmp17570

tmp17571 := MakeNative(func(__e *ControlFlow) {
V133 := __e.Get(1)
_ = V133
tmp17598 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V133)
}
__typedArg0 := V133
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17572 Obj

if True == tmp17598 {
tmp17573 := Call(__e, PrimFunc(symhead), V133)


let__17382 := tmp17573
_ = let__17382

tmp17574 := Call(__e, PrimFunc(symtail), V133)


let__17383 := tmp17574
_ = let__17383

tmp17575 := Call(__e, PrimFunc(symshen_4_5yaccsig_6), let__17383)


let__17384 := tmp17575
_ = let__17384

tmp17596 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17384)


var ifres17576 Obj

if True == tmp17596 {
tmp17577 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17576 = tmp17577


} else {
tmp17578 := Call(__e, PrimFunc(symshen_4_5_1out), let__17384)


let__17385 := tmp17578
_ = let__17385

tmp17579 := Call(__e, PrimFunc(symshen_4in_1_6), let__17384)


let__17386 := tmp17579
_ = let__17386

tmp17580 := Call(__e, PrimFunc(symshen_4_5c_1rules_6), let__17386)


let__17387 := tmp17580
_ = let__17387

tmp17595 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17387)


var ifres17581 Obj

if True == tmp17595 {
tmp17582 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17581 = tmp17582


} else {
tmp17583 := Call(__e, PrimFunc(symshen_4_5_1out), let__17387)


let__17388 := tmp17583
_ = let__17388

tmp17584 := Call(__e, PrimFunc(symshen_4in_1_6), let__17387)


let__17389 := tmp17584
_ = let__17389

tmp17585 := Call(__e, PrimFunc(symgensym), symS)


let__17390 := tmp17585
_ = let__17390

tmp17586 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17382, Nil)
}
__typedArg0 := let__17382
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17587 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefine, tmp17586)
}
__typedArg0 := symdefine
__typedArg1 := tmp17586
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17588 := Call(__e, PrimFunc(symshen_4c_1rules_1_6shen), let__17385, let__17390, let__17388)


tmp17589 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17588, Nil)
}
__typedArg0 := tmp17588
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17590 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_6, tmp17589)
}
__typedArg0 := sym_1_6
__typedArg1 := tmp17589
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17591 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17390, tmp17590)
}
__typedArg0 := let__17390
__typedArg1 := tmp17590
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17592 := Call(__e, PrimFunc(symappend), let__17385, tmp17591)


tmp17593 := Call(__e, PrimFunc(symappend), tmp17587, tmp17592)


let__17391 := tmp17593
_ = let__17391

tmp17594 := Call(__e, PrimFunc(symshen_4comb), let__17389, let__17391)


ifres17581 = tmp17594


}

ifres17576 = ifres17581


}

ifres17572 = ifres17576


} else {
tmp17597 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17572 = tmp17597


}

let__17381 := ifres17572
_ = let__17381

tmp17600 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17381)


if True == tmp17600 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__17381)
return
}


}, 1)

tmp17601 := Call(__e, ns2_1set, symshen_4_5yacc_6, tmp17571)


_ = tmp17601

tmp17602 := MakeNative(func(__e *ControlFlow) {
V145 := __e.Get(1)
_ = V145
tmp17659 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V145)
}
__typedArg0 := V145
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17603 Obj

if True == tmp17659 {
tmp17604 := Call(__e, PrimFunc(symhead), V145)


let__17393 := tmp17604
_ = let__17393

tmp17605 := Call(__e, PrimFunc(symtail), V145)


let__17394 := tmp17605
_ = let__17394

tmp17657 := Call(__e, PrimFunc(symshen_4ccons_2), let__17394)


var ifres17606 Obj

if True == tmp17657 {
tmp17607 := Call(__e, PrimFunc(symhead), let__17394)


let__17395 := tmp17607
_ = let__17395

tmp17608 := Call(__e, PrimFunc(symtail), let__17394)


let__17396 := tmp17608
_ = let__17396

tmp17655 := Call(__e, PrimFunc(symshen_4hds_a_2), let__17395, symlist)


var ifres17609 Obj

if True == tmp17655 {
tmp17610 := Call(__e, PrimFunc(symtail), let__17395)


let__17397 := tmp17610
_ = let__17397

tmp17653 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__17397)
}
__typedArg0 := let__17397
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17611 Obj

if True == tmp17653 {
tmp17612 := Call(__e, PrimFunc(symhead), let__17397)


let__17398 := tmp17612
_ = let__17398

tmp17613 := Call(__e, PrimFunc(symtail), let__17397)


let__17399 := tmp17613
_ = let__17399

tmp17614 := Call(__e, PrimFunc(sym_5end_6), let__17399)


let__17400 := tmp17614
_ = let__17400

tmp17651 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17400)


var ifres17615 Obj

if True == tmp17651 {
tmp17616 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17615 = tmp17616


} else {
tmp17617 := Call(__e, PrimFunc(symshen_4in_1_6), let__17400)


let__17401 := tmp17617
_ = let__17401

tmp17650 := Call(__e, PrimFunc(symshen_4hds_a_2), let__17396, sym_a_a_6)


var ifres17618 Obj

if True == tmp17650 {
tmp17619 := Call(__e, PrimFunc(symtail), let__17396)


let__17402 := tmp17619
_ = let__17402

tmp17648 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__17402)
}
__typedArg0 := let__17402
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17620 Obj

if True == tmp17648 {
tmp17621 := Call(__e, PrimFunc(symhead), let__17402)


let__17403 := tmp17621
_ = let__17403

tmp17622 := Call(__e, PrimFunc(symtail), let__17402)


let__17404 := tmp17622
_ = let__17404

tmp17646 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__17404)
}
__typedArg0 := let__17404
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17623 Obj

if True == tmp17646 {
tmp17624 := Call(__e, PrimFunc(symhead), let__17404)


let__17405 := tmp17624
_ = let__17405

tmp17625 := Call(__e, PrimFunc(symtail), let__17404)


let__17406 := tmp17625
_ = let__17406

tmp17644 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_i, let__17393)
}
__typedArg0 := sym_i
__typedArg1 := let__17393
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17641 Obj

if True == tmp17644 {
tmp17643 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_j, let__17405)
}
__typedArg0 := sym_j
__typedArg1 := let__17405
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17642 Obj

if True == tmp17643 {
ifres17642 = True


} else {
ifres17642 = False


}

ifres17641 = ifres17642


} else {
ifres17641 = False


}

var ifres17626 Obj

if True == ifres17641 {
tmp17627 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17398, Nil)
}
__typedArg0 := let__17398
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17628 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp17627)
}
__typedArg0 := symlist
__typedArg1 := tmp17627
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17629 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17398, Nil)
}
__typedArg0 := let__17398
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17630 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp17629)
}
__typedArg0 := symlist
__typedArg1 := tmp17629
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17631 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17403, Nil)
}
__typedArg0 := let__17403
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17632 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17630, tmp17631)
}
__typedArg0 := tmp17630
__typedArg1 := tmp17631
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17633 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstr, tmp17632)
}
__typedArg0 := symstr
__typedArg1 := tmp17632
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17634 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_j, Nil)
}
__typedArg0 := sym_j
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17635 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17633, tmp17634)
}
__typedArg0 := tmp17633
__typedArg1 := tmp17634
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17636 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp17635)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp17635
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17637 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17628, tmp17636)
}
__typedArg0 := tmp17628
__typedArg1 := tmp17636
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17638 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_i, tmp17637)
}
__typedArg0 := sym_i
__typedArg1 := tmp17637
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17639 := Call(__e, PrimFunc(symshen_4comb), let__17406, tmp17638)


ifres17626 = tmp17639


} else {
tmp17640 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17626 = tmp17640


}

ifres17623 = ifres17626


} else {
tmp17645 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17623 = tmp17645


}

ifres17620 = ifres17623


} else {
tmp17647 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17620 = tmp17647


}

ifres17618 = ifres17620


} else {
tmp17649 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17618 = tmp17649


}

ifres17615 = ifres17618


}

ifres17611 = ifres17615


} else {
tmp17652 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17611 = tmp17652


}

ifres17609 = ifres17611


} else {
tmp17654 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17609 = tmp17654


}

ifres17606 = ifres17609


} else {
tmp17656 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17606 = tmp17656


}

ifres17603 = ifres17606


} else {
tmp17658 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17603 = tmp17658


}

let__17392 := ifres17603
_ = let__17392

tmp17669 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17392)


if True == tmp17669 {
tmp17660 := Call(__e, PrimFunc(sym_5e_6), V145)


let__17408 := tmp17660
_ = let__17408

tmp17665 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17408)


var ifres17661 Obj

if True == tmp17665 {
tmp17662 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17661 = tmp17662


} else {
tmp17663 := Call(__e, PrimFunc(symshen_4in_1_6), let__17408)


let__17409 := tmp17663
_ = let__17409

tmp17664 := Call(__e, PrimFunc(symshen_4comb), let__17409, Nil)


ifres17661 = tmp17664


}

let__17407 := ifres17661
_ = let__17407

tmp17667 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17407)


if True == tmp17667 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__17407)
return
}


} else {
__e.Return(let__17392)
return
}


}, 1)

tmp17670 := Call(__e, ns2_1set, symshen_4_5yaccsig_6, tmp17602)


_ = tmp17670

tmp17671 := MakeNative(func(__e *ControlFlow) {
V164 := __e.Get(1)
_ = V164
tmp17672 := Call(__e, PrimFunc(symshen_4_5c_1rule_6), V164)


let__17411 := tmp17672
_ = let__17411

tmp17685 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17411)


var ifres17673 Obj

if True == tmp17685 {
tmp17674 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17673 = tmp17674


} else {
tmp17675 := Call(__e, PrimFunc(symshen_4_5_1out), let__17411)


let__17412 := tmp17675
_ = let__17412

tmp17676 := Call(__e, PrimFunc(symshen_4in_1_6), let__17411)


let__17413 := tmp17676
_ = let__17413

tmp17677 := Call(__e, PrimFunc(symshen_4_5c_1rules_6), let__17413)


let__17414 := tmp17677
_ = let__17414

tmp17684 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17414)


var ifres17678 Obj

if True == tmp17684 {
tmp17679 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17678 = tmp17679


} else {
tmp17680 := Call(__e, PrimFunc(symshen_4_5_1out), let__17414)


let__17415 := tmp17680
_ = let__17415

tmp17681 := Call(__e, PrimFunc(symshen_4in_1_6), let__17414)


let__17416 := tmp17681
_ = let__17416

tmp17682 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17412, let__17415)
}
__typedArg0 := let__17412
__typedArg1 := let__17415
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17683 := Call(__e, PrimFunc(symshen_4comb), let__17416, tmp17682)


ifres17678 = tmp17683


}

ifres17673 = ifres17678


}

let__17410 := ifres17673
_ = let__17410

tmp17701 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17410)


if True == tmp17701 {
tmp17686 := Call(__e, PrimFunc(sym_5_b_6), V164)


let__17418 := tmp17686
_ = let__17418

tmp17697 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17418)


var ifres17687 Obj

if True == tmp17697 {
tmp17688 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17687 = tmp17688


} else {
tmp17689 := Call(__e, PrimFunc(symshen_4_5_1out), let__17418)


let__17419 := tmp17689
_ = let__17419

tmp17690 := Call(__e, PrimFunc(symshen_4in_1_6), let__17418)


let__17420 := tmp17690
_ = let__17420

tmp17695 := Call(__e, PrimFunc(symempty_2), let__17419)


var ifres17691 Obj

if True == tmp17695 {
ifres17691 = Nil


} else {
tmp17692 := Call(__e, PrimFunc(symshen_4app), let__17419, MakeString("\n ..."), symshen_4r)


tmp17694 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("YACC syntax error here:\n "))
__typedS1, __typedOK1 := TypedString(tmp17692)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("YACC syntax error here:\n ")
__typedArg1 := tmp17692
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("YACC syntax error here:\n "))
__typedS1, __typedOK1 := TypedString(tmp17692)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("YACC syntax error here:\n ")
__typedArg1 := tmp17692
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})()

ifres17691 = tmp17694


}

tmp17696 := Call(__e, PrimFunc(symshen_4comb), let__17420, ifres17691)


ifres17687 = tmp17696


}

let__17417 := ifres17687
_ = let__17417

tmp17699 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17417)


if True == tmp17699 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__17417)
return
}


} else {
__e.Return(let__17410)
return
}


}, 1)

tmp17702 := Call(__e, ns2_1set, symshen_4_5c_1rules_6, tmp17671)


_ = tmp17702

tmp17703 := MakeNative(func(__e *ControlFlow) {
V176 := __e.Get(1)
_ = V176
tmp17704 := Call(__e, PrimFunc(symshen_4_5syntax_6), V176)


let__17422 := tmp17704
_ = let__17422

tmp17723 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17422)


var ifres17705 Obj

if True == tmp17723 {
tmp17706 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17705 = tmp17706


} else {
tmp17707 := Call(__e, PrimFunc(symshen_4_5_1out), let__17422)


let__17423 := tmp17707
_ = let__17423

tmp17708 := Call(__e, PrimFunc(symshen_4in_1_6), let__17422)


let__17424 := tmp17708
_ = let__17424

tmp17709 := Call(__e, PrimFunc(symshen_4_5semantics_6), let__17424)


let__17425 := tmp17709
_ = let__17425

tmp17722 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17425)


var ifres17710 Obj

if True == tmp17722 {
tmp17711 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17710 = tmp17711


} else {
tmp17712 := Call(__e, PrimFunc(symshen_4_5_1out), let__17425)


let__17426 := tmp17712
_ = let__17426

tmp17713 := Call(__e, PrimFunc(symshen_4in_1_6), let__17425)


let__17427 := tmp17713
_ = let__17427

tmp17714 := Call(__e, PrimFunc(symshen_4_5sc_6), let__17427)


let__17428 := tmp17714
_ = let__17428

tmp17721 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17428)


var ifres17715 Obj

if True == tmp17721 {
tmp17716 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17715 = tmp17716


} else {
tmp17717 := Call(__e, PrimFunc(symshen_4in_1_6), let__17428)


let__17429 := tmp17717
_ = let__17429

tmp17718 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17426, Nil)
}
__typedArg0 := let__17426
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17719 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17423, tmp17718)
}
__typedArg0 := let__17423
__typedArg1 := tmp17718
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17720 := Call(__e, PrimFunc(symshen_4comb), let__17429, tmp17719)


ifres17715 = tmp17720


}

ifres17710 = ifres17715


}

ifres17705 = ifres17710


}

let__17421 := ifres17705
_ = let__17421

tmp17742 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17421)


if True == tmp17742 {
tmp17724 := Call(__e, PrimFunc(symshen_4_5syntax_6), V176)


let__17431 := tmp17724
_ = let__17431

tmp17738 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17431)


var ifres17725 Obj

if True == tmp17738 {
tmp17726 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17725 = tmp17726


} else {
tmp17727 := Call(__e, PrimFunc(symshen_4_5_1out), let__17431)


let__17432 := tmp17727
_ = let__17432

tmp17728 := Call(__e, PrimFunc(symshen_4in_1_6), let__17431)


let__17433 := tmp17728
_ = let__17433

tmp17729 := Call(__e, PrimFunc(symshen_4_5sc_6), let__17433)


let__17434 := tmp17729
_ = let__17434

tmp17737 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17434)


var ifres17730 Obj

if True == tmp17737 {
tmp17731 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17730 = tmp17731


} else {
tmp17732 := Call(__e, PrimFunc(symshen_4in_1_6), let__17434)


let__17435 := tmp17732
_ = let__17435

tmp17733 := Call(__e, PrimFunc(symshen_4autocomplete), let__17432)


tmp17734 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17733, Nil)
}
__typedArg0 := tmp17733
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17735 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17432, tmp17734)
}
__typedArg0 := let__17432
__typedArg1 := tmp17734
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17736 := Call(__e, PrimFunc(symshen_4comb), let__17435, tmp17735)


ifres17730 = tmp17736


}

ifres17725 = ifres17730


}

let__17430 := ifres17725
_ = let__17430

tmp17740 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17430)


if True == tmp17740 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__17430)
return
}


} else {
__e.Return(let__17421)
return
}


}, 1)

tmp17743 := Call(__e, ns2_1set, symshen_4_5c_1rule_6, tmp17703)


_ = tmp17743

tmp17744 := MakeNative(func(__e *ControlFlow) {
V192 := __e.Get(1)
_ = V192
tmp17773 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V192)
}
__typedArg0 := V192
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17765 Obj

if True == tmp17773 {
tmp17771 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V192)
}
__typedArg0 := V192
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17772 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp17771)
}
__typedArg0 := Nil
__typedArg1 := tmp17771
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17767 Obj

if True == tmp17772 {
tmp17769 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V192)
}
__typedArg0 := V192
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17770 := Call(__e, PrimFunc(symshen_4non_1terminal_2), tmp17769)


var ifres17768 Obj

if True == tmp17770 {
ifres17768 = True


} else {
ifres17768 = False


}

ifres17767 = ifres17768


} else {
ifres17767 = False


}

var ifres17766 Obj

if True == ifres17767 {
ifres17766 = True


} else {
ifres17766 = False


}

ifres17765 = ifres17766


} else {
ifres17765 = False


}

if True == ifres17765 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V192)
}
__typedArg0 := V192
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return
} else {
tmp17763 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V192)
}
__typedArg0 := V192
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17759 Obj

if True == tmp17763 {
tmp17761 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V192)
}
__typedArg0 := V192
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17762 := Call(__e, PrimFunc(symshen_4non_1terminal_2), tmp17761)


var ifres17760 Obj

if True == tmp17762 {
ifres17760 = True


} else {
ifres17760 = False


}

ifres17759 = ifres17760


} else {
ifres17759 = False


}

if True == ifres17759 {
tmp17745 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V192)
}
__typedArg0 := V192
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17746 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V192)
}
__typedArg0 := V192
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17747 := Call(__e, PrimFunc(symshen_4autocomplete), tmp17746)


tmp17748 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17747, Nil)
}
__typedArg0 := tmp17747
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17749 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17745, tmp17748)
}
__typedArg0 := tmp17745
__typedArg1 := tmp17748
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symappend, tmp17749)
}
__typedArg0 := symappend
__typedArg1 := tmp17749
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp17757 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V192)
}
__typedArg0 := V192
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17757 {
tmp17750 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V192)
}
__typedArg0 := V192
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17751 := Call(__e, PrimFunc(symshen_4autocomplete), tmp17750)


tmp17752 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V192)
}
__typedArg0 := V192
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17753 := Call(__e, PrimFunc(symshen_4autocomplete), tmp17752)


tmp17754 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17753, Nil)
}
__typedArg0 := tmp17753
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17755 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17751, tmp17754)
}
__typedArg0 := tmp17751
__typedArg1 := tmp17754
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp17755)
}
__typedArg0 := symcons
__typedArg1 := tmp17755
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

tmp17774 := Call(__e, ns2_1set, symshen_4autocomplete, tmp17744)


_ = tmp17774

tmp17775 := MakeNative(func(__e *ControlFlow) {
V193 := __e.Get(1)
_ = V193
tmp17781 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsymbol_2) {
return PrimIsSymbol(V193)
}
__typedArg0 := V193
return Call(__e, PrimFunc(symsymbol_2), __typedArg0)
})()

if True == tmp17781 {
tmp17777 := Call(__e, PrimFunc(symexplode), V193)


let__17436 := tmp17777
_ = let__17436

tmp17778 := MakeNative(func(__e *ControlFlow) {
Z195 := __e.Get(1)
_ = Z195
__e.TailApply(PrimFunc(symshen_4_5non_1terminal_2_6), Z195)
return
}, 1)

tmp17779 := Call(__e, PrimFunc(symcompile), tmp17778, let__17436)


if True == tmp17779 {
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

tmp17782 := Call(__e, ns2_1set, symshen_4non_1terminal_2, tmp17775)


_ = tmp17782

tmp17783 := MakeNative(func(__e *ControlFlow) {
V196 := __e.Get(1)
_ = V196
tmp17784 := Call(__e, PrimFunc(symshen_4_5packagenames_6), V196)


let__17438 := tmp17784
_ = let__17438

tmp17794 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17438)


var ifres17785 Obj

if True == tmp17794 {
tmp17786 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17785 = tmp17786


} else {
tmp17787 := Call(__e, PrimFunc(symshen_4in_1_6), let__17438)


let__17439 := tmp17787
_ = let__17439

tmp17788 := Call(__e, PrimFunc(symshen_4_5non_1terminal_1name_6), let__17439)


let__17440 := tmp17788
_ = let__17440

tmp17793 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17440)


var ifres17789 Obj

if True == tmp17793 {
tmp17790 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17789 = tmp17790


} else {
tmp17791 := Call(__e, PrimFunc(symshen_4in_1_6), let__17440)


let__17441 := tmp17791
_ = let__17441

tmp17792 := Call(__e, PrimFunc(symshen_4comb), let__17441, True)


ifres17789 = tmp17792


}

ifres17785 = ifres17789


}

let__17437 := ifres17785
_ = let__17437

tmp17812 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17437)


if True == tmp17812 {
tmp17795 := Call(__e, PrimFunc(symshen_4_5non_1terminal_1name_6), V196)


let__17443 := tmp17795
_ = let__17443

tmp17800 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17443)


var ifres17796 Obj

if True == tmp17800 {
tmp17797 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17796 = tmp17797


} else {
tmp17798 := Call(__e, PrimFunc(symshen_4in_1_6), let__17443)


let__17444 := tmp17798
_ = let__17444

tmp17799 := Call(__e, PrimFunc(symshen_4comb), let__17444, True)


ifres17796 = tmp17799


}

let__17442 := ifres17796
_ = let__17442

tmp17810 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17442)


if True == tmp17810 {
tmp17801 := Call(__e, PrimFunc(sym_5_b_6), V196)


let__17446 := tmp17801
_ = let__17446

tmp17806 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17446)


var ifres17802 Obj

if True == tmp17806 {
tmp17803 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17802 = tmp17803


} else {
tmp17804 := Call(__e, PrimFunc(symshen_4in_1_6), let__17446)


let__17447 := tmp17804
_ = let__17447

tmp17805 := Call(__e, PrimFunc(symshen_4comb), let__17447, False)


ifres17802 = tmp17805


}

let__17445 := ifres17802
_ = let__17445

tmp17808 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17445)


if True == tmp17808 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__17445)
return
}


} else {
__e.Return(let__17442)
return
}


} else {
__e.Return(let__17437)
return
}


}, 1)

tmp17813 := Call(__e, ns2_1set, symshen_4_5non_1terminal_2_6, tmp17783)


_ = tmp17813

tmp17814 := MakeNative(func(__e *ControlFlow) {
V208 := __e.Get(1)
_ = V208
tmp17815 := Call(__e, PrimFunc(symshen_4_5packagename_6), V208)


let__17449 := tmp17815
_ = let__17449

tmp17829 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17449)


var ifres17816 Obj

if True == tmp17829 {
tmp17817 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17816 = tmp17817


} else {
tmp17818 := Call(__e, PrimFunc(symshen_4in_1_6), let__17449)


let__17450 := tmp17818
_ = let__17450

tmp17828 := Call(__e, PrimFunc(symshen_4hds_a_2), let__17450, MakeString("."))


var ifres17819 Obj

if True == tmp17828 {
tmp17820 := Call(__e, PrimFunc(symtail), let__17450)


let__17451 := tmp17820
_ = let__17451

tmp17821 := Call(__e, PrimFunc(symshen_4_5packagenames_6), let__17451)


let__17452 := tmp17821
_ = let__17452

tmp17826 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17452)


var ifres17822 Obj

if True == tmp17826 {
tmp17823 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17822 = tmp17823


} else {
tmp17824 := Call(__e, PrimFunc(symshen_4in_1_6), let__17452)


let__17453 := tmp17824
_ = let__17453

tmp17825 := Call(__e, PrimFunc(symshen_4comb), let__17453, symshen_4skip)


ifres17822 = tmp17825


}

ifres17819 = ifres17822


} else {
tmp17827 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17819 = tmp17827


}

ifres17816 = ifres17819


}

let__17448 := ifres17816
_ = let__17448

tmp17843 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17448)


if True == tmp17843 {
tmp17830 := Call(__e, PrimFunc(symshen_4_5packagename_6), V208)


let__17455 := tmp17830
_ = let__17455

tmp17839 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17455)


var ifres17831 Obj

if True == tmp17839 {
tmp17832 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17831 = tmp17832


} else {
tmp17833 := Call(__e, PrimFunc(symshen_4in_1_6), let__17455)


let__17456 := tmp17833
_ = let__17456

tmp17838 := Call(__e, PrimFunc(symshen_4hds_a_2), let__17456, MakeString("."))


var ifres17834 Obj

if True == tmp17838 {
tmp17835 := Call(__e, PrimFunc(symtail), let__17456)


let__17457 := tmp17835
_ = let__17457

tmp17836 := Call(__e, PrimFunc(symshen_4comb), let__17457, symshen_4skip)


ifres17834 = tmp17836


} else {
tmp17837 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17834 = tmp17837


}

ifres17831 = ifres17834


}

let__17454 := ifres17831
_ = let__17454

tmp17841 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17454)


if True == tmp17841 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__17454)
return
}


} else {
__e.Return(let__17448)
return
}


}, 1)

tmp17844 := Call(__e, ns2_1set, symshen_4_5packagenames_6, tmp17814)


_ = tmp17844

tmp17845 := MakeNative(func(__e *ControlFlow) {
V219 := __e.Get(1)
_ = V219
tmp17846 := Call(__e, PrimFunc(symshen_4_5packagechar_6), V219)


let__17459 := tmp17846
_ = let__17459

tmp17856 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17459)


var ifres17847 Obj

if True == tmp17856 {
tmp17848 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17847 = tmp17848


} else {
tmp17849 := Call(__e, PrimFunc(symshen_4in_1_6), let__17459)


let__17460 := tmp17849
_ = let__17460

tmp17850 := Call(__e, PrimFunc(symshen_4_5packagename_6), let__17460)


let__17461 := tmp17850
_ = let__17461

tmp17855 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17461)


var ifres17851 Obj

if True == tmp17855 {
tmp17852 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17851 = tmp17852


} else {
tmp17853 := Call(__e, PrimFunc(symshen_4in_1_6), let__17461)


let__17462 := tmp17853
_ = let__17462

tmp17854 := Call(__e, PrimFunc(symshen_4comb), let__17462, symshen_4skip)


ifres17851 = tmp17854


}

ifres17847 = ifres17851


}

let__17458 := ifres17847
_ = let__17458

tmp17866 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17458)


if True == tmp17866 {
tmp17857 := Call(__e, PrimFunc(sym_5e_6), V219)


let__17464 := tmp17857
_ = let__17464

tmp17862 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17464)


var ifres17858 Obj

if True == tmp17862 {
tmp17859 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17858 = tmp17859


} else {
tmp17860 := Call(__e, PrimFunc(symshen_4in_1_6), let__17464)


let__17465 := tmp17860
_ = let__17465

tmp17861 := Call(__e, PrimFunc(symshen_4comb), let__17465, symshen_4skip)


ifres17858 = tmp17861


}

let__17463 := ifres17858
_ = let__17463

tmp17864 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17463)


if True == tmp17864 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__17463)
return
}


} else {
__e.Return(let__17458)
return
}


}, 1)

tmp17867 := Call(__e, ns2_1set, symshen_4_5packagename_6, tmp17845)


_ = tmp17867

tmp17868 := MakeNative(func(__e *ControlFlow) {
V228 := __e.Get(1)
_ = V228
tmp17878 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V228)
}
__typedArg0 := V228
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17869 Obj

if True == tmp17878 {
tmp17870 := Call(__e, PrimFunc(symhead), V228)


let__17467 := tmp17870
_ = let__17467

tmp17871 := Call(__e, PrimFunc(symtail), V228)


let__17468 := tmp17871
_ = let__17468

tmp17875 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__17467, MakeString("."))
}
__typedArg0 := let__17467
__typedArg1 := MakeString(".")
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

tmp17876 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp17875)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp17875
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres17872 Obj

if True == tmp17876 {
tmp17873 := Call(__e, PrimFunc(symshen_4comb), let__17468, symshen_4skip)


ifres17872 = tmp17873


} else {
tmp17874 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17872 = tmp17874


}

ifres17869 = ifres17872


} else {
tmp17877 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17869 = tmp17877


}

let__17466 := ifres17869
_ = let__17466

tmp17880 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17466)


if True == tmp17880 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__17466)
return
}


}, 1)

tmp17881 := Call(__e, ns2_1set, symshen_4_5packagechar_6, tmp17868)


_ = tmp17881

tmp17882 := MakeNative(func(__e *ControlFlow) {
V232 := __e.Get(1)
_ = V232
tmp17901 := Call(__e, PrimFunc(symshen_4hds_a_2), V232, MakeString("<"))


var ifres17883 Obj

if True == tmp17901 {
tmp17884 := Call(__e, PrimFunc(symtail), V232)


let__17470 := tmp17884
_ = let__17470

tmp17885 := Call(__e, PrimFunc(sym_5_b_6), let__17470)


let__17471 := tmp17885
_ = let__17471

tmp17899 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17471)


var ifres17886 Obj

if True == tmp17899 {
tmp17887 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17886 = tmp17887


} else {
tmp17888 := Call(__e, PrimFunc(symshen_4_5_1out), let__17471)


let__17472 := tmp17888
_ = let__17472

tmp17889 := Call(__e, PrimFunc(symshen_4in_1_6), let__17471)


let__17473 := tmp17889
_ = let__17473

tmp17893 := Call(__e, PrimFunc(symreverse), let__17472)


let__17474 := tmp17893
_ = let__17474

tmp17898 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__17474)
}
__typedArg0 := let__17474
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17894 Obj

if True == tmp17898 {
tmp17896 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__17474)
}
__typedArg0 := let__17474
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17897 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp17896, MakeString(">"))
}
__typedArg0 := tmp17896
__typedArg1 := MakeString(">")
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17895 Obj

if True == tmp17897 {
ifres17895 = True


} else {
ifres17895 = False


}

ifres17894 = ifres17895


} else {
ifres17894 = False


}

var ifres17890 Obj

if True == ifres17894 {
tmp17891 := Call(__e, PrimFunc(symshen_4comb), let__17473, symshen_4skip)


ifres17890 = tmp17891


} else {
tmp17892 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17890 = tmp17892


}

ifres17886 = ifres17890


}

ifres17883 = ifres17886


} else {
tmp17900 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17883 = tmp17900


}

let__17469 := ifres17883
_ = let__17469

tmp17903 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17469)


if True == tmp17903 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__17469)
return
}


}, 1)

tmp17904 := Call(__e, ns2_1set, symshen_4_5non_1terminal_1name_6, tmp17882)


_ = tmp17904

tmp17905 := MakeNative(func(__e *ControlFlow) {
V239 := __e.Get(1)
_ = V239
tmp17906 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(";"))
}
__typedArg0 := MakeString(";")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V239, tmp17906)
}
__typedArg0 := V239
__typedArg1 := tmp17906
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp17907 := Call(__e, ns2_1set, symshen_4semicolon_2, tmp17905)


_ = tmp17907

tmp17908 := MakeNative(func(__e *ControlFlow) {
V240 := __e.Get(1)
_ = V240
tmp17917 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V240)
}
__typedArg0 := V240
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17909 Obj

if True == tmp17917 {
tmp17910 := Call(__e, PrimFunc(symhead), V240)


let__17476 := tmp17910
_ = let__17476

tmp17911 := Call(__e, PrimFunc(symtail), V240)


let__17477 := tmp17911
_ = let__17477

tmp17915 := Call(__e, PrimFunc(symshen_4colon_1equal_2), let__17476)


var ifres17912 Obj

if True == tmp17915 {
tmp17913 := Call(__e, PrimFunc(symshen_4comb), let__17477, symshen_4skip)


ifres17912 = tmp17913


} else {
tmp17914 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17912 = tmp17914


}

ifres17909 = ifres17912


} else {
tmp17916 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17909 = tmp17916


}

let__17475 := ifres17909
_ = let__17475

tmp17919 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17475)


if True == tmp17919 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__17475)
return
}


}, 1)

tmp17920 := Call(__e, ns2_1set, symshen_4_5colon_1equal_6, tmp17908)


_ = tmp17920

tmp17921 := MakeNative(func(__e *ControlFlow) {
V244 := __e.Get(1)
_ = V244
tmp17922 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":="))
}
__typedArg0 := MakeString(":=")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp17922, V244)
}
__typedArg0 := tmp17922
__typedArg1 := V244
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp17923 := Call(__e, ns2_1set, symshen_4colon_1equal_2, tmp17921)


_ = tmp17923

tmp17924 := MakeNative(func(__e *ControlFlow) {
V245 := __e.Get(1)
_ = V245
tmp17925 := Call(__e, PrimFunc(symshen_4_5syntax_1item_6), V245)


let__17479 := tmp17925
_ = let__17479

tmp17938 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17479)


var ifres17926 Obj

if True == tmp17938 {
tmp17927 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17926 = tmp17927


} else {
tmp17928 := Call(__e, PrimFunc(symshen_4_5_1out), let__17479)


let__17480 := tmp17928
_ = let__17480

tmp17929 := Call(__e, PrimFunc(symshen_4in_1_6), let__17479)


let__17481 := tmp17929
_ = let__17481

tmp17930 := Call(__e, PrimFunc(symshen_4_5syntax_6), let__17481)


let__17482 := tmp17930
_ = let__17482

tmp17937 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17482)


var ifres17931 Obj

if True == tmp17937 {
tmp17932 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17931 = tmp17932


} else {
tmp17933 := Call(__e, PrimFunc(symshen_4_5_1out), let__17482)


let__17483 := tmp17933
_ = let__17483

tmp17934 := Call(__e, PrimFunc(symshen_4in_1_6), let__17482)


let__17484 := tmp17934
_ = let__17484

tmp17935 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17480, let__17483)
}
__typedArg0 := let__17480
__typedArg1 := let__17483
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17936 := Call(__e, PrimFunc(symshen_4comb), let__17484, tmp17935)


ifres17931 = tmp17936


}

ifres17926 = ifres17931


}

let__17478 := ifres17926
_ = let__17478

tmp17950 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17478)


if True == tmp17950 {
tmp17939 := Call(__e, PrimFunc(symshen_4_5syntax_1item_6), V245)


let__17486 := tmp17939
_ = let__17486

tmp17946 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17486)


var ifres17940 Obj

if True == tmp17946 {
tmp17941 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17940 = tmp17941


} else {
tmp17942 := Call(__e, PrimFunc(symshen_4_5_1out), let__17486)


let__17487 := tmp17942
_ = let__17487

tmp17943 := Call(__e, PrimFunc(symshen_4in_1_6), let__17486)


let__17488 := tmp17943
_ = let__17488

tmp17944 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17487, Nil)
}
__typedArg0 := let__17487
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17945 := Call(__e, PrimFunc(symshen_4comb), let__17488, tmp17944)


ifres17940 = tmp17945


}

let__17485 := ifres17940
_ = let__17485

tmp17948 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17485)


if True == tmp17948 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__17485)
return
}


} else {
__e.Return(let__17478)
return
}


}, 1)

tmp17951 := Call(__e, ns2_1set, symshen_4_5syntax_6, tmp17924)


_ = tmp17951

tmp17952 := MakeNative(func(__e *ControlFlow) {
V257 := __e.Get(1)
_ = V257
tmp17961 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V257)
}
__typedArg0 := V257
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17953 Obj

if True == tmp17961 {
tmp17954 := Call(__e, PrimFunc(symhead), V257)


let__17490 := tmp17954
_ = let__17490

tmp17955 := Call(__e, PrimFunc(symtail), V257)


let__17491 := tmp17955
_ = let__17491

tmp17959 := Call(__e, PrimFunc(symshen_4syntax_1item_2), let__17490)


var ifres17956 Obj

if True == tmp17959 {
tmp17957 := Call(__e, PrimFunc(symshen_4comb), let__17491, let__17490)


ifres17956 = tmp17957


} else {
tmp17958 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17956 = tmp17958


}

ifres17953 = ifres17956


} else {
tmp17960 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17953 = tmp17960


}

let__17489 := ifres17953
_ = let__17489

tmp17963 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17489)


if True == tmp17963 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__17489)
return
}


}, 1)

tmp17964 := Call(__e, ns2_1set, symshen_4_5syntax_1item_6, tmp17952)


_ = tmp17964

tmp17965 := MakeNative(func(__e *ControlFlow) {
V263 := __e.Get(1)
_ = V263
tmp18001 := Call(__e, PrimFunc(symshen_4colon_1equal_2), V263)


if True == tmp18001 {
__e.Return(False)
return
} else {
tmp17999 := Call(__e, PrimFunc(symshen_4semicolon_2), V263)


if True == tmp17999 {
__e.Return(False)
return
} else {
tmp17997 := Call(__e, PrimFunc(symatom_2), V263)


if True == tmp17997 {
__e.Return(True)
return
} else {
tmp17995 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V263)
}
__typedArg0 := V263
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17976 Obj

if True == tmp17995 {
tmp17993 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V263)
}
__typedArg0 := V263
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17994 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcons, tmp17993)
}
__typedArg0 := symcons
__typedArg1 := tmp17993
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17978 Obj

if True == tmp17994 {
tmp17991 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V263)
}
__typedArg0 := V263
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17992 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp17991)
}
__typedArg0 := tmp17991
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17980 Obj

if True == tmp17992 {
tmp17988 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V263)
}
__typedArg0 := V263
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17989 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp17988)
}
__typedArg0 := tmp17988
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17990 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp17989)
}
__typedArg0 := tmp17989
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17982 Obj

if True == tmp17990 {
tmp17984 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V263)
}
__typedArg0 := V263
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17985 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp17984)
}
__typedArg0 := tmp17984
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17986 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp17985)
}
__typedArg0 := tmp17985
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17987 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp17986)
}
__typedArg0 := Nil
__typedArg1 := tmp17986
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17983 Obj

if True == tmp17987 {
ifres17983 = True


} else {
ifres17983 = False


}

ifres17982 = ifres17983


} else {
ifres17982 = False


}

var ifres17981 Obj

if True == ifres17982 {
ifres17981 = True


} else {
ifres17981 = False


}

ifres17980 = ifres17981


} else {
ifres17980 = False


}

var ifres17979 Obj

if True == ifres17980 {
ifres17979 = True


} else {
ifres17979 = False


}

ifres17978 = ifres17979


} else {
ifres17978 = False


}

var ifres17977 Obj

if True == ifres17978 {
ifres17977 = True


} else {
ifres17977 = False


}

ifres17976 = ifres17977


} else {
ifres17976 = False


}

if True == ifres17976 {
tmp17972 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V263)
}
__typedArg0 := V263
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17973 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp17972)
}
__typedArg0 := tmp17972
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17974 := Call(__e, PrimFunc(symshen_4syntax_1item_2), tmp17973)


if True == tmp17974 {
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp17968)
}
__typedArg0 := tmp17968
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17970 := Call(__e, PrimFunc(symshen_4syntax_1item_2), tmp17969)


if True == tmp17970 {
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

tmp18002 := Call(__e, ns2_1set, symshen_4syntax_1item_2, tmp17965)


_ = tmp18002

tmp18003 := MakeNative(func(__e *ControlFlow) {
V264 := __e.Get(1)
_ = V264
tmp18004 := Call(__e, PrimFunc(symshen_4_5colon_1equal_6), V264)


let__17493 := tmp18004
_ = let__17493

tmp18030 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17493)


var ifres18005 Obj

if True == tmp18030 {
tmp18006 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres18005 = tmp18006


} else {
tmp18007 := Call(__e, PrimFunc(symshen_4in_1_6), let__17493)


let__17494 := tmp18007
_ = let__17494

tmp18029 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__17494)
}
__typedArg0 := let__17494
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18008 Obj

if True == tmp18029 {
tmp18009 := Call(__e, PrimFunc(symhead), let__17494)


let__17495 := tmp18009
_ = let__17495

tmp18010 := Call(__e, PrimFunc(symtail), let__17494)


let__17496 := tmp18010
_ = let__17496

tmp18027 := Call(__e, PrimFunc(symshen_4hds_a_2), let__17496, symwhere)


var ifres18011 Obj

if True == tmp18027 {
tmp18012 := Call(__e, PrimFunc(symtail), let__17496)


let__17497 := tmp18012
_ = let__17497

tmp18025 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__17497)
}
__typedArg0 := let__17497
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18013 Obj

if True == tmp18025 {
tmp18014 := Call(__e, PrimFunc(symhead), let__17497)


let__17498 := tmp18014
_ = let__17498

tmp18015 := Call(__e, PrimFunc(symtail), let__17497)


let__17499 := tmp18015
_ = let__17499

tmp18022 := Call(__e, PrimFunc(symshen_4semicolon_2), let__17495)


tmp18023 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp18022)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp18022
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres18016 Obj

if True == tmp18023 {
tmp18017 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17495, Nil)
}
__typedArg0 := let__17495
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18018 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17498, tmp18017)
}
__typedArg0 := let__17498
__typedArg1 := tmp18017
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18019 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhere, tmp18018)
}
__typedArg0 := symwhere
__typedArg1 := tmp18018
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18020 := Call(__e, PrimFunc(symshen_4comb), let__17499, tmp18019)


ifres18016 = tmp18020


} else {
tmp18021 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres18016 = tmp18021


}

ifres18013 = ifres18016


} else {
tmp18024 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres18013 = tmp18024


}

ifres18011 = ifres18013


} else {
tmp18026 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres18011 = tmp18026


}

ifres18008 = ifres18011


} else {
tmp18028 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres18008 = tmp18028


}

ifres18005 = ifres18008


}

let__17492 := ifres18005
_ = let__17492

tmp18049 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17492)


if True == tmp18049 {
tmp18031 := Call(__e, PrimFunc(symshen_4_5colon_1equal_6), V264)


let__17501 := tmp18031
_ = let__17501

tmp18045 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17501)


var ifres18032 Obj

if True == tmp18045 {
tmp18033 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres18032 = tmp18033


} else {
tmp18034 := Call(__e, PrimFunc(symshen_4in_1_6), let__17501)


let__17502 := tmp18034
_ = let__17502

tmp18044 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__17502)
}
__typedArg0 := let__17502
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18035 Obj

if True == tmp18044 {
tmp18036 := Call(__e, PrimFunc(symhead), let__17502)


let__17503 := tmp18036
_ = let__17503

tmp18037 := Call(__e, PrimFunc(symtail), let__17502)


let__17504 := tmp18037
_ = let__17504

tmp18041 := Call(__e, PrimFunc(symshen_4semicolon_2), let__17503)


tmp18042 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp18041)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp18041
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres18038 Obj

if True == tmp18042 {
tmp18039 := Call(__e, PrimFunc(symshen_4comb), let__17504, let__17503)


ifres18038 = tmp18039


} else {
tmp18040 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres18038 = tmp18040


}

ifres18035 = ifres18038


} else {
tmp18043 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres18035 = tmp18043


}

ifres18032 = ifres18035


}

let__17500 := ifres18032
_ = let__17500

tmp18047 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17500)


if True == tmp18047 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__17500)
return
}


} else {
__e.Return(let__17492)
return
}


}, 1)

tmp18050 := Call(__e, ns2_1set, symshen_4_5semantics_6, tmp18003)


_ = tmp18050

tmp18051 := MakeNative(func(__e *ControlFlow) {
V286 := __e.Get(1)
_ = V286
V287 := __e.Get(2)
_ = V287
V288 := __e.Get(3)
_ = V288
tmp18059 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V288)
}
__typedArg0 := Nil
__typedArg1 := V288
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp18059 {
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
tmp18057 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V288)
}
__typedArg0 := V288
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp18057 {
tmp18052 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V288)
}
__typedArg0 := V288
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18053 := Call(__e, PrimFunc(symshen_4c_1rule_1_6shen), V286, tmp18052, V287)


tmp18054 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V288)
}
__typedArg0 := V288
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18055 := Call(__e, PrimFunc(symshen_4c_1rules_1_6shen), V286, V287, tmp18054)


__e.TailApply(PrimFunc(symshen_4combine_1c_1code), tmp18053, tmp18055)
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

tmp18060 := Call(__e, ns2_1set, symshen_4c_1rules_1_6shen, tmp18051)


_ = tmp18060

tmp18061 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symfail))
return
}, 0)

tmp18062 := Call(__e, ns2_1set, symshen_4parse_1failure, tmp18061)


_ = tmp18062

tmp18063 := MakeNative(func(__e *ControlFlow) {
V289 := __e.Get(1)
_ = V289
V290 := __e.Get(2)
_ = V290
tmp18064 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symResult, Nil)
}
__typedArg0 := symResult
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18065 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4parse_1failure_2, tmp18064)
}
__typedArg0 := symshen_4parse_1failure_2
__typedArg1 := tmp18064
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18066 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symResult, Nil)
}
__typedArg0 := symResult
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18067 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V290, tmp18066)
}
__typedArg0 := V290
__typedArg1 := tmp18066
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18068 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18065, tmp18067)
}
__typedArg0 := tmp18065
__typedArg1 := tmp18067
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18069 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp18068)
}
__typedArg0 := symif
__typedArg1 := tmp18068
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18070 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18069, Nil)
}
__typedArg0 := tmp18069
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18071 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V289, tmp18070)
}
__typedArg0 := V289
__typedArg1 := tmp18070
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18072 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symResult, tmp18071)
}
__typedArg0 := symResult
__typedArg1 := tmp18071
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp18072)
}
__typedArg0 := symlet
__typedArg1 := tmp18072
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 2)

tmp18073 := Call(__e, ns2_1set, symshen_4combine_1c_1code, tmp18063)


_ = tmp18073

tmp18074 := MakeNative(func(__e *ControlFlow) {
V297 := __e.Get(1)
_ = V297
V298 := __e.Get(2)
_ = V298
V299 := __e.Get(3)
_ = V299
tmp18088 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V298)
}
__typedArg0 := V298
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18079 Obj

if True == tmp18088 {
tmp18086 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V298)
}
__typedArg0 := V298
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18087 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18086)
}
__typedArg0 := tmp18086
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18081 Obj

if True == tmp18087 {
tmp18083 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V298)
}
__typedArg0 := V298
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18084 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18083)
}
__typedArg0 := tmp18083
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18085 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp18084)
}
__typedArg0 := Nil
__typedArg1 := tmp18084
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18082 Obj

if True == tmp18085 {
ifres18082 = True


} else {
ifres18082 = False


}

ifres18081 = ifres18082


} else {
ifres18081 = False


}

var ifres18080 Obj

if True == ifres18081 {
ifres18080 = True


} else {
ifres18080 = False


}

ifres18079 = ifres18080


} else {
ifres18079 = False


}

if True == ifres18079 {
tmp18075 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V298)
}
__typedArg0 := V298
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18076 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V298)
}
__typedArg0 := V298
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18077 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18076)
}
__typedArg0 := tmp18076
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4yacc_1syntax), V297, V299, tmp18075, tmp18077)
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

tmp18089 := Call(__e, ns2_1set, symshen_4c_1rule_1_6shen, tmp18074)


_ = tmp18089

tmp18090 := MakeNative(func(__e *ControlFlow) {
V308 := __e.Get(1)
_ = V308
V309 := __e.Get(2)
_ = V309
V310 := __e.Get(3)
_ = V310
V311 := __e.Get(4)
_ = V311
tmp18154 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V310)
}
__typedArg0 := Nil
__typedArg1 := V310
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18132 Obj

if True == tmp18154 {
tmp18153 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V311)
}
__typedArg0 := V311
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18134 Obj

if True == tmp18153 {
tmp18151 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V311)
}
__typedArg0 := V311
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18152 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symwhere, tmp18151)
}
__typedArg0 := symwhere
__typedArg1 := tmp18151
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18136 Obj

if True == tmp18152 {
tmp18149 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V311)
}
__typedArg0 := V311
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18150 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18149)
}
__typedArg0 := tmp18149
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18138 Obj

if True == tmp18150 {
tmp18146 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V311)
}
__typedArg0 := V311
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18147 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18146)
}
__typedArg0 := tmp18146
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18148 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18147)
}
__typedArg0 := tmp18147
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18140 Obj

if True == tmp18148 {
tmp18142 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V311)
}
__typedArg0 := V311
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18143 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18142)
}
__typedArg0 := tmp18142
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18144 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18143)
}
__typedArg0 := tmp18143
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18145 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp18144)
}
__typedArg0 := Nil
__typedArg1 := tmp18144
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18141 Obj

if True == tmp18145 {
ifres18141 = True


} else {
ifres18141 = False


}

ifres18140 = ifres18141


} else {
ifres18140 = False


}

var ifres18139 Obj

if True == ifres18140 {
ifres18139 = True


} else {
ifres18139 = False


}

ifres18138 = ifres18139


} else {
ifres18138 = False


}

var ifres18137 Obj

if True == ifres18138 {
ifres18137 = True


} else {
ifres18137 = False


}

ifres18136 = ifres18137


} else {
ifres18136 = False


}

var ifres18135 Obj

if True == ifres18136 {
ifres18135 = True


} else {
ifres18135 = False


}

ifres18134 = ifres18135


} else {
ifres18134 = False


}

var ifres18133 Obj

if True == ifres18134 {
ifres18133 = True


} else {
ifres18133 = False


}

ifres18132 = ifres18133


} else {
ifres18132 = False


}

if True == ifres18132 {
tmp18091 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V311)
}
__typedArg0 := V311
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18092 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18091)
}
__typedArg0 := tmp18091
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18093 := Call(__e, PrimFunc(symshen_4process_1yacc_1semantics), tmp18092)


tmp18094 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V311)
}
__typedArg0 := V311
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18095 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18094)
}
__typedArg0 := tmp18094
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18096 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18095)
}
__typedArg0 := tmp18095
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18097 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V308, V309, Nil, tmp18096)


tmp18098 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4parse_1failure, Nil)
}
__typedArg0 := symshen_4parse_1failure
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18099 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18098, Nil)
}
__typedArg0 := tmp18098
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18100 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18097, tmp18099)
}
__typedArg0 := tmp18097
__typedArg1 := tmp18099
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18101 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18093, tmp18100)
}
__typedArg0 := tmp18093
__typedArg1 := tmp18100
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp18101)
}
__typedArg0 := symif
__typedArg1 := tmp18101
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp18130 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V310)
}
__typedArg0 := Nil
__typedArg1 := V310
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp18130 {
__e.TailApply(PrimFunc(symshen_4yacc_1semantics), V308, V309, V311)
return
} else {
tmp18128 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp18128 {
tmp18125 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18126 := Call(__e, PrimFunc(symshen_4non_1terminal_2), tmp18125)


if True == tmp18126 {
tmp18102 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18103 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4non_1terminalcode), V308, V309, tmp18102, tmp18103, V311)
return


} else {
tmp18122 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18123 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(tmp18122)
}
__typedArg0 := tmp18122
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

if True == tmp18123 {
tmp18104 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18105 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4variablecode), V308, V309, tmp18104, tmp18105, V311)
return


} else {
tmp18119 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18120 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym__, tmp18119)
}
__typedArg0 := sym__
__typedArg1 := tmp18119
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp18120 {
tmp18106 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18107 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4wildcardcode), V308, V309, tmp18106, tmp18107, V311)
return


} else {
tmp18116 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18117 := Call(__e, PrimFunc(symatom_2), tmp18116)


if True == tmp18117 {
tmp18108 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18109 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4terminalcode), V308, V309, tmp18108, tmp18109, V311)
return


} else {
tmp18113 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18114 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18113)
}
__typedArg0 := tmp18113
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp18114 {
tmp18110 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18111 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4conscode), V308, V309, tmp18110, tmp18111, V311)
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

tmp18155 := Call(__e, ns2_1set, symshen_4yacc_1syntax, tmp18090)


_ = tmp18155

tmp18156 := MakeNative(func(__e *ControlFlow) {
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
tmp18157 := Call(__e, PrimFunc(symconcat), symParse, V314)


let__17505 := tmp18157
_ = let__17505

tmp18158 := Call(__e, PrimFunc(symconcat), symAction, V314)


let__17506 := tmp18158
_ = let__17506

tmp18159 := Call(__e, PrimFunc(symconcat), symRemainder, V314)


let__17507 := tmp18159
_ = let__17507

tmp18160 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V313, Nil)
}
__typedArg0 := V313
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18161 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V314, tmp18160)
}
__typedArg0 := V314
__typedArg1 := tmp18160
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18162 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17505, Nil)
}
__typedArg0 := let__17505
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18163 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4parse_1failure_2, tmp18162)
}
__typedArg0 := symshen_4parse_1failure_2
__typedArg1 := tmp18162
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18164 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4parse_1failure, Nil)
}
__typedArg0 := symshen_4parse_1failure
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18165 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17505, Nil)
}
__typedArg0 := let__17505
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18166 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4in_1_6, tmp18165)
}
__typedArg0 := symshen_4in_1_6
__typedArg1 := tmp18165
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18167 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V312, let__17507, V315, V316)


tmp18168 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18167, Nil)
}
__typedArg0 := tmp18167
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18169 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18166, tmp18168)
}
__typedArg0 := tmp18166
__typedArg1 := tmp18168
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18170 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17507, tmp18169)
}
__typedArg0 := let__17507
__typedArg1 := tmp18169
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18171 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp18170)
}
__typedArg0 := symlet
__typedArg1 := tmp18170
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__17508 := tmp18171
_ = let__17508

tmp18182 := Call(__e, PrimFunc(symshen_4occurs_1check_2), V314, V316)


var ifres18179 Obj

if True == tmp18182 {
ifres18179 = True


} else {
tmp18181 := Call(__e, PrimFunc(symshen_4occurs_1check_2), let__17506, V316)


var ifres18180 Obj

if True == tmp18181 {
ifres18180 = True


} else {
ifres18180 = False


}

ifres18179 = ifres18180


}

var ifres18172 Obj

if True == ifres18179 {
tmp18173 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17505, Nil)
}
__typedArg0 := let__17505
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18174 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4_5_1out, tmp18173)
}
__typedArg0 := symshen_4_5_1out
__typedArg1 := tmp18173
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18175 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17508, Nil)
}
__typedArg0 := let__17508
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18176 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18174, tmp18175)
}
__typedArg0 := tmp18174
__typedArg1 := tmp18175
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18177 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17506, tmp18176)
}
__typedArg0 := let__17506
__typedArg1 := tmp18176
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18178 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp18177)
}
__typedArg0 := symlet
__typedArg1 := tmp18177
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

ifres18172 = tmp18178


} else {
ifres18172 = let__17508


}

tmp18183 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(ifres18172, Nil)
}
__typedArg0 := ifres18172
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18184 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18164, tmp18183)
}
__typedArg0 := tmp18164
__typedArg1 := tmp18183
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18185 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18163, tmp18184)
}
__typedArg0 := tmp18163
__typedArg1 := tmp18184
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18186 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp18185)
}
__typedArg0 := symif
__typedArg1 := tmp18185
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18187 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18186, Nil)
}
__typedArg0 := tmp18186
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18188 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18161, tmp18187)
}
__typedArg0 := tmp18161
__typedArg1 := tmp18187
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18189 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17505, tmp18188)
}
__typedArg0 := let__17505
__typedArg1 := tmp18188
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp18189)
}
__typedArg0 := symlet
__typedArg1 := tmp18189
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 5)

tmp18190 := Call(__e, ns2_1set, symshen_4non_1terminalcode, tmp18156)


_ = tmp18190

tmp18191 := MakeNative(func(__e *ControlFlow) {
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
tmp18192 := Call(__e, PrimFunc(symgensym), symRemainder)


let__17509 := tmp18192
_ = let__17509

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
return PrimCons(symcons_2, tmp18193)
}
__typedArg0 := symcons_2
__typedArg1 := tmp18193
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18195 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V322, Nil)
}
__typedArg0 := V322
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18196 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtail, tmp18195)
}
__typedArg0 := symtail
__typedArg1 := tmp18195
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18197 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V321, let__17509, V324, V325)


tmp18198 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18197, Nil)
}
__typedArg0 := tmp18197
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18199 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18196, tmp18198)
}
__typedArg0 := tmp18196
__typedArg1 := tmp18198
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18200 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17509, tmp18199)
}
__typedArg0 := let__17509
__typedArg1 := tmp18199
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18201 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp18200)
}
__typedArg0 := symlet
__typedArg1 := tmp18200
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__17510 := tmp18201
_ = let__17510

tmp18209 := Call(__e, PrimFunc(symshen_4occurs_1check_2), V323, V325)


var ifres18202 Obj

if True == tmp18209 {
tmp18203 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V322, Nil)
}
__typedArg0 := V322
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18204 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhead, tmp18203)
}
__typedArg0 := symhead
__typedArg1 := tmp18203
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18205 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17510, Nil)
}
__typedArg0 := let__17510
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18206 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18204, tmp18205)
}
__typedArg0 := tmp18204
__typedArg1 := tmp18205
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18207 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V323, tmp18206)
}
__typedArg0 := V323
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

ifres18202 = tmp18208


} else {
ifres18202 = let__17510


}

tmp18210 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4parse_1failure, Nil)
}
__typedArg0 := symshen_4parse_1failure
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18211 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18210, Nil)
}
__typedArg0 := tmp18210
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18212 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(ifres18202, tmp18211)
}
__typedArg0 := ifres18202
__typedArg1 := tmp18211
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18213 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18194, tmp18212)
}
__typedArg0 := tmp18194
__typedArg1 := tmp18212
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp18213)
}
__typedArg0 := symif
__typedArg1 := tmp18213
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 5)

tmp18214 := Call(__e, ns2_1set, symshen_4variablecode, tmp18191)


_ = tmp18214

tmp18215 := MakeNative(func(__e *ControlFlow) {
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
tmp18216 := Call(__e, PrimFunc(symgensym), symRemainder)


let__17511 := tmp18216
_ = let__17511

tmp18217 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V329, Nil)
}
__typedArg0 := V329
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18218 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons_2, tmp18217)
}
__typedArg0 := symcons_2
__typedArg1 := tmp18217
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18219 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V329, Nil)
}
__typedArg0 := V329
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

tmp18221 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V328, let__17511, V331, V332)


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
return PrimCons(let__17511, tmp18223)
}
__typedArg0 := let__17511
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

tmp18230 := Call(__e, ns2_1set, symshen_4wildcardcode, tmp18215)


_ = tmp18230

tmp18231 := MakeNative(func(__e *ControlFlow) {
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
tmp18232 := Call(__e, PrimFunc(symgensym), symRemainder)


let__17512 := tmp18232
_ = let__17512

tmp18233 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V336, Nil)
}
__typedArg0 := V336
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18234 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V335, tmp18233)
}
__typedArg0 := V335
__typedArg1 := tmp18233
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18235 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4hds_a_2, tmp18234)
}
__typedArg0 := symshen_4hds_a_2
__typedArg1 := tmp18234
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18236 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V335, Nil)
}
__typedArg0 := V335
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18237 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtail, tmp18236)
}
__typedArg0 := symtail
__typedArg1 := tmp18236
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18238 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V334, let__17512, V337, V338)


tmp18239 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18238, Nil)
}
__typedArg0 := tmp18238
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18240 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18237, tmp18239)
}
__typedArg0 := tmp18237
__typedArg1 := tmp18239
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18241 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17512, tmp18240)
}
__typedArg0 := let__17512
__typedArg1 := tmp18240
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18242 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp18241)
}
__typedArg0 := symlet
__typedArg1 := tmp18241
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18243 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4parse_1failure, Nil)
}
__typedArg0 := symshen_4parse_1failure
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18244 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18243, Nil)
}
__typedArg0 := tmp18243
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18245 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18242, tmp18244)
}
__typedArg0 := tmp18242
__typedArg1 := tmp18244
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18246 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18235, tmp18245)
}
__typedArg0 := tmp18235
__typedArg1 := tmp18245
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp18246)
}
__typedArg0 := symif
__typedArg1 := tmp18246
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 5)

tmp18247 := Call(__e, ns2_1set, symshen_4terminalcode, tmp18231)


_ = tmp18247

tmp18248 := MakeNative(func(__e *ControlFlow) {
V347 := __e.Get(1)
_ = V347
V348 := __e.Get(2)
_ = V348
tmp18254 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V347)
}
__typedArg0 := V347
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18250 Obj

if True == tmp18254 {
tmp18252 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V347)
}
__typedArg0 := V347
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18253 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp18252, V348)
}
__typedArg0 := tmp18252
__typedArg1 := V348
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18251 Obj

if True == tmp18253 {
ifres18251 = True


} else {
ifres18251 = False


}

ifres18250 = ifres18251


} else {
ifres18250 = False


}

if True == ifres18250 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}, 2)

tmp18255 := Call(__e, ns2_1set, symshen_4hds_a_2, tmp18248)


_ = tmp18255

tmp18256 := MakeNative(func(__e *ControlFlow) {
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
tmp18257 := Call(__e, PrimFunc(symgensym), symRemainder)


let__17513 := tmp18257
_ = let__17513

tmp18258 := Call(__e, PrimFunc(symgensym), symHd)


let__17514 := tmp18258
_ = let__17514

tmp18259 := Call(__e, PrimFunc(symgensym), symTl)


let__17515 := tmp18259
_ = let__17515

tmp18260 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V350, Nil)
}
__typedArg0 := V350
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18261 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4ccons_2, tmp18260)
}
__typedArg0 := symshen_4ccons_2
__typedArg1 := tmp18260
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18262 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V350, Nil)
}
__typedArg0 := V350
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18263 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhead, tmp18262)
}
__typedArg0 := symhead
__typedArg1 := tmp18262
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18264 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V350, Nil)
}
__typedArg0 := V350
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18265 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtail, tmp18264)
}
__typedArg0 := symtail
__typedArg1 := tmp18264
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18266 := Call(__e, PrimFunc(symshen_4decons), V351)


tmp18267 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5end_6, Nil)
}
__typedArg0 := sym_5end_6
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18268 := Call(__e, PrimFunc(symappend), tmp18266, tmp18267)


tmp18269 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V349, let__17515, V352, V353)


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
return PrimCons(symshen_4processed, tmp18270)
}
__typedArg0 := symshen_4processed
__typedArg1 := tmp18270
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18272 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V349, let__17514, tmp18268, tmp18271)


tmp18273 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18272, Nil)
}
__typedArg0 := tmp18272
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18274 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18265, tmp18273)
}
__typedArg0 := tmp18265
__typedArg1 := tmp18273
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18275 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17515, tmp18274)
}
__typedArg0 := let__17515
__typedArg1 := tmp18274
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18276 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18263, tmp18275)
}
__typedArg0 := tmp18263
__typedArg1 := tmp18275
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18277 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17514, tmp18276)
}
__typedArg0 := let__17514
__typedArg1 := tmp18276
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18278 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp18277)
}
__typedArg0 := symlet
__typedArg1 := tmp18277
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18279 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4parse_1failure, Nil)
}
__typedArg0 := symshen_4parse_1failure
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18280 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18279, Nil)
}
__typedArg0 := tmp18279
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18281 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18278, tmp18280)
}
__typedArg0 := tmp18278
__typedArg1 := tmp18280
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18282 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18261, tmp18281)
}
__typedArg0 := tmp18261
__typedArg1 := tmp18281
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp18282)
}
__typedArg0 := symif
__typedArg1 := tmp18282
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 5)

tmp18283 := Call(__e, ns2_1set, symshen_4conscode, tmp18256)


_ = tmp18283

tmp18284 := MakeNative(func(__e *ControlFlow) {
V367 := __e.Get(1)
_ = V367
tmp18296 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V367)
}
__typedArg0 := V367
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18292 Obj

if True == tmp18296 {
tmp18294 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V367)
}
__typedArg0 := V367
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18295 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18294)
}
__typedArg0 := tmp18294
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18293 Obj

if True == tmp18295 {
ifres18293 = True


} else {
ifres18293 = False


}

ifres18292 = ifres18293


} else {
ifres18292 = False


}

if True == ifres18292 {
__e.Return(True)
return
} else {
tmp18290 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V367)
}
__typedArg0 := V367
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18286 Obj

if True == tmp18290 {
tmp18288 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V367)
}
__typedArg0 := V367
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18289 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp18288)
}
__typedArg0 := Nil
__typedArg1 := tmp18288
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18287 Obj

if True == tmp18289 {
ifres18287 = True


} else {
ifres18287 = False


}

ifres18286 = ifres18287


} else {
ifres18286 = False


}

if True == ifres18286 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp18297 := Call(__e, ns2_1set, symshen_4ccons_2, tmp18284)


_ = tmp18297

tmp18298 := MakeNative(func(__e *ControlFlow) {
V368 := __e.Get(1)
_ = V368
tmp18325 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V368)
}
__typedArg0 := V368
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18306 Obj

if True == tmp18325 {
tmp18323 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V368)
}
__typedArg0 := V368
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18324 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcons, tmp18323)
}
__typedArg0 := symcons
__typedArg1 := tmp18323
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18308 Obj

if True == tmp18324 {
tmp18321 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V368)
}
__typedArg0 := V368
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18322 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18321)
}
__typedArg0 := tmp18321
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18310 Obj

if True == tmp18322 {
tmp18318 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V368)
}
__typedArg0 := V368
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18319 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18318)
}
__typedArg0 := tmp18318
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18320 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18319)
}
__typedArg0 := tmp18319
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18312 Obj

if True == tmp18320 {
tmp18314 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V368)
}
__typedArg0 := V368
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18315 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18314)
}
__typedArg0 := tmp18314
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18316 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18315)
}
__typedArg0 := tmp18315
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18317 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp18316)
}
__typedArg0 := Nil
__typedArg1 := tmp18316
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18313 Obj

if True == tmp18317 {
ifres18313 = True


} else {
ifres18313 = False


}

ifres18312 = ifres18313


} else {
ifres18312 = False


}

var ifres18311 Obj

if True == ifres18312 {
ifres18311 = True


} else {
ifres18311 = False


}

ifres18310 = ifres18311


} else {
ifres18310 = False


}

var ifres18309 Obj

if True == ifres18310 {
ifres18309 = True


} else {
ifres18309 = False


}

ifres18308 = ifres18309


} else {
ifres18308 = False


}

var ifres18307 Obj

if True == ifres18308 {
ifres18307 = True


} else {
ifres18307 = False


}

ifres18306 = ifres18307


} else {
ifres18306 = False


}

if True == ifres18306 {
tmp18299 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V368)
}
__typedArg0 := V368
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18300 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18299)
}
__typedArg0 := tmp18299
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18302)
}
__typedArg0 := tmp18302
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18304 := Call(__e, PrimFunc(symshen_4decons), tmp18303)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18300, tmp18304)
}
__typedArg0 := tmp18300
__typedArg1 := tmp18304
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return(V368)
return
}


}, 1)

tmp18326 := Call(__e, ns2_1set, symshen_4decons, tmp18298)


_ = tmp18326

tmp18327 := MakeNative(func(__e *ControlFlow) {
V369 := __e.Get(1)
_ = V369
V370 := __e.Get(2)
_ = V370
tmp18328 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V370, Nil)
}
__typedArg0 := V370
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V369, tmp18328)
}
__typedArg0 := V369
__typedArg1 := tmp18328
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 2)

tmp18329 := Call(__e, ns2_1set, symshen_4comb, tmp18327)


_ = tmp18329

tmp18330 := MakeNative(func(__e *ControlFlow) {
V375 := __e.Get(1)
_ = V375
V376 := __e.Get(2)
_ = V376
V377 := __e.Get(3)
_ = V377
tmp18350 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V377)
}
__typedArg0 := V377
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18337 Obj

if True == tmp18350 {
tmp18348 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V377)
}
__typedArg0 := V377
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18349 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4processed, tmp18348)
}
__typedArg0 := symshen_4processed
__typedArg1 := tmp18348
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18339 Obj

if True == tmp18349 {
tmp18346 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V377)
}
__typedArg0 := V377
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18347 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18346)
}
__typedArg0 := tmp18346
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18341 Obj

if True == tmp18347 {
tmp18343 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V377)
}
__typedArg0 := V377
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp18344)
}
__typedArg0 := Nil
__typedArg1 := tmp18344
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18342 Obj

if True == tmp18345 {
ifres18342 = True


} else {
ifres18342 = False


}

ifres18341 = ifres18342


} else {
ifres18341 = False


}

var ifres18340 Obj

if True == ifres18341 {
ifres18340 = True


} else {
ifres18340 = False


}

ifres18339 = ifres18340


} else {
ifres18339 = False


}

var ifres18338 Obj

if True == ifres18339 {
ifres18338 = True


} else {
ifres18338 = False


}

ifres18337 = ifres18338


} else {
ifres18337 = False


}

if True == ifres18337 {
tmp18331 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V377)
}
__typedArg0 := V377
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18331)
}
__typedArg0 := tmp18331
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return


} else {
tmp18332 := Call(__e, PrimFunc(symshen_4process_1yacc_1semantics), V377)


let__17516 := tmp18332
_ = let__17516

tmp18333 := Call(__e, PrimFunc(symshen_4use_1type_1info), V375, let__17516)


let__17517 := tmp18333
_ = let__17517

tmp18334 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17517, Nil)
}
__typedArg0 := let__17517
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18335 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V376, tmp18334)
}
__typedArg0 := V376
__typedArg1 := tmp18334
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4comb, tmp18335)
}
__typedArg0 := symshen_4comb
__typedArg1 := tmp18335
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}


}, 3)

tmp18351 := Call(__e, ns2_1set, symshen_4yacc_1semantics, tmp18330)


_ = tmp18351

tmp18352 := MakeNative(func(__e *ControlFlow) {
V383 := __e.Get(1)
_ = V383
V384 := __e.Get(2)
_ = V384
tmp18540 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18361 Obj

if True == tmp18540 {
tmp18538 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18539 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_i, tmp18538)
}
__typedArg0 := sym_i
__typedArg1 := tmp18538
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18363 Obj

if True == tmp18539 {
tmp18536 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18537 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18536)
}
__typedArg0 := tmp18536
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18365 Obj

if True == tmp18537 {
tmp18533 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18534 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18533)
}
__typedArg0 := tmp18533
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18535 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18534)
}
__typedArg0 := tmp18534
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18367 Obj

if True == tmp18535 {
tmp18529 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18530 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18529)
}
__typedArg0 := tmp18529
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18531 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18530)
}
__typedArg0 := tmp18530
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18532 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlist, tmp18531)
}
__typedArg0 := symlist
__typedArg1 := tmp18531
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18369 Obj

if True == tmp18532 {
tmp18525 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18526 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18525)
}
__typedArg0 := tmp18525
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18527 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18526)
}
__typedArg0 := tmp18526
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18528 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18527)
}
__typedArg0 := tmp18527
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18371 Obj

if True == tmp18528 {
tmp18520 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18521 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18520)
}
__typedArg0 := tmp18520
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18522 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18521)
}
__typedArg0 := tmp18521
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18523 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18522)
}
__typedArg0 := tmp18522
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18524 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp18523)
}
__typedArg0 := Nil
__typedArg1 := tmp18523
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18373 Obj

if True == tmp18524 {
tmp18517 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18518 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18517)
}
__typedArg0 := tmp18517
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18519 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18518)
}
__typedArg0 := tmp18518
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18375 Obj

if True == tmp18519 {
tmp18513 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18514 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18513)
}
__typedArg0 := tmp18513
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18515 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18514)
}
__typedArg0 := tmp18514
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18516 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_1_1_6, tmp18515)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp18515
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18377 Obj

if True == tmp18516 {
tmp18509 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18510 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18509)
}
__typedArg0 := tmp18509
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18511 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18510)
}
__typedArg0 := tmp18510
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18512 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18511)
}
__typedArg0 := tmp18511
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18379 Obj

if True == tmp18512 {
tmp18504 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18506)
}
__typedArg0 := tmp18506
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18508 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18507)
}
__typedArg0 := tmp18507
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18381 Obj

if True == tmp18508 {
tmp18498 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18499 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18498)
}
__typedArg0 := tmp18498
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18500)
}
__typedArg0 := tmp18500
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18502 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18501)
}
__typedArg0 := tmp18501
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18503 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symstr, tmp18502)
}
__typedArg0 := symstr
__typedArg1 := tmp18502
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18383 Obj

if True == tmp18503 {
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18494)
}
__typedArg0 := tmp18494
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18496 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18495)
}
__typedArg0 := tmp18495
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18497 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18496)
}
__typedArg0 := tmp18496
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18385 Obj

if True == tmp18497 {
tmp18485 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18486 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18485)
}
__typedArg0 := tmp18485
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18487 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18486)
}
__typedArg0 := tmp18486
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18488 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18487)
}
__typedArg0 := tmp18487
return Call(__e, PrimFunc(symhd), __typedArg0)
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

var ifres18387 Obj

if True == tmp18491 {
tmp18477 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18478 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18477)
}
__typedArg0 := tmp18477
return Call(__e, PrimFunc(symtl), __typedArg0)
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18480)
}
__typedArg0 := tmp18480
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18482 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18481)
}
__typedArg0 := tmp18481
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18483 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18482)
}
__typedArg0 := tmp18482
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18484 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlist, tmp18483)
}
__typedArg0 := symlist
__typedArg1 := tmp18483
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18389 Obj

if True == tmp18484 {
tmp18469 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18472)
}
__typedArg0 := tmp18472
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18474 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18473)
}
__typedArg0 := tmp18473
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18475 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18474)
}
__typedArg0 := tmp18474
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18476 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18475)
}
__typedArg0 := tmp18475
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18391 Obj

if True == tmp18476 {
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18465)
}
__typedArg0 := tmp18465
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18467 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18466)
}
__typedArg0 := tmp18466
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18468 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp18467)
}
__typedArg0 := Nil
__typedArg1 := tmp18467
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18393 Obj

if True == tmp18468 {
tmp18453 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18454)
}
__typedArg0 := tmp18454
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18456 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18455)
}
__typedArg0 := tmp18455
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18457 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18456)
}
__typedArg0 := tmp18456
return Call(__e, PrimFunc(symtl), __typedArg0)
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

var ifres18395 Obj

if True == tmp18459 {
tmp18445 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18446 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18445)
}
__typedArg0 := tmp18445
return Call(__e, PrimFunc(symtl), __typedArg0)
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18450)
}
__typedArg0 := tmp18450
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18452 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp18451)
}
__typedArg0 := Nil
__typedArg1 := tmp18451
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18397 Obj

if True == tmp18452 {
tmp18440 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18441)
}
__typedArg0 := tmp18441
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18443 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18442)
}
__typedArg0 := tmp18442
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18444 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18443)
}
__typedArg0 := tmp18443
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18399 Obj

if True == tmp18444 {
tmp18434 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18435 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18434)
}
__typedArg0 := tmp18434
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_j, tmp18438)
}
__typedArg0 := sym_j
__typedArg1 := tmp18438
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18401 Obj

if True == tmp18439 {
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp18432)
}
__typedArg0 := Nil
__typedArg1 := tmp18432
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18403 Obj

if True == tmp18433 {
tmp18415 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18416)
}
__typedArg0 := tmp18416
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18418 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18417)
}
__typedArg0 := tmp18417
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18419 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18421)
}
__typedArg0 := tmp18421
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18423 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18422)
}
__typedArg0 := tmp18422
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18424 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18423)
}
__typedArg0 := tmp18423
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18425 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18424)
}
__typedArg0 := tmp18424
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18426 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18425)
}
__typedArg0 := tmp18425
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18427 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp18418, tmp18426)
}
__typedArg0 := tmp18418
__typedArg1 := tmp18426
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18405 Obj

if True == tmp18427 {
tmp18407 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18408 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18407)
}
__typedArg0 := tmp18407
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18409 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18408)
}
__typedArg0 := tmp18408
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18410 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18409)
}
__typedArg0 := tmp18409
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18411 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18410)
}
__typedArg0 := tmp18410
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18412)
}
__typedArg0 := tmp18412
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18414 := Call(__e, PrimFunc(symshen_4monomorphic_2), tmp18413)


var ifres18406 Obj

if True == tmp18414 {
ifres18406 = True


} else {
ifres18406 = False


}

ifres18405 = ifres18406


} else {
ifres18405 = False


}

var ifres18404 Obj

if True == ifres18405 {
ifres18404 = True


} else {
ifres18404 = False


}

ifres18403 = ifres18404


} else {
ifres18403 = False


}

var ifres18402 Obj

if True == ifres18403 {
ifres18402 = True


} else {
ifres18402 = False


}

ifres18401 = ifres18402


} else {
ifres18401 = False


}

var ifres18400 Obj

if True == ifres18401 {
ifres18400 = True


} else {
ifres18400 = False


}

ifres18399 = ifres18400


} else {
ifres18399 = False


}

var ifres18398 Obj

if True == ifres18399 {
ifres18398 = True


} else {
ifres18398 = False


}

ifres18397 = ifres18398


} else {
ifres18397 = False


}

var ifres18396 Obj

if True == ifres18397 {
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

if True == ifres18361 {
tmp18353 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18354 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18353)
}
__typedArg0 := tmp18353
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18355 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18354)
}
__typedArg0 := tmp18354
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18356 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18355)
}
__typedArg0 := tmp18355
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18357 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18356)
}
__typedArg0 := tmp18356
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18358 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18357)
}
__typedArg0 := tmp18357
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18359 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V384, tmp18358)
}
__typedArg0 := V384
__typedArg1 := tmp18358
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtype, tmp18359)
}
__typedArg0 := symtype
__typedArg1 := tmp18359
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return(V384)
return
}


}, 2)

tmp18541 := Call(__e, ns2_1set, symshen_4use_1type_1info, tmp18352)


_ = tmp18541

tmp18542 := MakeNative(func(__e *ControlFlow) {
V387 := __e.Get(1)
_ = V387
tmp18552 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(V387)
}
__typedArg0 := V387
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

if True == tmp18552 {
__e.Return(False)
return
} else {
tmp18550 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V387)
}
__typedArg0 := V387
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp18550 {
tmp18547 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V387)
}
__typedArg0 := V387
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18548 := Call(__e, PrimFunc(symshen_4monomorphic_2), tmp18547)


if True == tmp18548 {
tmp18544 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V387)
}
__typedArg0 := V387
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18545 := Call(__e, PrimFunc(symshen_4monomorphic_2), tmp18544)


if True == tmp18545 {
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

tmp18553 := Call(__e, ns2_1set, symshen_4monomorphic_2, tmp18542)


_ = tmp18553

tmp18554 := MakeNative(func(__e *ControlFlow) {
V388 := __e.Get(1)
_ = V388
tmp18580 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V388)
}
__typedArg0 := V388
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18562 Obj

if True == tmp18580 {
tmp18578 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V388)
}
__typedArg0 := V388
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18579 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symprotect, tmp18578)
}
__typedArg0 := symprotect
__typedArg1 := tmp18578
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18564 Obj

if True == tmp18579 {
tmp18576 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V388)
}
__typedArg0 := V388
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18577 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18576)
}
__typedArg0 := tmp18576
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18566 Obj

if True == tmp18577 {
tmp18573 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V388)
}
__typedArg0 := V388
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

var ifres18568 Obj

if True == tmp18575 {
tmp18570 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V388)
}
__typedArg0 := V388
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18571 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18570)
}
__typedArg0 := tmp18570
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18572 := Call(__e, PrimFunc(symshen_4non_1terminal_2), tmp18571)


var ifres18569 Obj

if True == tmp18572 {
ifres18569 = True


} else {
ifres18569 = False


}

ifres18568 = ifres18569


} else {
ifres18568 = False


}

var ifres18567 Obj

if True == ifres18568 {
ifres18567 = True


} else {
ifres18567 = False


}

ifres18566 = ifres18567


} else {
ifres18566 = False


}

var ifres18565 Obj

if True == ifres18566 {
ifres18565 = True


} else {
ifres18565 = False


}

ifres18564 = ifres18565


} else {
ifres18564 = False


}

var ifres18563 Obj

if True == ifres18564 {
ifres18563 = True


} else {
ifres18563 = False


}

ifres18562 = ifres18563


} else {
ifres18562 = False


}

if True == ifres18562 {
tmp18555 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V388)
}
__typedArg0 := V388
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18555)
}
__typedArg0 := tmp18555
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return


} else {
tmp18560 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V388)
}
__typedArg0 := V388
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp18560 {
tmp18556 := MakeNative(func(__e *ControlFlow) {
Z389 := __e.Get(1)
_ = Z389
__e.TailApply(PrimFunc(symshen_4process_1yacc_1semantics), Z389)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp18556, V388)
return


} else {
tmp18558 := Call(__e, PrimFunc(symshen_4non_1terminal_2), V388)


if True == tmp18558 {
__e.TailApply(PrimFunc(symconcat), symAction, V388)
return
} else {
__e.Return(V388)
return
}


}


}


}, 1)

tmp18581 := Call(__e, ns2_1set, symshen_4process_1yacc_1semantics, tmp18554)


_ = tmp18581

tmp18582 := MakeNative(func(__e *ControlFlow) {
V392 := __e.Get(1)
_ = V392
tmp18595 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V392)
}
__typedArg0 := V392
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18586 Obj

if True == tmp18595 {
tmp18593 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V392)
}
__typedArg0 := V392
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18594 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18593)
}
__typedArg0 := tmp18593
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18588 Obj

if True == tmp18594 {
tmp18590 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V392)
}
__typedArg0 := V392
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18591 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18590)
}
__typedArg0 := tmp18590
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18592 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp18591)
}
__typedArg0 := Nil
__typedArg1 := tmp18591
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18589 Obj

if True == tmp18592 {
ifres18589 = True


} else {
ifres18589 = False


}

ifres18588 = ifres18589


} else {
ifres18588 = False


}

var ifres18587 Obj

if True == ifres18588 {
ifres18587 = True


} else {
ifres18587 = False


}

ifres18586 = ifres18587


} else {
ifres18586 = False


}

if True == ifres18586 {
tmp18583 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V392)
}
__typedArg0 := V392
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18583)
}
__typedArg0 := tmp18583
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return


} else {
tmp18584 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V392)
}
__typedArg0 := V392
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18584)
}
__typedArg0 := tmp18584
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return


}


}, 1)

tmp18596 := Call(__e, ns2_1set, symshen_4_5_1out, tmp18582)


_ = tmp18596

tmp18597 := MakeNative(func(__e *ControlFlow) {
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

tmp18598 := Call(__e, ns2_1set, symshen_4in_1_6, tmp18597)


_ = tmp18598

tmp18599 := MakeNative(func(__e *ControlFlow) {
V394 := __e.Get(1)
_ = V394
tmp18600 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V394, Nil)
}
__typedArg0 := V394
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Nil, tmp18600)
}
__typedArg0 := Nil
__typedArg1 := tmp18600
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp18601 := Call(__e, ns2_1set, sym_5_b_6, tmp18599)


_ = tmp18601

tmp18602 := MakeNative(func(__e *ControlFlow) {
V395 := __e.Get(1)
_ = V395
tmp18603 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Nil, Nil)
}
__typedArg0 := Nil
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V395, tmp18603)
}
__typedArg0 := V395
__typedArg1 := tmp18603
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp18604 := Call(__e, ns2_1set, sym_5e_6, tmp18602)


_ = tmp18604

tmp18605 := MakeNative(func(__e *ControlFlow) {
V398 := __e.Get(1)
_ = V398
tmp18608 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V398)
}
__typedArg0 := Nil
__typedArg1 := V398
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp18608 {
tmp18606 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Nil, Nil)
}
__typedArg0 := Nil
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Nil, tmp18606)
}
__typedArg0 := Nil
__typedArg1 := tmp18606
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

__e.TailApply(ns2_1set, sym_5end_6, tmp18605)
return




}, 0)

