package main

import . "github.com/pyrex41/shen-go/kl"

var TopLevelMain = MakeNative(func(__e *ControlFlow) {
tmp8729 := MakeNative(func(__e *ControlFlow) {
tmp8730 := Call(__e, PrimFunc(symshen_4credits))


_ = tmp8730

__e.TailApply(PrimFunc(symshen_4loop))
return


}, 0)

tmp8731 := Call(__e, ns2_1set, symshen_4shen, tmp8729)


_ = tmp8731

tmp8732 := MakeNative(func(__e *ControlFlow) {
tmp8733 := Call(__e, PrimFunc(symshen_4initialise__environment))


_ = tmp8733

tmp8734 := Call(__e, PrimFunc(symshen_4prompt))


_ = tmp8734

tmp8735 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4read_1evaluate_1print))
return
}, 0)

tmp8736 := MakeNative(func(__e *ControlFlow) {
Z5262 := __e.Get(1)
_ = Z5262
tmp8737 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symerror_1to_1string) {
return PrimErrorToString(Z5262)
}
__typedArg0 := Z5262
return Call(__e, PrimFunc(symerror_1to_1string), __typedArg0)
})()

tmp8738 := Call(__e, PrimFunc(symstoutput))


tmp8739 := Call(__e, PrimFunc(sympr), tmp8737, tmp8738)


_ = tmp8739

__e.TailApply(PrimFunc(symnl), MakeNumber(0))
return


}, 1)

tmp8740 := Call(__e, try_1catch, tmp8735, tmp8736)


_ = tmp8740

__e.TailApply(PrimFunc(symshen_4loop))
return


}, 0)

tmp8741 := Call(__e, ns2_1set, symshen_4loop, tmp8732)


_ = tmp8741

tmp8742 := MakeNative(func(__e *ControlFlow) {
tmp8743 := Call(__e, PrimFunc(symstoutput))


tmp8744 := Call(__e, PrimFunc(sympr), MakeString("\nShen, www.shenlanguage.org, copyright (C) 2010-2024, Mark Tarver\n"), tmp8743)


_ = tmp8744

tmp8745 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dversion_d)
}
__typedArg0 := sym_dversion_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp8746 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dlanguage_d)
}
__typedArg0 := sym_dlanguage_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp8747 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dimplementation_d)
}
__typedArg0 := sym_dimplementation_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp8748 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_drelease_d)
}
__typedArg0 := sym_drelease_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp8749 := Call(__e, PrimFunc(symshen_4app), tmp8748, MakeString("\n"), symshen_4a)


tmp8751 := Call(__e, PrimFunc(symshen_4app), tmp8747, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(" "))
__typedS1, __typedOK1 := TypedString(tmp8749)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(" ")
__typedArg1 := tmp8749
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


tmp8753 := Call(__e, PrimFunc(symshen_4app), tmp8746, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(", platform: "))
__typedS1, __typedOK1 := TypedString(tmp8751)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(", platform: ")
__typedArg1 := tmp8751
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


tmp8755 := Call(__e, PrimFunc(symshen_4app), tmp8745, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(", language: "))
__typedS1, __typedOK1 := TypedString(tmp8753)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(", language: ")
__typedArg1 := tmp8753
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


tmp8756 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("version: S"))
__typedS1, __typedOK1 := TypedString(tmp8755)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("version: S")
__typedArg1 := tmp8755
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp8757 := Call(__e, PrimFunc(symstoutput))


tmp8758 := Call(__e, PrimFunc(sympr), tmp8756, tmp8757)


_ = tmp8758

tmp8759 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dport_d)
}
__typedArg0 := sym_dport_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp8760 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dporters_d)
}
__typedArg0 := sym_dporters_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp8761 := Call(__e, PrimFunc(symshen_4app), tmp8760, MakeString("\n\n"), symshen_4a)


tmp8763 := Call(__e, PrimFunc(symshen_4app), tmp8759, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(", ported by "))
__typedS1, __typedOK1 := TypedString(tmp8761)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(", ported by ")
__typedArg1 := tmp8761
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


tmp8764 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("port "))
__typedS1, __typedOK1 := TypedString(tmp8763)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("port ")
__typedArg1 := tmp8763
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp8765 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp8764, tmp8765)
return


}, 0)

tmp8766 := Call(__e, ns2_1set, symshen_4credits, tmp8742)


_ = tmp8766

tmp8767 := MakeNative(func(__e *ControlFlow) {
tmp8768 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dcall_d, MakeNumber(0))
}
__typedArg0 := symshen_4_dcall_d
__typedArg1 := MakeNumber(0)
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp8768

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

tmp8769 := Call(__e, ns2_1set, symshen_4initialise__environment, tmp8767)


_ = tmp8769

tmp8770 := MakeNative(func(__e *ControlFlow) {
tmp8782 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dtc_d)
}
__typedArg0 := symshen_4_dtc_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

if True == tmp8782 {
tmp8771 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dhistory_d)
}
__typedArg0 := symshen_4_dhistory_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp8772 := Call(__e, PrimFunc(symlength), tmp8771)


tmp8773 := Call(__e, PrimFunc(symshen_4app), tmp8772, MakeString("+) "), symshen_4a)


tmp8774 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("\n("))
__typedS1, __typedOK1 := TypedString(tmp8773)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("\n(")
__typedArg1 := tmp8773
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp8775 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp8774, tmp8775)
return


} else {
tmp8776 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dhistory_d)
}
__typedArg0 := symshen_4_dhistory_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp8777 := Call(__e, PrimFunc(symlength), tmp8776)


tmp8778 := Call(__e, PrimFunc(symshen_4app), tmp8777, MakeString("-) "), symshen_4a)


tmp8779 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("\n("))
__typedS1, __typedOK1 := TypedString(tmp8778)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("\n(")
__typedArg1 := tmp8778
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp8780 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp8779, tmp8780)
return


}


}, 0)

tmp8783 := Call(__e, ns2_1set, symshen_4prompt, tmp8770)


_ = tmp8783

tmp8784 := MakeNative(func(__e *ControlFlow) {
tmp8785 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dpackage_d)
}
__typedArg0 := symshen_4_dpackage_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

W52638713 := tmp8785
_ = W52638713

tmp8786 := Call(__e, PrimFunc(symstinput))


tmp8787 := Call(__e, PrimFunc(symlineread), tmp8786)


tmp8788 := Call(__e, PrimFunc(symshen_4package_1user_1input), W52638713, tmp8787)


W52648714 := tmp8788
_ = W52648714

tmp8789 := Call(__e, PrimFunc(symshen_4update_1history))


W52658715 := tmp8789
_ = W52658715

tmp8790 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dtc_d)
}
__typedArg0 := symshen_4_dtc_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4evaluate_1lineread), W52648714, W52658715, tmp8790)
return


}, 0)

tmp8791 := Call(__e, ns2_1set, symshen_4read_1evaluate_1print, tmp8784)


_ = tmp8791

tmp8792 := MakeNative(func(__e *ControlFlow) {
V5266 := __e.Get(1)
_ = V5266
V5267 := __e.Get(2)
_ = V5267
tmp8797 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symnull, V5266)
}
__typedArg0 := symnull
__typedArg1 := V5266
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp8797 {
__e.Return(V5267)
return
} else {
tmp8793 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(V5266)
}
__typedArg0 := V5266
return Call(__e, PrimFunc(symstr), __typedArg0)
})()

W52688716 := tmp8793
_ = W52688716

tmp8794 := Call(__e, PrimFunc(symexternal), V5266)


W52698717 := tmp8794
_ = W52698717

tmp8795 := MakeNative(func(__e *ControlFlow) {
Z5270 := __e.Get(1)
_ = Z5270
__e.TailApply(PrimFunc(symshen_4pui_1h), W52688716, W52698717, Z5270)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp8795, V5267)
return


}


}, 2)

tmp8798 := Call(__e, ns2_1set, symshen_4package_1user_1input, tmp8792)


_ = tmp8798

tmp8799 := MakeNative(func(__e *ControlFlow) {
V5275 := __e.Get(1)
_ = V5275
V5276 := __e.Get(2)
_ = V5276
V5277 := __e.Get(3)
_ = V5277
tmp8840 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5277)
}
__typedArg0 := V5277
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8827 Obj

if True == tmp8840 {
tmp8838 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5277)
}
__typedArg0 := V5277
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8839 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symfn, tmp8838)
}
__typedArg0 := symfn
__typedArg1 := tmp8838
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8829 Obj

if True == tmp8839 {
tmp8836 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5277)
}
__typedArg0 := V5277
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8837 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp8836)
}
__typedArg0 := tmp8836
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8831 Obj

if True == tmp8837 {
tmp8833 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5277)
}
__typedArg0 := V5277
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8834 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp8833)
}
__typedArg0 := tmp8833
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8835 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp8834)
}
__typedArg0 := Nil
__typedArg1 := tmp8834
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8832 Obj

if True == tmp8835 {
ifres8832 = True


} else {
ifres8832 = False


}

ifres8831 = ifres8832


} else {
ifres8831 = False


}

var ifres8830 Obj

if True == ifres8831 {
ifres8830 = True


} else {
ifres8830 = False


}

ifres8829 = ifres8830


} else {
ifres8829 = False


}

var ifres8828 Obj

if True == ifres8829 {
ifres8828 = True


} else {
ifres8828 = False


}

ifres8827 = ifres8828


} else {
ifres8827 = False


}

if True == ifres8827 {
tmp8805 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5277)
}
__typedArg0 := V5277
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8806 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp8805)
}
__typedArg0 := tmp8805
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8807 := Call(__e, PrimFunc(symshen_4internal_2), tmp8806, V5275, V5276)


if True == tmp8807 {
tmp8800 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5277)
}
__typedArg0 := V5277
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8801 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp8800)
}
__typedArg0 := tmp8800
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8802 := Call(__e, PrimFunc(symshen_4intern_1in_1package), V5275, tmp8801)


tmp8803 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8802, Nil)
}
__typedArg0 := tmp8802
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfn, tmp8803)
}
__typedArg0 := symfn
__typedArg1 := tmp8803
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return(V5277)
return
}


} else {
tmp8825 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5277)
}
__typedArg0 := V5277
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp8825 {
tmp8822 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5277)
}
__typedArg0 := V5277
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8823 := Call(__e, PrimFunc(symshen_4internal_2), tmp8822, V5275, V5276)


if True == tmp8823 {
tmp8808 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5277)
}
__typedArg0 := V5277
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8809 := Call(__e, PrimFunc(symshen_4intern_1in_1package), V5275, tmp8808)


tmp8810 := MakeNative(func(__e *ControlFlow) {
Z5278 := __e.Get(1)
_ = Z5278
__e.TailApply(PrimFunc(symshen_4pui_1h), V5275, V5276, Z5278)
return
}, 1)

tmp8811 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5277)
}
__typedArg0 := V5277
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8812 := Call(__e, PrimFunc(symmap), tmp8810, tmp8811)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8809, tmp8812)
}
__typedArg0 := tmp8809
__typedArg1 := tmp8812
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp8819 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5277)
}
__typedArg0 := V5277
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8820 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp8819)
}
__typedArg0 := tmp8819
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp8820 {
tmp8813 := MakeNative(func(__e *ControlFlow) {
Z5279 := __e.Get(1)
_ = Z5279
__e.TailApply(PrimFunc(symshen_4pui_1h), V5275, V5276, Z5279)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp8813, V5277)
return


} else {
tmp8814 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5277)
}
__typedArg0 := V5277
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8815 := MakeNative(func(__e *ControlFlow) {
Z5280 := __e.Get(1)
_ = Z5280
__e.TailApply(PrimFunc(symshen_4pui_1h), V5275, V5276, Z5280)
return
}, 1)

tmp8816 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5277)
}
__typedArg0 := V5277
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8817 := Call(__e, PrimFunc(symmap), tmp8815, tmp8816)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8814, tmp8817)
}
__typedArg0 := tmp8814
__typedArg1 := tmp8817
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

tmp8841 := Call(__e, ns2_1set, symshen_4pui_1h, tmp8799)


_ = tmp8841

tmp8842 := MakeNative(func(__e *ControlFlow) {
tmp8843 := Call(__e, PrimFunc(symit))


tmp8844 := Call(__e, PrimFunc(symshen_4trim_1it), tmp8843)


tmp8845 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dhistory_d)
}
__typedArg0 := symshen_4_dhistory_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp8846 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8844, tmp8845)
}
__typedArg0 := tmp8844
__typedArg1 := tmp8845
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dhistory_d, tmp8846)
}
__typedArg0 := symshen_4_dhistory_d
__typedArg1 := tmp8846
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return


}, 0)

tmp8847 := Call(__e, ns2_1set, symshen_4update_1history, tmp8842)


_ = tmp8847

tmp8848 := MakeNative(func(__e *ControlFlow) {
V5281 := __e.Get(1)
_ = V5281
tmp8856 := Call(__e, PrimFunc(symshen_4_7string_2), V5281)


var ifres8851 Obj

if True == tmp8856 {
tmp8853 := Call(__e, PrimFunc(symhdstr), V5281)


tmp8854 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_1_6n) {
return PrimStringToNumber(tmp8853)
}
__typedArg0 := tmp8853
return Call(__e, PrimFunc(symstring_1_6n), __typedArg0)
})()

tmp8855 := Call(__e, PrimFunc(symshen_4whitespace_2), tmp8854)


var ifres8852 Obj

if True == tmp8855 {
ifres8852 = True


} else {
ifres8852 = False


}

ifres8851 = ifres8852


} else {
ifres8851 = False


}

if True == ifres8851 {
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

tmp8857 := Call(__e, ns2_1set, symshen_4trim_1it, tmp8848)


_ = tmp8857

tmp8858 := MakeNative(func(__e *ControlFlow) {
V5300 := __e.Get(1)
_ = V5300
V5301 := __e.Get(2)
_ = V5301
V5302 := __e.Get(3)
_ = V5302
tmp8977 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5300)
}
__typedArg0 := V5300
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8946 Obj

if True == tmp8977 {
tmp8975 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5300)
}
__typedArg0 := V5300
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8976 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp8975)
}
__typedArg0 := Nil
__typedArg1 := tmp8975
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8948 Obj

if True == tmp8976 {
tmp8974 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5301)
}
__typedArg0 := V5301
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8950 Obj

if True == tmp8974 {
tmp8972 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5301)
}
__typedArg0 := V5301
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8973 := Call(__e, PrimFunc(symshen_4_7string_2), tmp8972)


var ifres8952 Obj

if True == tmp8973 {
tmp8969 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5301)
}
__typedArg0 := V5301
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8970 := Call(__e, PrimFunc(symhdstr), tmp8969)


tmp8971 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("!"), tmp8970)
}
__typedArg0 := MakeString("!")
__typedArg1 := tmp8970
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8954 Obj

if True == tmp8971 {
tmp8966 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5301)
}
__typedArg0 := V5301
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8968 := Call(__e, PrimFunc(symshen_4_7string_2), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(tmp8966)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := tmp8966
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


var ifres8956 Obj

if True == tmp8968 {
tmp8962 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5301)
}
__typedArg0 := V5301
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8964 := Call(__e, PrimFunc(symhdstr), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(tmp8962)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := tmp8962
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


tmp8965 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("!"), tmp8964)
}
__typedArg0 := MakeString("!")
__typedArg1 := tmp8964
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8958 Obj

if True == tmp8965 {
tmp8960 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5301)
}
__typedArg0 := V5301
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8961 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp8960)
}
__typedArg0 := tmp8960
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8959 Obj

if True == tmp8961 {
ifres8959 = True


} else {
ifres8959 = False


}

ifres8958 = ifres8959


} else {
ifres8958 = False


}

var ifres8957 Obj

if True == ifres8958 {
ifres8957 = True


} else {
ifres8957 = False


}

ifres8956 = ifres8957


} else {
ifres8956 = False


}

var ifres8955 Obj

if True == ifres8956 {
ifres8955 = True


} else {
ifres8955 = False


}

ifres8954 = ifres8955


} else {
ifres8954 = False


}

var ifres8953 Obj

if True == ifres8954 {
ifres8953 = True


} else {
ifres8953 = False


}

ifres8952 = ifres8953


} else {
ifres8952 = False


}

var ifres8951 Obj

if True == ifres8952 {
ifres8951 = True


} else {
ifres8951 = False


}

ifres8950 = ifres8951


} else {
ifres8950 = False


}

var ifres8949 Obj

if True == ifres8950 {
ifres8949 = True


} else {
ifres8949 = False


}

ifres8948 = ifres8949


} else {
ifres8948 = False


}

var ifres8947 Obj

if True == ifres8948 {
ifres8947 = True


} else {
ifres8947 = False


}

ifres8946 = ifres8947


} else {
ifres8946 = False


}

if True == ifres8946 {
tmp8859 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5301)
}
__typedArg0 := V5301
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8860 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp8859)
}
__typedArg0 := tmp8859
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8861 := Call(__e, PrimFunc(symread_1from_1string), tmp8860)


W53038718 := tmp8861
_ = W53038718

tmp8862 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5301)
}
__typedArg0 := V5301
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8863 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp8862)
}
__typedArg0 := tmp8862
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8864 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5301)
}
__typedArg0 := V5301
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8865 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8863, tmp8864)
}
__typedArg0 := tmp8863
__typedArg1 := tmp8864
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8866 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dhistory_d, tmp8865)
}
__typedArg0 := symshen_4_dhistory_d
__typedArg1 := tmp8865
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

W53048719 := tmp8866
_ = W53048719

tmp8867 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5301)
}
__typedArg0 := V5301
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8868 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp8867)
}
__typedArg0 := tmp8867
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8869 := Call(__e, PrimFunc(symshen_4app), tmp8868, MakeString("\n"), symshen_4a)


tmp8870 := Call(__e, PrimFunc(symstoutput))


tmp8871 := Call(__e, PrimFunc(sympr), tmp8869, tmp8870)


W53058720 := tmp8871
_ = W53058720

__e.TailApply(PrimFunc(symshen_4evaluate_1lineread), W53038718, W53048719, V5302)
return


} else {
tmp8944 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5300)
}
__typedArg0 := V5300
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8928 Obj

if True == tmp8944 {
tmp8942 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5300)
}
__typedArg0 := V5300
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8943 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp8942)
}
__typedArg0 := Nil
__typedArg1 := tmp8942
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8930 Obj

if True == tmp8943 {
tmp8941 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5301)
}
__typedArg0 := V5301
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8932 Obj

if True == tmp8941 {
tmp8939 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5301)
}
__typedArg0 := V5301
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8940 := Call(__e, PrimFunc(symshen_4_7string_2), tmp8939)


var ifres8934 Obj

if True == tmp8940 {
tmp8936 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5301)
}
__typedArg0 := V5301
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8937 := Call(__e, PrimFunc(symhdstr), tmp8936)


tmp8938 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("!"), tmp8937)
}
__typedArg0 := MakeString("!")
__typedArg1 := tmp8937
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8935 Obj

if True == tmp8938 {
ifres8935 = True


} else {
ifres8935 = False


}

ifres8934 = ifres8935


} else {
ifres8934 = False


}

var ifres8933 Obj

if True == ifres8934 {
ifres8933 = True


} else {
ifres8933 = False


}

ifres8932 = ifres8933


} else {
ifres8932 = False


}

var ifres8931 Obj

if True == ifres8932 {
ifres8931 = True


} else {
ifres8931 = False


}

ifres8930 = ifres8931


} else {
ifres8930 = False


}

var ifres8929 Obj

if True == ifres8930 {
ifres8929 = True


} else {
ifres8929 = False


}

ifres8928 = ifres8929


} else {
ifres8928 = False


}

if True == ifres8928 {
tmp8877 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5301)
}
__typedArg0 := V5301
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8879 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(tmp8877)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := tmp8877
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})(), MakeString(""))
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(tmp8877)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := tmp8877
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()
__typedArg1 := MakeString("")
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8872 Obj

if True == tmp8879 {
ifres8872 = Nil


} else {
tmp8873 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5301)
}
__typedArg0 := V5301
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8875 := Call(__e, PrimFunc(symread_1from_1string), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(tmp8873)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := tmp8873
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


tmp8876 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp8875)
}
__typedArg0 := tmp8875
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

ifres8872 = tmp8876


}

W53068721 := ifres8872
_ = W53068721

tmp8880 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5301)
}
__typedArg0 := V5301
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8881 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(tmp8880)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := tmp8880
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()

tmp8882 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5301)
}
__typedArg0 := V5301
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8883 := Call(__e, PrimFunc(symshen_4use_1history), W53068721, tmp8881, tmp8882)


W53078722 := tmp8883
_ = W53078722

tmp8884 := Call(__e, PrimFunc(symshen_4app), W53078722, MakeString("\n"), symshen_4a)


tmp8885 := Call(__e, PrimFunc(symstoutput))


tmp8886 := Call(__e, PrimFunc(sympr), tmp8884, tmp8885)


W53088723 := tmp8886
_ = W53088723

tmp8887 := Call(__e, PrimFunc(symread_1from_1string), W53078722)


W53098724 := tmp8887
_ = W53098724

tmp8888 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5301)
}
__typedArg0 := V5301
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8889 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W53078722, tmp8888)
}
__typedArg0 := W53078722
__typedArg1 := tmp8888
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8890 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dhistory_d, tmp8889)
}
__typedArg0 := symshen_4_dhistory_d
__typedArg1 := tmp8889
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

W53108725 := tmp8890
_ = W53108725

__e.TailApply(PrimFunc(symshen_4evaluate_1lineread), W53098724, W53108725, V5302)
return


} else {
tmp8926 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5300)
}
__typedArg0 := V5300
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8910 Obj

if True == tmp8926 {
tmp8924 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5300)
}
__typedArg0 := V5300
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8925 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp8924)
}
__typedArg0 := Nil
__typedArg1 := tmp8924
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8912 Obj

if True == tmp8925 {
tmp8923 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5301)
}
__typedArg0 := V5301
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8914 Obj

if True == tmp8923 {
tmp8921 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5301)
}
__typedArg0 := V5301
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8922 := Call(__e, PrimFunc(symshen_4_7string_2), tmp8921)


var ifres8916 Obj

if True == tmp8922 {
tmp8918 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5301)
}
__typedArg0 := V5301
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8919 := Call(__e, PrimFunc(symhdstr), tmp8918)


tmp8920 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("%"), tmp8919)
}
__typedArg0 := MakeString("%")
__typedArg1 := tmp8919
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8917 Obj

if True == tmp8920 {
ifres8917 = True


} else {
ifres8917 = False


}

ifres8916 = ifres8917


} else {
ifres8916 = False


}

var ifres8915 Obj

if True == ifres8916 {
ifres8915 = True


} else {
ifres8915 = False


}

ifres8914 = ifres8915


} else {
ifres8914 = False


}

var ifres8913 Obj

if True == ifres8914 {
ifres8913 = True


} else {
ifres8913 = False


}

ifres8912 = ifres8913


} else {
ifres8912 = False


}

var ifres8911 Obj

if True == ifres8912 {
ifres8911 = True


} else {
ifres8911 = False


}

ifres8910 = ifres8911


} else {
ifres8910 = False


}

if True == ifres8910 {
tmp8896 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5301)
}
__typedArg0 := V5301
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8898 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(tmp8896)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := tmp8896
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})(), MakeString(""))
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(tmp8896)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := tmp8896
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()
__typedArg1 := MakeString("")
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8891 Obj

if True == tmp8898 {
ifres8891 = Nil


} else {
tmp8892 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5301)
}
__typedArg0 := V5301
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8894 := Call(__e, PrimFunc(symread_1from_1string), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(tmp8892)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := tmp8892
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


tmp8895 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp8894)
}
__typedArg0 := tmp8894
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

ifres8891 = tmp8895


}

W53118726 := ifres8891
_ = W53118726

tmp8899 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5301)
}
__typedArg0 := V5301
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8900 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(tmp8899)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := tmp8899
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()

tmp8901 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5301)
}
__typedArg0 := V5301
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8902 := Call(__e, PrimFunc(symshen_4peek_1history), W53118726, tmp8900, tmp8901)


W53128727 := tmp8902
_ = W53128727

tmp8903 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5301)
}
__typedArg0 := V5301
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8904 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dhistory_d, tmp8903)
}
__typedArg0 := symshen_4_dhistory_d
__typedArg1 := tmp8903
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

W53138728 := tmp8904
_ = W53138728

__e.TailApply(PrimFunc(symabort))
return


} else {
tmp8908 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(True, V5302)
}
__typedArg0 := True
__typedArg1 := V5302
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp8908 {
__e.TailApply(PrimFunc(symshen_4check_1eval_1and_1print), V5300)
return
} else {
tmp8906 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(False, V5302)
}
__typedArg0 := False
__typedArg1 := V5302
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp8906 {
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

tmp8978 := Call(__e, ns2_1set, symshen_4evaluate_1lineread, tmp8858)


_ = tmp8978

tmp8979 := MakeNative(func(__e *ControlFlow) {
V5314 := __e.Get(1)
_ = V5314
V5315 := __e.Get(2)
_ = V5315
V5316 := __e.Get(3)
_ = V5316
tmp8985 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(syminteger_2) {
return PrimIsInteger(V5314)
}
__typedArg0 := V5314
return Call(__e, PrimFunc(syminteger_2), __typedArg0)
})()

if True == tmp8985 {
tmp8980 := (func() Obj {
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

tmp8981 := Call(__e, PrimFunc(symreverse), V5316)


__e.TailApply(PrimFunc(symnth), tmp8980, tmp8981)
return


} else {
tmp8983 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsymbol_2) {
return PrimIsSymbol(V5314)
}
__typedArg0 := V5314
return Call(__e, PrimFunc(symsymbol_2), __typedArg0)
})()

if True == tmp8983 {
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

tmp8986 := Call(__e, ns2_1set, symshen_4use_1history, tmp8979)


_ = tmp8986

tmp8987 := MakeNative(func(__e *ControlFlow) {
V5317 := __e.Get(1)
_ = V5317
V5318 := __e.Get(2)
_ = V5318
V5319 := __e.Get(3)
_ = V5319
tmp9001 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(syminteger_2) {
return PrimIsInteger(V5317)
}
__typedArg0 := V5317
return Call(__e, PrimFunc(syminteger_2), __typedArg0)
})()

if True == tmp9001 {
tmp8988 := (func() Obj {
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

tmp8989 := Call(__e, PrimFunc(symreverse), V5319)


tmp8990 := Call(__e, PrimFunc(symnth), tmp8988, tmp8989)


tmp8991 := Call(__e, PrimFunc(symshen_4app), tmp8990, MakeString(""), symshen_4a)


tmp8992 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("\n"))
__typedS1, __typedOK1 := TypedString(tmp8991)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("\n")
__typedArg1 := tmp8991
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp8993 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp8992, tmp8993)
return


} else {
tmp8999 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V5318, MakeString(""))
}
__typedArg0 := V5318
__typedArg1 := MakeString("")
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8996 Obj

if True == tmp8999 {
ifres8996 = True


} else {
tmp8998 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsymbol_2) {
return PrimIsSymbol(V5317)
}
__typedArg0 := V5317
return Call(__e, PrimFunc(symsymbol_2), __typedArg0)
})()

var ifres8997 Obj

if True == tmp8998 {
ifres8997 = True


} else {
ifres8997 = False


}

ifres8996 = ifres8997


}

if True == ifres8996 {
tmp8994 := Call(__e, PrimFunc(symreverse), V5319)


__e.TailApply(PrimFunc(symshen_4recursive_1string_1match), MakeNumber(0), V5318, tmp8994)
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

tmp9002 := Call(__e, ns2_1set, symshen_4peek_1history, tmp8987)


_ = tmp9002

tmp9003 := MakeNative(func(__e *ControlFlow) {
V5329 := __e.Get(1)
_ = V5329
V5330 := __e.Get(2)
_ = V5330
tmp9014 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V5330)
}
__typedArg0 := Nil
__typedArg1 := V5330
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9014 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("\ninput not found"))
}
__typedArg0 := MakeString("\ninput not found")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
} else {
tmp9012 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5330)
}
__typedArg0 := V5330
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9008 Obj

if True == tmp9012 {
tmp9010 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5330)
}
__typedArg0 := V5330
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9011 := Call(__e, PrimFunc(symshen_4string_1prefix_2), V5329, tmp9010)


var ifres9009 Obj

if True == tmp9011 {
ifres9009 = True


} else {
ifres9009 = False


}

ifres9008 = ifres9009


} else {
ifres9008 = False


}

if True == ifres9008 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5330)
}
__typedArg0 := V5330
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return
} else {
tmp9006 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5330)
}
__typedArg0 := V5330
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9006 {
tmp9004 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5330)
}
__typedArg0 := V5330
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4string_1match), V5329, tmp9004)
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

tmp9015 := Call(__e, ns2_1set, symshen_4string_1match, tmp9003)


_ = tmp9015

tmp9016 := MakeNative(func(__e *ControlFlow) {
V5338 := __e.Get(1)
_ = V5338
V5339 := __e.Get(2)
_ = V5339
tmp9053 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString(""), V5338)
}
__typedArg0 := MakeString("")
__typedArg1 := V5338
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9053 {
__e.Return(True)
return
} else {
tmp9051 := Call(__e, PrimFunc(symshen_4_7string_2), V5338)


var ifres9046 Obj

if True == tmp9051 {
tmp9048 := Call(__e, PrimFunc(symhdstr), V5338)


tmp9049 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_1_6n) {
return PrimStringToNumber(tmp9048)
}
__typedArg0 := tmp9048
return Call(__e, PrimFunc(symstring_1_6n), __typedArg0)
})()

tmp9050 := Call(__e, PrimFunc(symshen_4whitespace_2), tmp9049)


var ifres9047 Obj

if True == tmp9050 {
ifres9047 = True


} else {
ifres9047 = False


}

ifres9046 = ifres9047


} else {
ifres9046 = False


}

if True == ifres9046 {
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
tmp9044 := Call(__e, PrimFunc(symshen_4_7string_2), V5339)


var ifres9039 Obj

if True == tmp9044 {
tmp9041 := Call(__e, PrimFunc(symhdstr), V5339)


tmp9042 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_1_6n) {
return PrimStringToNumber(tmp9041)
}
__typedArg0 := tmp9041
return Call(__e, PrimFunc(symstring_1_6n), __typedArg0)
})()

tmp9043 := Call(__e, PrimFunc(symshen_4whitespace_2), tmp9042)


var ifres9040 Obj

if True == tmp9043 {
ifres9040 = True


} else {
ifres9040 = False


}

ifres9039 = ifres9040


} else {
ifres9039 = False


}

if True == ifres9039 {
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
tmp9037 := Call(__e, PrimFunc(symshen_4_7string_2), V5339)


var ifres9033 Obj

if True == tmp9037 {
tmp9035 := Call(__e, PrimFunc(symhdstr), V5339)


tmp9036 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("("), tmp9035)
}
__typedArg0 := MakeString("(")
__typedArg1 := tmp9035
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9034 Obj

if True == tmp9036 {
ifres9034 = True


} else {
ifres9034 = False


}

ifres9033 = ifres9034


} else {
ifres9033 = False


}

if True == ifres9033 {
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
tmp9031 := Call(__e, PrimFunc(symshen_4_7string_2), V5338)


var ifres9023 Obj

if True == tmp9031 {
tmp9030 := Call(__e, PrimFunc(symshen_4_7string_2), V5339)


var ifres9025 Obj

if True == tmp9030 {
tmp9027 := Call(__e, PrimFunc(symhdstr), V5338)


tmp9028 := Call(__e, PrimFunc(symhdstr), V5339)


tmp9029 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp9027, tmp9028)
}
__typedArg0 := tmp9027
__typedArg1 := tmp9028
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9026 Obj

if True == tmp9029 {
ifres9026 = True


} else {
ifres9026 = False


}

ifres9025 = ifres9026


} else {
ifres9025 = False


}

var ifres9024 Obj

if True == ifres9025 {
ifres9024 = True


} else {
ifres9024 = False


}

ifres9023 = ifres9024


} else {
ifres9023 = False


}

if True == ifres9023 {
tmp9020 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V5338)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V5338
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4string_1prefix_2), tmp9020, (func() Obj {
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

tmp9054 := Call(__e, ns2_1set, symshen_4string_1prefix_2, tmp9016)


_ = tmp9054

tmp9055 := MakeNative(func(__e *ControlFlow) {
V5350 := __e.Get(1)
_ = V5350
V5351 := __e.Get(2)
_ = V5351
V5352 := __e.Get(3)
_ = V5352
tmp9070 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V5352)
}
__typedArg0 := Nil
__typedArg1 := V5352
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9070 {
__e.Return(symshen_4skip)
return
} else {
tmp9068 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5352)
}
__typedArg0 := V5352
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9068 {
tmp9063 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5352)
}
__typedArg0 := V5352
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9064 := Call(__e, PrimFunc(symshen_4string_1prefix_2), V5351, tmp9063)


var ifres9056 Obj

if True == tmp9064 {
tmp9057 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5352)
}
__typedArg0 := V5352
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9058 := Call(__e, PrimFunc(symshen_4app), tmp9057, MakeString("\n"), symshen_4a)


tmp9060 := Call(__e, PrimFunc(symshen_4app), V5350, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(". "))
__typedS1, __typedOK1 := TypedString(tmp9058)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(". ")
__typedArg1 := tmp9058
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


tmp9061 := Call(__e, PrimFunc(symstoutput))


tmp9062 := Call(__e, PrimFunc(sympr), tmp9060, tmp9061)


ifres9056 = tmp9062


} else {
ifres9056 = symshen_4skip


}

_ = ifres9056

tmp9065 := (func() Obj {
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

tmp9066 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5352)
}
__typedArg0 := V5352
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4recursive_1string_1match), tmp9065, V5351, tmp9066)
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

__e.TailApply(ns2_1set, symshen_4recursive_1string_1match, tmp9055)
return




}, 0)

