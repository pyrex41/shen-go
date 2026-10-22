package main

import . "github.com/pyrex41/shen-go/kl"

var InitMain = MakeNative(func(__e *ControlFlow) {
tmp18993 := MakeNative(func(__e *ControlFlow) {
tmp18994 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dhistory_d, Nil)
}
__typedArg0 := symshen_4_dhistory_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp18994

tmp18995 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dtc_d, False)
}
__typedArg0 := symshen_4_dtc_d
__typedArg1 := False
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp18995

tmp18996 := Call(__e, PrimFunc(symshen_4dict), MakeInteger(20000))


tmp18997 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(sym_dproperty_1vector_d, tmp18996)
}
__typedArg0 := sym_dproperty_1vector_d
__typedArg1 := tmp18996
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp18997

tmp18998 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
__e.TailApply(PrimFunc(symshen_4macros), X)
return
}, 1)

tmp18999 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4macros, tmp18998)
}
__typedArg0 := symshen_4macros
__typedArg1 := tmp18998
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19000 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18999, Nil)
}
__typedArg0 := tmp18999
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19001 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(sym_dmacros_d, tmp19000)
}
__typedArg0 := sym_dmacros_d
__typedArg1 := tmp19000
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19001

tmp19002 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsource_1form_1handlers_d, Nil)
}
__typedArg0 := symshen_4_dsource_1form_1handlers_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19002

tmp19003 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dgensym_d, MakeInteger(0))
}
__typedArg0 := symshen_4_dgensym_d
__typedArg1 := MakeInteger(0)
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19003

tmp19004 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dtracking_d, Nil)
}
__typedArg0 := symshen_4_dtracking_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19004

tmp19005 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dprofiled_d, Nil)
}
__typedArg0 := symshen_4_dprofiled_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19005

tmp19006 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtype, Nil)
}
__typedArg0 := symtype
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19007 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4input_1h_7, tmp19006)
}
__typedArg0 := symshen_4input_1h_7
__typedArg1 := tmp19006
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19008 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symopen, tmp19007)
}
__typedArg0 := symopen
__typedArg1 := tmp19007
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19009 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symset, tmp19008)
}
__typedArg0 := symset
__typedArg1 := tmp19008
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19010 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhere, tmp19009)
}
__typedArg0 := symwhere
__typedArg1 := tmp19009
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19011 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp19010)
}
__typedArg0 := symlet
__typedArg1 := tmp19010
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19012 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp19011)
}
__typedArg0 := symlambda
__typedArg1 := tmp19011
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19013 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp19012)
}
__typedArg0 := symcons
__typedArg1 := tmp19012
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19014 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8v, tmp19013)
}
__typedArg0 := sym_8v
__typedArg1 := tmp19013
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19015 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8s, tmp19014)
}
__typedArg0 := sym_8s
__typedArg1 := tmp19014
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19016 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8p, tmp19015)
}
__typedArg0 := sym_8p
__typedArg1 := tmp19015
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19017 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dspecial_d, tmp19016)
}
__typedArg0 := symshen_4_dspecial_d
__typedArg1 := tmp19016
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19017

tmp19018 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dextraspecial_d, Nil)
}
__typedArg0 := symshen_4_dextraspecial_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19018

tmp19019 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dspy_d, False)
}
__typedArg0 := symshen_4_dspy_d
__typedArg1 := False
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19019

tmp19020 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_ddatatypes_d, Nil)
}
__typedArg0 := symshen_4_ddatatypes_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19020

tmp19021 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dalldatatypes_d, Nil)
}
__typedArg0 := symshen_4_dalldatatypes_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19021

tmp19022 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dshen_1type_1theory_1enabled_2_d, True)
}
__typedArg0 := symshen_4_dshen_1type_1theory_1enabled_2_d
__typedArg1 := True
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19022

tmp19023 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dpackage_d, symnull)
}
__typedArg0 := symshen_4_dpackage_d
__typedArg1 := symnull
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19023

tmp19024 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsynonyms_d, Nil)
}
__typedArg0 := symshen_4_dsynonyms_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19024

tmp19025 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_doccurs_d, True)
}
__typedArg0 := symshen_4_doccurs_d
__typedArg1 := True
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19025

tmp19026 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dfactorise_2_d, False)
}
__typedArg0 := symshen_4_dfactorise_2_d
__typedArg1 := False
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19026

tmp19027 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dmaxinferences_d, MakeInteger(1000000))
}
__typedArg0 := symshen_4_dmaxinferences_d
__typedArg1 := MakeInteger(1000000)
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19027

tmp19028 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(sym_dmaximum_1print_1sequence_1size_d, MakeInteger(20))
}
__typedArg0 := sym_dmaximum_1print_1sequence_1size_d
__typedArg1 := MakeInteger(20)
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19028

tmp19029 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dcall_d, MakeInteger(0))
}
__typedArg0 := symshen_4_dcall_d
__typedArg1 := MakeInteger(0)
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19029

tmp19030 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dinfs_d, MakeInteger(0))
}
__typedArg0 := symshen_4_dinfs_d
__typedArg1 := MakeInteger(0)
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19030

tmp19031 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(sym_dhush_d, False)
}
__typedArg0 := sym_dhush_d
__typedArg1 := False
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19031

tmp19032 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_doptimise_d, False)
}
__typedArg0 := symshen_4_doptimise_d
__typedArg1 := False
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19032

tmp19033 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(sym_dversion_d, MakeString("43.1"))
}
__typedArg0 := sym_dversion_d
__typedArg1 := MakeString("43.1")
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19033

tmp19034 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dnames_d, Nil)
}
__typedArg0 := symshen_4_dnames_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19034

tmp19035 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dstep_d, False)
}
__typedArg0 := symshen_4_dstep_d
__typedArg1 := False
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19035

tmp19036 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dit_d, MakeString(""))
}
__typedArg0 := symshen_4_dit_d
__typedArg1 := MakeString("")
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19036

tmp19037 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dresidue_d, Nil)
}
__typedArg0 := symshen_4_dresidue_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19037

tmp19038 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(sym_dabsolute_d, Nil)
}
__typedArg0 := sym_dabsolute_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19038

tmp19039 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dprolog_1memory_d, MakeInteger(1000))
}
__typedArg0 := symshen_4_dprolog_1memory_d
__typedArg1 := MakeInteger(1000)
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19039

tmp19040 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dloading_2_d, False)
}
__typedArg0 := symshen_4_dloading_2_d
__typedArg1 := False
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19040

tmp19041 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_duserdefs_d, Nil)
}
__typedArg0 := symshen_4_duserdefs_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19041

tmp19042 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
__e.Return(X)
return
}, 1)

tmp19043 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_ddemodulation_1function_d, tmp19042)
}
__typedArg0 := symshen_4_ddemodulation_1function_d
__typedArg1 := tmp19042
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19043

tmp19044 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dcustom_1pattern_1compiler_d, False)
}
__typedArg0 := symshen_4_dcustom_1pattern_1compiler_d
__typedArg1 := False
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19044

tmp19045 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dcustom_1pattern_1reducer_d, False)
}
__typedArg0 := symshen_4_dcustom_1pattern_1reducer_d
__typedArg1 := False
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19045

tmp19048 := Call(__e, PrimFunc(symbound_2), sym_dhome_1directory_d)


tmp19049 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp19048)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp19048
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres19046 Obj

if True == tmp19049 {
tmp19047 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(sym_dhome_1directory_d, MakeString(""))
}
__typedArg0 := sym_dhome_1directory_d
__typedArg1 := MakeString("")
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

ifres19046 = tmp19047


} else {
ifres19046 = symshen_4skip


}

_ = ifres19046

tmp19053 := Call(__e, PrimFunc(symbound_2), sym_dsterror_d)


tmp19054 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp19053)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp19053
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres19050 Obj

if True == tmp19054 {
tmp19051 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dstoutput_d)
}
__typedArg0 := sym_dstoutput_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp19052 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(sym_dsterror_d, tmp19051)
}
__typedArg0 := sym_dsterror_d
__typedArg1 := tmp19051
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

ifres19050 = tmp19052


} else {
ifres19050 = symshen_4skip


}

_ = ifres19050

tmp19055 := Call(__e, PrimFunc(symprolog_1memory), MakeInteger(10000))


_ = tmp19055

tmp19056 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dloading_2_d, False)
}
__typedArg0 := symshen_4_dloading_2_d
__typedArg1 := False
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19056

tmp19057 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), Nil)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19058 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8s, tmp19057)
}
__typedArg0 := sym_8s
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
return PrimCons(sym_8v, tmp19059)
}
__typedArg0 := sym_8v
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
return PrimCons(sym_8p, tmp19061)
}
__typedArg0 := sym_8p
__typedArg1 := tmp19061
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19063 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19062)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19062
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19064 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_b_6, tmp19063)
}
__typedArg0 := sym_5_b_6
__typedArg1 := tmp19063
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19065 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19064)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19064
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19066 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5end_6, tmp19065)
}
__typedArg0 := sym_5end_6
__typedArg1 := tmp19065
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19067 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19066)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19066
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19068 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5e_6, tmp19067)
}
__typedArg0 := sym_5e_6
__typedArg1 := tmp19067
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19069 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19068)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19068
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19070 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_a_a, tmp19069)
}
__typedArg0 := sym_a_a
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
return PrimCons(sym_1, tmp19071)
}
__typedArg0 := sym_1
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
return PrimCons(sym_c, tmp19073)
}
__typedArg0 := sym_c
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
return PrimCons(sym_d, tmp19075)
}
__typedArg0 := sym_d
__typedArg1 := tmp19075
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19077 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19076)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19076
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19078 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_7, tmp19077)
}
__typedArg0 := sym_7
__typedArg1 := tmp19077
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19079 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19078)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19078
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19080 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symy_1or_1n_2, tmp19079)
}
__typedArg0 := symy_1or_1n_2
__typedArg1 := tmp19079
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19081 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19080)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19080
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19082 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwrite_1to_1file, tmp19081)
}
__typedArg0 := symwrite_1to_1file
__typedArg1 := tmp19081
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19083 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19082)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19082
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19084 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwrite_1string, tmp19083)
}
__typedArg0 := symwrite_1string
__typedArg1 := tmp19083
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19085 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19084)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19084
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19086 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwrite_1byte, tmp19085)
}
__typedArg0 := symwrite_1byte
__typedArg1 := tmp19085
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19087 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(5), tmp19086)
}
__typedArg0 := MakeInteger(5)
__typedArg1 := tmp19086
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19088 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhen, tmp19087)
}
__typedArg0 := symwhen
__typedArg1 := tmp19087
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19089 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19088)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19088
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19090 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symversion, tmp19089)
}
__typedArg0 := symversion
__typedArg1 := tmp19089
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19091 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(5), tmp19090)
}
__typedArg0 := MakeInteger(5)
__typedArg1 := tmp19090
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19092 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvar_2, tmp19091)
}
__typedArg0 := symvar_2
__typedArg1 := tmp19091
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19093 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19092)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19092
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19094 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvariable_2, tmp19093)
}
__typedArg0 := symvariable_2
__typedArg1 := tmp19093
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19095 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19094)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19094
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19096 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvalue, tmp19095)
}
__typedArg0 := symvalue
__typedArg1 := tmp19095
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19097 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(3), tmp19096)
}
__typedArg0 := MakeInteger(3)
__typedArg1 := tmp19096
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19098 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector_1_6, tmp19097)
}
__typedArg0 := symvector_1_6
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
return PrimCons(symvector_2, tmp19099)
}
__typedArg0 := symvector_2
__typedArg1 := tmp19099
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19101 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19100)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19100
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19102 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp19101)
}
__typedArg0 := symvector
__typedArg1 := tmp19101
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19103 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19102)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19102
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19104 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symuserdefs, tmp19103)
}
__typedArg0 := symuserdefs
__typedArg1 := tmp19103
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19105 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19104)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19104
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19106 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symupdate_1lambda_1table, tmp19105)
}
__typedArg0 := symupdate_1lambda_1table
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
return PrimCons(symuntrack, tmp19109)
}
__typedArg0 := symuntrack
__typedArg1 := tmp19109
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19111 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19110)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19110
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19112 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunion, tmp19111)
}
__typedArg0 := symunion
__typedArg1 := tmp19111
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19113 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19112)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19112
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19114 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunprofile, tmp19113)
}
__typedArg0 := symunprofile
__typedArg1 := tmp19113
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19115 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(3), tmp19114)
}
__typedArg0 := MakeInteger(3)
__typedArg1 := tmp19114
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19116 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunput, tmp19115)
}
__typedArg0 := symunput
__typedArg1 := tmp19115
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19117 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19116)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19116
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19118 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symundefmacro, tmp19117)
}
__typedArg0 := symundefmacro
__typedArg1 := tmp19117
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19119 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19118)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19118
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19120 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunabsolute, tmp19119)
}
__typedArg0 := symunabsolute
__typedArg1 := tmp19119
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19121 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(5), tmp19120)
}
__typedArg0 := MakeInteger(5)
__typedArg1 := tmp19120
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19122 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symreturn, tmp19121)
}
__typedArg0 := symreturn
__typedArg1 := tmp19121
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19123 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19122)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19122
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19124 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtype, tmp19123)
}
__typedArg0 := symtype
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
return PrimCons(symtuple_2, tmp19125)
}
__typedArg0 := symtuple_2
__typedArg1 := tmp19125
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19127 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19126)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19126
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19128 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtrap_1error, tmp19127)
}
__typedArg0 := symtrap_1error
__typedArg1 := tmp19127
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19129 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19128)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19128
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19130 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtracked, tmp19129)
}
__typedArg0 := symtracked
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
return PrimCons(symtrack, tmp19131)
}
__typedArg0 := symtrack
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
return PrimCons(symtlstr, tmp19133)
}
__typedArg0 := symtlstr
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
return PrimCons(symthaw, tmp19135)
}
__typedArg0 := symthaw
__typedArg1 := tmp19135
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19137 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19136)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19136
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19138 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtc_2, tmp19137)
}
__typedArg0 := symtc_2
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
return PrimCons(symtc, tmp19139)
}
__typedArg0 := symtc
__typedArg1 := tmp19139
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19141 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19140)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19140
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19142 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtl, tmp19141)
}
__typedArg0 := symtl
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
return PrimCons(symtail, tmp19143)
}
__typedArg0 := symtail
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
return PrimCons(symsystemf, tmp19145)
}
__typedArg0 := symsystemf
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
return PrimCons(symsymbol_2, tmp19147)
}
__typedArg0 := symsymbol_2
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
return PrimCons(symsum, tmp19149)
}
__typedArg0 := symsum
__typedArg1 := tmp19149
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19151 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(3), tmp19150)
}
__typedArg0 := MakeInteger(3)
__typedArg1 := tmp19150
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19152 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsubst, tmp19151)
}
__typedArg0 := symsubst
__typedArg1 := tmp19151
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19153 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19152)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19152
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19154 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring_2, tmp19153)
}
__typedArg0 := symstring_2
__typedArg1 := tmp19153
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19155 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19154)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19154
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19156 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring_1_6symbol, tmp19155)
}
__typedArg0 := symstring_1_6symbol
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
return PrimCons(symstring_1_6n, tmp19157)
}
__typedArg0 := symstring_1_6n
__typedArg1 := tmp19157
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19159 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19158)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19158
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19160 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstr, tmp19159)
}
__typedArg0 := symstr
__typedArg1 := tmp19159
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19161 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19160)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19160
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19162 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstoutput, tmp19161)
}
__typedArg0 := symstoutput
__typedArg1 := tmp19161
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19163 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19162)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19162
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19164 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstinput, tmp19163)
}
__typedArg0 := symstinput
__typedArg1 := tmp19163
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19165 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19164)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19164
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19166 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstep_2, tmp19165)
}
__typedArg0 := symstep_2
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
return PrimCons(symstep, tmp19167)
}
__typedArg0 := symstep
__typedArg1 := tmp19167
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19169 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19168)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19168
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19170 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symspy_2, tmp19169)
}
__typedArg0 := symspy_2
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
return PrimCons(symspy, tmp19171)
}
__typedArg0 := symspy
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
return PrimCons(symspecialise, tmp19173)
}
__typedArg0 := symspecialise
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
return PrimCons(symsnd, tmp19175)
}
__typedArg0 := symsnd
__typedArg1 := tmp19175
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19177 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19176)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19176
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19178 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsimple_1error, tmp19177)
}
__typedArg0 := symsimple_1error
__typedArg1 := tmp19177
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19179 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19178)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19178
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19180 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symset, tmp19179)
}
__typedArg0 := symset
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
return PrimCons(symreverse, tmp19181)
}
__typedArg0 := symreverse
__typedArg1 := tmp19181
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19183 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19182)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19182
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19184 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symremove, tmp19183)
}
__typedArg0 := symremove
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
return PrimCons(symshen_4unregister_1source_1form, tmp19185)
}
__typedArg0 := symshen_4unregister_1source_1form
__typedArg1 := tmp19185
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19187 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19186)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19186
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19188 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4register_1source_1form, tmp19187)
}
__typedArg0 := symshen_4register_1source_1form
__typedArg1 := tmp19187
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19189 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19188)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19188
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19190 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symrelease, tmp19189)
}
__typedArg0 := symrelease
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
return PrimCons(symreceive, tmp19191)
}
__typedArg0 := symreceive
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
return PrimCons(symread_1unit_1string, tmp19193)
}
__typedArg0 := symread_1unit_1string
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
return PrimCons(symread_1from_1string_1unprocessed, tmp19195)
}
__typedArg0 := symread_1from_1string_1unprocessed
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
return PrimCons(symread_1from_1string, tmp19197)
}
__typedArg0 := symread_1from_1string
__typedArg1 := tmp19197
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19199 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19198)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19198
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19200 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1byte, tmp19199)
}
__typedArg0 := symread_1byte
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
return PrimCons(symread, tmp19201)
}
__typedArg0 := symread
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
return PrimCons(symread_1file, tmp19203)
}
__typedArg0 := symread_1file
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
return PrimCons(symread_1file_1as_1bytelist, tmp19205)
}
__typedArg0 := symread_1file_1as_1bytelist
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
return PrimCons(symread_1file_1as_1string, tmp19207)
}
__typedArg0 := symread_1file_1as_1string
__typedArg1 := tmp19207
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19209 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(4), tmp19208)
}
__typedArg0 := MakeInteger(4)
__typedArg1 := tmp19208
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19210 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symput, tmp19209)
}
__typedArg0 := symput
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
return PrimCons(symprotect, tmp19211)
}
__typedArg0 := symprotect
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
return PrimCons(sympreclude_1all_1but, tmp19213)
}
__typedArg0 := sympreclude_1all_1but
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
return PrimCons(sympreclude, tmp19215)
}
__typedArg0 := sympreclude
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
return PrimCons(symps, tmp19217)
}
__typedArg0 := symps
__typedArg1 := tmp19217
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19219 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19218)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19218
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19220 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympr, tmp19219)
}
__typedArg0 := sympr
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
return PrimCons(symprofile_1results, tmp19221)
}
__typedArg0 := symprofile_1results
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
return PrimCons(symprolog_1memory, tmp19223)
}
__typedArg0 := symprolog_1memory
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
return PrimCons(symshen_4printF, tmp19225)
}
__typedArg0 := symshen_4printF
__typedArg1 := tmp19225
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19227 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19226)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19226
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19228 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4print_1freshterm, tmp19227)
}
__typedArg0 := symshen_4print_1freshterm
__typedArg1 := tmp19227
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19229 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19228)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19228
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19230 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4print_1prolog_1vector, tmp19229)
}
__typedArg0 := symshen_4print_1prolog_1vector
__typedArg1 := tmp19229
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19231 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19230)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19230
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19232 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symprofile, tmp19231)
}
__typedArg0 := symprofile
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
return PrimCons(symprint, tmp19233)
}
__typedArg0 := symprint
__typedArg1 := tmp19233
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19235 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19234)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19234
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19236 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympreclude_1all_1but, tmp19235)
}
__typedArg0 := sympreclude_1all_1but
__typedArg1 := tmp19235
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19237 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19236)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19236
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19238 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympos, tmp19237)
}
__typedArg0 := sympos
__typedArg1 := tmp19237
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19239 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19238)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19238
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19240 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symporters, tmp19239)
}
__typedArg0 := symporters
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
return PrimCons(symport, tmp19241)
}
__typedArg0 := symport
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
return PrimCons(sympackage_2, tmp19243)
}
__typedArg0 := sympackage_2
__typedArg1 := tmp19243
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19245 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(3), tmp19244)
}
__typedArg0 := MakeInteger(3)
__typedArg1 := tmp19244
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19246 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympackage, tmp19245)
}
__typedArg0 := sympackage
__typedArg1 := tmp19245
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19247 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19246)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19246
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19248 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symos, tmp19247)
}
__typedArg0 := symos
__typedArg1 := tmp19247
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19249 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19248)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19248
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19250 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symor, tmp19249)
}
__typedArg0 := symor
__typedArg1 := tmp19249
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19251 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19250)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19250
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19252 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symoptimise_2, tmp19251)
}
__typedArg0 := symoptimise_2
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
return PrimCons(symoptimise, tmp19253)
}
__typedArg0 := symoptimise
__typedArg1 := tmp19253
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19255 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19254)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19254
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19256 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symopen, tmp19255)
}
__typedArg0 := symopen
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
return PrimCons(symoccurs_1check, tmp19257)
}
__typedArg0 := symoccurs_1check
__typedArg1 := tmp19257
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19259 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19258)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19258
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19260 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symoccurs_2, tmp19259)
}
__typedArg0 := symoccurs_2
__typedArg1 := tmp19259
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19261 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19260)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19260
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19262 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symoccurrences, tmp19261)
}
__typedArg0 := symoccurrences
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
return PrimCons(symoccurs_1check, tmp19263)
}
__typedArg0 := symoccurs_1check
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
return PrimCons(symnumber_2, tmp19265)
}
__typedArg0 := symnumber_2
__typedArg1 := tmp19265
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19267 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19266)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19266
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19268 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symn_1_6string, tmp19267)
}
__typedArg0 := symn_1_6string
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
return PrimCons(symnth, tmp19269)
}
__typedArg0 := symnth
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
return PrimCons(symnot, tmp19271)
}
__typedArg0 := symnot
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
return PrimCons(symnl, tmp19273)
}
__typedArg0 := symnl
__typedArg1 := tmp19273
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19275 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19274)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19274
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19276 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symmaxinferences, tmp19275)
}
__typedArg0 := symmaxinferences
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
return PrimCons(symmapcan, tmp19277)
}
__typedArg0 := symmapcan
__typedArg1 := tmp19277
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19279 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19278)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19278
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19280 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symmap, tmp19279)
}
__typedArg0 := symmap
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
return PrimCons(symmacroexpand, tmp19281)
}
__typedArg0 := symmacroexpand
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
return PrimCons(symvector, tmp19283)
}
__typedArg0 := symvector
__typedArg1 := tmp19283
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19285 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19284)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19284
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19286 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_a, tmp19285)
}
__typedArg0 := sym_5_a
__typedArg1 := tmp19285
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19287 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19286)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19286
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19288 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5, tmp19287)
}
__typedArg0 := sym_5
__typedArg1 := tmp19287
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19289 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19288)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19288
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19290 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symload, tmp19289)
}
__typedArg0 := symload
__typedArg1 := tmp19289
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19291 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19290)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19290
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19292 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19291)
}
__typedArg0 := symlist
__typedArg1 := tmp19291
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19293 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19292)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19292
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19294 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlineread, tmp19293)
}
__typedArg0 := symlineread
__typedArg1 := tmp19293
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19295 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19294)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19294
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19296 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlimit, tmp19295)
}
__typedArg0 := symlimit
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
return PrimCons(symlength, tmp19297)
}
__typedArg0 := symlength
__typedArg1 := tmp19297
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19299 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19298)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19298
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19300 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlanguage, tmp19299)
}
__typedArg0 := symlanguage
__typedArg1 := tmp19299
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19301 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(6), tmp19300)
}
__typedArg0 := MakeInteger(6)
__typedArg1 := tmp19300
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19302 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symis_b, tmp19301)
}
__typedArg0 := symis_b
__typedArg1 := tmp19301
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19303 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(6), tmp19302)
}
__typedArg0 := MakeInteger(6)
__typedArg1 := tmp19302
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19304 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symis, tmp19303)
}
__typedArg0 := symis
__typedArg1 := tmp19303
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19305 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19304)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19304
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19306 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symit, tmp19305)
}
__typedArg0 := symit
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
return PrimCons(syminternal, tmp19307)
}
__typedArg0 := syminternal
__typedArg1 := tmp19307
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19309 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19308)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19308
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19310 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symintersection, tmp19309)
}
__typedArg0 := symintersection
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
return PrimCons(syminclude_1all_1but, tmp19311)
}
__typedArg0 := syminclude_1all_1but
__typedArg1 := tmp19311
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19313 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19312)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19312
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19314 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symimplementation, tmp19313)
}
__typedArg0 := symimplementation
__typedArg1 := tmp19313
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19315 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19314)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19314
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19316 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminput_7, tmp19315)
}
__typedArg0 := syminput_7
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
return PrimCons(syminput, tmp19317)
}
__typedArg0 := syminput
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
return PrimCons(syminferences, tmp19319)
}
__typedArg0 := syminferences
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
return PrimCons(symintern, tmp19321)
}
__typedArg0 := symintern
__typedArg1 := tmp19321
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19323 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19322)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19322
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19324 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminternal, tmp19323)
}
__typedArg0 := syminternal
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
return PrimCons(syminteger_2, tmp19325)
}
__typedArg0 := syminteger_2
__typedArg1 := tmp19325
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19327 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19326)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19326
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19328 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symin_1package, tmp19327)
}
__typedArg0 := symin_1package
__typedArg1 := tmp19327
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19329 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19328)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19328
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19330 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symincluded, tmp19329)
}
__typedArg0 := symincluded
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
return PrimCons(syminclude, tmp19331)
}
__typedArg0 := syminclude
__typedArg1 := tmp19331
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19333 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(3), tmp19332)
}
__typedArg0 := MakeInteger(3)
__typedArg1 := tmp19332
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19334 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp19333)
}
__typedArg0 := symif
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
return PrimCons(symhush, tmp19335)
}
__typedArg0 := symhush
__typedArg1 := tmp19335
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19337 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19336)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19336
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19338 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhush_2, tmp19337)
}
__typedArg0 := symhush_2
__typedArg1 := tmp19337
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19339 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19338)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19338
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19340 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhead, tmp19339)
}
__typedArg0 := symhead
__typedArg1 := tmp19339
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19341 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19340)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19340
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19342 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhdstr, tmp19341)
}
__typedArg0 := symhdstr
__typedArg1 := tmp19341
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19343 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19342)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19342
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19344 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhdv, tmp19343)
}
__typedArg0 := symhdv
__typedArg1 := tmp19343
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19345 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19344)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19344
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19346 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhd, tmp19345)
}
__typedArg0 := symhd
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
return PrimCons(symhash, tmp19347)
}
__typedArg0 := symhash
__typedArg1 := tmp19347
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19349 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19348)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19348
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19350 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_a, tmp19349)
}
__typedArg0 := sym_a
__typedArg1 := tmp19349
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19351 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19350)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19350
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19352 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_6_a, tmp19351)
}
__typedArg0 := sym_6_a
__typedArg1 := tmp19351
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19353 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19352)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19352
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19354 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_6, tmp19353)
}
__typedArg0 := sym_6
__typedArg1 := tmp19353
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19355 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19354)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19354
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19356 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1vector, tmp19355)
}
__typedArg0 := sym_5_1vector
__typedArg1 := tmp19355
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19357 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19356)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19356
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19358 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1address, tmp19357)
}
__typedArg0 := sym_5_1address
__typedArg1 := tmp19357
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19359 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(3), tmp19358)
}
__typedArg0 := MakeInteger(3)
__typedArg1 := tmp19358
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19360 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symaddress_1_6, tmp19359)
}
__typedArg0 := symaddress_1_6
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
return PrimCons(symget_1time, tmp19361)
}
__typedArg0 := symget_1time
__typedArg1 := tmp19361
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19363 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(3), tmp19362)
}
__typedArg0 := MakeInteger(3)
__typedArg1 := tmp19362
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19364 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symget, tmp19363)
}
__typedArg0 := symget
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
return PrimCons(symgensym, tmp19365)
}
__typedArg0 := symgensym
__typedArg1 := tmp19365
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19367 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19366)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19366
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19368 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfunction, tmp19367)
}
__typedArg0 := symfunction
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
return PrimCons(symfn, tmp19369)
}
__typedArg0 := symfn
__typedArg1 := tmp19369
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19371 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19370)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19370
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19372 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfst, tmp19371)
}
__typedArg0 := symfst
__typedArg1 := tmp19371
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19373 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19372)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19372
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19374 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfresh, tmp19373)
}
__typedArg0 := symfresh
__typedArg1 := tmp19373
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19375 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19374)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19374
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19376 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfreeze, tmp19375)
}
__typedArg0 := symfreeze
__typedArg1 := tmp19375
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19377 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(5), tmp19376)
}
__typedArg0 := MakeInteger(5)
__typedArg1 := tmp19376
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19378 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfork, tmp19377)
}
__typedArg0 := symfork
__typedArg1 := tmp19377
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19379 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19378)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19378
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19380 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symforeign, tmp19379)
}
__typedArg0 := symforeign
__typedArg1 := tmp19379
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19381 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(7), tmp19380)
}
__typedArg0 := MakeInteger(7)
__typedArg1 := tmp19380
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19382 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfindall, tmp19381)
}
__typedArg0 := symfindall
__typedArg1 := tmp19381
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19383 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19382)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19382
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19384 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfix, tmp19383)
}
__typedArg0 := symfix
__typedArg1 := tmp19383
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19385 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19384)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19384
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19386 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfail, tmp19385)
}
__typedArg0 := symfail
__typedArg1 := tmp19385
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19387 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19386)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19386
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19388 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfail_1if, tmp19387)
}
__typedArg0 := symfail_1if
__typedArg1 := tmp19387
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19389 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19388)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19388
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19390 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfactorise_2, tmp19389)
}
__typedArg0 := symfactorise_2
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
return PrimCons(symfactorise, tmp19391)
}
__typedArg0 := symfactorise
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
return PrimCons(symexplode, tmp19395)
}
__typedArg0 := symexplode
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
return PrimCons(symeval_1kl, tmp19397)
}
__typedArg0 := symeval_1kl
__typedArg1 := tmp19397
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19399 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19398)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19398
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19400 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symeval, tmp19399)
}
__typedArg0 := symeval
__typedArg1 := tmp19399
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19401 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19400)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19400
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19402 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symerror_1to_1string, tmp19401)
}
__typedArg0 := symerror_1to_1string
__typedArg1 := tmp19401
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19403 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19402)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19402
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19404 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symexternal, tmp19403)
}
__typedArg0 := symexternal
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
return PrimCons(symenable_1type_1theory, tmp19405)
}
__typedArg0 := symenable_1type_1theory
__typedArg1 := tmp19405
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19407 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19406)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19406
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19408 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symempty_2, tmp19407)
}
__typedArg0 := symempty_2
__typedArg1 := tmp19407
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19409 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19408)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19408
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19410 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symelement_2, tmp19409)
}
__typedArg0 := symelement_2
__typedArg1 := tmp19409
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19411 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19410)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19410
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19412 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdo, tmp19411)
}
__typedArg0 := symdo
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
return PrimCons(symdifference, tmp19413)
}
__typedArg0 := symdifference
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
return PrimCons(symdestroy, tmp19415)
}
__typedArg0 := symdestroy
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
return PrimCons(symdeclare, tmp19417)
}
__typedArg0 := symdeclare
__typedArg1 := tmp19417
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19419 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19418)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19418
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19420 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdatatypes, tmp19419)
}
__typedArg0 := symdatatypes
__typedArg1 := tmp19419
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19421 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19420)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19420
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19422 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symclose, tmp19421)
}
__typedArg0 := symclose
__typedArg1 := tmp19421
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19423 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19422)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19422
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19424 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcn, tmp19423)
}
__typedArg0 := symcn
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
return PrimCons(symcons_2, tmp19425)
}
__typedArg0 := symcons_2
__typedArg1 := tmp19425
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19427 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19426)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19426
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19428 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp19427)
}
__typedArg0 := symcons
__typedArg1 := tmp19427
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19429 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19428)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19428
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19430 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symconcat, tmp19429)
}
__typedArg0 := symconcat
__typedArg1 := tmp19429
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19431 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19430)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19430
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19432 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcompile, tmp19431)
}
__typedArg0 := symcompile
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
return PrimCons(symchar_1stream_2, tmp19433)
}
__typedArg0 := symchar_1stream_2
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
return PrimCons(symcd, tmp19435)
}
__typedArg0 := symcd
__typedArg1 := tmp19435
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19437 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(5), tmp19436)
}
__typedArg0 := MakeInteger(5)
__typedArg1 := tmp19436
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19438 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcall, tmp19437)
}
__typedArg0 := symcall
__typedArg1 := tmp19437
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19439 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(6), tmp19438)
}
__typedArg0 := MakeInteger(6)
__typedArg1 := tmp19438
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19440 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symbind, tmp19439)
}
__typedArg0 := symbind
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
return PrimCons(symbound_2, tmp19441)
}
__typedArg0 := symbound_2
__typedArg1 := tmp19441
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19443 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19442)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19442
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19444 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symbootstrap, tmp19443)
}
__typedArg0 := symbootstrap
__typedArg1 := tmp19443
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19445 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19444)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19444
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19446 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean_2, tmp19445)
}
__typedArg0 := symboolean_2
__typedArg1 := tmp19445
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19447 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19446)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19446
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19448 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symatom_2, tmp19447)
}
__typedArg0 := symatom_2
__typedArg1 := tmp19447
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19449 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19448)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19448
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19450 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symassoc, tmp19449)
}
__typedArg0 := symassoc
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
return PrimCons(symarity, tmp19451)
}
__typedArg0 := symarity
__typedArg1 := tmp19451
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19453 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19452)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19452
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19454 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symappend, tmp19453)
}
__typedArg0 := symappend
__typedArg1 := tmp19453
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19455 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19454)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19454
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19456 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symand, tmp19455)
}
__typedArg0 := symand
__typedArg1 := tmp19455
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19457 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(2), tmp19456)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp19456
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19458 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symadjoin, tmp19457)
}
__typedArg0 := symadjoin
__typedArg1 := tmp19457
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19459 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(3), tmp19458)
}
__typedArg0 := MakeInteger(3)
__typedArg1 := tmp19458
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19460 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symaddress_1_6, tmp19459)
}
__typedArg0 := symaddress_1_6
__typedArg1 := tmp19459
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19461 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19460)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19460
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19462 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symabsvector, tmp19461)
}
__typedArg0 := symabsvector
__typedArg1 := tmp19461
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19463 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19462)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19462
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19464 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symabsvector_2, tmp19463)
}
__typedArg0 := symabsvector_2
__typedArg1 := tmp19463
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19465 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), tmp19464)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp19464
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19466 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symabsolute, tmp19465)
}
__typedArg0 := symabsolute
__typedArg1 := tmp19465
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19467 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp19466)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp19466
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19468 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symabort, tmp19467)
}
__typedArg0 := symabort
__typedArg1 := tmp19467
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19469 := Call(__e, PrimFunc(symshen_4initialise_1arity_1table), tmp19468)


_ = tmp19469

tmp19470 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp19471 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(";"))
}
__typedArg0 := MakeString(";")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp19472 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":="))
}
__typedArg0 := MakeString(":=")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp19473 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(","))
}
__typedArg0 := MakeString(",")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp19474 := Call(__e, PrimFunc(symvector), MakeInteger(0))


tmp19475 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString("bar!"))
}
__typedArg0 := MakeString("bar!")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp19476 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symabort, Nil)
}
__typedArg0 := symabort
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19477 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symabsolute, tmp19476)
}
__typedArg0 := symabsolute
__typedArg1 := tmp19476
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19478 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symabsvector, tmp19477)
}
__typedArg0 := symabsvector
__typedArg1 := tmp19477
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19479 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symabsvector_2, tmp19478)
}
__typedArg0 := symabsvector_2
__typedArg1 := tmp19478
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19480 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symaddress_1_6, tmp19479)
}
__typedArg0 := symaddress_1_6
__typedArg1 := tmp19479
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19481 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1address, tmp19480)
}
__typedArg0 := sym_5_1address
__typedArg1 := tmp19480
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19482 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symadjoin, tmp19481)
}
__typedArg0 := symadjoin
__typedArg1 := tmp19481
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19483 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symand, tmp19482)
}
__typedArg0 := symand
__typedArg1 := tmp19482
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19484 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symappend, tmp19483)
}
__typedArg0 := symappend
__typedArg1 := tmp19483
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19485 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symarity, tmp19484)
}
__typedArg0 := symarity
__typedArg1 := tmp19484
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19486 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symassoc, tmp19485)
}
__typedArg0 := symassoc
__typedArg1 := tmp19485
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19487 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symassertz, tmp19486)
}
__typedArg0 := symassertz
__typedArg1 := tmp19486
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19488 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symasserta, tmp19487)
}
__typedArg0 := symasserta
__typedArg1 := tmp19487
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19489 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symatom_2, tmp19488)
}
__typedArg0 := symatom_2
__typedArg1 := tmp19488
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19490 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstep_2, tmp19489)
}
__typedArg0 := symstep_2
__typedArg1 := tmp19489
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19491 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symspy_2, tmp19490)
}
__typedArg0 := symspy_2
__typedArg1 := tmp19490
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19492 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19475, tmp19491)
}
__typedArg0 := tmp19475
__typedArg1 := tmp19491
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19493 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symbootstrap, tmp19492)
}
__typedArg0 := symbootstrap
__typedArg1 := tmp19492
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19494 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, tmp19493)
}
__typedArg0 := symboolean
__typedArg1 := tmp19493
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19495 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean_2, tmp19494)
}
__typedArg0 := symboolean_2
__typedArg1 := tmp19494
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19496 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symbound_2, tmp19495)
}
__typedArg0 := symbound_2
__typedArg1 := tmp19495
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19497 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symbind, tmp19496)
}
__typedArg0 := symbind
__typedArg1 := tmp19496
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19498 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symclose, tmp19497)
}
__typedArg0 := symclose
__typedArg1 := tmp19497
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19499 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symchar_1stream_2, tmp19498)
}
__typedArg0 := symchar_1stream_2
__typedArg1 := tmp19498
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19500 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcall, tmp19499)
}
__typedArg0 := symcall
__typedArg1 := tmp19499
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19501 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcases, tmp19500)
}
__typedArg0 := symcases
__typedArg1 := tmp19500
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19502 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcd, tmp19501)
}
__typedArg0 := symcd
__typedArg1 := tmp19501
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19503 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcompile, tmp19502)
}
__typedArg0 := symcompile
__typedArg1 := tmp19502
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19504 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symconcat, tmp19503)
}
__typedArg0 := symconcat
__typedArg1 := tmp19503
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19505 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcond, tmp19504)
}
__typedArg0 := symcond
__typedArg1 := tmp19504
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19506 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp19505)
}
__typedArg0 := symcons
__typedArg1 := tmp19505
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19507 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons_2, tmp19506)
}
__typedArg0 := symcons_2
__typedArg1 := tmp19506
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19508 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcn, tmp19507)
}
__typedArg0 := symcn
__typedArg1 := tmp19507
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19509 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symctxt, tmp19508)
}
__typedArg0 := symctxt
__typedArg1 := tmp19508
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19510 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdatatypes, tmp19509)
}
__typedArg0 := symdatatypes
__typedArg1 := tmp19509
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19511 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdatatype, tmp19510)
}
__typedArg0 := symdatatype
__typedArg1 := tmp19510
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19512 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdeclare, tmp19511)
}
__typedArg0 := symdeclare
__typedArg1 := tmp19511
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19513 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefprolog, tmp19512)
}
__typedArg0 := symdefprolog
__typedArg1 := tmp19512
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19514 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefcc, tmp19513)
}
__typedArg0 := symdefcc
__typedArg1 := tmp19513
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19515 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefmacro, tmp19514)
}
__typedArg0 := symdefmacro
__typedArg1 := tmp19514
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19516 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefine, tmp19515)
}
__typedArg0 := symdefine
__typedArg1 := tmp19515
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19517 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefun, tmp19516)
}
__typedArg0 := symdefun
__typedArg1 := tmp19516
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19518 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdestroy, tmp19517)
}
__typedArg0 := symdestroy
__typedArg1 := tmp19517
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19519 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdifference, tmp19518)
}
__typedArg0 := symdifference
__typedArg1 := tmp19518
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19520 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdo, tmp19519)
}
__typedArg0 := symdo
__typedArg1 := tmp19519
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19521 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symelement_2, tmp19520)
}
__typedArg0 := symelement_2
__typedArg1 := tmp19520
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19522 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symempty_2, tmp19521)
}
__typedArg0 := symempty_2
__typedArg1 := tmp19521
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19523 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symerror, tmp19522)
}
__typedArg0 := symerror
__typedArg1 := tmp19522
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19524 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symerror_1to_1string, tmp19523)
}
__typedArg0 := symerror_1to_1string
__typedArg1 := tmp19523
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19525 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symeval, tmp19524)
}
__typedArg0 := symeval
__typedArg1 := tmp19524
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19526 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symeval_1kl, tmp19525)
}
__typedArg0 := symeval_1kl
__typedArg1 := tmp19525
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19527 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symexception, tmp19526)
}
__typedArg0 := symexception
__typedArg1 := tmp19526
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19528 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symexternal, tmp19527)
}
__typedArg0 := symexternal
__typedArg1 := tmp19527
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19529 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symexplode, tmp19528)
}
__typedArg0 := symexplode
__typedArg1 := tmp19528
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19530 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symenable_1type_1theory, tmp19529)
}
__typedArg0 := symenable_1type_1theory
__typedArg1 := tmp19529
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19531 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(False, tmp19530)
}
__typedArg0 := False
__typedArg1 := tmp19530
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19532 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfindall, tmp19531)
}
__typedArg0 := symfindall
__typedArg1 := tmp19531
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19533 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfactorise_2, tmp19532)
}
__typedArg0 := symfactorise_2
__typedArg1 := tmp19532
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19534 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfactorise, tmp19533)
}
__typedArg0 := symfactorise
__typedArg1 := tmp19533
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19535 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfail_1if, tmp19534)
}
__typedArg0 := symfail_1if
__typedArg1 := tmp19534
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19536 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfail, tmp19535)
}
__typedArg0 := symfail
__typedArg1 := tmp19535
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19537 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfile, tmp19536)
}
__typedArg0 := symfile
__typedArg1 := tmp19536
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19538 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfix, tmp19537)
}
__typedArg0 := symfix
__typedArg1 := tmp19537
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19539 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symforeign, tmp19538)
}
__typedArg0 := symforeign
__typedArg1 := tmp19538
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19540 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfork, tmp19539)
}
__typedArg0 := symfork
__typedArg1 := tmp19539
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19541 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfresh, tmp19540)
}
__typedArg0 := symfresh
__typedArg1 := tmp19540
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19542 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfreeze, tmp19541)
}
__typedArg0 := symfreeze
__typedArg1 := tmp19541
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19543 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfst, tmp19542)
}
__typedArg0 := symfst
__typedArg1 := tmp19542
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19544 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfunction, tmp19543)
}
__typedArg0 := symfunction
__typedArg1 := tmp19543
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19545 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfn, tmp19544)
}
__typedArg0 := symfn
__typedArg1 := tmp19544
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19546 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symgensym, tmp19545)
}
__typedArg0 := symgensym
__typedArg1 := tmp19545
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19547 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symget_1time, tmp19546)
}
__typedArg0 := symget_1time
__typedArg1 := tmp19546
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19548 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symget, tmp19547)
}
__typedArg0 := symget
__typedArg1 := tmp19547
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19549 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhash, tmp19548)
}
__typedArg0 := symhash
__typedArg1 := tmp19548
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19550 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhdstr, tmp19549)
}
__typedArg0 := symhdstr
__typedArg1 := tmp19549
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19551 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhdv, tmp19550)
}
__typedArg0 := symhdv
__typedArg1 := tmp19550
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19552 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhd, tmp19551)
}
__typedArg0 := symhd
__typedArg1 := tmp19551
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19553 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhead, tmp19552)
}
__typedArg0 := symhead
__typedArg1 := tmp19552
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19554 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhush, tmp19553)
}
__typedArg0 := symhush
__typedArg1 := tmp19553
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19555 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhush_2, tmp19554)
}
__typedArg0 := symhush_2
__typedArg1 := tmp19554
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19556 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp19555)
}
__typedArg0 := symif
__typedArg1 := tmp19555
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19557 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symimplementation, tmp19556)
}
__typedArg0 := symimplementation
__typedArg1 := tmp19556
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19558 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminternal, tmp19557)
}
__typedArg0 := syminternal
__typedArg1 := tmp19557
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19559 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symin_1package, tmp19558)
}
__typedArg0 := symin_1package
__typedArg1 := tmp19558
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19560 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symin, tmp19559)
}
__typedArg0 := symin
__typedArg1 := tmp19559
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19561 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symis_b, tmp19560)
}
__typedArg0 := symis_b
__typedArg1 := tmp19560
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19562 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symis, tmp19561)
}
__typedArg0 := symis
__typedArg1 := tmp19561
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19563 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symit, tmp19562)
}
__typedArg0 := symit
__typedArg1 := tmp19562
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19564 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminclude_1all_1but, tmp19563)
}
__typedArg0 := syminclude_1all_1but
__typedArg1 := tmp19563
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19565 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminclude, tmp19564)
}
__typedArg0 := syminclude
__typedArg1 := tmp19564
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19566 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symincluded, tmp19565)
}
__typedArg0 := symincluded
__typedArg1 := tmp19565
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19567 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminline, tmp19566)
}
__typedArg0 := syminline
__typedArg1 := tmp19566
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19568 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminput_7, tmp19567)
}
__typedArg0 := syminput_7
__typedArg1 := tmp19567
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19569 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminput, tmp19568)
}
__typedArg0 := syminput
__typedArg1 := tmp19568
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19570 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminteger_2, tmp19569)
}
__typedArg0 := syminteger_2
__typedArg1 := tmp19569
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19571 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symintern, tmp19570)
}
__typedArg0 := symintern
__typedArg1 := tmp19570
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19572 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminferences, tmp19571)
}
__typedArg0 := syminferences
__typedArg1 := tmp19571
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19573 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symintersection, tmp19572)
}
__typedArg0 := symintersection
__typedArg1 := tmp19572
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19574 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symis, tmp19573)
}
__typedArg0 := symis
__typedArg1 := tmp19573
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19575 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlanguage, tmp19574)
}
__typedArg0 := symlanguage
__typedArg1 := tmp19574
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19576 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp19575)
}
__typedArg0 := symlambda
__typedArg1 := tmp19575
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19577 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlazy, tmp19576)
}
__typedArg0 := symlazy
__typedArg1 := tmp19576
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19578 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp19577)
}
__typedArg0 := symlet
__typedArg1 := tmp19577
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19579 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlength, tmp19578)
}
__typedArg0 := symlength
__typedArg1 := tmp19578
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19580 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlimit, tmp19579)
}
__typedArg0 := symlimit
__typedArg1 := tmp19579
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19581 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlineread, tmp19580)
}
__typedArg0 := symlineread
__typedArg1 := tmp19580
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19582 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19581)
}
__typedArg0 := symlist
__typedArg1 := tmp19581
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19583 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symloaded, tmp19582)
}
__typedArg0 := symloaded
__typedArg1 := tmp19582
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19584 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symload, tmp19583)
}
__typedArg0 := symload
__typedArg1 := tmp19583
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19585 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symmake_1string, tmp19584)
}
__typedArg0 := symmake_1string
__typedArg1 := tmp19584
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19586 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symmap, tmp19585)
}
__typedArg0 := symmap
__typedArg1 := tmp19585
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19587 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symmapcan, tmp19586)
}
__typedArg0 := symmapcan
__typedArg1 := tmp19586
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19588 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symmaxinferences, tmp19587)
}
__typedArg0 := symmaxinferences
__typedArg1 := tmp19587
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19589 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symmacroexpand, tmp19588)
}
__typedArg0 := symmacroexpand
__typedArg1 := tmp19588
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19590 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symmode, tmp19589)
}
__typedArg0 := symmode
__typedArg1 := tmp19589
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19591 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnl, tmp19590)
}
__typedArg0 := symnl
__typedArg1 := tmp19590
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19592 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnot, tmp19591)
}
__typedArg0 := symnot
__typedArg1 := tmp19591
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19593 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnth, tmp19592)
}
__typedArg0 := symnth
__typedArg1 := tmp19592
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19594 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnull, tmp19593)
}
__typedArg0 := symnull
__typedArg1 := tmp19593
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19595 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp19594)
}
__typedArg0 := symnumber
__typedArg1 := tmp19594
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19596 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber_2, tmp19595)
}
__typedArg0 := symnumber_2
__typedArg1 := tmp19595
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19597 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symn_1_6string, tmp19596)
}
__typedArg0 := symn_1_6string
__typedArg1 := tmp19596
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19598 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symoccurs_2, tmp19597)
}
__typedArg0 := symoccurs_2
__typedArg1 := tmp19597
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19599 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symoccurs_1check, tmp19598)
}
__typedArg0 := symoccurs_1check
__typedArg1 := tmp19598
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19600 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symoccurrences, tmp19599)
}
__typedArg0 := symoccurrences
__typedArg1 := tmp19599
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19601 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symopen, tmp19600)
}
__typedArg0 := symopen
__typedArg1 := tmp19600
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19602 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symoptimise_2, tmp19601)
}
__typedArg0 := symoptimise_2
__typedArg1 := tmp19601
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19603 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symoptimise, tmp19602)
}
__typedArg0 := symoptimise
__typedArg1 := tmp19602
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19604 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symor, tmp19603)
}
__typedArg0 := symor
__typedArg1 := tmp19603
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19605 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symos, tmp19604)
}
__typedArg0 := symos
__typedArg1 := tmp19604
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19606 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symout, tmp19605)
}
__typedArg0 := symout
__typedArg1 := tmp19605
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19607 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symoutput, tmp19606)
}
__typedArg0 := symoutput
__typedArg1 := tmp19606
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19608 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympackage, tmp19607)
}
__typedArg0 := sympackage
__typedArg1 := tmp19607
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19609 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symport, tmp19608)
}
__typedArg0 := symport
__typedArg1 := tmp19608
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19610 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symporters, tmp19609)
}
__typedArg0 := symporters
__typedArg1 := tmp19609
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19611 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympos, tmp19610)
}
__typedArg0 := sympos
__typedArg1 := tmp19610
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19612 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympr, tmp19611)
}
__typedArg0 := sympr
__typedArg1 := tmp19611
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19613 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symprint, tmp19612)
}
__typedArg0 := symprint
__typedArg1 := tmp19612
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19614 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symprolog_1memory, tmp19613)
}
__typedArg0 := symprolog_1memory
__typedArg1 := tmp19613
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19615 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symprofile, tmp19614)
}
__typedArg0 := symprofile
__typedArg1 := tmp19614
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19616 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symprofile_1results, tmp19615)
}
__typedArg0 := symprofile_1results
__typedArg1 := tmp19615
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19617 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symprotect, tmp19616)
}
__typedArg0 := symprotect
__typedArg1 := tmp19616
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19618 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symprolog_2, tmp19617)
}
__typedArg0 := symprolog_2
__typedArg1 := tmp19617
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19619 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symps, tmp19618)
}
__typedArg0 := symps
__typedArg1 := tmp19618
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19620 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympreclude_1all_1but, tmp19619)
}
__typedArg0 := sympreclude_1all_1but
__typedArg1 := tmp19619
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19621 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympreclude, tmp19620)
}
__typedArg0 := sympreclude
__typedArg1 := tmp19620
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19622 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symput, tmp19621)
}
__typedArg0 := symput
__typedArg1 := tmp19621
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19623 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympackage_2, tmp19622)
}
__typedArg0 := sympackage_2
__typedArg1 := tmp19622
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19624 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1unit_1string, tmp19623)
}
__typedArg0 := symread_1unit_1string
__typedArg1 := tmp19623
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19625 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1from_1string_1unprocessed, tmp19624)
}
__typedArg0 := symread_1from_1string_1unprocessed
__typedArg1 := tmp19624
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19626 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1from_1string, tmp19625)
}
__typedArg0 := symread_1from_1string
__typedArg1 := tmp19625
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19627 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1byte, tmp19626)
}
__typedArg0 := symread_1byte
__typedArg1 := tmp19626
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19628 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1file_1as_1string, tmp19627)
}
__typedArg0 := symread_1file_1as_1string
__typedArg1 := tmp19627
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19629 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1file_1as_1bytelist, tmp19628)
}
__typedArg0 := symread_1file_1as_1bytelist
__typedArg1 := tmp19628
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19630 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1file, tmp19629)
}
__typedArg0 := symread_1file
__typedArg1 := tmp19629
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19631 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symreceive, tmp19630)
}
__typedArg0 := symreceive
__typedArg1 := tmp19630
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19632 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread, tmp19631)
}
__typedArg0 := symread
__typedArg1 := tmp19631
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19633 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symrelease, tmp19632)
}
__typedArg0 := symrelease
__typedArg1 := tmp19632
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19634 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symremove, tmp19633)
}
__typedArg0 := symremove
__typedArg1 := tmp19633
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19635 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symretract, tmp19634)
}
__typedArg0 := symretract
__typedArg1 := tmp19634
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19636 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symreverse, tmp19635)
}
__typedArg0 := symreverse
__typedArg1 := tmp19635
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19637 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symrun, tmp19636)
}
__typedArg0 := symrun
__typedArg1 := tmp19636
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19638 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstr, tmp19637)
}
__typedArg0 := symstr
__typedArg1 := tmp19637
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19639 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsave, tmp19638)
}
__typedArg0 := symsave
__typedArg1 := tmp19638
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19640 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symset, tmp19639)
}
__typedArg0 := symset
__typedArg1 := tmp19639
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19641 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsimple_1error, tmp19640)
}
__typedArg0 := symsimple_1error
__typedArg1 := tmp19640
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19642 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsnd, tmp19641)
}
__typedArg0 := symsnd
__typedArg1 := tmp19641
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19643 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symspecialise, tmp19642)
}
__typedArg0 := symspecialise
__typedArg1 := tmp19642
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19644 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symspy, tmp19643)
}
__typedArg0 := symspy
__typedArg1 := tmp19643
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19645 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsqts, tmp19644)
}
__typedArg0 := symsqts
__typedArg1 := tmp19644
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19646 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstep, tmp19645)
}
__typedArg0 := symstep
__typedArg1 := tmp19645
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19647 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstoutput, tmp19646)
}
__typedArg0 := symstoutput
__typedArg1 := tmp19646
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19648 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsterror, tmp19647)
}
__typedArg0 := symsterror
__typedArg1 := tmp19647
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19649 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstinput, tmp19648)
}
__typedArg0 := symstinput
__typedArg1 := tmp19648
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19650 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp19649)
}
__typedArg0 := symstring
__typedArg1 := tmp19649
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19651 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp19650)
}
__typedArg0 := symstream
__typedArg1 := tmp19650
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19652 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring_1_6n, tmp19651)
}
__typedArg0 := symstring_1_6n
__typedArg1 := tmp19651
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19653 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring_2, tmp19652)
}
__typedArg0 := symstring_2
__typedArg1 := tmp19652
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19654 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsubst, tmp19653)
}
__typedArg0 := symsubst
__typedArg1 := tmp19653
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19655 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsum, tmp19654)
}
__typedArg0 := symsum
__typedArg1 := tmp19654
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19656 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring_1_6symbol, tmp19655)
}
__typedArg0 := symstring_1_6symbol
__typedArg1 := tmp19655
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19657 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol_2, tmp19656)
}
__typedArg0 := symsymbol_2
__typedArg1 := tmp19656
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19658 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp19657)
}
__typedArg0 := symsymbol
__typedArg1 := tmp19657
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19659 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsystem_1S_2, tmp19658)
}
__typedArg0 := symsystem_1S_2
__typedArg1 := tmp19658
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19660 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsynonyms, tmp19659)
}
__typedArg0 := symsynonyms
__typedArg1 := tmp19659
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19661 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsystemf, tmp19660)
}
__typedArg0 := symsystemf
__typedArg1 := tmp19660
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19662 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtail, tmp19661)
}
__typedArg0 := symtail
__typedArg1 := tmp19661
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19663 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtlv, tmp19662)
}
__typedArg0 := symtlv
__typedArg1 := tmp19662
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19664 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtlstr, tmp19663)
}
__typedArg0 := symtlstr
__typedArg1 := tmp19663
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19665 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtl, tmp19664)
}
__typedArg0 := symtl
__typedArg1 := tmp19664
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19666 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtc, tmp19665)
}
__typedArg0 := symtc
__typedArg1 := tmp19665
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19667 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtc_2, tmp19666)
}
__typedArg0 := symtc_2
__typedArg1 := tmp19666
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19668 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symthaw, tmp19667)
}
__typedArg0 := symthaw
__typedArg1 := tmp19667
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19669 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtime, tmp19668)
}
__typedArg0 := symtime
__typedArg1 := tmp19668
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19670 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtrack, tmp19669)
}
__typedArg0 := symtrack
__typedArg1 := tmp19669
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19671 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtracked, tmp19670)
}
__typedArg0 := symtracked
__typedArg1 := tmp19670
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19672 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtrap_1error, tmp19671)
}
__typedArg0 := symtrap_1error
__typedArg1 := tmp19671
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19673 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(True, tmp19672)
}
__typedArg0 := True
__typedArg1 := tmp19672
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19674 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtuple_2, tmp19673)
}
__typedArg0 := symtuple_2
__typedArg1 := tmp19673
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19675 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtype, tmp19674)
}
__typedArg0 := symtype
__typedArg1 := tmp19674
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19676 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symreturn, tmp19675)
}
__typedArg0 := symreturn
__typedArg1 := tmp19675
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19677 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunabsolute, tmp19676)
}
__typedArg0 := symunabsolute
__typedArg1 := tmp19676
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19678 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symundefmacro, tmp19677)
}
__typedArg0 := symundefmacro
__typedArg1 := tmp19677
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19679 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunprofile, tmp19678)
}
__typedArg0 := symunprofile
__typedArg1 := tmp19678
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19680 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunput, tmp19679)
}
__typedArg0 := symunput
__typedArg1 := tmp19679
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19681 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunion, tmp19680)
}
__typedArg0 := symunion
__typedArg1 := tmp19680
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19682 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunix, tmp19681)
}
__typedArg0 := symunix
__typedArg1 := tmp19681
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19683 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunit, tmp19682)
}
__typedArg0 := symunit
__typedArg1 := tmp19682
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19684 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symuntrack, tmp19683)
}
__typedArg0 := symuntrack
__typedArg1 := tmp19683
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19685 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunspecialise, tmp19684)
}
__typedArg0 := symunspecialise
__typedArg1 := tmp19684
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19686 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symupdate_1lambda_1table, tmp19685)
}
__typedArg0 := symupdate_1lambda_1table
__typedArg1 := tmp19685
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19687 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symu_b, tmp19686)
}
__typedArg0 := symu_b
__typedArg1 := tmp19686
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19688 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symuserdefs, tmp19687)
}
__typedArg0 := symuserdefs
__typedArg1 := tmp19687
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19689 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector_2, tmp19688)
}
__typedArg0 := symvector_2
__typedArg1 := tmp19688
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19690 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp19689)
}
__typedArg0 := symvector
__typedArg1 := tmp19689
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19691 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1vector, tmp19690)
}
__typedArg0 := sym_5_1vector
__typedArg1 := tmp19690
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19692 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector_1_6, tmp19691)
}
__typedArg0 := symvector_1_6
__typedArg1 := tmp19691
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19693 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvalue, tmp19692)
}
__typedArg0 := symvalue
__typedArg1 := tmp19692
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19694 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvar_2, tmp19693)
}
__typedArg0 := symvar_2
__typedArg1 := tmp19693
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19695 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvariable_2, tmp19694)
}
__typedArg0 := symvariable_2
__typedArg1 := tmp19694
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19696 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symverified, tmp19695)
}
__typedArg0 := symverified
__typedArg1 := tmp19695
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19697 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symversion, tmp19696)
}
__typedArg0 := symversion
__typedArg1 := tmp19696
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19698 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwarn, tmp19697)
}
__typedArg0 := symwarn
__typedArg1 := tmp19697
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19699 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhen, tmp19698)
}
__typedArg0 := symwhen
__typedArg1 := tmp19698
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19700 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhere, tmp19699)
}
__typedArg0 := symwhere
__typedArg1 := tmp19699
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19701 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwrite_1byte, tmp19700)
}
__typedArg0 := symwrite_1byte
__typedArg1 := tmp19700
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19702 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwrite_1to_1file, tmp19701)
}
__typedArg0 := symwrite_1to_1file
__typedArg1 := tmp19701
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19703 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwrite_1string, tmp19702)
}
__typedArg0 := symwrite_1string
__typedArg1 := tmp19702
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19704 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symy_1or_1n_2, tmp19703)
}
__typedArg0 := symy_1or_1n_2
__typedArg1 := tmp19703
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19705 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19474, tmp19704)
}
__typedArg0 := tmp19474
__typedArg1 := tmp19704
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19706 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_6_6, tmp19705)
}
__typedArg0 := sym_6_6
__typedArg1 := tmp19705
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19707 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5, tmp19706)
}
__typedArg0 := sym_5
__typedArg1 := tmp19706
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19708 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_a, tmp19707)
}
__typedArg0 := sym_5_a
__typedArg1 := tmp19707
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19709 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_7, tmp19708)
}
__typedArg0 := sym_7
__typedArg1 := tmp19708
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19710 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_d, tmp19709)
}
__typedArg0 := sym_d
__typedArg1 := tmp19709
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19711 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_c, tmp19710)
}
__typedArg0 := sym_c
__typedArg1 := tmp19710
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19712 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1, tmp19711)
}
__typedArg0 := sym_1
__typedArg1 := tmp19711
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19713 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_3, tmp19712)
}
__typedArg0 := sym_3
__typedArg1 := tmp19712
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19714 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5end_6, tmp19713)
}
__typedArg0 := sym_5end_6
__typedArg1 := tmp19713
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19715 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_b_6, tmp19714)
}
__typedArg0 := sym_5_b_6
__typedArg1 := tmp19714
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19716 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_c_4, tmp19715)
}
__typedArg0 := sym_c_4
__typedArg1 := tmp19715
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19717 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_a_a_6, tmp19716)
}
__typedArg0 := sym_a_a_6
__typedArg1 := tmp19716
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19718 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_6, tmp19717)
}
__typedArg0 := sym_6
__typedArg1 := tmp19717
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19719 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_6_a, tmp19718)
}
__typedArg0 := sym_6_a
__typedArg1 := tmp19718
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19720 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_a, tmp19719)
}
__typedArg0 := sym_a
__typedArg1 := tmp19719
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19721 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_a_a, tmp19720)
}
__typedArg0 := sym_a_a
__typedArg1 := tmp19720
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19722 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5e_6, tmp19721)
}
__typedArg0 := sym_5e_6
__typedArg1 := tmp19721
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19723 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_6, tmp19722)
}
__typedArg0 := sym_1_6
__typedArg1 := tmp19722
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19724 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1, tmp19723)
}
__typedArg0 := sym_5_1
__typedArg1 := tmp19723
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19725 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dhush_d, tmp19724)
}
__typedArg0 := sym_dhush_d
__typedArg1 := tmp19724
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19726 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dporters_d, tmp19725)
}
__typedArg0 := sym_dporters_d
__typedArg1 := tmp19725
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19727 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dport_d, tmp19726)
}
__typedArg0 := sym_dport_d
__typedArg1 := tmp19726
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19728 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8s, tmp19727)
}
__typedArg0 := sym_8s
__typedArg1 := tmp19727
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19729 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8p, tmp19728)
}
__typedArg0 := sym_8p
__typedArg1 := tmp19728
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19730 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8v, tmp19729)
}
__typedArg0 := sym_8v
__typedArg1 := tmp19729
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19731 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dproperty_1vector_d, tmp19730)
}
__typedArg0 := sym_dproperty_1vector_d
__typedArg1 := tmp19730
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19732 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_drelease_d, tmp19731)
}
__typedArg0 := sym_drelease_d
__typedArg1 := tmp19731
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19733 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dos_d, tmp19732)
}
__typedArg0 := sym_dos_d
__typedArg1 := tmp19732
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19734 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dmacros_d, tmp19733)
}
__typedArg0 := sym_dmacros_d
__typedArg1 := tmp19733
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19735 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dmaximum_1print_1sequence_1size_d, tmp19734)
}
__typedArg0 := sym_dmaximum_1print_1sequence_1size_d
__typedArg1 := tmp19734
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19736 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dversion_d, tmp19735)
}
__typedArg0 := sym_dversion_d
__typedArg1 := tmp19735
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19737 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dhome_1directory_d, tmp19736)
}
__typedArg0 := sym_dhome_1directory_d
__typedArg1 := tmp19736
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19738 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dstoutput_d, tmp19737)
}
__typedArg0 := sym_dstoutput_d
__typedArg1 := tmp19737
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19739 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dsterror_d, tmp19738)
}
__typedArg0 := sym_dsterror_d
__typedArg1 := tmp19738
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19740 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dstinput_d, tmp19739)
}
__typedArg0 := sym_dstinput_d
__typedArg1 := tmp19739
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19741 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dimplementation_d, tmp19740)
}
__typedArg0 := sym_dimplementation_d
__typedArg1 := tmp19740
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19742 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dlanguage_d, tmp19741)
}
__typedArg0 := sym_dlanguage_d
__typedArg1 := tmp19741
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19743 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym__, tmp19742)
}
__typedArg0 := sym__
__typedArg1 := tmp19742
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19744 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19473, tmp19743)
}
__typedArg0 := tmp19473
__typedArg1 := tmp19743
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19745 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19472, tmp19744)
}
__typedArg0 := tmp19472
__typedArg1 := tmp19744
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19746 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19471, tmp19745)
}
__typedArg0 := tmp19471
__typedArg1 := tmp19745
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19747 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19470, tmp19746)
}
__typedArg0 := tmp19470
__typedArg1 := tmp19746
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19748 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_e_e, tmp19747)
}
__typedArg0 := sym_e_e
__typedArg1 := tmp19747
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19749 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1_1, tmp19748)
}
__typedArg0 := sym_5_1_1
__typedArg1 := tmp19748
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19750 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19749)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19749
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19751 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_i, tmp19750)
}
__typedArg0 := sym_i
__typedArg1 := tmp19750
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19752 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_j, tmp19751)
}
__typedArg0 := sym_j
__typedArg1 := tmp19751
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19753 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_b, tmp19752)
}
__typedArg0 := sym_b
__typedArg1 := tmp19752
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19754 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp19755 := Call(__e, PrimFunc(symput), symshen, symshen_4external_1symbols, tmp19753, tmp19754)


_ = tmp19755

tmp19756 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symabsvector) {
return PrimAbsvector(MakeInteger(0))
}
__typedArg0 := MakeInteger(0)
return Call(__e, PrimFunc(symabsvector), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dempty_1absvector_d, tmp19756)
}
__typedArg0 := symshen_4_dempty_1absvector_d
__typedArg1 := tmp19756
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return


}, 0)

tmp19757 := Call(__e, ns2_1set, symshen_4initialise_1environment, tmp18993)


_ = tmp19757

tmp19758 := MakeNative(func(__e *ControlFlow) {
tmp19759 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, Nil)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19759

tmp19760 := MakeNative(func(__e *ControlFlow) {
V6073 := __e.Get(1)
_ = V6073
__e.Return(MakeNative(func(__e *ControlFlow) {
B6069 := __e.Get(1)
_ = B6069
__e.Return(MakeNative(func(__e *ControlFlow) {
L6070 := __e.Get(1)
_ = L6070
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6071 := __e.Get(1)
_ = Key6071
__e.Return(MakeNative(func(__e *ControlFlow) {
C6072 := __e.Get(1)
_ = C6072
tmp19761 := Call(__e, PrimFunc(symshen_4newpv), B6069)


let__18906 := tmp19761
_ = let__18906

tmp19762 := Call(__e, PrimFunc(symshen_4newpv), B6069)


let__18907 := tmp19762
_ = let__18907

tmp19763 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18907, Nil)
}
__typedArg0 := let__18907
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19764 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19763)
}
__typedArg0 := symlist
__typedArg1 := tmp19763
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19765 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19764, Nil)
}
__typedArg0 := tmp19764
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19766 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19765)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19765
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19767 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18906, tmp19766)
}
__typedArg0 := let__18906
__typedArg1 := tmp19766
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19768 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19769 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19768)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19768
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19770 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19767, tmp19769)
}
__typedArg0 := tmp19767
__typedArg1 := tmp19769
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19771 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19770, Nil)
}
__typedArg0 := tmp19770
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19772 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19771)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19771
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19773 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp19772)
}
__typedArg0 := symsymbol
__typedArg1 := tmp19772
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19774 := Call(__e, PrimFunc(symis_b), V6073, tmp19773, B6069, L6070, Key6071, C6072)


tmp19775 := Call(__e, PrimFunc(symshen_4gc), B6069, tmp19774)


__e.TailApply(PrimFunc(symshen_4gc), B6069, tmp19775)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19776 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp19777 := Call(__e, PrimFunc(symshen_4assoc_1_6), symshen_4register_1source_1form, tmp19760, tmp19776)


tmp19778 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp19777)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp19777
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19778

tmp19779 := MakeNative(func(__e *ControlFlow) {
V6078 := __e.Get(1)
_ = V6078
__e.Return(MakeNative(func(__e *ControlFlow) {
B6074 := __e.Get(1)
_ = B6074
__e.Return(MakeNative(func(__e *ControlFlow) {
L6075 := __e.Get(1)
_ = L6075
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6076 := __e.Get(1)
_ = Key6076
__e.Return(MakeNative(func(__e *ControlFlow) {
C6077 := __e.Get(1)
_ = C6077
tmp19780 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19781 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19780)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19780
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19782 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp19781)
}
__typedArg0 := symsymbol
__typedArg1 := tmp19781
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6078, tmp19782, B6074, L6075, Key6076, C6077)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19783 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp19784 := Call(__e, PrimFunc(symshen_4assoc_1_6), symshen_4unregister_1source_1form, tmp19779, tmp19783)


tmp19785 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp19784)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp19784
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19785

tmp19786 := MakeNative(func(__e *ControlFlow) {
V6083 := __e.Get(1)
_ = V6083
__e.Return(MakeNative(func(__e *ControlFlow) {
B6079 := __e.Get(1)
_ = B6079
__e.Return(MakeNative(func(__e *ControlFlow) {
L6080 := __e.Get(1)
_ = L6080
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6081 := __e.Get(1)
_ = Key6081
__e.Return(MakeNative(func(__e *ControlFlow) {
C6082 := __e.Get(1)
_ = C6082
tmp19787 := Call(__e, PrimFunc(symshen_4newpv), B6079)


let__18908 := tmp19787
_ = let__18908

tmp19788 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18908, Nil)
}
__typedArg0 := let__18908
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

tmp19790 := Call(__e, PrimFunc(symis_b), V6083, tmp19789, B6079, L6080, Key6081, C6082)


__e.TailApply(PrimFunc(symshen_4gc), B6079, tmp19790)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19791 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp19792 := Call(__e, PrimFunc(symshen_4assoc_1_6), symabort, tmp19786, tmp19791)


tmp19793 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp19792)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp19792
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19793

tmp19794 := MakeNative(func(__e *ControlFlow) {
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
tmp19795 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19796 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19795)
}
__typedArg0 := symlist
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
return PrimCons(symstring, tmp19798)
}
__typedArg0 := symstring
__typedArg1 := tmp19798
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6088, tmp19799, B6084, L6085, Key6086, C6087)
return


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

tmp19801 := Call(__e, PrimFunc(symshen_4assoc_1_6), symabsolute, tmp19794, tmp19800)


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
tmp19804 := Call(__e, PrimFunc(symshen_4newpv), B6089)


let__18909 := tmp19804
_ = let__18909

tmp19805 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
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
return PrimCons(let__18909, tmp19806)
}
__typedArg0 := let__18909
__typedArg1 := tmp19806
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19808 := Call(__e, PrimFunc(symis_b), V6093, tmp19807, B6089, L6090, Key6091, C6092)


__e.TailApply(PrimFunc(symshen_4gc), B6089, tmp19808)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19809 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp19810 := Call(__e, PrimFunc(symshen_4assoc_1_6), symabsvector_2, tmp19803, tmp19809)


tmp19811 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp19810)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp19810
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19811

tmp19812 := MakeNative(func(__e *ControlFlow) {
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
tmp19813 := Call(__e, PrimFunc(symshen_4newpv), B6095)


let__18910 := tmp19813
_ = let__18910

tmp19814 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18910, Nil)
}
__typedArg0 := let__18910
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19815 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19814)
}
__typedArg0 := symlist
__typedArg1 := tmp19814
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19816 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18910, Nil)
}
__typedArg0 := let__18910
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19817 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19816)
}
__typedArg0 := symlist
__typedArg1 := tmp19816
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19818 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19817, Nil)
}
__typedArg0 := tmp19817
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19819 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19818)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19818
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19820 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19815, tmp19819)
}
__typedArg0 := tmp19815
__typedArg1 := tmp19819
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19821 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19820, Nil)
}
__typedArg0 := tmp19820
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19822 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19821)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19821
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19823 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18910, tmp19822)
}
__typedArg0 := let__18910
__typedArg1 := tmp19822
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19824 := Call(__e, PrimFunc(symis_b), V6099, tmp19823, B6095, L6096, Key6097, C6098)


__e.TailApply(PrimFunc(symshen_4gc), B6095, tmp19824)
return


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

tmp19826 := Call(__e, PrimFunc(symshen_4assoc_1_6), symadjoin, tmp19812, tmp19825)


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
tmp19829 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
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
return PrimCons(symboolean, tmp19830)
}
__typedArg0 := symboolean
__typedArg1 := tmp19830
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19832 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19831, Nil)
}
__typedArg0 := tmp19831
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19833 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19832)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19832
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19834 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, tmp19833)
}
__typedArg0 := symboolean
__typedArg1 := tmp19833
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6104, tmp19834, B6100, L6101, Key6102, C6103)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19835 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp19836 := Call(__e, PrimFunc(symshen_4assoc_1_6), symand, tmp19828, tmp19835)


tmp19837 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp19836)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp19836
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19837

tmp19838 := MakeNative(func(__e *ControlFlow) {
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
tmp19839 := Call(__e, PrimFunc(symshen_4newpv), B6105)


let__18911 := tmp19839
_ = let__18911

tmp19840 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19841 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19840)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19840
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19842 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp19841)
}
__typedArg0 := symsymbol
__typedArg1 := tmp19841
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19843 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19842, Nil)
}
__typedArg0 := tmp19842
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19844 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19843)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19843
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19845 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp19844)
}
__typedArg0 := symstring
__typedArg1 := tmp19844
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19846 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19845, Nil)
}
__typedArg0 := tmp19845
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19847 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19846)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19846
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19848 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18911, tmp19847)
}
__typedArg0 := let__18911
__typedArg1 := tmp19847
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19849 := Call(__e, PrimFunc(symis_b), V6109, tmp19848, B6105, L6106, Key6107, C6108)


__e.TailApply(PrimFunc(symshen_4gc), B6105, tmp19849)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19850 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp19851 := Call(__e, PrimFunc(symshen_4assoc_1_6), symshen_4app, tmp19838, tmp19850)


tmp19852 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp19851)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp19851
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19852

tmp19853 := MakeNative(func(__e *ControlFlow) {
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
tmp19854 := Call(__e, PrimFunc(symshen_4newpv), B6112)


let__18912 := tmp19854
_ = let__18912

tmp19855 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18912, Nil)
}
__typedArg0 := let__18912
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19856 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19855)
}
__typedArg0 := symlist
__typedArg1 := tmp19855
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19857 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18912, Nil)
}
__typedArg0 := let__18912
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19858 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19857)
}
__typedArg0 := symlist
__typedArg1 := tmp19857
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19859 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18912, Nil)
}
__typedArg0 := let__18912
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19860 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19859)
}
__typedArg0 := symlist
__typedArg1 := tmp19859
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19861 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19860, Nil)
}
__typedArg0 := tmp19860
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19862 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19861)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19861
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19863 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19858, tmp19862)
}
__typedArg0 := tmp19858
__typedArg1 := tmp19862
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19864 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19863, Nil)
}
__typedArg0 := tmp19863
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19865 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19864)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19864
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19866 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19856, tmp19865)
}
__typedArg0 := tmp19856
__typedArg1 := tmp19865
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19867 := Call(__e, PrimFunc(symis_b), V6116, tmp19866, B6112, L6113, Key6114, C6115)


__e.TailApply(PrimFunc(symshen_4gc), B6112, tmp19867)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19868 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp19869 := Call(__e, PrimFunc(symshen_4assoc_1_6), symappend, tmp19853, tmp19868)


tmp19870 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp19869)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp19869
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19870

tmp19871 := MakeNative(func(__e *ControlFlow) {
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
tmp19872 := Call(__e, PrimFunc(symshen_4newpv), B6117)


let__18913 := tmp19872
_ = let__18913

tmp19873 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19874 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19873)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19873
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19875 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18913, tmp19874)
}
__typedArg0 := let__18913
__typedArg1 := tmp19874
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19876 := Call(__e, PrimFunc(symis_b), V6121, tmp19875, B6117, L6118, Key6119, C6120)


__e.TailApply(PrimFunc(symshen_4gc), B6117, tmp19876)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19877 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp19878 := Call(__e, PrimFunc(symshen_4assoc_1_6), symarity, tmp19871, tmp19877)


tmp19879 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp19878)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp19878
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19879

tmp19880 := MakeNative(func(__e *ControlFlow) {
V6127 := __e.Get(1)
_ = V6127
__e.Return(MakeNative(func(__e *ControlFlow) {
B6123 := __e.Get(1)
_ = B6123
__e.Return(MakeNative(func(__e *ControlFlow) {
L6124 := __e.Get(1)
_ = L6124
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6125 := __e.Get(1)
_ = Key6125
__e.Return(MakeNative(func(__e *ControlFlow) {
C6126 := __e.Get(1)
_ = C6126
tmp19881 := Call(__e, PrimFunc(symshen_4newpv), B6123)


let__18914 := tmp19881
_ = let__18914

tmp19882 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18914, Nil)
}
__typedArg0 := let__18914
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
return PrimCons(symlist, tmp19884)
}
__typedArg0 := symlist
__typedArg1 := tmp19884
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19886 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18914, Nil)
}
__typedArg0 := let__18914
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19887 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19886)
}
__typedArg0 := symlist
__typedArg1 := tmp19886
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19888 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19887, Nil)
}
__typedArg0 := tmp19887
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19889 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19888)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19888
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19890 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19885, tmp19889)
}
__typedArg0 := tmp19885
__typedArg1 := tmp19889
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19891 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19890, Nil)
}
__typedArg0 := tmp19890
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19892 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19891)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19891
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19893 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18914, tmp19892)
}
__typedArg0 := let__18914
__typedArg1 := tmp19892
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19894 := Call(__e, PrimFunc(symis_b), V6127, tmp19893, B6123, L6124, Key6125, C6126)


__e.TailApply(PrimFunc(symshen_4gc), B6123, tmp19894)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19895 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp19896 := Call(__e, PrimFunc(symshen_4assoc_1_6), symassoc, tmp19880, tmp19895)


tmp19897 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp19896)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp19896
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19897

tmp19898 := MakeNative(func(__e *ControlFlow) {
V6132 := __e.Get(1)
_ = V6132
__e.Return(MakeNative(func(__e *ControlFlow) {
B6128 := __e.Get(1)
_ = B6128
__e.Return(MakeNative(func(__e *ControlFlow) {
L6129 := __e.Get(1)
_ = L6129
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6130 := __e.Get(1)
_ = Key6130
__e.Return(MakeNative(func(__e *ControlFlow) {
C6131 := __e.Get(1)
_ = C6131
tmp19899 := Call(__e, PrimFunc(symshen_4newpv), B6128)


let__18915 := tmp19899
_ = let__18915

tmp19900 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19901 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19900)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19900
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19902 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18915, tmp19901)
}
__typedArg0 := let__18915
__typedArg1 := tmp19901
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19903 := Call(__e, PrimFunc(symis_b), V6132, tmp19902, B6128, L6129, Key6130, C6131)


__e.TailApply(PrimFunc(symshen_4gc), B6128, tmp19903)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19904 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp19905 := Call(__e, PrimFunc(symshen_4assoc_1_6), symatom_2, tmp19898, tmp19904)


tmp19906 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp19905)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp19905
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19906

tmp19907 := MakeNative(func(__e *ControlFlow) {
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
tmp19908 := Call(__e, PrimFunc(symshen_4newpv), B6133)


let__18916 := tmp19908
_ = let__18916

tmp19909 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19910 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19909)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19909
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19911 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18916, tmp19910)
}
__typedArg0 := let__18916
__typedArg1 := tmp19910
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19912 := Call(__e, PrimFunc(symis_b), V6137, tmp19911, B6133, L6134, Key6135, C6136)


__e.TailApply(PrimFunc(symshen_4gc), B6133, tmp19912)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19913 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp19914 := Call(__e, PrimFunc(symshen_4assoc_1_6), symboolean_2, tmp19907, tmp19913)


tmp19915 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp19914)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp19914
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19915

tmp19916 := MakeNative(func(__e *ControlFlow) {
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
tmp19917 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19918 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19917)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19917
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19919 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp19918)
}
__typedArg0 := symstring
__typedArg1 := tmp19918
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6142, tmp19919, B6138, L6139, Key6140, C6141)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19920 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp19921 := Call(__e, PrimFunc(symshen_4assoc_1_6), symbootstrap, tmp19916, tmp19920)


tmp19922 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp19921)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp19921
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19922

tmp19923 := MakeNative(func(__e *ControlFlow) {
V6147 := __e.Get(1)
_ = V6147
__e.Return(MakeNative(func(__e *ControlFlow) {
B6143 := __e.Get(1)
_ = B6143
__e.Return(MakeNative(func(__e *ControlFlow) {
L6144 := __e.Get(1)
_ = L6144
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6145 := __e.Get(1)
_ = Key6145
__e.Return(MakeNative(func(__e *ControlFlow) {
C6146 := __e.Get(1)
_ = C6146
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
return PrimCons(symsymbol, tmp19925)
}
__typedArg0 := symsymbol
__typedArg1 := tmp19925
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6147, tmp19926, B6143, L6144, Key6145, C6146)
return


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

tmp19928 := Call(__e, PrimFunc(symshen_4assoc_1_6), symbound_2, tmp19923, tmp19927)


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
tmp19931 := Call(__e, PrimFunc(symshen_4newpv), B6148)


let__18917 := tmp19931
_ = let__18917

tmp19932 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18917, Nil)
}
__typedArg0 := let__18917
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19933 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19932)
}
__typedArg0 := symlist
__typedArg1 := tmp19932
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19934 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19935 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19934)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19934
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19936 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19933, tmp19935)
}
__typedArg0 := tmp19933
__typedArg1 := tmp19935
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19937 := Call(__e, PrimFunc(symis_b), V6152, tmp19936, B6148, L6149, Key6150, C6151)


__e.TailApply(PrimFunc(symshen_4gc), B6148, tmp19937)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19938 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp19939 := Call(__e, PrimFunc(symshen_4assoc_1_6), symshen_4ccons_2, tmp19930, tmp19938)


tmp19940 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp19939)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp19939
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19940

tmp19941 := MakeNative(func(__e *ControlFlow) {
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
tmp19942 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19943 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19942)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19942
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19944 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp19943)
}
__typedArg0 := symstring
__typedArg1 := tmp19943
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6157, tmp19944, B6153, L6154, Key6155, C6156)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19945 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp19946 := Call(__e, PrimFunc(symshen_4assoc_1_6), symcd, tmp19941, tmp19945)


tmp19947 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp19946)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp19946
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19947

tmp19948 := MakeNative(func(__e *ControlFlow) {
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
tmp19949 := Call(__e, PrimFunc(symshen_4newpv), B6159)


let__18918 := tmp19949
_ = let__18918

tmp19950 := Call(__e, PrimFunc(symshen_4newpv), B6159)


let__18919 := tmp19950
_ = let__18919

tmp19951 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18918, Nil)
}
__typedArg0 := let__18918
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19952 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp19951)
}
__typedArg0 := symstream
__typedArg1 := tmp19951
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19953 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18919, Nil)
}
__typedArg0 := let__18919
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19954 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19953)
}
__typedArg0 := symlist
__typedArg1 := tmp19953
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19955 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19954, Nil)
}
__typedArg0 := tmp19954
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19956 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19955)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19955
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19957 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19952, tmp19956)
}
__typedArg0 := tmp19952
__typedArg1 := tmp19956
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19958 := Call(__e, PrimFunc(symis_b), V6163, tmp19957, B6159, L6160, Key6161, C6162)


tmp19959 := Call(__e, PrimFunc(symshen_4gc), B6159, tmp19958)


__e.TailApply(PrimFunc(symshen_4gc), B6159, tmp19959)
return


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

tmp19961 := Call(__e, PrimFunc(symshen_4assoc_1_6), symclose, tmp19948, tmp19960)


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
tmp19964 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19965 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19964)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19964
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19966 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp19965)
}
__typedArg0 := symstring
__typedArg1 := tmp19965
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19967 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19966, Nil)
}
__typedArg0 := tmp19966
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
return PrimCons(symstring, tmp19968)
}
__typedArg0 := symstring
__typedArg1 := tmp19968
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6168, tmp19969, B6164, L6165, Key6166, C6167)
return


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

tmp19971 := Call(__e, PrimFunc(symshen_4assoc_1_6), symcn, tmp19963, tmp19970)


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
tmp19974 := Call(__e, PrimFunc(symshen_4newpv), B6169)


let__18920 := tmp19974
_ = let__18920

tmp19975 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18920, Nil)
}
__typedArg0 := let__18920
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19976 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp19975)
}
__typedArg0 := symstream
__typedArg1 := tmp19975
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19977 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19978 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19977)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19977
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19979 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19976, tmp19978)
}
__typedArg0 := tmp19976
__typedArg1 := tmp19978
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19980 := Call(__e, PrimFunc(symis_b), V6173, tmp19979, B6169, L6170, Key6171, C6172)


__e.TailApply(PrimFunc(symshen_4gc), B6169, tmp19980)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19981 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp19982 := Call(__e, PrimFunc(symshen_4assoc_1_6), symchar_1stream_2, tmp19973, tmp19981)


tmp19983 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp19982)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp19982
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19983

tmp19984 := MakeNative(func(__e *ControlFlow) {
V6180 := __e.Get(1)
_ = V6180
__e.Return(MakeNative(func(__e *ControlFlow) {
B6176 := __e.Get(1)
_ = B6176
__e.Return(MakeNative(func(__e *ControlFlow) {
L6177 := __e.Get(1)
_ = L6177
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6178 := __e.Get(1)
_ = Key6178
__e.Return(MakeNative(func(__e *ControlFlow) {
C6179 := __e.Get(1)
_ = C6179
tmp19985 := Call(__e, PrimFunc(symshen_4newpv), B6176)


let__18921 := tmp19985
_ = let__18921

tmp19986 := Call(__e, PrimFunc(symshen_4newpv), B6176)


let__18922 := tmp19986
_ = let__18922

tmp19987 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18921, Nil)
}
__typedArg0 := let__18921
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
return PrimCons(let__18921, Nil)
}
__typedArg0 := let__18921
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19990 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19989)
}
__typedArg0 := symlist
__typedArg1 := tmp19989
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19991 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18922, Nil)
}
__typedArg0 := let__18922
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19992 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19990, tmp19991)
}
__typedArg0 := tmp19990
__typedArg1 := tmp19991
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19993 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstr, tmp19992)
}
__typedArg0 := symstr
__typedArg1 := tmp19992
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19994 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19993, Nil)
}
__typedArg0 := tmp19993
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19995 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19994)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19994
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19996 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19988, tmp19995)
}
__typedArg0 := tmp19988
__typedArg1 := tmp19995
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19997 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18921, Nil)
}
__typedArg0 := let__18921
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19998 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19997)
}
__typedArg0 := symlist
__typedArg1 := tmp19997
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19999 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18922, Nil)
}
__typedArg0 := let__18922
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
return PrimCons(tmp19998, tmp20000)
}
__typedArg0 := tmp19998
__typedArg1 := tmp20000
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20002 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20001, Nil)
}
__typedArg0 := tmp20001
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
return PrimCons(tmp19996, tmp20003)
}
__typedArg0 := tmp19996
__typedArg1 := tmp20003
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20005 := Call(__e, PrimFunc(symis_b), V6180, tmp20004, B6176, L6177, Key6178, C6179)


tmp20006 := Call(__e, PrimFunc(symshen_4gc), B6176, tmp20005)


__e.TailApply(PrimFunc(symshen_4gc), B6176, tmp20006)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20007 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20008 := Call(__e, PrimFunc(symshen_4assoc_1_6), symcompile, tmp19984, tmp20007)


tmp20009 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20008)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20008
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20009

tmp20010 := MakeNative(func(__e *ControlFlow) {
V6185 := __e.Get(1)
_ = V6185
__e.Return(MakeNative(func(__e *ControlFlow) {
B6181 := __e.Get(1)
_ = B6181
__e.Return(MakeNative(func(__e *ControlFlow) {
L6182 := __e.Get(1)
_ = L6182
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6183 := __e.Get(1)
_ = Key6183
__e.Return(MakeNative(func(__e *ControlFlow) {
C6184 := __e.Get(1)
_ = C6184
tmp20011 := Call(__e, PrimFunc(symshen_4newpv), B6181)


let__18923 := tmp20011
_ = let__18923

tmp20012 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
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
return PrimCons(let__18923, tmp20013)
}
__typedArg0 := let__18923
__typedArg1 := tmp20013
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20015 := Call(__e, PrimFunc(symis_b), V6185, tmp20014, B6181, L6182, Key6183, C6184)


__e.TailApply(PrimFunc(symshen_4gc), B6181, tmp20015)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20016 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20017 := Call(__e, PrimFunc(symshen_4assoc_1_6), symcons_2, tmp20010, tmp20016)


tmp20018 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20017)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20017
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20018

tmp20019 := MakeNative(func(__e *ControlFlow) {
V6190 := __e.Get(1)
_ = V6190
__e.Return(MakeNative(func(__e *ControlFlow) {
B6186 := __e.Get(1)
_ = B6186
__e.Return(MakeNative(func(__e *ControlFlow) {
L6187 := __e.Get(1)
_ = L6187
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6188 := __e.Get(1)
_ = Key6188
__e.Return(MakeNative(func(__e *ControlFlow) {
C6189 := __e.Get(1)
_ = C6189
tmp20020 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
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
return PrimCons(tmp20021, Nil)
}
__typedArg0 := tmp20021
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

__e.TailApply(PrimFunc(symis_b), V6190, tmp20023, B6186, L6187, Key6188, C6189)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20024 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20025 := Call(__e, PrimFunc(symshen_4assoc_1_6), symdatatypes, tmp20019, tmp20024)


tmp20026 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20025)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20025
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20026

tmp20027 := MakeNative(func(__e *ControlFlow) {
V6195 := __e.Get(1)
_ = V6195
__e.Return(MakeNative(func(__e *ControlFlow) {
B6191 := __e.Get(1)
_ = B6191
__e.Return(MakeNative(func(__e *ControlFlow) {
L6192 := __e.Get(1)
_ = L6192
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6193 := __e.Get(1)
_ = Key6193
__e.Return(MakeNative(func(__e *ControlFlow) {
C6194 := __e.Get(1)
_ = C6194
tmp20028 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20029 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20028)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20028
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20030 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp20029)
}
__typedArg0 := symsymbol
__typedArg1 := tmp20029
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6195, tmp20030, B6191, L6192, Key6193, C6194)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20031 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20032 := Call(__e, PrimFunc(symshen_4assoc_1_6), symdestroy, tmp20027, tmp20031)


tmp20033 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20032)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20032
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20033

tmp20034 := MakeNative(func(__e *ControlFlow) {
V6202 := __e.Get(1)
_ = V6202
__e.Return(MakeNative(func(__e *ControlFlow) {
B6198 := __e.Get(1)
_ = B6198
__e.Return(MakeNative(func(__e *ControlFlow) {
L6199 := __e.Get(1)
_ = L6199
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6200 := __e.Get(1)
_ = Key6200
__e.Return(MakeNative(func(__e *ControlFlow) {
C6201 := __e.Get(1)
_ = C6201
tmp20035 := Call(__e, PrimFunc(symshen_4newpv), B6198)


let__18924 := tmp20035
_ = let__18924

tmp20036 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18924, Nil)
}
__typedArg0 := let__18924
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20037 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20036)
}
__typedArg0 := symlist
__typedArg1 := tmp20036
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20038 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18924, Nil)
}
__typedArg0 := let__18924
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20039 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20038)
}
__typedArg0 := symlist
__typedArg1 := tmp20038
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20040 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18924, Nil)
}
__typedArg0 := let__18924
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20041 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20040)
}
__typedArg0 := symlist
__typedArg1 := tmp20040
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20042 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20041, Nil)
}
__typedArg0 := tmp20041
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20043 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20042)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20042
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20044 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20039, tmp20043)
}
__typedArg0 := tmp20039
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

tmp20047 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20037, tmp20046)
}
__typedArg0 := tmp20037
__typedArg1 := tmp20046
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20048 := Call(__e, PrimFunc(symis_b), V6202, tmp20047, B6198, L6199, Key6200, C6201)


__e.TailApply(PrimFunc(symshen_4gc), B6198, tmp20048)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20049 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20050 := Call(__e, PrimFunc(symshen_4assoc_1_6), symdifference, tmp20034, tmp20049)


tmp20051 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20050)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20050
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20051

tmp20052 := MakeNative(func(__e *ControlFlow) {
V6207 := __e.Get(1)
_ = V6207
__e.Return(MakeNative(func(__e *ControlFlow) {
B6203 := __e.Get(1)
_ = B6203
__e.Return(MakeNative(func(__e *ControlFlow) {
L6204 := __e.Get(1)
_ = L6204
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6205 := __e.Get(1)
_ = Key6205
__e.Return(MakeNative(func(__e *ControlFlow) {
C6206 := __e.Get(1)
_ = C6206
tmp20053 := Call(__e, PrimFunc(symshen_4newpv), B6203)


let__18925 := tmp20053
_ = let__18925

tmp20054 := Call(__e, PrimFunc(symshen_4newpv), B6203)


let__18926 := tmp20054
_ = let__18926

tmp20055 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18926, Nil)
}
__typedArg0 := let__18926
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20056 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20055)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20055
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20057 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18926, tmp20056)
}
__typedArg0 := let__18926
__typedArg1 := tmp20056
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20058 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20057, Nil)
}
__typedArg0 := tmp20057
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20059 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20058)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20058
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20060 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18925, tmp20059)
}
__typedArg0 := let__18925
__typedArg1 := tmp20059
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20061 := Call(__e, PrimFunc(symis_b), V6207, tmp20060, B6203, L6204, Key6205, C6206)


tmp20062 := Call(__e, PrimFunc(symshen_4gc), B6203, tmp20061)


__e.TailApply(PrimFunc(symshen_4gc), B6203, tmp20062)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20063 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20064 := Call(__e, PrimFunc(symshen_4assoc_1_6), symdo, tmp20052, tmp20063)


tmp20065 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20064)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20064
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20065

tmp20066 := MakeNative(func(__e *ControlFlow) {
V6214 := __e.Get(1)
_ = V6214
__e.Return(MakeNative(func(__e *ControlFlow) {
B6210 := __e.Get(1)
_ = B6210
__e.Return(MakeNative(func(__e *ControlFlow) {
L6211 := __e.Get(1)
_ = L6211
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6212 := __e.Get(1)
_ = Key6212
__e.Return(MakeNative(func(__e *ControlFlow) {
C6213 := __e.Get(1)
_ = C6213
tmp20067 := Call(__e, PrimFunc(symshen_4newpv), B6210)


let__18927 := tmp20067
_ = let__18927

tmp20068 := Call(__e, PrimFunc(symshen_4newpv), B6210)


let__18928 := tmp20068
_ = let__18928

tmp20069 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18927, Nil)
}
__typedArg0 := let__18927
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20070 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20069)
}
__typedArg0 := symlist
__typedArg1 := tmp20069
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20071 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18927, Nil)
}
__typedArg0 := let__18927
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20072 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20071)
}
__typedArg0 := symlist
__typedArg1 := tmp20071
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20073 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18928, Nil)
}
__typedArg0 := let__18928
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20074 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20073)
}
__typedArg0 := symlist
__typedArg1 := tmp20073
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20075 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20074, Nil)
}
__typedArg0 := tmp20074
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20076 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20072, tmp20075)
}
__typedArg0 := tmp20072
__typedArg1 := tmp20075
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20077 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstr, tmp20076)
}
__typedArg0 := symstr
__typedArg1 := tmp20076
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20078 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20077, Nil)
}
__typedArg0 := tmp20077
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
return PrimCons(tmp20070, tmp20079)
}
__typedArg0 := tmp20070
__typedArg1 := tmp20079
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20081 := Call(__e, PrimFunc(symis_b), V6214, tmp20080, B6210, L6211, Key6212, C6213)


tmp20082 := Call(__e, PrimFunc(symshen_4gc), B6210, tmp20081)


__e.TailApply(PrimFunc(symshen_4gc), B6210, tmp20082)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20083 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20084 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_5e_6, tmp20066, tmp20083)


tmp20085 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20084)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20084
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20085

tmp20086 := MakeNative(func(__e *ControlFlow) {
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
tmp20087 := Call(__e, PrimFunc(symshen_4newpv), B6217)


let__18929 := tmp20087
_ = let__18929

tmp20088 := Call(__e, PrimFunc(symshen_4newpv), B6217)


let__18930 := tmp20088
_ = let__18930

tmp20089 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18930, Nil)
}
__typedArg0 := let__18930
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20090 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20089)
}
__typedArg0 := symlist
__typedArg1 := tmp20089
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20091 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18929, Nil)
}
__typedArg0 := let__18929
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20092 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20091)
}
__typedArg0 := symlist
__typedArg1 := tmp20091
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20093 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18930, Nil)
}
__typedArg0 := let__18930
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20094 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20093)
}
__typedArg0 := symlist
__typedArg1 := tmp20093
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20095 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20094, Nil)
}
__typedArg0 := tmp20094
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20096 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20092, tmp20095)
}
__typedArg0 := tmp20092
__typedArg1 := tmp20095
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20097 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstr, tmp20096)
}
__typedArg0 := symstr
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
return PrimCons(sym_1_1_6, tmp20098)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20098
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20100 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20090, tmp20099)
}
__typedArg0 := tmp20090
__typedArg1 := tmp20099
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20101 := Call(__e, PrimFunc(symis_b), V6221, tmp20100, B6217, L6218, Key6219, C6220)


tmp20102 := Call(__e, PrimFunc(symshen_4gc), B6217, tmp20101)


__e.TailApply(PrimFunc(symshen_4gc), B6217, tmp20102)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20103 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20104 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_5_b_6, tmp20086, tmp20103)


tmp20105 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20104)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20104
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20105

tmp20106 := MakeNative(func(__e *ControlFlow) {
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
tmp20107 := Call(__e, PrimFunc(symshen_4newpv), B6224)


let__18931 := tmp20107
_ = let__18931

tmp20108 := Call(__e, PrimFunc(symshen_4newpv), B6224)


let__18932 := tmp20108
_ = let__18932

tmp20109 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18931, Nil)
}
__typedArg0 := let__18931
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20110 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20109)
}
__typedArg0 := symlist
__typedArg1 := tmp20109
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20111 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18931, Nil)
}
__typedArg0 := let__18931
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20112 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20111)
}
__typedArg0 := symlist
__typedArg1 := tmp20111
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20113 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18932, Nil)
}
__typedArg0 := let__18932
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20114 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20113)
}
__typedArg0 := symlist
__typedArg1 := tmp20113
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20115 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20114, Nil)
}
__typedArg0 := tmp20114
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20116 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20112, tmp20115)
}
__typedArg0 := tmp20112
__typedArg1 := tmp20115
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20117 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstr, tmp20116)
}
__typedArg0 := symstr
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
return PrimCons(sym_1_1_6, tmp20118)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20118
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20120 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20110, tmp20119)
}
__typedArg0 := tmp20110
__typedArg1 := tmp20119
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20121 := Call(__e, PrimFunc(symis_b), V6228, tmp20120, B6224, L6225, Key6226, C6227)


tmp20122 := Call(__e, PrimFunc(symshen_4gc), B6224, tmp20121)


__e.TailApply(PrimFunc(symshen_4gc), B6224, tmp20122)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20123 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20124 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_5end_6, tmp20106, tmp20123)


tmp20125 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20124)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20124
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20125

tmp20126 := MakeNative(func(__e *ControlFlow) {
V6233 := __e.Get(1)
_ = V6233
__e.Return(MakeNative(func(__e *ControlFlow) {
B6229 := __e.Get(1)
_ = B6229
__e.Return(MakeNative(func(__e *ControlFlow) {
L6230 := __e.Get(1)
_ = L6230
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6231 := __e.Get(1)
_ = Key6231
__e.Return(MakeNative(func(__e *ControlFlow) {
C6232 := __e.Get(1)
_ = C6232
tmp20127 := Call(__e, PrimFunc(symshen_4newpv), B6229)


let__18933 := tmp20127
_ = let__18933

tmp20128 := Call(__e, PrimFunc(symshen_4newpv), B6229)


let__18934 := tmp20128
_ = let__18934

tmp20129 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18933, Nil)
}
__typedArg0 := let__18933
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20130 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20129)
}
__typedArg0 := symlist
__typedArg1 := tmp20129
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20131 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18934, Nil)
}
__typedArg0 := let__18934
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20132 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20130, tmp20131)
}
__typedArg0 := tmp20130
__typedArg1 := tmp20131
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20133 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstr, tmp20132)
}
__typedArg0 := symstr
__typedArg1 := tmp20132
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20134 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20135 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20134)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20134
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20136 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20133, tmp20135)
}
__typedArg0 := tmp20133
__typedArg1 := tmp20135
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20137 := Call(__e, PrimFunc(symis_b), V6233, tmp20136, B6229, L6230, Key6231, C6232)


tmp20138 := Call(__e, PrimFunc(symshen_4gc), B6229, tmp20137)


__e.TailApply(PrimFunc(symshen_4gc), B6229, tmp20138)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20139 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20140 := Call(__e, PrimFunc(symshen_4assoc_1_6), symshen_4parse_1failure_2, tmp20126, tmp20139)


tmp20141 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20140)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20140
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20141

tmp20142 := MakeNative(func(__e *ControlFlow) {
V6238 := __e.Get(1)
_ = V6238
__e.Return(MakeNative(func(__e *ControlFlow) {
B6234 := __e.Get(1)
_ = B6234
__e.Return(MakeNative(func(__e *ControlFlow) {
L6235 := __e.Get(1)
_ = L6235
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6236 := __e.Get(1)
_ = Key6236
__e.Return(MakeNative(func(__e *ControlFlow) {
C6237 := __e.Get(1)
_ = C6237
tmp20143 := Call(__e, PrimFunc(symshen_4newpv), B6234)


let__18935 := tmp20143
_ = let__18935

tmp20144 := Call(__e, PrimFunc(symshen_4newpv), B6234)


let__18936 := tmp20144
_ = let__18936

tmp20145 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18935, Nil)
}
__typedArg0 := let__18935
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20146 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20145)
}
__typedArg0 := symlist
__typedArg1 := tmp20145
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20147 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18936, Nil)
}
__typedArg0 := let__18936
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20148 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20146, tmp20147)
}
__typedArg0 := tmp20146
__typedArg1 := tmp20147
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20149 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstr, tmp20148)
}
__typedArg0 := symstr
__typedArg1 := tmp20148
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20150 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20149, Nil)
}
__typedArg0 := tmp20149
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20151 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20150)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20150
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20152 := Call(__e, PrimFunc(symis_b), V6238, tmp20151, B6234, L6235, Key6236, C6237)


tmp20153 := Call(__e, PrimFunc(symshen_4gc), B6234, tmp20152)


__e.TailApply(PrimFunc(symshen_4gc), B6234, tmp20153)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20154 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20155 := Call(__e, PrimFunc(symshen_4assoc_1_6), symshen_4parse_1failure, tmp20142, tmp20154)


tmp20156 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20155)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20155
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20156

tmp20157 := MakeNative(func(__e *ControlFlow) {
V6243 := __e.Get(1)
_ = V6243
__e.Return(MakeNative(func(__e *ControlFlow) {
B6239 := __e.Get(1)
_ = B6239
__e.Return(MakeNative(func(__e *ControlFlow) {
L6240 := __e.Get(1)
_ = L6240
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6241 := __e.Get(1)
_ = Key6241
__e.Return(MakeNative(func(__e *ControlFlow) {
C6242 := __e.Get(1)
_ = C6242
tmp20158 := Call(__e, PrimFunc(symshen_4newpv), B6239)


let__18937 := tmp20158
_ = let__18937

tmp20159 := Call(__e, PrimFunc(symshen_4newpv), B6239)


let__18938 := tmp20159
_ = let__18938

tmp20160 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18937, Nil)
}
__typedArg0 := let__18937
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20161 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20160)
}
__typedArg0 := symlist
__typedArg1 := tmp20160
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20162 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18938, Nil)
}
__typedArg0 := let__18938
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20163 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20161, tmp20162)
}
__typedArg0 := tmp20161
__typedArg1 := tmp20162
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20164 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstr, tmp20163)
}
__typedArg0 := symstr
__typedArg1 := tmp20163
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20165 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18938, Nil)
}
__typedArg0 := let__18938
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20166 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20165)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20165
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20167 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20164, tmp20166)
}
__typedArg0 := tmp20164
__typedArg1 := tmp20166
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20168 := Call(__e, PrimFunc(symis_b), V6243, tmp20167, B6239, L6240, Key6241, C6242)


tmp20169 := Call(__e, PrimFunc(symshen_4gc), B6239, tmp20168)


__e.TailApply(PrimFunc(symshen_4gc), B6239, tmp20169)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20170 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20171 := Call(__e, PrimFunc(symshen_4assoc_1_6), symshen_4_5_1out, tmp20157, tmp20170)


tmp20172 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20171)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20171
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20172

tmp20173 := MakeNative(func(__e *ControlFlow) {
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
tmp20174 := Call(__e, PrimFunc(symshen_4newpv), B6245)


let__18939 := tmp20174
_ = let__18939

tmp20175 := Call(__e, PrimFunc(symshen_4newpv), B6245)


let__18940 := tmp20175
_ = let__18940

tmp20176 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18940, Nil)
}
__typedArg0 := let__18940
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20177 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20176)
}
__typedArg0 := symlist
__typedArg1 := tmp20176
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20178 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18939, Nil)
}
__typedArg0 := let__18939
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20179 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20177, tmp20178)
}
__typedArg0 := tmp20177
__typedArg1 := tmp20178
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20180 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstr, tmp20179)
}
__typedArg0 := symstr
__typedArg1 := tmp20179
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20181 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18940, Nil)
}
__typedArg0 := let__18940
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20182 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20181)
}
__typedArg0 := symlist
__typedArg1 := tmp20181
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20183 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20182, Nil)
}
__typedArg0 := tmp20182
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20184 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20183)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20183
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20185 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20180, tmp20184)
}
__typedArg0 := tmp20180
__typedArg1 := tmp20184
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20186 := Call(__e, PrimFunc(symis_b), V6249, tmp20185, B6245, L6246, Key6247, C6248)


tmp20187 := Call(__e, PrimFunc(symshen_4gc), B6245, tmp20186)


__e.TailApply(PrimFunc(symshen_4gc), B6245, tmp20187)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20188 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20189 := Call(__e, PrimFunc(symshen_4assoc_1_6), symshen_4in_1_6, tmp20173, tmp20188)


tmp20190 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20189)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20189
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20190

tmp20191 := MakeNative(func(__e *ControlFlow) {
V6255 := __e.Get(1)
_ = V6255
__e.Return(MakeNative(func(__e *ControlFlow) {
B6251 := __e.Get(1)
_ = B6251
__e.Return(MakeNative(func(__e *ControlFlow) {
L6252 := __e.Get(1)
_ = L6252
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6253 := __e.Get(1)
_ = Key6253
__e.Return(MakeNative(func(__e *ControlFlow) {
C6254 := __e.Get(1)
_ = C6254
tmp20192 := Call(__e, PrimFunc(symshen_4newpv), B6251)


let__18941 := tmp20192
_ = let__18941

tmp20193 := Call(__e, PrimFunc(symshen_4newpv), B6251)


let__18942 := tmp20193
_ = let__18942

tmp20194 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18941, Nil)
}
__typedArg0 := let__18941
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20195 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20194)
}
__typedArg0 := symlist
__typedArg1 := tmp20194
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20196 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18941, Nil)
}
__typedArg0 := let__18941
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20197 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20196)
}
__typedArg0 := symlist
__typedArg1 := tmp20196
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20198 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18942, Nil)
}
__typedArg0 := let__18942
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20199 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20197, tmp20198)
}
__typedArg0 := tmp20197
__typedArg1 := tmp20198
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20200 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstr, tmp20199)
}
__typedArg0 := symstr
__typedArg1 := tmp20199
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20201 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20200, Nil)
}
__typedArg0 := tmp20200
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20202 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20201)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20201
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20203 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18942, tmp20202)
}
__typedArg0 := let__18942
__typedArg1 := tmp20202
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20204 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20203, Nil)
}
__typedArg0 := tmp20203
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20205 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20204)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20204
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20206 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20195, tmp20205)
}
__typedArg0 := tmp20195
__typedArg1 := tmp20205
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20207 := Call(__e, PrimFunc(symis_b), V6255, tmp20206, B6251, L6252, Key6253, C6254)


tmp20208 := Call(__e, PrimFunc(symshen_4gc), B6251, tmp20207)


__e.TailApply(PrimFunc(symshen_4gc), B6251, tmp20208)
return


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

tmp20210 := Call(__e, PrimFunc(symshen_4assoc_1_6), symshen_4comb, tmp20191, tmp20209)


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
V6260 := __e.Get(1)
_ = V6260
__e.Return(MakeNative(func(__e *ControlFlow) {
B6256 := __e.Get(1)
_ = B6256
__e.Return(MakeNative(func(__e *ControlFlow) {
L6257 := __e.Get(1)
_ = L6257
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6258 := __e.Get(1)
_ = Key6258
__e.Return(MakeNative(func(__e *ControlFlow) {
C6259 := __e.Get(1)
_ = C6259
tmp20213 := Call(__e, PrimFunc(symshen_4newpv), B6256)


let__18943 := tmp20213
_ = let__18943

tmp20214 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18943, Nil)
}
__typedArg0 := let__18943
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
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
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
return PrimCons(tmp20215, tmp20217)
}
__typedArg0 := tmp20215
__typedArg1 := tmp20217
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20219 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20218, Nil)
}
__typedArg0 := tmp20218
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20220 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20219)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20219
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20221 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18943, tmp20220)
}
__typedArg0 := let__18943
__typedArg1 := tmp20220
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20222 := Call(__e, PrimFunc(symis_b), V6260, tmp20221, B6256, L6257, Key6258, C6259)


__e.TailApply(PrimFunc(symshen_4gc), B6256, tmp20222)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20223 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20224 := Call(__e, PrimFunc(symshen_4assoc_1_6), symelement_2, tmp20212, tmp20223)


tmp20225 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20224)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20224
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20225

tmp20226 := MakeNative(func(__e *ControlFlow) {
V6265 := __e.Get(1)
_ = V6265
__e.Return(MakeNative(func(__e *ControlFlow) {
B6261 := __e.Get(1)
_ = B6261
__e.Return(MakeNative(func(__e *ControlFlow) {
L6262 := __e.Get(1)
_ = L6262
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6263 := __e.Get(1)
_ = Key6263
__e.Return(MakeNative(func(__e *ControlFlow) {
C6264 := __e.Get(1)
_ = C6264
tmp20227 := Call(__e, PrimFunc(symshen_4newpv), B6261)


let__18944 := tmp20227
_ = let__18944

tmp20228 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20229 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20228)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20228
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20230 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18944, tmp20229)
}
__typedArg0 := let__18944
__typedArg1 := tmp20229
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20231 := Call(__e, PrimFunc(symis_b), V6265, tmp20230, B6261, L6262, Key6263, C6264)


__e.TailApply(PrimFunc(symshen_4gc), B6261, tmp20231)
return


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

tmp20233 := Call(__e, PrimFunc(symshen_4assoc_1_6), symempty_2, tmp20226, tmp20232)


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
tmp20236 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20237 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20236)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20236
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20238 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp20237)
}
__typedArg0 := symsymbol
__typedArg1 := tmp20237
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6270, tmp20238, B6266, L6267, Key6268, C6269)
return


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

tmp20240 := Call(__e, PrimFunc(symshen_4assoc_1_6), symenable_1type_1theory, tmp20235, tmp20239)


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
V6275 := __e.Get(1)
_ = V6275
__e.Return(MakeNative(func(__e *ControlFlow) {
B6271 := __e.Get(1)
_ = B6271
__e.Return(MakeNative(func(__e *ControlFlow) {
L6272 := __e.Get(1)
_ = L6272
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6273 := __e.Get(1)
_ = Key6273
__e.Return(MakeNative(func(__e *ControlFlow) {
C6274 := __e.Get(1)
_ = C6274
tmp20243 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20244 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20243)
}
__typedArg0 := symlist
__typedArg1 := tmp20243
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20245 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20244, Nil)
}
__typedArg0 := tmp20244
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20246 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20245)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20245
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20247 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp20246)
}
__typedArg0 := symsymbol
__typedArg1 := tmp20246
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6275, tmp20247, B6271, L6272, Key6273, C6274)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20248 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20249 := Call(__e, PrimFunc(symshen_4assoc_1_6), symexternal, tmp20242, tmp20248)


tmp20250 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20249)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20249
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20250

tmp20251 := MakeNative(func(__e *ControlFlow) {
V6280 := __e.Get(1)
_ = V6280
__e.Return(MakeNative(func(__e *ControlFlow) {
B6276 := __e.Get(1)
_ = B6276
__e.Return(MakeNative(func(__e *ControlFlow) {
L6277 := __e.Get(1)
_ = L6277
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6278 := __e.Get(1)
_ = Key6278
__e.Return(MakeNative(func(__e *ControlFlow) {
C6279 := __e.Get(1)
_ = C6279
tmp20252 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
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
return PrimCons(symexception, tmp20253)
}
__typedArg0 := symexception
__typedArg1 := tmp20253
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6280, tmp20254, B6276, L6277, Key6278, C6279)
return


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

tmp20256 := Call(__e, PrimFunc(symshen_4assoc_1_6), symerror_1to_1string, tmp20251, tmp20255)


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
V6285 := __e.Get(1)
_ = V6285
__e.Return(MakeNative(func(__e *ControlFlow) {
B6281 := __e.Get(1)
_ = B6281
__e.Return(MakeNative(func(__e *ControlFlow) {
L6282 := __e.Get(1)
_ = L6282
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6283 := __e.Get(1)
_ = Key6283
__e.Return(MakeNative(func(__e *ControlFlow) {
C6284 := __e.Get(1)
_ = C6284
tmp20259 := Call(__e, PrimFunc(symshen_4newpv), B6281)


let__18945 := tmp20259
_ = let__18945

tmp20260 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20261 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20260)
}
__typedArg0 := symlist
__typedArg1 := tmp20260
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20262 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20261, Nil)
}
__typedArg0 := tmp20261
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20263 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20262)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20262
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20264 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18945, tmp20263)
}
__typedArg0 := let__18945
__typedArg1 := tmp20263
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20265 := Call(__e, PrimFunc(symis_b), V6285, tmp20264, B6281, L6282, Key6283, C6284)


__e.TailApply(PrimFunc(symshen_4gc), B6281, tmp20265)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20266 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20267 := Call(__e, PrimFunc(symshen_4assoc_1_6), symexplode, tmp20258, tmp20266)


tmp20268 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20267)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20267
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20268

tmp20269 := MakeNative(func(__e *ControlFlow) {
V6290 := __e.Get(1)
_ = V6290
__e.Return(MakeNative(func(__e *ControlFlow) {
B6286 := __e.Get(1)
_ = B6286
__e.Return(MakeNative(func(__e *ControlFlow) {
L6287 := __e.Get(1)
_ = L6287
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6288 := __e.Get(1)
_ = Key6288
__e.Return(MakeNative(func(__e *ControlFlow) {
C6289 := __e.Get(1)
_ = C6289
tmp20270 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20271 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20270)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20270
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20272 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp20271)
}
__typedArg0 := symsymbol
__typedArg1 := tmp20271
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6290, tmp20272, B6286, L6287, Key6288, C6289)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20273 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20274 := Call(__e, PrimFunc(symshen_4assoc_1_6), symfactorise, tmp20269, tmp20273)


tmp20275 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20274)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20274
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20275

tmp20276 := MakeNative(func(__e *ControlFlow) {
V6295 := __e.Get(1)
_ = V6295
__e.Return(MakeNative(func(__e *ControlFlow) {
B6291 := __e.Get(1)
_ = B6291
__e.Return(MakeNative(func(__e *ControlFlow) {
L6292 := __e.Get(1)
_ = L6292
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6293 := __e.Get(1)
_ = Key6293
__e.Return(MakeNative(func(__e *ControlFlow) {
C6294 := __e.Get(1)
_ = C6294
tmp20277 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20278 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20277)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20277
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6295, tmp20278, B6291, L6292, Key6293, C6294)
return


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

tmp20280 := Call(__e, PrimFunc(symshen_4assoc_1_6), symfactorise_2, tmp20276, tmp20279)


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
V6300 := __e.Get(1)
_ = V6300
__e.Return(MakeNative(func(__e *ControlFlow) {
B6296 := __e.Get(1)
_ = B6296
__e.Return(MakeNative(func(__e *ControlFlow) {
L6297 := __e.Get(1)
_ = L6297
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6298 := __e.Get(1)
_ = Key6298
__e.Return(MakeNative(func(__e *ControlFlow) {
C6299 := __e.Get(1)
_ = C6299
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

__e.TailApply(PrimFunc(symis_b), V6300, tmp20284, B6296, L6297, Key6298, C6299)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20285 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20286 := Call(__e, PrimFunc(symshen_4assoc_1_6), symfail, tmp20282, tmp20285)


tmp20287 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20286)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20286
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20287

tmp20288 := MakeNative(func(__e *ControlFlow) {
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
tmp20289 := Call(__e, PrimFunc(symshen_4newpv), B6302)


let__18946 := tmp20289
_ = let__18946

tmp20290 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18946, Nil)
}
__typedArg0 := let__18946
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20291 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20290)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20290
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20292 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18946, tmp20291)
}
__typedArg0 := let__18946
__typedArg1 := tmp20291
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20293 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18946, Nil)
}
__typedArg0 := let__18946
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
return PrimCons(let__18946, tmp20294)
}
__typedArg0 := let__18946
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

tmp20299 := Call(__e, PrimFunc(symis_b), V6306, tmp20298, B6302, L6303, Key6304, C6305)


__e.TailApply(PrimFunc(symshen_4gc), B6302, tmp20299)
return


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

tmp20301 := Call(__e, PrimFunc(symshen_4assoc_1_6), symfix, tmp20288, tmp20300)


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
tmp20304 := Call(__e, PrimFunc(symshen_4newpv), B6307)


let__18947 := tmp20304
_ = let__18947

tmp20305 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18947, Nil)
}
__typedArg0 := let__18947
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20306 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlazy, tmp20305)
}
__typedArg0 := symlazy
__typedArg1 := tmp20305
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20307 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20306, Nil)
}
__typedArg0 := tmp20306
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20308 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20307)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20307
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20309 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18947, tmp20308)
}
__typedArg0 := let__18947
__typedArg1 := tmp20308
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20310 := Call(__e, PrimFunc(symis_b), V6311, tmp20309, B6307, L6308, Key6309, C6310)


__e.TailApply(PrimFunc(symshen_4gc), B6307, tmp20310)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20311 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20312 := Call(__e, PrimFunc(symshen_4assoc_1_6), symfreeze, tmp20303, tmp20311)


tmp20313 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20312)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20312
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20313

tmp20314 := MakeNative(func(__e *ControlFlow) {
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
tmp20315 := Call(__e, PrimFunc(symshen_4newpv), B6312)


let__18948 := tmp20315
_ = let__18948

tmp20316 := Call(__e, PrimFunc(symshen_4newpv), B6312)


let__18949 := tmp20316
_ = let__18949

tmp20317 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18948, Nil)
}
__typedArg0 := let__18948
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20318 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_d, tmp20317)
}
__typedArg0 := sym_d
__typedArg1 := tmp20317
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20319 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18949, tmp20318)
}
__typedArg0 := let__18949
__typedArg1 := tmp20318
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20320 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18949, Nil)
}
__typedArg0 := let__18949
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
return PrimCons(tmp20319, tmp20321)
}
__typedArg0 := tmp20319
__typedArg1 := tmp20321
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20323 := Call(__e, PrimFunc(symis_b), V6316, tmp20322, B6312, L6313, Key6314, C6315)


tmp20324 := Call(__e, PrimFunc(symshen_4gc), B6312, tmp20323)


__e.TailApply(PrimFunc(symshen_4gc), B6312, tmp20324)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20325 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20326 := Call(__e, PrimFunc(symshen_4assoc_1_6), symfst, tmp20314, tmp20325)


tmp20327 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20326)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20326
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20327

tmp20328 := MakeNative(func(__e *ControlFlow) {
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
tmp20329 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
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
return PrimCons(symsymbol, tmp20330)
}
__typedArg0 := symsymbol
__typedArg1 := tmp20330
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6321, tmp20331, B6317, L6318, Key6319, C6320)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20332 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20333 := Call(__e, PrimFunc(symshen_4assoc_1_6), symgensym, tmp20328, tmp20332)


tmp20334 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20333)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20333
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20334

tmp20335 := MakeNative(func(__e *ControlFlow) {
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
tmp20336 := Call(__e, PrimFunc(symshen_4newpv), B6323)


let__18950 := tmp20336
_ = let__18950

tmp20337 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18950, Nil)
}
__typedArg0 := let__18950
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20338 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20337)
}
__typedArg0 := symlist
__typedArg1 := tmp20337
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20339 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
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
return PrimCons(let__18950, tmp20340)
}
__typedArg0 := let__18950
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
return PrimCons(tmp20338, tmp20343)
}
__typedArg0 := tmp20338
__typedArg1 := tmp20343
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20345 := Call(__e, PrimFunc(symis_b), V6327, tmp20344, B6323, L6324, Key6325, C6326)


__e.TailApply(PrimFunc(symshen_4gc), B6323, tmp20345)
return


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

tmp20347 := Call(__e, PrimFunc(symshen_4assoc_1_6), symshen_4hds_a_2, tmp20335, tmp20346)


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
tmp20350 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20351 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20350)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20350
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20352 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp20351)
}
__typedArg0 := symsymbol
__typedArg1 := tmp20351
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6332, tmp20352, B6328, L6329, Key6330, C6331)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20353 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20354 := Call(__e, PrimFunc(symshen_4assoc_1_6), symhush, tmp20349, tmp20353)


tmp20355 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20354)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20354
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20355

tmp20356 := MakeNative(func(__e *ControlFlow) {
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
tmp20357 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20358 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20357)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20357
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6337, tmp20358, B6333, L6334, Key6335, C6336)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20359 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20360 := Call(__e, PrimFunc(symshen_4assoc_1_6), symhush_2, tmp20356, tmp20359)


tmp20361 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20360)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20360
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20361

tmp20362 := MakeNative(func(__e *ControlFlow) {
V6343 := __e.Get(1)
_ = V6343
__e.Return(MakeNative(func(__e *ControlFlow) {
B6339 := __e.Get(1)
_ = B6339
__e.Return(MakeNative(func(__e *ControlFlow) {
L6340 := __e.Get(1)
_ = L6340
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6341 := __e.Get(1)
_ = Key6341
__e.Return(MakeNative(func(__e *ControlFlow) {
C6342 := __e.Get(1)
_ = C6342
tmp20363 := Call(__e, PrimFunc(symshen_4newpv), B6339)


let__18951 := tmp20363
_ = let__18951

tmp20364 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18951, Nil)
}
__typedArg0 := let__18951
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20365 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp20364)
}
__typedArg0 := symvector
__typedArg1 := tmp20364
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20366 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18951, Nil)
}
__typedArg0 := let__18951
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20367 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20366)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20366
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20368 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp20367)
}
__typedArg0 := symnumber
__typedArg1 := tmp20367
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20369 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20368, Nil)
}
__typedArg0 := tmp20368
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
return PrimCons(tmp20365, tmp20370)
}
__typedArg0 := tmp20365
__typedArg1 := tmp20370
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20372 := Call(__e, PrimFunc(symis_b), V6343, tmp20371, B6339, L6340, Key6341, C6342)


__e.TailApply(PrimFunc(symshen_4gc), B6339, tmp20372)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20373 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20374 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_5_1vector, tmp20362, tmp20373)


tmp20375 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20374)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20374
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20375

tmp20376 := MakeNative(func(__e *ControlFlow) {
V6349 := __e.Get(1)
_ = V6349
__e.Return(MakeNative(func(__e *ControlFlow) {
B6345 := __e.Get(1)
_ = B6345
__e.Return(MakeNative(func(__e *ControlFlow) {
L6346 := __e.Get(1)
_ = L6346
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6347 := __e.Get(1)
_ = Key6347
__e.Return(MakeNative(func(__e *ControlFlow) {
C6348 := __e.Get(1)
_ = C6348
tmp20377 := Call(__e, PrimFunc(symshen_4newpv), B6345)


let__18952 := tmp20377
_ = let__18952

tmp20378 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18952, Nil)
}
__typedArg0 := let__18952
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20379 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp20378)
}
__typedArg0 := symvector
__typedArg1 := tmp20378
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20380 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18952, Nil)
}
__typedArg0 := let__18952
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20381 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp20380)
}
__typedArg0 := symvector
__typedArg1 := tmp20380
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20382 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20381, Nil)
}
__typedArg0 := tmp20381
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20383 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20382)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20382
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20384 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18952, tmp20383)
}
__typedArg0 := let__18952
__typedArg1 := tmp20383
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20385 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20384, Nil)
}
__typedArg0 := tmp20384
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20386 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20385)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20385
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20387 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp20386)
}
__typedArg0 := symnumber
__typedArg1 := tmp20386
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20388 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20387, Nil)
}
__typedArg0 := tmp20387
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20389 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20388)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20388
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20390 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20379, tmp20389)
}
__typedArg0 := tmp20379
__typedArg1 := tmp20389
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20391 := Call(__e, PrimFunc(symis_b), V6349, tmp20390, B6345, L6346, Key6347, C6348)


__e.TailApply(PrimFunc(symshen_4gc), B6345, tmp20391)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20392 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20393 := Call(__e, PrimFunc(symshen_4assoc_1_6), symvector_1_6, tmp20376, tmp20392)


tmp20394 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20393)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20393
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20394

tmp20395 := MakeNative(func(__e *ControlFlow) {
V6354 := __e.Get(1)
_ = V6354
__e.Return(MakeNative(func(__e *ControlFlow) {
B6350 := __e.Get(1)
_ = B6350
__e.Return(MakeNative(func(__e *ControlFlow) {
L6351 := __e.Get(1)
_ = L6351
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6352 := __e.Get(1)
_ = Key6352
__e.Return(MakeNative(func(__e *ControlFlow) {
C6353 := __e.Get(1)
_ = C6353
tmp20396 := Call(__e, PrimFunc(symshen_4newpv), B6350)


let__18953 := tmp20396
_ = let__18953

tmp20397 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18953, Nil)
}
__typedArg0 := let__18953
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20398 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp20397)
}
__typedArg0 := symvector
__typedArg1 := tmp20397
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20399 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20398, Nil)
}
__typedArg0 := tmp20398
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20400 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20399)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20399
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20401 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp20400)
}
__typedArg0 := symnumber
__typedArg1 := tmp20400
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20402 := Call(__e, PrimFunc(symis_b), V6354, tmp20401, B6350, L6351, Key6352, C6353)


__e.TailApply(PrimFunc(symshen_4gc), B6350, tmp20402)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20403 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20404 := Call(__e, PrimFunc(symshen_4assoc_1_6), symvector, tmp20395, tmp20403)


tmp20405 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20404)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20404
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20405

tmp20406 := MakeNative(func(__e *ControlFlow) {
V6359 := __e.Get(1)
_ = V6359
__e.Return(MakeNative(func(__e *ControlFlow) {
B6355 := __e.Get(1)
_ = B6355
__e.Return(MakeNative(func(__e *ControlFlow) {
L6356 := __e.Get(1)
_ = L6356
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6357 := __e.Get(1)
_ = Key6357
__e.Return(MakeNative(func(__e *ControlFlow) {
C6358 := __e.Get(1)
_ = C6358
tmp20407 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20408 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20407)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20407
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20409 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp20408)
}
__typedArg0 := symsymbol
__typedArg1 := tmp20408
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6359, tmp20409, B6355, L6356, Key6357, C6358)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20410 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20411 := Call(__e, PrimFunc(symshen_4assoc_1_6), symget_1time, tmp20406, tmp20410)


tmp20412 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20411)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20411
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20412

tmp20413 := MakeNative(func(__e *ControlFlow) {
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
tmp20414 := Call(__e, PrimFunc(symshen_4newpv), B6360)


let__18954 := tmp20414
_ = let__18954

tmp20415 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20416 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20415)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20415
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20417 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp20416)
}
__typedArg0 := symnumber
__typedArg1 := tmp20416
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20418 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20417, Nil)
}
__typedArg0 := tmp20417
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20419 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20418)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20418
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20420 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18954, tmp20419)
}
__typedArg0 := let__18954
__typedArg1 := tmp20419
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20421 := Call(__e, PrimFunc(symis_b), V6364, tmp20420, B6360, L6361, Key6362, C6363)


__e.TailApply(PrimFunc(symshen_4gc), B6360, tmp20421)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20422 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20423 := Call(__e, PrimFunc(symshen_4assoc_1_6), symhash, tmp20413, tmp20422)


tmp20424 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20423)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20423
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20424

tmp20425 := MakeNative(func(__e *ControlFlow) {
V6369 := __e.Get(1)
_ = V6369
__e.Return(MakeNative(func(__e *ControlFlow) {
B6365 := __e.Get(1)
_ = B6365
__e.Return(MakeNative(func(__e *ControlFlow) {
L6366 := __e.Get(1)
_ = L6366
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6367 := __e.Get(1)
_ = Key6367
__e.Return(MakeNative(func(__e *ControlFlow) {
C6368 := __e.Get(1)
_ = C6368
tmp20426 := Call(__e, PrimFunc(symshen_4newpv), B6365)


let__18955 := tmp20426
_ = let__18955

tmp20427 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18955, Nil)
}
__typedArg0 := let__18955
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20428 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20427)
}
__typedArg0 := symlist
__typedArg1 := tmp20427
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20429 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18955, Nil)
}
__typedArg0 := let__18955
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20430 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20429)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20429
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20431 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20428, tmp20430)
}
__typedArg0 := tmp20428
__typedArg1 := tmp20430
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20432 := Call(__e, PrimFunc(symis_b), V6369, tmp20431, B6365, L6366, Key6367, C6368)


__e.TailApply(PrimFunc(symshen_4gc), B6365, tmp20432)
return


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

tmp20434 := Call(__e, PrimFunc(symshen_4assoc_1_6), symhead, tmp20425, tmp20433)


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
V6374 := __e.Get(1)
_ = V6374
__e.Return(MakeNative(func(__e *ControlFlow) {
B6370 := __e.Get(1)
_ = B6370
__e.Return(MakeNative(func(__e *ControlFlow) {
L6371 := __e.Get(1)
_ = L6371
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6372 := __e.Get(1)
_ = Key6372
__e.Return(MakeNative(func(__e *ControlFlow) {
C6373 := __e.Get(1)
_ = C6373
tmp20437 := Call(__e, PrimFunc(symshen_4newpv), B6370)


let__18956 := tmp20437
_ = let__18956

tmp20438 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18956, Nil)
}
__typedArg0 := let__18956
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20439 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp20438)
}
__typedArg0 := symvector
__typedArg1 := tmp20438
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20440 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18956, Nil)
}
__typedArg0 := let__18956
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20441 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20440)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20440
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20442 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20439, tmp20441)
}
__typedArg0 := tmp20439
__typedArg1 := tmp20441
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20443 := Call(__e, PrimFunc(symis_b), V6374, tmp20442, B6370, L6371, Key6372, C6373)


__e.TailApply(PrimFunc(symshen_4gc), B6370, tmp20443)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20444 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20445 := Call(__e, PrimFunc(symshen_4assoc_1_6), symhdv, tmp20436, tmp20444)


tmp20446 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20445)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20445
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20446

tmp20447 := MakeNative(func(__e *ControlFlow) {
V6379 := __e.Get(1)
_ = V6379
__e.Return(MakeNative(func(__e *ControlFlow) {
B6375 := __e.Get(1)
_ = B6375
__e.Return(MakeNative(func(__e *ControlFlow) {
L6376 := __e.Get(1)
_ = L6376
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6377 := __e.Get(1)
_ = Key6377
__e.Return(MakeNative(func(__e *ControlFlow) {
C6378 := __e.Get(1)
_ = C6378
tmp20448 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20449 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20448)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20448
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20450 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp20449)
}
__typedArg0 := symstring
__typedArg1 := tmp20449
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6379, tmp20450, B6375, L6376, Key6377, C6378)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20451 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20452 := Call(__e, PrimFunc(symshen_4assoc_1_6), symhdstr, tmp20447, tmp20451)


tmp20453 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20452)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20452
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20453

tmp20454 := MakeNative(func(__e *ControlFlow) {
V6384 := __e.Get(1)
_ = V6384
__e.Return(MakeNative(func(__e *ControlFlow) {
B6380 := __e.Get(1)
_ = B6380
__e.Return(MakeNative(func(__e *ControlFlow) {
L6381 := __e.Get(1)
_ = L6381
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6382 := __e.Get(1)
_ = Key6382
__e.Return(MakeNative(func(__e *ControlFlow) {
C6383 := __e.Get(1)
_ = C6383
tmp20455 := Call(__e, PrimFunc(symshen_4newpv), B6380)


let__18957 := tmp20455
_ = let__18957

tmp20456 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18957, Nil)
}
__typedArg0 := let__18957
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20457 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20456)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20456
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20458 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18957, tmp20457)
}
__typedArg0 := let__18957
__typedArg1 := tmp20457
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20459 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20458, Nil)
}
__typedArg0 := tmp20458
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20460 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20459)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20459
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20461 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18957, tmp20460)
}
__typedArg0 := let__18957
__typedArg1 := tmp20460
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20462 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20461, Nil)
}
__typedArg0 := tmp20461
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20463 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20462)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20462
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20464 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, tmp20463)
}
__typedArg0 := symboolean
__typedArg1 := tmp20463
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20465 := Call(__e, PrimFunc(symis_b), V6384, tmp20464, B6380, L6381, Key6382, C6383)


__e.TailApply(PrimFunc(symshen_4gc), B6380, tmp20465)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20466 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20467 := Call(__e, PrimFunc(symshen_4assoc_1_6), symif, tmp20454, tmp20466)


tmp20468 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20467)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20467
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20468

tmp20469 := MakeNative(func(__e *ControlFlow) {
V6389 := __e.Get(1)
_ = V6389
__e.Return(MakeNative(func(__e *ControlFlow) {
B6385 := __e.Get(1)
_ = B6385
__e.Return(MakeNative(func(__e *ControlFlow) {
L6386 := __e.Get(1)
_ = L6386
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6387 := __e.Get(1)
_ = Key6387
__e.Return(MakeNative(func(__e *ControlFlow) {
C6388 := __e.Get(1)
_ = C6388
tmp20470 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20471 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20470)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20470
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20472 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp20471)
}
__typedArg0 := symsymbol
__typedArg1 := tmp20471
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6389, tmp20472, B6385, L6386, Key6387, C6388)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20473 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20474 := Call(__e, PrimFunc(symshen_4assoc_1_6), symin_1package, tmp20469, tmp20473)


tmp20475 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20474)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20474
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20475

tmp20476 := MakeNative(func(__e *ControlFlow) {
V6394 := __e.Get(1)
_ = V6394
__e.Return(MakeNative(func(__e *ControlFlow) {
B6390 := __e.Get(1)
_ = B6390
__e.Return(MakeNative(func(__e *ControlFlow) {
L6391 := __e.Get(1)
_ = L6391
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6392 := __e.Get(1)
_ = Key6392
__e.Return(MakeNative(func(__e *ControlFlow) {
C6393 := __e.Get(1)
_ = C6393
tmp20477 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20478 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20477)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20477
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6394, tmp20478, B6390, L6391, Key6392, C6393)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20479 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20480 := Call(__e, PrimFunc(symshen_4assoc_1_6), symit, tmp20476, tmp20479)


tmp20481 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20480)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20480
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20481

tmp20482 := MakeNative(func(__e *ControlFlow) {
V6399 := __e.Get(1)
_ = V6399
__e.Return(MakeNative(func(__e *ControlFlow) {
B6395 := __e.Get(1)
_ = B6395
__e.Return(MakeNative(func(__e *ControlFlow) {
L6396 := __e.Get(1)
_ = L6396
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6397 := __e.Get(1)
_ = Key6397
__e.Return(MakeNative(func(__e *ControlFlow) {
C6398 := __e.Get(1)
_ = C6398
tmp20483 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
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

__e.TailApply(PrimFunc(symis_b), V6399, tmp20484, B6395, L6396, Key6397, C6398)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20485 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20486 := Call(__e, PrimFunc(symshen_4assoc_1_6), symimplementation, tmp20482, tmp20485)


tmp20487 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20486)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20486
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20487

tmp20488 := MakeNative(func(__e *ControlFlow) {
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
tmp20489 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20490 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20489)
}
__typedArg0 := symlist
__typedArg1 := tmp20489
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20491 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20492 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20491)
}
__typedArg0 := symlist
__typedArg1 := tmp20491
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20493 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20492, Nil)
}
__typedArg0 := tmp20492
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
return PrimCons(tmp20490, tmp20494)
}
__typedArg0 := tmp20490
__typedArg1 := tmp20494
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6405, tmp20495, B6401, L6402, Key6403, C6404)
return


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

tmp20497 := Call(__e, PrimFunc(symshen_4assoc_1_6), syminclude, tmp20488, tmp20496)


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
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20503 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20502)
}
__typedArg0 := symlist
__typedArg1 := tmp20502
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20504 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20503, Nil)
}
__typedArg0 := tmp20503
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20505 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20504)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20504
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20506 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20501, tmp20505)
}
__typedArg0 := tmp20501
__typedArg1 := tmp20505
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6411, tmp20506, B6407, L6408, Key6409, C6410)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20507 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20508 := Call(__e, PrimFunc(symshen_4assoc_1_6), syminclude_1all_1but, tmp20499, tmp20507)


tmp20509 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20508)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20508
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20509

tmp20510 := MakeNative(func(__e *ControlFlow) {
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
tmp20511 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20512 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20511)
}
__typedArg0 := symlist
__typedArg1 := tmp20511
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20513 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20512, Nil)
}
__typedArg0 := tmp20512
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20514 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20513)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20513
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6416, tmp20514, B6412, L6413, Key6414, C6415)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20515 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20516 := Call(__e, PrimFunc(symshen_4assoc_1_6), symincluded, tmp20510, tmp20515)


tmp20517 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20516)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20516
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20517

tmp20518 := MakeNative(func(__e *ControlFlow) {
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
tmp20519 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
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

__e.TailApply(PrimFunc(symis_b), V6421, tmp20520, B6417, L6418, Key6419, C6420)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20521 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20522 := Call(__e, PrimFunc(symshen_4assoc_1_6), syminferences, tmp20518, tmp20521)


tmp20523 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20522)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20522
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20523

tmp20524 := MakeNative(func(__e *ControlFlow) {
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
tmp20525 := Call(__e, PrimFunc(symshen_4newpv), B6422)


let__18958 := tmp20525
_ = let__18958

tmp20526 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20527 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20526)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20526
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20528 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp20527)
}
__typedArg0 := symstring
__typedArg1 := tmp20527
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20529 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20528, Nil)
}
__typedArg0 := tmp20528
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20530 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20529)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20529
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20531 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18958, tmp20530)
}
__typedArg0 := let__18958
__typedArg1 := tmp20530
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20532 := Call(__e, PrimFunc(symis_b), V6426, tmp20531, B6422, L6423, Key6424, C6425)


__e.TailApply(PrimFunc(symshen_4gc), B6422, tmp20532)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20533 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20534 := Call(__e, PrimFunc(symshen_4assoc_1_6), symshen_4insert, tmp20524, tmp20533)


tmp20535 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20534)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20534
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20535

tmp20536 := MakeNative(func(__e *ControlFlow) {
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
tmp20537 := Call(__e, PrimFunc(symshen_4newpv), B6427)


let__18959 := tmp20537
_ = let__18959

tmp20538 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
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
return PrimCons(let__18959, tmp20539)
}
__typedArg0 := let__18959
__typedArg1 := tmp20539
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20541 := Call(__e, PrimFunc(symis_b), V6431, tmp20540, B6427, L6428, Key6429, C6430)


__e.TailApply(PrimFunc(symshen_4gc), B6427, tmp20541)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20542 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20543 := Call(__e, PrimFunc(symshen_4assoc_1_6), syminteger_2, tmp20536, tmp20542)


tmp20544 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20543)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20543
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20544

tmp20545 := MakeNative(func(__e *ControlFlow) {
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
tmp20546 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
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
return PrimCons(tmp20547, Nil)
}
__typedArg0 := tmp20547
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20549 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20548)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20548
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20550 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp20549)
}
__typedArg0 := symsymbol
__typedArg1 := tmp20549
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6436, tmp20550, B6432, L6433, Key6434, C6435)
return


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

tmp20552 := Call(__e, PrimFunc(symshen_4assoc_1_6), syminternal, tmp20545, tmp20551)


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
V6443 := __e.Get(1)
_ = V6443
__e.Return(MakeNative(func(__e *ControlFlow) {
B6439 := __e.Get(1)
_ = B6439
__e.Return(MakeNative(func(__e *ControlFlow) {
L6440 := __e.Get(1)
_ = L6440
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6441 := __e.Get(1)
_ = Key6441
__e.Return(MakeNative(func(__e *ControlFlow) {
C6442 := __e.Get(1)
_ = C6442
tmp20555 := Call(__e, PrimFunc(symshen_4newpv), B6439)


let__18960 := tmp20555
_ = let__18960

tmp20556 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18960, Nil)
}
__typedArg0 := let__18960
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20557 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20556)
}
__typedArg0 := symlist
__typedArg1 := tmp20556
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20558 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18960, Nil)
}
__typedArg0 := let__18960
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20559 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20558)
}
__typedArg0 := symlist
__typedArg1 := tmp20558
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20560 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18960, Nil)
}
__typedArg0 := let__18960
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20561 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20560)
}
__typedArg0 := symlist
__typedArg1 := tmp20560
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20562 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20561, Nil)
}
__typedArg0 := tmp20561
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20563 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20562)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20562
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20564 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20559, tmp20563)
}
__typedArg0 := tmp20559
__typedArg1 := tmp20563
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20565 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20564, Nil)
}
__typedArg0 := tmp20564
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20566 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20565)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20565
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20567 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20557, tmp20566)
}
__typedArg0 := tmp20557
__typedArg1 := tmp20566
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20568 := Call(__e, PrimFunc(symis_b), V6443, tmp20567, B6439, L6440, Key6441, C6442)


__e.TailApply(PrimFunc(symshen_4gc), B6439, tmp20568)
return


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

tmp20570 := Call(__e, PrimFunc(symshen_4assoc_1_6), symintersection, tmp20554, tmp20569)


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
tmp20573 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20574 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20573)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20573
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6448, tmp20574, B6444, L6445, Key6446, C6447)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20575 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20576 := Call(__e, PrimFunc(symshen_4assoc_1_6), symlanguage, tmp20572, tmp20575)


tmp20577 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20576)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20576
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20577

tmp20578 := MakeNative(func(__e *ControlFlow) {
V6453 := __e.Get(1)
_ = V6453
__e.Return(MakeNative(func(__e *ControlFlow) {
B6449 := __e.Get(1)
_ = B6449
__e.Return(MakeNative(func(__e *ControlFlow) {
L6450 := __e.Get(1)
_ = L6450
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6451 := __e.Get(1)
_ = Key6451
__e.Return(MakeNative(func(__e *ControlFlow) {
C6452 := __e.Get(1)
_ = C6452
tmp20579 := Call(__e, PrimFunc(symshen_4newpv), B6449)


let__18961 := tmp20579
_ = let__18961

tmp20580 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18961, Nil)
}
__typedArg0 := let__18961
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
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
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
return PrimCons(tmp20581, tmp20583)
}
__typedArg0 := tmp20581
__typedArg1 := tmp20583
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20585 := Call(__e, PrimFunc(symis_b), V6453, tmp20584, B6449, L6450, Key6451, C6452)


__e.TailApply(PrimFunc(symshen_4gc), B6449, tmp20585)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20586 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20587 := Call(__e, PrimFunc(symshen_4assoc_1_6), symlength, tmp20578, tmp20586)


tmp20588 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20587)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20587
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20588

tmp20589 := MakeNative(func(__e *ControlFlow) {
V6458 := __e.Get(1)
_ = V6458
__e.Return(MakeNative(func(__e *ControlFlow) {
B6454 := __e.Get(1)
_ = B6454
__e.Return(MakeNative(func(__e *ControlFlow) {
L6455 := __e.Get(1)
_ = L6455
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6456 := __e.Get(1)
_ = Key6456
__e.Return(MakeNative(func(__e *ControlFlow) {
C6457 := __e.Get(1)
_ = C6457
tmp20590 := Call(__e, PrimFunc(symshen_4newpv), B6454)


let__18962 := tmp20590
_ = let__18962

tmp20591 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18962, Nil)
}
__typedArg0 := let__18962
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20592 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp20591)
}
__typedArg0 := symvector
__typedArg1 := tmp20591
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20593 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20594 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20593)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20593
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20595 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20592, tmp20594)
}
__typedArg0 := tmp20592
__typedArg1 := tmp20594
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20596 := Call(__e, PrimFunc(symis_b), V6458, tmp20595, B6454, L6455, Key6456, C6457)


__e.TailApply(PrimFunc(symshen_4gc), B6454, tmp20596)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20597 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20598 := Call(__e, PrimFunc(symshen_4assoc_1_6), symlimit, tmp20589, tmp20597)


tmp20599 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20598)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20598
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20599

tmp20600 := MakeNative(func(__e *ControlFlow) {
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
tmp20601 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symin, Nil)
}
__typedArg0 := symin
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20602 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp20601)
}
__typedArg0 := symstream
__typedArg1 := tmp20601
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20603 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunit, Nil)
}
__typedArg0 := symunit
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

__e.TailApply(PrimFunc(symis_b), V6464, tmp20607, B6460, L6461, Key6462, C6463)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20608 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20609 := Call(__e, PrimFunc(symshen_4assoc_1_6), symlineread, tmp20600, tmp20608)


tmp20610 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20609)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20609
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20610

tmp20611 := MakeNative(func(__e *ControlFlow) {
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
tmp20612 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20613 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20612)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20612
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20614 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp20613)
}
__typedArg0 := symstring
__typedArg1 := tmp20613
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6469, tmp20614, B6465, L6466, Key6467, C6468)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20615 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20616 := Call(__e, PrimFunc(symshen_4assoc_1_6), symload, tmp20611, tmp20615)


tmp20617 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20616)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20616
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20617

tmp20618 := MakeNative(func(__e *ControlFlow) {
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
tmp20619 := Call(__e, PrimFunc(symshen_4newpv), B6472)


let__18963 := tmp20619
_ = let__18963

tmp20620 := Call(__e, PrimFunc(symshen_4newpv), B6472)


let__18964 := tmp20620
_ = let__18964

tmp20621 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18964, Nil)
}
__typedArg0 := let__18964
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
return PrimCons(let__18963, tmp20622)
}
__typedArg0 := let__18963
__typedArg1 := tmp20622
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20624 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18963, Nil)
}
__typedArg0 := let__18963
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20625 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20624)
}
__typedArg0 := symlist
__typedArg1 := tmp20624
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20626 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18964, Nil)
}
__typedArg0 := let__18964
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

tmp20631 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20630, Nil)
}
__typedArg0 := tmp20630
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
return PrimCons(tmp20623, tmp20632)
}
__typedArg0 := tmp20623
__typedArg1 := tmp20632
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20634 := Call(__e, PrimFunc(symis_b), V6476, tmp20633, B6472, L6473, Key6474, C6475)


tmp20635 := Call(__e, PrimFunc(symshen_4gc), B6472, tmp20634)


__e.TailApply(PrimFunc(symshen_4gc), B6472, tmp20635)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20636 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20637 := Call(__e, PrimFunc(symshen_4assoc_1_6), symmap, tmp20618, tmp20636)


tmp20638 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20637)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20637
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20638

tmp20639 := MakeNative(func(__e *ControlFlow) {
V6483 := __e.Get(1)
_ = V6483
__e.Return(MakeNative(func(__e *ControlFlow) {
B6479 := __e.Get(1)
_ = B6479
__e.Return(MakeNative(func(__e *ControlFlow) {
L6480 := __e.Get(1)
_ = L6480
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6481 := __e.Get(1)
_ = Key6481
__e.Return(MakeNative(func(__e *ControlFlow) {
C6482 := __e.Get(1)
_ = C6482
tmp20640 := Call(__e, PrimFunc(symshen_4newpv), B6479)


let__18965 := tmp20640
_ = let__18965

tmp20641 := Call(__e, PrimFunc(symshen_4newpv), B6479)


let__18966 := tmp20641
_ = let__18966

tmp20642 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18966, Nil)
}
__typedArg0 := let__18966
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20643 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20642)
}
__typedArg0 := symlist
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
return PrimCons(let__18965, tmp20645)
}
__typedArg0 := let__18965
__typedArg1 := tmp20645
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20647 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18965, Nil)
}
__typedArg0 := let__18965
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
return PrimCons(let__18966, Nil)
}
__typedArg0 := let__18966
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

tmp20657 := Call(__e, PrimFunc(symis_b), V6483, tmp20656, B6479, L6480, Key6481, C6482)


tmp20658 := Call(__e, PrimFunc(symshen_4gc), B6479, tmp20657)


__e.TailApply(PrimFunc(symshen_4gc), B6479, tmp20658)
return


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

tmp20660 := Call(__e, PrimFunc(symshen_4assoc_1_6), symmapcan, tmp20639, tmp20659)


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
V6488 := __e.Get(1)
_ = V6488
__e.Return(MakeNative(func(__e *ControlFlow) {
B6484 := __e.Get(1)
_ = B6484
__e.Return(MakeNative(func(__e *ControlFlow) {
L6485 := __e.Get(1)
_ = L6485
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6486 := __e.Get(1)
_ = Key6486
__e.Return(MakeNative(func(__e *ControlFlow) {
C6487 := __e.Get(1)
_ = C6487
tmp20663 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20664 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20663)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20663
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20665 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp20664)
}
__typedArg0 := symnumber
__typedArg1 := tmp20664
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6488, tmp20665, B6484, L6485, Key6486, C6487)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20666 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20667 := Call(__e, PrimFunc(symshen_4assoc_1_6), symmaxinferences, tmp20662, tmp20666)


tmp20668 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20667)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20667
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20668

tmp20669 := MakeNative(func(__e *ControlFlow) {
V6493 := __e.Get(1)
_ = V6493
__e.Return(MakeNative(func(__e *ControlFlow) {
B6489 := __e.Get(1)
_ = B6489
__e.Return(MakeNative(func(__e *ControlFlow) {
L6490 := __e.Get(1)
_ = L6490
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6491 := __e.Get(1)
_ = Key6491
__e.Return(MakeNative(func(__e *ControlFlow) {
C6492 := __e.Get(1)
_ = C6492
tmp20670 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
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
return PrimCons(symnumber, tmp20671)
}
__typedArg0 := symnumber
__typedArg1 := tmp20671
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6493, tmp20672, B6489, L6490, Key6491, C6492)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20673 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20674 := Call(__e, PrimFunc(symshen_4assoc_1_6), symn_1_6string, tmp20669, tmp20673)


tmp20675 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20674)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20674
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20675

tmp20676 := MakeNative(func(__e *ControlFlow) {
V6498 := __e.Get(1)
_ = V6498
__e.Return(MakeNative(func(__e *ControlFlow) {
B6494 := __e.Get(1)
_ = B6494
__e.Return(MakeNative(func(__e *ControlFlow) {
L6495 := __e.Get(1)
_ = L6495
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6496 := __e.Get(1)
_ = Key6496
__e.Return(MakeNative(func(__e *ControlFlow) {
C6497 := __e.Get(1)
_ = C6497
tmp20677 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
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
return PrimCons(symnumber, tmp20678)
}
__typedArg0 := symnumber
__typedArg1 := tmp20678
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6498, tmp20679, B6494, L6495, Key6496, C6497)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20680 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20681 := Call(__e, PrimFunc(symshen_4assoc_1_6), symnl, tmp20676, tmp20680)


tmp20682 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20681)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20681
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20682

tmp20683 := MakeNative(func(__e *ControlFlow) {
V6503 := __e.Get(1)
_ = V6503
__e.Return(MakeNative(func(__e *ControlFlow) {
B6499 := __e.Get(1)
_ = B6499
__e.Return(MakeNative(func(__e *ControlFlow) {
L6500 := __e.Get(1)
_ = L6500
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6501 := __e.Get(1)
_ = Key6501
__e.Return(MakeNative(func(__e *ControlFlow) {
C6502 := __e.Get(1)
_ = C6502
tmp20684 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
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
return PrimCons(symboolean, tmp20685)
}
__typedArg0 := symboolean
__typedArg1 := tmp20685
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6503, tmp20686, B6499, L6500, Key6501, C6502)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20687 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20688 := Call(__e, PrimFunc(symshen_4assoc_1_6), symnot, tmp20683, tmp20687)


tmp20689 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20688)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20688
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20689

tmp20690 := MakeNative(func(__e *ControlFlow) {
V6508 := __e.Get(1)
_ = V6508
__e.Return(MakeNative(func(__e *ControlFlow) {
B6504 := __e.Get(1)
_ = B6504
__e.Return(MakeNative(func(__e *ControlFlow) {
L6505 := __e.Get(1)
_ = L6505
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6506 := __e.Get(1)
_ = Key6506
__e.Return(MakeNative(func(__e *ControlFlow) {
C6507 := __e.Get(1)
_ = C6507
tmp20691 := Call(__e, PrimFunc(symshen_4newpv), B6504)


let__18967 := tmp20691
_ = let__18967

tmp20692 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18967, Nil)
}
__typedArg0 := let__18967
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20693 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20692)
}
__typedArg0 := symlist
__typedArg1 := tmp20692
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20694 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18967, Nil)
}
__typedArg0 := let__18967
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20695 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20694)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20694
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20696 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20693, tmp20695)
}
__typedArg0 := tmp20693
__typedArg1 := tmp20695
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20697 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20696, Nil)
}
__typedArg0 := tmp20696
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20698 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20697)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20697
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20699 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp20698)
}
__typedArg0 := symnumber
__typedArg1 := tmp20698
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20700 := Call(__e, PrimFunc(symis_b), V6508, tmp20699, B6504, L6505, Key6506, C6507)


__e.TailApply(PrimFunc(symshen_4gc), B6504, tmp20700)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20701 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20702 := Call(__e, PrimFunc(symshen_4assoc_1_6), symnth, tmp20690, tmp20701)


tmp20703 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20702)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20702
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20703

tmp20704 := MakeNative(func(__e *ControlFlow) {
V6513 := __e.Get(1)
_ = V6513
__e.Return(MakeNative(func(__e *ControlFlow) {
B6509 := __e.Get(1)
_ = B6509
__e.Return(MakeNative(func(__e *ControlFlow) {
L6510 := __e.Get(1)
_ = L6510
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6511 := __e.Get(1)
_ = Key6511
__e.Return(MakeNative(func(__e *ControlFlow) {
C6512 := __e.Get(1)
_ = C6512
tmp20705 := Call(__e, PrimFunc(symshen_4newpv), B6509)


let__18968 := tmp20705
_ = let__18968

tmp20706 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
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
return PrimCons(let__18968, tmp20707)
}
__typedArg0 := let__18968
__typedArg1 := tmp20707
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20709 := Call(__e, PrimFunc(symis_b), V6513, tmp20708, B6509, L6510, Key6511, C6512)


__e.TailApply(PrimFunc(symshen_4gc), B6509, tmp20709)
return


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

tmp20711 := Call(__e, PrimFunc(symshen_4assoc_1_6), symnumber_2, tmp20704, tmp20710)


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
V6518 := __e.Get(1)
_ = V6518
__e.Return(MakeNative(func(__e *ControlFlow) {
B6514 := __e.Get(1)
_ = B6514
__e.Return(MakeNative(func(__e *ControlFlow) {
L6515 := __e.Get(1)
_ = L6515
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6516 := __e.Get(1)
_ = Key6516
__e.Return(MakeNative(func(__e *ControlFlow) {
C6517 := __e.Get(1)
_ = C6517
tmp20714 := Call(__e, PrimFunc(symshen_4newpv), B6514)


let__18969 := tmp20714
_ = let__18969

tmp20715 := Call(__e, PrimFunc(symshen_4newpv), B6514)


let__18970 := tmp20715
_ = let__18970

tmp20716 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20717 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20716)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20716
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20718 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18970, tmp20717)
}
__typedArg0 := let__18970
__typedArg1 := tmp20717
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20719 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20718, Nil)
}
__typedArg0 := tmp20718
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20720 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20719)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20719
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20721 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18969, tmp20720)
}
__typedArg0 := let__18969
__typedArg1 := tmp20720
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20722 := Call(__e, PrimFunc(symis_b), V6518, tmp20721, B6514, L6515, Key6516, C6517)


tmp20723 := Call(__e, PrimFunc(symshen_4gc), B6514, tmp20722)


__e.TailApply(PrimFunc(symshen_4gc), B6514, tmp20723)
return


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

tmp20725 := Call(__e, PrimFunc(symshen_4assoc_1_6), symoccurrences, tmp20713, tmp20724)


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
V6523 := __e.Get(1)
_ = V6523
__e.Return(MakeNative(func(__e *ControlFlow) {
B6519 := __e.Get(1)
_ = B6519
__e.Return(MakeNative(func(__e *ControlFlow) {
L6520 := __e.Get(1)
_ = L6520
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6521 := __e.Get(1)
_ = Key6521
__e.Return(MakeNative(func(__e *ControlFlow) {
C6522 := __e.Get(1)
_ = C6522
tmp20728 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20729 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20728)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20728
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20730 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp20729)
}
__typedArg0 := symsymbol
__typedArg1 := tmp20729
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6523, tmp20730, B6519, L6520, Key6521, C6522)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20731 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20732 := Call(__e, PrimFunc(symshen_4assoc_1_6), symoccurs_1check, tmp20727, tmp20731)


tmp20733 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20732)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20732
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20733

tmp20734 := MakeNative(func(__e *ControlFlow) {
V6528 := __e.Get(1)
_ = V6528
__e.Return(MakeNative(func(__e *ControlFlow) {
B6524 := __e.Get(1)
_ = B6524
__e.Return(MakeNative(func(__e *ControlFlow) {
L6525 := __e.Get(1)
_ = L6525
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6526 := __e.Get(1)
_ = Key6526
__e.Return(MakeNative(func(__e *ControlFlow) {
C6527 := __e.Get(1)
_ = C6527
tmp20735 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20736 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20735)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20735
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6528, tmp20736, B6524, L6525, Key6526, C6527)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20737 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20738 := Call(__e, PrimFunc(symshen_4assoc_1_6), symoccurs_2, tmp20734, tmp20737)


tmp20739 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20738)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20738
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20739

tmp20740 := MakeNative(func(__e *ControlFlow) {
V6533 := __e.Get(1)
_ = V6533
__e.Return(MakeNative(func(__e *ControlFlow) {
B6529 := __e.Get(1)
_ = B6529
__e.Return(MakeNative(func(__e *ControlFlow) {
L6530 := __e.Get(1)
_ = L6530
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6531 := __e.Get(1)
_ = Key6531
__e.Return(MakeNative(func(__e *ControlFlow) {
C6532 := __e.Get(1)
_ = C6532
tmp20741 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20742 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20741)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20741
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20743 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp20742)
}
__typedArg0 := symsymbol
__typedArg1 := tmp20742
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6533, tmp20743, B6529, L6530, Key6531, C6532)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20744 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20745 := Call(__e, PrimFunc(symshen_4assoc_1_6), symoptimise, tmp20740, tmp20744)


tmp20746 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20745)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20745
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20746

tmp20747 := MakeNative(func(__e *ControlFlow) {
V6538 := __e.Get(1)
_ = V6538
__e.Return(MakeNative(func(__e *ControlFlow) {
B6534 := __e.Get(1)
_ = B6534
__e.Return(MakeNative(func(__e *ControlFlow) {
L6535 := __e.Get(1)
_ = L6535
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6536 := __e.Get(1)
_ = Key6536
__e.Return(MakeNative(func(__e *ControlFlow) {
C6537 := __e.Get(1)
_ = C6537
tmp20748 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20749 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20748)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20748
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6538, tmp20749, B6534, L6535, Key6536, C6537)
return


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

tmp20751 := Call(__e, PrimFunc(symshen_4assoc_1_6), symoptimise_2, tmp20747, tmp20750)


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
V6543 := __e.Get(1)
_ = V6543
__e.Return(MakeNative(func(__e *ControlFlow) {
B6539 := __e.Get(1)
_ = B6539
__e.Return(MakeNative(func(__e *ControlFlow) {
L6540 := __e.Get(1)
_ = L6540
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6541 := __e.Get(1)
_ = Key6541
__e.Return(MakeNative(func(__e *ControlFlow) {
C6542 := __e.Get(1)
_ = C6542
tmp20754 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
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

tmp20756 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, tmp20755)
}
__typedArg0 := symboolean
__typedArg1 := tmp20755
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20757 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20756, Nil)
}
__typedArg0 := tmp20756
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20758 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20757)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20757
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20759 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, tmp20758)
}
__typedArg0 := symboolean
__typedArg1 := tmp20758
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6543, tmp20759, B6539, L6540, Key6541, C6542)
return


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

tmp20761 := Call(__e, PrimFunc(symshen_4assoc_1_6), symor, tmp20753, tmp20760)


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
V6548 := __e.Get(1)
_ = V6548
__e.Return(MakeNative(func(__e *ControlFlow) {
B6544 := __e.Get(1)
_ = B6544
__e.Return(MakeNative(func(__e *ControlFlow) {
L6545 := __e.Get(1)
_ = L6545
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6546 := __e.Get(1)
_ = Key6546
__e.Return(MakeNative(func(__e *ControlFlow) {
C6547 := __e.Get(1)
_ = C6547
tmp20764 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
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

__e.TailApply(PrimFunc(symis_b), V6548, tmp20765, B6544, L6545, Key6546, C6547)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20766 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20767 := Call(__e, PrimFunc(symshen_4assoc_1_6), symos, tmp20763, tmp20766)


tmp20768 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20767)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20767
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20768

tmp20769 := MakeNative(func(__e *ControlFlow) {
V6553 := __e.Get(1)
_ = V6553
__e.Return(MakeNative(func(__e *ControlFlow) {
B6549 := __e.Get(1)
_ = B6549
__e.Return(MakeNative(func(__e *ControlFlow) {
L6550 := __e.Get(1)
_ = L6550
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6551 := __e.Get(1)
_ = Key6551
__e.Return(MakeNative(func(__e *ControlFlow) {
C6552 := __e.Get(1)
_ = C6552
tmp20770 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20771 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20770)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20770
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20772 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp20771)
}
__typedArg0 := symsymbol
__typedArg1 := tmp20771
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6553, tmp20772, B6549, L6550, Key6551, C6552)
return


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

tmp20774 := Call(__e, PrimFunc(symshen_4assoc_1_6), sympackage_2, tmp20769, tmp20773)


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
V6558 := __e.Get(1)
_ = V6558
__e.Return(MakeNative(func(__e *ControlFlow) {
B6554 := __e.Get(1)
_ = B6554
__e.Return(MakeNative(func(__e *ControlFlow) {
L6555 := __e.Get(1)
_ = L6555
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6556 := __e.Get(1)
_ = Key6556
__e.Return(MakeNative(func(__e *ControlFlow) {
C6557 := __e.Get(1)
_ = C6557
tmp20777 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
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

__e.TailApply(PrimFunc(symis_b), V6558, tmp20778, B6554, L6555, Key6556, C6557)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20779 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20780 := Call(__e, PrimFunc(symshen_4assoc_1_6), symport, tmp20776, tmp20779)


tmp20781 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20780)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20780
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20781

tmp20782 := MakeNative(func(__e *ControlFlow) {
V6563 := __e.Get(1)
_ = V6563
__e.Return(MakeNative(func(__e *ControlFlow) {
B6559 := __e.Get(1)
_ = B6559
__e.Return(MakeNative(func(__e *ControlFlow) {
L6560 := __e.Get(1)
_ = L6560
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6561 := __e.Get(1)
_ = Key6561
__e.Return(MakeNative(func(__e *ControlFlow) {
C6562 := __e.Get(1)
_ = C6562
tmp20783 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20784 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20783)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20783
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6563, tmp20784, B6559, L6560, Key6561, C6562)
return


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

tmp20786 := Call(__e, PrimFunc(symshen_4assoc_1_6), symporters, tmp20782, tmp20785)


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
V6568 := __e.Get(1)
_ = V6568
__e.Return(MakeNative(func(__e *ControlFlow) {
B6564 := __e.Get(1)
_ = B6564
__e.Return(MakeNative(func(__e *ControlFlow) {
L6565 := __e.Get(1)
_ = L6565
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6566 := __e.Get(1)
_ = Key6566
__e.Return(MakeNative(func(__e *ControlFlow) {
C6567 := __e.Get(1)
_ = C6567
tmp20789 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20790 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20789)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20789
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20791 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp20790)
}
__typedArg0 := symnumber
__typedArg1 := tmp20790
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20792 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20791, Nil)
}
__typedArg0 := tmp20791
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

__e.TailApply(PrimFunc(symis_b), V6568, tmp20794, B6564, L6565, Key6566, C6567)
return


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

tmp20796 := Call(__e, PrimFunc(symshen_4assoc_1_6), sympos, tmp20788, tmp20795)


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
V6573 := __e.Get(1)
_ = V6573
__e.Return(MakeNative(func(__e *ControlFlow) {
B6569 := __e.Get(1)
_ = B6569
__e.Return(MakeNative(func(__e *ControlFlow) {
L6570 := __e.Get(1)
_ = L6570
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6571 := __e.Get(1)
_ = Key6571
__e.Return(MakeNative(func(__e *ControlFlow) {
C6572 := __e.Get(1)
_ = C6572
tmp20799 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symout, Nil)
}
__typedArg0 := symout
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20800 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp20799)
}
__typedArg0 := symstream
__typedArg1 := tmp20799
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20801 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20802 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20801)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20801
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20803 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20800, tmp20802)
}
__typedArg0 := tmp20800
__typedArg1 := tmp20802
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20804 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20803, Nil)
}
__typedArg0 := tmp20803
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20805 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20804)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20804
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20806 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp20805)
}
__typedArg0 := symstring
__typedArg1 := tmp20805
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6573, tmp20806, B6569, L6570, Key6571, C6572)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20807 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20808 := Call(__e, PrimFunc(symshen_4assoc_1_6), sympr, tmp20798, tmp20807)


tmp20809 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20808)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20808
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20809

tmp20810 := MakeNative(func(__e *ControlFlow) {
V6578 := __e.Get(1)
_ = V6578
__e.Return(MakeNative(func(__e *ControlFlow) {
B6574 := __e.Get(1)
_ = B6574
__e.Return(MakeNative(func(__e *ControlFlow) {
L6575 := __e.Get(1)
_ = L6575
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6576 := __e.Get(1)
_ = Key6576
__e.Return(MakeNative(func(__e *ControlFlow) {
C6577 := __e.Get(1)
_ = C6577
tmp20811 := Call(__e, PrimFunc(symshen_4newpv), B6574)


let__18971 := tmp20811
_ = let__18971

tmp20812 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18971, Nil)
}
__typedArg0 := let__18971
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
return PrimCons(let__18971, tmp20813)
}
__typedArg0 := let__18971
__typedArg1 := tmp20813
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20815 := Call(__e, PrimFunc(symis_b), V6578, tmp20814, B6574, L6575, Key6576, C6577)


__e.TailApply(PrimFunc(symshen_4gc), B6574, tmp20815)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20816 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20817 := Call(__e, PrimFunc(symshen_4assoc_1_6), symprint, tmp20810, tmp20816)


tmp20818 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20817)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20817
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20818

tmp20819 := MakeNative(func(__e *ControlFlow) {
V6583 := __e.Get(1)
_ = V6583
__e.Return(MakeNative(func(__e *ControlFlow) {
B6579 := __e.Get(1)
_ = B6579
__e.Return(MakeNative(func(__e *ControlFlow) {
L6580 := __e.Get(1)
_ = L6580
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6581 := __e.Get(1)
_ = Key6581
__e.Return(MakeNative(func(__e *ControlFlow) {
C6582 := __e.Get(1)
_ = C6582
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
return PrimCons(sym_1_1_6, tmp20820)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20820
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20822 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp20821)
}
__typedArg0 := symsymbol
__typedArg1 := tmp20821
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6583, tmp20822, B6579, L6580, Key6581, C6582)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20823 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20824 := Call(__e, PrimFunc(symshen_4assoc_1_6), symprofile, tmp20819, tmp20823)


tmp20825 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20824)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20824
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20825

tmp20826 := MakeNative(func(__e *ControlFlow) {
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
tmp20827 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20828 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20827)
}
__typedArg0 := symlist
__typedArg1 := tmp20827
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20829 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20830 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20829)
}
__typedArg0 := symlist
__typedArg1 := tmp20829
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20831 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20830, Nil)
}
__typedArg0 := tmp20830
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
return PrimCons(tmp20828, tmp20832)
}
__typedArg0 := tmp20828
__typedArg1 := tmp20832
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6589, tmp20833, B6585, L6586, Key6587, C6588)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20834 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20835 := Call(__e, PrimFunc(symshen_4assoc_1_6), sympreclude, tmp20826, tmp20834)


tmp20836 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20835)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20835
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20836

tmp20837 := MakeNative(func(__e *ControlFlow) {
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
tmp20838 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
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
return PrimCons(symstring, tmp20839)
}
__typedArg0 := symstring
__typedArg1 := tmp20839
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6594, tmp20840, B6590, L6591, Key6592, C6593)
return


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

tmp20842 := Call(__e, PrimFunc(symshen_4assoc_1_6), symshen_4proc_1nl, tmp20837, tmp20841)


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
tmp20845 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20846 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_d, tmp20845)
}
__typedArg0 := sym_d
__typedArg1 := tmp20845
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20847 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp20846)
}
__typedArg0 := symsymbol
__typedArg1 := tmp20846
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20848 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20847, Nil)
}
__typedArg0 := tmp20847
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20849 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20848)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20848
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20850 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp20849)
}
__typedArg0 := symsymbol
__typedArg1 := tmp20849
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6599, tmp20850, B6595, L6596, Key6597, C6598)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20851 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20852 := Call(__e, PrimFunc(symshen_4assoc_1_6), symprofile_1results, tmp20844, tmp20851)


tmp20853 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20852)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20852
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20853

tmp20854 := MakeNative(func(__e *ControlFlow) {
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
tmp20855 := Call(__e, PrimFunc(symshen_4newpv), B6600)


let__18972 := tmp20855
_ = let__18972

tmp20856 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18972, Nil)
}
__typedArg0 := let__18972
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20857 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20856)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20856
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20858 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18972, tmp20857)
}
__typedArg0 := let__18972
__typedArg1 := tmp20857
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20859 := Call(__e, PrimFunc(symis_b), V6604, tmp20858, B6600, L6601, Key6602, C6603)


__e.TailApply(PrimFunc(symshen_4gc), B6600, tmp20859)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20860 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20861 := Call(__e, PrimFunc(symshen_4assoc_1_6), symprotect, tmp20854, tmp20860)


tmp20862 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20861)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20861
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20862

tmp20863 := MakeNative(func(__e *ControlFlow) {
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
tmp20864 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20865 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20864)
}
__typedArg0 := symlist
__typedArg1 := tmp20864
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20866 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20867 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20866)
}
__typedArg0 := symlist
__typedArg1 := tmp20866
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20868 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20867, Nil)
}
__typedArg0 := tmp20867
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
return PrimCons(tmp20865, tmp20869)
}
__typedArg0 := tmp20865
__typedArg1 := tmp20869
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6610, tmp20870, B6606, L6607, Key6608, C6609)
return


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

tmp20872 := Call(__e, PrimFunc(symshen_4assoc_1_6), sympreclude_1all_1but, tmp20863, tmp20871)


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
tmp20875 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symout, Nil)
}
__typedArg0 := symout
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20876 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp20875)
}
__typedArg0 := symstream
__typedArg1 := tmp20875
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20877 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
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
return PrimCons(tmp20876, tmp20878)
}
__typedArg0 := tmp20876
__typedArg1 := tmp20878
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20880 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20879, Nil)
}
__typedArg0 := tmp20879
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20881 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20880)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20880
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20882 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp20881)
}
__typedArg0 := symstring
__typedArg1 := tmp20881
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6615, tmp20882, B6611, L6612, Key6613, C6614)
return


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

tmp20884 := Call(__e, PrimFunc(symshen_4assoc_1_6), symshen_4prhush, tmp20874, tmp20883)


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
tmp20887 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20888 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20887)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20887
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20889 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp20888)
}
__typedArg0 := symnumber
__typedArg1 := tmp20888
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6620, tmp20889, B6616, L6617, Key6618, C6619)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20890 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20891 := Call(__e, PrimFunc(symshen_4assoc_1_6), symprolog_1memory, tmp20886, tmp20890)


tmp20892 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20891)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20891
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20892

tmp20893 := MakeNative(func(__e *ControlFlow) {
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
tmp20894 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunit, Nil)
}
__typedArg0 := symunit
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20895 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20894)
}
__typedArg0 := symlist
__typedArg1 := tmp20894
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20896 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20895, Nil)
}
__typedArg0 := tmp20895
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20897 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20896)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20896
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20898 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp20897)
}
__typedArg0 := symsymbol
__typedArg1 := tmp20897
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6625, tmp20898, B6621, L6622, Key6623, C6624)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20899 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20900 := Call(__e, PrimFunc(symshen_4assoc_1_6), symps, tmp20893, tmp20899)


tmp20901 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20900)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20900
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20901

tmp20902 := MakeNative(func(__e *ControlFlow) {
V6630 := __e.Get(1)
_ = V6630
__e.Return(MakeNative(func(__e *ControlFlow) {
B6626 := __e.Get(1)
_ = B6626
__e.Return(MakeNative(func(__e *ControlFlow) {
L6627 := __e.Get(1)
_ = L6627
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6628 := __e.Get(1)
_ = Key6628
__e.Return(MakeNative(func(__e *ControlFlow) {
C6629 := __e.Get(1)
_ = C6629
tmp20903 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symin, Nil)
}
__typedArg0 := symin
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20904 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp20903)
}
__typedArg0 := symstream
__typedArg1 := tmp20903
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20905 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunit, Nil)
}
__typedArg0 := symunit
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20906 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20905)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20905
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20907 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20904, tmp20906)
}
__typedArg0 := tmp20904
__typedArg1 := tmp20906
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6630, tmp20907, B6626, L6627, Key6628, C6629)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20908 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20909 := Call(__e, PrimFunc(symshen_4assoc_1_6), symread, tmp20902, tmp20908)


tmp20910 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20909)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20909
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20910

tmp20911 := MakeNative(func(__e *ControlFlow) {
V6635 := __e.Get(1)
_ = V6635
__e.Return(MakeNative(func(__e *ControlFlow) {
B6631 := __e.Get(1)
_ = B6631
__e.Return(MakeNative(func(__e *ControlFlow) {
L6632 := __e.Get(1)
_ = L6632
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6633 := __e.Get(1)
_ = Key6633
__e.Return(MakeNative(func(__e *ControlFlow) {
C6634 := __e.Get(1)
_ = C6634
tmp20912 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symin, Nil)
}
__typedArg0 := symin
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20913 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp20912)
}
__typedArg0 := symstream
__typedArg1 := tmp20912
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20914 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20915 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20914)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20914
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20916 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20913, tmp20915)
}
__typedArg0 := tmp20913
__typedArg1 := tmp20915
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6635, tmp20916, B6631, L6632, Key6633, C6634)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20917 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20918 := Call(__e, PrimFunc(symshen_4assoc_1_6), symread_1unit_1string, tmp20911, tmp20917)


tmp20919 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20918)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20918
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20919

tmp20920 := MakeNative(func(__e *ControlFlow) {
V6640 := __e.Get(1)
_ = V6640
__e.Return(MakeNative(func(__e *ControlFlow) {
B6636 := __e.Get(1)
_ = B6636
__e.Return(MakeNative(func(__e *ControlFlow) {
L6637 := __e.Get(1)
_ = L6637
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6638 := __e.Get(1)
_ = Key6638
__e.Return(MakeNative(func(__e *ControlFlow) {
C6639 := __e.Get(1)
_ = C6639
tmp20921 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symin, Nil)
}
__typedArg0 := symin
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20922 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp20921)
}
__typedArg0 := symstream
__typedArg1 := tmp20921
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20923 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20924 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20923)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20923
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20925 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20922, tmp20924)
}
__typedArg0 := tmp20922
__typedArg1 := tmp20924
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6640, tmp20925, B6636, L6637, Key6638, C6639)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20926 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20927 := Call(__e, PrimFunc(symshen_4assoc_1_6), symread_1byte, tmp20920, tmp20926)


tmp20928 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20927)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20927
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20928

tmp20929 := MakeNative(func(__e *ControlFlow) {
V6645 := __e.Get(1)
_ = V6645
__e.Return(MakeNative(func(__e *ControlFlow) {
B6641 := __e.Get(1)
_ = B6641
__e.Return(MakeNative(func(__e *ControlFlow) {
L6642 := __e.Get(1)
_ = L6642
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6643 := __e.Get(1)
_ = Key6643
__e.Return(MakeNative(func(__e *ControlFlow) {
C6644 := __e.Get(1)
_ = C6644
tmp20930 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20931 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20930)
}
__typedArg0 := symlist
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
return PrimCons(symstring, tmp20933)
}
__typedArg0 := symstring
__typedArg1 := tmp20933
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6645, tmp20934, B6641, L6642, Key6643, C6644)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20935 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20936 := Call(__e, PrimFunc(symshen_4assoc_1_6), symread_1file_1as_1bytelist, tmp20929, tmp20935)


tmp20937 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20936)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20936
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20937

tmp20938 := MakeNative(func(__e *ControlFlow) {
V6650 := __e.Get(1)
_ = V6650
__e.Return(MakeNative(func(__e *ControlFlow) {
B6646 := __e.Get(1)
_ = B6646
__e.Return(MakeNative(func(__e *ControlFlow) {
L6647 := __e.Get(1)
_ = L6647
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6648 := __e.Get(1)
_ = Key6648
__e.Return(MakeNative(func(__e *ControlFlow) {
C6649 := __e.Get(1)
_ = C6649
tmp20939 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20940 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20939)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20939
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20941 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp20940)
}
__typedArg0 := symstring
__typedArg1 := tmp20940
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6650, tmp20941, B6646, L6647, Key6648, C6649)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20942 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20943 := Call(__e, PrimFunc(symshen_4assoc_1_6), symread_1file_1as_1string, tmp20938, tmp20942)


tmp20944 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20943)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20943
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20944

tmp20945 := MakeNative(func(__e *ControlFlow) {
V6655 := __e.Get(1)
_ = V6655
__e.Return(MakeNative(func(__e *ControlFlow) {
B6651 := __e.Get(1)
_ = B6651
__e.Return(MakeNative(func(__e *ControlFlow) {
L6652 := __e.Get(1)
_ = L6652
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6653 := __e.Get(1)
_ = Key6653
__e.Return(MakeNative(func(__e *ControlFlow) {
C6654 := __e.Get(1)
_ = C6654
tmp20946 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunit, Nil)
}
__typedArg0 := symunit
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20947 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20946)
}
__typedArg0 := symlist
__typedArg1 := tmp20946
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20948 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20947, Nil)
}
__typedArg0 := tmp20947
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20949 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20948)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20948
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20950 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp20949)
}
__typedArg0 := symstring
__typedArg1 := tmp20949
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6655, tmp20950, B6651, L6652, Key6653, C6654)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20951 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20952 := Call(__e, PrimFunc(symshen_4assoc_1_6), symread_1file, tmp20945, tmp20951)


tmp20953 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20952)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20952
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20953

tmp20954 := MakeNative(func(__e *ControlFlow) {
V6660 := __e.Get(1)
_ = V6660
__e.Return(MakeNative(func(__e *ControlFlow) {
B6656 := __e.Get(1)
_ = B6656
__e.Return(MakeNative(func(__e *ControlFlow) {
L6657 := __e.Get(1)
_ = L6657
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6658 := __e.Get(1)
_ = Key6658
__e.Return(MakeNative(func(__e *ControlFlow) {
C6659 := __e.Get(1)
_ = C6659
tmp20955 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunit, Nil)
}
__typedArg0 := symunit
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20956 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20955)
}
__typedArg0 := symlist
__typedArg1 := tmp20955
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20957 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20956, Nil)
}
__typedArg0 := tmp20956
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20958 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20957)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20957
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20959 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp20958)
}
__typedArg0 := symstring
__typedArg1 := tmp20958
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6660, tmp20959, B6656, L6657, Key6658, C6659)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20960 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20961 := Call(__e, PrimFunc(symshen_4assoc_1_6), symread_1from_1string, tmp20954, tmp20960)


tmp20962 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20961)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20961
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20962

tmp20963 := MakeNative(func(__e *ControlFlow) {
V6665 := __e.Get(1)
_ = V6665
__e.Return(MakeNative(func(__e *ControlFlow) {
B6661 := __e.Get(1)
_ = B6661
__e.Return(MakeNative(func(__e *ControlFlow) {
L6662 := __e.Get(1)
_ = L6662
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6663 := __e.Get(1)
_ = Key6663
__e.Return(MakeNative(func(__e *ControlFlow) {
C6664 := __e.Get(1)
_ = C6664
tmp20964 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunit, Nil)
}
__typedArg0 := symunit
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20965 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20964)
}
__typedArg0 := symlist
__typedArg1 := tmp20964
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20966 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20965, Nil)
}
__typedArg0 := tmp20965
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20967 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20966)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20966
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20968 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp20967)
}
__typedArg0 := symstring
__typedArg1 := tmp20967
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6665, tmp20968, B6661, L6662, Key6663, C6664)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20969 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20970 := Call(__e, PrimFunc(symshen_4assoc_1_6), symread_1from_1string_1unprocessed, tmp20963, tmp20969)


tmp20971 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20970)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20970
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20971

tmp20972 := MakeNative(func(__e *ControlFlow) {
V6670 := __e.Get(1)
_ = V6670
__e.Return(MakeNative(func(__e *ControlFlow) {
B6666 := __e.Get(1)
_ = B6666
__e.Return(MakeNative(func(__e *ControlFlow) {
L6667 := __e.Get(1)
_ = L6667
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6668 := __e.Get(1)
_ = Key6668
__e.Return(MakeNative(func(__e *ControlFlow) {
C6669 := __e.Get(1)
_ = C6669
tmp20973 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20974 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20973)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20973
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6670, tmp20974, B6666, L6667, Key6668, C6669)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20975 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20976 := Call(__e, PrimFunc(symshen_4assoc_1_6), symrelease, tmp20972, tmp20975)


tmp20977 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20976)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20976
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20977

tmp20978 := MakeNative(func(__e *ControlFlow) {
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
tmp20979 := Call(__e, PrimFunc(symshen_4newpv), B6672)


let__18973 := tmp20979
_ = let__18973

tmp20980 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18973, Nil)
}
__typedArg0 := let__18973
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20981 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20980)
}
__typedArg0 := symlist
__typedArg1 := tmp20980
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20982 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18973, Nil)
}
__typedArg0 := let__18973
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20983 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20982)
}
__typedArg0 := symlist
__typedArg1 := tmp20982
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20984 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20983, Nil)
}
__typedArg0 := tmp20983
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20985 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20984)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20984
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20986 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20981, tmp20985)
}
__typedArg0 := tmp20981
__typedArg1 := tmp20985
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20987 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20986, Nil)
}
__typedArg0 := tmp20986
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20988 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20987)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20987
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20989 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18973, tmp20988)
}
__typedArg0 := let__18973
__typedArg1 := tmp20988
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20990 := Call(__e, PrimFunc(symis_b), V6676, tmp20989, B6672, L6673, Key6674, C6675)


__e.TailApply(PrimFunc(symshen_4gc), B6672, tmp20990)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20991 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp20992 := Call(__e, PrimFunc(symshen_4assoc_1_6), symremove, tmp20978, tmp20991)


tmp20993 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp20992)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp20992
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20993

tmp20994 := MakeNative(func(__e *ControlFlow) {
V6682 := __e.Get(1)
_ = V6682
__e.Return(MakeNative(func(__e *ControlFlow) {
B6678 := __e.Get(1)
_ = B6678
__e.Return(MakeNative(func(__e *ControlFlow) {
L6679 := __e.Get(1)
_ = L6679
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6680 := __e.Get(1)
_ = Key6680
__e.Return(MakeNative(func(__e *ControlFlow) {
C6681 := __e.Get(1)
_ = C6681
tmp20995 := Call(__e, PrimFunc(symshen_4newpv), B6678)


let__18974 := tmp20995
_ = let__18974

tmp20996 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18974, Nil)
}
__typedArg0 := let__18974
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20997 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20996)
}
__typedArg0 := symlist
__typedArg1 := tmp20996
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20998 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18974, Nil)
}
__typedArg0 := let__18974
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20999 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20998)
}
__typedArg0 := symlist
__typedArg1 := tmp20998
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21000 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20999, Nil)
}
__typedArg0 := tmp20999
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21001 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21000)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21000
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21002 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20997, tmp21001)
}
__typedArg0 := tmp20997
__typedArg1 := tmp21001
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21003 := Call(__e, PrimFunc(symis_b), V6682, tmp21002, B6678, L6679, Key6680, C6681)


__e.TailApply(PrimFunc(symshen_4gc), B6678, tmp21003)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21004 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21005 := Call(__e, PrimFunc(symshen_4assoc_1_6), symreverse, tmp20994, tmp21004)


tmp21006 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21005)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21005
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21006

tmp21007 := MakeNative(func(__e *ControlFlow) {
V6687 := __e.Get(1)
_ = V6687
__e.Return(MakeNative(func(__e *ControlFlow) {
B6683 := __e.Get(1)
_ = B6683
__e.Return(MakeNative(func(__e *ControlFlow) {
L6684 := __e.Get(1)
_ = L6684
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6685 := __e.Get(1)
_ = Key6685
__e.Return(MakeNative(func(__e *ControlFlow) {
C6686 := __e.Get(1)
_ = C6686
tmp21008 := Call(__e, PrimFunc(symshen_4newpv), B6683)


let__18975 := tmp21008
_ = let__18975

tmp21009 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18975, Nil)
}
__typedArg0 := let__18975
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21010 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21009)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21009
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21011 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp21010)
}
__typedArg0 := symstring
__typedArg1 := tmp21010
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21012 := Call(__e, PrimFunc(symis_b), V6687, tmp21011, B6683, L6684, Key6685, C6686)


__e.TailApply(PrimFunc(symshen_4gc), B6683, tmp21012)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21013 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21014 := Call(__e, PrimFunc(symshen_4assoc_1_6), symsimple_1error, tmp21007, tmp21013)


tmp21015 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21014)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21014
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21015

tmp21016 := MakeNative(func(__e *ControlFlow) {
V6692 := __e.Get(1)
_ = V6692
__e.Return(MakeNative(func(__e *ControlFlow) {
B6688 := __e.Get(1)
_ = B6688
__e.Return(MakeNative(func(__e *ControlFlow) {
L6689 := __e.Get(1)
_ = L6689
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6690 := __e.Get(1)
_ = Key6690
__e.Return(MakeNative(func(__e *ControlFlow) {
C6691 := __e.Get(1)
_ = C6691
tmp21017 := Call(__e, PrimFunc(symshen_4newpv), B6688)


let__18976 := tmp21017
_ = let__18976

tmp21018 := Call(__e, PrimFunc(symshen_4newpv), B6688)


let__18977 := tmp21018
_ = let__18977

tmp21019 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18977, Nil)
}
__typedArg0 := let__18977
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21020 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_d, tmp21019)
}
__typedArg0 := sym_d
__typedArg1 := tmp21019
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21021 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18976, tmp21020)
}
__typedArg0 := let__18976
__typedArg1 := tmp21020
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21022 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18977, Nil)
}
__typedArg0 := let__18977
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

tmp21024 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21021, tmp21023)
}
__typedArg0 := tmp21021
__typedArg1 := tmp21023
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21025 := Call(__e, PrimFunc(symis_b), V6692, tmp21024, B6688, L6689, Key6690, C6691)


tmp21026 := Call(__e, PrimFunc(symshen_4gc), B6688, tmp21025)


__e.TailApply(PrimFunc(symshen_4gc), B6688, tmp21026)
return


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

tmp21028 := Call(__e, PrimFunc(symshen_4assoc_1_6), symsnd, tmp21016, tmp21027)


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
tmp21031 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
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
return PrimCons(symnumber, tmp21032)
}
__typedArg0 := symnumber
__typedArg1 := tmp21032
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21034 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21033, Nil)
}
__typedArg0 := tmp21033
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21035 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21034)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21034
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21036 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp21035)
}
__typedArg0 := symsymbol
__typedArg1 := tmp21035
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6697, tmp21036, B6693, L6694, Key6695, C6696)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21037 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21038 := Call(__e, PrimFunc(symshen_4assoc_1_6), symspecialise, tmp21030, tmp21037)


tmp21039 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21038)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21038
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21039

tmp21040 := MakeNative(func(__e *ControlFlow) {
V6702 := __e.Get(1)
_ = V6702
__e.Return(MakeNative(func(__e *ControlFlow) {
B6698 := __e.Get(1)
_ = B6698
__e.Return(MakeNative(func(__e *ControlFlow) {
L6699 := __e.Get(1)
_ = L6699
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6700 := __e.Get(1)
_ = Key6700
__e.Return(MakeNative(func(__e *ControlFlow) {
C6701 := __e.Get(1)
_ = C6701
tmp21041 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21042 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21041)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21041
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21043 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp21042)
}
__typedArg0 := symsymbol
__typedArg1 := tmp21042
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6702, tmp21043, B6698, L6699, Key6700, C6701)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21044 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21045 := Call(__e, PrimFunc(symshen_4assoc_1_6), symspy, tmp21040, tmp21044)


tmp21046 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21045)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21045
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21046

tmp21047 := MakeNative(func(__e *ControlFlow) {
V6707 := __e.Get(1)
_ = V6707
__e.Return(MakeNative(func(__e *ControlFlow) {
B6703 := __e.Get(1)
_ = B6703
__e.Return(MakeNative(func(__e *ControlFlow) {
L6704 := __e.Get(1)
_ = L6704
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6705 := __e.Get(1)
_ = Key6705
__e.Return(MakeNative(func(__e *ControlFlow) {
C6706 := __e.Get(1)
_ = C6706
tmp21048 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21049 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21048)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21048
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6707, tmp21049, B6703, L6704, Key6705, C6706)
return


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

tmp21051 := Call(__e, PrimFunc(symshen_4assoc_1_6), symspy_2, tmp21047, tmp21050)


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
V6712 := __e.Get(1)
_ = V6712
__e.Return(MakeNative(func(__e *ControlFlow) {
B6708 := __e.Get(1)
_ = B6708
__e.Return(MakeNative(func(__e *ControlFlow) {
L6709 := __e.Get(1)
_ = L6709
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6710 := __e.Get(1)
_ = Key6710
__e.Return(MakeNative(func(__e *ControlFlow) {
C6711 := __e.Get(1)
_ = C6711
tmp21054 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
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
return PrimCons(symsymbol, tmp21055)
}
__typedArg0 := symsymbol
__typedArg1 := tmp21055
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6712, tmp21056, B6708, L6709, Key6710, C6711)
return


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

tmp21058 := Call(__e, PrimFunc(symshen_4assoc_1_6), symstep, tmp21053, tmp21057)


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
V6717 := __e.Get(1)
_ = V6717
__e.Return(MakeNative(func(__e *ControlFlow) {
B6713 := __e.Get(1)
_ = B6713
__e.Return(MakeNative(func(__e *ControlFlow) {
L6714 := __e.Get(1)
_ = L6714
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6715 := __e.Get(1)
_ = Key6715
__e.Return(MakeNative(func(__e *ControlFlow) {
C6716 := __e.Get(1)
_ = C6716
tmp21061 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
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

__e.TailApply(PrimFunc(symis_b), V6717, tmp21062, B6713, L6714, Key6715, C6716)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21063 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21064 := Call(__e, PrimFunc(symshen_4assoc_1_6), symstep_2, tmp21060, tmp21063)


tmp21065 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21064)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21064
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21065

tmp21066 := MakeNative(func(__e *ControlFlow) {
V6722 := __e.Get(1)
_ = V6722
__e.Return(MakeNative(func(__e *ControlFlow) {
B6718 := __e.Get(1)
_ = B6718
__e.Return(MakeNative(func(__e *ControlFlow) {
L6719 := __e.Get(1)
_ = L6719
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6720 := __e.Get(1)
_ = Key6720
__e.Return(MakeNative(func(__e *ControlFlow) {
C6721 := __e.Get(1)
_ = C6721
tmp21067 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symin, Nil)
}
__typedArg0 := symin
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21068 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp21067)
}
__typedArg0 := symstream
__typedArg1 := tmp21067
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21069 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21068, Nil)
}
__typedArg0 := tmp21068
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21070 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21069)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21069
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6722, tmp21070, B6718, L6719, Key6720, C6721)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21071 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21072 := Call(__e, PrimFunc(symshen_4assoc_1_6), symstinput, tmp21066, tmp21071)


tmp21073 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21072)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21072
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21073

tmp21074 := MakeNative(func(__e *ControlFlow) {
V6727 := __e.Get(1)
_ = V6727
__e.Return(MakeNative(func(__e *ControlFlow) {
B6723 := __e.Get(1)
_ = B6723
__e.Return(MakeNative(func(__e *ControlFlow) {
L6724 := __e.Get(1)
_ = L6724
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6725 := __e.Get(1)
_ = Key6725
__e.Return(MakeNative(func(__e *ControlFlow) {
C6726 := __e.Get(1)
_ = C6726
tmp21075 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symout, Nil)
}
__typedArg0 := symout
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21076 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp21075)
}
__typedArg0 := symstream
__typedArg1 := tmp21075
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21077 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21076, Nil)
}
__typedArg0 := tmp21076
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

__e.TailApply(PrimFunc(symis_b), V6727, tmp21078, B6723, L6724, Key6725, C6726)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21079 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21080 := Call(__e, PrimFunc(symshen_4assoc_1_6), symsterror, tmp21074, tmp21079)


tmp21081 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21080)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21080
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21081

tmp21082 := MakeNative(func(__e *ControlFlow) {
V6732 := __e.Get(1)
_ = V6732
__e.Return(MakeNative(func(__e *ControlFlow) {
B6728 := __e.Get(1)
_ = B6728
__e.Return(MakeNative(func(__e *ControlFlow) {
L6729 := __e.Get(1)
_ = L6729
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6730 := __e.Get(1)
_ = Key6730
__e.Return(MakeNative(func(__e *ControlFlow) {
C6731 := __e.Get(1)
_ = C6731
tmp21083 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symout, Nil)
}
__typedArg0 := symout
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21084 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp21083)
}
__typedArg0 := symstream
__typedArg1 := tmp21083
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21085 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21084, Nil)
}
__typedArg0 := tmp21084
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21086 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21085)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21085
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6732, tmp21086, B6728, L6729, Key6730, C6731)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21087 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21088 := Call(__e, PrimFunc(symshen_4assoc_1_6), symstoutput, tmp21082, tmp21087)


tmp21089 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21088)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21088
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21089

tmp21090 := MakeNative(func(__e *ControlFlow) {
V6737 := __e.Get(1)
_ = V6737
__e.Return(MakeNative(func(__e *ControlFlow) {
B6733 := __e.Get(1)
_ = B6733
__e.Return(MakeNative(func(__e *ControlFlow) {
L6734 := __e.Get(1)
_ = L6734
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6735 := __e.Get(1)
_ = Key6735
__e.Return(MakeNative(func(__e *ControlFlow) {
C6736 := __e.Get(1)
_ = C6736
tmp21091 := Call(__e, PrimFunc(symshen_4newpv), B6733)


let__18978 := tmp21091
_ = let__18978

tmp21092 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
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

tmp21094 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18978, tmp21093)
}
__typedArg0 := let__18978
__typedArg1 := tmp21093
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21095 := Call(__e, PrimFunc(symis_b), V6737, tmp21094, B6733, L6734, Key6735, C6736)


__e.TailApply(PrimFunc(symshen_4gc), B6733, tmp21095)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21096 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21097 := Call(__e, PrimFunc(symshen_4assoc_1_6), symstring_2, tmp21090, tmp21096)


tmp21098 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21097)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21097
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21098

tmp21099 := MakeNative(func(__e *ControlFlow) {
V6742 := __e.Get(1)
_ = V6742
__e.Return(MakeNative(func(__e *ControlFlow) {
B6738 := __e.Get(1)
_ = B6738
__e.Return(MakeNative(func(__e *ControlFlow) {
L6739 := __e.Get(1)
_ = L6739
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6740 := __e.Get(1)
_ = Key6740
__e.Return(MakeNative(func(__e *ControlFlow) {
C6741 := __e.Get(1)
_ = C6741
tmp21100 := Call(__e, PrimFunc(symshen_4newpv), B6738)


let__18979 := tmp21100
_ = let__18979

tmp21101 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21102 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21101)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21101
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21103 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18979, tmp21102)
}
__typedArg0 := let__18979
__typedArg1 := tmp21102
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21104 := Call(__e, PrimFunc(symis_b), V6742, tmp21103, B6738, L6739, Key6740, C6741)


__e.TailApply(PrimFunc(symshen_4gc), B6738, tmp21104)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21105 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21106 := Call(__e, PrimFunc(symshen_4assoc_1_6), symstr, tmp21099, tmp21105)


tmp21107 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21106)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21106
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21107

tmp21108 := MakeNative(func(__e *ControlFlow) {
V6747 := __e.Get(1)
_ = V6747
__e.Return(MakeNative(func(__e *ControlFlow) {
B6743 := __e.Get(1)
_ = B6743
__e.Return(MakeNative(func(__e *ControlFlow) {
L6744 := __e.Get(1)
_ = L6744
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6745 := __e.Get(1)
_ = Key6745
__e.Return(MakeNative(func(__e *ControlFlow) {
C6746 := __e.Get(1)
_ = C6746
tmp21109 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21110 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21109)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21109
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21111 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp21110)
}
__typedArg0 := symstring
__typedArg1 := tmp21110
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6747, tmp21111, B6743, L6744, Key6745, C6746)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21112 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21113 := Call(__e, PrimFunc(symshen_4assoc_1_6), symstring_1_6n, tmp21108, tmp21112)


tmp21114 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21113)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21113
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21114

tmp21115 := MakeNative(func(__e *ControlFlow) {
V6752 := __e.Get(1)
_ = V6752
__e.Return(MakeNative(func(__e *ControlFlow) {
B6748 := __e.Get(1)
_ = B6748
__e.Return(MakeNative(func(__e *ControlFlow) {
L6749 := __e.Get(1)
_ = L6749
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6750 := __e.Get(1)
_ = Key6750
__e.Return(MakeNative(func(__e *ControlFlow) {
C6751 := __e.Get(1)
_ = C6751
tmp21116 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21117 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21116)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21116
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21118 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp21117)
}
__typedArg0 := symstring
__typedArg1 := tmp21117
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6752, tmp21118, B6748, L6749, Key6750, C6751)
return


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

tmp21120 := Call(__e, PrimFunc(symshen_4assoc_1_6), symstring_1_6symbol, tmp21115, tmp21119)


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
V6757 := __e.Get(1)
_ = V6757
__e.Return(MakeNative(func(__e *ControlFlow) {
B6753 := __e.Get(1)
_ = B6753
__e.Return(MakeNative(func(__e *ControlFlow) {
L6754 := __e.Get(1)
_ = L6754
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6755 := __e.Get(1)
_ = Key6755
__e.Return(MakeNative(func(__e *ControlFlow) {
C6756 := __e.Get(1)
_ = C6756
tmp21123 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21124 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp21123)
}
__typedArg0 := symlist
__typedArg1 := tmp21123
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21125 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21126 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21125)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21125
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21127 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21124, tmp21126)
}
__typedArg0 := tmp21124
__typedArg1 := tmp21126
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6757, tmp21127, B6753, L6754, Key6755, C6756)
return


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

tmp21129 := Call(__e, PrimFunc(symshen_4assoc_1_6), symsum, tmp21122, tmp21128)


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
V6762 := __e.Get(1)
_ = V6762
__e.Return(MakeNative(func(__e *ControlFlow) {
B6758 := __e.Get(1)
_ = B6758
__e.Return(MakeNative(func(__e *ControlFlow) {
L6759 := __e.Get(1)
_ = L6759
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6760 := __e.Get(1)
_ = Key6760
__e.Return(MakeNative(func(__e *ControlFlow) {
C6761 := __e.Get(1)
_ = C6761
tmp21132 := Call(__e, PrimFunc(symshen_4newpv), B6758)


let__18980 := tmp21132
_ = let__18980

tmp21133 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21134 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21133)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21133
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21135 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18980, tmp21134)
}
__typedArg0 := let__18980
__typedArg1 := tmp21134
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21136 := Call(__e, PrimFunc(symis_b), V6762, tmp21135, B6758, L6759, Key6760, C6761)


__e.TailApply(PrimFunc(symshen_4gc), B6758, tmp21136)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21137 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21138 := Call(__e, PrimFunc(symshen_4assoc_1_6), symsymbol_2, tmp21131, tmp21137)


tmp21139 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21138)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21138
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21139

tmp21140 := MakeNative(func(__e *ControlFlow) {
V6767 := __e.Get(1)
_ = V6767
__e.Return(MakeNative(func(__e *ControlFlow) {
B6763 := __e.Get(1)
_ = B6763
__e.Return(MakeNative(func(__e *ControlFlow) {
L6764 := __e.Get(1)
_ = L6764
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6765 := __e.Get(1)
_ = Key6765
__e.Return(MakeNative(func(__e *ControlFlow) {
C6766 := __e.Get(1)
_ = C6766
tmp21141 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21142 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21141)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21141
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21143 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp21142)
}
__typedArg0 := symsymbol
__typedArg1 := tmp21142
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6767, tmp21143, B6763, L6764, Key6765, C6766)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21144 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21145 := Call(__e, PrimFunc(symshen_4assoc_1_6), symsystemf, tmp21140, tmp21144)


tmp21146 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21145)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21145
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21146

tmp21147 := MakeNative(func(__e *ControlFlow) {
V6772 := __e.Get(1)
_ = V6772
__e.Return(MakeNative(func(__e *ControlFlow) {
B6768 := __e.Get(1)
_ = B6768
__e.Return(MakeNative(func(__e *ControlFlow) {
L6769 := __e.Get(1)
_ = L6769
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6770 := __e.Get(1)
_ = Key6770
__e.Return(MakeNative(func(__e *ControlFlow) {
C6771 := __e.Get(1)
_ = C6771
tmp21148 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
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

__e.TailApply(PrimFunc(symis_b), V6772, tmp21149, B6768, L6769, Key6770, C6771)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21150 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21151 := Call(__e, PrimFunc(symshen_4assoc_1_6), symsystem_1S_2, tmp21147, tmp21150)


tmp21152 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21151)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21151
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21152

tmp21153 := MakeNative(func(__e *ControlFlow) {
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
tmp21154 := Call(__e, PrimFunc(symshen_4newpv), B6774)


let__18981 := tmp21154
_ = let__18981

tmp21155 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18981, Nil)
}
__typedArg0 := let__18981
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21156 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp21155)
}
__typedArg0 := symlist
__typedArg1 := tmp21155
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21157 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18981, Nil)
}
__typedArg0 := let__18981
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21158 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp21157)
}
__typedArg0 := symlist
__typedArg1 := tmp21157
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21159 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21158, Nil)
}
__typedArg0 := tmp21158
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21160 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21159)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21159
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21161 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21156, tmp21160)
}
__typedArg0 := tmp21156
__typedArg1 := tmp21160
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21162 := Call(__e, PrimFunc(symis_b), V6778, tmp21161, B6774, L6775, Key6776, C6777)


__e.TailApply(PrimFunc(symshen_4gc), B6774, tmp21162)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21163 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21164 := Call(__e, PrimFunc(symshen_4assoc_1_6), symtail, tmp21153, tmp21163)


tmp21165 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21164)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21164
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21165

tmp21166 := MakeNative(func(__e *ControlFlow) {
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
tmp21167 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21168 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21167)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21167
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21169 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp21168)
}
__typedArg0 := symstring
__typedArg1 := tmp21168
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6783, tmp21169, B6779, L6780, Key6781, C6782)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21170 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21171 := Call(__e, PrimFunc(symshen_4assoc_1_6), symtlstr, tmp21166, tmp21170)


tmp21172 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21171)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21171
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21172

tmp21173 := MakeNative(func(__e *ControlFlow) {
V6789 := __e.Get(1)
_ = V6789
__e.Return(MakeNative(func(__e *ControlFlow) {
B6785 := __e.Get(1)
_ = B6785
__e.Return(MakeNative(func(__e *ControlFlow) {
L6786 := __e.Get(1)
_ = L6786
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6787 := __e.Get(1)
_ = Key6787
__e.Return(MakeNative(func(__e *ControlFlow) {
C6788 := __e.Get(1)
_ = C6788
tmp21174 := Call(__e, PrimFunc(symshen_4newpv), B6785)


let__18982 := tmp21174
_ = let__18982

tmp21175 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18982, Nil)
}
__typedArg0 := let__18982
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21176 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp21175)
}
__typedArg0 := symvector
__typedArg1 := tmp21175
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21177 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18982, Nil)
}
__typedArg0 := let__18982
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21178 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp21177)
}
__typedArg0 := symvector
__typedArg1 := tmp21177
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21179 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21178, Nil)
}
__typedArg0 := tmp21178
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21180 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21179)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21179
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21181 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21176, tmp21180)
}
__typedArg0 := tmp21176
__typedArg1 := tmp21180
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21182 := Call(__e, PrimFunc(symis_b), V6789, tmp21181, B6785, L6786, Key6787, C6788)


__e.TailApply(PrimFunc(symshen_4gc), B6785, tmp21182)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21183 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21184 := Call(__e, PrimFunc(symshen_4assoc_1_6), symtlv, tmp21173, tmp21183)


tmp21185 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21184)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21184
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21185

tmp21186 := MakeNative(func(__e *ControlFlow) {
V6794 := __e.Get(1)
_ = V6794
__e.Return(MakeNative(func(__e *ControlFlow) {
B6790 := __e.Get(1)
_ = B6790
__e.Return(MakeNative(func(__e *ControlFlow) {
L6791 := __e.Get(1)
_ = L6791
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6792 := __e.Get(1)
_ = Key6792
__e.Return(MakeNative(func(__e *ControlFlow) {
C6793 := __e.Get(1)
_ = C6793
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
return PrimCons(symsymbol, tmp21188)
}
__typedArg0 := symsymbol
__typedArg1 := tmp21188
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6794, tmp21189, B6790, L6791, Key6792, C6793)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21190 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21191 := Call(__e, PrimFunc(symshen_4assoc_1_6), symtc, tmp21186, tmp21190)


tmp21192 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21191)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21191
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21192

tmp21193 := MakeNative(func(__e *ControlFlow) {
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
tmp21194 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21195 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21194)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21194
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6799, tmp21195, B6795, L6796, Key6797, C6798)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21196 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21197 := Call(__e, PrimFunc(symshen_4assoc_1_6), symtc_2, tmp21193, tmp21196)


tmp21198 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21197)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21197
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21198

tmp21199 := MakeNative(func(__e *ControlFlow) {
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
tmp21200 := Call(__e, PrimFunc(symshen_4newpv), B6800)


let__18983 := tmp21200
_ = let__18983

tmp21201 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18983, Nil)
}
__typedArg0 := let__18983
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21202 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlazy, tmp21201)
}
__typedArg0 := symlazy
__typedArg1 := tmp21201
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21203 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18983, Nil)
}
__typedArg0 := let__18983
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21204 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21203)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21203
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21205 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21202, tmp21204)
}
__typedArg0 := tmp21202
__typedArg1 := tmp21204
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21206 := Call(__e, PrimFunc(symis_b), V6804, tmp21205, B6800, L6801, Key6802, C6803)


__e.TailApply(PrimFunc(symshen_4gc), B6800, tmp21206)
return


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

tmp21208 := Call(__e, PrimFunc(symshen_4assoc_1_6), symthaw, tmp21199, tmp21207)


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
V6809 := __e.Get(1)
_ = V6809
__e.Return(MakeNative(func(__e *ControlFlow) {
B6805 := __e.Get(1)
_ = B6805
__e.Return(MakeNative(func(__e *ControlFlow) {
L6806 := __e.Get(1)
_ = L6806
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6807 := __e.Get(1)
_ = Key6807
__e.Return(MakeNative(func(__e *ControlFlow) {
C6808 := __e.Get(1)
_ = C6808
tmp21211 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21212 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21211)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21211
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21213 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp21212)
}
__typedArg0 := symsymbol
__typedArg1 := tmp21212
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6809, tmp21213, B6805, L6806, Key6807, C6808)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21214 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21215 := Call(__e, PrimFunc(symshen_4assoc_1_6), symtrack, tmp21210, tmp21214)


tmp21216 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21215)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21215
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21216

tmp21217 := MakeNative(func(__e *ControlFlow) {
V6814 := __e.Get(1)
_ = V6814
__e.Return(MakeNative(func(__e *ControlFlow) {
B6810 := __e.Get(1)
_ = B6810
__e.Return(MakeNative(func(__e *ControlFlow) {
L6811 := __e.Get(1)
_ = L6811
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6812 := __e.Get(1)
_ = Key6812
__e.Return(MakeNative(func(__e *ControlFlow) {
C6813 := __e.Get(1)
_ = C6813
tmp21218 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21219 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp21218)
}
__typedArg0 := symlist
__typedArg1 := tmp21218
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21220 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21219, Nil)
}
__typedArg0 := tmp21219
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

__e.TailApply(PrimFunc(symis_b), V6814, tmp21221, B6810, L6811, Key6812, C6813)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21222 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21223 := Call(__e, PrimFunc(symshen_4assoc_1_6), symtracked, tmp21217, tmp21222)


tmp21224 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21223)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21223
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21224

tmp21225 := MakeNative(func(__e *ControlFlow) {
V6819 := __e.Get(1)
_ = V6819
__e.Return(MakeNative(func(__e *ControlFlow) {
B6815 := __e.Get(1)
_ = B6815
__e.Return(MakeNative(func(__e *ControlFlow) {
L6816 := __e.Get(1)
_ = L6816
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6817 := __e.Get(1)
_ = Key6817
__e.Return(MakeNative(func(__e *ControlFlow) {
C6818 := __e.Get(1)
_ = C6818
tmp21226 := Call(__e, PrimFunc(symshen_4newpv), B6815)


let__18984 := tmp21226
_ = let__18984

tmp21227 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18984, Nil)
}
__typedArg0 := let__18984
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21228 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21227)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21227
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21229 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symexception, tmp21228)
}
__typedArg0 := symexception
__typedArg1 := tmp21228
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21230 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18984, Nil)
}
__typedArg0 := let__18984
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21231 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21230)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21230
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21232 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21229, tmp21231)
}
__typedArg0 := tmp21229
__typedArg1 := tmp21231
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21233 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21232, Nil)
}
__typedArg0 := tmp21232
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21234 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21233)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21233
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21235 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18984, tmp21234)
}
__typedArg0 := let__18984
__typedArg1 := tmp21234
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21236 := Call(__e, PrimFunc(symis_b), V6819, tmp21235, B6815, L6816, Key6817, C6818)


__e.TailApply(PrimFunc(symshen_4gc), B6815, tmp21236)
return


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

tmp21238 := Call(__e, PrimFunc(symshen_4assoc_1_6), symtrap_1error, tmp21225, tmp21237)


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
V6824 := __e.Get(1)
_ = V6824
__e.Return(MakeNative(func(__e *ControlFlow) {
B6820 := __e.Get(1)
_ = B6820
__e.Return(MakeNative(func(__e *ControlFlow) {
L6821 := __e.Get(1)
_ = L6821
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6822 := __e.Get(1)
_ = Key6822
__e.Return(MakeNative(func(__e *ControlFlow) {
C6823 := __e.Get(1)
_ = C6823
tmp21241 := Call(__e, PrimFunc(symshen_4newpv), B6820)


let__18985 := tmp21241
_ = let__18985

tmp21242 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
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
return PrimCons(let__18985, tmp21243)
}
__typedArg0 := let__18985
__typedArg1 := tmp21243
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21245 := Call(__e, PrimFunc(symis_b), V6824, tmp21244, B6820, L6821, Key6822, C6823)


__e.TailApply(PrimFunc(symshen_4gc), B6820, tmp21245)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21246 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21247 := Call(__e, PrimFunc(symshen_4assoc_1_6), symtuple_2, tmp21240, tmp21246)


tmp21248 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21247)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21247
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21248

tmp21249 := MakeNative(func(__e *ControlFlow) {
V6829 := __e.Get(1)
_ = V6829
__e.Return(MakeNative(func(__e *ControlFlow) {
B6825 := __e.Get(1)
_ = B6825
__e.Return(MakeNative(func(__e *ControlFlow) {
L6826 := __e.Get(1)
_ = L6826
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6827 := __e.Get(1)
_ = Key6827
__e.Return(MakeNative(func(__e *ControlFlow) {
C6828 := __e.Get(1)
_ = C6828
tmp21250 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21251 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp21250)
}
__typedArg0 := symlist
__typedArg1 := tmp21250
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21252 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21251, Nil)
}
__typedArg0 := tmp21251
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
return PrimCons(symstring, tmp21253)
}
__typedArg0 := symstring
__typedArg1 := tmp21253
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6829, tmp21254, B6825, L6826, Key6827, C6828)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21255 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21256 := Call(__e, PrimFunc(symshen_4assoc_1_6), symunabsolute, tmp21249, tmp21255)


tmp21257 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21256)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21256
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21257

tmp21258 := MakeNative(func(__e *ControlFlow) {
V6834 := __e.Get(1)
_ = V6834
__e.Return(MakeNative(func(__e *ControlFlow) {
B6830 := __e.Get(1)
_ = B6830
__e.Return(MakeNative(func(__e *ControlFlow) {
L6831 := __e.Get(1)
_ = L6831
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6832 := __e.Get(1)
_ = Key6832
__e.Return(MakeNative(func(__e *ControlFlow) {
C6833 := __e.Get(1)
_ = C6833
tmp21259 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21260 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21259)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21259
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21261 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp21260)
}
__typedArg0 := symsymbol
__typedArg1 := tmp21260
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6834, tmp21261, B6830, L6831, Key6832, C6833)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21262 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21263 := Call(__e, PrimFunc(symshen_4assoc_1_6), symundefmacro, tmp21258, tmp21262)


tmp21264 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21263)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21263
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21264

tmp21265 := MakeNative(func(__e *ControlFlow) {
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
tmp21266 := Call(__e, PrimFunc(symshen_4newpv), B6837)


let__18986 := tmp21266
_ = let__18986

tmp21267 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18986, Nil)
}
__typedArg0 := let__18986
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21268 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp21267)
}
__typedArg0 := symlist
__typedArg1 := tmp21267
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21269 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18986, Nil)
}
__typedArg0 := let__18986
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21270 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp21269)
}
__typedArg0 := symlist
__typedArg1 := tmp21269
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21271 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18986, Nil)
}
__typedArg0 := let__18986
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21272 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp21271)
}
__typedArg0 := symlist
__typedArg1 := tmp21271
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21273 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21272, Nil)
}
__typedArg0 := tmp21272
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21274 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21273)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21273
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21275 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21270, tmp21274)
}
__typedArg0 := tmp21270
__typedArg1 := tmp21274
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21276 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21275, Nil)
}
__typedArg0 := tmp21275
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
return PrimCons(tmp21268, tmp21277)
}
__typedArg0 := tmp21268
__typedArg1 := tmp21277
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21279 := Call(__e, PrimFunc(symis_b), V6841, tmp21278, B6837, L6838, Key6839, C6840)


__e.TailApply(PrimFunc(symshen_4gc), B6837, tmp21279)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21280 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21281 := Call(__e, PrimFunc(symshen_4assoc_1_6), symunion, tmp21265, tmp21280)


tmp21282 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21281)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21281
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21282

tmp21283 := MakeNative(func(__e *ControlFlow) {
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
tmp21284 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21285 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21284)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21284
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21286 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp21285)
}
__typedArg0 := symsymbol
__typedArg1 := tmp21285
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6846, tmp21286, B6842, L6843, Key6844, C6845)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21287 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21288 := Call(__e, PrimFunc(symshen_4assoc_1_6), symunprofile, tmp21283, tmp21287)


tmp21289 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21288)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21288
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21289

tmp21290 := MakeNative(func(__e *ControlFlow) {
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
tmp21291 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21292 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21291)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21291
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21293 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp21292)
}
__typedArg0 := symsymbol
__typedArg1 := tmp21292
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6851, tmp21293, B6847, L6848, Key6849, C6850)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21294 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21295 := Call(__e, PrimFunc(symshen_4assoc_1_6), symuntrack, tmp21290, tmp21294)


tmp21296 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21295)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21295
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21296

tmp21297 := MakeNative(func(__e *ControlFlow) {
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
tmp21298 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21299 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp21298)
}
__typedArg0 := symlist
__typedArg1 := tmp21298
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21300 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21299, Nil)
}
__typedArg0 := tmp21299
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21301 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21300)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21300
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6856, tmp21301, B6852, L6853, Key6854, C6855)
return


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

tmp21303 := Call(__e, PrimFunc(symshen_4assoc_1_6), symuserdefs, tmp21297, tmp21302)


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
tmp21306 := Call(__e, PrimFunc(symshen_4newpv), B6857)


let__18987 := tmp21306
_ = let__18987

tmp21307 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
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
return PrimCons(let__18987, tmp21308)
}
__typedArg0 := let__18987
__typedArg1 := tmp21308
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21310 := Call(__e, PrimFunc(symis_b), V6861, tmp21309, B6857, L6858, Key6859, C6860)


__e.TailApply(PrimFunc(symshen_4gc), B6857, tmp21310)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21311 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21312 := Call(__e, PrimFunc(symshen_4assoc_1_6), symvariable_2, tmp21305, tmp21311)


tmp21313 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21312)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21312
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21313

tmp21314 := MakeNative(func(__e *ControlFlow) {
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
tmp21315 := Call(__e, PrimFunc(symshen_4newpv), B6862)


let__18988 := tmp21315
_ = let__18988

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
return PrimCons(let__18988, tmp21317)
}
__typedArg0 := let__18988
__typedArg1 := tmp21317
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21319 := Call(__e, PrimFunc(symis_b), V6866, tmp21318, B6862, L6863, Key6864, C6865)


__e.TailApply(PrimFunc(symshen_4gc), B6862, tmp21319)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21320 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21321 := Call(__e, PrimFunc(symshen_4assoc_1_6), symvector_2, tmp21314, tmp21320)


tmp21322 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21321)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21321
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21322

tmp21323 := MakeNative(func(__e *ControlFlow) {
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
tmp21324 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
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

__e.TailApply(PrimFunc(symis_b), V6871, tmp21325, B6867, L6868, Key6869, C6870)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21326 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21327 := Call(__e, PrimFunc(symshen_4assoc_1_6), symversion, tmp21323, tmp21326)


tmp21328 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21327)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21327
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21328

tmp21329 := MakeNative(func(__e *ControlFlow) {
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
tmp21330 := Call(__e, PrimFunc(symshen_4newpv), B6872)


let__18989 := tmp21330
_ = let__18989

tmp21331 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18989, Nil)
}
__typedArg0 := let__18989
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21332 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21331)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21331
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21333 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18989, tmp21332)
}
__typedArg0 := let__18989
__typedArg1 := tmp21332
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21334 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21333, Nil)
}
__typedArg0 := tmp21333
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
return PrimCons(symstring, tmp21335)
}
__typedArg0 := symstring
__typedArg1 := tmp21335
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21337 := Call(__e, PrimFunc(symis_b), V6876, tmp21336, B6872, L6873, Key6874, C6875)


__e.TailApply(PrimFunc(symshen_4gc), B6872, tmp21337)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21338 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21339 := Call(__e, PrimFunc(symshen_4assoc_1_6), symwrite_1to_1file, tmp21329, tmp21338)


tmp21340 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21339)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21339
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21340

tmp21341 := MakeNative(func(__e *ControlFlow) {
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
tmp21342 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symout, Nil)
}
__typedArg0 := symout
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21343 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp21342)
}
__typedArg0 := symstream
__typedArg1 := tmp21342
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21344 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21345 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21344)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21344
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21346 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21343, tmp21345)
}
__typedArg0 := tmp21343
__typedArg1 := tmp21345
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21347 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21346, Nil)
}
__typedArg0 := tmp21346
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
return PrimCons(symnumber, tmp21348)
}
__typedArg0 := symnumber
__typedArg1 := tmp21348
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6881, tmp21349, B6877, L6878, Key6879, C6880)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21350 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21351 := Call(__e, PrimFunc(symshen_4assoc_1_6), symwrite_1byte, tmp21341, tmp21350)


tmp21352 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21351)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21351
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21352

tmp21353 := MakeNative(func(__e *ControlFlow) {
V6886 := __e.Get(1)
_ = V6886
__e.Return(MakeNative(func(__e *ControlFlow) {
B6882 := __e.Get(1)
_ = B6882
__e.Return(MakeNative(func(__e *ControlFlow) {
L6883 := __e.Get(1)
_ = L6883
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6884 := __e.Get(1)
_ = Key6884
__e.Return(MakeNative(func(__e *ControlFlow) {
C6885 := __e.Get(1)
_ = C6885
tmp21354 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symout, Nil)
}
__typedArg0 := symout
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21355 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp21354)
}
__typedArg0 := symstream
__typedArg1 := tmp21354
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21356 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
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
return PrimCons(tmp21355, tmp21357)
}
__typedArg0 := tmp21355
__typedArg1 := tmp21357
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21359 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21358, Nil)
}
__typedArg0 := tmp21358
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21360 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21359)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21359
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21361 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp21360)
}
__typedArg0 := symstring
__typedArg1 := tmp21360
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6886, tmp21361, B6882, L6883, Key6884, C6885)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21362 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21363 := Call(__e, PrimFunc(symshen_4assoc_1_6), symwrite_1string, tmp21353, tmp21362)


tmp21364 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21363)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21363
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21364

tmp21365 := MakeNative(func(__e *ControlFlow) {
V6891 := __e.Get(1)
_ = V6891
__e.Return(MakeNative(func(__e *ControlFlow) {
B6887 := __e.Get(1)
_ = B6887
__e.Return(MakeNative(func(__e *ControlFlow) {
L6888 := __e.Get(1)
_ = L6888
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6889 := __e.Get(1)
_ = Key6889
__e.Return(MakeNative(func(__e *ControlFlow) {
C6890 := __e.Get(1)
_ = C6890
tmp21366 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
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
return PrimCons(symstring, tmp21367)
}
__typedArg0 := symstring
__typedArg1 := tmp21367
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symis_b), V6891, tmp21368, B6887, L6888, Key6889, C6890)
return


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

tmp21370 := Call(__e, PrimFunc(symshen_4assoc_1_6), symy_1or_1n_2, tmp21365, tmp21369)


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
V6896 := __e.Get(1)
_ = V6896
__e.Return(MakeNative(func(__e *ControlFlow) {
B6892 := __e.Get(1)
_ = B6892
__e.Return(MakeNative(func(__e *ControlFlow) {
L6893 := __e.Get(1)
_ = L6893
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6894 := __e.Get(1)
_ = Key6894
__e.Return(MakeNative(func(__e *ControlFlow) {
C6895 := __e.Get(1)
_ = C6895
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

__e.TailApply(PrimFunc(symis_b), V6896, tmp21378, B6892, L6893, Key6894, C6895)
return


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

tmp21380 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_6, tmp21372, tmp21379)


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
V6901 := __e.Get(1)
_ = V6901
__e.Return(MakeNative(func(__e *ControlFlow) {
B6897 := __e.Get(1)
_ = B6897
__e.Return(MakeNative(func(__e *ControlFlow) {
L6898 := __e.Get(1)
_ = L6898
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6899 := __e.Get(1)
_ = Key6899
__e.Return(MakeNative(func(__e *ControlFlow) {
C6900 := __e.Get(1)
_ = C6900
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

__e.TailApply(PrimFunc(symis_b), V6901, tmp21388, B6897, L6898, Key6899, C6900)
return


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

tmp21390 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_5, tmp21382, tmp21389)


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
V6906 := __e.Get(1)
_ = V6906
__e.Return(MakeNative(func(__e *ControlFlow) {
B6902 := __e.Get(1)
_ = B6902
__e.Return(MakeNative(func(__e *ControlFlow) {
L6903 := __e.Get(1)
_ = L6903
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6904 := __e.Get(1)
_ = Key6904
__e.Return(MakeNative(func(__e *ControlFlow) {
C6905 := __e.Get(1)
_ = C6905
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

__e.TailApply(PrimFunc(symis_b), V6906, tmp21398, B6902, L6903, Key6904, C6905)
return


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

tmp21400 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_6_a, tmp21392, tmp21399)


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
V6911 := __e.Get(1)
_ = V6911
__e.Return(MakeNative(func(__e *ControlFlow) {
B6907 := __e.Get(1)
_ = B6907
__e.Return(MakeNative(func(__e *ControlFlow) {
L6908 := __e.Get(1)
_ = L6908
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6909 := __e.Get(1)
_ = Key6909
__e.Return(MakeNative(func(__e *ControlFlow) {
C6910 := __e.Get(1)
_ = C6910
tmp21403 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
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
return PrimCons(symnumber, tmp21404)
}
__typedArg0 := symnumber
__typedArg1 := tmp21404
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21406 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21405, Nil)
}
__typedArg0 := tmp21405
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

__e.TailApply(PrimFunc(symis_b), V6911, tmp21408, B6907, L6908, Key6909, C6910)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21409 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21410 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_5_a, tmp21402, tmp21409)


tmp21411 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21410)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21410
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21411

tmp21412 := MakeNative(func(__e *ControlFlow) {
V6916 := __e.Get(1)
_ = V6916
__e.Return(MakeNative(func(__e *ControlFlow) {
B6912 := __e.Get(1)
_ = B6912
__e.Return(MakeNative(func(__e *ControlFlow) {
L6913 := __e.Get(1)
_ = L6913
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6914 := __e.Get(1)
_ = Key6914
__e.Return(MakeNative(func(__e *ControlFlow) {
C6915 := __e.Get(1)
_ = C6915
tmp21413 := Call(__e, PrimFunc(symshen_4newpv), B6912)


let__18990 := tmp21413
_ = let__18990

tmp21414 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21415 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21414)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21414
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21416 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18990, tmp21415)
}
__typedArg0 := let__18990
__typedArg1 := tmp21415
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21417 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21416, Nil)
}
__typedArg0 := tmp21416
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21418 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21417)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21417
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21419 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18990, tmp21418)
}
__typedArg0 := let__18990
__typedArg1 := tmp21418
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21420 := Call(__e, PrimFunc(symis_b), V6916, tmp21419, B6912, L6913, Key6914, C6915)


__e.TailApply(PrimFunc(symshen_4gc), B6912, tmp21420)
return


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

tmp21422 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_a, tmp21412, tmp21421)


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
V6921 := __e.Get(1)
_ = V6921
__e.Return(MakeNative(func(__e *ControlFlow) {
B6917 := __e.Get(1)
_ = B6917
__e.Return(MakeNative(func(__e *ControlFlow) {
L6918 := __e.Get(1)
_ = L6918
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6919 := __e.Get(1)
_ = Key6919
__e.Return(MakeNative(func(__e *ControlFlow) {
C6920 := __e.Get(1)
_ = C6920
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

__e.TailApply(PrimFunc(symis_b), V6921, tmp21430, B6917, L6918, Key6919, C6920)
return


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

tmp21432 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_7, tmp21424, tmp21431)


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
V6926 := __e.Get(1)
_ = V6926
__e.Return(MakeNative(func(__e *ControlFlow) {
B6922 := __e.Get(1)
_ = B6922
__e.Return(MakeNative(func(__e *ControlFlow) {
L6923 := __e.Get(1)
_ = L6923
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6924 := __e.Get(1)
_ = Key6924
__e.Return(MakeNative(func(__e *ControlFlow) {
C6925 := __e.Get(1)
_ = C6925
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

__e.TailApply(PrimFunc(symis_b), V6926, tmp21440, B6922, L6923, Key6924, C6925)
return


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

tmp21442 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_c, tmp21434, tmp21441)


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
V6931 := __e.Get(1)
_ = V6931
__e.Return(MakeNative(func(__e *ControlFlow) {
B6927 := __e.Get(1)
_ = B6927
__e.Return(MakeNative(func(__e *ControlFlow) {
L6928 := __e.Get(1)
_ = L6928
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6929 := __e.Get(1)
_ = Key6929
__e.Return(MakeNative(func(__e *ControlFlow) {
C6930 := __e.Get(1)
_ = C6930
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

__e.TailApply(PrimFunc(symis_b), V6931, tmp21450, B6927, L6928, Key6929, C6930)
return


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

tmp21452 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_1, tmp21444, tmp21451)


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
V6936 := __e.Get(1)
_ = V6936
__e.Return(MakeNative(func(__e *ControlFlow) {
B6932 := __e.Get(1)
_ = B6932
__e.Return(MakeNative(func(__e *ControlFlow) {
L6933 := __e.Get(1)
_ = L6933
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6934 := __e.Get(1)
_ = Key6934
__e.Return(MakeNative(func(__e *ControlFlow) {
C6935 := __e.Get(1)
_ = C6935
tmp21455 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21456 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21455)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21455
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21457 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp21456)
}
__typedArg0 := symnumber
__typedArg1 := tmp21456
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21458 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21457, Nil)
}
__typedArg0 := tmp21457
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

__e.TailApply(PrimFunc(symis_b), V6936, tmp21460, B6932, L6933, Key6934, C6935)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp21461 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21462 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_d, tmp21454, tmp21461)


tmp21463 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21462)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21462
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp21463

tmp21464 := MakeNative(func(__e *ControlFlow) {
V6941 := __e.Get(1)
_ = V6941
__e.Return(MakeNative(func(__e *ControlFlow) {
B6937 := __e.Get(1)
_ = B6937
__e.Return(MakeNative(func(__e *ControlFlow) {
L6938 := __e.Get(1)
_ = L6938
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6939 := __e.Get(1)
_ = Key6939
__e.Return(MakeNative(func(__e *ControlFlow) {
C6940 := __e.Get(1)
_ = C6940
tmp21465 := Call(__e, PrimFunc(symshen_4newpv), B6937)


let__18991 := tmp21465
_ = let__18991

tmp21466 := Call(__e, PrimFunc(symshen_4newpv), B6937)


let__18992 := tmp21466
_ = let__18992

tmp21467 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21468 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21467)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21467
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21469 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18992, tmp21468)
}
__typedArg0 := let__18992
__typedArg1 := tmp21468
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21470 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21469, Nil)
}
__typedArg0 := tmp21469
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21471 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp21470)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp21470
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21472 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18991, tmp21471)
}
__typedArg0 := let__18991
__typedArg1 := tmp21471
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21473 := Call(__e, PrimFunc(symis_b), V6941, tmp21472, B6937, L6938, Key6939, C6940)


tmp21474 := Call(__e, PrimFunc(symshen_4gc), B6937, tmp21473)


__e.TailApply(PrimFunc(symshen_4gc), B6937, tmp21474)
return


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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp21476 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_a_a, tmp21464, tmp21475)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp21476)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp21476
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return


}, 0)

tmp21477 := Call(__e, ns2_1set, symshen_4initialise_1signedfuncs, tmp19758)


_ = tmp21477

tmp21478 := MakeNative(func(__e *ControlFlow) {
tmp21479 := MakeNative(func(__e *ControlFlow) {
Y954 := __e.Get(1)
_ = Y954
__e.TailApply(PrimFunc(symshen_4tuple), Y954)
return
}, 1)

tmp21480 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4tuple, tmp21479)
}
__typedArg0 := symshen_4tuple
__typedArg1 := tmp21479
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21481 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21480)


_ = tmp21481

tmp21482 := MakeNative(func(__e *ControlFlow) {
Y955 := __e.Get(1)
_ = Y955
__e.TailApply(PrimFunc(symshen_4pvar), Y955)
return
}, 1)

tmp21483 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4pvar, tmp21482)
}
__typedArg0 := symshen_4pvar
__typedArg1 := tmp21482
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21484 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21483)


_ = tmp21484

tmp21485 := MakeNative(func(__e *ControlFlow) {
Y956 := __e.Get(1)
_ = Y956
__e.TailApply(PrimFunc(symshen_4dictionary), Y956)
return
}, 1)

tmp21486 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4dictionary, tmp21485)
}
__typedArg0 := symshen_4dictionary
__typedArg1 := tmp21485
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21487 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21486)


_ = tmp21487

tmp21488 := MakeNative(func(__e *ControlFlow) {
Y957 := __e.Get(1)
_ = Y957
__e.TailApply(PrimFunc(symshen_4print_1prolog_1vector), Y957)
return
}, 1)

tmp21489 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4print_1prolog_1vector, tmp21488)
}
__typedArg0 := symshen_4print_1prolog_1vector
__typedArg1 := tmp21488
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21490 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21489)


_ = tmp21490

tmp21491 := MakeNative(func(__e *ControlFlow) {
Y958 := __e.Get(1)
_ = Y958
__e.TailApply(PrimFunc(symshen_4print_1freshterm), Y958)
return
}, 1)

tmp21492 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4print_1freshterm, tmp21491)
}
__typedArg0 := symshen_4print_1freshterm
__typedArg1 := tmp21491
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21493 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21492)


_ = tmp21493

tmp21494 := MakeNative(func(__e *ControlFlow) {
Y959 := __e.Get(1)
_ = Y959
__e.TailApply(PrimFunc(symshen_4printF), Y959)
return
}, 1)

tmp21495 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4printF, tmp21494)
}
__typedArg0 := symshen_4printF
__typedArg1 := tmp21494
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21496 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21495)


_ = tmp21496

tmp21497 := MakeNative(func(__e *ControlFlow) {
Y960 := __e.Get(1)
_ = Y960
__e.TailApply(PrimFunc(symabsolute), Y960)
return
}, 1)

tmp21498 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symabsolute, tmp21497)
}
__typedArg0 := symabsolute
__typedArg1 := tmp21497
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21499 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21498)


_ = tmp21499

tmp21500 := MakeNative(func(__e *ControlFlow) {
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

tmp21501 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symabsvector_2, tmp21500)
}
__typedArg0 := symabsvector_2
__typedArg1 := tmp21500
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21502 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21501)


_ = tmp21502

tmp21503 := MakeNative(func(__e *ControlFlow) {
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

tmp21504 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symabsvector, tmp21503)
}
__typedArg0 := symabsvector
__typedArg1 := tmp21503
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21505 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21504)


_ = tmp21505

tmp21506 := MakeNative(func(__e *ControlFlow) {
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

tmp21507 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symaddress_1_6, tmp21506)
}
__typedArg0 := symaddress_1_6
__typedArg1 := tmp21506
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21508 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21507)


_ = tmp21508

tmp21509 := MakeNative(func(__e *ControlFlow) {
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

tmp21510 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symadjoin, tmp21509)
}
__typedArg0 := symadjoin
__typedArg1 := tmp21509
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21511 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21510)


_ = tmp21511

tmp21512 := MakeNative(func(__e *ControlFlow) {
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

tmp21515 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symand, tmp21512)
}
__typedArg0 := symand
__typedArg1 := tmp21512
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21516 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21515)


_ = tmp21516

tmp21517 := MakeNative(func(__e *ControlFlow) {
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

tmp21518 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symappend, tmp21517)
}
__typedArg0 := symappend
__typedArg1 := tmp21517
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21519 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21518)


_ = tmp21519

tmp21520 := MakeNative(func(__e *ControlFlow) {
Y972 := __e.Get(1)
_ = Y972
__e.TailApply(PrimFunc(symarity), Y972)
return
}, 1)

tmp21521 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symarity, tmp21520)
}
__typedArg0 := symarity
__typedArg1 := tmp21520
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21522 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21521)


_ = tmp21522

tmp21523 := MakeNative(func(__e *ControlFlow) {
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

tmp21524 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symassoc, tmp21523)
}
__typedArg0 := symassoc
__typedArg1 := tmp21523
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21525 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21524)


_ = tmp21525

tmp21526 := MakeNative(func(__e *ControlFlow) {
Y975 := __e.Get(1)
_ = Y975
__e.TailApply(PrimFunc(symatom_2), Y975)
return
}, 1)

tmp21527 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symatom_2, tmp21526)
}
__typedArg0 := symatom_2
__typedArg1 := tmp21526
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21528 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21527)


_ = tmp21528

tmp21529 := MakeNative(func(__e *ControlFlow) {
Y976 := __e.Get(1)
_ = Y976
__e.TailApply(PrimFunc(symboolean_2), Y976)
return
}, 1)

tmp21530 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean_2, tmp21529)
}
__typedArg0 := symboolean_2
__typedArg1 := tmp21529
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21531 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21530)


_ = tmp21531

tmp21532 := MakeNative(func(__e *ControlFlow) {
Y977 := __e.Get(1)
_ = Y977
__e.TailApply(PrimFunc(symbootstrap), Y977)
return
}, 1)

tmp21533 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symbootstrap, tmp21532)
}
__typedArg0 := symbootstrap
__typedArg1 := tmp21532
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21534 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21533)


_ = tmp21534

tmp21535 := MakeNative(func(__e *ControlFlow) {
Y978 := __e.Get(1)
_ = Y978
__e.TailApply(PrimFunc(symbound_2), Y978)
return
}, 1)

tmp21536 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symbound_2, tmp21535)
}
__typedArg0 := symbound_2
__typedArg1 := tmp21535
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21537 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21536)


_ = tmp21537

tmp21538 := MakeNative(func(__e *ControlFlow) {
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

tmp21539 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symbind, tmp21538)
}
__typedArg0 := symbind
__typedArg1 := tmp21538
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21540 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21539)


_ = tmp21540

tmp21541 := MakeNative(func(__e *ControlFlow) {
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

tmp21542 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcall, tmp21541)
}
__typedArg0 := symcall
__typedArg1 := tmp21541
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21543 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21542)


_ = tmp21543

tmp21544 := MakeNative(func(__e *ControlFlow) {
Y990 := __e.Get(1)
_ = Y990
__e.TailApply(PrimFunc(symcd), Y990)
return
}, 1)

tmp21545 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcd, tmp21544)
}
__typedArg0 := symcd
__typedArg1 := tmp21544
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21546 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21545)


_ = tmp21546

tmp21547 := MakeNative(func(__e *ControlFlow) {
Y991 := __e.Get(1)
_ = Y991
__e.TailApply(PrimFunc(symchar_1stream_2), Y991)
return
}, 1)

tmp21548 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symchar_1stream_2, tmp21547)
}
__typedArg0 := symchar_1stream_2
__typedArg1 := tmp21547
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21549 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21548)


_ = tmp21549

tmp21550 := MakeNative(func(__e *ControlFlow) {
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

tmp21551 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcompile, tmp21550)
}
__typedArg0 := symcompile
__typedArg1 := tmp21550
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21552 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21551)


_ = tmp21552

tmp21553 := MakeNative(func(__e *ControlFlow) {
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

tmp21554 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symconcat, tmp21553)
}
__typedArg0 := symconcat
__typedArg1 := tmp21553
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21555 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21554)


_ = tmp21555

tmp21556 := MakeNative(func(__e *ControlFlow) {
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

tmp21557 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp21556)
}
__typedArg0 := symcons
__typedArg1 := tmp21556
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21558 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21557)


_ = tmp21558

tmp21559 := MakeNative(func(__e *ControlFlow) {
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

tmp21560 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons_2, tmp21559)
}
__typedArg0 := symcons_2
__typedArg1 := tmp21559
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21561 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21560)


_ = tmp21561

tmp21562 := MakeNative(func(__e *ControlFlow) {
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

tmp21563 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcn, tmp21562)
}
__typedArg0 := symcn
__typedArg1 := tmp21562
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21564 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21563)


_ = tmp21564

tmp21565 := MakeNative(func(__e *ControlFlow) {
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

tmp21566 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symclose, tmp21565)
}
__typedArg0 := symclose
__typedArg1 := tmp21565
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21567 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21566)


_ = tmp21567

tmp21568 := MakeNative(func(__e *ControlFlow) {
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

tmp21569 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdeclare, tmp21568)
}
__typedArg0 := symdeclare
__typedArg1 := tmp21568
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21570 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21569)


_ = tmp21570

tmp21571 := MakeNative(func(__e *ControlFlow) {
Y1004 := __e.Get(1)
_ = Y1004
__e.TailApply(PrimFunc(symdestroy), Y1004)
return
}, 1)

tmp21572 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdestroy, tmp21571)
}
__typedArg0 := symdestroy
__typedArg1 := tmp21571
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21573 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21572)


_ = tmp21573

tmp21574 := MakeNative(func(__e *ControlFlow) {
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

tmp21575 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdifference, tmp21574)
}
__typedArg0 := symdifference
__typedArg1 := tmp21574
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21576 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21575)


_ = tmp21576

tmp21577 := MakeNative(func(__e *ControlFlow) {
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

tmp21578 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdo, tmp21577)
}
__typedArg0 := symdo
__typedArg1 := tmp21577
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21579 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21578)


_ = tmp21579

tmp21580 := MakeNative(func(__e *ControlFlow) {
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

tmp21581 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symelement_2, tmp21580)
}
__typedArg0 := symelement_2
__typedArg1 := tmp21580
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21582 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21581)


_ = tmp21582

tmp21583 := MakeNative(func(__e *ControlFlow) {
Y1011 := __e.Get(1)
_ = Y1011
__e.TailApply(PrimFunc(symempty_2), Y1011)
return
}, 1)

tmp21584 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symempty_2, tmp21583)
}
__typedArg0 := symempty_2
__typedArg1 := tmp21583
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21585 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21584)


_ = tmp21585

tmp21586 := MakeNative(func(__e *ControlFlow) {
Y1012 := __e.Get(1)
_ = Y1012
__e.TailApply(PrimFunc(symenable_1type_1theory), Y1012)
return
}, 1)

tmp21587 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symenable_1type_1theory, tmp21586)
}
__typedArg0 := symenable_1type_1theory
__typedArg1 := tmp21586
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21588 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21587)


_ = tmp21588

tmp21589 := MakeNative(func(__e *ControlFlow) {
Y1013 := __e.Get(1)
_ = Y1013
__e.TailApply(PrimFunc(symexternal), Y1013)
return
}, 1)

tmp21590 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symexternal, tmp21589)
}
__typedArg0 := symexternal
__typedArg1 := tmp21589
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21591 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21590)


_ = tmp21591

tmp21592 := MakeNative(func(__e *ControlFlow) {
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

tmp21593 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symerror_1to_1string, tmp21592)
}
__typedArg0 := symerror_1to_1string
__typedArg1 := tmp21592
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21594 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21593)


_ = tmp21594

tmp21595 := MakeNative(func(__e *ControlFlow) {
Y1015 := __e.Get(1)
_ = Y1015
__e.TailApply(PrimFunc(symeval), Y1015)
return
}, 1)

tmp21596 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symeval, tmp21595)
}
__typedArg0 := symeval
__typedArg1 := tmp21595
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21597 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21596)


_ = tmp21597

tmp21598 := MakeNative(func(__e *ControlFlow) {
Y1016 := __e.Get(1)
_ = Y1016
__e.TailApply(PrimFunc(symeval_1kl), Y1016)
return
}, 1)

tmp21599 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symeval_1kl, tmp21598)
}
__typedArg0 := symeval_1kl
__typedArg1 := tmp21598
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21600 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21599)


_ = tmp21600

tmp21601 := MakeNative(func(__e *ControlFlow) {
Y1017 := __e.Get(1)
_ = Y1017
__e.TailApply(PrimFunc(symexplode), Y1017)
return
}, 1)

tmp21602 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symexplode, tmp21601)
}
__typedArg0 := symexplode
__typedArg1 := tmp21601
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21603 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21602)


_ = tmp21603

tmp21604 := MakeNative(func(__e *ControlFlow) {
Y1018 := __e.Get(1)
_ = Y1018
__e.TailApply(PrimFunc(symexternal), Y1018)
return
}, 1)

tmp21605 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symexternal, tmp21604)
}
__typedArg0 := symexternal
__typedArg1 := tmp21604
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21606 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21605)


_ = tmp21606

tmp21607 := MakeNative(func(__e *ControlFlow) {
Y1019 := __e.Get(1)
_ = Y1019
__e.TailApply(PrimFunc(symfactorise), Y1019)
return
}, 1)

tmp21608 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfactorise, tmp21607)
}
__typedArg0 := symfactorise
__typedArg1 := tmp21607
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21609 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21608)


_ = tmp21609

tmp21610 := MakeNative(func(__e *ControlFlow) {
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

tmp21611 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfail_1if, tmp21610)
}
__typedArg0 := symfail_1if
__typedArg1 := tmp21610
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21612 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21611)


_ = tmp21612

tmp21613 := MakeNative(func(__e *ControlFlow) {
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

tmp21614 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfix, tmp21613)
}
__typedArg0 := symfix
__typedArg1 := tmp21613
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21615 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21614)


_ = tmp21615

tmp21616 := MakeNative(func(__e *ControlFlow) {
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

tmp21617 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfindall, tmp21616)
}
__typedArg0 := symfindall
__typedArg1 := tmp21616
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21618 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21617)


_ = tmp21618

tmp21619 := MakeNative(func(__e *ControlFlow) {
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

tmp21620 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfork, tmp21619)
}
__typedArg0 := symfork
__typedArg1 := tmp21619
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21621 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21620)


_ = tmp21621

tmp21622 := MakeNative(func(__e *ControlFlow) {
Y1036 := __e.Get(1)
_ = Y1036
__e.Return(MakeNative(func(__e *ControlFlow) {
__e.Return(Y1036)
return
}, 0))
return
}, 1)

tmp21623 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfreeze, tmp21622)
}
__typedArg0 := symfreeze
__typedArg1 := tmp21622
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21624 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21623)


_ = tmp21624

tmp21625 := MakeNative(func(__e *ControlFlow) {
Y1037 := __e.Get(1)
_ = Y1037
__e.TailApply(PrimFunc(symfst), Y1037)
return
}, 1)

tmp21626 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfst, tmp21625)
}
__typedArg0 := symfst
__typedArg1 := tmp21625
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21627 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21626)


_ = tmp21627

tmp21628 := MakeNative(func(__e *ControlFlow) {
Y1038 := __e.Get(1)
_ = Y1038
__e.TailApply(PrimFunc(symfn), Y1038)
return
}, 1)

tmp21629 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfn, tmp21628)
}
__typedArg0 := symfn
__typedArg1 := tmp21628
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21630 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21629)


_ = tmp21630

tmp21631 := MakeNative(func(__e *ControlFlow) {
Y1039 := __e.Get(1)
_ = Y1039
__e.TailApply(PrimFunc(symfunction), Y1039)
return
}, 1)

tmp21632 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfunction, tmp21631)
}
__typedArg0 := symfunction
__typedArg1 := tmp21631
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21633 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21632)


_ = tmp21633

tmp21634 := MakeNative(func(__e *ControlFlow) {
Y1040 := __e.Get(1)
_ = Y1040
__e.TailApply(PrimFunc(symgensym), Y1040)
return
}, 1)

tmp21635 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symgensym, tmp21634)
}
__typedArg0 := symgensym
__typedArg1 := tmp21634
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21636 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21635)


_ = tmp21636

tmp21637 := MakeNative(func(__e *ControlFlow) {
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

tmp21638 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symget, tmp21637)
}
__typedArg0 := symget
__typedArg1 := tmp21637
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21639 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21638)


_ = tmp21639

tmp21640 := MakeNative(func(__e *ControlFlow) {
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

tmp21641 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symget_1time, tmp21640)
}
__typedArg0 := symget_1time
__typedArg1 := tmp21640
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21642 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21641)


_ = tmp21642

tmp21643 := MakeNative(func(__e *ControlFlow) {
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

tmp21644 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symaddress_1_6, tmp21643)
}
__typedArg0 := symaddress_1_6
__typedArg1 := tmp21643
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21645 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21644)


_ = tmp21645

tmp21646 := MakeNative(func(__e *ControlFlow) {
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

tmp21647 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1address, tmp21646)
}
__typedArg0 := sym_5_1address
__typedArg1 := tmp21646
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21648 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21647)


_ = tmp21648

tmp21649 := MakeNative(func(__e *ControlFlow) {
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

tmp21650 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1vector, tmp21649)
}
__typedArg0 := sym_5_1vector
__typedArg1 := tmp21649
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21651 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21650)


_ = tmp21651

tmp21652 := MakeNative(func(__e *ControlFlow) {
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

tmp21653 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_6, tmp21652)
}
__typedArg0 := sym_6
__typedArg1 := tmp21652
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21654 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21653)


_ = tmp21654

tmp21655 := MakeNative(func(__e *ControlFlow) {
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

tmp21656 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_6_a, tmp21655)
}
__typedArg0 := sym_6_a
__typedArg1 := tmp21655
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21657 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21656)


_ = tmp21657

tmp21658 := MakeNative(func(__e *ControlFlow) {
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

tmp21659 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_a, tmp21658)
}
__typedArg0 := sym_a
__typedArg1 := tmp21658
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21660 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21659)


_ = tmp21660

tmp21661 := MakeNative(func(__e *ControlFlow) {
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

tmp21662 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhash, tmp21661)
}
__typedArg0 := symhash
__typedArg1 := tmp21661
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21663 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21662)


_ = tmp21663

tmp21664 := MakeNative(func(__e *ControlFlow) {
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

tmp21665 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhd, tmp21664)
}
__typedArg0 := symhd
__typedArg1 := tmp21664
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21666 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21665)


_ = tmp21666

tmp21667 := MakeNative(func(__e *ControlFlow) {
Y1061 := __e.Get(1)
_ = Y1061
__e.TailApply(PrimFunc(symhdv), Y1061)
return
}, 1)

tmp21668 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhdv, tmp21667)
}
__typedArg0 := symhdv
__typedArg1 := tmp21667
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21669 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21668)


_ = tmp21669

tmp21670 := MakeNative(func(__e *ControlFlow) {
Y1062 := __e.Get(1)
_ = Y1062
__e.TailApply(PrimFunc(symhdstr), Y1062)
return
}, 1)

tmp21671 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhdstr, tmp21670)
}
__typedArg0 := symhdstr
__typedArg1 := tmp21670
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21672 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21671)


_ = tmp21672

tmp21673 := MakeNative(func(__e *ControlFlow) {
Y1063 := __e.Get(1)
_ = Y1063
__e.TailApply(PrimFunc(symhead), Y1063)
return
}, 1)

tmp21674 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhead, tmp21673)
}
__typedArg0 := symhead
__typedArg1 := tmp21673
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21675 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21674)


_ = tmp21675

tmp21676 := MakeNative(func(__e *ControlFlow) {
Y1064 := __e.Get(1)
_ = Y1064
__e.TailApply(PrimFunc(symhush), Y1064)
return
}, 1)

tmp21677 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhush, tmp21676)
}
__typedArg0 := symhush
__typedArg1 := tmp21676
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21678 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21677)


_ = tmp21678

tmp21679 := MakeNative(func(__e *ControlFlow) {
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

tmp21681 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp21679)
}
__typedArg0 := symif
__typedArg1 := tmp21679
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21682 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21681)


_ = tmp21682

tmp21683 := MakeNative(func(__e *ControlFlow) {
Y1068 := __e.Get(1)
_ = Y1068
__e.TailApply(PrimFunc(syminclude), Y1068)
return
}, 1)

tmp21684 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminclude, tmp21683)
}
__typedArg0 := syminclude
__typedArg1 := tmp21683
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21685 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21684)


_ = tmp21685

tmp21686 := MakeNative(func(__e *ControlFlow) {
Y1069 := __e.Get(1)
_ = Y1069
__e.TailApply(PrimFunc(symin_1package), Y1069)
return
}, 1)

tmp21687 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symin_1package, tmp21686)
}
__typedArg0 := symin_1package
__typedArg1 := tmp21686
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21688 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21687)


_ = tmp21688

tmp21689 := MakeNative(func(__e *ControlFlow) {
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

tmp21690 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminteger_2, tmp21689)
}
__typedArg0 := syminteger_2
__typedArg1 := tmp21689
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21691 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21690)


_ = tmp21691

tmp21692 := MakeNative(func(__e *ControlFlow) {
Y1071 := __e.Get(1)
_ = Y1071
__e.TailApply(PrimFunc(syminternal), Y1071)
return
}, 1)

tmp21693 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminternal, tmp21692)
}
__typedArg0 := syminternal
__typedArg1 := tmp21692
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21694 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21693)


_ = tmp21694

tmp21695 := MakeNative(func(__e *ControlFlow) {
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

tmp21696 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symintern, tmp21695)
}
__typedArg0 := symintern
__typedArg1 := tmp21695
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21697 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21696)


_ = tmp21697

tmp21698 := MakeNative(func(__e *ControlFlow) {
Y1073 := __e.Get(1)
_ = Y1073
__e.TailApply(PrimFunc(syminput), Y1073)
return
}, 1)

tmp21699 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminput, tmp21698)
}
__typedArg0 := syminput
__typedArg1 := tmp21698
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21700 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21699)


_ = tmp21700

tmp21701 := MakeNative(func(__e *ControlFlow) {
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

tmp21702 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminput_7, tmp21701)
}
__typedArg0 := syminput_7
__typedArg1 := tmp21701
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21703 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21702)


_ = tmp21703

tmp21704 := MakeNative(func(__e *ControlFlow) {
Y1076 := __e.Get(1)
_ = Y1076
__e.TailApply(PrimFunc(syminclude_1all_1but), Y1076)
return
}, 1)

tmp21705 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminclude_1all_1but, tmp21704)
}
__typedArg0 := syminclude_1all_1but
__typedArg1 := tmp21704
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21706 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21705)


_ = tmp21706

tmp21707 := MakeNative(func(__e *ControlFlow) {
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

tmp21708 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symintersection, tmp21707)
}
__typedArg0 := symintersection
__typedArg1 := tmp21707
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21709 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21708)


_ = tmp21709

tmp21710 := MakeNative(func(__e *ControlFlow) {
Y1079 := __e.Get(1)
_ = Y1079
__e.TailApply(PrimFunc(syminternal), Y1079)
return
}, 1)

tmp21711 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminternal, tmp21710)
}
__typedArg0 := syminternal
__typedArg1 := tmp21710
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21712 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21711)


_ = tmp21712

tmp21713 := MakeNative(func(__e *ControlFlow) {
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

tmp21714 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symis, tmp21713)
}
__typedArg0 := symis
__typedArg1 := tmp21713
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21715 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21714)


_ = tmp21715

tmp21716 := MakeNative(func(__e *ControlFlow) {
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

tmp21717 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symis_b, tmp21716)
}
__typedArg0 := symis_b
__typedArg1 := tmp21716
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21718 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21717)


_ = tmp21718

tmp21719 := MakeNative(func(__e *ControlFlow) {
Y1092 := __e.Get(1)
_ = Y1092
__e.TailApply(PrimFunc(symlength), Y1092)
return
}, 1)

tmp21720 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlength, tmp21719)
}
__typedArg0 := symlength
__typedArg1 := tmp21719
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21721 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21720)


_ = tmp21721

tmp21722 := MakeNative(func(__e *ControlFlow) {
Y1093 := __e.Get(1)
_ = Y1093
__e.TailApply(PrimFunc(symlimit), Y1093)
return
}, 1)

tmp21723 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlimit, tmp21722)
}
__typedArg0 := symlimit
__typedArg1 := tmp21722
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21724 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21723)


_ = tmp21724

tmp21725 := MakeNative(func(__e *ControlFlow) {
Y1094 := __e.Get(1)
_ = Y1094
__e.TailApply(PrimFunc(symlineread), Y1094)
return
}, 1)

tmp21726 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlineread, tmp21725)
}
__typedArg0 := symlineread
__typedArg1 := tmp21725
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21727 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21726)


_ = tmp21727

tmp21728 := MakeNative(func(__e *ControlFlow) {
Y1095 := __e.Get(1)
_ = Y1095
__e.TailApply(PrimFunc(symload), Y1095)
return
}, 1)

tmp21729 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symload, tmp21728)
}
__typedArg0 := symload
__typedArg1 := tmp21728
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21730 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21729)


_ = tmp21730

tmp21731 := MakeNative(func(__e *ControlFlow) {
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

tmp21732 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5, tmp21731)
}
__typedArg0 := sym_5
__typedArg1 := tmp21731
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21733 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21732)


_ = tmp21733

tmp21734 := MakeNative(func(__e *ControlFlow) {
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

tmp21735 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_a, tmp21734)
}
__typedArg0 := sym_5_a
__typedArg1 := tmp21734
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21736 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21735)


_ = tmp21736

tmp21737 := MakeNative(func(__e *ControlFlow) {
Y1100 := __e.Get(1)
_ = Y1100
__e.TailApply(PrimFunc(symvector), Y1100)
return
}, 1)

tmp21738 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp21737)
}
__typedArg0 := symvector
__typedArg1 := tmp21737
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21739 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21738)


_ = tmp21739

tmp21740 := MakeNative(func(__e *ControlFlow) {
Y1101 := __e.Get(1)
_ = Y1101
__e.TailApply(PrimFunc(symmacroexpand), Y1101)
return
}, 1)

tmp21741 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symmacroexpand, tmp21740)
}
__typedArg0 := symmacroexpand
__typedArg1 := tmp21740
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21742 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21741)


_ = tmp21742

tmp21743 := MakeNative(func(__e *ControlFlow) {
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

tmp21744 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symmap, tmp21743)
}
__typedArg0 := symmap
__typedArg1 := tmp21743
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21745 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21744)


_ = tmp21745

tmp21746 := MakeNative(func(__e *ControlFlow) {
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

tmp21747 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symmapcan, tmp21746)
}
__typedArg0 := symmapcan
__typedArg1 := tmp21746
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21748 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21747)


_ = tmp21748

tmp21749 := MakeNative(func(__e *ControlFlow) {
Y1106 := __e.Get(1)
_ = Y1106
__e.TailApply(PrimFunc(symmaxinferences), Y1106)
return
}, 1)

tmp21750 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symmaxinferences, tmp21749)
}
__typedArg0 := symmaxinferences
__typedArg1 := tmp21749
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21751 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21750)


_ = tmp21751

tmp21752 := MakeNative(func(__e *ControlFlow) {
Y1107 := __e.Get(1)
_ = Y1107
__e.TailApply(PrimFunc(symnl), Y1107)
return
}, 1)

tmp21753 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnl, tmp21752)
}
__typedArg0 := symnl
__typedArg1 := tmp21752
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21754 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21753)


_ = tmp21754

tmp21755 := MakeNative(func(__e *ControlFlow) {
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

tmp21756 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnot, tmp21755)
}
__typedArg0 := symnot
__typedArg1 := tmp21755
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21757 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21756)


_ = tmp21757

tmp21758 := MakeNative(func(__e *ControlFlow) {
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

tmp21759 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnth, tmp21758)
}
__typedArg0 := symnth
__typedArg1 := tmp21758
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21760 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21759)


_ = tmp21760

tmp21761 := MakeNative(func(__e *ControlFlow) {
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

tmp21762 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symn_1_6string, tmp21761)
}
__typedArg0 := symn_1_6string
__typedArg1 := tmp21761
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21763 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21762)


_ = tmp21763

tmp21764 := MakeNative(func(__e *ControlFlow) {
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

tmp21765 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber_2, tmp21764)
}
__typedArg0 := symnumber_2
__typedArg1 := tmp21764
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21766 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21765)


_ = tmp21766

tmp21767 := MakeNative(func(__e *ControlFlow) {
Y1113 := __e.Get(1)
_ = Y1113
__e.TailApply(PrimFunc(symoccurs_1check), Y1113)
return
}, 1)

tmp21768 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symoccurs_1check, tmp21767)
}
__typedArg0 := symoccurs_1check
__typedArg1 := tmp21767
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21769 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21768)


_ = tmp21769

tmp21770 := MakeNative(func(__e *ControlFlow) {
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

tmp21771 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symoccurrences, tmp21770)
}
__typedArg0 := symoccurrences
__typedArg1 := tmp21770
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21772 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21771)


_ = tmp21772

tmp21773 := MakeNative(func(__e *ControlFlow) {
Y1116 := __e.Get(1)
_ = Y1116
__e.TailApply(PrimFunc(symoccurs_1check), Y1116)
return
}, 1)

tmp21774 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symoccurs_1check, tmp21773)
}
__typedArg0 := symoccurs_1check
__typedArg1 := tmp21773
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21775 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21774)


_ = tmp21775

tmp21776 := MakeNative(func(__e *ControlFlow) {
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

tmp21777 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symopen, tmp21776)
}
__typedArg0 := symopen
__typedArg1 := tmp21776
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21778 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21777)


_ = tmp21778

tmp21779 := MakeNative(func(__e *ControlFlow) {
Y1119 := __e.Get(1)
_ = Y1119
__e.TailApply(PrimFunc(symoptimise), Y1119)
return
}, 1)

tmp21780 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symoptimise, tmp21779)
}
__typedArg0 := symoptimise
__typedArg1 := tmp21779
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21781 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21780)


_ = tmp21781

tmp21782 := MakeNative(func(__e *ControlFlow) {
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

tmp21785 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symor, tmp21782)
}
__typedArg0 := symor
__typedArg1 := tmp21782
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21786 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21785)


_ = tmp21786

tmp21787 := MakeNative(func(__e *ControlFlow) {
Y1122 := __e.Get(1)
_ = Y1122
__e.TailApply(PrimFunc(sympackage_2), Y1122)
return
}, 1)

tmp21788 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympackage_2, tmp21787)
}
__typedArg0 := sympackage_2
__typedArg1 := tmp21787
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21789 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21788)


_ = tmp21789

tmp21790 := MakeNative(func(__e *ControlFlow) {
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

tmp21791 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympos, tmp21790)
}
__typedArg0 := sympos
__typedArg1 := tmp21790
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21792 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21791)


_ = tmp21792

tmp21793 := MakeNative(func(__e *ControlFlow) {
Y1125 := __e.Get(1)
_ = Y1125
__e.TailApply(PrimFunc(sympreclude_1all_1but), Y1125)
return
}, 1)

tmp21794 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympreclude_1all_1but, tmp21793)
}
__typedArg0 := sympreclude_1all_1but
__typedArg1 := tmp21793
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21795 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21794)


_ = tmp21795

tmp21796 := MakeNative(func(__e *ControlFlow) {
Y1126 := __e.Get(1)
_ = Y1126
__e.TailApply(PrimFunc(symprint), Y1126)
return
}, 1)

tmp21797 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symprint, tmp21796)
}
__typedArg0 := symprint
__typedArg1 := tmp21796
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21798 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21797)


_ = tmp21798

tmp21799 := MakeNative(func(__e *ControlFlow) {
Y1127 := __e.Get(1)
_ = Y1127
__e.TailApply(PrimFunc(symprofile), Y1127)
return
}, 1)

tmp21800 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symprofile, tmp21799)
}
__typedArg0 := symprofile
__typedArg1 := tmp21799
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21801 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21800)


_ = tmp21801

tmp21802 := MakeNative(func(__e *ControlFlow) {
Y1128 := __e.Get(1)
_ = Y1128
__e.TailApply(PrimFunc(symshen_4print_1prolog_1vector), Y1128)
return
}, 1)

tmp21803 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4print_1prolog_1vector, tmp21802)
}
__typedArg0 := symshen_4print_1prolog_1vector
__typedArg1 := tmp21802
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21804 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21803)


_ = tmp21804

tmp21805 := MakeNative(func(__e *ControlFlow) {
Y1129 := __e.Get(1)
_ = Y1129
__e.TailApply(PrimFunc(symshen_4print_1freshterm), Y1129)
return
}, 1)

tmp21806 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4print_1freshterm, tmp21805)
}
__typedArg0 := symshen_4print_1freshterm
__typedArg1 := tmp21805
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21807 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21806)


_ = tmp21807

tmp21808 := MakeNative(func(__e *ControlFlow) {
Y1130 := __e.Get(1)
_ = Y1130
__e.TailApply(PrimFunc(symshen_4printF), Y1130)
return
}, 1)

tmp21809 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4printF, tmp21808)
}
__typedArg0 := symshen_4printF
__typedArg1 := tmp21808
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21810 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21809)


_ = tmp21810

tmp21811 := MakeNative(func(__e *ControlFlow) {
Y1131 := __e.Get(1)
_ = Y1131
__e.TailApply(PrimFunc(symprolog_1memory), Y1131)
return
}, 1)

tmp21812 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symprolog_1memory, tmp21811)
}
__typedArg0 := symprolog_1memory
__typedArg1 := tmp21811
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21813 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21812)


_ = tmp21813

tmp21814 := MakeNative(func(__e *ControlFlow) {
Y1132 := __e.Get(1)
_ = Y1132
__e.TailApply(PrimFunc(symprofile_1results), Y1132)
return
}, 1)

tmp21815 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symprofile_1results, tmp21814)
}
__typedArg0 := symprofile_1results
__typedArg1 := tmp21814
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21816 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21815)


_ = tmp21816

tmp21817 := MakeNative(func(__e *ControlFlow) {
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

tmp21818 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympr, tmp21817)
}
__typedArg0 := sympr
__typedArg1 := tmp21817
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21819 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21818)


_ = tmp21819

tmp21820 := MakeNative(func(__e *ControlFlow) {
Y1135 := __e.Get(1)
_ = Y1135
__e.TailApply(PrimFunc(symps), Y1135)
return
}, 1)

tmp21821 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symps, tmp21820)
}
__typedArg0 := symps
__typedArg1 := tmp21820
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21822 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21821)


_ = tmp21822

tmp21823 := MakeNative(func(__e *ControlFlow) {
Y1136 := __e.Get(1)
_ = Y1136
__e.TailApply(PrimFunc(sympreclude), Y1136)
return
}, 1)

tmp21824 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympreclude, tmp21823)
}
__typedArg0 := sympreclude
__typedArg1 := tmp21823
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21825 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21824)


_ = tmp21825

tmp21826 := MakeNative(func(__e *ControlFlow) {
Y1137 := __e.Get(1)
_ = Y1137
__e.TailApply(PrimFunc(sympreclude_1all_1but), Y1137)
return
}, 1)

tmp21827 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympreclude_1all_1but, tmp21826)
}
__typedArg0 := sympreclude_1all_1but
__typedArg1 := tmp21826
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21828 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21827)


_ = tmp21828

tmp21829 := MakeNative(func(__e *ControlFlow) {
Y1138 := __e.Get(1)
_ = Y1138
__e.TailApply(PrimFunc(symprotect), Y1138)
return
}, 1)

tmp21830 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symprotect, tmp21829)
}
__typedArg0 := symprotect
__typedArg1 := tmp21829
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21831 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21830)


_ = tmp21831

tmp21832 := MakeNative(func(__e *ControlFlow) {
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

tmp21833 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symput, tmp21832)
}
__typedArg0 := symput
__typedArg1 := tmp21832
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21834 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21833)


_ = tmp21834

tmp21835 := MakeNative(func(__e *ControlFlow) {
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

tmp21836 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1file_1as_1string, tmp21835)
}
__typedArg0 := symread_1file_1as_1string
__typedArg1 := tmp21835
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21837 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21836)


_ = tmp21837

tmp21838 := MakeNative(func(__e *ControlFlow) {
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

tmp21839 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1file_1as_1bytelist, tmp21838)
}
__typedArg0 := symread_1file_1as_1bytelist
__typedArg1 := tmp21838
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21840 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21839)


_ = tmp21840

tmp21841 := MakeNative(func(__e *ControlFlow) {
Y1145 := __e.Get(1)
_ = Y1145
__e.TailApply(PrimFunc(symread_1file), Y1145)
return
}, 1)

tmp21842 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1file, tmp21841)
}
__typedArg0 := symread_1file
__typedArg1 := tmp21841
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21843 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21842)


_ = tmp21843

tmp21844 := MakeNative(func(__e *ControlFlow) {
Y1146 := __e.Get(1)
_ = Y1146
__e.TailApply(PrimFunc(symread), Y1146)
return
}, 1)

tmp21845 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread, tmp21844)
}
__typedArg0 := symread
__typedArg1 := tmp21844
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21846 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21845)


_ = tmp21846

tmp21847 := MakeNative(func(__e *ControlFlow) {
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

tmp21848 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1byte, tmp21847)
}
__typedArg0 := symread_1byte
__typedArg1 := tmp21847
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21849 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21848)


_ = tmp21849

tmp21850 := MakeNative(func(__e *ControlFlow) {
Y1148 := __e.Get(1)
_ = Y1148
__e.TailApply(PrimFunc(symread_1from_1string), Y1148)
return
}, 1)

tmp21851 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1from_1string, tmp21850)
}
__typedArg0 := symread_1from_1string
__typedArg1 := tmp21850
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21852 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21851)


_ = tmp21852

tmp21853 := MakeNative(func(__e *ControlFlow) {
Y1149 := __e.Get(1)
_ = Y1149
__e.TailApply(PrimFunc(symread_1from_1string_1unprocessed), Y1149)
return
}, 1)

tmp21854 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1from_1string_1unprocessed, tmp21853)
}
__typedArg0 := symread_1from_1string_1unprocessed
__typedArg1 := tmp21853
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21855 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21854)


_ = tmp21855

tmp21856 := MakeNative(func(__e *ControlFlow) {
Y1150 := __e.Get(1)
_ = Y1150
__e.TailApply(PrimFunc(symread_1unit_1string), Y1150)
return
}, 1)

tmp21857 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1unit_1string, tmp21856)
}
__typedArg0 := symread_1unit_1string
__typedArg1 := tmp21856
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21858 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21857)


_ = tmp21858

tmp21859 := MakeNative(func(__e *ControlFlow) {
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

tmp21860 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4register_1source_1form, tmp21859)
}
__typedArg0 := symshen_4register_1source_1form
__typedArg1 := tmp21859
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21861 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21860)


_ = tmp21861

tmp21862 := MakeNative(func(__e *ControlFlow) {
Y1153 := __e.Get(1)
_ = Y1153
__e.TailApply(PrimFunc(symshen_4unregister_1source_1form), Y1153)
return
}, 1)

tmp21863 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4unregister_1source_1form, tmp21862)
}
__typedArg0 := symshen_4unregister_1source_1form
__typedArg1 := tmp21862
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21864 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21863)


_ = tmp21864

tmp21865 := MakeNative(func(__e *ControlFlow) {
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

tmp21866 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symremove, tmp21865)
}
__typedArg0 := symremove
__typedArg1 := tmp21865
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21867 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21866)


_ = tmp21867

tmp21868 := MakeNative(func(__e *ControlFlow) {
Y1156 := __e.Get(1)
_ = Y1156
__e.TailApply(PrimFunc(symreverse), Y1156)
return
}, 1)

tmp21869 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symreverse, tmp21868)
}
__typedArg0 := symreverse
__typedArg1 := tmp21868
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21870 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21869)


_ = tmp21870

tmp21871 := MakeNative(func(__e *ControlFlow) {
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

tmp21872 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symset, tmp21871)
}
__typedArg0 := symset
__typedArg1 := tmp21871
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21873 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21872)


_ = tmp21873

tmp21874 := MakeNative(func(__e *ControlFlow) {
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

tmp21875 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsimple_1error, tmp21874)
}
__typedArg0 := symsimple_1error
__typedArg1 := tmp21874
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21876 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21875)


_ = tmp21876

tmp21877 := MakeNative(func(__e *ControlFlow) {
Y1160 := __e.Get(1)
_ = Y1160
__e.TailApply(PrimFunc(symsnd), Y1160)
return
}, 1)

tmp21878 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsnd, tmp21877)
}
__typedArg0 := symsnd
__typedArg1 := tmp21877
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21879 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21878)


_ = tmp21879

tmp21880 := MakeNative(func(__e *ControlFlow) {
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

tmp21881 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symspecialise, tmp21880)
}
__typedArg0 := symspecialise
__typedArg1 := tmp21880
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21882 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21881)


_ = tmp21882

tmp21883 := MakeNative(func(__e *ControlFlow) {
Y1163 := __e.Get(1)
_ = Y1163
__e.TailApply(PrimFunc(symspy), Y1163)
return
}, 1)

tmp21884 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symspy, tmp21883)
}
__typedArg0 := symspy
__typedArg1 := tmp21883
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21885 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21884)


_ = tmp21885

tmp21886 := MakeNative(func(__e *ControlFlow) {
Y1164 := __e.Get(1)
_ = Y1164
__e.TailApply(PrimFunc(symstep), Y1164)
return
}, 1)

tmp21887 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstep, tmp21886)
}
__typedArg0 := symstep
__typedArg1 := tmp21886
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21888 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21887)


_ = tmp21888

tmp21889 := MakeNative(func(__e *ControlFlow) {
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

tmp21890 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstr, tmp21889)
}
__typedArg0 := symstr
__typedArg1 := tmp21889
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21891 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21890)


_ = tmp21891

tmp21892 := MakeNative(func(__e *ControlFlow) {
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

tmp21893 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring_1_6n, tmp21892)
}
__typedArg0 := symstring_1_6n
__typedArg1 := tmp21892
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21894 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21893)


_ = tmp21894

tmp21895 := MakeNative(func(__e *ControlFlow) {
Y1167 := __e.Get(1)
_ = Y1167
__e.TailApply(PrimFunc(symstring_1_6symbol), Y1167)
return
}, 1)

tmp21896 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring_1_6symbol, tmp21895)
}
__typedArg0 := symstring_1_6symbol
__typedArg1 := tmp21895
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21897 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21896)


_ = tmp21897

tmp21898 := MakeNative(func(__e *ControlFlow) {
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

tmp21899 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring_2, tmp21898)
}
__typedArg0 := symstring_2
__typedArg1 := tmp21898
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21900 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21899)


_ = tmp21900

tmp21901 := MakeNative(func(__e *ControlFlow) {
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

tmp21902 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsubst, tmp21901)
}
__typedArg0 := symsubst
__typedArg1 := tmp21901
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21903 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21902)


_ = tmp21903

tmp21904 := MakeNative(func(__e *ControlFlow) {
Y1172 := __e.Get(1)
_ = Y1172
__e.TailApply(PrimFunc(symsum), Y1172)
return
}, 1)

tmp21905 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsum, tmp21904)
}
__typedArg0 := symsum
__typedArg1 := tmp21904
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21906 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21905)


_ = tmp21906

tmp21907 := MakeNative(func(__e *ControlFlow) {
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

tmp21908 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol_2, tmp21907)
}
__typedArg0 := symsymbol_2
__typedArg1 := tmp21907
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21909 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21908)


_ = tmp21909

tmp21910 := MakeNative(func(__e *ControlFlow) {
Y1174 := __e.Get(1)
_ = Y1174
__e.TailApply(PrimFunc(symsystemf), Y1174)
return
}, 1)

tmp21911 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsystemf, tmp21910)
}
__typedArg0 := symsystemf
__typedArg1 := tmp21910
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21912 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21911)


_ = tmp21912

tmp21913 := MakeNative(func(__e *ControlFlow) {
Y1175 := __e.Get(1)
_ = Y1175
__e.TailApply(PrimFunc(symtail), Y1175)
return
}, 1)

tmp21914 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtail, tmp21913)
}
__typedArg0 := symtail
__typedArg1 := tmp21913
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21915 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21914)


_ = tmp21915

tmp21916 := MakeNative(func(__e *ControlFlow) {
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

tmp21917 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtl, tmp21916)
}
__typedArg0 := symtl
__typedArg1 := tmp21916
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21918 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21917)


_ = tmp21918

tmp21919 := MakeNative(func(__e *ControlFlow) {
Y1177 := __e.Get(1)
_ = Y1177
__e.TailApply(PrimFunc(symtc), Y1177)
return
}, 1)

tmp21920 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtc, tmp21919)
}
__typedArg0 := symtc
__typedArg1 := tmp21919
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21921 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21920)


_ = tmp21921

tmp21922 := MakeNative(func(__e *ControlFlow) {
Y1178 := __e.Get(1)
_ = Y1178
__e.TailApply(PrimFunc(symthaw), Y1178)
return
}, 1)

tmp21923 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symthaw, tmp21922)
}
__typedArg0 := symthaw
__typedArg1 := tmp21922
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21924 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21923)


_ = tmp21924

tmp21925 := MakeNative(func(__e *ControlFlow) {
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

tmp21926 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtlstr, tmp21925)
}
__typedArg0 := symtlstr
__typedArg1 := tmp21925
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21927 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21926)


_ = tmp21927

tmp21928 := MakeNative(func(__e *ControlFlow) {
Y1180 := __e.Get(1)
_ = Y1180
__e.TailApply(PrimFunc(symtrack), Y1180)
return
}, 1)

tmp21929 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtrack, tmp21928)
}
__typedArg0 := symtrack
__typedArg1 := tmp21928
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21930 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21929)


_ = tmp21930

tmp21931 := MakeNative(func(__e *ControlFlow) {
Y1181 := __e.Get(1)
_ = Y1181
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1182 := __e.Get(1)
_ = Y1182
tmp21932 := MakeNative(func(__e *ControlFlow) {
__e.Return(Y1181)
return
}, 0)

__e.TailApply(try_1catch, tmp21932, Y1182)
return


}, 1))
return
}, 1)

tmp21933 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtrap_1error, tmp21931)
}
__typedArg0 := symtrap_1error
__typedArg1 := tmp21931
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21934 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21933)


_ = tmp21934

tmp21935 := MakeNative(func(__e *ControlFlow) {
Y1183 := __e.Get(1)
_ = Y1183
__e.TailApply(PrimFunc(symtuple_2), Y1183)
return
}, 1)

tmp21936 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtuple_2, tmp21935)
}
__typedArg0 := symtuple_2
__typedArg1 := tmp21935
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21937 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21936)


_ = tmp21937

tmp21938 := MakeNative(func(__e *ControlFlow) {
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

tmp21939 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtype, tmp21938)
}
__typedArg0 := symtype
__typedArg1 := tmp21938
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21940 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21939)


_ = tmp21940

tmp21941 := MakeNative(func(__e *ControlFlow) {
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

tmp21942 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symreturn, tmp21941)
}
__typedArg0 := symreturn
__typedArg1 := tmp21941
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21943 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21942)


_ = tmp21943

tmp21944 := MakeNative(func(__e *ControlFlow) {
Y1191 := __e.Get(1)
_ = Y1191
__e.TailApply(PrimFunc(symunabsolute), Y1191)
return
}, 1)

tmp21945 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunabsolute, tmp21944)
}
__typedArg0 := symunabsolute
__typedArg1 := tmp21944
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21946 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21945)


_ = tmp21946

tmp21947 := MakeNative(func(__e *ControlFlow) {
Y1192 := __e.Get(1)
_ = Y1192
__e.TailApply(PrimFunc(symundefmacro), Y1192)
return
}, 1)

tmp21948 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symundefmacro, tmp21947)
}
__typedArg0 := symundefmacro
__typedArg1 := tmp21947
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21949 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21948)


_ = tmp21949

tmp21950 := MakeNative(func(__e *ControlFlow) {
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

tmp21951 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunput, tmp21950)
}
__typedArg0 := symunput
__typedArg1 := tmp21950
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21952 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21951)


_ = tmp21952

tmp21953 := MakeNative(func(__e *ControlFlow) {
Y1196 := __e.Get(1)
_ = Y1196
__e.TailApply(PrimFunc(symunprofile), Y1196)
return
}, 1)

tmp21954 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunprofile, tmp21953)
}
__typedArg0 := symunprofile
__typedArg1 := tmp21953
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21955 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21954)


_ = tmp21955

tmp21956 := MakeNative(func(__e *ControlFlow) {
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

tmp21957 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunion, tmp21956)
}
__typedArg0 := symunion
__typedArg1 := tmp21956
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21958 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21957)


_ = tmp21958

tmp21959 := MakeNative(func(__e *ControlFlow) {
Y1199 := __e.Get(1)
_ = Y1199
__e.TailApply(PrimFunc(symuntrack), Y1199)
return
}, 1)

tmp21960 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symuntrack, tmp21959)
}
__typedArg0 := symuntrack
__typedArg1 := tmp21959
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21961 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21960)


_ = tmp21961

tmp21962 := MakeNative(func(__e *ControlFlow) {
Y1200 := __e.Get(1)
_ = Y1200
__e.TailApply(PrimFunc(symundefmacro), Y1200)
return
}, 1)

tmp21963 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symundefmacro, tmp21962)
}
__typedArg0 := symundefmacro
__typedArg1 := tmp21962
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21964 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21963)


_ = tmp21964

tmp21965 := MakeNative(func(__e *ControlFlow) {
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

tmp21966 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symupdate_1lambda_1table, tmp21965)
}
__typedArg0 := symupdate_1lambda_1table
__typedArg1 := tmp21965
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21967 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21966)


_ = tmp21967

tmp21968 := MakeNative(func(__e *ControlFlow) {
Y1203 := __e.Get(1)
_ = Y1203
__e.TailApply(PrimFunc(symvector), Y1203)
return
}, 1)

tmp21969 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp21968)
}
__typedArg0 := symvector
__typedArg1 := tmp21968
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21970 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21969)


_ = tmp21970

tmp21971 := MakeNative(func(__e *ControlFlow) {
Y1204 := __e.Get(1)
_ = Y1204
__e.TailApply(PrimFunc(symvector_2), Y1204)
return
}, 1)

tmp21972 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector_2, tmp21971)
}
__typedArg0 := symvector_2
__typedArg1 := tmp21971
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21973 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21972)


_ = tmp21973

tmp21974 := MakeNative(func(__e *ControlFlow) {
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

tmp21975 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector_1_6, tmp21974)
}
__typedArg0 := symvector_1_6
__typedArg1 := tmp21974
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21976 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21975)


_ = tmp21976

tmp21977 := MakeNative(func(__e *ControlFlow) {
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

tmp21978 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvalue, tmp21977)
}
__typedArg0 := symvalue
__typedArg1 := tmp21977
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21979 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21978)


_ = tmp21979

tmp21980 := MakeNative(func(__e *ControlFlow) {
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

tmp21981 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvariable_2, tmp21980)
}
__typedArg0 := symvariable_2
__typedArg1 := tmp21980
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21982 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21981)


_ = tmp21982

tmp21983 := MakeNative(func(__e *ControlFlow) {
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

tmp21984 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvar_2, tmp21983)
}
__typedArg0 := symvar_2
__typedArg1 := tmp21983
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21985 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21984)


_ = tmp21985

tmp21986 := MakeNative(func(__e *ControlFlow) {
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

tmp21987 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhen, tmp21986)
}
__typedArg0 := symwhen
__typedArg1 := tmp21986
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21988 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21987)


_ = tmp21988

tmp21989 := MakeNative(func(__e *ControlFlow) {
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

tmp21990 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwrite_1byte, tmp21989)
}
__typedArg0 := symwrite_1byte
__typedArg1 := tmp21989
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21991 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21990)


_ = tmp21991

tmp21992 := MakeNative(func(__e *ControlFlow) {
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

tmp21993 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwrite_1string, tmp21992)
}
__typedArg0 := symwrite_1string
__typedArg1 := tmp21992
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21994 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21993)


_ = tmp21994

tmp21995 := MakeNative(func(__e *ControlFlow) {
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

tmp21996 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwrite_1to_1file, tmp21995)
}
__typedArg0 := symwrite_1to_1file
__typedArg1 := tmp21995
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21997 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21996)


_ = tmp21997

tmp21998 := MakeNative(func(__e *ControlFlow) {
Y1226 := __e.Get(1)
_ = Y1226
__e.TailApply(PrimFunc(symy_1or_1n_2), Y1226)
return
}, 1)

tmp21999 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symy_1or_1n_2, tmp21998)
}
__typedArg0 := symy_1or_1n_2
__typedArg1 := tmp21998
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp22000 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp21999)


_ = tmp22000

tmp22001 := MakeNative(func(__e *ControlFlow) {
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

tmp22002 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_7, tmp22001)
}
__typedArg0 := sym_7
__typedArg1 := tmp22001
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp22003 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp22002)


_ = tmp22003

tmp22004 := MakeNative(func(__e *ControlFlow) {
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

tmp22005 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_d, tmp22004)
}
__typedArg0 := sym_d
__typedArg1 := tmp22004
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp22006 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp22005)


_ = tmp22006

tmp22007 := MakeNative(func(__e *ControlFlow) {
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

tmp22008 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_c, tmp22007)
}
__typedArg0 := sym_c
__typedArg1 := tmp22007
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp22009 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp22008)


_ = tmp22009

tmp22010 := MakeNative(func(__e *ControlFlow) {
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

tmp22011 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1, tmp22010)
}
__typedArg0 := sym_1
__typedArg1 := tmp22010
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp22012 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp22011)


_ = tmp22012

tmp22013 := MakeNative(func(__e *ControlFlow) {
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

tmp22014 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_a_a, tmp22013)
}
__typedArg0 := sym_a_a
__typedArg1 := tmp22013
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp22015 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp22014)


_ = tmp22015

tmp22016 := MakeNative(func(__e *ControlFlow) {
Y1237 := __e.Get(1)
_ = Y1237
__e.TailApply(PrimFunc(sym_5e_6), Y1237)
return
}, 1)

tmp22017 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5e_6, tmp22016)
}
__typedArg0 := sym_5e_6
__typedArg1 := tmp22016
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp22018 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp22017)


_ = tmp22018

tmp22019 := MakeNative(func(__e *ControlFlow) {
Y1238 := __e.Get(1)
_ = Y1238
__e.TailApply(PrimFunc(sym_5end_6), Y1238)
return
}, 1)

tmp22020 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5end_6, tmp22019)
}
__typedArg0 := sym_5end_6
__typedArg1 := tmp22019
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp22021 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp22020)


_ = tmp22021

tmp22022 := MakeNative(func(__e *ControlFlow) {
Y1239 := __e.Get(1)
_ = Y1239
__e.TailApply(PrimFunc(sym_5_b_6), Y1239)
return
}, 1)

tmp22023 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_b_6, tmp22022)
}
__typedArg0 := sym_5_b_6
__typedArg1 := tmp22022
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp22024 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp22023)


_ = tmp22024

tmp22025 := MakeNative(func(__e *ControlFlow) {
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

tmp22026 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8p, tmp22025)
}
__typedArg0 := sym_8p
__typedArg1 := tmp22025
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp22027 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp22026)


_ = tmp22027

tmp22028 := MakeNative(func(__e *ControlFlow) {
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

tmp22029 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8v, tmp22028)
}
__typedArg0 := sym_8v
__typedArg1 := tmp22028
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp22030 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp22029)


_ = tmp22030

tmp22031 := MakeNative(func(__e *ControlFlow) {
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

tmp22032 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8s, tmp22031)
}
__typedArg0 := sym_8s
__typedArg1 := tmp22031
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4set_1lambda_1form_1entry), tmp22032)
return


}, 0)

tmp22033 := Call(__e, ns2_1set, symshen_4initialise_1lambda_1forms, tmp21478)


_ = tmp22033

tmp22034 := MakeNative(func(__e *ControlFlow) {
tmp22035 := Call(__e, PrimFunc(symshen_4initialise_1environment))


_ = tmp22035

tmp22036 := Call(__e, PrimFunc(symshen_4initialise_1lambda_1forms))


_ = tmp22036

__e.TailApply(PrimFunc(symshen_4initialise_1signedfuncs))
return


}, 0)

__e.TailApply(ns2_1set, symshen_4initialise, tmp22034)
return




}, 0)

