package main

import . "github.com/pyrex41/shen-go/kl"

var SequentMain = MakeNative(func(__e *ControlFlow) {
tmp12354 := MakeNative(func(__e *ControlFlow) {
V3033 := __e.Get(1)
_ = V3033
tmp12355 := MakeNative(func(__e *ControlFlow) {
W3034 := __e.Get(1)
_ = W3034
tmp12357 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3034)


if True == tmp12357 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3034)
return
}


}, 1)

tmp12377 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3033)
}
__typedArg0 := V3033
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12358 Obj

if True == tmp12377 {
tmp12359 := MakeNative(func(__e *ControlFlow) {
W3035 := __e.Get(1)
_ = W3035
tmp12360 := MakeNative(func(__e *ControlFlow) {
W3036 := __e.Get(1)
_ = W3036
tmp12361 := MakeNative(func(__e *ControlFlow) {
W3037 := __e.Get(1)
_ = W3037
tmp12371 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3037)


if True == tmp12371 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12362 := MakeNative(func(__e *ControlFlow) {
W3038 := __e.Get(1)
_ = W3038
tmp12363 := MakeNative(func(__e *ControlFlow) {
W3039 := __e.Get(1)
_ = W3039
tmp12364 := MakeNative(func(__e *ControlFlow) {
W3040 := __e.Get(1)
_ = W3040
tmp12365 := Call(__e, PrimFunc(symfn), W3035)


__e.TailApply(PrimFunc(symshen_4remember_1datatype), W3035, tmp12365)
return


}, 1)

tmp12366 := Call(__e, PrimFunc(symshen_4rules_1_6prolog), W3035, W3038)


tmp12367 := Call(__e, tmp12364, tmp12366)


__e.TailApply(PrimFunc(symshen_4comb), W3039, tmp12367)
return


}, 1)

tmp12368 := Call(__e, PrimFunc(symshen_4in_1_6), W3037)


__e.TailApply(tmp12363, tmp12368)
return


}, 1)

tmp12369 := Call(__e, PrimFunc(symshen_4_5_1out), W3037)


__e.TailApply(tmp12362, tmp12369)
return


}


}, 1)

tmp12372 := Call(__e, PrimFunc(symshen_4_5datatype_1rules_6), W3036)


__e.TailApply(tmp12361, tmp12372)
return


}, 1)

tmp12373 := Call(__e, PrimFunc(symtail), V3033)


__e.TailApply(tmp12360, tmp12373)
return


}, 1)

tmp12374 := Call(__e, PrimFunc(symhead), V3033)


tmp12375 := Call(__e, tmp12359, tmp12374)


ifres12358 = tmp12375


} else {
tmp12376 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12358 = tmp12376


}

__e.TailApply(tmp12355, ifres12358)
return


}, 1)

tmp12378 := Call(__e, ns2_1set, symshen_4_5datatype_6, tmp12354)


_ = tmp12378

tmp12379 := MakeNative(func(__e *ControlFlow) {
V3041 := __e.Get(1)
_ = V3041
V3042 := __e.Get(2)
_ = V3042
tmp12380 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_ddatatypes_d)
}
__typedArg0 := symshen_4_ddatatypes_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp12381 := Call(__e, PrimFunc(symshen_4assoc_1_6), V3041, V3042, tmp12380)


tmp12382 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_ddatatypes_d, tmp12381)
}
__typedArg0 := symshen_4_ddatatypes_d
__typedArg1 := tmp12381
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp12382

tmp12383 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dalldatatypes_d)
}
__typedArg0 := symshen_4_dalldatatypes_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp12384 := Call(__e, PrimFunc(symshen_4assoc_1_6), V3041, V3042, tmp12383)


tmp12385 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dalldatatypes_d, tmp12384)
}
__typedArg0 := symshen_4_dalldatatypes_d
__typedArg1 := tmp12384
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp12385

__e.Return(V3041)
return


}, 2)

tmp12386 := Call(__e, ns2_1set, symshen_4remember_1datatype, tmp12379)


_ = tmp12386

tmp12387 := MakeNative(func(__e *ControlFlow) {
V3043 := __e.Get(1)
_ = V3043
tmp12388 := MakeNative(func(__e *ControlFlow) {
W3044 := __e.Get(1)
_ = W3044
tmp12407 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3044)


if True == tmp12407 {
tmp12389 := MakeNative(func(__e *ControlFlow) {
W3051 := __e.Get(1)
_ = W3051
tmp12391 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3051)


if True == tmp12391 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3051)
return
}


}, 1)

tmp12392 := MakeNative(func(__e *ControlFlow) {
W3052 := __e.Get(1)
_ = W3052
tmp12403 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3052)


if True == tmp12403 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12393 := MakeNative(func(__e *ControlFlow) {
W3053 := __e.Get(1)
_ = W3053
tmp12394 := MakeNative(func(__e *ControlFlow) {
W3054 := __e.Get(1)
_ = W3054
tmp12399 := Call(__e, PrimFunc(symempty_2), W3053)


var ifres12395 Obj

if True == tmp12399 {
ifres12395 = Nil


} else {
tmp12396 := Call(__e, PrimFunc(symshen_4app), W3053, MakeString("\n ..."), symshen_4r)


tmp12398 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("datatype syntax error here:\n "))
__typedS1, __typedOK1 := TypedString(tmp12396)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("datatype syntax error here:\n ")
__typedArg1 := tmp12396
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("datatype syntax error here:\n "))
__typedS1, __typedOK1 := TypedString(tmp12396)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("datatype syntax error here:\n ")
__typedArg1 := tmp12396
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})()

ifres12395 = tmp12398


}

__e.TailApply(PrimFunc(symshen_4comb), W3054, ifres12395)
return


}, 1)

tmp12400 := Call(__e, PrimFunc(symshen_4in_1_6), W3052)


__e.TailApply(tmp12394, tmp12400)
return


}, 1)

tmp12401 := Call(__e, PrimFunc(symshen_4_5_1out), W3052)


__e.TailApply(tmp12393, tmp12401)
return


}


}, 1)

tmp12404 := Call(__e, PrimFunc(sym_5_b_6), V3043)


tmp12405 := Call(__e, tmp12392, tmp12404)


__e.TailApply(tmp12389, tmp12405)
return


} else {
__e.Return(W3044)
return
}


}, 1)

tmp12408 := MakeNative(func(__e *ControlFlow) {
W3045 := __e.Get(1)
_ = W3045
tmp12423 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3045)


if True == tmp12423 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12409 := MakeNative(func(__e *ControlFlow) {
W3046 := __e.Get(1)
_ = W3046
tmp12410 := MakeNative(func(__e *ControlFlow) {
W3047 := __e.Get(1)
_ = W3047
tmp12411 := MakeNative(func(__e *ControlFlow) {
W3048 := __e.Get(1)
_ = W3048
tmp12418 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3048)


if True == tmp12418 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12412 := MakeNative(func(__e *ControlFlow) {
W3049 := __e.Get(1)
_ = W3049
tmp12413 := MakeNative(func(__e *ControlFlow) {
W3050 := __e.Get(1)
_ = W3050
tmp12414 := Call(__e, PrimFunc(symappend), W3046, W3049)


__e.TailApply(PrimFunc(symshen_4comb), W3050, tmp12414)
return


}, 1)

tmp12415 := Call(__e, PrimFunc(symshen_4in_1_6), W3048)


__e.TailApply(tmp12413, tmp12415)
return


}, 1)

tmp12416 := Call(__e, PrimFunc(symshen_4_5_1out), W3048)


__e.TailApply(tmp12412, tmp12416)
return


}


}, 1)

tmp12419 := Call(__e, PrimFunc(symshen_4_5datatype_1rules_6), W3047)


__e.TailApply(tmp12411, tmp12419)
return


}, 1)

tmp12420 := Call(__e, PrimFunc(symshen_4in_1_6), W3045)


__e.TailApply(tmp12410, tmp12420)
return


}, 1)

tmp12421 := Call(__e, PrimFunc(symshen_4_5_1out), W3045)


__e.TailApply(tmp12409, tmp12421)
return


}


}, 1)

tmp12424 := Call(__e, PrimFunc(symshen_4_5datatype_1rule_6), V3043)


tmp12425 := Call(__e, tmp12408, tmp12424)


__e.TailApply(tmp12388, tmp12425)
return


}, 1)

tmp12426 := Call(__e, ns2_1set, symshen_4_5datatype_1rules_6, tmp12387)


_ = tmp12426

tmp12427 := MakeNative(func(__e *ControlFlow) {
V3055 := __e.Get(1)
_ = V3055
tmp12428 := MakeNative(func(__e *ControlFlow) {
W3056 := __e.Get(1)
_ = W3056
tmp12442 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3056)


if True == tmp12442 {
tmp12429 := MakeNative(func(__e *ControlFlow) {
W3060 := __e.Get(1)
_ = W3060
tmp12431 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3060)


if True == tmp12431 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3060)
return
}


}, 1)

tmp12432 := MakeNative(func(__e *ControlFlow) {
W3061 := __e.Get(1)
_ = W3061
tmp12438 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3061)


if True == tmp12438 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12433 := MakeNative(func(__e *ControlFlow) {
W3062 := __e.Get(1)
_ = W3062
tmp12434 := MakeNative(func(__e *ControlFlow) {
W3063 := __e.Get(1)
_ = W3063
__e.TailApply(PrimFunc(symshen_4comb), W3063, W3062)
return
}, 1)

tmp12435 := Call(__e, PrimFunc(symshen_4in_1_6), W3061)


__e.TailApply(tmp12434, tmp12435)
return


}, 1)

tmp12436 := Call(__e, PrimFunc(symshen_4_5_1out), W3061)


__e.TailApply(tmp12433, tmp12436)
return


}


}, 1)

tmp12439 := Call(__e, PrimFunc(symshen_4_5double_6), V3055)


tmp12440 := Call(__e, tmp12432, tmp12439)


__e.TailApply(tmp12429, tmp12440)
return


} else {
__e.Return(W3056)
return
}


}, 1)

tmp12443 := MakeNative(func(__e *ControlFlow) {
W3057 := __e.Get(1)
_ = W3057
tmp12449 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3057)


if True == tmp12449 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12444 := MakeNative(func(__e *ControlFlow) {
W3058 := __e.Get(1)
_ = W3058
tmp12445 := MakeNative(func(__e *ControlFlow) {
W3059 := __e.Get(1)
_ = W3059
__e.TailApply(PrimFunc(symshen_4comb), W3059, W3058)
return
}, 1)

tmp12446 := Call(__e, PrimFunc(symshen_4in_1_6), W3057)


__e.TailApply(tmp12445, tmp12446)
return


}, 1)

tmp12447 := Call(__e, PrimFunc(symshen_4_5_1out), W3057)


__e.TailApply(tmp12444, tmp12447)
return


}


}, 1)

tmp12450 := Call(__e, PrimFunc(symshen_4_5single_6), V3055)


tmp12451 := Call(__e, tmp12443, tmp12450)


__e.TailApply(tmp12428, tmp12451)
return


}, 1)

tmp12452 := Call(__e, ns2_1set, symshen_4_5datatype_1rule_6, tmp12427)


_ = tmp12452

tmp12453 := MakeNative(func(__e *ControlFlow) {
V3064 := __e.Get(1)
_ = V3064
tmp12454 := MakeNative(func(__e *ControlFlow) {
W3065 := __e.Get(1)
_ = W3065
tmp12456 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3065)


if True == tmp12456 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3065)
return
}


}, 1)

tmp12457 := MakeNative(func(__e *ControlFlow) {
W3066 := __e.Get(1)
_ = W3066
tmp12495 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3066)


if True == tmp12495 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12458 := MakeNative(func(__e *ControlFlow) {
W3067 := __e.Get(1)
_ = W3067
tmp12459 := MakeNative(func(__e *ControlFlow) {
W3068 := __e.Get(1)
_ = W3068
tmp12460 := MakeNative(func(__e *ControlFlow) {
W3069 := __e.Get(1)
_ = W3069
tmp12490 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3069)


if True == tmp12490 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12461 := MakeNative(func(__e *ControlFlow) {
W3070 := __e.Get(1)
_ = W3070
tmp12462 := MakeNative(func(__e *ControlFlow) {
W3071 := __e.Get(1)
_ = W3071
tmp12463 := MakeNative(func(__e *ControlFlow) {
W3072 := __e.Get(1)
_ = W3072
tmp12485 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3072)


if True == tmp12485 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12464 := MakeNative(func(__e *ControlFlow) {
W3073 := __e.Get(1)
_ = W3073
tmp12465 := MakeNative(func(__e *ControlFlow) {
W3074 := __e.Get(1)
_ = W3074
tmp12481 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3074)


if True == tmp12481 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12466 := MakeNative(func(__e *ControlFlow) {
W3075 := __e.Get(1)
_ = W3075
tmp12467 := MakeNative(func(__e *ControlFlow) {
W3076 := __e.Get(1)
_ = W3076
tmp12468 := MakeNative(func(__e *ControlFlow) {
W3077 := __e.Get(1)
_ = W3077
tmp12476 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3077)


if True == tmp12476 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12469 := MakeNative(func(__e *ControlFlow) {
W3078 := __e.Get(1)
_ = W3078
tmp12470 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W3075, Nil)
}
__typedArg0 := W3075
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12471 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W3070, tmp12470)
}
__typedArg0 := W3070
__typedArg1 := tmp12470
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12472 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W3067, tmp12471)
}
__typedArg0 := W3067
__typedArg1 := tmp12471
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12473 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12472, Nil)
}
__typedArg0 := tmp12472
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4comb), W3078, tmp12473)
return


}, 1)

tmp12474 := Call(__e, PrimFunc(symshen_4in_1_6), W3077)


__e.TailApply(tmp12469, tmp12474)
return


}


}, 1)

tmp12477 := Call(__e, PrimFunc(symshen_4_5sc_6), W3076)


__e.TailApply(tmp12468, tmp12477)
return


}, 1)

tmp12478 := Call(__e, PrimFunc(symshen_4in_1_6), W3074)


__e.TailApply(tmp12467, tmp12478)
return


}, 1)

tmp12479 := Call(__e, PrimFunc(symshen_4_5_1out), W3074)


__e.TailApply(tmp12466, tmp12479)
return


}


}, 1)

tmp12482 := Call(__e, PrimFunc(symshen_4_5conc_6), W3073)


__e.TailApply(tmp12465, tmp12482)
return


}, 1)

tmp12483 := Call(__e, PrimFunc(symshen_4in_1_6), W3072)


__e.TailApply(tmp12464, tmp12483)
return


}


}, 1)

tmp12486 := Call(__e, PrimFunc(symshen_4_5sng_6), W3071)


__e.TailApply(tmp12463, tmp12486)
return


}, 1)

tmp12487 := Call(__e, PrimFunc(symshen_4in_1_6), W3069)


__e.TailApply(tmp12462, tmp12487)
return


}, 1)

tmp12488 := Call(__e, PrimFunc(symshen_4_5_1out), W3069)


__e.TailApply(tmp12461, tmp12488)
return


}


}, 1)

tmp12491 := Call(__e, PrimFunc(symshen_4_5prems_6), W3068)


__e.TailApply(tmp12460, tmp12491)
return


}, 1)

tmp12492 := Call(__e, PrimFunc(symshen_4in_1_6), W3066)


__e.TailApply(tmp12459, tmp12492)
return


}, 1)

tmp12493 := Call(__e, PrimFunc(symshen_4_5_1out), W3066)


__e.TailApply(tmp12458, tmp12493)
return


}


}, 1)

tmp12496 := Call(__e, PrimFunc(symshen_4_5sides_6), V3064)


tmp12497 := Call(__e, tmp12457, tmp12496)


__e.TailApply(tmp12454, tmp12497)
return


}, 1)

tmp12498 := Call(__e, ns2_1set, symshen_4_5single_6, tmp12453)


_ = tmp12498

tmp12499 := MakeNative(func(__e *ControlFlow) {
V3079 := __e.Get(1)
_ = V3079
tmp12500 := MakeNative(func(__e *ControlFlow) {
W3080 := __e.Get(1)
_ = W3080
tmp12502 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3080)


if True == tmp12502 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3080)
return
}


}, 1)

tmp12503 := MakeNative(func(__e *ControlFlow) {
W3081 := __e.Get(1)
_ = W3081
tmp12540 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3081)


if True == tmp12540 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12504 := MakeNative(func(__e *ControlFlow) {
W3082 := __e.Get(1)
_ = W3082
tmp12505 := MakeNative(func(__e *ControlFlow) {
W3083 := __e.Get(1)
_ = W3083
tmp12506 := MakeNative(func(__e *ControlFlow) {
W3084 := __e.Get(1)
_ = W3084
tmp12535 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3084)


if True == tmp12535 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12507 := MakeNative(func(__e *ControlFlow) {
W3085 := __e.Get(1)
_ = W3085
tmp12508 := MakeNative(func(__e *ControlFlow) {
W3086 := __e.Get(1)
_ = W3086
tmp12509 := MakeNative(func(__e *ControlFlow) {
W3087 := __e.Get(1)
_ = W3087
tmp12530 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3087)


if True == tmp12530 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12510 := MakeNative(func(__e *ControlFlow) {
W3088 := __e.Get(1)
_ = W3088
tmp12511 := MakeNative(func(__e *ControlFlow) {
W3089 := __e.Get(1)
_ = W3089
tmp12526 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3089)


if True == tmp12526 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12512 := MakeNative(func(__e *ControlFlow) {
W3090 := __e.Get(1)
_ = W3090
tmp12513 := MakeNative(func(__e *ControlFlow) {
W3091 := __e.Get(1)
_ = W3091
tmp12514 := MakeNative(func(__e *ControlFlow) {
W3092 := __e.Get(1)
_ = W3092
tmp12521 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3092)


if True == tmp12521 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12515 := MakeNative(func(__e *ControlFlow) {
W3093 := __e.Get(1)
_ = W3093
tmp12516 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W3090, Nil)
}
__typedArg0 := W3090
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12517 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Nil, tmp12516)
}
__typedArg0 := Nil
__typedArg1 := tmp12516
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12518 := Call(__e, PrimFunc(symshen_4lr_1rule), W3082, W3085, tmp12517)


__e.TailApply(PrimFunc(symshen_4comb), W3093, tmp12518)
return


}, 1)

tmp12519 := Call(__e, PrimFunc(symshen_4in_1_6), W3092)


__e.TailApply(tmp12515, tmp12519)
return


}


}, 1)

tmp12522 := Call(__e, PrimFunc(symshen_4_5sc_6), W3091)


__e.TailApply(tmp12514, tmp12522)
return


}, 1)

tmp12523 := Call(__e, PrimFunc(symshen_4in_1_6), W3089)


__e.TailApply(tmp12513, tmp12523)
return


}, 1)

tmp12524 := Call(__e, PrimFunc(symshen_4_5_1out), W3089)


__e.TailApply(tmp12512, tmp12524)
return


}


}, 1)

tmp12527 := Call(__e, PrimFunc(symshen_4_5formula_6), W3088)


__e.TailApply(tmp12511, tmp12527)
return


}, 1)

tmp12528 := Call(__e, PrimFunc(symshen_4in_1_6), W3087)


__e.TailApply(tmp12510, tmp12528)
return


}


}, 1)

tmp12531 := Call(__e, PrimFunc(symshen_4_5dbl_6), W3086)


__e.TailApply(tmp12509, tmp12531)
return


}, 1)

tmp12532 := Call(__e, PrimFunc(symshen_4in_1_6), W3084)


__e.TailApply(tmp12508, tmp12532)
return


}, 1)

tmp12533 := Call(__e, PrimFunc(symshen_4_5_1out), W3084)


__e.TailApply(tmp12507, tmp12533)
return


}


}, 1)

tmp12536 := Call(__e, PrimFunc(symshen_4_5formulae_6), W3083)


__e.TailApply(tmp12506, tmp12536)
return


}, 1)

tmp12537 := Call(__e, PrimFunc(symshen_4in_1_6), W3081)


__e.TailApply(tmp12505, tmp12537)
return


}, 1)

tmp12538 := Call(__e, PrimFunc(symshen_4_5_1out), W3081)


__e.TailApply(tmp12504, tmp12538)
return


}


}, 1)

tmp12541 := Call(__e, PrimFunc(symshen_4_5sides_6), V3079)


tmp12542 := Call(__e, tmp12503, tmp12541)


__e.TailApply(tmp12500, tmp12542)
return


}, 1)

tmp12543 := Call(__e, ns2_1set, symshen_4_5double_6, tmp12499)


_ = tmp12543

tmp12544 := MakeNative(func(__e *ControlFlow) {
V3094 := __e.Get(1)
_ = V3094
tmp12545 := MakeNative(func(__e *ControlFlow) {
W3095 := __e.Get(1)
_ = W3095
tmp12568 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3095)


if True == tmp12568 {
tmp12546 := MakeNative(func(__e *ControlFlow) {
W3104 := __e.Get(1)
_ = W3104
tmp12548 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3104)


if True == tmp12548 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3104)
return
}


}, 1)

tmp12549 := MakeNative(func(__e *ControlFlow) {
W3105 := __e.Get(1)
_ = W3105
tmp12564 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3105)


if True == tmp12564 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12550 := MakeNative(func(__e *ControlFlow) {
W3106 := __e.Get(1)
_ = W3106
tmp12551 := MakeNative(func(__e *ControlFlow) {
W3107 := __e.Get(1)
_ = W3107
tmp12552 := MakeNative(func(__e *ControlFlow) {
W3108 := __e.Get(1)
_ = W3108
tmp12559 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3108)


if True == tmp12559 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12553 := MakeNative(func(__e *ControlFlow) {
W3109 := __e.Get(1)
_ = W3109
tmp12554 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W3106, Nil)
}
__typedArg0 := W3106
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12555 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Nil, tmp12554)
}
__typedArg0 := Nil
__typedArg1 := tmp12554
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12556 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12555, Nil)
}
__typedArg0 := tmp12555
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4comb), W3109, tmp12556)
return


}, 1)

tmp12557 := Call(__e, PrimFunc(symshen_4in_1_6), W3108)


__e.TailApply(tmp12553, tmp12557)
return


}


}, 1)

tmp12560 := Call(__e, PrimFunc(symshen_4_5sc_6), W3107)


__e.TailApply(tmp12552, tmp12560)
return


}, 1)

tmp12561 := Call(__e, PrimFunc(symshen_4in_1_6), W3105)


__e.TailApply(tmp12551, tmp12561)
return


}, 1)

tmp12562 := Call(__e, PrimFunc(symshen_4_5_1out), W3105)


__e.TailApply(tmp12550, tmp12562)
return


}


}, 1)

tmp12565 := Call(__e, PrimFunc(symshen_4_5formula_6), V3094)


tmp12566 := Call(__e, tmp12549, tmp12565)


__e.TailApply(tmp12546, tmp12566)
return


} else {
__e.Return(W3095)
return
}


}, 1)

tmp12569 := MakeNative(func(__e *ControlFlow) {
W3096 := __e.Get(1)
_ = W3096
tmp12592 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3096)


if True == tmp12592 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12570 := MakeNative(func(__e *ControlFlow) {
W3097 := __e.Get(1)
_ = W3097
tmp12571 := MakeNative(func(__e *ControlFlow) {
W3098 := __e.Get(1)
_ = W3098
tmp12572 := MakeNative(func(__e *ControlFlow) {
W3099 := __e.Get(1)
_ = W3099
tmp12587 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3099)


if True == tmp12587 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12573 := MakeNative(func(__e *ControlFlow) {
W3100 := __e.Get(1)
_ = W3100
tmp12574 := MakeNative(func(__e *ControlFlow) {
W3101 := __e.Get(1)
_ = W3101
tmp12583 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3101)


if True == tmp12583 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12575 := MakeNative(func(__e *ControlFlow) {
W3102 := __e.Get(1)
_ = W3102
tmp12576 := MakeNative(func(__e *ControlFlow) {
W3103 := __e.Get(1)
_ = W3103
tmp12577 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W3097, Nil)
}
__typedArg0 := W3097
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12578 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Nil, tmp12577)
}
__typedArg0 := Nil
__typedArg1 := tmp12577
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12579 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12578, W3102)
}
__typedArg0 := tmp12578
__typedArg1 := W3102
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4comb), W3103, tmp12579)
return


}, 1)

tmp12580 := Call(__e, PrimFunc(symshen_4in_1_6), W3101)


__e.TailApply(tmp12576, tmp12580)
return


}, 1)

tmp12581 := Call(__e, PrimFunc(symshen_4_5_1out), W3101)


__e.TailApply(tmp12575, tmp12581)
return


}


}, 1)

tmp12584 := Call(__e, PrimFunc(symshen_4_5formulae_6), W3100)


__e.TailApply(tmp12574, tmp12584)
return


}, 1)

tmp12585 := Call(__e, PrimFunc(symshen_4in_1_6), W3099)


__e.TailApply(tmp12573, tmp12585)
return


}


}, 1)

tmp12588 := Call(__e, PrimFunc(symshen_4_5sc_6), W3098)


__e.TailApply(tmp12572, tmp12588)
return


}, 1)

tmp12589 := Call(__e, PrimFunc(symshen_4in_1_6), W3096)


__e.TailApply(tmp12571, tmp12589)
return


}, 1)

tmp12590 := Call(__e, PrimFunc(symshen_4_5_1out), W3096)


__e.TailApply(tmp12570, tmp12590)
return


}


}, 1)

tmp12593 := Call(__e, PrimFunc(symshen_4_5formula_6), V3094)


tmp12594 := Call(__e, tmp12569, tmp12593)


__e.TailApply(tmp12545, tmp12594)
return


}, 1)

tmp12595 := Call(__e, ns2_1set, symshen_4_5formulae_6, tmp12544)


_ = tmp12595

tmp12596 := MakeNative(func(__e *ControlFlow) {
V3110 := __e.Get(1)
_ = V3110
tmp12597 := MakeNative(func(__e *ControlFlow) {
W3111 := __e.Get(1)
_ = W3111
tmp12613 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3111)


if True == tmp12613 {
tmp12598 := MakeNative(func(__e *ControlFlow) {
W3119 := __e.Get(1)
_ = W3119
tmp12600 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3119)


if True == tmp12600 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3119)
return
}


}, 1)

tmp12601 := MakeNative(func(__e *ControlFlow) {
W3120 := __e.Get(1)
_ = W3120
tmp12609 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3120)


if True == tmp12609 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12602 := MakeNative(func(__e *ControlFlow) {
W3121 := __e.Get(1)
_ = W3121
tmp12603 := MakeNative(func(__e *ControlFlow) {
W3122 := __e.Get(1)
_ = W3122
tmp12604 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W3121, Nil)
}
__typedArg0 := W3121
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12605 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Nil, tmp12604)
}
__typedArg0 := Nil
__typedArg1 := tmp12604
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4comb), W3122, tmp12605)
return


}, 1)

tmp12606 := Call(__e, PrimFunc(symshen_4in_1_6), W3120)


__e.TailApply(tmp12603, tmp12606)
return


}, 1)

tmp12607 := Call(__e, PrimFunc(symshen_4_5_1out), W3120)


__e.TailApply(tmp12602, tmp12607)
return


}


}, 1)

tmp12610 := Call(__e, PrimFunc(symshen_4_5formula_6), V3110)


tmp12611 := Call(__e, tmp12601, tmp12610)


__e.TailApply(tmp12598, tmp12611)
return


} else {
__e.Return(W3111)
return
}


}, 1)

tmp12614 := MakeNative(func(__e *ControlFlow) {
W3112 := __e.Get(1)
_ = W3112
tmp12634 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3112)


if True == tmp12634 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12615 := MakeNative(func(__e *ControlFlow) {
W3113 := __e.Get(1)
_ = W3113
tmp12616 := MakeNative(func(__e *ControlFlow) {
W3114 := __e.Get(1)
_ = W3114
tmp12630 := Call(__e, PrimFunc(symshen_4hds_a_2), W3114, sym_6_6)


if True == tmp12630 {
tmp12617 := MakeNative(func(__e *ControlFlow) {
W3115 := __e.Get(1)
_ = W3115
tmp12618 := MakeNative(func(__e *ControlFlow) {
W3116 := __e.Get(1)
_ = W3116
tmp12626 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3116)


if True == tmp12626 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12619 := MakeNative(func(__e *ControlFlow) {
W3117 := __e.Get(1)
_ = W3117
tmp12620 := MakeNative(func(__e *ControlFlow) {
W3118 := __e.Get(1)
_ = W3118
tmp12621 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W3117, Nil)
}
__typedArg0 := W3117
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12622 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W3113, tmp12621)
}
__typedArg0 := W3113
__typedArg1 := tmp12621
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4comb), W3118, tmp12622)
return


}, 1)

tmp12623 := Call(__e, PrimFunc(symshen_4in_1_6), W3116)


__e.TailApply(tmp12620, tmp12623)
return


}, 1)

tmp12624 := Call(__e, PrimFunc(symshen_4_5_1out), W3116)


__e.TailApply(tmp12619, tmp12624)
return


}


}, 1)

tmp12627 := Call(__e, PrimFunc(symshen_4_5formula_6), W3115)


__e.TailApply(tmp12618, tmp12627)
return


}, 1)

tmp12628 := Call(__e, PrimFunc(symtail), W3114)


__e.TailApply(tmp12617, tmp12628)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp12631 := Call(__e, PrimFunc(symshen_4in_1_6), W3112)


__e.TailApply(tmp12616, tmp12631)
return


}, 1)

tmp12632 := Call(__e, PrimFunc(symshen_4_5_1out), W3112)


__e.TailApply(tmp12615, tmp12632)
return


}


}, 1)

tmp12635 := Call(__e, PrimFunc(symshen_4_5ass_6), V3110)


tmp12636 := Call(__e, tmp12614, tmp12635)


__e.TailApply(tmp12597, tmp12636)
return


}, 1)

tmp12637 := Call(__e, ns2_1set, symshen_4_5conc_6, tmp12596)


_ = tmp12637

tmp12638 := MakeNative(func(__e *ControlFlow) {
V3123 := __e.Get(1)
_ = V3123
tmp12639 := MakeNative(func(__e *ControlFlow) {
W3124 := __e.Get(1)
_ = W3124
tmp12651 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3124)


if True == tmp12651 {
tmp12640 := MakeNative(func(__e *ControlFlow) {
W3133 := __e.Get(1)
_ = W3133
tmp12642 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3133)


if True == tmp12642 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3133)
return
}


}, 1)

tmp12643 := MakeNative(func(__e *ControlFlow) {
W3134 := __e.Get(1)
_ = W3134
tmp12647 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3134)


if True == tmp12647 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12644 := MakeNative(func(__e *ControlFlow) {
W3135 := __e.Get(1)
_ = W3135
__e.TailApply(PrimFunc(symshen_4comb), W3135, Nil)
return
}, 1)

tmp12645 := Call(__e, PrimFunc(symshen_4in_1_6), W3134)


__e.TailApply(tmp12644, tmp12645)
return


}


}, 1)

tmp12648 := Call(__e, PrimFunc(sym_5e_6), V3123)


tmp12649 := Call(__e, tmp12643, tmp12648)


__e.TailApply(tmp12640, tmp12649)
return


} else {
__e.Return(W3124)
return
}


}, 1)

tmp12652 := MakeNative(func(__e *ControlFlow) {
W3125 := __e.Get(1)
_ = W3125
tmp12673 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3125)


if True == tmp12673 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12653 := MakeNative(func(__e *ControlFlow) {
W3126 := __e.Get(1)
_ = W3126
tmp12654 := MakeNative(func(__e *ControlFlow) {
W3127 := __e.Get(1)
_ = W3127
tmp12655 := MakeNative(func(__e *ControlFlow) {
W3128 := __e.Get(1)
_ = W3128
tmp12668 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3128)


if True == tmp12668 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12656 := MakeNative(func(__e *ControlFlow) {
W3129 := __e.Get(1)
_ = W3129
tmp12657 := MakeNative(func(__e *ControlFlow) {
W3130 := __e.Get(1)
_ = W3130
tmp12664 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3130)


if True == tmp12664 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12658 := MakeNative(func(__e *ControlFlow) {
W3131 := __e.Get(1)
_ = W3131
tmp12659 := MakeNative(func(__e *ControlFlow) {
W3132 := __e.Get(1)
_ = W3132
tmp12660 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W3126, W3131)
}
__typedArg0 := W3126
__typedArg1 := W3131
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4comb), W3132, tmp12660)
return


}, 1)

tmp12661 := Call(__e, PrimFunc(symshen_4in_1_6), W3130)


__e.TailApply(tmp12659, tmp12661)
return


}, 1)

tmp12662 := Call(__e, PrimFunc(symshen_4_5_1out), W3130)


__e.TailApply(tmp12658, tmp12662)
return


}


}, 1)

tmp12665 := Call(__e, PrimFunc(symshen_4_5prems_6), W3129)


__e.TailApply(tmp12657, tmp12665)
return


}, 1)

tmp12666 := Call(__e, PrimFunc(symshen_4in_1_6), W3128)


__e.TailApply(tmp12656, tmp12666)
return


}


}, 1)

tmp12669 := Call(__e, PrimFunc(symshen_4_5sc_6), W3127)


__e.TailApply(tmp12655, tmp12669)
return


}, 1)

tmp12670 := Call(__e, PrimFunc(symshen_4in_1_6), W3125)


__e.TailApply(tmp12654, tmp12670)
return


}, 1)

tmp12671 := Call(__e, PrimFunc(symshen_4_5_1out), W3125)


__e.TailApply(tmp12653, tmp12671)
return


}


}, 1)

tmp12674 := Call(__e, PrimFunc(symshen_4_5prem_6), V3123)


tmp12675 := Call(__e, tmp12652, tmp12674)


__e.TailApply(tmp12639, tmp12675)
return


}, 1)

tmp12676 := Call(__e, ns2_1set, symshen_4_5prems_6, tmp12638)


_ = tmp12676

tmp12677 := MakeNative(func(__e *ControlFlow) {
V3136 := __e.Get(1)
_ = V3136
tmp12678 := MakeNative(func(__e *ControlFlow) {
W3137 := __e.Get(1)
_ = W3137
tmp12720 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3137)


if True == tmp12720 {
tmp12679 := MakeNative(func(__e *ControlFlow) {
W3139 := __e.Get(1)
_ = W3139
tmp12695 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3139)


if True == tmp12695 {
tmp12680 := MakeNative(func(__e *ControlFlow) {
W3147 := __e.Get(1)
_ = W3147
tmp12682 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3147)


if True == tmp12682 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3147)
return
}


}, 1)

tmp12683 := MakeNative(func(__e *ControlFlow) {
W3148 := __e.Get(1)
_ = W3148
tmp12691 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3148)


if True == tmp12691 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12684 := MakeNative(func(__e *ControlFlow) {
W3149 := __e.Get(1)
_ = W3149
tmp12685 := MakeNative(func(__e *ControlFlow) {
W3150 := __e.Get(1)
_ = W3150
tmp12686 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W3149, Nil)
}
__typedArg0 := W3149
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12687 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Nil, tmp12686)
}
__typedArg0 := Nil
__typedArg1 := tmp12686
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4comb), W3150, tmp12687)
return


}, 1)

tmp12688 := Call(__e, PrimFunc(symshen_4in_1_6), W3148)


__e.TailApply(tmp12685, tmp12688)
return


}, 1)

tmp12689 := Call(__e, PrimFunc(symshen_4_5_1out), W3148)


__e.TailApply(tmp12684, tmp12689)
return


}


}, 1)

tmp12692 := Call(__e, PrimFunc(symshen_4_5formula_6), V3136)


tmp12693 := Call(__e, tmp12683, tmp12692)


__e.TailApply(tmp12680, tmp12693)
return


} else {
__e.Return(W3139)
return
}


}, 1)

tmp12696 := MakeNative(func(__e *ControlFlow) {
W3140 := __e.Get(1)
_ = W3140
tmp12716 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3140)


if True == tmp12716 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12697 := MakeNative(func(__e *ControlFlow) {
W3141 := __e.Get(1)
_ = W3141
tmp12698 := MakeNative(func(__e *ControlFlow) {
W3142 := __e.Get(1)
_ = W3142
tmp12712 := Call(__e, PrimFunc(symshen_4hds_a_2), W3142, sym_6_6)


if True == tmp12712 {
tmp12699 := MakeNative(func(__e *ControlFlow) {
W3143 := __e.Get(1)
_ = W3143
tmp12700 := MakeNative(func(__e *ControlFlow) {
W3144 := __e.Get(1)
_ = W3144
tmp12708 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3144)


if True == tmp12708 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12701 := MakeNative(func(__e *ControlFlow) {
W3145 := __e.Get(1)
_ = W3145
tmp12702 := MakeNative(func(__e *ControlFlow) {
W3146 := __e.Get(1)
_ = W3146
tmp12703 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W3145, Nil)
}
__typedArg0 := W3145
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12704 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W3141, tmp12703)
}
__typedArg0 := W3141
__typedArg1 := tmp12703
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4comb), W3146, tmp12704)
return


}, 1)

tmp12705 := Call(__e, PrimFunc(symshen_4in_1_6), W3144)


__e.TailApply(tmp12702, tmp12705)
return


}, 1)

tmp12706 := Call(__e, PrimFunc(symshen_4_5_1out), W3144)


__e.TailApply(tmp12701, tmp12706)
return


}


}, 1)

tmp12709 := Call(__e, PrimFunc(symshen_4_5formula_6), W3143)


__e.TailApply(tmp12700, tmp12709)
return


}, 1)

tmp12710 := Call(__e, PrimFunc(symtail), W3142)


__e.TailApply(tmp12699, tmp12710)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp12713 := Call(__e, PrimFunc(symshen_4in_1_6), W3140)


__e.TailApply(tmp12698, tmp12713)
return


}, 1)

tmp12714 := Call(__e, PrimFunc(symshen_4_5_1out), W3140)


__e.TailApply(tmp12697, tmp12714)
return


}


}, 1)

tmp12717 := Call(__e, PrimFunc(symshen_4_5ass_6), V3136)


tmp12718 := Call(__e, tmp12696, tmp12717)


__e.TailApply(tmp12679, tmp12718)
return


} else {
__e.Return(W3137)
return
}


}, 1)

tmp12726 := Call(__e, PrimFunc(symshen_4hds_a_2), V3136, sym_b)


var ifres12721 Obj

if True == tmp12726 {
tmp12722 := MakeNative(func(__e *ControlFlow) {
W3138 := __e.Get(1)
_ = W3138
__e.TailApply(PrimFunc(symshen_4comb), W3138, sym_b)
return
}, 1)

tmp12723 := Call(__e, PrimFunc(symtail), V3136)


tmp12724 := Call(__e, tmp12722, tmp12723)


ifres12721 = tmp12724


} else {
tmp12725 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12721 = tmp12725


}

__e.TailApply(tmp12678, ifres12721)
return


}, 1)

tmp12727 := Call(__e, ns2_1set, symshen_4_5prem_6, tmp12677)


_ = tmp12727

tmp12728 := MakeNative(func(__e *ControlFlow) {
V3151 := __e.Get(1)
_ = V3151
tmp12729 := MakeNative(func(__e *ControlFlow) {
W3152 := __e.Get(1)
_ = W3152
tmp12754 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3152)


if True == tmp12754 {
tmp12730 := MakeNative(func(__e *ControlFlow) {
W3161 := __e.Get(1)
_ = W3161
tmp12742 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3161)


if True == tmp12742 {
tmp12731 := MakeNative(func(__e *ControlFlow) {
W3165 := __e.Get(1)
_ = W3165
tmp12733 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3165)


if True == tmp12733 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3165)
return
}


}, 1)

tmp12734 := MakeNative(func(__e *ControlFlow) {
W3166 := __e.Get(1)
_ = W3166
tmp12738 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3166)


if True == tmp12738 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12735 := MakeNative(func(__e *ControlFlow) {
W3167 := __e.Get(1)
_ = W3167
__e.TailApply(PrimFunc(symshen_4comb), W3167, Nil)
return
}, 1)

tmp12736 := Call(__e, PrimFunc(symshen_4in_1_6), W3166)


__e.TailApply(tmp12735, tmp12736)
return


}


}, 1)

tmp12739 := Call(__e, PrimFunc(sym_5e_6), V3151)


tmp12740 := Call(__e, tmp12734, tmp12739)


__e.TailApply(tmp12731, tmp12740)
return


} else {
__e.Return(W3161)
return
}


}, 1)

tmp12743 := MakeNative(func(__e *ControlFlow) {
W3162 := __e.Get(1)
_ = W3162
tmp12750 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3162)


if True == tmp12750 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12744 := MakeNative(func(__e *ControlFlow) {
W3163 := __e.Get(1)
_ = W3163
tmp12745 := MakeNative(func(__e *ControlFlow) {
W3164 := __e.Get(1)
_ = W3164
tmp12746 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W3163, Nil)
}
__typedArg0 := W3163
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4comb), W3164, tmp12746)
return


}, 1)

tmp12747 := Call(__e, PrimFunc(symshen_4in_1_6), W3162)


__e.TailApply(tmp12745, tmp12747)
return


}, 1)

tmp12748 := Call(__e, PrimFunc(symshen_4_5_1out), W3162)


__e.TailApply(tmp12744, tmp12748)
return


}


}, 1)

tmp12751 := Call(__e, PrimFunc(symshen_4_5formula_6), V3151)


tmp12752 := Call(__e, tmp12743, tmp12751)


__e.TailApply(tmp12730, tmp12752)
return


} else {
__e.Return(W3152)
return
}


}, 1)

tmp12755 := MakeNative(func(__e *ControlFlow) {
W3153 := __e.Get(1)
_ = W3153
tmp12776 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3153)


if True == tmp12776 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12756 := MakeNative(func(__e *ControlFlow) {
W3154 := __e.Get(1)
_ = W3154
tmp12757 := MakeNative(func(__e *ControlFlow) {
W3155 := __e.Get(1)
_ = W3155
tmp12758 := MakeNative(func(__e *ControlFlow) {
W3156 := __e.Get(1)
_ = W3156
tmp12771 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3156)


if True == tmp12771 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12759 := MakeNative(func(__e *ControlFlow) {
W3157 := __e.Get(1)
_ = W3157
tmp12760 := MakeNative(func(__e *ControlFlow) {
W3158 := __e.Get(1)
_ = W3158
tmp12767 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3158)


if True == tmp12767 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12761 := MakeNative(func(__e *ControlFlow) {
W3159 := __e.Get(1)
_ = W3159
tmp12762 := MakeNative(func(__e *ControlFlow) {
W3160 := __e.Get(1)
_ = W3160
tmp12763 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W3154, W3159)
}
__typedArg0 := W3154
__typedArg1 := W3159
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4comb), W3160, tmp12763)
return


}, 1)

tmp12764 := Call(__e, PrimFunc(symshen_4in_1_6), W3158)


__e.TailApply(tmp12762, tmp12764)
return


}, 1)

tmp12765 := Call(__e, PrimFunc(symshen_4_5_1out), W3158)


__e.TailApply(tmp12761, tmp12765)
return


}


}, 1)

tmp12768 := Call(__e, PrimFunc(symshen_4_5ass_6), W3157)


__e.TailApply(tmp12760, tmp12768)
return


}, 1)

tmp12769 := Call(__e, PrimFunc(symshen_4in_1_6), W3156)


__e.TailApply(tmp12759, tmp12769)
return


}


}, 1)

tmp12772 := Call(__e, PrimFunc(symshen_4_5iscomma_6), W3155)


__e.TailApply(tmp12758, tmp12772)
return


}, 1)

tmp12773 := Call(__e, PrimFunc(symshen_4in_1_6), W3153)


__e.TailApply(tmp12757, tmp12773)
return


}, 1)

tmp12774 := Call(__e, PrimFunc(symshen_4_5_1out), W3153)


__e.TailApply(tmp12756, tmp12774)
return


}


}, 1)

tmp12777 := Call(__e, PrimFunc(symshen_4_5formula_6), V3151)


tmp12778 := Call(__e, tmp12755, tmp12777)


__e.TailApply(tmp12729, tmp12778)
return


}, 1)

tmp12779 := Call(__e, ns2_1set, symshen_4_5ass_6, tmp12728)


_ = tmp12779

tmp12780 := MakeNative(func(__e *ControlFlow) {
V3168 := __e.Get(1)
_ = V3168
tmp12781 := MakeNative(func(__e *ControlFlow) {
W3169 := __e.Get(1)
_ = W3169
tmp12783 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3169)


if True == tmp12783 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3169)
return
}


}, 1)

tmp12794 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3168)
}
__typedArg0 := V3168
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12784 Obj

if True == tmp12794 {
tmp12785 := MakeNative(func(__e *ControlFlow) {
W3170 := __e.Get(1)
_ = W3170
tmp12786 := MakeNative(func(__e *ControlFlow) {
W3171 := __e.Get(1)
_ = W3171
tmp12788 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(","))
}
__typedArg0 := MakeString(",")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp12789 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W3170, tmp12788)
}
__typedArg0 := W3170
__typedArg1 := tmp12788
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp12789 {
__e.TailApply(PrimFunc(symshen_4comb), W3171, symshen_4skip)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp12790 := Call(__e, PrimFunc(symtail), V3168)


__e.TailApply(tmp12786, tmp12790)
return


}, 1)

tmp12791 := Call(__e, PrimFunc(symhead), V3168)


tmp12792 := Call(__e, tmp12785, tmp12791)


ifres12784 = tmp12792


} else {
tmp12793 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12784 = tmp12793


}

__e.TailApply(tmp12781, ifres12784)
return


}, 1)

tmp12795 := Call(__e, ns2_1set, symshen_4_5iscomma_6, tmp12780)


_ = tmp12795

tmp12796 := MakeNative(func(__e *ControlFlow) {
V3172 := __e.Get(1)
_ = V3172
tmp12797 := MakeNative(func(__e *ControlFlow) {
W3173 := __e.Get(1)
_ = W3173
tmp12811 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3173)


if True == tmp12811 {
tmp12798 := MakeNative(func(__e *ControlFlow) {
W3182 := __e.Get(1)
_ = W3182
tmp12800 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3182)


if True == tmp12800 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3182)
return
}


}, 1)

tmp12801 := MakeNative(func(__e *ControlFlow) {
W3183 := __e.Get(1)
_ = W3183
tmp12807 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3183)


if True == tmp12807 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12802 := MakeNative(func(__e *ControlFlow) {
W3184 := __e.Get(1)
_ = W3184
tmp12803 := MakeNative(func(__e *ControlFlow) {
W3185 := __e.Get(1)
_ = W3185
__e.TailApply(PrimFunc(symshen_4comb), W3185, W3184)
return
}, 1)

tmp12804 := Call(__e, PrimFunc(symshen_4in_1_6), W3183)


__e.TailApply(tmp12803, tmp12804)
return


}, 1)

tmp12805 := Call(__e, PrimFunc(symshen_4_5_1out), W3183)


__e.TailApply(tmp12802, tmp12805)
return


}


}, 1)

tmp12808 := Call(__e, PrimFunc(symshen_4_5expr_6), V3172)


tmp12809 := Call(__e, tmp12801, tmp12808)


__e.TailApply(tmp12798, tmp12809)
return


} else {
__e.Return(W3173)
return
}


}, 1)

tmp12812 := MakeNative(func(__e *ControlFlow) {
W3174 := __e.Get(1)
_ = W3174
tmp12838 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3174)


if True == tmp12838 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12813 := MakeNative(func(__e *ControlFlow) {
W3175 := __e.Get(1)
_ = W3175
tmp12814 := MakeNative(func(__e *ControlFlow) {
W3176 := __e.Get(1)
_ = W3176
tmp12815 := MakeNative(func(__e *ControlFlow) {
W3177 := __e.Get(1)
_ = W3177
tmp12833 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3177)


if True == tmp12833 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12816 := MakeNative(func(__e *ControlFlow) {
W3178 := __e.Get(1)
_ = W3178
tmp12817 := MakeNative(func(__e *ControlFlow) {
W3179 := __e.Get(1)
_ = W3179
tmp12829 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3179)


if True == tmp12829 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12818 := MakeNative(func(__e *ControlFlow) {
W3180 := __e.Get(1)
_ = W3180
tmp12819 := MakeNative(func(__e *ControlFlow) {
W3181 := __e.Get(1)
_ = W3181
tmp12820 := Call(__e, PrimFunc(symshen_4curry), W3175)


tmp12821 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp12822 := Call(__e, PrimFunc(symshen_4rectify_1type), W3180)


tmp12823 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12822, Nil)
}
__typedArg0 := tmp12822
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12824 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12821, tmp12823)
}
__typedArg0 := tmp12821
__typedArg1 := tmp12823
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12825 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12820, tmp12824)
}
__typedArg0 := tmp12820
__typedArg1 := tmp12824
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4comb), W3181, tmp12825)
return


}, 1)

tmp12826 := Call(__e, PrimFunc(symshen_4in_1_6), W3179)


__e.TailApply(tmp12819, tmp12826)
return


}, 1)

tmp12827 := Call(__e, PrimFunc(symshen_4_5_1out), W3179)


__e.TailApply(tmp12818, tmp12827)
return


}


}, 1)

tmp12830 := Call(__e, PrimFunc(symshen_4_5type_6), W3178)


__e.TailApply(tmp12817, tmp12830)
return


}, 1)

tmp12831 := Call(__e, PrimFunc(symshen_4in_1_6), W3177)


__e.TailApply(tmp12816, tmp12831)
return


}


}, 1)

tmp12834 := Call(__e, PrimFunc(symshen_4_5iscolon_6), W3176)


__e.TailApply(tmp12815, tmp12834)
return


}, 1)

tmp12835 := Call(__e, PrimFunc(symshen_4in_1_6), W3174)


__e.TailApply(tmp12814, tmp12835)
return


}, 1)

tmp12836 := Call(__e, PrimFunc(symshen_4_5_1out), W3174)


__e.TailApply(tmp12813, tmp12836)
return


}


}, 1)

tmp12839 := Call(__e, PrimFunc(symshen_4_5expr_6), V3172)


tmp12840 := Call(__e, tmp12812, tmp12839)


__e.TailApply(tmp12797, tmp12840)
return


}, 1)

tmp12841 := Call(__e, ns2_1set, symshen_4_5formula_6, tmp12796)


_ = tmp12841

tmp12842 := MakeNative(func(__e *ControlFlow) {
V3186 := __e.Get(1)
_ = V3186
tmp12843 := MakeNative(func(__e *ControlFlow) {
W3187 := __e.Get(1)
_ = W3187
tmp12845 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3187)


if True == tmp12845 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3187)
return
}


}, 1)

tmp12856 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3186)
}
__typedArg0 := V3186
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12846 Obj

if True == tmp12856 {
tmp12847 := MakeNative(func(__e *ControlFlow) {
W3188 := __e.Get(1)
_ = W3188
tmp12848 := MakeNative(func(__e *ControlFlow) {
W3189 := __e.Get(1)
_ = W3189
tmp12850 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp12851 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W3188, tmp12850)
}
__typedArg0 := W3188
__typedArg1 := tmp12850
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp12851 {
__e.TailApply(PrimFunc(symshen_4comb), W3189, symshen_4skip)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp12852 := Call(__e, PrimFunc(symtail), V3186)


__e.TailApply(tmp12848, tmp12852)
return


}, 1)

tmp12853 := Call(__e, PrimFunc(symhead), V3186)


tmp12854 := Call(__e, tmp12847, tmp12853)


ifres12846 = tmp12854


} else {
tmp12855 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12846 = tmp12855


}

__e.TailApply(tmp12843, ifres12846)
return


}, 1)

tmp12857 := Call(__e, ns2_1set, symshen_4_5iscolon_6, tmp12842)


_ = tmp12857

tmp12858 := MakeNative(func(__e *ControlFlow) {
V3190 := __e.Get(1)
_ = V3190
tmp12859 := MakeNative(func(__e *ControlFlow) {
W3191 := __e.Get(1)
_ = W3191
tmp12871 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3191)


if True == tmp12871 {
tmp12860 := MakeNative(func(__e *ControlFlow) {
W3198 := __e.Get(1)
_ = W3198
tmp12862 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3198)


if True == tmp12862 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3198)
return
}


}, 1)

tmp12863 := MakeNative(func(__e *ControlFlow) {
W3199 := __e.Get(1)
_ = W3199
tmp12867 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3199)


if True == tmp12867 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12864 := MakeNative(func(__e *ControlFlow) {
W3200 := __e.Get(1)
_ = W3200
__e.TailApply(PrimFunc(symshen_4comb), W3200, Nil)
return
}, 1)

tmp12865 := Call(__e, PrimFunc(symshen_4in_1_6), W3199)


__e.TailApply(tmp12864, tmp12865)
return


}


}, 1)

tmp12868 := Call(__e, PrimFunc(sym_5e_6), V3190)


tmp12869 := Call(__e, tmp12863, tmp12868)


__e.TailApply(tmp12860, tmp12869)
return


} else {
__e.Return(W3191)
return
}


}, 1)

tmp12872 := MakeNative(func(__e *ControlFlow) {
W3192 := __e.Get(1)
_ = W3192
tmp12887 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3192)


if True == tmp12887 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12873 := MakeNative(func(__e *ControlFlow) {
W3193 := __e.Get(1)
_ = W3193
tmp12874 := MakeNative(func(__e *ControlFlow) {
W3194 := __e.Get(1)
_ = W3194
tmp12875 := MakeNative(func(__e *ControlFlow) {
W3195 := __e.Get(1)
_ = W3195
tmp12882 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3195)


if True == tmp12882 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12876 := MakeNative(func(__e *ControlFlow) {
W3196 := __e.Get(1)
_ = W3196
tmp12877 := MakeNative(func(__e *ControlFlow) {
W3197 := __e.Get(1)
_ = W3197
tmp12878 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W3193, W3196)
}
__typedArg0 := W3193
__typedArg1 := W3196
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4comb), W3197, tmp12878)
return


}, 1)

tmp12879 := Call(__e, PrimFunc(symshen_4in_1_6), W3195)


__e.TailApply(tmp12877, tmp12879)
return


}, 1)

tmp12880 := Call(__e, PrimFunc(symshen_4_5_1out), W3195)


__e.TailApply(tmp12876, tmp12880)
return


}


}, 1)

tmp12883 := Call(__e, PrimFunc(symshen_4_5sides_6), W3194)


__e.TailApply(tmp12875, tmp12883)
return


}, 1)

tmp12884 := Call(__e, PrimFunc(symshen_4in_1_6), W3192)


__e.TailApply(tmp12874, tmp12884)
return


}, 1)

tmp12885 := Call(__e, PrimFunc(symshen_4_5_1out), W3192)


__e.TailApply(tmp12873, tmp12885)
return


}


}, 1)

tmp12888 := Call(__e, PrimFunc(symshen_4_5side_6), V3190)


tmp12889 := Call(__e, tmp12872, tmp12888)


__e.TailApply(tmp12859, tmp12889)
return


}, 1)

tmp12890 := Call(__e, ns2_1set, symshen_4_5sides_6, tmp12858)


_ = tmp12890

tmp12891 := MakeNative(func(__e *ControlFlow) {
V3201 := __e.Get(1)
_ = V3201
tmp12892 := MakeNative(func(__e *ControlFlow) {
W3202 := __e.Get(1)
_ = W3202
tmp12956 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3202)


if True == tmp12956 {
tmp12893 := MakeNative(func(__e *ControlFlow) {
W3206 := __e.Get(1)
_ = W3206
tmp12933 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3206)


if True == tmp12933 {
tmp12894 := MakeNative(func(__e *ControlFlow) {
W3212 := __e.Get(1)
_ = W3212
tmp12915 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3212)


if True == tmp12915 {
tmp12895 := MakeNative(func(__e *ControlFlow) {
W3216 := __e.Get(1)
_ = W3216
tmp12897 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3216)


if True == tmp12897 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3216)
return
}


}, 1)

tmp12913 := Call(__e, PrimFunc(symshen_4hds_a_2), V3201, symsqts)


var ifres12898 Obj

if True == tmp12913 {
tmp12899 := MakeNative(func(__e *ControlFlow) {
W3217 := __e.Get(1)
_ = W3217
tmp12909 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W3217)
}
__typedArg0 := W3217
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp12909 {
tmp12900 := MakeNative(func(__e *ControlFlow) {
W3218 := __e.Get(1)
_ = W3218
tmp12901 := MakeNative(func(__e *ControlFlow) {
W3219 := __e.Get(1)
_ = W3219
tmp12905 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(W3218)
}
__typedArg0 := W3218
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

if True == tmp12905 {
tmp12902 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W3218, Nil)
}
__typedArg0 := W3218
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12903 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsqts, tmp12902)
}
__typedArg0 := symsqts
__typedArg1 := tmp12902
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4comb), W3219, tmp12903)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp12906 := Call(__e, PrimFunc(symtail), W3217)


__e.TailApply(tmp12901, tmp12906)
return


}, 1)

tmp12907 := Call(__e, PrimFunc(symhead), W3217)


__e.TailApply(tmp12900, tmp12907)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp12910 := Call(__e, PrimFunc(symtail), V3201)


tmp12911 := Call(__e, tmp12899, tmp12910)


ifres12898 = tmp12911


} else {
tmp12912 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12898 = tmp12912


}

__e.TailApply(tmp12895, ifres12898)
return


} else {
__e.Return(W3212)
return
}


}, 1)

tmp12931 := Call(__e, PrimFunc(symshen_4hds_a_2), V3201, symctxt)


var ifres12916 Obj

if True == tmp12931 {
tmp12917 := MakeNative(func(__e *ControlFlow) {
W3213 := __e.Get(1)
_ = W3213
tmp12927 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W3213)
}
__typedArg0 := W3213
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp12927 {
tmp12918 := MakeNative(func(__e *ControlFlow) {
W3214 := __e.Get(1)
_ = W3214
tmp12919 := MakeNative(func(__e *ControlFlow) {
W3215 := __e.Get(1)
_ = W3215
tmp12923 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(W3214)
}
__typedArg0 := W3214
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

if True == tmp12923 {
tmp12920 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W3214, Nil)
}
__typedArg0 := W3214
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12921 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symctxt, tmp12920)
}
__typedArg0 := symctxt
__typedArg1 := tmp12920
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4comb), W3215, tmp12921)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp12924 := Call(__e, PrimFunc(symtail), W3213)


__e.TailApply(tmp12919, tmp12924)
return


}, 1)

tmp12925 := Call(__e, PrimFunc(symhead), W3213)


__e.TailApply(tmp12918, tmp12925)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp12928 := Call(__e, PrimFunc(symtail), V3201)


tmp12929 := Call(__e, tmp12917, tmp12928)


ifres12916 = tmp12929


} else {
tmp12930 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12916 = tmp12930


}

__e.TailApply(tmp12894, ifres12916)
return


} else {
__e.Return(W3206)
return
}


}, 1)

tmp12954 := Call(__e, PrimFunc(symshen_4hds_a_2), V3201, symlet)


var ifres12934 Obj

if True == tmp12954 {
tmp12935 := MakeNative(func(__e *ControlFlow) {
W3207 := __e.Get(1)
_ = W3207
tmp12950 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W3207)
}
__typedArg0 := W3207
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp12950 {
tmp12936 := MakeNative(func(__e *ControlFlow) {
W3208 := __e.Get(1)
_ = W3208
tmp12937 := MakeNative(func(__e *ControlFlow) {
W3209 := __e.Get(1)
_ = W3209
tmp12946 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W3209)
}
__typedArg0 := W3209
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp12946 {
tmp12938 := MakeNative(func(__e *ControlFlow) {
W3210 := __e.Get(1)
_ = W3210
tmp12939 := MakeNative(func(__e *ControlFlow) {
W3211 := __e.Get(1)
_ = W3211
tmp12940 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W3210, Nil)
}
__typedArg0 := W3210
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12941 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W3208, tmp12940)
}
__typedArg0 := W3208
__typedArg1 := tmp12940
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12942 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp12941)
}
__typedArg0 := symlet
__typedArg1 := tmp12941
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4comb), W3211, tmp12942)
return


}, 1)

tmp12943 := Call(__e, PrimFunc(symtail), W3209)


__e.TailApply(tmp12939, tmp12943)
return


}, 1)

tmp12944 := Call(__e, PrimFunc(symhead), W3209)


__e.TailApply(tmp12938, tmp12944)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp12947 := Call(__e, PrimFunc(symtail), W3207)


__e.TailApply(tmp12937, tmp12947)
return


}, 1)

tmp12948 := Call(__e, PrimFunc(symhead), W3207)


__e.TailApply(tmp12936, tmp12948)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp12951 := Call(__e, PrimFunc(symtail), V3201)


tmp12952 := Call(__e, tmp12935, tmp12951)


ifres12934 = tmp12952


} else {
tmp12953 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12934 = tmp12953


}

__e.TailApply(tmp12893, ifres12934)
return


} else {
__e.Return(W3202)
return
}


}, 1)

tmp12970 := Call(__e, PrimFunc(symshen_4hds_a_2), V3201, symif)


var ifres12957 Obj

if True == tmp12970 {
tmp12958 := MakeNative(func(__e *ControlFlow) {
W3203 := __e.Get(1)
_ = W3203
tmp12966 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W3203)
}
__typedArg0 := W3203
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp12966 {
tmp12959 := MakeNative(func(__e *ControlFlow) {
W3204 := __e.Get(1)
_ = W3204
tmp12960 := MakeNative(func(__e *ControlFlow) {
W3205 := __e.Get(1)
_ = W3205
tmp12961 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W3204, Nil)
}
__typedArg0 := W3204
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12962 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp12961)
}
__typedArg0 := symif
__typedArg1 := tmp12961
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4comb), W3205, tmp12962)
return


}, 1)

tmp12963 := Call(__e, PrimFunc(symtail), W3203)


__e.TailApply(tmp12960, tmp12963)
return


}, 1)

tmp12964 := Call(__e, PrimFunc(symhead), W3203)


__e.TailApply(tmp12959, tmp12964)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp12967 := Call(__e, PrimFunc(symtail), V3201)


tmp12968 := Call(__e, tmp12958, tmp12967)


ifres12957 = tmp12968


} else {
tmp12969 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12957 = tmp12969


}

__e.TailApply(tmp12892, ifres12957)
return


}, 1)

tmp12971 := Call(__e, ns2_1set, symshen_4_5side_6, tmp12891)


_ = tmp12971

tmp12972 := MakeNative(func(__e *ControlFlow) {
V3226 := __e.Get(1)
_ = V3226
V3227 := __e.Get(2)
_ = V3227
V3228 := __e.Get(3)
_ = V3228
tmp13007 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3228)
}
__typedArg0 := V3228
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12994 Obj

if True == tmp13007 {
tmp13005 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3228)
}
__typedArg0 := V3228
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13006 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp13005)
}
__typedArg0 := Nil
__typedArg1 := tmp13005
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres12996 Obj

if True == tmp13006 {
tmp13003 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3228)
}
__typedArg0 := V3228
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13004 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13003)
}
__typedArg0 := tmp13003
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12998 Obj

if True == tmp13004 {
tmp13000 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3228)
}
__typedArg0 := V3228
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13001 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13000)
}
__typedArg0 := tmp13000
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13002 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp13001)
}
__typedArg0 := Nil
__typedArg1 := tmp13001
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres12999 Obj

if True == tmp13002 {
ifres12999 = True


} else {
ifres12999 = False


}

ifres12998 = ifres12999


} else {
ifres12998 = False


}

var ifres12997 Obj

if True == ifres12998 {
ifres12997 = True


} else {
ifres12997 = False


}

ifres12996 = ifres12997


} else {
ifres12996 = False


}

var ifres12995 Obj

if True == ifres12996 {
ifres12995 = True


} else {
ifres12995 = False


}

ifres12994 = ifres12995


} else {
ifres12994 = False


}

if True == ifres12994 {
tmp12973 := MakeNative(func(__e *ControlFlow) {
W3229 := __e.Get(1)
_ = W3229
tmp12974 := MakeNative(func(__e *ControlFlow) {
W3230 := __e.Get(1)
_ = W3230
tmp12975 := MakeNative(func(__e *ControlFlow) {
W3231 := __e.Get(1)
_ = W3231
tmp12976 := MakeNative(func(__e *ControlFlow) {
W3232 := __e.Get(1)
_ = W3232
tmp12977 := MakeNative(func(__e *ControlFlow) {
W3233 := __e.Get(1)
_ = W3233
tmp12978 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W3232, Nil)
}
__typedArg0 := W3232
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W3233, tmp12978)
}
__typedArg0 := W3233
__typedArg1 := tmp12978
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp12979 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3228, Nil)
}
__typedArg0 := V3228
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12980 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3227, tmp12979)
}
__typedArg0 := V3227
__typedArg1 := tmp12979
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12981 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3226, tmp12980)
}
__typedArg0 := V3226
__typedArg1 := tmp12980
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(tmp12977, tmp12981)
return


}, 1)

tmp12982 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W3231, Nil)
}
__typedArg0 := W3231
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12983 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W3230, Nil)
}
__typedArg0 := W3230
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12984 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12982, tmp12983)
}
__typedArg0 := tmp12982
__typedArg1 := tmp12983
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12985 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3226, tmp12984)
}
__typedArg0 := V3226
__typedArg1 := tmp12984
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(tmp12976, tmp12985)
return


}, 1)

tmp12986 := Call(__e, PrimFunc(symshen_4coll_1formulae), V3227)


tmp12987 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W3229, Nil)
}
__typedArg0 := W3229
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12988 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12986, tmp12987)
}
__typedArg0 := tmp12986
__typedArg1 := tmp12987
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(tmp12975, tmp12988)
return


}, 1)

tmp12989 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3228)
}
__typedArg0 := V3228
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12990 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W3229, Nil)
}
__typedArg0 := W3229
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12991 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12989, tmp12990)
}
__typedArg0 := tmp12989
__typedArg1 := tmp12990
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(tmp12974, tmp12991)
return


}, 1)

tmp12992 := Call(__e, PrimFunc(symgensym), symP)


__e.TailApply(tmp12973, tmp12992)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.lr-rule"))
}
__typedArg0 := MakeString("implementation error in shen.lr-rule")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}, 3)

tmp13008 := Call(__e, ns2_1set, symshen_4lr_1rule, tmp12972)


_ = tmp13008

tmp13009 := MakeNative(func(__e *ControlFlow) {
V3236 := __e.Get(1)
_ = V3236
tmp13038 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3236)
}
__typedArg0 := Nil
__typedArg1 := V3236
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13038 {
__e.Return(Nil)
return
} else {
tmp13036 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3236)
}
__typedArg0 := V3236
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13016 Obj

if True == tmp13036 {
tmp13034 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3236)
}
__typedArg0 := V3236
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13035 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13034)
}
__typedArg0 := tmp13034
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13018 Obj

if True == tmp13035 {
tmp13031 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3236)
}
__typedArg0 := V3236
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13032 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13031)
}
__typedArg0 := tmp13031
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13033 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp13032)
}
__typedArg0 := Nil
__typedArg1 := tmp13032
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13020 Obj

if True == tmp13033 {
tmp13028 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3236)
}
__typedArg0 := V3236
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13029 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13028)
}
__typedArg0 := tmp13028
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13030 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13029)
}
__typedArg0 := tmp13029
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13022 Obj

if True == tmp13030 {
tmp13024 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3236)
}
__typedArg0 := V3236
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13025 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13024)
}
__typedArg0 := tmp13024
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13026 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13025)
}
__typedArg0 := tmp13025
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13027 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp13026)
}
__typedArg0 := Nil
__typedArg1 := tmp13026
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13023 Obj

if True == tmp13027 {
ifres13023 = True


} else {
ifres13023 = False


}

ifres13022 = ifres13023


} else {
ifres13022 = False


}

var ifres13021 Obj

if True == ifres13022 {
ifres13021 = True


} else {
ifres13021 = False


}

ifres13020 = ifres13021


} else {
ifres13020 = False


}

var ifres13019 Obj

if True == ifres13020 {
ifres13019 = True


} else {
ifres13019 = False


}

ifres13018 = ifres13019


} else {
ifres13018 = False


}

var ifres13017 Obj

if True == ifres13018 {
ifres13017 = True


} else {
ifres13017 = False


}

ifres13016 = ifres13017


} else {
ifres13016 = False


}

if True == ifres13016 {
tmp13010 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3236)
}
__typedArg0 := V3236
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13011 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13010)
}
__typedArg0 := tmp13010
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13012 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13011)
}
__typedArg0 := tmp13011
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13013 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3236)
}
__typedArg0 := V3236
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13014 := Call(__e, PrimFunc(symshen_4coll_1formulae), tmp13013)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13012, tmp13014)
}
__typedArg0 := tmp13012
__typedArg1 := tmp13014
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.coll-formulae"))
}
__typedArg0 := MakeString("implementation error in shen.coll-formulae")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 1)

tmp13039 := Call(__e, ns2_1set, symshen_4coll_1formulae, tmp13009)


_ = tmp13039

tmp13040 := MakeNative(func(__e *ControlFlow) {
V3237 := __e.Get(1)
_ = V3237
tmp13041 := MakeNative(func(__e *ControlFlow) {
W3238 := __e.Get(1)
_ = W3238
tmp13043 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3238)


if True == tmp13043 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3238)
return
}


}, 1)

tmp13055 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3237)
}
__typedArg0 := V3237
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13044 Obj

if True == tmp13055 {
tmp13045 := MakeNative(func(__e *ControlFlow) {
W3239 := __e.Get(1)
_ = W3239
tmp13046 := MakeNative(func(__e *ControlFlow) {
W3240 := __e.Get(1)
_ = W3240
tmp13049 := Call(__e, PrimFunc(symshen_4key_1in_1sequent_1calculus_2), W3239)


if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp13049)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp13049
return Call(__e, PrimFunc(symnot), __typedArg0)
})() {
tmp13047 := Call(__e, PrimFunc(symmacroexpand), W3239)


__e.TailApply(PrimFunc(symshen_4comb), W3240, tmp13047)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp13051 := Call(__e, PrimFunc(symtail), V3237)


__e.TailApply(tmp13046, tmp13051)
return


}, 1)

tmp13052 := Call(__e, PrimFunc(symhead), V3237)


tmp13053 := Call(__e, tmp13045, tmp13052)


ifres13044 = tmp13053


} else {
tmp13054 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres13044 = tmp13054


}

__e.TailApply(tmp13041, ifres13044)
return


}, 1)

tmp13056 := Call(__e, ns2_1set, symshen_4_5expr_6, tmp13040)


_ = tmp13056

tmp13057 := MakeNative(func(__e *ControlFlow) {
V3241 := __e.Get(1)
_ = V3241
tmp13064 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(";"))
}
__typedArg0 := MakeString(";")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp13065 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(","))
}
__typedArg0 := MakeString(",")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp13066 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp13067 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1_1, Nil)
}
__typedArg0 := sym_5_1_1
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13068 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13066, tmp13067)
}
__typedArg0 := tmp13066
__typedArg1 := tmp13067
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13069 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13065, tmp13068)
}
__typedArg0 := tmp13065
__typedArg1 := tmp13068
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13070 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13064, tmp13069)
}
__typedArg0 := tmp13064
__typedArg1 := tmp13069
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13071 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_6_6, tmp13070)
}
__typedArg0 := sym_6_6
__typedArg1 := tmp13070
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13072 := Call(__e, PrimFunc(symelement_2), V3241, tmp13071)


if True == tmp13072 {
__e.Return(True)
return
} else {
tmp13062 := Call(__e, PrimFunc(symshen_4sng_2), V3241)


var ifres13059 Obj

if True == tmp13062 {
ifres13059 = True


} else {
tmp13061 := Call(__e, PrimFunc(symshen_4dbl_2), V3241)


var ifres13060 Obj

if True == tmp13061 {
ifres13060 = True


} else {
ifres13060 = False


}

ifres13059 = ifres13060


}

if True == ifres13059 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp13073 := Call(__e, ns2_1set, symshen_4key_1in_1sequent_1calculus_2, tmp13057)


_ = tmp13073

tmp13074 := MakeNative(func(__e *ControlFlow) {
V3242 := __e.Get(1)
_ = V3242
tmp13075 := MakeNative(func(__e *ControlFlow) {
W3243 := __e.Get(1)
_ = W3243
tmp13077 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3243)


if True == tmp13077 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3243)
return
}


}, 1)

tmp13078 := MakeNative(func(__e *ControlFlow) {
W3244 := __e.Get(1)
_ = W3244
tmp13084 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3244)


if True == tmp13084 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp13079 := MakeNative(func(__e *ControlFlow) {
W3245 := __e.Get(1)
_ = W3245
tmp13080 := MakeNative(func(__e *ControlFlow) {
W3246 := __e.Get(1)
_ = W3246
__e.TailApply(PrimFunc(symshen_4comb), W3246, W3245)
return
}, 1)

tmp13081 := Call(__e, PrimFunc(symshen_4in_1_6), W3244)


__e.TailApply(tmp13080, tmp13081)
return


}, 1)

tmp13082 := Call(__e, PrimFunc(symshen_4_5_1out), W3244)


__e.TailApply(tmp13079, tmp13082)
return


}


}, 1)

tmp13085 := Call(__e, PrimFunc(symshen_4_5expr_6), V3242)


tmp13086 := Call(__e, tmp13078, tmp13085)


__e.TailApply(tmp13075, tmp13086)
return


}, 1)

tmp13087 := Call(__e, ns2_1set, symshen_4_5type_6, tmp13074)


_ = tmp13087

tmp13088 := MakeNative(func(__e *ControlFlow) {
V3247 := __e.Get(1)
_ = V3247
tmp13089 := MakeNative(func(__e *ControlFlow) {
W3248 := __e.Get(1)
_ = W3248
tmp13091 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3248)


if True == tmp13091 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3248)
return
}


}, 1)

tmp13101 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3247)
}
__typedArg0 := V3247
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13092 Obj

if True == tmp13101 {
tmp13093 := MakeNative(func(__e *ControlFlow) {
W3249 := __e.Get(1)
_ = W3249
tmp13094 := MakeNative(func(__e *ControlFlow) {
W3250 := __e.Get(1)
_ = W3250
tmp13096 := Call(__e, PrimFunc(symshen_4dbl_2), W3249)


if True == tmp13096 {
__e.TailApply(PrimFunc(symshen_4comb), W3250, W3249)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp13097 := Call(__e, PrimFunc(symtail), V3247)


__e.TailApply(tmp13094, tmp13097)
return


}, 1)

tmp13098 := Call(__e, PrimFunc(symhead), V3247)


tmp13099 := Call(__e, tmp13093, tmp13098)


ifres13092 = tmp13099


} else {
tmp13100 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres13092 = tmp13100


}

__e.TailApply(tmp13089, ifres13092)
return


}, 1)

tmp13102 := Call(__e, ns2_1set, symshen_4_5dbl_6, tmp13088)


_ = tmp13102

tmp13103 := MakeNative(func(__e *ControlFlow) {
V3251 := __e.Get(1)
_ = V3251
tmp13104 := MakeNative(func(__e *ControlFlow) {
W3252 := __e.Get(1)
_ = W3252
tmp13106 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3252)


if True == tmp13106 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3252)
return
}


}, 1)

tmp13116 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3251)
}
__typedArg0 := V3251
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13107 Obj

if True == tmp13116 {
tmp13108 := MakeNative(func(__e *ControlFlow) {
W3253 := __e.Get(1)
_ = W3253
tmp13109 := MakeNative(func(__e *ControlFlow) {
W3254 := __e.Get(1)
_ = W3254
tmp13111 := Call(__e, PrimFunc(symshen_4sng_2), W3253)


if True == tmp13111 {
__e.TailApply(PrimFunc(symshen_4comb), W3254, W3253)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp13112 := Call(__e, PrimFunc(symtail), V3251)


__e.TailApply(tmp13109, tmp13112)
return


}, 1)

tmp13113 := Call(__e, PrimFunc(symhead), V3251)


tmp13114 := Call(__e, tmp13108, tmp13113)


ifres13107 = tmp13114


} else {
tmp13115 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres13107 = tmp13115


}

__e.TailApply(tmp13104, ifres13107)
return


}, 1)

tmp13117 := Call(__e, ns2_1set, symshen_4_5sng_6, tmp13103)


_ = tmp13117

tmp13118 := MakeNative(func(__e *ControlFlow) {
V3255 := __e.Get(1)
_ = V3255
tmp13123 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsymbol_2) {
return PrimIsSymbol(V3255)
}
__typedArg0 := V3255
return Call(__e, PrimFunc(symsymbol_2), __typedArg0)
})()

if True == tmp13123 {
tmp13120 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(V3255)
}
__typedArg0 := V3255
return Call(__e, PrimFunc(symstr), __typedArg0)
})()

tmp13121 := Call(__e, PrimFunc(symshen_4sng_1h_2), tmp13120)


if True == tmp13121 {
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

tmp13124 := Call(__e, ns2_1set, symshen_4sng_2, tmp13118)


_ = tmp13124

tmp13125 := MakeNative(func(__e *ControlFlow) {
V3258 := __e.Get(1)
_ = V3258
tmp13134 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("___"), V3258)
}
__typedArg0 := MakeString("___")
__typedArg1 := V3258
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13134 {
__e.Return(True)
return
} else {
tmp13132 := Call(__e, PrimFunc(symshen_4_7string_2), V3258)


var ifres13128 Obj

if True == tmp13132 {
tmp13130 := Call(__e, PrimFunc(symhdstr), V3258)


tmp13131 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("_"), tmp13130)
}
__typedArg0 := MakeString("_")
__typedArg1 := tmp13130
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13129 Obj

if True == tmp13131 {
ifres13129 = True


} else {
ifres13129 = False


}

ifres13128 = ifres13129


} else {
ifres13128 = False


}

if True == ifres13128 {
__e.TailApply(PrimFunc(symshen_4sng_1h_2), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V3258)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V3258
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp13135 := Call(__e, ns2_1set, symshen_4sng_1h_2, tmp13125)


_ = tmp13135

tmp13136 := MakeNative(func(__e *ControlFlow) {
V3259 := __e.Get(1)
_ = V3259
tmp13141 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsymbol_2) {
return PrimIsSymbol(V3259)
}
__typedArg0 := V3259
return Call(__e, PrimFunc(symsymbol_2), __typedArg0)
})()

if True == tmp13141 {
tmp13138 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(V3259)
}
__typedArg0 := V3259
return Call(__e, PrimFunc(symstr), __typedArg0)
})()

tmp13139 := Call(__e, PrimFunc(symshen_4dbl_1h_2), tmp13138)


if True == tmp13139 {
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

tmp13142 := Call(__e, ns2_1set, symshen_4dbl_2, tmp13136)


_ = tmp13142

tmp13143 := MakeNative(func(__e *ControlFlow) {
V3262 := __e.Get(1)
_ = V3262
tmp13152 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("==="), V3262)
}
__typedArg0 := MakeString("===")
__typedArg1 := V3262
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13152 {
__e.Return(True)
return
} else {
tmp13150 := Call(__e, PrimFunc(symshen_4_7string_2), V3262)


var ifres13146 Obj

if True == tmp13150 {
tmp13148 := Call(__e, PrimFunc(symhdstr), V3262)


tmp13149 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("="), tmp13148)
}
__typedArg0 := MakeString("=")
__typedArg1 := tmp13148
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13147 Obj

if True == tmp13149 {
ifres13147 = True


} else {
ifres13147 = False


}

ifres13146 = ifres13147


} else {
ifres13146 = False


}

if True == ifres13146 {
__e.TailApply(PrimFunc(symshen_4dbl_1h_2), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V3262)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V3262
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp13153 := Call(__e, ns2_1set, symshen_4dbl_1h_2, tmp13143)


_ = tmp13153

tmp13154 := MakeNative(func(__e *ControlFlow) {
V3263 := __e.Get(1)
_ = V3263
V3264 := __e.Get(2)
_ = V3264
tmp13155 := MakeNative(func(__e *ControlFlow) {
W3265 := __e.Get(1)
_ = W3265
tmp13156 := MakeNative(func(__e *ControlFlow) {
W3267 := __e.Get(1)
_ = W3267
__e.TailApply(PrimFunc(symeval), W3267)
return
}, 1)

tmp13157 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3263, W3265)
}
__typedArg0 := V3263
__typedArg1 := W3265
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13158 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefprolog, tmp13157)
}
__typedArg0 := symdefprolog
__typedArg1 := tmp13157
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(tmp13156, tmp13158)
return


}, 1)

tmp13159 := MakeNative(func(__e *ControlFlow) {
Z3266 := __e.Get(1)
_ = Z3266
__e.TailApply(PrimFunc(symshen_4rule_1_6clause), Z3266)
return
}, 1)

tmp13160 := Call(__e, PrimFunc(symmapcan), tmp13159, V3264)


__e.TailApply(tmp13155, tmp13160)
return


}, 2)

tmp13161 := Call(__e, ns2_1set, symshen_4rules_1_6prolog, tmp13154)


_ = tmp13161

tmp13162 := MakeNative(func(__e *ControlFlow) {
V3268 := __e.Get(1)
_ = V3268
tmp13223 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3268)
}
__typedArg0 := V3268
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13187 Obj

if True == tmp13223 {
tmp13221 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3268)
}
__typedArg0 := V3268
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13222 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13221)
}
__typedArg0 := tmp13221
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13189 Obj

if True == tmp13222 {
tmp13218 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3268)
}
__typedArg0 := V3268
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13219 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13218)
}
__typedArg0 := tmp13218
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13220 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13219)
}
__typedArg0 := tmp13219
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13191 Obj

if True == tmp13220 {
tmp13214 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3268)
}
__typedArg0 := V3268
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13215 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13214)
}
__typedArg0 := tmp13214
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13216 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13215)
}
__typedArg0 := tmp13215
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13217 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13216)
}
__typedArg0 := tmp13216
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13193 Obj

if True == tmp13217 {
tmp13209 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3268)
}
__typedArg0 := V3268
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13210 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13209)
}
__typedArg0 := tmp13209
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13211 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13210)
}
__typedArg0 := tmp13210
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13212 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13211)
}
__typedArg0 := tmp13211
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13213 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13212)
}
__typedArg0 := tmp13212
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13195 Obj

if True == tmp13213 {
tmp13203 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3268)
}
__typedArg0 := V3268
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13204 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13203)
}
__typedArg0 := tmp13203
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13205 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13204)
}
__typedArg0 := tmp13204
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13206 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13205)
}
__typedArg0 := tmp13205
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13207 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13206)
}
__typedArg0 := tmp13206
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13208 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp13207)
}
__typedArg0 := Nil
__typedArg1 := tmp13207
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13197 Obj

if True == tmp13208 {
tmp13199 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3268)
}
__typedArg0 := V3268
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13200 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13199)
}
__typedArg0 := tmp13199
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13201 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13200)
}
__typedArg0 := tmp13200
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13202 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp13201)
}
__typedArg0 := Nil
__typedArg1 := tmp13201
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13198 Obj

if True == tmp13202 {
ifres13198 = True


} else {
ifres13198 = False


}

ifres13197 = ifres13198


} else {
ifres13197 = False


}

var ifres13196 Obj

if True == ifres13197 {
ifres13196 = True


} else {
ifres13196 = False


}

ifres13195 = ifres13196


} else {
ifres13195 = False


}

var ifres13194 Obj

if True == ifres13195 {
ifres13194 = True


} else {
ifres13194 = False


}

ifres13193 = ifres13194


} else {
ifres13193 = False


}

var ifres13192 Obj

if True == ifres13193 {
ifres13192 = True


} else {
ifres13192 = False


}

ifres13191 = ifres13192


} else {
ifres13191 = False


}

var ifres13190 Obj

if True == ifres13191 {
ifres13190 = True


} else {
ifres13190 = False


}

ifres13189 = ifres13190


} else {
ifres13189 = False


}

var ifres13188 Obj

if True == ifres13189 {
ifres13188 = True


} else {
ifres13188 = False


}

ifres13187 = ifres13188


} else {
ifres13187 = False


}

if True == ifres13187 {
tmp13163 := MakeNative(func(__e *ControlFlow) {
W3269 := __e.Get(1)
_ = W3269
tmp13164 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3268)
}
__typedArg0 := V3268
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13165 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13164)
}
__typedArg0 := tmp13164
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13166 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13165)
}
__typedArg0 := tmp13165
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13167 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13166)
}
__typedArg0 := tmp13166
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13168 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13167)
}
__typedArg0 := tmp13167
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13169 := Call(__e, PrimFunc(symshen_4rule_1_6head), tmp13168)


tmp13170 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1_1, Nil)
}
__typedArg0 := sym_5_1_1
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13171 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3268)
}
__typedArg0 := V3268
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13172 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3268)
}
__typedArg0 := V3268
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13173 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13172)
}
__typedArg0 := tmp13172
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13174 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3268)
}
__typedArg0 := V3268
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13175 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13174)
}
__typedArg0 := tmp13174
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13176 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13175)
}
__typedArg0 := tmp13175
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13177 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13176)
}
__typedArg0 := tmp13176
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13178 := Call(__e, PrimFunc(symshen_4rule_1_6body), W3269, symAssumptions, tmp13171, tmp13173, tmp13177)


tmp13179 := Call(__e, PrimFunc(symappend), tmp13170, tmp13178)


__e.TailApply(PrimFunc(symappend), tmp13169, tmp13179)
return


}, 1)

tmp13180 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3268)
}
__typedArg0 := V3268
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13181 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13180)
}
__typedArg0 := tmp13180
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13182 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13181)
}
__typedArg0 := tmp13181
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13183 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13182)
}
__typedArg0 := tmp13182
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13184 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13183)
}
__typedArg0 := tmp13183
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13185 := Call(__e, PrimFunc(symshen_4extract_1vars), tmp13184)


__e.TailApply(tmp13163, tmp13185)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("partial function shen.rule->clause"))
}
__typedArg0 := MakeString("partial function shen.rule->clause")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}, 1)

tmp13224 := Call(__e, ns2_1set, symshen_4rule_1_6clause, tmp13162)


_ = tmp13224

tmp13225 := MakeNative(func(__e *ControlFlow) {
V3270 := __e.Get(1)
_ = V3270
tmp13226 := Call(__e, PrimFunc(symshen_4macro_1_8ch), V3270)


tmp13227 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symAssumptions, Nil)
}
__typedArg0 := symAssumptions
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13226, tmp13227)
}
__typedArg0 := tmp13226
__typedArg1 := tmp13227
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp13228 := Call(__e, ns2_1set, symshen_4rule_1_6head, tmp13225)


_ = tmp13228

tmp13229 := MakeNative(func(__e *ControlFlow) {
V3271 := __e.Get(1)
_ = V3271
tmp13230 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3271, Nil)
}
__typedArg0 := V3271
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4_8ch, tmp13230)
}
__typedArg0 := symshen_4_8ch
__typedArg1 := tmp13230
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp13231 := Call(__e, ns2_1set, symshen_4macro_1_8ch, tmp13229)


_ = tmp13231

tmp13232 := MakeNative(func(__e *ControlFlow) {
V3272 := __e.Get(1)
_ = V3272
tmp13233 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3272, Nil)
}
__typedArg0 := V3272
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4_8c, tmp13233)
}
__typedArg0 := symshen_4_8c
__typedArg1 := tmp13233
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp13234 := Call(__e, ns2_1set, symshen_4macro_1_8c, tmp13232)


_ = tmp13234

tmp13235 := MakeNative(func(__e *ControlFlow) {
V3273 := __e.Get(1)
_ = V3273
V3274 := __e.Get(2)
_ = V3274
V3275 := __e.Get(3)
_ = V3275
V3276 := __e.Get(4)
_ = V3276
V3277 := __e.Get(5)
_ = V3277
tmp13270 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3277)
}
__typedArg0 := Nil
__typedArg1 := V3277
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13270 {
__e.TailApply(PrimFunc(symshen_4side_1conditions_1_6goals), Nil, V3273, V3274, V3275, V3276)
return
} else {
tmp13268 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3276)
}
__typedArg0 := Nil
__typedArg1 := V3276
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13261 Obj

if True == tmp13268 {
tmp13267 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3277)
}
__typedArg0 := V3277
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13263 Obj

if True == tmp13267 {
tmp13265 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3277)
}
__typedArg0 := V3277
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13266 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp13265)
}
__typedArg0 := Nil
__typedArg1 := tmp13265
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13264 Obj

if True == tmp13266 {
ifres13264 = True


} else {
ifres13264 = False


}

ifres13263 = ifres13264


} else {
ifres13263 = False


}

var ifres13262 Obj

if True == ifres13263 {
ifres13262 = True


} else {
ifres13262 = False


}

ifres13261 = ifres13262


} else {
ifres13261 = False


}

if True == ifres13261 {
tmp13236 := MakeNative(func(__e *ControlFlow) {
W3278 := __e.Get(1)
_ = W3278
tmp13237 := MakeNative(func(__e *ControlFlow) {
W3279 := __e.Get(1)
_ = W3279
tmp13238 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3277)
}
__typedArg0 := V3277
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13239 := Call(__e, PrimFunc(symshen_4specialise_1member), tmp13238, V3274, W3279, W3278)


tmp13240 := Call(__e, PrimFunc(symshen_4side_1conditions_1_6goals), Nil, V3273, V3274, V3275, Nil)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13239, tmp13240)
}
__typedArg0 := tmp13239
__typedArg1 := tmp13240
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp13241 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3277)
}
__typedArg0 := V3277
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13242 := Call(__e, PrimFunc(symshen_4remove_1bystanders), V3273, tmp13241)


__e.TailApply(tmp13237, tmp13242)
return


}, 1)

tmp13243 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3277)
}
__typedArg0 := V3277
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13244 := Call(__e, PrimFunc(symshen_4passive_1variables), tmp13243, V3273)


__e.TailApply(tmp13236, tmp13244)
return


} else {
tmp13259 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3277)
}
__typedArg0 := V3277
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp13259 {
tmp13245 := MakeNative(func(__e *ControlFlow) {
W3280 := __e.Get(1)
_ = W3280
tmp13246 := MakeNative(func(__e *ControlFlow) {
W3281 := __e.Get(1)
_ = W3281
tmp13247 := MakeNative(func(__e *ControlFlow) {
W3282 := __e.Get(1)
_ = W3282
tmp13248 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3277)
}
__typedArg0 := V3277
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13249 := Call(__e, PrimFunc(symshen_4specialise_1consume), tmp13248, V3274, W3282, W3281, W3280)


tmp13250 := Call(__e, PrimFunc(symappend), V3273, W3281)


tmp13251 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3277)
}
__typedArg0 := V3277
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13252 := Call(__e, PrimFunc(symshen_4rule_1_6body), tmp13250, W3280, V3275, V3276, tmp13251)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13249, tmp13252)
}
__typedArg0 := tmp13249
__typedArg1 := tmp13252
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp13253 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3277)
}
__typedArg0 := V3277
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13254 := Call(__e, PrimFunc(symshen_4remove_1bystanders), V3273, tmp13253)


__e.TailApply(tmp13247, tmp13254)
return


}, 1)

tmp13255 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3277)
}
__typedArg0 := V3277
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13256 := Call(__e, PrimFunc(symshen_4passive_1variables), tmp13255, V3273)


__e.TailApply(tmp13246, tmp13256)
return


}, 1)

tmp13257 := Call(__e, PrimFunc(symgensym), symNewAssumptions)


__e.TailApply(tmp13245, tmp13257)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("partial function shen.rule->body"))
}
__typedArg0 := MakeString("partial function shen.rule->body")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}


}, 5)

tmp13271 := Call(__e, ns2_1set, symshen_4rule_1_6body, tmp13235)


_ = tmp13271

tmp13272 := MakeNative(func(__e *ControlFlow) {
V3283 := __e.Get(1)
_ = V3283
V3284 := __e.Get(2)
_ = V3284
V3285 := __e.Get(3)
_ = V3285
V3286 := __e.Get(4)
_ = V3286
tmp13273 := MakeNative(func(__e *ControlFlow) {
W3287 := __e.Get(1)
_ = W3287
tmp13274 := MakeNative(func(__e *ControlFlow) {
W3288 := __e.Get(1)
_ = W3288
tmp13275 := Call(__e, PrimFunc(symappend), V3285, V3286)


tmp13276 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3284, tmp13275)
}
__typedArg0 := V3284
__typedArg1 := tmp13275
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W3287, tmp13276)
}
__typedArg0 := W3287
__typedArg1 := tmp13276
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp13277 := Call(__e, PrimFunc(symshen_4member_1clause), W3287, V3283, V3285, V3286)


__e.TailApply(tmp13274, tmp13277)
return


}, 1)

tmp13278 := Call(__e, PrimFunc(symgensym), symshen_4member)


__e.TailApply(tmp13273, tmp13278)
return


}, 4)

tmp13279 := Call(__e, ns2_1set, symshen_4specialise_1member, tmp13272)


_ = tmp13279

tmp13280 := MakeNative(func(__e *ControlFlow) {
V3291 := __e.Get(1)
_ = V3291
V3292 := __e.Get(2)
_ = V3292
tmp13294 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3291)
}
__typedArg0 := Nil
__typedArg1 := V3291
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13294 {
__e.Return(Nil)
return
} else {
tmp13292 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3291)
}
__typedArg0 := V3291
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13288 Obj

if True == tmp13292 {
tmp13290 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3291)
}
__typedArg0 := V3291
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13291 := Call(__e, PrimFunc(symshen_4occurs_1check_2), tmp13290, V3292)


var ifres13289 Obj

if True == tmp13291 {
ifres13289 = True


} else {
ifres13289 = False


}

ifres13288 = ifres13289


} else {
ifres13288 = False


}

if True == ifres13288 {
tmp13281 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3291)
}
__typedArg0 := V3291
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13282 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3291)
}
__typedArg0 := V3291
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13283 := Call(__e, PrimFunc(symshen_4remove_1bystanders), tmp13282, V3292)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13281, tmp13283)
}
__typedArg0 := tmp13281
__typedArg1 := tmp13283
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp13286 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3291)
}
__typedArg0 := V3291
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp13286 {
tmp13284 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3291)
}
__typedArg0 := V3291
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4remove_1bystanders), tmp13284, V3292)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("partial function shen.remove-bystanders"))
}
__typedArg0 := MakeString("partial function shen.remove-bystanders")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}


}, 2)

tmp13295 := Call(__e, ns2_1set, symshen_4remove_1bystanders, tmp13280)


_ = tmp13295

tmp13296 := MakeNative(func(__e *ControlFlow) {
V3293 := __e.Get(1)
_ = V3293
V3294 := __e.Get(2)
_ = V3294
V3295 := __e.Get(3)
_ = V3295
V3296 := __e.Get(4)
_ = V3296
tmp13297 := MakeNative(func(__e *ControlFlow) {
W3297 := __e.Get(1)
_ = W3297
tmp13298 := MakeNative(func(__e *ControlFlow) {
W3298 := __e.Get(1)
_ = W3298
tmp13299 := MakeNative(func(__e *ControlFlow) {
W3299 := __e.Get(1)
_ = W3299
tmp13300 := MakeNative(func(__e *ControlFlow) {
W3304 := __e.Get(1)
_ = W3304
__e.TailApply(PrimFunc(symeval), W3304)
return
}, 1)

tmp13301 := Call(__e, PrimFunc(symappend), W3298, W3299)


tmp13302 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3293, tmp13301)
}
__typedArg0 := V3293
__typedArg1 := tmp13301
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13303 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefprolog, tmp13302)
}
__typedArg0 := symdefprolog
__typedArg1 := tmp13302
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(tmp13300, tmp13303)
return


}, 1)

tmp13304 := MakeNative(func(__e *ControlFlow) {
W3300 := __e.Get(1)
_ = W3300
tmp13305 := MakeNative(func(__e *ControlFlow) {
W3301 := __e.Get(1)
_ = W3301
tmp13306 := MakeNative(func(__e *ControlFlow) {
W3302 := __e.Get(1)
_ = W3302
tmp13307 := MakeNative(func(__e *ControlFlow) {
W3303 := __e.Get(1)
_ = W3303
tmp13308 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1_1, Nil)
}
__typedArg0 := sym_5_1_1
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13309 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(";"))
}
__typedArg0 := MakeString(";")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp13310 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13309, Nil)
}
__typedArg0 := tmp13309
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13311 := Call(__e, PrimFunc(symappend), W3303, tmp13310)


tmp13312 := Call(__e, PrimFunc(symappend), tmp13308, tmp13311)


__e.TailApply(PrimFunc(symappend), W3302, tmp13312)
return


}, 1)

tmp13313 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W3300, W3301)
}
__typedArg0 := W3300
__typedArg1 := W3301
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13314 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3293, tmp13313)
}
__typedArg0 := V3293
__typedArg1 := tmp13313
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13315 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13314, Nil)
}
__typedArg0 := tmp13314
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(tmp13307, tmp13315)
return


}, 1)

tmp13316 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W3300, Nil)
}
__typedArg0 := W3300
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13317 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym__, tmp13316)
}
__typedArg0 := sym__
__typedArg1 := tmp13316
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13318 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp13317)
}
__typedArg0 := symcons
__typedArg1 := tmp13317
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13319 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13318, Nil)
}
__typedArg0 := tmp13318
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13320 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1, tmp13319)
}
__typedArg0 := sym_1
__typedArg1 := tmp13319
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13321 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13320, Nil)
}
__typedArg0 := tmp13320
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13322 := Call(__e, PrimFunc(symappend), tmp13321, W3301)


__e.TailApply(tmp13306, tmp13322)
return


}, 1)

tmp13323 := Call(__e, PrimFunc(symappend), V3295, V3296)


__e.TailApply(tmp13305, tmp13323)
return


}, 1)

tmp13324 := Call(__e, PrimFunc(symgensym), symHypotheses)


tmp13325 := Call(__e, tmp13304, tmp13324)


__e.TailApply(tmp13299, tmp13325)
return


}, 1)

tmp13326 := Call(__e, PrimFunc(symshen_4macro_1_8ch), V3294)


tmp13327 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym__, Nil)
}
__typedArg0 := sym__
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13328 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13326, tmp13327)
}
__typedArg0 := tmp13326
__typedArg1 := tmp13327
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13329 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp13328)
}
__typedArg0 := symcons
__typedArg1 := tmp13328
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13330 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13329, Nil)
}
__typedArg0 := tmp13329
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13331 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1, tmp13330)
}
__typedArg0 := sym_1
__typedArg1 := tmp13330
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13332 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13331, Nil)
}
__typedArg0 := tmp13331
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13333 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1_1, Nil)
}
__typedArg0 := sym_5_1_1
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13334 := Call(__e, PrimFunc(symshen_4passive_1bind), V3296, W3297)


tmp13335 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(";"))
}
__typedArg0 := MakeString(";")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp13336 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13335, Nil)
}
__typedArg0 := tmp13335
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13337 := Call(__e, PrimFunc(symappend), tmp13334, tmp13336)


tmp13338 := Call(__e, PrimFunc(symappend), tmp13333, tmp13337)


tmp13339 := Call(__e, PrimFunc(symappend), W3297, tmp13338)


tmp13340 := Call(__e, PrimFunc(symappend), V3295, tmp13339)


tmp13341 := Call(__e, PrimFunc(symappend), tmp13332, tmp13340)


__e.TailApply(tmp13298, tmp13341)
return


}, 1)

tmp13342 := Call(__e, PrimFunc(symlength), V3296)


tmp13343 := Call(__e, PrimFunc(symshen_4nvars), tmp13342)


__e.TailApply(tmp13297, tmp13343)
return


}, 4)

tmp13344 := Call(__e, ns2_1set, symshen_4member_1clause, tmp13296)


_ = tmp13344

tmp13345 := MakeNative(func(__e *ControlFlow) {
V3305 := __e.Get(1)
_ = V3305
tmp13350 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeNumber(0), V3305)
}
__typedArg0 := MakeNumber(0)
__typedArg1 := V3305
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13350 {
__e.Return(Nil)
return
} else {
tmp13346 := Call(__e, PrimFunc(symgensym), symNewV)


tmp13348 := Call(__e, PrimFunc(symshen_4nvars), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(V3305)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := V3305
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})())


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13346, tmp13348)
}
__typedArg0 := tmp13346
__typedArg1 := tmp13348
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}


}, 1)

tmp13351 := Call(__e, ns2_1set, symshen_4nvars, tmp13345)


_ = tmp13351

tmp13352 := MakeNative(func(__e *ControlFlow) {
V3306 := __e.Get(1)
_ = V3306
V3307 := __e.Get(2)
_ = V3307
tmp13370 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3306)
}
__typedArg0 := Nil
__typedArg1 := V3306
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13367 Obj

if True == tmp13370 {
tmp13369 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3307)
}
__typedArg0 := Nil
__typedArg1 := V3307
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13368 Obj

if True == tmp13369 {
ifres13368 = True


} else {
ifres13368 = False


}

ifres13367 = ifres13368


} else {
ifres13367 = False


}

if True == ifres13367 {
__e.Return(Nil)
return
} else {
tmp13365 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3306)
}
__typedArg0 := V3306
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13362 Obj

if True == tmp13365 {
tmp13364 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3307)
}
__typedArg0 := V3307
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13363 Obj

if True == tmp13364 {
ifres13363 = True


} else {
ifres13363 = False


}

ifres13362 = ifres13363


} else {
ifres13362 = False


}

if True == ifres13362 {
tmp13353 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3307)
}
__typedArg0 := V3307
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13354 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3306)
}
__typedArg0 := V3306
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13355 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13354, Nil)
}
__typedArg0 := tmp13354
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13356 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13353, tmp13355)
}
__typedArg0 := tmp13353
__typedArg1 := tmp13355
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13357 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symbind, tmp13356)
}
__typedArg0 := symbind
__typedArg1 := tmp13356
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13358 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3306)
}
__typedArg0 := V3306
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13359 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3307)
}
__typedArg0 := V3307
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13360 := Call(__e, PrimFunc(symshen_4passive_1bind), tmp13358, tmp13359)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13357, tmp13360)
}
__typedArg0 := tmp13357
__typedArg1 := tmp13360
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("partial function shen.passive-bind"))
}
__typedArg0 := MakeString("partial function shen.passive-bind")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 2)

tmp13371 := Call(__e, ns2_1set, symshen_4passive_1bind, tmp13352)


_ = tmp13371

tmp13372 := MakeNative(func(__e *ControlFlow) {
V3308 := __e.Get(1)
_ = V3308
V3309 := __e.Get(2)
_ = V3309
V3310 := __e.Get(3)
_ = V3310
V3311 := __e.Get(4)
_ = V3311
V3312 := __e.Get(5)
_ = V3312
tmp13373 := MakeNative(func(__e *ControlFlow) {
W3313 := __e.Get(1)
_ = W3313
tmp13374 := MakeNative(func(__e *ControlFlow) {
W3314 := __e.Get(1)
_ = W3314
tmp13375 := Call(__e, PrimFunc(symappend), V3310, V3311)


tmp13376 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3312, tmp13375)
}
__typedArg0 := V3312
__typedArg1 := tmp13375
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13377 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3309, tmp13376)
}
__typedArg0 := V3309
__typedArg1 := tmp13376
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W3313, tmp13377)
}
__typedArg0 := W3313
__typedArg1 := tmp13377
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp13378 := Call(__e, PrimFunc(symshen_4consume_1clause), W3313, V3308, V3310, V3311, V3312)


__e.TailApply(tmp13374, tmp13378)
return


}, 1)

tmp13379 := Call(__e, PrimFunc(symgensym), symshen_4consume)


__e.TailApply(tmp13373, tmp13379)
return


}, 5)

tmp13380 := Call(__e, ns2_1set, symshen_4specialise_1consume, tmp13372)


_ = tmp13380

tmp13381 := MakeNative(func(__e *ControlFlow) {
V3315 := __e.Get(1)
_ = V3315
V3316 := __e.Get(2)
_ = V3316
V3317 := __e.Get(3)
_ = V3317
V3318 := __e.Get(4)
_ = V3318
V3319 := __e.Get(5)
_ = V3319
tmp13382 := MakeNative(func(__e *ControlFlow) {
W3320 := __e.Get(1)
_ = W3320
tmp13383 := MakeNative(func(__e *ControlFlow) {
W3321 := __e.Get(1)
_ = W3321
tmp13384 := MakeNative(func(__e *ControlFlow) {
W3322 := __e.Get(1)
_ = W3322
tmp13385 := MakeNative(func(__e *ControlFlow) {
W3323 := __e.Get(1)
_ = W3323
tmp13386 := MakeNative(func(__e *ControlFlow) {
W3329 := __e.Get(1)
_ = W3329
__e.TailApply(PrimFunc(symeval), W3329)
return
}, 1)

tmp13387 := Call(__e, PrimFunc(symappend), W3322, W3323)


tmp13388 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3315, tmp13387)
}
__typedArg0 := V3315
__typedArg1 := tmp13387
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13389 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefprolog, tmp13388)
}
__typedArg0 := symdefprolog
__typedArg1 := tmp13388
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(tmp13386, tmp13389)
return


}, 1)

tmp13390 := MakeNative(func(__e *ControlFlow) {
W3324 := __e.Get(1)
_ = W3324
tmp13391 := MakeNative(func(__e *ControlFlow) {
W3325 := __e.Get(1)
_ = W3325
tmp13392 := MakeNative(func(__e *ControlFlow) {
W3326 := __e.Get(1)
_ = W3326
tmp13393 := MakeNative(func(__e *ControlFlow) {
W3327 := __e.Get(1)
_ = W3327
tmp13394 := MakeNative(func(__e *ControlFlow) {
W3328 := __e.Get(1)
_ = W3328
tmp13395 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1_1, Nil)
}
__typedArg0 := sym_5_1_1
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13396 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(";"))
}
__typedArg0 := MakeString(";")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp13397 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13396, Nil)
}
__typedArg0 := tmp13396
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13398 := Call(__e, PrimFunc(symappend), W3328, tmp13397)


tmp13399 := Call(__e, PrimFunc(symappend), tmp13395, tmp13398)


__e.TailApply(PrimFunc(symappend), W3327, tmp13399)
return


}, 1)

tmp13400 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W3321, Nil)
}
__typedArg0 := W3321
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13401 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W3326, tmp13400)
}
__typedArg0 := W3326
__typedArg1 := tmp13400
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13402 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symbind, tmp13401)
}
__typedArg0 := symbind
__typedArg1 := tmp13401
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13403 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3319, W3325)
}
__typedArg0 := V3319
__typedArg1 := W3325
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13404 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W3324, tmp13403)
}
__typedArg0 := W3324
__typedArg1 := tmp13403
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13405 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3315, tmp13404)
}
__typedArg0 := V3315
__typedArg1 := tmp13404
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13406 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13405, Nil)
}
__typedArg0 := tmp13405
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13407 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13402, tmp13406)
}
__typedArg0 := tmp13402
__typedArg1 := tmp13406
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(tmp13394, tmp13407)
return


}, 1)

tmp13408 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W3324, Nil)
}
__typedArg0 := W3324
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13409 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W3321, tmp13408)
}
__typedArg0 := W3321
__typedArg1 := tmp13408
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13410 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp13409)
}
__typedArg0 := symcons
__typedArg1 := tmp13409
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13411 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13410, Nil)
}
__typedArg0 := tmp13410
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13412 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1, tmp13411)
}
__typedArg0 := sym_1
__typedArg1 := tmp13411
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13413 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3319, Nil)
}
__typedArg0 := V3319
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13414 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W3326, tmp13413)
}
__typedArg0 := W3326
__typedArg1 := tmp13413
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13415 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp13414)
}
__typedArg0 := symcons
__typedArg1 := tmp13414
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13416 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13415, W3325)
}
__typedArg0 := tmp13415
__typedArg1 := W3325
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13417 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13412, tmp13416)
}
__typedArg0 := tmp13412
__typedArg1 := tmp13416
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(tmp13393, tmp13417)
return


}, 1)

tmp13418 := Call(__e, PrimFunc(symgensym), symAssumptions)


__e.TailApply(tmp13392, tmp13418)
return


}, 1)

tmp13419 := Call(__e, PrimFunc(symappend), V3317, V3318)


__e.TailApply(tmp13391, tmp13419)
return


}, 1)

tmp13420 := Call(__e, PrimFunc(symgensym), symHypotheses)


tmp13421 := Call(__e, tmp13390, tmp13420)


__e.TailApply(tmp13385, tmp13421)
return


}, 1)

tmp13422 := Call(__e, PrimFunc(symshen_4macro_1_8ch), V3316)


tmp13423 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W3321, Nil)
}
__typedArg0 := W3321
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13424 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13422, tmp13423)
}
__typedArg0 := tmp13422
__typedArg1 := tmp13423
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13425 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp13424)
}
__typedArg0 := symcons
__typedArg1 := tmp13424
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13426 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13425, Nil)
}
__typedArg0 := tmp13425
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13427 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1, tmp13426)
}
__typedArg0 := sym_1
__typedArg1 := tmp13426
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13428 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1_1, Nil)
}
__typedArg0 := sym_5_1_1
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13429 := Call(__e, PrimFunc(symshen_4passive_1bind), V3318, W3320)


tmp13430 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W3321, Nil)
}
__typedArg0 := W3321
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13431 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3319, tmp13430)
}
__typedArg0 := V3319
__typedArg1 := tmp13430
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13432 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symbind, tmp13431)
}
__typedArg0 := symbind
__typedArg1 := tmp13431
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13433 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(";"))
}
__typedArg0 := MakeString(";")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp13434 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13433, Nil)
}
__typedArg0 := tmp13433
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13435 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13432, tmp13434)
}
__typedArg0 := tmp13432
__typedArg1 := tmp13434
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13436 := Call(__e, PrimFunc(symappend), tmp13429, tmp13435)


tmp13437 := Call(__e, PrimFunc(symappend), tmp13428, tmp13436)


tmp13438 := Call(__e, PrimFunc(symappend), W3320, tmp13437)


tmp13439 := Call(__e, PrimFunc(symappend), V3317, tmp13438)


tmp13440 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3319, tmp13439)
}
__typedArg0 := V3319
__typedArg1 := tmp13439
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13441 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13427, tmp13440)
}
__typedArg0 := tmp13427
__typedArg1 := tmp13440
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(tmp13384, tmp13441)
return


}, 1)

tmp13442 := Call(__e, PrimFunc(symgensym), symAssumption)


__e.TailApply(tmp13383, tmp13442)
return


}, 1)

tmp13443 := Call(__e, PrimFunc(symlength), V3318)


tmp13444 := Call(__e, PrimFunc(symshen_4nvars), tmp13443)


__e.TailApply(tmp13382, tmp13444)
return


}, 5)

tmp13445 := Call(__e, ns2_1set, symshen_4consume_1clause, tmp13381)


_ = tmp13445

tmp13446 := MakeNative(func(__e *ControlFlow) {
V3330 := __e.Get(1)
_ = V3330
V3331 := __e.Get(2)
_ = V3331
tmp13447 := Call(__e, PrimFunc(symshen_4extract_1vars), V3330)


__e.TailApply(PrimFunc(symdifference), tmp13447, V3331)
return


}, 2)

tmp13448 := Call(__e, ns2_1set, symshen_4passive_1variables, tmp13446)


_ = tmp13448

tmp13449 := MakeNative(func(__e *ControlFlow) {
V3336 := __e.Get(1)
_ = V3336
V3337 := __e.Get(2)
_ = V3337
V3338 := __e.Get(3)
_ = V3338
V3339 := __e.Get(4)
_ = V3339
V3340 := __e.Get(5)
_ = V3340
tmp13600 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3339)
}
__typedArg0 := Nil
__typedArg1 := V3339
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13600 {
__e.TailApply(PrimFunc(symshen_4premises_1_6goals), V3336, V3338, V3340)
return
} else {
tmp13598 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3339)
}
__typedArg0 := V3339
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13578 Obj

if True == tmp13598 {
tmp13596 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3339)
}
__typedArg0 := V3339
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13597 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13596)
}
__typedArg0 := tmp13596
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13580 Obj

if True == tmp13597 {
tmp13593 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3339)
}
__typedArg0 := V3339
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13594 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13593)
}
__typedArg0 := tmp13593
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13595 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symif, tmp13594)
}
__typedArg0 := symif
__typedArg1 := tmp13594
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13582 Obj

if True == tmp13595 {
tmp13590 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3339)
}
__typedArg0 := V3339
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13591 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13590)
}
__typedArg0 := tmp13590
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13592 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13591)
}
__typedArg0 := tmp13591
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13584 Obj

if True == tmp13592 {
tmp13586 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3339)
}
__typedArg0 := V3339
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13587 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13586)
}
__typedArg0 := tmp13586
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13588 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13587)
}
__typedArg0 := tmp13587
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13589 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp13588)
}
__typedArg0 := Nil
__typedArg1 := tmp13588
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13585 Obj

if True == tmp13589 {
ifres13585 = True


} else {
ifres13585 = False


}

ifres13584 = ifres13585


} else {
ifres13584 = False


}

var ifres13583 Obj

if True == ifres13584 {
ifres13583 = True


} else {
ifres13583 = False


}

ifres13582 = ifres13583


} else {
ifres13582 = False


}

var ifres13581 Obj

if True == ifres13582 {
ifres13581 = True


} else {
ifres13581 = False


}

ifres13580 = ifres13581


} else {
ifres13580 = False


}

var ifres13579 Obj

if True == ifres13580 {
ifres13579 = True


} else {
ifres13579 = False


}

ifres13578 = ifres13579


} else {
ifres13578 = False


}

if True == ifres13578 {
tmp13450 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3339)
}
__typedArg0 := V3339
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13451 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13450)
}
__typedArg0 := tmp13450
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13452 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhen, tmp13451)
}
__typedArg0 := symwhen
__typedArg1 := tmp13451
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13453 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3339)
}
__typedArg0 := V3339
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13454 := Call(__e, PrimFunc(symshen_4side_1conditions_1_6goals), V3336, V3337, V3338, tmp13453, V3340)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13452, tmp13454)
}
__typedArg0 := tmp13452
__typedArg1 := tmp13454
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp13576 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3339)
}
__typedArg0 := V3339
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13549 Obj

if True == tmp13576 {
tmp13574 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3339)
}
__typedArg0 := V3339
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13575 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13574)
}
__typedArg0 := tmp13574
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13551 Obj

if True == tmp13575 {
tmp13571 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3339)
}
__typedArg0 := V3339
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13572 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13571)
}
__typedArg0 := tmp13571
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13573 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlet, tmp13572)
}
__typedArg0 := symlet
__typedArg1 := tmp13572
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13553 Obj

if True == tmp13573 {
tmp13568 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3339)
}
__typedArg0 := V3339
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13569 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13568)
}
__typedArg0 := tmp13568
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13570 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13569)
}
__typedArg0 := tmp13569
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13555 Obj

if True == tmp13570 {
tmp13564 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3339)
}
__typedArg0 := V3339
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13565 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13564)
}
__typedArg0 := tmp13564
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13566 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13565)
}
__typedArg0 := tmp13565
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13567 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13566)
}
__typedArg0 := tmp13566
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13557 Obj

if True == tmp13567 {
tmp13559 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3339)
}
__typedArg0 := V3339
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13560 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13559)
}
__typedArg0 := tmp13559
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13561 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13560)
}
__typedArg0 := tmp13560
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13562 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13561)
}
__typedArg0 := tmp13561
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13563 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp13562)
}
__typedArg0 := Nil
__typedArg1 := tmp13562
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13558 Obj

if True == tmp13563 {
ifres13558 = True


} else {
ifres13558 = False


}

ifres13557 = ifres13558


} else {
ifres13557 = False


}

var ifres13556 Obj

if True == ifres13557 {
ifres13556 = True


} else {
ifres13556 = False


}

ifres13555 = ifres13556


} else {
ifres13555 = False


}

var ifres13554 Obj

if True == ifres13555 {
ifres13554 = True


} else {
ifres13554 = False


}

ifres13553 = ifres13554


} else {
ifres13553 = False


}

var ifres13552 Obj

if True == ifres13553 {
ifres13552 = True


} else {
ifres13552 = False


}

ifres13551 = ifres13552


} else {
ifres13551 = False


}

var ifres13550 Obj

if True == ifres13551 {
ifres13550 = True


} else {
ifres13550 = False


}

ifres13549 = ifres13550


} else {
ifres13549 = False


}

if True == ifres13549 {
tmp13470 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3339)
}
__typedArg0 := V3339
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13471 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13470)
}
__typedArg0 := tmp13470
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13472 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13471)
}
__typedArg0 := tmp13471
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13473 := Call(__e, PrimFunc(symelement_2), tmp13472, V3337)


if True == tmp13473 {
tmp13455 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3339)
}
__typedArg0 := V3339
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13456 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13455)
}
__typedArg0 := tmp13455
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13457 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symis_b, tmp13456)
}
__typedArg0 := symis_b
__typedArg1 := tmp13456
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13458 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3339)
}
__typedArg0 := V3339
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13459 := Call(__e, PrimFunc(symshen_4side_1conditions_1_6goals), V3336, V3337, V3338, tmp13458, V3340)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13457, tmp13459)
}
__typedArg0 := tmp13457
__typedArg1 := tmp13459
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp13460 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3339)
}
__typedArg0 := V3339
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13461 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13460)
}
__typedArg0 := tmp13460
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13462 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symbind, tmp13461)
}
__typedArg0 := symbind
__typedArg1 := tmp13461
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13463 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3339)
}
__typedArg0 := V3339
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13464 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13463)
}
__typedArg0 := tmp13463
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13465 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13464)
}
__typedArg0 := tmp13464
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13466 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13465, V3337)
}
__typedArg0 := tmp13465
__typedArg1 := V3337
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13467 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3339)
}
__typedArg0 := V3339
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13468 := Call(__e, PrimFunc(symshen_4side_1conditions_1_6goals), V3336, tmp13466, V3338, tmp13467, V3340)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13462, tmp13468)
}
__typedArg0 := tmp13462
__typedArg1 := tmp13468
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}


} else {
tmp13547 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3339)
}
__typedArg0 := V3339
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13527 Obj

if True == tmp13547 {
tmp13545 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3339)
}
__typedArg0 := V3339
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13546 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13545)
}
__typedArg0 := tmp13545
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13529 Obj

if True == tmp13546 {
tmp13542 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3339)
}
__typedArg0 := V3339
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13543 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13542)
}
__typedArg0 := tmp13542
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13544 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symctxt, tmp13543)
}
__typedArg0 := symctxt
__typedArg1 := tmp13543
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13531 Obj

if True == tmp13544 {
tmp13539 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3339)
}
__typedArg0 := V3339
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13540 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13539)
}
__typedArg0 := tmp13539
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13541 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13540)
}
__typedArg0 := tmp13540
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13533 Obj

if True == tmp13541 {
tmp13535 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3339)
}
__typedArg0 := V3339
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13536 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13535)
}
__typedArg0 := tmp13535
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13537 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13536)
}
__typedArg0 := tmp13536
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13538 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp13537)
}
__typedArg0 := Nil
__typedArg1 := tmp13537
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13534 Obj

if True == tmp13538 {
ifres13534 = True


} else {
ifres13534 = False


}

ifres13533 = ifres13534


} else {
ifres13533 = False


}

var ifres13532 Obj

if True == ifres13533 {
ifres13532 = True


} else {
ifres13532 = False


}

ifres13531 = ifres13532


} else {
ifres13531 = False


}

var ifres13530 Obj

if True == ifres13531 {
ifres13530 = True


} else {
ifres13530 = False


}

ifres13529 = ifres13530


} else {
ifres13529 = False


}

var ifres13528 Obj

if True == ifres13529 {
ifres13528 = True


} else {
ifres13528 = False


}

ifres13527 = ifres13528


} else {
ifres13527 = False


}

if True == ifres13527 {
tmp13499 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3339)
}
__typedArg0 := V3339
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13500 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13499)
}
__typedArg0 := tmp13499
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13501 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13500)
}
__typedArg0 := tmp13500
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13502 := Call(__e, PrimFunc(symelement_2), tmp13501, V3337)


if True == tmp13502 {
tmp13474 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3339)
}
__typedArg0 := V3339
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13475 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13474)
}
__typedArg0 := tmp13474
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13476 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13475)
}
__typedArg0 := tmp13475
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13477 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13476, V3336)
}
__typedArg0 := tmp13476
__typedArg1 := V3336
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13478 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3339)
}
__typedArg0 := V3339
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4side_1conditions_1_6goals), tmp13477, V3337, V3338, tmp13478, V3340)
return


} else {
tmp13479 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3339)
}
__typedArg0 := V3339
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13480 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13479)
}
__typedArg0 := tmp13479
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13481 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13480)
}
__typedArg0 := tmp13480
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13482 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3338, Nil)
}
__typedArg0 := V3338
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13483 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13481, tmp13482)
}
__typedArg0 := tmp13481
__typedArg1 := tmp13482
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13484 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symbind, tmp13483)
}
__typedArg0 := symbind
__typedArg1 := tmp13483
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13485 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3339)
}
__typedArg0 := V3339
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13486 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13485)
}
__typedArg0 := tmp13485
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13487 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13486)
}
__typedArg0 := tmp13486
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13488 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13487, V3336)
}
__typedArg0 := tmp13487
__typedArg1 := V3336
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13489 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3339)
}
__typedArg0 := V3339
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13490 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13489)
}
__typedArg0 := tmp13489
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13491 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13490)
}
__typedArg0 := tmp13490
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13492 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13491, V3337)
}
__typedArg0 := tmp13491
__typedArg1 := V3337
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13493 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3339)
}
__typedArg0 := V3339
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13494 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13493)
}
__typedArg0 := tmp13493
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13495 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13494)
}
__typedArg0 := tmp13494
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13496 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3339)
}
__typedArg0 := V3339
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13497 := Call(__e, PrimFunc(symshen_4side_1conditions_1_6goals), tmp13488, tmp13492, tmp13495, tmp13496, V3340)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13484, tmp13497)
}
__typedArg0 := tmp13484
__typedArg1 := tmp13497
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}


} else {
tmp13525 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3339)
}
__typedArg0 := V3339
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13505 Obj

if True == tmp13525 {
tmp13523 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3339)
}
__typedArg0 := V3339
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13524 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13523)
}
__typedArg0 := tmp13523
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13507 Obj

if True == tmp13524 {
tmp13520 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3339)
}
__typedArg0 := V3339
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13521 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13520)
}
__typedArg0 := tmp13520
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13522 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symsqts, tmp13521)
}
__typedArg0 := symsqts
__typedArg1 := tmp13521
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13509 Obj

if True == tmp13522 {
tmp13517 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3339)
}
__typedArg0 := V3339
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13518 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13517)
}
__typedArg0 := tmp13517
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13519 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13518)
}
__typedArg0 := tmp13518
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13511 Obj

if True == tmp13519 {
tmp13513 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3339)
}
__typedArg0 := V3339
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13514 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13513)
}
__typedArg0 := tmp13513
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13515 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13514)
}
__typedArg0 := tmp13514
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13516 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp13515)
}
__typedArg0 := Nil
__typedArg1 := tmp13515
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13512 Obj

if True == tmp13516 {
ifres13512 = True


} else {
ifres13512 = False


}

ifres13511 = ifres13512


} else {
ifres13511 = False


}

var ifres13510 Obj

if True == ifres13511 {
ifres13510 = True


} else {
ifres13510 = False


}

ifres13509 = ifres13510


} else {
ifres13509 = False


}

var ifres13508 Obj

if True == ifres13509 {
ifres13508 = True


} else {
ifres13508 = False


}

ifres13507 = ifres13508


} else {
ifres13507 = False


}

var ifres13506 Obj

if True == ifres13507 {
ifres13506 = True


} else {
ifres13506 = False


}

ifres13505 = ifres13506


} else {
ifres13505 = False


}

if True == ifres13505 {
tmp13503 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3339)
}
__typedArg0 := V3339
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4side_1conditions_1_6goals), V3336, V3337, V3338, tmp13503, V3340)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("partial function shen.side-conditions->goals"))
}
__typedArg0 := MakeString("partial function shen.side-conditions->goals")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}


}


}


}, 5)

tmp13601 := Call(__e, ns2_1set, symshen_4side_1conditions_1_6goals, tmp13449)


_ = tmp13601

tmp13602 := MakeNative(func(__e *ControlFlow) {
V3345 := __e.Get(1)
_ = V3345
V3346 := __e.Get(2)
_ = V3346
V3347 := __e.Get(3)
_ = V3347
tmp13652 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3347)
}
__typedArg0 := Nil
__typedArg1 := V3347
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13652 {
tmp13603 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(";"))
}
__typedArg0 := MakeString(";")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13603, Nil)
}
__typedArg0 := tmp13603
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp13650 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3347)
}
__typedArg0 := V3347
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13646 Obj

if True == tmp13650 {
tmp13648 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3347)
}
__typedArg0 := V3347
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13649 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_b, tmp13648)
}
__typedArg0 := sym_b
__typedArg1 := tmp13648
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13647 Obj

if True == tmp13649 {
ifres13647 = True


} else {
ifres13647 = False


}

ifres13646 = ifres13647


} else {
ifres13646 = False


}

if True == ifres13646 {
tmp13604 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3347)
}
__typedArg0 := V3347
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13605 := Call(__e, PrimFunc(symshen_4premises_1_6goals), V3345, V3346, tmp13604)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_b, tmp13605)
}
__typedArg0 := sym_b
__typedArg1 := tmp13605
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp13644 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3347)
}
__typedArg0 := V3347
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13640 Obj

if True == tmp13644 {
tmp13642 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3347)
}
__typedArg0 := V3347
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13643 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symfail, tmp13642)
}
__typedArg0 := symfail
__typedArg1 := tmp13642
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13641 Obj

if True == tmp13643 {
ifres13641 = True


} else {
ifres13641 = False


}

ifres13640 = ifres13641


} else {
ifres13640 = False


}

if True == ifres13640 {
tmp13606 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(False, Nil)
}
__typedArg0 := False
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13607 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhen, tmp13606)
}
__typedArg0 := symwhen
__typedArg1 := tmp13606
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13608 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3347)
}
__typedArg0 := V3347
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13609 := Call(__e, PrimFunc(symshen_4premises_1_6goals), V3345, V3346, tmp13608)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13607, tmp13609)
}
__typedArg0 := tmp13607
__typedArg1 := tmp13609
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp13638 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3347)
}
__typedArg0 := V3347
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13623 Obj

if True == tmp13638 {
tmp13636 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3347)
}
__typedArg0 := V3347
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13637 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13636)
}
__typedArg0 := tmp13636
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13625 Obj

if True == tmp13637 {
tmp13633 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3347)
}
__typedArg0 := V3347
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13634 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13633)
}
__typedArg0 := tmp13633
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13635 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13634)
}
__typedArg0 := tmp13634
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13627 Obj

if True == tmp13635 {
tmp13629 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3347)
}
__typedArg0 := V3347
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13630 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13629)
}
__typedArg0 := tmp13629
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13631 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13630)
}
__typedArg0 := tmp13630
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13632 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp13631)
}
__typedArg0 := Nil
__typedArg1 := tmp13631
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13628 Obj

if True == tmp13632 {
ifres13628 = True


} else {
ifres13628 = False


}

ifres13627 = ifres13628


} else {
ifres13627 = False


}

var ifres13626 Obj

if True == ifres13627 {
ifres13626 = True


} else {
ifres13626 = False


}

ifres13625 = ifres13626


} else {
ifres13625 = False


}

var ifres13624 Obj

if True == ifres13625 {
ifres13624 = True


} else {
ifres13624 = False


}

ifres13623 = ifres13624


} else {
ifres13623 = False


}

if True == ifres13623 {
tmp13610 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3347)
}
__typedArg0 := V3347
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13611 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13610)
}
__typedArg0 := tmp13610
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13612 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13611)
}
__typedArg0 := tmp13611
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13613 := Call(__e, PrimFunc(symshen_4macro_1_8c), tmp13612)


tmp13614 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3347)
}
__typedArg0 := V3347
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13615 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13614)
}
__typedArg0 := tmp13614
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13616 := Call(__e, PrimFunc(symshen_4construct_1context), V3345, tmp13615, V3346)


tmp13617 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13616, Nil)
}
__typedArg0 := tmp13616
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13618 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13613, tmp13617)
}
__typedArg0 := tmp13613
__typedArg1 := tmp13617
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13619 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4system_1S, tmp13618)
}
__typedArg0 := symshen_4system_1S
__typedArg1 := tmp13618
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13620 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3347)
}
__typedArg0 := V3347
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13621 := Call(__e, PrimFunc(symshen_4premises_1_6goals), V3345, V3346, tmp13620)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13619, tmp13621)
}
__typedArg0 := tmp13619
__typedArg1 := tmp13621
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("partial function shen.premises->goals"))
}
__typedArg0 := MakeString("partial function shen.premises->goals")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}


}


}, 3)

tmp13653 := Call(__e, ns2_1set, symshen_4premises_1_6goals, tmp13602)


_ = tmp13653

tmp13654 := MakeNative(func(__e *ControlFlow) {
V3351 := __e.Get(1)
_ = V3351
V3352 := __e.Get(2)
_ = V3352
V3353 := __e.Get(3)
_ = V3353
tmp13674 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3352)
}
__typedArg0 := Nil
__typedArg1 := V3352
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13674 {
__e.Return(V3353)
return
} else {
tmp13672 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3352)
}
__typedArg0 := V3352
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13664 Obj

if True == tmp13672 {
tmp13670 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3352)
}
__typedArg0 := V3352
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13671 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp13670)
}
__typedArg0 := Nil
__typedArg1 := tmp13670
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13666 Obj

if True == tmp13671 {
tmp13668 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3352)
}
__typedArg0 := V3352
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13669 := Call(__e, PrimFunc(symelement_2), tmp13668, V3351)


var ifres13667 Obj

if True == tmp13669 {
ifres13667 = True


} else {
ifres13667 = False


}

ifres13666 = ifres13667


} else {
ifres13666 = False


}

var ifres13665 Obj

if True == ifres13666 {
ifres13665 = True


} else {
ifres13665 = False


}

ifres13664 = ifres13665


} else {
ifres13664 = False


}

if True == ifres13664 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3352)
}
__typedArg0 := V3352
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return
} else {
tmp13662 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3352)
}
__typedArg0 := V3352
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp13662 {
tmp13655 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3352)
}
__typedArg0 := V3352
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13656 := Call(__e, PrimFunc(symshen_4macro_1_8c), tmp13655)


tmp13657 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3352)
}
__typedArg0 := V3352
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13658 := Call(__e, PrimFunc(symshen_4construct_1context), V3351, tmp13657, V3353)


tmp13659 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13658, Nil)
}
__typedArg0 := tmp13658
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13660 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13656, tmp13659)
}
__typedArg0 := tmp13656
__typedArg1 := tmp13659
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp13660)
}
__typedArg0 := symcons
__typedArg1 := tmp13660
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("partial function shen.construct-context"))
}
__typedArg0 := MakeString("partial function shen.construct-context")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}


}, 3)

tmp13675 := Call(__e, ns2_1set, symshen_4construct_1context, tmp13654)


_ = tmp13675

tmp13676 := MakeNative(func(__e *ControlFlow) {
V3354 := __e.Get(1)
_ = V3354
tmp13677 := MakeNative(func(__e *ControlFlow) {
W3355 := __e.Get(1)
_ = W3355
tmp13678 := MakeNative(func(__e *ControlFlow) {
W3357 := __e.Get(1)
_ = W3357
tmp13679 := MakeNative(func(__e *ControlFlow) {
W3358 := __e.Get(1)
_ = W3358
tmp13680 := MakeNative(func(__e *ControlFlow) {
W3359 := __e.Get(1)
_ = W3359
__e.TailApply(PrimFunc(symshen_4show_1datatypes), W3359)
return
}, 1)

tmp13681 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_ddatatypes_d, W3358)
}
__typedArg0 := symshen_4_ddatatypes_d
__typedArg1 := W3358
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

__e.TailApply(tmp13680, tmp13681)
return


}, 1)

tmp13682 := Call(__e, PrimFunc(symshen_4remove_1datatypes), W3355, W3357)


__e.TailApply(tmp13679, tmp13682)
return


}, 1)

tmp13683 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_ddatatypes_d)
}
__typedArg0 := symshen_4_ddatatypes_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(tmp13678, tmp13683)
return


}, 1)

tmp13684 := MakeNative(func(__e *ControlFlow) {
Z3356 := __e.Get(1)
_ = Z3356
__e.TailApply(PrimFunc(symshen_4intern_1type), Z3356)
return
}, 1)

tmp13685 := Call(__e, PrimFunc(symmap), tmp13684, V3354)


__e.TailApply(tmp13677, tmp13685)
return


}, 1)

tmp13686 := Call(__e, ns2_1set, sympreclude, tmp13676)


_ = tmp13686

tmp13687 := MakeNative(func(__e *ControlFlow) {
V3364 := __e.Get(1)
_ = V3364
V3365 := __e.Get(2)
_ = V3365
tmp13694 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3364)
}
__typedArg0 := Nil
__typedArg1 := V3364
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13694 {
__e.Return(V3365)
return
} else {
tmp13692 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3364)
}
__typedArg0 := V3364
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp13692 {
tmp13688 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3364)
}
__typedArg0 := V3364
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13689 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3364)
}
__typedArg0 := V3364
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13690 := Call(__e, PrimFunc(symshen_4unassoc), tmp13689, V3365)


__e.TailApply(PrimFunc(symshen_4remove_1datatypes), tmp13688, tmp13690)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.remove-datatypes"))
}
__typedArg0 := MakeString("implementation error in shen.remove-datatypes")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 2)

tmp13695 := Call(__e, ns2_1set, symshen_4remove_1datatypes, tmp13687)


_ = tmp13695

tmp13696 := MakeNative(func(__e *ControlFlow) {
V3366 := __e.Get(1)
_ = V3366
tmp13697 := MakeNative(func(__e *ControlFlow) {
Z3367 := __e.Get(1)
_ = Z3367
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(Z3367)
}
__typedArg0 := Z3367
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp13697, V3366)
return


}, 1)

tmp13698 := Call(__e, ns2_1set, symshen_4show_1datatypes, tmp13696)


_ = tmp13698

tmp13699 := MakeNative(func(__e *ControlFlow) {
V3368 := __e.Get(1)
_ = V3368
tmp13700 := MakeNative(func(__e *ControlFlow) {
W3369 := __e.Get(1)
_ = W3369
tmp13701 := MakeNative(func(__e *ControlFlow) {
W3371 := __e.Get(1)
_ = W3371
tmp13702 := MakeNative(func(__e *ControlFlow) {
W3373 := __e.Get(1)
_ = W3373
__e.TailApply(PrimFunc(symshen_4show_1datatypes), W3373)
return
}, 1)

tmp13703 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_ddatatypes_d)
}
__typedArg0 := symshen_4_ddatatypes_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(tmp13702, tmp13703)
return


}, 1)

tmp13704 := MakeNative(func(__e *ControlFlow) {
Z3372 := __e.Get(1)
_ = Z3372
tmp13705 := Call(__e, PrimFunc(symfn), Z3372)


__e.TailApply(PrimFunc(symshen_4remember_1datatype), Z3372, tmp13705)
return


}, 1)

tmp13706 := Call(__e, PrimFunc(symmap), tmp13704, W3369)


__e.TailApply(tmp13701, tmp13706)
return


}, 1)

tmp13707 := MakeNative(func(__e *ControlFlow) {
Z3370 := __e.Get(1)
_ = Z3370
__e.TailApply(PrimFunc(symshen_4intern_1type), Z3370)
return
}, 1)

tmp13708 := Call(__e, PrimFunc(symmap), tmp13707, V3368)


__e.TailApply(tmp13700, tmp13708)
return


}, 1)

tmp13709 := Call(__e, ns2_1set, syminclude, tmp13699)


_ = tmp13709

tmp13710 := MakeNative(func(__e *ControlFlow) {
V3374 := __e.Get(1)
_ = V3374
tmp13711 := MakeNative(func(__e *ControlFlow) {
W3375 := __e.Get(1)
_ = W3375
tmp13712 := MakeNative(func(__e *ControlFlow) {
W3376 := __e.Get(1)
_ = W3376
tmp13713 := MakeNative(func(__e *ControlFlow) {
W3378 := __e.Get(1)
_ = W3378
tmp13714 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_ddatatypes_d)
}
__typedArg0 := symshen_4_ddatatypes_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4show_1datatypes), tmp13714)
return


}, 1)

tmp13715 := MakeNative(func(__e *ControlFlow) {
Z3379 := __e.Get(1)
_ = Z3379
tmp13716 := Call(__e, PrimFunc(symfn), Z3379)


__e.TailApply(PrimFunc(symshen_4remember_1datatype), Z3379, tmp13716)
return


}, 1)

tmp13717 := Call(__e, PrimFunc(symmap), tmp13715, W3376)


__e.TailApply(tmp13713, tmp13717)
return


}, 1)

tmp13718 := MakeNative(func(__e *ControlFlow) {
Z3377 := __e.Get(1)
_ = Z3377
__e.TailApply(PrimFunc(symshen_4intern_1type), Z3377)
return
}, 1)

tmp13719 := Call(__e, PrimFunc(symmap), tmp13718, V3374)


__e.TailApply(tmp13712, tmp13719)
return


}, 1)

tmp13720 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_ddatatypes_d, Nil)
}
__typedArg0 := symshen_4_ddatatypes_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

__e.TailApply(tmp13711, tmp13720)
return


}, 1)

tmp13721 := Call(__e, ns2_1set, sympreclude_1all_1but, tmp13710)


_ = tmp13721

tmp13722 := MakeNative(func(__e *ControlFlow) {
V3380 := __e.Get(1)
_ = V3380
tmp13723 := MakeNative(func(__e *ControlFlow) {
W3381 := __e.Get(1)
_ = W3381
tmp13724 := MakeNative(func(__e *ControlFlow) {
W3383 := __e.Get(1)
_ = W3383
tmp13725 := MakeNative(func(__e *ControlFlow) {
W3384 := __e.Get(1)
_ = W3384
__e.TailApply(PrimFunc(symshen_4show_1datatypes), W3384)
return
}, 1)

tmp13726 := Call(__e, PrimFunc(symshen_4remove_1datatypes), W3381, W3383)


tmp13727 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_ddatatypes_d, tmp13726)
}
__typedArg0 := symshen_4_ddatatypes_d
__typedArg1 := tmp13726
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

__e.TailApply(tmp13725, tmp13727)
return


}, 1)

tmp13728 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dalldatatypes_d)
}
__typedArg0 := symshen_4_dalldatatypes_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(tmp13724, tmp13728)
return


}, 1)

tmp13729 := MakeNative(func(__e *ControlFlow) {
Z3382 := __e.Get(1)
_ = Z3382
__e.TailApply(PrimFunc(symshen_4intern_1type), Z3382)
return
}, 1)

tmp13730 := Call(__e, PrimFunc(symmap), tmp13729, V3380)


__e.TailApply(tmp13723, tmp13730)
return


}, 1)

__e.TailApply(ns2_1set, syminclude_1all_1but, tmp13722)
return




}, 0)

