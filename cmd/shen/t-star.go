package main

import . "github.com/pyrex41/shen-go/kl"

var TStarMain = MakeNative(func(__e *ControlFlow) {
tmp14987 := MakeNative(func(__e *ControlFlow) {
V4448 := __e.Get(1)
_ = V4448
V4449 := __e.Get(2)
_ = V4449
tmp14988 := Call(__e, PrimFunc(symshen_4extract_1vars), V4449)


W445014437 := tmp14988
_ = W445014437

tmp14989 := Call(__e, PrimFunc(symshen_4rectify_1type), V4449)


W445114438 := tmp14989
_ = W445114438

tmp14990 := Call(__e, PrimFunc(symshen_4curry), V4448)


W445214439 := tmp14990
_ = W445214439

tmp14991 := MakeNative(func(__e *ControlFlow) {
Z4453 := __e.Get(1)
_ = Z4453
__e.Return(MakeNative(func(__e *ControlFlow) {
Z4454 := __e.Get(1)
_ = Z4454
__e.Return(MakeNative(func(__e *ControlFlow) {
Z4455 := __e.Get(1)
_ = Z4455
__e.Return(MakeNative(func(__e *ControlFlow) {
Z4456 := __e.Get(1)
_ = Z4456
tmp14992 := Call(__e, PrimFunc(symshen_4newpv), Z4453)


W445714440 := tmp14992
_ = W445714440

tmp14993 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14993

tmp14994 := Call(__e, PrimFunc(symshen_4deref), W445014437, Z4453)


tmp14995 := Call(__e, PrimFunc(symreceive), tmp14994)


tmp14996 := Call(__e, PrimFunc(symshen_4deref), W445114438, Z4453)


tmp14997 := Call(__e, PrimFunc(symreceive), tmp14996)


tmp14998 := MakeNative(func(__e *ControlFlow) {
tmp14999 := Call(__e, PrimFunc(symshen_4deref), W445214439, Z4453)


tmp15000 := Call(__e, PrimFunc(symreceive), tmp14999)


tmp15001 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symreturn), W445714440, Z4453, Z4454, Z4455, Z4456)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4toplevel_1forms), tmp15000, W445714440, Z4453, Z4454, Z4455, tmp15001)
return


}, 0)

tmp15002 := Call(__e, PrimFunc(symshen_4insert_1prolog_1variables), tmp14995, tmp14997, W445714440, Z4453, Z4454, Z4455, tmp14998)


__e.TailApply(PrimFunc(symshen_4gc), Z4453, tmp15002)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp15003 := Call(__e, PrimFunc(symshen_4prolog_1vector))


tmp15004 := Call(__e, tmp14991, tmp15003)


tmp15005 := Call(__e, PrimFunc(symvector), MakeNumber(0))


tmp15006 := Call(__e, PrimFunc(sym_8v), MakeNumber(0), tmp15005)


tmp15007 := Call(__e, PrimFunc(sym_8v), True, tmp15006)


tmp15008 := Call(__e, tmp15004, tmp15007)


tmp15009 := Call(__e, tmp15008, MakeNumber(0))


tmp15010 := MakeNative(func(__e *ControlFlow) {
__e.Return(True)
return
}, 0)

__e.TailApply(tmp15009, tmp15010)
return


}, 2)

tmp15011 := Call(__e, ns2_1set, symshen_4typecheck, tmp14987)


_ = tmp15011

tmp15012 := MakeNative(func(__e *ControlFlow) {
V4458 := __e.Get(1)
_ = V4458
V4459 := __e.Get(2)
_ = V4459
V4460 := __e.Get(3)
_ = V4460
V4461 := __e.Get(4)
_ = V4461
V4462 := __e.Get(5)
_ = V4462
V4463 := __e.Get(6)
_ = V4463
V4464 := __e.Get(7)
_ = V4464
tmp15019 := Call(__e, PrimFunc(symshen_4unlocked_2), V4462)


var ifres15013 Obj

if True == tmp15019 {
tmp15014 := Call(__e, PrimFunc(symshen_4lazyderef), V4458, V4461)


W446614442 := tmp15014
_ = W446614442

tmp15018 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W446614442, Nil)
}
__typedArg0 := W446614442
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15015 Obj

if True == tmp15018 {
tmp15016 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15016

tmp15017 := Call(__e, PrimFunc(symis_b), V4459, V4460, V4461, V4462, V4463, V4464)


ifres15015 = tmp15017


} else {
ifres15015 = False


}

ifres15013 = ifres15015


} else {
ifres15013 = False


}

W446514441 := ifres15013
_ = W446514441

tmp15033 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W446514441, False)
}
__typedArg0 := W446514441
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15033 {
tmp15031 := Call(__e, PrimFunc(symshen_4unlocked_2), V4462)


if True == tmp15031 {
tmp15020 := Call(__e, PrimFunc(symshen_4lazyderef), V4458, V4461)


W446714443 := tmp15020
_ = W446714443

tmp15029 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W446714443)
}
__typedArg0 := W446714443
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15029 {
tmp15021 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W446714443)
}
__typedArg0 := W446714443
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

W446814444 := tmp15021
_ = W446814444

tmp15022 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W446714443)
}
__typedArg0 := W446714443
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

W446914445 := tmp15022
_ = W446914445

tmp15023 := Call(__e, PrimFunc(symshen_4newpv), V4461)


W447014446 := tmp15023
_ = W447014446

tmp15024 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15024

tmp15025 := Call(__e, PrimFunc(symshen_4deref), W447014446, V4461)


tmp15026 := Call(__e, PrimFunc(symsubst), tmp15025, W446814444, V4459)


tmp15027 := Call(__e, PrimFunc(symshen_4insert_1prolog_1variables), W446914445, tmp15026, V4460, V4461, V4462, V4463, V4464)


__e.TailApply(PrimFunc(symshen_4gc), V4461, tmp15027)
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
__e.Return(W446514441)
return
}


}, 7)

tmp15034 := Call(__e, ns2_1set, symshen_4insert_1prolog_1variables, tmp15012)


_ = tmp15034

tmp15035 := MakeNative(func(__e *ControlFlow) {
V4471 := __e.Get(1)
_ = V4471
V4472 := __e.Get(2)
_ = V4472
V4473 := __e.Get(3)
_ = V4473
V4474 := __e.Get(4)
_ = V4474
V4475 := __e.Get(5)
_ = V4475
V4476 := __e.Get(6)
_ = V4476
W447714447 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V4475)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V4475
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
_ = W447714447

tmp15060 := Call(__e, PrimFunc(symshen_4unlocked_2), V4474)


var ifres15037 Obj

if True == tmp15060 {
tmp15038 := Call(__e, PrimFunc(symshen_4lazyderef), V4471, V4473)


W447914449 := tmp15038
_ = W447914449

tmp15059 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W447914449)
}
__typedArg0 := W447914449
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15039 Obj

if True == tmp15059 {
tmp15040 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W447914449)
}
__typedArg0 := W447914449
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15041 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15040, V4473)


W448014450 := tmp15041
_ = W448014450

tmp15058 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W448014450, symdefine)
}
__typedArg0 := W448014450
__typedArg1 := symdefine
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15042 Obj

if True == tmp15058 {
tmp15043 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W447914449)
}
__typedArg0 := W447914449
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15044 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15043, V4473)


W448114451 := tmp15044
_ = W448114451

tmp15057 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W448114451)
}
__typedArg0 := W448114451
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15045 Obj

if True == tmp15057 {
tmp15046 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W448114451)
}
__typedArg0 := W448114451
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

W448214452 := tmp15046
_ = W448214452

tmp15047 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W448114451)
}
__typedArg0 := W448114451
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

W448314453 := tmp15047
_ = W448314453

tmp15048 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15048

tmp15049 := Call(__e, PrimFunc(symshen_4type_1theory_1enabled_2))


tmp15050 := MakeNative(func(__e *ControlFlow) {
tmp15051 := MakeNative(func(__e *ControlFlow) {
tmp15052 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dspy_d)
}
__typedArg0 := symshen_4_dspy_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp15053 := MakeNative(func(__e *ControlFlow) {
tmp15054 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W448214452, W448314453)
}
__typedArg0 := W448214452
__typedArg1 := W448314453
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15055 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefine, tmp15054)
}
__typedArg0 := symdefine
__typedArg1 := tmp15054
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4t_d), tmp15055, V4472, V4473, V4474, W447714447, V4476)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4signal_1def), tmp15052, W448214452, V4473, V4474, W447714447, tmp15053)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V4473, V4474, W447714447, tmp15051)
return


}, 0)

tmp15056 := Call(__e, PrimFunc(symwhen), tmp15049, V4473, V4474, W447714447, tmp15050)


ifres15045 = tmp15056


} else {
ifres15045 = False


}

ifres15042 = ifres15045


} else {
ifres15042 = False


}

ifres15039 = ifres15042


} else {
ifres15039 = False


}

ifres15037 = ifres15039


} else {
ifres15037 = False


}

W447814448 := ifres15037
_ = W447814448

tmp15072 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W447814448, False)
}
__typedArg0 := W447814448
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15072 {
tmp15068 := Call(__e, PrimFunc(symshen_4unlocked_2), V4474)


var ifres15061 Obj

if True == tmp15068 {
tmp15062 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15062

tmp15063 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp15064 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V4472, Nil)
}
__typedArg0 := V4472
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15065 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp15063, tmp15064)
}
__typedArg0 := tmp15063
__typedArg1 := tmp15064
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15066 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V4471, tmp15065)
}
__typedArg0 := V4471
__typedArg1 := tmp15065
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15067 := Call(__e, PrimFunc(symshen_4system_1S), tmp15066, Nil, V4473, V4474, W447714447, V4476)


ifres15061 = tmp15067


} else {
ifres15061 = False


}

W448414454 := ifres15061
_ = W448414454

tmp15070 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W448414454, False)
}
__typedArg0 := W448414454
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15070 {
__e.TailApply(PrimFunc(symshen_4unlock), V4474, W447714447)
return
} else {
__e.Return(W448414454)
return
}


} else {
__e.Return(W447814448)
return
}


}, 6)

tmp15073 := Call(__e, ns2_1set, symshen_4toplevel_1forms, tmp15035)


_ = tmp15073

tmp15074 := MakeNative(func(__e *ControlFlow) {
V4485 := __e.Get(1)
_ = V4485
V4486 := __e.Get(2)
_ = V4486
V4487 := __e.Get(3)
_ = V4487
V4488 := __e.Get(4)
_ = V4488
V4489 := __e.Get(5)
_ = V4489
V4490 := __e.Get(6)
_ = V4490
tmp15081 := Call(__e, PrimFunc(symshen_4unlocked_2), V4488)


var ifres15075 Obj

if True == tmp15081 {
tmp15076 := Call(__e, PrimFunc(symshen_4lazyderef), V4485, V4487)


W449214456 := tmp15076
_ = W449214456

tmp15080 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W449214456, False)
}
__typedArg0 := W449214456
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15077 Obj

if True == tmp15080 {
tmp15078 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15078

tmp15079 := Call(__e, PrimFunc(symthaw), V4490)


ifres15077 = tmp15079


} else {
ifres15077 = False


}

ifres15075 = ifres15077


} else {
ifres15075 = False


}

W449114455 := ifres15075
_ = W449114455

tmp15096 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W449114455, False)
}
__typedArg0 := W449114455
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15096 {
tmp15094 := Call(__e, PrimFunc(symshen_4unlocked_2), V4488)


if True == tmp15094 {
tmp15082 := Call(__e, PrimFunc(symshen_4lazyderef), V4485, V4487)


W449314457 := tmp15082
_ = W449314457

tmp15092 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W449314457, True)
}
__typedArg0 := W449314457
__typedArg1 := True
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15092 {
tmp15083 := Call(__e, PrimFunc(symshen_4newpv), V4487)


W449414458 := tmp15083
_ = W449414458

tmp15084 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15084

tmp15085 := Call(__e, PrimFunc(symshen_4deref), V4486, V4487)


tmp15086 := Call(__e, PrimFunc(symshen_4app), tmp15085, MakeString(")\n"), symshen_4a)


tmp15087 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("\ntypechecking (fn "))
__typedS1, __typedOK1 := TypedString(tmp15086)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("\ntypechecking (fn ")
__typedArg1 := tmp15086
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp15088 := Call(__e, PrimFunc(symstoutput))


tmp15089 := Call(__e, PrimFunc(sympr), tmp15087, tmp15088)


tmp15090 := Call(__e, PrimFunc(symis), W449414458, tmp15089, V4487, V4488, V4489, V4490)


__e.TailApply(PrimFunc(symshen_4gc), V4487, tmp15090)
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
__e.Return(W449114455)
return
}


}, 6)

tmp15097 := Call(__e, ns2_1set, symshen_4signal_1def, tmp15074)


_ = tmp15097

tmp15098 := MakeNative(func(__e *ControlFlow) {
V4495 := __e.Get(1)
_ = V4495
tmp15099 := Call(__e, PrimFunc(symshen_4curry_1type), V4495)


__e.TailApply(PrimFunc(symshen_4demodulate), tmp15099)
return


}, 1)

tmp15100 := Call(__e, ns2_1set, symshen_4rectify_1type, tmp15098)


_ = tmp15100

tmp15101 := MakeNative(func(__e *ControlFlow) {
V4496 := __e.Get(1)
_ = V4496
tmp15102 := MakeNative(func(__e *ControlFlow) {
tmp15103 := MakeNative(func(__e *ControlFlow) {
Z4498 := __e.Get(1)
_ = Z4498
__e.TailApply(PrimFunc(symshen_4demod), Z4498)
return
}, 1)

tmp15104 := Call(__e, PrimFunc(symshen_4walk), tmp15103, V4496)


W449714459 := tmp15104
_ = W449714459

tmp15106 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W449714459, V4496)
}
__typedArg0 := W449714459
__typedArg1 := V4496
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15106 {
__e.Return(V4496)
return
} else {
__e.TailApply(PrimFunc(symshen_4demodulate), W449714459)
return
}


}, 0)

tmp15107 := MakeNative(func(__e *ControlFlow) {
Z4499 := __e.Get(1)
_ = Z4499
__e.Return(V4496)
return
}, 1)

__e.TailApply(try_1catch, tmp15102, tmp15107)
return


}, 1)

tmp15108 := Call(__e, ns2_1set, symshen_4demodulate, tmp15101)


_ = tmp15108

tmp15109 := MakeNative(func(__e *ControlFlow) {
V4500 := __e.Get(1)
_ = V4500
tmp15233 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4500)
}
__typedArg0 := V4500
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15206 Obj

if True == tmp15233 {
tmp15231 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4500)
}
__typedArg0 := V4500
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15232 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp15231)
}
__typedArg0 := tmp15231
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15208 Obj

if True == tmp15232 {
tmp15228 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4500)
}
__typedArg0 := V4500
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15229 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp15228)
}
__typedArg0 := tmp15228
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15230 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_1_1_6, tmp15229)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp15229
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15210 Obj

if True == tmp15230 {
tmp15225 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4500)
}
__typedArg0 := V4500
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15226 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp15225)
}
__typedArg0 := tmp15225
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15227 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp15226)
}
__typedArg0 := tmp15226
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15212 Obj

if True == tmp15227 {
tmp15221 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4500)
}
__typedArg0 := V4500
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15222 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp15221)
}
__typedArg0 := tmp15221
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15223 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp15222)
}
__typedArg0 := tmp15222
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15224 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp15223)
}
__typedArg0 := tmp15223
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15214 Obj

if True == tmp15224 {
tmp15216 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4500)
}
__typedArg0 := V4500
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15217 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp15216)
}
__typedArg0 := tmp15216
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15218 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp15217)
}
__typedArg0 := tmp15217
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15219 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp15218)
}
__typedArg0 := tmp15218
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15220 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_1_1_6, tmp15219)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp15219
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15215 Obj

if True == tmp15220 {
ifres15215 = True


} else {
ifres15215 = False


}

ifres15214 = ifres15215


} else {
ifres15214 = False


}

var ifres15213 Obj

if True == ifres15214 {
ifres15213 = True


} else {
ifres15213 = False


}

ifres15212 = ifres15213


} else {
ifres15212 = False


}

var ifres15211 Obj

if True == ifres15212 {
ifres15211 = True


} else {
ifres15211 = False


}

ifres15210 = ifres15211


} else {
ifres15210 = False


}

var ifres15209 Obj

if True == ifres15210 {
ifres15209 = True


} else {
ifres15209 = False


}

ifres15208 = ifres15209


} else {
ifres15208 = False


}

var ifres15207 Obj

if True == ifres15208 {
ifres15207 = True


} else {
ifres15207 = False


}

ifres15206 = ifres15207


} else {
ifres15206 = False


}

if True == ifres15206 {
tmp15110 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4500)
}
__typedArg0 := V4500
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15111 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4500)
}
__typedArg0 := V4500
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15112 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp15111)
}
__typedArg0 := tmp15111
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15113 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp15112, Nil)
}
__typedArg0 := tmp15112
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15114 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp15113)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp15113
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15115 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp15110, tmp15114)
}
__typedArg0 := tmp15110
__typedArg1 := tmp15114
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4curry_1type), tmp15115)
return


} else {
tmp15204 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4500)
}
__typedArg0 := V4500
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15164 Obj

if True == tmp15204 {
tmp15202 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4500)
}
__typedArg0 := V4500
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15203 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp15202)
}
__typedArg0 := tmp15202
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15166 Obj

if True == tmp15203 {
tmp15199 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4500)
}
__typedArg0 := V4500
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15200 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp15199)
}
__typedArg0 := tmp15199
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15201 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlist, tmp15200)
}
__typedArg0 := symlist
__typedArg1 := tmp15200
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15168 Obj

if True == tmp15201 {
tmp15196 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4500)
}
__typedArg0 := V4500
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15197 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp15196)
}
__typedArg0 := tmp15196
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15198 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp15197)
}
__typedArg0 := tmp15197
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15170 Obj

if True == tmp15198 {
tmp15192 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4500)
}
__typedArg0 := V4500
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15193 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp15192)
}
__typedArg0 := tmp15192
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15194 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp15193)
}
__typedArg0 := tmp15193
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15195 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp15194)
}
__typedArg0 := Nil
__typedArg1 := tmp15194
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15172 Obj

if True == tmp15195 {
tmp15190 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4500)
}
__typedArg0 := V4500
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15191 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp15190)
}
__typedArg0 := tmp15190
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15174 Obj

if True == tmp15191 {
tmp15187 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4500)
}
__typedArg0 := V4500
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15188 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp15187)
}
__typedArg0 := tmp15187
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15189 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_a_a_6, tmp15188)
}
__typedArg0 := sym_a_a_6
__typedArg1 := tmp15188
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15176 Obj

if True == tmp15189 {
tmp15184 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4500)
}
__typedArg0 := V4500
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15185 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp15184)
}
__typedArg0 := tmp15184
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15186 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp15185)
}
__typedArg0 := tmp15185
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15178 Obj

if True == tmp15186 {
tmp15180 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4500)
}
__typedArg0 := V4500
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15181 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp15180)
}
__typedArg0 := tmp15180
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15182 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp15181)
}
__typedArg0 := tmp15181
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15183 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp15182)
}
__typedArg0 := Nil
__typedArg1 := tmp15182
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15179 Obj

if True == tmp15183 {
ifres15179 = True


} else {
ifres15179 = False


}

ifres15178 = ifres15179


} else {
ifres15178 = False


}

var ifres15177 Obj

if True == ifres15178 {
ifres15177 = True


} else {
ifres15177 = False


}

ifres15176 = ifres15177


} else {
ifres15176 = False


}

var ifres15175 Obj

if True == ifres15176 {
ifres15175 = True


} else {
ifres15175 = False


}

ifres15174 = ifres15175


} else {
ifres15174 = False


}

var ifres15173 Obj

if True == ifres15174 {
ifres15173 = True


} else {
ifres15173 = False


}

ifres15172 = ifres15173


} else {
ifres15172 = False


}

var ifres15171 Obj

if True == ifres15172 {
ifres15171 = True


} else {
ifres15171 = False


}

ifres15170 = ifres15171


} else {
ifres15170 = False


}

var ifres15169 Obj

if True == ifres15170 {
ifres15169 = True


} else {
ifres15169 = False


}

ifres15168 = ifres15169


} else {
ifres15168 = False


}

var ifres15167 Obj

if True == ifres15168 {
ifres15167 = True


} else {
ifres15167 = False


}

ifres15166 = ifres15167


} else {
ifres15166 = False


}

var ifres15165 Obj

if True == ifres15166 {
ifres15165 = True


} else {
ifres15165 = False


}

ifres15164 = ifres15165


} else {
ifres15164 = False


}

if True == ifres15164 {
tmp15116 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4500)
}
__typedArg0 := V4500
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15117 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4500)
}
__typedArg0 := V4500
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15118 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4500)
}
__typedArg0 := V4500
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15119 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp15118)
}
__typedArg0 := tmp15118
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15120 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp15117, tmp15119)
}
__typedArg0 := tmp15117
__typedArg1 := tmp15119
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15121 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstr, tmp15120)
}
__typedArg0 := symstr
__typedArg1 := tmp15120
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15122 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp15121, Nil)
}
__typedArg0 := tmp15121
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15123 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp15122)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp15122
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15124 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp15116, tmp15123)
}
__typedArg0 := tmp15116
__typedArg1 := tmp15123
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4curry_1type), tmp15124)
return


} else {
tmp15162 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4500)
}
__typedArg0 := V4500
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15135 Obj

if True == tmp15162 {
tmp15160 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4500)
}
__typedArg0 := V4500
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15161 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp15160)
}
__typedArg0 := tmp15160
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15137 Obj

if True == tmp15161 {
tmp15157 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4500)
}
__typedArg0 := V4500
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15158 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp15157)
}
__typedArg0 := tmp15157
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15159 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_d, tmp15158)
}
__typedArg0 := sym_d
__typedArg1 := tmp15158
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15139 Obj

if True == tmp15159 {
tmp15154 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4500)
}
__typedArg0 := V4500
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15155 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp15154)
}
__typedArg0 := tmp15154
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15156 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp15155)
}
__typedArg0 := tmp15155
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15141 Obj

if True == tmp15156 {
tmp15150 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4500)
}
__typedArg0 := V4500
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15151 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp15150)
}
__typedArg0 := tmp15150
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15152 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp15151)
}
__typedArg0 := tmp15151
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15153 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp15152)
}
__typedArg0 := tmp15152
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15143 Obj

if True == tmp15153 {
tmp15145 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4500)
}
__typedArg0 := V4500
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15146 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp15145)
}
__typedArg0 := tmp15145
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15147 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp15146)
}
__typedArg0 := tmp15146
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15148 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp15147)
}
__typedArg0 := tmp15147
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15149 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_d, tmp15148)
}
__typedArg0 := sym_d
__typedArg1 := tmp15148
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15144 Obj

if True == tmp15149 {
ifres15144 = True


} else {
ifres15144 = False


}

ifres15143 = ifres15144


} else {
ifres15143 = False


}

var ifres15142 Obj

if True == ifres15143 {
ifres15142 = True


} else {
ifres15142 = False


}

ifres15141 = ifres15142


} else {
ifres15141 = False


}

var ifres15140 Obj

if True == ifres15141 {
ifres15140 = True


} else {
ifres15140 = False


}

ifres15139 = ifres15140


} else {
ifres15139 = False


}

var ifres15138 Obj

if True == ifres15139 {
ifres15138 = True


} else {
ifres15138 = False


}

ifres15137 = ifres15138


} else {
ifres15137 = False


}

var ifres15136 Obj

if True == ifres15137 {
ifres15136 = True


} else {
ifres15136 = False


}

ifres15135 = ifres15136


} else {
ifres15135 = False


}

if True == ifres15135 {
tmp15125 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4500)
}
__typedArg0 := V4500
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15126 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4500)
}
__typedArg0 := V4500
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15127 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp15126)
}
__typedArg0 := tmp15126
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15128 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp15127, Nil)
}
__typedArg0 := tmp15127
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15129 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_d, tmp15128)
}
__typedArg0 := sym_d
__typedArg1 := tmp15128
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15130 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp15125, tmp15129)
}
__typedArg0 := tmp15125
__typedArg1 := tmp15129
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4curry_1type), tmp15130)
return


} else {
tmp15133 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4500)
}
__typedArg0 := V4500
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15133 {
tmp15131 := MakeNative(func(__e *ControlFlow) {
Z4501 := __e.Get(1)
_ = Z4501
__e.TailApply(PrimFunc(symshen_4curry_1type), Z4501)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp15131, V4500)
return


} else {
__e.Return(V4500)
return
}


}


}


}


}, 1)

tmp15234 := Call(__e, ns2_1set, symshen_4curry_1type, tmp15109)


_ = tmp15234

tmp15235 := MakeNative(func(__e *ControlFlow) {
V4502 := __e.Get(1)
_ = V4502
tmp15324 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4502)
}
__typedArg0 := V4502
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15316 Obj

if True == tmp15324 {
tmp15322 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4502)
}
__typedArg0 := V4502
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15323 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefine, tmp15322)
}
__typedArg0 := symdefine
__typedArg1 := tmp15322
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15318 Obj

if True == tmp15323 {
tmp15320 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4502)
}
__typedArg0 := V4502
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15321 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp15320)
}
__typedArg0 := tmp15320
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15319 Obj

if True == tmp15321 {
ifres15319 = True


} else {
ifres15319 = False


}

ifres15318 = ifres15319


} else {
ifres15318 = False


}

var ifres15317 Obj

if True == ifres15318 {
ifres15317 = True


} else {
ifres15317 = False


}

ifres15316 = ifres15317


} else {
ifres15316 = False


}

if True == ifres15316 {
__e.Return(V4502)
return
} else {
tmp15314 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4502)
}
__typedArg0 := V4502
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15295 Obj

if True == tmp15314 {
tmp15312 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4502)
}
__typedArg0 := V4502
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15313 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symtype, tmp15312)
}
__typedArg0 := symtype
__typedArg1 := tmp15312
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15297 Obj

if True == tmp15313 {
tmp15310 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4502)
}
__typedArg0 := V4502
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15311 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp15310)
}
__typedArg0 := tmp15310
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15299 Obj

if True == tmp15311 {
tmp15307 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4502)
}
__typedArg0 := V4502
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15308 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp15307)
}
__typedArg0 := tmp15307
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15309 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp15308)
}
__typedArg0 := tmp15308
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15301 Obj

if True == tmp15309 {
tmp15303 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4502)
}
__typedArg0 := V4502
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15304 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp15303)
}
__typedArg0 := tmp15303
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15305 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp15304)
}
__typedArg0 := tmp15304
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15306 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp15305)
}
__typedArg0 := Nil
__typedArg1 := tmp15305
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15302 Obj

if True == tmp15306 {
ifres15302 = True


} else {
ifres15302 = False


}

ifres15301 = ifres15302


} else {
ifres15301 = False


}

var ifres15300 Obj

if True == ifres15301 {
ifres15300 = True


} else {
ifres15300 = False


}

ifres15299 = ifres15300


} else {
ifres15299 = False


}

var ifres15298 Obj

if True == ifres15299 {
ifres15298 = True


} else {
ifres15298 = False


}

ifres15297 = ifres15298


} else {
ifres15297 = False


}

var ifres15296 Obj

if True == ifres15297 {
ifres15296 = True


} else {
ifres15296 = False


}

ifres15295 = ifres15296


} else {
ifres15295 = False


}

if True == ifres15295 {
tmp15236 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4502)
}
__typedArg0 := V4502
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15237 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp15236)
}
__typedArg0 := tmp15236
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15238 := Call(__e, PrimFunc(symshen_4curry), tmp15237)


tmp15239 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4502)
}
__typedArg0 := V4502
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15240 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp15239)
}
__typedArg0 := tmp15239
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15241 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp15238, tmp15240)
}
__typedArg0 := tmp15238
__typedArg1 := tmp15240
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtype, tmp15241)
}
__typedArg0 := symtype
__typedArg1 := tmp15241
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp15293 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4502)
}
__typedArg0 := V4502
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15289 Obj

if True == tmp15293 {
tmp15291 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4502)
}
__typedArg0 := V4502
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15292 := Call(__e, PrimFunc(symshen_4special_2), tmp15291)


var ifres15290 Obj

if True == tmp15292 {
ifres15290 = True


} else {
ifres15290 = False


}

ifres15289 = ifres15290


} else {
ifres15289 = False


}

if True == ifres15289 {
tmp15242 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4502)
}
__typedArg0 := V4502
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15243 := MakeNative(func(__e *ControlFlow) {
Z4503 := __e.Get(1)
_ = Z4503
__e.TailApply(PrimFunc(symshen_4curry), Z4503)
return
}, 1)

tmp15244 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4502)
}
__typedArg0 := V4502
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15245 := Call(__e, PrimFunc(symmap), tmp15243, tmp15244)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp15242, tmp15245)
}
__typedArg0 := tmp15242
__typedArg1 := tmp15245
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp15287 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4502)
}
__typedArg0 := V4502
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15283 Obj

if True == tmp15287 {
tmp15285 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4502)
}
__typedArg0 := V4502
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15286 := Call(__e, PrimFunc(symshen_4extraspecial_2), tmp15285)


var ifres15284 Obj

if True == tmp15286 {
ifres15284 = True


} else {
ifres15284 = False


}

ifres15283 = ifres15284


} else {
ifres15283 = False


}

if True == ifres15283 {
__e.Return(V4502)
return
} else {
tmp15281 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4502)
}
__typedArg0 := V4502
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15272 Obj

if True == tmp15281 {
tmp15279 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4502)
}
__typedArg0 := V4502
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15280 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp15279)
}
__typedArg0 := tmp15279
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15274 Obj

if True == tmp15280 {
tmp15276 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4502)
}
__typedArg0 := V4502
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15277 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp15276)
}
__typedArg0 := tmp15276
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15278 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp15277)
}
__typedArg0 := tmp15277
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15275 Obj

if True == tmp15278 {
ifres15275 = True


} else {
ifres15275 = False


}

ifres15274 = ifres15275


} else {
ifres15274 = False


}

var ifres15273 Obj

if True == ifres15274 {
ifres15273 = True


} else {
ifres15273 = False


}

ifres15272 = ifres15273


} else {
ifres15272 = False


}

if True == ifres15272 {
tmp15246 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4502)
}
__typedArg0 := V4502
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15247 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4502)
}
__typedArg0 := V4502
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15248 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp15247)
}
__typedArg0 := tmp15247
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15249 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp15248, Nil)
}
__typedArg0 := tmp15248
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15250 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp15246, tmp15249)
}
__typedArg0 := tmp15246
__typedArg1 := tmp15249
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15251 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4502)
}
__typedArg0 := V4502
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15252 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp15251)
}
__typedArg0 := tmp15251
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15253 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp15250, tmp15252)
}
__typedArg0 := tmp15250
__typedArg1 := tmp15252
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4curry), tmp15253)
return


} else {
tmp15270 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4502)
}
__typedArg0 := V4502
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15261 Obj

if True == tmp15270 {
tmp15268 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4502)
}
__typedArg0 := V4502
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15269 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp15268)
}
__typedArg0 := tmp15268
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15263 Obj

if True == tmp15269 {
tmp15265 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4502)
}
__typedArg0 := V4502
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15266 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp15265)
}
__typedArg0 := tmp15265
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15267 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp15266)
}
__typedArg0 := Nil
__typedArg1 := tmp15266
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15264 Obj

if True == tmp15267 {
ifres15264 = True


} else {
ifres15264 = False


}

ifres15263 = ifres15264


} else {
ifres15263 = False


}

var ifres15262 Obj

if True == ifres15263 {
ifres15262 = True


} else {
ifres15262 = False


}

ifres15261 = ifres15262


} else {
ifres15261 = False


}

if True == ifres15261 {
tmp15254 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4502)
}
__typedArg0 := V4502
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15255 := Call(__e, PrimFunc(symshen_4curry), tmp15254)


tmp15256 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4502)
}
__typedArg0 := V4502
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15257 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp15256)
}
__typedArg0 := tmp15256
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15258 := Call(__e, PrimFunc(symshen_4curry), tmp15257)


tmp15259 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp15258, Nil)
}
__typedArg0 := tmp15258
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp15255, tmp15259)
}
__typedArg0 := tmp15255
__typedArg1 := tmp15259
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return(V4502)
return
}


}


}


}


}


}


}, 1)

tmp15325 := Call(__e, ns2_1set, symshen_4curry, tmp15235)


_ = tmp15325

tmp15326 := MakeNative(func(__e *ControlFlow) {
V4504 := __e.Get(1)
_ = V4504
tmp15327 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dspecial_d)
}
__typedArg0 := symshen_4_dspecial_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symelement_2), V4504, tmp15327)
return


}, 1)

tmp15328 := Call(__e, ns2_1set, symshen_4special_2, tmp15326)


_ = tmp15328

tmp15329 := MakeNative(func(__e *ControlFlow) {
V4505 := __e.Get(1)
_ = V4505
tmp15330 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dextraspecial_d)
}
__typedArg0 := symshen_4_dextraspecial_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symelement_2), V4505, tmp15330)
return


}, 1)

tmp15331 := Call(__e, ns2_1set, symshen_4extraspecial_2, tmp15329)


_ = tmp15331

tmp15332 := MakeNative(func(__e *ControlFlow) {
V4506 := __e.Get(1)
_ = V4506
V4507 := __e.Get(2)
_ = V4507
V4508 := __e.Get(3)
_ = V4508
V4509 := __e.Get(4)
_ = V4509
V4510 := __e.Get(5)
_ = V4510
V4511 := __e.Get(6)
_ = V4511
W451214460 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V4510)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V4510
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
_ = W451214460

tmp15338 := Call(__e, PrimFunc(symshen_4unlocked_2), V4509)


var ifres15334 Obj

if True == tmp15338 {
tmp15335 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15335

tmp15336 := Call(__e, PrimFunc(symshen_4maxinfexceeded_2))


tmp15337 := Call(__e, PrimFunc(symwhen), tmp15336, V4508, V4509, W451214460, V4511)


ifres15334 = tmp15337


} else {
ifres15334 = False


}

W451314461 := ifres15334
_ = W451314461

tmp15386 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W451314461, False)
}
__typedArg0 := W451314461
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15386 {
tmp15367 := Call(__e, PrimFunc(symshen_4unlocked_2), V4509)


var ifres15339 Obj

if True == tmp15367 {
tmp15340 := Call(__e, PrimFunc(symshen_4lazyderef), V4506, V4508)


W451514463 := tmp15340
_ = W451514463

tmp15366 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W451514463)
}
__typedArg0 := W451514463
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15341 Obj

if True == tmp15366 {
tmp15342 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W451514463)
}
__typedArg0 := W451514463
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

W451614464 := tmp15342
_ = W451614464

tmp15343 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W451514463)
}
__typedArg0 := W451514463
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15344 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15343, V4508)


W451714465 := tmp15344
_ = W451714465

tmp15365 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W451714465)
}
__typedArg0 := W451714465
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15345 Obj

if True == tmp15365 {
tmp15346 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W451714465)
}
__typedArg0 := W451714465
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

W451814466 := tmp15346
_ = W451814466

tmp15347 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W451714465)
}
__typedArg0 := W451714465
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15348 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15347, V4508)


W451914467 := tmp15348
_ = W451914467

tmp15364 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W451914467)
}
__typedArg0 := W451914467
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15349 Obj

if True == tmp15364 {
tmp15350 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W451914467)
}
__typedArg0 := W451914467
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

W452014468 := tmp15350
_ = W452014468

tmp15351 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W451914467)
}
__typedArg0 := W451914467
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15352 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15351, V4508)


W452114469 := tmp15352
_ = W452114469

tmp15363 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W452114469, Nil)
}
__typedArg0 := W452114469
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15353 Obj

if True == tmp15363 {
tmp15354 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15354

tmp15355 := Call(__e, PrimFunc(symshen_4deref), W451814466, V4508)


tmp15356 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp15357 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp15355, tmp15356)
}
__typedArg0 := tmp15355
__typedArg1 := tmp15356
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

tmp15358 := MakeNative(func(__e *ControlFlow) {
tmp15359 := Call(__e, PrimFunc(symshen_4type_1theory_1enabled_2))


tmp15360 := MakeNative(func(__e *ControlFlow) {
tmp15361 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), W451614464, W452014468, V4507, V4508, V4509, W451214460, V4511)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V4508, V4509, W451214460, tmp15361)
return


}, 0)

__e.TailApply(PrimFunc(symwhen), tmp15359, V4508, V4509, W451214460, tmp15360)
return


}, 0)

tmp15362 := Call(__e, PrimFunc(symwhen), tmp15357, V4508, V4509, W451214460, tmp15358)


ifres15353 = tmp15362


} else {
ifres15353 = False


}

ifres15349 = ifres15353


} else {
ifres15349 = False


}

ifres15345 = ifres15349


} else {
ifres15345 = False


}

ifres15341 = ifres15345


} else {
ifres15341 = False


}

ifres15339 = ifres15341


} else {
ifres15339 = False


}

W451414462 := ifres15339
_ = W451414462

tmp15384 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W451414462, False)
}
__typedArg0 := W451414462
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15384 {
tmp15373 := Call(__e, PrimFunc(symshen_4unlocked_2), V4509)


var ifres15368 Obj

if True == tmp15373 {
tmp15369 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15369

tmp15370 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dspy_d)
}
__typedArg0 := symshen_4_dspy_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp15371 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4show), V4506, V4507, V4508, V4509, W451214460, V4511)
return
}, 0)

tmp15372 := Call(__e, PrimFunc(symwhen), tmp15370, V4508, V4509, W451214460, tmp15371)


ifres15368 = tmp15372


} else {
ifres15368 = False


}

W452214470 := ifres15368
_ = W452214470

tmp15382 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W452214470, False)
}
__typedArg0 := W452214470
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15382 {
tmp15378 := Call(__e, PrimFunc(symshen_4unlocked_2), V4509)


var ifres15374 Obj

if True == tmp15378 {
tmp15375 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15375

tmp15376 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_ddatatypes_d)
}
__typedArg0 := symshen_4_ddatatypes_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp15377 := Call(__e, PrimFunc(symshen_4search_1user_1datatypes), V4506, V4507, tmp15376, V4508, V4509, W451214460, V4511)


ifres15374 = tmp15377


} else {
ifres15374 = False


}

W452314471 := ifres15374
_ = W452314471

tmp15380 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W452314471, False)
}
__typedArg0 := W452314471
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15380 {
__e.TailApply(PrimFunc(symshen_4unlock), V4509, W451214460)
return
} else {
__e.Return(W452314471)
return
}


} else {
__e.Return(W452214470)
return
}


} else {
__e.Return(W451414462)
return
}


} else {
__e.Return(W451314461)
return
}


}, 6)

tmp15387 := Call(__e, ns2_1set, symshen_4system_1S, tmp15332)


_ = tmp15387

tmp15388 := MakeNative(func(__e *ControlFlow) {
V4530 := __e.Get(1)
_ = V4530
V4531 := __e.Get(2)
_ = V4531
V4532 := __e.Get(3)
_ = V4532
V4533 := __e.Get(4)
_ = V4533
V4534 := __e.Get(5)
_ = V4534
V4535 := __e.Get(6)
_ = V4535
tmp15389 := Call(__e, PrimFunc(symshen_4line))


_ = tmp15389

tmp15390 := Call(__e, PrimFunc(symshen_4deref), V4530, V4532)


tmp15391 := Call(__e, PrimFunc(symshen_4show_1p), tmp15390)


_ = tmp15391

tmp15392 := Call(__e, PrimFunc(symnl), MakeNumber(2))


_ = tmp15392

tmp15393 := Call(__e, PrimFunc(symshen_4deref), V4531, V4532)


tmp15394 := Call(__e, PrimFunc(symshen_4show_1assumptions), tmp15393, MakeNumber(1))


_ = tmp15394

tmp15395 := Call(__e, PrimFunc(symshen_4pause_1for_1user))


_ = tmp15395

__e.Return(False)
return


}, 6)

tmp15396 := Call(__e, ns2_1set, symshen_4show, tmp15388)


_ = tmp15396

tmp15397 := MakeNative(func(__e *ControlFlow) {
tmp15398 := Call(__e, PrimFunc(syminferences))


W453614472 := tmp15398
_ = W453614472

tmp15400 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeNumber(1), W453614472)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := W453614472
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15399 Obj

if True == tmp15400 {
ifres15399 = MakeString("")


} else {
ifres15399 = MakeString("s")


}

tmp15401 := Call(__e, PrimFunc(symshen_4app), ifres15399, MakeString(" \n?- "), symshen_4a)


tmp15403 := Call(__e, PrimFunc(symshen_4app), W453614472, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(" inference"))
__typedS1, __typedOK1 := TypedString(tmp15401)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(" inference")
__typedArg1 := tmp15401
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


tmp15404 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("____________________________________________________________ "))
__typedS1, __typedOK1 := TypedString(tmp15403)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("____________________________________________________________ ")
__typedArg1 := tmp15403
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp15405 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp15404, tmp15405)
return


}, 0)

tmp15406 := Call(__e, ns2_1set, symshen_4line, tmp15397)


_ = tmp15406

tmp15407 := MakeNative(func(__e *ControlFlow) {
V4537 := __e.Get(1)
_ = V4537
tmp15439 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4537)
}
__typedArg0 := V4537
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15418 Obj

if True == tmp15439 {
tmp15437 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4537)
}
__typedArg0 := V4537
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15438 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp15437)
}
__typedArg0 := tmp15437
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15420 Obj

if True == tmp15438 {
tmp15434 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4537)
}
__typedArg0 := V4537
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15435 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp15434)
}
__typedArg0 := tmp15434
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15436 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp15435)
}
__typedArg0 := tmp15435
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15422 Obj

if True == tmp15436 {
tmp15430 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4537)
}
__typedArg0 := V4537
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15431 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp15430)
}
__typedArg0 := tmp15430
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15432 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp15431)
}
__typedArg0 := tmp15431
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15433 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp15432)
}
__typedArg0 := Nil
__typedArg1 := tmp15432
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15424 Obj

if True == tmp15433 {
tmp15426 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4537)
}
__typedArg0 := V4537
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15427 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp15426)
}
__typedArg0 := tmp15426
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15428 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp15429 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp15427, tmp15428)
}
__typedArg0 := tmp15427
__typedArg1 := tmp15428
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15425 Obj

if True == tmp15429 {
ifres15425 = True


} else {
ifres15425 = False


}

ifres15424 = ifres15425


} else {
ifres15424 = False


}

var ifres15423 Obj

if True == ifres15424 {
ifres15423 = True


} else {
ifres15423 = False


}

ifres15422 = ifres15423


} else {
ifres15422 = False


}

var ifres15421 Obj

if True == ifres15422 {
ifres15421 = True


} else {
ifres15421 = False


}

ifres15420 = ifres15421


} else {
ifres15420 = False


}

var ifres15419 Obj

if True == ifres15420 {
ifres15419 = True


} else {
ifres15419 = False


}

ifres15418 = ifres15419


} else {
ifres15418 = False


}

if True == ifres15418 {
tmp15408 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4537)
}
__typedArg0 := V4537
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15409 := Call(__e, PrimFunc(symshen_4prterm), tmp15408)


_ = tmp15409

tmp15410 := Call(__e, PrimFunc(symstoutput))


tmp15411 := Call(__e, PrimFunc(sympr), MakeString(" : "), tmp15410)


_ = tmp15411

tmp15412 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4537)
}
__typedArg0 := V4537
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15413 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp15412)
}
__typedArg0 := tmp15412
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15414 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp15413)
}
__typedArg0 := tmp15413
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15415 := Call(__e, PrimFunc(symshen_4app), tmp15414, MakeString(""), symshen_4r)


tmp15416 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp15415, tmp15416)
return


} else {
__e.TailApply(PrimFunc(symshen_4prterm), V4537)
return
}


}, 1)

tmp15440 := Call(__e, ns2_1set, symshen_4show_1p, tmp15407)


_ = tmp15440

tmp15441 := MakeNative(func(__e *ControlFlow) {
V4538 := __e.Get(1)
_ = V4538
tmp15484 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4538)
}
__typedArg0 := V4538
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15465 Obj

if True == tmp15484 {
tmp15482 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4538)
}
__typedArg0 := V4538
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15483 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcons, tmp15482)
}
__typedArg0 := symcons
__typedArg1 := tmp15482
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15467 Obj

if True == tmp15483 {
tmp15480 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4538)
}
__typedArg0 := V4538
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15481 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp15480)
}
__typedArg0 := tmp15480
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15469 Obj

if True == tmp15481 {
tmp15477 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4538)
}
__typedArg0 := V4538
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15478 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp15477)
}
__typedArg0 := tmp15477
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15479 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp15478)
}
__typedArg0 := tmp15478
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15471 Obj

if True == tmp15479 {
tmp15473 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4538)
}
__typedArg0 := V4538
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15474 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp15473)
}
__typedArg0 := tmp15473
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15475 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp15474)
}
__typedArg0 := tmp15474
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15476 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp15475)
}
__typedArg0 := Nil
__typedArg1 := tmp15475
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15472 Obj

if True == tmp15476 {
ifres15472 = True


} else {
ifres15472 = False


}

ifres15471 = ifres15472


} else {
ifres15471 = False


}

var ifres15470 Obj

if True == ifres15471 {
ifres15470 = True


} else {
ifres15470 = False


}

ifres15469 = ifres15470


} else {
ifres15469 = False


}

var ifres15468 Obj

if True == ifres15469 {
ifres15468 = True


} else {
ifres15468 = False


}

ifres15467 = ifres15468


} else {
ifres15467 = False


}

var ifres15466 Obj

if True == ifres15467 {
ifres15466 = True


} else {
ifres15466 = False


}

ifres15465 = ifres15466


} else {
ifres15465 = False


}

if True == ifres15465 {
tmp15442 := Call(__e, PrimFunc(symstoutput))


tmp15443 := Call(__e, PrimFunc(sympr), MakeString("["), tmp15442)


_ = tmp15443

tmp15444 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4538)
}
__typedArg0 := V4538
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15445 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp15444)
}
__typedArg0 := tmp15444
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15446 := Call(__e, PrimFunc(symshen_4prterm), tmp15445)


_ = tmp15446

tmp15447 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4538)
}
__typedArg0 := V4538
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15448 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp15447)
}
__typedArg0 := tmp15447
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15449 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp15448)
}
__typedArg0 := tmp15448
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15450 := Call(__e, PrimFunc(symshen_4prtl), tmp15449)


_ = tmp15450

tmp15451 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), MakeString("]"), tmp15451)
return


} else {
tmp15463 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4538)
}
__typedArg0 := V4538
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15463 {
tmp15452 := Call(__e, PrimFunc(symstoutput))


tmp15453 := Call(__e, PrimFunc(sympr), MakeString("("), tmp15452)


_ = tmp15453

tmp15454 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4538)
}
__typedArg0 := V4538
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15455 := Call(__e, PrimFunc(symshen_4prterm), tmp15454)


_ = tmp15455

tmp15456 := MakeNative(func(__e *ControlFlow) {
Z4539 := __e.Get(1)
_ = Z4539
tmp15457 := Call(__e, PrimFunc(symstoutput))


tmp15458 := Call(__e, PrimFunc(sympr), MakeString(" "), tmp15457)


_ = tmp15458

__e.TailApply(PrimFunc(symshen_4prterm), Z4539)
return


}, 1)

tmp15459 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4538)
}
__typedArg0 := V4538
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15460 := Call(__e, PrimFunc(symmap), tmp15456, tmp15459)


_ = tmp15460

tmp15461 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), MakeString(")"), tmp15461)
return


} else {
__e.TailApply(PrimFunc(symprint), V4538)
return
}


}


}, 1)

tmp15485 := Call(__e, ns2_1set, symshen_4prterm, tmp15441)


_ = tmp15485

tmp15486 := MakeNative(func(__e *ControlFlow) {
V4540 := __e.Get(1)
_ = V4540
tmp15519 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V4540)
}
__typedArg0 := Nil
__typedArg1 := V4540
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15519 {
__e.Return(MakeString(""))
return
} else {
tmp15517 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4540)
}
__typedArg0 := V4540
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15498 Obj

if True == tmp15517 {
tmp15515 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4540)
}
__typedArg0 := V4540
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15516 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcons, tmp15515)
}
__typedArg0 := symcons
__typedArg1 := tmp15515
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15500 Obj

if True == tmp15516 {
tmp15513 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4540)
}
__typedArg0 := V4540
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15514 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp15513)
}
__typedArg0 := tmp15513
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15502 Obj

if True == tmp15514 {
tmp15510 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4540)
}
__typedArg0 := V4540
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15511 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp15510)
}
__typedArg0 := tmp15510
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15512 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp15511)
}
__typedArg0 := tmp15511
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15504 Obj

if True == tmp15512 {
tmp15506 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4540)
}
__typedArg0 := V4540
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15507 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp15506)
}
__typedArg0 := tmp15506
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15508 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp15507)
}
__typedArg0 := tmp15507
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15509 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp15508)
}
__typedArg0 := Nil
__typedArg1 := tmp15508
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15505 Obj

if True == tmp15509 {
ifres15505 = True


} else {
ifres15505 = False


}

ifres15504 = ifres15505


} else {
ifres15504 = False


}

var ifres15503 Obj

if True == ifres15504 {
ifres15503 = True


} else {
ifres15503 = False


}

ifres15502 = ifres15503


} else {
ifres15502 = False


}

var ifres15501 Obj

if True == ifres15502 {
ifres15501 = True


} else {
ifres15501 = False


}

ifres15500 = ifres15501


} else {
ifres15500 = False


}

var ifres15499 Obj

if True == ifres15500 {
ifres15499 = True


} else {
ifres15499 = False


}

ifres15498 = ifres15499


} else {
ifres15498 = False


}

if True == ifres15498 {
tmp15487 := Call(__e, PrimFunc(symstoutput))


tmp15488 := Call(__e, PrimFunc(sympr), MakeString(" "), tmp15487)


_ = tmp15488

tmp15489 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4540)
}
__typedArg0 := V4540
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15490 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp15489)
}
__typedArg0 := tmp15489
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15491 := Call(__e, PrimFunc(symshen_4prterm), tmp15490)


_ = tmp15491

tmp15492 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4540)
}
__typedArg0 := V4540
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15493 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp15492)
}
__typedArg0 := tmp15492
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15494 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp15493)
}
__typedArg0 := tmp15493
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4prtl), tmp15494)
return


} else {
tmp15495 := Call(__e, PrimFunc(symstoutput))


tmp15496 := Call(__e, PrimFunc(sympr), MakeString(" | "), tmp15495)


_ = tmp15496

__e.TailApply(PrimFunc(symshen_4prterm), V4540)
return


}


}


}, 1)

tmp15520 := Call(__e, ns2_1set, symshen_4prtl, tmp15486)


_ = tmp15520

tmp15521 := MakeNative(func(__e *ControlFlow) {
V4547 := __e.Get(1)
_ = V4547
V4548 := __e.Get(2)
_ = V4548
tmp15534 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V4547)
}
__typedArg0 := Nil
__typedArg1 := V4547
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15534 {
tmp15522 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), MakeString("\n> "), tmp15522)
return


} else {
tmp15532 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4547)
}
__typedArg0 := V4547
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15532 {
tmp15523 := Call(__e, PrimFunc(symshen_4app), V4548, MakeString(". "), symshen_4a)


tmp15524 := Call(__e, PrimFunc(symstoutput))


tmp15525 := Call(__e, PrimFunc(sympr), tmp15523, tmp15524)


_ = tmp15525

tmp15526 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4547)
}
__typedArg0 := V4547
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15527 := Call(__e, PrimFunc(symshen_4show_1p), tmp15526)


_ = tmp15527

tmp15528 := Call(__e, PrimFunc(symnl), MakeNumber(1))


_ = tmp15528

tmp15529 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4547)
}
__typedArg0 := V4547
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4show_1assumptions), tmp15529, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V4548)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V4548
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.show-assumptions"))
}
__typedArg0 := MakeString("implementation error in shen.show-assumptions")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 2)

tmp15535 := Call(__e, ns2_1set, symshen_4show_1assumptions, tmp15521)


_ = tmp15535

tmp15536 := MakeNative(func(__e *ControlFlow) {
tmp15537 := Call(__e, PrimFunc(symstinput))


tmp15538 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symread_1byte) {
return PrimReadByte(tmp15537)
}
__typedArg0 := tmp15537
return Call(__e, PrimFunc(symread_1byte), __typedArg0)
})()

W454914473 := tmp15538
_ = W454914473

tmp15540 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W454914473, MakeNumber(94))
}
__typedArg0 := W454914473
__typedArg1 := MakeNumber(94)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15540 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("input aborted\n"))
}
__typedArg0 := MakeString("input aborted\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
} else {
__e.TailApply(PrimFunc(symnl), MakeNumber(1))
return
}


}, 0)

tmp15541 := Call(__e, ns2_1set, symshen_4pause_1for_1user, tmp15536)


_ = tmp15541

tmp15542 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dshen_1type_1theory_1enabled_2_d)
}
__typedArg0 := symshen_4_dshen_1type_1theory_1enabled_2_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp15543 := Call(__e, ns2_1set, symshen_4type_1theory_1enabled_2, tmp15542)


_ = tmp15543

tmp15544 := MakeNative(func(__e *ControlFlow) {
tmp15546 := Call(__e, PrimFunc(syminferences))


tmp15547 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dmaxinferences_d)
}
__typedArg0 := symshen_4_dmaxinferences_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_6) {
__typedN0, __typedOK0 := TypedFloat64(tmp15546)
__typedN1, __typedOK1 := TypedFloat64(tmp15547)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_6) {
return TypedMaterializeBoolean((__typedN0 > __typedN1))
}}
__typedArg0 := tmp15546
__typedArg1 := tmp15547
return Call(__e, PrimFunc(sym_6), __typedArg0, __typedArg1)
})() {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("maximum inferences exceeded"))
}
__typedArg0 := MakeString("maximum inferences exceeded")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
} else {
__e.Return(False)
return
}


}, 0)

tmp15549 := Call(__e, ns2_1set, symshen_4maxinfexceeded_2, tmp15544)


_ = tmp15549

tmp15550 := MakeNative(func(__e *ControlFlow) {
V4550 := __e.Get(1)
_ = V4550
V4551 := __e.Get(2)
_ = V4551
V4552 := __e.Get(3)
_ = V4552
V4553 := __e.Get(4)
_ = V4553
V4554 := __e.Get(5)
_ = V4554
V4555 := __e.Get(6)
_ = V4555
V4556 := __e.Get(7)
_ = V4556
W455714474 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V4555)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V4555
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
_ = W455714474

tmp15561 := Call(__e, PrimFunc(symshen_4unlocked_2), V4554)


var ifres15552 Obj

if True == tmp15561 {
tmp15553 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15553

tmp15554 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dspy_d)
}
__typedArg0 := symshen_4_dspy_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp15555 := MakeNative(func(__e *ControlFlow) {
tmp15556 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp15557 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V4551, Nil)
}
__typedArg0 := V4551
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15558 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp15556, tmp15557)
}
__typedArg0 := tmp15556
__typedArg1 := tmp15557
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15559 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V4550, tmp15558)
}
__typedArg0 := V4550
__typedArg1 := tmp15558
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4show), tmp15559, V4552, V4553, V4554, W455714474, V4556)
return


}, 0)

tmp15560 := Call(__e, PrimFunc(symwhen), tmp15554, V4553, V4554, W455714474, tmp15555)


ifres15552 = tmp15560


} else {
ifres15552 = False


}

W455814475 := ifres15552
_ = W455814475

tmp16300 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W455814475, False)
}
__typedArg0 := W455814475
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16300 {
tmp15569 := Call(__e, PrimFunc(symshen_4unlocked_2), V4554)


var ifres15562 Obj

if True == tmp15569 {
tmp15563 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15563

tmp15564 := Call(__e, PrimFunc(symshen_4lazyderef), V4550, V4553)


tmp15565 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp15564)
}
__typedArg0 := tmp15564
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

tmp15566 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp15565)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp15565
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

tmp15567 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4primitive), V4550, V4551, V4553, V4554, W455714474, V4556)
return
}, 0)

tmp15568 := Call(__e, PrimFunc(symwhen), tmp15566, V4553, V4554, W455714474, tmp15567)


ifres15562 = tmp15568


} else {
ifres15562 = False


}

W455914476 := ifres15562
_ = W455914476

tmp16298 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W455914476, False)
}
__typedArg0 := W455914476
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16298 {
tmp15573 := Call(__e, PrimFunc(symshen_4unlocked_2), V4554)


var ifres15570 Obj

if True == tmp15573 {
tmp15571 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15571

tmp15572 := Call(__e, PrimFunc(symshen_4by_1hypothesis), V4550, V4551, V4552, V4553, V4554, W455714474, V4556)


ifres15570 = tmp15572


} else {
ifres15570 = False


}

W456014477 := ifres15570
_ = W456014477

tmp16296 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W456014477, False)
}
__typedArg0 := W456014477
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16296 {
tmp15587 := Call(__e, PrimFunc(symshen_4unlocked_2), V4554)


var ifres15574 Obj

if True == tmp15587 {
tmp15575 := Call(__e, PrimFunc(symshen_4lazyderef), V4550, V4553)


W456214479 := tmp15575
_ = W456214479

tmp15586 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W456214479)
}
__typedArg0 := W456214479
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15576 Obj

if True == tmp15586 {
tmp15577 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W456214479)
}
__typedArg0 := W456214479
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

W456314480 := tmp15577
_ = W456314480

tmp15578 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W456214479)
}
__typedArg0 := W456214479
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15579 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15578, V4553)


W456414481 := tmp15579
_ = W456414481

tmp15585 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W456414481, Nil)
}
__typedArg0 := W456414481
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15580 Obj

if True == tmp15585 {
tmp15581 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15581

tmp15582 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V4551, Nil)
}
__typedArg0 := V4551
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15583 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp15582)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp15582
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15584 := Call(__e, PrimFunc(symshen_4lookupsig), W456314480, tmp15583, V4553, V4554, W455714474, V4556)


ifres15580 = tmp15584


} else {
ifres15580 = False


}

ifres15576 = ifres15580


} else {
ifres15576 = False


}

ifres15574 = ifres15576


} else {
ifres15574 = False


}

W456114478 := ifres15574
_ = W456114478

tmp16294 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W456114478, False)
}
__typedArg0 := W456114478
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16294 {
tmp15613 := Call(__e, PrimFunc(symshen_4unlocked_2), V4554)


var ifres15588 Obj

if True == tmp15613 {
tmp15589 := Call(__e, PrimFunc(symshen_4lazyderef), V4550, V4553)


W456614483 := tmp15589
_ = W456614483

tmp15612 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W456614483)
}
__typedArg0 := W456614483
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15590 Obj

if True == tmp15612 {
tmp15591 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W456614483)
}
__typedArg0 := W456614483
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15592 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15591, V4553)


W456714484 := tmp15592
_ = W456714484

tmp15611 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W456714484, symfn)
}
__typedArg0 := W456714484
__typedArg1 := symfn
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15593 Obj

if True == tmp15611 {
tmp15594 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W456614483)
}
__typedArg0 := W456614483
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15595 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15594, V4553)


W456814485 := tmp15595
_ = W456814485

tmp15610 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W456814485)
}
__typedArg0 := W456814485
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15596 Obj

if True == tmp15610 {
tmp15597 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W456814485)
}
__typedArg0 := W456814485
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

W456914486 := tmp15597
_ = W456914486

tmp15598 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W456814485)
}
__typedArg0 := W456814485
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15599 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15598, V4553)


W457014487 := tmp15599
_ = W457014487

tmp15609 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W457014487, Nil)
}
__typedArg0 := W457014487
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15600 Obj

if True == tmp15609 {
tmp15601 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15601

tmp15602 := Call(__e, PrimFunc(symshen_4deref), W456914486, V4553)


tmp15603 := Call(__e, PrimFunc(symarity), tmp15602)


tmp15604 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp15603, MakeNumber(0))
}
__typedArg0 := tmp15603
__typedArg1 := MakeNumber(0)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

tmp15605 := MakeNative(func(__e *ControlFlow) {
tmp15606 := MakeNative(func(__e *ControlFlow) {
tmp15607 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W456914486, Nil)
}
__typedArg0 := W456914486
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4system_1S_1h), tmp15607, V4551, V4552, V4553, V4554, W455714474, V4556)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V4553, V4554, W455714474, tmp15606)
return


}, 0)

tmp15608 := Call(__e, PrimFunc(symwhen), tmp15604, V4553, V4554, W455714474, tmp15605)


ifres15600 = tmp15608


} else {
ifres15600 = False


}

ifres15596 = ifres15600


} else {
ifres15596 = False


}

ifres15593 = ifres15596


} else {
ifres15593 = False


}

ifres15590 = ifres15593


} else {
ifres15590 = False


}

ifres15588 = ifres15590


} else {
ifres15588 = False


}

W456514482 := ifres15588
_ = W456514482

tmp16292 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W456514482, False)
}
__typedArg0 := W456514482
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16292 {
tmp15633 := Call(__e, PrimFunc(symshen_4unlocked_2), V4554)


var ifres15614 Obj

if True == tmp15633 {
tmp15615 := Call(__e, PrimFunc(symshen_4lazyderef), V4550, V4553)


W457214489 := tmp15615
_ = W457214489

tmp15632 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W457214489)
}
__typedArg0 := W457214489
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15616 Obj

if True == tmp15632 {
tmp15617 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W457214489)
}
__typedArg0 := W457214489
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15618 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15617, V4553)


W457314490 := tmp15618
_ = W457314490

tmp15631 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W457314490, symfn)
}
__typedArg0 := W457314490
__typedArg1 := symfn
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15619 Obj

if True == tmp15631 {
tmp15620 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W457214489)
}
__typedArg0 := W457214489
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15621 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15620, V4553)


W457414491 := tmp15621
_ = W457414491

tmp15630 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W457414491)
}
__typedArg0 := W457414491
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15622 Obj

if True == tmp15630 {
tmp15623 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W457414491)
}
__typedArg0 := W457414491
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

W457514492 := tmp15623
_ = W457514492

tmp15624 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W457414491)
}
__typedArg0 := W457414491
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15625 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15624, V4553)


W457614493 := tmp15625
_ = W457614493

tmp15629 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W457614493, Nil)
}
__typedArg0 := W457614493
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15626 Obj

if True == tmp15629 {
tmp15627 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15627

tmp15628 := Call(__e, PrimFunc(symshen_4lookupsig), W457514492, V4551, V4553, V4554, W455714474, V4556)


ifres15626 = tmp15628


} else {
ifres15626 = False


}

ifres15622 = ifres15626


} else {
ifres15622 = False


}

ifres15619 = ifres15622


} else {
ifres15619 = False


}

ifres15616 = ifres15619


} else {
ifres15616 = False


}

ifres15614 = ifres15616


} else {
ifres15614 = False


}

W457114488 := ifres15614
_ = W457114488

tmp16290 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W457114488, False)
}
__typedArg0 := W457114488
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16290 {
tmp15660 := Call(__e, PrimFunc(symshen_4unlocked_2), V4554)


var ifres15634 Obj

if True == tmp15660 {
tmp15635 := Call(__e, PrimFunc(symshen_4lazyderef), V4550, V4553)


W457814495 := tmp15635
_ = W457814495

tmp15659 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W457814495)
}
__typedArg0 := W457814495
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15636 Obj

if True == tmp15659 {
tmp15637 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W457814495)
}
__typedArg0 := W457814495
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

W457914496 := tmp15637
_ = W457914496

tmp15638 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W457814495)
}
__typedArg0 := W457814495
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15639 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15638, V4553)


W458014497 := tmp15639
_ = W458014497

tmp15658 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W458014497)
}
__typedArg0 := W458014497
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15640 Obj

if True == tmp15658 {
tmp15641 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W458014497)
}
__typedArg0 := W458014497
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

W458114498 := tmp15641
_ = W458114498

tmp15642 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W458014497)
}
__typedArg0 := W458014497
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15643 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15642, V4553)


W458214499 := tmp15643
_ = W458214499

tmp15657 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W458214499, Nil)
}
__typedArg0 := W458214499
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15644 Obj

if True == tmp15657 {
tmp15645 := Call(__e, PrimFunc(symshen_4newpv), V4553)


W458314500 := tmp15645
_ = W458314500

tmp15646 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15646

tmp15647 := Call(__e, PrimFunc(symshen_4lazyderef), W457914496, V4553)


tmp15648 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp15647)
}
__typedArg0 := tmp15647
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

tmp15649 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp15648)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp15648
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

tmp15650 := MakeNative(func(__e *ControlFlow) {
tmp15651 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V4551, Nil)
}
__typedArg0 := V4551
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15652 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp15651)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp15651
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15653 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W458314500, tmp15652)
}
__typedArg0 := W458314500
__typedArg1 := tmp15652
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15654 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), W458114498, W458314500, V4552, V4553, V4554, W455714474, V4556)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4lookupsig), W457914496, tmp15653, V4553, V4554, W455714474, tmp15654)
return


}, 0)

tmp15655 := Call(__e, PrimFunc(symwhen), tmp15649, V4553, V4554, W455714474, tmp15650)


tmp15656 := Call(__e, PrimFunc(symshen_4gc), V4553, tmp15655)


ifres15644 = tmp15656


} else {
ifres15644 = False


}

ifres15640 = ifres15644


} else {
ifres15640 = False


}

ifres15636 = ifres15640


} else {
ifres15636 = False


}

ifres15634 = ifres15636


} else {
ifres15634 = False


}

W457714494 := ifres15634
_ = W457714494

tmp16288 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W457714494, False)
}
__typedArg0 := W457714494
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16288 {
tmp15683 := Call(__e, PrimFunc(symshen_4unlocked_2), V4554)


var ifres15661 Obj

if True == tmp15683 {
tmp15662 := Call(__e, PrimFunc(symshen_4lazyderef), V4550, V4553)


W458514502 := tmp15662
_ = W458514502

tmp15682 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W458514502)
}
__typedArg0 := W458514502
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15663 Obj

if True == tmp15682 {
tmp15664 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W458514502)
}
__typedArg0 := W458514502
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

W458614503 := tmp15664
_ = W458614503

tmp15665 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W458514502)
}
__typedArg0 := W458514502
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15666 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15665, V4553)


W458714504 := tmp15666
_ = W458714504

tmp15681 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W458714504)
}
__typedArg0 := W458714504
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15667 Obj

if True == tmp15681 {
tmp15668 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W458714504)
}
__typedArg0 := W458714504
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

W458814505 := tmp15668
_ = W458814505

tmp15669 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W458714504)
}
__typedArg0 := W458714504
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15670 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15669, V4553)


W458914506 := tmp15670
_ = W458914506

tmp15680 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W458914506, Nil)
}
__typedArg0 := W458914506
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15671 Obj

if True == tmp15680 {
tmp15672 := Call(__e, PrimFunc(symshen_4newpv), V4553)


W459014507 := tmp15672
_ = W459014507

tmp15673 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15673

tmp15674 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V4551, Nil)
}
__typedArg0 := V4551
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15675 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp15674)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp15674
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15676 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W459014507, tmp15675)
}
__typedArg0 := W459014507
__typedArg1 := tmp15675
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15677 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), W458814505, W459014507, V4552, V4553, V4554, W455714474, V4556)
return
}, 0)

tmp15678 := Call(__e, PrimFunc(symshen_4system_1S_1h), W458614503, tmp15676, V4552, V4553, V4554, W455714474, tmp15677)


tmp15679 := Call(__e, PrimFunc(symshen_4gc), V4553, tmp15678)


ifres15671 = tmp15679


} else {
ifres15671 = False


}

ifres15667 = ifres15671


} else {
ifres15667 = False


}

ifres15663 = ifres15667


} else {
ifres15663 = False


}

ifres15661 = ifres15663


} else {
ifres15661 = False


}

W458414501 := ifres15661
_ = W458414501

tmp16286 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W458414501, False)
}
__typedArg0 := W458414501
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16286 {
tmp15750 := Call(__e, PrimFunc(symshen_4unlocked_2), V4554)


var ifres15684 Obj

if True == tmp15750 {
tmp15685 := Call(__e, PrimFunc(symshen_4lazyderef), V4550, V4553)


W459214509 := tmp15685
_ = W459214509

tmp15749 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W459214509)
}
__typedArg0 := W459214509
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15686 Obj

if True == tmp15749 {
tmp15687 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W459214509)
}
__typedArg0 := W459214509
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15688 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15687, V4553)


W459314510 := tmp15688
_ = W459314510

tmp15748 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W459314510, symcons)
}
__typedArg0 := W459314510
__typedArg1 := symcons
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15689 Obj

if True == tmp15748 {
tmp15690 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W459214509)
}
__typedArg0 := W459214509
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15691 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15690, V4553)


W459414511 := tmp15691
_ = W459414511

tmp15747 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W459414511)
}
__typedArg0 := W459414511
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15692 Obj

if True == tmp15747 {
tmp15693 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W459414511)
}
__typedArg0 := W459414511
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

W459514512 := tmp15693
_ = W459514512

tmp15694 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W459414511)
}
__typedArg0 := W459414511
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15695 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15694, V4553)


W459614513 := tmp15695
_ = W459614513

tmp15746 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W459614513)
}
__typedArg0 := W459614513
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15696 Obj

if True == tmp15746 {
tmp15697 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W459614513)
}
__typedArg0 := W459614513
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

W459714514 := tmp15697
_ = W459714514

tmp15698 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W459614513)
}
__typedArg0 := W459614513
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15699 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15698, V4553)


W459814515 := tmp15699
_ = W459814515

tmp15745 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W459814515, Nil)
}
__typedArg0 := W459814515
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15700 Obj

if True == tmp15745 {
tmp15701 := Call(__e, PrimFunc(symshen_4lazyderef), V4551, V4553)


W459914516 := tmp15701
_ = W459914516

tmp15702 := MakeNative(func(__e *ControlFlow) {
Z4601 := __e.Get(1)
_ = Z4601
tmp15703 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15703

tmp15704 := MakeNative(func(__e *ControlFlow) {
tmp15705 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Z4601, Nil)
}
__typedArg0 := Z4601
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15706 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp15705)
}
__typedArg0 := symlist
__typedArg1 := tmp15705
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4system_1S_1h), W459714514, tmp15706, V4552, V4553, V4554, W455714474, V4556)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), W459514512, Z4601, V4552, V4553, V4554, W455714474, tmp15704)
return


}, 1)

W460014517 := tmp15702
_ = W460014517

tmp15744 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W459914516)
}
__typedArg0 := W459914516
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15707 Obj

if True == tmp15744 {
tmp15708 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W459914516)
}
__typedArg0 := W459914516
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15709 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15708, V4553)


W460214518 := tmp15709
_ = W460214518

tmp15710 := MakeNative(func(__e *ControlFlow) {
tmp15711 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W459914516)
}
__typedArg0 := W459914516
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15712 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15711, V4553)


W460414520 := tmp15712
_ = W460414520

tmp15713 := MakeNative(func(__e *ControlFlow) {
Z4606 := __e.Get(1)
_ = Z4606
__e.TailApply(W460014517, Z4606)
return
}, 1)

W460514521 := tmp15713
_ = W460514521

tmp15729 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W460414520)
}
__typedArg0 := W460414520
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15729 {
tmp15714 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W460414520)
}
__typedArg0 := W460414520
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

W460714522 := tmp15714
_ = W460714522

tmp15715 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W460414520)
}
__typedArg0 := W460414520
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15716 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15715, V4553)


W460814523 := tmp15716
_ = W460814523

tmp15717 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W460514521, W460714522)
return
}, 0)

W460914524 := tmp15717
_ = W460914524

tmp15721 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W460814523, Nil)
}
__typedArg0 := W460814523
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15721 {
__e.TailApply(PrimFunc(symthaw), W460914524)
return
} else {
tmp15719 := Call(__e, PrimFunc(symshen_4pvar_2), W460814523)


if True == tmp15719 {
__e.TailApply(PrimFunc(symshen_4bind_b), W460814523, Nil, V4553, W460914524)
return
} else {
__e.Return(False)
return
}


}


} else {
tmp15727 := Call(__e, PrimFunc(symshen_4pvar_2), W460414520)


if True == tmp15727 {
tmp15722 := Call(__e, PrimFunc(symshen_4newpv), V4553)


W461014525 := tmp15722
_ = W461014525

tmp15723 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W461014525, Nil)
}
__typedArg0 := W461014525
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15724 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W460514521, W461014525)
return
}, 0)

tmp15725 := Call(__e, PrimFunc(symshen_4bind_b), W460414520, tmp15723, V4553, tmp15724)


__e.TailApply(PrimFunc(symshen_4gc), V4553, tmp15725)
return


} else {
__e.Return(False)
return
}


}


}, 0)

W460314519 := tmp15710
_ = W460314519

tmp15735 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W460214518, symlist)
}
__typedArg0 := W460214518
__typedArg1 := symlist
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15730 Obj

if True == tmp15735 {
tmp15731 := Call(__e, PrimFunc(symthaw), W460314519)


ifres15730 = tmp15731


} else {
tmp15734 := Call(__e, PrimFunc(symshen_4pvar_2), W460214518)


var ifres15732 Obj

if True == tmp15734 {
tmp15733 := Call(__e, PrimFunc(symshen_4bind_b), W460214518, symlist, V4553, W460314519)


ifres15732 = tmp15733


} else {
ifres15732 = False


}

ifres15730 = ifres15732


}

ifres15707 = ifres15730


} else {
tmp15743 := Call(__e, PrimFunc(symshen_4pvar_2), W459914516)


var ifres15736 Obj

if True == tmp15743 {
tmp15737 := Call(__e, PrimFunc(symshen_4newpv), V4553)


W461114526 := tmp15737
_ = W461114526

tmp15738 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W461114526, Nil)
}
__typedArg0 := W461114526
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15739 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp15738)
}
__typedArg0 := symlist
__typedArg1 := tmp15738
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15740 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W460014517, W461114526)
return
}, 0)

tmp15741 := Call(__e, PrimFunc(symshen_4bind_b), W459914516, tmp15739, V4553, tmp15740)


tmp15742 := Call(__e, PrimFunc(symshen_4gc), V4553, tmp15741)


ifres15736 = tmp15742


} else {
ifres15736 = False


}

ifres15707 = ifres15736


}

ifres15700 = ifres15707


} else {
ifres15700 = False


}

ifres15696 = ifres15700


} else {
ifres15696 = False


}

ifres15692 = ifres15696


} else {
ifres15692 = False


}

ifres15689 = ifres15692


} else {
ifres15689 = False


}

ifres15686 = ifres15689


} else {
ifres15686 = False


}

ifres15684 = ifres15686


} else {
ifres15684 = False


}

W459114508 := ifres15684
_ = W459114508

tmp16284 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W459114508, False)
}
__typedArg0 := W459114508
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16284 {
tmp15834 := Call(__e, PrimFunc(symshen_4unlocked_2), V4554)


var ifres15751 Obj

if True == tmp15834 {
tmp15752 := Call(__e, PrimFunc(symshen_4lazyderef), V4550, V4553)


W461314528 := tmp15752
_ = W461314528

tmp15833 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W461314528)
}
__typedArg0 := W461314528
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15753 Obj

if True == tmp15833 {
tmp15754 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W461314528)
}
__typedArg0 := W461314528
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15755 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15754, V4553)


W461414529 := tmp15755
_ = W461414529

tmp15832 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W461414529, sym_8p)
}
__typedArg0 := W461414529
__typedArg1 := sym_8p
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15756 Obj

if True == tmp15832 {
tmp15757 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W461314528)
}
__typedArg0 := W461314528
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15758 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15757, V4553)


W461514530 := tmp15758
_ = W461514530

tmp15831 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W461514530)
}
__typedArg0 := W461514530
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15759 Obj

if True == tmp15831 {
tmp15760 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W461514530)
}
__typedArg0 := W461514530
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

W461614531 := tmp15760
_ = W461614531

tmp15761 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W461514530)
}
__typedArg0 := W461514530
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15762 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15761, V4553)


W461714532 := tmp15762
_ = W461714532

tmp15830 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W461714532)
}
__typedArg0 := W461714532
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15763 Obj

if True == tmp15830 {
tmp15764 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W461714532)
}
__typedArg0 := W461714532
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

W461814533 := tmp15764
_ = W461814533

tmp15765 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W461714532)
}
__typedArg0 := W461714532
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15766 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15765, V4553)


W461914534 := tmp15766
_ = W461914534

tmp15829 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W461914534, Nil)
}
__typedArg0 := W461914534
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15767 Obj

if True == tmp15829 {
tmp15768 := Call(__e, PrimFunc(symshen_4lazyderef), V4551, V4553)


W462014535 := tmp15768
_ = W462014535

tmp15769 := MakeNative(func(__e *ControlFlow) {
Z4622 := __e.Get(1)
_ = Z4622
__e.Return(MakeNative(func(__e *ControlFlow) {
Z4623 := __e.Get(1)
_ = Z4623
tmp15770 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15770

tmp15771 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), W461814533, Z4623, V4552, V4553, V4554, W455714474, V4556)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), W461614531, Z4622, V4552, V4553, V4554, W455714474, tmp15771)
return


}, 1))
return
}, 1)

W462114536 := tmp15769
_ = W462114536

tmp15828 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W462014535)
}
__typedArg0 := W462014535
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15772 Obj

if True == tmp15828 {
tmp15773 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W462014535)
}
__typedArg0 := W462014535
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

W462414537 := tmp15773
_ = W462414537

tmp15774 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W462014535)
}
__typedArg0 := W462014535
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15775 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15774, V4553)


W462514538 := tmp15775
_ = W462514538

tmp15776 := MakeNative(func(__e *ControlFlow) {
Z4627 := __e.Get(1)
_ = Z4627
tmp15777 := Call(__e, W462114536, W462414537)


__e.TailApply(tmp15777, Z4627)
return


}, 1)

W462614539 := tmp15776
_ = W462614539

tmp15815 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W462514538)
}
__typedArg0 := W462514538
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15778 Obj

if True == tmp15815 {
tmp15779 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W462514538)
}
__typedArg0 := W462514538
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15780 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15779, V4553)


W462814540 := tmp15780
_ = W462814540

tmp15781 := MakeNative(func(__e *ControlFlow) {
tmp15782 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W462514538)
}
__typedArg0 := W462514538
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15783 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15782, V4553)


W463014542 := tmp15783
_ = W463014542

tmp15784 := MakeNative(func(__e *ControlFlow) {
Z4632 := __e.Get(1)
_ = Z4632
__e.TailApply(W462614539, Z4632)
return
}, 1)

W463114543 := tmp15784
_ = W463114543

tmp15800 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W463014542)
}
__typedArg0 := W463014542
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15800 {
tmp15785 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W463014542)
}
__typedArg0 := W463014542
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

W463314544 := tmp15785
_ = W463314544

tmp15786 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W463014542)
}
__typedArg0 := W463014542
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15787 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15786, V4553)


W463414545 := tmp15787
_ = W463414545

tmp15788 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W463114543, W463314544)
return
}, 0)

W463514546 := tmp15788
_ = W463514546

tmp15792 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W463414545, Nil)
}
__typedArg0 := W463414545
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15792 {
__e.TailApply(PrimFunc(symthaw), W463514546)
return
} else {
tmp15790 := Call(__e, PrimFunc(symshen_4pvar_2), W463414545)


if True == tmp15790 {
__e.TailApply(PrimFunc(symshen_4bind_b), W463414545, Nil, V4553, W463514546)
return
} else {
__e.Return(False)
return
}


}


} else {
tmp15798 := Call(__e, PrimFunc(symshen_4pvar_2), W463014542)


if True == tmp15798 {
tmp15793 := Call(__e, PrimFunc(symshen_4newpv), V4553)


W463614547 := tmp15793
_ = W463614547

tmp15794 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W463614547, Nil)
}
__typedArg0 := W463614547
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15795 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W463114543, W463614547)
return
}, 0)

tmp15796 := Call(__e, PrimFunc(symshen_4bind_b), W463014542, tmp15794, V4553, tmp15795)


__e.TailApply(PrimFunc(symshen_4gc), V4553, tmp15796)
return


} else {
__e.Return(False)
return
}


}


}, 0)

W462914541 := tmp15781
_ = W462914541

tmp15806 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W462814540, sym_d)
}
__typedArg0 := W462814540
__typedArg1 := sym_d
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15801 Obj

if True == tmp15806 {
tmp15802 := Call(__e, PrimFunc(symthaw), W462914541)


ifres15801 = tmp15802


} else {
tmp15805 := Call(__e, PrimFunc(symshen_4pvar_2), W462814540)


var ifres15803 Obj

if True == tmp15805 {
tmp15804 := Call(__e, PrimFunc(symshen_4bind_b), W462814540, sym_d, V4553, W462914541)


ifres15803 = tmp15804


} else {
ifres15803 = False


}

ifres15801 = ifres15803


}

ifres15778 = ifres15801


} else {
tmp15814 := Call(__e, PrimFunc(symshen_4pvar_2), W462514538)


var ifres15807 Obj

if True == tmp15814 {
tmp15808 := Call(__e, PrimFunc(symshen_4newpv), V4553)


W463714548 := tmp15808
_ = W463714548

tmp15809 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W463714548, Nil)
}
__typedArg0 := W463714548
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15810 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_d, tmp15809)
}
__typedArg0 := sym_d
__typedArg1 := tmp15809
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15811 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W462614539, W463714548)
return
}, 0)

tmp15812 := Call(__e, PrimFunc(symshen_4bind_b), W462514538, tmp15810, V4553, tmp15811)


tmp15813 := Call(__e, PrimFunc(symshen_4gc), V4553, tmp15812)


ifres15807 = tmp15813


} else {
ifres15807 = False


}

ifres15778 = ifres15807


}

ifres15772 = ifres15778


} else {
tmp15827 := Call(__e, PrimFunc(symshen_4pvar_2), W462014535)


var ifres15816 Obj

if True == tmp15827 {
tmp15817 := Call(__e, PrimFunc(symshen_4newpv), V4553)


W463814549 := tmp15817
_ = W463814549

tmp15818 := Call(__e, PrimFunc(symshen_4newpv), V4553)


W463914550 := tmp15818
_ = W463914550

tmp15819 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W463914550, Nil)
}
__typedArg0 := W463914550
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15820 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_d, tmp15819)
}
__typedArg0 := sym_d
__typedArg1 := tmp15819
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15821 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W463814549, tmp15820)
}
__typedArg0 := W463814549
__typedArg1 := tmp15820
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15822 := MakeNative(func(__e *ControlFlow) {
tmp15823 := Call(__e, W462114536, W463814549)


__e.TailApply(tmp15823, W463914550)
return


}, 0)

tmp15824 := Call(__e, PrimFunc(symshen_4bind_b), W462014535, tmp15821, V4553, tmp15822)


tmp15825 := Call(__e, PrimFunc(symshen_4gc), V4553, tmp15824)


tmp15826 := Call(__e, PrimFunc(symshen_4gc), V4553, tmp15825)


ifres15816 = tmp15826


} else {
ifres15816 = False


}

ifres15772 = ifres15816


}

ifres15767 = ifres15772


} else {
ifres15767 = False


}

ifres15763 = ifres15767


} else {
ifres15763 = False


}

ifres15759 = ifres15763


} else {
ifres15759 = False


}

ifres15756 = ifres15759


} else {
ifres15756 = False


}

ifres15753 = ifres15756


} else {
ifres15753 = False


}

ifres15751 = ifres15753


} else {
ifres15751 = False


}

W461214527 := ifres15751
_ = W461214527

tmp16282 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W461214527, False)
}
__typedArg0 := W461214527
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16282 {
tmp15901 := Call(__e, PrimFunc(symshen_4unlocked_2), V4554)


var ifres15835 Obj

if True == tmp15901 {
tmp15836 := Call(__e, PrimFunc(symshen_4lazyderef), V4550, V4553)


W464114552 := tmp15836
_ = W464114552

tmp15900 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W464114552)
}
__typedArg0 := W464114552
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15837 Obj

if True == tmp15900 {
tmp15838 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W464114552)
}
__typedArg0 := W464114552
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15839 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15838, V4553)


W464214553 := tmp15839
_ = W464214553

tmp15899 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W464214553, sym_8v)
}
__typedArg0 := W464214553
__typedArg1 := sym_8v
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15840 Obj

if True == tmp15899 {
tmp15841 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W464114552)
}
__typedArg0 := W464114552
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15842 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15841, V4553)


W464314554 := tmp15842
_ = W464314554

tmp15898 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W464314554)
}
__typedArg0 := W464314554
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15843 Obj

if True == tmp15898 {
tmp15844 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W464314554)
}
__typedArg0 := W464314554
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

W464414555 := tmp15844
_ = W464414555

tmp15845 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W464314554)
}
__typedArg0 := W464314554
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15846 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15845, V4553)


W464514556 := tmp15846
_ = W464514556

tmp15897 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W464514556)
}
__typedArg0 := W464514556
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15847 Obj

if True == tmp15897 {
tmp15848 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W464514556)
}
__typedArg0 := W464514556
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

W464614557 := tmp15848
_ = W464614557

tmp15849 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W464514556)
}
__typedArg0 := W464514556
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15850 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15849, V4553)


W464714558 := tmp15850
_ = W464714558

tmp15896 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W464714558, Nil)
}
__typedArg0 := W464714558
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15851 Obj

if True == tmp15896 {
tmp15852 := Call(__e, PrimFunc(symshen_4lazyderef), V4551, V4553)


W464814559 := tmp15852
_ = W464814559

tmp15853 := MakeNative(func(__e *ControlFlow) {
Z4650 := __e.Get(1)
_ = Z4650
tmp15854 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15854

tmp15855 := MakeNative(func(__e *ControlFlow) {
tmp15856 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Z4650, Nil)
}
__typedArg0 := Z4650
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15857 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp15856)
}
__typedArg0 := symvector
__typedArg1 := tmp15856
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4system_1S_1h), W464614557, tmp15857, V4552, V4553, V4554, W455714474, V4556)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), W464414555, Z4650, V4552, V4553, V4554, W455714474, tmp15855)
return


}, 1)

W464914560 := tmp15853
_ = W464914560

tmp15895 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W464814559)
}
__typedArg0 := W464814559
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15858 Obj

if True == tmp15895 {
tmp15859 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W464814559)
}
__typedArg0 := W464814559
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15860 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15859, V4553)


W465114561 := tmp15860
_ = W465114561

tmp15861 := MakeNative(func(__e *ControlFlow) {
tmp15862 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W464814559)
}
__typedArg0 := W464814559
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15863 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15862, V4553)


W465314563 := tmp15863
_ = W465314563

tmp15864 := MakeNative(func(__e *ControlFlow) {
Z4655 := __e.Get(1)
_ = Z4655
__e.TailApply(W464914560, Z4655)
return
}, 1)

W465414564 := tmp15864
_ = W465414564

tmp15880 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W465314563)
}
__typedArg0 := W465314563
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15880 {
tmp15865 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W465314563)
}
__typedArg0 := W465314563
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

W465614565 := tmp15865
_ = W465614565

tmp15866 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W465314563)
}
__typedArg0 := W465314563
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15867 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15866, V4553)


W465714566 := tmp15867
_ = W465714566

tmp15868 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W465414564, W465614565)
return
}, 0)

W465814567 := tmp15868
_ = W465814567

tmp15872 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W465714566, Nil)
}
__typedArg0 := W465714566
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15872 {
__e.TailApply(PrimFunc(symthaw), W465814567)
return
} else {
tmp15870 := Call(__e, PrimFunc(symshen_4pvar_2), W465714566)


if True == tmp15870 {
__e.TailApply(PrimFunc(symshen_4bind_b), W465714566, Nil, V4553, W465814567)
return
} else {
__e.Return(False)
return
}


}


} else {
tmp15878 := Call(__e, PrimFunc(symshen_4pvar_2), W465314563)


if True == tmp15878 {
tmp15873 := Call(__e, PrimFunc(symshen_4newpv), V4553)


W465914568 := tmp15873
_ = W465914568

tmp15874 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W465914568, Nil)
}
__typedArg0 := W465914568
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15875 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W465414564, W465914568)
return
}, 0)

tmp15876 := Call(__e, PrimFunc(symshen_4bind_b), W465314563, tmp15874, V4553, tmp15875)


__e.TailApply(PrimFunc(symshen_4gc), V4553, tmp15876)
return


} else {
__e.Return(False)
return
}


}


}, 0)

W465214562 := tmp15861
_ = W465214562

tmp15886 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W465114561, symvector)
}
__typedArg0 := W465114561
__typedArg1 := symvector
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15881 Obj

if True == tmp15886 {
tmp15882 := Call(__e, PrimFunc(symthaw), W465214562)


ifres15881 = tmp15882


} else {
tmp15885 := Call(__e, PrimFunc(symshen_4pvar_2), W465114561)


var ifres15883 Obj

if True == tmp15885 {
tmp15884 := Call(__e, PrimFunc(symshen_4bind_b), W465114561, symvector, V4553, W465214562)


ifres15883 = tmp15884


} else {
ifres15883 = False


}

ifres15881 = ifres15883


}

ifres15858 = ifres15881


} else {
tmp15894 := Call(__e, PrimFunc(symshen_4pvar_2), W464814559)


var ifres15887 Obj

if True == tmp15894 {
tmp15888 := Call(__e, PrimFunc(symshen_4newpv), V4553)


W466014569 := tmp15888
_ = W466014569

tmp15889 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W466014569, Nil)
}
__typedArg0 := W466014569
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15890 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp15889)
}
__typedArg0 := symvector
__typedArg1 := tmp15889
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15891 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W464914560, W466014569)
return
}, 0)

tmp15892 := Call(__e, PrimFunc(symshen_4bind_b), W464814559, tmp15890, V4553, tmp15891)


tmp15893 := Call(__e, PrimFunc(symshen_4gc), V4553, tmp15892)


ifres15887 = tmp15893


} else {
ifres15887 = False


}

ifres15858 = ifres15887


}

ifres15851 = ifres15858


} else {
ifres15851 = False


}

ifres15847 = ifres15851


} else {
ifres15847 = False


}

ifres15843 = ifres15847


} else {
ifres15843 = False


}

ifres15840 = ifres15843


} else {
ifres15840 = False


}

ifres15837 = ifres15840


} else {
ifres15837 = False


}

ifres15835 = ifres15837


} else {
ifres15835 = False


}

W464014551 := ifres15835
_ = W464014551

tmp16280 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W464014551, False)
}
__typedArg0 := W464014551
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16280 {
tmp15934 := Call(__e, PrimFunc(symshen_4unlocked_2), V4554)


var ifres15902 Obj

if True == tmp15934 {
tmp15903 := Call(__e, PrimFunc(symshen_4lazyderef), V4550, V4553)


W466214571 := tmp15903
_ = W466214571

tmp15933 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W466214571)
}
__typedArg0 := W466214571
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15904 Obj

if True == tmp15933 {
tmp15905 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W466214571)
}
__typedArg0 := W466214571
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15906 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15905, V4553)


W466314572 := tmp15906
_ = W466314572

tmp15932 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W466314572, sym_8s)
}
__typedArg0 := W466314572
__typedArg1 := sym_8s
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15907 Obj

if True == tmp15932 {
tmp15908 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W466214571)
}
__typedArg0 := W466214571
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15909 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15908, V4553)


W466414573 := tmp15909
_ = W466414573

tmp15931 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W466414573)
}
__typedArg0 := W466414573
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15910 Obj

if True == tmp15931 {
tmp15911 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W466414573)
}
__typedArg0 := W466414573
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

W466514574 := tmp15911
_ = W466514574

tmp15912 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W466414573)
}
__typedArg0 := W466414573
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15913 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15912, V4553)


W466614575 := tmp15913
_ = W466614575

tmp15930 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W466614575)
}
__typedArg0 := W466614575
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15914 Obj

if True == tmp15930 {
tmp15915 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W466614575)
}
__typedArg0 := W466614575
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

W466714576 := tmp15915
_ = W466714576

tmp15916 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W466614575)
}
__typedArg0 := W466614575
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15917 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15916, V4553)


W466814577 := tmp15917
_ = W466814577

tmp15929 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W466814577, Nil)
}
__typedArg0 := W466814577
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15918 Obj

if True == tmp15929 {
tmp15919 := Call(__e, PrimFunc(symshen_4lazyderef), V4551, V4553)


W466914578 := tmp15919
_ = W466914578

tmp15920 := MakeNative(func(__e *ControlFlow) {
tmp15921 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15921

tmp15922 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), W466714576, symstring, V4552, V4553, V4554, W455714474, V4556)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), W466514574, symstring, V4552, V4553, V4554, W455714474, tmp15922)
return


}, 0)

W467014579 := tmp15920
_ = W467014579

tmp15928 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W466914578, symstring)
}
__typedArg0 := W466914578
__typedArg1 := symstring
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15923 Obj

if True == tmp15928 {
tmp15924 := Call(__e, PrimFunc(symthaw), W467014579)


ifres15923 = tmp15924


} else {
tmp15927 := Call(__e, PrimFunc(symshen_4pvar_2), W466914578)


var ifres15925 Obj

if True == tmp15927 {
tmp15926 := Call(__e, PrimFunc(symshen_4bind_b), W466914578, symstring, V4553, W467014579)


ifres15925 = tmp15926


} else {
ifres15925 = False


}

ifres15923 = ifres15925


}

ifres15918 = ifres15923


} else {
ifres15918 = False


}

ifres15914 = ifres15918


} else {
ifres15914 = False


}

ifres15910 = ifres15914


} else {
ifres15910 = False


}

ifres15907 = ifres15910


} else {
ifres15907 = False


}

ifres15904 = ifres15907


} else {
ifres15904 = False


}

ifres15902 = ifres15904


} else {
ifres15902 = False


}

W466114570 := ifres15902
_ = W466114570

tmp16278 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W466114570, False)
}
__typedArg0 := W466114570
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16278 {
tmp16034 := Call(__e, PrimFunc(symshen_4unlocked_2), V4554)


var ifres15935 Obj

if True == tmp16034 {
tmp15936 := Call(__e, PrimFunc(symshen_4lazyderef), V4550, V4553)


W467214581 := tmp15936
_ = W467214581

tmp16033 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W467214581)
}
__typedArg0 := W467214581
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15937 Obj

if True == tmp16033 {
tmp15938 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W467214581)
}
__typedArg0 := W467214581
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15939 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15938, V4553)


W467314582 := tmp15939
_ = W467314582

tmp16032 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W467314582, symlambda)
}
__typedArg0 := W467314582
__typedArg1 := symlambda
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15940 Obj

if True == tmp16032 {
tmp15941 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W467214581)
}
__typedArg0 := W467214581
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15942 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15941, V4553)


W467414583 := tmp15942
_ = W467414583

tmp16031 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W467414583)
}
__typedArg0 := W467414583
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15943 Obj

if True == tmp16031 {
tmp15944 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W467414583)
}
__typedArg0 := W467414583
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

W467514584 := tmp15944
_ = W467514584

tmp15945 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W467414583)
}
__typedArg0 := W467414583
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15946 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15945, V4553)


W467614585 := tmp15946
_ = W467614585

tmp16030 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W467614585)
}
__typedArg0 := W467614585
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15947 Obj

if True == tmp16030 {
tmp15948 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W467614585)
}
__typedArg0 := W467614585
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

W467714586 := tmp15948
_ = W467714586

tmp15949 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W467614585)
}
__typedArg0 := W467614585
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15950 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15949, V4553)


W467814587 := tmp15950
_ = W467814587

tmp16029 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W467814587, Nil)
}
__typedArg0 := W467814587
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15951 Obj

if True == tmp16029 {
tmp15952 := Call(__e, PrimFunc(symshen_4lazyderef), V4551, V4553)


W467914588 := tmp15952
_ = W467914588

tmp15953 := MakeNative(func(__e *ControlFlow) {
Z4681 := __e.Get(1)
_ = Z4681
__e.Return(MakeNative(func(__e *ControlFlow) {
Z4682 := __e.Get(1)
_ = Z4682
tmp15954 := Call(__e, PrimFunc(symshen_4newpv), V4553)


W468314590 := tmp15954
_ = W468314590

tmp15955 := Call(__e, PrimFunc(symshen_4newpv), V4553)


W468414591 := tmp15955
_ = W468414591

tmp15956 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15956

tmp15957 := Call(__e, PrimFunc(symshen_4lazyderef), W467514584, V4553)


tmp15958 := Call(__e, PrimFunc(symshen_4freshterm), tmp15957)


tmp15959 := MakeNative(func(__e *ControlFlow) {
tmp15960 := Call(__e, PrimFunc(symshen_4lazyderef), W467514584, V4553)


tmp15961 := Call(__e, PrimFunc(symshen_4deref), W468414591, V4553)


tmp15962 := Call(__e, PrimFunc(symshen_4deref), W467714586, V4553)


tmp15963 := Call(__e, PrimFunc(symshen_4beta), tmp15960, tmp15961, tmp15962)


tmp15964 := MakeNative(func(__e *ControlFlow) {
tmp15965 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp15966 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Z4681, Nil)
}
__typedArg0 := Z4681
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15967 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp15965, tmp15966)
}
__typedArg0 := tmp15965
__typedArg1 := tmp15966
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15968 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W468414591, tmp15967)
}
__typedArg0 := W468414591
__typedArg1 := tmp15967
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15969 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp15968, V4552)
}
__typedArg0 := tmp15968
__typedArg1 := V4552
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4system_1S_1h), W468314590, Z4682, tmp15969, V4553, V4554, W455714474, V4556)
return


}, 0)

__e.TailApply(PrimFunc(symbind), W468314590, tmp15963, V4553, V4554, W455714474, tmp15964)
return


}, 0)

tmp15970 := Call(__e, PrimFunc(symbind), W468414591, tmp15958, V4553, V4554, W455714474, tmp15959)


tmp15971 := Call(__e, PrimFunc(symshen_4gc), V4553, tmp15970)


__e.TailApply(PrimFunc(symshen_4gc), V4553, tmp15971)
return


}, 1))
return
}, 1)

W468014589 := tmp15953
_ = W468014589

tmp16028 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W467914588)
}
__typedArg0 := W467914588
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15972 Obj

if True == tmp16028 {
tmp15973 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W467914588)
}
__typedArg0 := W467914588
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

W468514592 := tmp15973
_ = W468514592

tmp15974 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W467914588)
}
__typedArg0 := W467914588
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15975 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15974, V4553)


W468614593 := tmp15975
_ = W468614593

tmp15976 := MakeNative(func(__e *ControlFlow) {
Z4688 := __e.Get(1)
_ = Z4688
tmp15977 := Call(__e, W468014589, W468514592)


__e.TailApply(tmp15977, Z4688)
return


}, 1)

W468714594 := tmp15976
_ = W468714594

tmp16015 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W468614593)
}
__typedArg0 := W468614593
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15978 Obj

if True == tmp16015 {
tmp15979 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W468614593)
}
__typedArg0 := W468614593
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15980 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15979, V4553)


W468914595 := tmp15980
_ = W468914595

tmp15981 := MakeNative(func(__e *ControlFlow) {
tmp15982 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W468614593)
}
__typedArg0 := W468614593
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15983 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15982, V4553)


W469114597 := tmp15983
_ = W469114597

tmp15984 := MakeNative(func(__e *ControlFlow) {
Z4693 := __e.Get(1)
_ = Z4693
__e.TailApply(W468714594, Z4693)
return
}, 1)

W469214598 := tmp15984
_ = W469214598

tmp16000 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W469114597)
}
__typedArg0 := W469114597
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp16000 {
tmp15985 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W469114597)
}
__typedArg0 := W469114597
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

W469414599 := tmp15985
_ = W469414599

tmp15986 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W469114597)
}
__typedArg0 := W469114597
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15987 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15986, V4553)


W469514600 := tmp15987
_ = W469514600

tmp15988 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W469214598, W469414599)
return
}, 0)

W469614601 := tmp15988
_ = W469614601

tmp15992 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W469514600, Nil)
}
__typedArg0 := W469514600
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15992 {
__e.TailApply(PrimFunc(symthaw), W469614601)
return
} else {
tmp15990 := Call(__e, PrimFunc(symshen_4pvar_2), W469514600)


if True == tmp15990 {
__e.TailApply(PrimFunc(symshen_4bind_b), W469514600, Nil, V4553, W469614601)
return
} else {
__e.Return(False)
return
}


}


} else {
tmp15998 := Call(__e, PrimFunc(symshen_4pvar_2), W469114597)


if True == tmp15998 {
tmp15993 := Call(__e, PrimFunc(symshen_4newpv), V4553)


W469714602 := tmp15993
_ = W469714602

tmp15994 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W469714602, Nil)
}
__typedArg0 := W469714602
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15995 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W469214598, W469714602)
return
}, 0)

tmp15996 := Call(__e, PrimFunc(symshen_4bind_b), W469114597, tmp15994, V4553, tmp15995)


__e.TailApply(PrimFunc(symshen_4gc), V4553, tmp15996)
return


} else {
__e.Return(False)
return
}


}


}, 0)

W469014596 := tmp15981
_ = W469014596

tmp16006 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W468914595, sym_1_1_6)
}
__typedArg0 := W468914595
__typedArg1 := sym_1_1_6
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16001 Obj

if True == tmp16006 {
tmp16002 := Call(__e, PrimFunc(symthaw), W469014596)


ifres16001 = tmp16002


} else {
tmp16005 := Call(__e, PrimFunc(symshen_4pvar_2), W468914595)


var ifres16003 Obj

if True == tmp16005 {
tmp16004 := Call(__e, PrimFunc(symshen_4bind_b), W468914595, sym_1_1_6, V4553, W469014596)


ifres16003 = tmp16004


} else {
ifres16003 = False


}

ifres16001 = ifres16003


}

ifres15978 = ifres16001


} else {
tmp16014 := Call(__e, PrimFunc(symshen_4pvar_2), W468614593)


var ifres16007 Obj

if True == tmp16014 {
tmp16008 := Call(__e, PrimFunc(symshen_4newpv), V4553)


W469814603 := tmp16008
_ = W469814603

tmp16009 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W469814603, Nil)
}
__typedArg0 := W469814603
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16010 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp16009)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp16009
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16011 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W468714594, W469814603)
return
}, 0)

tmp16012 := Call(__e, PrimFunc(symshen_4bind_b), W468614593, tmp16010, V4553, tmp16011)


tmp16013 := Call(__e, PrimFunc(symshen_4gc), V4553, tmp16012)


ifres16007 = tmp16013


} else {
ifres16007 = False


}

ifres15978 = ifres16007


}

ifres15972 = ifres15978


} else {
tmp16027 := Call(__e, PrimFunc(symshen_4pvar_2), W467914588)


var ifres16016 Obj

if True == tmp16027 {
tmp16017 := Call(__e, PrimFunc(symshen_4newpv), V4553)


W469914604 := tmp16017
_ = W469914604

tmp16018 := Call(__e, PrimFunc(symshen_4newpv), V4553)


W470014605 := tmp16018
_ = W470014605

tmp16019 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W470014605, Nil)
}
__typedArg0 := W470014605
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16020 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp16019)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp16019
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16021 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W469914604, tmp16020)
}
__typedArg0 := W469914604
__typedArg1 := tmp16020
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16022 := MakeNative(func(__e *ControlFlow) {
tmp16023 := Call(__e, W468014589, W469914604)


__e.TailApply(tmp16023, W470014605)
return


}, 0)

tmp16024 := Call(__e, PrimFunc(symshen_4bind_b), W467914588, tmp16021, V4553, tmp16022)


tmp16025 := Call(__e, PrimFunc(symshen_4gc), V4553, tmp16024)


tmp16026 := Call(__e, PrimFunc(symshen_4gc), V4553, tmp16025)


ifres16016 = tmp16026


} else {
ifres16016 = False


}

ifres15972 = ifres16016


}

ifres15951 = ifres15972


} else {
ifres15951 = False


}

ifres15947 = ifres15951


} else {
ifres15947 = False


}

ifres15943 = ifres15947


} else {
ifres15943 = False


}

ifres15940 = ifres15943


} else {
ifres15940 = False


}

ifres15937 = ifres15940


} else {
ifres15937 = False


}

ifres15935 = ifres15937


} else {
ifres15935 = False


}

W467114580 := ifres15935
_ = W467114580

tmp16276 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W467114580, False)
}
__typedArg0 := W467114580
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16276 {
tmp16084 := Call(__e, PrimFunc(symshen_4unlocked_2), V4554)


var ifres16035 Obj

if True == tmp16084 {
tmp16036 := Call(__e, PrimFunc(symshen_4lazyderef), V4550, V4553)


W470214607 := tmp16036
_ = W470214607

tmp16083 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W470214607)
}
__typedArg0 := W470214607
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16037 Obj

if True == tmp16083 {
tmp16038 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W470214607)
}
__typedArg0 := W470214607
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16039 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16038, V4553)


W470314608 := tmp16039
_ = W470314608

tmp16082 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W470314608, symlet)
}
__typedArg0 := W470314608
__typedArg1 := symlet
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16040 Obj

if True == tmp16082 {
tmp16041 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W470214607)
}
__typedArg0 := W470214607
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16042 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16041, V4553)


W470414609 := tmp16042
_ = W470414609

tmp16081 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W470414609)
}
__typedArg0 := W470414609
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16043 Obj

if True == tmp16081 {
tmp16044 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W470414609)
}
__typedArg0 := W470414609
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

W470514610 := tmp16044
_ = W470514610

tmp16045 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W470414609)
}
__typedArg0 := W470414609
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16046 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16045, V4553)


W470614611 := tmp16046
_ = W470614611

tmp16080 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W470614611)
}
__typedArg0 := W470614611
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16047 Obj

if True == tmp16080 {
tmp16048 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W470614611)
}
__typedArg0 := W470614611
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

W470714612 := tmp16048
_ = W470714612

tmp16049 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W470614611)
}
__typedArg0 := W470614611
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16050 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16049, V4553)


W470814613 := tmp16050
_ = W470814613

tmp16079 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W470814613)
}
__typedArg0 := W470814613
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16051 Obj

if True == tmp16079 {
tmp16052 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W470814613)
}
__typedArg0 := W470814613
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

W470914614 := tmp16052
_ = W470914614

tmp16053 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W470814613)
}
__typedArg0 := W470814613
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16054 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16053, V4553)


W471014615 := tmp16054
_ = W471014615

tmp16078 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W471014615, Nil)
}
__typedArg0 := W471014615
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16055 Obj

if True == tmp16078 {
tmp16056 := Call(__e, PrimFunc(symshen_4newpv), V4553)


W471114616 := tmp16056
_ = W471114616

tmp16057 := Call(__e, PrimFunc(symshen_4newpv), V4553)


W471214617 := tmp16057
_ = W471214617

tmp16058 := Call(__e, PrimFunc(symshen_4newpv), V4553)


W471314618 := tmp16058
_ = W471314618

tmp16059 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16059

tmp16060 := MakeNative(func(__e *ControlFlow) {
tmp16061 := Call(__e, PrimFunc(symshen_4lazyderef), W470514610, V4553)


tmp16062 := Call(__e, PrimFunc(symshen_4freshterm), tmp16061)


tmp16063 := MakeNative(func(__e *ControlFlow) {
tmp16064 := Call(__e, PrimFunc(symshen_4lazyderef), W470514610, V4553)


tmp16065 := Call(__e, PrimFunc(symshen_4lazyderef), W471214617, V4553)


tmp16066 := Call(__e, PrimFunc(symshen_4lazyderef), W470914614, V4553)


tmp16067 := Call(__e, PrimFunc(symshen_4beta), tmp16064, tmp16065, tmp16066)


tmp16068 := MakeNative(func(__e *ControlFlow) {
tmp16069 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp16070 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W471314618, Nil)
}
__typedArg0 := W471314618
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16071 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16069, tmp16070)
}
__typedArg0 := tmp16069
__typedArg1 := tmp16070
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16072 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W471214617, tmp16071)
}
__typedArg0 := W471214617
__typedArg1 := tmp16071
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16073 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16072, V4552)
}
__typedArg0 := tmp16072
__typedArg1 := V4552
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4system_1S_1h), W471114616, V4551, tmp16073, V4553, V4554, W455714474, V4556)
return


}, 0)

__e.TailApply(PrimFunc(symbind), W471114616, tmp16067, V4553, V4554, W455714474, tmp16068)
return


}, 0)

__e.TailApply(PrimFunc(symbind), W471214617, tmp16062, V4553, V4554, W455714474, tmp16063)
return


}, 0)

tmp16074 := Call(__e, PrimFunc(symshen_4system_1S_1h), W470714612, W471314618, V4552, V4553, V4554, W455714474, tmp16060)


tmp16075 := Call(__e, PrimFunc(symshen_4gc), V4553, tmp16074)


tmp16076 := Call(__e, PrimFunc(symshen_4gc), V4553, tmp16075)


tmp16077 := Call(__e, PrimFunc(symshen_4gc), V4553, tmp16076)


ifres16055 = tmp16077


} else {
ifres16055 = False


}

ifres16051 = ifres16055


} else {
ifres16051 = False


}

ifres16047 = ifres16051


} else {
ifres16047 = False


}

ifres16043 = ifres16047


} else {
ifres16043 = False


}

ifres16040 = ifres16043


} else {
ifres16040 = False


}

ifres16037 = ifres16040


} else {
ifres16037 = False


}

ifres16035 = ifres16037


} else {
ifres16035 = False


}

W470114606 := ifres16035
_ = W470114606

tmp16274 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W470114606, False)
}
__typedArg0 := W470114606
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16274 {
tmp16154 := Call(__e, PrimFunc(symshen_4unlocked_2), V4554)


var ifres16085 Obj

if True == tmp16154 {
tmp16086 := Call(__e, PrimFunc(symshen_4lazyderef), V4550, V4553)


W471514620 := tmp16086
_ = W471514620

tmp16153 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W471514620)
}
__typedArg0 := W471514620
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16087 Obj

if True == tmp16153 {
tmp16088 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W471514620)
}
__typedArg0 := W471514620
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16089 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16088, V4553)


W471614621 := tmp16089
_ = W471614621

tmp16152 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W471614621, symopen)
}
__typedArg0 := W471614621
__typedArg1 := symopen
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16090 Obj

if True == tmp16152 {
tmp16091 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W471514620)
}
__typedArg0 := W471514620
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16092 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16091, V4553)


W471714622 := tmp16092
_ = W471714622

tmp16151 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W471714622)
}
__typedArg0 := W471714622
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16093 Obj

if True == tmp16151 {
tmp16094 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W471714622)
}
__typedArg0 := W471714622
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

W471814623 := tmp16094
_ = W471814623

tmp16095 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W471714622)
}
__typedArg0 := W471714622
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16096 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16095, V4553)


W471914624 := tmp16096
_ = W471914624

tmp16150 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W471914624)
}
__typedArg0 := W471914624
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16097 Obj

if True == tmp16150 {
tmp16098 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W471914624)
}
__typedArg0 := W471914624
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

W472014625 := tmp16098
_ = W472014625

tmp16099 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W471914624)
}
__typedArg0 := W471914624
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16100 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16099, V4553)


W472114626 := tmp16100
_ = W472114626

tmp16149 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W472114626, Nil)
}
__typedArg0 := W472114626
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16101 Obj

if True == tmp16149 {
tmp16102 := Call(__e, PrimFunc(symshen_4lazyderef), V4551, V4553)


W472214627 := tmp16102
_ = W472214627

tmp16103 := MakeNative(func(__e *ControlFlow) {
Z4724 := __e.Get(1)
_ = Z4724
tmp16104 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16104

tmp16105 := MakeNative(func(__e *ControlFlow) {
tmp16106 := Call(__e, PrimFunc(symshen_4lazyderef), Z4724, V4553)


tmp16107 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symout, Nil)
}
__typedArg0 := symout
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16108 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symin, tmp16107)
}
__typedArg0 := symin
__typedArg1 := tmp16107
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16109 := Call(__e, PrimFunc(symelement_2), tmp16106, tmp16108)


tmp16110 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), W471814623, symstring, V4552, V4553, V4554, W455714474, V4556)
return
}, 0)

__e.TailApply(PrimFunc(symwhen), tmp16109, V4553, V4554, W455714474, tmp16110)
return


}, 0)

__e.TailApply(PrimFunc(symis_b), W472014625, Z4724, V4553, V4554, W455714474, tmp16105)
return


}, 1)

W472314628 := tmp16103
_ = W472314628

tmp16148 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W472214627)
}
__typedArg0 := W472214627
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16111 Obj

if True == tmp16148 {
tmp16112 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W472214627)
}
__typedArg0 := W472214627
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16113 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16112, V4553)


W472514629 := tmp16113
_ = W472514629

tmp16114 := MakeNative(func(__e *ControlFlow) {
tmp16115 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W472214627)
}
__typedArg0 := W472214627
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16116 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16115, V4553)


W472714631 := tmp16116
_ = W472714631

tmp16117 := MakeNative(func(__e *ControlFlow) {
Z4729 := __e.Get(1)
_ = Z4729
__e.TailApply(W472314628, Z4729)
return
}, 1)

W472814632 := tmp16117
_ = W472814632

tmp16133 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W472714631)
}
__typedArg0 := W472714631
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp16133 {
tmp16118 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W472714631)
}
__typedArg0 := W472714631
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

W473014633 := tmp16118
_ = W473014633

tmp16119 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W472714631)
}
__typedArg0 := W472714631
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16120 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16119, V4553)


W473114634 := tmp16120
_ = W473114634

tmp16121 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W472814632, W473014633)
return
}, 0)

W473214635 := tmp16121
_ = W473214635

tmp16125 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W473114634, Nil)
}
__typedArg0 := W473114634
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16125 {
__e.TailApply(PrimFunc(symthaw), W473214635)
return
} else {
tmp16123 := Call(__e, PrimFunc(symshen_4pvar_2), W473114634)


if True == tmp16123 {
__e.TailApply(PrimFunc(symshen_4bind_b), W473114634, Nil, V4553, W473214635)
return
} else {
__e.Return(False)
return
}


}


} else {
tmp16131 := Call(__e, PrimFunc(symshen_4pvar_2), W472714631)


if True == tmp16131 {
tmp16126 := Call(__e, PrimFunc(symshen_4newpv), V4553)


W473314636 := tmp16126
_ = W473314636

tmp16127 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W473314636, Nil)
}
__typedArg0 := W473314636
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16128 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W472814632, W473314636)
return
}, 0)

tmp16129 := Call(__e, PrimFunc(symshen_4bind_b), W472714631, tmp16127, V4553, tmp16128)


__e.TailApply(PrimFunc(symshen_4gc), V4553, tmp16129)
return


} else {
__e.Return(False)
return
}


}


}, 0)

W472614630 := tmp16114
_ = W472614630

tmp16139 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W472514629, symstream)
}
__typedArg0 := W472514629
__typedArg1 := symstream
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16134 Obj

if True == tmp16139 {
tmp16135 := Call(__e, PrimFunc(symthaw), W472614630)


ifres16134 = tmp16135


} else {
tmp16138 := Call(__e, PrimFunc(symshen_4pvar_2), W472514629)


var ifres16136 Obj

if True == tmp16138 {
tmp16137 := Call(__e, PrimFunc(symshen_4bind_b), W472514629, symstream, V4553, W472614630)


ifres16136 = tmp16137


} else {
ifres16136 = False


}

ifres16134 = ifres16136


}

ifres16111 = ifres16134


} else {
tmp16147 := Call(__e, PrimFunc(symshen_4pvar_2), W472214627)


var ifres16140 Obj

if True == tmp16147 {
tmp16141 := Call(__e, PrimFunc(symshen_4newpv), V4553)


W473414637 := tmp16141
_ = W473414637

tmp16142 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W473414637, Nil)
}
__typedArg0 := W473414637
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16143 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp16142)
}
__typedArg0 := symstream
__typedArg1 := tmp16142
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16144 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W472314628, W473414637)
return
}, 0)

tmp16145 := Call(__e, PrimFunc(symshen_4bind_b), W472214627, tmp16143, V4553, tmp16144)


tmp16146 := Call(__e, PrimFunc(symshen_4gc), V4553, tmp16145)


ifres16140 = tmp16146


} else {
ifres16140 = False


}

ifres16111 = ifres16140


}

ifres16101 = ifres16111


} else {
ifres16101 = False


}

ifres16097 = ifres16101


} else {
ifres16097 = False


}

ifres16093 = ifres16097


} else {
ifres16093 = False


}

ifres16090 = ifres16093


} else {
ifres16090 = False


}

ifres16087 = ifres16090


} else {
ifres16087 = False


}

ifres16085 = ifres16087


} else {
ifres16085 = False


}

W471414619 := ifres16085
_ = W471414619

tmp16272 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W471414619, False)
}
__typedArg0 := W471414619
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16272 {
tmp16183 := Call(__e, PrimFunc(symshen_4unlocked_2), V4554)


var ifres16155 Obj

if True == tmp16183 {
tmp16156 := Call(__e, PrimFunc(symshen_4lazyderef), V4550, V4553)


W473614639 := tmp16156
_ = W473614639

tmp16182 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W473614639)
}
__typedArg0 := W473614639
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16157 Obj

if True == tmp16182 {
tmp16158 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W473614639)
}
__typedArg0 := W473614639
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16159 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16158, V4553)


W473714640 := tmp16159
_ = W473714640

tmp16181 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W473714640, symtype)
}
__typedArg0 := W473714640
__typedArg1 := symtype
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16160 Obj

if True == tmp16181 {
tmp16161 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W473614639)
}
__typedArg0 := W473614639
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16162 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16161, V4553)


W473814641 := tmp16162
_ = W473814641

tmp16180 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W473814641)
}
__typedArg0 := W473814641
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16163 Obj

if True == tmp16180 {
tmp16164 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W473814641)
}
__typedArg0 := W473814641
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

W473914642 := tmp16164
_ = W473914642

tmp16165 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W473814641)
}
__typedArg0 := W473814641
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16166 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16165, V4553)


W474014643 := tmp16166
_ = W474014643

tmp16179 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W474014643)
}
__typedArg0 := W474014643
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16167 Obj

if True == tmp16179 {
tmp16168 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W474014643)
}
__typedArg0 := W474014643
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

W474114644 := tmp16168
_ = W474114644

tmp16169 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W474014643)
}
__typedArg0 := W474014643
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16170 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16169, V4553)


W474214645 := tmp16170
_ = W474214645

tmp16178 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W474214645, Nil)
}
__typedArg0 := W474214645
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16171 Obj

if True == tmp16178 {
tmp16172 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16172

tmp16173 := MakeNative(func(__e *ControlFlow) {
tmp16174 := Call(__e, PrimFunc(symshen_4deref), W474114644, V4553)


tmp16175 := Call(__e, PrimFunc(symshen_4rectify_1type), tmp16174)


tmp16176 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), W473914642, V4551, V4552, V4553, V4554, W455714474, V4556)
return
}, 0)

__e.TailApply(PrimFunc(symis_b), tmp16175, V4551, V4553, V4554, W455714474, tmp16176)
return


}, 0)

tmp16177 := Call(__e, PrimFunc(symshen_4cut), V4553, V4554, W455714474, tmp16173)


ifres16171 = tmp16177


} else {
ifres16171 = False


}

ifres16167 = ifres16171


} else {
ifres16167 = False


}

ifres16163 = ifres16167


} else {
ifres16163 = False


}

ifres16160 = ifres16163


} else {
ifres16160 = False


}

ifres16157 = ifres16160


} else {
ifres16157 = False


}

ifres16155 = ifres16157


} else {
ifres16155 = False


}

W473514638 := ifres16155
_ = W473514638

tmp16270 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W473514638, False)
}
__typedArg0 := W473514638
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16270 {
tmp16214 := Call(__e, PrimFunc(symshen_4unlocked_2), V4554)


var ifres16184 Obj

if True == tmp16214 {
tmp16185 := Call(__e, PrimFunc(symshen_4lazyderef), V4550, V4553)


W474414647 := tmp16185
_ = W474414647

tmp16213 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W474414647)
}
__typedArg0 := W474414647
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16186 Obj

if True == tmp16213 {
tmp16187 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W474414647)
}
__typedArg0 := W474414647
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16188 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16187, V4553)


W474514648 := tmp16188
_ = W474514648

tmp16212 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W474514648, symshen_4input_1h_7)
}
__typedArg0 := W474514648
__typedArg1 := symshen_4input_1h_7
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16189 Obj

if True == tmp16212 {
tmp16190 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W474414647)
}
__typedArg0 := W474414647
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16191 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16190, V4553)


W474614649 := tmp16191
_ = W474614649

tmp16211 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W474614649)
}
__typedArg0 := W474614649
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16192 Obj

if True == tmp16211 {
tmp16193 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W474614649)
}
__typedArg0 := W474614649
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

W474714650 := tmp16193
_ = W474714650

tmp16194 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W474614649)
}
__typedArg0 := W474614649
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16195 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16194, V4553)


W474814651 := tmp16195
_ = W474814651

tmp16210 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W474814651)
}
__typedArg0 := W474814651
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16196 Obj

if True == tmp16210 {
tmp16197 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W474814651)
}
__typedArg0 := W474814651
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

W474914652 := tmp16197
_ = W474914652

tmp16198 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W474814651)
}
__typedArg0 := W474814651
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16199 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16198, V4553)


W475014653 := tmp16199
_ = W475014653

tmp16209 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W475014653, Nil)
}
__typedArg0 := W475014653
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16200 Obj

if True == tmp16209 {
tmp16201 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16201

tmp16202 := Call(__e, PrimFunc(symshen_4deref), W474714650, V4553)


tmp16203 := Call(__e, PrimFunc(symshen_4rdecons), tmp16202)


tmp16204 := Call(__e, PrimFunc(symshen_4rectify_1type), tmp16203)


tmp16205 := MakeNative(func(__e *ControlFlow) {
tmp16206 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symin, Nil)
}
__typedArg0 := symin
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16207 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp16206)
}
__typedArg0 := symstream
__typedArg1 := tmp16206
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4system_1S_1h), W474914652, tmp16207, V4552, V4553, V4554, W455714474, V4556)
return


}, 0)

tmp16208 := Call(__e, PrimFunc(symis), V4551, tmp16204, V4553, V4554, W455714474, tmp16205)


ifres16200 = tmp16208


} else {
ifres16200 = False


}

ifres16196 = ifres16200


} else {
ifres16196 = False


}

ifres16192 = ifres16196


} else {
ifres16192 = False


}

ifres16189 = ifres16192


} else {
ifres16189 = False


}

ifres16186 = ifres16189


} else {
ifres16186 = False


}

ifres16184 = ifres16186


} else {
ifres16184 = False


}

W474314646 := ifres16184
_ = W474314646

tmp16268 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W474314646, False)
}
__typedArg0 := W474314646
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16268 {
tmp16243 := Call(__e, PrimFunc(symshen_4unlocked_2), V4554)


var ifres16215 Obj

if True == tmp16243 {
tmp16216 := Call(__e, PrimFunc(symshen_4lazyderef), V4550, V4553)


W475214655 := tmp16216
_ = W475214655

tmp16242 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W475214655)
}
__typedArg0 := W475214655
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16217 Obj

if True == tmp16242 {
tmp16218 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W475214655)
}
__typedArg0 := W475214655
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16219 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16218, V4553)


W475314656 := tmp16219
_ = W475314656

tmp16241 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W475314656, symset)
}
__typedArg0 := W475314656
__typedArg1 := symset
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16220 Obj

if True == tmp16241 {
tmp16221 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W475214655)
}
__typedArg0 := W475214655
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16222 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16221, V4553)


W475414657 := tmp16222
_ = W475414657

tmp16240 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W475414657)
}
__typedArg0 := W475414657
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16223 Obj

if True == tmp16240 {
tmp16224 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W475414657)
}
__typedArg0 := W475414657
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

W475514658 := tmp16224
_ = W475514658

tmp16225 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W475414657)
}
__typedArg0 := W475414657
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16226 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16225, V4553)


W475614659 := tmp16226
_ = W475614659

tmp16239 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W475614659)
}
__typedArg0 := W475614659
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16227 Obj

if True == tmp16239 {
tmp16228 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W475614659)
}
__typedArg0 := W475614659
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

W475714660 := tmp16228
_ = W475714660

tmp16229 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W475614659)
}
__typedArg0 := W475614659
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16230 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16229, V4553)


W475814661 := tmp16230
_ = W475814661

tmp16238 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W475814661, Nil)
}
__typedArg0 := W475814661
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16231 Obj

if True == tmp16238 {
tmp16232 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16232

tmp16233 := MakeNative(func(__e *ControlFlow) {
tmp16234 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W475514658, Nil)
}
__typedArg0 := W475514658
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16235 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvalue, tmp16234)
}
__typedArg0 := symvalue
__typedArg1 := tmp16234
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16236 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), W475714660, V4551, V4552, V4553, V4554, W455714474, V4556)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), tmp16235, V4551, V4552, V4553, V4554, W455714474, tmp16236)
return


}, 0)

tmp16237 := Call(__e, PrimFunc(symshen_4system_1S_1h), W475514658, symsymbol, V4552, V4553, V4554, W455714474, tmp16233)


ifres16231 = tmp16237


} else {
ifres16231 = False


}

ifres16227 = ifres16231


} else {
ifres16227 = False


}

ifres16223 = ifres16227


} else {
ifres16223 = False


}

ifres16220 = ifres16223


} else {
ifres16220 = False


}

ifres16217 = ifres16220


} else {
ifres16217 = False


}

ifres16215 = ifres16217


} else {
ifres16215 = False


}

W475114654 := ifres16215
_ = W475114654

tmp16266 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W475114654, False)
}
__typedArg0 := W475114654
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16266 {
tmp16251 := Call(__e, PrimFunc(symshen_4unlocked_2), V4554)


var ifres16244 Obj

if True == tmp16251 {
tmp16245 := Call(__e, PrimFunc(symshen_4newpv), V4553)


W476014663 := tmp16245
_ = W476014663

tmp16246 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16246

tmp16247 := MakeNative(func(__e *ControlFlow) {
tmp16248 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), V4550, V4551, W476014663, V4553, V4554, W455714474, V4556)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V4553, V4554, W455714474, tmp16248)
return


}, 0)

tmp16249 := Call(__e, PrimFunc(symshen_4l_1rules), V4552, W476014663, False, V4553, V4554, W455714474, tmp16247)


tmp16250 := Call(__e, PrimFunc(symshen_4gc), V4553, tmp16249)


ifres16244 = tmp16250


} else {
ifres16244 = False


}

W475914662 := ifres16244
_ = W475914662

tmp16264 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W475914662, False)
}
__typedArg0 := W475914662
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16264 {
tmp16260 := Call(__e, PrimFunc(symshen_4unlocked_2), V4554)


var ifres16252 Obj

if True == tmp16260 {
tmp16253 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16253

tmp16254 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp16255 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V4551, Nil)
}
__typedArg0 := V4551
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16256 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16254, tmp16255)
}
__typedArg0 := tmp16254
__typedArg1 := tmp16255
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16257 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V4550, tmp16256)
}
__typedArg0 := V4550
__typedArg1 := tmp16256
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16258 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_ddatatypes_d)
}
__typedArg0 := symshen_4_ddatatypes_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp16259 := Call(__e, PrimFunc(symshen_4search_1user_1datatypes), tmp16257, V4552, tmp16258, V4553, V4554, W455714474, V4556)


ifres16252 = tmp16259


} else {
ifres16252 = False


}

W476114664 := ifres16252
_ = W476114664

tmp16262 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W476114664, False)
}
__typedArg0 := W476114664
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16262 {
__e.TailApply(PrimFunc(symshen_4unlock), V4554, W455714474)
return
} else {
__e.Return(W476114664)
return
}


} else {
__e.Return(W475914662)
return
}


} else {
__e.Return(W475114654)
return
}


} else {
__e.Return(W474314646)
return
}


} else {
__e.Return(W473514638)
return
}


} else {
__e.Return(W471414619)
return
}


} else {
__e.Return(W470114606)
return
}


} else {
__e.Return(W467114580)
return
}


} else {
__e.Return(W466114570)
return
}


} else {
__e.Return(W464014551)
return
}


} else {
__e.Return(W461214527)
return
}


} else {
__e.Return(W459114508)
return
}


} else {
__e.Return(W458414501)
return
}


} else {
__e.Return(W457714494)
return
}


} else {
__e.Return(W457114488)
return
}


} else {
__e.Return(W456514482)
return
}


} else {
__e.Return(W456114478)
return
}


} else {
__e.Return(W456014477)
return
}


} else {
__e.Return(W455914476)
return
}


} else {
__e.Return(W455814475)
return
}


}, 7)

tmp16301 := Call(__e, ns2_1set, symshen_4system_1S_1h, tmp15550)


_ = tmp16301

tmp16302 := MakeNative(func(__e *ControlFlow) {
V4762 := __e.Get(1)
_ = V4762
tmp16336 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4762)
}
__typedArg0 := V4762
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16317 Obj

if True == tmp16336 {
tmp16334 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4762)
}
__typedArg0 := V4762
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16335 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcons, tmp16334)
}
__typedArg0 := symcons
__typedArg1 := tmp16334
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16319 Obj

if True == tmp16335 {
tmp16332 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4762)
}
__typedArg0 := V4762
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16333 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp16332)
}
__typedArg0 := tmp16332
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16321 Obj

if True == tmp16333 {
tmp16329 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4762)
}
__typedArg0 := V4762
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16330 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16329)
}
__typedArg0 := tmp16329
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16331 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp16330)
}
__typedArg0 := tmp16330
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16323 Obj

if True == tmp16331 {
tmp16325 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4762)
}
__typedArg0 := V4762
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16326 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16325)
}
__typedArg0 := tmp16325
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16327 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16326)
}
__typedArg0 := tmp16326
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16328 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp16327)
}
__typedArg0 := Nil
__typedArg1 := tmp16327
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16324 Obj

if True == tmp16328 {
ifres16324 = True


} else {
ifres16324 = False


}

ifres16323 = ifres16324


} else {
ifres16323 = False


}

var ifres16322 Obj

if True == ifres16323 {
ifres16322 = True


} else {
ifres16322 = False


}

ifres16321 = ifres16322


} else {
ifres16321 = False


}

var ifres16320 Obj

if True == ifres16321 {
ifres16320 = True


} else {
ifres16320 = False


}

ifres16319 = ifres16320


} else {
ifres16319 = False


}

var ifres16318 Obj

if True == ifres16319 {
ifres16318 = True


} else {
ifres16318 = False


}

ifres16317 = ifres16318


} else {
ifres16317 = False


}

if True == ifres16317 {
tmp16303 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4762)
}
__typedArg0 := V4762
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16304 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp16303)
}
__typedArg0 := tmp16303
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16305 := Call(__e, PrimFunc(symshen_4rdecons), tmp16304)


tmp16306 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4762)
}
__typedArg0 := V4762
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16307 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16306)
}
__typedArg0 := tmp16306
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16308 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp16307)
}
__typedArg0 := tmp16307
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16309 := Call(__e, PrimFunc(symshen_4rdecons), tmp16308)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16305, tmp16309)
}
__typedArg0 := tmp16305
__typedArg1 := tmp16309
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp16315 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4762)
}
__typedArg0 := V4762
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp16315 {
tmp16310 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4762)
}
__typedArg0 := V4762
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16311 := Call(__e, PrimFunc(symshen_4rdecons), tmp16310)


tmp16312 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4762)
}
__typedArg0 := V4762
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16313 := Call(__e, PrimFunc(symshen_4rdecons), tmp16312)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16311, tmp16313)
}
__typedArg0 := tmp16311
__typedArg1 := tmp16313
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return(V4762)
return
}


}


}, 1)

tmp16337 := Call(__e, ns2_1set, symshen_4rdecons, tmp16302)


_ = tmp16337

tmp16338 := MakeNative(func(__e *ControlFlow) {
V4763 := __e.Get(1)
_ = V4763
V4764 := __e.Get(2)
_ = V4764
V4765 := __e.Get(3)
_ = V4765
V4766 := __e.Get(4)
_ = V4766
V4767 := __e.Get(5)
_ = V4767
V4768 := __e.Get(6)
_ = V4768
tmp16351 := Call(__e, PrimFunc(symshen_4unlocked_2), V4766)


var ifres16339 Obj

if True == tmp16351 {
tmp16340 := Call(__e, PrimFunc(symshen_4lazyderef), V4764, V4765)


W477014666 := tmp16340
_ = W477014666

tmp16341 := MakeNative(func(__e *ControlFlow) {
tmp16342 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16342

tmp16343 := Call(__e, PrimFunc(symshen_4lazyderef), V4763, V4765)


tmp16344 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnumber_2) {
return PrimIsNumber(tmp16343)
}
__typedArg0 := tmp16343
return Call(__e, PrimFunc(symnumber_2), __typedArg0)
})()

__e.TailApply(PrimFunc(symwhen), tmp16344, V4765, V4766, V4767, V4768)
return


}, 0)

W477114667 := tmp16341
_ = W477114667

tmp16350 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W477014666, symnumber)
}
__typedArg0 := W477014666
__typedArg1 := symnumber
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16345 Obj

if True == tmp16350 {
tmp16346 := Call(__e, PrimFunc(symthaw), W477114667)


ifres16345 = tmp16346


} else {
tmp16349 := Call(__e, PrimFunc(symshen_4pvar_2), W477014666)


var ifres16347 Obj

if True == tmp16349 {
tmp16348 := Call(__e, PrimFunc(symshen_4bind_b), W477014666, symnumber, V4765, W477114667)


ifres16347 = tmp16348


} else {
ifres16347 = False


}

ifres16345 = ifres16347


}

ifres16339 = ifres16345


} else {
ifres16339 = False


}

W476914665 := ifres16339
_ = W476914665

tmp16441 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W476914665, False)
}
__typedArg0 := W476914665
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16441 {
tmp16364 := Call(__e, PrimFunc(symshen_4unlocked_2), V4766)


var ifres16352 Obj

if True == tmp16364 {
tmp16353 := Call(__e, PrimFunc(symshen_4lazyderef), V4764, V4765)


W477314669 := tmp16353
_ = W477314669

tmp16354 := MakeNative(func(__e *ControlFlow) {
tmp16355 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16355

tmp16356 := Call(__e, PrimFunc(symshen_4lazyderef), V4763, V4765)


tmp16357 := Call(__e, PrimFunc(symboolean_2), tmp16356)


__e.TailApply(PrimFunc(symwhen), tmp16357, V4765, V4766, V4767, V4768)
return


}, 0)

W477414670 := tmp16354
_ = W477414670

tmp16363 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W477314669, symboolean)
}
__typedArg0 := W477314669
__typedArg1 := symboolean
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16358 Obj

if True == tmp16363 {
tmp16359 := Call(__e, PrimFunc(symthaw), W477414670)


ifres16358 = tmp16359


} else {
tmp16362 := Call(__e, PrimFunc(symshen_4pvar_2), W477314669)


var ifres16360 Obj

if True == tmp16362 {
tmp16361 := Call(__e, PrimFunc(symshen_4bind_b), W477314669, symboolean, V4765, W477414670)


ifres16360 = tmp16361


} else {
ifres16360 = False


}

ifres16358 = ifres16360


}

ifres16352 = ifres16358


} else {
ifres16352 = False


}

W477214668 := ifres16352
_ = W477214668

tmp16439 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W477214668, False)
}
__typedArg0 := W477214668
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16439 {
tmp16377 := Call(__e, PrimFunc(symshen_4unlocked_2), V4766)


var ifres16365 Obj

if True == tmp16377 {
tmp16366 := Call(__e, PrimFunc(symshen_4lazyderef), V4764, V4765)


W477614672 := tmp16366
_ = W477614672

tmp16367 := MakeNative(func(__e *ControlFlow) {
tmp16368 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16368

tmp16369 := Call(__e, PrimFunc(symshen_4lazyderef), V4763, V4765)


tmp16370 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_2) {
return PrimIsString(tmp16369)
}
__typedArg0 := tmp16369
return Call(__e, PrimFunc(symstring_2), __typedArg0)
})()

__e.TailApply(PrimFunc(symwhen), tmp16370, V4765, V4766, V4767, V4768)
return


}, 0)

W477714673 := tmp16367
_ = W477714673

tmp16376 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W477614672, symstring)
}
__typedArg0 := W477614672
__typedArg1 := symstring
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16371 Obj

if True == tmp16376 {
tmp16372 := Call(__e, PrimFunc(symthaw), W477714673)


ifres16371 = tmp16372


} else {
tmp16375 := Call(__e, PrimFunc(symshen_4pvar_2), W477614672)


var ifres16373 Obj

if True == tmp16375 {
tmp16374 := Call(__e, PrimFunc(symshen_4bind_b), W477614672, symstring, V4765, W477714673)


ifres16373 = tmp16374


} else {
ifres16373 = False


}

ifres16371 = ifres16373


}

ifres16365 = ifres16371


} else {
ifres16365 = False


}

W477514671 := ifres16365
_ = W477514671

tmp16437 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W477514671, False)
}
__typedArg0 := W477514671
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16437 {
tmp16390 := Call(__e, PrimFunc(symshen_4unlocked_2), V4766)


var ifres16378 Obj

if True == tmp16390 {
tmp16379 := Call(__e, PrimFunc(symshen_4lazyderef), V4764, V4765)


W477914675 := tmp16379
_ = W477914675

tmp16380 := MakeNative(func(__e *ControlFlow) {
tmp16381 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16381

tmp16382 := Call(__e, PrimFunc(symshen_4lazyderef), V4763, V4765)


tmp16383 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsymbol_2) {
return PrimIsSymbol(tmp16382)
}
__typedArg0 := tmp16382
return Call(__e, PrimFunc(symsymbol_2), __typedArg0)
})()

__e.TailApply(PrimFunc(symwhen), tmp16383, V4765, V4766, V4767, V4768)
return


}, 0)

W478014676 := tmp16380
_ = W478014676

tmp16389 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W477914675, symsymbol)
}
__typedArg0 := W477914675
__typedArg1 := symsymbol
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16384 Obj

if True == tmp16389 {
tmp16385 := Call(__e, PrimFunc(symthaw), W478014676)


ifres16384 = tmp16385


} else {
tmp16388 := Call(__e, PrimFunc(symshen_4pvar_2), W477914675)


var ifres16386 Obj

if True == tmp16388 {
tmp16387 := Call(__e, PrimFunc(symshen_4bind_b), W477914675, symsymbol, V4765, W478014676)


ifres16386 = tmp16387


} else {
ifres16386 = False


}

ifres16384 = ifres16386


}

ifres16378 = ifres16384


} else {
ifres16378 = False


}

W477814674 := ifres16378
_ = W477814674

tmp16435 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W477814674, False)
}
__typedArg0 := W477814674
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16435 {
tmp16433 := Call(__e, PrimFunc(symshen_4unlocked_2), V4766)


if True == tmp16433 {
tmp16391 := Call(__e, PrimFunc(symshen_4lazyderef), V4763, V4765)


W478114677 := tmp16391
_ = W478114677

tmp16431 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W478114677, Nil)
}
__typedArg0 := W478114677
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16431 {
tmp16392 := Call(__e, PrimFunc(symshen_4lazyderef), V4764, V4765)


W478214678 := tmp16392
_ = W478214678

tmp16393 := MakeNative(func(__e *ControlFlow) {
Z4784 := __e.Get(1)
_ = Z4784
tmp16394 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16394

__e.TailApply(PrimFunc(symthaw), V4768)
return


}, 1)

W478314679 := tmp16393
_ = W478314679

tmp16429 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W478214678)
}
__typedArg0 := W478214678
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp16429 {
tmp16395 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W478214678)
}
__typedArg0 := W478214678
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16396 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16395, V4765)


W478514680 := tmp16396
_ = W478514680

tmp16397 := MakeNative(func(__e *ControlFlow) {
tmp16398 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W478214678)
}
__typedArg0 := W478214678
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16399 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16398, V4765)


W478714682 := tmp16399
_ = W478714682

tmp16400 := MakeNative(func(__e *ControlFlow) {
Z4789 := __e.Get(1)
_ = Z4789
__e.TailApply(W478314679, Z4789)
return
}, 1)

W478814683 := tmp16400
_ = W478814683

tmp16416 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W478714682)
}
__typedArg0 := W478714682
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp16416 {
tmp16401 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W478714682)
}
__typedArg0 := W478714682
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

W479014684 := tmp16401
_ = W479014684

tmp16402 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W478714682)
}
__typedArg0 := W478714682
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16403 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16402, V4765)


W479114685 := tmp16403
_ = W479114685

tmp16404 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W478814683, W479014684)
return
}, 0)

W479214686 := tmp16404
_ = W479214686

tmp16408 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W479114685, Nil)
}
__typedArg0 := W479114685
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16408 {
__e.TailApply(PrimFunc(symthaw), W479214686)
return
} else {
tmp16406 := Call(__e, PrimFunc(symshen_4pvar_2), W479114685)


if True == tmp16406 {
__e.TailApply(PrimFunc(symshen_4bind_b), W479114685, Nil, V4765, W479214686)
return
} else {
__e.Return(False)
return
}


}


} else {
tmp16414 := Call(__e, PrimFunc(symshen_4pvar_2), W478714682)


if True == tmp16414 {
tmp16409 := Call(__e, PrimFunc(symshen_4newpv), V4765)


W479314687 := tmp16409
_ = W479314687

tmp16410 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W479314687, Nil)
}
__typedArg0 := W479314687
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16411 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W478814683, W479314687)
return
}, 0)

tmp16412 := Call(__e, PrimFunc(symshen_4bind_b), W478714682, tmp16410, V4765, tmp16411)


__e.TailApply(PrimFunc(symshen_4gc), V4765, tmp16412)
return


} else {
__e.Return(False)
return
}


}


}, 0)

W478614681 := tmp16397
_ = W478614681

tmp16420 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W478514680, symlist)
}
__typedArg0 := W478514680
__typedArg1 := symlist
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16420 {
__e.TailApply(PrimFunc(symthaw), W478614681)
return
} else {
tmp16418 := Call(__e, PrimFunc(symshen_4pvar_2), W478514680)


if True == tmp16418 {
__e.TailApply(PrimFunc(symshen_4bind_b), W478514680, symlist, V4765, W478614681)
return
} else {
__e.Return(False)
return
}


}


} else {
tmp16427 := Call(__e, PrimFunc(symshen_4pvar_2), W478214678)


if True == tmp16427 {
tmp16421 := Call(__e, PrimFunc(symshen_4newpv), V4765)


W479414688 := tmp16421
_ = W479414688

tmp16422 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W479414688, Nil)
}
__typedArg0 := W479414688
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16423 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp16422)
}
__typedArg0 := symlist
__typedArg1 := tmp16422
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16424 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W478314679, W479414688)
return
}, 0)

tmp16425 := Call(__e, PrimFunc(symshen_4bind_b), W478214678, tmp16423, V4765, tmp16424)


__e.TailApply(PrimFunc(symshen_4gc), V4765, tmp16425)
return


} else {
__e.Return(False)
return
}


}


} else {
__e.Return(False)
return
}


} else {
__e.Return(False)
return
}


} else {
__e.Return(W477814674)
return
}


} else {
__e.Return(W477514671)
return
}


} else {
__e.Return(W477214668)
return
}


} else {
__e.Return(W476914665)
return
}


}, 6)

tmp16442 := Call(__e, ns2_1set, symshen_4primitive, tmp16338)


_ = tmp16442

tmp16443 := MakeNative(func(__e *ControlFlow) {
V4795 := __e.Get(1)
_ = V4795
V4796 := __e.Get(2)
_ = V4796
V4797 := __e.Get(3)
_ = V4797
V4798 := __e.Get(4)
_ = V4798
V4799 := __e.Get(5)
_ = V4799
V4800 := __e.Get(6)
_ = V4800
V4801 := __e.Get(7)
_ = V4801
tmp16477 := Call(__e, PrimFunc(symshen_4unlocked_2), V4799)


var ifres16444 Obj

if True == tmp16477 {
tmp16445 := Call(__e, PrimFunc(symshen_4lazyderef), V4797, V4798)


W480314690 := tmp16445
_ = W480314690

tmp16476 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W480314690)
}
__typedArg0 := W480314690
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16446 Obj

if True == tmp16476 {
tmp16447 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W480314690)
}
__typedArg0 := W480314690
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16448 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16447, V4798)


W480414691 := tmp16448
_ = W480414691

tmp16475 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W480414691)
}
__typedArg0 := W480414691
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16449 Obj

if True == tmp16475 {
tmp16450 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W480414691)
}
__typedArg0 := W480414691
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

W480514692 := tmp16450
_ = W480514692

tmp16451 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W480414691)
}
__typedArg0 := W480414691
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16452 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16451, V4798)


W480614693 := tmp16452
_ = W480614693

tmp16474 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W480614693)
}
__typedArg0 := W480614693
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16453 Obj

if True == tmp16474 {
tmp16454 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W480614693)
}
__typedArg0 := W480614693
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

W480714694 := tmp16454
_ = W480714694

tmp16455 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W480614693)
}
__typedArg0 := W480614693
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16456 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16455, V4798)


W480814695 := tmp16456
_ = W480814695

tmp16473 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W480814695)
}
__typedArg0 := W480814695
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16457 Obj

if True == tmp16473 {
tmp16458 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W480814695)
}
__typedArg0 := W480814695
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

W480914696 := tmp16458
_ = W480914696

tmp16459 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W480814695)
}
__typedArg0 := W480814695
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16460 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16459, V4798)


W481014697 := tmp16460
_ = W481014697

tmp16472 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W481014697, Nil)
}
__typedArg0 := W481014697
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16461 Obj

if True == tmp16472 {
tmp16462 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16462

tmp16463 := Call(__e, PrimFunc(symshen_4deref), W480714694, V4798)


tmp16464 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp16465 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp16463, tmp16464)
}
__typedArg0 := tmp16463
__typedArg1 := tmp16464
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

tmp16466 := MakeNative(func(__e *ControlFlow) {
tmp16467 := Call(__e, PrimFunc(symshen_4deref), V4795, V4798)


tmp16468 := Call(__e, PrimFunc(symshen_4deref), W480514692, V4798)


tmp16469 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp16467, tmp16468)
}
__typedArg0 := tmp16467
__typedArg1 := tmp16468
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

tmp16470 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symis_b), V4796, W480914696, V4798, V4799, V4800, V4801)
return
}, 0)

__e.TailApply(PrimFunc(symwhen), tmp16469, V4798, V4799, V4800, tmp16470)
return


}, 0)

tmp16471 := Call(__e, PrimFunc(symwhen), tmp16465, V4798, V4799, V4800, tmp16466)


ifres16461 = tmp16471


} else {
ifres16461 = False


}

ifres16457 = ifres16461


} else {
ifres16457 = False


}

ifres16453 = ifres16457


} else {
ifres16453 = False


}

ifres16449 = ifres16453


} else {
ifres16449 = False


}

ifres16446 = ifres16449


} else {
ifres16446 = False


}

ifres16444 = ifres16446


} else {
ifres16444 = False


}

W480214689 := ifres16444
_ = W480214689

tmp16486 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W480214689, False)
}
__typedArg0 := W480214689
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16486 {
tmp16484 := Call(__e, PrimFunc(symshen_4unlocked_2), V4799)


if True == tmp16484 {
tmp16478 := Call(__e, PrimFunc(symshen_4lazyderef), V4797, V4798)


W481114698 := tmp16478
_ = W481114698

tmp16482 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W481114698)
}
__typedArg0 := W481114698
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp16482 {
tmp16479 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W481114698)
}
__typedArg0 := W481114698
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

W481214699 := tmp16479
_ = W481214699

tmp16480 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16480

__e.TailApply(PrimFunc(symshen_4by_1hypothesis), V4795, V4796, W481214699, V4798, V4799, V4800, V4801)
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
__e.Return(W480214689)
return
}


}, 7)

tmp16487 := Call(__e, ns2_1set, symshen_4by_1hypothesis, tmp16443)


_ = tmp16487

tmp16488 := MakeNative(func(__e *ControlFlow) {
V4813 := __e.Get(1)
_ = V4813
V4814 := __e.Get(2)
_ = V4814
V4815 := __e.Get(3)
_ = V4815
V4816 := __e.Get(4)
_ = V4816
V4817 := __e.Get(5)
_ = V4817
V4818 := __e.Get(6)
_ = V4818
tmp16493 := Call(__e, PrimFunc(symshen_4unlocked_2), V4816)


if True == tmp16493 {
tmp16489 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16489

tmp16490 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp16491 := Call(__e, PrimFunc(symassoc), V4813, tmp16490)


__e.TailApply(PrimFunc(symshen_4sigf), tmp16491, V4814, V4815, V4816, V4817, V4818)
return


} else {
__e.Return(False)
return
}


}, 6)

tmp16494 := Call(__e, ns2_1set, symshen_4lookupsig, tmp16488)


_ = tmp16494

tmp16495 := MakeNative(func(__e *ControlFlow) {
V4833 := __e.Get(1)
_ = V4833
V4834 := __e.Get(2)
_ = V4834
V4835 := __e.Get(3)
_ = V4835
V4836 := __e.Get(4)
_ = V4836
V4837 := __e.Get(5)
_ = V4837
V4838 := __e.Get(6)
_ = V4838
tmp16502 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4833)
}
__typedArg0 := V4833
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp16502 {
tmp16496 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4833)
}
__typedArg0 := V4833
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16497 := Call(__e, tmp16496, V4834)


tmp16498 := Call(__e, tmp16497, V4835)


tmp16499 := Call(__e, tmp16498, V4836)


tmp16500 := Call(__e, tmp16499, V4837)


__e.TailApply(tmp16500, V4838)
return


} else {
__e.Return(False)
return
}


}, 6)

tmp16503 := Call(__e, ns2_1set, symshen_4sigf, tmp16495)


_ = tmp16503

tmp16504 := MakeNative(func(__e *ControlFlow) {
V4839 := __e.Get(1)
_ = V4839
tmp16505 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symabsvector) {
return PrimAbsvector(MakeNumber(3))
}
__typedArg0 := MakeNumber(3)
return Call(__e, PrimFunc(symabsvector), __typedArg0)
})()

W484014700 := tmp16505
_ = W484014700

tmp16506 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(W484014700, MakeNumber(0), symshen_4print_1freshterm)
}
__typedArg0 := W484014700
__typedArg1 := MakeNumber(0)
__typedArg2 := symshen_4print_1freshterm
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

W484114701 := tmp16506
_ = W484114701

tmp16507 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(W484114701, MakeNumber(1), V4839)
}
__typedArg0 := W484114701
__typedArg1 := MakeNumber(1)
__typedArg2 := V4839
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

W484214702 := tmp16507
_ = W484214702

tmp16508 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dgensym_d)
}
__typedArg0 := symshen_4_dgensym_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp16510 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dgensym_d, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(1))
__typedN1, __typedOK1 := TypedFloat64(tmp16508)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp16508
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())
}
__typedArg0 := symshen_4_dgensym_d
__typedArg1 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(1))
__typedN1, __typedOK1 := TypedFloat64(tmp16508)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp16508
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

tmp16511 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(W484214702, MakeNumber(2), tmp16510)
}
__typedArg0 := W484214702
__typedArg1 := MakeNumber(2)
__typedArg2 := tmp16510
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

W484314703 := tmp16511
_ = W484314703

__e.Return(W484314703)
return


}, 1)

tmp16512 := Call(__e, ns2_1set, symshen_4freshterm, tmp16504)


_ = tmp16512

tmp16513 := MakeNative(func(__e *ControlFlow) {
V4844 := __e.Get(1)
_ = V4844
tmp16514 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V4844, MakeNumber(1))
}
__typedArg0 := V4844
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})()

tmp16515 := Call(__e, PrimFunc(symshen_4app), tmp16514, MakeString(""), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("&&"))
__typedS1, __typedOK1 := TypedString(tmp16515)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("&&")
__typedArg1 := tmp16515
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp16516 := Call(__e, ns2_1set, symshen_4print_1freshterm, tmp16513)


_ = tmp16516

tmp16517 := MakeNative(func(__e *ControlFlow) {
V4845 := __e.Get(1)
_ = V4845
V4846 := __e.Get(2)
_ = V4846
V4847 := __e.Get(3)
_ = V4847
V4848 := __e.Get(4)
_ = V4848
V4849 := __e.Get(5)
_ = V4849
V4850 := __e.Get(6)
_ = V4850
V4851 := __e.Get(7)
_ = V4851
tmp16534 := Call(__e, PrimFunc(symshen_4unlocked_2), V4849)


var ifres16518 Obj

if True == tmp16534 {
tmp16519 := Call(__e, PrimFunc(symshen_4lazyderef), V4847, V4848)


W485314705 := tmp16519
_ = W485314705

tmp16533 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W485314705)
}
__typedArg0 := W485314705
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16520 Obj

if True == tmp16533 {
tmp16521 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W485314705)
}
__typedArg0 := W485314705
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16522 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16521, V4848)


W485414706 := tmp16522
_ = W485414706

tmp16532 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W485414706)
}
__typedArg0 := W485414706
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16523 Obj

if True == tmp16532 {
tmp16524 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W485414706)
}
__typedArg0 := W485414706
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

W485514707 := tmp16524
_ = W485514707

tmp16525 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16525

tmp16526 := Call(__e, PrimFunc(symshen_4deref), W485514707, V4848)


tmp16527 := Call(__e, PrimFunc(symshen_4deref), V4845, V4848)


tmp16528 := Call(__e, tmp16526, tmp16527)


tmp16529 := Call(__e, PrimFunc(symshen_4deref), V4846, V4848)


tmp16530 := Call(__e, tmp16528, tmp16529)


tmp16531 := Call(__e, PrimFunc(symcall), tmp16530, V4848, V4849, V4850, V4851)


ifres16523 = tmp16531


} else {
ifres16523 = False


}

ifres16520 = ifres16523


} else {
ifres16520 = False


}

ifres16518 = ifres16520


} else {
ifres16518 = False


}

W485214704 := ifres16518
_ = W485214704

tmp16543 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W485214704, False)
}
__typedArg0 := W485214704
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16543 {
tmp16541 := Call(__e, PrimFunc(symshen_4unlocked_2), V4849)


if True == tmp16541 {
tmp16535 := Call(__e, PrimFunc(symshen_4lazyderef), V4847, V4848)


W485614708 := tmp16535
_ = W485614708

tmp16539 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W485614708)
}
__typedArg0 := W485614708
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp16539 {
tmp16536 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W485614708)
}
__typedArg0 := W485614708
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

W485714709 := tmp16536
_ = W485714709

tmp16537 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16537

__e.TailApply(PrimFunc(symshen_4search_1user_1datatypes), V4845, V4846, W485714709, V4848, V4849, V4850, V4851)
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
__e.Return(W485214704)
return
}


}, 7)

tmp16544 := Call(__e, ns2_1set, symshen_4search_1user_1datatypes, tmp16517)


_ = tmp16544

tmp16545 := MakeNative(func(__e *ControlFlow) {
V4858 := __e.Get(1)
_ = V4858
V4859 := __e.Get(2)
_ = V4859
V4860 := __e.Get(3)
_ = V4860
V4861 := __e.Get(4)
_ = V4861
V4862 := __e.Get(5)
_ = V4862
V4863 := __e.Get(6)
_ = V4863
V4864 := __e.Get(7)
_ = V4864
W486514710 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V4863)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V4863
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
_ = W486514710

tmp16557 := Call(__e, PrimFunc(symshen_4unlocked_2), V4862)


var ifres16547 Obj

if True == tmp16557 {
tmp16548 := Call(__e, PrimFunc(symshen_4lazyderef), V4858, V4861)


W486714712 := tmp16548
_ = W486714712

tmp16556 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W486714712, Nil)
}
__typedArg0 := W486714712
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16549 Obj

if True == tmp16556 {
tmp16550 := Call(__e, PrimFunc(symshen_4lazyderef), V4860, V4861)


W486814713 := tmp16550
_ = W486814713

tmp16555 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W486814713, True)
}
__typedArg0 := W486814713
__typedArg1 := True
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16551 Obj

if True == tmp16555 {
tmp16552 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16552

tmp16553 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symbind), V4859, Nil, V4861, V4862, W486514710, V4864)
return
}, 0)

tmp16554 := Call(__e, PrimFunc(symshen_4cut), V4861, V4862, W486514710, tmp16553)


ifres16551 = tmp16554


} else {
ifres16551 = False


}

ifres16549 = ifres16551


} else {
ifres16549 = False


}

ifres16547 = ifres16549


} else {
ifres16547 = False


}

W486614711 := ifres16547
_ = W486614711

tmp16900 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W486614711, False)
}
__typedArg0 := W486614711
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16900 {
tmp16636 := Call(__e, PrimFunc(symshen_4unlocked_2), V4862)


var ifres16558 Obj

if True == tmp16636 {
tmp16559 := Call(__e, PrimFunc(symshen_4lazyderef), V4858, V4861)


W487014715 := tmp16559
_ = W487014715

tmp16635 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W487014715)
}
__typedArg0 := W487014715
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16560 Obj

if True == tmp16635 {
tmp16561 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W487014715)
}
__typedArg0 := W487014715
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16562 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16561, V4861)


W487114716 := tmp16562
_ = W487114716

tmp16634 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W487114716)
}
__typedArg0 := W487114716
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16563 Obj

if True == tmp16634 {
tmp16564 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W487114716)
}
__typedArg0 := W487114716
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16565 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16564, V4861)


W487214717 := tmp16565
_ = W487214717

tmp16633 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W487214717)
}
__typedArg0 := W487214717
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16566 Obj

if True == tmp16633 {
tmp16567 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W487214717)
}
__typedArg0 := W487214717
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16568 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16567, V4861)


W487314718 := tmp16568
_ = W487314718

tmp16632 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W487314718, symcons)
}
__typedArg0 := W487314718
__typedArg1 := symcons
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16569 Obj

if True == tmp16632 {
tmp16570 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W487214717)
}
__typedArg0 := W487214717
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16571 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16570, V4861)


W487414719 := tmp16571
_ = W487414719

tmp16631 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W487414719)
}
__typedArg0 := W487414719
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16572 Obj

if True == tmp16631 {
tmp16573 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W487414719)
}
__typedArg0 := W487414719
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

W487514720 := tmp16573
_ = W487514720

tmp16574 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W487414719)
}
__typedArg0 := W487414719
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16575 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16574, V4861)


W487614721 := tmp16575
_ = W487614721

tmp16630 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W487614721)
}
__typedArg0 := W487614721
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16576 Obj

if True == tmp16630 {
tmp16577 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W487614721)
}
__typedArg0 := W487614721
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

W487714722 := tmp16577
_ = W487714722

tmp16578 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W487614721)
}
__typedArg0 := W487614721
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16579 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16578, V4861)


W487814723 := tmp16579
_ = W487814723

tmp16629 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W487814723, Nil)
}
__typedArg0 := W487814723
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16580 Obj

if True == tmp16629 {
tmp16581 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W487114716)
}
__typedArg0 := W487114716
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16582 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16581, V4861)


W487914724 := tmp16582
_ = W487914724

tmp16628 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W487914724)
}
__typedArg0 := W487914724
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16583 Obj

if True == tmp16628 {
tmp16584 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W487914724)
}
__typedArg0 := W487914724
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

W488014725 := tmp16584
_ = W488014725

tmp16585 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W487914724)
}
__typedArg0 := W487914724
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16586 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16585, V4861)


W488114726 := tmp16586
_ = W488114726

tmp16627 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W488114726)
}
__typedArg0 := W488114726
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16587 Obj

if True == tmp16627 {
tmp16588 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W488114726)
}
__typedArg0 := W488114726
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16589 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16588, V4861)


W488214727 := tmp16589
_ = W488214727

tmp16626 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W488214727)
}
__typedArg0 := W488214727
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16590 Obj

if True == tmp16626 {
tmp16591 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W488214727)
}
__typedArg0 := W488214727
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16592 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16591, V4861)


W488314728 := tmp16592
_ = W488314728

tmp16625 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W488314728, symlist)
}
__typedArg0 := W488314728
__typedArg1 := symlist
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16593 Obj

if True == tmp16625 {
tmp16594 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W488214727)
}
__typedArg0 := W488214727
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16595 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16594, V4861)


W488414729 := tmp16595
_ = W488414729

tmp16624 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W488414729)
}
__typedArg0 := W488414729
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16596 Obj

if True == tmp16624 {
tmp16597 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W488414729)
}
__typedArg0 := W488414729
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

W488514730 := tmp16597
_ = W488514730

tmp16598 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W488414729)
}
__typedArg0 := W488414729
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16599 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16598, V4861)


W488614731 := tmp16599
_ = W488614731

tmp16623 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W488614731, Nil)
}
__typedArg0 := W488614731
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16600 Obj

if True == tmp16623 {
tmp16601 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W488114726)
}
__typedArg0 := W488114726
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16602 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16601, V4861)


W488714732 := tmp16602
_ = W488714732

tmp16622 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W488714732, Nil)
}
__typedArg0 := W488714732
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16603 Obj

if True == tmp16622 {
tmp16604 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W487014715)
}
__typedArg0 := W487014715
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

W488814733 := tmp16604
_ = W488814733

tmp16605 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16605

tmp16606 := Call(__e, PrimFunc(symshen_4deref), W488014725, V4861)


tmp16607 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp16608 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp16606, tmp16607)
}
__typedArg0 := tmp16606
__typedArg1 := tmp16607
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

tmp16609 := MakeNative(func(__e *ControlFlow) {
tmp16610 := MakeNative(func(__e *ControlFlow) {
tmp16611 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W488514730, Nil)
}
__typedArg0 := W488514730
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16612 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W488014725, tmp16611)
}
__typedArg0 := W488014725
__typedArg1 := tmp16611
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16613 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W487514720, tmp16612)
}
__typedArg0 := W487514720
__typedArg1 := tmp16612
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16614 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W488514730, Nil)
}
__typedArg0 := W488514730
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16615 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp16614)
}
__typedArg0 := symlist
__typedArg1 := tmp16614
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16616 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16615, Nil)
}
__typedArg0 := tmp16615
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16617 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W488014725, tmp16616)
}
__typedArg0 := W488014725
__typedArg1 := tmp16616
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16618 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W487714722, tmp16617)
}
__typedArg0 := W487714722
__typedArg1 := tmp16617
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16619 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16618, W488814733)
}
__typedArg0 := tmp16618
__typedArg1 := W488814733
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16620 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16613, tmp16619)
}
__typedArg0 := tmp16613
__typedArg1 := tmp16619
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4l_1rules), tmp16620, V4859, True, V4861, V4862, W486514710, V4864)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V4861, V4862, W486514710, tmp16610)
return


}, 0)

tmp16621 := Call(__e, PrimFunc(symwhen), tmp16608, V4861, V4862, W486514710, tmp16609)


ifres16603 = tmp16621


} else {
ifres16603 = False


}

ifres16600 = ifres16603


} else {
ifres16600 = False


}

ifres16596 = ifres16600


} else {
ifres16596 = False


}

ifres16593 = ifres16596


} else {
ifres16593 = False


}

ifres16590 = ifres16593


} else {
ifres16590 = False


}

ifres16587 = ifres16590


} else {
ifres16587 = False


}

ifres16583 = ifres16587


} else {
ifres16583 = False


}

ifres16580 = ifres16583


} else {
ifres16580 = False


}

ifres16576 = ifres16580


} else {
ifres16576 = False


}

ifres16572 = ifres16576


} else {
ifres16572 = False


}

ifres16569 = ifres16572


} else {
ifres16569 = False


}

ifres16566 = ifres16569


} else {
ifres16566 = False


}

ifres16563 = ifres16566


} else {
ifres16563 = False


}

ifres16560 = ifres16563


} else {
ifres16560 = False


}

ifres16558 = ifres16560


} else {
ifres16558 = False


}

W486914714 := ifres16558
_ = W486914714

tmp16898 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W486914714, False)
}
__typedArg0 := W486914714
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16898 {
tmp16718 := Call(__e, PrimFunc(symshen_4unlocked_2), V4862)


var ifres16637 Obj

if True == tmp16718 {
tmp16638 := Call(__e, PrimFunc(symshen_4lazyderef), V4858, V4861)


W489014735 := tmp16638
_ = W489014735

tmp16717 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W489014735)
}
__typedArg0 := W489014735
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16639 Obj

if True == tmp16717 {
tmp16640 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W489014735)
}
__typedArg0 := W489014735
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16641 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16640, V4861)


W489114736 := tmp16641
_ = W489114736

tmp16716 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W489114736)
}
__typedArg0 := W489114736
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16642 Obj

if True == tmp16716 {
tmp16643 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W489114736)
}
__typedArg0 := W489114736
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16644 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16643, V4861)


W489214737 := tmp16644
_ = W489214737

tmp16715 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W489214737)
}
__typedArg0 := W489214737
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16645 Obj

if True == tmp16715 {
tmp16646 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W489214737)
}
__typedArg0 := W489214737
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16647 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16646, V4861)


W489314738 := tmp16647
_ = W489314738

tmp16714 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W489314738, sym_8p)
}
__typedArg0 := W489314738
__typedArg1 := sym_8p
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16648 Obj

if True == tmp16714 {
tmp16649 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W489214737)
}
__typedArg0 := W489214737
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16650 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16649, V4861)


W489414739 := tmp16650
_ = W489414739

tmp16713 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W489414739)
}
__typedArg0 := W489414739
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16651 Obj

if True == tmp16713 {
tmp16652 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W489414739)
}
__typedArg0 := W489414739
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

W489514740 := tmp16652
_ = W489514740

tmp16653 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W489414739)
}
__typedArg0 := W489414739
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16654 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16653, V4861)


W489614741 := tmp16654
_ = W489614741

tmp16712 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W489614741)
}
__typedArg0 := W489614741
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16655 Obj

if True == tmp16712 {
tmp16656 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W489614741)
}
__typedArg0 := W489614741
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

W489714742 := tmp16656
_ = W489714742

tmp16657 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W489614741)
}
__typedArg0 := W489614741
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16658 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16657, V4861)


W489814743 := tmp16658
_ = W489814743

tmp16711 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W489814743, Nil)
}
__typedArg0 := W489814743
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16659 Obj

if True == tmp16711 {
tmp16660 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W489114736)
}
__typedArg0 := W489114736
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16661 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16660, V4861)


W489914744 := tmp16661
_ = W489914744

tmp16710 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W489914744)
}
__typedArg0 := W489914744
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16662 Obj

if True == tmp16710 {
tmp16663 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W489914744)
}
__typedArg0 := W489914744
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

W490014745 := tmp16663
_ = W490014745

tmp16664 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W489914744)
}
__typedArg0 := W489914744
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16665 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16664, V4861)


W490114746 := tmp16665
_ = W490114746

tmp16709 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W490114746)
}
__typedArg0 := W490114746
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16666 Obj

if True == tmp16709 {
tmp16667 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W490114746)
}
__typedArg0 := W490114746
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16668 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16667, V4861)


W490214747 := tmp16668
_ = W490214747

tmp16708 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W490214747)
}
__typedArg0 := W490214747
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16669 Obj

if True == tmp16708 {
tmp16670 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W490214747)
}
__typedArg0 := W490214747
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

W490314748 := tmp16670
_ = W490314748

tmp16671 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W490214747)
}
__typedArg0 := W490214747
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16672 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16671, V4861)


W490414749 := tmp16672
_ = W490414749

tmp16707 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W490414749)
}
__typedArg0 := W490414749
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16673 Obj

if True == tmp16707 {
tmp16674 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W490414749)
}
__typedArg0 := W490414749
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16675 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16674, V4861)


W490514750 := tmp16675
_ = W490514750

tmp16706 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W490514750, sym_d)
}
__typedArg0 := W490514750
__typedArg1 := sym_d
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16676 Obj

if True == tmp16706 {
tmp16677 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W490414749)
}
__typedArg0 := W490414749
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16678 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16677, V4861)


W490614751 := tmp16678
_ = W490614751

tmp16705 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W490614751)
}
__typedArg0 := W490614751
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16679 Obj

if True == tmp16705 {
tmp16680 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W490614751)
}
__typedArg0 := W490614751
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

W490714752 := tmp16680
_ = W490714752

tmp16681 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W490614751)
}
__typedArg0 := W490614751
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16682 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16681, V4861)


W490814753 := tmp16682
_ = W490814753

tmp16704 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W490814753, Nil)
}
__typedArg0 := W490814753
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16683 Obj

if True == tmp16704 {
tmp16684 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W490114746)
}
__typedArg0 := W490114746
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16685 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16684, V4861)


W490914754 := tmp16685
_ = W490914754

tmp16703 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W490914754, Nil)
}
__typedArg0 := W490914754
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16686 Obj

if True == tmp16703 {
tmp16687 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W489014735)
}
__typedArg0 := W489014735
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

W491014755 := tmp16687
_ = W491014755

tmp16688 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16688

tmp16689 := Call(__e, PrimFunc(symshen_4deref), W490014745, V4861)


tmp16690 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp16691 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp16689, tmp16690)
}
__typedArg0 := tmp16689
__typedArg1 := tmp16690
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

tmp16692 := MakeNative(func(__e *ControlFlow) {
tmp16693 := MakeNative(func(__e *ControlFlow) {
tmp16694 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W490314748, Nil)
}
__typedArg0 := W490314748
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16695 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W490014745, tmp16694)
}
__typedArg0 := W490014745
__typedArg1 := tmp16694
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16696 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W489514740, tmp16695)
}
__typedArg0 := W489514740
__typedArg1 := tmp16695
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16697 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W490714752, Nil)
}
__typedArg0 := W490714752
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16698 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W490014745, tmp16697)
}
__typedArg0 := W490014745
__typedArg1 := tmp16697
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16699 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W489714742, tmp16698)
}
__typedArg0 := W489714742
__typedArg1 := tmp16698
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16700 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16699, W491014755)
}
__typedArg0 := tmp16699
__typedArg1 := W491014755
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16701 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16696, tmp16700)
}
__typedArg0 := tmp16696
__typedArg1 := tmp16700
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4l_1rules), tmp16701, V4859, True, V4861, V4862, W486514710, V4864)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V4861, V4862, W486514710, tmp16693)
return


}, 0)

tmp16702 := Call(__e, PrimFunc(symwhen), tmp16691, V4861, V4862, W486514710, tmp16692)


ifres16686 = tmp16702


} else {
ifres16686 = False


}

ifres16683 = ifres16686


} else {
ifres16683 = False


}

ifres16679 = ifres16683


} else {
ifres16679 = False


}

ifres16676 = ifres16679


} else {
ifres16676 = False


}

ifres16673 = ifres16676


} else {
ifres16673 = False


}

ifres16669 = ifres16673


} else {
ifres16669 = False


}

ifres16666 = ifres16669


} else {
ifres16666 = False


}

ifres16662 = ifres16666


} else {
ifres16662 = False


}

ifres16659 = ifres16662


} else {
ifres16659 = False


}

ifres16655 = ifres16659


} else {
ifres16655 = False


}

ifres16651 = ifres16655


} else {
ifres16651 = False


}

ifres16648 = ifres16651


} else {
ifres16648 = False


}

ifres16645 = ifres16648


} else {
ifres16645 = False


}

ifres16642 = ifres16645


} else {
ifres16642 = False


}

ifres16639 = ifres16642


} else {
ifres16639 = False


}

ifres16637 = ifres16639


} else {
ifres16637 = False


}

W488914734 := ifres16637
_ = W488914734

tmp16896 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W488914734, False)
}
__typedArg0 := W488914734
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16896 {
tmp16782 := Call(__e, PrimFunc(symshen_4unlocked_2), V4862)


var ifres16719 Obj

if True == tmp16782 {
tmp16720 := Call(__e, PrimFunc(symshen_4lazyderef), V4858, V4861)


W491214757 := tmp16720
_ = W491214757

tmp16781 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W491214757)
}
__typedArg0 := W491214757
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16721 Obj

if True == tmp16781 {
tmp16722 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W491214757)
}
__typedArg0 := W491214757
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16723 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16722, V4861)


W491314758 := tmp16723
_ = W491314758

tmp16780 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W491314758)
}
__typedArg0 := W491314758
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16724 Obj

if True == tmp16780 {
tmp16725 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W491314758)
}
__typedArg0 := W491314758
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16726 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16725, V4861)


W491414759 := tmp16726
_ = W491414759

tmp16779 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W491414759)
}
__typedArg0 := W491414759
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16727 Obj

if True == tmp16779 {
tmp16728 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W491414759)
}
__typedArg0 := W491414759
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16729 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16728, V4861)


W491514760 := tmp16729
_ = W491514760

tmp16778 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W491514760, sym_8s)
}
__typedArg0 := W491514760
__typedArg1 := sym_8s
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16730 Obj

if True == tmp16778 {
tmp16731 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W491414759)
}
__typedArg0 := W491414759
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16732 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16731, V4861)


W491614761 := tmp16732
_ = W491614761

tmp16777 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W491614761)
}
__typedArg0 := W491614761
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16733 Obj

if True == tmp16777 {
tmp16734 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W491614761)
}
__typedArg0 := W491614761
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

W491714762 := tmp16734
_ = W491714762

tmp16735 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W491614761)
}
__typedArg0 := W491614761
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16736 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16735, V4861)


W491814763 := tmp16736
_ = W491814763

tmp16776 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W491814763)
}
__typedArg0 := W491814763
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16737 Obj

if True == tmp16776 {
tmp16738 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W491814763)
}
__typedArg0 := W491814763
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

W491914764 := tmp16738
_ = W491914764

tmp16739 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W491814763)
}
__typedArg0 := W491814763
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16740 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16739, V4861)


W492014765 := tmp16740
_ = W492014765

tmp16775 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W492014765, Nil)
}
__typedArg0 := W492014765
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16741 Obj

if True == tmp16775 {
tmp16742 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W491314758)
}
__typedArg0 := W491314758
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16743 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16742, V4861)


W492114766 := tmp16743
_ = W492114766

tmp16774 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W492114766)
}
__typedArg0 := W492114766
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16744 Obj

if True == tmp16774 {
tmp16745 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W492114766)
}
__typedArg0 := W492114766
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

W492214767 := tmp16745
_ = W492214767

tmp16746 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W492114766)
}
__typedArg0 := W492114766
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16747 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16746, V4861)


W492314768 := tmp16747
_ = W492314768

tmp16773 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W492314768)
}
__typedArg0 := W492314768
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16748 Obj

if True == tmp16773 {
tmp16749 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W492314768)
}
__typedArg0 := W492314768
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16750 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16749, V4861)


W492414769 := tmp16750
_ = W492414769

tmp16772 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W492414769, symstring)
}
__typedArg0 := W492414769
__typedArg1 := symstring
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16751 Obj

if True == tmp16772 {
tmp16752 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W492314768)
}
__typedArg0 := W492314768
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16753 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16752, V4861)


W492514770 := tmp16753
_ = W492514770

tmp16771 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W492514770, Nil)
}
__typedArg0 := W492514770
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16754 Obj

if True == tmp16771 {
tmp16755 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W491214757)
}
__typedArg0 := W491214757
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

W492614771 := tmp16755
_ = W492614771

tmp16756 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16756

tmp16757 := Call(__e, PrimFunc(symshen_4deref), W492214767, V4861)


tmp16758 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp16759 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp16757, tmp16758)
}
__typedArg0 := tmp16757
__typedArg1 := tmp16758
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

tmp16760 := MakeNative(func(__e *ControlFlow) {
tmp16761 := MakeNative(func(__e *ControlFlow) {
tmp16762 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16763 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W492214767, tmp16762)
}
__typedArg0 := W492214767
__typedArg1 := tmp16762
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16764 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W491714762, tmp16763)
}
__typedArg0 := W491714762
__typedArg1 := tmp16763
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16765 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16766 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W492214767, tmp16765)
}
__typedArg0 := W492214767
__typedArg1 := tmp16765
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16767 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W491914764, tmp16766)
}
__typedArg0 := W491914764
__typedArg1 := tmp16766
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16768 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16767, W492614771)
}
__typedArg0 := tmp16767
__typedArg1 := W492614771
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16769 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16764, tmp16768)
}
__typedArg0 := tmp16764
__typedArg1 := tmp16768
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4l_1rules), tmp16769, V4859, True, V4861, V4862, W486514710, V4864)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V4861, V4862, W486514710, tmp16761)
return


}, 0)

tmp16770 := Call(__e, PrimFunc(symwhen), tmp16759, V4861, V4862, W486514710, tmp16760)


ifres16754 = tmp16770


} else {
ifres16754 = False


}

ifres16751 = ifres16754


} else {
ifres16751 = False


}

ifres16748 = ifres16751


} else {
ifres16748 = False


}

ifres16744 = ifres16748


} else {
ifres16744 = False


}

ifres16741 = ifres16744


} else {
ifres16741 = False


}

ifres16737 = ifres16741


} else {
ifres16737 = False


}

ifres16733 = ifres16737


} else {
ifres16733 = False


}

ifres16730 = ifres16733


} else {
ifres16730 = False


}

ifres16727 = ifres16730


} else {
ifres16727 = False


}

ifres16724 = ifres16727


} else {
ifres16724 = False


}

ifres16721 = ifres16724


} else {
ifres16721 = False


}

ifres16719 = ifres16721


} else {
ifres16719 = False


}

W491114756 := ifres16719
_ = W491114756

tmp16894 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W491114756, False)
}
__typedArg0 := W491114756
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16894 {
tmp16861 := Call(__e, PrimFunc(symshen_4unlocked_2), V4862)


var ifres16783 Obj

if True == tmp16861 {
tmp16784 := Call(__e, PrimFunc(symshen_4lazyderef), V4858, V4861)


W492814773 := tmp16784
_ = W492814773

tmp16860 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W492814773)
}
__typedArg0 := W492814773
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16785 Obj

if True == tmp16860 {
tmp16786 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W492814773)
}
__typedArg0 := W492814773
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16787 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16786, V4861)


W492914774 := tmp16787
_ = W492914774

tmp16859 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W492914774)
}
__typedArg0 := W492914774
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16788 Obj

if True == tmp16859 {
tmp16789 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W492914774)
}
__typedArg0 := W492914774
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16790 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16789, V4861)


W493014775 := tmp16790
_ = W493014775

tmp16858 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W493014775)
}
__typedArg0 := W493014775
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16791 Obj

if True == tmp16858 {
tmp16792 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W493014775)
}
__typedArg0 := W493014775
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16793 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16792, V4861)


W493114776 := tmp16793
_ = W493114776

tmp16857 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W493114776, sym_8v)
}
__typedArg0 := W493114776
__typedArg1 := sym_8v
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16794 Obj

if True == tmp16857 {
tmp16795 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W493014775)
}
__typedArg0 := W493014775
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16796 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16795, V4861)


W493214777 := tmp16796
_ = W493214777

tmp16856 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W493214777)
}
__typedArg0 := W493214777
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16797 Obj

if True == tmp16856 {
tmp16798 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W493214777)
}
__typedArg0 := W493214777
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

W493314778 := tmp16798
_ = W493314778

tmp16799 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W493214777)
}
__typedArg0 := W493214777
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16800 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16799, V4861)


W493414779 := tmp16800
_ = W493414779

tmp16855 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W493414779)
}
__typedArg0 := W493414779
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16801 Obj

if True == tmp16855 {
tmp16802 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W493414779)
}
__typedArg0 := W493414779
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

W493514780 := tmp16802
_ = W493514780

tmp16803 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W493414779)
}
__typedArg0 := W493414779
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16804 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16803, V4861)


W493614781 := tmp16804
_ = W493614781

tmp16854 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W493614781, Nil)
}
__typedArg0 := W493614781
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16805 Obj

if True == tmp16854 {
tmp16806 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W492914774)
}
__typedArg0 := W492914774
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16807 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16806, V4861)


W493714782 := tmp16807
_ = W493714782

tmp16853 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W493714782)
}
__typedArg0 := W493714782
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16808 Obj

if True == tmp16853 {
tmp16809 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W493714782)
}
__typedArg0 := W493714782
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

W493814783 := tmp16809
_ = W493814783

tmp16810 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W493714782)
}
__typedArg0 := W493714782
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16811 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16810, V4861)


W493914784 := tmp16811
_ = W493914784

tmp16852 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W493914784)
}
__typedArg0 := W493914784
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16812 Obj

if True == tmp16852 {
tmp16813 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W493914784)
}
__typedArg0 := W493914784
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16814 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16813, V4861)


W494014785 := tmp16814
_ = W494014785

tmp16851 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W494014785)
}
__typedArg0 := W494014785
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16815 Obj

if True == tmp16851 {
tmp16816 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W494014785)
}
__typedArg0 := W494014785
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16817 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16816, V4861)


W494114786 := tmp16817
_ = W494114786

tmp16850 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W494114786, symvector)
}
__typedArg0 := W494114786
__typedArg1 := symvector
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16818 Obj

if True == tmp16850 {
tmp16819 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W494014785)
}
__typedArg0 := W494014785
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16820 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16819, V4861)


W494214787 := tmp16820
_ = W494214787

tmp16849 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W494214787)
}
__typedArg0 := W494214787
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16821 Obj

if True == tmp16849 {
tmp16822 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W494214787)
}
__typedArg0 := W494214787
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

W494314788 := tmp16822
_ = W494314788

tmp16823 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W494214787)
}
__typedArg0 := W494214787
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16824 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16823, V4861)


W494414789 := tmp16824
_ = W494414789

tmp16848 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W494414789, Nil)
}
__typedArg0 := W494414789
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16825 Obj

if True == tmp16848 {
tmp16826 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W493914784)
}
__typedArg0 := W493914784
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16827 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16826, V4861)


W494514790 := tmp16827
_ = W494514790

tmp16847 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W494514790, Nil)
}
__typedArg0 := W494514790
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16828 Obj

if True == tmp16847 {
tmp16829 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W492814773)
}
__typedArg0 := W492814773
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

W494614791 := tmp16829
_ = W494614791

tmp16830 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16830

tmp16831 := Call(__e, PrimFunc(symshen_4deref), W493814783, V4861)


tmp16832 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp16833 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp16831, tmp16832)
}
__typedArg0 := tmp16831
__typedArg1 := tmp16832
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

tmp16834 := MakeNative(func(__e *ControlFlow) {
tmp16835 := MakeNative(func(__e *ControlFlow) {
tmp16836 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W494314788, Nil)
}
__typedArg0 := W494314788
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16837 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W493814783, tmp16836)
}
__typedArg0 := W493814783
__typedArg1 := tmp16836
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16838 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W493314778, tmp16837)
}
__typedArg0 := W493314778
__typedArg1 := tmp16837
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16839 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W494314788, Nil)
}
__typedArg0 := W494314788
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16840 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp16839)
}
__typedArg0 := symvector
__typedArg1 := tmp16839
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16841 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16840, Nil)
}
__typedArg0 := tmp16840
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16842 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W493814783, tmp16841)
}
__typedArg0 := W493814783
__typedArg1 := tmp16841
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16843 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W493514780, tmp16842)
}
__typedArg0 := W493514780
__typedArg1 := tmp16842
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16844 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16843, W494614791)
}
__typedArg0 := tmp16843
__typedArg1 := W494614791
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16845 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16838, tmp16844)
}
__typedArg0 := tmp16838
__typedArg1 := tmp16844
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4l_1rules), tmp16845, V4859, True, V4861, V4862, W486514710, V4864)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V4861, V4862, W486514710, tmp16835)
return


}, 0)

tmp16846 := Call(__e, PrimFunc(symwhen), tmp16833, V4861, V4862, W486514710, tmp16834)


ifres16828 = tmp16846


} else {
ifres16828 = False


}

ifres16825 = ifres16828


} else {
ifres16825 = False


}

ifres16821 = ifres16825


} else {
ifres16821 = False


}

ifres16818 = ifres16821


} else {
ifres16818 = False


}

ifres16815 = ifres16818


} else {
ifres16815 = False


}

ifres16812 = ifres16815


} else {
ifres16812 = False


}

ifres16808 = ifres16812


} else {
ifres16808 = False


}

ifres16805 = ifres16808


} else {
ifres16805 = False


}

ifres16801 = ifres16805


} else {
ifres16801 = False


}

ifres16797 = ifres16801


} else {
ifres16797 = False


}

ifres16794 = ifres16797


} else {
ifres16794 = False


}

ifres16791 = ifres16794


} else {
ifres16791 = False


}

ifres16788 = ifres16791


} else {
ifres16788 = False


}

ifres16785 = ifres16788


} else {
ifres16785 = False


}

ifres16783 = ifres16785


} else {
ifres16783 = False


}

W492714772 := ifres16783
_ = W492714772

tmp16892 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W492714772, False)
}
__typedArg0 := W492714772
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16892 {
tmp16888 := Call(__e, PrimFunc(symshen_4unlocked_2), V4862)


var ifres16862 Obj

if True == tmp16888 {
tmp16863 := Call(__e, PrimFunc(symshen_4lazyderef), V4858, V4861)


W494814793 := tmp16863
_ = W494814793

tmp16887 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W494814793)
}
__typedArg0 := W494814793
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16864 Obj

if True == tmp16887 {
tmp16865 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W494814793)
}
__typedArg0 := W494814793
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

W494914794 := tmp16865
_ = W494914794

tmp16866 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W494814793)
}
__typedArg0 := W494814793
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

W495014795 := tmp16866
_ = W495014795

tmp16867 := Call(__e, PrimFunc(symshen_4lazyderef), V4859, V4861)


W495114796 := tmp16867
_ = W495114796

tmp16868 := MakeNative(func(__e *ControlFlow) {
Z4953 := __e.Get(1)
_ = Z4953
__e.Return(MakeNative(func(__e *ControlFlow) {
Z4954 := __e.Get(1)
_ = Z4954
tmp16869 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16869

tmp16870 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4l_1rules), W495014795, Z4954, V4860, V4861, V4862, W486514710, V4864)
return
}, 0)

__e.TailApply(PrimFunc(symbind), Z4953, W494914794, V4861, V4862, W486514710, tmp16870)
return


}, 1))
return
}, 1)

W495214797 := tmp16868
_ = W495214797

tmp16886 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W495114796)
}
__typedArg0 := W495114796
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16871 Obj

if True == tmp16886 {
tmp16872 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W495114796)
}
__typedArg0 := W495114796
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

W495514798 := tmp16872
_ = W495514798

tmp16873 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W495114796)
}
__typedArg0 := W495114796
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

W495614799 := tmp16873
_ = W495614799

tmp16874 := Call(__e, W495214797, W495514798)


tmp16875 := Call(__e, tmp16874, W495614799)


ifres16871 = tmp16875


} else {
tmp16885 := Call(__e, PrimFunc(symshen_4pvar_2), W495114796)


var ifres16876 Obj

if True == tmp16885 {
tmp16877 := Call(__e, PrimFunc(symshen_4newpv), V4861)


W495714800 := tmp16877
_ = W495714800

tmp16878 := Call(__e, PrimFunc(symshen_4newpv), V4861)


W495814801 := tmp16878
_ = W495814801

tmp16879 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W495714800, W495814801)
}
__typedArg0 := W495714800
__typedArg1 := W495814801
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16880 := MakeNative(func(__e *ControlFlow) {
tmp16881 := Call(__e, W495214797, W495714800)


__e.TailApply(tmp16881, W495814801)
return


}, 0)

tmp16882 := Call(__e, PrimFunc(symshen_4bind_b), W495114796, tmp16879, V4861, tmp16880)


tmp16883 := Call(__e, PrimFunc(symshen_4gc), V4861, tmp16882)


tmp16884 := Call(__e, PrimFunc(symshen_4gc), V4861, tmp16883)


ifres16876 = tmp16884


} else {
ifres16876 = False


}

ifres16871 = ifres16876


}

ifres16864 = ifres16871


} else {
ifres16864 = False


}

ifres16862 = ifres16864


} else {
ifres16862 = False


}

W494714792 := ifres16862
_ = W494714792

tmp16890 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W494714792, False)
}
__typedArg0 := W494714792
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16890 {
__e.TailApply(PrimFunc(symshen_4unlock), V4862, W486514710)
return
} else {
__e.Return(W494714792)
return
}


} else {
__e.Return(W492714772)
return
}


} else {
__e.Return(W491114756)
return
}


} else {
__e.Return(W488914734)
return
}


} else {
__e.Return(W486914714)
return
}


} else {
__e.Return(W486614711)
return
}


}, 7)

tmp16901 := Call(__e, ns2_1set, symshen_4l_1rules, tmp16545)


_ = tmp16901

tmp16902 := MakeNative(func(__e *ControlFlow) {
V4959 := __e.Get(1)
_ = V4959
V4960 := __e.Get(2)
_ = V4960
V4961 := __e.Get(3)
_ = V4961
V4962 := __e.Get(4)
_ = V4962
V4963 := __e.Get(5)
_ = V4963
V4964 := __e.Get(6)
_ = V4964
W496514802 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V4963)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V4963
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
_ = W496514802

tmp16942 := Call(__e, PrimFunc(symshen_4unlocked_2), V4962)


var ifres16904 Obj

if True == tmp16942 {
tmp16905 := Call(__e, PrimFunc(symshen_4lazyderef), V4959, V4961)


W496714804 := tmp16905
_ = W496714804

tmp16941 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W496714804)
}
__typedArg0 := W496714804
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16906 Obj

if True == tmp16941 {
tmp16907 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W496714804)
}
__typedArg0 := W496714804
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16908 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16907, V4961)


W496814805 := tmp16908
_ = W496814805

tmp16940 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W496814805, symdefine)
}
__typedArg0 := W496814805
__typedArg1 := symdefine
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16909 Obj

if True == tmp16940 {
tmp16910 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W496714804)
}
__typedArg0 := W496714804
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16911 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16910, V4961)


W496914806 := tmp16911
_ = W496914806

tmp16939 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W496914806)
}
__typedArg0 := W496914806
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16912 Obj

if True == tmp16939 {
tmp16913 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W496914806)
}
__typedArg0 := W496914806
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

W497014807 := tmp16913
_ = W497014807

tmp16914 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W496914806)
}
__typedArg0 := W496914806
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

W497114808 := tmp16914
_ = W497114808

tmp16915 := Call(__e, PrimFunc(symshen_4newpv), V4961)


W497214809 := tmp16915
_ = W497214809

tmp16916 := Call(__e, PrimFunc(symshen_4newpv), V4961)


W497314810 := tmp16916
_ = W497314810

tmp16917 := Call(__e, PrimFunc(symshen_4newpv), V4961)


W497414811 := tmp16917
_ = W497414811

tmp16918 := Call(__e, PrimFunc(symshen_4newpv), V4961)


W497514812 := tmp16918
_ = W497514812

tmp16919 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16919

tmp16920 := MakeNative(func(__e *ControlFlow) {
tmp16921 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W497014807, W497114808)
}
__typedArg0 := W497014807
__typedArg1 := W497114808
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16922 := Call(__e, PrimFunc(symshen_4sigxrules), tmp16921)


tmp16923 := MakeNative(func(__e *ControlFlow) {
tmp16924 := Call(__e, PrimFunc(symshen_4lazyderef), W497214809, V4961)


tmp16925 := Call(__e, PrimFunc(symfst), tmp16924)


tmp16926 := MakeNative(func(__e *ControlFlow) {
tmp16927 := Call(__e, PrimFunc(symshen_4lazyderef), W497214809, V4961)


tmp16928 := Call(__e, PrimFunc(symsnd), tmp16927)


tmp16929 := MakeNative(func(__e *ControlFlow) {
tmp16930 := Call(__e, PrimFunc(symshen_4deref), W497514812, V4961)


tmp16931 := Call(__e, PrimFunc(symshen_4freshen_1sig), tmp16930)


tmp16932 := MakeNative(func(__e *ControlFlow) {
tmp16933 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symis), W497514812, V4960, V4961, V4962, W496514802, V4964)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4t_d_1rules), W497014807, W497314810, W497414811, MakeNumber(1), V4961, V4962, W496514802, tmp16933)
return


}, 0)

__e.TailApply(PrimFunc(symbind), W497414811, tmp16931, V4961, V4962, W496514802, tmp16932)
return


}, 0)

__e.TailApply(PrimFunc(symbind), W497314810, tmp16928, V4961, V4962, W496514802, tmp16929)
return


}, 0)

__e.TailApply(PrimFunc(symbind), W497514812, tmp16925, V4961, V4962, W496514802, tmp16926)
return


}, 0)

__e.TailApply(PrimFunc(symbind), W497214809, tmp16922, V4961, V4962, W496514802, tmp16923)
return


}, 0)

tmp16934 := Call(__e, PrimFunc(symshen_4cut), V4961, V4962, W496514802, tmp16920)


tmp16935 := Call(__e, PrimFunc(symshen_4gc), V4961, tmp16934)


tmp16936 := Call(__e, PrimFunc(symshen_4gc), V4961, tmp16935)


tmp16937 := Call(__e, PrimFunc(symshen_4gc), V4961, tmp16936)


tmp16938 := Call(__e, PrimFunc(symshen_4gc), V4961, tmp16937)


ifres16912 = tmp16938


} else {
ifres16912 = False


}

ifres16909 = ifres16912


} else {
ifres16909 = False


}

ifres16906 = ifres16909


} else {
ifres16906 = False


}

ifres16904 = ifres16906


} else {
ifres16904 = False


}

W496614803 := ifres16904
_ = W496614803

tmp16944 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W496614803, False)
}
__typedArg0 := W496614803
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16944 {
__e.TailApply(PrimFunc(symshen_4unlock), V4962, W496514802)
return
} else {
__e.Return(W496614803)
return
}


}, 6)

tmp16945 := Call(__e, ns2_1set, symshen_4t_d, tmp16902)


_ = tmp16945

tmp16946 := MakeNative(func(__e *ControlFlow) {
V4976 := __e.Get(1)
_ = V4976
tmp16947 := MakeNative(func(__e *ControlFlow) {
Z4977 := __e.Get(1)
_ = Z4977
__e.TailApply(PrimFunc(symshen_4_5sig_drules_6), Z4977)
return
}, 1)

__e.TailApply(PrimFunc(symcompile), tmp16947, V4976)
return


}, 1)

tmp16948 := Call(__e, ns2_1set, symshen_4sigxrules, tmp16946)


_ = tmp16948

tmp16949 := MakeNative(func(__e *ControlFlow) {
V4978 := __e.Get(1)
_ = V4978
tmp16976 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4978)
}
__typedArg0 := V4978
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16950 Obj

if True == tmp16976 {
tmp16951 := Call(__e, PrimFunc(symtail), V4978)


W498014814 := tmp16951
_ = W498014814

tmp16974 := Call(__e, PrimFunc(symshen_4hds_a_2), W498014814, sym_i)


var ifres16952 Obj

if True == tmp16974 {
tmp16953 := Call(__e, PrimFunc(symtail), W498014814)


W498114815 := tmp16953
_ = W498114815

tmp16954 := Call(__e, PrimFunc(symshen_4_5signature_6), W498114815)


W498214816 := tmp16954
_ = W498214816

tmp16972 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W498214816)


var ifres16955 Obj

if True == tmp16972 {
tmp16956 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16955 = tmp16956


} else {
tmp16957 := Call(__e, PrimFunc(symshen_4_5_1out), W498214816)


W498314817 := tmp16957
_ = W498314817

tmp16958 := Call(__e, PrimFunc(symshen_4in_1_6), W498214816)


W498414818 := tmp16958
_ = W498414818

tmp16971 := Call(__e, PrimFunc(symshen_4hds_a_2), W498414818, sym_j)


var ifres16959 Obj

if True == tmp16971 {
tmp16960 := Call(__e, PrimFunc(symtail), W498414818)


W498514819 := tmp16960
_ = W498514819

tmp16961 := Call(__e, PrimFunc(symshen_4_5rules_d_6), W498514819)


W498614820 := tmp16961
_ = W498614820

tmp16969 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W498614820)


var ifres16962 Obj

if True == tmp16969 {
tmp16963 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16962 = tmp16963


} else {
tmp16964 := Call(__e, PrimFunc(symshen_4_5_1out), W498614820)


W498714821 := tmp16964
_ = W498714821

tmp16965 := Call(__e, PrimFunc(symshen_4in_1_6), W498614820)


W498814822 := tmp16965
_ = W498814822

tmp16966 := Call(__e, PrimFunc(symshen_4rectify_1type), W498314817)


W498914823 := tmp16966
_ = W498914823

tmp16967 := Call(__e, PrimFunc(sym_8p), W498914823, W498714821)


tmp16968 := Call(__e, PrimFunc(symshen_4comb), W498814822, tmp16967)


ifres16962 = tmp16968


}

ifres16959 = ifres16962


} else {
tmp16970 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16959 = tmp16970


}

ifres16955 = ifres16959


}

ifres16952 = ifres16955


} else {
tmp16973 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16952 = tmp16973


}

ifres16950 = ifres16952


} else {
tmp16975 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16950 = tmp16975


}

W497914813 := ifres16950
_ = W497914813

tmp16978 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W497914813)


if True == tmp16978 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W497914813)
return
}


}, 1)

tmp16979 := Call(__e, ns2_1set, symshen_4_5sig_drules_6, tmp16949)


_ = tmp16979

tmp16980 := MakeNative(func(__e *ControlFlow) {
V4990 := __e.Get(1)
_ = V4990
tmp16981 := Call(__e, PrimFunc(symshen_4extract_1vars), V4990)


W499114824 := tmp16981
_ = W499114824

tmp16982 := MakeNative(func(__e *ControlFlow) {
Z4993 := __e.Get(1)
_ = Z4993
tmp16983 := Call(__e, PrimFunc(symconcat), sym_e, Z4993)


tmp16984 := Call(__e, PrimFunc(symshen_4freshterm), tmp16983)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Z4993, tmp16984)
}
__typedArg0 := Z4993
__typedArg1 := tmp16984
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp16985 := Call(__e, PrimFunc(symmap), tmp16982, W499114824)


W499214825 := tmp16985
_ = W499214825

__e.TailApply(PrimFunc(symshen_4freshen_1type), W499214825, V4990)
return


}, 1)

tmp16986 := Call(__e, ns2_1set, symshen_4freshen_1sig, tmp16980)


_ = tmp16986

tmp16987 := MakeNative(func(__e *ControlFlow) {
V4994 := __e.Get(1)
_ = V4994
V4995 := __e.Get(2)
_ = V4995
tmp17001 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V4994)
}
__typedArg0 := Nil
__typedArg1 := V4994
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17001 {
__e.Return(V4995)
return
} else {
tmp16999 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4994)
}
__typedArg0 := V4994
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16995 Obj

if True == tmp16999 {
tmp16997 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4994)
}
__typedArg0 := V4994
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16998 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp16997)
}
__typedArg0 := tmp16997
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16996 Obj

if True == tmp16998 {
ifres16996 = True


} else {
ifres16996 = False


}

ifres16995 = ifres16996


} else {
ifres16995 = False


}

if True == ifres16995 {
tmp16988 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4994)
}
__typedArg0 := V4994
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16989 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4994)
}
__typedArg0 := V4994
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16990 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16989)
}
__typedArg0 := tmp16989
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16991 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4994)
}
__typedArg0 := V4994
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16992 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp16991)
}
__typedArg0 := tmp16991
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16993 := Call(__e, PrimFunc(symsubst), tmp16990, tmp16992, V4995)


__e.TailApply(PrimFunc(symshen_4freshen_1type), tmp16988, tmp16993)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("partial function shen.freshen-type"))
}
__typedArg0 := MakeString("partial function shen.freshen-type")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 2)

tmp17002 := Call(__e, ns2_1set, symshen_4freshen_1type, tmp16987)


_ = tmp17002

tmp17003 := MakeNative(func(__e *ControlFlow) {
V4996 := __e.Get(1)
_ = V4996
tmp17004 := Call(__e, PrimFunc(symshen_4_5rule_d_6), V4996)


W499814827 := tmp17004
_ = W499814827

tmp17017 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W499814827)


var ifres17005 Obj

if True == tmp17017 {
tmp17006 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17005 = tmp17006


} else {
tmp17007 := Call(__e, PrimFunc(symshen_4_5_1out), W499814827)


W499914828 := tmp17007
_ = W499914828

tmp17008 := Call(__e, PrimFunc(symshen_4in_1_6), W499814827)


W500014829 := tmp17008
_ = W500014829

tmp17009 := Call(__e, PrimFunc(symshen_4_5rules_d_6), W500014829)


W500114830 := tmp17009
_ = W500114830

tmp17016 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W500114830)


var ifres17010 Obj

if True == tmp17016 {
tmp17011 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17010 = tmp17011


} else {
tmp17012 := Call(__e, PrimFunc(symshen_4_5_1out), W500114830)


W500214831 := tmp17012
_ = W500214831

tmp17013 := Call(__e, PrimFunc(symshen_4in_1_6), W500114830)


W500314832 := tmp17013
_ = W500314832

tmp17014 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W499914828, W500214831)
}
__typedArg0 := W499914828
__typedArg1 := W500214831
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17015 := Call(__e, PrimFunc(symshen_4comb), W500314832, tmp17014)


ifres17010 = tmp17015


}

ifres17005 = ifres17010


}

W499714826 := ifres17005
_ = W499714826

tmp17029 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W499714826)


if True == tmp17029 {
tmp17018 := Call(__e, PrimFunc(symshen_4_5rule_d_6), V4996)


W500514834 := tmp17018
_ = W500514834

tmp17025 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W500514834)


var ifres17019 Obj

if True == tmp17025 {
tmp17020 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17019 = tmp17020


} else {
tmp17021 := Call(__e, PrimFunc(symshen_4_5_1out), W500514834)


W500614835 := tmp17021
_ = W500614835

tmp17022 := Call(__e, PrimFunc(symshen_4in_1_6), W500514834)


W500714836 := tmp17022
_ = W500714836

tmp17023 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W500614835, Nil)
}
__typedArg0 := W500614835
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17024 := Call(__e, PrimFunc(symshen_4comb), W500714836, tmp17023)


ifres17019 = tmp17024


}

W500414833 := ifres17019
_ = W500414833

tmp17027 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W500414833)


if True == tmp17027 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W500414833)
return
}


} else {
__e.Return(W499714826)
return
}


}, 1)

tmp17030 := Call(__e, ns2_1set, symshen_4_5rules_d_6, tmp17003)


_ = tmp17030

tmp17031 := MakeNative(func(__e *ControlFlow) {
V5008 := __e.Get(1)
_ = V5008
tmp17032 := Call(__e, PrimFunc(symshen_4_5patterns_6), V5008)


W501014838 := tmp17032
_ = W501014838

tmp17060 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W501014838)


var ifres17033 Obj

if True == tmp17060 {
tmp17034 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17033 = tmp17034


} else {
tmp17035 := Call(__e, PrimFunc(symshen_4_5_1out), W501014838)


W501114839 := tmp17035
_ = W501114839

tmp17036 := Call(__e, PrimFunc(symshen_4in_1_6), W501014838)


W501214840 := tmp17036
_ = W501214840

tmp17059 := Call(__e, PrimFunc(symshen_4hds_a_2), W501214840, sym_1_6)


var ifres17037 Obj

if True == tmp17059 {
tmp17038 := Call(__e, PrimFunc(symtail), W501214840)


W501314841 := tmp17038
_ = W501314841

tmp17057 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W501314841)
}
__typedArg0 := W501314841
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17039 Obj

if True == tmp17057 {
tmp17040 := Call(__e, PrimFunc(symhead), W501314841)


W501414842 := tmp17040
_ = W501414842

tmp17041 := Call(__e, PrimFunc(symtail), W501314841)


W501514843 := tmp17041
_ = W501514843

tmp17055 := Call(__e, PrimFunc(symshen_4hds_a_2), W501514843, symwhere)


var ifres17042 Obj

if True == tmp17055 {
tmp17043 := Call(__e, PrimFunc(symtail), W501514843)


W501614844 := tmp17043
_ = W501614844

tmp17053 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W501614844)
}
__typedArg0 := W501614844
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17044 Obj

if True == tmp17053 {
tmp17045 := Call(__e, PrimFunc(symhead), W501614844)


W501714845 := tmp17045
_ = W501714845

tmp17046 := Call(__e, PrimFunc(symtail), W501614844)


W501814846 := tmp17046
_ = W501814846

tmp17047 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W501414842, Nil)
}
__typedArg0 := W501414842
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17048 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W501714845, tmp17047)
}
__typedArg0 := W501714845
__typedArg1 := tmp17047
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17049 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhere, tmp17048)
}
__typedArg0 := symwhere
__typedArg1 := tmp17048
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17050 := Call(__e, PrimFunc(sym_8p), W501114839, tmp17049)


tmp17051 := Call(__e, PrimFunc(symshen_4comb), W501814846, tmp17050)


ifres17044 = tmp17051


} else {
tmp17052 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17044 = tmp17052


}

ifres17042 = ifres17044


} else {
tmp17054 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17042 = tmp17054


}

ifres17039 = ifres17042


} else {
tmp17056 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17039 = tmp17056


}

ifres17037 = ifres17039


} else {
tmp17058 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17037 = tmp17058


}

ifres17033 = ifres17037


}

W500914837 := ifres17033
_ = W500914837

tmp17133 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W500914837)


if True == tmp17133 {
tmp17061 := Call(__e, PrimFunc(symshen_4_5patterns_6), V5008)


W502014848 := tmp17061
_ = W502014848

tmp17090 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W502014848)


var ifres17062 Obj

if True == tmp17090 {
tmp17063 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17062 = tmp17063


} else {
tmp17064 := Call(__e, PrimFunc(symshen_4_5_1out), W502014848)


W502114849 := tmp17064
_ = W502114849

tmp17065 := Call(__e, PrimFunc(symshen_4in_1_6), W502014848)


W502214850 := tmp17065
_ = W502214850

tmp17089 := Call(__e, PrimFunc(symshen_4hds_a_2), W502214850, sym_5_1)


var ifres17066 Obj

if True == tmp17089 {
tmp17067 := Call(__e, PrimFunc(symtail), W502214850)


W502314851 := tmp17067
_ = W502314851

tmp17087 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W502314851)
}
__typedArg0 := W502314851
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17068 Obj

if True == tmp17087 {
tmp17069 := Call(__e, PrimFunc(symhead), W502314851)


W502414852 := tmp17069
_ = W502414852

tmp17070 := Call(__e, PrimFunc(symtail), W502314851)


W502514853 := tmp17070
_ = W502514853

tmp17085 := Call(__e, PrimFunc(symshen_4hds_a_2), W502514853, symwhere)


var ifres17071 Obj

if True == tmp17085 {
tmp17072 := Call(__e, PrimFunc(symtail), W502514853)


W502614854 := tmp17072
_ = W502614854

tmp17083 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W502614854)
}
__typedArg0 := W502614854
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17073 Obj

if True == tmp17083 {
tmp17074 := Call(__e, PrimFunc(symhead), W502614854)


W502714855 := tmp17074
_ = W502714855

tmp17075 := Call(__e, PrimFunc(symtail), W502614854)


W502814856 := tmp17075
_ = W502814856

tmp17076 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W502414852, Nil)
}
__typedArg0 := W502414852
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17077 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W502714855, tmp17076)
}
__typedArg0 := W502714855
__typedArg1 := tmp17076
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17078 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhere, tmp17077)
}
__typedArg0 := symwhere
__typedArg1 := tmp17077
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17079 := Call(__e, PrimFunc(symshen_4correct), tmp17078)


tmp17080 := Call(__e, PrimFunc(sym_8p), W502114849, tmp17079)


tmp17081 := Call(__e, PrimFunc(symshen_4comb), W502814856, tmp17080)


ifres17073 = tmp17081


} else {
tmp17082 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17073 = tmp17082


}

ifres17071 = ifres17073


} else {
tmp17084 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17071 = tmp17084


}

ifres17068 = ifres17071


} else {
tmp17086 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17068 = tmp17086


}

ifres17066 = ifres17068


} else {
tmp17088 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17066 = tmp17088


}

ifres17062 = ifres17066


}

W501914847 := ifres17062
_ = W501914847

tmp17131 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W501914847)


if True == tmp17131 {
tmp17091 := Call(__e, PrimFunc(symshen_4_5patterns_6), V5008)


W503014858 := tmp17091
_ = W503014858

tmp17108 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W503014858)


var ifres17092 Obj

if True == tmp17108 {
tmp17093 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17092 = tmp17093


} else {
tmp17094 := Call(__e, PrimFunc(symshen_4_5_1out), W503014858)


W503114859 := tmp17094
_ = W503114859

tmp17095 := Call(__e, PrimFunc(symshen_4in_1_6), W503014858)


W503214860 := tmp17095
_ = W503214860

tmp17107 := Call(__e, PrimFunc(symshen_4hds_a_2), W503214860, sym_5_1)


var ifres17096 Obj

if True == tmp17107 {
tmp17097 := Call(__e, PrimFunc(symtail), W503214860)


W503314861 := tmp17097
_ = W503314861

tmp17105 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W503314861)
}
__typedArg0 := W503314861
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17098 Obj

if True == tmp17105 {
tmp17099 := Call(__e, PrimFunc(symhead), W503314861)


W503414862 := tmp17099
_ = W503414862

tmp17100 := Call(__e, PrimFunc(symtail), W503314861)


W503514863 := tmp17100
_ = W503514863

tmp17101 := Call(__e, PrimFunc(symshen_4correct), W503414862)


tmp17102 := Call(__e, PrimFunc(sym_8p), W503114859, tmp17101)


tmp17103 := Call(__e, PrimFunc(symshen_4comb), W503514863, tmp17102)


ifres17098 = tmp17103


} else {
tmp17104 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17098 = tmp17104


}

ifres17096 = ifres17098


} else {
tmp17106 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17096 = tmp17106


}

ifres17092 = ifres17096


}

W502914857 := ifres17092
_ = W502914857

tmp17129 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W502914857)


if True == tmp17129 {
tmp17109 := Call(__e, PrimFunc(symshen_4_5patterns_6), V5008)


W503714865 := tmp17109
_ = W503714865

tmp17125 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W503714865)


var ifres17110 Obj

if True == tmp17125 {
tmp17111 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17110 = tmp17111


} else {
tmp17112 := Call(__e, PrimFunc(symshen_4_5_1out), W503714865)


W503814866 := tmp17112
_ = W503814866

tmp17113 := Call(__e, PrimFunc(symshen_4in_1_6), W503714865)


W503914867 := tmp17113
_ = W503914867

tmp17124 := Call(__e, PrimFunc(symshen_4hds_a_2), W503914867, sym_1_6)


var ifres17114 Obj

if True == tmp17124 {
tmp17115 := Call(__e, PrimFunc(symtail), W503914867)


W504014868 := tmp17115
_ = W504014868

tmp17122 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W504014868)
}
__typedArg0 := W504014868
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17116 Obj

if True == tmp17122 {
tmp17117 := Call(__e, PrimFunc(symhead), W504014868)


W504114869 := tmp17117
_ = W504114869

tmp17118 := Call(__e, PrimFunc(symtail), W504014868)


W504214870 := tmp17118
_ = W504214870

tmp17119 := Call(__e, PrimFunc(sym_8p), W503814866, W504114869)


tmp17120 := Call(__e, PrimFunc(symshen_4comb), W504214870, tmp17119)


ifres17116 = tmp17120


} else {
tmp17121 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17116 = tmp17121


}

ifres17114 = ifres17116


} else {
tmp17123 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17114 = tmp17123


}

ifres17110 = ifres17114


}

W503614864 := ifres17110
_ = W503614864

tmp17127 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W503614864)


if True == tmp17127 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W503614864)
return
}


} else {
__e.Return(W502914857)
return
}


} else {
__e.Return(W501914847)
return
}


} else {
__e.Return(W500914837)
return
}


}, 1)

tmp17134 := Call(__e, ns2_1set, symshen_4_5rule_d_6, tmp17031)


_ = tmp17134

tmp17135 := MakeNative(func(__e *ControlFlow) {
V5043 := __e.Get(1)
_ = V5043
tmp17283 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5043)
}
__typedArg0 := V5043
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17227 Obj

if True == tmp17283 {
tmp17281 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5043)
}
__typedArg0 := V5043
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17282 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symwhere, tmp17281)
}
__typedArg0 := symwhere
__typedArg1 := tmp17281
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17229 Obj

if True == tmp17282 {
tmp17279 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5043)
}
__typedArg0 := V5043
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17280 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp17279)
}
__typedArg0 := tmp17279
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17231 Obj

if True == tmp17280 {
tmp17276 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5043)
}
__typedArg0 := V5043
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17277 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp17276)
}
__typedArg0 := tmp17276
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17278 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp17277)
}
__typedArg0 := tmp17277
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17233 Obj

if True == tmp17278 {
tmp17272 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5043)
}
__typedArg0 := V5043
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17273 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp17272)
}
__typedArg0 := tmp17272
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17274 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp17273)
}
__typedArg0 := tmp17273
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17275 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp17274)
}
__typedArg0 := tmp17274
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17235 Obj

if True == tmp17275 {
tmp17267 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5043)
}
__typedArg0 := V5043
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17268 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp17267)
}
__typedArg0 := tmp17267
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17269 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp17268)
}
__typedArg0 := tmp17268
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17270 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp17269)
}
__typedArg0 := tmp17269
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17271 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symfail_1if, tmp17270)
}
__typedArg0 := symfail_1if
__typedArg1 := tmp17270
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17237 Obj

if True == tmp17271 {
tmp17262 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5043)
}
__typedArg0 := V5043
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17263 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp17262)
}
__typedArg0 := tmp17262
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17264 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp17263)
}
__typedArg0 := tmp17263
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17265 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp17264)
}
__typedArg0 := tmp17264
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17266 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp17265)
}
__typedArg0 := tmp17265
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17239 Obj

if True == tmp17266 {
tmp17256 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5043)
}
__typedArg0 := V5043
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17257 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp17256)
}
__typedArg0 := tmp17256
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17258 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp17257)
}
__typedArg0 := tmp17257
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17259 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp17258)
}
__typedArg0 := tmp17258
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17260 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp17259)
}
__typedArg0 := tmp17259
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17261 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp17260)
}
__typedArg0 := tmp17260
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17241 Obj

if True == tmp17261 {
tmp17249 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5043)
}
__typedArg0 := V5043
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17250 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp17249)
}
__typedArg0 := tmp17249
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17251 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp17250)
}
__typedArg0 := tmp17250
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17252 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp17251)
}
__typedArg0 := tmp17251
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17253 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp17252)
}
__typedArg0 := tmp17252
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17254 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp17253)
}
__typedArg0 := tmp17253
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17255 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp17254)
}
__typedArg0 := Nil
__typedArg1 := tmp17254
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17243 Obj

if True == tmp17255 {
tmp17245 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5043)
}
__typedArg0 := V5043
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17246 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp17245)
}
__typedArg0 := tmp17245
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17247 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp17246)
}
__typedArg0 := tmp17246
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17248 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp17247)
}
__typedArg0 := Nil
__typedArg1 := tmp17247
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17244 Obj

if True == tmp17248 {
ifres17244 = True


} else {
ifres17244 = False


}

ifres17243 = ifres17244


} else {
ifres17243 = False


}

var ifres17242 Obj

if True == ifres17243 {
ifres17242 = True


} else {
ifres17242 = False


}

ifres17241 = ifres17242


} else {
ifres17241 = False


}

var ifres17240 Obj

if True == ifres17241 {
ifres17240 = True


} else {
ifres17240 = False


}

ifres17239 = ifres17240


} else {
ifres17239 = False


}

var ifres17238 Obj

if True == ifres17239 {
ifres17238 = True


} else {
ifres17238 = False


}

ifres17237 = ifres17238


} else {
ifres17237 = False


}

var ifres17236 Obj

if True == ifres17237 {
ifres17236 = True


} else {
ifres17236 = False


}

ifres17235 = ifres17236


} else {
ifres17235 = False


}

var ifres17234 Obj

if True == ifres17235 {
ifres17234 = True


} else {
ifres17234 = False


}

ifres17233 = ifres17234


} else {
ifres17233 = False


}

var ifres17232 Obj

if True == ifres17233 {
ifres17232 = True


} else {
ifres17232 = False


}

ifres17231 = ifres17232


} else {
ifres17231 = False


}

var ifres17230 Obj

if True == ifres17231 {
ifres17230 = True


} else {
ifres17230 = False


}

ifres17229 = ifres17230


} else {
ifres17229 = False


}

var ifres17228 Obj

if True == ifres17229 {
ifres17228 = True


} else {
ifres17228 = False


}

ifres17227 = ifres17228


} else {
ifres17227 = False


}

if True == ifres17227 {
tmp17136 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5043)
}
__typedArg0 := V5043
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17137 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp17136)
}
__typedArg0 := tmp17136
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17138 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5043)
}
__typedArg0 := V5043
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17139 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp17138)
}
__typedArg0 := tmp17138
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17140 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp17139)
}
__typedArg0 := tmp17139
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17141 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp17140)
}
__typedArg0 := tmp17140
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17142 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17141, Nil)
}
__typedArg0 := tmp17141
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17143 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnot, tmp17142)
}
__typedArg0 := symnot
__typedArg1 := tmp17142
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17144 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17143, Nil)
}
__typedArg0 := tmp17143
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17145 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17137, tmp17144)
}
__typedArg0 := tmp17137
__typedArg1 := tmp17144
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17146 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symand, tmp17145)
}
__typedArg0 := symand
__typedArg1 := tmp17145
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17147 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5043)
}
__typedArg0 := V5043
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17148 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp17147)
}
__typedArg0 := tmp17147
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17149 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp17148)
}
__typedArg0 := tmp17148
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17150 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp17149)
}
__typedArg0 := tmp17149
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17151 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp17150)
}
__typedArg0 := tmp17150
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17152 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17146, tmp17151)
}
__typedArg0 := tmp17146
__typedArg1 := tmp17151
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhere, tmp17152)
}
__typedArg0 := symwhere
__typedArg1 := tmp17152
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp17225 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5043)
}
__typedArg0 := V5043
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17206 Obj

if True == tmp17225 {
tmp17223 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5043)
}
__typedArg0 := V5043
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17224 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symwhere, tmp17223)
}
__typedArg0 := symwhere
__typedArg1 := tmp17223
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17208 Obj

if True == tmp17224 {
tmp17221 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5043)
}
__typedArg0 := V5043
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17222 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp17221)
}
__typedArg0 := tmp17221
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17210 Obj

if True == tmp17222 {
tmp17218 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5043)
}
__typedArg0 := V5043
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17219 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp17218)
}
__typedArg0 := tmp17218
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17220 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp17219)
}
__typedArg0 := tmp17219
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17212 Obj

if True == tmp17220 {
tmp17214 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5043)
}
__typedArg0 := V5043
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17215 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp17214)
}
__typedArg0 := tmp17214
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17216 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp17215)
}
__typedArg0 := tmp17215
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17217 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp17216)
}
__typedArg0 := Nil
__typedArg1 := tmp17216
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17213 Obj

if True == tmp17217 {
ifres17213 = True


} else {
ifres17213 = False


}

ifres17212 = ifres17213


} else {
ifres17212 = False


}

var ifres17211 Obj

if True == ifres17212 {
ifres17211 = True


} else {
ifres17211 = False


}

ifres17210 = ifres17211


} else {
ifres17210 = False


}

var ifres17209 Obj

if True == ifres17210 {
ifres17209 = True


} else {
ifres17209 = False


}

ifres17208 = ifres17209


} else {
ifres17208 = False


}

var ifres17207 Obj

if True == ifres17208 {
ifres17207 = True


} else {
ifres17207 = False


}

ifres17206 = ifres17207


} else {
ifres17206 = False


}

if True == ifres17206 {
tmp17153 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5043)
}
__typedArg0 := V5043
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17154 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp17153)
}
__typedArg0 := tmp17153
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17155 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5043)
}
__typedArg0 := V5043
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17156 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp17155)
}
__typedArg0 := tmp17155
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17157 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp17156)
}
__typedArg0 := tmp17156
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17158 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfail, Nil)
}
__typedArg0 := symfail
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17159 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17158, Nil)
}
__typedArg0 := tmp17158
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17160 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17157, tmp17159)
}
__typedArg0 := tmp17157
__typedArg1 := tmp17159
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17161 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_a_a, tmp17160)
}
__typedArg0 := sym_a_a
__typedArg1 := tmp17160
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17162 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17161, Nil)
}
__typedArg0 := tmp17161
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17163 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnot, tmp17162)
}
__typedArg0 := symnot
__typedArg1 := tmp17162
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17164 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17163, Nil)
}
__typedArg0 := tmp17163
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17165 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17154, tmp17164)
}
__typedArg0 := tmp17154
__typedArg1 := tmp17164
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17166 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symand, tmp17165)
}
__typedArg0 := symand
__typedArg1 := tmp17165
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17167 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5043)
}
__typedArg0 := V5043
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17168 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp17167)
}
__typedArg0 := tmp17167
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17169 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17166, tmp17168)
}
__typedArg0 := tmp17166
__typedArg1 := tmp17168
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhere, tmp17169)
}
__typedArg0 := symwhere
__typedArg1 := tmp17169
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp17204 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5043)
}
__typedArg0 := V5043
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17185 Obj

if True == tmp17204 {
tmp17202 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5043)
}
__typedArg0 := V5043
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17203 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symfail_1if, tmp17202)
}
__typedArg0 := symfail_1if
__typedArg1 := tmp17202
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17187 Obj

if True == tmp17203 {
tmp17200 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5043)
}
__typedArg0 := V5043
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17201 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp17200)
}
__typedArg0 := tmp17200
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17189 Obj

if True == tmp17201 {
tmp17197 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5043)
}
__typedArg0 := V5043
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17198 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp17197)
}
__typedArg0 := tmp17197
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17199 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp17198)
}
__typedArg0 := tmp17198
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17191 Obj

if True == tmp17199 {
tmp17193 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5043)
}
__typedArg0 := V5043
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17194 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp17193)
}
__typedArg0 := tmp17193
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17195 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp17194)
}
__typedArg0 := tmp17194
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17196 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp17195)
}
__typedArg0 := Nil
__typedArg1 := tmp17195
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17192 Obj

if True == tmp17196 {
ifres17192 = True


} else {
ifres17192 = False


}

ifres17191 = ifres17192


} else {
ifres17191 = False


}

var ifres17190 Obj

if True == ifres17191 {
ifres17190 = True


} else {
ifres17190 = False


}

ifres17189 = ifres17190


} else {
ifres17189 = False


}

var ifres17188 Obj

if True == ifres17189 {
ifres17188 = True


} else {
ifres17188 = False


}

ifres17187 = ifres17188


} else {
ifres17187 = False


}

var ifres17186 Obj

if True == ifres17187 {
ifres17186 = True


} else {
ifres17186 = False


}

ifres17185 = ifres17186


} else {
ifres17185 = False


}

if True == ifres17185 {
tmp17170 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5043)
}
__typedArg0 := V5043
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17171 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17170, Nil)
}
__typedArg0 := tmp17170
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17172 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnot, tmp17171)
}
__typedArg0 := symnot
__typedArg1 := tmp17171
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17173 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5043)
}
__typedArg0 := V5043
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17174 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp17173)
}
__typedArg0 := tmp17173
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17175 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17172, tmp17174)
}
__typedArg0 := tmp17172
__typedArg1 := tmp17174
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhere, tmp17175)
}
__typedArg0 := symwhere
__typedArg1 := tmp17175
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp17176 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfail, Nil)
}
__typedArg0 := symfail
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17177 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17176, Nil)
}
__typedArg0 := tmp17176
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17178 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5043, tmp17177)
}
__typedArg0 := V5043
__typedArg1 := tmp17177
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17179 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_a_a, tmp17178)
}
__typedArg0 := sym_a_a
__typedArg1 := tmp17178
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17180 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17179, Nil)
}
__typedArg0 := tmp17179
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17181 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnot, tmp17180)
}
__typedArg0 := symnot
__typedArg1 := tmp17180
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17182 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5043, Nil)
}
__typedArg0 := V5043
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17183 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17181, tmp17182)
}
__typedArg0 := tmp17181
__typedArg1 := tmp17182
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhere, tmp17183)
}
__typedArg0 := symwhere
__typedArg1 := tmp17183
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}


}


}


}, 1)

tmp17284 := Call(__e, ns2_1set, symshen_4correct, tmp17135)


_ = tmp17284

tmp17285 := MakeNative(func(__e *ControlFlow) {
V5044 := __e.Get(1)
_ = V5044
V5045 := __e.Get(2)
_ = V5045
V5046 := __e.Get(3)
_ = V5046
V5047 := __e.Get(4)
_ = V5047
V5048 := __e.Get(5)
_ = V5048
V5049 := __e.Get(6)
_ = V5049
V5050 := __e.Get(7)
_ = V5050
V5051 := __e.Get(8)
_ = V5051
W505214871 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V5050)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V5050
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
_ = W505214871

tmp17293 := Call(__e, PrimFunc(symshen_4unlocked_2), V5049)


var ifres17287 Obj

if True == tmp17293 {
tmp17288 := Call(__e, PrimFunc(symshen_4lazyderef), V5045, V5048)


W505414873 := tmp17288
_ = W505414873

tmp17292 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W505414873, Nil)
}
__typedArg0 := W505414873
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17289 Obj

if True == tmp17292 {
tmp17290 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp17290

tmp17291 := Call(__e, PrimFunc(symthaw), V5051)


ifres17289 = tmp17291


} else {
ifres17289 = False


}

ifres17287 = ifres17289


} else {
ifres17287 = False


}

W505314872 := ifres17287
_ = W505314872

tmp17318 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W505314872, False)
}
__typedArg0 := W505314872
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17318 {
tmp17314 := Call(__e, PrimFunc(symshen_4unlocked_2), V5049)


var ifres17294 Obj

if True == tmp17314 {
tmp17295 := Call(__e, PrimFunc(symshen_4lazyderef), V5045, V5048)


W505614875 := tmp17295
_ = W505614875

tmp17313 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W505614875)
}
__typedArg0 := W505614875
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17296 Obj

if True == tmp17313 {
tmp17297 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W505614875)
}
__typedArg0 := W505614875
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

W505714876 := tmp17297
_ = W505714876

tmp17298 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W505614875)
}
__typedArg0 := W505614875
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

W505814877 := tmp17298
_ = W505814877

tmp17299 := Call(__e, PrimFunc(symshen_4newpv), V5048)


W505914878 := tmp17299
_ = W505914878

tmp17300 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp17300

tmp17301 := Call(__e, PrimFunc(symshen_4deref), W505714876, V5048)


tmp17302 := Call(__e, PrimFunc(symshen_4freshen_1rule), tmp17301)


tmp17303 := MakeNative(func(__e *ControlFlow) {
tmp17304 := Call(__e, PrimFunc(symshen_4lazyderef), W505914878, V5048)


tmp17305 := Call(__e, PrimFunc(symfst), tmp17304)


tmp17306 := Call(__e, PrimFunc(symshen_4lazyderef), W505914878, V5048)


tmp17307 := Call(__e, PrimFunc(symsnd), tmp17306)


tmp17308 := MakeNative(func(__e *ControlFlow) {
tmp17309 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4t_d_1rules), V5044, W505814877, V5046, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V5047)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V5047
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})(), V5048, V5049, W505214871, V5051)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V5048, V5049, W505214871, tmp17309)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4t_d_1rule), V5044, V5047, tmp17305, tmp17307, V5046, V5048, V5049, W505214871, tmp17308)
return


}, 0)

tmp17311 := Call(__e, PrimFunc(symbind), W505914878, tmp17302, V5048, V5049, W505214871, tmp17303)


tmp17312 := Call(__e, PrimFunc(symshen_4gc), V5048, tmp17311)


ifres17296 = tmp17312


} else {
ifres17296 = False


}

ifres17294 = ifres17296


} else {
ifres17294 = False


}

W505514874 := ifres17294
_ = W505514874

tmp17316 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W505514874, False)
}
__typedArg0 := W505514874
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17316 {
__e.TailApply(PrimFunc(symshen_4unlock), V5049, W505214871)
return
} else {
__e.Return(W505514874)
return
}


} else {
__e.Return(W505314872)
return
}


}, 8)

tmp17319 := Call(__e, ns2_1set, symshen_4t_d_1rules, tmp17285)


_ = tmp17319

tmp17320 := MakeNative(func(__e *ControlFlow) {
V5060 := __e.Get(1)
_ = V5060
tmp17331 := Call(__e, PrimFunc(symtuple_2), V5060)


if True == tmp17331 {
tmp17321 := Call(__e, PrimFunc(symfst), V5060)


tmp17322 := Call(__e, PrimFunc(symshen_4extract_1vars), tmp17321)


W506114879 := tmp17322
_ = W506114879

tmp17323 := MakeNative(func(__e *ControlFlow) {
Z5063 := __e.Get(1)
_ = Z5063
tmp17324 := Call(__e, PrimFunc(symshen_4freshterm), Z5063)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Z5063, tmp17324)
}
__typedArg0 := Z5063
__typedArg1 := tmp17324
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp17325 := Call(__e, PrimFunc(symmap), tmp17323, W506114879)


W506214880 := tmp17325
_ = W506214880

tmp17326 := Call(__e, PrimFunc(symfst), V5060)


tmp17327 := Call(__e, PrimFunc(symshen_4freshen), W506214880, tmp17326)


tmp17328 := Call(__e, PrimFunc(symsnd), V5060)


tmp17329 := Call(__e, PrimFunc(symshen_4freshen), W506214880, tmp17328)


__e.TailApply(PrimFunc(sym_8p), tmp17327, tmp17329)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("partial function shen.freshen-rule"))
}
__typedArg0 := MakeString("partial function shen.freshen-rule")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}, 1)

tmp17332 := Call(__e, ns2_1set, symshen_4freshen_1rule, tmp17320)


_ = tmp17332

tmp17333 := MakeNative(func(__e *ControlFlow) {
V5064 := __e.Get(1)
_ = V5064
V5065 := __e.Get(2)
_ = V5065
tmp17347 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V5064)
}
__typedArg0 := Nil
__typedArg1 := V5064
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17347 {
__e.Return(V5065)
return
} else {
tmp17345 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5064)
}
__typedArg0 := V5064
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17341 Obj

if True == tmp17345 {
tmp17343 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5064)
}
__typedArg0 := V5064
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17344 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp17343)
}
__typedArg0 := tmp17343
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17342 Obj

if True == tmp17344 {
ifres17342 = True


} else {
ifres17342 = False


}

ifres17341 = ifres17342


} else {
ifres17341 = False


}

if True == ifres17341 {
tmp17334 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5064)
}
__typedArg0 := V5064
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17335 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5064)
}
__typedArg0 := V5064
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17336 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp17335)
}
__typedArg0 := tmp17335
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17337 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5064)
}
__typedArg0 := V5064
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17338 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp17337)
}
__typedArg0 := tmp17337
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17339 := Call(__e, PrimFunc(symshen_4beta), tmp17336, tmp17338, V5065)


__e.TailApply(PrimFunc(symshen_4freshen), tmp17334, tmp17339)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("partial function shen.freshen"))
}
__typedArg0 := MakeString("partial function shen.freshen")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 2)

tmp17348 := Call(__e, ns2_1set, symshen_4freshen, tmp17333)


_ = tmp17348

tmp17349 := MakeNative(func(__e *ControlFlow) {
V5066 := __e.Get(1)
_ = V5066
V5067 := __e.Get(2)
_ = V5067
V5068 := __e.Get(3)
_ = V5068
V5069 := __e.Get(4)
_ = V5069
V5070 := __e.Get(5)
_ = V5070
V5071 := __e.Get(6)
_ = V5071
V5072 := __e.Get(7)
_ = V5072
V5073 := __e.Get(8)
_ = V5073
V5074 := __e.Get(9)
_ = V5074
tmp17353 := Call(__e, PrimFunc(symshen_4unlocked_2), V5072)


var ifres17350 Obj

if True == tmp17353 {
tmp17351 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp17351

tmp17352 := Call(__e, PrimFunc(symshen_4t_d_1rule_1h), V5068, V5069, V5070, V5071, V5072, V5073, V5074)


ifres17350 = tmp17352


} else {
ifres17350 = False


}

W507514881 := ifres17350
_ = W507514881

tmp17365 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W507514881, False)
}
__typedArg0 := W507514881
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17365 {
tmp17363 := Call(__e, PrimFunc(symshen_4unlocked_2), V5072)


if True == tmp17363 {
tmp17354 := Call(__e, PrimFunc(symshen_4newpv), V5071)


W507614882 := tmp17354
_ = W507614882

tmp17355 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp17355

tmp17356 := Call(__e, PrimFunc(symshen_4app), V5066, MakeString("\n"), symshen_4a)


tmp17358 := Call(__e, PrimFunc(symshen_4app), V5067, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(" of "))
__typedS1, __typedOK1 := TypedString(tmp17356)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(" of ")
__typedArg1 := tmp17356
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


tmp17360 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("type error in rule "))
__typedS1, __typedOK1 := TypedString(tmp17358)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("type error in rule ")
__typedArg1 := tmp17358
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("type error in rule "))
__typedS1, __typedOK1 := TypedString(tmp17358)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("type error in rule ")
__typedArg1 := tmp17358
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})()

tmp17361 := Call(__e, PrimFunc(symbind), W507614882, tmp17360, V5071, V5072, V5073, V5074)


__e.TailApply(PrimFunc(symshen_4gc), V5071, tmp17361)
return


} else {
__e.Return(False)
return
}


} else {
__e.Return(W507514881)
return
}


}, 9)

tmp17366 := Call(__e, ns2_1set, symshen_4t_d_1rule, tmp17349)


_ = tmp17366

tmp17367 := MakeNative(func(__e *ControlFlow) {
V5077 := __e.Get(1)
_ = V5077
V5078 := __e.Get(2)
_ = V5078
V5079 := __e.Get(3)
_ = V5079
V5080 := __e.Get(4)
_ = V5080
V5081 := __e.Get(5)
_ = V5081
V5082 := __e.Get(6)
_ = V5082
V5083 := __e.Get(7)
_ = V5083
W508414883 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V5082)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V5082
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
_ = W508414883

tmp17392 := Call(__e, PrimFunc(symshen_4unlocked_2), V5081)


var ifres17369 Obj

if True == tmp17392 {
tmp17370 := Call(__e, PrimFunc(symshen_4lazyderef), V5077, V5080)


W508614885 := tmp17370
_ = W508614885

tmp17391 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W508614885, Nil)
}
__typedArg0 := W508614885
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17371 Obj

if True == tmp17391 {
tmp17372 := Call(__e, PrimFunc(symshen_4lazyderef), V5079, V5080)


W508714886 := tmp17372
_ = W508714886

tmp17390 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W508714886)
}
__typedArg0 := W508714886
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17373 Obj

if True == tmp17390 {
tmp17374 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W508714886)
}
__typedArg0 := W508714886
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17375 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17374, V5080)


W508814887 := tmp17375
_ = W508814887

tmp17389 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W508814887, sym_1_1_6)
}
__typedArg0 := W508814887
__typedArg1 := sym_1_1_6
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17376 Obj

if True == tmp17389 {
tmp17377 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W508714886)
}
__typedArg0 := W508714886
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17378 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17377, V5080)


W508914888 := tmp17378
_ = W508914888

tmp17388 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W508914888)
}
__typedArg0 := W508914888
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17379 Obj

if True == tmp17388 {
tmp17380 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W508914888)
}
__typedArg0 := W508914888
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

W509014889 := tmp17380
_ = W509014889

tmp17381 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W508914888)
}
__typedArg0 := W508914888
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17382 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17381, V5080)


W509114890 := tmp17382
_ = W509114890

tmp17387 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W509114890, Nil)
}
__typedArg0 := W509114890
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17383 Obj

if True == tmp17387 {
tmp17384 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp17384

tmp17385 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4t_d_1correct), V5078, W509014889, Nil, V5080, V5081, W508414883, V5083)
return
}, 0)

tmp17386 := Call(__e, PrimFunc(symshen_4cut), V5080, V5081, W508414883, tmp17385)


ifres17383 = tmp17386


} else {
ifres17383 = False


}

ifres17379 = ifres17383


} else {
ifres17379 = False


}

ifres17376 = ifres17379


} else {
ifres17376 = False


}

ifres17373 = ifres17376


} else {
ifres17373 = False


}

ifres17371 = ifres17373


} else {
ifres17371 = False


}

ifres17369 = ifres17371


} else {
ifres17369 = False


}

W508514884 := ifres17369
_ = W508514884

tmp17411 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W508514884, False)
}
__typedArg0 := W508514884
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17411 {
tmp17407 := Call(__e, PrimFunc(symshen_4unlocked_2), V5081)


var ifres17393 Obj

if True == tmp17407 {
tmp17394 := Call(__e, PrimFunc(symshen_4newpv), V5080)


W509314892 := tmp17394
_ = W509314892

tmp17395 := Call(__e, PrimFunc(symshen_4newpv), V5080)


W509414893 := tmp17395
_ = W509414893

tmp17396 := Call(__e, PrimFunc(symshen_4newpv), V5080)


W509514894 := tmp17396
_ = W509514894

tmp17397 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp17397

tmp17398 := Call(__e, PrimFunc(symshen_4freshterms), V5077)


tmp17399 := MakeNative(func(__e *ControlFlow) {
tmp17400 := MakeNative(func(__e *ControlFlow) {
tmp17401 := MakeNative(func(__e *ControlFlow) {
tmp17402 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4t_d_1correct), V5078, W509414893, W509514894, V5080, V5081, W508414883, V5083)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4myassume), V5077, V5079, W509514894, V5080, V5081, W508414883, tmp17402)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V5080, V5081, W508414883, tmp17401)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4t_d_1integrity), V5077, V5079, W509314892, W509414893, V5080, V5081, W508414883, tmp17400)
return


}, 0)

tmp17403 := Call(__e, PrimFunc(symshen_4p_1hyps), tmp17398, W509314892, V5080, V5081, W508414883, tmp17399)


tmp17404 := Call(__e, PrimFunc(symshen_4gc), V5080, tmp17403)


tmp17405 := Call(__e, PrimFunc(symshen_4gc), V5080, tmp17404)


tmp17406 := Call(__e, PrimFunc(symshen_4gc), V5080, tmp17405)


ifres17393 = tmp17406


} else {
ifres17393 = False


}

W509214891 := ifres17393
_ = W509214891

tmp17409 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W509214891, False)
}
__typedArg0 := W509214891
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17409 {
__e.TailApply(PrimFunc(symshen_4unlock), V5081, W508414883)
return
} else {
__e.Return(W509214891)
return
}


} else {
__e.Return(W508514884)
return
}


}, 7)

tmp17412 := Call(__e, ns2_1set, symshen_4t_d_1rule_1h, tmp17367)


_ = tmp17412

tmp17413 := MakeNative(func(__e *ControlFlow) {
V5096 := __e.Get(1)
_ = V5096
V5097 := __e.Get(2)
_ = V5097
V5098 := __e.Get(3)
_ = V5098
V5099 := __e.Get(4)
_ = V5099
V5100 := __e.Get(5)
_ = V5100
V5101 := __e.Get(6)
_ = V5101
V5102 := __e.Get(7)
_ = V5102
tmp17427 := Call(__e, PrimFunc(symshen_4unlocked_2), V5100)


var ifres17414 Obj

if True == tmp17427 {
tmp17415 := Call(__e, PrimFunc(symshen_4lazyderef), V5096, V5099)


W510414896 := tmp17415
_ = W510414896

tmp17426 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W510414896, Nil)
}
__typedArg0 := W510414896
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17416 Obj

if True == tmp17426 {
tmp17417 := Call(__e, PrimFunc(symshen_4lazyderef), V5098, V5099)


W510514897 := tmp17417
_ = W510514897

tmp17418 := MakeNative(func(__e *ControlFlow) {
tmp17419 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp17419

__e.TailApply(PrimFunc(symthaw), V5102)
return


}, 0)

W510614898 := tmp17418
_ = W510614898

tmp17425 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W510514897, Nil)
}
__typedArg0 := W510514897
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17420 Obj

if True == tmp17425 {
tmp17421 := Call(__e, PrimFunc(symthaw), W510614898)


ifres17420 = tmp17421


} else {
tmp17424 := Call(__e, PrimFunc(symshen_4pvar_2), W510514897)


var ifres17422 Obj

if True == tmp17424 {
tmp17423 := Call(__e, PrimFunc(symshen_4bind_b), W510514897, Nil, V5099, W510614898)


ifres17422 = tmp17423


} else {
ifres17422 = False


}

ifres17420 = ifres17422


}

ifres17416 = ifres17420


} else {
ifres17416 = False


}

ifres17414 = ifres17416


} else {
ifres17414 = False


}

W510314895 := ifres17414
_ = W510314895

tmp17546 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W510314895, False)
}
__typedArg0 := W510314895
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17546 {
tmp17544 := Call(__e, PrimFunc(symshen_4unlocked_2), V5100)


if True == tmp17544 {
tmp17428 := Call(__e, PrimFunc(symshen_4lazyderef), V5096, V5099)


W510714899 := tmp17428
_ = W510714899

tmp17542 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W510714899)
}
__typedArg0 := W510714899
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17542 {
tmp17429 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W510714899)
}
__typedArg0 := W510714899
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

W510814900 := tmp17429
_ = W510814900

tmp17430 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W510714899)
}
__typedArg0 := W510714899
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

W510914901 := tmp17430
_ = W510914901

tmp17431 := Call(__e, PrimFunc(symshen_4lazyderef), V5097, V5099)


W511014902 := tmp17431
_ = W511014902

tmp17540 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W511014902)
}
__typedArg0 := W511014902
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17540 {
tmp17432 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W511014902)
}
__typedArg0 := W511014902
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

W511114903 := tmp17432
_ = W511114903

tmp17433 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W511014902)
}
__typedArg0 := W511014902
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17434 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17433, V5099)


W511214904 := tmp17434
_ = W511214904

tmp17538 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W511214904)
}
__typedArg0 := W511214904
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17538 {
tmp17435 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W511214904)
}
__typedArg0 := W511214904
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17436 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17435, V5099)


W511314905 := tmp17436
_ = W511314905

tmp17536 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W511314905, sym_1_1_6)
}
__typedArg0 := W511314905
__typedArg1 := sym_1_1_6
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17536 {
tmp17437 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W511214904)
}
__typedArg0 := W511214904
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17438 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17437, V5099)


W511414906 := tmp17438
_ = W511414906

tmp17534 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W511414906)
}
__typedArg0 := W511414906
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17534 {
tmp17439 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W511414906)
}
__typedArg0 := W511414906
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

W511514907 := tmp17439
_ = W511514907

tmp17440 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W511414906)
}
__typedArg0 := W511414906
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17441 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17440, V5099)


W511614908 := tmp17441
_ = W511614908

tmp17532 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W511614908, Nil)
}
__typedArg0 := W511614908
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17532 {
tmp17442 := Call(__e, PrimFunc(symshen_4lazyderef), V5098, V5099)


W511714909 := tmp17442
_ = W511714909

tmp17443 := MakeNative(func(__e *ControlFlow) {
Z5119 := __e.Get(1)
_ = Z5119
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5120 := __e.Get(1)
_ = Z5120
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5121 := __e.Get(1)
_ = Z5121
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5122 := __e.Get(1)
_ = Z5122
tmp17444 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp17444

tmp17445 := MakeNative(func(__e *ControlFlow) {
tmp17446 := MakeNative(func(__e *ControlFlow) {
tmp17447 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp17448 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4myassume), W510914901, W511514907, Z5122, V5099, V5100, V5101, V5102)
return
}, 0)

__e.TailApply(PrimFunc(symbind), Z5120, tmp17447, V5099, V5100, V5101, tmp17448)
return


}, 0)

__e.TailApply(PrimFunc(symis_b), W510814900, Z5119, V5099, V5100, V5101, tmp17446)
return


}, 0)

__e.TailApply(PrimFunc(symis_b), W511114903, Z5121, V5099, V5100, V5101, tmp17445)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

W511814910 := tmp17443
_ = W511814910

tmp17530 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W511714909)
}
__typedArg0 := W511714909
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17530 {
tmp17449 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W511714909)
}
__typedArg0 := W511714909
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17450 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17449, V5099)


W512314911 := tmp17450
_ = W512314911

tmp17451 := MakeNative(func(__e *ControlFlow) {
Z5125 := __e.Get(1)
_ = Z5125
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5126 := __e.Get(1)
_ = Z5126
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5127 := __e.Get(1)
_ = Z5127
tmp17452 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W511714909)
}
__typedArg0 := W511714909
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

W512814913 := tmp17452
_ = W512814913

tmp17453 := Call(__e, W511814910, Z5125)


tmp17454 := Call(__e, tmp17453, Z5126)


tmp17455 := Call(__e, tmp17454, Z5127)


__e.TailApply(tmp17455, W512814913)
return


}, 1))
return
}, 1))
return
}, 1)

W512414912 := tmp17451
_ = W512414912

tmp17510 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W512314911)
}
__typedArg0 := W512314911
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17510 {
tmp17456 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W512314911)
}
__typedArg0 := W512314911
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

W512914914 := tmp17456
_ = W512914914

tmp17457 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W512314911)
}
__typedArg0 := W512314911
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17458 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17457, V5099)


W513014915 := tmp17458
_ = W513014915

tmp17459 := MakeNative(func(__e *ControlFlow) {
Z5132 := __e.Get(1)
_ = Z5132
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5133 := __e.Get(1)
_ = Z5133
tmp17460 := Call(__e, W512414912, W512914914)


tmp17461 := Call(__e, tmp17460, Z5132)


__e.TailApply(tmp17461, Z5133)
return


}, 1))
return
}, 1)

W513114916 := tmp17459
_ = W513114916

tmp17494 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W513014915)
}
__typedArg0 := W513014915
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17494 {
tmp17462 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W513014915)
}
__typedArg0 := W513014915
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

W513414917 := tmp17462
_ = W513414917

tmp17463 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W513014915)
}
__typedArg0 := W513014915
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17464 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17463, V5099)


W513514918 := tmp17464
_ = W513514918

tmp17465 := MakeNative(func(__e *ControlFlow) {
Z5137 := __e.Get(1)
_ = Z5137
tmp17466 := Call(__e, W513114916, W513414917)


__e.TailApply(tmp17466, Z5137)
return


}, 1)

W513614919 := tmp17465
_ = W513614919

tmp17482 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W513514918)
}
__typedArg0 := W513514918
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17482 {
tmp17467 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W513514918)
}
__typedArg0 := W513514918
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

W513814920 := tmp17467
_ = W513814920

tmp17468 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W513514918)
}
__typedArg0 := W513514918
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17469 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17468, V5099)


W513914921 := tmp17469
_ = W513914921

tmp17470 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W513614919, W513814920)
return
}, 0)

W514014922 := tmp17470
_ = W514014922

tmp17474 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W513914921, Nil)
}
__typedArg0 := W513914921
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17474 {
__e.TailApply(PrimFunc(symthaw), W514014922)
return
} else {
tmp17472 := Call(__e, PrimFunc(symshen_4pvar_2), W513914921)


if True == tmp17472 {
__e.TailApply(PrimFunc(symshen_4bind_b), W513914921, Nil, V5099, W514014922)
return
} else {
__e.Return(False)
return
}


}


} else {
tmp17480 := Call(__e, PrimFunc(symshen_4pvar_2), W513514918)


if True == tmp17480 {
tmp17475 := Call(__e, PrimFunc(symshen_4newpv), V5099)


W514114923 := tmp17475
_ = W514114923

tmp17476 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W514114923, Nil)
}
__typedArg0 := W514114923
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17477 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W513614919, W514114923)
return
}, 0)

tmp17478 := Call(__e, PrimFunc(symshen_4bind_b), W513514918, tmp17476, V5099, tmp17477)


__e.TailApply(PrimFunc(symshen_4gc), V5099, tmp17478)
return


} else {
__e.Return(False)
return
}


}


} else {
tmp17492 := Call(__e, PrimFunc(symshen_4pvar_2), W513014915)


if True == tmp17492 {
tmp17483 := Call(__e, PrimFunc(symshen_4newpv), V5099)


W514214924 := tmp17483
_ = W514214924

tmp17484 := Call(__e, PrimFunc(symshen_4newpv), V5099)


W514314925 := tmp17484
_ = W514314925

tmp17485 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W514314925, Nil)
}
__typedArg0 := W514314925
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17486 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W514214924, tmp17485)
}
__typedArg0 := W514214924
__typedArg1 := tmp17485
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17487 := MakeNative(func(__e *ControlFlow) {
tmp17488 := Call(__e, W513114916, W514214924)


__e.TailApply(tmp17488, W514314925)
return


}, 0)

tmp17489 := Call(__e, PrimFunc(symshen_4bind_b), W513014915, tmp17486, V5099, tmp17487)


tmp17490 := Call(__e, PrimFunc(symshen_4gc), V5099, tmp17489)


__e.TailApply(PrimFunc(symshen_4gc), V5099, tmp17490)
return


} else {
__e.Return(False)
return
}


}


} else {
tmp17508 := Call(__e, PrimFunc(symshen_4pvar_2), W512314911)


if True == tmp17508 {
tmp17495 := Call(__e, PrimFunc(symshen_4newpv), V5099)


W514414926 := tmp17495
_ = W514414926

tmp17496 := Call(__e, PrimFunc(symshen_4newpv), V5099)


W514514927 := tmp17496
_ = W514514927

tmp17497 := Call(__e, PrimFunc(symshen_4newpv), V5099)


W514614928 := tmp17497
_ = W514614928

tmp17498 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W514614928, Nil)
}
__typedArg0 := W514614928
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17499 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W514514927, tmp17498)
}
__typedArg0 := W514514927
__typedArg1 := tmp17498
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17500 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W514414926, tmp17499)
}
__typedArg0 := W514414926
__typedArg1 := tmp17499
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17501 := MakeNative(func(__e *ControlFlow) {
tmp17502 := Call(__e, W512414912, W514414926)


tmp17503 := Call(__e, tmp17502, W514514927)


__e.TailApply(tmp17503, W514614928)
return


}, 0)

tmp17504 := Call(__e, PrimFunc(symshen_4bind_b), W512314911, tmp17500, V5099, tmp17501)


tmp17505 := Call(__e, PrimFunc(symshen_4gc), V5099, tmp17504)


tmp17506 := Call(__e, PrimFunc(symshen_4gc), V5099, tmp17505)


__e.TailApply(PrimFunc(symshen_4gc), V5099, tmp17506)
return


} else {
__e.Return(False)
return
}


}


} else {
tmp17528 := Call(__e, PrimFunc(symshen_4pvar_2), W511714909)


if True == tmp17528 {
tmp17511 := Call(__e, PrimFunc(symshen_4newpv), V5099)


W514714929 := tmp17511
_ = W514714929

tmp17512 := Call(__e, PrimFunc(symshen_4newpv), V5099)


W514814930 := tmp17512
_ = W514814930

tmp17513 := Call(__e, PrimFunc(symshen_4newpv), V5099)


W514914931 := tmp17513
_ = W514914931

tmp17514 := Call(__e, PrimFunc(symshen_4newpv), V5099)


W515014932 := tmp17514
_ = W515014932

tmp17515 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W514914931, Nil)
}
__typedArg0 := W514914931
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17516 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W514814930, tmp17515)
}
__typedArg0 := W514814930
__typedArg1 := tmp17515
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17517 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W514714929, tmp17516)
}
__typedArg0 := W514714929
__typedArg1 := tmp17516
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17518 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17517, W515014932)
}
__typedArg0 := tmp17517
__typedArg1 := W515014932
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17519 := MakeNative(func(__e *ControlFlow) {
tmp17520 := Call(__e, W511814910, W514714929)


tmp17521 := Call(__e, tmp17520, W514814930)


tmp17522 := Call(__e, tmp17521, W514914931)


__e.TailApply(tmp17522, W515014932)
return


}, 0)

tmp17523 := Call(__e, PrimFunc(symshen_4bind_b), W511714909, tmp17518, V5099, tmp17519)


tmp17524 := Call(__e, PrimFunc(symshen_4gc), V5099, tmp17523)


tmp17525 := Call(__e, PrimFunc(symshen_4gc), V5099, tmp17524)


tmp17526 := Call(__e, PrimFunc(symshen_4gc), V5099, tmp17525)


__e.TailApply(PrimFunc(symshen_4gc), V5099, tmp17526)
return


} else {
__e.Return(False)
return
}


}


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


} else {
__e.Return(False)
return
}


} else {
__e.Return(W510314895)
return
}


}, 7)

tmp17547 := Call(__e, ns2_1set, symshen_4myassume, tmp17413)


_ = tmp17547

tmp17548 := MakeNative(func(__e *ControlFlow) {
V5153 := __e.Get(1)
_ = V5153
tmp17571 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V5153)
}
__typedArg0 := Nil
__typedArg1 := V5153
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17571 {
__e.Return(Nil)
return
} else {
tmp17569 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5153)
}
__typedArg0 := V5153
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17565 Obj

if True == tmp17569 {
tmp17567 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5153)
}
__typedArg0 := V5153
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17568 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp17567)
}
__typedArg0 := tmp17567
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17566 Obj

if True == tmp17568 {
ifres17566 = True


} else {
ifres17566 = False


}

ifres17565 = ifres17566


} else {
ifres17565 = False


}

if True == ifres17565 {
tmp17549 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5153)
}
__typedArg0 := V5153
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17550 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5153)
}
__typedArg0 := V5153
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17551 := Call(__e, PrimFunc(symappend), tmp17549, tmp17550)


__e.TailApply(PrimFunc(symshen_4freshterms), tmp17551)
return


} else {
tmp17563 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5153)
}
__typedArg0 := V5153
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17559 Obj

if True == tmp17563 {
tmp17561 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5153)
}
__typedArg0 := V5153
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17562 := Call(__e, PrimFunc(symshen_4freshterm_2), tmp17561)


var ifres17560 Obj

if True == tmp17562 {
ifres17560 = True


} else {
ifres17560 = False


}

ifres17559 = ifres17560


} else {
ifres17559 = False


}

if True == ifres17559 {
tmp17552 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5153)
}
__typedArg0 := V5153
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17553 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5153)
}
__typedArg0 := V5153
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17554 := Call(__e, PrimFunc(symshen_4freshterms), tmp17553)


__e.TailApply(PrimFunc(symadjoin), tmp17552, tmp17554)
return


} else {
tmp17557 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5153)
}
__typedArg0 := V5153
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17557 {
tmp17555 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5153)
}
__typedArg0 := V5153
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4freshterms), tmp17555)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("partial function shen.freshterms"))
}
__typedArg0 := MakeString("partial function shen.freshterms")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}


}


}, 1)

tmp17572 := Call(__e, ns2_1set, symshen_4freshterms, tmp17548)


_ = tmp17572

tmp17573 := MakeNative(func(__e *ControlFlow) {
V5154 := __e.Get(1)
_ = V5154
V5155 := __e.Get(2)
_ = V5155
V5156 := __e.Get(3)
_ = V5156
V5157 := __e.Get(4)
_ = V5157
V5158 := __e.Get(5)
_ = V5158
V5159 := __e.Get(6)
_ = V5159
tmp17587 := Call(__e, PrimFunc(symshen_4unlocked_2), V5157)


var ifres17574 Obj

if True == tmp17587 {
tmp17575 := Call(__e, PrimFunc(symshen_4lazyderef), V5154, V5156)


W516114934 := tmp17575
_ = W516114934

tmp17586 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W516114934, Nil)
}
__typedArg0 := W516114934
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17576 Obj

if True == tmp17586 {
tmp17577 := Call(__e, PrimFunc(symshen_4lazyderef), V5155, V5156)


W516214935 := tmp17577
_ = W516214935

tmp17578 := MakeNative(func(__e *ControlFlow) {
tmp17579 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp17579

__e.TailApply(PrimFunc(symthaw), V5159)
return


}, 0)

W516314936 := tmp17578
_ = W516314936

tmp17585 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W516214935, Nil)
}
__typedArg0 := W516214935
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17580 Obj

if True == tmp17585 {
tmp17581 := Call(__e, PrimFunc(symthaw), W516314936)


ifres17580 = tmp17581


} else {
tmp17584 := Call(__e, PrimFunc(symshen_4pvar_2), W516214935)


var ifres17582 Obj

if True == tmp17584 {
tmp17583 := Call(__e, PrimFunc(symshen_4bind_b), W516214935, Nil, V5156, W516314936)


ifres17582 = tmp17583


} else {
ifres17582 = False


}

ifres17580 = ifres17582


}

ifres17576 = ifres17580


} else {
ifres17576 = False


}

ifres17574 = ifres17576


} else {
ifres17574 = False


}

W516014933 := ifres17574
_ = W516014933

tmp17684 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W516014933, False)
}
__typedArg0 := W516014933
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17684 {
tmp17682 := Call(__e, PrimFunc(symshen_4unlocked_2), V5157)


if True == tmp17682 {
tmp17588 := Call(__e, PrimFunc(symshen_4lazyderef), V5154, V5156)


W516414937 := tmp17588
_ = W516414937

tmp17680 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W516414937)
}
__typedArg0 := W516414937
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17680 {
tmp17589 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W516414937)
}
__typedArg0 := W516414937
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

W516514938 := tmp17589
_ = W516514938

tmp17590 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W516414937)
}
__typedArg0 := W516414937
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

W516614939 := tmp17590
_ = W516614939

tmp17591 := Call(__e, PrimFunc(symshen_4lazyderef), V5155, V5156)


W516714940 := tmp17591
_ = W516714940

tmp17592 := MakeNative(func(__e *ControlFlow) {
Z5169 := __e.Get(1)
_ = Z5169
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5170 := __e.Get(1)
_ = Z5170
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5171 := __e.Get(1)
_ = Z5171
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5172 := __e.Get(1)
_ = Z5172
tmp17593 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp17593

tmp17594 := MakeNative(func(__e *ControlFlow) {
tmp17595 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp17596 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4p_1hyps), W516614939, Z5172, V5156, V5157, V5158, V5159)
return
}, 0)

__e.TailApply(PrimFunc(symbind), Z5170, tmp17595, V5156, V5157, V5158, tmp17596)
return


}, 0)

__e.TailApply(PrimFunc(symbind), Z5169, W516514938, V5156, V5157, V5158, tmp17594)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

W516814941 := tmp17592
_ = W516814941

tmp17678 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W516714940)
}
__typedArg0 := W516714940
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17678 {
tmp17597 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W516714940)
}
__typedArg0 := W516714940
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17598 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17597, V5156)


W517314942 := tmp17598
_ = W517314942

tmp17599 := MakeNative(func(__e *ControlFlow) {
Z5175 := __e.Get(1)
_ = Z5175
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5176 := __e.Get(1)
_ = Z5176
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5177 := __e.Get(1)
_ = Z5177
tmp17600 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W516714940)
}
__typedArg0 := W516714940
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

W517814944 := tmp17600
_ = W517814944

tmp17601 := Call(__e, W516814941, Z5175)


tmp17602 := Call(__e, tmp17601, Z5176)


tmp17603 := Call(__e, tmp17602, Z5177)


__e.TailApply(tmp17603, W517814944)
return


}, 1))
return
}, 1))
return
}, 1)

W517414943 := tmp17599
_ = W517414943

tmp17658 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W517314942)
}
__typedArg0 := W517314942
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17658 {
tmp17604 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W517314942)
}
__typedArg0 := W517314942
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

W517914945 := tmp17604
_ = W517914945

tmp17605 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W517314942)
}
__typedArg0 := W517314942
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17606 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17605, V5156)


W518014946 := tmp17606
_ = W518014946

tmp17607 := MakeNative(func(__e *ControlFlow) {
Z5182 := __e.Get(1)
_ = Z5182
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5183 := __e.Get(1)
_ = Z5183
tmp17608 := Call(__e, W517414943, W517914945)


tmp17609 := Call(__e, tmp17608, Z5182)


__e.TailApply(tmp17609, Z5183)
return


}, 1))
return
}, 1)

W518114947 := tmp17607
_ = W518114947

tmp17642 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W518014946)
}
__typedArg0 := W518014946
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17642 {
tmp17610 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W518014946)
}
__typedArg0 := W518014946
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

W518414948 := tmp17610
_ = W518414948

tmp17611 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W518014946)
}
__typedArg0 := W518014946
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17612 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17611, V5156)


W518514949 := tmp17612
_ = W518514949

tmp17613 := MakeNative(func(__e *ControlFlow) {
Z5187 := __e.Get(1)
_ = Z5187
tmp17614 := Call(__e, W518114947, W518414948)


__e.TailApply(tmp17614, Z5187)
return


}, 1)

W518614950 := tmp17613
_ = W518614950

tmp17630 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W518514949)
}
__typedArg0 := W518514949
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17630 {
tmp17615 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W518514949)
}
__typedArg0 := W518514949
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

W518814951 := tmp17615
_ = W518814951

tmp17616 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W518514949)
}
__typedArg0 := W518514949
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17617 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17616, V5156)


W518914952 := tmp17617
_ = W518914952

tmp17618 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W518614950, W518814951)
return
}, 0)

W519014953 := tmp17618
_ = W519014953

tmp17622 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W518914952, Nil)
}
__typedArg0 := W518914952
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17622 {
__e.TailApply(PrimFunc(symthaw), W519014953)
return
} else {
tmp17620 := Call(__e, PrimFunc(symshen_4pvar_2), W518914952)


if True == tmp17620 {
__e.TailApply(PrimFunc(symshen_4bind_b), W518914952, Nil, V5156, W519014953)
return
} else {
__e.Return(False)
return
}


}


} else {
tmp17628 := Call(__e, PrimFunc(symshen_4pvar_2), W518514949)


if True == tmp17628 {
tmp17623 := Call(__e, PrimFunc(symshen_4newpv), V5156)


W519114954 := tmp17623
_ = W519114954

tmp17624 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W519114954, Nil)
}
__typedArg0 := W519114954
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17625 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W518614950, W519114954)
return
}, 0)

tmp17626 := Call(__e, PrimFunc(symshen_4bind_b), W518514949, tmp17624, V5156, tmp17625)


__e.TailApply(PrimFunc(symshen_4gc), V5156, tmp17626)
return


} else {
__e.Return(False)
return
}


}


} else {
tmp17640 := Call(__e, PrimFunc(symshen_4pvar_2), W518014946)


if True == tmp17640 {
tmp17631 := Call(__e, PrimFunc(symshen_4newpv), V5156)


W519214955 := tmp17631
_ = W519214955

tmp17632 := Call(__e, PrimFunc(symshen_4newpv), V5156)


W519314956 := tmp17632
_ = W519314956

tmp17633 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W519314956, Nil)
}
__typedArg0 := W519314956
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17634 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W519214955, tmp17633)
}
__typedArg0 := W519214955
__typedArg1 := tmp17633
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17635 := MakeNative(func(__e *ControlFlow) {
tmp17636 := Call(__e, W518114947, W519214955)


__e.TailApply(tmp17636, W519314956)
return


}, 0)

tmp17637 := Call(__e, PrimFunc(symshen_4bind_b), W518014946, tmp17634, V5156, tmp17635)


tmp17638 := Call(__e, PrimFunc(symshen_4gc), V5156, tmp17637)


__e.TailApply(PrimFunc(symshen_4gc), V5156, tmp17638)
return


} else {
__e.Return(False)
return
}


}


} else {
tmp17656 := Call(__e, PrimFunc(symshen_4pvar_2), W517314942)


if True == tmp17656 {
tmp17643 := Call(__e, PrimFunc(symshen_4newpv), V5156)


W519414957 := tmp17643
_ = W519414957

tmp17644 := Call(__e, PrimFunc(symshen_4newpv), V5156)


W519514958 := tmp17644
_ = W519514958

tmp17645 := Call(__e, PrimFunc(symshen_4newpv), V5156)


W519614959 := tmp17645
_ = W519614959

tmp17646 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W519614959, Nil)
}
__typedArg0 := W519614959
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17647 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W519514958, tmp17646)
}
__typedArg0 := W519514958
__typedArg1 := tmp17646
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17648 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W519414957, tmp17647)
}
__typedArg0 := W519414957
__typedArg1 := tmp17647
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17649 := MakeNative(func(__e *ControlFlow) {
tmp17650 := Call(__e, W517414943, W519414957)


tmp17651 := Call(__e, tmp17650, W519514958)


__e.TailApply(tmp17651, W519614959)
return


}, 0)

tmp17652 := Call(__e, PrimFunc(symshen_4bind_b), W517314942, tmp17648, V5156, tmp17649)


tmp17653 := Call(__e, PrimFunc(symshen_4gc), V5156, tmp17652)


tmp17654 := Call(__e, PrimFunc(symshen_4gc), V5156, tmp17653)


__e.TailApply(PrimFunc(symshen_4gc), V5156, tmp17654)
return


} else {
__e.Return(False)
return
}


}


} else {
tmp17676 := Call(__e, PrimFunc(symshen_4pvar_2), W516714940)


if True == tmp17676 {
tmp17659 := Call(__e, PrimFunc(symshen_4newpv), V5156)


W519714960 := tmp17659
_ = W519714960

tmp17660 := Call(__e, PrimFunc(symshen_4newpv), V5156)


W519814961 := tmp17660
_ = W519814961

tmp17661 := Call(__e, PrimFunc(symshen_4newpv), V5156)


W519914962 := tmp17661
_ = W519914962

tmp17662 := Call(__e, PrimFunc(symshen_4newpv), V5156)


W520014963 := tmp17662
_ = W520014963

tmp17663 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W519914962, Nil)
}
__typedArg0 := W519914962
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17664 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W519814961, tmp17663)
}
__typedArg0 := W519814961
__typedArg1 := tmp17663
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17665 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W519714960, tmp17664)
}
__typedArg0 := W519714960
__typedArg1 := tmp17664
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17666 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17665, W520014963)
}
__typedArg0 := tmp17665
__typedArg1 := W520014963
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17667 := MakeNative(func(__e *ControlFlow) {
tmp17668 := Call(__e, W516814941, W519714960)


tmp17669 := Call(__e, tmp17668, W519814961)


tmp17670 := Call(__e, tmp17669, W519914962)


__e.TailApply(tmp17670, W520014963)
return


}, 0)

tmp17671 := Call(__e, PrimFunc(symshen_4bind_b), W516714940, tmp17666, V5156, tmp17667)


tmp17672 := Call(__e, PrimFunc(symshen_4gc), V5156, tmp17671)


tmp17673 := Call(__e, PrimFunc(symshen_4gc), V5156, tmp17672)


tmp17674 := Call(__e, PrimFunc(symshen_4gc), V5156, tmp17673)


__e.TailApply(PrimFunc(symshen_4gc), V5156, tmp17674)
return


} else {
__e.Return(False)
return
}


}


} else {
__e.Return(False)
return
}


} else {
__e.Return(False)
return
}


} else {
__e.Return(W516014933)
return
}


}, 6)

tmp17685 := Call(__e, ns2_1set, symshen_4p_1hyps, tmp17573)


_ = tmp17685

tmp17686 := MakeNative(func(__e *ControlFlow) {
V5201 := __e.Get(1)
_ = V5201
V5202 := __e.Get(2)
_ = V5202
V5203 := __e.Get(3)
_ = V5203
V5204 := __e.Get(4)
_ = V5204
V5205 := __e.Get(5)
_ = V5205
V5206 := __e.Get(6)
_ = V5206
V5207 := __e.Get(7)
_ = V5207
W520814964 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V5206)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V5206
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
_ = W520814964

tmp17724 := Call(__e, PrimFunc(symshen_4unlocked_2), V5205)


var ifres17688 Obj

if True == tmp17724 {
tmp17689 := Call(__e, PrimFunc(symshen_4lazyderef), V5201, V5204)


W521014966 := tmp17689
_ = W521014966

tmp17723 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W521014966)
}
__typedArg0 := W521014966
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17690 Obj

if True == tmp17723 {
tmp17691 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W521014966)
}
__typedArg0 := W521014966
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17692 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17691, V5204)


W521114967 := tmp17692
_ = W521114967

tmp17722 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W521114967, symwhere)
}
__typedArg0 := W521114967
__typedArg1 := symwhere
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17693 Obj

if True == tmp17722 {
tmp17694 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W521014966)
}
__typedArg0 := W521014966
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17695 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17694, V5204)


W521214968 := tmp17695
_ = W521214968

tmp17721 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W521214968)
}
__typedArg0 := W521214968
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17696 Obj

if True == tmp17721 {
tmp17697 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W521214968)
}
__typedArg0 := W521214968
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

W521314969 := tmp17697
_ = W521314969

tmp17698 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W521214968)
}
__typedArg0 := W521214968
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17699 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17698, V5204)


W521414970 := tmp17699
_ = W521414970

tmp17720 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W521414970)
}
__typedArg0 := W521414970
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17700 Obj

if True == tmp17720 {
tmp17701 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W521414970)
}
__typedArg0 := W521414970
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

W521514971 := tmp17701
_ = W521514971

tmp17702 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W521414970)
}
__typedArg0 := W521414970
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17703 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17702, V5204)


W521614972 := tmp17703
_ = W521614972

tmp17719 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W521614972, Nil)
}
__typedArg0 := W521614972
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17704 Obj

if True == tmp17719 {
tmp17705 := Call(__e, PrimFunc(symshen_4newpv), V5204)


W521714973 := tmp17705
_ = W521714973

tmp17706 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp17706

tmp17707 := MakeNative(func(__e *ControlFlow) {
tmp17708 := Call(__e, PrimFunc(symshen_4curry), W521314969)


tmp17709 := MakeNative(func(__e *ControlFlow) {
tmp17710 := MakeNative(func(__e *ControlFlow) {
tmp17711 := MakeNative(func(__e *ControlFlow) {
tmp17712 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp17713 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symverified, Nil)
}
__typedArg0 := symverified
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17714 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17712, tmp17713)
}
__typedArg0 := tmp17712
__typedArg1 := tmp17713
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17715 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W521714973, tmp17714)
}
__typedArg0 := W521714973
__typedArg1 := tmp17714
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17716 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17715, V5203)
}
__typedArg0 := tmp17715
__typedArg1 := V5203
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4t_d_1correct), W521514971, V5202, tmp17716, V5204, V5205, W520814964, V5207)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V5204, V5205, W520814964, tmp17711)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), W521714973, symboolean, V5203, V5204, V5205, W520814964, tmp17710)
return


}, 0)

__e.TailApply(PrimFunc(symbind), W521714973, tmp17708, V5204, V5205, W520814964, tmp17709)
return


}, 0)

tmp17717 := Call(__e, PrimFunc(symshen_4cut), V5204, V5205, W520814964, tmp17707)


tmp17718 := Call(__e, PrimFunc(symshen_4gc), V5204, tmp17717)


ifres17704 = tmp17718


} else {
ifres17704 = False


}

ifres17700 = ifres17704


} else {
ifres17700 = False


}

ifres17696 = ifres17700


} else {
ifres17696 = False


}

ifres17693 = ifres17696


} else {
ifres17693 = False


}

ifres17690 = ifres17693


} else {
ifres17690 = False


}

ifres17688 = ifres17690


} else {
ifres17688 = False


}

W520914965 := ifres17688
_ = W520914965

tmp17733 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W520914965, False)
}
__typedArg0 := W520914965
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17733 {
tmp17729 := Call(__e, PrimFunc(symshen_4unlocked_2), V5205)


var ifres17725 Obj

if True == tmp17729 {
tmp17726 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp17726

tmp17727 := Call(__e, PrimFunc(symshen_4curry), V5201)


tmp17728 := Call(__e, PrimFunc(symshen_4system_1S_1h), tmp17727, V5202, V5203, V5204, V5205, W520814964, V5207)


ifres17725 = tmp17728


} else {
ifres17725 = False


}

W521814974 := ifres17725
_ = W521814974

tmp17731 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W521814974, False)
}
__typedArg0 := W521814974
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17731 {
__e.TailApply(PrimFunc(symshen_4unlock), V5205, W520814964)
return
} else {
__e.Return(W521814974)
return
}


} else {
__e.Return(W520914965)
return
}


}, 7)

tmp17734 := Call(__e, ns2_1set, symshen_4t_d_1correct, tmp17686)


_ = tmp17734

tmp17735 := MakeNative(func(__e *ControlFlow) {
V5219 := __e.Get(1)
_ = V5219
V5220 := __e.Get(2)
_ = V5220
V5221 := __e.Get(3)
_ = V5221
V5222 := __e.Get(4)
_ = V5222
V5223 := __e.Get(5)
_ = V5223
V5224 := __e.Get(6)
_ = V5224
V5225 := __e.Get(7)
_ = V5225
V5226 := __e.Get(8)
_ = V5226
tmp17742 := Call(__e, PrimFunc(symshen_4unlocked_2), V5224)


var ifres17736 Obj

if True == tmp17742 {
tmp17737 := Call(__e, PrimFunc(symshen_4lazyderef), V5219, V5223)


W522814976 := tmp17737
_ = W522814976

tmp17741 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W522814976, Nil)
}
__typedArg0 := W522814976
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17738 Obj

if True == tmp17741 {
tmp17739 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp17739

tmp17740 := Call(__e, PrimFunc(symis_b), V5220, V5222, V5223, V5224, V5225, V5226)


ifres17738 = tmp17740


} else {
ifres17738 = False


}

ifres17736 = ifres17738


} else {
ifres17736 = False


}

W522714975 := ifres17736
_ = W522714975

tmp17774 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W522714975, False)
}
__typedArg0 := W522714975
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17774 {
tmp17772 := Call(__e, PrimFunc(symshen_4unlocked_2), V5224)


if True == tmp17772 {
tmp17743 := Call(__e, PrimFunc(symshen_4lazyderef), V5219, V5223)


W522914977 := tmp17743
_ = W522914977

tmp17770 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W522914977)
}
__typedArg0 := W522914977
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17770 {
tmp17744 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W522914977)
}
__typedArg0 := W522914977
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

W523014978 := tmp17744
_ = W523014978

tmp17745 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W522914977)
}
__typedArg0 := W522914977
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

W523114979 := tmp17745
_ = W523114979

tmp17746 := Call(__e, PrimFunc(symshen_4lazyderef), V5220, V5223)


W523214980 := tmp17746
_ = W523214980

tmp17768 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W523214980)
}
__typedArg0 := W523214980
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17768 {
tmp17747 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W523214980)
}
__typedArg0 := W523214980
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

W523314981 := tmp17747
_ = W523314981

tmp17748 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W523214980)
}
__typedArg0 := W523214980
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17749 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17748, V5223)


W523414982 := tmp17749
_ = W523414982

tmp17766 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W523414982)
}
__typedArg0 := W523414982
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17766 {
tmp17750 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W523414982)
}
__typedArg0 := W523414982
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17751 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17750, V5223)


W523514983 := tmp17751
_ = W523514983

tmp17764 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W523514983, sym_1_1_6)
}
__typedArg0 := W523514983
__typedArg1 := sym_1_1_6
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17764 {
tmp17752 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W523414982)
}
__typedArg0 := W523414982
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17753 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17752, V5223)


W523614984 := tmp17753
_ = W523614984

tmp17762 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W523614984)
}
__typedArg0 := W523614984
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17762 {
tmp17754 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W523614984)
}
__typedArg0 := W523614984
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

W523714985 := tmp17754
_ = W523714985

tmp17755 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W523614984)
}
__typedArg0 := W523614984
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17756 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17755, V5223)


W523814986 := tmp17756
_ = W523814986

tmp17760 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W523814986, Nil)
}
__typedArg0 := W523814986
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17760 {
tmp17757 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp17757

tmp17758 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4t_d_1integrity), W523114979, W523714985, V5221, V5222, V5223, V5224, V5225, V5226)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), W523014978, W523314981, V5221, V5223, V5224, V5225, tmp17758)
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


} else {
__e.Return(False)
return
}


} else {
__e.Return(W522714975)
return
}


}, 8)

tmp17775 := Call(__e, ns2_1set, symshen_4t_d_1integrity, tmp17735)


_ = tmp17775

tmp17776 := MakeNative(func(__e *ControlFlow) {
V5239 := __e.Get(1)
_ = V5239
tmp17785 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symabsvector_2) {
return PrimIsVector(V5239)
}
__typedArg0 := V5239
return Call(__e, PrimFunc(symabsvector_2), __typedArg0)
})()

if True == tmp17785 {
tmp17782 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_2) {
return PrimIsString(V5239)
}
__typedArg0 := V5239
return Call(__e, PrimFunc(symstring_2), __typedArg0)
})()

tmp17783 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp17782)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp17782
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres17778 Obj

if True == tmp17783 {
tmp17780 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V5239, MakeNumber(0))
}
__typedArg0 := V5239
__typedArg1 := MakeNumber(0)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})()

tmp17781 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp17780, symshen_4print_1freshterm)
}
__typedArg0 := tmp17780
__typedArg1 := symshen_4print_1freshterm
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17779 Obj

if True == tmp17781 {
ifres17779 = True


} else {
ifres17779 = False


}

ifres17778 = ifres17779


} else {
ifres17778 = False


}

if True == ifres17778 {
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

__e.TailApply(ns2_1set, symshen_4freshterm_2, tmp17776)
return




}, 0)

