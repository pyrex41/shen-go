package main

import . "github.com/pyrex41/shen-go/kl"

var InitMain = MakeNative(func(__e *ControlFlow) {
tmp18983 := MakeNative(func(__e *ControlFlow) {
tmp18984 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dhistory_d, Nil)
}
__typedArg0 := symshen_4_dhistory_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp18984

tmp18985 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dtc_d, False)
}
__typedArg0 := symshen_4_dtc_d
__typedArg1 := False
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp18985

tmp18986 := Call(__e, PrimFunc(symshen_4dict), MakeInteger(20000))


tmp18987 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(sym_dproperty_1vector_d, tmp18986)
}
__typedArg0 := sym_dproperty_1vector_d
__typedArg1 := tmp18986
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp18987

tmp18988 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
__e.TailApply(PrimFunc(symshen_4macros), X)
return
}, 1)

tmp18989 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4macros, tmp18988)
}
__typedArg0 := symshen_4macros
__typedArg1 := tmp18988
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18990 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18989, Nil)
}
__typedArg0 := tmp18989
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18991 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(sym_dmacros_d, tmp18990)
}
__typedArg0 := sym_dmacros_d
__typedArg1 := tmp18990
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp18991

tmp18992 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsource_1form_1handlers_d, Nil)
}
__typedArg0 := symshen_4_dsource_1form_1handlers_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp18992

tmp18993 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dgensym_d, MakeInteger(0))
}
__typedArg0 := symshen_4_dgensym_d
__typedArg1 := MakeInteger(0)
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp18993

tmp18994 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dtracking_d, Nil)
}
__typedArg0 := symshen_4_dtracking_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp18994

tmp18995 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dprofiled_d, Nil)
}
__typedArg0 := symshen_4_dprofiled_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp18995

tmp18996 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtype, Nil)
}
__typedArg0 := symtype
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18997 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4input_1h_7, tmp18996)
}
__typedArg0 := symshen_4input_1h_7
__typedArg1 := tmp18996
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18998 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symopen, tmp18997)
}
__typedArg0 := symopen
__typedArg1 := tmp18997
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18999 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symset, tmp18998)
}
__typedArg0 := symset
__typedArg1 := tmp18998
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19000 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhere, tmp18999)
}
__typedArg0 := symwhere
__typedArg1 := tmp18999
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19001 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp19000)
}
__typedArg0 := symlet
__typedArg1 := tmp19000
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19002 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp19001)
}
__typedArg0 := symlambda
__typedArg1 := tmp19001
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19003 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp19002)
}
__typedArg0 := symcons
__typedArg1 := tmp19002
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19004 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8v, tmp19003)
}
__typedArg0 := sym_8v
__typedArg1 := tmp19003
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19005 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8s, tmp19004)
}
__typedArg0 := sym_8s
__typedArg1 := tmp19004
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19006 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8p, tmp19005)
}
__typedArg0 := sym_8p
__typedArg1 := tmp19005
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19007 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dspecial_d, tmp19006)
}
__typedArg0 := symshen_4_dspecial_d
__typedArg1 := tmp19006
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19007

tmp19008 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dextraspecial_d, Nil)
}
__typedArg0 := symshen_4_dextraspecial_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19008

tmp19009 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dspy_d, False)
}
__typedArg0 := symshen_4_dspy_d
__typedArg1 := False
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19009

tmp19010 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_ddatatypes_d, Nil)
}
__typedArg0 := symshen_4_ddatatypes_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19010

tmp19011 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dalldatatypes_d, Nil)
}
__typedArg0 := symshen_4_dalldatatypes_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19011

tmp19012 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dshen_1type_1theory_1enabled_2_d, True)
}
__typedArg0 := symshen_4_dshen_1type_1theory_1enabled_2_d
__typedArg1 := True
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19012

tmp19013 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dpackage_d, symnull)
}
__typedArg0 := symshen_4_dpackage_d
__typedArg1 := symnull
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19013

tmp19014 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsynonyms_d, Nil)
}
__typedArg0 := symshen_4_dsynonyms_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19014

tmp19015 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_doccurs_d, True)
}
__typedArg0 := symshen_4_doccurs_d
__typedArg1 := True
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19015

tmp19016 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dfactorise_2_d, False)
}
__typedArg0 := symshen_4_dfactorise_2_d
__typedArg1 := False
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19016

tmp19017 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dmaxinferences_d, MakeInteger(1000000))
}
__typedArg0 := symshen_4_dmaxinferences_d
__typedArg1 := MakeInteger(1000000)
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19017

tmp19018 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(sym_dmaximum_1print_1sequence_1size_d, MakeInteger(20))
}
__typedArg0 := sym_dmaximum_1print_1sequence_1size_d
__typedArg1 := MakeInteger(20)
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19018

tmp19019 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dcall_d, MakeInteger(0))
}
__typedArg0 := symshen_4_dcall_d
__typedArg1 := MakeInteger(0)
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19019

tmp19020 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dinfs_d, MakeInteger(0))
}
__typedArg0 := symshen_4_dinfs_d
__typedArg1 := MakeInteger(0)
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19020

tmp19021 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(sym_dhush_d, False)
}
__typedArg0 := sym_dhush_d
__typedArg1 := False
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19021

tmp19022 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_doptimise_d, False)
}
__typedArg0 := symshen_4_doptimise_d
__typedArg1 := False
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19022

tmp19023 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(sym_dversion_d, MakeString("43.1"))
}
__typedArg0 := sym_dversion_d
__typedArg1 := MakeString("43.1")
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19023

tmp19024 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dnames_d, Nil)
}
__typedArg0 := symshen_4_dnames_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19024

tmp19025 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dstep_d, False)
}
__typedArg0 := symshen_4_dstep_d
__typedArg1 := False
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19025

tmp19026 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dit_d, MakeString(""))
}
__typedArg0 := symshen_4_dit_d
__typedArg1 := MakeString("")
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19026

tmp19027 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dresidue_d, Nil)
}
__typedArg0 := symshen_4_dresidue_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19027

tmp19028 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(sym_dabsolute_d, Nil)
}
__typedArg0 := sym_dabsolute_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19028

tmp19029 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dprolog_1memory_d, MakeInteger(1000))
}
__typedArg0 := symshen_4_dprolog_1memory_d
__typedArg1 := MakeInteger(1000)
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19029

tmp19030 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dloading_2_d, False)
}
__typedArg0 := symshen_4_dloading_2_d
__typedArg1 := False
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19030

tmp19031 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_duserdefs_d, Nil)
}
__typedArg0 := symshen_4_duserdefs_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19031

tmp19032 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
__e.Return(X)
return
}, 1)

tmp19033 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_ddemodulation_1function_d, tmp19032)
}
__typedArg0 := symshen_4_ddemodulation_1function_d
__typedArg1 := tmp19032
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19033

tmp19034 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dcustom_1pattern_1compiler_d, False)
}
__typedArg0 := symshen_4_dcustom_1pattern_1compiler_d
__typedArg1 := False
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19034

tmp19035 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dcustom_1pattern_1reducer_d, False)
}
__typedArg0 := symshen_4_dcustom_1pattern_1reducer_d
__typedArg1 := False
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19035

tmp19038 := Call(__e, PrimFunc(symbound_2), sym_dhome_1directory_d)


tmp19039 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp19038)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp19038
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres19036 Obj

if True == tmp19039 {
tmp19037 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(sym_dhome_1directory_d, MakeString(""))
}
__typedArg0 := sym_dhome_1directory_d
__typedArg1 := MakeString("")
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

ifres19036 = tmp19037


} else {
ifres19036 = symshen_4skip


}

_ = ifres19036

tmp19043 := Call(__e, PrimFunc(symbound_2), sym_dsterror_d)


tmp19044 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp19043)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp19043
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres19040 Obj

if True == tmp19044 {
tmp19041 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dstoutput_d)
}
__typedArg0 := sym_dstoutput_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp19042 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(sym_dsterror_d, tmp19041)
}
__typedArg0 := sym_dsterror_d
__typedArg1 := tmp19041
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

ifres19040 = tmp19042


} else {
ifres19040 = symshen_4skip


}

_ = ifres19040

tmp19045 := Call(__e, PrimFunc(symprolog_1memory), MakeInteger(10000))


_ = tmp19045

tmp19046 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dloading_2_d, False)
}
__typedArg0 := symshen_4_dloading_2_d
__typedArg1 := False
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19046

tmp19047 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), Nil)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19048 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8s, tmp19047)
}
__typedArg0 := sym_8s
__typedArg1 := tmp19047
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19049 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19048)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19048
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19050 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8v, tmp19049)
}
__typedArg0 := sym_8v
__typedArg1 := tmp19049
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19051 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19050)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19050
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19052 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8p, tmp19051)
}
__typedArg0 := sym_8p
__typedArg1 := tmp19051
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19053 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19052)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19052
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19054 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_b_6, tmp19053)
}
__typedArg0 := sym_5_b_6
__typedArg1 := tmp19053
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19055 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19054)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19054
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19056 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5end_6, tmp19055)
}
__typedArg0 := sym_5end_6
__typedArg1 := tmp19055
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19057 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19056)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19056
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19058 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5e_6, tmp19057)
}
__typedArg0 := sym_5e_6
__typedArg1 := tmp19057
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19059 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19058)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19058
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19060 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_a_a, tmp19059)
}
__typedArg0 := sym_a_a
__typedArg1 := tmp19059
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19061 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19060)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19060
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19062 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1, tmp19061)
}
__typedArg0 := sym_1
__typedArg1 := tmp19061
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19063 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19062)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19062
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19064 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_c, tmp19063)
}
__typedArg0 := sym_c
__typedArg1 := tmp19063
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19065 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19064)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19064
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19066 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_d, tmp19065)
}
__typedArg0 := sym_d
__typedArg1 := tmp19065
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19067 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19066)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19066
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19068 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_7, tmp19067)
}
__typedArg0 := sym_7
__typedArg1 := tmp19067
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19069 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19068)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19068
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19070 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symy_1or_1n_2, tmp19069)
}
__typedArg0 := symy_1or_1n_2
__typedArg1 := tmp19069
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19071 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19070)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19070
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19072 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwrite_1to_1file, tmp19071)
}
__typedArg0 := symwrite_1to_1file
__typedArg1 := tmp19071
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19073 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19072)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19072
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19074 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwrite_1string, tmp19073)
}
__typedArg0 := symwrite_1string
__typedArg1 := tmp19073
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19075 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19074)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19074
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19076 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwrite_1byte, tmp19075)
}
__typedArg0 := symwrite_1byte
__typedArg1 := tmp19075
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19077 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(5), tmp19076)
}
__typedArg0 := MakeInteger(5)
__typedArg1 := tmp19076
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19078 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhen, tmp19077)
}
__typedArg0 := symwhen
__typedArg1 := tmp19077
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19079 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19078)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19078
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19080 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symversion, tmp19079)
}
__typedArg0 := symversion
__typedArg1 := tmp19079
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19081 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(5), tmp19080)
}
__typedArg0 := MakeInteger(5)
__typedArg1 := tmp19080
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19082 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvar_2, tmp19081)
}
__typedArg0 := symvar_2
__typedArg1 := tmp19081
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19083 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19082)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19082
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19084 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvariable_2, tmp19083)
}
__typedArg0 := symvariable_2
__typedArg1 := tmp19083
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19085 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19084)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19084
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19086 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvalue, tmp19085)
}
__typedArg0 := symvalue
__typedArg1 := tmp19085
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19087 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(3), tmp19086)
}
__typedArg0 := MakeInteger(3)
__typedArg1 := tmp19086
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19088 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector_1_6, tmp19087)
}
__typedArg0 := symvector_1_6
__typedArg1 := tmp19087
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19089 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19088)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19088
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19090 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector_2, tmp19089)
}
__typedArg0 := symvector_2
__typedArg1 := tmp19089
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19091 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19090)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19090
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19092 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp19091)
}
__typedArg0 := symvector
__typedArg1 := tmp19091
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19093 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19092)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19092
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19094 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symuserdefs, tmp19093)
}
__typedArg0 := symuserdefs
__typedArg1 := tmp19093
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19095 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19094)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19094
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19096 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symupdate_1lambda_1table, tmp19095)
}
__typedArg0 := symupdate_1lambda_1table
__typedArg1 := tmp19095
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19097 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19096)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19096
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19098 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symundefmacro, tmp19097)
}
__typedArg0 := symundefmacro
__typedArg1 := tmp19097
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19099 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19098)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19098
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19100 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symuntrack, tmp19099)
}
__typedArg0 := symuntrack
__typedArg1 := tmp19099
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19101 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19100)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19100
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19102 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunion, tmp19101)
}
__typedArg0 := symunion
__typedArg1 := tmp19101
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19103 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19102)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19102
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19104 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunprofile, tmp19103)
}
__typedArg0 := symunprofile
__typedArg1 := tmp19103
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19105 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(3), tmp19104)
}
__typedArg0 := MakeInteger(3)
__typedArg1 := tmp19104
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19106 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunput, tmp19105)
}
__typedArg0 := symunput
__typedArg1 := tmp19105
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19107 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19106)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19106
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19108 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symundefmacro, tmp19107)
}
__typedArg0 := symundefmacro
__typedArg1 := tmp19107
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19109 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19108)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19108
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19110 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunabsolute, tmp19109)
}
__typedArg0 := symunabsolute
__typedArg1 := tmp19109
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19111 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(5), tmp19110)
}
__typedArg0 := MakeInteger(5)
__typedArg1 := tmp19110
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19112 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symreturn, tmp19111)
}
__typedArg0 := symreturn
__typedArg1 := tmp19111
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19113 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19112)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19112
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19114 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtype, tmp19113)
}
__typedArg0 := symtype
__typedArg1 := tmp19113
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19115 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19114)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19114
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19116 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtuple_2, tmp19115)
}
__typedArg0 := symtuple_2
__typedArg1 := tmp19115
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19117 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19116)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19116
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19118 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtrap_1error, tmp19117)
}
__typedArg0 := symtrap_1error
__typedArg1 := tmp19117
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19119 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19118)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19118
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19120 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtracked, tmp19119)
}
__typedArg0 := symtracked
__typedArg1 := tmp19119
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19121 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19120)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19120
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19122 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtrack, tmp19121)
}
__typedArg0 := symtrack
__typedArg1 := tmp19121
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19123 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19122)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19122
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19124 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtlstr, tmp19123)
}
__typedArg0 := symtlstr
__typedArg1 := tmp19123
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19125 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19124)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19124
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19126 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symthaw, tmp19125)
}
__typedArg0 := symthaw
__typedArg1 := tmp19125
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19127 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19126)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19126
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19128 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtc_2, tmp19127)
}
__typedArg0 := symtc_2
__typedArg1 := tmp19127
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19129 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19128)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19128
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19130 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtc, tmp19129)
}
__typedArg0 := symtc
__typedArg1 := tmp19129
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19131 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19130)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19130
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19132 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtl, tmp19131)
}
__typedArg0 := symtl
__typedArg1 := tmp19131
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19133 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19132)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19132
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19134 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtail, tmp19133)
}
__typedArg0 := symtail
__typedArg1 := tmp19133
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19135 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19134)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19134
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19136 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsystemf, tmp19135)
}
__typedArg0 := symsystemf
__typedArg1 := tmp19135
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19137 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19136)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19136
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19138 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol_2, tmp19137)
}
__typedArg0 := symsymbol_2
__typedArg1 := tmp19137
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19139 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19138)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19138
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19140 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsum, tmp19139)
}
__typedArg0 := symsum
__typedArg1 := tmp19139
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19141 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(3), tmp19140)
}
__typedArg0 := MakeInteger(3)
__typedArg1 := tmp19140
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19142 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsubst, tmp19141)
}
__typedArg0 := symsubst
__typedArg1 := tmp19141
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19143 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19142)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19142
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19144 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring_2, tmp19143)
}
__typedArg0 := symstring_2
__typedArg1 := tmp19143
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19145 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19144)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19144
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19146 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring_1_6symbol, tmp19145)
}
__typedArg0 := symstring_1_6symbol
__typedArg1 := tmp19145
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19147 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19146)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19146
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19148 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring_1_6n, tmp19147)
}
__typedArg0 := symstring_1_6n
__typedArg1 := tmp19147
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19149 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19148)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19148
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19150 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstr, tmp19149)
}
__typedArg0 := symstr
__typedArg1 := tmp19149
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19151 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19150)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19150
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19152 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstoutput, tmp19151)
}
__typedArg0 := symstoutput
__typedArg1 := tmp19151
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19153 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19152)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19152
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19154 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstinput, tmp19153)
}
__typedArg0 := symstinput
__typedArg1 := tmp19153
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19155 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19154)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19154
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19156 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstep_2, tmp19155)
}
__typedArg0 := symstep_2
__typedArg1 := tmp19155
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19157 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19156)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19156
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19158 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstep, tmp19157)
}
__typedArg0 := symstep
__typedArg1 := tmp19157
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19159 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19158)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19158
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19160 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symspy_2, tmp19159)
}
__typedArg0 := symspy_2
__typedArg1 := tmp19159
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19161 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19160)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19160
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19162 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symspy, tmp19161)
}
__typedArg0 := symspy
__typedArg1 := tmp19161
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19163 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19162)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19162
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19164 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symspecialise, tmp19163)
}
__typedArg0 := symspecialise
__typedArg1 := tmp19163
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19165 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19164)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19164
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19166 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsnd, tmp19165)
}
__typedArg0 := symsnd
__typedArg1 := tmp19165
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19167 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19166)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19166
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19168 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsimple_1error, tmp19167)
}
__typedArg0 := symsimple_1error
__typedArg1 := tmp19167
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19169 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19168)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19168
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19170 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symset, tmp19169)
}
__typedArg0 := symset
__typedArg1 := tmp19169
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19171 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19170)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19170
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19172 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symreverse, tmp19171)
}
__typedArg0 := symreverse
__typedArg1 := tmp19171
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19173 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19172)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19172
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19174 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symremove, tmp19173)
}
__typedArg0 := symremove
__typedArg1 := tmp19173
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19175 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19174)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19174
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19176 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4unregister_1source_1form, tmp19175)
}
__typedArg0 := symshen_4unregister_1source_1form
__typedArg1 := tmp19175
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19177 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19176)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19176
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19178 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4register_1source_1form, tmp19177)
}
__typedArg0 := symshen_4register_1source_1form
__typedArg1 := tmp19177
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19179 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19178)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19178
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19180 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symrelease, tmp19179)
}
__typedArg0 := symrelease
__typedArg1 := tmp19179
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19181 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19180)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19180
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19182 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symreceive, tmp19181)
}
__typedArg0 := symreceive
__typedArg1 := tmp19181
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19183 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19182)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19182
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19184 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1unit_1string, tmp19183)
}
__typedArg0 := symread_1unit_1string
__typedArg1 := tmp19183
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19185 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19184)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19184
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19186 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1from_1string_1unprocessed, tmp19185)
}
__typedArg0 := symread_1from_1string_1unprocessed
__typedArg1 := tmp19185
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19187 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19186)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19186
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19188 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1from_1string, tmp19187)
}
__typedArg0 := symread_1from_1string
__typedArg1 := tmp19187
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19189 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19188)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19188
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19190 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1byte, tmp19189)
}
__typedArg0 := symread_1byte
__typedArg1 := tmp19189
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19191 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19190)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19190
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19192 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread, tmp19191)
}
__typedArg0 := symread
__typedArg1 := tmp19191
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19193 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19192)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19192
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19194 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1file, tmp19193)
}
__typedArg0 := symread_1file
__typedArg1 := tmp19193
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19195 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19194)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19194
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19196 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1file_1as_1bytelist, tmp19195)
}
__typedArg0 := symread_1file_1as_1bytelist
__typedArg1 := tmp19195
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19197 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19196)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19196
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19198 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1file_1as_1string, tmp19197)
}
__typedArg0 := symread_1file_1as_1string
__typedArg1 := tmp19197
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19199 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(4), tmp19198)
}
__typedArg0 := MakeInteger(4)
__typedArg1 := tmp19198
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19200 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symput, tmp19199)
}
__typedArg0 := symput
__typedArg1 := tmp19199
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19201 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19200)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19200
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19202 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symprotect, tmp19201)
}
__typedArg0 := symprotect
__typedArg1 := tmp19201
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19203 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19202)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19202
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19204 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympreclude_1all_1but, tmp19203)
}
__typedArg0 := sympreclude_1all_1but
__typedArg1 := tmp19203
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19205 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19204)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19204
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19206 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympreclude, tmp19205)
}
__typedArg0 := sympreclude
__typedArg1 := tmp19205
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19207 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19206)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19206
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19208 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symps, tmp19207)
}
__typedArg0 := symps
__typedArg1 := tmp19207
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19209 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19208)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19208
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19210 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympr, tmp19209)
}
__typedArg0 := sympr
__typedArg1 := tmp19209
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19211 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19210)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19210
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19212 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symprofile_1results, tmp19211)
}
__typedArg0 := symprofile_1results
__typedArg1 := tmp19211
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19213 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19212)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19212
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19214 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symprolog_1memory, tmp19213)
}
__typedArg0 := symprolog_1memory
__typedArg1 := tmp19213
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19215 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19214)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19214
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19216 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4printF, tmp19215)
}
__typedArg0 := symshen_4printF
__typedArg1 := tmp19215
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19217 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19216)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19216
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19218 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4print_1freshterm, tmp19217)
}
__typedArg0 := symshen_4print_1freshterm
__typedArg1 := tmp19217
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19219 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19218)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19218
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19220 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4print_1prolog_1vector, tmp19219)
}
__typedArg0 := symshen_4print_1prolog_1vector
__typedArg1 := tmp19219
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19221 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19220)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19220
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19222 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symprofile, tmp19221)
}
__typedArg0 := symprofile
__typedArg1 := tmp19221
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19223 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19222)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19222
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19224 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symprint, tmp19223)
}
__typedArg0 := symprint
__typedArg1 := tmp19223
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19225 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19224)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19224
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19226 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympreclude_1all_1but, tmp19225)
}
__typedArg0 := sympreclude_1all_1but
__typedArg1 := tmp19225
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19227 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19226)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19226
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19228 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympos, tmp19227)
}
__typedArg0 := sympos
__typedArg1 := tmp19227
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19229 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19228)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19228
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19230 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symporters, tmp19229)
}
__typedArg0 := symporters
__typedArg1 := tmp19229
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19231 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19230)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19230
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19232 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symport, tmp19231)
}
__typedArg0 := symport
__typedArg1 := tmp19231
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19233 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19232)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19232
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19234 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympackage_2, tmp19233)
}
__typedArg0 := sympackage_2
__typedArg1 := tmp19233
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19235 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(3), tmp19234)
}
__typedArg0 := MakeInteger(3)
__typedArg1 := tmp19234
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19236 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympackage, tmp19235)
}
__typedArg0 := sympackage
__typedArg1 := tmp19235
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19237 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19236)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19236
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19238 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symos, tmp19237)
}
__typedArg0 := symos
__typedArg1 := tmp19237
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19239 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19238)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19238
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19240 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symor, tmp19239)
}
__typedArg0 := symor
__typedArg1 := tmp19239
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19241 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19240)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19240
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19242 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symoptimise_2, tmp19241)
}
__typedArg0 := symoptimise_2
__typedArg1 := tmp19241
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19243 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19242)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19242
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19244 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symoptimise, tmp19243)
}
__typedArg0 := symoptimise
__typedArg1 := tmp19243
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19245 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19244)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19244
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19246 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symopen, tmp19245)
}
__typedArg0 := symopen
__typedArg1 := tmp19245
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19247 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19246)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19246
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19248 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symoccurs_1check, tmp19247)
}
__typedArg0 := symoccurs_1check
__typedArg1 := tmp19247
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19249 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19248)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19248
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19250 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symoccurs_2, tmp19249)
}
__typedArg0 := symoccurs_2
__typedArg1 := tmp19249
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19251 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19250)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19250
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19252 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symoccurrences, tmp19251)
}
__typedArg0 := symoccurrences
__typedArg1 := tmp19251
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19253 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19252)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19252
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19254 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symoccurs_1check, tmp19253)
}
__typedArg0 := symoccurs_1check
__typedArg1 := tmp19253
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19255 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19254)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19254
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19256 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber_2, tmp19255)
}
__typedArg0 := symnumber_2
__typedArg1 := tmp19255
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19257 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19256)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19256
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19258 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symn_1_6string, tmp19257)
}
__typedArg0 := symn_1_6string
__typedArg1 := tmp19257
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19259 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19258)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19258
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19260 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnth, tmp19259)
}
__typedArg0 := symnth
__typedArg1 := tmp19259
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19261 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19260)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19260
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19262 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnot, tmp19261)
}
__typedArg0 := symnot
__typedArg1 := tmp19261
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19263 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19262)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19262
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19264 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnl, tmp19263)
}
__typedArg0 := symnl
__typedArg1 := tmp19263
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19265 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19264)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19264
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19266 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symmaxinferences, tmp19265)
}
__typedArg0 := symmaxinferences
__typedArg1 := tmp19265
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19267 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19266)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19266
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19268 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symmapcan, tmp19267)
}
__typedArg0 := symmapcan
__typedArg1 := tmp19267
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19269 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19268)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19268
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19270 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symmap, tmp19269)
}
__typedArg0 := symmap
__typedArg1 := tmp19269
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19271 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19270)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19270
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19272 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symmacroexpand, tmp19271)
}
__typedArg0 := symmacroexpand
__typedArg1 := tmp19271
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19273 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19272)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19272
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19274 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp19273)
}
__typedArg0 := symvector
__typedArg1 := tmp19273
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19275 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19274)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19274
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19276 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_a, tmp19275)
}
__typedArg0 := sym_5_a
__typedArg1 := tmp19275
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19277 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19276)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19276
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19278 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5, tmp19277)
}
__typedArg0 := sym_5
__typedArg1 := tmp19277
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19279 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19278)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19278
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19280 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symload, tmp19279)
}
__typedArg0 := symload
__typedArg1 := tmp19279
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19281 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19280)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19280
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19282 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19281)
}
__typedArg0 := symlist
__typedArg1 := tmp19281
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19283 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19282)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19282
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19284 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlineread, tmp19283)
}
__typedArg0 := symlineread
__typedArg1 := tmp19283
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19285 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19284)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19284
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19286 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlimit, tmp19285)
}
__typedArg0 := symlimit
__typedArg1 := tmp19285
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19287 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19286)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19286
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19288 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlength, tmp19287)
}
__typedArg0 := symlength
__typedArg1 := tmp19287
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19289 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19288)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19288
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19290 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlanguage, tmp19289)
}
__typedArg0 := symlanguage
__typedArg1 := tmp19289
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19291 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(6), tmp19290)
}
__typedArg0 := MakeInteger(6)
__typedArg1 := tmp19290
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19292 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symis_b, tmp19291)
}
__typedArg0 := symis_b
__typedArg1 := tmp19291
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19293 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(6), tmp19292)
}
__typedArg0 := MakeInteger(6)
__typedArg1 := tmp19292
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19294 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symis, tmp19293)
}
__typedArg0 := symis
__typedArg1 := tmp19293
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19295 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19294)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19294
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19296 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symit, tmp19295)
}
__typedArg0 := symit
__typedArg1 := tmp19295
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19297 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19296)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19296
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19298 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminternal, tmp19297)
}
__typedArg0 := syminternal
__typedArg1 := tmp19297
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19299 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19298)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19298
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19300 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symintersection, tmp19299)
}
__typedArg0 := symintersection
__typedArg1 := tmp19299
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19301 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19300)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19300
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19302 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminclude_1all_1but, tmp19301)
}
__typedArg0 := syminclude_1all_1but
__typedArg1 := tmp19301
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19303 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19302)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19302
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19304 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symimplementation, tmp19303)
}
__typedArg0 := symimplementation
__typedArg1 := tmp19303
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19305 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19304)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19304
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19306 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminput_7, tmp19305)
}
__typedArg0 := syminput_7
__typedArg1 := tmp19305
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19307 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19306)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19306
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19308 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminput, tmp19307)
}
__typedArg0 := syminput
__typedArg1 := tmp19307
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19309 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19308)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19308
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19310 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminferences, tmp19309)
}
__typedArg0 := syminferences
__typedArg1 := tmp19309
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19311 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19310)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19310
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19312 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symintern, tmp19311)
}
__typedArg0 := symintern
__typedArg1 := tmp19311
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19313 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19312)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19312
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19314 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminternal, tmp19313)
}
__typedArg0 := syminternal
__typedArg1 := tmp19313
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19315 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19314)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19314
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19316 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminteger_2, tmp19315)
}
__typedArg0 := syminteger_2
__typedArg1 := tmp19315
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19317 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19316)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19316
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19318 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symin_1package, tmp19317)
}
__typedArg0 := symin_1package
__typedArg1 := tmp19317
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19319 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19318)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19318
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19320 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symincluded, tmp19319)
}
__typedArg0 := symincluded
__typedArg1 := tmp19319
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19321 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19320)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19320
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19322 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminclude, tmp19321)
}
__typedArg0 := syminclude
__typedArg1 := tmp19321
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19323 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(3), tmp19322)
}
__typedArg0 := MakeInteger(3)
__typedArg1 := tmp19322
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19324 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp19323)
}
__typedArg0 := symif
__typedArg1 := tmp19323
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19325 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19324)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19324
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19326 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhush, tmp19325)
}
__typedArg0 := symhush
__typedArg1 := tmp19325
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19327 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19326)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19326
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19328 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhush_2, tmp19327)
}
__typedArg0 := symhush_2
__typedArg1 := tmp19327
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19329 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19328)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19328
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19330 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhead, tmp19329)
}
__typedArg0 := symhead
__typedArg1 := tmp19329
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19331 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19330)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19330
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19332 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhdstr, tmp19331)
}
__typedArg0 := symhdstr
__typedArg1 := tmp19331
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19333 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19332)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19332
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19334 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhdv, tmp19333)
}
__typedArg0 := symhdv
__typedArg1 := tmp19333
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19335 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19334)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19334
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19336 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhd, tmp19335)
}
__typedArg0 := symhd
__typedArg1 := tmp19335
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19337 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19336)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19336
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19338 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhash, tmp19337)
}
__typedArg0 := symhash
__typedArg1 := tmp19337
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19339 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19338)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19338
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19340 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_a, tmp19339)
}
__typedArg0 := sym_a
__typedArg1 := tmp19339
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19341 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19340)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19340
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19342 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_6_a, tmp19341)
}
__typedArg0 := sym_6_a
__typedArg1 := tmp19341
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19343 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19342)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19342
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19344 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_6, tmp19343)
}
__typedArg0 := sym_6
__typedArg1 := tmp19343
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19345 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19344)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19344
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19346 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1vector, tmp19345)
}
__typedArg0 := sym_5_1vector
__typedArg1 := tmp19345
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19347 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19346)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19346
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19348 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1address, tmp19347)
}
__typedArg0 := sym_5_1address
__typedArg1 := tmp19347
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19349 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(3), tmp19348)
}
__typedArg0 := MakeInteger(3)
__typedArg1 := tmp19348
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19350 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symaddress_1_6, tmp19349)
}
__typedArg0 := symaddress_1_6
__typedArg1 := tmp19349
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19351 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19350)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19350
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19352 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symget_1time, tmp19351)
}
__typedArg0 := symget_1time
__typedArg1 := tmp19351
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19353 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(3), tmp19352)
}
__typedArg0 := MakeInteger(3)
__typedArg1 := tmp19352
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19354 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symget, tmp19353)
}
__typedArg0 := symget
__typedArg1 := tmp19353
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19355 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19354)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19354
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19356 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symgensym, tmp19355)
}
__typedArg0 := symgensym
__typedArg1 := tmp19355
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19357 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19356)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19356
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19358 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfunction, tmp19357)
}
__typedArg0 := symfunction
__typedArg1 := tmp19357
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19359 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19358)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19358
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19360 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfn, tmp19359)
}
__typedArg0 := symfn
__typedArg1 := tmp19359
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19361 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19360)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19360
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19362 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfst, tmp19361)
}
__typedArg0 := symfst
__typedArg1 := tmp19361
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19363 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19362)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19362
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19364 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfresh, tmp19363)
}
__typedArg0 := symfresh
__typedArg1 := tmp19363
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19365 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19364)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19364
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19366 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfreeze, tmp19365)
}
__typedArg0 := symfreeze
__typedArg1 := tmp19365
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19367 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(5), tmp19366)
}
__typedArg0 := MakeInteger(5)
__typedArg1 := tmp19366
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19368 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfork, tmp19367)
}
__typedArg0 := symfork
__typedArg1 := tmp19367
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19369 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19368)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19368
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19370 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symforeign, tmp19369)
}
__typedArg0 := symforeign
__typedArg1 := tmp19369
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19371 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(7), tmp19370)
}
__typedArg0 := MakeInteger(7)
__typedArg1 := tmp19370
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19372 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfindall, tmp19371)
}
__typedArg0 := symfindall
__typedArg1 := tmp19371
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19373 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19372)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19372
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19374 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfix, tmp19373)
}
__typedArg0 := symfix
__typedArg1 := tmp19373
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19375 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19374)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19374
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19376 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfail, tmp19375)
}
__typedArg0 := symfail
__typedArg1 := tmp19375
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19377 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19376)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19376
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19378 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfail_1if, tmp19377)
}
__typedArg0 := symfail_1if
__typedArg1 := tmp19377
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19379 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19378)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19378
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19380 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfactorise_2, tmp19379)
}
__typedArg0 := symfactorise_2
__typedArg1 := tmp19379
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19381 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19380)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19380
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19382 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfactorise, tmp19381)
}
__typedArg0 := symfactorise
__typedArg1 := tmp19381
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19383 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19382)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19382
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19384 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symexternal, tmp19383)
}
__typedArg0 := symexternal
__typedArg1 := tmp19383
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19385 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19384)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19384
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19386 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symexplode, tmp19385)
}
__typedArg0 := symexplode
__typedArg1 := tmp19385
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19387 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19386)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19386
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19388 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symeval_1kl, tmp19387)
}
__typedArg0 := symeval_1kl
__typedArg1 := tmp19387
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19389 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19388)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19388
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19390 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symeval, tmp19389)
}
__typedArg0 := symeval
__typedArg1 := tmp19389
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19391 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19390)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19390
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19392 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symerror_1to_1string, tmp19391)
}
__typedArg0 := symerror_1to_1string
__typedArg1 := tmp19391
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19393 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19392)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19392
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19394 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symexternal, tmp19393)
}
__typedArg0 := symexternal
__typedArg1 := tmp19393
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19395 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19394)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19394
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19396 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symenable_1type_1theory, tmp19395)
}
__typedArg0 := symenable_1type_1theory
__typedArg1 := tmp19395
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19397 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19396)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19396
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19398 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symempty_2, tmp19397)
}
__typedArg0 := symempty_2
__typedArg1 := tmp19397
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19399 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19398)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19398
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19400 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symelement_2, tmp19399)
}
__typedArg0 := symelement_2
__typedArg1 := tmp19399
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19401 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19400)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19400
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19402 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdo, tmp19401)
}
__typedArg0 := symdo
__typedArg1 := tmp19401
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19403 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19402)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19402
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19404 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdifference, tmp19403)
}
__typedArg0 := symdifference
__typedArg1 := tmp19403
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19405 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19404)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19404
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19406 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdestroy, tmp19405)
}
__typedArg0 := symdestroy
__typedArg1 := tmp19405
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19407 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19406)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19406
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19408 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdeclare, tmp19407)
}
__typedArg0 := symdeclare
__typedArg1 := tmp19407
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19409 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19408)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19408
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19410 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdatatypes, tmp19409)
}
__typedArg0 := symdatatypes
__typedArg1 := tmp19409
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19411 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19410)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19410
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19412 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symclose, tmp19411)
}
__typedArg0 := symclose
__typedArg1 := tmp19411
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19413 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19412)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19412
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19414 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcn, tmp19413)
}
__typedArg0 := symcn
__typedArg1 := tmp19413
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19415 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19414)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19414
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19416 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons_2, tmp19415)
}
__typedArg0 := symcons_2
__typedArg1 := tmp19415
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19417 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19416)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19416
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19418 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp19417)
}
__typedArg0 := symcons
__typedArg1 := tmp19417
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19419 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19418)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19418
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19420 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symconcat, tmp19419)
}
__typedArg0 := symconcat
__typedArg1 := tmp19419
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19421 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19420)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19420
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19422 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcompile, tmp19421)
}
__typedArg0 := symcompile
__typedArg1 := tmp19421
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19423 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19422)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19422
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19424 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symchar_1stream_2, tmp19423)
}
__typedArg0 := symchar_1stream_2
__typedArg1 := tmp19423
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19425 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19424)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19424
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19426 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcd, tmp19425)
}
__typedArg0 := symcd
__typedArg1 := tmp19425
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19427 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(5), tmp19426)
}
__typedArg0 := MakeInteger(5)
__typedArg1 := tmp19426
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19428 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcall, tmp19427)
}
__typedArg0 := symcall
__typedArg1 := tmp19427
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19429 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(6), tmp19428)
}
__typedArg0 := MakeInteger(6)
__typedArg1 := tmp19428
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19430 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symbind, tmp19429)
}
__typedArg0 := symbind
__typedArg1 := tmp19429
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19431 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19430)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19430
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19432 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symbound_2, tmp19431)
}
__typedArg0 := symbound_2
__typedArg1 := tmp19431
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19433 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19432)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19432
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19434 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symbootstrap, tmp19433)
}
__typedArg0 := symbootstrap
__typedArg1 := tmp19433
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19435 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19434)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19434
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19436 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean_2, tmp19435)
}
__typedArg0 := symboolean_2
__typedArg1 := tmp19435
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19437 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19436)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19436
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19438 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symatom_2, tmp19437)
}
__typedArg0 := symatom_2
__typedArg1 := tmp19437
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19439 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19438)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19438
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19440 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symassoc, tmp19439)
}
__typedArg0 := symassoc
__typedArg1 := tmp19439
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19441 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19440)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19440
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19442 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symarity, tmp19441)
}
__typedArg0 := symarity
__typedArg1 := tmp19441
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19443 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19442)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19442
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19444 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symappend, tmp19443)
}
__typedArg0 := symappend
__typedArg1 := tmp19443
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19445 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19444)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19444
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19446 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symand, tmp19445)
}
__typedArg0 := symand
__typedArg1 := tmp19445
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19447 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19446)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19446
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19448 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symadjoin, tmp19447)
}
__typedArg0 := symadjoin
__typedArg1 := tmp19447
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19449 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(3), tmp19448)
}
__typedArg0 := MakeInteger(3)
__typedArg1 := tmp19448
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19450 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symaddress_1_6, tmp19449)
}
__typedArg0 := symaddress_1_6
__typedArg1 := tmp19449
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19451 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19450)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19450
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19452 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symabsvector, tmp19451)
}
__typedArg0 := symabsvector
__typedArg1 := tmp19451
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19453 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19452)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19452
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19454 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symabsvector_2, tmp19453)
}
__typedArg0 := symabsvector_2
__typedArg1 := tmp19453
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19455 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19454)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19454
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19456 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symabsolute, tmp19455)
}
__typedArg0 := symabsolute
__typedArg1 := tmp19455
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19457 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19456)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19456
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19458 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symabort, tmp19457)
}
__typedArg0 := symabort
__typedArg1 := tmp19457
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19459 := Call(__e, PrimFunc(symshen_4initialise_1arity_1table), tmp19458)


_ = tmp19459

tmp19460 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp19461 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(";"))
}
__typedArg0 := MakeString(";")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp19462 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":="))
}
__typedArg0 := MakeString(":=")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp19463 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(","))
}
__typedArg0 := MakeString(",")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp19464 := Call(__e, PrimFunc(symvector), MakeInteger(0))


tmp19465 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString("bar!"))
}
__typedArg0 := MakeString("bar!")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp19466 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symabort, Nil)
}
__typedArg0 := symabort
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19467 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symabsolute, tmp19466)
}
__typedArg0 := symabsolute
__typedArg1 := tmp19466
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19468 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symabsvector, tmp19467)
}
__typedArg0 := symabsvector
__typedArg1 := tmp19467
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19469 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symabsvector_2, tmp19468)
}
__typedArg0 := symabsvector_2
__typedArg1 := tmp19468
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19470 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symaddress_1_6, tmp19469)
}
__typedArg0 := symaddress_1_6
__typedArg1 := tmp19469
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19471 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1address, tmp19470)
}
__typedArg0 := sym_5_1address
__typedArg1 := tmp19470
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19472 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symadjoin, tmp19471)
}
__typedArg0 := symadjoin
__typedArg1 := tmp19471
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19473 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symand, tmp19472)
}
__typedArg0 := symand
__typedArg1 := tmp19472
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19474 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symappend, tmp19473)
}
__typedArg0 := symappend
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
return PrimCons(symassoc, tmp19475)
}
__typedArg0 := symassoc
__typedArg1 := tmp19475
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19477 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symassertz, tmp19476)
}
__typedArg0 := symassertz
__typedArg1 := tmp19476
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19478 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symasserta, tmp19477)
}
__typedArg0 := symasserta
__typedArg1 := tmp19477
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19479 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symatom_2, tmp19478)
}
__typedArg0 := symatom_2
__typedArg1 := tmp19478
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19480 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstep_2, tmp19479)
}
__typedArg0 := symstep_2
__typedArg1 := tmp19479
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19481 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symspy_2, tmp19480)
}
__typedArg0 := symspy_2
__typedArg1 := tmp19480
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19482 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19465, tmp19481)
}
__typedArg0 := tmp19465
__typedArg1 := tmp19481
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19483 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symbootstrap, tmp19482)
}
__typedArg0 := symbootstrap
__typedArg1 := tmp19482
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19484 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, tmp19483)
}
__typedArg0 := symboolean
__typedArg1 := tmp19483
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19485 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean_2, tmp19484)
}
__typedArg0 := symboolean_2
__typedArg1 := tmp19484
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19486 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symbound_2, tmp19485)
}
__typedArg0 := symbound_2
__typedArg1 := tmp19485
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19487 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symbind, tmp19486)
}
__typedArg0 := symbind
__typedArg1 := tmp19486
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19488 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symclose, tmp19487)
}
__typedArg0 := symclose
__typedArg1 := tmp19487
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19489 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symchar_1stream_2, tmp19488)
}
__typedArg0 := symchar_1stream_2
__typedArg1 := tmp19488
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19490 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcall, tmp19489)
}
__typedArg0 := symcall
__typedArg1 := tmp19489
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19491 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcases, tmp19490)
}
__typedArg0 := symcases
__typedArg1 := tmp19490
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19492 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcd, tmp19491)
}
__typedArg0 := symcd
__typedArg1 := tmp19491
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19493 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcompile, tmp19492)
}
__typedArg0 := symcompile
__typedArg1 := tmp19492
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19494 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symconcat, tmp19493)
}
__typedArg0 := symconcat
__typedArg1 := tmp19493
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19495 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcond, tmp19494)
}
__typedArg0 := symcond
__typedArg1 := tmp19494
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19496 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp19495)
}
__typedArg0 := symcons
__typedArg1 := tmp19495
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19497 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons_2, tmp19496)
}
__typedArg0 := symcons_2
__typedArg1 := tmp19496
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19498 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcn, tmp19497)
}
__typedArg0 := symcn
__typedArg1 := tmp19497
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19499 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symctxt, tmp19498)
}
__typedArg0 := symctxt
__typedArg1 := tmp19498
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19500 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdatatypes, tmp19499)
}
__typedArg0 := symdatatypes
__typedArg1 := tmp19499
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19501 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdatatype, tmp19500)
}
__typedArg0 := symdatatype
__typedArg1 := tmp19500
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19502 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdeclare, tmp19501)
}
__typedArg0 := symdeclare
__typedArg1 := tmp19501
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19503 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefprolog, tmp19502)
}
__typedArg0 := symdefprolog
__typedArg1 := tmp19502
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19504 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefcc, tmp19503)
}
__typedArg0 := symdefcc
__typedArg1 := tmp19503
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19505 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefmacro, tmp19504)
}
__typedArg0 := symdefmacro
__typedArg1 := tmp19504
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19506 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefine, tmp19505)
}
__typedArg0 := symdefine
__typedArg1 := tmp19505
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19507 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefun, tmp19506)
}
__typedArg0 := symdefun
__typedArg1 := tmp19506
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19508 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdestroy, tmp19507)
}
__typedArg0 := symdestroy
__typedArg1 := tmp19507
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19509 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdifference, tmp19508)
}
__typedArg0 := symdifference
__typedArg1 := tmp19508
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19510 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdo, tmp19509)
}
__typedArg0 := symdo
__typedArg1 := tmp19509
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19511 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symelement_2, tmp19510)
}
__typedArg0 := symelement_2
__typedArg1 := tmp19510
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19512 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symempty_2, tmp19511)
}
__typedArg0 := symempty_2
__typedArg1 := tmp19511
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19513 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symerror, tmp19512)
}
__typedArg0 := symerror
__typedArg1 := tmp19512
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19514 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symerror_1to_1string, tmp19513)
}
__typedArg0 := symerror_1to_1string
__typedArg1 := tmp19513
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19515 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symeval, tmp19514)
}
__typedArg0 := symeval
__typedArg1 := tmp19514
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19516 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symeval_1kl, tmp19515)
}
__typedArg0 := symeval_1kl
__typedArg1 := tmp19515
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19517 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symexception, tmp19516)
}
__typedArg0 := symexception
__typedArg1 := tmp19516
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19518 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symexternal, tmp19517)
}
__typedArg0 := symexternal
__typedArg1 := tmp19517
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19519 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symexplode, tmp19518)
}
__typedArg0 := symexplode
__typedArg1 := tmp19518
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19520 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symenable_1type_1theory, tmp19519)
}
__typedArg0 := symenable_1type_1theory
__typedArg1 := tmp19519
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19521 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(False, tmp19520)
}
__typedArg0 := False
__typedArg1 := tmp19520
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19522 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfindall, tmp19521)
}
__typedArg0 := symfindall
__typedArg1 := tmp19521
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19523 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfactorise_2, tmp19522)
}
__typedArg0 := symfactorise_2
__typedArg1 := tmp19522
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19524 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfactorise, tmp19523)
}
__typedArg0 := symfactorise
__typedArg1 := tmp19523
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19525 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfail_1if, tmp19524)
}
__typedArg0 := symfail_1if
__typedArg1 := tmp19524
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19526 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfail, tmp19525)
}
__typedArg0 := symfail
__typedArg1 := tmp19525
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19527 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfile, tmp19526)
}
__typedArg0 := symfile
__typedArg1 := tmp19526
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19528 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfix, tmp19527)
}
__typedArg0 := symfix
__typedArg1 := tmp19527
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19529 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symforeign, tmp19528)
}
__typedArg0 := symforeign
__typedArg1 := tmp19528
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19530 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfork, tmp19529)
}
__typedArg0 := symfork
__typedArg1 := tmp19529
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19531 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfresh, tmp19530)
}
__typedArg0 := symfresh
__typedArg1 := tmp19530
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19532 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfreeze, tmp19531)
}
__typedArg0 := symfreeze
__typedArg1 := tmp19531
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19533 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfst, tmp19532)
}
__typedArg0 := symfst
__typedArg1 := tmp19532
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19534 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfunction, tmp19533)
}
__typedArg0 := symfunction
__typedArg1 := tmp19533
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19535 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfn, tmp19534)
}
__typedArg0 := symfn
__typedArg1 := tmp19534
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19536 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symgensym, tmp19535)
}
__typedArg0 := symgensym
__typedArg1 := tmp19535
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19537 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symget_1time, tmp19536)
}
__typedArg0 := symget_1time
__typedArg1 := tmp19536
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19538 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symget, tmp19537)
}
__typedArg0 := symget
__typedArg1 := tmp19537
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19539 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhash, tmp19538)
}
__typedArg0 := symhash
__typedArg1 := tmp19538
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19540 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhdstr, tmp19539)
}
__typedArg0 := symhdstr
__typedArg1 := tmp19539
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19541 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhdv, tmp19540)
}
__typedArg0 := symhdv
__typedArg1 := tmp19540
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19542 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhd, tmp19541)
}
__typedArg0 := symhd
__typedArg1 := tmp19541
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19543 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhead, tmp19542)
}
__typedArg0 := symhead
__typedArg1 := tmp19542
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19544 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhush, tmp19543)
}
__typedArg0 := symhush
__typedArg1 := tmp19543
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19545 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhush_2, tmp19544)
}
__typedArg0 := symhush_2
__typedArg1 := tmp19544
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19546 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp19545)
}
__typedArg0 := symif
__typedArg1 := tmp19545
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19547 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symimplementation, tmp19546)
}
__typedArg0 := symimplementation
__typedArg1 := tmp19546
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19548 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminternal, tmp19547)
}
__typedArg0 := syminternal
__typedArg1 := tmp19547
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19549 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symin_1package, tmp19548)
}
__typedArg0 := symin_1package
__typedArg1 := tmp19548
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19550 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symin, tmp19549)
}
__typedArg0 := symin
__typedArg1 := tmp19549
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19551 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symis_b, tmp19550)
}
__typedArg0 := symis_b
__typedArg1 := tmp19550
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19552 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symis, tmp19551)
}
__typedArg0 := symis
__typedArg1 := tmp19551
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19553 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symit, tmp19552)
}
__typedArg0 := symit
__typedArg1 := tmp19552
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19554 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminclude_1all_1but, tmp19553)
}
__typedArg0 := syminclude_1all_1but
__typedArg1 := tmp19553
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19555 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminclude, tmp19554)
}
__typedArg0 := syminclude
__typedArg1 := tmp19554
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19556 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symincluded, tmp19555)
}
__typedArg0 := symincluded
__typedArg1 := tmp19555
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19557 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminline, tmp19556)
}
__typedArg0 := syminline
__typedArg1 := tmp19556
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19558 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminput_7, tmp19557)
}
__typedArg0 := syminput_7
__typedArg1 := tmp19557
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19559 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminput, tmp19558)
}
__typedArg0 := syminput
__typedArg1 := tmp19558
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19560 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminteger_2, tmp19559)
}
__typedArg0 := syminteger_2
__typedArg1 := tmp19559
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19561 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symintern, tmp19560)
}
__typedArg0 := symintern
__typedArg1 := tmp19560
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19562 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminferences, tmp19561)
}
__typedArg0 := syminferences
__typedArg1 := tmp19561
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19563 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symintersection, tmp19562)
}
__typedArg0 := symintersection
__typedArg1 := tmp19562
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19564 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symis, tmp19563)
}
__typedArg0 := symis
__typedArg1 := tmp19563
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19565 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlanguage, tmp19564)
}
__typedArg0 := symlanguage
__typedArg1 := tmp19564
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19566 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp19565)
}
__typedArg0 := symlambda
__typedArg1 := tmp19565
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19567 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlazy, tmp19566)
}
__typedArg0 := symlazy
__typedArg1 := tmp19566
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19568 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp19567)
}
__typedArg0 := symlet
__typedArg1 := tmp19567
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19569 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlength, tmp19568)
}
__typedArg0 := symlength
__typedArg1 := tmp19568
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19570 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlimit, tmp19569)
}
__typedArg0 := symlimit
__typedArg1 := tmp19569
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19571 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlineread, tmp19570)
}
__typedArg0 := symlineread
__typedArg1 := tmp19570
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19572 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19571)
}
__typedArg0 := symlist
__typedArg1 := tmp19571
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19573 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symloaded, tmp19572)
}
__typedArg0 := symloaded
__typedArg1 := tmp19572
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19574 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symload, tmp19573)
}
__typedArg0 := symload
__typedArg1 := tmp19573
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19575 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symmake_1string, tmp19574)
}
__typedArg0 := symmake_1string
__typedArg1 := tmp19574
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19576 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symmap, tmp19575)
}
__typedArg0 := symmap
__typedArg1 := tmp19575
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19577 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symmapcan, tmp19576)
}
__typedArg0 := symmapcan
__typedArg1 := tmp19576
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19578 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symmaxinferences, tmp19577)
}
__typedArg0 := symmaxinferences
__typedArg1 := tmp19577
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19579 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symmacroexpand, tmp19578)
}
__typedArg0 := symmacroexpand
__typedArg1 := tmp19578
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19580 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symmode, tmp19579)
}
__typedArg0 := symmode
__typedArg1 := tmp19579
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19581 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnl, tmp19580)
}
__typedArg0 := symnl
__typedArg1 := tmp19580
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19582 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnot, tmp19581)
}
__typedArg0 := symnot
__typedArg1 := tmp19581
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19583 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnth, tmp19582)
}
__typedArg0 := symnth
__typedArg1 := tmp19582
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19584 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnull, tmp19583)
}
__typedArg0 := symnull
__typedArg1 := tmp19583
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19585 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp19584)
}
__typedArg0 := symnumber
__typedArg1 := tmp19584
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19586 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber_2, tmp19585)
}
__typedArg0 := symnumber_2
__typedArg1 := tmp19585
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19587 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symn_1_6string, tmp19586)
}
__typedArg0 := symn_1_6string
__typedArg1 := tmp19586
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19588 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symoccurs_2, tmp19587)
}
__typedArg0 := symoccurs_2
__typedArg1 := tmp19587
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19589 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symoccurs_1check, tmp19588)
}
__typedArg0 := symoccurs_1check
__typedArg1 := tmp19588
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19590 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symoccurrences, tmp19589)
}
__typedArg0 := symoccurrences
__typedArg1 := tmp19589
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19591 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symopen, tmp19590)
}
__typedArg0 := symopen
__typedArg1 := tmp19590
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19592 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symoptimise_2, tmp19591)
}
__typedArg0 := symoptimise_2
__typedArg1 := tmp19591
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19593 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symoptimise, tmp19592)
}
__typedArg0 := symoptimise
__typedArg1 := tmp19592
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19594 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symor, tmp19593)
}
__typedArg0 := symor
__typedArg1 := tmp19593
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19595 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symos, tmp19594)
}
__typedArg0 := symos
__typedArg1 := tmp19594
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19596 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symout, tmp19595)
}
__typedArg0 := symout
__typedArg1 := tmp19595
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19597 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symoutput, tmp19596)
}
__typedArg0 := symoutput
__typedArg1 := tmp19596
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19598 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympackage, tmp19597)
}
__typedArg0 := sympackage
__typedArg1 := tmp19597
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19599 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symport, tmp19598)
}
__typedArg0 := symport
__typedArg1 := tmp19598
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19600 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symporters, tmp19599)
}
__typedArg0 := symporters
__typedArg1 := tmp19599
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19601 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympos, tmp19600)
}
__typedArg0 := sympos
__typedArg1 := tmp19600
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19602 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympr, tmp19601)
}
__typedArg0 := sympr
__typedArg1 := tmp19601
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19603 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symprint, tmp19602)
}
__typedArg0 := symprint
__typedArg1 := tmp19602
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19604 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symprolog_1memory, tmp19603)
}
__typedArg0 := symprolog_1memory
__typedArg1 := tmp19603
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19605 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symprofile, tmp19604)
}
__typedArg0 := symprofile
__typedArg1 := tmp19604
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19606 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symprofile_1results, tmp19605)
}
__typedArg0 := symprofile_1results
__typedArg1 := tmp19605
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19607 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symprotect, tmp19606)
}
__typedArg0 := symprotect
__typedArg1 := tmp19606
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19608 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symprolog_2, tmp19607)
}
__typedArg0 := symprolog_2
__typedArg1 := tmp19607
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19609 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symps, tmp19608)
}
__typedArg0 := symps
__typedArg1 := tmp19608
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19610 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympreclude_1all_1but, tmp19609)
}
__typedArg0 := sympreclude_1all_1but
__typedArg1 := tmp19609
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19611 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympreclude, tmp19610)
}
__typedArg0 := sympreclude
__typedArg1 := tmp19610
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19612 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symput, tmp19611)
}
__typedArg0 := symput
__typedArg1 := tmp19611
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19613 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympackage_2, tmp19612)
}
__typedArg0 := sympackage_2
__typedArg1 := tmp19612
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19614 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1unit_1string, tmp19613)
}
__typedArg0 := symread_1unit_1string
__typedArg1 := tmp19613
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19615 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1from_1string_1unprocessed, tmp19614)
}
__typedArg0 := symread_1from_1string_1unprocessed
__typedArg1 := tmp19614
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19616 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1from_1string, tmp19615)
}
__typedArg0 := symread_1from_1string
__typedArg1 := tmp19615
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19617 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1byte, tmp19616)
}
__typedArg0 := symread_1byte
__typedArg1 := tmp19616
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19618 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1file_1as_1string, tmp19617)
}
__typedArg0 := symread_1file_1as_1string
__typedArg1 := tmp19617
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19619 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1file_1as_1bytelist, tmp19618)
}
__typedArg0 := symread_1file_1as_1bytelist
__typedArg1 := tmp19618
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19620 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1file, tmp19619)
}
__typedArg0 := symread_1file
__typedArg1 := tmp19619
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19621 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symreceive, tmp19620)
}
__typedArg0 := symreceive
__typedArg1 := tmp19620
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19622 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread, tmp19621)
}
__typedArg0 := symread
__typedArg1 := tmp19621
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19623 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symrelease, tmp19622)
}
__typedArg0 := symrelease
__typedArg1 := tmp19622
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19624 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symremove, tmp19623)
}
__typedArg0 := symremove
__typedArg1 := tmp19623
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19625 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symretract, tmp19624)
}
__typedArg0 := symretract
__typedArg1 := tmp19624
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19626 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symreverse, tmp19625)
}
__typedArg0 := symreverse
__typedArg1 := tmp19625
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19627 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symrun, tmp19626)
}
__typedArg0 := symrun
__typedArg1 := tmp19626
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19628 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstr, tmp19627)
}
__typedArg0 := symstr
__typedArg1 := tmp19627
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19629 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsave, tmp19628)
}
__typedArg0 := symsave
__typedArg1 := tmp19628
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19630 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symset, tmp19629)
}
__typedArg0 := symset
__typedArg1 := tmp19629
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19631 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsimple_1error, tmp19630)
}
__typedArg0 := symsimple_1error
__typedArg1 := tmp19630
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19632 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsnd, tmp19631)
}
__typedArg0 := symsnd
__typedArg1 := tmp19631
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19633 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symspecialise, tmp19632)
}
__typedArg0 := symspecialise
__typedArg1 := tmp19632
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19634 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symspy, tmp19633)
}
__typedArg0 := symspy
__typedArg1 := tmp19633
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19635 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsqts, tmp19634)
}
__typedArg0 := symsqts
__typedArg1 := tmp19634
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19636 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstep, tmp19635)
}
__typedArg0 := symstep
__typedArg1 := tmp19635
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19637 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstoutput, tmp19636)
}
__typedArg0 := symstoutput
__typedArg1 := tmp19636
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19638 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsterror, tmp19637)
}
__typedArg0 := symsterror
__typedArg1 := tmp19637
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19639 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstinput, tmp19638)
}
__typedArg0 := symstinput
__typedArg1 := tmp19638
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19640 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp19639)
}
__typedArg0 := symstring
__typedArg1 := tmp19639
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19641 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp19640)
}
__typedArg0 := symstream
__typedArg1 := tmp19640
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19642 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring_1_6n, tmp19641)
}
__typedArg0 := symstring_1_6n
__typedArg1 := tmp19641
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19643 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring_2, tmp19642)
}
__typedArg0 := symstring_2
__typedArg1 := tmp19642
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19644 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsubst, tmp19643)
}
__typedArg0 := symsubst
__typedArg1 := tmp19643
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19645 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsum, tmp19644)
}
__typedArg0 := symsum
__typedArg1 := tmp19644
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19646 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring_1_6symbol, tmp19645)
}
__typedArg0 := symstring_1_6symbol
__typedArg1 := tmp19645
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19647 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol_2, tmp19646)
}
__typedArg0 := symsymbol_2
__typedArg1 := tmp19646
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19648 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp19647)
}
__typedArg0 := symsymbol
__typedArg1 := tmp19647
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19649 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsystem_1S_2, tmp19648)
}
__typedArg0 := symsystem_1S_2
__typedArg1 := tmp19648
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19650 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsynonyms, tmp19649)
}
__typedArg0 := symsynonyms
__typedArg1 := tmp19649
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19651 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsystemf, tmp19650)
}
__typedArg0 := symsystemf
__typedArg1 := tmp19650
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19652 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtail, tmp19651)
}
__typedArg0 := symtail
__typedArg1 := tmp19651
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19653 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtlv, tmp19652)
}
__typedArg0 := symtlv
__typedArg1 := tmp19652
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19654 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtlstr, tmp19653)
}
__typedArg0 := symtlstr
__typedArg1 := tmp19653
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19655 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtl, tmp19654)
}
__typedArg0 := symtl
__typedArg1 := tmp19654
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19656 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtc, tmp19655)
}
__typedArg0 := symtc
__typedArg1 := tmp19655
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19657 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtc_2, tmp19656)
}
__typedArg0 := symtc_2
__typedArg1 := tmp19656
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19658 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symthaw, tmp19657)
}
__typedArg0 := symthaw
__typedArg1 := tmp19657
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19659 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtime, tmp19658)
}
__typedArg0 := symtime
__typedArg1 := tmp19658
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19660 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtrack, tmp19659)
}
__typedArg0 := symtrack
__typedArg1 := tmp19659
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19661 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtracked, tmp19660)
}
__typedArg0 := symtracked
__typedArg1 := tmp19660
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19662 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtrap_1error, tmp19661)
}
__typedArg0 := symtrap_1error
__typedArg1 := tmp19661
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19663 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(True, tmp19662)
}
__typedArg0 := True
__typedArg1 := tmp19662
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19664 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtuple_2, tmp19663)
}
__typedArg0 := symtuple_2
__typedArg1 := tmp19663
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19665 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtype, tmp19664)
}
__typedArg0 := symtype
__typedArg1 := tmp19664
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19666 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symreturn, tmp19665)
}
__typedArg0 := symreturn
__typedArg1 := tmp19665
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19667 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunabsolute, tmp19666)
}
__typedArg0 := symunabsolute
__typedArg1 := tmp19666
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19668 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symundefmacro, tmp19667)
}
__typedArg0 := symundefmacro
__typedArg1 := tmp19667
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19669 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunprofile, tmp19668)
}
__typedArg0 := symunprofile
__typedArg1 := tmp19668
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19670 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunput, tmp19669)
}
__typedArg0 := symunput
__typedArg1 := tmp19669
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19671 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunion, tmp19670)
}
__typedArg0 := symunion
__typedArg1 := tmp19670
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19672 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunix, tmp19671)
}
__typedArg0 := symunix
__typedArg1 := tmp19671
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19673 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunit, tmp19672)
}
__typedArg0 := symunit
__typedArg1 := tmp19672
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19674 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symuntrack, tmp19673)
}
__typedArg0 := symuntrack
__typedArg1 := tmp19673
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19675 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunspecialise, tmp19674)
}
__typedArg0 := symunspecialise
__typedArg1 := tmp19674
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19676 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symupdate_1lambda_1table, tmp19675)
}
__typedArg0 := symupdate_1lambda_1table
__typedArg1 := tmp19675
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19677 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symu_b, tmp19676)
}
__typedArg0 := symu_b
__typedArg1 := tmp19676
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19678 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symuserdefs, tmp19677)
}
__typedArg0 := symuserdefs
__typedArg1 := tmp19677
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19679 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector_2, tmp19678)
}
__typedArg0 := symvector_2
__typedArg1 := tmp19678
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19680 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp19679)
}
__typedArg0 := symvector
__typedArg1 := tmp19679
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19681 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1vector, tmp19680)
}
__typedArg0 := sym_5_1vector
__typedArg1 := tmp19680
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19682 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector_1_6, tmp19681)
}
__typedArg0 := symvector_1_6
__typedArg1 := tmp19681
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19683 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvalue, tmp19682)
}
__typedArg0 := symvalue
__typedArg1 := tmp19682
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19684 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvar_2, tmp19683)
}
__typedArg0 := symvar_2
__typedArg1 := tmp19683
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19685 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvariable_2, tmp19684)
}
__typedArg0 := symvariable_2
__typedArg1 := tmp19684
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19686 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symverified, tmp19685)
}
__typedArg0 := symverified
__typedArg1 := tmp19685
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19687 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symversion, tmp19686)
}
__typedArg0 := symversion
__typedArg1 := tmp19686
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19688 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwarn, tmp19687)
}
__typedArg0 := symwarn
__typedArg1 := tmp19687
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19689 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhen, tmp19688)
}
__typedArg0 := symwhen
__typedArg1 := tmp19688
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19690 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhere, tmp19689)
}
__typedArg0 := symwhere
__typedArg1 := tmp19689
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19691 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwrite_1byte, tmp19690)
}
__typedArg0 := symwrite_1byte
__typedArg1 := tmp19690
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19692 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwrite_1to_1file, tmp19691)
}
__typedArg0 := symwrite_1to_1file
__typedArg1 := tmp19691
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19693 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwrite_1string, tmp19692)
}
__typedArg0 := symwrite_1string
__typedArg1 := tmp19692
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19694 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symy_1or_1n_2, tmp19693)
}
__typedArg0 := symy_1or_1n_2
__typedArg1 := tmp19693
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19695 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19464, tmp19694)
}
__typedArg0 := tmp19464
__typedArg1 := tmp19694
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19696 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_6_6, tmp19695)
}
__typedArg0 := sym_6_6
__typedArg1 := tmp19695
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19697 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5, tmp19696)
}
__typedArg0 := sym_5
__typedArg1 := tmp19696
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19698 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_a, tmp19697)
}
__typedArg0 := sym_5_a
__typedArg1 := tmp19697
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19699 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_7, tmp19698)
}
__typedArg0 := sym_7
__typedArg1 := tmp19698
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19700 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_d, tmp19699)
}
__typedArg0 := sym_d
__typedArg1 := tmp19699
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19701 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_c, tmp19700)
}
__typedArg0 := sym_c
__typedArg1 := tmp19700
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19702 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1, tmp19701)
}
__typedArg0 := sym_1
__typedArg1 := tmp19701
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19703 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_3, tmp19702)
}
__typedArg0 := sym_3
__typedArg1 := tmp19702
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19704 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5end_6, tmp19703)
}
__typedArg0 := sym_5end_6
__typedArg1 := tmp19703
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19705 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_b_6, tmp19704)
}
__typedArg0 := sym_5_b_6
__typedArg1 := tmp19704
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19706 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_c_4, tmp19705)
}
__typedArg0 := sym_c_4
__typedArg1 := tmp19705
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19707 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_a_a_6, tmp19706)
}
__typedArg0 := sym_a_a_6
__typedArg1 := tmp19706
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19708 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_6, tmp19707)
}
__typedArg0 := sym_6
__typedArg1 := tmp19707
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19709 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_6_a, tmp19708)
}
__typedArg0 := sym_6_a
__typedArg1 := tmp19708
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19710 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_a, tmp19709)
}
__typedArg0 := sym_a
__typedArg1 := tmp19709
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19711 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_a_a, tmp19710)
}
__typedArg0 := sym_a_a
__typedArg1 := tmp19710
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19712 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5e_6, tmp19711)
}
__typedArg0 := sym_5e_6
__typedArg1 := tmp19711
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19713 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_6, tmp19712)
}
__typedArg0 := sym_1_6
__typedArg1 := tmp19712
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19714 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1, tmp19713)
}
__typedArg0 := sym_5_1
__typedArg1 := tmp19713
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19715 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dhush_d, tmp19714)
}
__typedArg0 := sym_dhush_d
__typedArg1 := tmp19714
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19716 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dporters_d, tmp19715)
}
__typedArg0 := sym_dporters_d
__typedArg1 := tmp19715
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19717 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dport_d, tmp19716)
}
__typedArg0 := sym_dport_d
__typedArg1 := tmp19716
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19718 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8s, tmp19717)
}
__typedArg0 := sym_8s
__typedArg1 := tmp19717
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19719 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8p, tmp19718)
}
__typedArg0 := sym_8p
__typedArg1 := tmp19718
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19720 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8v, tmp19719)
}
__typedArg0 := sym_8v
__typedArg1 := tmp19719
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19721 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dproperty_1vector_d, tmp19720)
}
__typedArg0 := sym_dproperty_1vector_d
__typedArg1 := tmp19720
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19722 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_drelease_d, tmp19721)
}
__typedArg0 := sym_drelease_d
__typedArg1 := tmp19721
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19723 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dos_d, tmp19722)
}
__typedArg0 := sym_dos_d
__typedArg1 := tmp19722
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19724 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dmacros_d, tmp19723)
}
__typedArg0 := sym_dmacros_d
__typedArg1 := tmp19723
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19725 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dmaximum_1print_1sequence_1size_d, tmp19724)
}
__typedArg0 := sym_dmaximum_1print_1sequence_1size_d
__typedArg1 := tmp19724
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19726 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dversion_d, tmp19725)
}
__typedArg0 := sym_dversion_d
__typedArg1 := tmp19725
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19727 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dhome_1directory_d, tmp19726)
}
__typedArg0 := sym_dhome_1directory_d
__typedArg1 := tmp19726
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19728 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dstoutput_d, tmp19727)
}
__typedArg0 := sym_dstoutput_d
__typedArg1 := tmp19727
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19729 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dsterror_d, tmp19728)
}
__typedArg0 := sym_dsterror_d
__typedArg1 := tmp19728
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19730 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dstinput_d, tmp19729)
}
__typedArg0 := sym_dstinput_d
__typedArg1 := tmp19729
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19731 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dimplementation_d, tmp19730)
}
__typedArg0 := sym_dimplementation_d
__typedArg1 := tmp19730
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19732 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dlanguage_d, tmp19731)
}
__typedArg0 := sym_dlanguage_d
__typedArg1 := tmp19731
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19733 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym__, tmp19732)
}
__typedArg0 := sym__
__typedArg1 := tmp19732
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19734 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19463, tmp19733)
}
__typedArg0 := tmp19463
__typedArg1 := tmp19733
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19735 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19462, tmp19734)
}
__typedArg0 := tmp19462
__typedArg1 := tmp19734
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19736 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19461, tmp19735)
}
__typedArg0 := tmp19461
__typedArg1 := tmp19735
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19737 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19460, tmp19736)
}
__typedArg0 := tmp19460
__typedArg1 := tmp19736
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19738 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_e_e, tmp19737)
}
__typedArg0 := sym_e_e
__typedArg1 := tmp19737
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19739 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1_1, tmp19738)
}
__typedArg0 := sym_5_1_1
__typedArg1 := tmp19738
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19740 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19739)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19739
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19741 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_i, tmp19740)
}
__typedArg0 := sym_i
__typedArg1 := tmp19740
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19742 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_j, tmp19741)
}
__typedArg0 := sym_j
__typedArg1 := tmp19741
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19743 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_b, tmp19742)
}
__typedArg0 := sym_b
__typedArg1 := tmp19742
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19744 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp19745 := Call(__e, PrimFunc(symput), symshen, symshen_4external_1symbols, tmp19743, tmp19744)


_ = tmp19745

tmp19746 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symabsvector) {
return PrimAbsvector(MakeInteger(0))
}
__typedArg0 := MakeInteger(0)
return Call(__e, PrimFunc(symabsvector), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dempty_1absvector_d, tmp19746)
}
__typedArg0 := symshen_4_dempty_1absvector_d
__typedArg1 := tmp19746
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return


}, 0)

tmp19747 := Call(__e, ns2_1set, symshen_4initialise_1environment, tmp18983)


_ = tmp19747

tmp19748 := MakeNative(func(__e *ControlFlow) {
tmp19749 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, Nil)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19749

tmp19750 := MakeNative(func(__e *ControlFlow) {
V6045 := __e.Get(1)
_ = V6045
__e.Return(MakeNative(func(__e *ControlFlow) {
B6041 := __e.Get(1)
_ = B6041
__e.Return(MakeNative(func(__e *ControlFlow) {
L6042 := __e.Get(1)
_ = L6042
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6043 := __e.Get(1)
_ = Key6043
__e.Return(MakeNative(func(__e *ControlFlow) {
C6044 := __e.Get(1)
_ = C6044
tmp19751 := Call(__e, PrimFunc(symshen_4newpv), B6041)


let__18896 := tmp19751
_ = let__18896

tmp19752 := Call(__e, PrimFunc(symshen_4newpv), B6041)


let__18897 := tmp19752
_ = let__18897

tmp19753 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18897, Nil)
}
__typedArg0 := let__18897
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19754 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19753)
}
__typedArg0 := symlist
__typedArg1 := tmp19753
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19755 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19754, Nil)
}
__typedArg0 := tmp19754
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19756 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19755)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19755
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19757 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18896, tmp19756)
}
__typedArg0 := let__18896
__typedArg1 := tmp19756
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19758 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19759 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19758)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19758
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19760 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19757, tmp19759)
}
__typedArg0 := tmp19757
__typedArg1 := tmp19759
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19761 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19760, Nil)
}
__typedArg0 := tmp19760
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19762 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19761)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19761
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19763 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp19762)
}
__typedArg0 := symsymbol
__typedArg1 := tmp19762
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19764 := Call(__e, PrimFunc(symis_b), V6045, tmp19763, B6041, L6042, Key6043, C6044)


tmp19765 := Call(__e, PrimFunc(symshen_4gc), B6041, tmp19764)


__e.TailApply(PrimFunc(symshen_4gc), B6041, tmp19765)
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

tmp19766 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp19767 := Call(__e, PrimFunc(symshen_4assoc_1_6), symshen_4register_1source_1form, tmp19750, tmp19766)


tmp19768 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp19767)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp19767
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19768

tmp19769 := MakeNative(func(__e *ControlFlow) {
V6050 := __e.Get(1)
_ = V6050
__e.Return(MakeNative(func(__e *ControlFlow) {
B6046 := __e.Get(1)
_ = B6046
__e.Return(MakeNative(func(__e *ControlFlow) {
L6047 := __e.Get(1)
_ = L6047
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6048 := __e.Get(1)
_ = Key6048
__e.Return(MakeNative(func(__e *ControlFlow) {
C6049 := __e.Get(1)
_ = C6049
tmp19770 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19771 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19770)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19770
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19772 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp19771)
}
__typedArg0 := symsymbol
__typedArg1 := tmp19771
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6050, tmp19772, B6046, L6047, Key6048, C6049)
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

tmp19773 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp19774 := Call(__e, PrimFunc(symshen_4assoc_1_6), symshen_4unregister_1source_1form, tmp19769, tmp19773)


tmp19775 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp19774)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp19774
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19775

tmp19776 := MakeNative(func(__e *ControlFlow) {
V6055 := __e.Get(1)
_ = V6055
__e.Return(MakeNative(func(__e *ControlFlow) {
B6051 := __e.Get(1)
_ = B6051
__e.Return(MakeNative(func(__e *ControlFlow) {
L6052 := __e.Get(1)
_ = L6052
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6053 := __e.Get(1)
_ = Key6053
__e.Return(MakeNative(func(__e *ControlFlow) {
C6054 := __e.Get(1)
_ = C6054
tmp19777 := Call(__e, PrimFunc(symshen_4newpv), B6051)


let__18898 := tmp19777
_ = let__18898

tmp19778 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18898, Nil)
}
__typedArg0 := let__18898
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19779 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19778)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19778
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19780 := Call(__e, PrimFunc(symis_b), V6055, tmp19779, B6051, L6052, Key6053, C6054)


__e.TailApply(PrimFunc(symshen_4gc), B6051, tmp19780)
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

tmp19781 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp19782 := Call(__e, PrimFunc(symshen_4assoc_1_6), symabort, tmp19776, tmp19781)


tmp19783 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp19782)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp19782
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19783

tmp19784 := MakeNative(func(__e *ControlFlow) {
V6060 := __e.Get(1)
_ = V6060
__e.Return(MakeNative(func(__e *ControlFlow) {
B6056 := __e.Get(1)
_ = B6056
__e.Return(MakeNative(func(__e *ControlFlow) {
L6057 := __e.Get(1)
_ = L6057
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6058 := __e.Get(1)
_ = Key6058
__e.Return(MakeNative(func(__e *ControlFlow) {
C6059 := __e.Get(1)
_ = C6059
tmp19785 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19786 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19785)
}
__typedArg0 := symlist
__typedArg1 := tmp19785
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19787 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19786, Nil)
}
__typedArg0 := tmp19786
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19788 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19787)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19787
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19789 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp19788)
}
__typedArg0 := symstring
__typedArg1 := tmp19788
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6060, tmp19789, B6056, L6057, Key6058, C6059)
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

tmp19790 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp19791 := Call(__e, PrimFunc(symshen_4assoc_1_6), symabsolute, tmp19784, tmp19790)


tmp19792 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp19791)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp19791
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19792

tmp19793 := MakeNative(func(__e *ControlFlow) {
V6065 := __e.Get(1)
_ = V6065
__e.Return(MakeNative(func(__e *ControlFlow) {
B6061 := __e.Get(1)
_ = B6061
__e.Return(MakeNative(func(__e *ControlFlow) {
L6062 := __e.Get(1)
_ = L6062
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6063 := __e.Get(1)
_ = Key6063
__e.Return(MakeNative(func(__e *ControlFlow) {
C6064 := __e.Get(1)
_ = C6064
tmp19794 := Call(__e, PrimFunc(symshen_4newpv), B6061)


let__18899 := tmp19794
_ = let__18899

tmp19795 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19796 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19795)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19795
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19797 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18899, tmp19796)
}
__typedArg0 := let__18899
__typedArg1 := tmp19796
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19798 := Call(__e, PrimFunc(symis_b), V6065, tmp19797, B6061, L6062, Key6063, C6064)


__e.TailApply(PrimFunc(symshen_4gc), B6061, tmp19798)
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

tmp19800 := Call(__e, PrimFunc(symshen_4assoc_1_6), symabsvector_2, tmp19793, tmp19799)


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
V6071 := __e.Get(1)
_ = V6071
__e.Return(MakeNative(func(__e *ControlFlow) {
B6067 := __e.Get(1)
_ = B6067
__e.Return(MakeNative(func(__e *ControlFlow) {
L6068 := __e.Get(1)
_ = L6068
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6069 := __e.Get(1)
_ = Key6069
__e.Return(MakeNative(func(__e *ControlFlow) {
C6070 := __e.Get(1)
_ = C6070
tmp19803 := Call(__e, PrimFunc(symshen_4newpv), B6067)


let__18900 := tmp19803
_ = let__18900

tmp19804 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18900, Nil)
}
__typedArg0 := let__18900
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19805 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19804)
}
__typedArg0 := symlist
__typedArg1 := tmp19804
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19806 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18900, Nil)
}
__typedArg0 := let__18900
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19807 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19806)
}
__typedArg0 := symlist
__typedArg1 := tmp19806
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19808 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19807, Nil)
}
__typedArg0 := tmp19807
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19809 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19808)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19808
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19810 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19805, tmp19809)
}
__typedArg0 := tmp19805
__typedArg1 := tmp19809
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19811 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19810, Nil)
}
__typedArg0 := tmp19810
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

tmp19813 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18900, tmp19812)
}
__typedArg0 := let__18900
__typedArg1 := tmp19812
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19814 := Call(__e, PrimFunc(symis_b), V6071, tmp19813, B6067, L6068, Key6069, C6070)


__e.TailApply(PrimFunc(symshen_4gc), B6067, tmp19814)
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

tmp19815 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp19816 := Call(__e, PrimFunc(symshen_4assoc_1_6), symadjoin, tmp19802, tmp19815)


tmp19817 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp19816)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp19816
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19817

tmp19818 := MakeNative(func(__e *ControlFlow) {
V6076 := __e.Get(1)
_ = V6076
__e.Return(MakeNative(func(__e *ControlFlow) {
B6072 := __e.Get(1)
_ = B6072
__e.Return(MakeNative(func(__e *ControlFlow) {
L6073 := __e.Get(1)
_ = L6073
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6074 := __e.Get(1)
_ = Key6074
__e.Return(MakeNative(func(__e *ControlFlow) {
C6075 := __e.Get(1)
_ = C6075
tmp19819 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19820 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19819)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19819
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19821 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, tmp19820)
}
__typedArg0 := symboolean
__typedArg1 := tmp19820
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19822 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19821, Nil)
}
__typedArg0 := tmp19821
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19823 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19822)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19822
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19824 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, tmp19823)
}
__typedArg0 := symboolean
__typedArg1 := tmp19823
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6076, tmp19824, B6072, L6073, Key6074, C6075)
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

tmp19825 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp19826 := Call(__e, PrimFunc(symshen_4assoc_1_6), symand, tmp19818, tmp19825)


tmp19827 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp19826)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp19826
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19827

tmp19828 := MakeNative(func(__e *ControlFlow) {
V6081 := __e.Get(1)
_ = V6081
__e.Return(MakeNative(func(__e *ControlFlow) {
B6077 := __e.Get(1)
_ = B6077
__e.Return(MakeNative(func(__e *ControlFlow) {
L6078 := __e.Get(1)
_ = L6078
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6079 := __e.Get(1)
_ = Key6079
__e.Return(MakeNative(func(__e *ControlFlow) {
C6080 := __e.Get(1)
_ = C6080
tmp19829 := Call(__e, PrimFunc(symshen_4newpv), B6077)


let__18901 := tmp19829
_ = let__18901

tmp19830 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19831 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19830)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19830
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19832 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp19831)
}
__typedArg0 := symsymbol
__typedArg1 := tmp19831
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19833 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19832, Nil)
}
__typedArg0 := tmp19832
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19834 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19833)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19833
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19835 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp19834)
}
__typedArg0 := symstring
__typedArg1 := tmp19834
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19836 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19835, Nil)
}
__typedArg0 := tmp19835
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19837 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19836)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19836
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19838 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18901, tmp19837)
}
__typedArg0 := let__18901
__typedArg1 := tmp19837
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19839 := Call(__e, PrimFunc(symis_b), V6081, tmp19838, B6077, L6078, Key6079, C6080)


__e.TailApply(PrimFunc(symshen_4gc), B6077, tmp19839)
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

tmp19840 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp19841 := Call(__e, PrimFunc(symshen_4assoc_1_6), symshen_4app, tmp19828, tmp19840)


tmp19842 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp19841)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp19841
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19842

tmp19843 := MakeNative(func(__e *ControlFlow) {
V6088 := __e.Get(1)
_ = V6088
__e.Return(MakeNative(func(__e *ControlFlow) {
B6084 := __e.Get(1)
_ = B6084
__e.Return(MakeNative(func(__e *ControlFlow) {
L6085 := __e.Get(1)
_ = L6085
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6086 := __e.Get(1)
_ = Key6086
__e.Return(MakeNative(func(__e *ControlFlow) {
C6087 := __e.Get(1)
_ = C6087
tmp19844 := Call(__e, PrimFunc(symshen_4newpv), B6084)


let__18902 := tmp19844
_ = let__18902

tmp19845 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18902, Nil)
}
__typedArg0 := let__18902
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19846 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19845)
}
__typedArg0 := symlist
__typedArg1 := tmp19845
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19847 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18902, Nil)
}
__typedArg0 := let__18902
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19848 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19847)
}
__typedArg0 := symlist
__typedArg1 := tmp19847
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19849 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18902, Nil)
}
__typedArg0 := let__18902
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19850 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19849)
}
__typedArg0 := symlist
__typedArg1 := tmp19849
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19851 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19850, Nil)
}
__typedArg0 := tmp19850
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19852 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19851)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19851
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19853 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19848, tmp19852)
}
__typedArg0 := tmp19848
__typedArg1 := tmp19852
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19854 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19853, Nil)
}
__typedArg0 := tmp19853
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19855 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19854)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19854
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19856 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19846, tmp19855)
}
__typedArg0 := tmp19846
__typedArg1 := tmp19855
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19857 := Call(__e, PrimFunc(symis_b), V6088, tmp19856, B6084, L6085, Key6086, C6087)


__e.TailApply(PrimFunc(symshen_4gc), B6084, tmp19857)
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

tmp19859 := Call(__e, PrimFunc(symshen_4assoc_1_6), symappend, tmp19843, tmp19858)


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
V6093 := __e.Get(1)
_ = V6093
__e.Return(MakeNative(func(__e *ControlFlow) {
B6089 := __e.Get(1)
_ = B6089
__e.Return(MakeNative(func(__e *ControlFlow) {
L6090 := __e.Get(1)
_ = L6090
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6091 := __e.Get(1)
_ = Key6091
__e.Return(MakeNative(func(__e *ControlFlow) {
C6092 := __e.Get(1)
_ = C6092
tmp19862 := Call(__e, PrimFunc(symshen_4newpv), B6089)


let__18903 := tmp19862
_ = let__18903

tmp19863 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
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
return PrimCons(let__18903, tmp19864)
}
__typedArg0 := let__18903
__typedArg1 := tmp19864
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19866 := Call(__e, PrimFunc(symis_b), V6093, tmp19865, B6089, L6090, Key6091, C6092)


__e.TailApply(PrimFunc(symshen_4gc), B6089, tmp19866)
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

tmp19867 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp19868 := Call(__e, PrimFunc(symshen_4assoc_1_6), symarity, tmp19861, tmp19867)


tmp19869 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp19868)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp19868
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19869

tmp19870 := MakeNative(func(__e *ControlFlow) {
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
tmp19871 := Call(__e, PrimFunc(symshen_4newpv), B6095)


let__18904 := tmp19871
_ = let__18904

tmp19872 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18904, Nil)
}
__typedArg0 := let__18904
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19873 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19872)
}
__typedArg0 := symlist
__typedArg1 := tmp19872
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19874 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19873, Nil)
}
__typedArg0 := tmp19873
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19875 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19874)
}
__typedArg0 := symlist
__typedArg1 := tmp19874
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19876 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18904, Nil)
}
__typedArg0 := let__18904
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19877 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19876)
}
__typedArg0 := symlist
__typedArg1 := tmp19876
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19878 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19877, Nil)
}
__typedArg0 := tmp19877
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19879 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19878)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19878
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19880 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19875, tmp19879)
}
__typedArg0 := tmp19875
__typedArg1 := tmp19879
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19881 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19880, Nil)
}
__typedArg0 := tmp19880
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19882 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19881)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19881
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19883 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18904, tmp19882)
}
__typedArg0 := let__18904
__typedArg1 := tmp19882
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19884 := Call(__e, PrimFunc(symis_b), V6099, tmp19883, B6095, L6096, Key6097, C6098)


__e.TailApply(PrimFunc(symshen_4gc), B6095, tmp19884)
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

tmp19885 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp19886 := Call(__e, PrimFunc(symshen_4assoc_1_6), symassoc, tmp19870, tmp19885)


tmp19887 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp19886)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp19886
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19887

tmp19888 := MakeNative(func(__e *ControlFlow) {
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
tmp19889 := Call(__e, PrimFunc(symshen_4newpv), B6100)


let__18905 := tmp19889
_ = let__18905

tmp19890 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19891 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19890)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19890
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19892 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18905, tmp19891)
}
__typedArg0 := let__18905
__typedArg1 := tmp19891
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19893 := Call(__e, PrimFunc(symis_b), V6104, tmp19892, B6100, L6101, Key6102, C6103)


__e.TailApply(PrimFunc(symshen_4gc), B6100, tmp19893)
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

tmp19894 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp19895 := Call(__e, PrimFunc(symshen_4assoc_1_6), symatom_2, tmp19888, tmp19894)


tmp19896 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp19895)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp19895
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19896

tmp19897 := MakeNative(func(__e *ControlFlow) {
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
tmp19898 := Call(__e, PrimFunc(symshen_4newpv), B6105)


let__18906 := tmp19898
_ = let__18906

tmp19899 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19900 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19899)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19899
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19901 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18906, tmp19900)
}
__typedArg0 := let__18906
__typedArg1 := tmp19900
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19902 := Call(__e, PrimFunc(symis_b), V6109, tmp19901, B6105, L6106, Key6107, C6108)


__e.TailApply(PrimFunc(symshen_4gc), B6105, tmp19902)
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

tmp19903 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp19904 := Call(__e, PrimFunc(symshen_4assoc_1_6), symboolean_2, tmp19897, tmp19903)


tmp19905 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp19904)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp19904
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19905

tmp19906 := MakeNative(func(__e *ControlFlow) {
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
tmp19907 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19908 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19907)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19907
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19909 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp19908)
}
__typedArg0 := symstring
__typedArg1 := tmp19908
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6114, tmp19909, B6110, L6111, Key6112, C6113)
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

tmp19910 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp19911 := Call(__e, PrimFunc(symshen_4assoc_1_6), symbootstrap, tmp19906, tmp19910)


tmp19912 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp19911)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp19911
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19912

tmp19913 := MakeNative(func(__e *ControlFlow) {
V6119 := __e.Get(1)
_ = V6119
__e.Return(MakeNative(func(__e *ControlFlow) {
B6115 := __e.Get(1)
_ = B6115
__e.Return(MakeNative(func(__e *ControlFlow) {
L6116 := __e.Get(1)
_ = L6116
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6117 := __e.Get(1)
_ = Key6117
__e.Return(MakeNative(func(__e *ControlFlow) {
C6118 := __e.Get(1)
_ = C6118
tmp19914 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
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
return PrimCons(symsymbol, tmp19915)
}
__typedArg0 := symsymbol
__typedArg1 := tmp19915
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6119, tmp19916, B6115, L6116, Key6117, C6118)
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

tmp19917 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp19918 := Call(__e, PrimFunc(symshen_4assoc_1_6), symbound_2, tmp19913, tmp19917)


tmp19919 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp19918)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp19918
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19919

tmp19920 := MakeNative(func(__e *ControlFlow) {
V6124 := __e.Get(1)
_ = V6124
__e.Return(MakeNative(func(__e *ControlFlow) {
B6120 := __e.Get(1)
_ = B6120
__e.Return(MakeNative(func(__e *ControlFlow) {
L6121 := __e.Get(1)
_ = L6121
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6122 := __e.Get(1)
_ = Key6122
__e.Return(MakeNative(func(__e *ControlFlow) {
C6123 := __e.Get(1)
_ = C6123
tmp19921 := Call(__e, PrimFunc(symshen_4newpv), B6120)


let__18907 := tmp19921
_ = let__18907

tmp19922 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18907, Nil)
}
__typedArg0 := let__18907
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19923 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19922)
}
__typedArg0 := symlist
__typedArg1 := tmp19922
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19924 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19925 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19924)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19924
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19926 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19923, tmp19925)
}
__typedArg0 := tmp19923
__typedArg1 := tmp19925
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19927 := Call(__e, PrimFunc(symis_b), V6124, tmp19926, B6120, L6121, Key6122, C6123)


__e.TailApply(PrimFunc(symshen_4gc), B6120, tmp19927)
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

tmp19928 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp19929 := Call(__e, PrimFunc(symshen_4assoc_1_6), symshen_4ccons_2, tmp19920, tmp19928)


tmp19930 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp19929)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp19929
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19930

tmp19931 := MakeNative(func(__e *ControlFlow) {
V6129 := __e.Get(1)
_ = V6129
__e.Return(MakeNative(func(__e *ControlFlow) {
B6125 := __e.Get(1)
_ = B6125
__e.Return(MakeNative(func(__e *ControlFlow) {
L6126 := __e.Get(1)
_ = L6126
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6127 := __e.Get(1)
_ = Key6127
__e.Return(MakeNative(func(__e *ControlFlow) {
C6128 := __e.Get(1)
_ = C6128
tmp19932 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
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
return PrimCons(symstring, tmp19933)
}
__typedArg0 := symstring
__typedArg1 := tmp19933
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6129, tmp19934, B6125, L6126, Key6127, C6128)
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

tmp19935 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp19936 := Call(__e, PrimFunc(symshen_4assoc_1_6), symcd, tmp19931, tmp19935)


tmp19937 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp19936)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp19936
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19937

tmp19938 := MakeNative(func(__e *ControlFlow) {
V6135 := __e.Get(1)
_ = V6135
__e.Return(MakeNative(func(__e *ControlFlow) {
B6131 := __e.Get(1)
_ = B6131
__e.Return(MakeNative(func(__e *ControlFlow) {
L6132 := __e.Get(1)
_ = L6132
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6133 := __e.Get(1)
_ = Key6133
__e.Return(MakeNative(func(__e *ControlFlow) {
C6134 := __e.Get(1)
_ = C6134
tmp19939 := Call(__e, PrimFunc(symshen_4newpv), B6131)


let__18908 := tmp19939
_ = let__18908

tmp19940 := Call(__e, PrimFunc(symshen_4newpv), B6131)


let__18909 := tmp19940
_ = let__18909

tmp19941 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18908, Nil)
}
__typedArg0 := let__18908
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19942 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp19941)
}
__typedArg0 := symstream
__typedArg1 := tmp19941
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19943 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18909, Nil)
}
__typedArg0 := let__18909
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19944 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19943)
}
__typedArg0 := symlist
__typedArg1 := tmp19943
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19945 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19944, Nil)
}
__typedArg0 := tmp19944
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19946 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19945)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19945
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19947 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19942, tmp19946)
}
__typedArg0 := tmp19942
__typedArg1 := tmp19946
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19948 := Call(__e, PrimFunc(symis_b), V6135, tmp19947, B6131, L6132, Key6133, C6134)


tmp19949 := Call(__e, PrimFunc(symshen_4gc), B6131, tmp19948)


__e.TailApply(PrimFunc(symshen_4gc), B6131, tmp19949)
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

tmp19951 := Call(__e, PrimFunc(symshen_4assoc_1_6), symclose, tmp19938, tmp19950)


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
V6140 := __e.Get(1)
_ = V6140
__e.Return(MakeNative(func(__e *ControlFlow) {
B6136 := __e.Get(1)
_ = B6136
__e.Return(MakeNative(func(__e *ControlFlow) {
L6137 := __e.Get(1)
_ = L6137
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6138 := __e.Get(1)
_ = Key6138
__e.Return(MakeNative(func(__e *ControlFlow) {
C6139 := __e.Get(1)
_ = C6139
tmp19954 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19955 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19954)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19954
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19956 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp19955)
}
__typedArg0 := symstring
__typedArg1 := tmp19955
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19957 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19956, Nil)
}
__typedArg0 := tmp19956
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
return PrimCons(symstring, tmp19958)
}
__typedArg0 := symstring
__typedArg1 := tmp19958
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6140, tmp19959, B6136, L6137, Key6138, C6139)
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

tmp19960 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp19961 := Call(__e, PrimFunc(symshen_4assoc_1_6), symcn, tmp19953, tmp19960)


tmp19962 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp19961)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp19961
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19962

tmp19963 := MakeNative(func(__e *ControlFlow) {
V6145 := __e.Get(1)
_ = V6145
__e.Return(MakeNative(func(__e *ControlFlow) {
B6141 := __e.Get(1)
_ = B6141
__e.Return(MakeNative(func(__e *ControlFlow) {
L6142 := __e.Get(1)
_ = L6142
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6143 := __e.Get(1)
_ = Key6143
__e.Return(MakeNative(func(__e *ControlFlow) {
C6144 := __e.Get(1)
_ = C6144
tmp19964 := Call(__e, PrimFunc(symshen_4newpv), B6141)


let__18910 := tmp19964
_ = let__18910

tmp19965 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18910, Nil)
}
__typedArg0 := let__18910
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19966 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp19965)
}
__typedArg0 := symstream
__typedArg1 := tmp19965
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19967 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19968 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19967)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19967
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19969 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19966, tmp19968)
}
__typedArg0 := tmp19966
__typedArg1 := tmp19968
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19970 := Call(__e, PrimFunc(symis_b), V6145, tmp19969, B6141, L6142, Key6143, C6144)


__e.TailApply(PrimFunc(symshen_4gc), B6141, tmp19970)
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

tmp19971 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp19972 := Call(__e, PrimFunc(symshen_4assoc_1_6), symchar_1stream_2, tmp19963, tmp19971)


tmp19973 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp19972)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp19972
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19973

tmp19974 := MakeNative(func(__e *ControlFlow) {
V6152 := __e.Get(1)
_ = V6152
__e.Return(MakeNative(func(__e *ControlFlow) {
B6148 := __e.Get(1)
_ = B6148
__e.Return(MakeNative(func(__e *ControlFlow) {
L6149 := __e.Get(1)
_ = L6149
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6150 := __e.Get(1)
_ = Key6150
__e.Return(MakeNative(func(__e *ControlFlow) {
C6151 := __e.Get(1)
_ = C6151
tmp19975 := Call(__e, PrimFunc(symshen_4newpv), B6148)


let__18911 := tmp19975
_ = let__18911

tmp19976 := Call(__e, PrimFunc(symshen_4newpv), B6148)


let__18912 := tmp19976
_ = let__18912

tmp19977 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18911, Nil)
}
__typedArg0 := let__18911
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19978 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19977)
}
__typedArg0 := symlist
__typedArg1 := tmp19977
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19979 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18911, Nil)
}
__typedArg0 := let__18911
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19980 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19979)
}
__typedArg0 := symlist
__typedArg1 := tmp19979
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19981 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18912, Nil)
}
__typedArg0 := let__18912
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19982 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19980, tmp19981)
}
__typedArg0 := tmp19980
__typedArg1 := tmp19981
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19983 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstr, tmp19982)
}
__typedArg0 := symstr
__typedArg1 := tmp19982
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19984 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19983, Nil)
}
__typedArg0 := tmp19983
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19985 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19984)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19984
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19986 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19978, tmp19985)
}
__typedArg0 := tmp19978
__typedArg1 := tmp19985
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19987 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18911, Nil)
}
__typedArg0 := let__18911
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19988 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19987)
}
__typedArg0 := symlist
__typedArg1 := tmp19987
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19989 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18912, Nil)
}
__typedArg0 := let__18912
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19990 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19989)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19989
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19991 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19988, tmp19990)
}
__typedArg0 := tmp19988
__typedArg1 := tmp19990
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19992 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19991, Nil)
}
__typedArg0 := tmp19991
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19993 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19992)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19992
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19994 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19986, tmp19993)
}
__typedArg0 := tmp19986
__typedArg1 := tmp19993
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19995 := Call(__e, PrimFunc(symis_b), V6152, tmp19994, B6148, L6149, Key6150, C6151)


tmp19996 := Call(__e, PrimFunc(symshen_4gc), B6148, tmp19995)


__e.TailApply(PrimFunc(symshen_4gc), B6148, tmp19996)
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

tmp19997 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp19998 := Call(__e, PrimFunc(symshen_4assoc_1_6), symcompile, tmp19974, tmp19997)


tmp19999 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp19998)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp19998
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19999

tmp20000 := MakeNative(func(__e *ControlFlow) {
V6157 := __e.Get(1)
_ = V6157
__e.Return(MakeNative(func(__e *ControlFlow) {
B6153 := __e.Get(1)
_ = B6153
__e.Return(MakeNative(func(__e *ControlFlow) {
L6154 := __e.Get(1)
_ = L6154
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6155 := __e.Get(1)
_ = Key6155
__e.Return(MakeNative(func(__e *ControlFlow) {
C6156 := __e.Get(1)
_ = C6156
tmp20001 := Call(__e, PrimFunc(symshen_4newpv), B6153)


let__18913 := tmp20001
_ = let__18913

tmp20002 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20003 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20002)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20002
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20004 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18913, tmp20003)
}
__typedArg0 := let__18913
__typedArg1 := tmp20003
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20005 := Call(__e, PrimFunc(symis_b), V6157, tmp20004, B6153, L6154, Key6155, C6156)


__e.TailApply(PrimFunc(symshen_4gc), B6153, tmp20005)
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

tmp20006 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20007 := Call(__e, PrimFunc(symshen_4assoc_1_6), symcons_2, tmp20000, tmp20006)


tmp20008 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20007)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20007
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20008

tmp20009 := MakeNative(func(__e *ControlFlow) {
V6162 := __e.Get(1)
_ = V6162
__e.Return(MakeNative(func(__e *ControlFlow) {
B6158 := __e.Get(1)
_ = B6158
__e.Return(MakeNative(func(__e *ControlFlow) {
L6159 := __e.Get(1)
_ = L6159
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6160 := __e.Get(1)
_ = Key6160
__e.Return(MakeNative(func(__e *ControlFlow) {
C6161 := __e.Get(1)
_ = C6161
tmp20010 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
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
return PrimCons(tmp20011, Nil)
}
__typedArg0 := tmp20011
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20013 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20012)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20012
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6162, tmp20013, B6158, L6159, Key6160, C6161)
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

tmp20014 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20015 := Call(__e, PrimFunc(symshen_4assoc_1_6), symdatatypes, tmp20009, tmp20014)


tmp20016 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20015)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20015
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20016

tmp20017 := MakeNative(func(__e *ControlFlow) {
V6167 := __e.Get(1)
_ = V6167
__e.Return(MakeNative(func(__e *ControlFlow) {
B6163 := __e.Get(1)
_ = B6163
__e.Return(MakeNative(func(__e *ControlFlow) {
L6164 := __e.Get(1)
_ = L6164
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6165 := __e.Get(1)
_ = Key6165
__e.Return(MakeNative(func(__e *ControlFlow) {
C6166 := __e.Get(1)
_ = C6166
tmp20018 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20019 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20018)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20018
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20020 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp20019)
}
__typedArg0 := symsymbol
__typedArg1 := tmp20019
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6167, tmp20020, B6163, L6164, Key6165, C6166)
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

tmp20021 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20022 := Call(__e, PrimFunc(symshen_4assoc_1_6), symdestroy, tmp20017, tmp20021)


tmp20023 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20022)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20022
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20023

tmp20024 := MakeNative(func(__e *ControlFlow) {
V6174 := __e.Get(1)
_ = V6174
__e.Return(MakeNative(func(__e *ControlFlow) {
B6170 := __e.Get(1)
_ = B6170
__e.Return(MakeNative(func(__e *ControlFlow) {
L6171 := __e.Get(1)
_ = L6171
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6172 := __e.Get(1)
_ = Key6172
__e.Return(MakeNative(func(__e *ControlFlow) {
C6173 := __e.Get(1)
_ = C6173
tmp20025 := Call(__e, PrimFunc(symshen_4newpv), B6170)


let__18914 := tmp20025
_ = let__18914

tmp20026 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18914, Nil)
}
__typedArg0 := let__18914
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20027 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20026)
}
__typedArg0 := symlist
__typedArg1 := tmp20026
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20028 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18914, Nil)
}
__typedArg0 := let__18914
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20029 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20028)
}
__typedArg0 := symlist
__typedArg1 := tmp20028
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20030 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18914, Nil)
}
__typedArg0 := let__18914
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20031 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20030)
}
__typedArg0 := symlist
__typedArg1 := tmp20030
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20032 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20031, Nil)
}
__typedArg0 := tmp20031
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20033 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20032)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20032
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20034 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20029, tmp20033)
}
__typedArg0 := tmp20029
__typedArg1 := tmp20033
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20035 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20034, Nil)
}
__typedArg0 := tmp20034
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
return PrimCons(tmp20027, tmp20036)
}
__typedArg0 := tmp20027
__typedArg1 := tmp20036
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20038 := Call(__e, PrimFunc(symis_b), V6174, tmp20037, B6170, L6171, Key6172, C6173)


__e.TailApply(PrimFunc(symshen_4gc), B6170, tmp20038)
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

tmp20040 := Call(__e, PrimFunc(symshen_4assoc_1_6), symdifference, tmp20024, tmp20039)


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
V6179 := __e.Get(1)
_ = V6179
__e.Return(MakeNative(func(__e *ControlFlow) {
B6175 := __e.Get(1)
_ = B6175
__e.Return(MakeNative(func(__e *ControlFlow) {
L6176 := __e.Get(1)
_ = L6176
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6177 := __e.Get(1)
_ = Key6177
__e.Return(MakeNative(func(__e *ControlFlow) {
C6178 := __e.Get(1)
_ = C6178
tmp20043 := Call(__e, PrimFunc(symshen_4newpv), B6175)


let__18915 := tmp20043
_ = let__18915

tmp20044 := Call(__e, PrimFunc(symshen_4newpv), B6175)


let__18916 := tmp20044
_ = let__18916

tmp20045 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18916, Nil)
}
__typedArg0 := let__18916
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

tmp20047 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18916, tmp20046)
}
__typedArg0 := let__18916
__typedArg1 := tmp20046
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20048 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20047, Nil)
}
__typedArg0 := tmp20047
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20049 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20048)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20048
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20050 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18915, tmp20049)
}
__typedArg0 := let__18915
__typedArg1 := tmp20049
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20051 := Call(__e, PrimFunc(symis_b), V6179, tmp20050, B6175, L6176, Key6177, C6178)


tmp20052 := Call(__e, PrimFunc(symshen_4gc), B6175, tmp20051)


__e.TailApply(PrimFunc(symshen_4gc), B6175, tmp20052)
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

tmp20053 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20054 := Call(__e, PrimFunc(symshen_4assoc_1_6), symdo, tmp20042, tmp20053)


tmp20055 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20054)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20054
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20055

tmp20056 := MakeNative(func(__e *ControlFlow) {
V6186 := __e.Get(1)
_ = V6186
__e.Return(MakeNative(func(__e *ControlFlow) {
B6182 := __e.Get(1)
_ = B6182
__e.Return(MakeNative(func(__e *ControlFlow) {
L6183 := __e.Get(1)
_ = L6183
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6184 := __e.Get(1)
_ = Key6184
__e.Return(MakeNative(func(__e *ControlFlow) {
C6185 := __e.Get(1)
_ = C6185
tmp20057 := Call(__e, PrimFunc(symshen_4newpv), B6182)


let__18917 := tmp20057
_ = let__18917

tmp20058 := Call(__e, PrimFunc(symshen_4newpv), B6182)


let__18918 := tmp20058
_ = let__18918

tmp20059 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18917, Nil)
}
__typedArg0 := let__18917
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
return PrimCons(let__18917, Nil)
}
__typedArg0 := let__18917
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
return PrimCons(let__18918, Nil)
}
__typedArg0 := let__18918
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
return PrimCons(tmp20062, tmp20065)
}
__typedArg0 := tmp20062
__typedArg1 := tmp20065
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20067 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstr, tmp20066)
}
__typedArg0 := symstr
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

tmp20071 := Call(__e, PrimFunc(symis_b), V6186, tmp20070, B6182, L6183, Key6184, C6185)


tmp20072 := Call(__e, PrimFunc(symshen_4gc), B6182, tmp20071)


__e.TailApply(PrimFunc(symshen_4gc), B6182, tmp20072)
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

tmp20073 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20074 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_5e_6, tmp20056, tmp20073)


tmp20075 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20074)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20074
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20075

tmp20076 := MakeNative(func(__e *ControlFlow) {
V6193 := __e.Get(1)
_ = V6193
__e.Return(MakeNative(func(__e *ControlFlow) {
B6189 := __e.Get(1)
_ = B6189
__e.Return(MakeNative(func(__e *ControlFlow) {
L6190 := __e.Get(1)
_ = L6190
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6191 := __e.Get(1)
_ = Key6191
__e.Return(MakeNative(func(__e *ControlFlow) {
C6192 := __e.Get(1)
_ = C6192
tmp20077 := Call(__e, PrimFunc(symshen_4newpv), B6189)


let__18919 := tmp20077
_ = let__18919

tmp20078 := Call(__e, PrimFunc(symshen_4newpv), B6189)


let__18920 := tmp20078
_ = let__18920

tmp20079 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18920, Nil)
}
__typedArg0 := let__18920
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20080 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20079)
}
__typedArg0 := symlist
__typedArg1 := tmp20079
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20081 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18919, Nil)
}
__typedArg0 := let__18919
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20082 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20081)
}
__typedArg0 := symlist
__typedArg1 := tmp20081
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20083 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18920, Nil)
}
__typedArg0 := let__18920
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20084 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20083)
}
__typedArg0 := symlist
__typedArg1 := tmp20083
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20085 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20084, Nil)
}
__typedArg0 := tmp20084
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20086 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20082, tmp20085)
}
__typedArg0 := tmp20082
__typedArg1 := tmp20085
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20087 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstr, tmp20086)
}
__typedArg0 := symstr
__typedArg1 := tmp20086
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20088 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20087, Nil)
}
__typedArg0 := tmp20087
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20089 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20088)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20088
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20090 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20080, tmp20089)
}
__typedArg0 := tmp20080
__typedArg1 := tmp20089
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20091 := Call(__e, PrimFunc(symis_b), V6193, tmp20090, B6189, L6190, Key6191, C6192)


tmp20092 := Call(__e, PrimFunc(symshen_4gc), B6189, tmp20091)


__e.TailApply(PrimFunc(symshen_4gc), B6189, tmp20092)
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

tmp20093 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20094 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_5_b_6, tmp20076, tmp20093)


tmp20095 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20094)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20094
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20095

tmp20096 := MakeNative(func(__e *ControlFlow) {
V6200 := __e.Get(1)
_ = V6200
__e.Return(MakeNative(func(__e *ControlFlow) {
B6196 := __e.Get(1)
_ = B6196
__e.Return(MakeNative(func(__e *ControlFlow) {
L6197 := __e.Get(1)
_ = L6197
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6198 := __e.Get(1)
_ = Key6198
__e.Return(MakeNative(func(__e *ControlFlow) {
C6199 := __e.Get(1)
_ = C6199
tmp20097 := Call(__e, PrimFunc(symshen_4newpv), B6196)


let__18921 := tmp20097
_ = let__18921

tmp20098 := Call(__e, PrimFunc(symshen_4newpv), B6196)


let__18922 := tmp20098
_ = let__18922

tmp20099 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18921, Nil)
}
__typedArg0 := let__18921
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20100 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20099)
}
__typedArg0 := symlist
__typedArg1 := tmp20099
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20101 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18921, Nil)
}
__typedArg0 := let__18921
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20102 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20101)
}
__typedArg0 := symlist
__typedArg1 := tmp20101
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20103 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18922, Nil)
}
__typedArg0 := let__18922
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20104 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20103)
}
__typedArg0 := symlist
__typedArg1 := tmp20103
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20105 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20104, Nil)
}
__typedArg0 := tmp20104
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20106 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20102, tmp20105)
}
__typedArg0 := tmp20102
__typedArg1 := tmp20105
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20107 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstr, tmp20106)
}
__typedArg0 := symstr
__typedArg1 := tmp20106
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20108 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20107, Nil)
}
__typedArg0 := tmp20107
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20109 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20108)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20108
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20110 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20100, tmp20109)
}
__typedArg0 := tmp20100
__typedArg1 := tmp20109
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20111 := Call(__e, PrimFunc(symis_b), V6200, tmp20110, B6196, L6197, Key6198, C6199)


tmp20112 := Call(__e, PrimFunc(symshen_4gc), B6196, tmp20111)


__e.TailApply(PrimFunc(symshen_4gc), B6196, tmp20112)
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

tmp20113 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20114 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_5end_6, tmp20096, tmp20113)


tmp20115 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20114)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20114
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20115

tmp20116 := MakeNative(func(__e *ControlFlow) {
V6205 := __e.Get(1)
_ = V6205
__e.Return(MakeNative(func(__e *ControlFlow) {
B6201 := __e.Get(1)
_ = B6201
__e.Return(MakeNative(func(__e *ControlFlow) {
L6202 := __e.Get(1)
_ = L6202
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6203 := __e.Get(1)
_ = Key6203
__e.Return(MakeNative(func(__e *ControlFlow) {
C6204 := __e.Get(1)
_ = C6204
tmp20117 := Call(__e, PrimFunc(symshen_4newpv), B6201)


let__18923 := tmp20117
_ = let__18923

tmp20118 := Call(__e, PrimFunc(symshen_4newpv), B6201)


let__18924 := tmp20118
_ = let__18924

tmp20119 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18923, Nil)
}
__typedArg0 := let__18923
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20120 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20119)
}
__typedArg0 := symlist
__typedArg1 := tmp20119
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20121 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18924, Nil)
}
__typedArg0 := let__18924
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20122 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20120, tmp20121)
}
__typedArg0 := tmp20120
__typedArg1 := tmp20121
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20123 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstr, tmp20122)
}
__typedArg0 := symstr
__typedArg1 := tmp20122
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20124 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20125 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20124)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20124
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20126 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20123, tmp20125)
}
__typedArg0 := tmp20123
__typedArg1 := tmp20125
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20127 := Call(__e, PrimFunc(symis_b), V6205, tmp20126, B6201, L6202, Key6203, C6204)


tmp20128 := Call(__e, PrimFunc(symshen_4gc), B6201, tmp20127)


__e.TailApply(PrimFunc(symshen_4gc), B6201, tmp20128)
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

tmp20129 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20130 := Call(__e, PrimFunc(symshen_4assoc_1_6), symshen_4parse_1failure_2, tmp20116, tmp20129)


tmp20131 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20130)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20130
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20131

tmp20132 := MakeNative(func(__e *ControlFlow) {
V6210 := __e.Get(1)
_ = V6210
__e.Return(MakeNative(func(__e *ControlFlow) {
B6206 := __e.Get(1)
_ = B6206
__e.Return(MakeNative(func(__e *ControlFlow) {
L6207 := __e.Get(1)
_ = L6207
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6208 := __e.Get(1)
_ = Key6208
__e.Return(MakeNative(func(__e *ControlFlow) {
C6209 := __e.Get(1)
_ = C6209
tmp20133 := Call(__e, PrimFunc(symshen_4newpv), B6206)


let__18925 := tmp20133
_ = let__18925

tmp20134 := Call(__e, PrimFunc(symshen_4newpv), B6206)


let__18926 := tmp20134
_ = let__18926

tmp20135 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18925, Nil)
}
__typedArg0 := let__18925
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20136 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20135)
}
__typedArg0 := symlist
__typedArg1 := tmp20135
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20137 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18926, Nil)
}
__typedArg0 := let__18926
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20138 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20136, tmp20137)
}
__typedArg0 := tmp20136
__typedArg1 := tmp20137
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20139 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstr, tmp20138)
}
__typedArg0 := symstr
__typedArg1 := tmp20138
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20140 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20139, Nil)
}
__typedArg0 := tmp20139
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20141 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20140)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20140
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20142 := Call(__e, PrimFunc(symis_b), V6210, tmp20141, B6206, L6207, Key6208, C6209)


tmp20143 := Call(__e, PrimFunc(symshen_4gc), B6206, tmp20142)


__e.TailApply(PrimFunc(symshen_4gc), B6206, tmp20143)
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

tmp20144 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20145 := Call(__e, PrimFunc(symshen_4assoc_1_6), symshen_4parse_1failure, tmp20132, tmp20144)


tmp20146 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20145)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20145
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20146

tmp20147 := MakeNative(func(__e *ControlFlow) {
V6215 := __e.Get(1)
_ = V6215
__e.Return(MakeNative(func(__e *ControlFlow) {
B6211 := __e.Get(1)
_ = B6211
__e.Return(MakeNative(func(__e *ControlFlow) {
L6212 := __e.Get(1)
_ = L6212
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6213 := __e.Get(1)
_ = Key6213
__e.Return(MakeNative(func(__e *ControlFlow) {
C6214 := __e.Get(1)
_ = C6214
tmp20148 := Call(__e, PrimFunc(symshen_4newpv), B6211)


let__18927 := tmp20148
_ = let__18927

tmp20149 := Call(__e, PrimFunc(symshen_4newpv), B6211)


let__18928 := tmp20149
_ = let__18928

tmp20150 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18927, Nil)
}
__typedArg0 := let__18927
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20151 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20150)
}
__typedArg0 := symlist
__typedArg1 := tmp20150
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20152 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18928, Nil)
}
__typedArg0 := let__18928
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20153 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20151, tmp20152)
}
__typedArg0 := tmp20151
__typedArg1 := tmp20152
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20154 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstr, tmp20153)
}
__typedArg0 := symstr
__typedArg1 := tmp20153
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20155 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18928, Nil)
}
__typedArg0 := let__18928
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20156 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20155)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20155
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20157 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20154, tmp20156)
}
__typedArg0 := tmp20154
__typedArg1 := tmp20156
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20158 := Call(__e, PrimFunc(symis_b), V6215, tmp20157, B6211, L6212, Key6213, C6214)


tmp20159 := Call(__e, PrimFunc(symshen_4gc), B6211, tmp20158)


__e.TailApply(PrimFunc(symshen_4gc), B6211, tmp20159)
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

tmp20160 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20161 := Call(__e, PrimFunc(symshen_4assoc_1_6), symshen_4_5_1out, tmp20147, tmp20160)


tmp20162 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20161)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20161
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20162

tmp20163 := MakeNative(func(__e *ControlFlow) {
V6221 := __e.Get(1)
_ = V6221
__e.Return(MakeNative(func(__e *ControlFlow) {
B6217 := __e.Get(1)
_ = B6217
__e.Return(MakeNative(func(__e *ControlFlow) {
L6218 := __e.Get(1)
_ = L6218
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6219 := __e.Get(1)
_ = Key6219
__e.Return(MakeNative(func(__e *ControlFlow) {
C6220 := __e.Get(1)
_ = C6220
tmp20164 := Call(__e, PrimFunc(symshen_4newpv), B6217)


let__18929 := tmp20164
_ = let__18929

tmp20165 := Call(__e, PrimFunc(symshen_4newpv), B6217)


let__18930 := tmp20165
_ = let__18930

tmp20166 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18930, Nil)
}
__typedArg0 := let__18930
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20167 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20166)
}
__typedArg0 := symlist
__typedArg1 := tmp20166
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20168 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18929, Nil)
}
__typedArg0 := let__18929
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20169 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20167, tmp20168)
}
__typedArg0 := tmp20167
__typedArg1 := tmp20168
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20170 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstr, tmp20169)
}
__typedArg0 := symstr
__typedArg1 := tmp20169
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20171 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18930, Nil)
}
__typedArg0 := let__18930
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20172 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20171)
}
__typedArg0 := symlist
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

tmp20175 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20170, tmp20174)
}
__typedArg0 := tmp20170
__typedArg1 := tmp20174
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20176 := Call(__e, PrimFunc(symis_b), V6221, tmp20175, B6217, L6218, Key6219, C6220)


tmp20177 := Call(__e, PrimFunc(symshen_4gc), B6217, tmp20176)


__e.TailApply(PrimFunc(symshen_4gc), B6217, tmp20177)
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

tmp20178 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20179 := Call(__e, PrimFunc(symshen_4assoc_1_6), symshen_4in_1_6, tmp20163, tmp20178)


tmp20180 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20179)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20179
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20180

tmp20181 := MakeNative(func(__e *ControlFlow) {
V6227 := __e.Get(1)
_ = V6227
__e.Return(MakeNative(func(__e *ControlFlow) {
B6223 := __e.Get(1)
_ = B6223
__e.Return(MakeNative(func(__e *ControlFlow) {
L6224 := __e.Get(1)
_ = L6224
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6225 := __e.Get(1)
_ = Key6225
__e.Return(MakeNative(func(__e *ControlFlow) {
C6226 := __e.Get(1)
_ = C6226
tmp20182 := Call(__e, PrimFunc(symshen_4newpv), B6223)


let__18931 := tmp20182
_ = let__18931

tmp20183 := Call(__e, PrimFunc(symshen_4newpv), B6223)


let__18932 := tmp20183
_ = let__18932

tmp20184 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18931, Nil)
}
__typedArg0 := let__18931
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20185 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20184)
}
__typedArg0 := symlist
__typedArg1 := tmp20184
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20186 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18931, Nil)
}
__typedArg0 := let__18931
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20187 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20186)
}
__typedArg0 := symlist
__typedArg1 := tmp20186
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20188 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18932, Nil)
}
__typedArg0 := let__18932
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20189 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20187, tmp20188)
}
__typedArg0 := tmp20187
__typedArg1 := tmp20188
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20190 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstr, tmp20189)
}
__typedArg0 := symstr
__typedArg1 := tmp20189
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20191 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20190, Nil)
}
__typedArg0 := tmp20190
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20192 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20191)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20191
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20193 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18932, tmp20192)
}
__typedArg0 := let__18932
__typedArg1 := tmp20192
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20194 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20193, Nil)
}
__typedArg0 := tmp20193
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20195 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20194)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20194
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20196 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20185, tmp20195)
}
__typedArg0 := tmp20185
__typedArg1 := tmp20195
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20197 := Call(__e, PrimFunc(symis_b), V6227, tmp20196, B6223, L6224, Key6225, C6226)


tmp20198 := Call(__e, PrimFunc(symshen_4gc), B6223, tmp20197)


__e.TailApply(PrimFunc(symshen_4gc), B6223, tmp20198)
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

tmp20199 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20200 := Call(__e, PrimFunc(symshen_4assoc_1_6), symshen_4comb, tmp20181, tmp20199)


tmp20201 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20200)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20200
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20201

tmp20202 := MakeNative(func(__e *ControlFlow) {
V6232 := __e.Get(1)
_ = V6232
__e.Return(MakeNative(func(__e *ControlFlow) {
B6228 := __e.Get(1)
_ = B6228
__e.Return(MakeNative(func(__e *ControlFlow) {
L6229 := __e.Get(1)
_ = L6229
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6230 := __e.Get(1)
_ = Key6230
__e.Return(MakeNative(func(__e *ControlFlow) {
C6231 := __e.Get(1)
_ = C6231
tmp20203 := Call(__e, PrimFunc(symshen_4newpv), B6228)


let__18933 := tmp20203
_ = let__18933

tmp20204 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18933, Nil)
}
__typedArg0 := let__18933
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
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
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
return PrimCons(tmp20205, tmp20207)
}
__typedArg0 := tmp20205
__typedArg1 := tmp20207
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20209 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20208, Nil)
}
__typedArg0 := tmp20208
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20210 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20209)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20209
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20211 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18933, tmp20210)
}
__typedArg0 := let__18933
__typedArg1 := tmp20210
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20212 := Call(__e, PrimFunc(symis_b), V6232, tmp20211, B6228, L6229, Key6230, C6231)


__e.TailApply(PrimFunc(symshen_4gc), B6228, tmp20212)
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

tmp20213 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20214 := Call(__e, PrimFunc(symshen_4assoc_1_6), symelement_2, tmp20202, tmp20213)


tmp20215 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20214)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20214
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20215

tmp20216 := MakeNative(func(__e *ControlFlow) {
V6237 := __e.Get(1)
_ = V6237
__e.Return(MakeNative(func(__e *ControlFlow) {
B6233 := __e.Get(1)
_ = B6233
__e.Return(MakeNative(func(__e *ControlFlow) {
L6234 := __e.Get(1)
_ = L6234
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6235 := __e.Get(1)
_ = Key6235
__e.Return(MakeNative(func(__e *ControlFlow) {
C6236 := __e.Get(1)
_ = C6236
tmp20217 := Call(__e, PrimFunc(symshen_4newpv), B6233)


let__18934 := tmp20217
_ = let__18934

tmp20218 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20219 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20218)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20218
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20220 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18934, tmp20219)
}
__typedArg0 := let__18934
__typedArg1 := tmp20219
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20221 := Call(__e, PrimFunc(symis_b), V6237, tmp20220, B6233, L6234, Key6235, C6236)


__e.TailApply(PrimFunc(symshen_4gc), B6233, tmp20221)
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

tmp20222 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20223 := Call(__e, PrimFunc(symshen_4assoc_1_6), symempty_2, tmp20216, tmp20222)


tmp20224 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20223)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20223
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20224

tmp20225 := MakeNative(func(__e *ControlFlow) {
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
tmp20226 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20227 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20226)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20226
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20228 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp20227)
}
__typedArg0 := symsymbol
__typedArg1 := tmp20227
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6242, tmp20228, B6238, L6239, Key6240, C6241)
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

tmp20229 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20230 := Call(__e, PrimFunc(symshen_4assoc_1_6), symenable_1type_1theory, tmp20225, tmp20229)


tmp20231 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20230)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20230
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20231

tmp20232 := MakeNative(func(__e *ControlFlow) {
V6247 := __e.Get(1)
_ = V6247
__e.Return(MakeNative(func(__e *ControlFlow) {
B6243 := __e.Get(1)
_ = B6243
__e.Return(MakeNative(func(__e *ControlFlow) {
L6244 := __e.Get(1)
_ = L6244
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6245 := __e.Get(1)
_ = Key6245
__e.Return(MakeNative(func(__e *ControlFlow) {
C6246 := __e.Get(1)
_ = C6246
tmp20233 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20234 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20233)
}
__typedArg0 := symlist
__typedArg1 := tmp20233
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20235 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20234, Nil)
}
__typedArg0 := tmp20234
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20236 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20235)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20235
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20237 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp20236)
}
__typedArg0 := symsymbol
__typedArg1 := tmp20236
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6247, tmp20237, B6243, L6244, Key6245, C6246)
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

tmp20238 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20239 := Call(__e, PrimFunc(symshen_4assoc_1_6), symexternal, tmp20232, tmp20238)


tmp20240 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20239)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20239
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20240

tmp20241 := MakeNative(func(__e *ControlFlow) {
V6252 := __e.Get(1)
_ = V6252
__e.Return(MakeNative(func(__e *ControlFlow) {
B6248 := __e.Get(1)
_ = B6248
__e.Return(MakeNative(func(__e *ControlFlow) {
L6249 := __e.Get(1)
_ = L6249
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6250 := __e.Get(1)
_ = Key6250
__e.Return(MakeNative(func(__e *ControlFlow) {
C6251 := __e.Get(1)
_ = C6251
tmp20242 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
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
return PrimCons(symexception, tmp20243)
}
__typedArg0 := symexception
__typedArg1 := tmp20243
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6252, tmp20244, B6248, L6249, Key6250, C6251)
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

tmp20245 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20246 := Call(__e, PrimFunc(symshen_4assoc_1_6), symerror_1to_1string, tmp20241, tmp20245)


tmp20247 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20246)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20246
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20247

tmp20248 := MakeNative(func(__e *ControlFlow) {
V6257 := __e.Get(1)
_ = V6257
__e.Return(MakeNative(func(__e *ControlFlow) {
B6253 := __e.Get(1)
_ = B6253
__e.Return(MakeNative(func(__e *ControlFlow) {
L6254 := __e.Get(1)
_ = L6254
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6255 := __e.Get(1)
_ = Key6255
__e.Return(MakeNative(func(__e *ControlFlow) {
C6256 := __e.Get(1)
_ = C6256
tmp20249 := Call(__e, PrimFunc(symshen_4newpv), B6253)


let__18935 := tmp20249
_ = let__18935

tmp20250 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20251 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20250)
}
__typedArg0 := symlist
__typedArg1 := tmp20250
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20252 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20251, Nil)
}
__typedArg0 := tmp20251
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20253 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20252)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20252
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20254 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18935, tmp20253)
}
__typedArg0 := let__18935
__typedArg1 := tmp20253
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20255 := Call(__e, PrimFunc(symis_b), V6257, tmp20254, B6253, L6254, Key6255, C6256)


__e.TailApply(PrimFunc(symshen_4gc), B6253, tmp20255)
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

tmp20256 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20257 := Call(__e, PrimFunc(symshen_4assoc_1_6), symexplode, tmp20248, tmp20256)


tmp20258 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20257)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20257
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20258

tmp20259 := MakeNative(func(__e *ControlFlow) {
V6262 := __e.Get(1)
_ = V6262
__e.Return(MakeNative(func(__e *ControlFlow) {
B6258 := __e.Get(1)
_ = B6258
__e.Return(MakeNative(func(__e *ControlFlow) {
L6259 := __e.Get(1)
_ = L6259
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6260 := __e.Get(1)
_ = Key6260
__e.Return(MakeNative(func(__e *ControlFlow) {
C6261 := __e.Get(1)
_ = C6261
tmp20260 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20261 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20260)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20260
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20262 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp20261)
}
__typedArg0 := symsymbol
__typedArg1 := tmp20261
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6262, tmp20262, B6258, L6259, Key6260, C6261)
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

tmp20263 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20264 := Call(__e, PrimFunc(symshen_4assoc_1_6), symfactorise, tmp20259, tmp20263)


tmp20265 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20264)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20264
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20265

tmp20266 := MakeNative(func(__e *ControlFlow) {
V6267 := __e.Get(1)
_ = V6267
__e.Return(MakeNative(func(__e *ControlFlow) {
B6263 := __e.Get(1)
_ = B6263
__e.Return(MakeNative(func(__e *ControlFlow) {
L6264 := __e.Get(1)
_ = L6264
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6265 := __e.Get(1)
_ = Key6265
__e.Return(MakeNative(func(__e *ControlFlow) {
C6266 := __e.Get(1)
_ = C6266
tmp20267 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20268 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20267)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20267
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6267, tmp20268, B6263, L6264, Key6265, C6266)
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

tmp20269 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20270 := Call(__e, PrimFunc(symshen_4assoc_1_6), symfactorise_2, tmp20266, tmp20269)


tmp20271 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20270)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20270
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20271

tmp20272 := MakeNative(func(__e *ControlFlow) {
V6272 := __e.Get(1)
_ = V6272
__e.Return(MakeNative(func(__e *ControlFlow) {
B6268 := __e.Get(1)
_ = B6268
__e.Return(MakeNative(func(__e *ControlFlow) {
L6269 := __e.Get(1)
_ = L6269
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6270 := __e.Get(1)
_ = Key6270
__e.Return(MakeNative(func(__e *ControlFlow) {
C6271 := __e.Get(1)
_ = C6271
tmp20273 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20274 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20273)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20273
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6272, tmp20274, B6268, L6269, Key6270, C6271)
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

tmp20275 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20276 := Call(__e, PrimFunc(symshen_4assoc_1_6), symfail, tmp20272, tmp20275)


tmp20277 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20276)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20276
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20277

tmp20278 := MakeNative(func(__e *ControlFlow) {
V6278 := __e.Get(1)
_ = V6278
__e.Return(MakeNative(func(__e *ControlFlow) {
B6274 := __e.Get(1)
_ = B6274
__e.Return(MakeNative(func(__e *ControlFlow) {
L6275 := __e.Get(1)
_ = L6275
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6276 := __e.Get(1)
_ = Key6276
__e.Return(MakeNative(func(__e *ControlFlow) {
C6277 := __e.Get(1)
_ = C6277
tmp20279 := Call(__e, PrimFunc(symshen_4newpv), B6274)


let__18936 := tmp20279
_ = let__18936

tmp20280 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18936, Nil)
}
__typedArg0 := let__18936
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20281 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20280)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20280
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20282 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18936, tmp20281)
}
__typedArg0 := let__18936
__typedArg1 := tmp20281
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20283 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18936, Nil)
}
__typedArg0 := let__18936
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20284 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20283)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20283
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20285 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18936, tmp20284)
}
__typedArg0 := let__18936
__typedArg1 := tmp20284
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20286 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20285, Nil)
}
__typedArg0 := tmp20285
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20287 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20286)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20286
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20288 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20282, tmp20287)
}
__typedArg0 := tmp20282
__typedArg1 := tmp20287
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20289 := Call(__e, PrimFunc(symis_b), V6278, tmp20288, B6274, L6275, Key6276, C6277)


__e.TailApply(PrimFunc(symshen_4gc), B6274, tmp20289)
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

tmp20290 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20291 := Call(__e, PrimFunc(symshen_4assoc_1_6), symfix, tmp20278, tmp20290)


tmp20292 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20291)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20291
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20292

tmp20293 := MakeNative(func(__e *ControlFlow) {
V6283 := __e.Get(1)
_ = V6283
__e.Return(MakeNative(func(__e *ControlFlow) {
B6279 := __e.Get(1)
_ = B6279
__e.Return(MakeNative(func(__e *ControlFlow) {
L6280 := __e.Get(1)
_ = L6280
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6281 := __e.Get(1)
_ = Key6281
__e.Return(MakeNative(func(__e *ControlFlow) {
C6282 := __e.Get(1)
_ = C6282
tmp20294 := Call(__e, PrimFunc(symshen_4newpv), B6279)


let__18937 := tmp20294
_ = let__18937

tmp20295 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18937, Nil)
}
__typedArg0 := let__18937
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20296 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlazy, tmp20295)
}
__typedArg0 := symlazy
__typedArg1 := tmp20295
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20297 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20296, Nil)
}
__typedArg0 := tmp20296
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20298 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20297)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20297
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20299 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18937, tmp20298)
}
__typedArg0 := let__18937
__typedArg1 := tmp20298
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20300 := Call(__e, PrimFunc(symis_b), V6283, tmp20299, B6279, L6280, Key6281, C6282)


__e.TailApply(PrimFunc(symshen_4gc), B6279, tmp20300)
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

tmp20301 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20302 := Call(__e, PrimFunc(symshen_4assoc_1_6), symfreeze, tmp20293, tmp20301)


tmp20303 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20302)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20302
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20303

tmp20304 := MakeNative(func(__e *ControlFlow) {
V6288 := __e.Get(1)
_ = V6288
__e.Return(MakeNative(func(__e *ControlFlow) {
B6284 := __e.Get(1)
_ = B6284
__e.Return(MakeNative(func(__e *ControlFlow) {
L6285 := __e.Get(1)
_ = L6285
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6286 := __e.Get(1)
_ = Key6286
__e.Return(MakeNative(func(__e *ControlFlow) {
C6287 := __e.Get(1)
_ = C6287
tmp20305 := Call(__e, PrimFunc(symshen_4newpv), B6284)


let__18938 := tmp20305
_ = let__18938

tmp20306 := Call(__e, PrimFunc(symshen_4newpv), B6284)


let__18939 := tmp20306
_ = let__18939

tmp20307 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18938, Nil)
}
__typedArg0 := let__18938
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20308 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_d, tmp20307)
}
__typedArg0 := sym_d
__typedArg1 := tmp20307
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20309 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18939, tmp20308)
}
__typedArg0 := let__18939
__typedArg1 := tmp20308
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20310 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18939, Nil)
}
__typedArg0 := let__18939
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20311 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20310)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20310
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20312 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20309, tmp20311)
}
__typedArg0 := tmp20309
__typedArg1 := tmp20311
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20313 := Call(__e, PrimFunc(symis_b), V6288, tmp20312, B6284, L6285, Key6286, C6287)


tmp20314 := Call(__e, PrimFunc(symshen_4gc), B6284, tmp20313)


__e.TailApply(PrimFunc(symshen_4gc), B6284, tmp20314)
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

tmp20315 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20316 := Call(__e, PrimFunc(symshen_4assoc_1_6), symfst, tmp20304, tmp20315)


tmp20317 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20316)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20316
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20317

tmp20318 := MakeNative(func(__e *ControlFlow) {
V6293 := __e.Get(1)
_ = V6293
__e.Return(MakeNative(func(__e *ControlFlow) {
B6289 := __e.Get(1)
_ = B6289
__e.Return(MakeNative(func(__e *ControlFlow) {
L6290 := __e.Get(1)
_ = L6290
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6291 := __e.Get(1)
_ = Key6291
__e.Return(MakeNative(func(__e *ControlFlow) {
C6292 := __e.Get(1)
_ = C6292
tmp20319 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
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
return PrimCons(symsymbol, tmp20320)
}
__typedArg0 := symsymbol
__typedArg1 := tmp20320
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6293, tmp20321, B6289, L6290, Key6291, C6292)
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

tmp20322 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20323 := Call(__e, PrimFunc(symshen_4assoc_1_6), symgensym, tmp20318, tmp20322)


tmp20324 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20323)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20323
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20324

tmp20325 := MakeNative(func(__e *ControlFlow) {
V6299 := __e.Get(1)
_ = V6299
__e.Return(MakeNative(func(__e *ControlFlow) {
B6295 := __e.Get(1)
_ = B6295
__e.Return(MakeNative(func(__e *ControlFlow) {
L6296 := __e.Get(1)
_ = L6296
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6297 := __e.Get(1)
_ = Key6297
__e.Return(MakeNative(func(__e *ControlFlow) {
C6298 := __e.Get(1)
_ = C6298
tmp20326 := Call(__e, PrimFunc(symshen_4newpv), B6295)


let__18940 := tmp20326
_ = let__18940

tmp20327 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18940, Nil)
}
__typedArg0 := let__18940
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20328 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20327)
}
__typedArg0 := symlist
__typedArg1 := tmp20327
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20329 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20330 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20329)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20329
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20331 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18940, tmp20330)
}
__typedArg0 := let__18940
__typedArg1 := tmp20330
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20332 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20331, Nil)
}
__typedArg0 := tmp20331
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20333 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20332)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20332
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20334 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20328, tmp20333)
}
__typedArg0 := tmp20328
__typedArg1 := tmp20333
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20335 := Call(__e, PrimFunc(symis_b), V6299, tmp20334, B6295, L6296, Key6297, C6298)


__e.TailApply(PrimFunc(symshen_4gc), B6295, tmp20335)
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

tmp20336 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20337 := Call(__e, PrimFunc(symshen_4assoc_1_6), symshen_4hds_a_2, tmp20325, tmp20336)


tmp20338 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20337)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20337
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20338

tmp20339 := MakeNative(func(__e *ControlFlow) {
V6304 := __e.Get(1)
_ = V6304
__e.Return(MakeNative(func(__e *ControlFlow) {
B6300 := __e.Get(1)
_ = B6300
__e.Return(MakeNative(func(__e *ControlFlow) {
L6301 := __e.Get(1)
_ = L6301
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6302 := __e.Get(1)
_ = Key6302
__e.Return(MakeNative(func(__e *ControlFlow) {
C6303 := __e.Get(1)
_ = C6303
tmp20340 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20341 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20340)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20340
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20342 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp20341)
}
__typedArg0 := symsymbol
__typedArg1 := tmp20341
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6304, tmp20342, B6300, L6301, Key6302, C6303)
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

tmp20343 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20344 := Call(__e, PrimFunc(symshen_4assoc_1_6), symhush, tmp20339, tmp20343)


tmp20345 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20344)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20344
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20345

tmp20346 := MakeNative(func(__e *ControlFlow) {
V6309 := __e.Get(1)
_ = V6309
__e.Return(MakeNative(func(__e *ControlFlow) {
B6305 := __e.Get(1)
_ = B6305
__e.Return(MakeNative(func(__e *ControlFlow) {
L6306 := __e.Get(1)
_ = L6306
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6307 := __e.Get(1)
_ = Key6307
__e.Return(MakeNative(func(__e *ControlFlow) {
C6308 := __e.Get(1)
_ = C6308
tmp20347 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20348 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20347)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20347
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6309, tmp20348, B6305, L6306, Key6307, C6308)
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

tmp20349 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20350 := Call(__e, PrimFunc(symshen_4assoc_1_6), symhush_2, tmp20346, tmp20349)


tmp20351 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20350)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20350
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20351

tmp20352 := MakeNative(func(__e *ControlFlow) {
V6315 := __e.Get(1)
_ = V6315
__e.Return(MakeNative(func(__e *ControlFlow) {
B6311 := __e.Get(1)
_ = B6311
__e.Return(MakeNative(func(__e *ControlFlow) {
L6312 := __e.Get(1)
_ = L6312
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6313 := __e.Get(1)
_ = Key6313
__e.Return(MakeNative(func(__e *ControlFlow) {
C6314 := __e.Get(1)
_ = C6314
tmp20353 := Call(__e, PrimFunc(symshen_4newpv), B6311)


let__18941 := tmp20353
_ = let__18941

tmp20354 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18941, Nil)
}
__typedArg0 := let__18941
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20355 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp20354)
}
__typedArg0 := symvector
__typedArg1 := tmp20354
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20356 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18941, Nil)
}
__typedArg0 := let__18941
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20357 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20356)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20356
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20358 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp20357)
}
__typedArg0 := symnumber
__typedArg1 := tmp20357
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20359 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20358, Nil)
}
__typedArg0 := tmp20358
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20360 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20359)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20359
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20361 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20355, tmp20360)
}
__typedArg0 := tmp20355
__typedArg1 := tmp20360
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20362 := Call(__e, PrimFunc(symis_b), V6315, tmp20361, B6311, L6312, Key6313, C6314)


__e.TailApply(PrimFunc(symshen_4gc), B6311, tmp20362)
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

tmp20363 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20364 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_5_1vector, tmp20352, tmp20363)


tmp20365 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20364)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20364
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20365

tmp20366 := MakeNative(func(__e *ControlFlow) {
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
tmp20367 := Call(__e, PrimFunc(symshen_4newpv), B6317)


let__18942 := tmp20367
_ = let__18942

tmp20368 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18942, Nil)
}
__typedArg0 := let__18942
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20369 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp20368)
}
__typedArg0 := symvector
__typedArg1 := tmp20368
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20370 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18942, Nil)
}
__typedArg0 := let__18942
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20371 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp20370)
}
__typedArg0 := symvector
__typedArg1 := tmp20370
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20372 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20371, Nil)
}
__typedArg0 := tmp20371
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20373 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20372)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20372
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20374 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18942, tmp20373)
}
__typedArg0 := let__18942
__typedArg1 := tmp20373
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20375 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20374, Nil)
}
__typedArg0 := tmp20374
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20376 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20375)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20375
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20377 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp20376)
}
__typedArg0 := symnumber
__typedArg1 := tmp20376
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20378 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20377, Nil)
}
__typedArg0 := tmp20377
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20379 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20378)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20378
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20380 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20369, tmp20379)
}
__typedArg0 := tmp20369
__typedArg1 := tmp20379
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20381 := Call(__e, PrimFunc(symis_b), V6321, tmp20380, B6317, L6318, Key6319, C6320)


__e.TailApply(PrimFunc(symshen_4gc), B6317, tmp20381)
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

tmp20383 := Call(__e, PrimFunc(symshen_4assoc_1_6), symvector_1_6, tmp20366, tmp20382)


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
V6326 := __e.Get(1)
_ = V6326
__e.Return(MakeNative(func(__e *ControlFlow) {
B6322 := __e.Get(1)
_ = B6322
__e.Return(MakeNative(func(__e *ControlFlow) {
L6323 := __e.Get(1)
_ = L6323
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6324 := __e.Get(1)
_ = Key6324
__e.Return(MakeNative(func(__e *ControlFlow) {
C6325 := __e.Get(1)
_ = C6325
tmp20386 := Call(__e, PrimFunc(symshen_4newpv), B6322)


let__18943 := tmp20386
_ = let__18943

tmp20387 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18943, Nil)
}
__typedArg0 := let__18943
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
return PrimCons(tmp20388, Nil)
}
__typedArg0 := tmp20388
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

tmp20392 := Call(__e, PrimFunc(symis_b), V6326, tmp20391, B6322, L6323, Key6324, C6325)


__e.TailApply(PrimFunc(symshen_4gc), B6322, tmp20392)
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

tmp20393 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20394 := Call(__e, PrimFunc(symshen_4assoc_1_6), symvector, tmp20385, tmp20393)


tmp20395 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20394)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20394
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20395

tmp20396 := MakeNative(func(__e *ControlFlow) {
V6331 := __e.Get(1)
_ = V6331
__e.Return(MakeNative(func(__e *ControlFlow) {
B6327 := __e.Get(1)
_ = B6327
__e.Return(MakeNative(func(__e *ControlFlow) {
L6328 := __e.Get(1)
_ = L6328
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6329 := __e.Get(1)
_ = Key6329
__e.Return(MakeNative(func(__e *ControlFlow) {
C6330 := __e.Get(1)
_ = C6330
tmp20397 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20398 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20397)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20397
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20399 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp20398)
}
__typedArg0 := symsymbol
__typedArg1 := tmp20398
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6331, tmp20399, B6327, L6328, Key6329, C6330)
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

tmp20400 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20401 := Call(__e, PrimFunc(symshen_4assoc_1_6), symget_1time, tmp20396, tmp20400)


tmp20402 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20401)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20401
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20402

tmp20403 := MakeNative(func(__e *ControlFlow) {
V6336 := __e.Get(1)
_ = V6336
__e.Return(MakeNative(func(__e *ControlFlow) {
B6332 := __e.Get(1)
_ = B6332
__e.Return(MakeNative(func(__e *ControlFlow) {
L6333 := __e.Get(1)
_ = L6333
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6334 := __e.Get(1)
_ = Key6334
__e.Return(MakeNative(func(__e *ControlFlow) {
C6335 := __e.Get(1)
_ = C6335
tmp20404 := Call(__e, PrimFunc(symshen_4newpv), B6332)


let__18944 := tmp20404
_ = let__18944

tmp20405 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
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
return PrimCons(symnumber, tmp20406)
}
__typedArg0 := symnumber
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
return PrimCons(let__18944, tmp20409)
}
__typedArg0 := let__18944
__typedArg1 := tmp20409
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20411 := Call(__e, PrimFunc(symis_b), V6336, tmp20410, B6332, L6333, Key6334, C6335)


__e.TailApply(PrimFunc(symshen_4gc), B6332, tmp20411)
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

tmp20412 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20413 := Call(__e, PrimFunc(symshen_4assoc_1_6), symhash, tmp20403, tmp20412)


tmp20414 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20413)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20413
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20414

tmp20415 := MakeNative(func(__e *ControlFlow) {
V6341 := __e.Get(1)
_ = V6341
__e.Return(MakeNative(func(__e *ControlFlow) {
B6337 := __e.Get(1)
_ = B6337
__e.Return(MakeNative(func(__e *ControlFlow) {
L6338 := __e.Get(1)
_ = L6338
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6339 := __e.Get(1)
_ = Key6339
__e.Return(MakeNative(func(__e *ControlFlow) {
C6340 := __e.Get(1)
_ = C6340
tmp20416 := Call(__e, PrimFunc(symshen_4newpv), B6337)


let__18945 := tmp20416
_ = let__18945

tmp20417 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18945, Nil)
}
__typedArg0 := let__18945
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20418 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20417)
}
__typedArg0 := symlist
__typedArg1 := tmp20417
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20419 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18945, Nil)
}
__typedArg0 := let__18945
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20420 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20419)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20419
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20421 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20418, tmp20420)
}
__typedArg0 := tmp20418
__typedArg1 := tmp20420
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20422 := Call(__e, PrimFunc(symis_b), V6341, tmp20421, B6337, L6338, Key6339, C6340)


__e.TailApply(PrimFunc(symshen_4gc), B6337, tmp20422)
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

tmp20423 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20424 := Call(__e, PrimFunc(symshen_4assoc_1_6), symhead, tmp20415, tmp20423)


tmp20425 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20424)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20424
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20425

tmp20426 := MakeNative(func(__e *ControlFlow) {
V6346 := __e.Get(1)
_ = V6346
__e.Return(MakeNative(func(__e *ControlFlow) {
B6342 := __e.Get(1)
_ = B6342
__e.Return(MakeNative(func(__e *ControlFlow) {
L6343 := __e.Get(1)
_ = L6343
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6344 := __e.Get(1)
_ = Key6344
__e.Return(MakeNative(func(__e *ControlFlow) {
C6345 := __e.Get(1)
_ = C6345
tmp20427 := Call(__e, PrimFunc(symshen_4newpv), B6342)


let__18946 := tmp20427
_ = let__18946

tmp20428 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18946, Nil)
}
__typedArg0 := let__18946
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20429 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp20428)
}
__typedArg0 := symvector
__typedArg1 := tmp20428
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20430 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18946, Nil)
}
__typedArg0 := let__18946
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
return PrimCons(tmp20429, tmp20431)
}
__typedArg0 := tmp20429
__typedArg1 := tmp20431
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20433 := Call(__e, PrimFunc(symis_b), V6346, tmp20432, B6342, L6343, Key6344, C6345)


__e.TailApply(PrimFunc(symshen_4gc), B6342, tmp20433)
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

tmp20434 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20435 := Call(__e, PrimFunc(symshen_4assoc_1_6), symhdv, tmp20426, tmp20434)


tmp20436 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20435)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20435
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20436

tmp20437 := MakeNative(func(__e *ControlFlow) {
V6351 := __e.Get(1)
_ = V6351
__e.Return(MakeNative(func(__e *ControlFlow) {
B6347 := __e.Get(1)
_ = B6347
__e.Return(MakeNative(func(__e *ControlFlow) {
L6348 := __e.Get(1)
_ = L6348
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6349 := __e.Get(1)
_ = Key6349
__e.Return(MakeNative(func(__e *ControlFlow) {
C6350 := __e.Get(1)
_ = C6350
tmp20438 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
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
return PrimCons(symstring, tmp20439)
}
__typedArg0 := symstring
__typedArg1 := tmp20439
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6351, tmp20440, B6347, L6348, Key6349, C6350)
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

tmp20441 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20442 := Call(__e, PrimFunc(symshen_4assoc_1_6), symhdstr, tmp20437, tmp20441)


tmp20443 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20442)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20442
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20443

tmp20444 := MakeNative(func(__e *ControlFlow) {
V6356 := __e.Get(1)
_ = V6356
__e.Return(MakeNative(func(__e *ControlFlow) {
B6352 := __e.Get(1)
_ = B6352
__e.Return(MakeNative(func(__e *ControlFlow) {
L6353 := __e.Get(1)
_ = L6353
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6354 := __e.Get(1)
_ = Key6354
__e.Return(MakeNative(func(__e *ControlFlow) {
C6355 := __e.Get(1)
_ = C6355
tmp20445 := Call(__e, PrimFunc(symshen_4newpv), B6352)


let__18947 := tmp20445
_ = let__18947

tmp20446 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18947, Nil)
}
__typedArg0 := let__18947
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20447 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20446)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20446
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20448 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18947, tmp20447)
}
__typedArg0 := let__18947
__typedArg1 := tmp20447
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20449 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20448, Nil)
}
__typedArg0 := tmp20448
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20450 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20449)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20449
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20451 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18947, tmp20450)
}
__typedArg0 := let__18947
__typedArg1 := tmp20450
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20452 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20451, Nil)
}
__typedArg0 := tmp20451
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
return PrimCons(symboolean, tmp20453)
}
__typedArg0 := symboolean
__typedArg1 := tmp20453
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20455 := Call(__e, PrimFunc(symis_b), V6356, tmp20454, B6352, L6353, Key6354, C6355)


__e.TailApply(PrimFunc(symshen_4gc), B6352, tmp20455)
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

tmp20457 := Call(__e, PrimFunc(symshen_4assoc_1_6), symif, tmp20444, tmp20456)


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
V6361 := __e.Get(1)
_ = V6361
__e.Return(MakeNative(func(__e *ControlFlow) {
B6357 := __e.Get(1)
_ = B6357
__e.Return(MakeNative(func(__e *ControlFlow) {
L6358 := __e.Get(1)
_ = L6358
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6359 := __e.Get(1)
_ = Key6359
__e.Return(MakeNative(func(__e *ControlFlow) {
C6360 := __e.Get(1)
_ = C6360
tmp20460 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20461 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20460)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20460
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20462 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp20461)
}
__typedArg0 := symsymbol
__typedArg1 := tmp20461
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6361, tmp20462, B6357, L6358, Key6359, C6360)
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

tmp20463 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20464 := Call(__e, PrimFunc(symshen_4assoc_1_6), symin_1package, tmp20459, tmp20463)


tmp20465 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20464)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20464
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20465

tmp20466 := MakeNative(func(__e *ControlFlow) {
V6366 := __e.Get(1)
_ = V6366
__e.Return(MakeNative(func(__e *ControlFlow) {
B6362 := __e.Get(1)
_ = B6362
__e.Return(MakeNative(func(__e *ControlFlow) {
L6363 := __e.Get(1)
_ = L6363
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6364 := __e.Get(1)
_ = Key6364
__e.Return(MakeNative(func(__e *ControlFlow) {
C6365 := __e.Get(1)
_ = C6365
tmp20467 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20468 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20467)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20467
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6366, tmp20468, B6362, L6363, Key6364, C6365)
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

tmp20469 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20470 := Call(__e, PrimFunc(symshen_4assoc_1_6), symit, tmp20466, tmp20469)


tmp20471 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20470)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20470
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20471

tmp20472 := MakeNative(func(__e *ControlFlow) {
V6371 := __e.Get(1)
_ = V6371
__e.Return(MakeNative(func(__e *ControlFlow) {
B6367 := __e.Get(1)
_ = B6367
__e.Return(MakeNative(func(__e *ControlFlow) {
L6368 := __e.Get(1)
_ = L6368
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6369 := __e.Get(1)
_ = Key6369
__e.Return(MakeNative(func(__e *ControlFlow) {
C6370 := __e.Get(1)
_ = C6370
tmp20473 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20474 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20473)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20473
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6371, tmp20474, B6367, L6368, Key6369, C6370)
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

tmp20475 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20476 := Call(__e, PrimFunc(symshen_4assoc_1_6), symimplementation, tmp20472, tmp20475)


tmp20477 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20476)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20476
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20477

tmp20478 := MakeNative(func(__e *ControlFlow) {
V6377 := __e.Get(1)
_ = V6377
__e.Return(MakeNative(func(__e *ControlFlow) {
B6373 := __e.Get(1)
_ = B6373
__e.Return(MakeNative(func(__e *ControlFlow) {
L6374 := __e.Get(1)
_ = L6374
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6375 := __e.Get(1)
_ = Key6375
__e.Return(MakeNative(func(__e *ControlFlow) {
C6376 := __e.Get(1)
_ = C6376
tmp20479 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20480 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20479)
}
__typedArg0 := symlist
__typedArg1 := tmp20479
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20481 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20482 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20481)
}
__typedArg0 := symlist
__typedArg1 := tmp20481
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20483 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20482, Nil)
}
__typedArg0 := tmp20482
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20484 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20483)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20483
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20485 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20480, tmp20484)
}
__typedArg0 := tmp20480
__typedArg1 := tmp20484
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6377, tmp20485, B6373, L6374, Key6375, C6376)
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

tmp20486 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20487 := Call(__e, PrimFunc(symshen_4assoc_1_6), syminclude, tmp20478, tmp20486)


tmp20488 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20487)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20487
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20488

tmp20489 := MakeNative(func(__e *ControlFlow) {
V6383 := __e.Get(1)
_ = V6383
__e.Return(MakeNative(func(__e *ControlFlow) {
B6379 := __e.Get(1)
_ = B6379
__e.Return(MakeNative(func(__e *ControlFlow) {
L6380 := __e.Get(1)
_ = L6380
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6381 := __e.Get(1)
_ = Key6381
__e.Return(MakeNative(func(__e *ControlFlow) {
C6382 := __e.Get(1)
_ = C6382
tmp20490 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20491 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20490)
}
__typedArg0 := symlist
__typedArg1 := tmp20490
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20492 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20493 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20492)
}
__typedArg0 := symlist
__typedArg1 := tmp20492
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20494 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20493, Nil)
}
__typedArg0 := tmp20493
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20495 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20494)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20494
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20496 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20491, tmp20495)
}
__typedArg0 := tmp20491
__typedArg1 := tmp20495
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6383, tmp20496, B6379, L6380, Key6381, C6382)
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

tmp20497 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20498 := Call(__e, PrimFunc(symshen_4assoc_1_6), syminclude_1all_1but, tmp20489, tmp20497)


tmp20499 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20498)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20498
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20499

tmp20500 := MakeNative(func(__e *ControlFlow) {
V6388 := __e.Get(1)
_ = V6388
__e.Return(MakeNative(func(__e *ControlFlow) {
B6384 := __e.Get(1)
_ = B6384
__e.Return(MakeNative(func(__e *ControlFlow) {
L6385 := __e.Get(1)
_ = L6385
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6386 := __e.Get(1)
_ = Key6386
__e.Return(MakeNative(func(__e *ControlFlow) {
C6387 := __e.Get(1)
_ = C6387
tmp20501 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20502 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20501)
}
__typedArg0 := symlist
__typedArg1 := tmp20501
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20503 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20502, Nil)
}
__typedArg0 := tmp20502
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20504 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20503)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20503
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6388, tmp20504, B6384, L6385, Key6386, C6387)
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

tmp20505 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20506 := Call(__e, PrimFunc(symshen_4assoc_1_6), symincluded, tmp20500, tmp20505)


tmp20507 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20506)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20506
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20507

tmp20508 := MakeNative(func(__e *ControlFlow) {
V6393 := __e.Get(1)
_ = V6393
__e.Return(MakeNative(func(__e *ControlFlow) {
B6389 := __e.Get(1)
_ = B6389
__e.Return(MakeNative(func(__e *ControlFlow) {
L6390 := __e.Get(1)
_ = L6390
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6391 := __e.Get(1)
_ = Key6391
__e.Return(MakeNative(func(__e *ControlFlow) {
C6392 := __e.Get(1)
_ = C6392
tmp20509 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20510 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20509)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20509
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6393, tmp20510, B6389, L6390, Key6391, C6392)
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

tmp20511 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20512 := Call(__e, PrimFunc(symshen_4assoc_1_6), syminferences, tmp20508, tmp20511)


tmp20513 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20512)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20512
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20513

tmp20514 := MakeNative(func(__e *ControlFlow) {
V6398 := __e.Get(1)
_ = V6398
__e.Return(MakeNative(func(__e *ControlFlow) {
B6394 := __e.Get(1)
_ = B6394
__e.Return(MakeNative(func(__e *ControlFlow) {
L6395 := __e.Get(1)
_ = L6395
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6396 := __e.Get(1)
_ = Key6396
__e.Return(MakeNative(func(__e *ControlFlow) {
C6397 := __e.Get(1)
_ = C6397
tmp20515 := Call(__e, PrimFunc(symshen_4newpv), B6394)


let__18948 := tmp20515
_ = let__18948

tmp20516 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
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
return PrimCons(symstring, tmp20517)
}
__typedArg0 := symstring
__typedArg1 := tmp20517
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20519 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20518, Nil)
}
__typedArg0 := tmp20518
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20520 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20519)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20519
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20521 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18948, tmp20520)
}
__typedArg0 := let__18948
__typedArg1 := tmp20520
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20522 := Call(__e, PrimFunc(symis_b), V6398, tmp20521, B6394, L6395, Key6396, C6397)


__e.TailApply(PrimFunc(symshen_4gc), B6394, tmp20522)
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

tmp20523 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20524 := Call(__e, PrimFunc(symshen_4assoc_1_6), symshen_4insert, tmp20514, tmp20523)


tmp20525 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20524)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20524
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20525

tmp20526 := MakeNative(func(__e *ControlFlow) {
V6403 := __e.Get(1)
_ = V6403
__e.Return(MakeNative(func(__e *ControlFlow) {
B6399 := __e.Get(1)
_ = B6399
__e.Return(MakeNative(func(__e *ControlFlow) {
L6400 := __e.Get(1)
_ = L6400
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6401 := __e.Get(1)
_ = Key6401
__e.Return(MakeNative(func(__e *ControlFlow) {
C6402 := __e.Get(1)
_ = C6402
tmp20527 := Call(__e, PrimFunc(symshen_4newpv), B6399)


let__18949 := tmp20527
_ = let__18949

tmp20528 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20529 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20528)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20528
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20530 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18949, tmp20529)
}
__typedArg0 := let__18949
__typedArg1 := tmp20529
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20531 := Call(__e, PrimFunc(symis_b), V6403, tmp20530, B6399, L6400, Key6401, C6402)


__e.TailApply(PrimFunc(symshen_4gc), B6399, tmp20531)
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

tmp20532 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20533 := Call(__e, PrimFunc(symshen_4assoc_1_6), syminteger_2, tmp20526, tmp20532)


tmp20534 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20533)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20533
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20534

tmp20535 := MakeNative(func(__e *ControlFlow) {
V6408 := __e.Get(1)
_ = V6408
__e.Return(MakeNative(func(__e *ControlFlow) {
B6404 := __e.Get(1)
_ = B6404
__e.Return(MakeNative(func(__e *ControlFlow) {
L6405 := __e.Get(1)
_ = L6405
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6406 := __e.Get(1)
_ = Key6406
__e.Return(MakeNative(func(__e *ControlFlow) {
C6407 := __e.Get(1)
_ = C6407
tmp20536 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20537 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20536)
}
__typedArg0 := symlist
__typedArg1 := tmp20536
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20538 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20537, Nil)
}
__typedArg0 := tmp20537
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20539 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20538)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20538
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20540 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp20539)
}
__typedArg0 := symsymbol
__typedArg1 := tmp20539
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6408, tmp20540, B6404, L6405, Key6406, C6407)
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

tmp20541 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20542 := Call(__e, PrimFunc(symshen_4assoc_1_6), syminternal, tmp20535, tmp20541)


tmp20543 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20542)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20542
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20543

tmp20544 := MakeNative(func(__e *ControlFlow) {
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
tmp20545 := Call(__e, PrimFunc(symshen_4newpv), B6411)


let__18950 := tmp20545
_ = let__18950

tmp20546 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18950, Nil)
}
__typedArg0 := let__18950
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20547 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20546)
}
__typedArg0 := symlist
__typedArg1 := tmp20546
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20548 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18950, Nil)
}
__typedArg0 := let__18950
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20549 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20548)
}
__typedArg0 := symlist
__typedArg1 := tmp20548
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20550 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18950, Nil)
}
__typedArg0 := let__18950
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20551 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20550)
}
__typedArg0 := symlist
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
return PrimCons(tmp20549, tmp20553)
}
__typedArg0 := tmp20549
__typedArg1 := tmp20553
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20555 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20554, Nil)
}
__typedArg0 := tmp20554
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20556 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20555)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20555
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20557 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20547, tmp20556)
}
__typedArg0 := tmp20547
__typedArg1 := tmp20556
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20558 := Call(__e, PrimFunc(symis_b), V6415, tmp20557, B6411, L6412, Key6413, C6414)


__e.TailApply(PrimFunc(symshen_4gc), B6411, tmp20558)
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

tmp20559 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20560 := Call(__e, PrimFunc(symshen_4assoc_1_6), symintersection, tmp20544, tmp20559)


tmp20561 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20560)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20560
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20561

tmp20562 := MakeNative(func(__e *ControlFlow) {
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
tmp20563 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20564 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20563)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20563
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6420, tmp20564, B6416, L6417, Key6418, C6419)
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

tmp20566 := Call(__e, PrimFunc(symshen_4assoc_1_6), symlanguage, tmp20562, tmp20565)


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
V6425 := __e.Get(1)
_ = V6425
__e.Return(MakeNative(func(__e *ControlFlow) {
B6421 := __e.Get(1)
_ = B6421
__e.Return(MakeNative(func(__e *ControlFlow) {
L6422 := __e.Get(1)
_ = L6422
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6423 := __e.Get(1)
_ = Key6423
__e.Return(MakeNative(func(__e *ControlFlow) {
C6424 := __e.Get(1)
_ = C6424
tmp20569 := Call(__e, PrimFunc(symshen_4newpv), B6421)


let__18951 := tmp20569
_ = let__18951

tmp20570 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18951, Nil)
}
__typedArg0 := let__18951
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20571 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20570)
}
__typedArg0 := symlist
__typedArg1 := tmp20570
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20572 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20573 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20572)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20572
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20574 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20571, tmp20573)
}
__typedArg0 := tmp20571
__typedArg1 := tmp20573
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20575 := Call(__e, PrimFunc(symis_b), V6425, tmp20574, B6421, L6422, Key6423, C6424)


__e.TailApply(PrimFunc(symshen_4gc), B6421, tmp20575)
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

tmp20576 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20577 := Call(__e, PrimFunc(symshen_4assoc_1_6), symlength, tmp20568, tmp20576)


tmp20578 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20577)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20577
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20578

tmp20579 := MakeNative(func(__e *ControlFlow) {
V6430 := __e.Get(1)
_ = V6430
__e.Return(MakeNative(func(__e *ControlFlow) {
B6426 := __e.Get(1)
_ = B6426
__e.Return(MakeNative(func(__e *ControlFlow) {
L6427 := __e.Get(1)
_ = L6427
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6428 := __e.Get(1)
_ = Key6428
__e.Return(MakeNative(func(__e *ControlFlow) {
C6429 := __e.Get(1)
_ = C6429
tmp20580 := Call(__e, PrimFunc(symshen_4newpv), B6426)


let__18952 := tmp20580
_ = let__18952

tmp20581 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18952, Nil)
}
__typedArg0 := let__18952
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20582 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp20581)
}
__typedArg0 := symvector
__typedArg1 := tmp20581
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20583 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20584 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20583)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20583
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20585 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20582, tmp20584)
}
__typedArg0 := tmp20582
__typedArg1 := tmp20584
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20586 := Call(__e, PrimFunc(symis_b), V6430, tmp20585, B6426, L6427, Key6428, C6429)


__e.TailApply(PrimFunc(symshen_4gc), B6426, tmp20586)
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

tmp20587 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20588 := Call(__e, PrimFunc(symshen_4assoc_1_6), symlimit, tmp20579, tmp20587)


tmp20589 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20588)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20588
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20589

tmp20590 := MakeNative(func(__e *ControlFlow) {
V6436 := __e.Get(1)
_ = V6436
__e.Return(MakeNative(func(__e *ControlFlow) {
B6432 := __e.Get(1)
_ = B6432
__e.Return(MakeNative(func(__e *ControlFlow) {
L6433 := __e.Get(1)
_ = L6433
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6434 := __e.Get(1)
_ = Key6434
__e.Return(MakeNative(func(__e *ControlFlow) {
C6435 := __e.Get(1)
_ = C6435
tmp20591 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symin, Nil)
}
__typedArg0 := symin
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20592 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp20591)
}
__typedArg0 := symstream
__typedArg1 := tmp20591
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20593 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunit, Nil)
}
__typedArg0 := symunit
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20594 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20593)
}
__typedArg0 := symlist
__typedArg1 := tmp20593
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20595 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20594, Nil)
}
__typedArg0 := tmp20594
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20596 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20595)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20595
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20597 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20592, tmp20596)
}
__typedArg0 := tmp20592
__typedArg1 := tmp20596
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6436, tmp20597, B6432, L6433, Key6434, C6435)
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

tmp20599 := Call(__e, PrimFunc(symshen_4assoc_1_6), symlineread, tmp20590, tmp20598)


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
V6441 := __e.Get(1)
_ = V6441
__e.Return(MakeNative(func(__e *ControlFlow) {
B6437 := __e.Get(1)
_ = B6437
__e.Return(MakeNative(func(__e *ControlFlow) {
L6438 := __e.Get(1)
_ = L6438
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6439 := __e.Get(1)
_ = Key6439
__e.Return(MakeNative(func(__e *ControlFlow) {
C6440 := __e.Get(1)
_ = C6440
tmp20602 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20603 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20602)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20602
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20604 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp20603)
}
__typedArg0 := symstring
__typedArg1 := tmp20603
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6441, tmp20604, B6437, L6438, Key6439, C6440)
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

tmp20605 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20606 := Call(__e, PrimFunc(symshen_4assoc_1_6), symload, tmp20601, tmp20605)


tmp20607 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20606)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20606
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20607

tmp20608 := MakeNative(func(__e *ControlFlow) {
V6448 := __e.Get(1)
_ = V6448
__e.Return(MakeNative(func(__e *ControlFlow) {
B6444 := __e.Get(1)
_ = B6444
__e.Return(MakeNative(func(__e *ControlFlow) {
L6445 := __e.Get(1)
_ = L6445
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6446 := __e.Get(1)
_ = Key6446
__e.Return(MakeNative(func(__e *ControlFlow) {
C6447 := __e.Get(1)
_ = C6447
tmp20609 := Call(__e, PrimFunc(symshen_4newpv), B6444)


let__18953 := tmp20609
_ = let__18953

tmp20610 := Call(__e, PrimFunc(symshen_4newpv), B6444)


let__18954 := tmp20610
_ = let__18954

tmp20611 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18954, Nil)
}
__typedArg0 := let__18954
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20612 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20611)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20611
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20613 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18953, tmp20612)
}
__typedArg0 := let__18953
__typedArg1 := tmp20612
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20614 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18953, Nil)
}
__typedArg0 := let__18953
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20615 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20614)
}
__typedArg0 := symlist
__typedArg1 := tmp20614
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20616 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18954, Nil)
}
__typedArg0 := let__18954
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20617 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20616)
}
__typedArg0 := symlist
__typedArg1 := tmp20616
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20618 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20617, Nil)
}
__typedArg0 := tmp20617
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20619 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20618)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20618
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20620 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20615, tmp20619)
}
__typedArg0 := tmp20615
__typedArg1 := tmp20619
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20621 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20620, Nil)
}
__typedArg0 := tmp20620
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20622 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20621)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20621
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20623 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20613, tmp20622)
}
__typedArg0 := tmp20613
__typedArg1 := tmp20622
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20624 := Call(__e, PrimFunc(symis_b), V6448, tmp20623, B6444, L6445, Key6446, C6447)


tmp20625 := Call(__e, PrimFunc(symshen_4gc), B6444, tmp20624)


__e.TailApply(PrimFunc(symshen_4gc), B6444, tmp20625)
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

tmp20626 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20627 := Call(__e, PrimFunc(symshen_4assoc_1_6), symmap, tmp20608, tmp20626)


tmp20628 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20627)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20627
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20628

tmp20629 := MakeNative(func(__e *ControlFlow) {
V6455 := __e.Get(1)
_ = V6455
__e.Return(MakeNative(func(__e *ControlFlow) {
B6451 := __e.Get(1)
_ = B6451
__e.Return(MakeNative(func(__e *ControlFlow) {
L6452 := __e.Get(1)
_ = L6452
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6453 := __e.Get(1)
_ = Key6453
__e.Return(MakeNative(func(__e *ControlFlow) {
C6454 := __e.Get(1)
_ = C6454
tmp20630 := Call(__e, PrimFunc(symshen_4newpv), B6451)


let__18955 := tmp20630
_ = let__18955

tmp20631 := Call(__e, PrimFunc(symshen_4newpv), B6451)


let__18956 := tmp20631
_ = let__18956

tmp20632 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18956, Nil)
}
__typedArg0 := let__18956
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20633 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20632)
}
__typedArg0 := symlist
__typedArg1 := tmp20632
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20634 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20633, Nil)
}
__typedArg0 := tmp20633
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20635 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20634)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20634
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20636 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18955, tmp20635)
}
__typedArg0 := let__18955
__typedArg1 := tmp20635
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20637 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18955, Nil)
}
__typedArg0 := let__18955
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20638 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20637)
}
__typedArg0 := symlist
__typedArg1 := tmp20637
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20639 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18956, Nil)
}
__typedArg0 := let__18956
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20640 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20639)
}
__typedArg0 := symlist
__typedArg1 := tmp20639
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20641 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20640, Nil)
}
__typedArg0 := tmp20640
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20642 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20641)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20641
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20643 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20638, tmp20642)
}
__typedArg0 := tmp20638
__typedArg1 := tmp20642
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20644 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20643, Nil)
}
__typedArg0 := tmp20643
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
return PrimCons(tmp20636, tmp20645)
}
__typedArg0 := tmp20636
__typedArg1 := tmp20645
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20647 := Call(__e, PrimFunc(symis_b), V6455, tmp20646, B6451, L6452, Key6453, C6454)


tmp20648 := Call(__e, PrimFunc(symshen_4gc), B6451, tmp20647)


__e.TailApply(PrimFunc(symshen_4gc), B6451, tmp20648)
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

tmp20649 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20650 := Call(__e, PrimFunc(symshen_4assoc_1_6), symmapcan, tmp20629, tmp20649)


tmp20651 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20650)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20650
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20651

tmp20652 := MakeNative(func(__e *ControlFlow) {
V6460 := __e.Get(1)
_ = V6460
__e.Return(MakeNative(func(__e *ControlFlow) {
B6456 := __e.Get(1)
_ = B6456
__e.Return(MakeNative(func(__e *ControlFlow) {
L6457 := __e.Get(1)
_ = L6457
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6458 := __e.Get(1)
_ = Key6458
__e.Return(MakeNative(func(__e *ControlFlow) {
C6459 := __e.Get(1)
_ = C6459
tmp20653 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20654 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20653)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20653
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20655 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp20654)
}
__typedArg0 := symnumber
__typedArg1 := tmp20654
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6460, tmp20655, B6456, L6457, Key6458, C6459)
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

tmp20656 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20657 := Call(__e, PrimFunc(symshen_4assoc_1_6), symmaxinferences, tmp20652, tmp20656)


tmp20658 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20657)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20657
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20658

tmp20659 := MakeNative(func(__e *ControlFlow) {
V6465 := __e.Get(1)
_ = V6465
__e.Return(MakeNative(func(__e *ControlFlow) {
B6461 := __e.Get(1)
_ = B6461
__e.Return(MakeNative(func(__e *ControlFlow) {
L6462 := __e.Get(1)
_ = L6462
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6463 := __e.Get(1)
_ = Key6463
__e.Return(MakeNative(func(__e *ControlFlow) {
C6464 := __e.Get(1)
_ = C6464
tmp20660 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20661 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20660)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20660
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20662 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp20661)
}
__typedArg0 := symnumber
__typedArg1 := tmp20661
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6465, tmp20662, B6461, L6462, Key6463, C6464)
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

tmp20663 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20664 := Call(__e, PrimFunc(symshen_4assoc_1_6), symn_1_6string, tmp20659, tmp20663)


tmp20665 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20664)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20664
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20665

tmp20666 := MakeNative(func(__e *ControlFlow) {
V6470 := __e.Get(1)
_ = V6470
__e.Return(MakeNative(func(__e *ControlFlow) {
B6466 := __e.Get(1)
_ = B6466
__e.Return(MakeNative(func(__e *ControlFlow) {
L6467 := __e.Get(1)
_ = L6467
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6468 := __e.Get(1)
_ = Key6468
__e.Return(MakeNative(func(__e *ControlFlow) {
C6469 := __e.Get(1)
_ = C6469
tmp20667 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
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
return PrimCons(symnumber, tmp20668)
}
__typedArg0 := symnumber
__typedArg1 := tmp20668
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6470, tmp20669, B6466, L6467, Key6468, C6469)
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

tmp20670 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20671 := Call(__e, PrimFunc(symshen_4assoc_1_6), symnl, tmp20666, tmp20670)


tmp20672 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20671)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20671
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20672

tmp20673 := MakeNative(func(__e *ControlFlow) {
V6475 := __e.Get(1)
_ = V6475
__e.Return(MakeNative(func(__e *ControlFlow) {
B6471 := __e.Get(1)
_ = B6471
__e.Return(MakeNative(func(__e *ControlFlow) {
L6472 := __e.Get(1)
_ = L6472
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6473 := __e.Get(1)
_ = Key6473
__e.Return(MakeNative(func(__e *ControlFlow) {
C6474 := __e.Get(1)
_ = C6474
tmp20674 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
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
return PrimCons(symboolean, tmp20675)
}
__typedArg0 := symboolean
__typedArg1 := tmp20675
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6475, tmp20676, B6471, L6472, Key6473, C6474)
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

tmp20677 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20678 := Call(__e, PrimFunc(symshen_4assoc_1_6), symnot, tmp20673, tmp20677)


tmp20679 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20678)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20678
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20679

tmp20680 := MakeNative(func(__e *ControlFlow) {
V6480 := __e.Get(1)
_ = V6480
__e.Return(MakeNative(func(__e *ControlFlow) {
B6476 := __e.Get(1)
_ = B6476
__e.Return(MakeNative(func(__e *ControlFlow) {
L6477 := __e.Get(1)
_ = L6477
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6478 := __e.Get(1)
_ = Key6478
__e.Return(MakeNative(func(__e *ControlFlow) {
C6479 := __e.Get(1)
_ = C6479
tmp20681 := Call(__e, PrimFunc(symshen_4newpv), B6476)


let__18957 := tmp20681
_ = let__18957

tmp20682 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18957, Nil)
}
__typedArg0 := let__18957
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20683 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20682)
}
__typedArg0 := symlist
__typedArg1 := tmp20682
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20684 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18957, Nil)
}
__typedArg0 := let__18957
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20685 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20684)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20684
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20686 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20683, tmp20685)
}
__typedArg0 := tmp20683
__typedArg1 := tmp20685
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20687 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20686, Nil)
}
__typedArg0 := tmp20686
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20688 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20687)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20687
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20689 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp20688)
}
__typedArg0 := symnumber
__typedArg1 := tmp20688
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20690 := Call(__e, PrimFunc(symis_b), V6480, tmp20689, B6476, L6477, Key6478, C6479)


__e.TailApply(PrimFunc(symshen_4gc), B6476, tmp20690)
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

tmp20691 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20692 := Call(__e, PrimFunc(symshen_4assoc_1_6), symnth, tmp20680, tmp20691)


tmp20693 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20692)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20692
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20693

tmp20694 := MakeNative(func(__e *ControlFlow) {
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
tmp20695 := Call(__e, PrimFunc(symshen_4newpv), B6481)


let__18958 := tmp20695
_ = let__18958

tmp20696 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20697 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20696)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20696
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20698 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18958, tmp20697)
}
__typedArg0 := let__18958
__typedArg1 := tmp20697
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20699 := Call(__e, PrimFunc(symis_b), V6485, tmp20698, B6481, L6482, Key6483, C6484)


__e.TailApply(PrimFunc(symshen_4gc), B6481, tmp20699)
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

tmp20700 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20701 := Call(__e, PrimFunc(symshen_4assoc_1_6), symnumber_2, tmp20694, tmp20700)


tmp20702 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20701)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20701
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20702

tmp20703 := MakeNative(func(__e *ControlFlow) {
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
tmp20704 := Call(__e, PrimFunc(symshen_4newpv), B6486)


let__18959 := tmp20704
_ = let__18959

tmp20705 := Call(__e, PrimFunc(symshen_4newpv), B6486)


let__18960 := tmp20705
_ = let__18960

tmp20706 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20707 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20706)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20706
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20708 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18960, tmp20707)
}
__typedArg0 := let__18960
__typedArg1 := tmp20707
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20709 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20708, Nil)
}
__typedArg0 := tmp20708
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20710 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20709)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20709
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20711 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18959, tmp20710)
}
__typedArg0 := let__18959
__typedArg1 := tmp20710
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20712 := Call(__e, PrimFunc(symis_b), V6490, tmp20711, B6486, L6487, Key6488, C6489)


tmp20713 := Call(__e, PrimFunc(symshen_4gc), B6486, tmp20712)


__e.TailApply(PrimFunc(symshen_4gc), B6486, tmp20713)
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

tmp20714 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20715 := Call(__e, PrimFunc(symshen_4assoc_1_6), symoccurrences, tmp20703, tmp20714)


tmp20716 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20715)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20715
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20716

tmp20717 := MakeNative(func(__e *ControlFlow) {
V6495 := __e.Get(1)
_ = V6495
__e.Return(MakeNative(func(__e *ControlFlow) {
B6491 := __e.Get(1)
_ = B6491
__e.Return(MakeNative(func(__e *ControlFlow) {
L6492 := __e.Get(1)
_ = L6492
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6493 := __e.Get(1)
_ = Key6493
__e.Return(MakeNative(func(__e *ControlFlow) {
C6494 := __e.Get(1)
_ = C6494
tmp20718 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20719 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20718)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20718
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20720 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp20719)
}
__typedArg0 := symsymbol
__typedArg1 := tmp20719
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6495, tmp20720, B6491, L6492, Key6493, C6494)
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

tmp20721 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20722 := Call(__e, PrimFunc(symshen_4assoc_1_6), symoccurs_1check, tmp20717, tmp20721)


tmp20723 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20722)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20722
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20723

tmp20724 := MakeNative(func(__e *ControlFlow) {
V6500 := __e.Get(1)
_ = V6500
__e.Return(MakeNative(func(__e *ControlFlow) {
B6496 := __e.Get(1)
_ = B6496
__e.Return(MakeNative(func(__e *ControlFlow) {
L6497 := __e.Get(1)
_ = L6497
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6498 := __e.Get(1)
_ = Key6498
__e.Return(MakeNative(func(__e *ControlFlow) {
C6499 := __e.Get(1)
_ = C6499
tmp20725 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20726 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20725)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20725
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6500, tmp20726, B6496, L6497, Key6498, C6499)
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

tmp20727 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20728 := Call(__e, PrimFunc(symshen_4assoc_1_6), symoccurs_2, tmp20724, tmp20727)


tmp20729 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20728)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20728
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20729

tmp20730 := MakeNative(func(__e *ControlFlow) {
V6505 := __e.Get(1)
_ = V6505
__e.Return(MakeNative(func(__e *ControlFlow) {
B6501 := __e.Get(1)
_ = B6501
__e.Return(MakeNative(func(__e *ControlFlow) {
L6502 := __e.Get(1)
_ = L6502
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6503 := __e.Get(1)
_ = Key6503
__e.Return(MakeNative(func(__e *ControlFlow) {
C6504 := __e.Get(1)
_ = C6504
tmp20731 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20732 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20731)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20731
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20733 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp20732)
}
__typedArg0 := symsymbol
__typedArg1 := tmp20732
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6505, tmp20733, B6501, L6502, Key6503, C6504)
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

tmp20734 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20735 := Call(__e, PrimFunc(symshen_4assoc_1_6), symoptimise, tmp20730, tmp20734)


tmp20736 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20735)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20735
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20736

tmp20737 := MakeNative(func(__e *ControlFlow) {
V6510 := __e.Get(1)
_ = V6510
__e.Return(MakeNative(func(__e *ControlFlow) {
B6506 := __e.Get(1)
_ = B6506
__e.Return(MakeNative(func(__e *ControlFlow) {
L6507 := __e.Get(1)
_ = L6507
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6508 := __e.Get(1)
_ = Key6508
__e.Return(MakeNative(func(__e *ControlFlow) {
C6509 := __e.Get(1)
_ = C6509
tmp20738 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20739 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20738)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20738
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6510, tmp20739, B6506, L6507, Key6508, C6509)
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

tmp20740 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20741 := Call(__e, PrimFunc(symshen_4assoc_1_6), symoptimise_2, tmp20737, tmp20740)


tmp20742 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20741)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20741
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20742

tmp20743 := MakeNative(func(__e *ControlFlow) {
V6515 := __e.Get(1)
_ = V6515
__e.Return(MakeNative(func(__e *ControlFlow) {
B6511 := __e.Get(1)
_ = B6511
__e.Return(MakeNative(func(__e *ControlFlow) {
L6512 := __e.Get(1)
_ = L6512
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6513 := __e.Get(1)
_ = Key6513
__e.Return(MakeNative(func(__e *ControlFlow) {
C6514 := __e.Get(1)
_ = C6514
tmp20744 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20745 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20744)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20744
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20746 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, tmp20745)
}
__typedArg0 := symboolean
__typedArg1 := tmp20745
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20747 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20746, Nil)
}
__typedArg0 := tmp20746
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20748 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20747)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20747
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20749 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, tmp20748)
}
__typedArg0 := symboolean
__typedArg1 := tmp20748
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6515, tmp20749, B6511, L6512, Key6513, C6514)
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

tmp20750 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20751 := Call(__e, PrimFunc(symshen_4assoc_1_6), symor, tmp20743, tmp20750)


tmp20752 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20751)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20751
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20752

tmp20753 := MakeNative(func(__e *ControlFlow) {
V6520 := __e.Get(1)
_ = V6520
__e.Return(MakeNative(func(__e *ControlFlow) {
B6516 := __e.Get(1)
_ = B6516
__e.Return(MakeNative(func(__e *ControlFlow) {
L6517 := __e.Get(1)
_ = L6517
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6518 := __e.Get(1)
_ = Key6518
__e.Return(MakeNative(func(__e *ControlFlow) {
C6519 := __e.Get(1)
_ = C6519
tmp20754 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20755 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20754)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20754
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6520, tmp20755, B6516, L6517, Key6518, C6519)
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

tmp20756 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20757 := Call(__e, PrimFunc(symshen_4assoc_1_6), symos, tmp20753, tmp20756)


tmp20758 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20757)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20757
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20758

tmp20759 := MakeNative(func(__e *ControlFlow) {
V6525 := __e.Get(1)
_ = V6525
__e.Return(MakeNative(func(__e *ControlFlow) {
B6521 := __e.Get(1)
_ = B6521
__e.Return(MakeNative(func(__e *ControlFlow) {
L6522 := __e.Get(1)
_ = L6522
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6523 := __e.Get(1)
_ = Key6523
__e.Return(MakeNative(func(__e *ControlFlow) {
C6524 := __e.Get(1)
_ = C6524
tmp20760 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20761 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20760)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20760
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20762 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp20761)
}
__typedArg0 := symsymbol
__typedArg1 := tmp20761
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6525, tmp20762, B6521, L6522, Key6523, C6524)
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

tmp20763 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20764 := Call(__e, PrimFunc(symshen_4assoc_1_6), sympackage_2, tmp20759, tmp20763)


tmp20765 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20764)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20764
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20765

tmp20766 := MakeNative(func(__e *ControlFlow) {
V6530 := __e.Get(1)
_ = V6530
__e.Return(MakeNative(func(__e *ControlFlow) {
B6526 := __e.Get(1)
_ = B6526
__e.Return(MakeNative(func(__e *ControlFlow) {
L6527 := __e.Get(1)
_ = L6527
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6528 := __e.Get(1)
_ = Key6528
__e.Return(MakeNative(func(__e *ControlFlow) {
C6529 := __e.Get(1)
_ = C6529
tmp20767 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20768 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20767)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20767
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6530, tmp20768, B6526, L6527, Key6528, C6529)
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

tmp20769 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20770 := Call(__e, PrimFunc(symshen_4assoc_1_6), symport, tmp20766, tmp20769)


tmp20771 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20770)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20770
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20771

tmp20772 := MakeNative(func(__e *ControlFlow) {
V6535 := __e.Get(1)
_ = V6535
__e.Return(MakeNative(func(__e *ControlFlow) {
B6531 := __e.Get(1)
_ = B6531
__e.Return(MakeNative(func(__e *ControlFlow) {
L6532 := __e.Get(1)
_ = L6532
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6533 := __e.Get(1)
_ = Key6533
__e.Return(MakeNative(func(__e *ControlFlow) {
C6534 := __e.Get(1)
_ = C6534
tmp20773 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20774 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20773)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20773
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6535, tmp20774, B6531, L6532, Key6533, C6534)
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

tmp20775 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20776 := Call(__e, PrimFunc(symshen_4assoc_1_6), symporters, tmp20772, tmp20775)


tmp20777 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20776)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20776
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20777

tmp20778 := MakeNative(func(__e *ControlFlow) {
V6540 := __e.Get(1)
_ = V6540
__e.Return(MakeNative(func(__e *ControlFlow) {
B6536 := __e.Get(1)
_ = B6536
__e.Return(MakeNative(func(__e *ControlFlow) {
L6537 := __e.Get(1)
_ = L6537
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6538 := __e.Get(1)
_ = Key6538
__e.Return(MakeNative(func(__e *ControlFlow) {
C6539 := __e.Get(1)
_ = C6539
tmp20779 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20780 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20779)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20779
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20781 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp20780)
}
__typedArg0 := symnumber
__typedArg1 := tmp20780
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20782 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20781, Nil)
}
__typedArg0 := tmp20781
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20783 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20782)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20782
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20784 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp20783)
}
__typedArg0 := symstring
__typedArg1 := tmp20783
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6540, tmp20784, B6536, L6537, Key6538, C6539)
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

tmp20785 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20786 := Call(__e, PrimFunc(symshen_4assoc_1_6), sympos, tmp20778, tmp20785)


tmp20787 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20786)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20786
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20787

tmp20788 := MakeNative(func(__e *ControlFlow) {
V6545 := __e.Get(1)
_ = V6545
__e.Return(MakeNative(func(__e *ControlFlow) {
B6541 := __e.Get(1)
_ = B6541
__e.Return(MakeNative(func(__e *ControlFlow) {
L6542 := __e.Get(1)
_ = L6542
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6543 := __e.Get(1)
_ = Key6543
__e.Return(MakeNative(func(__e *ControlFlow) {
C6544 := __e.Get(1)
_ = C6544
tmp20789 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symout, Nil)
}
__typedArg0 := symout
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20790 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp20789)
}
__typedArg0 := symstream
__typedArg1 := tmp20789
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20791 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20792 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20791)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20791
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20793 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20790, tmp20792)
}
__typedArg0 := tmp20790
__typedArg1 := tmp20792
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20794 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20793, Nil)
}
__typedArg0 := tmp20793
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20795 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20794)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20794
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20796 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp20795)
}
__typedArg0 := symstring
__typedArg1 := tmp20795
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6545, tmp20796, B6541, L6542, Key6543, C6544)
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

tmp20797 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20798 := Call(__e, PrimFunc(symshen_4assoc_1_6), sympr, tmp20788, tmp20797)


tmp20799 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20798)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20798
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20799

tmp20800 := MakeNative(func(__e *ControlFlow) {
V6550 := __e.Get(1)
_ = V6550
__e.Return(MakeNative(func(__e *ControlFlow) {
B6546 := __e.Get(1)
_ = B6546
__e.Return(MakeNative(func(__e *ControlFlow) {
L6547 := __e.Get(1)
_ = L6547
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6548 := __e.Get(1)
_ = Key6548
__e.Return(MakeNative(func(__e *ControlFlow) {
C6549 := __e.Get(1)
_ = C6549
tmp20801 := Call(__e, PrimFunc(symshen_4newpv), B6546)


let__18961 := tmp20801
_ = let__18961

tmp20802 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18961, Nil)
}
__typedArg0 := let__18961
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20803 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20802)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20802
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20804 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18961, tmp20803)
}
__typedArg0 := let__18961
__typedArg1 := tmp20803
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20805 := Call(__e, PrimFunc(symis_b), V6550, tmp20804, B6546, L6547, Key6548, C6549)


__e.TailApply(PrimFunc(symshen_4gc), B6546, tmp20805)
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

tmp20806 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20807 := Call(__e, PrimFunc(symshen_4assoc_1_6), symprint, tmp20800, tmp20806)


tmp20808 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20807)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20807
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20808

tmp20809 := MakeNative(func(__e *ControlFlow) {
V6555 := __e.Get(1)
_ = V6555
__e.Return(MakeNative(func(__e *ControlFlow) {
B6551 := __e.Get(1)
_ = B6551
__e.Return(MakeNative(func(__e *ControlFlow) {
L6552 := __e.Get(1)
_ = L6552
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6553 := __e.Get(1)
_ = Key6553
__e.Return(MakeNative(func(__e *ControlFlow) {
C6554 := __e.Get(1)
_ = C6554
tmp20810 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20811 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20810)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20810
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20812 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp20811)
}
__typedArg0 := symsymbol
__typedArg1 := tmp20811
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6555, tmp20812, B6551, L6552, Key6553, C6554)
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

tmp20813 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20814 := Call(__e, PrimFunc(symshen_4assoc_1_6), symprofile, tmp20809, tmp20813)


tmp20815 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20814)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20814
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20815

tmp20816 := MakeNative(func(__e *ControlFlow) {
V6561 := __e.Get(1)
_ = V6561
__e.Return(MakeNative(func(__e *ControlFlow) {
B6557 := __e.Get(1)
_ = B6557
__e.Return(MakeNative(func(__e *ControlFlow) {
L6558 := __e.Get(1)
_ = L6558
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6559 := __e.Get(1)
_ = Key6559
__e.Return(MakeNative(func(__e *ControlFlow) {
C6560 := __e.Get(1)
_ = C6560
tmp20817 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20818 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20817)
}
__typedArg0 := symlist
__typedArg1 := tmp20817
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20819 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20820 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20819)
}
__typedArg0 := symlist
__typedArg1 := tmp20819
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20821 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20820, Nil)
}
__typedArg0 := tmp20820
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20822 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20821)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20821
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20823 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20818, tmp20822)
}
__typedArg0 := tmp20818
__typedArg1 := tmp20822
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6561, tmp20823, B6557, L6558, Key6559, C6560)
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

tmp20824 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20825 := Call(__e, PrimFunc(symshen_4assoc_1_6), sympreclude, tmp20816, tmp20824)


tmp20826 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20825)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20825
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20826

tmp20827 := MakeNative(func(__e *ControlFlow) {
V6566 := __e.Get(1)
_ = V6566
__e.Return(MakeNative(func(__e *ControlFlow) {
B6562 := __e.Get(1)
_ = B6562
__e.Return(MakeNative(func(__e *ControlFlow) {
L6563 := __e.Get(1)
_ = L6563
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6564 := __e.Get(1)
_ = Key6564
__e.Return(MakeNative(func(__e *ControlFlow) {
C6565 := __e.Get(1)
_ = C6565
tmp20828 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20829 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20828)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20828
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20830 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp20829)
}
__typedArg0 := symstring
__typedArg1 := tmp20829
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6566, tmp20830, B6562, L6563, Key6564, C6565)
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

tmp20831 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20832 := Call(__e, PrimFunc(symshen_4assoc_1_6), symshen_4proc_1nl, tmp20827, tmp20831)


tmp20833 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20832)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20832
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20833

tmp20834 := MakeNative(func(__e *ControlFlow) {
V6571 := __e.Get(1)
_ = V6571
__e.Return(MakeNative(func(__e *ControlFlow) {
B6567 := __e.Get(1)
_ = B6567
__e.Return(MakeNative(func(__e *ControlFlow) {
L6568 := __e.Get(1)
_ = L6568
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6569 := __e.Get(1)
_ = Key6569
__e.Return(MakeNative(func(__e *ControlFlow) {
C6570 := __e.Get(1)
_ = C6570
tmp20835 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20836 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_d, tmp20835)
}
__typedArg0 := sym_d
__typedArg1 := tmp20835
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20837 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp20836)
}
__typedArg0 := symsymbol
__typedArg1 := tmp20836
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20838 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20837, Nil)
}
__typedArg0 := tmp20837
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20839 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20838)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20838
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20840 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp20839)
}
__typedArg0 := symsymbol
__typedArg1 := tmp20839
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6571, tmp20840, B6567, L6568, Key6569, C6570)
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

tmp20841 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20842 := Call(__e, PrimFunc(symshen_4assoc_1_6), symprofile_1results, tmp20834, tmp20841)


tmp20843 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20842)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20842
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20843

tmp20844 := MakeNative(func(__e *ControlFlow) {
V6576 := __e.Get(1)
_ = V6576
__e.Return(MakeNative(func(__e *ControlFlow) {
B6572 := __e.Get(1)
_ = B6572
__e.Return(MakeNative(func(__e *ControlFlow) {
L6573 := __e.Get(1)
_ = L6573
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6574 := __e.Get(1)
_ = Key6574
__e.Return(MakeNative(func(__e *ControlFlow) {
C6575 := __e.Get(1)
_ = C6575
tmp20845 := Call(__e, PrimFunc(symshen_4newpv), B6572)


let__18962 := tmp20845
_ = let__18962

tmp20846 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18962, Nil)
}
__typedArg0 := let__18962
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20847 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20846)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20846
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20848 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18962, tmp20847)
}
__typedArg0 := let__18962
__typedArg1 := tmp20847
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20849 := Call(__e, PrimFunc(symis_b), V6576, tmp20848, B6572, L6573, Key6574, C6575)


__e.TailApply(PrimFunc(symshen_4gc), B6572, tmp20849)
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

tmp20850 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20851 := Call(__e, PrimFunc(symshen_4assoc_1_6), symprotect, tmp20844, tmp20850)


tmp20852 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20851)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20851
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20852

tmp20853 := MakeNative(func(__e *ControlFlow) {
V6582 := __e.Get(1)
_ = V6582
__e.Return(MakeNative(func(__e *ControlFlow) {
B6578 := __e.Get(1)
_ = B6578
__e.Return(MakeNative(func(__e *ControlFlow) {
L6579 := __e.Get(1)
_ = L6579
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6580 := __e.Get(1)
_ = Key6580
__e.Return(MakeNative(func(__e *ControlFlow) {
C6581 := __e.Get(1)
_ = C6581
tmp20854 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20855 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20854)
}
__typedArg0 := symlist
__typedArg1 := tmp20854
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20856 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20857 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20856)
}
__typedArg0 := symlist
__typedArg1 := tmp20856
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20858 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20857, Nil)
}
__typedArg0 := tmp20857
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20859 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20858)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20858
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20860 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20855, tmp20859)
}
__typedArg0 := tmp20855
__typedArg1 := tmp20859
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6582, tmp20860, B6578, L6579, Key6580, C6581)
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

tmp20861 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20862 := Call(__e, PrimFunc(symshen_4assoc_1_6), sympreclude_1all_1but, tmp20853, tmp20861)


tmp20863 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20862)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20862
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20863

tmp20864 := MakeNative(func(__e *ControlFlow) {
V6587 := __e.Get(1)
_ = V6587
__e.Return(MakeNative(func(__e *ControlFlow) {
B6583 := __e.Get(1)
_ = B6583
__e.Return(MakeNative(func(__e *ControlFlow) {
L6584 := __e.Get(1)
_ = L6584
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6585 := __e.Get(1)
_ = Key6585
__e.Return(MakeNative(func(__e *ControlFlow) {
C6586 := __e.Get(1)
_ = C6586
tmp20865 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symout, Nil)
}
__typedArg0 := symout
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20866 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp20865)
}
__typedArg0 := symstream
__typedArg1 := tmp20865
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20867 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20868 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20867)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20867
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20869 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20866, tmp20868)
}
__typedArg0 := tmp20866
__typedArg1 := tmp20868
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20870 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20869, Nil)
}
__typedArg0 := tmp20869
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20871 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20870)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20870
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20872 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp20871)
}
__typedArg0 := symstring
__typedArg1 := tmp20871
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6587, tmp20872, B6583, L6584, Key6585, C6586)
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

tmp20873 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20874 := Call(__e, PrimFunc(symshen_4assoc_1_6), symshen_4prhush, tmp20864, tmp20873)


tmp20875 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20874)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20874
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20875

tmp20876 := MakeNative(func(__e *ControlFlow) {
V6592 := __e.Get(1)
_ = V6592
__e.Return(MakeNative(func(__e *ControlFlow) {
B6588 := __e.Get(1)
_ = B6588
__e.Return(MakeNative(func(__e *ControlFlow) {
L6589 := __e.Get(1)
_ = L6589
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6590 := __e.Get(1)
_ = Key6590
__e.Return(MakeNative(func(__e *ControlFlow) {
C6591 := __e.Get(1)
_ = C6591
tmp20877 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20878 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20877)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20877
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20879 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp20878)
}
__typedArg0 := symnumber
__typedArg1 := tmp20878
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6592, tmp20879, B6588, L6589, Key6590, C6591)
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

tmp20880 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20881 := Call(__e, PrimFunc(symshen_4assoc_1_6), symprolog_1memory, tmp20876, tmp20880)


tmp20882 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20881)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20881
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20882

tmp20883 := MakeNative(func(__e *ControlFlow) {
V6597 := __e.Get(1)
_ = V6597
__e.Return(MakeNative(func(__e *ControlFlow) {
B6593 := __e.Get(1)
_ = B6593
__e.Return(MakeNative(func(__e *ControlFlow) {
L6594 := __e.Get(1)
_ = L6594
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6595 := __e.Get(1)
_ = Key6595
__e.Return(MakeNative(func(__e *ControlFlow) {
C6596 := __e.Get(1)
_ = C6596
tmp20884 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunit, Nil)
}
__typedArg0 := symunit
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20885 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20884)
}
__typedArg0 := symlist
__typedArg1 := tmp20884
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20886 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20885, Nil)
}
__typedArg0 := tmp20885
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20887 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20886)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20886
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20888 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp20887)
}
__typedArg0 := symsymbol
__typedArg1 := tmp20887
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6597, tmp20888, B6593, L6594, Key6595, C6596)
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

tmp20889 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20890 := Call(__e, PrimFunc(symshen_4assoc_1_6), symps, tmp20883, tmp20889)


tmp20891 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20890)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20890
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20891

tmp20892 := MakeNative(func(__e *ControlFlow) {
V6602 := __e.Get(1)
_ = V6602
__e.Return(MakeNative(func(__e *ControlFlow) {
B6598 := __e.Get(1)
_ = B6598
__e.Return(MakeNative(func(__e *ControlFlow) {
L6599 := __e.Get(1)
_ = L6599
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6600 := __e.Get(1)
_ = Key6600
__e.Return(MakeNative(func(__e *ControlFlow) {
C6601 := __e.Get(1)
_ = C6601
tmp20893 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symin, Nil)
}
__typedArg0 := symin
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20894 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp20893)
}
__typedArg0 := symstream
__typedArg1 := tmp20893
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20895 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunit, Nil)
}
__typedArg0 := symunit
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20896 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20895)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20895
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20897 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20894, tmp20896)
}
__typedArg0 := tmp20894
__typedArg1 := tmp20896
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6602, tmp20897, B6598, L6599, Key6600, C6601)
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

tmp20898 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20899 := Call(__e, PrimFunc(symshen_4assoc_1_6), symread, tmp20892, tmp20898)


tmp20900 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20899)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20899
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20900

tmp20901 := MakeNative(func(__e *ControlFlow) {
V6607 := __e.Get(1)
_ = V6607
__e.Return(MakeNative(func(__e *ControlFlow) {
B6603 := __e.Get(1)
_ = B6603
__e.Return(MakeNative(func(__e *ControlFlow) {
L6604 := __e.Get(1)
_ = L6604
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6605 := __e.Get(1)
_ = Key6605
__e.Return(MakeNative(func(__e *ControlFlow) {
C6606 := __e.Get(1)
_ = C6606
tmp20902 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symin, Nil)
}
__typedArg0 := symin
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20903 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp20902)
}
__typedArg0 := symstream
__typedArg1 := tmp20902
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20904 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20905 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20904)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20904
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20906 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20903, tmp20905)
}
__typedArg0 := tmp20903
__typedArg1 := tmp20905
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6607, tmp20906, B6603, L6604, Key6605, C6606)
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

tmp20907 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20908 := Call(__e, PrimFunc(symshen_4assoc_1_6), symread_1unit_1string, tmp20901, tmp20907)


tmp20909 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20908)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20908
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20909

tmp20910 := MakeNative(func(__e *ControlFlow) {
V6612 := __e.Get(1)
_ = V6612
__e.Return(MakeNative(func(__e *ControlFlow) {
B6608 := __e.Get(1)
_ = B6608
__e.Return(MakeNative(func(__e *ControlFlow) {
L6609 := __e.Get(1)
_ = L6609
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6610 := __e.Get(1)
_ = Key6610
__e.Return(MakeNative(func(__e *ControlFlow) {
C6611 := __e.Get(1)
_ = C6611
tmp20911 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symin, Nil)
}
__typedArg0 := symin
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20912 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp20911)
}
__typedArg0 := symstream
__typedArg1 := tmp20911
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20913 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20914 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20913)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20913
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20915 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20912, tmp20914)
}
__typedArg0 := tmp20912
__typedArg1 := tmp20914
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6612, tmp20915, B6608, L6609, Key6610, C6611)
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

tmp20916 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20917 := Call(__e, PrimFunc(symshen_4assoc_1_6), symread_1byte, tmp20910, tmp20916)


tmp20918 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20917)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20917
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20918

tmp20919 := MakeNative(func(__e *ControlFlow) {
V6617 := __e.Get(1)
_ = V6617
__e.Return(MakeNative(func(__e *ControlFlow) {
B6613 := __e.Get(1)
_ = B6613
__e.Return(MakeNative(func(__e *ControlFlow) {
L6614 := __e.Get(1)
_ = L6614
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6615 := __e.Get(1)
_ = Key6615
__e.Return(MakeNative(func(__e *ControlFlow) {
C6616 := __e.Get(1)
_ = C6616
tmp20920 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20921 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20920)
}
__typedArg0 := symlist
__typedArg1 := tmp20920
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20922 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20921, Nil)
}
__typedArg0 := tmp20921
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20923 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20922)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20922
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20924 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp20923)
}
__typedArg0 := symstring
__typedArg1 := tmp20923
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6617, tmp20924, B6613, L6614, Key6615, C6616)
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

tmp20925 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20926 := Call(__e, PrimFunc(symshen_4assoc_1_6), symread_1file_1as_1bytelist, tmp20919, tmp20925)


tmp20927 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20926)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20926
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20927

tmp20928 := MakeNative(func(__e *ControlFlow) {
V6622 := __e.Get(1)
_ = V6622
__e.Return(MakeNative(func(__e *ControlFlow) {
B6618 := __e.Get(1)
_ = B6618
__e.Return(MakeNative(func(__e *ControlFlow) {
L6619 := __e.Get(1)
_ = L6619
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6620 := __e.Get(1)
_ = Key6620
__e.Return(MakeNative(func(__e *ControlFlow) {
C6621 := __e.Get(1)
_ = C6621
tmp20929 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20930 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20929)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20929
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20931 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp20930)
}
__typedArg0 := symstring
__typedArg1 := tmp20930
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6622, tmp20931, B6618, L6619, Key6620, C6621)
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

tmp20932 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20933 := Call(__e, PrimFunc(symshen_4assoc_1_6), symread_1file_1as_1string, tmp20928, tmp20932)


tmp20934 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20933)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20933
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20934

tmp20935 := MakeNative(func(__e *ControlFlow) {
V6627 := __e.Get(1)
_ = V6627
__e.Return(MakeNative(func(__e *ControlFlow) {
B6623 := __e.Get(1)
_ = B6623
__e.Return(MakeNative(func(__e *ControlFlow) {
L6624 := __e.Get(1)
_ = L6624
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6625 := __e.Get(1)
_ = Key6625
__e.Return(MakeNative(func(__e *ControlFlow) {
C6626 := __e.Get(1)
_ = C6626
tmp20936 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunit, Nil)
}
__typedArg0 := symunit
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20937 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20936)
}
__typedArg0 := symlist
__typedArg1 := tmp20936
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20938 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20937, Nil)
}
__typedArg0 := tmp20937
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20939 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20938)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20938
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20940 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp20939)
}
__typedArg0 := symstring
__typedArg1 := tmp20939
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6627, tmp20940, B6623, L6624, Key6625, C6626)
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

tmp20941 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20942 := Call(__e, PrimFunc(symshen_4assoc_1_6), symread_1file, tmp20935, tmp20941)


tmp20943 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20942)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20942
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20943

tmp20944 := MakeNative(func(__e *ControlFlow) {
V6632 := __e.Get(1)
_ = V6632
__e.Return(MakeNative(func(__e *ControlFlow) {
B6628 := __e.Get(1)
_ = B6628
__e.Return(MakeNative(func(__e *ControlFlow) {
L6629 := __e.Get(1)
_ = L6629
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6630 := __e.Get(1)
_ = Key6630
__e.Return(MakeNative(func(__e *ControlFlow) {
C6631 := __e.Get(1)
_ = C6631
tmp20945 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunit, Nil)
}
__typedArg0 := symunit
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20946 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20945)
}
__typedArg0 := symlist
__typedArg1 := tmp20945
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20947 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20946, Nil)
}
__typedArg0 := tmp20946
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20948 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20947)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20947
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20949 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp20948)
}
__typedArg0 := symstring
__typedArg1 := tmp20948
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6632, tmp20949, B6628, L6629, Key6630, C6631)
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

tmp20950 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20951 := Call(__e, PrimFunc(symshen_4assoc_1_6), symread_1from_1string, tmp20944, tmp20950)


tmp20952 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20951)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20951
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20952

tmp20953 := MakeNative(func(__e *ControlFlow) {
V6637 := __e.Get(1)
_ = V6637
__e.Return(MakeNative(func(__e *ControlFlow) {
B6633 := __e.Get(1)
_ = B6633
__e.Return(MakeNative(func(__e *ControlFlow) {
L6634 := __e.Get(1)
_ = L6634
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6635 := __e.Get(1)
_ = Key6635
__e.Return(MakeNative(func(__e *ControlFlow) {
C6636 := __e.Get(1)
_ = C6636
tmp20954 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunit, Nil)
}
__typedArg0 := symunit
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20955 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20954)
}
__typedArg0 := symlist
__typedArg1 := tmp20954
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20956 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20955, Nil)
}
__typedArg0 := tmp20955
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20957 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20956)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20956
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20958 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp20957)
}
__typedArg0 := symstring
__typedArg1 := tmp20957
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6637, tmp20958, B6633, L6634, Key6635, C6636)
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

tmp20959 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20960 := Call(__e, PrimFunc(symshen_4assoc_1_6), symread_1from_1string_1unprocessed, tmp20953, tmp20959)


tmp20961 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20960)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20960
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20961

tmp20962 := MakeNative(func(__e *ControlFlow) {
V6642 := __e.Get(1)
_ = V6642
__e.Return(MakeNative(func(__e *ControlFlow) {
B6638 := __e.Get(1)
_ = B6638
__e.Return(MakeNative(func(__e *ControlFlow) {
L6639 := __e.Get(1)
_ = L6639
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6640 := __e.Get(1)
_ = Key6640
__e.Return(MakeNative(func(__e *ControlFlow) {
C6641 := __e.Get(1)
_ = C6641
tmp20963 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20964 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20963)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20963
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6642, tmp20964, B6638, L6639, Key6640, C6641)
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

tmp20966 := Call(__e, PrimFunc(symshen_4assoc_1_6), symrelease, tmp20962, tmp20965)


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
V6648 := __e.Get(1)
_ = V6648
__e.Return(MakeNative(func(__e *ControlFlow) {
B6644 := __e.Get(1)
_ = B6644
__e.Return(MakeNative(func(__e *ControlFlow) {
L6645 := __e.Get(1)
_ = L6645
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6646 := __e.Get(1)
_ = Key6646
__e.Return(MakeNative(func(__e *ControlFlow) {
C6647 := __e.Get(1)
_ = C6647
tmp20969 := Call(__e, PrimFunc(symshen_4newpv), B6644)


let__18963 := tmp20969
_ = let__18963

tmp20970 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18963, Nil)
}
__typedArg0 := let__18963
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20971 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20970)
}
__typedArg0 := symlist
__typedArg1 := tmp20970
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20972 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18963, Nil)
}
__typedArg0 := let__18963
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20973 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20972)
}
__typedArg0 := symlist
__typedArg1 := tmp20972
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20974 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20973, Nil)
}
__typedArg0 := tmp20973
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20975 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20974)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20974
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20976 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20971, tmp20975)
}
__typedArg0 := tmp20971
__typedArg1 := tmp20975
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20977 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20976, Nil)
}
__typedArg0 := tmp20976
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20978 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20977)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20977
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20979 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18963, tmp20978)
}
__typedArg0 := let__18963
__typedArg1 := tmp20978
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20980 := Call(__e, PrimFunc(symis_b), V6648, tmp20979, B6644, L6645, Key6646, C6647)


__e.TailApply(PrimFunc(symshen_4gc), B6644, tmp20980)
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

tmp20981 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20982 := Call(__e, PrimFunc(symshen_4assoc_1_6), symremove, tmp20968, tmp20981)


tmp20983 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20982)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20982
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20983

tmp20984 := MakeNative(func(__e *ControlFlow) {
V6654 := __e.Get(1)
_ = V6654
__e.Return(MakeNative(func(__e *ControlFlow) {
B6650 := __e.Get(1)
_ = B6650
__e.Return(MakeNative(func(__e *ControlFlow) {
L6651 := __e.Get(1)
_ = L6651
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6652 := __e.Get(1)
_ = Key6652
__e.Return(MakeNative(func(__e *ControlFlow) {
C6653 := __e.Get(1)
_ = C6653
tmp20985 := Call(__e, PrimFunc(symshen_4newpv), B6650)


let__18964 := tmp20985
_ = let__18964

tmp20986 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18964, Nil)
}
__typedArg0 := let__18964
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20987 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20986)
}
__typedArg0 := symlist
__typedArg1 := tmp20986
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20988 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18964, Nil)
}
__typedArg0 := let__18964
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20989 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20988)
}
__typedArg0 := symlist
__typedArg1 := tmp20988
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20990 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20989, Nil)
}
__typedArg0 := tmp20989
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20991 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20990)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20990
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20992 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20987, tmp20991)
}
__typedArg0 := tmp20987
__typedArg1 := tmp20991
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20993 := Call(__e, PrimFunc(symis_b), V6654, tmp20992, B6650, L6651, Key6652, C6653)


__e.TailApply(PrimFunc(symshen_4gc), B6650, tmp20993)
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

tmp20994 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20995 := Call(__e, PrimFunc(symshen_4assoc_1_6), symreverse, tmp20984, tmp20994)


tmp20996 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20995)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20995
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20996

tmp20997 := MakeNative(func(__e *ControlFlow) {
V6659 := __e.Get(1)
_ = V6659
__e.Return(MakeNative(func(__e *ControlFlow) {
B6655 := __e.Get(1)
_ = B6655
__e.Return(MakeNative(func(__e *ControlFlow) {
L6656 := __e.Get(1)
_ = L6656
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6657 := __e.Get(1)
_ = Key6657
__e.Return(MakeNative(func(__e *ControlFlow) {
C6658 := __e.Get(1)
_ = C6658
tmp20998 := Call(__e, PrimFunc(symshen_4newpv), B6655)


let__18965 := tmp20998
_ = let__18965

tmp20999 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18965, Nil)
}
__typedArg0 := let__18965
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21000 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20999)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20999
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21001 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp21000)
}
__typedArg0 := symstring
__typedArg1 := tmp21000
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21002 := Call(__e, PrimFunc(symis_b), V6659, tmp21001, B6655, L6656, Key6657, C6658)


__e.TailApply(PrimFunc(symshen_4gc), B6655, tmp21002)
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

tmp21003 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21004 := Call(__e, PrimFunc(symshen_4assoc_1_6), symsimple_1error, tmp20997, tmp21003)


tmp21005 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21004)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21004
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21005

tmp21006 := MakeNative(func(__e *ControlFlow) {
V6664 := __e.Get(1)
_ = V6664
__e.Return(MakeNative(func(__e *ControlFlow) {
B6660 := __e.Get(1)
_ = B6660
__e.Return(MakeNative(func(__e *ControlFlow) {
L6661 := __e.Get(1)
_ = L6661
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6662 := __e.Get(1)
_ = Key6662
__e.Return(MakeNative(func(__e *ControlFlow) {
C6663 := __e.Get(1)
_ = C6663
tmp21007 := Call(__e, PrimFunc(symshen_4newpv), B6660)


let__18966 := tmp21007
_ = let__18966

tmp21008 := Call(__e, PrimFunc(symshen_4newpv), B6660)


let__18967 := tmp21008
_ = let__18967

tmp21009 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18967, Nil)
}
__typedArg0 := let__18967
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21010 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_d, tmp21009)
}
__typedArg0 := sym_d
__typedArg1 := tmp21009
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21011 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18966, tmp21010)
}
__typedArg0 := let__18966
__typedArg1 := tmp21010
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21012 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18967, Nil)
}
__typedArg0 := let__18967
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21013 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21012)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21012
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21014 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21011, tmp21013)
}
__typedArg0 := tmp21011
__typedArg1 := tmp21013
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21015 := Call(__e, PrimFunc(symis_b), V6664, tmp21014, B6660, L6661, Key6662, C6663)


tmp21016 := Call(__e, PrimFunc(symshen_4gc), B6660, tmp21015)


__e.TailApply(PrimFunc(symshen_4gc), B6660, tmp21016)
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

tmp21017 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21018 := Call(__e, PrimFunc(symshen_4assoc_1_6), symsnd, tmp21006, tmp21017)


tmp21019 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21018)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21018
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21019

tmp21020 := MakeNative(func(__e *ControlFlow) {
V6669 := __e.Get(1)
_ = V6669
__e.Return(MakeNative(func(__e *ControlFlow) {
B6665 := __e.Get(1)
_ = B6665
__e.Return(MakeNative(func(__e *ControlFlow) {
L6666 := __e.Get(1)
_ = L6666
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6667 := __e.Get(1)
_ = Key6667
__e.Return(MakeNative(func(__e *ControlFlow) {
C6668 := __e.Get(1)
_ = C6668
tmp21021 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21022 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21021)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21021
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21023 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp21022)
}
__typedArg0 := symnumber
__typedArg1 := tmp21022
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21024 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21023, Nil)
}
__typedArg0 := tmp21023
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21025 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21024)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21024
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21026 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp21025)
}
__typedArg0 := symsymbol
__typedArg1 := tmp21025
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6669, tmp21026, B6665, L6666, Key6667, C6668)
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

tmp21028 := Call(__e, PrimFunc(symshen_4assoc_1_6), symspecialise, tmp21020, tmp21027)


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
V6674 := __e.Get(1)
_ = V6674
__e.Return(MakeNative(func(__e *ControlFlow) {
B6670 := __e.Get(1)
_ = B6670
__e.Return(MakeNative(func(__e *ControlFlow) {
L6671 := __e.Get(1)
_ = L6671
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6672 := __e.Get(1)
_ = Key6672
__e.Return(MakeNative(func(__e *ControlFlow) {
C6673 := __e.Get(1)
_ = C6673
tmp21031 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21032 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21031)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21031
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21033 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp21032)
}
__typedArg0 := symsymbol
__typedArg1 := tmp21032
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6674, tmp21033, B6670, L6671, Key6672, C6673)
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

tmp21034 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21035 := Call(__e, PrimFunc(symshen_4assoc_1_6), symspy, tmp21030, tmp21034)


tmp21036 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21035)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21035
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21036

tmp21037 := MakeNative(func(__e *ControlFlow) {
V6679 := __e.Get(1)
_ = V6679
__e.Return(MakeNative(func(__e *ControlFlow) {
B6675 := __e.Get(1)
_ = B6675
__e.Return(MakeNative(func(__e *ControlFlow) {
L6676 := __e.Get(1)
_ = L6676
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6677 := __e.Get(1)
_ = Key6677
__e.Return(MakeNative(func(__e *ControlFlow) {
C6678 := __e.Get(1)
_ = C6678
tmp21038 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21039 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21038)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21038
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6679, tmp21039, B6675, L6676, Key6677, C6678)
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

tmp21040 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21041 := Call(__e, PrimFunc(symshen_4assoc_1_6), symspy_2, tmp21037, tmp21040)


tmp21042 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21041)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21041
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21042

tmp21043 := MakeNative(func(__e *ControlFlow) {
V6684 := __e.Get(1)
_ = V6684
__e.Return(MakeNative(func(__e *ControlFlow) {
B6680 := __e.Get(1)
_ = B6680
__e.Return(MakeNative(func(__e *ControlFlow) {
L6681 := __e.Get(1)
_ = L6681
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6682 := __e.Get(1)
_ = Key6682
__e.Return(MakeNative(func(__e *ControlFlow) {
C6683 := __e.Get(1)
_ = C6683
tmp21044 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21045 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21044)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21044
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21046 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp21045)
}
__typedArg0 := symsymbol
__typedArg1 := tmp21045
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6684, tmp21046, B6680, L6681, Key6682, C6683)
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

tmp21047 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21048 := Call(__e, PrimFunc(symshen_4assoc_1_6), symstep, tmp21043, tmp21047)


tmp21049 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21048)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21048
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21049

tmp21050 := MakeNative(func(__e *ControlFlow) {
V6689 := __e.Get(1)
_ = V6689
__e.Return(MakeNative(func(__e *ControlFlow) {
B6685 := __e.Get(1)
_ = B6685
__e.Return(MakeNative(func(__e *ControlFlow) {
L6686 := __e.Get(1)
_ = L6686
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6687 := __e.Get(1)
_ = Key6687
__e.Return(MakeNative(func(__e *ControlFlow) {
C6688 := __e.Get(1)
_ = C6688
tmp21051 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21052 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21051)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21051
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6689, tmp21052, B6685, L6686, Key6687, C6688)
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

tmp21053 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21054 := Call(__e, PrimFunc(symshen_4assoc_1_6), symstep_2, tmp21050, tmp21053)


tmp21055 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21054)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21054
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21055

tmp21056 := MakeNative(func(__e *ControlFlow) {
V6694 := __e.Get(1)
_ = V6694
__e.Return(MakeNative(func(__e *ControlFlow) {
B6690 := __e.Get(1)
_ = B6690
__e.Return(MakeNative(func(__e *ControlFlow) {
L6691 := __e.Get(1)
_ = L6691
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6692 := __e.Get(1)
_ = Key6692
__e.Return(MakeNative(func(__e *ControlFlow) {
C6693 := __e.Get(1)
_ = C6693
tmp21057 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symin, Nil)
}
__typedArg0 := symin
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21058 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp21057)
}
__typedArg0 := symstream
__typedArg1 := tmp21057
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21059 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21058, Nil)
}
__typedArg0 := tmp21058
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21060 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21059)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21059
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6694, tmp21060, B6690, L6691, Key6692, C6693)
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

tmp21061 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21062 := Call(__e, PrimFunc(symshen_4assoc_1_6), symstinput, tmp21056, tmp21061)


tmp21063 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21062)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21062
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21063

tmp21064 := MakeNative(func(__e *ControlFlow) {
V6699 := __e.Get(1)
_ = V6699
__e.Return(MakeNative(func(__e *ControlFlow) {
B6695 := __e.Get(1)
_ = B6695
__e.Return(MakeNative(func(__e *ControlFlow) {
L6696 := __e.Get(1)
_ = L6696
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6697 := __e.Get(1)
_ = Key6697
__e.Return(MakeNative(func(__e *ControlFlow) {
C6698 := __e.Get(1)
_ = C6698
tmp21065 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symout, Nil)
}
__typedArg0 := symout
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21066 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp21065)
}
__typedArg0 := symstream
__typedArg1 := tmp21065
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21067 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21066, Nil)
}
__typedArg0 := tmp21066
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21068 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21067)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21067
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6699, tmp21068, B6695, L6696, Key6697, C6698)
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

tmp21069 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21070 := Call(__e, PrimFunc(symshen_4assoc_1_6), symsterror, tmp21064, tmp21069)


tmp21071 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21070)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21070
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21071

tmp21072 := MakeNative(func(__e *ControlFlow) {
V6704 := __e.Get(1)
_ = V6704
__e.Return(MakeNative(func(__e *ControlFlow) {
B6700 := __e.Get(1)
_ = B6700
__e.Return(MakeNative(func(__e *ControlFlow) {
L6701 := __e.Get(1)
_ = L6701
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6702 := __e.Get(1)
_ = Key6702
__e.Return(MakeNative(func(__e *ControlFlow) {
C6703 := __e.Get(1)
_ = C6703
tmp21073 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symout, Nil)
}
__typedArg0 := symout
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21074 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp21073)
}
__typedArg0 := symstream
__typedArg1 := tmp21073
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21075 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21074, Nil)
}
__typedArg0 := tmp21074
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21076 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21075)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21075
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6704, tmp21076, B6700, L6701, Key6702, C6703)
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

tmp21077 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21078 := Call(__e, PrimFunc(symshen_4assoc_1_6), symstoutput, tmp21072, tmp21077)


tmp21079 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21078)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21078
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21079

tmp21080 := MakeNative(func(__e *ControlFlow) {
V6709 := __e.Get(1)
_ = V6709
__e.Return(MakeNative(func(__e *ControlFlow) {
B6705 := __e.Get(1)
_ = B6705
__e.Return(MakeNative(func(__e *ControlFlow) {
L6706 := __e.Get(1)
_ = L6706
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6707 := __e.Get(1)
_ = Key6707
__e.Return(MakeNative(func(__e *ControlFlow) {
C6708 := __e.Get(1)
_ = C6708
tmp21081 := Call(__e, PrimFunc(symshen_4newpv), B6705)


let__18968 := tmp21081
_ = let__18968

tmp21082 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21083 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21082)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21082
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21084 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18968, tmp21083)
}
__typedArg0 := let__18968
__typedArg1 := tmp21083
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21085 := Call(__e, PrimFunc(symis_b), V6709, tmp21084, B6705, L6706, Key6707, C6708)


__e.TailApply(PrimFunc(symshen_4gc), B6705, tmp21085)
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

tmp21087 := Call(__e, PrimFunc(symshen_4assoc_1_6), symstring_2, tmp21080, tmp21086)


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
V6714 := __e.Get(1)
_ = V6714
__e.Return(MakeNative(func(__e *ControlFlow) {
B6710 := __e.Get(1)
_ = B6710
__e.Return(MakeNative(func(__e *ControlFlow) {
L6711 := __e.Get(1)
_ = L6711
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6712 := __e.Get(1)
_ = Key6712
__e.Return(MakeNative(func(__e *ControlFlow) {
C6713 := __e.Get(1)
_ = C6713
tmp21090 := Call(__e, PrimFunc(symshen_4newpv), B6710)


let__18969 := tmp21090
_ = let__18969

tmp21091 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21092 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21091)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21091
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21093 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18969, tmp21092)
}
__typedArg0 := let__18969
__typedArg1 := tmp21092
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21094 := Call(__e, PrimFunc(symis_b), V6714, tmp21093, B6710, L6711, Key6712, C6713)


__e.TailApply(PrimFunc(symshen_4gc), B6710, tmp21094)
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

tmp21095 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21096 := Call(__e, PrimFunc(symshen_4assoc_1_6), symstr, tmp21089, tmp21095)


tmp21097 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21096)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21096
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21097

tmp21098 := MakeNative(func(__e *ControlFlow) {
V6719 := __e.Get(1)
_ = V6719
__e.Return(MakeNative(func(__e *ControlFlow) {
B6715 := __e.Get(1)
_ = B6715
__e.Return(MakeNative(func(__e *ControlFlow) {
L6716 := __e.Get(1)
_ = L6716
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6717 := __e.Get(1)
_ = Key6717
__e.Return(MakeNative(func(__e *ControlFlow) {
C6718 := __e.Get(1)
_ = C6718
tmp21099 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21100 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21099)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21099
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21101 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp21100)
}
__typedArg0 := symstring
__typedArg1 := tmp21100
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6719, tmp21101, B6715, L6716, Key6717, C6718)
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

tmp21103 := Call(__e, PrimFunc(symshen_4assoc_1_6), symstring_1_6n, tmp21098, tmp21102)


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
V6724 := __e.Get(1)
_ = V6724
__e.Return(MakeNative(func(__e *ControlFlow) {
B6720 := __e.Get(1)
_ = B6720
__e.Return(MakeNative(func(__e *ControlFlow) {
L6721 := __e.Get(1)
_ = L6721
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6722 := __e.Get(1)
_ = Key6722
__e.Return(MakeNative(func(__e *ControlFlow) {
C6723 := __e.Get(1)
_ = C6723
tmp21106 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21107 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21106)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21106
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21108 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp21107)
}
__typedArg0 := symstring
__typedArg1 := tmp21107
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6724, tmp21108, B6720, L6721, Key6722, C6723)
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

tmp21109 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21110 := Call(__e, PrimFunc(symshen_4assoc_1_6), symstring_1_6symbol, tmp21105, tmp21109)


tmp21111 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21110)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21110
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21111

tmp21112 := MakeNative(func(__e *ControlFlow) {
V6729 := __e.Get(1)
_ = V6729
__e.Return(MakeNative(func(__e *ControlFlow) {
B6725 := __e.Get(1)
_ = B6725
__e.Return(MakeNative(func(__e *ControlFlow) {
L6726 := __e.Get(1)
_ = L6726
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6727 := __e.Get(1)
_ = Key6727
__e.Return(MakeNative(func(__e *ControlFlow) {
C6728 := __e.Get(1)
_ = C6728
tmp21113 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21114 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp21113)
}
__typedArg0 := symlist
__typedArg1 := tmp21113
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21115 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
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
return PrimCons(tmp21114, tmp21116)
}
__typedArg0 := tmp21114
__typedArg1 := tmp21116
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6729, tmp21117, B6725, L6726, Key6727, C6728)
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

tmp21118 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21119 := Call(__e, PrimFunc(symshen_4assoc_1_6), symsum, tmp21112, tmp21118)


tmp21120 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21119)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21119
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21120

tmp21121 := MakeNative(func(__e *ControlFlow) {
V6734 := __e.Get(1)
_ = V6734
__e.Return(MakeNative(func(__e *ControlFlow) {
B6730 := __e.Get(1)
_ = B6730
__e.Return(MakeNative(func(__e *ControlFlow) {
L6731 := __e.Get(1)
_ = L6731
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6732 := __e.Get(1)
_ = Key6732
__e.Return(MakeNative(func(__e *ControlFlow) {
C6733 := __e.Get(1)
_ = C6733
tmp21122 := Call(__e, PrimFunc(symshen_4newpv), B6730)


let__18970 := tmp21122
_ = let__18970

tmp21123 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21124 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21123)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21123
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21125 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18970, tmp21124)
}
__typedArg0 := let__18970
__typedArg1 := tmp21124
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21126 := Call(__e, PrimFunc(symis_b), V6734, tmp21125, B6730, L6731, Key6732, C6733)


__e.TailApply(PrimFunc(symshen_4gc), B6730, tmp21126)
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

tmp21127 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21128 := Call(__e, PrimFunc(symshen_4assoc_1_6), symsymbol_2, tmp21121, tmp21127)


tmp21129 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21128)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21128
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21129

tmp21130 := MakeNative(func(__e *ControlFlow) {
V6739 := __e.Get(1)
_ = V6739
__e.Return(MakeNative(func(__e *ControlFlow) {
B6735 := __e.Get(1)
_ = B6735
__e.Return(MakeNative(func(__e *ControlFlow) {
L6736 := __e.Get(1)
_ = L6736
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6737 := __e.Get(1)
_ = Key6737
__e.Return(MakeNative(func(__e *ControlFlow) {
C6738 := __e.Get(1)
_ = C6738
tmp21131 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21132 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21131)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21131
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21133 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp21132)
}
__typedArg0 := symsymbol
__typedArg1 := tmp21132
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6739, tmp21133, B6735, L6736, Key6737, C6738)
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

tmp21134 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21135 := Call(__e, PrimFunc(symshen_4assoc_1_6), symsystemf, tmp21130, tmp21134)


tmp21136 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21135)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21135
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21136

tmp21137 := MakeNative(func(__e *ControlFlow) {
V6744 := __e.Get(1)
_ = V6744
__e.Return(MakeNative(func(__e *ControlFlow) {
B6740 := __e.Get(1)
_ = B6740
__e.Return(MakeNative(func(__e *ControlFlow) {
L6741 := __e.Get(1)
_ = L6741
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6742 := __e.Get(1)
_ = Key6742
__e.Return(MakeNative(func(__e *ControlFlow) {
C6743 := __e.Get(1)
_ = C6743
tmp21138 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21139 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21138)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21138
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6744, tmp21139, B6740, L6741, Key6742, C6743)
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

tmp21140 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21141 := Call(__e, PrimFunc(symshen_4assoc_1_6), symsystem_1S_2, tmp21137, tmp21140)


tmp21142 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21141)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21141
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21142

tmp21143 := MakeNative(func(__e *ControlFlow) {
V6750 := __e.Get(1)
_ = V6750
__e.Return(MakeNative(func(__e *ControlFlow) {
B6746 := __e.Get(1)
_ = B6746
__e.Return(MakeNative(func(__e *ControlFlow) {
L6747 := __e.Get(1)
_ = L6747
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6748 := __e.Get(1)
_ = Key6748
__e.Return(MakeNative(func(__e *ControlFlow) {
C6749 := __e.Get(1)
_ = C6749
tmp21144 := Call(__e, PrimFunc(symshen_4newpv), B6746)


let__18971 := tmp21144
_ = let__18971

tmp21145 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18971, Nil)
}
__typedArg0 := let__18971
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21146 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp21145)
}
__typedArg0 := symlist
__typedArg1 := tmp21145
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21147 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18971, Nil)
}
__typedArg0 := let__18971
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21148 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp21147)
}
__typedArg0 := symlist
__typedArg1 := tmp21147
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21149 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21148, Nil)
}
__typedArg0 := tmp21148
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21150 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21149)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21149
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21151 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21146, tmp21150)
}
__typedArg0 := tmp21146
__typedArg1 := tmp21150
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21152 := Call(__e, PrimFunc(symis_b), V6750, tmp21151, B6746, L6747, Key6748, C6749)


__e.TailApply(PrimFunc(symshen_4gc), B6746, tmp21152)
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

tmp21153 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21154 := Call(__e, PrimFunc(symshen_4assoc_1_6), symtail, tmp21143, tmp21153)


tmp21155 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21154)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21154
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21155

tmp21156 := MakeNative(func(__e *ControlFlow) {
V6755 := __e.Get(1)
_ = V6755
__e.Return(MakeNative(func(__e *ControlFlow) {
B6751 := __e.Get(1)
_ = B6751
__e.Return(MakeNative(func(__e *ControlFlow) {
L6752 := __e.Get(1)
_ = L6752
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6753 := __e.Get(1)
_ = Key6753
__e.Return(MakeNative(func(__e *ControlFlow) {
C6754 := __e.Get(1)
_ = C6754
tmp21157 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21158 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21157)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21157
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21159 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp21158)
}
__typedArg0 := symstring
__typedArg1 := tmp21158
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6755, tmp21159, B6751, L6752, Key6753, C6754)
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

tmp21161 := Call(__e, PrimFunc(symshen_4assoc_1_6), symtlstr, tmp21156, tmp21160)


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
V6761 := __e.Get(1)
_ = V6761
__e.Return(MakeNative(func(__e *ControlFlow) {
B6757 := __e.Get(1)
_ = B6757
__e.Return(MakeNative(func(__e *ControlFlow) {
L6758 := __e.Get(1)
_ = L6758
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6759 := __e.Get(1)
_ = Key6759
__e.Return(MakeNative(func(__e *ControlFlow) {
C6760 := __e.Get(1)
_ = C6760
tmp21164 := Call(__e, PrimFunc(symshen_4newpv), B6757)


let__18972 := tmp21164
_ = let__18972

tmp21165 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18972, Nil)
}
__typedArg0 := let__18972
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21166 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp21165)
}
__typedArg0 := symvector
__typedArg1 := tmp21165
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21167 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18972, Nil)
}
__typedArg0 := let__18972
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21168 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp21167)
}
__typedArg0 := symvector
__typedArg1 := tmp21167
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21169 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21168, Nil)
}
__typedArg0 := tmp21168
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21170 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21169)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21169
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21171 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21166, tmp21170)
}
__typedArg0 := tmp21166
__typedArg1 := tmp21170
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21172 := Call(__e, PrimFunc(symis_b), V6761, tmp21171, B6757, L6758, Key6759, C6760)


__e.TailApply(PrimFunc(symshen_4gc), B6757, tmp21172)
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

tmp21174 := Call(__e, PrimFunc(symshen_4assoc_1_6), symtlv, tmp21163, tmp21173)


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
V6766 := __e.Get(1)
_ = V6766
__e.Return(MakeNative(func(__e *ControlFlow) {
B6762 := __e.Get(1)
_ = B6762
__e.Return(MakeNative(func(__e *ControlFlow) {
L6763 := __e.Get(1)
_ = L6763
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6764 := __e.Get(1)
_ = Key6764
__e.Return(MakeNative(func(__e *ControlFlow) {
C6765 := __e.Get(1)
_ = C6765
tmp21177 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21178 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21177)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21177
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21179 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp21178)
}
__typedArg0 := symsymbol
__typedArg1 := tmp21178
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6766, tmp21179, B6762, L6763, Key6764, C6765)
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

tmp21180 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21181 := Call(__e, PrimFunc(symshen_4assoc_1_6), symtc, tmp21176, tmp21180)


tmp21182 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21181)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21181
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21182

tmp21183 := MakeNative(func(__e *ControlFlow) {
V6771 := __e.Get(1)
_ = V6771
__e.Return(MakeNative(func(__e *ControlFlow) {
B6767 := __e.Get(1)
_ = B6767
__e.Return(MakeNative(func(__e *ControlFlow) {
L6768 := __e.Get(1)
_ = L6768
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6769 := __e.Get(1)
_ = Key6769
__e.Return(MakeNative(func(__e *ControlFlow) {
C6770 := __e.Get(1)
_ = C6770
tmp21184 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21185 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21184)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21184
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6771, tmp21185, B6767, L6768, Key6769, C6770)
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

tmp21187 := Call(__e, PrimFunc(symshen_4assoc_1_6), symtc_2, tmp21183, tmp21186)


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
V6776 := __e.Get(1)
_ = V6776
__e.Return(MakeNative(func(__e *ControlFlow) {
B6772 := __e.Get(1)
_ = B6772
__e.Return(MakeNative(func(__e *ControlFlow) {
L6773 := __e.Get(1)
_ = L6773
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6774 := __e.Get(1)
_ = Key6774
__e.Return(MakeNative(func(__e *ControlFlow) {
C6775 := __e.Get(1)
_ = C6775
tmp21190 := Call(__e, PrimFunc(symshen_4newpv), B6772)


let__18973 := tmp21190
_ = let__18973

tmp21191 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18973, Nil)
}
__typedArg0 := let__18973
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21192 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlazy, tmp21191)
}
__typedArg0 := symlazy
__typedArg1 := tmp21191
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21193 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18973, Nil)
}
__typedArg0 := let__18973
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21194 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21193)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21193
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21195 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21192, tmp21194)
}
__typedArg0 := tmp21192
__typedArg1 := tmp21194
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21196 := Call(__e, PrimFunc(symis_b), V6776, tmp21195, B6772, L6773, Key6774, C6775)


__e.TailApply(PrimFunc(symshen_4gc), B6772, tmp21196)
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

tmp21197 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21198 := Call(__e, PrimFunc(symshen_4assoc_1_6), symthaw, tmp21189, tmp21197)


tmp21199 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21198)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21198
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21199

tmp21200 := MakeNative(func(__e *ControlFlow) {
V6781 := __e.Get(1)
_ = V6781
__e.Return(MakeNative(func(__e *ControlFlow) {
B6777 := __e.Get(1)
_ = B6777
__e.Return(MakeNative(func(__e *ControlFlow) {
L6778 := __e.Get(1)
_ = L6778
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6779 := __e.Get(1)
_ = Key6779
__e.Return(MakeNative(func(__e *ControlFlow) {
C6780 := __e.Get(1)
_ = C6780
tmp21201 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21202 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21201)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21201
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21203 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp21202)
}
__typedArg0 := symsymbol
__typedArg1 := tmp21202
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6781, tmp21203, B6777, L6778, Key6779, C6780)
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

tmp21204 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21205 := Call(__e, PrimFunc(symshen_4assoc_1_6), symtrack, tmp21200, tmp21204)


tmp21206 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21205)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21205
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21206

tmp21207 := MakeNative(func(__e *ControlFlow) {
V6786 := __e.Get(1)
_ = V6786
__e.Return(MakeNative(func(__e *ControlFlow) {
B6782 := __e.Get(1)
_ = B6782
__e.Return(MakeNative(func(__e *ControlFlow) {
L6783 := __e.Get(1)
_ = L6783
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6784 := __e.Get(1)
_ = Key6784
__e.Return(MakeNative(func(__e *ControlFlow) {
C6785 := __e.Get(1)
_ = C6785
tmp21208 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21209 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp21208)
}
__typedArg0 := symlist
__typedArg1 := tmp21208
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21210 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21209, Nil)
}
__typedArg0 := tmp21209
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

__e.TailApply(PrimFunc(symis_b), V6786, tmp21211, B6782, L6783, Key6784, C6785)
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

tmp21212 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21213 := Call(__e, PrimFunc(symshen_4assoc_1_6), symtracked, tmp21207, tmp21212)


tmp21214 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21213)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21213
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21214

tmp21215 := MakeNative(func(__e *ControlFlow) {
V6791 := __e.Get(1)
_ = V6791
__e.Return(MakeNative(func(__e *ControlFlow) {
B6787 := __e.Get(1)
_ = B6787
__e.Return(MakeNative(func(__e *ControlFlow) {
L6788 := __e.Get(1)
_ = L6788
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6789 := __e.Get(1)
_ = Key6789
__e.Return(MakeNative(func(__e *ControlFlow) {
C6790 := __e.Get(1)
_ = C6790
tmp21216 := Call(__e, PrimFunc(symshen_4newpv), B6787)


let__18974 := tmp21216
_ = let__18974

tmp21217 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18974, Nil)
}
__typedArg0 := let__18974
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

tmp21219 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symexception, tmp21218)
}
__typedArg0 := symexception
__typedArg1 := tmp21218
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21220 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18974, Nil)
}
__typedArg0 := let__18974
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21221 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21220)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21220
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21222 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21219, tmp21221)
}
__typedArg0 := tmp21219
__typedArg1 := tmp21221
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21223 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21222, Nil)
}
__typedArg0 := tmp21222
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21224 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21223)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21223
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21225 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18974, tmp21224)
}
__typedArg0 := let__18974
__typedArg1 := tmp21224
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21226 := Call(__e, PrimFunc(symis_b), V6791, tmp21225, B6787, L6788, Key6789, C6790)


__e.TailApply(PrimFunc(symshen_4gc), B6787, tmp21226)
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

tmp21227 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21228 := Call(__e, PrimFunc(symshen_4assoc_1_6), symtrap_1error, tmp21215, tmp21227)


tmp21229 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21228)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21228
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21229

tmp21230 := MakeNative(func(__e *ControlFlow) {
V6796 := __e.Get(1)
_ = V6796
__e.Return(MakeNative(func(__e *ControlFlow) {
B6792 := __e.Get(1)
_ = B6792
__e.Return(MakeNative(func(__e *ControlFlow) {
L6793 := __e.Get(1)
_ = L6793
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6794 := __e.Get(1)
_ = Key6794
__e.Return(MakeNative(func(__e *ControlFlow) {
C6795 := __e.Get(1)
_ = C6795
tmp21231 := Call(__e, PrimFunc(symshen_4newpv), B6792)


let__18975 := tmp21231
_ = let__18975

tmp21232 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21233 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21232)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21232
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21234 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18975, tmp21233)
}
__typedArg0 := let__18975
__typedArg1 := tmp21233
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21235 := Call(__e, PrimFunc(symis_b), V6796, tmp21234, B6792, L6793, Key6794, C6795)


__e.TailApply(PrimFunc(symshen_4gc), B6792, tmp21235)
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

tmp21236 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21237 := Call(__e, PrimFunc(symshen_4assoc_1_6), symtuple_2, tmp21230, tmp21236)


tmp21238 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21237)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21237
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21238

tmp21239 := MakeNative(func(__e *ControlFlow) {
V6801 := __e.Get(1)
_ = V6801
__e.Return(MakeNative(func(__e *ControlFlow) {
B6797 := __e.Get(1)
_ = B6797
__e.Return(MakeNative(func(__e *ControlFlow) {
L6798 := __e.Get(1)
_ = L6798
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6799 := __e.Get(1)
_ = Key6799
__e.Return(MakeNative(func(__e *ControlFlow) {
C6800 := __e.Get(1)
_ = C6800
tmp21240 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21241 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp21240)
}
__typedArg0 := symlist
__typedArg1 := tmp21240
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21242 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21241, Nil)
}
__typedArg0 := tmp21241
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21243 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21242)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21242
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21244 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp21243)
}
__typedArg0 := symstring
__typedArg1 := tmp21243
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6801, tmp21244, B6797, L6798, Key6799, C6800)
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

tmp21246 := Call(__e, PrimFunc(symshen_4assoc_1_6), symunabsolute, tmp21239, tmp21245)


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
V6806 := __e.Get(1)
_ = V6806
__e.Return(MakeNative(func(__e *ControlFlow) {
B6802 := __e.Get(1)
_ = B6802
__e.Return(MakeNative(func(__e *ControlFlow) {
L6803 := __e.Get(1)
_ = L6803
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6804 := __e.Get(1)
_ = Key6804
__e.Return(MakeNative(func(__e *ControlFlow) {
C6805 := __e.Get(1)
_ = C6805
tmp21249 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21250 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21249)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21249
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21251 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp21250)
}
__typedArg0 := symsymbol
__typedArg1 := tmp21250
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6806, tmp21251, B6802, L6803, Key6804, C6805)
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

tmp21252 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21253 := Call(__e, PrimFunc(symshen_4assoc_1_6), symundefmacro, tmp21248, tmp21252)


tmp21254 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21253)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21253
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21254

tmp21255 := MakeNative(func(__e *ControlFlow) {
V6813 := __e.Get(1)
_ = V6813
__e.Return(MakeNative(func(__e *ControlFlow) {
B6809 := __e.Get(1)
_ = B6809
__e.Return(MakeNative(func(__e *ControlFlow) {
L6810 := __e.Get(1)
_ = L6810
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6811 := __e.Get(1)
_ = Key6811
__e.Return(MakeNative(func(__e *ControlFlow) {
C6812 := __e.Get(1)
_ = C6812
tmp21256 := Call(__e, PrimFunc(symshen_4newpv), B6809)


let__18976 := tmp21256
_ = let__18976

tmp21257 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18976, Nil)
}
__typedArg0 := let__18976
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21258 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp21257)
}
__typedArg0 := symlist
__typedArg1 := tmp21257
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21259 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18976, Nil)
}
__typedArg0 := let__18976
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21260 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp21259)
}
__typedArg0 := symlist
__typedArg1 := tmp21259
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21261 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18976, Nil)
}
__typedArg0 := let__18976
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21262 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp21261)
}
__typedArg0 := symlist
__typedArg1 := tmp21261
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21263 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21262, Nil)
}
__typedArg0 := tmp21262
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21264 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21263)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21263
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21265 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21260, tmp21264)
}
__typedArg0 := tmp21260
__typedArg1 := tmp21264
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21266 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21265, Nil)
}
__typedArg0 := tmp21265
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21267 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21266)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21266
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21268 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21258, tmp21267)
}
__typedArg0 := tmp21258
__typedArg1 := tmp21267
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21269 := Call(__e, PrimFunc(symis_b), V6813, tmp21268, B6809, L6810, Key6811, C6812)


__e.TailApply(PrimFunc(symshen_4gc), B6809, tmp21269)
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

tmp21270 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21271 := Call(__e, PrimFunc(symshen_4assoc_1_6), symunion, tmp21255, tmp21270)


tmp21272 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21271)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21271
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21272

tmp21273 := MakeNative(func(__e *ControlFlow) {
V6818 := __e.Get(1)
_ = V6818
__e.Return(MakeNative(func(__e *ControlFlow) {
B6814 := __e.Get(1)
_ = B6814
__e.Return(MakeNative(func(__e *ControlFlow) {
L6815 := __e.Get(1)
_ = L6815
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6816 := __e.Get(1)
_ = Key6816
__e.Return(MakeNative(func(__e *ControlFlow) {
C6817 := __e.Get(1)
_ = C6817
tmp21274 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21275 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21274)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21274
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21276 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp21275)
}
__typedArg0 := symsymbol
__typedArg1 := tmp21275
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6818, tmp21276, B6814, L6815, Key6816, C6817)
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

tmp21277 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21278 := Call(__e, PrimFunc(symshen_4assoc_1_6), symunprofile, tmp21273, tmp21277)


tmp21279 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21278)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21278
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21279

tmp21280 := MakeNative(func(__e *ControlFlow) {
V6823 := __e.Get(1)
_ = V6823
__e.Return(MakeNative(func(__e *ControlFlow) {
B6819 := __e.Get(1)
_ = B6819
__e.Return(MakeNative(func(__e *ControlFlow) {
L6820 := __e.Get(1)
_ = L6820
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6821 := __e.Get(1)
_ = Key6821
__e.Return(MakeNative(func(__e *ControlFlow) {
C6822 := __e.Get(1)
_ = C6822
tmp21281 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21282 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21281)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21281
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21283 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp21282)
}
__typedArg0 := symsymbol
__typedArg1 := tmp21282
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6823, tmp21283, B6819, L6820, Key6821, C6822)
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

tmp21284 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21285 := Call(__e, PrimFunc(symshen_4assoc_1_6), symuntrack, tmp21280, tmp21284)


tmp21286 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21285)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21285
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21286

tmp21287 := MakeNative(func(__e *ControlFlow) {
V6828 := __e.Get(1)
_ = V6828
__e.Return(MakeNative(func(__e *ControlFlow) {
B6824 := __e.Get(1)
_ = B6824
__e.Return(MakeNative(func(__e *ControlFlow) {
L6825 := __e.Get(1)
_ = L6825
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6826 := __e.Get(1)
_ = Key6826
__e.Return(MakeNative(func(__e *ControlFlow) {
C6827 := __e.Get(1)
_ = C6827
tmp21288 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21289 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp21288)
}
__typedArg0 := symlist
__typedArg1 := tmp21288
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21290 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21289, Nil)
}
__typedArg0 := tmp21289
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21291 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21290)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21290
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6828, tmp21291, B6824, L6825, Key6826, C6827)
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

tmp21292 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21293 := Call(__e, PrimFunc(symshen_4assoc_1_6), symuserdefs, tmp21287, tmp21292)


tmp21294 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21293)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21293
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21294

tmp21295 := MakeNative(func(__e *ControlFlow) {
V6833 := __e.Get(1)
_ = V6833
__e.Return(MakeNative(func(__e *ControlFlow) {
B6829 := __e.Get(1)
_ = B6829
__e.Return(MakeNative(func(__e *ControlFlow) {
L6830 := __e.Get(1)
_ = L6830
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6831 := __e.Get(1)
_ = Key6831
__e.Return(MakeNative(func(__e *ControlFlow) {
C6832 := __e.Get(1)
_ = C6832
tmp21296 := Call(__e, PrimFunc(symshen_4newpv), B6829)


let__18977 := tmp21296
_ = let__18977

tmp21297 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21298 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21297)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21297
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21299 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18977, tmp21298)
}
__typedArg0 := let__18977
__typedArg1 := tmp21298
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21300 := Call(__e, PrimFunc(symis_b), V6833, tmp21299, B6829, L6830, Key6831, C6832)


__e.TailApply(PrimFunc(symshen_4gc), B6829, tmp21300)
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

tmp21301 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21302 := Call(__e, PrimFunc(symshen_4assoc_1_6), symvariable_2, tmp21295, tmp21301)


tmp21303 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21302)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21302
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21303

tmp21304 := MakeNative(func(__e *ControlFlow) {
V6838 := __e.Get(1)
_ = V6838
__e.Return(MakeNative(func(__e *ControlFlow) {
B6834 := __e.Get(1)
_ = B6834
__e.Return(MakeNative(func(__e *ControlFlow) {
L6835 := __e.Get(1)
_ = L6835
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6836 := __e.Get(1)
_ = Key6836
__e.Return(MakeNative(func(__e *ControlFlow) {
C6837 := __e.Get(1)
_ = C6837
tmp21305 := Call(__e, PrimFunc(symshen_4newpv), B6834)


let__18978 := tmp21305
_ = let__18978

tmp21306 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21307 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21306)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21306
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21308 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18978, tmp21307)
}
__typedArg0 := let__18978
__typedArg1 := tmp21307
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21309 := Call(__e, PrimFunc(symis_b), V6838, tmp21308, B6834, L6835, Key6836, C6837)


__e.TailApply(PrimFunc(symshen_4gc), B6834, tmp21309)
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

tmp21311 := Call(__e, PrimFunc(symshen_4assoc_1_6), symvector_2, tmp21304, tmp21310)


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
V6843 := __e.Get(1)
_ = V6843
__e.Return(MakeNative(func(__e *ControlFlow) {
B6839 := __e.Get(1)
_ = B6839
__e.Return(MakeNative(func(__e *ControlFlow) {
L6840 := __e.Get(1)
_ = L6840
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6841 := __e.Get(1)
_ = Key6841
__e.Return(MakeNative(func(__e *ControlFlow) {
C6842 := __e.Get(1)
_ = C6842
tmp21314 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
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

__e.TailApply(PrimFunc(symis_b), V6843, tmp21315, B6839, L6840, Key6841, C6842)
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

tmp21316 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21317 := Call(__e, PrimFunc(symshen_4assoc_1_6), symversion, tmp21313, tmp21316)


tmp21318 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21317)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21317
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21318

tmp21319 := MakeNative(func(__e *ControlFlow) {
V6848 := __e.Get(1)
_ = V6848
__e.Return(MakeNative(func(__e *ControlFlow) {
B6844 := __e.Get(1)
_ = B6844
__e.Return(MakeNative(func(__e *ControlFlow) {
L6845 := __e.Get(1)
_ = L6845
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6846 := __e.Get(1)
_ = Key6846
__e.Return(MakeNative(func(__e *ControlFlow) {
C6847 := __e.Get(1)
_ = C6847
tmp21320 := Call(__e, PrimFunc(symshen_4newpv), B6844)


let__18979 := tmp21320
_ = let__18979

tmp21321 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18979, Nil)
}
__typedArg0 := let__18979
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21322 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21321)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21321
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21323 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18979, tmp21322)
}
__typedArg0 := let__18979
__typedArg1 := tmp21322
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21324 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21323, Nil)
}
__typedArg0 := tmp21323
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21325 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21324)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21324
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21326 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp21325)
}
__typedArg0 := symstring
__typedArg1 := tmp21325
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21327 := Call(__e, PrimFunc(symis_b), V6848, tmp21326, B6844, L6845, Key6846, C6847)


__e.TailApply(PrimFunc(symshen_4gc), B6844, tmp21327)
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

tmp21328 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21329 := Call(__e, PrimFunc(symshen_4assoc_1_6), symwrite_1to_1file, tmp21319, tmp21328)


tmp21330 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21329)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21329
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21330

tmp21331 := MakeNative(func(__e *ControlFlow) {
V6853 := __e.Get(1)
_ = V6853
__e.Return(MakeNative(func(__e *ControlFlow) {
B6849 := __e.Get(1)
_ = B6849
__e.Return(MakeNative(func(__e *ControlFlow) {
L6850 := __e.Get(1)
_ = L6850
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6851 := __e.Get(1)
_ = Key6851
__e.Return(MakeNative(func(__e *ControlFlow) {
C6852 := __e.Get(1)
_ = C6852
tmp21332 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symout, Nil)
}
__typedArg0 := symout
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21333 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp21332)
}
__typedArg0 := symstream
__typedArg1 := tmp21332
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21334 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21335 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21334)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21334
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21336 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21333, tmp21335)
}
__typedArg0 := tmp21333
__typedArg1 := tmp21335
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21337 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21336, Nil)
}
__typedArg0 := tmp21336
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21338 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21337)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21337
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21339 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp21338)
}
__typedArg0 := symnumber
__typedArg1 := tmp21338
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6853, tmp21339, B6849, L6850, Key6851, C6852)
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

tmp21340 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21341 := Call(__e, PrimFunc(symshen_4assoc_1_6), symwrite_1byte, tmp21331, tmp21340)


tmp21342 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21341)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21341
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21342

tmp21343 := MakeNative(func(__e *ControlFlow) {
V6858 := __e.Get(1)
_ = V6858
__e.Return(MakeNative(func(__e *ControlFlow) {
B6854 := __e.Get(1)
_ = B6854
__e.Return(MakeNative(func(__e *ControlFlow) {
L6855 := __e.Get(1)
_ = L6855
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6856 := __e.Get(1)
_ = Key6856
__e.Return(MakeNative(func(__e *ControlFlow) {
C6857 := __e.Get(1)
_ = C6857
tmp21344 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symout, Nil)
}
__typedArg0 := symout
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21345 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp21344)
}
__typedArg0 := symstream
__typedArg1 := tmp21344
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21346 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21347 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21346)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21346
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21348 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21345, tmp21347)
}
__typedArg0 := tmp21345
__typedArg1 := tmp21347
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21349 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21348, Nil)
}
__typedArg0 := tmp21348
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21350 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21349)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21349
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21351 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp21350)
}
__typedArg0 := symstring
__typedArg1 := tmp21350
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6858, tmp21351, B6854, L6855, Key6856, C6857)
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

tmp21352 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21353 := Call(__e, PrimFunc(symshen_4assoc_1_6), symwrite_1string, tmp21343, tmp21352)


tmp21354 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21353)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21353
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21354

tmp21355 := MakeNative(func(__e *ControlFlow) {
V6863 := __e.Get(1)
_ = V6863
__e.Return(MakeNative(func(__e *ControlFlow) {
B6859 := __e.Get(1)
_ = B6859
__e.Return(MakeNative(func(__e *ControlFlow) {
L6860 := __e.Get(1)
_ = L6860
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6861 := __e.Get(1)
_ = Key6861
__e.Return(MakeNative(func(__e *ControlFlow) {
C6862 := __e.Get(1)
_ = C6862
tmp21356 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21357 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21356)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21356
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21358 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp21357)
}
__typedArg0 := symstring
__typedArg1 := tmp21357
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6863, tmp21358, B6859, L6860, Key6861, C6862)
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

tmp21359 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21360 := Call(__e, PrimFunc(symshen_4assoc_1_6), symy_1or_1n_2, tmp21355, tmp21359)


tmp21361 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21360)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21360
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21361

tmp21362 := MakeNative(func(__e *ControlFlow) {
V6868 := __e.Get(1)
_ = V6868
__e.Return(MakeNative(func(__e *ControlFlow) {
B6864 := __e.Get(1)
_ = B6864
__e.Return(MakeNative(func(__e *ControlFlow) {
L6865 := __e.Get(1)
_ = L6865
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6866 := __e.Get(1)
_ = Key6866
__e.Return(MakeNative(func(__e *ControlFlow) {
C6867 := __e.Get(1)
_ = C6867
tmp21363 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21364 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21363)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21363
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21365 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp21364)
}
__typedArg0 := symnumber
__typedArg1 := tmp21364
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21366 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21365, Nil)
}
__typedArg0 := tmp21365
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21367 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21366)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21366
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21368 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp21367)
}
__typedArg0 := symnumber
__typedArg1 := tmp21367
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6868, tmp21368, B6864, L6865, Key6866, C6867)
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

tmp21369 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21370 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_6, tmp21362, tmp21369)


tmp21371 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21370)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21370
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21371

tmp21372 := MakeNative(func(__e *ControlFlow) {
V6873 := __e.Get(1)
_ = V6873
__e.Return(MakeNative(func(__e *ControlFlow) {
B6869 := __e.Get(1)
_ = B6869
__e.Return(MakeNative(func(__e *ControlFlow) {
L6870 := __e.Get(1)
_ = L6870
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6871 := __e.Get(1)
_ = Key6871
__e.Return(MakeNative(func(__e *ControlFlow) {
C6872 := __e.Get(1)
_ = C6872
tmp21373 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21374 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21373)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21373
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21375 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp21374)
}
__typedArg0 := symnumber
__typedArg1 := tmp21374
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21376 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21375, Nil)
}
__typedArg0 := tmp21375
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21377 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21376)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21376
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21378 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp21377)
}
__typedArg0 := symnumber
__typedArg1 := tmp21377
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6873, tmp21378, B6869, L6870, Key6871, C6872)
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

tmp21379 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21380 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_5, tmp21372, tmp21379)


tmp21381 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21380)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21380
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21381

tmp21382 := MakeNative(func(__e *ControlFlow) {
V6878 := __e.Get(1)
_ = V6878
__e.Return(MakeNative(func(__e *ControlFlow) {
B6874 := __e.Get(1)
_ = B6874
__e.Return(MakeNative(func(__e *ControlFlow) {
L6875 := __e.Get(1)
_ = L6875
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6876 := __e.Get(1)
_ = Key6876
__e.Return(MakeNative(func(__e *ControlFlow) {
C6877 := __e.Get(1)
_ = C6877
tmp21383 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21384 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21383)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21383
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21385 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp21384)
}
__typedArg0 := symnumber
__typedArg1 := tmp21384
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21386 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21385, Nil)
}
__typedArg0 := tmp21385
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21387 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21386)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21386
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21388 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp21387)
}
__typedArg0 := symnumber
__typedArg1 := tmp21387
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6878, tmp21388, B6874, L6875, Key6876, C6877)
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

tmp21389 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21390 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_6_a, tmp21382, tmp21389)


tmp21391 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21390)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21390
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21391

tmp21392 := MakeNative(func(__e *ControlFlow) {
V6883 := __e.Get(1)
_ = V6883
__e.Return(MakeNative(func(__e *ControlFlow) {
B6879 := __e.Get(1)
_ = B6879
__e.Return(MakeNative(func(__e *ControlFlow) {
L6880 := __e.Get(1)
_ = L6880
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6881 := __e.Get(1)
_ = Key6881
__e.Return(MakeNative(func(__e *ControlFlow) {
C6882 := __e.Get(1)
_ = C6882
tmp21393 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21394 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21393)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21393
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21395 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp21394)
}
__typedArg0 := symnumber
__typedArg1 := tmp21394
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21396 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21395, Nil)
}
__typedArg0 := tmp21395
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

__e.TailApply(PrimFunc(symis_b), V6883, tmp21398, B6879, L6880, Key6881, C6882)
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

tmp21399 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21400 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_5_a, tmp21392, tmp21399)


tmp21401 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21400)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21400
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21401

tmp21402 := MakeNative(func(__e *ControlFlow) {
V6888 := __e.Get(1)
_ = V6888
__e.Return(MakeNative(func(__e *ControlFlow) {
B6884 := __e.Get(1)
_ = B6884
__e.Return(MakeNative(func(__e *ControlFlow) {
L6885 := __e.Get(1)
_ = L6885
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6886 := __e.Get(1)
_ = Key6886
__e.Return(MakeNative(func(__e *ControlFlow) {
C6887 := __e.Get(1)
_ = C6887
tmp21403 := Call(__e, PrimFunc(symshen_4newpv), B6884)


let__18980 := tmp21403
_ = let__18980

tmp21404 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21405 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21404)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21404
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21406 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18980, tmp21405)
}
__typedArg0 := let__18980
__typedArg1 := tmp21405
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21407 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21406, Nil)
}
__typedArg0 := tmp21406
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21408 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21407)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21407
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21409 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18980, tmp21408)
}
__typedArg0 := let__18980
__typedArg1 := tmp21408
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21410 := Call(__e, PrimFunc(symis_b), V6888, tmp21409, B6884, L6885, Key6886, C6887)


__e.TailApply(PrimFunc(symshen_4gc), B6884, tmp21410)
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

tmp21411 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21412 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_a, tmp21402, tmp21411)


tmp21413 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21412)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21412
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21413

tmp21414 := MakeNative(func(__e *ControlFlow) {
V6893 := __e.Get(1)
_ = V6893
__e.Return(MakeNative(func(__e *ControlFlow) {
B6889 := __e.Get(1)
_ = B6889
__e.Return(MakeNative(func(__e *ControlFlow) {
L6890 := __e.Get(1)
_ = L6890
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6891 := __e.Get(1)
_ = Key6891
__e.Return(MakeNative(func(__e *ControlFlow) {
C6892 := __e.Get(1)
_ = C6892
tmp21415 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21416 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21415)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21415
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21417 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp21416)
}
__typedArg0 := symnumber
__typedArg1 := tmp21416
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21418 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21417, Nil)
}
__typedArg0 := tmp21417
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21419 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21418)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21418
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21420 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp21419)
}
__typedArg0 := symnumber
__typedArg1 := tmp21419
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6893, tmp21420, B6889, L6890, Key6891, C6892)
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

tmp21421 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21422 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_7, tmp21414, tmp21421)


tmp21423 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21422)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21422
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21423

tmp21424 := MakeNative(func(__e *ControlFlow) {
V6898 := __e.Get(1)
_ = V6898
__e.Return(MakeNative(func(__e *ControlFlow) {
B6894 := __e.Get(1)
_ = B6894
__e.Return(MakeNative(func(__e *ControlFlow) {
L6895 := __e.Get(1)
_ = L6895
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6896 := __e.Get(1)
_ = Key6896
__e.Return(MakeNative(func(__e *ControlFlow) {
C6897 := __e.Get(1)
_ = C6897
tmp21425 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21426 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21425)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21425
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21427 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp21426)
}
__typedArg0 := symnumber
__typedArg1 := tmp21426
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21428 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21427, Nil)
}
__typedArg0 := tmp21427
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21429 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21428)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21428
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21430 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp21429)
}
__typedArg0 := symnumber
__typedArg1 := tmp21429
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6898, tmp21430, B6894, L6895, Key6896, C6897)
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

tmp21431 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21432 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_c, tmp21424, tmp21431)


tmp21433 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21432)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21432
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21433

tmp21434 := MakeNative(func(__e *ControlFlow) {
V6903 := __e.Get(1)
_ = V6903
__e.Return(MakeNative(func(__e *ControlFlow) {
B6899 := __e.Get(1)
_ = B6899
__e.Return(MakeNative(func(__e *ControlFlow) {
L6900 := __e.Get(1)
_ = L6900
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6901 := __e.Get(1)
_ = Key6901
__e.Return(MakeNative(func(__e *ControlFlow) {
C6902 := __e.Get(1)
_ = C6902
tmp21435 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21436 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21435)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21435
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21437 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp21436)
}
__typedArg0 := symnumber
__typedArg1 := tmp21436
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21438 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21437, Nil)
}
__typedArg0 := tmp21437
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21439 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21438)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21438
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21440 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp21439)
}
__typedArg0 := symnumber
__typedArg1 := tmp21439
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6903, tmp21440, B6899, L6900, Key6901, C6902)
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

tmp21441 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21442 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_1, tmp21434, tmp21441)


tmp21443 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21442)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21442
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21443

tmp21444 := MakeNative(func(__e *ControlFlow) {
V6908 := __e.Get(1)
_ = V6908
__e.Return(MakeNative(func(__e *ControlFlow) {
B6904 := __e.Get(1)
_ = B6904
__e.Return(MakeNative(func(__e *ControlFlow) {
L6905 := __e.Get(1)
_ = L6905
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6906 := __e.Get(1)
_ = Key6906
__e.Return(MakeNative(func(__e *ControlFlow) {
C6907 := __e.Get(1)
_ = C6907
tmp21445 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21446 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21445)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21445
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21447 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp21446)
}
__typedArg0 := symnumber
__typedArg1 := tmp21446
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21448 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21447, Nil)
}
__typedArg0 := tmp21447
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

__e.TailApply(PrimFunc(symis_b), V6908, tmp21450, B6904, L6905, Key6906, C6907)
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

tmp21451 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21452 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_d, tmp21444, tmp21451)


tmp21453 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21452)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21452
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21453

tmp21454 := MakeNative(func(__e *ControlFlow) {
V6913 := __e.Get(1)
_ = V6913
__e.Return(MakeNative(func(__e *ControlFlow) {
B6909 := __e.Get(1)
_ = B6909
__e.Return(MakeNative(func(__e *ControlFlow) {
L6910 := __e.Get(1)
_ = L6910
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6911 := __e.Get(1)
_ = Key6911
__e.Return(MakeNative(func(__e *ControlFlow) {
C6912 := __e.Get(1)
_ = C6912
tmp21455 := Call(__e, PrimFunc(symshen_4newpv), B6909)


let__18981 := tmp21455
_ = let__18981

tmp21456 := Call(__e, PrimFunc(symshen_4newpv), B6909)


let__18982 := tmp21456
_ = let__18982

tmp21457 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21458 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21457)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21457
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21459 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18982, tmp21458)
}
__typedArg0 := let__18982
__typedArg1 := tmp21458
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21460 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21459, Nil)
}
__typedArg0 := tmp21459
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21461 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21460)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21460
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21462 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18981, tmp21461)
}
__typedArg0 := let__18981
__typedArg1 := tmp21461
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21463 := Call(__e, PrimFunc(symis_b), V6913, tmp21462, B6909, L6910, Key6911, C6912)


tmp21464 := Call(__e, PrimFunc(symshen_4gc), B6909, tmp21463)


__e.TailApply(PrimFunc(symshen_4gc), B6909, tmp21464)
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

tmp21465 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21466 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_a_a, tmp21454, tmp21465)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21466)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21466
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return


}, 0)

tmp21467 := Call(__e, ns2_1set, symshen_4initialise_1signedfuncs, tmp19748)


_ = tmp21467

tmp21468 := MakeNative(func(__e *ControlFlow) {
tmp21469 := MakeNative(func(__e *ControlFlow) {
Y954 := __e.Get(1)
_ = Y954
__e.TailApply(PrimFunc(symshen_4tuple), Y954)
return
}, 1)

tmp21470 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4tuple, tmp21469)
}
__typedArg0 := symshen_4tuple
__typedArg1 := tmp21469
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21471 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21470)


_ = tmp21471

tmp21472 := MakeNative(func(__e *ControlFlow) {
Y955 := __e.Get(1)
_ = Y955
__e.TailApply(PrimFunc(symshen_4pvar), Y955)
return
}, 1)

tmp21473 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4pvar, tmp21472)
}
__typedArg0 := symshen_4pvar
__typedArg1 := tmp21472
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21474 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21473)


_ = tmp21474

tmp21475 := MakeNative(func(__e *ControlFlow) {
Y956 := __e.Get(1)
_ = Y956
__e.TailApply(PrimFunc(symshen_4dictionary), Y956)
return
}, 1)

tmp21476 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4dictionary, tmp21475)
}
__typedArg0 := symshen_4dictionary
__typedArg1 := tmp21475
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21477 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21476)


_ = tmp21477

tmp21478 := MakeNative(func(__e *ControlFlow) {
Y957 := __e.Get(1)
_ = Y957
__e.TailApply(PrimFunc(symshen_4print_1prolog_1vector), Y957)
return
}, 1)

tmp21479 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4print_1prolog_1vector, tmp21478)
}
__typedArg0 := symshen_4print_1prolog_1vector
__typedArg1 := tmp21478
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21480 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21479)


_ = tmp21480

tmp21481 := MakeNative(func(__e *ControlFlow) {
Y958 := __e.Get(1)
_ = Y958
__e.TailApply(PrimFunc(symshen_4print_1freshterm), Y958)
return
}, 1)

tmp21482 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4print_1freshterm, tmp21481)
}
__typedArg0 := symshen_4print_1freshterm
__typedArg1 := tmp21481
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21483 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21482)


_ = tmp21483

tmp21484 := MakeNative(func(__e *ControlFlow) {
Y959 := __e.Get(1)
_ = Y959
__e.TailApply(PrimFunc(symshen_4printF), Y959)
return
}, 1)

tmp21485 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4printF, tmp21484)
}
__typedArg0 := symshen_4printF
__typedArg1 := tmp21484
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21486 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21485)


_ = tmp21486

tmp21487 := MakeNative(func(__e *ControlFlow) {
Y960 := __e.Get(1)
_ = Y960
__e.TailApply(PrimFunc(symabsolute), Y960)
return
}, 1)

tmp21488 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symabsolute, tmp21487)
}
__typedArg0 := symabsolute
__typedArg1 := tmp21487
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21489 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21488)


_ = tmp21489

tmp21490 := MakeNative(func(__e *ControlFlow) {
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

tmp21491 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symabsvector_2, tmp21490)
}
__typedArg0 := symabsvector_2
__typedArg1 := tmp21490
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21492 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21491)


_ = tmp21492

tmp21493 := MakeNative(func(__e *ControlFlow) {
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

tmp21494 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symabsvector, tmp21493)
}
__typedArg0 := symabsvector
__typedArg1 := tmp21493
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21495 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21494)


_ = tmp21495

tmp21496 := MakeNative(func(__e *ControlFlow) {
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

tmp21497 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symaddress_1_6, tmp21496)
}
__typedArg0 := symaddress_1_6
__typedArg1 := tmp21496
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21498 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21497)


_ = tmp21498

tmp21499 := MakeNative(func(__e *ControlFlow) {
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

tmp21500 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symadjoin, tmp21499)
}
__typedArg0 := symadjoin
__typedArg1 := tmp21499
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21501 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21500)


_ = tmp21501

tmp21502 := MakeNative(func(__e *ControlFlow) {
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

tmp21505 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symand, tmp21502)
}
__typedArg0 := symand
__typedArg1 := tmp21502
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21506 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21505)


_ = tmp21506

tmp21507 := MakeNative(func(__e *ControlFlow) {
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

tmp21508 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symappend, tmp21507)
}
__typedArg0 := symappend
__typedArg1 := tmp21507
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21509 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21508)


_ = tmp21509

tmp21510 := MakeNative(func(__e *ControlFlow) {
Y972 := __e.Get(1)
_ = Y972
__e.TailApply(PrimFunc(symarity), Y972)
return
}, 1)

tmp21511 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symarity, tmp21510)
}
__typedArg0 := symarity
__typedArg1 := tmp21510
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21512 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21511)


_ = tmp21512

tmp21513 := MakeNative(func(__e *ControlFlow) {
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

tmp21514 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symassoc, tmp21513)
}
__typedArg0 := symassoc
__typedArg1 := tmp21513
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21515 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21514)


_ = tmp21515

tmp21516 := MakeNative(func(__e *ControlFlow) {
Y975 := __e.Get(1)
_ = Y975
__e.TailApply(PrimFunc(symatom_2), Y975)
return
}, 1)

tmp21517 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symatom_2, tmp21516)
}
__typedArg0 := symatom_2
__typedArg1 := tmp21516
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21518 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21517)


_ = tmp21518

tmp21519 := MakeNative(func(__e *ControlFlow) {
Y976 := __e.Get(1)
_ = Y976
__e.TailApply(PrimFunc(symboolean_2), Y976)
return
}, 1)

tmp21520 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean_2, tmp21519)
}
__typedArg0 := symboolean_2
__typedArg1 := tmp21519
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21521 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21520)


_ = tmp21521

tmp21522 := MakeNative(func(__e *ControlFlow) {
Y977 := __e.Get(1)
_ = Y977
__e.TailApply(PrimFunc(symbootstrap), Y977)
return
}, 1)

tmp21523 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symbootstrap, tmp21522)
}
__typedArg0 := symbootstrap
__typedArg1 := tmp21522
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21524 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21523)


_ = tmp21524

tmp21525 := MakeNative(func(__e *ControlFlow) {
Y978 := __e.Get(1)
_ = Y978
__e.TailApply(PrimFunc(symbound_2), Y978)
return
}, 1)

tmp21526 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symbound_2, tmp21525)
}
__typedArg0 := symbound_2
__typedArg1 := tmp21525
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21527 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21526)


_ = tmp21527

tmp21528 := MakeNative(func(__e *ControlFlow) {
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

tmp21529 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symbind, tmp21528)
}
__typedArg0 := symbind
__typedArg1 := tmp21528
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21530 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21529)


_ = tmp21530

tmp21531 := MakeNative(func(__e *ControlFlow) {
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

tmp21532 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcall, tmp21531)
}
__typedArg0 := symcall
__typedArg1 := tmp21531
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21533 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21532)


_ = tmp21533

tmp21534 := MakeNative(func(__e *ControlFlow) {
Y990 := __e.Get(1)
_ = Y990
__e.TailApply(PrimFunc(symcd), Y990)
return
}, 1)

tmp21535 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcd, tmp21534)
}
__typedArg0 := symcd
__typedArg1 := tmp21534
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21536 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21535)


_ = tmp21536

tmp21537 := MakeNative(func(__e *ControlFlow) {
Y991 := __e.Get(1)
_ = Y991
__e.TailApply(PrimFunc(symchar_1stream_2), Y991)
return
}, 1)

tmp21538 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symchar_1stream_2, tmp21537)
}
__typedArg0 := symchar_1stream_2
__typedArg1 := tmp21537
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21539 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21538)


_ = tmp21539

tmp21540 := MakeNative(func(__e *ControlFlow) {
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

tmp21541 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcompile, tmp21540)
}
__typedArg0 := symcompile
__typedArg1 := tmp21540
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21542 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21541)


_ = tmp21542

tmp21543 := MakeNative(func(__e *ControlFlow) {
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

tmp21544 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symconcat, tmp21543)
}
__typedArg0 := symconcat
__typedArg1 := tmp21543
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21545 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21544)


_ = tmp21545

tmp21546 := MakeNative(func(__e *ControlFlow) {
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

tmp21547 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp21546)
}
__typedArg0 := symcons
__typedArg1 := tmp21546
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21548 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21547)


_ = tmp21548

tmp21549 := MakeNative(func(__e *ControlFlow) {
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

tmp21550 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons_2, tmp21549)
}
__typedArg0 := symcons_2
__typedArg1 := tmp21549
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21551 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21550)


_ = tmp21551

tmp21552 := MakeNative(func(__e *ControlFlow) {
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

tmp21553 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcn, tmp21552)
}
__typedArg0 := symcn
__typedArg1 := tmp21552
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21554 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21553)


_ = tmp21554

tmp21555 := MakeNative(func(__e *ControlFlow) {
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

tmp21556 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symclose, tmp21555)
}
__typedArg0 := symclose
__typedArg1 := tmp21555
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21557 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21556)


_ = tmp21557

tmp21558 := MakeNative(func(__e *ControlFlow) {
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

tmp21559 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdeclare, tmp21558)
}
__typedArg0 := symdeclare
__typedArg1 := tmp21558
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21560 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21559)


_ = tmp21560

tmp21561 := MakeNative(func(__e *ControlFlow) {
Y1004 := __e.Get(1)
_ = Y1004
__e.TailApply(PrimFunc(symdestroy), Y1004)
return
}, 1)

tmp21562 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdestroy, tmp21561)
}
__typedArg0 := symdestroy
__typedArg1 := tmp21561
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21563 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21562)


_ = tmp21563

tmp21564 := MakeNative(func(__e *ControlFlow) {
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

tmp21565 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdifference, tmp21564)
}
__typedArg0 := symdifference
__typedArg1 := tmp21564
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21566 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21565)


_ = tmp21566

tmp21567 := MakeNative(func(__e *ControlFlow) {
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

tmp21568 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdo, tmp21567)
}
__typedArg0 := symdo
__typedArg1 := tmp21567
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21569 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21568)


_ = tmp21569

tmp21570 := MakeNative(func(__e *ControlFlow) {
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

tmp21571 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symelement_2, tmp21570)
}
__typedArg0 := symelement_2
__typedArg1 := tmp21570
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21572 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21571)


_ = tmp21572

tmp21573 := MakeNative(func(__e *ControlFlow) {
Y1011 := __e.Get(1)
_ = Y1011
__e.TailApply(PrimFunc(symempty_2), Y1011)
return
}, 1)

tmp21574 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symempty_2, tmp21573)
}
__typedArg0 := symempty_2
__typedArg1 := tmp21573
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21575 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21574)


_ = tmp21575

tmp21576 := MakeNative(func(__e *ControlFlow) {
Y1012 := __e.Get(1)
_ = Y1012
__e.TailApply(PrimFunc(symenable_1type_1theory), Y1012)
return
}, 1)

tmp21577 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symenable_1type_1theory, tmp21576)
}
__typedArg0 := symenable_1type_1theory
__typedArg1 := tmp21576
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21578 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21577)


_ = tmp21578

tmp21579 := MakeNative(func(__e *ControlFlow) {
Y1013 := __e.Get(1)
_ = Y1013
__e.TailApply(PrimFunc(symexternal), Y1013)
return
}, 1)

tmp21580 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symexternal, tmp21579)
}
__typedArg0 := symexternal
__typedArg1 := tmp21579
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21581 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21580)


_ = tmp21581

tmp21582 := MakeNative(func(__e *ControlFlow) {
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

tmp21583 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symerror_1to_1string, tmp21582)
}
__typedArg0 := symerror_1to_1string
__typedArg1 := tmp21582
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21584 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21583)


_ = tmp21584

tmp21585 := MakeNative(func(__e *ControlFlow) {
Y1015 := __e.Get(1)
_ = Y1015
__e.TailApply(PrimFunc(symeval), Y1015)
return
}, 1)

tmp21586 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symeval, tmp21585)
}
__typedArg0 := symeval
__typedArg1 := tmp21585
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21587 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21586)


_ = tmp21587

tmp21588 := MakeNative(func(__e *ControlFlow) {
Y1016 := __e.Get(1)
_ = Y1016
__e.TailApply(PrimFunc(symeval_1kl), Y1016)
return
}, 1)

tmp21589 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symeval_1kl, tmp21588)
}
__typedArg0 := symeval_1kl
__typedArg1 := tmp21588
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21590 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21589)


_ = tmp21590

tmp21591 := MakeNative(func(__e *ControlFlow) {
Y1017 := __e.Get(1)
_ = Y1017
__e.TailApply(PrimFunc(symexplode), Y1017)
return
}, 1)

tmp21592 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symexplode, tmp21591)
}
__typedArg0 := symexplode
__typedArg1 := tmp21591
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21593 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21592)


_ = tmp21593

tmp21594 := MakeNative(func(__e *ControlFlow) {
Y1018 := __e.Get(1)
_ = Y1018
__e.TailApply(PrimFunc(symexternal), Y1018)
return
}, 1)

tmp21595 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symexternal, tmp21594)
}
__typedArg0 := symexternal
__typedArg1 := tmp21594
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21596 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21595)


_ = tmp21596

tmp21597 := MakeNative(func(__e *ControlFlow) {
Y1019 := __e.Get(1)
_ = Y1019
__e.TailApply(PrimFunc(symfactorise), Y1019)
return
}, 1)

tmp21598 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfactorise, tmp21597)
}
__typedArg0 := symfactorise
__typedArg1 := tmp21597
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21599 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21598)


_ = tmp21599

tmp21600 := MakeNative(func(__e *ControlFlow) {
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

tmp21601 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfail_1if, tmp21600)
}
__typedArg0 := symfail_1if
__typedArg1 := tmp21600
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21602 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21601)


_ = tmp21602

tmp21603 := MakeNative(func(__e *ControlFlow) {
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

tmp21604 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfix, tmp21603)
}
__typedArg0 := symfix
__typedArg1 := tmp21603
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21605 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21604)


_ = tmp21605

tmp21606 := MakeNative(func(__e *ControlFlow) {
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

tmp21607 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfindall, tmp21606)
}
__typedArg0 := symfindall
__typedArg1 := tmp21606
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21608 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21607)


_ = tmp21608

tmp21609 := MakeNative(func(__e *ControlFlow) {
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

tmp21610 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfork, tmp21609)
}
__typedArg0 := symfork
__typedArg1 := tmp21609
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21611 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21610)


_ = tmp21611

tmp21612 := MakeNative(func(__e *ControlFlow) {
Y1036 := __e.Get(1)
_ = Y1036
__e.Return(MakeNative(func(__e *ControlFlow) {
__e.Return(Y1036)
return
}, 0))
return
}, 1)

tmp21613 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfreeze, tmp21612)
}
__typedArg0 := symfreeze
__typedArg1 := tmp21612
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21614 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21613)


_ = tmp21614

tmp21615 := MakeNative(func(__e *ControlFlow) {
Y1037 := __e.Get(1)
_ = Y1037
__e.TailApply(PrimFunc(symfst), Y1037)
return
}, 1)

tmp21616 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfst, tmp21615)
}
__typedArg0 := symfst
__typedArg1 := tmp21615
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21617 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21616)


_ = tmp21617

tmp21618 := MakeNative(func(__e *ControlFlow) {
Y1038 := __e.Get(1)
_ = Y1038
__e.TailApply(PrimFunc(symfn), Y1038)
return
}, 1)

tmp21619 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfn, tmp21618)
}
__typedArg0 := symfn
__typedArg1 := tmp21618
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21620 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21619)


_ = tmp21620

tmp21621 := MakeNative(func(__e *ControlFlow) {
Y1039 := __e.Get(1)
_ = Y1039
__e.TailApply(PrimFunc(symfunction), Y1039)
return
}, 1)

tmp21622 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfunction, tmp21621)
}
__typedArg0 := symfunction
__typedArg1 := tmp21621
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21623 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21622)


_ = tmp21623

tmp21624 := MakeNative(func(__e *ControlFlow) {
Y1040 := __e.Get(1)
_ = Y1040
__e.TailApply(PrimFunc(symgensym), Y1040)
return
}, 1)

tmp21625 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symgensym, tmp21624)
}
__typedArg0 := symgensym
__typedArg1 := tmp21624
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21626 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21625)


_ = tmp21626

tmp21627 := MakeNative(func(__e *ControlFlow) {
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

tmp21628 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symget, tmp21627)
}
__typedArg0 := symget
__typedArg1 := tmp21627
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21629 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21628)


_ = tmp21629

tmp21630 := MakeNative(func(__e *ControlFlow) {
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

tmp21631 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symget_1time, tmp21630)
}
__typedArg0 := symget_1time
__typedArg1 := tmp21630
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21632 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21631)


_ = tmp21632

tmp21633 := MakeNative(func(__e *ControlFlow) {
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

tmp21634 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symaddress_1_6, tmp21633)
}
__typedArg0 := symaddress_1_6
__typedArg1 := tmp21633
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21635 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21634)


_ = tmp21635

tmp21636 := MakeNative(func(__e *ControlFlow) {
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

tmp21637 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1address, tmp21636)
}
__typedArg0 := sym_5_1address
__typedArg1 := tmp21636
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21638 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21637)


_ = tmp21638

tmp21639 := MakeNative(func(__e *ControlFlow) {
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

tmp21640 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1vector, tmp21639)
}
__typedArg0 := sym_5_1vector
__typedArg1 := tmp21639
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21641 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21640)


_ = tmp21641

tmp21642 := MakeNative(func(__e *ControlFlow) {
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

tmp21643 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_6, tmp21642)
}
__typedArg0 := sym_6
__typedArg1 := tmp21642
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21644 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21643)


_ = tmp21644

tmp21645 := MakeNative(func(__e *ControlFlow) {
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

tmp21646 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_6_a, tmp21645)
}
__typedArg0 := sym_6_a
__typedArg1 := tmp21645
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21647 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21646)


_ = tmp21647

tmp21648 := MakeNative(func(__e *ControlFlow) {
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

tmp21649 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_a, tmp21648)
}
__typedArg0 := sym_a
__typedArg1 := tmp21648
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21650 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21649)


_ = tmp21650

tmp21651 := MakeNative(func(__e *ControlFlow) {
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

tmp21652 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhash, tmp21651)
}
__typedArg0 := symhash
__typedArg1 := tmp21651
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21653 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21652)


_ = tmp21653

tmp21654 := MakeNative(func(__e *ControlFlow) {
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

tmp21655 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhd, tmp21654)
}
__typedArg0 := symhd
__typedArg1 := tmp21654
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21656 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21655)


_ = tmp21656

tmp21657 := MakeNative(func(__e *ControlFlow) {
Y1061 := __e.Get(1)
_ = Y1061
__e.TailApply(PrimFunc(symhdv), Y1061)
return
}, 1)

tmp21658 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhdv, tmp21657)
}
__typedArg0 := symhdv
__typedArg1 := tmp21657
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21659 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21658)


_ = tmp21659

tmp21660 := MakeNative(func(__e *ControlFlow) {
Y1062 := __e.Get(1)
_ = Y1062
__e.TailApply(PrimFunc(symhdstr), Y1062)
return
}, 1)

tmp21661 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhdstr, tmp21660)
}
__typedArg0 := symhdstr
__typedArg1 := tmp21660
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21662 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21661)


_ = tmp21662

tmp21663 := MakeNative(func(__e *ControlFlow) {
Y1063 := __e.Get(1)
_ = Y1063
__e.TailApply(PrimFunc(symhead), Y1063)
return
}, 1)

tmp21664 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhead, tmp21663)
}
__typedArg0 := symhead
__typedArg1 := tmp21663
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21665 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21664)


_ = tmp21665

tmp21666 := MakeNative(func(__e *ControlFlow) {
Y1064 := __e.Get(1)
_ = Y1064
__e.TailApply(PrimFunc(symhush), Y1064)
return
}, 1)

tmp21667 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhush, tmp21666)
}
__typedArg0 := symhush
__typedArg1 := tmp21666
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21668 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21667)


_ = tmp21668

tmp21669 := MakeNative(func(__e *ControlFlow) {
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

tmp21671 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp21669)
}
__typedArg0 := symif
__typedArg1 := tmp21669
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21672 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21671)


_ = tmp21672

tmp21673 := MakeNative(func(__e *ControlFlow) {
Y1068 := __e.Get(1)
_ = Y1068
__e.TailApply(PrimFunc(syminclude), Y1068)
return
}, 1)

tmp21674 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminclude, tmp21673)
}
__typedArg0 := syminclude
__typedArg1 := tmp21673
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21675 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21674)


_ = tmp21675

tmp21676 := MakeNative(func(__e *ControlFlow) {
Y1069 := __e.Get(1)
_ = Y1069
__e.TailApply(PrimFunc(symin_1package), Y1069)
return
}, 1)

tmp21677 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symin_1package, tmp21676)
}
__typedArg0 := symin_1package
__typedArg1 := tmp21676
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21678 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21677)


_ = tmp21678

tmp21679 := MakeNative(func(__e *ControlFlow) {
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

tmp21680 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminteger_2, tmp21679)
}
__typedArg0 := syminteger_2
__typedArg1 := tmp21679
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21681 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21680)


_ = tmp21681

tmp21682 := MakeNative(func(__e *ControlFlow) {
Y1071 := __e.Get(1)
_ = Y1071
__e.TailApply(PrimFunc(syminternal), Y1071)
return
}, 1)

tmp21683 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminternal, tmp21682)
}
__typedArg0 := syminternal
__typedArg1 := tmp21682
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21684 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21683)


_ = tmp21684

tmp21685 := MakeNative(func(__e *ControlFlow) {
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

tmp21686 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symintern, tmp21685)
}
__typedArg0 := symintern
__typedArg1 := tmp21685
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21687 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21686)


_ = tmp21687

tmp21688 := MakeNative(func(__e *ControlFlow) {
Y1073 := __e.Get(1)
_ = Y1073
__e.TailApply(PrimFunc(syminput), Y1073)
return
}, 1)

tmp21689 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminput, tmp21688)
}
__typedArg0 := syminput
__typedArg1 := tmp21688
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21690 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21689)


_ = tmp21690

tmp21691 := MakeNative(func(__e *ControlFlow) {
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

tmp21692 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminput_7, tmp21691)
}
__typedArg0 := syminput_7
__typedArg1 := tmp21691
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21693 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21692)


_ = tmp21693

tmp21694 := MakeNative(func(__e *ControlFlow) {
Y1076 := __e.Get(1)
_ = Y1076
__e.TailApply(PrimFunc(syminclude_1all_1but), Y1076)
return
}, 1)

tmp21695 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminclude_1all_1but, tmp21694)
}
__typedArg0 := syminclude_1all_1but
__typedArg1 := tmp21694
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21696 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21695)


_ = tmp21696

tmp21697 := MakeNative(func(__e *ControlFlow) {
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

tmp21698 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symintersection, tmp21697)
}
__typedArg0 := symintersection
__typedArg1 := tmp21697
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21699 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21698)


_ = tmp21699

tmp21700 := MakeNative(func(__e *ControlFlow) {
Y1079 := __e.Get(1)
_ = Y1079
__e.TailApply(PrimFunc(syminternal), Y1079)
return
}, 1)

tmp21701 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminternal, tmp21700)
}
__typedArg0 := syminternal
__typedArg1 := tmp21700
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21702 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21701)


_ = tmp21702

tmp21703 := MakeNative(func(__e *ControlFlow) {
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

tmp21704 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symis, tmp21703)
}
__typedArg0 := symis
__typedArg1 := tmp21703
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21705 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21704)


_ = tmp21705

tmp21706 := MakeNative(func(__e *ControlFlow) {
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

tmp21707 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symis_b, tmp21706)
}
__typedArg0 := symis_b
__typedArg1 := tmp21706
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21708 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21707)


_ = tmp21708

tmp21709 := MakeNative(func(__e *ControlFlow) {
Y1092 := __e.Get(1)
_ = Y1092
__e.TailApply(PrimFunc(symlength), Y1092)
return
}, 1)

tmp21710 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlength, tmp21709)
}
__typedArg0 := symlength
__typedArg1 := tmp21709
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21711 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21710)


_ = tmp21711

tmp21712 := MakeNative(func(__e *ControlFlow) {
Y1093 := __e.Get(1)
_ = Y1093
__e.TailApply(PrimFunc(symlimit), Y1093)
return
}, 1)

tmp21713 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlimit, tmp21712)
}
__typedArg0 := symlimit
__typedArg1 := tmp21712
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21714 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21713)


_ = tmp21714

tmp21715 := MakeNative(func(__e *ControlFlow) {
Y1094 := __e.Get(1)
_ = Y1094
__e.TailApply(PrimFunc(symlineread), Y1094)
return
}, 1)

tmp21716 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlineread, tmp21715)
}
__typedArg0 := symlineread
__typedArg1 := tmp21715
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21717 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21716)


_ = tmp21717

tmp21718 := MakeNative(func(__e *ControlFlow) {
Y1095 := __e.Get(1)
_ = Y1095
__e.TailApply(PrimFunc(symload), Y1095)
return
}, 1)

tmp21719 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symload, tmp21718)
}
__typedArg0 := symload
__typedArg1 := tmp21718
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21720 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21719)


_ = tmp21720

tmp21721 := MakeNative(func(__e *ControlFlow) {
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

tmp21722 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5, tmp21721)
}
__typedArg0 := sym_5
__typedArg1 := tmp21721
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21723 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21722)


_ = tmp21723

tmp21724 := MakeNative(func(__e *ControlFlow) {
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

tmp21725 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_a, tmp21724)
}
__typedArg0 := sym_5_a
__typedArg1 := tmp21724
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21726 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21725)


_ = tmp21726

tmp21727 := MakeNative(func(__e *ControlFlow) {
Y1100 := __e.Get(1)
_ = Y1100
__e.TailApply(PrimFunc(symvector), Y1100)
return
}, 1)

tmp21728 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp21727)
}
__typedArg0 := symvector
__typedArg1 := tmp21727
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21729 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21728)


_ = tmp21729

tmp21730 := MakeNative(func(__e *ControlFlow) {
Y1101 := __e.Get(1)
_ = Y1101
__e.TailApply(PrimFunc(symmacroexpand), Y1101)
return
}, 1)

tmp21731 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symmacroexpand, tmp21730)
}
__typedArg0 := symmacroexpand
__typedArg1 := tmp21730
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21732 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21731)


_ = tmp21732

tmp21733 := MakeNative(func(__e *ControlFlow) {
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

tmp21734 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symmap, tmp21733)
}
__typedArg0 := symmap
__typedArg1 := tmp21733
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21735 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21734)


_ = tmp21735

tmp21736 := MakeNative(func(__e *ControlFlow) {
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

tmp21737 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symmapcan, tmp21736)
}
__typedArg0 := symmapcan
__typedArg1 := tmp21736
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21738 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21737)


_ = tmp21738

tmp21739 := MakeNative(func(__e *ControlFlow) {
Y1106 := __e.Get(1)
_ = Y1106
__e.TailApply(PrimFunc(symmaxinferences), Y1106)
return
}, 1)

tmp21740 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symmaxinferences, tmp21739)
}
__typedArg0 := symmaxinferences
__typedArg1 := tmp21739
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21741 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21740)


_ = tmp21741

tmp21742 := MakeNative(func(__e *ControlFlow) {
Y1107 := __e.Get(1)
_ = Y1107
__e.TailApply(PrimFunc(symnl), Y1107)
return
}, 1)

tmp21743 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnl, tmp21742)
}
__typedArg0 := symnl
__typedArg1 := tmp21742
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21744 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21743)


_ = tmp21744

tmp21745 := MakeNative(func(__e *ControlFlow) {
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

tmp21746 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnot, tmp21745)
}
__typedArg0 := symnot
__typedArg1 := tmp21745
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21747 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21746)


_ = tmp21747

tmp21748 := MakeNative(func(__e *ControlFlow) {
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

tmp21749 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnth, tmp21748)
}
__typedArg0 := symnth
__typedArg1 := tmp21748
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21750 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21749)


_ = tmp21750

tmp21751 := MakeNative(func(__e *ControlFlow) {
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

tmp21752 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symn_1_6string, tmp21751)
}
__typedArg0 := symn_1_6string
__typedArg1 := tmp21751
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21753 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21752)


_ = tmp21753

tmp21754 := MakeNative(func(__e *ControlFlow) {
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

tmp21755 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber_2, tmp21754)
}
__typedArg0 := symnumber_2
__typedArg1 := tmp21754
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21756 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21755)


_ = tmp21756

tmp21757 := MakeNative(func(__e *ControlFlow) {
Y1113 := __e.Get(1)
_ = Y1113
__e.TailApply(PrimFunc(symoccurs_1check), Y1113)
return
}, 1)

tmp21758 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symoccurs_1check, tmp21757)
}
__typedArg0 := symoccurs_1check
__typedArg1 := tmp21757
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21759 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21758)


_ = tmp21759

tmp21760 := MakeNative(func(__e *ControlFlow) {
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

tmp21761 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symoccurrences, tmp21760)
}
__typedArg0 := symoccurrences
__typedArg1 := tmp21760
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21762 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21761)


_ = tmp21762

tmp21763 := MakeNative(func(__e *ControlFlow) {
Y1116 := __e.Get(1)
_ = Y1116
__e.TailApply(PrimFunc(symoccurs_1check), Y1116)
return
}, 1)

tmp21764 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symoccurs_1check, tmp21763)
}
__typedArg0 := symoccurs_1check
__typedArg1 := tmp21763
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21765 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21764)


_ = tmp21765

tmp21766 := MakeNative(func(__e *ControlFlow) {
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

tmp21767 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symopen, tmp21766)
}
__typedArg0 := symopen
__typedArg1 := tmp21766
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21768 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21767)


_ = tmp21768

tmp21769 := MakeNative(func(__e *ControlFlow) {
Y1119 := __e.Get(1)
_ = Y1119
__e.TailApply(PrimFunc(symoptimise), Y1119)
return
}, 1)

tmp21770 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symoptimise, tmp21769)
}
__typedArg0 := symoptimise
__typedArg1 := tmp21769
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21771 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21770)


_ = tmp21771

tmp21772 := MakeNative(func(__e *ControlFlow) {
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

tmp21775 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symor, tmp21772)
}
__typedArg0 := symor
__typedArg1 := tmp21772
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21776 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21775)


_ = tmp21776

tmp21777 := MakeNative(func(__e *ControlFlow) {
Y1122 := __e.Get(1)
_ = Y1122
__e.TailApply(PrimFunc(sympackage_2), Y1122)
return
}, 1)

tmp21778 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympackage_2, tmp21777)
}
__typedArg0 := sympackage_2
__typedArg1 := tmp21777
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21779 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21778)


_ = tmp21779

tmp21780 := MakeNative(func(__e *ControlFlow) {
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

tmp21781 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympos, tmp21780)
}
__typedArg0 := sympos
__typedArg1 := tmp21780
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21782 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21781)


_ = tmp21782

tmp21783 := MakeNative(func(__e *ControlFlow) {
Y1125 := __e.Get(1)
_ = Y1125
__e.TailApply(PrimFunc(sympreclude_1all_1but), Y1125)
return
}, 1)

tmp21784 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympreclude_1all_1but, tmp21783)
}
__typedArg0 := sympreclude_1all_1but
__typedArg1 := tmp21783
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21785 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21784)


_ = tmp21785

tmp21786 := MakeNative(func(__e *ControlFlow) {
Y1126 := __e.Get(1)
_ = Y1126
__e.TailApply(PrimFunc(symprint), Y1126)
return
}, 1)

tmp21787 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symprint, tmp21786)
}
__typedArg0 := symprint
__typedArg1 := tmp21786
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21788 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21787)


_ = tmp21788

tmp21789 := MakeNative(func(__e *ControlFlow) {
Y1127 := __e.Get(1)
_ = Y1127
__e.TailApply(PrimFunc(symprofile), Y1127)
return
}, 1)

tmp21790 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symprofile, tmp21789)
}
__typedArg0 := symprofile
__typedArg1 := tmp21789
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21791 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21790)


_ = tmp21791

tmp21792 := MakeNative(func(__e *ControlFlow) {
Y1128 := __e.Get(1)
_ = Y1128
__e.TailApply(PrimFunc(symshen_4print_1prolog_1vector), Y1128)
return
}, 1)

tmp21793 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4print_1prolog_1vector, tmp21792)
}
__typedArg0 := symshen_4print_1prolog_1vector
__typedArg1 := tmp21792
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21794 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21793)


_ = tmp21794

tmp21795 := MakeNative(func(__e *ControlFlow) {
Y1129 := __e.Get(1)
_ = Y1129
__e.TailApply(PrimFunc(symshen_4print_1freshterm), Y1129)
return
}, 1)

tmp21796 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4print_1freshterm, tmp21795)
}
__typedArg0 := symshen_4print_1freshterm
__typedArg1 := tmp21795
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21797 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21796)


_ = tmp21797

tmp21798 := MakeNative(func(__e *ControlFlow) {
Y1130 := __e.Get(1)
_ = Y1130
__e.TailApply(PrimFunc(symshen_4printF), Y1130)
return
}, 1)

tmp21799 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4printF, tmp21798)
}
__typedArg0 := symshen_4printF
__typedArg1 := tmp21798
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21800 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21799)


_ = tmp21800

tmp21801 := MakeNative(func(__e *ControlFlow) {
Y1131 := __e.Get(1)
_ = Y1131
__e.TailApply(PrimFunc(symprolog_1memory), Y1131)
return
}, 1)

tmp21802 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symprolog_1memory, tmp21801)
}
__typedArg0 := symprolog_1memory
__typedArg1 := tmp21801
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21803 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21802)


_ = tmp21803

tmp21804 := MakeNative(func(__e *ControlFlow) {
Y1132 := __e.Get(1)
_ = Y1132
__e.TailApply(PrimFunc(symprofile_1results), Y1132)
return
}, 1)

tmp21805 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symprofile_1results, tmp21804)
}
__typedArg0 := symprofile_1results
__typedArg1 := tmp21804
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21806 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21805)


_ = tmp21806

tmp21807 := MakeNative(func(__e *ControlFlow) {
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

tmp21808 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympr, tmp21807)
}
__typedArg0 := sympr
__typedArg1 := tmp21807
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21809 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21808)


_ = tmp21809

tmp21810 := MakeNative(func(__e *ControlFlow) {
Y1135 := __e.Get(1)
_ = Y1135
__e.TailApply(PrimFunc(symps), Y1135)
return
}, 1)

tmp21811 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symps, tmp21810)
}
__typedArg0 := symps
__typedArg1 := tmp21810
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21812 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21811)


_ = tmp21812

tmp21813 := MakeNative(func(__e *ControlFlow) {
Y1136 := __e.Get(1)
_ = Y1136
__e.TailApply(PrimFunc(sympreclude), Y1136)
return
}, 1)

tmp21814 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympreclude, tmp21813)
}
__typedArg0 := sympreclude
__typedArg1 := tmp21813
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21815 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21814)


_ = tmp21815

tmp21816 := MakeNative(func(__e *ControlFlow) {
Y1137 := __e.Get(1)
_ = Y1137
__e.TailApply(PrimFunc(sympreclude_1all_1but), Y1137)
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
Y1138 := __e.Get(1)
_ = Y1138
__e.TailApply(PrimFunc(symprotect), Y1138)
return
}, 1)

tmp21820 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symprotect, tmp21819)
}
__typedArg0 := symprotect
__typedArg1 := tmp21819
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21821 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21820)


_ = tmp21821

tmp21822 := MakeNative(func(__e *ControlFlow) {
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

tmp21823 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symput, tmp21822)
}
__typedArg0 := symput
__typedArg1 := tmp21822
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21824 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21823)


_ = tmp21824

tmp21825 := MakeNative(func(__e *ControlFlow) {
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

tmp21826 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1file_1as_1string, tmp21825)
}
__typedArg0 := symread_1file_1as_1string
__typedArg1 := tmp21825
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21827 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21826)


_ = tmp21827

tmp21828 := MakeNative(func(__e *ControlFlow) {
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

tmp21829 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1file_1as_1bytelist, tmp21828)
}
__typedArg0 := symread_1file_1as_1bytelist
__typedArg1 := tmp21828
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21830 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21829)


_ = tmp21830

tmp21831 := MakeNative(func(__e *ControlFlow) {
Y1145 := __e.Get(1)
_ = Y1145
__e.TailApply(PrimFunc(symread_1file), Y1145)
return
}, 1)

tmp21832 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1file, tmp21831)
}
__typedArg0 := symread_1file
__typedArg1 := tmp21831
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21833 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21832)


_ = tmp21833

tmp21834 := MakeNative(func(__e *ControlFlow) {
Y1146 := __e.Get(1)
_ = Y1146
__e.TailApply(PrimFunc(symread), Y1146)
return
}, 1)

tmp21835 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread, tmp21834)
}
__typedArg0 := symread
__typedArg1 := tmp21834
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21836 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21835)


_ = tmp21836

tmp21837 := MakeNative(func(__e *ControlFlow) {
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

tmp21838 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1byte, tmp21837)
}
__typedArg0 := symread_1byte
__typedArg1 := tmp21837
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21839 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21838)


_ = tmp21839

tmp21840 := MakeNative(func(__e *ControlFlow) {
Y1148 := __e.Get(1)
_ = Y1148
__e.TailApply(PrimFunc(symread_1from_1string), Y1148)
return
}, 1)

tmp21841 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1from_1string, tmp21840)
}
__typedArg0 := symread_1from_1string
__typedArg1 := tmp21840
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21842 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21841)


_ = tmp21842

tmp21843 := MakeNative(func(__e *ControlFlow) {
Y1149 := __e.Get(1)
_ = Y1149
__e.TailApply(PrimFunc(symread_1from_1string_1unprocessed), Y1149)
return
}, 1)

tmp21844 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1from_1string_1unprocessed, tmp21843)
}
__typedArg0 := symread_1from_1string_1unprocessed
__typedArg1 := tmp21843
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21845 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21844)


_ = tmp21845

tmp21846 := MakeNative(func(__e *ControlFlow) {
Y1150 := __e.Get(1)
_ = Y1150
__e.TailApply(PrimFunc(symread_1unit_1string), Y1150)
return
}, 1)

tmp21847 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1unit_1string, tmp21846)
}
__typedArg0 := symread_1unit_1string
__typedArg1 := tmp21846
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21848 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21847)


_ = tmp21848

tmp21849 := MakeNative(func(__e *ControlFlow) {
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

tmp21850 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4register_1source_1form, tmp21849)
}
__typedArg0 := symshen_4register_1source_1form
__typedArg1 := tmp21849
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21851 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21850)


_ = tmp21851

tmp21852 := MakeNative(func(__e *ControlFlow) {
Y1153 := __e.Get(1)
_ = Y1153
__e.TailApply(PrimFunc(symshen_4unregister_1source_1form), Y1153)
return
}, 1)

tmp21853 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4unregister_1source_1form, tmp21852)
}
__typedArg0 := symshen_4unregister_1source_1form
__typedArg1 := tmp21852
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21854 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21853)


_ = tmp21854

tmp21855 := MakeNative(func(__e *ControlFlow) {
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

tmp21856 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symremove, tmp21855)
}
__typedArg0 := symremove
__typedArg1 := tmp21855
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21857 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21856)


_ = tmp21857

tmp21858 := MakeNative(func(__e *ControlFlow) {
Y1156 := __e.Get(1)
_ = Y1156
__e.TailApply(PrimFunc(symreverse), Y1156)
return
}, 1)

tmp21859 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symreverse, tmp21858)
}
__typedArg0 := symreverse
__typedArg1 := tmp21858
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21860 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21859)


_ = tmp21860

tmp21861 := MakeNative(func(__e *ControlFlow) {
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

tmp21862 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symset, tmp21861)
}
__typedArg0 := symset
__typedArg1 := tmp21861
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21863 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21862)


_ = tmp21863

tmp21864 := MakeNative(func(__e *ControlFlow) {
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

tmp21865 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsimple_1error, tmp21864)
}
__typedArg0 := symsimple_1error
__typedArg1 := tmp21864
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21866 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21865)


_ = tmp21866

tmp21867 := MakeNative(func(__e *ControlFlow) {
Y1160 := __e.Get(1)
_ = Y1160
__e.TailApply(PrimFunc(symsnd), Y1160)
return
}, 1)

tmp21868 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsnd, tmp21867)
}
__typedArg0 := symsnd
__typedArg1 := tmp21867
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21869 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21868)


_ = tmp21869

tmp21870 := MakeNative(func(__e *ControlFlow) {
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

tmp21871 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symspecialise, tmp21870)
}
__typedArg0 := symspecialise
__typedArg1 := tmp21870
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21872 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21871)


_ = tmp21872

tmp21873 := MakeNative(func(__e *ControlFlow) {
Y1163 := __e.Get(1)
_ = Y1163
__e.TailApply(PrimFunc(symspy), Y1163)
return
}, 1)

tmp21874 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symspy, tmp21873)
}
__typedArg0 := symspy
__typedArg1 := tmp21873
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21875 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21874)


_ = tmp21875

tmp21876 := MakeNative(func(__e *ControlFlow) {
Y1164 := __e.Get(1)
_ = Y1164
__e.TailApply(PrimFunc(symstep), Y1164)
return
}, 1)

tmp21877 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstep, tmp21876)
}
__typedArg0 := symstep
__typedArg1 := tmp21876
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21878 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21877)


_ = tmp21878

tmp21879 := MakeNative(func(__e *ControlFlow) {
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

tmp21880 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstr, tmp21879)
}
__typedArg0 := symstr
__typedArg1 := tmp21879
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21881 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21880)


_ = tmp21881

tmp21882 := MakeNative(func(__e *ControlFlow) {
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

tmp21883 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring_1_6n, tmp21882)
}
__typedArg0 := symstring_1_6n
__typedArg1 := tmp21882
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21884 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21883)


_ = tmp21884

tmp21885 := MakeNative(func(__e *ControlFlow) {
Y1167 := __e.Get(1)
_ = Y1167
__e.TailApply(PrimFunc(symstring_1_6symbol), Y1167)
return
}, 1)

tmp21886 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring_1_6symbol, tmp21885)
}
__typedArg0 := symstring_1_6symbol
__typedArg1 := tmp21885
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21887 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21886)


_ = tmp21887

tmp21888 := MakeNative(func(__e *ControlFlow) {
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

tmp21889 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring_2, tmp21888)
}
__typedArg0 := symstring_2
__typedArg1 := tmp21888
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21890 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21889)


_ = tmp21890

tmp21891 := MakeNative(func(__e *ControlFlow) {
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

tmp21892 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsubst, tmp21891)
}
__typedArg0 := symsubst
__typedArg1 := tmp21891
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21893 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21892)


_ = tmp21893

tmp21894 := MakeNative(func(__e *ControlFlow) {
Y1172 := __e.Get(1)
_ = Y1172
__e.TailApply(PrimFunc(symsum), Y1172)
return
}, 1)

tmp21895 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsum, tmp21894)
}
__typedArg0 := symsum
__typedArg1 := tmp21894
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21896 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21895)


_ = tmp21896

tmp21897 := MakeNative(func(__e *ControlFlow) {
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

tmp21898 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol_2, tmp21897)
}
__typedArg0 := symsymbol_2
__typedArg1 := tmp21897
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21899 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21898)


_ = tmp21899

tmp21900 := MakeNative(func(__e *ControlFlow) {
Y1174 := __e.Get(1)
_ = Y1174
__e.TailApply(PrimFunc(symsystemf), Y1174)
return
}, 1)

tmp21901 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsystemf, tmp21900)
}
__typedArg0 := symsystemf
__typedArg1 := tmp21900
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21902 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21901)


_ = tmp21902

tmp21903 := MakeNative(func(__e *ControlFlow) {
Y1175 := __e.Get(1)
_ = Y1175
__e.TailApply(PrimFunc(symtail), Y1175)
return
}, 1)

tmp21904 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtail, tmp21903)
}
__typedArg0 := symtail
__typedArg1 := tmp21903
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21905 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21904)


_ = tmp21905

tmp21906 := MakeNative(func(__e *ControlFlow) {
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

tmp21907 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtl, tmp21906)
}
__typedArg0 := symtl
__typedArg1 := tmp21906
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21908 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21907)


_ = tmp21908

tmp21909 := MakeNative(func(__e *ControlFlow) {
Y1177 := __e.Get(1)
_ = Y1177
__e.TailApply(PrimFunc(symtc), Y1177)
return
}, 1)

tmp21910 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtc, tmp21909)
}
__typedArg0 := symtc
__typedArg1 := tmp21909
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21911 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21910)


_ = tmp21911

tmp21912 := MakeNative(func(__e *ControlFlow) {
Y1178 := __e.Get(1)
_ = Y1178
__e.TailApply(PrimFunc(symthaw), Y1178)
return
}, 1)

tmp21913 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symthaw, tmp21912)
}
__typedArg0 := symthaw
__typedArg1 := tmp21912
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21914 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21913)


_ = tmp21914

tmp21915 := MakeNative(func(__e *ControlFlow) {
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

tmp21916 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtlstr, tmp21915)
}
__typedArg0 := symtlstr
__typedArg1 := tmp21915
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21917 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21916)


_ = tmp21917

tmp21918 := MakeNative(func(__e *ControlFlow) {
Y1180 := __e.Get(1)
_ = Y1180
__e.TailApply(PrimFunc(symtrack), Y1180)
return
}, 1)

tmp21919 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtrack, tmp21918)
}
__typedArg0 := symtrack
__typedArg1 := tmp21918
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21920 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21919)


_ = tmp21920

tmp21921 := MakeNative(func(__e *ControlFlow) {
Y1181 := __e.Get(1)
_ = Y1181
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1182 := __e.Get(1)
_ = Y1182
tmp21922 := MakeNative(func(__e *ControlFlow) {
__e.Return(Y1181)
return
}, 0)

__e.TailApply(try_1catch, tmp21922, Y1182)
return


}, 1))
return
}, 1)

tmp21923 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtrap_1error, tmp21921)
}
__typedArg0 := symtrap_1error
__typedArg1 := tmp21921
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21924 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21923)


_ = tmp21924

tmp21925 := MakeNative(func(__e *ControlFlow) {
Y1183 := __e.Get(1)
_ = Y1183
__e.TailApply(PrimFunc(symtuple_2), Y1183)
return
}, 1)

tmp21926 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtuple_2, tmp21925)
}
__typedArg0 := symtuple_2
__typedArg1 := tmp21925
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21927 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21926)


_ = tmp21927

tmp21928 := MakeNative(func(__e *ControlFlow) {
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

tmp21929 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtype, tmp21928)
}
__typedArg0 := symtype
__typedArg1 := tmp21928
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21930 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21929)


_ = tmp21930

tmp21931 := MakeNative(func(__e *ControlFlow) {
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

tmp21932 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symreturn, tmp21931)
}
__typedArg0 := symreturn
__typedArg1 := tmp21931
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21933 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21932)


_ = tmp21933

tmp21934 := MakeNative(func(__e *ControlFlow) {
Y1191 := __e.Get(1)
_ = Y1191
__e.TailApply(PrimFunc(symunabsolute), Y1191)
return
}, 1)

tmp21935 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunabsolute, tmp21934)
}
__typedArg0 := symunabsolute
__typedArg1 := tmp21934
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21936 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21935)


_ = tmp21936

tmp21937 := MakeNative(func(__e *ControlFlow) {
Y1192 := __e.Get(1)
_ = Y1192
__e.TailApply(PrimFunc(symundefmacro), Y1192)
return
}, 1)

tmp21938 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symundefmacro, tmp21937)
}
__typedArg0 := symundefmacro
__typedArg1 := tmp21937
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21939 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21938)


_ = tmp21939

tmp21940 := MakeNative(func(__e *ControlFlow) {
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

tmp21941 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunput, tmp21940)
}
__typedArg0 := symunput
__typedArg1 := tmp21940
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21942 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21941)


_ = tmp21942

tmp21943 := MakeNative(func(__e *ControlFlow) {
Y1196 := __e.Get(1)
_ = Y1196
__e.TailApply(PrimFunc(symunprofile), Y1196)
return
}, 1)

tmp21944 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunprofile, tmp21943)
}
__typedArg0 := symunprofile
__typedArg1 := tmp21943
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21945 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21944)


_ = tmp21945

tmp21946 := MakeNative(func(__e *ControlFlow) {
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

tmp21947 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunion, tmp21946)
}
__typedArg0 := symunion
__typedArg1 := tmp21946
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21948 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21947)


_ = tmp21948

tmp21949 := MakeNative(func(__e *ControlFlow) {
Y1199 := __e.Get(1)
_ = Y1199
__e.TailApply(PrimFunc(symuntrack), Y1199)
return
}, 1)

tmp21950 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symuntrack, tmp21949)
}
__typedArg0 := symuntrack
__typedArg1 := tmp21949
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21951 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21950)


_ = tmp21951

tmp21952 := MakeNative(func(__e *ControlFlow) {
Y1200 := __e.Get(1)
_ = Y1200
__e.TailApply(PrimFunc(symundefmacro), Y1200)
return
}, 1)

tmp21953 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symundefmacro, tmp21952)
}
__typedArg0 := symundefmacro
__typedArg1 := tmp21952
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21954 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21953)


_ = tmp21954

tmp21955 := MakeNative(func(__e *ControlFlow) {
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

tmp21956 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symupdate_1lambda_1table, tmp21955)
}
__typedArg0 := symupdate_1lambda_1table
__typedArg1 := tmp21955
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21957 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21956)


_ = tmp21957

tmp21958 := MakeNative(func(__e *ControlFlow) {
Y1203 := __e.Get(1)
_ = Y1203
__e.TailApply(PrimFunc(symvector), Y1203)
return
}, 1)

tmp21959 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp21958)
}
__typedArg0 := symvector
__typedArg1 := tmp21958
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21960 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21959)


_ = tmp21960

tmp21961 := MakeNative(func(__e *ControlFlow) {
Y1204 := __e.Get(1)
_ = Y1204
__e.TailApply(PrimFunc(symvector_2), Y1204)
return
}, 1)

tmp21962 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector_2, tmp21961)
}
__typedArg0 := symvector_2
__typedArg1 := tmp21961
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21963 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21962)


_ = tmp21963

tmp21964 := MakeNative(func(__e *ControlFlow) {
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

tmp21965 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector_1_6, tmp21964)
}
__typedArg0 := symvector_1_6
__typedArg1 := tmp21964
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21966 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21965)


_ = tmp21966

tmp21967 := MakeNative(func(__e *ControlFlow) {
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

tmp21968 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvalue, tmp21967)
}
__typedArg0 := symvalue
__typedArg1 := tmp21967
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21969 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21968)


_ = tmp21969

tmp21970 := MakeNative(func(__e *ControlFlow) {
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

tmp21971 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvariable_2, tmp21970)
}
__typedArg0 := symvariable_2
__typedArg1 := tmp21970
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21972 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21971)


_ = tmp21972

tmp21973 := MakeNative(func(__e *ControlFlow) {
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

tmp21974 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvar_2, tmp21973)
}
__typedArg0 := symvar_2
__typedArg1 := tmp21973
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21975 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21974)


_ = tmp21975

tmp21976 := MakeNative(func(__e *ControlFlow) {
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

tmp21977 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhen, tmp21976)
}
__typedArg0 := symwhen
__typedArg1 := tmp21976
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21978 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21977)


_ = tmp21978

tmp21979 := MakeNative(func(__e *ControlFlow) {
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

tmp21980 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwrite_1byte, tmp21979)
}
__typedArg0 := symwrite_1byte
__typedArg1 := tmp21979
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21981 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21980)


_ = tmp21981

tmp21982 := MakeNative(func(__e *ControlFlow) {
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

tmp21983 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwrite_1string, tmp21982)
}
__typedArg0 := symwrite_1string
__typedArg1 := tmp21982
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21984 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21983)


_ = tmp21984

tmp21985 := MakeNative(func(__e *ControlFlow) {
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

tmp21986 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwrite_1to_1file, tmp21985)
}
__typedArg0 := symwrite_1to_1file
__typedArg1 := tmp21985
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21987 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21986)


_ = tmp21987

tmp21988 := MakeNative(func(__e *ControlFlow) {
Y1226 := __e.Get(1)
_ = Y1226
__e.TailApply(PrimFunc(symy_1or_1n_2), Y1226)
return
}, 1)

tmp21989 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symy_1or_1n_2, tmp21988)
}
__typedArg0 := symy_1or_1n_2
__typedArg1 := tmp21988
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21990 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21989)


_ = tmp21990

tmp21991 := MakeNative(func(__e *ControlFlow) {
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

tmp21992 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_7, tmp21991)
}
__typedArg0 := sym_7
__typedArg1 := tmp21991
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21993 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21992)


_ = tmp21993

tmp21994 := MakeNative(func(__e *ControlFlow) {
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

tmp21995 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_d, tmp21994)
}
__typedArg0 := sym_d
__typedArg1 := tmp21994
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21996 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21995)


_ = tmp21996

tmp21997 := MakeNative(func(__e *ControlFlow) {
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

tmp21998 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_c, tmp21997)
}
__typedArg0 := sym_c
__typedArg1 := tmp21997
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21999 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21998)


_ = tmp21999

tmp22000 := MakeNative(func(__e *ControlFlow) {
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

tmp22001 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1, tmp22000)
}
__typedArg0 := sym_1
__typedArg1 := tmp22000
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp22002 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp22001)


_ = tmp22002

tmp22003 := MakeNative(func(__e *ControlFlow) {
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

tmp22004 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_a_a, tmp22003)
}
__typedArg0 := sym_a_a
__typedArg1 := tmp22003
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp22005 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp22004)


_ = tmp22005

tmp22006 := MakeNative(func(__e *ControlFlow) {
Y1237 := __e.Get(1)
_ = Y1237
__e.TailApply(PrimFunc(sym_5e_6), Y1237)
return
}, 1)

tmp22007 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5e_6, tmp22006)
}
__typedArg0 := sym_5e_6
__typedArg1 := tmp22006
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp22008 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp22007)


_ = tmp22008

tmp22009 := MakeNative(func(__e *ControlFlow) {
Y1238 := __e.Get(1)
_ = Y1238
__e.TailApply(PrimFunc(sym_5end_6), Y1238)
return
}, 1)

tmp22010 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5end_6, tmp22009)
}
__typedArg0 := sym_5end_6
__typedArg1 := tmp22009
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp22011 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp22010)


_ = tmp22011

tmp22012 := MakeNative(func(__e *ControlFlow) {
Y1239 := __e.Get(1)
_ = Y1239
__e.TailApply(PrimFunc(sym_5_b_6), Y1239)
return
}, 1)

tmp22013 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_b_6, tmp22012)
}
__typedArg0 := sym_5_b_6
__typedArg1 := tmp22012
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp22014 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp22013)


_ = tmp22014

tmp22015 := MakeNative(func(__e *ControlFlow) {
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

tmp22016 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8p, tmp22015)
}
__typedArg0 := sym_8p
__typedArg1 := tmp22015
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp22017 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp22016)


_ = tmp22017

tmp22018 := MakeNative(func(__e *ControlFlow) {
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

tmp22019 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8v, tmp22018)
}
__typedArg0 := sym_8v
__typedArg1 := tmp22018
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp22020 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp22019)


_ = tmp22020

tmp22021 := MakeNative(func(__e *ControlFlow) {
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

tmp22022 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8s, tmp22021)
}
__typedArg0 := sym_8s
__typedArg1 := tmp22021
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4set_1lambda_1form_1entry), tmp22022)
return


}, 0)

tmp22023 := Call(__e, ns2_1set, symshen_4initialise_1lambda_1forms, tmp21468)


_ = tmp22023

tmp22024 := MakeNative(func(__e *ControlFlow) {
tmp22025 := Call(__e, PrimFunc(symshen_4initialise_1environment))


_ = tmp22025

tmp22026 := Call(__e, PrimFunc(symshen_4initialise_1lambda_1forms))


_ = tmp22026

__e.TailApply(PrimFunc(symshen_4initialise_1signedfuncs))
return


}, 0)

__e.TailApply(ns2_1set, symshen_4initialise, tmp22024)
return




}, 0)

