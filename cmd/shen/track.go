package main

import . "github.com/pyrex41/shen-go/kl"

var TrackMain = MakeNative(func(__e *ControlFlow) {
tmp14003 := MakeNative(func(__e *ControlFlow) {
V5384 := __e.Get(1)
_ = V5384
tmp14004 := Call(__e, PrimFunc(symshen_4app), V5384, MakeString(";\n"), symshen_4a)


tmp14005 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("partial function "))
__typedS1, __typedOK1 := TypedString(tmp14004)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("partial function ")
__typedArg1 := tmp14004
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp14006 := Call(__e, PrimFunc(symstoutput))


tmp14007 := Call(__e, PrimFunc(sympr), tmp14005, tmp14006)


_ = tmp14007

tmp14016 := Call(__e, PrimFunc(symshen_4tracked_2), V5384)


tmp14017 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp14016)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp14016
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres14011 Obj

if True == tmp14017 {
tmp14013 := Call(__e, PrimFunc(symshen_4app), V5384, MakeString("? "), symshen_4a)


tmp14015 := Call(__e, PrimFunc(symy_1or_1n_2), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("track "))
__typedS1, __typedOK1 := TypedString(tmp14013)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("track ")
__typedArg1 := tmp14013
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())


var ifres14012 Obj

if True == tmp14015 {
ifres14012 = True


} else {
ifres14012 = False


}

ifres14011 = ifres14012


} else {
ifres14011 = False


}

var ifres14008 Obj

if True == ifres14011 {
tmp14009 := Call(__e, PrimFunc(symps), V5384)


tmp14010 := Call(__e, PrimFunc(symshen_4track_1function), tmp14009)


ifres14008 = tmp14010


} else {
ifres14008 = symshen_4ok


}

_ = ifres14008

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("aborted"))
}
__typedArg0 := MakeString("aborted")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}, 1)

tmp14018 := Call(__e, ns2_1set, symshen_4f_1error, tmp14003)


_ = tmp14018

tmp14019 := MakeNative(func(__e *ControlFlow) {
V5385 := __e.Get(1)
_ = V5385
tmp14020 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dtracking_d)
}
__typedArg0 := symshen_4_dtracking_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symelement_2), V5385, tmp14020)
return


}, 1)

tmp14021 := Call(__e, ns2_1set, symshen_4tracked_2, tmp14019)


_ = tmp14021

tmp14022 := MakeNative(func(__e *ControlFlow) {
V5386 := __e.Get(1)
_ = V5386
tmp14023 := Call(__e, PrimFunc(symps), V5386)


let__13992 := tmp14023
_ = let__13992

__e.TailApply(PrimFunc(symshen_4track_1function), let__13992)
return


}, 1)

tmp14024 := Call(__e, ns2_1set, symtrack, tmp14022)


_ = tmp14024

tmp14025 := MakeNative(func(__e *ControlFlow) {
V5390 := __e.Get(1)
_ = V5390
tmp14079 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5390)
}
__typedArg0 := V5390
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14053 Obj

if True == tmp14079 {
tmp14077 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5390)
}
__typedArg0 := V5390
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14078 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefun, tmp14077)
}
__typedArg0 := symdefun
__typedArg1 := tmp14077
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14055 Obj

if True == tmp14078 {
tmp14075 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5390)
}
__typedArg0 := V5390
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14076 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14075)
}
__typedArg0 := tmp14075
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14057 Obj

if True == tmp14076 {
tmp14072 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5390)
}
__typedArg0 := V5390
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14073 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14072)
}
__typedArg0 := tmp14072
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14074 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14073)
}
__typedArg0 := tmp14073
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14059 Obj

if True == tmp14074 {
tmp14068 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5390)
}
__typedArg0 := V5390
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14069 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14068)
}
__typedArg0 := tmp14068
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14070 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14069)
}
__typedArg0 := tmp14069
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14071 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14070)
}
__typedArg0 := tmp14070
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14061 Obj

if True == tmp14071 {
tmp14063 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5390)
}
__typedArg0 := V5390
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14064 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14063)
}
__typedArg0 := tmp14063
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14065 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14064)
}
__typedArg0 := tmp14064
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14066 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14065)
}
__typedArg0 := tmp14065
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14067 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp14066)
}
__typedArg0 := Nil
__typedArg1 := tmp14066
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14062 Obj

if True == tmp14067 {
ifres14062 = True


} else {
ifres14062 = False


}

ifres14061 = ifres14062


} else {
ifres14061 = False


}

var ifres14060 Obj

if True == ifres14061 {
ifres14060 = True


} else {
ifres14060 = False


}

ifres14059 = ifres14060


} else {
ifres14059 = False


}

var ifres14058 Obj

if True == ifres14059 {
ifres14058 = True


} else {
ifres14058 = False


}

ifres14057 = ifres14058


} else {
ifres14057 = False


}

var ifres14056 Obj

if True == ifres14057 {
ifres14056 = True


} else {
ifres14056 = False


}

ifres14055 = ifres14056


} else {
ifres14055 = False


}

var ifres14054 Obj

if True == ifres14055 {
ifres14054 = True


} else {
ifres14054 = False


}

ifres14053 = ifres14054


} else {
ifres14053 = False


}

if True == ifres14053 {
tmp14026 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5390)
}
__typedArg0 := V5390
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14027 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp14026)
}
__typedArg0 := tmp14026
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14028 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5390)
}
__typedArg0 := V5390
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14029 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14028)
}
__typedArg0 := tmp14028
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14030 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp14029)
}
__typedArg0 := tmp14029
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14031 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5390)
}
__typedArg0 := V5390
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14032 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp14031)
}
__typedArg0 := tmp14031
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14033 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5390)
}
__typedArg0 := V5390
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14034 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14033)
}
__typedArg0 := tmp14033
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14035 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp14034)
}
__typedArg0 := tmp14034
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14036 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5390)
}
__typedArg0 := V5390
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14037 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14036)
}
__typedArg0 := tmp14036
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14038 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14037)
}
__typedArg0 := tmp14037
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14039 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp14038)
}
__typedArg0 := tmp14038
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14040 := Call(__e, PrimFunc(symshen_4insert_1tracking_1code), tmp14032, tmp14035, tmp14039)


tmp14041 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14040, Nil)
}
__typedArg0 := tmp14040
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14042 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14030, tmp14041)
}
__typedArg0 := tmp14030
__typedArg1 := tmp14041
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14043 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14027, tmp14042)
}
__typedArg0 := tmp14027
__typedArg1 := tmp14042
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14044 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefun, tmp14043)
}
__typedArg0 := symdefun
__typedArg1 := tmp14043
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__13993 := tmp14044
_ = let__13993

tmp14045 := Call(__e, PrimFunc(symeval_1kl), let__13993)


let__13994 := tmp14045
_ = let__13994

tmp14046 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5390)
}
__typedArg0 := V5390
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14047 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp14046)
}
__typedArg0 := tmp14046
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14048 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dtracking_d)
}
__typedArg0 := symshen_4_dtracking_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp14049 := Call(__e, PrimFunc(symadjoin), tmp14047, tmp14048)


tmp14050 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dtracking_d, tmp14049)
}
__typedArg0 := symshen_4_dtracking_d
__typedArg1 := tmp14049
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

let__13995 := tmp14050
_ = let__13995

tmp14051 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5390)
}
__typedArg0 := V5390
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp14051)
}
__typedArg0 := tmp14051
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.track-function"))
}
__typedArg0 := MakeString("implementation error in shen.track-function")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}, 1)

tmp14080 := Call(__e, ns2_1set, symshen_4track_1function, tmp14025)


_ = tmp14080

tmp14081 := MakeNative(func(__e *ControlFlow) {
V5394 := __e.Get(1)
_ = V5394
V5395 := __e.Get(2)
_ = V5395
V5396 := __e.Get(3)
_ = V5396
tmp14082 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4_dcall_d, Nil)
}
__typedArg0 := symshen_4_dcall_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14083 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvalue, tmp14082)
}
__typedArg0 := symvalue
__typedArg1 := tmp14082
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14084 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), Nil)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14085 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14083, tmp14084)
}
__typedArg0 := tmp14083
__typedArg1 := tmp14084
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14086 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_7, tmp14085)
}
__typedArg0 := sym_7
__typedArg1 := tmp14085
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14087 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14086, Nil)
}
__typedArg0 := tmp14086
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14088 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4_dcall_d, tmp14087)
}
__typedArg0 := symshen_4_dcall_d
__typedArg1 := tmp14087
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14089 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symset, tmp14088)
}
__typedArg0 := symset
__typedArg1 := tmp14088
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14090 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4_dcall_d, Nil)
}
__typedArg0 := symshen_4_dcall_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14091 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvalue, tmp14090)
}
__typedArg0 := symvalue
__typedArg1 := tmp14090
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14092 := Call(__e, PrimFunc(symshen_4prolog_1track), V5396, V5395)


tmp14093 := Call(__e, PrimFunc(symshen_4cons_1form), tmp14092)


tmp14094 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14093, Nil)
}
__typedArg0 := tmp14093
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14095 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5394, tmp14094)
}
__typedArg0 := V5394
__typedArg1 := tmp14094
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14096 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14091, tmp14095)
}
__typedArg0 := tmp14091
__typedArg1 := tmp14095
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14097 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4input_1track, tmp14096)
}
__typedArg0 := symshen_4input_1track
__typedArg1 := tmp14096
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14098 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4terpri_1or_1read_1char, Nil)
}
__typedArg0 := symshen_4terpri_1or_1read_1char
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14099 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4_dcall_d, Nil)
}
__typedArg0 := symshen_4_dcall_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14100 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvalue, tmp14099)
}
__typedArg0 := symvalue
__typedArg1 := tmp14099
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14101 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symResult, Nil)
}
__typedArg0 := symResult
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14102 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5394, tmp14101)
}
__typedArg0 := V5394
__typedArg1 := tmp14101
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14103 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14100, tmp14102)
}
__typedArg0 := tmp14100
__typedArg1 := tmp14102
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14104 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4output_1track, tmp14103)
}
__typedArg0 := symshen_4output_1track
__typedArg1 := tmp14103
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14105 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4_dcall_d, Nil)
}
__typedArg0 := symshen_4_dcall_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14106 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvalue, tmp14105)
}
__typedArg0 := symvalue
__typedArg1 := tmp14105
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14107 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), Nil)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14108 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14106, tmp14107)
}
__typedArg0 := tmp14106
__typedArg1 := tmp14107
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14109 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1, tmp14108)
}
__typedArg0 := sym_1
__typedArg1 := tmp14108
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14110 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14109, Nil)
}
__typedArg0 := tmp14109
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14111 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4_dcall_d, tmp14110)
}
__typedArg0 := symshen_4_dcall_d
__typedArg1 := tmp14110
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14112 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symset, tmp14111)
}
__typedArg0 := symset
__typedArg1 := tmp14111
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14113 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4terpri_1or_1read_1char, Nil)
}
__typedArg0 := symshen_4terpri_1or_1read_1char
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14114 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symResult, Nil)
}
__typedArg0 := symResult
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14115 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14113, tmp14114)
}
__typedArg0 := tmp14113
__typedArg1 := tmp14114
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14116 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdo, tmp14115)
}
__typedArg0 := symdo
__typedArg1 := tmp14115
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14117 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14116, Nil)
}
__typedArg0 := tmp14116
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14118 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14112, tmp14117)
}
__typedArg0 := tmp14112
__typedArg1 := tmp14117
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14119 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdo, tmp14118)
}
__typedArg0 := symdo
__typedArg1 := tmp14118
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14120 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14119, Nil)
}
__typedArg0 := tmp14119
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14121 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14104, tmp14120)
}
__typedArg0 := tmp14104
__typedArg1 := tmp14120
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14122 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdo, tmp14121)
}
__typedArg0 := symdo
__typedArg1 := tmp14121
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14123 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14122, Nil)
}
__typedArg0 := tmp14122
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14124 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5396, tmp14123)
}
__typedArg0 := V5396
__typedArg1 := tmp14123
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14125 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symResult, tmp14124)
}
__typedArg0 := symResult
__typedArg1 := tmp14124
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14126 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp14125)
}
__typedArg0 := symlet
__typedArg1 := tmp14125
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14127 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14126, Nil)
}
__typedArg0 := tmp14126
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14128 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14098, tmp14127)
}
__typedArg0 := tmp14098
__typedArg1 := tmp14127
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14129 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdo, tmp14128)
}
__typedArg0 := symdo
__typedArg1 := tmp14128
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14130 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14129, Nil)
}
__typedArg0 := tmp14129
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14131 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14097, tmp14130)
}
__typedArg0 := tmp14097
__typedArg1 := tmp14130
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14132 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdo, tmp14131)
}
__typedArg0 := symdo
__typedArg1 := tmp14131
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14133 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14132, Nil)
}
__typedArg0 := tmp14132
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14134 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14089, tmp14133)
}
__typedArg0 := tmp14089
__typedArg1 := tmp14133
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdo, tmp14134)
}
__typedArg0 := symdo
__typedArg1 := tmp14134
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 3)

tmp14135 := Call(__e, ns2_1set, symshen_4insert_1tracking_1code, tmp14081)


_ = tmp14135

tmp14136 := MakeNative(func(__e *ControlFlow) {
V5397 := __e.Get(1)
_ = V5397
V5398 := __e.Get(2)
_ = V5398
tmp14139 := Call(__e, PrimFunc(symoccurrences), symshen_4incinfs, V5397)


tmp14140 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp14139, MakeNumber(0))
}
__typedArg0 := tmp14139
__typedArg1 := MakeNumber(0)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp14140 {
__e.Return(V5398)
return
} else {
tmp14137 := Call(__e, PrimFunc(symshen_4vector_1parameter), V5398)


__e.TailApply(PrimFunc(symshen_4vector_1dereference), V5398, tmp14137)
return


}


}, 2)

tmp14141 := Call(__e, ns2_1set, symshen_4prolog_1track, tmp14136)


_ = tmp14141

tmp14142 := MakeNative(func(__e *ControlFlow) {
V5401 := __e.Get(1)
_ = V5401
tmp14171 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V5401)
}
__typedArg0 := Nil
__typedArg1 := V5401
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp14171 {
__e.Return(Nil)
return
} else {
tmp14169 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5401)
}
__typedArg0 := V5401
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14147 Obj

if True == tmp14169 {
tmp14167 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5401)
}
__typedArg0 := V5401
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14168 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14167)
}
__typedArg0 := tmp14167
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14149 Obj

if True == tmp14168 {
tmp14164 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5401)
}
__typedArg0 := V5401
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14165 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14164)
}
__typedArg0 := tmp14164
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14166 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14165)
}
__typedArg0 := tmp14165
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14151 Obj

if True == tmp14166 {
tmp14160 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5401)
}
__typedArg0 := V5401
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14161 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14160)
}
__typedArg0 := tmp14160
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14162 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14161)
}
__typedArg0 := tmp14161
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14163 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14162)
}
__typedArg0 := tmp14162
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14153 Obj

if True == tmp14163 {
tmp14155 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5401)
}
__typedArg0 := V5401
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14156 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14155)
}
__typedArg0 := tmp14155
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14157 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14156)
}
__typedArg0 := tmp14156
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14158 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14157)
}
__typedArg0 := tmp14157
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14159 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp14158)
}
__typedArg0 := Nil
__typedArg1 := tmp14158
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14154 Obj

if True == tmp14159 {
ifres14154 = True


} else {
ifres14154 = False


}

ifres14153 = ifres14154


} else {
ifres14153 = False


}

var ifres14152 Obj

if True == ifres14153 {
ifres14152 = True


} else {
ifres14152 = False


}

ifres14151 = ifres14152


} else {
ifres14151 = False


}

var ifres14150 Obj

if True == ifres14151 {
ifres14150 = True


} else {
ifres14150 = False


}

ifres14149 = ifres14150


} else {
ifres14149 = False


}

var ifres14148 Obj

if True == ifres14149 {
ifres14148 = True


} else {
ifres14148 = False


}

ifres14147 = ifres14148


} else {
ifres14147 = False


}

if True == ifres14147 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5401)
}
__typedArg0 := V5401
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return
} else {
tmp14145 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5401)
}
__typedArg0 := V5401
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp14145 {
tmp14143 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5401)
}
__typedArg0 := V5401
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4vector_1parameter), tmp14143)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("partial function shen.vector-parameter"))
}
__typedArg0 := MakeString("partial function shen.vector-parameter")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}


}, 1)

tmp14172 := Call(__e, ns2_1set, symshen_4vector_1parameter, tmp14142)


_ = tmp14172

tmp14173 := MakeNative(func(__e *ControlFlow) {
V5404 := __e.Get(1)
_ = V5404
V5405 := __e.Get(2)
_ = V5405
tmp14207 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V5405)
}
__typedArg0 := Nil
__typedArg1 := V5405
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp14207 {
__e.Return(V5404)
return
} else {
tmp14205 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5404)
}
__typedArg0 := V5404
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14183 Obj

if True == tmp14205 {
tmp14203 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5404)
}
__typedArg0 := V5404
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14204 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14203)
}
__typedArg0 := tmp14203
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14185 Obj

if True == tmp14204 {
tmp14200 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5404)
}
__typedArg0 := V5404
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14201 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14200)
}
__typedArg0 := tmp14200
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14202 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14201)
}
__typedArg0 := tmp14201
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14187 Obj

if True == tmp14202 {
tmp14196 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5404)
}
__typedArg0 := V5404
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14197 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14196)
}
__typedArg0 := tmp14196
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14198 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14197)
}
__typedArg0 := tmp14197
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14199 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14198)
}
__typedArg0 := tmp14198
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14189 Obj

if True == tmp14199 {
tmp14191 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5404)
}
__typedArg0 := V5404
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14192 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14191)
}
__typedArg0 := tmp14191
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14193 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14192)
}
__typedArg0 := tmp14192
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14194 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14193)
}
__typedArg0 := tmp14193
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14195 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp14194)
}
__typedArg0 := Nil
__typedArg1 := tmp14194
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14190 Obj

if True == tmp14195 {
ifres14190 = True


} else {
ifres14190 = False


}

ifres14189 = ifres14190


} else {
ifres14189 = False


}

var ifres14188 Obj

if True == ifres14189 {
ifres14188 = True


} else {
ifres14188 = False


}

ifres14187 = ifres14188


} else {
ifres14187 = False


}

var ifres14186 Obj

if True == ifres14187 {
ifres14186 = True


} else {
ifres14186 = False


}

ifres14185 = ifres14186


} else {
ifres14185 = False


}

var ifres14184 Obj

if True == ifres14185 {
ifres14184 = True


} else {
ifres14184 = False


}

ifres14183 = ifres14184


} else {
ifres14183 = False


}

if True == ifres14183 {
__e.Return(V5404)
return
} else {
tmp14181 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5404)
}
__typedArg0 := V5404
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp14181 {
tmp14174 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5404)
}
__typedArg0 := V5404
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14175 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5405, Nil)
}
__typedArg0 := V5405
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14176 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14174, tmp14175)
}
__typedArg0 := tmp14174
__typedArg1 := tmp14175
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14177 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4deref, tmp14176)
}
__typedArg0 := symshen_4deref
__typedArg1 := tmp14176
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14178 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5404)
}
__typedArg0 := V5404
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14179 := Call(__e, PrimFunc(symshen_4vector_1dereference), tmp14178, V5405)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14177, tmp14179)
}
__typedArg0 := tmp14177
__typedArg1 := tmp14179
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("partial function shen.vector-dereference"))
}
__typedArg0 := MakeString("partial function shen.vector-dereference")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}


}, 2)

tmp14208 := Call(__e, ns2_1set, symshen_4vector_1dereference, tmp14173)


_ = tmp14208

tmp14209 := MakeNative(func(__e *ControlFlow) {
V5408 := __e.Get(1)
_ = V5408
tmp14213 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_7, V5408)
}
__typedArg0 := sym_7
__typedArg1 := V5408
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp14213 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dstep_d, True)
}
__typedArg0 := symshen_4_dstep_d
__typedArg1 := True
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return
} else {
tmp14211 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_1, V5408)
}
__typedArg0 := sym_1
__typedArg1 := V5408
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp14211 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dstep_d, False)
}
__typedArg0 := symshen_4_dstep_d
__typedArg1 := False
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return
} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("step expects a + or a -.\n"))
}
__typedArg0 := MakeString("step expects a + or a -.\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 1)

tmp14214 := Call(__e, ns2_1set, symstep, tmp14209)


_ = tmp14214

tmp14215 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dstep_d)
}
__typedArg0 := symshen_4_dstep_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp14216 := Call(__e, ns2_1set, symstep_2, tmp14215)


_ = tmp14216

tmp14217 := MakeNative(func(__e *ControlFlow) {
V5411 := __e.Get(1)
_ = V5411
tmp14221 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_7, V5411)
}
__typedArg0 := sym_7
__typedArg1 := V5411
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp14221 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dspy_d, True)
}
__typedArg0 := symshen_4_dspy_d
__typedArg1 := True
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return
} else {
tmp14219 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_1, V5411)
}
__typedArg0 := sym_1
__typedArg1 := V5411
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp14219 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dspy_d, False)
}
__typedArg0 := symshen_4_dspy_d
__typedArg1 := False
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return
} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("spy expects a + or a -.\n"))
}
__typedArg0 := MakeString("spy expects a + or a -.\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 1)

tmp14222 := Call(__e, ns2_1set, symspy, tmp14217)


_ = tmp14222

tmp14223 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dspy_d)
}
__typedArg0 := symshen_4_dspy_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp14224 := Call(__e, ns2_1set, symspy_2, tmp14223)


_ = tmp14224

tmp14225 := MakeNative(func(__e *ControlFlow) {
tmp14229 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dstep_d)
}
__typedArg0 := symshen_4_dstep_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

if True == tmp14229 {
tmp14226 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dstinput_d)
}
__typedArg0 := sym_dstinput_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp14227 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symread_1byte) {
return PrimReadByte(tmp14226)
}
__typedArg0 := tmp14226
return Call(__e, PrimFunc(symread_1byte), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4check_1byte), tmp14227)
return


} else {
__e.TailApply(PrimFunc(symnl), MakeNumber(1))
return
}


}, 0)

tmp14230 := Call(__e, ns2_1set, symshen_4terpri_1or_1read_1char, tmp14225)


_ = tmp14230

tmp14231 := MakeNative(func(__e *ControlFlow) {
V5414 := __e.Get(1)
_ = V5414
tmp14233 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeNumber(94), V5414)
}
__typedArg0 := MakeNumber(94)
__typedArg1 := V5414
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp14233 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("aborted"))
}
__typedArg0 := MakeString("aborted")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
} else {
__e.Return(True)
return
}


}, 1)

tmp14234 := Call(__e, ns2_1set, symshen_4check_1byte, tmp14231)


_ = tmp14234

tmp14235 := MakeNative(func(__e *ControlFlow) {
V5415 := __e.Get(1)
_ = V5415
V5416 := __e.Get(2)
_ = V5416
V5417 := __e.Get(3)
_ = V5417
tmp14236 := Call(__e, PrimFunc(symshen_4spaces), V5415)


tmp14237 := Call(__e, PrimFunc(symshen_4spaces), V5415)


tmp14238 := Call(__e, PrimFunc(symshen_4app), tmp14237, MakeString(""), symshen_4a)


tmp14240 := Call(__e, PrimFunc(symshen_4app), V5416, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(" \n"))
__typedS1, __typedOK1 := TypedString(tmp14238)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(" \n")
__typedArg1 := tmp14238
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


tmp14242 := Call(__e, PrimFunc(symshen_4app), V5415, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("> Inputs to "))
__typedS1, __typedOK1 := TypedString(tmp14240)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("> Inputs to ")
__typedArg1 := tmp14240
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


tmp14244 := Call(__e, PrimFunc(symshen_4app), tmp14236, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("<"))
__typedS1, __typedOK1 := TypedString(tmp14242)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("<")
__typedArg1 := tmp14242
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


tmp14245 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("\n"))
__typedS1, __typedOK1 := TypedString(tmp14244)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("\n")
__typedArg1 := tmp14244
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp14246 := Call(__e, PrimFunc(symstoutput))


tmp14247 := Call(__e, PrimFunc(sympr), tmp14245, tmp14246)


_ = tmp14247

__e.TailApply(PrimFunc(symshen_4recursively_1print), V5417)
return


}, 3)

tmp14248 := Call(__e, ns2_1set, symshen_4input_1track, tmp14235)


_ = tmp14248

tmp14249 := MakeNative(func(__e *ControlFlow) {
V5420 := __e.Get(1)
_ = V5420
tmp14259 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V5420)
}
__typedArg0 := Nil
__typedArg1 := V5420
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp14259 {
tmp14250 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), MakeString(" ==>"), tmp14250)
return


} else {
tmp14257 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5420)
}
__typedArg0 := V5420
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp14257 {
tmp14251 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5420)
}
__typedArg0 := V5420
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14252 := Call(__e, PrimFunc(symprint), tmp14251)


_ = tmp14252

tmp14253 := Call(__e, PrimFunc(symstoutput))


tmp14254 := Call(__e, PrimFunc(sympr), MakeString(", "), tmp14253)


_ = tmp14254

tmp14255 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5420)
}
__typedArg0 := V5420
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4recursively_1print), tmp14255)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.recursively-print"))
}
__typedArg0 := MakeString("implementation error in shen.recursively-print")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 1)

tmp14260 := Call(__e, ns2_1set, symshen_4recursively_1print, tmp14249)


_ = tmp14260

tmp14261 := MakeNative(func(__e *ControlFlow) {
V5421 := __e.Get(1)
_ = V5421
tmp14265 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeNumber(0), V5421)
}
__typedArg0 := MakeNumber(0)
__typedArg1 := V5421
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp14265 {
__e.Return(MakeString(""))
return
} else {
tmp14263 := Call(__e, PrimFunc(symshen_4spaces), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(V5421)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := V5421
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})())


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(" "))
__typedS1, __typedOK1 := TypedString(tmp14263)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(" ")
__typedArg1 := tmp14263
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
return


}


}, 1)

tmp14266 := Call(__e, ns2_1set, symshen_4spaces, tmp14261)


_ = tmp14266

tmp14267 := MakeNative(func(__e *ControlFlow) {
V5422 := __e.Get(1)
_ = V5422
V5423 := __e.Get(2)
_ = V5423
V5424 := __e.Get(3)
_ = V5424
tmp14268 := Call(__e, PrimFunc(symshen_4spaces), V5422)


tmp14269 := Call(__e, PrimFunc(symshen_4spaces), V5422)


tmp14270 := Call(__e, PrimFunc(symshen_4app), V5424, MakeString(""), symshen_4s)


tmp14272 := Call(__e, PrimFunc(symshen_4app), tmp14269, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("==> "))
__typedS1, __typedOK1 := TypedString(tmp14270)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("==> ")
__typedArg1 := tmp14270
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


tmp14274 := Call(__e, PrimFunc(symshen_4app), V5423, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(" \n"))
__typedS1, __typedOK1 := TypedString(tmp14272)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(" \n")
__typedArg1 := tmp14272
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


tmp14276 := Call(__e, PrimFunc(symshen_4app), V5422, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("> Output from "))
__typedS1, __typedOK1 := TypedString(tmp14274)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("> Output from ")
__typedArg1 := tmp14274
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


tmp14278 := Call(__e, PrimFunc(symshen_4app), tmp14268, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("<"))
__typedS1, __typedOK1 := TypedString(tmp14276)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("<")
__typedArg1 := tmp14276
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


tmp14279 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("\n"))
__typedS1, __typedOK1 := TypedString(tmp14278)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("\n")
__typedArg1 := tmp14278
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp14280 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp14279, tmp14280)
return


}, 3)

tmp14281 := Call(__e, ns2_1set, symshen_4output_1track, tmp14267)


_ = tmp14281

tmp14282 := MakeNative(func(__e *ControlFlow) {
V5425 := __e.Get(1)
_ = V5425
tmp14283 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dtracking_d)
}
__typedArg0 := symshen_4_dtracking_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp14284 := Call(__e, PrimFunc(symremove), V5425, tmp14283)


tmp14285 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dtracking_d, tmp14284)
}
__typedArg0 := symshen_4_dtracking_d
__typedArg1 := tmp14284
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp14285

tmp14286 := MakeNative(func(__e *ControlFlow) {
tmp14287 := Call(__e, PrimFunc(symps), V5425)


__e.TailApply(PrimFunc(symeval), tmp14287)
return


}, 0)

tmp14288 := MakeNative(func(__e *ControlFlow) {
Z5426 := __e.Get(1)
_ = Z5426
__e.Return(V5425)
return
}, 1)

tmp14289 := Call(__e, try_1catch, tmp14286, tmp14288)


_ = tmp14289

__e.Return(V5425)
return


}, 1)

tmp14290 := Call(__e, ns2_1set, symuntrack, tmp14282)


_ = tmp14290

tmp14291 := MakeNative(func(__e *ControlFlow) {
V5427 := __e.Get(1)
_ = V5427
V5428 := __e.Get(2)
_ = V5428
__e.TailApply(PrimFunc(symshen_4remove_1h), V5427, V5428, Nil)
return
}, 2)

tmp14292 := Call(__e, ns2_1set, symremove, tmp14291)


_ = tmp14292

tmp14293 := MakeNative(func(__e *ControlFlow) {
V5438 := __e.Get(1)
_ = V5438
V5439 := __e.Get(2)
_ = V5439
V5440 := __e.Get(3)
_ = V5440
tmp14308 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V5439)
}
__typedArg0 := Nil
__typedArg1 := V5439
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp14308 {
__e.TailApply(PrimFunc(symreverse), V5440)
return
} else {
tmp14306 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5439)
}
__typedArg0 := V5439
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14302 Obj

if True == tmp14306 {
tmp14304 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5439)
}
__typedArg0 := V5439
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14305 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V5438, tmp14304)
}
__typedArg0 := V5438
__typedArg1 := tmp14304
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14303 Obj

if True == tmp14305 {
ifres14303 = True


} else {
ifres14303 = False


}

ifres14302 = ifres14303


} else {
ifres14302 = False


}

if True == ifres14302 {
tmp14294 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5439)
}
__typedArg0 := V5439
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14295 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5439)
}
__typedArg0 := V5439
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4remove_1h), tmp14294, tmp14295, V5440)
return


} else {
tmp14300 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5439)
}
__typedArg0 := V5439
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp14300 {
tmp14296 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5439)
}
__typedArg0 := V5439
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14297 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5439)
}
__typedArg0 := V5439
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14298 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14297, V5440)
}
__typedArg0 := tmp14297
__typedArg1 := V5440
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4remove_1h), V5438, tmp14296, tmp14298)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.remove-h"))
}
__typedArg0 := MakeString("implementation error in shen.remove-h")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}


}, 3)

tmp14309 := Call(__e, ns2_1set, symshen_4remove_1h, tmp14293)


_ = tmp14309

tmp14310 := MakeNative(func(__e *ControlFlow) {
V5441 := __e.Get(1)
_ = V5441
tmp14311 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dprofiled_d)
}
__typedArg0 := symshen_4_dprofiled_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp14312 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5441, tmp14311)
}
__typedArg0 := V5441
__typedArg1 := tmp14311
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14313 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dprofiled_d, tmp14312)
}
__typedArg0 := symshen_4_dprofiled_d
__typedArg1 := tmp14312
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp14313

tmp14314 := Call(__e, PrimFunc(symps), V5441)


__e.TailApply(PrimFunc(symshen_4profile_1help), tmp14314)
return


}, 1)

tmp14315 := Call(__e, ns2_1set, symprofile, tmp14310)


_ = tmp14315

tmp14316 := MakeNative(func(__e *ControlFlow) {
V5444 := __e.Get(1)
_ = V5444
tmp14381 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5444)
}
__typedArg0 := V5444
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14355 Obj

if True == tmp14381 {
tmp14379 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5444)
}
__typedArg0 := V5444
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14380 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefun, tmp14379)
}
__typedArg0 := symdefun
__typedArg1 := tmp14379
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14357 Obj

if True == tmp14380 {
tmp14377 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5444)
}
__typedArg0 := V5444
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14378 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14377)
}
__typedArg0 := tmp14377
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14359 Obj

if True == tmp14378 {
tmp14374 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5444)
}
__typedArg0 := V5444
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14375 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14374)
}
__typedArg0 := tmp14374
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14376 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14375)
}
__typedArg0 := tmp14375
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14361 Obj

if True == tmp14376 {
tmp14370 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5444)
}
__typedArg0 := V5444
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14371 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14370)
}
__typedArg0 := tmp14370
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14372 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14371)
}
__typedArg0 := tmp14371
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14373 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14372)
}
__typedArg0 := tmp14372
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14363 Obj

if True == tmp14373 {
tmp14365 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5444)
}
__typedArg0 := V5444
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14366 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14365)
}
__typedArg0 := tmp14365
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14367 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14366)
}
__typedArg0 := tmp14366
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14368 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14367)
}
__typedArg0 := tmp14367
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14369 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp14368)
}
__typedArg0 := Nil
__typedArg1 := tmp14368
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14364 Obj

if True == tmp14369 {
ifres14364 = True


} else {
ifres14364 = False


}

ifres14363 = ifres14364


} else {
ifres14363 = False


}

var ifres14362 Obj

if True == ifres14363 {
ifres14362 = True


} else {
ifres14362 = False


}

ifres14361 = ifres14362


} else {
ifres14361 = False


}

var ifres14360 Obj

if True == ifres14361 {
ifres14360 = True


} else {
ifres14360 = False


}

ifres14359 = ifres14360


} else {
ifres14359 = False


}

var ifres14358 Obj

if True == ifres14359 {
ifres14358 = True


} else {
ifres14358 = False


}

ifres14357 = ifres14358


} else {
ifres14357 = False


}

var ifres14356 Obj

if True == ifres14357 {
ifres14356 = True


} else {
ifres14356 = False


}

ifres14355 = ifres14356


} else {
ifres14355 = False


}

if True == ifres14355 {
tmp14317 := Call(__e, PrimFunc(symgensym), symshen_4f)


let__13996 := tmp14317
_ = let__13996

tmp14318 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5444)
}
__typedArg0 := V5444
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14319 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp14318)
}
__typedArg0 := tmp14318
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14320 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5444)
}
__typedArg0 := V5444
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14321 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14320)
}
__typedArg0 := tmp14320
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14322 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp14321)
}
__typedArg0 := tmp14321
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14323 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5444)
}
__typedArg0 := V5444
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14324 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp14323)
}
__typedArg0 := tmp14323
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14325 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5444)
}
__typedArg0 := V5444
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14326 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14325)
}
__typedArg0 := tmp14325
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14327 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp14326)
}
__typedArg0 := tmp14326
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14328 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5444)
}
__typedArg0 := V5444
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14329 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14328)
}
__typedArg0 := tmp14328
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14330 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp14329)
}
__typedArg0 := tmp14329
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14331 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__13996, tmp14330)
}
__typedArg0 := let__13996
__typedArg1 := tmp14330
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14332 := Call(__e, PrimFunc(symshen_4profile_1func), tmp14324, tmp14327, tmp14331)


tmp14333 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14332, Nil)
}
__typedArg0 := tmp14332
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14334 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14322, tmp14333)
}
__typedArg0 := tmp14322
__typedArg1 := tmp14333
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14335 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14319, tmp14334)
}
__typedArg0 := tmp14319
__typedArg1 := tmp14334
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14336 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefun, tmp14335)
}
__typedArg0 := symdefun
__typedArg1 := tmp14335
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__13997 := tmp14336
_ = let__13997

tmp14337 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5444)
}
__typedArg0 := V5444
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14338 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14337)
}
__typedArg0 := tmp14337
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14339 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp14338)
}
__typedArg0 := tmp14338
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14340 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5444)
}
__typedArg0 := V5444
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14341 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp14340)
}
__typedArg0 := tmp14340
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14342 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5444)
}
__typedArg0 := V5444
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14343 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14342)
}
__typedArg0 := tmp14342
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14344 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14343)
}
__typedArg0 := tmp14343
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14345 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp14344)
}
__typedArg0 := tmp14344
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14346 := Call(__e, PrimFunc(symsubst), let__13996, tmp14341, tmp14345)


tmp14347 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14346, Nil)
}
__typedArg0 := tmp14346
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14348 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14339, tmp14347)
}
__typedArg0 := tmp14339
__typedArg1 := tmp14347
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14349 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__13996, tmp14348)
}
__typedArg0 := let__13996
__typedArg1 := tmp14348
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14350 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefun, tmp14349)
}
__typedArg0 := symdefun
__typedArg1 := tmp14349
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__13998 := tmp14350
_ = let__13998

tmp14351 := Call(__e, PrimFunc(symeval_1kl), let__13997)


let__13999 := tmp14351
_ = let__13999

tmp14352 := Call(__e, PrimFunc(symeval_1kl), let__13998)


let__14000 := tmp14352
_ = let__14000

tmp14353 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5444)
}
__typedArg0 := V5444
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp14353)
}
__typedArg0 := tmp14353
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("Cannot profile.\n"))
}
__typedArg0 := MakeString("Cannot profile.\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}, 1)

tmp14382 := Call(__e, ns2_1set, symshen_4profile_1help, tmp14316)


_ = tmp14382

tmp14383 := MakeNative(func(__e *ControlFlow) {
V5450 := __e.Get(1)
_ = V5450
tmp14384 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dprofiled_d)
}
__typedArg0 := symshen_4_dprofiled_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp14385 := Call(__e, PrimFunc(symremove), V5450, tmp14384)


tmp14386 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dprofiled_d, tmp14385)
}
__typedArg0 := symshen_4_dprofiled_d
__typedArg1 := tmp14385
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp14386

tmp14387 := MakeNative(func(__e *ControlFlow) {
tmp14388 := Call(__e, PrimFunc(symps), V5450)


__e.TailApply(PrimFunc(symeval), tmp14388)
return


}, 0)

tmp14389 := MakeNative(func(__e *ControlFlow) {
Z5451 := __e.Get(1)
_ = Z5451
__e.Return(V5450)
return
}, 1)

__e.TailApply(try_1catch, tmp14387, tmp14389)
return


}, 1)

tmp14390 := Call(__e, ns2_1set, symunprofile, tmp14383)


_ = tmp14390

tmp14391 := MakeNative(func(__e *ControlFlow) {
V5452 := __e.Get(1)
_ = V5452
tmp14392 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dprofiled_d)
}
__typedArg0 := symshen_4_dprofiled_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symelement_2), V5452, tmp14392)
return


}, 1)

tmp14393 := Call(__e, ns2_1set, symshen_4profiled_2, tmp14391)


_ = tmp14393

tmp14394 := MakeNative(func(__e *ControlFlow) {
V5453 := __e.Get(1)
_ = V5453
V5454 := __e.Get(2)
_ = V5454
V5455 := __e.Get(3)
_ = V5455
tmp14395 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symrun, Nil)
}
__typedArg0 := symrun
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14396 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symget_1time, tmp14395)
}
__typedArg0 := symget_1time
__typedArg1 := tmp14395
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14397 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symrun, Nil)
}
__typedArg0 := symrun
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14398 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symget_1time, tmp14397)
}
__typedArg0 := symget_1time
__typedArg1 := tmp14397
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14399 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symStart, Nil)
}
__typedArg0 := symStart
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14400 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14398, tmp14399)
}
__typedArg0 := tmp14398
__typedArg1 := tmp14399
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14401 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1, tmp14400)
}
__typedArg0 := sym_1
__typedArg1 := tmp14400
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14402 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5453, Nil)
}
__typedArg0 := V5453
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14403 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4get_1profile, tmp14402)
}
__typedArg0 := symshen_4get_1profile
__typedArg1 := tmp14402
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14404 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symFinish, Nil)
}
__typedArg0 := symFinish
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14405 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14403, tmp14404)
}
__typedArg0 := tmp14403
__typedArg1 := tmp14404
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14406 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_7, tmp14405)
}
__typedArg0 := sym_7
__typedArg1 := tmp14405
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14407 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14406, Nil)
}
__typedArg0 := tmp14406
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14408 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5453, tmp14407)
}
__typedArg0 := V5453
__typedArg1 := tmp14407
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14409 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4put_1profile, tmp14408)
}
__typedArg0 := symshen_4put_1profile
__typedArg1 := tmp14408
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14410 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symResult, Nil)
}
__typedArg0 := symResult
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14411 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14409, tmp14410)
}
__typedArg0 := tmp14409
__typedArg1 := tmp14410
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14412 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symRecord, tmp14411)
}
__typedArg0 := symRecord
__typedArg1 := tmp14411
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14413 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp14412)
}
__typedArg0 := symlet
__typedArg1 := tmp14412
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14414 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14413, Nil)
}
__typedArg0 := tmp14413
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14415 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14401, tmp14414)
}
__typedArg0 := tmp14401
__typedArg1 := tmp14414
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14416 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symFinish, tmp14415)
}
__typedArg0 := symFinish
__typedArg1 := tmp14415
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14417 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp14416)
}
__typedArg0 := symlet
__typedArg1 := tmp14416
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14418 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14417, Nil)
}
__typedArg0 := tmp14417
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14419 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5455, tmp14418)
}
__typedArg0 := V5455
__typedArg1 := tmp14418
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14420 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symResult, tmp14419)
}
__typedArg0 := symResult
__typedArg1 := tmp14419
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14421 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp14420)
}
__typedArg0 := symlet
__typedArg1 := tmp14420
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14422 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14421, Nil)
}
__typedArg0 := tmp14421
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14423 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14396, tmp14422)
}
__typedArg0 := tmp14396
__typedArg1 := tmp14422
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14424 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symStart, tmp14423)
}
__typedArg0 := symStart
__typedArg1 := tmp14423
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp14424)
}
__typedArg0 := symlet
__typedArg1 := tmp14424
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 3)

tmp14425 := Call(__e, ns2_1set, symshen_4profile_1func, tmp14394)


_ = tmp14425

tmp14426 := MakeNative(func(__e *ControlFlow) {
V5456 := __e.Get(1)
_ = V5456
tmp14427 := Call(__e, PrimFunc(symshen_4get_1profile), V5456)


let__14001 := tmp14427
_ = let__14001

tmp14428 := Call(__e, PrimFunc(symshen_4put_1profile), V5456, MakeNumber(0))


let__14002 := tmp14428
_ = let__14002

__e.TailApply(PrimFunc(sym_8p), V5456, let__14001)
return


}, 1)

tmp14429 := Call(__e, ns2_1set, symprofile_1results, tmp14426)


_ = tmp14429

tmp14430 := MakeNative(func(__e *ControlFlow) {
V5459 := __e.Get(1)
_ = V5459
tmp14431 := MakeNative(func(__e *ControlFlow) {
tmp14432 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symget), V5459, symprofile, tmp14432)
return


}, 0)

tmp14433 := MakeNative(func(__e *ControlFlow) {
Z5460 := __e.Get(1)
_ = Z5460
__e.Return(MakeNumber(0))
return
}, 1)

__e.TailApply(try_1catch, tmp14431, tmp14433)
return


}, 1)

tmp14434 := Call(__e, ns2_1set, symshen_4get_1profile, tmp14430)


_ = tmp14434

tmp14435 := MakeNative(func(__e *ControlFlow) {
V5461 := __e.Get(1)
_ = V5461
V5462 := __e.Get(2)
_ = V5462
tmp14436 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symput), V5461, symprofile, V5462, tmp14436)
return


}, 2)

__e.TailApply(ns2_1set, symshen_4put_1profile, tmp14435)
return




}, 0)

