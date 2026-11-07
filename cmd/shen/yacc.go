package main

import . "github.com/pyrex41/shen-go/kl"

var YaccMain = MakeNative(func(__e *ControlFlow) {
tmp17538 := MakeNative(func(__e *ControlFlow) {
V112 := __e.Get(1)
_ = V112
V113 := __e.Get(2)
_ = V113
tmp17539 := Call(__e, V112, V113)


let__17400 := tmp17539
_ = let__17400

tmp17546 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17400)


if True == tmp17546 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("parse failure\n"))
}
__typedArg0 := MakeString("parse failure\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
} else {
tmp17544 := Call(__e, PrimFunc(symshen_4partial_1parse_1failure_2), let__17400)


if True == tmp17544 {
tmp17540 := Call(__e, PrimFunc(symshen_4in_1_6), let__17400)


tmp17541 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dresidue_d, tmp17540)
}
__typedArg0 := symshen_4_dresidue_d
__typedArg1 := tmp17540
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp17541

tmp17542 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dresidue_d)
}
__typedArg0 := symshen_4_dresidue_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4raise_1syntax_1error), tmp17542)
return


} else {
__e.TailApply(PrimFunc(symshen_4_5_1out), let__17400)
return
}


}


}, 2)

tmp17547 := Call(__e, ns2_1set, symcompile, tmp17538)


_ = tmp17547

tmp17548 := MakeNative(func(__e *ControlFlow) {
V115 := __e.Get(1)
_ = V115
tmp17549 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dmaximum_1print_1sequence_1size_d)
}
__typedArg0 := sym_dmaximum_1print_1sequence_1size_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp17550 := Call(__e, PrimFunc(symshen_4syntax_1error_1message), tmp17549, MakeInteger(0), V115)


tmp17552 := Call(__e, PrimFunc(symshen_4proc_1nl), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("syntax error here: "))
__typedS1, __typedOK1 := TypedString(tmp17550)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("syntax error here: ")
__typedArg1 := tmp17550
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(tmp17552)
}
__typedArg0 := tmp17552
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}, 1)

tmp17553 := Call(__e, ns2_1set, symshen_4raise_1syntax_1error, tmp17548)


_ = tmp17553

tmp17554 := MakeNative(func(__e *ControlFlow) {
V123 := __e.Get(1)
_ = V123
V124 := __e.Get(2)
_ = V124
V125 := __e.Get(3)
_ = V125
tmp17565 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V125)
}
__typedArg0 := Nil
__typedArg1 := V125
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17565 {
__e.Return(MakeString("\n"))
return
} else {
tmp17563 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V123, V124)
}
__typedArg0 := V123
__typedArg1 := V124
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17563 {
__e.Return(MakeString("...etc \n"))
return
} else {
tmp17561 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V125)
}
__typedArg0 := V125
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17561 {
tmp17555 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V125)
}
__typedArg0 := V125
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17556 := Call(__e, PrimFunc(symshen_4app), tmp17555, MakeString(" "), symshen_4s)


tmp17557 := (func() Obj {
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

tmp17558 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V125)
}
__typedArg0 := V125
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17559 := Call(__e, PrimFunc(symshen_4syntax_1error_1message), V123, tmp17557, tmp17558)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(tmp17556)
__typedS1, __typedOK1 := TypedString(tmp17559)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := tmp17556
__typedArg1 := tmp17559
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

tmp17566 := Call(__e, ns2_1set, symshen_4syntax_1error_1message, tmp17554)


_ = tmp17566

tmp17567 := MakeNative(func(__e *ControlFlow) {
V126 := __e.Get(1)
_ = V126
tmp17568 := Call(__e, PrimFunc(symfail))


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V126, tmp17568)
}
__typedArg0 := V126
__typedArg1 := tmp17568
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp17569 := Call(__e, ns2_1set, symshen_4parse_1failure_2, tmp17567)


_ = tmp17569

tmp17570 := MakeNative(func(__e *ControlFlow) {
V127 := __e.Get(1)
_ = V127
tmp17571 := Call(__e, PrimFunc(symshen_4in_1_6), V127)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp17571)
}
__typedArg0 := tmp17571
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})())
return


}, 1)

tmp17572 := Call(__e, ns2_1set, symshen_4partial_1parse_1failure_2, tmp17570)


_ = tmp17572

tmp17573 := MakeNative(func(__e *ControlFlow) {
V130 := __e.Get(1)
_ = V130
tmp17586 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V130)
}
__typedArg0 := V130
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17577 Obj

if True == tmp17586 {
tmp17584 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V130)
}
__typedArg0 := V130
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17585 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp17584)
}
__typedArg0 := tmp17584
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17579 Obj

if True == tmp17585 {
tmp17581 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V130)
}
__typedArg0 := V130
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17582 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp17581)
}
__typedArg0 := tmp17581
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17583 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp17582)
}
__typedArg0 := Nil
__typedArg1 := tmp17582
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17580 Obj

if True == tmp17583 {
ifres17580 = True


} else {
ifres17580 = False


}

ifres17579 = ifres17580


} else {
ifres17579 = False


}

var ifres17578 Obj

if True == ifres17579 {
ifres17578 = True


} else {
ifres17578 = False


}

ifres17577 = ifres17578


} else {
ifres17577 = False


}

if True == ifres17577 {
tmp17574 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V130)
}
__typedArg0 := V130
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp17574)
}
__typedArg0 := tmp17574
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return


} else {
tmp17575 := Call(__e, PrimFunc(symshen_4app), V130, MakeString(" is not a YACC stream\n"), symshen_4s)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(tmp17575)
}
__typedArg0 := tmp17575
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}


}, 1)

tmp17587 := Call(__e, ns2_1set, symshen_4objectcode, tmp17573)


_ = tmp17587

tmp17588 := MakeNative(func(__e *ControlFlow) {
V131 := __e.Get(1)
_ = V131
tmp17589 := MakeNative(func(__e *ControlFlow) {
Z132 := __e.Get(1)
_ = Z132
__e.TailApply(PrimFunc(symshen_4_5yacc_6), Z132)
return
}, 1)

__e.TailApply(PrimFunc(symcompile), tmp17589, V131)
return


}, 1)

tmp17590 := Call(__e, ns2_1set, symshen_4yacc_1_6shen, tmp17588)


_ = tmp17590

tmp17591 := MakeNative(func(__e *ControlFlow) {
V133 := __e.Get(1)
_ = V133
tmp17618 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V133)
}
__typedArg0 := V133
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17592 Obj

if True == tmp17618 {
tmp17593 := Call(__e, PrimFunc(symhead), V133)


let__17402 := tmp17593
_ = let__17402

tmp17594 := Call(__e, PrimFunc(symtail), V133)


let__17403 := tmp17594
_ = let__17403

tmp17595 := Call(__e, PrimFunc(symshen_4_5yaccsig_6), let__17403)


let__17404 := tmp17595
_ = let__17404

tmp17616 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17404)


var ifres17596 Obj

if True == tmp17616 {
tmp17597 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17596 = tmp17597


} else {
tmp17598 := Call(__e, PrimFunc(symshen_4_5_1out), let__17404)


let__17405 := tmp17598
_ = let__17405

tmp17599 := Call(__e, PrimFunc(symshen_4in_1_6), let__17404)


let__17406 := tmp17599
_ = let__17406

tmp17600 := Call(__e, PrimFunc(symshen_4_5c_1rules_6), let__17406)


let__17407 := tmp17600
_ = let__17407

tmp17615 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17407)


var ifres17601 Obj

if True == tmp17615 {
tmp17602 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17601 = tmp17602


} else {
tmp17603 := Call(__e, PrimFunc(symshen_4_5_1out), let__17407)


let__17408 := tmp17603
_ = let__17408

tmp17604 := Call(__e, PrimFunc(symshen_4in_1_6), let__17407)


let__17409 := tmp17604
_ = let__17409

tmp17605 := Call(__e, PrimFunc(symgensym), symS)


let__17410 := tmp17605
_ = let__17410

tmp17606 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17402, Nil)
}
__typedArg0 := let__17402
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17607 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefine, tmp17606)
}
__typedArg0 := symdefine
__typedArg1 := tmp17606
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17608 := Call(__e, PrimFunc(symshen_4c_1rules_1_6shen), let__17405, let__17410, let__17408)


tmp17609 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17608, Nil)
}
__typedArg0 := tmp17608
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17610 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_6, tmp17609)
}
__typedArg0 := sym_1_6
__typedArg1 := tmp17609
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17611 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17410, tmp17610)
}
__typedArg0 := let__17410
__typedArg1 := tmp17610
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17612 := Call(__e, PrimFunc(symappend), let__17405, tmp17611)


tmp17613 := Call(__e, PrimFunc(symappend), tmp17607, tmp17612)


let__17411 := tmp17613
_ = let__17411

tmp17614 := Call(__e, PrimFunc(symshen_4comb), let__17409, let__17411)


ifres17601 = tmp17614


}

ifres17596 = ifres17601


}

ifres17592 = ifres17596


} else {
tmp17617 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17592 = tmp17617


}

let__17401 := ifres17592
_ = let__17401

tmp17620 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17401)


if True == tmp17620 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__17401)
return
}


}, 1)

tmp17621 := Call(__e, ns2_1set, symshen_4_5yacc_6, tmp17591)


_ = tmp17621

tmp17622 := MakeNative(func(__e *ControlFlow) {
V145 := __e.Get(1)
_ = V145
tmp17679 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V145)
}
__typedArg0 := V145
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17623 Obj

if True == tmp17679 {
tmp17624 := Call(__e, PrimFunc(symhead), V145)


let__17413 := tmp17624
_ = let__17413

tmp17625 := Call(__e, PrimFunc(symtail), V145)


let__17414 := tmp17625
_ = let__17414

tmp17677 := Call(__e, PrimFunc(symshen_4ccons_2), let__17414)


var ifres17626 Obj

if True == tmp17677 {
tmp17627 := Call(__e, PrimFunc(symhead), let__17414)


let__17415 := tmp17627
_ = let__17415

tmp17628 := Call(__e, PrimFunc(symtail), let__17414)


let__17416 := tmp17628
_ = let__17416

tmp17675 := Call(__e, PrimFunc(symshen_4hds_a_2), let__17415, symlist)


var ifres17629 Obj

if True == tmp17675 {
tmp17630 := Call(__e, PrimFunc(symtail), let__17415)


let__17417 := tmp17630
_ = let__17417

tmp17673 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__17417)
}
__typedArg0 := let__17417
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17631 Obj

if True == tmp17673 {
tmp17632 := Call(__e, PrimFunc(symhead), let__17417)


let__17418 := tmp17632
_ = let__17418

tmp17633 := Call(__e, PrimFunc(symtail), let__17417)


let__17419 := tmp17633
_ = let__17419

tmp17634 := Call(__e, PrimFunc(sym_5end_6), let__17419)


let__17420 := tmp17634
_ = let__17420

tmp17671 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17420)


var ifres17635 Obj

if True == tmp17671 {
tmp17636 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17635 = tmp17636


} else {
tmp17637 := Call(__e, PrimFunc(symshen_4in_1_6), let__17420)


let__17421 := tmp17637
_ = let__17421

tmp17670 := Call(__e, PrimFunc(symshen_4hds_a_2), let__17416, sym_a_a_6)


var ifres17638 Obj

if True == tmp17670 {
tmp17639 := Call(__e, PrimFunc(symtail), let__17416)


let__17422 := tmp17639
_ = let__17422

tmp17668 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__17422)
}
__typedArg0 := let__17422
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17640 Obj

if True == tmp17668 {
tmp17641 := Call(__e, PrimFunc(symhead), let__17422)


let__17423 := tmp17641
_ = let__17423

tmp17642 := Call(__e, PrimFunc(symtail), let__17422)


let__17424 := tmp17642
_ = let__17424

tmp17666 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__17424)
}
__typedArg0 := let__17424
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17643 Obj

if True == tmp17666 {
tmp17644 := Call(__e, PrimFunc(symhead), let__17424)


let__17425 := tmp17644
_ = let__17425

tmp17645 := Call(__e, PrimFunc(symtail), let__17424)


let__17426 := tmp17645
_ = let__17426

tmp17664 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_i, let__17413)
}
__typedArg0 := sym_i
__typedArg1 := let__17413
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17661 Obj

if True == tmp17664 {
tmp17663 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_j, let__17425)
}
__typedArg0 := sym_j
__typedArg1 := let__17425
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17662 Obj

if True == tmp17663 {
ifres17662 = True


} else {
ifres17662 = False


}

ifres17661 = ifres17662


} else {
ifres17661 = False


}

var ifres17646 Obj

if True == ifres17661 {
tmp17647 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17418, Nil)
}
__typedArg0 := let__17418
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17648 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp17647)
}
__typedArg0 := symlist
__typedArg1 := tmp17647
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17649 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17418, Nil)
}
__typedArg0 := let__17418
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17650 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp17649)
}
__typedArg0 := symlist
__typedArg1 := tmp17649
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17651 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17423, Nil)
}
__typedArg0 := let__17423
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17652 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17650, tmp17651)
}
__typedArg0 := tmp17650
__typedArg1 := tmp17651
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17653 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstr, tmp17652)
}
__typedArg0 := symstr
__typedArg1 := tmp17652
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17654 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_j, Nil)
}
__typedArg0 := sym_j
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17655 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17653, tmp17654)
}
__typedArg0 := tmp17653
__typedArg1 := tmp17654
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17656 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp17655)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp17655
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17657 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17648, tmp17656)
}
__typedArg0 := tmp17648
__typedArg1 := tmp17656
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17658 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_i, tmp17657)
}
__typedArg0 := sym_i
__typedArg1 := tmp17657
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17659 := Call(__e, PrimFunc(symshen_4comb), let__17426, tmp17658)


ifres17646 = tmp17659


} else {
tmp17660 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17646 = tmp17660


}

ifres17643 = ifres17646


} else {
tmp17665 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17643 = tmp17665


}

ifres17640 = ifres17643


} else {
tmp17667 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17640 = tmp17667


}

ifres17638 = ifres17640


} else {
tmp17669 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17638 = tmp17669


}

ifres17635 = ifres17638


}

ifres17631 = ifres17635


} else {
tmp17672 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17631 = tmp17672


}

ifres17629 = ifres17631


} else {
tmp17674 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17629 = tmp17674


}

ifres17626 = ifres17629


} else {
tmp17676 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17626 = tmp17676


}

ifres17623 = ifres17626


} else {
tmp17678 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17623 = tmp17678


}

let__17412 := ifres17623
_ = let__17412

tmp17689 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17412)


if True == tmp17689 {
tmp17680 := Call(__e, PrimFunc(sym_5e_6), V145)


let__17428 := tmp17680
_ = let__17428

tmp17685 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17428)


var ifres17681 Obj

if True == tmp17685 {
tmp17682 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17681 = tmp17682


} else {
tmp17683 := Call(__e, PrimFunc(symshen_4in_1_6), let__17428)


let__17429 := tmp17683
_ = let__17429

tmp17684 := Call(__e, PrimFunc(symshen_4comb), let__17429, Nil)


ifres17681 = tmp17684


}

let__17427 := ifres17681
_ = let__17427

tmp17687 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17427)


if True == tmp17687 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__17427)
return
}


} else {
__e.Return(let__17412)
return
}


}, 1)

tmp17690 := Call(__e, ns2_1set, symshen_4_5yaccsig_6, tmp17622)


_ = tmp17690

tmp17691 := MakeNative(func(__e *ControlFlow) {
V164 := __e.Get(1)
_ = V164
tmp17692 := Call(__e, PrimFunc(symshen_4_5c_1rule_6), V164)


let__17431 := tmp17692
_ = let__17431

tmp17705 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17431)


var ifres17693 Obj

if True == tmp17705 {
tmp17694 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17693 = tmp17694


} else {
tmp17695 := Call(__e, PrimFunc(symshen_4_5_1out), let__17431)


let__17432 := tmp17695
_ = let__17432

tmp17696 := Call(__e, PrimFunc(symshen_4in_1_6), let__17431)


let__17433 := tmp17696
_ = let__17433

tmp17697 := Call(__e, PrimFunc(symshen_4_5c_1rules_6), let__17433)


let__17434 := tmp17697
_ = let__17434

tmp17704 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17434)


var ifres17698 Obj

if True == tmp17704 {
tmp17699 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17698 = tmp17699


} else {
tmp17700 := Call(__e, PrimFunc(symshen_4_5_1out), let__17434)


let__17435 := tmp17700
_ = let__17435

tmp17701 := Call(__e, PrimFunc(symshen_4in_1_6), let__17434)


let__17436 := tmp17701
_ = let__17436

tmp17702 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17432, let__17435)
}
__typedArg0 := let__17432
__typedArg1 := let__17435
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17703 := Call(__e, PrimFunc(symshen_4comb), let__17436, tmp17702)


ifres17698 = tmp17703


}

ifres17693 = ifres17698


}

let__17430 := ifres17693
_ = let__17430

tmp17721 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17430)


if True == tmp17721 {
tmp17706 := Call(__e, PrimFunc(sym_5_b_6), V164)


let__17438 := tmp17706
_ = let__17438

tmp17717 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17438)


var ifres17707 Obj

if True == tmp17717 {
tmp17708 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17707 = tmp17708


} else {
tmp17709 := Call(__e, PrimFunc(symshen_4_5_1out), let__17438)


let__17439 := tmp17709
_ = let__17439

tmp17710 := Call(__e, PrimFunc(symshen_4in_1_6), let__17438)


let__17440 := tmp17710
_ = let__17440

tmp17715 := Call(__e, PrimFunc(symempty_2), let__17439)


var ifres17711 Obj

if True == tmp17715 {
ifres17711 = Nil


} else {
tmp17712 := Call(__e, PrimFunc(symshen_4app), let__17439, MakeString("\n ..."), symshen_4r)


tmp17714 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("YACC syntax error here:\n "))
__typedS1, __typedOK1 := TypedString(tmp17712)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("YACC syntax error here:\n ")
__typedArg1 := tmp17712
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("YACC syntax error here:\n "))
__typedS1, __typedOK1 := TypedString(tmp17712)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("YACC syntax error here:\n ")
__typedArg1 := tmp17712
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})()

ifres17711 = tmp17714


}

tmp17716 := Call(__e, PrimFunc(symshen_4comb), let__17440, ifres17711)


ifres17707 = tmp17716


}

let__17437 := ifres17707
_ = let__17437

tmp17719 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17437)


if True == tmp17719 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__17437)
return
}


} else {
__e.Return(let__17430)
return
}


}, 1)

tmp17722 := Call(__e, ns2_1set, symshen_4_5c_1rules_6, tmp17691)


_ = tmp17722

tmp17723 := MakeNative(func(__e *ControlFlow) {
V176 := __e.Get(1)
_ = V176
tmp17724 := Call(__e, PrimFunc(symshen_4_5syntax_6), V176)


let__17442 := tmp17724
_ = let__17442

tmp17743 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17442)


var ifres17725 Obj

if True == tmp17743 {
tmp17726 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17725 = tmp17726


} else {
tmp17727 := Call(__e, PrimFunc(symshen_4_5_1out), let__17442)


let__17443 := tmp17727
_ = let__17443

tmp17728 := Call(__e, PrimFunc(symshen_4in_1_6), let__17442)


let__17444 := tmp17728
_ = let__17444

tmp17729 := Call(__e, PrimFunc(symshen_4_5semantics_6), let__17444)


let__17445 := tmp17729
_ = let__17445

tmp17742 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17445)


var ifres17730 Obj

if True == tmp17742 {
tmp17731 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17730 = tmp17731


} else {
tmp17732 := Call(__e, PrimFunc(symshen_4_5_1out), let__17445)


let__17446 := tmp17732
_ = let__17446

tmp17733 := Call(__e, PrimFunc(symshen_4in_1_6), let__17445)


let__17447 := tmp17733
_ = let__17447

tmp17734 := Call(__e, PrimFunc(symshen_4_5sc_6), let__17447)


let__17448 := tmp17734
_ = let__17448

tmp17741 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17448)


var ifres17735 Obj

if True == tmp17741 {
tmp17736 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17735 = tmp17736


} else {
tmp17737 := Call(__e, PrimFunc(symshen_4in_1_6), let__17448)


let__17449 := tmp17737
_ = let__17449

tmp17738 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17446, Nil)
}
__typedArg0 := let__17446
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17739 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17443, tmp17738)
}
__typedArg0 := let__17443
__typedArg1 := tmp17738
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17740 := Call(__e, PrimFunc(symshen_4comb), let__17449, tmp17739)


ifres17735 = tmp17740


}

ifres17730 = ifres17735


}

ifres17725 = ifres17730


}

let__17441 := ifres17725
_ = let__17441

tmp17762 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17441)


if True == tmp17762 {
tmp17744 := Call(__e, PrimFunc(symshen_4_5syntax_6), V176)


let__17451 := tmp17744
_ = let__17451

tmp17758 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17451)


var ifres17745 Obj

if True == tmp17758 {
tmp17746 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17745 = tmp17746


} else {
tmp17747 := Call(__e, PrimFunc(symshen_4_5_1out), let__17451)


let__17452 := tmp17747
_ = let__17452

tmp17748 := Call(__e, PrimFunc(symshen_4in_1_6), let__17451)


let__17453 := tmp17748
_ = let__17453

tmp17749 := Call(__e, PrimFunc(symshen_4_5sc_6), let__17453)


let__17454 := tmp17749
_ = let__17454

tmp17757 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17454)


var ifres17750 Obj

if True == tmp17757 {
tmp17751 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17750 = tmp17751


} else {
tmp17752 := Call(__e, PrimFunc(symshen_4in_1_6), let__17454)


let__17455 := tmp17752
_ = let__17455

tmp17753 := Call(__e, PrimFunc(symshen_4autocomplete), let__17452)


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
return PrimCons(let__17452, tmp17754)
}
__typedArg0 := let__17452
__typedArg1 := tmp17754
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17756 := Call(__e, PrimFunc(symshen_4comb), let__17455, tmp17755)


ifres17750 = tmp17756


}

ifres17745 = ifres17750


}

let__17450 := ifres17745
_ = let__17450

tmp17760 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17450)


if True == tmp17760 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__17450)
return
}


} else {
__e.Return(let__17441)
return
}


}, 1)

tmp17763 := Call(__e, ns2_1set, symshen_4_5c_1rule_6, tmp17723)


_ = tmp17763

tmp17764 := MakeNative(func(__e *ControlFlow) {
V192 := __e.Get(1)
_ = V192
tmp17793 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V192)
}
__typedArg0 := V192
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17785 Obj

if True == tmp17793 {
tmp17791 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V192)
}
__typedArg0 := V192
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17792 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp17791)
}
__typedArg0 := Nil
__typedArg1 := tmp17791
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17787 Obj

if True == tmp17792 {
tmp17789 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V192)
}
__typedArg0 := V192
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17790 := Call(__e, PrimFunc(symshen_4non_1terminal_2), tmp17789)


var ifres17788 Obj

if True == tmp17790 {
ifres17788 = True


} else {
ifres17788 = False


}

ifres17787 = ifres17788


} else {
ifres17787 = False


}

var ifres17786 Obj

if True == ifres17787 {
ifres17786 = True


} else {
ifres17786 = False


}

ifres17785 = ifres17786


} else {
ifres17785 = False


}

if True == ifres17785 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V192)
}
__typedArg0 := V192
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return
} else {
tmp17783 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V192)
}
__typedArg0 := V192
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17779 Obj

if True == tmp17783 {
tmp17781 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V192)
}
__typedArg0 := V192
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17782 := Call(__e, PrimFunc(symshen_4non_1terminal_2), tmp17781)


var ifres17780 Obj

if True == tmp17782 {
ifres17780 = True


} else {
ifres17780 = False


}

ifres17779 = ifres17780


} else {
ifres17779 = False


}

if True == ifres17779 {
tmp17765 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V192)
}
__typedArg0 := V192
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17766 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V192)
}
__typedArg0 := V192
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17767 := Call(__e, PrimFunc(symshen_4autocomplete), tmp17766)


tmp17768 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17767, Nil)
}
__typedArg0 := tmp17767
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17769 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17765, tmp17768)
}
__typedArg0 := tmp17765
__typedArg1 := tmp17768
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symappend, tmp17769)
}
__typedArg0 := symappend
__typedArg1 := tmp17769
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp17777 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V192)
}
__typedArg0 := V192
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17777 {
tmp17770 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V192)
}
__typedArg0 := V192
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17771 := Call(__e, PrimFunc(symshen_4autocomplete), tmp17770)


tmp17772 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V192)
}
__typedArg0 := V192
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17773 := Call(__e, PrimFunc(symshen_4autocomplete), tmp17772)


tmp17774 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17773, Nil)
}
__typedArg0 := tmp17773
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17775 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17771, tmp17774)
}
__typedArg0 := tmp17771
__typedArg1 := tmp17774
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp17775)
}
__typedArg0 := symcons
__typedArg1 := tmp17775
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

tmp17794 := Call(__e, ns2_1set, symshen_4autocomplete, tmp17764)


_ = tmp17794

tmp17795 := MakeNative(func(__e *ControlFlow) {
V193 := __e.Get(1)
_ = V193
tmp17801 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsymbol_2) {
return PrimIsSymbol(V193)
}
__typedArg0 := V193
return Call(__e, PrimFunc(symsymbol_2), __typedArg0)
})()

if True == tmp17801 {
tmp17797 := Call(__e, PrimFunc(symexplode), V193)


let__17456 := tmp17797
_ = let__17456

tmp17798 := MakeNative(func(__e *ControlFlow) {
Z195 := __e.Get(1)
_ = Z195
__e.TailApply(PrimFunc(symshen_4_5non_1terminal_2_6), Z195)
return
}, 1)

tmp17799 := Call(__e, PrimFunc(symcompile), tmp17798, let__17456)


if True == tmp17799 {
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

tmp17802 := Call(__e, ns2_1set, symshen_4non_1terminal_2, tmp17795)


_ = tmp17802

tmp17803 := MakeNative(func(__e *ControlFlow) {
V196 := __e.Get(1)
_ = V196
tmp17804 := Call(__e, PrimFunc(symshen_4_5packagenames_6), V196)


let__17458 := tmp17804
_ = let__17458

tmp17814 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17458)


var ifres17805 Obj

if True == tmp17814 {
tmp17806 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17805 = tmp17806


} else {
tmp17807 := Call(__e, PrimFunc(symshen_4in_1_6), let__17458)


let__17459 := tmp17807
_ = let__17459

tmp17808 := Call(__e, PrimFunc(symshen_4_5non_1terminal_1name_6), let__17459)


let__17460 := tmp17808
_ = let__17460

tmp17813 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17460)


var ifres17809 Obj

if True == tmp17813 {
tmp17810 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17809 = tmp17810


} else {
tmp17811 := Call(__e, PrimFunc(symshen_4in_1_6), let__17460)


let__17461 := tmp17811
_ = let__17461

tmp17812 := Call(__e, PrimFunc(symshen_4comb), let__17461, True)


ifres17809 = tmp17812


}

ifres17805 = ifres17809


}

let__17457 := ifres17805
_ = let__17457

tmp17832 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17457)


if True == tmp17832 {
tmp17815 := Call(__e, PrimFunc(symshen_4_5non_1terminal_1name_6), V196)


let__17463 := tmp17815
_ = let__17463

tmp17820 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17463)


var ifres17816 Obj

if True == tmp17820 {
tmp17817 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17816 = tmp17817


} else {
tmp17818 := Call(__e, PrimFunc(symshen_4in_1_6), let__17463)


let__17464 := tmp17818
_ = let__17464

tmp17819 := Call(__e, PrimFunc(symshen_4comb), let__17464, True)


ifres17816 = tmp17819


}

let__17462 := ifres17816
_ = let__17462

tmp17830 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17462)


if True == tmp17830 {
tmp17821 := Call(__e, PrimFunc(sym_5_b_6), V196)


let__17466 := tmp17821
_ = let__17466

tmp17826 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17466)


var ifres17822 Obj

if True == tmp17826 {
tmp17823 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17822 = tmp17823


} else {
tmp17824 := Call(__e, PrimFunc(symshen_4in_1_6), let__17466)


let__17467 := tmp17824
_ = let__17467

tmp17825 := Call(__e, PrimFunc(symshen_4comb), let__17467, False)


ifres17822 = tmp17825


}

let__17465 := ifres17822
_ = let__17465

tmp17828 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17465)


if True == tmp17828 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__17465)
return
}


} else {
__e.Return(let__17462)
return
}


} else {
__e.Return(let__17457)
return
}


}, 1)

tmp17833 := Call(__e, ns2_1set, symshen_4_5non_1terminal_2_6, tmp17803)


_ = tmp17833

tmp17834 := MakeNative(func(__e *ControlFlow) {
V208 := __e.Get(1)
_ = V208
tmp17835 := Call(__e, PrimFunc(symshen_4_5packagename_6), V208)


let__17469 := tmp17835
_ = let__17469

tmp17849 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17469)


var ifres17836 Obj

if True == tmp17849 {
tmp17837 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17836 = tmp17837


} else {
tmp17838 := Call(__e, PrimFunc(symshen_4in_1_6), let__17469)


let__17470 := tmp17838
_ = let__17470

tmp17848 := Call(__e, PrimFunc(symshen_4hds_a_2), let__17470, MakeString("."))


var ifres17839 Obj

if True == tmp17848 {
tmp17840 := Call(__e, PrimFunc(symtail), let__17470)


let__17471 := tmp17840
_ = let__17471

tmp17841 := Call(__e, PrimFunc(symshen_4_5packagenames_6), let__17471)


let__17472 := tmp17841
_ = let__17472

tmp17846 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17472)


var ifres17842 Obj

if True == tmp17846 {
tmp17843 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17842 = tmp17843


} else {
tmp17844 := Call(__e, PrimFunc(symshen_4in_1_6), let__17472)


let__17473 := tmp17844
_ = let__17473

tmp17845 := Call(__e, PrimFunc(symshen_4comb), let__17473, symshen_4skip)


ifres17842 = tmp17845


}

ifres17839 = ifres17842


} else {
tmp17847 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17839 = tmp17847


}

ifres17836 = ifres17839


}

let__17468 := ifres17836
_ = let__17468

tmp17863 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17468)


if True == tmp17863 {
tmp17850 := Call(__e, PrimFunc(symshen_4_5packagename_6), V208)


let__17475 := tmp17850
_ = let__17475

tmp17859 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17475)


var ifres17851 Obj

if True == tmp17859 {
tmp17852 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17851 = tmp17852


} else {
tmp17853 := Call(__e, PrimFunc(symshen_4in_1_6), let__17475)


let__17476 := tmp17853
_ = let__17476

tmp17858 := Call(__e, PrimFunc(symshen_4hds_a_2), let__17476, MakeString("."))


var ifres17854 Obj

if True == tmp17858 {
tmp17855 := Call(__e, PrimFunc(symtail), let__17476)


let__17477 := tmp17855
_ = let__17477

tmp17856 := Call(__e, PrimFunc(symshen_4comb), let__17477, symshen_4skip)


ifres17854 = tmp17856


} else {
tmp17857 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17854 = tmp17857


}

ifres17851 = ifres17854


}

let__17474 := ifres17851
_ = let__17474

tmp17861 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17474)


if True == tmp17861 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__17474)
return
}


} else {
__e.Return(let__17468)
return
}


}, 1)

tmp17864 := Call(__e, ns2_1set, symshen_4_5packagenames_6, tmp17834)


_ = tmp17864

tmp17865 := MakeNative(func(__e *ControlFlow) {
V219 := __e.Get(1)
_ = V219
tmp17866 := Call(__e, PrimFunc(symshen_4_5packagechar_6), V219)


let__17479 := tmp17866
_ = let__17479

tmp17876 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17479)


var ifres17867 Obj

if True == tmp17876 {
tmp17868 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17867 = tmp17868


} else {
tmp17869 := Call(__e, PrimFunc(symshen_4in_1_6), let__17479)


let__17480 := tmp17869
_ = let__17480

tmp17870 := Call(__e, PrimFunc(symshen_4_5packagename_6), let__17480)


let__17481 := tmp17870
_ = let__17481

tmp17875 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17481)


var ifres17871 Obj

if True == tmp17875 {
tmp17872 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17871 = tmp17872


} else {
tmp17873 := Call(__e, PrimFunc(symshen_4in_1_6), let__17481)


let__17482 := tmp17873
_ = let__17482

tmp17874 := Call(__e, PrimFunc(symshen_4comb), let__17482, symshen_4skip)


ifres17871 = tmp17874


}

ifres17867 = ifres17871


}

let__17478 := ifres17867
_ = let__17478

tmp17886 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17478)


if True == tmp17886 {
tmp17877 := Call(__e, PrimFunc(sym_5e_6), V219)


let__17484 := tmp17877
_ = let__17484

tmp17882 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17484)


var ifres17878 Obj

if True == tmp17882 {
tmp17879 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17878 = tmp17879


} else {
tmp17880 := Call(__e, PrimFunc(symshen_4in_1_6), let__17484)


let__17485 := tmp17880
_ = let__17485

tmp17881 := Call(__e, PrimFunc(symshen_4comb), let__17485, symshen_4skip)


ifres17878 = tmp17881


}

let__17483 := ifres17878
_ = let__17483

tmp17884 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17483)


if True == tmp17884 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__17483)
return
}


} else {
__e.Return(let__17478)
return
}


}, 1)

tmp17887 := Call(__e, ns2_1set, symshen_4_5packagename_6, tmp17865)


_ = tmp17887

tmp17888 := MakeNative(func(__e *ControlFlow) {
V228 := __e.Get(1)
_ = V228
tmp17898 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V228)
}
__typedArg0 := V228
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17889 Obj

if True == tmp17898 {
tmp17890 := Call(__e, PrimFunc(symhead), V228)


let__17487 := tmp17890
_ = let__17487

tmp17891 := Call(__e, PrimFunc(symtail), V228)


let__17488 := tmp17891
_ = let__17488

tmp17895 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__17487, MakeString("."))
}
__typedArg0 := let__17487
__typedArg1 := MakeString(".")
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

tmp17896 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp17895)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp17895
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres17892 Obj

if True == tmp17896 {
tmp17893 := Call(__e, PrimFunc(symshen_4comb), let__17488, symshen_4skip)


ifres17892 = tmp17893


} else {
tmp17894 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17892 = tmp17894


}

ifres17889 = ifres17892


} else {
tmp17897 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17889 = tmp17897


}

let__17486 := ifres17889
_ = let__17486

tmp17900 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17486)


if True == tmp17900 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__17486)
return
}


}, 1)

tmp17901 := Call(__e, ns2_1set, symshen_4_5packagechar_6, tmp17888)


_ = tmp17901

tmp17902 := MakeNative(func(__e *ControlFlow) {
V232 := __e.Get(1)
_ = V232
tmp17921 := Call(__e, PrimFunc(symshen_4hds_a_2), V232, MakeString("<"))


var ifres17903 Obj

if True == tmp17921 {
tmp17904 := Call(__e, PrimFunc(symtail), V232)


let__17490 := tmp17904
_ = let__17490

tmp17905 := Call(__e, PrimFunc(sym_5_b_6), let__17490)


let__17491 := tmp17905
_ = let__17491

tmp17919 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17491)


var ifres17906 Obj

if True == tmp17919 {
tmp17907 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17906 = tmp17907


} else {
tmp17908 := Call(__e, PrimFunc(symshen_4_5_1out), let__17491)


let__17492 := tmp17908
_ = let__17492

tmp17909 := Call(__e, PrimFunc(symshen_4in_1_6), let__17491)


let__17493 := tmp17909
_ = let__17493

tmp17913 := Call(__e, PrimFunc(symreverse), let__17492)


let__17494 := tmp17913
_ = let__17494

tmp17918 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__17494)
}
__typedArg0 := let__17494
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17914 Obj

if True == tmp17918 {
tmp17916 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__17494)
}
__typedArg0 := let__17494
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17917 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp17916, MakeString(">"))
}
__typedArg0 := tmp17916
__typedArg1 := MakeString(">")
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17915 Obj

if True == tmp17917 {
ifres17915 = True


} else {
ifres17915 = False


}

ifres17914 = ifres17915


} else {
ifres17914 = False


}

var ifres17910 Obj

if True == ifres17914 {
tmp17911 := Call(__e, PrimFunc(symshen_4comb), let__17493, symshen_4skip)


ifres17910 = tmp17911


} else {
tmp17912 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17910 = tmp17912


}

ifres17906 = ifres17910


}

ifres17903 = ifres17906


} else {
tmp17920 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17903 = tmp17920


}

let__17489 := ifres17903
_ = let__17489

tmp17923 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17489)


if True == tmp17923 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__17489)
return
}


}, 1)

tmp17924 := Call(__e, ns2_1set, symshen_4_5non_1terminal_1name_6, tmp17902)


_ = tmp17924

tmp17925 := MakeNative(func(__e *ControlFlow) {
V239 := __e.Get(1)
_ = V239
tmp17926 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(";"))
}
__typedArg0 := MakeString(";")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V239, tmp17926)
}
__typedArg0 := V239
__typedArg1 := tmp17926
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp17927 := Call(__e, ns2_1set, symshen_4semicolon_2, tmp17925)


_ = tmp17927

tmp17928 := MakeNative(func(__e *ControlFlow) {
V240 := __e.Get(1)
_ = V240
tmp17937 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V240)
}
__typedArg0 := V240
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17929 Obj

if True == tmp17937 {
tmp17930 := Call(__e, PrimFunc(symhead), V240)


let__17496 := tmp17930
_ = let__17496

tmp17931 := Call(__e, PrimFunc(symtail), V240)


let__17497 := tmp17931
_ = let__17497

tmp17935 := Call(__e, PrimFunc(symshen_4colon_1equal_2), let__17496)


var ifres17932 Obj

if True == tmp17935 {
tmp17933 := Call(__e, PrimFunc(symshen_4comb), let__17497, symshen_4skip)


ifres17932 = tmp17933


} else {
tmp17934 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17932 = tmp17934


}

ifres17929 = ifres17932


} else {
tmp17936 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17929 = tmp17936


}

let__17495 := ifres17929
_ = let__17495

tmp17939 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17495)


if True == tmp17939 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__17495)
return
}


}, 1)

tmp17940 := Call(__e, ns2_1set, symshen_4_5colon_1equal_6, tmp17928)


_ = tmp17940

tmp17941 := MakeNative(func(__e *ControlFlow) {
V244 := __e.Get(1)
_ = V244
tmp17942 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":="))
}
__typedArg0 := MakeString(":=")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp17942, V244)
}
__typedArg0 := tmp17942
__typedArg1 := V244
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp17943 := Call(__e, ns2_1set, symshen_4colon_1equal_2, tmp17941)


_ = tmp17943

tmp17944 := MakeNative(func(__e *ControlFlow) {
V245 := __e.Get(1)
_ = V245
tmp17945 := Call(__e, PrimFunc(symshen_4_5syntax_1item_6), V245)


let__17499 := tmp17945
_ = let__17499

tmp17958 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17499)


var ifres17946 Obj

if True == tmp17958 {
tmp17947 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17946 = tmp17947


} else {
tmp17948 := Call(__e, PrimFunc(symshen_4_5_1out), let__17499)


let__17500 := tmp17948
_ = let__17500

tmp17949 := Call(__e, PrimFunc(symshen_4in_1_6), let__17499)


let__17501 := tmp17949
_ = let__17501

tmp17950 := Call(__e, PrimFunc(symshen_4_5syntax_6), let__17501)


let__17502 := tmp17950
_ = let__17502

tmp17957 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17502)


var ifres17951 Obj

if True == tmp17957 {
tmp17952 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17951 = tmp17952


} else {
tmp17953 := Call(__e, PrimFunc(symshen_4_5_1out), let__17502)


let__17503 := tmp17953
_ = let__17503

tmp17954 := Call(__e, PrimFunc(symshen_4in_1_6), let__17502)


let__17504 := tmp17954
_ = let__17504

tmp17955 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17500, let__17503)
}
__typedArg0 := let__17500
__typedArg1 := let__17503
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17956 := Call(__e, PrimFunc(symshen_4comb), let__17504, tmp17955)


ifres17951 = tmp17956


}

ifres17946 = ifres17951


}

let__17498 := ifres17946
_ = let__17498

tmp17970 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17498)


if True == tmp17970 {
tmp17959 := Call(__e, PrimFunc(symshen_4_5syntax_1item_6), V245)


let__17506 := tmp17959
_ = let__17506

tmp17966 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17506)


var ifres17960 Obj

if True == tmp17966 {
tmp17961 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17960 = tmp17961


} else {
tmp17962 := Call(__e, PrimFunc(symshen_4_5_1out), let__17506)


let__17507 := tmp17962
_ = let__17507

tmp17963 := Call(__e, PrimFunc(symshen_4in_1_6), let__17506)


let__17508 := tmp17963
_ = let__17508

tmp17964 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17507, Nil)
}
__typedArg0 := let__17507
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17965 := Call(__e, PrimFunc(symshen_4comb), let__17508, tmp17964)


ifres17960 = tmp17965


}

let__17505 := ifres17960
_ = let__17505

tmp17968 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17505)


if True == tmp17968 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__17505)
return
}


} else {
__e.Return(let__17498)
return
}


}, 1)

tmp17971 := Call(__e, ns2_1set, symshen_4_5syntax_6, tmp17944)


_ = tmp17971

tmp17972 := MakeNative(func(__e *ControlFlow) {
V257 := __e.Get(1)
_ = V257
tmp17981 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V257)
}
__typedArg0 := V257
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17973 Obj

if True == tmp17981 {
tmp17974 := Call(__e, PrimFunc(symhead), V257)


let__17510 := tmp17974
_ = let__17510

tmp17975 := Call(__e, PrimFunc(symtail), V257)


let__17511 := tmp17975
_ = let__17511

tmp17979 := Call(__e, PrimFunc(symshen_4syntax_1item_2), let__17510)


var ifres17976 Obj

if True == tmp17979 {
tmp17977 := Call(__e, PrimFunc(symshen_4comb), let__17511, let__17510)


ifres17976 = tmp17977


} else {
tmp17978 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17976 = tmp17978


}

ifres17973 = ifres17976


} else {
tmp17980 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17973 = tmp17980


}

let__17509 := ifres17973
_ = let__17509

tmp17983 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17509)


if True == tmp17983 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__17509)
return
}


}, 1)

tmp17984 := Call(__e, ns2_1set, symshen_4_5syntax_1item_6, tmp17972)


_ = tmp17984

tmp17985 := MakeNative(func(__e *ControlFlow) {
V263 := __e.Get(1)
_ = V263
tmp18021 := Call(__e, PrimFunc(symshen_4colon_1equal_2), V263)


if True == tmp18021 {
__e.Return(False)
return
} else {
tmp18019 := Call(__e, PrimFunc(symshen_4semicolon_2), V263)


if True == tmp18019 {
__e.Return(False)
return
} else {
tmp18017 := Call(__e, PrimFunc(symatom_2), V263)


if True == tmp18017 {
__e.Return(True)
return
} else {
tmp18015 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V263)
}
__typedArg0 := V263
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17996 Obj

if True == tmp18015 {
tmp18013 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V263)
}
__typedArg0 := V263
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18014 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcons, tmp18013)
}
__typedArg0 := symcons
__typedArg1 := tmp18013
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17998 Obj

if True == tmp18014 {
tmp18011 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V263)
}
__typedArg0 := V263
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18012 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18011)
}
__typedArg0 := tmp18011
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18000 Obj

if True == tmp18012 {
tmp18008 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V263)
}
__typedArg0 := V263
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18009 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18008)
}
__typedArg0 := tmp18008
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18010 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18009)
}
__typedArg0 := tmp18009
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18002 Obj

if True == tmp18010 {
tmp18004 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V263)
}
__typedArg0 := V263
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18005 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18004)
}
__typedArg0 := tmp18004
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18006 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18005)
}
__typedArg0 := tmp18005
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18007 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp18006)
}
__typedArg0 := Nil
__typedArg1 := tmp18006
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18003 Obj

if True == tmp18007 {
ifres18003 = True


} else {
ifres18003 = False


}

ifres18002 = ifres18003


} else {
ifres18002 = False


}

var ifres18001 Obj

if True == ifres18002 {
ifres18001 = True


} else {
ifres18001 = False


}

ifres18000 = ifres18001


} else {
ifres18000 = False


}

var ifres17999 Obj

if True == ifres18000 {
ifres17999 = True


} else {
ifres17999 = False


}

ifres17998 = ifres17999


} else {
ifres17998 = False


}

var ifres17997 Obj

if True == ifres17998 {
ifres17997 = True


} else {
ifres17997 = False


}

ifres17996 = ifres17997


} else {
ifres17996 = False


}

if True == ifres17996 {
tmp17992 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V263)
}
__typedArg0 := V263
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17993 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp17992)
}
__typedArg0 := tmp17992
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17994 := Call(__e, PrimFunc(symshen_4syntax_1item_2), tmp17993)


if True == tmp17994 {
tmp17987 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V263)
}
__typedArg0 := V263
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17988 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp17987)
}
__typedArg0 := tmp17987
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17989 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp17988)
}
__typedArg0 := tmp17988
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17990 := Call(__e, PrimFunc(symshen_4syntax_1item_2), tmp17989)


if True == tmp17990 {
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

tmp18022 := Call(__e, ns2_1set, symshen_4syntax_1item_2, tmp17985)


_ = tmp18022

tmp18023 := MakeNative(func(__e *ControlFlow) {
V264 := __e.Get(1)
_ = V264
tmp18024 := Call(__e, PrimFunc(symshen_4_5colon_1equal_6), V264)


let__17513 := tmp18024
_ = let__17513

tmp18050 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17513)


var ifres18025 Obj

if True == tmp18050 {
tmp18026 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres18025 = tmp18026


} else {
tmp18027 := Call(__e, PrimFunc(symshen_4in_1_6), let__17513)


let__17514 := tmp18027
_ = let__17514

tmp18049 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__17514)
}
__typedArg0 := let__17514
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18028 Obj

if True == tmp18049 {
tmp18029 := Call(__e, PrimFunc(symhead), let__17514)


let__17515 := tmp18029
_ = let__17515

tmp18030 := Call(__e, PrimFunc(symtail), let__17514)


let__17516 := tmp18030
_ = let__17516

tmp18047 := Call(__e, PrimFunc(symshen_4hds_a_2), let__17516, symwhere)


var ifres18031 Obj

if True == tmp18047 {
tmp18032 := Call(__e, PrimFunc(symtail), let__17516)


let__17517 := tmp18032
_ = let__17517

tmp18045 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__17517)
}
__typedArg0 := let__17517
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18033 Obj

if True == tmp18045 {
tmp18034 := Call(__e, PrimFunc(symhead), let__17517)


let__17518 := tmp18034
_ = let__17518

tmp18035 := Call(__e, PrimFunc(symtail), let__17517)


let__17519 := tmp18035
_ = let__17519

tmp18042 := Call(__e, PrimFunc(symshen_4semicolon_2), let__17515)


tmp18043 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp18042)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp18042
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres18036 Obj

if True == tmp18043 {
tmp18037 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17515, Nil)
}
__typedArg0 := let__17515
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18038 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17518, tmp18037)
}
__typedArg0 := let__17518
__typedArg1 := tmp18037
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18039 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhere, tmp18038)
}
__typedArg0 := symwhere
__typedArg1 := tmp18038
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18040 := Call(__e, PrimFunc(symshen_4comb), let__17519, tmp18039)


ifres18036 = tmp18040


} else {
tmp18041 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres18036 = tmp18041


}

ifres18033 = ifres18036


} else {
tmp18044 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres18033 = tmp18044


}

ifres18031 = ifres18033


} else {
tmp18046 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres18031 = tmp18046


}

ifres18028 = ifres18031


} else {
tmp18048 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres18028 = tmp18048


}

ifres18025 = ifres18028


}

let__17512 := ifres18025
_ = let__17512

tmp18069 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17512)


if True == tmp18069 {
tmp18051 := Call(__e, PrimFunc(symshen_4_5colon_1equal_6), V264)


let__17521 := tmp18051
_ = let__17521

tmp18065 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17521)


var ifres18052 Obj

if True == tmp18065 {
tmp18053 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres18052 = tmp18053


} else {
tmp18054 := Call(__e, PrimFunc(symshen_4in_1_6), let__17521)


let__17522 := tmp18054
_ = let__17522

tmp18064 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__17522)
}
__typedArg0 := let__17522
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18055 Obj

if True == tmp18064 {
tmp18056 := Call(__e, PrimFunc(symhead), let__17522)


let__17523 := tmp18056
_ = let__17523

tmp18057 := Call(__e, PrimFunc(symtail), let__17522)


let__17524 := tmp18057
_ = let__17524

tmp18061 := Call(__e, PrimFunc(symshen_4semicolon_2), let__17523)


tmp18062 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp18061)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp18061
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres18058 Obj

if True == tmp18062 {
tmp18059 := Call(__e, PrimFunc(symshen_4comb), let__17524, let__17523)


ifres18058 = tmp18059


} else {
tmp18060 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres18058 = tmp18060


}

ifres18055 = ifres18058


} else {
tmp18063 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres18055 = tmp18063


}

ifres18052 = ifres18055


}

let__17520 := ifres18052
_ = let__17520

tmp18067 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17520)


if True == tmp18067 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__17520)
return
}


} else {
__e.Return(let__17512)
return
}


}, 1)

tmp18070 := Call(__e, ns2_1set, symshen_4_5semantics_6, tmp18023)


_ = tmp18070

tmp18071 := MakeNative(func(__e *ControlFlow) {
V286 := __e.Get(1)
_ = V286
V287 := __e.Get(2)
_ = V287
V288 := __e.Get(3)
_ = V288
tmp18079 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V288)
}
__typedArg0 := Nil
__typedArg1 := V288
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp18079 {
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
tmp18077 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V288)
}
__typedArg0 := V288
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp18077 {
tmp18072 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V288)
}
__typedArg0 := V288
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18073 := Call(__e, PrimFunc(symshen_4c_1rule_1_6shen), V286, tmp18072, V287)


tmp18074 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V288)
}
__typedArg0 := V288
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18075 := Call(__e, PrimFunc(symshen_4c_1rules_1_6shen), V286, V287, tmp18074)


__e.TailApply(PrimFunc(symshen_4combine_1c_1code), tmp18073, tmp18075)
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

tmp18080 := Call(__e, ns2_1set, symshen_4c_1rules_1_6shen, tmp18071)


_ = tmp18080

tmp18081 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symfail))
return
}, 0)

tmp18082 := Call(__e, ns2_1set, symshen_4parse_1failure, tmp18081)


_ = tmp18082

tmp18083 := MakeNative(func(__e *ControlFlow) {
V289 := __e.Get(1)
_ = V289
V290 := __e.Get(2)
_ = V290
tmp18084 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symResult, Nil)
}
__typedArg0 := symResult
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18085 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4parse_1failure_2, tmp18084)
}
__typedArg0 := symshen_4parse_1failure_2
__typedArg1 := tmp18084
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18086 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symResult, Nil)
}
__typedArg0 := symResult
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18087 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V290, tmp18086)
}
__typedArg0 := V290
__typedArg1 := tmp18086
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18088 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18085, tmp18087)
}
__typedArg0 := tmp18085
__typedArg1 := tmp18087
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18089 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp18088)
}
__typedArg0 := symif
__typedArg1 := tmp18088
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18090 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18089, Nil)
}
__typedArg0 := tmp18089
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18091 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V289, tmp18090)
}
__typedArg0 := V289
__typedArg1 := tmp18090
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18092 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symResult, tmp18091)
}
__typedArg0 := symResult
__typedArg1 := tmp18091
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp18092)
}
__typedArg0 := symlet
__typedArg1 := tmp18092
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 2)

tmp18093 := Call(__e, ns2_1set, symshen_4combine_1c_1code, tmp18083)


_ = tmp18093

tmp18094 := MakeNative(func(__e *ControlFlow) {
V297 := __e.Get(1)
_ = V297
V298 := __e.Get(2)
_ = V298
V299 := __e.Get(3)
_ = V299
tmp18108 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V298)
}
__typedArg0 := V298
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18099 Obj

if True == tmp18108 {
tmp18106 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V298)
}
__typedArg0 := V298
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18107 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18106)
}
__typedArg0 := tmp18106
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18101 Obj

if True == tmp18107 {
tmp18103 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V298)
}
__typedArg0 := V298
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18104 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18103)
}
__typedArg0 := tmp18103
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18105 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp18104)
}
__typedArg0 := Nil
__typedArg1 := tmp18104
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18102 Obj

if True == tmp18105 {
ifres18102 = True


} else {
ifres18102 = False


}

ifres18101 = ifres18102


} else {
ifres18101 = False


}

var ifres18100 Obj

if True == ifres18101 {
ifres18100 = True


} else {
ifres18100 = False


}

ifres18099 = ifres18100


} else {
ifres18099 = False


}

if True == ifres18099 {
tmp18095 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V298)
}
__typedArg0 := V298
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18096 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V298)
}
__typedArg0 := V298
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18097 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18096)
}
__typedArg0 := tmp18096
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4yacc_1syntax), V297, V299, tmp18095, tmp18097)
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

tmp18109 := Call(__e, ns2_1set, symshen_4c_1rule_1_6shen, tmp18094)


_ = tmp18109

tmp18110 := MakeNative(func(__e *ControlFlow) {
V308 := __e.Get(1)
_ = V308
V309 := __e.Get(2)
_ = V309
V310 := __e.Get(3)
_ = V310
V311 := __e.Get(4)
_ = V311
tmp18174 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V310)
}
__typedArg0 := Nil
__typedArg1 := V310
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18152 Obj

if True == tmp18174 {
tmp18173 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V311)
}
__typedArg0 := V311
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18154 Obj

if True == tmp18173 {
tmp18171 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V311)
}
__typedArg0 := V311
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18172 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symwhere, tmp18171)
}
__typedArg0 := symwhere
__typedArg1 := tmp18171
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18156 Obj

if True == tmp18172 {
tmp18169 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V311)
}
__typedArg0 := V311
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18170 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18169)
}
__typedArg0 := tmp18169
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18158 Obj

if True == tmp18170 {
tmp18166 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V311)
}
__typedArg0 := V311
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18167 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18166)
}
__typedArg0 := tmp18166
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18168 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18167)
}
__typedArg0 := tmp18167
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18160 Obj

if True == tmp18168 {
tmp18162 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V311)
}
__typedArg0 := V311
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18163 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18162)
}
__typedArg0 := tmp18162
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18164 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18163)
}
__typedArg0 := tmp18163
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18165 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp18164)
}
__typedArg0 := Nil
__typedArg1 := tmp18164
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18161 Obj

if True == tmp18165 {
ifres18161 = True


} else {
ifres18161 = False


}

ifres18160 = ifres18161


} else {
ifres18160 = False


}

var ifres18159 Obj

if True == ifres18160 {
ifres18159 = True


} else {
ifres18159 = False


}

ifres18158 = ifres18159


} else {
ifres18158 = False


}

var ifres18157 Obj

if True == ifres18158 {
ifres18157 = True


} else {
ifres18157 = False


}

ifres18156 = ifres18157


} else {
ifres18156 = False


}

var ifres18155 Obj

if True == ifres18156 {
ifres18155 = True


} else {
ifres18155 = False


}

ifres18154 = ifres18155


} else {
ifres18154 = False


}

var ifres18153 Obj

if True == ifres18154 {
ifres18153 = True


} else {
ifres18153 = False


}

ifres18152 = ifres18153


} else {
ifres18152 = False


}

if True == ifres18152 {
tmp18111 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V311)
}
__typedArg0 := V311
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18112 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18111)
}
__typedArg0 := tmp18111
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18113 := Call(__e, PrimFunc(symshen_4process_1yacc_1semantics), tmp18112)


tmp18114 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V311)
}
__typedArg0 := V311
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18115 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18114)
}
__typedArg0 := tmp18114
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18116 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18115)
}
__typedArg0 := tmp18115
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18117 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V308, V309, Nil, tmp18116)


tmp18118 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4parse_1failure, Nil)
}
__typedArg0 := symshen_4parse_1failure
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18119 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18118, Nil)
}
__typedArg0 := tmp18118
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18120 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18117, tmp18119)
}
__typedArg0 := tmp18117
__typedArg1 := tmp18119
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18121 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18113, tmp18120)
}
__typedArg0 := tmp18113
__typedArg1 := tmp18120
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp18121)
}
__typedArg0 := symif
__typedArg1 := tmp18121
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp18150 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V310)
}
__typedArg0 := Nil
__typedArg1 := V310
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp18150 {
__e.TailApply(PrimFunc(symshen_4yacc_1semantics), V308, V309, V311)
return
} else {
tmp18148 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp18148 {
tmp18145 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18146 := Call(__e, PrimFunc(symshen_4non_1terminal_2), tmp18145)


if True == tmp18146 {
tmp18122 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18123 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4non_1terminalcode), V308, V309, tmp18122, tmp18123, V311)
return


} else {
tmp18142 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18143 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(tmp18142)
}
__typedArg0 := tmp18142
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

if True == tmp18143 {
tmp18124 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18125 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4variablecode), V308, V309, tmp18124, tmp18125, V311)
return


} else {
tmp18139 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18140 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym__, tmp18139)
}
__typedArg0 := sym__
__typedArg1 := tmp18139
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp18140 {
tmp18126 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18127 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4wildcardcode), V308, V309, tmp18126, tmp18127, V311)
return


} else {
tmp18136 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18137 := Call(__e, PrimFunc(symatom_2), tmp18136)


if True == tmp18137 {
tmp18128 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18129 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4terminalcode), V308, V309, tmp18128, tmp18129, V311)
return


} else {
tmp18133 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18134 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18133)
}
__typedArg0 := tmp18133
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp18134 {
tmp18130 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18131 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4conscode), V308, V309, tmp18130, tmp18131, V311)
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

tmp18175 := Call(__e, ns2_1set, symshen_4yacc_1syntax, tmp18110)


_ = tmp18175

tmp18176 := MakeNative(func(__e *ControlFlow) {
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
tmp18177 := Call(__e, PrimFunc(symconcat), symParse, V314)


let__17525 := tmp18177
_ = let__17525

tmp18178 := Call(__e, PrimFunc(symconcat), symAction, V314)


let__17526 := tmp18178
_ = let__17526

tmp18179 := Call(__e, PrimFunc(symconcat), symRemainder, V314)


let__17527 := tmp18179
_ = let__17527

tmp18180 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V313, Nil)
}
__typedArg0 := V313
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18181 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V314, tmp18180)
}
__typedArg0 := V314
__typedArg1 := tmp18180
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18182 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17525, Nil)
}
__typedArg0 := let__17525
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18183 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4parse_1failure_2, tmp18182)
}
__typedArg0 := symshen_4parse_1failure_2
__typedArg1 := tmp18182
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18184 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4parse_1failure, Nil)
}
__typedArg0 := symshen_4parse_1failure
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18185 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17525, Nil)
}
__typedArg0 := let__17525
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18186 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4in_1_6, tmp18185)
}
__typedArg0 := symshen_4in_1_6
__typedArg1 := tmp18185
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18187 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V312, let__17527, V315, V316)


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
return PrimCons(let__17527, tmp18189)
}
__typedArg0 := let__17527
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

let__17528 := tmp18191
_ = let__17528

tmp18202 := Call(__e, PrimFunc(symshen_4occurs_1check_2), V314, V316)


var ifres18199 Obj

if True == tmp18202 {
ifres18199 = True


} else {
tmp18201 := Call(__e, PrimFunc(symshen_4occurs_1check_2), let__17526, V316)


var ifres18200 Obj

if True == tmp18201 {
ifres18200 = True


} else {
ifres18200 = False


}

ifres18199 = ifres18200


}

var ifres18192 Obj

if True == ifres18199 {
tmp18193 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17525, Nil)
}
__typedArg0 := let__17525
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18194 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4_5_1out, tmp18193)
}
__typedArg0 := symshen_4_5_1out
__typedArg1 := tmp18193
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18195 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17528, Nil)
}
__typedArg0 := let__17528
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
return PrimCons(let__17526, tmp18196)
}
__typedArg0 := let__17526
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
ifres18192 = let__17528


}

tmp18203 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(ifres18192, Nil)
}
__typedArg0 := ifres18192
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18204 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18184, tmp18203)
}
__typedArg0 := tmp18184
__typedArg1 := tmp18203
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18205 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18183, tmp18204)
}
__typedArg0 := tmp18183
__typedArg1 := tmp18204
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18206 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp18205)
}
__typedArg0 := symif
__typedArg1 := tmp18205
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18207 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18206, Nil)
}
__typedArg0 := tmp18206
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18208 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18181, tmp18207)
}
__typedArg0 := tmp18181
__typedArg1 := tmp18207
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18209 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17525, tmp18208)
}
__typedArg0 := let__17525
__typedArg1 := tmp18208
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp18209)
}
__typedArg0 := symlet
__typedArg1 := tmp18209
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 5)

tmp18210 := Call(__e, ns2_1set, symshen_4non_1terminalcode, tmp18176)


_ = tmp18210

tmp18211 := MakeNative(func(__e *ControlFlow) {
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
tmp18212 := Call(__e, PrimFunc(symgensym), symRemainder)


let__17529 := tmp18212
_ = let__17529

tmp18213 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V322, Nil)
}
__typedArg0 := V322
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18214 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons_2, tmp18213)
}
__typedArg0 := symcons_2
__typedArg1 := tmp18213
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18215 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V322, Nil)
}
__typedArg0 := V322
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18216 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtail, tmp18215)
}
__typedArg0 := symtail
__typedArg1 := tmp18215
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18217 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V321, let__17529, V324, V325)


tmp18218 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18217, Nil)
}
__typedArg0 := tmp18217
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18219 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18216, tmp18218)
}
__typedArg0 := tmp18216
__typedArg1 := tmp18218
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18220 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17529, tmp18219)
}
__typedArg0 := let__17529
__typedArg1 := tmp18219
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18221 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp18220)
}
__typedArg0 := symlet
__typedArg1 := tmp18220
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__17530 := tmp18221
_ = let__17530

tmp18229 := Call(__e, PrimFunc(symshen_4occurs_1check_2), V323, V325)


var ifres18222 Obj

if True == tmp18229 {
tmp18223 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V322, Nil)
}
__typedArg0 := V322
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18224 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhead, tmp18223)
}
__typedArg0 := symhead
__typedArg1 := tmp18223
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18225 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17530, Nil)
}
__typedArg0 := let__17530
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18226 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18224, tmp18225)
}
__typedArg0 := tmp18224
__typedArg1 := tmp18225
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18227 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V323, tmp18226)
}
__typedArg0 := V323
__typedArg1 := tmp18226
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18228 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp18227)
}
__typedArg0 := symlet
__typedArg1 := tmp18227
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

ifres18222 = tmp18228


} else {
ifres18222 = let__17530


}

tmp18230 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4parse_1failure, Nil)
}
__typedArg0 := symshen_4parse_1failure
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18231 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18230, Nil)
}
__typedArg0 := tmp18230
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18232 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(ifres18222, tmp18231)
}
__typedArg0 := ifres18222
__typedArg1 := tmp18231
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18233 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18214, tmp18232)
}
__typedArg0 := tmp18214
__typedArg1 := tmp18232
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp18233)
}
__typedArg0 := symif
__typedArg1 := tmp18233
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 5)

tmp18234 := Call(__e, ns2_1set, symshen_4variablecode, tmp18211)


_ = tmp18234

tmp18235 := MakeNative(func(__e *ControlFlow) {
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
tmp18236 := Call(__e, PrimFunc(symgensym), symRemainder)


let__17531 := tmp18236
_ = let__17531

tmp18237 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V329, Nil)
}
__typedArg0 := V329
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18238 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons_2, tmp18237)
}
__typedArg0 := symcons_2
__typedArg1 := tmp18237
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18239 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V329, Nil)
}
__typedArg0 := V329
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18240 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtail, tmp18239)
}
__typedArg0 := symtail
__typedArg1 := tmp18239
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18241 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V328, let__17531, V331, V332)


tmp18242 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18241, Nil)
}
__typedArg0 := tmp18241
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18243 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18240, tmp18242)
}
__typedArg0 := tmp18240
__typedArg1 := tmp18242
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18244 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17531, tmp18243)
}
__typedArg0 := let__17531
__typedArg1 := tmp18243
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18245 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp18244)
}
__typedArg0 := symlet
__typedArg1 := tmp18244
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18246 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4parse_1failure, Nil)
}
__typedArg0 := symshen_4parse_1failure
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18247 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18246, Nil)
}
__typedArg0 := tmp18246
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18248 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18245, tmp18247)
}
__typedArg0 := tmp18245
__typedArg1 := tmp18247
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18249 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18238, tmp18248)
}
__typedArg0 := tmp18238
__typedArg1 := tmp18248
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp18249)
}
__typedArg0 := symif
__typedArg1 := tmp18249
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 5)

tmp18250 := Call(__e, ns2_1set, symshen_4wildcardcode, tmp18235)


_ = tmp18250

tmp18251 := MakeNative(func(__e *ControlFlow) {
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
tmp18252 := Call(__e, PrimFunc(symgensym), symRemainder)


let__17532 := tmp18252
_ = let__17532

tmp18253 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V336, Nil)
}
__typedArg0 := V336
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18254 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V335, tmp18253)
}
__typedArg0 := V335
__typedArg1 := tmp18253
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18255 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4hds_a_2, tmp18254)
}
__typedArg0 := symshen_4hds_a_2
__typedArg1 := tmp18254
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18256 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V335, Nil)
}
__typedArg0 := V335
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18257 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtail, tmp18256)
}
__typedArg0 := symtail
__typedArg1 := tmp18256
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18258 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V334, let__17532, V337, V338)


tmp18259 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18258, Nil)
}
__typedArg0 := tmp18258
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18260 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18257, tmp18259)
}
__typedArg0 := tmp18257
__typedArg1 := tmp18259
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18261 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17532, tmp18260)
}
__typedArg0 := let__17532
__typedArg1 := tmp18260
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18262 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp18261)
}
__typedArg0 := symlet
__typedArg1 := tmp18261
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18263 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4parse_1failure, Nil)
}
__typedArg0 := symshen_4parse_1failure
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18264 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18263, Nil)
}
__typedArg0 := tmp18263
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18265 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18262, tmp18264)
}
__typedArg0 := tmp18262
__typedArg1 := tmp18264
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18266 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18255, tmp18265)
}
__typedArg0 := tmp18255
__typedArg1 := tmp18265
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp18266)
}
__typedArg0 := symif
__typedArg1 := tmp18266
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 5)

tmp18267 := Call(__e, ns2_1set, symshen_4terminalcode, tmp18251)


_ = tmp18267

tmp18268 := MakeNative(func(__e *ControlFlow) {
V347 := __e.Get(1)
_ = V347
V348 := __e.Get(2)
_ = V348
tmp18274 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V347)
}
__typedArg0 := V347
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18270 Obj

if True == tmp18274 {
tmp18272 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V347)
}
__typedArg0 := V347
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18273 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp18272, V348)
}
__typedArg0 := tmp18272
__typedArg1 := V348
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18271 Obj

if True == tmp18273 {
ifres18271 = True


} else {
ifres18271 = False


}

ifres18270 = ifres18271


} else {
ifres18270 = False


}

if True == ifres18270 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}, 2)

tmp18275 := Call(__e, ns2_1set, symshen_4hds_a_2, tmp18268)


_ = tmp18275

tmp18276 := MakeNative(func(__e *ControlFlow) {
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
tmp18277 := Call(__e, PrimFunc(symgensym), symRemainder)


let__17533 := tmp18277
_ = let__17533

tmp18278 := Call(__e, PrimFunc(symgensym), symHd)


let__17534 := tmp18278
_ = let__17534

tmp18279 := Call(__e, PrimFunc(symgensym), symTl)


let__17535 := tmp18279
_ = let__17535

tmp18280 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V350, Nil)
}
__typedArg0 := V350
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18281 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4ccons_2, tmp18280)
}
__typedArg0 := symshen_4ccons_2
__typedArg1 := tmp18280
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18282 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V350, Nil)
}
__typedArg0 := V350
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18283 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhead, tmp18282)
}
__typedArg0 := symhead
__typedArg1 := tmp18282
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18284 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V350, Nil)
}
__typedArg0 := V350
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18285 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtail, tmp18284)
}
__typedArg0 := symtail
__typedArg1 := tmp18284
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18286 := Call(__e, PrimFunc(symshen_4decons), V351)


tmp18287 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5end_6, Nil)
}
__typedArg0 := sym_5end_6
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18288 := Call(__e, PrimFunc(symappend), tmp18286, tmp18287)


tmp18289 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V349, let__17535, V352, V353)


tmp18290 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18289, Nil)
}
__typedArg0 := tmp18289
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18291 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4processed, tmp18290)
}
__typedArg0 := symshen_4processed
__typedArg1 := tmp18290
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18292 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V349, let__17534, tmp18288, tmp18291)


tmp18293 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18292, Nil)
}
__typedArg0 := tmp18292
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18294 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18285, tmp18293)
}
__typedArg0 := tmp18285
__typedArg1 := tmp18293
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18295 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17535, tmp18294)
}
__typedArg0 := let__17535
__typedArg1 := tmp18294
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18296 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18283, tmp18295)
}
__typedArg0 := tmp18283
__typedArg1 := tmp18295
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18297 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17534, tmp18296)
}
__typedArg0 := let__17534
__typedArg1 := tmp18296
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18298 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp18297)
}
__typedArg0 := symlet
__typedArg1 := tmp18297
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18299 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4parse_1failure, Nil)
}
__typedArg0 := symshen_4parse_1failure
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18300 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18299, Nil)
}
__typedArg0 := tmp18299
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18301 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18298, tmp18300)
}
__typedArg0 := tmp18298
__typedArg1 := tmp18300
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18302 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18281, tmp18301)
}
__typedArg0 := tmp18281
__typedArg1 := tmp18301
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp18302)
}
__typedArg0 := symif
__typedArg1 := tmp18302
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 5)

tmp18303 := Call(__e, ns2_1set, symshen_4conscode, tmp18276)


_ = tmp18303

tmp18304 := MakeNative(func(__e *ControlFlow) {
V367 := __e.Get(1)
_ = V367
tmp18316 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V367)
}
__typedArg0 := V367
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18312 Obj

if True == tmp18316 {
tmp18314 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V367)
}
__typedArg0 := V367
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18315 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18314)
}
__typedArg0 := tmp18314
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18313 Obj

if True == tmp18315 {
ifres18313 = True


} else {
ifres18313 = False


}

ifres18312 = ifres18313


} else {
ifres18312 = False


}

if True == ifres18312 {
__e.Return(True)
return
} else {
tmp18310 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V367)
}
__typedArg0 := V367
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18306 Obj

if True == tmp18310 {
tmp18308 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V367)
}
__typedArg0 := V367
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18309 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp18308)
}
__typedArg0 := Nil
__typedArg1 := tmp18308
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18307 Obj

if True == tmp18309 {
ifres18307 = True


} else {
ifres18307 = False


}

ifres18306 = ifres18307


} else {
ifres18306 = False


}

if True == ifres18306 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp18317 := Call(__e, ns2_1set, symshen_4ccons_2, tmp18304)


_ = tmp18317

tmp18318 := MakeNative(func(__e *ControlFlow) {
V368 := __e.Get(1)
_ = V368
tmp18345 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V368)
}
__typedArg0 := V368
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18326 Obj

if True == tmp18345 {
tmp18343 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V368)
}
__typedArg0 := V368
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18344 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcons, tmp18343)
}
__typedArg0 := symcons
__typedArg1 := tmp18343
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18328 Obj

if True == tmp18344 {
tmp18341 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V368)
}
__typedArg0 := V368
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18342 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18341)
}
__typedArg0 := tmp18341
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18330 Obj

if True == tmp18342 {
tmp18338 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V368)
}
__typedArg0 := V368
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18339 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18338)
}
__typedArg0 := tmp18338
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18340 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18339)
}
__typedArg0 := tmp18339
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18332 Obj

if True == tmp18340 {
tmp18334 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V368)
}
__typedArg0 := V368
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18335 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18334)
}
__typedArg0 := tmp18334
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18336 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18335)
}
__typedArg0 := tmp18335
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18337 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp18336)
}
__typedArg0 := Nil
__typedArg1 := tmp18336
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18333 Obj

if True == tmp18337 {
ifres18333 = True


} else {
ifres18333 = False


}

ifres18332 = ifres18333


} else {
ifres18332 = False


}

var ifres18331 Obj

if True == ifres18332 {
ifres18331 = True


} else {
ifres18331 = False


}

ifres18330 = ifres18331


} else {
ifres18330 = False


}

var ifres18329 Obj

if True == ifres18330 {
ifres18329 = True


} else {
ifres18329 = False


}

ifres18328 = ifres18329


} else {
ifres18328 = False


}

var ifres18327 Obj

if True == ifres18328 {
ifres18327 = True


} else {
ifres18327 = False


}

ifres18326 = ifres18327


} else {
ifres18326 = False


}

if True == ifres18326 {
tmp18319 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V368)
}
__typedArg0 := V368
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18320 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18319)
}
__typedArg0 := tmp18319
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18321 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V368)
}
__typedArg0 := V368
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18322 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18321)
}
__typedArg0 := tmp18321
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18323 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18322)
}
__typedArg0 := tmp18322
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18324 := Call(__e, PrimFunc(symshen_4decons), tmp18323)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18320, tmp18324)
}
__typedArg0 := tmp18320
__typedArg1 := tmp18324
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return(V368)
return
}


}, 1)

tmp18346 := Call(__e, ns2_1set, symshen_4decons, tmp18318)


_ = tmp18346

tmp18347 := MakeNative(func(__e *ControlFlow) {
V369 := __e.Get(1)
_ = V369
V370 := __e.Get(2)
_ = V370
tmp18348 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V370, Nil)
}
__typedArg0 := V370
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V369, tmp18348)
}
__typedArg0 := V369
__typedArg1 := tmp18348
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 2)

tmp18349 := Call(__e, ns2_1set, symshen_4comb, tmp18347)


_ = tmp18349

tmp18350 := MakeNative(func(__e *ControlFlow) {
V375 := __e.Get(1)
_ = V375
V376 := __e.Get(2)
_ = V376
V377 := __e.Get(3)
_ = V377
tmp18370 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V377)
}
__typedArg0 := V377
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18357 Obj

if True == tmp18370 {
tmp18368 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V377)
}
__typedArg0 := V377
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18369 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4processed, tmp18368)
}
__typedArg0 := symshen_4processed
__typedArg1 := tmp18368
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18359 Obj

if True == tmp18369 {
tmp18366 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V377)
}
__typedArg0 := V377
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18367 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18366)
}
__typedArg0 := tmp18366
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18361 Obj

if True == tmp18367 {
tmp18363 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V377)
}
__typedArg0 := V377
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18364 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18363)
}
__typedArg0 := tmp18363
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18365 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp18364)
}
__typedArg0 := Nil
__typedArg1 := tmp18364
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18362 Obj

if True == tmp18365 {
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

if True == ifres18357 {
tmp18351 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V377)
}
__typedArg0 := V377
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18351)
}
__typedArg0 := tmp18351
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return


} else {
tmp18352 := Call(__e, PrimFunc(symshen_4process_1yacc_1semantics), V377)


let__17536 := tmp18352
_ = let__17536

tmp18353 := Call(__e, PrimFunc(symshen_4use_1type_1info), V375, let__17536)


let__17537 := tmp18353
_ = let__17537

tmp18354 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17537, Nil)
}
__typedArg0 := let__17537
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18355 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V376, tmp18354)
}
__typedArg0 := V376
__typedArg1 := tmp18354
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4comb, tmp18355)
}
__typedArg0 := symshen_4comb
__typedArg1 := tmp18355
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}


}, 3)

tmp18371 := Call(__e, ns2_1set, symshen_4yacc_1semantics, tmp18350)


_ = tmp18371

tmp18372 := MakeNative(func(__e *ControlFlow) {
V383 := __e.Get(1)
_ = V383
V384 := __e.Get(2)
_ = V384
tmp18560 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18381 Obj

if True == tmp18560 {
tmp18558 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18559 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_i, tmp18558)
}
__typedArg0 := sym_i
__typedArg1 := tmp18558
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18383 Obj

if True == tmp18559 {
tmp18556 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18557 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18556)
}
__typedArg0 := tmp18556
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18385 Obj

if True == tmp18557 {
tmp18553 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18554 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18553)
}
__typedArg0 := tmp18553
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18555 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18554)
}
__typedArg0 := tmp18554
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18387 Obj

if True == tmp18555 {
tmp18549 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18550 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18549)
}
__typedArg0 := tmp18549
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18551 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18550)
}
__typedArg0 := tmp18550
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18552 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlist, tmp18551)
}
__typedArg0 := symlist
__typedArg1 := tmp18551
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18389 Obj

if True == tmp18552 {
tmp18545 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18546 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18545)
}
__typedArg0 := tmp18545
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18547 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18546)
}
__typedArg0 := tmp18546
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18548 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18547)
}
__typedArg0 := tmp18547
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18391 Obj

if True == tmp18548 {
tmp18540 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18541 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18540)
}
__typedArg0 := tmp18540
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18542 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18541)
}
__typedArg0 := tmp18541
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18543 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18542)
}
__typedArg0 := tmp18542
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18544 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp18543)
}
__typedArg0 := Nil
__typedArg1 := tmp18543
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18393 Obj

if True == tmp18544 {
tmp18537 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18538 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18537)
}
__typedArg0 := tmp18537
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18539 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18538)
}
__typedArg0 := tmp18538
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18395 Obj

if True == tmp18539 {
tmp18533 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18534 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18533)
}
__typedArg0 := tmp18533
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18535 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18534)
}
__typedArg0 := tmp18534
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18536 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_1_1_6, tmp18535)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp18535
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18397 Obj

if True == tmp18536 {
tmp18529 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18530 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18529)
}
__typedArg0 := tmp18529
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18531 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18530)
}
__typedArg0 := tmp18530
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18532 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18531)
}
__typedArg0 := tmp18531
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18399 Obj

if True == tmp18532 {
tmp18524 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18525 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18524)
}
__typedArg0 := tmp18524
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18526 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18525)
}
__typedArg0 := tmp18525
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18527 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18526)
}
__typedArg0 := tmp18526
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18528 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18527)
}
__typedArg0 := tmp18527
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18401 Obj

if True == tmp18528 {
tmp18518 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18519 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18518)
}
__typedArg0 := tmp18518
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18520 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18519)
}
__typedArg0 := tmp18519
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18521)
}
__typedArg0 := tmp18521
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18523 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symstr, tmp18522)
}
__typedArg0 := symstr
__typedArg1 := tmp18522
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18403 Obj

if True == tmp18523 {
tmp18512 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18515)
}
__typedArg0 := tmp18515
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18517 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18516)
}
__typedArg0 := tmp18516
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18405 Obj

if True == tmp18517 {
tmp18505 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18506)
}
__typedArg0 := tmp18506
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18508 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18507)
}
__typedArg0 := tmp18507
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18509 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18508)
}
__typedArg0 := tmp18508
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18510 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18509)
}
__typedArg0 := tmp18509
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18511 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18510)
}
__typedArg0 := tmp18510
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18407 Obj

if True == tmp18511 {
tmp18497 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18498 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18497)
}
__typedArg0 := tmp18497
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18499)
}
__typedArg0 := tmp18499
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18501 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18500)
}
__typedArg0 := tmp18500
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18502 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18501)
}
__typedArg0 := tmp18501
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18503 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18502)
}
__typedArg0 := tmp18502
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18504 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlist, tmp18503)
}
__typedArg0 := symlist
__typedArg1 := tmp18503
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18409 Obj

if True == tmp18504 {
tmp18489 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18490)
}
__typedArg0 := tmp18490
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18492 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18491)
}
__typedArg0 := tmp18491
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18493 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18492)
}
__typedArg0 := tmp18492
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18494 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18493)
}
__typedArg0 := tmp18493
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18495 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18494)
}
__typedArg0 := tmp18494
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18496 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18495)
}
__typedArg0 := tmp18495
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18411 Obj

if True == tmp18496 {
tmp18480 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18481 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18480)
}
__typedArg0 := tmp18480
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18482)
}
__typedArg0 := tmp18482
return Call(__e, PrimFunc(symhd), __typedArg0)
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18486)
}
__typedArg0 := tmp18486
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18488 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp18487)
}
__typedArg0 := Nil
__typedArg1 := tmp18487
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18413 Obj

if True == tmp18488 {
tmp18473 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18474 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18473)
}
__typedArg0 := tmp18473
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18475 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18474)
}
__typedArg0 := tmp18474
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18476 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18475)
}
__typedArg0 := tmp18475
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18477 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18476)
}
__typedArg0 := tmp18476
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18478)
}
__typedArg0 := tmp18478
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18415 Obj

if True == tmp18479 {
tmp18465 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18467)
}
__typedArg0 := tmp18467
return Call(__e, PrimFunc(symhd), __typedArg0)
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18470)
}
__typedArg0 := tmp18470
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18472 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp18471)
}
__typedArg0 := Nil
__typedArg1 := tmp18471
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18417 Obj

if True == tmp18472 {
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18462)
}
__typedArg0 := tmp18462
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18464 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18463)
}
__typedArg0 := tmp18463
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18419 Obj

if True == tmp18464 {
tmp18454 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18457)
}
__typedArg0 := tmp18457
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18459 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_j, tmp18458)
}
__typedArg0 := sym_j
__typedArg1 := tmp18458
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18421 Obj

if True == tmp18459 {
tmp18448 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18451)
}
__typedArg0 := tmp18451
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18453 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp18452)
}
__typedArg0 := Nil
__typedArg1 := tmp18452
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18423 Obj

if True == tmp18453 {
tmp18435 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18436 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18435)
}
__typedArg0 := tmp18435
return Call(__e, PrimFunc(symhd), __typedArg0)
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
return PrimTail(V383)
}
__typedArg0 := V383
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18441)
}
__typedArg0 := tmp18441
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18443 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18442)
}
__typedArg0 := tmp18442
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18444 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18443)
}
__typedArg0 := tmp18443
return Call(__e, PrimFunc(symhd), __typedArg0)
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp18438, tmp18446)
}
__typedArg0 := tmp18438
__typedArg1 := tmp18446
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18425 Obj

if True == tmp18447 {
tmp18427 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18428 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18427)
}
__typedArg0 := tmp18427
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18429)
}
__typedArg0 := tmp18429
return Call(__e, PrimFunc(symhd), __typedArg0)
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18432)
}
__typedArg0 := tmp18432
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18434 := Call(__e, PrimFunc(symshen_4monomorphic_2), tmp18433)


var ifres18426 Obj

if True == tmp18434 {
ifres18426 = True


} else {
ifres18426 = False


}

ifres18425 = ifres18426


} else {
ifres18425 = False


}

var ifres18424 Obj

if True == ifres18425 {
ifres18424 = True


} else {
ifres18424 = False


}

ifres18423 = ifres18424


} else {
ifres18423 = False


}

var ifres18422 Obj

if True == ifres18423 {
ifres18422 = True


} else {
ifres18422 = False


}

ifres18421 = ifres18422


} else {
ifres18421 = False


}

var ifres18420 Obj

if True == ifres18421 {
ifres18420 = True


} else {
ifres18420 = False


}

ifres18419 = ifres18420


} else {
ifres18419 = False


}

var ifres18418 Obj

if True == ifres18419 {
ifres18418 = True


} else {
ifres18418 = False


}

ifres18417 = ifres18418


} else {
ifres18417 = False


}

var ifres18416 Obj

if True == ifres18417 {
ifres18416 = True


} else {
ifres18416 = False


}

ifres18415 = ifres18416


} else {
ifres18415 = False


}

var ifres18414 Obj

if True == ifres18415 {
ifres18414 = True


} else {
ifres18414 = False


}

ifres18413 = ifres18414


} else {
ifres18413 = False


}

var ifres18412 Obj

if True == ifres18413 {
ifres18412 = True


} else {
ifres18412 = False


}

ifres18411 = ifres18412


} else {
ifres18411 = False


}

var ifres18410 Obj

if True == ifres18411 {
ifres18410 = True


} else {
ifres18410 = False


}

ifres18409 = ifres18410


} else {
ifres18409 = False


}

var ifres18408 Obj

if True == ifres18409 {
ifres18408 = True


} else {
ifres18408 = False


}

ifres18407 = ifres18408


} else {
ifres18407 = False


}

var ifres18406 Obj

if True == ifres18407 {
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

if True == ifres18381 {
tmp18373 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18374 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18373)
}
__typedArg0 := tmp18373
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18375 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18374)
}
__typedArg0 := tmp18374
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18376 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18375)
}
__typedArg0 := tmp18375
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18377 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18376)
}
__typedArg0 := tmp18376
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18378 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18377)
}
__typedArg0 := tmp18377
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18379 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V384, tmp18378)
}
__typedArg0 := V384
__typedArg1 := tmp18378
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtype, tmp18379)
}
__typedArg0 := symtype
__typedArg1 := tmp18379
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return(V384)
return
}


}, 2)

tmp18561 := Call(__e, ns2_1set, symshen_4use_1type_1info, tmp18372)


_ = tmp18561

tmp18562 := MakeNative(func(__e *ControlFlow) {
V387 := __e.Get(1)
_ = V387
tmp18572 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(V387)
}
__typedArg0 := V387
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

if True == tmp18572 {
__e.Return(False)
return
} else {
tmp18570 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V387)
}
__typedArg0 := V387
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp18570 {
tmp18567 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V387)
}
__typedArg0 := V387
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18568 := Call(__e, PrimFunc(symshen_4monomorphic_2), tmp18567)


if True == tmp18568 {
tmp18564 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V387)
}
__typedArg0 := V387
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18565 := Call(__e, PrimFunc(symshen_4monomorphic_2), tmp18564)


if True == tmp18565 {
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

tmp18573 := Call(__e, ns2_1set, symshen_4monomorphic_2, tmp18562)


_ = tmp18573

tmp18574 := MakeNative(func(__e *ControlFlow) {
V388 := __e.Get(1)
_ = V388
tmp18600 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V388)
}
__typedArg0 := V388
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18582 Obj

if True == tmp18600 {
tmp18598 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V388)
}
__typedArg0 := V388
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18599 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symprotect, tmp18598)
}
__typedArg0 := symprotect
__typedArg1 := tmp18598
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18584 Obj

if True == tmp18599 {
tmp18596 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V388)
}
__typedArg0 := V388
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18597 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18596)
}
__typedArg0 := tmp18596
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18586 Obj

if True == tmp18597 {
tmp18593 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V388)
}
__typedArg0 := V388
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18594 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18593)
}
__typedArg0 := tmp18593
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18595 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp18594)
}
__typedArg0 := Nil
__typedArg1 := tmp18594
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18588 Obj

if True == tmp18595 {
tmp18590 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V388)
}
__typedArg0 := V388
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18591 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18590)
}
__typedArg0 := tmp18590
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18592 := Call(__e, PrimFunc(symshen_4non_1terminal_2), tmp18591)


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

var ifres18585 Obj

if True == ifres18586 {
ifres18585 = True


} else {
ifres18585 = False


}

ifres18584 = ifres18585


} else {
ifres18584 = False


}

var ifres18583 Obj

if True == ifres18584 {
ifres18583 = True


} else {
ifres18583 = False


}

ifres18582 = ifres18583


} else {
ifres18582 = False


}

if True == ifres18582 {
tmp18575 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V388)
}
__typedArg0 := V388
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18575)
}
__typedArg0 := tmp18575
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return


} else {
tmp18580 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V388)
}
__typedArg0 := V388
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp18580 {
tmp18576 := MakeNative(func(__e *ControlFlow) {
Z389 := __e.Get(1)
_ = Z389
__e.TailApply(PrimFunc(symshen_4process_1yacc_1semantics), Z389)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp18576, V388)
return


} else {
tmp18578 := Call(__e, PrimFunc(symshen_4non_1terminal_2), V388)


if True == tmp18578 {
__e.TailApply(PrimFunc(symconcat), symAction, V388)
return
} else {
__e.Return(V388)
return
}


}


}


}, 1)

tmp18601 := Call(__e, ns2_1set, symshen_4process_1yacc_1semantics, tmp18574)


_ = tmp18601

tmp18602 := MakeNative(func(__e *ControlFlow) {
V392 := __e.Get(1)
_ = V392
tmp18615 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V392)
}
__typedArg0 := V392
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18606 Obj

if True == tmp18615 {
tmp18613 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V392)
}
__typedArg0 := V392
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18614 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18613)
}
__typedArg0 := tmp18613
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18608 Obj

if True == tmp18614 {
tmp18610 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V392)
}
__typedArg0 := V392
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18611 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18610)
}
__typedArg0 := tmp18610
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18612 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp18611)
}
__typedArg0 := Nil
__typedArg1 := tmp18611
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18609 Obj

if True == tmp18612 {
ifres18609 = True


} else {
ifres18609 = False


}

ifres18608 = ifres18609


} else {
ifres18608 = False


}

var ifres18607 Obj

if True == ifres18608 {
ifres18607 = True


} else {
ifres18607 = False


}

ifres18606 = ifres18607


} else {
ifres18606 = False


}

if True == ifres18606 {
tmp18603 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V392)
}
__typedArg0 := V392
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18603)
}
__typedArg0 := tmp18603
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return


} else {
tmp18604 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V392)
}
__typedArg0 := V392
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18604)
}
__typedArg0 := tmp18604
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return


}


}, 1)

tmp18616 := Call(__e, ns2_1set, symshen_4_5_1out, tmp18602)


_ = tmp18616

tmp18617 := MakeNative(func(__e *ControlFlow) {
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

tmp18618 := Call(__e, ns2_1set, symshen_4in_1_6, tmp18617)


_ = tmp18618

tmp18619 := MakeNative(func(__e *ControlFlow) {
V394 := __e.Get(1)
_ = V394
tmp18620 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V394, Nil)
}
__typedArg0 := V394
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Nil, tmp18620)
}
__typedArg0 := Nil
__typedArg1 := tmp18620
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp18621 := Call(__e, ns2_1set, sym_5_b_6, tmp18619)


_ = tmp18621

tmp18622 := MakeNative(func(__e *ControlFlow) {
V395 := __e.Get(1)
_ = V395
tmp18623 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Nil, Nil)
}
__typedArg0 := Nil
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V395, tmp18623)
}
__typedArg0 := V395
__typedArg1 := tmp18623
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp18624 := Call(__e, ns2_1set, sym_5e_6, tmp18622)


_ = tmp18624

tmp18625 := MakeNative(func(__e *ControlFlow) {
V398 := __e.Get(1)
_ = V398
tmp18628 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V398)
}
__typedArg0 := Nil
__typedArg1 := V398
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp18628 {
tmp18626 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Nil, Nil)
}
__typedArg0 := Nil
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Nil, tmp18626)
}
__typedArg0 := Nil
__typedArg1 := tmp18626
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

__e.TailApply(ns2_1set, sym_5end_6, tmp18625)
return




}, 0)

