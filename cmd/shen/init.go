package main

import . "github.com/pyrex41/shen-go/kl"

var InitMain = MakeNative(func(__e *ControlFlow) {
tmp18975 := MakeNative(func(__e *ControlFlow) {
tmp18976 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dhistory_d, Nil)
}
__typedArg0 := symshen_4_dhistory_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp18976

tmp18977 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dtc_d, False)
}
__typedArg0 := symshen_4_dtc_d
__typedArg1 := False
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp18977

tmp18978 := Call(__e, PrimFunc(symshen_4dict), MakeInteger(20000))


tmp18979 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(sym_dproperty_1vector_d, tmp18978)
}
__typedArg0 := sym_dproperty_1vector_d
__typedArg1 := tmp18978
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp18979

tmp18980 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
__e.TailApply(PrimFunc(symshen_4macros), X)
return
}, 1)

tmp18981 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4macros, tmp18980)
}
__typedArg0 := symshen_4macros
__typedArg1 := tmp18980
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18982 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18981, Nil)
}
__typedArg0 := tmp18981
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18983 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(sym_dmacros_d, tmp18982)
}
__typedArg0 := sym_dmacros_d
__typedArg1 := tmp18982
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp18983

tmp18984 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsource_1form_1handlers_d, Nil)
}
__typedArg0 := symshen_4_dsource_1form_1handlers_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp18984

tmp18985 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dgensym_d, MakeInteger(0))
}
__typedArg0 := symshen_4_dgensym_d
__typedArg1 := MakeInteger(0)
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp18985

tmp18986 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dtracking_d, Nil)
}
__typedArg0 := symshen_4_dtracking_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp18986

tmp18987 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dprofiled_d, Nil)
}
__typedArg0 := symshen_4_dprofiled_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp18987

tmp18988 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtype, Nil)
}
__typedArg0 := symtype
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18989 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4input_1h_7, tmp18988)
}
__typedArg0 := symshen_4input_1h_7
__typedArg1 := tmp18988
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18990 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symopen, tmp18989)
}
__typedArg0 := symopen
__typedArg1 := tmp18989
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18991 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symset, tmp18990)
}
__typedArg0 := symset
__typedArg1 := tmp18990
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18992 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhere, tmp18991)
}
__typedArg0 := symwhere
__typedArg1 := tmp18991
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18993 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp18992)
}
__typedArg0 := symlet
__typedArg1 := tmp18992
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18994 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp18993)
}
__typedArg0 := symlambda
__typedArg1 := tmp18993
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18995 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp18994)
}
__typedArg0 := symcons
__typedArg1 := tmp18994
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18996 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8v, tmp18995)
}
__typedArg0 := sym_8v
__typedArg1 := tmp18995
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18997 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8s, tmp18996)
}
__typedArg0 := sym_8s
__typedArg1 := tmp18996
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18998 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8p, tmp18997)
}
__typedArg0 := sym_8p
__typedArg1 := tmp18997
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18999 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dspecial_d, tmp18998)
}
__typedArg0 := symshen_4_dspecial_d
__typedArg1 := tmp18998
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp18999

tmp19000 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dextraspecial_d, Nil)
}
__typedArg0 := symshen_4_dextraspecial_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19000

tmp19001 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dspy_d, False)
}
__typedArg0 := symshen_4_dspy_d
__typedArg1 := False
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19001

tmp19002 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_ddatatypes_d, Nil)
}
__typedArg0 := symshen_4_ddatatypes_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19002

tmp19003 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dalldatatypes_d, Nil)
}
__typedArg0 := symshen_4_dalldatatypes_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19003

tmp19004 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dshen_1type_1theory_1enabled_2_d, True)
}
__typedArg0 := symshen_4_dshen_1type_1theory_1enabled_2_d
__typedArg1 := True
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19004

tmp19005 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dpackage_d, symnull)
}
__typedArg0 := symshen_4_dpackage_d
__typedArg1 := symnull
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19005

tmp19006 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsynonyms_d, Nil)
}
__typedArg0 := symshen_4_dsynonyms_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19006

tmp19007 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsystem_d, Nil)
}
__typedArg0 := symshen_4_dsystem_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19007

tmp19008 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_doccurs_d, True)
}
__typedArg0 := symshen_4_doccurs_d
__typedArg1 := True
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19008

tmp19009 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dfactorise_2_d, False)
}
__typedArg0 := symshen_4_dfactorise_2_d
__typedArg1 := False
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19009

tmp19010 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dmaxinferences_d, MakeInteger(1000000))
}
__typedArg0 := symshen_4_dmaxinferences_d
__typedArg1 := MakeInteger(1000000)
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19010

tmp19011 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(sym_dmaximum_1print_1sequence_1size_d, MakeInteger(20))
}
__typedArg0 := sym_dmaximum_1print_1sequence_1size_d
__typedArg1 := MakeInteger(20)
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19011

tmp19012 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dcall_d, MakeInteger(0))
}
__typedArg0 := symshen_4_dcall_d
__typedArg1 := MakeInteger(0)
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19012

tmp19013 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dinfs_d, MakeInteger(0))
}
__typedArg0 := symshen_4_dinfs_d
__typedArg1 := MakeInteger(0)
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19013

tmp19014 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(sym_dhush_d, False)
}
__typedArg0 := sym_dhush_d
__typedArg1 := False
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19014

tmp19015 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_doptimise_d, False)
}
__typedArg0 := symshen_4_doptimise_d
__typedArg1 := False
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19015

tmp19016 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(sym_dversion_d, MakeString("42.2"))
}
__typedArg0 := sym_dversion_d
__typedArg1 := MakeString("42.2")
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19016

tmp19017 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dnames_d, Nil)
}
__typedArg0 := symshen_4_dnames_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19017

tmp19018 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dstep_d, False)
}
__typedArg0 := symshen_4_dstep_d
__typedArg1 := False
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19018

tmp19019 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dit_d, MakeString(""))
}
__typedArg0 := symshen_4_dit_d
__typedArg1 := MakeString("")
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19019

tmp19020 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dresidue_d, Nil)
}
__typedArg0 := symshen_4_dresidue_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19020

tmp19021 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(sym_dabsolute_d, Nil)
}
__typedArg0 := sym_dabsolute_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19021

tmp19022 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dprolog_1memory_d, MakeInteger(1000))
}
__typedArg0 := symshen_4_dprolog_1memory_d
__typedArg1 := MakeInteger(1000)
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19022

tmp19023 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dloading_2_d, False)
}
__typedArg0 := symshen_4_dloading_2_d
__typedArg1 := False
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19023

tmp19024 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_duserdefs_d, Nil)
}
__typedArg0 := symshen_4_duserdefs_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19024

tmp19025 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
__e.Return(X)
return
}, 1)

tmp19026 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_ddemodulation_1function_d, tmp19025)
}
__typedArg0 := symshen_4_ddemodulation_1function_d
__typedArg1 := tmp19025
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19026

tmp19027 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dcustom_1pattern_1compiler_d, False)
}
__typedArg0 := symshen_4_dcustom_1pattern_1compiler_d
__typedArg1 := False
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19027

tmp19028 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dcustom_1pattern_1reducer_d, False)
}
__typedArg0 := symshen_4_dcustom_1pattern_1reducer_d
__typedArg1 := False
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19028

tmp19031 := Call(__e, PrimFunc(symbound_2), sym_dhome_1directory_d)


tmp19032 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp19031)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp19031
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres19029 Obj

if True == tmp19032 {
tmp19030 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(sym_dhome_1directory_d, MakeString(""))
}
__typedArg0 := sym_dhome_1directory_d
__typedArg1 := MakeString("")
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

ifres19029 = tmp19030


} else {
ifres19029 = symshen_4skip


}

_ = ifres19029

tmp19036 := Call(__e, PrimFunc(symbound_2), sym_dsterror_d)


tmp19037 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp19036)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp19036
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres19033 Obj

if True == tmp19037 {
tmp19034 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dstoutput_d)
}
__typedArg0 := sym_dstoutput_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp19035 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(sym_dsterror_d, tmp19034)
}
__typedArg0 := sym_dsterror_d
__typedArg1 := tmp19034
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

ifres19033 = tmp19035


} else {
ifres19033 = symshen_4skip


}

_ = ifres19033

tmp19038 := Call(__e, PrimFunc(symprolog_1memory), MakeInteger(10000))


_ = tmp19038

tmp19039 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dloading_2_d, False)
}
__typedArg0 := symshen_4_dloading_2_d
__typedArg1 := False
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19039

tmp19040 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), Nil)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19041 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8s, tmp19040)
}
__typedArg0 := sym_8s
__typedArg1 := tmp19040
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19042 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19041)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19041
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19043 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8v, tmp19042)
}
__typedArg0 := sym_8v
__typedArg1 := tmp19042
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19044 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19043)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19043
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19045 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8p, tmp19044)
}
__typedArg0 := sym_8p
__typedArg1 := tmp19044
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19046 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19045)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19045
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19047 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_b_6, tmp19046)
}
__typedArg0 := sym_5_b_6
__typedArg1 := tmp19046
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19048 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19047)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19047
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19049 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5end_6, tmp19048)
}
__typedArg0 := sym_5end_6
__typedArg1 := tmp19048
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19050 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19049)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19049
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19051 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5e_6, tmp19050)
}
__typedArg0 := sym_5e_6
__typedArg1 := tmp19050
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19052 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19051)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19051
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19053 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_a_a, tmp19052)
}
__typedArg0 := sym_a_a
__typedArg1 := tmp19052
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19054 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19053)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19053
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19055 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1, tmp19054)
}
__typedArg0 := sym_1
__typedArg1 := tmp19054
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19056 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19055)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19055
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19057 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_c, tmp19056)
}
__typedArg0 := sym_c
__typedArg1 := tmp19056
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19058 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19057)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19057
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19059 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_d, tmp19058)
}
__typedArg0 := sym_d
__typedArg1 := tmp19058
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19060 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19059)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19059
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19061 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_7, tmp19060)
}
__typedArg0 := sym_7
__typedArg1 := tmp19060
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19062 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19061)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19061
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19063 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symy_1or_1n_2, tmp19062)
}
__typedArg0 := symy_1or_1n_2
__typedArg1 := tmp19062
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19064 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19063)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19063
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19065 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwrite_1to_1file, tmp19064)
}
__typedArg0 := symwrite_1to_1file
__typedArg1 := tmp19064
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19066 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19065)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19065
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19067 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwrite_1byte, tmp19066)
}
__typedArg0 := symwrite_1byte
__typedArg1 := tmp19066
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19068 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(5), tmp19067)
}
__typedArg0 := MakeInteger(5)
__typedArg1 := tmp19067
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19069 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhen, tmp19068)
}
__typedArg0 := symwhen
__typedArg1 := tmp19068
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19070 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19069)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19069
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19071 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symversion, tmp19070)
}
__typedArg0 := symversion
__typedArg1 := tmp19070
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19072 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(5), tmp19071)
}
__typedArg0 := MakeInteger(5)
__typedArg1 := tmp19071
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19073 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvar_2, tmp19072)
}
__typedArg0 := symvar_2
__typedArg1 := tmp19072
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19074 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19073)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19073
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19075 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvariable_2, tmp19074)
}
__typedArg0 := symvariable_2
__typedArg1 := tmp19074
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19076 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19075)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19075
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19077 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvalue, tmp19076)
}
__typedArg0 := symvalue
__typedArg1 := tmp19076
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19078 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(3), tmp19077)
}
__typedArg0 := MakeInteger(3)
__typedArg1 := tmp19077
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19079 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector_1_6, tmp19078)
}
__typedArg0 := symvector_1_6
__typedArg1 := tmp19078
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19080 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19079)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19079
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19081 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector_2, tmp19080)
}
__typedArg0 := symvector_2
__typedArg1 := tmp19080
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19082 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19081)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19081
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19083 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp19082)
}
__typedArg0 := symvector
__typedArg1 := tmp19082
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19084 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19083)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19083
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19085 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symuserdefs, tmp19084)
}
__typedArg0 := symuserdefs
__typedArg1 := tmp19084
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19086 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19085)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19085
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19087 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symupdate_1lambda_1table, tmp19086)
}
__typedArg0 := symupdate_1lambda_1table
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
return PrimCons(symundefmacro, tmp19088)
}
__typedArg0 := symundefmacro
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
return PrimCons(symuntrack, tmp19090)
}
__typedArg0 := symuntrack
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
return PrimCons(symunion, tmp19092)
}
__typedArg0 := symunion
__typedArg1 := tmp19092
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19094 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19093)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19093
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19095 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunprofile, tmp19094)
}
__typedArg0 := symunprofile
__typedArg1 := tmp19094
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19096 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(3), tmp19095)
}
__typedArg0 := MakeInteger(3)
__typedArg1 := tmp19095
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19097 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunput, tmp19096)
}
__typedArg0 := symunput
__typedArg1 := tmp19096
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19098 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19097)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19097
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19099 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symundefmacro, tmp19098)
}
__typedArg0 := symundefmacro
__typedArg1 := tmp19098
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19100 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19099)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19099
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19101 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunabsolute, tmp19100)
}
__typedArg0 := symunabsolute
__typedArg1 := tmp19100
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19102 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(5), tmp19101)
}
__typedArg0 := MakeInteger(5)
__typedArg1 := tmp19101
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19103 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symreturn, tmp19102)
}
__typedArg0 := symreturn
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
return PrimCons(symtype, tmp19104)
}
__typedArg0 := symtype
__typedArg1 := tmp19104
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19106 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19105)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19105
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19107 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtuple_2, tmp19106)
}
__typedArg0 := symtuple_2
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
return PrimCons(symtrap_1error, tmp19108)
}
__typedArg0 := symtrap_1error
__typedArg1 := tmp19108
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19110 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19109)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19109
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19111 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtracked, tmp19110)
}
__typedArg0 := symtracked
__typedArg1 := tmp19110
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19112 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19111)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19111
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19113 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtrack, tmp19112)
}
__typedArg0 := symtrack
__typedArg1 := tmp19112
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19114 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19113)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19113
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19115 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtlstr, tmp19114)
}
__typedArg0 := symtlstr
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
return PrimCons(symthaw, tmp19116)
}
__typedArg0 := symthaw
__typedArg1 := tmp19116
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19118 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19117)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19117
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19119 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtc_2, tmp19118)
}
__typedArg0 := symtc_2
__typedArg1 := tmp19118
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19120 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19119)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19119
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19121 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtc, tmp19120)
}
__typedArg0 := symtc
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
return PrimCons(symtl, tmp19122)
}
__typedArg0 := symtl
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
return PrimCons(symtail, tmp19124)
}
__typedArg0 := symtail
__typedArg1 := tmp19124
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19126 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19125)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19125
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19127 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsystemf, tmp19126)
}
__typedArg0 := symsystemf
__typedArg1 := tmp19126
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19128 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19127)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19127
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19129 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol_2, tmp19128)
}
__typedArg0 := symsymbol_2
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
return PrimCons(symsum, tmp19130)
}
__typedArg0 := symsum
__typedArg1 := tmp19130
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19132 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(3), tmp19131)
}
__typedArg0 := MakeInteger(3)
__typedArg1 := tmp19131
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19133 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsubst, tmp19132)
}
__typedArg0 := symsubst
__typedArg1 := tmp19132
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19134 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19133)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19133
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19135 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring_2, tmp19134)
}
__typedArg0 := symstring_2
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
return PrimCons(symstring_1_6symbol, tmp19136)
}
__typedArg0 := symstring_1_6symbol
__typedArg1 := tmp19136
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19138 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19137)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19137
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19139 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring_1_6n, tmp19138)
}
__typedArg0 := symstring_1_6n
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
return PrimCons(symstr, tmp19140)
}
__typedArg0 := symstr
__typedArg1 := tmp19140
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19142 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19141)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19141
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19143 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstoutput, tmp19142)
}
__typedArg0 := symstoutput
__typedArg1 := tmp19142
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19144 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19143)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19143
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19145 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstinput, tmp19144)
}
__typedArg0 := symstinput
__typedArg1 := tmp19144
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19146 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19145)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19145
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19147 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4step_2, tmp19146)
}
__typedArg0 := symshen_4step_2
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
return PrimCons(symstep, tmp19148)
}
__typedArg0 := symstep
__typedArg1 := tmp19148
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19150 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19149)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19149
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19151 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4spy_2, tmp19150)
}
__typedArg0 := symshen_4spy_2
__typedArg1 := tmp19150
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19152 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19151)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19151
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19153 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symspy, tmp19152)
}
__typedArg0 := symspy
__typedArg1 := tmp19152
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19154 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19153)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19153
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19155 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symspecialise, tmp19154)
}
__typedArg0 := symspecialise
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
return PrimCons(symsnd, tmp19156)
}
__typedArg0 := symsnd
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
return PrimCons(symsimple_1error, tmp19158)
}
__typedArg0 := symsimple_1error
__typedArg1 := tmp19158
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19160 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19159)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19159
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19161 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symset, tmp19160)
}
__typedArg0 := symset
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
return PrimCons(symreverse, tmp19162)
}
__typedArg0 := symreverse
__typedArg1 := tmp19162
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19164 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19163)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19163
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19165 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symremove, tmp19164)
}
__typedArg0 := symremove
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
return PrimCons(symshen_4unregister_1source_1form, tmp19166)
}
__typedArg0 := symshen_4unregister_1source_1form
__typedArg1 := tmp19166
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19168 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19167)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19167
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19169 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4register_1source_1form, tmp19168)
}
__typedArg0 := symshen_4register_1source_1form
__typedArg1 := tmp19168
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19170 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19169)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19169
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19171 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symrelease, tmp19170)
}
__typedArg0 := symrelease
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
return PrimCons(symreceive, tmp19172)
}
__typedArg0 := symreceive
__typedArg1 := tmp19172
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19174 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19173)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19173
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19175 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4read_1unit_1string, tmp19174)
}
__typedArg0 := symshen_4read_1unit_1string
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
return PrimCons(symread_1from_1string_1unprocessed, tmp19176)
}
__typedArg0 := symread_1from_1string_1unprocessed
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
return PrimCons(symread_1from_1string, tmp19178)
}
__typedArg0 := symread_1from_1string
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
return PrimCons(symread_1byte, tmp19180)
}
__typedArg0 := symread_1byte
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
return PrimCons(symread, tmp19182)
}
__typedArg0 := symread
__typedArg1 := tmp19182
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19184 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19183)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19183
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19185 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1file, tmp19184)
}
__typedArg0 := symread_1file
__typedArg1 := tmp19184
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19186 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19185)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19185
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19187 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1file_1as_1bytelist, tmp19186)
}
__typedArg0 := symread_1file_1as_1bytelist
__typedArg1 := tmp19186
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19188 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19187)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19187
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19189 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1file_1as_1string, tmp19188)
}
__typedArg0 := symread_1file_1as_1string
__typedArg1 := tmp19188
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19190 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(4), tmp19189)
}
__typedArg0 := MakeInteger(4)
__typedArg1 := tmp19189
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19191 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symput, tmp19190)
}
__typedArg0 := symput
__typedArg1 := tmp19190
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19192 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19191)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19191
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19193 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symprotect, tmp19192)
}
__typedArg0 := symprotect
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
return PrimCons(sympreclude_1all_1but, tmp19194)
}
__typedArg0 := sympreclude_1all_1but
__typedArg1 := tmp19194
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19196 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19195)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19195
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19197 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympreclude, tmp19196)
}
__typedArg0 := sympreclude
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
return PrimCons(symps, tmp19198)
}
__typedArg0 := symps
__typedArg1 := tmp19198
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19200 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19199)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19199
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19201 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympr, tmp19200)
}
__typedArg0 := sympr
__typedArg1 := tmp19200
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19202 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19201)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19201
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19203 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symprofile_1results, tmp19202)
}
__typedArg0 := symprofile_1results
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
return PrimCons(symprolog_1memory, tmp19204)
}
__typedArg0 := symprolog_1memory
__typedArg1 := tmp19204
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19206 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19205)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19205
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19207 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4printF, tmp19206)
}
__typedArg0 := symshen_4printF
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
return PrimCons(symshen_4print_1freshterm, tmp19208)
}
__typedArg0 := symshen_4print_1freshterm
__typedArg1 := tmp19208
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19210 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19209)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19209
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19211 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4print_1prolog_1vector, tmp19210)
}
__typedArg0 := symshen_4print_1prolog_1vector
__typedArg1 := tmp19210
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19212 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19211)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19211
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19213 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symprofile, tmp19212)
}
__typedArg0 := symprofile
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
return PrimCons(symprint, tmp19214)
}
__typedArg0 := symprint
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
return PrimCons(sympreclude_1all_1but, tmp19216)
}
__typedArg0 := sympreclude_1all_1but
__typedArg1 := tmp19216
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19218 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19217)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19217
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19219 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympos, tmp19218)
}
__typedArg0 := sympos
__typedArg1 := tmp19218
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19220 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19219)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19219
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19221 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symporters, tmp19220)
}
__typedArg0 := symporters
__typedArg1 := tmp19220
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19222 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19221)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19221
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19223 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symport, tmp19222)
}
__typedArg0 := symport
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
return PrimCons(sympackage_2, tmp19224)
}
__typedArg0 := sympackage_2
__typedArg1 := tmp19224
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19226 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(3), tmp19225)
}
__typedArg0 := MakeInteger(3)
__typedArg1 := tmp19225
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19227 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympackage, tmp19226)
}
__typedArg0 := sympackage
__typedArg1 := tmp19226
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19228 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19227)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19227
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19229 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symos, tmp19228)
}
__typedArg0 := symos
__typedArg1 := tmp19228
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19230 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19229)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19229
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19231 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symor, tmp19230)
}
__typedArg0 := symor
__typedArg1 := tmp19230
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19232 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19231)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19231
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19233 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symoptimise_2, tmp19232)
}
__typedArg0 := symoptimise_2
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
return PrimCons(symoptimise, tmp19234)
}
__typedArg0 := symoptimise
__typedArg1 := tmp19234
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19236 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19235)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19235
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19237 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symopen, tmp19236)
}
__typedArg0 := symopen
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
return PrimCons(symoccurs_1check, tmp19238)
}
__typedArg0 := symoccurs_1check
__typedArg1 := tmp19238
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19240 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19239)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19239
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19241 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symoccurs_2, tmp19240)
}
__typedArg0 := symoccurs_2
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
return PrimCons(symoccurrences, tmp19242)
}
__typedArg0 := symoccurrences
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
return PrimCons(symoccurs_1check, tmp19244)
}
__typedArg0 := symoccurs_1check
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
return PrimCons(symnumber_2, tmp19246)
}
__typedArg0 := symnumber_2
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
return PrimCons(symn_1_6string, tmp19248)
}
__typedArg0 := symn_1_6string
__typedArg1 := tmp19248
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19250 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19249)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19249
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19251 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnth, tmp19250)
}
__typedArg0 := symnth
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
return PrimCons(symnot, tmp19252)
}
__typedArg0 := symnot
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
return PrimCons(symnl, tmp19254)
}
__typedArg0 := symnl
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
return PrimCons(symmaxinferences, tmp19256)
}
__typedArg0 := symmaxinferences
__typedArg1 := tmp19256
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19258 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19257)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19257
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19259 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symmapcan, tmp19258)
}
__typedArg0 := symmapcan
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
return PrimCons(symmap, tmp19260)
}
__typedArg0 := symmap
__typedArg1 := tmp19260
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19262 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19261)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19261
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19263 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symmacroexpand, tmp19262)
}
__typedArg0 := symmacroexpand
__typedArg1 := tmp19262
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19264 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19263)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19263
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19265 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp19264)
}
__typedArg0 := symvector
__typedArg1 := tmp19264
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19266 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19265)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19265
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19267 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_a, tmp19266)
}
__typedArg0 := sym_5_a
__typedArg1 := tmp19266
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19268 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19267)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19267
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19269 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5, tmp19268)
}
__typedArg0 := sym_5
__typedArg1 := tmp19268
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19270 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19269)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19269
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19271 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symload, tmp19270)
}
__typedArg0 := symload
__typedArg1 := tmp19270
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19272 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19271)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19271
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19273 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19272)
}
__typedArg0 := symlist
__typedArg1 := tmp19272
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19274 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19273)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19273
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19275 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlineread, tmp19274)
}
__typedArg0 := symlineread
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
return PrimCons(symlimit, tmp19276)
}
__typedArg0 := symlimit
__typedArg1 := tmp19276
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19278 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19277)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19277
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19279 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlength, tmp19278)
}
__typedArg0 := symlength
__typedArg1 := tmp19278
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19280 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19279)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19279
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19281 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlanguage, tmp19280)
}
__typedArg0 := symlanguage
__typedArg1 := tmp19280
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19282 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(6), tmp19281)
}
__typedArg0 := MakeInteger(6)
__typedArg1 := tmp19281
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19283 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symis_b, tmp19282)
}
__typedArg0 := symis_b
__typedArg1 := tmp19282
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19284 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(6), tmp19283)
}
__typedArg0 := MakeInteger(6)
__typedArg1 := tmp19283
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19285 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symis, tmp19284)
}
__typedArg0 := symis
__typedArg1 := tmp19284
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19286 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19285)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19285
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19287 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symit, tmp19286)
}
__typedArg0 := symit
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
return PrimCons(syminternal, tmp19288)
}
__typedArg0 := syminternal
__typedArg1 := tmp19288
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19290 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19289)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19289
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19291 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symintersection, tmp19290)
}
__typedArg0 := symintersection
__typedArg1 := tmp19290
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19292 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19291)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19291
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19293 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminclude_1all_1but, tmp19292)
}
__typedArg0 := syminclude_1all_1but
__typedArg1 := tmp19292
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19294 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19293)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19293
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19295 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symimplementation, tmp19294)
}
__typedArg0 := symimplementation
__typedArg1 := tmp19294
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19296 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19295)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19295
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19297 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminput_7, tmp19296)
}
__typedArg0 := syminput_7
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
return PrimCons(syminput, tmp19298)
}
__typedArg0 := syminput
__typedArg1 := tmp19298
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19300 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19299)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19299
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19301 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminferences, tmp19300)
}
__typedArg0 := syminferences
__typedArg1 := tmp19300
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19302 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19301)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19301
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19303 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symintern, tmp19302)
}
__typedArg0 := symintern
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
return PrimCons(syminternal, tmp19304)
}
__typedArg0 := syminternal
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
return PrimCons(syminteger_2, tmp19306)
}
__typedArg0 := syminteger_2
__typedArg1 := tmp19306
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19308 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19307)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19307
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19309 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symin_1package, tmp19308)
}
__typedArg0 := symin_1package
__typedArg1 := tmp19308
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19310 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19309)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19309
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19311 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4included, tmp19310)
}
__typedArg0 := symshen_4included
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
return PrimCons(syminclude, tmp19312)
}
__typedArg0 := syminclude
__typedArg1 := tmp19312
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19314 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(3), tmp19313)
}
__typedArg0 := MakeInteger(3)
__typedArg1 := tmp19313
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19315 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp19314)
}
__typedArg0 := symif
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
return PrimCons(symhush, tmp19316)
}
__typedArg0 := symhush
__typedArg1 := tmp19316
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19318 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19317)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19317
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19319 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhush_2, tmp19318)
}
__typedArg0 := symhush_2
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
return PrimCons(symhead, tmp19320)
}
__typedArg0 := symhead
__typedArg1 := tmp19320
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19322 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19321)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19321
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19323 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhdstr, tmp19322)
}
__typedArg0 := symhdstr
__typedArg1 := tmp19322
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19324 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19323)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19323
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19325 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhdv, tmp19324)
}
__typedArg0 := symhdv
__typedArg1 := tmp19324
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19326 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19325)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19325
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19327 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhd, tmp19326)
}
__typedArg0 := symhd
__typedArg1 := tmp19326
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19328 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19327)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19327
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19329 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhash, tmp19328)
}
__typedArg0 := symhash
__typedArg1 := tmp19328
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19330 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19329)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19329
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19331 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_a, tmp19330)
}
__typedArg0 := sym_a
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
return PrimCons(sym_6_a, tmp19332)
}
__typedArg0 := sym_6_a
__typedArg1 := tmp19332
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19334 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19333)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19333
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19335 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_6, tmp19334)
}
__typedArg0 := sym_6
__typedArg1 := tmp19334
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19336 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19335)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19335
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19337 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1vector, tmp19336)
}
__typedArg0 := sym_5_1vector
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
return PrimCons(sym_5_1address, tmp19338)
}
__typedArg0 := sym_5_1address
__typedArg1 := tmp19338
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19340 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(3), tmp19339)
}
__typedArg0 := MakeInteger(3)
__typedArg1 := tmp19339
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19341 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symaddress_1_6, tmp19340)
}
__typedArg0 := symaddress_1_6
__typedArg1 := tmp19340
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19342 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19341)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19341
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19343 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symget_1time, tmp19342)
}
__typedArg0 := symget_1time
__typedArg1 := tmp19342
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19344 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(3), tmp19343)
}
__typedArg0 := MakeInteger(3)
__typedArg1 := tmp19343
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19345 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symget, tmp19344)
}
__typedArg0 := symget
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
return PrimCons(symgensym, tmp19346)
}
__typedArg0 := symgensym
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
return PrimCons(symfunction, tmp19348)
}
__typedArg0 := symfunction
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
return PrimCons(symfn, tmp19350)
}
__typedArg0 := symfn
__typedArg1 := tmp19350
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19352 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19351)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19351
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19353 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfst, tmp19352)
}
__typedArg0 := symfst
__typedArg1 := tmp19352
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19354 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19353)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19353
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19355 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfresh, tmp19354)
}
__typedArg0 := symfresh
__typedArg1 := tmp19354
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19356 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19355)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19355
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19357 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfreeze, tmp19356)
}
__typedArg0 := symfreeze
__typedArg1 := tmp19356
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19358 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(5), tmp19357)
}
__typedArg0 := MakeInteger(5)
__typedArg1 := tmp19357
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19359 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfork, tmp19358)
}
__typedArg0 := symfork
__typedArg1 := tmp19358
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19360 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19359)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19359
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19361 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symforeign, tmp19360)
}
__typedArg0 := symforeign
__typedArg1 := tmp19360
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19362 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(7), tmp19361)
}
__typedArg0 := MakeInteger(7)
__typedArg1 := tmp19361
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19363 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfindall, tmp19362)
}
__typedArg0 := symfindall
__typedArg1 := tmp19362
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19364 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19363)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19363
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19365 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfix, tmp19364)
}
__typedArg0 := symfix
__typedArg1 := tmp19364
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19366 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19365)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19365
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19367 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfail, tmp19366)
}
__typedArg0 := symfail
__typedArg1 := tmp19366
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19368 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19367)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19367
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19369 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfail_1if, tmp19368)
}
__typedArg0 := symfail_1if
__typedArg1 := tmp19368
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19370 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19369)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19369
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19371 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfactorise_2, tmp19370)
}
__typedArg0 := symfactorise_2
__typedArg1 := tmp19370
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19372 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19371)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19371
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19373 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfactorise, tmp19372)
}
__typedArg0 := symfactorise
__typedArg1 := tmp19372
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19374 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19373)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19373
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19375 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symexternal, tmp19374)
}
__typedArg0 := symexternal
__typedArg1 := tmp19374
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19376 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19375)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19375
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19377 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symexplode, tmp19376)
}
__typedArg0 := symexplode
__typedArg1 := tmp19376
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19378 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19377)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19377
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19379 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symeval_1kl, tmp19378)
}
__typedArg0 := symeval_1kl
__typedArg1 := tmp19378
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19380 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19379)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19379
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19381 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symeval, tmp19380)
}
__typedArg0 := symeval
__typedArg1 := tmp19380
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19382 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19381)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19381
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19383 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symerror_1to_1string, tmp19382)
}
__typedArg0 := symerror_1to_1string
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
return PrimCons(symexternal, tmp19384)
}
__typedArg0 := symexternal
__typedArg1 := tmp19384
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19386 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19385)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19385
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19387 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symenable_1type_1theory, tmp19386)
}
__typedArg0 := symenable_1type_1theory
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
return PrimCons(symempty_2, tmp19388)
}
__typedArg0 := symempty_2
__typedArg1 := tmp19388
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19390 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19389)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19389
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19391 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symelement_2, tmp19390)
}
__typedArg0 := symelement_2
__typedArg1 := tmp19390
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19392 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19391)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19391
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19393 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdo, tmp19392)
}
__typedArg0 := symdo
__typedArg1 := tmp19392
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19394 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19393)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19393
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19395 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdifference, tmp19394)
}
__typedArg0 := symdifference
__typedArg1 := tmp19394
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19396 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19395)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19395
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19397 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdestroy, tmp19396)
}
__typedArg0 := symdestroy
__typedArg1 := tmp19396
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19398 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19397)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19397
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19399 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdeclare, tmp19398)
}
__typedArg0 := symdeclare
__typedArg1 := tmp19398
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19400 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19399)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19399
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19401 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdatatypes, tmp19400)
}
__typedArg0 := symdatatypes
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
return PrimCons(symclose, tmp19402)
}
__typedArg0 := symclose
__typedArg1 := tmp19402
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19404 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19403)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19403
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19405 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcn, tmp19404)
}
__typedArg0 := symcn
__typedArg1 := tmp19404
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19406 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19405)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19405
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19407 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons_2, tmp19406)
}
__typedArg0 := symcons_2
__typedArg1 := tmp19406
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19408 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19407)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19407
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19409 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp19408)
}
__typedArg0 := symcons
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
return PrimCons(symconcat, tmp19410)
}
__typedArg0 := symconcat
__typedArg1 := tmp19410
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19412 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19411)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19411
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19413 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcompile, tmp19412)
}
__typedArg0 := symcompile
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
return PrimCons(symcd, tmp19414)
}
__typedArg0 := symcd
__typedArg1 := tmp19414
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19416 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(5), tmp19415)
}
__typedArg0 := MakeInteger(5)
__typedArg1 := tmp19415
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19417 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcall, tmp19416)
}
__typedArg0 := symcall
__typedArg1 := tmp19416
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19418 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(6), tmp19417)
}
__typedArg0 := MakeInteger(6)
__typedArg1 := tmp19417
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19419 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symbind, tmp19418)
}
__typedArg0 := symbind
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
return PrimCons(symbound_2, tmp19420)
}
__typedArg0 := symbound_2
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
return PrimCons(symbootstrap, tmp19422)
}
__typedArg0 := symbootstrap
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
return PrimCons(symboolean_2, tmp19424)
}
__typedArg0 := symboolean_2
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
return PrimCons(symatom_2, tmp19426)
}
__typedArg0 := symatom_2
__typedArg1 := tmp19426
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19428 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19427)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19427
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19429 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symassoc, tmp19428)
}
__typedArg0 := symassoc
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
return PrimCons(symarity, tmp19430)
}
__typedArg0 := symarity
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
return PrimCons(symappend, tmp19432)
}
__typedArg0 := symappend
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
return PrimCons(symand, tmp19434)
}
__typedArg0 := symand
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
return PrimCons(symadjoin, tmp19436)
}
__typedArg0 := symadjoin
__typedArg1 := tmp19436
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19438 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(3), tmp19437)
}
__typedArg0 := MakeInteger(3)
__typedArg1 := tmp19437
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19439 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symaddress_1_6, tmp19438)
}
__typedArg0 := symaddress_1_6
__typedArg1 := tmp19438
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19440 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19439)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19439
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19441 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symabsvector, tmp19440)
}
__typedArg0 := symabsvector
__typedArg1 := tmp19440
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19442 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19441)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19441
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19443 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symabsvector_2, tmp19442)
}
__typedArg0 := symabsvector_2
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
return PrimCons(symabsolute, tmp19444)
}
__typedArg0 := symabsolute
__typedArg1 := tmp19444
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19446 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19445)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19445
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19447 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symabort, tmp19446)
}
__typedArg0 := symabort
__typedArg1 := tmp19446
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19448 := Call(__e, PrimFunc(symshen_4initialise_1arity_1table), tmp19447)


_ = tmp19448

tmp19449 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp19450 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(";"))
}
__typedArg0 := MakeString(";")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp19451 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":="))
}
__typedArg0 := MakeString(":=")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp19452 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(","))
}
__typedArg0 := MakeString(",")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp19453 := Call(__e, PrimFunc(symvector), MakeInteger(0))


tmp19454 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString("bar!"))
}
__typedArg0 := MakeString("bar!")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp19455 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symabort, Nil)
}
__typedArg0 := symabort
__typedArg1 := Nil
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
return PrimCons(symabsvector, tmp19456)
}
__typedArg0 := symabsvector
__typedArg1 := tmp19456
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19458 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symabsvector_2, tmp19457)
}
__typedArg0 := symabsvector_2
__typedArg1 := tmp19457
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19459 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symaddress_1_6, tmp19458)
}
__typedArg0 := symaddress_1_6
__typedArg1 := tmp19458
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19460 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1address, tmp19459)
}
__typedArg0 := sym_5_1address
__typedArg1 := tmp19459
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19461 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symadjoin, tmp19460)
}
__typedArg0 := symadjoin
__typedArg1 := tmp19460
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19462 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symand, tmp19461)
}
__typedArg0 := symand
__typedArg1 := tmp19461
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19463 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symappend, tmp19462)
}
__typedArg0 := symappend
__typedArg1 := tmp19462
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19464 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symarity, tmp19463)
}
__typedArg0 := symarity
__typedArg1 := tmp19463
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19465 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symassoc, tmp19464)
}
__typedArg0 := symassoc
__typedArg1 := tmp19464
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19466 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symassertz, tmp19465)
}
__typedArg0 := symassertz
__typedArg1 := tmp19465
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19467 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symasserta, tmp19466)
}
__typedArg0 := symasserta
__typedArg1 := tmp19466
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19468 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symatom_2, tmp19467)
}
__typedArg0 := symatom_2
__typedArg1 := tmp19467
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19469 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19454, tmp19468)
}
__typedArg0 := tmp19454
__typedArg1 := tmp19468
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19470 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symbootstrap, tmp19469)
}
__typedArg0 := symbootstrap
__typedArg1 := tmp19469
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19471 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, tmp19470)
}
__typedArg0 := symboolean
__typedArg1 := tmp19470
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19472 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean_2, tmp19471)
}
__typedArg0 := symboolean_2
__typedArg1 := tmp19471
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19473 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symbound_2, tmp19472)
}
__typedArg0 := symbound_2
__typedArg1 := tmp19472
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19474 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symbind, tmp19473)
}
__typedArg0 := symbind
__typedArg1 := tmp19473
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19475 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symclose, tmp19474)
}
__typedArg0 := symclose
__typedArg1 := tmp19474
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19476 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcall, tmp19475)
}
__typedArg0 := symcall
__typedArg1 := tmp19475
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19477 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcases, tmp19476)
}
__typedArg0 := symcases
__typedArg1 := tmp19476
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19478 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcd, tmp19477)
}
__typedArg0 := symcd
__typedArg1 := tmp19477
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19479 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcompile, tmp19478)
}
__typedArg0 := symcompile
__typedArg1 := tmp19478
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19480 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symconcat, tmp19479)
}
__typedArg0 := symconcat
__typedArg1 := tmp19479
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19481 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcond, tmp19480)
}
__typedArg0 := symcond
__typedArg1 := tmp19480
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19482 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp19481)
}
__typedArg0 := symcons
__typedArg1 := tmp19481
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19483 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons_2, tmp19482)
}
__typedArg0 := symcons_2
__typedArg1 := tmp19482
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19484 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcn, tmp19483)
}
__typedArg0 := symcn
__typedArg1 := tmp19483
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19485 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4ctxt, tmp19484)
}
__typedArg0 := symshen_4ctxt
__typedArg1 := tmp19484
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19486 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdatatypes, tmp19485)
}
__typedArg0 := symdatatypes
__typedArg1 := tmp19485
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19487 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdatatype, tmp19486)
}
__typedArg0 := symdatatype
__typedArg1 := tmp19486
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19488 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdeclare, tmp19487)
}
__typedArg0 := symdeclare
__typedArg1 := tmp19487
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19489 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefprolog, tmp19488)
}
__typedArg0 := symdefprolog
__typedArg1 := tmp19488
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19490 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefcc, tmp19489)
}
__typedArg0 := symdefcc
__typedArg1 := tmp19489
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19491 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefmacro, tmp19490)
}
__typedArg0 := symdefmacro
__typedArg1 := tmp19490
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19492 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefine, tmp19491)
}
__typedArg0 := symdefine
__typedArg1 := tmp19491
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19493 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefun, tmp19492)
}
__typedArg0 := symdefun
__typedArg1 := tmp19492
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19494 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdestroy, tmp19493)
}
__typedArg0 := symdestroy
__typedArg1 := tmp19493
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19495 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdifference, tmp19494)
}
__typedArg0 := symdifference
__typedArg1 := tmp19494
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19496 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdo, tmp19495)
}
__typedArg0 := symdo
__typedArg1 := tmp19495
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19497 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symelement_2, tmp19496)
}
__typedArg0 := symelement_2
__typedArg1 := tmp19496
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19498 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symempty_2, tmp19497)
}
__typedArg0 := symempty_2
__typedArg1 := tmp19497
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19499 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symerror, tmp19498)
}
__typedArg0 := symerror
__typedArg1 := tmp19498
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19500 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symerror_1to_1string, tmp19499)
}
__typedArg0 := symerror_1to_1string
__typedArg1 := tmp19499
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19501 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symeval, tmp19500)
}
__typedArg0 := symeval
__typedArg1 := tmp19500
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19502 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symeval_1kl, tmp19501)
}
__typedArg0 := symeval_1kl
__typedArg1 := tmp19501
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19503 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symexception, tmp19502)
}
__typedArg0 := symexception
__typedArg1 := tmp19502
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19504 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symexternal, tmp19503)
}
__typedArg0 := symexternal
__typedArg1 := tmp19503
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19505 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symexplode, tmp19504)
}
__typedArg0 := symexplode
__typedArg1 := tmp19504
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19506 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symenable_1type_1theory, tmp19505)
}
__typedArg0 := symenable_1type_1theory
__typedArg1 := tmp19505
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19507 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(False, tmp19506)
}
__typedArg0 := False
__typedArg1 := tmp19506
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19508 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfindall, tmp19507)
}
__typedArg0 := symfindall
__typedArg1 := tmp19507
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19509 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfactorise, tmp19508)
}
__typedArg0 := symfactorise
__typedArg1 := tmp19508
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19510 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfail_1if, tmp19509)
}
__typedArg0 := symfail_1if
__typedArg1 := tmp19509
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19511 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfail, tmp19510)
}
__typedArg0 := symfail
__typedArg1 := tmp19510
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19512 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfile, tmp19511)
}
__typedArg0 := symfile
__typedArg1 := tmp19511
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19513 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfix, tmp19512)
}
__typedArg0 := symfix
__typedArg1 := tmp19512
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19514 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symforeign, tmp19513)
}
__typedArg0 := symforeign
__typedArg1 := tmp19513
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19515 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfork, tmp19514)
}
__typedArg0 := symfork
__typedArg1 := tmp19514
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19516 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfresh, tmp19515)
}
__typedArg0 := symfresh
__typedArg1 := tmp19515
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19517 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfreeze, tmp19516)
}
__typedArg0 := symfreeze
__typedArg1 := tmp19516
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19518 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfst, tmp19517)
}
__typedArg0 := symfst
__typedArg1 := tmp19517
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19519 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfunction, tmp19518)
}
__typedArg0 := symfunction
__typedArg1 := tmp19518
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19520 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfn, tmp19519)
}
__typedArg0 := symfn
__typedArg1 := tmp19519
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19521 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symgensym, tmp19520)
}
__typedArg0 := symgensym
__typedArg1 := tmp19520
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19522 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symget_1time, tmp19521)
}
__typedArg0 := symget_1time
__typedArg1 := tmp19521
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19523 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symget, tmp19522)
}
__typedArg0 := symget
__typedArg1 := tmp19522
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19524 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhash, tmp19523)
}
__typedArg0 := symhash
__typedArg1 := tmp19523
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19525 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhdstr, tmp19524)
}
__typedArg0 := symhdstr
__typedArg1 := tmp19524
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19526 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhdv, tmp19525)
}
__typedArg0 := symhdv
__typedArg1 := tmp19525
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19527 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhd, tmp19526)
}
__typedArg0 := symhd
__typedArg1 := tmp19526
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19528 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhead, tmp19527)
}
__typedArg0 := symhead
__typedArg1 := tmp19527
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19529 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp19528)
}
__typedArg0 := symif
__typedArg1 := tmp19528
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19530 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symimplementation, tmp19529)
}
__typedArg0 := symimplementation
__typedArg1 := tmp19529
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19531 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminternal, tmp19530)
}
__typedArg0 := syminternal
__typedArg1 := tmp19530
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19532 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symin_1package, tmp19531)
}
__typedArg0 := symin_1package
__typedArg1 := tmp19531
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19533 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symin, tmp19532)
}
__typedArg0 := symin
__typedArg1 := tmp19532
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19534 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symis_b, tmp19533)
}
__typedArg0 := symis_b
__typedArg1 := tmp19533
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19535 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symis, tmp19534)
}
__typedArg0 := symis
__typedArg1 := tmp19534
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19536 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symit, tmp19535)
}
__typedArg0 := symit
__typedArg1 := tmp19535
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19537 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminclude_1all_1but, tmp19536)
}
__typedArg0 := syminclude_1all_1but
__typedArg1 := tmp19536
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19538 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminclude, tmp19537)
}
__typedArg0 := syminclude
__typedArg1 := tmp19537
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19539 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminline, tmp19538)
}
__typedArg0 := syminline
__typedArg1 := tmp19538
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19540 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminput_7, tmp19539)
}
__typedArg0 := syminput_7
__typedArg1 := tmp19539
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19541 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminput, tmp19540)
}
__typedArg0 := syminput
__typedArg1 := tmp19540
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19542 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminteger_2, tmp19541)
}
__typedArg0 := syminteger_2
__typedArg1 := tmp19541
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19543 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symintern, tmp19542)
}
__typedArg0 := symintern
__typedArg1 := tmp19542
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19544 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminferences, tmp19543)
}
__typedArg0 := syminferences
__typedArg1 := tmp19543
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19545 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symintersection, tmp19544)
}
__typedArg0 := symintersection
__typedArg1 := tmp19544
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19546 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symis, tmp19545)
}
__typedArg0 := symis
__typedArg1 := tmp19545
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19547 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlanguage, tmp19546)
}
__typedArg0 := symlanguage
__typedArg1 := tmp19546
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19548 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp19547)
}
__typedArg0 := symlambda
__typedArg1 := tmp19547
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19549 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlazy, tmp19548)
}
__typedArg0 := symlazy
__typedArg1 := tmp19548
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19550 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp19549)
}
__typedArg0 := symlet
__typedArg1 := tmp19549
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19551 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlength, tmp19550)
}
__typedArg0 := symlength
__typedArg1 := tmp19550
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19552 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlimit, tmp19551)
}
__typedArg0 := symlimit
__typedArg1 := tmp19551
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19553 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlineread, tmp19552)
}
__typedArg0 := symlineread
__typedArg1 := tmp19552
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19554 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19553)
}
__typedArg0 := symlist
__typedArg1 := tmp19553
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19555 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symloaded, tmp19554)
}
__typedArg0 := symloaded
__typedArg1 := tmp19554
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19556 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symload, tmp19555)
}
__typedArg0 := symload
__typedArg1 := tmp19555
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19557 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symmake_1string, tmp19556)
}
__typedArg0 := symmake_1string
__typedArg1 := tmp19556
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19558 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symmap, tmp19557)
}
__typedArg0 := symmap
__typedArg1 := tmp19557
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19559 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symmapcan, tmp19558)
}
__typedArg0 := symmapcan
__typedArg1 := tmp19558
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19560 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symmaxinferences, tmp19559)
}
__typedArg0 := symmaxinferences
__typedArg1 := tmp19559
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19561 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symmacroexpand, tmp19560)
}
__typedArg0 := symmacroexpand
__typedArg1 := tmp19560
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19562 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symmode, tmp19561)
}
__typedArg0 := symmode
__typedArg1 := tmp19561
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19563 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnl, tmp19562)
}
__typedArg0 := symnl
__typedArg1 := tmp19562
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19564 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnot, tmp19563)
}
__typedArg0 := symnot
__typedArg1 := tmp19563
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19565 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnth, tmp19564)
}
__typedArg0 := symnth
__typedArg1 := tmp19564
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19566 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnull, tmp19565)
}
__typedArg0 := symnull
__typedArg1 := tmp19565
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19567 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp19566)
}
__typedArg0 := symnumber
__typedArg1 := tmp19566
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19568 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber_2, tmp19567)
}
__typedArg0 := symnumber_2
__typedArg1 := tmp19567
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19569 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symn_1_6string, tmp19568)
}
__typedArg0 := symn_1_6string
__typedArg1 := tmp19568
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19570 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symoccurs_1check, tmp19569)
}
__typedArg0 := symoccurs_1check
__typedArg1 := tmp19569
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19571 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symoccurrences, tmp19570)
}
__typedArg0 := symoccurrences
__typedArg1 := tmp19570
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19572 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symopen, tmp19571)
}
__typedArg0 := symopen
__typedArg1 := tmp19571
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19573 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symoptimise, tmp19572)
}
__typedArg0 := symoptimise
__typedArg1 := tmp19572
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19574 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symor, tmp19573)
}
__typedArg0 := symor
__typedArg1 := tmp19573
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19575 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symos, tmp19574)
}
__typedArg0 := symos
__typedArg1 := tmp19574
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19576 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symout, tmp19575)
}
__typedArg0 := symout
__typedArg1 := tmp19575
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19577 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symoutput, tmp19576)
}
__typedArg0 := symoutput
__typedArg1 := tmp19576
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19578 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympackage, tmp19577)
}
__typedArg0 := sympackage
__typedArg1 := tmp19577
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19579 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symport, tmp19578)
}
__typedArg0 := symport
__typedArg1 := tmp19578
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19580 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symporters, tmp19579)
}
__typedArg0 := symporters
__typedArg1 := tmp19579
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19581 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympos, tmp19580)
}
__typedArg0 := sympos
__typedArg1 := tmp19580
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19582 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympr, tmp19581)
}
__typedArg0 := sympr
__typedArg1 := tmp19581
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19583 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symprint, tmp19582)
}
__typedArg0 := symprint
__typedArg1 := tmp19582
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19584 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symprolog_1memory, tmp19583)
}
__typedArg0 := symprolog_1memory
__typedArg1 := tmp19583
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19585 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symprofile, tmp19584)
}
__typedArg0 := symprofile
__typedArg1 := tmp19584
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19586 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symprofile_1results, tmp19585)
}
__typedArg0 := symprofile_1results
__typedArg1 := tmp19585
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19587 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symprotect, tmp19586)
}
__typedArg0 := symprotect
__typedArg1 := tmp19586
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19588 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symprolog_2, tmp19587)
}
__typedArg0 := symprolog_2
__typedArg1 := tmp19587
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19589 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symps, tmp19588)
}
__typedArg0 := symps
__typedArg1 := tmp19588
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19590 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympreclude_1all_1but, tmp19589)
}
__typedArg0 := sympreclude_1all_1but
__typedArg1 := tmp19589
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19591 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympreclude, tmp19590)
}
__typedArg0 := sympreclude
__typedArg1 := tmp19590
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19592 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symput, tmp19591)
}
__typedArg0 := symput
__typedArg1 := tmp19591
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19593 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympackage_2, tmp19592)
}
__typedArg0 := sympackage_2
__typedArg1 := tmp19592
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19594 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1from_1string_1unprocessed, tmp19593)
}
__typedArg0 := symread_1from_1string_1unprocessed
__typedArg1 := tmp19593
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19595 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1from_1string, tmp19594)
}
__typedArg0 := symread_1from_1string
__typedArg1 := tmp19594
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19596 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1byte, tmp19595)
}
__typedArg0 := symread_1byte
__typedArg1 := tmp19595
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19597 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1file_1as_1string, tmp19596)
}
__typedArg0 := symread_1file_1as_1string
__typedArg1 := tmp19596
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19598 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1file_1as_1bytelist, tmp19597)
}
__typedArg0 := symread_1file_1as_1bytelist
__typedArg1 := tmp19597
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19599 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1file, tmp19598)
}
__typedArg0 := symread_1file
__typedArg1 := tmp19598
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19600 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symreceive, tmp19599)
}
__typedArg0 := symreceive
__typedArg1 := tmp19599
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19601 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread, tmp19600)
}
__typedArg0 := symread
__typedArg1 := tmp19600
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19602 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symrelease, tmp19601)
}
__typedArg0 := symrelease
__typedArg1 := tmp19601
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19603 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symremove, tmp19602)
}
__typedArg0 := symremove
__typedArg1 := tmp19602
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19604 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symretract, tmp19603)
}
__typedArg0 := symretract
__typedArg1 := tmp19603
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19605 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symreverse, tmp19604)
}
__typedArg0 := symreverse
__typedArg1 := tmp19604
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19606 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symrun, tmp19605)
}
__typedArg0 := symrun
__typedArg1 := tmp19605
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19607 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstr, tmp19606)
}
__typedArg0 := symstr
__typedArg1 := tmp19606
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19608 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsave, tmp19607)
}
__typedArg0 := symsave
__typedArg1 := tmp19607
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19609 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symset, tmp19608)
}
__typedArg0 := symset
__typedArg1 := tmp19608
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19610 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsimple_1error, tmp19609)
}
__typedArg0 := symsimple_1error
__typedArg1 := tmp19609
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19611 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsnd, tmp19610)
}
__typedArg0 := symsnd
__typedArg1 := tmp19610
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19612 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symspecialise, tmp19611)
}
__typedArg0 := symspecialise
__typedArg1 := tmp19611
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19613 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symspy, tmp19612)
}
__typedArg0 := symspy
__typedArg1 := tmp19612
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19614 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsqts, tmp19613)
}
__typedArg0 := symsqts
__typedArg1 := tmp19613
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19615 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstep, tmp19614)
}
__typedArg0 := symstep
__typedArg1 := tmp19614
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19616 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstoutput, tmp19615)
}
__typedArg0 := symstoutput
__typedArg1 := tmp19615
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19617 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsterror, tmp19616)
}
__typedArg0 := symsterror
__typedArg1 := tmp19616
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19618 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstinput, tmp19617)
}
__typedArg0 := symstinput
__typedArg1 := tmp19617
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19619 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp19618)
}
__typedArg0 := symstring
__typedArg1 := tmp19618
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19620 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp19619)
}
__typedArg0 := symstream
__typedArg1 := tmp19619
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19621 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring_1_6n, tmp19620)
}
__typedArg0 := symstring_1_6n
__typedArg1 := tmp19620
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19622 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring_2, tmp19621)
}
__typedArg0 := symstring_2
__typedArg1 := tmp19621
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19623 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsubst, tmp19622)
}
__typedArg0 := symsubst
__typedArg1 := tmp19622
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19624 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsum, tmp19623)
}
__typedArg0 := symsum
__typedArg1 := tmp19623
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19625 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring_1_6symbol, tmp19624)
}
__typedArg0 := symstring_1_6symbol
__typedArg1 := tmp19624
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19626 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol_2, tmp19625)
}
__typedArg0 := symsymbol_2
__typedArg1 := tmp19625
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19627 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp19626)
}
__typedArg0 := symsymbol
__typedArg1 := tmp19626
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19628 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsynonyms, tmp19627)
}
__typedArg0 := symsynonyms
__typedArg1 := tmp19627
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19629 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsystemf, tmp19628)
}
__typedArg0 := symsystemf
__typedArg1 := tmp19628
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19630 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtail, tmp19629)
}
__typedArg0 := symtail
__typedArg1 := tmp19629
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19631 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtlv, tmp19630)
}
__typedArg0 := symtlv
__typedArg1 := tmp19630
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19632 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtlstr, tmp19631)
}
__typedArg0 := symtlstr
__typedArg1 := tmp19631
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19633 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtl, tmp19632)
}
__typedArg0 := symtl
__typedArg1 := tmp19632
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19634 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtc, tmp19633)
}
__typedArg0 := symtc
__typedArg1 := tmp19633
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19635 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtc_2, tmp19634)
}
__typedArg0 := symtc_2
__typedArg1 := tmp19634
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19636 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symthaw, tmp19635)
}
__typedArg0 := symthaw
__typedArg1 := tmp19635
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19637 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtime, tmp19636)
}
__typedArg0 := symtime
__typedArg1 := tmp19636
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19638 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtrack, tmp19637)
}
__typedArg0 := symtrack
__typedArg1 := tmp19637
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19639 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtrap_1error, tmp19638)
}
__typedArg0 := symtrap_1error
__typedArg1 := tmp19638
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19640 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(True, tmp19639)
}
__typedArg0 := True
__typedArg1 := tmp19639
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19641 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtuple_2, tmp19640)
}
__typedArg0 := symtuple_2
__typedArg1 := tmp19640
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19642 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtype, tmp19641)
}
__typedArg0 := symtype
__typedArg1 := tmp19641
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19643 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symreturn, tmp19642)
}
__typedArg0 := symreturn
__typedArg1 := tmp19642
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19644 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunabsolute, tmp19643)
}
__typedArg0 := symunabsolute
__typedArg1 := tmp19643
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19645 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symundefmacro, tmp19644)
}
__typedArg0 := symundefmacro
__typedArg1 := tmp19644
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19646 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunprofile, tmp19645)
}
__typedArg0 := symunprofile
__typedArg1 := tmp19645
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19647 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunput, tmp19646)
}
__typedArg0 := symunput
__typedArg1 := tmp19646
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19648 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunion, tmp19647)
}
__typedArg0 := symunion
__typedArg1 := tmp19647
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19649 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4unix, tmp19648)
}
__typedArg0 := symshen_4unix
__typedArg1 := tmp19648
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19650 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunit, tmp19649)
}
__typedArg0 := symunit
__typedArg1 := tmp19649
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19651 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symuntrack, tmp19650)
}
__typedArg0 := symuntrack
__typedArg1 := tmp19650
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19652 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunspecialise, tmp19651)
}
__typedArg0 := symunspecialise
__typedArg1 := tmp19651
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19653 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symupdate_1lambda_1table, tmp19652)
}
__typedArg0 := symupdate_1lambda_1table
__typedArg1 := tmp19652
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19654 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symu_b, tmp19653)
}
__typedArg0 := symu_b
__typedArg1 := tmp19653
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19655 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector_2, tmp19654)
}
__typedArg0 := symvector_2
__typedArg1 := tmp19654
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19656 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp19655)
}
__typedArg0 := symvector
__typedArg1 := tmp19655
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19657 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1vector, tmp19656)
}
__typedArg0 := sym_5_1vector
__typedArg1 := tmp19656
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19658 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector_1_6, tmp19657)
}
__typedArg0 := symvector_1_6
__typedArg1 := tmp19657
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19659 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvalue, tmp19658)
}
__typedArg0 := symvalue
__typedArg1 := tmp19658
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19660 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvar_2, tmp19659)
}
__typedArg0 := symvar_2
__typedArg1 := tmp19659
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19661 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvariable_2, tmp19660)
}
__typedArg0 := symvariable_2
__typedArg1 := tmp19660
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19662 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symverified, tmp19661)
}
__typedArg0 := symverified
__typedArg1 := tmp19661
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19663 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symversion, tmp19662)
}
__typedArg0 := symversion
__typedArg1 := tmp19662
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19664 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwarn, tmp19663)
}
__typedArg0 := symwarn
__typedArg1 := tmp19663
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19665 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhen, tmp19664)
}
__typedArg0 := symwhen
__typedArg1 := tmp19664
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19666 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhere, tmp19665)
}
__typedArg0 := symwhere
__typedArg1 := tmp19665
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19667 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwrite_1byte, tmp19666)
}
__typedArg0 := symwrite_1byte
__typedArg1 := tmp19666
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19668 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwrite_1to_1file, tmp19667)
}
__typedArg0 := symwrite_1to_1file
__typedArg1 := tmp19667
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19669 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symy_1or_1n_2, tmp19668)
}
__typedArg0 := symy_1or_1n_2
__typedArg1 := tmp19668
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19670 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19453, tmp19669)
}
__typedArg0 := tmp19453
__typedArg1 := tmp19669
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19671 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_6_6, tmp19670)
}
__typedArg0 := sym_6_6
__typedArg1 := tmp19670
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19672 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5, tmp19671)
}
__typedArg0 := sym_5
__typedArg1 := tmp19671
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19673 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_a, tmp19672)
}
__typedArg0 := sym_5_a
__typedArg1 := tmp19672
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19674 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_7, tmp19673)
}
__typedArg0 := sym_7
__typedArg1 := tmp19673
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19675 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_d, tmp19674)
}
__typedArg0 := sym_d
__typedArg1 := tmp19674
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19676 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_c, tmp19675)
}
__typedArg0 := sym_c
__typedArg1 := tmp19675
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19677 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1, tmp19676)
}
__typedArg0 := sym_1
__typedArg1 := tmp19676
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19678 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_3, tmp19677)
}
__typedArg0 := sym_3
__typedArg1 := tmp19677
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19679 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5end_6, tmp19678)
}
__typedArg0 := sym_5end_6
__typedArg1 := tmp19678
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19680 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_b_6, tmp19679)
}
__typedArg0 := sym_5_b_6
__typedArg1 := tmp19679
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19681 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_c_4, tmp19680)
}
__typedArg0 := sym_c_4
__typedArg1 := tmp19680
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19682 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_a_a_6, tmp19681)
}
__typedArg0 := sym_a_a_6
__typedArg1 := tmp19681
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19683 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_6, tmp19682)
}
__typedArg0 := sym_6
__typedArg1 := tmp19682
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19684 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_6_a, tmp19683)
}
__typedArg0 := sym_6_a
__typedArg1 := tmp19683
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19685 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_a, tmp19684)
}
__typedArg0 := sym_a
__typedArg1 := tmp19684
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19686 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_a_a, tmp19685)
}
__typedArg0 := sym_a_a
__typedArg1 := tmp19685
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19687 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5e_6, tmp19686)
}
__typedArg0 := sym_5e_6
__typedArg1 := tmp19686
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19688 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_6, tmp19687)
}
__typedArg0 := sym_1_6
__typedArg1 := tmp19687
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19689 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1, tmp19688)
}
__typedArg0 := sym_5_1
__typedArg1 := tmp19688
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19690 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dhush_d, tmp19689)
}
__typedArg0 := sym_dhush_d
__typedArg1 := tmp19689
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19691 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dporters_d, tmp19690)
}
__typedArg0 := sym_dporters_d
__typedArg1 := tmp19690
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19692 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dport_d, tmp19691)
}
__typedArg0 := sym_dport_d
__typedArg1 := tmp19691
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19693 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8s, tmp19692)
}
__typedArg0 := sym_8s
__typedArg1 := tmp19692
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19694 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8p, tmp19693)
}
__typedArg0 := sym_8p
__typedArg1 := tmp19693
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19695 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8v, tmp19694)
}
__typedArg0 := sym_8v
__typedArg1 := tmp19694
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19696 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dproperty_1vector_d, tmp19695)
}
__typedArg0 := sym_dproperty_1vector_d
__typedArg1 := tmp19695
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19697 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_drelease_d, tmp19696)
}
__typedArg0 := sym_drelease_d
__typedArg1 := tmp19696
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19698 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dos_d, tmp19697)
}
__typedArg0 := sym_dos_d
__typedArg1 := tmp19697
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19699 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dmacros_d, tmp19698)
}
__typedArg0 := sym_dmacros_d
__typedArg1 := tmp19698
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19700 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dmaximum_1print_1sequence_1size_d, tmp19699)
}
__typedArg0 := sym_dmaximum_1print_1sequence_1size_d
__typedArg1 := tmp19699
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19701 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dversion_d, tmp19700)
}
__typedArg0 := sym_dversion_d
__typedArg1 := tmp19700
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19702 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dhome_1directory_d, tmp19701)
}
__typedArg0 := sym_dhome_1directory_d
__typedArg1 := tmp19701
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19703 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dstoutput_d, tmp19702)
}
__typedArg0 := sym_dstoutput_d
__typedArg1 := tmp19702
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19704 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dsterror_d, tmp19703)
}
__typedArg0 := sym_dsterror_d
__typedArg1 := tmp19703
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19705 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dstinput_d, tmp19704)
}
__typedArg0 := sym_dstinput_d
__typedArg1 := tmp19704
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19706 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dimplementation_d, tmp19705)
}
__typedArg0 := sym_dimplementation_d
__typedArg1 := tmp19705
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19707 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dlanguage_d, tmp19706)
}
__typedArg0 := sym_dlanguage_d
__typedArg1 := tmp19706
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19708 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym__, tmp19707)
}
__typedArg0 := sym__
__typedArg1 := tmp19707
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19709 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19452, tmp19708)
}
__typedArg0 := tmp19452
__typedArg1 := tmp19708
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19710 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19451, tmp19709)
}
__typedArg0 := tmp19451
__typedArg1 := tmp19709
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19711 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19450, tmp19710)
}
__typedArg0 := tmp19450
__typedArg1 := tmp19710
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19712 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19449, tmp19711)
}
__typedArg0 := tmp19449
__typedArg1 := tmp19711
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19713 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_e_e, tmp19712)
}
__typedArg0 := sym_e_e
__typedArg1 := tmp19712
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19714 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1_1, tmp19713)
}
__typedArg0 := sym_5_1_1
__typedArg1 := tmp19713
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19715 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19714)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19714
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19716 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_i, tmp19715)
}
__typedArg0 := sym_i
__typedArg1 := tmp19715
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19717 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_j, tmp19716)
}
__typedArg0 := sym_j
__typedArg1 := tmp19716
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19718 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_b, tmp19717)
}
__typedArg0 := sym_b
__typedArg1 := tmp19717
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19719 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp19720 := Call(__e, PrimFunc(symput), symshen, symshen_4external_1symbols, tmp19718, tmp19719)


_ = tmp19720

tmp19721 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symabsvector) {
return PrimAbsvector(MakeInteger(0))
}
__typedArg0 := MakeInteger(0)
return Call(__e, PrimFunc(symabsvector), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dempty_1absvector_d, tmp19721)
}
__typedArg0 := symshen_4_dempty_1absvector_d
__typedArg1 := tmp19721
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return


}, 0)

tmp19722 := Call(__e, ns2_1set, symshen_4initialise_1environment, tmp18975)


_ = tmp19722

tmp19723 := MakeNative(func(__e *ControlFlow) {
tmp19724 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, Nil)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19724

tmp19725 := MakeNative(func(__e *ControlFlow) {
V6066 := __e.Get(1)
_ = V6066
__e.Return(MakeNative(func(__e *ControlFlow) {
B6062 := __e.Get(1)
_ = B6062
__e.Return(MakeNative(func(__e *ControlFlow) {
L6063 := __e.Get(1)
_ = L6063
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6064 := __e.Get(1)
_ = Key6064
__e.Return(MakeNative(func(__e *ControlFlow) {
C6065 := __e.Get(1)
_ = C6065
tmp19726 := Call(__e, PrimFunc(symshen_4newpv), B6062)


let__18889 := tmp19726
_ = let__18889

tmp19727 := Call(__e, PrimFunc(symshen_4newpv), B6062)


let__18890 := tmp19727
_ = let__18890

tmp19728 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18890, Nil)
}
__typedArg0 := let__18890
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19729 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19728)
}
__typedArg0 := symlist
__typedArg1 := tmp19728
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19730 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19729, Nil)
}
__typedArg0 := tmp19729
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19731 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19730)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19730
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19732 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18889, tmp19731)
}
__typedArg0 := let__18889
__typedArg1 := tmp19731
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19733 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19734 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19733)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19733
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19735 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19732, tmp19734)
}
__typedArg0 := tmp19732
__typedArg1 := tmp19734
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19736 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19735, Nil)
}
__typedArg0 := tmp19735
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19737 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19736)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19736
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19738 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp19737)
}
__typedArg0 := symsymbol
__typedArg1 := tmp19737
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19739 := Call(__e, PrimFunc(symis_b), V6066, tmp19738, B6062, L6063, Key6064, C6065)


tmp19740 := Call(__e, PrimFunc(symshen_4gc), B6062, tmp19739)


__e.TailApply(PrimFunc(symshen_4gc), B6062, tmp19740)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19741 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp19742 := Call(__e, PrimFunc(symshen_4assoc_1_6), symshen_4register_1source_1form, tmp19725, tmp19741)


tmp19743 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp19742)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp19742
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19743

tmp19744 := MakeNative(func(__e *ControlFlow) {
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
tmp19745 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19746 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19745)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19745
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19747 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp19746)
}
__typedArg0 := symsymbol
__typedArg1 := tmp19746
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6071, tmp19747, B6067, L6068, Key6069, C6070)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19748 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp19749 := Call(__e, PrimFunc(symshen_4assoc_1_6), symshen_4unregister_1source_1form, tmp19744, tmp19748)


tmp19750 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp19749)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp19749
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19750

tmp19751 := MakeNative(func(__e *ControlFlow) {
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
tmp19752 := Call(__e, PrimFunc(symshen_4newpv), B6072)


let__18891 := tmp19752
_ = let__18891

tmp19753 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18891, Nil)
}
__typedArg0 := let__18891
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19754 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19753)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19753
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19755 := Call(__e, PrimFunc(symis_b), V6076, tmp19754, B6072, L6073, Key6074, C6075)


__e.TailApply(PrimFunc(symshen_4gc), B6072, tmp19755)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19756 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp19757 := Call(__e, PrimFunc(symshen_4assoc_1_6), symabort, tmp19751, tmp19756)


tmp19758 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp19757)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp19757
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19758

tmp19759 := MakeNative(func(__e *ControlFlow) {
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
tmp19760 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19761 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19760)
}
__typedArg0 := symlist
__typedArg1 := tmp19760
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19762 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19761, Nil)
}
__typedArg0 := tmp19761
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19763 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19762)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19762
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19764 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp19763)
}
__typedArg0 := symstring
__typedArg1 := tmp19763
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6081, tmp19764, B6077, L6078, Key6079, C6080)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19765 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp19766 := Call(__e, PrimFunc(symshen_4assoc_1_6), symabsolute, tmp19759, tmp19765)


tmp19767 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp19766)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp19766
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19767

tmp19768 := MakeNative(func(__e *ControlFlow) {
V6086 := __e.Get(1)
_ = V6086
__e.Return(MakeNative(func(__e *ControlFlow) {
B6082 := __e.Get(1)
_ = B6082
__e.Return(MakeNative(func(__e *ControlFlow) {
L6083 := __e.Get(1)
_ = L6083
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6084 := __e.Get(1)
_ = Key6084
__e.Return(MakeNative(func(__e *ControlFlow) {
C6085 := __e.Get(1)
_ = C6085
tmp19769 := Call(__e, PrimFunc(symshen_4newpv), B6082)


let__18892 := tmp19769
_ = let__18892

tmp19770 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
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
return PrimCons(let__18892, tmp19771)
}
__typedArg0 := let__18892
__typedArg1 := tmp19771
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19773 := Call(__e, PrimFunc(symis_b), V6086, tmp19772, B6082, L6083, Key6084, C6085)


__e.TailApply(PrimFunc(symshen_4gc), B6082, tmp19773)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19774 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp19775 := Call(__e, PrimFunc(symshen_4assoc_1_6), symabsvector_2, tmp19768, tmp19774)


tmp19776 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp19775)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp19775
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19776

tmp19777 := MakeNative(func(__e *ControlFlow) {
V6091 := __e.Get(1)
_ = V6091
__e.Return(MakeNative(func(__e *ControlFlow) {
B6087 := __e.Get(1)
_ = B6087
__e.Return(MakeNative(func(__e *ControlFlow) {
L6088 := __e.Get(1)
_ = L6088
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6089 := __e.Get(1)
_ = Key6089
__e.Return(MakeNative(func(__e *ControlFlow) {
C6090 := __e.Get(1)
_ = C6090
tmp19778 := Call(__e, PrimFunc(symshen_4newpv), B6087)


let__18893 := tmp19778
_ = let__18893

tmp19779 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18893, Nil)
}
__typedArg0 := let__18893
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19780 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19779)
}
__typedArg0 := symlist
__typedArg1 := tmp19779
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19781 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18893, Nil)
}
__typedArg0 := let__18893
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19782 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19781)
}
__typedArg0 := symlist
__typedArg1 := tmp19781
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19783 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19782, Nil)
}
__typedArg0 := tmp19782
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19784 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19783)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19783
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19785 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19780, tmp19784)
}
__typedArg0 := tmp19780
__typedArg1 := tmp19784
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19786 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19785, Nil)
}
__typedArg0 := tmp19785
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19787 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19786)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19786
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19788 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18893, tmp19787)
}
__typedArg0 := let__18893
__typedArg1 := tmp19787
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19789 := Call(__e, PrimFunc(symis_b), V6091, tmp19788, B6087, L6088, Key6089, C6090)


__e.TailApply(PrimFunc(symshen_4gc), B6087, tmp19789)
return


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

tmp19791 := Call(__e, PrimFunc(symshen_4assoc_1_6), symadjoin, tmp19777, tmp19790)


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
V6096 := __e.Get(1)
_ = V6096
__e.Return(MakeNative(func(__e *ControlFlow) {
B6092 := __e.Get(1)
_ = B6092
__e.Return(MakeNative(func(__e *ControlFlow) {
L6093 := __e.Get(1)
_ = L6093
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6094 := __e.Get(1)
_ = Key6094
__e.Return(MakeNative(func(__e *ControlFlow) {
C6095 := __e.Get(1)
_ = C6095
tmp19794 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
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
return PrimCons(symboolean, tmp19795)
}
__typedArg0 := symboolean
__typedArg1 := tmp19795
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19797 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19796, Nil)
}
__typedArg0 := tmp19796
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19798 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19797)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19797
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19799 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, tmp19798)
}
__typedArg0 := symboolean
__typedArg1 := tmp19798
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6096, tmp19799, B6092, L6093, Key6094, C6095)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19800 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp19801 := Call(__e, PrimFunc(symshen_4assoc_1_6), symand, tmp19793, tmp19800)


tmp19802 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp19801)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp19801
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19802

tmp19803 := MakeNative(func(__e *ControlFlow) {
V6101 := __e.Get(1)
_ = V6101
__e.Return(MakeNative(func(__e *ControlFlow) {
B6097 := __e.Get(1)
_ = B6097
__e.Return(MakeNative(func(__e *ControlFlow) {
L6098 := __e.Get(1)
_ = L6098
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6099 := __e.Get(1)
_ = Key6099
__e.Return(MakeNative(func(__e *ControlFlow) {
C6100 := __e.Get(1)
_ = C6100
tmp19804 := Call(__e, PrimFunc(symshen_4newpv), B6097)


let__18894 := tmp19804
_ = let__18894

tmp19805 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19806 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19805)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19805
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19807 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp19806)
}
__typedArg0 := symsymbol
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
return PrimCons(symstring, tmp19809)
}
__typedArg0 := symstring
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
return PrimCons(let__18894, tmp19812)
}
__typedArg0 := let__18894
__typedArg1 := tmp19812
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19814 := Call(__e, PrimFunc(symis_b), V6101, tmp19813, B6097, L6098, Key6099, C6100)


__e.TailApply(PrimFunc(symshen_4gc), B6097, tmp19814)
return


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

tmp19816 := Call(__e, PrimFunc(symshen_4assoc_1_6), symshen_4app, tmp19803, tmp19815)


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
V6106 := __e.Get(1)
_ = V6106
__e.Return(MakeNative(func(__e *ControlFlow) {
B6102 := __e.Get(1)
_ = B6102
__e.Return(MakeNative(func(__e *ControlFlow) {
L6103 := __e.Get(1)
_ = L6103
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6104 := __e.Get(1)
_ = Key6104
__e.Return(MakeNative(func(__e *ControlFlow) {
C6105 := __e.Get(1)
_ = C6105
tmp19819 := Call(__e, PrimFunc(symshen_4newpv), B6102)


let__18895 := tmp19819
_ = let__18895

tmp19820 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18895, Nil)
}
__typedArg0 := let__18895
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19821 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19820)
}
__typedArg0 := symlist
__typedArg1 := tmp19820
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19822 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18895, Nil)
}
__typedArg0 := let__18895
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19823 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19822)
}
__typedArg0 := symlist
__typedArg1 := tmp19822
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19824 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18895, Nil)
}
__typedArg0 := let__18895
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19825 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19824)
}
__typedArg0 := symlist
__typedArg1 := tmp19824
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19826 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19825, Nil)
}
__typedArg0 := tmp19825
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19827 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19826)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19826
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19828 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19823, tmp19827)
}
__typedArg0 := tmp19823
__typedArg1 := tmp19827
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19829 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19828, Nil)
}
__typedArg0 := tmp19828
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19830 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19829)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19829
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19831 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19821, tmp19830)
}
__typedArg0 := tmp19821
__typedArg1 := tmp19830
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19832 := Call(__e, PrimFunc(symis_b), V6106, tmp19831, B6102, L6103, Key6104, C6105)


__e.TailApply(PrimFunc(symshen_4gc), B6102, tmp19832)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19833 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp19834 := Call(__e, PrimFunc(symshen_4assoc_1_6), symappend, tmp19818, tmp19833)


tmp19835 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp19834)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp19834
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19835

tmp19836 := MakeNative(func(__e *ControlFlow) {
V6111 := __e.Get(1)
_ = V6111
__e.Return(MakeNative(func(__e *ControlFlow) {
B6107 := __e.Get(1)
_ = B6107
__e.Return(MakeNative(func(__e *ControlFlow) {
L6108 := __e.Get(1)
_ = L6108
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6109 := __e.Get(1)
_ = Key6109
__e.Return(MakeNative(func(__e *ControlFlow) {
C6110 := __e.Get(1)
_ = C6110
tmp19837 := Call(__e, PrimFunc(symshen_4newpv), B6107)


let__18896 := tmp19837
_ = let__18896

tmp19838 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19839 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19838)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19838
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19840 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18896, tmp19839)
}
__typedArg0 := let__18896
__typedArg1 := tmp19839
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19841 := Call(__e, PrimFunc(symis_b), V6111, tmp19840, B6107, L6108, Key6109, C6110)


__e.TailApply(PrimFunc(symshen_4gc), B6107, tmp19841)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19842 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp19843 := Call(__e, PrimFunc(symshen_4assoc_1_6), symarity, tmp19836, tmp19842)


tmp19844 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp19843)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp19843
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19844

tmp19845 := MakeNative(func(__e *ControlFlow) {
V6116 := __e.Get(1)
_ = V6116
__e.Return(MakeNative(func(__e *ControlFlow) {
B6112 := __e.Get(1)
_ = B6112
__e.Return(MakeNative(func(__e *ControlFlow) {
L6113 := __e.Get(1)
_ = L6113
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6114 := __e.Get(1)
_ = Key6114
__e.Return(MakeNative(func(__e *ControlFlow) {
C6115 := __e.Get(1)
_ = C6115
tmp19846 := Call(__e, PrimFunc(symshen_4newpv), B6112)


let__18897 := tmp19846
_ = let__18897

tmp19847 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18897, Nil)
}
__typedArg0 := let__18897
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
return PrimCons(tmp19848, Nil)
}
__typedArg0 := tmp19848
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
return PrimCons(let__18897, Nil)
}
__typedArg0 := let__18897
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19852 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19851)
}
__typedArg0 := symlist
__typedArg1 := tmp19851
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19853 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19852, Nil)
}
__typedArg0 := tmp19852
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19854 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19853)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19853
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19855 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19850, tmp19854)
}
__typedArg0 := tmp19850
__typedArg1 := tmp19854
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19856 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19855, Nil)
}
__typedArg0 := tmp19855
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19857 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19856)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19856
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19858 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18897, tmp19857)
}
__typedArg0 := let__18897
__typedArg1 := tmp19857
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19859 := Call(__e, PrimFunc(symis_b), V6116, tmp19858, B6112, L6113, Key6114, C6115)


__e.TailApply(PrimFunc(symshen_4gc), B6112, tmp19859)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19860 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp19861 := Call(__e, PrimFunc(symshen_4assoc_1_6), symassoc, tmp19845, tmp19860)


tmp19862 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp19861)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp19861
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19862

tmp19863 := MakeNative(func(__e *ControlFlow) {
V6121 := __e.Get(1)
_ = V6121
__e.Return(MakeNative(func(__e *ControlFlow) {
B6117 := __e.Get(1)
_ = B6117
__e.Return(MakeNative(func(__e *ControlFlow) {
L6118 := __e.Get(1)
_ = L6118
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6119 := __e.Get(1)
_ = Key6119
__e.Return(MakeNative(func(__e *ControlFlow) {
C6120 := __e.Get(1)
_ = C6120
tmp19864 := Call(__e, PrimFunc(symshen_4newpv), B6117)


let__18898 := tmp19864
_ = let__18898

tmp19865 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19866 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19865)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19865
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19867 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18898, tmp19866)
}
__typedArg0 := let__18898
__typedArg1 := tmp19866
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19868 := Call(__e, PrimFunc(symis_b), V6121, tmp19867, B6117, L6118, Key6119, C6120)


__e.TailApply(PrimFunc(symshen_4gc), B6117, tmp19868)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19869 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp19870 := Call(__e, PrimFunc(symshen_4assoc_1_6), symatom_2, tmp19863, tmp19869)


tmp19871 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp19870)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp19870
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19871

tmp19872 := MakeNative(func(__e *ControlFlow) {
V6126 := __e.Get(1)
_ = V6126
__e.Return(MakeNative(func(__e *ControlFlow) {
B6122 := __e.Get(1)
_ = B6122
__e.Return(MakeNative(func(__e *ControlFlow) {
L6123 := __e.Get(1)
_ = L6123
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6124 := __e.Get(1)
_ = Key6124
__e.Return(MakeNative(func(__e *ControlFlow) {
C6125 := __e.Get(1)
_ = C6125
tmp19873 := Call(__e, PrimFunc(symshen_4newpv), B6122)


let__18899 := tmp19873
_ = let__18899

tmp19874 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19875 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19874)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19874
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19876 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18899, tmp19875)
}
__typedArg0 := let__18899
__typedArg1 := tmp19875
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19877 := Call(__e, PrimFunc(symis_b), V6126, tmp19876, B6122, L6123, Key6124, C6125)


__e.TailApply(PrimFunc(symshen_4gc), B6122, tmp19877)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19878 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp19879 := Call(__e, PrimFunc(symshen_4assoc_1_6), symboolean_2, tmp19872, tmp19878)


tmp19880 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp19879)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp19879
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19880

tmp19881 := MakeNative(func(__e *ControlFlow) {
V6131 := __e.Get(1)
_ = V6131
__e.Return(MakeNative(func(__e *ControlFlow) {
B6127 := __e.Get(1)
_ = B6127
__e.Return(MakeNative(func(__e *ControlFlow) {
L6128 := __e.Get(1)
_ = L6128
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6129 := __e.Get(1)
_ = Key6129
__e.Return(MakeNative(func(__e *ControlFlow) {
C6130 := __e.Get(1)
_ = C6130
tmp19882 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19883 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19882)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19882
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19884 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp19883)
}
__typedArg0 := symstring
__typedArg1 := tmp19883
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6131, tmp19884, B6127, L6128, Key6129, C6130)
return


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

tmp19886 := Call(__e, PrimFunc(symshen_4assoc_1_6), symbootstrap, tmp19881, tmp19885)


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
V6136 := __e.Get(1)
_ = V6136
__e.Return(MakeNative(func(__e *ControlFlow) {
B6132 := __e.Get(1)
_ = B6132
__e.Return(MakeNative(func(__e *ControlFlow) {
L6133 := __e.Get(1)
_ = L6133
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6134 := __e.Get(1)
_ = Key6134
__e.Return(MakeNative(func(__e *ControlFlow) {
C6135 := __e.Get(1)
_ = C6135
tmp19889 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19890 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19889)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19889
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19891 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp19890)
}
__typedArg0 := symsymbol
__typedArg1 := tmp19890
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6136, tmp19891, B6132, L6133, Key6134, C6135)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19892 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp19893 := Call(__e, PrimFunc(symshen_4assoc_1_6), symbound_2, tmp19888, tmp19892)


tmp19894 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp19893)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp19893
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19894

tmp19895 := MakeNative(func(__e *ControlFlow) {
V6141 := __e.Get(1)
_ = V6141
__e.Return(MakeNative(func(__e *ControlFlow) {
B6137 := __e.Get(1)
_ = B6137
__e.Return(MakeNative(func(__e *ControlFlow) {
L6138 := __e.Get(1)
_ = L6138
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6139 := __e.Get(1)
_ = Key6139
__e.Return(MakeNative(func(__e *ControlFlow) {
C6140 := __e.Get(1)
_ = C6140
tmp19896 := Call(__e, PrimFunc(symshen_4newpv), B6137)


let__18900 := tmp19896
_ = let__18900

tmp19897 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18900, Nil)
}
__typedArg0 := let__18900
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19898 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19897)
}
__typedArg0 := symlist
__typedArg1 := tmp19897
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

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
return PrimCons(tmp19898, tmp19900)
}
__typedArg0 := tmp19898
__typedArg1 := tmp19900
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19902 := Call(__e, PrimFunc(symis_b), V6141, tmp19901, B6137, L6138, Key6139, C6140)


__e.TailApply(PrimFunc(symshen_4gc), B6137, tmp19902)
return


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

tmp19904 := Call(__e, PrimFunc(symshen_4assoc_1_6), symshen_4ccons_2, tmp19895, tmp19903)


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
V6146 := __e.Get(1)
_ = V6146
__e.Return(MakeNative(func(__e *ControlFlow) {
B6142 := __e.Get(1)
_ = B6142
__e.Return(MakeNative(func(__e *ControlFlow) {
L6143 := __e.Get(1)
_ = L6143
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6144 := __e.Get(1)
_ = Key6144
__e.Return(MakeNative(func(__e *ControlFlow) {
C6145 := __e.Get(1)
_ = C6145
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

__e.TailApply(PrimFunc(symis_b), V6146, tmp19909, B6142, L6143, Key6144, C6145)
return


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

tmp19911 := Call(__e, PrimFunc(symshen_4assoc_1_6), symcd, tmp19906, tmp19910)


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
V6151 := __e.Get(1)
_ = V6151
__e.Return(MakeNative(func(__e *ControlFlow) {
B6147 := __e.Get(1)
_ = B6147
__e.Return(MakeNative(func(__e *ControlFlow) {
L6148 := __e.Get(1)
_ = L6148
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6149 := __e.Get(1)
_ = Key6149
__e.Return(MakeNative(func(__e *ControlFlow) {
C6150 := __e.Get(1)
_ = C6150
tmp19914 := Call(__e, PrimFunc(symshen_4newpv), B6147)


let__18901 := tmp19914
_ = let__18901

tmp19915 := Call(__e, PrimFunc(symshen_4newpv), B6147)


let__18902 := tmp19915
_ = let__18902

tmp19916 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18901, Nil)
}
__typedArg0 := let__18901
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19917 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp19916)
}
__typedArg0 := symstream
__typedArg1 := tmp19916
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19918 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18902, Nil)
}
__typedArg0 := let__18902
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19919 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19918)
}
__typedArg0 := symlist
__typedArg1 := tmp19918
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19920 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19919, Nil)
}
__typedArg0 := tmp19919
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19921 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19920)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19920
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19922 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19917, tmp19921)
}
__typedArg0 := tmp19917
__typedArg1 := tmp19921
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19923 := Call(__e, PrimFunc(symis_b), V6151, tmp19922, B6147, L6148, Key6149, C6150)


tmp19924 := Call(__e, PrimFunc(symshen_4gc), B6147, tmp19923)


__e.TailApply(PrimFunc(symshen_4gc), B6147, tmp19924)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19925 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp19926 := Call(__e, PrimFunc(symshen_4assoc_1_6), symclose, tmp19913, tmp19925)


tmp19927 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp19926)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp19926
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19927

tmp19928 := MakeNative(func(__e *ControlFlow) {
V6156 := __e.Get(1)
_ = V6156
__e.Return(MakeNative(func(__e *ControlFlow) {
B6152 := __e.Get(1)
_ = B6152
__e.Return(MakeNative(func(__e *ControlFlow) {
L6153 := __e.Get(1)
_ = L6153
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6154 := __e.Get(1)
_ = Key6154
__e.Return(MakeNative(func(__e *ControlFlow) {
C6155 := __e.Get(1)
_ = C6155
tmp19929 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19930 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19929)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19929
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19931 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp19930)
}
__typedArg0 := symstring
__typedArg1 := tmp19930
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19932 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19931, Nil)
}
__typedArg0 := tmp19931
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

__e.TailApply(PrimFunc(symis_b), V6156, tmp19934, B6152, L6153, Key6154, C6155)
return


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

tmp19936 := Call(__e, PrimFunc(symshen_4assoc_1_6), symcn, tmp19928, tmp19935)


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
V6161 := __e.Get(1)
_ = V6161
__e.Return(MakeNative(func(__e *ControlFlow) {
B6157 := __e.Get(1)
_ = B6157
__e.Return(MakeNative(func(__e *ControlFlow) {
L6158 := __e.Get(1)
_ = L6158
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6159 := __e.Get(1)
_ = Key6159
__e.Return(MakeNative(func(__e *ControlFlow) {
C6160 := __e.Get(1)
_ = C6160
tmp19939 := Call(__e, PrimFunc(symshen_4newpv), B6157)


let__18903 := tmp19939
_ = let__18903

tmp19940 := Call(__e, PrimFunc(symshen_4newpv), B6157)


let__18904 := tmp19940
_ = let__18904

tmp19941 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18903, Nil)
}
__typedArg0 := let__18903
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19942 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19941)
}
__typedArg0 := symlist
__typedArg1 := tmp19941
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19943 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18903, Nil)
}
__typedArg0 := let__18903
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
return PrimCons(let__18904, Nil)
}
__typedArg0 := let__18904
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19946 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19944, tmp19945)
}
__typedArg0 := tmp19944
__typedArg1 := tmp19945
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19947 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstr, tmp19946)
}
__typedArg0 := symstr
__typedArg1 := tmp19946
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19948 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19947, Nil)
}
__typedArg0 := tmp19947
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19949 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19948)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19948
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19950 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19942, tmp19949)
}
__typedArg0 := tmp19942
__typedArg1 := tmp19949
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19951 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18903, Nil)
}
__typedArg0 := let__18903
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19952 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19951)
}
__typedArg0 := symlist
__typedArg1 := tmp19951
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19953 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18904, Nil)
}
__typedArg0 := let__18904
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19954 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19953)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19953
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19955 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19952, tmp19954)
}
__typedArg0 := tmp19952
__typedArg1 := tmp19954
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19956 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19955, Nil)
}
__typedArg0 := tmp19955
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19957 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19956)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19956
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19958 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19950, tmp19957)
}
__typedArg0 := tmp19950
__typedArg1 := tmp19957
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19959 := Call(__e, PrimFunc(symis_b), V6161, tmp19958, B6157, L6158, Key6159, C6160)


tmp19960 := Call(__e, PrimFunc(symshen_4gc), B6157, tmp19959)


__e.TailApply(PrimFunc(symshen_4gc), B6157, tmp19960)
return


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

tmp19962 := Call(__e, PrimFunc(symshen_4assoc_1_6), symcompile, tmp19938, tmp19961)


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
V6166 := __e.Get(1)
_ = V6166
__e.Return(MakeNative(func(__e *ControlFlow) {
B6162 := __e.Get(1)
_ = B6162
__e.Return(MakeNative(func(__e *ControlFlow) {
L6163 := __e.Get(1)
_ = L6163
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6164 := __e.Get(1)
_ = Key6164
__e.Return(MakeNative(func(__e *ControlFlow) {
C6165 := __e.Get(1)
_ = C6165
tmp19965 := Call(__e, PrimFunc(symshen_4newpv), B6162)


let__18905 := tmp19965
_ = let__18905

tmp19966 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19967 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19966)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19966
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19968 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18905, tmp19967)
}
__typedArg0 := let__18905
__typedArg1 := tmp19967
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19969 := Call(__e, PrimFunc(symis_b), V6166, tmp19968, B6162, L6163, Key6164, C6165)


__e.TailApply(PrimFunc(symshen_4gc), B6162, tmp19969)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19970 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp19971 := Call(__e, PrimFunc(symshen_4assoc_1_6), symcons_2, tmp19964, tmp19970)


tmp19972 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp19971)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp19971
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19972

tmp19973 := MakeNative(func(__e *ControlFlow) {
V6171 := __e.Get(1)
_ = V6171
__e.Return(MakeNative(func(__e *ControlFlow) {
B6167 := __e.Get(1)
_ = B6167
__e.Return(MakeNative(func(__e *ControlFlow) {
L6168 := __e.Get(1)
_ = L6168
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6169 := __e.Get(1)
_ = Key6169
__e.Return(MakeNative(func(__e *ControlFlow) {
C6170 := __e.Get(1)
_ = C6170
tmp19974 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19975 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19974)
}
__typedArg0 := symlist
__typedArg1 := tmp19974
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19976 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19975, Nil)
}
__typedArg0 := tmp19975
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19977 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19976)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19976
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6171, tmp19977, B6167, L6168, Key6169, C6170)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19978 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp19979 := Call(__e, PrimFunc(symshen_4assoc_1_6), symdatatypes, tmp19973, tmp19978)


tmp19980 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp19979)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp19979
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19980

tmp19981 := MakeNative(func(__e *ControlFlow) {
V6176 := __e.Get(1)
_ = V6176
__e.Return(MakeNative(func(__e *ControlFlow) {
B6172 := __e.Get(1)
_ = B6172
__e.Return(MakeNative(func(__e *ControlFlow) {
L6173 := __e.Get(1)
_ = L6173
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6174 := __e.Get(1)
_ = Key6174
__e.Return(MakeNative(func(__e *ControlFlow) {
C6175 := __e.Get(1)
_ = C6175
tmp19982 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19983 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19982)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19982
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19984 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp19983)
}
__typedArg0 := symsymbol
__typedArg1 := tmp19983
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6176, tmp19984, B6172, L6173, Key6174, C6175)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19985 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp19986 := Call(__e, PrimFunc(symshen_4assoc_1_6), symdestroy, tmp19981, tmp19985)


tmp19987 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp19986)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp19986
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19987

tmp19988 := MakeNative(func(__e *ControlFlow) {
V6181 := __e.Get(1)
_ = V6181
__e.Return(MakeNative(func(__e *ControlFlow) {
B6177 := __e.Get(1)
_ = B6177
__e.Return(MakeNative(func(__e *ControlFlow) {
L6178 := __e.Get(1)
_ = L6178
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6179 := __e.Get(1)
_ = Key6179
__e.Return(MakeNative(func(__e *ControlFlow) {
C6180 := __e.Get(1)
_ = C6180
tmp19989 := Call(__e, PrimFunc(symshen_4newpv), B6177)


let__18906 := tmp19989
_ = let__18906

tmp19990 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18906, Nil)
}
__typedArg0 := let__18906
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19991 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19990)
}
__typedArg0 := symlist
__typedArg1 := tmp19990
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19992 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18906, Nil)
}
__typedArg0 := let__18906
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19993 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19992)
}
__typedArg0 := symlist
__typedArg1 := tmp19992
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19994 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18906, Nil)
}
__typedArg0 := let__18906
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19995 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19994)
}
__typedArg0 := symlist
__typedArg1 := tmp19994
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19996 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19995, Nil)
}
__typedArg0 := tmp19995
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19997 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19996)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19996
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19998 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19993, tmp19997)
}
__typedArg0 := tmp19993
__typedArg1 := tmp19997
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19999 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19998, Nil)
}
__typedArg0 := tmp19998
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20000 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19999)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19999
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20001 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19991, tmp20000)
}
__typedArg0 := tmp19991
__typedArg1 := tmp20000
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20002 := Call(__e, PrimFunc(symis_b), V6181, tmp20001, B6177, L6178, Key6179, C6180)


__e.TailApply(PrimFunc(symshen_4gc), B6177, tmp20002)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20003 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20004 := Call(__e, PrimFunc(symshen_4assoc_1_6), symdifference, tmp19988, tmp20003)


tmp20005 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20004)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20004
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20005

tmp20006 := MakeNative(func(__e *ControlFlow) {
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
tmp20007 := Call(__e, PrimFunc(symshen_4newpv), B6182)


let__18907 := tmp20007
_ = let__18907

tmp20008 := Call(__e, PrimFunc(symshen_4newpv), B6182)


let__18908 := tmp20008
_ = let__18908

tmp20009 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18908, Nil)
}
__typedArg0 := let__18908
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20010 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20009)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20009
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20011 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18908, tmp20010)
}
__typedArg0 := let__18908
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

tmp20014 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18907, tmp20013)
}
__typedArg0 := let__18907
__typedArg1 := tmp20013
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20015 := Call(__e, PrimFunc(symis_b), V6186, tmp20014, B6182, L6183, Key6184, C6185)


tmp20016 := Call(__e, PrimFunc(symshen_4gc), B6182, tmp20015)


__e.TailApply(PrimFunc(symshen_4gc), B6182, tmp20016)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20017 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20018 := Call(__e, PrimFunc(symshen_4assoc_1_6), symdo, tmp20006, tmp20017)


tmp20019 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20018)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20018
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20019

tmp20020 := MakeNative(func(__e *ControlFlow) {
V6191 := __e.Get(1)
_ = V6191
__e.Return(MakeNative(func(__e *ControlFlow) {
B6187 := __e.Get(1)
_ = B6187
__e.Return(MakeNative(func(__e *ControlFlow) {
L6188 := __e.Get(1)
_ = L6188
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6189 := __e.Get(1)
_ = Key6189
__e.Return(MakeNative(func(__e *ControlFlow) {
C6190 := __e.Get(1)
_ = C6190
tmp20021 := Call(__e, PrimFunc(symshen_4newpv), B6187)


let__18909 := tmp20021
_ = let__18909

tmp20022 := Call(__e, PrimFunc(symshen_4newpv), B6187)


let__18910 := tmp20022
_ = let__18910

tmp20023 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18909, Nil)
}
__typedArg0 := let__18909
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20024 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20023)
}
__typedArg0 := symlist
__typedArg1 := tmp20023
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20025 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18909, Nil)
}
__typedArg0 := let__18909
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20026 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20025)
}
__typedArg0 := symlist
__typedArg1 := tmp20025
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20027 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18910, Nil)
}
__typedArg0 := let__18910
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20028 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20027)
}
__typedArg0 := symlist
__typedArg1 := tmp20027
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20029 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20028, Nil)
}
__typedArg0 := tmp20028
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20030 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20026, tmp20029)
}
__typedArg0 := tmp20026
__typedArg1 := tmp20029
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20031 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstr, tmp20030)
}
__typedArg0 := symstr
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
return PrimCons(tmp20024, tmp20033)
}
__typedArg0 := tmp20024
__typedArg1 := tmp20033
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20035 := Call(__e, PrimFunc(symis_b), V6191, tmp20034, B6187, L6188, Key6189, C6190)


tmp20036 := Call(__e, PrimFunc(symshen_4gc), B6187, tmp20035)


__e.TailApply(PrimFunc(symshen_4gc), B6187, tmp20036)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20037 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20038 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_5e_6, tmp20020, tmp20037)


tmp20039 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20038)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20038
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20039

tmp20040 := MakeNative(func(__e *ControlFlow) {
V6196 := __e.Get(1)
_ = V6196
__e.Return(MakeNative(func(__e *ControlFlow) {
B6192 := __e.Get(1)
_ = B6192
__e.Return(MakeNative(func(__e *ControlFlow) {
L6193 := __e.Get(1)
_ = L6193
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6194 := __e.Get(1)
_ = Key6194
__e.Return(MakeNative(func(__e *ControlFlow) {
C6195 := __e.Get(1)
_ = C6195
tmp20041 := Call(__e, PrimFunc(symshen_4newpv), B6192)


let__18911 := tmp20041
_ = let__18911

tmp20042 := Call(__e, PrimFunc(symshen_4newpv), B6192)


let__18912 := tmp20042
_ = let__18912

tmp20043 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18912, Nil)
}
__typedArg0 := let__18912
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
return PrimCons(let__18911, Nil)
}
__typedArg0 := let__18911
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20046 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20045)
}
__typedArg0 := symlist
__typedArg1 := tmp20045
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20047 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18912, Nil)
}
__typedArg0 := let__18912
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20048 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20047)
}
__typedArg0 := symlist
__typedArg1 := tmp20047
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20049 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20048, Nil)
}
__typedArg0 := tmp20048
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20050 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20046, tmp20049)
}
__typedArg0 := tmp20046
__typedArg1 := tmp20049
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20051 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstr, tmp20050)
}
__typedArg0 := symstr
__typedArg1 := tmp20050
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20052 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20051, Nil)
}
__typedArg0 := tmp20051
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20053 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20052)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20052
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20054 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20044, tmp20053)
}
__typedArg0 := tmp20044
__typedArg1 := tmp20053
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20055 := Call(__e, PrimFunc(symis_b), V6196, tmp20054, B6192, L6193, Key6194, C6195)


tmp20056 := Call(__e, PrimFunc(symshen_4gc), B6192, tmp20055)


__e.TailApply(PrimFunc(symshen_4gc), B6192, tmp20056)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20057 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20058 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_5_b_6, tmp20040, tmp20057)


tmp20059 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20058)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20058
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20059

tmp20060 := MakeNative(func(__e *ControlFlow) {
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
tmp20061 := Call(__e, PrimFunc(symshen_4newpv), B6197)


let__18913 := tmp20061
_ = let__18913

tmp20062 := Call(__e, PrimFunc(symshen_4newpv), B6197)


let__18914 := tmp20062
_ = let__18914

tmp20063 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18913, Nil)
}
__typedArg0 := let__18913
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
return PrimCons(let__18913, Nil)
}
__typedArg0 := let__18913
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20066 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20065)
}
__typedArg0 := symlist
__typedArg1 := tmp20065
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20067 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18914, Nil)
}
__typedArg0 := let__18914
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20068 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20067)
}
__typedArg0 := symlist
__typedArg1 := tmp20067
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20069 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20068, Nil)
}
__typedArg0 := tmp20068
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20070 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20066, tmp20069)
}
__typedArg0 := tmp20066
__typedArg1 := tmp20069
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20071 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstr, tmp20070)
}
__typedArg0 := symstr
__typedArg1 := tmp20070
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20072 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20071, Nil)
}
__typedArg0 := tmp20071
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20073 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20072)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20072
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20074 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20064, tmp20073)
}
__typedArg0 := tmp20064
__typedArg1 := tmp20073
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20075 := Call(__e, PrimFunc(symis_b), V6201, tmp20074, B6197, L6198, Key6199, C6200)


tmp20076 := Call(__e, PrimFunc(symshen_4gc), B6197, tmp20075)


__e.TailApply(PrimFunc(symshen_4gc), B6197, tmp20076)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20077 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20078 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_5end_6, tmp20060, tmp20077)


tmp20079 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20078)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20078
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20079

tmp20080 := MakeNative(func(__e *ControlFlow) {
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
tmp20081 := Call(__e, PrimFunc(symshen_4newpv), B6202)


let__18915 := tmp20081
_ = let__18915

tmp20082 := Call(__e, PrimFunc(symshen_4newpv), B6202)


let__18916 := tmp20082
_ = let__18916

tmp20083 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18915, Nil)
}
__typedArg0 := let__18915
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
return PrimCons(let__18916, Nil)
}
__typedArg0 := let__18916
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20086 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20084, tmp20085)
}
__typedArg0 := tmp20084
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
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
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
return PrimCons(tmp20087, tmp20089)
}
__typedArg0 := tmp20087
__typedArg1 := tmp20089
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20091 := Call(__e, PrimFunc(symis_b), V6206, tmp20090, B6202, L6203, Key6204, C6205)


tmp20092 := Call(__e, PrimFunc(symshen_4gc), B6202, tmp20091)


__e.TailApply(PrimFunc(symshen_4gc), B6202, tmp20092)
return


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

tmp20094 := Call(__e, PrimFunc(symshen_4assoc_1_6), symshen_4parse_1failure_2, tmp20080, tmp20093)


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
tmp20097 := Call(__e, PrimFunc(symshen_4newpv), B6207)


let__18917 := tmp20097
_ = let__18917

tmp20098 := Call(__e, PrimFunc(symshen_4newpv), B6207)


let__18918 := tmp20098
_ = let__18918

tmp20099 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18917, Nil)
}
__typedArg0 := let__18917
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
return PrimCons(let__18918, Nil)
}
__typedArg0 := let__18918
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20102 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20100, tmp20101)
}
__typedArg0 := tmp20100
__typedArg1 := tmp20101
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20103 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstr, tmp20102)
}
__typedArg0 := symstr
__typedArg1 := tmp20102
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20104 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20103, Nil)
}
__typedArg0 := tmp20103
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20105 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20104)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20104
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20106 := Call(__e, PrimFunc(symis_b), V6211, tmp20105, B6207, L6208, Key6209, C6210)


tmp20107 := Call(__e, PrimFunc(symshen_4gc), B6207, tmp20106)


__e.TailApply(PrimFunc(symshen_4gc), B6207, tmp20107)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20108 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20109 := Call(__e, PrimFunc(symshen_4assoc_1_6), symshen_4parse_1failure, tmp20096, tmp20108)


tmp20110 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20109)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20109
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20110

tmp20111 := MakeNative(func(__e *ControlFlow) {
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
tmp20112 := Call(__e, PrimFunc(symshen_4newpv), B6212)


let__18919 := tmp20112
_ = let__18919

tmp20113 := Call(__e, PrimFunc(symshen_4newpv), B6212)


let__18920 := tmp20113
_ = let__18920

tmp20114 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18919, Nil)
}
__typedArg0 := let__18919
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
return PrimCons(let__18920, Nil)
}
__typedArg0 := let__18920
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20117 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20115, tmp20116)
}
__typedArg0 := tmp20115
__typedArg1 := tmp20116
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20118 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstr, tmp20117)
}
__typedArg0 := symstr
__typedArg1 := tmp20117
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20119 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18920, Nil)
}
__typedArg0 := let__18920
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20120 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20119)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20119
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20121 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20118, tmp20120)
}
__typedArg0 := tmp20118
__typedArg1 := tmp20120
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20122 := Call(__e, PrimFunc(symis_b), V6216, tmp20121, B6212, L6213, Key6214, C6215)


tmp20123 := Call(__e, PrimFunc(symshen_4gc), B6212, tmp20122)


__e.TailApply(PrimFunc(symshen_4gc), B6212, tmp20123)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20124 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20125 := Call(__e, PrimFunc(symshen_4assoc_1_6), symshen_4_5_1out, tmp20111, tmp20124)


tmp20126 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20125)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20125
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20126

tmp20127 := MakeNative(func(__e *ControlFlow) {
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
tmp20128 := Call(__e, PrimFunc(symshen_4newpv), B6217)


let__18921 := tmp20128
_ = let__18921

tmp20129 := Call(__e, PrimFunc(symshen_4newpv), B6217)


let__18922 := tmp20129
_ = let__18922

tmp20130 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18922, Nil)
}
__typedArg0 := let__18922
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20131 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20130)
}
__typedArg0 := symlist
__typedArg1 := tmp20130
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20132 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18921, Nil)
}
__typedArg0 := let__18921
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20133 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20131, tmp20132)
}
__typedArg0 := tmp20131
__typedArg1 := tmp20132
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20134 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstr, tmp20133)
}
__typedArg0 := symstr
__typedArg1 := tmp20133
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20135 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18922, Nil)
}
__typedArg0 := let__18922
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
return PrimCons(tmp20136, Nil)
}
__typedArg0 := tmp20136
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20138 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20137)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20137
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20139 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20134, tmp20138)
}
__typedArg0 := tmp20134
__typedArg1 := tmp20138
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20140 := Call(__e, PrimFunc(symis_b), V6221, tmp20139, B6217, L6218, Key6219, C6220)


tmp20141 := Call(__e, PrimFunc(symshen_4gc), B6217, tmp20140)


__e.TailApply(PrimFunc(symshen_4gc), B6217, tmp20141)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20142 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20143 := Call(__e, PrimFunc(symshen_4assoc_1_6), symshen_4in_1_6, tmp20127, tmp20142)


tmp20144 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20143)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20143
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20144

tmp20145 := MakeNative(func(__e *ControlFlow) {
V6226 := __e.Get(1)
_ = V6226
__e.Return(MakeNative(func(__e *ControlFlow) {
B6222 := __e.Get(1)
_ = B6222
__e.Return(MakeNative(func(__e *ControlFlow) {
L6223 := __e.Get(1)
_ = L6223
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6224 := __e.Get(1)
_ = Key6224
__e.Return(MakeNative(func(__e *ControlFlow) {
C6225 := __e.Get(1)
_ = C6225
tmp20146 := Call(__e, PrimFunc(symshen_4newpv), B6222)


let__18923 := tmp20146
_ = let__18923

tmp20147 := Call(__e, PrimFunc(symshen_4newpv), B6222)


let__18924 := tmp20147
_ = let__18924

tmp20148 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18923, Nil)
}
__typedArg0 := let__18923
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20149 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20148)
}
__typedArg0 := symlist
__typedArg1 := tmp20148
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20150 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18923, Nil)
}
__typedArg0 := let__18923
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
return PrimCons(let__18924, Nil)
}
__typedArg0 := let__18924
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
return PrimCons(tmp20154, Nil)
}
__typedArg0 := tmp20154
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
return PrimCons(let__18924, tmp20156)
}
__typedArg0 := let__18924
__typedArg1 := tmp20156
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20158 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20157, Nil)
}
__typedArg0 := tmp20157
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20159 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20158)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20158
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20160 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20149, tmp20159)
}
__typedArg0 := tmp20149
__typedArg1 := tmp20159
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20161 := Call(__e, PrimFunc(symis_b), V6226, tmp20160, B6222, L6223, Key6224, C6225)


tmp20162 := Call(__e, PrimFunc(symshen_4gc), B6222, tmp20161)


__e.TailApply(PrimFunc(symshen_4gc), B6222, tmp20162)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20163 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20164 := Call(__e, PrimFunc(symshen_4assoc_1_6), symshen_4comb, tmp20145, tmp20163)


tmp20165 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20164)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20164
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20165

tmp20166 := MakeNative(func(__e *ControlFlow) {
V6231 := __e.Get(1)
_ = V6231
__e.Return(MakeNative(func(__e *ControlFlow) {
B6227 := __e.Get(1)
_ = B6227
__e.Return(MakeNative(func(__e *ControlFlow) {
L6228 := __e.Get(1)
_ = L6228
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6229 := __e.Get(1)
_ = Key6229
__e.Return(MakeNative(func(__e *ControlFlow) {
C6230 := __e.Get(1)
_ = C6230
tmp20167 := Call(__e, PrimFunc(symshen_4newpv), B6227)


let__18925 := tmp20167
_ = let__18925

tmp20168 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18925, Nil)
}
__typedArg0 := let__18925
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
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20171 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20170)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20170
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20172 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20169, tmp20171)
}
__typedArg0 := tmp20169
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
return PrimCons(let__18925, tmp20174)
}
__typedArg0 := let__18925
__typedArg1 := tmp20174
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20176 := Call(__e, PrimFunc(symis_b), V6231, tmp20175, B6227, L6228, Key6229, C6230)


__e.TailApply(PrimFunc(symshen_4gc), B6227, tmp20176)
return


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

tmp20178 := Call(__e, PrimFunc(symshen_4assoc_1_6), symelement_2, tmp20166, tmp20177)


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
V6236 := __e.Get(1)
_ = V6236
__e.Return(MakeNative(func(__e *ControlFlow) {
B6232 := __e.Get(1)
_ = B6232
__e.Return(MakeNative(func(__e *ControlFlow) {
L6233 := __e.Get(1)
_ = L6233
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6234 := __e.Get(1)
_ = Key6234
__e.Return(MakeNative(func(__e *ControlFlow) {
C6235 := __e.Get(1)
_ = C6235
tmp20181 := Call(__e, PrimFunc(symshen_4newpv), B6232)


let__18926 := tmp20181
_ = let__18926

tmp20182 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20183 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20182)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20182
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20184 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18926, tmp20183)
}
__typedArg0 := let__18926
__typedArg1 := tmp20183
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20185 := Call(__e, PrimFunc(symis_b), V6236, tmp20184, B6232, L6233, Key6234, C6235)


__e.TailApply(PrimFunc(symshen_4gc), B6232, tmp20185)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20186 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20187 := Call(__e, PrimFunc(symshen_4assoc_1_6), symempty_2, tmp20180, tmp20186)


tmp20188 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20187)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20187
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20188

tmp20189 := MakeNative(func(__e *ControlFlow) {
V6241 := __e.Get(1)
_ = V6241
__e.Return(MakeNative(func(__e *ControlFlow) {
B6237 := __e.Get(1)
_ = B6237
__e.Return(MakeNative(func(__e *ControlFlow) {
L6238 := __e.Get(1)
_ = L6238
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6239 := __e.Get(1)
_ = Key6239
__e.Return(MakeNative(func(__e *ControlFlow) {
C6240 := __e.Get(1)
_ = C6240
tmp20190 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20191 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20190)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20190
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20192 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp20191)
}
__typedArg0 := symsymbol
__typedArg1 := tmp20191
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6241, tmp20192, B6237, L6238, Key6239, C6240)
return


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

tmp20194 := Call(__e, PrimFunc(symshen_4assoc_1_6), symenable_1type_1theory, tmp20189, tmp20193)


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
V6246 := __e.Get(1)
_ = V6246
__e.Return(MakeNative(func(__e *ControlFlow) {
B6242 := __e.Get(1)
_ = B6242
__e.Return(MakeNative(func(__e *ControlFlow) {
L6243 := __e.Get(1)
_ = L6243
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6244 := __e.Get(1)
_ = Key6244
__e.Return(MakeNative(func(__e *ControlFlow) {
C6245 := __e.Get(1)
_ = C6245
tmp20197 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20198 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20197)
}
__typedArg0 := symlist
__typedArg1 := tmp20197
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20199 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20198, Nil)
}
__typedArg0 := tmp20198
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20200 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20199)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20199
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20201 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp20200)
}
__typedArg0 := symsymbol
__typedArg1 := tmp20200
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6246, tmp20201, B6242, L6243, Key6244, C6245)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20202 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20203 := Call(__e, PrimFunc(symshen_4assoc_1_6), symexternal, tmp20196, tmp20202)


tmp20204 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20203)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20203
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20204

tmp20205 := MakeNative(func(__e *ControlFlow) {
V6251 := __e.Get(1)
_ = V6251
__e.Return(MakeNative(func(__e *ControlFlow) {
B6247 := __e.Get(1)
_ = B6247
__e.Return(MakeNative(func(__e *ControlFlow) {
L6248 := __e.Get(1)
_ = L6248
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6249 := __e.Get(1)
_ = Key6249
__e.Return(MakeNative(func(__e *ControlFlow) {
C6250 := __e.Get(1)
_ = C6250
tmp20206 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
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
return PrimCons(symexception, tmp20207)
}
__typedArg0 := symexception
__typedArg1 := tmp20207
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6251, tmp20208, B6247, L6248, Key6249, C6250)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20209 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20210 := Call(__e, PrimFunc(symshen_4assoc_1_6), symerror_1to_1string, tmp20205, tmp20209)


tmp20211 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20210)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20210
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20211

tmp20212 := MakeNative(func(__e *ControlFlow) {
V6256 := __e.Get(1)
_ = V6256
__e.Return(MakeNative(func(__e *ControlFlow) {
B6252 := __e.Get(1)
_ = B6252
__e.Return(MakeNative(func(__e *ControlFlow) {
L6253 := __e.Get(1)
_ = L6253
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6254 := __e.Get(1)
_ = Key6254
__e.Return(MakeNative(func(__e *ControlFlow) {
C6255 := __e.Get(1)
_ = C6255
tmp20213 := Call(__e, PrimFunc(symshen_4newpv), B6252)


let__18927 := tmp20213
_ = let__18927

tmp20214 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20215 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20214)
}
__typedArg0 := symlist
__typedArg1 := tmp20214
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20216 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20215, Nil)
}
__typedArg0 := tmp20215
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20217 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20216)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20216
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20218 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18927, tmp20217)
}
__typedArg0 := let__18927
__typedArg1 := tmp20217
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20219 := Call(__e, PrimFunc(symis_b), V6256, tmp20218, B6252, L6253, Key6254, C6255)


__e.TailApply(PrimFunc(symshen_4gc), B6252, tmp20219)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20220 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20221 := Call(__e, PrimFunc(symshen_4assoc_1_6), symexplode, tmp20212, tmp20220)


tmp20222 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20221)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20221
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20222

tmp20223 := MakeNative(func(__e *ControlFlow) {
V6261 := __e.Get(1)
_ = V6261
__e.Return(MakeNative(func(__e *ControlFlow) {
B6257 := __e.Get(1)
_ = B6257
__e.Return(MakeNative(func(__e *ControlFlow) {
L6258 := __e.Get(1)
_ = L6258
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6259 := __e.Get(1)
_ = Key6259
__e.Return(MakeNative(func(__e *ControlFlow) {
C6260 := __e.Get(1)
_ = C6260
tmp20224 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
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
return PrimCons(symsymbol, tmp20225)
}
__typedArg0 := symsymbol
__typedArg1 := tmp20225
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6261, tmp20226, B6257, L6258, Key6259, C6260)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20227 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20228 := Call(__e, PrimFunc(symshen_4assoc_1_6), symfactorise, tmp20223, tmp20227)


tmp20229 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20228)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20228
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20229

tmp20230 := MakeNative(func(__e *ControlFlow) {
V6266 := __e.Get(1)
_ = V6266
__e.Return(MakeNative(func(__e *ControlFlow) {
B6262 := __e.Get(1)
_ = B6262
__e.Return(MakeNative(func(__e *ControlFlow) {
L6263 := __e.Get(1)
_ = L6263
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6264 := __e.Get(1)
_ = Key6264
__e.Return(MakeNative(func(__e *ControlFlow) {
C6265 := __e.Get(1)
_ = C6265
tmp20231 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20232 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20231)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20231
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6266, tmp20232, B6262, L6263, Key6264, C6265)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20233 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20234 := Call(__e, PrimFunc(symshen_4assoc_1_6), symfactorise_2, tmp20230, tmp20233)


tmp20235 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20234)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20234
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20235

tmp20236 := MakeNative(func(__e *ControlFlow) {
V6271 := __e.Get(1)
_ = V6271
__e.Return(MakeNative(func(__e *ControlFlow) {
B6267 := __e.Get(1)
_ = B6267
__e.Return(MakeNative(func(__e *ControlFlow) {
L6268 := __e.Get(1)
_ = L6268
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6269 := __e.Get(1)
_ = Key6269
__e.Return(MakeNative(func(__e *ControlFlow) {
C6270 := __e.Get(1)
_ = C6270
tmp20237 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20238 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20237)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20237
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6271, tmp20238, B6267, L6268, Key6269, C6270)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20239 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20240 := Call(__e, PrimFunc(symshen_4assoc_1_6), symfail, tmp20236, tmp20239)


tmp20241 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20240)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20240
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20241

tmp20242 := MakeNative(func(__e *ControlFlow) {
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
tmp20243 := Call(__e, PrimFunc(symshen_4newpv), B6272)


let__18928 := tmp20243
_ = let__18928

tmp20244 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18928, Nil)
}
__typedArg0 := let__18928
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20245 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20244)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20244
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20246 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18928, tmp20245)
}
__typedArg0 := let__18928
__typedArg1 := tmp20245
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20247 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18928, Nil)
}
__typedArg0 := let__18928
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20248 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20247)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20247
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20249 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18928, tmp20248)
}
__typedArg0 := let__18928
__typedArg1 := tmp20248
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20250 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20249, Nil)
}
__typedArg0 := tmp20249
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20251 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20250)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20250
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20252 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20246, tmp20251)
}
__typedArg0 := tmp20246
__typedArg1 := tmp20251
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20253 := Call(__e, PrimFunc(symis_b), V6276, tmp20252, B6272, L6273, Key6274, C6275)


__e.TailApply(PrimFunc(symshen_4gc), B6272, tmp20253)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20254 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20255 := Call(__e, PrimFunc(symshen_4assoc_1_6), symfix, tmp20242, tmp20254)


tmp20256 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20255)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20255
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20256

tmp20257 := MakeNative(func(__e *ControlFlow) {
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
tmp20258 := Call(__e, PrimFunc(symshen_4newpv), B6277)


let__18929 := tmp20258
_ = let__18929

tmp20259 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18929, Nil)
}
__typedArg0 := let__18929
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20260 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlazy, tmp20259)
}
__typedArg0 := symlazy
__typedArg1 := tmp20259
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20261 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20260, Nil)
}
__typedArg0 := tmp20260
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20262 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20261)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20261
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20263 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18929, tmp20262)
}
__typedArg0 := let__18929
__typedArg1 := tmp20262
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20264 := Call(__e, PrimFunc(symis_b), V6281, tmp20263, B6277, L6278, Key6279, C6280)


__e.TailApply(PrimFunc(symshen_4gc), B6277, tmp20264)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20265 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20266 := Call(__e, PrimFunc(symshen_4assoc_1_6), symfreeze, tmp20257, tmp20265)


tmp20267 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20266)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20266
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20267

tmp20268 := MakeNative(func(__e *ControlFlow) {
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
tmp20269 := Call(__e, PrimFunc(symshen_4newpv), B6282)


let__18930 := tmp20269
_ = let__18930

tmp20270 := Call(__e, PrimFunc(symshen_4newpv), B6282)


let__18931 := tmp20270
_ = let__18931

tmp20271 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18930, Nil)
}
__typedArg0 := let__18930
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20272 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_d, tmp20271)
}
__typedArg0 := sym_d
__typedArg1 := tmp20271
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20273 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18931, tmp20272)
}
__typedArg0 := let__18931
__typedArg1 := tmp20272
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20274 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18931, Nil)
}
__typedArg0 := let__18931
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20275 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20274)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20274
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20276 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20273, tmp20275)
}
__typedArg0 := tmp20273
__typedArg1 := tmp20275
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20277 := Call(__e, PrimFunc(symis_b), V6286, tmp20276, B6282, L6283, Key6284, C6285)


tmp20278 := Call(__e, PrimFunc(symshen_4gc), B6282, tmp20277)


__e.TailApply(PrimFunc(symshen_4gc), B6282, tmp20278)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20279 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20280 := Call(__e, PrimFunc(symshen_4assoc_1_6), symfst, tmp20268, tmp20279)


tmp20281 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20280)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20280
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20281

tmp20282 := MakeNative(func(__e *ControlFlow) {
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
tmp20283 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
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
return PrimCons(symsymbol, tmp20284)
}
__typedArg0 := symsymbol
__typedArg1 := tmp20284
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6291, tmp20285, B6287, L6288, Key6289, C6290)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20286 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20287 := Call(__e, PrimFunc(symshen_4assoc_1_6), symgensym, tmp20282, tmp20286)


tmp20288 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20287)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20287
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20288

tmp20289 := MakeNative(func(__e *ControlFlow) {
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
tmp20290 := Call(__e, PrimFunc(symshen_4newpv), B6292)


let__18932 := tmp20290
_ = let__18932

tmp20291 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18932, Nil)
}
__typedArg0 := let__18932
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20292 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20291)
}
__typedArg0 := symlist
__typedArg1 := tmp20291
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20293 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
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
return PrimCons(let__18932, tmp20294)
}
__typedArg0 := let__18932
__typedArg1 := tmp20294
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20296 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20295, Nil)
}
__typedArg0 := tmp20295
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20297 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20296)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20296
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20298 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20292, tmp20297)
}
__typedArg0 := tmp20292
__typedArg1 := tmp20297
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20299 := Call(__e, PrimFunc(symis_b), V6296, tmp20298, B6292, L6293, Key6294, C6295)


__e.TailApply(PrimFunc(symshen_4gc), B6292, tmp20299)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20300 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20301 := Call(__e, PrimFunc(symshen_4assoc_1_6), symshen_4hds_a_2, tmp20289, tmp20300)


tmp20302 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20301)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20301
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20302

tmp20303 := MakeNative(func(__e *ControlFlow) {
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
tmp20304 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20305 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20304)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20304
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20306 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp20305)
}
__typedArg0 := symsymbol
__typedArg1 := tmp20305
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6301, tmp20306, B6297, L6298, Key6299, C6300)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20307 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20308 := Call(__e, PrimFunc(symshen_4assoc_1_6), symhush, tmp20303, tmp20307)


tmp20309 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20308)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20308
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20309

tmp20310 := MakeNative(func(__e *ControlFlow) {
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
tmp20311 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20312 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20311)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20311
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6306, tmp20312, B6302, L6303, Key6304, C6305)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20313 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20314 := Call(__e, PrimFunc(symshen_4assoc_1_6), symhush_2, tmp20310, tmp20313)


tmp20315 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20314)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20314
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20315

tmp20316 := MakeNative(func(__e *ControlFlow) {
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
tmp20317 := Call(__e, PrimFunc(symshen_4newpv), B6307)


let__18933 := tmp20317
_ = let__18933

tmp20318 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18933, Nil)
}
__typedArg0 := let__18933
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20319 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp20318)
}
__typedArg0 := symvector
__typedArg1 := tmp20318
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20320 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18933, Nil)
}
__typedArg0 := let__18933
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20321 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20320)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20320
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20322 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp20321)
}
__typedArg0 := symnumber
__typedArg1 := tmp20321
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20323 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20322, Nil)
}
__typedArg0 := tmp20322
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20324 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20323)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20323
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20325 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20319, tmp20324)
}
__typedArg0 := tmp20319
__typedArg1 := tmp20324
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20326 := Call(__e, PrimFunc(symis_b), V6311, tmp20325, B6307, L6308, Key6309, C6310)


__e.TailApply(PrimFunc(symshen_4gc), B6307, tmp20326)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20327 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20328 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_5_1vector, tmp20316, tmp20327)


tmp20329 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20328)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20328
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20329

tmp20330 := MakeNative(func(__e *ControlFlow) {
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
tmp20331 := Call(__e, PrimFunc(symshen_4newpv), B6312)


let__18934 := tmp20331
_ = let__18934

tmp20332 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18934, Nil)
}
__typedArg0 := let__18934
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20333 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp20332)
}
__typedArg0 := symvector
__typedArg1 := tmp20332
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20334 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18934, Nil)
}
__typedArg0 := let__18934
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20335 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp20334)
}
__typedArg0 := symvector
__typedArg1 := tmp20334
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20336 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20335, Nil)
}
__typedArg0 := tmp20335
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20337 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20336)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20336
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20338 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18934, tmp20337)
}
__typedArg0 := let__18934
__typedArg1 := tmp20337
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20339 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20338, Nil)
}
__typedArg0 := tmp20338
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20340 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20339)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20339
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20341 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp20340)
}
__typedArg0 := symnumber
__typedArg1 := tmp20340
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20342 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20341, Nil)
}
__typedArg0 := tmp20341
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20343 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20342)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20342
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20344 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20333, tmp20343)
}
__typedArg0 := tmp20333
__typedArg1 := tmp20343
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20345 := Call(__e, PrimFunc(symis_b), V6316, tmp20344, B6312, L6313, Key6314, C6315)


__e.TailApply(PrimFunc(symshen_4gc), B6312, tmp20345)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20346 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20347 := Call(__e, PrimFunc(symshen_4assoc_1_6), symvector_1_6, tmp20330, tmp20346)


tmp20348 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20347)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20347
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20348

tmp20349 := MakeNative(func(__e *ControlFlow) {
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
tmp20350 := Call(__e, PrimFunc(symshen_4newpv), B6317)


let__18935 := tmp20350
_ = let__18935

tmp20351 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18935, Nil)
}
__typedArg0 := let__18935
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20352 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp20351)
}
__typedArg0 := symvector
__typedArg1 := tmp20351
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20353 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20352, Nil)
}
__typedArg0 := tmp20352
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20354 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20353)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20353
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20355 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp20354)
}
__typedArg0 := symnumber
__typedArg1 := tmp20354
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20356 := Call(__e, PrimFunc(symis_b), V6321, tmp20355, B6317, L6318, Key6319, C6320)


__e.TailApply(PrimFunc(symshen_4gc), B6317, tmp20356)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20357 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20358 := Call(__e, PrimFunc(symshen_4assoc_1_6), symvector, tmp20349, tmp20357)


tmp20359 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20358)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20358
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20359

tmp20360 := MakeNative(func(__e *ControlFlow) {
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
tmp20361 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20362 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20361)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20361
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20363 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp20362)
}
__typedArg0 := symsymbol
__typedArg1 := tmp20362
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6326, tmp20363, B6322, L6323, Key6324, C6325)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20364 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20365 := Call(__e, PrimFunc(symshen_4assoc_1_6), symget_1time, tmp20360, tmp20364)


tmp20366 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20365)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20365
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20366

tmp20367 := MakeNative(func(__e *ControlFlow) {
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
tmp20368 := Call(__e, PrimFunc(symshen_4newpv), B6327)


let__18936 := tmp20368
_ = let__18936

tmp20369 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20370 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20369)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20369
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20371 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp20370)
}
__typedArg0 := symnumber
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
return PrimCons(let__18936, tmp20373)
}
__typedArg0 := let__18936
__typedArg1 := tmp20373
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20375 := Call(__e, PrimFunc(symis_b), V6331, tmp20374, B6327, L6328, Key6329, C6330)


__e.TailApply(PrimFunc(symshen_4gc), B6327, tmp20375)
return


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

tmp20377 := Call(__e, PrimFunc(symshen_4assoc_1_6), symhash, tmp20367, tmp20376)


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
tmp20380 := Call(__e, PrimFunc(symshen_4newpv), B6332)


let__18937 := tmp20380
_ = let__18937

tmp20381 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18937, Nil)
}
__typedArg0 := let__18937
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20382 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20381)
}
__typedArg0 := symlist
__typedArg1 := tmp20381
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20383 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18937, Nil)
}
__typedArg0 := let__18937
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20384 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20383)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20383
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20385 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20382, tmp20384)
}
__typedArg0 := tmp20382
__typedArg1 := tmp20384
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20386 := Call(__e, PrimFunc(symis_b), V6336, tmp20385, B6332, L6333, Key6334, C6335)


__e.TailApply(PrimFunc(symshen_4gc), B6332, tmp20386)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20387 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20388 := Call(__e, PrimFunc(symshen_4assoc_1_6), symhead, tmp20379, tmp20387)


tmp20389 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20388)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20388
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20389

tmp20390 := MakeNative(func(__e *ControlFlow) {
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
tmp20391 := Call(__e, PrimFunc(symshen_4newpv), B6337)


let__18938 := tmp20391
_ = let__18938

tmp20392 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18938, Nil)
}
__typedArg0 := let__18938
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20393 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp20392)
}
__typedArg0 := symvector
__typedArg1 := tmp20392
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20394 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18938, Nil)
}
__typedArg0 := let__18938
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20395 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20394)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20394
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20396 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20393, tmp20395)
}
__typedArg0 := tmp20393
__typedArg1 := tmp20395
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20397 := Call(__e, PrimFunc(symis_b), V6341, tmp20396, B6337, L6338, Key6339, C6340)


__e.TailApply(PrimFunc(symshen_4gc), B6337, tmp20397)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20398 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20399 := Call(__e, PrimFunc(symshen_4assoc_1_6), symhdv, tmp20390, tmp20398)


tmp20400 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20399)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20399
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20400

tmp20401 := MakeNative(func(__e *ControlFlow) {
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
tmp20402 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20403 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20402)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20402
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20404 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp20403)
}
__typedArg0 := symstring
__typedArg1 := tmp20403
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6346, tmp20404, B6342, L6343, Key6344, C6345)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20405 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20406 := Call(__e, PrimFunc(symshen_4assoc_1_6), symhdstr, tmp20401, tmp20405)


tmp20407 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20406)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20406
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20407

tmp20408 := MakeNative(func(__e *ControlFlow) {
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
tmp20409 := Call(__e, PrimFunc(symshen_4newpv), B6347)


let__18939 := tmp20409
_ = let__18939

tmp20410 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18939, Nil)
}
__typedArg0 := let__18939
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20411 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20410)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20410
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20412 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18939, tmp20411)
}
__typedArg0 := let__18939
__typedArg1 := tmp20411
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20413 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20412, Nil)
}
__typedArg0 := tmp20412
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20414 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20413)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20413
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20415 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18939, tmp20414)
}
__typedArg0 := let__18939
__typedArg1 := tmp20414
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20416 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20415, Nil)
}
__typedArg0 := tmp20415
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20417 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20416)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20416
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20418 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, tmp20417)
}
__typedArg0 := symboolean
__typedArg1 := tmp20417
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20419 := Call(__e, PrimFunc(symis_b), V6351, tmp20418, B6347, L6348, Key6349, C6350)


__e.TailApply(PrimFunc(symshen_4gc), B6347, tmp20419)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20420 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20421 := Call(__e, PrimFunc(symshen_4assoc_1_6), symif, tmp20408, tmp20420)


tmp20422 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20421)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20421
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20422

tmp20423 := MakeNative(func(__e *ControlFlow) {
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
tmp20424 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20425 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20424)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20424
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20426 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp20425)
}
__typedArg0 := symsymbol
__typedArg1 := tmp20425
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6356, tmp20426, B6352, L6353, Key6354, C6355)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20427 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20428 := Call(__e, PrimFunc(symshen_4assoc_1_6), symin_1package, tmp20423, tmp20427)


tmp20429 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20428)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20428
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20429

tmp20430 := MakeNative(func(__e *ControlFlow) {
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
tmp20431 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20432 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20431)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20431
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6361, tmp20432, B6357, L6358, Key6359, C6360)
return


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

tmp20434 := Call(__e, PrimFunc(symshen_4assoc_1_6), symit, tmp20430, tmp20433)


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
tmp20437 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20438 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20437)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20437
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6366, tmp20438, B6362, L6363, Key6364, C6365)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20439 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20440 := Call(__e, PrimFunc(symshen_4assoc_1_6), symimplementation, tmp20436, tmp20439)


tmp20441 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20440)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20440
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20441

tmp20442 := MakeNative(func(__e *ControlFlow) {
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
tmp20443 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20444 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20443)
}
__typedArg0 := symlist
__typedArg1 := tmp20443
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20445 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20446 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20445)
}
__typedArg0 := symlist
__typedArg1 := tmp20445
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20447 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20446, Nil)
}
__typedArg0 := tmp20446
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20448 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20447)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20447
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20449 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20444, tmp20448)
}
__typedArg0 := tmp20444
__typedArg1 := tmp20448
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6371, tmp20449, B6367, L6368, Key6369, C6370)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20450 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20451 := Call(__e, PrimFunc(symshen_4assoc_1_6), syminclude, tmp20442, tmp20450)


tmp20452 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20451)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20451
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20452

tmp20453 := MakeNative(func(__e *ControlFlow) {
V6376 := __e.Get(1)
_ = V6376
__e.Return(MakeNative(func(__e *ControlFlow) {
B6372 := __e.Get(1)
_ = B6372
__e.Return(MakeNative(func(__e *ControlFlow) {
L6373 := __e.Get(1)
_ = L6373
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6374 := __e.Get(1)
_ = Key6374
__e.Return(MakeNative(func(__e *ControlFlow) {
C6375 := __e.Get(1)
_ = C6375
tmp20454 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20455 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20454)
}
__typedArg0 := symlist
__typedArg1 := tmp20454
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20456 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20457 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20456)
}
__typedArg0 := symlist
__typedArg1 := tmp20456
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20458 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20457, Nil)
}
__typedArg0 := tmp20457
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20459 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20458)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20458
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20460 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20455, tmp20459)
}
__typedArg0 := tmp20455
__typedArg1 := tmp20459
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6376, tmp20460, B6372, L6373, Key6374, C6375)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20461 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20462 := Call(__e, PrimFunc(symshen_4assoc_1_6), syminclude_1all_1but, tmp20453, tmp20461)


tmp20463 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20462)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20462
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20463

tmp20464 := MakeNative(func(__e *ControlFlow) {
V6381 := __e.Get(1)
_ = V6381
__e.Return(MakeNative(func(__e *ControlFlow) {
B6377 := __e.Get(1)
_ = B6377
__e.Return(MakeNative(func(__e *ControlFlow) {
L6378 := __e.Get(1)
_ = L6378
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6379 := __e.Get(1)
_ = Key6379
__e.Return(MakeNative(func(__e *ControlFlow) {
C6380 := __e.Get(1)
_ = C6380
tmp20465 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20466 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20465)
}
__typedArg0 := symlist
__typedArg1 := tmp20465
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20467 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20466, Nil)
}
__typedArg0 := tmp20466
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

__e.TailApply(PrimFunc(symis_b), V6381, tmp20468, B6377, L6378, Key6379, C6380)
return


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

tmp20470 := Call(__e, PrimFunc(symshen_4assoc_1_6), symincluded, tmp20464, tmp20469)


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
V6386 := __e.Get(1)
_ = V6386
__e.Return(MakeNative(func(__e *ControlFlow) {
B6382 := __e.Get(1)
_ = B6382
__e.Return(MakeNative(func(__e *ControlFlow) {
L6383 := __e.Get(1)
_ = L6383
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6384 := __e.Get(1)
_ = Key6384
__e.Return(MakeNative(func(__e *ControlFlow) {
C6385 := __e.Get(1)
_ = C6385
tmp20473 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
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

__e.TailApply(PrimFunc(symis_b), V6386, tmp20474, B6382, L6383, Key6384, C6385)
return


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

tmp20476 := Call(__e, PrimFunc(symshen_4assoc_1_6), syminferences, tmp20472, tmp20475)


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
V6391 := __e.Get(1)
_ = V6391
__e.Return(MakeNative(func(__e *ControlFlow) {
B6387 := __e.Get(1)
_ = B6387
__e.Return(MakeNative(func(__e *ControlFlow) {
L6388 := __e.Get(1)
_ = L6388
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6389 := __e.Get(1)
_ = Key6389
__e.Return(MakeNative(func(__e *ControlFlow) {
C6390 := __e.Get(1)
_ = C6390
tmp20479 := Call(__e, PrimFunc(symshen_4newpv), B6387)


let__18940 := tmp20479
_ = let__18940

tmp20480 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20481 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20480)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20480
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20482 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp20481)
}
__typedArg0 := symstring
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
return PrimCons(let__18940, tmp20484)
}
__typedArg0 := let__18940
__typedArg1 := tmp20484
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20486 := Call(__e, PrimFunc(symis_b), V6391, tmp20485, B6387, L6388, Key6389, C6390)


__e.TailApply(PrimFunc(symshen_4gc), B6387, tmp20486)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20487 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20488 := Call(__e, PrimFunc(symshen_4assoc_1_6), symshen_4insert, tmp20478, tmp20487)


tmp20489 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20488)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20488
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20489

tmp20490 := MakeNative(func(__e *ControlFlow) {
V6396 := __e.Get(1)
_ = V6396
__e.Return(MakeNative(func(__e *ControlFlow) {
B6392 := __e.Get(1)
_ = B6392
__e.Return(MakeNative(func(__e *ControlFlow) {
L6393 := __e.Get(1)
_ = L6393
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6394 := __e.Get(1)
_ = Key6394
__e.Return(MakeNative(func(__e *ControlFlow) {
C6395 := __e.Get(1)
_ = C6395
tmp20491 := Call(__e, PrimFunc(symshen_4newpv), B6392)


let__18941 := tmp20491
_ = let__18941

tmp20492 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20493 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20492)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20492
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20494 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18941, tmp20493)
}
__typedArg0 := let__18941
__typedArg1 := tmp20493
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20495 := Call(__e, PrimFunc(symis_b), V6396, tmp20494, B6392, L6393, Key6394, C6395)


__e.TailApply(PrimFunc(symshen_4gc), B6392, tmp20495)
return


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

tmp20497 := Call(__e, PrimFunc(symshen_4assoc_1_6), syminteger_2, tmp20490, tmp20496)


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
V6401 := __e.Get(1)
_ = V6401
__e.Return(MakeNative(func(__e *ControlFlow) {
B6397 := __e.Get(1)
_ = B6397
__e.Return(MakeNative(func(__e *ControlFlow) {
L6398 := __e.Get(1)
_ = L6398
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6399 := __e.Get(1)
_ = Key6399
__e.Return(MakeNative(func(__e *ControlFlow) {
C6400 := __e.Get(1)
_ = C6400
tmp20500 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20501 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20500)
}
__typedArg0 := symlist
__typedArg1 := tmp20500
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20502 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20501, Nil)
}
__typedArg0 := tmp20501
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20503 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20502)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20502
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20504 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp20503)
}
__typedArg0 := symsymbol
__typedArg1 := tmp20503
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6401, tmp20504, B6397, L6398, Key6399, C6400)
return


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

tmp20506 := Call(__e, PrimFunc(symshen_4assoc_1_6), syminternal, tmp20499, tmp20505)


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
V6406 := __e.Get(1)
_ = V6406
__e.Return(MakeNative(func(__e *ControlFlow) {
B6402 := __e.Get(1)
_ = B6402
__e.Return(MakeNative(func(__e *ControlFlow) {
L6403 := __e.Get(1)
_ = L6403
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6404 := __e.Get(1)
_ = Key6404
__e.Return(MakeNative(func(__e *ControlFlow) {
C6405 := __e.Get(1)
_ = C6405
tmp20509 := Call(__e, PrimFunc(symshen_4newpv), B6402)


let__18942 := tmp20509
_ = let__18942

tmp20510 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18942, Nil)
}
__typedArg0 := let__18942
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20511 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20510)
}
__typedArg0 := symlist
__typedArg1 := tmp20510
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20512 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18942, Nil)
}
__typedArg0 := let__18942
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
return PrimCons(let__18942, Nil)
}
__typedArg0 := let__18942
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
return PrimCons(tmp20511, tmp20520)
}
__typedArg0 := tmp20511
__typedArg1 := tmp20520
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20522 := Call(__e, PrimFunc(symis_b), V6406, tmp20521, B6402, L6403, Key6404, C6405)


__e.TailApply(PrimFunc(symshen_4gc), B6402, tmp20522)
return


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

tmp20524 := Call(__e, PrimFunc(symshen_4assoc_1_6), symintersection, tmp20508, tmp20523)


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
V6411 := __e.Get(1)
_ = V6411
__e.Return(MakeNative(func(__e *ControlFlow) {
B6407 := __e.Get(1)
_ = B6407
__e.Return(MakeNative(func(__e *ControlFlow) {
L6408 := __e.Get(1)
_ = L6408
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6409 := __e.Get(1)
_ = Key6409
__e.Return(MakeNative(func(__e *ControlFlow) {
C6410 := __e.Get(1)
_ = C6410
tmp20527 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
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

__e.TailApply(PrimFunc(symis_b), V6411, tmp20528, B6407, L6408, Key6409, C6410)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20529 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20530 := Call(__e, PrimFunc(symshen_4assoc_1_6), symlanguage, tmp20526, tmp20529)


tmp20531 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20530)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20530
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20531

tmp20532 := MakeNative(func(__e *ControlFlow) {
V6416 := __e.Get(1)
_ = V6416
__e.Return(MakeNative(func(__e *ControlFlow) {
B6412 := __e.Get(1)
_ = B6412
__e.Return(MakeNative(func(__e *ControlFlow) {
L6413 := __e.Get(1)
_ = L6413
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6414 := __e.Get(1)
_ = Key6414
__e.Return(MakeNative(func(__e *ControlFlow) {
C6415 := __e.Get(1)
_ = C6415
tmp20533 := Call(__e, PrimFunc(symshen_4newpv), B6412)


let__18943 := tmp20533
_ = let__18943

tmp20534 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18943, Nil)
}
__typedArg0 := let__18943
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
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
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

tmp20538 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20535, tmp20537)
}
__typedArg0 := tmp20535
__typedArg1 := tmp20537
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20539 := Call(__e, PrimFunc(symis_b), V6416, tmp20538, B6412, L6413, Key6414, C6415)


__e.TailApply(PrimFunc(symshen_4gc), B6412, tmp20539)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20540 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20541 := Call(__e, PrimFunc(symshen_4assoc_1_6), symlength, tmp20532, tmp20540)


tmp20542 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20541)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20541
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20542

tmp20543 := MakeNative(func(__e *ControlFlow) {
V6421 := __e.Get(1)
_ = V6421
__e.Return(MakeNative(func(__e *ControlFlow) {
B6417 := __e.Get(1)
_ = B6417
__e.Return(MakeNative(func(__e *ControlFlow) {
L6418 := __e.Get(1)
_ = L6418
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6419 := __e.Get(1)
_ = Key6419
__e.Return(MakeNative(func(__e *ControlFlow) {
C6420 := __e.Get(1)
_ = C6420
tmp20544 := Call(__e, PrimFunc(symshen_4newpv), B6417)


let__18944 := tmp20544
_ = let__18944

tmp20545 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18944, Nil)
}
__typedArg0 := let__18944
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20546 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp20545)
}
__typedArg0 := symvector
__typedArg1 := tmp20545
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20547 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20548 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20547)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20547
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20549 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20546, tmp20548)
}
__typedArg0 := tmp20546
__typedArg1 := tmp20548
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20550 := Call(__e, PrimFunc(symis_b), V6421, tmp20549, B6417, L6418, Key6419, C6420)


__e.TailApply(PrimFunc(symshen_4gc), B6417, tmp20550)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20551 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20552 := Call(__e, PrimFunc(symshen_4assoc_1_6), symlimit, tmp20543, tmp20551)


tmp20553 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20552)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20552
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20553

tmp20554 := MakeNative(func(__e *ControlFlow) {
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
tmp20555 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symin, Nil)
}
__typedArg0 := symin
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20556 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp20555)
}
__typedArg0 := symstream
__typedArg1 := tmp20555
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20557 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunit, Nil)
}
__typedArg0 := symunit
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20558 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20557)
}
__typedArg0 := symlist
__typedArg1 := tmp20557
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20559 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20558, Nil)
}
__typedArg0 := tmp20558
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20560 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20559)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20559
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20561 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20556, tmp20560)
}
__typedArg0 := tmp20556
__typedArg1 := tmp20560
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6426, tmp20561, B6422, L6423, Key6424, C6425)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20562 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20563 := Call(__e, PrimFunc(symshen_4assoc_1_6), symlineread, tmp20554, tmp20562)


tmp20564 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20563)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20563
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20564

tmp20565 := MakeNative(func(__e *ControlFlow) {
V6431 := __e.Get(1)
_ = V6431
__e.Return(MakeNative(func(__e *ControlFlow) {
B6427 := __e.Get(1)
_ = B6427
__e.Return(MakeNative(func(__e *ControlFlow) {
L6428 := __e.Get(1)
_ = L6428
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6429 := __e.Get(1)
_ = Key6429
__e.Return(MakeNative(func(__e *ControlFlow) {
C6430 := __e.Get(1)
_ = C6430
tmp20566 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20567 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20566)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20566
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20568 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp20567)
}
__typedArg0 := symstring
__typedArg1 := tmp20567
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6431, tmp20568, B6427, L6428, Key6429, C6430)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20569 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20570 := Call(__e, PrimFunc(symshen_4assoc_1_6), symload, tmp20565, tmp20569)


tmp20571 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20570)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20570
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20571

tmp20572 := MakeNative(func(__e *ControlFlow) {
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
tmp20573 := Call(__e, PrimFunc(symshen_4newpv), B6432)


let__18945 := tmp20573
_ = let__18945

tmp20574 := Call(__e, PrimFunc(symshen_4newpv), B6432)


let__18946 := tmp20574
_ = let__18946

tmp20575 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18946, Nil)
}
__typedArg0 := let__18946
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20576 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20575)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20575
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20577 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18945, tmp20576)
}
__typedArg0 := let__18945
__typedArg1 := tmp20576
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20578 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18945, Nil)
}
__typedArg0 := let__18945
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20579 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20578)
}
__typedArg0 := symlist
__typedArg1 := tmp20578
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20580 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18946, Nil)
}
__typedArg0 := let__18946
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20581 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20580)
}
__typedArg0 := symlist
__typedArg1 := tmp20580
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20582 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20581, Nil)
}
__typedArg0 := tmp20581
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20583 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20582)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20582
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20584 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20579, tmp20583)
}
__typedArg0 := tmp20579
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
return PrimCons(tmp20577, tmp20586)
}
__typedArg0 := tmp20577
__typedArg1 := tmp20586
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20588 := Call(__e, PrimFunc(symis_b), V6436, tmp20587, B6432, L6433, Key6434, C6435)


tmp20589 := Call(__e, PrimFunc(symshen_4gc), B6432, tmp20588)


__e.TailApply(PrimFunc(symshen_4gc), B6432, tmp20589)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20590 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20591 := Call(__e, PrimFunc(symshen_4assoc_1_6), symmap, tmp20572, tmp20590)


tmp20592 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20591)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20591
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20592

tmp20593 := MakeNative(func(__e *ControlFlow) {
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
tmp20594 := Call(__e, PrimFunc(symshen_4newpv), B6437)


let__18947 := tmp20594
_ = let__18947

tmp20595 := Call(__e, PrimFunc(symshen_4newpv), B6437)


let__18948 := tmp20595
_ = let__18948

tmp20596 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18948, Nil)
}
__typedArg0 := let__18948
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20597 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20596)
}
__typedArg0 := symlist
__typedArg1 := tmp20596
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20598 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20597, Nil)
}
__typedArg0 := tmp20597
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20599 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20598)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20598
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20600 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18947, tmp20599)
}
__typedArg0 := let__18947
__typedArg1 := tmp20599
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20601 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18947, Nil)
}
__typedArg0 := let__18947
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20602 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20601)
}
__typedArg0 := symlist
__typedArg1 := tmp20601
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20603 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18948, Nil)
}
__typedArg0 := let__18948
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
return PrimCons(tmp20604, Nil)
}
__typedArg0 := tmp20604
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
return PrimCons(tmp20602, tmp20606)
}
__typedArg0 := tmp20602
__typedArg1 := tmp20606
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20608 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20607, Nil)
}
__typedArg0 := tmp20607
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20609 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20608)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20608
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20610 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20600, tmp20609)
}
__typedArg0 := tmp20600
__typedArg1 := tmp20609
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20611 := Call(__e, PrimFunc(symis_b), V6441, tmp20610, B6437, L6438, Key6439, C6440)


tmp20612 := Call(__e, PrimFunc(symshen_4gc), B6437, tmp20611)


__e.TailApply(PrimFunc(symshen_4gc), B6437, tmp20612)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20613 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20614 := Call(__e, PrimFunc(symshen_4assoc_1_6), symmapcan, tmp20593, tmp20613)


tmp20615 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20614)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20614
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20615

tmp20616 := MakeNative(func(__e *ControlFlow) {
V6446 := __e.Get(1)
_ = V6446
__e.Return(MakeNative(func(__e *ControlFlow) {
B6442 := __e.Get(1)
_ = B6442
__e.Return(MakeNative(func(__e *ControlFlow) {
L6443 := __e.Get(1)
_ = L6443
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6444 := __e.Get(1)
_ = Key6444
__e.Return(MakeNative(func(__e *ControlFlow) {
C6445 := __e.Get(1)
_ = C6445
tmp20617 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20618 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20617)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20617
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20619 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp20618)
}
__typedArg0 := symnumber
__typedArg1 := tmp20618
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6446, tmp20619, B6442, L6443, Key6444, C6445)
return


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

tmp20621 := Call(__e, PrimFunc(symshen_4assoc_1_6), symmaxinferences, tmp20616, tmp20620)


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
V6451 := __e.Get(1)
_ = V6451
__e.Return(MakeNative(func(__e *ControlFlow) {
B6447 := __e.Get(1)
_ = B6447
__e.Return(MakeNative(func(__e *ControlFlow) {
L6448 := __e.Get(1)
_ = L6448
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6449 := __e.Get(1)
_ = Key6449
__e.Return(MakeNative(func(__e *ControlFlow) {
C6450 := __e.Get(1)
_ = C6450
tmp20624 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20625 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20624)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20624
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20626 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp20625)
}
__typedArg0 := symnumber
__typedArg1 := tmp20625
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6451, tmp20626, B6447, L6448, Key6449, C6450)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20627 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20628 := Call(__e, PrimFunc(symshen_4assoc_1_6), symn_1_6string, tmp20623, tmp20627)


tmp20629 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20628)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20628
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20629

tmp20630 := MakeNative(func(__e *ControlFlow) {
V6456 := __e.Get(1)
_ = V6456
__e.Return(MakeNative(func(__e *ControlFlow) {
B6452 := __e.Get(1)
_ = B6452
__e.Return(MakeNative(func(__e *ControlFlow) {
L6453 := __e.Get(1)
_ = L6453
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6454 := __e.Get(1)
_ = Key6454
__e.Return(MakeNative(func(__e *ControlFlow) {
C6455 := __e.Get(1)
_ = C6455
tmp20631 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20632 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20631)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20631
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20633 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp20632)
}
__typedArg0 := symnumber
__typedArg1 := tmp20632
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6456, tmp20633, B6452, L6453, Key6454, C6455)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20634 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20635 := Call(__e, PrimFunc(symshen_4assoc_1_6), symnl, tmp20630, tmp20634)


tmp20636 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20635)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20635
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20636

tmp20637 := MakeNative(func(__e *ControlFlow) {
V6461 := __e.Get(1)
_ = V6461
__e.Return(MakeNative(func(__e *ControlFlow) {
B6457 := __e.Get(1)
_ = B6457
__e.Return(MakeNative(func(__e *ControlFlow) {
L6458 := __e.Get(1)
_ = L6458
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6459 := __e.Get(1)
_ = Key6459
__e.Return(MakeNative(func(__e *ControlFlow) {
C6460 := __e.Get(1)
_ = C6460
tmp20638 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20639 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20638)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20638
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20640 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, tmp20639)
}
__typedArg0 := symboolean
__typedArg1 := tmp20639
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6461, tmp20640, B6457, L6458, Key6459, C6460)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20641 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20642 := Call(__e, PrimFunc(symshen_4assoc_1_6), symnot, tmp20637, tmp20641)


tmp20643 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20642)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20642
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20643

tmp20644 := MakeNative(func(__e *ControlFlow) {
V6466 := __e.Get(1)
_ = V6466
__e.Return(MakeNative(func(__e *ControlFlow) {
B6462 := __e.Get(1)
_ = B6462
__e.Return(MakeNative(func(__e *ControlFlow) {
L6463 := __e.Get(1)
_ = L6463
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6464 := __e.Get(1)
_ = Key6464
__e.Return(MakeNative(func(__e *ControlFlow) {
C6465 := __e.Get(1)
_ = C6465
tmp20645 := Call(__e, PrimFunc(symshen_4newpv), B6462)


let__18949 := tmp20645
_ = let__18949

tmp20646 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18949, Nil)
}
__typedArg0 := let__18949
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20647 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20646)
}
__typedArg0 := symlist
__typedArg1 := tmp20646
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20648 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18949, Nil)
}
__typedArg0 := let__18949
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20649 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20648)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20648
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20650 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20647, tmp20649)
}
__typedArg0 := tmp20647
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
return PrimCons(symnumber, tmp20652)
}
__typedArg0 := symnumber
__typedArg1 := tmp20652
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20654 := Call(__e, PrimFunc(symis_b), V6466, tmp20653, B6462, L6463, Key6464, C6465)


__e.TailApply(PrimFunc(symshen_4gc), B6462, tmp20654)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20655 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20656 := Call(__e, PrimFunc(symshen_4assoc_1_6), symnth, tmp20644, tmp20655)


tmp20657 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20656)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20656
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20657

tmp20658 := MakeNative(func(__e *ControlFlow) {
V6471 := __e.Get(1)
_ = V6471
__e.Return(MakeNative(func(__e *ControlFlow) {
B6467 := __e.Get(1)
_ = B6467
__e.Return(MakeNative(func(__e *ControlFlow) {
L6468 := __e.Get(1)
_ = L6468
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6469 := __e.Get(1)
_ = Key6469
__e.Return(MakeNative(func(__e *ControlFlow) {
C6470 := __e.Get(1)
_ = C6470
tmp20659 := Call(__e, PrimFunc(symshen_4newpv), B6467)


let__18950 := tmp20659
_ = let__18950

tmp20660 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
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
return PrimCons(let__18950, tmp20661)
}
__typedArg0 := let__18950
__typedArg1 := tmp20661
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20663 := Call(__e, PrimFunc(symis_b), V6471, tmp20662, B6467, L6468, Key6469, C6470)


__e.TailApply(PrimFunc(symshen_4gc), B6467, tmp20663)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20664 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20665 := Call(__e, PrimFunc(symshen_4assoc_1_6), symnumber_2, tmp20658, tmp20664)


tmp20666 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20665)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20665
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20666

tmp20667 := MakeNative(func(__e *ControlFlow) {
V6476 := __e.Get(1)
_ = V6476
__e.Return(MakeNative(func(__e *ControlFlow) {
B6472 := __e.Get(1)
_ = B6472
__e.Return(MakeNative(func(__e *ControlFlow) {
L6473 := __e.Get(1)
_ = L6473
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6474 := __e.Get(1)
_ = Key6474
__e.Return(MakeNative(func(__e *ControlFlow) {
C6475 := __e.Get(1)
_ = C6475
tmp20668 := Call(__e, PrimFunc(symshen_4newpv), B6472)


let__18951 := tmp20668
_ = let__18951

tmp20669 := Call(__e, PrimFunc(symshen_4newpv), B6472)


let__18952 := tmp20669
_ = let__18952

tmp20670 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20671 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20670)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20670
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20672 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18952, tmp20671)
}
__typedArg0 := let__18952
__typedArg1 := tmp20671
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20673 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20672, Nil)
}
__typedArg0 := tmp20672
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20674 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20673)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20673
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20675 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18951, tmp20674)
}
__typedArg0 := let__18951
__typedArg1 := tmp20674
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20676 := Call(__e, PrimFunc(symis_b), V6476, tmp20675, B6472, L6473, Key6474, C6475)


tmp20677 := Call(__e, PrimFunc(symshen_4gc), B6472, tmp20676)


__e.TailApply(PrimFunc(symshen_4gc), B6472, tmp20677)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20678 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20679 := Call(__e, PrimFunc(symshen_4assoc_1_6), symoccurrences, tmp20667, tmp20678)


tmp20680 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20679)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20679
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20680

tmp20681 := MakeNative(func(__e *ControlFlow) {
V6481 := __e.Get(1)
_ = V6481
__e.Return(MakeNative(func(__e *ControlFlow) {
B6477 := __e.Get(1)
_ = B6477
__e.Return(MakeNative(func(__e *ControlFlow) {
L6478 := __e.Get(1)
_ = L6478
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6479 := __e.Get(1)
_ = Key6479
__e.Return(MakeNative(func(__e *ControlFlow) {
C6480 := __e.Get(1)
_ = C6480
tmp20682 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20683 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20682)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20682
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20684 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp20683)
}
__typedArg0 := symsymbol
__typedArg1 := tmp20683
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6481, tmp20684, B6477, L6478, Key6479, C6480)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20685 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20686 := Call(__e, PrimFunc(symshen_4assoc_1_6), symoccurs_1check, tmp20681, tmp20685)


tmp20687 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20686)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20686
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20687

tmp20688 := MakeNative(func(__e *ControlFlow) {
V6486 := __e.Get(1)
_ = V6486
__e.Return(MakeNative(func(__e *ControlFlow) {
B6482 := __e.Get(1)
_ = B6482
__e.Return(MakeNative(func(__e *ControlFlow) {
L6483 := __e.Get(1)
_ = L6483
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6484 := __e.Get(1)
_ = Key6484
__e.Return(MakeNative(func(__e *ControlFlow) {
C6485 := __e.Get(1)
_ = C6485
tmp20689 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20690 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20689)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20689
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6486, tmp20690, B6482, L6483, Key6484, C6485)
return


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

tmp20692 := Call(__e, PrimFunc(symshen_4assoc_1_6), symoccurs_2, tmp20688, tmp20691)


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
V6491 := __e.Get(1)
_ = V6491
__e.Return(MakeNative(func(__e *ControlFlow) {
B6487 := __e.Get(1)
_ = B6487
__e.Return(MakeNative(func(__e *ControlFlow) {
L6488 := __e.Get(1)
_ = L6488
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6489 := __e.Get(1)
_ = Key6489
__e.Return(MakeNative(func(__e *ControlFlow) {
C6490 := __e.Get(1)
_ = C6490
tmp20695 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20696 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20695)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20695
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20697 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp20696)
}
__typedArg0 := symsymbol
__typedArg1 := tmp20696
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6491, tmp20697, B6487, L6488, Key6489, C6490)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20698 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20699 := Call(__e, PrimFunc(symshen_4assoc_1_6), symoptimise, tmp20694, tmp20698)


tmp20700 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20699)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20699
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20700

tmp20701 := MakeNative(func(__e *ControlFlow) {
V6496 := __e.Get(1)
_ = V6496
__e.Return(MakeNative(func(__e *ControlFlow) {
B6492 := __e.Get(1)
_ = B6492
__e.Return(MakeNative(func(__e *ControlFlow) {
L6493 := __e.Get(1)
_ = L6493
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6494 := __e.Get(1)
_ = Key6494
__e.Return(MakeNative(func(__e *ControlFlow) {
C6495 := __e.Get(1)
_ = C6495
tmp20702 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20703 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20702)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20702
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6496, tmp20703, B6492, L6493, Key6494, C6495)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20704 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20705 := Call(__e, PrimFunc(symshen_4assoc_1_6), symoptimise_2, tmp20701, tmp20704)


tmp20706 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20705)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20705
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20706

tmp20707 := MakeNative(func(__e *ControlFlow) {
V6501 := __e.Get(1)
_ = V6501
__e.Return(MakeNative(func(__e *ControlFlow) {
B6497 := __e.Get(1)
_ = B6497
__e.Return(MakeNative(func(__e *ControlFlow) {
L6498 := __e.Get(1)
_ = L6498
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6499 := __e.Get(1)
_ = Key6499
__e.Return(MakeNative(func(__e *ControlFlow) {
C6500 := __e.Get(1)
_ = C6500
tmp20708 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20709 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20708)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20708
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20710 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, tmp20709)
}
__typedArg0 := symboolean
__typedArg1 := tmp20709
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20711 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20710, Nil)
}
__typedArg0 := tmp20710
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20712 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20711)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20711
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20713 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, tmp20712)
}
__typedArg0 := symboolean
__typedArg1 := tmp20712
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6501, tmp20713, B6497, L6498, Key6499, C6500)
return


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

tmp20715 := Call(__e, PrimFunc(symshen_4assoc_1_6), symor, tmp20707, tmp20714)


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
V6506 := __e.Get(1)
_ = V6506
__e.Return(MakeNative(func(__e *ControlFlow) {
B6502 := __e.Get(1)
_ = B6502
__e.Return(MakeNative(func(__e *ControlFlow) {
L6503 := __e.Get(1)
_ = L6503
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6504 := __e.Get(1)
_ = Key6504
__e.Return(MakeNative(func(__e *ControlFlow) {
C6505 := __e.Get(1)
_ = C6505
tmp20718 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
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

__e.TailApply(PrimFunc(symis_b), V6506, tmp20719, B6502, L6503, Key6504, C6505)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20720 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20721 := Call(__e, PrimFunc(symshen_4assoc_1_6), symos, tmp20717, tmp20720)


tmp20722 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20721)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20721
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20722

tmp20723 := MakeNative(func(__e *ControlFlow) {
V6511 := __e.Get(1)
_ = V6511
__e.Return(MakeNative(func(__e *ControlFlow) {
B6507 := __e.Get(1)
_ = B6507
__e.Return(MakeNative(func(__e *ControlFlow) {
L6508 := __e.Get(1)
_ = L6508
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6509 := __e.Get(1)
_ = Key6509
__e.Return(MakeNative(func(__e *ControlFlow) {
C6510 := __e.Get(1)
_ = C6510
tmp20724 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20725 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20724)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20724
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20726 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp20725)
}
__typedArg0 := symsymbol
__typedArg1 := tmp20725
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6511, tmp20726, B6507, L6508, Key6509, C6510)
return


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

tmp20728 := Call(__e, PrimFunc(symshen_4assoc_1_6), sympackage_2, tmp20723, tmp20727)


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
V6516 := __e.Get(1)
_ = V6516
__e.Return(MakeNative(func(__e *ControlFlow) {
B6512 := __e.Get(1)
_ = B6512
__e.Return(MakeNative(func(__e *ControlFlow) {
L6513 := __e.Get(1)
_ = L6513
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6514 := __e.Get(1)
_ = Key6514
__e.Return(MakeNative(func(__e *ControlFlow) {
C6515 := __e.Get(1)
_ = C6515
tmp20731 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
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

__e.TailApply(PrimFunc(symis_b), V6516, tmp20732, B6512, L6513, Key6514, C6515)
return


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

tmp20734 := Call(__e, PrimFunc(symshen_4assoc_1_6), symport, tmp20730, tmp20733)


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
V6521 := __e.Get(1)
_ = V6521
__e.Return(MakeNative(func(__e *ControlFlow) {
B6517 := __e.Get(1)
_ = B6517
__e.Return(MakeNative(func(__e *ControlFlow) {
L6518 := __e.Get(1)
_ = L6518
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6519 := __e.Get(1)
_ = Key6519
__e.Return(MakeNative(func(__e *ControlFlow) {
C6520 := __e.Get(1)
_ = C6520
tmp20737 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20738 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20737)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20737
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6521, tmp20738, B6517, L6518, Key6519, C6520)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20739 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20740 := Call(__e, PrimFunc(symshen_4assoc_1_6), symporters, tmp20736, tmp20739)


tmp20741 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20740)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20740
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20741

tmp20742 := MakeNative(func(__e *ControlFlow) {
V6526 := __e.Get(1)
_ = V6526
__e.Return(MakeNative(func(__e *ControlFlow) {
B6522 := __e.Get(1)
_ = B6522
__e.Return(MakeNative(func(__e *ControlFlow) {
L6523 := __e.Get(1)
_ = L6523
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6524 := __e.Get(1)
_ = Key6524
__e.Return(MakeNative(func(__e *ControlFlow) {
C6525 := __e.Get(1)
_ = C6525
tmp20743 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20744 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20743)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20743
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20745 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp20744)
}
__typedArg0 := symnumber
__typedArg1 := tmp20744
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20746 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20745, Nil)
}
__typedArg0 := tmp20745
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20747 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20746)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20746
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20748 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp20747)
}
__typedArg0 := symstring
__typedArg1 := tmp20747
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6526, tmp20748, B6522, L6523, Key6524, C6525)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20749 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20750 := Call(__e, PrimFunc(symshen_4assoc_1_6), sympos, tmp20742, tmp20749)


tmp20751 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20750)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20750
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20751

tmp20752 := MakeNative(func(__e *ControlFlow) {
V6531 := __e.Get(1)
_ = V6531
__e.Return(MakeNative(func(__e *ControlFlow) {
B6527 := __e.Get(1)
_ = B6527
__e.Return(MakeNative(func(__e *ControlFlow) {
L6528 := __e.Get(1)
_ = L6528
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6529 := __e.Get(1)
_ = Key6529
__e.Return(MakeNative(func(__e *ControlFlow) {
C6530 := __e.Get(1)
_ = C6530
tmp20753 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symout, Nil)
}
__typedArg0 := symout
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20754 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp20753)
}
__typedArg0 := symstream
__typedArg1 := tmp20753
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20755 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20756 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20755)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20755
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20757 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20754, tmp20756)
}
__typedArg0 := tmp20754
__typedArg1 := tmp20756
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20758 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20757, Nil)
}
__typedArg0 := tmp20757
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

tmp20760 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp20759)
}
__typedArg0 := symstring
__typedArg1 := tmp20759
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6531, tmp20760, B6527, L6528, Key6529, C6530)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20761 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20762 := Call(__e, PrimFunc(symshen_4assoc_1_6), sympr, tmp20752, tmp20761)


tmp20763 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20762)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20762
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20763

tmp20764 := MakeNative(func(__e *ControlFlow) {
V6536 := __e.Get(1)
_ = V6536
__e.Return(MakeNative(func(__e *ControlFlow) {
B6532 := __e.Get(1)
_ = B6532
__e.Return(MakeNative(func(__e *ControlFlow) {
L6533 := __e.Get(1)
_ = L6533
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6534 := __e.Get(1)
_ = Key6534
__e.Return(MakeNative(func(__e *ControlFlow) {
C6535 := __e.Get(1)
_ = C6535
tmp20765 := Call(__e, PrimFunc(symshen_4newpv), B6532)


let__18953 := tmp20765
_ = let__18953

tmp20766 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18953, Nil)
}
__typedArg0 := let__18953
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20767 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20766)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20766
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20768 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18953, tmp20767)
}
__typedArg0 := let__18953
__typedArg1 := tmp20767
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20769 := Call(__e, PrimFunc(symis_b), V6536, tmp20768, B6532, L6533, Key6534, C6535)


__e.TailApply(PrimFunc(symshen_4gc), B6532, tmp20769)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20770 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20771 := Call(__e, PrimFunc(symshen_4assoc_1_6), symprint, tmp20764, tmp20770)


tmp20772 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20771)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20771
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20772

tmp20773 := MakeNative(func(__e *ControlFlow) {
V6541 := __e.Get(1)
_ = V6541
__e.Return(MakeNative(func(__e *ControlFlow) {
B6537 := __e.Get(1)
_ = B6537
__e.Return(MakeNative(func(__e *ControlFlow) {
L6538 := __e.Get(1)
_ = L6538
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6539 := __e.Get(1)
_ = Key6539
__e.Return(MakeNative(func(__e *ControlFlow) {
C6540 := __e.Get(1)
_ = C6540
tmp20774 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20775 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20774)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20774
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20776 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp20775)
}
__typedArg0 := symsymbol
__typedArg1 := tmp20775
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6541, tmp20776, B6537, L6538, Key6539, C6540)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20777 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20778 := Call(__e, PrimFunc(symshen_4assoc_1_6), symprofile, tmp20773, tmp20777)


tmp20779 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20778)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20778
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20779

tmp20780 := MakeNative(func(__e *ControlFlow) {
V6546 := __e.Get(1)
_ = V6546
__e.Return(MakeNative(func(__e *ControlFlow) {
B6542 := __e.Get(1)
_ = B6542
__e.Return(MakeNative(func(__e *ControlFlow) {
L6543 := __e.Get(1)
_ = L6543
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6544 := __e.Get(1)
_ = Key6544
__e.Return(MakeNative(func(__e *ControlFlow) {
C6545 := __e.Get(1)
_ = C6545
tmp20781 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20782 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20781)
}
__typedArg0 := symlist
__typedArg1 := tmp20781
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20783 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20784 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20783)
}
__typedArg0 := symlist
__typedArg1 := tmp20783
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20785 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20784, Nil)
}
__typedArg0 := tmp20784
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20786 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20785)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20785
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20787 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20782, tmp20786)
}
__typedArg0 := tmp20782
__typedArg1 := tmp20786
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6546, tmp20787, B6542, L6543, Key6544, C6545)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20788 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20789 := Call(__e, PrimFunc(symshen_4assoc_1_6), sympreclude, tmp20780, tmp20788)


tmp20790 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20789)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20789
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20790

tmp20791 := MakeNative(func(__e *ControlFlow) {
V6551 := __e.Get(1)
_ = V6551
__e.Return(MakeNative(func(__e *ControlFlow) {
B6547 := __e.Get(1)
_ = B6547
__e.Return(MakeNative(func(__e *ControlFlow) {
L6548 := __e.Get(1)
_ = L6548
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6549 := __e.Get(1)
_ = Key6549
__e.Return(MakeNative(func(__e *ControlFlow) {
C6550 := __e.Get(1)
_ = C6550
tmp20792 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20793 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20792)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20792
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20794 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp20793)
}
__typedArg0 := symstring
__typedArg1 := tmp20793
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6551, tmp20794, B6547, L6548, Key6549, C6550)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20795 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20796 := Call(__e, PrimFunc(symshen_4assoc_1_6), symshen_4proc_1nl, tmp20791, tmp20795)


tmp20797 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20796)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20796
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20797

tmp20798 := MakeNative(func(__e *ControlFlow) {
V6556 := __e.Get(1)
_ = V6556
__e.Return(MakeNative(func(__e *ControlFlow) {
B6552 := __e.Get(1)
_ = B6552
__e.Return(MakeNative(func(__e *ControlFlow) {
L6553 := __e.Get(1)
_ = L6553
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6554 := __e.Get(1)
_ = Key6554
__e.Return(MakeNative(func(__e *ControlFlow) {
C6555 := __e.Get(1)
_ = C6555
tmp20799 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20800 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_d, tmp20799)
}
__typedArg0 := sym_d
__typedArg1 := tmp20799
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20801 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp20800)
}
__typedArg0 := symsymbol
__typedArg1 := tmp20800
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20802 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20801, Nil)
}
__typedArg0 := tmp20801
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
return PrimCons(symsymbol, tmp20803)
}
__typedArg0 := symsymbol
__typedArg1 := tmp20803
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6556, tmp20804, B6552, L6553, Key6554, C6555)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20805 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20806 := Call(__e, PrimFunc(symshen_4assoc_1_6), symprofile_1results, tmp20798, tmp20805)


tmp20807 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20806)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20806
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20807

tmp20808 := MakeNative(func(__e *ControlFlow) {
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
tmp20809 := Call(__e, PrimFunc(symshen_4newpv), B6557)


let__18954 := tmp20809
_ = let__18954

tmp20810 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18954, Nil)
}
__typedArg0 := let__18954
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
return PrimCons(let__18954, tmp20811)
}
__typedArg0 := let__18954
__typedArg1 := tmp20811
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20813 := Call(__e, PrimFunc(symis_b), V6561, tmp20812, B6557, L6558, Key6559, C6560)


__e.TailApply(PrimFunc(symshen_4gc), B6557, tmp20813)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20814 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20815 := Call(__e, PrimFunc(symshen_4assoc_1_6), symprotect, tmp20808, tmp20814)


tmp20816 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20815)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20815
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20816

tmp20817 := MakeNative(func(__e *ControlFlow) {
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
tmp20818 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20819 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20818)
}
__typedArg0 := symlist
__typedArg1 := tmp20818
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20820 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20821 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20820)
}
__typedArg0 := symlist
__typedArg1 := tmp20820
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20822 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20821, Nil)
}
__typedArg0 := tmp20821
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20823 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20822)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20822
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20824 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20819, tmp20823)
}
__typedArg0 := tmp20819
__typedArg1 := tmp20823
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6566, tmp20824, B6562, L6563, Key6564, C6565)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20825 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20826 := Call(__e, PrimFunc(symshen_4assoc_1_6), sympreclude_1all_1but, tmp20817, tmp20825)


tmp20827 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20826)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20826
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20827

tmp20828 := MakeNative(func(__e *ControlFlow) {
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
tmp20829 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symout, Nil)
}
__typedArg0 := symout
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20830 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp20829)
}
__typedArg0 := symstream
__typedArg1 := tmp20829
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20831 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20832 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20831)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20831
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20833 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20830, tmp20832)
}
__typedArg0 := tmp20830
__typedArg1 := tmp20832
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20834 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20833, Nil)
}
__typedArg0 := tmp20833
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20835 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20834)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20834
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20836 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp20835)
}
__typedArg0 := symstring
__typedArg1 := tmp20835
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6571, tmp20836, B6567, L6568, Key6569, C6570)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20837 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20838 := Call(__e, PrimFunc(symshen_4assoc_1_6), symshen_4prhush, tmp20828, tmp20837)


tmp20839 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20838)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20838
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20839

tmp20840 := MakeNative(func(__e *ControlFlow) {
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
tmp20841 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20842 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20841)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20841
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20843 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp20842)
}
__typedArg0 := symnumber
__typedArg1 := tmp20842
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6576, tmp20843, B6572, L6573, Key6574, C6575)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20844 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20845 := Call(__e, PrimFunc(symshen_4assoc_1_6), symprolog_1memory, tmp20840, tmp20844)


tmp20846 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20845)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20845
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20846

tmp20847 := MakeNative(func(__e *ControlFlow) {
V6581 := __e.Get(1)
_ = V6581
__e.Return(MakeNative(func(__e *ControlFlow) {
B6577 := __e.Get(1)
_ = B6577
__e.Return(MakeNative(func(__e *ControlFlow) {
L6578 := __e.Get(1)
_ = L6578
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6579 := __e.Get(1)
_ = Key6579
__e.Return(MakeNative(func(__e *ControlFlow) {
C6580 := __e.Get(1)
_ = C6580
tmp20848 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunit, Nil)
}
__typedArg0 := symunit
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20849 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20848)
}
__typedArg0 := symlist
__typedArg1 := tmp20848
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20850 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20849, Nil)
}
__typedArg0 := tmp20849
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20851 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20850)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20850
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20852 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp20851)
}
__typedArg0 := symsymbol
__typedArg1 := tmp20851
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6581, tmp20852, B6577, L6578, Key6579, C6580)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20853 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20854 := Call(__e, PrimFunc(symshen_4assoc_1_6), symps, tmp20847, tmp20853)


tmp20855 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20854)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20854
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20855

tmp20856 := MakeNative(func(__e *ControlFlow) {
V6586 := __e.Get(1)
_ = V6586
__e.Return(MakeNative(func(__e *ControlFlow) {
B6582 := __e.Get(1)
_ = B6582
__e.Return(MakeNative(func(__e *ControlFlow) {
L6583 := __e.Get(1)
_ = L6583
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6584 := __e.Get(1)
_ = Key6584
__e.Return(MakeNative(func(__e *ControlFlow) {
C6585 := __e.Get(1)
_ = C6585
tmp20857 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symin, Nil)
}
__typedArg0 := symin
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20858 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp20857)
}
__typedArg0 := symstream
__typedArg1 := tmp20857
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20859 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunit, Nil)
}
__typedArg0 := symunit
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20860 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20859)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20859
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20861 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20858, tmp20860)
}
__typedArg0 := tmp20858
__typedArg1 := tmp20860
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6586, tmp20861, B6582, L6583, Key6584, C6585)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20862 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20863 := Call(__e, PrimFunc(symshen_4assoc_1_6), symread, tmp20856, tmp20862)


tmp20864 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20863)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20863
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20864

tmp20865 := MakeNative(func(__e *ControlFlow) {
V6591 := __e.Get(1)
_ = V6591
__e.Return(MakeNative(func(__e *ControlFlow) {
B6587 := __e.Get(1)
_ = B6587
__e.Return(MakeNative(func(__e *ControlFlow) {
L6588 := __e.Get(1)
_ = L6588
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6589 := __e.Get(1)
_ = Key6589
__e.Return(MakeNative(func(__e *ControlFlow) {
C6590 := __e.Get(1)
_ = C6590
tmp20866 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symin, Nil)
}
__typedArg0 := symin
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20867 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp20866)
}
__typedArg0 := symstream
__typedArg1 := tmp20866
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

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
return PrimCons(sym_1_1_6, tmp20868)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20868
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20870 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20867, tmp20869)
}
__typedArg0 := tmp20867
__typedArg1 := tmp20869
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6591, tmp20870, B6587, L6588, Key6589, C6590)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20871 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20872 := Call(__e, PrimFunc(symshen_4assoc_1_6), symread_1byte, tmp20865, tmp20871)


tmp20873 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20872)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20872
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20873

tmp20874 := MakeNative(func(__e *ControlFlow) {
V6596 := __e.Get(1)
_ = V6596
__e.Return(MakeNative(func(__e *ControlFlow) {
B6592 := __e.Get(1)
_ = B6592
__e.Return(MakeNative(func(__e *ControlFlow) {
L6593 := __e.Get(1)
_ = L6593
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6594 := __e.Get(1)
_ = Key6594
__e.Return(MakeNative(func(__e *ControlFlow) {
C6595 := __e.Get(1)
_ = C6595
tmp20875 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20876 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20875)
}
__typedArg0 := symlist
__typedArg1 := tmp20875
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20877 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20876, Nil)
}
__typedArg0 := tmp20876
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
return PrimCons(symstring, tmp20878)
}
__typedArg0 := symstring
__typedArg1 := tmp20878
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6596, tmp20879, B6592, L6593, Key6594, C6595)
return


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

tmp20881 := Call(__e, PrimFunc(symshen_4assoc_1_6), symread_1file_1as_1bytelist, tmp20874, tmp20880)


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
V6601 := __e.Get(1)
_ = V6601
__e.Return(MakeNative(func(__e *ControlFlow) {
B6597 := __e.Get(1)
_ = B6597
__e.Return(MakeNative(func(__e *ControlFlow) {
L6598 := __e.Get(1)
_ = L6598
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6599 := __e.Get(1)
_ = Key6599
__e.Return(MakeNative(func(__e *ControlFlow) {
C6600 := __e.Get(1)
_ = C6600
tmp20884 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20885 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20884)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20884
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20886 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp20885)
}
__typedArg0 := symstring
__typedArg1 := tmp20885
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6601, tmp20886, B6597, L6598, Key6599, C6600)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20887 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20888 := Call(__e, PrimFunc(symshen_4assoc_1_6), symread_1file_1as_1string, tmp20883, tmp20887)


tmp20889 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20888)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20888
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20889

tmp20890 := MakeNative(func(__e *ControlFlow) {
V6606 := __e.Get(1)
_ = V6606
__e.Return(MakeNative(func(__e *ControlFlow) {
B6602 := __e.Get(1)
_ = B6602
__e.Return(MakeNative(func(__e *ControlFlow) {
L6603 := __e.Get(1)
_ = L6603
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6604 := __e.Get(1)
_ = Key6604
__e.Return(MakeNative(func(__e *ControlFlow) {
C6605 := __e.Get(1)
_ = C6605
tmp20891 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunit, Nil)
}
__typedArg0 := symunit
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20892 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20891)
}
__typedArg0 := symlist
__typedArg1 := tmp20891
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20893 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20892, Nil)
}
__typedArg0 := tmp20892
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20894 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20893)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20893
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20895 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp20894)
}
__typedArg0 := symstring
__typedArg1 := tmp20894
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6606, tmp20895, B6602, L6603, Key6604, C6605)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20896 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20897 := Call(__e, PrimFunc(symshen_4assoc_1_6), symread_1file, tmp20890, tmp20896)


tmp20898 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20897)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20897
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20898

tmp20899 := MakeNative(func(__e *ControlFlow) {
V6611 := __e.Get(1)
_ = V6611
__e.Return(MakeNative(func(__e *ControlFlow) {
B6607 := __e.Get(1)
_ = B6607
__e.Return(MakeNative(func(__e *ControlFlow) {
L6608 := __e.Get(1)
_ = L6608
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6609 := __e.Get(1)
_ = Key6609
__e.Return(MakeNative(func(__e *ControlFlow) {
C6610 := __e.Get(1)
_ = C6610
tmp20900 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunit, Nil)
}
__typedArg0 := symunit
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20901 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20900)
}
__typedArg0 := symlist
__typedArg1 := tmp20900
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20902 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20901, Nil)
}
__typedArg0 := tmp20901
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20903 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20902)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20902
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20904 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp20903)
}
__typedArg0 := symstring
__typedArg1 := tmp20903
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6611, tmp20904, B6607, L6608, Key6609, C6610)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20905 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20906 := Call(__e, PrimFunc(symshen_4assoc_1_6), symread_1from_1string, tmp20899, tmp20905)


tmp20907 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20906)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20906
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20907

tmp20908 := MakeNative(func(__e *ControlFlow) {
V6616 := __e.Get(1)
_ = V6616
__e.Return(MakeNative(func(__e *ControlFlow) {
B6612 := __e.Get(1)
_ = B6612
__e.Return(MakeNative(func(__e *ControlFlow) {
L6613 := __e.Get(1)
_ = L6613
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6614 := __e.Get(1)
_ = Key6614
__e.Return(MakeNative(func(__e *ControlFlow) {
C6615 := __e.Get(1)
_ = C6615
tmp20909 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunit, Nil)
}
__typedArg0 := symunit
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20910 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20909)
}
__typedArg0 := symlist
__typedArg1 := tmp20909
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20911 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20910, Nil)
}
__typedArg0 := tmp20910
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20912 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20911)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20911
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20913 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp20912)
}
__typedArg0 := symstring
__typedArg1 := tmp20912
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6616, tmp20913, B6612, L6613, Key6614, C6615)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20914 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20915 := Call(__e, PrimFunc(symshen_4assoc_1_6), symread_1from_1string_1unprocessed, tmp20908, tmp20914)


tmp20916 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20915)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20915
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20916

tmp20917 := MakeNative(func(__e *ControlFlow) {
V6621 := __e.Get(1)
_ = V6621
__e.Return(MakeNative(func(__e *ControlFlow) {
B6617 := __e.Get(1)
_ = B6617
__e.Return(MakeNative(func(__e *ControlFlow) {
L6618 := __e.Get(1)
_ = L6618
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6619 := __e.Get(1)
_ = Key6619
__e.Return(MakeNative(func(__e *ControlFlow) {
C6620 := __e.Get(1)
_ = C6620
tmp20918 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20919 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20918)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20918
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6621, tmp20919, B6617, L6618, Key6619, C6620)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20920 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20921 := Call(__e, PrimFunc(symshen_4assoc_1_6), symrelease, tmp20917, tmp20920)


tmp20922 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20921)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20921
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20922

tmp20923 := MakeNative(func(__e *ControlFlow) {
V6626 := __e.Get(1)
_ = V6626
__e.Return(MakeNative(func(__e *ControlFlow) {
B6622 := __e.Get(1)
_ = B6622
__e.Return(MakeNative(func(__e *ControlFlow) {
L6623 := __e.Get(1)
_ = L6623
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6624 := __e.Get(1)
_ = Key6624
__e.Return(MakeNative(func(__e *ControlFlow) {
C6625 := __e.Get(1)
_ = C6625
tmp20924 := Call(__e, PrimFunc(symshen_4newpv), B6622)


let__18955 := tmp20924
_ = let__18955

tmp20925 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18955, Nil)
}
__typedArg0 := let__18955
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20926 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20925)
}
__typedArg0 := symlist
__typedArg1 := tmp20925
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20927 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18955, Nil)
}
__typedArg0 := let__18955
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20928 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20927)
}
__typedArg0 := symlist
__typedArg1 := tmp20927
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20929 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20928, Nil)
}
__typedArg0 := tmp20928
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
return PrimCons(tmp20926, tmp20930)
}
__typedArg0 := tmp20926
__typedArg1 := tmp20930
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20932 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20931, Nil)
}
__typedArg0 := tmp20931
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20933 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20932)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20932
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20934 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18955, tmp20933)
}
__typedArg0 := let__18955
__typedArg1 := tmp20933
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20935 := Call(__e, PrimFunc(symis_b), V6626, tmp20934, B6622, L6623, Key6624, C6625)


__e.TailApply(PrimFunc(symshen_4gc), B6622, tmp20935)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20936 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20937 := Call(__e, PrimFunc(symshen_4assoc_1_6), symremove, tmp20923, tmp20936)


tmp20938 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20937)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20937
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20938

tmp20939 := MakeNative(func(__e *ControlFlow) {
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
tmp20940 := Call(__e, PrimFunc(symshen_4newpv), B6627)


let__18956 := tmp20940
_ = let__18956

tmp20941 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18956, Nil)
}
__typedArg0 := let__18956
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20942 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20941)
}
__typedArg0 := symlist
__typedArg1 := tmp20941
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20943 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18956, Nil)
}
__typedArg0 := let__18956
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20944 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20943)
}
__typedArg0 := symlist
__typedArg1 := tmp20943
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20945 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20944, Nil)
}
__typedArg0 := tmp20944
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20946 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20945)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20945
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20947 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20942, tmp20946)
}
__typedArg0 := tmp20942
__typedArg1 := tmp20946
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20948 := Call(__e, PrimFunc(symis_b), V6631, tmp20947, B6627, L6628, Key6629, C6630)


__e.TailApply(PrimFunc(symshen_4gc), B6627, tmp20948)
return


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

tmp20950 := Call(__e, PrimFunc(symshen_4assoc_1_6), symreverse, tmp20939, tmp20949)


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
tmp20953 := Call(__e, PrimFunc(symshen_4newpv), B6632)


let__18957 := tmp20953
_ = let__18957

tmp20954 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18957, Nil)
}
__typedArg0 := let__18957
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20955 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20954)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20954
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20956 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp20955)
}
__typedArg0 := symstring
__typedArg1 := tmp20955
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20957 := Call(__e, PrimFunc(symis_b), V6636, tmp20956, B6632, L6633, Key6634, C6635)


__e.TailApply(PrimFunc(symshen_4gc), B6632, tmp20957)
return


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

tmp20959 := Call(__e, PrimFunc(symshen_4assoc_1_6), symsimple_1error, tmp20952, tmp20958)


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
tmp20962 := Call(__e, PrimFunc(symshen_4newpv), B6637)


let__18958 := tmp20962
_ = let__18958

tmp20963 := Call(__e, PrimFunc(symshen_4newpv), B6637)


let__18959 := tmp20963
_ = let__18959

tmp20964 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18959, Nil)
}
__typedArg0 := let__18959
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20965 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_d, tmp20964)
}
__typedArg0 := sym_d
__typedArg1 := tmp20964
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20966 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18958, tmp20965)
}
__typedArg0 := let__18958
__typedArg1 := tmp20965
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20967 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18959, Nil)
}
__typedArg0 := let__18959
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20968 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20967)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20967
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20969 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20966, tmp20968)
}
__typedArg0 := tmp20966
__typedArg1 := tmp20968
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20970 := Call(__e, PrimFunc(symis_b), V6641, tmp20969, B6637, L6638, Key6639, C6640)


tmp20971 := Call(__e, PrimFunc(symshen_4gc), B6637, tmp20970)


__e.TailApply(PrimFunc(symshen_4gc), B6637, tmp20971)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20972 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20973 := Call(__e, PrimFunc(symshen_4assoc_1_6), symsnd, tmp20961, tmp20972)


tmp20974 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20973)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20973
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20974

tmp20975 := MakeNative(func(__e *ControlFlow) {
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
tmp20976 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20977 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20976)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20976
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20978 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp20977)
}
__typedArg0 := symnumber
__typedArg1 := tmp20977
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20979 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20978, Nil)
}
__typedArg0 := tmp20978
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20980 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20979)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20979
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20981 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp20980)
}
__typedArg0 := symsymbol
__typedArg1 := tmp20980
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6646, tmp20981, B6642, L6643, Key6644, C6645)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20982 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20983 := Call(__e, PrimFunc(symshen_4assoc_1_6), symspecialise, tmp20975, tmp20982)


tmp20984 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20983)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20983
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20984

tmp20985 := MakeNative(func(__e *ControlFlow) {
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
tmp20986 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20987 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20986)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20986
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20988 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp20987)
}
__typedArg0 := symsymbol
__typedArg1 := tmp20987
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6651, tmp20988, B6647, L6648, Key6649, C6650)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20989 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20990 := Call(__e, PrimFunc(symshen_4assoc_1_6), symspy, tmp20985, tmp20989)


tmp20991 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20990)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20990
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20991

tmp20992 := MakeNative(func(__e *ControlFlow) {
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
tmp20993 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20994 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20993)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20993
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6656, tmp20994, B6652, L6653, Key6654, C6655)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20995 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20996 := Call(__e, PrimFunc(symshen_4assoc_1_6), symshen_4spy_2, tmp20992, tmp20995)


tmp20997 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20996)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20996
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20997

tmp20998 := MakeNative(func(__e *ControlFlow) {
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
tmp20999 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
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
return PrimCons(symsymbol, tmp21000)
}
__typedArg0 := symsymbol
__typedArg1 := tmp21000
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6661, tmp21001, B6657, L6658, Key6659, C6660)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21002 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21003 := Call(__e, PrimFunc(symshen_4assoc_1_6), symstep, tmp20998, tmp21002)


tmp21004 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21003)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21003
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21004

tmp21005 := MakeNative(func(__e *ControlFlow) {
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
tmp21006 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21007 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21006)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21006
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6666, tmp21007, B6662, L6663, Key6664, C6665)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21008 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21009 := Call(__e, PrimFunc(symshen_4assoc_1_6), symshen_4step_2, tmp21005, tmp21008)


tmp21010 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21009)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21009
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21010

tmp21011 := MakeNative(func(__e *ControlFlow) {
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
tmp21012 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symin, Nil)
}
__typedArg0 := symin
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21013 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp21012)
}
__typedArg0 := symstream
__typedArg1 := tmp21012
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21014 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21013, Nil)
}
__typedArg0 := tmp21013
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21015 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21014)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21014
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6671, tmp21015, B6667, L6668, Key6669, C6670)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21016 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21017 := Call(__e, PrimFunc(symshen_4assoc_1_6), symstinput, tmp21011, tmp21016)


tmp21018 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21017)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21017
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21018

tmp21019 := MakeNative(func(__e *ControlFlow) {
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
tmp21020 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symout, Nil)
}
__typedArg0 := symout
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21021 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp21020)
}
__typedArg0 := symstream
__typedArg1 := tmp21020
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21022 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21021, Nil)
}
__typedArg0 := tmp21021
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21023 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21022)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21022
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6676, tmp21023, B6672, L6673, Key6674, C6675)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21024 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21025 := Call(__e, PrimFunc(symshen_4assoc_1_6), symsterror, tmp21019, tmp21024)


tmp21026 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21025)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21025
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21026

tmp21027 := MakeNative(func(__e *ControlFlow) {
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
tmp21028 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symout, Nil)
}
__typedArg0 := symout
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21029 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp21028)
}
__typedArg0 := symstream
__typedArg1 := tmp21028
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21030 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21029, Nil)
}
__typedArg0 := tmp21029
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21031 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21030)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21030
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6681, tmp21031, B6677, L6678, Key6679, C6680)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21032 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21033 := Call(__e, PrimFunc(symshen_4assoc_1_6), symstoutput, tmp21027, tmp21032)


tmp21034 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21033)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21033
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21034

tmp21035 := MakeNative(func(__e *ControlFlow) {
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
tmp21036 := Call(__e, PrimFunc(symshen_4newpv), B6682)


let__18960 := tmp21036
_ = let__18960

tmp21037 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21038 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21037)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21037
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21039 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18960, tmp21038)
}
__typedArg0 := let__18960
__typedArg1 := tmp21038
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21040 := Call(__e, PrimFunc(symis_b), V6686, tmp21039, B6682, L6683, Key6684, C6685)


__e.TailApply(PrimFunc(symshen_4gc), B6682, tmp21040)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21041 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21042 := Call(__e, PrimFunc(symshen_4assoc_1_6), symstring_2, tmp21035, tmp21041)


tmp21043 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21042)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21042
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21043

tmp21044 := MakeNative(func(__e *ControlFlow) {
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
tmp21045 := Call(__e, PrimFunc(symshen_4newpv), B6687)


let__18961 := tmp21045
_ = let__18961

tmp21046 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21047 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21046)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21046
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21048 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18961, tmp21047)
}
__typedArg0 := let__18961
__typedArg1 := tmp21047
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21049 := Call(__e, PrimFunc(symis_b), V6691, tmp21048, B6687, L6688, Key6689, C6690)


__e.TailApply(PrimFunc(symshen_4gc), B6687, tmp21049)
return


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

tmp21051 := Call(__e, PrimFunc(symshen_4assoc_1_6), symstr, tmp21044, tmp21050)


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
V6696 := __e.Get(1)
_ = V6696
__e.Return(MakeNative(func(__e *ControlFlow) {
B6692 := __e.Get(1)
_ = B6692
__e.Return(MakeNative(func(__e *ControlFlow) {
L6693 := __e.Get(1)
_ = L6693
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6694 := __e.Get(1)
_ = Key6694
__e.Return(MakeNative(func(__e *ControlFlow) {
C6695 := __e.Get(1)
_ = C6695
tmp21054 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
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
return PrimCons(symstring, tmp21055)
}
__typedArg0 := symstring
__typedArg1 := tmp21055
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6696, tmp21056, B6692, L6693, Key6694, C6695)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21057 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21058 := Call(__e, PrimFunc(symshen_4assoc_1_6), symstring_1_6n, tmp21053, tmp21057)


tmp21059 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21058)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21058
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21059

tmp21060 := MakeNative(func(__e *ControlFlow) {
V6701 := __e.Get(1)
_ = V6701
__e.Return(MakeNative(func(__e *ControlFlow) {
B6697 := __e.Get(1)
_ = B6697
__e.Return(MakeNative(func(__e *ControlFlow) {
L6698 := __e.Get(1)
_ = L6698
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6699 := __e.Get(1)
_ = Key6699
__e.Return(MakeNative(func(__e *ControlFlow) {
C6700 := __e.Get(1)
_ = C6700
tmp21061 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21062 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21061)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21061
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21063 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp21062)
}
__typedArg0 := symstring
__typedArg1 := tmp21062
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6701, tmp21063, B6697, L6698, Key6699, C6700)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21064 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21065 := Call(__e, PrimFunc(symshen_4assoc_1_6), symstring_1_6symbol, tmp21060, tmp21064)


tmp21066 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21065)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21065
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21066

tmp21067 := MakeNative(func(__e *ControlFlow) {
V6706 := __e.Get(1)
_ = V6706
__e.Return(MakeNative(func(__e *ControlFlow) {
B6702 := __e.Get(1)
_ = B6702
__e.Return(MakeNative(func(__e *ControlFlow) {
L6703 := __e.Get(1)
_ = L6703
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6704 := __e.Get(1)
_ = Key6704
__e.Return(MakeNative(func(__e *ControlFlow) {
C6705 := __e.Get(1)
_ = C6705
tmp21068 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21069 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp21068)
}
__typedArg0 := symlist
__typedArg1 := tmp21068
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21070 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21071 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21070)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21070
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21072 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21069, tmp21071)
}
__typedArg0 := tmp21069
__typedArg1 := tmp21071
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6706, tmp21072, B6702, L6703, Key6704, C6705)
return


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

tmp21074 := Call(__e, PrimFunc(symshen_4assoc_1_6), symsum, tmp21067, tmp21073)


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
V6711 := __e.Get(1)
_ = V6711
__e.Return(MakeNative(func(__e *ControlFlow) {
B6707 := __e.Get(1)
_ = B6707
__e.Return(MakeNative(func(__e *ControlFlow) {
L6708 := __e.Get(1)
_ = L6708
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6709 := __e.Get(1)
_ = Key6709
__e.Return(MakeNative(func(__e *ControlFlow) {
C6710 := __e.Get(1)
_ = C6710
tmp21077 := Call(__e, PrimFunc(symshen_4newpv), B6707)


let__18962 := tmp21077
_ = let__18962

tmp21078 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21079 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21078)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21078
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21080 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18962, tmp21079)
}
__typedArg0 := let__18962
__typedArg1 := tmp21079
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21081 := Call(__e, PrimFunc(symis_b), V6711, tmp21080, B6707, L6708, Key6709, C6710)


__e.TailApply(PrimFunc(symshen_4gc), B6707, tmp21081)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21082 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21083 := Call(__e, PrimFunc(symshen_4assoc_1_6), symsymbol_2, tmp21076, tmp21082)


tmp21084 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21083)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21083
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21084

tmp21085 := MakeNative(func(__e *ControlFlow) {
V6716 := __e.Get(1)
_ = V6716
__e.Return(MakeNative(func(__e *ControlFlow) {
B6712 := __e.Get(1)
_ = B6712
__e.Return(MakeNative(func(__e *ControlFlow) {
L6713 := __e.Get(1)
_ = L6713
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6714 := __e.Get(1)
_ = Key6714
__e.Return(MakeNative(func(__e *ControlFlow) {
C6715 := __e.Get(1)
_ = C6715
tmp21086 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21087 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21086)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21086
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21088 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp21087)
}
__typedArg0 := symsymbol
__typedArg1 := tmp21087
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6716, tmp21088, B6712, L6713, Key6714, C6715)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21089 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21090 := Call(__e, PrimFunc(symshen_4assoc_1_6), symsystemf, tmp21085, tmp21089)


tmp21091 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21090)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21090
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21091

tmp21092 := MakeNative(func(__e *ControlFlow) {
V6721 := __e.Get(1)
_ = V6721
__e.Return(MakeNative(func(__e *ControlFlow) {
B6717 := __e.Get(1)
_ = B6717
__e.Return(MakeNative(func(__e *ControlFlow) {
L6718 := __e.Get(1)
_ = L6718
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6719 := __e.Get(1)
_ = Key6719
__e.Return(MakeNative(func(__e *ControlFlow) {
C6720 := __e.Get(1)
_ = C6720
tmp21093 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21094 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21093)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21093
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6721, tmp21094, B6717, L6718, Key6719, C6720)
return


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

tmp21096 := Call(__e, PrimFunc(symshen_4assoc_1_6), symsystem_1S_2, tmp21092, tmp21095)


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
V6726 := __e.Get(1)
_ = V6726
__e.Return(MakeNative(func(__e *ControlFlow) {
B6722 := __e.Get(1)
_ = B6722
__e.Return(MakeNative(func(__e *ControlFlow) {
L6723 := __e.Get(1)
_ = L6723
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6724 := __e.Get(1)
_ = Key6724
__e.Return(MakeNative(func(__e *ControlFlow) {
C6725 := __e.Get(1)
_ = C6725
tmp21099 := Call(__e, PrimFunc(symshen_4newpv), B6722)


let__18963 := tmp21099
_ = let__18963

tmp21100 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18963, Nil)
}
__typedArg0 := let__18963
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21101 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp21100)
}
__typedArg0 := symlist
__typedArg1 := tmp21100
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21102 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18963, Nil)
}
__typedArg0 := let__18963
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21103 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp21102)
}
__typedArg0 := symlist
__typedArg1 := tmp21102
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21104 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21103, Nil)
}
__typedArg0 := tmp21103
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21105 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21104)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21104
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21106 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21101, tmp21105)
}
__typedArg0 := tmp21101
__typedArg1 := tmp21105
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21107 := Call(__e, PrimFunc(symis_b), V6726, tmp21106, B6722, L6723, Key6724, C6725)


__e.TailApply(PrimFunc(symshen_4gc), B6722, tmp21107)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21108 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21109 := Call(__e, PrimFunc(symshen_4assoc_1_6), symtail, tmp21098, tmp21108)


tmp21110 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21109)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21109
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21110

tmp21111 := MakeNative(func(__e *ControlFlow) {
V6731 := __e.Get(1)
_ = V6731
__e.Return(MakeNative(func(__e *ControlFlow) {
B6727 := __e.Get(1)
_ = B6727
__e.Return(MakeNative(func(__e *ControlFlow) {
L6728 := __e.Get(1)
_ = L6728
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6729 := __e.Get(1)
_ = Key6729
__e.Return(MakeNative(func(__e *ControlFlow) {
C6730 := __e.Get(1)
_ = C6730
tmp21112 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21113 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21112)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21112
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21114 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp21113)
}
__typedArg0 := symstring
__typedArg1 := tmp21113
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6731, tmp21114, B6727, L6728, Key6729, C6730)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21115 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21116 := Call(__e, PrimFunc(symshen_4assoc_1_6), symtlstr, tmp21111, tmp21115)


tmp21117 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21116)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21116
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21117

tmp21118 := MakeNative(func(__e *ControlFlow) {
V6736 := __e.Get(1)
_ = V6736
__e.Return(MakeNative(func(__e *ControlFlow) {
B6732 := __e.Get(1)
_ = B6732
__e.Return(MakeNative(func(__e *ControlFlow) {
L6733 := __e.Get(1)
_ = L6733
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6734 := __e.Get(1)
_ = Key6734
__e.Return(MakeNative(func(__e *ControlFlow) {
C6735 := __e.Get(1)
_ = C6735
tmp21119 := Call(__e, PrimFunc(symshen_4newpv), B6732)


let__18964 := tmp21119
_ = let__18964

tmp21120 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18964, Nil)
}
__typedArg0 := let__18964
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21121 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp21120)
}
__typedArg0 := symvector
__typedArg1 := tmp21120
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21122 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18964, Nil)
}
__typedArg0 := let__18964
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21123 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp21122)
}
__typedArg0 := symvector
__typedArg1 := tmp21122
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21124 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21123, Nil)
}
__typedArg0 := tmp21123
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
return PrimCons(tmp21121, tmp21125)
}
__typedArg0 := tmp21121
__typedArg1 := tmp21125
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21127 := Call(__e, PrimFunc(symis_b), V6736, tmp21126, B6732, L6733, Key6734, C6735)


__e.TailApply(PrimFunc(symshen_4gc), B6732, tmp21127)
return


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

tmp21129 := Call(__e, PrimFunc(symshen_4assoc_1_6), symtlv, tmp21118, tmp21128)


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
V6741 := __e.Get(1)
_ = V6741
__e.Return(MakeNative(func(__e *ControlFlow) {
B6737 := __e.Get(1)
_ = B6737
__e.Return(MakeNative(func(__e *ControlFlow) {
L6738 := __e.Get(1)
_ = L6738
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6739 := __e.Get(1)
_ = Key6739
__e.Return(MakeNative(func(__e *ControlFlow) {
C6740 := __e.Get(1)
_ = C6740
tmp21132 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
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
return PrimCons(symsymbol, tmp21133)
}
__typedArg0 := symsymbol
__typedArg1 := tmp21133
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6741, tmp21134, B6737, L6738, Key6739, C6740)
return


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

tmp21136 := Call(__e, PrimFunc(symshen_4assoc_1_6), symtc, tmp21131, tmp21135)


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
V6746 := __e.Get(1)
_ = V6746
__e.Return(MakeNative(func(__e *ControlFlow) {
B6742 := __e.Get(1)
_ = B6742
__e.Return(MakeNative(func(__e *ControlFlow) {
L6743 := __e.Get(1)
_ = L6743
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6744 := __e.Get(1)
_ = Key6744
__e.Return(MakeNative(func(__e *ControlFlow) {
C6745 := __e.Get(1)
_ = C6745
tmp21139 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
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

__e.TailApply(PrimFunc(symis_b), V6746, tmp21140, B6742, L6743, Key6744, C6745)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21141 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21142 := Call(__e, PrimFunc(symshen_4assoc_1_6), symtc_2, tmp21138, tmp21141)


tmp21143 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21142)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21142
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21143

tmp21144 := MakeNative(func(__e *ControlFlow) {
V6751 := __e.Get(1)
_ = V6751
__e.Return(MakeNative(func(__e *ControlFlow) {
B6747 := __e.Get(1)
_ = B6747
__e.Return(MakeNative(func(__e *ControlFlow) {
L6748 := __e.Get(1)
_ = L6748
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6749 := __e.Get(1)
_ = Key6749
__e.Return(MakeNative(func(__e *ControlFlow) {
C6750 := __e.Get(1)
_ = C6750
tmp21145 := Call(__e, PrimFunc(symshen_4newpv), B6747)


let__18965 := tmp21145
_ = let__18965

tmp21146 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18965, Nil)
}
__typedArg0 := let__18965
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21147 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlazy, tmp21146)
}
__typedArg0 := symlazy
__typedArg1 := tmp21146
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21148 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18965, Nil)
}
__typedArg0 := let__18965
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

tmp21151 := Call(__e, PrimFunc(symis_b), V6751, tmp21150, B6747, L6748, Key6749, C6750)


__e.TailApply(PrimFunc(symshen_4gc), B6747, tmp21151)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21152 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21153 := Call(__e, PrimFunc(symshen_4assoc_1_6), symthaw, tmp21144, tmp21152)


tmp21154 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21153)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21153
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21154

tmp21155 := MakeNative(func(__e *ControlFlow) {
V6756 := __e.Get(1)
_ = V6756
__e.Return(MakeNative(func(__e *ControlFlow) {
B6752 := __e.Get(1)
_ = B6752
__e.Return(MakeNative(func(__e *ControlFlow) {
L6753 := __e.Get(1)
_ = L6753
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6754 := __e.Get(1)
_ = Key6754
__e.Return(MakeNative(func(__e *ControlFlow) {
C6755 := __e.Get(1)
_ = C6755
tmp21156 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
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
return PrimCons(symsymbol, tmp21157)
}
__typedArg0 := symsymbol
__typedArg1 := tmp21157
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6756, tmp21158, B6752, L6753, Key6754, C6755)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21159 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21160 := Call(__e, PrimFunc(symshen_4assoc_1_6), symtrack, tmp21155, tmp21159)


tmp21161 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21160)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21160
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21161

tmp21162 := MakeNative(func(__e *ControlFlow) {
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
tmp21163 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21164 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp21163)
}
__typedArg0 := symlist
__typedArg1 := tmp21163
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21165 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21164, Nil)
}
__typedArg0 := tmp21164
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21166 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21165)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21165
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6761, tmp21166, B6757, L6758, Key6759, C6760)
return


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

tmp21168 := Call(__e, PrimFunc(symshen_4assoc_1_6), symtracked, tmp21162, tmp21167)


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
tmp21171 := Call(__e, PrimFunc(symshen_4newpv), B6762)


let__18966 := tmp21171
_ = let__18966

tmp21172 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18966, Nil)
}
__typedArg0 := let__18966
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21173 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21172)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21172
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21174 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symexception, tmp21173)
}
__typedArg0 := symexception
__typedArg1 := tmp21173
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21175 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18966, Nil)
}
__typedArg0 := let__18966
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21176 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21175)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21175
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21177 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21174, tmp21176)
}
__typedArg0 := tmp21174
__typedArg1 := tmp21176
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21178 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21177, Nil)
}
__typedArg0 := tmp21177
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21179 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21178)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21178
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21180 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18966, tmp21179)
}
__typedArg0 := let__18966
__typedArg1 := tmp21179
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21181 := Call(__e, PrimFunc(symis_b), V6766, tmp21180, B6762, L6763, Key6764, C6765)


__e.TailApply(PrimFunc(symshen_4gc), B6762, tmp21181)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21182 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21183 := Call(__e, PrimFunc(symshen_4assoc_1_6), symtrap_1error, tmp21170, tmp21182)


tmp21184 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21183)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21183
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21184

tmp21185 := MakeNative(func(__e *ControlFlow) {
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
tmp21186 := Call(__e, PrimFunc(symshen_4newpv), B6767)


let__18967 := tmp21186
_ = let__18967

tmp21187 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21188 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21187)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21187
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21189 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18967, tmp21188)
}
__typedArg0 := let__18967
__typedArg1 := tmp21188
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21190 := Call(__e, PrimFunc(symis_b), V6771, tmp21189, B6767, L6768, Key6769, C6770)


__e.TailApply(PrimFunc(symshen_4gc), B6767, tmp21190)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21191 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21192 := Call(__e, PrimFunc(symshen_4assoc_1_6), symtuple_2, tmp21185, tmp21191)


tmp21193 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21192)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21192
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21193

tmp21194 := MakeNative(func(__e *ControlFlow) {
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
tmp21195 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21196 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp21195)
}
__typedArg0 := symlist
__typedArg1 := tmp21195
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21197 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21196, Nil)
}
__typedArg0 := tmp21196
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21198 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21197)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21197
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21199 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp21198)
}
__typedArg0 := symstring
__typedArg1 := tmp21198
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6776, tmp21199, B6772, L6773, Key6774, C6775)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21200 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21201 := Call(__e, PrimFunc(symshen_4assoc_1_6), symunabsolute, tmp21194, tmp21200)


tmp21202 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21201)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21201
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21202

tmp21203 := MakeNative(func(__e *ControlFlow) {
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
tmp21204 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21205 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21204)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21204
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21206 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp21205)
}
__typedArg0 := symsymbol
__typedArg1 := tmp21205
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6781, tmp21206, B6777, L6778, Key6779, C6780)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21207 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21208 := Call(__e, PrimFunc(symshen_4assoc_1_6), symundefmacro, tmp21203, tmp21207)


tmp21209 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21208)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21208
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21209

tmp21210 := MakeNative(func(__e *ControlFlow) {
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
tmp21211 := Call(__e, PrimFunc(symshen_4newpv), B6782)


let__18968 := tmp21211
_ = let__18968

tmp21212 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18968, Nil)
}
__typedArg0 := let__18968
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21213 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp21212)
}
__typedArg0 := symlist
__typedArg1 := tmp21212
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21214 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18968, Nil)
}
__typedArg0 := let__18968
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21215 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp21214)
}
__typedArg0 := symlist
__typedArg1 := tmp21214
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21216 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18968, Nil)
}
__typedArg0 := let__18968
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21217 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp21216)
}
__typedArg0 := symlist
__typedArg1 := tmp21216
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21218 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21217, Nil)
}
__typedArg0 := tmp21217
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21219 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21218)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21218
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21220 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21215, tmp21219)
}
__typedArg0 := tmp21215
__typedArg1 := tmp21219
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21221 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21220, Nil)
}
__typedArg0 := tmp21220
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21222 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21221)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21221
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21223 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21213, tmp21222)
}
__typedArg0 := tmp21213
__typedArg1 := tmp21222
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21224 := Call(__e, PrimFunc(symis_b), V6786, tmp21223, B6782, L6783, Key6784, C6785)


__e.TailApply(PrimFunc(symshen_4gc), B6782, tmp21224)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21225 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21226 := Call(__e, PrimFunc(symshen_4assoc_1_6), symunion, tmp21210, tmp21225)


tmp21227 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21226)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21226
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21227

tmp21228 := MakeNative(func(__e *ControlFlow) {
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
tmp21229 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21230 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21229)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21229
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21231 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp21230)
}
__typedArg0 := symsymbol
__typedArg1 := tmp21230
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6791, tmp21231, B6787, L6788, Key6789, C6790)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21232 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21233 := Call(__e, PrimFunc(symshen_4assoc_1_6), symunprofile, tmp21228, tmp21232)


tmp21234 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21233)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21233
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21234

tmp21235 := MakeNative(func(__e *ControlFlow) {
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
tmp21236 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21237 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21236)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21236
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21238 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp21237)
}
__typedArg0 := symsymbol
__typedArg1 := tmp21237
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6796, tmp21238, B6792, L6793, Key6794, C6795)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21239 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21240 := Call(__e, PrimFunc(symshen_4assoc_1_6), symuntrack, tmp21235, tmp21239)


tmp21241 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21240)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21240
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21241

tmp21242 := MakeNative(func(__e *ControlFlow) {
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
tmp21243 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21244 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp21243)
}
__typedArg0 := symlist
__typedArg1 := tmp21243
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21245 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21244, Nil)
}
__typedArg0 := tmp21244
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21246 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21245)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21245
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6801, tmp21246, B6797, L6798, Key6799, C6800)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21247 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21248 := Call(__e, PrimFunc(symshen_4assoc_1_6), symuserdefs, tmp21242, tmp21247)


tmp21249 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21248)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21248
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21249

tmp21250 := MakeNative(func(__e *ControlFlow) {
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
tmp21251 := Call(__e, PrimFunc(symshen_4newpv), B6802)


let__18969 := tmp21251
_ = let__18969

tmp21252 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21253 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21252)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21252
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21254 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18969, tmp21253)
}
__typedArg0 := let__18969
__typedArg1 := tmp21253
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21255 := Call(__e, PrimFunc(symis_b), V6806, tmp21254, B6802, L6803, Key6804, C6805)


__e.TailApply(PrimFunc(symshen_4gc), B6802, tmp21255)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21256 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21257 := Call(__e, PrimFunc(symshen_4assoc_1_6), symvariable_2, tmp21250, tmp21256)


tmp21258 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21257)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21257
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21258

tmp21259 := MakeNative(func(__e *ControlFlow) {
V6811 := __e.Get(1)
_ = V6811
__e.Return(MakeNative(func(__e *ControlFlow) {
B6807 := __e.Get(1)
_ = B6807
__e.Return(MakeNative(func(__e *ControlFlow) {
L6808 := __e.Get(1)
_ = L6808
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6809 := __e.Get(1)
_ = Key6809
__e.Return(MakeNative(func(__e *ControlFlow) {
C6810 := __e.Get(1)
_ = C6810
tmp21260 := Call(__e, PrimFunc(symshen_4newpv), B6807)


let__18970 := tmp21260
_ = let__18970

tmp21261 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21262 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21261)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21261
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21263 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18970, tmp21262)
}
__typedArg0 := let__18970
__typedArg1 := tmp21262
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21264 := Call(__e, PrimFunc(symis_b), V6811, tmp21263, B6807, L6808, Key6809, C6810)


__e.TailApply(PrimFunc(symshen_4gc), B6807, tmp21264)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21265 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21266 := Call(__e, PrimFunc(symshen_4assoc_1_6), symvector_2, tmp21259, tmp21265)


tmp21267 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21266)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21266
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21267

tmp21268 := MakeNative(func(__e *ControlFlow) {
V6816 := __e.Get(1)
_ = V6816
__e.Return(MakeNative(func(__e *ControlFlow) {
B6812 := __e.Get(1)
_ = B6812
__e.Return(MakeNative(func(__e *ControlFlow) {
L6813 := __e.Get(1)
_ = L6813
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6814 := __e.Get(1)
_ = Key6814
__e.Return(MakeNative(func(__e *ControlFlow) {
C6815 := __e.Get(1)
_ = C6815
tmp21269 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21270 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21269)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21269
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6816, tmp21270, B6812, L6813, Key6814, C6815)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21271 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21272 := Call(__e, PrimFunc(symshen_4assoc_1_6), symversion, tmp21268, tmp21271)


tmp21273 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21272)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21272
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21273

tmp21274 := MakeNative(func(__e *ControlFlow) {
V6821 := __e.Get(1)
_ = V6821
__e.Return(MakeNative(func(__e *ControlFlow) {
B6817 := __e.Get(1)
_ = B6817
__e.Return(MakeNative(func(__e *ControlFlow) {
L6818 := __e.Get(1)
_ = L6818
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6819 := __e.Get(1)
_ = Key6819
__e.Return(MakeNative(func(__e *ControlFlow) {
C6820 := __e.Get(1)
_ = C6820
tmp21275 := Call(__e, PrimFunc(symshen_4newpv), B6817)


let__18971 := tmp21275
_ = let__18971

tmp21276 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18971, Nil)
}
__typedArg0 := let__18971
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21277 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21276)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21276
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21278 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18971, tmp21277)
}
__typedArg0 := let__18971
__typedArg1 := tmp21277
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21279 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21278, Nil)
}
__typedArg0 := tmp21278
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21280 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21279)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21279
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21281 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp21280)
}
__typedArg0 := symstring
__typedArg1 := tmp21280
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21282 := Call(__e, PrimFunc(symis_b), V6821, tmp21281, B6817, L6818, Key6819, C6820)


__e.TailApply(PrimFunc(symshen_4gc), B6817, tmp21282)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21283 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21284 := Call(__e, PrimFunc(symshen_4assoc_1_6), symwrite_1to_1file, tmp21274, tmp21283)


tmp21285 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21284)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21284
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21285

tmp21286 := MakeNative(func(__e *ControlFlow) {
V6826 := __e.Get(1)
_ = V6826
__e.Return(MakeNative(func(__e *ControlFlow) {
B6822 := __e.Get(1)
_ = B6822
__e.Return(MakeNative(func(__e *ControlFlow) {
L6823 := __e.Get(1)
_ = L6823
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6824 := __e.Get(1)
_ = Key6824
__e.Return(MakeNative(func(__e *ControlFlow) {
C6825 := __e.Get(1)
_ = C6825
tmp21287 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symout, Nil)
}
__typedArg0 := symout
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21288 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp21287)
}
__typedArg0 := symstream
__typedArg1 := tmp21287
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21289 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21290 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21289)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21289
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21291 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21288, tmp21290)
}
__typedArg0 := tmp21288
__typedArg1 := tmp21290
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21292 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21291, Nil)
}
__typedArg0 := tmp21291
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21293 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21292)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21292
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21294 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp21293)
}
__typedArg0 := symnumber
__typedArg1 := tmp21293
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6826, tmp21294, B6822, L6823, Key6824, C6825)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21295 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21296 := Call(__e, PrimFunc(symshen_4assoc_1_6), symwrite_1byte, tmp21286, tmp21295)


tmp21297 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21296)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21296
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21297

tmp21298 := MakeNative(func(__e *ControlFlow) {
V6831 := __e.Get(1)
_ = V6831
__e.Return(MakeNative(func(__e *ControlFlow) {
B6827 := __e.Get(1)
_ = B6827
__e.Return(MakeNative(func(__e *ControlFlow) {
L6828 := __e.Get(1)
_ = L6828
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6829 := __e.Get(1)
_ = Key6829
__e.Return(MakeNative(func(__e *ControlFlow) {
C6830 := __e.Get(1)
_ = C6830
tmp21299 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
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
return PrimCons(symstring, tmp21300)
}
__typedArg0 := symstring
__typedArg1 := tmp21300
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6831, tmp21301, B6827, L6828, Key6829, C6830)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21302 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21303 := Call(__e, PrimFunc(symshen_4assoc_1_6), symy_1or_1n_2, tmp21298, tmp21302)


tmp21304 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21303)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21303
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21304

tmp21305 := MakeNative(func(__e *ControlFlow) {
V6836 := __e.Get(1)
_ = V6836
__e.Return(MakeNative(func(__e *ControlFlow) {
B6832 := __e.Get(1)
_ = B6832
__e.Return(MakeNative(func(__e *ControlFlow) {
L6833 := __e.Get(1)
_ = L6833
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6834 := __e.Get(1)
_ = Key6834
__e.Return(MakeNative(func(__e *ControlFlow) {
C6835 := __e.Get(1)
_ = C6835
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
return PrimCons(symnumber, tmp21307)
}
__typedArg0 := symnumber
__typedArg1 := tmp21307
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21309 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21308, Nil)
}
__typedArg0 := tmp21308
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21310 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21309)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21309
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21311 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp21310)
}
__typedArg0 := symnumber
__typedArg1 := tmp21310
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6836, tmp21311, B6832, L6833, Key6834, C6835)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21312 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21313 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_6, tmp21305, tmp21312)


tmp21314 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21313)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21313
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21314

tmp21315 := MakeNative(func(__e *ControlFlow) {
V6841 := __e.Get(1)
_ = V6841
__e.Return(MakeNative(func(__e *ControlFlow) {
B6837 := __e.Get(1)
_ = B6837
__e.Return(MakeNative(func(__e *ControlFlow) {
L6838 := __e.Get(1)
_ = L6838
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6839 := __e.Get(1)
_ = Key6839
__e.Return(MakeNative(func(__e *ControlFlow) {
C6840 := __e.Get(1)
_ = C6840
tmp21316 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21317 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21316)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21316
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21318 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp21317)
}
__typedArg0 := symnumber
__typedArg1 := tmp21317
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21319 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21318, Nil)
}
__typedArg0 := tmp21318
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21320 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21319)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21319
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21321 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp21320)
}
__typedArg0 := symnumber
__typedArg1 := tmp21320
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6841, tmp21321, B6837, L6838, Key6839, C6840)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21322 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21323 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_5, tmp21315, tmp21322)


tmp21324 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21323)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21323
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21324

tmp21325 := MakeNative(func(__e *ControlFlow) {
V6846 := __e.Get(1)
_ = V6846
__e.Return(MakeNative(func(__e *ControlFlow) {
B6842 := __e.Get(1)
_ = B6842
__e.Return(MakeNative(func(__e *ControlFlow) {
L6843 := __e.Get(1)
_ = L6843
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6844 := __e.Get(1)
_ = Key6844
__e.Return(MakeNative(func(__e *ControlFlow) {
C6845 := __e.Get(1)
_ = C6845
tmp21326 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21327 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21326)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21326
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21328 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp21327)
}
__typedArg0 := symnumber
__typedArg1 := tmp21327
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21329 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21328, Nil)
}
__typedArg0 := tmp21328
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21330 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21329)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21329
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21331 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp21330)
}
__typedArg0 := symnumber
__typedArg1 := tmp21330
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6846, tmp21331, B6842, L6843, Key6844, C6845)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21332 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21333 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_6_a, tmp21325, tmp21332)


tmp21334 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21333)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21333
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21334

tmp21335 := MakeNative(func(__e *ControlFlow) {
V6851 := __e.Get(1)
_ = V6851
__e.Return(MakeNative(func(__e *ControlFlow) {
B6847 := __e.Get(1)
_ = B6847
__e.Return(MakeNative(func(__e *ControlFlow) {
L6848 := __e.Get(1)
_ = L6848
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6849 := __e.Get(1)
_ = Key6849
__e.Return(MakeNative(func(__e *ControlFlow) {
C6850 := __e.Get(1)
_ = C6850
tmp21336 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21337 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21336)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21336
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21338 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp21337)
}
__typedArg0 := symnumber
__typedArg1 := tmp21337
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21339 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21338, Nil)
}
__typedArg0 := tmp21338
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
return PrimCons(symnumber, tmp21340)
}
__typedArg0 := symnumber
__typedArg1 := tmp21340
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6851, tmp21341, B6847, L6848, Key6849, C6850)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21342 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21343 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_5_a, tmp21335, tmp21342)


tmp21344 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21343)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21343
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21344

tmp21345 := MakeNative(func(__e *ControlFlow) {
V6856 := __e.Get(1)
_ = V6856
__e.Return(MakeNative(func(__e *ControlFlow) {
B6852 := __e.Get(1)
_ = B6852
__e.Return(MakeNative(func(__e *ControlFlow) {
L6853 := __e.Get(1)
_ = L6853
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6854 := __e.Get(1)
_ = Key6854
__e.Return(MakeNative(func(__e *ControlFlow) {
C6855 := __e.Get(1)
_ = C6855
tmp21346 := Call(__e, PrimFunc(symshen_4newpv), B6852)


let__18972 := tmp21346
_ = let__18972

tmp21347 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
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

tmp21349 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18972, tmp21348)
}
__typedArg0 := let__18972
__typedArg1 := tmp21348
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21350 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21349, Nil)
}
__typedArg0 := tmp21349
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21351 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21350)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21350
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21352 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18972, tmp21351)
}
__typedArg0 := let__18972
__typedArg1 := tmp21351
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21353 := Call(__e, PrimFunc(symis_b), V6856, tmp21352, B6852, L6853, Key6854, C6855)


__e.TailApply(PrimFunc(symshen_4gc), B6852, tmp21353)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21354 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21355 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_a, tmp21345, tmp21354)


tmp21356 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21355)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21355
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21356

tmp21357 := MakeNative(func(__e *ControlFlow) {
V6861 := __e.Get(1)
_ = V6861
__e.Return(MakeNative(func(__e *ControlFlow) {
B6857 := __e.Get(1)
_ = B6857
__e.Return(MakeNative(func(__e *ControlFlow) {
L6858 := __e.Get(1)
_ = L6858
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6859 := __e.Get(1)
_ = Key6859
__e.Return(MakeNative(func(__e *ControlFlow) {
C6860 := __e.Get(1)
_ = C6860
tmp21358 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21359 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21358)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21358
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21360 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp21359)
}
__typedArg0 := symnumber
__typedArg1 := tmp21359
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21361 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21360, Nil)
}
__typedArg0 := tmp21360
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21362 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21361)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21361
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21363 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp21362)
}
__typedArg0 := symnumber
__typedArg1 := tmp21362
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6861, tmp21363, B6857, L6858, Key6859, C6860)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21364 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21365 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_7, tmp21357, tmp21364)


tmp21366 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21365)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21365
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21366

tmp21367 := MakeNative(func(__e *ControlFlow) {
V6866 := __e.Get(1)
_ = V6866
__e.Return(MakeNative(func(__e *ControlFlow) {
B6862 := __e.Get(1)
_ = B6862
__e.Return(MakeNative(func(__e *ControlFlow) {
L6863 := __e.Get(1)
_ = L6863
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6864 := __e.Get(1)
_ = Key6864
__e.Return(MakeNative(func(__e *ControlFlow) {
C6865 := __e.Get(1)
_ = C6865
tmp21368 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21369 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21368)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21368
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21370 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp21369)
}
__typedArg0 := symnumber
__typedArg1 := tmp21369
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21371 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21370, Nil)
}
__typedArg0 := tmp21370
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21372 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21371)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21371
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21373 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp21372)
}
__typedArg0 := symnumber
__typedArg1 := tmp21372
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6866, tmp21373, B6862, L6863, Key6864, C6865)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21374 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21375 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_c, tmp21367, tmp21374)


tmp21376 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21375)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21375
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21376

tmp21377 := MakeNative(func(__e *ControlFlow) {
V6871 := __e.Get(1)
_ = V6871
__e.Return(MakeNative(func(__e *ControlFlow) {
B6867 := __e.Get(1)
_ = B6867
__e.Return(MakeNative(func(__e *ControlFlow) {
L6868 := __e.Get(1)
_ = L6868
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6869 := __e.Get(1)
_ = Key6869
__e.Return(MakeNative(func(__e *ControlFlow) {
C6870 := __e.Get(1)
_ = C6870
tmp21378 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21379 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21378)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21378
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21380 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp21379)
}
__typedArg0 := symnumber
__typedArg1 := tmp21379
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21381 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21380, Nil)
}
__typedArg0 := tmp21380
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21382 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21381)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21381
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21383 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp21382)
}
__typedArg0 := symnumber
__typedArg1 := tmp21382
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6871, tmp21383, B6867, L6868, Key6869, C6870)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21384 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21385 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_1, tmp21377, tmp21384)


tmp21386 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21385)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21385
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21386

tmp21387 := MakeNative(func(__e *ControlFlow) {
V6876 := __e.Get(1)
_ = V6876
__e.Return(MakeNative(func(__e *ControlFlow) {
B6872 := __e.Get(1)
_ = B6872
__e.Return(MakeNative(func(__e *ControlFlow) {
L6873 := __e.Get(1)
_ = L6873
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6874 := __e.Get(1)
_ = Key6874
__e.Return(MakeNative(func(__e *ControlFlow) {
C6875 := __e.Get(1)
_ = C6875
tmp21388 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21389 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21388)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21388
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21390 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp21389)
}
__typedArg0 := symnumber
__typedArg1 := tmp21389
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21391 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21390, Nil)
}
__typedArg0 := tmp21390
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21392 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21391)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21391
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21393 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp21392)
}
__typedArg0 := symnumber
__typedArg1 := tmp21392
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6876, tmp21393, B6872, L6873, Key6874, C6875)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21394 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21395 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_d, tmp21387, tmp21394)


tmp21396 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21395)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21395
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21396

tmp21397 := MakeNative(func(__e *ControlFlow) {
V6881 := __e.Get(1)
_ = V6881
__e.Return(MakeNative(func(__e *ControlFlow) {
B6877 := __e.Get(1)
_ = B6877
__e.Return(MakeNative(func(__e *ControlFlow) {
L6878 := __e.Get(1)
_ = L6878
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6879 := __e.Get(1)
_ = Key6879
__e.Return(MakeNative(func(__e *ControlFlow) {
C6880 := __e.Get(1)
_ = C6880
tmp21398 := Call(__e, PrimFunc(symshen_4newpv), B6877)


let__18973 := tmp21398
_ = let__18973

tmp21399 := Call(__e, PrimFunc(symshen_4newpv), B6877)


let__18974 := tmp21399
_ = let__18974

tmp21400 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21401 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21400)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21400
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21402 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18974, tmp21401)
}
__typedArg0 := let__18974
__typedArg1 := tmp21401
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21403 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21402, Nil)
}
__typedArg0 := tmp21402
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21404 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21403)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21403
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21405 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18973, tmp21404)
}
__typedArg0 := let__18973
__typedArg1 := tmp21404
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21406 := Call(__e, PrimFunc(symis_b), V6881, tmp21405, B6877, L6878, Key6879, C6880)


tmp21407 := Call(__e, PrimFunc(symshen_4gc), B6877, tmp21406)


__e.TailApply(PrimFunc(symshen_4gc), B6877, tmp21407)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21408 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21409 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_a_a, tmp21397, tmp21408)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21409)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21409
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return


}, 0)

tmp21410 := Call(__e, ns2_1set, symshen_4initialise_1signedfuncs, tmp19723)


_ = tmp21410

tmp21411 := MakeNative(func(__e *ControlFlow) {
tmp21412 := MakeNative(func(__e *ControlFlow) {
Y1242 := __e.Get(1)
_ = Y1242
__e.TailApply(PrimFunc(symshen_4tuple), Y1242)
return
}, 1)

tmp21413 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4tuple, tmp21412)
}
__typedArg0 := symshen_4tuple
__typedArg1 := tmp21412
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21414 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21413)


_ = tmp21414

tmp21415 := MakeNative(func(__e *ControlFlow) {
Y1241 := __e.Get(1)
_ = Y1241
__e.TailApply(PrimFunc(symshen_4pvar), Y1241)
return
}, 1)

tmp21416 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4pvar, tmp21415)
}
__typedArg0 := symshen_4pvar
__typedArg1 := tmp21415
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21417 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21416)


_ = tmp21417

tmp21418 := MakeNative(func(__e *ControlFlow) {
Y1240 := __e.Get(1)
_ = Y1240
__e.TailApply(PrimFunc(symshen_4dictionary), Y1240)
return
}, 1)

tmp21419 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4dictionary, tmp21418)
}
__typedArg0 := symshen_4dictionary
__typedArg1 := tmp21418
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21420 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21419)


_ = tmp21420

tmp21421 := MakeNative(func(__e *ControlFlow) {
Y1239 := __e.Get(1)
_ = Y1239
__e.TailApply(PrimFunc(symshen_4print_1prolog_1vector), Y1239)
return
}, 1)

tmp21422 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4print_1prolog_1vector, tmp21421)
}
__typedArg0 := symshen_4print_1prolog_1vector
__typedArg1 := tmp21421
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21423 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21422)


_ = tmp21423

tmp21424 := MakeNative(func(__e *ControlFlow) {
Y1238 := __e.Get(1)
_ = Y1238
__e.TailApply(PrimFunc(symshen_4print_1freshterm), Y1238)
return
}, 1)

tmp21425 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4print_1freshterm, tmp21424)
}
__typedArg0 := symshen_4print_1freshterm
__typedArg1 := tmp21424
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21426 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21425)


_ = tmp21426

tmp21427 := MakeNative(func(__e *ControlFlow) {
Y1237 := __e.Get(1)
_ = Y1237
__e.TailApply(PrimFunc(symshen_4printF), Y1237)
return
}, 1)

tmp21428 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4printF, tmp21427)
}
__typedArg0 := symshen_4printF
__typedArg1 := tmp21427
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21429 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21428)


_ = tmp21429

tmp21430 := MakeNative(func(__e *ControlFlow) {
Y1236 := __e.Get(1)
_ = Y1236
__e.TailApply(PrimFunc(symabsolute), Y1236)
return
}, 1)

tmp21431 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symabsolute, tmp21430)
}
__typedArg0 := symabsolute
__typedArg1 := tmp21430
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21432 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21431)


_ = tmp21432

tmp21433 := MakeNative(func(__e *ControlFlow) {
Y1235 := __e.Get(1)
_ = Y1235
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symabsvector_2) {
return PrimIsVector(Y1235)
}
__typedArg0 := Y1235
return Call(__e, PrimFunc(symabsvector_2), __typedArg0)
})())
return
}, 1)

tmp21434 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symabsvector_2, tmp21433)
}
__typedArg0 := symabsvector_2
__typedArg1 := tmp21433
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21435 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21434)


_ = tmp21435

tmp21436 := MakeNative(func(__e *ControlFlow) {
Y1234 := __e.Get(1)
_ = Y1234
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symabsvector) {
return PrimAbsvector(Y1234)
}
__typedArg0 := Y1234
return Call(__e, PrimFunc(symabsvector), __typedArg0)
})())
return
}, 1)

tmp21437 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symabsvector, tmp21436)
}
__typedArg0 := symabsvector
__typedArg1 := tmp21436
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21438 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21437)


_ = tmp21438

tmp21439 := MakeNative(func(__e *ControlFlow) {
Y1231 := __e.Get(1)
_ = Y1231
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1232 := __e.Get(1)
_ = Y1232
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1233 := __e.Get(1)
_ = Y1233
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(Y1231, Y1232, Y1233)
}
__typedArg0 := Y1231
__typedArg1 := Y1232
__typedArg2 := Y1233
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})())
return
}, 1))
return
}, 1))
return
}, 1)

tmp21440 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symaddress_1_6, tmp21439)
}
__typedArg0 := symaddress_1_6
__typedArg1 := tmp21439
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21441 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21440)


_ = tmp21441

tmp21442 := MakeNative(func(__e *ControlFlow) {
Y1229 := __e.Get(1)
_ = Y1229
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1230 := __e.Get(1)
_ = Y1230
__e.TailApply(PrimFunc(symadjoin), Y1229, Y1230)
return
}, 1))
return
}, 1)

tmp21443 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symadjoin, tmp21442)
}
__typedArg0 := symadjoin
__typedArg1 := tmp21442
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21444 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21443)


_ = tmp21444

tmp21445 := MakeNative(func(__e *ControlFlow) {
Y1227 := __e.Get(1)
_ = Y1227
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1228 := __e.Get(1)
_ = Y1228
if True == Y1227 {
if True == Y1228 {
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

tmp21448 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symand, tmp21445)
}
__typedArg0 := symand
__typedArg1 := tmp21445
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21449 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21448)


_ = tmp21449

tmp21450 := MakeNative(func(__e *ControlFlow) {
Y1225 := __e.Get(1)
_ = Y1225
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1226 := __e.Get(1)
_ = Y1226
__e.TailApply(PrimFunc(symappend), Y1225, Y1226)
return
}, 1))
return
}, 1)

tmp21451 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symappend, tmp21450)
}
__typedArg0 := symappend
__typedArg1 := tmp21450
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21452 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21451)


_ = tmp21452

tmp21453 := MakeNative(func(__e *ControlFlow) {
Y1224 := __e.Get(1)
_ = Y1224
__e.TailApply(PrimFunc(symarity), Y1224)
return
}, 1)

tmp21454 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symarity, tmp21453)
}
__typedArg0 := symarity
__typedArg1 := tmp21453
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21455 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21454)


_ = tmp21455

tmp21456 := MakeNative(func(__e *ControlFlow) {
Y1222 := __e.Get(1)
_ = Y1222
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1223 := __e.Get(1)
_ = Y1223
__e.TailApply(PrimFunc(symassoc), Y1222, Y1223)
return
}, 1))
return
}, 1)

tmp21457 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symassoc, tmp21456)
}
__typedArg0 := symassoc
__typedArg1 := tmp21456
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21458 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21457)


_ = tmp21458

tmp21459 := MakeNative(func(__e *ControlFlow) {
Y1221 := __e.Get(1)
_ = Y1221
__e.TailApply(PrimFunc(symatom_2), Y1221)
return
}, 1)

tmp21460 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symatom_2, tmp21459)
}
__typedArg0 := symatom_2
__typedArg1 := tmp21459
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21461 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21460)


_ = tmp21461

tmp21462 := MakeNative(func(__e *ControlFlow) {
Y1220 := __e.Get(1)
_ = Y1220
__e.TailApply(PrimFunc(symboolean_2), Y1220)
return
}, 1)

tmp21463 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean_2, tmp21462)
}
__typedArg0 := symboolean_2
__typedArg1 := tmp21462
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21464 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21463)


_ = tmp21464

tmp21465 := MakeNative(func(__e *ControlFlow) {
Y1219 := __e.Get(1)
_ = Y1219
__e.TailApply(PrimFunc(symbootstrap), Y1219)
return
}, 1)

tmp21466 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symbootstrap, tmp21465)
}
__typedArg0 := symbootstrap
__typedArg1 := tmp21465
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21467 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21466)


_ = tmp21467

tmp21468 := MakeNative(func(__e *ControlFlow) {
Y1218 := __e.Get(1)
_ = Y1218
__e.TailApply(PrimFunc(symbound_2), Y1218)
return
}, 1)

tmp21469 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symbound_2, tmp21468)
}
__typedArg0 := symbound_2
__typedArg1 := tmp21468
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21470 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21469)


_ = tmp21470

tmp21471 := MakeNative(func(__e *ControlFlow) {
Y1212 := __e.Get(1)
_ = Y1212
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1213 := __e.Get(1)
_ = Y1213
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1214 := __e.Get(1)
_ = Y1214
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1215 := __e.Get(1)
_ = Y1215
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1216 := __e.Get(1)
_ = Y1216
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1217 := __e.Get(1)
_ = Y1217
__e.TailApply(PrimFunc(symbind), Y1212, Y1213, Y1214, Y1215, Y1216, Y1217)
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

tmp21472 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symbind, tmp21471)
}
__typedArg0 := symbind
__typedArg1 := tmp21471
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21473 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21472)


_ = tmp21473

tmp21474 := MakeNative(func(__e *ControlFlow) {
Y1207 := __e.Get(1)
_ = Y1207
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1208 := __e.Get(1)
_ = Y1208
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1209 := __e.Get(1)
_ = Y1209
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1210 := __e.Get(1)
_ = Y1210
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1211 := __e.Get(1)
_ = Y1211
__e.TailApply(PrimFunc(symcall), Y1207, Y1208, Y1209, Y1210, Y1211)
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21475 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcall, tmp21474)
}
__typedArg0 := symcall
__typedArg1 := tmp21474
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21476 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21475)


_ = tmp21476

tmp21477 := MakeNative(func(__e *ControlFlow) {
Y1206 := __e.Get(1)
_ = Y1206
__e.TailApply(PrimFunc(symcd), Y1206)
return
}, 1)

tmp21478 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcd, tmp21477)
}
__typedArg0 := symcd
__typedArg1 := tmp21477
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21479 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21478)


_ = tmp21479

tmp21480 := MakeNative(func(__e *ControlFlow) {
Y1204 := __e.Get(1)
_ = Y1204
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1205 := __e.Get(1)
_ = Y1205
__e.TailApply(PrimFunc(symcompile), Y1204, Y1205)
return
}, 1))
return
}, 1)

tmp21481 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcompile, tmp21480)
}
__typedArg0 := symcompile
__typedArg1 := tmp21480
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21482 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21481)


_ = tmp21482

tmp21483 := MakeNative(func(__e *ControlFlow) {
Y1202 := __e.Get(1)
_ = Y1202
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1203 := __e.Get(1)
_ = Y1203
__e.TailApply(PrimFunc(symconcat), Y1202, Y1203)
return
}, 1))
return
}, 1)

tmp21484 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symconcat, tmp21483)
}
__typedArg0 := symconcat
__typedArg1 := tmp21483
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21485 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21484)


_ = tmp21485

tmp21486 := MakeNative(func(__e *ControlFlow) {
Y1200 := __e.Get(1)
_ = Y1200
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1201 := __e.Get(1)
_ = Y1201
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Y1200, Y1201)
}
__typedArg0 := Y1200
__typedArg1 := Y1201
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return
}, 1))
return
}, 1)

tmp21487 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp21486)
}
__typedArg0 := symcons
__typedArg1 := tmp21486
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21488 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21487)


_ = tmp21488

tmp21489 := MakeNative(func(__e *ControlFlow) {
Y1199 := __e.Get(1)
_ = Y1199
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(Y1199)
}
__typedArg0 := Y1199
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})())
return
}, 1)

tmp21490 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons_2, tmp21489)
}
__typedArg0 := symcons_2
__typedArg1 := tmp21489
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21491 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21490)


_ = tmp21491

tmp21492 := MakeNative(func(__e *ControlFlow) {
Y1197 := __e.Get(1)
_ = Y1197
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1198 := __e.Get(1)
_ = Y1198
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(Y1197)
__typedS1, __typedOK1 := TypedString(Y1198)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := Y1197
__typedArg1 := Y1198
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
return
}, 1))
return
}, 1)

tmp21493 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcn, tmp21492)
}
__typedArg0 := symcn
__typedArg1 := tmp21492
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21494 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21493)


_ = tmp21494

tmp21495 := MakeNative(func(__e *ControlFlow) {
Y1196 := __e.Get(1)
_ = Y1196
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symclose) {
return PrimCloseStream(Y1196)
}
__typedArg0 := Y1196
return Call(__e, PrimFunc(symclose), __typedArg0)
})())
return
}, 1)

tmp21496 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symclose, tmp21495)
}
__typedArg0 := symclose
__typedArg1 := tmp21495
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21497 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21496)


_ = tmp21497

tmp21498 := MakeNative(func(__e *ControlFlow) {
Y1194 := __e.Get(1)
_ = Y1194
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1195 := __e.Get(1)
_ = Y1195
__e.TailApply(PrimFunc(symdeclare), Y1194, Y1195)
return
}, 1))
return
}, 1)

tmp21499 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdeclare, tmp21498)
}
__typedArg0 := symdeclare
__typedArg1 := tmp21498
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21500 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21499)


_ = tmp21500

tmp21501 := MakeNative(func(__e *ControlFlow) {
Y1193 := __e.Get(1)
_ = Y1193
__e.TailApply(PrimFunc(symdestroy), Y1193)
return
}, 1)

tmp21502 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdestroy, tmp21501)
}
__typedArg0 := symdestroy
__typedArg1 := tmp21501
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21503 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21502)


_ = tmp21503

tmp21504 := MakeNative(func(__e *ControlFlow) {
Y1191 := __e.Get(1)
_ = Y1191
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1192 := __e.Get(1)
_ = Y1192
__e.TailApply(PrimFunc(symdifference), Y1191, Y1192)
return
}, 1))
return
}, 1)

tmp21505 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdifference, tmp21504)
}
__typedArg0 := symdifference
__typedArg1 := tmp21504
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21506 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21505)


_ = tmp21506

tmp21507 := MakeNative(func(__e *ControlFlow) {
Y1189 := __e.Get(1)
_ = Y1189
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1190 := __e.Get(1)
_ = Y1190
_ = Y1189

__e.Return(Y1190)
return


}, 1))
return
}, 1)

tmp21508 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdo, tmp21507)
}
__typedArg0 := symdo
__typedArg1 := tmp21507
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21509 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21508)


_ = tmp21509

tmp21510 := MakeNative(func(__e *ControlFlow) {
Y1187 := __e.Get(1)
_ = Y1187
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1188 := __e.Get(1)
_ = Y1188
__e.TailApply(PrimFunc(symelement_2), Y1187, Y1188)
return
}, 1))
return
}, 1)

tmp21511 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symelement_2, tmp21510)
}
__typedArg0 := symelement_2
__typedArg1 := tmp21510
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21512 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21511)


_ = tmp21512

tmp21513 := MakeNative(func(__e *ControlFlow) {
Y1186 := __e.Get(1)
_ = Y1186
__e.TailApply(PrimFunc(symempty_2), Y1186)
return
}, 1)

tmp21514 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symempty_2, tmp21513)
}
__typedArg0 := symempty_2
__typedArg1 := tmp21513
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21515 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21514)


_ = tmp21515

tmp21516 := MakeNative(func(__e *ControlFlow) {
Y1185 := __e.Get(1)
_ = Y1185
__e.TailApply(PrimFunc(symenable_1type_1theory), Y1185)
return
}, 1)

tmp21517 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symenable_1type_1theory, tmp21516)
}
__typedArg0 := symenable_1type_1theory
__typedArg1 := tmp21516
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21518 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21517)


_ = tmp21518

tmp21519 := MakeNative(func(__e *ControlFlow) {
Y1184 := __e.Get(1)
_ = Y1184
__e.TailApply(PrimFunc(symexternal), Y1184)
return
}, 1)

tmp21520 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symexternal, tmp21519)
}
__typedArg0 := symexternal
__typedArg1 := tmp21519
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21521 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21520)


_ = tmp21521

tmp21522 := MakeNative(func(__e *ControlFlow) {
Y1183 := __e.Get(1)
_ = Y1183
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symerror_1to_1string) {
return PrimErrorToString(Y1183)
}
__typedArg0 := Y1183
return Call(__e, PrimFunc(symerror_1to_1string), __typedArg0)
})())
return
}, 1)

tmp21523 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symerror_1to_1string, tmp21522)
}
__typedArg0 := symerror_1to_1string
__typedArg1 := tmp21522
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21524 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21523)


_ = tmp21524

tmp21525 := MakeNative(func(__e *ControlFlow) {
Y1182 := __e.Get(1)
_ = Y1182
__e.TailApply(PrimFunc(symeval), Y1182)
return
}, 1)

tmp21526 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symeval, tmp21525)
}
__typedArg0 := symeval
__typedArg1 := tmp21525
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21527 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21526)


_ = tmp21527

tmp21528 := MakeNative(func(__e *ControlFlow) {
Y1181 := __e.Get(1)
_ = Y1181
__e.TailApply(PrimFunc(symeval_1kl), Y1181)
return
}, 1)

tmp21529 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symeval_1kl, tmp21528)
}
__typedArg0 := symeval_1kl
__typedArg1 := tmp21528
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21530 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21529)


_ = tmp21530

tmp21531 := MakeNative(func(__e *ControlFlow) {
Y1180 := __e.Get(1)
_ = Y1180
__e.TailApply(PrimFunc(symexplode), Y1180)
return
}, 1)

tmp21532 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symexplode, tmp21531)
}
__typedArg0 := symexplode
__typedArg1 := tmp21531
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21533 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21532)


_ = tmp21533

tmp21534 := MakeNative(func(__e *ControlFlow) {
Y1179 := __e.Get(1)
_ = Y1179
__e.TailApply(PrimFunc(symexternal), Y1179)
return
}, 1)

tmp21535 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symexternal, tmp21534)
}
__typedArg0 := symexternal
__typedArg1 := tmp21534
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21536 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21535)


_ = tmp21536

tmp21537 := MakeNative(func(__e *ControlFlow) {
Y1178 := __e.Get(1)
_ = Y1178
__e.TailApply(PrimFunc(symfactorise), Y1178)
return
}, 1)

tmp21538 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfactorise, tmp21537)
}
__typedArg0 := symfactorise
__typedArg1 := tmp21537
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21539 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21538)


_ = tmp21539

tmp21540 := MakeNative(func(__e *ControlFlow) {
Y1176 := __e.Get(1)
_ = Y1176
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1177 := __e.Get(1)
_ = Y1177
__e.TailApply(PrimFunc(symfail_1if), Y1176, Y1177)
return
}, 1))
return
}, 1)

tmp21541 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfail_1if, tmp21540)
}
__typedArg0 := symfail_1if
__typedArg1 := tmp21540
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21542 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21541)


_ = tmp21542

tmp21543 := MakeNative(func(__e *ControlFlow) {
Y1174 := __e.Get(1)
_ = Y1174
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1175 := __e.Get(1)
_ = Y1175
__e.TailApply(PrimFunc(symfix), Y1174, Y1175)
return
}, 1))
return
}, 1)

tmp21544 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfix, tmp21543)
}
__typedArg0 := symfix
__typedArg1 := tmp21543
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21545 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21544)


_ = tmp21545

tmp21546 := MakeNative(func(__e *ControlFlow) {
Y1167 := __e.Get(1)
_ = Y1167
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1168 := __e.Get(1)
_ = Y1168
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1169 := __e.Get(1)
_ = Y1169
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1170 := __e.Get(1)
_ = Y1170
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1171 := __e.Get(1)
_ = Y1171
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1172 := __e.Get(1)
_ = Y1172
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1173 := __e.Get(1)
_ = Y1173
__e.TailApply(PrimFunc(symfindall), Y1167, Y1168, Y1169, Y1170, Y1171, Y1172, Y1173)
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

tmp21547 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfindall, tmp21546)
}
__typedArg0 := symfindall
__typedArg1 := tmp21546
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21548 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21547)


_ = tmp21548

tmp21549 := MakeNative(func(__e *ControlFlow) {
Y1162 := __e.Get(1)
_ = Y1162
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1163 := __e.Get(1)
_ = Y1163
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1164 := __e.Get(1)
_ = Y1164
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1165 := __e.Get(1)
_ = Y1165
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1166 := __e.Get(1)
_ = Y1166
__e.TailApply(PrimFunc(symfork), Y1162, Y1163, Y1164, Y1165, Y1166)
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21550 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfork, tmp21549)
}
__typedArg0 := symfork
__typedArg1 := tmp21549
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21551 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21550)


_ = tmp21551

tmp21552 := MakeNative(func(__e *ControlFlow) {
Y1161 := __e.Get(1)
_ = Y1161
__e.Return(MakeNative(func(__e *ControlFlow) {
__e.Return(Y1161)
return
}, 0))
return
}, 1)

tmp21553 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfreeze, tmp21552)
}
__typedArg0 := symfreeze
__typedArg1 := tmp21552
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21554 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21553)


_ = tmp21554

tmp21555 := MakeNative(func(__e *ControlFlow) {
Y1160 := __e.Get(1)
_ = Y1160
__e.TailApply(PrimFunc(symfst), Y1160)
return
}, 1)

tmp21556 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfst, tmp21555)
}
__typedArg0 := symfst
__typedArg1 := tmp21555
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21557 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21556)


_ = tmp21557

tmp21558 := MakeNative(func(__e *ControlFlow) {
Y1159 := __e.Get(1)
_ = Y1159
__e.TailApply(PrimFunc(symfn), Y1159)
return
}, 1)

tmp21559 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfn, tmp21558)
}
__typedArg0 := symfn
__typedArg1 := tmp21558
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21560 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21559)


_ = tmp21560

tmp21561 := MakeNative(func(__e *ControlFlow) {
Y1158 := __e.Get(1)
_ = Y1158
__e.TailApply(PrimFunc(symfunction), Y1158)
return
}, 1)

tmp21562 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfunction, tmp21561)
}
__typedArg0 := symfunction
__typedArg1 := tmp21561
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21563 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21562)


_ = tmp21563

tmp21564 := MakeNative(func(__e *ControlFlow) {
Y1157 := __e.Get(1)
_ = Y1157
__e.TailApply(PrimFunc(symgensym), Y1157)
return
}, 1)

tmp21565 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symgensym, tmp21564)
}
__typedArg0 := symgensym
__typedArg1 := tmp21564
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21566 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21565)


_ = tmp21566

tmp21567 := MakeNative(func(__e *ControlFlow) {
Y1154 := __e.Get(1)
_ = Y1154
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1155 := __e.Get(1)
_ = Y1155
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1156 := __e.Get(1)
_ = Y1156
__e.TailApply(PrimFunc(symget), Y1154, Y1155, Y1156)
return
}, 1))
return
}, 1))
return
}, 1)

tmp21568 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symget, tmp21567)
}
__typedArg0 := symget
__typedArg1 := tmp21567
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21569 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21568)


_ = tmp21569

tmp21570 := MakeNative(func(__e *ControlFlow) {
Y1153 := __e.Get(1)
_ = Y1153
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symget_1time) {
return PrimGetTime(Y1153)
}
__typedArg0 := Y1153
return Call(__e, PrimFunc(symget_1time), __typedArg0)
})())
return
}, 1)

tmp21571 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symget_1time, tmp21570)
}
__typedArg0 := symget_1time
__typedArg1 := tmp21570
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21572 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21571)


_ = tmp21572

tmp21573 := MakeNative(func(__e *ControlFlow) {
Y1150 := __e.Get(1)
_ = Y1150
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1151 := __e.Get(1)
_ = Y1151
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1152 := __e.Get(1)
_ = Y1152
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(Y1150, Y1151, Y1152)
}
__typedArg0 := Y1150
__typedArg1 := Y1151
__typedArg2 := Y1152
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})())
return
}, 1))
return
}, 1))
return
}, 1)

tmp21574 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symaddress_1_6, tmp21573)
}
__typedArg0 := symaddress_1_6
__typedArg1 := tmp21573
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21575 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21574)


_ = tmp21575

tmp21576 := MakeNative(func(__e *ControlFlow) {
Y1148 := __e.Get(1)
_ = Y1148
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1149 := __e.Get(1)
_ = Y1149
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(Y1148, Y1149)
}
__typedArg0 := Y1148
__typedArg1 := Y1149
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})())
return
}, 1))
return
}, 1)

tmp21577 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1address, tmp21576)
}
__typedArg0 := sym_5_1address
__typedArg1 := tmp21576
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21578 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21577)


_ = tmp21578

tmp21579 := MakeNative(func(__e *ControlFlow) {
Y1146 := __e.Get(1)
_ = Y1146
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1147 := __e.Get(1)
_ = Y1147
__e.TailApply(PrimFunc(sym_5_1vector), Y1146, Y1147)
return
}, 1))
return
}, 1)

tmp21580 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1vector, tmp21579)
}
__typedArg0 := sym_5_1vector
__typedArg1 := tmp21579
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21581 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21580)


_ = tmp21581

tmp21582 := MakeNative(func(__e *ControlFlow) {
Y1144 := __e.Get(1)
_ = Y1144
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1145 := __e.Get(1)
_ = Y1145
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_6) {
__typedN0, __typedOK0 := TypedFloat64(Y1144)
__typedN1, __typedOK1 := TypedFloat64(Y1145)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_6) {
return TypedMaterializeBoolean((__typedN0 > __typedN1))
}}
__typedArg0 := Y1144
__typedArg1 := Y1145
return Call(__e, PrimFunc(sym_6), __typedArg0, __typedArg1)
})())
return
}, 1))
return
}, 1)

tmp21583 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_6, tmp21582)
}
__typedArg0 := sym_6
__typedArg1 := tmp21582
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21584 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21583)


_ = tmp21584

tmp21585 := MakeNative(func(__e *ControlFlow) {
Y1142 := __e.Get(1)
_ = Y1142
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1143 := __e.Get(1)
_ = Y1143
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_6_a) {
__typedN0, __typedOK0 := TypedFloat64(Y1142)
__typedN1, __typedOK1 := TypedFloat64(Y1143)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_6_a) {
return TypedMaterializeBoolean((__typedN0 >= __typedN1))
}}
__typedArg0 := Y1142
__typedArg1 := Y1143
return Call(__e, PrimFunc(sym_6_a), __typedArg0, __typedArg1)
})())
return
}, 1))
return
}, 1)

tmp21586 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_6_a, tmp21585)
}
__typedArg0 := sym_6_a
__typedArg1 := tmp21585
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21587 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21586)


_ = tmp21587

tmp21588 := MakeNative(func(__e *ControlFlow) {
Y1140 := __e.Get(1)
_ = Y1140
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1141 := __e.Get(1)
_ = Y1141
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Y1140, Y1141)
}
__typedArg0 := Y1140
__typedArg1 := Y1141
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})())
return
}, 1))
return
}, 1)

tmp21589 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_a, tmp21588)
}
__typedArg0 := sym_a
__typedArg1 := tmp21588
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21590 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21589)


_ = tmp21590

tmp21591 := MakeNative(func(__e *ControlFlow) {
Y1138 := __e.Get(1)
_ = Y1138
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1139 := __e.Get(1)
_ = Y1139
__e.TailApply(PrimFunc(symhash), Y1138, Y1139)
return
}, 1))
return
}, 1)

tmp21592 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhash, tmp21591)
}
__typedArg0 := symhash
__typedArg1 := tmp21591
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21593 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21592)


_ = tmp21593

tmp21594 := MakeNative(func(__e *ControlFlow) {
Y1137 := __e.Get(1)
_ = Y1137
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(Y1137)
}
__typedArg0 := Y1137
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return
}, 1)

tmp21595 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhd, tmp21594)
}
__typedArg0 := symhd
__typedArg1 := tmp21594
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21596 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21595)


_ = tmp21596

tmp21597 := MakeNative(func(__e *ControlFlow) {
Y1136 := __e.Get(1)
_ = Y1136
__e.TailApply(PrimFunc(symhdv), Y1136)
return
}, 1)

tmp21598 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhdv, tmp21597)
}
__typedArg0 := symhdv
__typedArg1 := tmp21597
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21599 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21598)


_ = tmp21599

tmp21600 := MakeNative(func(__e *ControlFlow) {
Y1135 := __e.Get(1)
_ = Y1135
__e.TailApply(PrimFunc(symhdstr), Y1135)
return
}, 1)

tmp21601 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhdstr, tmp21600)
}
__typedArg0 := symhdstr
__typedArg1 := tmp21600
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21602 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21601)


_ = tmp21602

tmp21603 := MakeNative(func(__e *ControlFlow) {
Y1134 := __e.Get(1)
_ = Y1134
__e.TailApply(PrimFunc(symhead), Y1134)
return
}, 1)

tmp21604 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhead, tmp21603)
}
__typedArg0 := symhead
__typedArg1 := tmp21603
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21605 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21604)


_ = tmp21605

tmp21606 := MakeNative(func(__e *ControlFlow) {
Y1133 := __e.Get(1)
_ = Y1133
__e.TailApply(PrimFunc(symhush), Y1133)
return
}, 1)

tmp21607 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhush, tmp21606)
}
__typedArg0 := symhush
__typedArg1 := tmp21606
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21608 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21607)


_ = tmp21608

tmp21609 := MakeNative(func(__e *ControlFlow) {
Y1130 := __e.Get(1)
_ = Y1130
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1131 := __e.Get(1)
_ = Y1131
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1132 := __e.Get(1)
_ = Y1132
if True == Y1130 {
__e.Return(Y1131)
return
} else {
__e.Return(Y1132)
return
}
}, 1))
return
}, 1))
return
}, 1)

tmp21611 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp21609)
}
__typedArg0 := symif
__typedArg1 := tmp21609
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21612 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21611)


_ = tmp21612

tmp21613 := MakeNative(func(__e *ControlFlow) {
Y1129 := __e.Get(1)
_ = Y1129
__e.TailApply(PrimFunc(syminclude), Y1129)
return
}, 1)

tmp21614 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminclude, tmp21613)
}
__typedArg0 := syminclude
__typedArg1 := tmp21613
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21615 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21614)


_ = tmp21615

tmp21616 := MakeNative(func(__e *ControlFlow) {
Y1128 := __e.Get(1)
_ = Y1128
__e.TailApply(PrimFunc(symin_1package), Y1128)
return
}, 1)

tmp21617 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symin_1package, tmp21616)
}
__typedArg0 := symin_1package
__typedArg1 := tmp21616
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21618 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21617)


_ = tmp21618

tmp21619 := MakeNative(func(__e *ControlFlow) {
Y1127 := __e.Get(1)
_ = Y1127
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(syminteger_2) {
return PrimIsInteger(Y1127)
}
__typedArg0 := Y1127
return Call(__e, PrimFunc(syminteger_2), __typedArg0)
})())
return
}, 1)

tmp21620 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminteger_2, tmp21619)
}
__typedArg0 := syminteger_2
__typedArg1 := tmp21619
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21621 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21620)


_ = tmp21621

tmp21622 := MakeNative(func(__e *ControlFlow) {
Y1126 := __e.Get(1)
_ = Y1126
__e.TailApply(PrimFunc(syminternal), Y1126)
return
}, 1)

tmp21623 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminternal, tmp21622)
}
__typedArg0 := syminternal
__typedArg1 := tmp21622
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21624 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21623)


_ = tmp21624

tmp21625 := MakeNative(func(__e *ControlFlow) {
Y1125 := __e.Get(1)
_ = Y1125
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(Y1125)
}
__typedArg0 := Y1125
return Call(__e, PrimFunc(symintern), __typedArg0)
})())
return
}, 1)

tmp21626 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symintern, tmp21625)
}
__typedArg0 := symintern
__typedArg1 := tmp21625
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21627 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21626)


_ = tmp21627

tmp21628 := MakeNative(func(__e *ControlFlow) {
Y1124 := __e.Get(1)
_ = Y1124
__e.TailApply(PrimFunc(syminput), Y1124)
return
}, 1)

tmp21629 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminput, tmp21628)
}
__typedArg0 := syminput
__typedArg1 := tmp21628
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21630 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21629)


_ = tmp21630

tmp21631 := MakeNative(func(__e *ControlFlow) {
Y1122 := __e.Get(1)
_ = Y1122
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1123 := __e.Get(1)
_ = Y1123
__e.TailApply(PrimFunc(symshen_4input_1h_7), Y1122, Y1123)
return
}, 1))
return
}, 1)

tmp21632 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminput_7, tmp21631)
}
__typedArg0 := syminput_7
__typedArg1 := tmp21631
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21633 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21632)


_ = tmp21633

tmp21634 := MakeNative(func(__e *ControlFlow) {
Y1121 := __e.Get(1)
_ = Y1121
__e.TailApply(PrimFunc(syminclude_1all_1but), Y1121)
return
}, 1)

tmp21635 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminclude_1all_1but, tmp21634)
}
__typedArg0 := syminclude_1all_1but
__typedArg1 := tmp21634
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21636 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21635)


_ = tmp21636

tmp21637 := MakeNative(func(__e *ControlFlow) {
Y1119 := __e.Get(1)
_ = Y1119
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1120 := __e.Get(1)
_ = Y1120
__e.TailApply(PrimFunc(symintersection), Y1119, Y1120)
return
}, 1))
return
}, 1)

tmp21638 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symintersection, tmp21637)
}
__typedArg0 := symintersection
__typedArg1 := tmp21637
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21639 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21638)


_ = tmp21639

tmp21640 := MakeNative(func(__e *ControlFlow) {
Y1118 := __e.Get(1)
_ = Y1118
__e.TailApply(PrimFunc(syminternal), Y1118)
return
}, 1)

tmp21641 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminternal, tmp21640)
}
__typedArg0 := syminternal
__typedArg1 := tmp21640
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21642 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21641)


_ = tmp21642

tmp21643 := MakeNative(func(__e *ControlFlow) {
Y1112 := __e.Get(1)
_ = Y1112
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1113 := __e.Get(1)
_ = Y1113
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1114 := __e.Get(1)
_ = Y1114
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1115 := __e.Get(1)
_ = Y1115
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1116 := __e.Get(1)
_ = Y1116
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1117 := __e.Get(1)
_ = Y1117
__e.TailApply(PrimFunc(symis), Y1112, Y1113, Y1114, Y1115, Y1116, Y1117)
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

tmp21644 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symis, tmp21643)
}
__typedArg0 := symis
__typedArg1 := tmp21643
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21645 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21644)


_ = tmp21645

tmp21646 := MakeNative(func(__e *ControlFlow) {
Y1106 := __e.Get(1)
_ = Y1106
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1107 := __e.Get(1)
_ = Y1107
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1108 := __e.Get(1)
_ = Y1108
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1109 := __e.Get(1)
_ = Y1109
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1110 := __e.Get(1)
_ = Y1110
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1111 := __e.Get(1)
_ = Y1111
__e.TailApply(PrimFunc(symis_b), Y1106, Y1107, Y1108, Y1109, Y1110, Y1111)
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

tmp21647 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symis_b, tmp21646)
}
__typedArg0 := symis_b
__typedArg1 := tmp21646
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21648 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21647)


_ = tmp21648

tmp21649 := MakeNative(func(__e *ControlFlow) {
Y1105 := __e.Get(1)
_ = Y1105
__e.TailApply(PrimFunc(symlength), Y1105)
return
}, 1)

tmp21650 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlength, tmp21649)
}
__typedArg0 := symlength
__typedArg1 := tmp21649
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21651 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21650)


_ = tmp21651

tmp21652 := MakeNative(func(__e *ControlFlow) {
Y1104 := __e.Get(1)
_ = Y1104
__e.TailApply(PrimFunc(symlimit), Y1104)
return
}, 1)

tmp21653 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlimit, tmp21652)
}
__typedArg0 := symlimit
__typedArg1 := tmp21652
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21654 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21653)


_ = tmp21654

tmp21655 := MakeNative(func(__e *ControlFlow) {
Y1103 := __e.Get(1)
_ = Y1103
__e.TailApply(PrimFunc(symlineread), Y1103)
return
}, 1)

tmp21656 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlineread, tmp21655)
}
__typedArg0 := symlineread
__typedArg1 := tmp21655
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21657 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21656)


_ = tmp21657

tmp21658 := MakeNative(func(__e *ControlFlow) {
Y1102 := __e.Get(1)
_ = Y1102
__e.TailApply(PrimFunc(symload), Y1102)
return
}, 1)

tmp21659 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symload, tmp21658)
}
__typedArg0 := symload
__typedArg1 := tmp21658
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21660 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21659)


_ = tmp21660

tmp21661 := MakeNative(func(__e *ControlFlow) {
Y1100 := __e.Get(1)
_ = Y1100
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1101 := __e.Get(1)
_ = Y1101
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5) {
__typedN0, __typedOK0 := TypedFloat64(Y1100)
__typedN1, __typedOK1 := TypedFloat64(Y1101)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_5) {
return TypedMaterializeBoolean((__typedN0 < __typedN1))
}}
__typedArg0 := Y1100
__typedArg1 := Y1101
return Call(__e, PrimFunc(sym_5), __typedArg0, __typedArg1)
})())
return
}, 1))
return
}, 1)

tmp21662 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5, tmp21661)
}
__typedArg0 := sym_5
__typedArg1 := tmp21661
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21663 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21662)


_ = tmp21663

tmp21664 := MakeNative(func(__e *ControlFlow) {
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

tmp21665 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_a, tmp21664)
}
__typedArg0 := sym_5_a
__typedArg1 := tmp21664
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21666 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21665)


_ = tmp21666

tmp21667 := MakeNative(func(__e *ControlFlow) {
Y1097 := __e.Get(1)
_ = Y1097
__e.TailApply(PrimFunc(symvector), Y1097)
return
}, 1)

tmp21668 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp21667)
}
__typedArg0 := symvector
__typedArg1 := tmp21667
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21669 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21668)


_ = tmp21669

tmp21670 := MakeNative(func(__e *ControlFlow) {
Y1096 := __e.Get(1)
_ = Y1096
__e.TailApply(PrimFunc(symmacroexpand), Y1096)
return
}, 1)

tmp21671 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symmacroexpand, tmp21670)
}
__typedArg0 := symmacroexpand
__typedArg1 := tmp21670
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21672 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21671)


_ = tmp21672

tmp21673 := MakeNative(func(__e *ControlFlow) {
Y1094 := __e.Get(1)
_ = Y1094
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1095 := __e.Get(1)
_ = Y1095
__e.TailApply(PrimFunc(symmap), Y1094, Y1095)
return
}, 1))
return
}, 1)

tmp21674 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symmap, tmp21673)
}
__typedArg0 := symmap
__typedArg1 := tmp21673
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21675 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21674)


_ = tmp21675

tmp21676 := MakeNative(func(__e *ControlFlow) {
Y1092 := __e.Get(1)
_ = Y1092
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1093 := __e.Get(1)
_ = Y1093
__e.TailApply(PrimFunc(symmapcan), Y1092, Y1093)
return
}, 1))
return
}, 1)

tmp21677 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symmapcan, tmp21676)
}
__typedArg0 := symmapcan
__typedArg1 := tmp21676
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21678 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21677)


_ = tmp21678

tmp21679 := MakeNative(func(__e *ControlFlow) {
Y1091 := __e.Get(1)
_ = Y1091
__e.TailApply(PrimFunc(symmaxinferences), Y1091)
return
}, 1)

tmp21680 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symmaxinferences, tmp21679)
}
__typedArg0 := symmaxinferences
__typedArg1 := tmp21679
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21681 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21680)


_ = tmp21681

tmp21682 := MakeNative(func(__e *ControlFlow) {
Y1090 := __e.Get(1)
_ = Y1090
__e.TailApply(PrimFunc(symnl), Y1090)
return
}, 1)

tmp21683 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnl, tmp21682)
}
__typedArg0 := symnl
__typedArg1 := tmp21682
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21684 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21683)


_ = tmp21684

tmp21685 := MakeNative(func(__e *ControlFlow) {
Y1089 := __e.Get(1)
_ = Y1089
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(Y1089)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := Y1089
return Call(__e, PrimFunc(symnot), __typedArg0)
})())
return
}, 1)

tmp21686 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnot, tmp21685)
}
__typedArg0 := symnot
__typedArg1 := tmp21685
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21687 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21686)


_ = tmp21687

tmp21688 := MakeNative(func(__e *ControlFlow) {
Y1087 := __e.Get(1)
_ = Y1087
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1088 := __e.Get(1)
_ = Y1088
__e.TailApply(PrimFunc(symnth), Y1087, Y1088)
return
}, 1))
return
}, 1)

tmp21689 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnth, tmp21688)
}
__typedArg0 := symnth
__typedArg1 := tmp21688
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21690 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21689)


_ = tmp21690

tmp21691 := MakeNative(func(__e *ControlFlow) {
Y1086 := __e.Get(1)
_ = Y1086
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symn_1_6string) {
return PrimNumberToString(Y1086)
}
__typedArg0 := Y1086
return Call(__e, PrimFunc(symn_1_6string), __typedArg0)
})())
return
}, 1)

tmp21692 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symn_1_6string, tmp21691)
}
__typedArg0 := symn_1_6string
__typedArg1 := tmp21691
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21693 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21692)


_ = tmp21693

tmp21694 := MakeNative(func(__e *ControlFlow) {
Y1085 := __e.Get(1)
_ = Y1085
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnumber_2) {
return PrimIsNumber(Y1085)
}
__typedArg0 := Y1085
return Call(__e, PrimFunc(symnumber_2), __typedArg0)
})())
return
}, 1)

tmp21695 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber_2, tmp21694)
}
__typedArg0 := symnumber_2
__typedArg1 := tmp21694
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21696 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21695)


_ = tmp21696

tmp21697 := MakeNative(func(__e *ControlFlow) {
Y1084 := __e.Get(1)
_ = Y1084
__e.TailApply(PrimFunc(symoccurs_1check), Y1084)
return
}, 1)

tmp21698 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symoccurs_1check, tmp21697)
}
__typedArg0 := symoccurs_1check
__typedArg1 := tmp21697
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21699 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21698)


_ = tmp21699

tmp21700 := MakeNative(func(__e *ControlFlow) {
Y1082 := __e.Get(1)
_ = Y1082
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1083 := __e.Get(1)
_ = Y1083
__e.TailApply(PrimFunc(symoccurrences), Y1082, Y1083)
return
}, 1))
return
}, 1)

tmp21701 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symoccurrences, tmp21700)
}
__typedArg0 := symoccurrences
__typedArg1 := tmp21700
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21702 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21701)


_ = tmp21702

tmp21703 := MakeNative(func(__e *ControlFlow) {
Y1081 := __e.Get(1)
_ = Y1081
__e.TailApply(PrimFunc(symoccurs_1check), Y1081)
return
}, 1)

tmp21704 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symoccurs_1check, tmp21703)
}
__typedArg0 := symoccurs_1check
__typedArg1 := tmp21703
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21705 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21704)


_ = tmp21705

tmp21706 := MakeNative(func(__e *ControlFlow) {
Y1079 := __e.Get(1)
_ = Y1079
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1080 := __e.Get(1)
_ = Y1080
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symopen) {
return PrimOpenStream(Y1079, Y1080)
}
__typedArg0 := Y1079
__typedArg1 := Y1080
return Call(__e, PrimFunc(symopen), __typedArg0, __typedArg1)
})())
return
}, 1))
return
}, 1)

tmp21707 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symopen, tmp21706)
}
__typedArg0 := symopen
__typedArg1 := tmp21706
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21708 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21707)


_ = tmp21708

tmp21709 := MakeNative(func(__e *ControlFlow) {
Y1078 := __e.Get(1)
_ = Y1078
__e.TailApply(PrimFunc(symoptimise), Y1078)
return
}, 1)

tmp21710 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symoptimise, tmp21709)
}
__typedArg0 := symoptimise
__typedArg1 := tmp21709
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21711 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21710)


_ = tmp21711

tmp21712 := MakeNative(func(__e *ControlFlow) {
Y1076 := __e.Get(1)
_ = Y1076
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1077 := __e.Get(1)
_ = Y1077
if True == Y1076 {
__e.Return(True)
return
} else {
if True == Y1077 {
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

tmp21715 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symor, tmp21712)
}
__typedArg0 := symor
__typedArg1 := tmp21712
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21716 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21715)


_ = tmp21716

tmp21717 := MakeNative(func(__e *ControlFlow) {
Y1075 := __e.Get(1)
_ = Y1075
__e.TailApply(PrimFunc(sympackage_2), Y1075)
return
}, 1)

tmp21718 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympackage_2, tmp21717)
}
__typedArg0 := sympackage_2
__typedArg1 := tmp21717
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21719 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21718)


_ = tmp21719

tmp21720 := MakeNative(func(__e *ControlFlow) {
Y1073 := __e.Get(1)
_ = Y1073
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1074 := __e.Get(1)
_ = Y1074
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sympos) {
return PrimPos(Y1073, Y1074)
}
__typedArg0 := Y1073
__typedArg1 := Y1074
return Call(__e, PrimFunc(sympos), __typedArg0, __typedArg1)
})())
return
}, 1))
return
}, 1)

tmp21721 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympos, tmp21720)
}
__typedArg0 := sympos
__typedArg1 := tmp21720
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21722 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21721)


_ = tmp21722

tmp21723 := MakeNative(func(__e *ControlFlow) {
Y1072 := __e.Get(1)
_ = Y1072
__e.TailApply(PrimFunc(sympreclude_1all_1but), Y1072)
return
}, 1)

tmp21724 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympreclude_1all_1but, tmp21723)
}
__typedArg0 := sympreclude_1all_1but
__typedArg1 := tmp21723
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21725 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21724)


_ = tmp21725

tmp21726 := MakeNative(func(__e *ControlFlow) {
Y1071 := __e.Get(1)
_ = Y1071
__e.TailApply(PrimFunc(symprint), Y1071)
return
}, 1)

tmp21727 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symprint, tmp21726)
}
__typedArg0 := symprint
__typedArg1 := tmp21726
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21728 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21727)


_ = tmp21728

tmp21729 := MakeNative(func(__e *ControlFlow) {
Y1070 := __e.Get(1)
_ = Y1070
__e.TailApply(PrimFunc(symprofile), Y1070)
return
}, 1)

tmp21730 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symprofile, tmp21729)
}
__typedArg0 := symprofile
__typedArg1 := tmp21729
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21731 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21730)


_ = tmp21731

tmp21732 := MakeNative(func(__e *ControlFlow) {
Y1069 := __e.Get(1)
_ = Y1069
__e.TailApply(PrimFunc(symshen_4print_1prolog_1vector), Y1069)
return
}, 1)

tmp21733 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4print_1prolog_1vector, tmp21732)
}
__typedArg0 := symshen_4print_1prolog_1vector
__typedArg1 := tmp21732
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21734 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21733)


_ = tmp21734

tmp21735 := MakeNative(func(__e *ControlFlow) {
Y1068 := __e.Get(1)
_ = Y1068
__e.TailApply(PrimFunc(symshen_4print_1freshterm), Y1068)
return
}, 1)

tmp21736 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4print_1freshterm, tmp21735)
}
__typedArg0 := symshen_4print_1freshterm
__typedArg1 := tmp21735
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21737 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21736)


_ = tmp21737

tmp21738 := MakeNative(func(__e *ControlFlow) {
Y1067 := __e.Get(1)
_ = Y1067
__e.TailApply(PrimFunc(symshen_4printF), Y1067)
return
}, 1)

tmp21739 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4printF, tmp21738)
}
__typedArg0 := symshen_4printF
__typedArg1 := tmp21738
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21740 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21739)


_ = tmp21740

tmp21741 := MakeNative(func(__e *ControlFlow) {
Y1066 := __e.Get(1)
_ = Y1066
__e.TailApply(PrimFunc(symprolog_1memory), Y1066)
return
}, 1)

tmp21742 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symprolog_1memory, tmp21741)
}
__typedArg0 := symprolog_1memory
__typedArg1 := tmp21741
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21743 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21742)


_ = tmp21743

tmp21744 := MakeNative(func(__e *ControlFlow) {
Y1065 := __e.Get(1)
_ = Y1065
__e.TailApply(PrimFunc(symprofile_1results), Y1065)
return
}, 1)

tmp21745 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symprofile_1results, tmp21744)
}
__typedArg0 := symprofile_1results
__typedArg1 := tmp21744
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21746 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21745)


_ = tmp21746

tmp21747 := MakeNative(func(__e *ControlFlow) {
Y1063 := __e.Get(1)
_ = Y1063
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1064 := __e.Get(1)
_ = Y1064
__e.TailApply(PrimFunc(sympr), Y1063, Y1064)
return
}, 1))
return
}, 1)

tmp21748 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympr, tmp21747)
}
__typedArg0 := sympr
__typedArg1 := tmp21747
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21749 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21748)


_ = tmp21749

tmp21750 := MakeNative(func(__e *ControlFlow) {
Y1062 := __e.Get(1)
_ = Y1062
__e.TailApply(PrimFunc(symps), Y1062)
return
}, 1)

tmp21751 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symps, tmp21750)
}
__typedArg0 := symps
__typedArg1 := tmp21750
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21752 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21751)


_ = tmp21752

tmp21753 := MakeNative(func(__e *ControlFlow) {
Y1061 := __e.Get(1)
_ = Y1061
__e.TailApply(PrimFunc(sympreclude), Y1061)
return
}, 1)

tmp21754 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympreclude, tmp21753)
}
__typedArg0 := sympreclude
__typedArg1 := tmp21753
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21755 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21754)


_ = tmp21755

tmp21756 := MakeNative(func(__e *ControlFlow) {
Y1060 := __e.Get(1)
_ = Y1060
__e.TailApply(PrimFunc(sympreclude_1all_1but), Y1060)
return
}, 1)

tmp21757 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympreclude_1all_1but, tmp21756)
}
__typedArg0 := sympreclude_1all_1but
__typedArg1 := tmp21756
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21758 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21757)


_ = tmp21758

tmp21759 := MakeNative(func(__e *ControlFlow) {
Y1059 := __e.Get(1)
_ = Y1059
__e.TailApply(PrimFunc(symprotect), Y1059)
return
}, 1)

tmp21760 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symprotect, tmp21759)
}
__typedArg0 := symprotect
__typedArg1 := tmp21759
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21761 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21760)


_ = tmp21761

tmp21762 := MakeNative(func(__e *ControlFlow) {
Y1055 := __e.Get(1)
_ = Y1055
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1056 := __e.Get(1)
_ = Y1056
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1057 := __e.Get(1)
_ = Y1057
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1058 := __e.Get(1)
_ = Y1058
__e.TailApply(PrimFunc(symput), Y1055, Y1056, Y1057, Y1058)
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21763 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symput, tmp21762)
}
__typedArg0 := symput
__typedArg1 := tmp21762
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21764 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21763)


_ = tmp21764

tmp21765 := MakeNative(func(__e *ControlFlow) {
Y1054 := __e.Get(1)
_ = Y1054
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symread_1file_1as_1string) {
return PrimReadFileAsString(Y1054)
}
__typedArg0 := Y1054
return Call(__e, PrimFunc(symread_1file_1as_1string), __typedArg0)
})())
return
}, 1)

tmp21766 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1file_1as_1string, tmp21765)
}
__typedArg0 := symread_1file_1as_1string
__typedArg1 := tmp21765
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21767 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21766)


_ = tmp21767

tmp21768 := MakeNative(func(__e *ControlFlow) {
Y1053 := __e.Get(1)
_ = Y1053
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symread_1file_1as_1bytelist) {
return PrimReadFileAsByteList(Y1053)
}
__typedArg0 := Y1053
return Call(__e, PrimFunc(symread_1file_1as_1bytelist), __typedArg0)
})())
return
}, 1)

tmp21769 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1file_1as_1bytelist, tmp21768)
}
__typedArg0 := symread_1file_1as_1bytelist
__typedArg1 := tmp21768
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21770 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21769)


_ = tmp21770

tmp21771 := MakeNative(func(__e *ControlFlow) {
Y1052 := __e.Get(1)
_ = Y1052
__e.TailApply(PrimFunc(symread_1file), Y1052)
return
}, 1)

tmp21772 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1file, tmp21771)
}
__typedArg0 := symread_1file
__typedArg1 := tmp21771
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21773 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21772)


_ = tmp21773

tmp21774 := MakeNative(func(__e *ControlFlow) {
Y1051 := __e.Get(1)
_ = Y1051
__e.TailApply(PrimFunc(symread), Y1051)
return
}, 1)

tmp21775 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread, tmp21774)
}
__typedArg0 := symread
__typedArg1 := tmp21774
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21776 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21775)


_ = tmp21776

tmp21777 := MakeNative(func(__e *ControlFlow) {
Y1050 := __e.Get(1)
_ = Y1050
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symread_1byte) {
return PrimReadByte(Y1050)
}
__typedArg0 := Y1050
return Call(__e, PrimFunc(symread_1byte), __typedArg0)
})())
return
}, 1)

tmp21778 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1byte, tmp21777)
}
__typedArg0 := symread_1byte
__typedArg1 := tmp21777
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21779 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21778)


_ = tmp21779

tmp21780 := MakeNative(func(__e *ControlFlow) {
Y1049 := __e.Get(1)
_ = Y1049
__e.TailApply(PrimFunc(symread_1from_1string), Y1049)
return
}, 1)

tmp21781 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1from_1string, tmp21780)
}
__typedArg0 := symread_1from_1string
__typedArg1 := tmp21780
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21782 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21781)


_ = tmp21782

tmp21783 := MakeNative(func(__e *ControlFlow) {
Y1048 := __e.Get(1)
_ = Y1048
__e.TailApply(PrimFunc(symread_1from_1string_1unprocessed), Y1048)
return
}, 1)

tmp21784 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1from_1string_1unprocessed, tmp21783)
}
__typedArg0 := symread_1from_1string_1unprocessed
__typedArg1 := tmp21783
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21785 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21784)


_ = tmp21785

tmp21786 := MakeNative(func(__e *ControlFlow) {
Y1047 := __e.Get(1)
_ = Y1047
__e.TailApply(PrimFunc(symshen_4read_1unit_1string), Y1047)
return
}, 1)

tmp21787 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4read_1unit_1string, tmp21786)
}
__typedArg0 := symshen_4read_1unit_1string
__typedArg1 := tmp21786
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21788 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21787)


_ = tmp21788

tmp21789 := MakeNative(func(__e *ControlFlow) {
Y1045 := __e.Get(1)
_ = Y1045
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1046 := __e.Get(1)
_ = Y1046
__e.TailApply(PrimFunc(symshen_4register_1source_1form), Y1045, Y1046)
return
}, 1))
return
}, 1)

tmp21790 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4register_1source_1form, tmp21789)
}
__typedArg0 := symshen_4register_1source_1form
__typedArg1 := tmp21789
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21791 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21790)


_ = tmp21791

tmp21792 := MakeNative(func(__e *ControlFlow) {
Y1044 := __e.Get(1)
_ = Y1044
__e.TailApply(PrimFunc(symshen_4unregister_1source_1form), Y1044)
return
}, 1)

tmp21793 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4unregister_1source_1form, tmp21792)
}
__typedArg0 := symshen_4unregister_1source_1form
__typedArg1 := tmp21792
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21794 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21793)


_ = tmp21794

tmp21795 := MakeNative(func(__e *ControlFlow) {
Y1042 := __e.Get(1)
_ = Y1042
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1043 := __e.Get(1)
_ = Y1043
__e.TailApply(PrimFunc(symremove), Y1042, Y1043)
return
}, 1))
return
}, 1)

tmp21796 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symremove, tmp21795)
}
__typedArg0 := symremove
__typedArg1 := tmp21795
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21797 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21796)


_ = tmp21797

tmp21798 := MakeNative(func(__e *ControlFlow) {
Y1041 := __e.Get(1)
_ = Y1041
__e.TailApply(PrimFunc(symreverse), Y1041)
return
}, 1)

tmp21799 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symreverse, tmp21798)
}
__typedArg0 := symreverse
__typedArg1 := tmp21798
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21800 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21799)


_ = tmp21800

tmp21801 := MakeNative(func(__e *ControlFlow) {
Y1039 := __e.Get(1)
_ = Y1039
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1040 := __e.Get(1)
_ = Y1040
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(Y1039, Y1040)
}
__typedArg0 := Y1039
__typedArg1 := Y1040
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return
}, 1))
return
}, 1)

tmp21802 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symset, tmp21801)
}
__typedArg0 := symset
__typedArg1 := tmp21801
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21803 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21802)


_ = tmp21803

tmp21804 := MakeNative(func(__e *ControlFlow) {
Y1038 := __e.Get(1)
_ = Y1038
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(Y1038)
}
__typedArg0 := Y1038
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}, 1)

tmp21805 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsimple_1error, tmp21804)
}
__typedArg0 := symsimple_1error
__typedArg1 := tmp21804
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21806 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21805)


_ = tmp21806

tmp21807 := MakeNative(func(__e *ControlFlow) {
Y1037 := __e.Get(1)
_ = Y1037
__e.TailApply(PrimFunc(symsnd), Y1037)
return
}, 1)

tmp21808 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsnd, tmp21807)
}
__typedArg0 := symsnd
__typedArg1 := tmp21807
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21809 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21808)


_ = tmp21809

tmp21810 := MakeNative(func(__e *ControlFlow) {
Y1035 := __e.Get(1)
_ = Y1035
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1036 := __e.Get(1)
_ = Y1036
__e.TailApply(PrimFunc(symspecialise), Y1035, Y1036)
return
}, 1))
return
}, 1)

tmp21811 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symspecialise, tmp21810)
}
__typedArg0 := symspecialise
__typedArg1 := tmp21810
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21812 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21811)


_ = tmp21812

tmp21813 := MakeNative(func(__e *ControlFlow) {
Y1034 := __e.Get(1)
_ = Y1034
__e.TailApply(PrimFunc(symspy), Y1034)
return
}, 1)

tmp21814 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symspy, tmp21813)
}
__typedArg0 := symspy
__typedArg1 := tmp21813
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21815 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21814)


_ = tmp21815

tmp21816 := MakeNative(func(__e *ControlFlow) {
Y1033 := __e.Get(1)
_ = Y1033
__e.TailApply(PrimFunc(symstep), Y1033)
return
}, 1)

tmp21817 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstep, tmp21816)
}
__typedArg0 := symstep
__typedArg1 := tmp21816
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21818 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21817)


_ = tmp21818

tmp21819 := MakeNative(func(__e *ControlFlow) {
Y1032 := __e.Get(1)
_ = Y1032
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(Y1032)
}
__typedArg0 := Y1032
return Call(__e, PrimFunc(symstr), __typedArg0)
})())
return
}, 1)

tmp21820 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstr, tmp21819)
}
__typedArg0 := symstr
__typedArg1 := tmp21819
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21821 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21820)


_ = tmp21821

tmp21822 := MakeNative(func(__e *ControlFlow) {
Y1031 := __e.Get(1)
_ = Y1031
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_1_6n) {
return PrimStringToNumber(Y1031)
}
__typedArg0 := Y1031
return Call(__e, PrimFunc(symstring_1_6n), __typedArg0)
})())
return
}, 1)

tmp21823 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring_1_6n, tmp21822)
}
__typedArg0 := symstring_1_6n
__typedArg1 := tmp21822
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21824 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21823)


_ = tmp21824

tmp21825 := MakeNative(func(__e *ControlFlow) {
Y1030 := __e.Get(1)
_ = Y1030
__e.TailApply(PrimFunc(symstring_1_6symbol), Y1030)
return
}, 1)

tmp21826 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring_1_6symbol, tmp21825)
}
__typedArg0 := symstring_1_6symbol
__typedArg1 := tmp21825
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21827 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21826)


_ = tmp21827

tmp21828 := MakeNative(func(__e *ControlFlow) {
Y1029 := __e.Get(1)
_ = Y1029
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_2) {
return PrimIsString(Y1029)
}
__typedArg0 := Y1029
return Call(__e, PrimFunc(symstring_2), __typedArg0)
})())
return
}, 1)

tmp21829 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring_2, tmp21828)
}
__typedArg0 := symstring_2
__typedArg1 := tmp21828
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21830 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21829)


_ = tmp21830

tmp21831 := MakeNative(func(__e *ControlFlow) {
Y1026 := __e.Get(1)
_ = Y1026
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1027 := __e.Get(1)
_ = Y1027
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1028 := __e.Get(1)
_ = Y1028
__e.TailApply(PrimFunc(symsubst), Y1026, Y1027, Y1028)
return
}, 1))
return
}, 1))
return
}, 1)

tmp21832 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsubst, tmp21831)
}
__typedArg0 := symsubst
__typedArg1 := tmp21831
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21833 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21832)


_ = tmp21833

tmp21834 := MakeNative(func(__e *ControlFlow) {
Y1025 := __e.Get(1)
_ = Y1025
__e.TailApply(PrimFunc(symsum), Y1025)
return
}, 1)

tmp21835 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsum, tmp21834)
}
__typedArg0 := symsum
__typedArg1 := tmp21834
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21836 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21835)


_ = tmp21836

tmp21837 := MakeNative(func(__e *ControlFlow) {
Y1024 := __e.Get(1)
_ = Y1024
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsymbol_2) {
return PrimIsSymbol(Y1024)
}
__typedArg0 := Y1024
return Call(__e, PrimFunc(symsymbol_2), __typedArg0)
})())
return
}, 1)

tmp21838 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol_2, tmp21837)
}
__typedArg0 := symsymbol_2
__typedArg1 := tmp21837
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21839 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21838)


_ = tmp21839

tmp21840 := MakeNative(func(__e *ControlFlow) {
Y1023 := __e.Get(1)
_ = Y1023
__e.TailApply(PrimFunc(symsystemf), Y1023)
return
}, 1)

tmp21841 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsystemf, tmp21840)
}
__typedArg0 := symsystemf
__typedArg1 := tmp21840
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21842 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21841)


_ = tmp21842

tmp21843 := MakeNative(func(__e *ControlFlow) {
Y1022 := __e.Get(1)
_ = Y1022
__e.TailApply(PrimFunc(symtail), Y1022)
return
}, 1)

tmp21844 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtail, tmp21843)
}
__typedArg0 := symtail
__typedArg1 := tmp21843
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21845 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21844)


_ = tmp21845

tmp21846 := MakeNative(func(__e *ControlFlow) {
Y1021 := __e.Get(1)
_ = Y1021
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Y1021)
}
__typedArg0 := Y1021
return Call(__e, PrimFunc(symtl), __typedArg0)
})())
return
}, 1)

tmp21847 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtl, tmp21846)
}
__typedArg0 := symtl
__typedArg1 := tmp21846
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21848 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21847)


_ = tmp21848

tmp21849 := MakeNative(func(__e *ControlFlow) {
Y1020 := __e.Get(1)
_ = Y1020
__e.TailApply(PrimFunc(symtc), Y1020)
return
}, 1)

tmp21850 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtc, tmp21849)
}
__typedArg0 := symtc
__typedArg1 := tmp21849
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21851 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21850)


_ = tmp21851

tmp21852 := MakeNative(func(__e *ControlFlow) {
Y1019 := __e.Get(1)
_ = Y1019
__e.TailApply(PrimFunc(symthaw), Y1019)
return
}, 1)

tmp21853 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symthaw, tmp21852)
}
__typedArg0 := symthaw
__typedArg1 := tmp21852
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21854 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21853)


_ = tmp21854

tmp21855 := MakeNative(func(__e *ControlFlow) {
Y1018 := __e.Get(1)
_ = Y1018
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(Y1018)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := Y1018
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())
return
}, 1)

tmp21856 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtlstr, tmp21855)
}
__typedArg0 := symtlstr
__typedArg1 := tmp21855
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21857 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21856)


_ = tmp21857

tmp21858 := MakeNative(func(__e *ControlFlow) {
Y1017 := __e.Get(1)
_ = Y1017
__e.TailApply(PrimFunc(symtrack), Y1017)
return
}, 1)

tmp21859 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtrack, tmp21858)
}
__typedArg0 := symtrack
__typedArg1 := tmp21858
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21860 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21859)


_ = tmp21860

tmp21861 := MakeNative(func(__e *ControlFlow) {
Y1015 := __e.Get(1)
_ = Y1015
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1016 := __e.Get(1)
_ = Y1016
tmp21862 := MakeNative(func(__e *ControlFlow) {
__e.Return(Y1015)
return
}, 0)

__e.TailApply(try_1catch, tmp21862, Y1016)
return


}, 1))
return
}, 1)

tmp21863 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtrap_1error, tmp21861)
}
__typedArg0 := symtrap_1error
__typedArg1 := tmp21861
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21864 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21863)


_ = tmp21864

tmp21865 := MakeNative(func(__e *ControlFlow) {
Y1014 := __e.Get(1)
_ = Y1014
__e.TailApply(PrimFunc(symtuple_2), Y1014)
return
}, 1)

tmp21866 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtuple_2, tmp21865)
}
__typedArg0 := symtuple_2
__typedArg1 := tmp21865
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21867 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21866)


_ = tmp21867

tmp21868 := MakeNative(func(__e *ControlFlow) {
Y1012 := __e.Get(1)
_ = Y1012
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1013 := __e.Get(1)
_ = Y1013
__e.Return(Y1013)
return
}, 1))
return
}, 1)

tmp21869 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtype, tmp21868)
}
__typedArg0 := symtype
__typedArg1 := tmp21868
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21870 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21869)


_ = tmp21870

tmp21871 := MakeNative(func(__e *ControlFlow) {
Y1007 := __e.Get(1)
_ = Y1007
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1008 := __e.Get(1)
_ = Y1008
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1009 := __e.Get(1)
_ = Y1009
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1010 := __e.Get(1)
_ = Y1010
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1011 := __e.Get(1)
_ = Y1011
__e.TailApply(PrimFunc(symreturn), Y1007, Y1008, Y1009, Y1010, Y1011)
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21872 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symreturn, tmp21871)
}
__typedArg0 := symreturn
__typedArg1 := tmp21871
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21873 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21872)


_ = tmp21873

tmp21874 := MakeNative(func(__e *ControlFlow) {
Y1006 := __e.Get(1)
_ = Y1006
__e.TailApply(PrimFunc(symunabsolute), Y1006)
return
}, 1)

tmp21875 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunabsolute, tmp21874)
}
__typedArg0 := symunabsolute
__typedArg1 := tmp21874
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21876 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21875)


_ = tmp21876

tmp21877 := MakeNative(func(__e *ControlFlow) {
Y1005 := __e.Get(1)
_ = Y1005
__e.TailApply(PrimFunc(symundefmacro), Y1005)
return
}, 1)

tmp21878 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symundefmacro, tmp21877)
}
__typedArg0 := symundefmacro
__typedArg1 := tmp21877
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21879 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21878)


_ = tmp21879

tmp21880 := MakeNative(func(__e *ControlFlow) {
Y1002 := __e.Get(1)
_ = Y1002
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1003 := __e.Get(1)
_ = Y1003
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1004 := __e.Get(1)
_ = Y1004
__e.TailApply(PrimFunc(symunput), Y1002, Y1003, Y1004)
return
}, 1))
return
}, 1))
return
}, 1)

tmp21881 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunput, tmp21880)
}
__typedArg0 := symunput
__typedArg1 := tmp21880
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21882 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21881)


_ = tmp21882

tmp21883 := MakeNative(func(__e *ControlFlow) {
Y1001 := __e.Get(1)
_ = Y1001
__e.TailApply(PrimFunc(symunprofile), Y1001)
return
}, 1)

tmp21884 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunprofile, tmp21883)
}
__typedArg0 := symunprofile
__typedArg1 := tmp21883
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21885 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21884)


_ = tmp21885

tmp21886 := MakeNative(func(__e *ControlFlow) {
Y999 := __e.Get(1)
_ = Y999
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1000 := __e.Get(1)
_ = Y1000
__e.TailApply(PrimFunc(symunion), Y999, Y1000)
return
}, 1))
return
}, 1)

tmp21887 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunion, tmp21886)
}
__typedArg0 := symunion
__typedArg1 := tmp21886
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21888 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21887)


_ = tmp21888

tmp21889 := MakeNative(func(__e *ControlFlow) {
Y998 := __e.Get(1)
_ = Y998
__e.TailApply(PrimFunc(symuntrack), Y998)
return
}, 1)

tmp21890 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symuntrack, tmp21889)
}
__typedArg0 := symuntrack
__typedArg1 := tmp21889
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21891 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21890)


_ = tmp21891

tmp21892 := MakeNative(func(__e *ControlFlow) {
Y997 := __e.Get(1)
_ = Y997
__e.TailApply(PrimFunc(symundefmacro), Y997)
return
}, 1)

tmp21893 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symundefmacro, tmp21892)
}
__typedArg0 := symundefmacro
__typedArg1 := tmp21892
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21894 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21893)


_ = tmp21894

tmp21895 := MakeNative(func(__e *ControlFlow) {
Y995 := __e.Get(1)
_ = Y995
__e.Return(MakeNative(func(__e *ControlFlow) {
Y996 := __e.Get(1)
_ = Y996
__e.TailApply(PrimFunc(symupdate_1lambda_1table), Y995, Y996)
return
}, 1))
return
}, 1)

tmp21896 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symupdate_1lambda_1table, tmp21895)
}
__typedArg0 := symupdate_1lambda_1table
__typedArg1 := tmp21895
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21897 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21896)


_ = tmp21897

tmp21898 := MakeNative(func(__e *ControlFlow) {
Y994 := __e.Get(1)
_ = Y994
__e.TailApply(PrimFunc(symvector), Y994)
return
}, 1)

tmp21899 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp21898)
}
__typedArg0 := symvector
__typedArg1 := tmp21898
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21900 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21899)


_ = tmp21900

tmp21901 := MakeNative(func(__e *ControlFlow) {
Y993 := __e.Get(1)
_ = Y993
__e.TailApply(PrimFunc(symvector_2), Y993)
return
}, 1)

tmp21902 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector_2, tmp21901)
}
__typedArg0 := symvector_2
__typedArg1 := tmp21901
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21903 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21902)


_ = tmp21903

tmp21904 := MakeNative(func(__e *ControlFlow) {
Y990 := __e.Get(1)
_ = Y990
__e.Return(MakeNative(func(__e *ControlFlow) {
Y991 := __e.Get(1)
_ = Y991
__e.Return(MakeNative(func(__e *ControlFlow) {
Y992 := __e.Get(1)
_ = Y992
__e.TailApply(PrimFunc(symvector_1_6), Y990, Y991, Y992)
return
}, 1))
return
}, 1))
return
}, 1)

tmp21905 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector_1_6, tmp21904)
}
__typedArg0 := symvector_1_6
__typedArg1 := tmp21904
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21906 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21905)


_ = tmp21906

tmp21907 := MakeNative(func(__e *ControlFlow) {
Y989 := __e.Get(1)
_ = Y989
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(Y989)
}
__typedArg0 := Y989
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 1)

tmp21908 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvalue, tmp21907)
}
__typedArg0 := symvalue
__typedArg1 := tmp21907
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21909 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21908)


_ = tmp21909

tmp21910 := MakeNative(func(__e *ControlFlow) {
Y988 := __e.Get(1)
_ = Y988
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(Y988)
}
__typedArg0 := Y988
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})())
return
}, 1)

tmp21911 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvariable_2, tmp21910)
}
__typedArg0 := symvariable_2
__typedArg1 := tmp21910
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21912 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21911)


_ = tmp21912

tmp21913 := MakeNative(func(__e *ControlFlow) {
Y983 := __e.Get(1)
_ = Y983
__e.Return(MakeNative(func(__e *ControlFlow) {
Y984 := __e.Get(1)
_ = Y984
__e.Return(MakeNative(func(__e *ControlFlow) {
Y985 := __e.Get(1)
_ = Y985
__e.Return(MakeNative(func(__e *ControlFlow) {
Y986 := __e.Get(1)
_ = Y986
__e.Return(MakeNative(func(__e *ControlFlow) {
Y987 := __e.Get(1)
_ = Y987
__e.TailApply(PrimFunc(symvar_2), Y983, Y984, Y985, Y986, Y987)
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21914 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvar_2, tmp21913)
}
__typedArg0 := symvar_2
__typedArg1 := tmp21913
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21915 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21914)


_ = tmp21915

tmp21916 := MakeNative(func(__e *ControlFlow) {
Y978 := __e.Get(1)
_ = Y978
__e.Return(MakeNative(func(__e *ControlFlow) {
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
__e.TailApply(PrimFunc(symwhen), Y978, Y979, Y980, Y981, Y982)
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21917 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhen, tmp21916)
}
__typedArg0 := symwhen
__typedArg1 := tmp21916
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21918 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21917)


_ = tmp21918

tmp21919 := MakeNative(func(__e *ControlFlow) {
Y976 := __e.Get(1)
_ = Y976
__e.Return(MakeNative(func(__e *ControlFlow) {
Y977 := __e.Get(1)
_ = Y977
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symwrite_1byte) {
return PrimWriteByte(Y976, Y977)
}
__typedArg0 := Y976
__typedArg1 := Y977
return Call(__e, PrimFunc(symwrite_1byte), __typedArg0, __typedArg1)
})())
return
}, 1))
return
}, 1)

tmp21920 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwrite_1byte, tmp21919)
}
__typedArg0 := symwrite_1byte
__typedArg1 := tmp21919
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21921 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21920)


_ = tmp21921

tmp21922 := MakeNative(func(__e *ControlFlow) {
Y974 := __e.Get(1)
_ = Y974
__e.Return(MakeNative(func(__e *ControlFlow) {
Y975 := __e.Get(1)
_ = Y975
__e.TailApply(PrimFunc(symwrite_1to_1file), Y974, Y975)
return
}, 1))
return
}, 1)

tmp21923 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwrite_1to_1file, tmp21922)
}
__typedArg0 := symwrite_1to_1file
__typedArg1 := tmp21922
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21924 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21923)


_ = tmp21924

tmp21925 := MakeNative(func(__e *ControlFlow) {
Y973 := __e.Get(1)
_ = Y973
__e.TailApply(PrimFunc(symy_1or_1n_2), Y973)
return
}, 1)

tmp21926 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symy_1or_1n_2, tmp21925)
}
__typedArg0 := symy_1or_1n_2
__typedArg1 := tmp21925
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21927 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21926)


_ = tmp21927

tmp21928 := MakeNative(func(__e *ControlFlow) {
Y971 := __e.Get(1)
_ = Y971
__e.Return(MakeNative(func(__e *ControlFlow) {
Y972 := __e.Get(1)
_ = Y972
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(Y971)
__typedN1, __typedOK1 := TypedFloat64(Y972)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := Y971
__typedArg1 := Y972
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())
return
}, 1))
return
}, 1)

tmp21929 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_7, tmp21928)
}
__typedArg0 := sym_7
__typedArg1 := tmp21928
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21930 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21929)


_ = tmp21930

tmp21931 := MakeNative(func(__e *ControlFlow) {
Y969 := __e.Get(1)
_ = Y969
__e.Return(MakeNative(func(__e *ControlFlow) {
Y970 := __e.Get(1)
_ = Y970
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_d) {
__typedN0, __typedOK0 := TypedFloat64(Y969)
__typedN1, __typedOK1 := TypedFloat64(Y970)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_d) {
return TypedMaterializeNumber((__typedN0 * __typedN1))
}}
__typedArg0 := Y969
__typedArg1 := Y970
return Call(__e, PrimFunc(sym_d), __typedArg0, __typedArg1)
})())
return
}, 1))
return
}, 1)

tmp21932 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_d, tmp21931)
}
__typedArg0 := sym_d
__typedArg1 := tmp21931
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21933 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21932)


_ = tmp21933

tmp21934 := MakeNative(func(__e *ControlFlow) {
Y967 := __e.Get(1)
_ = Y967
__e.Return(MakeNative(func(__e *ControlFlow) {
Y968 := __e.Get(1)
_ = Y968
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_c) {
__typedN0, __typedOK0 := TypedFloat64(Y967)
__typedN1, __typedOK1 := TypedFloat64(Y968)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_c) {
return TypedMaterializeNumber(TypedDivideValue(__typedN0, __typedN1))
}}
__typedArg0 := Y967
__typedArg1 := Y968
return Call(__e, PrimFunc(sym_c), __typedArg0, __typedArg1)
})())
return
}, 1))
return
}, 1)

tmp21935 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_c, tmp21934)
}
__typedArg0 := sym_c
__typedArg1 := tmp21934
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21936 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21935)


_ = tmp21936

tmp21937 := MakeNative(func(__e *ControlFlow) {
Y965 := __e.Get(1)
_ = Y965
__e.Return(MakeNative(func(__e *ControlFlow) {
Y966 := __e.Get(1)
_ = Y966
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(Y965)
__typedN1, __typedOK1 := TypedFloat64(Y966)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := Y965
__typedArg1 := Y966
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})())
return
}, 1))
return
}, 1)

tmp21938 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1, tmp21937)
}
__typedArg0 := sym_1
__typedArg1 := tmp21937
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21939 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21938)


_ = tmp21939

tmp21940 := MakeNative(func(__e *ControlFlow) {
Y963 := __e.Get(1)
_ = Y963
__e.Return(MakeNative(func(__e *ControlFlow) {
Y964 := __e.Get(1)
_ = Y964
__e.TailApply(PrimFunc(sym_a_a), Y963, Y964)
return
}, 1))
return
}, 1)

tmp21941 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_a_a, tmp21940)
}
__typedArg0 := sym_a_a
__typedArg1 := tmp21940
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21942 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21941)


_ = tmp21942

tmp21943 := MakeNative(func(__e *ControlFlow) {
Y962 := __e.Get(1)
_ = Y962
__e.TailApply(PrimFunc(sym_5e_6), Y962)
return
}, 1)

tmp21944 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5e_6, tmp21943)
}
__typedArg0 := sym_5e_6
__typedArg1 := tmp21943
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21945 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21944)


_ = tmp21945

tmp21946 := MakeNative(func(__e *ControlFlow) {
Y961 := __e.Get(1)
_ = Y961
__e.TailApply(PrimFunc(sym_5end_6), Y961)
return
}, 1)

tmp21947 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5end_6, tmp21946)
}
__typedArg0 := sym_5end_6
__typedArg1 := tmp21946
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21948 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21947)


_ = tmp21948

tmp21949 := MakeNative(func(__e *ControlFlow) {
Y960 := __e.Get(1)
_ = Y960
__e.TailApply(PrimFunc(sym_5_b_6), Y960)
return
}, 1)

tmp21950 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_b_6, tmp21949)
}
__typedArg0 := sym_5_b_6
__typedArg1 := tmp21949
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21951 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21950)


_ = tmp21951

tmp21952 := MakeNative(func(__e *ControlFlow) {
Y958 := __e.Get(1)
_ = Y958
__e.Return(MakeNative(func(__e *ControlFlow) {
Y959 := __e.Get(1)
_ = Y959
__e.TailApply(PrimFunc(sym_8p), Y958, Y959)
return
}, 1))
return
}, 1)

tmp21953 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8p, tmp21952)
}
__typedArg0 := sym_8p
__typedArg1 := tmp21952
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21954 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21953)


_ = tmp21954

tmp21955 := MakeNative(func(__e *ControlFlow) {
Y956 := __e.Get(1)
_ = Y956
__e.Return(MakeNative(func(__e *ControlFlow) {
Y957 := __e.Get(1)
_ = Y957
__e.TailApply(PrimFunc(sym_8v), Y956, Y957)
return
}, 1))
return
}, 1)

tmp21956 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8v, tmp21955)
}
__typedArg0 := sym_8v
__typedArg1 := tmp21955
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21957 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21956)


_ = tmp21957

tmp21958 := MakeNative(func(__e *ControlFlow) {
Y954 := __e.Get(1)
_ = Y954
__e.Return(MakeNative(func(__e *ControlFlow) {
Y955 := __e.Get(1)
_ = Y955
__e.TailApply(PrimFunc(sym_8s), Y954, Y955)
return
}, 1))
return
}, 1)

tmp21959 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8s, tmp21958)
}
__typedArg0 := sym_8s
__typedArg1 := tmp21958
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21959)
return


}, 0)

tmp21960 := Call(__e, ns2_1set, symshen_4initialise_1lambda_1forms, tmp21411)


_ = tmp21960

tmp21961 := MakeNative(func(__e *ControlFlow) {
tmp21962 := Call(__e, PrimFunc(symshen_4initialise_1environment))


_ = tmp21962

tmp21963 := Call(__e, PrimFunc(symshen_4initialise_1lambda_1forms))


_ = tmp21963

__e.TailApply(PrimFunc(symshen_4initialise_1signedfuncs))
return


}, 0)

__e.TailApply(ns2_1set, symshen_4initialise, tmp21961)
return




}, 0)

