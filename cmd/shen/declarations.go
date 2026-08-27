package main

import . "github.com/pyrex41/shen-go/kl"

var DeclarationsMain = MakeNative(func(__e *ControlFlow) {
tmp7706 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dhistory_d, Nil)
}
__typedArg0 := symshen_4_dhistory_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp7706

tmp7707 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dtc_d, False)
}
__typedArg0 := symshen_4_dtc_d
__typedArg1 := False
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp7707

tmp7708 := Call(__e, PrimFunc(symvector), MakeNumber(20000))


tmp7709 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(sym_dproperty_1vector_d, tmp7708)
}
__typedArg0 := sym_dproperty_1vector_d
__typedArg1 := tmp7708
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp7709

tmp7710 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
__e.TailApply(PrimFunc(symshen_4macros), X)
return
}, 1)

tmp7711 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4macros, tmp7710)
}
__typedArg0 := symshen_4macros
__typedArg1 := tmp7710
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7712 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7711, Nil)
}
__typedArg0 := tmp7711
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7713 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(sym_dmacros_d, tmp7712)
}
__typedArg0 := sym_dmacros_d
__typedArg1 := tmp7712
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp7713

tmp7714 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dgensym_d, MakeNumber(0))
}
__typedArg0 := symshen_4_dgensym_d
__typedArg1 := MakeNumber(0)
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp7714

tmp7715 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dtracking_d, Nil)
}
__typedArg0 := symshen_4_dtracking_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp7715

tmp7716 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dprofiled_d, Nil)
}
__typedArg0 := symshen_4_dprofiled_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp7716

tmp7717 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(sym_dhome_1directory_d, MakeString(""))
}
__typedArg0 := sym_dhome_1directory_d
__typedArg1 := MakeString("")
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp7717

tmp7718 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtype, Nil)
}
__typedArg0 := symtype
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7719 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4input_1h_7, tmp7718)
}
__typedArg0 := symshen_4input_1h_7
__typedArg1 := tmp7718
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7720 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symopen, tmp7719)
}
__typedArg0 := symopen
__typedArg1 := tmp7719
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7721 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symset, tmp7720)
}
__typedArg0 := symset
__typedArg1 := tmp7720
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7722 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhere, tmp7721)
}
__typedArg0 := symwhere
__typedArg1 := tmp7721
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7723 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp7722)
}
__typedArg0 := symlet
__typedArg1 := tmp7722
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7724 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp7723)
}
__typedArg0 := symlambda
__typedArg1 := tmp7723
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7725 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp7724)
}
__typedArg0 := symcons
__typedArg1 := tmp7724
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7726 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8v, tmp7725)
}
__typedArg0 := sym_8v
__typedArg1 := tmp7725
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7727 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8s, tmp7726)
}
__typedArg0 := sym_8s
__typedArg1 := tmp7726
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7728 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8p, tmp7727)
}
__typedArg0 := sym_8p
__typedArg1 := tmp7727
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7729 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dspecial_d, tmp7728)
}
__typedArg0 := symshen_4_dspecial_d
__typedArg1 := tmp7728
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp7729

tmp7730 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dextraspecial_d, Nil)
}
__typedArg0 := symshen_4_dextraspecial_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp7730

tmp7731 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dspy_d, False)
}
__typedArg0 := symshen_4_dspy_d
__typedArg1 := False
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp7731

tmp7732 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_ddatatypes_d, Nil)
}
__typedArg0 := symshen_4_ddatatypes_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp7732

tmp7733 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dalldatatypes_d, Nil)
}
__typedArg0 := symshen_4_dalldatatypes_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp7733

tmp7734 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dshen_1type_1theory_1enabled_2_d, True)
}
__typedArg0 := symshen_4_dshen_1type_1theory_1enabled_2_d
__typedArg1 := True
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp7734

tmp7735 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dpackage_d, symnull)
}
__typedArg0 := symshen_4_dpackage_d
__typedArg1 := symnull
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp7735

tmp7736 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsynonyms_d, Nil)
}
__typedArg0 := symshen_4_dsynonyms_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp7736

tmp7737 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsystem_d, Nil)
}
__typedArg0 := symshen_4_dsystem_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp7737

tmp7738 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, Nil)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp7738

tmp7739 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_doccurs_d, True)
}
__typedArg0 := symshen_4_doccurs_d
__typedArg1 := True
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp7739

tmp7740 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dfactorise_2_d, False)
}
__typedArg0 := symshen_4_dfactorise_2_d
__typedArg1 := False
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp7740

tmp7741 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dmaxinferences_d, MakeNumber(1e+06))
}
__typedArg0 := symshen_4_dmaxinferences_d
__typedArg1 := MakeNumber(1e+06)
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp7741

tmp7742 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(sym_dmaximum_1print_1sequence_1size_d, MakeNumber(20))
}
__typedArg0 := sym_dmaximum_1print_1sequence_1size_d
__typedArg1 := MakeNumber(20)
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp7742

tmp7743 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dcall_d, MakeNumber(0))
}
__typedArg0 := symshen_4_dcall_d
__typedArg1 := MakeNumber(0)
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp7743

tmp7744 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dinfs_d, MakeNumber(0))
}
__typedArg0 := symshen_4_dinfs_d
__typedArg1 := MakeNumber(0)
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp7744

tmp7745 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(sym_dhush_d, False)
}
__typedArg0 := sym_dhush_d
__typedArg1 := False
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp7745

tmp7746 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_doptimise_d, False)
}
__typedArg0 := symshen_4_doptimise_d
__typedArg1 := False
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp7746

tmp7747 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(sym_dversion_d, MakeString("42"))
}
__typedArg0 := sym_dversion_d
__typedArg1 := MakeString("42")
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp7747

tmp7748 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dnames_d, Nil)
}
__typedArg0 := symshen_4_dnames_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp7748

tmp7749 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dstep_d, False)
}
__typedArg0 := symshen_4_dstep_d
__typedArg1 := False
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp7749

tmp7750 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dit_d, MakeString(""))
}
__typedArg0 := symshen_4_dit_d
__typedArg1 := MakeString("")
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp7750

tmp7751 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dresidue_d, Nil)
}
__typedArg0 := symshen_4_dresidue_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp7751

tmp7752 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dprolog_1memory_d, MakeNumber(1000))
}
__typedArg0 := symshen_4_dprolog_1memory_d
__typedArg1 := MakeNumber(1000)
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp7752

tmp7753 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dloading_2_d, False)
}
__typedArg0 := symshen_4_dloading_2_d
__typedArg1 := False
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp7753

tmp7754 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_duserdefs_d, Nil)
}
__typedArg0 := symshen_4_duserdefs_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp7754

tmp7755 := MakeNative(func(__e *ControlFlow) {
tmp7756 := MakeNative(func(__e *ControlFlow) {
Z5759 := __e.Get(1)
_ = Z5759
__e.TailApply(PrimFunc(symshen_4typename), Z5759)
return
}, 1)

tmp7757 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dalldatatypes_d)
}
__typedArg0 := symshen_4_dalldatatypes_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symmap), tmp7756, tmp7757)
return


}, 0)

tmp7758 := Call(__e, ns2_1set, symdatatypes, tmp7755)


_ = tmp7758

tmp7759 := MakeNative(func(__e *ControlFlow) {
tmp7760 := MakeNative(func(__e *ControlFlow) {
Z5760 := __e.Get(1)
_ = Z5760
__e.TailApply(PrimFunc(symshen_4typename), Z5760)
return
}, 1)

tmp7761 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_ddatatypes_d)
}
__typedArg0 := symshen_4_ddatatypes_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symmap), tmp7760, tmp7761)
return


}, 0)

tmp7762 := Call(__e, ns2_1set, symincluded, tmp7759)


_ = tmp7762

tmp7763 := MakeNative(func(__e *ControlFlow) {
V5763 := __e.Get(1)
_ = V5763
tmp7768 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5763)
}
__typedArg0 := V5763
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp7768 {
tmp7764 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5763)
}
__typedArg0 := V5763
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7765 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(tmp7764)
}
__typedArg0 := tmp7764
return Call(__e, PrimFunc(symstr), __typedArg0)
})()

tmp7766 := Call(__e, PrimFunc(symshen_4typename_1h), tmp7765)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(tmp7766)
}
__typedArg0 := tmp7766
return Call(__e, PrimFunc(symintern), __typedArg0)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("partial function shen.typename"))
}
__typedArg0 := MakeString("partial function shen.typename")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}, 1)

tmp7769 := Call(__e, ns2_1set, symshen_4typename, tmp7763)


_ = tmp7769

tmp7770 := MakeNative(func(__e *ControlFlow) {
V5764 := __e.Get(1)
_ = V5764
tmp7777 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("#type"), V5764)
}
__typedArg0 := MakeString("#type")
__typedArg1 := V5764
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp7777 {
__e.Return(MakeString(""))
return
} else {
tmp7775 := Call(__e, PrimFunc(symshen_4_7string_2), V5764)


if True == tmp7775 {
tmp7771 := Call(__e, PrimFunc(symhdstr), V5764)


tmp7773 := Call(__e, PrimFunc(symshen_4typename_1h), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V5764)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V5764
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(tmp7771)
__typedS1, __typedOK1 := TypedString(tmp7773)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := tmp7771
__typedArg1 := tmp7773
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("partial function shen.typename-h"))
}
__typedArg0 := MakeString("partial function shen.typename-h")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 1)

tmp7778 := Call(__e, ns2_1set, symshen_4typename_1h, tmp7770)


_ = tmp7778

tmp7779 := MakeNative(func(__e *ControlFlow) {
V5765 := __e.Get(1)
_ = V5765
if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5) {
__typedN0, __typedOK0 := TypedFloat64(V5765)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(0))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_5) {
return TypedMaterializeBoolean((__typedN0 < __typedN1))
}}
__typedArg0 := V5765
__typedArg1 := MakeNumber(0)
return Call(__e, PrimFunc(sym_5), __typedArg0, __typedArg1)
})() {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dprolog_1memory_d)
}
__typedArg0 := symshen_4_dprolog_1memory_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
} else {
tmp7781 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(syminteger_2) {
return PrimIsInteger(V5765)
}
__typedArg0 := V5765
return Call(__e, PrimFunc(syminteger_2), __typedArg0)
})()

if True == tmp7781 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dprolog_1memory_d, V5765)
}
__typedArg0 := symshen_4_dprolog_1memory_d
__typedArg1 := V5765
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return
} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("prolog memory expects an integer value\n"))
}
__typedArg0 := MakeString("prolog memory expects an integer value\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 1)

tmp7784 := Call(__e, ns2_1set, symprolog_1memory, tmp7779)


_ = tmp7784

tmp7785 := MakeNative(func(__e *ControlFlow) {
V5768 := __e.Get(1)
_ = V5768
tmp7801 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V5768)
}
__typedArg0 := Nil
__typedArg1 := V5768
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp7801 {
__e.Return(Nil)
return
} else {
tmp7799 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5768)
}
__typedArg0 := V5768
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7795 Obj

if True == tmp7799 {
tmp7797 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5768)
}
__typedArg0 := V5768
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7798 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7797)
}
__typedArg0 := tmp7797
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7796 Obj

if True == tmp7798 {
ifres7796 = True


} else {
ifres7796 = False


}

ifres7795 = ifres7796


} else {
ifres7795 = False


}

if True == ifres7795 {
tmp7786 := MakeNative(func(__e *ControlFlow) {
W5769 := __e.Get(1)
_ = W5769
tmp7787 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5768)
}
__typedArg0 := V5768
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7788 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7787)
}
__typedArg0 := tmp7787
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4initialise_1arity_1table), tmp7788)
return


}, 1)

tmp7789 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5768)
}
__typedArg0 := V5768
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7790 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5768)
}
__typedArg0 := V5768
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7791 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7790)
}
__typedArg0 := tmp7790
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7792 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp7793 := Call(__e, PrimFunc(symput), tmp7789, symarity, tmp7791, tmp7792)


__e.TailApply(tmp7786, tmp7793)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.initialise-arity-table"))
}
__typedArg0 := MakeString("implementation error in shen.initialise-arity-table")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 1)

tmp7802 := Call(__e, ns2_1set, symshen_4initialise_1lambda_1tables, tmp7785)


_ = tmp7802

tmp7803 := MakeNative(func(__e *ControlFlow) {
V5770 := __e.Get(1)
_ = V5770
tmp7804 := MakeNative(func(__e *ControlFlow) {
tmp7805 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symget), V5770, symarity, tmp7805)
return


}, 0)

tmp7806 := MakeNative(func(__e *ControlFlow) {
Z5771 := __e.Get(1)
_ = Z5771
__e.Return(MakeNumber(-1))
return
}, 1)

__e.TailApply(try_1catch, tmp7804, tmp7806)
return


}, 1)

tmp7807 := Call(__e, ns2_1set, symarity, tmp7803)


_ = tmp7807

tmp7808 := MakeNative(func(__e *ControlFlow) {
V5774 := __e.Get(1)
_ = V5774
tmp7824 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V5774)
}
__typedArg0 := Nil
__typedArg1 := V5774
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp7824 {
__e.Return(Nil)
return
} else {
tmp7822 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5774)
}
__typedArg0 := V5774
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7818 Obj

if True == tmp7822 {
tmp7820 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5774)
}
__typedArg0 := V5774
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7821 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7820)
}
__typedArg0 := tmp7820
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7819 Obj

if True == tmp7821 {
ifres7819 = True


} else {
ifres7819 = False


}

ifres7818 = ifres7819


} else {
ifres7818 = False


}

if True == ifres7818 {
tmp7809 := MakeNative(func(__e *ControlFlow) {
W5775 := __e.Get(1)
_ = W5775
tmp7810 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5774)
}
__typedArg0 := V5774
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7811 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7810)
}
__typedArg0 := tmp7810
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4initialise_1arity_1table), tmp7811)
return


}, 1)

tmp7812 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5774)
}
__typedArg0 := V5774
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7813 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5774)
}
__typedArg0 := V5774
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7814 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7813)
}
__typedArg0 := tmp7813
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7815 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp7816 := Call(__e, PrimFunc(symput), tmp7812, symarity, tmp7814, tmp7815)


__e.TailApply(tmp7809, tmp7816)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.initialise-arity-table"))
}
__typedArg0 := MakeString("implementation error in shen.initialise-arity-table")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 1)

tmp7825 := Call(__e, ns2_1set, symshen_4initialise_1arity_1table, tmp7808)


_ = tmp7825

tmp7826 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(2), Nil)
}
__typedArg0 := MakeNumber(2)
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7827 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8s, tmp7826)
}
__typedArg0 := sym_8s
__typedArg1 := tmp7826
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7828 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(2), tmp7827)
}
__typedArg0 := MakeNumber(2)
__typedArg1 := tmp7827
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7829 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8v, tmp7828)
}
__typedArg0 := sym_8v
__typedArg1 := tmp7828
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7830 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(2), tmp7829)
}
__typedArg0 := MakeNumber(2)
__typedArg1 := tmp7829
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7831 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8p, tmp7830)
}
__typedArg0 := sym_8p
__typedArg1 := tmp7830
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7832 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp7831)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp7831
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7833 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_b_6, tmp7832)
}
__typedArg0 := sym_5_b_6
__typedArg1 := tmp7832
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7834 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp7833)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp7833
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7835 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5end_6, tmp7834)
}
__typedArg0 := sym_5end_6
__typedArg1 := tmp7834
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7836 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp7835)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp7835
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7837 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5e_6, tmp7836)
}
__typedArg0 := sym_5e_6
__typedArg1 := tmp7836
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7838 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(2), tmp7837)
}
__typedArg0 := MakeNumber(2)
__typedArg1 := tmp7837
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7839 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_a_a, tmp7838)
}
__typedArg0 := sym_a_a
__typedArg1 := tmp7838
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7840 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(2), tmp7839)
}
__typedArg0 := MakeNumber(2)
__typedArg1 := tmp7839
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7841 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1, tmp7840)
}
__typedArg0 := sym_1
__typedArg1 := tmp7840
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7842 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(2), tmp7841)
}
__typedArg0 := MakeNumber(2)
__typedArg1 := tmp7841
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7843 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_c, tmp7842)
}
__typedArg0 := sym_c
__typedArg1 := tmp7842
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7844 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(2), tmp7843)
}
__typedArg0 := MakeNumber(2)
__typedArg1 := tmp7843
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7845 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_d, tmp7844)
}
__typedArg0 := sym_d
__typedArg1 := tmp7844
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7846 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(2), tmp7845)
}
__typedArg0 := MakeNumber(2)
__typedArg1 := tmp7845
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7847 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_7, tmp7846)
}
__typedArg0 := sym_7
__typedArg1 := tmp7846
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7848 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp7847)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp7847
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7849 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symy_1or_1n_2, tmp7848)
}
__typedArg0 := symy_1or_1n_2
__typedArg1 := tmp7848
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7850 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(2), tmp7849)
}
__typedArg0 := MakeNumber(2)
__typedArg1 := tmp7849
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7851 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwrite_1to_1file, tmp7850)
}
__typedArg0 := symwrite_1to_1file
__typedArg1 := tmp7850
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7852 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(2), tmp7851)
}
__typedArg0 := MakeNumber(2)
__typedArg1 := tmp7851
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7853 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwrite_1byte, tmp7852)
}
__typedArg0 := symwrite_1byte
__typedArg1 := tmp7852
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7854 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(5), tmp7853)
}
__typedArg0 := MakeNumber(5)
__typedArg1 := tmp7853
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7855 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhen, tmp7854)
}
__typedArg0 := symwhen
__typedArg1 := tmp7854
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7856 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(0), tmp7855)
}
__typedArg0 := MakeNumber(0)
__typedArg1 := tmp7855
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7857 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symversion, tmp7856)
}
__typedArg0 := symversion
__typedArg1 := tmp7856
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7858 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(5), tmp7857)
}
__typedArg0 := MakeNumber(5)
__typedArg1 := tmp7857
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7859 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvar_2, tmp7858)
}
__typedArg0 := symvar_2
__typedArg1 := tmp7858
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7860 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp7859)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp7859
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7861 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvariable_2, tmp7860)
}
__typedArg0 := symvariable_2
__typedArg1 := tmp7860
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7862 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp7861)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp7861
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7863 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvalue, tmp7862)
}
__typedArg0 := symvalue
__typedArg1 := tmp7862
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7864 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(3), tmp7863)
}
__typedArg0 := MakeNumber(3)
__typedArg1 := tmp7863
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7865 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector_1_6, tmp7864)
}
__typedArg0 := symvector_1_6
__typedArg1 := tmp7864
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7866 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp7865)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp7865
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7867 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector_2, tmp7866)
}
__typedArg0 := symvector_2
__typedArg1 := tmp7866
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7868 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp7867)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp7867
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7869 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp7868)
}
__typedArg0 := symvector
__typedArg1 := tmp7868
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7870 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(0), tmp7869)
}
__typedArg0 := MakeNumber(0)
__typedArg1 := tmp7869
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7871 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symuserdefs, tmp7870)
}
__typedArg0 := symuserdefs
__typedArg1 := tmp7870
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7872 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(2), tmp7871)
}
__typedArg0 := MakeNumber(2)
__typedArg1 := tmp7871
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7873 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symupdate_1lambda_1table, tmp7872)
}
__typedArg0 := symupdate_1lambda_1table
__typedArg1 := tmp7872
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7874 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp7873)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp7873
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7875 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symundefmacro, tmp7874)
}
__typedArg0 := symundefmacro
__typedArg1 := tmp7874
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7876 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp7875)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp7875
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7877 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symuntrack, tmp7876)
}
__typedArg0 := symuntrack
__typedArg1 := tmp7876
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7878 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(2), tmp7877)
}
__typedArg0 := MakeNumber(2)
__typedArg1 := tmp7877
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7879 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunion, tmp7878)
}
__typedArg0 := symunion
__typedArg1 := tmp7878
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7880 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp7879)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp7879
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7881 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunprofile, tmp7880)
}
__typedArg0 := symunprofile
__typedArg1 := tmp7880
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7882 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(3), tmp7881)
}
__typedArg0 := MakeNumber(3)
__typedArg1 := tmp7881
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7883 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunput, tmp7882)
}
__typedArg0 := symunput
__typedArg1 := tmp7882
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7884 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp7883)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp7883
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7885 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symundefmacro, tmp7884)
}
__typedArg0 := symundefmacro
__typedArg1 := tmp7884
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7886 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp7885)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp7885
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7887 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunabsolute, tmp7886)
}
__typedArg0 := symunabsolute
__typedArg1 := tmp7886
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7888 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(5), tmp7887)
}
__typedArg0 := MakeNumber(5)
__typedArg1 := tmp7887
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7889 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symreturn, tmp7888)
}
__typedArg0 := symreturn
__typedArg1 := tmp7888
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7890 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(2), tmp7889)
}
__typedArg0 := MakeNumber(2)
__typedArg1 := tmp7889
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7891 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtype, tmp7890)
}
__typedArg0 := symtype
__typedArg1 := tmp7890
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7892 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp7891)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp7891
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7893 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtuple_2, tmp7892)
}
__typedArg0 := symtuple_2
__typedArg1 := tmp7892
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7894 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(2), tmp7893)
}
__typedArg0 := MakeNumber(2)
__typedArg1 := tmp7893
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7895 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtrap_1error, tmp7894)
}
__typedArg0 := symtrap_1error
__typedArg1 := tmp7894
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7896 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(0), tmp7895)
}
__typedArg0 := MakeNumber(0)
__typedArg1 := tmp7895
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7897 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtracked, tmp7896)
}
__typedArg0 := symtracked
__typedArg1 := tmp7896
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7898 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp7897)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp7897
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7899 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtrack, tmp7898)
}
__typedArg0 := symtrack
__typedArg1 := tmp7898
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7900 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp7899)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp7899
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7901 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtlstr, tmp7900)
}
__typedArg0 := symtlstr
__typedArg1 := tmp7900
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7902 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp7901)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp7901
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7903 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symthaw, tmp7902)
}
__typedArg0 := symthaw
__typedArg1 := tmp7902
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7904 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(0), tmp7903)
}
__typedArg0 := MakeNumber(0)
__typedArg1 := tmp7903
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7905 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtc_2, tmp7904)
}
__typedArg0 := symtc_2
__typedArg1 := tmp7904
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7906 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp7905)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp7905
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7907 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtc, tmp7906)
}
__typedArg0 := symtc
__typedArg1 := tmp7906
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7908 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp7907)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp7907
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7909 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtl, tmp7908)
}
__typedArg0 := symtl
__typedArg1 := tmp7908
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7910 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp7909)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp7909
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7911 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtail, tmp7910)
}
__typedArg0 := symtail
__typedArg1 := tmp7910
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7912 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp7911)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp7911
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7913 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsystemf, tmp7912)
}
__typedArg0 := symsystemf
__typedArg1 := tmp7912
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7914 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp7913)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp7913
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7915 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol_2, tmp7914)
}
__typedArg0 := symsymbol_2
__typedArg1 := tmp7914
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7916 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp7915)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp7915
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7917 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsum, tmp7916)
}
__typedArg0 := symsum
__typedArg1 := tmp7916
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7918 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(3), tmp7917)
}
__typedArg0 := MakeNumber(3)
__typedArg1 := tmp7917
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7919 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsubst, tmp7918)
}
__typedArg0 := symsubst
__typedArg1 := tmp7918
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7920 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp7919)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp7919
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7921 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring_2, tmp7920)
}
__typedArg0 := symstring_2
__typedArg1 := tmp7920
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7922 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp7921)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp7921
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7923 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring_1_6symbol, tmp7922)
}
__typedArg0 := symstring_1_6symbol
__typedArg1 := tmp7922
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7924 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp7923)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp7923
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7925 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring_1_6n, tmp7924)
}
__typedArg0 := symstring_1_6n
__typedArg1 := tmp7924
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7926 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp7925)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp7925
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7927 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstr, tmp7926)
}
__typedArg0 := symstr
__typedArg1 := tmp7926
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7928 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(0), tmp7927)
}
__typedArg0 := MakeNumber(0)
__typedArg1 := tmp7927
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7929 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstoutput, tmp7928)
}
__typedArg0 := symstoutput
__typedArg1 := tmp7928
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7930 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(0), tmp7929)
}
__typedArg0 := MakeNumber(0)
__typedArg1 := tmp7929
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7931 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstinput, tmp7930)
}
__typedArg0 := symstinput
__typedArg1 := tmp7930
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7932 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(0), tmp7931)
}
__typedArg0 := MakeNumber(0)
__typedArg1 := tmp7931
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7933 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstep_2, tmp7932)
}
__typedArg0 := symstep_2
__typedArg1 := tmp7932
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7934 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp7933)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp7933
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7935 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstep, tmp7934)
}
__typedArg0 := symstep
__typedArg1 := tmp7934
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7936 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(0), tmp7935)
}
__typedArg0 := MakeNumber(0)
__typedArg1 := tmp7935
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7937 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symspy_2, tmp7936)
}
__typedArg0 := symspy_2
__typedArg1 := tmp7936
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7938 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp7937)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp7937
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7939 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symspy, tmp7938)
}
__typedArg0 := symspy
__typedArg1 := tmp7938
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7940 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(2), tmp7939)
}
__typedArg0 := MakeNumber(2)
__typedArg1 := tmp7939
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7941 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symspecialise, tmp7940)
}
__typedArg0 := symspecialise
__typedArg1 := tmp7940
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7942 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp7941)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp7941
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7943 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsnd, tmp7942)
}
__typedArg0 := symsnd
__typedArg1 := tmp7942
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7944 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp7943)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp7943
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7945 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsimple_1error, tmp7944)
}
__typedArg0 := symsimple_1error
__typedArg1 := tmp7944
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7946 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(2), tmp7945)
}
__typedArg0 := MakeNumber(2)
__typedArg1 := tmp7945
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7947 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symset, tmp7946)
}
__typedArg0 := symset
__typedArg1 := tmp7946
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7948 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp7947)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp7947
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7949 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symreverse, tmp7948)
}
__typedArg0 := symreverse
__typedArg1 := tmp7948
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7950 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(2), tmp7949)
}
__typedArg0 := MakeNumber(2)
__typedArg1 := tmp7949
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7951 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symremove, tmp7950)
}
__typedArg0 := symremove
__typedArg1 := tmp7950
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7952 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(0), tmp7951)
}
__typedArg0 := MakeNumber(0)
__typedArg1 := tmp7951
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7953 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symrelease, tmp7952)
}
__typedArg0 := symrelease
__typedArg1 := tmp7952
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7954 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp7953)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp7953
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7955 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symreceive, tmp7954)
}
__typedArg0 := symreceive
__typedArg1 := tmp7954
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7956 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp7955)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp7955
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7957 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4read_1unit_1string, tmp7956)
}
__typedArg0 := symshen_4read_1unit_1string
__typedArg1 := tmp7956
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7958 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp7957)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp7957
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7959 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1from_1string_1unprocessed, tmp7958)
}
__typedArg0 := symread_1from_1string_1unprocessed
__typedArg1 := tmp7958
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7960 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp7959)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp7959
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7961 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1from_1string, tmp7960)
}
__typedArg0 := symread_1from_1string
__typedArg1 := tmp7960
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7962 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp7961)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp7961
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7963 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1byte, tmp7962)
}
__typedArg0 := symread_1byte
__typedArg1 := tmp7962
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7964 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp7963)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp7963
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7965 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread, tmp7964)
}
__typedArg0 := symread
__typedArg1 := tmp7964
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7966 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp7965)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp7965
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7967 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1file, tmp7966)
}
__typedArg0 := symread_1file
__typedArg1 := tmp7966
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7968 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp7967)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp7967
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7969 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1file_1as_1bytelist, tmp7968)
}
__typedArg0 := symread_1file_1as_1bytelist
__typedArg1 := tmp7968
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7970 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp7969)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp7969
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7971 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1file_1as_1string, tmp7970)
}
__typedArg0 := symread_1file_1as_1string
__typedArg1 := tmp7970
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7972 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(4), tmp7971)
}
__typedArg0 := MakeNumber(4)
__typedArg1 := tmp7971
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7973 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symput, tmp7972)
}
__typedArg0 := symput
__typedArg1 := tmp7972
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7974 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp7973)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp7973
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7975 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symprotect, tmp7974)
}
__typedArg0 := symprotect
__typedArg1 := tmp7974
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7976 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp7975)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp7975
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7977 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympreclude_1all_1but, tmp7976)
}
__typedArg0 := sympreclude_1all_1but
__typedArg1 := tmp7976
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7978 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp7977)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp7977
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7979 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympreclude, tmp7978)
}
__typedArg0 := sympreclude
__typedArg1 := tmp7978
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7980 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp7979)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp7979
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7981 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symps, tmp7980)
}
__typedArg0 := symps
__typedArg1 := tmp7980
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7982 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(2), tmp7981)
}
__typedArg0 := MakeNumber(2)
__typedArg1 := tmp7981
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7983 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympr, tmp7982)
}
__typedArg0 := sympr
__typedArg1 := tmp7982
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7984 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp7983)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp7983
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7985 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symprofile_1results, tmp7984)
}
__typedArg0 := symprofile_1results
__typedArg1 := tmp7984
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7986 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp7985)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp7985
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7987 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symprolog_1memory, tmp7986)
}
__typedArg0 := symprolog_1memory
__typedArg1 := tmp7986
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7988 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp7987)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp7987
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7989 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4printF, tmp7988)
}
__typedArg0 := symshen_4printF
__typedArg1 := tmp7988
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7990 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp7989)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp7989
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7991 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4print_1freshterm, tmp7990)
}
__typedArg0 := symshen_4print_1freshterm
__typedArg1 := tmp7990
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7992 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp7991)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp7991
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7993 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4print_1prolog_1vector, tmp7992)
}
__typedArg0 := symshen_4print_1prolog_1vector
__typedArg1 := tmp7992
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7994 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp7993)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp7993
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7995 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symprofile, tmp7994)
}
__typedArg0 := symprofile
__typedArg1 := tmp7994
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7996 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp7995)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp7995
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7997 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symprint, tmp7996)
}
__typedArg0 := symprint
__typedArg1 := tmp7996
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7998 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp7997)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp7997
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7999 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympreclude_1all_1but, tmp7998)
}
__typedArg0 := sympreclude_1all_1but
__typedArg1 := tmp7998
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8000 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(2), tmp7999)
}
__typedArg0 := MakeNumber(2)
__typedArg1 := tmp7999
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8001 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympos, tmp8000)
}
__typedArg0 := sympos
__typedArg1 := tmp8000
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8002 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(0), tmp8001)
}
__typedArg0 := MakeNumber(0)
__typedArg1 := tmp8001
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8003 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symporters, tmp8002)
}
__typedArg0 := symporters
__typedArg1 := tmp8002
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8004 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(0), tmp8003)
}
__typedArg0 := MakeNumber(0)
__typedArg1 := tmp8003
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8005 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symport, tmp8004)
}
__typedArg0 := symport
__typedArg1 := tmp8004
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8006 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp8005)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp8005
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8007 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympackage_2, tmp8006)
}
__typedArg0 := sympackage_2
__typedArg1 := tmp8006
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8008 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(3), tmp8007)
}
__typedArg0 := MakeNumber(3)
__typedArg1 := tmp8007
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8009 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympackage, tmp8008)
}
__typedArg0 := sympackage
__typedArg1 := tmp8008
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8010 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(0), tmp8009)
}
__typedArg0 := MakeNumber(0)
__typedArg1 := tmp8009
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8011 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symos, tmp8010)
}
__typedArg0 := symos
__typedArg1 := tmp8010
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8012 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(2), tmp8011)
}
__typedArg0 := MakeNumber(2)
__typedArg1 := tmp8011
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8013 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symor, tmp8012)
}
__typedArg0 := symor
__typedArg1 := tmp8012
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8014 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(0), tmp8013)
}
__typedArg0 := MakeNumber(0)
__typedArg1 := tmp8013
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8015 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symoptimise_2, tmp8014)
}
__typedArg0 := symoptimise_2
__typedArg1 := tmp8014
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8016 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp8015)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp8015
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8017 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symoptimise, tmp8016)
}
__typedArg0 := symoptimise
__typedArg1 := tmp8016
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8018 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(2), tmp8017)
}
__typedArg0 := MakeNumber(2)
__typedArg1 := tmp8017
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8019 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symopen, tmp8018)
}
__typedArg0 := symopen
__typedArg1 := tmp8018
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8020 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp8019)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp8019
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8021 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symoccurs_1check, tmp8020)
}
__typedArg0 := symoccurs_1check
__typedArg1 := tmp8020
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8022 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(0), tmp8021)
}
__typedArg0 := MakeNumber(0)
__typedArg1 := tmp8021
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8023 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symoccurs_2, tmp8022)
}
__typedArg0 := symoccurs_2
__typedArg1 := tmp8022
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8024 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(2), tmp8023)
}
__typedArg0 := MakeNumber(2)
__typedArg1 := tmp8023
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8025 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symoccurrences, tmp8024)
}
__typedArg0 := symoccurrences
__typedArg1 := tmp8024
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8026 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp8025)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp8025
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8027 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symoccurs_1check, tmp8026)
}
__typedArg0 := symoccurs_1check
__typedArg1 := tmp8026
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8028 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp8027)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp8027
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8029 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber_2, tmp8028)
}
__typedArg0 := symnumber_2
__typedArg1 := tmp8028
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8030 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp8029)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp8029
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8031 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symn_1_6string, tmp8030)
}
__typedArg0 := symn_1_6string
__typedArg1 := tmp8030
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8032 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(2), tmp8031)
}
__typedArg0 := MakeNumber(2)
__typedArg1 := tmp8031
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8033 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnth, tmp8032)
}
__typedArg0 := symnth
__typedArg1 := tmp8032
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8034 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp8033)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp8033
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8035 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnot, tmp8034)
}
__typedArg0 := symnot
__typedArg1 := tmp8034
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8036 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp8035)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp8035
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8037 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnl, tmp8036)
}
__typedArg0 := symnl
__typedArg1 := tmp8036
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8038 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp8037)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp8037
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8039 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symmaxinferences, tmp8038)
}
__typedArg0 := symmaxinferences
__typedArg1 := tmp8038
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8040 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(2), tmp8039)
}
__typedArg0 := MakeNumber(2)
__typedArg1 := tmp8039
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8041 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symmapcan, tmp8040)
}
__typedArg0 := symmapcan
__typedArg1 := tmp8040
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8042 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(2), tmp8041)
}
__typedArg0 := MakeNumber(2)
__typedArg1 := tmp8041
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8043 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symmap, tmp8042)
}
__typedArg0 := symmap
__typedArg1 := tmp8042
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8044 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp8043)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp8043
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8045 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symmacroexpand, tmp8044)
}
__typedArg0 := symmacroexpand
__typedArg1 := tmp8044
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8046 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp8045)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp8045
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8047 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp8046)
}
__typedArg0 := symvector
__typedArg1 := tmp8046
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8048 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(2), tmp8047)
}
__typedArg0 := MakeNumber(2)
__typedArg1 := tmp8047
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8049 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_a, tmp8048)
}
__typedArg0 := sym_5_a
__typedArg1 := tmp8048
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8050 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(2), tmp8049)
}
__typedArg0 := MakeNumber(2)
__typedArg1 := tmp8049
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8051 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5, tmp8050)
}
__typedArg0 := sym_5
__typedArg1 := tmp8050
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8052 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp8051)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp8051
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8053 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symload, tmp8052)
}
__typedArg0 := symload
__typedArg1 := tmp8052
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8054 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp8053)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp8053
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8055 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp8054)
}
__typedArg0 := symlist
__typedArg1 := tmp8054
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8056 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp8055)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp8055
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8057 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlineread, tmp8056)
}
__typedArg0 := symlineread
__typedArg1 := tmp8056
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8058 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp8057)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp8057
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8059 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlimit, tmp8058)
}
__typedArg0 := symlimit
__typedArg1 := tmp8058
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8060 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp8059)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp8059
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8061 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlength, tmp8060)
}
__typedArg0 := symlength
__typedArg1 := tmp8060
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8062 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(0), tmp8061)
}
__typedArg0 := MakeNumber(0)
__typedArg1 := tmp8061
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8063 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlanguage, tmp8062)
}
__typedArg0 := symlanguage
__typedArg1 := tmp8062
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8064 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(6), tmp8063)
}
__typedArg0 := MakeNumber(6)
__typedArg1 := tmp8063
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8065 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symis_b, tmp8064)
}
__typedArg0 := symis_b
__typedArg1 := tmp8064
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8066 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(6), tmp8065)
}
__typedArg0 := MakeNumber(6)
__typedArg1 := tmp8065
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8067 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symis, tmp8066)
}
__typedArg0 := symis
__typedArg1 := tmp8066
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8068 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(0), tmp8067)
}
__typedArg0 := MakeNumber(0)
__typedArg1 := tmp8067
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8069 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symit, tmp8068)
}
__typedArg0 := symit
__typedArg1 := tmp8068
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8070 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp8069)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp8069
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8071 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminternal, tmp8070)
}
__typedArg0 := syminternal
__typedArg1 := tmp8070
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8072 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(2), tmp8071)
}
__typedArg0 := MakeNumber(2)
__typedArg1 := tmp8071
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8073 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symintersection, tmp8072)
}
__typedArg0 := symintersection
__typedArg1 := tmp8072
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8074 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp8073)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp8073
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8075 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminclude_1all_1but, tmp8074)
}
__typedArg0 := syminclude_1all_1but
__typedArg1 := tmp8074
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8076 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(0), tmp8075)
}
__typedArg0 := MakeNumber(0)
__typedArg1 := tmp8075
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8077 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symimplementation, tmp8076)
}
__typedArg0 := symimplementation
__typedArg1 := tmp8076
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8078 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(2), tmp8077)
}
__typedArg0 := MakeNumber(2)
__typedArg1 := tmp8077
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8079 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminput_7, tmp8078)
}
__typedArg0 := syminput_7
__typedArg1 := tmp8078
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8080 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp8079)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp8079
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8081 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminput, tmp8080)
}
__typedArg0 := syminput
__typedArg1 := tmp8080
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8082 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(0), tmp8081)
}
__typedArg0 := MakeNumber(0)
__typedArg1 := tmp8081
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8083 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminferences, tmp8082)
}
__typedArg0 := syminferences
__typedArg1 := tmp8082
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8084 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp8083)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp8083
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8085 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symintern, tmp8084)
}
__typedArg0 := symintern
__typedArg1 := tmp8084
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8086 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp8085)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp8085
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8087 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminternal, tmp8086)
}
__typedArg0 := syminternal
__typedArg1 := tmp8086
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8088 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp8087)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp8087
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8089 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminteger_2, tmp8088)
}
__typedArg0 := syminteger_2
__typedArg1 := tmp8088
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8090 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp8089)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp8089
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8091 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symin_1package, tmp8090)
}
__typedArg0 := symin_1package
__typedArg1 := tmp8090
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8092 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(0), tmp8091)
}
__typedArg0 := MakeNumber(0)
__typedArg1 := tmp8091
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8093 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symincluded, tmp8092)
}
__typedArg0 := symincluded
__typedArg1 := tmp8092
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8094 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp8093)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp8093
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8095 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminclude, tmp8094)
}
__typedArg0 := syminclude
__typedArg1 := tmp8094
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8096 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(3), tmp8095)
}
__typedArg0 := MakeNumber(3)
__typedArg1 := tmp8095
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8097 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp8096)
}
__typedArg0 := symif
__typedArg1 := tmp8096
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8098 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp8097)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp8097
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8099 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhush, tmp8098)
}
__typedArg0 := symhush
__typedArg1 := tmp8098
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8100 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(0), tmp8099)
}
__typedArg0 := MakeNumber(0)
__typedArg1 := tmp8099
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8101 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhush_2, tmp8100)
}
__typedArg0 := symhush_2
__typedArg1 := tmp8100
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8102 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp8101)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp8101
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8103 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhead, tmp8102)
}
__typedArg0 := symhead
__typedArg1 := tmp8102
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8104 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp8103)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp8103
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8105 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhdstr, tmp8104)
}
__typedArg0 := symhdstr
__typedArg1 := tmp8104
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8106 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp8105)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp8105
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8107 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhdv, tmp8106)
}
__typedArg0 := symhdv
__typedArg1 := tmp8106
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8108 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp8107)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp8107
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8109 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhd, tmp8108)
}
__typedArg0 := symhd
__typedArg1 := tmp8108
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8110 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(2), tmp8109)
}
__typedArg0 := MakeNumber(2)
__typedArg1 := tmp8109
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8111 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhash, tmp8110)
}
__typedArg0 := symhash
__typedArg1 := tmp8110
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8112 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(2), tmp8111)
}
__typedArg0 := MakeNumber(2)
__typedArg1 := tmp8111
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8113 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_a, tmp8112)
}
__typedArg0 := sym_a
__typedArg1 := tmp8112
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8114 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(2), tmp8113)
}
__typedArg0 := MakeNumber(2)
__typedArg1 := tmp8113
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8115 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_6_a, tmp8114)
}
__typedArg0 := sym_6_a
__typedArg1 := tmp8114
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8116 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(2), tmp8115)
}
__typedArg0 := MakeNumber(2)
__typedArg1 := tmp8115
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8117 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_6, tmp8116)
}
__typedArg0 := sym_6
__typedArg1 := tmp8116
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8118 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(2), tmp8117)
}
__typedArg0 := MakeNumber(2)
__typedArg1 := tmp8117
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8119 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1vector, tmp8118)
}
__typedArg0 := sym_5_1vector
__typedArg1 := tmp8118
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8120 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(2), tmp8119)
}
__typedArg0 := MakeNumber(2)
__typedArg1 := tmp8119
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8121 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1address, tmp8120)
}
__typedArg0 := sym_5_1address
__typedArg1 := tmp8120
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8122 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(3), tmp8121)
}
__typedArg0 := MakeNumber(3)
__typedArg1 := tmp8121
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8123 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symaddress_1_6, tmp8122)
}
__typedArg0 := symaddress_1_6
__typedArg1 := tmp8122
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8124 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp8123)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp8123
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8125 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symget_1time, tmp8124)
}
__typedArg0 := symget_1time
__typedArg1 := tmp8124
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8126 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(3), tmp8125)
}
__typedArg0 := MakeNumber(3)
__typedArg1 := tmp8125
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8127 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symget, tmp8126)
}
__typedArg0 := symget
__typedArg1 := tmp8126
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8128 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp8127)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp8127
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8129 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symgensym, tmp8128)
}
__typedArg0 := symgensym
__typedArg1 := tmp8128
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8130 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp8129)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp8129
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8131 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfunction, tmp8130)
}
__typedArg0 := symfunction
__typedArg1 := tmp8130
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8132 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp8131)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp8131
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8133 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfn, tmp8132)
}
__typedArg0 := symfn
__typedArg1 := tmp8132
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8134 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp8133)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp8133
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8135 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfst, tmp8134)
}
__typedArg0 := symfst
__typedArg1 := tmp8134
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8136 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(0), tmp8135)
}
__typedArg0 := MakeNumber(0)
__typedArg1 := tmp8135
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8137 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfresh, tmp8136)
}
__typedArg0 := symfresh
__typedArg1 := tmp8136
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8138 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp8137)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp8137
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8139 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfreeze, tmp8138)
}
__typedArg0 := symfreeze
__typedArg1 := tmp8138
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8140 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(5), tmp8139)
}
__typedArg0 := MakeNumber(5)
__typedArg1 := tmp8139
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8141 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfork, tmp8140)
}
__typedArg0 := symfork
__typedArg1 := tmp8140
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8142 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp8141)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp8141
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8143 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symforeign, tmp8142)
}
__typedArg0 := symforeign
__typedArg1 := tmp8142
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8144 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(7), tmp8143)
}
__typedArg0 := MakeNumber(7)
__typedArg1 := tmp8143
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8145 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfindall, tmp8144)
}
__typedArg0 := symfindall
__typedArg1 := tmp8144
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8146 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(2), tmp8145)
}
__typedArg0 := MakeNumber(2)
__typedArg1 := tmp8145
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8147 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfix, tmp8146)
}
__typedArg0 := symfix
__typedArg1 := tmp8146
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8148 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(0), tmp8147)
}
__typedArg0 := MakeNumber(0)
__typedArg1 := tmp8147
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8149 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfail, tmp8148)
}
__typedArg0 := symfail
__typedArg1 := tmp8148
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8150 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(2), tmp8149)
}
__typedArg0 := MakeNumber(2)
__typedArg1 := tmp8149
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8151 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfail_1if, tmp8150)
}
__typedArg0 := symfail_1if
__typedArg1 := tmp8150
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8152 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(0), tmp8151)
}
__typedArg0 := MakeNumber(0)
__typedArg1 := tmp8151
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8153 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfactorise_2, tmp8152)
}
__typedArg0 := symfactorise_2
__typedArg1 := tmp8152
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8154 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp8153)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp8153
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8155 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfactorise, tmp8154)
}
__typedArg0 := symfactorise
__typedArg1 := tmp8154
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8156 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp8155)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp8155
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8157 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symexternal, tmp8156)
}
__typedArg0 := symexternal
__typedArg1 := tmp8156
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8158 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp8157)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp8157
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8159 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symexplode, tmp8158)
}
__typedArg0 := symexplode
__typedArg1 := tmp8158
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8160 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp8159)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp8159
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8161 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symeval_1kl, tmp8160)
}
__typedArg0 := symeval_1kl
__typedArg1 := tmp8160
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8162 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp8161)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp8161
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8163 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symeval, tmp8162)
}
__typedArg0 := symeval
__typedArg1 := tmp8162
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8164 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp8163)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp8163
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8165 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symerror_1to_1string, tmp8164)
}
__typedArg0 := symerror_1to_1string
__typedArg1 := tmp8164
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8166 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp8165)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp8165
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8167 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symexternal, tmp8166)
}
__typedArg0 := symexternal
__typedArg1 := tmp8166
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8168 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp8167)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp8167
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8169 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symenable_1type_1theory, tmp8168)
}
__typedArg0 := symenable_1type_1theory
__typedArg1 := tmp8168
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8170 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp8169)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp8169
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8171 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symempty_2, tmp8170)
}
__typedArg0 := symempty_2
__typedArg1 := tmp8170
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8172 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(2), tmp8171)
}
__typedArg0 := MakeNumber(2)
__typedArg1 := tmp8171
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8173 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symelement_2, tmp8172)
}
__typedArg0 := symelement_2
__typedArg1 := tmp8172
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8174 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(2), tmp8173)
}
__typedArg0 := MakeNumber(2)
__typedArg1 := tmp8173
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8175 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdo, tmp8174)
}
__typedArg0 := symdo
__typedArg1 := tmp8174
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8176 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(2), tmp8175)
}
__typedArg0 := MakeNumber(2)
__typedArg1 := tmp8175
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8177 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdifference, tmp8176)
}
__typedArg0 := symdifference
__typedArg1 := tmp8176
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8178 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp8177)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp8177
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8179 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdestroy, tmp8178)
}
__typedArg0 := symdestroy
__typedArg1 := tmp8178
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8180 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(2), tmp8179)
}
__typedArg0 := MakeNumber(2)
__typedArg1 := tmp8179
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8181 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdeclare, tmp8180)
}
__typedArg0 := symdeclare
__typedArg1 := tmp8180
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8182 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(0), tmp8181)
}
__typedArg0 := MakeNumber(0)
__typedArg1 := tmp8181
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8183 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdatatypes, tmp8182)
}
__typedArg0 := symdatatypes
__typedArg1 := tmp8182
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8184 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp8183)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp8183
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8185 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symclose, tmp8184)
}
__typedArg0 := symclose
__typedArg1 := tmp8184
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8186 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(2), tmp8185)
}
__typedArg0 := MakeNumber(2)
__typedArg1 := tmp8185
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8187 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcn, tmp8186)
}
__typedArg0 := symcn
__typedArg1 := tmp8186
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8188 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp8187)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp8187
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8189 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons_2, tmp8188)
}
__typedArg0 := symcons_2
__typedArg1 := tmp8188
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8190 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(2), tmp8189)
}
__typedArg0 := MakeNumber(2)
__typedArg1 := tmp8189
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8191 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp8190)
}
__typedArg0 := symcons
__typedArg1 := tmp8190
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8192 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(2), tmp8191)
}
__typedArg0 := MakeNumber(2)
__typedArg1 := tmp8191
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8193 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symconcat, tmp8192)
}
__typedArg0 := symconcat
__typedArg1 := tmp8192
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8194 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(2), tmp8193)
}
__typedArg0 := MakeNumber(2)
__typedArg1 := tmp8193
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8195 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcompile, tmp8194)
}
__typedArg0 := symcompile
__typedArg1 := tmp8194
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8196 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp8195)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp8195
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8197 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcd, tmp8196)
}
__typedArg0 := symcd
__typedArg1 := tmp8196
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8198 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(5), tmp8197)
}
__typedArg0 := MakeNumber(5)
__typedArg1 := tmp8197
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8199 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcall, tmp8198)
}
__typedArg0 := symcall
__typedArg1 := tmp8198
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8200 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(6), tmp8199)
}
__typedArg0 := MakeNumber(6)
__typedArg1 := tmp8199
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8201 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symbind, tmp8200)
}
__typedArg0 := symbind
__typedArg1 := tmp8200
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8202 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp8201)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp8201
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8203 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symbound_2, tmp8202)
}
__typedArg0 := symbound_2
__typedArg1 := tmp8202
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8204 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp8203)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp8203
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8205 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symbootstrap, tmp8204)
}
__typedArg0 := symbootstrap
__typedArg1 := tmp8204
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8206 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp8205)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp8205
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8207 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean_2, tmp8206)
}
__typedArg0 := symboolean_2
__typedArg1 := tmp8206
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8208 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp8207)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp8207
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8209 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symatom_2, tmp8208)
}
__typedArg0 := symatom_2
__typedArg1 := tmp8208
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8210 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(2), tmp8209)
}
__typedArg0 := MakeNumber(2)
__typedArg1 := tmp8209
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8211 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symassoc, tmp8210)
}
__typedArg0 := symassoc
__typedArg1 := tmp8210
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8212 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp8211)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp8211
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8213 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symarity, tmp8212)
}
__typedArg0 := symarity
__typedArg1 := tmp8212
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8214 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(2), tmp8213)
}
__typedArg0 := MakeNumber(2)
__typedArg1 := tmp8213
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8215 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symappend, tmp8214)
}
__typedArg0 := symappend
__typedArg1 := tmp8214
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8216 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(2), tmp8215)
}
__typedArg0 := MakeNumber(2)
__typedArg1 := tmp8215
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8217 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symand, tmp8216)
}
__typedArg0 := symand
__typedArg1 := tmp8216
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8218 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(2), tmp8217)
}
__typedArg0 := MakeNumber(2)
__typedArg1 := tmp8217
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8219 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symadjoin, tmp8218)
}
__typedArg0 := symadjoin
__typedArg1 := tmp8218
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8220 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(3), tmp8219)
}
__typedArg0 := MakeNumber(3)
__typedArg1 := tmp8219
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8221 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symaddress_1_6, tmp8220)
}
__typedArg0 := symaddress_1_6
__typedArg1 := tmp8220
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8222 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp8221)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp8221
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8223 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symabsvector, tmp8222)
}
__typedArg0 := symabsvector
__typedArg1 := tmp8222
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8224 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp8223)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp8223
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8225 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symabsvector_2, tmp8224)
}
__typedArg0 := symabsvector_2
__typedArg1 := tmp8224
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8226 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), tmp8225)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp8225
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8227 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symabsolute, tmp8226)
}
__typedArg0 := symabsolute
__typedArg1 := tmp8226
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8228 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(0), tmp8227)
}
__typedArg0 := MakeNumber(0)
__typedArg1 := tmp8227
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8229 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symabort, tmp8228)
}
__typedArg0 := symabort
__typedArg1 := tmp8228
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8230 := Call(__e, PrimFunc(symshen_4initialise_1arity_1table), tmp8229)


_ = tmp8230

tmp8231 := MakeNative(func(__e *ControlFlow) {
V5776 := __e.Get(1)
_ = V5776
tmp8232 := MakeNative(func(__e *ControlFlow) {
W5777 := __e.Get(1)
_ = W5777
tmp8233 := MakeNative(func(__e *ControlFlow) {
W5778 := __e.Get(1)
_ = W5778
__e.Return(V5776)
return
}, 1)

tmp8234 := Call(__e, PrimFunc(symadjoin), V5776, W5777)


tmp8235 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp8236 := Call(__e, PrimFunc(symput), symshen, symshen_4external_1symbols, tmp8234, tmp8235)


__e.TailApply(tmp8233, tmp8236)
return


}, 1)

tmp8237 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp8238 := Call(__e, PrimFunc(symget), symshen, symshen_4external_1symbols, tmp8237)


__e.TailApply(tmp8232, tmp8238)
return


}, 1)

tmp8239 := Call(__e, ns2_1set, symsystemf, tmp8231)


_ = tmp8239

tmp8240 := MakeNative(func(__e *ControlFlow) {
V5779 := __e.Get(1)
_ = V5779
V5780 := __e.Get(2)
_ = V5780
tmp8242 := Call(__e, PrimFunc(symelement_2), V5779, V5780)


if True == tmp8242 {
__e.Return(V5780)
return
} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5779, V5780)
}
__typedArg0 := V5779
__typedArg1 := V5780
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return
}


}, 2)

tmp8243 := Call(__e, ns2_1set, symadjoin, tmp8240)


_ = tmp8243

tmp8244 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp8245 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(";"))
}
__typedArg0 := MakeString(";")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp8246 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":="))
}
__typedArg0 := MakeString(":=")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp8247 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(","))
}
__typedArg0 := MakeString(",")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp8248 := Call(__e, PrimFunc(symvector), MakeNumber(0))


tmp8249 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString("bar!"))
}
__typedArg0 := MakeString("bar!")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp8250 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symabort, Nil)
}
__typedArg0 := symabort
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8251 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symabsolute, tmp8250)
}
__typedArg0 := symabsolute
__typedArg1 := tmp8250
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8252 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symabsvector, tmp8251)
}
__typedArg0 := symabsvector
__typedArg1 := tmp8251
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8253 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symabsvector_2, tmp8252)
}
__typedArg0 := symabsvector_2
__typedArg1 := tmp8252
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8254 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symaddress_1_6, tmp8253)
}
__typedArg0 := symaddress_1_6
__typedArg1 := tmp8253
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8255 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1address, tmp8254)
}
__typedArg0 := sym_5_1address
__typedArg1 := tmp8254
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8256 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symadjoin, tmp8255)
}
__typedArg0 := symadjoin
__typedArg1 := tmp8255
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8257 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symand, tmp8256)
}
__typedArg0 := symand
__typedArg1 := tmp8256
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8258 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symappend, tmp8257)
}
__typedArg0 := symappend
__typedArg1 := tmp8257
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8259 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symarity, tmp8258)
}
__typedArg0 := symarity
__typedArg1 := tmp8258
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8260 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symassoc, tmp8259)
}
__typedArg0 := symassoc
__typedArg1 := tmp8259
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8261 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symassertz, tmp8260)
}
__typedArg0 := symassertz
__typedArg1 := tmp8260
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8262 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symasserta, tmp8261)
}
__typedArg0 := symasserta
__typedArg1 := tmp8261
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8263 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symatom_2, tmp8262)
}
__typedArg0 := symatom_2
__typedArg1 := tmp8262
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8264 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstep_2, tmp8263)
}
__typedArg0 := symstep_2
__typedArg1 := tmp8263
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8265 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symspy_2, tmp8264)
}
__typedArg0 := symspy_2
__typedArg1 := tmp8264
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8266 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8249, tmp8265)
}
__typedArg0 := tmp8249
__typedArg1 := tmp8265
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8267 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symbootstrap, tmp8266)
}
__typedArg0 := symbootstrap
__typedArg1 := tmp8266
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8268 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, tmp8267)
}
__typedArg0 := symboolean
__typedArg1 := tmp8267
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8269 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean_2, tmp8268)
}
__typedArg0 := symboolean_2
__typedArg1 := tmp8268
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8270 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symbound_2, tmp8269)
}
__typedArg0 := symbound_2
__typedArg1 := tmp8269
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8271 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symbind, tmp8270)
}
__typedArg0 := symbind
__typedArg1 := tmp8270
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8272 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symclose, tmp8271)
}
__typedArg0 := symclose
__typedArg1 := tmp8271
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8273 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcall, tmp8272)
}
__typedArg0 := symcall
__typedArg1 := tmp8272
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8274 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcases, tmp8273)
}
__typedArg0 := symcases
__typedArg1 := tmp8273
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8275 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcd, tmp8274)
}
__typedArg0 := symcd
__typedArg1 := tmp8274
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8276 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcompile, tmp8275)
}
__typedArg0 := symcompile
__typedArg1 := tmp8275
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8277 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symconcat, tmp8276)
}
__typedArg0 := symconcat
__typedArg1 := tmp8276
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8278 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcond, tmp8277)
}
__typedArg0 := symcond
__typedArg1 := tmp8277
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8279 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp8278)
}
__typedArg0 := symcons
__typedArg1 := tmp8278
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8280 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons_2, tmp8279)
}
__typedArg0 := symcons_2
__typedArg1 := tmp8279
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8281 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcn, tmp8280)
}
__typedArg0 := symcn
__typedArg1 := tmp8280
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8282 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symctxt, tmp8281)
}
__typedArg0 := symctxt
__typedArg1 := tmp8281
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8283 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdatatypes, tmp8282)
}
__typedArg0 := symdatatypes
__typedArg1 := tmp8282
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8284 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdatatype, tmp8283)
}
__typedArg0 := symdatatype
__typedArg1 := tmp8283
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8285 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdeclare, tmp8284)
}
__typedArg0 := symdeclare
__typedArg1 := tmp8284
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8286 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefprolog, tmp8285)
}
__typedArg0 := symdefprolog
__typedArg1 := tmp8285
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8287 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefcc, tmp8286)
}
__typedArg0 := symdefcc
__typedArg1 := tmp8286
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8288 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefmacro, tmp8287)
}
__typedArg0 := symdefmacro
__typedArg1 := tmp8287
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8289 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefine, tmp8288)
}
__typedArg0 := symdefine
__typedArg1 := tmp8288
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8290 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefun, tmp8289)
}
__typedArg0 := symdefun
__typedArg1 := tmp8289
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8291 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdestroy, tmp8290)
}
__typedArg0 := symdestroy
__typedArg1 := tmp8290
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8292 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdifference, tmp8291)
}
__typedArg0 := symdifference
__typedArg1 := tmp8291
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8293 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdo, tmp8292)
}
__typedArg0 := symdo
__typedArg1 := tmp8292
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8294 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symelement_2, tmp8293)
}
__typedArg0 := symelement_2
__typedArg1 := tmp8293
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8295 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symempty_2, tmp8294)
}
__typedArg0 := symempty_2
__typedArg1 := tmp8294
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8296 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symerror, tmp8295)
}
__typedArg0 := symerror
__typedArg1 := tmp8295
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8297 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symerror_1to_1string, tmp8296)
}
__typedArg0 := symerror_1to_1string
__typedArg1 := tmp8296
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8298 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symeval, tmp8297)
}
__typedArg0 := symeval
__typedArg1 := tmp8297
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8299 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symeval_1kl, tmp8298)
}
__typedArg0 := symeval_1kl
__typedArg1 := tmp8298
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8300 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symexception, tmp8299)
}
__typedArg0 := symexception
__typedArg1 := tmp8299
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8301 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symexternal, tmp8300)
}
__typedArg0 := symexternal
__typedArg1 := tmp8300
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8302 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symexplode, tmp8301)
}
__typedArg0 := symexplode
__typedArg1 := tmp8301
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8303 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symenable_1type_1theory, tmp8302)
}
__typedArg0 := symenable_1type_1theory
__typedArg1 := tmp8302
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8304 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(False, tmp8303)
}
__typedArg0 := False
__typedArg1 := tmp8303
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8305 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfindall, tmp8304)
}
__typedArg0 := symfindall
__typedArg1 := tmp8304
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8306 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfactorise_2, tmp8305)
}
__typedArg0 := symfactorise_2
__typedArg1 := tmp8305
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8307 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfactorise, tmp8306)
}
__typedArg0 := symfactorise
__typedArg1 := tmp8306
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8308 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfail_1if, tmp8307)
}
__typedArg0 := symfail_1if
__typedArg1 := tmp8307
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8309 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfail, tmp8308)
}
__typedArg0 := symfail
__typedArg1 := tmp8308
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8310 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfile, tmp8309)
}
__typedArg0 := symfile
__typedArg1 := tmp8309
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8311 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfix, tmp8310)
}
__typedArg0 := symfix
__typedArg1 := tmp8310
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8312 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symforeign, tmp8311)
}
__typedArg0 := symforeign
__typedArg1 := tmp8311
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8313 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfork, tmp8312)
}
__typedArg0 := symfork
__typedArg1 := tmp8312
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8314 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfresh, tmp8313)
}
__typedArg0 := symfresh
__typedArg1 := tmp8313
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8315 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfreeze, tmp8314)
}
__typedArg0 := symfreeze
__typedArg1 := tmp8314
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8316 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfst, tmp8315)
}
__typedArg0 := symfst
__typedArg1 := tmp8315
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8317 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfunction, tmp8316)
}
__typedArg0 := symfunction
__typedArg1 := tmp8316
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8318 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfn, tmp8317)
}
__typedArg0 := symfn
__typedArg1 := tmp8317
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8319 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symgensym, tmp8318)
}
__typedArg0 := symgensym
__typedArg1 := tmp8318
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8320 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symget_1time, tmp8319)
}
__typedArg0 := symget_1time
__typedArg1 := tmp8319
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8321 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symget, tmp8320)
}
__typedArg0 := symget
__typedArg1 := tmp8320
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8322 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhash, tmp8321)
}
__typedArg0 := symhash
__typedArg1 := tmp8321
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8323 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhdstr, tmp8322)
}
__typedArg0 := symhdstr
__typedArg1 := tmp8322
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8324 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhdv, tmp8323)
}
__typedArg0 := symhdv
__typedArg1 := tmp8323
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8325 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhd, tmp8324)
}
__typedArg0 := symhd
__typedArg1 := tmp8324
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8326 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhead, tmp8325)
}
__typedArg0 := symhead
__typedArg1 := tmp8325
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8327 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhush_2, tmp8326)
}
__typedArg0 := symhush_2
__typedArg1 := tmp8326
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8328 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhush_2, tmp8327)
}
__typedArg0 := symhush_2
__typedArg1 := tmp8327
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8329 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp8328)
}
__typedArg0 := symif
__typedArg1 := tmp8328
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8330 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symimplementation, tmp8329)
}
__typedArg0 := symimplementation
__typedArg1 := tmp8329
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8331 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminternal, tmp8330)
}
__typedArg0 := syminternal
__typedArg1 := tmp8330
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8332 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symin_1package, tmp8331)
}
__typedArg0 := symin_1package
__typedArg1 := tmp8331
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8333 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symin, tmp8332)
}
__typedArg0 := symin
__typedArg1 := tmp8332
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8334 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symis_b, tmp8333)
}
__typedArg0 := symis_b
__typedArg1 := tmp8333
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8335 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symis, tmp8334)
}
__typedArg0 := symis
__typedArg1 := tmp8334
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8336 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symit, tmp8335)
}
__typedArg0 := symit
__typedArg1 := tmp8335
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8337 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminclude_1all_1but, tmp8336)
}
__typedArg0 := syminclude_1all_1but
__typedArg1 := tmp8336
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8338 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminclude, tmp8337)
}
__typedArg0 := syminclude
__typedArg1 := tmp8337
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8339 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symincluded, tmp8338)
}
__typedArg0 := symincluded
__typedArg1 := tmp8338
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8340 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminput_7, tmp8339)
}
__typedArg0 := syminput_7
__typedArg1 := tmp8339
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8341 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminput, tmp8340)
}
__typedArg0 := syminput
__typedArg1 := tmp8340
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8342 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminteger_2, tmp8341)
}
__typedArg0 := syminteger_2
__typedArg1 := tmp8341
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8343 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symintern, tmp8342)
}
__typedArg0 := symintern
__typedArg1 := tmp8342
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8344 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminferences, tmp8343)
}
__typedArg0 := syminferences
__typedArg1 := tmp8343
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8345 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symintersection, tmp8344)
}
__typedArg0 := symintersection
__typedArg1 := tmp8344
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8346 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symis, tmp8345)
}
__typedArg0 := symis
__typedArg1 := tmp8345
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8347 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlanguage, tmp8346)
}
__typedArg0 := symlanguage
__typedArg1 := tmp8346
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8348 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp8347)
}
__typedArg0 := symlambda
__typedArg1 := tmp8347
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8349 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlazy, tmp8348)
}
__typedArg0 := symlazy
__typedArg1 := tmp8348
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8350 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp8349)
}
__typedArg0 := symlet
__typedArg1 := tmp8349
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8351 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlength, tmp8350)
}
__typedArg0 := symlength
__typedArg1 := tmp8350
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8352 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlimit, tmp8351)
}
__typedArg0 := symlimit
__typedArg1 := tmp8351
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8353 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlineread, tmp8352)
}
__typedArg0 := symlineread
__typedArg1 := tmp8352
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8354 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp8353)
}
__typedArg0 := symlist
__typedArg1 := tmp8353
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8355 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symloaded, tmp8354)
}
__typedArg0 := symloaded
__typedArg1 := tmp8354
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8356 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symload, tmp8355)
}
__typedArg0 := symload
__typedArg1 := tmp8355
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8357 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symmake_1string, tmp8356)
}
__typedArg0 := symmake_1string
__typedArg1 := tmp8356
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8358 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symmap, tmp8357)
}
__typedArg0 := symmap
__typedArg1 := tmp8357
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8359 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symmapcan, tmp8358)
}
__typedArg0 := symmapcan
__typedArg1 := tmp8358
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8360 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symmaxinferences, tmp8359)
}
__typedArg0 := symmaxinferences
__typedArg1 := tmp8359
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8361 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symmacroexpand, tmp8360)
}
__typedArg0 := symmacroexpand
__typedArg1 := tmp8360
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8362 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symmode, tmp8361)
}
__typedArg0 := symmode
__typedArg1 := tmp8361
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8363 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnl, tmp8362)
}
__typedArg0 := symnl
__typedArg1 := tmp8362
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8364 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnot, tmp8363)
}
__typedArg0 := symnot
__typedArg1 := tmp8363
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8365 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnth, tmp8364)
}
__typedArg0 := symnth
__typedArg1 := tmp8364
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8366 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnull, tmp8365)
}
__typedArg0 := symnull
__typedArg1 := tmp8365
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8367 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp8366)
}
__typedArg0 := symnumber
__typedArg1 := tmp8366
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8368 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber_2, tmp8367)
}
__typedArg0 := symnumber_2
__typedArg1 := tmp8367
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8369 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symn_1_6string, tmp8368)
}
__typedArg0 := symn_1_6string
__typedArg1 := tmp8368
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8370 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symoccurs_2, tmp8369)
}
__typedArg0 := symoccurs_2
__typedArg1 := tmp8369
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8371 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symoccurs_1check, tmp8370)
}
__typedArg0 := symoccurs_1check
__typedArg1 := tmp8370
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8372 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symoccurrences, tmp8371)
}
__typedArg0 := symoccurrences
__typedArg1 := tmp8371
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8373 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symopen, tmp8372)
}
__typedArg0 := symopen
__typedArg1 := tmp8372
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8374 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symoptimise_2, tmp8373)
}
__typedArg0 := symoptimise_2
__typedArg1 := tmp8373
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8375 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symoptimise, tmp8374)
}
__typedArg0 := symoptimise
__typedArg1 := tmp8374
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8376 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symor, tmp8375)
}
__typedArg0 := symor
__typedArg1 := tmp8375
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8377 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symos, tmp8376)
}
__typedArg0 := symos
__typedArg1 := tmp8376
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8378 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symout, tmp8377)
}
__typedArg0 := symout
__typedArg1 := tmp8377
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8379 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symoutput, tmp8378)
}
__typedArg0 := symoutput
__typedArg1 := tmp8378
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8380 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympackage, tmp8379)
}
__typedArg0 := sympackage
__typedArg1 := tmp8379
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8381 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symport, tmp8380)
}
__typedArg0 := symport
__typedArg1 := tmp8380
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8382 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symporters, tmp8381)
}
__typedArg0 := symporters
__typedArg1 := tmp8381
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8383 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympos, tmp8382)
}
__typedArg0 := sympos
__typedArg1 := tmp8382
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8384 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympr, tmp8383)
}
__typedArg0 := sympr
__typedArg1 := tmp8383
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8385 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symprint, tmp8384)
}
__typedArg0 := symprint
__typedArg1 := tmp8384
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8386 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symprolog_1memory, tmp8385)
}
__typedArg0 := symprolog_1memory
__typedArg1 := tmp8385
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8387 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symprofile, tmp8386)
}
__typedArg0 := symprofile
__typedArg1 := tmp8386
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8388 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symprofile_1results, tmp8387)
}
__typedArg0 := symprofile_1results
__typedArg1 := tmp8387
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8389 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symprotect, tmp8388)
}
__typedArg0 := symprotect
__typedArg1 := tmp8388
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8390 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symprolog_2, tmp8389)
}
__typedArg0 := symprolog_2
__typedArg1 := tmp8389
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8391 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symps, tmp8390)
}
__typedArg0 := symps
__typedArg1 := tmp8390
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8392 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympreclude_1all_1but, tmp8391)
}
__typedArg0 := sympreclude_1all_1but
__typedArg1 := tmp8391
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8393 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympreclude, tmp8392)
}
__typedArg0 := sympreclude
__typedArg1 := tmp8392
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8394 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symput, tmp8393)
}
__typedArg0 := symput
__typedArg1 := tmp8393
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8395 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympackage_2, tmp8394)
}
__typedArg0 := sympackage_2
__typedArg1 := tmp8394
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8396 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1from_1string_1unprocessed, tmp8395)
}
__typedArg0 := symread_1from_1string_1unprocessed
__typedArg1 := tmp8395
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8397 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1from_1string, tmp8396)
}
__typedArg0 := symread_1from_1string
__typedArg1 := tmp8396
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8398 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1byte, tmp8397)
}
__typedArg0 := symread_1byte
__typedArg1 := tmp8397
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8399 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1file_1as_1string, tmp8398)
}
__typedArg0 := symread_1file_1as_1string
__typedArg1 := tmp8398
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8400 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1file_1as_1bytelist, tmp8399)
}
__typedArg0 := symread_1file_1as_1bytelist
__typedArg1 := tmp8399
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8401 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1file, tmp8400)
}
__typedArg0 := symread_1file
__typedArg1 := tmp8400
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8402 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symreceive, tmp8401)
}
__typedArg0 := symreceive
__typedArg1 := tmp8401
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8403 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread, tmp8402)
}
__typedArg0 := symread
__typedArg1 := tmp8402
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8404 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symrelease, tmp8403)
}
__typedArg0 := symrelease
__typedArg1 := tmp8403
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8405 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symremove, tmp8404)
}
__typedArg0 := symremove
__typedArg1 := tmp8404
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8406 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symretract, tmp8405)
}
__typedArg0 := symretract
__typedArg1 := tmp8405
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8407 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symreverse, tmp8406)
}
__typedArg0 := symreverse
__typedArg1 := tmp8406
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8408 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symrun, tmp8407)
}
__typedArg0 := symrun
__typedArg1 := tmp8407
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8409 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstr, tmp8408)
}
__typedArg0 := symstr
__typedArg1 := tmp8408
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8410 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsave, tmp8409)
}
__typedArg0 := symsave
__typedArg1 := tmp8409
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8411 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symset, tmp8410)
}
__typedArg0 := symset
__typedArg1 := tmp8410
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8412 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsimple_1error, tmp8411)
}
__typedArg0 := symsimple_1error
__typedArg1 := tmp8411
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8413 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsnd, tmp8412)
}
__typedArg0 := symsnd
__typedArg1 := tmp8412
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8414 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symspecialise, tmp8413)
}
__typedArg0 := symspecialise
__typedArg1 := tmp8413
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8415 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symspy, tmp8414)
}
__typedArg0 := symspy
__typedArg1 := tmp8414
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8416 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsqts, tmp8415)
}
__typedArg0 := symsqts
__typedArg1 := tmp8415
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8417 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstep, tmp8416)
}
__typedArg0 := symstep
__typedArg1 := tmp8416
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8418 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstoutput, tmp8417)
}
__typedArg0 := symstoutput
__typedArg1 := tmp8417
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8419 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstinput, tmp8418)
}
__typedArg0 := symstinput
__typedArg1 := tmp8418
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8420 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp8419)
}
__typedArg0 := symstring
__typedArg1 := tmp8419
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8421 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp8420)
}
__typedArg0 := symstream
__typedArg1 := tmp8420
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8422 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring_1_6n, tmp8421)
}
__typedArg0 := symstring_1_6n
__typedArg1 := tmp8421
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8423 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring_2, tmp8422)
}
__typedArg0 := symstring_2
__typedArg1 := tmp8422
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8424 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsubst, tmp8423)
}
__typedArg0 := symsubst
__typedArg1 := tmp8423
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8425 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsum, tmp8424)
}
__typedArg0 := symsum
__typedArg1 := tmp8424
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8426 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring_1_6symbol, tmp8425)
}
__typedArg0 := symstring_1_6symbol
__typedArg1 := tmp8425
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8427 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol_2, tmp8426)
}
__typedArg0 := symsymbol_2
__typedArg1 := tmp8426
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8428 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp8427)
}
__typedArg0 := symsymbol
__typedArg1 := tmp8427
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8429 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsystem_1S_2, tmp8428)
}
__typedArg0 := symsystem_1S_2
__typedArg1 := tmp8428
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8430 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsynonyms, tmp8429)
}
__typedArg0 := symsynonyms
__typedArg1 := tmp8429
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8431 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsystemf, tmp8430)
}
__typedArg0 := symsystemf
__typedArg1 := tmp8430
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8432 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtail, tmp8431)
}
__typedArg0 := symtail
__typedArg1 := tmp8431
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8433 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtlv, tmp8432)
}
__typedArg0 := symtlv
__typedArg1 := tmp8432
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8434 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtlstr, tmp8433)
}
__typedArg0 := symtlstr
__typedArg1 := tmp8433
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8435 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtl, tmp8434)
}
__typedArg0 := symtl
__typedArg1 := tmp8434
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8436 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtc, tmp8435)
}
__typedArg0 := symtc
__typedArg1 := tmp8435
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8437 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtc_2, tmp8436)
}
__typedArg0 := symtc_2
__typedArg1 := tmp8436
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8438 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symthaw, tmp8437)
}
__typedArg0 := symthaw
__typedArg1 := tmp8437
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8439 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtime, tmp8438)
}
__typedArg0 := symtime
__typedArg1 := tmp8438
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8440 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtrack, tmp8439)
}
__typedArg0 := symtrack
__typedArg1 := tmp8439
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8441 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtracked, tmp8440)
}
__typedArg0 := symtracked
__typedArg1 := tmp8440
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8442 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtrap_1error, tmp8441)
}
__typedArg0 := symtrap_1error
__typedArg1 := tmp8441
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8443 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(True, tmp8442)
}
__typedArg0 := True
__typedArg1 := tmp8442
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8444 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtuple_2, tmp8443)
}
__typedArg0 := symtuple_2
__typedArg1 := tmp8443
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8445 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtype, tmp8444)
}
__typedArg0 := symtype
__typedArg1 := tmp8444
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8446 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symreturn, tmp8445)
}
__typedArg0 := symreturn
__typedArg1 := tmp8445
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8447 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunabsolute, tmp8446)
}
__typedArg0 := symunabsolute
__typedArg1 := tmp8446
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8448 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symundefmacro, tmp8447)
}
__typedArg0 := symundefmacro
__typedArg1 := tmp8447
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8449 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunprofile, tmp8448)
}
__typedArg0 := symunprofile
__typedArg1 := tmp8448
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8450 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunput, tmp8449)
}
__typedArg0 := symunput
__typedArg1 := tmp8449
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8451 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunion, tmp8450)
}
__typedArg0 := symunion
__typedArg1 := tmp8450
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8452 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunix, tmp8451)
}
__typedArg0 := symunix
__typedArg1 := tmp8451
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8453 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunit, tmp8452)
}
__typedArg0 := symunit
__typedArg1 := tmp8452
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8454 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symuntrack, tmp8453)
}
__typedArg0 := symuntrack
__typedArg1 := tmp8453
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8455 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunspecialise, tmp8454)
}
__typedArg0 := symunspecialise
__typedArg1 := tmp8454
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8456 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symupdate_1lambda_1table, tmp8455)
}
__typedArg0 := symupdate_1lambda_1table
__typedArg1 := tmp8455
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8457 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symu_b, tmp8456)
}
__typedArg0 := symu_b
__typedArg1 := tmp8456
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8458 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symuserdefs, tmp8457)
}
__typedArg0 := symuserdefs
__typedArg1 := tmp8457
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8459 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector_2, tmp8458)
}
__typedArg0 := symvector_2
__typedArg1 := tmp8458
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8460 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp8459)
}
__typedArg0 := symvector
__typedArg1 := tmp8459
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8461 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1vector, tmp8460)
}
__typedArg0 := sym_5_1vector
__typedArg1 := tmp8460
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8462 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector_1_6, tmp8461)
}
__typedArg0 := symvector_1_6
__typedArg1 := tmp8461
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8463 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvalue, tmp8462)
}
__typedArg0 := symvalue
__typedArg1 := tmp8462
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8464 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvar_2, tmp8463)
}
__typedArg0 := symvar_2
__typedArg1 := tmp8463
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8465 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvariable_2, tmp8464)
}
__typedArg0 := symvariable_2
__typedArg1 := tmp8464
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8466 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symverified, tmp8465)
}
__typedArg0 := symverified
__typedArg1 := tmp8465
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8467 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symversion, tmp8466)
}
__typedArg0 := symversion
__typedArg1 := tmp8466
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8468 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhen, tmp8467)
}
__typedArg0 := symwhen
__typedArg1 := tmp8467
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8469 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhere, tmp8468)
}
__typedArg0 := symwhere
__typedArg1 := tmp8468
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8470 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwrite_1byte, tmp8469)
}
__typedArg0 := symwrite_1byte
__typedArg1 := tmp8469
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8471 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwrite_1to_1file, tmp8470)
}
__typedArg0 := symwrite_1to_1file
__typedArg1 := tmp8470
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8472 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symy_1or_1n_2, tmp8471)
}
__typedArg0 := symy_1or_1n_2
__typedArg1 := tmp8471
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8473 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8248, tmp8472)
}
__typedArg0 := tmp8248
__typedArg1 := tmp8472
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8474 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_6_6, tmp8473)
}
__typedArg0 := sym_6_6
__typedArg1 := tmp8473
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8475 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5, tmp8474)
}
__typedArg0 := sym_5
__typedArg1 := tmp8474
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8476 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_a, tmp8475)
}
__typedArg0 := sym_5_a
__typedArg1 := tmp8475
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8477 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_7, tmp8476)
}
__typedArg0 := sym_7
__typedArg1 := tmp8476
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8478 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_d, tmp8477)
}
__typedArg0 := sym_d
__typedArg1 := tmp8477
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8479 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_c, tmp8478)
}
__typedArg0 := sym_c
__typedArg1 := tmp8478
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8480 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1, tmp8479)
}
__typedArg0 := sym_1
__typedArg1 := tmp8479
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8481 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_3, tmp8480)
}
__typedArg0 := sym_3
__typedArg1 := tmp8480
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8482 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5end_6, tmp8481)
}
__typedArg0 := sym_5end_6
__typedArg1 := tmp8481
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8483 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_b_6, tmp8482)
}
__typedArg0 := sym_5_b_6
__typedArg1 := tmp8482
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8484 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_c_4, tmp8483)
}
__typedArg0 := sym_c_4
__typedArg1 := tmp8483
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8485 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_a_a_6, tmp8484)
}
__typedArg0 := sym_a_a_6
__typedArg1 := tmp8484
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8486 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_6, tmp8485)
}
__typedArg0 := sym_6
__typedArg1 := tmp8485
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8487 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_6_a, tmp8486)
}
__typedArg0 := sym_6_a
__typedArg1 := tmp8486
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8488 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_a, tmp8487)
}
__typedArg0 := sym_a
__typedArg1 := tmp8487
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8489 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_a_a, tmp8488)
}
__typedArg0 := sym_a_a
__typedArg1 := tmp8488
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8490 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5e_6, tmp8489)
}
__typedArg0 := sym_5e_6
__typedArg1 := tmp8489
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8491 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_6, tmp8490)
}
__typedArg0 := sym_1_6
__typedArg1 := tmp8490
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8492 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1, tmp8491)
}
__typedArg0 := sym_5_1
__typedArg1 := tmp8491
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8493 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dhush_d, tmp8492)
}
__typedArg0 := sym_dhush_d
__typedArg1 := tmp8492
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8494 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dporters_d, tmp8493)
}
__typedArg0 := sym_dporters_d
__typedArg1 := tmp8493
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8495 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dport_d, tmp8494)
}
__typedArg0 := sym_dport_d
__typedArg1 := tmp8494
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8496 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8s, tmp8495)
}
__typedArg0 := sym_8s
__typedArg1 := tmp8495
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8497 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8p, tmp8496)
}
__typedArg0 := sym_8p
__typedArg1 := tmp8496
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8498 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8v, tmp8497)
}
__typedArg0 := sym_8v
__typedArg1 := tmp8497
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8499 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dproperty_1vector_d, tmp8498)
}
__typedArg0 := sym_dproperty_1vector_d
__typedArg1 := tmp8498
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8500 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_drelease_d, tmp8499)
}
__typedArg0 := sym_drelease_d
__typedArg1 := tmp8499
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8501 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dos_d, tmp8500)
}
__typedArg0 := sym_dos_d
__typedArg1 := tmp8500
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8502 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dmacros_d, tmp8501)
}
__typedArg0 := sym_dmacros_d
__typedArg1 := tmp8501
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8503 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dmaximum_1print_1sequence_1size_d, tmp8502)
}
__typedArg0 := sym_dmaximum_1print_1sequence_1size_d
__typedArg1 := tmp8502
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8504 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dversion_d, tmp8503)
}
__typedArg0 := sym_dversion_d
__typedArg1 := tmp8503
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8505 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dhome_1directory_d, tmp8504)
}
__typedArg0 := sym_dhome_1directory_d
__typedArg1 := tmp8504
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8506 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dstoutput_d, tmp8505)
}
__typedArg0 := sym_dstoutput_d
__typedArg1 := tmp8505
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8507 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dstinput_d, tmp8506)
}
__typedArg0 := sym_dstinput_d
__typedArg1 := tmp8506
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8508 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dimplementation_d, tmp8507)
}
__typedArg0 := sym_dimplementation_d
__typedArg1 := tmp8507
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8509 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dlanguage_d, tmp8508)
}
__typedArg0 := sym_dlanguage_d
__typedArg1 := tmp8508
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8510 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym__, tmp8509)
}
__typedArg0 := sym__
__typedArg1 := tmp8509
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8511 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8247, tmp8510)
}
__typedArg0 := tmp8247
__typedArg1 := tmp8510
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8512 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8246, tmp8511)
}
__typedArg0 := tmp8246
__typedArg1 := tmp8511
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8513 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8245, tmp8512)
}
__typedArg0 := tmp8245
__typedArg1 := tmp8512
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8514 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8244, tmp8513)
}
__typedArg0 := tmp8244
__typedArg1 := tmp8513
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8515 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_e_e, tmp8514)
}
__typedArg0 := sym_e_e
__typedArg1 := tmp8514
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8516 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1_1, tmp8515)
}
__typedArg0 := sym_5_1_1
__typedArg1 := tmp8515
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8517 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp8516)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp8516
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8518 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_i, tmp8517)
}
__typedArg0 := sym_i
__typedArg1 := tmp8517
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8519 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_j, tmp8518)
}
__typedArg0 := sym_j
__typedArg1 := tmp8518
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8520 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_b, tmp8519)
}
__typedArg0 := sym_b
__typedArg1 := tmp8519
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8521 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp8522 := Call(__e, PrimFunc(symput), symshen, symshen_4external_1symbols, tmp8520, tmp8521)


_ = tmp8522

tmp8523 := MakeNative(func(__e *ControlFlow) {
V5781 := __e.Get(1)
_ = V5781
tmp8524 := MakeNative(func(__e *ControlFlow) {
W5782 := __e.Get(1)
_ = W5782
tmp8532 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W5782, MakeNumber(-1))
}
__typedArg0 := W5782
__typedArg1 := MakeNumber(-1)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8529 Obj

if True == tmp8532 {
ifres8529 = True


} else {
tmp8531 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W5782, MakeNumber(0))
}
__typedArg0 := W5782
__typedArg1 := MakeNumber(0)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8530 Obj

if True == tmp8531 {
ifres8530 = True


} else {
ifres8530 = False


}

ifres8529 = ifres8530


}

if True == ifres8529 {
__e.Return(Nil)
return
} else {
tmp8525 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5781, Nil)
}
__typedArg0 := V5781
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8526 := Call(__e, PrimFunc(symshen_4lambda_1function), tmp8525, W5782)


tmp8527 := Call(__e, PrimFunc(symeval_1kl), tmp8526)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5781, tmp8527)
}
__typedArg0 := V5781
__typedArg1 := tmp8527
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}


}, 1)

tmp8533 := Call(__e, PrimFunc(symarity), V5781)


__e.TailApply(tmp8524, tmp8533)
return


}, 1)

tmp8534 := Call(__e, ns2_1set, symshen_4lambda_1entry, tmp8523)


_ = tmp8534

tmp8535 := MakeNative(func(__e *ControlFlow) {
V5783 := __e.Get(1)
_ = V5783
tmp8536 := MakeNative(func(__e *ControlFlow) {
W5784 := __e.Get(1)
_ = W5784
tmp8537 := MakeNative(func(__e *ControlFlow) {
Z5786 := __e.Get(1)
_ = Z5786
__e.TailApply(PrimFunc(symshen_4tuple), Z5786)
return
}, 1)

tmp8538 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4tuple, tmp8537)
}
__typedArg0 := symshen_4tuple
__typedArg1 := tmp8537
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8539 := MakeNative(func(__e *ControlFlow) {
Z5787 := __e.Get(1)
_ = Z5787
__e.TailApply(PrimFunc(symshen_4pvar), Z5787)
return
}, 1)

tmp8540 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4pvar, tmp8539)
}
__typedArg0 := symshen_4pvar
__typedArg1 := tmp8539
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8541 := MakeNative(func(__e *ControlFlow) {
Z5788 := __e.Get(1)
_ = Z5788
__e.TailApply(PrimFunc(symshen_4print_1prolog_1vector), Z5788)
return
}, 1)

tmp8542 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4print_1prolog_1vector, tmp8541)
}
__typedArg0 := symshen_4print_1prolog_1vector
__typedArg1 := tmp8541
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8543 := MakeNative(func(__e *ControlFlow) {
Z5789 := __e.Get(1)
_ = Z5789
__e.TailApply(PrimFunc(symshen_4print_1freshterm), Z5789)
return
}, 1)

tmp8544 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4print_1freshterm, tmp8543)
}
__typedArg0 := symshen_4print_1freshterm
__typedArg1 := tmp8543
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8545 := MakeNative(func(__e *ControlFlow) {
Z5790 := __e.Get(1)
_ = Z5790
__e.TailApply(PrimFunc(symshen_4printF), Z5790)
return
}, 1)

tmp8546 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4printF, tmp8545)
}
__typedArg0 := symshen_4printF
__typedArg1 := tmp8545
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8547 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8546, W5784)
}
__typedArg0 := tmp8546
__typedArg1 := W5784
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8548 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8544, tmp8547)
}
__typedArg0 := tmp8544
__typedArg1 := tmp8547
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8549 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8542, tmp8548)
}
__typedArg0 := tmp8542
__typedArg1 := tmp8548
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8550 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8540, tmp8549)
}
__typedArg0 := tmp8540
__typedArg1 := tmp8549
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8551 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8538, tmp8550)
}
__typedArg0 := tmp8538
__typedArg1 := tmp8550
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dlambdatable_d, tmp8551)
}
__typedArg0 := symshen_4_dlambdatable_d
__typedArg1 := tmp8551
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp8552 := MakeNative(func(__e *ControlFlow) {
Z5785 := __e.Get(1)
_ = Z5785
__e.TailApply(PrimFunc(symshen_4lambda_1entry), Z5785)
return
}, 1)

tmp8553 := Call(__e, PrimFunc(symmap), tmp8552, V5783)


__e.TailApply(tmp8536, tmp8553)
return


}, 1)

tmp8554 := Call(__e, ns2_1set, symshen_4build_1lambda_1table, tmp8535)


_ = tmp8554

tmp8555 := Call(__e, PrimFunc(symexternal), symshen)


__e.TailApply(PrimFunc(symshen_4build_1lambda_1table), tmp8555)
return




}, 0)

