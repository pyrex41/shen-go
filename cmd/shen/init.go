package main

import . "github.com/pyrex41/shen-go/kl"

var InitMain = MakeNative(func(__e *ControlFlow) {
tmp19016 := MakeNative(func(__e *ControlFlow) {
tmp19017 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dhistory_d, Nil)
}
__typedArg0 := symshen_4_dhistory_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19017

tmp19018 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dtc_d, False)
}
__typedArg0 := symshen_4_dtc_d
__typedArg1 := False
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19018

tmp19019 := Call(__e, PrimFunc(symshen_4dict), MakeInteger(20000))


tmp19020 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(sym_dproperty_1vector_d, tmp19019)
}
__typedArg0 := sym_dproperty_1vector_d
__typedArg1 := tmp19019
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19020

tmp19021 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
__e.TailApply(PrimFunc(symshen_4macros), X)
return
}, 1)

tmp19022 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4macros, tmp19021)
}
__typedArg0 := symshen_4macros
__typedArg1 := tmp19021
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19023 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19022, Nil)
}
__typedArg0 := tmp19022
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19024 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(sym_dmacros_d, tmp19023)
}
__typedArg0 := sym_dmacros_d
__typedArg1 := tmp19023
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19024

tmp19025 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsource_1form_1handlers_d, Nil)
}
__typedArg0 := symshen_4_dsource_1form_1handlers_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19025

tmp19026 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dgensym_d, MakeInteger(0))
}
__typedArg0 := symshen_4_dgensym_d
__typedArg1 := MakeInteger(0)
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19026

tmp19027 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dtracking_d, Nil)
}
__typedArg0 := symshen_4_dtracking_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19027

tmp19028 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dprofiled_d, Nil)
}
__typedArg0 := symshen_4_dprofiled_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19028

tmp19029 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtype, Nil)
}
__typedArg0 := symtype
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19030 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4input_1h_7, tmp19029)
}
__typedArg0 := symshen_4input_1h_7
__typedArg1 := tmp19029
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19031 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symopen, tmp19030)
}
__typedArg0 := symopen
__typedArg1 := tmp19030
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19032 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symset, tmp19031)
}
__typedArg0 := symset
__typedArg1 := tmp19031
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19033 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhere, tmp19032)
}
__typedArg0 := symwhere
__typedArg1 := tmp19032
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19034 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp19033)
}
__typedArg0 := symlet
__typedArg1 := tmp19033
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19035 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp19034)
}
__typedArg0 := symlambda
__typedArg1 := tmp19034
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19036 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp19035)
}
__typedArg0 := symcons
__typedArg1 := tmp19035
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19037 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8v, tmp19036)
}
__typedArg0 := sym_8v
__typedArg1 := tmp19036
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19038 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8s, tmp19037)
}
__typedArg0 := sym_8s
__typedArg1 := tmp19037
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19039 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8p, tmp19038)
}
__typedArg0 := sym_8p
__typedArg1 := tmp19038
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19040 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dspecial_d, tmp19039)
}
__typedArg0 := symshen_4_dspecial_d
__typedArg1 := tmp19039
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19040

tmp19041 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dextraspecial_d, Nil)
}
__typedArg0 := symshen_4_dextraspecial_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19041

tmp19042 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dspy_d, False)
}
__typedArg0 := symshen_4_dspy_d
__typedArg1 := False
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19042

tmp19043 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_ddatatypes_d, Nil)
}
__typedArg0 := symshen_4_ddatatypes_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19043

tmp19044 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dalldatatypes_d, Nil)
}
__typedArg0 := symshen_4_dalldatatypes_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19044

tmp19045 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dshen_1type_1theory_1enabled_2_d, True)
}
__typedArg0 := symshen_4_dshen_1type_1theory_1enabled_2_d
__typedArg1 := True
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19045

tmp19046 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dpackage_d, symnull)
}
__typedArg0 := symshen_4_dpackage_d
__typedArg1 := symnull
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19046

tmp19047 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsynonyms_d, Nil)
}
__typedArg0 := symshen_4_dsynonyms_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19047

tmp19048 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_doccurs_d, True)
}
__typedArg0 := symshen_4_doccurs_d
__typedArg1 := True
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19048

tmp19049 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dfactorise_2_d, False)
}
__typedArg0 := symshen_4_dfactorise_2_d
__typedArg1 := False
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19049

tmp19050 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dmaxinferences_d, MakeInteger(1000000))
}
__typedArg0 := symshen_4_dmaxinferences_d
__typedArg1 := MakeInteger(1000000)
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19050

tmp19051 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(sym_dmaximum_1print_1sequence_1size_d, MakeInteger(20))
}
__typedArg0 := sym_dmaximum_1print_1sequence_1size_d
__typedArg1 := MakeInteger(20)
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19051

tmp19052 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dcall_d, MakeInteger(0))
}
__typedArg0 := symshen_4_dcall_d
__typedArg1 := MakeInteger(0)
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19052

tmp19053 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dinfs_d, MakeInteger(0))
}
__typedArg0 := symshen_4_dinfs_d
__typedArg1 := MakeInteger(0)
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19053

tmp19054 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(sym_dhush_d, False)
}
__typedArg0 := sym_dhush_d
__typedArg1 := False
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19054

tmp19055 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_doptimise_d, False)
}
__typedArg0 := symshen_4_doptimise_d
__typedArg1 := False
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19055

tmp19056 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(sym_dversion_d, MakeString("43.4"))
}
__typedArg0 := sym_dversion_d
__typedArg1 := MakeString("43.4")
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19056

tmp19057 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dnames_d, Nil)
}
__typedArg0 := symshen_4_dnames_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19057

tmp19058 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dstep_d, False)
}
__typedArg0 := symshen_4_dstep_d
__typedArg1 := False
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19058

tmp19059 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dit_d, MakeString(""))
}
__typedArg0 := symshen_4_dit_d
__typedArg1 := MakeString("")
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19059

tmp19060 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dresidue_d, Nil)
}
__typedArg0 := symshen_4_dresidue_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19060

tmp19061 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(sym_dabsolute_d, Nil)
}
__typedArg0 := sym_dabsolute_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19061

tmp19062 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dprolog_1memory_d, MakeInteger(1000))
}
__typedArg0 := symshen_4_dprolog_1memory_d
__typedArg1 := MakeInteger(1000)
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19062

tmp19063 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dloading_2_d, False)
}
__typedArg0 := symshen_4_dloading_2_d
__typedArg1 := False
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19063

tmp19064 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_duserdefs_d, Nil)
}
__typedArg0 := symshen_4_duserdefs_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19064

tmp19065 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
__e.Return(X)
return
}, 1)

tmp19066 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_ddemodulation_1function_d, tmp19065)
}
__typedArg0 := symshen_4_ddemodulation_1function_d
__typedArg1 := tmp19065
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19066

tmp19067 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dcustom_1pattern_1compiler_d, False)
}
__typedArg0 := symshen_4_dcustom_1pattern_1compiler_d
__typedArg1 := False
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19067

tmp19068 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dcustom_1pattern_1reducer_d, False)
}
__typedArg0 := symshen_4_dcustom_1pattern_1reducer_d
__typedArg1 := False
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19068

tmp19071 := Call(__e, PrimFunc(symbound_2), sym_dhome_1directory_d)


tmp19072 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp19071)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp19071
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres19069 Obj

if True == tmp19072 {
tmp19070 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(sym_dhome_1directory_d, MakeString(""))
}
__typedArg0 := sym_dhome_1directory_d
__typedArg1 := MakeString("")
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

ifres19069 = tmp19070


} else {
ifres19069 = symshen_4skip


}

_ = ifres19069

tmp19076 := Call(__e, PrimFunc(symbound_2), sym_dsterror_d)


tmp19077 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp19076)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp19076
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres19073 Obj

if True == tmp19077 {
tmp19074 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dstoutput_d)
}
__typedArg0 := sym_dstoutput_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp19075 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(sym_dsterror_d, tmp19074)
}
__typedArg0 := sym_dsterror_d
__typedArg1 := tmp19074
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

ifres19073 = tmp19075


} else {
ifres19073 = symshen_4skip


}

_ = ifres19073

tmp19078 := Call(__e, PrimFunc(symprolog_1memory), MakeInteger(10000))


_ = tmp19078

tmp19079 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dloading_2_d, False)
}
__typedArg0 := symshen_4_dloading_2_d
__typedArg1 := False
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19079

tmp19080 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), Nil)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19081 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8s, tmp19080)
}
__typedArg0 := sym_8s
__typedArg1 := tmp19080
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19082 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19081)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19081
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19083 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8v, tmp19082)
}
__typedArg0 := sym_8v
__typedArg1 := tmp19082
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19084 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19083)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19083
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19085 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8p, tmp19084)
}
__typedArg0 := sym_8p
__typedArg1 := tmp19084
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19086 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19085)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19085
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19087 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_b_6, tmp19086)
}
__typedArg0 := sym_5_b_6
__typedArg1 := tmp19086
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19088 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19087)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19087
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19089 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5end_6, tmp19088)
}
__typedArg0 := sym_5end_6
__typedArg1 := tmp19088
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19090 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19089)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19089
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19091 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5e_6, tmp19090)
}
__typedArg0 := sym_5e_6
__typedArg1 := tmp19090
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19092 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19091)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19091
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19093 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_a_a, tmp19092)
}
__typedArg0 := sym_a_a
__typedArg1 := tmp19092
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19094 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19093)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19093
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19095 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1, tmp19094)
}
__typedArg0 := sym_1
__typedArg1 := tmp19094
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19096 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19095)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19095
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19097 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_c, tmp19096)
}
__typedArg0 := sym_c
__typedArg1 := tmp19096
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19098 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19097)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19097
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19099 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_d, tmp19098)
}
__typedArg0 := sym_d
__typedArg1 := tmp19098
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19100 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19099)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19099
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19101 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_7, tmp19100)
}
__typedArg0 := sym_7
__typedArg1 := tmp19100
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19102 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19101)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19101
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19103 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symy_1or_1n_2, tmp19102)
}
__typedArg0 := symy_1or_1n_2
__typedArg1 := tmp19102
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19104 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19103)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19103
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19105 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwrite_1to_1file, tmp19104)
}
__typedArg0 := symwrite_1to_1file
__typedArg1 := tmp19104
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19106 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19105)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19105
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19107 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwrite_1string, tmp19106)
}
__typedArg0 := symwrite_1string
__typedArg1 := tmp19106
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19108 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19107)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19107
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19109 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwrite_1byte, tmp19108)
}
__typedArg0 := symwrite_1byte
__typedArg1 := tmp19108
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19110 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(5), tmp19109)
}
__typedArg0 := MakeInteger(5)
__typedArg1 := tmp19109
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19111 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhen, tmp19110)
}
__typedArg0 := symwhen
__typedArg1 := tmp19110
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19112 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19111)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19111
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19113 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symversion, tmp19112)
}
__typedArg0 := symversion
__typedArg1 := tmp19112
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19114 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(5), tmp19113)
}
__typedArg0 := MakeInteger(5)
__typedArg1 := tmp19113
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19115 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvar_2, tmp19114)
}
__typedArg0 := symvar_2
__typedArg1 := tmp19114
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19116 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19115)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19115
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19117 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvariable_2, tmp19116)
}
__typedArg0 := symvariable_2
__typedArg1 := tmp19116
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19118 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19117)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19117
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19119 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvalue, tmp19118)
}
__typedArg0 := symvalue
__typedArg1 := tmp19118
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19120 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(3), tmp19119)
}
__typedArg0 := MakeInteger(3)
__typedArg1 := tmp19119
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19121 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector_1_6, tmp19120)
}
__typedArg0 := symvector_1_6
__typedArg1 := tmp19120
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19122 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19121)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19121
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19123 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector_2, tmp19122)
}
__typedArg0 := symvector_2
__typedArg1 := tmp19122
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19124 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19123)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19123
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19125 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp19124)
}
__typedArg0 := symvector
__typedArg1 := tmp19124
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19126 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19125)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19125
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19127 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symuserdefs, tmp19126)
}
__typedArg0 := symuserdefs
__typedArg1 := tmp19126
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19128 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19127)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19127
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19129 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symupdate_1lambda_1table, tmp19128)
}
__typedArg0 := symupdate_1lambda_1table
__typedArg1 := tmp19128
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19130 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19129)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19129
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19131 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symundefmacro, tmp19130)
}
__typedArg0 := symundefmacro
__typedArg1 := tmp19130
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19132 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19131)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19131
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19133 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symuntrack, tmp19132)
}
__typedArg0 := symuntrack
__typedArg1 := tmp19132
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19134 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19133)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19133
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19135 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunion, tmp19134)
}
__typedArg0 := symunion
__typedArg1 := tmp19134
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19136 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19135)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19135
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19137 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunprofile, tmp19136)
}
__typedArg0 := symunprofile
__typedArg1 := tmp19136
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19138 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(3), tmp19137)
}
__typedArg0 := MakeInteger(3)
__typedArg1 := tmp19137
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19139 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunput, tmp19138)
}
__typedArg0 := symunput
__typedArg1 := tmp19138
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19140 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19139)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19139
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19141 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symundefmacro, tmp19140)
}
__typedArg0 := symundefmacro
__typedArg1 := tmp19140
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19142 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19141)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19141
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19143 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunabsolute, tmp19142)
}
__typedArg0 := symunabsolute
__typedArg1 := tmp19142
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19144 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(5), tmp19143)
}
__typedArg0 := MakeInteger(5)
__typedArg1 := tmp19143
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19145 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symreturn, tmp19144)
}
__typedArg0 := symreturn
__typedArg1 := tmp19144
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19146 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19145)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19145
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19147 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtype, tmp19146)
}
__typedArg0 := symtype
__typedArg1 := tmp19146
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19148 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19147)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19147
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19149 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtuple_2, tmp19148)
}
__typedArg0 := symtuple_2
__typedArg1 := tmp19148
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19150 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19149)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19149
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19151 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtrap_1error, tmp19150)
}
__typedArg0 := symtrap_1error
__typedArg1 := tmp19150
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19152 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19151)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19151
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19153 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtracked, tmp19152)
}
__typedArg0 := symtracked
__typedArg1 := tmp19152
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19154 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19153)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19153
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19155 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtrack, tmp19154)
}
__typedArg0 := symtrack
__typedArg1 := tmp19154
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19156 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19155)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19155
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19157 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtlstr, tmp19156)
}
__typedArg0 := symtlstr
__typedArg1 := tmp19156
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19158 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19157)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19157
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19159 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symthaw, tmp19158)
}
__typedArg0 := symthaw
__typedArg1 := tmp19158
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19160 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19159)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19159
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19161 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtc_2, tmp19160)
}
__typedArg0 := symtc_2
__typedArg1 := tmp19160
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19162 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19161)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19161
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19163 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtc, tmp19162)
}
__typedArg0 := symtc
__typedArg1 := tmp19162
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19164 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19163)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19163
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19165 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtl, tmp19164)
}
__typedArg0 := symtl
__typedArg1 := tmp19164
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19166 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19165)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19165
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19167 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtail, tmp19166)
}
__typedArg0 := symtail
__typedArg1 := tmp19166
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19168 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19167)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19167
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19169 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsystemf, tmp19168)
}
__typedArg0 := symsystemf
__typedArg1 := tmp19168
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19170 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19169)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19169
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19171 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol_2, tmp19170)
}
__typedArg0 := symsymbol_2
__typedArg1 := tmp19170
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19172 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19171)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19171
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19173 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsum, tmp19172)
}
__typedArg0 := symsum
__typedArg1 := tmp19172
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19174 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(3), tmp19173)
}
__typedArg0 := MakeInteger(3)
__typedArg1 := tmp19173
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19175 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsubst, tmp19174)
}
__typedArg0 := symsubst
__typedArg1 := tmp19174
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19176 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19175)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19175
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19177 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring_2, tmp19176)
}
__typedArg0 := symstring_2
__typedArg1 := tmp19176
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19178 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19177)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19177
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19179 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring_1_6symbol, tmp19178)
}
__typedArg0 := symstring_1_6symbol
__typedArg1 := tmp19178
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19180 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19179)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19179
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19181 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring_1_6n, tmp19180)
}
__typedArg0 := symstring_1_6n
__typedArg1 := tmp19180
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19182 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19181)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19181
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19183 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstr, tmp19182)
}
__typedArg0 := symstr
__typedArg1 := tmp19182
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19184 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19183)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19183
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19185 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstoutput, tmp19184)
}
__typedArg0 := symstoutput
__typedArg1 := tmp19184
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19186 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19185)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19185
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19187 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstinput, tmp19186)
}
__typedArg0 := symstinput
__typedArg1 := tmp19186
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19188 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19187)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19187
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19189 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstep_2, tmp19188)
}
__typedArg0 := symstep_2
__typedArg1 := tmp19188
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19190 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19189)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19189
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19191 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstep, tmp19190)
}
__typedArg0 := symstep
__typedArg1 := tmp19190
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19192 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19191)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19191
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19193 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symspy_2, tmp19192)
}
__typedArg0 := symspy_2
__typedArg1 := tmp19192
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19194 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19193)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19193
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19195 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symspy, tmp19194)
}
__typedArg0 := symspy
__typedArg1 := tmp19194
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19196 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19195)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19195
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19197 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symspecialise, tmp19196)
}
__typedArg0 := symspecialise
__typedArg1 := tmp19196
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19198 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19197)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19197
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19199 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsnd, tmp19198)
}
__typedArg0 := symsnd
__typedArg1 := tmp19198
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19200 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19199)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19199
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19201 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsimple_1error, tmp19200)
}
__typedArg0 := symsimple_1error
__typedArg1 := tmp19200
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19202 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19201)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19201
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19203 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symset, tmp19202)
}
__typedArg0 := symset
__typedArg1 := tmp19202
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19204 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19203)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19203
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19205 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symreverse, tmp19204)
}
__typedArg0 := symreverse
__typedArg1 := tmp19204
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19206 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19205)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19205
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19207 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symremove, tmp19206)
}
__typedArg0 := symremove
__typedArg1 := tmp19206
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19208 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19207)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19207
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19209 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4unregister_1source_1form, tmp19208)
}
__typedArg0 := symshen_4unregister_1source_1form
__typedArg1 := tmp19208
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19210 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19209)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19209
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19211 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4register_1source_1form, tmp19210)
}
__typedArg0 := symshen_4register_1source_1form
__typedArg1 := tmp19210
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19212 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19211)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19211
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19213 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symrelease, tmp19212)
}
__typedArg0 := symrelease
__typedArg1 := tmp19212
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19214 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19213)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19213
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19215 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symreceive, tmp19214)
}
__typedArg0 := symreceive
__typedArg1 := tmp19214
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19216 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19215)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19215
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19217 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1unit_1string, tmp19216)
}
__typedArg0 := symread_1unit_1string
__typedArg1 := tmp19216
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19218 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19217)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19217
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19219 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1from_1string_1unprocessed, tmp19218)
}
__typedArg0 := symread_1from_1string_1unprocessed
__typedArg1 := tmp19218
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19220 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19219)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19219
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19221 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1from_1string, tmp19220)
}
__typedArg0 := symread_1from_1string
__typedArg1 := tmp19220
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19222 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19221)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19221
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19223 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1byte, tmp19222)
}
__typedArg0 := symread_1byte
__typedArg1 := tmp19222
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19224 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19223)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19223
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19225 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread, tmp19224)
}
__typedArg0 := symread
__typedArg1 := tmp19224
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19226 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19225)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19225
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19227 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1file, tmp19226)
}
__typedArg0 := symread_1file
__typedArg1 := tmp19226
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19228 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19227)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19227
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19229 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1file_1as_1bytelist, tmp19228)
}
__typedArg0 := symread_1file_1as_1bytelist
__typedArg1 := tmp19228
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19230 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19229)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19229
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19231 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1file_1as_1string, tmp19230)
}
__typedArg0 := symread_1file_1as_1string
__typedArg1 := tmp19230
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19232 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(4), tmp19231)
}
__typedArg0 := MakeInteger(4)
__typedArg1 := tmp19231
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19233 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symput, tmp19232)
}
__typedArg0 := symput
__typedArg1 := tmp19232
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19234 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19233)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19233
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19235 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symprotect, tmp19234)
}
__typedArg0 := symprotect
__typedArg1 := tmp19234
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19236 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19235)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19235
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19237 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympreclude_1all_1but, tmp19236)
}
__typedArg0 := sympreclude_1all_1but
__typedArg1 := tmp19236
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19238 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19237)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19237
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19239 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympreclude, tmp19238)
}
__typedArg0 := sympreclude
__typedArg1 := tmp19238
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19240 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19239)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19239
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19241 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symps, tmp19240)
}
__typedArg0 := symps
__typedArg1 := tmp19240
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19242 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19241)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19241
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19243 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympr, tmp19242)
}
__typedArg0 := sympr
__typedArg1 := tmp19242
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19244 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19243)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19243
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19245 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symprofile_1results, tmp19244)
}
__typedArg0 := symprofile_1results
__typedArg1 := tmp19244
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19246 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19245)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19245
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19247 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symprolog_1memory, tmp19246)
}
__typedArg0 := symprolog_1memory
__typedArg1 := tmp19246
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19248 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19247)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19247
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19249 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4printF, tmp19248)
}
__typedArg0 := symshen_4printF
__typedArg1 := tmp19248
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19250 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19249)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19249
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19251 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4print_1freshterm, tmp19250)
}
__typedArg0 := symshen_4print_1freshterm
__typedArg1 := tmp19250
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19252 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19251)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19251
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19253 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4print_1prolog_1vector, tmp19252)
}
__typedArg0 := symshen_4print_1prolog_1vector
__typedArg1 := tmp19252
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19254 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19253)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19253
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19255 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symprofile, tmp19254)
}
__typedArg0 := symprofile
__typedArg1 := tmp19254
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19256 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19255)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19255
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19257 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symprint, tmp19256)
}
__typedArg0 := symprint
__typedArg1 := tmp19256
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19258 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19257)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19257
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19259 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympreclude_1all_1but, tmp19258)
}
__typedArg0 := sympreclude_1all_1but
__typedArg1 := tmp19258
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19260 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19259)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19259
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19261 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympos, tmp19260)
}
__typedArg0 := sympos
__typedArg1 := tmp19260
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19262 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19261)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19261
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19263 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symporters, tmp19262)
}
__typedArg0 := symporters
__typedArg1 := tmp19262
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19264 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19263)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19263
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19265 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symport, tmp19264)
}
__typedArg0 := symport
__typedArg1 := tmp19264
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19266 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19265)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19265
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19267 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympackage_2, tmp19266)
}
__typedArg0 := sympackage_2
__typedArg1 := tmp19266
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19268 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(3), tmp19267)
}
__typedArg0 := MakeInteger(3)
__typedArg1 := tmp19267
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19269 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympackage, tmp19268)
}
__typedArg0 := sympackage
__typedArg1 := tmp19268
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19270 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19269)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19269
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19271 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symos, tmp19270)
}
__typedArg0 := symos
__typedArg1 := tmp19270
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19272 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19271)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19271
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19273 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symor, tmp19272)
}
__typedArg0 := symor
__typedArg1 := tmp19272
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19274 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19273)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19273
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19275 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symoptimise_2, tmp19274)
}
__typedArg0 := symoptimise_2
__typedArg1 := tmp19274
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19276 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19275)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19275
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19277 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symoptimise, tmp19276)
}
__typedArg0 := symoptimise
__typedArg1 := tmp19276
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19278 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19277)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19277
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19279 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symopen, tmp19278)
}
__typedArg0 := symopen
__typedArg1 := tmp19278
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19280 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19279)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19279
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19281 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symoccurs_1check, tmp19280)
}
__typedArg0 := symoccurs_1check
__typedArg1 := tmp19280
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19282 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19281)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19281
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19283 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symoccurs_2, tmp19282)
}
__typedArg0 := symoccurs_2
__typedArg1 := tmp19282
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19284 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19283)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19283
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19285 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symoccurrences, tmp19284)
}
__typedArg0 := symoccurrences
__typedArg1 := tmp19284
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19286 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19285)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19285
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19287 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symoccurs_1check, tmp19286)
}
__typedArg0 := symoccurs_1check
__typedArg1 := tmp19286
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19288 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19287)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19287
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19289 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber_2, tmp19288)
}
__typedArg0 := symnumber_2
__typedArg1 := tmp19288
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19290 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19289)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19289
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19291 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symn_1_6string, tmp19290)
}
__typedArg0 := symn_1_6string
__typedArg1 := tmp19290
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19292 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19291)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19291
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19293 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnth, tmp19292)
}
__typedArg0 := symnth
__typedArg1 := tmp19292
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19294 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19293)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19293
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19295 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnot, tmp19294)
}
__typedArg0 := symnot
__typedArg1 := tmp19294
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19296 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19295)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19295
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19297 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnl, tmp19296)
}
__typedArg0 := symnl
__typedArg1 := tmp19296
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19298 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19297)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19297
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19299 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symmaxinferences, tmp19298)
}
__typedArg0 := symmaxinferences
__typedArg1 := tmp19298
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19300 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19299)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19299
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19301 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symmapcan, tmp19300)
}
__typedArg0 := symmapcan
__typedArg1 := tmp19300
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19302 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19301)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19301
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19303 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symmap, tmp19302)
}
__typedArg0 := symmap
__typedArg1 := tmp19302
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19304 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19303)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19303
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19305 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symmacroexpand, tmp19304)
}
__typedArg0 := symmacroexpand
__typedArg1 := tmp19304
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19306 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19305)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19305
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19307 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp19306)
}
__typedArg0 := symvector
__typedArg1 := tmp19306
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19308 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19307)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19307
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19309 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_a, tmp19308)
}
__typedArg0 := sym_5_a
__typedArg1 := tmp19308
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19310 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19309)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19309
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19311 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5, tmp19310)
}
__typedArg0 := sym_5
__typedArg1 := tmp19310
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19312 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19311)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19311
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19313 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symload, tmp19312)
}
__typedArg0 := symload
__typedArg1 := tmp19312
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19314 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19313)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19313
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19315 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19314)
}
__typedArg0 := symlist
__typedArg1 := tmp19314
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19316 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19315)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19315
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19317 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlineread, tmp19316)
}
__typedArg0 := symlineread
__typedArg1 := tmp19316
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19318 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19317)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19317
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19319 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlimit, tmp19318)
}
__typedArg0 := symlimit
__typedArg1 := tmp19318
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19320 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19319)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19319
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19321 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlength, tmp19320)
}
__typedArg0 := symlength
__typedArg1 := tmp19320
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19322 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19321)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19321
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19323 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlanguage, tmp19322)
}
__typedArg0 := symlanguage
__typedArg1 := tmp19322
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19324 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(6), tmp19323)
}
__typedArg0 := MakeInteger(6)
__typedArg1 := tmp19323
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19325 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symis_b, tmp19324)
}
__typedArg0 := symis_b
__typedArg1 := tmp19324
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19326 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(6), tmp19325)
}
__typedArg0 := MakeInteger(6)
__typedArg1 := tmp19325
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19327 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symis, tmp19326)
}
__typedArg0 := symis
__typedArg1 := tmp19326
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19328 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19327)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19327
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19329 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symit, tmp19328)
}
__typedArg0 := symit
__typedArg1 := tmp19328
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19330 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19329)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19329
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19331 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminternal, tmp19330)
}
__typedArg0 := syminternal
__typedArg1 := tmp19330
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19332 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19331)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19331
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19333 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symintersection, tmp19332)
}
__typedArg0 := symintersection
__typedArg1 := tmp19332
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19334 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19333)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19333
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19335 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminclude_1all_1but, tmp19334)
}
__typedArg0 := syminclude_1all_1but
__typedArg1 := tmp19334
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19336 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19335)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19335
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19337 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symimplementation, tmp19336)
}
__typedArg0 := symimplementation
__typedArg1 := tmp19336
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19338 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19337)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19337
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19339 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminput_7, tmp19338)
}
__typedArg0 := syminput_7
__typedArg1 := tmp19338
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19340 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19339)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19339
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19341 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminput, tmp19340)
}
__typedArg0 := syminput
__typedArg1 := tmp19340
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19342 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19341)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19341
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19343 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminferences, tmp19342)
}
__typedArg0 := syminferences
__typedArg1 := tmp19342
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19344 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19343)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19343
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19345 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symintern, tmp19344)
}
__typedArg0 := symintern
__typedArg1 := tmp19344
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19346 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19345)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19345
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19347 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminternal, tmp19346)
}
__typedArg0 := syminternal
__typedArg1 := tmp19346
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19348 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19347)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19347
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19349 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminteger_2, tmp19348)
}
__typedArg0 := syminteger_2
__typedArg1 := tmp19348
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19350 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19349)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19349
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19351 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symin_1package, tmp19350)
}
__typedArg0 := symin_1package
__typedArg1 := tmp19350
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19352 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19351)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19351
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19353 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symincluded, tmp19352)
}
__typedArg0 := symincluded
__typedArg1 := tmp19352
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19354 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19353)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19353
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19355 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminclude, tmp19354)
}
__typedArg0 := syminclude
__typedArg1 := tmp19354
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19356 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(3), tmp19355)
}
__typedArg0 := MakeInteger(3)
__typedArg1 := tmp19355
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19357 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp19356)
}
__typedArg0 := symif
__typedArg1 := tmp19356
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19358 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19357)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19357
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19359 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhush, tmp19358)
}
__typedArg0 := symhush
__typedArg1 := tmp19358
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19360 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19359)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19359
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19361 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhush_2, tmp19360)
}
__typedArg0 := symhush_2
__typedArg1 := tmp19360
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19362 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19361)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19361
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19363 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhead, tmp19362)
}
__typedArg0 := symhead
__typedArg1 := tmp19362
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19364 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19363)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19363
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19365 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhdstr, tmp19364)
}
__typedArg0 := symhdstr
__typedArg1 := tmp19364
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19366 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19365)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19365
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19367 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhdv, tmp19366)
}
__typedArg0 := symhdv
__typedArg1 := tmp19366
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19368 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19367)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19367
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19369 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhd, tmp19368)
}
__typedArg0 := symhd
__typedArg1 := tmp19368
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19370 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19369)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19369
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19371 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhash, tmp19370)
}
__typedArg0 := symhash
__typedArg1 := tmp19370
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19372 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19371)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19371
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19373 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_a, tmp19372)
}
__typedArg0 := sym_a
__typedArg1 := tmp19372
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19374 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19373)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19373
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19375 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_6_a, tmp19374)
}
__typedArg0 := sym_6_a
__typedArg1 := tmp19374
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19376 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19375)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19375
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19377 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_6, tmp19376)
}
__typedArg0 := sym_6
__typedArg1 := tmp19376
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19378 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19377)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19377
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19379 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1vector, tmp19378)
}
__typedArg0 := sym_5_1vector
__typedArg1 := tmp19378
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19380 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19379)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19379
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19381 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1address, tmp19380)
}
__typedArg0 := sym_5_1address
__typedArg1 := tmp19380
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19382 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(3), tmp19381)
}
__typedArg0 := MakeInteger(3)
__typedArg1 := tmp19381
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19383 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symaddress_1_6, tmp19382)
}
__typedArg0 := symaddress_1_6
__typedArg1 := tmp19382
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19384 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19383)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19383
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19385 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symget_1time, tmp19384)
}
__typedArg0 := symget_1time
__typedArg1 := tmp19384
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19386 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(3), tmp19385)
}
__typedArg0 := MakeInteger(3)
__typedArg1 := tmp19385
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19387 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symget, tmp19386)
}
__typedArg0 := symget
__typedArg1 := tmp19386
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19388 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19387)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19387
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19389 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symgensym, tmp19388)
}
__typedArg0 := symgensym
__typedArg1 := tmp19388
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19390 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19389)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19389
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19391 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfunction, tmp19390)
}
__typedArg0 := symfunction
__typedArg1 := tmp19390
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19392 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19391)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19391
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19393 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfn, tmp19392)
}
__typedArg0 := symfn
__typedArg1 := tmp19392
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19394 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19393)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19393
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19395 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfst, tmp19394)
}
__typedArg0 := symfst
__typedArg1 := tmp19394
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19396 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19395)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19395
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19397 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfresh, tmp19396)
}
__typedArg0 := symfresh
__typedArg1 := tmp19396
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19398 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19397)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19397
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19399 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfreeze, tmp19398)
}
__typedArg0 := symfreeze
__typedArg1 := tmp19398
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19400 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(5), tmp19399)
}
__typedArg0 := MakeInteger(5)
__typedArg1 := tmp19399
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19401 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfork, tmp19400)
}
__typedArg0 := symfork
__typedArg1 := tmp19400
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19402 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19401)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19401
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19403 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symforeign, tmp19402)
}
__typedArg0 := symforeign
__typedArg1 := tmp19402
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19404 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(7), tmp19403)
}
__typedArg0 := MakeInteger(7)
__typedArg1 := tmp19403
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19405 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfindall, tmp19404)
}
__typedArg0 := symfindall
__typedArg1 := tmp19404
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19406 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19405)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19405
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19407 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfix, tmp19406)
}
__typedArg0 := symfix
__typedArg1 := tmp19406
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19408 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19407)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19407
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19409 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfail, tmp19408)
}
__typedArg0 := symfail
__typedArg1 := tmp19408
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19410 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19409)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19409
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19411 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfail_1if, tmp19410)
}
__typedArg0 := symfail_1if
__typedArg1 := tmp19410
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19412 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19411)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19411
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19413 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfactorise_2, tmp19412)
}
__typedArg0 := symfactorise_2
__typedArg1 := tmp19412
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19414 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19413)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19413
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19415 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfactorise, tmp19414)
}
__typedArg0 := symfactorise
__typedArg1 := tmp19414
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19416 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19415)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19415
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19417 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symexternal, tmp19416)
}
__typedArg0 := symexternal
__typedArg1 := tmp19416
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19418 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19417)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19417
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19419 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symexplode, tmp19418)
}
__typedArg0 := symexplode
__typedArg1 := tmp19418
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19420 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19419)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19419
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19421 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symeval_1kl, tmp19420)
}
__typedArg0 := symeval_1kl
__typedArg1 := tmp19420
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19422 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19421)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19421
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19423 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symeval, tmp19422)
}
__typedArg0 := symeval
__typedArg1 := tmp19422
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19424 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19423)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19423
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19425 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symerror_1to_1string, tmp19424)
}
__typedArg0 := symerror_1to_1string
__typedArg1 := tmp19424
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19426 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19425)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19425
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19427 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symexternal, tmp19426)
}
__typedArg0 := symexternal
__typedArg1 := tmp19426
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19428 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19427)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19427
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19429 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symenable_1type_1theory, tmp19428)
}
__typedArg0 := symenable_1type_1theory
__typedArg1 := tmp19428
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19430 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19429)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19429
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19431 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symempty_2, tmp19430)
}
__typedArg0 := symempty_2
__typedArg1 := tmp19430
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19432 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19431)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19431
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19433 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symelement_2, tmp19432)
}
__typedArg0 := symelement_2
__typedArg1 := tmp19432
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19434 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19433)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19433
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19435 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdo, tmp19434)
}
__typedArg0 := symdo
__typedArg1 := tmp19434
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19436 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19435)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19435
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19437 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdifference, tmp19436)
}
__typedArg0 := symdifference
__typedArg1 := tmp19436
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19438 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19437)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19437
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19439 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdestroy, tmp19438)
}
__typedArg0 := symdestroy
__typedArg1 := tmp19438
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19440 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19439)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19439
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19441 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdeclare, tmp19440)
}
__typedArg0 := symdeclare
__typedArg1 := tmp19440
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19442 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19441)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19441
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19443 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdatatypes, tmp19442)
}
__typedArg0 := symdatatypes
__typedArg1 := tmp19442
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19444 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19443)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19443
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19445 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symclose, tmp19444)
}
__typedArg0 := symclose
__typedArg1 := tmp19444
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19446 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19445)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19445
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19447 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcn, tmp19446)
}
__typedArg0 := symcn
__typedArg1 := tmp19446
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19448 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19447)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19447
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19449 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons_2, tmp19448)
}
__typedArg0 := symcons_2
__typedArg1 := tmp19448
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19450 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19449)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19449
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19451 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp19450)
}
__typedArg0 := symcons
__typedArg1 := tmp19450
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19452 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19451)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19451
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19453 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symconcat, tmp19452)
}
__typedArg0 := symconcat
__typedArg1 := tmp19452
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19454 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19453)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19453
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19455 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcompile, tmp19454)
}
__typedArg0 := symcompile
__typedArg1 := tmp19454
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19456 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19455)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19455
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19457 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symchar_1stream_2, tmp19456)
}
__typedArg0 := symchar_1stream_2
__typedArg1 := tmp19456
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19458 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19457)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19457
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19459 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcd, tmp19458)
}
__typedArg0 := symcd
__typedArg1 := tmp19458
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19460 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(5), tmp19459)
}
__typedArg0 := MakeInteger(5)
__typedArg1 := tmp19459
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19461 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcall, tmp19460)
}
__typedArg0 := symcall
__typedArg1 := tmp19460
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19462 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(6), tmp19461)
}
__typedArg0 := MakeInteger(6)
__typedArg1 := tmp19461
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19463 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symbind, tmp19462)
}
__typedArg0 := symbind
__typedArg1 := tmp19462
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19464 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19463)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19463
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19465 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symbound_2, tmp19464)
}
__typedArg0 := symbound_2
__typedArg1 := tmp19464
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19466 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19465)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19465
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19467 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symbootstrap, tmp19466)
}
__typedArg0 := symbootstrap
__typedArg1 := tmp19466
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19468 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19467)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19467
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19469 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean_2, tmp19468)
}
__typedArg0 := symboolean_2
__typedArg1 := tmp19468
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19470 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19469)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19469
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19471 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symatom_2, tmp19470)
}
__typedArg0 := symatom_2
__typedArg1 := tmp19470
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19472 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19471)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19471
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19473 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symassoc, tmp19472)
}
__typedArg0 := symassoc
__typedArg1 := tmp19472
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19474 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19473)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19473
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19475 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symarity, tmp19474)
}
__typedArg0 := symarity
__typedArg1 := tmp19474
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19476 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19475)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19475
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19477 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symappend, tmp19476)
}
__typedArg0 := symappend
__typedArg1 := tmp19476
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19478 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19477)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19477
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19479 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symand, tmp19478)
}
__typedArg0 := symand
__typedArg1 := tmp19478
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19480 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19479)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19479
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19481 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symadjoin, tmp19480)
}
__typedArg0 := symadjoin
__typedArg1 := tmp19480
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19482 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(3), tmp19481)
}
__typedArg0 := MakeInteger(3)
__typedArg1 := tmp19481
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19483 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symaddress_1_6, tmp19482)
}
__typedArg0 := symaddress_1_6
__typedArg1 := tmp19482
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19484 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19483)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19483
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19485 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symabsvector, tmp19484)
}
__typedArg0 := symabsvector
__typedArg1 := tmp19484
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19486 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19485)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19485
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19487 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symabsvector_2, tmp19486)
}
__typedArg0 := symabsvector_2
__typedArg1 := tmp19486
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19488 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19487)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19487
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19489 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symabsolute, tmp19488)
}
__typedArg0 := symabsolute
__typedArg1 := tmp19488
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19490 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19489)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19489
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19491 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symabort, tmp19490)
}
__typedArg0 := symabort
__typedArg1 := tmp19490
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19492 := Call(__e, PrimFunc(symshen_4initialise_1arity_1table), tmp19491)


_ = tmp19492

tmp19493 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp19494 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(";"))
}
__typedArg0 := MakeString(";")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp19495 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":="))
}
__typedArg0 := MakeString(":=")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp19496 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(","))
}
__typedArg0 := MakeString(",")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp19497 := Call(__e, PrimFunc(symvector), MakeInteger(0))


tmp19498 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString("bar!"))
}
__typedArg0 := MakeString("bar!")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp19499 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symabort, Nil)
}
__typedArg0 := symabort
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19500 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symabsolute, tmp19499)
}
__typedArg0 := symabsolute
__typedArg1 := tmp19499
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19501 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symabsvector, tmp19500)
}
__typedArg0 := symabsvector
__typedArg1 := tmp19500
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19502 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symabsvector_2, tmp19501)
}
__typedArg0 := symabsvector_2
__typedArg1 := tmp19501
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19503 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symaddress_1_6, tmp19502)
}
__typedArg0 := symaddress_1_6
__typedArg1 := tmp19502
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19504 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1address, tmp19503)
}
__typedArg0 := sym_5_1address
__typedArg1 := tmp19503
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19505 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symadjoin, tmp19504)
}
__typedArg0 := symadjoin
__typedArg1 := tmp19504
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19506 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symand, tmp19505)
}
__typedArg0 := symand
__typedArg1 := tmp19505
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19507 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symappend, tmp19506)
}
__typedArg0 := symappend
__typedArg1 := tmp19506
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19508 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symarity, tmp19507)
}
__typedArg0 := symarity
__typedArg1 := tmp19507
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19509 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symassoc, tmp19508)
}
__typedArg0 := symassoc
__typedArg1 := tmp19508
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19510 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symassertz, tmp19509)
}
__typedArg0 := symassertz
__typedArg1 := tmp19509
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19511 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symasserta, tmp19510)
}
__typedArg0 := symasserta
__typedArg1 := tmp19510
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19512 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symatom_2, tmp19511)
}
__typedArg0 := symatom_2
__typedArg1 := tmp19511
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19513 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstep_2, tmp19512)
}
__typedArg0 := symstep_2
__typedArg1 := tmp19512
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19514 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symspy_2, tmp19513)
}
__typedArg0 := symspy_2
__typedArg1 := tmp19513
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19515 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19498, tmp19514)
}
__typedArg0 := tmp19498
__typedArg1 := tmp19514
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19516 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symbootstrap, tmp19515)
}
__typedArg0 := symbootstrap
__typedArg1 := tmp19515
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19517 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, tmp19516)
}
__typedArg0 := symboolean
__typedArg1 := tmp19516
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19518 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean_2, tmp19517)
}
__typedArg0 := symboolean_2
__typedArg1 := tmp19517
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19519 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symbound_2, tmp19518)
}
__typedArg0 := symbound_2
__typedArg1 := tmp19518
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19520 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symbind, tmp19519)
}
__typedArg0 := symbind
__typedArg1 := tmp19519
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19521 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symclose, tmp19520)
}
__typedArg0 := symclose
__typedArg1 := tmp19520
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19522 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symchar_1stream_2, tmp19521)
}
__typedArg0 := symchar_1stream_2
__typedArg1 := tmp19521
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19523 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcall, tmp19522)
}
__typedArg0 := symcall
__typedArg1 := tmp19522
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19524 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcases, tmp19523)
}
__typedArg0 := symcases
__typedArg1 := tmp19523
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19525 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcd, tmp19524)
}
__typedArg0 := symcd
__typedArg1 := tmp19524
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19526 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcompile, tmp19525)
}
__typedArg0 := symcompile
__typedArg1 := tmp19525
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19527 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symconcat, tmp19526)
}
__typedArg0 := symconcat
__typedArg1 := tmp19526
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19528 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcond, tmp19527)
}
__typedArg0 := symcond
__typedArg1 := tmp19527
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19529 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp19528)
}
__typedArg0 := symcons
__typedArg1 := tmp19528
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19530 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons_2, tmp19529)
}
__typedArg0 := symcons_2
__typedArg1 := tmp19529
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19531 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcn, tmp19530)
}
__typedArg0 := symcn
__typedArg1 := tmp19530
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19532 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symctxt, tmp19531)
}
__typedArg0 := symctxt
__typedArg1 := tmp19531
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19533 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdatatypes, tmp19532)
}
__typedArg0 := symdatatypes
__typedArg1 := tmp19532
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19534 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdatatype, tmp19533)
}
__typedArg0 := symdatatype
__typedArg1 := tmp19533
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19535 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdeclare, tmp19534)
}
__typedArg0 := symdeclare
__typedArg1 := tmp19534
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19536 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefprolog, tmp19535)
}
__typedArg0 := symdefprolog
__typedArg1 := tmp19535
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19537 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefcc, tmp19536)
}
__typedArg0 := symdefcc
__typedArg1 := tmp19536
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19538 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefmacro, tmp19537)
}
__typedArg0 := symdefmacro
__typedArg1 := tmp19537
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19539 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefine, tmp19538)
}
__typedArg0 := symdefine
__typedArg1 := tmp19538
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19540 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefun, tmp19539)
}
__typedArg0 := symdefun
__typedArg1 := tmp19539
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19541 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdestroy, tmp19540)
}
__typedArg0 := symdestroy
__typedArg1 := tmp19540
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19542 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdifference, tmp19541)
}
__typedArg0 := symdifference
__typedArg1 := tmp19541
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19543 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdo, tmp19542)
}
__typedArg0 := symdo
__typedArg1 := tmp19542
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19544 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symelement_2, tmp19543)
}
__typedArg0 := symelement_2
__typedArg1 := tmp19543
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19545 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symempty_2, tmp19544)
}
__typedArg0 := symempty_2
__typedArg1 := tmp19544
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19546 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symerror, tmp19545)
}
__typedArg0 := symerror
__typedArg1 := tmp19545
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19547 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symerror_1to_1string, tmp19546)
}
__typedArg0 := symerror_1to_1string
__typedArg1 := tmp19546
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19548 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symeval, tmp19547)
}
__typedArg0 := symeval
__typedArg1 := tmp19547
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19549 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symeval_1kl, tmp19548)
}
__typedArg0 := symeval_1kl
__typedArg1 := tmp19548
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19550 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symexception, tmp19549)
}
__typedArg0 := symexception
__typedArg1 := tmp19549
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19551 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symexternal, tmp19550)
}
__typedArg0 := symexternal
__typedArg1 := tmp19550
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19552 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symexplode, tmp19551)
}
__typedArg0 := symexplode
__typedArg1 := tmp19551
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19553 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symenable_1type_1theory, tmp19552)
}
__typedArg0 := symenable_1type_1theory
__typedArg1 := tmp19552
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19554 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(False, tmp19553)
}
__typedArg0 := False
__typedArg1 := tmp19553
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19555 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfindall, tmp19554)
}
__typedArg0 := symfindall
__typedArg1 := tmp19554
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19556 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfactorise_2, tmp19555)
}
__typedArg0 := symfactorise_2
__typedArg1 := tmp19555
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19557 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfactorise, tmp19556)
}
__typedArg0 := symfactorise
__typedArg1 := tmp19556
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19558 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfail_1if, tmp19557)
}
__typedArg0 := symfail_1if
__typedArg1 := tmp19557
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19559 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfail, tmp19558)
}
__typedArg0 := symfail
__typedArg1 := tmp19558
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19560 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfile, tmp19559)
}
__typedArg0 := symfile
__typedArg1 := tmp19559
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19561 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfix, tmp19560)
}
__typedArg0 := symfix
__typedArg1 := tmp19560
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19562 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symforeign, tmp19561)
}
__typedArg0 := symforeign
__typedArg1 := tmp19561
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19563 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfork, tmp19562)
}
__typedArg0 := symfork
__typedArg1 := tmp19562
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19564 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfresh, tmp19563)
}
__typedArg0 := symfresh
__typedArg1 := tmp19563
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19565 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfreeze, tmp19564)
}
__typedArg0 := symfreeze
__typedArg1 := tmp19564
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19566 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfst, tmp19565)
}
__typedArg0 := symfst
__typedArg1 := tmp19565
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19567 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfunction, tmp19566)
}
__typedArg0 := symfunction
__typedArg1 := tmp19566
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19568 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfn, tmp19567)
}
__typedArg0 := symfn
__typedArg1 := tmp19567
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19569 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symgensym, tmp19568)
}
__typedArg0 := symgensym
__typedArg1 := tmp19568
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19570 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symget_1time, tmp19569)
}
__typedArg0 := symget_1time
__typedArg1 := tmp19569
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19571 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symget, tmp19570)
}
__typedArg0 := symget
__typedArg1 := tmp19570
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19572 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhash, tmp19571)
}
__typedArg0 := symhash
__typedArg1 := tmp19571
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19573 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhdstr, tmp19572)
}
__typedArg0 := symhdstr
__typedArg1 := tmp19572
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19574 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhdv, tmp19573)
}
__typedArg0 := symhdv
__typedArg1 := tmp19573
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19575 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhd, tmp19574)
}
__typedArg0 := symhd
__typedArg1 := tmp19574
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19576 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhead, tmp19575)
}
__typedArg0 := symhead
__typedArg1 := tmp19575
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19577 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhush, tmp19576)
}
__typedArg0 := symhush
__typedArg1 := tmp19576
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19578 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhush_2, tmp19577)
}
__typedArg0 := symhush_2
__typedArg1 := tmp19577
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19579 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp19578)
}
__typedArg0 := symif
__typedArg1 := tmp19578
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19580 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symimplementation, tmp19579)
}
__typedArg0 := symimplementation
__typedArg1 := tmp19579
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19581 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminternal, tmp19580)
}
__typedArg0 := syminternal
__typedArg1 := tmp19580
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19582 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symin_1package, tmp19581)
}
__typedArg0 := symin_1package
__typedArg1 := tmp19581
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19583 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symin, tmp19582)
}
__typedArg0 := symin
__typedArg1 := tmp19582
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19584 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symis_b, tmp19583)
}
__typedArg0 := symis_b
__typedArg1 := tmp19583
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19585 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symis, tmp19584)
}
__typedArg0 := symis
__typedArg1 := tmp19584
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19586 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symit, tmp19585)
}
__typedArg0 := symit
__typedArg1 := tmp19585
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19587 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminclude_1all_1but, tmp19586)
}
__typedArg0 := syminclude_1all_1but
__typedArg1 := tmp19586
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19588 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminclude, tmp19587)
}
__typedArg0 := syminclude
__typedArg1 := tmp19587
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19589 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symincluded, tmp19588)
}
__typedArg0 := symincluded
__typedArg1 := tmp19588
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19590 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminline, tmp19589)
}
__typedArg0 := syminline
__typedArg1 := tmp19589
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19591 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminput_7, tmp19590)
}
__typedArg0 := syminput_7
__typedArg1 := tmp19590
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19592 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminput, tmp19591)
}
__typedArg0 := syminput
__typedArg1 := tmp19591
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19593 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminteger_2, tmp19592)
}
__typedArg0 := syminteger_2
__typedArg1 := tmp19592
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19594 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symintern, tmp19593)
}
__typedArg0 := symintern
__typedArg1 := tmp19593
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19595 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminferences, tmp19594)
}
__typedArg0 := syminferences
__typedArg1 := tmp19594
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19596 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symintersection, tmp19595)
}
__typedArg0 := symintersection
__typedArg1 := tmp19595
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19597 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symis, tmp19596)
}
__typedArg0 := symis
__typedArg1 := tmp19596
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19598 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlanguage, tmp19597)
}
__typedArg0 := symlanguage
__typedArg1 := tmp19597
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19599 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp19598)
}
__typedArg0 := symlambda
__typedArg1 := tmp19598
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19600 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlazy, tmp19599)
}
__typedArg0 := symlazy
__typedArg1 := tmp19599
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19601 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp19600)
}
__typedArg0 := symlet
__typedArg1 := tmp19600
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19602 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlength, tmp19601)
}
__typedArg0 := symlength
__typedArg1 := tmp19601
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19603 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlimit, tmp19602)
}
__typedArg0 := symlimit
__typedArg1 := tmp19602
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19604 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlineread, tmp19603)
}
__typedArg0 := symlineread
__typedArg1 := tmp19603
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19605 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19604)
}
__typedArg0 := symlist
__typedArg1 := tmp19604
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19606 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symloaded, tmp19605)
}
__typedArg0 := symloaded
__typedArg1 := tmp19605
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19607 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symload, tmp19606)
}
__typedArg0 := symload
__typedArg1 := tmp19606
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19608 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symmake_1string, tmp19607)
}
__typedArg0 := symmake_1string
__typedArg1 := tmp19607
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19609 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symmap, tmp19608)
}
__typedArg0 := symmap
__typedArg1 := tmp19608
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19610 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symmapcan, tmp19609)
}
__typedArg0 := symmapcan
__typedArg1 := tmp19609
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19611 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symmaxinferences, tmp19610)
}
__typedArg0 := symmaxinferences
__typedArg1 := tmp19610
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19612 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symmacroexpand, tmp19611)
}
__typedArg0 := symmacroexpand
__typedArg1 := tmp19611
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19613 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symmode, tmp19612)
}
__typedArg0 := symmode
__typedArg1 := tmp19612
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19614 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnl, tmp19613)
}
__typedArg0 := symnl
__typedArg1 := tmp19613
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19615 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnot, tmp19614)
}
__typedArg0 := symnot
__typedArg1 := tmp19614
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19616 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnth, tmp19615)
}
__typedArg0 := symnth
__typedArg1 := tmp19615
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19617 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnull, tmp19616)
}
__typedArg0 := symnull
__typedArg1 := tmp19616
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19618 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp19617)
}
__typedArg0 := symnumber
__typedArg1 := tmp19617
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19619 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber_2, tmp19618)
}
__typedArg0 := symnumber_2
__typedArg1 := tmp19618
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19620 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symn_1_6string, tmp19619)
}
__typedArg0 := symn_1_6string
__typedArg1 := tmp19619
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19621 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symoccurs_2, tmp19620)
}
__typedArg0 := symoccurs_2
__typedArg1 := tmp19620
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19622 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symoccurs_1check, tmp19621)
}
__typedArg0 := symoccurs_1check
__typedArg1 := tmp19621
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19623 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symoccurrences, tmp19622)
}
__typedArg0 := symoccurrences
__typedArg1 := tmp19622
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19624 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symopen, tmp19623)
}
__typedArg0 := symopen
__typedArg1 := tmp19623
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19625 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symoptimise_2, tmp19624)
}
__typedArg0 := symoptimise_2
__typedArg1 := tmp19624
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19626 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symoptimise, tmp19625)
}
__typedArg0 := symoptimise
__typedArg1 := tmp19625
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19627 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symor, tmp19626)
}
__typedArg0 := symor
__typedArg1 := tmp19626
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19628 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symos, tmp19627)
}
__typedArg0 := symos
__typedArg1 := tmp19627
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19629 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symout, tmp19628)
}
__typedArg0 := symout
__typedArg1 := tmp19628
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19630 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symoutput, tmp19629)
}
__typedArg0 := symoutput
__typedArg1 := tmp19629
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19631 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympackage, tmp19630)
}
__typedArg0 := sympackage
__typedArg1 := tmp19630
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19632 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symport, tmp19631)
}
__typedArg0 := symport
__typedArg1 := tmp19631
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19633 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symporters, tmp19632)
}
__typedArg0 := symporters
__typedArg1 := tmp19632
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19634 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympos, tmp19633)
}
__typedArg0 := sympos
__typedArg1 := tmp19633
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19635 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympr, tmp19634)
}
__typedArg0 := sympr
__typedArg1 := tmp19634
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19636 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symprint, tmp19635)
}
__typedArg0 := symprint
__typedArg1 := tmp19635
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19637 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symprolog_1memory, tmp19636)
}
__typedArg0 := symprolog_1memory
__typedArg1 := tmp19636
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19638 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symprofile, tmp19637)
}
__typedArg0 := symprofile
__typedArg1 := tmp19637
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19639 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symprofile_1results, tmp19638)
}
__typedArg0 := symprofile_1results
__typedArg1 := tmp19638
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19640 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symprotect, tmp19639)
}
__typedArg0 := symprotect
__typedArg1 := tmp19639
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19641 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symprolog_2, tmp19640)
}
__typedArg0 := symprolog_2
__typedArg1 := tmp19640
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19642 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symps, tmp19641)
}
__typedArg0 := symps
__typedArg1 := tmp19641
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19643 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympreclude_1all_1but, tmp19642)
}
__typedArg0 := sympreclude_1all_1but
__typedArg1 := tmp19642
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19644 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympreclude, tmp19643)
}
__typedArg0 := sympreclude
__typedArg1 := tmp19643
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19645 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symput, tmp19644)
}
__typedArg0 := symput
__typedArg1 := tmp19644
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19646 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympackage_2, tmp19645)
}
__typedArg0 := sympackage_2
__typedArg1 := tmp19645
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19647 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1unit_1string, tmp19646)
}
__typedArg0 := symread_1unit_1string
__typedArg1 := tmp19646
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19648 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1from_1string_1unprocessed, tmp19647)
}
__typedArg0 := symread_1from_1string_1unprocessed
__typedArg1 := tmp19647
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19649 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1from_1string, tmp19648)
}
__typedArg0 := symread_1from_1string
__typedArg1 := tmp19648
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19650 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1byte, tmp19649)
}
__typedArg0 := symread_1byte
__typedArg1 := tmp19649
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19651 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1file_1as_1string, tmp19650)
}
__typedArg0 := symread_1file_1as_1string
__typedArg1 := tmp19650
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19652 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1file_1as_1bytelist, tmp19651)
}
__typedArg0 := symread_1file_1as_1bytelist
__typedArg1 := tmp19651
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19653 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1file, tmp19652)
}
__typedArg0 := symread_1file
__typedArg1 := tmp19652
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19654 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symreceive, tmp19653)
}
__typedArg0 := symreceive
__typedArg1 := tmp19653
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19655 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread, tmp19654)
}
__typedArg0 := symread
__typedArg1 := tmp19654
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19656 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symrelease, tmp19655)
}
__typedArg0 := symrelease
__typedArg1 := tmp19655
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19657 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symremove, tmp19656)
}
__typedArg0 := symremove
__typedArg1 := tmp19656
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19658 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symretract, tmp19657)
}
__typedArg0 := symretract
__typedArg1 := tmp19657
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19659 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symreverse, tmp19658)
}
__typedArg0 := symreverse
__typedArg1 := tmp19658
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19660 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symrun, tmp19659)
}
__typedArg0 := symrun
__typedArg1 := tmp19659
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19661 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstr, tmp19660)
}
__typedArg0 := symstr
__typedArg1 := tmp19660
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19662 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsave, tmp19661)
}
__typedArg0 := symsave
__typedArg1 := tmp19661
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19663 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symset, tmp19662)
}
__typedArg0 := symset
__typedArg1 := tmp19662
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19664 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsimple_1error, tmp19663)
}
__typedArg0 := symsimple_1error
__typedArg1 := tmp19663
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19665 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsnd, tmp19664)
}
__typedArg0 := symsnd
__typedArg1 := tmp19664
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19666 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symspecialise, tmp19665)
}
__typedArg0 := symspecialise
__typedArg1 := tmp19665
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19667 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symspy, tmp19666)
}
__typedArg0 := symspy
__typedArg1 := tmp19666
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19668 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsqts, tmp19667)
}
__typedArg0 := symsqts
__typedArg1 := tmp19667
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19669 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstep, tmp19668)
}
__typedArg0 := symstep
__typedArg1 := tmp19668
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19670 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstoutput, tmp19669)
}
__typedArg0 := symstoutput
__typedArg1 := tmp19669
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19671 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsterror, tmp19670)
}
__typedArg0 := symsterror
__typedArg1 := tmp19670
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19672 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstinput, tmp19671)
}
__typedArg0 := symstinput
__typedArg1 := tmp19671
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19673 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp19672)
}
__typedArg0 := symstring
__typedArg1 := tmp19672
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19674 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp19673)
}
__typedArg0 := symstream
__typedArg1 := tmp19673
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19675 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring_1_6n, tmp19674)
}
__typedArg0 := symstring_1_6n
__typedArg1 := tmp19674
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19676 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring_2, tmp19675)
}
__typedArg0 := symstring_2
__typedArg1 := tmp19675
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19677 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsubst, tmp19676)
}
__typedArg0 := symsubst
__typedArg1 := tmp19676
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19678 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsum, tmp19677)
}
__typedArg0 := symsum
__typedArg1 := tmp19677
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19679 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring_1_6symbol, tmp19678)
}
__typedArg0 := symstring_1_6symbol
__typedArg1 := tmp19678
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19680 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol_2, tmp19679)
}
__typedArg0 := symsymbol_2
__typedArg1 := tmp19679
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19681 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp19680)
}
__typedArg0 := symsymbol
__typedArg1 := tmp19680
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19682 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsystem_1S_2, tmp19681)
}
__typedArg0 := symsystem_1S_2
__typedArg1 := tmp19681
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19683 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsynonyms, tmp19682)
}
__typedArg0 := symsynonyms
__typedArg1 := tmp19682
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19684 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsystemf, tmp19683)
}
__typedArg0 := symsystemf
__typedArg1 := tmp19683
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19685 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtail, tmp19684)
}
__typedArg0 := symtail
__typedArg1 := tmp19684
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19686 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtlv, tmp19685)
}
__typedArg0 := symtlv
__typedArg1 := tmp19685
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19687 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtlstr, tmp19686)
}
__typedArg0 := symtlstr
__typedArg1 := tmp19686
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19688 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtl, tmp19687)
}
__typedArg0 := symtl
__typedArg1 := tmp19687
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19689 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtc, tmp19688)
}
__typedArg0 := symtc
__typedArg1 := tmp19688
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19690 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtc_2, tmp19689)
}
__typedArg0 := symtc_2
__typedArg1 := tmp19689
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19691 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symthaw, tmp19690)
}
__typedArg0 := symthaw
__typedArg1 := tmp19690
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19692 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtime, tmp19691)
}
__typedArg0 := symtime
__typedArg1 := tmp19691
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19693 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtrack, tmp19692)
}
__typedArg0 := symtrack
__typedArg1 := tmp19692
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19694 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtracked, tmp19693)
}
__typedArg0 := symtracked
__typedArg1 := tmp19693
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19695 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtrap_1error, tmp19694)
}
__typedArg0 := symtrap_1error
__typedArg1 := tmp19694
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19696 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(True, tmp19695)
}
__typedArg0 := True
__typedArg1 := tmp19695
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19697 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtuple_2, tmp19696)
}
__typedArg0 := symtuple_2
__typedArg1 := tmp19696
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19698 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtype, tmp19697)
}
__typedArg0 := symtype
__typedArg1 := tmp19697
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19699 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symreturn, tmp19698)
}
__typedArg0 := symreturn
__typedArg1 := tmp19698
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19700 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunabsolute, tmp19699)
}
__typedArg0 := symunabsolute
__typedArg1 := tmp19699
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19701 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symundefmacro, tmp19700)
}
__typedArg0 := symundefmacro
__typedArg1 := tmp19700
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19702 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunprofile, tmp19701)
}
__typedArg0 := symunprofile
__typedArg1 := tmp19701
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19703 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunput, tmp19702)
}
__typedArg0 := symunput
__typedArg1 := tmp19702
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19704 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunion, tmp19703)
}
__typedArg0 := symunion
__typedArg1 := tmp19703
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19705 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunix, tmp19704)
}
__typedArg0 := symunix
__typedArg1 := tmp19704
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19706 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunit, tmp19705)
}
__typedArg0 := symunit
__typedArg1 := tmp19705
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19707 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symuntrack, tmp19706)
}
__typedArg0 := symuntrack
__typedArg1 := tmp19706
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19708 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunspecialise, tmp19707)
}
__typedArg0 := symunspecialise
__typedArg1 := tmp19707
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19709 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symupdate_1lambda_1table, tmp19708)
}
__typedArg0 := symupdate_1lambda_1table
__typedArg1 := tmp19708
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19710 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symu_b, tmp19709)
}
__typedArg0 := symu_b
__typedArg1 := tmp19709
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19711 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symuserdefs, tmp19710)
}
__typedArg0 := symuserdefs
__typedArg1 := tmp19710
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19712 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector_2, tmp19711)
}
__typedArg0 := symvector_2
__typedArg1 := tmp19711
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19713 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp19712)
}
__typedArg0 := symvector
__typedArg1 := tmp19712
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19714 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1vector, tmp19713)
}
__typedArg0 := sym_5_1vector
__typedArg1 := tmp19713
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19715 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector_1_6, tmp19714)
}
__typedArg0 := symvector_1_6
__typedArg1 := tmp19714
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19716 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvalue, tmp19715)
}
__typedArg0 := symvalue
__typedArg1 := tmp19715
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19717 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvar_2, tmp19716)
}
__typedArg0 := symvar_2
__typedArg1 := tmp19716
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19718 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvariable_2, tmp19717)
}
__typedArg0 := symvariable_2
__typedArg1 := tmp19717
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19719 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symverified, tmp19718)
}
__typedArg0 := symverified
__typedArg1 := tmp19718
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19720 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symversion, tmp19719)
}
__typedArg0 := symversion
__typedArg1 := tmp19719
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19721 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwarn, tmp19720)
}
__typedArg0 := symwarn
__typedArg1 := tmp19720
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19722 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhen, tmp19721)
}
__typedArg0 := symwhen
__typedArg1 := tmp19721
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19723 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhere, tmp19722)
}
__typedArg0 := symwhere
__typedArg1 := tmp19722
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19724 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwrite_1byte, tmp19723)
}
__typedArg0 := symwrite_1byte
__typedArg1 := tmp19723
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19725 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwrite_1to_1file, tmp19724)
}
__typedArg0 := symwrite_1to_1file
__typedArg1 := tmp19724
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19726 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwrite_1string, tmp19725)
}
__typedArg0 := symwrite_1string
__typedArg1 := tmp19725
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19727 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symy_1or_1n_2, tmp19726)
}
__typedArg0 := symy_1or_1n_2
__typedArg1 := tmp19726
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19728 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19497, tmp19727)
}
__typedArg0 := tmp19497
__typedArg1 := tmp19727
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19729 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_6_6, tmp19728)
}
__typedArg0 := sym_6_6
__typedArg1 := tmp19728
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19730 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5, tmp19729)
}
__typedArg0 := sym_5
__typedArg1 := tmp19729
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19731 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_a, tmp19730)
}
__typedArg0 := sym_5_a
__typedArg1 := tmp19730
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19732 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_7, tmp19731)
}
__typedArg0 := sym_7
__typedArg1 := tmp19731
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19733 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_d, tmp19732)
}
__typedArg0 := sym_d
__typedArg1 := tmp19732
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19734 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_c, tmp19733)
}
__typedArg0 := sym_c
__typedArg1 := tmp19733
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19735 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1, tmp19734)
}
__typedArg0 := sym_1
__typedArg1 := tmp19734
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19736 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_3, tmp19735)
}
__typedArg0 := sym_3
__typedArg1 := tmp19735
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19737 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5end_6, tmp19736)
}
__typedArg0 := sym_5end_6
__typedArg1 := tmp19736
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19738 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_b_6, tmp19737)
}
__typedArg0 := sym_5_b_6
__typedArg1 := tmp19737
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19739 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_c_4, tmp19738)
}
__typedArg0 := sym_c_4
__typedArg1 := tmp19738
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19740 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_a_a_6, tmp19739)
}
__typedArg0 := sym_a_a_6
__typedArg1 := tmp19739
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19741 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_6, tmp19740)
}
__typedArg0 := sym_6
__typedArg1 := tmp19740
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19742 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_6_a, tmp19741)
}
__typedArg0 := sym_6_a
__typedArg1 := tmp19741
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19743 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_a, tmp19742)
}
__typedArg0 := sym_a
__typedArg1 := tmp19742
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19744 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_a_a, tmp19743)
}
__typedArg0 := sym_a_a
__typedArg1 := tmp19743
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19745 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5e_6, tmp19744)
}
__typedArg0 := sym_5e_6
__typedArg1 := tmp19744
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19746 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_6, tmp19745)
}
__typedArg0 := sym_1_6
__typedArg1 := tmp19745
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19747 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1, tmp19746)
}
__typedArg0 := sym_5_1
__typedArg1 := tmp19746
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19748 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dhush_d, tmp19747)
}
__typedArg0 := sym_dhush_d
__typedArg1 := tmp19747
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19749 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dporters_d, tmp19748)
}
__typedArg0 := sym_dporters_d
__typedArg1 := tmp19748
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19750 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dport_d, tmp19749)
}
__typedArg0 := sym_dport_d
__typedArg1 := tmp19749
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19751 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8s, tmp19750)
}
__typedArg0 := sym_8s
__typedArg1 := tmp19750
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19752 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8p, tmp19751)
}
__typedArg0 := sym_8p
__typedArg1 := tmp19751
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19753 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8v, tmp19752)
}
__typedArg0 := sym_8v
__typedArg1 := tmp19752
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19754 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dproperty_1vector_d, tmp19753)
}
__typedArg0 := sym_dproperty_1vector_d
__typedArg1 := tmp19753
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19755 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_drelease_d, tmp19754)
}
__typedArg0 := sym_drelease_d
__typedArg1 := tmp19754
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19756 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dos_d, tmp19755)
}
__typedArg0 := sym_dos_d
__typedArg1 := tmp19755
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19757 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dmacros_d, tmp19756)
}
__typedArg0 := sym_dmacros_d
__typedArg1 := tmp19756
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19758 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dmaximum_1print_1sequence_1size_d, tmp19757)
}
__typedArg0 := sym_dmaximum_1print_1sequence_1size_d
__typedArg1 := tmp19757
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19759 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dversion_d, tmp19758)
}
__typedArg0 := sym_dversion_d
__typedArg1 := tmp19758
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19760 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dhome_1directory_d, tmp19759)
}
__typedArg0 := sym_dhome_1directory_d
__typedArg1 := tmp19759
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19761 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dstoutput_d, tmp19760)
}
__typedArg0 := sym_dstoutput_d
__typedArg1 := tmp19760
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19762 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dsterror_d, tmp19761)
}
__typedArg0 := sym_dsterror_d
__typedArg1 := tmp19761
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19763 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dstinput_d, tmp19762)
}
__typedArg0 := sym_dstinput_d
__typedArg1 := tmp19762
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19764 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dimplementation_d, tmp19763)
}
__typedArg0 := sym_dimplementation_d
__typedArg1 := tmp19763
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19765 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dlanguage_d, tmp19764)
}
__typedArg0 := sym_dlanguage_d
__typedArg1 := tmp19764
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19766 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym__, tmp19765)
}
__typedArg0 := sym__
__typedArg1 := tmp19765
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19767 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19496, tmp19766)
}
__typedArg0 := tmp19496
__typedArg1 := tmp19766
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19768 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19495, tmp19767)
}
__typedArg0 := tmp19495
__typedArg1 := tmp19767
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19769 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19494, tmp19768)
}
__typedArg0 := tmp19494
__typedArg1 := tmp19768
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19770 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19493, tmp19769)
}
__typedArg0 := tmp19493
__typedArg1 := tmp19769
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19771 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_e_e, tmp19770)
}
__typedArg0 := sym_e_e
__typedArg1 := tmp19770
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19772 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1_1, tmp19771)
}
__typedArg0 := sym_5_1_1
__typedArg1 := tmp19771
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19773 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19772)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19772
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19774 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_i, tmp19773)
}
__typedArg0 := sym_i
__typedArg1 := tmp19773
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19775 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_j, tmp19774)
}
__typedArg0 := sym_j
__typedArg1 := tmp19774
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19776 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_b, tmp19775)
}
__typedArg0 := sym_b
__typedArg1 := tmp19775
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19777 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp19778 := Call(__e, PrimFunc(symput), symshen, symshen_4external_1symbols, tmp19776, tmp19777)


_ = tmp19778

tmp19779 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symabsvector) {
return PrimAbsvector(MakeInteger(0))
}
__typedArg0 := MakeInteger(0)
return Call(__e, PrimFunc(symabsvector), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dempty_1absvector_d, tmp19779)
}
__typedArg0 := symshen_4_dempty_1absvector_d
__typedArg1 := tmp19779
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return


}, 0)

tmp19780 := Call(__e, ns2_1set, symshen_4initialise_1environment, tmp19016)


_ = tmp19780

tmp19781 := MakeNative(func(__e *ControlFlow) {
tmp19782 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, Nil)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19782

tmp19783 := MakeNative(func(__e *ControlFlow) {
V6094 := __e.Get(1)
_ = V6094
__e.Return(MakeNative(func(__e *ControlFlow) {
B6090 := __e.Get(1)
_ = B6090
__e.Return(MakeNative(func(__e *ControlFlow) {
L6091 := __e.Get(1)
_ = L6091
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6092 := __e.Get(1)
_ = Key6092
__e.Return(MakeNative(func(__e *ControlFlow) {
C6093 := __e.Get(1)
_ = C6093
tmp19784 := Call(__e, PrimFunc(symshen_4newpv), B6090)


let__18929 := tmp19784
_ = let__18929

tmp19785 := Call(__e, PrimFunc(symshen_4newpv), B6090)


let__18930 := tmp19785
_ = let__18930

tmp19786 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18930, Nil)
}
__typedArg0 := let__18930
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19787 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19786)
}
__typedArg0 := symlist
__typedArg1 := tmp19786
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19788 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19787, Nil)
}
__typedArg0 := tmp19787
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19789 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19788)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19788
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19790 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18929, tmp19789)
}
__typedArg0 := let__18929
__typedArg1 := tmp19789
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19791 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19792 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19791)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19791
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19793 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19790, tmp19792)
}
__typedArg0 := tmp19790
__typedArg1 := tmp19792
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19794 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19793, Nil)
}
__typedArg0 := tmp19793
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19795 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19794)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19794
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19796 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp19795)
}
__typedArg0 := symsymbol
__typedArg1 := tmp19795
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19797 := Call(__e, PrimFunc(symis_b), V6094, tmp19796, B6090, L6091, Key6092, C6093)


tmp19798 := Call(__e, PrimFunc(symshen_4gc), B6090, tmp19797)


__e.TailApply(PrimFunc(symshen_4gc), B6090, tmp19798)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19799 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp19800 := Call(__e, PrimFunc(symshen_4assoc_1_6), symshen_4register_1source_1form, tmp19783, tmp19799)


tmp19801 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp19800)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp19800
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19801

tmp19802 := MakeNative(func(__e *ControlFlow) {
V6099 := __e.Get(1)
_ = V6099
__e.Return(MakeNative(func(__e *ControlFlow) {
B6095 := __e.Get(1)
_ = B6095
__e.Return(MakeNative(func(__e *ControlFlow) {
L6096 := __e.Get(1)
_ = L6096
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6097 := __e.Get(1)
_ = Key6097
__e.Return(MakeNative(func(__e *ControlFlow) {
C6098 := __e.Get(1)
_ = C6098
tmp19803 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19804 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19803)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19803
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19805 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp19804)
}
__typedArg0 := symsymbol
__typedArg1 := tmp19804
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6099, tmp19805, B6095, L6096, Key6097, C6098)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19806 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp19807 := Call(__e, PrimFunc(symshen_4assoc_1_6), symshen_4unregister_1source_1form, tmp19802, tmp19806)


tmp19808 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp19807)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp19807
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19808

tmp19809 := MakeNative(func(__e *ControlFlow) {
V6104 := __e.Get(1)
_ = V6104
__e.Return(MakeNative(func(__e *ControlFlow) {
B6100 := __e.Get(1)
_ = B6100
__e.Return(MakeNative(func(__e *ControlFlow) {
L6101 := __e.Get(1)
_ = L6101
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6102 := __e.Get(1)
_ = Key6102
__e.Return(MakeNative(func(__e *ControlFlow) {
C6103 := __e.Get(1)
_ = C6103
tmp19810 := Call(__e, PrimFunc(symshen_4newpv), B6100)


let__18931 := tmp19810
_ = let__18931

tmp19811 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18931, Nil)
}
__typedArg0 := let__18931
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19812 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19811)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19811
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19813 := Call(__e, PrimFunc(symis_b), V6104, tmp19812, B6100, L6101, Key6102, C6103)


__e.TailApply(PrimFunc(symshen_4gc), B6100, tmp19813)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19814 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp19815 := Call(__e, PrimFunc(symshen_4assoc_1_6), symabort, tmp19809, tmp19814)


tmp19816 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp19815)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp19815
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19816

tmp19817 := MakeNative(func(__e *ControlFlow) {
V6109 := __e.Get(1)
_ = V6109
__e.Return(MakeNative(func(__e *ControlFlow) {
B6105 := __e.Get(1)
_ = B6105
__e.Return(MakeNative(func(__e *ControlFlow) {
L6106 := __e.Get(1)
_ = L6106
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6107 := __e.Get(1)
_ = Key6107
__e.Return(MakeNative(func(__e *ControlFlow) {
C6108 := __e.Get(1)
_ = C6108
tmp19818 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19819 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19818)
}
__typedArg0 := symlist
__typedArg1 := tmp19818
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19820 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19819, Nil)
}
__typedArg0 := tmp19819
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19821 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19820)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19820
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19822 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp19821)
}
__typedArg0 := symstring
__typedArg1 := tmp19821
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6109, tmp19822, B6105, L6106, Key6107, C6108)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19823 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp19824 := Call(__e, PrimFunc(symshen_4assoc_1_6), symabsolute, tmp19817, tmp19823)


tmp19825 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp19824)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp19824
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19825

tmp19826 := MakeNative(func(__e *ControlFlow) {
V6114 := __e.Get(1)
_ = V6114
__e.Return(MakeNative(func(__e *ControlFlow) {
B6110 := __e.Get(1)
_ = B6110
__e.Return(MakeNative(func(__e *ControlFlow) {
L6111 := __e.Get(1)
_ = L6111
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6112 := __e.Get(1)
_ = Key6112
__e.Return(MakeNative(func(__e *ControlFlow) {
C6113 := __e.Get(1)
_ = C6113
tmp19827 := Call(__e, PrimFunc(symshen_4newpv), B6110)


let__18932 := tmp19827
_ = let__18932

tmp19828 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19829 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19828)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19828
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19830 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18932, tmp19829)
}
__typedArg0 := let__18932
__typedArg1 := tmp19829
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19831 := Call(__e, PrimFunc(symis_b), V6114, tmp19830, B6110, L6111, Key6112, C6113)


__e.TailApply(PrimFunc(symshen_4gc), B6110, tmp19831)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19832 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp19833 := Call(__e, PrimFunc(symshen_4assoc_1_6), symabsvector_2, tmp19826, tmp19832)


tmp19834 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp19833)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp19833
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19834

tmp19835 := MakeNative(func(__e *ControlFlow) {
V6120 := __e.Get(1)
_ = V6120
__e.Return(MakeNative(func(__e *ControlFlow) {
B6116 := __e.Get(1)
_ = B6116
__e.Return(MakeNative(func(__e *ControlFlow) {
L6117 := __e.Get(1)
_ = L6117
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6118 := __e.Get(1)
_ = Key6118
__e.Return(MakeNative(func(__e *ControlFlow) {
C6119 := __e.Get(1)
_ = C6119
tmp19836 := Call(__e, PrimFunc(symshen_4newpv), B6116)


let__18933 := tmp19836
_ = let__18933

tmp19837 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18933, Nil)
}
__typedArg0 := let__18933
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19838 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19837)
}
__typedArg0 := symlist
__typedArg1 := tmp19837
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19839 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18933, Nil)
}
__typedArg0 := let__18933
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19840 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19839)
}
__typedArg0 := symlist
__typedArg1 := tmp19839
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19841 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19840, Nil)
}
__typedArg0 := tmp19840
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19842 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19841)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19841
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19843 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19838, tmp19842)
}
__typedArg0 := tmp19838
__typedArg1 := tmp19842
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19844 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19843, Nil)
}
__typedArg0 := tmp19843
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19845 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19844)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19844
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19846 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18933, tmp19845)
}
__typedArg0 := let__18933
__typedArg1 := tmp19845
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19847 := Call(__e, PrimFunc(symis_b), V6120, tmp19846, B6116, L6117, Key6118, C6119)


__e.TailApply(PrimFunc(symshen_4gc), B6116, tmp19847)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19848 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp19849 := Call(__e, PrimFunc(symshen_4assoc_1_6), symadjoin, tmp19835, tmp19848)


tmp19850 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp19849)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp19849
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19850

tmp19851 := MakeNative(func(__e *ControlFlow) {
V6125 := __e.Get(1)
_ = V6125
__e.Return(MakeNative(func(__e *ControlFlow) {
B6121 := __e.Get(1)
_ = B6121
__e.Return(MakeNative(func(__e *ControlFlow) {
L6122 := __e.Get(1)
_ = L6122
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6123 := __e.Get(1)
_ = Key6123
__e.Return(MakeNative(func(__e *ControlFlow) {
C6124 := __e.Get(1)
_ = C6124
tmp19852 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19853 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19852)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19852
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19854 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, tmp19853)
}
__typedArg0 := symboolean
__typedArg1 := tmp19853
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19855 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19854, Nil)
}
__typedArg0 := tmp19854
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19856 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19855)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19855
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19857 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, tmp19856)
}
__typedArg0 := symboolean
__typedArg1 := tmp19856
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6125, tmp19857, B6121, L6122, Key6123, C6124)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19858 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp19859 := Call(__e, PrimFunc(symshen_4assoc_1_6), symand, tmp19851, tmp19858)


tmp19860 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp19859)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp19859
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19860

tmp19861 := MakeNative(func(__e *ControlFlow) {
V6130 := __e.Get(1)
_ = V6130
__e.Return(MakeNative(func(__e *ControlFlow) {
B6126 := __e.Get(1)
_ = B6126
__e.Return(MakeNative(func(__e *ControlFlow) {
L6127 := __e.Get(1)
_ = L6127
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6128 := __e.Get(1)
_ = Key6128
__e.Return(MakeNative(func(__e *ControlFlow) {
C6129 := __e.Get(1)
_ = C6129
tmp19862 := Call(__e, PrimFunc(symshen_4newpv), B6126)


let__18934 := tmp19862
_ = let__18934

tmp19863 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19864 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19863)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19863
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19865 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp19864)
}
__typedArg0 := symsymbol
__typedArg1 := tmp19864
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19866 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19865, Nil)
}
__typedArg0 := tmp19865
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19867 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19866)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19866
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19868 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp19867)
}
__typedArg0 := symstring
__typedArg1 := tmp19867
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19869 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19868, Nil)
}
__typedArg0 := tmp19868
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19870 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19869)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19869
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19871 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18934, tmp19870)
}
__typedArg0 := let__18934
__typedArg1 := tmp19870
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19872 := Call(__e, PrimFunc(symis_b), V6130, tmp19871, B6126, L6127, Key6128, C6129)


__e.TailApply(PrimFunc(symshen_4gc), B6126, tmp19872)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19873 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp19874 := Call(__e, PrimFunc(symshen_4assoc_1_6), symshen_4app, tmp19861, tmp19873)


tmp19875 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp19874)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp19874
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19875

tmp19876 := MakeNative(func(__e *ControlFlow) {
V6137 := __e.Get(1)
_ = V6137
__e.Return(MakeNative(func(__e *ControlFlow) {
B6133 := __e.Get(1)
_ = B6133
__e.Return(MakeNative(func(__e *ControlFlow) {
L6134 := __e.Get(1)
_ = L6134
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6135 := __e.Get(1)
_ = Key6135
__e.Return(MakeNative(func(__e *ControlFlow) {
C6136 := __e.Get(1)
_ = C6136
tmp19877 := Call(__e, PrimFunc(symshen_4newpv), B6133)


let__18935 := tmp19877
_ = let__18935

tmp19878 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18935, Nil)
}
__typedArg0 := let__18935
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19879 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19878)
}
__typedArg0 := symlist
__typedArg1 := tmp19878
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19880 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18935, Nil)
}
__typedArg0 := let__18935
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19881 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19880)
}
__typedArg0 := symlist
__typedArg1 := tmp19880
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19882 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18935, Nil)
}
__typedArg0 := let__18935
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19883 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19882)
}
__typedArg0 := symlist
__typedArg1 := tmp19882
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19884 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19883, Nil)
}
__typedArg0 := tmp19883
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19885 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19884)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19884
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19886 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19881, tmp19885)
}
__typedArg0 := tmp19881
__typedArg1 := tmp19885
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19887 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19886, Nil)
}
__typedArg0 := tmp19886
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19888 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19887)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19887
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19889 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19879, tmp19888)
}
__typedArg0 := tmp19879
__typedArg1 := tmp19888
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19890 := Call(__e, PrimFunc(symis_b), V6137, tmp19889, B6133, L6134, Key6135, C6136)


__e.TailApply(PrimFunc(symshen_4gc), B6133, tmp19890)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19891 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp19892 := Call(__e, PrimFunc(symshen_4assoc_1_6), symappend, tmp19876, tmp19891)


tmp19893 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp19892)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp19892
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19893

tmp19894 := MakeNative(func(__e *ControlFlow) {
V6142 := __e.Get(1)
_ = V6142
__e.Return(MakeNative(func(__e *ControlFlow) {
B6138 := __e.Get(1)
_ = B6138
__e.Return(MakeNative(func(__e *ControlFlow) {
L6139 := __e.Get(1)
_ = L6139
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6140 := __e.Get(1)
_ = Key6140
__e.Return(MakeNative(func(__e *ControlFlow) {
C6141 := __e.Get(1)
_ = C6141
tmp19895 := Call(__e, PrimFunc(symshen_4newpv), B6138)


let__18936 := tmp19895
_ = let__18936

tmp19896 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19897 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19896)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19896
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19898 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18936, tmp19897)
}
__typedArg0 := let__18936
__typedArg1 := tmp19897
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19899 := Call(__e, PrimFunc(symis_b), V6142, tmp19898, B6138, L6139, Key6140, C6141)


__e.TailApply(PrimFunc(symshen_4gc), B6138, tmp19899)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19900 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp19901 := Call(__e, PrimFunc(symshen_4assoc_1_6), symarity, tmp19894, tmp19900)


tmp19902 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp19901)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp19901
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19902

tmp19903 := MakeNative(func(__e *ControlFlow) {
V6148 := __e.Get(1)
_ = V6148
__e.Return(MakeNative(func(__e *ControlFlow) {
B6144 := __e.Get(1)
_ = B6144
__e.Return(MakeNative(func(__e *ControlFlow) {
L6145 := __e.Get(1)
_ = L6145
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6146 := __e.Get(1)
_ = Key6146
__e.Return(MakeNative(func(__e *ControlFlow) {
C6147 := __e.Get(1)
_ = C6147
tmp19904 := Call(__e, PrimFunc(symshen_4newpv), B6144)


let__18937 := tmp19904
_ = let__18937

tmp19905 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18937, Nil)
}
__typedArg0 := let__18937
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19906 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19905)
}
__typedArg0 := symlist
__typedArg1 := tmp19905
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19907 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19906, Nil)
}
__typedArg0 := tmp19906
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19908 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19907)
}
__typedArg0 := symlist
__typedArg1 := tmp19907
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19909 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18937, Nil)
}
__typedArg0 := let__18937
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19910 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19909)
}
__typedArg0 := symlist
__typedArg1 := tmp19909
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19911 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19910, Nil)
}
__typedArg0 := tmp19910
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19912 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19911)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19911
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19913 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19908, tmp19912)
}
__typedArg0 := tmp19908
__typedArg1 := tmp19912
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19914 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19913, Nil)
}
__typedArg0 := tmp19913
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19915 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19914)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19914
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19916 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18937, tmp19915)
}
__typedArg0 := let__18937
__typedArg1 := tmp19915
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19917 := Call(__e, PrimFunc(symis_b), V6148, tmp19916, B6144, L6145, Key6146, C6147)


__e.TailApply(PrimFunc(symshen_4gc), B6144, tmp19917)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19918 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp19919 := Call(__e, PrimFunc(symshen_4assoc_1_6), symassoc, tmp19903, tmp19918)


tmp19920 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp19919)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp19919
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19920

tmp19921 := MakeNative(func(__e *ControlFlow) {
V6153 := __e.Get(1)
_ = V6153
__e.Return(MakeNative(func(__e *ControlFlow) {
B6149 := __e.Get(1)
_ = B6149
__e.Return(MakeNative(func(__e *ControlFlow) {
L6150 := __e.Get(1)
_ = L6150
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6151 := __e.Get(1)
_ = Key6151
__e.Return(MakeNative(func(__e *ControlFlow) {
C6152 := __e.Get(1)
_ = C6152
tmp19922 := Call(__e, PrimFunc(symshen_4newpv), B6149)


let__18938 := tmp19922
_ = let__18938

tmp19923 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19924 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19923)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19923
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19925 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18938, tmp19924)
}
__typedArg0 := let__18938
__typedArg1 := tmp19924
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19926 := Call(__e, PrimFunc(symis_b), V6153, tmp19925, B6149, L6150, Key6151, C6152)


__e.TailApply(PrimFunc(symshen_4gc), B6149, tmp19926)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19927 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp19928 := Call(__e, PrimFunc(symshen_4assoc_1_6), symatom_2, tmp19921, tmp19927)


tmp19929 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp19928)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp19928
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19929

tmp19930 := MakeNative(func(__e *ControlFlow) {
V6158 := __e.Get(1)
_ = V6158
__e.Return(MakeNative(func(__e *ControlFlow) {
B6154 := __e.Get(1)
_ = B6154
__e.Return(MakeNative(func(__e *ControlFlow) {
L6155 := __e.Get(1)
_ = L6155
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6156 := __e.Get(1)
_ = Key6156
__e.Return(MakeNative(func(__e *ControlFlow) {
C6157 := __e.Get(1)
_ = C6157
tmp19931 := Call(__e, PrimFunc(symshen_4newpv), B6154)


let__18939 := tmp19931
_ = let__18939

tmp19932 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19933 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19932)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19932
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19934 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18939, tmp19933)
}
__typedArg0 := let__18939
__typedArg1 := tmp19933
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19935 := Call(__e, PrimFunc(symis_b), V6158, tmp19934, B6154, L6155, Key6156, C6157)


__e.TailApply(PrimFunc(symshen_4gc), B6154, tmp19935)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19936 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp19937 := Call(__e, PrimFunc(symshen_4assoc_1_6), symboolean_2, tmp19930, tmp19936)


tmp19938 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp19937)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp19937
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19938

tmp19939 := MakeNative(func(__e *ControlFlow) {
V6163 := __e.Get(1)
_ = V6163
__e.Return(MakeNative(func(__e *ControlFlow) {
B6159 := __e.Get(1)
_ = B6159
__e.Return(MakeNative(func(__e *ControlFlow) {
L6160 := __e.Get(1)
_ = L6160
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6161 := __e.Get(1)
_ = Key6161
__e.Return(MakeNative(func(__e *ControlFlow) {
C6162 := __e.Get(1)
_ = C6162
tmp19940 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19941 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19940)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19940
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19942 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp19941)
}
__typedArg0 := symstring
__typedArg1 := tmp19941
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6163, tmp19942, B6159, L6160, Key6161, C6162)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19943 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp19944 := Call(__e, PrimFunc(symshen_4assoc_1_6), symbootstrap, tmp19939, tmp19943)


tmp19945 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp19944)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp19944
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19945

tmp19946 := MakeNative(func(__e *ControlFlow) {
V6168 := __e.Get(1)
_ = V6168
__e.Return(MakeNative(func(__e *ControlFlow) {
B6164 := __e.Get(1)
_ = B6164
__e.Return(MakeNative(func(__e *ControlFlow) {
L6165 := __e.Get(1)
_ = L6165
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6166 := __e.Get(1)
_ = Key6166
__e.Return(MakeNative(func(__e *ControlFlow) {
C6167 := __e.Get(1)
_ = C6167
tmp19947 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19948 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19947)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19947
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19949 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp19948)
}
__typedArg0 := symsymbol
__typedArg1 := tmp19948
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6168, tmp19949, B6164, L6165, Key6166, C6167)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19950 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp19951 := Call(__e, PrimFunc(symshen_4assoc_1_6), symbound_2, tmp19946, tmp19950)


tmp19952 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp19951)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp19951
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19952

tmp19953 := MakeNative(func(__e *ControlFlow) {
V6173 := __e.Get(1)
_ = V6173
__e.Return(MakeNative(func(__e *ControlFlow) {
B6169 := __e.Get(1)
_ = B6169
__e.Return(MakeNative(func(__e *ControlFlow) {
L6170 := __e.Get(1)
_ = L6170
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6171 := __e.Get(1)
_ = Key6171
__e.Return(MakeNative(func(__e *ControlFlow) {
C6172 := __e.Get(1)
_ = C6172
tmp19954 := Call(__e, PrimFunc(symshen_4newpv), B6169)


let__18940 := tmp19954
_ = let__18940

tmp19955 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18940, Nil)
}
__typedArg0 := let__18940
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19956 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19955)
}
__typedArg0 := symlist
__typedArg1 := tmp19955
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19957 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19958 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19957)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19957
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19959 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19956, tmp19958)
}
__typedArg0 := tmp19956
__typedArg1 := tmp19958
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19960 := Call(__e, PrimFunc(symis_b), V6173, tmp19959, B6169, L6170, Key6171, C6172)


__e.TailApply(PrimFunc(symshen_4gc), B6169, tmp19960)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19961 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp19962 := Call(__e, PrimFunc(symshen_4assoc_1_6), symshen_4ccons_2, tmp19953, tmp19961)


tmp19963 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp19962)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp19962
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19963

tmp19964 := MakeNative(func(__e *ControlFlow) {
V6178 := __e.Get(1)
_ = V6178
__e.Return(MakeNative(func(__e *ControlFlow) {
B6174 := __e.Get(1)
_ = B6174
__e.Return(MakeNative(func(__e *ControlFlow) {
L6175 := __e.Get(1)
_ = L6175
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6176 := __e.Get(1)
_ = Key6176
__e.Return(MakeNative(func(__e *ControlFlow) {
C6177 := __e.Get(1)
_ = C6177
tmp19965 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19966 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19965)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19965
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19967 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp19966)
}
__typedArg0 := symstring
__typedArg1 := tmp19966
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6178, tmp19967, B6174, L6175, Key6176, C6177)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19968 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp19969 := Call(__e, PrimFunc(symshen_4assoc_1_6), symcd, tmp19964, tmp19968)


tmp19970 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp19969)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp19969
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19970

tmp19971 := MakeNative(func(__e *ControlFlow) {
V6184 := __e.Get(1)
_ = V6184
__e.Return(MakeNative(func(__e *ControlFlow) {
B6180 := __e.Get(1)
_ = B6180
__e.Return(MakeNative(func(__e *ControlFlow) {
L6181 := __e.Get(1)
_ = L6181
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6182 := __e.Get(1)
_ = Key6182
__e.Return(MakeNative(func(__e *ControlFlow) {
C6183 := __e.Get(1)
_ = C6183
tmp19972 := Call(__e, PrimFunc(symshen_4newpv), B6180)


let__18941 := tmp19972
_ = let__18941

tmp19973 := Call(__e, PrimFunc(symshen_4newpv), B6180)


let__18942 := tmp19973
_ = let__18942

tmp19974 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18941, Nil)
}
__typedArg0 := let__18941
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19975 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp19974)
}
__typedArg0 := symstream
__typedArg1 := tmp19974
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19976 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18942, Nil)
}
__typedArg0 := let__18942
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19977 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19976)
}
__typedArg0 := symlist
__typedArg1 := tmp19976
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19978 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19977, Nil)
}
__typedArg0 := tmp19977
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19979 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19978)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19978
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19980 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19975, tmp19979)
}
__typedArg0 := tmp19975
__typedArg1 := tmp19979
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19981 := Call(__e, PrimFunc(symis_b), V6184, tmp19980, B6180, L6181, Key6182, C6183)


tmp19982 := Call(__e, PrimFunc(symshen_4gc), B6180, tmp19981)


__e.TailApply(PrimFunc(symshen_4gc), B6180, tmp19982)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19983 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp19984 := Call(__e, PrimFunc(symshen_4assoc_1_6), symclose, tmp19971, tmp19983)


tmp19985 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp19984)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp19984
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19985

tmp19986 := MakeNative(func(__e *ControlFlow) {
V6189 := __e.Get(1)
_ = V6189
__e.Return(MakeNative(func(__e *ControlFlow) {
B6185 := __e.Get(1)
_ = B6185
__e.Return(MakeNative(func(__e *ControlFlow) {
L6186 := __e.Get(1)
_ = L6186
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6187 := __e.Get(1)
_ = Key6187
__e.Return(MakeNative(func(__e *ControlFlow) {
C6188 := __e.Get(1)
_ = C6188
tmp19987 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19988 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19987)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19987
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19989 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp19988)
}
__typedArg0 := symstring
__typedArg1 := tmp19988
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19990 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19989, Nil)
}
__typedArg0 := tmp19989
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19991 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19990)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19990
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19992 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp19991)
}
__typedArg0 := symstring
__typedArg1 := tmp19991
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6189, tmp19992, B6185, L6186, Key6187, C6188)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19993 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp19994 := Call(__e, PrimFunc(symshen_4assoc_1_6), symcn, tmp19986, tmp19993)


tmp19995 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp19994)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp19994
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19995

tmp19996 := MakeNative(func(__e *ControlFlow) {
V6194 := __e.Get(1)
_ = V6194
__e.Return(MakeNative(func(__e *ControlFlow) {
B6190 := __e.Get(1)
_ = B6190
__e.Return(MakeNative(func(__e *ControlFlow) {
L6191 := __e.Get(1)
_ = L6191
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6192 := __e.Get(1)
_ = Key6192
__e.Return(MakeNative(func(__e *ControlFlow) {
C6193 := __e.Get(1)
_ = C6193
tmp19997 := Call(__e, PrimFunc(symshen_4newpv), B6190)


let__18943 := tmp19997
_ = let__18943

tmp19998 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18943, Nil)
}
__typedArg0 := let__18943
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19999 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp19998)
}
__typedArg0 := symstream
__typedArg1 := tmp19998
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20000 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20001 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20000)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20000
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20002 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19999, tmp20001)
}
__typedArg0 := tmp19999
__typedArg1 := tmp20001
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20003 := Call(__e, PrimFunc(symis_b), V6194, tmp20002, B6190, L6191, Key6192, C6193)


__e.TailApply(PrimFunc(symshen_4gc), B6190, tmp20003)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20004 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20005 := Call(__e, PrimFunc(symshen_4assoc_1_6), symchar_1stream_2, tmp19996, tmp20004)


tmp20006 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20005)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20005
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20006

tmp20007 := MakeNative(func(__e *ControlFlow) {
V6201 := __e.Get(1)
_ = V6201
__e.Return(MakeNative(func(__e *ControlFlow) {
B6197 := __e.Get(1)
_ = B6197
__e.Return(MakeNative(func(__e *ControlFlow) {
L6198 := __e.Get(1)
_ = L6198
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6199 := __e.Get(1)
_ = Key6199
__e.Return(MakeNative(func(__e *ControlFlow) {
C6200 := __e.Get(1)
_ = C6200
tmp20008 := Call(__e, PrimFunc(symshen_4newpv), B6197)


let__18944 := tmp20008
_ = let__18944

tmp20009 := Call(__e, PrimFunc(symshen_4newpv), B6197)


let__18945 := tmp20009
_ = let__18945

tmp20010 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18944, Nil)
}
__typedArg0 := let__18944
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20011 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20010)
}
__typedArg0 := symlist
__typedArg1 := tmp20010
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20012 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18944, Nil)
}
__typedArg0 := let__18944
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20013 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20012)
}
__typedArg0 := symlist
__typedArg1 := tmp20012
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20014 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18945, Nil)
}
__typedArg0 := let__18945
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20015 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20013, tmp20014)
}
__typedArg0 := tmp20013
__typedArg1 := tmp20014
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20016 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstr, tmp20015)
}
__typedArg0 := symstr
__typedArg1 := tmp20015
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20017 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20016, Nil)
}
__typedArg0 := tmp20016
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20018 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20017)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20017
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20019 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20011, tmp20018)
}
__typedArg0 := tmp20011
__typedArg1 := tmp20018
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20020 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18944, Nil)
}
__typedArg0 := let__18944
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20021 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20020)
}
__typedArg0 := symlist
__typedArg1 := tmp20020
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20022 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18945, Nil)
}
__typedArg0 := let__18945
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20023 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20022)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20022
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20024 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20021, tmp20023)
}
__typedArg0 := tmp20021
__typedArg1 := tmp20023
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20025 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20024, Nil)
}
__typedArg0 := tmp20024
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20026 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20025)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20025
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20027 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20019, tmp20026)
}
__typedArg0 := tmp20019
__typedArg1 := tmp20026
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20028 := Call(__e, PrimFunc(symis_b), V6201, tmp20027, B6197, L6198, Key6199, C6200)


tmp20029 := Call(__e, PrimFunc(symshen_4gc), B6197, tmp20028)


__e.TailApply(PrimFunc(symshen_4gc), B6197, tmp20029)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20030 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20031 := Call(__e, PrimFunc(symshen_4assoc_1_6), symcompile, tmp20007, tmp20030)


tmp20032 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20031)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20031
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20032

tmp20033 := MakeNative(func(__e *ControlFlow) {
V6206 := __e.Get(1)
_ = V6206
__e.Return(MakeNative(func(__e *ControlFlow) {
B6202 := __e.Get(1)
_ = B6202
__e.Return(MakeNative(func(__e *ControlFlow) {
L6203 := __e.Get(1)
_ = L6203
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6204 := __e.Get(1)
_ = Key6204
__e.Return(MakeNative(func(__e *ControlFlow) {
C6205 := __e.Get(1)
_ = C6205
tmp20034 := Call(__e, PrimFunc(symshen_4newpv), B6202)


let__18946 := tmp20034
_ = let__18946

tmp20035 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20036 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20035)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20035
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20037 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18946, tmp20036)
}
__typedArg0 := let__18946
__typedArg1 := tmp20036
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20038 := Call(__e, PrimFunc(symis_b), V6206, tmp20037, B6202, L6203, Key6204, C6205)


__e.TailApply(PrimFunc(symshen_4gc), B6202, tmp20038)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20039 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20040 := Call(__e, PrimFunc(symshen_4assoc_1_6), symcons_2, tmp20033, tmp20039)


tmp20041 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20040)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20040
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20041

tmp20042 := MakeNative(func(__e *ControlFlow) {
V6211 := __e.Get(1)
_ = V6211
__e.Return(MakeNative(func(__e *ControlFlow) {
B6207 := __e.Get(1)
_ = B6207
__e.Return(MakeNative(func(__e *ControlFlow) {
L6208 := __e.Get(1)
_ = L6208
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6209 := __e.Get(1)
_ = Key6209
__e.Return(MakeNative(func(__e *ControlFlow) {
C6210 := __e.Get(1)
_ = C6210
tmp20043 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20044 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20043)
}
__typedArg0 := symlist
__typedArg1 := tmp20043
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20045 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20044, Nil)
}
__typedArg0 := tmp20044
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20046 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20045)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20045
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6211, tmp20046, B6207, L6208, Key6209, C6210)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20047 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20048 := Call(__e, PrimFunc(symshen_4assoc_1_6), symdatatypes, tmp20042, tmp20047)


tmp20049 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20048)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20048
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20049

tmp20050 := MakeNative(func(__e *ControlFlow) {
V6216 := __e.Get(1)
_ = V6216
__e.Return(MakeNative(func(__e *ControlFlow) {
B6212 := __e.Get(1)
_ = B6212
__e.Return(MakeNative(func(__e *ControlFlow) {
L6213 := __e.Get(1)
_ = L6213
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6214 := __e.Get(1)
_ = Key6214
__e.Return(MakeNative(func(__e *ControlFlow) {
C6215 := __e.Get(1)
_ = C6215
tmp20051 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20052 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20051)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20051
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20053 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp20052)
}
__typedArg0 := symsymbol
__typedArg1 := tmp20052
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6216, tmp20053, B6212, L6213, Key6214, C6215)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20054 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20055 := Call(__e, PrimFunc(symshen_4assoc_1_6), symdestroy, tmp20050, tmp20054)


tmp20056 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20055)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20055
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20056

tmp20057 := MakeNative(func(__e *ControlFlow) {
V6223 := __e.Get(1)
_ = V6223
__e.Return(MakeNative(func(__e *ControlFlow) {
B6219 := __e.Get(1)
_ = B6219
__e.Return(MakeNative(func(__e *ControlFlow) {
L6220 := __e.Get(1)
_ = L6220
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6221 := __e.Get(1)
_ = Key6221
__e.Return(MakeNative(func(__e *ControlFlow) {
C6222 := __e.Get(1)
_ = C6222
tmp20058 := Call(__e, PrimFunc(symshen_4newpv), B6219)


let__18947 := tmp20058
_ = let__18947

tmp20059 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18947, Nil)
}
__typedArg0 := let__18947
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20060 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20059)
}
__typedArg0 := symlist
__typedArg1 := tmp20059
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20061 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18947, Nil)
}
__typedArg0 := let__18947
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20062 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20061)
}
__typedArg0 := symlist
__typedArg1 := tmp20061
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20063 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18947, Nil)
}
__typedArg0 := let__18947
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20064 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20063)
}
__typedArg0 := symlist
__typedArg1 := tmp20063
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20065 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20064, Nil)
}
__typedArg0 := tmp20064
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20066 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20065)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20065
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20067 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20062, tmp20066)
}
__typedArg0 := tmp20062
__typedArg1 := tmp20066
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20068 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20067, Nil)
}
__typedArg0 := tmp20067
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20069 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20068)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20068
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20070 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20060, tmp20069)
}
__typedArg0 := tmp20060
__typedArg1 := tmp20069
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20071 := Call(__e, PrimFunc(symis_b), V6223, tmp20070, B6219, L6220, Key6221, C6222)


__e.TailApply(PrimFunc(symshen_4gc), B6219, tmp20071)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20072 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20073 := Call(__e, PrimFunc(symshen_4assoc_1_6), symdifference, tmp20057, tmp20072)


tmp20074 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20073)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20073
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20074

tmp20075 := MakeNative(func(__e *ControlFlow) {
V6228 := __e.Get(1)
_ = V6228
__e.Return(MakeNative(func(__e *ControlFlow) {
B6224 := __e.Get(1)
_ = B6224
__e.Return(MakeNative(func(__e *ControlFlow) {
L6225 := __e.Get(1)
_ = L6225
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6226 := __e.Get(1)
_ = Key6226
__e.Return(MakeNative(func(__e *ControlFlow) {
C6227 := __e.Get(1)
_ = C6227
tmp20076 := Call(__e, PrimFunc(symshen_4newpv), B6224)


let__18948 := tmp20076
_ = let__18948

tmp20077 := Call(__e, PrimFunc(symshen_4newpv), B6224)


let__18949 := tmp20077
_ = let__18949

tmp20078 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18949, Nil)
}
__typedArg0 := let__18949
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20079 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20078)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20078
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20080 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18949, tmp20079)
}
__typedArg0 := let__18949
__typedArg1 := tmp20079
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20081 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20080, Nil)
}
__typedArg0 := tmp20080
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20082 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20081)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20081
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20083 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18948, tmp20082)
}
__typedArg0 := let__18948
__typedArg1 := tmp20082
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20084 := Call(__e, PrimFunc(symis_b), V6228, tmp20083, B6224, L6225, Key6226, C6227)


tmp20085 := Call(__e, PrimFunc(symshen_4gc), B6224, tmp20084)


__e.TailApply(PrimFunc(symshen_4gc), B6224, tmp20085)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20086 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20087 := Call(__e, PrimFunc(symshen_4assoc_1_6), symdo, tmp20075, tmp20086)


tmp20088 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20087)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20087
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20088

tmp20089 := MakeNative(func(__e *ControlFlow) {
V6235 := __e.Get(1)
_ = V6235
__e.Return(MakeNative(func(__e *ControlFlow) {
B6231 := __e.Get(1)
_ = B6231
__e.Return(MakeNative(func(__e *ControlFlow) {
L6232 := __e.Get(1)
_ = L6232
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6233 := __e.Get(1)
_ = Key6233
__e.Return(MakeNative(func(__e *ControlFlow) {
C6234 := __e.Get(1)
_ = C6234
tmp20090 := Call(__e, PrimFunc(symshen_4newpv), B6231)


let__18950 := tmp20090
_ = let__18950

tmp20091 := Call(__e, PrimFunc(symshen_4newpv), B6231)


let__18951 := tmp20091
_ = let__18951

tmp20092 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18950, Nil)
}
__typedArg0 := let__18950
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20093 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20092)
}
__typedArg0 := symlist
__typedArg1 := tmp20092
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20094 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18950, Nil)
}
__typedArg0 := let__18950
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20095 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20094)
}
__typedArg0 := symlist
__typedArg1 := tmp20094
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20096 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18951, Nil)
}
__typedArg0 := let__18951
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20097 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20096)
}
__typedArg0 := symlist
__typedArg1 := tmp20096
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20098 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20097, Nil)
}
__typedArg0 := tmp20097
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20099 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20095, tmp20098)
}
__typedArg0 := tmp20095
__typedArg1 := tmp20098
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20100 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstr, tmp20099)
}
__typedArg0 := symstr
__typedArg1 := tmp20099
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20101 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20100, Nil)
}
__typedArg0 := tmp20100
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20102 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20101)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20101
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20103 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20093, tmp20102)
}
__typedArg0 := tmp20093
__typedArg1 := tmp20102
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20104 := Call(__e, PrimFunc(symis_b), V6235, tmp20103, B6231, L6232, Key6233, C6234)


tmp20105 := Call(__e, PrimFunc(symshen_4gc), B6231, tmp20104)


__e.TailApply(PrimFunc(symshen_4gc), B6231, tmp20105)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20106 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20107 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_5e_6, tmp20089, tmp20106)


tmp20108 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20107)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20107
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20108

tmp20109 := MakeNative(func(__e *ControlFlow) {
V6242 := __e.Get(1)
_ = V6242
__e.Return(MakeNative(func(__e *ControlFlow) {
B6238 := __e.Get(1)
_ = B6238
__e.Return(MakeNative(func(__e *ControlFlow) {
L6239 := __e.Get(1)
_ = L6239
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6240 := __e.Get(1)
_ = Key6240
__e.Return(MakeNative(func(__e *ControlFlow) {
C6241 := __e.Get(1)
_ = C6241
tmp20110 := Call(__e, PrimFunc(symshen_4newpv), B6238)


let__18952 := tmp20110
_ = let__18952

tmp20111 := Call(__e, PrimFunc(symshen_4newpv), B6238)


let__18953 := tmp20111
_ = let__18953

tmp20112 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18953, Nil)
}
__typedArg0 := let__18953
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20113 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20112)
}
__typedArg0 := symlist
__typedArg1 := tmp20112
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20114 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18952, Nil)
}
__typedArg0 := let__18952
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20115 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20114)
}
__typedArg0 := symlist
__typedArg1 := tmp20114
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20116 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18953, Nil)
}
__typedArg0 := let__18953
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20117 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20116)
}
__typedArg0 := symlist
__typedArg1 := tmp20116
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20118 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20117, Nil)
}
__typedArg0 := tmp20117
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20119 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20115, tmp20118)
}
__typedArg0 := tmp20115
__typedArg1 := tmp20118
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20120 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstr, tmp20119)
}
__typedArg0 := symstr
__typedArg1 := tmp20119
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20121 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20120, Nil)
}
__typedArg0 := tmp20120
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20122 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20121)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20121
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20123 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20113, tmp20122)
}
__typedArg0 := tmp20113
__typedArg1 := tmp20122
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20124 := Call(__e, PrimFunc(symis_b), V6242, tmp20123, B6238, L6239, Key6240, C6241)


tmp20125 := Call(__e, PrimFunc(symshen_4gc), B6238, tmp20124)


__e.TailApply(PrimFunc(symshen_4gc), B6238, tmp20125)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20126 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20127 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_5_b_6, tmp20109, tmp20126)


tmp20128 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20127)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20127
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20128

tmp20129 := MakeNative(func(__e *ControlFlow) {
V6249 := __e.Get(1)
_ = V6249
__e.Return(MakeNative(func(__e *ControlFlow) {
B6245 := __e.Get(1)
_ = B6245
__e.Return(MakeNative(func(__e *ControlFlow) {
L6246 := __e.Get(1)
_ = L6246
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6247 := __e.Get(1)
_ = Key6247
__e.Return(MakeNative(func(__e *ControlFlow) {
C6248 := __e.Get(1)
_ = C6248
tmp20130 := Call(__e, PrimFunc(symshen_4newpv), B6245)


let__18954 := tmp20130
_ = let__18954

tmp20131 := Call(__e, PrimFunc(symshen_4newpv), B6245)


let__18955 := tmp20131
_ = let__18955

tmp20132 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18954, Nil)
}
__typedArg0 := let__18954
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20133 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20132)
}
__typedArg0 := symlist
__typedArg1 := tmp20132
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20134 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18954, Nil)
}
__typedArg0 := let__18954
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20135 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20134)
}
__typedArg0 := symlist
__typedArg1 := tmp20134
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20136 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18955, Nil)
}
__typedArg0 := let__18955
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20137 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20136)
}
__typedArg0 := symlist
__typedArg1 := tmp20136
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20138 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20137, Nil)
}
__typedArg0 := tmp20137
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20139 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20135, tmp20138)
}
__typedArg0 := tmp20135
__typedArg1 := tmp20138
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20140 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstr, tmp20139)
}
__typedArg0 := symstr
__typedArg1 := tmp20139
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20141 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20140, Nil)
}
__typedArg0 := tmp20140
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20142 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20141)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20141
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20143 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20133, tmp20142)
}
__typedArg0 := tmp20133
__typedArg1 := tmp20142
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20144 := Call(__e, PrimFunc(symis_b), V6249, tmp20143, B6245, L6246, Key6247, C6248)


tmp20145 := Call(__e, PrimFunc(symshen_4gc), B6245, tmp20144)


__e.TailApply(PrimFunc(symshen_4gc), B6245, tmp20145)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20146 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20147 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_5end_6, tmp20129, tmp20146)


tmp20148 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20147)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20147
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20148

tmp20149 := MakeNative(func(__e *ControlFlow) {
V6254 := __e.Get(1)
_ = V6254
__e.Return(MakeNative(func(__e *ControlFlow) {
B6250 := __e.Get(1)
_ = B6250
__e.Return(MakeNative(func(__e *ControlFlow) {
L6251 := __e.Get(1)
_ = L6251
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6252 := __e.Get(1)
_ = Key6252
__e.Return(MakeNative(func(__e *ControlFlow) {
C6253 := __e.Get(1)
_ = C6253
tmp20150 := Call(__e, PrimFunc(symshen_4newpv), B6250)


let__18956 := tmp20150
_ = let__18956

tmp20151 := Call(__e, PrimFunc(symshen_4newpv), B6250)


let__18957 := tmp20151
_ = let__18957

tmp20152 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18956, Nil)
}
__typedArg0 := let__18956
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20153 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20152)
}
__typedArg0 := symlist
__typedArg1 := tmp20152
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20154 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18957, Nil)
}
__typedArg0 := let__18957
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20155 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20153, tmp20154)
}
__typedArg0 := tmp20153
__typedArg1 := tmp20154
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20156 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstr, tmp20155)
}
__typedArg0 := symstr
__typedArg1 := tmp20155
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20157 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20158 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20157)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20157
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20159 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20156, tmp20158)
}
__typedArg0 := tmp20156
__typedArg1 := tmp20158
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20160 := Call(__e, PrimFunc(symis_b), V6254, tmp20159, B6250, L6251, Key6252, C6253)


tmp20161 := Call(__e, PrimFunc(symshen_4gc), B6250, tmp20160)


__e.TailApply(PrimFunc(symshen_4gc), B6250, tmp20161)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20162 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20163 := Call(__e, PrimFunc(symshen_4assoc_1_6), symshen_4parse_1failure_2, tmp20149, tmp20162)


tmp20164 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20163)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20163
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20164

tmp20165 := MakeNative(func(__e *ControlFlow) {
V6259 := __e.Get(1)
_ = V6259
__e.Return(MakeNative(func(__e *ControlFlow) {
B6255 := __e.Get(1)
_ = B6255
__e.Return(MakeNative(func(__e *ControlFlow) {
L6256 := __e.Get(1)
_ = L6256
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6257 := __e.Get(1)
_ = Key6257
__e.Return(MakeNative(func(__e *ControlFlow) {
C6258 := __e.Get(1)
_ = C6258
tmp20166 := Call(__e, PrimFunc(symshen_4newpv), B6255)


let__18958 := tmp20166
_ = let__18958

tmp20167 := Call(__e, PrimFunc(symshen_4newpv), B6255)


let__18959 := tmp20167
_ = let__18959

tmp20168 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18958, Nil)
}
__typedArg0 := let__18958
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20169 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20168)
}
__typedArg0 := symlist
__typedArg1 := tmp20168
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20170 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18959, Nil)
}
__typedArg0 := let__18959
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20171 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20169, tmp20170)
}
__typedArg0 := tmp20169
__typedArg1 := tmp20170
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20172 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstr, tmp20171)
}
__typedArg0 := symstr
__typedArg1 := tmp20171
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20173 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20172, Nil)
}
__typedArg0 := tmp20172
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20174 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20173)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20173
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20175 := Call(__e, PrimFunc(symis_b), V6259, tmp20174, B6255, L6256, Key6257, C6258)


tmp20176 := Call(__e, PrimFunc(symshen_4gc), B6255, tmp20175)


__e.TailApply(PrimFunc(symshen_4gc), B6255, tmp20176)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20177 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20178 := Call(__e, PrimFunc(symshen_4assoc_1_6), symshen_4parse_1failure, tmp20165, tmp20177)


tmp20179 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20178)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20178
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20179

tmp20180 := MakeNative(func(__e *ControlFlow) {
V6264 := __e.Get(1)
_ = V6264
__e.Return(MakeNative(func(__e *ControlFlow) {
B6260 := __e.Get(1)
_ = B6260
__e.Return(MakeNative(func(__e *ControlFlow) {
L6261 := __e.Get(1)
_ = L6261
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6262 := __e.Get(1)
_ = Key6262
__e.Return(MakeNative(func(__e *ControlFlow) {
C6263 := __e.Get(1)
_ = C6263
tmp20181 := Call(__e, PrimFunc(symshen_4newpv), B6260)


let__18960 := tmp20181
_ = let__18960

tmp20182 := Call(__e, PrimFunc(symshen_4newpv), B6260)


let__18961 := tmp20182
_ = let__18961

tmp20183 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18960, Nil)
}
__typedArg0 := let__18960
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20184 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20183)
}
__typedArg0 := symlist
__typedArg1 := tmp20183
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20185 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18961, Nil)
}
__typedArg0 := let__18961
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20186 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20184, tmp20185)
}
__typedArg0 := tmp20184
__typedArg1 := tmp20185
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20187 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstr, tmp20186)
}
__typedArg0 := symstr
__typedArg1 := tmp20186
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20188 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18961, Nil)
}
__typedArg0 := let__18961
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20189 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20188)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20188
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20190 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20187, tmp20189)
}
__typedArg0 := tmp20187
__typedArg1 := tmp20189
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20191 := Call(__e, PrimFunc(symis_b), V6264, tmp20190, B6260, L6261, Key6262, C6263)


tmp20192 := Call(__e, PrimFunc(symshen_4gc), B6260, tmp20191)


__e.TailApply(PrimFunc(symshen_4gc), B6260, tmp20192)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20193 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20194 := Call(__e, PrimFunc(symshen_4assoc_1_6), symshen_4_5_1out, tmp20180, tmp20193)


tmp20195 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20194)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20194
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20195

tmp20196 := MakeNative(func(__e *ControlFlow) {
V6270 := __e.Get(1)
_ = V6270
__e.Return(MakeNative(func(__e *ControlFlow) {
B6266 := __e.Get(1)
_ = B6266
__e.Return(MakeNative(func(__e *ControlFlow) {
L6267 := __e.Get(1)
_ = L6267
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6268 := __e.Get(1)
_ = Key6268
__e.Return(MakeNative(func(__e *ControlFlow) {
C6269 := __e.Get(1)
_ = C6269
tmp20197 := Call(__e, PrimFunc(symshen_4newpv), B6266)


let__18962 := tmp20197
_ = let__18962

tmp20198 := Call(__e, PrimFunc(symshen_4newpv), B6266)


let__18963 := tmp20198
_ = let__18963

tmp20199 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18963, Nil)
}
__typedArg0 := let__18963
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20200 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20199)
}
__typedArg0 := symlist
__typedArg1 := tmp20199
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20201 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18962, Nil)
}
__typedArg0 := let__18962
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20202 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20200, tmp20201)
}
__typedArg0 := tmp20200
__typedArg1 := tmp20201
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20203 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstr, tmp20202)
}
__typedArg0 := symstr
__typedArg1 := tmp20202
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20204 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18963, Nil)
}
__typedArg0 := let__18963
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20205 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20204)
}
__typedArg0 := symlist
__typedArg1 := tmp20204
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20206 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20205, Nil)
}
__typedArg0 := tmp20205
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20207 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20206)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20206
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20208 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20203, tmp20207)
}
__typedArg0 := tmp20203
__typedArg1 := tmp20207
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20209 := Call(__e, PrimFunc(symis_b), V6270, tmp20208, B6266, L6267, Key6268, C6269)


tmp20210 := Call(__e, PrimFunc(symshen_4gc), B6266, tmp20209)


__e.TailApply(PrimFunc(symshen_4gc), B6266, tmp20210)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20211 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20212 := Call(__e, PrimFunc(symshen_4assoc_1_6), symshen_4in_1_6, tmp20196, tmp20211)


tmp20213 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20212)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20212
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20213

tmp20214 := MakeNative(func(__e *ControlFlow) {
V6276 := __e.Get(1)
_ = V6276
__e.Return(MakeNative(func(__e *ControlFlow) {
B6272 := __e.Get(1)
_ = B6272
__e.Return(MakeNative(func(__e *ControlFlow) {
L6273 := __e.Get(1)
_ = L6273
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6274 := __e.Get(1)
_ = Key6274
__e.Return(MakeNative(func(__e *ControlFlow) {
C6275 := __e.Get(1)
_ = C6275
tmp20215 := Call(__e, PrimFunc(symshen_4newpv), B6272)


let__18964 := tmp20215
_ = let__18964

tmp20216 := Call(__e, PrimFunc(symshen_4newpv), B6272)


let__18965 := tmp20216
_ = let__18965

tmp20217 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18964, Nil)
}
__typedArg0 := let__18964
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20218 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20217)
}
__typedArg0 := symlist
__typedArg1 := tmp20217
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20219 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18964, Nil)
}
__typedArg0 := let__18964
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20220 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20219)
}
__typedArg0 := symlist
__typedArg1 := tmp20219
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20221 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18965, Nil)
}
__typedArg0 := let__18965
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20222 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20220, tmp20221)
}
__typedArg0 := tmp20220
__typedArg1 := tmp20221
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20223 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstr, tmp20222)
}
__typedArg0 := symstr
__typedArg1 := tmp20222
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20224 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20223, Nil)
}
__typedArg0 := tmp20223
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20225 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20224)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20224
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20226 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18965, tmp20225)
}
__typedArg0 := let__18965
__typedArg1 := tmp20225
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20227 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20226, Nil)
}
__typedArg0 := tmp20226
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20228 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20227)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20227
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20229 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20218, tmp20228)
}
__typedArg0 := tmp20218
__typedArg1 := tmp20228
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20230 := Call(__e, PrimFunc(symis_b), V6276, tmp20229, B6272, L6273, Key6274, C6275)


tmp20231 := Call(__e, PrimFunc(symshen_4gc), B6272, tmp20230)


__e.TailApply(PrimFunc(symshen_4gc), B6272, tmp20231)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20232 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20233 := Call(__e, PrimFunc(symshen_4assoc_1_6), symshen_4comb, tmp20214, tmp20232)


tmp20234 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20233)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20233
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20234

tmp20235 := MakeNative(func(__e *ControlFlow) {
V6281 := __e.Get(1)
_ = V6281
__e.Return(MakeNative(func(__e *ControlFlow) {
B6277 := __e.Get(1)
_ = B6277
__e.Return(MakeNative(func(__e *ControlFlow) {
L6278 := __e.Get(1)
_ = L6278
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6279 := __e.Get(1)
_ = Key6279
__e.Return(MakeNative(func(__e *ControlFlow) {
C6280 := __e.Get(1)
_ = C6280
tmp20236 := Call(__e, PrimFunc(symshen_4newpv), B6277)


let__18966 := tmp20236
_ = let__18966

tmp20237 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18966, Nil)
}
__typedArg0 := let__18966
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20238 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20237)
}
__typedArg0 := symlist
__typedArg1 := tmp20237
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20239 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20240 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20239)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20239
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20241 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20238, tmp20240)
}
__typedArg0 := tmp20238
__typedArg1 := tmp20240
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20242 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20241, Nil)
}
__typedArg0 := tmp20241
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20243 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20242)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20242
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20244 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18966, tmp20243)
}
__typedArg0 := let__18966
__typedArg1 := tmp20243
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20245 := Call(__e, PrimFunc(symis_b), V6281, tmp20244, B6277, L6278, Key6279, C6280)


__e.TailApply(PrimFunc(symshen_4gc), B6277, tmp20245)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20246 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20247 := Call(__e, PrimFunc(symshen_4assoc_1_6), symelement_2, tmp20235, tmp20246)


tmp20248 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20247)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20247
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20248

tmp20249 := MakeNative(func(__e *ControlFlow) {
V6286 := __e.Get(1)
_ = V6286
__e.Return(MakeNative(func(__e *ControlFlow) {
B6282 := __e.Get(1)
_ = B6282
__e.Return(MakeNative(func(__e *ControlFlow) {
L6283 := __e.Get(1)
_ = L6283
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6284 := __e.Get(1)
_ = Key6284
__e.Return(MakeNative(func(__e *ControlFlow) {
C6285 := __e.Get(1)
_ = C6285
tmp20250 := Call(__e, PrimFunc(symshen_4newpv), B6282)


let__18967 := tmp20250
_ = let__18967

tmp20251 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20252 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20251)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20251
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20253 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18967, tmp20252)
}
__typedArg0 := let__18967
__typedArg1 := tmp20252
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20254 := Call(__e, PrimFunc(symis_b), V6286, tmp20253, B6282, L6283, Key6284, C6285)


__e.TailApply(PrimFunc(symshen_4gc), B6282, tmp20254)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20255 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20256 := Call(__e, PrimFunc(symshen_4assoc_1_6), symempty_2, tmp20249, tmp20255)


tmp20257 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20256)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20256
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20257

tmp20258 := MakeNative(func(__e *ControlFlow) {
V6291 := __e.Get(1)
_ = V6291
__e.Return(MakeNative(func(__e *ControlFlow) {
B6287 := __e.Get(1)
_ = B6287
__e.Return(MakeNative(func(__e *ControlFlow) {
L6288 := __e.Get(1)
_ = L6288
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6289 := __e.Get(1)
_ = Key6289
__e.Return(MakeNative(func(__e *ControlFlow) {
C6290 := __e.Get(1)
_ = C6290
tmp20259 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20260 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20259)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20259
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20261 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp20260)
}
__typedArg0 := symsymbol
__typedArg1 := tmp20260
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6291, tmp20261, B6287, L6288, Key6289, C6290)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20262 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20263 := Call(__e, PrimFunc(symshen_4assoc_1_6), symenable_1type_1theory, tmp20258, tmp20262)


tmp20264 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20263)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20263
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20264

tmp20265 := MakeNative(func(__e *ControlFlow) {
V6296 := __e.Get(1)
_ = V6296
__e.Return(MakeNative(func(__e *ControlFlow) {
B6292 := __e.Get(1)
_ = B6292
__e.Return(MakeNative(func(__e *ControlFlow) {
L6293 := __e.Get(1)
_ = L6293
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6294 := __e.Get(1)
_ = Key6294
__e.Return(MakeNative(func(__e *ControlFlow) {
C6295 := __e.Get(1)
_ = C6295
tmp20266 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20267 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20266)
}
__typedArg0 := symlist
__typedArg1 := tmp20266
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20268 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20267, Nil)
}
__typedArg0 := tmp20267
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20269 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20268)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20268
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20270 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp20269)
}
__typedArg0 := symsymbol
__typedArg1 := tmp20269
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6296, tmp20270, B6292, L6293, Key6294, C6295)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20271 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20272 := Call(__e, PrimFunc(symshen_4assoc_1_6), symexternal, tmp20265, tmp20271)


tmp20273 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20272)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20272
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20273

tmp20274 := MakeNative(func(__e *ControlFlow) {
V6301 := __e.Get(1)
_ = V6301
__e.Return(MakeNative(func(__e *ControlFlow) {
B6297 := __e.Get(1)
_ = B6297
__e.Return(MakeNative(func(__e *ControlFlow) {
L6298 := __e.Get(1)
_ = L6298
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6299 := __e.Get(1)
_ = Key6299
__e.Return(MakeNative(func(__e *ControlFlow) {
C6300 := __e.Get(1)
_ = C6300
tmp20275 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20276 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20275)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20275
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20277 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symexception, tmp20276)
}
__typedArg0 := symexception
__typedArg1 := tmp20276
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6301, tmp20277, B6297, L6298, Key6299, C6300)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20278 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20279 := Call(__e, PrimFunc(symshen_4assoc_1_6), symerror_1to_1string, tmp20274, tmp20278)


tmp20280 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20279)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20279
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20280

tmp20281 := MakeNative(func(__e *ControlFlow) {
V6306 := __e.Get(1)
_ = V6306
__e.Return(MakeNative(func(__e *ControlFlow) {
B6302 := __e.Get(1)
_ = B6302
__e.Return(MakeNative(func(__e *ControlFlow) {
L6303 := __e.Get(1)
_ = L6303
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6304 := __e.Get(1)
_ = Key6304
__e.Return(MakeNative(func(__e *ControlFlow) {
C6305 := __e.Get(1)
_ = C6305
tmp20282 := Call(__e, PrimFunc(symshen_4newpv), B6302)


let__18968 := tmp20282
_ = let__18968

tmp20283 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20284 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20283)
}
__typedArg0 := symlist
__typedArg1 := tmp20283
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20285 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20284, Nil)
}
__typedArg0 := tmp20284
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20286 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20285)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20285
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20287 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18968, tmp20286)
}
__typedArg0 := let__18968
__typedArg1 := tmp20286
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20288 := Call(__e, PrimFunc(symis_b), V6306, tmp20287, B6302, L6303, Key6304, C6305)


__e.TailApply(PrimFunc(symshen_4gc), B6302, tmp20288)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20289 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20290 := Call(__e, PrimFunc(symshen_4assoc_1_6), symexplode, tmp20281, tmp20289)


tmp20291 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20290)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20290
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20291

tmp20292 := MakeNative(func(__e *ControlFlow) {
V6311 := __e.Get(1)
_ = V6311
__e.Return(MakeNative(func(__e *ControlFlow) {
B6307 := __e.Get(1)
_ = B6307
__e.Return(MakeNative(func(__e *ControlFlow) {
L6308 := __e.Get(1)
_ = L6308
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6309 := __e.Get(1)
_ = Key6309
__e.Return(MakeNative(func(__e *ControlFlow) {
C6310 := __e.Get(1)
_ = C6310
tmp20293 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20294 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20293)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20293
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20295 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp20294)
}
__typedArg0 := symsymbol
__typedArg1 := tmp20294
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6311, tmp20295, B6307, L6308, Key6309, C6310)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20296 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20297 := Call(__e, PrimFunc(symshen_4assoc_1_6), symfactorise, tmp20292, tmp20296)


tmp20298 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20297)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20297
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20298

tmp20299 := MakeNative(func(__e *ControlFlow) {
V6316 := __e.Get(1)
_ = V6316
__e.Return(MakeNative(func(__e *ControlFlow) {
B6312 := __e.Get(1)
_ = B6312
__e.Return(MakeNative(func(__e *ControlFlow) {
L6313 := __e.Get(1)
_ = L6313
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6314 := __e.Get(1)
_ = Key6314
__e.Return(MakeNative(func(__e *ControlFlow) {
C6315 := __e.Get(1)
_ = C6315
tmp20300 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20301 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20300)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20300
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6316, tmp20301, B6312, L6313, Key6314, C6315)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20302 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20303 := Call(__e, PrimFunc(symshen_4assoc_1_6), symfactorise_2, tmp20299, tmp20302)


tmp20304 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20303)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20303
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20304

tmp20305 := MakeNative(func(__e *ControlFlow) {
V6321 := __e.Get(1)
_ = V6321
__e.Return(MakeNative(func(__e *ControlFlow) {
B6317 := __e.Get(1)
_ = B6317
__e.Return(MakeNative(func(__e *ControlFlow) {
L6318 := __e.Get(1)
_ = L6318
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6319 := __e.Get(1)
_ = Key6319
__e.Return(MakeNative(func(__e *ControlFlow) {
C6320 := __e.Get(1)
_ = C6320
tmp20306 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20307 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20306)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20306
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6321, tmp20307, B6317, L6318, Key6319, C6320)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20308 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20309 := Call(__e, PrimFunc(symshen_4assoc_1_6), symfail, tmp20305, tmp20308)


tmp20310 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20309)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20309
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20310

tmp20311 := MakeNative(func(__e *ControlFlow) {
V6327 := __e.Get(1)
_ = V6327
__e.Return(MakeNative(func(__e *ControlFlow) {
B6323 := __e.Get(1)
_ = B6323
__e.Return(MakeNative(func(__e *ControlFlow) {
L6324 := __e.Get(1)
_ = L6324
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6325 := __e.Get(1)
_ = Key6325
__e.Return(MakeNative(func(__e *ControlFlow) {
C6326 := __e.Get(1)
_ = C6326
tmp20312 := Call(__e, PrimFunc(symshen_4newpv), B6323)


let__18969 := tmp20312
_ = let__18969

tmp20313 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18969, Nil)
}
__typedArg0 := let__18969
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20314 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20313)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20313
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20315 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18969, tmp20314)
}
__typedArg0 := let__18969
__typedArg1 := tmp20314
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20316 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18969, Nil)
}
__typedArg0 := let__18969
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20317 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20316)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20316
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20318 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18969, tmp20317)
}
__typedArg0 := let__18969
__typedArg1 := tmp20317
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20319 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20318, Nil)
}
__typedArg0 := tmp20318
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20320 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20319)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20319
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20321 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20315, tmp20320)
}
__typedArg0 := tmp20315
__typedArg1 := tmp20320
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20322 := Call(__e, PrimFunc(symis_b), V6327, tmp20321, B6323, L6324, Key6325, C6326)


__e.TailApply(PrimFunc(symshen_4gc), B6323, tmp20322)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20323 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20324 := Call(__e, PrimFunc(symshen_4assoc_1_6), symfix, tmp20311, tmp20323)


tmp20325 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20324)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20324
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20325

tmp20326 := MakeNative(func(__e *ControlFlow) {
V6332 := __e.Get(1)
_ = V6332
__e.Return(MakeNative(func(__e *ControlFlow) {
B6328 := __e.Get(1)
_ = B6328
__e.Return(MakeNative(func(__e *ControlFlow) {
L6329 := __e.Get(1)
_ = L6329
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6330 := __e.Get(1)
_ = Key6330
__e.Return(MakeNative(func(__e *ControlFlow) {
C6331 := __e.Get(1)
_ = C6331
tmp20327 := Call(__e, PrimFunc(symshen_4newpv), B6328)


let__18970 := tmp20327
_ = let__18970

tmp20328 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18970, Nil)
}
__typedArg0 := let__18970
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20329 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlazy, tmp20328)
}
__typedArg0 := symlazy
__typedArg1 := tmp20328
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20330 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20329, Nil)
}
__typedArg0 := tmp20329
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20331 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20330)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20330
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20332 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18970, tmp20331)
}
__typedArg0 := let__18970
__typedArg1 := tmp20331
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20333 := Call(__e, PrimFunc(symis_b), V6332, tmp20332, B6328, L6329, Key6330, C6331)


__e.TailApply(PrimFunc(symshen_4gc), B6328, tmp20333)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20334 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20335 := Call(__e, PrimFunc(symshen_4assoc_1_6), symfreeze, tmp20326, tmp20334)


tmp20336 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20335)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20335
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20336

tmp20337 := MakeNative(func(__e *ControlFlow) {
V6337 := __e.Get(1)
_ = V6337
__e.Return(MakeNative(func(__e *ControlFlow) {
B6333 := __e.Get(1)
_ = B6333
__e.Return(MakeNative(func(__e *ControlFlow) {
L6334 := __e.Get(1)
_ = L6334
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6335 := __e.Get(1)
_ = Key6335
__e.Return(MakeNative(func(__e *ControlFlow) {
C6336 := __e.Get(1)
_ = C6336
tmp20338 := Call(__e, PrimFunc(symshen_4newpv), B6333)


let__18971 := tmp20338
_ = let__18971

tmp20339 := Call(__e, PrimFunc(symshen_4newpv), B6333)


let__18972 := tmp20339
_ = let__18972

tmp20340 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18971, Nil)
}
__typedArg0 := let__18971
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20341 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_d, tmp20340)
}
__typedArg0 := sym_d
__typedArg1 := tmp20340
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20342 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18972, tmp20341)
}
__typedArg0 := let__18972
__typedArg1 := tmp20341
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20343 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18972, Nil)
}
__typedArg0 := let__18972
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20344 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20343)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20343
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20345 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20342, tmp20344)
}
__typedArg0 := tmp20342
__typedArg1 := tmp20344
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20346 := Call(__e, PrimFunc(symis_b), V6337, tmp20345, B6333, L6334, Key6335, C6336)


tmp20347 := Call(__e, PrimFunc(symshen_4gc), B6333, tmp20346)


__e.TailApply(PrimFunc(symshen_4gc), B6333, tmp20347)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20348 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20349 := Call(__e, PrimFunc(symshen_4assoc_1_6), symfst, tmp20337, tmp20348)


tmp20350 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20349)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20349
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20350

tmp20351 := MakeNative(func(__e *ControlFlow) {
V6342 := __e.Get(1)
_ = V6342
__e.Return(MakeNative(func(__e *ControlFlow) {
B6338 := __e.Get(1)
_ = B6338
__e.Return(MakeNative(func(__e *ControlFlow) {
L6339 := __e.Get(1)
_ = L6339
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6340 := __e.Get(1)
_ = Key6340
__e.Return(MakeNative(func(__e *ControlFlow) {
C6341 := __e.Get(1)
_ = C6341
tmp20352 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20353 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20352)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20352
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20354 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp20353)
}
__typedArg0 := symsymbol
__typedArg1 := tmp20353
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6342, tmp20354, B6338, L6339, Key6340, C6341)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20355 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20356 := Call(__e, PrimFunc(symshen_4assoc_1_6), symgensym, tmp20351, tmp20355)


tmp20357 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20356)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20356
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20357

tmp20358 := MakeNative(func(__e *ControlFlow) {
V6348 := __e.Get(1)
_ = V6348
__e.Return(MakeNative(func(__e *ControlFlow) {
B6344 := __e.Get(1)
_ = B6344
__e.Return(MakeNative(func(__e *ControlFlow) {
L6345 := __e.Get(1)
_ = L6345
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6346 := __e.Get(1)
_ = Key6346
__e.Return(MakeNative(func(__e *ControlFlow) {
C6347 := __e.Get(1)
_ = C6347
tmp20359 := Call(__e, PrimFunc(symshen_4newpv), B6344)


let__18973 := tmp20359
_ = let__18973

tmp20360 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18973, Nil)
}
__typedArg0 := let__18973
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20361 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20360)
}
__typedArg0 := symlist
__typedArg1 := tmp20360
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20362 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20363 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20362)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20362
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20364 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18973, tmp20363)
}
__typedArg0 := let__18973
__typedArg1 := tmp20363
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20365 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20364, Nil)
}
__typedArg0 := tmp20364
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20366 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20365)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20365
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20367 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20361, tmp20366)
}
__typedArg0 := tmp20361
__typedArg1 := tmp20366
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20368 := Call(__e, PrimFunc(symis_b), V6348, tmp20367, B6344, L6345, Key6346, C6347)


__e.TailApply(PrimFunc(symshen_4gc), B6344, tmp20368)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20369 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20370 := Call(__e, PrimFunc(symshen_4assoc_1_6), symshen_4hds_a_2, tmp20358, tmp20369)


tmp20371 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20370)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20370
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20371

tmp20372 := MakeNative(func(__e *ControlFlow) {
V6353 := __e.Get(1)
_ = V6353
__e.Return(MakeNative(func(__e *ControlFlow) {
B6349 := __e.Get(1)
_ = B6349
__e.Return(MakeNative(func(__e *ControlFlow) {
L6350 := __e.Get(1)
_ = L6350
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6351 := __e.Get(1)
_ = Key6351
__e.Return(MakeNative(func(__e *ControlFlow) {
C6352 := __e.Get(1)
_ = C6352
tmp20373 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20374 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20373)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20373
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20375 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp20374)
}
__typedArg0 := symsymbol
__typedArg1 := tmp20374
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6353, tmp20375, B6349, L6350, Key6351, C6352)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20376 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20377 := Call(__e, PrimFunc(symshen_4assoc_1_6), symhush, tmp20372, tmp20376)


tmp20378 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20377)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20377
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20378

tmp20379 := MakeNative(func(__e *ControlFlow) {
V6358 := __e.Get(1)
_ = V6358
__e.Return(MakeNative(func(__e *ControlFlow) {
B6354 := __e.Get(1)
_ = B6354
__e.Return(MakeNative(func(__e *ControlFlow) {
L6355 := __e.Get(1)
_ = L6355
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6356 := __e.Get(1)
_ = Key6356
__e.Return(MakeNative(func(__e *ControlFlow) {
C6357 := __e.Get(1)
_ = C6357
tmp20380 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20381 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20380)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20380
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6358, tmp20381, B6354, L6355, Key6356, C6357)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20382 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20383 := Call(__e, PrimFunc(symshen_4assoc_1_6), symhush_2, tmp20379, tmp20382)


tmp20384 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20383)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20383
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20384

tmp20385 := MakeNative(func(__e *ControlFlow) {
V6364 := __e.Get(1)
_ = V6364
__e.Return(MakeNative(func(__e *ControlFlow) {
B6360 := __e.Get(1)
_ = B6360
__e.Return(MakeNative(func(__e *ControlFlow) {
L6361 := __e.Get(1)
_ = L6361
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6362 := __e.Get(1)
_ = Key6362
__e.Return(MakeNative(func(__e *ControlFlow) {
C6363 := __e.Get(1)
_ = C6363
tmp20386 := Call(__e, PrimFunc(symshen_4newpv), B6360)


let__18974 := tmp20386
_ = let__18974

tmp20387 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18974, Nil)
}
__typedArg0 := let__18974
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20388 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp20387)
}
__typedArg0 := symvector
__typedArg1 := tmp20387
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20389 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18974, Nil)
}
__typedArg0 := let__18974
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20390 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20389)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20389
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20391 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp20390)
}
__typedArg0 := symnumber
__typedArg1 := tmp20390
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20392 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20391, Nil)
}
__typedArg0 := tmp20391
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20393 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20392)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20392
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20394 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20388, tmp20393)
}
__typedArg0 := tmp20388
__typedArg1 := tmp20393
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20395 := Call(__e, PrimFunc(symis_b), V6364, tmp20394, B6360, L6361, Key6362, C6363)


__e.TailApply(PrimFunc(symshen_4gc), B6360, tmp20395)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20396 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20397 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_5_1vector, tmp20385, tmp20396)


tmp20398 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20397)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20397
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20398

tmp20399 := MakeNative(func(__e *ControlFlow) {
V6370 := __e.Get(1)
_ = V6370
__e.Return(MakeNative(func(__e *ControlFlow) {
B6366 := __e.Get(1)
_ = B6366
__e.Return(MakeNative(func(__e *ControlFlow) {
L6367 := __e.Get(1)
_ = L6367
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6368 := __e.Get(1)
_ = Key6368
__e.Return(MakeNative(func(__e *ControlFlow) {
C6369 := __e.Get(1)
_ = C6369
tmp20400 := Call(__e, PrimFunc(symshen_4newpv), B6366)


let__18975 := tmp20400
_ = let__18975

tmp20401 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18975, Nil)
}
__typedArg0 := let__18975
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20402 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp20401)
}
__typedArg0 := symvector
__typedArg1 := tmp20401
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20403 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18975, Nil)
}
__typedArg0 := let__18975
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20404 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp20403)
}
__typedArg0 := symvector
__typedArg1 := tmp20403
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20405 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20404, Nil)
}
__typedArg0 := tmp20404
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20406 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20405)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20405
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20407 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18975, tmp20406)
}
__typedArg0 := let__18975
__typedArg1 := tmp20406
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20408 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20407, Nil)
}
__typedArg0 := tmp20407
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20409 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20408)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20408
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20410 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp20409)
}
__typedArg0 := symnumber
__typedArg1 := tmp20409
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20411 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20410, Nil)
}
__typedArg0 := tmp20410
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20412 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20411)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20411
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20413 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20402, tmp20412)
}
__typedArg0 := tmp20402
__typedArg1 := tmp20412
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20414 := Call(__e, PrimFunc(symis_b), V6370, tmp20413, B6366, L6367, Key6368, C6369)


__e.TailApply(PrimFunc(symshen_4gc), B6366, tmp20414)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20415 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20416 := Call(__e, PrimFunc(symshen_4assoc_1_6), symvector_1_6, tmp20399, tmp20415)


tmp20417 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20416)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20416
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20417

tmp20418 := MakeNative(func(__e *ControlFlow) {
V6375 := __e.Get(1)
_ = V6375
__e.Return(MakeNative(func(__e *ControlFlow) {
B6371 := __e.Get(1)
_ = B6371
__e.Return(MakeNative(func(__e *ControlFlow) {
L6372 := __e.Get(1)
_ = L6372
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6373 := __e.Get(1)
_ = Key6373
__e.Return(MakeNative(func(__e *ControlFlow) {
C6374 := __e.Get(1)
_ = C6374
tmp20419 := Call(__e, PrimFunc(symshen_4newpv), B6371)


let__18976 := tmp20419
_ = let__18976

tmp20420 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18976, Nil)
}
__typedArg0 := let__18976
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20421 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp20420)
}
__typedArg0 := symvector
__typedArg1 := tmp20420
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20422 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20421, Nil)
}
__typedArg0 := tmp20421
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20423 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20422)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20422
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20424 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp20423)
}
__typedArg0 := symnumber
__typedArg1 := tmp20423
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20425 := Call(__e, PrimFunc(symis_b), V6375, tmp20424, B6371, L6372, Key6373, C6374)


__e.TailApply(PrimFunc(symshen_4gc), B6371, tmp20425)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20426 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20427 := Call(__e, PrimFunc(symshen_4assoc_1_6), symvector, tmp20418, tmp20426)


tmp20428 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20427)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20427
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20428

tmp20429 := MakeNative(func(__e *ControlFlow) {
V6380 := __e.Get(1)
_ = V6380
__e.Return(MakeNative(func(__e *ControlFlow) {
B6376 := __e.Get(1)
_ = B6376
__e.Return(MakeNative(func(__e *ControlFlow) {
L6377 := __e.Get(1)
_ = L6377
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6378 := __e.Get(1)
_ = Key6378
__e.Return(MakeNative(func(__e *ControlFlow) {
C6379 := __e.Get(1)
_ = C6379
tmp20430 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20431 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20430)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20430
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20432 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp20431)
}
__typedArg0 := symsymbol
__typedArg1 := tmp20431
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6380, tmp20432, B6376, L6377, Key6378, C6379)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20433 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20434 := Call(__e, PrimFunc(symshen_4assoc_1_6), symget_1time, tmp20429, tmp20433)


tmp20435 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20434)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20434
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20435

tmp20436 := MakeNative(func(__e *ControlFlow) {
V6385 := __e.Get(1)
_ = V6385
__e.Return(MakeNative(func(__e *ControlFlow) {
B6381 := __e.Get(1)
_ = B6381
__e.Return(MakeNative(func(__e *ControlFlow) {
L6382 := __e.Get(1)
_ = L6382
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6383 := __e.Get(1)
_ = Key6383
__e.Return(MakeNative(func(__e *ControlFlow) {
C6384 := __e.Get(1)
_ = C6384
tmp20437 := Call(__e, PrimFunc(symshen_4newpv), B6381)


let__18977 := tmp20437
_ = let__18977

tmp20438 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20439 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20438)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20438
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20440 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp20439)
}
__typedArg0 := symnumber
__typedArg1 := tmp20439
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20441 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20440, Nil)
}
__typedArg0 := tmp20440
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20442 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20441)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20441
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20443 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18977, tmp20442)
}
__typedArg0 := let__18977
__typedArg1 := tmp20442
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20444 := Call(__e, PrimFunc(symis_b), V6385, tmp20443, B6381, L6382, Key6383, C6384)


__e.TailApply(PrimFunc(symshen_4gc), B6381, tmp20444)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20445 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20446 := Call(__e, PrimFunc(symshen_4assoc_1_6), symhash, tmp20436, tmp20445)


tmp20447 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20446)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20446
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20447

tmp20448 := MakeNative(func(__e *ControlFlow) {
V6390 := __e.Get(1)
_ = V6390
__e.Return(MakeNative(func(__e *ControlFlow) {
B6386 := __e.Get(1)
_ = B6386
__e.Return(MakeNative(func(__e *ControlFlow) {
L6387 := __e.Get(1)
_ = L6387
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6388 := __e.Get(1)
_ = Key6388
__e.Return(MakeNative(func(__e *ControlFlow) {
C6389 := __e.Get(1)
_ = C6389
tmp20449 := Call(__e, PrimFunc(symshen_4newpv), B6386)


let__18978 := tmp20449
_ = let__18978

tmp20450 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18978, Nil)
}
__typedArg0 := let__18978
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20451 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20450)
}
__typedArg0 := symlist
__typedArg1 := tmp20450
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20452 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18978, Nil)
}
__typedArg0 := let__18978
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20453 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20452)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20452
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20454 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20451, tmp20453)
}
__typedArg0 := tmp20451
__typedArg1 := tmp20453
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20455 := Call(__e, PrimFunc(symis_b), V6390, tmp20454, B6386, L6387, Key6388, C6389)


__e.TailApply(PrimFunc(symshen_4gc), B6386, tmp20455)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20456 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20457 := Call(__e, PrimFunc(symshen_4assoc_1_6), symhead, tmp20448, tmp20456)


tmp20458 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20457)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20457
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20458

tmp20459 := MakeNative(func(__e *ControlFlow) {
V6395 := __e.Get(1)
_ = V6395
__e.Return(MakeNative(func(__e *ControlFlow) {
B6391 := __e.Get(1)
_ = B6391
__e.Return(MakeNative(func(__e *ControlFlow) {
L6392 := __e.Get(1)
_ = L6392
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6393 := __e.Get(1)
_ = Key6393
__e.Return(MakeNative(func(__e *ControlFlow) {
C6394 := __e.Get(1)
_ = C6394
tmp20460 := Call(__e, PrimFunc(symshen_4newpv), B6391)


let__18979 := tmp20460
_ = let__18979

tmp20461 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18979, Nil)
}
__typedArg0 := let__18979
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20462 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp20461)
}
__typedArg0 := symvector
__typedArg1 := tmp20461
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20463 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18979, Nil)
}
__typedArg0 := let__18979
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20464 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20463)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20463
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20465 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20462, tmp20464)
}
__typedArg0 := tmp20462
__typedArg1 := tmp20464
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20466 := Call(__e, PrimFunc(symis_b), V6395, tmp20465, B6391, L6392, Key6393, C6394)


__e.TailApply(PrimFunc(symshen_4gc), B6391, tmp20466)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20467 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20468 := Call(__e, PrimFunc(symshen_4assoc_1_6), symhdv, tmp20459, tmp20467)


tmp20469 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20468)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20468
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20469

tmp20470 := MakeNative(func(__e *ControlFlow) {
V6400 := __e.Get(1)
_ = V6400
__e.Return(MakeNative(func(__e *ControlFlow) {
B6396 := __e.Get(1)
_ = B6396
__e.Return(MakeNative(func(__e *ControlFlow) {
L6397 := __e.Get(1)
_ = L6397
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6398 := __e.Get(1)
_ = Key6398
__e.Return(MakeNative(func(__e *ControlFlow) {
C6399 := __e.Get(1)
_ = C6399
tmp20471 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20472 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20471)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20471
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20473 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp20472)
}
__typedArg0 := symstring
__typedArg1 := tmp20472
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6400, tmp20473, B6396, L6397, Key6398, C6399)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20474 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20475 := Call(__e, PrimFunc(symshen_4assoc_1_6), symhdstr, tmp20470, tmp20474)


tmp20476 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20475)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20475
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20476

tmp20477 := MakeNative(func(__e *ControlFlow) {
V6405 := __e.Get(1)
_ = V6405
__e.Return(MakeNative(func(__e *ControlFlow) {
B6401 := __e.Get(1)
_ = B6401
__e.Return(MakeNative(func(__e *ControlFlow) {
L6402 := __e.Get(1)
_ = L6402
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6403 := __e.Get(1)
_ = Key6403
__e.Return(MakeNative(func(__e *ControlFlow) {
C6404 := __e.Get(1)
_ = C6404
tmp20478 := Call(__e, PrimFunc(symshen_4newpv), B6401)


let__18980 := tmp20478
_ = let__18980

tmp20479 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18980, Nil)
}
__typedArg0 := let__18980
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20480 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20479)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20479
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20481 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18980, tmp20480)
}
__typedArg0 := let__18980
__typedArg1 := tmp20480
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20482 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20481, Nil)
}
__typedArg0 := tmp20481
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20483 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20482)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20482
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20484 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18980, tmp20483)
}
__typedArg0 := let__18980
__typedArg1 := tmp20483
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20485 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20484, Nil)
}
__typedArg0 := tmp20484
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20486 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20485)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20485
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20487 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, tmp20486)
}
__typedArg0 := symboolean
__typedArg1 := tmp20486
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20488 := Call(__e, PrimFunc(symis_b), V6405, tmp20487, B6401, L6402, Key6403, C6404)


__e.TailApply(PrimFunc(symshen_4gc), B6401, tmp20488)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20489 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20490 := Call(__e, PrimFunc(symshen_4assoc_1_6), symif, tmp20477, tmp20489)


tmp20491 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20490)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20490
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20491

tmp20492 := MakeNative(func(__e *ControlFlow) {
V6410 := __e.Get(1)
_ = V6410
__e.Return(MakeNative(func(__e *ControlFlow) {
B6406 := __e.Get(1)
_ = B6406
__e.Return(MakeNative(func(__e *ControlFlow) {
L6407 := __e.Get(1)
_ = L6407
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6408 := __e.Get(1)
_ = Key6408
__e.Return(MakeNative(func(__e *ControlFlow) {
C6409 := __e.Get(1)
_ = C6409
tmp20493 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20494 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20493)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20493
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20495 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp20494)
}
__typedArg0 := symsymbol
__typedArg1 := tmp20494
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6410, tmp20495, B6406, L6407, Key6408, C6409)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20496 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20497 := Call(__e, PrimFunc(symshen_4assoc_1_6), symin_1package, tmp20492, tmp20496)


tmp20498 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20497)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20497
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20498

tmp20499 := MakeNative(func(__e *ControlFlow) {
V6415 := __e.Get(1)
_ = V6415
__e.Return(MakeNative(func(__e *ControlFlow) {
B6411 := __e.Get(1)
_ = B6411
__e.Return(MakeNative(func(__e *ControlFlow) {
L6412 := __e.Get(1)
_ = L6412
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6413 := __e.Get(1)
_ = Key6413
__e.Return(MakeNative(func(__e *ControlFlow) {
C6414 := __e.Get(1)
_ = C6414
tmp20500 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20501 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20500)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20500
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6415, tmp20501, B6411, L6412, Key6413, C6414)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20502 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20503 := Call(__e, PrimFunc(symshen_4assoc_1_6), symit, tmp20499, tmp20502)


tmp20504 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20503)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20503
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20504

tmp20505 := MakeNative(func(__e *ControlFlow) {
V6420 := __e.Get(1)
_ = V6420
__e.Return(MakeNative(func(__e *ControlFlow) {
B6416 := __e.Get(1)
_ = B6416
__e.Return(MakeNative(func(__e *ControlFlow) {
L6417 := __e.Get(1)
_ = L6417
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6418 := __e.Get(1)
_ = Key6418
__e.Return(MakeNative(func(__e *ControlFlow) {
C6419 := __e.Get(1)
_ = C6419
tmp20506 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20507 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20506)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20506
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6420, tmp20507, B6416, L6417, Key6418, C6419)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20508 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20509 := Call(__e, PrimFunc(symshen_4assoc_1_6), symimplementation, tmp20505, tmp20508)


tmp20510 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20509)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20509
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20510

tmp20511 := MakeNative(func(__e *ControlFlow) {
V6426 := __e.Get(1)
_ = V6426
__e.Return(MakeNative(func(__e *ControlFlow) {
B6422 := __e.Get(1)
_ = B6422
__e.Return(MakeNative(func(__e *ControlFlow) {
L6423 := __e.Get(1)
_ = L6423
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6424 := __e.Get(1)
_ = Key6424
__e.Return(MakeNative(func(__e *ControlFlow) {
C6425 := __e.Get(1)
_ = C6425
tmp20512 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20513 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20512)
}
__typedArg0 := symlist
__typedArg1 := tmp20512
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20514 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20515 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20514)
}
__typedArg0 := symlist
__typedArg1 := tmp20514
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20516 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20515, Nil)
}
__typedArg0 := tmp20515
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20517 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20516)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20516
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20518 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20513, tmp20517)
}
__typedArg0 := tmp20513
__typedArg1 := tmp20517
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6426, tmp20518, B6422, L6423, Key6424, C6425)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20519 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20520 := Call(__e, PrimFunc(symshen_4assoc_1_6), syminclude, tmp20511, tmp20519)


tmp20521 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20520)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20520
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20521

tmp20522 := MakeNative(func(__e *ControlFlow) {
V6432 := __e.Get(1)
_ = V6432
__e.Return(MakeNative(func(__e *ControlFlow) {
B6428 := __e.Get(1)
_ = B6428
__e.Return(MakeNative(func(__e *ControlFlow) {
L6429 := __e.Get(1)
_ = L6429
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6430 := __e.Get(1)
_ = Key6430
__e.Return(MakeNative(func(__e *ControlFlow) {
C6431 := __e.Get(1)
_ = C6431
tmp20523 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20524 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20523)
}
__typedArg0 := symlist
__typedArg1 := tmp20523
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20525 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20526 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20525)
}
__typedArg0 := symlist
__typedArg1 := tmp20525
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20527 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20526, Nil)
}
__typedArg0 := tmp20526
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20528 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20527)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20527
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20529 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20524, tmp20528)
}
__typedArg0 := tmp20524
__typedArg1 := tmp20528
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6432, tmp20529, B6428, L6429, Key6430, C6431)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20530 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20531 := Call(__e, PrimFunc(symshen_4assoc_1_6), syminclude_1all_1but, tmp20522, tmp20530)


tmp20532 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20531)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20531
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20532

tmp20533 := MakeNative(func(__e *ControlFlow) {
V6437 := __e.Get(1)
_ = V6437
__e.Return(MakeNative(func(__e *ControlFlow) {
B6433 := __e.Get(1)
_ = B6433
__e.Return(MakeNative(func(__e *ControlFlow) {
L6434 := __e.Get(1)
_ = L6434
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6435 := __e.Get(1)
_ = Key6435
__e.Return(MakeNative(func(__e *ControlFlow) {
C6436 := __e.Get(1)
_ = C6436
tmp20534 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20535 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20534)
}
__typedArg0 := symlist
__typedArg1 := tmp20534
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20536 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20535, Nil)
}
__typedArg0 := tmp20535
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20537 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20536)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20536
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6437, tmp20537, B6433, L6434, Key6435, C6436)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20538 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20539 := Call(__e, PrimFunc(symshen_4assoc_1_6), symincluded, tmp20533, tmp20538)


tmp20540 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20539)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20539
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20540

tmp20541 := MakeNative(func(__e *ControlFlow) {
V6442 := __e.Get(1)
_ = V6442
__e.Return(MakeNative(func(__e *ControlFlow) {
B6438 := __e.Get(1)
_ = B6438
__e.Return(MakeNative(func(__e *ControlFlow) {
L6439 := __e.Get(1)
_ = L6439
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6440 := __e.Get(1)
_ = Key6440
__e.Return(MakeNative(func(__e *ControlFlow) {
C6441 := __e.Get(1)
_ = C6441
tmp20542 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20543 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20542)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20542
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6442, tmp20543, B6438, L6439, Key6440, C6441)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20544 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20545 := Call(__e, PrimFunc(symshen_4assoc_1_6), syminferences, tmp20541, tmp20544)


tmp20546 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20545)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20545
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20546

tmp20547 := MakeNative(func(__e *ControlFlow) {
V6447 := __e.Get(1)
_ = V6447
__e.Return(MakeNative(func(__e *ControlFlow) {
B6443 := __e.Get(1)
_ = B6443
__e.Return(MakeNative(func(__e *ControlFlow) {
L6444 := __e.Get(1)
_ = L6444
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6445 := __e.Get(1)
_ = Key6445
__e.Return(MakeNative(func(__e *ControlFlow) {
C6446 := __e.Get(1)
_ = C6446
tmp20548 := Call(__e, PrimFunc(symshen_4newpv), B6443)


let__18981 := tmp20548
_ = let__18981

tmp20549 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20550 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20549)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20549
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20551 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp20550)
}
__typedArg0 := symstring
__typedArg1 := tmp20550
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20552 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20551, Nil)
}
__typedArg0 := tmp20551
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20553 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20552)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20552
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20554 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18981, tmp20553)
}
__typedArg0 := let__18981
__typedArg1 := tmp20553
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20555 := Call(__e, PrimFunc(symis_b), V6447, tmp20554, B6443, L6444, Key6445, C6446)


__e.TailApply(PrimFunc(symshen_4gc), B6443, tmp20555)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20556 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20557 := Call(__e, PrimFunc(symshen_4assoc_1_6), symshen_4insert, tmp20547, tmp20556)


tmp20558 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20557)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20557
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20558

tmp20559 := MakeNative(func(__e *ControlFlow) {
V6452 := __e.Get(1)
_ = V6452
__e.Return(MakeNative(func(__e *ControlFlow) {
B6448 := __e.Get(1)
_ = B6448
__e.Return(MakeNative(func(__e *ControlFlow) {
L6449 := __e.Get(1)
_ = L6449
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6450 := __e.Get(1)
_ = Key6450
__e.Return(MakeNative(func(__e *ControlFlow) {
C6451 := __e.Get(1)
_ = C6451
tmp20560 := Call(__e, PrimFunc(symshen_4newpv), B6448)


let__18982 := tmp20560
_ = let__18982

tmp20561 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20562 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20561)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20561
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20563 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18982, tmp20562)
}
__typedArg0 := let__18982
__typedArg1 := tmp20562
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20564 := Call(__e, PrimFunc(symis_b), V6452, tmp20563, B6448, L6449, Key6450, C6451)


__e.TailApply(PrimFunc(symshen_4gc), B6448, tmp20564)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20565 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20566 := Call(__e, PrimFunc(symshen_4assoc_1_6), syminteger_2, tmp20559, tmp20565)


tmp20567 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20566)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20566
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20567

tmp20568 := MakeNative(func(__e *ControlFlow) {
V6457 := __e.Get(1)
_ = V6457
__e.Return(MakeNative(func(__e *ControlFlow) {
B6453 := __e.Get(1)
_ = B6453
__e.Return(MakeNative(func(__e *ControlFlow) {
L6454 := __e.Get(1)
_ = L6454
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6455 := __e.Get(1)
_ = Key6455
__e.Return(MakeNative(func(__e *ControlFlow) {
C6456 := __e.Get(1)
_ = C6456
tmp20569 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20570 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20569)
}
__typedArg0 := symlist
__typedArg1 := tmp20569
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20571 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20570, Nil)
}
__typedArg0 := tmp20570
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20572 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20571)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20571
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20573 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp20572)
}
__typedArg0 := symsymbol
__typedArg1 := tmp20572
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6457, tmp20573, B6453, L6454, Key6455, C6456)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20574 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20575 := Call(__e, PrimFunc(symshen_4assoc_1_6), syminternal, tmp20568, tmp20574)


tmp20576 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20575)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20575
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20576

tmp20577 := MakeNative(func(__e *ControlFlow) {
V6464 := __e.Get(1)
_ = V6464
__e.Return(MakeNative(func(__e *ControlFlow) {
B6460 := __e.Get(1)
_ = B6460
__e.Return(MakeNative(func(__e *ControlFlow) {
L6461 := __e.Get(1)
_ = L6461
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6462 := __e.Get(1)
_ = Key6462
__e.Return(MakeNative(func(__e *ControlFlow) {
C6463 := __e.Get(1)
_ = C6463
tmp20578 := Call(__e, PrimFunc(symshen_4newpv), B6460)


let__18983 := tmp20578
_ = let__18983

tmp20579 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18983, Nil)
}
__typedArg0 := let__18983
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20580 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20579)
}
__typedArg0 := symlist
__typedArg1 := tmp20579
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20581 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18983, Nil)
}
__typedArg0 := let__18983
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20582 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20581)
}
__typedArg0 := symlist
__typedArg1 := tmp20581
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20583 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18983, Nil)
}
__typedArg0 := let__18983
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20584 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20583)
}
__typedArg0 := symlist
__typedArg1 := tmp20583
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20585 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20584, Nil)
}
__typedArg0 := tmp20584
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20586 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20585)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20585
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20587 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20582, tmp20586)
}
__typedArg0 := tmp20582
__typedArg1 := tmp20586
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20588 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20587, Nil)
}
__typedArg0 := tmp20587
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20589 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20588)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20588
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20590 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20580, tmp20589)
}
__typedArg0 := tmp20580
__typedArg1 := tmp20589
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20591 := Call(__e, PrimFunc(symis_b), V6464, tmp20590, B6460, L6461, Key6462, C6463)


__e.TailApply(PrimFunc(symshen_4gc), B6460, tmp20591)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20592 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20593 := Call(__e, PrimFunc(symshen_4assoc_1_6), symintersection, tmp20577, tmp20592)


tmp20594 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20593)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20593
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20594

tmp20595 := MakeNative(func(__e *ControlFlow) {
V6469 := __e.Get(1)
_ = V6469
__e.Return(MakeNative(func(__e *ControlFlow) {
B6465 := __e.Get(1)
_ = B6465
__e.Return(MakeNative(func(__e *ControlFlow) {
L6466 := __e.Get(1)
_ = L6466
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6467 := __e.Get(1)
_ = Key6467
__e.Return(MakeNative(func(__e *ControlFlow) {
C6468 := __e.Get(1)
_ = C6468
tmp20596 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20597 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20596)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20596
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6469, tmp20597, B6465, L6466, Key6467, C6468)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20598 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20599 := Call(__e, PrimFunc(symshen_4assoc_1_6), symlanguage, tmp20595, tmp20598)


tmp20600 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20599)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20599
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20600

tmp20601 := MakeNative(func(__e *ControlFlow) {
V6474 := __e.Get(1)
_ = V6474
__e.Return(MakeNative(func(__e *ControlFlow) {
B6470 := __e.Get(1)
_ = B6470
__e.Return(MakeNative(func(__e *ControlFlow) {
L6471 := __e.Get(1)
_ = L6471
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6472 := __e.Get(1)
_ = Key6472
__e.Return(MakeNative(func(__e *ControlFlow) {
C6473 := __e.Get(1)
_ = C6473
tmp20602 := Call(__e, PrimFunc(symshen_4newpv), B6470)


let__18984 := tmp20602
_ = let__18984

tmp20603 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18984, Nil)
}
__typedArg0 := let__18984
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20604 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20603)
}
__typedArg0 := symlist
__typedArg1 := tmp20603
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20605 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20606 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20605)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20605
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20607 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20604, tmp20606)
}
__typedArg0 := tmp20604
__typedArg1 := tmp20606
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20608 := Call(__e, PrimFunc(symis_b), V6474, tmp20607, B6470, L6471, Key6472, C6473)


__e.TailApply(PrimFunc(symshen_4gc), B6470, tmp20608)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20609 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20610 := Call(__e, PrimFunc(symshen_4assoc_1_6), symlength, tmp20601, tmp20609)


tmp20611 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20610)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20610
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20611

tmp20612 := MakeNative(func(__e *ControlFlow) {
V6479 := __e.Get(1)
_ = V6479
__e.Return(MakeNative(func(__e *ControlFlow) {
B6475 := __e.Get(1)
_ = B6475
__e.Return(MakeNative(func(__e *ControlFlow) {
L6476 := __e.Get(1)
_ = L6476
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6477 := __e.Get(1)
_ = Key6477
__e.Return(MakeNative(func(__e *ControlFlow) {
C6478 := __e.Get(1)
_ = C6478
tmp20613 := Call(__e, PrimFunc(symshen_4newpv), B6475)


let__18985 := tmp20613
_ = let__18985

tmp20614 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18985, Nil)
}
__typedArg0 := let__18985
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20615 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp20614)
}
__typedArg0 := symvector
__typedArg1 := tmp20614
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20616 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20617 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20616)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20616
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20618 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20615, tmp20617)
}
__typedArg0 := tmp20615
__typedArg1 := tmp20617
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20619 := Call(__e, PrimFunc(symis_b), V6479, tmp20618, B6475, L6476, Key6477, C6478)


__e.TailApply(PrimFunc(symshen_4gc), B6475, tmp20619)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20620 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20621 := Call(__e, PrimFunc(symshen_4assoc_1_6), symlimit, tmp20612, tmp20620)


tmp20622 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20621)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20621
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20622

tmp20623 := MakeNative(func(__e *ControlFlow) {
V6485 := __e.Get(1)
_ = V6485
__e.Return(MakeNative(func(__e *ControlFlow) {
B6481 := __e.Get(1)
_ = B6481
__e.Return(MakeNative(func(__e *ControlFlow) {
L6482 := __e.Get(1)
_ = L6482
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6483 := __e.Get(1)
_ = Key6483
__e.Return(MakeNative(func(__e *ControlFlow) {
C6484 := __e.Get(1)
_ = C6484
tmp20624 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symin, Nil)
}
__typedArg0 := symin
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20625 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp20624)
}
__typedArg0 := symstream
__typedArg1 := tmp20624
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20626 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunit, Nil)
}
__typedArg0 := symunit
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20627 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20626)
}
__typedArg0 := symlist
__typedArg1 := tmp20626
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20628 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20627, Nil)
}
__typedArg0 := tmp20627
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20629 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20628)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20628
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20630 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20625, tmp20629)
}
__typedArg0 := tmp20625
__typedArg1 := tmp20629
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6485, tmp20630, B6481, L6482, Key6483, C6484)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20631 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20632 := Call(__e, PrimFunc(symshen_4assoc_1_6), symlineread, tmp20623, tmp20631)


tmp20633 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20632)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20632
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20633

tmp20634 := MakeNative(func(__e *ControlFlow) {
V6490 := __e.Get(1)
_ = V6490
__e.Return(MakeNative(func(__e *ControlFlow) {
B6486 := __e.Get(1)
_ = B6486
__e.Return(MakeNative(func(__e *ControlFlow) {
L6487 := __e.Get(1)
_ = L6487
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6488 := __e.Get(1)
_ = Key6488
__e.Return(MakeNative(func(__e *ControlFlow) {
C6489 := __e.Get(1)
_ = C6489
tmp20635 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20636 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20635)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20635
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20637 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp20636)
}
__typedArg0 := symstring
__typedArg1 := tmp20636
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6490, tmp20637, B6486, L6487, Key6488, C6489)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20638 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20639 := Call(__e, PrimFunc(symshen_4assoc_1_6), symload, tmp20634, tmp20638)


tmp20640 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20639)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20639
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20640

tmp20641 := MakeNative(func(__e *ControlFlow) {
V6497 := __e.Get(1)
_ = V6497
__e.Return(MakeNative(func(__e *ControlFlow) {
B6493 := __e.Get(1)
_ = B6493
__e.Return(MakeNative(func(__e *ControlFlow) {
L6494 := __e.Get(1)
_ = L6494
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6495 := __e.Get(1)
_ = Key6495
__e.Return(MakeNative(func(__e *ControlFlow) {
C6496 := __e.Get(1)
_ = C6496
tmp20642 := Call(__e, PrimFunc(symshen_4newpv), B6493)


let__18986 := tmp20642
_ = let__18986

tmp20643 := Call(__e, PrimFunc(symshen_4newpv), B6493)


let__18987 := tmp20643
_ = let__18987

tmp20644 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18987, Nil)
}
__typedArg0 := let__18987
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20645 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20644)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20644
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20646 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18986, tmp20645)
}
__typedArg0 := let__18986
__typedArg1 := tmp20645
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20647 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18986, Nil)
}
__typedArg0 := let__18986
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20648 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20647)
}
__typedArg0 := symlist
__typedArg1 := tmp20647
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20649 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18987, Nil)
}
__typedArg0 := let__18987
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20650 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20649)
}
__typedArg0 := symlist
__typedArg1 := tmp20649
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20651 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20650, Nil)
}
__typedArg0 := tmp20650
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20652 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20651)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20651
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20653 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20648, tmp20652)
}
__typedArg0 := tmp20648
__typedArg1 := tmp20652
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20654 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20653, Nil)
}
__typedArg0 := tmp20653
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20655 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20654)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20654
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20656 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20646, tmp20655)
}
__typedArg0 := tmp20646
__typedArg1 := tmp20655
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20657 := Call(__e, PrimFunc(symis_b), V6497, tmp20656, B6493, L6494, Key6495, C6496)


tmp20658 := Call(__e, PrimFunc(symshen_4gc), B6493, tmp20657)


__e.TailApply(PrimFunc(symshen_4gc), B6493, tmp20658)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20659 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20660 := Call(__e, PrimFunc(symshen_4assoc_1_6), symmap, tmp20641, tmp20659)


tmp20661 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20660)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20660
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20661

tmp20662 := MakeNative(func(__e *ControlFlow) {
V6504 := __e.Get(1)
_ = V6504
__e.Return(MakeNative(func(__e *ControlFlow) {
B6500 := __e.Get(1)
_ = B6500
__e.Return(MakeNative(func(__e *ControlFlow) {
L6501 := __e.Get(1)
_ = L6501
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6502 := __e.Get(1)
_ = Key6502
__e.Return(MakeNative(func(__e *ControlFlow) {
C6503 := __e.Get(1)
_ = C6503
tmp20663 := Call(__e, PrimFunc(symshen_4newpv), B6500)


let__18988 := tmp20663
_ = let__18988

tmp20664 := Call(__e, PrimFunc(symshen_4newpv), B6500)


let__18989 := tmp20664
_ = let__18989

tmp20665 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18989, Nil)
}
__typedArg0 := let__18989
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20666 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20665)
}
__typedArg0 := symlist
__typedArg1 := tmp20665
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20667 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20666, Nil)
}
__typedArg0 := tmp20666
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20668 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20667)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20667
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20669 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18988, tmp20668)
}
__typedArg0 := let__18988
__typedArg1 := tmp20668
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20670 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18988, Nil)
}
__typedArg0 := let__18988
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20671 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20670)
}
__typedArg0 := symlist
__typedArg1 := tmp20670
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20672 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18989, Nil)
}
__typedArg0 := let__18989
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20673 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20672)
}
__typedArg0 := symlist
__typedArg1 := tmp20672
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20674 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20673, Nil)
}
__typedArg0 := tmp20673
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20675 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20674)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20674
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20676 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20671, tmp20675)
}
__typedArg0 := tmp20671
__typedArg1 := tmp20675
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20677 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20676, Nil)
}
__typedArg0 := tmp20676
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20678 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20677)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20677
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20679 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20669, tmp20678)
}
__typedArg0 := tmp20669
__typedArg1 := tmp20678
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20680 := Call(__e, PrimFunc(symis_b), V6504, tmp20679, B6500, L6501, Key6502, C6503)


tmp20681 := Call(__e, PrimFunc(symshen_4gc), B6500, tmp20680)


__e.TailApply(PrimFunc(symshen_4gc), B6500, tmp20681)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20682 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20683 := Call(__e, PrimFunc(symshen_4assoc_1_6), symmapcan, tmp20662, tmp20682)


tmp20684 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20683)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20683
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20684

tmp20685 := MakeNative(func(__e *ControlFlow) {
V6509 := __e.Get(1)
_ = V6509
__e.Return(MakeNative(func(__e *ControlFlow) {
B6505 := __e.Get(1)
_ = B6505
__e.Return(MakeNative(func(__e *ControlFlow) {
L6506 := __e.Get(1)
_ = L6506
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6507 := __e.Get(1)
_ = Key6507
__e.Return(MakeNative(func(__e *ControlFlow) {
C6508 := __e.Get(1)
_ = C6508
tmp20686 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20687 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20686)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20686
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20688 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp20687)
}
__typedArg0 := symnumber
__typedArg1 := tmp20687
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6509, tmp20688, B6505, L6506, Key6507, C6508)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20689 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20690 := Call(__e, PrimFunc(symshen_4assoc_1_6), symmaxinferences, tmp20685, tmp20689)


tmp20691 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20690)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20690
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20691

tmp20692 := MakeNative(func(__e *ControlFlow) {
V6514 := __e.Get(1)
_ = V6514
__e.Return(MakeNative(func(__e *ControlFlow) {
B6510 := __e.Get(1)
_ = B6510
__e.Return(MakeNative(func(__e *ControlFlow) {
L6511 := __e.Get(1)
_ = L6511
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6512 := __e.Get(1)
_ = Key6512
__e.Return(MakeNative(func(__e *ControlFlow) {
C6513 := __e.Get(1)
_ = C6513
tmp20693 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20694 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20693)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20693
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20695 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp20694)
}
__typedArg0 := symnumber
__typedArg1 := tmp20694
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6514, tmp20695, B6510, L6511, Key6512, C6513)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20696 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20697 := Call(__e, PrimFunc(symshen_4assoc_1_6), symn_1_6string, tmp20692, tmp20696)


tmp20698 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20697)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20697
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20698

tmp20699 := MakeNative(func(__e *ControlFlow) {
V6519 := __e.Get(1)
_ = V6519
__e.Return(MakeNative(func(__e *ControlFlow) {
B6515 := __e.Get(1)
_ = B6515
__e.Return(MakeNative(func(__e *ControlFlow) {
L6516 := __e.Get(1)
_ = L6516
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6517 := __e.Get(1)
_ = Key6517
__e.Return(MakeNative(func(__e *ControlFlow) {
C6518 := __e.Get(1)
_ = C6518
tmp20700 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20701 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20700)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20700
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20702 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp20701)
}
__typedArg0 := symnumber
__typedArg1 := tmp20701
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6519, tmp20702, B6515, L6516, Key6517, C6518)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20703 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20704 := Call(__e, PrimFunc(symshen_4assoc_1_6), symnl, tmp20699, tmp20703)


tmp20705 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20704)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20704
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20705

tmp20706 := MakeNative(func(__e *ControlFlow) {
V6524 := __e.Get(1)
_ = V6524
__e.Return(MakeNative(func(__e *ControlFlow) {
B6520 := __e.Get(1)
_ = B6520
__e.Return(MakeNative(func(__e *ControlFlow) {
L6521 := __e.Get(1)
_ = L6521
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6522 := __e.Get(1)
_ = Key6522
__e.Return(MakeNative(func(__e *ControlFlow) {
C6523 := __e.Get(1)
_ = C6523
tmp20707 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20708 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20707)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20707
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20709 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, tmp20708)
}
__typedArg0 := symboolean
__typedArg1 := tmp20708
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6524, tmp20709, B6520, L6521, Key6522, C6523)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20710 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20711 := Call(__e, PrimFunc(symshen_4assoc_1_6), symnot, tmp20706, tmp20710)


tmp20712 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20711)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20711
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20712

tmp20713 := MakeNative(func(__e *ControlFlow) {
V6529 := __e.Get(1)
_ = V6529
__e.Return(MakeNative(func(__e *ControlFlow) {
B6525 := __e.Get(1)
_ = B6525
__e.Return(MakeNative(func(__e *ControlFlow) {
L6526 := __e.Get(1)
_ = L6526
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6527 := __e.Get(1)
_ = Key6527
__e.Return(MakeNative(func(__e *ControlFlow) {
C6528 := __e.Get(1)
_ = C6528
tmp20714 := Call(__e, PrimFunc(symshen_4newpv), B6525)


let__18990 := tmp20714
_ = let__18990

tmp20715 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18990, Nil)
}
__typedArg0 := let__18990
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20716 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20715)
}
__typedArg0 := symlist
__typedArg1 := tmp20715
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20717 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18990, Nil)
}
__typedArg0 := let__18990
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20718 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20717)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20717
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20719 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20716, tmp20718)
}
__typedArg0 := tmp20716
__typedArg1 := tmp20718
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20720 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20719, Nil)
}
__typedArg0 := tmp20719
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20721 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20720)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20720
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20722 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp20721)
}
__typedArg0 := symnumber
__typedArg1 := tmp20721
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20723 := Call(__e, PrimFunc(symis_b), V6529, tmp20722, B6525, L6526, Key6527, C6528)


__e.TailApply(PrimFunc(symshen_4gc), B6525, tmp20723)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20724 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20725 := Call(__e, PrimFunc(symshen_4assoc_1_6), symnth, tmp20713, tmp20724)


tmp20726 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20725)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20725
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20726

tmp20727 := MakeNative(func(__e *ControlFlow) {
V6534 := __e.Get(1)
_ = V6534
__e.Return(MakeNative(func(__e *ControlFlow) {
B6530 := __e.Get(1)
_ = B6530
__e.Return(MakeNative(func(__e *ControlFlow) {
L6531 := __e.Get(1)
_ = L6531
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6532 := __e.Get(1)
_ = Key6532
__e.Return(MakeNative(func(__e *ControlFlow) {
C6533 := __e.Get(1)
_ = C6533
tmp20728 := Call(__e, PrimFunc(symshen_4newpv), B6530)


let__18991 := tmp20728
_ = let__18991

tmp20729 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20730 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20729)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20729
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20731 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18991, tmp20730)
}
__typedArg0 := let__18991
__typedArg1 := tmp20730
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20732 := Call(__e, PrimFunc(symis_b), V6534, tmp20731, B6530, L6531, Key6532, C6533)


__e.TailApply(PrimFunc(symshen_4gc), B6530, tmp20732)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20733 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20734 := Call(__e, PrimFunc(symshen_4assoc_1_6), symnumber_2, tmp20727, tmp20733)


tmp20735 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20734)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20734
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20735

tmp20736 := MakeNative(func(__e *ControlFlow) {
V6539 := __e.Get(1)
_ = V6539
__e.Return(MakeNative(func(__e *ControlFlow) {
B6535 := __e.Get(1)
_ = B6535
__e.Return(MakeNative(func(__e *ControlFlow) {
L6536 := __e.Get(1)
_ = L6536
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6537 := __e.Get(1)
_ = Key6537
__e.Return(MakeNative(func(__e *ControlFlow) {
C6538 := __e.Get(1)
_ = C6538
tmp20737 := Call(__e, PrimFunc(symshen_4newpv), B6535)


let__18992 := tmp20737
_ = let__18992

tmp20738 := Call(__e, PrimFunc(symshen_4newpv), B6535)


let__18993 := tmp20738
_ = let__18993

tmp20739 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20740 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20739)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20739
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20741 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18993, tmp20740)
}
__typedArg0 := let__18993
__typedArg1 := tmp20740
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20742 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20741, Nil)
}
__typedArg0 := tmp20741
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20743 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20742)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20742
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20744 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18992, tmp20743)
}
__typedArg0 := let__18992
__typedArg1 := tmp20743
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20745 := Call(__e, PrimFunc(symis_b), V6539, tmp20744, B6535, L6536, Key6537, C6538)


tmp20746 := Call(__e, PrimFunc(symshen_4gc), B6535, tmp20745)


__e.TailApply(PrimFunc(symshen_4gc), B6535, tmp20746)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20747 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20748 := Call(__e, PrimFunc(symshen_4assoc_1_6), symoccurrences, tmp20736, tmp20747)


tmp20749 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20748)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20748
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20749

tmp20750 := MakeNative(func(__e *ControlFlow) {
V6544 := __e.Get(1)
_ = V6544
__e.Return(MakeNative(func(__e *ControlFlow) {
B6540 := __e.Get(1)
_ = B6540
__e.Return(MakeNative(func(__e *ControlFlow) {
L6541 := __e.Get(1)
_ = L6541
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6542 := __e.Get(1)
_ = Key6542
__e.Return(MakeNative(func(__e *ControlFlow) {
C6543 := __e.Get(1)
_ = C6543
tmp20751 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20752 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20751)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20751
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20753 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp20752)
}
__typedArg0 := symsymbol
__typedArg1 := tmp20752
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6544, tmp20753, B6540, L6541, Key6542, C6543)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20754 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20755 := Call(__e, PrimFunc(symshen_4assoc_1_6), symoccurs_1check, tmp20750, tmp20754)


tmp20756 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20755)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20755
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20756

tmp20757 := MakeNative(func(__e *ControlFlow) {
V6549 := __e.Get(1)
_ = V6549
__e.Return(MakeNative(func(__e *ControlFlow) {
B6545 := __e.Get(1)
_ = B6545
__e.Return(MakeNative(func(__e *ControlFlow) {
L6546 := __e.Get(1)
_ = L6546
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6547 := __e.Get(1)
_ = Key6547
__e.Return(MakeNative(func(__e *ControlFlow) {
C6548 := __e.Get(1)
_ = C6548
tmp20758 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20759 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20758)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20758
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6549, tmp20759, B6545, L6546, Key6547, C6548)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20760 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20761 := Call(__e, PrimFunc(symshen_4assoc_1_6), symoccurs_2, tmp20757, tmp20760)


tmp20762 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20761)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20761
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20762

tmp20763 := MakeNative(func(__e *ControlFlow) {
V6554 := __e.Get(1)
_ = V6554
__e.Return(MakeNative(func(__e *ControlFlow) {
B6550 := __e.Get(1)
_ = B6550
__e.Return(MakeNative(func(__e *ControlFlow) {
L6551 := __e.Get(1)
_ = L6551
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6552 := __e.Get(1)
_ = Key6552
__e.Return(MakeNative(func(__e *ControlFlow) {
C6553 := __e.Get(1)
_ = C6553
tmp20764 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20765 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20764)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20764
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20766 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp20765)
}
__typedArg0 := symsymbol
__typedArg1 := tmp20765
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6554, tmp20766, B6550, L6551, Key6552, C6553)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20767 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20768 := Call(__e, PrimFunc(symshen_4assoc_1_6), symoptimise, tmp20763, tmp20767)


tmp20769 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20768)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20768
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20769

tmp20770 := MakeNative(func(__e *ControlFlow) {
V6559 := __e.Get(1)
_ = V6559
__e.Return(MakeNative(func(__e *ControlFlow) {
B6555 := __e.Get(1)
_ = B6555
__e.Return(MakeNative(func(__e *ControlFlow) {
L6556 := __e.Get(1)
_ = L6556
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6557 := __e.Get(1)
_ = Key6557
__e.Return(MakeNative(func(__e *ControlFlow) {
C6558 := __e.Get(1)
_ = C6558
tmp20771 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20772 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20771)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20771
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6559, tmp20772, B6555, L6556, Key6557, C6558)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20773 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20774 := Call(__e, PrimFunc(symshen_4assoc_1_6), symoptimise_2, tmp20770, tmp20773)


tmp20775 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20774)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20774
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20775

tmp20776 := MakeNative(func(__e *ControlFlow) {
V6564 := __e.Get(1)
_ = V6564
__e.Return(MakeNative(func(__e *ControlFlow) {
B6560 := __e.Get(1)
_ = B6560
__e.Return(MakeNative(func(__e *ControlFlow) {
L6561 := __e.Get(1)
_ = L6561
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6562 := __e.Get(1)
_ = Key6562
__e.Return(MakeNative(func(__e *ControlFlow) {
C6563 := __e.Get(1)
_ = C6563
tmp20777 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20778 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20777)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20777
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20779 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, tmp20778)
}
__typedArg0 := symboolean
__typedArg1 := tmp20778
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20780 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20779, Nil)
}
__typedArg0 := tmp20779
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20781 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20780)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20780
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20782 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, tmp20781)
}
__typedArg0 := symboolean
__typedArg1 := tmp20781
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6564, tmp20782, B6560, L6561, Key6562, C6563)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20783 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20784 := Call(__e, PrimFunc(symshen_4assoc_1_6), symor, tmp20776, tmp20783)


tmp20785 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20784)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20784
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20785

tmp20786 := MakeNative(func(__e *ControlFlow) {
V6569 := __e.Get(1)
_ = V6569
__e.Return(MakeNative(func(__e *ControlFlow) {
B6565 := __e.Get(1)
_ = B6565
__e.Return(MakeNative(func(__e *ControlFlow) {
L6566 := __e.Get(1)
_ = L6566
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6567 := __e.Get(1)
_ = Key6567
__e.Return(MakeNative(func(__e *ControlFlow) {
C6568 := __e.Get(1)
_ = C6568
tmp20787 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20788 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20787)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20787
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6569, tmp20788, B6565, L6566, Key6567, C6568)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20789 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20790 := Call(__e, PrimFunc(symshen_4assoc_1_6), symos, tmp20786, tmp20789)


tmp20791 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20790)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20790
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20791

tmp20792 := MakeNative(func(__e *ControlFlow) {
V6574 := __e.Get(1)
_ = V6574
__e.Return(MakeNative(func(__e *ControlFlow) {
B6570 := __e.Get(1)
_ = B6570
__e.Return(MakeNative(func(__e *ControlFlow) {
L6571 := __e.Get(1)
_ = L6571
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6572 := __e.Get(1)
_ = Key6572
__e.Return(MakeNative(func(__e *ControlFlow) {
C6573 := __e.Get(1)
_ = C6573
tmp20793 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20794 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20793)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20793
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20795 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp20794)
}
__typedArg0 := symsymbol
__typedArg1 := tmp20794
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6574, tmp20795, B6570, L6571, Key6572, C6573)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20796 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20797 := Call(__e, PrimFunc(symshen_4assoc_1_6), sympackage_2, tmp20792, tmp20796)


tmp20798 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20797)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20797
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20798

tmp20799 := MakeNative(func(__e *ControlFlow) {
V6579 := __e.Get(1)
_ = V6579
__e.Return(MakeNative(func(__e *ControlFlow) {
B6575 := __e.Get(1)
_ = B6575
__e.Return(MakeNative(func(__e *ControlFlow) {
L6576 := __e.Get(1)
_ = L6576
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6577 := __e.Get(1)
_ = Key6577
__e.Return(MakeNative(func(__e *ControlFlow) {
C6578 := __e.Get(1)
_ = C6578
tmp20800 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20801 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20800)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20800
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6579, tmp20801, B6575, L6576, Key6577, C6578)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20802 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20803 := Call(__e, PrimFunc(symshen_4assoc_1_6), symport, tmp20799, tmp20802)


tmp20804 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20803)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20803
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20804

tmp20805 := MakeNative(func(__e *ControlFlow) {
V6584 := __e.Get(1)
_ = V6584
__e.Return(MakeNative(func(__e *ControlFlow) {
B6580 := __e.Get(1)
_ = B6580
__e.Return(MakeNative(func(__e *ControlFlow) {
L6581 := __e.Get(1)
_ = L6581
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6582 := __e.Get(1)
_ = Key6582
__e.Return(MakeNative(func(__e *ControlFlow) {
C6583 := __e.Get(1)
_ = C6583
tmp20806 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20807 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20806)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20806
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6584, tmp20807, B6580, L6581, Key6582, C6583)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20808 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20809 := Call(__e, PrimFunc(symshen_4assoc_1_6), symporters, tmp20805, tmp20808)


tmp20810 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20809)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20809
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20810

tmp20811 := MakeNative(func(__e *ControlFlow) {
V6589 := __e.Get(1)
_ = V6589
__e.Return(MakeNative(func(__e *ControlFlow) {
B6585 := __e.Get(1)
_ = B6585
__e.Return(MakeNative(func(__e *ControlFlow) {
L6586 := __e.Get(1)
_ = L6586
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6587 := __e.Get(1)
_ = Key6587
__e.Return(MakeNative(func(__e *ControlFlow) {
C6588 := __e.Get(1)
_ = C6588
tmp20812 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20813 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20812)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20812
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20814 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp20813)
}
__typedArg0 := symnumber
__typedArg1 := tmp20813
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20815 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20814, Nil)
}
__typedArg0 := tmp20814
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20816 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20815)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20815
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20817 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp20816)
}
__typedArg0 := symstring
__typedArg1 := tmp20816
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6589, tmp20817, B6585, L6586, Key6587, C6588)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20818 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20819 := Call(__e, PrimFunc(symshen_4assoc_1_6), sympos, tmp20811, tmp20818)


tmp20820 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20819)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20819
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20820

tmp20821 := MakeNative(func(__e *ControlFlow) {
V6594 := __e.Get(1)
_ = V6594
__e.Return(MakeNative(func(__e *ControlFlow) {
B6590 := __e.Get(1)
_ = B6590
__e.Return(MakeNative(func(__e *ControlFlow) {
L6591 := __e.Get(1)
_ = L6591
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6592 := __e.Get(1)
_ = Key6592
__e.Return(MakeNative(func(__e *ControlFlow) {
C6593 := __e.Get(1)
_ = C6593
tmp20822 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symout, Nil)
}
__typedArg0 := symout
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20823 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp20822)
}
__typedArg0 := symstream
__typedArg1 := tmp20822
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20824 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20825 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20824)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20824
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20826 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20823, tmp20825)
}
__typedArg0 := tmp20823
__typedArg1 := tmp20825
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20827 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20826, Nil)
}
__typedArg0 := tmp20826
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20828 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20827)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20827
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20829 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp20828)
}
__typedArg0 := symstring
__typedArg1 := tmp20828
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6594, tmp20829, B6590, L6591, Key6592, C6593)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20830 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20831 := Call(__e, PrimFunc(symshen_4assoc_1_6), sympr, tmp20821, tmp20830)


tmp20832 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20831)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20831
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20832

tmp20833 := MakeNative(func(__e *ControlFlow) {
V6599 := __e.Get(1)
_ = V6599
__e.Return(MakeNative(func(__e *ControlFlow) {
B6595 := __e.Get(1)
_ = B6595
__e.Return(MakeNative(func(__e *ControlFlow) {
L6596 := __e.Get(1)
_ = L6596
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6597 := __e.Get(1)
_ = Key6597
__e.Return(MakeNative(func(__e *ControlFlow) {
C6598 := __e.Get(1)
_ = C6598
tmp20834 := Call(__e, PrimFunc(symshen_4newpv), B6595)


let__18994 := tmp20834
_ = let__18994

tmp20835 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18994, Nil)
}
__typedArg0 := let__18994
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20836 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20835)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20835
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20837 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18994, tmp20836)
}
__typedArg0 := let__18994
__typedArg1 := tmp20836
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20838 := Call(__e, PrimFunc(symis_b), V6599, tmp20837, B6595, L6596, Key6597, C6598)


__e.TailApply(PrimFunc(symshen_4gc), B6595, tmp20838)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20839 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20840 := Call(__e, PrimFunc(symshen_4assoc_1_6), symprint, tmp20833, tmp20839)


tmp20841 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20840)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20840
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20841

tmp20842 := MakeNative(func(__e *ControlFlow) {
V6604 := __e.Get(1)
_ = V6604
__e.Return(MakeNative(func(__e *ControlFlow) {
B6600 := __e.Get(1)
_ = B6600
__e.Return(MakeNative(func(__e *ControlFlow) {
L6601 := __e.Get(1)
_ = L6601
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6602 := __e.Get(1)
_ = Key6602
__e.Return(MakeNative(func(__e *ControlFlow) {
C6603 := __e.Get(1)
_ = C6603
tmp20843 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20844 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20843)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20843
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20845 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp20844)
}
__typedArg0 := symsymbol
__typedArg1 := tmp20844
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6604, tmp20845, B6600, L6601, Key6602, C6603)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20846 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20847 := Call(__e, PrimFunc(symshen_4assoc_1_6), symprofile, tmp20842, tmp20846)


tmp20848 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20847)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20847
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20848

tmp20849 := MakeNative(func(__e *ControlFlow) {
V6610 := __e.Get(1)
_ = V6610
__e.Return(MakeNative(func(__e *ControlFlow) {
B6606 := __e.Get(1)
_ = B6606
__e.Return(MakeNative(func(__e *ControlFlow) {
L6607 := __e.Get(1)
_ = L6607
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6608 := __e.Get(1)
_ = Key6608
__e.Return(MakeNative(func(__e *ControlFlow) {
C6609 := __e.Get(1)
_ = C6609
tmp20850 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20851 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20850)
}
__typedArg0 := symlist
__typedArg1 := tmp20850
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20852 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20853 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20852)
}
__typedArg0 := symlist
__typedArg1 := tmp20852
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20854 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20853, Nil)
}
__typedArg0 := tmp20853
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20855 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20854)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20854
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20856 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20851, tmp20855)
}
__typedArg0 := tmp20851
__typedArg1 := tmp20855
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6610, tmp20856, B6606, L6607, Key6608, C6609)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20857 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20858 := Call(__e, PrimFunc(symshen_4assoc_1_6), sympreclude, tmp20849, tmp20857)


tmp20859 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20858)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20858
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20859

tmp20860 := MakeNative(func(__e *ControlFlow) {
V6615 := __e.Get(1)
_ = V6615
__e.Return(MakeNative(func(__e *ControlFlow) {
B6611 := __e.Get(1)
_ = B6611
__e.Return(MakeNative(func(__e *ControlFlow) {
L6612 := __e.Get(1)
_ = L6612
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6613 := __e.Get(1)
_ = Key6613
__e.Return(MakeNative(func(__e *ControlFlow) {
C6614 := __e.Get(1)
_ = C6614
tmp20861 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20862 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20861)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20861
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20863 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp20862)
}
__typedArg0 := symstring
__typedArg1 := tmp20862
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6615, tmp20863, B6611, L6612, Key6613, C6614)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20864 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20865 := Call(__e, PrimFunc(symshen_4assoc_1_6), symshen_4proc_1nl, tmp20860, tmp20864)


tmp20866 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20865)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20865
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20866

tmp20867 := MakeNative(func(__e *ControlFlow) {
V6620 := __e.Get(1)
_ = V6620
__e.Return(MakeNative(func(__e *ControlFlow) {
B6616 := __e.Get(1)
_ = B6616
__e.Return(MakeNative(func(__e *ControlFlow) {
L6617 := __e.Get(1)
_ = L6617
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6618 := __e.Get(1)
_ = Key6618
__e.Return(MakeNative(func(__e *ControlFlow) {
C6619 := __e.Get(1)
_ = C6619
tmp20868 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20869 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_d, tmp20868)
}
__typedArg0 := sym_d
__typedArg1 := tmp20868
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20870 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp20869)
}
__typedArg0 := symsymbol
__typedArg1 := tmp20869
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20871 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20870, Nil)
}
__typedArg0 := tmp20870
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20872 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20871)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20871
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20873 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp20872)
}
__typedArg0 := symsymbol
__typedArg1 := tmp20872
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6620, tmp20873, B6616, L6617, Key6618, C6619)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20874 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20875 := Call(__e, PrimFunc(symshen_4assoc_1_6), symprofile_1results, tmp20867, tmp20874)


tmp20876 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20875)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20875
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20876

tmp20877 := MakeNative(func(__e *ControlFlow) {
V6625 := __e.Get(1)
_ = V6625
__e.Return(MakeNative(func(__e *ControlFlow) {
B6621 := __e.Get(1)
_ = B6621
__e.Return(MakeNative(func(__e *ControlFlow) {
L6622 := __e.Get(1)
_ = L6622
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6623 := __e.Get(1)
_ = Key6623
__e.Return(MakeNative(func(__e *ControlFlow) {
C6624 := __e.Get(1)
_ = C6624
tmp20878 := Call(__e, PrimFunc(symshen_4newpv), B6621)


let__18995 := tmp20878
_ = let__18995

tmp20879 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18995, Nil)
}
__typedArg0 := let__18995
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20880 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20879)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20879
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20881 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18995, tmp20880)
}
__typedArg0 := let__18995
__typedArg1 := tmp20880
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20882 := Call(__e, PrimFunc(symis_b), V6625, tmp20881, B6621, L6622, Key6623, C6624)


__e.TailApply(PrimFunc(symshen_4gc), B6621, tmp20882)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20883 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20884 := Call(__e, PrimFunc(symshen_4assoc_1_6), symprotect, tmp20877, tmp20883)


tmp20885 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20884)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20884
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20885

tmp20886 := MakeNative(func(__e *ControlFlow) {
V6631 := __e.Get(1)
_ = V6631
__e.Return(MakeNative(func(__e *ControlFlow) {
B6627 := __e.Get(1)
_ = B6627
__e.Return(MakeNative(func(__e *ControlFlow) {
L6628 := __e.Get(1)
_ = L6628
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6629 := __e.Get(1)
_ = Key6629
__e.Return(MakeNative(func(__e *ControlFlow) {
C6630 := __e.Get(1)
_ = C6630
tmp20887 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20888 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20887)
}
__typedArg0 := symlist
__typedArg1 := tmp20887
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20889 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20890 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20889)
}
__typedArg0 := symlist
__typedArg1 := tmp20889
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20891 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20890, Nil)
}
__typedArg0 := tmp20890
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20892 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20891)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20891
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20893 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20888, tmp20892)
}
__typedArg0 := tmp20888
__typedArg1 := tmp20892
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6631, tmp20893, B6627, L6628, Key6629, C6630)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20894 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20895 := Call(__e, PrimFunc(symshen_4assoc_1_6), sympreclude_1all_1but, tmp20886, tmp20894)


tmp20896 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20895)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20895
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20896

tmp20897 := MakeNative(func(__e *ControlFlow) {
V6636 := __e.Get(1)
_ = V6636
__e.Return(MakeNative(func(__e *ControlFlow) {
B6632 := __e.Get(1)
_ = B6632
__e.Return(MakeNative(func(__e *ControlFlow) {
L6633 := __e.Get(1)
_ = L6633
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6634 := __e.Get(1)
_ = Key6634
__e.Return(MakeNative(func(__e *ControlFlow) {
C6635 := __e.Get(1)
_ = C6635
tmp20898 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symout, Nil)
}
__typedArg0 := symout
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20899 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp20898)
}
__typedArg0 := symstream
__typedArg1 := tmp20898
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20900 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20901 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20900)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20900
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20902 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20899, tmp20901)
}
__typedArg0 := tmp20899
__typedArg1 := tmp20901
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20903 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20902, Nil)
}
__typedArg0 := tmp20902
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20904 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20903)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20903
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20905 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp20904)
}
__typedArg0 := symstring
__typedArg1 := tmp20904
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6636, tmp20905, B6632, L6633, Key6634, C6635)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20906 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20907 := Call(__e, PrimFunc(symshen_4assoc_1_6), symshen_4prhush, tmp20897, tmp20906)


tmp20908 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20907)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20907
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20908

tmp20909 := MakeNative(func(__e *ControlFlow) {
V6641 := __e.Get(1)
_ = V6641
__e.Return(MakeNative(func(__e *ControlFlow) {
B6637 := __e.Get(1)
_ = B6637
__e.Return(MakeNative(func(__e *ControlFlow) {
L6638 := __e.Get(1)
_ = L6638
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6639 := __e.Get(1)
_ = Key6639
__e.Return(MakeNative(func(__e *ControlFlow) {
C6640 := __e.Get(1)
_ = C6640
tmp20910 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20911 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20910)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20910
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20912 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp20911)
}
__typedArg0 := symnumber
__typedArg1 := tmp20911
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6641, tmp20912, B6637, L6638, Key6639, C6640)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20913 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20914 := Call(__e, PrimFunc(symshen_4assoc_1_6), symprolog_1memory, tmp20909, tmp20913)


tmp20915 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20914)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20914
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20915

tmp20916 := MakeNative(func(__e *ControlFlow) {
V6646 := __e.Get(1)
_ = V6646
__e.Return(MakeNative(func(__e *ControlFlow) {
B6642 := __e.Get(1)
_ = B6642
__e.Return(MakeNative(func(__e *ControlFlow) {
L6643 := __e.Get(1)
_ = L6643
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6644 := __e.Get(1)
_ = Key6644
__e.Return(MakeNative(func(__e *ControlFlow) {
C6645 := __e.Get(1)
_ = C6645
tmp20917 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunit, Nil)
}
__typedArg0 := symunit
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20918 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20917)
}
__typedArg0 := symlist
__typedArg1 := tmp20917
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20919 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20918, Nil)
}
__typedArg0 := tmp20918
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20920 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20919)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20919
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20921 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp20920)
}
__typedArg0 := symsymbol
__typedArg1 := tmp20920
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6646, tmp20921, B6642, L6643, Key6644, C6645)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20922 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20923 := Call(__e, PrimFunc(symshen_4assoc_1_6), symps, tmp20916, tmp20922)


tmp20924 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20923)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20923
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20924

tmp20925 := MakeNative(func(__e *ControlFlow) {
V6651 := __e.Get(1)
_ = V6651
__e.Return(MakeNative(func(__e *ControlFlow) {
B6647 := __e.Get(1)
_ = B6647
__e.Return(MakeNative(func(__e *ControlFlow) {
L6648 := __e.Get(1)
_ = L6648
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6649 := __e.Get(1)
_ = Key6649
__e.Return(MakeNative(func(__e *ControlFlow) {
C6650 := __e.Get(1)
_ = C6650
tmp20926 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symin, Nil)
}
__typedArg0 := symin
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20927 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp20926)
}
__typedArg0 := symstream
__typedArg1 := tmp20926
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20928 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunit, Nil)
}
__typedArg0 := symunit
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20929 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20928)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20928
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20930 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20927, tmp20929)
}
__typedArg0 := tmp20927
__typedArg1 := tmp20929
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6651, tmp20930, B6647, L6648, Key6649, C6650)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20931 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20932 := Call(__e, PrimFunc(symshen_4assoc_1_6), symread, tmp20925, tmp20931)


tmp20933 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20932)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20932
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20933

tmp20934 := MakeNative(func(__e *ControlFlow) {
V6656 := __e.Get(1)
_ = V6656
__e.Return(MakeNative(func(__e *ControlFlow) {
B6652 := __e.Get(1)
_ = B6652
__e.Return(MakeNative(func(__e *ControlFlow) {
L6653 := __e.Get(1)
_ = L6653
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6654 := __e.Get(1)
_ = Key6654
__e.Return(MakeNative(func(__e *ControlFlow) {
C6655 := __e.Get(1)
_ = C6655
tmp20935 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symin, Nil)
}
__typedArg0 := symin
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20936 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp20935)
}
__typedArg0 := symstream
__typedArg1 := tmp20935
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20937 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20938 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20937)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20937
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20939 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20936, tmp20938)
}
__typedArg0 := tmp20936
__typedArg1 := tmp20938
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6656, tmp20939, B6652, L6653, Key6654, C6655)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20940 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20941 := Call(__e, PrimFunc(symshen_4assoc_1_6), symread_1unit_1string, tmp20934, tmp20940)


tmp20942 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20941)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20941
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20942

tmp20943 := MakeNative(func(__e *ControlFlow) {
V6661 := __e.Get(1)
_ = V6661
__e.Return(MakeNative(func(__e *ControlFlow) {
B6657 := __e.Get(1)
_ = B6657
__e.Return(MakeNative(func(__e *ControlFlow) {
L6658 := __e.Get(1)
_ = L6658
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6659 := __e.Get(1)
_ = Key6659
__e.Return(MakeNative(func(__e *ControlFlow) {
C6660 := __e.Get(1)
_ = C6660
tmp20944 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symin, Nil)
}
__typedArg0 := symin
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20945 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp20944)
}
__typedArg0 := symstream
__typedArg1 := tmp20944
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20946 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20947 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20946)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20946
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20948 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20945, tmp20947)
}
__typedArg0 := tmp20945
__typedArg1 := tmp20947
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6661, tmp20948, B6657, L6658, Key6659, C6660)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20949 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20950 := Call(__e, PrimFunc(symshen_4assoc_1_6), symread_1byte, tmp20943, tmp20949)


tmp20951 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20950)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20950
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20951

tmp20952 := MakeNative(func(__e *ControlFlow) {
V6666 := __e.Get(1)
_ = V6666
__e.Return(MakeNative(func(__e *ControlFlow) {
B6662 := __e.Get(1)
_ = B6662
__e.Return(MakeNative(func(__e *ControlFlow) {
L6663 := __e.Get(1)
_ = L6663
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6664 := __e.Get(1)
_ = Key6664
__e.Return(MakeNative(func(__e *ControlFlow) {
C6665 := __e.Get(1)
_ = C6665
tmp20953 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20954 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20953)
}
__typedArg0 := symlist
__typedArg1 := tmp20953
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20955 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20954, Nil)
}
__typedArg0 := tmp20954
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20956 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20955)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20955
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20957 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp20956)
}
__typedArg0 := symstring
__typedArg1 := tmp20956
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6666, tmp20957, B6662, L6663, Key6664, C6665)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20958 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20959 := Call(__e, PrimFunc(symshen_4assoc_1_6), symread_1file_1as_1bytelist, tmp20952, tmp20958)


tmp20960 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20959)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20959
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20960

tmp20961 := MakeNative(func(__e *ControlFlow) {
V6671 := __e.Get(1)
_ = V6671
__e.Return(MakeNative(func(__e *ControlFlow) {
B6667 := __e.Get(1)
_ = B6667
__e.Return(MakeNative(func(__e *ControlFlow) {
L6668 := __e.Get(1)
_ = L6668
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6669 := __e.Get(1)
_ = Key6669
__e.Return(MakeNative(func(__e *ControlFlow) {
C6670 := __e.Get(1)
_ = C6670
tmp20962 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20963 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20962)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20962
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20964 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp20963)
}
__typedArg0 := symstring
__typedArg1 := tmp20963
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6671, tmp20964, B6667, L6668, Key6669, C6670)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20965 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20966 := Call(__e, PrimFunc(symshen_4assoc_1_6), symread_1file_1as_1string, tmp20961, tmp20965)


tmp20967 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20966)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20966
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20967

tmp20968 := MakeNative(func(__e *ControlFlow) {
V6676 := __e.Get(1)
_ = V6676
__e.Return(MakeNative(func(__e *ControlFlow) {
B6672 := __e.Get(1)
_ = B6672
__e.Return(MakeNative(func(__e *ControlFlow) {
L6673 := __e.Get(1)
_ = L6673
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6674 := __e.Get(1)
_ = Key6674
__e.Return(MakeNative(func(__e *ControlFlow) {
C6675 := __e.Get(1)
_ = C6675
tmp20969 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunit, Nil)
}
__typedArg0 := symunit
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20970 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20969)
}
__typedArg0 := symlist
__typedArg1 := tmp20969
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20971 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20970, Nil)
}
__typedArg0 := tmp20970
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20972 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20971)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20971
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20973 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp20972)
}
__typedArg0 := symstring
__typedArg1 := tmp20972
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6676, tmp20973, B6672, L6673, Key6674, C6675)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20974 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20975 := Call(__e, PrimFunc(symshen_4assoc_1_6), symread_1file, tmp20968, tmp20974)


tmp20976 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20975)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20975
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20976

tmp20977 := MakeNative(func(__e *ControlFlow) {
V6681 := __e.Get(1)
_ = V6681
__e.Return(MakeNative(func(__e *ControlFlow) {
B6677 := __e.Get(1)
_ = B6677
__e.Return(MakeNative(func(__e *ControlFlow) {
L6678 := __e.Get(1)
_ = L6678
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6679 := __e.Get(1)
_ = Key6679
__e.Return(MakeNative(func(__e *ControlFlow) {
C6680 := __e.Get(1)
_ = C6680
tmp20978 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunit, Nil)
}
__typedArg0 := symunit
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20979 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20978)
}
__typedArg0 := symlist
__typedArg1 := tmp20978
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20980 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20979, Nil)
}
__typedArg0 := tmp20979
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20981 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20980)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20980
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20982 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp20981)
}
__typedArg0 := symstring
__typedArg1 := tmp20981
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6681, tmp20982, B6677, L6678, Key6679, C6680)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20983 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20984 := Call(__e, PrimFunc(symshen_4assoc_1_6), symread_1from_1string, tmp20977, tmp20983)


tmp20985 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20984)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20984
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20985

tmp20986 := MakeNative(func(__e *ControlFlow) {
V6686 := __e.Get(1)
_ = V6686
__e.Return(MakeNative(func(__e *ControlFlow) {
B6682 := __e.Get(1)
_ = B6682
__e.Return(MakeNative(func(__e *ControlFlow) {
L6683 := __e.Get(1)
_ = L6683
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6684 := __e.Get(1)
_ = Key6684
__e.Return(MakeNative(func(__e *ControlFlow) {
C6685 := __e.Get(1)
_ = C6685
tmp20987 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunit, Nil)
}
__typedArg0 := symunit
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20988 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20987)
}
__typedArg0 := symlist
__typedArg1 := tmp20987
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20989 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20988, Nil)
}
__typedArg0 := tmp20988
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20990 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20989)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20989
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20991 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp20990)
}
__typedArg0 := symstring
__typedArg1 := tmp20990
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6686, tmp20991, B6682, L6683, Key6684, C6685)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20992 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20993 := Call(__e, PrimFunc(symshen_4assoc_1_6), symread_1from_1string_1unprocessed, tmp20986, tmp20992)


tmp20994 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20993)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20993
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20994

tmp20995 := MakeNative(func(__e *ControlFlow) {
V6691 := __e.Get(1)
_ = V6691
__e.Return(MakeNative(func(__e *ControlFlow) {
B6687 := __e.Get(1)
_ = B6687
__e.Return(MakeNative(func(__e *ControlFlow) {
L6688 := __e.Get(1)
_ = L6688
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6689 := __e.Get(1)
_ = Key6689
__e.Return(MakeNative(func(__e *ControlFlow) {
C6690 := __e.Get(1)
_ = C6690
tmp20996 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20997 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20996)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20996
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6691, tmp20997, B6687, L6688, Key6689, C6690)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20998 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20999 := Call(__e, PrimFunc(symshen_4assoc_1_6), symrelease, tmp20995, tmp20998)


tmp21000 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20999)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20999
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21000

tmp21001 := MakeNative(func(__e *ControlFlow) {
V6697 := __e.Get(1)
_ = V6697
__e.Return(MakeNative(func(__e *ControlFlow) {
B6693 := __e.Get(1)
_ = B6693
__e.Return(MakeNative(func(__e *ControlFlow) {
L6694 := __e.Get(1)
_ = L6694
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6695 := __e.Get(1)
_ = Key6695
__e.Return(MakeNative(func(__e *ControlFlow) {
C6696 := __e.Get(1)
_ = C6696
tmp21002 := Call(__e, PrimFunc(symshen_4newpv), B6693)


let__18996 := tmp21002
_ = let__18996

tmp21003 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18996, Nil)
}
__typedArg0 := let__18996
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21004 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp21003)
}
__typedArg0 := symlist
__typedArg1 := tmp21003
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21005 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18996, Nil)
}
__typedArg0 := let__18996
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21006 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp21005)
}
__typedArg0 := symlist
__typedArg1 := tmp21005
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21007 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21006, Nil)
}
__typedArg0 := tmp21006
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21008 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21007)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21007
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21009 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21004, tmp21008)
}
__typedArg0 := tmp21004
__typedArg1 := tmp21008
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21010 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21009, Nil)
}
__typedArg0 := tmp21009
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21011 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21010)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21010
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21012 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18996, tmp21011)
}
__typedArg0 := let__18996
__typedArg1 := tmp21011
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21013 := Call(__e, PrimFunc(symis_b), V6697, tmp21012, B6693, L6694, Key6695, C6696)


__e.TailApply(PrimFunc(symshen_4gc), B6693, tmp21013)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21014 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21015 := Call(__e, PrimFunc(symshen_4assoc_1_6), symremove, tmp21001, tmp21014)


tmp21016 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21015)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21015
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21016

tmp21017 := MakeNative(func(__e *ControlFlow) {
V6703 := __e.Get(1)
_ = V6703
__e.Return(MakeNative(func(__e *ControlFlow) {
B6699 := __e.Get(1)
_ = B6699
__e.Return(MakeNative(func(__e *ControlFlow) {
L6700 := __e.Get(1)
_ = L6700
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6701 := __e.Get(1)
_ = Key6701
__e.Return(MakeNative(func(__e *ControlFlow) {
C6702 := __e.Get(1)
_ = C6702
tmp21018 := Call(__e, PrimFunc(symshen_4newpv), B6699)


let__18997 := tmp21018
_ = let__18997

tmp21019 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18997, Nil)
}
__typedArg0 := let__18997
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21020 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp21019)
}
__typedArg0 := symlist
__typedArg1 := tmp21019
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21021 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18997, Nil)
}
__typedArg0 := let__18997
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21022 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp21021)
}
__typedArg0 := symlist
__typedArg1 := tmp21021
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21023 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21022, Nil)
}
__typedArg0 := tmp21022
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21024 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21023)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21023
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21025 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21020, tmp21024)
}
__typedArg0 := tmp21020
__typedArg1 := tmp21024
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21026 := Call(__e, PrimFunc(symis_b), V6703, tmp21025, B6699, L6700, Key6701, C6702)


__e.TailApply(PrimFunc(symshen_4gc), B6699, tmp21026)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21027 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21028 := Call(__e, PrimFunc(symshen_4assoc_1_6), symreverse, tmp21017, tmp21027)


tmp21029 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21028)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21028
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21029

tmp21030 := MakeNative(func(__e *ControlFlow) {
V6708 := __e.Get(1)
_ = V6708
__e.Return(MakeNative(func(__e *ControlFlow) {
B6704 := __e.Get(1)
_ = B6704
__e.Return(MakeNative(func(__e *ControlFlow) {
L6705 := __e.Get(1)
_ = L6705
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6706 := __e.Get(1)
_ = Key6706
__e.Return(MakeNative(func(__e *ControlFlow) {
C6707 := __e.Get(1)
_ = C6707
tmp21031 := Call(__e, PrimFunc(symshen_4newpv), B6704)


let__18998 := tmp21031
_ = let__18998

tmp21032 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18998, Nil)
}
__typedArg0 := let__18998
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21033 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21032)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21032
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21034 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp21033)
}
__typedArg0 := symstring
__typedArg1 := tmp21033
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21035 := Call(__e, PrimFunc(symis_b), V6708, tmp21034, B6704, L6705, Key6706, C6707)


__e.TailApply(PrimFunc(symshen_4gc), B6704, tmp21035)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21036 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21037 := Call(__e, PrimFunc(symshen_4assoc_1_6), symsimple_1error, tmp21030, tmp21036)


tmp21038 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21037)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21037
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21038

tmp21039 := MakeNative(func(__e *ControlFlow) {
V6713 := __e.Get(1)
_ = V6713
__e.Return(MakeNative(func(__e *ControlFlow) {
B6709 := __e.Get(1)
_ = B6709
__e.Return(MakeNative(func(__e *ControlFlow) {
L6710 := __e.Get(1)
_ = L6710
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6711 := __e.Get(1)
_ = Key6711
__e.Return(MakeNative(func(__e *ControlFlow) {
C6712 := __e.Get(1)
_ = C6712
tmp21040 := Call(__e, PrimFunc(symshen_4newpv), B6709)


let__18999 := tmp21040
_ = let__18999

tmp21041 := Call(__e, PrimFunc(symshen_4newpv), B6709)


let__19000 := tmp21041
_ = let__19000

tmp21042 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__19000, Nil)
}
__typedArg0 := let__19000
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21043 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_d, tmp21042)
}
__typedArg0 := sym_d
__typedArg1 := tmp21042
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21044 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18999, tmp21043)
}
__typedArg0 := let__18999
__typedArg1 := tmp21043
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21045 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__19000, Nil)
}
__typedArg0 := let__19000
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21046 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21045)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21045
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21047 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21044, tmp21046)
}
__typedArg0 := tmp21044
__typedArg1 := tmp21046
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21048 := Call(__e, PrimFunc(symis_b), V6713, tmp21047, B6709, L6710, Key6711, C6712)


tmp21049 := Call(__e, PrimFunc(symshen_4gc), B6709, tmp21048)


__e.TailApply(PrimFunc(symshen_4gc), B6709, tmp21049)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21050 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21051 := Call(__e, PrimFunc(symshen_4assoc_1_6), symsnd, tmp21039, tmp21050)


tmp21052 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21051)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21051
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21052

tmp21053 := MakeNative(func(__e *ControlFlow) {
V6718 := __e.Get(1)
_ = V6718
__e.Return(MakeNative(func(__e *ControlFlow) {
B6714 := __e.Get(1)
_ = B6714
__e.Return(MakeNative(func(__e *ControlFlow) {
L6715 := __e.Get(1)
_ = L6715
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6716 := __e.Get(1)
_ = Key6716
__e.Return(MakeNative(func(__e *ControlFlow) {
C6717 := __e.Get(1)
_ = C6717
tmp21054 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21055 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21054)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21054
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21056 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp21055)
}
__typedArg0 := symnumber
__typedArg1 := tmp21055
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21057 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21056, Nil)
}
__typedArg0 := tmp21056
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21058 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21057)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21057
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21059 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp21058)
}
__typedArg0 := symsymbol
__typedArg1 := tmp21058
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6718, tmp21059, B6714, L6715, Key6716, C6717)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21060 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21061 := Call(__e, PrimFunc(symshen_4assoc_1_6), symspecialise, tmp21053, tmp21060)


tmp21062 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21061)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21061
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21062

tmp21063 := MakeNative(func(__e *ControlFlow) {
V6723 := __e.Get(1)
_ = V6723
__e.Return(MakeNative(func(__e *ControlFlow) {
B6719 := __e.Get(1)
_ = B6719
__e.Return(MakeNative(func(__e *ControlFlow) {
L6720 := __e.Get(1)
_ = L6720
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6721 := __e.Get(1)
_ = Key6721
__e.Return(MakeNative(func(__e *ControlFlow) {
C6722 := __e.Get(1)
_ = C6722
tmp21064 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21065 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21064)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21064
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21066 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp21065)
}
__typedArg0 := symsymbol
__typedArg1 := tmp21065
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6723, tmp21066, B6719, L6720, Key6721, C6722)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21067 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21068 := Call(__e, PrimFunc(symshen_4assoc_1_6), symspy, tmp21063, tmp21067)


tmp21069 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21068)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21068
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21069

tmp21070 := MakeNative(func(__e *ControlFlow) {
V6728 := __e.Get(1)
_ = V6728
__e.Return(MakeNative(func(__e *ControlFlow) {
B6724 := __e.Get(1)
_ = B6724
__e.Return(MakeNative(func(__e *ControlFlow) {
L6725 := __e.Get(1)
_ = L6725
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6726 := __e.Get(1)
_ = Key6726
__e.Return(MakeNative(func(__e *ControlFlow) {
C6727 := __e.Get(1)
_ = C6727
tmp21071 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21072 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21071)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21071
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6728, tmp21072, B6724, L6725, Key6726, C6727)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21073 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21074 := Call(__e, PrimFunc(symshen_4assoc_1_6), symspy_2, tmp21070, tmp21073)


tmp21075 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21074)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21074
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21075

tmp21076 := MakeNative(func(__e *ControlFlow) {
V6733 := __e.Get(1)
_ = V6733
__e.Return(MakeNative(func(__e *ControlFlow) {
B6729 := __e.Get(1)
_ = B6729
__e.Return(MakeNative(func(__e *ControlFlow) {
L6730 := __e.Get(1)
_ = L6730
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6731 := __e.Get(1)
_ = Key6731
__e.Return(MakeNative(func(__e *ControlFlow) {
C6732 := __e.Get(1)
_ = C6732
tmp21077 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21078 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21077)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21077
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21079 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp21078)
}
__typedArg0 := symsymbol
__typedArg1 := tmp21078
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6733, tmp21079, B6729, L6730, Key6731, C6732)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21080 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21081 := Call(__e, PrimFunc(symshen_4assoc_1_6), symstep, tmp21076, tmp21080)


tmp21082 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21081)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21081
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21082

tmp21083 := MakeNative(func(__e *ControlFlow) {
V6738 := __e.Get(1)
_ = V6738
__e.Return(MakeNative(func(__e *ControlFlow) {
B6734 := __e.Get(1)
_ = B6734
__e.Return(MakeNative(func(__e *ControlFlow) {
L6735 := __e.Get(1)
_ = L6735
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6736 := __e.Get(1)
_ = Key6736
__e.Return(MakeNative(func(__e *ControlFlow) {
C6737 := __e.Get(1)
_ = C6737
tmp21084 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21085 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21084)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21084
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6738, tmp21085, B6734, L6735, Key6736, C6737)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21086 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21087 := Call(__e, PrimFunc(symshen_4assoc_1_6), symstep_2, tmp21083, tmp21086)


tmp21088 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21087)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21087
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21088

tmp21089 := MakeNative(func(__e *ControlFlow) {
V6743 := __e.Get(1)
_ = V6743
__e.Return(MakeNative(func(__e *ControlFlow) {
B6739 := __e.Get(1)
_ = B6739
__e.Return(MakeNative(func(__e *ControlFlow) {
L6740 := __e.Get(1)
_ = L6740
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6741 := __e.Get(1)
_ = Key6741
__e.Return(MakeNative(func(__e *ControlFlow) {
C6742 := __e.Get(1)
_ = C6742
tmp21090 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symin, Nil)
}
__typedArg0 := symin
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21091 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp21090)
}
__typedArg0 := symstream
__typedArg1 := tmp21090
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21092 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21091, Nil)
}
__typedArg0 := tmp21091
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21093 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21092)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21092
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6743, tmp21093, B6739, L6740, Key6741, C6742)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21094 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21095 := Call(__e, PrimFunc(symshen_4assoc_1_6), symstinput, tmp21089, tmp21094)


tmp21096 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21095)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21095
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21096

tmp21097 := MakeNative(func(__e *ControlFlow) {
V6748 := __e.Get(1)
_ = V6748
__e.Return(MakeNative(func(__e *ControlFlow) {
B6744 := __e.Get(1)
_ = B6744
__e.Return(MakeNative(func(__e *ControlFlow) {
L6745 := __e.Get(1)
_ = L6745
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6746 := __e.Get(1)
_ = Key6746
__e.Return(MakeNative(func(__e *ControlFlow) {
C6747 := __e.Get(1)
_ = C6747
tmp21098 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symout, Nil)
}
__typedArg0 := symout
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21099 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp21098)
}
__typedArg0 := symstream
__typedArg1 := tmp21098
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21100 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21099, Nil)
}
__typedArg0 := tmp21099
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21101 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21100)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21100
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6748, tmp21101, B6744, L6745, Key6746, C6747)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21102 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21103 := Call(__e, PrimFunc(symshen_4assoc_1_6), symsterror, tmp21097, tmp21102)


tmp21104 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21103)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21103
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21104

tmp21105 := MakeNative(func(__e *ControlFlow) {
V6753 := __e.Get(1)
_ = V6753
__e.Return(MakeNative(func(__e *ControlFlow) {
B6749 := __e.Get(1)
_ = B6749
__e.Return(MakeNative(func(__e *ControlFlow) {
L6750 := __e.Get(1)
_ = L6750
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6751 := __e.Get(1)
_ = Key6751
__e.Return(MakeNative(func(__e *ControlFlow) {
C6752 := __e.Get(1)
_ = C6752
tmp21106 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symout, Nil)
}
__typedArg0 := symout
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21107 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp21106)
}
__typedArg0 := symstream
__typedArg1 := tmp21106
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21108 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21107, Nil)
}
__typedArg0 := tmp21107
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21109 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21108)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21108
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6753, tmp21109, B6749, L6750, Key6751, C6752)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21110 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21111 := Call(__e, PrimFunc(symshen_4assoc_1_6), symstoutput, tmp21105, tmp21110)


tmp21112 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21111)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21111
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21112

tmp21113 := MakeNative(func(__e *ControlFlow) {
V6758 := __e.Get(1)
_ = V6758
__e.Return(MakeNative(func(__e *ControlFlow) {
B6754 := __e.Get(1)
_ = B6754
__e.Return(MakeNative(func(__e *ControlFlow) {
L6755 := __e.Get(1)
_ = L6755
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6756 := __e.Get(1)
_ = Key6756
__e.Return(MakeNative(func(__e *ControlFlow) {
C6757 := __e.Get(1)
_ = C6757
tmp21114 := Call(__e, PrimFunc(symshen_4newpv), B6754)


let__19001 := tmp21114
_ = let__19001

tmp21115 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21116 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21115)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21115
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21117 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__19001, tmp21116)
}
__typedArg0 := let__19001
__typedArg1 := tmp21116
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21118 := Call(__e, PrimFunc(symis_b), V6758, tmp21117, B6754, L6755, Key6756, C6757)


__e.TailApply(PrimFunc(symshen_4gc), B6754, tmp21118)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21119 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21120 := Call(__e, PrimFunc(symshen_4assoc_1_6), symstring_2, tmp21113, tmp21119)


tmp21121 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21120)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21120
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21121

tmp21122 := MakeNative(func(__e *ControlFlow) {
V6763 := __e.Get(1)
_ = V6763
__e.Return(MakeNative(func(__e *ControlFlow) {
B6759 := __e.Get(1)
_ = B6759
__e.Return(MakeNative(func(__e *ControlFlow) {
L6760 := __e.Get(1)
_ = L6760
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6761 := __e.Get(1)
_ = Key6761
__e.Return(MakeNative(func(__e *ControlFlow) {
C6762 := __e.Get(1)
_ = C6762
tmp21123 := Call(__e, PrimFunc(symshen_4newpv), B6759)


let__19002 := tmp21123
_ = let__19002

tmp21124 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21125 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21124)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21124
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21126 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__19002, tmp21125)
}
__typedArg0 := let__19002
__typedArg1 := tmp21125
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21127 := Call(__e, PrimFunc(symis_b), V6763, tmp21126, B6759, L6760, Key6761, C6762)


__e.TailApply(PrimFunc(symshen_4gc), B6759, tmp21127)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21128 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21129 := Call(__e, PrimFunc(symshen_4assoc_1_6), symstr, tmp21122, tmp21128)


tmp21130 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21129)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21129
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21130

tmp21131 := MakeNative(func(__e *ControlFlow) {
V6768 := __e.Get(1)
_ = V6768
__e.Return(MakeNative(func(__e *ControlFlow) {
B6764 := __e.Get(1)
_ = B6764
__e.Return(MakeNative(func(__e *ControlFlow) {
L6765 := __e.Get(1)
_ = L6765
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6766 := __e.Get(1)
_ = Key6766
__e.Return(MakeNative(func(__e *ControlFlow) {
C6767 := __e.Get(1)
_ = C6767
tmp21132 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21133 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21132)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21132
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21134 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp21133)
}
__typedArg0 := symstring
__typedArg1 := tmp21133
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6768, tmp21134, B6764, L6765, Key6766, C6767)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21135 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21136 := Call(__e, PrimFunc(symshen_4assoc_1_6), symstring_1_6n, tmp21131, tmp21135)


tmp21137 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21136)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21136
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21137

tmp21138 := MakeNative(func(__e *ControlFlow) {
V6773 := __e.Get(1)
_ = V6773
__e.Return(MakeNative(func(__e *ControlFlow) {
B6769 := __e.Get(1)
_ = B6769
__e.Return(MakeNative(func(__e *ControlFlow) {
L6770 := __e.Get(1)
_ = L6770
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6771 := __e.Get(1)
_ = Key6771
__e.Return(MakeNative(func(__e *ControlFlow) {
C6772 := __e.Get(1)
_ = C6772
tmp21139 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21140 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21139)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21139
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21141 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp21140)
}
__typedArg0 := symstring
__typedArg1 := tmp21140
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6773, tmp21141, B6769, L6770, Key6771, C6772)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21142 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21143 := Call(__e, PrimFunc(symshen_4assoc_1_6), symstring_1_6symbol, tmp21138, tmp21142)


tmp21144 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21143)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21143
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21144

tmp21145 := MakeNative(func(__e *ControlFlow) {
V6778 := __e.Get(1)
_ = V6778
__e.Return(MakeNative(func(__e *ControlFlow) {
B6774 := __e.Get(1)
_ = B6774
__e.Return(MakeNative(func(__e *ControlFlow) {
L6775 := __e.Get(1)
_ = L6775
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6776 := __e.Get(1)
_ = Key6776
__e.Return(MakeNative(func(__e *ControlFlow) {
C6777 := __e.Get(1)
_ = C6777
tmp21146 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21147 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp21146)
}
__typedArg0 := symlist
__typedArg1 := tmp21146
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21148 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21149 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21148)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21148
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21150 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21147, tmp21149)
}
__typedArg0 := tmp21147
__typedArg1 := tmp21149
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6778, tmp21150, B6774, L6775, Key6776, C6777)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21151 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21152 := Call(__e, PrimFunc(symshen_4assoc_1_6), symsum, tmp21145, tmp21151)


tmp21153 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21152)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21152
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21153

tmp21154 := MakeNative(func(__e *ControlFlow) {
V6783 := __e.Get(1)
_ = V6783
__e.Return(MakeNative(func(__e *ControlFlow) {
B6779 := __e.Get(1)
_ = B6779
__e.Return(MakeNative(func(__e *ControlFlow) {
L6780 := __e.Get(1)
_ = L6780
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6781 := __e.Get(1)
_ = Key6781
__e.Return(MakeNative(func(__e *ControlFlow) {
C6782 := __e.Get(1)
_ = C6782
tmp21155 := Call(__e, PrimFunc(symshen_4newpv), B6779)


let__19003 := tmp21155
_ = let__19003

tmp21156 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21157 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21156)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21156
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21158 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__19003, tmp21157)
}
__typedArg0 := let__19003
__typedArg1 := tmp21157
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21159 := Call(__e, PrimFunc(symis_b), V6783, tmp21158, B6779, L6780, Key6781, C6782)


__e.TailApply(PrimFunc(symshen_4gc), B6779, tmp21159)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21160 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21161 := Call(__e, PrimFunc(symshen_4assoc_1_6), symsymbol_2, tmp21154, tmp21160)


tmp21162 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21161)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21161
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21162

tmp21163 := MakeNative(func(__e *ControlFlow) {
V6788 := __e.Get(1)
_ = V6788
__e.Return(MakeNative(func(__e *ControlFlow) {
B6784 := __e.Get(1)
_ = B6784
__e.Return(MakeNative(func(__e *ControlFlow) {
L6785 := __e.Get(1)
_ = L6785
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6786 := __e.Get(1)
_ = Key6786
__e.Return(MakeNative(func(__e *ControlFlow) {
C6787 := __e.Get(1)
_ = C6787
tmp21164 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21165 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21164)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21164
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21166 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp21165)
}
__typedArg0 := symsymbol
__typedArg1 := tmp21165
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6788, tmp21166, B6784, L6785, Key6786, C6787)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21167 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21168 := Call(__e, PrimFunc(symshen_4assoc_1_6), symsystemf, tmp21163, tmp21167)


tmp21169 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21168)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21168
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21169

tmp21170 := MakeNative(func(__e *ControlFlow) {
V6793 := __e.Get(1)
_ = V6793
__e.Return(MakeNative(func(__e *ControlFlow) {
B6789 := __e.Get(1)
_ = B6789
__e.Return(MakeNative(func(__e *ControlFlow) {
L6790 := __e.Get(1)
_ = L6790
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6791 := __e.Get(1)
_ = Key6791
__e.Return(MakeNative(func(__e *ControlFlow) {
C6792 := __e.Get(1)
_ = C6792
tmp21171 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21172 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21171)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21171
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6793, tmp21172, B6789, L6790, Key6791, C6792)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21173 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21174 := Call(__e, PrimFunc(symshen_4assoc_1_6), symsystem_1S_2, tmp21170, tmp21173)


tmp21175 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21174)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21174
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21175

tmp21176 := MakeNative(func(__e *ControlFlow) {
V6799 := __e.Get(1)
_ = V6799
__e.Return(MakeNative(func(__e *ControlFlow) {
B6795 := __e.Get(1)
_ = B6795
__e.Return(MakeNative(func(__e *ControlFlow) {
L6796 := __e.Get(1)
_ = L6796
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6797 := __e.Get(1)
_ = Key6797
__e.Return(MakeNative(func(__e *ControlFlow) {
C6798 := __e.Get(1)
_ = C6798
tmp21177 := Call(__e, PrimFunc(symshen_4newpv), B6795)


let__19004 := tmp21177
_ = let__19004

tmp21178 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__19004, Nil)
}
__typedArg0 := let__19004
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21179 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp21178)
}
__typedArg0 := symlist
__typedArg1 := tmp21178
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21180 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__19004, Nil)
}
__typedArg0 := let__19004
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21181 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp21180)
}
__typedArg0 := symlist
__typedArg1 := tmp21180
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21182 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21181, Nil)
}
__typedArg0 := tmp21181
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21183 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21182)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21182
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21184 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21179, tmp21183)
}
__typedArg0 := tmp21179
__typedArg1 := tmp21183
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21185 := Call(__e, PrimFunc(symis_b), V6799, tmp21184, B6795, L6796, Key6797, C6798)


__e.TailApply(PrimFunc(symshen_4gc), B6795, tmp21185)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21186 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21187 := Call(__e, PrimFunc(symshen_4assoc_1_6), symtail, tmp21176, tmp21186)


tmp21188 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21187)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21187
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21188

tmp21189 := MakeNative(func(__e *ControlFlow) {
V6804 := __e.Get(1)
_ = V6804
__e.Return(MakeNative(func(__e *ControlFlow) {
B6800 := __e.Get(1)
_ = B6800
__e.Return(MakeNative(func(__e *ControlFlow) {
L6801 := __e.Get(1)
_ = L6801
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6802 := __e.Get(1)
_ = Key6802
__e.Return(MakeNative(func(__e *ControlFlow) {
C6803 := __e.Get(1)
_ = C6803
tmp21190 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21191 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21190)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21190
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21192 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp21191)
}
__typedArg0 := symstring
__typedArg1 := tmp21191
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6804, tmp21192, B6800, L6801, Key6802, C6803)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21193 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21194 := Call(__e, PrimFunc(symshen_4assoc_1_6), symtlstr, tmp21189, tmp21193)


tmp21195 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21194)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21194
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21195

tmp21196 := MakeNative(func(__e *ControlFlow) {
V6810 := __e.Get(1)
_ = V6810
__e.Return(MakeNative(func(__e *ControlFlow) {
B6806 := __e.Get(1)
_ = B6806
__e.Return(MakeNative(func(__e *ControlFlow) {
L6807 := __e.Get(1)
_ = L6807
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6808 := __e.Get(1)
_ = Key6808
__e.Return(MakeNative(func(__e *ControlFlow) {
C6809 := __e.Get(1)
_ = C6809
tmp21197 := Call(__e, PrimFunc(symshen_4newpv), B6806)


let__19005 := tmp21197
_ = let__19005

tmp21198 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__19005, Nil)
}
__typedArg0 := let__19005
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21199 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp21198)
}
__typedArg0 := symvector
__typedArg1 := tmp21198
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21200 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__19005, Nil)
}
__typedArg0 := let__19005
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21201 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp21200)
}
__typedArg0 := symvector
__typedArg1 := tmp21200
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21202 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21201, Nil)
}
__typedArg0 := tmp21201
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21203 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21202)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21202
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21204 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21199, tmp21203)
}
__typedArg0 := tmp21199
__typedArg1 := tmp21203
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21205 := Call(__e, PrimFunc(symis_b), V6810, tmp21204, B6806, L6807, Key6808, C6809)


__e.TailApply(PrimFunc(symshen_4gc), B6806, tmp21205)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21206 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21207 := Call(__e, PrimFunc(symshen_4assoc_1_6), symtlv, tmp21196, tmp21206)


tmp21208 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21207)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21207
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21208

tmp21209 := MakeNative(func(__e *ControlFlow) {
V6815 := __e.Get(1)
_ = V6815
__e.Return(MakeNative(func(__e *ControlFlow) {
B6811 := __e.Get(1)
_ = B6811
__e.Return(MakeNative(func(__e *ControlFlow) {
L6812 := __e.Get(1)
_ = L6812
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6813 := __e.Get(1)
_ = Key6813
__e.Return(MakeNative(func(__e *ControlFlow) {
C6814 := __e.Get(1)
_ = C6814
tmp21210 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21211 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21210)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21210
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21212 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp21211)
}
__typedArg0 := symsymbol
__typedArg1 := tmp21211
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6815, tmp21212, B6811, L6812, Key6813, C6814)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21213 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21214 := Call(__e, PrimFunc(symshen_4assoc_1_6), symtc, tmp21209, tmp21213)


tmp21215 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21214)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21214
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21215

tmp21216 := MakeNative(func(__e *ControlFlow) {
V6820 := __e.Get(1)
_ = V6820
__e.Return(MakeNative(func(__e *ControlFlow) {
B6816 := __e.Get(1)
_ = B6816
__e.Return(MakeNative(func(__e *ControlFlow) {
L6817 := __e.Get(1)
_ = L6817
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6818 := __e.Get(1)
_ = Key6818
__e.Return(MakeNative(func(__e *ControlFlow) {
C6819 := __e.Get(1)
_ = C6819
tmp21217 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21218 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21217)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21217
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6820, tmp21218, B6816, L6817, Key6818, C6819)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21219 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21220 := Call(__e, PrimFunc(symshen_4assoc_1_6), symtc_2, tmp21216, tmp21219)


tmp21221 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21220)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21220
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21221

tmp21222 := MakeNative(func(__e *ControlFlow) {
V6825 := __e.Get(1)
_ = V6825
__e.Return(MakeNative(func(__e *ControlFlow) {
B6821 := __e.Get(1)
_ = B6821
__e.Return(MakeNative(func(__e *ControlFlow) {
L6822 := __e.Get(1)
_ = L6822
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6823 := __e.Get(1)
_ = Key6823
__e.Return(MakeNative(func(__e *ControlFlow) {
C6824 := __e.Get(1)
_ = C6824
tmp21223 := Call(__e, PrimFunc(symshen_4newpv), B6821)


let__19006 := tmp21223
_ = let__19006

tmp21224 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__19006, Nil)
}
__typedArg0 := let__19006
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21225 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlazy, tmp21224)
}
__typedArg0 := symlazy
__typedArg1 := tmp21224
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21226 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__19006, Nil)
}
__typedArg0 := let__19006
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21227 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21226)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21226
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21228 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21225, tmp21227)
}
__typedArg0 := tmp21225
__typedArg1 := tmp21227
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21229 := Call(__e, PrimFunc(symis_b), V6825, tmp21228, B6821, L6822, Key6823, C6824)


__e.TailApply(PrimFunc(symshen_4gc), B6821, tmp21229)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21230 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21231 := Call(__e, PrimFunc(symshen_4assoc_1_6), symthaw, tmp21222, tmp21230)


tmp21232 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21231)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21231
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21232

tmp21233 := MakeNative(func(__e *ControlFlow) {
V6830 := __e.Get(1)
_ = V6830
__e.Return(MakeNative(func(__e *ControlFlow) {
B6826 := __e.Get(1)
_ = B6826
__e.Return(MakeNative(func(__e *ControlFlow) {
L6827 := __e.Get(1)
_ = L6827
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6828 := __e.Get(1)
_ = Key6828
__e.Return(MakeNative(func(__e *ControlFlow) {
C6829 := __e.Get(1)
_ = C6829
tmp21234 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21235 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21234)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21234
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21236 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp21235)
}
__typedArg0 := symsymbol
__typedArg1 := tmp21235
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6830, tmp21236, B6826, L6827, Key6828, C6829)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21237 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21238 := Call(__e, PrimFunc(symshen_4assoc_1_6), symtrack, tmp21233, tmp21237)


tmp21239 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21238)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21238
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21239

tmp21240 := MakeNative(func(__e *ControlFlow) {
V6835 := __e.Get(1)
_ = V6835
__e.Return(MakeNative(func(__e *ControlFlow) {
B6831 := __e.Get(1)
_ = B6831
__e.Return(MakeNative(func(__e *ControlFlow) {
L6832 := __e.Get(1)
_ = L6832
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6833 := __e.Get(1)
_ = Key6833
__e.Return(MakeNative(func(__e *ControlFlow) {
C6834 := __e.Get(1)
_ = C6834
tmp21241 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21242 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp21241)
}
__typedArg0 := symlist
__typedArg1 := tmp21241
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21243 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21242, Nil)
}
__typedArg0 := tmp21242
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21244 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21243)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21243
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6835, tmp21244, B6831, L6832, Key6833, C6834)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21245 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21246 := Call(__e, PrimFunc(symshen_4assoc_1_6), symtracked, tmp21240, tmp21245)


tmp21247 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21246)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21246
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21247

tmp21248 := MakeNative(func(__e *ControlFlow) {
V6840 := __e.Get(1)
_ = V6840
__e.Return(MakeNative(func(__e *ControlFlow) {
B6836 := __e.Get(1)
_ = B6836
__e.Return(MakeNative(func(__e *ControlFlow) {
L6837 := __e.Get(1)
_ = L6837
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6838 := __e.Get(1)
_ = Key6838
__e.Return(MakeNative(func(__e *ControlFlow) {
C6839 := __e.Get(1)
_ = C6839
tmp21249 := Call(__e, PrimFunc(symshen_4newpv), B6836)


let__19007 := tmp21249
_ = let__19007

tmp21250 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__19007, Nil)
}
__typedArg0 := let__19007
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21251 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21250)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21250
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21252 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symexception, tmp21251)
}
__typedArg0 := symexception
__typedArg1 := tmp21251
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21253 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__19007, Nil)
}
__typedArg0 := let__19007
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21254 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21253)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21253
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21255 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21252, tmp21254)
}
__typedArg0 := tmp21252
__typedArg1 := tmp21254
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21256 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21255, Nil)
}
__typedArg0 := tmp21255
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21257 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21256)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21256
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21258 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__19007, tmp21257)
}
__typedArg0 := let__19007
__typedArg1 := tmp21257
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21259 := Call(__e, PrimFunc(symis_b), V6840, tmp21258, B6836, L6837, Key6838, C6839)


__e.TailApply(PrimFunc(symshen_4gc), B6836, tmp21259)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21260 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21261 := Call(__e, PrimFunc(symshen_4assoc_1_6), symtrap_1error, tmp21248, tmp21260)


tmp21262 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21261)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21261
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21262

tmp21263 := MakeNative(func(__e *ControlFlow) {
V6845 := __e.Get(1)
_ = V6845
__e.Return(MakeNative(func(__e *ControlFlow) {
B6841 := __e.Get(1)
_ = B6841
__e.Return(MakeNative(func(__e *ControlFlow) {
L6842 := __e.Get(1)
_ = L6842
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6843 := __e.Get(1)
_ = Key6843
__e.Return(MakeNative(func(__e *ControlFlow) {
C6844 := __e.Get(1)
_ = C6844
tmp21264 := Call(__e, PrimFunc(symshen_4newpv), B6841)


let__19008 := tmp21264
_ = let__19008

tmp21265 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21266 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21265)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21265
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21267 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__19008, tmp21266)
}
__typedArg0 := let__19008
__typedArg1 := tmp21266
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21268 := Call(__e, PrimFunc(symis_b), V6845, tmp21267, B6841, L6842, Key6843, C6844)


__e.TailApply(PrimFunc(symshen_4gc), B6841, tmp21268)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21269 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21270 := Call(__e, PrimFunc(symshen_4assoc_1_6), symtuple_2, tmp21263, tmp21269)


tmp21271 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21270)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21270
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21271

tmp21272 := MakeNative(func(__e *ControlFlow) {
V6850 := __e.Get(1)
_ = V6850
__e.Return(MakeNative(func(__e *ControlFlow) {
B6846 := __e.Get(1)
_ = B6846
__e.Return(MakeNative(func(__e *ControlFlow) {
L6847 := __e.Get(1)
_ = L6847
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6848 := __e.Get(1)
_ = Key6848
__e.Return(MakeNative(func(__e *ControlFlow) {
C6849 := __e.Get(1)
_ = C6849
tmp21273 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21274 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp21273)
}
__typedArg0 := symlist
__typedArg1 := tmp21273
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21275 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21274, Nil)
}
__typedArg0 := tmp21274
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21276 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21275)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21275
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21277 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp21276)
}
__typedArg0 := symstring
__typedArg1 := tmp21276
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6850, tmp21277, B6846, L6847, Key6848, C6849)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21278 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21279 := Call(__e, PrimFunc(symshen_4assoc_1_6), symunabsolute, tmp21272, tmp21278)


tmp21280 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21279)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21279
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21280

tmp21281 := MakeNative(func(__e *ControlFlow) {
V6855 := __e.Get(1)
_ = V6855
__e.Return(MakeNative(func(__e *ControlFlow) {
B6851 := __e.Get(1)
_ = B6851
__e.Return(MakeNative(func(__e *ControlFlow) {
L6852 := __e.Get(1)
_ = L6852
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6853 := __e.Get(1)
_ = Key6853
__e.Return(MakeNative(func(__e *ControlFlow) {
C6854 := __e.Get(1)
_ = C6854
tmp21282 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21283 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21282)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21282
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21284 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp21283)
}
__typedArg0 := symsymbol
__typedArg1 := tmp21283
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6855, tmp21284, B6851, L6852, Key6853, C6854)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21285 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21286 := Call(__e, PrimFunc(symshen_4assoc_1_6), symundefmacro, tmp21281, tmp21285)


tmp21287 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21286)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21286
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21287

tmp21288 := MakeNative(func(__e *ControlFlow) {
V6862 := __e.Get(1)
_ = V6862
__e.Return(MakeNative(func(__e *ControlFlow) {
B6858 := __e.Get(1)
_ = B6858
__e.Return(MakeNative(func(__e *ControlFlow) {
L6859 := __e.Get(1)
_ = L6859
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6860 := __e.Get(1)
_ = Key6860
__e.Return(MakeNative(func(__e *ControlFlow) {
C6861 := __e.Get(1)
_ = C6861
tmp21289 := Call(__e, PrimFunc(symshen_4newpv), B6858)


let__19009 := tmp21289
_ = let__19009

tmp21290 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__19009, Nil)
}
__typedArg0 := let__19009
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21291 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp21290)
}
__typedArg0 := symlist
__typedArg1 := tmp21290
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21292 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__19009, Nil)
}
__typedArg0 := let__19009
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21293 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp21292)
}
__typedArg0 := symlist
__typedArg1 := tmp21292
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21294 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__19009, Nil)
}
__typedArg0 := let__19009
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21295 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp21294)
}
__typedArg0 := symlist
__typedArg1 := tmp21294
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21296 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21295, Nil)
}
__typedArg0 := tmp21295
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21297 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21296)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21296
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21298 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21293, tmp21297)
}
__typedArg0 := tmp21293
__typedArg1 := tmp21297
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21299 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21298, Nil)
}
__typedArg0 := tmp21298
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21300 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21299)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21299
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21301 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21291, tmp21300)
}
__typedArg0 := tmp21291
__typedArg1 := tmp21300
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21302 := Call(__e, PrimFunc(symis_b), V6862, tmp21301, B6858, L6859, Key6860, C6861)


__e.TailApply(PrimFunc(symshen_4gc), B6858, tmp21302)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21303 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21304 := Call(__e, PrimFunc(symshen_4assoc_1_6), symunion, tmp21288, tmp21303)


tmp21305 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21304)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21304
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21305

tmp21306 := MakeNative(func(__e *ControlFlow) {
V6867 := __e.Get(1)
_ = V6867
__e.Return(MakeNative(func(__e *ControlFlow) {
B6863 := __e.Get(1)
_ = B6863
__e.Return(MakeNative(func(__e *ControlFlow) {
L6864 := __e.Get(1)
_ = L6864
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6865 := __e.Get(1)
_ = Key6865
__e.Return(MakeNative(func(__e *ControlFlow) {
C6866 := __e.Get(1)
_ = C6866
tmp21307 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21308 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21307)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21307
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21309 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp21308)
}
__typedArg0 := symsymbol
__typedArg1 := tmp21308
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6867, tmp21309, B6863, L6864, Key6865, C6866)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21310 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21311 := Call(__e, PrimFunc(symshen_4assoc_1_6), symunprofile, tmp21306, tmp21310)


tmp21312 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21311)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21311
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21312

tmp21313 := MakeNative(func(__e *ControlFlow) {
V6872 := __e.Get(1)
_ = V6872
__e.Return(MakeNative(func(__e *ControlFlow) {
B6868 := __e.Get(1)
_ = B6868
__e.Return(MakeNative(func(__e *ControlFlow) {
L6869 := __e.Get(1)
_ = L6869
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6870 := __e.Get(1)
_ = Key6870
__e.Return(MakeNative(func(__e *ControlFlow) {
C6871 := __e.Get(1)
_ = C6871
tmp21314 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21315 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21314)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21314
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21316 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp21315)
}
__typedArg0 := symsymbol
__typedArg1 := tmp21315
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6872, tmp21316, B6868, L6869, Key6870, C6871)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21317 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21318 := Call(__e, PrimFunc(symshen_4assoc_1_6), symuntrack, tmp21313, tmp21317)


tmp21319 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21318)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21318
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21319

tmp21320 := MakeNative(func(__e *ControlFlow) {
V6877 := __e.Get(1)
_ = V6877
__e.Return(MakeNative(func(__e *ControlFlow) {
B6873 := __e.Get(1)
_ = B6873
__e.Return(MakeNative(func(__e *ControlFlow) {
L6874 := __e.Get(1)
_ = L6874
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6875 := __e.Get(1)
_ = Key6875
__e.Return(MakeNative(func(__e *ControlFlow) {
C6876 := __e.Get(1)
_ = C6876
tmp21321 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21322 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp21321)
}
__typedArg0 := symlist
__typedArg1 := tmp21321
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21323 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21322, Nil)
}
__typedArg0 := tmp21322
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21324 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21323)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21323
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6877, tmp21324, B6873, L6874, Key6875, C6876)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21325 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21326 := Call(__e, PrimFunc(symshen_4assoc_1_6), symuserdefs, tmp21320, tmp21325)


tmp21327 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21326)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21326
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21327

tmp21328 := MakeNative(func(__e *ControlFlow) {
V6882 := __e.Get(1)
_ = V6882
__e.Return(MakeNative(func(__e *ControlFlow) {
B6878 := __e.Get(1)
_ = B6878
__e.Return(MakeNative(func(__e *ControlFlow) {
L6879 := __e.Get(1)
_ = L6879
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6880 := __e.Get(1)
_ = Key6880
__e.Return(MakeNative(func(__e *ControlFlow) {
C6881 := __e.Get(1)
_ = C6881
tmp21329 := Call(__e, PrimFunc(symshen_4newpv), B6878)


let__19010 := tmp21329
_ = let__19010

tmp21330 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21331 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21330)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21330
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21332 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__19010, tmp21331)
}
__typedArg0 := let__19010
__typedArg1 := tmp21331
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21333 := Call(__e, PrimFunc(symis_b), V6882, tmp21332, B6878, L6879, Key6880, C6881)


__e.TailApply(PrimFunc(symshen_4gc), B6878, tmp21333)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21334 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21335 := Call(__e, PrimFunc(symshen_4assoc_1_6), symvariable_2, tmp21328, tmp21334)


tmp21336 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21335)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21335
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21336

tmp21337 := MakeNative(func(__e *ControlFlow) {
V6887 := __e.Get(1)
_ = V6887
__e.Return(MakeNative(func(__e *ControlFlow) {
B6883 := __e.Get(1)
_ = B6883
__e.Return(MakeNative(func(__e *ControlFlow) {
L6884 := __e.Get(1)
_ = L6884
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6885 := __e.Get(1)
_ = Key6885
__e.Return(MakeNative(func(__e *ControlFlow) {
C6886 := __e.Get(1)
_ = C6886
tmp21338 := Call(__e, PrimFunc(symshen_4newpv), B6883)


let__19011 := tmp21338
_ = let__19011

tmp21339 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21340 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21339)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21339
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21341 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__19011, tmp21340)
}
__typedArg0 := let__19011
__typedArg1 := tmp21340
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21342 := Call(__e, PrimFunc(symis_b), V6887, tmp21341, B6883, L6884, Key6885, C6886)


__e.TailApply(PrimFunc(symshen_4gc), B6883, tmp21342)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21343 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21344 := Call(__e, PrimFunc(symshen_4assoc_1_6), symvector_2, tmp21337, tmp21343)


tmp21345 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21344)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21344
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21345

tmp21346 := MakeNative(func(__e *ControlFlow) {
V6892 := __e.Get(1)
_ = V6892
__e.Return(MakeNative(func(__e *ControlFlow) {
B6888 := __e.Get(1)
_ = B6888
__e.Return(MakeNative(func(__e *ControlFlow) {
L6889 := __e.Get(1)
_ = L6889
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6890 := __e.Get(1)
_ = Key6890
__e.Return(MakeNative(func(__e *ControlFlow) {
C6891 := __e.Get(1)
_ = C6891
tmp21347 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21348 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21347)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21347
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6892, tmp21348, B6888, L6889, Key6890, C6891)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21349 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21350 := Call(__e, PrimFunc(symshen_4assoc_1_6), symversion, tmp21346, tmp21349)


tmp21351 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21350)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21350
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21351

tmp21352 := MakeNative(func(__e *ControlFlow) {
V6897 := __e.Get(1)
_ = V6897
__e.Return(MakeNative(func(__e *ControlFlow) {
B6893 := __e.Get(1)
_ = B6893
__e.Return(MakeNative(func(__e *ControlFlow) {
L6894 := __e.Get(1)
_ = L6894
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6895 := __e.Get(1)
_ = Key6895
__e.Return(MakeNative(func(__e *ControlFlow) {
C6896 := __e.Get(1)
_ = C6896
tmp21353 := Call(__e, PrimFunc(symshen_4newpv), B6893)


let__19012 := tmp21353
_ = let__19012

tmp21354 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__19012, Nil)
}
__typedArg0 := let__19012
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21355 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21354)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21354
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21356 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__19012, tmp21355)
}
__typedArg0 := let__19012
__typedArg1 := tmp21355
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21357 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21356, Nil)
}
__typedArg0 := tmp21356
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21358 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21357)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21357
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21359 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp21358)
}
__typedArg0 := symstring
__typedArg1 := tmp21358
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21360 := Call(__e, PrimFunc(symis_b), V6897, tmp21359, B6893, L6894, Key6895, C6896)


__e.TailApply(PrimFunc(symshen_4gc), B6893, tmp21360)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21361 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21362 := Call(__e, PrimFunc(symshen_4assoc_1_6), symwrite_1to_1file, tmp21352, tmp21361)


tmp21363 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21362)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21362
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21363

tmp21364 := MakeNative(func(__e *ControlFlow) {
V6902 := __e.Get(1)
_ = V6902
__e.Return(MakeNative(func(__e *ControlFlow) {
B6898 := __e.Get(1)
_ = B6898
__e.Return(MakeNative(func(__e *ControlFlow) {
L6899 := __e.Get(1)
_ = L6899
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6900 := __e.Get(1)
_ = Key6900
__e.Return(MakeNative(func(__e *ControlFlow) {
C6901 := __e.Get(1)
_ = C6901
tmp21365 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symout, Nil)
}
__typedArg0 := symout
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21366 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp21365)
}
__typedArg0 := symstream
__typedArg1 := tmp21365
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21367 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21368 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21367)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21367
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21369 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21366, tmp21368)
}
__typedArg0 := tmp21366
__typedArg1 := tmp21368
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21370 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21369, Nil)
}
__typedArg0 := tmp21369
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21371 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21370)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21370
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21372 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp21371)
}
__typedArg0 := symnumber
__typedArg1 := tmp21371
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6902, tmp21372, B6898, L6899, Key6900, C6901)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21373 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21374 := Call(__e, PrimFunc(symshen_4assoc_1_6), symwrite_1byte, tmp21364, tmp21373)


tmp21375 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21374)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21374
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21375

tmp21376 := MakeNative(func(__e *ControlFlow) {
V6907 := __e.Get(1)
_ = V6907
__e.Return(MakeNative(func(__e *ControlFlow) {
B6903 := __e.Get(1)
_ = B6903
__e.Return(MakeNative(func(__e *ControlFlow) {
L6904 := __e.Get(1)
_ = L6904
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6905 := __e.Get(1)
_ = Key6905
__e.Return(MakeNative(func(__e *ControlFlow) {
C6906 := __e.Get(1)
_ = C6906
tmp21377 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symout, Nil)
}
__typedArg0 := symout
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21378 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp21377)
}
__typedArg0 := symstream
__typedArg1 := tmp21377
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21379 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21380 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21379)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21379
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21381 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21378, tmp21380)
}
__typedArg0 := tmp21378
__typedArg1 := tmp21380
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21382 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21381, Nil)
}
__typedArg0 := tmp21381
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21383 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21382)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21382
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21384 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp21383)
}
__typedArg0 := symstring
__typedArg1 := tmp21383
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6907, tmp21384, B6903, L6904, Key6905, C6906)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21385 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21386 := Call(__e, PrimFunc(symshen_4assoc_1_6), symwrite_1string, tmp21376, tmp21385)


tmp21387 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21386)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21386
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21387

tmp21388 := MakeNative(func(__e *ControlFlow) {
V6912 := __e.Get(1)
_ = V6912
__e.Return(MakeNative(func(__e *ControlFlow) {
B6908 := __e.Get(1)
_ = B6908
__e.Return(MakeNative(func(__e *ControlFlow) {
L6909 := __e.Get(1)
_ = L6909
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6910 := __e.Get(1)
_ = Key6910
__e.Return(MakeNative(func(__e *ControlFlow) {
C6911 := __e.Get(1)
_ = C6911
tmp21389 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21390 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21389)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21389
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21391 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp21390)
}
__typedArg0 := symstring
__typedArg1 := tmp21390
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6912, tmp21391, B6908, L6909, Key6910, C6911)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21392 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21393 := Call(__e, PrimFunc(symshen_4assoc_1_6), symy_1or_1n_2, tmp21388, tmp21392)


tmp21394 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21393)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21393
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21394

tmp21395 := MakeNative(func(__e *ControlFlow) {
V6917 := __e.Get(1)
_ = V6917
__e.Return(MakeNative(func(__e *ControlFlow) {
B6913 := __e.Get(1)
_ = B6913
__e.Return(MakeNative(func(__e *ControlFlow) {
L6914 := __e.Get(1)
_ = L6914
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6915 := __e.Get(1)
_ = Key6915
__e.Return(MakeNative(func(__e *ControlFlow) {
C6916 := __e.Get(1)
_ = C6916
tmp21396 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21397 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21396)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21396
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21398 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp21397)
}
__typedArg0 := symnumber
__typedArg1 := tmp21397
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21399 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21398, Nil)
}
__typedArg0 := tmp21398
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21400 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21399)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21399
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21401 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp21400)
}
__typedArg0 := symnumber
__typedArg1 := tmp21400
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6917, tmp21401, B6913, L6914, Key6915, C6916)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21402 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21403 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_6, tmp21395, tmp21402)


tmp21404 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21403)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21403
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21404

tmp21405 := MakeNative(func(__e *ControlFlow) {
V6922 := __e.Get(1)
_ = V6922
__e.Return(MakeNative(func(__e *ControlFlow) {
B6918 := __e.Get(1)
_ = B6918
__e.Return(MakeNative(func(__e *ControlFlow) {
L6919 := __e.Get(1)
_ = L6919
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6920 := __e.Get(1)
_ = Key6920
__e.Return(MakeNative(func(__e *ControlFlow) {
C6921 := __e.Get(1)
_ = C6921
tmp21406 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21407 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21406)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21406
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21408 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp21407)
}
__typedArg0 := symnumber
__typedArg1 := tmp21407
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21409 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21408, Nil)
}
__typedArg0 := tmp21408
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21410 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21409)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21409
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21411 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp21410)
}
__typedArg0 := symnumber
__typedArg1 := tmp21410
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6922, tmp21411, B6918, L6919, Key6920, C6921)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21412 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21413 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_5, tmp21405, tmp21412)


tmp21414 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21413)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21413
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21414

tmp21415 := MakeNative(func(__e *ControlFlow) {
V6927 := __e.Get(1)
_ = V6927
__e.Return(MakeNative(func(__e *ControlFlow) {
B6923 := __e.Get(1)
_ = B6923
__e.Return(MakeNative(func(__e *ControlFlow) {
L6924 := __e.Get(1)
_ = L6924
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6925 := __e.Get(1)
_ = Key6925
__e.Return(MakeNative(func(__e *ControlFlow) {
C6926 := __e.Get(1)
_ = C6926
tmp21416 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21417 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21416)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21416
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21418 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp21417)
}
__typedArg0 := symnumber
__typedArg1 := tmp21417
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21419 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21418, Nil)
}
__typedArg0 := tmp21418
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21420 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21419)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21419
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21421 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp21420)
}
__typedArg0 := symnumber
__typedArg1 := tmp21420
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6927, tmp21421, B6923, L6924, Key6925, C6926)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21422 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21423 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_6_a, tmp21415, tmp21422)


tmp21424 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21423)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21423
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21424

tmp21425 := MakeNative(func(__e *ControlFlow) {
V6932 := __e.Get(1)
_ = V6932
__e.Return(MakeNative(func(__e *ControlFlow) {
B6928 := __e.Get(1)
_ = B6928
__e.Return(MakeNative(func(__e *ControlFlow) {
L6929 := __e.Get(1)
_ = L6929
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6930 := __e.Get(1)
_ = Key6930
__e.Return(MakeNative(func(__e *ControlFlow) {
C6931 := __e.Get(1)
_ = C6931
tmp21426 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21427 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21426)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21426
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21428 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp21427)
}
__typedArg0 := symnumber
__typedArg1 := tmp21427
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21429 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21428, Nil)
}
__typedArg0 := tmp21428
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21430 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21429)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21429
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21431 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp21430)
}
__typedArg0 := symnumber
__typedArg1 := tmp21430
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6932, tmp21431, B6928, L6929, Key6930, C6931)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21432 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21433 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_5_a, tmp21425, tmp21432)


tmp21434 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21433)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21433
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21434

tmp21435 := MakeNative(func(__e *ControlFlow) {
V6937 := __e.Get(1)
_ = V6937
__e.Return(MakeNative(func(__e *ControlFlow) {
B6933 := __e.Get(1)
_ = B6933
__e.Return(MakeNative(func(__e *ControlFlow) {
L6934 := __e.Get(1)
_ = L6934
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6935 := __e.Get(1)
_ = Key6935
__e.Return(MakeNative(func(__e *ControlFlow) {
C6936 := __e.Get(1)
_ = C6936
tmp21436 := Call(__e, PrimFunc(symshen_4newpv), B6933)


let__19013 := tmp21436
_ = let__19013

tmp21437 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21438 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21437)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21437
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21439 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__19013, tmp21438)
}
__typedArg0 := let__19013
__typedArg1 := tmp21438
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21440 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21439, Nil)
}
__typedArg0 := tmp21439
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21441 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21440)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21440
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21442 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__19013, tmp21441)
}
__typedArg0 := let__19013
__typedArg1 := tmp21441
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21443 := Call(__e, PrimFunc(symis_b), V6937, tmp21442, B6933, L6934, Key6935, C6936)


__e.TailApply(PrimFunc(symshen_4gc), B6933, tmp21443)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21444 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21445 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_a, tmp21435, tmp21444)


tmp21446 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21445)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21445
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21446

tmp21447 := MakeNative(func(__e *ControlFlow) {
V6942 := __e.Get(1)
_ = V6942
__e.Return(MakeNative(func(__e *ControlFlow) {
B6938 := __e.Get(1)
_ = B6938
__e.Return(MakeNative(func(__e *ControlFlow) {
L6939 := __e.Get(1)
_ = L6939
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6940 := __e.Get(1)
_ = Key6940
__e.Return(MakeNative(func(__e *ControlFlow) {
C6941 := __e.Get(1)
_ = C6941
tmp21448 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21449 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21448)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21448
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21450 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp21449)
}
__typedArg0 := symnumber
__typedArg1 := tmp21449
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21451 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21450, Nil)
}
__typedArg0 := tmp21450
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21452 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21451)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21451
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21453 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp21452)
}
__typedArg0 := symnumber
__typedArg1 := tmp21452
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6942, tmp21453, B6938, L6939, Key6940, C6941)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21454 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21455 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_7, tmp21447, tmp21454)


tmp21456 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21455)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21455
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21456

tmp21457 := MakeNative(func(__e *ControlFlow) {
V6947 := __e.Get(1)
_ = V6947
__e.Return(MakeNative(func(__e *ControlFlow) {
B6943 := __e.Get(1)
_ = B6943
__e.Return(MakeNative(func(__e *ControlFlow) {
L6944 := __e.Get(1)
_ = L6944
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6945 := __e.Get(1)
_ = Key6945
__e.Return(MakeNative(func(__e *ControlFlow) {
C6946 := __e.Get(1)
_ = C6946
tmp21458 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21459 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21458)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21458
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21460 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp21459)
}
__typedArg0 := symnumber
__typedArg1 := tmp21459
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21461 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21460, Nil)
}
__typedArg0 := tmp21460
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21462 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21461)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21461
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21463 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp21462)
}
__typedArg0 := symnumber
__typedArg1 := tmp21462
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6947, tmp21463, B6943, L6944, Key6945, C6946)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21464 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21465 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_c, tmp21457, tmp21464)


tmp21466 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21465)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21465
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21466

tmp21467 := MakeNative(func(__e *ControlFlow) {
V6952 := __e.Get(1)
_ = V6952
__e.Return(MakeNative(func(__e *ControlFlow) {
B6948 := __e.Get(1)
_ = B6948
__e.Return(MakeNative(func(__e *ControlFlow) {
L6949 := __e.Get(1)
_ = L6949
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6950 := __e.Get(1)
_ = Key6950
__e.Return(MakeNative(func(__e *ControlFlow) {
C6951 := __e.Get(1)
_ = C6951
tmp21468 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21469 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21468)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21468
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21470 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp21469)
}
__typedArg0 := symnumber
__typedArg1 := tmp21469
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21471 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21470, Nil)
}
__typedArg0 := tmp21470
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21472 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21471)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21471
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21473 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp21472)
}
__typedArg0 := symnumber
__typedArg1 := tmp21472
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6952, tmp21473, B6948, L6949, Key6950, C6951)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21474 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21475 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_1, tmp21467, tmp21474)


tmp21476 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21475)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21475
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21476

tmp21477 := MakeNative(func(__e *ControlFlow) {
V6957 := __e.Get(1)
_ = V6957
__e.Return(MakeNative(func(__e *ControlFlow) {
B6953 := __e.Get(1)
_ = B6953
__e.Return(MakeNative(func(__e *ControlFlow) {
L6954 := __e.Get(1)
_ = L6954
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6955 := __e.Get(1)
_ = Key6955
__e.Return(MakeNative(func(__e *ControlFlow) {
C6956 := __e.Get(1)
_ = C6956
tmp21478 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21479 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21478)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21478
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21480 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp21479)
}
__typedArg0 := symnumber
__typedArg1 := tmp21479
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21481 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21480, Nil)
}
__typedArg0 := tmp21480
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21482 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21481)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21481
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21483 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp21482)
}
__typedArg0 := symnumber
__typedArg1 := tmp21482
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6957, tmp21483, B6953, L6954, Key6955, C6956)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21484 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21485 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_d, tmp21477, tmp21484)


tmp21486 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21485)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21485
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21486

tmp21487 := MakeNative(func(__e *ControlFlow) {
V6962 := __e.Get(1)
_ = V6962
__e.Return(MakeNative(func(__e *ControlFlow) {
B6958 := __e.Get(1)
_ = B6958
__e.Return(MakeNative(func(__e *ControlFlow) {
L6959 := __e.Get(1)
_ = L6959
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6960 := __e.Get(1)
_ = Key6960
__e.Return(MakeNative(func(__e *ControlFlow) {
C6961 := __e.Get(1)
_ = C6961
tmp21488 := Call(__e, PrimFunc(symshen_4newpv), B6958)


let__19014 := tmp21488
_ = let__19014

tmp21489 := Call(__e, PrimFunc(symshen_4newpv), B6958)


let__19015 := tmp21489
_ = let__19015

tmp21490 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21491 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21490)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21490
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21492 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__19015, tmp21491)
}
__typedArg0 := let__19015
__typedArg1 := tmp21491
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21493 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21492, Nil)
}
__typedArg0 := tmp21492
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21494 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21493)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21493
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21495 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__19014, tmp21494)
}
__typedArg0 := let__19014
__typedArg1 := tmp21494
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21496 := Call(__e, PrimFunc(symis_b), V6962, tmp21495, B6958, L6959, Key6960, C6961)


tmp21497 := Call(__e, PrimFunc(symshen_4gc), B6958, tmp21496)


__e.TailApply(PrimFunc(symshen_4gc), B6958, tmp21497)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21498 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21499 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_a_a, tmp21487, tmp21498)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21499)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21499
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return


}, 0)

tmp21500 := Call(__e, ns2_1set, symshen_4initialise_1signedfuncs, tmp19781)


_ = tmp21500

tmp21501 := MakeNative(func(__e *ControlFlow) {
tmp21502 := MakeNative(func(__e *ControlFlow) {
Y954 := __e.Get(1)
_ = Y954
__e.TailApply(PrimFunc(symshen_4tuple), Y954)
return
}, 1)

tmp21503 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4tuple, tmp21502)
}
__typedArg0 := symshen_4tuple
__typedArg1 := tmp21502
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21504 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21503)


_ = tmp21504

tmp21505 := MakeNative(func(__e *ControlFlow) {
Y955 := __e.Get(1)
_ = Y955
__e.TailApply(PrimFunc(symshen_4pvar), Y955)
return
}, 1)

tmp21506 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4pvar, tmp21505)
}
__typedArg0 := symshen_4pvar
__typedArg1 := tmp21505
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21507 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21506)


_ = tmp21507

tmp21508 := MakeNative(func(__e *ControlFlow) {
Y956 := __e.Get(1)
_ = Y956
__e.TailApply(PrimFunc(symshen_4dictionary), Y956)
return
}, 1)

tmp21509 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4dictionary, tmp21508)
}
__typedArg0 := symshen_4dictionary
__typedArg1 := tmp21508
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21510 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21509)


_ = tmp21510

tmp21511 := MakeNative(func(__e *ControlFlow) {
Y957 := __e.Get(1)
_ = Y957
__e.TailApply(PrimFunc(symshen_4print_1prolog_1vector), Y957)
return
}, 1)

tmp21512 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4print_1prolog_1vector, tmp21511)
}
__typedArg0 := symshen_4print_1prolog_1vector
__typedArg1 := tmp21511
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21513 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21512)


_ = tmp21513

tmp21514 := MakeNative(func(__e *ControlFlow) {
Y958 := __e.Get(1)
_ = Y958
__e.TailApply(PrimFunc(symshen_4print_1freshterm), Y958)
return
}, 1)

tmp21515 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4print_1freshterm, tmp21514)
}
__typedArg0 := symshen_4print_1freshterm
__typedArg1 := tmp21514
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21516 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21515)


_ = tmp21516

tmp21517 := MakeNative(func(__e *ControlFlow) {
Y959 := __e.Get(1)
_ = Y959
__e.TailApply(PrimFunc(symshen_4printF), Y959)
return
}, 1)

tmp21518 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4printF, tmp21517)
}
__typedArg0 := symshen_4printF
__typedArg1 := tmp21517
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21519 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21518)


_ = tmp21519

tmp21520 := MakeNative(func(__e *ControlFlow) {
Y960 := __e.Get(1)
_ = Y960
__e.TailApply(PrimFunc(symabsolute), Y960)
return
}, 1)

tmp21521 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symabsolute, tmp21520)
}
__typedArg0 := symabsolute
__typedArg1 := tmp21520
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21522 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21521)


_ = tmp21522

tmp21523 := MakeNative(func(__e *ControlFlow) {
Y961 := __e.Get(1)
_ = Y961
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symabsvector_2) {
return PrimIsVector(Y961)
}
__typedArg0 := Y961
return Call(__e, PrimFunc(symabsvector_2), __typedArg0)
})())
return
}, 1)

tmp21524 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symabsvector_2, tmp21523)
}
__typedArg0 := symabsvector_2
__typedArg1 := tmp21523
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21525 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21524)


_ = tmp21525

tmp21526 := MakeNative(func(__e *ControlFlow) {
Y962 := __e.Get(1)
_ = Y962
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symabsvector) {
return PrimAbsvector(Y962)
}
__typedArg0 := Y962
return Call(__e, PrimFunc(symabsvector), __typedArg0)
})())
return
}, 1)

tmp21527 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symabsvector, tmp21526)
}
__typedArg0 := symabsvector
__typedArg1 := tmp21526
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21528 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21527)


_ = tmp21528

tmp21529 := MakeNative(func(__e *ControlFlow) {
Y963 := __e.Get(1)
_ = Y963
__e.Return(MakeNative(func(__e *ControlFlow) {
Y964 := __e.Get(1)
_ = Y964
__e.Return(MakeNative(func(__e *ControlFlow) {
Y965 := __e.Get(1)
_ = Y965
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(Y963, Y964, Y965)
}
__typedArg0 := Y963
__typedArg1 := Y964
__typedArg2 := Y965
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})())
return
}, 1))
return
}, 1))
return
}, 1)

tmp21530 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symaddress_1_6, tmp21529)
}
__typedArg0 := symaddress_1_6
__typedArg1 := tmp21529
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21531 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21530)


_ = tmp21531

tmp21532 := MakeNative(func(__e *ControlFlow) {
Y966 := __e.Get(1)
_ = Y966
__e.Return(MakeNative(func(__e *ControlFlow) {
Y967 := __e.Get(1)
_ = Y967
__e.TailApply(PrimFunc(symadjoin), Y966, Y967)
return
}, 1))
return
}, 1)

tmp21533 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symadjoin, tmp21532)
}
__typedArg0 := symadjoin
__typedArg1 := tmp21532
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21534 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21533)


_ = tmp21534

tmp21535 := MakeNative(func(__e *ControlFlow) {
Y968 := __e.Get(1)
_ = Y968
__e.Return(MakeNative(func(__e *ControlFlow) {
Y969 := __e.Get(1)
_ = Y969
if True == Y968 {
if True == Y969 {
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
}, 1))
return
}, 1)

tmp21538 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symand, tmp21535)
}
__typedArg0 := symand
__typedArg1 := tmp21535
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21539 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21538)


_ = tmp21539

tmp21540 := MakeNative(func(__e *ControlFlow) {
Y970 := __e.Get(1)
_ = Y970
__e.Return(MakeNative(func(__e *ControlFlow) {
Y971 := __e.Get(1)
_ = Y971
__e.TailApply(PrimFunc(symappend), Y970, Y971)
return
}, 1))
return
}, 1)

tmp21541 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symappend, tmp21540)
}
__typedArg0 := symappend
__typedArg1 := tmp21540
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21542 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21541)


_ = tmp21542

tmp21543 := MakeNative(func(__e *ControlFlow) {
Y972 := __e.Get(1)
_ = Y972
__e.TailApply(PrimFunc(symarity), Y972)
return
}, 1)

tmp21544 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symarity, tmp21543)
}
__typedArg0 := symarity
__typedArg1 := tmp21543
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21545 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21544)


_ = tmp21545

tmp21546 := MakeNative(func(__e *ControlFlow) {
Y973 := __e.Get(1)
_ = Y973
__e.Return(MakeNative(func(__e *ControlFlow) {
Y974 := __e.Get(1)
_ = Y974
__e.TailApply(PrimFunc(symassoc), Y973, Y974)
return
}, 1))
return
}, 1)

tmp21547 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symassoc, tmp21546)
}
__typedArg0 := symassoc
__typedArg1 := tmp21546
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21548 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21547)


_ = tmp21548

tmp21549 := MakeNative(func(__e *ControlFlow) {
Y975 := __e.Get(1)
_ = Y975
__e.TailApply(PrimFunc(symatom_2), Y975)
return
}, 1)

tmp21550 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symatom_2, tmp21549)
}
__typedArg0 := symatom_2
__typedArg1 := tmp21549
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21551 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21550)


_ = tmp21551

tmp21552 := MakeNative(func(__e *ControlFlow) {
Y976 := __e.Get(1)
_ = Y976
__e.TailApply(PrimFunc(symboolean_2), Y976)
return
}, 1)

tmp21553 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean_2, tmp21552)
}
__typedArg0 := symboolean_2
__typedArg1 := tmp21552
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21554 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21553)


_ = tmp21554

tmp21555 := MakeNative(func(__e *ControlFlow) {
Y977 := __e.Get(1)
_ = Y977
__e.TailApply(PrimFunc(symbootstrap), Y977)
return
}, 1)

tmp21556 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symbootstrap, tmp21555)
}
__typedArg0 := symbootstrap
__typedArg1 := tmp21555
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21557 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21556)


_ = tmp21557

tmp21558 := MakeNative(func(__e *ControlFlow) {
Y978 := __e.Get(1)
_ = Y978
__e.TailApply(PrimFunc(symbound_2), Y978)
return
}, 1)

tmp21559 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symbound_2, tmp21558)
}
__typedArg0 := symbound_2
__typedArg1 := tmp21558
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21560 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21559)


_ = tmp21560

tmp21561 := MakeNative(func(__e *ControlFlow) {
Y979 := __e.Get(1)
_ = Y979
__e.Return(MakeNative(func(__e *ControlFlow) {
Y980 := __e.Get(1)
_ = Y980
__e.Return(MakeNative(func(__e *ControlFlow) {
Y981 := __e.Get(1)
_ = Y981
__e.Return(MakeNative(func(__e *ControlFlow) {
Y982 := __e.Get(1)
_ = Y982
__e.Return(MakeNative(func(__e *ControlFlow) {
Y983 := __e.Get(1)
_ = Y983
__e.Return(MakeNative(func(__e *ControlFlow) {
Y984 := __e.Get(1)
_ = Y984
__e.TailApply(PrimFunc(symbind), Y979, Y980, Y981, Y982, Y983, Y984)
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21562 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symbind, tmp21561)
}
__typedArg0 := symbind
__typedArg1 := tmp21561
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21563 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21562)


_ = tmp21563

tmp21564 := MakeNative(func(__e *ControlFlow) {
Y985 := __e.Get(1)
_ = Y985
__e.Return(MakeNative(func(__e *ControlFlow) {
Y986 := __e.Get(1)
_ = Y986
__e.Return(MakeNative(func(__e *ControlFlow) {
Y987 := __e.Get(1)
_ = Y987
__e.Return(MakeNative(func(__e *ControlFlow) {
Y988 := __e.Get(1)
_ = Y988
__e.Return(MakeNative(func(__e *ControlFlow) {
Y989 := __e.Get(1)
_ = Y989
__e.TailApply(PrimFunc(symcall), Y985, Y986, Y987, Y988, Y989)
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21565 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcall, tmp21564)
}
__typedArg0 := symcall
__typedArg1 := tmp21564
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21566 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21565)


_ = tmp21566

tmp21567 := MakeNative(func(__e *ControlFlow) {
Y990 := __e.Get(1)
_ = Y990
__e.TailApply(PrimFunc(symcd), Y990)
return
}, 1)

tmp21568 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcd, tmp21567)
}
__typedArg0 := symcd
__typedArg1 := tmp21567
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21569 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21568)


_ = tmp21569

tmp21570 := MakeNative(func(__e *ControlFlow) {
Y991 := __e.Get(1)
_ = Y991
__e.TailApply(PrimFunc(symchar_1stream_2), Y991)
return
}, 1)

tmp21571 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symchar_1stream_2, tmp21570)
}
__typedArg0 := symchar_1stream_2
__typedArg1 := tmp21570
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21572 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21571)


_ = tmp21572

tmp21573 := MakeNative(func(__e *ControlFlow) {
Y992 := __e.Get(1)
_ = Y992
__e.Return(MakeNative(func(__e *ControlFlow) {
Y993 := __e.Get(1)
_ = Y993
__e.TailApply(PrimFunc(symcompile), Y992, Y993)
return
}, 1))
return
}, 1)

tmp21574 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcompile, tmp21573)
}
__typedArg0 := symcompile
__typedArg1 := tmp21573
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21575 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21574)


_ = tmp21575

tmp21576 := MakeNative(func(__e *ControlFlow) {
Y994 := __e.Get(1)
_ = Y994
__e.Return(MakeNative(func(__e *ControlFlow) {
Y995 := __e.Get(1)
_ = Y995
__e.TailApply(PrimFunc(symconcat), Y994, Y995)
return
}, 1))
return
}, 1)

tmp21577 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symconcat, tmp21576)
}
__typedArg0 := symconcat
__typedArg1 := tmp21576
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21578 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21577)


_ = tmp21578

tmp21579 := MakeNative(func(__e *ControlFlow) {
Y996 := __e.Get(1)
_ = Y996
__e.Return(MakeNative(func(__e *ControlFlow) {
Y997 := __e.Get(1)
_ = Y997
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Y996, Y997)
}
__typedArg0 := Y996
__typedArg1 := Y997
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return
}, 1))
return
}, 1)

tmp21580 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp21579)
}
__typedArg0 := symcons
__typedArg1 := tmp21579
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21581 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21580)


_ = tmp21581

tmp21582 := MakeNative(func(__e *ControlFlow) {
Y998 := __e.Get(1)
_ = Y998
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(Y998)
}
__typedArg0 := Y998
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})())
return
}, 1)

tmp21583 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons_2, tmp21582)
}
__typedArg0 := symcons_2
__typedArg1 := tmp21582
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21584 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21583)


_ = tmp21584

tmp21585 := MakeNative(func(__e *ControlFlow) {
Y999 := __e.Get(1)
_ = Y999
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1000 := __e.Get(1)
_ = Y1000
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(Y999)
__typedS1, __typedOK1 := TypedString(Y1000)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := Y999
__typedArg1 := Y1000
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
return
}, 1))
return
}, 1)

tmp21586 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcn, tmp21585)
}
__typedArg0 := symcn
__typedArg1 := tmp21585
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21587 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21586)


_ = tmp21587

tmp21588 := MakeNative(func(__e *ControlFlow) {
Y1001 := __e.Get(1)
_ = Y1001
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symclose) {
return PrimCloseStream(Y1001)
}
__typedArg0 := Y1001
return Call(__e, PrimFunc(symclose), __typedArg0)
})())
return
}, 1)

tmp21589 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symclose, tmp21588)
}
__typedArg0 := symclose
__typedArg1 := tmp21588
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21590 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21589)


_ = tmp21590

tmp21591 := MakeNative(func(__e *ControlFlow) {
Y1002 := __e.Get(1)
_ = Y1002
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1003 := __e.Get(1)
_ = Y1003
__e.TailApply(PrimFunc(symdeclare), Y1002, Y1003)
return
}, 1))
return
}, 1)

tmp21592 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdeclare, tmp21591)
}
__typedArg0 := symdeclare
__typedArg1 := tmp21591
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21593 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21592)


_ = tmp21593

tmp21594 := MakeNative(func(__e *ControlFlow) {
Y1004 := __e.Get(1)
_ = Y1004
__e.TailApply(PrimFunc(symdestroy), Y1004)
return
}, 1)

tmp21595 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdestroy, tmp21594)
}
__typedArg0 := symdestroy
__typedArg1 := tmp21594
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21596 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21595)


_ = tmp21596

tmp21597 := MakeNative(func(__e *ControlFlow) {
Y1005 := __e.Get(1)
_ = Y1005
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1006 := __e.Get(1)
_ = Y1006
__e.TailApply(PrimFunc(symdifference), Y1005, Y1006)
return
}, 1))
return
}, 1)

tmp21598 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdifference, tmp21597)
}
__typedArg0 := symdifference
__typedArg1 := tmp21597
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21599 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21598)


_ = tmp21599

tmp21600 := MakeNative(func(__e *ControlFlow) {
Y1007 := __e.Get(1)
_ = Y1007
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1008 := __e.Get(1)
_ = Y1008
_ = Y1007

__e.Return(Y1008)
return


}, 1))
return
}, 1)

tmp21601 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdo, tmp21600)
}
__typedArg0 := symdo
__typedArg1 := tmp21600
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21602 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21601)


_ = tmp21602

tmp21603 := MakeNative(func(__e *ControlFlow) {
Y1009 := __e.Get(1)
_ = Y1009
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1010 := __e.Get(1)
_ = Y1010
__e.TailApply(PrimFunc(symelement_2), Y1009, Y1010)
return
}, 1))
return
}, 1)

tmp21604 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symelement_2, tmp21603)
}
__typedArg0 := symelement_2
__typedArg1 := tmp21603
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21605 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21604)


_ = tmp21605

tmp21606 := MakeNative(func(__e *ControlFlow) {
Y1011 := __e.Get(1)
_ = Y1011
__e.TailApply(PrimFunc(symempty_2), Y1011)
return
}, 1)

tmp21607 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symempty_2, tmp21606)
}
__typedArg0 := symempty_2
__typedArg1 := tmp21606
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21608 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21607)


_ = tmp21608

tmp21609 := MakeNative(func(__e *ControlFlow) {
Y1012 := __e.Get(1)
_ = Y1012
__e.TailApply(PrimFunc(symenable_1type_1theory), Y1012)
return
}, 1)

tmp21610 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symenable_1type_1theory, tmp21609)
}
__typedArg0 := symenable_1type_1theory
__typedArg1 := tmp21609
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21611 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21610)


_ = tmp21611

tmp21612 := MakeNative(func(__e *ControlFlow) {
Y1013 := __e.Get(1)
_ = Y1013
__e.TailApply(PrimFunc(symexternal), Y1013)
return
}, 1)

tmp21613 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symexternal, tmp21612)
}
__typedArg0 := symexternal
__typedArg1 := tmp21612
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21614 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21613)


_ = tmp21614

tmp21615 := MakeNative(func(__e *ControlFlow) {
Y1014 := __e.Get(1)
_ = Y1014
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symerror_1to_1string) {
return PrimErrorToString(Y1014)
}
__typedArg0 := Y1014
return Call(__e, PrimFunc(symerror_1to_1string), __typedArg0)
})())
return
}, 1)

tmp21616 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symerror_1to_1string, tmp21615)
}
__typedArg0 := symerror_1to_1string
__typedArg1 := tmp21615
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21617 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21616)


_ = tmp21617

tmp21618 := MakeNative(func(__e *ControlFlow) {
Y1015 := __e.Get(1)
_ = Y1015
__e.TailApply(PrimFunc(symeval), Y1015)
return
}, 1)

tmp21619 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symeval, tmp21618)
}
__typedArg0 := symeval
__typedArg1 := tmp21618
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21620 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21619)


_ = tmp21620

tmp21621 := MakeNative(func(__e *ControlFlow) {
Y1016 := __e.Get(1)
_ = Y1016
__e.TailApply(PrimFunc(symeval_1kl), Y1016)
return
}, 1)

tmp21622 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symeval_1kl, tmp21621)
}
__typedArg0 := symeval_1kl
__typedArg1 := tmp21621
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21623 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21622)


_ = tmp21623

tmp21624 := MakeNative(func(__e *ControlFlow) {
Y1017 := __e.Get(1)
_ = Y1017
__e.TailApply(PrimFunc(symexplode), Y1017)
return
}, 1)

tmp21625 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symexplode, tmp21624)
}
__typedArg0 := symexplode
__typedArg1 := tmp21624
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21626 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21625)


_ = tmp21626

tmp21627 := MakeNative(func(__e *ControlFlow) {
Y1018 := __e.Get(1)
_ = Y1018
__e.TailApply(PrimFunc(symexternal), Y1018)
return
}, 1)

tmp21628 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symexternal, tmp21627)
}
__typedArg0 := symexternal
__typedArg1 := tmp21627
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21629 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21628)


_ = tmp21629

tmp21630 := MakeNative(func(__e *ControlFlow) {
Y1019 := __e.Get(1)
_ = Y1019
__e.TailApply(PrimFunc(symfactorise), Y1019)
return
}, 1)

tmp21631 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfactorise, tmp21630)
}
__typedArg0 := symfactorise
__typedArg1 := tmp21630
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21632 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21631)


_ = tmp21632

tmp21633 := MakeNative(func(__e *ControlFlow) {
Y1020 := __e.Get(1)
_ = Y1020
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1021 := __e.Get(1)
_ = Y1021
__e.TailApply(PrimFunc(symfail_1if), Y1020, Y1021)
return
}, 1))
return
}, 1)

tmp21634 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfail_1if, tmp21633)
}
__typedArg0 := symfail_1if
__typedArg1 := tmp21633
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21635 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21634)


_ = tmp21635

tmp21636 := MakeNative(func(__e *ControlFlow) {
Y1022 := __e.Get(1)
_ = Y1022
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1023 := __e.Get(1)
_ = Y1023
__e.TailApply(PrimFunc(symfix), Y1022, Y1023)
return
}, 1))
return
}, 1)

tmp21637 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfix, tmp21636)
}
__typedArg0 := symfix
__typedArg1 := tmp21636
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21638 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21637)


_ = tmp21638

tmp21639 := MakeNative(func(__e *ControlFlow) {
Y1024 := __e.Get(1)
_ = Y1024
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1025 := __e.Get(1)
_ = Y1025
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1026 := __e.Get(1)
_ = Y1026
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1027 := __e.Get(1)
_ = Y1027
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1028 := __e.Get(1)
_ = Y1028
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1029 := __e.Get(1)
_ = Y1029
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1030 := __e.Get(1)
_ = Y1030
__e.TailApply(PrimFunc(symfindall), Y1024, Y1025, Y1026, Y1027, Y1028, Y1029, Y1030)
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21640 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfindall, tmp21639)
}
__typedArg0 := symfindall
__typedArg1 := tmp21639
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21641 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21640)


_ = tmp21641

tmp21642 := MakeNative(func(__e *ControlFlow) {
Y1031 := __e.Get(1)
_ = Y1031
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1032 := __e.Get(1)
_ = Y1032
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1033 := __e.Get(1)
_ = Y1033
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1034 := __e.Get(1)
_ = Y1034
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1035 := __e.Get(1)
_ = Y1035
__e.TailApply(PrimFunc(symfork), Y1031, Y1032, Y1033, Y1034, Y1035)
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21643 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfork, tmp21642)
}
__typedArg0 := symfork
__typedArg1 := tmp21642
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21644 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21643)


_ = tmp21644

tmp21645 := MakeNative(func(__e *ControlFlow) {
Y1036 := __e.Get(1)
_ = Y1036
__e.Return(MakeNative(func(__e *ControlFlow) {
__e.Return(Y1036)
return
}, 0))
return
}, 1)

tmp21646 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfreeze, tmp21645)
}
__typedArg0 := symfreeze
__typedArg1 := tmp21645
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21647 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21646)


_ = tmp21647

tmp21648 := MakeNative(func(__e *ControlFlow) {
Y1037 := __e.Get(1)
_ = Y1037
__e.TailApply(PrimFunc(symfst), Y1037)
return
}, 1)

tmp21649 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfst, tmp21648)
}
__typedArg0 := symfst
__typedArg1 := tmp21648
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21650 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21649)


_ = tmp21650

tmp21651 := MakeNative(func(__e *ControlFlow) {
Y1038 := __e.Get(1)
_ = Y1038
__e.TailApply(PrimFunc(symfn), Y1038)
return
}, 1)

tmp21652 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfn, tmp21651)
}
__typedArg0 := symfn
__typedArg1 := tmp21651
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21653 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21652)


_ = tmp21653

tmp21654 := MakeNative(func(__e *ControlFlow) {
Y1039 := __e.Get(1)
_ = Y1039
__e.TailApply(PrimFunc(symfunction), Y1039)
return
}, 1)

tmp21655 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfunction, tmp21654)
}
__typedArg0 := symfunction
__typedArg1 := tmp21654
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21656 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21655)


_ = tmp21656

tmp21657 := MakeNative(func(__e *ControlFlow) {
Y1040 := __e.Get(1)
_ = Y1040
__e.TailApply(PrimFunc(symgensym), Y1040)
return
}, 1)

tmp21658 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symgensym, tmp21657)
}
__typedArg0 := symgensym
__typedArg1 := tmp21657
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21659 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21658)


_ = tmp21659

tmp21660 := MakeNative(func(__e *ControlFlow) {
Y1041 := __e.Get(1)
_ = Y1041
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1042 := __e.Get(1)
_ = Y1042
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1043 := __e.Get(1)
_ = Y1043
__e.TailApply(PrimFunc(symget), Y1041, Y1042, Y1043)
return
}, 1))
return
}, 1))
return
}, 1)

tmp21661 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symget, tmp21660)
}
__typedArg0 := symget
__typedArg1 := tmp21660
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21662 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21661)


_ = tmp21662

tmp21663 := MakeNative(func(__e *ControlFlow) {
Y1044 := __e.Get(1)
_ = Y1044
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symget_1time) {
return PrimGetTime(Y1044)
}
__typedArg0 := Y1044
return Call(__e, PrimFunc(symget_1time), __typedArg0)
})())
return
}, 1)

tmp21664 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symget_1time, tmp21663)
}
__typedArg0 := symget_1time
__typedArg1 := tmp21663
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21665 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21664)


_ = tmp21665

tmp21666 := MakeNative(func(__e *ControlFlow) {
Y1045 := __e.Get(1)
_ = Y1045
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1046 := __e.Get(1)
_ = Y1046
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1047 := __e.Get(1)
_ = Y1047
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(Y1045, Y1046, Y1047)
}
__typedArg0 := Y1045
__typedArg1 := Y1046
__typedArg2 := Y1047
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})())
return
}, 1))
return
}, 1))
return
}, 1)

tmp21667 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symaddress_1_6, tmp21666)
}
__typedArg0 := symaddress_1_6
__typedArg1 := tmp21666
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21668 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21667)


_ = tmp21668

tmp21669 := MakeNative(func(__e *ControlFlow) {
Y1048 := __e.Get(1)
_ = Y1048
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1049 := __e.Get(1)
_ = Y1049
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(Y1048, Y1049)
}
__typedArg0 := Y1048
__typedArg1 := Y1049
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})())
return
}, 1))
return
}, 1)

tmp21670 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1address, tmp21669)
}
__typedArg0 := sym_5_1address
__typedArg1 := tmp21669
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21671 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21670)


_ = tmp21671

tmp21672 := MakeNative(func(__e *ControlFlow) {
Y1050 := __e.Get(1)
_ = Y1050
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1051 := __e.Get(1)
_ = Y1051
__e.TailApply(PrimFunc(sym_5_1vector), Y1050, Y1051)
return
}, 1))
return
}, 1)

tmp21673 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1vector, tmp21672)
}
__typedArg0 := sym_5_1vector
__typedArg1 := tmp21672
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21674 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21673)


_ = tmp21674

tmp21675 := MakeNative(func(__e *ControlFlow) {
Y1052 := __e.Get(1)
_ = Y1052
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1053 := __e.Get(1)
_ = Y1053
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_6) {
__typedN0, __typedOK0 := TypedFloat64(Y1052)
__typedN1, __typedOK1 := TypedFloat64(Y1053)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_6) {
return TypedMaterializeBoolean((__typedN0 > __typedN1))
}}
__typedArg0 := Y1052
__typedArg1 := Y1053
return Call(__e, PrimFunc(sym_6), __typedArg0, __typedArg1)
})())
return
}, 1))
return
}, 1)

tmp21676 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_6, tmp21675)
}
__typedArg0 := sym_6
__typedArg1 := tmp21675
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21677 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21676)


_ = tmp21677

tmp21678 := MakeNative(func(__e *ControlFlow) {
Y1054 := __e.Get(1)
_ = Y1054
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1055 := __e.Get(1)
_ = Y1055
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_6_a) {
__typedN0, __typedOK0 := TypedFloat64(Y1054)
__typedN1, __typedOK1 := TypedFloat64(Y1055)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_6_a) {
return TypedMaterializeBoolean((__typedN0 >= __typedN1))
}}
__typedArg0 := Y1054
__typedArg1 := Y1055
return Call(__e, PrimFunc(sym_6_a), __typedArg0, __typedArg1)
})())
return
}, 1))
return
}, 1)

tmp21679 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_6_a, tmp21678)
}
__typedArg0 := sym_6_a
__typedArg1 := tmp21678
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21680 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21679)


_ = tmp21680

tmp21681 := MakeNative(func(__e *ControlFlow) {
Y1056 := __e.Get(1)
_ = Y1056
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1057 := __e.Get(1)
_ = Y1057
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Y1056, Y1057)
}
__typedArg0 := Y1056
__typedArg1 := Y1057
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})())
return
}, 1))
return
}, 1)

tmp21682 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_a, tmp21681)
}
__typedArg0 := sym_a
__typedArg1 := tmp21681
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21683 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21682)


_ = tmp21683

tmp21684 := MakeNative(func(__e *ControlFlow) {
Y1058 := __e.Get(1)
_ = Y1058
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1059 := __e.Get(1)
_ = Y1059
__e.TailApply(PrimFunc(symhash), Y1058, Y1059)
return
}, 1))
return
}, 1)

tmp21685 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhash, tmp21684)
}
__typedArg0 := symhash
__typedArg1 := tmp21684
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21686 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21685)


_ = tmp21686

tmp21687 := MakeNative(func(__e *ControlFlow) {
Y1060 := __e.Get(1)
_ = Y1060
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(Y1060)
}
__typedArg0 := Y1060
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return
}, 1)

tmp21688 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhd, tmp21687)
}
__typedArg0 := symhd
__typedArg1 := tmp21687
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21689 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21688)


_ = tmp21689

tmp21690 := MakeNative(func(__e *ControlFlow) {
Y1061 := __e.Get(1)
_ = Y1061
__e.TailApply(PrimFunc(symhdv), Y1061)
return
}, 1)

tmp21691 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhdv, tmp21690)
}
__typedArg0 := symhdv
__typedArg1 := tmp21690
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21692 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21691)


_ = tmp21692

tmp21693 := MakeNative(func(__e *ControlFlow) {
Y1062 := __e.Get(1)
_ = Y1062
__e.TailApply(PrimFunc(symhdstr), Y1062)
return
}, 1)

tmp21694 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhdstr, tmp21693)
}
__typedArg0 := symhdstr
__typedArg1 := tmp21693
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21695 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21694)


_ = tmp21695

tmp21696 := MakeNative(func(__e *ControlFlow) {
Y1063 := __e.Get(1)
_ = Y1063
__e.TailApply(PrimFunc(symhead), Y1063)
return
}, 1)

tmp21697 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhead, tmp21696)
}
__typedArg0 := symhead
__typedArg1 := tmp21696
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21698 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21697)


_ = tmp21698

tmp21699 := MakeNative(func(__e *ControlFlow) {
Y1064 := __e.Get(1)
_ = Y1064
__e.TailApply(PrimFunc(symhush), Y1064)
return
}, 1)

tmp21700 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhush, tmp21699)
}
__typedArg0 := symhush
__typedArg1 := tmp21699
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21701 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21700)


_ = tmp21701

tmp21702 := MakeNative(func(__e *ControlFlow) {
Y1065 := __e.Get(1)
_ = Y1065
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1066 := __e.Get(1)
_ = Y1066
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1067 := __e.Get(1)
_ = Y1067
if True == Y1065 {
__e.Return(Y1066)
return
} else {
__e.Return(Y1067)
return
}
}, 1))
return
}, 1))
return
}, 1)

tmp21704 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp21702)
}
__typedArg0 := symif
__typedArg1 := tmp21702
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21705 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21704)


_ = tmp21705

tmp21706 := MakeNative(func(__e *ControlFlow) {
Y1068 := __e.Get(1)
_ = Y1068
__e.TailApply(PrimFunc(syminclude), Y1068)
return
}, 1)

tmp21707 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminclude, tmp21706)
}
__typedArg0 := syminclude
__typedArg1 := tmp21706
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21708 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21707)


_ = tmp21708

tmp21709 := MakeNative(func(__e *ControlFlow) {
Y1069 := __e.Get(1)
_ = Y1069
__e.TailApply(PrimFunc(symin_1package), Y1069)
return
}, 1)

tmp21710 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symin_1package, tmp21709)
}
__typedArg0 := symin_1package
__typedArg1 := tmp21709
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21711 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21710)


_ = tmp21711

tmp21712 := MakeNative(func(__e *ControlFlow) {
Y1070 := __e.Get(1)
_ = Y1070
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(syminteger_2) {
return PrimIsInteger(Y1070)
}
__typedArg0 := Y1070
return Call(__e, PrimFunc(syminteger_2), __typedArg0)
})())
return
}, 1)

tmp21713 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminteger_2, tmp21712)
}
__typedArg0 := syminteger_2
__typedArg1 := tmp21712
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21714 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21713)


_ = tmp21714

tmp21715 := MakeNative(func(__e *ControlFlow) {
Y1071 := __e.Get(1)
_ = Y1071
__e.TailApply(PrimFunc(syminternal), Y1071)
return
}, 1)

tmp21716 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminternal, tmp21715)
}
__typedArg0 := syminternal
__typedArg1 := tmp21715
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21717 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21716)


_ = tmp21717

tmp21718 := MakeNative(func(__e *ControlFlow) {
Y1072 := __e.Get(1)
_ = Y1072
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(Y1072)
}
__typedArg0 := Y1072
return Call(__e, PrimFunc(symintern), __typedArg0)
})())
return
}, 1)

tmp21719 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symintern, tmp21718)
}
__typedArg0 := symintern
__typedArg1 := tmp21718
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21720 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21719)


_ = tmp21720

tmp21721 := MakeNative(func(__e *ControlFlow) {
Y1073 := __e.Get(1)
_ = Y1073
__e.TailApply(PrimFunc(syminput), Y1073)
return
}, 1)

tmp21722 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminput, tmp21721)
}
__typedArg0 := syminput
__typedArg1 := tmp21721
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21723 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21722)


_ = tmp21723

tmp21724 := MakeNative(func(__e *ControlFlow) {
Y1074 := __e.Get(1)
_ = Y1074
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1075 := __e.Get(1)
_ = Y1075
__e.TailApply(PrimFunc(symshen_4input_1h_7), Y1074, Y1075)
return
}, 1))
return
}, 1)

tmp21725 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminput_7, tmp21724)
}
__typedArg0 := syminput_7
__typedArg1 := tmp21724
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21726 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21725)


_ = tmp21726

tmp21727 := MakeNative(func(__e *ControlFlow) {
Y1076 := __e.Get(1)
_ = Y1076
__e.TailApply(PrimFunc(syminclude_1all_1but), Y1076)
return
}, 1)

tmp21728 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminclude_1all_1but, tmp21727)
}
__typedArg0 := syminclude_1all_1but
__typedArg1 := tmp21727
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21729 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21728)


_ = tmp21729

tmp21730 := MakeNative(func(__e *ControlFlow) {
Y1077 := __e.Get(1)
_ = Y1077
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1078 := __e.Get(1)
_ = Y1078
__e.TailApply(PrimFunc(symintersection), Y1077, Y1078)
return
}, 1))
return
}, 1)

tmp21731 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symintersection, tmp21730)
}
__typedArg0 := symintersection
__typedArg1 := tmp21730
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21732 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21731)


_ = tmp21732

tmp21733 := MakeNative(func(__e *ControlFlow) {
Y1079 := __e.Get(1)
_ = Y1079
__e.TailApply(PrimFunc(syminternal), Y1079)
return
}, 1)

tmp21734 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminternal, tmp21733)
}
__typedArg0 := syminternal
__typedArg1 := tmp21733
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21735 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21734)


_ = tmp21735

tmp21736 := MakeNative(func(__e *ControlFlow) {
Y1080 := __e.Get(1)
_ = Y1080
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1081 := __e.Get(1)
_ = Y1081
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1082 := __e.Get(1)
_ = Y1082
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1083 := __e.Get(1)
_ = Y1083
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1084 := __e.Get(1)
_ = Y1084
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1085 := __e.Get(1)
_ = Y1085
__e.TailApply(PrimFunc(symis), Y1080, Y1081, Y1082, Y1083, Y1084, Y1085)
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21737 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symis, tmp21736)
}
__typedArg0 := symis
__typedArg1 := tmp21736
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21738 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21737)


_ = tmp21738

tmp21739 := MakeNative(func(__e *ControlFlow) {
Y1086 := __e.Get(1)
_ = Y1086
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1087 := __e.Get(1)
_ = Y1087
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1088 := __e.Get(1)
_ = Y1088
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1089 := __e.Get(1)
_ = Y1089
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1090 := __e.Get(1)
_ = Y1090
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1091 := __e.Get(1)
_ = Y1091
__e.TailApply(PrimFunc(symis_b), Y1086, Y1087, Y1088, Y1089, Y1090, Y1091)
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21740 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symis_b, tmp21739)
}
__typedArg0 := symis_b
__typedArg1 := tmp21739
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21741 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21740)


_ = tmp21741

tmp21742 := MakeNative(func(__e *ControlFlow) {
Y1092 := __e.Get(1)
_ = Y1092
__e.TailApply(PrimFunc(symlength), Y1092)
return
}, 1)

tmp21743 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlength, tmp21742)
}
__typedArg0 := symlength
__typedArg1 := tmp21742
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21744 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21743)


_ = tmp21744

tmp21745 := MakeNative(func(__e *ControlFlow) {
Y1093 := __e.Get(1)
_ = Y1093
__e.TailApply(PrimFunc(symlimit), Y1093)
return
}, 1)

tmp21746 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlimit, tmp21745)
}
__typedArg0 := symlimit
__typedArg1 := tmp21745
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21747 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21746)


_ = tmp21747

tmp21748 := MakeNative(func(__e *ControlFlow) {
Y1094 := __e.Get(1)
_ = Y1094
__e.TailApply(PrimFunc(symlineread), Y1094)
return
}, 1)

tmp21749 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlineread, tmp21748)
}
__typedArg0 := symlineread
__typedArg1 := tmp21748
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21750 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21749)


_ = tmp21750

tmp21751 := MakeNative(func(__e *ControlFlow) {
Y1095 := __e.Get(1)
_ = Y1095
__e.TailApply(PrimFunc(symload), Y1095)
return
}, 1)

tmp21752 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symload, tmp21751)
}
__typedArg0 := symload
__typedArg1 := tmp21751
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21753 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21752)


_ = tmp21753

tmp21754 := MakeNative(func(__e *ControlFlow) {
Y1096 := __e.Get(1)
_ = Y1096
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1097 := __e.Get(1)
_ = Y1097
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5) {
__typedN0, __typedOK0 := TypedFloat64(Y1096)
__typedN1, __typedOK1 := TypedFloat64(Y1097)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_5) {
return TypedMaterializeBoolean((__typedN0 < __typedN1))
}}
__typedArg0 := Y1096
__typedArg1 := Y1097
return Call(__e, PrimFunc(sym_5), __typedArg0, __typedArg1)
})())
return
}, 1))
return
}, 1)

tmp21755 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5, tmp21754)
}
__typedArg0 := sym_5
__typedArg1 := tmp21754
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21756 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21755)


_ = tmp21756

tmp21757 := MakeNative(func(__e *ControlFlow) {
Y1098 := __e.Get(1)
_ = Y1098
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1099 := __e.Get(1)
_ = Y1099
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_a) {
__typedN0, __typedOK0 := TypedFloat64(Y1098)
__typedN1, __typedOK1 := TypedFloat64(Y1099)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_5_a) {
return TypedMaterializeBoolean((__typedN0 <= __typedN1))
}}
__typedArg0 := Y1098
__typedArg1 := Y1099
return Call(__e, PrimFunc(sym_5_a), __typedArg0, __typedArg1)
})())
return
}, 1))
return
}, 1)

tmp21758 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_a, tmp21757)
}
__typedArg0 := sym_5_a
__typedArg1 := tmp21757
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21759 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21758)


_ = tmp21759

tmp21760 := MakeNative(func(__e *ControlFlow) {
Y1100 := __e.Get(1)
_ = Y1100
__e.TailApply(PrimFunc(symvector), Y1100)
return
}, 1)

tmp21761 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp21760)
}
__typedArg0 := symvector
__typedArg1 := tmp21760
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21762 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21761)


_ = tmp21762

tmp21763 := MakeNative(func(__e *ControlFlow) {
Y1101 := __e.Get(1)
_ = Y1101
__e.TailApply(PrimFunc(symmacroexpand), Y1101)
return
}, 1)

tmp21764 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symmacroexpand, tmp21763)
}
__typedArg0 := symmacroexpand
__typedArg1 := tmp21763
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21765 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21764)


_ = tmp21765

tmp21766 := MakeNative(func(__e *ControlFlow) {
Y1102 := __e.Get(1)
_ = Y1102
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1103 := __e.Get(1)
_ = Y1103
__e.TailApply(PrimFunc(symmap), Y1102, Y1103)
return
}, 1))
return
}, 1)

tmp21767 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symmap, tmp21766)
}
__typedArg0 := symmap
__typedArg1 := tmp21766
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21768 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21767)


_ = tmp21768

tmp21769 := MakeNative(func(__e *ControlFlow) {
Y1104 := __e.Get(1)
_ = Y1104
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1105 := __e.Get(1)
_ = Y1105
__e.TailApply(PrimFunc(symmapcan), Y1104, Y1105)
return
}, 1))
return
}, 1)

tmp21770 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symmapcan, tmp21769)
}
__typedArg0 := symmapcan
__typedArg1 := tmp21769
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21771 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21770)


_ = tmp21771

tmp21772 := MakeNative(func(__e *ControlFlow) {
Y1106 := __e.Get(1)
_ = Y1106
__e.TailApply(PrimFunc(symmaxinferences), Y1106)
return
}, 1)

tmp21773 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symmaxinferences, tmp21772)
}
__typedArg0 := symmaxinferences
__typedArg1 := tmp21772
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21774 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21773)


_ = tmp21774

tmp21775 := MakeNative(func(__e *ControlFlow) {
Y1107 := __e.Get(1)
_ = Y1107
__e.TailApply(PrimFunc(symnl), Y1107)
return
}, 1)

tmp21776 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnl, tmp21775)
}
__typedArg0 := symnl
__typedArg1 := tmp21775
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21777 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21776)


_ = tmp21777

tmp21778 := MakeNative(func(__e *ControlFlow) {
Y1108 := __e.Get(1)
_ = Y1108
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(Y1108)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := Y1108
return Call(__e, PrimFunc(symnot), __typedArg0)
})())
return
}, 1)

tmp21779 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnot, tmp21778)
}
__typedArg0 := symnot
__typedArg1 := tmp21778
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21780 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21779)


_ = tmp21780

tmp21781 := MakeNative(func(__e *ControlFlow) {
Y1109 := __e.Get(1)
_ = Y1109
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1110 := __e.Get(1)
_ = Y1110
__e.TailApply(PrimFunc(symnth), Y1109, Y1110)
return
}, 1))
return
}, 1)

tmp21782 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnth, tmp21781)
}
__typedArg0 := symnth
__typedArg1 := tmp21781
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21783 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21782)


_ = tmp21783

tmp21784 := MakeNative(func(__e *ControlFlow) {
Y1111 := __e.Get(1)
_ = Y1111
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symn_1_6string) {
return PrimNumberToString(Y1111)
}
__typedArg0 := Y1111
return Call(__e, PrimFunc(symn_1_6string), __typedArg0)
})())
return
}, 1)

tmp21785 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symn_1_6string, tmp21784)
}
__typedArg0 := symn_1_6string
__typedArg1 := tmp21784
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21786 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21785)


_ = tmp21786

tmp21787 := MakeNative(func(__e *ControlFlow) {
Y1112 := __e.Get(1)
_ = Y1112
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnumber_2) {
return PrimIsNumber(Y1112)
}
__typedArg0 := Y1112
return Call(__e, PrimFunc(symnumber_2), __typedArg0)
})())
return
}, 1)

tmp21788 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber_2, tmp21787)
}
__typedArg0 := symnumber_2
__typedArg1 := tmp21787
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21789 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21788)


_ = tmp21789

tmp21790 := MakeNative(func(__e *ControlFlow) {
Y1113 := __e.Get(1)
_ = Y1113
__e.TailApply(PrimFunc(symoccurs_1check), Y1113)
return
}, 1)

tmp21791 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symoccurs_1check, tmp21790)
}
__typedArg0 := symoccurs_1check
__typedArg1 := tmp21790
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21792 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21791)


_ = tmp21792

tmp21793 := MakeNative(func(__e *ControlFlow) {
Y1114 := __e.Get(1)
_ = Y1114
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1115 := __e.Get(1)
_ = Y1115
__e.TailApply(PrimFunc(symoccurrences), Y1114, Y1115)
return
}, 1))
return
}, 1)

tmp21794 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symoccurrences, tmp21793)
}
__typedArg0 := symoccurrences
__typedArg1 := tmp21793
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21795 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21794)


_ = tmp21795

tmp21796 := MakeNative(func(__e *ControlFlow) {
Y1116 := __e.Get(1)
_ = Y1116
__e.TailApply(PrimFunc(symoccurs_1check), Y1116)
return
}, 1)

tmp21797 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symoccurs_1check, tmp21796)
}
__typedArg0 := symoccurs_1check
__typedArg1 := tmp21796
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21798 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21797)


_ = tmp21798

tmp21799 := MakeNative(func(__e *ControlFlow) {
Y1117 := __e.Get(1)
_ = Y1117
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1118 := __e.Get(1)
_ = Y1118
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symopen) {
return PrimOpenStream(Y1117, Y1118)
}
__typedArg0 := Y1117
__typedArg1 := Y1118
return Call(__e, PrimFunc(symopen), __typedArg0, __typedArg1)
})())
return
}, 1))
return
}, 1)

tmp21800 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symopen, tmp21799)
}
__typedArg0 := symopen
__typedArg1 := tmp21799
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21801 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21800)


_ = tmp21801

tmp21802 := MakeNative(func(__e *ControlFlow) {
Y1119 := __e.Get(1)
_ = Y1119
__e.TailApply(PrimFunc(symoptimise), Y1119)
return
}, 1)

tmp21803 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symoptimise, tmp21802)
}
__typedArg0 := symoptimise
__typedArg1 := tmp21802
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21804 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21803)


_ = tmp21804

tmp21805 := MakeNative(func(__e *ControlFlow) {
Y1120 := __e.Get(1)
_ = Y1120
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1121 := __e.Get(1)
_ = Y1121
if True == Y1120 {
__e.Return(True)
return
} else {
if True == Y1121 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}
}
}, 1))
return
}, 1)

tmp21808 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symor, tmp21805)
}
__typedArg0 := symor
__typedArg1 := tmp21805
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21809 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21808)


_ = tmp21809

tmp21810 := MakeNative(func(__e *ControlFlow) {
Y1122 := __e.Get(1)
_ = Y1122
__e.TailApply(PrimFunc(sympackage_2), Y1122)
return
}, 1)

tmp21811 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympackage_2, tmp21810)
}
__typedArg0 := sympackage_2
__typedArg1 := tmp21810
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21812 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21811)


_ = tmp21812

tmp21813 := MakeNative(func(__e *ControlFlow) {
Y1123 := __e.Get(1)
_ = Y1123
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1124 := __e.Get(1)
_ = Y1124
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sympos) {
return PrimPos(Y1123, Y1124)
}
__typedArg0 := Y1123
__typedArg1 := Y1124
return Call(__e, PrimFunc(sympos), __typedArg0, __typedArg1)
})())
return
}, 1))
return
}, 1)

tmp21814 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympos, tmp21813)
}
__typedArg0 := sympos
__typedArg1 := tmp21813
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21815 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21814)


_ = tmp21815

tmp21816 := MakeNative(func(__e *ControlFlow) {
Y1125 := __e.Get(1)
_ = Y1125
__e.TailApply(PrimFunc(sympreclude_1all_1but), Y1125)
return
}, 1)

tmp21817 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympreclude_1all_1but, tmp21816)
}
__typedArg0 := sympreclude_1all_1but
__typedArg1 := tmp21816
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21818 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21817)


_ = tmp21818

tmp21819 := MakeNative(func(__e *ControlFlow) {
Y1126 := __e.Get(1)
_ = Y1126
__e.TailApply(PrimFunc(symprint), Y1126)
return
}, 1)

tmp21820 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symprint, tmp21819)
}
__typedArg0 := symprint
__typedArg1 := tmp21819
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21821 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21820)


_ = tmp21821

tmp21822 := MakeNative(func(__e *ControlFlow) {
Y1127 := __e.Get(1)
_ = Y1127
__e.TailApply(PrimFunc(symprofile), Y1127)
return
}, 1)

tmp21823 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symprofile, tmp21822)
}
__typedArg0 := symprofile
__typedArg1 := tmp21822
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21824 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21823)


_ = tmp21824

tmp21825 := MakeNative(func(__e *ControlFlow) {
Y1128 := __e.Get(1)
_ = Y1128
__e.TailApply(PrimFunc(symshen_4print_1prolog_1vector), Y1128)
return
}, 1)

tmp21826 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4print_1prolog_1vector, tmp21825)
}
__typedArg0 := symshen_4print_1prolog_1vector
__typedArg1 := tmp21825
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21827 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21826)


_ = tmp21827

tmp21828 := MakeNative(func(__e *ControlFlow) {
Y1129 := __e.Get(1)
_ = Y1129
__e.TailApply(PrimFunc(symshen_4print_1freshterm), Y1129)
return
}, 1)

tmp21829 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4print_1freshterm, tmp21828)
}
__typedArg0 := symshen_4print_1freshterm
__typedArg1 := tmp21828
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21830 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21829)


_ = tmp21830

tmp21831 := MakeNative(func(__e *ControlFlow) {
Y1130 := __e.Get(1)
_ = Y1130
__e.TailApply(PrimFunc(symshen_4printF), Y1130)
return
}, 1)

tmp21832 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4printF, tmp21831)
}
__typedArg0 := symshen_4printF
__typedArg1 := tmp21831
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21833 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21832)


_ = tmp21833

tmp21834 := MakeNative(func(__e *ControlFlow) {
Y1131 := __e.Get(1)
_ = Y1131
__e.TailApply(PrimFunc(symprolog_1memory), Y1131)
return
}, 1)

tmp21835 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symprolog_1memory, tmp21834)
}
__typedArg0 := symprolog_1memory
__typedArg1 := tmp21834
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21836 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21835)


_ = tmp21836

tmp21837 := MakeNative(func(__e *ControlFlow) {
Y1132 := __e.Get(1)
_ = Y1132
__e.TailApply(PrimFunc(symprofile_1results), Y1132)
return
}, 1)

tmp21838 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symprofile_1results, tmp21837)
}
__typedArg0 := symprofile_1results
__typedArg1 := tmp21837
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21839 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21838)


_ = tmp21839

tmp21840 := MakeNative(func(__e *ControlFlow) {
Y1133 := __e.Get(1)
_ = Y1133
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1134 := __e.Get(1)
_ = Y1134
__e.TailApply(PrimFunc(sympr), Y1133, Y1134)
return
}, 1))
return
}, 1)

tmp21841 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympr, tmp21840)
}
__typedArg0 := sympr
__typedArg1 := tmp21840
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21842 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21841)


_ = tmp21842

tmp21843 := MakeNative(func(__e *ControlFlow) {
Y1135 := __e.Get(1)
_ = Y1135
__e.TailApply(PrimFunc(symps), Y1135)
return
}, 1)

tmp21844 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symps, tmp21843)
}
__typedArg0 := symps
__typedArg1 := tmp21843
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21845 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21844)


_ = tmp21845

tmp21846 := MakeNative(func(__e *ControlFlow) {
Y1136 := __e.Get(1)
_ = Y1136
__e.TailApply(PrimFunc(sympreclude), Y1136)
return
}, 1)

tmp21847 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympreclude, tmp21846)
}
__typedArg0 := sympreclude
__typedArg1 := tmp21846
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21848 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21847)


_ = tmp21848

tmp21849 := MakeNative(func(__e *ControlFlow) {
Y1137 := __e.Get(1)
_ = Y1137
__e.TailApply(PrimFunc(sympreclude_1all_1but), Y1137)
return
}, 1)

tmp21850 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympreclude_1all_1but, tmp21849)
}
__typedArg0 := sympreclude_1all_1but
__typedArg1 := tmp21849
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21851 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21850)


_ = tmp21851

tmp21852 := MakeNative(func(__e *ControlFlow) {
Y1138 := __e.Get(1)
_ = Y1138
__e.TailApply(PrimFunc(symprotect), Y1138)
return
}, 1)

tmp21853 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symprotect, tmp21852)
}
__typedArg0 := symprotect
__typedArg1 := tmp21852
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21854 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21853)


_ = tmp21854

tmp21855 := MakeNative(func(__e *ControlFlow) {
Y1139 := __e.Get(1)
_ = Y1139
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1140 := __e.Get(1)
_ = Y1140
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1141 := __e.Get(1)
_ = Y1141
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1142 := __e.Get(1)
_ = Y1142
__e.TailApply(PrimFunc(symput), Y1139, Y1140, Y1141, Y1142)
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21856 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symput, tmp21855)
}
__typedArg0 := symput
__typedArg1 := tmp21855
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21857 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21856)


_ = tmp21857

tmp21858 := MakeNative(func(__e *ControlFlow) {
Y1143 := __e.Get(1)
_ = Y1143
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symread_1file_1as_1string) {
return PrimReadFileAsString(Y1143)
}
__typedArg0 := Y1143
return Call(__e, PrimFunc(symread_1file_1as_1string), __typedArg0)
})())
return
}, 1)

tmp21859 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1file_1as_1string, tmp21858)
}
__typedArg0 := symread_1file_1as_1string
__typedArg1 := tmp21858
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21860 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21859)


_ = tmp21860

tmp21861 := MakeNative(func(__e *ControlFlow) {
Y1144 := __e.Get(1)
_ = Y1144
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symread_1file_1as_1bytelist) {
return PrimReadFileAsByteList(Y1144)
}
__typedArg0 := Y1144
return Call(__e, PrimFunc(symread_1file_1as_1bytelist), __typedArg0)
})())
return
}, 1)

tmp21862 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1file_1as_1bytelist, tmp21861)
}
__typedArg0 := symread_1file_1as_1bytelist
__typedArg1 := tmp21861
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21863 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21862)


_ = tmp21863

tmp21864 := MakeNative(func(__e *ControlFlow) {
Y1145 := __e.Get(1)
_ = Y1145
__e.TailApply(PrimFunc(symread_1file), Y1145)
return
}, 1)

tmp21865 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1file, tmp21864)
}
__typedArg0 := symread_1file
__typedArg1 := tmp21864
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21866 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21865)


_ = tmp21866

tmp21867 := MakeNative(func(__e *ControlFlow) {
Y1146 := __e.Get(1)
_ = Y1146
__e.TailApply(PrimFunc(symread), Y1146)
return
}, 1)

tmp21868 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread, tmp21867)
}
__typedArg0 := symread
__typedArg1 := tmp21867
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21869 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21868)


_ = tmp21869

tmp21870 := MakeNative(func(__e *ControlFlow) {
Y1147 := __e.Get(1)
_ = Y1147
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symread_1byte) {
return PrimReadByte(Y1147)
}
__typedArg0 := Y1147
return Call(__e, PrimFunc(symread_1byte), __typedArg0)
})())
return
}, 1)

tmp21871 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1byte, tmp21870)
}
__typedArg0 := symread_1byte
__typedArg1 := tmp21870
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21872 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21871)


_ = tmp21872

tmp21873 := MakeNative(func(__e *ControlFlow) {
Y1148 := __e.Get(1)
_ = Y1148
__e.TailApply(PrimFunc(symread_1from_1string), Y1148)
return
}, 1)

tmp21874 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1from_1string, tmp21873)
}
__typedArg0 := symread_1from_1string
__typedArg1 := tmp21873
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21875 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21874)


_ = tmp21875

tmp21876 := MakeNative(func(__e *ControlFlow) {
Y1149 := __e.Get(1)
_ = Y1149
__e.TailApply(PrimFunc(symread_1from_1string_1unprocessed), Y1149)
return
}, 1)

tmp21877 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1from_1string_1unprocessed, tmp21876)
}
__typedArg0 := symread_1from_1string_1unprocessed
__typedArg1 := tmp21876
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21878 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21877)


_ = tmp21878

tmp21879 := MakeNative(func(__e *ControlFlow) {
Y1150 := __e.Get(1)
_ = Y1150
__e.TailApply(PrimFunc(symread_1unit_1string), Y1150)
return
}, 1)

tmp21880 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1unit_1string, tmp21879)
}
__typedArg0 := symread_1unit_1string
__typedArg1 := tmp21879
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21881 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21880)


_ = tmp21881

tmp21882 := MakeNative(func(__e *ControlFlow) {
Y1151 := __e.Get(1)
_ = Y1151
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1152 := __e.Get(1)
_ = Y1152
__e.TailApply(PrimFunc(symshen_4register_1source_1form), Y1151, Y1152)
return
}, 1))
return
}, 1)

tmp21883 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4register_1source_1form, tmp21882)
}
__typedArg0 := symshen_4register_1source_1form
__typedArg1 := tmp21882
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21884 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21883)


_ = tmp21884

tmp21885 := MakeNative(func(__e *ControlFlow) {
Y1153 := __e.Get(1)
_ = Y1153
__e.TailApply(PrimFunc(symshen_4unregister_1source_1form), Y1153)
return
}, 1)

tmp21886 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4unregister_1source_1form, tmp21885)
}
__typedArg0 := symshen_4unregister_1source_1form
__typedArg1 := tmp21885
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21887 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21886)


_ = tmp21887

tmp21888 := MakeNative(func(__e *ControlFlow) {
Y1154 := __e.Get(1)
_ = Y1154
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1155 := __e.Get(1)
_ = Y1155
__e.TailApply(PrimFunc(symremove), Y1154, Y1155)
return
}, 1))
return
}, 1)

tmp21889 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symremove, tmp21888)
}
__typedArg0 := symremove
__typedArg1 := tmp21888
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21890 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21889)


_ = tmp21890

tmp21891 := MakeNative(func(__e *ControlFlow) {
Y1156 := __e.Get(1)
_ = Y1156
__e.TailApply(PrimFunc(symreverse), Y1156)
return
}, 1)

tmp21892 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symreverse, tmp21891)
}
__typedArg0 := symreverse
__typedArg1 := tmp21891
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21893 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21892)


_ = tmp21893

tmp21894 := MakeNative(func(__e *ControlFlow) {
Y1157 := __e.Get(1)
_ = Y1157
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1158 := __e.Get(1)
_ = Y1158
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(Y1157, Y1158)
}
__typedArg0 := Y1157
__typedArg1 := Y1158
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return
}, 1))
return
}, 1)

tmp21895 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symset, tmp21894)
}
__typedArg0 := symset
__typedArg1 := tmp21894
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21896 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21895)


_ = tmp21896

tmp21897 := MakeNative(func(__e *ControlFlow) {
Y1159 := __e.Get(1)
_ = Y1159
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(Y1159)
}
__typedArg0 := Y1159
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}, 1)

tmp21898 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsimple_1error, tmp21897)
}
__typedArg0 := symsimple_1error
__typedArg1 := tmp21897
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21899 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21898)


_ = tmp21899

tmp21900 := MakeNative(func(__e *ControlFlow) {
Y1160 := __e.Get(1)
_ = Y1160
__e.TailApply(PrimFunc(symsnd), Y1160)
return
}, 1)

tmp21901 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsnd, tmp21900)
}
__typedArg0 := symsnd
__typedArg1 := tmp21900
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21902 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21901)


_ = tmp21902

tmp21903 := MakeNative(func(__e *ControlFlow) {
Y1161 := __e.Get(1)
_ = Y1161
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1162 := __e.Get(1)
_ = Y1162
__e.TailApply(PrimFunc(symspecialise), Y1161, Y1162)
return
}, 1))
return
}, 1)

tmp21904 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symspecialise, tmp21903)
}
__typedArg0 := symspecialise
__typedArg1 := tmp21903
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21905 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21904)


_ = tmp21905

tmp21906 := MakeNative(func(__e *ControlFlow) {
Y1163 := __e.Get(1)
_ = Y1163
__e.TailApply(PrimFunc(symspy), Y1163)
return
}, 1)

tmp21907 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symspy, tmp21906)
}
__typedArg0 := symspy
__typedArg1 := tmp21906
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21908 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21907)


_ = tmp21908

tmp21909 := MakeNative(func(__e *ControlFlow) {
Y1164 := __e.Get(1)
_ = Y1164
__e.TailApply(PrimFunc(symstep), Y1164)
return
}, 1)

tmp21910 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstep, tmp21909)
}
__typedArg0 := symstep
__typedArg1 := tmp21909
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21911 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21910)


_ = tmp21911

tmp21912 := MakeNative(func(__e *ControlFlow) {
Y1165 := __e.Get(1)
_ = Y1165
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(Y1165)
}
__typedArg0 := Y1165
return Call(__e, PrimFunc(symstr), __typedArg0)
})())
return
}, 1)

tmp21913 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstr, tmp21912)
}
__typedArg0 := symstr
__typedArg1 := tmp21912
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21914 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21913)


_ = tmp21914

tmp21915 := MakeNative(func(__e *ControlFlow) {
Y1166 := __e.Get(1)
_ = Y1166
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_1_6n) {
return PrimStringToNumber(Y1166)
}
__typedArg0 := Y1166
return Call(__e, PrimFunc(symstring_1_6n), __typedArg0)
})())
return
}, 1)

tmp21916 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring_1_6n, tmp21915)
}
__typedArg0 := symstring_1_6n
__typedArg1 := tmp21915
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21917 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21916)


_ = tmp21917

tmp21918 := MakeNative(func(__e *ControlFlow) {
Y1167 := __e.Get(1)
_ = Y1167
__e.TailApply(PrimFunc(symstring_1_6symbol), Y1167)
return
}, 1)

tmp21919 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring_1_6symbol, tmp21918)
}
__typedArg0 := symstring_1_6symbol
__typedArg1 := tmp21918
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21920 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21919)


_ = tmp21920

tmp21921 := MakeNative(func(__e *ControlFlow) {
Y1168 := __e.Get(1)
_ = Y1168
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_2) {
return PrimIsString(Y1168)
}
__typedArg0 := Y1168
return Call(__e, PrimFunc(symstring_2), __typedArg0)
})())
return
}, 1)

tmp21922 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring_2, tmp21921)
}
__typedArg0 := symstring_2
__typedArg1 := tmp21921
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21923 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21922)


_ = tmp21923

tmp21924 := MakeNative(func(__e *ControlFlow) {
Y1169 := __e.Get(1)
_ = Y1169
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1170 := __e.Get(1)
_ = Y1170
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1171 := __e.Get(1)
_ = Y1171
__e.TailApply(PrimFunc(symsubst), Y1169, Y1170, Y1171)
return
}, 1))
return
}, 1))
return
}, 1)

tmp21925 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsubst, tmp21924)
}
__typedArg0 := symsubst
__typedArg1 := tmp21924
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21926 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21925)


_ = tmp21926

tmp21927 := MakeNative(func(__e *ControlFlow) {
Y1172 := __e.Get(1)
_ = Y1172
__e.TailApply(PrimFunc(symsum), Y1172)
return
}, 1)

tmp21928 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsum, tmp21927)
}
__typedArg0 := symsum
__typedArg1 := tmp21927
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21929 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21928)


_ = tmp21929

tmp21930 := MakeNative(func(__e *ControlFlow) {
Y1173 := __e.Get(1)
_ = Y1173
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsymbol_2) {
return PrimIsSymbol(Y1173)
}
__typedArg0 := Y1173
return Call(__e, PrimFunc(symsymbol_2), __typedArg0)
})())
return
}, 1)

tmp21931 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol_2, tmp21930)
}
__typedArg0 := symsymbol_2
__typedArg1 := tmp21930
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21932 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21931)


_ = tmp21932

tmp21933 := MakeNative(func(__e *ControlFlow) {
Y1174 := __e.Get(1)
_ = Y1174
__e.TailApply(PrimFunc(symsystemf), Y1174)
return
}, 1)

tmp21934 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsystemf, tmp21933)
}
__typedArg0 := symsystemf
__typedArg1 := tmp21933
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21935 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21934)


_ = tmp21935

tmp21936 := MakeNative(func(__e *ControlFlow) {
Y1175 := __e.Get(1)
_ = Y1175
__e.TailApply(PrimFunc(symtail), Y1175)
return
}, 1)

tmp21937 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtail, tmp21936)
}
__typedArg0 := symtail
__typedArg1 := tmp21936
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21938 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21937)


_ = tmp21938

tmp21939 := MakeNative(func(__e *ControlFlow) {
Y1176 := __e.Get(1)
_ = Y1176
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Y1176)
}
__typedArg0 := Y1176
return Call(__e, PrimFunc(symtl), __typedArg0)
})())
return
}, 1)

tmp21940 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtl, tmp21939)
}
__typedArg0 := symtl
__typedArg1 := tmp21939
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21941 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21940)


_ = tmp21941

tmp21942 := MakeNative(func(__e *ControlFlow) {
Y1177 := __e.Get(1)
_ = Y1177
__e.TailApply(PrimFunc(symtc), Y1177)
return
}, 1)

tmp21943 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtc, tmp21942)
}
__typedArg0 := symtc
__typedArg1 := tmp21942
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21944 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21943)


_ = tmp21944

tmp21945 := MakeNative(func(__e *ControlFlow) {
Y1178 := __e.Get(1)
_ = Y1178
__e.TailApply(PrimFunc(symthaw), Y1178)
return
}, 1)

tmp21946 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symthaw, tmp21945)
}
__typedArg0 := symthaw
__typedArg1 := tmp21945
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21947 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21946)


_ = tmp21947

tmp21948 := MakeNative(func(__e *ControlFlow) {
Y1179 := __e.Get(1)
_ = Y1179
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(Y1179)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := Y1179
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())
return
}, 1)

tmp21949 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtlstr, tmp21948)
}
__typedArg0 := symtlstr
__typedArg1 := tmp21948
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21950 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21949)


_ = tmp21950

tmp21951 := MakeNative(func(__e *ControlFlow) {
Y1180 := __e.Get(1)
_ = Y1180
__e.TailApply(PrimFunc(symtrack), Y1180)
return
}, 1)

tmp21952 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtrack, tmp21951)
}
__typedArg0 := symtrack
__typedArg1 := tmp21951
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21953 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21952)


_ = tmp21953

tmp21954 := MakeNative(func(__e *ControlFlow) {
Y1181 := __e.Get(1)
_ = Y1181
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1182 := __e.Get(1)
_ = Y1182
tmp21955 := MakeNative(func(__e *ControlFlow) {
__e.Return(Y1181)
return
}, 0)

__e.TailApply(try_1catch, tmp21955, Y1182)
return


}, 1))
return
}, 1)

tmp21956 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtrap_1error, tmp21954)
}
__typedArg0 := symtrap_1error
__typedArg1 := tmp21954
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21957 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21956)


_ = tmp21957

tmp21958 := MakeNative(func(__e *ControlFlow) {
Y1183 := __e.Get(1)
_ = Y1183
__e.TailApply(PrimFunc(symtuple_2), Y1183)
return
}, 1)

tmp21959 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtuple_2, tmp21958)
}
__typedArg0 := symtuple_2
__typedArg1 := tmp21958
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21960 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21959)


_ = tmp21960

tmp21961 := MakeNative(func(__e *ControlFlow) {
Y1184 := __e.Get(1)
_ = Y1184
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1185 := __e.Get(1)
_ = Y1185
__e.Return(Y1185)
return
}, 1))
return
}, 1)

tmp21962 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtype, tmp21961)
}
__typedArg0 := symtype
__typedArg1 := tmp21961
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21963 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21962)


_ = tmp21963

tmp21964 := MakeNative(func(__e *ControlFlow) {
Y1186 := __e.Get(1)
_ = Y1186
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1187 := __e.Get(1)
_ = Y1187
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1188 := __e.Get(1)
_ = Y1188
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1189 := __e.Get(1)
_ = Y1189
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1190 := __e.Get(1)
_ = Y1190
__e.TailApply(PrimFunc(symreturn), Y1186, Y1187, Y1188, Y1189, Y1190)
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21965 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symreturn, tmp21964)
}
__typedArg0 := symreturn
__typedArg1 := tmp21964
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21966 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21965)


_ = tmp21966

tmp21967 := MakeNative(func(__e *ControlFlow) {
Y1191 := __e.Get(1)
_ = Y1191
__e.TailApply(PrimFunc(symunabsolute), Y1191)
return
}, 1)

tmp21968 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunabsolute, tmp21967)
}
__typedArg0 := symunabsolute
__typedArg1 := tmp21967
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21969 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21968)


_ = tmp21969

tmp21970 := MakeNative(func(__e *ControlFlow) {
Y1192 := __e.Get(1)
_ = Y1192
__e.TailApply(PrimFunc(symundefmacro), Y1192)
return
}, 1)

tmp21971 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symundefmacro, tmp21970)
}
__typedArg0 := symundefmacro
__typedArg1 := tmp21970
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21972 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21971)


_ = tmp21972

tmp21973 := MakeNative(func(__e *ControlFlow) {
Y1193 := __e.Get(1)
_ = Y1193
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1194 := __e.Get(1)
_ = Y1194
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1195 := __e.Get(1)
_ = Y1195
__e.TailApply(PrimFunc(symunput), Y1193, Y1194, Y1195)
return
}, 1))
return
}, 1))
return
}, 1)

tmp21974 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunput, tmp21973)
}
__typedArg0 := symunput
__typedArg1 := tmp21973
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21975 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21974)


_ = tmp21975

tmp21976 := MakeNative(func(__e *ControlFlow) {
Y1196 := __e.Get(1)
_ = Y1196
__e.TailApply(PrimFunc(symunprofile), Y1196)
return
}, 1)

tmp21977 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunprofile, tmp21976)
}
__typedArg0 := symunprofile
__typedArg1 := tmp21976
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21978 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21977)


_ = tmp21978

tmp21979 := MakeNative(func(__e *ControlFlow) {
Y1197 := __e.Get(1)
_ = Y1197
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1198 := __e.Get(1)
_ = Y1198
__e.TailApply(PrimFunc(symunion), Y1197, Y1198)
return
}, 1))
return
}, 1)

tmp21980 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunion, tmp21979)
}
__typedArg0 := symunion
__typedArg1 := tmp21979
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21981 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21980)


_ = tmp21981

tmp21982 := MakeNative(func(__e *ControlFlow) {
Y1199 := __e.Get(1)
_ = Y1199
__e.TailApply(PrimFunc(symuntrack), Y1199)
return
}, 1)

tmp21983 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symuntrack, tmp21982)
}
__typedArg0 := symuntrack
__typedArg1 := tmp21982
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21984 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21983)


_ = tmp21984

tmp21985 := MakeNative(func(__e *ControlFlow) {
Y1200 := __e.Get(1)
_ = Y1200
__e.TailApply(PrimFunc(symundefmacro), Y1200)
return
}, 1)

tmp21986 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symundefmacro, tmp21985)
}
__typedArg0 := symundefmacro
__typedArg1 := tmp21985
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21987 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21986)


_ = tmp21987

tmp21988 := MakeNative(func(__e *ControlFlow) {
Y1201 := __e.Get(1)
_ = Y1201
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1202 := __e.Get(1)
_ = Y1202
__e.TailApply(PrimFunc(symupdate_1lambda_1table), Y1201, Y1202)
return
}, 1))
return
}, 1)

tmp21989 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symupdate_1lambda_1table, tmp21988)
}
__typedArg0 := symupdate_1lambda_1table
__typedArg1 := tmp21988
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21990 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21989)


_ = tmp21990

tmp21991 := MakeNative(func(__e *ControlFlow) {
Y1203 := __e.Get(1)
_ = Y1203
__e.TailApply(PrimFunc(symvector), Y1203)
return
}, 1)

tmp21992 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp21991)
}
__typedArg0 := symvector
__typedArg1 := tmp21991
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21993 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21992)


_ = tmp21993

tmp21994 := MakeNative(func(__e *ControlFlow) {
Y1204 := __e.Get(1)
_ = Y1204
__e.TailApply(PrimFunc(symvector_2), Y1204)
return
}, 1)

tmp21995 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector_2, tmp21994)
}
__typedArg0 := symvector_2
__typedArg1 := tmp21994
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21996 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21995)


_ = tmp21996

tmp21997 := MakeNative(func(__e *ControlFlow) {
Y1205 := __e.Get(1)
_ = Y1205
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1206 := __e.Get(1)
_ = Y1206
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1207 := __e.Get(1)
_ = Y1207
__e.TailApply(PrimFunc(symvector_1_6), Y1205, Y1206, Y1207)
return
}, 1))
return
}, 1))
return
}, 1)

tmp21998 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector_1_6, tmp21997)
}
__typedArg0 := symvector_1_6
__typedArg1 := tmp21997
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21999 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21998)


_ = tmp21999

tmp22000 := MakeNative(func(__e *ControlFlow) {
Y1208 := __e.Get(1)
_ = Y1208
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(Y1208)
}
__typedArg0 := Y1208
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 1)

tmp22001 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvalue, tmp22000)
}
__typedArg0 := symvalue
__typedArg1 := tmp22000
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp22002 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp22001)


_ = tmp22002

tmp22003 := MakeNative(func(__e *ControlFlow) {
Y1209 := __e.Get(1)
_ = Y1209
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(Y1209)
}
__typedArg0 := Y1209
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})())
return
}, 1)

tmp22004 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvariable_2, tmp22003)
}
__typedArg0 := symvariable_2
__typedArg1 := tmp22003
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp22005 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp22004)


_ = tmp22005

tmp22006 := MakeNative(func(__e *ControlFlow) {
Y1210 := __e.Get(1)
_ = Y1210
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1211 := __e.Get(1)
_ = Y1211
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1212 := __e.Get(1)
_ = Y1212
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1213 := __e.Get(1)
_ = Y1213
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1214 := __e.Get(1)
_ = Y1214
__e.TailApply(PrimFunc(symvar_2), Y1210, Y1211, Y1212, Y1213, Y1214)
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp22007 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvar_2, tmp22006)
}
__typedArg0 := symvar_2
__typedArg1 := tmp22006
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp22008 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp22007)


_ = tmp22008

tmp22009 := MakeNative(func(__e *ControlFlow) {
Y1215 := __e.Get(1)
_ = Y1215
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1216 := __e.Get(1)
_ = Y1216
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1217 := __e.Get(1)
_ = Y1217
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1218 := __e.Get(1)
_ = Y1218
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1219 := __e.Get(1)
_ = Y1219
__e.TailApply(PrimFunc(symwhen), Y1215, Y1216, Y1217, Y1218, Y1219)
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp22010 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhen, tmp22009)
}
__typedArg0 := symwhen
__typedArg1 := tmp22009
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp22011 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp22010)


_ = tmp22011

tmp22012 := MakeNative(func(__e *ControlFlow) {
Y1220 := __e.Get(1)
_ = Y1220
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1221 := __e.Get(1)
_ = Y1221
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symwrite_1byte) {
return PrimWriteByte(Y1220, Y1221)
}
__typedArg0 := Y1220
__typedArg1 := Y1221
return Call(__e, PrimFunc(symwrite_1byte), __typedArg0, __typedArg1)
})())
return
}, 1))
return
}, 1)

tmp22013 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwrite_1byte, tmp22012)
}
__typedArg0 := symwrite_1byte
__typedArg1 := tmp22012
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp22014 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp22013)


_ = tmp22014

tmp22015 := MakeNative(func(__e *ControlFlow) {
Y1222 := __e.Get(1)
_ = Y1222
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1223 := __e.Get(1)
_ = Y1223
__e.TailApply(PrimFunc(symwrite_1string), Y1222, Y1223)
return
}, 1))
return
}, 1)

tmp22016 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwrite_1string, tmp22015)
}
__typedArg0 := symwrite_1string
__typedArg1 := tmp22015
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp22017 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp22016)


_ = tmp22017

tmp22018 := MakeNative(func(__e *ControlFlow) {
Y1224 := __e.Get(1)
_ = Y1224
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1225 := __e.Get(1)
_ = Y1225
__e.TailApply(PrimFunc(symwrite_1to_1file), Y1224, Y1225)
return
}, 1))
return
}, 1)

tmp22019 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwrite_1to_1file, tmp22018)
}
__typedArg0 := symwrite_1to_1file
__typedArg1 := tmp22018
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp22020 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp22019)


_ = tmp22020

tmp22021 := MakeNative(func(__e *ControlFlow) {
Y1226 := __e.Get(1)
_ = Y1226
__e.TailApply(PrimFunc(symy_1or_1n_2), Y1226)
return
}, 1)

tmp22022 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symy_1or_1n_2, tmp22021)
}
__typedArg0 := symy_1or_1n_2
__typedArg1 := tmp22021
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp22023 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp22022)


_ = tmp22023

tmp22024 := MakeNative(func(__e *ControlFlow) {
Y1227 := __e.Get(1)
_ = Y1227
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1228 := __e.Get(1)
_ = Y1228
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(Y1227)
__typedN1, __typedOK1 := TypedFloat64(Y1228)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := Y1227
__typedArg1 := Y1228
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())
return
}, 1))
return
}, 1)

tmp22025 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_7, tmp22024)
}
__typedArg0 := sym_7
__typedArg1 := tmp22024
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp22026 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp22025)


_ = tmp22026

tmp22027 := MakeNative(func(__e *ControlFlow) {
Y1229 := __e.Get(1)
_ = Y1229
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1230 := __e.Get(1)
_ = Y1230
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_d) {
__typedN0, __typedOK0 := TypedFloat64(Y1229)
__typedN1, __typedOK1 := TypedFloat64(Y1230)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_d) {
return TypedMaterializeNumber((__typedN0 * __typedN1))
}}
__typedArg0 := Y1229
__typedArg1 := Y1230
return Call(__e, PrimFunc(sym_d), __typedArg0, __typedArg1)
})())
return
}, 1))
return
}, 1)

tmp22028 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_d, tmp22027)
}
__typedArg0 := sym_d
__typedArg1 := tmp22027
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp22029 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp22028)


_ = tmp22029

tmp22030 := MakeNative(func(__e *ControlFlow) {
Y1231 := __e.Get(1)
_ = Y1231
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1232 := __e.Get(1)
_ = Y1232
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_c) {
__typedN0, __typedOK0 := TypedFloat64(Y1231)
__typedN1, __typedOK1 := TypedFloat64(Y1232)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_c) {
return TypedMaterializeNumber(TypedDivideValue(__typedN0, __typedN1))
}}
__typedArg0 := Y1231
__typedArg1 := Y1232
return Call(__e, PrimFunc(sym_c), __typedArg0, __typedArg1)
})())
return
}, 1))
return
}, 1)

tmp22031 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_c, tmp22030)
}
__typedArg0 := sym_c
__typedArg1 := tmp22030
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp22032 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp22031)


_ = tmp22032

tmp22033 := MakeNative(func(__e *ControlFlow) {
Y1233 := __e.Get(1)
_ = Y1233
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1234 := __e.Get(1)
_ = Y1234
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(Y1233)
__typedN1, __typedOK1 := TypedFloat64(Y1234)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := Y1233
__typedArg1 := Y1234
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})())
return
}, 1))
return
}, 1)

tmp22034 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1, tmp22033)
}
__typedArg0 := sym_1
__typedArg1 := tmp22033
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp22035 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp22034)


_ = tmp22035

tmp22036 := MakeNative(func(__e *ControlFlow) {
Y1235 := __e.Get(1)
_ = Y1235
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1236 := __e.Get(1)
_ = Y1236
__e.TailApply(PrimFunc(sym_a_a), Y1235, Y1236)
return
}, 1))
return
}, 1)

tmp22037 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_a_a, tmp22036)
}
__typedArg0 := sym_a_a
__typedArg1 := tmp22036
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp22038 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp22037)


_ = tmp22038

tmp22039 := MakeNative(func(__e *ControlFlow) {
Y1237 := __e.Get(1)
_ = Y1237
__e.TailApply(PrimFunc(sym_5e_6), Y1237)
return
}, 1)

tmp22040 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5e_6, tmp22039)
}
__typedArg0 := sym_5e_6
__typedArg1 := tmp22039
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp22041 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp22040)


_ = tmp22041

tmp22042 := MakeNative(func(__e *ControlFlow) {
Y1238 := __e.Get(1)
_ = Y1238
__e.TailApply(PrimFunc(sym_5end_6), Y1238)
return
}, 1)

tmp22043 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5end_6, tmp22042)
}
__typedArg0 := sym_5end_6
__typedArg1 := tmp22042
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp22044 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp22043)


_ = tmp22044

tmp22045 := MakeNative(func(__e *ControlFlow) {
Y1239 := __e.Get(1)
_ = Y1239
__e.TailApply(PrimFunc(sym_5_b_6), Y1239)
return
}, 1)

tmp22046 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_b_6, tmp22045)
}
__typedArg0 := sym_5_b_6
__typedArg1 := tmp22045
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp22047 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp22046)


_ = tmp22047

tmp22048 := MakeNative(func(__e *ControlFlow) {
Y1240 := __e.Get(1)
_ = Y1240
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1241 := __e.Get(1)
_ = Y1241
__e.TailApply(PrimFunc(sym_8p), Y1240, Y1241)
return
}, 1))
return
}, 1)

tmp22049 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8p, tmp22048)
}
__typedArg0 := sym_8p
__typedArg1 := tmp22048
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp22050 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp22049)


_ = tmp22050

tmp22051 := MakeNative(func(__e *ControlFlow) {
Y1242 := __e.Get(1)
_ = Y1242
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1243 := __e.Get(1)
_ = Y1243
__e.TailApply(PrimFunc(sym_8v), Y1242, Y1243)
return
}, 1))
return
}, 1)

tmp22052 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8v, tmp22051)
}
__typedArg0 := sym_8v
__typedArg1 := tmp22051
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp22053 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp22052)


_ = tmp22053

tmp22054 := MakeNative(func(__e *ControlFlow) {
Y1244 := __e.Get(1)
_ = Y1244
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1245 := __e.Get(1)
_ = Y1245
__e.TailApply(PrimFunc(sym_8s), Y1244, Y1245)
return
}, 1))
return
}, 1)

tmp22055 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8s, tmp22054)
}
__typedArg0 := sym_8s
__typedArg1 := tmp22054
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4set_1lambda_1form_1entry), tmp22055)
return


}, 0)

tmp22056 := Call(__e, ns2_1set, symshen_4initialise_1lambda_1forms, tmp21501)


_ = tmp22056

tmp22057 := MakeNative(func(__e *ControlFlow) {
tmp22058 := Call(__e, PrimFunc(symshen_4initialise_1environment))


_ = tmp22058

tmp22059 := Call(__e, PrimFunc(symshen_4initialise_1lambda_1forms))


_ = tmp22059

__e.TailApply(PrimFunc(symshen_4initialise_1signedfuncs))
return


}, 0)

__e.TailApply(ns2_1set, symshen_4initialise, tmp22057)
return




}, 0)

