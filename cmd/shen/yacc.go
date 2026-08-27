package main

import . "github.com/pyrex41/shen-go/kl"

var YaccMain = MakeNative(func(__e *ControlFlow) {
tmp17482 := MakeNative(func(__e *ControlFlow) {
V112 := __e.Get(1)
_ = V112
V113 := __e.Get(2)
_ = V113
tmp17483 := MakeNative(func(__e *ControlFlow) {
W114 := __e.Get(1)
_ = W114
tmp17490 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W114)


if True == tmp17490 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("parse failure\n"))
}
__typedArg0 := MakeString("parse failure\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
} else {
tmp17488 := Call(__e, PrimFunc(symshen_4partial_1parse_1failure_2), W114)


if True == tmp17488 {
tmp17484 := Call(__e, PrimFunc(symshen_4in_1_6), W114)


tmp17485 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dresidue_d, tmp17484)
}
__typedArg0 := symshen_4_dresidue_d
__typedArg1 := tmp17484
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp17485

tmp17486 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dresidue_d)
}
__typedArg0 := symshen_4_dresidue_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4raise_1syntax_1error), tmp17486)
return


} else {
__e.TailApply(PrimFunc(symshen_4_5_1out), W114)
return
}


}


}, 1)

tmp17491 := Call(__e, V112, V113)


__e.TailApply(tmp17483, tmp17491)
return


}, 2)

tmp17492 := Call(__e, ns2_1set, symcompile, tmp17482)


_ = tmp17492

tmp17493 := MakeNative(func(__e *ControlFlow) {
V115 := __e.Get(1)
_ = V115
tmp17494 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dmaximum_1print_1sequence_1size_d)
}
__typedArg0 := sym_dmaximum_1print_1sequence_1size_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp17495 := Call(__e, PrimFunc(symshen_4syntax_1error_1message), tmp17494, MakeNumber(0), V115)


tmp17497 := Call(__e, PrimFunc(symshen_4proc_1nl), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("syntax error here: "))
__typedS1, __typedOK1 := TypedString(tmp17495)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("syntax error here: ")
__typedArg1 := tmp17495
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(tmp17497)
}
__typedArg0 := tmp17497
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}, 1)

tmp17498 := Call(__e, ns2_1set, symshen_4raise_1syntax_1error, tmp17493)


_ = tmp17498

tmp17499 := MakeNative(func(__e *ControlFlow) {
V123 := __e.Get(1)
_ = V123
V124 := __e.Get(2)
_ = V124
V125 := __e.Get(3)
_ = V125
tmp17510 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V125)
}
__typedArg0 := Nil
__typedArg1 := V125
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17510 {
__e.Return(MakeString("\n"))
return
} else {
tmp17508 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V123, V124)
}
__typedArg0 := V123
__typedArg1 := V124
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17508 {
__e.Return(MakeString("...etc \n"))
return
} else {
tmp17506 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V125)
}
__typedArg0 := V125
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17506 {
tmp17500 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V125)
}
__typedArg0 := V125
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17501 := Call(__e, PrimFunc(symshen_4app), tmp17500, MakeString(" "), symshen_4s)


tmp17502 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V124)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V124
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()

tmp17503 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V125)
}
__typedArg0 := V125
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17504 := Call(__e, PrimFunc(symshen_4syntax_1error_1message), V123, tmp17502, tmp17503)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(tmp17501)
__typedS1, __typedOK1 := TypedString(tmp17504)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := tmp17501
__typedArg1 := tmp17504
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("partial function shen.syntax-error-message"))
}
__typedArg0 := MakeString("partial function shen.syntax-error-message")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}


}, 3)

tmp17511 := Call(__e, ns2_1set, symshen_4syntax_1error_1message, tmp17499)


_ = tmp17511

tmp17512 := MakeNative(func(__e *ControlFlow) {
V126 := __e.Get(1)
_ = V126
tmp17513 := Call(__e, PrimFunc(symfail))


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V126, tmp17513)
}
__typedArg0 := V126
__typedArg1 := tmp17513
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp17514 := Call(__e, ns2_1set, symshen_4parse_1failure_2, tmp17512)


_ = tmp17514

tmp17515 := MakeNative(func(__e *ControlFlow) {
V127 := __e.Get(1)
_ = V127
tmp17516 := Call(__e, PrimFunc(symshen_4in_1_6), V127)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp17516)
}
__typedArg0 := tmp17516
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})())
return


}, 1)

tmp17517 := Call(__e, ns2_1set, symshen_4partial_1parse_1failure_2, tmp17515)


_ = tmp17517

tmp17518 := MakeNative(func(__e *ControlFlow) {
V130 := __e.Get(1)
_ = V130
tmp17531 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V130)
}
__typedArg0 := V130
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17522 Obj

if True == tmp17531 {
tmp17529 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V130)
}
__typedArg0 := V130
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17530 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp17529)
}
__typedArg0 := tmp17529
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17524 Obj

if True == tmp17530 {
tmp17526 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V130)
}
__typedArg0 := V130
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17527 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp17526)
}
__typedArg0 := tmp17526
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17528 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp17527)
}
__typedArg0 := Nil
__typedArg1 := tmp17527
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17525 Obj

if True == tmp17528 {
ifres17525 = True


} else {
ifres17525 = False


}

ifres17524 = ifres17525


} else {
ifres17524 = False


}

var ifres17523 Obj

if True == ifres17524 {
ifres17523 = True


} else {
ifres17523 = False


}

ifres17522 = ifres17523


} else {
ifres17522 = False


}

if True == ifres17522 {
tmp17519 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V130)
}
__typedArg0 := V130
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp17519)
}
__typedArg0 := tmp17519
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return


} else {
tmp17520 := Call(__e, PrimFunc(symshen_4app), V130, MakeString(" is not a YACC stream\n"), symshen_4s)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(tmp17520)
}
__typedArg0 := tmp17520
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}


}, 1)

tmp17532 := Call(__e, ns2_1set, symshen_4objectcode, tmp17518)


_ = tmp17532

tmp17533 := MakeNative(func(__e *ControlFlow) {
V131 := __e.Get(1)
_ = V131
tmp17534 := MakeNative(func(__e *ControlFlow) {
Z132 := __e.Get(1)
_ = Z132
__e.TailApply(PrimFunc(symshen_4_5yacc_6), Z132)
return
}, 1)

__e.TailApply(PrimFunc(symcompile), tmp17534, V131)
return


}, 1)

tmp17535 := Call(__e, ns2_1set, symshen_4yacc_1_6shen, tmp17533)


_ = tmp17535

tmp17536 := MakeNative(func(__e *ControlFlow) {
V133 := __e.Get(1)
_ = V133
tmp17537 := MakeNative(func(__e *ControlFlow) {
W134 := __e.Get(1)
_ = W134
tmp17539 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W134)


if True == tmp17539 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W134)
return
}


}, 1)

tmp17575 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V133)
}
__typedArg0 := V133
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17540 Obj

if True == tmp17575 {
tmp17541 := MakeNative(func(__e *ControlFlow) {
W135 := __e.Get(1)
_ = W135
tmp17542 := MakeNative(func(__e *ControlFlow) {
W136 := __e.Get(1)
_ = W136
tmp17543 := MakeNative(func(__e *ControlFlow) {
W137 := __e.Get(1)
_ = W137
tmp17569 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W137)


if True == tmp17569 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17544 := MakeNative(func(__e *ControlFlow) {
W138 := __e.Get(1)
_ = W138
tmp17545 := MakeNative(func(__e *ControlFlow) {
W139 := __e.Get(1)
_ = W139
tmp17546 := MakeNative(func(__e *ControlFlow) {
W140 := __e.Get(1)
_ = W140
tmp17564 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W140)


if True == tmp17564 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17547 := MakeNative(func(__e *ControlFlow) {
W141 := __e.Get(1)
_ = W141
tmp17548 := MakeNative(func(__e *ControlFlow) {
W142 := __e.Get(1)
_ = W142
tmp17549 := MakeNative(func(__e *ControlFlow) {
W143 := __e.Get(1)
_ = W143
tmp17550 := MakeNative(func(__e *ControlFlow) {
W144 := __e.Get(1)
_ = W144
__e.Return(W144)
return
}, 1)

tmp17551 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W135, Nil)
}
__typedArg0 := W135
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17552 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefine, tmp17551)
}
__typedArg0 := symdefine
__typedArg1 := tmp17551
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17553 := Call(__e, PrimFunc(symshen_4c_1rules_1_6shen), W138, W143, W141)


tmp17554 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17553, Nil)
}
__typedArg0 := tmp17553
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17555 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_6, tmp17554)
}
__typedArg0 := sym_1_6
__typedArg1 := tmp17554
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17556 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W143, tmp17555)
}
__typedArg0 := W143
__typedArg1 := tmp17555
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17557 := Call(__e, PrimFunc(symappend), W138, tmp17556)


tmp17558 := Call(__e, PrimFunc(symappend), tmp17552, tmp17557)


__e.TailApply(tmp17550, tmp17558)
return


}, 1)

tmp17559 := Call(__e, PrimFunc(symgensym), symS)


tmp17560 := Call(__e, tmp17549, tmp17559)


__e.TailApply(PrimFunc(symshen_4comb), W142, tmp17560)
return


}, 1)

tmp17561 := Call(__e, PrimFunc(symshen_4in_1_6), W140)


__e.TailApply(tmp17548, tmp17561)
return


}, 1)

tmp17562 := Call(__e, PrimFunc(symshen_4_5_1out), W140)


__e.TailApply(tmp17547, tmp17562)
return


}


}, 1)

tmp17565 := Call(__e, PrimFunc(symshen_4_5c_1rules_6), W139)


__e.TailApply(tmp17546, tmp17565)
return


}, 1)

tmp17566 := Call(__e, PrimFunc(symshen_4in_1_6), W137)


__e.TailApply(tmp17545, tmp17566)
return


}, 1)

tmp17567 := Call(__e, PrimFunc(symshen_4_5_1out), W137)


__e.TailApply(tmp17544, tmp17567)
return


}


}, 1)

tmp17570 := Call(__e, PrimFunc(symshen_4_5yaccsig_6), W136)


__e.TailApply(tmp17543, tmp17570)
return


}, 1)

tmp17571 := Call(__e, PrimFunc(symtail), V133)


__e.TailApply(tmp17542, tmp17571)
return


}, 1)

tmp17572 := Call(__e, PrimFunc(symhead), V133)


tmp17573 := Call(__e, tmp17541, tmp17572)


ifres17540 = tmp17573


} else {
tmp17574 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17540 = tmp17574


}

__e.TailApply(tmp17537, ifres17540)
return


}, 1)

tmp17576 := Call(__e, ns2_1set, symshen_4_5yacc_6, tmp17536)


_ = tmp17576

tmp17577 := MakeNative(func(__e *ControlFlow) {
V145 := __e.Get(1)
_ = V145
tmp17578 := MakeNative(func(__e *ControlFlow) {
W146 := __e.Get(1)
_ = W146
tmp17590 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W146)


if True == tmp17590 {
tmp17579 := MakeNative(func(__e *ControlFlow) {
W161 := __e.Get(1)
_ = W161
tmp17581 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W161)


if True == tmp17581 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W161)
return
}


}, 1)

tmp17582 := MakeNative(func(__e *ControlFlow) {
W162 := __e.Get(1)
_ = W162
tmp17586 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W162)


if True == tmp17586 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17583 := MakeNative(func(__e *ControlFlow) {
W163 := __e.Get(1)
_ = W163
__e.TailApply(PrimFunc(symshen_4comb), W163, Nil)
return
}, 1)

tmp17584 := Call(__e, PrimFunc(symshen_4in_1_6), W162)


__e.TailApply(tmp17583, tmp17584)
return


}


}, 1)

tmp17587 := Call(__e, PrimFunc(sym_5e_6), V145)


tmp17588 := Call(__e, tmp17582, tmp17587)


__e.TailApply(tmp17579, tmp17588)
return


} else {
__e.Return(W146)
return
}


}, 1)

tmp17653 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V145)
}
__typedArg0 := V145
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17591 Obj

if True == tmp17653 {
tmp17592 := MakeNative(func(__e *ControlFlow) {
W147 := __e.Get(1)
_ = W147
tmp17593 := MakeNative(func(__e *ControlFlow) {
W148 := __e.Get(1)
_ = W148
tmp17648 := Call(__e, PrimFunc(symshen_4ccons_2), W148)


if True == tmp17648 {
tmp17594 := MakeNative(func(__e *ControlFlow) {
W149 := __e.Get(1)
_ = W149
tmp17595 := MakeNative(func(__e *ControlFlow) {
W150 := __e.Get(1)
_ = W150
tmp17644 := Call(__e, PrimFunc(symshen_4hds_a_2), W149, symlist)


if True == tmp17644 {
tmp17596 := MakeNative(func(__e *ControlFlow) {
W151 := __e.Get(1)
_ = W151
tmp17641 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W151)
}
__typedArg0 := W151
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17641 {
tmp17597 := MakeNative(func(__e *ControlFlow) {
W152 := __e.Get(1)
_ = W152
tmp17598 := MakeNative(func(__e *ControlFlow) {
W153 := __e.Get(1)
_ = W153
tmp17599 := MakeNative(func(__e *ControlFlow) {
W154 := __e.Get(1)
_ = W154
tmp17636 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W154)


if True == tmp17636 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17600 := MakeNative(func(__e *ControlFlow) {
W155 := __e.Get(1)
_ = W155
tmp17633 := Call(__e, PrimFunc(symshen_4hds_a_2), W150, sym_a_a_6)


if True == tmp17633 {
tmp17601 := MakeNative(func(__e *ControlFlow) {
W156 := __e.Get(1)
_ = W156
tmp17630 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W156)
}
__typedArg0 := W156
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17630 {
tmp17602 := MakeNative(func(__e *ControlFlow) {
W157 := __e.Get(1)
_ = W157
tmp17603 := MakeNative(func(__e *ControlFlow) {
W158 := __e.Get(1)
_ = W158
tmp17626 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W158)
}
__typedArg0 := W158
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17626 {
tmp17604 := MakeNative(func(__e *ControlFlow) {
W159 := __e.Get(1)
_ = W159
tmp17605 := MakeNative(func(__e *ControlFlow) {
W160 := __e.Get(1)
_ = W160
tmp17622 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_i, W147)
}
__typedArg0 := sym_i
__typedArg1 := W147
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17619 Obj

if True == tmp17622 {
tmp17621 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_j, W159)
}
__typedArg0 := sym_j
__typedArg1 := W159
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17620 Obj

if True == tmp17621 {
ifres17620 = True


} else {
ifres17620 = False


}

ifres17619 = ifres17620


} else {
ifres17619 = False


}

if True == ifres17619 {
tmp17606 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W152, Nil)
}
__typedArg0 := W152
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17607 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp17606)
}
__typedArg0 := symlist
__typedArg1 := tmp17606
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17608 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W152, Nil)
}
__typedArg0 := W152
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17609 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp17608)
}
__typedArg0 := symlist
__typedArg1 := tmp17608
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17610 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W157, Nil)
}
__typedArg0 := W157
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17611 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17609, tmp17610)
}
__typedArg0 := tmp17609
__typedArg1 := tmp17610
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17612 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstr, tmp17611)
}
__typedArg0 := symstr
__typedArg1 := tmp17611
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17613 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_j, Nil)
}
__typedArg0 := sym_j
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17614 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17612, tmp17613)
}
__typedArg0 := tmp17612
__typedArg1 := tmp17613
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17615 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp17614)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp17614
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17616 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17607, tmp17615)
}
__typedArg0 := tmp17607
__typedArg1 := tmp17615
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17617 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_i, tmp17616)
}
__typedArg0 := sym_i
__typedArg1 := tmp17616
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4comb), W160, tmp17617)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp17623 := Call(__e, PrimFunc(symtail), W158)


__e.TailApply(tmp17605, tmp17623)
return


}, 1)

tmp17624 := Call(__e, PrimFunc(symhead), W158)


__e.TailApply(tmp17604, tmp17624)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp17627 := Call(__e, PrimFunc(symtail), W156)


__e.TailApply(tmp17603, tmp17627)
return


}, 1)

tmp17628 := Call(__e, PrimFunc(symhead), W156)


__e.TailApply(tmp17602, tmp17628)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp17631 := Call(__e, PrimFunc(symtail), W150)


__e.TailApply(tmp17601, tmp17631)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp17634 := Call(__e, PrimFunc(symshen_4in_1_6), W154)


__e.TailApply(tmp17600, tmp17634)
return


}


}, 1)

tmp17637 := Call(__e, PrimFunc(sym_5end_6), W153)


__e.TailApply(tmp17599, tmp17637)
return


}, 1)

tmp17638 := Call(__e, PrimFunc(symtail), W151)


__e.TailApply(tmp17598, tmp17638)
return


}, 1)

tmp17639 := Call(__e, PrimFunc(symhead), W151)


__e.TailApply(tmp17597, tmp17639)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp17642 := Call(__e, PrimFunc(symtail), W149)


__e.TailApply(tmp17596, tmp17642)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp17645 := Call(__e, PrimFunc(symtail), W148)


__e.TailApply(tmp17595, tmp17645)
return


}, 1)

tmp17646 := Call(__e, PrimFunc(symhead), W148)


__e.TailApply(tmp17594, tmp17646)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp17649 := Call(__e, PrimFunc(symtail), V145)


__e.TailApply(tmp17593, tmp17649)
return


}, 1)

tmp17650 := Call(__e, PrimFunc(symhead), V145)


tmp17651 := Call(__e, tmp17592, tmp17650)


ifres17591 = tmp17651


} else {
tmp17652 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17591 = tmp17652


}

__e.TailApply(tmp17578, ifres17591)
return


}, 1)

tmp17654 := Call(__e, ns2_1set, symshen_4_5yaccsig_6, tmp17577)


_ = tmp17654

tmp17655 := MakeNative(func(__e *ControlFlow) {
V164 := __e.Get(1)
_ = V164
tmp17656 := MakeNative(func(__e *ControlFlow) {
W165 := __e.Get(1)
_ = W165
tmp17675 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W165)


if True == tmp17675 {
tmp17657 := MakeNative(func(__e *ControlFlow) {
W172 := __e.Get(1)
_ = W172
tmp17659 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W172)


if True == tmp17659 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W172)
return
}


}, 1)

tmp17660 := MakeNative(func(__e *ControlFlow) {
W173 := __e.Get(1)
_ = W173
tmp17671 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W173)


if True == tmp17671 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17661 := MakeNative(func(__e *ControlFlow) {
W174 := __e.Get(1)
_ = W174
tmp17662 := MakeNative(func(__e *ControlFlow) {
W175 := __e.Get(1)
_ = W175
tmp17667 := Call(__e, PrimFunc(symempty_2), W174)


var ifres17663 Obj

if True == tmp17667 {
ifres17663 = Nil


} else {
tmp17664 := Call(__e, PrimFunc(symshen_4app), W174, MakeString("\n ..."), symshen_4r)


tmp17666 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("YACC syntax error here:\n "))
__typedS1, __typedOK1 := TypedString(tmp17664)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("YACC syntax error here:\n ")
__typedArg1 := tmp17664
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("YACC syntax error here:\n "))
__typedS1, __typedOK1 := TypedString(tmp17664)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("YACC syntax error here:\n ")
__typedArg1 := tmp17664
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})()

ifres17663 = tmp17666


}

__e.TailApply(PrimFunc(symshen_4comb), W175, ifres17663)
return


}, 1)

tmp17668 := Call(__e, PrimFunc(symshen_4in_1_6), W173)


__e.TailApply(tmp17662, tmp17668)
return


}, 1)

tmp17669 := Call(__e, PrimFunc(symshen_4_5_1out), W173)


__e.TailApply(tmp17661, tmp17669)
return


}


}, 1)

tmp17672 := Call(__e, PrimFunc(sym_5_b_6), V164)


tmp17673 := Call(__e, tmp17660, tmp17672)


__e.TailApply(tmp17657, tmp17673)
return


} else {
__e.Return(W165)
return
}


}, 1)

tmp17676 := MakeNative(func(__e *ControlFlow) {
W166 := __e.Get(1)
_ = W166
tmp17691 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W166)


if True == tmp17691 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17677 := MakeNative(func(__e *ControlFlow) {
W167 := __e.Get(1)
_ = W167
tmp17678 := MakeNative(func(__e *ControlFlow) {
W168 := __e.Get(1)
_ = W168
tmp17679 := MakeNative(func(__e *ControlFlow) {
W169 := __e.Get(1)
_ = W169
tmp17686 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W169)


if True == tmp17686 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17680 := MakeNative(func(__e *ControlFlow) {
W170 := __e.Get(1)
_ = W170
tmp17681 := MakeNative(func(__e *ControlFlow) {
W171 := __e.Get(1)
_ = W171
tmp17682 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W167, W170)
}
__typedArg0 := W167
__typedArg1 := W170
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4comb), W171, tmp17682)
return


}, 1)

tmp17683 := Call(__e, PrimFunc(symshen_4in_1_6), W169)


__e.TailApply(tmp17681, tmp17683)
return


}, 1)

tmp17684 := Call(__e, PrimFunc(symshen_4_5_1out), W169)


__e.TailApply(tmp17680, tmp17684)
return


}


}, 1)

tmp17687 := Call(__e, PrimFunc(symshen_4_5c_1rules_6), W168)


__e.TailApply(tmp17679, tmp17687)
return


}, 1)

tmp17688 := Call(__e, PrimFunc(symshen_4in_1_6), W166)


__e.TailApply(tmp17678, tmp17688)
return


}, 1)

tmp17689 := Call(__e, PrimFunc(symshen_4_5_1out), W166)


__e.TailApply(tmp17677, tmp17689)
return


}


}, 1)

tmp17692 := Call(__e, PrimFunc(symshen_4_5c_1rule_6), V164)


tmp17693 := Call(__e, tmp17676, tmp17692)


__e.TailApply(tmp17656, tmp17693)
return


}, 1)

tmp17694 := Call(__e, ns2_1set, symshen_4_5c_1rules_6, tmp17655)


_ = tmp17694

tmp17695 := MakeNative(func(__e *ControlFlow) {
V176 := __e.Get(1)
_ = V176
tmp17696 := MakeNative(func(__e *ControlFlow) {
W177 := __e.Get(1)
_ = W177
tmp17719 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W177)


if True == tmp17719 {
tmp17697 := MakeNative(func(__e *ControlFlow) {
W186 := __e.Get(1)
_ = W186
tmp17699 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W186)


if True == tmp17699 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W186)
return
}


}, 1)

tmp17700 := MakeNative(func(__e *ControlFlow) {
W187 := __e.Get(1)
_ = W187
tmp17715 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W187)


if True == tmp17715 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17701 := MakeNative(func(__e *ControlFlow) {
W188 := __e.Get(1)
_ = W188
tmp17702 := MakeNative(func(__e *ControlFlow) {
W189 := __e.Get(1)
_ = W189
tmp17703 := MakeNative(func(__e *ControlFlow) {
W190 := __e.Get(1)
_ = W190
tmp17710 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W190)


if True == tmp17710 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17704 := MakeNative(func(__e *ControlFlow) {
W191 := __e.Get(1)
_ = W191
tmp17705 := Call(__e, PrimFunc(symshen_4autocomplete), W188)


tmp17706 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17705, Nil)
}
__typedArg0 := tmp17705
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17707 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W188, tmp17706)
}
__typedArg0 := W188
__typedArg1 := tmp17706
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4comb), W191, tmp17707)
return


}, 1)

tmp17708 := Call(__e, PrimFunc(symshen_4in_1_6), W190)


__e.TailApply(tmp17704, tmp17708)
return


}


}, 1)

tmp17711 := Call(__e, PrimFunc(symshen_4_5sc_6), W189)


__e.TailApply(tmp17703, tmp17711)
return


}, 1)

tmp17712 := Call(__e, PrimFunc(symshen_4in_1_6), W187)


__e.TailApply(tmp17702, tmp17712)
return


}, 1)

tmp17713 := Call(__e, PrimFunc(symshen_4_5_1out), W187)


__e.TailApply(tmp17701, tmp17713)
return


}


}, 1)

tmp17716 := Call(__e, PrimFunc(symshen_4_5syntax_6), V176)


tmp17717 := Call(__e, tmp17700, tmp17716)


__e.TailApply(tmp17697, tmp17717)
return


} else {
__e.Return(W177)
return
}


}, 1)

tmp17720 := MakeNative(func(__e *ControlFlow) {
W178 := __e.Get(1)
_ = W178
tmp17742 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W178)


if True == tmp17742 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17721 := MakeNative(func(__e *ControlFlow) {
W179 := __e.Get(1)
_ = W179
tmp17722 := MakeNative(func(__e *ControlFlow) {
W180 := __e.Get(1)
_ = W180
tmp17723 := MakeNative(func(__e *ControlFlow) {
W181 := __e.Get(1)
_ = W181
tmp17737 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W181)


if True == tmp17737 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17724 := MakeNative(func(__e *ControlFlow) {
W182 := __e.Get(1)
_ = W182
tmp17725 := MakeNative(func(__e *ControlFlow) {
W183 := __e.Get(1)
_ = W183
tmp17726 := MakeNative(func(__e *ControlFlow) {
W184 := __e.Get(1)
_ = W184
tmp17732 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W184)


if True == tmp17732 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17727 := MakeNative(func(__e *ControlFlow) {
W185 := __e.Get(1)
_ = W185
tmp17728 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W182, Nil)
}
__typedArg0 := W182
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17729 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W179, tmp17728)
}
__typedArg0 := W179
__typedArg1 := tmp17728
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4comb), W185, tmp17729)
return


}, 1)

tmp17730 := Call(__e, PrimFunc(symshen_4in_1_6), W184)


__e.TailApply(tmp17727, tmp17730)
return


}


}, 1)

tmp17733 := Call(__e, PrimFunc(symshen_4_5sc_6), W183)


__e.TailApply(tmp17726, tmp17733)
return


}, 1)

tmp17734 := Call(__e, PrimFunc(symshen_4in_1_6), W181)


__e.TailApply(tmp17725, tmp17734)
return


}, 1)

tmp17735 := Call(__e, PrimFunc(symshen_4_5_1out), W181)


__e.TailApply(tmp17724, tmp17735)
return


}


}, 1)

tmp17738 := Call(__e, PrimFunc(symshen_4_5semantics_6), W180)


__e.TailApply(tmp17723, tmp17738)
return


}, 1)

tmp17739 := Call(__e, PrimFunc(symshen_4in_1_6), W178)


__e.TailApply(tmp17722, tmp17739)
return


}, 1)

tmp17740 := Call(__e, PrimFunc(symshen_4_5_1out), W178)


__e.TailApply(tmp17721, tmp17740)
return


}


}, 1)

tmp17743 := Call(__e, PrimFunc(symshen_4_5syntax_6), V176)


tmp17744 := Call(__e, tmp17720, tmp17743)


__e.TailApply(tmp17696, tmp17744)
return


}, 1)

tmp17745 := Call(__e, ns2_1set, symshen_4_5c_1rule_6, tmp17695)


_ = tmp17745

tmp17746 := MakeNative(func(__e *ControlFlow) {
V192 := __e.Get(1)
_ = V192
tmp17775 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V192)
}
__typedArg0 := V192
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17767 Obj

if True == tmp17775 {
tmp17773 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V192)
}
__typedArg0 := V192
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17774 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp17773)
}
__typedArg0 := Nil
__typedArg1 := tmp17773
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17769 Obj

if True == tmp17774 {
tmp17771 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V192)
}
__typedArg0 := V192
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17772 := Call(__e, PrimFunc(symshen_4non_1terminal_2), tmp17771)


var ifres17770 Obj

if True == tmp17772 {
ifres17770 = True


} else {
ifres17770 = False


}

ifres17769 = ifres17770


} else {
ifres17769 = False


}

var ifres17768 Obj

if True == ifres17769 {
ifres17768 = True


} else {
ifres17768 = False


}

ifres17767 = ifres17768


} else {
ifres17767 = False


}

if True == ifres17767 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V192)
}
__typedArg0 := V192
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return
} else {
tmp17765 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V192)
}
__typedArg0 := V192
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17761 Obj

if True == tmp17765 {
tmp17763 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V192)
}
__typedArg0 := V192
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17764 := Call(__e, PrimFunc(symshen_4non_1terminal_2), tmp17763)


var ifres17762 Obj

if True == tmp17764 {
ifres17762 = True


} else {
ifres17762 = False


}

ifres17761 = ifres17762


} else {
ifres17761 = False


}

if True == ifres17761 {
tmp17747 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V192)
}
__typedArg0 := V192
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17748 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V192)
}
__typedArg0 := V192
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17749 := Call(__e, PrimFunc(symshen_4autocomplete), tmp17748)


tmp17750 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17749, Nil)
}
__typedArg0 := tmp17749
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17751 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17747, tmp17750)
}
__typedArg0 := tmp17747
__typedArg1 := tmp17750
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symappend, tmp17751)
}
__typedArg0 := symappend
__typedArg1 := tmp17751
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp17759 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V192)
}
__typedArg0 := V192
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17759 {
tmp17752 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V192)
}
__typedArg0 := V192
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17753 := Call(__e, PrimFunc(symshen_4autocomplete), tmp17752)


tmp17754 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V192)
}
__typedArg0 := V192
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17755 := Call(__e, PrimFunc(symshen_4autocomplete), tmp17754)


tmp17756 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17755, Nil)
}
__typedArg0 := tmp17755
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17757 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17753, tmp17756)
}
__typedArg0 := tmp17753
__typedArg1 := tmp17756
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp17757)
}
__typedArg0 := symcons
__typedArg1 := tmp17757
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

tmp17776 := Call(__e, ns2_1set, symshen_4autocomplete, tmp17746)


_ = tmp17776

tmp17777 := MakeNative(func(__e *ControlFlow) {
V193 := __e.Get(1)
_ = V193
tmp17784 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsymbol_2) {
return PrimIsSymbol(V193)
}
__typedArg0 := V193
return Call(__e, PrimFunc(symsymbol_2), __typedArg0)
})()

if True == tmp17784 {
tmp17779 := MakeNative(func(__e *ControlFlow) {
W194 := __e.Get(1)
_ = W194
tmp17780 := MakeNative(func(__e *ControlFlow) {
Z195 := __e.Get(1)
_ = Z195
__e.TailApply(PrimFunc(symshen_4_5non_1terminal_2_6), Z195)
return
}, 1)

__e.TailApply(PrimFunc(symcompile), tmp17780, W194)
return


}, 1)

tmp17781 := Call(__e, PrimFunc(symexplode), V193)


tmp17782 := Call(__e, tmp17779, tmp17781)


if True == tmp17782 {
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

tmp17785 := Call(__e, ns2_1set, symshen_4non_1terminal_2, tmp17777)


_ = tmp17785

tmp17786 := MakeNative(func(__e *ControlFlow) {
V196 := __e.Get(1)
_ = V196
tmp17787 := MakeNative(func(__e *ControlFlow) {
W197 := __e.Get(1)
_ = W197
tmp17809 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W197)


if True == tmp17809 {
tmp17788 := MakeNative(func(__e *ControlFlow) {
W202 := __e.Get(1)
_ = W202
tmp17800 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W202)


if True == tmp17800 {
tmp17789 := MakeNative(func(__e *ControlFlow) {
W205 := __e.Get(1)
_ = W205
tmp17791 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W205)


if True == tmp17791 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W205)
return
}


}, 1)

tmp17792 := MakeNative(func(__e *ControlFlow) {
W206 := __e.Get(1)
_ = W206
tmp17796 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W206)


if True == tmp17796 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17793 := MakeNative(func(__e *ControlFlow) {
W207 := __e.Get(1)
_ = W207
__e.TailApply(PrimFunc(symshen_4comb), W207, False)
return
}, 1)

tmp17794 := Call(__e, PrimFunc(symshen_4in_1_6), W206)


__e.TailApply(tmp17793, tmp17794)
return


}


}, 1)

tmp17797 := Call(__e, PrimFunc(sym_5_b_6), V196)


tmp17798 := Call(__e, tmp17792, tmp17797)


__e.TailApply(tmp17789, tmp17798)
return


} else {
__e.Return(W202)
return
}


}, 1)

tmp17801 := MakeNative(func(__e *ControlFlow) {
W203 := __e.Get(1)
_ = W203
tmp17805 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W203)


if True == tmp17805 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17802 := MakeNative(func(__e *ControlFlow) {
W204 := __e.Get(1)
_ = W204
__e.TailApply(PrimFunc(symshen_4comb), W204, True)
return
}, 1)

tmp17803 := Call(__e, PrimFunc(symshen_4in_1_6), W203)


__e.TailApply(tmp17802, tmp17803)
return


}


}, 1)

tmp17806 := Call(__e, PrimFunc(symshen_4_5non_1terminal_1name_6), V196)


tmp17807 := Call(__e, tmp17801, tmp17806)


__e.TailApply(tmp17788, tmp17807)
return


} else {
__e.Return(W197)
return
}


}, 1)

tmp17810 := MakeNative(func(__e *ControlFlow) {
W198 := __e.Get(1)
_ = W198
tmp17820 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W198)


if True == tmp17820 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17811 := MakeNative(func(__e *ControlFlow) {
W199 := __e.Get(1)
_ = W199
tmp17812 := MakeNative(func(__e *ControlFlow) {
W200 := __e.Get(1)
_ = W200
tmp17816 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W200)


if True == tmp17816 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17813 := MakeNative(func(__e *ControlFlow) {
W201 := __e.Get(1)
_ = W201
__e.TailApply(PrimFunc(symshen_4comb), W201, True)
return
}, 1)

tmp17814 := Call(__e, PrimFunc(symshen_4in_1_6), W200)


__e.TailApply(tmp17813, tmp17814)
return


}


}, 1)

tmp17817 := Call(__e, PrimFunc(symshen_4_5non_1terminal_1name_6), W199)


__e.TailApply(tmp17812, tmp17817)
return


}, 1)

tmp17818 := Call(__e, PrimFunc(symshen_4in_1_6), W198)


__e.TailApply(tmp17811, tmp17818)
return


}


}, 1)

tmp17821 := Call(__e, PrimFunc(symshen_4_5packagenames_6), V196)


tmp17822 := Call(__e, tmp17810, tmp17821)


__e.TailApply(tmp17787, tmp17822)
return


}, 1)

tmp17823 := Call(__e, ns2_1set, symshen_4_5non_1terminal_2_6, tmp17786)


_ = tmp17823

tmp17824 := MakeNative(func(__e *ControlFlow) {
V208 := __e.Get(1)
_ = V208
tmp17825 := MakeNative(func(__e *ControlFlow) {
W209 := __e.Get(1)
_ = W209
tmp17841 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W209)


if True == tmp17841 {
tmp17826 := MakeNative(func(__e *ControlFlow) {
W215 := __e.Get(1)
_ = W215
tmp17828 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W215)


if True == tmp17828 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W215)
return
}


}, 1)

tmp17829 := MakeNative(func(__e *ControlFlow) {
W216 := __e.Get(1)
_ = W216
tmp17837 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W216)


if True == tmp17837 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17830 := MakeNative(func(__e *ControlFlow) {
W217 := __e.Get(1)
_ = W217
tmp17834 := Call(__e, PrimFunc(symshen_4hds_a_2), W217, MakeString("."))


if True == tmp17834 {
tmp17831 := MakeNative(func(__e *ControlFlow) {
W218 := __e.Get(1)
_ = W218
__e.TailApply(PrimFunc(symshen_4comb), W218, symshen_4skip)
return
}, 1)

tmp17832 := Call(__e, PrimFunc(symtail), W217)


__e.TailApply(tmp17831, tmp17832)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp17835 := Call(__e, PrimFunc(symshen_4in_1_6), W216)


__e.TailApply(tmp17830, tmp17835)
return


}


}, 1)

tmp17838 := Call(__e, PrimFunc(symshen_4_5packagename_6), V208)


tmp17839 := Call(__e, tmp17829, tmp17838)


__e.TailApply(tmp17826, tmp17839)
return


} else {
__e.Return(W209)
return
}


}, 1)

tmp17842 := MakeNative(func(__e *ControlFlow) {
W210 := __e.Get(1)
_ = W210
tmp17856 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W210)


if True == tmp17856 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17843 := MakeNative(func(__e *ControlFlow) {
W211 := __e.Get(1)
_ = W211
tmp17853 := Call(__e, PrimFunc(symshen_4hds_a_2), W211, MakeString("."))


if True == tmp17853 {
tmp17844 := MakeNative(func(__e *ControlFlow) {
W212 := __e.Get(1)
_ = W212
tmp17845 := MakeNative(func(__e *ControlFlow) {
W213 := __e.Get(1)
_ = W213
tmp17849 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W213)


if True == tmp17849 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17846 := MakeNative(func(__e *ControlFlow) {
W214 := __e.Get(1)
_ = W214
__e.TailApply(PrimFunc(symshen_4comb), W214, symshen_4skip)
return
}, 1)

tmp17847 := Call(__e, PrimFunc(symshen_4in_1_6), W213)


__e.TailApply(tmp17846, tmp17847)
return


}


}, 1)

tmp17850 := Call(__e, PrimFunc(symshen_4_5packagenames_6), W212)


__e.TailApply(tmp17845, tmp17850)
return


}, 1)

tmp17851 := Call(__e, PrimFunc(symtail), W211)


__e.TailApply(tmp17844, tmp17851)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp17854 := Call(__e, PrimFunc(symshen_4in_1_6), W210)


__e.TailApply(tmp17843, tmp17854)
return


}


}, 1)

tmp17857 := Call(__e, PrimFunc(symshen_4_5packagename_6), V208)


tmp17858 := Call(__e, tmp17842, tmp17857)


__e.TailApply(tmp17825, tmp17858)
return


}, 1)

tmp17859 := Call(__e, ns2_1set, symshen_4_5packagenames_6, tmp17824)


_ = tmp17859

tmp17860 := MakeNative(func(__e *ControlFlow) {
V219 := __e.Get(1)
_ = V219
tmp17861 := MakeNative(func(__e *ControlFlow) {
W220 := __e.Get(1)
_ = W220
tmp17873 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W220)


if True == tmp17873 {
tmp17862 := MakeNative(func(__e *ControlFlow) {
W225 := __e.Get(1)
_ = W225
tmp17864 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W225)


if True == tmp17864 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W225)
return
}


}, 1)

tmp17865 := MakeNative(func(__e *ControlFlow) {
W226 := __e.Get(1)
_ = W226
tmp17869 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W226)


if True == tmp17869 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17866 := MakeNative(func(__e *ControlFlow) {
W227 := __e.Get(1)
_ = W227
__e.TailApply(PrimFunc(symshen_4comb), W227, symshen_4skip)
return
}, 1)

tmp17867 := Call(__e, PrimFunc(symshen_4in_1_6), W226)


__e.TailApply(tmp17866, tmp17867)
return


}


}, 1)

tmp17870 := Call(__e, PrimFunc(sym_5e_6), V219)


tmp17871 := Call(__e, tmp17865, tmp17870)


__e.TailApply(tmp17862, tmp17871)
return


} else {
__e.Return(W220)
return
}


}, 1)

tmp17874 := MakeNative(func(__e *ControlFlow) {
W221 := __e.Get(1)
_ = W221
tmp17884 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W221)


if True == tmp17884 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17875 := MakeNative(func(__e *ControlFlow) {
W222 := __e.Get(1)
_ = W222
tmp17876 := MakeNative(func(__e *ControlFlow) {
W223 := __e.Get(1)
_ = W223
tmp17880 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W223)


if True == tmp17880 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17877 := MakeNative(func(__e *ControlFlow) {
W224 := __e.Get(1)
_ = W224
__e.TailApply(PrimFunc(symshen_4comb), W224, symshen_4skip)
return
}, 1)

tmp17878 := Call(__e, PrimFunc(symshen_4in_1_6), W223)


__e.TailApply(tmp17877, tmp17878)
return


}


}, 1)

tmp17881 := Call(__e, PrimFunc(symshen_4_5packagename_6), W222)


__e.TailApply(tmp17876, tmp17881)
return


}, 1)

tmp17882 := Call(__e, PrimFunc(symshen_4in_1_6), W221)


__e.TailApply(tmp17875, tmp17882)
return


}


}, 1)

tmp17885 := Call(__e, PrimFunc(symshen_4_5packagechar_6), V219)


tmp17886 := Call(__e, tmp17874, tmp17885)


__e.TailApply(tmp17861, tmp17886)
return


}, 1)

tmp17887 := Call(__e, ns2_1set, symshen_4_5packagename_6, tmp17860)


_ = tmp17887

tmp17888 := MakeNative(func(__e *ControlFlow) {
V228 := __e.Get(1)
_ = V228
tmp17889 := MakeNative(func(__e *ControlFlow) {
W229 := __e.Get(1)
_ = W229
tmp17891 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W229)


if True == tmp17891 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W229)
return
}


}, 1)

tmp17902 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V228)
}
__typedArg0 := V228
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17892 Obj

if True == tmp17902 {
tmp17893 := MakeNative(func(__e *ControlFlow) {
W230 := __e.Get(1)
_ = W230
tmp17894 := MakeNative(func(__e *ControlFlow) {
W231 := __e.Get(1)
_ = W231
tmp17896 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W230, MakeString("."))
}
__typedArg0 := W230
__typedArg1 := MakeString(".")
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp17896)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp17896
return Call(__e, PrimFunc(symnot), __typedArg0)
})() {
__e.TailApply(PrimFunc(symshen_4comb), W231, symshen_4skip)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp17898 := Call(__e, PrimFunc(symtail), V228)


__e.TailApply(tmp17894, tmp17898)
return


}, 1)

tmp17899 := Call(__e, PrimFunc(symhead), V228)


tmp17900 := Call(__e, tmp17893, tmp17899)


ifres17892 = tmp17900


} else {
tmp17901 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17892 = tmp17901


}

__e.TailApply(tmp17889, ifres17892)
return


}, 1)

tmp17903 := Call(__e, ns2_1set, symshen_4_5packagechar_6, tmp17888)


_ = tmp17903

tmp17904 := MakeNative(func(__e *ControlFlow) {
V232 := __e.Get(1)
_ = V232
tmp17905 := MakeNative(func(__e *ControlFlow) {
W233 := __e.Get(1)
_ = W233
tmp17907 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W233)


if True == tmp17907 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W233)
return
}


}, 1)

tmp17930 := Call(__e, PrimFunc(symshen_4hds_a_2), V232, MakeString("<"))


var ifres17908 Obj

if True == tmp17930 {
tmp17909 := MakeNative(func(__e *ControlFlow) {
W234 := __e.Get(1)
_ = W234
tmp17910 := MakeNative(func(__e *ControlFlow) {
W235 := __e.Get(1)
_ = W235
tmp17925 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W235)


if True == tmp17925 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17911 := MakeNative(func(__e *ControlFlow) {
W236 := __e.Get(1)
_ = W236
tmp17912 := MakeNative(func(__e *ControlFlow) {
W237 := __e.Get(1)
_ = W237
tmp17914 := MakeNative(func(__e *ControlFlow) {
W238 := __e.Get(1)
_ = W238
tmp17919 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W238)
}
__typedArg0 := W238
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17919 {
tmp17916 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W238)
}
__typedArg0 := W238
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

if True == tmp17917 {
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

tmp17920 := Call(__e, PrimFunc(symreverse), W236)


tmp17921 := Call(__e, tmp17914, tmp17920)


if True == tmp17921 {
__e.TailApply(PrimFunc(symshen_4comb), W237, symshen_4skip)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp17922 := Call(__e, PrimFunc(symshen_4in_1_6), W235)


__e.TailApply(tmp17912, tmp17922)
return


}, 1)

tmp17923 := Call(__e, PrimFunc(symshen_4_5_1out), W235)


__e.TailApply(tmp17911, tmp17923)
return


}


}, 1)

tmp17926 := Call(__e, PrimFunc(sym_5_b_6), W234)


__e.TailApply(tmp17910, tmp17926)
return


}, 1)

tmp17927 := Call(__e, PrimFunc(symtail), V232)


tmp17928 := Call(__e, tmp17909, tmp17927)


ifres17908 = tmp17928


} else {
tmp17929 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17908 = tmp17929


}

__e.TailApply(tmp17905, ifres17908)
return


}, 1)

tmp17931 := Call(__e, ns2_1set, symshen_4_5non_1terminal_1name_6, tmp17904)


_ = tmp17931

tmp17932 := MakeNative(func(__e *ControlFlow) {
V239 := __e.Get(1)
_ = V239
tmp17933 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(";"))
}
__typedArg0 := MakeString(";")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V239, tmp17933)
}
__typedArg0 := V239
__typedArg1 := tmp17933
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp17934 := Call(__e, ns2_1set, symshen_4semicolon_2, tmp17932)


_ = tmp17934

tmp17935 := MakeNative(func(__e *ControlFlow) {
V240 := __e.Get(1)
_ = V240
tmp17936 := MakeNative(func(__e *ControlFlow) {
W241 := __e.Get(1)
_ = W241
tmp17938 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W241)


if True == tmp17938 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W241)
return
}


}, 1)

tmp17948 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V240)
}
__typedArg0 := V240
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17939 Obj

if True == tmp17948 {
tmp17940 := MakeNative(func(__e *ControlFlow) {
W242 := __e.Get(1)
_ = W242
tmp17941 := MakeNative(func(__e *ControlFlow) {
W243 := __e.Get(1)
_ = W243
tmp17943 := Call(__e, PrimFunc(symshen_4colon_1equal_2), W242)


if True == tmp17943 {
__e.TailApply(PrimFunc(symshen_4comb), W243, symshen_4skip)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp17944 := Call(__e, PrimFunc(symtail), V240)


__e.TailApply(tmp17941, tmp17944)
return


}, 1)

tmp17945 := Call(__e, PrimFunc(symhead), V240)


tmp17946 := Call(__e, tmp17940, tmp17945)


ifres17939 = tmp17946


} else {
tmp17947 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17939 = tmp17947


}

__e.TailApply(tmp17936, ifres17939)
return


}, 1)

tmp17949 := Call(__e, ns2_1set, symshen_4_5colon_1equal_6, tmp17935)


_ = tmp17949

tmp17950 := MakeNative(func(__e *ControlFlow) {
V244 := __e.Get(1)
_ = V244
tmp17951 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":="))
}
__typedArg0 := MakeString(":=")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp17951, V244)
}
__typedArg0 := tmp17951
__typedArg1 := V244
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp17952 := Call(__e, ns2_1set, symshen_4colon_1equal_2, tmp17950)


_ = tmp17952

tmp17953 := MakeNative(func(__e *ControlFlow) {
V245 := __e.Get(1)
_ = V245
tmp17954 := MakeNative(func(__e *ControlFlow) {
W246 := __e.Get(1)
_ = W246
tmp17969 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W246)


if True == tmp17969 {
tmp17955 := MakeNative(func(__e *ControlFlow) {
W253 := __e.Get(1)
_ = W253
tmp17957 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W253)


if True == tmp17957 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W253)
return
}


}, 1)

tmp17958 := MakeNative(func(__e *ControlFlow) {
W254 := __e.Get(1)
_ = W254
tmp17965 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W254)


if True == tmp17965 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17959 := MakeNative(func(__e *ControlFlow) {
W255 := __e.Get(1)
_ = W255
tmp17960 := MakeNative(func(__e *ControlFlow) {
W256 := __e.Get(1)
_ = W256
tmp17961 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W255, Nil)
}
__typedArg0 := W255
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4comb), W256, tmp17961)
return


}, 1)

tmp17962 := Call(__e, PrimFunc(symshen_4in_1_6), W254)


__e.TailApply(tmp17960, tmp17962)
return


}, 1)

tmp17963 := Call(__e, PrimFunc(symshen_4_5_1out), W254)


__e.TailApply(tmp17959, tmp17963)
return


}


}, 1)

tmp17966 := Call(__e, PrimFunc(symshen_4_5syntax_1item_6), V245)


tmp17967 := Call(__e, tmp17958, tmp17966)


__e.TailApply(tmp17955, tmp17967)
return


} else {
__e.Return(W246)
return
}


}, 1)

tmp17970 := MakeNative(func(__e *ControlFlow) {
W247 := __e.Get(1)
_ = W247
tmp17985 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W247)


if True == tmp17985 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17971 := MakeNative(func(__e *ControlFlow) {
W248 := __e.Get(1)
_ = W248
tmp17972 := MakeNative(func(__e *ControlFlow) {
W249 := __e.Get(1)
_ = W249
tmp17973 := MakeNative(func(__e *ControlFlow) {
W250 := __e.Get(1)
_ = W250
tmp17980 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W250)


if True == tmp17980 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17974 := MakeNative(func(__e *ControlFlow) {
W251 := __e.Get(1)
_ = W251
tmp17975 := MakeNative(func(__e *ControlFlow) {
W252 := __e.Get(1)
_ = W252
tmp17976 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W248, W251)
}
__typedArg0 := W248
__typedArg1 := W251
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4comb), W252, tmp17976)
return


}, 1)

tmp17977 := Call(__e, PrimFunc(symshen_4in_1_6), W250)


__e.TailApply(tmp17975, tmp17977)
return


}, 1)

tmp17978 := Call(__e, PrimFunc(symshen_4_5_1out), W250)


__e.TailApply(tmp17974, tmp17978)
return


}


}, 1)

tmp17981 := Call(__e, PrimFunc(symshen_4_5syntax_6), W249)


__e.TailApply(tmp17973, tmp17981)
return


}, 1)

tmp17982 := Call(__e, PrimFunc(symshen_4in_1_6), W247)


__e.TailApply(tmp17972, tmp17982)
return


}, 1)

tmp17983 := Call(__e, PrimFunc(symshen_4_5_1out), W247)


__e.TailApply(tmp17971, tmp17983)
return


}


}, 1)

tmp17986 := Call(__e, PrimFunc(symshen_4_5syntax_1item_6), V245)


tmp17987 := Call(__e, tmp17970, tmp17986)


__e.TailApply(tmp17954, tmp17987)
return


}, 1)

tmp17988 := Call(__e, ns2_1set, symshen_4_5syntax_6, tmp17953)


_ = tmp17988

tmp17989 := MakeNative(func(__e *ControlFlow) {
V257 := __e.Get(1)
_ = V257
tmp17990 := MakeNative(func(__e *ControlFlow) {
W258 := __e.Get(1)
_ = W258
tmp17992 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W258)


if True == tmp17992 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W258)
return
}


}, 1)

tmp18002 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V257)
}
__typedArg0 := V257
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17993 Obj

if True == tmp18002 {
tmp17994 := MakeNative(func(__e *ControlFlow) {
W259 := __e.Get(1)
_ = W259
tmp17995 := MakeNative(func(__e *ControlFlow) {
W260 := __e.Get(1)
_ = W260
tmp17997 := Call(__e, PrimFunc(symshen_4syntax_1item_2), W259)


if True == tmp17997 {
__e.TailApply(PrimFunc(symshen_4comb), W260, W259)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp17998 := Call(__e, PrimFunc(symtail), V257)


__e.TailApply(tmp17995, tmp17998)
return


}, 1)

tmp17999 := Call(__e, PrimFunc(symhead), V257)


tmp18000 := Call(__e, tmp17994, tmp17999)


ifres17993 = tmp18000


} else {
tmp18001 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17993 = tmp18001


}

__e.TailApply(tmp17990, ifres17993)
return


}, 1)

tmp18003 := Call(__e, ns2_1set, symshen_4_5syntax_1item_6, tmp17989)


_ = tmp18003

tmp18004 := MakeNative(func(__e *ControlFlow) {
V263 := __e.Get(1)
_ = V263
tmp18040 := Call(__e, PrimFunc(symshen_4colon_1equal_2), V263)


if True == tmp18040 {
__e.Return(False)
return
} else {
tmp18038 := Call(__e, PrimFunc(symshen_4semicolon_2), V263)


if True == tmp18038 {
__e.Return(False)
return
} else {
tmp18036 := Call(__e, PrimFunc(symatom_2), V263)


if True == tmp18036 {
__e.Return(True)
return
} else {
tmp18034 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V263)
}
__typedArg0 := V263
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18015 Obj

if True == tmp18034 {
tmp18032 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V263)
}
__typedArg0 := V263
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18033 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcons, tmp18032)
}
__typedArg0 := symcons
__typedArg1 := tmp18032
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18017 Obj

if True == tmp18033 {
tmp18030 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V263)
}
__typedArg0 := V263
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18031 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18030)
}
__typedArg0 := tmp18030
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18019 Obj

if True == tmp18031 {
tmp18027 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V263)
}
__typedArg0 := V263
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18028 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18027)
}
__typedArg0 := tmp18027
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18029 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18028)
}
__typedArg0 := tmp18028
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18021 Obj

if True == tmp18029 {
tmp18023 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V263)
}
__typedArg0 := V263
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18024 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18023)
}
__typedArg0 := tmp18023
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18025 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18024)
}
__typedArg0 := tmp18024
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18026 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp18025)
}
__typedArg0 := Nil
__typedArg1 := tmp18025
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18022 Obj

if True == tmp18026 {
ifres18022 = True


} else {
ifres18022 = False


}

ifres18021 = ifres18022


} else {
ifres18021 = False


}

var ifres18020 Obj

if True == ifres18021 {
ifres18020 = True


} else {
ifres18020 = False


}

ifres18019 = ifres18020


} else {
ifres18019 = False


}

var ifres18018 Obj

if True == ifres18019 {
ifres18018 = True


} else {
ifres18018 = False


}

ifres18017 = ifres18018


} else {
ifres18017 = False


}

var ifres18016 Obj

if True == ifres18017 {
ifres18016 = True


} else {
ifres18016 = False


}

ifres18015 = ifres18016


} else {
ifres18015 = False


}

if True == ifres18015 {
tmp18011 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V263)
}
__typedArg0 := V263
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18012 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18011)
}
__typedArg0 := tmp18011
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18013 := Call(__e, PrimFunc(symshen_4syntax_1item_2), tmp18012)


if True == tmp18013 {
tmp18006 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V263)
}
__typedArg0 := V263
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18007 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18006)
}
__typedArg0 := tmp18006
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18008 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18007)
}
__typedArg0 := tmp18007
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18009 := Call(__e, PrimFunc(symshen_4syntax_1item_2), tmp18008)


if True == tmp18009 {
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

tmp18041 := Call(__e, ns2_1set, symshen_4syntax_1item_2, tmp18004)


_ = tmp18041

tmp18042 := MakeNative(func(__e *ControlFlow) {
V264 := __e.Get(1)
_ = V264
tmp18043 := MakeNative(func(__e *ControlFlow) {
W265 := __e.Get(1)
_ = W265
tmp18064 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W265)


if True == tmp18064 {
tmp18044 := MakeNative(func(__e *ControlFlow) {
W273 := __e.Get(1)
_ = W273
tmp18046 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W273)


if True == tmp18046 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W273)
return
}


}, 1)

tmp18047 := MakeNative(func(__e *ControlFlow) {
W274 := __e.Get(1)
_ = W274
tmp18060 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W274)


if True == tmp18060 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp18048 := MakeNative(func(__e *ControlFlow) {
W275 := __e.Get(1)
_ = W275
tmp18057 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W275)
}
__typedArg0 := W275
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp18057 {
tmp18049 := MakeNative(func(__e *ControlFlow) {
W276 := __e.Get(1)
_ = W276
tmp18050 := MakeNative(func(__e *ControlFlow) {
W277 := __e.Get(1)
_ = W277
tmp18052 := Call(__e, PrimFunc(symshen_4semicolon_2), W276)


if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp18052)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp18052
return Call(__e, PrimFunc(symnot), __typedArg0)
})() {
__e.TailApply(PrimFunc(symshen_4comb), W277, W276)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp18054 := Call(__e, PrimFunc(symtail), W275)


__e.TailApply(tmp18050, tmp18054)
return


}, 1)

tmp18055 := Call(__e, PrimFunc(symhead), W275)


__e.TailApply(tmp18049, tmp18055)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp18058 := Call(__e, PrimFunc(symshen_4in_1_6), W274)


__e.TailApply(tmp18048, tmp18058)
return


}


}, 1)

tmp18061 := Call(__e, PrimFunc(symshen_4_5colon_1equal_6), V264)


tmp18062 := Call(__e, tmp18047, tmp18061)


__e.TailApply(tmp18044, tmp18062)
return


} else {
__e.Return(W265)
return
}


}, 1)

tmp18065 := MakeNative(func(__e *ControlFlow) {
W266 := __e.Get(1)
_ = W266
tmp18091 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W266)


if True == tmp18091 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp18066 := MakeNative(func(__e *ControlFlow) {
W267 := __e.Get(1)
_ = W267
tmp18088 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W267)
}
__typedArg0 := W267
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp18088 {
tmp18067 := MakeNative(func(__e *ControlFlow) {
W268 := __e.Get(1)
_ = W268
tmp18068 := MakeNative(func(__e *ControlFlow) {
W269 := __e.Get(1)
_ = W269
tmp18084 := Call(__e, PrimFunc(symshen_4hds_a_2), W269, symwhere)


if True == tmp18084 {
tmp18069 := MakeNative(func(__e *ControlFlow) {
W270 := __e.Get(1)
_ = W270
tmp18081 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W270)
}
__typedArg0 := W270
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp18081 {
tmp18070 := MakeNative(func(__e *ControlFlow) {
W271 := __e.Get(1)
_ = W271
tmp18071 := MakeNative(func(__e *ControlFlow) {
W272 := __e.Get(1)
_ = W272
tmp18076 := Call(__e, PrimFunc(symshen_4semicolon_2), W268)


if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp18076)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp18076
return Call(__e, PrimFunc(symnot), __typedArg0)
})() {
tmp18072 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W268, Nil)
}
__typedArg0 := W268
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18073 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W271, tmp18072)
}
__typedArg0 := W271
__typedArg1 := tmp18072
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18074 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhere, tmp18073)
}
__typedArg0 := symwhere
__typedArg1 := tmp18073
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4comb), W272, tmp18074)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp18078 := Call(__e, PrimFunc(symtail), W270)


__e.TailApply(tmp18071, tmp18078)
return


}, 1)

tmp18079 := Call(__e, PrimFunc(symhead), W270)


__e.TailApply(tmp18070, tmp18079)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp18082 := Call(__e, PrimFunc(symtail), W269)


__e.TailApply(tmp18069, tmp18082)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp18085 := Call(__e, PrimFunc(symtail), W267)


__e.TailApply(tmp18068, tmp18085)
return


}, 1)

tmp18086 := Call(__e, PrimFunc(symhead), W267)


__e.TailApply(tmp18067, tmp18086)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp18089 := Call(__e, PrimFunc(symshen_4in_1_6), W266)


__e.TailApply(tmp18066, tmp18089)
return


}


}, 1)

tmp18092 := Call(__e, PrimFunc(symshen_4_5colon_1equal_6), V264)


tmp18093 := Call(__e, tmp18065, tmp18092)


__e.TailApply(tmp18043, tmp18093)
return


}, 1)

tmp18094 := Call(__e, ns2_1set, symshen_4_5semantics_6, tmp18042)


_ = tmp18094

tmp18095 := MakeNative(func(__e *ControlFlow) {
V286 := __e.Get(1)
_ = V286
V287 := __e.Get(2)
_ = V287
V288 := __e.Get(3)
_ = V288
tmp18103 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V288)
}
__typedArg0 := Nil
__typedArg1 := V288
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp18103 {
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
tmp18101 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V288)
}
__typedArg0 := V288
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp18101 {
tmp18096 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V288)
}
__typedArg0 := V288
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18097 := Call(__e, PrimFunc(symshen_4c_1rule_1_6shen), V286, tmp18096, V287)


tmp18098 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V288)
}
__typedArg0 := V288
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18099 := Call(__e, PrimFunc(symshen_4c_1rules_1_6shen), V286, V287, tmp18098)


__e.TailApply(PrimFunc(symshen_4combine_1c_1code), tmp18097, tmp18099)
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

tmp18104 := Call(__e, ns2_1set, symshen_4c_1rules_1_6shen, tmp18095)


_ = tmp18104

tmp18105 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symfail))
return
}, 0)

tmp18106 := Call(__e, ns2_1set, symshen_4parse_1failure, tmp18105)


_ = tmp18106

tmp18107 := MakeNative(func(__e *ControlFlow) {
V289 := __e.Get(1)
_ = V289
V290 := __e.Get(2)
_ = V290
tmp18108 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symResult, Nil)
}
__typedArg0 := symResult
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18109 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4parse_1failure_2, tmp18108)
}
__typedArg0 := symshen_4parse_1failure_2
__typedArg1 := tmp18108
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18110 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symResult, Nil)
}
__typedArg0 := symResult
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18111 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V290, tmp18110)
}
__typedArg0 := V290
__typedArg1 := tmp18110
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18112 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18109, tmp18111)
}
__typedArg0 := tmp18109
__typedArg1 := tmp18111
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18113 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp18112)
}
__typedArg0 := symif
__typedArg1 := tmp18112
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18114 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18113, Nil)
}
__typedArg0 := tmp18113
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18115 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V289, tmp18114)
}
__typedArg0 := V289
__typedArg1 := tmp18114
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18116 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symResult, tmp18115)
}
__typedArg0 := symResult
__typedArg1 := tmp18115
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp18116)
}
__typedArg0 := symlet
__typedArg1 := tmp18116
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 2)

tmp18117 := Call(__e, ns2_1set, symshen_4combine_1c_1code, tmp18107)


_ = tmp18117

tmp18118 := MakeNative(func(__e *ControlFlow) {
V297 := __e.Get(1)
_ = V297
V298 := __e.Get(2)
_ = V298
V299 := __e.Get(3)
_ = V299
tmp18132 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V298)
}
__typedArg0 := V298
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18123 Obj

if True == tmp18132 {
tmp18130 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V298)
}
__typedArg0 := V298
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18131 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18130)
}
__typedArg0 := tmp18130
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18125 Obj

if True == tmp18131 {
tmp18127 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V298)
}
__typedArg0 := V298
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18128 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18127)
}
__typedArg0 := tmp18127
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18129 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp18128)
}
__typedArg0 := Nil
__typedArg1 := tmp18128
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18126 Obj

if True == tmp18129 {
ifres18126 = True


} else {
ifres18126 = False


}

ifres18125 = ifres18126


} else {
ifres18125 = False


}

var ifres18124 Obj

if True == ifres18125 {
ifres18124 = True


} else {
ifres18124 = False


}

ifres18123 = ifres18124


} else {
ifres18123 = False


}

if True == ifres18123 {
tmp18119 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V298)
}
__typedArg0 := V298
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18120 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V298)
}
__typedArg0 := V298
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18121 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18120)
}
__typedArg0 := tmp18120
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4yacc_1syntax), V297, V299, tmp18119, tmp18121)
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

tmp18133 := Call(__e, ns2_1set, symshen_4c_1rule_1_6shen, tmp18118)


_ = tmp18133

tmp18134 := MakeNative(func(__e *ControlFlow) {
V308 := __e.Get(1)
_ = V308
V309 := __e.Get(2)
_ = V309
V310 := __e.Get(3)
_ = V310
V311 := __e.Get(4)
_ = V311
tmp18198 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V310)
}
__typedArg0 := Nil
__typedArg1 := V310
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18176 Obj

if True == tmp18198 {
tmp18197 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V311)
}
__typedArg0 := V311
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18178 Obj

if True == tmp18197 {
tmp18195 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V311)
}
__typedArg0 := V311
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18196 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symwhere, tmp18195)
}
__typedArg0 := symwhere
__typedArg1 := tmp18195
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18180 Obj

if True == tmp18196 {
tmp18193 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V311)
}
__typedArg0 := V311
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18194 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18193)
}
__typedArg0 := tmp18193
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18182 Obj

if True == tmp18194 {
tmp18190 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V311)
}
__typedArg0 := V311
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18191 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18190)
}
__typedArg0 := tmp18190
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18192 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18191)
}
__typedArg0 := tmp18191
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18184 Obj

if True == tmp18192 {
tmp18186 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V311)
}
__typedArg0 := V311
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18187 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18186)
}
__typedArg0 := tmp18186
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18188 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18187)
}
__typedArg0 := tmp18187
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18189 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp18188)
}
__typedArg0 := Nil
__typedArg1 := tmp18188
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18185 Obj

if True == tmp18189 {
ifres18185 = True


} else {
ifres18185 = False


}

ifres18184 = ifres18185


} else {
ifres18184 = False


}

var ifres18183 Obj

if True == ifres18184 {
ifres18183 = True


} else {
ifres18183 = False


}

ifres18182 = ifres18183


} else {
ifres18182 = False


}

var ifres18181 Obj

if True == ifres18182 {
ifres18181 = True


} else {
ifres18181 = False


}

ifres18180 = ifres18181


} else {
ifres18180 = False


}

var ifres18179 Obj

if True == ifres18180 {
ifres18179 = True


} else {
ifres18179 = False


}

ifres18178 = ifres18179


} else {
ifres18178 = False


}

var ifres18177 Obj

if True == ifres18178 {
ifres18177 = True


} else {
ifres18177 = False


}

ifres18176 = ifres18177


} else {
ifres18176 = False


}

if True == ifres18176 {
tmp18135 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V311)
}
__typedArg0 := V311
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18136 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18135)
}
__typedArg0 := tmp18135
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18137 := Call(__e, PrimFunc(symshen_4process_1yacc_1semantics), tmp18136)


tmp18138 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V311)
}
__typedArg0 := V311
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18139 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18138)
}
__typedArg0 := tmp18138
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18140 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18139)
}
__typedArg0 := tmp18139
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18141 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V308, V309, Nil, tmp18140)


tmp18142 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4parse_1failure, Nil)
}
__typedArg0 := symshen_4parse_1failure
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18143 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18142, Nil)
}
__typedArg0 := tmp18142
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18144 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18141, tmp18143)
}
__typedArg0 := tmp18141
__typedArg1 := tmp18143
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18145 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18137, tmp18144)
}
__typedArg0 := tmp18137
__typedArg1 := tmp18144
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp18145)
}
__typedArg0 := symif
__typedArg1 := tmp18145
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp18174 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V310)
}
__typedArg0 := Nil
__typedArg1 := V310
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp18174 {
__e.TailApply(PrimFunc(symshen_4yacc_1semantics), V308, V309, V311)
return
} else {
tmp18172 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp18172 {
tmp18169 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18170 := Call(__e, PrimFunc(symshen_4non_1terminal_2), tmp18169)


if True == tmp18170 {
tmp18146 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18147 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4non_1terminalcode), V308, V309, tmp18146, tmp18147, V311)
return


} else {
tmp18166 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18167 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(tmp18166)
}
__typedArg0 := tmp18166
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

if True == tmp18167 {
tmp18148 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18149 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4variablecode), V308, V309, tmp18148, tmp18149, V311)
return


} else {
tmp18163 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18164 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym__, tmp18163)
}
__typedArg0 := sym__
__typedArg1 := tmp18163
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp18164 {
tmp18150 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18151 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4wildcardcode), V308, V309, tmp18150, tmp18151, V311)
return


} else {
tmp18160 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18161 := Call(__e, PrimFunc(symatom_2), tmp18160)


if True == tmp18161 {
tmp18152 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18153 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4terminalcode), V308, V309, tmp18152, tmp18153, V311)
return


} else {
tmp18157 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18158 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18157)
}
__typedArg0 := tmp18157
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp18158 {
tmp18154 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18155 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4conscode), V308, V309, tmp18154, tmp18155, V311)
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

tmp18199 := Call(__e, ns2_1set, symshen_4yacc_1syntax, tmp18134)


_ = tmp18199

tmp18200 := MakeNative(func(__e *ControlFlow) {
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
tmp18201 := MakeNative(func(__e *ControlFlow) {
W317 := __e.Get(1)
_ = W317
tmp18202 := MakeNative(func(__e *ControlFlow) {
W318 := __e.Get(1)
_ = W318
tmp18203 := MakeNative(func(__e *ControlFlow) {
W319 := __e.Get(1)
_ = W319
tmp18204 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V313, Nil)
}
__typedArg0 := V313
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18205 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V314, tmp18204)
}
__typedArg0 := V314
__typedArg1 := tmp18204
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18206 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W317, Nil)
}
__typedArg0 := W317
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18207 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4parse_1failure_2, tmp18206)
}
__typedArg0 := symshen_4parse_1failure_2
__typedArg1 := tmp18206
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18208 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4parse_1failure, Nil)
}
__typedArg0 := symshen_4parse_1failure
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18209 := MakeNative(func(__e *ControlFlow) {
W320 := __e.Get(1)
_ = W320
tmp18219 := Call(__e, PrimFunc(symshen_4occurs_1check_2), V314, V316)


var ifres18216 Obj

if True == tmp18219 {
ifres18216 = True


} else {
tmp18218 := Call(__e, PrimFunc(symshen_4occurs_1check_2), W318, V316)


var ifres18217 Obj

if True == tmp18218 {
ifres18217 = True


} else {
ifres18217 = False


}

ifres18216 = ifres18217


}

if True == ifres18216 {
tmp18210 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W317, Nil)
}
__typedArg0 := W317
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18211 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4_5_1out, tmp18210)
}
__typedArg0 := symshen_4_5_1out
__typedArg1 := tmp18210
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18212 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W320, Nil)
}
__typedArg0 := W320
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18213 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18211, tmp18212)
}
__typedArg0 := tmp18211
__typedArg1 := tmp18212
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18214 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W318, tmp18213)
}
__typedArg0 := W318
__typedArg1 := tmp18213
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp18214)
}
__typedArg0 := symlet
__typedArg1 := tmp18214
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return(W320)
return
}


}, 1)

tmp18220 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W317, Nil)
}
__typedArg0 := W317
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18221 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4in_1_6, tmp18220)
}
__typedArg0 := symshen_4in_1_6
__typedArg1 := tmp18220
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18222 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V312, W319, V315, V316)


tmp18223 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18222, Nil)
}
__typedArg0 := tmp18222
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18224 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18221, tmp18223)
}
__typedArg0 := tmp18221
__typedArg1 := tmp18223
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18225 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W319, tmp18224)
}
__typedArg0 := W319
__typedArg1 := tmp18224
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18226 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp18225)
}
__typedArg0 := symlet
__typedArg1 := tmp18225
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18227 := Call(__e, tmp18209, tmp18226)


tmp18228 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18227, Nil)
}
__typedArg0 := tmp18227
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18229 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18208, tmp18228)
}
__typedArg0 := tmp18208
__typedArg1 := tmp18228
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18230 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18207, tmp18229)
}
__typedArg0 := tmp18207
__typedArg1 := tmp18229
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18231 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp18230)
}
__typedArg0 := symif
__typedArg1 := tmp18230
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18232 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18231, Nil)
}
__typedArg0 := tmp18231
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18233 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18205, tmp18232)
}
__typedArg0 := tmp18205
__typedArg1 := tmp18232
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18234 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W317, tmp18233)
}
__typedArg0 := W317
__typedArg1 := tmp18233
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp18234)
}
__typedArg0 := symlet
__typedArg1 := tmp18234
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp18235 := Call(__e, PrimFunc(symconcat), symRemainder, V314)


__e.TailApply(tmp18203, tmp18235)
return


}, 1)

tmp18236 := Call(__e, PrimFunc(symconcat), symAction, V314)


__e.TailApply(tmp18202, tmp18236)
return


}, 1)

tmp18237 := Call(__e, PrimFunc(symconcat), symParse, V314)


__e.TailApply(tmp18201, tmp18237)
return


}, 5)

tmp18238 := Call(__e, ns2_1set, symshen_4non_1terminalcode, tmp18200)


_ = tmp18238

tmp18239 := MakeNative(func(__e *ControlFlow) {
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
tmp18240 := MakeNative(func(__e *ControlFlow) {
W326 := __e.Get(1)
_ = W326
tmp18241 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V322, Nil)
}
__typedArg0 := V322
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18242 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons_2, tmp18241)
}
__typedArg0 := symcons_2
__typedArg1 := tmp18241
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18243 := MakeNative(func(__e *ControlFlow) {
W327 := __e.Get(1)
_ = W327
tmp18250 := Call(__e, PrimFunc(symshen_4occurs_1check_2), V323, V325)


if True == tmp18250 {
tmp18244 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V322, Nil)
}
__typedArg0 := V322
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18245 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhead, tmp18244)
}
__typedArg0 := symhead
__typedArg1 := tmp18244
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18246 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W327, Nil)
}
__typedArg0 := W327
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18247 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18245, tmp18246)
}
__typedArg0 := tmp18245
__typedArg1 := tmp18246
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18248 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V323, tmp18247)
}
__typedArg0 := V323
__typedArg1 := tmp18247
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp18248)
}
__typedArg0 := symlet
__typedArg1 := tmp18248
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return(W327)
return
}


}, 1)

tmp18251 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V322, Nil)
}
__typedArg0 := V322
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18252 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtail, tmp18251)
}
__typedArg0 := symtail
__typedArg1 := tmp18251
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18253 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V321, W326, V324, V325)


tmp18254 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18253, Nil)
}
__typedArg0 := tmp18253
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18255 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18252, tmp18254)
}
__typedArg0 := tmp18252
__typedArg1 := tmp18254
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18256 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W326, tmp18255)
}
__typedArg0 := W326
__typedArg1 := tmp18255
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18257 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp18256)
}
__typedArg0 := symlet
__typedArg1 := tmp18256
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18258 := Call(__e, tmp18243, tmp18257)


tmp18259 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4parse_1failure, Nil)
}
__typedArg0 := symshen_4parse_1failure
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

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
return PrimCons(tmp18258, tmp18260)
}
__typedArg0 := tmp18258
__typedArg1 := tmp18260
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18262 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18242, tmp18261)
}
__typedArg0 := tmp18242
__typedArg1 := tmp18261
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp18262)
}
__typedArg0 := symif
__typedArg1 := tmp18262
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp18263 := Call(__e, PrimFunc(symgensym), symRemainder)


__e.TailApply(tmp18240, tmp18263)
return


}, 5)

tmp18264 := Call(__e, ns2_1set, symshen_4variablecode, tmp18239)


_ = tmp18264

tmp18265 := MakeNative(func(__e *ControlFlow) {
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
tmp18266 := MakeNative(func(__e *ControlFlow) {
W333 := __e.Get(1)
_ = W333
tmp18267 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V329, Nil)
}
__typedArg0 := V329
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18268 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons_2, tmp18267)
}
__typedArg0 := symcons_2
__typedArg1 := tmp18267
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18269 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V329, Nil)
}
__typedArg0 := V329
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18270 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtail, tmp18269)
}
__typedArg0 := symtail
__typedArg1 := tmp18269
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18271 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V328, W333, V331, V332)


tmp18272 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18271, Nil)
}
__typedArg0 := tmp18271
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18273 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18270, tmp18272)
}
__typedArg0 := tmp18270
__typedArg1 := tmp18272
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18274 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W333, tmp18273)
}
__typedArg0 := W333
__typedArg1 := tmp18273
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18275 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp18274)
}
__typedArg0 := symlet
__typedArg1 := tmp18274
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18276 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4parse_1failure, Nil)
}
__typedArg0 := symshen_4parse_1failure
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18277 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18276, Nil)
}
__typedArg0 := tmp18276
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18278 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18275, tmp18277)
}
__typedArg0 := tmp18275
__typedArg1 := tmp18277
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18279 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18268, tmp18278)
}
__typedArg0 := tmp18268
__typedArg1 := tmp18278
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp18279)
}
__typedArg0 := symif
__typedArg1 := tmp18279
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp18280 := Call(__e, PrimFunc(symgensym), symRemainder)


__e.TailApply(tmp18266, tmp18280)
return


}, 5)

tmp18281 := Call(__e, ns2_1set, symshen_4wildcardcode, tmp18265)


_ = tmp18281

tmp18282 := MakeNative(func(__e *ControlFlow) {
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
tmp18283 := MakeNative(func(__e *ControlFlow) {
W339 := __e.Get(1)
_ = W339
tmp18284 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V336, Nil)
}
__typedArg0 := V336
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18285 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V335, tmp18284)
}
__typedArg0 := V335
__typedArg1 := tmp18284
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18286 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4hds_a_2, tmp18285)
}
__typedArg0 := symshen_4hds_a_2
__typedArg1 := tmp18285
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18287 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V335, Nil)
}
__typedArg0 := V335
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18288 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtail, tmp18287)
}
__typedArg0 := symtail
__typedArg1 := tmp18287
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18289 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V334, W339, V337, V338)


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
return PrimCons(tmp18288, tmp18290)
}
__typedArg0 := tmp18288
__typedArg1 := tmp18290
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18292 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W339, tmp18291)
}
__typedArg0 := W339
__typedArg1 := tmp18291
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18293 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp18292)
}
__typedArg0 := symlet
__typedArg1 := tmp18292
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18294 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4parse_1failure, Nil)
}
__typedArg0 := symshen_4parse_1failure
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18295 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18294, Nil)
}
__typedArg0 := tmp18294
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18296 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18293, tmp18295)
}
__typedArg0 := tmp18293
__typedArg1 := tmp18295
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18297 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18286, tmp18296)
}
__typedArg0 := tmp18286
__typedArg1 := tmp18296
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp18297)
}
__typedArg0 := symif
__typedArg1 := tmp18297
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp18298 := Call(__e, PrimFunc(symgensym), symRemainder)


__e.TailApply(tmp18283, tmp18298)
return


}, 5)

tmp18299 := Call(__e, ns2_1set, symshen_4terminalcode, tmp18282)


_ = tmp18299

tmp18300 := MakeNative(func(__e *ControlFlow) {
V347 := __e.Get(1)
_ = V347
V348 := __e.Get(2)
_ = V348
tmp18306 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V347)
}
__typedArg0 := V347
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18302 Obj

if True == tmp18306 {
tmp18304 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V347)
}
__typedArg0 := V347
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18305 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp18304, V348)
}
__typedArg0 := tmp18304
__typedArg1 := V348
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18303 Obj

if True == tmp18305 {
ifres18303 = True


} else {
ifres18303 = False


}

ifres18302 = ifres18303


} else {
ifres18302 = False


}

if True == ifres18302 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}, 2)

tmp18307 := Call(__e, ns2_1set, symshen_4hds_a_2, tmp18300)


_ = tmp18307

tmp18308 := MakeNative(func(__e *ControlFlow) {
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
tmp18309 := MakeNative(func(__e *ControlFlow) {
W354 := __e.Get(1)
_ = W354
tmp18310 := MakeNative(func(__e *ControlFlow) {
W355 := __e.Get(1)
_ = W355
tmp18311 := MakeNative(func(__e *ControlFlow) {
W356 := __e.Get(1)
_ = W356
tmp18312 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V350, Nil)
}
__typedArg0 := V350
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18313 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4ccons_2, tmp18312)
}
__typedArg0 := symshen_4ccons_2
__typedArg1 := tmp18312
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18314 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V350, Nil)
}
__typedArg0 := V350
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18315 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhead, tmp18314)
}
__typedArg0 := symhead
__typedArg1 := tmp18314
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18316 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V350, Nil)
}
__typedArg0 := V350
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18317 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtail, tmp18316)
}
__typedArg0 := symtail
__typedArg1 := tmp18316
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18318 := Call(__e, PrimFunc(symshen_4decons), V351)


tmp18319 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5end_6, Nil)
}
__typedArg0 := sym_5end_6
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18320 := Call(__e, PrimFunc(symappend), tmp18318, tmp18319)


tmp18321 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V349, W356, V352, V353)


tmp18322 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18321, Nil)
}
__typedArg0 := tmp18321
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18323 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4processed, tmp18322)
}
__typedArg0 := symshen_4processed
__typedArg1 := tmp18322
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18324 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V349, W355, tmp18320, tmp18323)


tmp18325 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18324, Nil)
}
__typedArg0 := tmp18324
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18326 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18317, tmp18325)
}
__typedArg0 := tmp18317
__typedArg1 := tmp18325
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18327 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W356, tmp18326)
}
__typedArg0 := W356
__typedArg1 := tmp18326
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18328 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18315, tmp18327)
}
__typedArg0 := tmp18315
__typedArg1 := tmp18327
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18329 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W355, tmp18328)
}
__typedArg0 := W355
__typedArg1 := tmp18328
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18330 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp18329)
}
__typedArg0 := symlet
__typedArg1 := tmp18329
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18331 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4parse_1failure, Nil)
}
__typedArg0 := symshen_4parse_1failure
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18332 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18331, Nil)
}
__typedArg0 := tmp18331
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18333 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18330, tmp18332)
}
__typedArg0 := tmp18330
__typedArg1 := tmp18332
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18334 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18313, tmp18333)
}
__typedArg0 := tmp18313
__typedArg1 := tmp18333
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp18334)
}
__typedArg0 := symif
__typedArg1 := tmp18334
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp18335 := Call(__e, PrimFunc(symgensym), symTl)


__e.TailApply(tmp18311, tmp18335)
return


}, 1)

tmp18336 := Call(__e, PrimFunc(symgensym), symHd)


__e.TailApply(tmp18310, tmp18336)
return


}, 1)

tmp18337 := Call(__e, PrimFunc(symgensym), symRemainder)


__e.TailApply(tmp18309, tmp18337)
return


}, 5)

tmp18338 := Call(__e, ns2_1set, symshen_4conscode, tmp18308)


_ = tmp18338

tmp18339 := MakeNative(func(__e *ControlFlow) {
V367 := __e.Get(1)
_ = V367
tmp18351 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V367)
}
__typedArg0 := V367
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18347 Obj

if True == tmp18351 {
tmp18349 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V367)
}
__typedArg0 := V367
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18350 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18349)
}
__typedArg0 := tmp18349
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18348 Obj

if True == tmp18350 {
ifres18348 = True


} else {
ifres18348 = False


}

ifres18347 = ifres18348


} else {
ifres18347 = False


}

if True == ifres18347 {
__e.Return(True)
return
} else {
tmp18345 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V367)
}
__typedArg0 := V367
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18341 Obj

if True == tmp18345 {
tmp18343 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V367)
}
__typedArg0 := V367
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18344 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp18343)
}
__typedArg0 := Nil
__typedArg1 := tmp18343
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18342 Obj

if True == tmp18344 {
ifres18342 = True


} else {
ifres18342 = False


}

ifres18341 = ifres18342


} else {
ifres18341 = False


}

if True == ifres18341 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp18352 := Call(__e, ns2_1set, symshen_4ccons_2, tmp18339)


_ = tmp18352

tmp18353 := MakeNative(func(__e *ControlFlow) {
V368 := __e.Get(1)
_ = V368
tmp18380 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V368)
}
__typedArg0 := V368
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18361 Obj

if True == tmp18380 {
tmp18378 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V368)
}
__typedArg0 := V368
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18379 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcons, tmp18378)
}
__typedArg0 := symcons
__typedArg1 := tmp18378
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18363 Obj

if True == tmp18379 {
tmp18376 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V368)
}
__typedArg0 := V368
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18377 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18376)
}
__typedArg0 := tmp18376
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18365 Obj

if True == tmp18377 {
tmp18373 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V368)
}
__typedArg0 := V368
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18374)
}
__typedArg0 := tmp18374
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18367 Obj

if True == tmp18375 {
tmp18369 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V368)
}
__typedArg0 := V368
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18370 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18369)
}
__typedArg0 := tmp18369
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18371 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18370)
}
__typedArg0 := tmp18370
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18372 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp18371)
}
__typedArg0 := Nil
__typedArg1 := tmp18371
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18368 Obj

if True == tmp18372 {
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
tmp18354 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V368)
}
__typedArg0 := V368
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18355 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18354)
}
__typedArg0 := tmp18354
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18356 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V368)
}
__typedArg0 := V368
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18357 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18356)
}
__typedArg0 := tmp18356
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18358 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18357)
}
__typedArg0 := tmp18357
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18359 := Call(__e, PrimFunc(symshen_4decons), tmp18358)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18355, tmp18359)
}
__typedArg0 := tmp18355
__typedArg1 := tmp18359
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return(V368)
return
}


}, 1)

tmp18381 := Call(__e, ns2_1set, symshen_4decons, tmp18353)


_ = tmp18381

tmp18382 := MakeNative(func(__e *ControlFlow) {
V369 := __e.Get(1)
_ = V369
V370 := __e.Get(2)
_ = V370
tmp18383 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V370, Nil)
}
__typedArg0 := V370
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V369, tmp18383)
}
__typedArg0 := V369
__typedArg1 := tmp18383
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 2)

tmp18384 := Call(__e, ns2_1set, symshen_4comb, tmp18382)


_ = tmp18384

tmp18385 := MakeNative(func(__e *ControlFlow) {
V375 := __e.Get(1)
_ = V375
V376 := __e.Get(2)
_ = V376
V377 := __e.Get(3)
_ = V377
tmp18407 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V377)
}
__typedArg0 := V377
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18394 Obj

if True == tmp18407 {
tmp18405 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V377)
}
__typedArg0 := V377
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18406 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4processed, tmp18405)
}
__typedArg0 := symshen_4processed
__typedArg1 := tmp18405
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18396 Obj

if True == tmp18406 {
tmp18403 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V377)
}
__typedArg0 := V377
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18404 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18403)
}
__typedArg0 := tmp18403
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18398 Obj

if True == tmp18404 {
tmp18400 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V377)
}
__typedArg0 := V377
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18401 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18400)
}
__typedArg0 := tmp18400
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18402 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp18401)
}
__typedArg0 := Nil
__typedArg1 := tmp18401
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18399 Obj

if True == tmp18402 {
ifres18399 = True


} else {
ifres18399 = False


}

ifres18398 = ifres18399


} else {
ifres18398 = False


}

var ifres18397 Obj

if True == ifres18398 {
ifres18397 = True


} else {
ifres18397 = False


}

ifres18396 = ifres18397


} else {
ifres18396 = False


}

var ifres18395 Obj

if True == ifres18396 {
ifres18395 = True


} else {
ifres18395 = False


}

ifres18394 = ifres18395


} else {
ifres18394 = False


}

if True == ifres18394 {
tmp18386 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V377)
}
__typedArg0 := V377
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18386)
}
__typedArg0 := tmp18386
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return


} else {
tmp18387 := MakeNative(func(__e *ControlFlow) {
W378 := __e.Get(1)
_ = W378
tmp18388 := MakeNative(func(__e *ControlFlow) {
W379 := __e.Get(1)
_ = W379
tmp18389 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W379, Nil)
}
__typedArg0 := W379
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18390 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V376, tmp18389)
}
__typedArg0 := V376
__typedArg1 := tmp18389
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4comb, tmp18390)
}
__typedArg0 := symshen_4comb
__typedArg1 := tmp18390
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp18391 := Call(__e, PrimFunc(symshen_4use_1type_1info), V375, W378)


__e.TailApply(tmp18388, tmp18391)
return


}, 1)

tmp18392 := Call(__e, PrimFunc(symshen_4process_1yacc_1semantics), V377)


__e.TailApply(tmp18387, tmp18392)
return


}


}, 3)

tmp18408 := Call(__e, ns2_1set, symshen_4yacc_1semantics, tmp18385)


_ = tmp18408

tmp18409 := MakeNative(func(__e *ControlFlow) {
V383 := __e.Get(1)
_ = V383
V384 := __e.Get(2)
_ = V384
tmp18597 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18418 Obj

if True == tmp18597 {
tmp18595 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18596 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_i, tmp18595)
}
__typedArg0 := sym_i
__typedArg1 := tmp18595
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18420 Obj

if True == tmp18596 {
tmp18593 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18594 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18593)
}
__typedArg0 := tmp18593
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18422 Obj

if True == tmp18594 {
tmp18590 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18591 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18590)
}
__typedArg0 := tmp18590
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18592 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18591)
}
__typedArg0 := tmp18591
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18424 Obj

if True == tmp18592 {
tmp18586 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18587 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18586)
}
__typedArg0 := tmp18586
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18588 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18587)
}
__typedArg0 := tmp18587
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18589 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlist, tmp18588)
}
__typedArg0 := symlist
__typedArg1 := tmp18588
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18426 Obj

if True == tmp18589 {
tmp18582 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18583 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18582)
}
__typedArg0 := tmp18582
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18584 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18583)
}
__typedArg0 := tmp18583
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18585 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18584)
}
__typedArg0 := tmp18584
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18428 Obj

if True == tmp18585 {
tmp18577 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18578 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18577)
}
__typedArg0 := tmp18577
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18579 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18578)
}
__typedArg0 := tmp18578
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18580 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18579)
}
__typedArg0 := tmp18579
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18581 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp18580)
}
__typedArg0 := Nil
__typedArg1 := tmp18580
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18430 Obj

if True == tmp18581 {
tmp18574 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18575 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18574)
}
__typedArg0 := tmp18574
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18576 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18575)
}
__typedArg0 := tmp18575
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18432 Obj

if True == tmp18576 {
tmp18570 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18571 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18570)
}
__typedArg0 := tmp18570
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18572 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18571)
}
__typedArg0 := tmp18571
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18573 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_1_1_6, tmp18572)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp18572
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18434 Obj

if True == tmp18573 {
tmp18566 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18567 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18566)
}
__typedArg0 := tmp18566
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18568 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18567)
}
__typedArg0 := tmp18567
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18569 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18568)
}
__typedArg0 := tmp18568
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18436 Obj

if True == tmp18569 {
tmp18561 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18562 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18561)
}
__typedArg0 := tmp18561
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18563 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18562)
}
__typedArg0 := tmp18562
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18564 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18563)
}
__typedArg0 := tmp18563
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18565 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18564)
}
__typedArg0 := tmp18564
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18438 Obj

if True == tmp18565 {
tmp18555 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18556 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18555)
}
__typedArg0 := tmp18555
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18557)
}
__typedArg0 := tmp18557
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18559 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18558)
}
__typedArg0 := tmp18558
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18560 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symstr, tmp18559)
}
__typedArg0 := symstr
__typedArg1 := tmp18559
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18440 Obj

if True == tmp18560 {
tmp18549 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18550 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18549)
}
__typedArg0 := tmp18549
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18551 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18550)
}
__typedArg0 := tmp18550
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18552 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18551)
}
__typedArg0 := tmp18551
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18553 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18552)
}
__typedArg0 := tmp18552
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18554 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18553)
}
__typedArg0 := tmp18553
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18442 Obj

if True == tmp18554 {
tmp18542 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18543)
}
__typedArg0 := tmp18543
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18545 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18544)
}
__typedArg0 := tmp18544
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18546 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18545)
}
__typedArg0 := tmp18545
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18547 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18546)
}
__typedArg0 := tmp18546
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18548 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18547)
}
__typedArg0 := tmp18547
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18444 Obj

if True == tmp18548 {
tmp18534 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18535 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18534)
}
__typedArg0 := tmp18534
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18536 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18535)
}
__typedArg0 := tmp18535
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18537 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18536)
}
__typedArg0 := tmp18536
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18538 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18537)
}
__typedArg0 := tmp18537
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18539 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18538)
}
__typedArg0 := tmp18538
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18540 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18539)
}
__typedArg0 := tmp18539
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18541 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlist, tmp18540)
}
__typedArg0 := symlist
__typedArg1 := tmp18540
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18446 Obj

if True == tmp18541 {
tmp18526 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18527 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18526)
}
__typedArg0 := tmp18526
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18528 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18527)
}
__typedArg0 := tmp18527
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18529 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18528)
}
__typedArg0 := tmp18528
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18530 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18529)
}
__typedArg0 := tmp18529
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18531 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18530)
}
__typedArg0 := tmp18530
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18532 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18531)
}
__typedArg0 := tmp18531
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18533 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18532)
}
__typedArg0 := tmp18532
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18448 Obj

if True == tmp18533 {
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18518)
}
__typedArg0 := tmp18518
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18520)
}
__typedArg0 := tmp18520
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18522 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18521)
}
__typedArg0 := tmp18521
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18523 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18522)
}
__typedArg0 := tmp18522
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18524 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18523)
}
__typedArg0 := tmp18523
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18525 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp18524)
}
__typedArg0 := Nil
__typedArg1 := tmp18524
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18450 Obj

if True == tmp18525 {
tmp18510 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18511)
}
__typedArg0 := tmp18511
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18513)
}
__typedArg0 := tmp18513
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18515 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18514)
}
__typedArg0 := tmp18514
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18516 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18515)
}
__typedArg0 := tmp18515
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18452 Obj

if True == tmp18516 {
tmp18502 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18503 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18502)
}
__typedArg0 := tmp18502
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18507)
}
__typedArg0 := tmp18507
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18509 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp18508)
}
__typedArg0 := Nil
__typedArg1 := tmp18508
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18454 Obj

if True == tmp18509 {
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18499)
}
__typedArg0 := tmp18499
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18501 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18500)
}
__typedArg0 := tmp18500
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18456 Obj

if True == tmp18501 {
tmp18491 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18492 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18491)
}
__typedArg0 := tmp18491
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_j, tmp18495)
}
__typedArg0 := sym_j
__typedArg1 := tmp18495
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18458 Obj

if True == tmp18496 {
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp18489)
}
__typedArg0 := Nil
__typedArg1 := tmp18489
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18460 Obj

if True == tmp18490 {
tmp18472 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18473)
}
__typedArg0 := tmp18473
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18475 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18474)
}
__typedArg0 := tmp18474
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18476 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18477)
}
__typedArg0 := tmp18477
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18479 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18478)
}
__typedArg0 := tmp18478
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18480 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18479)
}
__typedArg0 := tmp18479
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18481 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18480)
}
__typedArg0 := tmp18480
return Call(__e, PrimFunc(symhd), __typedArg0)
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp18475, tmp18483)
}
__typedArg0 := tmp18475
__typedArg1 := tmp18483
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18462 Obj

if True == tmp18484 {
tmp18464 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18465 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18464)
}
__typedArg0 := tmp18464
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18466)
}
__typedArg0 := tmp18466
return Call(__e, PrimFunc(symhd), __typedArg0)
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

tmp18471 := Call(__e, PrimFunc(symshen_4monomorphic_2), tmp18470)


var ifres18463 Obj

if True == tmp18471 {
ifres18463 = True


} else {
ifres18463 = False


}

ifres18462 = ifres18463


} else {
ifres18462 = False


}

var ifres18461 Obj

if True == ifres18462 {
ifres18461 = True


} else {
ifres18461 = False


}

ifres18460 = ifres18461


} else {
ifres18460 = False


}

var ifres18459 Obj

if True == ifres18460 {
ifres18459 = True


} else {
ifres18459 = False


}

ifres18458 = ifres18459


} else {
ifres18458 = False


}

var ifres18457 Obj

if True == ifres18458 {
ifres18457 = True


} else {
ifres18457 = False


}

ifres18456 = ifres18457


} else {
ifres18456 = False


}

var ifres18455 Obj

if True == ifres18456 {
ifres18455 = True


} else {
ifres18455 = False


}

ifres18454 = ifres18455


} else {
ifres18454 = False


}

var ifres18453 Obj

if True == ifres18454 {
ifres18453 = True


} else {
ifres18453 = False


}

ifres18452 = ifres18453


} else {
ifres18452 = False


}

var ifres18451 Obj

if True == ifres18452 {
ifres18451 = True


} else {
ifres18451 = False


}

ifres18450 = ifres18451


} else {
ifres18450 = False


}

var ifres18449 Obj

if True == ifres18450 {
ifres18449 = True


} else {
ifres18449 = False


}

ifres18448 = ifres18449


} else {
ifres18448 = False


}

var ifres18447 Obj

if True == ifres18448 {
ifres18447 = True


} else {
ifres18447 = False


}

ifres18446 = ifres18447


} else {
ifres18446 = False


}

var ifres18445 Obj

if True == ifres18446 {
ifres18445 = True


} else {
ifres18445 = False


}

ifres18444 = ifres18445


} else {
ifres18444 = False


}

var ifres18443 Obj

if True == ifres18444 {
ifres18443 = True


} else {
ifres18443 = False


}

ifres18442 = ifres18443


} else {
ifres18442 = False


}

var ifres18441 Obj

if True == ifres18442 {
ifres18441 = True


} else {
ifres18441 = False


}

ifres18440 = ifres18441


} else {
ifres18440 = False


}

var ifres18439 Obj

if True == ifres18440 {
ifres18439 = True


} else {
ifres18439 = False


}

ifres18438 = ifres18439


} else {
ifres18438 = False


}

var ifres18437 Obj

if True == ifres18438 {
ifres18437 = True


} else {
ifres18437 = False


}

ifres18436 = ifres18437


} else {
ifres18436 = False


}

var ifres18435 Obj

if True == ifres18436 {
ifres18435 = True


} else {
ifres18435 = False


}

ifres18434 = ifres18435


} else {
ifres18434 = False


}

var ifres18433 Obj

if True == ifres18434 {
ifres18433 = True


} else {
ifres18433 = False


}

ifres18432 = ifres18433


} else {
ifres18432 = False


}

var ifres18431 Obj

if True == ifres18432 {
ifres18431 = True


} else {
ifres18431 = False


}

ifres18430 = ifres18431


} else {
ifres18430 = False


}

var ifres18429 Obj

if True == ifres18430 {
ifres18429 = True


} else {
ifres18429 = False


}

ifres18428 = ifres18429


} else {
ifres18428 = False


}

var ifres18427 Obj

if True == ifres18428 {
ifres18427 = True


} else {
ifres18427 = False


}

ifres18426 = ifres18427


} else {
ifres18426 = False


}

var ifres18425 Obj

if True == ifres18426 {
ifres18425 = True


} else {
ifres18425 = False


}

ifres18424 = ifres18425


} else {
ifres18424 = False


}

var ifres18423 Obj

if True == ifres18424 {
ifres18423 = True


} else {
ifres18423 = False


}

ifres18422 = ifres18423


} else {
ifres18422 = False


}

var ifres18421 Obj

if True == ifres18422 {
ifres18421 = True


} else {
ifres18421 = False


}

ifres18420 = ifres18421


} else {
ifres18420 = False


}

var ifres18419 Obj

if True == ifres18420 {
ifres18419 = True


} else {
ifres18419 = False


}

ifres18418 = ifres18419


} else {
ifres18418 = False


}

if True == ifres18418 {
tmp18410 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V384, tmp18415)
}
__typedArg0 := V384
__typedArg1 := tmp18415
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtype, tmp18416)
}
__typedArg0 := symtype
__typedArg1 := tmp18416
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return(V384)
return
}


}, 2)

tmp18598 := Call(__e, ns2_1set, symshen_4use_1type_1info, tmp18409)


_ = tmp18598

tmp18599 := MakeNative(func(__e *ControlFlow) {
V387 := __e.Get(1)
_ = V387
tmp18609 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(V387)
}
__typedArg0 := V387
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

if True == tmp18609 {
__e.Return(False)
return
} else {
tmp18607 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V387)
}
__typedArg0 := V387
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp18607 {
tmp18604 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V387)
}
__typedArg0 := V387
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18605 := Call(__e, PrimFunc(symshen_4monomorphic_2), tmp18604)


if True == tmp18605 {
tmp18601 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V387)
}
__typedArg0 := V387
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18602 := Call(__e, PrimFunc(symshen_4monomorphic_2), tmp18601)


if True == tmp18602 {
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

tmp18610 := Call(__e, ns2_1set, symshen_4monomorphic_2, tmp18599)


_ = tmp18610

tmp18611 := MakeNative(func(__e *ControlFlow) {
V388 := __e.Get(1)
_ = V388
tmp18637 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V388)
}
__typedArg0 := V388
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18619 Obj

if True == tmp18637 {
tmp18635 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V388)
}
__typedArg0 := V388
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18636 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symprotect, tmp18635)
}
__typedArg0 := symprotect
__typedArg1 := tmp18635
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18621 Obj

if True == tmp18636 {
tmp18633 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V388)
}
__typedArg0 := V388
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18634 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18633)
}
__typedArg0 := tmp18633
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18623 Obj

if True == tmp18634 {
tmp18630 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V388)
}
__typedArg0 := V388
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18631 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18630)
}
__typedArg0 := tmp18630
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18632 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp18631)
}
__typedArg0 := Nil
__typedArg1 := tmp18631
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18625 Obj

if True == tmp18632 {
tmp18627 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V388)
}
__typedArg0 := V388
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18628 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18627)
}
__typedArg0 := tmp18627
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18629 := Call(__e, PrimFunc(symshen_4non_1terminal_2), tmp18628)


var ifres18626 Obj

if True == tmp18629 {
ifres18626 = True


} else {
ifres18626 = False


}

ifres18625 = ifres18626


} else {
ifres18625 = False


}

var ifres18624 Obj

if True == ifres18625 {
ifres18624 = True


} else {
ifres18624 = False


}

ifres18623 = ifres18624


} else {
ifres18623 = False


}

var ifres18622 Obj

if True == ifres18623 {
ifres18622 = True


} else {
ifres18622 = False


}

ifres18621 = ifres18622


} else {
ifres18621 = False


}

var ifres18620 Obj

if True == ifres18621 {
ifres18620 = True


} else {
ifres18620 = False


}

ifres18619 = ifres18620


} else {
ifres18619 = False


}

if True == ifres18619 {
tmp18612 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V388)
}
__typedArg0 := V388
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18612)
}
__typedArg0 := tmp18612
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return


} else {
tmp18617 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V388)
}
__typedArg0 := V388
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp18617 {
tmp18613 := MakeNative(func(__e *ControlFlow) {
Z389 := __e.Get(1)
_ = Z389
__e.TailApply(PrimFunc(symshen_4process_1yacc_1semantics), Z389)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp18613, V388)
return


} else {
tmp18615 := Call(__e, PrimFunc(symshen_4non_1terminal_2), V388)


if True == tmp18615 {
__e.TailApply(PrimFunc(symconcat), symAction, V388)
return
} else {
__e.Return(V388)
return
}


}


}


}, 1)

tmp18638 := Call(__e, ns2_1set, symshen_4process_1yacc_1semantics, tmp18611)


_ = tmp18638

tmp18639 := MakeNative(func(__e *ControlFlow) {
V390 := __e.Get(1)
_ = V390
tmp18640 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V390)
}
__typedArg0 := V390
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18640)
}
__typedArg0 := tmp18640
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return


}, 1)

tmp18641 := Call(__e, ns2_1set, symshen_4_5_1out, tmp18639)


_ = tmp18641

tmp18642 := MakeNative(func(__e *ControlFlow) {
V391 := __e.Get(1)
_ = V391
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V391)
}
__typedArg0 := V391
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return
}, 1)

tmp18643 := Call(__e, ns2_1set, symshen_4in_1_6, tmp18642)


_ = tmp18643

tmp18644 := MakeNative(func(__e *ControlFlow) {
V392 := __e.Get(1)
_ = V392
tmp18645 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V392, Nil)
}
__typedArg0 := V392
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Nil, tmp18645)
}
__typedArg0 := Nil
__typedArg1 := tmp18645
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp18646 := Call(__e, ns2_1set, sym_5_b_6, tmp18644)


_ = tmp18646

tmp18647 := MakeNative(func(__e *ControlFlow) {
V393 := __e.Get(1)
_ = V393
tmp18648 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Nil, Nil)
}
__typedArg0 := Nil
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V393, tmp18648)
}
__typedArg0 := V393
__typedArg1 := tmp18648
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp18649 := Call(__e, ns2_1set, sym_5e_6, tmp18647)


_ = tmp18649

tmp18650 := MakeNative(func(__e *ControlFlow) {
V396 := __e.Get(1)
_ = V396
tmp18653 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V396)
}
__typedArg0 := Nil
__typedArg1 := V396
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp18653 {
tmp18651 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Nil, Nil)
}
__typedArg0 := Nil
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Nil, tmp18651)
}
__typedArg0 := Nil
__typedArg1 := tmp18651
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

__e.TailApply(ns2_1set, sym_5end_6, tmp18650)
return




}, 0)

