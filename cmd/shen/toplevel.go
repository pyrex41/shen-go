package main

import . "github.com/pyrex41/shen-go/kl"

var TopLevelMain = MakeNative(func(__e *ControlFlow) {
tmp8556 := MakeNative(func(__e *ControlFlow) {
tmp8557 := Call(__e, PrimFunc(symshen_4credits))


_ = tmp8557

__e.TailApply(PrimFunc(symshen_4loop))
return


}, 0)

tmp8558 := Call(__e, ns2_1set, symshen_4shen, tmp8556)


_ = tmp8558

tmp8559 := MakeNative(func(__e *ControlFlow) {
tmp8560 := Call(__e, PrimFunc(symshen_4initialise__environment))


_ = tmp8560

tmp8561 := Call(__e, PrimFunc(symshen_4prompt))


_ = tmp8561

tmp8562 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4read_1evaluate_1print))
return
}, 0)

tmp8563 := MakeNative(func(__e *ControlFlow) {
Z5262 := __e.Get(1)
_ = Z5262
tmp8564 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symerror_1to_1string) {
return PrimErrorToString(Z5262)
}
__typedArg0 := Z5262
return Call(__e, PrimFunc(symerror_1to_1string), __typedArg0)
})()

tmp8565 := Call(__e, PrimFunc(symstoutput))


tmp8566 := Call(__e, PrimFunc(sympr), tmp8564, tmp8565)


_ = tmp8566

__e.TailApply(PrimFunc(symnl), MakeNumber(0))
return


}, 1)

tmp8567 := Call(__e, try_1catch, tmp8562, tmp8563)


_ = tmp8567

__e.TailApply(PrimFunc(symshen_4loop))
return


}, 0)

tmp8568 := Call(__e, ns2_1set, symshen_4loop, tmp8559)


_ = tmp8568

tmp8569 := MakeNative(func(__e *ControlFlow) {
tmp8570 := Call(__e, PrimFunc(symstoutput))


tmp8571 := Call(__e, PrimFunc(sympr), MakeString("\nShen, www.shenlanguage.org, copyright (C) 2010-2024, Mark Tarver\n"), tmp8570)


_ = tmp8571

tmp8572 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dversion_d)
}
__typedArg0 := sym_dversion_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp8573 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dlanguage_d)
}
__typedArg0 := sym_dlanguage_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp8574 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dimplementation_d)
}
__typedArg0 := sym_dimplementation_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp8575 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_drelease_d)
}
__typedArg0 := sym_drelease_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp8576 := Call(__e, PrimFunc(symshen_4app), tmp8575, MakeString("\n"), symshen_4a)


tmp8578 := Call(__e, PrimFunc(symshen_4app), tmp8574, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(" "))
__typedS1, __typedOK1 := TypedString(tmp8576)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(" ")
__typedArg1 := tmp8576
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


tmp8580 := Call(__e, PrimFunc(symshen_4app), tmp8573, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(", platform: "))
__typedS1, __typedOK1 := TypedString(tmp8578)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(", platform: ")
__typedArg1 := tmp8578
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


tmp8582 := Call(__e, PrimFunc(symshen_4app), tmp8572, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(", language: "))
__typedS1, __typedOK1 := TypedString(tmp8580)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(", language: ")
__typedArg1 := tmp8580
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


tmp8583 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("version: S"))
__typedS1, __typedOK1 := TypedString(tmp8582)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("version: S")
__typedArg1 := tmp8582
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp8584 := Call(__e, PrimFunc(symstoutput))


tmp8585 := Call(__e, PrimFunc(sympr), tmp8583, tmp8584)


_ = tmp8585

tmp8586 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dport_d)
}
__typedArg0 := sym_dport_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp8587 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dporters_d)
}
__typedArg0 := sym_dporters_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp8588 := Call(__e, PrimFunc(symshen_4app), tmp8587, MakeString("\n\n"), symshen_4a)


tmp8590 := Call(__e, PrimFunc(symshen_4app), tmp8586, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(", ported by "))
__typedS1, __typedOK1 := TypedString(tmp8588)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(", ported by ")
__typedArg1 := tmp8588
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


tmp8591 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("port "))
__typedS1, __typedOK1 := TypedString(tmp8590)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("port ")
__typedArg1 := tmp8590
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp8592 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp8591, tmp8592)
return


}, 0)

tmp8593 := Call(__e, ns2_1set, symshen_4credits, tmp8569)


_ = tmp8593

tmp8594 := MakeNative(func(__e *ControlFlow) {
tmp8595 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dcall_d, MakeNumber(0))
}
__typedArg0 := symshen_4_dcall_d
__typedArg1 := MakeNumber(0)
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp8595

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dinfs_d, MakeNumber(0))
}
__typedArg0 := symshen_4_dinfs_d
__typedArg1 := MakeNumber(0)
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return


}, 0)

tmp8596 := Call(__e, ns2_1set, symshen_4initialise__environment, tmp8594)


_ = tmp8596

tmp8597 := MakeNative(func(__e *ControlFlow) {
tmp8609 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dtc_d)
}
__typedArg0 := symshen_4_dtc_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

if True == tmp8609 {
tmp8598 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dhistory_d)
}
__typedArg0 := symshen_4_dhistory_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp8599 := Call(__e, PrimFunc(symlength), tmp8598)


tmp8600 := Call(__e, PrimFunc(symshen_4app), tmp8599, MakeString("+) "), symshen_4a)


tmp8601 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("\n("))
__typedS1, __typedOK1 := TypedString(tmp8600)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("\n(")
__typedArg1 := tmp8600
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp8602 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp8601, tmp8602)
return


} else {
tmp8603 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dhistory_d)
}
__typedArg0 := symshen_4_dhistory_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp8604 := Call(__e, PrimFunc(symlength), tmp8603)


tmp8605 := Call(__e, PrimFunc(symshen_4app), tmp8604, MakeString("-) "), symshen_4a)


tmp8606 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("\n("))
__typedS1, __typedOK1 := TypedString(tmp8605)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("\n(")
__typedArg1 := tmp8605
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp8607 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp8606, tmp8607)
return


}


}, 0)

tmp8610 := Call(__e, ns2_1set, symshen_4prompt, tmp8597)


_ = tmp8610

tmp8611 := MakeNative(func(__e *ControlFlow) {
tmp8612 := MakeNative(func(__e *ControlFlow) {
W5263 := __e.Get(1)
_ = W5263
tmp8613 := MakeNative(func(__e *ControlFlow) {
W5264 := __e.Get(1)
_ = W5264
tmp8614 := MakeNative(func(__e *ControlFlow) {
W5265 := __e.Get(1)
_ = W5265
tmp8615 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dtc_d)
}
__typedArg0 := symshen_4_dtc_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4evaluate_1lineread), W5264, W5265, tmp8615)
return


}, 1)

tmp8616 := Call(__e, PrimFunc(symshen_4update_1history))


__e.TailApply(tmp8614, tmp8616)
return


}, 1)

tmp8617 := Call(__e, PrimFunc(symstinput))


tmp8618 := Call(__e, PrimFunc(symlineread), tmp8617)


tmp8619 := Call(__e, PrimFunc(symshen_4package_1user_1input), W5263, tmp8618)


__e.TailApply(tmp8613, tmp8619)
return


}, 1)

tmp8620 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dpackage_d)
}
__typedArg0 := symshen_4_dpackage_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(tmp8612, tmp8620)
return


}, 0)

tmp8621 := Call(__e, ns2_1set, symshen_4read_1evaluate_1print, tmp8611)


_ = tmp8621

tmp8622 := MakeNative(func(__e *ControlFlow) {
V5266 := __e.Get(1)
_ = V5266
V5267 := __e.Get(2)
_ = V5267
tmp8629 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symnull, V5266)
}
__typedArg0 := symnull
__typedArg1 := V5266
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp8629 {
__e.Return(V5267)
return
} else {
tmp8623 := MakeNative(func(__e *ControlFlow) {
W5268 := __e.Get(1)
_ = W5268
tmp8624 := MakeNative(func(__e *ControlFlow) {
W5269 := __e.Get(1)
_ = W5269
tmp8625 := MakeNative(func(__e *ControlFlow) {
Z5270 := __e.Get(1)
_ = Z5270
__e.TailApply(PrimFunc(symshen_4pui_1h), W5268, W5269, Z5270)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp8625, V5267)
return


}, 1)

tmp8626 := Call(__e, PrimFunc(symexternal), V5266)


__e.TailApply(tmp8624, tmp8626)
return


}, 1)

tmp8627 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(V5266)
}
__typedArg0 := V5266
return Call(__e, PrimFunc(symstr), __typedArg0)
})()

__e.TailApply(tmp8623, tmp8627)
return


}


}, 2)

tmp8630 := Call(__e, ns2_1set, symshen_4package_1user_1input, tmp8622)


_ = tmp8630

tmp8631 := MakeNative(func(__e *ControlFlow) {
V5275 := __e.Get(1)
_ = V5275
V5276 := __e.Get(2)
_ = V5276
V5277 := __e.Get(3)
_ = V5277
tmp8672 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5277)
}
__typedArg0 := V5277
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8659 Obj

if True == tmp8672 {
tmp8670 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5277)
}
__typedArg0 := V5277
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8671 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symfn, tmp8670)
}
__typedArg0 := symfn
__typedArg1 := tmp8670
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8661 Obj

if True == tmp8671 {
tmp8668 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5277)
}
__typedArg0 := V5277
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8669 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp8668)
}
__typedArg0 := tmp8668
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8663 Obj

if True == tmp8669 {
tmp8665 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5277)
}
__typedArg0 := V5277
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8666 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp8665)
}
__typedArg0 := tmp8665
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8667 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp8666)
}
__typedArg0 := Nil
__typedArg1 := tmp8666
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8664 Obj

if True == tmp8667 {
ifres8664 = True


} else {
ifres8664 = False


}

ifres8663 = ifres8664


} else {
ifres8663 = False


}

var ifres8662 Obj

if True == ifres8663 {
ifres8662 = True


} else {
ifres8662 = False


}

ifres8661 = ifres8662


} else {
ifres8661 = False


}

var ifres8660 Obj

if True == ifres8661 {
ifres8660 = True


} else {
ifres8660 = False


}

ifres8659 = ifres8660


} else {
ifres8659 = False


}

if True == ifres8659 {
tmp8637 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5277)
}
__typedArg0 := V5277
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8638 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp8637)
}
__typedArg0 := tmp8637
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8639 := Call(__e, PrimFunc(symshen_4internal_2), tmp8638, V5275, V5276)


if True == tmp8639 {
tmp8632 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5277)
}
__typedArg0 := V5277
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8633 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp8632)
}
__typedArg0 := tmp8632
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8634 := Call(__e, PrimFunc(symshen_4intern_1in_1package), V5275, tmp8633)


tmp8635 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8634, Nil)
}
__typedArg0 := tmp8634
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfn, tmp8635)
}
__typedArg0 := symfn
__typedArg1 := tmp8635
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return(V5277)
return
}


} else {
tmp8657 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5277)
}
__typedArg0 := V5277
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp8657 {
tmp8654 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5277)
}
__typedArg0 := V5277
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8655 := Call(__e, PrimFunc(symshen_4internal_2), tmp8654, V5275, V5276)


if True == tmp8655 {
tmp8640 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5277)
}
__typedArg0 := V5277
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8641 := Call(__e, PrimFunc(symshen_4intern_1in_1package), V5275, tmp8640)


tmp8642 := MakeNative(func(__e *ControlFlow) {
Z5278 := __e.Get(1)
_ = Z5278
__e.TailApply(PrimFunc(symshen_4pui_1h), V5275, V5276, Z5278)
return
}, 1)

tmp8643 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5277)
}
__typedArg0 := V5277
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8644 := Call(__e, PrimFunc(symmap), tmp8642, tmp8643)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8641, tmp8644)
}
__typedArg0 := tmp8641
__typedArg1 := tmp8644
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp8651 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5277)
}
__typedArg0 := V5277
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8652 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp8651)
}
__typedArg0 := tmp8651
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp8652 {
tmp8645 := MakeNative(func(__e *ControlFlow) {
Z5279 := __e.Get(1)
_ = Z5279
__e.TailApply(PrimFunc(symshen_4pui_1h), V5275, V5276, Z5279)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp8645, V5277)
return


} else {
tmp8646 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5277)
}
__typedArg0 := V5277
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8647 := MakeNative(func(__e *ControlFlow) {
Z5280 := __e.Get(1)
_ = Z5280
__e.TailApply(PrimFunc(symshen_4pui_1h), V5275, V5276, Z5280)
return
}, 1)

tmp8648 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5277)
}
__typedArg0 := V5277
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8649 := Call(__e, PrimFunc(symmap), tmp8647, tmp8648)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8646, tmp8649)
}
__typedArg0 := tmp8646
__typedArg1 := tmp8649
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}


}


} else {
__e.Return(V5277)
return
}


}


}, 3)

tmp8673 := Call(__e, ns2_1set, symshen_4pui_1h, tmp8631)


_ = tmp8673

tmp8674 := MakeNative(func(__e *ControlFlow) {
tmp8675 := Call(__e, PrimFunc(symit))


tmp8676 := Call(__e, PrimFunc(symshen_4trim_1it), tmp8675)


tmp8677 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dhistory_d)
}
__typedArg0 := symshen_4_dhistory_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp8678 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8676, tmp8677)
}
__typedArg0 := tmp8676
__typedArg1 := tmp8677
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dhistory_d, tmp8678)
}
__typedArg0 := symshen_4_dhistory_d
__typedArg1 := tmp8678
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return


}, 0)

tmp8679 := Call(__e, ns2_1set, symshen_4update_1history, tmp8674)


_ = tmp8679

tmp8680 := MakeNative(func(__e *ControlFlow) {
V5281 := __e.Get(1)
_ = V5281
tmp8688 := Call(__e, PrimFunc(symshen_4_7string_2), V5281)


var ifres8683 Obj

if True == tmp8688 {
tmp8685 := Call(__e, PrimFunc(symhdstr), V5281)


tmp8686 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_1_6n) {
return PrimStringToNumber(tmp8685)
}
__typedArg0 := tmp8685
return Call(__e, PrimFunc(symstring_1_6n), __typedArg0)
})()

tmp8687 := Call(__e, PrimFunc(symshen_4whitespace_2), tmp8686)


var ifres8684 Obj

if True == tmp8687 {
ifres8684 = True


} else {
ifres8684 = False


}

ifres8683 = ifres8684


} else {
ifres8683 = False


}

if True == ifres8683 {
__e.TailApply(PrimFunc(symshen_4trim_1it), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V5281)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V5281
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())
return


} else {
__e.Return(V5281)
return
}


}, 1)

tmp8689 := Call(__e, ns2_1set, symshen_4trim_1it, tmp8680)


_ = tmp8689

tmp8690 := MakeNative(func(__e *ControlFlow) {
V5300 := __e.Get(1)
_ = V5300
V5301 := __e.Get(2)
_ = V5301
V5302 := __e.Get(3)
_ = V5302
tmp8820 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5300)
}
__typedArg0 := V5300
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8789 Obj

if True == tmp8820 {
tmp8818 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5300)
}
__typedArg0 := V5300
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8819 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp8818)
}
__typedArg0 := Nil
__typedArg1 := tmp8818
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8791 Obj

if True == tmp8819 {
tmp8817 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5301)
}
__typedArg0 := V5301
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8793 Obj

if True == tmp8817 {
tmp8815 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5301)
}
__typedArg0 := V5301
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8816 := Call(__e, PrimFunc(symshen_4_7string_2), tmp8815)


var ifres8795 Obj

if True == tmp8816 {
tmp8812 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5301)
}
__typedArg0 := V5301
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8813 := Call(__e, PrimFunc(symhdstr), tmp8812)


tmp8814 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("!"), tmp8813)
}
__typedArg0 := MakeString("!")
__typedArg1 := tmp8813
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8797 Obj

if True == tmp8814 {
tmp8809 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5301)
}
__typedArg0 := V5301
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8811 := Call(__e, PrimFunc(symshen_4_7string_2), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(tmp8809)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := tmp8809
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


var ifres8799 Obj

if True == tmp8811 {
tmp8805 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5301)
}
__typedArg0 := V5301
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8807 := Call(__e, PrimFunc(symhdstr), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(tmp8805)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := tmp8805
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


tmp8808 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("!"), tmp8807)
}
__typedArg0 := MakeString("!")
__typedArg1 := tmp8807
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8801 Obj

if True == tmp8808 {
tmp8803 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5301)
}
__typedArg0 := V5301
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8804 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp8803)
}
__typedArg0 := tmp8803
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8802 Obj

if True == tmp8804 {
ifres8802 = True


} else {
ifres8802 = False


}

ifres8801 = ifres8802


} else {
ifres8801 = False


}

var ifres8800 Obj

if True == ifres8801 {
ifres8800 = True


} else {
ifres8800 = False


}

ifres8799 = ifres8800


} else {
ifres8799 = False


}

var ifres8798 Obj

if True == ifres8799 {
ifres8798 = True


} else {
ifres8798 = False


}

ifres8797 = ifres8798


} else {
ifres8797 = False


}

var ifres8796 Obj

if True == ifres8797 {
ifres8796 = True


} else {
ifres8796 = False


}

ifres8795 = ifres8796


} else {
ifres8795 = False


}

var ifres8794 Obj

if True == ifres8795 {
ifres8794 = True


} else {
ifres8794 = False


}

ifres8793 = ifres8794


} else {
ifres8793 = False


}

var ifres8792 Obj

if True == ifres8793 {
ifres8792 = True


} else {
ifres8792 = False


}

ifres8791 = ifres8792


} else {
ifres8791 = False


}

var ifres8790 Obj

if True == ifres8791 {
ifres8790 = True


} else {
ifres8790 = False


}

ifres8789 = ifres8790


} else {
ifres8789 = False


}

if True == ifres8789 {
tmp8691 := MakeNative(func(__e *ControlFlow) {
W5303 := __e.Get(1)
_ = W5303
tmp8692 := MakeNative(func(__e *ControlFlow) {
W5304 := __e.Get(1)
_ = W5304
tmp8693 := MakeNative(func(__e *ControlFlow) {
W5305 := __e.Get(1)
_ = W5305
__e.TailApply(PrimFunc(symshen_4evaluate_1lineread), W5303, W5304, V5302)
return
}, 1)

tmp8694 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5301)
}
__typedArg0 := V5301
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8695 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp8694)
}
__typedArg0 := tmp8694
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8696 := Call(__e, PrimFunc(symshen_4app), tmp8695, MakeString("\n"), symshen_4a)


tmp8697 := Call(__e, PrimFunc(symstoutput))


tmp8698 := Call(__e, PrimFunc(sympr), tmp8696, tmp8697)


__e.TailApply(tmp8693, tmp8698)
return


}, 1)

tmp8699 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5301)
}
__typedArg0 := V5301
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8700 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp8699)
}
__typedArg0 := tmp8699
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8701 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5301)
}
__typedArg0 := V5301
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8702 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8700, tmp8701)
}
__typedArg0 := tmp8700
__typedArg1 := tmp8701
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8703 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dhistory_d, tmp8702)
}
__typedArg0 := symshen_4_dhistory_d
__typedArg1 := tmp8702
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

__e.TailApply(tmp8692, tmp8703)
return


}, 1)

tmp8704 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5301)
}
__typedArg0 := V5301
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8705 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp8704)
}
__typedArg0 := tmp8704
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8706 := Call(__e, PrimFunc(symread_1from_1string), tmp8705)


__e.TailApply(tmp8691, tmp8706)
return


} else {
tmp8787 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5300)
}
__typedArg0 := V5300
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8771 Obj

if True == tmp8787 {
tmp8785 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5300)
}
__typedArg0 := V5300
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8786 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp8785)
}
__typedArg0 := Nil
__typedArg1 := tmp8785
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8773 Obj

if True == tmp8786 {
tmp8784 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5301)
}
__typedArg0 := V5301
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8775 Obj

if True == tmp8784 {
tmp8782 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5301)
}
__typedArg0 := V5301
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8783 := Call(__e, PrimFunc(symshen_4_7string_2), tmp8782)


var ifres8777 Obj

if True == tmp8783 {
tmp8779 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5301)
}
__typedArg0 := V5301
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8780 := Call(__e, PrimFunc(symhdstr), tmp8779)


tmp8781 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("!"), tmp8780)
}
__typedArg0 := MakeString("!")
__typedArg1 := tmp8780
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8778 Obj

if True == tmp8781 {
ifres8778 = True


} else {
ifres8778 = False


}

ifres8777 = ifres8778


} else {
ifres8777 = False


}

var ifres8776 Obj

if True == ifres8777 {
ifres8776 = True


} else {
ifres8776 = False


}

ifres8775 = ifres8776


} else {
ifres8775 = False


}

var ifres8774 Obj

if True == ifres8775 {
ifres8774 = True


} else {
ifres8774 = False


}

ifres8773 = ifres8774


} else {
ifres8773 = False


}

var ifres8772 Obj

if True == ifres8773 {
ifres8772 = True


} else {
ifres8772 = False


}

ifres8771 = ifres8772


} else {
ifres8771 = False


}

if True == ifres8771 {
tmp8707 := MakeNative(func(__e *ControlFlow) {
W5306 := __e.Get(1)
_ = W5306
tmp8708 := MakeNative(func(__e *ControlFlow) {
W5307 := __e.Get(1)
_ = W5307
tmp8709 := MakeNative(func(__e *ControlFlow) {
W5308 := __e.Get(1)
_ = W5308
tmp8710 := MakeNative(func(__e *ControlFlow) {
W5309 := __e.Get(1)
_ = W5309
tmp8711 := MakeNative(func(__e *ControlFlow) {
W5310 := __e.Get(1)
_ = W5310
__e.TailApply(PrimFunc(symshen_4evaluate_1lineread), W5309, W5310, V5302)
return
}, 1)

tmp8712 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5301)
}
__typedArg0 := V5301
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8713 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W5307, tmp8712)
}
__typedArg0 := W5307
__typedArg1 := tmp8712
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8714 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dhistory_d, tmp8713)
}
__typedArg0 := symshen_4_dhistory_d
__typedArg1 := tmp8713
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

__e.TailApply(tmp8711, tmp8714)
return


}, 1)

tmp8715 := Call(__e, PrimFunc(symread_1from_1string), W5307)


__e.TailApply(tmp8710, tmp8715)
return


}, 1)

tmp8716 := Call(__e, PrimFunc(symshen_4app), W5307, MakeString("\n"), symshen_4a)


tmp8717 := Call(__e, PrimFunc(symstoutput))


tmp8718 := Call(__e, PrimFunc(sympr), tmp8716, tmp8717)


__e.TailApply(tmp8709, tmp8718)
return


}, 1)

tmp8719 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5301)
}
__typedArg0 := V5301
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8720 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(tmp8719)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := tmp8719
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()

tmp8721 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5301)
}
__typedArg0 := V5301
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8722 := Call(__e, PrimFunc(symshen_4use_1history), W5306, tmp8720, tmp8721)


__e.TailApply(tmp8708, tmp8722)
return


}, 1)

tmp8728 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5301)
}
__typedArg0 := V5301
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8730 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(tmp8728)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := tmp8728
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})(), MakeString(""))
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(tmp8728)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := tmp8728
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()
__typedArg1 := MakeString("")
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8723 Obj

if True == tmp8730 {
ifres8723 = Nil


} else {
tmp8724 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5301)
}
__typedArg0 := V5301
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8726 := Call(__e, PrimFunc(symread_1from_1string), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(tmp8724)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := tmp8724
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


tmp8727 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp8726)
}
__typedArg0 := tmp8726
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

ifres8723 = tmp8727


}

__e.TailApply(tmp8707, ifres8723)
return


} else {
tmp8769 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5300)
}
__typedArg0 := V5300
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8753 Obj

if True == tmp8769 {
tmp8767 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5300)
}
__typedArg0 := V5300
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8768 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp8767)
}
__typedArg0 := Nil
__typedArg1 := tmp8767
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8755 Obj

if True == tmp8768 {
tmp8766 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5301)
}
__typedArg0 := V5301
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8757 Obj

if True == tmp8766 {
tmp8764 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5301)
}
__typedArg0 := V5301
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8765 := Call(__e, PrimFunc(symshen_4_7string_2), tmp8764)


var ifres8759 Obj

if True == tmp8765 {
tmp8761 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5301)
}
__typedArg0 := V5301
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8762 := Call(__e, PrimFunc(symhdstr), tmp8761)


tmp8763 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("%"), tmp8762)
}
__typedArg0 := MakeString("%")
__typedArg1 := tmp8762
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8760 Obj

if True == tmp8763 {
ifres8760 = True


} else {
ifres8760 = False


}

ifres8759 = ifres8760


} else {
ifres8759 = False


}

var ifres8758 Obj

if True == ifres8759 {
ifres8758 = True


} else {
ifres8758 = False


}

ifres8757 = ifres8758


} else {
ifres8757 = False


}

var ifres8756 Obj

if True == ifres8757 {
ifres8756 = True


} else {
ifres8756 = False


}

ifres8755 = ifres8756


} else {
ifres8755 = False


}

var ifres8754 Obj

if True == ifres8755 {
ifres8754 = True


} else {
ifres8754 = False


}

ifres8753 = ifres8754


} else {
ifres8753 = False


}

if True == ifres8753 {
tmp8731 := MakeNative(func(__e *ControlFlow) {
W5311 := __e.Get(1)
_ = W5311
tmp8732 := MakeNative(func(__e *ControlFlow) {
W5312 := __e.Get(1)
_ = W5312
tmp8733 := MakeNative(func(__e *ControlFlow) {
W5313 := __e.Get(1)
_ = W5313
__e.TailApply(PrimFunc(symabort))
return
}, 1)

tmp8734 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5301)
}
__typedArg0 := V5301
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8735 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dhistory_d, tmp8734)
}
__typedArg0 := symshen_4_dhistory_d
__typedArg1 := tmp8734
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

__e.TailApply(tmp8733, tmp8735)
return


}, 1)

tmp8736 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5301)
}
__typedArg0 := V5301
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8737 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(tmp8736)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := tmp8736
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()

tmp8738 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5301)
}
__typedArg0 := V5301
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8739 := Call(__e, PrimFunc(symshen_4peek_1history), W5311, tmp8737, tmp8738)


__e.TailApply(tmp8732, tmp8739)
return


}, 1)

tmp8745 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5301)
}
__typedArg0 := V5301
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8747 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(tmp8745)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := tmp8745
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})(), MakeString(""))
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(tmp8745)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := tmp8745
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()
__typedArg1 := MakeString("")
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8740 Obj

if True == tmp8747 {
ifres8740 = Nil


} else {
tmp8741 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5301)
}
__typedArg0 := V5301
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8743 := Call(__e, PrimFunc(symread_1from_1string), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(tmp8741)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := tmp8741
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


tmp8744 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp8743)
}
__typedArg0 := tmp8743
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

ifres8740 = tmp8744


}

__e.TailApply(tmp8731, ifres8740)
return


} else {
tmp8751 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(True, V5302)
}
__typedArg0 := True
__typedArg1 := V5302
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp8751 {
__e.TailApply(PrimFunc(symshen_4check_1eval_1and_1print), V5300)
return
} else {
tmp8749 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(False, V5302)
}
__typedArg0 := False
__typedArg1 := V5302
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp8749 {
__e.TailApply(PrimFunc(symshen_4eval_1and_1print), V5300)
return
} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.evaluate-lineread"))
}
__typedArg0 := MakeString("implementation error in shen.evaluate-lineread")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}


}


}


}, 3)

tmp8821 := Call(__e, ns2_1set, symshen_4evaluate_1lineread, tmp8690)


_ = tmp8821

tmp8822 := MakeNative(func(__e *ControlFlow) {
V5314 := __e.Get(1)
_ = V5314
V5315 := __e.Get(2)
_ = V5315
V5316 := __e.Get(3)
_ = V5316
tmp8828 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(syminteger_2) {
return PrimIsInteger(V5314)
}
__typedArg0 := V5314
return Call(__e, PrimFunc(syminteger_2), __typedArg0)
})()

if True == tmp8828 {
tmp8823 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(1))
__typedN1, __typedOK1 := TypedFloat64(V5314)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := MakeNumber(1)
__typedArg1 := V5314
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()

tmp8824 := Call(__e, PrimFunc(symreverse), V5316)


__e.TailApply(PrimFunc(symnth), tmp8823, tmp8824)
return


} else {
tmp8826 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsymbol_2) {
return PrimIsSymbol(V5314)
}
__typedArg0 := V5314
return Call(__e, PrimFunc(symsymbol_2), __typedArg0)
})()

if True == tmp8826 {
__e.TailApply(PrimFunc(symshen_4string_1match), V5315, V5316)
return
} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("! expects a number or a symbol\n"))
}
__typedArg0 := MakeString("! expects a number or a symbol\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 3)

tmp8829 := Call(__e, ns2_1set, symshen_4use_1history, tmp8822)


_ = tmp8829

tmp8830 := MakeNative(func(__e *ControlFlow) {
V5317 := __e.Get(1)
_ = V5317
V5318 := __e.Get(2)
_ = V5318
V5319 := __e.Get(3)
_ = V5319
tmp8844 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(syminteger_2) {
return PrimIsInteger(V5317)
}
__typedArg0 := V5317
return Call(__e, PrimFunc(syminteger_2), __typedArg0)
})()

if True == tmp8844 {
tmp8831 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(1))
__typedN1, __typedOK1 := TypedFloat64(V5317)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := MakeNumber(1)
__typedArg1 := V5317
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()

tmp8832 := Call(__e, PrimFunc(symreverse), V5319)


tmp8833 := Call(__e, PrimFunc(symnth), tmp8831, tmp8832)


tmp8834 := Call(__e, PrimFunc(symshen_4app), tmp8833, MakeString(""), symshen_4a)


tmp8835 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("\n"))
__typedS1, __typedOK1 := TypedString(tmp8834)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("\n")
__typedArg1 := tmp8834
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp8836 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp8835, tmp8836)
return


} else {
tmp8842 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V5318, MakeString(""))
}
__typedArg0 := V5318
__typedArg1 := MakeString("")
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8839 Obj

if True == tmp8842 {
ifres8839 = True


} else {
tmp8841 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsymbol_2) {
return PrimIsSymbol(V5317)
}
__typedArg0 := V5317
return Call(__e, PrimFunc(symsymbol_2), __typedArg0)
})()

var ifres8840 Obj

if True == tmp8841 {
ifres8840 = True


} else {
ifres8840 = False


}

ifres8839 = ifres8840


}

if True == ifres8839 {
tmp8837 := Call(__e, PrimFunc(symreverse), V5319)


__e.TailApply(PrimFunc(symshen_4recursive_1string_1match), MakeNumber(0), V5318, tmp8837)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("% expects a number or a symbol\n"))
}
__typedArg0 := MakeString("% expects a number or a symbol\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 3)

tmp8845 := Call(__e, ns2_1set, symshen_4peek_1history, tmp8830)


_ = tmp8845

tmp8846 := MakeNative(func(__e *ControlFlow) {
V5329 := __e.Get(1)
_ = V5329
V5330 := __e.Get(2)
_ = V5330
tmp8857 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V5330)
}
__typedArg0 := Nil
__typedArg1 := V5330
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp8857 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("\ninput not found"))
}
__typedArg0 := MakeString("\ninput not found")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
} else {
tmp8855 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5330)
}
__typedArg0 := V5330
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8851 Obj

if True == tmp8855 {
tmp8853 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5330)
}
__typedArg0 := V5330
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8854 := Call(__e, PrimFunc(symshen_4string_1prefix_2), V5329, tmp8853)


var ifres8852 Obj

if True == tmp8854 {
ifres8852 = True


} else {
ifres8852 = False


}

ifres8851 = ifres8852


} else {
ifres8851 = False


}

if True == ifres8851 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5330)
}
__typedArg0 := V5330
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return
} else {
tmp8849 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5330)
}
__typedArg0 := V5330
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp8849 {
tmp8847 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5330)
}
__typedArg0 := V5330
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4string_1match), V5329, tmp8847)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.string-match"))
}
__typedArg0 := MakeString("implementation error in shen.string-match")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}


}, 2)

tmp8858 := Call(__e, ns2_1set, symshen_4string_1match, tmp8846)


_ = tmp8858

tmp8859 := MakeNative(func(__e *ControlFlow) {
V5338 := __e.Get(1)
_ = V5338
V5339 := __e.Get(2)
_ = V5339
tmp8896 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString(""), V5338)
}
__typedArg0 := MakeString("")
__typedArg1 := V5338
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp8896 {
__e.Return(True)
return
} else {
tmp8894 := Call(__e, PrimFunc(symshen_4_7string_2), V5338)


var ifres8889 Obj

if True == tmp8894 {
tmp8891 := Call(__e, PrimFunc(symhdstr), V5338)


tmp8892 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_1_6n) {
return PrimStringToNumber(tmp8891)
}
__typedArg0 := tmp8891
return Call(__e, PrimFunc(symstring_1_6n), __typedArg0)
})()

tmp8893 := Call(__e, PrimFunc(symshen_4whitespace_2), tmp8892)


var ifres8890 Obj

if True == tmp8893 {
ifres8890 = True


} else {
ifres8890 = False


}

ifres8889 = ifres8890


} else {
ifres8889 = False


}

if True == ifres8889 {
__e.TailApply(PrimFunc(symshen_4string_1prefix_2), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V5338)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V5338
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})(), V5339)
return


} else {
tmp8887 := Call(__e, PrimFunc(symshen_4_7string_2), V5339)


var ifres8882 Obj

if True == tmp8887 {
tmp8884 := Call(__e, PrimFunc(symhdstr), V5339)


tmp8885 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_1_6n) {
return PrimStringToNumber(tmp8884)
}
__typedArg0 := tmp8884
return Call(__e, PrimFunc(symstring_1_6n), __typedArg0)
})()

tmp8886 := Call(__e, PrimFunc(symshen_4whitespace_2), tmp8885)


var ifres8883 Obj

if True == tmp8886 {
ifres8883 = True


} else {
ifres8883 = False


}

ifres8882 = ifres8883


} else {
ifres8882 = False


}

if True == ifres8882 {
__e.TailApply(PrimFunc(symshen_4string_1prefix_2), V5338, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V5339)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V5339
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())
return


} else {
tmp8880 := Call(__e, PrimFunc(symshen_4_7string_2), V5339)


var ifres8876 Obj

if True == tmp8880 {
tmp8878 := Call(__e, PrimFunc(symhdstr), V5339)


tmp8879 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("("), tmp8878)
}
__typedArg0 := MakeString("(")
__typedArg1 := tmp8878
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8877 Obj

if True == tmp8879 {
ifres8877 = True


} else {
ifres8877 = False


}

ifres8876 = ifres8877


} else {
ifres8876 = False


}

if True == ifres8876 {
__e.TailApply(PrimFunc(symshen_4string_1prefix_2), V5338, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V5339)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V5339
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())
return


} else {
tmp8874 := Call(__e, PrimFunc(symshen_4_7string_2), V5338)


var ifres8866 Obj

if True == tmp8874 {
tmp8873 := Call(__e, PrimFunc(symshen_4_7string_2), V5339)


var ifres8868 Obj

if True == tmp8873 {
tmp8870 := Call(__e, PrimFunc(symhdstr), V5338)


tmp8871 := Call(__e, PrimFunc(symhdstr), V5339)


tmp8872 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp8870, tmp8871)
}
__typedArg0 := tmp8870
__typedArg1 := tmp8871
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8869 Obj

if True == tmp8872 {
ifres8869 = True


} else {
ifres8869 = False


}

ifres8868 = ifres8869


} else {
ifres8868 = False


}

var ifres8867 Obj

if True == ifres8868 {
ifres8867 = True


} else {
ifres8867 = False


}

ifres8866 = ifres8867


} else {
ifres8866 = False


}

if True == ifres8866 {
tmp8863 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V5338)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V5338
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4string_1prefix_2), tmp8863, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V5339)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V5339
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())
return


} else {
__e.Return(False)
return
}


}


}


}


}


}, 2)

tmp8897 := Call(__e, ns2_1set, symshen_4string_1prefix_2, tmp8859)


_ = tmp8897

tmp8898 := MakeNative(func(__e *ControlFlow) {
V5350 := __e.Get(1)
_ = V5350
V5351 := __e.Get(2)
_ = V5351
V5352 := __e.Get(3)
_ = V5352
tmp8913 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V5352)
}
__typedArg0 := Nil
__typedArg1 := V5352
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp8913 {
__e.Return(symshen_4skip)
return
} else {
tmp8911 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5352)
}
__typedArg0 := V5352
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp8911 {
tmp8906 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5352)
}
__typedArg0 := V5352
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8907 := Call(__e, PrimFunc(symshen_4string_1prefix_2), V5351, tmp8906)


var ifres8899 Obj

if True == tmp8907 {
tmp8900 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5352)
}
__typedArg0 := V5352
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8901 := Call(__e, PrimFunc(symshen_4app), tmp8900, MakeString("\n"), symshen_4a)


tmp8903 := Call(__e, PrimFunc(symshen_4app), V5350, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(". "))
__typedS1, __typedOK1 := TypedString(tmp8901)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(". ")
__typedArg1 := tmp8901
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


tmp8904 := Call(__e, PrimFunc(symstoutput))


tmp8905 := Call(__e, PrimFunc(sympr), tmp8903, tmp8904)


ifres8899 = tmp8905


} else {
ifres8899 = symshen_4skip


}

_ = ifres8899

tmp8908 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V5350)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V5350
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()

tmp8909 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5352)
}
__typedArg0 := V5352
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4recursive_1string_1match), tmp8908, V5351, tmp8909)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.recursive-string-match"))
}
__typedArg0 := MakeString("implementation error in shen.recursive-string-match")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 3)

__e.TailApply(ns2_1set, symshen_4recursive_1string_1match, tmp8898)
return




}, 0)

